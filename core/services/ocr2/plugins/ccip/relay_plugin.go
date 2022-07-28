package ccip

import (
	"encoding/json"
	"math/big"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/core/chains/evm"
	eth "github.com/smartcontractkit/chainlink/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/blob_verifier"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/evm_2_evm_subscription_onramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/evm_2_evm_toll_onramp"
	type_and_version "github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/type_and_version_interface_wrapper"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins"
	ccipconfig "github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/config"
)

type CCIPRelay struct {
	lggr                logger.Logger
	spec                *job.OCR2OracleSpec
	sourceChainPoller   logpoller.LogPoller
	blobVerifier        *blob_verifier.BlobVerifier
	onRamps             []common.Address
	onRampSeqParsers    map[common.Address]func(log logpoller.Log) (uint64, error)
	onRampToReqEventSig map[common.Address]common.Hash
}

var _ plugins.OraclePlugin = &CCIPRelay{}

type ContractType string

var (
	EVM2EVMTollOnRamp          ContractType = "EVM2EVMTollOnRamp"
	Any2EVMTollOffRamp         ContractType = "Any2EVMTollOffRamp"
	EVM2EVMSubscriptionOnRamp  ContractType = "EVM2EVMSubscriptionOnRamp"
	Any2EVMSubscriptionOffRamp ContractType = "Any2EVMSubscriptionOffRamp"
	ContractTypes                           = map[ContractType]struct{}{
		EVM2EVMTollOnRamp:          {},
		Any2EVMTollOffRamp:         {},
		EVM2EVMSubscriptionOnRamp:  {},
		Any2EVMSubscriptionOffRamp: {},
	}
)

func typeAndVersion(addr common.Address, client eth.Client) (ContractType, semver.Version, error) {
	tv, err := type_and_version.NewTypeAndVersionInterface(addr, client)
	if err != nil {
		return "", semver.Version{}, errors.Wrap(err, "failed creating a type and version")
	}
	tvStr, err := tv.TypeAndVersion(nil)
	if err != nil {
		return "", semver.Version{}, errors.Wrap(err, "failed to call type and version")
	}
	typeAndVersion := strings.Split(tvStr, " ")
	contractType, version := typeAndVersion[0], typeAndVersion[1]
	v, err := semver.NewVersion(version)
	if err != nil {
		return "", semver.Version{}, err
	}
	if _, ok := ContractTypes[ContractType(contractType)]; !ok {
		return "", semver.Version{}, errors.Errorf("unrecognized contract type %v", contractType)
	}
	return ContractType(contractType), *v, nil
}

func NewCCIPRelay(lggr logger.Logger, spec *job.OCR2OracleSpec, chainSet evm.ChainSet) (*CCIPRelay, error) {
	var pluginConfig ccipconfig.RelayPluginConfig
	err := json.Unmarshal(spec.PluginConfig.Bytes(), &pluginConfig)
	if err != nil {
		return &CCIPRelay{}, err
	}
	err = pluginConfig.ValidateRelayPluginConfig()
	if err != nil {
		return &CCIPRelay{}, err
	}
	lggr.Infof("CCIP relay plugin initialized with offchainConfig: %+v", pluginConfig)

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
	blobVerifier, err := blob_verifier.NewBlobVerifier(common.HexToAddress(spec.ContractID), destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed creating a new offramp")
	}
	onRampSeqParsers := make(map[common.Address]func(log logpoller.Log) (uint64, error))
	onRampToReqEventSig := make(map[common.Address]common.Hash)
	var onRamps []common.Address
	for _, onRampID := range pluginConfig.OnRampIDs {
		addr := common.HexToAddress(onRampID)
		onRamps = append(onRamps, addr)
		contractType, _, _ := typeAndVersion(addr, sourceChain.Client())
		switch contractType {
		case EVM2EVMTollOnRamp:
			onRamp, err := evm_2_evm_toll_onramp.NewEVM2EVMTollOnRamp(addr, sourceChain.Client())
			if err != nil {
				return nil, errors.Wrap(err, "failed creating a new onramp")
			}
			onRampSeqParsers[common.HexToAddress(onRampID)] = func(log logpoller.Log) (uint64, error) {
				req, err := onRamp.ParseCCIPSendRequested(types.Log{Data: log.Data, Topics: log.GetTopics()})
				if err != nil {
					return 0, err
				}
				return req.Message.SequenceNumber, nil
			}
			// Subscribe to all relevant relay logs.
			sourceChain.LogPoller().MergeFilter([]common.Hash{CCIPSendRequested}, onRamp.Address())
			onRampToReqEventSig[onRamp.Address()] = CCIPSendRequested
		case EVM2EVMSubscriptionOnRamp:
			onRamp, err := evm_2_evm_subscription_onramp.NewEVM2EVMSubscriptionOnRamp(addr, sourceChain.Client())
			if err != nil {
				return nil, errors.Wrap(err, "failed creating a new onramp")
			}
			onRampSeqParsers[common.HexToAddress(onRampID)] = func(log logpoller.Log) (uint64, error) {
				req, err := onRamp.ParseCCIPSendRequested(types.Log{Data: log.Data, Topics: log.GetTopics()})
				if err != nil {
					return 0, err
				}
				return req.Message.SequenceNumber, nil
			}
			// Subscribe to all relevant relay logs.
			sourceChain.LogPoller().MergeFilter([]common.Hash{CCIPSubSendRequested}, onRamp.Address())
			onRampToReqEventSig[onRamp.Address()] = CCIPSubSendRequested
		default:
			return nil, errors.Errorf("unrecognized onramp %v", onRampID)
		}
	}
	return &CCIPRelay{
		lggr:                lggr,
		blobVerifier:        blobVerifier,
		onRampSeqParsers:    onRampSeqParsers,
		onRampToReqEventSig: onRampToReqEventSig,
		sourceChainPoller:   sourceChain.LogPoller(),
		onRamps:             onRamps,
	}, nil
}

func (c *CCIPRelay) GetPluginFactory() (plugin ocrtypes.ReportingPluginFactory, err error) {
	return NewRelayReportingPluginFactory(c.lggr, c.sourceChainPoller, c.blobVerifier, c.onRampSeqParsers, c.onRampToReqEventSig, c.onRamps), nil
}

func (c *CCIPRelay) GetServices() ([]job.ServiceCtx, error) {
	return []job.ServiceCtx{}, nil
}
