package ccip

import (
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/logger"
)

const (
	TOLL_CONSTANT_MESSAGE_PART_BYTES = 20 + // receiver
		20 + // sender
		2 + // chain id
		8 + // sequence number
		32 + // gas limit
		20 + // fee token address
		32 // fee token amount
	TOLL_EXECUTION_STATE_PROCESSING_OVERHEAD_GAS = 2_100 + // COLD_SLOAD_COST for first reading the state
		20_000 + // SSTORE_SET_GAS for writing from 0 (untouched) to non-zero (in-progress)
		100 //# SLOAD_GAS = WARM_STORAGE_READ_COST for rewriting from non-zero (in-progress) to non-zero (success/failure)
)

// Onchain: we bill deterministically for tolls so that we can notify clients how much of a refund they get.
// Offchain: we compute the max overhead gas to determine msg executability.
func overheadGasToll(merkleGasShare uint64, tollMsg evm_2_evm_toll_onramp.CCIPEVM2EVMTollMessage) uint64 {
	messageBytes := TOLL_CONSTANT_MESSAGE_PART_BYTES +
		(EVM_ADDRESS_LENGTH_BYTES+EVM_WORD_BYTES)*len(tollMsg.TokensAndAmounts) + // token address (address) + token amount (uint256)
		len(tollMsg.Data)
	messageCallDataGas := uint64(messageBytes * CALLDATA_GAS_PER_BYTE)

	// Rate limiter only limits value in tokens. It's not called if there are no
	// tokens in the message.
	rateLimiterOverhead := uint64(0)
	if len(tollMsg.TokensAndAmounts) >= 1 {
		rateLimiterOverhead = RATE_LIMITER_OVERHEAD_GAS
	}

	return messageCallDataGas +
		merkleGasShare +
		TOLL_EXECUTION_STATE_PROCESSING_OVERHEAD_GAS +
		PER_TOKEN_OVERHEAD_GAS*uint64(len(tollMsg.TokensAndAmounts)+1) + // All tokens plus fee token
		rateLimiterOverhead +
		EXTERNAL_CALL_OVERHEAD_GAS
}

func maxGasOverHeadGasToll(numMsgs int, tollMsg evm_2_evm_toll_onramp.CCIPEVM2EVMTollMessage) uint64 {
	merkleProofBytes := (math.Ceil(math.Log2(float64(numMsgs)))+2)*32 +
		(1+2)*32 // only ever one outer root hash
	merkleGasShare := uint64(merkleProofBytes * CALLDATA_GAS_PER_BYTE)
	return overheadGasToll(merkleGasShare, tollMsg)
}

func maxTollCharge(gasPrice *big.Int, subTokenPerFeeCoin *big.Int, totalGasLimit uint64) *big.Int {
	return new(big.Int).Div(new(big.Int).Mul(new(big.Int).Mul(big.NewInt(int64(totalGasLimit)), gasPrice), subTokenPerFeeCoin), big.NewInt(1e18))
}

type TollBatchBuilder struct {
	tollABI         abi.ABI
	eventSignatures EventSignatures
	lggr            logger.Logger
}

func NewTollBatchBuilder(lggr logger.Logger, eventSignatures EventSignatures) *TollBatchBuilder {
	tollABI, _ := abi.JSON(strings.NewReader(evm_2_evm_toll_onramp.EVM2EVMTollOnRampABI))
	return &TollBatchBuilder{
		tollABI:         tollABI,
		eventSignatures: eventSignatures,
		lggr:            lggr,
	}
}

func (tb *TollBatchBuilder) parseLog(log types.Log) (*evm_2_evm_toll_onramp.EVM2EVMTollOnRampCCIPSendRequested, error) {
	event := new(evm_2_evm_toll_onramp.EVM2EVMTollOnRampCCIPSendRequested)
	err := bind.NewBoundContract(common.Address{}, tb.tollABI, nil, nil, nil).UnpackLog(event, "CCIPSendRequested", log)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (tb *TollBatchBuilder) BuildBatch(
	srcToDst map[common.Address]common.Address,
	msgs []logpoller.Log,
	executed map[uint64]struct{},
	batchGasLimit uint64,
	gasPrice *big.Int,
	tollTokensPerFeeCoin map[common.Address]*big.Int,
	inflight []InflightExecutionReport,
	aggregateTokenLimit *big.Int,
	tokenLimitPrices map[common.Address]*big.Int,
) (executableSeqNrs []uint64, executedAllMessages bool) {
	inflightSeqNrs, inflightAggregateValue, err := tb.inflight(inflight, tokenLimitPrices, srcToDst)
	if err != nil {
		tb.lggr.Errorw("Unexpected error computing inflight values", "err", err)
		return []uint64{}, false
	}
	aggregateTokenLimit.Sub(aggregateTokenLimit, inflightAggregateValue)
	executedAllMessages = true
	for _, msg := range msgs {
		tollMsg, err2 := tb.parseLog(types.Log{
			Topics: msg.GetTopics(),
			Data:   msg.Data,
		})
		if err2 != nil {
			tb.lggr.Errorw("unable to parse message", "err", err2, "msg", msg)
			// Unable to parse so don't mark as executed
			executedAllMessages = false
			continue
		}
		if _, executed := executed[tollMsg.Message.SequenceNumber]; executed {
			tb.lggr.Infow("Skipping message already executed", "seqNr", tollMsg.Message.SequenceNumber)
			continue
		}
		executedAllMessages = false
		if _, inflight := inflightSeqNrs[tollMsg.Message.SequenceNumber]; inflight {
			tb.lggr.Infow("Skipping message already inflight", "seqNr", tollMsg.Message.SequenceNumber)
			continue
		}
		var tokens []common.Address
		var amounts []*big.Int
		for i := 0; i < len(tollMsg.Message.TokensAndAmounts); i++ {
			tokens = append(tokens, tollMsg.Message.TokensAndAmounts[i].Token)
			amounts = append(amounts, tollMsg.Message.TokensAndAmounts[i].Amount)
		}
		msgValue, err := aggregateTokenValue(tokenLimitPrices, srcToDst, tokens, amounts)
		if err != nil {
			tb.lggr.Errorw("Skipping message unable to compute aggregate value", "err", err)
			continue
		}
		// if token limit is smaller than message value skip message
		if aggregateTokenLimit.Cmp(msgValue) == -1 {
			continue
		}
		// Check solvency
		maxGasOverhead := maxGasOverHeadGasToll(len(msgs), tollMsg.Message)
		totalGasLimit := tollMsg.Message.GasLimit.Uint64() + maxGasOverhead
		// Check sufficient gas in batch
		if batchGasLimit < totalGasLimit {
			tb.lggr.Infow("Insufficient remaining gas in batch limit", "gasLimit", batchGasLimit, "totalGasLimit", totalGasLimit)
			continue
		}
		if _, ok := srcToDst[tollMsg.Message.FeeTokenAndAmount.Token]; !ok {
			tb.lggr.Errorw("Unknown fee token", "token", tollMsg.Message.FeeTokenAndAmount, "supported", srcToDst)
			continue
		}
		maxCharge := maxTollCharge(gasPrice, tollTokensPerFeeCoin[srcToDst[tollMsg.Message.FeeTokenAndAmount.Token]], totalGasLimit)
		if tollMsg.Message.FeeTokenAndAmount.Amount.Cmp(maxCharge) < 0 {
			tb.lggr.Infow("Insufficient fee token to execute msg", "balance", tollMsg.Message.FeeTokenAndAmount.Amount, "maxCharge", maxCharge, "maxGasOverhead", maxGasOverhead)
			continue
		}
		batchGasLimit -= totalGasLimit
		aggregateTokenLimit.Sub(aggregateTokenLimit, msgValue)
		tb.lggr.Infow("Adding toll msg to batch", "seqNum", tollMsg.Message.SequenceNumber, "maxCharge", maxCharge, "maxGasOverhead", maxGasOverhead)
		executableSeqNrs = append(executableSeqNrs, tollMsg.Message.SequenceNumber)
	}
	return executableSeqNrs, executedAllMessages
}

func (tb *TollBatchBuilder) inflight(
	inflight []InflightExecutionReport,
	tokenLimitPrices map[common.Address]*big.Int,
	srcToDst map[common.Address]common.Address,
) (map[uint64]struct{}, *big.Int, error) {
	inflightSeqNrs := make(map[uint64]struct{})
	inflightAggregateValue := big.NewInt(0)
	for _, rep := range inflight {
		for _, seqNr := range rep.report.SequenceNumbers {
			inflightSeqNrs[seqNr] = struct{}{}
		}
		for _, encMsg := range rep.report.EncodedMessages {
			msg, err := tb.parseLog(types.Log{
				// Note this needs to change if we start indexing things.
				Topics: []common.Hash{tb.eventSignatures.SendRequested},
				Data:   encMsg,
			})
			if err != nil {
				return nil, nil, err
			}
			var tokens []common.Address
			var amounts []*big.Int
			for i := 0; i < len(msg.Message.TokensAndAmounts); i++ {
				tokens = append(tokens, msg.Message.TokensAndAmounts[i].Token)
				amounts = append(amounts, msg.Message.TokensAndAmounts[i].Amount)
			}
			msgValue, err := aggregateTokenValue(tokenLimitPrices, srcToDst, tokens, amounts)
			if err != nil {
				return nil, nil, err
			}
			inflightAggregateValue.Add(inflightAggregateValue, msgValue)
		}
	}
	return inflightSeqNrs, inflightAggregateValue, nil
}
