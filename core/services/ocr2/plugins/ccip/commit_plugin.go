package ccip

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	pluginConfig, err := ParseAndVerifyPluginConfig(spec.PluginConfig)
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

	sourceChain, err := chainSet.Get(big.NewInt(0).SetUint64(pluginConfig.SourceChainID))
	if err != nil {
		return nil, errors.Wrap(err, "unable to open source chain")
	}

	inflightCacheExpiry := DefaultInflightCacheExpiry
	if pluginConfig.InflightCacheExpiry.Duration() != 0 {
		inflightCacheExpiry = pluginConfig.InflightCacheExpiry.Duration()
	}

	commitStore, err := LoadCommitStore(common.HexToAddress(spec.ContractID), destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed loading commitStore")
	}
	onRamp, err := LoadOnRamp(common.HexToAddress(pluginConfig.OnRampID), sourceChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed loading onRamp")
	}
	offRamp, err := LoadOffRamp(common.HexToAddress(pluginConfig.OffRampID), destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed loading offRamp")
	}

	staticConfig, err := commitStore.GetStaticConfig(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "failed getting the static config from the commitStore")
	}
	if staticConfig.OnRamp != onRamp.Address() {
		return nil, errors.Errorf("Wrong onRamp got %s expected from jobspec %s", staticConfig.OnRamp, onRamp.Address())
	}
	if staticConfig.SourceChainId != pluginConfig.SourceChainID {
		return nil, errors.Errorf("Wrong source chain ID got %d expected from jobspec %d", staticConfig.SourceChainId, pluginConfig.SourceChainID)
	}
	if staticConfig.ChainId != uint64(destChainID) {
		return nil, errors.Errorf("Wrong dest chain ID got %d expected from jobspec %d", staticConfig.ChainId, destChainID)
	}
	dynamicConfig, err := commitStore.GetDynamicConfig(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "failed getting the dynamic config from the commitStore")
	}

	seqParsers := func(log logpoller.Log) (uint64, error) {
		req, err2 := onRamp.ParseCCIPSendRequested(types.Log{Data: log.Data, Topics: log.GetTopics()})
		if err2 != nil {
			lggr.Warnf("failed to parse log: %+v", log)
			return 0, err2
		}
		return req.Message.SequenceNumber, nil
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

	leafHasher := NewLeafHasher(pluginConfig.SourceChainID, uint64(destChainID), onRamp.Address(), hasher.NewKeccakCtx())
	wrappedPluginFactory := NewCommitReportingPluginFactory(
		CommitPluginConfig{
			lggr:                lggr,
			source:              sourceChain.LogPoller(),
			dest:                destChain.LogPoller(),
			seqParsers:          seqParsers,
			reqEventSig:         eventSigs,
			onRamp:              onRamp.Address(),
			offRamp:             offRamp,
			priceRegistry:       priceRegistry,
			priceGetter:         priceGetterObject,
			sourceNative:        sourceNative,
			sourceFeeEstimator:  sourceChain.TxManager().GetGasEstimator(),
			sourceChainID:       pluginConfig.SourceChainID,
			commitStore:         commitStore,
			hasher:              leafHasher,
			inflightCacheExpiry: inflightCacheExpiry,
		})
	argsNoPlugin.ReportingPluginFactory = promwrapper.NewPromFactory(wrappedPluginFactory, "CCIPCommit", string(spec.Relay), destChain.ID())
	argsNoPlugin.Logger = logger.NewOCRWrapper(lggr.Named("CCIPCommit").With(
		"srcChain", ChainName(int64(pluginConfig.SourceChainID)), "dstChain", ChainName(destChainID)), true, logError)
	oracle, err := libocr2.NewOracle(argsNoPlugin)
	if err != nil {
		return nil, err
	}
	// If this is a brand-new job, then we make use of the start blocks. If not then we're rebooting and log poller will pick up where we left off.
	if new {
		return []job.ServiceCtx{&BackfilledOracle{
			srcStartBlock: pluginConfig.SourceStartBlock,
			dstStartBlock: pluginConfig.DestStartBlock,
			src:           sourceChain.LogPoller(),
			dst:           destChain.LogPoller(),
			oracle:        job.NewServiceAdapter(oracle),
			lggr:          lggr,
		}}, nil
	}
	return []job.ServiceCtx{job.NewServiceAdapter(oracle)}, nil
}

func ParseAndVerifyPluginConfig(jsonConfig job.JSONConfig) (ccipconfig.CommitPluginConfig, error) {
	var pluginConfig ccipconfig.CommitPluginConfig
	err := json.Unmarshal(jsonConfig.Bytes(), &pluginConfig)
	if err != nil {
		return ccipconfig.CommitPluginConfig{}, err
	}
	err = pluginConfig.ValidateCommitPluginConfig()
	if err != nil {
		return ccipconfig.CommitPluginConfig{}, err
	}
	return pluginConfig, nil
}
