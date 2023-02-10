package ccip

import (
	"encoding/json"
	"math/big"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	libocr2 "github.com/smartcontractkit/libocr/offchainreporting2"

	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_onramp"
	type_and_version "github.com/smartcontractkit/chainlink/core/gethwrappers/generated/type_and_version_interface_wrapper"
	"github.com/smartcontractkit/chainlink/core/logger"
	hlp "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/core/services/job"
	ccipconfig "github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/promwrapper"
)

type ContractType string

var (
	EVM2EVMOnRamp  ContractType = "EVM2EVMOnRamp"
	EVM2EVMOffRamp ContractType = "EVM2EVMOffRamp"
	CommitStore    ContractType = "CommitStore"
	Router         ContractType = "Router"
	ContractTypes               = map[ContractType]struct{}{
		EVM2EVMOffRamp: {},
		EVM2EVMOnRamp:  {},
		CommitStore:    {},
	}
)

func TypeAndVersion(addr common.Address, client bind.ContractBackend) (ContractType, semver.Version, error) {
	tv, err := type_and_version.NewTypeAndVersionInterface(addr, client)
	if err != nil {
		return "", semver.Version{}, errors.Wrap(err, "failed creating a type and version")
	}
	tvStr, err := tv.TypeAndVersion(nil)
	if err != nil {
		return "", semver.Version{}, errors.Wrap(err, "failed to call type and version")
	}
	typeAndVersionValues := strings.Split(tvStr, " ")
	contractType, version := typeAndVersionValues[0], typeAndVersionValues[1]
	v, err := semver.NewVersion(version)
	if err != nil {
		return "", semver.Version{}, err
	}
	if _, ok := ContractTypes[ContractType(contractType)]; !ok {
		return "", semver.Version{}, errors.Errorf("unrecognized contract type %v", contractType)
	}
	return ContractType(contractType), *v, nil
}

func NewCommitServices(lggr logger.Logger, spec *job.OCR2OracleSpec, chainSet evm.ChainSet, new bool, argsNoPlugin libocr2.OracleArgs) ([]job.ServiceCtx, error) {
	var pluginConfig ccipconfig.CommitPluginConfig
	err := json.Unmarshal(spec.PluginConfig.Bytes(), &pluginConfig)
	if err != nil {
		return nil, err
	}
	err = pluginConfig.ValidateCommitPluginConfig()
	if err != nil {
		return nil, err
	}
	lggr.Infof("CCIP commit plugin initialized with offchainConfig: %+v", pluginConfig)

	sourceChain, err := chainSet.Get(big.NewInt(0).SetUint64(pluginConfig.SourceChainID))
	if err != nil {
		return nil, errors.Wrap(err, "unable to open source chain")
	}
	destChain, err := chainSet.Get(big.NewInt(0).SetUint64(pluginConfig.DestChainID))
	if err != nil {
		return nil, errors.Wrap(err, "unable to open destination chain")
	}

	lggr = lggr.With("srcChain", hlp.ChainName(int64(pluginConfig.SourceChainID)), "dstChain", hlp.ChainName(int64(pluginConfig.DestChainID)))

	inflightCacheExpiry := DefaultInflightCacheExpiry
	if pluginConfig.InflightCacheExpiry.Duration() != 0 {
		inflightCacheExpiry = pluginConfig.InflightCacheExpiry.Duration()
	}

	if !common.IsHexAddress(spec.ContractID) {
		return nil, errors.Wrap(err, "spec.ContractID is not a valid hex address")
	}
	commitStore, err := commit_store.NewCommitStore(common.HexToAddress(spec.ContractID), destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed loading the commitStore")
	}
	onRampSeqParsers := make(map[common.Address]func(log logpoller.Log) (uint64, error))
	onRampToReqEventSig := make(map[common.Address]EventSignatures)
	var onRamps []common.Address
	var onRampToHasher = make(map[common.Address]LeafHasherInterface[[32]byte])
	hashingCtx := hasher.NewKeccakCtx()

	for _, onRampID := range pluginConfig.OnRampIDs {
		addr := common.HexToAddress(onRampID)
		onRamps = append(onRamps, addr)
		contractType, _, err2 := TypeAndVersion(addr, sourceChain.Client())
		if err2 != nil {
			return nil, errors.Errorf("failed getting type and version %v", err2)
		}

		switch contractType {
		case EVM2EVMOnRamp:
			onRamp, err3 := evm_2_evm_onramp.NewEVM2EVMOnRamp(addr, sourceChain.Client())
			if err3 != nil {
				return nil, errors.Wrap(err3, "failed creating a new onramp")
			}
			onRampSeqParsers[addr] = func(log logpoller.Log) (uint64, error) {
				req, err4 := onRamp.ParseCCIPSendRequested(types.Log{Data: log.Data, Topics: log.GetTopics()})
				if err4 != nil {
					lggr.Warnf("failed to parse log: %+v", log)
					return 0, err4
				}
				return req.Message.SequenceNumber, nil
			}
			onRampToReqEventSig[addr] = GetEventSignatures()
			onRampToHasher[addr] = NewLeafHasher(pluginConfig.SourceChainID, pluginConfig.DestChainID, addr, hashingCtx)
		default:
			return nil, errors.Errorf("unrecognized onramp %v", onRampID)
		}
		// Subscribe to all relevant commit logs.
		_, err = sourceChain.LogPoller().RegisterFilter(logpoller.Filter{EventSigs: []common.Hash{onRampToReqEventSig[addr].SendRequested}, Addresses: []common.Address{addr}})
		if err != nil {
			return nil, err
		}
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
	wrappedPluginFactory := NewCommitReportingPluginFactory(lggr, sourceChain.LogPoller(), commitStore, onRampSeqParsers, onRampToReqEventSig, onRamps, onRampToHasher, inflightCacheExpiry)
	argsNoPlugin.ReportingPluginFactory = promwrapper.NewPromFactory(wrappedPluginFactory, "CCIPCommit", string(spec.Relay), chain.ID())
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
