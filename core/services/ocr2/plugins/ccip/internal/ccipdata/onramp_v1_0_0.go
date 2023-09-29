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
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp_1_0_0"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/hashlib"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

const (
	CCIPSendRequestedEventNameV1_0_0 = "CCIPSendRequested"
	MetaDataHashPrefixV1_0_0         = "EVM2EVMMessageEvent"
)

var leafDomainSeparator = [1]byte{0x00}

type LeafHasherV1_0_0 struct {
	metaDataHash [32]byte
	ctx          hashlib.Ctx[[32]byte]
	onRamp       *evm_2_evm_onramp_1_0_0.EVM2EVMOnRamp
}

func getMetaDataHash[H hashlib.Hash](ctx hashlib.Ctx[H], prefix [32]byte, sourceChainSelector uint64, onRampId common.Address, destChainSelector uint64) H {
	paddedOnRamp := onRampId.Hash()
	return ctx.Hash(utils.ConcatBytes(prefix[:], math.U256Bytes(big.NewInt(0).SetUint64(sourceChainSelector)), math.U256Bytes(big.NewInt(0).SetUint64(destChainSelector)), paddedOnRamp[:]))
}

func NewLeafHasherV1_0_0(sourceChainSelector uint64, destChainSelector uint64, onRampId common.Address, ctx hashlib.Ctx[[32]byte], onRamp *evm_2_evm_onramp_1_0_0.EVM2EVMOnRamp) *LeafHasherV1_0_0 {
	return &LeafHasherV1_0_0{
		metaDataHash: getMetaDataHash(ctx, ctx.Hash([]byte(MetaDataHashPrefixV1_0_0)), sourceChainSelector, onRampId, destChainSelector),
		ctx:          ctx,
		onRamp:       onRamp,
	}
}

func (t *LeafHasherV1_0_0) HashLeaf(log types.Log) ([32]byte, error) {
	message, err := t.onRamp.ParseCCIPSendRequested(log)
	if err != nil {
		return [32]byte{}, err
	}
	encodedTokens, err := utils.ABIEncode(
		`[
{"components": [{"name":"token","type":"address"},{"name":"amount","type":"uint256"}], "type":"tuple[]"}]`, message.Message.TokenAmounts)
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
		leafDomainSeparator,
		t.metaDataHash,
		message.Message.SequenceNumber,
		message.Message.Nonce,
		message.Message.Sender,
		message.Message.Receiver,
		t.ctx.Hash(message.Message.Data),
		t.ctx.Hash(encodedTokens),
		message.Message.GasLimit,
		message.Message.Strict,
		message.Message.FeeToken,
		message.Message.FeeTokenAmount,
	)
	if err != nil {
		return [32]byte{}, err
	}
	return t.ctx.Hash(packedValues), nil
}

var _ OnRampReader = &OnRampV1_0_0{}

type OnRampV1_0_0 struct {
	address                    common.Address
	onRamp                     *evm_2_evm_onramp_1_0_0.EVM2EVMOnRamp
	finalityTags               bool
	lp                         logpoller.LogPoller
	lggr                       logger.Logger
	client                     client.Client
	leafHasher                 LeafHasherInterface[[32]byte]
	filterName                 string
	sendRequestedEventSig      common.Hash
	sendRequestedSeqNumberWord int
}

func (o *OnRampV1_0_0) GetLastUSDCMessagePriorToLogIndexInTx(ctx context.Context, logIndex int64, txHash common.Hash) ([]byte, error) {
	return nil, errors.New("USDC not supported in < 1.2.0")
}

func NewOnRampV1_0_0(lggr logger.Logger, sourceSelector, destSelector uint64, onRampAddress common.Address, sourceLP logpoller.LogPoller, source client.Client, finalityTags bool) (*OnRampV1_0_0, error) {
	onRamp, err := evm_2_evm_onramp_1_0_0.NewEVM2EVMOnRamp(onRampAddress, source)
	if err != nil {
		return nil, err
	}
	onRampABI, err := abi.JSON(strings.NewReader(evm_2_evm_onramp_1_0_0.EVM2EVMOnRampABI))
	if err != nil {
		return nil, err
	}
	// Subscribe to the relevant logs
	name := logpoller.FilterName(COMMIT_CCIP_SENDS, onRampAddress)
	eventSig := abihelpers.GetIDOrPanic(CCIPSendRequestedEventNameV1_0_0, onRampABI)
	err = sourceLP.RegisterFilter(logpoller.Filter{
		Name:      name,
		EventSigs: []common.Hash{eventSig},
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
		leafHasher:   NewLeafHasherV1_0_0(sourceSelector, destSelector, onRampAddress, hashlib.NewKeccakCtx(), onRamp),
		filterName:   name,
		// offset || sourceChainID || seqNum || ...
		sendRequestedSeqNumberWord: 2,
		sendRequestedEventSig:      eventSig,
	}, nil
}

func (o *OnRampV1_0_0) logToMessage(log types.Log) (*EVM2EVMMessage, error) {
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
}

func (o *OnRampV1_0_0) GetSendRequestsGteSeqNum(ctx context.Context, seqNum uint64, confs int) ([]Event[EVM2EVMMessage], error) {
	if !o.finalityTags {
		logs, err2 := o.lp.LogsDataWordGreaterThan(
			o.sendRequestedEventSig,
			o.address,
			o.sendRequestedSeqNumberWord,
			abihelpers.EvmWord(seqNum),
			confs,
			pg.WithParentCtx(ctx),
		)
		if err2 != nil {
			return nil, fmt.Errorf("logs data word greater than: %w", err2)
		}
		return parseLogs[EVM2EVMMessage](logs, o.lggr, o.logToMessage)
	}
	latestFinalizedHash, err := latestFinalizedBlockHash(ctx, o.client)
	if err != nil {
		return nil, err
	}
	logs, err := o.lp.LogsUntilBlockHashDataWordGreaterThan(
		o.sendRequestedEventSig,
		o.address,
		o.sendRequestedSeqNumberWord,
		abihelpers.EvmWord(seqNum),
		latestFinalizedHash,
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("logs until block hash data word greater than: %w", err)
	}
	return parseLogs[EVM2EVMMessage](logs, o.lggr, o.logToMessage)
}

func (o *OnRampV1_0_0) RouterAddress() common.Address {
	config, _ := o.onRamp.GetDynamicConfig(nil)
	return config.Router
}

func (o *OnRampV1_0_0) GetSendRequestsBetweenSeqNums(ctx context.Context, seqNumMin, seqNumMax uint64, confs int) ([]Event[EVM2EVMMessage], error) {
	logs, err := o.lp.LogsDataWordRange(
		o.sendRequestedEventSig,
		o.address,
		o.sendRequestedSeqNumberWord,
		logpoller.EvmWord(seqNumMin),
		logpoller.EvmWord(seqNumMax),
		confs,
		pg.WithParentCtx(ctx))
	if err != nil {
		return nil, err
	}
	return parseLogs[EVM2EVMMessage](logs, o.lggr, o.logToMessage)
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
