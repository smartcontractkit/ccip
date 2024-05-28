package commit

import (
	"context"

	//cache "github.com/smartcontractkit/ccipocr3/internal/copypaste/commit_roots_cache"
	"github.com/smartcontractkit/ccipocr3/internal/model"
	"github.com/smartcontractkit/ccipocr3/internal/reader"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

// Interface compatibility checks.
var (
	_ ocr3types.ReportingPlugin[[]byte] = &Plugin{}
)

// StaticConfig contains configuration derived from the job spec that is shared across all instances of the plugin.
type StaticConfig struct {
}

// Plugin implements the main ocr3 plugin logic.
type Plugin struct {
	StaticConfig

	destChain model.ChainSelector

	reader reader.CCIP

	//commitRootsCache cache.CommitsRootsCache
	lastReportBlock uint64
}

func NewPlugin(config StaticConfig) *Plugin {
	return &Plugin{
		StaticConfig: config,
		//commitRootsCache: cache.NewCommitRootsCache(lggr, onchainConfig.PermissionLessExecutionThresholdSeconds, offchainConfig.RootSnoozeTime.Duration()),
	}
}

func (p *Plugin) Query(ctx context.Context, outctx ocr3types.OutcomeContext) (types.Query, error) {
	return types.Query{}, nil
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
func (p *Plugin) Observation(ctx context.Context, outctx ocr3types.OutcomeContext, query types.Query) (types.Observation, error) {
	previousOutcome, err := model.DecodeExecutePluginOutcome(outctx.PreviousOutcome)
	if err != nil {
		return types.Observation{}, err
	}

	// Phase 1: Gather commit reports from the destination chain and determine which messages are required to build a valid execution report.
	// TODO: filter out "cannot read p.destChain" errors? Or avoid calling it in the first place?
	commitReports, err := p.reader.ReportsFromBlockNum(ctx, p.destChain, p.lastReportBlock, 1000)
	if err != nil {
		return types.Observation{}, err
	}

	// Phase 2: Gather messages from the source chains and build the execution report.
	messages := make(map[model.ChainSelector][]model.CCIPMessage)
	if len(previousOutcome.Messages) == 0 {
		// No messages to execute.
		// This is expected after a cold start.
	} else {
		for selector, reports := range previousOutcome.NextCommits {
			for _, report := range reports {
				msgs, err := p.reader.MsgsBetweenSeqNums(ctx, []model.ChainSelector{selector}, report.SequenceNumberRange)
				if err != nil {
					return types.Observation{}, err
				}
				var convert []model.CCIPMessage
				for _, msg := range msgs {
					convert = append(convert, model.CCIPMessage{
						SequenceNumber: msg.SeqNum,
						Message:        msg.ID[:],
					})
				}
				messages[selector] = convert
			}
		}
	}

	return model.NewExecutePluginObservation(commitReports, messages).Encode()
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
