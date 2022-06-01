package ccip

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/onramp"
	"github.com/smartcontractkit/chainlink/core/utils"
)

var (
	// offset || sourceChainID || seqNum || ...
	CCIPSendRequested common.Hash
	// merkleRoot || minSeqNum || maxSeqNum
	ReportAccepted common.Hash
	// SeqNum
	CrossChainMessageExecuted common.Hash
	ConfigSet                 common.Hash
)

// Zero indexed
const (
	SendRequestedSequenceNumberIndex             = 2
	ReportAcceptedMinSequenceNumberIndex         = 1
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
	onRampABI, err := abi.JSON(strings.NewReader(onramp.OnRampABI))
	if err != nil {
		panic(err)
	}
	offRampABI, err := abi.JSON(strings.NewReader(offramp.OffRampABI))
	if err != nil {
		panic(err)
	}
	CCIPSendRequested = getIDOrPanic("CCIPSendRequested", onRampABI)
	ReportAccepted = getIDOrPanic("ReportAccepted", offRampABI)
	CrossChainMessageExecuted = getIDOrPanic("CrossChainMessageExecuted", offRampABI)
	ConfigSet = getIDOrPanic("ConfigSet", offRampABI)
}

// DecodeCCIPMessage decodes the bytecode message into an offramp.CCIPAnyToEVMTollMessage
// This function returns an error if there is no message in the bytecode or
// when the payload is malformed.
func DecodeCCIPMessage(b []byte) (*offramp.CCIPAnyToEVMTollMessage, error) {
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
	return &offramp.CCIPAnyToEVMTollMessage{
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

func EVMToEVMTollEventToMessage(event onramp.CCIPEVMToEVMTollEvent) Message {
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

type ExecutionReport struct {
	Messages      []Message  `json:"messages"`
	Proofs        [][32]byte `json:"proofs"`
	ProofFlagBits *big.Int   `json:"proofFlagBits"`
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
					Name: "Messages",
					Type: "tuple[]",
					Components: []abi.ArgumentMarshaling{
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
					},
				},
				{
					Name: "Proofs",
					Type: "bytes32[]",
				},
				{
					Name: "ProofFlagBits",
					Type: "uint256",
				},
			}),
		},
	}
}

func makeRelayReportArgs() abi.Arguments {
	return []abi.Argument{
		{
			Name: "merkleRoot",
			Type: utils.MustAbiType("bytes32", nil),
		},
		{
			Name: "minSequenceNumber",
			Type: utils.MustAbiType("uint64", nil),
		},
		{
			Name: "maxSequenceNumber",
			Type: utils.MustAbiType("uint64", nil),
		},
	}
}
