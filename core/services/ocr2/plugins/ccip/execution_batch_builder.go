package ccip

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/blob_verifier"
)

const (
	PERMISSIONLESS_EXECUTION_THRESHOLD_SECONDS = 2 * 7 * 24 * 60 * 60
)

type BatchBuilder interface {
	BuildBatch(
		srcToDst map[common.Address]common.Address,
		msgs []logpoller.Log,
		executed map[uint64]struct{},
		gasLimit uint64,
		gasPrice uint64,
		tokensPerFeeCoin map[common.Address]uint64,
		inflight []InflightExecutionReport) []uint64
}

type ExecutionBatchBuilder struct {
	gasLimit                   uint64
	snoozeTime                 time.Duration
	builder                    BatchBuilder
	blobVerifier               *blob_verifier.BlobVerifier
	onRamp                     common.Address
	offRampAddr                common.Address
	offRamp                    OffRamp
	srcLogPoller, dstLogPoller logpoller.LogPoller
	config                     OffchainConfig
	snoozedRoots               map[[32]byte]time.Time
}

func NewExecutionBatchBuilder(gasLimit uint64, snoozeTime time.Duration, blobVerifier *blob_verifier.BlobVerifier, onRamp, offRampAddr common.Address, srcLogPoller, dstLogPoller logpoller.LogPoller, builder BatchBuilder, config OffchainConfig, offRamp OffRamp) *ExecutionBatchBuilder {
	return &ExecutionBatchBuilder{
		gasLimit:     gasLimit,
		snoozeTime:   snoozeTime,
		builder:      builder,
		blobVerifier: blobVerifier,
		dstLogPoller: dstLogPoller,
		srcLogPoller: srcLogPoller,
		offRamp:      offRamp,
		onRamp:       onRamp,
		offRampAddr:  offRampAddr,
		config:       config,
		snoozedRoots: make(map[[32]byte]time.Time),
	}
}

func (eb *ExecutionBatchBuilder) relayedReport(seqNr uint64) (blob_verifier.CCIPRelayReport, error) {
	latest, err := eb.dstLogPoller.LatestBlock()
	if err != nil {
		return blob_verifier.CCIPRelayReport{}, err
	}
	// Since the report accepted logs now contain intervals per onramp, we don't have a simple way of looking
	// up the relayed report for a given sequence number from the chain.
	// TODO: Follow up with a more efficient way, ideally we use the chain only to obtain natural reorg self-healing.
	// One option is to emit a log per onramp (i.e. ReportAccepted(root, onRamp, min, max)) so we could easily search for the relevant log?
	logs, err := eb.dstLogPoller.Logs(1, latest, ReportAccepted, eb.blobVerifier.Address())
	if err != nil {
		return blob_verifier.CCIPRelayReport{}, err
	}
	if len(logs) == 0 {
		return blob_verifier.CCIPRelayReport{}, errors.Errorf("seq number not relayed, nothing relayed")
	}
	for _, log := range logs {
		reportAccepted, err := eb.blobVerifier.ParseReportAccepted(types.Log{
			Topics: log.GetTopics(),
			Data:   log.Data,
		})
		if err != nil {
			return blob_verifier.CCIPRelayReport{}, err
		}
		for i, onRamp := range reportAccepted.Report.OnRamps {
			if onRamp == eb.onRamp {
				if reportAccepted.Report.Intervals[i].Min <= seqNr && seqNr <= reportAccepted.Report.Intervals[i].Max {
					return reportAccepted.Report, nil
				}
			}
		}
	}
	return blob_verifier.CCIPRelayReport{}, errors.Errorf("seq number not relayed")
}

func (eb *ExecutionBatchBuilder) getUnexpiredRelayReports() ([]blob_verifier.CCIPRelayReport, error) {
	logs, err := eb.dstLogPoller.LogsCreatedAfter(ReportAccepted, eb.blobVerifier.Address(), time.Now().Add(-PERMISSIONLESS_EXECUTION_THRESHOLD_SECONDS*time.Second))
	if err != nil {
		return nil, err
	}
	var reports []blob_verifier.CCIPRelayReport
	for _, log := range logs {
		reportAccepted, err := eb.blobVerifier.ParseReportAccepted(types.Log{
			Topics: log.GetTopics(),
			Data:   log.Data,
		})
		if err != nil {
			return nil, err
		}
		reports = append(reports, reportAccepted.Report)
	}
	return reports, nil
}

func (eb *ExecutionBatchBuilder) getExecutedSeqNrsInRange(min, max uint64) (map[uint64]struct{}, error) {
	executedLogs, err := eb.dstLogPoller.IndexedLogsTopicRange(CrossChainMessageExecuted, eb.offRampAddr, CrossChainMessageExecutedSequenceNumberIndex, logpoller.EvmWord(min), logpoller.EvmWord(max), int(eb.config.DestIncomingConfirmations))
	if err != nil {
		return nil, err
	}
	executedMp := make(map[uint64]struct{})
	for _, executedLog := range executedLogs {
		e, err := eb.offRamp.ParseExecutionCompleted(types.Log{Data: executedLog.Data, Topics: executedLog.GetTopics()})
		if err != nil {
			return nil, err
		}
		executedMp[e.SequenceNumber] = struct{}{}
	}
	return executedMp, nil
}

func (eb *ExecutionBatchBuilder) getExecutableSeqNrs(maxGasPrice uint64, tokensPerFeeCoin map[common.Address]uint64, inflight []InflightExecutionReport) ([]uint64, error) {
	unexpiredReports, err := eb.getUnexpiredRelayReports()
	if err != nil {
		return nil, err
	}
	for _, unexpiredReport := range unexpiredReports {
		var idx int
		for i, onRamp := range unexpiredReport.OnRamps {
			if onRamp == eb.onRamp {
				idx = i
			}
		}
		snoozeUntil, haveSnoozed := eb.snoozedRoots[unexpiredReport.MerkleRoots[idx]]
		if haveSnoozed && time.Now().Before(snoozeUntil) {
			continue
		}
		// Check this root for executable messages
		srcLogs, err := eb.srcLogPoller.LogsDataWordRange(CCIPSendRequested, eb.onRamp, SendRequestedSequenceNumberIndex, logpoller.EvmWord(unexpiredReport.Intervals[idx].Min), logpoller.EvmWord(unexpiredReport.Intervals[idx].Max), int(eb.config.SourceIncomingConfirmations))
		if err != nil {
			return nil, err
		}
		if len(srcLogs) == 0 {
			return nil, errors.Errorf("unexpected empty log set for root %v", unexpiredReport.MerkleRoots[idx])
		}
		executedMp, err := eb.getExecutedSeqNrsInRange(unexpiredReport.Intervals[idx].Min, unexpiredReport.Intervals[idx].Max)
		if err != nil {
			return nil, err
		}
		srcToDst := make(map[common.Address]common.Address)
		sourceTokens, err := eb.offRamp.GetPoolTokens(nil)
		if err != nil {
			return nil, err
		}
		for _, sourceToken := range sourceTokens {
			dst, err := eb.offRamp.GetPool(nil, sourceToken)
			if err != nil {
				return nil, err
			}
			srcToDst[sourceToken] = dst
		}
		batch := eb.builder.BuildBatch(srcToDst, srcLogs, executedMp, eb.gasLimit, maxGasPrice, tokensPerFeeCoin, inflight)
		if len(batch) != 0 {
			return batch, nil
		}
		eb.snoozedRoots[unexpiredReport.MerkleRoots[idx]] = time.Now().Add(eb.snoozeTime)
	}
	return []uint64{}, nil
}
