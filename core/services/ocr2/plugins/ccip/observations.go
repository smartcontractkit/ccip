package ccip

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/core/logger"
)

func evmWord(i uint64) common.Hash {
	var b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return common.BigToHash(big.NewInt(0).SetBytes(b))
}

type Observation struct {
	MinSeqNum uint64 `json:"minSeqNum"`
	MaxSeqNum uint64 `json:"maxSeqNum"`
}

// getNonEmptyObservations checks the given observations for formatting and value errors.
// It returns all valid observations, potentially being an empty list. It will log
// malformed observations but never error.
func getNonEmptyObservations(l logger.Logger, observations []types.AttributedObservation) (nonEmptyObservations []Observation) {
	for _, ao := range observations {
		if len(ao.Observation) == 0 {
			// Empty observation
			l.Tracew("Discarded empty observation %+v", ao)
			continue
		}
		var ob Observation
		err := json.Unmarshal(ao.Observation, &ob)
		if err != nil {
			l.Errorw("Received unmarshallable observation", "err", err, "observation", string(ao.Observation))
			continue
		}
		nonEmptyObservations = append(nonEmptyObservations, ob)
	}
	return nonEmptyObservations
}

// getMinMaxSequenceNumbers retrieves the minimum and maximum sequence numbers for
// a given set of observations and F. F is an upper bound on the number of faulty nodes.
// This function can return an error on bad input or invalid results.
// Note this mutates the observation slice by sorting it.
func getMinMaxSequenceNumbers(observations []Observation, F int) (minSeqNum uint64, maxSeqNum uint64, err error) {
	if len(observations) <= F {
		return 0, 0, fmt.Errorf("number of observations (%d) too low for given F (%d)", len(observations), F)
	}
	// Extract the min and max
	sort.Slice(observations, func(i, j int) bool {
		return observations[i].MinSeqNum < observations[j].MinSeqNum
	})
	// r.F < len(nonEmptyObservations) because of the check above and therefore this is safe
	minSeqNum = observations[F].MinSeqNum

	sort.Slice(observations, func(i, j int) bool {
		return observations[i].MaxSeqNum < observations[j].MaxSeqNum
	})
	// We use a conservative maximum. If we pick a value that some honest oracles might not
	// have seen they’ll end up not agreeing on a msg, stalling the protocol.
	maxSeqNum = observations[F].MaxSeqNum

	if maxSeqNum < minSeqNum {
		return 0, 0, errors.New("max seq num smaller than min")
	}
	return
}
