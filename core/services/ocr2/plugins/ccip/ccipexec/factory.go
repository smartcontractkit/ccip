package ccipexec

import (
	"context"
	"fmt"
	"sync"

	"go.uber.org/multierr"

	"github.com/smartcontractkit/chainlink-common/pkg/services"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
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
	lggr               logger.Logger
	services           []services.Service
}

func (rf *ExecutionReportingPluginFactory) Name() string {
	return rf.lggr.Name()
}

// Start is used to run chainHealthcheck and tokenDataWorker, which were previously passed
// back to the delegate as top level job.ServiceCtx to be managed in core alongside the reporting
// plugin factory
func (rf *ExecutionReportingPluginFactory) Start(ctx context.Context) (err error) {
	rf.readersMu.Lock()
	defer rf.readersMu.Unlock()
	for _, service := range rf.services {
		serviceErr := service.Start(ctx)
		err = multierr.Append(err, serviceErr)
	}
	return
}

func (rf *ExecutionReportingPluginFactory) Close() (err error) {
	rf.readersMu.Lock()
	defer rf.readersMu.Unlock()
	for _, service := range rf.services {
		closeErr := service.Close()
		err = multierr.Append(err, closeErr)
	}

	return
}

func (rf *ExecutionReportingPluginFactory) Ready() error {
	return nil
}

func (rf *ExecutionReportingPluginFactory) HealthReport() map[string]error {
	return make(map[string]error)
}

func NewExecutionReportingPluginFactoryV2(ctx context.Context, lggr logger.Logger, sourceTokenAddress string, srcChainID int64, dstChainID int64, srcProvider commontypes.CCIPExecProvider, dstProvider commontypes.CCIPExecProvider) (*ExecutionReportingPluginFactory, error) {
	// NewOffRampReader doesn't need addr param when provided in job spec
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
	tokenDataProviders[cciptypes.Address(sourceTokenAddress)] = usdcReader
	if err2 != nil {
		// in order to not wire the attestation API through this factory, we wire it through the provider
		// when the provider is created. In some cases the attestation API can be nil, which means we
		// don't want any token data providers. This should not cause creating the job to fail, so we
		// give an empty map and move on.
		if err2.Error() != "empty USDC attestation API" {
			return nil, fmt.Errorf("new usdc reader: %w", err2)
		}
		tokenDataProviders = make(map[cciptypes.Address]tokendata.Reader)
	}

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
			logger.With(lggr, "onramp", offRampConfig.OnRamp,
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
			lggr:                          logger.Sugared(lggr),
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
		services:  []services.Service{chainHealthcheck, tokenBackgroundWorker},
		readersMu: &sync.Mutex{},
		lggr:      logger.Sugared(lggr),

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

func (rf *ExecutionReportingPluginFactory) UpdateDynamicReaders(_ context.Context, newPriceRegAddr cciptypes.Address) error {
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
	maxRetries := rf.config.newReportingPluginRetryConfig.MaxRetries

	pluginAndInfo, err := ccipcommon.RetryUntilSuccess(
		rf.NewReportingPluginFn(config), initialRetryDelay, maxDelay, maxRetries,
	)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	return pluginAndInfo.plugin, pluginAndInfo.pluginInfo, err
}

// NewReportingPluginFn implements the NewReportingPlugin logic. It is defined as a function so that it can easily be
// retried via RetryUntilSuccess. NewReportingPlugin must return successfully in order for the Exec plugin to function,
// hence why we can only keep retrying it until it succeeds.
func (rf *ExecutionReportingPluginFactory) NewReportingPluginFn(config types.ReportingPluginConfig) func() (reportingPluginAndInfo, error) {
	newReportingPluginFn := func() (reportingPluginAndInfo, error) {
		ctx := context.Background() // todo: consider setting a timeout

		// Start the chainHealthcheck and tokenDataWorker services
		// Using Start, while a bit more obtuse, allows us to manage these services
		// in the same process as the plugin factory in LOOP mode
		err := rf.Start(ctx)
		if err != nil {
			return reportingPluginAndInfo{}, err
		}

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

		batchingStrategy, err := NewBatchingStrategy(offchainConfig.BatchingStrategyID, rf.config.txmStatusChecker)
		if err != nil {
			return reportingPluginAndInfo{}, fmt.Errorf("get batching strategy: %w", err)
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
			commitRootsCache:            cache.NewCommitRootsCache(lggr, rf.config.commitStoreReader, msgVisibilityInterval, offchainConfig.RootSnoozeTime.Duration()),
			metricsCollector:            rf.config.metricsCollector,
			chainHealthcheck:            rf.config.chainHealthcheck,
			batchingStrategy:            batchingStrategy,
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

	return func() (reportingPluginAndInfo, error) {
		result, err := newReportingPluginFn()
		if err != nil {
			rf.config.lggr.Errorw("NewReportingPlugin failed", "err", err)
			rf.config.metricsCollector.NewReportingPluginError()
		}

		return result, err
	}
}
