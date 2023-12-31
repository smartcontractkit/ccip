package ccipexec

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/url"
	"strconv"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	chainselectors "github.com/smartcontractkit/chain-selectors"
	libocr2 "github.com/smartcontractkit/libocr/offchainreporting2plus"

	commonlogger "github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr"
	"github.com/smartcontractkit/chainlink/v2/core/chains/legacyevm"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipcommon"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/ccipdataprovider"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/factory"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/observability"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/oraclelib"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/tokendata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/tokendata/usdc"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/promwrapper"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

const numTokenDataWorkers = 5

func jobSpecToExecPluginConfig(ctx context.Context, lggr logger.Logger, jb job.Job, chainSet legacyevm.LegacyChainContainer) (*ExecutionPluginStaticConfig, *ccipcommon.BackfillArgs, error) {
	if jb.OCR2OracleSpec == nil {
		return nil, nil, errors.New("spec is nil")
	}
	spec := jb.OCR2OracleSpec
	var pluginConfig ccipconfig.ExecutionPluginJobSpecConfig
	err := json.Unmarshal(spec.PluginConfig.Bytes(), &pluginConfig)
	if err != nil {
		return nil, nil, err
	}

	destChain, destChainID, err := ccipconfig.GetChainFromSpec(spec, chainSet)
	if err != nil {
		return nil, nil, err
	}

	// Create the offRamp reader.
	offRampAddress := common.HexToAddress(spec.ContractID)
	offRampReader, err := factory.NewOffRampReader(lggr, offRampAddress, destChain.Client(), destChain.LogPoller(), destChain.GasEstimator())
	if err != nil {
		return nil, nil, errors.Wrap(err, "create offRampReader")
	}
	offRampConfig, err := offRampReader.GetStaticConfig(context.Background())
	if err != nil {
		return nil, nil, errors.Wrap(err, "get offRamp static config")
	}

	chainID, err := chainselectors.ChainIdFromSelector(offRampConfig.SourceChainSelector)
	if err != nil {
		return nil, nil, err
	}

	sourceChain, err := chainSet.Get(strconv.FormatUint(chainID, 10))
	if err != nil {
		return nil, nil, errors.Wrap(err, "open source chain")
	}

	sourceChainName, destChainName, err := ccipconfig.ResolveChainNames(int64(chainID), destChainID)
	if err != nil {
		return nil, nil, err
	}
	execLggr := lggr.Named("CCIPExecution").With("sourceChain", sourceChainName, "destChain", destChainName)
	onRampReader, err := factory.NewOnRampReader(execLggr, offRampConfig.SourceChainSelector, offRampConfig.ChainSelector, offRampConfig.OnRamp, sourceChain.LogPoller(), sourceChain.Client())
	if err != nil {
		return nil, nil, errors.Wrap(err, "create onramp reader")
	}
	dynamicOnRampConfig, err := onRampReader.GetDynamicConfig()
	if err != nil {
		return nil, nil, errors.Wrap(err, "get onramp dynamic config")
	}

	sourceRouter, err := router.NewRouter(dynamicOnRampConfig.Router, sourceChain.Client())
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed loading source router")
	}
	sourceWrappedNative, err := sourceRouter.GetWrappedNative(&bind.CallOpts{})
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not get source native token")
	}

	// TODO: we don't support onramp source registry changes without a reboot yet?
	sourcePriceRegistry, err := factory.NewPriceRegistryReader(lggr, dynamicOnRampConfig.PriceRegistry, sourceChain.LogPoller(), sourceChain.Client())
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not load source registry")
	}

	commitStoreReader, err := factory.NewCommitStoreReader(lggr, offRampConfig.CommitStore, destChain.Client(), destChain.LogPoller(), sourceChain.GasEstimator())
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not load commitStoreReader reader")
	}

	tokenDataProviders, err := getTokenDataProviders(lggr, pluginConfig, sourceChain.LogPoller())
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not get token data providers")
	}

	// Prom wrappers
	onRampReader = observability.NewObservedOnRampReader(onRampReader, int64(chainID), ccip.ExecPluginLabel)
	sourcePriceRegistry = observability.NewPriceRegistryReader(sourcePriceRegistry, int64(chainID), ccip.ExecPluginLabel)
	commitStoreReader = observability.NewObservedCommitStoreReader(commitStoreReader, destChainID, ccip.ExecPluginLabel)
	offRampReader = observability.NewObservedOffRampReader(offRampReader, destChainID, ccip.ExecPluginLabel)

	destChainSelector, err := chainselectors.SelectorFromChainId(uint64(destChainID))
	if err != nil {
		return nil, nil, fmt.Errorf("get chain %d selector: %w", destChainID, err)
	}

	execLggr.Infow("Initialized exec plugin",
		"pluginConfig", pluginConfig,
		"onRampAddress", offRampConfig.OnRamp,
		"sourcePriceRegistry", sourcePriceRegistry.Address(),
		"dynamicOnRampConfig", dynamicOnRampConfig,
		"sourceNative", sourceWrappedNative,
		"sourceRouter", sourceRouter.Address())

	return &ExecutionPluginStaticConfig{
			lggr:                     execLggr,
			onRampReader:             onRampReader,
			commitStoreReader:        commitStoreReader,
			offRampReader:            offRampReader,
			sourcePriceRegistry:      sourcePriceRegistry,
			sourceWrappedNativeToken: sourceWrappedNative,
			destChainSelector:        destChainSelector,
			priceRegistryProvider:    ccipdataprovider.NewEvmPriceRegistry(destChain.LogPoller(), destChain.Client(), execLggr, ccip.ExecPluginLabel),
			tokenDataWorker: tokendata.NewBackgroundWorker(
				ctx,
				tokenDataProviders,
				numTokenDataWorkers,
				5*time.Second,
				offRampReader.OnchainConfig().PermissionLessExecutionThresholdSeconds,
			),
		}, &ccipcommon.BackfillArgs{
			SourceLP:         sourceChain.LogPoller(),
			DestLP:           destChain.LogPoller(),
			SourceStartBlock: pluginConfig.SourceStartBlock,
			DestStartBlock:   pluginConfig.DestStartBlock,
		}, nil
}

func NewExecutionServices(ctx context.Context, lggr logger.Logger, jb job.Job, chainSet legacyevm.LegacyChainContainer, new bool, argsNoPlugin libocr2.OCR2OracleArgs, logError func(string), qopts ...pg.QOpt) ([]job.ServiceCtx, error) {
	execPluginConfig, backfillArgs, err := jobSpecToExecPluginConfig(ctx, lggr, jb, chainSet)
	if err != nil {
		return nil, err
	}
	wrappedPluginFactory := NewExecutionReportingPluginFactory(*execPluginConfig)

	if err1 := execPluginConfig.offRampReader.RegisterFilters(qopts...); err1 != nil {
		return nil, err1
	}
	if err1 := execPluginConfig.onRampReader.RegisterFilters(qopts...); err1 != nil {
		return nil, err1
	}
	if err1 := execPluginConfig.commitStoreReader.RegisterFilters(qopts...); err1 != nil {
		return nil, err1
	}
	for _, dp := range execPluginConfig.tokenDataWorker.GetReaders() {
		if err1 := dp.RegisterFilters(qopts...); err1 != nil {
			return nil, err1
		}
	}
	destChainID, err := chainselectors.ChainIdFromSelector(execPluginConfig.destChainSelector)
	if err != nil {
		return nil, err
	}
	argsNoPlugin.ReportingPluginFactory = promwrapper.NewPromFactory(wrappedPluginFactory, "CCIPExecution", jb.OCR2OracleSpec.Relay, big.NewInt(0).SetUint64(destChainID))
	argsNoPlugin.Logger = commonlogger.NewOCRWrapper(execPluginConfig.lggr, true, logError)
	oracle, err := libocr2.NewOracle(argsNoPlugin)
	if err != nil {
		return nil, err
	}
	// If this is a brand-new job, then we make use of the start blocks. If not then we're rebooting and log poller will pick up where we left off.
	if new {
		return []job.ServiceCtx{
			oraclelib.NewBackfilledOracle(
				execPluginConfig.lggr,
				backfillArgs.SourceLP,
				backfillArgs.DestLP,
				backfillArgs.SourceStartBlock,
				backfillArgs.DestStartBlock,
				job.NewServiceAdapter(oracle)),
		}, nil
	}
	return []job.ServiceCtx{job.NewServiceAdapter(oracle)}, nil
}

func getTokenDataProviders(lggr logger.Logger, pluginConfig ccipconfig.ExecutionPluginJobSpecConfig, sourceLP logpoller.LogPoller) (map[common.Address]tokendata.Reader, error) {
	tokenDataProviders := make(map[common.Address]tokendata.Reader)

	if pluginConfig.USDCConfig.AttestationAPI != "" {
		lggr.Infof("USDC token data provider enabled")
		err := pluginConfig.USDCConfig.ValidateUSDCConfig()
		if err != nil {
			return nil, err
		}

		attestationURI, err := url.ParseRequestURI(pluginConfig.USDCConfig.AttestationAPI)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse USDC attestation API")
		}

		usdcReader, err := ccipdata.NewUSDCReader(lggr, pluginConfig.USDCConfig.SourceMessageTransmitterAddress, sourceLP)
		if err != nil {
			return nil, err
		}
		tokenDataProviders[pluginConfig.USDCConfig.SourceTokenAddress] = usdc.NewUSDCTokenDataReader(
			lggr,
			usdcReader,
			attestationURI,
			pluginConfig.USDCConfig.AttestationAPITimeoutSeconds,
		)

	}

	return tokenDataProviders, nil
}

// UnregisterExecPluginLpFilters unregisters all the registered filters for both source and dest chains.
// See comment in UnregisterCommitPluginLpFilters
// It MUST mirror the filters registered in NewExecutionServices.
func UnregisterExecPluginLpFilters(ctx context.Context, lggr logger.Logger, jb job.Job, chainSet legacyevm.LegacyChainContainer, qopts ...pg.QOpt) error {
	execPluginConfig, _, err := jobSpecToExecPluginConfig(ctx, lggr, jb, chainSet)
	if err != nil {
		return err
	}
	if err := execPluginConfig.onRampReader.Close(qopts...); err != nil {
		return err
	}
	for _, tokenReader := range execPluginConfig.tokenDataWorker.GetReaders() {
		if err := tokenReader.Close(qopts...); err != nil {
			return err
		}
	}
	if err := execPluginConfig.offRampReader.Close(qopts...); err != nil {
		return err
	}
	return execPluginConfig.commitStoreReader.Close(qopts...)
}

// ExecReportToEthTxMeta generates a txmgr.EthTxMeta from the given report.
// Only MessageIDs will be populated in the TxMeta.
func ExecReportToEthTxMeta(typ ccipconfig.ContractType, ver semver.Version) (func(report []byte) (*txmgr.TxMeta, error), error) {
	return factory.ExecReportToEthTxMeta(typ, ver)
}
