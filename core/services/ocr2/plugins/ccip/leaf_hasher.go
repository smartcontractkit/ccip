package ccip

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_subscription_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/core/utils"
)

type LeafHasher[H hasher.Hash] interface {
	HashLeaf(log types.Log) (H, error)
}

var (
	LeafDomainSeparator = [1]byte{0x00}
)

type TollLeafHasher struct {
	tollABI      abi.ABI
	metaDataHash [32]byte
	ctx          hasher.Ctx[[32]byte]
}

var _ LeafHasher[[32]byte] = &TollLeafHasher{}

func NewTollLeafHasher(sourceChainId *big.Int, destChainId *big.Int, onRampId common.Address, ctx hasher.Ctx[[32]byte]) *TollLeafHasher {
	tollABI, _ := abi.JSON(strings.NewReader(evm_2_evm_toll_onramp.EVM2EVMTollOnRampABI))
	return &TollLeafHasher{
		tollABI:      tollABI,
		metaDataHash: getMetaDataHash(ctx, ctx.Hash([]byte("EVM2EVMTollMessagePlus")), sourceChainId, onRampId, destChainId),
		ctx:          ctx,
	}
}

func (t *TollLeafHasher) HashLeaf(log types.Log) ([32]byte, error) {
	event, err := t.ParseEVM2EVMTollLog(log)
	if err != nil {
		return [32]byte{}, err
	}
	encodedTokens, err := utils.ABIEncode(`[{"components": [{"name": "token","type": "address"}, {"name": "amount", "type": "uint256"}],"type": "tuple[]"}]`, event.Message.TokensAndAmounts)
	if err != nil {
		return [32]byte{}, err
	}

	packedValues, err := utils.ABIEncode(
		`[{"type":"bytes1"},{"type":"bytes32"},{"type":"uint64"},{"type":"address"},{"type":"address"},{"type":"bytes32"},{"type":"bytes32"},{"type":"uint256"},{"components": [{"name": "token","type": "address"}, {"name": "amount", "type": "uint256"}],"name": "feeToken","type": "tuple"}]`,
		LeafDomainSeparator,
		t.metaDataHash,
		event.Message.SequenceNumber,
		event.Message.Sender,
		event.Message.Receiver,
		t.ctx.Hash(event.Message.Data),
		t.ctx.Hash(encodedTokens),
		event.Message.GasLimit,
		event.Message.FeeTokenAndAmount,
	)
	if err != nil {
		return [32]byte{}, err
	}
	return t.ctx.Hash(packedValues), nil
}

func (t *TollLeafHasher) ParseEVM2EVMTollLog(log types.Log) (*evm_2_evm_toll_onramp.EVM2EVMTollOnRampCCIPSendRequested, error) {
	event := new(evm_2_evm_toll_onramp.EVM2EVMTollOnRampCCIPSendRequested)
	err := bind.NewBoundContract(common.Address{}, t.tollABI, nil, nil, nil).UnpackLog(event, "CCIPSendRequested", log)
	return event, err
}

type SubscriptionLeafHasher struct {
	subABI       abi.ABI
	metaDataHash [32]byte
	ctx          hasher.Ctx[[32]byte]
}

var _ LeafHasher[[32]byte] = &SubscriptionLeafHasher{}

func NewSubscriptionLeafHasher(sourceChainId *big.Int, destChainId *big.Int, onRampId common.Address, ctx hasher.Ctx[[32]byte]) *SubscriptionLeafHasher {
	subABI, _ := abi.JSON(strings.NewReader(evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampABI))
	return &SubscriptionLeafHasher{
		subABI:       subABI,
		metaDataHash: getMetaDataHash(ctx, ctx.Hash([]byte("EVM2EVMSubscriptionMessagePlus")), sourceChainId, onRampId, destChainId),
		ctx:          ctx,
	}
}

func (s *SubscriptionLeafHasher) HashLeaf(log types.Log) ([32]byte, error) {
	event, err := s.ParseEVM2EVMSubscriptionLog(log)
	if err != nil {
		return [32]byte{}, err
	}
	encodedTokens, err := utils.ABIEncode(`[{"components": [{"name": "token","type": "address"}, {"name": "amount", "type": "uint256"}], "type": "tuple[]"}]`, event.Message.TokensAndAmounts)
	if err != nil {
		return [32]byte{}, err
	}
	packedValues, err := utils.ABIEncode(
		`[{"type":"bytes1"},{"type":"bytes32"},{"type":"uint64"},{"type":"address"},{"type":"address"},{"type":"bytes32"},{"type":"bytes32"},{"type":"uint256"},{"type":"uint64"}]`,
		LeafDomainSeparator,
		s.metaDataHash,
		event.Message.SequenceNumber,
		event.Message.Sender,
		event.Message.Receiver,
		s.ctx.Hash(event.Message.Data),
		s.ctx.Hash(encodedTokens),
		event.Message.GasLimit,
		event.Message.Nonce,
	)
	if err != nil {
		return [32]byte{}, err
	}
	return s.ctx.Hash(packedValues), nil
}

func (s *SubscriptionLeafHasher) ParseEVM2EVMSubscriptionLog(log types.Log) (*evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampCCIPSendRequested, error) {
	event := new(evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampCCIPSendRequested)
	err := bind.NewBoundContract(common.Address{}, s.subABI, nil, nil, nil).UnpackLog(event, "CCIPSendRequested", log)
	return event, err
}

func getMetaDataHash[H hasher.Hash](ctx hasher.Ctx[H], prefix [32]byte, sourceChainId *big.Int, onRampId common.Address, destChainId *big.Int) H {
	paddedOnRamp := onRampId.Hash()
	return ctx.Hash(utils.ConcatBytes(prefix[:], math.U256Bytes(sourceChainId), math.U256Bytes(destChainId), paddedOnRamp[:]))
}

func LogPollerLogToEthLog(log logpoller.Log) types.Log {
	return types.Log{
		Topics: log.GetTopics(),
		Data:   log.Data,
	}
}
