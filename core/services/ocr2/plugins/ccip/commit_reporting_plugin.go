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
	evmclient "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/merklemulti"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

const MaxCommitReportLength = 1000

var (
	_ types.ReportingPluginFactory = &CommitReportingPluginFactory{}
	_ types.ReportingPlugin        = &CommitReportingPlugin{}
)

type InflightReport struct {
	report    *commit_store.CommitStoreCommitReport
	createdAt time.Time
}

type InflightPriceUpdate struct {
	priceUpdates commit_store.InternalPriceUpdates
	createdAt    time.Time
}

type update struct {
	timestamp time.Time
	value     *big.Int
}

type CommitPluginConfig struct {
	lggr                logger.Logger
	sourceLP, destLP    logpoller.LogPoller
	offRamp             evm_2_evm_offramp.EVM2EVMOffRampInterface
	onRampAddress       common.Address
	commitStore         commit_store.CommitStoreInterface
	priceGetter         PriceGetter
	sourceChainSelector uint64
	sourceNative        common.Address
	sourceFeeEstimator  txmgrtypes.FeeEstimator[*evmtypes.Head, gas.EvmFee, *assets.Wei, common.Hash]
	destClient          evmclient.Client
	leafHasher          hasher.LeafHasherInterface[[32]byte]
	getSeqNumFromLog    func(log logpoller.Log) (uint64, error)
}

type CommitReportingPlugin struct {
	config CommitPluginConfig
	F      int
	lggr   logger.Logger
	// We need to synchronize access to the inflight structure
	// as reporting plugin methods may be called from separate goroutines,
	// e.g. reporting vs transmission protocol.
	inFlightMu              sync.RWMutex
	inFlight                map[[32]byte]InflightReport
	inFlightPriceUpdates    []InflightPriceUpdate
	priceRegistry           price_registry.PriceRegistryInterface
	offchainConfig          ccipconfig.CommitOffchainConfig
	onchainConfig           ccipconfig.CommitOnchainConfig
	tokenToDecimalMappingMu sync.RWMutex
	tokenToDecimalMapping   map[common.Address]uint8
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
	onchainConfig, err := abihelpers.DecodeAbiStruct[ccipconfig.CommitOnchainConfig](config.OnchainConfig)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	offchainConfig, err := ccipconfig.DecodeOffchainConfig[ccipconfig.CommitOffchainConfig](config.OffchainConfig)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	priceRegistry, err := price_registry.NewPriceRegistry(onchainConfig.PriceRegistry, rf.config.destClient)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}

	err = rf.config.destLP.RegisterFilter(logpoller.Filter{
		Name:      logpoller.FilterName(COMMIT_PRICE_UPDATES, onchainConfig.PriceRegistry.String()),
		EventSigs: []common.Hash{abihelpers.EventSignatures.UsdPerUnitGasUpdated, abihelpers.EventSignatures.UsdPerTokenUpdated},
		Addresses: []common.Address{onchainConfig.PriceRegistry}})
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	err = rf.config.sourceLP.RegisterFilter(logpoller.Filter{
		Name:      logpoller.FilterName(COMMIT_CCIP_SENDS, rf.config.onRampAddress.String()),
		EventSigs: []common.Hash{abihelpers.EventSignatures.SendRequested},
		Addresses: []common.Address{rf.config.onRampAddress}})
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	tokenToDecimalMapping, err := generateTokenToDecimalMapping(context.Background(), rf.config, map[common.Address]uint8{}, priceRegistry)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}

	rf.config.lggr.Infow("Starting commit plugin", "offchainConfig", offchainConfig, "onchainConfig", onchainConfig)

	return &CommitReportingPlugin{
			config:                rf.config,
			F:                     config.F,
			lggr:                  rf.config.lggr.Named("CommitReportingPlugin"),
			inFlight:              make(map[[32]byte]InflightReport),
			priceRegistry:         priceRegistry,
			onchainConfig:         onchainConfig,
			offchainConfig:        offchainConfig,
			tokenToDecimalMapping: tokenToDecimalMapping,
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

// Query is not used by the CCIP Commit plugin.
func (r *CommitReportingPlugin) Query(context.Context, types.ReportTimestamp) (types.Query, error) {
	return types.Query{}, nil
}

// Observation calculates the sequence number interval ready to be committed and
// the token and gas price updates required. A valid report could contain a merkle
// root and/or price updates.
func (r *CommitReportingPlugin) Observation(ctx context.Context, _ types.ReportTimestamp, _ types.Query) (types.Observation, error) {
	lggr := r.lggr.Named("CommitObservation")
	// If the commit store is down the protocol should halt.
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

// expireInflight removed any expired entries from the inflight cache.
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
			// Sad path: inflight report reverts onchain, we remove it from inflight, onchain state does not reflect the chains, so we retry.
			lggr.Infow("Inflight price update expired", "updates", inFlightFeeUpdate.priceUpdates)
			stillInflight = append(stillInflight, inFlightFeeUpdate)
		}
	}
	r.inFlightPriceUpdates = stillInflight
}

func (r *CommitReportingPlugin) calculateMinMaxSequenceNumbers(ctx context.Context, lggr logger.Logger) (uint64, uint64, error) {
	nextMin, err := r.nextMinSeqNum(ctx)
	if err != nil {
		return 0, 0, err
	}
	// All available messages that have not been committed yet and have sufficient confirmations.
	lggr.Infof("Looking for requests with sig %v and nextMin %d on onRampAddr %v", abihelpers.EventSignatures.SendRequested, nextMin, r.config.onRampAddress)
	reqs, err := r.config.sourceLP.LogsDataWordGreaterThan(
		abihelpers.EventSignatures.SendRequested,
		r.config.onRampAddress,
		abihelpers.EventSignatures.SendRequestedSequenceNumberWord,
		abihelpers.EvmWord(nextMin),
		int(r.offchainConfig.SourceIncomingConfirmations),
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return 0, 0, err
	}
	lggr.Infof("%d requests found for onRampAddr %v", len(reqs), r.config.onRampAddress)
	if len(reqs) == 0 {
		return 0, 0, nil
	}
	var seqNrs []uint64
	for _, req := range reqs {
		seqNr, err2 := r.config.getSeqNumFromLog(req)
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
		lggr.Warnf("Missing sequence number range [%d-%d] for onRamp %v", nextMin, min, r.config.onRampAddress)
	}
	if !contiguousReqs(lggr, min, max, seqNrs) {
		return 0, 0, errors.New("unexpected gap in seq nums")
	}
	lggr.Infof("OnRamp %v: min %v max %v", r.config.onRampAddress, min, max)
	return min, max, nil
}

func (r *CommitReportingPlugin) nextMinSeqNum(ctx context.Context) (uint64, error) {
	nextMin, err := r.config.commitStore.GetExpectedNextSequenceNumber(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, err
	}
	// loop through sorted inFlightReports
	// only increment nextMin if the report build ontop the running nextMin
	r.inFlightMu.RLock()
	defer r.inFlightMu.RUnlock()
	for _, report := range r.inFlight {
		if report.report.Interval.Max >= nextMin {
			nextMin = report.report.Interval.Max + 1
		}
	}
	return nextMin, nil
}

// All prices are USD ($1=1e18) denominated. We only generate prices we think should be updated;
// otherwise, omitting values means voting to skip updating them
func (r *CommitReportingPlugin) generatePriceUpdates(
	ctx context.Context,
	now time.Time,
) (sourceGasPriceUSD *big.Int, tokenPricesUSD map[common.Address]*big.Int, err error) {
	// Detect token changes and update decimals mapping if needed, so we are up-to-date with supported tokens
	if err = r.updateTokenToDecimalMapping(ctx); err != nil {
		return nil, nil, err
	}

	r.tokenToDecimalMappingMu.RLock()
	defer r.tokenToDecimalMappingMu.RUnlock()

	tokensWithDecimal := make([]common.Address, 0, len(r.tokenToDecimalMapping))
	for token := range r.tokenToDecimalMapping {
		tokensWithDecimal = append(tokensWithDecimal, token)
	}

	queryTokens := append([]common.Address{r.config.sourceNative}, tokensWithDecimal...)
	// Include wrapped native in our token query as way to identify the source native USD price.
	// notice USD is in 1e18 scale, i.e. $1 = 1e18
	rawTokenPricesUSD, err := r.config.priceGetter.TokenPricesUSD(ctx, queryTokens)
	if err != nil {
		return nil, nil, err
	}
	for _, token := range queryTokens {
		if rawTokenPricesUSD[token] == nil {
			return nil, nil, errors.Errorf("missing token price: %+v", token)
		}
	}

	sourceNativePriceUSD := rawTokenPricesUSD[r.config.sourceNative]
	tokenPricesUSD = make(map[common.Address]*big.Int, len(rawTokenPricesUSD))
	for token := range rawTokenPricesUSD {
		if !slices.Contains(tokensWithDecimal, token) {
			// do not include any address which isn't a supported token on dest chain, including sourceNative
			continue
		}

		decimals, ok := r.tokenToDecimalMapping[token]
		if !ok {
			return nil, nil, errors.Errorf("missing token decimals: %+v", token)
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

	sourceGasPriceUSD = calculateUsdPerUnitGas(gasPrice, sourceNativePriceUSD)

	gasPriceUpdate, err := r.getLatestGasPriceUpdate(ctx, now, false)
	if err != nil {
		return nil, nil, err
	}

	if gasPriceUpdate.value != nil && now.Sub(gasPriceUpdate.timestamp) < r.offchainConfig.FeeUpdateHeartBeat.Duration() && !deviates(sourceGasPriceUSD, gasPriceUpdate.value, int64(r.offchainConfig.FeeUpdateDeviationPPB)) {
		// vote skip gasPrice update by leaving it nil
		sourceGasPriceUSD = nil
	}

	tokenPriceUpdates, err := r.getLatestTokenPriceUpdates(ctx, now, false)
	if err != nil {
		return nil, nil, err
	}

	for token, price := range tokenPricesUSD {
		tokenUpdate := tokenPriceUpdates[token]
		if tokenUpdate.value != nil && now.Sub(tokenUpdate.timestamp) < r.offchainConfig.FeeUpdateHeartBeat.Duration() && !deviates(price, tokenUpdate.value, int64(r.offchainConfig.FeeUpdateDeviationPPB)) {
			// vote skip tokenPrice update by not including it in price map
			delete(tokenPricesUSD, token)
		}
	}

	// either may be empty
	return sourceGasPriceUSD, tokenPricesUSD, nil
}

// Input price is USD per full token, in base units 1e18
// Result price is USD per 1e18 of smallest token denomination, in base units 1e18
// Example: 1 USDC = 1.00 USD per full token, each full token is 6 decimals -> 1 * 1e18 * 1e18 / 1e6 = 1e30
func calculateUsdPer1e18TokenAmount(price *big.Int, decimals uint8) *big.Int {
	tmp := big.NewInt(0).Mul(price, big.NewInt(1e18))
	return tmp.Div(tmp, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
}

// Checks for updates, updates the token to decimal mapping.
func (r *CommitReportingPlugin) updateTokenToDecimalMapping(ctx context.Context) error {
	newTokenToDecimalMapping, err := generateTokenToDecimalMapping(ctx, r.config, r.tokenToDecimalMapping, r.priceRegistry)
	if err != nil {
		return err
	}

	// Pre-emptively guarding plugin state changes
	// Going forward, it may be possible for Should* OCR2 functions for call this from another thread
	r.tokenToDecimalMappingMu.Lock()
	r.tokenToDecimalMapping = newTokenToDecimalMapping
	r.tokenToDecimalMappingMu.Unlock()

	return nil
}

// Generates the token to decimal mapping for dest tokens and fee tokens.
// NOTE: this queries token decimals n times, where n is the number of tokens whose decimals are not already cached.
func generateTokenToDecimalMapping(ctx context.Context, config CommitPluginConfig, curMapping map[common.Address]uint8, priceRegistry price_registry.PriceRegistryInterface) (map[common.Address]uint8, error) {
	newMapping := make(map[common.Address]uint8)

	destTokens, err := config.offRamp.GetDestinationTokens(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	feeTokens, err := priceRegistry.GetFeeTokens(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	// In case if a fee token is not an offramp dest token, we still want to update its decimals and price
	for _, feeToken := range feeTokens {
		if !slices.Contains(destTokens, feeToken) {
			destTokens = append(destTokens, feeToken)
		}
	}

	for _, token := range destTokens {
		if curDecimal, ok := curMapping[token]; ok {
			// If token already in mapping, no need to call decimals again, decimals should be immutable
			newMapping[token] = curDecimal
			continue
		}
		tokenContract, err := link_token_interface.NewLinkToken(token, config.destClient)
		if err != nil {
			return nil, err
		}

		decimal, err := tokenContract.Decimals(&bind.CallOpts{Context: ctx})
		if err != nil {
			return nil, err
		}
		newMapping[token] = decimal
	}
	return newMapping, nil
}

// Gets the latest token price updates based on logs within the heartbeat
func (r *CommitReportingPlugin) getLatestTokenPriceUpdates(ctx context.Context, now time.Time, skipInflight bool) (map[common.Address]update, error) {
	tokenUpdatesWithinHeartBeat, err := r.config.destLP.LogsCreatedAfter(abihelpers.EventSignatures.UsdPerTokenUpdated, r.priceRegistry.Address(), now.Add(-r.offchainConfig.FeeUpdateHeartBeat.Duration()), pg.WithParentCtx(ctx))
	latestUpdates := make(map[common.Address]update)

	if err != nil {
		return nil, err
	}
	for _, log := range tokenUpdatesWithinHeartBeat {
		// Ordered by ascending timestamps
		tokenUpdate, err := r.priceRegistry.ParseUsdPerTokenUpdated(log.GetGethLog())
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

// Gets the latest gas price updates based on logs within the heartbeat
func (r *CommitReportingPlugin) getLatestGasPriceUpdate(ctx context.Context, now time.Time, skipInflight bool) (gasPriceUpdate update, error error) {
	gasUpdatesWithinHeartBeat, err := r.config.destLP.IndexedLogsCreatedAfter(
		abihelpers.EventSignatures.UsdPerUnitGasUpdated,
		r.priceRegistry.Address(),
		1,
		[]common.Hash{abihelpers.EvmWord(r.config.sourceChainSelector)},
		now.Add(-r.offchainConfig.FeeUpdateHeartBeat.Duration()),
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return update{}, err
	}

	for _, log := range gasUpdatesWithinHeartBeat {
		// Ordered by ascending timestamps
		priceUpdate, err2 := r.priceRegistry.ParseUsdPerUnitGasUpdated(log.GetGethLog())
		if err2 != nil {
			return update{}, err2
		}
		timestamp := time.Unix(priceUpdate.Timestamp.Int64(), 0)
		if !timestamp.Before(gasPriceUpdate.timestamp) {
			gasPriceUpdate = update{
				timestamp: timestamp,
				value:     priceUpdate.Value,
			}
		}
	}

	if skipInflight {
		return gasPriceUpdate, nil
	}

	r.inFlightMu.RLock()
	defer r.inFlightMu.RUnlock()
	for _, inflight := range r.inFlightPriceUpdates {
		if inflight.priceUpdates.DestChainSelector != 0 && !inflight.createdAt.Before(gasPriceUpdate.timestamp) {
			gasPriceUpdate = update{
				timestamp: inflight.createdAt,
				value:     inflight.priceUpdates.UsdPerUnitGas,
			}
		}
	}

	return gasPriceUpdate, nil
}

// deviation_parts_per_billion = ((x2 - x1) / x1) * 1e9
func deviates(x1, x2 *big.Int, feeUpdateDeviationPPB int64) bool {
	// if x1 == 0, deviates if x2 != x1, to avoid the relative division by 0 error
	if x1.BitLen() == 0 {
		return x1.Cmp(x2) != 0
	}
	diff := big.NewInt(0).Sub(x1, x2)
	diff.Mul(diff, big.NewInt(1e9))
	diff.Div(diff, x1)
	return diff.CmpAbs(big.NewInt(feeUpdateDeviationPPB)) > 0
}

func (r *CommitReportingPlugin) Report(ctx context.Context, _ types.ReportTimestamp, _ types.Query, observations []types.AttributedObservation) (bool, types.Report, error) {
	lggr := r.lggr.Named("Report")
	parsableObservations := getParsableObservations[CommitObservation](lggr, observations)
	var intervals []commit_store.CommitStoreInterval
	for _, obs := range parsableObservations {
		intervals = append(intervals, obs.Interval)
	}

	agreedInterval, err := calculateIntervalConsensus(intervals, r.F)
	if err != nil {
		return false, nil, err
	}

	priceUpdates := calculatePriceUpdates(r.config.sourceChainSelector, parsableObservations, r.F)
	// If there are no fee updates and the interval is zero there is no report to produce.
	if len(priceUpdates.TokenPriceUpdates) == 0 && priceUpdates.DestChainSelector == 0 && agreedInterval.Min == 0 {
		return false, nil, nil
	}

	report, err := r.buildReport(ctx, agreedInterval, priceUpdates)
	if err != nil {
		return false, nil, err
	}
	encodedReport, err := abihelpers.EncodeCommitReport(report)
	if err != nil {
		return false, nil, err
	}
	lggr.Infow("Built report", "interval", agreedInterval)
	return true, encodedReport, nil
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

// Note priceUpdates must be deterministic.
func calculatePriceUpdates(destChainSelector uint64, observations []CommitObservation, f int) commit_store.InternalPriceUpdates {
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
		destChainSelector = 0
	} else {
		usdPerUnitGas = median(sourceGasObservations)
	}

	return commit_store.InternalPriceUpdates{
		TokenPriceUpdates: priceUpdates,
		// Sending zero is ok, UsdPerUnitGas update is skipped
		DestChainSelector: destChainSelector,
		UsdPerUnitGas:     usdPerUnitGas,
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

// buildReport assumes there is at least one message in reqs.
func (r *CommitReportingPlugin) buildReport(ctx context.Context, interval commit_store.CommitStoreInterval, priceUpdates commit_store.InternalPriceUpdates) (commit_store.CommitStoreCommitReport, error) {
	lggr := r.lggr.Named("BuildReport")

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
	logs, err := r.config.sourceLP.LogsDataWordRange(
		abihelpers.EventSignatures.SendRequested,
		r.config.onRampAddress,
		abihelpers.EventSignatures.SendRequestedSequenceNumberWord,
		logpoller.EvmWord(interval.Min),
		logpoller.EvmWord(interval.Max),
		int(r.offchainConfig.SourceIncomingConfirmations),
		pg.WithParentCtx(ctx))
	if err != nil {
		return commit_store.CommitStoreCommitReport{}, err
	}
	leaves, err := leavesFromIntervals(lggr, r.config.getSeqNumFromLog, interval, r.config.leafHasher, logs)
	if err != nil {
		return commit_store.CommitStoreCommitReport{}, err
	}

	if len(leaves) == 0 {
		return commit_store.CommitStoreCommitReport{}, fmt.Errorf("tried building a tree without leaves for onRampAddr %v. %+v", r.config.onRampAddress, leaves)
	}
	tree, err := merklemulti.NewTree(hasher.NewKeccakCtx(), leaves)
	if err != nil {
		return commit_store.CommitStoreCommitReport{}, err
	}

	return commit_store.CommitStoreCommitReport{
		PriceUpdates: priceUpdates,
		MerkleRoot:   tree.Root(),
		Interval:     interval,
	}, nil
}

func (r *CommitReportingPlugin) ShouldAcceptFinalizedReport(ctx context.Context, _ types.ReportTimestamp, report types.Report) (bool, error) {
	lggr := r.lggr.Named("ShouldAcceptFinalizedReport")
	parsedReport, err := abihelpers.DecodeCommitReport(report)
	if err != nil {
		return false, err
	}
	// Empty report, should not be put on chain
	if parsedReport.MerkleRoot == [32]byte{} && parsedReport.PriceUpdates.DestChainSelector == 0 && len(parsedReport.PriceUpdates.TokenPriceUpdates) == 0 {
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

	r.addToInflight(lggr, &parsedReport)
	lggr.Infow("Accepting finalized report", "merkleRoot", hexutil.Encode(parsedReport.MerkleRoot[:]))
	return true, nil
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

	if report.PriceUpdates.DestChainSelector != 0 || len(report.PriceUpdates.TokenPriceUpdates) != 0 {
		lggr.Infow("Adding to inflight fee updates", "priceUpdates", report.PriceUpdates)
		r.inFlightPriceUpdates = append(r.inFlightPriceUpdates, InflightPriceUpdate{
			priceUpdates: report.PriceUpdates,
			createdAt:    time.Now(),
		})
	}
}

// ShouldTransmitAcceptedReport checks if the report is stale, if it is it should not be
// transmitted.
func (r *CommitReportingPlugin) ShouldTransmitAcceptedReport(ctx context.Context, _ types.ReportTimestamp, report types.Report) (bool, error) {
	parsedReport, err := abihelpers.DecodeCommitReport(report)
	if err != nil {
		return false, err
	}
	// If report is not stale we transmit.
	// When the commitTransmitter enqueues the tx for tx manager,
	// we mark it as fulfilled, effectively removing it from the set of inflight messages.
	return !r.isStaleReport(ctx, parsedReport), nil
}

// isStaleReport checks a report to see if the contents have become stale.
// It does so in three ways:
//  1. if there is a merkle root, check if the sequence numbers match up with onchain data
//  2. if there is a gas price update check to see if the value is different from the last
//     reported value
//  3. if there are token prices check to see if the values are different from the last
//     reported values.
//
// If there is a merkle root present, staleness is only measured based on the merkle root
// If there is no merkle root but there is a gas update, only this gas update is used
// for staleness checks.
// If only price updates are included, only price updates are used for staleness
// If nothing is included the report is always considered stale.
func (r *CommitReportingPlugin) isStaleReport(ctx context.Context, report commit_store.CommitStoreCommitReport) bool {
	// There could be a report with only price updates, in that case ignore sequence number staleness
	if report.MerkleRoot != [32]byte{} {
		nextMin, err := r.config.commitStore.GetExpectedNextSequenceNumber(&bind.CallOpts{Context: ctx})
		if err != nil {
			// Assume it's a transient issue getting the last report
			// Will try again on the next round
			return true
		}
		// If the next min is already greater than this reports min, this report is stale.
		if nextMin > report.Interval.Min {
			r.lggr.Infow("report is stale", "onchain min", nextMin, "report min", report.Interval.Min)
			return true
		}
		return false
	}

	if report.PriceUpdates.DestChainSelector != 0 {
		gasPriceUpdate, err := r.getLatestGasPriceUpdate(ctx, time.Now(), true)
		if err != nil {
			return true
		}

		if gasPriceUpdate.value != nil && gasPriceUpdate.value.Cmp(report.PriceUpdates.UsdPerUnitGas) == 0 {
			r.lggr.Infow("gasPriceUpdate-only report is stale", "latest gasPrice", gasPriceUpdate.value, "destChainSelector", report.PriceUpdates.DestChainSelector)
			return true
		}
		return false
	}

	if len(report.PriceUpdates.TokenPriceUpdates) > 0 {
		// getting the last price updates without including inflight is like querying
		// current prices onchain, but uses logpoller's data to save on the RPC requests
		tokenPriceUpdates, err := r.getLatestTokenPriceUpdates(ctx, time.Now(), true)
		if err != nil {
			return true
		}
		// check first token
		tokenUpdate := report.PriceUpdates.TokenPriceUpdates[0]
		if latestUpdate, ok := tokenPriceUpdates[tokenUpdate.SourceToken]; ok && latestUpdate.value.Cmp(tokenUpdate.UsdPerToken) == 0 {
			r.lggr.Infow("tokenPriceUpdate-only report is stale", "latest tokenPrice", latestUpdate.value, "token", tokenUpdate.SourceToken)
			return true
		}
		return false
	}
	// Can only get here if there is no merkle root, no gas price update and no token price update
	// If so, we don't want to write anything onchain, so we consider this report stale.
	return true
}

func (r *CommitReportingPlugin) Close() error {
	return nil
}

// CommitReportToEthTxMeta generates a txmgr.EthTxMeta from the given commit report.
// sequence numbers of the committed messages will be added to tx metadata
func CommitReportToEthTxMeta(report []byte) (*txmgr.EthTxMeta, error) {
	commitReport, err := abihelpers.DecodeCommitReport(report)
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
