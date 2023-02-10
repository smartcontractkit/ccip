package ccip

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/fee_manager"
	"github.com/smartcontractkit/chainlink/core/utils"
)

var (
	// merkleRoot || minSeqNum || maxSeqNum
	ReportAccepted common.Hash
	// FeeManager
	GasFeeUpdated common.Hash
)

// MessageExecutionState defines the execution states of CCIP messages.
type MessageExecutionState uint64

const (
	Untouched MessageExecutionState = iota
	InProgress
	Success
	Failure
)

func getIDOrPanic(name string, abi2 abi.ABI) common.Hash {
	event, ok := abi2.Events[name]
	if !ok {
		panic(fmt.Sprintf("missing event %s", name))
	}
	return event.ID
}

func init() {
	commitStoreABI, err := abi.JSON(strings.NewReader(commit_store.CommitStoreABI))
	if err != nil {
		panic(err)
	}
	ReportAccepted = getIDOrPanic("ReportAccepted", commitStoreABI)

	feeManagerABI, err := abi.JSON(strings.NewReader(fee_manager.FeeManagerABI))
	if err != nil {
		panic(err)
	}
	GasFeeUpdated = getIDOrPanic("GasFeeUpdated", feeManagerABI)
}

func GetEventSignatures() EventSignatures {
	onRampABI, err := abi.JSON(strings.NewReader(evm_2_evm_onramp.EVM2EVMOnRampABI))
	if err != nil {
		panic(err)
	}
	CCIPSendRequested := getIDOrPanic("CCIPSendRequested", onRampABI)

	offRampABI, err := abi.JSON(strings.NewReader(evm_2_evm_offramp.EVM2EVMOffRampABI))
	if err != nil {
		panic(err)
	}
	ExecutionStateChanged := getIDOrPanic("ExecutionStateChanged", offRampABI)

	return EventSignatures{
		// offset || sourceChainID || seqNum || ...
		SendRequested:                    CCIPSendRequested,
		SendRequestedSequenceNumberIndex: 2,
		// sig || seqNum || messageId || ...
		ExecutionStateChanged:                    ExecutionStateChanged,
		ExecutionStateChangedSequenceNumberIndex: 1,
	}
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
		SourceChainId    uint64         `json:"sourceChainId"`
		SequenceNumber   uint64         `json:"sequenceNumber"`
		FeeTokenAmount   *big.Int       `json:"feeTokenAmount"`
		Sender           common.Address `json:"sender"`
		Nonce            uint64         `json:"nonce"`
		GasLimit         *big.Int       `json:"gasLimit"`
		Strict           bool           `json:"strict"`
		Receiver         common.Address `json:"receiver"`
		Data             []uint8        `json:"data"`
		TokensAndAmounts []struct {
			Token  common.Address `json:"token"`
			Amount *big.Int       `json:"amount"`
		} `json:"tokensAndAmounts"`
		FeeToken  common.Address `json:"feeToken"`
		MessageId [32]byte       `json:"messageId"`
	})
	if !ok {
		return nil, fmt.Errorf("invalid format have %T want %T", unpacked[0], receivedCp)
	}
	var tokensAndAmounts []evm_2_evm_onramp.CommonEVMTokenAndAmount
	for _, tokenAndAmount := range receivedCp.TokensAndAmounts {
		tokensAndAmounts = append(tokensAndAmounts, evm_2_evm_onramp.CommonEVMTokenAndAmount{
			Token:  tokenAndAmount.Token,
			Amount: tokenAndAmount.Amount,
		})
	}

	return &evm_2_evm_onramp.InternalEVM2EVMMessage{
		SourceChainId:    receivedCp.SourceChainId,
		SequenceNumber:   receivedCp.SequenceNumber,
		FeeTokenAmount:   receivedCp.FeeTokenAmount,
		Sender:           receivedCp.Sender,
		Nonce:            receivedCp.Nonce,
		GasLimit:         receivedCp.GasLimit,
		Strict:           receivedCp.Strict,
		Receiver:         receivedCp.Receiver,
		Data:             receivedCp.Data,
		TokensAndAmounts: tokensAndAmounts,
		FeeToken:         receivedCp.FeeToken,
		MessageId:        receivedCp.MessageId,
	}, nil
}

func MakeMessageArgs() abi.Arguments {
	var tuples = []abi.ArgumentMarshaling{
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
			Name: "tokensAndAmounts",
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
					Name: "feeUpdates",
					Type: "tuple[]",
					Components: []abi.ArgumentMarshaling{
						{
							Name: "sourceFeeToken",
							Type: "address",
						},
						{
							Name: "destChainId",
							Type: "uint64",
						},
						{
							Name: "feeTokenBaseUnitsPerUnitGas",
							Type: "uint256",
						},
					},
				},
				{
					Name: "encodedMessages",
					Type: "bytes[]",
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
