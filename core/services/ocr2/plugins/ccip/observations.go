package ccip

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

type CommitObservation struct {
	Interval          commit_store.CommitStoreInterval `json:"interval"`
	TokenPricesUSD    map[common.Address]*big.Int      `json:"tokensPerFeeCoin"`
	SourceGasPriceUSD *big.Int                         `json:"sourceGasPrice"`
}

func (o CommitObservation) Marshal() ([]byte, error) {
	return json.Marshal(&o)
}

// ExecutionObservation stores messages as a map pointing from a sequence number (uint) to the message payload (MsgData)
// Having it structured this way is critical because:
// * it prevents having duplicated sequence numbers within a single ExecutionObservation (compared to the list representation)
// * prevents malicious actors from passing multiple messages with the same sequence number
type ExecutionObservation struct {
	Messages map[uint64]MsgData `json:"messages"`
}

type MsgData struct {
	TokenData [][]byte `json:"tokenData"`
}

// ObservedMessage is a transient struct used for processing convenience within the plugin. It's easier to process observed messages
// when all properties are flattened into a single structure.
// It should not be serialized and returned from types.ReportingPlugin functions, please serialize/deserialize to/from ExecutionObservation instead using NewObservedMessage
type ObservedMessage struct {
	SeqNr uint64
	MsgData
}

func NewExecutionObservation(observations []ObservedMessage) ExecutionObservation {
	denormalized := make(map[uint64]MsgData, len(observations))
	for _, o := range observations {
		denormalized[o.SeqNr] = MsgData{TokenData: o.TokenData}
	}
	return ExecutionObservation{Messages: denormalized}
}

func NewObservedMessage(seqNr uint64, tokenData [][]byte) ObservedMessage {
	return ObservedMessage{
		SeqNr:   seqNr,
		MsgData: MsgData{TokenData: tokenData},
	}
}

func (o ExecutionObservation) Marshal() ([]byte, error) {
	return json.Marshal(&o)
}

// getParsableObservations checks the given observations for formatting and value errors.
// It returns all valid observations, potentially being an empty list. It will log
// malformed observations but never error.
func getParsableObservations[O CommitObservation | ExecutionObservation](l logger.Logger, observations []types.AttributedObservation) []O {
	var parseableObservations []O
	for _, ao := range observations {
		if len(ao.Observation) == 0 {
			// Empty observation
			l.Infow("Discarded empty observation %+v", ao)
			continue
		}
		var ob O
		err := json.Unmarshal(ao.Observation, &ob)
		if err != nil {
			l.Errorw("Received unmarshallable observation", "err", err, "observation", string(ao.Observation))
			continue
		}
		parseableObservations = append(parseableObservations, ob)
	}
	return parseableObservations
}
