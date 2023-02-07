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
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/core/utils"
)

type LeafHasherInterface[H hasher.Hash] interface {
	HashLeaf(log types.Log) (H, error)
}

var (
	LeafDomainSeparator = [1]byte{0x00}
)

func getMetaDataHash[H hasher.Hash](ctx hasher.Ctx[H], prefix [32]byte, sourceChainId uint64, onRampId common.Address, destChainId uint64) H {
	paddedOnRamp := onRampId.Hash()
	return ctx.Hash(utils.ConcatBytes(prefix[:], math.U256Bytes(big.NewInt(0).SetUint64(sourceChainId)), math.U256Bytes(big.NewInt(0).SetUint64(destChainId)), paddedOnRamp[:]))
}

func LogPollerLogToEthLog(log logpoller.Log) types.Log {
	return types.Log{
		Topics: log.GetTopics(),
		Data:   log.Data,
	}
}

type LeafHasher struct {
	geABI        abi.ABI
	metaDataHash [32]byte
	ctx          hasher.Ctx[[32]byte]
}

func NewLeafHasher(sourceChainId uint64, destChainId uint64, onRampId common.Address, ctx hasher.Ctx[[32]byte]) *LeafHasher {
	geABI, _ := abi.JSON(strings.NewReader(evm_2_evm_onramp.EVM2EVMOnRampABI))
	return &LeafHasher{
		geABI:        geABI,
		metaDataHash: getMetaDataHash(ctx, ctx.Hash([]byte("EVM2EVMMessageEvent")), sourceChainId, onRampId, destChainId),
		ctx:          ctx,
	}
}

var _ LeafHasherInterface[[32]byte] = &LeafHasher{}

func (t *LeafHasher) HashLeaf(log types.Log) ([32]byte, error) {
	event, err := t.ParseEVM2EVMLog(log)
	if err != nil {
		return [32]byte{}, err
	}
	encodedTokens, err := utils.ABIEncode(`[{"components": [{"name": "token","type": "address"}, {"name": "amount", "type": "uint256"}],"type": "tuple[]"}]`, event.Message.TokensAndAmounts)
	if err != nil {
		return [32]byte{}, err
	}

	packedValues, err := utils.ABIEncode(
		`[
{"name": "leafDomainSeparator","type":"bytes1"},
{"name": "metadataHash", "type":"bytes32"},
{"name": "sequenceNumber", "type":"uint64"},
{"name": "nonce", "type":"uint64"},
{"name": "sender", "type":"address"},
{"name": "receiver", "type":"address"},
{"name": "dataHash", "type":"bytes32"},
{"name": "tokenAmountsHash", "type":"bytes32"},
{"name": "gasLimit", "type":"uint256"},
{"name": "strict", "type":"bool"},
{"name": "feeToken","type": "address"},
{"name": "feeTokenAmount","type": "uint256"}
]`,
		LeafDomainSeparator,
		t.metaDataHash,
		event.Message.SequenceNumber,
		event.Message.Nonce,
		event.Message.Sender,
		event.Message.Receiver,
		t.ctx.Hash(event.Message.Data),
		t.ctx.Hash(encodedTokens),
		event.Message.GasLimit,
		event.Message.Strict,
		event.Message.FeeToken,
		event.Message.FeeTokenAmount,
	)
	if err != nil {
		return [32]byte{}, err
	}
	return t.ctx.Hash(packedValues), nil
}

func (t *LeafHasher) ParseEVM2EVMLog(log types.Log) (*evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested, error) {
	event := new(evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested)
	err := bind.NewBoundContract(common.Address{}, t.geABI, nil, nil, nil).UnpackLog(event, "CCIPSendRequested", log)
	return event, err
}
