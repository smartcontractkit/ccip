package ccip

import (
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_subscription_onramp"
	"github.com/smartcontractkit/chainlink/core/logger"
)

const (
	SUBSCRIPTION_CONSTANT_MESSAGE_PART_BYTES = (20 + // receiver
		20 + // sender
		2 + // chain id
		8 + // sequence number
		32 + // gas limit
		32) //  nonce
	SUBSCRIPTION_EXECUTION_STATE_PROCESSING_OVERHEAD_GAS = (2_100 + // COLD_SLOAD_COST for first reading the state
		20_000 + // SSTORE_SET_GAS for writing from 0 (untouched) to non-zero (in-progress)
		100 + // SLOAD_GAS = WARM_STORAGE_READ_COST for rewriting from non-zero (in-progress) to non-zero (success/failure)
		2_100 + // COLD_SLOAD_COST for reading the nonce
		5_000) // SSTORE_RESET_GAS for incrementing the nonce from non-zero to non-zero
	SUBSCRIPTION_FEE_CHARGING = 5_000 // SSTORE_RESET_GAS for decreasing sub balance from non-zero to non-zero
)

type SubscriptionCache struct {
	router  *any_2_evm_subscription_offramp_router.Any2EVMSubscriptionOffRampRouter
	offramp *any_2_evm_subscription_offramp.EVM2EVMSubscriptionOffRamp
	strict  map[common.Address]bool
	lggr    logger.Logger
}

func NewSubscriptionCache(
	router *any_2_evm_subscription_offramp_router.Any2EVMSubscriptionOffRampRouter,
	offramp *any_2_evm_subscription_offramp.EVM2EVMSubscriptionOffRamp,
	lggr logger.Logger,
) SubscriptionCache {
	return SubscriptionCache{
		router:  router,
		offramp: offramp,
		strict:  make(map[common.Address]bool),
		lggr:    lggr,
	}
}

func (sbc *SubscriptionCache) Balance(addr common.Address) *big.Int {
	sub, err := sbc.router.GetSubscription(nil, addr)
	if err != nil {
		sbc.lggr.Errorw("No sub found", "addr", addr)
		return big.NewInt(0)
	}
	return sub.Balance
}

func (sbc *SubscriptionCache) Nonce(addr common.Address) uint64 {
	nonce, err := sbc.offramp.SReceiverToNonce(nil, addr)
	if err != nil {
		sbc.lggr.Errorw("Unable to get nonce for sub", "addr", addr)
		return 0
	}
	return nonce
}

func (sbc *SubscriptionCache) IsStrict(addr common.Address) (bool, error) {
	if _, ok := sbc.strict[addr]; !ok {
		sub, _ := sbc.router.GetSubscription(nil, addr)
		if sub.Receiver == [common.AddressLength]byte{} {
			return false, errors.Errorf("subscription does not exist for addr %v", addr)
		}
		sbc.strict[addr] = sub.StrictSequencing
	}
	return sbc.strict[addr], nil
}

func (sbc *SubscriptionCache) MostRecentExecution(addr common.Address) MessageExecutionState {
	// TODO as part of https://app.shortcut.com/chainlinklabs/story/51129/efficient-report-from-seq-num-lookup
	return MessageStateSuccess
}

func maxSubCharge(maxGasPrice uint64, subTokenPerFeeCoin *big.Int, totalGasLimit uint64) *big.Int {
	return new(big.Int).Div(new(big.Int).Mul(new(big.Int).Mul(big.NewInt(int64(totalGasLimit)), big.NewInt(int64(maxGasPrice))), subTokenPerFeeCoin), big.NewInt(1e18))
}

// Note: this is only used offchain.
// Onchain: we simply measure the gas usage and bill accordingly
// Offchain: we compute the max overhead gas to determine msg executability.
func overheadGasSubscription(merkleGasShare uint64, subMsg *evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampCCIPSendRequested) uint64 {
	messageBytes := SUBSCRIPTION_CONSTANT_MESSAGE_PART_BYTES +
		(EVM_ADDRESS_LENGTH_BYTES+EVM_WORD_BYTES)*len(subMsg.Message.Tokens) + // token address (address) + token amount (uint256)
		len(subMsg.Message.Data)
	messageCallDataGas := uint64(messageBytes * CALLDATA_GAS_PER_BYTE)
	return messageCallDataGas +
		merkleGasShare +
		SUBSCRIPTION_EXECUTION_STATE_PROCESSING_OVERHEAD_GAS +
		PER_TOKEN_OVERHEAD_GAS*uint64(len(subMsg.Message.Tokens)) +
		SUBSCRIPTION_FEE_CHARGING +
		RATE_LIMITER_OVERHEAD_GAS +
		EXTERNAL_CALL_OVERHEAD_GAS
}

func maxGasOverHeadGasSubscription(numMsgs int, subMsg *evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampCCIPSendRequested) uint64 {
	merkleProofBytes := (math.Ceil(math.Log2(float64(numMsgs)))+2)*32 +
		(1+2)*32 // only ever one outer root hash
	merkleGasShare := uint64(merkleProofBytes * CALLDATA_GAS_PER_BYTE)
	return overheadGasSubscription(merkleGasShare, subMsg)
}

type SubscriptionBatchBuilder struct {
	lggr        logger.Logger
	subABI      abi.ABI
	subFeeToken common.Address
	subCache    SubscriptionCache
}

func NewSubscriptionBatchBuilder(lggr logger.Logger, subFeeToken common.Address, subOffRamp *subOffRamp) *SubscriptionBatchBuilder {
	subABI, _ := abi.JSON(strings.NewReader(evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampABI))
	subCache := NewSubscriptionCache(subOffRamp.router, subOffRamp.EVM2EVMSubscriptionOffRamp, lggr)
	return &SubscriptionBatchBuilder{lggr: lggr, subABI: subABI, subFeeToken: subFeeToken, subCache: subCache}
}

func (sb *SubscriptionBatchBuilder) BuildBatch(
	srcToDst map[common.Address]common.Address,
	msgs []logpoller.Log,
	executed map[uint64]struct{},
	batchGasLimit uint64,
	gasPrice uint64,
	tokensPerFeeCoin map[common.Address]*big.Int,
	inflight []InflightExecutionReport,
) []uint64 {
	subTokenPerFeeCoin := tokensPerFeeCoin[sb.subFeeToken]
	if subTokenPerFeeCoin == nil {
		sb.lggr.Errorf("Fee token price not found for token: %s", sb.subFeeToken.Hex())
	}
	inflightSeqNrs, reserved, nonces, err := sb.inflight(gasPrice, subTokenPerFeeCoin, inflight, len(msgs))
	if err != nil {
		sb.lggr.Errorw("Unexpected error computing inflight values", "err", err)
		return []uint64{}
	}
	stalledSub := make(map[common.Address]struct{})
	subscriptionBalances := make(map[common.Address]*big.Int)
	var executableSeqNrs []uint64
	for _, msg := range msgs {
		subMsg, err := sb.parseLog(types.Log{
			Topics: msg.GetTopics(),
			Data:   msg.Data,
		})
		if err != nil {
			sb.lggr.Errorw("unable to parse message", "err", err, "msg", msg)
			continue
		}
		// Skip inflight
		if _, inflight := inflightSeqNrs[subMsg.Message.SequenceNumber]; inflight {
			continue
		}
		// Skip executed
		if _, executed := executed[subMsg.Message.SequenceNumber]; executed {
			continue
		}
		// Skip if sub is stalled
		if _, stalled := stalledSub[subMsg.Message.Receiver]; stalled {
			continue
		}
		strict, err := sb.subCache.IsStrict(subMsg.Message.Receiver)
		if err != nil {
			// Skip if we're unable to determine its strictness.
			continue
		}
		if strict {
			if sb.subCache.MostRecentExecution(subMsg.Message.Receiver) == MessageStateFailure {
				stalledSub[subMsg.Message.Receiver] = struct{}{}
				continue
			}
		}
		maxOverhead := maxGasOverHeadGasSubscription(len(msgs), subMsg)
		totalGasLimit := maxOverhead + subMsg.Message.GasLimit.Uint64()
		// Skip if insufficient gas left in the batch
		if batchGasLimit < totalGasLimit {
			sb.lggr.Infow("Insufficient remaining gas in batch limit", "batchGasLimit", batchGasLimit, "totalGasLimit", totalGasLimit)
			continue
		}
		maxCharge := maxSubCharge(gasPrice, subTokenPerFeeCoin, totalGasLimit)
		if _, ok := subscriptionBalances[subMsg.Message.Receiver]; !ok {
			reservedBalance, ok := reserved[subMsg.Message.Receiver]
			if !ok {
				reservedBalance = big.NewInt(0)
			}
			subscriptionBalances[subMsg.Message.Receiver] = big.NewInt(0).Sub(sb.subCache.Balance(subMsg.Message.Receiver), reservedBalance)
		}
		// Skip if insufficient balance
		if subscriptionBalances[subMsg.Message.Receiver].Cmp(maxCharge) < 0 {
			sb.lggr.Infow("Insufficient sub balance to execute msg", "balance", subscriptionBalances[subMsg.Message.Receiver], "maxCharge", maxCharge)
			continue
		}
		if strict {
			if _, ok := nonces[subMsg.Message.Receiver]; !ok {
				nonces[subMsg.Message.Receiver] = sb.subCache.Nonce(subMsg.Message.Receiver)
			}
			if subMsg.Message.Nonce != nonces[subMsg.Message.Receiver]+1 {
				sb.lggr.Infow("Invalid nonce", "nonce", subMsg.Message.Nonce, "expectedNonce", nonces[subMsg.Message.Receiver])
				continue
			}
			// We have the correct nonce, increment.
			nonces[subMsg.Message.Receiver] = subMsg.Message.Nonce
		}
		subscriptionBalances[subMsg.Message.Receiver] = big.NewInt(0).Sub(subscriptionBalances[subMsg.Message.Receiver], maxCharge)
		batchGasLimit -= totalGasLimit
		sb.lggr.Infow("Adding sub msg to batch", "seqNum", subMsg.Message.SequenceNumber, "maxCharge", maxCharge, "maxGasOverhead", maxOverhead, "strict", strict)
		executableSeqNrs = append(executableSeqNrs, subMsg.Message.SequenceNumber)
	}
	return executableSeqNrs
}

func (sb *SubscriptionBatchBuilder) parseLog(log types.Log) (*evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampCCIPSendRequested, error) {
	event := new(evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampCCIPSendRequested)
	err := bind.NewBoundContract(common.Address{}, sb.subABI, nil, nil, nil).UnpackLog(event, "CCIPSendRequested", log)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (sb *SubscriptionBatchBuilder) inflight(
	maxGasPrice uint64,
	subTokenPerFeeCoin *big.Int,
	inflight []InflightExecutionReport,
	numMsgsInRoot int,
) (map[uint64]struct{}, map[common.Address]*big.Int, map[common.Address]uint64, error) {
	reserved := make(map[common.Address]*big.Int)
	inflightSeqNrs := make(map[uint64]struct{})
	nonces := make(map[common.Address]uint64)
	// TODO: aggregate token value once we have
	// https://app.shortcut.com/chainlinklabs/story/41867/contract-scaffolding-aggregatetokenlimiter-contract
	for _, r := range inflight {
		for _, encMsg := range r.report.EncodedMessages {
			msg, err := sb.parseLog(types.Log{
				// Note this needs to change if we start indexing things.
				Topics: []common.Hash{CCIPSubSendRequested},
				Data:   encMsg,
			})
			if err != nil {
				return nil, nil, nil, err
			}
			totalGasLimit := maxGasOverHeadGasSubscription(numMsgsInRoot, msg) + msg.Message.GasLimit.Uint64()
			if reserved[msg.Message.Receiver] == nil {
				reserved[msg.Message.Receiver] = maxSubCharge(maxGasPrice, subTokenPerFeeCoin, totalGasLimit)
			} else {
				reserved[msg.Message.Receiver].Add(reserved[msg.Message.Receiver], maxSubCharge(maxGasPrice, subTokenPerFeeCoin, totalGasLimit))
			}
			if _, ok := nonces[msg.Message.Receiver]; !ok {
				nonces[msg.Message.Receiver] = msg.Message.Nonce
			}
			if msg.Message.Nonce > nonces[msg.Message.Receiver] {
				// Save max inflight nonce
				nonces[msg.Message.Receiver] = msg.Message.Nonce
			}
			inflightSeqNrs[msg.Message.SequenceNumber] = struct{}{}
		}
	}
	return inflightSeqNrs, reserved, nonces, nil
}
