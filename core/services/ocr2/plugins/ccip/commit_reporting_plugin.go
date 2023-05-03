package ccip

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"
	"golang.org/x/exp/slices"

	txmgrtypes "github.com/smartcontractkit/chainlink/v2/common/txmgr/types"
	"github.com/smartcontractkit/chainlink/v2/core/assets"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/merklemulti"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

const MaxCommitReportLength = 1000

var (
	_ types.ReportingPluginFactory = &CommitReportingPluginFactory{}
	_ types.ReportingPlugin        = &CommitReportingPlugin{}
)

// CommitReportToEthTxMeta generates a txmgr.EthTxMeta from the given commit report.
// sequence numbers of the committed messages will be added to tx metadata
func CommitReportToEthTxMeta(report []byte) (*txmgr.EthTxMeta, error) {
	commitReport, err := DecodeCommitReport(report)
	if err != nil {
		return nil, err
	}
	n := int(commitReport.Interval.Max-commitReport.Interval.Min) + 1
	seqRange := make([]uint64, n)
	for i := 0; i < n; i++ {
		seqRange[i] = uint64(i) + commitReport.Interval.Min
	}
	return &txmgr.EthTxMeta{
		SeqNumbers: seqRange,
	}, nil
}

// EncodeCommitReport abi encodes an offramp.InternalCommitReport.
func EncodeCommitReport(commitReport *commit_store.CommitStoreCommitReport) (types.Report, error) {
	report, err := makeCommitReportArgs().PackValues([]interface{}{
		commitReport,
	})
	if err != nil {
		return nil, err
	}
	return report, nil
}

// DecodeCommitReport abi decodes a types.Report to an CommitStoreCommitReport
func DecodeCommitReport(report types.Report) (*commit_store.CommitStoreCommitReport, error) {
	unpacked, err := makeCommitReportArgs().Unpack(report)
	if err != nil {
		return nil, err
	}
	if len(unpacked) != 1 {
		return nil, errors.New("expected single struct value")
	}

	commitReport, ok := unpacked[0].(struct {
		PriceUpdates struct {
			TokenPriceUpdates []struct {
				SourceToken common.Address `json:"sourceToken"`
				UsdPerToken *big.Int       `json:"usdPerToken"`
			} `json:"tokenPriceUpdates"`
			DestChainId   uint64   `json:"destChainId"`
			UsdPerUnitGas *big.Int `json:"usdPerUnitGas"`
		} `json:"priceUpdates"`
		Interval struct {
			Min uint64 `json:"min"`
			Max uint64 `json:"max"`
		} `json:"interval"`
		MerkleRoot [32]byte `json:"merkleRoot"`
	})
	if !ok {
		return nil, errors.Errorf("invalid commit report got %T", unpacked[0])
	}

	var tokenPriceUpdates []commit_store.InternalTokenPriceUpdate
	for _, u := range commitReport.PriceUpdates.TokenPriceUpdates {
		tokenPriceUpdates = append(tokenPriceUpdates, commit_store.InternalTokenPriceUpdate{
			SourceToken: u.SourceToken,
			UsdPerToken: u.UsdPerToken,
		})
	}

	return &commit_store.CommitStoreCommitReport{
		PriceUpdates: commit_store.InternalPriceUpdates{
			DestChainId:       commitReport.PriceUpdates.DestChainId,
			UsdPerUnitGas:     commitReport.PriceUpdates.UsdPerUnitGas,
			TokenPriceUpdates: tokenPriceUpdates,
		},
		Interval: commit_store.CommitStoreInterval{
			Min: commitReport.Interval.Min,
			Max: commitReport.Interval.Max,
		},
		MerkleRoot: commitReport.MerkleRoot,
	}, nil
}

type InflightReport struct {
	report    *commit_store.CommitStoreCommitReport
	createdAt time.Time
}

type InflightPriceUpdate struct {
	priceUpdates commit_store.InternalPriceUpdates
	createdAt    time.Time
}

type CommitPluginConfig struct {
	lggr               logger.Logger
	sourceLP, destLP   logpoller.LogPoller
	onRamp             *evm_2_evm_onramp.EVM2EVMOnRamp
	priceRegistry      *price_registry.PriceRegistry
	priceGetter        PriceGetter
	sourceNative       common.Address
	sourceFeeEstimator txmgrtypes.FeeEstimator[*evmtypes.Head, gas.EvmFee, *assets.Wei, evmtypes.TxHash]
	sourceChainID      uint64
	commitStore        *commit_store.CommitStore
	hasher             LeafHasherInterface[[32]byte]
}

type CommitReportingPluginFactory struct {
	config CommitPluginConfig
}

// NewCommitReportingPluginFactory return a new CommitReportingPluginFactory.
func NewCommitReportingPluginFactory(config CommitPluginConfig) types.ReportingPluginFactory {
	return &CommitReportingPluginFactory{config: config}
}

// NewReportingPlugin returns the ccip CommitReportingPlugin and satisfies the ReportingPluginFactory interface.
func (rf *CommitReportingPluginFactory) NewReportingPlugin(config types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	offchainConfig, err := DecodeOffchainConfig[CommitOffchainConfig](config.OffchainConfig)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}

	onchainConfig, err := DecodeAbiStruct[CommitOnchainConfig](config.OnchainConfig)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}

	rf.config.lggr.Infow("Starting commit plugin", "offchainConfig", offchainConfig)

	return &CommitReportingPlugin{
			config:         rf.config,
			F:              config.F,
			inFlight:       make(map[[32]byte]InflightReport),
			onchainConfig:  onchainConfig,
			offchainConfig: offchainConfig,
		},
		types.ReportingPluginInfo{
			Name:          "CCIPCommit",
			UniqueReports: true,
			Limits: types.ReportingPluginLimits{
				MaxQueryLength:       MaxQueryLength,
				MaxObservationLength: MaxObservationLength,
				MaxReportLength:      MaxCommitReportLength,
			},
		}, nil
}

type CommitReportingPlugin struct {
	config CommitPluginConfig
	F      int
	// We need to synchronize access to the inflight structure
	// as reporting plugin methods may be called from separate goroutines,
	// e.g. reporting vs transmission protocol.
	inFlightMu           sync.RWMutex
	inFlight             map[[32]byte]InflightReport
	inFlightPriceUpdates []InflightPriceUpdate
	offchainConfig       CommitOffchainConfig
	onchainConfig        CommitOnchainConfig
}

func (r *CommitReportingPlugin) seqParser(log logpoller.Log) (uint64, error) {
	req, err := r.config.onRamp.ParseCCIPSendRequested(log.GetGethLog())
	if err != nil {
		r.config.lggr.Warnf("failed to parse log: %+v", log)
		return 0, err
	}
	return req.Message.SequenceNumber, nil
}

func (r *CommitReportingPlugin) nextMinSeqNum(ctx context.Context) (uint64, error) {
	nextMin, err := r.config.commitStore.GetExpectedNextSequenceNumber(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, err
	}
	r.inFlightMu.RLock()
	for _, report := range r.inFlight {
		if report.report.Interval.Max >= nextMin {
			nextMin = report.report.Interval.Max + 1
		}
	}
	r.inFlightMu.RUnlock()
	return nextMin, nil
}

func (r *CommitReportingPlugin) Query(context.Context, types.ReportTimestamp) (types.Query, error) {
	return types.Query{}, nil
}

func calculateUsdPerUnitGas(sourceGasPrice *big.Int, usdPerFeeCoin *big.Int) *big.Int {
	// (wei / gas) * (usd / eth) * (1 eth / 1e18 wei)  = usd/gas
	tmp := big.NewInt(0).Mul(sourceGasPrice, usdPerFeeCoin)
	return tmp.Div(tmp, big.NewInt(1e18))
}

// deviation_parts_per_billion = ((x2 - x1) / x1) * 1e9
func (r *CommitReportingPlugin) deviates(x1, x2 *big.Int) bool {
	// if x1 == 0, deviates if x2 != x1, to avoid the relative division by 0 error
	if x1.BitLen() == 0 {
		return x1.Cmp(x2) != 0
	}
	diff := big.NewInt(0).Sub(x1, x2)
	diff.Mul(diff, big.NewInt(1e9))
	diff.Div(diff, x1)
	return diff.CmpAbs(big.NewInt(int64(r.offchainConfig.FeeUpdateDeviationPPB))) > 0
}

type update = struct {
	timestamp time.Time
	value     *big.Int
}

// latest gasPrice update is returned in addressZero (common.Address{}); the other keys are tokens price updates
func (r *CommitReportingPlugin) getLatestPriceUpdates(ctx context.Context, now time.Time, skipInflight bool) (latestUpdates map[common.Address]update, err error) {
	latestUpdates = make(map[common.Address]update)
	gasUpdatesWithinHeartBeat, err := r.config.destLP.IndexedLogsCreatedAfter(EventSignatures.UsdPerUnitGasUpdated, r.config.priceRegistry.Address(), 1, []common.Hash{EvmWord(r.config.sourceChainID)}, now.Add(-r.offchainConfig.FeeUpdateHeartBeat.Duration()), pg.WithParentCtx(ctx))
	if err != nil {
		return nil, err
	}
	for _, log := range gasUpdatesWithinHeartBeat {
		// Ordered by ascending timestamps
		priceUpdate, err2 := r.config.priceRegistry.ParseUsdPerUnitGasUpdated(log.GetGethLog())
		if err2 != nil {
			return nil, err2
		}
		timestamp := time.Unix(priceUpdate.Timestamp.Int64(), 0)
		if !timestamp.Before(latestUpdates[common.Address{}].timestamp) {
			latestUpdates[common.Address{}] = update{
				timestamp: timestamp,
				value:     priceUpdate.Value,
			}
		}
	}

	tokenUpdatesWithinHeartBeat, err := r.config.destLP.LogsCreatedAfter(EventSignatures.UsdPerTokenUpdated, r.config.priceRegistry.Address(), now.Add(-r.offchainConfig.FeeUpdateHeartBeat.Duration()), pg.WithParentCtx(ctx))
	if err != nil {
		return nil, err
	}
	for _, log := range tokenUpdatesWithinHeartBeat {
		// Ordered by ascending timestamps
		tokenUpdate, err := r.config.priceRegistry.ParseUsdPerTokenUpdated(log.GetGethLog())
		if err != nil {
			return nil, err
		}
		timestamp := time.Unix(tokenUpdate.Timestamp.Int64(), 0)
		if !timestamp.Before(latestUpdates[tokenUpdate.Token].timestamp) {
			latestUpdates[tokenUpdate.Token] = update{
				timestamp: timestamp,
				value:     tokenUpdate.Value,
			}
		}
	}
	if skipInflight {
		return latestUpdates, nil
	}

	r.inFlightMu.RLock()
	defer r.inFlightMu.RUnlock()
	for _, inflight := range r.inFlightPriceUpdates {
		if inflight.priceUpdates.DestChainId != 0 && !inflight.createdAt.Before(latestUpdates[common.Address{}].timestamp) {
			latestUpdates[common.Address{}] = update{
				timestamp: inflight.createdAt,
				value:     inflight.priceUpdates.UsdPerUnitGas,
			}
		}

		for _, inflightTokenUpdate := range inflight.priceUpdates.TokenPriceUpdates {
			if !inflight.createdAt.Before(latestUpdates[inflightTokenUpdate.SourceToken].timestamp) {
				latestUpdates[inflightTokenUpdate.SourceToken] = update{
					timestamp: inflight.createdAt,
					value:     inflightTokenUpdate.UsdPerToken,
				}
			}
		}
	}

	return latestUpdates, nil
}

// All prices are USD ($1=1e18) denominated. We only generate prices we think should be updated; otherwise, omitting values means voting to skip updating them
func (r *CommitReportingPlugin) generatePriceUpdates(ctx context.Context, now time.Time) (sourceGasPriceUSD *big.Int, tokenPricesUSD map[common.Address]*big.Int, err error) {
	// fetch feeTokens every observation, so we're automatically up-to-date if new feeTokens are added or removed
	feeTokens, err := r.config.priceRegistry.GetFeeTokens(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, nil, err
	}

	queryTokens := append([]common.Address{r.config.sourceNative}, feeTokens...)
	// Include wrapped native in our token query as way to identify the source native USD price.
	// notice USD is in 1e18 scale, i.e. $1 = 1e18
	tokenPricesUSD, err = r.config.priceGetter.TokenPricesUSD(ctx, queryTokens)
	if err != nil {
		return nil, nil, err
	}
	for _, token := range queryTokens {
		if tokenPricesUSD[token] == nil {
			return nil, nil, errors.Errorf("missing token price: %+v", token)
		}
	}
	sourceNativePriceUSD := tokenPricesUSD[r.config.sourceNative]
	for token := range tokenPricesUSD {
		if !slices.Contains(feeTokens, token) {
			// clean tokenPricesUSD of any address which isn't a feeToken, including sourceNative
			delete(tokenPricesUSD, token)
		}
	}

	// Observe a source chain price for pricing.
	sourceGasPriceWei, _, err := r.config.sourceFeeEstimator.GetFee(ctx, nil, 0, assets.NewWei(big.NewInt(int64(r.offchainConfig.MaxGasPrice))))
	if err != nil {
		return nil, nil, err
	}
	// Use legacy if no dynamic is available.
	gasPrice := sourceGasPriceWei.Legacy.ToInt()
	if sourceGasPriceWei.Dynamic != nil && sourceGasPriceWei.Dynamic.FeeCap != nil {
		gasPrice = sourceGasPriceWei.Dynamic.FeeCap.ToInt()
	}
	if gasPrice == nil {
		return nil, nil, fmt.Errorf("missing gas price %+v", sourceGasPriceWei)
	}

	sourceGasPriceUSD = calculateUsdPerUnitGas(gasPrice, sourceNativePriceUSD)

	latestUpdates, err := r.getLatestPriceUpdates(ctx, now, false)
	if err != nil {
		return nil, nil, err
	}

	if gasUpdate := latestUpdates[common.Address{}]; gasUpdate.value != nil && now.Sub(gasUpdate.timestamp) < r.offchainConfig.FeeUpdateHeartBeat.Duration() && !r.deviates(sourceGasPriceUSD, gasUpdate.value) {
		// vote skip gasPrice update by leaving it nil
		sourceGasPriceUSD = nil
	}

	for token, price := range tokenPricesUSD {
		if tokenUpdate := latestUpdates[token]; tokenUpdate.value != nil && now.Sub(tokenUpdate.timestamp) < r.offchainConfig.FeeUpdateHeartBeat.Duration() && !r.deviates(price, tokenUpdate.value) {
			// vote skip tokenPrice update by not including it in price map
			delete(tokenPricesUSD, token)
		}
	}

	// either may be empty
	return sourceGasPriceUSD, tokenPricesUSD, nil
}

func (r *CommitReportingPlugin) Observation(ctx context.Context, _ types.ReportTimestamp, _ types.Query) (types.Observation, error) {
	lggr := r.config.lggr.Named("CommitObservation")
	if isCommitStoreDownNow(ctx, lggr, r.config.commitStore) {
		return nil, ErrCommitStoreIsDown
	}
	r.expireInflight(lggr)

	// Will return 0,0 if no messages are found. This is a valid case as the report could
	// still contain fee updates.
	min, max, err := r.calculateMinMaxSequenceNumbers(ctx, lggr)
	if err != nil {
		return nil, err
	}

	sourceGasPriceUSD, tokenPricesUSD, err := r.generatePriceUpdates(ctx, time.Now())
	if err != nil {
		return nil, err
	}

	return CommitObservation{
		Interval: commit_store.CommitStoreInterval{
			Min: min,
			Max: max,
		},
		TokenPricesUSD:    tokenPricesUSD,
		SourceGasPriceUSD: sourceGasPriceUSD,
	}.Marshal()
}

func (r *CommitReportingPlugin) calculateMinMaxSequenceNumbers(ctx context.Context, lggr logger.Logger) (uint64, uint64, error) {
	nextMin, err := r.nextMinSeqNum(ctx)
	if err != nil {
		return 0, 0, err
	}
	// All available messages that have not been committed yet and have sufficient confirmations.
	lggr.Infof("Looking for requests with sig %v and nextMin %d on onRampAddr %v", EventSignatures.SendRequested, nextMin, r.config.onRamp.Address())
	reqs, err := r.config.sourceLP.LogsDataWordGreaterThan(
		EventSignatures.SendRequested,
		r.config.onRamp.Address(),
		EventSignatures.SendRequestedSequenceNumberWord,
		EvmWord(nextMin),
		int(r.offchainConfig.SourceIncomingConfirmations),
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return 0, 0, err
	}
	lggr.Infof("%d requests found for onRampAddr %v", len(reqs), r.config.onRamp.Address())
	if len(reqs) == 0 {
		return 0, 0, nil
	}
	var seqNrs []uint64
	for _, req := range reqs {
		seqNr, err2 := r.seqParser(req)
		if err2 != nil {
			lggr.Errorw("error parsing seq num", "err", err2)
			continue
		}
		seqNrs = append(seqNrs, seqNr)
	}
	min := seqNrs[0]
	max := seqNrs[len(seqNrs)-1]
	if min != nextMin {
		// Still report the observation as even partial reports have value e.g. all nodes are
		// missing a single, different log each, they would still be able to produce a valid report.
		lggr.Warnf("Missing sequence number range [%d-%d] for onRamp %v", nextMin, min, r.config.onRamp.Address())
	}
	if !contiguousReqs(lggr, min, max, seqNrs) {
		return 0, 0, errors.New("unexpected gap in seq nums")
	}
	lggr.Infof("OnRamp %v: min %v max %v", r.config.onRamp.Address(), min, max)
	return min, max, nil
}

// buildReport assumes there is at least one message in reqs.
func (r *CommitReportingPlugin) buildReport(ctx context.Context, interval commit_store.CommitStoreInterval, priceUpdates commit_store.InternalPriceUpdates) (*commit_store.CommitStoreCommitReport, error) {
	lggr := r.config.lggr.Named("BuildReport")

	// If no messages are needed only include fee updates
	if interval.Min == 0 {
		return &commit_store.CommitStoreCommitReport{
			PriceUpdates: priceUpdates,
			MerkleRoot:   [32]byte{},
			Interval:     interval,
		}, nil
	}

	// Logs are guaranteed to be in order of seq num, since these are finalized logs only
	// and the contract's seq num is auto-incrementing.
	logs, err := r.config.sourceLP.LogsDataWordRange(
		EventSignatures.SendRequested,
		r.config.onRamp.Address(),
		EventSignatures.SendRequestedSequenceNumberWord,
		logpoller.EvmWord(interval.Min),
		logpoller.EvmWord(interval.Max),
		int(r.offchainConfig.SourceIncomingConfirmations),
		pg.WithParentCtx(ctx))
	if err != nil {
		return nil, err
	}
	leaves, err := leavesFromIntervals(lggr, r.seqParser, interval, r.config.hasher, logs)
	if err != nil {
		return nil, err
	}

	if len(leaves) == 0 {
		return nil, fmt.Errorf("tried building a tree without leaves for onRampAddr %v. %+v", r.config.onRamp.Address(), leaves)
	}
	tree, err := merklemulti.NewTree(hasher.NewKeccakCtx(), leaves)
	if err != nil {
		return nil, err
	}

	return &commit_store.CommitStoreCommitReport{
		PriceUpdates: priceUpdates,
		MerkleRoot:   tree.Root(),
		Interval:     interval,
	}, nil
}

func (r *CommitReportingPlugin) Report(ctx context.Context, _ types.ReportTimestamp, _ types.Query, observations []types.AttributedObservation) (bool, types.Report, error) {
	lggr := r.config.lggr.Named("Report")
	nonEmptyObservations := getNonEmptyObservations[CommitObservation](lggr, observations)
	var intervals []commit_store.CommitStoreInterval
	for _, obs := range nonEmptyObservations {
		intervals = append(intervals, obs.Interval)
	}

	agreedInterval, err := calculateIntervalConsensus(intervals, r.F)
	if err != nil {
		return false, nil, err
	}

	priceUpdates := calculatePriceUpdates(r.config.sourceChainID, nonEmptyObservations, r.F)
	// If there are no fee updates and the interval is zero there is no report to produce.
	if len(priceUpdates.TokenPriceUpdates) == 0 && priceUpdates.DestChainId == 0 && agreedInterval.Min == 0 {
		return false, nil, nil
	}

	report, err := r.buildReport(ctx, agreedInterval, priceUpdates)
	if err != nil {
		return false, nil, err
	}
	encodedReport, err := EncodeCommitReport(report)
	if err != nil {
		return false, nil, err
	}
	lggr.Infow("Built report", "interval", agreedInterval)
	return true, encodedReport, nil
}

// Note priceUpdates must be deterministic.
func calculatePriceUpdates(destChainId uint64, observations []CommitObservation, f int) commit_store.InternalPriceUpdates {
	priceObservations := make(map[common.Address][]*big.Int)
	var sourceGasObservations []*big.Int

	for _, obs := range observations {
		if obs.SourceGasPriceUSD != nil {
			// Add only non-nil source gas price
			sourceGasObservations = append(sourceGasObservations, obs.SourceGasPriceUSD)
		}
		// iterate over any token which price is included in observations
		for token, price := range obs.TokenPricesUSD {
			if price == nil {
				continue
			}
			priceObservations[token] = append(priceObservations[token], price)
		}
	}

	var priceUpdates []commit_store.InternalTokenPriceUpdate
	for token, tokenPriceObservations := range priceObservations {
		// If majority report a token price, include it in the update
		if len(tokenPriceObservations) <= f {
			continue
		}
		medianPrice := median(tokenPriceObservations)
		priceUpdates = append(priceUpdates, commit_store.InternalTokenPriceUpdate{
			SourceToken: token,
			UsdPerToken: medianPrice,
		})
	}

	// Determinism required.
	sort.Slice(priceUpdates, func(i, j int) bool {
		return bytes.Compare(priceUpdates[i].SourceToken[:], priceUpdates[j].SourceToken[:]) == -1
	})

	// Must never be nil
	usdPerUnitGas := big.NewInt(0)
	// If majority report a gas price, include it in the update
	if len(sourceGasObservations) <= f {
		destChainId = 0
	} else {
		usdPerUnitGas = median(sourceGasObservations)
	}

	return commit_store.InternalPriceUpdates{
		TokenPriceUpdates: priceUpdates,
		// Sending zero is ok, UsdPerUnitGas update is skipped
		DestChainId:   destChainId,
		UsdPerUnitGas: usdPerUnitGas,
	}
}

func median(vals []*big.Int) *big.Int {
	valsCopy := make([]*big.Int, len(vals))
	copy(valsCopy[:], vals[:])
	sort.Slice(valsCopy, func(i, j int) bool {
		return valsCopy[i].Cmp(valsCopy[j]) == -1
	})
	return valsCopy[len(valsCopy)/2]
}

// Assumed at least f+1 valid observations
func calculateIntervalConsensus(intervals []commit_store.CommitStoreInterval, f int) (commit_store.CommitStoreInterval, error) {
	if len(intervals) <= f {
		return commit_store.CommitStoreInterval{}, errors.Errorf("Not enough intervals to form consensus: #obs=%d, f=%d", len(intervals), f)
	}
	// Extract the min and max
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Min < intervals[j].Min
	})
	minSeqNum := intervals[f].Min

	// The only way a report could have a minSeqNum of 0 is when there are no messages to report
	// and the report is potentially still valid for gas fee updates.
	if minSeqNum == 0 {
		return commit_store.CommitStoreInterval{Min: 0, Max: 0}, nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Max < intervals[j].Max
	})
	// We use a conservative maximum. If we pick a value that some honest oracles might not
	// have seen they’ll end up not agreeing on a msg, stalling the protocol.
	maxSeqNum := intervals[f].Max
	// TODO: Do we for sure want to fail everything here?
	if maxSeqNum < minSeqNum {
		return commit_store.CommitStoreInterval{}, errors.New("max seq num smaller than min")
	}

	return commit_store.CommitStoreInterval{
		Min: minSeqNum,
		Max: maxSeqNum,
	}, nil
}

func (r *CommitReportingPlugin) expireInflight(lggr logger.Logger) {
	r.inFlightMu.Lock()
	defer r.inFlightMu.Unlock()
	// Reap any expired entries from inflight.
	for root, inFlightReport := range r.inFlight {
		if time.Since(inFlightReport.createdAt) > r.offchainConfig.InflightCacheExpiry.Duration() {
			// Happy path: inflight report was successfully transmitted onchain, we remove it from inflight and onchain state reflects inflight.
			// Sad path: inflight report reverts onchain, we remove it from inflight, onchain state does not reflect the chains so we retry.
			lggr.Infow("Inflight report expired", "rootOfRoots", hexutil.Encode(inFlightReport.report.MerkleRoot[:]))
			delete(r.inFlight, root)
		}
	}
	var stillInflight []InflightPriceUpdate
	for _, inFlightFeeUpdate := range r.inFlightPriceUpdates {
		if time.Since(inFlightFeeUpdate.createdAt) > r.offchainConfig.InflightCacheExpiry.Duration() {
			// Happy path: inflight report was successfully transmitted onchain, we remove it from inflight and onchain state reflects inflight.
			// Sad path: inflight report reverts onchain, we remove it from inflight, onchain state does not reflect the chains so we retry.
			lggr.Infow("Inflight price update expired", "updates", inFlightFeeUpdate.priceUpdates)
			stillInflight = append(stillInflight, inFlightFeeUpdate)
		}
	}
	r.inFlightPriceUpdates = stillInflight
}

func (r *CommitReportingPlugin) addToInflight(lggr logger.Logger, report *commit_store.CommitStoreCommitReport) {
	r.inFlightMu.Lock()
	defer r.inFlightMu.Unlock()

	if report.MerkleRoot != [32]byte{} {
		// Set new inflight ones as pending
		lggr.Infow("Adding to inflight report", "rootOfRoots", hexutil.Encode(report.MerkleRoot[:]))
		r.inFlight[report.MerkleRoot] = InflightReport{
			report:    report,
			createdAt: time.Now(),
		}
	}

	if report.PriceUpdates.DestChainId != 0 || len(report.PriceUpdates.TokenPriceUpdates) != 0 {
		lggr.Infow("Adding to inflight fee updates", "priceUpdates", report.PriceUpdates)
		r.inFlightPriceUpdates = append(r.inFlightPriceUpdates, InflightPriceUpdate{
			priceUpdates: report.PriceUpdates,
			createdAt:    time.Now(),
		})
	}
}

func (r *CommitReportingPlugin) ShouldAcceptFinalizedReport(ctx context.Context, _ types.ReportTimestamp, report types.Report) (bool, error) {
	lggr := r.config.lggr.Named("ShouldAcceptFinalizedReport")
	parsedReport, err := DecodeCommitReport(report)
	if err != nil {
		return false, err
	}
	if parsedReport.MerkleRoot == [32]byte{} && parsedReport.PriceUpdates.DestChainId == 0 && len(parsedReport.PriceUpdates.TokenPriceUpdates) == 0 {
		// Empty report, should not be put on chain
		return false, nil
	}

	if parsedReport.MerkleRoot != [32]byte{} {
		// Note it's ok to leave the unstarted requests behind, since the
		// 'Observe' is always based on the last reports onchain min seq num.
		// This is a stricter isStaleReport, which considers inFlight requests and accepts only
		// reports starting at nextMinSeqNum
		nextInflightMin, err := r.nextMinSeqNum(ctx)
		if err != nil {
			return false, err
		}
		if nextInflightMin != parsedReport.Interval.Min {
			// There are sequence numbers missing between the commitStore/inflight txs and the proposed report.
			// The report will fail onchain unless the inflight cache is in an incorrect state. A state like this
			// could happen for various reasons, e.g. a reboot of the node emptying the caches, and should be self-healing.
			// We do not submit a tx and wait for the protocol to self-heal by updating the caches or invalidating
			// inflight caches over time.
			lggr.Errorw("Next inflight min is not equal to the proposed min of the report", "nextInflightMin", nextInflightMin, "proposed min", parsedReport.Interval.Min)
			return false, errors.New("Next inflight min is not equal to the proposed min of the report")
		}
	}

	r.addToInflight(lggr, parsedReport)
	lggr.Infow("Accepting finalized report", "merkleRoot", hexutil.Encode(parsedReport.MerkleRoot[:]))
	return true, nil
}

func (r *CommitReportingPlugin) ShouldTransmitAcceptedReport(ctx context.Context, _ types.ReportTimestamp, report types.Report) (bool, error) {
	parsedReport, err := DecodeCommitReport(report)
	if err != nil {
		return false, err
	}
	// If report is not stale we transmit.
	// When the commitTransmitter enqueues the tx for tx manager,
	// we mark it as fulfilled, effectively removing it from the set of inflight messages.
	return !r.isStaleReport(ctx, parsedReport), nil
}

func (r *CommitReportingPlugin) isStaleReport(ctx context.Context, report *commit_store.CommitStoreCommitReport) bool {
	if report.MerkleRoot != [32]byte{} {
		nextMin, err := r.config.commitStore.GetExpectedNextSequenceNumber(&bind.CallOpts{Context: ctx})
		if err != nil {
			// Assume it's a transient issue getting the last report
			// Will try again on the next round
			return true
		}
		// If the next min is already greater than this reports min,
		// this report is stale.
		if nextMin > report.Interval.Min {
			r.config.lggr.Infow("report is stale", "onchain min", nextMin, "report min", report.Interval.Min)
			return true
		}
		return false
	}

	// getting the last price updates without including inflight is like querying
	// current prices onchain, but uses logpoller's data to save on the RPC requests
	latestUpdates, err := r.getLatestPriceUpdates(ctx, time.Now(), true)
	if err != nil {
		return true
	}

	if report.PriceUpdates.DestChainId != 0 {
		if latestUpdate, ok := latestUpdates[common.Address{}]; ok && latestUpdate.value.Cmp(report.PriceUpdates.UsdPerUnitGas) == 0 {
			r.config.lggr.Infow("gasPriceUpdate-only report is stale", "latest gasPrice", latestUpdate.value, "destChainID", report.PriceUpdates.DestChainId)
			return true
		}
	} else if len(report.PriceUpdates.TokenPriceUpdates) > 0 { // check first token
		tokenUpdate := report.PriceUpdates.TokenPriceUpdates[0]
		if latestUpdate, ok := latestUpdates[tokenUpdate.SourceToken]; ok && latestUpdate.value.Cmp(tokenUpdate.UsdPerToken) == 0 {
			r.config.lggr.Infow("tokenPriceUpdate-only report is stale", "latest tokenPrice", latestUpdate.value, "token", tokenUpdate.SourceToken)
			return true
		}
	} else {
		return true
	}
	return false
}

func (r *CommitReportingPlugin) Close() error {
	return nil
}
