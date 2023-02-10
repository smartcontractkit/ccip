package load

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/chainlink-testing-framework/client"
	"github.com/smartcontractkit/chainlink-testing-framework/testreporters"
	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/router"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/testhelpers"
	"github.com/smartcontractkit/chainlink/core/utils"
	bigmath "github.com/smartcontractkit/chainlink/core/utils/big_math"
	"github.com/smartcontractkit/chainlink/integration-tests/actions"
)

type phase string
type status string

const (
	E2E                     phase  = "v)OverallCommitAndExecution"
	TX                      phase  = "i)SendTxBlockConfirmation"
	CCIPSendRe              phase  = "ii)CCIPSendRequested Event"
	SeqNumAndRepAccIncrease phase  = "iii)ReportAcceptedByCommitStore(Commit)"
	ExecStateChanged        phase  = "iv)ExecutionStateChanged Event(Execution)"
	success                 status = "✅"
	fail                    status = "❌"
	TokenTransfer           string = "WithToken"
	DataOnlyTransfer        string = "WithoutToken"
)

type CCIPE2ELoad struct {
	t                     *testing.T
	Source                *actions.SourceCCIPModule // all source contracts
	Destination           *actions.DestCCIPModule   // all destination contracts
	NoOfReq               int64                     // no of Request fired - required for balance assertion at the end
	totalGEFee            *big.Int
	BalanceStats          BalanceStats  // balance assertion details
	CurrentMsgSerialNo    *atomic.Int64 // current msg serial number in the load sequence
	InitialSourceBlockNum uint64
	InitialDestBlockNum   uint64 // blocknumber before the first message is fired in the load sequence
	sentMsgMu             *sync.Mutex
	SentMsg               map[uint64]evm_2_evm_onramp.InternalEVM2EVMMessage // track the messages by seq num for debugging purpose
	CallTimeOut           time.Duration                                      // max time to wait for various on-chain events
	TickerDuration        time.Duration                                      // poll frequency while waiting for on-chain events
	callStatsMu           *sync.Mutex
	callStats             map[int64]map[phase]StatParams // keeps track of various phase related metrics
	seqNumCommittedMu     *sync.Mutex
	seqNumCommitted       map[uint64]uint64 // key : seqNumber in the ReportAccepted event, value : blocknumber for corresponding event
	msg                   router.ConsumerEVM2AnyMessage
}

type StatParams struct {
	SeqNum    string  `json:"SequenceNumber,omitempty"`
	Duration  float64 `json:"duration,omitempty"`
	ReqStatus status  `json:"success"`
}

type AvgStatParams struct {
	Duration float64 `json:"averageDurationForSuccessfulRequests,omitempty"`
	Failed   int     `json:"failedCount,omitempty"`
	Success  int     `json:"successCount"`
}

type BalanceStats struct {
	SourceBalanceReq        map[string]*big.Int
	SourceBalanceAssertions []testhelpers.BalanceAssertion
	DestBalanceReq          map[string]*big.Int
	DestBalanceAssertions   []testhelpers.BalanceAssertion
}

type JsonStats struct {
	MsgSerialNumber int64                `json:"msgSerialNumber"`
	Stats           map[phase]StatParams `json:"phaseDetails"`
}

type SlackStats struct {
	AvgE2EDuration                  float64 `json:"avgOverallCommitAndExecution,omitempty"`
	AvgCommitDuration               float64 `json:"avgCommit,omitempty"`
	AvgExecDuration                 float64 `json:"avgExecution,omitempty"`
	LongestE2EDuration              float64 `json:"longestOverallCommitAndExecution,omitempty"`
	LongestCommitDuration           float64 `json:"longestCommit,omitempty"`
	LongestExecDuration             float64 `json:"longestExecution,omitempty"`
	FastestE2EDuration              float64 `json:"fastestOverallCommitAndExecution,omitempty"`
	FastestCommitDuration           float64 `json:"fastestCommit,omitempty"`
	FastestExecDuration             float64 `json:"fastestExecution,omitempty"`
	TotalNumberOfFailedRequests     int     `json:"totalFailedRequests,omitempty"`
	TotalNumberOfSuccessfulRequests int     `json:"totalSuccessfulRequests,omitempty"`
	FailedCommit                    int     `json:"noOfFailedCommit,omitempty"`
	FailedExecution                 int     `json:"noOfFailedExecution,omitempty"`
	FailedSendTransaction           int     `json:"noOfFailedSendTransaction,omitempty"`
	FailedCCIPSendRequested         int     `json:"noOfFailedCCIPSendRequested,omitempty"`
}

func NewCCIPLoad(t *testing.T, source *actions.SourceCCIPModule, dest *actions.DestCCIPModule, timeout time.Duration, noOfReq int64) *CCIPE2ELoad {
	return &CCIPE2ELoad{
		t:                  t,
		Source:             source,
		Destination:        dest,
		CurrentMsgSerialNo: atomic.NewInt64(1),
		sentMsgMu:          &sync.Mutex{},
		SentMsg:            make(map[uint64]evm_2_evm_onramp.InternalEVM2EVMMessage),
		TickerDuration:     time.Second,
		CallTimeOut:        timeout,
		NoOfReq:            noOfReq,
		callStats:          make(map[int64]map[phase]StatParams),
		callStatsMu:        &sync.Mutex{},
		seqNumCommittedMu:  &sync.Mutex{},
		seqNumCommitted:    make(map[uint64]uint64),
	}
}

// BeforeAllCall funds subscription, approves the token transfer amount.
// Needs to be called before load sequence is started.
// Needs to approve and fund for the entire sequence.
func (c *CCIPE2ELoad) BeforeAllCall() {
	sourceCCIP := c.Source
	destCCIP := c.Destination
	var tokenAndAmounts []router.CommonEVMTokenAndAmount
	for i, token := range sourceCCIP.Common.BridgeTokens {
		tokenAndAmounts = append(tokenAndAmounts, router.CommonEVMTokenAndAmount{
			Token: common.HexToAddress(token.Address()), Amount: c.Source.TransferAmount[i],
		})
		// approve the onramp router so that it caninitiate transferring the token

		err := token.Approve(c.Source.Common.Router.Address(), bigmath.Mul(c.Source.TransferAmount[i], big.NewInt(c.NoOfReq)))
		require.NoError(c.t, err, "Could not approve permissions for the onRamp router "+
			"on the source link token contract")
	}

	err := sourceCCIP.Common.ChainClient.WaitForEvents()
	require.NoError(c.t, err, "Failed to wait for events")

	// save the current block numbers to use in various filter log requests
	currentBlockOnSource, err := sourceCCIP.Common.ChainClient.LatestBlockNumber(context.Background())
	require.NoError(c.t, err, "failed to fetch latest source block num")
	currentBlockOnDest, err := destCCIP.Common.ChainClient.LatestBlockNumber(context.Background())
	require.NoError(c.t, err, "failed to fetch latest dest block num")
	c.InitialDestBlockNum = currentBlockOnDest
	c.InitialSourceBlockNum = currentBlockOnSource
	// collect the balance requirement to verify balances after transfer
	sourceBalances, err := testhelpers.GetBalances(sourceCCIP.CollectBalanceRequirements(c.t))
	require.NoError(c.t, err, "fetching source balance")
	destBalances, err := testhelpers.GetBalances(destCCIP.CollectBalanceRequirements(c.t))
	require.NoError(c.t, err, "fetching dest balance")
	c.BalanceStats = BalanceStats{
		SourceBalanceReq: sourceBalances,
		DestBalanceReq:   destBalances,
	}
	extraArgsV1, err := testhelpers.GetEVMExtraArgsV1(big.NewInt(100_000), false)
	require.NoError(c.t, err, "Failed encoding the options field")

	receiver, err := utils.ABIEncode(`[{"type":"address"}]`, destCCIP.ReceiverDapp.EthAddress)
	require.NoError(c.t, err, "Failed encoding the receiver address")
	c.msg = router.ConsumerEVM2AnyMessage{
		Receiver:  receiver,
		ExtraArgs: extraArgsV1,
		FeeToken:  common.HexToAddress(sourceCCIP.Common.FeeToken.Address()),
	}
	// calculate approx fee
	fee, err := sourceCCIP.Common.Router.GetFee(sourceCCIP.DestinationChainId, c.msg)
	require.NoError(c.t, err)
	// Approve sufficient fee amount
	err = sourceCCIP.Common.FeeToken.Approve(sourceCCIP.Common.Router.Address(), bigmath.Mul(fee, big.NewInt(c.NoOfReq)))
	require.NoError(c.t, err)
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	require.NoError(c.t, err, "Failed to wait for events")

	sourceCCIP.Common.ChainClient.ParallelTransactions(false)
	destCCIP.Common.ChainClient.ParallelTransactions(false)
}

func (c *CCIPE2ELoad) AfterAllCall() {
	c.BalanceStats.DestBalanceAssertions = c.Destination.BalanceAssertions(
		c.t,
		c.BalanceStats.DestBalanceReq,
		c.Source.TransferAmount,
		c.NoOfReq,
	)
	c.BalanceStats.SourceBalanceAssertions = c.Source.BalanceAssertions(c.t, c.BalanceStats.SourceBalanceReq, c.NoOfReq, c.totalGEFee)
	actions.AssertBalances(c.t, c.BalanceStats.DestBalanceAssertions)
	actions.AssertBalances(c.t, c.BalanceStats.SourceBalanceAssertions)
}

func (c *CCIPE2ELoad) Call(msgType interface{}) client.CallResult {
	var res client.CallResult
	sourceCCIP := c.Source
	destCCIP := c.Destination
	msgSerialNo := c.CurrentMsgSerialNo.Load()
	c.CurrentMsgSerialNo.Inc()
	var tokenAndAmounts []router.CommonEVMTokenAndAmount
	for i, token := range sourceCCIP.Common.BridgeTokens {
		tokenAndAmounts = append(tokenAndAmounts, router.CommonEVMTokenAndAmount{
			Token: common.HexToAddress(token.Address()), Amount: c.Source.TransferAmount[i],
		})
	}

	// form the message for transfer
	msgStr := fmt.Sprintf("message with Id %d", msgSerialNo)
	c.msg.Data = []byte(msgStr)

	if msgType == TokenTransfer {
		c.msg.TokensAndAmounts = tokenAndAmounts
	}

	startTime := time.Now()
	// initiate the transfer
	log.Debug().Int("msg Number", int(msgSerialNo)).Str("triggeredAt", time.Now().GoString()).Msg("triggering transfer")
	sendTx, err := sourceCCIP.Common.Router.CCIPSend(destCCIP.Common.ChainClient.GetChainID().Uint64(), c.msg)

	if err != nil {
		c.updatestats(msgSerialNo, "", TX, time.Since(startTime), fail)
		res.Error = fmt.Errorf("CCIPSend request error %v for msg ID %d", err, msgSerialNo)
		return res
	}
	c.updatestats(msgSerialNo, "", TX, time.Since(startTime), success)

	// wait for
	// - CCIPSendRequested Event log to be generated,
	ticker := time.NewTicker(c.TickerDuration)
	defer ticker.Stop()
	sentMsg, err := c.waitForCCIPSendRequested(ticker, c.InitialSourceBlockNum, msgSerialNo, sendTx.Hash().Hex(), time.Now())
	if err != nil {
		res.Error = err
		return res
	}
	commitStartTime := time.Now()
	seqNum := sentMsg.SequenceNumber
	messageID := sentMsg.MessageId
	if bytes.Compare(sentMsg.Data, []byte(msgStr)) == 0 {
		c.updateSentMsgQueue(seqNum, sentMsg)
	} else {
		res.Error = fmt.Errorf("the message byte didnot match expected %s received %s msg ID %d", msgStr, string(sentMsg.Data), msgSerialNo)
		return res
	}

	// wait for
	// - CommitStore to increase the seq number,
	err = c.waitForSeqNumberIncrease(ticker, seqNum, msgSerialNo, commitStartTime)
	if err != nil {
		res.Error = err
		return res
	}
	// wait for ReportAccepted event
	err = c.waitForReportAccepted(ticker, msgSerialNo, seqNum, c.InitialDestBlockNum, commitStartTime)
	if err != nil {
		res.Error = err
		return res
	}

	// wait for ExecutionStateChanged event
	err = c.waitForExecStateChange(ticker, []uint64{seqNum}, [][32]byte{messageID}, c.seqNumCommitted[seqNum]-2, msgSerialNo, time.Now())
	if err != nil {
		res.Error = err
		return res
	}
	c.updatestats(msgSerialNo, fmt.Sprint(seqNum), E2E, time.Since(commitStartTime), success)
	res.Error = nil
	res.Data = c.SentMsg[seqNum]
	return res
}

func (c *CCIPE2ELoad) updateSentMsgQueue(seqNum uint64, sentMsg evm_2_evm_onramp.InternalEVM2EVMMessage) {
	c.sentMsgMu.Lock()
	defer c.sentMsgMu.Unlock()
	c.SentMsg[seqNum] = sentMsg
}

func (c *CCIPE2ELoad) updateSeqNumCommitted(seqNum []uint64, blockNum uint64) {
	c.seqNumCommittedMu.Lock()
	defer c.seqNumCommittedMu.Unlock()
	for _, num := range seqNum {
		if _, ok := c.seqNumCommitted[num]; ok {
			return
		}
		c.seqNumCommitted[num] = blockNum
	}
}

func (c *CCIPE2ELoad) updatestats(msgSerialNo int64, seqNum string, step phase, duration time.Duration, state status) {
	c.callStatsMu.Lock()
	defer c.callStatsMu.Unlock()
	if _, ok := c.callStats[msgSerialNo]; !ok {
		c.callStats[msgSerialNo] = make(map[phase]StatParams)
		c.callStats[msgSerialNo][step] = StatParams{
			SeqNum:    seqNum,
			Duration:  duration.Seconds(),
			ReqStatus: state,
		}
	} else {
		c.callStats[msgSerialNo][step] = StatParams{
			SeqNum:    seqNum,
			Duration:  c.callStats[msgSerialNo][step].Duration + duration.Seconds(),
			ReqStatus: state,
		}
	}
	// if any of the phase fails mark the E2E as failed
	if state == fail {
		c.callStats[msgSerialNo][E2E] = StatParams{
			SeqNum:    seqNum,
			ReqStatus: state,
		}
	} else {
		log.Info().Str(fmt.Sprint(step), fmt.Sprint(state)).Msgf("seq num %s", seqNum)
	}
}

func (c *CCIPE2ELoad) PrintStats(rps int, duration float64) {
	if _, err := os.Stat("./logs/stats"); os.IsNotExist(err) {
		os.MkdirAll("./logs/stats", 0700)
	}
	tempFile, err := os.Create("./logs/stats/CCIPLoad_complete.json")
	defer tempFile.Close()
	require.NoError(c.t, err, "creating stat file")

	tempStatFile, err := os.Create("./logs/stats/CCIPLoad_avg.json")
	defer tempStatFile.Close()
	require.NoError(c.t, err, "creating stat file")

	slackFile, err := os.Create("./logs/payload-slack-content.json")
	defer slackFile.Close()
	require.NoError(c.t, err, "creating stat file")

	log.Info().Msg("Msg Stats")
	keys := make([]int64, 0, len(c.callStats))
	for msgSerialNo := range c.callStats {
		keys = append(keys, msgSerialNo)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	durationSumMap := make(map[phase]float64)
	durationMap := make(map[phase][]float64)
	successCount := make(map[phase]int)
	failureCount := make(map[phase]int)
	var jsonstats []JsonStats

	for _, msgSerialNo := range keys {
		stats := c.callStats[msgSerialNo]
		event := log.Info()
		jsonstats = append(jsonstats, JsonStats{
			MsgSerialNumber: msgSerialNo,
			Stats:           stats,
		})
		for step, stat := range stats {
			event.Float64(fmt.Sprintf("%s%s", step, stat.ReqStatus), stat.Duration)
			if stat.ReqStatus == success {
				durationSumMap[step] += stat.Duration
				if _, ok := durationMap[step]; !ok {
					durationMap[step] = []float64{stat.Duration}
				} else {
					durationMap[step] = append(durationMap[step], stat.Duration)
				}
				successCount[step] += 1
			} else {
				failureCount[step] += 1
			}
		}
		event.Msgf("Msg stats for msg Id %d", msgSerialNo)
	}
	event := log.Info()
	overallStats := make(map[phase]AvgStatParams)
	for step, d := range durationSumMap {
		avg := d / float64(successCount[step])
		overallStats[step] = AvgStatParams{
			Duration: avg,
			Failed:   failureCount[step],
			Success:  successCount[step],
		}
		event.Float64(string(step), avg)
	}
	// if all the requests fail for a particular step
	for step, f := range failureCount {
		if _, ok := durationSumMap[step]; !ok {
			overallStats[step] = AvgStatParams{
				Failed: f,
			}
		}
	}
	for step, _ := range durationMap {
		sort.Slice(durationMap[step], func(i, j int) bool {
			return durationMap[step][i] < durationMap[step][j]
		})
	}
	slackStats := SlackStats{
		AvgE2EDuration:                  overallStats[E2E].Duration,
		AvgCommitDuration:               overallStats[SeqNumAndRepAccIncrease].Duration,
		AvgExecDuration:                 overallStats[ExecStateChanged].Duration,
		LongestE2EDuration:              durationMap[E2E][len(durationMap[E2E])-1],
		LongestCommitDuration:           durationMap[SeqNumAndRepAccIncrease][len(durationMap[SeqNumAndRepAccIncrease])-1],
		LongestExecDuration:             durationMap[ExecStateChanged][len(durationMap[ExecStateChanged])-1],
		FastestE2EDuration:              durationMap[E2E][0],
		FastestCommitDuration:           durationMap[SeqNumAndRepAccIncrease][0],
		FastestExecDuration:             durationMap[ExecStateChanged][0],
		TotalNumberOfFailedRequests:     failureCount[E2E],
		TotalNumberOfSuccessfulRequests: successCount[E2E],
		FailedCommit:                    failureCount[SeqNumAndRepAccIncrease],
		FailedExecution:                 failureCount[ExecStateChanged],
		FailedSendTransaction:           failureCount[TX],
		FailedCCIPSendRequested:         failureCount[CCIPSendRe],
	}

	stats, err := json.MarshalIndent(overallStats, "", "  ")
	require.NoError(c.t, err, "marshal overallStats")
	_, err = tempStatFile.Write(stats)
	require.NoError(c.t, err, "writing overallStats")

	stats, err = json.MarshalIndent(jsonstats, "", "  ")
	require.NoError(c.t, err, "marshal avg stats")
	_, err = tempFile.Write(stats)
	require.NoError(c.t, err, "writing avg stats")

	headerText := ":white_check_mark: CCIP Load Test PASSED :white_check_mark:"
	if c.t.Failed() {
		headerText = ":x: CCIP Load Test FAILED :x:"
	}
	stats, err = json.MarshalIndent(slackStats, "", "  ")
	require.NoError(c.t, err)

	runUrl := os.Getenv("GH_RUN_URL")
	if runUrl != "" {
		testreporters.SlackNotifyBlocks(headerText, []string{fmt.Sprintf(
			"Load sequence ran for %.0fm sending a total of %d transactions at a rate of %d tx(s) per second",
			duration, successCount[E2E]+failureCount[E2E], rps),
			fmt.Sprintf("<%s|Detailed Run Results are available in artifacts>", runUrl),
			"\nLoad Run Summary:",
			string(stats)}, slackFile)
	}

	event.Int("No of Successful Requests", successCount[E2E])
	event.Msgf("Average Duration for successful requests")
	log.Info().Msg("Commit Report stats")
	it, err := c.Destination.CommitStore.FilterReportAccepted(c.InitialDestBlockNum)
	require.NoError(c.t, err, "report committed result")
	i := 1
	event = log.Info()
	for it.Next() {
		event.Interface(fmt.Sprintf("%d Report Intervals", i), it.Event.Report.Intervals)
		i++
	}
	event.Msgf("CommitStore-Reports Accepted")
}

func (c *CCIPE2ELoad) waitForExecStateChange(ticker *time.Ticker, seqNums []uint64, messageID [][32]byte, currentBlockOnDest uint64, msgSerialNo int64, timeNow time.Time) error {
	log.Info().Int("msg Number", int(msgSerialNo)).Msgf(
		"waiting for ExecutionStateChanged for seqNums %v", seqNums)
	ctx, cancel := context.WithTimeout(context.Background(), c.CallTimeOut)
	defer cancel()
	for {
		select {
		case <-ticker.C:
			iterator, err := c.Destination.OffRamp.FilterExecutionStateChanged(seqNums, messageID, currentBlockOnDest)
			if err != nil {
				for _, seqNum := range seqNums {
					c.updatestats(msgSerialNo, fmt.Sprint(seqNum), ExecStateChanged, time.Since(timeNow), fail)
				}
				return fmt.Errorf("filtering event ExecutionStateChanged returned error %v msg ID %d and seqNum %v", err, msgSerialNo, seqNums)
			}
			for iterator.Next() {
				switch ccip.MessageExecutionState(iterator.Event.State) {
				case ccip.Success:
					for _, seqNum := range seqNums {
						c.updatestats(msgSerialNo, fmt.Sprint(seqNum), ExecStateChanged, time.Since(timeNow), success)
					}
					return nil
				case ccip.Failure:
					for _, seqNum := range seqNums {
						c.updatestats(msgSerialNo, fmt.Sprint(seqNum), ExecStateChanged, time.Since(timeNow), fail)
					}
					return fmt.Errorf("ExecutionStateChanged event returned failure for seq num %v msg ID %d", seqNums, msgSerialNo)
				}
			}
		case <-ctx.Done():
			for _, seqNum := range seqNums {
				c.updatestats(msgSerialNo, fmt.Sprint(seqNum), ExecStateChanged, time.Since(timeNow), fail)
			}
			return fmt.Errorf("ExecutionStateChanged event not found for seq num %v msg ID %d", seqNums, msgSerialNo)
		}
	}
}

func (c *CCIPE2ELoad) waitForSeqNumberIncrease(ticker *time.Ticker, seqNum uint64, msgSerialNo int64, timeNow time.Time) error {
	log.Info().Int("msg Number", int(msgSerialNo)).Msgf("waiting for seq number %d to get increased", int(seqNum))
	ctx, cancel := context.WithTimeout(context.Background(), c.CallTimeOut)
	defer cancel()
	for {
		select {
		case <-ticker.C:
			seqNumberAfter, err := c.Destination.CommitStore.GetNextSeqNumber()
			if err != nil {
				c.updatestats(msgSerialNo, fmt.Sprint(seqNum), SeqNumAndRepAccIncrease, time.Since(timeNow), fail)
				return fmt.Errorf("error %v in GetNextExpectedSeqNumber by commitStore for msg ID %d", err, msgSerialNo)
			}
			if seqNumberAfter > seqNum {
				return nil
			}
		case <-ctx.Done():
			c.updatestats(msgSerialNo, fmt.Sprint(seqNum), SeqNumAndRepAccIncrease, time.Since(timeNow), fail)
			return fmt.Errorf("sequence number is not increased for seq num %d msg ID %d", seqNum, msgSerialNo)
		}
	}
}

func (c *CCIPE2ELoad) waitForReportAccepted(ticker *time.Ticker, msgSerialNo int64, seqNum uint64, currentBlockOnDest uint64, timeNow time.Time) error {
	log.Info().Int("seq Number", int(seqNum)).Msg("waiting for ReportAccepted")
	ctx, cancel := context.WithTimeout(context.Background(), c.CallTimeOut)
	defer cancel()
	for {
		select {
		case <-ticker.C:
			// skip calling FilterReportAccepted if the seqNum is present in the map
			if _, ok := c.seqNumCommitted[seqNum]; ok {
				c.updatestats(msgSerialNo, fmt.Sprint(seqNum), SeqNumAndRepAccIncrease, time.Since(timeNow), success)
				return nil
			}
			it, err := c.Destination.CommitStore.FilterReportAccepted(currentBlockOnDest)
			if err != nil {
				c.updatestats(msgSerialNo, fmt.Sprint(seqNum), SeqNumAndRepAccIncrease, time.Since(timeNow), fail)
				return fmt.Errorf("error %v in filtering by ReportAccepted event for seq num %d", err, seqNum)
			}
			for it.Next() {
				for _, in := range it.Event.Report.Intervals {
					seqNums := make([]uint64, in.Max-in.Min+1)
					var i uint64
					for range seqNums {
						seqNums[i] = in.Min + i
						i++
					}
					// update SeqNumCommitted map for all seqNums in the emitted ReportAccepted event
					c.updateSeqNumCommitted(seqNums, it.Event.Raw.BlockNumber)
					if in.Max >= seqNum && in.Min <= seqNum {
						c.updatestats(msgSerialNo, fmt.Sprint(seqNum), SeqNumAndRepAccIncrease, time.Since(timeNow), success)
						return nil
					}
				}
			}
		case <-ctx.Done():
			c.updatestats(msgSerialNo, fmt.Sprint(seqNum), SeqNumAndRepAccIncrease, time.Since(timeNow), fail)
			return fmt.Errorf("ReportAccepted is not found for seq num %d", seqNum)
		}
	}
}

func (c *CCIPE2ELoad) waitForCCIPSendRequested(
	ticker *time.Ticker,
	currentBlockOnSource uint64,
	msgSerialNo int64,
	txHash string,
	timeNow time.Time,
) (evm_2_evm_onramp.InternalEVM2EVMMessage, error) {
	var sentmsg evm_2_evm_onramp.InternalEVM2EVMMessage
	log.Info().Int("msg Number", int(msgSerialNo)).Msg("waiting for CCIPSendRequested")
	ctx, cancel := context.WithTimeout(context.Background(), c.CallTimeOut)
	defer cancel()
	for {
		select {
		case <-ticker.C:
			iterator, err := c.Source.OnRamp.FilterCCIPSendRequested(currentBlockOnSource)
			if err != nil {
				c.updatestats(msgSerialNo, "", CCIPSendRe, time.Since(timeNow), fail)
				return sentmsg, fmt.Errorf("error %v in filtering CCIPSendRequested event for msg ID %d tx %s", err, msgSerialNo, txHash)
			}
			for iterator.Next() {
				if iterator.Event.Raw.TxHash.Hex() == txHash {
					sentmsg = iterator.Event.Message
					c.updatestats(msgSerialNo, fmt.Sprint(sentmsg.SequenceNumber), CCIPSendRe, time.Since(timeNow), success)
					return sentmsg, nil
				}
			}
		case <-ctx.Done():
			c.updatestats(msgSerialNo, "", CCIPSendRe, time.Since(timeNow), fail)
			latest, _ := c.Source.Common.ChainClient.LatestBlockNumber(context.Background())
			return sentmsg, fmt.Errorf("CCIPSendRequested event is not found for msg ID %d tx %s startblock %d latestblock %d", msgSerialNo, txHash, currentBlockOnSource, latest)
		}
	}
}
