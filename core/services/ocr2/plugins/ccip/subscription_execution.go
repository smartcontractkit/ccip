package ccip

import (
	"math/big"
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
	balances map[common.Address]*big.Int
}

func (sbc SubscriptionBalanceCache) Balance(addr common.Address) *big.Int {
	return sbc.balances[addr]
}

type SubscriptionBatch struct {
	remainingGasLimit  uint64
	maxGasPrice        uint64
	subTokenPerFeeCoin *big.Int
	cache              SubscriptionBalanceCache
	reservedBalance    map[common.Address]*big.Int
	seqNrs             []uint64
}

func NewSubscriptionBatch(maxGasLimit uint64, maxGasPrice uint64, subTokenPerFeeCoin *big.Int, cache SubscriptionBalanceCache, reservedBalance map[common.Address]*big.Int) *SubscriptionBatch {
	return &SubscriptionBatch{
		remainingGasLimit:  maxGasLimit,
		maxGasPrice:        maxGasPrice,
		subTokenPerFeeCoin: subTokenPerFeeCoin,
		cache:              cache,
		reservedBalance:    reservedBalance,
	}
}

// TODO: Need sub message
func (seb *SubscriptionBatch) Add(msg *evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampCCIPSendRequested) bool {
	if seb.remainingGasLimit < (msg.Message.GasLimit.Uint64() + MAX_OVERHEAD_GAS_SUBSCRIPTION) {
		return false
	}
	subBalance := seb.cache.Balance(msg.Message.Receiver)
	if subBalance == nil {
		// TODO error/ Implement proper sub cache
		subBalance = big.NewInt(1e18)
	}
	reserved, hasReserved := seb.reservedBalance[msg.Message.Receiver]
	if hasReserved {
		subBalance.Sub(subBalance, reserved)
	}
	if subBalance.Cmp(maxSubCharge(seb.maxGasPrice, seb.subTokenPerFeeCoin, msg)) == -1 {
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

func (sb *SubscriptionBatchBuilder) BuildBatch(srcToDst map[common.Address]common.Address, msgs []logpoller.Log, executed map[uint64]struct{}, gasLimit uint64, gasPrice uint64, tokensPerFeeCoin map[common.Address]*big.Int, inflight []InflightExecutionReport) []uint64 {
	subTokenPerFeeCoin := tokensPerFeeCoin[sb.subFeeToken]
	if subTokenPerFeeCoin == nil {
		sb.lggr.Errorf("Fee token price not found for token: %s", sb.subFeeToken.Hex())
	}
	inflightSeqNrs, reserved, err := sb.inflightAndReservedBalances(gasPrice, subTokenPerFeeCoin, inflight)
	if err != nil {
		// Log error
		return []uint64{}
	}
	balanceCache := SubscriptionBalanceCache{
		balances: map[common.Address]*big.Int{},
	}
	subBatch := NewSubscriptionBatch(gasLimit, gasPrice, subTokenPerFeeCoin, balanceCache, reserved)
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

func maxSubCharge(maxGasPrice uint64, subTokenPerFeeCoin *big.Int, msg *evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampCCIPSendRequested) *big.Int {
	totalGasLimit := new(big.Int).Add(msg.Message.GasLimit, big.NewInt(MAX_OVERHEAD_GAS_SUBSCRIPTION))
	// TODO: Adjust once subscription contracts available.

	return new(big.Int).Div(new(big.Int).Mul(new(big.Int).Mul(totalGasLimit, big.NewInt(int64(maxGasPrice))), subTokenPerFeeCoin), big.NewInt(1e18))
}

func (sb *SubscriptionBatchBuilder) inflightAndReservedBalances(maxGasPrice uint64, subTokenPerFeeCoin *big.Int, inflight []InflightExecutionReport) (map[uint64]struct{}, map[common.Address]*big.Int, error) {
	reserved := make(map[common.Address]*big.Int)
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

			if reserved[subMsg.Message.Receiver] == nil {
				reserved[subMsg.Message.Receiver] = maxSubCharge(maxGasPrice, subTokenPerFeeCoin, subMsg)
			} else {
				reserved[subMsg.Message.Receiver].Add(reserved[subMsg.Message.Receiver], maxSubCharge(maxGasPrice, subTokenPerFeeCoin, subMsg))
			}
		}
	}
	return inflightSeqNrs, reserved, nil
}
