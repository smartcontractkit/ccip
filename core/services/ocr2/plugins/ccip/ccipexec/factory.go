package ccipexec

import (
	"context"
	"fmt"
	"sync"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/observability"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/tokendata"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipcommon"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/cache"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
)

type ExecutionReportingPluginFactory struct {
	// Config derived from job specs and does not change between instances.
	config ExecutionPluginStaticConfig

	destPriceRegReader ccipdata.PriceRegistryReader
	destPriceRegAddr   cciptypes.Address
	readersMu          *sync.Mutex
}

func (rf *ExecutionReportingPluginFactory) Name() string {
	//TODO implement me
	panic("implement me")
}

func (rf *ExecutionReportingPluginFactory) Start(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (rf *ExecutionReportingPluginFactory) Close() error {
	//TODO implement me
	panic("implement me")
}

func (rf *ExecutionReportingPluginFactory) Ready() error {
	//TODO implement me
	panic("implement me")
}

func (rf *ExecutionReportingPluginFactory) HealthReport() map[string]error {
	//TODO implement me
	panic("implement me")
}

func NewExecutionReportingPluginFactoryV2(ctx context.Context, lggr logger.Logger, sourceTokenAddress string, srcChainID int64, dstChainID int64, srcProvider commontypes.CCIPExecProvider, dstProvider commontypes.CCIPExecProvider) (*ExecutionReportingPluginFactory, error) {
	// TODO: common logger is a subset of core logger.
	// what's the golden path for passing a logger through from the plugin to the LOOP reporting plugin factory?
	if lggr == nil {
		lggr, _ = logger.NewLogger()
	}

	// TODO: NewOffRampReader doesn't need addr param when provided in job spec
	offRampReader, err := dstProvider.NewOffRampReader(ctx, "")
	if err != nil {
		return nil, fmt.Errorf("create offRampReader: %w", err)
	}

	offRampConfig, err := offRampReader.GetStaticConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("get offRamp static config: %w", err)
	}

	srcChainSelector := offRampConfig.SourceChainSelector
	dstChainSelector := offRampConfig.ChainSelector
	onRampReader, err := srcProvider.NewOnRampReader(ctx, offRampConfig.OnRamp, srcChainSelector, dstChainSelector)
	if err != nil {
		return nil, fmt.Errorf("create onRampReader: %w", err)
	}

	dynamicOnRampConfig, err := onRampReader.GetDynamicConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("get onramp dynamic config: %w", err)
	}

	sourceWrappedNative, err := srcProvider.SourceNativeToken(ctx, dynamicOnRampConfig.Router)
	if err != nil {
		return nil, fmt.Errorf("get source wrapped native token: %w", err)
	}

	srcCommitStore, err := srcProvider.NewCommitStoreReader(ctx, offRampConfig.CommitStore)
	if err != nil {
		return nil, fmt.Errorf("could not create src commitStoreReader reader: %w", err)
	}

	dstCommitStore, err := dstProvider.NewCommitStoreReader(ctx, offRampConfig.CommitStore)
	if err != nil {
		return nil, fmt.Errorf("could not create dst commitStoreReader reader: %w", err)
	}

	var commitStoreReader ccipdata.CommitStoreReader
	commitStoreReader = ccip.NewProviderProxyCommitStoreReader(srcCommitStore, dstCommitStore)

	tokenDataProviders := make(map[cciptypes.Address]tokendata.Reader)
	// init usdc token data provider
	usdcReader, err2 := srcProvider.NewTokenDataReader(ctx, "")
	if err2 != nil {
		return nil, fmt.Errorf("new usdc reader: %w", err2)
	}
	tokenDataProviders[cciptypes.Address(sourceTokenAddress)] = usdcReader

	// Prom wrappers
	onRampReader = observability.NewObservedOnRampReader(onRampReader, srcChainID, ccip.ExecPluginLabel)
	commitStoreReader = observability.NewObservedCommitStoreReader(commitStoreReader, dstChainID, ccip.ExecPluginLabel)
	offRampReader = observability.NewObservedOffRampReader(offRampReader, dstChainID, ccip.ExecPluginLabel)
	metricsCollector := ccip.NewPluginMetricsCollector(ccip.ExecPluginLabel, srcChainID, dstChainID)

	tokenPoolBatchedReader, err := dstProvider.NewTokenPoolBatchedReader(ctx, "", srcChainSelector)
	if err != nil {
		return nil, fmt.Errorf("new token pool batched reader: %w", err)
	}

	chainHealthcheck := cache.NewObservedChainHealthCheck(
		cache.NewChainHealthcheck(
			// Adding more details to Logger to make healthcheck logs more informative
			// It's safe because healthcheck logs only in case of unhealthy state
			lggr.With(
				"onramp", offRampConfig.OnRamp,
				"commitStore", offRampConfig.CommitStore,
			),
			onRampReader,
			commitStoreReader,
		),
		ccip.ExecPluginLabel,
		srcChainID,
		dstChainID,
		offRampConfig.OnRamp,
	)

	tokenBackgroundWorker := tokendata.NewBackgroundWorker(
		tokenDataProviders,
		tokenDataWorkerNumWorkers,
		tokenDataWorkerTimeout,
		2*tokenDataWorkerTimeout,
	)

	return &ExecutionReportingPluginFactory{
		config: ExecutionPluginStaticConfig{
			lggr:                          lggr,
			onRampReader:                  onRampReader,
			commitStoreReader:             commitStoreReader,
			offRampReader:                 offRampReader,
			sourcePriceRegistryProvider:   ccip.NewChainAgnosticPriceRegistry(srcProvider),
			sourceWrappedNativeToken:      sourceWrappedNative,
			destChainSelector:             dstChainSelector,
			priceRegistryProvider:         ccip.NewChainAgnosticPriceRegistry(dstProvider),
			tokenPoolBatchedReader:        tokenPoolBatchedReader,
			tokenDataWorker:               tokenBackgroundWorker,
			metricsCollector:              metricsCollector,
			chainHealthcheck:              chainHealthcheck,
			newReportingPluginRetryConfig: defaultNewReportingPluginRetryConfig,
		},
		readersMu: &sync.Mutex{},

		// the fields below are initially empty and populated on demand
		destPriceRegReader: nil,
		destPriceRegAddr:   "",
	}, nil
}

func NewExecutionReportingPluginFactory(config ExecutionPluginStaticConfig) *ExecutionReportingPluginFactory {
	return &ExecutionReportingPluginFactory{
		config:    config,
		readersMu: &sync.Mutex{},

		// the fields below are initially empty and populated on demand
		destPriceRegReader: nil,
		destPriceRegAddr:   "",
	}
}

func (rf *ExecutionReportingPluginFactory) UpdateDynamicReaders(ctx context.Context, newPriceRegAddr cciptypes.Address) error {
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

	destPriceRegistryReader, err := rf.config.priceRegistryProvider.NewPriceRegistryReader(context.Background(), newPriceRegAddr)
	if err != nil {
		return err
	}
	rf.destPriceRegReader = destPriceRegistryReader
	rf.destPriceRegAddr = newPriceRegAddr
	return nil
}

type reportingPluginAndInfo struct {
	plugin     types.ReportingPlugin
	pluginInfo types.ReportingPluginInfo
}

// NewReportingPlugin registers a new ReportingPlugin
func (rf *ExecutionReportingPluginFactory) NewReportingPlugin(config types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	initialRetryDelay := rf.config.newReportingPluginRetryConfig.InitialDelay
	maxDelay := rf.config.newReportingPluginRetryConfig.MaxDelay

	pluginAndInfo, err := ccipcommon.RetryUntilSuccess(rf.NewReportingPluginFn(config), initialRetryDelay, maxDelay)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	return pluginAndInfo.plugin, pluginAndInfo.pluginInfo, err
}

// NewReportingPluginFn implements the NewReportingPlugin logic. It is defined as a function so that it can easily be
// retried via RetryUntilSuccess. NewReportingPlugin must return successfully in order for the Exec plugin to function,
// hence why we can only keep retrying it until it succeeds.
func (rf *ExecutionReportingPluginFactory) NewReportingPluginFn(config types.ReportingPluginConfig) func() (reportingPluginAndInfo, error) {
	return func() (reportingPluginAndInfo, error) {
		ctx := context.Background() // todo: consider setting a timeout

		destPriceRegistry, destWrappedNative, err := rf.config.offRampReader.ChangeConfig(ctx, config.OnchainConfig, config.OffchainConfig)
		if err != nil {
			return reportingPluginAndInfo{}, err
		}

		// Open dynamic readers
		err = rf.UpdateDynamicReaders(ctx, destPriceRegistry)
		if err != nil {
			return reportingPluginAndInfo{}, err
		}

		offchainConfig, err := rf.config.offRampReader.OffchainConfig(ctx)
		if err != nil {
			return reportingPluginAndInfo{}, fmt.Errorf("get offchain config from offramp: %w", err)
		}

		gasPriceEstimator, err := rf.config.offRampReader.GasPriceEstimator(ctx)
		if err != nil {
			return reportingPluginAndInfo{}, fmt.Errorf("get gas price estimator from offramp: %w", err)
		}

		onchainConfig, err := rf.config.offRampReader.OnchainConfig(ctx)
		if err != nil {
			return reportingPluginAndInfo{}, fmt.Errorf("get onchain config from offramp: %w", err)
		}

		msgVisibilityInterval := offchainConfig.MessageVisibilityInterval.Duration()
		if msgVisibilityInterval.Seconds() == 0 {
			rf.config.lggr.Info("MessageVisibilityInterval not set, falling back to PermissionLessExecutionThreshold")
			msgVisibilityInterval = onchainConfig.PermissionLessExecutionThresholdSeconds
		}
		rf.config.lggr.Infof("MessageVisibilityInterval set to: %s", msgVisibilityInterval)

		lggr := rf.config.lggr.Named("ExecutionReportingPlugin")
		plugin := &ExecutionReportingPlugin{
			F:                           config.F,
			lggr:                        lggr,
			offchainConfig:              offchainConfig,
			tokenDataWorker:             rf.config.tokenDataWorker,
			gasPriceEstimator:           gasPriceEstimator,
			sourcePriceRegistryProvider: rf.config.sourcePriceRegistryProvider,
			sourcePriceRegistryLock:     sync.RWMutex{},
			sourceWrappedNativeToken:    rf.config.sourceWrappedNativeToken,
			onRampReader:                rf.config.onRampReader,
			commitStoreReader:           rf.config.commitStoreReader,
			destPriceRegistry:           rf.destPriceRegReader,
			destWrappedNative:           destWrappedNative,
			onchainConfig:               onchainConfig,
			offRampReader:               rf.config.offRampReader,
			tokenPoolBatchedReader:      rf.config.tokenPoolBatchedReader,
			inflightReports:             newInflightExecReportsContainer(offchainConfig.InflightCacheExpiry.Duration()),
			commitRootsCache:            cache.NewCommitRootsCache(lggr, msgVisibilityInterval, offchainConfig.RootSnoozeTime.Duration()),
			metricsCollector:            rf.config.metricsCollector,
			chainHealthcheck:            rf.config.chainHealthcheck,
		}

		pluginInfo := types.ReportingPluginInfo{
			Name: "CCIPExecution",
			// Setting this to false saves on calldata since OffRamp doesn't require agreement between NOPs
			// (OffRamp is only able to execute committed messages).
			UniqueReports: false,
			Limits: types.ReportingPluginLimits{
				MaxObservationLength: ccip.MaxObservationLength,
				MaxReportLength:      MaxExecutionReportLength,
			},
		}

		return reportingPluginAndInfo{plugin, pluginInfo}, nil
	}
}
