package ccip

import (
	"bytes"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/blob_verifier"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_subscription_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins"
	ccipconfig "github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/core/services/pipeline"
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
	reqEventSig                        common.Hash
	priceGetter                        PriceGetter
}

var _ plugins.OraclePlugin = &CCIPExecution{}

func NewCCIPExecution(lggr logger.Logger, jb job.Job, chainSet evm.ChainSet, pr pipeline.Runner) (*CCIPExecution, error) {
	spec := jb.OCR2OracleSpec
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
	var reqEventSig common.Hash
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
		sourceChain.LogPoller().MergeFilter([]common.Hash{CCIPTollSendRequested}, []common.Address{onRampAddr})
		reqEventSig = CCIPTollSendRequested
	case EVM2EVMSubscriptionOnRamp:
		onRamp, err2 := evm_2_evm_subscription_onramp.NewEVM2EVMSubscriptionOnRamp(onRampAddr, sourceChain.Client())
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
		sourceChain.LogPoller().MergeFilter([]common.Hash{CCIPSubSendRequested}, []common.Address{onRampAddr})
		reqEventSig = CCIPSubSendRequested
	default:
		return nil, errors.Errorf("unrecognized onramp, is %v the correct onramp address?", onRampAddr)
	}
	destChain.LogPoller().MergeFilter([]common.Hash{ReportAccepted}, []common.Address{verifier.Address()})
	offRampType, _, _ := typeAndVersion(common.HexToAddress(spec.ContractID), destChain.Client())
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
		batchBuilder = NewTollBatchBuilder(lggr)
		offRamp, err2 = NewTollOffRamp(common.HexToAddress(spec.ContractID), destChain)
	case EVM2EVMSubscriptionOffRamp:
		var subFeeToken common.Address
		offRamp, subFeeToken, err2 = NewSubOffRamp(common.HexToAddress(spec.ContractID), destChain)
		batchBuilder = NewSubscriptionBatchBuilder(lggr, subFeeToken, offRamp.(*subOffRamp))
	default:
		return nil, errors.Errorf("unrecognized offramp, is %v the correct offramp address?", spec.ContractID)
	}
	if err2 != nil {
		return nil, err
	}
	destChain.LogPoller().MergeFilter([]common.Hash{ExecutionStateChanged}, []common.Address{offRamp.Address()})
	// TODO: Can also check the on/offramp pair is compatible
	priceGetter, err := NewPriceGetter(pluginConfig.TokensPerFeeCoinPipeline, pr, jb.ID, jb.ExternalJobID, jb.Name.ValueOrZero(), lggr)
	if err2 != nil {
		return nil, err
	}
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
		reqEventSig:       reqEventSig,
		priceGetter:       priceGetter,
	}, nil
}

type OffRamp interface {
	GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error)
	GetDestinationTokens(opts *bind.CallOpts) ([]common.Address, error)
	GetDestinationToken(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)
	GetPool(opts *bind.CallOpts, sourceToken common.Address) (common.Address, error)
	GetExecutionState(opts *bind.CallOpts, arg0 uint64) (uint8, error)
	ParseSeqNumFromExecutionStateChanged(log types.Log) (uint64, error)
	Address() common.Address
	// Destination chain addresses.
	// Toll: dest pool addresses
	// Sub:  dest sub token address (not necessarily in a pool)
	GetSupportedTokensForExecutionFee() ([]common.Address, error)
}

type subOffRamp struct {
	*any_2_evm_subscription_offramp.EVM2EVMSubscriptionOffRamp
	router *any_2_evm_subscription_offramp_router.Any2EVMSubscriptionOffRampRouter
}

func (s subOffRamp) GetSupportedTokensForExecutionFee() ([]common.Address, error) {
	return s.router.GetSupportedTokensForExecutionFee(nil)
}

func (s subOffRamp) ParseSeqNumFromExecutionStateChanged(log types.Log) (uint64, error) {
	ec, err := s.ParseExecutionStateChanged(log)
	if err != nil {
		return 0, err
	}
	return ec.SequenceNumber, nil
}

func NewSubOffRamp(addr common.Address, destChain evm.Chain) (OffRamp, common.Address, error) {
	offRamp, err := any_2_evm_subscription_offramp.NewEVM2EVMSubscriptionOffRamp(addr, destChain.Client())
	if err != nil {
		return nil, common.Address{}, err
	}
	routerAddr, err := offRamp.SRouter(nil)
	if err != nil {
		return nil, common.Address{}, err
	}
	if bytes.Equal(routerAddr.Bytes(), common.Address{}.Bytes()) {
		return nil, common.Address{}, errors.New("router unset")
	}
	router, err := any_2_evm_subscription_offramp_router.NewAny2EVMSubscriptionOffRampRouter(routerAddr, destChain.Client())
	if err != nil {
		return nil, common.Address{}, err
	}
	subFeeToken, err := router.GetFeeToken(nil)
	if err != nil {
		return nil, common.Address{}, err
	}
	return &subOffRamp{EVM2EVMSubscriptionOffRamp: offRamp, router: router}, subFeeToken, nil
}

type tollOffRamp struct {
	*any_2_evm_toll_offramp.EVM2EVMTollOffRamp
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

func NewTollOffRamp(addr common.Address, destChain evm.Chain) (OffRamp, error) {
	offRamp, err := any_2_evm_toll_offramp.NewEVM2EVMTollOffRamp(addr, destChain.Client())
	if err != nil {
		return nil, err
	}
	return &tollOffRamp{offRamp}, nil
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
		c.reqEventSig,
		c.priceGetter,
	), nil
}

func (c *CCIPExecution) GetServices() ([]job.ServiceCtx, error) {
	return []job.ServiceCtx{}, nil
}
