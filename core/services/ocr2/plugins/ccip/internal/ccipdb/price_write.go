package db

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	cciporm "github.com/smartcontractkit/chainlink/v2/core/services/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipcommon"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/pricegetter"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/prices"
)

// PriceWrite writes latest gas and token prices into db. To make debugging and updating easier,
// new prices are always inserted to the DB, as opposed to overwriting the old ones.
type PriceWrite interface {
	job.ServiceCtx
	UpdateDynamicConfig(prices.GasPriceEstimatorCommit, ccipdata.PriceRegistryReader) error
}

var _ PriceWrite = (*priceWrite)(nil)

const (
	// Prices can be updated every 1 minute, to be consistent with existing CommitDON round time
	priceUpdateInterval = 60 * time.Second
)

type priceWrite struct {
	updateInterval time.Duration

	lggr              logger.Logger
	orm               cciporm.ORM
	jobId             int32
	destChainSelector uint64

	sourceChainSelector     uint64
	sourceNative            cciptypes.Address
	priceGetter             pricegetter.PriceGetter
	offRampReader           ccipdata.OffRampReader
	gasPriceEstimator       prices.GasPriceEstimatorCommit
	destPriceRegistryReader ccipdata.PriceRegistryReader

	services.StateMachine
	wg               *sync.WaitGroup
	backgroundCtx    context.Context
	backgroundCancel context.CancelFunc
	dynamicConfigMu  *sync.Mutex
}

func NewPriceWrite(
	lggr logger.Logger,
	orm cciporm.ORM,
	jobId int32,
	destChainSelector uint64,
	sourceChainSelector uint64,

	sourceNative cciptypes.Address,
	priceGetter pricegetter.PriceGetter,
	offRampReader ccipdata.OffRampReader,
	gasPriceEstimator prices.GasPriceEstimatorCommit,
	destPriceRegistryReader ccipdata.PriceRegistryReader,
) PriceWrite {
	ctx, cancel := context.WithCancel(context.Background())

	pw := &priceWrite{
		updateInterval: priceUpdateInterval,

		lggr:              lggr,
		orm:               orm,
		jobId:             jobId,
		destChainSelector: destChainSelector,

		sourceChainSelector:     sourceChainSelector,
		sourceNative:            sourceNative,
		priceGetter:             priceGetter,
		offRampReader:           offRampReader,
		gasPriceEstimator:       gasPriceEstimator,
		destPriceRegistryReader: destPriceRegistryReader,

		wg:               new(sync.WaitGroup),
		backgroundCtx:    ctx,
		backgroundCancel: cancel,
		dynamicConfigMu:  &sync.Mutex{},
	}
	return pw
}

func (c *priceWrite) Start(context.Context) error {
	return c.StateMachine.StartOnce("PriceWrite", func() error {
		c.lggr.Info("Starting PriceWrite")
		c.wg.Add(1)
		c.run()
		return nil
	})
}

func (c *priceWrite) Close() error {
	return c.StateMachine.StopOnce("PriceWrite", func() error {
		c.lggr.Info("Closing PriceWrite")
		c.backgroundCancel()
		c.wg.Wait()
		return nil
	})
}

func (c *priceWrite) run() {
	ticker := time.NewTicker(c.updateInterval)
	go func() {
		defer c.wg.Done()

		for {
			select {
			case <-c.backgroundCtx.Done():
				return
			case <-ticker.C:
				err := c.executeUpdate(c.backgroundCtx)
				if err != nil {
					c.lggr.Errorw("Failed to write in-db prices in the background", "err", err)
				}
			}
		}
	}()
}

func (c *priceWrite) UpdateDynamicConfig(gasPriceEstimator prices.GasPriceEstimatorCommit, destPriceRegistryReader ccipdata.PriceRegistryReader) error {
	c.dynamicConfigMu.Lock()
	defer c.dynamicConfigMu.Unlock()

	c.gasPriceEstimator = gasPriceEstimator
	c.destPriceRegistryReader = destPriceRegistryReader

	// Config update may substantially change the prices, refresh the prices immediately, this also makes testing easier
	// for not having to wait to the full update interval.
	err := c.updatePrices(c.backgroundCtx)
	if err != nil {
		c.lggr.Errorw("Failed to write in-db prices after config update", "err", err)
	}

	return nil
}

func (c *priceWrite) executeUpdate(ctx context.Context) error {
	// Price updates happen infrequently - once every `priceUpdateInterval` seconds or DynamicConfig update.
	// It does not happen on any code path that is performance sensitive.
	// We can afford to have non-performant concurrency protection here that is simple and safe.
	c.dynamicConfigMu.Lock()
	defer c.dynamicConfigMu.Unlock()

	return c.updatePrices(ctx)
}

func (c *priceWrite) updatePrices(ctx context.Context) error {
	if c.gasPriceEstimator == nil || c.destPriceRegistryReader == nil {
		return fmt.Errorf("PriceWrite ORM skipping price update, gasPriceEstimator and/or destPriceRegistry is not set yet")
	}

	sourceGasPriceUSD, tokenPricesUSD, err := c.observePriceUpdates(ctx, c.lggr)
	if err != nil {
		return err
	}

	c.lggr.Infow("PriceWrite updatePrices",
		"sourceGasPriceUSD", sourceGasPriceUSD,
		"tokenPricesUSD", tokenPricesUSD,
	)

	return c.writePricesToDB(ctx, sourceGasPriceUSD, tokenPricesUSD)
}

func (c *priceWrite) observePriceUpdates(
	ctx context.Context,
	lggr logger.Logger,
) (sourceGasPriceUSD *big.Int, tokenPricesUSD map[cciptypes.Address]*big.Int, err error) {
	if c.destPriceRegistryReader == nil {
		return nil, nil, fmt.Errorf("skipping price update, destPriceRegistryReader not set yet")
	}
	// It is ok to lock/unlock here for a shorter critical section, the follow-up use for PriceRegistryReader is a call to decimals.
	// If a new PriceRegistryReader is set in between, either it continues to work or errors, both are acceptable
	sortedLaneTokens, filteredLaneTokens, err := ccipcommon.GetFilteredSortedLaneTokens(ctx, c.offRampReader, c.destPriceRegistryReader, c.priceGetter)

	lggr.Debugw("Filtered bridgeable tokens with no configured price getter", "filteredLaneTokens", filteredLaneTokens)

	if err != nil {
		return nil, nil, fmt.Errorf("get destination tokens: %w", err)
	}

	return c.generatePriceUpdates(ctx, lggr, sortedLaneTokens)
}

// All prices are USD ($1=1e18) denominated. All prices must be not nil.
// Return token prices should contain the exact same tokens as in tokenDecimals.
func (c *priceWrite) generatePriceUpdates(
	ctx context.Context,
	lggr logger.Logger,
	sortedLaneTokens []cciptypes.Address,
) (sourceGasPriceUSD *big.Int, tokenPricesUSD map[cciptypes.Address]*big.Int, err error) {
	// Include wrapped native in our token query as way to identify the source native USD price.
	// notice USD is in 1e18 scale, i.e. $1 = 1e18
	queryTokens := ccipcommon.FlattenUniqueSlice([]cciptypes.Address{c.sourceNative}, sortedLaneTokens)

	rawTokenPricesUSD, err := c.priceGetter.TokenPricesUSD(ctx, queryTokens)
	if err != nil {
		return nil, nil, err
	}
	lggr.Infow("Raw token prices", "rawTokenPrices", rawTokenPricesUSD)

	// make sure that we got prices for all the tokens of our query
	for _, token := range queryTokens {
		if rawTokenPricesUSD[token] == nil {
			return nil, nil, fmt.Errorf("missing token price: %+v", token)
		}
	}

	sourceNativePriceUSD, exists := rawTokenPricesUSD[c.sourceNative]
	if !exists {
		return nil, nil, fmt.Errorf("missing source native (%s) price", c.sourceNative)
	}

	destTokensDecimals, err := c.destPriceRegistryReader.GetTokensDecimals(ctx, sortedLaneTokens)
	if err != nil {
		return nil, nil, fmt.Errorf("get tokens decimals: %w", err)
	}

	tokenPricesUSD = make(map[cciptypes.Address]*big.Int, len(rawTokenPricesUSD))
	for i, token := range sortedLaneTokens {
		tokenPricesUSD[token] = calculateUsdPer1e18TokenAmount(rawTokenPricesUSD[token], destTokensDecimals[i])
	}

	if c.gasPriceEstimator == nil {
		return nil, nil, fmt.Errorf("skipping price update, gasPriceEstimator not set yet")
	}
	sourceGasPrice, err := c.gasPriceEstimator.GetGasPrice(ctx)
	if err != nil {
		return nil, nil, err
	}
	if sourceGasPrice == nil {
		return nil, nil, fmt.Errorf("missing gas price")
	}
	sourceGasPriceUSD, err = c.gasPriceEstimator.DenoteInUSD(sourceGasPrice, sourceNativePriceUSD)
	if err != nil {
		return nil, nil, err
	}

	lggr.Infow("Observing gas price", "observedGasPriceWei", sourceGasPrice, "observedGasPriceUSD", sourceGasPriceUSD)
	lggr.Infow("Observing token prices", "tokenPrices", tokenPricesUSD, "sourceNativePriceUSD", sourceNativePriceUSD)
	return sourceGasPriceUSD, tokenPricesUSD, nil
}

func (c *priceWrite) writePricesToDB(
	ctx context.Context,
	sourceGasPriceUSD *big.Int,
	tokenPricesUSD map[cciptypes.Address]*big.Int,
) (err error) {
	eg := new(errgroup.Group)

	if sourceGasPriceUSD != nil {
		eg.Go(func() error {
			return c.orm.InsertGasPricesForDestChain(ctx, c.destChainSelector, c.jobId, []cciporm.GasPriceUpdate{
				{
					SourceChainSelector: c.sourceChainSelector,
					GasPrice:            assets.NewWei(sourceGasPriceUSD),
				},
			})
		})
	}

	if tokenPricesUSD != nil {
		var tokenPrices []cciporm.TokenPriceUpdate

		for token, price := range tokenPricesUSD {
			tokenPrices = append(tokenPrices, cciporm.TokenPriceUpdate{
				TokenAddr:  string(token),
				TokenPrice: assets.NewWei(price),
			})
		}

		// Sort token by addr to make price updates ordering deterministic, easier to testing and debugging
		sort.Slice(tokenPrices, func(i, j int) bool {
			return tokenPrices[i].TokenAddr < tokenPrices[j].TokenAddr
		})

		eg.Go(func() error {
			return c.orm.InsertTokenPricesForDestChain(ctx, c.destChainSelector, c.jobId, tokenPrices)
		})
	}

	return eg.Wait()
}

// Input price is USD per full token, with 18 decimal precision
// Result price is USD per 1e18 of smallest token denomination, with 18 decimal precision
// Example: 1 USDC = 1.00 USD per full token, each full token is 6 decimals -> 1 * 1e18 * 1e18 / 1e6 = 1e30
func calculateUsdPer1e18TokenAmount(price *big.Int, decimals uint8) *big.Int {
	tmp := big.NewInt(0).Mul(price, big.NewInt(1e18))
	return tmp.Div(tmp, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
}
