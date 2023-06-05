package ccip

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	libocr2 "github.com/smartcontractkit/libocr/offchainreporting2"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
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
	chainIDInterface, ok := spec.RelayConfig["chainID"]
	if !ok {
		return nil, errors.New("chainID must be provided in relay config")
	}
	destChainID := int64(chainIDInterface.(float64))
	destChain, err := chainSet.Get(big.NewInt(destChainID))
	if err != nil {
		return nil, errors.Wrap(err, "get chainset")
	}
	commitStore, err := LoadCommitStore(common.HexToAddress(spec.ContractID), CommitPluginLabel, destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed loading commitStore")
	}
	staticConfig, err := commitStore.GetStaticConfig(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "failed getting the static config from the commitStore")
	}
	sourceChain, err := chainSet.Get(big.NewInt(0).SetUint64(uint64(pluginConfig.SourceEvmChainId)))
	if err != nil {
		return nil, errors.Wrap(err, "unable to open source chain")
	}
	offRamp, err := LoadOffRamp(common.HexToAddress(pluginConfig.OffRamp), CommitPluginLabel, destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed loading offRamp")
	}
	onRamp, err := LoadOnRamp(staticConfig.OnRamp, CommitPluginLabel, sourceChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed loading onRamp")
	}
	priceGetterObject, err := NewPriceGetter(pluginConfig.TokenPricesUSDPipeline, pr, jb.ID, jb.ExternalJobID, jb.Name.ValueOrZero(), lggr)
	if err != nil {
		return nil, err
	}
	dynamicOnRampConfig, err := onRamp.GetDynamicConfig(nil)
	if err != nil {
		return nil, err
	}
	sourceRouter, err := router.NewRouter(dynamicOnRampConfig.Router, sourceChain.Client())
	if err != nil {
		return nil, err
	}
	sourceNative, err := sourceRouter.GetWrappedNative(nil)
	if err != nil {
		return nil, err
	}

	leafHasher := hasher.NewLeafHasher(staticConfig.SourceChainSelector, staticConfig.ChainSelector, onRamp.Address(), hasher.NewKeccakCtx())
	// Note that lggr already has the jobName and contractID (commit store)
	commitLggr := lggr.Named("CCIPCommit").With(
		"sourceChain", ChainName(pluginConfig.SourceEvmChainId),
		"destChain", ChainName(destChainID))
	wrappedPluginFactory := NewCommitReportingPluginFactory(
		CommitPluginConfig{
			lggr:                commitLggr,
			sourceLP:            sourceChain.LogPoller(),
			destLP:              destChain.LogPoller(),
			offRamp:             offRamp,
			onRampAddress:       onRamp.Address(),
			priceGetter:         priceGetterObject,
			sourceNative:        sourceNative,
			sourceFeeEstimator:  sourceChain.GasEstimator(),
			sourceChainSelector: staticConfig.SourceChainSelector,
			destClient:          destChain.Client(),
			commitStore:         commitStore,
			leafHasher:          leafHasher,
			getSeqNumFromLog:    getSeqNumFromLog(onRamp),
		})
	argsNoPlugin.ReportingPluginFactory = promwrapper.NewPromFactory(wrappedPluginFactory, "CCIPCommit", string(spec.Relay), destChain.ID())
	argsNoPlugin.Logger = logger.NewOCRWrapper(commitLggr, true, logError)
	oracle, err := libocr2.NewOracle(argsNoPlugin)
	if err != nil {
		return nil, err
	}
	commitLggr.Infow("NewCommitServices",
		"pluginConfig", pluginConfig,
		"staticConfig", staticConfig,
		"dynamicOnRampConfig", dynamicOnRampConfig,
		"sourceNative", sourceNative,
		"sourceRouter", sourceRouter.Address())
	// If this is a brand-new job, then we make use of the start blocks. If not then we're rebooting and log poller will pick up where we left off.
	if new {
		return []job.ServiceCtx{NewBackfilledOracle(
			commitLggr,
			sourceChain.LogPoller(),
			destChain.LogPoller(),
			pluginConfig.SourceStartBlock,
			pluginConfig.DestStartBlock,
			job.NewServiceAdapter(oracle)),
		}, nil
	}
	return []job.ServiceCtx{job.NewServiceAdapter(oracle)}, nil
}

func getSeqNumFromLog(onRamp evm_2_evm_onramp.EVM2EVMOnRampInterface) func(log logpoller.Log) (uint64, error) {
	return func(log logpoller.Log) (uint64, error) {
		req, err := onRamp.ParseCCIPSendRequested(log.GetGethLog())
		if err != nil {
			return 0, err
		}
		return req.Message.SequenceNumber, nil
	}
}

// CommitReportToEthTxMeta generates a txmgr.EthTxMeta from the given commit report.
// sequence numbers of the committed messages will be added to tx metadata
func CommitReportToEthTxMeta(report []byte) (*txmgr.EthTxMeta, error) {
	commitReport, err := abihelpers.DecodeCommitReport(report)
	if err != nil {
		return nil, err
	}
	n := int(commitReport.Interval.Max-commitReport.Interval.Min) + 1
	seqRange := make([]uint64, n)
	for i := 0; i < n; i++ {
		seqRange[i] = uint64(i) + commitReport.Interval.Min
	}
	return &txmgr.EthTxMeta{
		SeqNumbers: seqRange,
	}, nil
}

func getCommitPluginSourceLpFilters(onRamp common.Address) []logpoller.Filter {
	return []logpoller.Filter{
		{
			Name:      logpoller.FilterName(COMMIT_CCIP_SENDS, onRamp.String()),
			EventSigs: []common.Hash{abihelpers.EventSignatures.SendRequested},
			Addresses: []common.Address{onRamp},
		},
	}
}

func getCommitPluginDestLpFilters(priceRegistry common.Address) []logpoller.Filter {
	return []logpoller.Filter{
		{
			Name:      logpoller.FilterName(COMMIT_PRICE_UPDATES, priceRegistry.String()),
			EventSigs: []common.Hash{abihelpers.EventSignatures.UsdPerUnitGasUpdated, abihelpers.EventSignatures.UsdPerTokenUpdated},
			Addresses: []common.Address{priceRegistry},
		},
	}
}

// GetCommitPluginFilterNamesFromSpec returns all the registered filter names, for both source and dest log pollers.
func GetCommitPluginFilterNamesFromSpec(ctx context.Context, spec *job.OCR2OracleSpec, chainSet evm.ChainSet) ([]string, error) {
	if spec == nil {
		return nil, errors.New("spec is nil")
	}
	if !common.IsHexAddress(spec.ContractID) {
		return nil, fmt.Errorf("invalid contract id address: %s", spec.ContractID)
	}

	destChainIDInterface, ok := spec.RelayConfig["chainID"]
	if !ok {
		return nil, errors.New("chainID must be provided in relay config")
	}
	destChainIDf64, is := destChainIDInterface.(float64)
	if !is {
		return nil, fmt.Errorf("chain id '%v' is not float64", destChainIDInterface)
	}
	destChainID := int64(destChainIDf64)
	destChain, err := chainSet.Get(big.NewInt(destChainID))
	if err != nil {
		return nil, err
	}

	commitStore, err := LoadCommitStore(common.HexToAddress(spec.ContractID), CommitPluginLabel, destChain.Client())
	if err != nil {
		return nil, err
	}

	return getCommitPluginFilterNames(ctx, commitStore)
}

func getCommitPluginFilterNames(ctx context.Context, dstCommitStore commit_store.CommitStoreInterface) ([]string, error) {
	staticCfg, err := dstCommitStore.GetStaticConfig(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	dynamicCfg, err := dstCommitStore.GetDynamicConfig(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	return append(
		getLpFilterNames(getCommitPluginSourceLpFilters(staticCfg.OnRamp)),
		getLpFilterNames(getCommitPluginDestLpFilters(dynamicCfg.PriceRegistry))...,
	), nil
}
