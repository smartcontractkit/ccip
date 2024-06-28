package execute

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"sort"
	"sync/atomic"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/ccipocr3/internal/libs/slicelib"
)

// maxReportSizeBytes that should be returned as an execution report payload.
const maxReportSizeBytes = 250_000

// Plugin implements the main ocr3 plugin logic.
type Plugin struct {
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

// buildSingleChainReportMaxSize generates the largest report which fits into maxSizeBytes.
// See buildSingleChainReport for more details about how a report is built.
func buildSingleChainReportMaxSize(
	ctx context.Context,
	lggr logger.Logger,
	hasher cciptypes.MessageHasher,
	tokenDataReader TokenDataReader,
	encoder cciptypes.ExecutePluginCodec,
	report cciptypes.ExecutePluginCommitDataWithMessages,
	maxSizeBytes int,
) (cciptypes.ExecutePluginReportSingleChain, int, cciptypes.ExecutePluginCommitDataWithMessages, error) {
	finalReport, encodedSize, err :=
		buildSingleChainReport(ctx, lggr, hasher, tokenDataReader, encoder, report, 0)
	if err != nil {
		return cciptypes.ExecutePluginReportSingleChain{},
			0,
			cciptypes.ExecutePluginCommitDataWithMessages{},
			fmt.Errorf("unable to build a single chain report (max): %w", err)
	}

	// return fully executed report
	if encodedSize <= maxSizeBytes {
		report = markNewMessagesExecuted(finalReport, report)
		return finalReport, encodedSize, report, nil
	}

	var searchErr error
	idx := sort.Search(len(report.Messages), func(mid int) bool {
		if searchErr != nil {
			return false
		}
		finalReport2, encodedSize2, err :=
			buildSingleChainReport(ctx, lggr, hasher, tokenDataReader, encoder, report, mid)
		if searchErr != nil {
			searchErr = fmt.Errorf("unable to build a single chain report (messages %d): %w", mid, err)
		}

		if (encodedSize2) <= maxSizeBytes {
			// mid is a valid report size, try something bigger next iteration.
			finalReport = finalReport2
			encodedSize = encodedSize2
			return false // not full
		}
		return true // full
	})
	if searchErr != nil {
		return cciptypes.ExecutePluginReportSingleChain{}, 0, cciptypes.ExecutePluginCommitDataWithMessages{}, searchErr
	}

	// No messages fit into the report.
	if idx <= 0 {
		return cciptypes.ExecutePluginReportSingleChain{},
			0,
			cciptypes.ExecutePluginCommitDataWithMessages{},
			errNothingExecuted
	}

	report = markNewMessagesExecuted(finalReport, report)
	return finalReport, encodedSize, report, nil
}

// buildSingleChainReport converts the on-chain event data stored in cciptypes.ExecutePluginCommitDataWithMessages into
// the final on-chain report format.
//
// The hasher and encoding codec are provided as arguments to allow for chain-specific formats to be used.
//
// The maxMessages argument is used to limit the number of messages that are included in the report. If maxMessages is
// set to 0, all messages will be included. This allows the caller to create smaller reports if needed.
func buildSingleChainReport(
	ctx context.Context,
	lggr logger.Logger,
	hasher cciptypes.MessageHasher,
	tokenDataReader TokenDataReader,
	encoder cciptypes.ExecutePluginCodec,
	report cciptypes.ExecutePluginCommitDataWithMessages,
	maxMessages int,
) (cciptypes.ExecutePluginReportSingleChain, int, error) {
	// TODO: maxMessages selects messages in FIFO order which may not yield the optimal message size. One message with a
	//       maximum data size could push the report over a size limit even if several smaller messages could have fit.
	if maxMessages == 0 {
		maxMessages = len(report.Messages)
	}

	tree, err := constructMerkleTree(ctx, hasher, report)
	if err != nil {
		return cciptypes.ExecutePluginReportSingleChain{}, 0,
			fmt.Errorf("unable to construct merkle tree from messages: %w", err)
	}
	lggr.Debugw(
		"constructing merkle tree",
		"sourceChain", report.SourceChain,
		"treeLeaves", len(report.Messages))
	numMsgs := len(report.Messages)

	// Iterate sequence range and executed messages to select messages to execute.
	var toExecute []int
	var offchainTokenData [][][]byte
	var msgInRoot []cciptypes.CCIPMsg
	executedIdx := 0
	for i := 0; i < numMsgs && len(toExecute) <= maxMessages; i++ {
		seqNum := report.SequenceNumberRange.Start() + cciptypes.SeqNum(i)
		// Skip messages which are already executed
		if executedIdx < len(report.ExecutedMessages) && report.ExecutedMessages[executedIdx] == seqNum {
			executedIdx++
		} else {
			msg := report.Messages[i]
			tokenData, err := tokenDataReader.ReadTokenData(context.Background(), report.SourceChain, msg.SeqNum)
			if err != nil {
				// TODO: skip message instead of failing the whole thing.
				//       that might mean moving the token data reading out of the loop.
				lggr.Infow(
					"unable to read token data",
					"source-chain", report.SourceChain,
					"seq-num", msg.SeqNum,
					"error", err)
				return cciptypes.ExecutePluginReportSingleChain{}, 0, fmt.Errorf(
					"unable to read token data for message %d: %w", msg.SeqNum, err)
			}

			lggr.Debugw(
				"read token data",
				"source-chain", report.SourceChain,
				"seq-num", msg.SeqNum,
				"data", tokenData)
			offchainTokenData = append(offchainTokenData, tokenData)
			toExecute = append(toExecute, i)
			msgInRoot = append(msgInRoot, msg)
		}
	}

	lggr.Infow(
		"selected messages from commit report for execution",
		"sourceChain", report.SourceChain,
		"commitRoot", report.MerkleRoot.String(),
		"numMessages", numMsgs,
		"toExecute", len(toExecute))
	proof, err := tree.Prove(toExecute)
	if err != nil {
		return cciptypes.ExecutePluginReportSingleChain{}, 0,
			fmt.Errorf("unable to prove messages for report %s: %w", report.MerkleRoot.String(), err)
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
	encoded, err := encoder.Encode(
		ctx,
		cciptypes.ExecutePluginReport{
			ChainReports: []cciptypes.ExecutePluginReportSingleChain{finalReport},
		},
	)
	if err != nil {
		lggr.Errorw("unable to encode report", "err", err, "report", finalReport)
		return cciptypes.ExecutePluginReportSingleChain{}, 0, fmt.Errorf("unable to encode report: %w", err)
	}

	return finalReport, len(encoded), nil
}

// selectReport takes a list of reports in execution order and selects the first reports that fit within the
// maxReportSizeBytes. Individual messages in a commit report may be skipped for various reasons, for example if an
// out-of-order execution is detected or the message requires additional off-chain metadata which is not yet available.
// If there is not enough space in the final report, it may be partially executed by searching for a subset of messages
// which can fit in the final report.
func selectReport(
	ctx context.Context,
	lggr logger.Logger,
	hasher cciptypes.MessageHasher,
	encoder cciptypes.ExecutePluginCodec,
	tokenDataReader TokenDataReader,
	reports []cciptypes.ExecutePluginCommitDataWithMessages,
	maxReportSizeBytes int,
) ([]cciptypes.ExecutePluginReportSingleChain, []cciptypes.ExecutePluginCommitDataWithMessages, error) {
	// TODO: It may be desirable for this entire function to be an interface so that
	//       different selection algorithms can be used.

	// count number of fully executed reports so that they can be removed after iterating the reports.
	fullyExecuted := 0
	accumulatedSize := 0
	var finalReports []cciptypes.ExecutePluginReportSingleChain
	for reportIdx, report := range reports {
		execReport, encodedSize, updatedReport, err :=
			buildSingleChainReportMaxSize(ctx, lggr, hasher, tokenDataReader, encoder,
				report, maxReportSizeBytes-accumulatedSize)
		// No messages fit into the report, stop adding more reports.
		if errors.Is(err, errNothingExecuted) {
			break
		}
		if err != nil {
			return nil, nil, fmt.Errorf("unable to build single chain report: %w", err)
		}
		reports[reportIdx] = updatedReport
		accumulatedSize += encodedSize
		finalReports = append(finalReports, execReport)

		// partially executed report detected, stop adding more reports.
		// TODO: do not break if messages were intentionally skipped.
		if len(updatedReport.Messages) != len(updatedReport.ExecutedMessages) {
			break
		}
		fullyExecuted++
	}

	// Remove reports that are about to be executed.
	if fullyExecuted == len(reports) {
		reports = nil
	} else {
		reports = reports[fullyExecuted:]
	}

	lggr.Infow(
		"selected commit reports for execution report",
		"numReports", len(finalReports),
		"size", accumulatedSize,
		"incompleteReports", len(reports),
		"maxSize", maxReportSizeBytes)

	return finalReports, reports, nil
}

// Outcome collects the reports from the two phases and constructs the final outcome. Part of the outcome is a fully
// formed report that will be encoded for final transmission in the reporting phase.
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

	// TODO: this function should be pure, a context should not be needed.
	outcomeReports, commitReports, err :=
		selectReport(context.Background(), p.lggr, p.msgHasher, p.reportCodec, p.tokenDataReader,
			commitReports, maxReportSizeBytes)
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

	// TODO: this function should be pure, a context should not be needed.
	encoded, err := p.reportCodec.Encode(context.Background(), decodedOutcome.Report)
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
