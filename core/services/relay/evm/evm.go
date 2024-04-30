package evm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/ccipcommit"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/x_internal/cache"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/x_internal/ccipcalc"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/x_internal/ccipcommon"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/x_internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/x_internal/ccipdata/ccipdataprovider"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/x_internal/ccipdata/factory"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/x_internal/observability"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/x_internal/pricegetter"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/x_internal/rpclib"
	"github.com/smartcontractkit/chainlink/v2/core/services/pipeline"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	pkgerrors "github.com/pkg/errors"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/libocr/gethwrappers2/ocr2aggregator"
	"github.com/smartcontractkit/libocr/offchainreporting2/reportingplugin/median"
	"github.com/smartcontractkit/libocr/offchainreporting2/reportingplugin/median/evmreportcodec"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/services"
	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"

	txmgrcommon "github.com/smartcontractkit/chainlink/v2/common/txmgr"
	txm "github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/chains/legacyevm"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore"
	"github.com/smartcontractkit/chainlink/v2/core/services/llo"
	"github.com/smartcontractkit/chainlink/v2/core/services/llo/bm"
	lloconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/llo/config"
	mercuryconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/mercury/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/functions"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/mercury"
	mercuryutils "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/mercury/utils"
	reportcodecv1 "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/mercury/v1/reportcodec"
	reportcodecv2 "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/mercury/v2/reportcodec"
	reportcodecv3 "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/mercury/v3/reportcodec"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/mercury/wsrpc"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
)

var (
	OCR2AggregatorTransmissionContractABI abi.ABI
	OCR2AggregatorLogDecoder              LogDecoder
	ChannelVerifierLogDecoder             LogDecoder
)

func init() {
	var err error
	OCR2AggregatorTransmissionContractABI, err = abi.JSON(strings.NewReader(ocr2aggregator.OCR2AggregatorMetaData.ABI))
	if err != nil {
		panic(err)
	}
	OCR2AggregatorLogDecoder, err = newOCR2AggregatorLogDecoder()
	if err != nil {
		panic(err)
	}
	ChannelVerifierLogDecoder, err = newChannelVerifierLogDecoder()
	if err != nil {
		panic(err)
	}
}

var _ commontypes.Relayer = &Relayer{} //nolint:staticcheck

type Relayer struct {
	db          *sqlx.DB
	chain       legacyevm.Chain
	lggr        logger.Logger
	ks          CSAETHKeystore
	mercuryPool wsrpc.Pool
	pgCfg       pg.QConfig
	chainReader commontypes.ChainReader
	codec       commontypes.Codec

	// Mercury
	mercuryORM mercury.ORM

	// LLO/data streams
	cdcFactory llo.ChannelDefinitionCacheFactory
	lloORM     llo.ORM
}

type CSAETHKeystore interface {
	CSA() keystore.CSA
	Eth() keystore.Eth
}

type RelayerOpts struct {
	*sqlx.DB
	pg.QConfig
	CSAETHKeystore
	MercuryPool wsrpc.Pool
}

func (c RelayerOpts) Validate() error {
	var err error
	if c.DB == nil {
		err = errors.Join(err, errors.New("nil DB"))
	}
	if c.QConfig == nil {
		err = errors.Join(err, errors.New("nil QConfig"))
	}
	if c.CSAETHKeystore == nil {
		err = errors.Join(err, errors.New("nil Keystore"))
	}

	if err != nil {
		err = fmt.Errorf("invalid RelayerOpts: %w", err)
	}
	return err
}

func NewRelayer(lggr logger.Logger, chain legacyevm.Chain, opts RelayerOpts) (*Relayer, error) {
	err := opts.Validate()
	if err != nil {
		return nil, fmt.Errorf("cannot create evm relayer: %w", err)
	}
	lggr = lggr.Named("Relayer")

	mercuryORM := mercury.NewORM(opts.DB, lggr, opts.QConfig)
	lloORM := llo.NewORM(pg.NewQ(opts.DB, lggr, opts.QConfig), chain.ID())
	cdcFactory := llo.NewChannelDefinitionCacheFactory(lggr, lloORM, chain.LogPoller())
	return &Relayer{
		db:          opts.DB,
		chain:       chain,
		lggr:        lggr,
		ks:          opts.CSAETHKeystore,
		mercuryPool: opts.MercuryPool,
		pgCfg:       opts.QConfig,
		cdcFactory:  cdcFactory,
		lloORM:      lloORM,
		mercuryORM:  mercuryORM,
	}, nil
}

func (r *Relayer) Name() string {
	return r.lggr.Name()
}

// Start does noop: no subservices started on relay start, but when the first job is started
func (r *Relayer) Start(context.Context) error {
	return nil
}

func (r *Relayer) Close() error {
	return nil
}

// Ready does noop: always ready
func (r *Relayer) Ready() error {
	return r.chain.Ready()
}

func (r *Relayer) HealthReport() (report map[string]error) {
	report = make(map[string]error)
	maps.Copy(report, r.chain.HealthReport())
	return
}

func (r *Relayer) NewPluginProvider(rargs commontypes.RelayArgs, pargs commontypes.PluginArgs) (commontypes.PluginProvider, error) {

	// TODO https://smartcontract-it.atlassian.net/browse/BCF-2887
	ctx := context.Background()

	lggr := r.lggr.Named("PluginProvider").Named(rargs.ExternalJobID.String())

	configWatcher, err := newStandardConfigProvider(ctx, r.lggr, r.chain, types.NewRelayOpts(rargs))
	if err != nil {
		return nil, err
	}

	transmitter, err := newOnChainContractTransmitter(ctx, r.lggr, rargs, pargs.TransmitterID, r.ks.Eth(), configWatcher, configTransmitterOpts{}, OCR2AggregatorTransmissionContractABI, nil)
	if err != nil {
		return nil, err
	}

	return NewPluginProvider(
		r.chainReader,
		r.codec,
		transmitter,
		configWatcher,
		lggr,
	), nil
}

func (r *Relayer) NewMercuryProvider(rargs commontypes.RelayArgs, pargs commontypes.PluginArgs) (commontypes.MercuryProvider, error) {
	// TODO https://smartcontract-it.atlassian.net/browse/BCF-2887
	ctx := context.Background()
	lggr := r.lggr.Named("MercuryProvider").Named(rargs.ExternalJobID.String())
	relayOpts := types.NewRelayOpts(rargs)
	relayConfig, err := relayOpts.RelayConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get relay config: %w", err)
	}

	var mercuryConfig mercuryconfig.PluginConfig
	if err = json.Unmarshal(pargs.PluginConfig, &mercuryConfig); err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	if relayConfig.FeedID == nil {
		return nil, pkgerrors.New("FeedID must be specified")
	}
	feedID := mercuryutils.FeedID(*relayConfig.FeedID)

	if relayConfig.ChainID.String() != r.chain.ID().String() {
		return nil, fmt.Errorf("internal error: chain id in spec does not match this relayer's chain: have %s expected %s", relayConfig.ChainID.String(), r.chain.ID().String())
	}
	cp, err := newMercuryConfigProvider(ctx, lggr, r.chain, relayOpts)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	if !relayConfig.EffectiveTransmitterID.Valid {
		return nil, pkgerrors.New("EffectiveTransmitterID must be specified")
	}
	privKey, err := r.ks.CSA().Get(relayConfig.EffectiveTransmitterID.String)
	if err != nil {
		return nil, pkgerrors.Wrap(err, "failed to get CSA key for mercury connection")
	}

	clients := make(map[string]wsrpc.Client)
	for _, server := range mercuryConfig.GetServers() {
		client, err := r.mercuryPool.Checkout(context.Background(), privKey, server.PubKey, server.URL)
		if err != nil {
			return nil, err
		}
		clients[server.URL] = client
	}

	// FIXME: We actually know the version here since it's in the feed ID, can
	// we use generics to avoid passing three of this?
	// https://smartcontract-it.atlassian.net/browse/MERC-1414
	reportCodecV1 := reportcodecv1.NewReportCodec(*relayConfig.FeedID, lggr.Named("ReportCodecV1"))
	reportCodecV2 := reportcodecv2.NewReportCodec(*relayConfig.FeedID, lggr.Named("ReportCodecV2"))
	reportCodecV3 := reportcodecv3.NewReportCodec(*relayConfig.FeedID, lggr.Named("ReportCodecV3"))

	var transmitterCodec mercury.TransmitterReportDecoder
	switch feedID.Version() {
	case 1:
		transmitterCodec = reportCodecV1
	case 2:
		transmitterCodec = reportCodecV2
	case 3:
		transmitterCodec = reportCodecV3
	default:
		return nil, fmt.Errorf("invalid feed version %d", feedID.Version())
	}
	transmitter := mercury.NewTransmitter(lggr, clients, privKey.PublicKey, rargs.JobID, *relayConfig.FeedID, r.mercuryORM, transmitterCodec)

	return NewMercuryProvider(cp, r.chainReader, r.codec, NewMercuryChainReader(r.chain.HeadTracker()), transmitter, reportCodecV1, reportCodecV2, reportCodecV3, lggr), nil
}

func (r *Relayer) NewLLOProvider(rargs commontypes.RelayArgs, pargs commontypes.PluginArgs) (commontypes.LLOProvider, error) {

	// TODO https://smartcontract-it.atlassian.net/browse/BCF-2887
	ctx := context.Background()

	relayOpts := types.NewRelayOpts(rargs)
	var relayConfig types.RelayConfig
	{
		var err error
		relayConfig, err = relayOpts.RelayConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to get relay config: %w", err)
		}
	}

	var lloCfg lloconfig.PluginConfig
	if err := json.Unmarshal(pargs.PluginConfig, &lloCfg); err != nil {
		return nil, pkgerrors.WithStack(err)
	}
	if err := lloCfg.Validate(); err != nil {
		return nil, err
	}

	if relayConfig.ChainID.String() != r.chain.ID().String() {
		return nil, fmt.Errorf("internal error: chain id in spec does not match this relayer's chain: have %s expected %s", relayConfig.ChainID.String(), r.chain.ID().String())
	}
	cp, err := newLLOConfigProvider(ctx, r.lggr, r.chain, relayOpts)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	if !relayConfig.EffectiveTransmitterID.Valid {
		return nil, pkgerrors.New("EffectiveTransmitterID must be specified")
	}
	privKey, err := r.ks.CSA().Get(relayConfig.EffectiveTransmitterID.String)
	if err != nil {
		return nil, pkgerrors.Wrap(err, "failed to get CSA key for mercury connection")
	}

	// FIXME: Remove after benchmarking is done
	// https://smartcontract-it.atlassian.net/browse/MERC-3487
	var transmitter llo.Transmitter
	if lloCfg.BenchmarkMode {
		r.lggr.Info("Benchmark mode enabled, using dummy transmitter. NOTE: THIS WILL NOT TRANSMIT ANYTHING")
		transmitter = bm.NewTransmitter(r.lggr, privKey.PublicKey)
	} else {
		var client wsrpc.Client
		client, err = r.mercuryPool.Checkout(context.Background(), privKey, lloCfg.ServerPubKey, lloCfg.ServerURL())
		if err != nil {
			return nil, err
		}
		transmitter = llo.NewTransmitter(r.lggr, client, privKey.PublicKey)
	}

	cdc, err := r.cdcFactory.NewCache(lloCfg)
	if err != nil {
		return nil, err
	}
	return NewLLOProvider(cp, transmitter, r.lggr, cdc), nil
}

func (r *Relayer) NewFunctionsProvider(rargs commontypes.RelayArgs, pargs commontypes.PluginArgs) (commontypes.FunctionsProvider, error) {

	// TODO https://smartcontract-it.atlassian.net/browse/BCF-2887
	ctx := context.Background()

	lggr := r.lggr.Named("FunctionsProvider").Named(rargs.ExternalJobID.String())
	// TODO(FUN-668): Not ready yet (doesn't implement FunctionsEvents() properly)
	return NewFunctionsProvider(ctx, r.chain, rargs, pargs, lggr, r.ks.Eth(), functions.FunctionsPlugin)
}

// huiepatr TODO: chainSet likely means we need two relayers, one for source chain and one for dest chain
// pipeline runner needs to go
// jobname needs to go
// qopts need to go (should replace with context anyways)
func (r *Relayer) NewCCIPCommitProvider(rargs commontypes.RelayArgs, pargs commontypes.PluginArgs, chainSet legacyevm.LegacyChainContainer, pr pipeline.Runner, jobName string, qopts ...pg.QOpt) (commontypes.CCIPCommitProvider, error) {
	// TODO https://smartcontract-it.atlassian.net/browse/BCF-2887
	ctx := context.Background()

	// huiepatr TODO: this validation will probably happen before we convert job/spec to rargs + pargs
	//if jb.OCR2OracleSpec == nil {
	//	return nil, errors.New("spec is nil")
	//}
	//spec := jb.OCR2OracleSpec

	var pluginConfig ccipconfig.CommitPluginJobSpecConfig
	err := json.Unmarshal(pargs.PluginConfig, &pluginConfig)
	if err != nil {
		return nil, err
	}
	// ensure addresses are formatted properly - (lowercase to eip55 for evm)
	pluginConfig.OffRamp = cciptypes.Address(common.HexToAddress(string(pluginConfig.OffRamp)).String())

	destChain := r.chain

	gethCommitStoreAddress := common.HexToAddress(rargs.ContractID) // TODO: is this the right contract ID?

	// huiepatr TODO: figure out how we get FetchCommitStoreStaticConfig in here
	staticConfig, err := ccipdata.FetchCommitStoreStaticConfig(gethCommitStoreAddress, destChain.Client())
	if err != nil {
		return nil, fmt.Errorf("get commit store static config: %w", err)
	}

	sourceChain, _, err := ccipconfig.GetChainByChainSelector(chainSet, staticConfig.SourceChainSelector)
	if err != nil {
		return nil, err
	}

	// return value of extractJobSpecParams
	//return &jobSpecParams{
	//	pluginConfig:         pluginConfig,
	//	commitStoreAddress:   cciptypes.Address(commitStoreAddress.String()),
	//	commitStoreStaticCfg: staticConfig,
	//	sourceChain:          sourceChain,
	//	destChain:            destChain,
	//}, nil

	commitStoreAddress := cciptypes.Address(gethCommitStoreAddress.String())
	commitStoreStaticCfg := staticConfig

	lggr := r.lggr

	lggr.Infow("Initializing commit plugin",
		"CommitStore", commitStoreAddress,
		"OffRamp", pluginConfig.OffRamp,
		"OnRamp", commitStoreStaticCfg.OnRamp,
		"ArmProxy", commitStoreStaticCfg.ArmProxy,
		"SourceChainSelector", commitStoreStaticCfg.SourceChainSelector,
		"DestChainSelector", commitStoreStaticCfg.ChainSelector)

	versionFinder := factory.NewEvmVersionFinder()
	commitStoreReader, err := factory.NewCommitStoreReader(lggr, versionFinder, commitStoreAddress, destChain.Client(), destChain.LogPoller(), sourceChain.GasEstimator(), sourceChain.Config().EVM().GasEstimator().PriceMax().ToInt(), qopts...)
	if err != nil {
		return nil, fmt.Errorf("could not create commitStore reader: %w", err)
	}
	sourceChainName, destChainName, err := ccipconfig.ResolveChainNames(sourceChain.ID().Int64(), destChain.ID().Int64())
	if err != nil {
		return nil, err
	}
	commitLggr := r.lggr.Named("CCIPCommit").With("sourceChain", sourceChainName, "destChain", destChainName)

	var priceGetter pricegetter.PriceGetter
	withPipeline := strings.Trim(pluginConfig.TokenPricesUSDPipeline, "\n\t ") != ""
	if withPipeline {
		priceGetter, err = pricegetter.NewPipelineGetter(pluginConfig.TokenPricesUSDPipeline, pr, rargs.JobID, rargs.ExternalJobID, jobName, r.lggr)
		if err != nil {
			return nil, fmt.Errorf("creating pipeline price getter: %w", err)
		}
	} else {
		// Use dynamic price getter.
		if pluginConfig.PriceGetterConfig == nil {
			return nil, fmt.Errorf("priceGetterConfig is nil")
		}

		// Build price getter clients for all chains specified in the aggregator configurations.
		// Some lanes (e.g. Wemix/Kroma) requires other clients than source and destination, since they use feeds from other chains.
		priceGetterClients := map[uint64]pricegetter.DynamicPriceGetterClient{}
		for _, aggCfg := range pluginConfig.PriceGetterConfig.AggregatorPrices {
			chainID := aggCfg.ChainID
			// Retrieve the chain.
			chain, _, err2 := ccipconfig.GetChainByChainID(chainSet, chainID)
			if err2 != nil {
				return nil, fmt.Errorf("retrieving chain for chainID %d: %w", chainID, err2)
			}
			caller := rpclib.NewDynamicLimitedBatchCaller(
				lggr,
				chain.Client(),
				rpclib.DefaultRpcBatchSizeLimit,
				rpclib.DefaultRpcBatchBackOffMultiplier,
				rpclib.DefaultMaxParallelRpcCalls,
			)
			priceGetterClients[chainID] = pricegetter.NewDynamicPriceGetterClient(caller)
		}

		priceGetter, err = pricegetter.NewDynamicPriceGetter(*pluginConfig.PriceGetterConfig, priceGetterClients)
		if err != nil {
			return nil, fmt.Errorf("creating dynamic price getter: %w", err)
		}
	}

	// Load all the readers relevant for this plugin.
	onrampAddress := cciptypes.Address(commitStoreStaticCfg.OnRamp.String())
	onRampReader, err := factory.NewOnRampReader(commitLggr, versionFinder, commitStoreStaticCfg.SourceChainSelector, commitStoreStaticCfg.ChainSelector, onrampAddress, sourceChain.LogPoller(), sourceChain.Client(), qopts...)
	if err != nil {
		return nil, pkgerrors.Wrap(err, "failed onramp reader")
	}
	offRampReader, err := factory.NewOffRampReader(commitLggr, versionFinder, pluginConfig.OffRamp, destChain.Client(), destChain.LogPoller(), destChain.GasEstimator(), destChain.Config().EVM().GasEstimator().PriceMax().ToInt(), true, qopts...)
	if err != nil {
		return nil, pkgerrors.Wrap(err, "failed offramp reader")
	}
	// Look up all destination offRamps connected to the same router
	destRouterAddr, err := offRampReader.GetRouter(ctx)
	if err != nil {
		return nil, err
	}
	destRouterEvmAddr, err := ccipcalc.GenericAddrToEvm(destRouterAddr)
	if err != nil {
		return nil, err
	}
	destRouter, err := router.NewRouter(destRouterEvmAddr, destChain.Client())
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
			commitLggr,
			versionFinder,
			destOffRampAddr,
			destChain.Client(),
			destChain.LogPoller(),
			destChain.GasEstimator(),
			destChain.Config().EVM().GasEstimator().PriceMax().ToInt(),
			true,
			// qopts..., huiepatr TODO
		)
		if err2 != nil {
			return nil, err2
		}

		destOffRampReaders = append(destOffRampReaders, destOffRampReader)
	}

	onRampRouterAddr, err := onRampReader.RouterAddress(ctx)
	if err != nil {
		return nil, err
	}
	sourceRouterAddr, err := ccipcalc.GenericAddrToEvm(onRampRouterAddr)
	if err != nil {
		return nil, err
	}
	sourceRouter, err := router.NewRouter(sourceRouterAddr, sourceChain.Client())
	if err != nil {
		return nil, err
	}
	sourceNative, err := sourceRouter.GetWrappedNative(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	// Prom wrappers
	onRampReader = observability.NewObservedOnRampReader(onRampReader, sourceChain.ID().Int64(), ccip.CommitPluginLabel)
	commitStoreReader = observability.NewObservedCommitStoreReader(commitStoreReader, destChain.ID().Int64(), ccip.CommitPluginLabel)
	metricsCollector := ccip.NewPluginMetricsCollector(ccip.CommitPluginLabel, sourceChain.ID().Int64(), destChain.ID().Int64())
	for i, o := range destOffRampReaders {
		destOffRampReaders[i] = observability.NewObservedOffRampReader(o, destChain.ID().Int64(), ccip.CommitPluginLabel)
	}

	chainHealthcheck := cache.NewObservedChainHealthCheck(
		cache.NewChainHealthcheck(
			// Adding more details to Logger to make healthcheck logs more informative
			// It's safe because healthcheck logs only in case of unhealthy state
			lggr.With(
				"onramp", onrampAddress,
				"commitStore", commitStoreAddress,
				"offramp", pluginConfig.OffRamp,
			),
			onRampReader,
			commitStoreReader,
		),
		ccip.CommitPluginLabel,
		sourceChain.ID().Int64(),
		destChain.ID().Int64(),
		onrampAddress,
	)

	commitLggr.Infow("NewCommitServices",
		"pluginConfig", pluginConfig,
		"staticConfig", commitStoreStaticCfg,
		// TODO bring back
		//"dynamicOnRampConfig", dynamicOnRampConfig,
		"sourceNative", sourceNative,
		"sourceRouter", sourceRouter.Address())

	// jobSpecToCommitPluginConfig
	//return &CommitPluginStaticConfig{
	//		lggr:                  commitLggr,
	//		onRampReader:          onRampReader,
	//		offRamps:              destOffRampReaders,
	//		sourceNative:          ccipcalc.EvmAddrToGeneric(sourceNative),
	//		priceGetter:           priceGetter,
	//		sourceChainSelector:   commitStoreStaticCfg.SourceChainSelector,
	//		destChainSelector:     commitStoreStaticCfg.ChainSelector,
	//		commitStore:           commitStoreReader,
	//		priceRegistryProvider: ccipdataprovider.NewEvmPriceRegistry(destChain.LogPoller(), destChain.Client(), commitLggr, ccip.CommitPluginLabel),
	//		metricsCollector:      metricsCollector,
	//		chainHealthcheck:      chainHealthcheck,
	//	}, &ccipcommon.BackfillArgs{
	//		SourceLP:         sourceChain.LogPoller(),
	//		DestLP:           destChain.LogPoller(),
	//		SourceStartBlock: pluginConfig.SourceStartBlock,
	//		DestStartBlock:   pluginConfig.DestStartBlock,
	//	},
	//	chainHealthcheck,
	//	nil

	ccipPluginConfig := ccipcommit.CommitPluginStaticConfig{
		Lggr:                  commitLggr,
		OnRampReader:          onRampReader,
		OffRamps:              destOffRampReaders,
		SourceNative:          ccipcalc.EvmAddrToGeneric(sourceNative),
		PriceGetter:           priceGetter,
		SourceChainSelector:   commitStoreStaticCfg.SourceChainSelector,
		DestChainSelector:     commitStoreStaticCfg.ChainSelector,
		CommitStore:           commitStoreReader,
		PriceRegistryProvider: ccipdataprovider.NewEvmPriceRegistry(destChain.LogPoller(), destChain.Client(), commitLggr, ccip.CommitPluginLabel),
		MetricsCollector:      metricsCollector,
		ChainHealthcheck:      chainHealthcheck,
	}

	backfillArgs := ccipcommon.BackfillArgs{
		SourceLP:         sourceChain.LogPoller(),
		DestLP:           destChain.LogPoller(),
		SourceStartBlock: pluginConfig.SourceStartBlock,
		DestStartBlock:   pluginConfig.DestStartBlock,
	}

	return XXXCreateEVMCCIPCommitProvider(ccipPluginConfig, backfillArgs, chainHealthcheck), nil

}

type EVMCCIPCommitProviderImpl struct{}

func XXXCreateEVMCCIPCommitProvider(cpsc ccipcommit.CommitPluginStaticConfig, backfillArgs ccipcommon.BackfillArgs, chainHealthcheck *cache.ObservedChainHealthcheck) EVMCCIPCommitProviderImpl {
	// huiepatr TODO: do something sensible
	return EVMCCIPCommitProviderImpl{}
}

func (E EVMCCIPCommitProviderImpl) Name() string {
	return "EVMCCIPCommitProvider"
}

func (E EVMCCIPCommitProviderImpl) Start(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl) Close() error {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl) Ready() error {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl) HealthReport() map[string]error {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl) OffchainConfigDigester() ocrtypes.OffchainConfigDigester {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl) ContractConfigTracker() ocrtypes.ContractConfigTracker {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl) ContractTransmitter() ocrtypes.ContractTransmitter {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl) ChainReader() commontypes.ChainReader {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl) Codec() commontypes.Codec {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl) NewCommitStoreReader(ctx context.Context, addr cciptypes.Address) (cciptypes.CommitStoreReader, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl) NewOffRampReader(ctx context.Context, addr cciptypes.Address) (cciptypes.OffRampReader, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl) NewOnRampReader(ctx context.Context, addr cciptypes.Address) (cciptypes.OnRampReader, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl) NewPriceGetter(ctx context.Context) (cciptypes.PriceGetter, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl) NewPriceRegistryReader(ctx context.Context, addr cciptypes.Address) (cciptypes.PriceRegistryReader, error) {
	//TODO implement me
	panic("implement me")
}

func (E EVMCCIPCommitProviderImpl) SourceNativeToken(ctx context.Context) (cciptypes.Address, error) {
	//TODO implement me
	panic("implement me")
}

type CommitStoreStaticConfig struct {
	ChainSelector       uint64
	SourceChainSelector uint64
	OnRamp              common.Address
	ArmProxy            common.Address
}

// NewConfigProvider is called by bootstrap jobs
func (r *Relayer) NewConfigProvider(args commontypes.RelayArgs) (configProvider commontypes.ConfigProvider, err error) {
	// TODO https://smartcontract-it.atlassian.net/browse/BCF-2887
	ctx := context.Background()

	lggr := r.lggr.Named("ConfigProvider").Named(args.ExternalJobID.String())
	relayOpts := types.NewRelayOpts(args)
	relayConfig, err := relayOpts.RelayConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get relay config: %w", err)
	}
	expectedChainID := relayConfig.ChainID.String()
	if expectedChainID != r.chain.ID().String() {
		return nil, fmt.Errorf("internal error: chain id in spec does not match this relayer's chain: have %s expected %s", relayConfig.ChainID.String(), r.chain.ID().String())
	}

	// Handle legacy jobs which did not yet specify provider type and
	// switched between median/mercury based on presence of feed ID
	if args.ProviderType == "" {
		if relayConfig.FeedID == nil {
			args.ProviderType = "median"
		} else {
			args.ProviderType = "mercury"
		}
	}

	switch args.ProviderType {
	case "median":
		configProvider, err = newStandardConfigProvider(ctx, lggr, r.chain, relayOpts)
	case "mercury":
		configProvider, err = newMercuryConfigProvider(ctx, lggr, r.chain, relayOpts)
	case "llo":
		configProvider, err = newLLOConfigProvider(ctx, lggr, r.chain, relayOpts)
	default:
		return nil, fmt.Errorf("unrecognized provider type: %q", args.ProviderType)
	}

	if err != nil {
		// Never return (*configProvider)(nil)
		return nil, err
	}
	return configProvider, err
}

func FilterNamesFromRelayArgs(args commontypes.RelayArgs) (filterNames []string, err error) {
	var addr evmtypes.EIP55Address
	if addr, err = evmtypes.NewEIP55Address(args.ContractID); err != nil {
		return nil, err
	}
	var relayConfig types.RelayConfig
	if err = json.Unmarshal(args.RelayConfig, &relayConfig); err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	if relayConfig.FeedID != nil {
		filterNames = []string{mercury.FilterName(addr.Address(), *relayConfig.FeedID)}
	} else {
		filterNames = []string{configPollerFilterName(addr.Address()), transmitterFilterName(addr.Address())}
	}
	return filterNames, err
}

type configWatcher struct {
	services.StateMachine
	lggr             logger.Logger
	contractAddress  common.Address
	offchainDigester ocrtypes.OffchainConfigDigester
	configPoller     types.ConfigPoller
	chain            legacyevm.Chain
	runReplay        bool
	fromBlock        uint64
	replayCtx        context.Context
	replayCancel     context.CancelFunc
	wg               sync.WaitGroup
}

func newConfigWatcher(lggr logger.Logger,
	contractAddress common.Address,
	offchainDigester ocrtypes.OffchainConfigDigester,
	configPoller types.ConfigPoller,
	chain legacyevm.Chain,
	fromBlock uint64,
	runReplay bool,
) *configWatcher {
	replayCtx, replayCancel := context.WithCancel(context.Background())
	return &configWatcher{
		lggr:             lggr.Named("ConfigWatcher").Named(contractAddress.String()),
		contractAddress:  contractAddress,
		offchainDigester: offchainDigester,
		configPoller:     configPoller,
		chain:            chain,
		runReplay:        runReplay,
		fromBlock:        fromBlock,
		replayCtx:        replayCtx,
		replayCancel:     replayCancel,
	}

}

func (c *configWatcher) Name() string {
	return c.lggr.Name()
}

func (c *configWatcher) Start(ctx context.Context) error {
	return c.StartOnce(fmt.Sprintf("configWatcher %x", c.contractAddress), func() error {
		if c.runReplay && c.fromBlock != 0 {
			// Only replay if it's a brand runReplay job.
			c.wg.Add(1)
			go func() {
				defer c.wg.Done()
				c.lggr.Infow("starting replay for config", "fromBlock", c.fromBlock)
				if err := c.configPoller.Replay(c.replayCtx, int64(c.fromBlock)); err != nil {
					c.lggr.Errorf("error replaying for config", "err", err)
				} else {
					c.lggr.Infow("completed replaying for config", "fromBlock", c.fromBlock)
				}
			}()
		}
		c.configPoller.Start()
		return nil
	})
}

func (c *configWatcher) Close() error {
	return c.StopOnce(fmt.Sprintf("configWatcher %x", c.contractAddress), func() error {
		c.replayCancel()
		c.wg.Wait()
		return c.configPoller.Close()
	})
}

func (c *configWatcher) HealthReport() map[string]error {
	return map[string]error{c.Name(): c.Healthy()}
}

func (c *configWatcher) OffchainConfigDigester() ocrtypes.OffchainConfigDigester {
	return c.offchainDigester
}

func (c *configWatcher) ContractConfigTracker() ocrtypes.ContractConfigTracker {
	return c.configPoller
}

type configTransmitterOpts struct {
	// pluginGasLimit overrides the gas limit default provided in the config watcher.
	pluginGasLimit *uint32
	// subjectID overrides the queueing subject id (the job external id will be used by default).
	subjectID *uuid.UUID
}

// newOnChainContractTransmitter creates a new contract transmitter.
func newOnChainContractTransmitter(ctx context.Context, lggr logger.Logger, rargs commontypes.RelayArgs, transmitterID string, ethKeystore keystore.Eth, configWatcher *configWatcher, opts configTransmitterOpts, transmissionContractABI abi.ABI, reportToEthMeta ReportToEthMetadata) (*contractTransmitter, error) {
	var relayConfig types.RelayConfig
	if err := json.Unmarshal(rargs.RelayConfig, &relayConfig); err != nil {
		return nil, err
	}
	var fromAddresses []common.Address
	sendingKeys := relayConfig.SendingKeys
	if !relayConfig.EffectiveTransmitterID.Valid {
		return nil, pkgerrors.New("EffectiveTransmitterID must be specified")
	}
	effectiveTransmitterAddress := common.HexToAddress(relayConfig.EffectiveTransmitterID.String)

	sendingKeysLength := len(sendingKeys)
	if sendingKeysLength == 0 {
		return nil, pkgerrors.New("no sending keys provided")
	}

	// If we are using multiple sending keys, then a forwarder is needed to rotate transmissions.
	// Ensure that this forwarder is not set to a local sending key, and ensure our sending keys are enabled.
	for _, s := range sendingKeys {
		if sendingKeysLength > 1 && s == effectiveTransmitterAddress.String() {
			return nil, pkgerrors.New("the transmitter is a local sending key with transaction forwarding enabled")
		}
		if err := ethKeystore.CheckEnabled(ctx, common.HexToAddress(s), configWatcher.chain.Config().EVM().ChainID()); err != nil {
			return nil, pkgerrors.Wrap(err, "one of the sending keys given is not enabled")
		}
		fromAddresses = append(fromAddresses, common.HexToAddress(s))
	}

	subject := rargs.ExternalJobID
	if opts.subjectID != nil {
		subject = *opts.subjectID
	}
	scoped := configWatcher.chain.Config()
	strategy := txmgrcommon.NewQueueingTxStrategy(subject, scoped.OCR2().DefaultTransactionQueueDepth(), scoped.Database().DefaultQueryTimeout())

	var checker txm.TransmitCheckerSpec
	if configWatcher.chain.Config().OCR2().SimulateTransactions() {
		checker.CheckerType = txm.TransmitCheckerTypeSimulate
	}

	gasLimit := configWatcher.chain.Config().EVM().GasEstimator().LimitDefault()
	ocr2Limit := configWatcher.chain.Config().EVM().GasEstimator().LimitJobType().OCR2()
	if ocr2Limit != nil {
		gasLimit = uint64(*ocr2Limit)
	}
	if opts.pluginGasLimit != nil {
		gasLimit = uint64(*opts.pluginGasLimit)
	}

	transmitter, err := ocrcommon.NewTransmitter(
		configWatcher.chain.TxManager(),
		fromAddresses,
		gasLimit,
		effectiveTransmitterAddress,
		strategy,
		checker,
		configWatcher.chain.ID(),
		ethKeystore,
	)

	if err != nil {
		return nil, pkgerrors.Wrap(err, "failed to create transmitter")
	}

	return NewOCRContractTransmitter(
		ctx,
		configWatcher.contractAddress,
		configWatcher.chain.Client(),
		transmissionContractABI,
		transmitter,
		configWatcher.chain.LogPoller(),
		lggr,
		reportToEthMeta,
	)
}

func (r *Relayer) NewMedianProvider(rargs commontypes.RelayArgs, pargs commontypes.PluginArgs) (commontypes.MedianProvider, error) {
	// TODO https://smartcontract-it.atlassian.net/browse/BCF-2887
	ctx := context.Background()

	lggr := r.lggr.Named("MedianProvider").Named(rargs.ExternalJobID.String())
	relayOpts := types.NewRelayOpts(rargs)
	relayConfig, err := relayOpts.RelayConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get relay config: %w", err)
	}
	expectedChainID := relayConfig.ChainID.String()
	if expectedChainID != r.chain.ID().String() {
		return nil, fmt.Errorf("internal error: chain id in spec does not match this relayer's chain: have %s expected %s", relayConfig.ChainID.String(), r.chain.ID().String())
	}
	if !common.IsHexAddress(relayOpts.ContractID) {
		return nil, fmt.Errorf("invalid contractID %s, expected hex address", relayOpts.ContractID)
	}
	contractID := common.HexToAddress(relayOpts.ContractID)

	configWatcher, err := newStandardConfigProvider(ctx, lggr, r.chain, relayOpts)
	if err != nil {
		return nil, err
	}

	reportCodec := evmreportcodec.ReportCodec{}

	contractTransmitter, err := newOnChainContractTransmitter(ctx, lggr, rargs, pargs.TransmitterID, r.ks.Eth(), configWatcher, configTransmitterOpts{}, OCR2AggregatorTransmissionContractABI, nil)
	if err != nil {
		return nil, err
	}

	medianContract, err := newMedianContract(configWatcher.ContractConfigTracker(), configWatcher.contractAddress, configWatcher.chain, rargs.JobID, r.db, lggr)
	if err != nil {
		return nil, err
	}

	medianProvider := medianProvider{
		lggr:                lggr.Named("MedianProvider"),
		configWatcher:       configWatcher,
		reportCodec:         reportCodec,
		contractTransmitter: contractTransmitter,
		medianContract:      medianContract,
	}

	// allow fallback until chain reader is default and median contract is removed, but still log just in case
	var chainReaderService ChainReaderService
	if relayConfig.ChainReader != nil {
		if chainReaderService, err = NewChainReaderService(ctx, lggr, r.chain.LogPoller(), r.chain, *relayConfig.ChainReader); err != nil {
			return nil, err
		}

		boundContracts := []commontypes.BoundContract{{Name: "median", Pending: true, Address: contractID.String()}}
		if err = chainReaderService.Bind(context.Background(), boundContracts); err != nil {
			return nil, err
		}
	} else {
		lggr.Info("ChainReader missing from RelayConfig; falling back to internal MedianContract")
	}
	medianProvider.chainReader = chainReaderService

	if relayConfig.Codec != nil {
		medianProvider.codec, err = NewCodec(*relayConfig.Codec)
		if err != nil {
			return nil, err
		}
	} else {
		lggr.Info("Codec missing from RelayConfig; falling back to internal MedianContract")
	}

	return &medianProvider, nil
}

func (r *Relayer) NewAutomationProvider(rargs commontypes.RelayArgs, pargs commontypes.PluginArgs) (commontypes.AutomationProvider, error) {
	lggr := r.lggr.Named("AutomationProvider").Named(rargs.ExternalJobID.String())
	ocr2keeperRelayer := NewOCR2KeeperRelayer(r.db, r.chain, lggr.Named("OCR2KeeperRelayer"), r.ks.Eth(), r.pgCfg)

	return ocr2keeperRelayer.NewOCR2KeeperProvider(rargs, pargs)
}

var _ commontypes.MedianProvider = (*medianProvider)(nil)

type medianProvider struct {
	lggr                logger.Logger
	configWatcher       *configWatcher
	contractTransmitter ContractTransmitter
	reportCodec         median.ReportCodec
	medianContract      *medianContract
	chainReader         ChainReaderService
	codec               commontypes.Codec
	ms                  services.MultiStart
}

func (p *medianProvider) Name() string { return p.lggr.Name() }

func (p *medianProvider) Start(ctx context.Context) error {
	srvcs := []services.StartClose{p.configWatcher, p.contractTransmitter, p.medianContract}
	if p.chainReader != nil {
		srvcs = append(srvcs, p.chainReader)
	}

	return p.ms.Start(ctx, srvcs...)
}

func (p *medianProvider) Close() error { return p.ms.Close() }

func (p *medianProvider) Ready() error { return nil }

func (p *medianProvider) HealthReport() map[string]error {
	hp := map[string]error{p.Name(): p.Ready()}
	services.CopyHealth(hp, p.configWatcher.HealthReport())
	services.CopyHealth(hp, p.contractTransmitter.HealthReport())
	services.CopyHealth(hp, p.medianContract.HealthReport())
	return hp
}

func (p *medianProvider) ContractTransmitter() ocrtypes.ContractTransmitter {
	return p.contractTransmitter
}

func (p *medianProvider) ReportCodec() median.ReportCodec {
	return p.reportCodec
}

func (p *medianProvider) MedianContract() median.MedianContract {
	return p.medianContract
}

func (p *medianProvider) OnchainConfigCodec() median.OnchainConfigCodec {
	return median.StandardOnchainConfigCodec{}
}

func (p *medianProvider) OffchainConfigDigester() ocrtypes.OffchainConfigDigester {
	return p.configWatcher.OffchainConfigDigester()
}

func (p *medianProvider) ContractConfigTracker() ocrtypes.ContractConfigTracker {
	return p.configWatcher.ContractConfigTracker()
}

func (p *medianProvider) ChainReader() commontypes.ChainReader {
	return p.chainReader
}

func (p *medianProvider) Codec() commontypes.Codec {
	return p.codec
}
