package commit

import (
	"sort"

	"github.com/pkg/errors"

	"github.com/smartcontractkit/ccipocr3/internal/model"
)

var ErrOverlappingRanges = errors.New("overlapping sequence numbers in reports")

// computeRanges takes a slice of reports and computes the smallest number of contiguous ranges
// that cover all the sequence numbers in the reports.
func computeRanges(reports []model.ExecutePluginCommitData) ([]model.SeqNumRange, error) {
	var ranges []model.SeqNumRange

	if len(reports) == 0 {
		return nil, nil
	}

	var seqRange model.SeqNumRange
	for i, report := range reports {
		if i == 0 {
			// initialize
			seqRange = model.NewSeqNumRange(report.SequenceNumberRange.Start(), report.SequenceNumberRange.End())
		} else if seqRange.End()+1 == report.SequenceNumberRange.Start() {
			// extend the contiguous range
			seqRange.SetEnd(report.SequenceNumberRange.End())
		} else if report.SequenceNumberRange.Start() < seqRange.End() {
			return nil, ErrOverlappingRanges
		} else {
			ranges = append(ranges, seqRange)

			// Reset the range.
			seqRange = model.NewSeqNumRange(report.SequenceNumberRange.Start(), report.SequenceNumberRange.End())
		}
	}
	// add final range
	ranges = append(ranges, seqRange)

	return ranges, nil
}

func groupByChainSelector(reports []model.CommitPluginReportWithMeta) model.ExecutePluginCommitObservations {
	commitReportCache := make(map[model.ChainSelector][]model.ExecutePluginCommitData)
	for _, report := range reports {
		for _, singleReport := range report.Report.MerkleRoots {
			commitReportCache[singleReport.ChainSel] = append(commitReportCache[singleReport.ChainSel], model.ExecutePluginCommitData{
				Timestamp:           report.Timestamp,
				BlockNum:            report.BlockNum,
				MerkleRoot:          singleReport.MerkleRoot,
				SequenceNumberRange: singleReport.SeqNumsRange,
				ExecutedMessages:    nil,
			})
		}
	}
	return commitReportCache
}

// filterOutFullyExecutedMessages returns a new reports slice with fully executed messages removed.
func filterOutFullyExecutedMessages(reports []model.ExecutePluginCommitData, executedMessages []model.SeqNumRange) ([]model.ExecutePluginCommitData, error) {
	// If none are executed, return the input.
	if len(executedMessages) == 0 {
		return reports, nil
	}

	sort.Slice(executedMessages, func(i, j int) bool {
		return executedMessages[i].Start() < executedMessages[j].Start()
	})
	// Make sure they do not overlap
	previousMax := model.SeqNum(0)
	for _, seqRange := range executedMessages {
		if seqRange.Start() < previousMax {
			return nil, ErrOverlappingRanges
		}
		previousMax = seqRange.End()
	}

	var filtered []model.ExecutePluginCommitData

	reportIdx := 0
	for _, executed := range executedMessages {
		for i := reportIdx; i < len(reports); i++ {
			reportRange := reports[i].SequenceNumberRange
			if executed.End() < reportRange.Start() {
				// need to go to the next set of executed messages.
				break
			}

			reportIdx++

			if reportRange.Start() >= executed.Start() && reportRange.End() <= executed.End() {
				// fully executed report
				continue
			}

			filtered = append(filtered, reports[i])
		}
	}

	// Add any remaining reports that were not fully executed.
	for i := reportIdx; i < len(reports); i++ {
		filtered = append(filtered, reports[i])
	}

	return filtered, nil
}
