package load

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/logging"
	"github.com/smartcontractkit/wasp"

	"github.com/smartcontractkit/chainlink/integration-tests/ccip-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/ccip-tests/contracts"
	"github.com/smartcontractkit/chainlink/integration-tests/ccip-tests/testreporters"
	"github.com/smartcontractkit/chainlink/integration-tests/ccip-tests/testsetups"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/testhelpers"
)

// CCIPMultiCallLoadGenerator represents a load generator for the CCIP lanes originating from same network
// The purpose of this load generator is to group ccip-send calls for the CCIP lanes originating from same network
// This is to avoid the scenario of hitting rpc rate limit for the same network if the load generator is sending
// too many ccip-send calls to the same network hitting the rpc rate limit
type CCIPMultiCallLoadGenerator struct {
	t                       *testing.T
	logger                  zerolog.Logger
	client                  blockchain.EVMClient
	E2ELoads                map[string]*CCIPE2ELoad
	MultiCall               string
	NoOfRequestsPerUnitTime int
}

type ReturnValues struct {
	Msgs  []contracts.CCIPMsgData
	Stats []*testreporters.RequestStat
}

func NewMultiCallLoadGenerator(t *testing.T, lanes []*actions.CCIPLane) (*CCIPMultiCallLoadGenerator, error) {
	// check if all lanes are from same network
	source := lanes[0].SourceChain.GetChainID()
	multiCall := lanes[0].SrcNetworkLaneCfg.Multicall
	if multiCall == "" {
		return nil, fmt.Errorf("multicall address cannot be empty")
	}
	for i := 0; i < len(lanes); i++ {
		if source != lanes[i].SourceChain.GetChainID() {
			return nil, fmt.Errorf("all lanes should be from same network")
		}
		if lanes[i].SrcNetworkLaneCfg.Multicall != multiCall {
			return nil, fmt.Errorf("multicall address should be same for all lanes")
		}
	}
	client := lanes[0].SourceChain
	lggr := logging.GetTestLogger(t).With().Str("Source Network", client.GetNetworkName()).Logger()
	return &CCIPMultiCallLoadGenerator{
		t:         t,
		client:    client,
		MultiCall: multiCall,
		logger:    lggr,
	}, nil
}

func (m *CCIPMultiCallLoadGenerator) BeforeAll(testCfg *testsetups.CCIPTestConfig, lanes []*actions.CCIPLane) {
	for _, lane := range lanes {
		ccipLoad := NewCCIPLoad(testCfg.Test, lane, testCfg.TestGroupInput.PhaseTimeout.Duration(), 100000)
		ccipLoad.BeforeAllCall(testCfg.TestGroupInput.MsgType)
		m.E2ELoads[fmt.Sprintf("%s-%s", lane.SourceNetworkName, lane.DestNetworkName)] = ccipLoad
	}
}

func (m *CCIPMultiCallLoadGenerator) Call(_ *wasp.Generator) *wasp.CallResult {
	res := &wasp.CallResult{}
	msgs, returnValuesByDest, allStats, err := m.MergeCalls()
	if err != nil {
		res.Error = err.Error()
		res.Failed = true
		return res
	}
	startTime := time.Now().UTC()
	// for now we are using all ccip-sends with native
	sendTx, err := contracts.MultiCallCCIP(m.client, m.MultiCall, msgs, true)
	if err != nil {
		res.Error = err.Error()
		res.Failed = true
		res.Data = allStats
		return res
	}

	lggr := m.logger.With().Str("Msg Tx", sendTx.Hash().String()).Logger()
	txConfirmationTime := time.Now().UTC()
	rcpt, err1 := bind.WaitMined(context.Background(), m.client.DeployBackend(), sendTx)
	if err1 == nil {
		hdr, err1 := m.client.HeaderByNumber(context.Background(), rcpt.BlockNumber)
		if err1 == nil {
			txConfirmationTime = hdr.Timestamp
		}
	}
	var gasUsed uint64
	if rcpt != nil {
		gasUsed = rcpt.GasUsed
	}
	for i, allStats := range stats {
		for _, stat := range allStats {
			stat.UpdateState(lggr, 0, testreporters.TX, startTime.Sub(txConfirmationTime), testreporters.Success,
				testreporters.TransactionStats{
					Fee:                msgs[i].Fee.String(),
					GasUsed:            gasUsed,
					TxHash:             sendTx.Hash().Hex(),
					NoOfTokensSent:     len(msgs[i].Msg.TokenAmounts),
					MessageBytesLength: len(msgs[i].Msg.Data),
				})
		}
	}

	// wait for
	// - CCIPSendRequested Event log to be generated,
	for _, allstates := range stats {
		key := fmt.Sprintf("%s-%s", stat.SourceNetwork, stat.DestNetwork)
		c, ok := m.E2ELoads[key]
		if !ok {
			res.Error = fmt.Sprintf("load for %s not found", key)
			res.Failed = true
			return res
		}
		msgLogs, sourceLogTime, err := c.Lane.Source.AssertEventCCIPSendRequested(lggr, sendTx.Hash().Hex(), c.CallTimeOut, txConfirmationTime, []*testreporters.RequestStat{stat})

		if err != nil || msgLogs == nil || len(msgLogs) == 0 {
			res.Error = err.Error()
			res.Data = allStats
			res.Failed = true
			return res
		}
		msgLog := msgLogs[0]
		sentMsg := msgLog.Message
		seqNum := sentMsg.SequenceNumber
		lggr = lggr.With().Str("msgId ", fmt.Sprintf("0x%x", sentMsg.MessageId[:])).Logger()

		lstFinalizedBlock := c.LastFinalizedTxBlock.Load()
		var sourceLogFinalizedAt time.Time
		// if the finality tag is enabled and the last finalized block is greater than the block number of the message
		// consider the message finalized
		if c.Lane.Source.Common.ChainClient.GetNetworkConfig().FinalityDepth == 0 &&
			lstFinalizedBlock != 0 && lstFinalizedBlock > msgLog.Raw.BlockNumber {
			sourceLogFinalizedAt = c.LastFinalizedTimestamp.Load()
			stat.UpdateState(lggr, seqNum, testreporters.SourceLogFinalized,
				sourceLogFinalizedAt.Sub(sourceLogTime), testreporters.Success,
				testreporters.TransactionStats{
					TxHash:           msgLog.Raw.TxHash.String(),
					FinalizedByBlock: strconv.FormatUint(lstFinalizedBlock, 10),
					FinalizedAt:      sourceLogFinalizedAt.String(),
				})
		} else {
			var finalizingBlock uint64
			sourceLogFinalizedAt, finalizingBlock, err = c.Lane.Source.AssertSendRequestedLogFinalized(
				lggr, sendTx.Hash(), sourceLogTime, []*testreporters.RequestStat{stats})
			if err != nil {
				res.Error = err.Error()
				res.Data = allStats
				res.Failed = true
				return res
			}
			c.LastFinalizedTxBlock.Store(finalizingBlock)
			c.LastFinalizedTimestamp.Store(sourceLogFinalizedAt)
		}

		// wait for
		// - CommitStore to increase the seq number,
		err = c.Lane.Dest.AssertSeqNumberExecuted(lggr, seqNum, c.CallTimeOut, sourceLogFinalizedAt, stat)
		if err != nil {
			res.Error = err.Error()
			res.Data = allStats
			res.Failed = true
			return res
		}
		// wait for ReportAccepted event
		commitReport, reportAcceptedAt, err := c.Lane.Dest.AssertEventReportAccepted(lggr, seqNum, c.CallTimeOut, sourceLogFinalizedAt, stat)
		if err != nil || commitReport == nil {
			res.Error = err.Error()
			res.Data = allStats
			res.Failed = true
			return res
		}
		blessedAt, err := c.Lane.Dest.AssertReportBlessed(lggr, seqNum, c.CallTimeOut, *commitReport, reportAcceptedAt, stat)
		if err != nil {
			res.Error = err.Error()
			res.Data = allStats
			res.Failed = true
			return res
		}
		_, err = c.Lane.Dest.AssertEventExecutionStateChanged(lggr, seqNum, c.CallTimeOut, blessedAt, stat, testhelpers.ExecutionStateSuccess)
		if err != nil {
			res.Error = err.Error()
			res.Data = allStats
			res.Failed = true
			return res
		}
	}

	res.Data = allStats
	return res
}

func (m *CCIPMultiCallLoadGenerator) MergeCalls() ([]contracts.CCIPMsgData, map[string]ReturnValues, []*testreporters.RequestStat, error) {
	var ccipMsgs []contracts.CCIPMsgData
	statDetails := make(map[string]ReturnValues)
	var allStats []*testreporters.RequestStat
	for _, e2eLoad := range m.E2ELoads {
		destChainSelector, err := chain_selectors.SelectorFromChainId(e2eLoad.Lane.Source.DestinationChainId)
		if err != nil {
			return ccipMsgs, statDetails, allStats, err
		}

		allFee := big.NewInt(0)
		var allStatsForDest []*testreporters.RequestStat
		var allMsgsForDest []contracts.CCIPMsgData
		for i := 0; i < m.NoOfRequestsPerUnitTime; i++ {
			msg, stats := e2eLoad.CCIPMsg()
			msg.FeeToken = common.Address{}
			fee, err := e2eLoad.Lane.Source.Common.Router.GetFee(destChainSelector, msg)
			if err != nil {
				return ccipMsgs, statDetails, allStats, err
			}
			// transfer fee to the multicall address
			if msg.FeeToken != (common.Address{}) {
				allFee = new(big.Int).Add(allFee, fee)
			}
			msgData := contracts.CCIPMsgData{
				RouterAddr:    e2eLoad.Lane.Source.Common.Router.EthAddress,
				ChainSelector: destChainSelector,
				Msg:           msg,
				Fee:           fee,
			}
			ccipMsgs = append(ccipMsgs, msgData)
			allStats = append(allStats, stats)
			allStatsForDest = append(allStatsForDest, stats)
			allMsgsForDest = append(allMsgsForDest, msgData)
		}
		statDetails[e2eLoad.Lane.DestNetworkName] = ReturnValues{
			Stats: allStatsForDest,
			Msgs:  allMsgsForDest,
		}
		// transfer fee to the multicall address
		if allFee.Cmp(big.NewInt(0)) > 0 {
			if err := e2eLoad.Lane.Source.Common.FeeToken.Transfer(e2eLoad.Lane.Source.Common.MulticallContract.Hex(), allFee); err != nil {
				return ccipMsgs, statDetails, allStats, err
			}
		}
	}
	return ccipMsgs, statDetails, allStats, nil
}
