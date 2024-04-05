package ccipcommit

import (
	"context"
	"maps"
	"math/big"
	"sort"
	"sync"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipcommon"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/pricegetter"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/prices"
)

const (
	priceGetterBatchSizeLimit = 15
)

type CommitCrossLaneManager interface {
	// GetTokenPrices(ctx context.Context, tokens map[cciptypes.Address][]cciptypes.Address) (map[cciptypes.Address]*big.Int, error)
}

type laneAsset struct {
	pricegetter       pricegetter.PriceGetter
	sourceNative      cciptypes.Address
	offRampReader     ccipdata.OffRampReader
	gasPriceEstimator prices.GasPriceEstimatorCommit
}

// commitCrossLaneManager manages assets such as tokens and price getters across lanes.
type commitCrossLaneManager struct {
	lanes map[int32]laneAsset // jobID -> lane assets
	lggr  logger.Logger

	priceRegistry cciptypes.PriceRegistryReader

	mu sync.RWMutex
}

// NewCommitCrossLaneManager creates a new instance of CrossLaneManager.
func NewCommitCrossLaneManager(lggr logger.Logger, jobs []job.Job, priceRegistry cciptypes.PriceRegistryReader) CommitCrossLaneManager {
	return &commitCrossLaneManager{
		lanes: make(map[int32]laneAsset),
		lggr:  lggr,
	}
}

// AddLane removes a lane from the manager.
// This should be called immediately after the assets are initialized at Commit service creation.
func (s *commitCrossLaneManager) AddLane(jb job.Job, pricegetter pricegetter.PriceGetter, sourceNative cciptypes.Address, priceRegistry cciptypes.PriceRegistryReader, offRampReader ccipdata.OffRampReader, gasPriceEstimator prices.GasPriceEstimatorCommit) {
	la := laneAsset{
		pricegetter:       pricegetter,
		sourceNative:      sourceNative,
		offRampReader:     offRampReader,
		gasPriceEstimator: gasPriceEstimator,
	}

	s.lanes[jb.ID] = la
	s.priceRegistry = priceRegistry
}

// DeleteLane removes a lane from the manager.
// This should be called immediately before the filters are unregistered at job deletion.
func (s *commitCrossLaneManager) DeleteLane(jb job.Job) {
	delete(s.lanes, jb.ID)
}

// GetAllTokens returns all tokens across all lanes, including all supported tokens, fee tokens, and all source natives
func (s *commitCrossLaneManager) GetAllTokens(ctx context.Context) (allTokens []cciptypes.Address, err error) {
	destTokens, err := s.GetDestTokens(ctx)
	if err != nil {
		return nil, err
	}

	var sourceNatives []cciptypes.Address
	for _, la := range s.lanes {
		sourceNatives = append(sourceNatives, la.sourceNative)
	}
	// sort source natives for deterministic result
	sort.Slice(sourceNatives, func(i, j int) bool {
		return sourceNatives[i] < sourceNatives[j]
	})

	return ccipcommon.FlattenUniqueSlice(sourceNatives, destTokens), nil
}

// GetDestTokens returns all dest tokens across all lanes, including all supported tokens and fee tokens
func (s *commitCrossLaneManager) GetDestTokens(ctx context.Context) (allTokens []cciptypes.Address, err error) {
	var offRampReaders []ccipdata.OffRampReader
	for _, la := range s.lanes {
		offRampReaders = append(offRampReaders, la.offRampReader)
	}

	return ccipcommon.GetSortedChainTokens(ctx, offRampReaders, s.priceRegistry)
}

// GetAllTokenPrices looks up token prices from OffRamp's corresponding price getters.
func (s *commitCrossLaneManager) GetAllTokenPrices(ctx context.Context) (map[cciptypes.Address]*big.Int, error) {
	return s.getAllTokenPrices(ctx, priceGetterBatchSizeLimit)
}

func (s *commitCrossLaneManager) getAllTokenPrices(ctx context.Context, batchSize int) (map[cciptypes.Address]*big.Int, error) {

	if err := s.syncPriceGetters(); err != nil {
		return nil, err
	}

	combinedPrices := make(map[cciptypes.Address]*big.Int)
	for offRamp, tokens := range tokensPerOffRamp {
		// short circuit if there are no tokens to get prices for on a lane
		if len(tokens) == 0 {
			continue
		}
		priceGetter, exists := s.priceGetters[offRamp]
		if !exists {
			// Matt TODO
			// missing offramp means the
			continue
		}

		prices, err := priceGetter.TokenPricesUSD(ctx, tokens)
		if err != nil {
			return nil, err
		}

		maps.Copy(combinedPrices, prices)
	}

	return combinedPrices, nil
}

func (s *commitCrossLaneManager) getAllTokenPrices(ctx context.Context, batchSize int) (map[cciptypes.Address]*big.Int, error) {

}
