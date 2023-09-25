package ccipdata

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp_1_0_0"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/hashlib"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

var leafDomainSeparator = [1]byte{0x00}

type LeafHasherV1_0_0 struct {
	metaDataHash [32]byte
	ctx          hashlib.Ctx[[32]byte]
}

func getMetaDataHash[H hashlib.Hash](ctx hashlib.Ctx[H], prefix [32]byte, sourceChainSelector uint64, onRampId common.Address, destChainSelector uint64) H {
	paddedOnRamp := onRampId.Hash()
	return ctx.Hash(utils.ConcatBytes(prefix[:], math.U256Bytes(big.NewInt(0).SetUint64(sourceChainSelector)), math.U256Bytes(big.NewInt(0).SetUint64(destChainSelector)), paddedOnRamp[:]))
}

func NewLeafHasherV1_0_0(sourceChainSelector uint64, destChainSelector uint64, onRampId common.Address, ctx hashlib.Ctx[[32]byte]) *LeafHasherV1_0_0 {
	return &LeafHasherV1_0_0{
		metaDataHash: getMetaDataHash(ctx, ctx.Hash([]byte("EVM2EVMMessageEvent")), sourceChainSelector, onRampId, destChainSelector),
		ctx:          ctx,
	}
}

func (t *LeafHasherV1_0_0) HashLeaf(log types.Log) ([32]byte, error) {
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

var _ OnRampReader = &OnRampV1_0_0{}

type OnRampV1_0_0 struct {
	address      common.Address
	onRamp       *evm_2_evm_onramp_1_0_0.EVM2EVMOnRamp
	finalityTags bool
	lp           logpoller.LogPoller
	lggr         logger.Logger
	client       client.Client
	leafHasher   LeafHasherInterface[[32]byte]
	filterName   string
}

func NewOnRampV1_0_0(lggr logger.Logger, sourceSelector, destSelector uint64, onRampAddress common.Address, sourceLP logpoller.LogPoller, source client.Client, finalityTags bool) (*OnRampV1_0_0, error) {
	onRamp, err := evm_2_evm_onramp_1_0_0.NewEVM2EVMOnRamp(onRampAddress, source)
	if err != nil {
		return nil, err
	}
	onRampABI, err := abi.JSON(strings.NewReader(evm_2_evm_onramp.EVM2EVMOnRampABI))
	if err != nil {
		return nil, err
	}
	// Subscribe to the relevant logs
	name := logpoller.FilterName(COMMIT_CCIP_SENDS, onRampAddress)
	err = sourceLP.RegisterFilter(logpoller.Filter{
		Name:      name,
		EventSigs: []common.Hash{abihelpers.GetIDOrPanic("CCIPSendRequested", onRampABI)},
		Addresses: []common.Address{onRampAddress},
	})
	if err != nil {
		return nil, err
	}
	return &OnRampV1_0_0{
		lggr:         lggr,
		address:      onRampAddress,
		onRamp:       onRamp,
		lp:           sourceLP,
		finalityTags: finalityTags,
		leafHasher:   NewLeafHasherV1_0_0(sourceSelector, destSelector, onRampAddress, hashlib.NewKeccakCtx()),
		filterName:   name,
	}, nil
}

func (o *OnRampV1_0_0) GetSendRequestsGteSeqNum(ctx context.Context, seqNum uint64, confs int) ([]Event[EVM2EVMMessage], error) {
	if !o.finalityTags {
		logs, err2 := o.lp.LogsDataWordGreaterThan(
			abihelpers.EventSignatures.SendRequested,
			o.address,
			abihelpers.EventSignatures.SendRequestedSequenceNumberWord,
			abihelpers.EvmWord(seqNum),
			confs,
			pg.WithParentCtx(ctx),
		)
		if err2 != nil {
			return nil, fmt.Errorf("logs data word greater than: %w", err2)
		}
		return parseLogs[EVM2EVMMessage](
			logs,
			o.lggr,
			func(log types.Log) (*EVM2EVMMessage, error) {
				msg, err := o.onRamp.ParseCCIPSendRequested(log)
				if err != nil {
					return nil, err
				}
				h, err := o.leafHasher.HashLeaf(log)
				if err != nil {
					return nil, err
				}
				return &EVM2EVMMessage{
					SequenceNumber: msg.Message.SequenceNumber,
					GasLimit:       msg.Message.GasLimit,
					Nonce:          msg.Message.Nonce,
					Hash:           h,
					Log:            log,
				}, nil
			},
		)
	}

	// If the chain is based on explicit finality we only examine logs less than or equal to the latest finalized block number.
	// NOTE: there appears to be a bug in ethclient whereby BlockByNumber fails with "unsupported txtype" when trying to parse the block
	// when querying L2s, headers however work.
	// TODO (CCIP-778): Migrate to core finalized tags, below doesn't work for some chains e.g. Celo.
	latestFinalizedHeader, err := o.client.HeaderByNumber(
		ctx,
		big.NewInt(rpc.FinalizedBlockNumber.Int64()),
	)
	if err != nil {
		return nil, err
	}

	if latestFinalizedHeader == nil {
		return nil, errors.New("latest finalized header is nil")
	}
	if latestFinalizedHeader.Number == nil {
		return nil, errors.New("latest finalized number is nil")
	}
	logs, err := o.lp.LogsUntilBlockHashDataWordGreaterThan(
		abihelpers.EventSignatures.SendRequested,
		o.address,
		abihelpers.EventSignatures.SendRequestedSequenceNumberWord,
		abihelpers.EvmWord(seqNum),
		latestFinalizedHeader.Hash(),
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("logs until block hash data word greater than: %w", err)
	}

	return parseLogs[EVM2EVMMessage](
		logs,
		o.lggr,
		func(log types.Log) (*EVM2EVMMessage, error) {
			msg, err := o.onRamp.ParseCCIPSendRequested(log)
			if err != nil {
				return nil, err
			}
			h, err := o.leafHasher.HashLeaf(log)
			if err != nil {
				return nil, err
			}
			return &EVM2EVMMessage{
				SequenceNumber: msg.Message.SequenceNumber,
				GasLimit:       msg.Message.GasLimit,
				Nonce:          msg.Message.Nonce,
				Hash:           h,
			}, nil
		},
	)
}

func (o *OnRampV1_0_0) Router() common.Address {
	config, _ := o.onRamp.GetDynamicConfig(nil)
	return config.Router
}

func (o *OnRampV1_0_0) GetSendRequestsBetweenSeqNums(ctx context.Context, seqNumMin, seqNumMax uint64, confs int) ([]Event[EVM2EVMMessage], error) {
	logs, err := o.lp.LogsDataWordRange(
		abihelpers.EventSignatures.SendRequested,
		o.address,
		abihelpers.EventSignatures.SendRequestedSequenceNumberWord,
		logpoller.EvmWord(seqNumMin),
		logpoller.EvmWord(seqNumMax),
		confs,
		pg.WithParentCtx(ctx))
	if err != nil {
		return nil, err
	}

	return parseLogs[EVM2EVMMessage](
		logs,
		o.lggr,
		func(log types.Log) (*EVM2EVMMessage, error) {
			msg, err := o.onRamp.ParseCCIPSendRequested(log)
			if err != nil {
				return nil, err
			}
			h, err := o.leafHasher.HashLeaf(log)
			if err != nil {
				return nil, err
			}
			return &EVM2EVMMessage{
				SequenceNumber: msg.Message.SequenceNumber,
				GasLimit:       msg.Message.GasLimit,
				Nonce:          msg.Message.Nonce,
				Hash:           h,
			}, nil
		},
	)
}

// TODO: follow up with offramp version abstraction
func (o *OnRampV1_0_0) ToOffRampMessage(message EVM2EVMMessage) (*evm_2_evm_offramp.InternalEVM2EVMMessage, error) {
	m, err := o.onRamp.ParseCCIPSendRequested(message.Log)
	if err != nil {
		return nil, err
	}
	tokensAndAmounts := make([]evm_2_evm_offramp.ClientEVMTokenAmount, len(m.Message.TokenAmounts))
	for i, tokenAndAmount := range m.Message.TokenAmounts {
		tokensAndAmounts[i] = evm_2_evm_offramp.ClientEVMTokenAmount{
			Token:  tokenAndAmount.Token,
			Amount: tokenAndAmount.Amount,
		}
	}
	return &evm_2_evm_offramp.InternalEVM2EVMMessage{
		SourceChainSelector: m.Message.SourceChainSelector,
		Sender:              m.Message.Sender,
		Receiver:            m.Message.Receiver,
		SequenceNumber:      m.Message.SequenceNumber,
		GasLimit:            m.Message.GasLimit,
		Strict:              m.Message.Strict,
		Nonce:               m.Message.Nonce,
		FeeToken:            m.Message.FeeToken,
		FeeTokenAmount:      m.Message.FeeTokenAmount,
		Data:                m.Message.Data,
		TokenAmounts:        tokensAndAmounts,
		SourceTokenData:     make([][]byte, len(m.Message.TokenAmounts)),
		MessageId:           m.Message.MessageId,
	}, nil
}

func (o *OnRampV1_0_0) Close() error {
	return o.lp.UnregisterFilter(o.filterName)
}
