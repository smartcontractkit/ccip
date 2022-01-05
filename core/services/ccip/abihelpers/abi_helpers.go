package abihelpers

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
)

func DecodeCCIPMessage(b []byte) (*single_token_offramp.CCIPMessage, error) {
	unpacked, err := MakeCCIPMsgArgs().Unpack(b)
	if err != nil {
		return nil, err
	}
	// Note must use unnamed type here
	receivedCp, ok := unpacked[0].(struct {
		SequenceNumber     *big.Int       `json:"sequenceNumber"`
		SourceChainId      *big.Int       `json:"sourceChainId"`
		DestinationChainId *big.Int       `json:"destinationChainId"`
		Sender             common.Address `json:"sender"`
		Payload            struct {
			Receiver common.Address   `json:"receiver"`
			Data     []uint8          `json:"data"`
			Tokens   []common.Address `json:"tokens"`
			Amounts  []*big.Int       `json:"amounts"`
			Executor common.Address   `json:"executor"`
			Options  []uint8          `json:"options"`
		} `json:"payload"`
	})
	if !ok {
		return nil, fmt.Errorf("invalid format have %T want %T", unpacked[0], receivedCp)
	}
	return &single_token_offramp.CCIPMessage{
		SequenceNumber:     receivedCp.SequenceNumber,
		SourceChainId:      receivedCp.SourceChainId,
		DestinationChainId: receivedCp.DestinationChainId,
		Sender:             receivedCp.Sender,
		Payload: single_token_offramp.CCIPMessagePayload{
			Receiver: receivedCp.Payload.Receiver,
			Data:     receivedCp.Payload.Data,
			Tokens:   receivedCp.Payload.Tokens,
			Amounts:  receivedCp.Payload.Amounts,
			Executor: receivedCp.Payload.Executor,
			Options:  receivedCp.Payload.Options,
		},
	}, nil
}

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
			Name: "destinationChainId",
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
					Name: "executor",
					Type: "address",
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
