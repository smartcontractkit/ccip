package oraclecreator

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	chainsel "github.com/smartcontractkit/chain-selectors"
	commitocr3 "github.com/smartcontractkit/chainlink-ccip/commit"
	ccipreaderpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/commontypes"
	libocr3 "github.com/smartcontractkit/libocr/offchainreporting2plus"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/loop"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/crconfigs"
	"github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/ocrimpls"
	cctypes "github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/types"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/chaintype"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/ccipevm"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/v2/core/services/synchronization"
	"github.com/smartcontractkit/chainlink/v2/core/services/telemetry"
)

var _ cctypes.OracleCreator = &inprocessOracleCreator{}

// inprocessOracleCreator creates oracles that reference plugins running
// in the same process as the chainlink node, i.e not LOOPPs.
type inprocessOracleCreator struct {
	ocrKeyBundles         map[chaintype.ChainType]ocr2key.KeyBundle
	transmitters          map[types.RelayID][]string
	relayers              map[types.RelayID]loop.Relayer
	peerWrapper           *ocrcommon.SingletonPeerWrapper
	externalJobID         uuid.UUID
	jobID                 int32
	isNewlyCreatedJob     bool
	pluginConfig          job.JSONConfig
	db                    ocr3types.Database
	lggr                  logger.Logger
	monitoringEndpointGen telemetry.MonitoringEndpointGenerator
	bootstrapperLocators  []commontypes.BootstrapperLocator
	homeChainReader       ccipreaderpkg.HomeChain
}

func New(
	ocrKeyBundles map[chaintype.ChainType]ocr2key.KeyBundle,
	transmitters map[types.RelayID][]string,
	relayers map[types.RelayID]loop.Relayer,
	peerWrapper *ocrcommon.SingletonPeerWrapper,
	externalJobID uuid.UUID,
	jobID int32,
	isNewlyCreatedJob bool,
	pluginConfig job.JSONConfig,
	db ocr3types.Database,
	lggr logger.Logger,
	monitoringEndpointGen telemetry.MonitoringEndpointGenerator,
	bootstrapperLocators []commontypes.BootstrapperLocator,
) cctypes.OracleCreator {
	return &inprocessOracleCreator{
		ocrKeyBundles:         ocrKeyBundles,
		transmitters:          transmitters,
		relayers:              relayers,
		peerWrapper:           peerWrapper,
		externalJobID:         externalJobID,
		jobID:                 jobID,
		isNewlyCreatedJob:     isNewlyCreatedJob,
		pluginConfig:          pluginConfig,
		db:                    db,
		lggr:                  lggr,
		monitoringEndpointGen: monitoringEndpointGen,
		bootstrapperLocators:  bootstrapperLocators,
	}
}

// CreateBootstrapOracle implements types.OracleCreator.
func (i *inprocessOracleCreator) CreateBootstrapOracle(config cctypes.OCR3ConfigWithMeta) (cctypes.CCIPOracle, error) {
	// Assuming that the chain selector is referring to an evm chain for now.
	// TODO: add an api that returns chain family.
	chainID, err := chainsel.ChainIdFromSelector(uint64(config.Config.ChainSelector))
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID from selector: %w", err)
	}

	destChainFamily := chaintype.EVM
	destRelayID := types.NewRelayID(string(destChainFamily), fmt.Sprintf("%d", chainID))

	bootstrapperArgs := libocr3.BootstrapperArgs{
		BootstrapperFactory:   i.peerWrapper.Peer2,
		V2Bootstrappers:       i.bootstrapperLocators,
		ContractConfigTracker: ocrimpls.NewConfigTracker(config),
		Database:              i.db,
		LocalConfig: ocrtypes.LocalConfig{
			BlockchainTimeout: 10 * time.Second,

			// Config tracking is handled by the launcher, since we're doing blue-green
			// deployments we're not going to be using OCR's built-in config switching,
			// which always shuts down the previous instance.
			ContractConfigConfirmations:        1,
			SkipContractConfigConfirmations:    true,
			ContractConfigTrackerPollInterval:  10 * time.Second,
			ContractTransmitterTransmitTimeout: 10 * time.Second,
			DatabaseTimeout:                    10 * time.Second,
			MinOCR2MaxDurationQuery:            1 * time.Second,
			DevelopmentMode:                    "false",
		},
		Logger: ocrcommon.NewOCRWrapper(
			i.lggr.
				Named("CCIPBootstrap").
				Named(destRelayID.String()).
				Named(config.Config.ChainSelector.String()).
				Named(hexutil.Encode(config.Config.OfframpAddress)),
			false, /* traceLogging */
			func(ctx context.Context, msg string) {}),
		MonitoringEndpoint: i.monitoringEndpointGen.GenMonitoringEndpoint(
			string(destChainFamily),
			destRelayID.ChainID,
			hexutil.Encode(config.Config.OfframpAddress),
			synchronization.OCR3CCIPBootstrap,
		),
		OffchainConfigDigester: ocrimpls.NewConfigDigester(config.ConfigDigest),
	}
	bootstrapper, err := libocr3.NewBootstrapper(bootstrapperArgs)
	if err != nil {
		return nil, err
	}
	return bootstrapper, nil
}

// CreatePluginOracle implements types.OracleCreator.
func (i *inprocessOracleCreator) CreatePluginOracle(pluginType cctypes.PluginType, config cctypes.OCR3ConfigWithMeta) (cctypes.CCIPOracle, error) {
	// this is so that we can use the msg hasher and report encoder from that dest chain relayer's provider.
	contractReaders := make(map[cciptypes.ChainSelector]types.ContractReader)
	for relayID, relayer := range i.relayers {
		cr, err := relayer.NewContractReader(context.Background(), crconfigs.MustCCIPReaderConfig())
		if err != nil {
			return nil, fmt.Errorf("failed to create contract reader for relay %s: %w", relayID, err)
		}

		chainID, err := strconv.ParseUint(relayID.ChainID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse chain ID %s: %w", relayID.ChainID, err)
		}

		chainSelector, ok := chainsel.EvmChainIdToChainSelector()[chainID]
		if !ok {
			return nil, fmt.Errorf("failed to get chain selector from chain ID %d", chainID)
		}

		contractReaders[cciptypes.ChainSelector(chainSelector)] = cr
	}

	// Assuming that the chain selector is referring to an evm chain for now.
	// TODO: add an api that returns chain family.
	destChainID, err := chainsel.ChainIdFromSelector(uint64(config.Config.ChainSelector))
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID from selector: %w", err)
	}
	destChainFamily := chaintype.EVM
	destRelayID := types.NewRelayID(string(destChainFamily), fmt.Sprintf("%d", destChainID))

	// build the onchain keyring. it will be the signing key for the destination chain family.
	keybundle, ok := i.ocrKeyBundles[destChainFamily]
	if !ok {
		return nil, fmt.Errorf("no OCR key bundle found for chain family %s, forgot to create one?", destChainFamily)
	}
	onchainKeyring := ocrcommon.NewOCR3OnchainKeyringAdapter(keybundle)

	oracleArgs := libocr3.OCR3OracleArgs[[]byte]{
		BinaryNetworkEndpointFactory: i.peerWrapper.Peer2,
		Database:                     i.db,
		V2Bootstrappers:              i.bootstrapperLocators,
		ContractConfigTracker:        ocrimpls.NewConfigTracker(config),
		ContractTransmitter:          nil, // TODO
		LocalConfig: ocrtypes.LocalConfig{
			BlockchainTimeout: 10 * time.Second,
			// Config tracking is handled by the launcher, since we're doing blue-green
			// deployments we're not going to be using OCR's built-in config switching,
			// which always shuts down the previous instance.
			ContractConfigConfirmations:        1,
			SkipContractConfigConfirmations:    true,
			ContractConfigTrackerPollInterval:  10 * time.Second,
			ContractTransmitterTransmitTimeout: 10 * time.Second,
			DatabaseTimeout:                    10 * time.Second,
			MinOCR2MaxDurationQuery:            1 * time.Second,
			DevelopmentMode:                    "false",
		},
		Logger: ocrcommon.NewOCRWrapper(
			i.lggr.
				Named("CCIPCommitOCR3").
				Named(destRelayID.String()).
				Named(hexutil.Encode(config.Config.OfframpAddress)),
			false,
			func(ctx context.Context, msg string) {}),
		MetricsRegisterer: prometheus.WrapRegistererWith(map[string]string{"name": fmt.Sprintf("commit-%d", config.Config.ChainSelector)}, prometheus.DefaultRegisterer),
		MonitoringEndpoint: i.monitoringEndpointGen.GenMonitoringEndpoint(
			string(destChainFamily),
			destRelayID.ChainID,
			string(config.Config.OfframpAddress),
			synchronization.OCR3CCIPCommit,
		),
		OffchainConfigDigester: ocrimpls.NewConfigDigester(config.ConfigDigest),
		OffchainKeyring:        keybundle,
		OnchainKeyring:         onchainKeyring,
		ReportingPluginFactory: commitocr3.NewPluginFactory(
			i.lggr.
				Named("CCIPCommitPlugin").
				Named(destRelayID.String()).
				Named(hexutil.Encode(config.Config.OfframpAddress)),
			ccipreaderpkg.OCR3ConfigWithMeta(config),
			ccipevm.NewCommitPluginCodecV1(),
			ccipevm.NewMessageHasherV1(),
			i.homeChainReader,
			contractReaders,
		),
	}
	oracle, err := libocr3.NewOracle(oracleArgs)
	if err != nil {
		return nil, err
	}
	return oracle, nil
}
