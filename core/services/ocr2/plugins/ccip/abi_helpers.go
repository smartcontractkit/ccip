package ccip

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp"
)

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
		SequenceNumber *big.Int       `json:"sequenceNumber"`
		SourceChainId  *big.Int       `json:"sourceChainId"`
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
		SequenceNumber: receivedCp.SequenceNumber,
		SourceChainId:  receivedCp.SourceChainId,
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
			Name: "sequenceNumber",
			Type: "uint256",
		},
		{
			Name: "sourceChainId",
			Type: "uint256",
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
