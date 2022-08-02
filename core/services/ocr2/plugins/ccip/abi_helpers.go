package ccip

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/any_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/blob_verifier"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/evm_2_evm_subscription_onramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/utils"
)

var (
	// offset || sourceChainID || seqNum || ...
	CCIPSendRequested    common.Hash
	CCIPSubSendRequested common.Hash
	// merkleRoot || minSeqNum || maxSeqNum
	ReportAccepted common.Hash
	// sig || SeqNum || ...
	CrossChainMessageExecuted common.Hash
	ConfigSet                 common.Hash
)

// Zero indexed
const (
	SendRequestedSequenceNumberIndex             = 2 // Valid for both toll and sub
	ReportAcceptedMinSequenceNumberIndex         = 1
	ReportAcceptedMaxSequenceNumberIndex         = 2
	CrossChainMessageExecutedSequenceNumberIndex = 1
)

func init() {
	getIDOrPanic := func(name string, abi2 abi.ABI) common.Hash {
		event, ok := abi2.Events[name]
		if !ok {
			panic(fmt.Sprintf("missing event %s", name))
		}
		return event.ID
	}
	onRampABI, err := abi.JSON(strings.NewReader(evm_2_evm_toll_onramp.EVM2EVMTollOnRampABI))
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
	blobVerifierABI, err := abi.JSON(strings.NewReader(blob_verifier.BlobVerifierABI))
	if err != nil {
		panic(err)
	}
	CCIPSendRequested = getIDOrPanic("CCIPSendRequested", onRampABI)
	CCIPSubSendRequested = getIDOrPanic("CCIPSendRequested", subOnRampABI)
	ReportAccepted = getIDOrPanic("ReportAccepted", blobVerifierABI)
	CrossChainMessageExecuted = getIDOrPanic("ExecutionStateChanged", offRampABI)
	ConfigSet = getIDOrPanic("ConfigSet", blobVerifierABI)
}

// DecodeCCIPMessage decodes the bytecode message into a blob_verifier.CCIPAny2EVMTollMessage
// This function returns an error if there is no message in the bytecode or
// when the payload is malformed.
func DecodeCCIPMessage(b []byte) (*evm_2_evm_toll_onramp.CCIPEVM2EVMTollEvent, error) {
	unpacked, err := MakeCCIPMsgArgs().Unpack(b)
	if err != nil {
		return nil, err
	}
	if len(unpacked) == 0 {
		return nil, fmt.Errorf("no message found when unpacking")
	}
	// Note must use unnamed type here
	receivedCp, ok := unpacked[0].(struct {
		SourceChainId  *big.Int         `json:"sourceChainId"`
		SequenceNumber uint64           `json:"sequenceNumber"`
		Sender         common.Address   `json:"sender"`
		Receiver       common.Address   `json:"receiver"`
		Data           []uint8          `json:"data"`
		Tokens         []common.Address `json:"tokens"`
		Amounts        []*big.Int       `json:"amounts"`
		FeeToken       common.Address   `json:"feeToken"`
		FeeTokenAmount *big.Int         `json:"feeTokenAmount"`
		GasLimit       *big.Int         `json:"gasLimit"`
	})
	if !ok {
		return nil, fmt.Errorf("invalid format have %T want %T", unpacked[0], receivedCp)
	}
	return &evm_2_evm_toll_onramp.CCIPEVM2EVMTollEvent{
		SourceChainId:  receivedCp.SourceChainId,
		SequenceNumber: receivedCp.SequenceNumber,
		Sender:         receivedCp.Sender,
		Receiver:       receivedCp.Receiver,
		Data:           receivedCp.Data,
		Tokens:         receivedCp.Tokens,
		Amounts:        receivedCp.Amounts,
		FeeToken:       receivedCp.FeeToken,
		FeeTokenAmount: receivedCp.FeeTokenAmount,
		GasLimit:       receivedCp.GasLimit,
	}, nil
}

func EVM2EVMTollEventToMessage(event evm_2_evm_toll_onramp.CCIPEVM2EVMTollEvent) Message {
	return Message{
		SourceChainId:  event.SourceChainId,
		SequenceNumber: event.SequenceNumber,
		Sender:         event.Sender,
		Receiver:       event.Receiver,
		Data:           event.Data,
		Tokens:         event.Tokens,
		Amounts:        event.Amounts,
		FeeToken:       event.FeeToken,
		FeeTokenAmount: event.FeeTokenAmount,
		GasLimit:       event.GasLimit,
	}
}

// MakeCCIPMsgArgs is a static function that always returns the abi.Arguments
// for a CCIP message.
func MakeCCIPMsgArgs() abi.Arguments {
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
			Name: "tokens",
			Type: "address[]",
		},
		{
			Name: "amounts",
			Type: "uint256[]",
		},
		{
			Name: "feeToken",
			Type: "address",
		},
		{
			Name: "feeTokenAmount",
			Type: "uint256",
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
	SourceChainId  *big.Int         `json:"sourceChainId"`
	SequenceNumber uint64           `json:"sequenceNumber"`
	Sender         common.Address   `json:"sender"`
	Receiver       common.Address   `json:"receiver"`
	Data           []uint8          `json:"data"`
	Tokens         []common.Address `json:"tokens"`
	Amounts        []*big.Int       `json:"amounts"`
	FeeToken       common.Address   `json:"feeToken"`
	FeeTokenAmount *big.Int         `json:"feeTokenAmount"`
	GasLimit       *big.Int         `json:"gasLimit"`
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

func makeRelayReportArgs() abi.Arguments {
	return []abi.Argument{
		{
			Name: "RelayReport",
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
