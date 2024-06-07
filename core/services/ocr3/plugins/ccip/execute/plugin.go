package execute

import (
	"context"
	"fmt"
	"slices"
	"sort"
	"sync/atomic"
	"time"

	"golang.org/x/crypto/sha3"

	"github.com/smartcontractkit/ccipocr3/execute/internal/validation"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// Plugin implements the main ocr3 plugin logic.
type Plugin struct {
	nodeID     commontypes.OracleID
	cfg        cciptypes.ExecutePluginConfig
	ccipReader cciptypes.CCIPReader

	//commitRootsCache cache.CommitsRootsCache
	lastReportTS *atomic.Int64
}

func NewPlugin(
	_ context.Context,
	nodeID commontypes.OracleID,
	cfg cciptypes.ExecutePluginConfig,
	ccipReader cciptypes.CCIPReader,
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

func getPendingExecutedReports(ctx context.Context, ccipReader cciptypes.CCIPReader, dest cciptypes.ChainSelector, ts time.Time) (cciptypes.ExecutePluginCommitObservations, time.Time, error) {
	latestReport := time.Time{}
	commitReports, err := ccipReader.CommitReportsGTETimestamp(ctx, dest, ts, 1000)
	if err != nil {
		return nil, time.Time{}, err
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

		var executedMessages []cciptypes.SeqNumRange
		for _, seqRange := range ranges {
			executedMessagesForRange, err2 := ccipReader.ExecutedMessageRanges(ctx, selector, dest, seqRange)
			if err2 != nil {
				return nil, time.Time{}, err2
			}
			executedMessages = append(executedMessages, executedMessagesForRange...)
		}

		// Remove fully executed reports.
		groupedCommits[selector], err = filterOutExecutedMessages(reports, executedMessages)
		if err != nil {
			return nil, time.Time{}, err
		}
	}

	// Put back in a single slice

	sort.Slice(commitReports, func(i, j int) bool {
		return commitReports[i].BlockNum < commitReports[j].BlockNum
	})
	if len(commitReports) > 0 {
		latestReport = commitReports[len(commitReports)-1].Timestamp
	}

	return groupedCommits, latestReport, nil
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
	previousOutcome, err := cciptypes.DecodeExecutePluginOutcome(outctx.PreviousOutcome)
	if err != nil {
		return types.Observation{}, err
	}

	// Phase 1: Gather commit reports from the destination chain and determine which messages are required to build a valid execution report.
	ownConfig := p.cfg.ObserverInfo[p.nodeID]
	var groupedCommits cciptypes.ExecutePluginCommitObservations
	if slices.Contains(ownConfig.Reads, p.cfg.DestChain) {
		var oldestReport time.Time
		groupedCommits, oldestReport, err = getPendingExecutedReports(ctx, p.ccipReader, p.cfg.DestChain, time.UnixMilli(p.lastReportTS.Load()))
		if err != nil {
			return types.Observation{}, err
		}
		// Update timestamp to the last report.
		p.lastReportTS.Store(oldestReport.UnixMilli())

		// TODO: truncate grouped commits to a maximum observation size.
		//       Cache everything which is not executed.
	}

	// Phase 2: Gather messages from the source chains and build the execution report.
	messages := make(cciptypes.ExecutePluginMessageObservations)
	if len(previousOutcome.Messages) == 0 {
		fmt.Println("TODO: No messages to execute. This is expected after a cold start.")
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
				msgs, err := p.ccipReader.MsgsBetweenSeqNums(ctx, selector, seqRange)
				if err != nil {
					return nil, err
				}
				for _, msg := range msgs {
					messages[selector][msg.SeqNum] = msg.ID
				}
			}
		}
	}

	// TODO: Fire off messages for an attestation check service.

	return cciptypes.NewExecutePluginObservation(groupedCommits, messages).Encode()
}

func (p *Plugin) ValidateObservation(outctx ocr3types.OutcomeContext, query types.Query, ao types.AttributedObservation) error {
	decodedObservation, err := cciptypes.DecodeExecutePluginObservation(ao.Observation)
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

// validatedCommitObservations merges all observations which reach the fChain threshold into a single result.
// Any observations, or subsets of observations, which do not reach the threshold are ignored.
func validatedCommitObservations(aos []decodedAttributedObservation, fChain map[cciptypes.ChainSelector]int) (cciptypes.ExecutePluginCommitObservations, error) {
	// TODO: validate and merge decoded observations into a single observation.

	// Merge commit reports.
	// Merge messages.
	// Ensure f_Chain observations for all.

	// Create a validator for each chain
	validators := make(map[cciptypes.ChainSelector]validation.Validator[cciptypes.ExecutePluginCommitData])
	idFunc := func(data cciptypes.ExecutePluginCommitData) [32]byte {
		return sha3.Sum256([]byte(fmt.Sprintf("%v", data)))
	}
	for selector, f := range fChain {
		validators[selector] = validation.NewValidator[cciptypes.ExecutePluginCommitData](2*f+1, idFunc)
	}

	// Need some sort of cache to store the reports, there could be invalid reports, so we need to keep a tally
	// of all versions to figure out which one is the most common and whether or not there are fChain observations.
	// TODO: this may be easier if executions were stored separately from the reports.
	for _, ao := range aos {
		for selector, commitReports := range ao.Observation.CommitReports {
			validator, ok := validators[selector]
			if !ok {
				return cciptypes.ExecutePluginCommitObservations{}, fmt.Errorf("no validator for chain %d", selector)
			}
			// Add reports
			for _, commitReport := range commitReports {
				if err := validator.AddReport(commitReport); err != nil {
					return cciptypes.ExecutePluginCommitObservations{}, err
				}
			}
		}
	}

	results := make(cciptypes.ExecutePluginCommitObservations)
	for selector, validator := range validators {
		var err error
		results[selector], err = validator.GetValidatedReports()
		if err != nil {
			return cciptypes.ExecutePluginCommitObservations{}, err
		}
	}

	return results, nil
}

func (p *Plugin) Outcome(outctx ocr3types.OutcomeContext, query types.Query, aos []types.AttributedObservation) (ocr3types.Outcome, error) {
	decodedObservations, err := decodeAttributedObservations(aos)
	if err != nil {
		return ocr3types.Outcome{}, err

	}
	if len(decodedObservations) < p.cfg.F {
		return ocr3types.Outcome{}, fmt.Errorf("below F threshold")
	}

	mergedCommitObservations, err := validatedCommitObservations(decodedObservations, p.cfg.FChain)
	if err != nil {
		return ocr3types.Outcome{}, err
	}

	// TODO: flatten results and sort by timestamp

	// TODO: validatedMessageObservations

	// add reports one by one into the outcome until

	fmt.Println(mergedCommitObservations, err)

	// Reports from previous outcome
	previousOutcome, err := cciptypes.DecodeExecutePluginOutcome(outctx.PreviousOutcome)
	if err != nil {
		return ocr3types.Outcome{}, err
	}
	for selector, reports := range previousOutcome.NextCommits {
		// TODO: if we have all of the messages for the previous requested reports, build the proof.
		fmt.Println(selector, reports)
	}

	// TODO: carry over reports which weren't executed by adding them to mergedCommitObservations.

	// TODO: do messages even go into the outcome??? I don't think so. Just the Observation.
	//       on the other hand, outcomes can be large since they aren't gossipped.

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
