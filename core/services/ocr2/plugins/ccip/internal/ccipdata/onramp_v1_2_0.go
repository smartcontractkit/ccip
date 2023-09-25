package ccipdata

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/hashlib"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

type LeafHasherV1_2_0 struct {
	metaDataHash [32]byte
	ctx          hashlib.Ctx[[32]byte]
}

func NewLeafHasherV1_2_0(sourceChainSelector uint64, destChainSelector uint64, onRampId common.Address, ctx hashlib.Ctx[[32]byte]) *LeafHasherV1_2_0 {
	return &LeafHasherV1_2_0{
		metaDataHash: getMetaDataHash(ctx, ctx.Hash([]byte("EVM2EVMMessageHashV2")), sourceChainSelector, onRampId, destChainSelector),
		ctx:          ctx,
	}
}

func (t *LeafHasherV1_2_0) HashLeaf(log types.Log) ([32]byte, error) {
	message, err := abihelpers.DecodeOffRampMessage(log.Data)
	if err != nil {
		return [32]byte{}, err
	}

	encodedTokens, err := abihelpers.TokenAmountsArgs.PackValues([]interface{}{message.TokenAmounts})
	if err != nil {
		return [32]byte{}, err
	}

	bytesArray, err := abi.NewType("bytes[]", "bytes[]", nil)
	if err != nil {
		return [32]byte{}, err
	}

	encodedSourceTokenData, err := abi.Arguments{abi.Argument{Type: bytesArray}}.PackValues([]interface{}{message.SourceTokenData})
	if err != nil {
		return [32]byte{}, err
	}

	// TODO: Update according 1.2
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
{"name": "sourceTokenDataHash", "type":"bytes32"},
{"name": "gasLimit", "type":"uint256"},
{"name": "strict", "type":"bool"},
{"name": "feeToken","type": "address"},
{"name": "feeTokenAmount","type": "uint256"}
]`,
		leafDomainSeparator,
		t.metaDataHash,
		message.SequenceNumber,
		message.Nonce,
		message.Sender,
		message.Receiver,
		t.ctx.Hash(message.Data),
		t.ctx.Hash(encodedTokens),
		t.ctx.Hash(encodedSourceTokenData),
		message.GasLimit,
		message.Strict,
		message.FeeToken,
		message.FeeTokenAmount,
	)
	if err != nil {
		return [32]byte{}, err
	}
	return t.ctx.Hash(packedValues), nil
}

var _ OnRampReader = &OnRampV1_2_0{}

// Significant change in 1.2:
// - CCIPSendRequested event signature has changed
type OnRampV1_2_0 struct {
	onRamp     *evm_2_evm_onramp.EVM2EVMOnRamp
	leafHasher LeafHasherInterface[[32]byte]
}

func (o *OnRampV1_2_0) GetSendRequestsGteSeqNum(ctx context.Context, seqNum uint64, confs int) ([]Event[EVM2EVMMessage], error) {
	//TODO implement me
	panic("implement me")
}

func (o *OnRampV1_2_0) GetSendRequestsBetweenSeqNums(ctx context.Context, seqNumMin, seqNumMax uint64, confs int) ([]Event[EVM2EVMMessage], error) {
	//TODO implement me
	panic("implement me")
}

func (o *OnRampV1_2_0) Router() common.Address {
	//TODO implement me
	panic("implement me")
}

func (o *OnRampV1_2_0) ToOffRampMessage(message EVM2EVMMessage) (*evm_2_evm_offramp.InternalEVM2EVMMessage, error) {
	//TODO implement me
	panic("implement me")
}

func (o *OnRampV1_2_0) Close() error {
	//TODO implement me
	panic("implement me")
}

func NewOnRampV1_2_0(lggr logger.Logger, sourceSelector, destSelector uint64, onRampAddress common.Address, sourceLP logpoller.LogPoller, source client.Client, finalityTags bool) (*OnRampV1_2_0, error) {
	onRamp, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(onRampAddress, source)
	if err != nil {
		panic(err) // ABI failure ok to panic
	}
	return &OnRampV1_2_0{
		leafHasher: NewLeafHasherV1_2_0(sourceSelector, destSelector, onRampAddress, hashlib.NewKeccakCtx()),
		onRamp:     onRamp,
	}, nil
}
