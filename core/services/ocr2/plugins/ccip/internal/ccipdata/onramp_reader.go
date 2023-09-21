package ccipdata

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp_1_0_0"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp_1_1_0"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/hashlib"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

// EVM2EVMMessage is the interface for a message sent from the offramp to the onramp
// Plugin can operate against any lane version which has a message satisfying this interface.
type EVM2EVMMessage struct {
	SequenceNumber uint64
	GasLimit       *big.Int
	Nonce          uint64
	Hash           [32]byte
}

type OnRampReader interface {
	// GetSendRequestsGteSeqNum returns all the message send requests with sequence number greater than or equal to the provided.
	// If checkFinalityTags is set to true then confs param is ignored, the latest finalized block is used in the query.
	GetSendRequestsGteSeqNum(ctx context.Context, seqNum uint64, confs int) ([]Event[EVM2EVMMessage], error)

	// GetSendRequestsBetweenSeqNums returns all the message send requests in the provided sequence numbers range (inclusive).
	GetSendRequestsBetweenSeqNums(ctx context.Context, seqNumMin, seqNumMax uint64, confs int) ([]Event[EVM2EVMMessage], error)

	// Get router configured in the onRamp
	Router() common.Address
}

var _ OnRampReader = &OnRamp1_0_0{}

type OnRamp1_0_0 struct {
	address      common.Address
	onRamp       *evm_2_evm_onramp_1_0_0.EVM2EVMOnRamp
	finalityTags bool
	lp           logpoller.LogPoller
	lggr         logger.Logger
	client       client.Client
	leafHasher   hashlib.LeafHasherInterface[[32]byte]
}

func NewOnRamp1_0_0(lggr logger.Logger, sourceSelector, destSelector uint64, onRampAddress common.Address, sourceLP logpoller.LogPoller, source client.Client, finalityTags bool) *OnRamp1_0_0 {
	onRamp, err := evm_2_evm_onramp_1_0_0.NewEVM2EVMOnRamp(onRampAddress, source)
	if err != nil {
		panic(err) // ABI failure ok to panic
	}
	return &OnRamp1_0_0{
		lggr:         lggr,
		address:      onRampAddress,
		onRamp:       onRamp,
		lp:           sourceLP,
		finalityTags: finalityTags,
		leafHasher:   hashlib.NewLeafHasher(sourceSelector, destSelector, onRampAddress, hashlib.NewKeccakCtx()),
	}
}

func (o *OnRamp1_0_0) GetSendRequestsGteSeqNum(ctx context.Context, seqNum uint64, confs int) ([]Event[EVM2EVMMessage], error) {
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

func (o *OnRamp1_0_0) Router() common.Address {
	config, _ := o.onRamp.GetDynamicConfig(nil)
	return config.Router
}

func (o *OnRamp1_0_0) GetSendRequestsBetweenSeqNums(ctx context.Context, seqNumMin, seqNumMax uint64, confs int) ([]Event[EVM2EVMMessage], error) {
	//TODO implement me
	panic("implement me")
}

var _ OnRampReader = &OnRampV1_1_0{}

// OnRampV1_1_0 The only difference that the plugins care about in 1.1 is that the dynamic config struct has changed.
type OnRampV1_1_0 struct {
	*OnRamp1_0_0
	onRamp *evm_2_evm_onramp_1_1_0.EVM2EVMOnRamp
}

func NewOnRamp1_1_0(lggr logger.Logger, sourceSelector, destSelector uint64, onRampAddress common.Address, sourceLP logpoller.LogPoller, source client.Client, finalityTags bool) *OnRampV1_1_0 {
	onRamp, err := evm_2_evm_onramp_1_1_0.NewEVM2EVMOnRamp(onRampAddress, source)
	if err != nil {
		panic(err) // ABI failure ok to panic
	}
	return &OnRampV1_1_0{
		OnRamp1_0_0: NewOnRamp1_0_0(lggr, sourceSelector, destSelector, onRampAddress, sourceLP, source, finalityTags),
		onRamp:      onRamp,
	}
}

func (o *OnRampV1_1_0) Router() common.Address {
	config, _ := o.onRamp.GetDynamicConfig(nil)
	return config.Router
}

// NewOnRampReader determines the appropriate version of the onramp and returns a reader for it
func NewOnRampReader(lggr logger.Logger, sourceSelector, destSelector uint64, onRampAddress common.Address, sourceLP logpoller.LogPoller, source client.Client, finalityTags bool) (OnRampReader, error) {
	contractType, version, err := ccipconfig.TypeAndVersion(onRampAddress, source)
	if err != nil {
		return nil, errors.Errorf("expected %v got %v", ccipconfig.EVM2EVMOnRamp, contractType)
	}
	switch version.String() {
	case "1.0.0":
		return NewOnRamp1_0_0(lggr, sourceSelector, destSelector, onRampAddress, sourceLP, source, finalityTags), nil
	case "1.1.0":
		return NewOnRamp1_1_0(lggr, sourceSelector, destSelector, onRampAddress, sourceLP, source, finalityTags), nil
	default:
		return nil, errors.Errorf("expected version 1.0.0 got %v", version.String())
	}
}
