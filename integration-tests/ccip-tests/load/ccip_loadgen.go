package load

import (
	"context"
	crypto_rand "crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/wasp"
	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"

	"github.com/smartcontractkit/chainlink-common/pkg/config"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/testhelpers"

	"github.com/smartcontractkit/chainlink/integration-tests/ccip-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/ccip-tests/testreporters"
)

// CCIPLaneOptimized is a light-weight version of CCIPLane, It only contains elements which are used during load triggering and validation
type CCIPLaneOptimized struct {
	Logger            zerolog.Logger
	SourceNetworkName string
	DestNetworkName   string
	Source            *actions.SourceCCIPModule
	Dest              *actions.DestCCIPModule
	Reports           *testreporters.CCIPLaneStats
}

type CCIPE2ELoad struct {
	t                                          *testing.T
	Lane                                       *CCIPLaneOptimized
	NoOfReq                                    int64         // approx no of Request fired
	CurrentMsgSerialNo                         *atomic.Int64 // current msg serial number in the load sequence
	CallTimeOut                                time.Duration // max time to wait for various on-chain events
	msg                                        router.ClientEVM2AnyMessage
	MaxDataBytes                               uint32
	SendMaxDataIntermittentlyInMsgCount        int64
	SkipRequestIfAnotherRequestTriggeredWithin *config.Duration
	LastFinalizedTxBlock                       atomic.Uint64
	LastFinalizedTimestamp                     atomic.Time
}

func NewCCIPLoad(
	t *testing.T,
	lane *actions.CCIPLane,
	timeout time.Duration,
	noOfReq int64,
	sendMaxDataIntermittentlyInEveryMsgCount int64,
	SkipRequestIfAnotherRequestTriggeredWithin *config.Duration,
) *CCIPE2ELoad {
	// to avoid holding extra data
	loadLane := &CCIPLaneOptimized{
		Logger:            lane.Logger,
		SourceNetworkName: lane.SourceNetworkName,
		DestNetworkName:   lane.DestNetworkName,
		Source:            lane.Source,
		Dest:              lane.Dest,
		Reports:           lane.Reports,
	}

	return &CCIPE2ELoad{
		t:                                   t,
		Lane:                                loadLane,
		CurrentMsgSerialNo:                  atomic.NewInt64(1),
		CallTimeOut:                         timeout,
		NoOfReq:                             noOfReq,
		SendMaxDataIntermittentlyInMsgCount: sendMaxDataIntermittentlyInEveryMsgCount,
		SkipRequestIfAnotherRequestTriggeredWithin: SkipRequestIfAnotherRequestTriggeredWithin,
	}
}

// BeforeAllCall funds subscription, approves the token transfer amount.
// Needs to be called before load sequence is started.
// Needs to approve and fund for the entire sequence.
func (c *CCIPE2ELoad) BeforeAllCall(isTokenTranfer bool, gasLimit *big.Int) {
	sourceCCIP := c.Lane.Source
	destCCIP := c.Lane.Dest
	extraArgsV1, err := testhelpers.GetEVMExtraArgsV1(gasLimit, false)
	require.NoError(c.t, err, "Failed encoding the options field")

	receiver, err := utils.ABIEncode(`[{"type":"address"}]`, destCCIP.ReceiverDapp.EthAddress)
	require.NoError(c.t, err, "Failed encoding the receiver address")
	c.msg = router.ClientEVM2AnyMessage{
		Receiver:  receiver,
		ExtraArgs: extraArgsV1,
		FeeToken:  common.HexToAddress(sourceCCIP.Common.FeeToken.Address()),
		Data:      []byte("message with Id 1"),
	}
	var tokenAndAmounts []router.ClientEVMTokenAmount
	if isTokenTranfer && len(c.Lane.Source.Common.BridgeTokens) > 0 {
		for i := range c.Lane.Source.TransferAmount {
			// if length of sourceCCIP.TransferAmount is more than available bridge token use first bridge token
			token := sourceCCIP.Common.BridgeTokens[0]
			if i < len(sourceCCIP.Common.BridgeTokens) {
				token = sourceCCIP.Common.BridgeTokens[i]
			}
			tokenAndAmounts = append(tokenAndAmounts, router.ClientEVMTokenAmount{
				Token: common.HexToAddress(token.Address()), Amount: c.Lane.Source.TransferAmount[i],
			})
		}
		c.msg.TokenAmounts = tokenAndAmounts
	}

	if c.SendMaxDataIntermittentlyInMsgCount > 0 {
		dCfg, err := sourceCCIP.OnRamp.Instance.GetDynamicConfig(nil)
		require.NoError(c.t, err, "failed to fetch dynamic config")
		c.MaxDataBytes = dCfg.MaxDataBytes
	}
	// if the msg is sent via multicall, transfer the token transfer amount to multicall contract
	if sourceCCIP.Common.MulticallEnabled &&
		sourceCCIP.Common.MulticallContract != (common.Address{}) &&
		isTokenTranfer {
		for i, amount := range sourceCCIP.TransferAmount {
			// if length of sourceCCIP.TransferAmount is more than available bridge token use first bridge token
			token := sourceCCIP.Common.BridgeTokens[0]
			if i < len(sourceCCIP.Common.BridgeTokens) {
				token = sourceCCIP.Common.BridgeTokens[i]
			}
			amountToApprove := new(big.Int).Mul(amount, big.NewInt(c.NoOfReq))
			bal, err := token.BalanceOf(context.Background(), sourceCCIP.Common.MulticallContract.Hex())
			require.NoError(c.t, err, "Failed to get token balance")
			if bal.Cmp(amountToApprove) < 0 {
				err := token.Transfer(sourceCCIP.Common.MulticallContract.Hex(), amountToApprove)
				require.NoError(c.t, err, "Failed to approve token transfer amount")
			}
		}
	}

	c.LastFinalizedTxBlock.Store(c.Lane.Source.NewFinalizedBlockNum.Load())
	c.LastFinalizedTimestamp.Store(c.Lane.Source.NewFinalizedBlockTimestamp.Load())

	sourceCCIP.Common.ChainClient.ParallelTransactions(false)
	destCCIP.Common.ChainClient.ParallelTransactions(false)
}

func (c *CCIPE2ELoad) CCIPMsg() (router.ClientEVM2AnyMessage, *testreporters.RequestStat, error) {
	msgSerialNo := c.CurrentMsgSerialNo.Load()
	c.CurrentMsgSerialNo.Inc()

	stats := testreporters.NewCCIPRequestStats(msgSerialNo, c.Lane.SourceNetworkName, c.Lane.DestNetworkName)
	// form the message for transfer
	msgLength := c.Lane.Source.MsgDataLength
	msg := c.msg
	if msgLength > 0 {
		if c.SendMaxDataIntermittentlyInMsgCount > 0 {
			// every SendMaxDataIntermittentlyInMsgCount message will have extra data with almost MaxDataBytes
			if msgSerialNo%c.SendMaxDataIntermittentlyInMsgCount == 0 {
				msgLength = int64(c.MaxDataBytes - 1)
			}
		}
		b := make([]byte, msgLength)
		_, err := crypto_rand.Read(b)
		if err != nil {
			return router.ClientEVM2AnyMessage{}, stats, fmt.Errorf("failed to generate random string %w", err)
		}
		randomString := base64.URLEncoding.EncodeToString(b)
		msg.Data = []byte(randomString[:msgLength])
	}

	return msg, stats, nil
}

func (c *CCIPE2ELoad) Call(_ *wasp.Generator) *wasp.Response {
	res := &wasp.Response{}
	sourceCCIP := c.Lane.Source
	recentRequestFoundAt := sourceCCIP.IsRequestTriggeredWithinTimeframe(c.SkipRequestIfAnotherRequestTriggeredWithin)
	if recentRequestFoundAt != nil {
		c.Lane.Logger.
			Info().
			Str("Found At=", recentRequestFoundAt.String()).
			Msgf("Skipping ...Another Request found within given timeframe %s", c.SkipRequestIfAnotherRequestTriggeredWithin.String())
		return res
	}
	msg, stats, err := c.CCIPMsg()
	if err != nil {
		res.Error = err.Error()
		res.Failed = true
		return res
	}
	msgSerialNo := stats.ReqNo
	lggr := c.Lane.Logger.With().Int64("msg Number", stats.ReqNo).Logger()

	feeToken := sourceCCIP.Common.FeeToken.EthAddress
	// initiate the transfer
	lggr.Debug().Str("triggeredAt", time.Now().GoString()).Msg("triggering transfer")
	var sendTx *types.Transaction

	destChainSelector, err := chain_selectors.SelectorFromChainId(sourceCCIP.DestinationChainId)
	if err != nil {
		res.Error = err.Error()
		res.Failed = true
		return res
	}
	// initiate the transfer
	// if the token address is 0x0 it will use Native as fee token and the fee amount should be mentioned in bind.TransactOpts's value

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
		// add a bit buffer to fee
		sendTx, err = sourceCCIP.Common.Router.CCIPSend(destChainSelector, msg, new(big.Int).Add(big.NewInt(1e5), fee))
	}
	if err != nil {
		stats.UpdateState(lggr, 0, testreporters.TX, time.Since(startTime), testreporters.Failure)
		res.Error = err.Error()
		res.Data = stats.StatusByPhase
		res.Failed = true
		return res
	}

	err = sourceCCIP.Common.ChainClient.MarkTxAsSentOnL2(sendTx)

	if err != nil {
		stats.UpdateState(lggr, 0, testreporters.TX, time.Since(startTime), testreporters.Failure)
		res.Error = fmt.Sprintf("ccip-send tx error %+v for msg ID %d", err, msgSerialNo)
		res.Data = stats.StatusByPhase
		res.Failed = true
		return res
	}

	txConfirmationTime := time.Now().UTC()
	rcpt, err1 := bind.WaitMined(context.Background(), sourceCCIP.Common.ChainClient.DeployBackend(), sendTx)
	if err1 == nil {
		hdr, err1 := c.Lane.Source.Common.ChainClient.HeaderByNumber(context.Background(), rcpt.BlockNumber)
		if err1 == nil {
			txConfirmationTime = hdr.Timestamp
		}
	}
	lggr = lggr.With().Str("Msg Tx", sendTx.Hash().String()).Logger()
	var gasUsed uint64
	if rcpt != nil {
		gasUsed = rcpt.GasUsed
	}
	if rcpt.Status != types.ReceiptStatusSuccessful {
		stats.UpdateState(lggr, 0, testreporters.TX, startTime.Sub(txConfirmationTime), testreporters.Failure,
			testreporters.TransactionStats{
				Fee:                fee.String(),
				GasUsed:            gasUsed,
				TxHash:             sendTx.Hash().Hex(),
				NoOfTokensSent:     len(msg.TokenAmounts),
				MessageBytesLength: int64(len(msg.Data)),
			})
		errReason, v, err := c.Lane.Source.Common.ChainClient.RevertReasonFromTx(rcpt.TxHash, router.RouterABI)
		if err != nil {
			errReason = "could not decode"
		}
		res.Error = fmt.Sprintf("ccip-send request receipt is not successful, errReason=%s, args =%v", errReason, v)
		res.Failed = true
		res.Data = stats.StatusByPhase
		return res
	}
	stats.UpdateState(lggr, 0, testreporters.TX, startTime.Sub(txConfirmationTime), testreporters.Success,
		testreporters.TransactionStats{
			Fee:                fee.String(),
			GasUsed:            gasUsed,
			TxHash:             sendTx.Hash().Hex(),
			NoOfTokensSent:     len(msg.TokenAmounts),
			MessageBytesLength: int64(len(msg.Data)),
		})
	err = c.Validate(lggr, sendTx, txConfirmationTime, []*testreporters.RequestStat{stats})
	if err != nil {
		res.Error = err.Error()
		res.Failed = true
		res.Data = stats.StatusByPhase
		return res
	}
	res.Data = stats.StatusByPhase
	return res
}

func (c *CCIPE2ELoad) Validate(lggr zerolog.Logger, sendTx *types.Transaction, txConfirmationTime time.Time, stats []*testreporters.RequestStat) error {
	// wait for
	// - CCIPSendRequested Event log to be generated,
	msgLogs, sourceLogTime, err := c.Lane.Source.AssertEventCCIPSendRequested(lggr, sendTx.Hash().Hex(), c.CallTimeOut, txConfirmationTime, stats)

	if err != nil || msgLogs == nil || len(msgLogs) == 0 {
		return err
	}

	lstFinalizedBlock := c.LastFinalizedTxBlock.Load()
	var sourceLogFinalizedAt time.Time
	// if the finality tag is enabled and the last finalized block is greater than the block number of the message
	// consider the message finalized
	if c.Lane.Source.Common.ChainClient.GetNetworkConfig().FinalityDepth == 0 &&
		lstFinalizedBlock != 0 && lstFinalizedBlock > msgLogs[0].Raw.BlockNumber {
		sourceLogFinalizedAt = c.LastFinalizedTimestamp.Load()
		for _, stat := range stats {
			stat.UpdateState(lggr, stat.SeqNum, testreporters.SourceLogFinalized,
				sourceLogFinalizedAt.Sub(sourceLogTime), testreporters.Success,
				testreporters.TransactionStats{
					TxHash:           msgLogs[0].Raw.TxHash.String(),
					FinalizedByBlock: strconv.FormatUint(lstFinalizedBlock, 10),
					FinalizedAt:      sourceLogFinalizedAt.String(),
				})
		}
	} else {
		var finalizingBlock uint64
		sourceLogFinalizedAt, finalizingBlock, err = c.Lane.Source.AssertSendRequestedLogFinalized(
			lggr, sendTx.Hash(), sourceLogTime, stats)
		if err != nil {
			return err
		}
		c.LastFinalizedTxBlock.Store(finalizingBlock)
		c.LastFinalizedTimestamp.Store(sourceLogFinalizedAt)
	}

	for _, msgLog := range msgLogs {
		seqNum := msgLog.Message.SequenceNumber
		var reqStat *testreporters.RequestStat
		lggr = lggr.With().Str("msgId ", fmt.Sprintf("0x%x", msgLog.Message.MessageId[:])).Logger()
		for _, stat := range stats {
			if stat.SeqNum == seqNum {
				reqStat = stat
				break
			}
		}
		if reqStat == nil {
			return fmt.Errorf("could not find request stat for seq number %d", seqNum)
		}
		// wait for
		// - CommitStore to increase the seq number,
		err = c.Lane.Dest.AssertSeqNumberExecuted(lggr, seqNum, c.CallTimeOut, sourceLogFinalizedAt, reqStat)
		if err != nil {
			return err
		}
		// wait for ReportAccepted event
		commitReport, reportAcceptedAt, err := c.Lane.Dest.AssertEventReportAccepted(lggr, seqNum, c.CallTimeOut, sourceLogFinalizedAt, reqStat)
		if err != nil || commitReport == nil {
			return err
		}
		blessedAt, err := c.Lane.Dest.AssertReportBlessed(lggr, seqNum, c.CallTimeOut, *commitReport, reportAcceptedAt, reqStat)
		if err != nil {
			return err
		}
		_, err = c.Lane.Dest.AssertEventExecutionStateChanged(lggr, seqNum, c.CallTimeOut, blessedAt, reqStat, testhelpers.ExecutionStateSuccess)
		if err != nil {
			return err
		}
	}

	return nil
}
