package ccip

import (
	"encoding/json"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	libocr2 "github.com/smartcontractkit/libocr/offchainreporting2"

	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_ge_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_ge_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/job"
	ccipconfig "github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/hasher"
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
	sourceChainId, destChainId := big.NewInt(0).SetUint64(pluginConfig.SourceChainID), big.NewInt(0).SetUint64(pluginConfig.DestChainID)
	sourceChain, err := chainSet.Get(sourceChainId)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open source chain")
	}
	destChain, err := chainSet.Get(destChainId)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open destination chain")
	}

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
	var onRampSeqParser func(log logpoller.Log) (uint64, error)
	var eventSignatures EventSignatures
	var onRampToHasher = make(map[common.Address]LeafHasher[[32]byte])
	hashingCtx := hasher.NewKeccakCtx()

	switch onRampType {
	case EVM2EVMTollOnRamp:
		onRamp, err2 := evm_2_evm_toll_onramp.NewEVM2EVMTollOnRamp(onRampAddr, sourceChain.Client())
		if err2 != nil {
			return nil, err2
		}
		onRampSeqParser = func(log logpoller.Log) (uint64, error) {
			req, err3 := onRamp.ParseCCIPSendRequested(types.Log{Data: log.Data, Topics: log.GetTopics()})
			if err3 != nil {
				return 0, err3
			}
			return req.Message.SequenceNumber, nil
		}
		eventSignatures = GetTollEventSignatures()
		onRampToHasher[onRampAddr] = NewTollLeafHasher(sourceChainId, destChainId, onRampAddr, hashingCtx)
	case EVM2EVMGEOnRamp:
		onRamp, err2 := evm_2_evm_ge_onramp.NewEVM2EVMGEOnRamp(onRampAddr, sourceChain.Client())
		if err2 != nil {
			return nil, err2
		}
		onRampSeqParser = func(log logpoller.Log) (uint64, error) {
			req, err3 := onRamp.ParseCCIPSendRequested(types.Log{Data: log.Data, Topics: log.GetTopics()})
			if err3 != nil {
				return 0, err3
			}
			return req.Message.SequenceNumber, nil
		}
		eventSignatures = GetGEEventSignatures()
		onRampToHasher[onRampAddr] = NewGELeafHasher(sourceChainId, destChainId, onRampAddr, hashingCtx)
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
	offRampType, _, err := TypeAndVersion(common.HexToAddress(spec.ContractID), destChain.Client())
	if err != nil {
		return nil, err
	}
	var (
		batchBuilder BatchBuilder
		offRamp      OffRamp
		err2         error
	)
	switch offRampType {
	case EVM2EVMTollOffRamp:
		batchBuilder = NewTollBatchBuilder(lggr, eventSignatures)
		offRamp, err2 = NewTollOffRamp(common.HexToAddress(spec.ContractID), destChain)
	case EVM2EVMGEOffRamp:
		batchBuilder = NewGEBatchBuilder(lggr, eventSignatures)
		offRamp, err2 = NewGEOffRamp(common.HexToAddress(spec.ContractID), destChain)
	default:
		return nil, errors.Errorf("unrecognized offramp, is %v the correct offramp address?", spec.ContractID)
	}
	if err2 != nil {
		return nil, err
	}
	_, err = destChain.LogPoller().RegisterFilter(logpoller.Filter{EventSigs: []common.Hash{eventSignatures.ExecutionStateChanged}, Addresses: []common.Address{offRamp.Address()}})
	if err != nil {
		return nil, err
	}
	// TODO: Can also check the on/offramp pair is compatible
	priceGetterObject, err := NewPriceGetter(pluginConfig.TokensPerFeeCoinPipeline, pr, jb.ID, jb.ExternalJobID, jb.Name.ValueOrZero(), lggr)
	if err2 != nil {
		return nil, err
	}

	argsNoPlugin.ReportingPluginFactory = NewExecutionReportingPluginFactory(
		lggr,
		onRampAddr,
		verifier,
		sourceChain.LogPoller(), destChain.LogPoller(),
		common.HexToAddress(spec.ContractID),
		offRamp,
		batchBuilder,
		onRampSeqParser,
		eventSignatures,
		priceGetterObject,
		onRampToHasher,
		rootSnoozeTime,
		inflightCacheExpiry,
		sourceChain.TxManager().GetGasEstimator(),
		destChain.TxManager().GetGasEstimator(),
		sourceChainId,
	)
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

type OffRamp interface {
	GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error)
	GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)
	GetExecutionState(opts *bind.CallOpts, arg0 uint64) (uint8, error)
	ParseSeqNumFromExecutionStateChanged(log types.Log) (uint64, error)
	Address() common.Address
	// Destination chain addresses.
	// Toll: dest pool addresses
	GetSupportedTokensForExecutionFee() ([]common.Address, error)
	GetAllowedTokensAmount(opts *bind.CallOpts) (*big.Int, error)
	GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error)
}

type tollOffRamp struct {
	*evm_2_evm_toll_offramp.EVM2EVMTollOffRamp
}

func (s tollOffRamp) ParseSeqNumFromExecutionStateChanged(log types.Log) (uint64, error) {
	ec, err := s.ParseExecutionStateChanged(log)
	if err != nil {
		return 0, err
	}
	return ec.SequenceNumber, nil
}

func (s tollOffRamp) GetSupportedTokensForExecutionFee() ([]common.Address, error) {
	// TODO: Toll offramp contract is missing ExecConfig?
	// for now support all source tokens as fee tokens
	return s.EVM2EVMTollOffRamp.GetDestinationTokens(nil)
}

func (s tollOffRamp) GetAllowedTokensAmount(opts *bind.CallOpts) (*big.Int, error) {
	bucket, err := s.EVM2EVMTollOffRamp.CalculateCurrentTokenBucketState(opts)
	if err != nil {
		return nil, err
	}
	return bucket.Tokens, nil
}

func NewTollOffRamp(addr common.Address, destChain evm.Chain) (OffRamp, error) {
	offRamp, err := evm_2_evm_toll_offramp.NewEVM2EVMTollOffRamp(addr, destChain.Client())
	if err != nil {
		return nil, err
	}
	return &tollOffRamp{offRamp}, nil
}

func NewGEOffRamp(addr common.Address, destChain evm.Chain) (OffRamp, error) {
	offRamp, err := evm_2_evm_ge_offramp.NewEVM2EVMGEOffRamp(addr, destChain.Client())
	if err != nil {
		return nil, err
	}
	return &geOffRamp{offRamp}, nil
}

type geOffRamp struct {
	*evm_2_evm_ge_offramp.EVM2EVMGEOffRamp
}

func (s geOffRamp) ParseSeqNumFromExecutionStateChanged(log types.Log) (uint64, error) {
	ec, err := s.ParseExecutionStateChanged(log)
	if err != nil {
		return 0, err
	}
	return ec.SequenceNumber, nil
}

func (s geOffRamp) GetSupportedTokensForExecutionFee() ([]common.Address, error) {
	// TODO: Toll offramp contract is missing ExecConfig?
	// for now support all source tokens as fee tokens
	return s.EVM2EVMGEOffRamp.GetDestinationTokens(nil)
}

func (s geOffRamp) GetAllowedTokensAmount(opts *bind.CallOpts) (*big.Int, error) {
	bucket, err := s.EVM2EVMGEOffRamp.CalculateCurrentTokenBucketState(opts)
	if err != nil {
		return nil, err
	}
	return bucket.Tokens, nil
}
