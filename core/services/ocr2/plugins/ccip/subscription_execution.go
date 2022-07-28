package ccip

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/evm_2_evm_subscription_onramp"
	"github.com/smartcontractkit/chainlink/core/logger"
)

const (
	MAX_OVERHEAD_GAS_SUBSCRIPTION = 1 // TODO
)

type SubscriptionBalanceCache struct {
	balances map[common.Address]uint64
}

func (sbc SubscriptionBalanceCache) Balance(addr common.Address) uint64 {
	return sbc.balances[addr]
}

type SubscriptionBatch struct {
	remainingGasLimit  uint64
	maxGasPrice        uint64
	subTokenPerFeeCoin uint64
	cache              SubscriptionBalanceCache
	reservedBalance    map[common.Address]uint64
	seqNrs             []uint64
}

func NewSubscriptionBatch(maxGasLimit uint64, maxGasPrice uint64, subTokenPerFeeCoin uint64, cache SubscriptionBalanceCache, reservedBalance map[common.Address]uint64) *SubscriptionBatch {
	return &SubscriptionBatch{
		remainingGasLimit:  maxGasLimit,
		maxGasPrice:        maxGasPrice,
		subTokenPerFeeCoin: subTokenPerFeeCoin,
		cache:              cache,
		reservedBalance:    reservedBalance,
	}
}

func (seb *SubscriptionBatch) Add(msg *evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampCCIPSendRequested) bool {
	if seb.remainingGasLimit-(msg.Message.GasLimit.Uint64()+MAX_OVERHEAD_GAS_SUBSCRIPTION) < 0 {
		return false
	}
	subBalance := seb.cache.Balance(msg.Message.Receiver)
	reserved, hasReserved := seb.reservedBalance[msg.Message.Receiver]
	if hasReserved {
		subBalance -= reserved
	}
	if subBalance < maxSubCharge(seb.maxGasPrice, seb.subTokenPerFeeCoin, msg) {
		return false
	}
	seb.remainingGasLimit -= msg.Message.GasLimit.Uint64()
	seb.seqNrs = append(seb.seqNrs, msg.Message.SequenceNumber)
	return true
}

func (seb *SubscriptionBatch) SeqNrs() []uint64 {
	return seb.seqNrs
}

type SubscriptionBatchBuilder struct {
	lggr        logger.Logger
	subABI      abi.ABI
	subFeeToken common.Address
}

func NewSubscriptionBatchBuilder(lggr logger.Logger, subFeeToken common.Address) *SubscriptionBatchBuilder {
	subABI, _ := abi.JSON(strings.NewReader(evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampABI))
	return &SubscriptionBatchBuilder{lggr: lggr, subABI: subABI, subFeeToken: subFeeToken}
}

func (sb *SubscriptionBatchBuilder) BuildBatch(srcToDst map[common.Address]common.Address, msgs []logpoller.Log, executed map[uint64]struct{}, gasLimit uint64, gasPrice uint64, tokensPerFeeCoin map[common.Address]uint64, inflight []InflightExecutionReport) []uint64 {
	subTokenPerFeeCoin := tokensPerFeeCoin[sb.subFeeToken]
	inflightSeqNrs, reserved, err := sb.inflightAndReservedBalances(gasPrice, subTokenPerFeeCoin, inflight)
	if err != nil {
		// Log error
		return []uint64{}
	}
	subBatch := NewSubscriptionBatch(gasLimit, gasPrice, subTokenPerFeeCoin, SubscriptionBalanceCache{}, reserved)
	haveOne := false
	for _, msg := range msgs {
		subMsg, err := sb.parseLog(types.Log{
			Topics: msg.GetTopics(),
			Data:   msg.Data,
		})
		if err != nil {
			sb.lggr.Errorw("unable to parse message", "err", err, "msg", msg)
			continue
		}
		if _, inflight := inflightSeqNrs[subMsg.Message.SequenceNumber]; inflight {
			continue
		}
		if _, executed := executed[subMsg.Message.SequenceNumber]; executed {
			continue
		}
		added := subBatch.Add(subMsg)
		if !added && haveOne {
			break
		}
		if added && !haveOne {
			haveOne = true
		}
	}
	return subBatch.SeqNrs()
}

func (sb *SubscriptionBatchBuilder) parseLog(log types.Log) (*evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampCCIPSendRequested, error) {
	event := new(evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampCCIPSendRequested)
	err := bind.NewBoundContract(common.Address{}, sb.subABI, nil, nil, nil).UnpackLog(event, "CCIPSendRequested", log)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func maxSubCharge(maxGasPrice uint64, subTokenPerFeeCoin uint64, msg *evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampCCIPSendRequested) uint64 {
	totalGasLimit := msg.Message.GasLimit.Uint64() + MAX_OVERHEAD_GAS_SUBSCRIPTION
	// TODO: Adjust once subscription contracts available.
	return totalGasLimit * maxGasPrice * subTokenPerFeeCoin
}

func (sb *SubscriptionBatchBuilder) inflightAndReservedBalances(maxGasPrice uint64, subTokenPerFeeCoin uint64, inflight []InflightExecutionReport) (map[uint64]struct{}, map[common.Address]uint64, error) {
	reserved := make(map[common.Address]uint64)
	inflightSeqNrs := make(map[uint64]struct{})
	for _, r := range inflight {
		for _, encMsg := range r.report.EncodedMessages {
			subMsg, err := sb.parseLog(types.Log{
				// Note this needs to change if we start indexing things.
				Topics: []common.Hash{CCIPSubSendRequested},
				Data:   encMsg,
			})
			if err != nil {
				return nil, nil, err
			}
			inflightSeqNrs[subMsg.Message.SequenceNumber] = struct{}{}
			reserved[subMsg.Message.Receiver] += maxSubCharge(maxGasPrice, subTokenPerFeeCoin, subMsg)
		}
	}
	return inflightSeqNrs, reserved, nil
}
