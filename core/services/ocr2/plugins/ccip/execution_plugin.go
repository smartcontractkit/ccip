package ccip

import (
	"encoding/json"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	libocr2 "github.com/smartcontractkit/libocr/offchainreporting2"

	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/fee_manager"
	"github.com/smartcontractkit/chainlink/core/logger"
	hlp "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/core/services/job"
	ccipconfig "github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/promwrapper"
	"github.com/smartcontractkit/chainlink/core/services/pipeline"
)

const (
	DefaultInflightCacheExpiry = 3 * time.Minute
	DefaultRootSnoozeTime      = 10 * time.Minute
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
	offRamp, err := LoadOffRamp(common.HexToAddress(spec.ContractID), destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed loading offRamp")
	}

	lggr = lggr.With("srcChain", hlp.ChainName(int64(pluginConfig.SourceChainID)), "dstChain", hlp.ChainName(destChainID))

	rootSnoozeTime := DefaultRootSnoozeTime
	if pluginConfig.RootSnoozeTime.Duration() != 0 {
		rootSnoozeTime = pluginConfig.RootSnoozeTime.Duration()
	}
	inflightCacheExpiry := DefaultInflightCacheExpiry
	if pluginConfig.InflightCacheExpiry.Duration() != 0 {
		inflightCacheExpiry = pluginConfig.InflightCacheExpiry.Duration()
	}
	priceGetterObject, err := NewPriceGetter(pluginConfig.TokensPerFeeCoinPipeline, pr, jb.ID, jb.ExternalJobID, jb.Name.ValueOrZero(), lggr)
	if err != nil {
		return nil, err
	}

	// subscribe for GasFeeUpdated logs, but FeeManager is only available as part of onchain offramp's config
	// TODO: how to detect if OffRampConfig.FeeManager changes on-chain? Currently, we expect a plugin/job/node restart
	offRampConfig, err := offRamp.GetOffRampConfig(nil)
	if err != nil {
		return nil, err
	}
	feeManager, err := fee_manager.NewFeeManager(offRampConfig.FeeManager, destChain.Client())
	if err != nil {
		return nil, err
	}

	eventSignatures := GetEventSignatures()
	wrappedPluginFactory := NewExecutionReportingPluginFactory(
		ExecutionPluginConfig{
			lggr:                lggr,
			source:              sourceChain.LogPoller(),
			dest:                destChain.LogPoller(),
			offRamp:             offRamp,
			onRamp:              onRamp,
			commitStore:         commitStore,
			feeManager:          feeManager,
			builder:             NewBatchBuilder(lggr, eventSignatures, offRamp),
			eventSignatures:     eventSignatures,
			priceGetter:         priceGetterObject,
			leafHasher:          NewLeafHasher(pluginConfig.SourceChainID, uint64(destChainID), onRamp.Address(), hasher.NewKeccakCtx()),
			snoozeTime:          rootSnoozeTime,
			inflightCacheExpiry: inflightCacheExpiry,
			sourceChainID:       pluginConfig.SourceChainID,
			gasLimit:            BatchGasLimit,
			destGasEstimator:    destChain.TxManager().GetGasEstimator(),
			sourceGasEstimator:  sourceChain.TxManager().GetGasEstimator(),
		})
	// Subscribe to all relevant commit logs.
	_, err = sourceChain.LogPoller().RegisterFilter(logpoller.Filter{EventSigs: []common.Hash{eventSignatures.SendRequested}, Addresses: []common.Address{onRamp.Address()}})
	if err != nil {
		return nil, err
	}
	_, err = destChain.LogPoller().RegisterFilter(logpoller.Filter{EventSigs: []common.Hash{ReportAccepted}, Addresses: []common.Address{commitStore.Address()}})
	if err != nil {
		return nil, err
	}
	_, err = destChain.LogPoller().RegisterFilter(logpoller.Filter{EventSigs: []common.Hash{GasFeeUpdated}, Addresses: []common.Address{offRampConfig.FeeManager}})
	if err != nil {
		return nil, err
	}
	_, err = destChain.LogPoller().RegisterFilter(logpoller.Filter{EventSigs: []common.Hash{eventSignatures.ExecutionStateChanged}, Addresses: []common.Address{offRamp.Address()}})
	if err != nil {
		return nil, err
	}

	argsNoPlugin.ReportingPluginFactory = promwrapper.NewPromFactory(wrappedPluginFactory, "CCIPExecution", string(spec.Relay), destChain.ID())
	argsNoPlugin.Logger = logger.NewOCRWrapper(lggr.Named("CCIPExecution").With(
		"srcChain", hlp.ChainName(int64(pluginConfig.SourceChainID)), "dstChain", hlp.ChainName(destChainID)), true, logError)
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
