package ccip

import (
	"context"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

// EncodeCommitReport abi encodes an offramp.InternalCommitReport.
func EncodeCommitReport(commitReport *commit_store.InternalCommitReport) (types.Report, error) {
	report, err := makeCommitReportArgs().PackValues([]interface{}{
		commitReport,
	})
	if err != nil {
		return nil, err
	}
	return report, nil
}

// DecodeCommitReport abi decodes a types.Report to an offramp.InternalCommitReport
func DecodeCommitReport(report types.Report) (*commit_store.InternalCommitReport, error) {
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
	var intervalsF []commit_store.InternalInterval
	for i := range commitReport.Intervals {
		intervalsF = append(intervalsF, commit_store.InternalInterval{
			Min: commitReport.Intervals[i].Min,
			Max: commitReport.Intervals[i].Max,
		})
	}
	return &commit_store.InternalCommitReport{
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
	report    *commit_store.InternalCommitReport
	createdAt time.Time
}

type CommitReportingPluginFactory struct {
	lggr                logger.Logger
	source              logpoller.LogPoller
	onRampSeqParsers    map[common.Address]func(log logpoller.Log) (uint64, error)
	onRampToReqEventSig map[common.Address]EventSignatures
	onRamps             []common.Address
	commitStore         *commit_store.CommitStore
	onRampToHasher      map[common.Address]LeafHasherInterface[[32]byte]
	inflightCacheExpiry time.Duration
}

// NewCommitReportingPluginFactory return a new CommitReportingPluginFactory.
func NewCommitReportingPluginFactory(
	lggr logger.Logger,
	source logpoller.LogPoller,
	commitStore *commit_store.CommitStore,
	onRampSeqParsers map[common.Address]func(log logpoller.Log) (uint64, error),
	onRampToReqEventSig map[common.Address]EventSignatures,
	onRamps []common.Address,
	onRampToHasher map[common.Address]LeafHasherInterface[[32]byte],
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
	onRampToReqEventSig map[common.Address]EventSignatures
	onRampSeqParsers    map[common.Address]func(log logpoller.Log) (uint64, error)
	commitStore         *commit_store.CommitStore
	// We need to synchronize access to the inflight structure
	// as reporting plugin methods may be called from separate goroutines,
	// e.g. reporting vs transmission protocol.
	inFlightMu          sync.RWMutex
	inFlight            map[[32]byte]InflightReport
	offchainConfig      OffchainConfig
	onRampToHasher      map[common.Address]LeafHasherInterface[[32]byte]
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
	r.expireInflight(lggr)
	intervalsByOnRamp := make(map[common.Address]commit_store.InternalInterval)
	for _, onRamp := range r.onRamps {
		nextMin, err := r.nextMinSeqNum(onRamp)
		if err != nil {
			return nil, err
		}
		// All available messages that have not been committed yet and have sufficient confirmations.
		lggr.Infof("Looking for requests with sig %s and nextMin %d on onRampAddr %s", r.onRampToReqEventSig[onRamp].SendRequested.Hex(), nextMin, onRamp.Hex())
		reqs, err := r.source.LogsDataWordGreaterThan(r.onRampToReqEventSig[onRamp].SendRequested, onRamp, r.onRampToReqEventSig[onRamp].SendRequestedSequenceNumberIndex, EvmWord(nextMin), int(r.offchainConfig.SourceIncomingConfirmations))
		if err != nil {
			return nil, err
		}
		lggr.Infof("%d requests found for onRampAddr %s", len(reqs), onRamp.Hex())
		if len(reqs) == 0 {
			lggr.Infow("no requests", "onRampAddr", onRamp)
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
		if min != nextMin {
			// Still report the observation as even partial reports have value e.g. all nodes are
			// missing a single, different log each, they would still be able to produce a valid report.
			lggr.Warnf("Missing sequence number range [%d-%d] for onRamp %s", nextMin, min, onRamp.Hex())
		}
		if !contiguousReqs(lggr, min, max, seqNrs) {
			return nil, errors.New("unexpected gap in seq nums")
		}
		intervalsByOnRamp[onRamp] = commit_store.InternalInterval{
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
func (r *CommitReportingPlugin) buildReport(intervalByOnRamp map[common.Address]commit_store.InternalInterval) (*commit_store.InternalCommitReport, error) {
	lggr := r.lggr.Named("BuildReport")
	leafsByOnRamp, err := leafsFromIntervals(lggr, r.onRampToReqEventSig, r.onRampSeqParsers, intervalByOnRamp, r.source, r.onRampToHasher, int(r.offchainConfig.SourceIncomingConfirmations))
	if err != nil {
		return nil, err
	}
	// Produce a root for each onramp, then a root of roots.
	var (
		onRamps   []common.Address
		roots     [][32]byte
		intervals []commit_store.InternalInterval
	)
	mctx := hasher.NewKeccakCtx()
	for onRamp, leaves := range leafsByOnRamp {
		if len(leaves) == 0 {
			lggr.Warnf("Tried building a tree without leaves for onRampAddr %s. %+v", onRamp.Hex(), leafsByOnRamp)
			continue
		}
		tree, err2 := merklemulti.NewTree(mctx, leaves)
		if err2 != nil {
			return nil, err2
		}
		roots = append(roots, tree.Root())
		onRamps = append(onRamps, onRamp)
		interval := intervalByOnRamp[onRamp]
		intervals = append(intervals, commit_store.InternalInterval{
			Min: interval.Min,
			Max: interval.Max,
		})
	}
	if len(roots) == 0 {
		lggr.Warn("No valid roots found")
		return &commit_store.InternalCommitReport{}, errors.New("No valid roots found")
	}
	// Make a root of roots
	outerTree, err := merklemulti.NewTree(mctx, roots)
	if err != nil {
		return nil, err
	}
	return &commit_store.InternalCommitReport{
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
	intervalsByOnRamp := make(map[common.Address][]commit_store.InternalInterval)
	for _, obs := range nonEmptyObservations {
		for onRamp, interval := range obs.IntervalsByOnRamp {
			intervalsByOnRamp[onRamp] = append(intervalsByOnRamp[onRamp], interval)
		}
	}
	intervalByOnRamp := make(map[common.Address]commit_store.InternalInterval)
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
		intervalByOnRamp[onRamp] = commit_store.InternalInterval{
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

func (r *CommitReportingPlugin) expireInflight(lggr logger.Logger) {
	r.inFlightMu.Lock()
	defer r.inFlightMu.Unlock()
	// Reap any expired entries from inflight.
	for root, inFlightReport := range r.inFlight {
		if time.Since(inFlightReport.createdAt) > r.inflightCacheExpiry {
			// Happy path: inflight report was successfully transmitted onchain, we remove it from inflight and onchain state reflects inflight.
			// Sad path: inflight report reverts onchain, we remove it from inflight, onchain state does not reflect the chains so we retry.
			lggr.Infow("Inflight report expired", "rootOfRoots", hexutil.Encode(inFlightReport.report.RootOfRoots[:]))
			delete(r.inFlight, root)
		}
	}
}

func (r *CommitReportingPlugin) addToInflight(lggr logger.Logger, report *commit_store.InternalCommitReport) {
	r.inFlightMu.Lock()
	defer r.inFlightMu.Unlock()
	// Set new inflight ones as pending
	lggr.Infow("Adding to inflight report", "rootOfRoots", hexutil.Encode(report.RootOfRoots[:]))
	r.inFlight[report.RootOfRoots] = InflightReport{
		report:    report,
		createdAt: time.Now(),
	}
}

func (r *CommitReportingPlugin) ShouldAcceptFinalizedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	lggr := r.lggr.Named("ShouldAcceptFinalizedReport")
	parsedReport, err := DecodeCommitReport(report)
	if err != nil {
		return false, err
	}
	// Note it's ok to leave the unstarted requests behind, since the
	// 'Observe' is always based on the last reports onchain min seq num.
	if r.isStaleReport(parsedReport) {
		return false, nil
	}
	for i, onRamp := range parsedReport.OnRamps {
		nextInflightMin, err := r.nextMinSeqNum(onRamp)
		if err != nil {
			return false, err
		}
		if nextInflightMin != parsedReport.Intervals[i].Min {
			// There are sequence numbers missing between the commitStore/inflight txs and the proposed report.
			// The report will fail onchain unless the inflight cache is in an incorrect state. A state like this
			// could happen for various reasons, e.g. a reboot of the node emptying the caches, and should be self-healing.
			// We do not submit a tx and wait for the protocol to self-heal by updating the caches or invalidating
			// inflight caches over time.
			r.lggr.Errorw("Next inflight min is not equal to the proposed min of the report", "nextInflightMin", nextInflightMin, "proposed min", parsedReport.Intervals[i].Min, "onRamp", onRamp.Hex())
			return false, errors.New("Next inflight min is not equal to the proposed min of the report")
		}
	}
	r.addToInflight(lggr, parsedReport)
	lggr.Infow("Accepting finalized report", "rootOfRoots", hexutil.Encode(parsedReport.RootOfRoots[:]))
	return true, nil
}

func (r *CommitReportingPlugin) ShouldTransmitAcceptedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	parsedReport, err := DecodeCommitReport(report)
	if err != nil {
		return false, err
	}
	// If report is not stale we transmit.
	// When the commitTransmitter enqueues the tx for bptxm,
	// we mark it as fulfilled, effectively removing it from the set of inflight messages.
	return !r.isStaleReport(parsedReport), nil
}

func (r *CommitReportingPlugin) isStaleReport(report *commit_store.InternalCommitReport) bool {
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
