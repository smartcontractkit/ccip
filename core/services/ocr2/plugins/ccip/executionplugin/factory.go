package executionplugin

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/custom_token_pool"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/cache"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
)

type Factory struct {
	// Config derived from job specs and does not change between instances.
	config StaticConfig

	destPriceRegReader ccipdata.PriceRegistryReader
	destPriceRegAddr   common.Address
	readersMu          *sync.Mutex
}

func newFactory(config StaticConfig) *Factory {
	return &Factory{
		config:    config,
		readersMu: &sync.Mutex{},
	}
}

func (rf *Factory) UpdateDynamicReaders(newPriceRegAddr common.Address) error {
	rf.readersMu.Lock()
	defer rf.readersMu.Unlock()
	// TODO: Investigate use of Close() to cleanup.
	// TODO: a true price registry upgrade on an existing lane may want some kind of start block in its config? Right now we
	// essentially assume that plugins don't care about historical price reg logs.
	if rf.destPriceRegAddr == newPriceRegAddr {
		// No-op
		return nil
	}
	// Close old reader (if present) and open new reader if address changed.
	if rf.destPriceRegReader != nil {
		if err := rf.destPriceRegReader.Close(); err != nil {
			return err
		}
	}
	destPriceRegistryReader, err := ccipdata.NewPriceRegistryReader(rf.config.lggr, newPriceRegAddr, rf.config.destLP, rf.config.destClient)
	if err != nil {
		return err
	}
	rf.destPriceRegReader = destPriceRegistryReader
	rf.destPriceRegAddr = newPriceRegAddr
	return nil
}

func (rf *Factory) NewReportingPlugin(config types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	destPriceRegistry, destWrappedNative, err := rf.config.offRampReader.ChangeConfig(config.OnchainConfig, config.OffchainConfig)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	// Open dynamic readers
	err = rf.UpdateDynamicReaders(destPriceRegistry)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}

	offchainConfig := rf.config.offRampReader.OffchainConfig()
	cachedSourceFeeTokens := cache.NewCachedFeeTokens(rf.config.sourceLP, rf.config.sourcePriceRegistry, int64(offchainConfig.SourceFinalityDepth))
	cachedDestTokens := cache.NewCachedSupportedTokens(rf.config.destLP, rf.config.offRampReader, rf.destPriceRegReader, int64(offchainConfig.DestOptimisticConfirmations))

	cachedTokenPools := cache.NewTokenPools(rf.config.lggr, rf.config.destLP, rf.config.offRampReader, int64(offchainConfig.DestOptimisticConfirmations), 5)

	return &ExecutionReportingPlugin{
			config:                rf.config,
			F:                     config.F,
			lggr:                  rf.config.lggr.Named("ExecutionReportingPlugin"),
			snoozedRoots:          cache.NewSnoozedRoots(rf.config.offRampReader.OnchainConfig().PermissionLessExecutionThresholdSeconds, offchainConfig.RootSnoozeTime.Duration()),
			inflightReports:       newInflightExecReportsContainer(offchainConfig.InflightCacheExpiry.Duration()),
			destPriceRegistry:     rf.destPriceRegReader,
			destWrappedNative:     destWrappedNative,
			onchainConfig:         rf.config.offRampReader.OnchainConfig(),
			offchainConfig:        offchainConfig,
			cachedDestTokens:      cachedDestTokens,
			cachedSourceFeeTokens: cachedSourceFeeTokens,
			cachedTokenPools:      cachedTokenPools,
			customTokenPoolFactory: func(ctx context.Context, poolAddress common.Address, contractBackend bind.ContractBackend) (custom_token_pool.CustomTokenPoolInterface, error) {
				return custom_token_pool.NewCustomTokenPool(poolAddress, contractBackend)
			},
			gasPriceEstimator: rf.config.offRampReader.GasPriceEstimator(),
		}, types.ReportingPluginInfo{
			Name: "CCIPExecution",
			// Setting this to false saves on calldata since OffRamp doesn't require agreement between NOPs
			// (OffRamp is only able to execute committed messages).
			UniqueReports: false,
			Limits: types.ReportingPluginLimits{
				MaxObservationLength: ccip.MaxObservationLength,
				MaxReportLength:      MaxExecutionReportLength,
			},
		}, nil
}
