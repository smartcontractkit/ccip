package ccip

import (
	"context"
	"encoding/json"
	"math/big"
	"sort"
	"sync"
	"time"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/onramp"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/merklemulti"
)

const RelayMaxInflightTimeSeconds = 180

var (
	_ types.ReportingPluginFactory = &RelayReportingPluginFactory{}
	_ types.ReportingPlugin        = &RelayReportingPlugin{}
)

// EncodeRelayReport abi encodes an offramp.CCIPRelayReport.
func EncodeRelayReport(relayReport *offramp.CCIPRelayReport) (types.Report, error) {
	report, err := makeRelayReportArgs().PackValues([]interface{}{relayReport.MerkleRoot, relayReport.MinSequenceNumber, relayReport.MaxSequenceNumber})
	if err != nil {
		return nil, err
	}
	return report, nil
}

// DecodeRelayReport abi decodes a types.Report to an offramp.CCIPRelayReport
func DecodeRelayReport(report types.Report) (*offramp.CCIPRelayReport, error) {
	unpacked, err := makeRelayReportArgs().Unpack(report)
	if err != nil {
		return nil, err
	}
	if len(unpacked) != 3 {
		return nil, errors.New("invalid num fields in report")
	}
	root, ok := unpacked[0].([32]byte)
	if !ok {
		return nil, errors.New("invalid root")
	}
	min, ok := unpacked[1].(uint64)
	if !ok {
		return nil, errors.New("invalid min")
	}
	max, ok := unpacked[2].(uint64)
	if !ok {
		return nil, errors.New("invalid max")
	}
	return &offramp.CCIPRelayReport{
		MerkleRoot:        root,
		MinSequenceNumber: min,
		MaxSequenceNumber: max,
	}, nil
}

// parseLogs preserves the log order.
func parseLogs(onRamp *onramp.OnRamp, logs []logpoller.Log) ([]onramp.OnRampCrossChainSendRequested, error) {
	var unstartedReqs []onramp.OnRampCrossChainSendRequested
	for _, log := range logs {
		reqParsed, err := onRamp.ParseCrossChainSendRequested(gethtypes.Log{
			Data:   log.Data,
			Topics: log.GetTopics(),
		})
		if err != nil {
			return nil, err
		}
		unstartedReqs = append(unstartedReqs, *reqParsed)
	}
	return unstartedReqs, nil
}

func isOffRampDownNow(lggr logger.Logger, offRamp *offramp.OffRamp) bool {
	paused, err := offRamp.Paused(nil)
	if err != nil {
		// Air on side of caution by halting if we cannot read the state?
		lggr.Errorw("Unable to read offramp paused", "err", err)
		return true
	}
	healthy, err := offRamp.IsHealthy(nil, big.NewInt(time.Now().Unix()))
	if err != nil {
		lggr.Errorw("Unable to read offramp afn", "err", err)
		return true
	}
	return paused || !healthy
}

type InflightReport struct {
	report    *offramp.CCIPRelayReport
	createdAt time.Time
}

type RelayReportingPluginFactory struct {
	lggr         logger.Logger
	source       *logpoller.LogPoller
	onRamp       *onramp.OnRamp
	offRamp      *offramp.OffRamp
	configPoller *ConfigPoller
}

// NewRelayReportingPluginFactory return a new RelayReportingPluginFactory.
func NewRelayReportingPluginFactory(
	lggr logger.Logger,
	source *logpoller.LogPoller,
	offRamp *offramp.OffRamp,
	onRamp *onramp.OnRamp,
	configPoller *ConfigPoller,
) types.ReportingPluginFactory {
	return &RelayReportingPluginFactory{lggr: lggr, offRamp: offRamp, onRamp: onRamp, source: source, configPoller: configPoller}
}

// NewReportingPlugin returns the ccip RelayReportingPlugin and satisfies the ReportingPluginFactory interface.
func (rf *RelayReportingPluginFactory) NewReportingPlugin(config types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	return &RelayReportingPlugin{
			lggr:         rf.lggr.Named("RelayReportingPlugin"),
			F:            config.F,
			source:       rf.source,
			onRamp:       rf.onRamp,
			offRamp:      rf.offRamp,
			inFlight:     make(map[[32]byte]InflightReport),
			configPoller: rf.configPoller,
		},
		types.ReportingPluginInfo{
			Name:          "CCIPRelay",
			UniqueReports: true,
			Limits: types.ReportingPluginLimits{
				MaxQueryLength: 0,
				// TODO: https://app.shortcut.com/chainlinklabs/story/30171/define-report-plugin-limits
				MaxObservationLength: 100000, // TODO
				MaxReportLength:      100000, // TODO
			},
		}, nil
}

type RelayReportingPlugin struct {
	lggr    logger.Logger
	F       int
	source  *logpoller.LogPoller
	onRamp  *onramp.OnRamp
	offRamp *offramp.OffRamp
	// We need to synchronize access to the inflight structure
	// as reporting plugin methods may be called from separate goroutines,
	// e.g. reporting vs transmission protocol.
	inFlightMu   sync.RWMutex
	inFlight     map[[32]byte]InflightReport
	configPoller *ConfigPoller
}

func (r *RelayReportingPlugin) nextMinSeqNumForOffRamp() (uint64, error) {
	lastReport, err := r.offRamp.GetLastReport(nil)
	if err != nil {
		return 0, err
	}
	if lastReport.MerkleRoot == [32]byte{} {
		return 1, nil
	}
	return lastReport.MaxSequenceNumber + 1, nil
}

func (r *RelayReportingPlugin) nextMinSeqNumForInFlight() uint64 {
	r.inFlightMu.RLock()
	defer r.inFlightMu.RUnlock()
	max := uint64(0)
	for _, report := range r.inFlight {
		if report.report.MaxSequenceNumber > max {
			max = report.report.MaxSequenceNumber
		}
	}
	return max + 1
}

func (r *RelayReportingPlugin) nextMinSeqNum() (uint64, error) {
	nextMin, err := r.nextMinSeqNumForOffRamp()
	if err != nil {
		return 0, err
	}
	nextMinInFlight := r.nextMinSeqNumForInFlight()
	if nextMinInFlight > nextMin {
		nextMin = nextMinInFlight
	}
	return nextMin, nil
}

func (r *RelayReportingPlugin) contiguousReqs(min, max uint64, reqs []onramp.OnRampCrossChainSendRequested) bool {
	for i, j := min, 0; i < max && j < len(reqs); i, j = i+1, j+1 {
		if reqs[j].Message.SequenceNumber != i {
			r.lggr.Errorw("unexpected gap in seq nums", "seq", i)
			return false
		}
	}
	return true
}

func (r *RelayReportingPlugin) Query(ctx context.Context, timestamp types.ReportTimestamp) (types.Query, error) {
	return types.Query{}, nil
}

func (r *RelayReportingPlugin) Observation(ctx context.Context, timestamp types.ReportTimestamp, query types.Query) (types.Observation, error) {
	lggr := r.lggr.Named("Observation")
	if isOffRampDownNow(lggr, r.offRamp) {
		return nil, errors.New("offRamp is down")
	}
	minSeqNum, err := r.nextMinSeqNum()
	if err != nil {
		return nil, err
	}
	// All available messages that have not been relayed yet and have sufficient confirmations.
	reqs, err := r.source.LogsDataWordGreaterThan(CrossChainSendRequested, r.onRamp.Address(), 2, EvmWord(minSeqNum), r.configPoller.sourceConfs())
	if err != nil {
		return nil, err
	}
	unstartedReqs, err := parseLogs(r.onRamp, reqs)
	if err != nil {
		return nil, err
	}
	if len(unstartedReqs) == 0 {
		return []byte{}, nil
	}
	min := unstartedReqs[0].Message.SequenceNumber
	max := unstartedReqs[len(unstartedReqs)-1].Message.SequenceNumber
	if !r.contiguousReqs(min, max, unstartedReqs) {
		return nil, errors.New("unexpected gap in seq nums")
	}
	lggr.Infof("Messages %v", unstartedReqs)
	return json.Marshal(&Observation{
		MinSeqNum: min,
		MaxSeqNum: max,
	})
}

// buildReport assumes there is at least one message in reqs.
func (r *RelayReportingPlugin) buildReport(minSeqNum, maxSeqNum uint64) (*offramp.CCIPRelayReport, error) {
	// Logs are guaranteed to be in order of seq num, since these are finalized logs only
	// and the contract's seq num is auto-incrementing.
	logs, err := r.source.LogsDataWordRange(CrossChainSendRequested, r.onRamp.Address(), 2, EvmWord(minSeqNum), EvmWord(maxSeqNum), r.configPoller.sourceConfs())
	if err != nil {
		return nil, err
	}
	reqs, err := parseLogs(r.onRamp, logs)
	if err != nil {
		return nil, err
	}
	if !r.contiguousReqs(minSeqNum, maxSeqNum, reqs) {
		return nil, errors.New("unexpected gap in seq nums")
	}
	// Take all these request and produce a merkle root of them
	mctx := merklemulti.NewKeccakCtx()
	var leaves [][32]byte
	for _, req := range reqs {
		leaves = append(leaves, mctx.HashLeaf(req.Raw.Data))
	}
	tree := merklemulti.NewTree(mctx, leaves)
	return &offramp.CCIPRelayReport{
		MerkleRoot:        tree.Root(),
		MinSequenceNumber: reqs[0].Message.SequenceNumber,
		MaxSequenceNumber: reqs[len(reqs)-1].Message.SequenceNumber,
	}, nil
}

func (r *RelayReportingPlugin) Report(ctx context.Context, timestamp types.ReportTimestamp, query types.Query, observations []types.AttributedObservation) (bool, types.Report, error) {
	lggr := r.lggr.Named("Report")
	if isOffRampDownNow(lggr, r.offRamp) {
		return false, nil, errors.New("offRamp is down")
	}
	var nonEmptyObservations = getNonEmptyObservations(r.lggr, observations)
	// Need at least F+1 valid observations
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
	// have seen they’ll end up not agreeing on a msg, stalling the protocol.
	maxSeqNum := nonEmptyObservations[r.F].MaxSeqNum
	if maxSeqNum < minSeqNum {
		return false, nil, errors.New("max seq num smaller than min")
	}
	nextMin, err := r.nextMinSeqNumForOffRamp()
	if err != nil {
		return false, nil, err
	}
	// Contract would revert
	if nextMin > minSeqNum {
		return false, nil, errors.Errorf("invalid min seq number got %v want %v", minSeqNum, nextMin)
	}
	report, err := r.buildReport(minSeqNum, maxSeqNum)
	if err != nil {
		return false, nil, err
	}
	encodedReport, err := EncodeRelayReport(report)
	if err != nil {
		return false, nil, err
	}
	lggr.Infow("Built report", "min", minSeqNum, "max", maxSeqNum)
	return true, encodedReport, nil
}

func (r *RelayReportingPlugin) updateInflight(lggr logger.Logger, report *offramp.CCIPRelayReport) {
	r.inFlightMu.Lock()
	defer r.inFlightMu.Unlock()
	// Reap any expired entries from inflight.
	for root, inFlightReport := range r.inFlight {
		if time.Since(inFlightReport.createdAt) > RelayMaxInflightTimeSeconds {
			lggr.Warnw("Inflight report expired, retrying", "min", inFlightReport.report.MinSequenceNumber, "max", inFlightReport.report.MaxSequenceNumber)
			delete(r.inFlight, root)
		}
	}
	// Set new inflight ones as pending
	r.inFlight[report.MerkleRoot] = InflightReport{
		report:    report,
		createdAt: time.Now(),
	}
}

func (r *RelayReportingPlugin) ShouldAcceptFinalizedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	lggr := r.lggr.Named("ShouldAcceptFinalizedReport")
	parsedReport, err := DecodeRelayReport(report)
	if err != nil {
		return false, nil
	}
	// Note it's ok to leave the unstarted requests behind, since the
	// 'Observe' is always based on the last reports onchain min seq num.
	if r.isStaleReport(parsedReport) {
		return false, nil
	}
	r.updateInflight(lggr, parsedReport)
	lggr.Infow("Accepting finalized report", "min", parsedReport.MinSequenceNumber, "max", parsedReport.MaxSequenceNumber)
	return true, nil
}

func (r *RelayReportingPlugin) ShouldTransmitAcceptedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	parsedReport, err := DecodeRelayReport(report)
	if err != nil {
		return false, nil
	}
	// If report is not stale we transmit.
	// When the relayTransmitter enqueues the tx for bptxm,
	// we mark it as fulfilled, effectively removing it from the set of inflight messages.
	return !r.isStaleReport(parsedReport), nil
}

func (r *RelayReportingPlugin) isStaleReport(report *offramp.CCIPRelayReport) bool {
	if isOffRampDownNow(r.lggr, r.offRamp) {
		return true
	}
	nextMin, err := r.nextMinSeqNumForOffRamp()
	if err != nil {
		// Assume it's a transient issue getting the last report
		// Will try again on the next round
		return true
	}
	// If the next min is already greater than this reports min,
	// this report is stale.
	if nextMin > report.MinSequenceNumber {
		r.lggr.Warnw("report is stale", "onchain min", nextMin, "report min", report.MinSequenceNumber)
		return true
	}
	return false
}

func (r *RelayReportingPlugin) Close() error {
	return nil
}
