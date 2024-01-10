package ccipcommit

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	chainselectors "github.com/smartcontractkit/chain-selectors"
	libocr2 "github.com/smartcontractkit/libocr/offchainreporting2plus"

	commonlogger "github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipcommon"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr"
	"github.com/smartcontractkit/chainlink/v2/core/chains/legacyevm"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/ccipdataprovider"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/factory"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/observability"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/oraclelib"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/pricegetter"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/promwrapper"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"github.com/smartcontractkit/chainlink/v2/core/services/pipeline"
)

func jobSpecToCommitPluginConfig(lggr logger.Logger, jb job.Job, pr pipeline.Runner, chainSet legacyevm.LegacyChainContainer, qopts ...pg.QOpt) (*CommitPluginStaticConfig, *ccipcommon.BackfillArgs, error) {
	if jb.OCR2OracleSpec == nil {
		return nil, nil, errors.New("spec is nil")
	}
	spec := jb.OCR2OracleSpec
	var pluginConfig ccipconfig.CommitPluginJobSpecConfig
	err := json.Unmarshal(spec.PluginConfig.Bytes(), &pluginConfig)
	if err != nil {
		return nil, nil, err
	}

	destChain, destChainID, err := ccipconfig.GetChainFromSpec(spec, chainSet)
	if err != nil {
		return nil, nil, err
	}

	commitStoreAddress := common.HexToAddress(spec.ContractID)
	staticConfig, err := ccipdata.FetchCommitStoreStaticConfig(commitStoreAddress, destChain.Client())
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed getting the static config from the commitStore")
	}
	sourceChain, sourceChainID, err := ccipconfig.GetChainByChainSelector(chainSet, staticConfig.SourceChainSelector)
	if err != nil {
		return nil, nil, err
	}
	commitStoreReader, err := factory.NewCommitStoreReader(lggr, commitStoreAddress, destChain.Client(), destChain.LogPoller(), sourceChain.GasEstimator())
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not create commitStore reader")
	}
	sourceChainName, destChainName, err := ccipconfig.ResolveChainNames(sourceChainID, destChainID)
	if err != nil {
		return nil, nil, err
	}
	commitLggr := lggr.Named("CCIPCommit").With("sourceChain", sourceChainName, "destChain", destChainName)
	pipelinePriceGetter, err := pricegetter.NewPipelineGetter(pluginConfig.TokenPricesUSDPipeline, pr, jb.ID, jb.ExternalJobID, jb.Name.ValueOrZero(), lggr)
	if err != nil {
		return nil, nil, err
	}

	// Load all the readers relevant for this plugin.
	onRampReader, err := factory.NewOnRampReader(commitLggr, staticConfig.SourceChainSelector, staticConfig.ChainSelector, staticConfig.OnRamp, sourceChain.LogPoller(), sourceChain.Client())
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed onramp reader")
	}
	offRampReader, err := factory.NewOffRampReader(commitLggr, pluginConfig.OffRamp, destChain.Client(), destChain.LogPoller(), destChain.GasEstimator())
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed offramp reader")
	}
	onRampRouterAddr, err := onRampReader.RouterAddress()
	if err != nil {
		return nil, nil, err
	}
	sourceRouter, err := router.NewRouter(onRampRouterAddr, sourceChain.Client())
	if err != nil {
		return nil, nil, err
	}
	sourceNative, err := sourceRouter.GetWrappedNative(nil)
	if err != nil {
		return nil, nil, err
	}

	// Prom wrappers
	onRampReader = observability.NewObservedOnRampReader(onRampReader, sourceChainID, ccip.CommitPluginLabel)
	offRampReader = observability.NewObservedOffRampReader(offRampReader, destChainID, ccip.CommitPluginLabel)
	commitStoreReader = observability.NewObservedCommitStoreReader(commitStoreReader, destChainID, ccip.CommitPluginLabel)
	metricsCollector := ccip.NewPluginMetricsCollector(ccip.CommitPluginLabel, big.NewInt(sourceChainID), big.NewInt(destChainID))

	lggr.Infow("NewCommitServices",
		"pluginConfig", pluginConfig,
		"staticConfig", staticConfig,
		// TODO bring back
		//"dynamicOnRampConfig", dynamicOnRampConfig,
		"sourceNative", sourceNative,
		"sourceRouter", sourceRouter.Address())
	return &CommitPluginStaticConfig{
			lggr:                  commitLggr,
			onRampReader:          onRampReader,
			offRamp:               offRampReader,
			priceGetter:           pipelinePriceGetter,
			sourceNative:          sourceNative,
			sourceChainSelector:   staticConfig.SourceChainSelector,
			destChainSelector:     staticConfig.ChainSelector,
			commitStore:           commitStoreReader,
			priceRegistryProvider: ccipdataprovider.NewEvmPriceRegistry(destChain.LogPoller(), destChain.Client(), commitLggr, ccip.CommitPluginLabel),
			metricsCollector:      metricsCollector,
		}, &ccipcommon.BackfillArgs{
			SourceLP:         sourceChain.LogPoller(),
			DestLP:           destChain.LogPoller(),
			SourceStartBlock: pluginConfig.SourceStartBlock,
			DestStartBlock:   pluginConfig.DestStartBlock,
		}, nil
}

func NewCommitServices(ctx context.Context, lggr logger.Logger, jb job.Job, chainSet legacyevm.LegacyChainContainer, new bool, pr pipeline.Runner, argsNoPlugin libocr2.OCR2OracleArgs, logError func(string), qopts ...pg.QOpt) ([]job.ServiceCtx, error) {
	pluginConfig, backfillArgs, err := jobSpecToCommitPluginConfig(lggr, jb, pr, chainSet, qopts...)
	if err != nil {
		return nil, err
	}
	wrappedPluginFactory := NewCommitReportingPluginFactory(*pluginConfig)

	if err1 := pluginConfig.onRampReader.RegisterFilters(qopts...); err1 != nil {
		return nil, err1
	}

	if err1 := pluginConfig.commitStore.RegisterFilters(qopts...); err1 != nil {
		return nil, err1
	}

	if err1 := pluginConfig.offRamp.RegisterFilters(qopts...); err1 != nil {
		return nil, err1
	}

	destChainID, err := chainselectors.ChainIdFromSelector(pluginConfig.destChainSelector)
	if err != nil {
		return nil, err
	}
	argsNoPlugin.ReportingPluginFactory = promwrapper.NewPromFactory(wrappedPluginFactory, "CCIPCommit", jb.OCR2OracleSpec.Relay, big.NewInt(0).SetUint64(destChainID))
	argsNoPlugin.Logger = commonlogger.NewOCRWrapper(pluginConfig.lggr, true, logError)
	oracle, err := libocr2.NewOracle(argsNoPlugin)
	if err != nil {
		return nil, err
	}
	// If this is a brand-new job, then we make use of the start blocks. If not then we're rebooting and log poller will pick up where we left off.
	if new {
		return []job.ServiceCtx{oraclelib.NewBackfilledOracle(
			pluginConfig.lggr,
			backfillArgs.SourceLP,
			backfillArgs.DestLP,
			backfillArgs.SourceStartBlock,
			backfillArgs.DestStartBlock,
			job.NewServiceAdapter(oracle)),
		}, nil
	}
	return []job.ServiceCtx{job.NewServiceAdapter(oracle)}, nil
}

func CommitReportToEthTxMeta(typ ccipconfig.ContractType, ver semver.Version) (func(report []byte) (*txmgr.TxMeta, error), error) {
	return factory.CommitReportToEthTxMeta(typ, ver)
}

// UnregisterCommitPluginLpFilters unregisters all the registered filters for both source and dest chains.
// NOTE: The transaction MUST be used here for CLO's monster tx to function as expected
// https://github.com/smartcontractkit/ccip/blob/68e2197472fb017dd4e5630d21e7878d58bc2a44/core/services/feeds/service.go#L716
// TODO once that transaction is broken up, we should be able to simply rely on oracle.Close() to cleanup the filters.
// Until then we have to deterministically reload the readers from the spec (and thus their filters) and close them.
func UnregisterCommitPluginLpFilters(ctx context.Context, lggr logger.Logger, jb job.Job, pr pipeline.Runner, chainSet legacyevm.LegacyChainContainer, qopts ...pg.QOpt) error {
	commitPluginConfig, _, err := jobSpecToCommitPluginConfig(lggr, jb, pr, chainSet)
	if err != nil {
		return errors.New("spec is nil")
	}
	if err := commitPluginConfig.onRampReader.Close(qopts...); err != nil {
		return err
	}
	if err := commitPluginConfig.commitStore.Close(qopts...); err != nil {
		return err
	}
	return commitPluginConfig.offRamp.Close(qopts...)
}
