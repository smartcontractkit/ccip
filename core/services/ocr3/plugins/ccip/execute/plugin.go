package commit

import (
	"context"

	//cache "github.com/smartcontractkit/ccipocr3/internal/copypaste/commit_roots_cache"
	"github.com/smartcontractkit/ccipocr3/internal/model"
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

	//commitRootsCache cache.CommitsRootsCache
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

type Observation struct {
	NodeID model.NodeID
	// slice of messages for each chain
	Msgs map[model.ChainSelector][]model.CCIPMsgBaseDetails

	// slice of (oldest?) reports from destination
	Reports [][]byte
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

	return model.NewExecutePluginObservation("", nil).Encode()
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
