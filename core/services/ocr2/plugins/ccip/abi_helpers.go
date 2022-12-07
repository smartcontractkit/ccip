package ccip

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_ge_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_ge_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/utils"
)

var (
	// merkleRoot || minSeqNum || maxSeqNum
	ReportAccepted common.Hash
	ConfigSet      common.Hash
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
	ConfigSet = getIDOrPanic("ConfigSet", commitStoreABI)
}

func GetTollEventSignatures() EventSignatures {
	tollOnRampABI, err := abi.JSON(strings.NewReader(evm_2_evm_toll_onramp.EVM2EVMTollOnRampABI))
	if err != nil {
		panic(err)
	}
	CCIPSendRequested := getIDOrPanic("CCIPSendRequested", tollOnRampABI)

	tollOffRampABI, err := abi.JSON(strings.NewReader(evm_2_evm_toll_offramp.EVM2EVMTollOffRampABI))
	if err != nil {
		panic(err)
	}
	ExecutionStateChanged := getIDOrPanic("ExecutionStateChanged", tollOffRampABI)

	return EventSignatures{
		// offset || sourceChainID || seqNum || ...
		SendRequested:                    CCIPSendRequested,
		SendRequestedSequenceNumberIndex: 2,
		// sig || seqNum || ...
		ExecutionStateChanged:                    ExecutionStateChanged,
		ExecutionStateChangedSequenceNumberIndex: 1,
	}
}

func GetGEEventSignatures() EventSignatures {
	geOnRampABI, err := abi.JSON(strings.NewReader(evm_2_evm_ge_onramp.EVM2EVMGEOnRampABI))
	if err != nil {
		panic(err)
	}
	CCIPSendRequested := getIDOrPanic("CCIPSendRequested", geOnRampABI)

	offRampABI, err := abi.JSON(strings.NewReader(evm_2_evm_ge_offramp.EVM2EVMGEOffRampABI))
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

// DecodeCCIPMessage decodes the bytecode message into a commit_store.CCIPAny2EVMTollMessage
// This function returns an error if there is no message in the bytecode or
// when the payload is malformed.
func DecodeCCIPMessage(b []byte) (*evm_2_evm_toll_onramp.TollEVM2EVMTollMessage, error) {
	unpacked, err := MakeTollCCIPMsgArgs().Unpack(b)
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
		Sender           common.Address `json:"sender"`
		Receiver         common.Address `json:"receiver"`
		Data             []uint8        `json:"data"`
		TokensAndAmounts []struct {
			Token  common.Address `json:"token"`
			Amount *big.Int       `json:"amount"`
		} `json:"tokensAndAmounts"`
		FeeTokenAndAmount struct {
			Token  common.Address `json:"token"`
			Amount *big.Int       `json:"amount"`
		} `json:"feeTokenAndAmount"`
		GasLimit *big.Int `json:"gasLimit"`
	})
	if !ok {
		return nil, fmt.Errorf("invalid format have %T want %T", unpacked[0], receivedCp)
	}

	var tokensAndAmounts []evm_2_evm_toll_onramp.CommonEVMTokenAndAmount

	for _, tokenAndAmount := range receivedCp.TokensAndAmounts {
		tokensAndAmounts = append(tokensAndAmounts, evm_2_evm_toll_onramp.CommonEVMTokenAndAmount{
			Token:  tokenAndAmount.Token,
			Amount: tokenAndAmount.Amount,
		})
	}

	return &evm_2_evm_toll_onramp.TollEVM2EVMTollMessage{
		SourceChainId:    receivedCp.SourceChainId,
		SequenceNumber:   receivedCp.SequenceNumber,
		Sender:           receivedCp.Sender,
		Receiver:         receivedCp.Receiver,
		Data:             receivedCp.Data,
		TokensAndAmounts: tokensAndAmounts,
		FeeTokenAndAmount: evm_2_evm_toll_onramp.CommonEVMTokenAndAmount{
			Token:  receivedCp.FeeTokenAndAmount.Token,
			Amount: receivedCp.FeeTokenAndAmount.Amount,
		},
		GasLimit: receivedCp.GasLimit,
	}, nil
}

// MakeTollCCIPMsgArgs is a static function that always returns the abi.Arguments
// for a CCIP message.
func MakeTollCCIPMsgArgs() abi.Arguments {
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
			Name: "sender",
			Type: "address",
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
			Name: "feeTokenAndAmount",
			Type: "tuple",
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
			Name: "gasLimit",
			Type: "uint256",
		},
	}
	ty, _ := abi.NewType("tuple", "", tuples)
	return abi.Arguments{
		{
			Type: ty,
		},
	}
}

func ProofFlagsToBits(proofFlags []bool) *big.Int {
	// TODO: Support larger than int64
	var a int64
	for i := 0; i < len(proofFlags); i++ {
		if proofFlags[i] {
			a |= 1 << i
		}
	}
	return big.NewInt(a)
}

func makeTollExecutionReportArgs() abi.Arguments {
	return []abi.Argument{
		{
			Name: "ExecutionReport",
			Type: utils.MustAbiType("tuple", []abi.ArgumentMarshaling{
				{
					Name: "sequenceNumbers",
					Type: "uint64[]",
				},
				{
					Name: "tokenPerFeeCoinAddresses",
					Type: "address[]",
				},
				{
					Name: "tokenPerFeeCoin",
					Type: "uint256[]",
				},
				{
					Name: "encodedMessages",
					Type: "bytes[]",
				},
				{
					Name: "innerProofs",
					Type: "bytes32[]",
				},
				{
					Name: "innerProofFlagBits",
					Type: "uint256",
				},
				{
					Name: "outerProofs",
					Type: "bytes32[]",
				},
				{
					Name: "outerProofFlagBits",
					Type: "uint256",
				},
			}),
		},
	}
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
					Name: "tokenPerFeeCoinAddresses",
					Type: "address[]",
				},
				{
					Name: "tokenPerFeeCoin",
					Type: "uint256[]",
				},
				{
					Name: "feeUpdates",
					Type: "tuple[]",
					Components: []abi.ArgumentMarshaling{
						{
							Name: "chainId",
							Type: "uint64",
						},
						{
							Name: "linkPerUnitGas",
							Type: "uint256",
						},
					},
				},
				{
					Name: "encodedMessages",
					Type: "bytes[]",
				},
				{
					Name: "innerProofs",
					Type: "bytes32[]",
				},
				{
					Name: "innerProofFlagBits",
					Type: "uint256",
				},
				{
					Name: "outerProofs",
					Type: "bytes32[]",
				},
				{
					Name: "outerProofFlagBits",
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
					Name: "onRamps",
					Type: "address[]",
				},
				{
					Name: "intervals",
					Type: "tuple[]",
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
					Name: "merkleRoots",
					Type: "bytes32[]",
				},
				{
					Name: "rootOfRoots",
					Type: "bytes32",
				},
			}),
		},
	}
}
