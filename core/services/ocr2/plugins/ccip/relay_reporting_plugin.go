package ccip

import (
	"context"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/blob_verifier"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/merklemulti"
)

const (
	RelayMaxInflightTimeSeconds = 180
	MaxRelayReportLength        = 1000 // TODO: Need to rethink this based on root of roots report.
)

var (
	_ types.ReportingPluginFactory = &RelayReportingPluginFactory{}
	_ types.ReportingPlugin        = &RelayReportingPlugin{}
)

// EncodeRelayReport abi encodes an offramp.CCIPRelayReport.
func EncodeRelayReport(relayReport *blob_verifier.CCIPRelayReport) (types.Report, error) {
	report, err := makeRelayReportArgs().PackValues([]interface{}{
		relayReport,
	})
	if err != nil {
		return nil, err
	}
	return report, nil
}

// DecodeRelayReport abi decodes a types.Report to an offramp.CCIPRelayReport
func DecodeRelayReport(report types.Report) (*blob_verifier.CCIPRelayReport, error) {
	unpacked, err := makeRelayReportArgs().Unpack(report)
	if err != nil {
		return nil, err
	}
	if len(unpacked) != 1 {
		return nil, errors.New("expected single struct value")
	}
	relayReport, ok := unpacked[0].(struct {
		OnRamps   []common.Address `json:"onRamps"`
		Intervals []struct {
			Min uint64 `json:"min"`
			Max uint64 `json:"max"`
		} `json:"intervals"`
		MerkleRoots [][32]byte `json:"merkleRoots"`
		RootOfRoots [32]byte   `json:"rootOfRoots"`
	})
	if !ok {
		return nil, errors.Errorf("invalid relay report got %T", unpacked[0])
	}
	var intervalsF []blob_verifier.CCIPInterval
	for i := range relayReport.Intervals {
		intervalsF = append(intervalsF, blob_verifier.CCIPInterval{
			Min: relayReport.Intervals[i].Min,
			Max: relayReport.Intervals[i].Max,
		})
	}
	return &blob_verifier.CCIPRelayReport{
		OnRamps:     relayReport.OnRamps,
		Intervals:   intervalsF,
		MerkleRoots: relayReport.MerkleRoots,
		RootOfRoots: relayReport.RootOfRoots,
	}, nil
}

func isBlobVerifierDownNow(lggr logger.Logger, blobVerifier *blob_verifier.BlobVerifier) bool {
	paused, err := blobVerifier.Paused(nil)
	if err != nil {
		// Air on side of caution by halting if we cannot read the state?
		lggr.Errorw("Unable to read offramp paused", "err", err)
		return true
	}
	healthy, err := blobVerifier.IsAFNHealthy(nil)
	if err != nil {
		lggr.Errorw("Unable to read offramp afn", "err", err)
		return true
	}
	return paused || !healthy
}

type InflightReport struct {
	report    *blob_verifier.CCIPRelayReport
	createdAt time.Time
}

type RelayReportingPluginFactory struct {
	lggr                logger.Logger
	source              logpoller.LogPoller
	onRampSeqParsers    map[common.Address]func(log logpoller.Log) (uint64, error)
	onRampToReqEventSig map[common.Address]common.Hash
	onRamps             []common.Address
	blobVerifier        *blob_verifier.BlobVerifier
	onRampToHasher      map[common.Address]LeafHasher[[32]byte]
}

// NewRelayReportingPluginFactory return a new RelayReportingPluginFactory.
func NewRelayReportingPluginFactory(
	lggr logger.Logger,
	source logpoller.LogPoller,
	blobVerifier *blob_verifier.BlobVerifier,
	onRampSeqParsers map[common.Address]func(log logpoller.Log) (uint64, error),
	onRampToReqEventSig map[common.Address]common.Hash,
	onRamps []common.Address,
	onRampToHasher map[common.Address]LeafHasher[[32]byte],
) types.ReportingPluginFactory {
	return &RelayReportingPluginFactory{lggr: lggr, blobVerifier: blobVerifier, onRampToReqEventSig: onRampToReqEventSig, onRampSeqParsers: onRampSeqParsers, onRamps: onRamps, source: source, onRampToHasher: onRampToHasher}
}

// NewReportingPlugin returns the ccip RelayReportingPlugin and satisfies the ReportingPluginFactory interface.
func (rf *RelayReportingPluginFactory) NewReportingPlugin(config types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	offchainConfig, err := Decode(config.OffchainConfig)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	return &RelayReportingPlugin{
			lggr:                rf.lggr.Named("RelayReportingPlugin"),
			F:                   config.F,
			source:              rf.source,
			onRampSeqParsers:    rf.onRampSeqParsers,
			onRampToReqEventSig: rf.onRampToReqEventSig,
			onRamps:             rf.onRamps,
			blobVerifier:        rf.blobVerifier,
			inFlight:            make(map[[32]byte]InflightReport),
			offchainConfig:      offchainConfig,
			onRampToHasher:      rf.onRampToHasher,
		},
		types.ReportingPluginInfo{
			Name:          "CCIPRelay",
			UniqueReports: true,
			Limits: types.ReportingPluginLimits{
				MaxQueryLength:       0,
				MaxObservationLength: MaxObservationLength,
				MaxReportLength:      MaxRelayReportLength,
			},
		}, nil
}

type RelayReportingPlugin struct {
	lggr                logger.Logger
	F                   int
	source              logpoller.LogPoller
	onRamps             []common.Address
	onRampToReqEventSig map[common.Address]common.Hash
	onRampSeqParsers    map[common.Address]func(log logpoller.Log) (uint64, error)
	blobVerifier        *blob_verifier.BlobVerifier
	// We need to synchronize access to the inflight structure
	// as reporting plugin methods may be called from separate goroutines,
	// e.g. reporting vs transmission protocol.
	inFlightMu     sync.RWMutex
	inFlight       map[[32]byte]InflightReport
	offchainConfig OffchainConfig
	onRampToHasher map[common.Address]LeafHasher[[32]byte]
}

func (r *RelayReportingPlugin) nextMinSeqNumForOffRamp(onRamp common.Address) (uint64, error) {
	return r.blobVerifier.GetExpectedNextSequenceNumber(nil, onRamp)
}

func (r *RelayReportingPlugin) nextMinSeqNumForInFlight(onRamp common.Address) uint64 {
	r.inFlightMu.RLock()
	defer r.inFlightMu.RUnlock()
	max := uint64(0)
	for _, report := range r.inFlight {
		// TODO: it is more ergonomic to make it a struct
		for i, or := range report.report.OnRamps {
			if or == onRamp {
				if report.report.Intervals[i].Max > max {
					max = report.report.Intervals[i].Max
				}
			}
		}
	}
	return max + 1
}

func (r *RelayReportingPlugin) nextMinSeqNum(onRamp common.Address) (uint64, error) {
	nextMin, err := r.nextMinSeqNumForOffRamp(onRamp)
	if err != nil {
		return 0, err
	}
	nextMinInFlight := r.nextMinSeqNumForInFlight(onRamp)
	if nextMinInFlight > nextMin {
		nextMin = nextMinInFlight
	}
	return nextMin, nil
}

func (r *RelayReportingPlugin) Query(ctx context.Context, timestamp types.ReportTimestamp) (types.Query, error) {
	return types.Query{}, nil
}

func (r *RelayReportingPlugin) Observation(ctx context.Context, timestamp types.ReportTimestamp, query types.Query) (types.Observation, error) {
	lggr := r.lggr.Named("RelayObservation")
	if isBlobVerifierDownNow(lggr, r.blobVerifier) {
		return nil, ErrBlobVerifierIsDown
	}
	intervalsByOnRamp := make(map[common.Address]blob_verifier.CCIPInterval)
	for _, onRamp := range r.onRamps {
		nextMin, err := r.nextMinSeqNum(onRamp)
		if err != nil {
			return nil, err
		}
		// All available messages that have not been relayed yet and have sufficient confirmations.
		lggr.Infof("Looking for requests with sig %s and nextMin %d on tollOnRamp %s", r.onRampToReqEventSig[onRamp].Hex(), nextMin, onRamp.Hex())
		reqs, err := r.source.LogsDataWordGreaterThan(r.onRampToReqEventSig[onRamp], onRamp, SendRequestedSequenceNumberIndex, EvmWord(nextMin), int(r.offchainConfig.SourceIncomingConfirmations))
		if err != nil {
			return nil, err
		}
		lggr.Infof("%d requests found for tollOnRamp %s", len(reqs), onRamp.Hex())
		if len(reqs) == 0 {
			r.lggr.Infow("no requests", "tollOnRamp", onRamp)
			continue
		}
		var seqNrs []uint64
		for _, req := range reqs {
			seqNr, err := r.onRampSeqParsers[onRamp](req)
			if err != nil {
				r.lggr.Errorw("error parsing seq num", "err", err)
				continue
			}
			seqNrs = append(seqNrs, seqNr)
		}
		min := seqNrs[0]
		max := seqNrs[len(seqNrs)-1]
		if !contiguousReqs(r.lggr, min, max, seqNrs) {
			return nil, errors.New("unexpected gap in seq nums")
		}
		intervalsByOnRamp[onRamp] = blob_verifier.CCIPInterval{
			Min: min,
			Max: max,
		}
		lggr.Infof("tollOnRamp %v: min %v max %v", onRamp, min, max)
	}
	if len(intervalsByOnRamp) == 0 {
		r.lggr.Infow("No observations")
		return []byte{}, nil
	}
	return RelayObservation{
		IntervalsByOnRamp: intervalsByOnRamp,
	}.Marshal()
}

// buildReport assumes there is at least one message in reqs.
func (r *RelayReportingPlugin) buildReport(intervalByOnRamp map[common.Address]blob_verifier.CCIPInterval) (*blob_verifier.CCIPRelayReport, error) {
	leafsByOnRamp, err := leafsFromIntervals(r.lggr, r.onRampToReqEventSig, r.onRampSeqParsers, intervalByOnRamp, r.source, r.onRampToHasher)
	if err != nil {
		return nil, err
	}
	// Produce a root for each onramp, then a root of roots.
	var (
		onRamps   []common.Address
		roots     [][32]byte
		intervals []blob_verifier.CCIPInterval
	)
	mctx := hasher.NewKeccakCtx()
	for onRamp, leaves := range leafsByOnRamp {
		tree := merklemulti.NewTree(mctx, leaves)
		roots = append(roots, tree.Root())
		onRamps = append(onRamps, onRamp)
		interval := intervalByOnRamp[onRamp]
		intervals = append(intervals, blob_verifier.CCIPInterval{
			Min: interval.Min,
			Max: interval.Max,
		})
	}
	// Make a root of roots
	outerTree := merklemulti.NewTree(mctx, roots)
	return &blob_verifier.CCIPRelayReport{
		MerkleRoots: roots,
		Intervals:   intervals,
		OnRamps:     onRamps,
		RootOfRoots: outerTree.Root(),
	}, nil
}

func (r *RelayReportingPlugin) Report(ctx context.Context, timestamp types.ReportTimestamp, query types.Query, observations []types.AttributedObservation) (bool, types.Report, error) {
	lggr := r.lggr.Named("Report")
	if isBlobVerifierDownNow(lggr, r.blobVerifier) {
		return false, nil, ErrBlobVerifierIsDown
	}
	nonEmptyObservations := getNonEmptyObservations[RelayObservation](r.lggr, observations)
	// Need at least F+1 valid observations
	if len(nonEmptyObservations) <= r.F {
		lggr.Debugf("Non-empty observations <= F, need at least F+1 to continue")
		return false, nil, nil
	}
	// Group intervals by onramp.
	intervalsByOnRamp := make(map[common.Address][]blob_verifier.CCIPInterval)
	for _, obs := range nonEmptyObservations {
		for onRamp, interval := range obs.IntervalsByOnRamp {
			intervalsByOnRamp[onRamp] = append(intervalsByOnRamp[onRamp], interval)
		}
	}
	intervalByOnRamp := make(map[common.Address]blob_verifier.CCIPInterval)
	for onRamp, intervals := range intervalsByOnRamp {
		if len(intervals) <= r.F {
			lggr.Debugf("Observations for tollOnRamp %s 1 < #obs <= F, need at least F+1 to continue", onRamp.Hex())
			continue
		}

		// We have at least F+1 valid observations for the given tollOnRamp
		// Extract the min and max
		sort.Slice(intervals, func(i, j int) bool {
			return intervals[i].Min < intervals[j].Min
		})
		// r.F < len(intervals) because of the check above and therefore this is safe
		minSeqNum := intervals[r.F].Min
		sort.Slice(intervals, func(i, j int) bool {
			return intervals[i].Max < intervals[j].Max
		})
		// We use a conservative maximum. If we pick a value that some honest oracles might not
		// have seen they’ll end up not agreeing on a msg, stalling the protocol.
		maxSeqNum := intervals[r.F].Max
		// TODO: Do we for sure want to fail everything here?
		if maxSeqNum < minSeqNum {
			return false, nil, errors.New("max seq num smaller than min")
		}
		nextMin, err := r.nextMinSeqNumForOffRamp(onRamp)
		if err != nil {
			return false, nil, err
		}
		// Contract would revert
		if nextMin > minSeqNum {
			return false, nil, errors.Errorf("invalid min seq number got %v want %v", minSeqNum, nextMin)
		}
		intervalByOnRamp[onRamp] = blob_verifier.CCIPInterval{
			Min: minSeqNum,
			Max: maxSeqNum,
		}
	}
	report, err := r.buildReport(intervalByOnRamp)
	if err != nil {
		return false, nil, err
	}
	encodedReport, err := EncodeRelayReport(report)
	if err != nil {
		return false, nil, err
	}
	lggr.Infow("Built report", "intervalByOnRamp", intervalByOnRamp)
	return true, encodedReport, nil
}

func (r *RelayReportingPlugin) updateInflight(lggr logger.Logger, report *blob_verifier.CCIPRelayReport) {
	r.inFlightMu.Lock()
	defer r.inFlightMu.Unlock()
	// Reap any expired entries from inflight.
	for root, inFlightReport := range r.inFlight {
		if time.Since(inFlightReport.createdAt) > RelayMaxInflightTimeSeconds {
			lggr.Warnw("Inflight report expired, retrying")
			delete(r.inFlight, root)
		}
	}
	// Set new inflight ones as pending
	r.inFlight[report.RootOfRoots] = InflightReport{
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
	lggr.Infow("Accepting finalized report")
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

func (r *RelayReportingPlugin) isStaleReport(report *blob_verifier.CCIPRelayReport) bool {
	if isBlobVerifierDownNow(r.lggr, r.blobVerifier) {
		return true
	}
	for i, onRamp := range report.OnRamps {
		nextMin, err := r.nextMinSeqNumForOffRamp(onRamp)
		if err != nil {
			// Assume it's a transient issue getting the last report
			// Will try again on the next round
			return true
		}
		// If the next min is already greater than this reports min,
		// this report is stale.
		if nextMin > report.Intervals[i].Min {
			r.lggr.Warnw("report is stale", "onchain min", nextMin, "report min", report.Intervals[i].Min)
			return true
		}
	}
	return false
}

func (r *RelayReportingPlugin) Close() error {
	return nil
}
