package load

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/wasp"
	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"

	"github.com/smartcontractkit/chainlink/integration-tests/testreporters"
	"github.com/smartcontractkit/chainlink/integration-tests/testsetups"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/testhelpers"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

type CCIPE2ELoad struct {
	t                     *testing.T
	Lane                  *actions.CCIPLane
	NoOfReq               int64 // no of Request fired - required for balance assertion at the end
	totalGEFee            *big.Int
	BalanceStats          BalanceStats  // balance assertion details
	CurrentMsgSerialNo    *atomic.Int64 // current msg serial number in the load sequence
	InitialSourceBlockNum uint64
	InitialDestBlockNum   uint64        // blocknumber before the first message is fired in the load sequence
	CallTimeOut           time.Duration // max time to wait for various on-chain events
	reports               *testreporters.CCIPLaneStats
	msg                   router.ClientEVM2AnyMessage
}
type BalanceStats struct {
	SourceBalanceReq        map[string]*big.Int
	SourceBalanceAssertions []testhelpers.BalanceAssertion
	DestBalanceReq          map[string]*big.Int
	DestBalanceAssertions   []testhelpers.BalanceAssertion
}

func NewCCIPLoad(t *testing.T, lane *actions.CCIPLane, timeout time.Duration, noOfReq int64, reporter *testreporters.CCIPLaneStats) *CCIPE2ELoad {
	return &CCIPE2ELoad{
		t:                  t,
		Lane:               lane,
		CurrentMsgSerialNo: atomic.NewInt64(1),
		CallTimeOut:        timeout,
		NoOfReq:            noOfReq,
		reports:            reporter,
	}
}

// BeforeAllCall funds subscription, approves the token transfer amount.
// Needs to be called before load sequence is started.
// Needs to approve and fund for the entire sequence.
func (c *CCIPE2ELoad) BeforeAllCall(msgType string) {
	sourceCCIP := c.Lane.Source
	destCCIP := c.Lane.Dest
	var tokenAndAmounts []router.ClientEVMTokenAmount
	for i := range c.Lane.Source.TransferAmount {
		token := sourceCCIP.Common.BridgeTokens[i]
		tokenAndAmounts = append(tokenAndAmounts, router.ClientEVMTokenAmount{
			Token: common.HexToAddress(token.Address()), Amount: c.Lane.Source.TransferAmount[i],
		})
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
	sourceBalances, err := testhelpers.GetBalances(c.t, sourceCCIP.CollectBalanceRequirements())
	require.NoError(c.t, err, "fetching source balance")
	destBalances, err := testhelpers.GetBalances(c.t, destCCIP.CollectBalanceRequirements())
	require.NoError(c.t, err, "fetching dest balance")
	c.BalanceStats = BalanceStats{
		SourceBalanceReq: sourceBalances,
		DestBalanceReq:   destBalances,
	}
	extraArgsV1, err := testhelpers.GetEVMExtraArgsV1(big.NewInt(100_000), false)
	require.NoError(c.t, err, "Failed encoding the options field")

	receiver, err := utils.ABIEncode(`[{"type":"address"}]`, destCCIP.ReceiverDapp.EthAddress)
	require.NoError(c.t, err, "Failed encoding the receiver address")
	c.msg = router.ClientEVM2AnyMessage{
		Receiver:  receiver,
		ExtraArgs: extraArgsV1,
		FeeToken:  common.HexToAddress(sourceCCIP.Common.FeeToken.Address()),
		Data:      []byte("message with Id 1"),
	}
	if msgType == testsetups.TokenTransfer {
		c.msg.TokenAmounts = tokenAndAmounts
	}

	sourceCCIP.Common.ChainClient.ParallelTransactions(false)
	destCCIP.Common.ChainClient.ParallelTransactions(false)
}

func (c *CCIPE2ELoad) Call(_ *wasp.Generator) wasp.CallResult {
	var res wasp.CallResult
	sourceCCIP := c.Lane.Source
	msgSerialNo := c.CurrentMsgSerialNo.Load()
	c.CurrentMsgSerialNo.Inc()

	lggr := c.Lane.Logger.With().Int("msg Number", int(msgSerialNo)).Logger()

	// form the message for transfer
	msgStr := fmt.Sprintf("message with Id %d", msgSerialNo)
	msg := c.msg
	msg.Data = []byte(msgStr)

	feeToken := sourceCCIP.Common.FeeToken.EthAddress
	// initiate the transfer
	lggr.Debug().Str("triggeredAt", time.Now().GoString()).Msg("triggering transfer")
	var sendTx *types.Transaction
	var err error

	// initiate the transfer
	// if the token address is 0x0 it will use Native as fee token and the fee amount should be mentioned in bind.TransactOpts's value

	destChainSelector, err := actions.EvmChainIdToChainSelector(sourceCCIP.DestinationChainId, c.Lane.Dest.Common.ChainClient.NetworkSimulated())
	if err != nil {
		res.Error = err.Error()
		res.Failed = true
		return res
	}
	fee, err := sourceCCIP.Common.Router.GetFee(destChainSelector, msg)
	if err != nil {
		res.Error = err.Error()
		res.Failed = true
		return res
	}
	startTime := time.Now()
	if feeToken != common.HexToAddress("0x0") {
		sendTx, err = sourceCCIP.Common.Router.CCIPSend(destChainSelector, msg, nil)
	} else {
		sendTx, err = sourceCCIP.Common.Router.CCIPSend(destChainSelector, msg, fee)
	}

	if err != nil {
		c.reports.UpdatePhaseStats(msgSerialNo, 0, testreporters.TX, time.Since(startTime), testreporters.Failure)
		lggr.Err(err).Msg("ccip-send tx error for msg ID")
		res.Error = fmt.Sprintf("ccip-send tx error %+v for msg ID %d", err, msgSerialNo)
		res.Data = c.reports.GetPhaseStatsForRequest(msgSerialNo)
		res.Failed = true
		return res
	}
	lggr = lggr.With().Str("Msg Tx", sendTx.Hash().String()).Logger()
	txConfirmationTime := time.Now().UTC()
	rcpt, err1 := c.Lane.Source.Common.ChainClient.GetTxReceipt(sendTx.Hash())
	if err1 == nil {
		hdr, err1 := c.Lane.Source.Common.ChainClient.HeaderByNumber(context.Background(), rcpt.BlockNumber)
		if err1 == nil {
			txConfirmationTime = hdr.Timestamp
		}
	}
	c.reports.UpdatePhaseStats(msgSerialNo, 0, testreporters.TX, startTime.Sub(txConfirmationTime), testreporters.Success,
		testreporters.SendTransactionStats{
			Fee:                fee.String(),
			GasUsed:            rcpt.GasUsed,
			TxHash:             sendTx.Hash().Hex(),
			NoOfTokensSent:     len(msg.TokenAmounts),
			MessageBytesLength: len(msg.Data),
		})
	// wait for
	// - CCIPSendRequested Event log to be generated,
	msgLog, sourceLogTime, err := c.Lane.Source.AssertEventCCIPSendRequested(
		lggr, msgSerialNo, sendTx.Hash().Hex(), c.CallTimeOut, txConfirmationTime, c.reports)
	if err != nil || msgLog == nil {
		lggr.Err(err).Msgf("CCIPSendRequested event error for msg ID %d", msgSerialNo)
		res.Error = err.Error()
		res.Data = c.reports.GetPhaseStatsForRequest(msgSerialNo)
		res.Failed = true
		return res
	}
	sentMsg := msgLog.Message
	seqNum := sentMsg.SequenceNumber

	if bytes.Compare(sentMsg.Data, []byte(msgStr)) != 0 {
		res.Error = fmt.Sprintf("the message byte didnot match expected %s received %s msg ID %d", msgStr, string(sentMsg.Data), msgSerialNo)
		res.Data = c.reports.GetPhaseStatsForRequest(msgSerialNo)
		res.Failed = true
		return res
	}
	sourceLogFinalizedAt, err := c.Lane.Source.AssertSendRequestedLogFinalized(
		lggr, msgSerialNo, seqNum, msgLog, sourceLogTime, c.reports)
	if err != nil {
		lggr.Err(err).Msgf("waiting for source log to be finalized for msg ID %d", msgSerialNo)
		res.Error = err.Error()
		res.Data = c.reports.GetPhaseStatsForRequest(msgSerialNo)
		res.Failed = true
		return res
	}

	// wait for
	// - CommitStore to increase the seq number,
	err = c.Lane.Dest.AssertSeqNumberExecuted(lggr, msgSerialNo, seqNum, c.CallTimeOut, sourceLogFinalizedAt, c.reports)
	if err != nil {
		lggr.Err(err).Msgf("waiting for seq num increase for msg ID %d", msgSerialNo)
		res.Error = err.Error()
		res.Data = c.reports.GetPhaseStatsForRequest(msgSerialNo)
		res.Failed = true
		return res
	}
	// wait for ReportAccepted event
	commitReport, reportAcceptedAt, err := c.Lane.Dest.AssertEventReportAccepted(lggr, msgSerialNo, seqNum, c.CallTimeOut, sourceLogFinalizedAt, c.reports)
	if err != nil || commitReport == nil {
		lggr.Err(err).Msgf("waiting for ReportAcceptedEvent for msg ID %d", msgSerialNo)
		res.Error = err.Error()
		res.Data = c.reports.GetPhaseStatsForRequest(msgSerialNo)
		res.Failed = true
		return res
	}
	blessedAt, err := c.Lane.Dest.AssertReportBlessed(lggr, msgSerialNo, seqNum, c.CallTimeOut, *commitReport, reportAcceptedAt, c.reports)
	if err != nil {
		lggr.Err(err).Msgf("waiting for ReportBlessedEvent for msg ID %d root=%x", msgSerialNo, commitReport.MerkleRoot)
		res.Error = err.Error()
		res.Data = c.reports.GetPhaseStatsForRequest(msgSerialNo)
		res.Failed = true
		return res
	}
	err = c.Lane.Dest.AssertEventExecutionStateChanged(lggr, msgSerialNo, seqNum, c.CallTimeOut, blessedAt, c.reports)
	if err != nil {
		lggr.Err(err).Msgf("waiting for ExecutionStateChangedEvent for msg ID %d", msgSerialNo)
		res.Error = err.Error()
		res.Data = c.reports.GetPhaseStatsForRequest(msgSerialNo)
		res.Failed = true
		return res
	}

	res.Data = c.reports.GetPhaseStatsForRequest(msgSerialNo)
	return res
}

func (c *CCIPE2ELoad) ReportAcceptedLog() {
	c.Lane.Logger.Info().Msg("Commit Report stats")
	it, err := c.Lane.Dest.CommitStore.Instance.FilterReportAccepted(&bind.FilterOpts{Start: c.InitialDestBlockNum})
	require.NoError(c.t, err, "report committed result")
	i := 1
	event := c.Lane.Logger.Info()
	for it.Next() {
		event.Interface(fmt.Sprintf("%d Report Intervals", i), it.Event.Report.Interval)
		i++
	}
	event.Msgf("CommitStore-Reports Accepted")
}
