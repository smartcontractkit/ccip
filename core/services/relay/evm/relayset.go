package evm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/x_internal/ccipcalc"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/x_internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/x_internal/ccipdata/factory"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/x_internal/pricegetter"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/x_internal/rpclib"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"golang.org/x/exp/maps"
	"math/big"
	"strconv"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config" //TODO: does this break the dependency graph?
)

type RelayID struct {
	Network string // TODO: consider removing?
	ChainID string
}

type RelayerSet interface {
	Get(RelayID) (Relayer, error)
	GetAll() (map[RelayID]Relayer, error)
}

type CCIPRelayerSet struct {
	relayerMap    map[RelayID]Relayer
	sourceRelayer Relayer
	destRelayer   Relayer
}

func chainIDToRelayerID(rs RelayerSet, chainID string) (RelayID, error) {
	relayerMap, err := rs.GetAll()
	if err != nil {
		return RelayID{}, err
	}

	rids := maps.Keys(relayerMap)
	for _, rid := range rids {
		if rid.ChainID == chainID {
			return rid, nil
		}
	}

	return RelayID{}, fmt.Errorf("chain ID '%s' not found", chainID)

}

func (rs *CCIPRelayerSet) Get(relayID RelayID) (Relayer, error) {
	return rs.relayerMap[relayID], nil
}

func (rs *CCIPRelayerSet) GetAll() (map[RelayID]Relayer, error) {
	return rs.relayerMap, nil
}

func (rs *CCIPRelayerSet) NewCCIPCommitProvider_V2(context context.Context, rargs commontypes.RelayArgs, pargs commontypes.PluginArgs) (commontypes.CCIPCommitProvider, error) {
	var pluginConfig ccipconfig.CommitPluginJobSpecConfig
	err := json.Unmarshal(pargs.PluginConfig, &pluginConfig)
	if err != nil {
		return nil, err
	}

	offRampAddress := string(pluginConfig.OffRamp)

	// Build price getter clients for all chains specified in the aggregator configurations.
	// Some lanes (e.g. Wemix/Kroma) requires other clients than source and destination, since they use feeds from other chains.
	// TODO: Double check i've wired this in from Delegate
	priceGetterClients := map[uint64]pricegetter.DynamicPriceGetterClient{}
	for _, aggCfg := range pluginConfig.PriceGetterConfig.AggregatorPrices {
		chainID := aggCfg.ChainID
		// Retrieve the chain.
		aggRelayerID, err2 := chainIDToRelayerID(rs, strconv.FormatUint(chainID, 10))
		if err2 != nil {
			return nil, fmt.Errorf("retrieving chain for chainID %d: %w", chainID, err2)
		}

		aggRelayer, err2 := rs.Get(aggRelayerID)
		if err2 != nil {
			return nil, fmt.Errorf("retrieving relayer for relayerID %v: %w", aggRelayerID, err2)
		}

		caller := rpclib.NewDynamicLimitedBatchCaller(
			lggr,
			aggRelayer.chain.Client(),
			rpclib.DefaultRpcBatchSizeLimit,
			rpclib.DefaultRpcBatchBackOffMultiplier,
			rpclib.DefaultMaxParallelRpcCalls,
		)
		priceGetterClients[chainID] = pricegetter.NewDynamicPriceGetterClient(caller)
	}

	return &EVMCCIPCommitProviderImpl_V2{
		offRampAddress:     offRampAddress,
		priceGetterClients: priceGetterClients,
		priceGetterConfig:  *pluginConfig.PriceGetterConfig,
	}, nil
}

type EVMCCIPCommitProviderImpl_V2 struct {
	lggr               logger.Logger
	sourceLP           logpoller.LogPoller
	destLP             logpoller.LogPoller
	sourceStartBlock   uint64
	destStartBlock     uint64
	commitStoreAddress common.Address
	offRampAddress     string
	new                bool
	sourceClient       client.Client
	destClient         client.Client
	sourceGasEstimator gas.EvmFeeEstimator
	destGasEstimator   gas.EvmFeeEstimator
	sourceMaxGasPrice  big.Int
	destMaxGasPrice    big.Int
	priceGetterClients map[uint64]pricegetter.DynamicPriceGetterClient
	priceGetterConfig  config.DynamicPriceGetterConfig
	versionFinder      factory.VersionFinder
}

func (E EVMCCIPCommitProviderImpl_V2) Name() string {
	return "EVMCCIPCommitProvider"
}

func (E EVMCCIPCommitProviderImpl_V2) Start(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl_V2) Close() error {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl_V2) Ready() error {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl_V2) HealthReport() map[string]error {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl_V2) OffchainConfigDigester() ocrtypes.OffchainConfigDigester {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl_V2) ContractConfigTracker() ocrtypes.ContractConfigTracker {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl_V2) ContractTransmitter() ocrtypes.ContractTransmitter {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl_V2) ChainReader() commontypes.ChainReader {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl_V2) Codec() commontypes.Codec {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl_V2) NewCommitStoreReader(ctx context.Context, addr cciptypes.Address) (commitStoreReader cciptypes.CommitStoreReader, err error) {
	commitStoreReader, err = factory.NewCommitStoreReader(E.lggr, E.versionFinder, addr, E.destClient, E.destLP, E.sourceGasEstimator, &E.sourceMaxGasPrice, nil)
	return
}

func (E EVMCCIPCommitProviderImpl_V2) NewOffRampReader(ctx context.Context, offRampAddr cciptypes.Address) (offRampReader cciptypes.OffRampReader, err error) {
	offRampReader, err = factory.NewOffRampReader(E.lggr, E.versionFinder, offRampAddr, E.destClient, E.destLP, E.destGasEstimator, &E.destMaxGasPrice, true, nil)
	return
}

func (E EVMCCIPCommitProviderImpl_V2) NewOnRampReader(ctx context.Context, onRampAddress cciptypes.Address, cfg cciptypes.CommitStoreStaticConfig) (onRampReader cciptypes.OnRampReader, err error) {
	versionFinder := factory.NewEvmVersionFinder()
	onRampReader, err = factory.NewOnRampReader(E.lggr, versionFinder, cfg.SourceChainSelector, cfg.ChainSelector, onRampAddress, E.sourceLP, E.sourceClient, nil)
	return
}

// Dynamic Price Getter for CCIP commit service
func (E EVMCCIPCommitProviderImpl_V2) NewPriceGetter(ctx context.Context) (priceGetter cciptypes.PriceGetter, err error) {
	priceGetter, err = pricegetter.NewDynamicPriceGetter(E.priceGetterConfig, E.priceGetterClients)
	return
}

func (E EVMCCIPCommitProviderImpl_V2) NewPriceRegistryReader(ctx context.Context, addr cciptypes.Address) (cciptypes.PriceRegistryReader, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl_V2) SourceNativeToken(ctx context.Context, sourceRouterAddr cciptypes.Address) (cciptypes.Address, error) {
	sourceRouterAddrHex := ccipAddressToCCIPAddress(sourceRouterAddr)
	sourceRouter, err := router.NewRouter(sourceRouterAddrHex, E.sourceClient)
	if err != nil {
		return "", err
	}
	sourceNative, err := sourceRouter.GetWrappedNative(&bind.CallOpts{Context: ctx})
	if err != nil {
		return "", err
	}

	return commonAddressToCCIPAddress(sourceNative)
}

func (E EVMCCIPCommitProviderImpl_V2) NewOffRampReaders(ctx context.Context, offRampReader cciptypes.OffRampReader) (offRampReaders []cciptypes.OffRampReader, err error) {
	// Look up all destination offRamps connected to the same router
	destRouterAddr, err := offRampReader.GetRouter(ctx)
	if err != nil {
		return nil, err
	}
	destRouterEvmAddr, err := ccipcalc.GenericAddrToEvm(destRouterAddr)
	if err != nil {
		return nil, err
	}
	destRouter, err := router.NewRouter(destRouterEvmAddr, E.destClient)
	if err != nil {
		return nil, err
	}
	destRouterOffRamps, err := destRouter.GetOffRamps(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	var destOffRampReaders []ccipdata.OffRampReader
	for _, o := range destRouterOffRamps {
		destOffRampAddr := cciptypes.Address(o.OffRamp.String())
		destOffRampReader, err2 := factory.NewOffRampReader(
			E.lggr,
			E.versionFinder,
			destOffRampAddr,
			E.destClient,
			E.destLP,
			E.destGasEstimator,
			&E.destMaxGasPrice,
			true,
			nil,
		)
		if err2 != nil {
			return nil, err2
		}

		destOffRampReaders = append(destOffRampReaders, destOffRampReader)
	}

	// convert internal CCIP OffRampReader type to common type
	offRampReaders = make([]cciptypes.OffRampReader, 0, len(destOffRampReaders))
	for _, d := range destOffRampReaders {
		offRampReaders = append(offRampReaders, d)
	}
	return
}

func (E EVMCCIPCommitProviderImpl_V2) GetStaticConfig(ctx context.Context, addr cciptypes.Address) (cciptypes.CommitStoreStaticConfig, error) {
	commitStoreAddress := common.HexToAddress(string(addr))
	staticConfig, err := ccipdata.FetchCommitStoreStaticConfig(commitStoreAddress, E.destClient)
	if err != nil {
		return cciptypes.CommitStoreStaticConfig{}, fmt.Errorf("get commit store static config: %w", err)
	}

	return cciptypes.CommitStoreStaticConfig{
		ChainSelector:       staticConfig.ChainSelector,
		SourceChainSelector: staticConfig.SourceChainSelector,
		OnRamp:              cciptypes.Address(staticConfig.OnRamp.String()),
		ArmProxy:            cciptypes.Address(staticConfig.ArmProxy.String()),
	}, nil
}
