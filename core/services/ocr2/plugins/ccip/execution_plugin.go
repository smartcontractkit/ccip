package ccip

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/any_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/blob_verifier"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins"
	ccipconfig "github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/config"
)

type CCIPExecution struct {
	lggr                               logger.Logger
	spec                               *job.OCR2OracleSpec
	sourceChainPoller, destChainPoller logpoller.LogPoller
	destChain                          evm.Chain
	blobVerifier                       *blob_verifier.BlobVerifier
	onRamp                             common.Address
	offRamp                            OffRamp
	batchBuilder                       BatchBuilder
	onRampSeqParser                    func(log logpoller.Log) (uint64, error)
}

var _ plugins.OraclePlugin = &CCIPExecution{}

func NewCCIPExecution(lggr logger.Logger, spec *job.OCR2OracleSpec, chainSet evm.ChainSet) (*CCIPExecution, error) {
	var pluginConfig ccipconfig.ExecutionPluginConfig
	err := json.Unmarshal(spec.PluginConfig.Bytes(), &pluginConfig)
	if err != nil {
		return &CCIPExecution{}, err
	}
	err = pluginConfig.ValidateExecutionPluginConfig()
	if err != nil {
		return &CCIPExecution{}, err
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
	if !common.IsHexAddress(spec.ContractID) {
		return nil, errors.Wrap(err, "spec.OffRampID is not a valid hex address")
	}
	verifier, err := blob_verifier.NewBlobVerifier(common.HexToAddress(pluginConfig.BlobVerifierID), destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed creating a new onramp")
	}
	// Subscribe to the correct logs based on onramp type.
	onRampAddr := common.HexToAddress(pluginConfig.OnRampID)
	onRampType, _, err := typeAndVersion(onRampAddr, sourceChain.Client())
	if err != nil {
		return nil, err
	}
	var onRampSeqParser func(log logpoller.Log) (uint64, error)
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
		// Subscribe to all relevant relay logs.
		sourceChain.LogPoller().MergeFilter([]common.Hash{CCIPSendRequested}, onRampAddr)
	case EVM2EVMSubscriptionOnRamp:
		// TODO: need event sigs for sub onramp
	default:
		return nil, errors.Errorf("unrecognized onramp, is %v the correct onramp address?", onRampAddr)
	}
	destChain.LogPoller().MergeFilter([]common.Hash{ReportAccepted}, verifier.Address())
	offRampType, _, _ := typeAndVersion(common.HexToAddress(spec.ContractID), destChain.Client())
	if err != nil {
		return nil, err
	}
	var batchBuilder BatchBuilder
	var offRamp OffRamp
	switch offRampType {
	case Any2EVMTollOffRamp:
		batchBuilder = NewTollBatchBuilder(lggr)
		offRampAddr := common.HexToAddress(spec.ContractID)
		tollOffRamp, err := any_2_evm_toll_offramp.NewAny2EVMTollOffRamp(offRampAddr, destChain.Client())
		if err != nil {
			return nil, err
		}
		offRamp = tollOffRamp
		destChain.LogPoller().MergeFilter([]common.Hash{CrossChainMessageExecuted}, offRampAddr) // May be common to all offramps?
	case Any2EVMSubscriptionOffRamp:
		// TODO: get sub fee token from subscription contract itself
		batchBuilder = NewSubscriptionBatchBuilder(lggr, common.Address{})
	default:
		return nil, errors.Errorf("unrecognized offramp, is %v the correct offramp address?", spec.ContractID)
	}
	// TODO: Can also check the on/offramp pair is compatible
	return &CCIPExecution{
		lggr:              lggr,
		spec:              spec,
		blobVerifier:      verifier,
		onRamp:            common.HexToAddress(pluginConfig.OnRampID),
		offRamp:           offRamp,
		sourceChainPoller: sourceChain.LogPoller(),
		destChainPoller:   destChain.LogPoller(),
		batchBuilder:      batchBuilder,
		onRampSeqParser:   onRampSeqParser,
	}, nil
}

type OffRamp interface {
	GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error)
	GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)
	GetExecutionState(opts *bind.CallOpts, arg0 uint64) (uint8, error)
	// TODO: make generic for sub offramp.
	ParseExecutionCompleted(log types.Log) (*any_2_evm_toll_offramp.Any2EVMTollOffRampExecutionCompleted, error)
}

func (c *CCIPExecution) GetPluginFactory() (plugin ocrtypes.ReportingPluginFactory, err error) {
	return NewExecutionReportingPluginFactory(
		c.lggr,
		c.onRamp,
		c.blobVerifier,
		c.sourceChainPoller,
		c.destChainPoller,
		common.HexToAddress(c.spec.ContractID),
		c.offRamp,
		c.batchBuilder,
		c.onRampSeqParser,
	), nil
}

func (c *CCIPExecution) GetServices() ([]job.ServiceCtx, error) {
	return []job.ServiceCtx{}, nil
}
