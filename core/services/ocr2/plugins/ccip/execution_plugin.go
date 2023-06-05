package ccip

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	libocr2 "github.com/smartcontractkit/libocr/offchainreporting2"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/promwrapper"
)

const (
	EXEC_CCIP_SENDS              = "Exec ccip sends"
	EXEC_REPORT_ACCEPTS          = "Exec report accepts"
	EXEC_EXECUTION_STATE_CHANGES = "Exec execution state changes"
	EXEC_TOKEN_POOL_ADDED        = "Token pool added"
	EXEC_TOKEN_POOL_REMOVED      = "Token pool removed"
)

func NewExecutionServices(lggr logger.Logger, jb job.Job, chainSet evm.ChainSet, new bool, argsNoPlugin libocr2.OracleArgs, logError func(string)) ([]job.ServiceCtx, error) {
	spec := jb.OCR2OracleSpec
	var pluginConfig ccipconfig.ExecutionPluginConfig
	err := json.Unmarshal(spec.PluginConfig.Bytes(), &pluginConfig)
	if err != nil {
		return nil, err
	}

	chainIDInterface, ok := spec.RelayConfig["chainID"]
	if !ok {
		return nil, errors.New("chainID must be provided in relay config")
	}
	destChainID := int64(chainIDInterface.(float64))
	destChain, err := chainSet.Get(big.NewInt(destChainID))
	if err != nil {
		return nil, errors.Wrap(err, "get chainset")
	}
	offRamp, err := LoadOffRamp(common.HexToAddress(spec.ContractID), destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed loading offRamp")
	}
	offRampConfig, err := offRamp.GetStaticConfig(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	sourceChain, err := chainSet.Get(big.NewInt(0).SetUint64(uint64(pluginConfig.SourceEvmChainId)))
	if err != nil {
		return nil, errors.Wrap(err, "unable to open source chain")
	}
	commitStore, err := LoadCommitStore(offRampConfig.CommitStore, destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed loading commitStore")
	}
	onRamp, err := LoadOnRamp(offRampConfig.OnRamp, sourceChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed loading onRamp")
	}
	dynamicOnRampConfig, err := onRamp.GetDynamicConfig(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "failed loading onRamp config")
	}
	srcRouter, err := router.NewRouter(dynamicOnRampConfig.Router, sourceChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed loading source router")
	}
	sourceWrappedNative, err := srcRouter.GetWrappedNative(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "could not get source native token")
	}
	srcPriceRegistry, err := price_registry.NewPriceRegistry(dynamicOnRampConfig.PriceRegistry, sourceChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "could not create source price registry")
	}

	lggr.Named("CCIPExecution").With(
		"srcChain", ChainName(pluginConfig.SourceEvmChainId), "dstChain", ChainName(destChainID))

	wrappedPluginFactory := NewExecutionReportingPluginFactory(
		ExecutionPluginConfig{
			lggr:                  lggr,
			sourceLP:              sourceChain.LogPoller(),
			destLP:                destChain.LogPoller(),
			onRamp:                onRamp,
			offRamp:               offRamp,
			commitStore:           commitStore,
			srcPriceRegistry:      srcPriceRegistry,
			srcWrappedNativeToken: sourceWrappedNative,
			destClient:            destChain.Client(),
			destGasEstimator:      destChain.GasEstimator(),
			leafHasher:            hasher.NewLeafHasher(offRampConfig.SourceChainSelector, offRampConfig.ChainSelector, onRamp.Address(), hasher.NewKeccakCtx()),
		})

	// Subscribe to all relevant logs.
	err = sourceChain.LogPoller().RegisterFilter(logpoller.Filter{
		Name:      logpoller.FilterName(EXEC_CCIP_SENDS, onRamp.Address().String()),
		EventSigs: []common.Hash{abihelpers.EventSignatures.SendRequested},
		Addresses: []common.Address{onRamp.Address()},
	})
	if err != nil {
		return nil, err
	}
	err = destChain.LogPoller().RegisterFilter(logpoller.Filter{
		Name:      logpoller.FilterName(EXEC_REPORT_ACCEPTS, commitStore.Address().String()),
		EventSigs: []common.Hash{abihelpers.EventSignatures.ReportAccepted},
		Addresses: []common.Address{commitStore.Address()},
	})
	if err != nil {
		return nil, err
	}
	err = destChain.LogPoller().RegisterFilter(logpoller.Filter{
		Name:      logpoller.FilterName(EXEC_EXECUTION_STATE_CHANGES, offRamp.Address().String()),
		EventSigs: []common.Hash{abihelpers.EventSignatures.ExecutionStateChanged},
		Addresses: []common.Address{offRamp.Address()},
	})
	if err != nil {
		return nil, err
	}
	err = destChain.LogPoller().RegisterFilter(logpoller.Filter{
		Name:      logpoller.FilterName(EXEC_TOKEN_POOL_ADDED, offRamp.Address().String()),
		EventSigs: []common.Hash{abihelpers.EventSignatures.PoolAdded},
		Addresses: []common.Address{offRamp.Address()},
	})
	if err != nil {
		return nil, err
	}
	err = destChain.LogPoller().RegisterFilter(logpoller.Filter{
		Name:      logpoller.FilterName(EXEC_TOKEN_POOL_REMOVED, offRamp.Address().String()),
		EventSigs: []common.Hash{abihelpers.EventSignatures.PoolRemoved},
		Addresses: []common.Address{offRamp.Address()},
	})
	if err != nil {
		return nil, err
	}

	argsNoPlugin.ReportingPluginFactory = promwrapper.NewPromFactory(wrappedPluginFactory, "CCIPExecution", string(spec.Relay), destChain.ID())
	argsNoPlugin.Logger = logger.NewOCRWrapper(lggr, true, logError)
	oracle, err := libocr2.NewOracle(argsNoPlugin)
	if err != nil {
		return nil, err
	}
	lggr.Infof("Initialized exec plugin",
		"pluginConfig", pluginConfig,
		"onRampAddress", onRamp.Address(),
		"sourcePriceRegistry", srcPriceRegistry.Address(),
		"dynamicOnRampConfig", dynamicOnRampConfig,
		"sourceNative", sourceWrappedNative,
		"sourceRouter", srcRouter.Address())
	// If this is a brand-new job, then we make use of the start blocks. If not then we're rebooting and log poller will pick up where we left off.
	if new {
		return []job.ServiceCtx{
			NewBackfilledOracle(
				lggr,
				sourceChain.LogPoller(),
				destChain.LogPoller(),
				pluginConfig.SourceStartBlock,
				pluginConfig.DestStartBlock,
				job.NewServiceAdapter(oracle)),
		}, nil
	}
	return []job.ServiceCtx{job.NewServiceAdapter(oracle)}, nil
}
