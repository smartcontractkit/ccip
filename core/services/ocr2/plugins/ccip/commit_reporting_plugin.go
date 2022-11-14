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
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/merklemulti"
)

const MaxCommitReportLength = 1000 // TODO: Need to rethink this based on root of roots report.

var (
	_ types.ReportingPluginFactory = &CommitReportingPluginFactory{}
	_ types.ReportingPlugin        = &CommitReportingPlugin{}
)

// EncodeCommitReport abi encodes an offramp.CCIPCommitReport.
func EncodeCommitReport(commitReport *commit_store.CCIPCommitReport) (types.Report, error) {
	report, err := makeCommitReportArgs().PackValues([]interface{}{
		commitReport,
	})
	if err != nil {
		return nil, err
	}
	return report, nil
}

// DecodeCommitReport abi decodes a types.Report to an offramp.CCIPCommitReport
func DecodeCommitReport(report types.Report) (*commit_store.CCIPCommitReport, error) {
	unpacked, err := makeCommitReportArgs().Unpack(report)
	if err != nil {
		return nil, err
	}
	if len(unpacked) != 1 {
		return nil, errors.New("expected single struct value")
	}
	commitReport, ok := unpacked[0].(struct {
		OnRamps   []common.Address `json:"onRamps"`
		Intervals []struct {
			Min uint64 `json:"min"`
			Max uint64 `json:"max"`
		} `json:"intervals"`
		MerkleRoots [][32]byte `json:"merkleRoots"`
		RootOfRoots [32]byte   `json:"rootOfRoots"`
	})
	if !ok {
		return nil, errors.Errorf("invalid commit report got %T", unpacked[0])
	}
	var intervalsF []commit_store.CCIPInterval
	for i := range commitReport.Intervals {
		intervalsF = append(intervalsF, commit_store.CCIPInterval{
			Min: commitReport.Intervals[i].Min,
			Max: commitReport.Intervals[i].Max,
		})
	}
	return &commit_store.CCIPCommitReport{
		OnRamps:     commitReport.OnRamps,
		Intervals:   intervalsF,
		MerkleRoots: commitReport.MerkleRoots,
		RootOfRoots: commitReport.RootOfRoots,
	}, nil
}

func isCommitStoreDownNow(lggr logger.Logger, commitStore *commit_store.CommitStore) bool {
	paused, err := commitStore.Paused(nil)
	if err != nil {
		// Air on side of caution by halting if we cannot read the state?
		lggr.Errorw("Unable to read offramp paused", "err", err)
		return true
	}
	healthy, err := commitStore.IsAFNHealthy(nil)
	if err != nil {
		lggr.Errorw("Unable to read offramp afn", "err", err)
		return true
	}
	return paused || !healthy
}

type InflightReport struct {
	report    *commit_store.CCIPCommitReport
	createdAt time.Time
}

type CommitReportingPluginFactory struct {
	lggr                logger.Logger
	source              logpoller.LogPoller
	onRampSeqParsers    map[common.Address]func(log logpoller.Log) (uint64, error)
	onRampToReqEventSig map[common.Address]common.Hash
	onRamps             []common.Address
	commitStore         *commit_store.CommitStore
	onRampToHasher      map[common.Address]LeafHasher[[32]byte]
	inflightCacheExpiry time.Duration
}

// NewCommitReportingPluginFactory return a new CommitReportingPluginFactory.
func NewCommitReportingPluginFactory(
	lggr logger.Logger,
	source logpoller.LogPoller,
	commitStore *commit_store.CommitStore,
	onRampSeqParsers map[common.Address]func(log logpoller.Log) (uint64, error),
	onRampToReqEventSig map[common.Address]common.Hash,
	onRamps []common.Address,
	onRampToHasher map[common.Address]LeafHasher[[32]byte],
	inflightCacheExpiry time.Duration,
) types.ReportingPluginFactory {
	return &CommitReportingPluginFactory{lggr: lggr, commitStore: commitStore, onRampToReqEventSig: onRampToReqEventSig, onRampSeqParsers: onRampSeqParsers, onRamps: onRamps, source: source, onRampToHasher: onRampToHasher, inflightCacheExpiry: inflightCacheExpiry}
}

// NewReportingPlugin returns the ccip CommitReportingPlugin and satisfies the ReportingPluginFactory interface.
func (rf *CommitReportingPluginFactory) NewReportingPlugin(config types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	offchainConfig, err := Decode(config.OffchainConfig)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	return &CommitReportingPlugin{
			lggr:                rf.lggr.Named("CommitReportingPlugin"),
			F:                   config.F,
			source:              rf.source,
			onRampSeqParsers:    rf.onRampSeqParsers,
			onRampToReqEventSig: rf.onRampToReqEventSig,
			onRamps:             rf.onRamps,
			commitStore:         rf.commitStore,
			inFlight:            make(map[[32]byte]InflightReport),
			offchainConfig:      offchainConfig,
			onRampToHasher:      rf.onRampToHasher,
			inflightCacheExpiry: rf.inflightCacheExpiry,
		},
		types.ReportingPluginInfo{
			Name:          "CCIPCommit",
			UniqueReports: true,
			Limits: types.ReportingPluginLimits{
				MaxQueryLength:       0,
				MaxObservationLength: MaxObservationLength,
				MaxReportLength:      MaxCommitReportLength,
			},
		}, nil
}

type CommitReportingPlugin struct {
	lggr                logger.Logger
	F                   int
	source              logpoller.LogPoller
	onRamps             []common.Address
	onRampToReqEventSig map[common.Address]common.Hash
	onRampSeqParsers    map[common.Address]func(log logpoller.Log) (uint64, error)
	commitStore         *commit_store.CommitStore
	// We need to synchronize access to the inflight structure
	// as reporting plugin methods may be called from separate goroutines,
	// e.g. reporting vs transmission protocol.
	inFlightMu          sync.RWMutex
	inFlight            map[[32]byte]InflightReport
	offchainConfig      OffchainConfig
	onRampToHasher      map[common.Address]LeafHasher[[32]byte]
	inflightCacheExpiry time.Duration
}

func (r *CommitReportingPlugin) nextMinSeqNumForOffRamp(onRamp common.Address) (uint64, error) {
	return r.commitStore.GetExpectedNextSequenceNumber(nil, onRamp)
}

func (r *CommitReportingPlugin) nextMinSeqNumForInFlight(onRamp common.Address) uint64 {
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

func (r *CommitReportingPlugin) nextMinSeqNum(onRamp common.Address) (uint64, error) {
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

func (r *CommitReportingPlugin) Query(ctx context.Context, timestamp types.ReportTimestamp) (types.Query, error) {
	return types.Query{}, nil
}

func (r *CommitReportingPlugin) Observation(ctx context.Context, timestamp types.ReportTimestamp, query types.Query) (types.Observation, error) {
	lggr := r.lggr.Named("CommitObservation")
	if isCommitStoreDownNow(lggr, r.commitStore) {
		return nil, ErrCommitStoreIsDown
	}
	intervalsByOnRamp := make(map[common.Address]commit_store.CCIPInterval)
	for _, onRamp := range r.onRamps {
		nextMin, err := r.nextMinSeqNum(onRamp)
		if err != nil {
			return nil, err
		}
		// All available messages that have not been committed yet and have sufficient confirmations.
		lggr.Infof("Looking for requests with sig %s and nextMin %d on onRamp %s", r.onRampToReqEventSig[onRamp].Hex(), nextMin, onRamp.Hex())
		reqs, err := r.source.LogsDataWordGreaterThan(r.onRampToReqEventSig[onRamp], onRamp, SendRequestedSequenceNumberIndex, EvmWord(nextMin), int(r.offchainConfig.SourceIncomingConfirmations))
		if err != nil {
			return nil, err
		}
		lggr.Infof("%d requests found for onRamp %s", len(reqs), onRamp.Hex())
		if len(reqs) == 0 {
			lggr.Infow("no requests", "onRamp", onRamp)
			continue
		}
		var seqNrs []uint64
		for _, req := range reqs {
			seqNr, err := r.onRampSeqParsers[onRamp](req)
			if err != nil {
				lggr.Errorw("error parsing seq num", "err", err)
				continue
			}
			seqNrs = append(seqNrs, seqNr)
		}
		min := seqNrs[0]
		max := seqNrs[len(seqNrs)-1]
		if !contiguousReqs(lggr, min, max, seqNrs) {
			return nil, errors.New("unexpected gap in seq nums")
		}
		intervalsByOnRamp[onRamp] = commit_store.CCIPInterval{
			Min: min,
			Max: max,
		}
		lggr.Infof("OnRamp %v: min %v max %v", onRamp, min, max)
	}
	if len(intervalsByOnRamp) == 0 {
		lggr.Infow("No observations")
		return []byte{}, nil
	}
	return CommitObservation{
		IntervalsByOnRamp: intervalsByOnRamp,
	}.Marshal()
}

// buildReport assumes there is at least one message in reqs.
func (r *CommitReportingPlugin) buildReport(intervalByOnRamp map[common.Address]commit_store.CCIPInterval) (*commit_store.CCIPCommitReport, error) {
	lggr := r.lggr.Named("BuildReport")
	leafsByOnRamp, err := leafsFromIntervals(lggr, r.onRampToReqEventSig, r.onRampSeqParsers, intervalByOnRamp, r.source, r.onRampToHasher, int(r.offchainConfig.SourceIncomingConfirmations))
	if err != nil {
		return nil, err
	}
	// Produce a root for each onramp, then a root of roots.
	var (
		onRamps   []common.Address
		roots     [][32]byte
		intervals []commit_store.CCIPInterval
	)
	mctx := hasher.NewKeccakCtx()
	for onRamp, leaves := range leafsByOnRamp {
		if len(leaves) == 0 {
			lggr.Warnf("Tried building a tree without leaves for onRamp %s. %+v", onRamp.Hex(), leafsByOnRamp)
			continue
		}
		tree, err2 := merklemulti.NewTree(mctx, leaves)
		if err2 != nil {
			return nil, err2
		}
		roots = append(roots, tree.Root())
		onRamps = append(onRamps, onRamp)
		interval := intervalByOnRamp[onRamp]
		intervals = append(intervals, commit_store.CCIPInterval{
			Min: interval.Min,
			Max: interval.Max,
		})
	}
	if len(roots) == 0 {
		lggr.Warn("No valid roots found")
		return &commit_store.CCIPCommitReport{}, errors.New("No valid roots found")
	}
	// Make a root of roots
	outerTree, err := merklemulti.NewTree(mctx, roots)
	if err != nil {
		return nil, err
	}
	return &commit_store.CCIPCommitReport{
		MerkleRoots: roots,
		Intervals:   intervals,
		OnRamps:     onRamps,
		RootOfRoots: outerTree.Root(),
	}, nil
}

func (r *CommitReportingPlugin) Report(ctx context.Context, timestamp types.ReportTimestamp, query types.Query, observations []types.AttributedObservation) (bool, types.Report, error) {
	lggr := r.lggr.Named("Report")
	if isCommitStoreDownNow(lggr, r.commitStore) {
		return false, nil, ErrCommitStoreIsDown
	}
	nonEmptyObservations := getNonEmptyObservations[CommitObservation](lggr, observations)
	// Need at least F+1 valid observations
	if len(nonEmptyObservations) <= r.F {
		lggr.Debugf("Non-empty observations <= F, need at least F+1 to continue")
		return false, nil, nil
	}
	// Group intervals by onramp.
	intervalsByOnRamp := make(map[common.Address][]commit_store.CCIPInterval)
	for _, obs := range nonEmptyObservations {
		for onRamp, interval := range obs.IntervalsByOnRamp {
			intervalsByOnRamp[onRamp] = append(intervalsByOnRamp[onRamp], interval)
		}
	}
	intervalByOnRamp := make(map[common.Address]commit_store.CCIPInterval)
	for onRamp, intervals := range intervalsByOnRamp {
		if len(intervals) <= r.F {
			lggr.Debugf("Observations for OnRamp %s 1 < #obs <= F, need at least F+1 to continue", onRamp.Hex())
			continue
		}

		// We have at least F+1 valid observations for the given OnRamp
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
		intervalByOnRamp[onRamp] = commit_store.CCIPInterval{
			Min: minSeqNum,
			Max: maxSeqNum,
		}
	}
	report, err := r.buildReport(intervalByOnRamp)
	if err != nil {
		return false, nil, err
	}
	encodedReport, err := EncodeCommitReport(report)
	if err != nil {
		return false, nil, err
	}
	lggr.Infow("Built report", "intervalByOnRamp", intervalByOnRamp)
	return true, encodedReport, nil
}

func (r *CommitReportingPlugin) updateInflight(lggr logger.Logger, report *commit_store.CCIPCommitReport) {
	r.inFlightMu.Lock()
	defer r.inFlightMu.Unlock()
	// Reap any expired entries from inflight.
	for root, inFlightReport := range r.inFlight {
		if time.Since(inFlightReport.createdAt) > r.inflightCacheExpiry {
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

func (r *CommitReportingPlugin) ShouldAcceptFinalizedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	lggr := r.lggr.Named("ShouldAcceptFinalizedReport")
	parsedReport, err := DecodeCommitReport(report)
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

func (r *CommitReportingPlugin) ShouldTransmitAcceptedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	parsedReport, err := DecodeCommitReport(report)
	if err != nil {
		return false, nil
	}
	// If report is not stale we transmit.
	// When the commitTransmitter enqueues the tx for bptxm,
	// we mark it as fulfilled, effectively removing it from the set of inflight messages.
	return !r.isStaleReport(parsedReport), nil
}

func (r *CommitReportingPlugin) isStaleReport(report *commit_store.CCIPCommitReport) bool {
	if isCommitStoreDownNow(r.lggr, r.commitStore) {
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

func (r *CommitReportingPlugin) Close() error {
	return nil
}
