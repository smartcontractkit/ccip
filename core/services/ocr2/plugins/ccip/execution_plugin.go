package ccip

import (
	"encoding/json"
	"math/big"
	"time"

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
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/promwrapper"
	"github.com/smartcontractkit/chainlink/v2/core/services/pipeline"
)

const (
	DefaultInflightCacheExpiry   = 3 * time.Minute
	DefaultRootSnoozeTime        = 10 * time.Minute
	EXEC_CCIP_SENDS              = "Exec ccip sends"
	EXEC_REPORT_ACCEPTS          = "Exec report accepts"
	EXEC_EXECUTION_STATE_CHANGES = "Exec execution state changes"
)

func NewExecutionServices(lggr logger.Logger, jb job.Job, chainSet evm.ChainSet, new bool, pr pipeline.Runner, argsNoPlugin libocr2.OracleArgs, logError func(string)) ([]job.ServiceCtx, error) {
	spec := jb.OCR2OracleSpec
	var pluginConfig ccipconfig.ExecutionPluginConfig
	err := json.Unmarshal(spec.PluginConfig.Bytes(), &pluginConfig)
	if err != nil {
		return nil, err
	}
	err = pluginConfig.ValidateExecutionPluginConfig()
	if err != nil {
		return nil, err
	}
	lggr.Infof("CCIP execution plugin initialized with offchainConfig: %+v", pluginConfig)

	sourceChain, err := chainSet.Get(big.NewInt(0).SetUint64(pluginConfig.SourceChainID))
	if err != nil {
		return nil, errors.Wrap(err, "unable to open source chain")
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

	commitStore, err := LoadCommitStore(common.HexToAddress(pluginConfig.CommitStoreID), destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed loading commitStore")
	}

	onRamp, err := LoadOnRamp(common.HexToAddress(pluginConfig.OnRampID), sourceChain.Client())
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

	offRamp, err := LoadOffRamp(common.HexToAddress(spec.ContractID), destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed loading offRamp")
	}
	dynamicOffRampConfig, err := offRamp.GetDynamicConfig(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "failed loading offRamp config")
	}
	destRouter, err := router.NewRouter(dynamicOffRampConfig.Router, destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed loading dest router")
	}
	destWrappedNative, err := destRouter.GetWrappedNative(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "could not get destination token")
	}

	dynamicConfig, err := commitStore.GetDynamicConfig(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "failed getting the dynamic config from the commitStore")
	}
	destPriceRegistry, err := price_registry.NewPriceRegistry(dynamicConfig.PriceRegistry, destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "could not create dest price registry")
	}

	srcPriceRegistry, err := price_registry.NewPriceRegistry(dynamicOnRampConfig.PriceRegistry, sourceChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "could not create source price registry")
	}

	lggr = lggr.With("srcChain", ChainName(int64(pluginConfig.SourceChainID)), "dstChain", ChainName(destChainID))

	rootSnoozeTime := DefaultRootSnoozeTime
	if pluginConfig.RootSnoozeTime.Duration() != 0 {
		rootSnoozeTime = pluginConfig.RootSnoozeTime.Duration()
	}
	inflightCacheExpiry := DefaultInflightCacheExpiry
	if pluginConfig.InflightCacheExpiry.Duration() != 0 {
		inflightCacheExpiry = pluginConfig.InflightCacheExpiry.Duration()
	}

	eventSignatures := GetEventSignatures()
	wrappedPluginFactory := NewExecutionReportingPluginFactory(
		ExecutionPluginConfig{
			lggr:                   lggr,
			source:                 sourceChain.LogPoller(),
			dest:                   destChain.LogPoller(),
			offRamp:                offRamp,
			onRamp:                 onRamp,
			commitStore:            commitStore,
			eventSignatures:        eventSignatures,
			leafHasher:             NewLeafHasher(pluginConfig.SourceChainID, uint64(destChainID), onRamp.Address(), hasher.NewKeccakCtx()),
			snoozeTime:             rootSnoozeTime,
			inflightCacheExpiry:    inflightCacheExpiry,
			destPriceRegistry:      destPriceRegistry,
			srcPriceRegistry:       srcPriceRegistry,
			destGasEstimator:       destChain.GasEstimator(),
			destWrappedNativeToken: destWrappedNative,
			srcWrappedNativeToken:  sourceWrappedNative,
		})

	// Subscribe to all relevant commit logs.
	err = sourceChain.LogPoller().RegisterFilter(logpoller.Filter{Name: logpoller.FilterName(EXEC_CCIP_SENDS, onRamp.Address().String()),
		EventSigs: []common.Hash{eventSignatures.SendRequested}, Addresses: []common.Address{onRamp.Address()}})
	if err != nil {
		return nil, err
	}
	err = destChain.LogPoller().RegisterFilter(logpoller.Filter{Name: logpoller.FilterName(EXEC_REPORT_ACCEPTS, commitStore.Address().String()),
		EventSigs: []common.Hash{ReportAccepted}, Addresses: []common.Address{commitStore.Address()}})
	if err != nil {
		return nil, err
	}
	err = destChain.LogPoller().RegisterFilter(logpoller.Filter{Name: logpoller.FilterName(EXEC_EXECUTION_STATE_CHANGES, offRamp.Address().String()),
		EventSigs: []common.Hash{eventSignatures.ExecutionStateChanged}, Addresses: []common.Address{offRamp.Address()}})
	if err != nil {
		return nil, err
	}

	argsNoPlugin.ReportingPluginFactory = promwrapper.NewPromFactory(wrappedPluginFactory, "CCIPExecution", string(spec.Relay), destChain.ID())
	argsNoPlugin.Logger = logger.NewOCRWrapper(lggr.Named("CCIPExecution").With(
		"srcChain", ChainName(int64(pluginConfig.SourceChainID)), "dstChain", ChainName(destChainID)), true, logError)
	oracle, err := libocr2.NewOracle(argsNoPlugin)
	if err != nil {
		return nil, err
	}
	// If this is a brand-new job, then we make use of the start blocks. If not then we're rebooting and log poller will pick up where we left off.
	if new {
		return []job.ServiceCtx{NewBackfilledOracle(
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
