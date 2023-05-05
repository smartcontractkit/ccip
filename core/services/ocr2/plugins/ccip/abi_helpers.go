package ccip

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

// MessageExecutionState defines the execution states of CCIP messages.
type MessageExecutionState uint8

const (
	ExecutionStateUntouched MessageExecutionState = iota
	ExecutionStateInProgress
	ExecutionStateSuccess
	ExecutionStateFailure
)

var EventSignatures struct {
	// OnRamp
	SendRequested common.Hash
	// CommitStore
	ReportAccepted common.Hash
	// OffRamp
	ExecutionStateChanged common.Hash
	PoolAdded             common.Hash
	PoolRemoved           common.Hash

	// PriceRegistry
	UsdPerUnitGasUpdated common.Hash
	UsdPerTokenUpdated   common.Hash

	// offset || sourceChainID || seqNum || ...
	SendRequestedSequenceNumberWord int
	// offset || priceUpdatesOffset || minSeqNum || maxSeqNum || merkleRoot
	ReportAcceptedMaxSequenceNumberWord int
	// sig || seqNum || messageId || ...
	ExecutionStateChangedSequenceNumberIndex int
}

func getIDOrPanic(name string, abi2 abi.ABI) common.Hash {
	event, ok := abi2.Events[name]
	if !ok {
		panic(fmt.Sprintf("missing event %s", name))
	}
	return event.ID
}

func init() {
	onRampABI, err := abi.JSON(strings.NewReader(evm_2_evm_onramp.EVM2EVMOnRampABI))
	if err != nil {
		panic(err)
	}
	EventSignatures.SendRequested = getIDOrPanic("CCIPSendRequested", onRampABI)
	EventSignatures.SendRequestedSequenceNumberWord = 2

	commitStoreABI, err := abi.JSON(strings.NewReader(commit_store.CommitStoreABI))
	if err != nil {
		panic(err)
	}
	EventSignatures.ReportAccepted = getIDOrPanic("ReportAccepted", commitStoreABI)
	EventSignatures.ReportAcceptedMaxSequenceNumberWord = 3

	offRampABI, err := abi.JSON(strings.NewReader(evm_2_evm_offramp.EVM2EVMOffRampABI))
	if err != nil {
		panic(err)
	}
	EventSignatures.ExecutionStateChanged = getIDOrPanic("ExecutionStateChanged", offRampABI)
	EventSignatures.ExecutionStateChangedSequenceNumberIndex = 1
	EventSignatures.PoolAdded = getIDOrPanic("PoolAdded", offRampABI)
	EventSignatures.PoolRemoved = getIDOrPanic("PoolRemoved", offRampABI)

	priceRegistryABI, err := abi.JSON(strings.NewReader(price_registry.PriceRegistryABI))
	if err != nil {
		panic(err)
	}
	EventSignatures.UsdPerUnitGasUpdated = getIDOrPanic("UsdPerUnitGasUpdated", priceRegistryABI)
	EventSignatures.UsdPerTokenUpdated = getIDOrPanic("UsdPerTokenUpdated", priceRegistryABI)
}

func DecodeMessage(b []byte) (*evm_2_evm_onramp.InternalEVM2EVMMessage, error) {
	unpacked, err := MakeMessageArgs().Unpack(b)
	if err != nil {
		return nil, err
	}
	if len(unpacked) == 0 {
		return nil, fmt.Errorf("no message found when unpacking")
	}

	// Note must use unnamed type here
	receivedCp, ok := unpacked[0].(struct {
		SourceChainId  uint64         `json:"sourceChainId"`
		SequenceNumber uint64         `json:"sequenceNumber"`
		FeeTokenAmount *big.Int       `json:"feeTokenAmount"`
		Sender         common.Address `json:"sender"`
		Nonce          uint64         `json:"nonce"`
		GasLimit       *big.Int       `json:"gasLimit"`
		Strict         bool           `json:"strict"`
		Receiver       common.Address `json:"receiver"`
		Data           []uint8        `json:"data"`
		TokenAmounts   []struct {
			Token  common.Address `json:"token"`
			Amount *big.Int       `json:"amount"`
		} `json:"tokenAmounts"`
		FeeToken  common.Address `json:"feeToken"`
		MessageId [32]byte       `json:"messageId"`
	})
	if !ok {
		return nil, fmt.Errorf("invalid format have %T want %T", unpacked[0], receivedCp)
	}
	var tokensAndAmounts []evm_2_evm_onramp.ClientEVMTokenAmount
	for _, tokenAndAmount := range receivedCp.TokenAmounts {
		tokensAndAmounts = append(tokensAndAmounts, evm_2_evm_onramp.ClientEVMTokenAmount{
			Token:  tokenAndAmount.Token,
			Amount: tokenAndAmount.Amount,
		})
	}

	return &evm_2_evm_onramp.InternalEVM2EVMMessage{
		SourceChainId:  receivedCp.SourceChainId,
		SequenceNumber: receivedCp.SequenceNumber,
		FeeTokenAmount: receivedCp.FeeTokenAmount,
		Sender:         receivedCp.Sender,
		Nonce:          receivedCp.Nonce,
		GasLimit:       receivedCp.GasLimit,
		Strict:         receivedCp.Strict,
		Receiver:       receivedCp.Receiver,
		Data:           receivedCp.Data,
		TokenAmounts:   tokensAndAmounts,
		FeeToken:       receivedCp.FeeToken,
		MessageId:      receivedCp.MessageId,
	}, nil
}

func MakeMessageArgs() abi.Arguments {
	tuples := []abi.ArgumentMarshaling{
		{
			Name: "sourceChainId",
			Type: "uint64",
		},
		{
			Name: "sequenceNumber",
			Type: "uint64",
		},
		{
			Name: "feeTokenAmount",
			Type: "uint256",
		},
		{
			Name: "sender",
			Type: "address",
		},
		{
			Name: "nonce",
			Type: "uint64",
		},
		{
			Name: "gasLimit",
			Type: "uint256",
		},
		{
			Name: "strict",
			Type: "bool",
		},
		{
			Name: "receiver",
			Type: "address",
		},
		{
			Name: "data",
			Type: "bytes",
		},
		{
			Name: "tokenAmounts",
			Type: "tuple[]",
			Components: []abi.ArgumentMarshaling{
				{
					Name: "token",
					Type: "address",
				},
				{
					Name: "amount",
					Type: "uint256",
				},
			},
		},
		{
			Name: "feeToken",
			Type: "address",
		},
		{
			Name: "messageId",
			Type: "bytes32",
		},
	}
	ty, _ := abi.NewType("tuple", "", tuples)
	return abi.Arguments{
		{
			Type: ty,
		},
	}
}

// ProofFlagsToBits transforms a list of boolean proof flags to a *big.Int
// encoded number.
func ProofFlagsToBits(proofFlags []bool) *big.Int {
	encodedFlags := big.NewInt(0)
	for i := 0; i < len(proofFlags); i++ {
		if proofFlags[i] {
			encodedFlags.SetBit(encodedFlags, i, 1)
		}
	}
	return encodedFlags
}

func makeExecutionReportArgs() abi.Arguments {
	return []abi.Argument{
		{
			Name: "ExecutionReport",
			Type: utils.MustAbiType("tuple", []abi.ArgumentMarshaling{
				{
					Name: "sequenceNumbers",
					Type: "uint64[]",
				},
				{
					Name: "encodedMessages",
					Type: "bytes[]",
				},
				{
					Name: "offchainTokenData",
					Type: "bytes[][]",
				},
				{
					Name: "proofs",
					Type: "bytes32[]",
				},
				{
					Name: "proofFlagBits",
					Type: "uint256",
				},
			}),
		},
	}
}

func makeCommitReportArgs() abi.Arguments {
	return []abi.Argument{
		{
			Name: "CommitReport",
			Type: utils.MustAbiType("tuple", []abi.ArgumentMarshaling{
				{
					Name: "priceUpdates",
					Type: "tuple",
					Components: []abi.ArgumentMarshaling{
						{
							Name: "tokenPriceUpdates",
							Type: "tuple[]",
							Components: []abi.ArgumentMarshaling{
								{
									Name: "sourceToken",
									Type: "address",
								},
								{
									Name: "usdPerToken",
									Type: "uint128",
								},
							},
						},
						{
							Name: "destChainId",
							Type: "uint64",
						},
						{
							Name: "usdPerUnitGas",
							Type: "uint128",
						},
					},
				},
				{
					Name: "interval",
					Type: "tuple",
					Components: []abi.ArgumentMarshaling{
						{
							Name: "min",
							Type: "uint64",
						},
						{
							Name: "max",
							Type: "uint64",
						},
					},
				},
				{
					Name: "merkleRoot",
					Type: "bytes32",
				},
			}),
		},
	}
}

type AbiDefined interface {
	AbiString() string
}

type AbiDefinedValid interface {
	AbiDefined
	Validate() error
}

func EncodeAbiStruct[T AbiDefined](decoded T) ([]byte, error) {
	encoded, err := utils.ABIEncode(decoded.AbiString(), decoded)
	if err != nil {
		return nil, err
	}
	return encoded, nil
}

func DecodeAbiStruct[T AbiDefinedValid](encoded []byte) (T, error) {
	var empty T

	decoded, err := utils.ABIDecode(empty.AbiString(), encoded)
	if err != nil {
		return empty, err
	}

	converted := abi.ConvertType(decoded[0], &empty)
	if casted, ok := converted.(*T); ok {
		return *casted, (*casted).Validate()
	}
	return empty, fmt.Errorf("can't cast from %T to %T", converted, empty)
}
