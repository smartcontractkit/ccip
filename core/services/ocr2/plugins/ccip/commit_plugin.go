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
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/promwrapper"
	"github.com/smartcontractkit/chainlink/v2/core/services/pipeline"
)

const (
	COMMIT_PRICE_UPDATES = "Commit price updates"
	COMMIT_CCIP_SENDS    = "Commit ccip sends"
)

func NewCommitServices(lggr logger.Logger, jb job.Job, chainSet evm.ChainSet, new bool, pr pipeline.Runner, argsNoPlugin libocr2.OracleArgs, logError func(string)) ([]job.ServiceCtx, error) {
	spec := jb.OCR2OracleSpec

	var pluginConfig ccipconfig.CommitPluginConfig
	err := json.Unmarshal(spec.PluginConfig.Bytes(), &pluginConfig)
	if err != nil {
		return nil, err
	}
	lggr.Infof("CCIP commit plugin initialized with offchainConfig: %+v", pluginConfig)

	chainIDInterface, ok := spec.RelayConfig["chainID"]
	if !ok {
		return nil, errors.New("chainID must be provided in relay config")
	}
	destChainID := int64(chainIDInterface.(float64))
	destChain, err := chainSet.Get(big.NewInt(destChainID))
	if err != nil {
		return nil, errors.Wrap(err, "get chainset")
	}
	commitStore, err := LoadCommitStore(common.HexToAddress(spec.ContractID), destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed loading commitStore")
	}
	staticConfig, err := commitStore.GetStaticConfig(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "failed getting the static config from the commitStore")
	}
	sourceChain, err := chainSet.Get(big.NewInt(0).SetUint64(staticConfig.SourceChainId))
	if err != nil {
		return nil, errors.Wrap(err, "unable to open source chain")
	}
	onRamp, err := LoadOnRamp(staticConfig.OnRamp, sourceChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed loading onRamp")
	}
	if staticConfig.ChainId != uint64(destChainID) {
		return nil, errors.Errorf("Wrong dest chain ID got %d expected from jobspec %d", staticConfig.ChainId, destChainID)
	}
	// TODO DynamicConfig call, remove
	dynamicConfig, err := commitStore.GetDynamicConfig(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "failed getting the dynamic config from the commitStore")
	}

	priceGetterObject, err := NewPriceGetter(pluginConfig.TokenPricesUSDPipeline, pr, jb.ID, jb.ExternalJobID, jb.Name.ValueOrZero(), lggr)
	if err != nil {
		return nil, err
	}

	// subscribe for GasFeeUpdated logs, but the PriceRegistry is only available as part of onchain commitStore's config
	// TODO: how to detect if commitStoreConfig.PriceRegistry changes on-chain? Currently, we expect a plugin/job/node restart
	priceRegistry, err := price_registry.NewPriceRegistry(dynamicConfig.PriceRegistry, destChain.Client())
	if err != nil {
		return nil, err
	}
	dynamicOnRampConfig, err := onRamp.GetDynamicConfig(nil)
	if err != nil {
		return nil, err
	}
	router, err := router.NewRouter(dynamicOnRampConfig.Router, sourceChain.Client())
	if err != nil {
		return nil, err
	}
	sourceNative, err := router.GetWrappedNative(nil)
	if err != nil {
		return nil, err
	}

	eventSigs := GetEventSignatures()
	err = destChain.LogPoller().RegisterFilter(logpoller.Filter{Name: logpoller.FilterName(COMMIT_PRICE_UPDATES, dynamicConfig.PriceRegistry.String()),
		EventSigs: []common.Hash{UsdPerUnitGasUpdated, UsdPerTokenUpdated}, Addresses: []common.Address{dynamicConfig.PriceRegistry}})
	if err != nil {
		return nil, err
	}
	err = sourceChain.LogPoller().RegisterFilter(logpoller.Filter{Name: logpoller.FilterName(COMMIT_CCIP_SENDS, onRamp.Address().String()), EventSigs: []common.Hash{eventSigs.SendRequested}, Addresses: []common.Address{onRamp.Address()}})
	if err != nil {
		return nil, err
	}

	leafHasher := NewLeafHasher(staticConfig.SourceChainId, uint64(destChainID), onRamp.Address(), hasher.NewKeccakCtx())
	wrappedPluginFactory := NewCommitReportingPluginFactory(
		CommitPluginConfig{
			lggr:               lggr,
			source:             sourceChain.LogPoller(),
			dest:               destChain.LogPoller(),
			reqEventSig:        eventSigs,
			onRamp:             onRamp,
			priceRegistry:      priceRegistry,
			priceGetter:        priceGetterObject,
			sourceNative:       sourceNative,
			sourceFeeEstimator: sourceChain.GasEstimator(),
			sourceChainID:      staticConfig.SourceChainId,
			commitStore:        commitStore,
			hasher:             leafHasher,
		})
	argsNoPlugin.ReportingPluginFactory = promwrapper.NewPromFactory(wrappedPluginFactory, "CCIPCommit", string(spec.Relay), destChain.ID())
	argsNoPlugin.Logger = logger.NewOCRWrapper(lggr.Named("CCIPCommit").With(
		"srcChain", ChainName(int64(staticConfig.SourceChainId)), "dstChain", ChainName(destChainID)), true, logError)
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
