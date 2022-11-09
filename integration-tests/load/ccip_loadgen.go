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
	"time"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/chainlink-testing-framework/client"
	"github.com/smartcontractkit/chainlink-testing-framework/testreporters"
	"go.uber.org/atomic"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_subscription_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_subscription_onramp"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/testhelpers"
	"github.com/smartcontractkit/chainlink/core/utils"
	bigmath "github.com/smartcontractkit/chainlink/core/utils/big_math"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
)

type phase string
type status string

const (
	E2E                     phase  = "i)OverallRelayAndExecution"
	TX                      phase  = "ii)SendTxBlockConfirmation"
	CCIPSendRe              phase  = "iii)CCIPSendRequested Event"
	SeqNumAndRepAccIncrease phase  = "iv)ReportAcceptedByBlobVerifier(Relay)"
	ExecStateChanged        phase  = "v)ExecutionStateChanged Event(Execution)"
	success                 status = "✅"
	fail                    status = "❌"
	TokenTransfer           string = "WithToken"
	DataOnlyTransfer        string = "WithoutToken"
)

type CCIPE2ELoad struct {
	Source                *actions.SourceCCIPModule // all source contracts
	Destination           *actions.DestCCIPModule   // all destination contracts
	Model                 actions.BillingModel      // toll Or sub
	NoOfReq               int64                     // no of Request fired - required for balance assertion at the end
	BalanceStats          BalanceStats              // balance assertion details
	CurrentMsgID          *atomic.Int64             // current msg serial number in the load sequence
	InitialSourceBlockNum uint64
	InitialDestBlockNum   uint64 // blocknumber before the first message is fired in the load sequence
	sentMsgMu             *sync.Mutex
	SentMsg               map[uint64]evm_2_evm_subscription_onramp.CCIPEVM2EVMSubscriptionMessage // track the messages by seq num for debugging purpose
	CallTimeOut           time.Duration                                                           // max time to wait for various on-chain events
	TickerDuration        time.Duration                                                           // poll frequency while waiting for on-chain events
	callStatsMu           *sync.Mutex
	callStats             map[int64]map[phase]StatParams // keeps track of various phase related metrics
	seqNumRelayedMu       *sync.Mutex
	seqNumRelayed         map[uint64]uint64 // key : seqNumber in the ReportAccepted event, value : blocknumber for corresponding event
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
	MsgId int64                `json:"msgSerialNumber"`
	Stats map[phase]StatParams `json:"phaseDetails"`
}

type SlackStats struct {
	AvgE2EDuration                  float64 `json:"avgOverallRelayAndExecution,omitempty"`
	AvgRelayDuration                float64 `json:"avgRelay,omitempty"`
	AvgExecDuration                 float64 `json:"avgExecution,omitempty"`
	LongestE2EDuration              float64 `json:"longestOverallRelayAndExecution,omitempty"`
	LongestRelayDuration            float64 `json:"longestRelay,omitempty"`
	LongestExecDuration             float64 `json:"longestExecution,omitempty"`
	FastestE2EDuration              float64 `json:"fastestOverallRelayAndExecution,omitempty"`
	FastestRelayDuration            float64 `json:"fastestRelay,omitempty"`
	FastestExecDuration             float64 `json:"fastestExecution,omitempty"`
	TotalNumberOfFailedRequests     int     `json:"totalFailedRequests,omitempty"`
	TotalNumberOfSuccessfulRequests int     `json:"totalSuccessfulRequests,omitempty"`
	FailedRelay                     int     `json:"noOfFailedRelay,omitempty"`
	FailedExecution                 int     `json:"noOfFailedExecution,omitempty"`
	FailedSendTransaction           int     `json:"noOfFailedSendTransaction,omitempty"`
	FailedCCIPSendRequested         int     `json:"noOfFailedCCIPSendRequested,omitempty"`
}

func NewCCIPLoad(source *actions.SourceCCIPModule, dest *actions.DestCCIPModule, model actions.BillingModel, timeout time.Duration, noOfReq int64) *CCIPE2ELoad {
	return &CCIPE2ELoad{
		Source:          source,
		Destination:     dest,
		Model:           model,
		CurrentMsgID:    atomic.NewInt64(1),
		sentMsgMu:       &sync.Mutex{},
		SentMsg:         make(map[uint64]evm_2_evm_subscription_onramp.CCIPEVM2EVMSubscriptionMessage),
		TickerDuration:  time.Second,
		CallTimeOut:     timeout,
		NoOfReq:         noOfReq,
		callStats:       make(map[int64]map[phase]StatParams),
		callStatsMu:     &sync.Mutex{},
		seqNumRelayedMu: &sync.Mutex{},
		seqNumRelayed:   make(map[uint64]uint64),
	}
}

// BeforeAllCall funds subscription, approves the token transfer amount.
// Needs to be called before load sequence is started.
// Needs to approve and fund for the entire sequence.
func (c *CCIPE2ELoad) BeforeAllCall() {
	var sourceTokens []common.Address
	sourceCCIP := c.Source
	destCCIP := c.Destination
	for i, token := range sourceCCIP.Common.BridgeTokens {
		sourceTokens = append(sourceTokens, common.HexToAddress(token.Address()))
		// approve the onramp router so that it can initiate transferring the token
		err := token.Approve(sourceCCIP.SubOnRampRouter.Address(), bigmath.Mul(sourceCCIP.TransferAmount[i], c.NoOfReq))
		Expect(err).ShouldNot(HaveOccurred(), "Could not approve permissions for the onRamp router "+
			"on the source link token contract")
	}
	err := sourceCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Failed to wait for events")

	if c.Model == actions.SUB {
		actions.CreateAndFundSubscription(*sourceCCIP, *destCCIP, bigmath.Mul(big.NewInt(c.NoOfReq*100), big.NewInt(1e18)), c.NoOfReq)
	}
	// save the current block numbers to use in various filter log requests
	currentBlockOnSource, err := sourceCCIP.Common.ChainClient.LatestBlockNumber(context.Background())
	Expect(err).ShouldNot(HaveOccurred(), "failed to fetch latest source block num")
	currentBlockOnDest, err := destCCIP.Common.ChainClient.LatestBlockNumber(context.Background())
	Expect(err).ShouldNot(HaveOccurred(), "failed to fetch latest dest block num")
	c.InitialDestBlockNum = currentBlockOnDest
	c.InitialSourceBlockNum = currentBlockOnSource
	// collect the balance requirement to verify balances after transfer
	sourceBalances, err := testhelpers.GetBalances(sourceCCIP.CollectBalanceRequirements(c.Model))
	Expect(err).ShouldNot(HaveOccurred(), "fetching source balance")
	destBalances, err := testhelpers.GetBalances(destCCIP.CollectBalanceRequirements(c.Model))
	Expect(err).ShouldNot(HaveOccurred(), "fetching dest balance")
	c.BalanceStats = BalanceStats{
		SourceBalanceReq: sourceBalances,
		DestBalanceReq:   destBalances,
	}
}

func (c *CCIPE2ELoad) AfterAllCall() {
	c.BalanceStats.DestBalanceAssertions = c.Destination.BalanceAssertions(
		c.Model, c.BalanceStats.DestBalanceReq, c.Source.TransferAmount, big.NewInt(0),
		c.NoOfReq, bigmath.Mul(big.NewInt(c.NoOfReq), big.NewInt(0.79e18)))
	c.BalanceStats.SourceBalanceAssertions = c.Source.BalanceAssertions(c.Model, c.BalanceStats.SourceBalanceReq, c.NoOfReq)
	actions.AssertBalances(c.BalanceStats.DestBalanceAssertions)
	actions.AssertBalances(c.BalanceStats.SourceBalanceAssertions)
}

func (c *CCIPE2ELoad) Call(msgType interface{}) client.CallResult {
	var res client.CallResult
	sourceCCIP := c.Source
	destCCIP := c.Destination
	msgID := c.CurrentMsgID.Load()
	c.CurrentMsgID.Inc()
	var sourceTokensAndAmounts []evm_2_any_subscription_onramp_router.CCIPEVMTokenAndAmount
	for i, token := range sourceCCIP.Common.BridgeTokens {
		sourceTokensAndAmounts = append(sourceTokensAndAmounts, evm_2_any_subscription_onramp_router.CCIPEVMTokenAndAmount{
			Token:  common.HexToAddress(token.Address()),
			Amount: sourceCCIP.TransferAmount[i],
		})
	}

	extraArgsV1, err := testhelpers.GetEVMExtraArgsV1(big.NewInt(100_000))
	Expect(err).ShouldNot(HaveOccurred(), "Failed encoding the options field")

	// form the message for transfer
	msgStr := fmt.Sprintf("message with Id %d", msgID)
	receiver, err := utils.ABIEncode(`[{"type":"address"}]`, destCCIP.ReceiverDapp.EthAddress)
	Expect(err).ShouldNot(HaveOccurred(), "Failed encoding the receiver address")
	msg := evm_2_any_subscription_onramp_router.CCIPEVM2AnySubscriptionMessage{
		Receiver:  receiver,
		Data:      []byte(msgStr),
		ExtraArgs: extraArgsV1,
	}
	if msgType == TokenTransfer {
		msg.TokensAndAmounts = sourceTokensAndAmounts
	}
	startTime := time.Now()
	// initiate the transfer
	log.Debug().Int("msg Number", int(msgID)).Str("triggeredAt", time.Now().GoString()).Msg("triggering transfer")
	sendTx, err := sourceCCIP.SubOnRampRouter.CCIPSend(destCCIP.Common.ChainClient.GetChainID(), msg)
	if err != nil {
		c.updatestats(msgID, "", TX, time.Since(startTime), fail)
		res.Error = fmt.Errorf("CCIPSend request error %v for msg ID %d", err, msgID)
		return res
	}
	c.updatestats(msgID, "", TX, time.Since(startTime), success)

	// wait for
	// - CCIPSendRequested Event log to be generated,
	ticker := time.NewTicker(c.TickerDuration)
	defer ticker.Stop()
	sentMsg, err := c.waitForCCIPSendRequested(ticker, c.InitialSourceBlockNum, msgID, sendTx.Hash().Hex(), time.Now())
	if err != nil {
		res.Error = err
		return res
	}
	relayStartTime := time.Now()
	seqNum := sentMsg.SequenceNumber
	if bytes.Compare(sentMsg.Data, []byte(msgStr)) == 0 {
		c.updateSentMsgQueue(seqNum, sentMsg)
	} else {
		res.Error = fmt.Errorf("the message byte didnot match expected %s received %s msg ID %d", msgStr, string(sentMsg.Data), msgID)
		return res
	}

	// wait for
	// - BlobVerifier to increase the seq number,
	err = c.waitForSeqNumberIncrease(ticker, seqNum, msgID, relayStartTime)
	if err != nil {
		res.Error = err
		return res
	}
	// wait for ReportAccepted event
	err = c.waitForReportAccepted(ticker, msgID, seqNum, c.InitialDestBlockNum, relayStartTime)
	if err != nil {
		res.Error = err
		return res
	}

	// wait for ExecutionStateChanged event
	err = c.waitForExecStateChange(ticker, []uint64{seqNum}, c.seqNumRelayed[seqNum]-2, msgID, time.Now())
	if err != nil {
		res.Error = err
		return res
	}
	c.updatestats(msgID, fmt.Sprint(seqNum), E2E, time.Since(relayStartTime), success)
	res.Error = nil
	res.Data = c.SentMsg[seqNum]
	return res
}

func (c *CCIPE2ELoad) updateSentMsgQueue(seqNum uint64, sentMsg evm_2_evm_subscription_onramp.CCIPEVM2EVMSubscriptionMessage) {
	c.sentMsgMu.Lock()
	defer c.sentMsgMu.Unlock()
	c.SentMsg[seqNum] = sentMsg
}

func (c *CCIPE2ELoad) updateSeqNumRelayed(seqNum []uint64, blockNum uint64) {
	c.seqNumRelayedMu.Lock()
	defer c.seqNumRelayedMu.Unlock()
	for _, num := range seqNum {
		if _, ok := c.seqNumRelayed[num]; ok {
			return
		}
		c.seqNumRelayed[num] = blockNum
	}
}

func (c *CCIPE2ELoad) updatestats(msgId int64, seqNum string, step phase, duration time.Duration, state status) {
	c.callStatsMu.Lock()
	defer c.callStatsMu.Unlock()
	if _, ok := c.callStats[msgId]; !ok {
		c.callStats[msgId] = make(map[phase]StatParams)
		c.callStats[msgId][step] = StatParams{
			SeqNum:    seqNum,
			Duration:  duration.Seconds(),
			ReqStatus: state,
		}
	} else {
		c.callStats[msgId][step] = StatParams{
			SeqNum:    seqNum,
			Duration:  c.callStats[msgId][step].Duration + duration.Seconds(),
			ReqStatus: state,
		}
	}
	// if any of the phase fails mark the E2E as failed
	if state == fail {
		c.callStats[msgId][E2E] = StatParams{
			SeqNum:    seqNum,
			ReqStatus: state,
		}
	} else {
		log.Info().Str(fmt.Sprint(step), fmt.Sprint(state)).Msgf("seq num %s", seqNum)
	}
}

func (c *CCIPE2ELoad) PrintStats(failed bool, rps int, duration float64) {
	if _, err := os.Stat("./logs/stats"); os.IsNotExist(err) {
		os.MkdirAll("./logs/stats", 0700)
	}
	tempFile, err := os.Create("./logs/stats/CCIPLoad_complete.json")
	defer tempFile.Close()
	Expect(err).ShouldNot(HaveOccurred(), "creating stat file")

	tempStatFile, err := os.Create("./logs/stats/CCIPLoad_avg.json")
	defer tempStatFile.Close()
	Expect(err).ShouldNot(HaveOccurred(), "creating stat file")

	slackFile, err := os.Create("./logs/payload-slack-content.json")
	defer slackFile.Close()
	Expect(err).ShouldNot(HaveOccurred(), "creating stat file")

	log.Info().Msg("Msg Stats")
	keys := make([]int64, 0, len(c.callStats))
	for msgId := range c.callStats {
		keys = append(keys, msgId)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	durationSumMap := make(map[phase]float64)
	durationMap := make(map[phase][]float64)
	successCount := make(map[phase]int)
	failureCount := make(map[phase]int)
	var jsonstats []JsonStats

	for _, msgId := range keys {
		stats := c.callStats[msgId]
		event := log.Info()
		jsonstats = append(jsonstats, JsonStats{
			MsgId: msgId,
			Stats: stats,
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
		event.Msgf("Msg stats for seq num %d", msgId)
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
		AvgRelayDuration:                overallStats[SeqNumAndRepAccIncrease].Duration,
		AvgExecDuration:                 overallStats[ExecStateChanged].Duration,
		LongestE2EDuration:              durationMap[E2E][len(durationMap[E2E])-1],
		LongestRelayDuration:            durationMap[SeqNumAndRepAccIncrease][len(durationMap[SeqNumAndRepAccIncrease])-1],
		LongestExecDuration:             durationMap[ExecStateChanged][len(durationMap[ExecStateChanged])-1],
		FastestE2EDuration:              durationMap[E2E][0],
		FastestRelayDuration:            durationMap[SeqNumAndRepAccIncrease][0],
		FastestExecDuration:             durationMap[ExecStateChanged][0],
		TotalNumberOfFailedRequests:     failureCount[E2E],
		TotalNumberOfSuccessfulRequests: successCount[E2E],
		FailedRelay:                     failureCount[SeqNumAndRepAccIncrease],
		FailedExecution:                 failureCount[ExecStateChanged],
		FailedSendTransaction:           failureCount[TX],
		FailedCCIPSendRequested:         failureCount[CCIPSendRe],
	}

	stats, err := json.MarshalIndent(overallStats, "", "  ")
	Expect(err).ShouldNot(HaveOccurred(), "marshal overallStats")
	_, err = tempStatFile.Write(stats)
	Expect(err).ShouldNot(HaveOccurred(), "writing overallStats")

	stats, err = json.MarshalIndent(jsonstats, "", "  ")
	Expect(err).ShouldNot(HaveOccurred(), "marshal avg stats")
	_, err = tempFile.Write(stats)
	Expect(err).ShouldNot(HaveOccurred(), "writing avg stats")

	headerText := ":white_check_mark: CCIP Load Test PASSED :white_check_mark:"
	if failed {
		headerText = ":x: CCIP Load Test FAILED :x:"
	}
	stats, err = json.MarshalIndent(slackStats, "", "  ")
	Expect(err).ShouldNot(HaveOccurred())

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
	log.Info().Msg("Relay Report stats")
	it, err := c.Destination.BlobVerifier.FilterReportAccepted(c.InitialDestBlockNum)
	Expect(err).ShouldNot(HaveOccurred(), "report relayed result")
	i := 1
	event = log.Info()
	for it.Next() {
		event.Interface(fmt.Sprintf("%d Report Intervals", i), it.Event.Report.Intervals)
		i++
	}
	event.Msgf("BlobVerifier-Reports Accepted")
}

func (c *CCIPE2ELoad) waitForExecStateChange(ticker *time.Ticker, seqNums []uint64, currentBlockOnDest uint64, msgID int64, timeNow time.Time) error {
	log.Info().Int("msg Number", int(msgID)).Msgf(
		"waiting for ExecutionStateChanged for seqNums %v", seqNums)
	ctx, cancel := context.WithTimeout(context.Background(), c.CallTimeOut)
	defer cancel()
	for {
		select {
		case <-ticker.C:
			iterator, err := c.Destination.SubOffRamp.FilterExecutionStateChanged(seqNums, currentBlockOnDest)
			if err != nil {
				for _, seqNum := range seqNums {
					c.updatestats(msgID, fmt.Sprint(seqNum), ExecStateChanged, time.Since(timeNow), fail)
				}
				return fmt.Errorf("filtering event ExecutionStateChanged returned error %v msg ID %d and seqNum %v", err, msgID, seqNums)
			}
			for iterator.Next() {
				switch ccip.MessageExecutionState(iterator.Event.State) {
				case ccip.Success:
					for _, seqNum := range seqNums {
						c.updatestats(msgID, fmt.Sprint(seqNum), ExecStateChanged, time.Since(timeNow), success)
					}
					return nil
				case ccip.Failure:
					for _, seqNum := range seqNums {
						c.updatestats(msgID, fmt.Sprint(seqNum), ExecStateChanged, time.Since(timeNow), fail)
					}
					return fmt.Errorf("ExecutionStateChanged event returned failure for seq num %v msg ID %d", seqNums, msgID)
				}
			}
		case <-ctx.Done():
			for _, seqNum := range seqNums {
				c.updatestats(msgID, fmt.Sprint(seqNum), ExecStateChanged, time.Since(timeNow), fail)
			}
			return fmt.Errorf("ExecutionStateChanged event not found for seq num %v msg ID %d", seqNums, msgID)
		}
	}
}

func (c *CCIPE2ELoad) waitForSeqNumberIncrease(ticker *time.Ticker, seqNum uint64, msgID int64, timeNow time.Time) error {
	log.Info().Int("msg Number", int(msgID)).Msgf("waiting for seq number %d to get increased", int(seqNum))
	ctx, cancel := context.WithTimeout(context.Background(), c.CallTimeOut)
	defer cancel()
	for {
		select {
		case <-ticker.C:
			seqNumberAfter, err := c.Destination.BlobVerifier.GetNextSeqNumber(c.Source.SubOnRamp.EthAddress)
			if err != nil {
				c.updatestats(msgID, fmt.Sprint(seqNum), SeqNumAndRepAccIncrease, time.Since(timeNow), fail)
				return fmt.Errorf("error %v in GetNextExpectedSeqNumber by blobverifier for msg ID %d", err, msgID)
			}
			if seqNumberAfter > seqNum {
				return nil
			}
		case <-ctx.Done():
			c.updatestats(msgID, fmt.Sprint(seqNum), SeqNumAndRepAccIncrease, time.Since(timeNow), fail)
			return fmt.Errorf("sequence number is not increased for seq num %d msg ID %d", seqNum, msgID)
		}
	}
}

func (c *CCIPE2ELoad) waitForReportAccepted(ticker *time.Ticker, msgID int64, seqNum uint64, currentBlockOnDest uint64, timeNow time.Time) error {
	log.Info().Int("seq Number", int(seqNum)).Msg("waiting for ReportAccepted")
	ctx, cancel := context.WithTimeout(context.Background(), c.CallTimeOut)
	defer cancel()
	for {
		select {
		case <-ticker.C:
			// skip calling FilterReportAccepted if the seqNum is present in the map
			if _, ok := c.seqNumRelayed[seqNum]; ok {
				c.updatestats(msgID, fmt.Sprint(seqNum), SeqNumAndRepAccIncrease, time.Since(timeNow), success)
				return nil
			}
			it, err := c.Destination.BlobVerifier.FilterReportAccepted(currentBlockOnDest)
			if err != nil {
				c.updatestats(msgID, fmt.Sprint(seqNum), SeqNumAndRepAccIncrease, time.Since(timeNow), fail)
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
					// update SeqNumRelayed map for all seqNums in the emitted ReportAccepted event
					c.updateSeqNumRelayed(seqNums, it.Event.Raw.BlockNumber)
					if in.Max >= seqNum && in.Min <= seqNum {
						c.updatestats(msgID, fmt.Sprint(seqNum), SeqNumAndRepAccIncrease, time.Since(timeNow), success)
						return nil
					}
				}
			}
		case <-ctx.Done():
			c.updatestats(msgID, fmt.Sprint(seqNum), SeqNumAndRepAccIncrease, time.Since(timeNow), fail)
			return fmt.Errorf("ReportAccepted is not found for seq num %d", seqNum)
		}
	}
}

func (c *CCIPE2ELoad) waitForCCIPSendRequested(
	ticker *time.Ticker,
	currentBlockOnSource uint64,
	msgID int64,
	txHash string,
	timeNow time.Time,
) (evm_2_evm_subscription_onramp.CCIPEVM2EVMSubscriptionMessage, error) {
	var sentmsg evm_2_evm_subscription_onramp.CCIPEVM2EVMSubscriptionMessage
	log.Info().Int("msg Number", int(msgID)).Msg("waiting for CCIPSendRequested")
	ctx, cancel := context.WithTimeout(context.Background(), c.CallTimeOut)
	defer cancel()
	for {
		select {
		case <-ticker.C:
			iterator, err := c.Source.SubOnRamp.FilterCCIPSendRequested(currentBlockOnSource)
			if err != nil {
				c.updatestats(msgID, "", CCIPSendRe, time.Since(timeNow), fail)
				return sentmsg, fmt.Errorf("error %v in filtering CCIPSendRequested event for msg ID %d tx %s", err, msgID, txHash)
			}
			for iterator.Next() {
				if iterator.Event.Raw.TxHash.Hex() == txHash {
					sentmsg = iterator.Event.Message
					c.updatestats(msgID, fmt.Sprint(sentmsg.SequenceNumber), CCIPSendRe, time.Since(timeNow), success)
					return sentmsg, nil
				}
			}
		case <-ctx.Done():
			c.updatestats(msgID, "", CCIPSendRe, time.Since(timeNow), fail)
			latest, _ := c.Source.Common.ChainClient.LatestBlockNumber(context.Background())
			return sentmsg, fmt.Errorf("CCIPSendRequested event is not found for msg ID %d tx %s startblock %d latestblock %d", msgID, txHash, currentBlockOnSource, latest)
		}
	}
}
