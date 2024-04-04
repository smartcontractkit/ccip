package ccipcommit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"math/big"
	"strings"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"

	"github.com/smartcontractkit/chainlink/v2/core/chains/legacyevm"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipcalc"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/pricegetter"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/rpclib"
	"github.com/smartcontractkit/chainlink/v2/core/services/pipeline"
)

type MultiLanePriceGetter interface {
	GetTokenPrices(ctx context.Context, tokens map[cciptypes.Address][]cciptypes.Address) (map[cciptypes.Address]*big.Int, error)
}

// multiLanePriceGetter is a collection of price getters, one for each lane.
type multiLanePriceGetter struct {
	priceGetters map[cciptypes.Address]pricegetter.PriceGetter

	jobCache *ocr2.SharedJobCache
	pr       pipeline.Runner
	chainSet legacyevm.LegacyChainContainer

	lggr logger.Logger
}

// NewMultiLanePriceGetter creates a new instance of NewMultiLanePriceGetter.
func NewMultiLanePriceGetter(lggr logger.Logger, jobCache *ocr2.SharedJobCache, pr pipeline.Runner, chainSet legacyevm.LegacyChainContainer) MultiLanePriceGetter {
	return &multiLanePriceGetter{
		priceGetters: make(map[cciptypes.Address]pricegetter.PriceGetter),

		jobCache: jobCache,
		pr:       pr,
		chainSet: chainSet,

		lggr: lggr,
	}
}

func (s *multiLanePriceGetter) syncPriceGetters() error {
	jobs := s.jobCache.Get()
	priceGetters := make(map[cciptypes.Address]pricegetter.PriceGetter)
	for _, jb := range jobs {
		pluginConfig, err := extractPluginConfig(jb)
		if err != nil {
			return err
		}

		offRamp := pluginConfig.OffRamp
		// if PriceGetting matching the job's OffRamp already exists, reuse the PriceGetter.
		if _, exists := s.priceGetters[offRamp]; exists {
			priceGetters[offRamp] = s.priceGetters[offRamp]
			continue
		}

		var priceGetter pricegetter.PriceGetter
		withPipeline := strings.Trim(pluginConfig.TokenPricesUSDPipeline, "\n\t ") != ""
		if withPipeline {
			priceGetter, err = pricegetter.NewPipelineGetter(pluginConfig.TokenPricesUSDPipeline, s.pr, jb.ID, jb.ExternalJobID, jb.Name.ValueOrZero(), s.lggr)
			if err != nil {
				return fmt.Errorf("creating pipeline price getter: %w", err)
			}
		} else {
			// Use dynamic price getter.
			if pluginConfig.PriceGetterConfig == nil {
				return fmt.Errorf("priceGetterConfig is nil")
			}

			// Build price getter clients for all chains specified in the aggregator configurations.
			// Some lanes (e.g. Wemix/Kroma) requires other clients than source and destination, since they use feeds from other chains.
			priceGetterClients := map[uint64]pricegetter.DynamicPriceGetterClient{}
			for _, aggCfg := range pluginConfig.PriceGetterConfig.AggregatorPrices {
				chainID := aggCfg.ChainID
				// Retrieve the chain.
				chain, _, err2 := ccipconfig.GetChainByChainID(s.chainSet, chainID)
				if err2 != nil {
					return fmt.Errorf("retrieving chain for chainID %d: %w", chainID, err2)
				}
				caller := rpclib.NewDynamicLimitedBatchCaller(
					s.lggr,
					chain.Client(),
					rpclib.DefaultRpcBatchSizeLimit,
					rpclib.DefaultRpcBatchBackOffMultiplier,
					rpclib.DefaultMaxParallelRpcCalls,
				)
				priceGetterClients[chainID] = pricegetter.NewDynamicPriceGetterClient(caller)
			}

			priceGetter, err = pricegetter.NewDynamicPriceGetter(*pluginConfig.PriceGetterConfig, priceGetterClients)
			if err != nil {
				return fmt.Errorf("creating dynamic price getter: %w", err)
			}
		}

		priceGetters[offRamp] = priceGetter
	}

	s.priceGetters = priceGetters

	return nil
}

func extractPluginConfig(jb job.Job) (*ccipconfig.CommitPluginJobSpecConfig, error) {
	if jb.OCR2OracleSpec == nil {
		return nil, errors.New("spec is nil")
	}
	spec := jb.OCR2OracleSpec

	var pluginConfig ccipconfig.CommitPluginJobSpecConfig
	err := json.Unmarshal(spec.PluginConfig.Bytes(), &pluginConfig)
	if err != nil {
		return nil, err
	}
	// ensure addresses are formatted properly - (lowercase to eip55 for evm)
	pluginConfig.OffRamp = ccipcalc.HexToAddress(string(pluginConfig.OffRamp))

	return &pluginConfig, nil
}

// GetTokenPrices looks up token prices from OffRamp's corresponding price getters.
// Matt TODO parallelize this
func (s *multiLanePriceGetter) GetTokenPrices(ctx context.Context, tokens map[cciptypes.Address][]cciptypes.Address) (map[cciptypes.Address]*big.Int, error) {
	if err := s.syncPriceGetters(); err != nil {
		return nil, err
	}

	combinedPrices := make(map[cciptypes.Address]*big.Int)
	for offRamp, tokens := range tokens {
		priceGetter, exists := s.priceGetters[offRamp]
		if !exists {
			return nil, fmt.Errorf("priceGetter for OffRamp %s does not exist", offRamp)
		}

		prices, err := priceGetter.TokenPricesUSD(ctx, tokens)
		if err != nil {
			return nil, err
		}

		maps.Copy(combinedPrices, prices)
	}

	return combinedPrices, nil
}
