package commit

import (
	"context"
	"time"

	//cache "github.com/smartcontractkit/ccipocr3/internal/copypaste/commit_roots_cache"
	"github.com/smartcontractkit/ccipocr3/internal/model"
	"github.com/smartcontractkit/ccipocr3/internal/reader"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

// Plugin implements the main ocr3 plugin logic.
type Plugin struct {
	nodeID     commontypes.OracleID
	cfg        model.ExecutePluginConfig
	ccipReader reader.CCIP

	//commitRootsCache cache.CommitsRootsCache
	//lastReportBlock uint64
	lastReportTS time.Time
}

func NewPlugin(
	_ context.Context,
	nodeID commontypes.OracleID,
	cfg model.ExecutePluginConfig,
	ccipReader reader.CCIP,
) *Plugin {
	return &Plugin{
		nodeID:       nodeID,
		cfg:          cfg,
		ccipReader:   ccipReader,
		lastReportTS: time.Now().Add(-8 * time.Hour), // TODO: get this from a config
	}
}

func (p *Plugin) Query(ctx context.Context, outctx ocr3types.OutcomeContext) (types.Query, error) {
	return types.Query{}, nil
}

func groupByChainSelector(reports []model.CommitPluginReport) map[model.ChainSelector][]model.ExecutePluginCommitData {
	commitReportCache := make(map[model.ChainSelector][]model.ExecutePluginCommitData)
	for _, reports := range reports {
		for _, singleReport := range reports.MerkleRoots {
			commitReportCache[singleReport.ChainSel] = append(commitReportCache[singleReport.ChainSel], model.ExecutePluginCommitData{
				MerkleRoot:          singleReport.MerkleRoot,
				SequenceNumberRange: singleReport.SeqNumsRange,
				ExecutedMessages:    nil,
			})
		}
	}
	return commitReportCache
}

func readAndAppendNextRange(ctx context.Context, ccipReader reader.CCIP, messages []model.ExecutePluginCCIPData, selector model.ChainSelector, seqRange model.SeqNumRange) ([]model.ExecutePluginCCIPData, error) {
	msgs, err := ccipReader.MsgsBetweenSeqNums(ctx, []model.ChainSelector{selector}, seqRange)
	if err != nil {
		return nil, err
	}
	var convert []model.ExecutePluginCCIPData
	for _, msg := range msgs {
		convert = append(convert, model.ExecutePluginCCIPData{
			SequenceNumber: msg.SeqNum,
			Message:        msg.ID,
		})
	}
	return messages, nil
}

// Observation collects data across two phases which happen in separate rounds.
// These phases happen continuously so that except for the first round, every
// subsequent round can have a new execution report.
//
// Phase 1: Gather commit reports from the destination chain and determine
// which messages are required to build a valid execution report.
//
// Phase 2: Gather messages from the source chains and build the execution
// report.
func (p *Plugin) Observation(ctx context.Context, outctx ocr3types.OutcomeContext, _ types.Query) (types.Observation, error) {
	previousOutcome, err := model.DecodeExecutePluginOutcome(outctx.PreviousOutcome)
	if err != nil {
		return types.Observation{}, err
	}

	// Phase 1: Gather commit reports from the destination chain and determine which messages are required to build a valid execution report.
	// TODO: filter out "cannot read p.destChain" errors? Or avoid calling it in the first place?
	commitReports, err := p.ccipReader.CommitReportsGTETimestamp(ctx, p.cfg.DestChain, p.lastReportTS, 1000)
	if err != nil {
		return types.Observation{}, err
	}
	if len(commitReports) > 0 {
		//lastReport := commitReports[len(commitReports)-1]
		//p.lastReportTS = lastReport.
		// TODO: Need a way to get a timestamp of the report.
	}

	// Phase 2: Gather messages from the source chains and build the execution report.
	messages := make(map[model.ChainSelector][]model.ExecutePluginCCIPData)
	if len(previousOutcome.Messages) == 0 {
		// No messages to execute.
		// This is expected after a cold start.
	} else {
		for selector, reports := range previousOutcome.NextCommits {
			if len(reports) == 0 {
				continue
			}

			// The total number of reads are minimized by grouping together contiguous ranges.
			// For new reports, we expect all sequence numbers to be sequential. Handling for
			// non-contiguous ranges is also implemented to handle older reports when necessary.
			var seqRange model.SeqNumRange
			for i, report := range reports {
				if i == 0 {
					// initialize
					seqRange.SetStart(report.SequenceNumberRange.Start())
					seqRange.SetEnd(report.SequenceNumberRange.End())
				} else if report.SequenceNumberRange.Start()-1 == seqRange.End() {
					// extend the contiguous range
					seqRange.SetEnd(report.SequenceNumberRange.End())
				} else {
					// non-contiguous range detected, make a request for the contiguous range.
					messages[selector], err = readAndAppendNextRange(ctx, p.ccipReader, messages[selector], selector, seqRange)
					if err != nil {
						return types.Observation{}, err
					}

					// Reset the range.
					seqRange.SetStart(report.SequenceNumberRange.Start())
					seqRange.SetEnd(report.SequenceNumberRange.End())
				}
			}

			// Append the last contiguous range.
			messages[selector], err = readAndAppendNextRange(ctx, p.ccipReader, messages[selector], selector, seqRange)
			if err != nil {
				return types.Observation{}, err
			}
		}
	}

	return model.NewExecutePluginObservation(groupByChainSelector(commitReports), messages).Encode()
}

func (p *Plugin) ValidateObservation(outctx ocr3types.OutcomeContext, query types.Query, ao types.AttributedObservation) error {
	// TODO: do "readers" need to be configured?
	//       for security, it doesn't matter. A merkle root is generated which must be consisted with the commit report.

	panic("implement me")
}

func (p *Plugin) ObservationQuorum(outctx ocr3types.OutcomeContext, query types.Query) (ocr3types.Quorum, error) {
	// TODO: should we use f+1 (or less) instead of 2f+1 because it is not needed for security?
	return ocr3types.QuorumFPlusOne, nil
}

func (p *Plugin) Outcome(outctx ocr3types.OutcomeContext, query types.Query, aos []types.AttributedObservation) (ocr3types.Outcome, error) {
	// aggregated list of observations?
	// TODO: whats the difference between this and the Report?
	//       just the seqNr it seems, attach that to the outcome to make a report?
	panic("implement me")
}

func (p *Plugin) Reports(seqNr uint64, outcome ocr3types.Outcome) ([]ocr3types.ReportWithInfo[[]byte], error) {
	panic("implement me")
}

func (p *Plugin) ShouldAcceptAttestedReport(ctx context.Context, u uint64, r ocr3types.ReportWithInfo[[]byte]) (bool, error) {
	panic("implement me")
}

func (p *Plugin) ShouldTransmitAcceptedReport(ctx context.Context, u uint64, r ocr3types.ReportWithInfo[[]byte]) (bool, error) {
	panic("implement me")
}

func (p *Plugin) Close() error {
	panic("implement me")
}

// Interface compatibility checks.
var _ ocr3types.ReportingPlugin[[]byte] = &Plugin{}
