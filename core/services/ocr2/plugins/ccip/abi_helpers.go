package ccip

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_subscription_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/utils"
)

var (
	// offset || sourceChainID || seqNum || ...
	CCIPTollSendRequested common.Hash
	CCIPSubSendRequested  common.Hash
	// merkleRoot || minSeqNum || maxSeqNum
	ReportAccepted common.Hash
	// sig || SeqNum || ...
	ExecutionStateChanged common.Hash
	ConfigSet             common.Hash
)

// Zero indexed
const (
	SendRequestedSequenceNumberIndex             = 2 // Valid for both toll and sub
	ReportAcceptedMinSequenceNumberIndex         = 1
	ReportAcceptedMaxSequenceNumberIndex         = 2
	CrossChainMessageExecutedSequenceNumberIndex = 1
)

// MessageExecutionState defines the execution states of CCIP messages.
type MessageExecutionState uint64

const (
	Untouched MessageExecutionState = iota
	InProgress
	Success
	Failure
)

func init() {
	getIDOrPanic := func(name string, abi2 abi.ABI) common.Hash {
		event, ok := abi2.Events[name]
		if !ok {
			panic(fmt.Sprintf("missing event %s", name))
		}
		return event.ID
	}
	tollOnRampABI, err := abi.JSON(strings.NewReader(evm_2_evm_toll_onramp.EVM2EVMTollOnRampABI))
	if err != nil {
		panic(err)
	}
	subOnRampABI, err := abi.JSON(strings.NewReader(evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampABI))
	if err != nil {
		panic(err)
	}
	offRampABI, err := abi.JSON(strings.NewReader(any_2_evm_toll_offramp.EVM2EVMTollOffRampABI))
	if err != nil {
		panic(err)
	}
	commitStoreABI, err := abi.JSON(strings.NewReader(commit_store.CommitStoreABI))
	if err != nil {
		panic(err)
	}
	CCIPTollSendRequested = getIDOrPanic("CCIPSendRequested", tollOnRampABI)
	CCIPSubSendRequested = getIDOrPanic("CCIPSendRequested", subOnRampABI)
	ReportAccepted = getIDOrPanic("ReportAccepted", commitStoreABI)
	ExecutionStateChanged = getIDOrPanic("ExecutionStateChanged", offRampABI)
	ConfigSet = getIDOrPanic("ConfigSet", commitStoreABI)
}

// DecodeCCIPMessage decodes the bytecode message into a commit_store.CCIPAny2EVMTollMessage
// This function returns an error if there is no message in the bytecode or
// when the payload is malformed.
func DecodeCCIPMessage(b []byte) (*evm_2_evm_toll_onramp.CCIPEVM2EVMTollMessage, error) {
	unpacked, err := MakeTollCCIPMsgArgs().Unpack(b)
	if err != nil {
		return nil, err
	}
	if len(unpacked) == 0 {
		return nil, fmt.Errorf("no message found when unpacking")
	}
	// Note must use unnamed type here
	receivedCp, ok := unpacked[0].(struct {
		SourceChainId    *big.Int       `json:"sourceChainId"`
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

	var tokensAndAmounts []evm_2_evm_toll_onramp.CCIPEVMTokenAndAmount

	for _, tokenAndAmount := range receivedCp.TokensAndAmounts {
		tokensAndAmounts = append(tokensAndAmounts, evm_2_evm_toll_onramp.CCIPEVMTokenAndAmount{
			Token:  tokenAndAmount.Token,
			Amount: tokenAndAmount.Amount,
		})
	}

	return &evm_2_evm_toll_onramp.CCIPEVM2EVMTollMessage{
		SourceChainId:    receivedCp.SourceChainId,
		SequenceNumber:   receivedCp.SequenceNumber,
		Sender:           receivedCp.Sender,
		Receiver:         receivedCp.Receiver,
		Data:             receivedCp.Data,
		TokensAndAmounts: tokensAndAmounts,
		FeeTokenAndAmount: evm_2_evm_toll_onramp.CCIPEVMTokenAndAmount{
			Token:  receivedCp.FeeTokenAndAmount.Token,
			Amount: receivedCp.FeeTokenAndAmount.Amount,
		},
		GasLimit: receivedCp.GasLimit,
	}, nil
}

func DecodeCCIPSubMessage(b []byte) (*evm_2_evm_subscription_onramp.CCIPEVM2EVMSubscriptionMessage, error) {
	unpacked, err := MakeSubscriptionCCIPMsgArgs().Unpack(b)
	if err != nil {
		return nil, err
	}
	if len(unpacked) == 0 {
		return nil, fmt.Errorf("no message found when unpacking")
	}
	// Note must use unnamed type here
	receivedCp, ok := unpacked[0].(struct {
		SourceChainId    *big.Int       `json:"sourceChainId"`
		SequenceNumber   uint64         `json:"sequenceNumber"`
		Sender           common.Address `json:"sender"`
		Receiver         common.Address `json:"receiver"`
		Nonce            uint64         `json:"nonce"`
		Data             []byte         `json:"data"`
		TokensAndAmounts []struct {
			Token  common.Address `json:"token"`
			Amount *big.Int       `json:"amount"`
		} `json:"tokensAndAmounts"`
		GasLimit *big.Int `json:"gasLimit"`
	})
	if !ok {
		return nil, fmt.Errorf("invalid format have %T want %T", unpacked[0], receivedCp)
	}

	var tokensAndAmounts []evm_2_evm_subscription_onramp.CCIPEVMTokenAndAmount

	for _, tokenAndAmount := range receivedCp.TokensAndAmounts {
		tokensAndAmounts = append(tokensAndAmounts, evm_2_evm_subscription_onramp.CCIPEVMTokenAndAmount{
			Token:  tokenAndAmount.Token,
			Amount: tokenAndAmount.Amount,
		})
	}

	return &evm_2_evm_subscription_onramp.CCIPEVM2EVMSubscriptionMessage{
		SourceChainId:    receivedCp.SourceChainId,
		SequenceNumber:   receivedCp.SequenceNumber,
		Sender:           receivedCp.Sender,
		Receiver:         receivedCp.Receiver,
		Nonce:            receivedCp.Nonce,
		Data:             receivedCp.Data,
		TokensAndAmounts: tokensAndAmounts,
		GasLimit:         receivedCp.GasLimit,
	}, nil
}

// MakeTollCCIPMsgArgs is a static function that always returns the abi.Arguments
// for a CCIP message.
func MakeTollCCIPMsgArgs() abi.Arguments {
	var tuples = []abi.ArgumentMarshaling{
		{
			Name: "sourceChainId",
			Type: "uint256",
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

func MakeSubscriptionCCIPMsgArgs() abi.Arguments {
	var tuples = []abi.ArgumentMarshaling{
		{
			Name: "sourceChainId",
			Type: "uint256",
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
			Name: "nonce",
			Type: "uint64",
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

// Message contains the data from a cross chain message
type Message struct {
	SourceChainId     *big.Int                                      `json:"sourceChainId"`
	SequenceNumber    uint64                                        `json:"sequenceNumber"`
	Sender            common.Address                                `json:"sender"`
	Receiver          common.Address                                `json:"receiver"`
	Data              []uint8                                       `json:"data"`
	TokensAndAmounts  []evm_2_evm_toll_onramp.CCIPEVMTokenAndAmount `json:"tokensAndAmounts"`
	FeeTokenAndAmount evm_2_evm_toll_onramp.CCIPEVMTokenAndAmount   `json:"feeTokenAndAmount"`
	GasLimit          *big.Int                                      `json:"gasLimit"`
}

func ProofFlagsToBits(proofFlags []bool) *big.Int {
	// TODO: Support larger than int64?
	var a int64
	for i := 0; i < len(proofFlags); i++ {
		if proofFlags[i] {
			a |= 1 << i
		}
	}
	return big.NewInt(a)
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
							Type: "uint256",
						},
						{
							Name: "gasPrice",
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
