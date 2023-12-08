package ccip

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	libocr2 "github.com/smartcontractkit/libocr/offchainreporting2plus"

	relaylogger "github.com/smartcontractkit/chainlink-relay/pkg/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/serviceprovider"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/oraclelib"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/promwrapper"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"github.com/smartcontractkit/chainlink/v2/core/services/pipeline"
)

type BackfillArgs struct {
	sourceLP, destLP                 logpoller.LogPoller
	sourceStartBlock, destStartBlock int64
}

func parseJobSpec(lggr logger.Logger, jb job.Job, pr pipeline.Runner, chainSet evm.LegacyChainContainer, qopts ...pg.QOpt) (
	logger.Logger, *serviceprovider.EvmServiceProvider, *serviceprovider.EvmServiceProvider, *CommitPluginStaticConfig, *BackfillArgs, error) {
	if jb.OCR2OracleSpec == nil {
		return nil, nil, nil, nil, nil, errors.New("spec is nil")
	}
	spec := jb.OCR2OracleSpec
	var pluginConfig ccipconfig.CommitPluginJobSpecConfig
	err := json.Unmarshal(spec.PluginConfig.Bytes(), &pluginConfig)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	destChain, destChainID, err := ccipconfig.GetChainFromSpec(spec, chainSet)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	commitStoreAddress := common.HexToAddress(spec.ContractID)
	staticConfig, err := ccipdata.FetchCommitStoreStaticConfig(commitStoreAddress, destChain.Client())
	if err != nil {
		return nil, nil, nil, nil, nil, errors.Wrap(err, "failed getting the static config from the commitStore")
	}
	sourceChain, sourceChainID, err := ccipconfig.GetChainByChainSelector(chainSet, staticConfig.SourceChainSelector)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	commitLggr := lggr.Named("CCIPCommit").With(
		"sourceChain", ChainName(sourceChainID),
		"destChain", ChainName(destChainID))

	sourceServiceProvider := serviceprovider.NewEvmServiceProvider(
		CommitPluginLabel,
		sourceChain.LogPoller(),
		sourceChain.Client(),
		&serviceprovider.PriceGetterArgs{
			Source:        pluginConfig.TokenPricesUSDPipeline,
			Runner:        pr,
			JobID:         jb.ID,
			ExternalJobID: jb.ExternalJobID,
			Name:          jb.Name.ValueOrZero(),
			Lggr:          commitLggr,
		},
		&serviceprovider.OnRampReaderArgs{
			Lggr:           commitLggr,
			SourceSelector: staticConfig.SourceChainSelector,
			DestSelector:   staticConfig.ChainSelector,
		},
		nil,
		nil,
		nil,
	)
	onRampReader, err := sourceServiceProvider.NewOnRampReader(context.Background(), staticConfig.OnRamp)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	destServiceProvider := serviceprovider.NewEvmServiceProvider(
		CommitPluginLabel,
		destChain.LogPoller(),
		destChain.Client(),
		nil,
		nil,
		&serviceprovider.CommitStoreReaderArgs{Lggr: commitLggr, Estimator: sourceChain.GasEstimator()},
		&serviceprovider.OffRampReaderArgs{Lggr: commitLggr, Estimator: destChain.GasEstimator()},
		&serviceprovider.PriceRegistryReaderArgs{Lggr: commitLggr},
	)

	onRampRouterAddr, err := onRampReader.RouterAddress()
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	sourceRouter, err := router.NewRouter(onRampRouterAddr, sourceChain.Client())
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	sourceNative, err := sourceRouter.GetWrappedNative(nil)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	lggr.Infow("NewCommitServices",
		"pluginConfig", pluginConfig,
		"staticConfig", staticConfig,
		// TODO bring back
		//"dynamicOnRampConfig", dynamicOnRampConfig,
		"sourceNative", sourceNative,
		"sourceRouter", sourceRouter.Address())

	return commitLggr,
		sourceServiceProvider,
		destServiceProvider,
		&CommitPluginStaticConfig{
			OnRampAddress:       staticConfig.OnRamp,
			OffRampAddress:      common.HexToAddress(pluginConfig.OffRamp),
			CommitStoreAddress:  commitStoreAddress,
			SourceChainSelector: staticConfig.SourceChainSelector,
			SourceNative:        sourceNative,
			DestChainID:         big.NewInt(destChainID),
		},
		&BackfillArgs{
			sourceLP:         sourceChain.LogPoller(),
			destLP:           destChain.LogPoller(),
			sourceStartBlock: pluginConfig.SourceStartBlock,
			destStartBlock:   pluginConfig.DestStartBlock,
		},
		nil
}

func NewCommitServices(lggr logger.Logger, jb job.Job, chainSet evm.LegacyChainContainer, new bool, pr pipeline.Runner, argsNoPlugin libocr2.OCR2OracleArgs, logError func(string), qopts ...pg.QOpt) ([]job.ServiceCtx, error) {
	ctx := context.Background()

	lggr, sourceServiceProvider, destServiceProvider, pluginConfig, backfillArgs, err := parseJobSpec(lggr, jb, pr, chainSet, qopts...)
	if err != nil {
		return nil, err
	}

	// preload the readers - ensures filters are registered
	if _, err := sourceServiceProvider.NewPriceGetter(ctx); err != nil {
		return nil, fmt.Errorf("init price getter: %w", err)
	}
	if _, err := sourceServiceProvider.NewOnRampReader(ctx, pluginConfig.OnRampAddress); err != nil {
		return nil, fmt.Errorf("init on ramp reader: %w", err)
	}
	if _, err := destServiceProvider.NewCommitStoreReader(ctx, pluginConfig.CommitStoreAddress); err != nil {
		return nil, fmt.Errorf("init commit store reader: %w", err)
	}
	if _, err := destServiceProvider.NewOffRampReader(ctx, pluginConfig.OffRampAddress); err != nil {
		return nil, fmt.Errorf("init off ramp reader: %w", err)
	}

	commitPluginFactory := NewCommitReportingPluginFactory(
		lggr,
		*pluginConfig,
		sourceServiceProvider,
		destServiceProvider,
		sourceServiceProvider,
	)

	argsNoPlugin.ReportingPluginFactory = promwrapper.NewPromFactory(
		commitPluginFactory, "CCIPCommit", jb.OCR2OracleSpec.Relay, pluginConfig.DestChainID)
	argsNoPlugin.Logger = relaylogger.NewOCRWrapper(lggr, true, logError)
	oracle, err := libocr2.NewOracle(argsNoPlugin)
	if err != nil {
		return nil, err
	}
	// If this is a brand-new job, then we make use of the start blocks. If not then we're rebooting and log poller will pick up where we left off.
	if new {
		return []job.ServiceCtx{oraclelib.NewBackfilledOracle(
			lggr,
			backfillArgs.sourceLP,
			backfillArgs.destLP,
			backfillArgs.sourceStartBlock,
			backfillArgs.destStartBlock,
			job.NewServiceAdapter(oracle)),
		}, nil
	}
	return []job.ServiceCtx{job.NewServiceAdapter(oracle)}, nil
}

func CommitReportToEthTxMeta(typ ccipconfig.ContractType, ver semver.Version) (func(report []byte) (*txmgr.TxMeta, error), error) {
	return ccipdata.CommitReportToEthTxMeta(typ, ver)
}

// UnregisterCommitPluginLpFilters unregisters all the registered filters for both source and dest chains.
// NOTE: The transaction MUST be used here for CLO's monster tx to function as expected
// https://github.com/smartcontractkit/ccip/blob/68e2197472fb017dd4e5630d21e7878d58bc2a44/core/services/feeds/service.go#L716
// TODO once that transaction is broken up, we should be able to simply rely on oracle.Close() to cleanup the filters.
// Until then we have to deterministically reload the readers from the spec (and thus their filters) and close them.
func UnregisterCommitPluginLpFilters(ctx context.Context, lggr logger.Logger, jb job.Job, pr pipeline.Runner, chainSet evm.LegacyChainContainer, qopts ...pg.QOpt) error {
	_, sourceProvider, destProvider, commitPluginConfig, _, err := parseJobSpec(lggr, jb, pr, chainSet)
	if err != nil {
		return fmt.Errorf("parse job spec: %w", err)
	}

	onRampReader, err := ccipdata.NewOnRampReader(
		sourceProvider.OnRampReaderArgs.Lggr,
		sourceProvider.OnRampReaderArgs.SourceSelector,
		sourceProvider.OnRampReaderArgs.DestSelector,
		commitPluginConfig.OnRampAddress,
		sourceProvider.LP,
		sourceProvider.EC,
	)
	if err != nil {
		return err
	}
	if err := onRampReader.Close(qopts...); err != nil {
		return err
	}

	commitStoreReader, err := ccipdata.NewCommitStoreReader(
		destProvider.CommitStoreReaderArgs.Lggr,
		commitPluginConfig.CommitStoreAddress,
		destProvider.EC,
		destProvider.LP,
		destProvider.CommitStoreReaderArgs.Estimator,
	)
	if err != nil {
		return err
	}
	if err := commitStoreReader.Close(qopts...); err != nil {
		return err
	}

	offRampReader, err := ccipdata.NewOffRampReader(
		destProvider.OffRampReaderArgs.Lggr,
		commitPluginConfig.OffRampAddress,
		destProvider.EC,
		destProvider.LP,
		destProvider.OffRampReaderArgs.Estimator,
	)
	if err != nil {
		return err
	}
	return offRampReader.Close(qopts...)
}
