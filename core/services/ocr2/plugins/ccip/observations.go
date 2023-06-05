package ccip

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
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

type ExecutionObservation struct {
	Messages []ObservedMessage `json:"messages"`
}

type ObservedMessage struct {
	SeqNr     uint64   `json:"seqNr"`
	TokenData [][]byte `json:"tokenData"`
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
