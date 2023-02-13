package ccip

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	libocr2 "github.com/smartcontractkit/libocr/offchainreporting2"

	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/core/logger"
	hlp "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/core/services/job"
	ccipconfig "github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/promwrapper"
)

func NewCommitServices(lggr logger.Logger, spec *job.OCR2OracleSpec, chainSet evm.ChainSet, new bool, argsNoPlugin libocr2.OracleArgs, logError func(string)) ([]job.ServiceCtx, error) {
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

	lggr = lggr.With("srcChain", hlp.ChainName(int64(pluginConfig.SourceChainID)), "dstChain", hlp.ChainName(destChainID))

	inflightCacheExpiry := DefaultInflightCacheExpiry
	if pluginConfig.InflightCacheExpiry.Duration() != 0 {
		inflightCacheExpiry = pluginConfig.InflightCacheExpiry.Duration()
	}

	commitStoreAddress := common.HexToAddress(spec.ContractID)
	onRampAddress := common.HexToAddress(pluginConfig.OnRampID)

	err = ccipconfig.VerifyTypeAndVersion(onRampAddress, sourceChain.Client(), ccipconfig.EVM2EVMOnRamp)
	if err != nil {
		return nil, err
	}
	err = ccipconfig.VerifyTypeAndVersion(commitStoreAddress, destChain.Client(), ccipconfig.CommitStore)
	if err != nil {
		return nil, err
	}

	commitStore, err := commit_store.NewCommitStore(commitStoreAddress, destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed loading the commitStore")
	}
	commitStoreConfig, err := commitStore.GetConfig(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "failed getting the config from the commitStore")
	}
	if commitStoreConfig.OnRamp != onRampAddress {
		return nil, errors.Errorf("Wrong onRamp got %s expected from jobspec %s", commitStoreConfig.OnRamp, onRampAddress)
	}
	if commitStoreConfig.SourceChainId != pluginConfig.SourceChainID {
		return nil, errors.Errorf("Wrong source chain ID got %d expected from jobspec %d", commitStoreConfig.SourceChainId, pluginConfig.SourceChainID)
	}
	if commitStoreConfig.ChainId != uint64(destChainID) {
		return nil, errors.Errorf("Wrong dest chain ID got %d expected from jobspec %d", commitStoreConfig.ChainId, destChainID)
	}

	onRamp, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(onRampAddress, sourceChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed loading the onRamp")
	}

	seqParsers := func(log logpoller.Log) (uint64, error) {
		req, err2 := onRamp.ParseCCIPSendRequested(types.Log{Data: log.Data, Topics: log.GetTopics()})
		if err2 != nil {
			lggr.Warnf("failed to parse log: %+v", log)
			return 0, err2
		}
		return req.Message.SequenceNumber, nil
	}

	eventSigs := GetEventSignatures()
	_, err = sourceChain.LogPoller().RegisterFilter(logpoller.Filter{EventSigs: []common.Hash{eventSigs.SendRequested}, Addresses: []common.Address{onRampAddress}})
	if err != nil {
		return nil, err
	}

	leafHasher := NewLeafHasher(pluginConfig.SourceChainID, uint64(destChainID), onRampAddress, hasher.NewKeccakCtx())
	wrappedPluginFactory := NewCommitReportingPluginFactory(lggr, sourceChain.LogPoller(), commitStore, seqParsers, eventSigs, onRampAddress, leafHasher, inflightCacheExpiry)
	argsNoPlugin.ReportingPluginFactory = promwrapper.NewPromFactory(wrappedPluginFactory, "CCIPCommit", string(spec.Relay), destChain.ID())
	argsNoPlugin.Logger = logger.NewOCRWrapper(lggr.Named("CCIPCommit").With(
		"srcChain", hlp.ChainName(int64(pluginConfig.SourceChainID)), "dstChain", hlp.ChainName(destChainID)), true, logError)
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
