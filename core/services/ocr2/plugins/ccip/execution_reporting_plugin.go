package ccip

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/onramp"
	"github.com/smartcontractkit/chainlink/core/logger"
)

const (
	ExecutionMaxInflightTimeSeconds = 180
	MaxNumMessagesInExecutionReport = 100
)

var (
	_                types.ReportingPluginFactory = &ExecutionReportingPluginFactory{}
	_                types.ReportingPlugin        = &ExecutionReportingPlugin{}
	ErrOffRampIsDown                              = errors.New("offramp is down")
)

func EncodeExecutionReport(ems []ExecutableMessage) (types.Report, error) {
	report, err := makeExecutionReportArgs().PackValues([]interface{}{ems})
	if err != nil {
		return nil, err
	}
	return report, nil
}

func DecodeExecutionReport(report types.Report) ([]ExecutableMessage, error) {
	unpacked, err := makeExecutionReportArgs().Unpack(report)
	if err != nil {
		return nil, err
	}
	if len(unpacked) == 0 {
		return nil, errors.New("assumptionViolation: expected at least one element")
	}

	// Must be anonymous struct here
	msgs, ok := unpacked[0].([]struct {
		Path    [][32]uint8 `json:"Path"`
		Index   *big.Int    `json:"Index"`
		Message struct {
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
				Options            []uint8          `json:"options"`
			} `json:"payload"`
		} `json:"Message"`
	})
	if !ok {
		return nil, fmt.Errorf("got %T", unpacked[0])
	}
	if len(msgs) == 0 {
		return nil, errors.New("assumptionViolation: expected at least one element")
	}
	var ems []ExecutableMessage
	for _, emi := range msgs {
		ems = append(ems, ExecutableMessage{
			Path:    emi.Path,
			Index:   emi.Index,
			Message: emi.Message,
		})
	}
	return ems, nil
}

//go:generate mockery --name OffRampLastReporter --output ./mocks/lastreporter --case=underscore
type OffRampLastReporter interface {
	GetLastReport(opts *bind.CallOpts) (offramp.CCIPRelayReport, error)
}

type ExecutionReportingPluginFactory struct {
	lggr         logger.Logger
	lastReporter OffRampLastReporter
	source, dest *logpoller.LogPoller
	executor     common.Address
	onRamp       *onramp.OnRamp
	offRamp      *offramp.OffRamp
	configPoller *ConfigPoller
}

func NewExecutionReportingPluginFactory(
	lggr logger.Logger,
	onRamp *onramp.OnRamp,
	offRamp *offramp.OffRamp,
	source, dest *logpoller.LogPoller,
	executor common.Address,
	lastReporter OffRampLastReporter,
	configPoller *ConfigPoller,
) types.ReportingPluginFactory {
	return &ExecutionReportingPluginFactory{lggr: lggr, onRamp: onRamp, offRamp: offRamp, executor: executor, source: source, dest: dest, lastReporter: lastReporter, configPoller: configPoller}
}

func (rf *ExecutionReportingPluginFactory) NewReportingPlugin(config types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	return &ExecutionReportingPlugin{
			lggr:         rf.lggr.Named("ExecutionReportingPlugin"),
			F:            config.F,
			executor:     rf.executor,
			onRamp:       rf.onRamp,
			offRamp:      rf.offRamp,
			source:       rf.source,
			dest:         rf.dest,
			lastReporter: rf.lastReporter,
			configPoller: rf.configPoller,
		}, types.ReportingPluginInfo{
			Name:          "CCIPExecution",
			UniqueReports: true,
			Limits: types.ReportingPluginLimits{
				MaxQueryLength: 0,
				// TODO: https://app.shortcut.com/chainlinklabs/story/30171/define-report-plugin-limits
				MaxObservationLength: 100000, // TODO
				MaxReportLength:      100000, // TODO
			},
		}, nil
}

type ExecutionReportingPlugin struct {
	lggr         logger.Logger
	F            int
	executor     common.Address
	onRamp       *onramp.OnRamp
	offRamp      *offramp.OffRamp
	source, dest *logpoller.LogPoller
	// We also use the offramp for defensive checks
	lastReporter OffRampLastReporter
	inFlight     []InflightExecutionReport
	configPoller *ConfigPoller
}

type InflightExecutionReport struct {
	createdAt time.Time
	msgs      []ExecutableMessage
}

func (r *ExecutionReportingPlugin) Query(ctx context.Context, timestamp types.ReportTimestamp) (types.Query, error) {
	// We don't use a query for this reporting plugin, so we can just leave it empty here
	return types.Query{}, nil
}

// getRelayedReports returns them in sorted order.
func (r *ExecutionReportingPlugin) getRelayedReports(min, max uint64) ([]offramp.OffRampReportAccepted, error) {
	// Get all reports where minSeqNum is >= min as a lower bound.
	reportLogs, err := r.dest.LogsDataWordGreaterThan(ReportAccepted, r.offRamp.Address(), 1, evmWord(min), 1)
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
	inFlightSeqNums := make(map[uint64]struct{})
	for _, report := range r.inFlight {
		for _, msg := range report.msgs {
			inFlightSeqNums[msg.Message.SequenceNumber] = struct{}{}
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
	rep, err := r.lastReporter.GetLastReport(nil)
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
	return json.Marshal(&Observation{
		MinSeqNum: executable[0],
		MaxSeqNum: executable[len(executable)-1],
	})
}

func (r *ExecutionReportingPlugin) getMessagesInRangeWithExecutor(min, max uint64, executor common.Address) ([]onramp.OnRampCrossChainSendRequested, error) {
	msgs, err := r.source.LogsDataWordRange(CrossChainSendRequested, r.onRamp.Address(), 2, evmWord(min), evmWord(max), r.configPoller.sourceConfs())
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

// Assumes non-empty msgs
func (r *ExecutionReportingPlugin) buildReport(lggr logger.Logger, msgs []onramp.OnRampCrossChainSendRequested) ([]byte, error) {
	// Bound the execution report size.
	msgs = msgs[:min(uint64(len(msgs)), MaxNumMessagesInExecutionReport)]
	minActualSeqNum, maxActualSeqNum := msgs[0].Message.SequenceNumber, msgs[len(msgs)-1].Message.SequenceNumber
	// Find the root for each message by looking at all the relayed reports
	// Assumes the return relayed reports are sorted.
	reports, err := r.getRelayedReports(minActualSeqNum, maxActualSeqNum)
	if err != nil {
		return nil, err
	}
	var ems []ExecutableMessage
	i := int64(0)
	for _, report := range reports {
		for ; i < int64(len(msgs)); i++ {
			if msgs[i].Message.SequenceNumber > report.Report.MaxSequenceNumber {
				break
			}
			// its in the range of the msgs, get all messages in the msgs
			allMsgs, err2 := r.getMessagesInRangeWithExecutor(report.Report.MinSequenceNumber, report.Report.MaxSequenceNumber, [20]byte{})
			if err2 != nil {
				return nil, err2
			}
			if len(allMsgs) != int(report.Report.MaxSequenceNumber-report.Report.MinSequenceNumber+1) {
				return nil, errors.Errorf("do not have all messages, have %d want %d", len(allMsgs), int(report.Report.MaxSequenceNumber-report.Report.MinSequenceNumber+1))
			}
			var leaves [][]byte
			for _, msg := range allMsgs {
				leaves = append(leaves, msg.Raw.Data)
			}
			index := msgs[i].Message.SequenceNumber - report.Report.MinSequenceNumber
			if index < 0 {
				return nil, errors.New("unexpected invalid index")
			}
			root, proof := GenerateMerkleProof(32, leaves, int(index))
			if !bytes.Equal(root[:], report.Report.MerkleRoot[:]) {
				lggr.Errorw("Invalid merkle root generated", "have", hexutil.Encode(root[:]), "want", hexutil.Encode(report.Report.MerkleRoot[:]), "proving", msgs[i])
				continue
			}
			ems = append(ems, ExecutableMessage{
				Path: proof.PathForExecute(),
				Message: Message{
					SequenceNumber: msgs[i].Message.SequenceNumber,
					SourceChainId:  msgs[i].Message.SourceChainId,
					Sender:         msgs[i].Message.Sender,
					Payload: struct {
						Tokens             []common.Address `json:"tokens"`
						Amounts            []*big.Int       `json:"amounts"`
						DestinationChainId *big.Int         `json:"destinationChainId"`
						Receiver           common.Address   `json:"receiver"`
						Executor           common.Address   `json:"executor"`
						Data               []uint8          `json:"data"`
						Options            []uint8          `json:"options"`
					}{
						Tokens:             msgs[i].Message.Payload.Tokens,
						Amounts:            msgs[i].Message.Payload.Amounts,
						DestinationChainId: msgs[i].Message.Payload.DestinationChainId,
						Receiver:           msgs[i].Message.Payload.Receiver,
						Executor:           msgs[i].Message.Payload.Executor,
						Data:               msgs[i].Message.Payload.Data,
						Options:            msgs[i].Message.Payload.Options,
					},
				},
				Index: proof.Index(),
			})
		}
	}
	if len(ems) == 0 {
		return nil, errors.New("no executed messages created")
	}
	report, err := EncodeExecutionReport(ems)
	if err != nil {
		return nil, err
	}
	return report, nil
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
	// have seen they’ll end up not agreeing on a msgs, stalling the protocol.
	maxSeqNum := nonEmptyObservations[r.F].MaxSeqNum
	if maxSeqNum < minSeqNum {
		return false, nil, errors.New("max seq num smaller than min")
	}
	lastRep, err := r.lastReporter.GetLastReport(nil)
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
	report, err := r.buildReport(lggr, msgs)
	if err != nil {
		return false, nil, err
	}
	return true, report, nil
}

func (r *ExecutionReportingPlugin) ShouldAcceptFinalizedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	lggr := r.lggr.Named("ShouldAcceptFinalizedReport")
	ems, err := DecodeExecutionReport(report)
	if err != nil {
		return false, nil
	}
	if len(ems) == 0 {
		r.lggr.Warnw("Received empty report")
		return false, nil
	}
	var seqNums []uint64
	for i := range ems {
		seqNums = append(seqNums, ems[i].Message.SequenceNumber)
	}
	lggr.Infof("Seq nums %v", seqNums)
	// If the first message is executed already, this execution msgs is stale, and we do not accept it.
	stale, err := r.isStaleReport(seqNums[0])
	if err != nil {
		return !stale, err
	}
	if stale {
		return false, err
	}
	// Reap old inflights and check if any messages in the report are inflight.
	var stillInFlight []InflightExecutionReport
	for _, report := range r.inFlight {
		// TODO: Think about if this fails in reorgs
		if report.msgs[0].Message.SequenceNumber == ems[0].Message.SequenceNumber {
			return false, errors.Errorf("report is already in flight")
		}
		if time.Since(report.createdAt) < ExecutionMaxInflightTimeSeconds {
			stillInFlight = append(stillInFlight, report)
		} else {
			lggr.Warnw("Inflight report expired, retrying", "min", report.msgs[0].Message.SequenceNumber, "max", report.msgs[len(report.msgs)-1].Message.SequenceNumber)
		}
	}
	// Add new inflight
	r.inFlight = append(stillInFlight, InflightExecutionReport{
		createdAt: time.Now(),
		msgs:      ems,
	})
	return true, nil
}

func (r *ExecutionReportingPlugin) ShouldTransmitAcceptedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	parsedReport, err := DecodeExecutionReport(report)
	if err != nil {
		return false, nil
	}
	// If msgs is not stale we transmit.
	// When the executeTransmitter enqueues the tx for bptxm,
	// we mark it as execution_sent, removing it from the set of inflight messages.
	stale, err := r.isStaleReport(parsedReport[0].Message.SequenceNumber)
	return !stale, err
}

func (r *ExecutionReportingPlugin) isStaleReport(min uint64) (bool, error) {
	// If the first message is executed already, this execution msgs is stale.
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
