package ccip

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/onramp"
)

var (
	// offset || sourceChainID || seqNum || ...
	CrossChainSendRequested common.Hash
	// merkleRoot || minSeqNum || maxSeqNum
	ReportAccepted common.Hash
	// SeqNum
	CrossChainMessageExecuted common.Hash
	ConfigSet                 common.Hash
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
	CrossChainSendRequested = getIDOrPanic("CrossChainSendRequested", onRampABI)
	ReportAccepted = getIDOrPanic("ReportAccepted", offRampABI)
	CrossChainMessageExecuted = getIDOrPanic("CrossChainMessageExecuted", offRampABI)
	ConfigSet = getIDOrPanic("ConfigSet", offRampABI)
}

// DecodeCCIPMessage decodes the bytecode message into an offramp.CCIPMessage
// This function returns an error if there is no message in the bytecode or
// when the payload is malformed.
func DecodeCCIPMessage(b []byte) (*offramp.CCIPMessage, error) {
	unpacked, err := MakeCCIPMsgArgs().Unpack(b)
	if err != nil {
		return nil, err
	}
	if len(unpacked) == 0 {
		return nil, fmt.Errorf("no message found when unpacking")
	}
	// Note must use unnamed type here
	receivedCp, ok := unpacked[0].(struct {
		SourceChainId  *big.Int       `json:"sourceChainId"`
		SequenceNumber uint64         `json:"sequenceNumber"`
		Sender         common.Address `json:"sender"`
		Payload        struct {
			Tokens             []common.Address `json:"tokens"`
			Amounts            []*big.Int       `json:"amounts"`
			DestinationChainId *big.Int         `json:"destinationChainId"`
			Receiver           common.Address   `json:"receiver"`
			Executor           common.Address   `json:"executor"`
			Data               []uint8          `json:"data"`
			Options            []uint8          `json:"options"`
		} `json:"payload"`
	})
	if !ok {
		return nil, fmt.Errorf("invalid format have %T want %T", unpacked[0], receivedCp)
	}
	return &offramp.CCIPMessage{
		SourceChainId:  receivedCp.SourceChainId,
		SequenceNumber: receivedCp.SequenceNumber,
		Sender:         receivedCp.Sender,
		Payload: offramp.CCIPMessagePayload{
			DestinationChainId: receivedCp.Payload.DestinationChainId,
			Receiver:           receivedCp.Payload.Receiver,
			Data:               receivedCp.Payload.Data,
			Tokens:             receivedCp.Payload.Tokens,
			Amounts:            receivedCp.Payload.Amounts,
			Executor:           receivedCp.Payload.Executor,
			Options:            receivedCp.Payload.Options,
		},
	}, nil
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
			Name: "payload",
			Type: "tuple",
			Components: []abi.ArgumentMarshaling{
				{
					Name: "tokens",
					Type: "address[]",
				},
				{
					Name: "amounts",
					Type: "uint256[]",
				},
				{
					Name: "destinationChainId",
					Type: "uint256",
				},
				{
					Name: "receiver",
					Type: "address",
				},
				{
					Name: "executor",
					Type: "address",
				},
				{
					Name: "data",
					Type: "bytes",
				},
				{
					Name: "options",
					Type: "bytes",
				},
			},
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
	SourceChainId  *big.Int       `json:"sourceChainId"`
	SequenceNumber uint64         `json:"sequenceNumber"`
	Sender         common.Address `json:"sender"`
	Payload        struct {
		Tokens             []common.Address `json:"tokens"`
		Amounts            []*big.Int       `json:"amounts"`
		DestinationChainId *big.Int         `json:"destinationChainId"`
		Receiver           common.Address   `json:"receiver"`
		Executor           common.Address   `json:"executor"`
		Data               []uint8          `json:"data"`
		Options            []uint8          `json:"options"`
	} `json:"payload"`
}

type ExecutableMessage struct {
	Path    [][32]byte `json:"path"`
	Index   *big.Int   `json:"index"`
	Message Message    `json:"message"`
}

func makeExecutionReportArgs() abi.Arguments {
	mustType := func(ts string, components []abi.ArgumentMarshaling) abi.Type {
		ty, _ := abi.NewType(ts, "", components)
		return ty
	}
	return []abi.Argument{
		{
			Name: "executableMessages",
			Type: mustType("tuple[]", []abi.ArgumentMarshaling{
				{
					Name: "Path",
					Type: "bytes32[]",
				},
				{
					Name: "Index",
					Type: "uint256",
				},
				{
					Name: "Message",
					Type: "tuple",
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
							Name: "payload",
							Type: "tuple",
							Components: []abi.ArgumentMarshaling{
								{
									Name: "tokens",
									Type: "address[]",
								},
								{
									Name: "amounts",
									Type: "uint256[]",
								},
								{
									Name: "destinationChainId",
									Type: "uint256",
								},
								{
									Name: "receiver",
									Type: "address",
								},
								{
									Name: "executor",
									Type: "address",
								},
								{
									Name: "data",
									Type: "bytes",
								},
								{
									Name: "options",
									Type: "bytes",
								},
							},
						},
					},
				},
			}),
		},
	}
}

func makeRelayReportArgs() abi.Arguments {
	mustType := func(ts string) abi.Type {
		ty, _ := abi.NewType(ts, "", nil)
		return ty
	}
	return []abi.Argument{
		{
			Name: "merkleRoot",
			Type: mustType("bytes32"),
		},
		{
			Name: "minSequenceNumber",
			Type: mustType("uint64"),
		},
		{
			Name: "maxSequenceNumber",
			Type: mustType("uint64"),
		},
	}
}
