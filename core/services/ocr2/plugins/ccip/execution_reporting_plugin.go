package ccip

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/onramp"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/merklemulti"
)

const (
	ExecutionMaxInflightTimeSeconds = 180
	// Note user research is required for setting (MaxPayloadLength, MaxTokensPerMessage).
	// TODO: If we really want this to be constant and not dynamic, then we need to wait
	// until we have gas limits per message and ensure the block gas limit constraint is respected
	// as well as the tx size limit.
	MaxNumMessagesInExecutionReport = 70
	MaxPayloadLength                = 1000
	MaxTokensPerMessage             = 5
	// NOTE: If execution report format changes, this has to change.
	// See makeExecutionReportArgs. Note for each dynamic type, there's a offset + length word.
	// We explicitly do not include struct packing here as its an upper bound.
	MaxMessageLength = 32 + // len of message struct
		32*6 + // sourceChainId, seqNum, sender, destChainId, executor, receiver
		32*2 + // length of payload struct
		32*2 + // len, offset for data
		MaxPayloadLength +
		32*2 + // len, offset for tokens
		32*2 + // len, offset for amounts
		MaxTokensPerMessage*(2*32) // per token, token and amount
	MaxExecutionReportLength = 32 + // len of report struct
		32*2 + // len, offset for messages
		MaxMessageLength*MaxNumMessagesInExecutionReport + // messages
		32*2 + // len, offset proofs
		32 + // proof, only one in the case of all messages included
		32 // proofFlagBits
)

var (
	_                types.ReportingPluginFactory = &ExecutionReportingPluginFactory{}
	_                types.ReportingPlugin        = &ExecutionReportingPlugin{}
	ErrOffRampIsDown                              = errors.New("offramp is down")
)

func EncodeExecutionReport(msgs []Message, proofs [][32]byte, proofSourceFlags []bool) (types.Report, error) {
	report, err := makeExecutionReportArgs().PackValues([]interface{}{ExecutionReport{
		Messages:      msgs,
		Proofs:        proofs,
		ProofFlagBits: ProofFlagsToBits(proofSourceFlags),
	}})
	if err != nil {
		return nil, err
	}
	return report, nil
}

func DecodeExecutionReport(report types.Report) (*ExecutionReport, error) {
	unpacked, err := makeExecutionReportArgs().Unpack(report)
	if err != nil {
		return nil, err
	}
	if len(unpacked) == 0 {
		return nil, errors.New("assumptionViolation: expected at least one element")
	}

	// Must be anonymous struct here
	erStruct, ok := unpacked[0].(struct {
		Messages []struct {
			SourceChainId  *big.Int       `json:"sourceChainId"`
			SequenceNumber uint64         `json:"sequenceNumber"`
			Sender         common.Address `json:"sender"`
			Payload        struct {
				Tokens             []common.Address `json:"tokens"`
				Amounts            []*big.Int       `json:"amounts"`
				DestinationChainId *big.Int         `json:"destinationChainId"`
				Receiver           common.Address   `json:"receiver"`
				Executor           common.Address   `json:"executor"`
				Data               []uint8          `json:"data"`
			} `json:"payload"`
		} `json:"Messages"`
		Proofs        [][32]uint8 `json:"Proofs"`
		ProofFlagBits *big.Int    `json:"ProofFlagBits"`
	})
	if !ok {
		return nil, fmt.Errorf("got %T", unpacked[0])
	}
	if len(erStruct.Messages) == 0 {
		return nil, errors.New("assumptionViolation: expected at least one element")
	}
	var er ExecutionReport
	for _, msg := range erStruct.Messages {
		er.Messages = append(er.Messages, msg)
	}
	for _, proof := range erStruct.Proofs {
		er.Proofs = append(er.Proofs, proof)
	}
	er.ProofFlagBits = erStruct.ProofFlagBits
	return &er, nil
}

type ExecutionReportingPluginFactory struct {
	lggr         logger.Logger
	source, dest logpoller.LogPoller
	executor     common.Address
	onRamp       *onramp.OnRamp
	offRamp      *offramp.OffRamp
}

func NewExecutionReportingPluginFactory(
	lggr logger.Logger,
	onRamp *onramp.OnRamp,
	offRamp *offramp.OffRamp,
	source, dest logpoller.LogPoller,
	executor common.Address,
) types.ReportingPluginFactory {
	return &ExecutionReportingPluginFactory{lggr: lggr, onRamp: onRamp, offRamp: offRamp, executor: executor, source: source, dest: dest}
}

func (rf *ExecutionReportingPluginFactory) NewReportingPlugin(config types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	offchainConfig, err := Decode(config.OffchainConfig)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	return &ExecutionReportingPlugin{
			lggr:           rf.lggr.Named("ExecutionReportingPlugin"),
			F:              config.F,
			executor:       rf.executor,
			onRamp:         rf.onRamp,
			offRamp:        rf.offRamp,
			source:         rf.source,
			dest:           rf.dest,
			offchainConfig: offchainConfig,
		}, types.ReportingPluginInfo{
			Name:          "CCIPExecution",
			UniqueReports: true,
			Limits: types.ReportingPluginLimits{
				MaxQueryLength:       0,
				MaxObservationLength: MaxObservationLength,
				MaxReportLength:      MaxExecutionReportLength,
			},
		}, nil
}

type ExecutionReportingPlugin struct {
	lggr         logger.Logger
	F            int
	executor     common.Address
	onRamp       *onramp.OnRamp
	offRamp      *offramp.OffRamp
	source, dest logpoller.LogPoller
	// We need to synchronize access to the inflight structure
	// as reporting plugin methods may be called from separate goroutines,
	// e.g. reporting vs transmission protocol.
	inFlightMu     sync.RWMutex
	inFlight       []InflightExecutionReport
	offchainConfig OffchainConfig
}

type InflightExecutionReport struct {
	createdAt time.Time
	report    ExecutionReport
}

func (r *ExecutionReportingPlugin) Query(ctx context.Context, timestamp types.ReportTimestamp) (types.Query, error) {
	// We don't use a query for this reporting plugin, so we can just leave it empty here
	return types.Query{}, nil
}

// getRelayedReports returns them in sorted order.
func (r *ExecutionReportingPlugin) getRelayedReports(min, max uint64) ([]offramp.OffRampReportAccepted, error) {
	// Get all reports where minSeqNum is >= min as a lower bound.
	reportLogs, err := r.dest.LogsDataWordGreaterThan(ReportAccepted, r.offRamp.Address(), 1, EvmWord(min), 1)
	if err != nil {
		return nil, err
	}
	var reports []offramp.OffRampReportAccepted
	for _, reportLog := range reportLogs {
		report, err := r.offRamp.ParseReportAccepted(gethtypes.Log{Data: reportLog.Data, Topics: reportLog.GetTopics()})
		if err != nil {
			return nil, err
		}
		if report.Report.MinSequenceNumber >= min && report.Report.MaxSequenceNumber <= max {
			reports = append(reports, *report)
		}
	}
	return reports, nil
}

func (r *ExecutionReportingPlugin) inflightSeqNums() map[uint64]struct{} {
	r.inFlightMu.RLock()
	defer r.inFlightMu.RUnlock()
	inFlightSeqNums := make(map[uint64]struct{})
	for _, report := range r.inFlight {
		for _, msg := range report.report.Messages {
			inFlightSeqNums[msg.SequenceNumber] = struct{}{}
		}
	}
	return inFlightSeqNums
}

func (r *ExecutionReportingPlugin) getExecutedMessages() (map[uint64]struct{}, error) {
	blk, err := r.source.LatestBlock()
	if err != nil {
		return nil, err
	}
	// TODO: This scans all logs in the history of the offramp.
	// To optimize, we only need to scan finalized blocks once and remember the set of unexecuted.
	executedLogs, err := r.dest.Logs(1, blk, CrossChainMessageExecuted, r.offRamp.Address())
	if err != nil {
		return nil, err
	}
	var executedMp = make(map[uint64]struct{})
	for _, executedLog := range executedLogs {
		e, err := r.offRamp.ParseCrossChainMessageExecuted(gethtypes.Log{Data: executedLog.Data, Topics: executedLog.GetTopics()})
		if err != nil {
			return nil, err
		}
		executedMp[e.SequenceNumber] = struct{}{}
	}
	return executedMp, nil
}

func (r *ExecutionReportingPlugin) Observation(ctx context.Context, timestamp types.ReportTimestamp, query types.Query) (types.Observation, error) {
	lggr := r.lggr.Named("Observation")
	if isOffRampDownNow(r.lggr, r.offRamp) {
		return nil, ErrOffRampIsDown
	}
	rep, err := r.offRamp.GetLastReport(nil)
	if err != nil {
		return nil, err
	}
	// Find the set of unexecuted seq nums. i.e. ones which have a report accepted but do not have a cross chain executed
	// and are pinned to us as an executor.
	executedMp, err := r.getExecutedMessages()
	if err != nil {
		return nil, err
	}
	inFlightExecutions := r.inflightSeqNums()
	var executable []uint64
	for i := uint64(1); i <= rep.MaxSequenceNumber; i++ {
		_, executed := executedMp[i]
		_, inflight := inFlightExecutions[i]
		lggr.Debugw("Seq num", "num", i, "executed", executed, "inflight", inflight)
		if !executed && !inflight {
			executable = append(executable, i)
		}
	}
	lggr.Infof("Executable messages %v", executable)
	if len(executable) == 0 {
		return []byte{}, nil
	}
	return Observation{
		MinSeqNum: executable[0],
		MaxSeqNum: executable[len(executable)-1],
	}.Marshal()
}

func (r *ExecutionReportingPlugin) getMessagesInRangeWithExecutor(min, max uint64, executor common.Address) ([]onramp.OnRampCrossChainSendRequested, error) {
	msgs, err := r.source.LogsDataWordRange(CrossChainSendRequested, r.onRamp.Address(), 2, EvmWord(min), EvmWord(max), int(r.offchainConfig.SourceIncomingConfirmations))
	if err != nil {
		return nil, err
	}
	var reqs []onramp.OnRampCrossChainSendRequested
	for _, msg := range msgs {
		req, err := r.onRamp.ParseCrossChainSendRequested(gethtypes.Log{Data: msg.Data, Topics: msg.GetTopics()})
		if err != nil {
			return nil, err
		}
		if executor == [20]byte{} || req.Message.Payload.Executor == executor {
			reqs = append(reqs, *req)
		}
	}
	return reqs, nil
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

// Assumes non-empty report. Messages to execute can be span more than one report, but are assumed to be in order of increasing
// sequence number.
func (r *ExecutionReportingPlugin) buildReport(lggr logger.Logger, report offramp.OffRampReportAccepted, msgsToExecute []onramp.OnRampCrossChainSendRequested) ([]byte, error) {
	allMsgs, err2 := r.getMessagesInRangeWithExecutor(report.Report.MinSequenceNumber, report.Report.MaxSequenceNumber, [20]byte{})
	if err2 != nil {
		return nil, err2
	}
	if len(allMsgs) != int(report.Report.MaxSequenceNumber-report.Report.MinSequenceNumber+1) {
		return nil, errors.Errorf("do not have all messages, have %d want %d", len(allMsgs), int(report.Report.MaxSequenceNumber-report.Report.MinSequenceNumber+1))
	}
	mctx := merklemulti.NewKeccakCtx()
	var leaves [][32]byte
	for _, msg := range allMsgs {
		leaves = append(leaves, mctx.HashLeaf(msg.Raw.Data))
	}
	var messages []Message
	var indices []int
	tree := merklemulti.NewTree(mctx, leaves)
	for i := int64(0); i < int64(len(msgsToExecute)); i++ {
		if msgsToExecute[i].Message.SequenceNumber > report.Report.MaxSequenceNumber {
			// Again we only execute one report at a time.
			break
		}
		index := msgsToExecute[i].Message.SequenceNumber - report.Report.MinSequenceNumber
		if index < 0 {
			return nil, errors.New("unexpected invalid index")
		}
		indices = append(indices, int(index))
		messages = append(messages, Message{
			SequenceNumber: msgsToExecute[i].Message.SequenceNumber,
			SourceChainId:  msgsToExecute[i].Message.SourceChainId,
			Sender:         msgsToExecute[i].Message.Sender,
			Payload: struct {
				Tokens             []common.Address `json:"tokens"`
				Amounts            []*big.Int       `json:"amounts"`
				DestinationChainId *big.Int         `json:"destinationChainId"`
				Receiver           common.Address   `json:"receiver"`
				Executor           common.Address   `json:"executor"`
				Data               []uint8          `json:"data"`
			}{
				Tokens:             msgsToExecute[i].Message.Payload.Tokens,
				Amounts:            msgsToExecute[i].Message.Payload.Amounts,
				DestinationChainId: msgsToExecute[i].Message.Payload.DestinationChainId,
				Receiver:           msgsToExecute[i].Message.Payload.Receiver,
				Executor:           msgsToExecute[i].Message.Payload.Executor,
				Data:               msgsToExecute[i].Message.Payload.Data,
			},
		})
	}
	if len(messages) == 0 {
		return nil, errors.New("no executed messages created")
	}
	proof := tree.Prove(indices)
	er, err := EncodeExecutionReport(messages, proof.Hashes, proof.SourceFlags)
	if err != nil {
		return nil, err
	}
	return er, nil
}

func (r *ExecutionReportingPlugin) Report(ctx context.Context, timestamp types.ReportTimestamp, query types.Query, observations []types.AttributedObservation) (bool, types.Report, error) {
	lggr := r.lggr.Named("Report")
	if isOffRampDownNow(lggr, r.offRamp) {
		return false, nil, ErrOffRampIsDown
	}
	var nonEmptyObservations = getNonEmptyObservations(r.lggr, observations)
	// Need at least F+1 observations
	if len(nonEmptyObservations) <= r.F {
		lggr.Tracew("Non-empty observations <= F, need at least F+1 to continue")
		return false, nil, nil
	}
	// We have at least F+1 valid observations
	// Extract the min and max
	sort.Slice(nonEmptyObservations, func(i, j int) bool {
		return nonEmptyObservations[i].MinSeqNum < nonEmptyObservations[j].MinSeqNum
	})
	// r.F < len(nonEmptyObservations) because of the check above and therefore this is safe
	minSeqNum := nonEmptyObservations[r.F].MinSeqNum
	sort.Slice(nonEmptyObservations, func(i, j int) bool {
		return nonEmptyObservations[i].MaxSeqNum < nonEmptyObservations[j].MaxSeqNum
	})
	// We use a conservative maximum. If we pick a value that some honest oracles might not
	// have seen they’ll end up not agreeing on a report, stalling the protocol.
	maxSeqNum := nonEmptyObservations[r.F].MaxSeqNum
	if maxSeqNum < minSeqNum {
		return false, nil, errors.New("max seq num smaller than min")
	}
	lastRep, err := r.offRamp.GetLastReport(nil)
	if err != nil {
		return false, nil, err
	}
	if lastRep.MerkleRoot == [32]byte{} {
		return false, nil, errors.New("no relayed report")
	}
	if minSeqNum > lastRep.MaxSequenceNumber {
		return false, nil, errors.New("min seq num greater than max relayed seq num")
	}
	msgs, err := r.getMessagesInRangeWithExecutor(minSeqNum, maxSeqNum, r.executor)
	if err != nil {
		return false, nil, err
	}
	if len(msgs) == 0 {
		lggr.Infow("No messages to execute")
		return false, nil, nil
	}
	msgs = msgs[:min(uint64(len(msgs)), MaxNumMessagesInExecutionReport)]
	minActualSeqNum, maxActualSeqNum := msgs[0].Message.SequenceNumber, msgs[len(msgs)-1].Message.SequenceNumber
	// Find the root for each message by looking at all the relayed reports
	// Assumes the return relayed reports are sorted.
	reports, err := r.getRelayedReports(minActualSeqNum, maxActualSeqNum)
	if err != nil {
		return false, nil, err
	}
	if len(reports) == 0 {
		lggr.Infow("Executable message present, but reports not relayed yet", "numExecutable", len(msgs))
		return false, nil, nil
	}
	// We only operate on one report at a time
	report, err := r.buildReport(lggr, reports[0], msgs)
	if err != nil {
		return false, nil, err
	}
	return true, report, nil
}

func (r *ExecutionReportingPlugin) updateInFlight(lggr logger.Logger, er ExecutionReport) error {
	r.inFlightMu.Lock()
	defer r.inFlightMu.Unlock()
	// Reap old inflights and check if any messages in the report are inflight.
	var stillInFlight []InflightExecutionReport
	for _, report := range r.inFlight {
		// TODO: Think about if this fails in reorgs
		if report.report.Messages[0].SequenceNumber == er.Messages[0].SequenceNumber {
			return errors.Errorf("report is already in flight")
		}
		if time.Since(report.createdAt) < ExecutionMaxInflightTimeSeconds {
			stillInFlight = append(stillInFlight, report)
		} else {
			lggr.Warnw("Inflight report expired, retrying", "min", report.report.Messages[0].SequenceNumber, "max", report.report.Messages[len(report.report.Messages)-1].SequenceNumber)
		}
	}
	// Add new inflight
	r.inFlight = append(stillInFlight, InflightExecutionReport{
		createdAt: time.Now(),
		report:    er,
	})
	return nil
}

func (r *ExecutionReportingPlugin) ShouldAcceptFinalizedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	lggr := r.lggr.Named("ShouldAcceptFinalizedReport")
	er, err := DecodeExecutionReport(report)
	if err != nil {
		return false, nil
	}
	if len(er.Messages) == 0 {
		r.lggr.Warnw("Received empty report")
		return false, nil
	}
	var seqNums []uint64
	for i := range er.Messages {
		seqNums = append(seqNums, er.Messages[i].SequenceNumber)
		lggr.Infof("msg amounts %s", er.Messages[i].Payload.Amounts[0].String())
	}
	lggr.Infof("Seq nums %v proofs %+v proof bits %s", seqNums, er.Proofs, er.ProofFlagBits.String())
	// If the first message is executed already, this execution report is stale, and we do not accept it.
	stale, err := r.isStaleReport(seqNums[0])
	if err != nil {
		return !stale, err
	}
	if stale {
		return false, err
	}
	if err := r.updateInFlight(lggr, *er); err != nil {
		return false, err
	}
	return true, nil
}

func (r *ExecutionReportingPlugin) ShouldTransmitAcceptedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	parsedReport, err := DecodeExecutionReport(report)
	if err != nil {
		return false, nil
	}
	// If report is not stale we transmit.
	// When the executeTransmitter enqueues the tx for bptxm,
	// we mark it as execution_sent, removing it from the set of inflight messages.
	stale, err := r.isStaleReport(parsedReport.Messages[0].SequenceNumber)
	return !stale, err
}

func (r *ExecutionReportingPlugin) isStaleReport(min uint64) (bool, error) {
	// If the first message is executed already, this execution report is stale.
	executedMap, err := r.getExecutedMessages()
	if err != nil {
		return true, err
	}
	if _, ok := executedMap[min]; ok {
		return true, nil
	}
	return false, nil
}

func (r *ExecutionReportingPlugin) Close() error {
	return nil
}
