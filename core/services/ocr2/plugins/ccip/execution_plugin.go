package ccip

import (
	"encoding/json"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	libocr2 "github.com/smartcontractkit/libocr/offchainreporting2"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_onramp"
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

func NewExecutionServices(lggr logger.Logger, jb job.Job, chainSet evm.ChainSet, new bool, pr pipeline.Runner, argsNoPlugin libocr2.OracleArgs) ([]job.ServiceCtx, error) {
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
	destChain, err := chainSet.Get(big.NewInt(0).SetUint64(pluginConfig.DestChainID))
	if err != nil {
		return nil, errors.Wrap(err, "unable to open destination chain")
	}

	lggr = lggr.With("srcChain", hlp.ChainName(int64(pluginConfig.SourceChainID)), "dstChain", hlp.ChainName(int64(pluginConfig.DestChainID)))

	rootSnoozeTime := DefaultRootSnoozeTime
	if pluginConfig.RootSnoozeTime.Duration() != 0 {
		rootSnoozeTime = pluginConfig.RootSnoozeTime.Duration()
	}
	inflightCacheExpiry := DefaultInflightCacheExpiry
	if pluginConfig.InflightCacheExpiry.Duration() != 0 {
		inflightCacheExpiry = pluginConfig.InflightCacheExpiry.Duration()
	}
	if !common.IsHexAddress(spec.ContractID) {
		return nil, errors.Wrap(err, "spec.OffRampID is not a valid hex address")
	}
	verifier, err := commit_store.NewCommitStore(common.HexToAddress(pluginConfig.CommitStoreID), destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed creating a new onramp")
	}
	// Subscribe to the correct logs based on onramp type.
	onRampAddr := common.HexToAddress(pluginConfig.OnRampID)
	onRampType, _, err := TypeAndVersion(onRampAddr, sourceChain.Client())
	if err != nil {
		return nil, err
	}
	offRampAddr := common.HexToAddress(spec.ContractID)
	offRampType, _, err := TypeAndVersion(offRampAddr, destChain.Client())
	if err != nil {
		return nil, err
	}
	priceGetterObject, err := NewPriceGetter(pluginConfig.TokensPerFeeCoinPipeline, pr, jb.ID, jb.ExternalJobID, jb.Name.ValueOrZero(), lggr)
	if err != nil {
		return nil, err
	}
	var eventSignatures EventSignatures
	var wrappedPluginFactory ocrtypes.ReportingPluginFactory
	var offRampConfig evm_2_evm_offramp.IEVM2EVMOffRampOffRampConfig
	hashingCtx := hasher.NewKeccakCtx()
	switch onRampType {
	case EVM2EVMOnRamp:
		if offRampType != EVM2EVMOffRamp {
			return nil, errors.Errorf("invalid ramp combination %v and %v", onRampType, offRampType)
		}
		onRamp, err2 := evm_2_evm_onramp.NewEVM2EVMOnRamp(onRampAddr, sourceChain.Client())
		if err2 != nil {
			return nil, err2
		}
		offRamp, err2 := evm_2_evm_offramp.NewEVM2EVMOffRamp(offRampAddr, destChain.Client())
		if err2 != nil {
			return nil, err2
		}

		// subscribe for GasFeeUpdated logs, but FeeManager is only available as part of onchain offramp's config
		// TODO: how to detect if OffRampConfig.FeeManager changes on-chain? Currently, we expect a plugin/job/node restart
		offRampConfig, err2 = offRamp.GetOffRampConfig(nil)
		if err2 != nil {
			return nil, err2
		}
		feeManager, err2 := fee_manager.NewFeeManager(offRampConfig.FeeManager, destChain.Client())
		if err2 != nil {
			return nil, err2
		}

		eventSignatures = GetEventSignatures()
		wrappedPluginFactory = NewExecutionReportingPluginFactory(
			ExecutionPluginConfig{
				lggr:                lggr,
				source:              sourceChain.LogPoller(),
				dest:                destChain.LogPoller(),
				offRamp:             offRamp,
				onRamp:              onRamp,
				commitStore:         verifier,
				feeManager:          feeManager,
				builder:             NewBatchBuilder(lggr, eventSignatures, offRamp),
				eventSignatures:     eventSignatures,
				priceGetter:         priceGetterObject,
				leafHasher:          NewLeafHasher(pluginConfig.SourceChainID, pluginConfig.DestChainID, onRampAddr, hashingCtx),
				snoozeTime:          rootSnoozeTime,
				inflightCacheExpiry: inflightCacheExpiry,
				sourceChainID:       pluginConfig.SourceChainID,
				gasLimit:            BatchGasLimit,
				destGasEstimator:    destChain.TxManager().GetGasEstimator(),
				sourceGasEstimator:  sourceChain.TxManager().GetGasEstimator(),
			})
	default:
		return nil, errors.Errorf("unrecognized onramp, is %v the correct onramp address?", onRampAddr)
	}
	// Subscribe to all relevant commit logs.
	_, err = sourceChain.LogPoller().RegisterFilter(logpoller.Filter{EventSigs: []common.Hash{eventSignatures.SendRequested}, Addresses: []common.Address{onRampAddr}})
	if err != nil {
		return nil, err
	}
	_, err = destChain.LogPoller().RegisterFilter(logpoller.Filter{EventSigs: []common.Hash{ReportAccepted}, Addresses: []common.Address{verifier.Address()}})
	if err != nil {
		return nil, err
	}
	_, err = destChain.LogPoller().RegisterFilter(logpoller.Filter{EventSigs: []common.Hash{GasFeeUpdated}, Addresses: []common.Address{offRampConfig.FeeManager}})
	if err != nil {
		return nil, err
	}
	_, err = destChain.LogPoller().RegisterFilter(logpoller.Filter{EventSigs: []common.Hash{eventSignatures.ExecutionStateChanged}, Addresses: []common.Address{offRampAddr}})
	if err != nil {
		return nil, err
	}
	chainIDInterface, ok := spec.RelayConfig["chainID"]
	if !ok {
		return nil, errors.New("chainID must be provided in relay config")
	}
	chainID := int64(chainIDInterface.(float64))

	chain, err2 := chainSet.Get(big.NewInt(chainID))
	if err2 != nil {
		return nil, errors.Wrap(err2, "get chainset")
	}
	argsNoPlugin.ReportingPluginFactory = promwrapper.NewPromFactory(wrappedPluginFactory, "CCIPExecution", string(spec.Relay), chain.ID())
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
