package commit

import (
	"context"
	"fmt"
	"sync/atomic"
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
	lastReportTS *atomic.Int64
}

func NewPlugin(
	_ context.Context,
	nodeID commontypes.OracleID,
	cfg model.ExecutePluginConfig,
	ccipReader reader.CCIP,
) *Plugin {
	lastReportTS := &atomic.Int64{}
	lastReportTS.Store(time.Now().Add(-cfg.MessageVisibilityInterval).UnixMilli())

	return &Plugin{
		nodeID:       nodeID,
		cfg:          cfg,
		ccipReader:   ccipReader,
		lastReportTS: lastReportTS,
	}
}

func (p *Plugin) Query(ctx context.Context, outctx ocr3types.OutcomeContext) (types.Query, error) {
	return types.Query{}, nil
}

func getPendingExecutedReports(ctx context.Context, ccipReader reader.CCIP, dest model.ChainSelector, ts time.Time) (model.ExecutePluginCommitObservations, time.Time, error) {
	// TODO: filter out "cannot read p.destChain" errors? Or avoid calling it in the first place?
	commitReports, err := ccipReader.CommitReportsGTETimestamp(ctx, dest, ts, 1000)
	if err != nil {
		return nil, time.Time{}, err
	}
	if len(commitReports) > 0 {
		//lastReport := commitReports[len(commitReports)-1]
		//p.lastReportTS = lastReport.
		// TODO: Need a way to get a timestamp of the report.
	}

	groupedCommits := groupByChainSelector(commitReports)

	// Remove fully executed reports.
	for selector, reports := range groupedCommits {
		if len(reports) == 0 {
			continue
		}

		ranges, err := computeRanges(reports)
		if err != nil {
			return nil, time.Time{}, err
		}

		var executedMessages []model.SeqNumRange
		for _, seqRange := range ranges {
			executedMessagesForRange, err := ccipReader.ExecutedMessageRanges(ctx, selector, dest, seqRange)
			if err != nil {
				return nil, time.Time{}, err
			}
			executedMessages = append(executedMessages, executedMessagesForRange...)
		}

		// Remove fully executed reports.
		groupedCommits[selector], err = filterOutFullyExecutedMessages(reports, executedMessages)
		if err != nil {
			return nil, time.Time{}, err
		}
	}

	return groupedCommits, time.Time{}, nil
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
	groupedCommits, _, err := getPendingExecutedReports(ctx, p.ccipReader, p.cfg.DestChain, time.UnixMilli(p.lastReportTS.Load()))
	// TODO: Need a way to get a timestamp of the report.

	// Phase 2: Gather messages from the source chains and build the execution report.
	messages := make(model.ExecutePluginMessageObservations)
	if len(previousOutcome.Messages) == 0 {
		// No messages to execute.
		// This is expected after a cold start.
	} else {
		for selector, reports := range previousOutcome.NextCommits {
			if len(reports) == 0 {
				continue
			}

			ranges, err := computeRanges(reports)
			if err != nil {
				return types.Observation{}, err
			}

			// Read messages for each range.
			for _, seqRange := range ranges {
				msgs, err := p.ccipReader.MsgsBetweenSeqNums(ctx, []model.ChainSelector{selector}, seqRange)
				if err != nil {
					return nil, err
				}
				for _, msg := range msgs {
					messages[selector][msg.SeqNum] = msg.ID
				}

				if err != nil {
					return types.Observation{}, err
				}
			}
		}
	}

	// TODO: Fire off messages for an attestation check service.

	return model.NewExecutePluginObservation(groupedCommits, messages).Encode()
}

func (p *Plugin) ValidateObservation(outctx ocr3types.OutcomeContext, query types.Query, ao types.AttributedObservation) error {
	decodedObservation, err := model.DecodeExecutePluginObservation(ao.Observation)
	if err != nil {
		return fmt.Errorf("decode observation: %w", err)
	}

	if err := validateObserverReadingEligibility(p.nodeID, p.cfg.ObserverInfo, decodedObservation.Messages); err != nil {
		return fmt.Errorf("validate observer reading eligibility: %w", err)
	}

	if err := validateObservedSequenceNumbers(decodedObservation.CommitReports); err != nil {
		return fmt.Errorf("validate observed sequence numbers: %w", err)
	}

	return nil
}

func (p *Plugin) ObservationQuorum(outctx ocr3types.OutcomeContext, query types.Query) (ocr3types.Quorum, error) {
	// TODO: should we use f+1 (or less) instead of 2f+1 because it is not needed for security?
	return ocr3types.QuorumFPlusOne, nil
}

func (p *Plugin) Outcome(outctx ocr3types.OutcomeContext, query types.Query, aos []types.AttributedObservation) (ocr3types.Outcome, error) {
	// TODO: do we care about f_chain here? I believe only commit is needs true consensus.
	//       if we do, it would mainly be to prevent bad participants from invalidating the proofs with bad data.
	// Aggregate messages from the current observations
	aggregatedMessages := make(map[model.ChainSelector]map[model.SeqNum]model.Bytes32)
	for _, ao := range aos {
		obs, err := model.DecodeExecutePluginObservation(ao.Observation)
		if err != nil {
			return ocr3types.Outcome{}, err
		}

		for selector, messages := range obs.Messages {
			for seqNr, message := range messages {
				aggregatedMessages[selector][seqNr] = message
			}
		}
	}

	// Reports from previous outcome
	// TODO: Build the proof
	/*
		previousOutcome, err := model.DecodeExecutePluginOutcome(outctx.PreviousOutcome)
		if err != nil {
			return ocr3types.Outcome{}, err
		}
		for selector, report := range previousOutcome.NextCommits {
			// if we have all of the messages, build the proof.
		}
	*/

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
