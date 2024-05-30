package commit

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/ccipocr3/internal/model"
	"github.com/smartcontractkit/libocr/commontypes"
)

// validateObserverReadingEligibility checks if the observer is eligible to observe the messages it observed.
func validateObserverReadingEligibility(
	observer commontypes.OracleID,
	observerCfg map[commontypes.OracleID]model.ObserverInfo,
	observedMsgs map[model.ChainSelector][]model.ExecutePluginCCIPData,
) error {
	observerInfo, exists := observerCfg[observer]
	if !exists {
		return fmt.Errorf("observer not found in config")
	}

	observerReadChains := mapset.NewSet(observerInfo.Reads...)

	for chainSel, msgs := range observedMsgs {
		if len(msgs) == 0 {
			continue
		}

		if !observerReadChains.Contains(chainSel) {
			return fmt.Errorf("observer not allowed to read from chain %d", chainSel)
		}
	}

	return nil
}

// validateObservedSequenceNumbers checks if the sequence numbers of the provided messages are unique for each chain and
// that they match the observed max sequence numbers.
func validateObservedSequenceNumbers(observedData map[model.ChainSelector][]model.ExecutePluginCommitData) error {
	for _, commitData := range observedData {
		// observed commitData must not contain duplicates

		observedMerkleRoots := mapset.NewSet[string]()
		observedRanges := make([]model.SeqNumRange, 0)

		for _, data := range commitData {
			rootStr := data.MerkleRoot.String()
			if observedMerkleRoots.Contains(rootStr) {
				return fmt.Errorf("duplicate merkle root %s observed", rootStr)
			}
			observedMerkleRoots.Add(rootStr)

			for _, rng := range observedRanges {
				if rng.Overlaps(data.SequenceNumberRange) {
					return fmt.Errorf("sequence number range %v overlaps with %v", data.SequenceNumberRange, rng)
				}
			}
			observedRanges = append(observedRanges, data.SequenceNumberRange)

			// Executed sequence numbers should belong in the observed range.
			for _, seqNum := range data.ExecutedMessages {
				if !data.SequenceNumberRange.Contains(seqNum) {
					return fmt.Errorf("executed message %d not in observed range %v", seqNum, data.SequenceNumberRange)
				}
			}
		}
	}

	return nil
}
