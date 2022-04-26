package ccip

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/core/logger"
)

// NoRequestsToProcess indicates an empty observation. We use -1 as any value below zero would
// indicate a failure and therefore this number range is safe to use.
var NoRequestsToProcess = int64(-1)

type Observation struct {
	MinSeqNum int64 `json:"minSeqNum"`
	MaxSeqNum int64 `json:"maxSeqNum"`
}

// getNonEmptyObservations checks the given observations for formatting and value errors.
// It returns all valid observations, potentially being an empty list. It will log
// malformed observations but never error.
func getNonEmptyObservations(l logger.Logger, observations []types.AttributedObservation) (nonEmptyObservations []Observation) {
	for _, ao := range observations {
		var ob Observation
		err := json.Unmarshal(ao.Observation, &ob)
		if err != nil {
			l.Errorw("Received unmarshallable observation", "err", err, "observation", string(ao.Observation))
			continue
		}
		if ob.MinSeqNum < 0 {
			if ob.MinSeqNum == NoRequestsToProcess {
				l.Tracew("Discarded empty observation %+v", ao)
			} else {
				l.Warnf("Discarded invalid observation %+v", ao)
			}
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
func getMinMaxSequenceNumbers(observations []Observation, F int) (minSeqNum int64, maxSeqNum int64, err error) {
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
	// have seen they’ll end up not agreeing on a report, stalling the protocol.
	maxSeqNum = observations[F].MaxSeqNum

	if maxSeqNum < minSeqNum {
		return 0, 0, errors.New("max seq num smaller than min")
	}
	return
}
