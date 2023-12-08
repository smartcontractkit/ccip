package serviceprovider

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/observability"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/pricegetter"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"github.com/smartcontractkit/chainlink/v2/core/services/pipeline"
)

type EvmServiceProvider struct {
	mu                      *sync.Mutex
	pluginLabel             string
	LP                      logpoller.LogPoller
	EC                      client.Client
	priceGetter             pricegetter.PriceGetter
	PriceGetterArgs         *PriceGetterArgs
	onRampReader            ccipdata.OnRampReader
	OnRampReaderArgs        *OnRampReaderArgs
	commitStoreReader       ccipdata.CommitStoreReader
	CommitStoreReaderArgs   *CommitStoreReaderArgs
	offRampReader           ccipdata.OffRampReader
	OffRampReaderArgs       *OffRampReaderArgs
	priceRegistryReader     ccipdata.PriceRegistryReader
	PriceRegistryReaderArgs *PriceRegistryReaderArgs
}

func NewEvmServiceProvider(
	pluginLabel string,
	lp logpoller.LogPoller,
	ec client.Client,
	priceGetterArgs *PriceGetterArgs,
	onRampReaderArgs *OnRampReaderArgs,
	commitStoreReaderArgs *CommitStoreReaderArgs,
	offRampReaderArgs *OffRampReaderArgs,
	priceRegistryReaderArgs *PriceRegistryReaderArgs,
) *EvmServiceProvider {
	return &EvmServiceProvider{
		pluginLabel:             pluginLabel,
		mu:                      &sync.Mutex{},
		LP:                      lp,
		EC:                      ec,
		PriceGetterArgs:         priceGetterArgs,
		OnRampReaderArgs:        onRampReaderArgs,
		CommitStoreReaderArgs:   commitStoreReaderArgs,
		OffRampReaderArgs:       offRampReaderArgs,
		PriceRegistryReaderArgs: priceRegistryReaderArgs,
	}
}

type PriceGetterArgs struct {
	Source        string
	Runner        pipeline.Runner
	JobID         int32
	ExternalJobID uuid.UUID
	Name          string
	Lggr          logger.Logger
}

type OnRampReaderArgs struct {
	Lggr           logger.Logger
	SourceSelector uint64
	DestSelector   uint64
}

type CommitStoreReaderArgs struct {
	Lggr      logger.Logger
	Estimator gas.EvmFeeEstimator
}

type OffRampReaderArgs struct {
	Lggr      logger.Logger
	Estimator gas.EvmFeeEstimator
}

type PriceRegistryReaderArgs struct {
	Lggr logger.Logger
}

func (c *EvmServiceProvider) NewPriceGetter(ctx context.Context) (pricegetter.PriceGetter, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.priceGetter != nil {
		return c.priceGetter, nil
	}
	priceGetter, err := pricegetter.NewPipelineGetter(
		c.PriceGetterArgs.Source,
		c.PriceGetterArgs.Runner,
		c.PriceGetterArgs.JobID,
		c.PriceGetterArgs.ExternalJobID,
		c.PriceGetterArgs.Name,
		c.PriceGetterArgs.Lggr,
	)
	if err != nil {
		return nil, err
	}
	c.priceGetter = priceGetter
	return priceGetter, nil
}

func (c *EvmServiceProvider) NewPriceRegistryReader(ctx context.Context, addr common.Address) (ccipdata.PriceRegistryReader, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.priceRegistryReader != nil {
		existingAddr := c.priceRegistryReader.Address()
		if existingAddr == addr {
			return c.priceRegistryReader, nil
		}
	}

	// logpoller filters are registered on the constructor
	priceRegistryReader, err := ccipdata.NewPriceRegistryReader(
		c.PriceRegistryReaderArgs.Lggr,
		addr,
		c.LP,
		c.EC,
	)
	if err != nil {
		return nil, err
	}

	if c.priceRegistryReader != nil {
		if err := c.priceRegistryReader.Close(pg.WithParentCtx(ctx)); err != nil {
			return nil, err
		}
	}

	c.priceRegistryReader = observability.NewPriceRegistryReader(
		priceRegistryReader, c.EC.ConfiguredChainID().Int64(), c.pluginLabel)
	return c.priceRegistryReader, nil
}

func (c *EvmServiceProvider) NewOffRampReader(ctx context.Context, addr common.Address) (ccipdata.OffRampReader, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.offRampReader != nil {
		existingAddr := c.offRampReader.Address()
		if existingAddr == addr {
			return c.offRampReader, nil
		}
	}

	offRampReader, err := ccipdata.NewOffRampReader(
		c.OffRampReaderArgs.Lggr,
		addr,
		c.EC,
		c.LP,
		c.OffRampReaderArgs.Estimator,
	)
	if err != nil {
		return nil, err
	}

	if err := offRampReader.RegisterFilters(pg.WithParentCtx(ctx)); err != nil {
		return nil, err
	}

	if c.offRampReader != nil {
		if err := c.offRampReader.Close(pg.WithParentCtx(ctx)); err != nil {
			return nil, err
		}
	}

	c.offRampReader = observability.NewObservedOffRampReader(
		offRampReader, c.EC.ConfiguredChainID().Int64(), c.pluginLabel)
	return c.offRampReader, nil
}

func (c *EvmServiceProvider) NewCommitStoreReader(ctx context.Context, addr common.Address) (ccipdata.CommitStoreReader, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.commitStoreReader != nil {
		existingAddr := c.commitStoreReader.Address()
		if existingAddr == addr {
			return c.commitStoreReader, nil
		}
	}

	commitStoreReader, err := ccipdata.NewCommitStoreReader(
		c.CommitStoreReaderArgs.Lggr,
		addr,
		c.EC,
		c.LP,
		c.CommitStoreReaderArgs.Estimator,
	)
	if err != nil {
		return nil, err
	}

	if err := commitStoreReader.RegisterFilters(pg.WithParentCtx(ctx)); err != nil {
		return nil, err
	}

	if c.commitStoreReader != nil {
		if err := c.commitStoreReader.Close(pg.WithParentCtx(ctx)); err != nil {
			return nil, err
		}
	}

	c.commitStoreReader = observability.NewObservedCommitStoreReader(
		commitStoreReader, c.EC.ConfiguredChainID().Int64(), c.pluginLabel)
	return c.commitStoreReader, nil
}

func (c *EvmServiceProvider) NewOnRampReader(ctx context.Context, addr common.Address) (ccipdata.OnRampReader, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.onRampReader != nil {
		existingAddr, err := c.onRampReader.Address()
		if err != nil {
			return nil, err
		}
		if existingAddr == addr {
			return c.onRampReader, nil
		}
	}

	onRampReader, err := ccipdata.NewOnRampReader(
		c.OnRampReaderArgs.Lggr,
		c.OnRampReaderArgs.SourceSelector,
		c.OnRampReaderArgs.DestSelector,
		addr,
		c.LP,
		c.EC,
	)
	if err != nil {
		return nil, err
	}

	if err := onRampReader.RegisterFilters(pg.WithParentCtx(ctx)); err != nil {
		return nil, err
	}

	if c.onRampReader != nil {
		if err := c.onRampReader.Close(pg.WithParentCtx(ctx)); err != nil {
			return nil, err
		}
	}

	c.onRampReader = observability.NewObservedOnRampReader(
		onRampReader, c.EC.ConfiguredChainID().Int64(), c.pluginLabel)
	return c.onRampReader, nil
}
