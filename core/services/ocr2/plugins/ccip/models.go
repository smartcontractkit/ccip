package ccip

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lib/pq"

	"github.com/smartcontractkit/chainlink/core/utils"
)

type Request struct {
	SeqNum        utils.Big
	SourceChainID string // TODO Note this will be some super set which includes evm_chain_id
	DestChainID   string // TODO Note this will be some super set which includes evm_chain_id
	OnRamp        common.Address
	OffRamp       common.Address
	Sender        common.Address
	Receiver      common.Address
	Data          []byte
	Tokens        pq.StringArray
	Amounts       pq.StringArray
	Executor      common.Address
	Options       []byte
	Raw           []byte // Full ABI-encoded event for merkle tree
	Status        RequestStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func mustStringToBigInt(s string) *big.Int {
	i, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		panic(fmt.Sprintf("invalid big.Int string %v", s))
	}
	return i
}

func (r Request) ToMessage() Message {
	var tokens []common.Address
	for _, t := range r.Tokens {
		tokens = append(tokens, common.HexToAddress(t))
	}
	var amounts []*big.Int
	for _, a := range r.Amounts {
		amounts = append(amounts, mustStringToBigInt(a))
	}
	return Message{
		SequenceNumber: r.SeqNum.ToInt(),
		SourceChainId:  mustStringToBigInt(r.SourceChainID),
		Sender:         r.Sender,
		Payload: struct {
			Tokens             []common.Address `json:"tokens"`
			Amounts            []*big.Int       `json:"amounts"`
			DestinationChainId *big.Int         `json:"destinationChainId"`
			Receiver           common.Address   `json:"receiver"`
			Executor           common.Address   `json:"executor"`
			Data               []uint8          `json:"data"`
			Options            []uint8          `json:"options"`
		}{
			Tokens:             tokens,
			Amounts:            amounts,
			DestinationChainId: mustStringToBigInt(r.DestChainID),
			Receiver:           r.Receiver,
			Executor:           r.Executor,
			Data:               r.Data,
			Options:            r.Options,
		},
	}
}

type RequestStatus string

const (
	RequestStatusUnstarted    RequestStatus = "unstarted"
	RequestStatusRelayPending RequestStatus = "relay_pending"
	// RequestStatusRelayConfirmed is used after we've seen the report accepted log with sufficient
	// number of confirmations
	RequestStatusRelayConfirmed   RequestStatus = "relay_confirmed"
	RequestStatusExecutionPending RequestStatus = "execution_pending"
	// RequestStatusExecutionConfirmed is used after we've seen the Message executed log with sufficient
	// number of confirmations
	RequestStatusExecutionConfirmed RequestStatus = "execution_confirmed"
)

type RelayReport struct {
	Root      []byte
	MinSeqNum utils.Big
	MaxSeqNum utils.Big
	CreatedAt time.Time
}
