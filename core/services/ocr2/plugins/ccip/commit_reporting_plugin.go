package ccip

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink/v2/core/assets"
	evmclient "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipcalc"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/contractutil"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/logpollerutil"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/pricegetter"
	"github.com/smartcontractkit/chainlink/v2/core/utils/mathutil"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/cache"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/hashlib"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/merklemulti"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

const (
	// only dynamic field in CommitReport is tokens PriceUpdates, and we don't expect to need to update thousands of tokens in a single tx
	MaxCommitReportLength = 10_000
	// Maximum inflight seq number range before we consider reports to be failing to get included entirely
	// and restart from the chain's minSeqNum. Want to set it high to allow for large throughput,
	// but low enough to minimize wasted revert cost.
	MaxInflightSeqNumGap = 500
)

var (
	_ types.ReportingPluginFactory = &CommitReportingPluginFactory{}
	_ types.ReportingPlugin        = &CommitReportingPlugin{}
)

type update struct {
	timestamp time.Time
	value     *big.Int
}

type CommitPluginConfig struct {
	lggr logger.Logger
	// Source
	onRampReader        ccipdata.OnRampReader
	sourceChainSelector uint64
	sourceNative        common.Address
	sourceFeeEstimator  gas.EvmFeeEstimator
	// Dest
	destLP         logpoller.LogPoller
	destReader     ccipdata.Reader
	offRamp        evm_2_evm_offramp.EVM2EVMOffRampInterface
	commitStore    commit_store.CommitStoreInterface
	destClient     evmclient.Client
	destChainEVMID *big.Int
	// Offchain
	priceGetter pricegetter.PriceGetter
}

type CommitReportingPlugin struct {
	config             CommitPluginConfig
	F                  int
	lggr               logger.Logger
	inflightReports    *inflightCommitReportsContainer
	destPriceRegistry  price_registry.PriceRegistryInterface
	offchainConfig     ccipconfig.CommitOffchainConfig
	onchainConfig      ccipconfig.CommitOnchainConfig
	tokenDecimalsCache cache.AutoSync[map[common.Address]uint8]
}

type CommitReportingPluginFactory struct {
	config CommitPluginConfig

	// We keep track of the registered filters
	// TODO: Can push this down into the readers
	destChainFilters []logpoller.Filter
	filtersMu        *sync.Mutex
}

// NewCommitReportingPluginFactory return a new CommitReportingPluginFactory.
func NewCommitReportingPluginFactory(config CommitPluginConfig) *CommitReportingPluginFactory {
	return &CommitReportingPluginFactory{
		config:    config,
		filtersMu: &sync.Mutex{},
	}
}

// NewReportingPlugin returns the ccip CommitReportingPlugin and satisfies the ReportingPluginFactory interface.
func (rf *CommitReportingPluginFactory) NewReportingPlugin(config types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	onchainConfig, err := abihelpers.DecodeAbiStruct[ccipconfig.CommitOnchainConfig](config.OnchainConfig)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	offchainConfig, err := ccipconfig.DecodeOffchainConfig[ccipconfig.CommitOffchainConfig](config.OffchainConfig)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	destPriceRegistry, err := price_registry.NewPriceRegistry(onchainConfig.PriceRegistry, rf.config.destClient)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}

	if err = rf.UpdateLogPollerFilters(onchainConfig.PriceRegistry); err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}

	rf.config.lggr.Infow("NewReportingPlugin",
		"offchainConfig", offchainConfig,
		"onchainConfig", onchainConfig,
	)

	return &CommitReportingPlugin{
			config:            rf.config,
			F:                 config.F,
			lggr:              rf.config.lggr.Named("CommitReportingPlugin"),
			inflightReports:   newInflightCommitReportsContainer(offchainConfig.InflightCacheExpiry.Duration()),
			destPriceRegistry: destPriceRegistry,
			onchainConfig:     onchainConfig,
			offchainConfig:    offchainConfig,
			tokenDecimalsCache: cache.NewTokenToDecimals(
				rf.config.lggr,
				rf.config.destLP,
				rf.config.offRamp,
				destPriceRegistry,
				rf.config.destClient,
				int64(offchainConfig.DestFinalityDepth),
			),
		},
		types.ReportingPluginInfo{
			Name:          "CCIPCommit",
			UniqueReports: false, // See comment in CommitStore constructor.
			Limits: types.ReportingPluginLimits{
				MaxQueryLength:       MaxQueryLength,
				MaxObservationLength: MaxObservationLength,
				MaxReportLength:      MaxCommitReportLength,
			},
		}, nil
}

// Query is not used by the CCIP Commit plugin.
func (r *CommitReportingPlugin) Query(context.Context, types.ReportTimestamp) (types.Query, error) {
	return types.Query{}, nil
}

// Observation calculates the sequence number interval ready to be committed and
// the token and gas price updates required. A valid report could contain a merkle
// root and/or price updates.
func (r *CommitReportingPlugin) Observation(ctx context.Context, epochAndRound types.ReportTimestamp, _ types.Query) (types.Observation, error) {
	lggr := r.lggr.Named("CommitObservation")
	// If the commit store is down the protocol should halt.
	if contractutil.IsCommitStoreDownNow(ctx, lggr, r.config.commitStore) {
		return nil, ErrCommitStoreIsDown
	}
	r.inflightReports.expire(lggr)

	// Will return 0,0 if no messages are found. This is a valid case as the report could
	// still contain fee updates.
	min, max, err := r.calculateMinMaxSequenceNumbers(ctx, lggr)
	if err != nil {
		return nil, err
	}

	tokenDecimals, err := r.tokenDecimalsCache.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("get token decimals: %w", err)
	}

	sourceGasPriceUSD, tokenPricesUSD, err := r.generatePriceUpdates(ctx, lggr, tokenDecimals)
	if err != nil {
		return nil, err
	}

	lggr.Infow("Observation",
		"minSeqNr", min,
		"maxSeqNr", max,
		"sourceGasPriceUSD", sourceGasPriceUSD,
		"tokenPricesUSD", tokenPricesUSD,
		"epochAndRound", epochAndRound)
	// Even if all values are empty we still want to communicate our observation
	// with the other nodes, therefore, we always return the observed values.
	return CommitObservation{
		Interval: commit_store.CommitStoreInterval{
			Min: min,
			Max: max,
		},
		TokenPricesUSD:    tokenPricesUSD,
		SourceGasPriceUSD: sourceGasPriceUSD,
	}.Marshal()
}

// UpdateLogPollerFilters updates the log poller filters for the source and destination chains.
// pass zeroAddress if destPriceRegistry is unknown, filters with zero address are omitted.
// TODO: Should be able to Close and re-create readers to abstract filters.
func (rf *CommitReportingPluginFactory) UpdateLogPollerFilters(destPriceRegistry common.Address, qopts ...pg.QOpt) error {
	rf.filtersMu.Lock()
	defer rf.filtersMu.Unlock()

	// destination chain filters
	destFiltersBefore, destFiltersNow := rf.destChainFilters, getCommitPluginDestLpFilters(destPriceRegistry, rf.config.offRamp.Address())
	created, deleted := logpollerutil.FiltersDiff(destFiltersBefore, destFiltersNow)
	if err := logpollerutil.UnregisterLpFilters(rf.config.destLP, deleted, qopts...); err != nil {
		return err
	}
	if err := logpollerutil.RegisterLpFilters(rf.config.destLP, created, qopts...); err != nil {
		return err
	}
	rf.destChainFilters = destFiltersNow

	return nil
}

func (r *CommitReportingPlugin) calculateMinMaxSequenceNumbers(ctx context.Context, lggr logger.Logger) (uint64, uint64, error) {
	nextInflightMin, _, err := r.nextMinSeqNum(ctx, lggr)
	if err != nil {
		return 0, 0, err
	}

	msgRequests, err := r.config.onRampReader.GetSendRequestsGteSeqNum(ctx, nextInflightMin, int(r.offchainConfig.SourceFinalityDepth))
	if err != nil {
		return 0, 0, err
	}
	if len(msgRequests) == 0 {
		lggr.Infow("No new requests", "minSeqNr", nextInflightMin)
		return 0, 0, nil
	}
	seqNrs := make([]uint64, 0, len(msgRequests))
	for _, msgReq := range msgRequests {
		seqNrs = append(seqNrs, msgReq.Data.SequenceNumber)
	}

	minSeqNr := seqNrs[0]
	maxSeqNr := seqNrs[len(seqNrs)-1]
	if minSeqNr != nextInflightMin {
		// Still report the observation as even partial reports have value e.g. all nodes are
		// missing a single, different log each, they would still be able to produce a valid report.
		lggr.Warnf("Missing sequence number range [%d-%d]", nextInflightMin, minSeqNr)
	}
	if !ccipcalc.ContiguousReqs(lggr, minSeqNr, maxSeqNr, seqNrs) {
		return 0, 0, errors.New("unexpected gap in seq nums")
	}
	return minSeqNr, maxSeqNr, nil
}

func (r *CommitReportingPlugin) nextMinSeqNum(ctx context.Context, lggr logger.Logger) (inflightMin, onChainMin uint64, err error) {
	nextMinOnChain, err := r.config.commitStore.GetExpectedNextSequenceNumber(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, 0, err
	}
	// There are several scenarios to consider here for nextMin and inflight intervals.
	// 1. nextMin=2, inflight=[[2,3],[4,5]]. Node is waiting for [2,3] and [4,5] to be included, should return 6 to build on top.
	// 2. nextMin=2, inflight=[[4,5]]. [2,3] is expired but not yet visible onchain (means our cache expiry
	// was too low). In this case still want to return 6.
	// 3. nextMin=2, inflight=[] but other nodes have inflight=[2,3]. Say our node restarted and lost its cache. In this case
	// we still return the chain's nextMin, other oracles will ignore our observation. Other nodes however will build [4,5]
	// and then we'll add that to our cache in ShouldAcceptFinalizedReport, putting us into the previous position at which point
	// we can start contributing again.
	// 4. nextMin=4, inflight=[[2,3],[4,5]]. We see the onchain update, but haven't expired from our cache yet. Should happen
	// regularly and we just return 6.
	// 5. nextMin=2, inflight=[[4,5]]. [2,3] failed to get onchain for some reason. We'll return 6 and continue building even though
	// subsequent reports will revert, but eventually they will all expire OR we'll hit MaxInflightSeqNumGap and forcibly
	// expire them all. This scenario can also occur if there is a reorg which reorders the reports such that one reverts.
	maxInflight := r.inflightReports.maxInflightSeqNr()
	if (maxInflight > nextMinOnChain) && ((maxInflight - nextMinOnChain) > MaxInflightSeqNumGap) {
		r.inflightReports.reset(lggr)
		return nextMinOnChain, nextMinOnChain, nil
	}
	return mathutil.Max(nextMinOnChain, maxInflight+1), nextMinOnChain, nil
}

// All prices are USD ($1=1e18) denominated. We only generate prices we think should be updated;
// otherwise, omitting values means voting to skip updating them
func (r *CommitReportingPlugin) generatePriceUpdates(
	ctx context.Context,
	lggr logger.Logger,
	tokenDecimals map[common.Address]uint8,
) (sourceGasPriceUSD *big.Int, tokenPricesUSD map[common.Address]*big.Int, err error) {
	tokensWithDecimal := make([]common.Address, 0, len(tokenDecimals))
	for token := range tokenDecimals {
		tokensWithDecimal = append(tokensWithDecimal, token)
	}

	// Include wrapped native in our token query as way to identify the source native USD price.
	// notice USD is in 1e18 scale, i.e. $1 = 1e18
	queryTokens := append([]common.Address{r.config.sourceNative}, tokensWithDecimal...)
	sort.Slice(queryTokens, func(i, j int) bool { return queryTokens[i].String() < queryTokens[j].String() }) // make the query deterministic

	rawTokenPricesUSD, err := r.config.priceGetter.TokenPricesUSD(ctx, queryTokens)
	if err != nil {
		return nil, nil, err
	}
	lggr.Infow("Raw token prices", "rawTokenPrices", rawTokenPricesUSD)

	// make sure that we got prices for all the tokens of our query
	for _, token := range queryTokens {
		if rawTokenPricesUSD[token] == nil {
			return nil, nil, errors.Errorf("missing token price: %+v", token)
		}
	}

	sourceNativePriceUSD, exists := rawTokenPricesUSD[r.config.sourceNative]
	if !exists {
		return nil, nil, fmt.Errorf("missing source native (%s) price", r.config.sourceNative)
	}

	tokenPricesUSD = make(map[common.Address]*big.Int, len(rawTokenPricesUSD))
	for token := range rawTokenPricesUSD {
		decimals, exists := tokenDecimals[token]
		if !exists {
			// do not include any address which isn't a supported token on dest chain, including sourceNative
			lggr.Infow("Skipping token not supported on dest chain", "token", token)
			continue
		}
		tokenPricesUSD[token] = calculateUsdPer1e18TokenAmount(rawTokenPricesUSD[token], decimals)
	}

	// Observe a source chain price for pricing.
	sourceGasPriceWei, _, err := r.config.sourceFeeEstimator.GetFee(ctx, nil, 0, assets.NewWei(big.NewInt(int64(r.offchainConfig.MaxGasPrice))))
	if err != nil {
		return nil, nil, err
	}
	// Use legacy if no dynamic is available.
	gasPrice := sourceGasPriceWei.Legacy.ToInt()
	if sourceGasPriceWei.DynamicFeeCap != nil {
		gasPrice = sourceGasPriceWei.DynamicFeeCap.ToInt()
	}
	if gasPrice == nil {
		return nil, nil, fmt.Errorf("missing gas price %+v", sourceGasPriceWei)
	}
	sourceGasPriceUSD = ccipcalc.CalculateUsdPerUnitGas(gasPrice, sourceNativePriceUSD)

	lggr.Infow("Observing gas price", "observedGasPriceWei", gasPrice, "observedGasPriceUSD", sourceGasPriceUSD)
	lggr.Infow("Observing token prices", "tokenPrices", tokenPricesUSD, "sourceNativePriceUSD", sourceNativePriceUSD)
	return sourceGasPriceUSD, tokenPricesUSD, nil
}

// Input price is USD per full token, with 18 decimal precision
// Result price is USD per 1e18 of smallest token denomination, with 18 decimal precision
// Example: 1 USDC = 1.00 USD per full token, each full token is 6 decimals -> 1 * 1e18 * 1e18 / 1e6 = 1e30
func calculateUsdPer1e18TokenAmount(price *big.Int, decimals uint8) *big.Int {
	tmp := big.NewInt(0).Mul(price, big.NewInt(1e18))
	return tmp.Div(tmp, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
}

// Gets the latest token price updates based on logs within the heartbeat
// The updates returned by this function are guaranteed to not contain nil values.
func (r *CommitReportingPlugin) getLatestTokenPriceUpdates(ctx context.Context, now time.Time, checkInflight bool) (map[common.Address]update, error) {
	tokenPriceUpdates, err := r.config.destReader.GetTokenPriceUpdatesCreatedAfter(
		ctx,
		r.destPriceRegistry.Address(),
		now.Add(-r.offchainConfig.FeeUpdateHeartBeat.Duration()),
		0,
	)
	if err != nil {
		return nil, err
	}

	latestUpdates := make(map[common.Address]update)
	for _, tokenPriceUpdate := range tokenPriceUpdates {
		priceUpdate := tokenPriceUpdate.Data
		// Ordered by ascending timestamps
		timestamp := time.Unix(priceUpdate.Timestamp.Int64(), 0)
		if priceUpdate.Value != nil && !timestamp.Before(latestUpdates[priceUpdate.Token].timestamp) {
			latestUpdates[priceUpdate.Token] = update{
				timestamp: timestamp,
				value:     priceUpdate.Value,
			}
		}
	}
	if !checkInflight {
		return latestUpdates, nil
	}

	// todo this comparison is faulty, as a previously-sent update's onchain timestamp can be higher than inflight timestamp
	// to properly fix, need a solution to map from onchain request to offchain timestamp
	// leaving it as is, as token prices are updated infrequently, so this should not cause many issues
	latestInflightTokenPriceUpdates := r.inflightReports.latestInflightTokenPriceUpdates()
	for inflightToken, latestInflightUpdate := range latestInflightTokenPriceUpdates {
		if latestInflightUpdate.timestamp.After(latestUpdates[inflightToken].timestamp) && latestInflightUpdate.value != nil {
			latestUpdates[inflightToken] = latestInflightUpdate
		}
	}

	return latestUpdates, nil
}

// Gets the latest gas price updates based on logs within the heartbeat
func (r *CommitReportingPlugin) getLatestGasPriceUpdate(ctx context.Context, now time.Time, checkInflight bool) (gasPriceUpdate update, error error) {
	if checkInflight {
		latestInflightGasPriceUpdate := r.inflightReports.getLatestInflightGasPriceUpdate()
		if latestInflightGasPriceUpdate != nil && latestInflightGasPriceUpdate.timestamp.After(gasPriceUpdate.timestamp) {
			gasPriceUpdate = *latestInflightGasPriceUpdate
		}

		if gasPriceUpdate.value != nil {
			r.lggr.Infow("Latest gas price from inflight", "gasPriceUpdateVal", gasPriceUpdate.value, "gasPriceUpdateTs", gasPriceUpdate.timestamp)
			// Gas price can fluctuate frequently, many updates may be in flight.
			// If there is gas price update inflight, use it as source of truth, no need to check onchain.
			return gasPriceUpdate, nil
		}
	}

	// If there are no price updates inflight, check latest prices onchain
	gasPriceUpdates, err := r.config.destReader.GetGasPriceUpdatesCreatedAfter(
		ctx,
		r.destPriceRegistry.Address(),
		r.config.sourceChainSelector,
		now.Add(-r.offchainConfig.FeeUpdateHeartBeat.Duration()),
		0,
	)
	if err != nil {
		return update{}, err
	}

	for _, priceUpdate := range gasPriceUpdates {
		// Ordered by ascending timestamps
		timestamp := time.Unix(priceUpdate.Data.Timestamp.Int64(), 0)
		if !timestamp.Before(gasPriceUpdate.timestamp) {
			gasPriceUpdate = update{
				timestamp: timestamp,
				value:     priceUpdate.Data.Value,
			}
		}
	}

	if gasPriceUpdate.value != nil {
		r.lggr.Infow("Latest gas price from log poller", "gasPriceUpdateVal", gasPriceUpdate.value, "gasPriceUpdateTs", gasPriceUpdate.timestamp)
	}

	return gasPriceUpdate, nil
}

func (r *CommitReportingPlugin) Report(ctx context.Context, epochAndRound types.ReportTimestamp, _ types.Query, observations []types.AttributedObservation) (bool, types.Report, error) {
	now := time.Now()
	lggr := r.lggr.Named("CommitReport")
	parsableObservations := getParsableObservations[CommitObservation](lggr, observations)
	var intervals []commit_store.CommitStoreInterval
	for _, obs := range parsableObservations {
		intervals = append(intervals, obs.Interval)
	}

	agreedInterval, err := calculateIntervalConsensus(intervals, r.F, merklemulti.MaxNumberTreeLeaves)
	if err != nil {
		return false, nil, err
	}

	latestGasPrice, err := r.getLatestGasPriceUpdate(ctx, now, true)
	if err != nil {
		return false, nil, err
	}

	latestTokenPrices, err := r.getLatestTokenPriceUpdates(ctx, now, true)
	if err != nil {
		return false, nil, err
	}

	priceUpdates := r.calculatePriceUpdates(parsableObservations, latestGasPrice, latestTokenPrices)
	// If there are no fee updates and the interval is zero there is no report to produce.
	if len(priceUpdates.TokenPriceUpdates) == 0 && priceUpdates.DestChainSelector == 0 && agreedInterval.Min == 0 {
		lggr.Infow("Empty report, skipping")
		return false, nil, nil
	}

	report, err := r.buildReport(ctx, lggr, agreedInterval, priceUpdates)
	if err != nil {
		return false, nil, err
	}
	encodedReport, err := abihelpers.EncodeCommitReport(report)
	if err != nil {
		return false, nil, err
	}
	lggr.Infow("Report",
		"merkleRoot", hex.EncodeToString(report.MerkleRoot[:]),
		"minSeqNr", report.Interval.Min,
		"maxSeqNr", report.Interval.Max,
		"tokenPriceUpdates", report.PriceUpdates.TokenPriceUpdates,
		"destChainSelector", report.PriceUpdates.DestChainSelector,
		"sourceUsdPerUnitGas", report.PriceUpdates.UsdPerUnitGas,
		"epochAndRound", epochAndRound,
	)
	return true, encodedReport, nil
}

// calculateIntervalConsensus compresses a set of intervals into one interval
// taking into account f which is the maximum number of faults across the whole DON.
// OCR itself won't call Report unless there are 2*f+1 observations
// https://github.com/smartcontractkit/libocr/blob/master/offchainreporting2/internal/protocol/report_generation_follower.go#L415
// and f of those observations may be either unparseable or adversarially set values. That means
// we'll either have f+1 parsed honest values here, 2f+1 parsed values with f adversarial values or somewhere
// in between.
// rangeLimit is the maximum range of the interval. If the interval is larger than this, it will be truncated. Zero means no limit.
func calculateIntervalConsensus(intervals []commit_store.CommitStoreInterval, f int, rangeLimit uint64) (commit_store.CommitStoreInterval, error) {
	// We require at least f+1 parsed values. This corresponds to the scenario where f of the 2f+1 are faulty
	// in the sense that they are unparseable.
	if len(intervals) <= f {
		return commit_store.CommitStoreInterval{}, errors.Errorf("Not enough intervals to form consensus: #obs=%d, f=%d", len(intervals), f)
	}

	// To understand min/max selection here, we need to consider an adversary that controls f values
	// and is intentionally trying to stall the protocol or influence the value returned. For simplicity
	// consider f=1 and n=4 nodes. In that case adversary may try to bias the min or max high/low.
	// We could end up (2f+1=3) with sorted_mins=[1,1,1e9] or [-1e9,1,1] as examples. Selecting
	// sorted_mins[f] ensures:
	// - At least one honest node has seen this value, so adversary cannot bias the value lower which
	// would cause reverts
	// - If an honest oracle reports sorted_min[f] which happens to be stale i.e. that oracle
	// has a delayed view of the chain, then the report will revert onchain but still succeed upon retry
	// - We minimize the risk of naturally hitting the error condition minSeqNum > maxSeqNum due to oracles
	// delayed views of the chain (would be an issue with taking sorted_mins[-f])
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Min < intervals[j].Min
	})
	minSeqNum := intervals[f].Min

	// The only way a report could have a minSeqNum of 0 is when there are no messages to report
	// and the report is potentially still valid for gas fee updates.
	if minSeqNum == 0 {
		return commit_store.CommitStoreInterval{Min: 0, Max: 0}, nil
	}
	// Consider a similar example to the sorted_mins one above except where they are maxes.
	// We choose the more "conservative" sorted_maxes[f] so:
	// - We are ensured that at least one honest oracle has seen the max, so adversary cannot set it lower and
	// cause the maxSeqNum < minSeqNum errors
	// - If an honest oracle reports sorted_max[f] which happens to be stale i.e. that oracle
	// has a delayed view of the source chain, then we simply lose a little bit of throughput.
	// - If we were to pick sorted_max[-f] i.e. the maximum honest node view (a more "aggressive" setting in terms of throughput),
	// then an adversary can continually send high values e.g. imagine we have observations from all 4 nodes
	// [honest 1, honest 1, honest 2, malicious 2], in this case we pick 2, but it's not enough to be able
	// to build a report since the first 2 honest nodes are unaware of message 2.
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Max < intervals[j].Max
	})
	maxSeqNum := intervals[f].Max
	if maxSeqNum < minSeqNum {
		// If the consensus report is invalid for onchain acceptance, we do not vote for it as
		// an early termination step.
		return commit_store.CommitStoreInterval{}, errors.New("max seq num smaller than min")
	}

	// If the range is too large, truncate it.
	if rangeLimit > 0 && maxSeqNum-minSeqNum+1 > rangeLimit {
		maxSeqNum = minSeqNum + rangeLimit - 1
	}

	return commit_store.CommitStoreInterval{
		Min: minSeqNum,
		Max: maxSeqNum,
	}, nil
}

// Note priceUpdates must be deterministic.
// The provided latestTokenPrices should not contain nil values.
func (r *CommitReportingPlugin) calculatePriceUpdates(observations []CommitObservation, latestGasPrice update, latestTokenPrices map[common.Address]update) commit_store.InternalPriceUpdates {
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

	var tokenPriceUpdates []commit_store.InternalTokenPriceUpdate
	for token, tokenPriceObservations := range priceObservations {
		// If majority report a token price, include it in the update
		if len(tokenPriceObservations) <= r.F {
			continue
		}
		medianPrice := ccipcalc.BigIntMedian(tokenPriceObservations)

		latestTokenPrice, exists := latestTokenPrices[token]
		if exists {
			tokenPriceUpdatedRecently := time.Since(latestTokenPrice.timestamp) < r.offchainConfig.FeeUpdateHeartBeat.Duration()
			tokenPriceNotChanged := !ccipcalc.Deviates(medianPrice, latestTokenPrice.value, int64(r.offchainConfig.FeeUpdateDeviationPPB))
			if tokenPriceUpdatedRecently && tokenPriceNotChanged {
				r.lggr.Debugw("price was updated recently, skipping the update",
					"token", token, "newPrice", medianPrice, "existingPrice", latestTokenPrice.value)
				continue // skip the update if we recently had a price update close to the new value
			}
		}

		tokenPriceUpdates = append(tokenPriceUpdates, commit_store.InternalTokenPriceUpdate{
			SourceToken: token,
			UsdPerToken: medianPrice,
		})
	}

	// Determinism required.
	sort.Slice(tokenPriceUpdates, func(i, j int) bool {
		return bytes.Compare(tokenPriceUpdates[i].SourceToken[:], tokenPriceUpdates[j].SourceToken[:]) == -1
	})

	usdPerUnitGas := big.NewInt(0)
	destChainSelector := uint64(0)

	if len(sourceGasObservations) > r.F {
		usdPerUnitGas = ccipcalc.BigIntMedian(sourceGasObservations) // Compute the median price
		destChainSelector = r.config.sourceChainSelector             // Assuming plugin lane is A->B, we write to B the gas price of A

		if latestGasPrice.value != nil {
			gasPriceUpdatedRecently := time.Since(latestGasPrice.timestamp) < r.offchainConfig.FeeUpdateHeartBeat.Duration()
			gasPriceNotChanged := !ccipcalc.Deviates(usdPerUnitGas, latestGasPrice.value, int64(r.offchainConfig.FeeUpdateDeviationPPB))
			if gasPriceUpdatedRecently && gasPriceNotChanged {
				usdPerUnitGas = big.NewInt(0)
				destChainSelector = uint64(0)
			}
		}
	}

	return commit_store.InternalPriceUpdates{
		TokenPriceUpdates: tokenPriceUpdates,
		DestChainSelector: destChainSelector,
		UsdPerUnitGas:     usdPerUnitGas, // we MUST pass zero to skip the update (never nil)
	}
}

// buildReport assumes there is at least one message in reqs.
func (r *CommitReportingPlugin) buildReport(ctx context.Context, lggr logger.Logger, interval commit_store.CommitStoreInterval, priceUpdates commit_store.InternalPriceUpdates) (commit_store.CommitStoreCommitReport, error) {
	// If no messages are needed only include fee updates
	if interval.Min == 0 {
		return commit_store.CommitStoreCommitReport{
			PriceUpdates: priceUpdates,
			MerkleRoot:   [32]byte{},
			Interval:     interval,
		}, nil
	}

	// Logs are guaranteed to be in order of seq num, since these are finalized logs only
	// and the contract's seq num is auto-incrementing.
	sendRequests, err := r.config.onRampReader.GetSendRequestsBetweenSeqNums(
		ctx,
		interval.Min,
		interval.Max,
		int(r.offchainConfig.SourceFinalityDepth),
	)
	if err != nil {
		return commit_store.CommitStoreCommitReport{}, err
	}
	if len(sendRequests) == 0 {
		lggr.Warn("No messages found in interval",
			"minSeqNr", interval.Min,
			"maxSeqNr", interval.Max)
		return commit_store.CommitStoreCommitReport{}, fmt.Errorf("tried building a tree without leaves")
	}

	leaves := make([][32]byte, 0, len(sendRequests))
	var seqNrs []uint64
	for _, req := range sendRequests {
		leaves = append(leaves, req.Data.Hash)
		seqNrs = append(seqNrs, req.Data.SequenceNumber)
	}
	if !ccipcalc.ContiguousReqs(lggr, interval.Min, interval.Max, seqNrs) {
		return commit_store.CommitStoreCommitReport{}, errors.Errorf("do not have full range [%v, %v] have %v", interval.Min, interval.Max, seqNrs)
	}
	tree, err := merklemulti.NewTree(hashlib.NewKeccakCtx(), leaves)
	if err != nil {
		return commit_store.CommitStoreCommitReport{}, err
	}

	return commit_store.CommitStoreCommitReport{
		PriceUpdates: priceUpdates,
		MerkleRoot:   tree.Root(),
		Interval:     interval,
	}, nil
}

func (r *CommitReportingPlugin) ShouldAcceptFinalizedReport(ctx context.Context, reportTimestamp types.ReportTimestamp, report types.Report) (bool, error) {
	parsedReport, err := abihelpers.DecodeCommitReport(report)
	if err != nil {
		return false, err
	}
	lggr := r.lggr.Named("CommitShouldAcceptFinalizedReport").With(
		"merkleRoot", parsedReport.MerkleRoot,
		"minSeqNum", parsedReport.Interval.Min,
		"maxSeqNum", parsedReport.Interval.Max,
		"destChainSelector", parsedReport.PriceUpdates.DestChainSelector,
		"usdPerUnitGas", parsedReport.PriceUpdates.UsdPerUnitGas,
		"tokenPriceUpdates", parsedReport.PriceUpdates.TokenPriceUpdates,
		"reportTimestamp", reportTimestamp,
	)
	// Empty report, should not be put on chain
	if parsedReport.MerkleRoot == [32]byte{} && parsedReport.PriceUpdates.DestChainSelector == 0 && len(parsedReport.PriceUpdates.TokenPriceUpdates) == 0 {
		lggr.Warn("Empty report, should not be put on chain")
		return false, nil
	}

	if r.isStaleReport(ctx, lggr, parsedReport, true, reportTimestamp) {
		lggr.Infow("Rejecting stale report")
		return false, nil
	}

	epochAndRound := ccipcalc.MergeEpochAndRound(reportTimestamp.Epoch, reportTimestamp.Round)
	if err := r.inflightReports.add(lggr, parsedReport, epochAndRound); err != nil {
		return false, err
	}
	lggr.Infow("Accepting finalized report", "merkleRoot", hexutil.Encode(parsedReport.MerkleRoot[:]))
	return true, nil
}

// ShouldTransmitAcceptedReport checks if the report is stale, if it is it should not be transmitted.
func (r *CommitReportingPlugin) ShouldTransmitAcceptedReport(ctx context.Context, reportTimestamp types.ReportTimestamp, report types.Report) (bool, error) {
	lggr := r.lggr.Named("CommitShouldTransmitAcceptedReport")
	parsedReport, err := abihelpers.DecodeCommitReport(report)
	if err != nil {
		return false, err
	}
	// If report is not stale we transmit.
	// When the commitTransmitter enqueues the tx for tx manager,
	// we mark it as fulfilled, effectively removing it from the set of inflight messages.
	shouldTransmit := !r.isStaleReport(ctx, lggr, parsedReport, false, reportTimestamp)

	lggr.Infow("ShouldTransmitAcceptedReport",
		"shouldTransmit", shouldTransmit,
		"reportTimestamp", reportTimestamp)
	return shouldTransmit, nil
}

// isStaleReport checks a report to see if the contents have become stale.
// It does so in four ways:
//  1. if there is a merkle root, check if the sequence numbers match up with onchain data
//  2. if there is no merkle root, check if current price's epoch and round is after onchain epoch and round
//  3. if there is a gas price update check to see if the value is different from the last
//     reported value
//  4. if there are token prices check to see if the values are different from the last
//     reported values.
//
// If there is a merkle root present, staleness is only measured based on the merkle root
// If there is no merkle root but there is a gas update, only this gas update is used for staleness checks.
// If only price updates are included, the price updates are used to check for staleness
// If nothing is included the report is always considered stale.
func (r *CommitReportingPlugin) isStaleReport(ctx context.Context, lggr logger.Logger, report commit_store.CommitStoreCommitReport, checkInflight bool, reportTimestamp types.ReportTimestamp) bool {
	// If there is a merkle root, ignore all other staleness checks and only check for sequence number staleness
	if report.MerkleRoot != [32]byte{} {
		return r.isStaleMerkleRoot(ctx, lggr, report.Interval, checkInflight)
	}

	hasGasPriceUpdate := report.PriceUpdates.DestChainSelector != 0
	hasTokenPriceUpdates := len(report.PriceUpdates.TokenPriceUpdates) > 0

	// If there is no merkle root, no gas price update and no token price update
	// we don't want to write anything on-chain, so we consider this report stale.
	if !hasGasPriceUpdate && !hasTokenPriceUpdates {
		return true
	}

	// We consider a price update as stale when, there isn't an update or there is an update that is stale.
	gasPriceStale := !hasGasPriceUpdate || r.isStaleGasPrice(ctx, lggr, report.PriceUpdates, checkInflight)
	tokenPricesStale := !hasTokenPriceUpdates || r.isStaleTokenPrices(ctx, lggr, report.PriceUpdates.TokenPriceUpdates, checkInflight)

	if gasPriceStale && tokenPricesStale {
		return true
	}

	// If report only has price update, check if its epoch and round lags behind the latest onchain
	lastPriceEpochAndRound, err := r.config.commitStore.GetLatestPriceEpochAndRound(&bind.CallOpts{Context: ctx})
	if err != nil {
		// Assume it's a transient issue getting the last report and try again on the next round
		return true
	}

	thisEpochAndRound := ccipcalc.MergeEpochAndRound(reportTimestamp.Epoch, reportTimestamp.Round)
	return lastPriceEpochAndRound >= thisEpochAndRound
}

func (r *CommitReportingPlugin) isStaleMerkleRoot(ctx context.Context, lggr logger.Logger, reportInterval commit_store.CommitStoreInterval, checkInflight bool) bool {
	nextInflightMin, nextOnChainMin, err := r.nextMinSeqNum(ctx, lggr)
	if err != nil {
		// Assume it's a transient issue getting the last report and try again on the next round
		return true
	}

	if checkInflight && nextInflightMin != reportInterval.Min {
		// There are sequence numbers missing between the commitStore/inflight txs and the proposed report.
		// The report will fail onchain unless the inflight cache is in an incorrect state. A state like this
		// could happen for various reasons, e.g. a reboot of the node emptying the caches, and should be self-healing.
		// We do not submit a tx and wait for the protocol to self-heal by updating the caches or invalidating
		// inflight caches over time.
		lggr.Errorw("Next inflight min is not equal to the proposed min of the report", "nextInflightMin", nextInflightMin)
		return true
	}

	if !checkInflight && nextOnChainMin > reportInterval.Min {
		// If the next min is already greater than this reports min, this report is stale.
		lggr.Infow("Report is stale because of root", "onchain min", nextOnChainMin, "report min", reportInterval.Min)
		return true
	}

	// If a report has root and valid sequence number, the report should be submitted, regardless of price staleness
	return false
}

func (r *CommitReportingPlugin) isStaleGasPrice(ctx context.Context, lggr logger.Logger, priceUpdates commit_store.InternalPriceUpdates, checkInflight bool) bool {
	gasPriceUpdate, err := r.getLatestGasPriceUpdate(ctx, time.Now(), checkInflight)
	if err != nil {
		return true
	}

	if gasPriceUpdate.value != nil && !ccipcalc.Deviates(priceUpdates.UsdPerUnitGas, gasPriceUpdate.value, int64(r.offchainConfig.FeeUpdateDeviationPPB)) {
		lggr.Infow("Report is stale because of gas price",
			"latestGasPriceUpdate", gasPriceUpdate.value,
			"usdPerUnitGas", priceUpdates.UsdPerUnitGas,
			"destChainSelector", priceUpdates.DestChainSelector)
		return true
	}

	return false
}

func (r *CommitReportingPlugin) isStaleTokenPrices(ctx context.Context, lggr logger.Logger, priceUpdates []commit_store.InternalTokenPriceUpdate, checkInflight bool) bool {
	// getting the last price updates without including inflight is like querying
	// current prices onchain, but uses logpoller's data to save on the RPC requests
	latestTokenPriceUpdates, err := r.getLatestTokenPriceUpdates(ctx, time.Now(), checkInflight)
	if err != nil {
		return true
	}

	for _, tokenUpdate := range priceUpdates {
		latestUpdate, ok := latestTokenPriceUpdates[tokenUpdate.SourceToken]
		priceEqual := ok && !ccipcalc.Deviates(tokenUpdate.UsdPerToken, latestUpdate.value, int64(r.offchainConfig.FeeUpdateDeviationPPB))

		if !priceEqual {
			lggr.Infow("Found non-stale token price", "token", tokenUpdate.SourceToken, "usdPerToken", tokenUpdate.UsdPerToken, "latestUpdate", latestUpdate.value)
			return false
		}
		lggr.Infow("Token price is stale", "latestTokenPrice", latestUpdate.value, "usdPerToken", tokenUpdate.UsdPerToken, "token", tokenUpdate.SourceToken)
	}

	lggr.Infow("All token prices are stale")
	return true
}

func (r *CommitReportingPlugin) Close() error {
	return nil
}
