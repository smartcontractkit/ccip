package execute

import (
	"context"
	"fmt"
	"slices"
	"sort"
	"sync/atomic"
	"time"

	"github.com/smartcontractkit/ccipocr3/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/merklemulti"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

// maxReportSize that should be returned as an execution report payload.
const maxReportSize = 250_000

// Plugin implements the main ocr3 plugin logic.
type Plugin struct {
	ctx             context.Context
	reportingCfg    ocr3types.ReportingPluginConfig
	cfg             cciptypes.ExecutePluginConfig
	ccipReader      cciptypes.CCIPReader
	reportCodec     cciptypes.ExecutePluginCodec
	msgHasher       cciptypes.MessageHasher
	tokenDataReader TokenDataReader

	lastReportTS *atomic.Int64

	lggr logger.Logger
}

func NewPlugin(
	ctx context.Context,
	reportingCfg ocr3types.ReportingPluginConfig,
	cfg cciptypes.ExecutePluginConfig,
	ccipReader cciptypes.CCIPReader,
	reportCodec cciptypes.ExecutePluginCodec,
	msgHasher cciptypes.MessageHasher,
	lggr logger.Logger,
) *Plugin {
	lastReportTS := &atomic.Int64{}
	lastReportTS.Store(time.Now().Add(-cfg.MessageVisibilityInterval).UnixMilli())

	// TODO: initialize tokenDataReader.

	return &Plugin{
		ctx:          ctx,
		reportingCfg: reportingCfg,
		cfg:          cfg,
		ccipReader:   ccipReader,
		reportCodec:  reportCodec,
		msgHasher:    msgHasher,
		lastReportTS: lastReportTS,
		lggr:         lggr,
	}
}

func (p *Plugin) Query(ctx context.Context, outctx ocr3types.OutcomeContext) (types.Query, error) {
	return types.Query{}, nil
}

func getPendingExecutedReports(
	ctx context.Context, ccipReader cciptypes.CCIPReader, dest cciptypes.ChainSelector, ts time.Time,
) (cciptypes.ExecutePluginCommitObservations, time.Time, error) {
	latestReportTS := time.Time{}
	commitReports, err := ccipReader.CommitReportsGTETimestamp(ctx, dest, ts, 1000)
	if err != nil {
		return nil, time.Time{}, err
	}
	// TODO: this could be more efficient. reports is also traversed in 'filterOutExecutedMessages' function.
	for _, report := range commitReports {
		if report.Timestamp.After(latestReportTS) {
			latestReportTS = report.Timestamp
		}
	}

	// TODO: this could be more efficient. commitReports is also traversed in 'groupByChainSelector'.
	for _, report := range commitReports {
		if report.Timestamp.After(latestReportTS) {
			latestReportTS = report.Timestamp
		}
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

	return groupedCommits, latestReportTS, nil
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
func (p *Plugin) Observation(
	ctx context.Context, outctx ocr3types.OutcomeContext, _ types.Query,
) (types.Observation, error) {
	previousOutcome, err := cciptypes.DecodeExecutePluginOutcome(outctx.PreviousOutcome)
	if err != nil {
		return types.Observation{}, err
	}

	// Phase 1: Gather commit reports from the destination chain and determine which messages are required to build a
	//          valid execution report.
	ownConfig := p.cfg.ObserverInfo[p.reportingCfg.OracleID]
	var groupedCommits cciptypes.ExecutePluginCommitObservations
	if slices.Contains(ownConfig.Reads, p.cfg.DestChain) {
		var latestReportTS time.Time
		groupedCommits, latestReportTS, err =
			getPendingExecutedReports(ctx, p.ccipReader, p.cfg.DestChain, time.UnixMilli(p.lastReportTS.Load()))
		if err != nil {
			return types.Observation{}, err
		}
		// Update timestamp to the last report.
		p.lastReportTS.Store(latestReportTS.UnixMilli())

		// TODO: truncate grouped commits to a maximum observation size.
		//       Cache everything which is not executed.
	}

	// Phase 2: Gather messages from the source chains and build the execution report.
	messages := make(cciptypes.ExecutePluginMessageObservations)
	if len(previousOutcome.PendingCommitReports) == 0 {
		fmt.Println("TODO: No reports to execute. This is expected after a cold start.")
		// No reports to execute.
		// This is expected after a cold start.
	} else {
		commitReportCache := make(map[cciptypes.ChainSelector][]cciptypes.ExecutePluginCommitDataWithMessages)
		for _, report := range previousOutcome.PendingCommitReports {
			commitReportCache[report.SourceChain] = append(commitReportCache[report.SourceChain], report)
		}

		for selector, reports := range commitReportCache {
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
					messages[selector][msg.SeqNum] = msg
				}
			}
		}
	}

	// TODO: Fire off messages for an attestation check service.

	return cciptypes.NewExecutePluginObservation(groupedCommits, messages).Encode()
}

func (p *Plugin) ValidateObservation(
	outctx ocr3types.OutcomeContext, query types.Query, ao types.AttributedObservation,
) error {
	decodedObservation, err := cciptypes.DecodeExecutePluginObservation(ao.Observation)
	if err != nil {
		return fmt.Errorf("decode observation: %w", err)
	}

	err = validateObserverReadingEligibility(p.reportingCfg.OracleID, p.cfg.ObserverInfo, decodedObservation.Messages)
	if err != nil {
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

// TokenDataReader is an interface for reading extra token data from an async process.
// TODO: Build a token data reading process.
type TokenDataReader interface {
	ReadTokenData(ctx context.Context, srcChain cciptypes.ChainSelector, num cciptypes.SeqNum) ([][]byte, error)
}

// buildSingleChainReport constructs a single chain report from a commit report.
func buildSingleChainReport(ctx context.Context, lggr logger.Logger, hasher cciptypes.MessageHasher, codec cciptypes.ExecutePluginCodec, tokenDataReader TokenDataReader, report cciptypes.ExecutePluginCommitDataWithMessages, maxReportSize, maxMessages int) (cciptypes.ExecutePluginReportSingleChain /* encoded size */, int, error) {
	if maxMessages == 0 {
		maxMessages = len(report.Messages)
	}

	numMsg := int(report.SequenceNumberRange.End() - report.SequenceNumberRange.Start() + 1)
	if numMsg != len(report.Messages) {
		return cciptypes.ExecutePluginReportSingleChain{}, 0, fmt.Errorf("malformed report %s, unexpected number of messages: expected %d , got %d", report.MerkleRoot.String(), numMsg, len(report.Messages))
	}

	treeLeaves := make([][32]byte, 0)
	for _, msg := range report.Messages {
		if !report.SequenceNumberRange.Contains(msg.SeqNum) {
			return cciptypes.ExecutePluginReportSingleChain{}, 0, fmt.Errorf("malformed report %s, message with sequence number %d outside of report range %s", report.MerkleRoot.String(), msg.SeqNum, report.SequenceNumberRange)
		}
		// TODO: pass in a hasher to construct the chain specific merkle tree.
		leaf, err := hasher.Hash(ctx, msg)
		if err != nil {
			return cciptypes.ExecutePluginReportSingleChain{}, 0, fmt.Errorf("unable to hash message (%d, %d): %w", msg.SourceChain, msg.SeqNum, err)
		}
		treeLeaves = append(treeLeaves, leaf)
	}

	lggr.Debugw("constructing merkle tree", "sourceChain", report.SourceChain, "treeLeaves", len(treeLeaves))
	tree, err := merklemulti.NewTree(hashutil.NewKeccak(), treeLeaves)
	if err != nil {
		return cciptypes.ExecutePluginReportSingleChain{}, 0, fmt.Errorf("unable to constructing merkle tree from messages: %w", err)
	}

	// Iterate sequence range and executed messages to select messages to execute.
	var toExecute []int
	var offchainTokenData [][][]byte
	var msgInRoot []cciptypes.CCIPMsg
	executedIdx := 0
	for i := 0; i < numMsg && len(toExecute) <= maxMessages; i++ {
		seqNum := report.SequenceNumberRange.Start() + cciptypes.SeqNum(i)
		// Skip messages which are already executed
		if executedIdx < len(report.ExecutedMessages) && report.ExecutedMessages[executedIdx] == seqNum {
			executedIdx++
		} else {
			msg := report.Messages[i]
			tokenData, err := tokenDataReader.ReadTokenData(context.Background(), report.SourceChain, msg.SeqNum)
			if err != nil {
				lggr.Info("unable to read token data", "source-chain", report.SourceChain, "seq-num", msg.SeqNum, "error", err)
				offchainTokenData = append(offchainTokenData, nil)
			} else {
				lggr.Debugw("read token data", "source-chain", report.SourceChain, "seq-num", msg.SeqNum, "data", tokenData)
				offchainTokenData = append(offchainTokenData, tokenData)
			}
			toExecute = append(toExecute, i)
			msgInRoot = append(msgInRoot, msg)
		}
	}

	lggr.Infow("selected messages from commit report for execution", "sourceChain", report.SourceChain, "commitRoot", report.MerkleRoot.String(), "numMessages", len(toExecute), "totalMessages", numMsg, "toExecuteu", len(toExecute))
	proof, err := tree.Prove(toExecute)
	if err != nil {
		return cciptypes.ExecutePluginReportSingleChain{}, 0, fmt.Errorf("unable to prove messages for report %s: %w", report.MerkleRoot.String(), err)
	}

	var proofsCast []cciptypes.Bytes32
	for _, p := range proof.Hashes {
		proofsCast = append(proofsCast, p)
	}

	finalReport := cciptypes.ExecutePluginReportSingleChain{
		SourceChainSelector: report.SourceChain,
		Messages:            msgInRoot,
		OffchainTokenData:   offchainTokenData,
		Proofs:              proofsCast,
		ProofFlagBits:       cciptypes.BigInt{Int: slicelib.BoolsToBitFlags(proof.SourceFlags)},
	}

	// Note: ExecutePluginReport is a strict array of data, so wrapping the final report
	//       does not add any additional overhead to the size being computed here.

	// Compute the size of the encoded report.
	encoded, err := codec.Encode(ctx, cciptypes.ExecutePluginReport{ChainReports: []cciptypes.ExecutePluginReportSingleChain{finalReport}})
	if err != nil {
		lggr.Errorw("unable to encode report", "err", err, "report", finalReport)
		return cciptypes.ExecutePluginReportSingleChain{}, 0, fmt.Errorf("unable to encode report: %w", err)
	}

	return finalReport, len(encoded), nil
}

// selectReport takes an ordered list of reports and selects the first reports that fit within the maxReportSize.
func selectReport(ctx context.Context, lggr logger.Logger, hasher cciptypes.MessageHasher, codec cciptypes.ExecutePluginCodec, tokenDataReader TokenDataReader, reports []cciptypes.ExecutePluginCommitDataWithMessages, maxReportSize int) ([]cciptypes.ExecutePluginReportSingleChain, []cciptypes.ExecutePluginCommitDataWithMessages, error) {
	// TODO: It may be desirable for this entire function to be an interface so that
	//       different selection algorithms can be used.

	size := 0
	fullyExecuted := 0
	partialReport := false
	var finalReports []cciptypes.ExecutePluginReportSingleChain
	for reportIdx, report := range reports {
		finalReport, encodedSize, err := buildSingleChainReport(ctx, lggr, hasher, codec, tokenDataReader, report, maxReportSize, 0)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to build a single chain report (max): %w", err)
		}

		// The full report is too large, binary search for best report size
		if (size + encodedSize) >= maxReportSize {
			partialReport = true
			low := 1
			high := len(report.Messages) - 1
			for low <= high {
				mid := low + ((high - low) / 2)

				finalReport2, encodedSize2, err2 := buildSingleChainReport(ctx, lggr, hasher, codec, tokenDataReader, report, maxReportSize, mid)
				if err2 != nil {
					return nil, nil, fmt.Errorf("unable to build a single chain report (messages %d): %w", mid, err2)
				}

				if (size + encodedSize2) <= maxReportSize {
					// mid is a valid report size, try something bigger next iteration.
					finalReport = finalReport2
					encodedSize = encodedSize2
					low = mid + 1
				} else {
					// mid is invalid, try something smaller next iteration.
					high = mid - 1
				}
			}

			// Mark new messages executed.
			for i := 0; i < len(finalReport.Messages); i++ {
				reports[reportIdx].ExecutedMessages = append(reports[reportIdx].ExecutedMessages, report.Messages[i].SeqNum-report.SequenceNumberRange.Start())
			}
			sort.Slice(report.ExecutedMessages, func(i, j int) bool { return report.ExecutedMessages[i] < report.ExecutedMessages[j] })
		} else {
			fullyExecuted++
		}

		size += encodedSize
		finalReports = append(finalReports, finalReport)
		if partialReport {
			break
		}
	}

	// Remove reports that are about to be executed.
	if fullyExecuted == len(reports) {
		reports = nil
	} else {
		reports = reports[fullyExecuted:]
	}

	lggr.Infow("selected commit reports for execution report", "numReports", len(finalReports), "size", size, "incompleteReports", len(reports), "maxSize", maxReportSize)

	return finalReports, reports, nil
}

func (p *Plugin) Outcome(
	outctx ocr3types.OutcomeContext, query types.Query, aos []types.AttributedObservation,
) (ocr3types.Outcome, error) {
	decodedObservations, err := decodeAttributedObservations(aos)
	if err != nil {
		return ocr3types.Outcome{}, err

	}
	if len(decodedObservations) < p.reportingCfg.F {
		return ocr3types.Outcome{}, fmt.Errorf("below F threshold")
	}

	mergedCommitObservations, err := mergeCommitObservations(decodedObservations, p.cfg.FChain)
	if err != nil {
		return ocr3types.Outcome{}, err
	}

	mergedMessageObservations, err := mergeMessageObservations(decodedObservations, p.cfg.FChain)
	if err != nil {
		return ocr3types.Outcome{}, err
	}

	observation := cciptypes.NewExecutePluginObservation(
		mergedCommitObservations,
		mergedMessageObservations)

	// flatten commit reports and sort by timestamp.
	var commitReports []cciptypes.ExecutePluginCommitDataWithMessages
	for _, report := range observation.CommitReports {
		commitReports = append(commitReports, report...)
	}
	sort.Slice(commitReports, func(i, j int) bool {
		return commitReports[i].Timestamp.Before(commitReports[j].Timestamp)
	})

	// add messages to their commitReports.
	for _, report := range commitReports {
		report.Messages = nil
		for i := report.SequenceNumberRange.Start(); i <= report.SequenceNumberRange.End(); i++ {
			if msg, ok := observation.Messages[report.SourceChain][i]; ok {
				report.Messages = append(report.Messages, msg)
			}
		}
	}

	outcomeReports, commitReports, err := selectReport(p.ctx, p.lggr, p.msgHasher, p.reportCodec, nil, commitReports, maxReportSize)
	if err != nil {
		return ocr3types.Outcome{}, fmt.Errorf("unable to extract proofs: %w", err)
	}

	execReport := cciptypes.ExecutePluginReport{
		ChainReports: outcomeReports,
	}

	return cciptypes.NewExecutePluginOutcome(commitReports, execReport).Encode()
}

func (p *Plugin) Reports(seqNr uint64, outcome ocr3types.Outcome) ([]ocr3types.ReportWithInfo[[]byte], error) {
	decodedOutcome, err := cciptypes.DecodeExecutePluginOutcome(outcome)
	if err != nil {
		return nil, err
	}

	encoded, err := p.reportCodec.Encode(p.ctx, decodedOutcome.Report)
	if err != nil {
		return nil, err
	}

	report := []ocr3types.ReportWithInfo[[]byte]{{
		Report: encoded,
		Info:   nil,
	}}

	return report, nil
}

func (p *Plugin) ShouldAcceptAttestedReport(
	ctx context.Context, u uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	panic("implement me")
}

func (p *Plugin) ShouldTransmitAcceptedReport(
	ctx context.Context, u uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	panic("implement me")
}

func (p *Plugin) Close() error {
	panic("implement me")
}

// Interface compatibility checks.
var _ ocr3types.ReportingPlugin[[]byte] = &Plugin{}
