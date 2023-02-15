package ccip

import (
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/core/logger"
)

type BatchBuilder struct {
	abi             abi.ABI
	eventSignatures EventSignatures
	lggr            logger.Logger
	ramp            *evm_2_evm_offramp.EVM2EVMOffRamp
}

const (
	CONSTANT_MESSAGE_PART_BYTES = 32 + // sourceChainId
		32 + // feeTokenAmount
		8 + // sequenceNumber
		20 + // sender
		32 + // gas limit
		8 + // nonce
		1 + // strict
		20 + // receiver
		32 // fee token
	EXECUTION_STATE_PROCESSING_OVERHEAD_GAS = 2_100 + // COLD_SLOAD_COST for first reading the state
		20_000 + // SSTORE_SET_GAS for writing from 0 (untouched) to non-zero (in-progress)
		100 //# SLOAD_GAS = WARM_STORAGE_READ_COST for rewriting from non-zero (in-progress) to non-zero (success/failure)
)

// Offchain: we compute the max overhead gas to determine msg executability.
func overheadGas(geMsg evm_2_evm_onramp.InternalEVM2EVMMessage) uint64 {
	messageBytes := CONSTANT_MESSAGE_PART_BYTES +
		(EVM_ADDRESS_LENGTH_BYTES+EVM_WORD_BYTES)*len(geMsg.TokenAmounts) + // token address (address) + token amount (uint256)
		len(geMsg.Data)
	messageCallDataGas := uint64(messageBytes * CALLDATA_GAS_PER_BYTE)

	// Rate limiter only limits value in tokens. It's not called if there are no
	// tokens in the message.
	rateLimiterOverhead := uint64(0)
	if len(geMsg.TokenAmounts) >= 1 {
		rateLimiterOverhead = RATE_LIMITER_OVERHEAD_GAS
	}

	return messageCallDataGas +
		EXECUTION_STATE_PROCESSING_OVERHEAD_GAS +
		PER_TOKEN_OVERHEAD_GAS*uint64(len(geMsg.TokenAmounts)) +
		rateLimiterOverhead +
		EXTERNAL_CALL_OVERHEAD_GAS
}

func maxGasOverHeadGas(numMsgs int, geMsg evm_2_evm_onramp.InternalEVM2EVMMessage) uint64 {
	merkleProofBytes := (math.Ceil(math.Log2(float64(numMsgs)))+2)*32 + (1+2)*32 // only ever one outer root hash
	merkleGasShare := uint64(merkleProofBytes * CALLDATA_GAS_PER_BYTE)
	gasFeeShare := uint64(PER_TOKEN_OVERHEAD_GAS / numMsgs)

	return overheadGas(geMsg) + merkleGasShare + gasFeeShare
}

func NewBatchBuilder(lggr logger.Logger, eventSignatures EventSignatures, ramp *evm_2_evm_offramp.EVM2EVMOffRamp) *BatchBuilder {
	geABI, _ := abi.JSON(strings.NewReader(evm_2_evm_onramp.EVM2EVMOnRampABI))
	return &BatchBuilder{
		abi:             geABI,
		eventSignatures: eventSignatures,
		lggr:            lggr,
		ramp:            ramp,
	}
}

func (tb *BatchBuilder) parseLog(log types.Log) (*evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested, error) {
	event := new(evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested)
	err := bind.NewBoundContract(common.Address{}, tb.abi, nil, nil, nil).UnpackLog(event, "CCIPSendRequested", log)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (tb *BatchBuilder) BuildBatch(
	srcToDst map[common.Address]common.Address,
	msgs []logpoller.Log,
	executed map[uint64]struct{},
	batchGasLimit uint64,
	gasPrice *big.Int,
	tokensPerFeeCoin map[common.Address]*big.Int,
	inflight []InflightInternalExecutionReport,
	aggregateTokenLimit *big.Int,
	tokenLimitPrices map[common.Address]*big.Int,
) (executableSeqNrs []uint64, executedAllMessages bool) {
	inflightSeqNrs, inflightAggregateValue, maxInflightSenderNonces, err := tb.inflight(inflight, tokenLimitPrices, srcToDst)
	if err != nil {
		tb.lggr.Errorw("Unexpected error computing inflight values", "err", err)
		return []uint64{}, false
	}
	aggregateTokenLimit.Sub(aggregateTokenLimit, inflightAggregateValue)
	executedAllMessages = true
	expectedNonces := make(map[common.Address]uint64)
	for _, msg := range msgs {
		msg, err2 := tb.parseLog(types.Log{
			Topics: msg.GetTopics(),
			Data:   msg.Data,
		})
		if err2 != nil {
			tb.lggr.Errorw("unable to parse message", "err", err2, "msg", msg)
			// Unable to parse so don't mark as executed
			executedAllMessages = false
			continue
		}
		lggr := tb.lggr.With("messageID", hexutil.Encode(msg.Message.MessageId[:]))
		if _, executed := executed[msg.Message.SequenceNumber]; executed {
			lggr.Infow("Skipping message already executed", "seqNr", msg.Message.SequenceNumber)
			continue
		}
		executedAllMessages = false
		if _, inflight := inflightSeqNrs[msg.Message.SequenceNumber]; inflight {
			lggr.Infow("Skipping message already inflight", "seqNr", msg.Message.SequenceNumber)
			continue
		}
		if _, ok := expectedNonces[msg.Message.Sender]; !ok {
			// First message in batch, need to populate expected nonce
			if maxInflight, ok := maxInflightSenderNonces[msg.Message.Sender]; ok {
				// Sender already has inflight nonce, populate from there
				expectedNonces[msg.Message.Sender] = maxInflight + 1
			} else {
				// Nothing inflight take from chain.
				// Chain holds expected next nonce.
				nonce, err := tb.ramp.GetSenderNonce(nil, msg.Message.Sender)
				if err != nil {
					lggr.Errorw("unable to get sender nonce", "err", err)
					continue
				}
				expectedNonces[msg.Message.Sender] = nonce + 1
			}
		}
		// Check expected nonce is valid

		if msg.Message.Nonce != expectedNonces[msg.Message.Sender] {
			lggr.Warnw("Skipping message invalid nonce", "have", msg.Message.Nonce, "want", expectedNonces[msg.Message.Sender])
			continue
		}

		var tokens []common.Address
		var amounts []*big.Int
		for i := 0; i < len(msg.Message.TokenAmounts); i++ {
			tokens = append(tokens, msg.Message.TokenAmounts[i].Token)
			amounts = append(amounts, msg.Message.TokenAmounts[i].Amount)
		}
		msgValue, err := aggregateTokenValue(tokenLimitPrices, srcToDst, tokens, amounts)
		if err != nil {
			lggr.Errorw("Skipping message unable to compute aggregate value", "err", err)
			continue
		}
		// if token limit is smaller than message value skip message
		if aggregateTokenLimit.Cmp(msgValue) == -1 {
			continue
		}
		// TODO: fee boosting check, loss protection etc. For now we are just executing regardless
		totalGasLimit := msg.Message.GasLimit.Uint64() + maxGasOverHeadGas(len(msgs), msg.Message)
		// Check sufficient gas in batch
		if batchGasLimit < totalGasLimit {

			lggr.Infow("Insufficient remaining gas in batch limit", "gasLimit", batchGasLimit, "totalGasLimit", totalGasLimit)
			continue
		}
		batchGasLimit -= totalGasLimit
		aggregateTokenLimit.Sub(aggregateTokenLimit, msgValue)

		lggr.Infow("Adding msg to batch", "seqNum", msg.Message.SequenceNumber, "nonce", msg.Message.Nonce)
		executableSeqNrs = append(executableSeqNrs, msg.Message.SequenceNumber)
		expectedNonces[msg.Message.Sender] = msg.Message.Nonce + 1
	}
	return executableSeqNrs, executedAllMessages
}

func (tb *BatchBuilder) inflight(
	inflight []InflightInternalExecutionReport,
	tokenLimitPrices map[common.Address]*big.Int,
	srcToDst map[common.Address]common.Address,
) (map[uint64]struct{}, *big.Int, map[common.Address]uint64, error) {
	inflightSeqNrs := make(map[uint64]struct{})
	inflightAggregateValue := big.NewInt(0)
	maxInflightSenderNonces := make(map[common.Address]uint64)
	for _, rep := range inflight {
		for _, seqNr := range rep.seqNrs {
			inflightSeqNrs[seqNr] = struct{}{}
		}
		for _, encMsg := range rep.encMessages {
			msg, err := tb.parseLog(types.Log{
				// Note this needs to change if we start indexing things.
				Topics: []common.Hash{tb.eventSignatures.SendRequested},
				Data:   encMsg,
			})
			if err != nil {
				return nil, nil, nil, err
			}
			var tokens []common.Address
			var amounts []*big.Int
			for i := 0; i < len(msg.Message.TokenAmounts); i++ {
				tokens = append(tokens, msg.Message.TokenAmounts[i].Token)
				amounts = append(amounts, msg.Message.TokenAmounts[i].Amount)
			}
			msgValue, err := aggregateTokenValue(tokenLimitPrices, srcToDst, tokens, amounts)
			if err != nil {
				return nil, nil, nil, err
			}
			inflightAggregateValue.Add(inflightAggregateValue, msgValue)
			maxInflightSenderNonce, ok := maxInflightSenderNonces[msg.Message.Sender]
			if !ok || msg.Message.Nonce > maxInflightSenderNonce {
				maxInflightSenderNonces[msg.Message.Sender] = msg.Message.Nonce
			}
		}
	}
	return inflightSeqNrs, inflightAggregateValue, maxInflightSenderNonces, nil
}
