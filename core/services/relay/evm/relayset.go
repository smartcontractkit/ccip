package evm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	types "github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
	"github.com/smartcontractkit/chainlink-common/pkg/types/core"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"go.uber.org/multierr"
	"golang.org/x/exp/maps"
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config" //TODO: does this break the dependency graph?
)

type RelayerSet interface {
	Get(ctx context.Context, relayID types.RelayID) (Relayer, error)

	// List lists the relayers corresponding to `...types.RelayID`
	// returning all relayers if len(...types.RelayID) == 0.
	List(ctx context.Context, relayIDs ...types.RelayID) (map[types.RelayID]Relayer, error)
	NewCrossRelayerPluginProvider(context.Context, types.RelayArgs, types.PluginArgs) (types.PluginProvider, error)
}

type EVMRelayerSet struct {
	relayerMap map[types.RelayID]Relayer
	lggr       logger.Logger
}

func NewEVMRelayerSet(relayerMap map[types.RelayID]Relayer, lggr logger.Logger) *EVMRelayerSet {
	return &EVMRelayerSet{
		relayerMap: relayerMap,
		lggr:       lggr,
	}
}

func chainIDToRelayerID(ctx context.Context, rs core.RelayerSet, chainID string) (types.RelayID, error) {
	relayerMap, err := rs.List(ctx)
	if err != nil {
		return types.RelayID{}, err
	}

	rids := maps.Keys(relayerMap)
	for _, rid := range rids {
		if rid.ChainID == chainID {
			return rid, nil
		}
	}

	return types.RelayID{}, fmt.Errorf("chain ID '%s' not found", chainID)

}

func (rs *EVMRelayerSet) Get(ctx context.Context, relayID types.RelayID) (Relayer, error) {
	return rs.relayerMap[relayID], nil
}

func (rs *EVMRelayerSet) List(ctx context.Context, relayIDs ...types.RelayID) (map[types.RelayID]Relayer, error) {
	if len(relayIDs) == 0 {
		return rs.relayerMap, nil
	}

	subset := make(map[types.RelayID]Relayer)
	// subset relayer ids
	for _, srid := range relayIDs {
		subset[srid] = Relayer{}
	}

	// target relayer ids
	for trid := range rs.relayerMap {
		_, ok := subset[trid]
		if ok {
			subset[trid] = rs.relayerMap[trid]
		}
	}

	return subset, nil
}

func (rs *EVMRelayerSet) NewCrossRelayerPluginProvider(ctx context.Context, rargs types.RelayArgs, pargs types.PluginArgs) (types.PluginProvider, error) {
	switch rargs.ProviderType {
	case string(types.CCIPCommit):
		return rs.NewCCIPCommitProvider(ctx, rargs, pargs)
	}

	return nil, fmt.Errorf("unsupported relayer type: %s", rargs.ProviderType)
}

func (rs *EVMRelayerSet) NewCCIPCommitProvider(ctx context.Context, rargs types.RelayArgs, pargs types.PluginArgs) (types.CCIPCommitProvider, error) {
	var commitPluginConfig ccipconfig.CommitPluginConfig
	err := json.Unmarshal(pargs.PluginConfig, &commitPluginConfig)
	if err != nil {
		return nil, err
	}

	sourceRelayer, err := rs.Get(ctx, commitPluginConfig.SourceRelayerID)
	if err != nil {
		return nil, err
	}

	destRelayer, err := rs.Get(ctx, commitPluginConfig.DestRelayerID)
	if err != nil {
		return nil, err
	}

	lggr := rs.lggr
	commitStoreAddress := rargs.ContractID

	var pluginConfig ccipconfig.CommitPluginJobSpecConfig
	err = json.Unmarshal(pargs.PluginConfig, &pluginConfig)
	if err != nil {
		return nil, err
	}

	offRampAddress := string(pluginConfig.OffRamp)
	sourceStartBlock := pluginConfig.SourceStartBlock
	destStartBlock := pluginConfig.DestStartBlock

	// Build price getter clients for all chains specified in the aggregator configurations.
	// Some lanes (e.g. Wemix/Kroma) requires other clients than source and destination, since they use feeds from other chains.
	// TODO: Double check i've wired this in from Delegate (i.e. creating the RelayerSet with all the necessary relayers)
	priceGetterClients := map[uint64]ccip.DynamicPriceGetterClient{}
	for _, aggCfg := range pluginConfig.PriceGetterConfig.AggregatorPrices {
		chainID := aggCfg.ChainID
		// Retrieve the chain.
		aggRelayerID, err2 := chainIDToRelayerID(ctx, rs, strconv.FormatUint(chainID, 10))
		if err2 != nil {
			return nil, fmt.Errorf("retrieving chain for chainID %d: %w", chainID, err2)
		}

		aggRelayer, err2 := rs.Get(ctx, aggRelayerID)
		if err2 != nil {
			return nil, fmt.Errorf("retrieving relayer for relayerID %v: %w", aggRelayerID, err2)
		}

		caller := ccip.NewDynamicLimitedBatchCaller(
			lggr,
			aggRelayer.chain.Client(),
			uint(ccip.DefaultRpcBatchSizeLimit),
			uint(ccip.DefaultRpcBatchBackOffMultiplier),
			uint(ccip.DefaultMaxParallelRpcCalls),
		)
		priceGetterClients[chainID] = ccip.NewDynamicPriceGetterClient(caller)
	}

	return EVMCCIPCommitProviderImpl_V2{
		lggr:               lggr,
		sourceLP:           sourceRelayer.chain.LogPoller(),
		destLP:             destRelayer.chain.LogPoller(),
		sourceStartBlock:   sourceStartBlock,
		destStartBlock:     destStartBlock,
		commitStoreAddress: commitStoreAddress,
		offRampAddress:     offRampAddress,
		sourceClient:       sourceRelayer.chain.Client(),
		destClient:         destRelayer.chain.Client(),
		sourceGasEstimator: sourceRelayer.chain.GasEstimator(),
		destGasEstimator:   destRelayer.chain.GasEstimator(),
		sourceMaxGasPrice:  *sourceRelayer.chain.Config().EVM().GasEstimator().PriceMax().ToInt(),
		destMaxGasPrice:    *destRelayer.chain.Config().EVM().GasEstimator().PriceMax().ToInt(),
		sourceCodec:        sourceRelayer.codec,
		sourceChainReader:  sourceRelayer.chainReader,
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
	commitStoreAddress string
	offRampAddress     string
	sourceClient       client.Client
	destClient         client.Client
	sourceGasEstimator gas.EvmFeeEstimator
	destGasEstimator   gas.EvmFeeEstimator
	sourceMaxGasPrice  big.Int
	destMaxGasPrice    big.Int
	sourceCodec        types.Codec
	sourceChainReader  types.ChainReader
	priceGetterClients map[uint64]ccip.DynamicPriceGetterClient
	priceGetterConfig  config.DynamicPriceGetterConfig
	versionFinder      ccip.VersionFinder
	s                  services.Service
	cp                 types.ConfigProvider
}

func (E EVMCCIPCommitProviderImpl_V2) Name() string {
	return "EVMCCIPCommitProvider"
}

func (E EVMCCIPCommitProviderImpl_V2) Start(ctx context.Context) error {
	err := E.s.Start(ctx)
	if err != nil {
		return err
	}

	var errMu sync.Mutex
	var wg sync.WaitGroup
	// Replay in parallel if both requested.
	if E.sourceStartBlock != 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s := time.Now()
			E.lggr.Infow("start replaying src chain", "fromBlock", E.sourceStartBlock)
			srcReplayErr := E.sourceLP.Replay(ctx, int64(E.sourceStartBlock))
			errMu.Lock()
			err = multierr.Combine(err, srcReplayErr)
			errMu.Unlock()
			E.lggr.Infow("finished replaying src chain", "time", time.Since(s))
		}()
	}
	if E.destStartBlock != 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s := time.Now()
			E.lggr.Infow("start replaying dst chain", "fromBlock", E.destStartBlock)
			dstReplayErr := E.destLP.Replay(ctx, int64(E.destStartBlock))
			errMu.Lock()
			err = multierr.Combine(err, dstReplayErr)
			errMu.Unlock()
			E.lggr.Infow("finished replaying dst chain", "time", time.Since(s))
		}()
	}
	wg.Wait()
	if err != nil {
		E.lggr.Criticalw("unexpected error replaying, continuing plugin boot without all the logs backfilled", "err", err)
	}
	if err := ctx.Err(); err != nil {
		E.lggr.Errorw("context already cancelled", "err", err)
		return err
	}
	return nil
}

func (E EVMCCIPCommitProviderImpl_V2) Close() error {
	return E.s.Close()
}

func (E EVMCCIPCommitProviderImpl_V2) Ready() error {
	return E.s.Ready()
}

func (E EVMCCIPCommitProviderImpl_V2) HealthReport() map[string]error {
	report := map[string]error{}
	services.CopyHealth(report, E.cp.HealthReport())
	services.CopyHealth(report, E.s.HealthReport())
	return report
}

func (E EVMCCIPCommitProviderImpl_V2) OffchainConfigDigester() ocrtypes.OffchainConfigDigester {
	return E.cp.OffchainConfigDigester()
}

func (E EVMCCIPCommitProviderImpl_V2) ContractConfigTracker() ocrtypes.ContractConfigTracker {
	return E.cp.ContractConfigTracker()
}

func (E EVMCCIPCommitProviderImpl_V2) ContractTransmitter() ocrtypes.ContractTransmitter {
	panic("not implemented")
}

func (E EVMCCIPCommitProviderImpl_V2) ChainReader() types.ChainReader {
	return E.sourceChainReader
}

func (E EVMCCIPCommitProviderImpl_V2) Codec() types.Codec {
	return E.sourceCodec
}

func (E EVMCCIPCommitProviderImpl_V2) NewCommitStoreReader(ctx context.Context, _ cciptypes.Address) (commitStoreReader cciptypes.CommitStoreReader, err error) {
	commitStoreReader, err = ccip.NewCommitStoreReader(E.lggr, E.versionFinder, cciptypes.Address(E.commitStoreAddress), E.destClient, E.destLP, E.sourceGasEstimator, &E.sourceMaxGasPrice, nil)
	return
}

func (E EVMCCIPCommitProviderImpl_V2) NewOffRampReader(ctx context.Context, offRampAddr cciptypes.Address) (offRampReader cciptypes.OffRampReader, err error) {
	offRampReader, err = ccip.NewOffRampReader(E.lggr, E.versionFinder, offRampAddr, E.destClient, E.destLP, E.destGasEstimator, &E.destMaxGasPrice, true, nil)
	return
}

func (E EVMCCIPCommitProviderImpl_V2) NewOnRampReader(ctx context.Context, onRampAddress cciptypes.Address, sourceChainSelector uint64, destChainSelector uint64) (onRampReader cciptypes.OnRampReader, err error) {
	versionFinder := ccip.NewEvmVersionFinder()
	onRampReader, err = ccip.NewOnRampReader(E.lggr, versionFinder, sourceChainSelector, destChainSelector, onRampAddress, E.sourceLP, E.sourceClient, nil)
	return
}

// Dynamic Price Getter for CCIP commit service
func (E EVMCCIPCommitProviderImpl_V2) NewPriceGetter(ctx context.Context) (priceGetter cciptypes.PriceGetter, err error) {
	priceGetter, err = ccip.NewDynamicPriceGetter(E.priceGetterConfig, E.priceGetterClients)
	return
}

func (E EVMCCIPCommitProviderImpl_V2) NewPriceRegistryReader(ctx context.Context, addr cciptypes.Address) (priceRegistryReader cciptypes.PriceRegistryReader, err error) {
	destPriceRegistry := ccip.NewEvmPriceRegistry(E.destLP, E.destClient, E.lggr, ccip.CommitPluginLabel)
	priceRegistryReader, err = destPriceRegistry.NewPriceRegistryReader(ctx, addr)
	return
}

func (E EVMCCIPCommitProviderImpl_V2) SourceNativeToken(ctx context.Context, sourceRouterAddr cciptypes.Address) (cciptypes.Address, error) {
	sourceRouterAddrHex := sourceRouterAddr.ToCommonAddress()
	sourceRouter, err := router.NewRouter(sourceRouterAddrHex, E.sourceClient)
	if err != nil {
		return "", err
	}
	sourceNative, err := sourceRouter.GetWrappedNative(&bind.CallOpts{Context: ctx})
	if err != nil {
		return "", err
	}

	return cciptypes.FromCommonAddress(sourceNative), nil
}

func (E EVMCCIPCommitProviderImpl_V2) NewOffRampReaders(ctx context.Context, destRouterAddr cciptypes.Address) (offRampReaders []cciptypes.OffRampReader, err error) {
	// Look up all destination offRamps connected to the same router
	destRouterEvmAddr, err := ccip.GenericAddrToEvm(destRouterAddr)
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
	var destOffRampReaders []ccip.OffRampReader
	for _, o := range destRouterOffRamps {
		destOffRampAddr := cciptypes.Address(o.OffRamp.String())
		destOffRampReader, err2 := ccip.NewOffRampReader(
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
