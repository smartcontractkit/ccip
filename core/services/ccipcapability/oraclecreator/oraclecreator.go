package oraclecreator

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/loop"
	"github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	cctypes "github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/types"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/chaintype"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/v2/core/services/synchronization"
	"github.com/smartcontractkit/chainlink/v2/core/services/telemetry"
	libocr3 "github.com/smartcontractkit/libocr/offchainreporting2plus"
)

type oracleCreator struct {
	ocrKeyBundles         map[chaintype.ChainType]ocr2key.KeyBundle
	transmitters          map[types.RelayID][]string
	relayers              map[types.RelayID]loop.Relayer
	peerWrapper           *ocrcommon.SingletonPeerWrapper
	externalJobID         uuid.UUID
	jobID                 int32
	isNewlyCreatedJob     bool
	relayConfigs          map[string]job.JSONConfig
	pluginConfig          job.JSONConfig
	db                    ocr3types.Database
	lggr                  logger.Logger
	monitoringEndpointGen telemetry.MonitoringEndpointGenerator
	bootstrapperLocators  []commontypes.BootstrapperLocator
}

func New(
	ocrKeyBundles map[chaintype.ChainType]ocr2key.KeyBundle,
	transmitters map[types.RelayID][]string,
	relayers map[types.RelayID]loop.Relayer,
	peerWrapper *ocrcommon.SingletonPeerWrapper,
	externalJobID uuid.UUID,
	jobID int32,
	isNewlyCreatedJob bool,
	relayConfigs map[string]job.JSONConfig,
	pluginConfig job.JSONConfig,
	db ocr3types.Database,
	lggr logger.Logger,
	monitoringEndpointGen telemetry.MonitoringEndpointGenerator,
) cctypes.OracleCreator {
	return &oracleCreator{
		ocrKeyBundles:         ocrKeyBundles,
		transmitters:          transmitters,
		relayers:              relayers,
		peerWrapper:           peerWrapper,
		externalJobID:         externalJobID,
		jobID:                 jobID,
		isNewlyCreatedJob:     isNewlyCreatedJob,
		relayConfigs:          relayConfigs,
		pluginConfig:          pluginConfig,
		db:                    db,
		lggr:                  lggr,
		monitoringEndpointGen: monitoringEndpointGen,
	}
}

// CreateCommitOracle implements types.OracleCreator.
func (o *oracleCreator) CreateCommitOracle(config cctypes.OCRConfig) (cctypes.CCIPOracle, error) {
	// for now we assume that we have a relayer for the destination chain.
	// this is so that we can use the msg hasher and report encoder from that dest chain relayer's provider.
	providers := make(map[types.RelayID]types.CCIPOCR3CommitProvider)
	for relayID, relayer := range o.relayers {
		provider, err := relayer.NewPluginProvider(context.Background(), types.RelayArgs{
			ExternalJobID: o.externalJobID,
			JobID:         o.jobID,
			ContractID:    "", // TODO: figure out contract ID
			New:           o.isNewlyCreatedJob,
			RelayConfig:   o.relayConfigs[relayID.String()].Bytes(),
			ProviderType:  "CCIPOCR3CommitProvider",
		}, types.PluginArgs{
			TransmitterID: o.transmitters[relayID][0],
			PluginConfig:  o.pluginConfig.Bytes(),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create ocr3 commit plugin provider for relay %s: %w", relayID.String(), err)
		}

		commitProvider, ok := provider.(types.CCIPOCR3CommitProvider)
		if !ok {
			return nil, fmt.Errorf("expected CCIPOCR3CommitProvider, got %T", provider)
		}

		providers[relayID] = commitProvider
	}

	// Assuming that the chain selector is referring to an evm chain for now.
	// TODO: add an api that returns chain family.
	chainID, err := chainsel.ChainIdFromSelector(config.ChainSelector())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID from selector: %w", err)
	}

	destChainFamily := chaintype.EVM
	destRelayID := types.NewRelayID(string(destChainFamily), fmt.Sprintf("%d", chainID))
	destProvider, ok := providers[destRelayID]
	if !ok {
		return nil, fmt.Errorf("no provider found for destination chain %s in providers map", destRelayID.String())
	}

	// Adapt the provider's contract transmitter for OCR3, unless
	// the provider exposes an OCR3ContractTransmitter interface, in which case
	// we'll use that instead.
	contractTransmitter := ocr3types.ContractTransmitter[[]byte](
		ocrcommon.NewOCR3ContractTransmitterAdapter(destProvider.ContractTransmitter()),
	)
	if ocr3Provider, ok := destProvider.(types.OCR3ContractTransmitter); ok {
		contractTransmitter = ocr3Provider.OCR3ContractTransmitter()
	}

	// build the onchain keyring. it will be the signing key for the destination chain family.
	keybundle, ok := o.ocrKeyBundles[destChainFamily]
	if !ok {
		return nil, fmt.Errorf("no OCR key bundle found for chain family %s, forgot to create one?", destChainFamily)
	}
	onchainKeyring := ocrcommon.NewOCR3OnchainKeyringAdapter(keybundle)

	ocrLogger := ocrcommon.NewOCRWrapper(o.lggr, false /* traceLogging */, func(ctx context.Context, msg string) {
		// o.lggr.ErrorIf(d.jobORM.RecordError(ctx, jb.ID, msg), "unable to record error") // TODO
	})
	oracleArgs := libocr3.OCR3OracleArgs[[]byte]{
		BinaryNetworkEndpointFactory: o.peerWrapper.Peer2,
		Database:                     o.db,
		V2Bootstrappers:              o.bootstrapperLocators,
		ContractConfigTracker:        nil, // TODO
		ContractTransmitter:          contractTransmitter,
		LocalConfig: ocrtypes.LocalConfig{
			BlockchainTimeout:                  10 * time.Second,
			ContractConfigConfirmations:        0,
			SkipContractConfigConfirmations:    false,
			ContractConfigTrackerPollInterval:  10 * time.Second,
			ContractTransmitterTransmitTimeout: 10 * time.Second,
			DatabaseTimeout:                    10 * time.Second,
			MinOCR2MaxDurationQuery:            1 * time.Second,
			DevelopmentMode:                    "false",
		},
		Logger:            ocrLogger,
		MetricsRegisterer: prometheus.WrapRegistererWith(map[string]string{"name": fmt.Sprintf("commit-%d", config.ChainSelector())}, prometheus.DefaultRegisterer),
		MonitoringEndpoint: o.monitoringEndpointGen.GenMonitoringEndpoint(
			string(destChainFamily),
			destRelayID.ChainID,
			string(config.OfframpAddress()),
			synchronization.OCR3CCIPCommit,
		),
		OffchainConfigDigester: nil, // TODO
		OffchainKeyring:        keybundle,
		OnchainKeyring:         onchainKeyring,
		ReportingPluginFactory: nil, // TODO
	}
	oracle, err := libocr3.NewOracle(oracleArgs)
	if err != nil {
		return nil, err
	}
	return oracle, nil
}

// CreateExecOracle implements types.OracleCreator.
func (o *oracleCreator) CreateExecOracle(config cctypes.OCRConfig) (cctypes.CCIPOracle, error) {
	// for now we assume that we have a relayer for the destination chain.
	// this is so that we can use the msg hasher and report encoder from that dest chain relayer's provider.
	providers := make(map[types.RelayID]types.CCIPOCR3ExecuteProvider)
	for relayID, relayer := range o.relayers {
		provider, err := relayer.NewPluginProvider(context.Background(), types.RelayArgs{
			ExternalJobID: o.externalJobID,
			JobID:         o.jobID,
			ContractID:    "", // TODO: figure out contract ID
			New:           o.isNewlyCreatedJob,
			RelayConfig:   o.relayConfigs[relayID.String()].Bytes(),
			ProviderType:  "CCIPOCR3ExecuteProvider",
		}, types.PluginArgs{
			TransmitterID: o.transmitters[relayID][0],
			PluginConfig:  o.pluginConfig.Bytes(),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create ocr3 exec plugin provider for relay %s: %w", relayID.String(), err)
		}

		execProvider, ok := provider.(types.CCIPOCR3ExecuteProvider)
		if !ok {
			return nil, fmt.Errorf("expected CCIPOCR3CommitProvider, got %T", provider)
		}

		providers[relayID] = execProvider
	}

	// Assuming that the chain selector is referring to an evm chain for now.
	// TODO: add an api that returns chain family.
	chainID, err := chainsel.ChainIdFromSelector(config.ChainSelector())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID from selector: %w", err)
	}

	destChainFamily := chaintype.EVM
	destRelayID := types.NewRelayID(string(destChainFamily), fmt.Sprintf("%d", chainID))
	destProvider, ok := providers[destRelayID]
	if !ok {
		return nil, fmt.Errorf("no provider found for destination chain %s in providers map", destRelayID.String())
	}

	// Adapt the provider's contract transmitter for OCR3, unless
	// the provider exposes an OCR3ContractTransmitter interface, in which case
	// we'll use that instead.
	contractTransmitter := ocr3types.ContractTransmitter[[]byte](
		ocrcommon.NewOCR3ContractTransmitterAdapter(destProvider.ContractTransmitter()),
	)
	if ocr3Provider, ok := destProvider.(types.OCR3ContractTransmitter); ok {
		contractTransmitter = ocr3Provider.OCR3ContractTransmitter()
	}

	// build the onchain keyring. it will be the signing key for the destination chain family.
	keybundle, ok := o.ocrKeyBundles[destChainFamily]
	if !ok {
		return nil, fmt.Errorf("no OCR key bundle found for chain family %s, forgot to create one?", destChainFamily)
	}
	onchainKeyring := ocrcommon.NewOCR3OnchainKeyringAdapter(keybundle)

	ocrLogger := ocrcommon.NewOCRWrapper(o.lggr, false /* traceLogging */, func(ctx context.Context, msg string) {
		// o.lggr.ErrorIf(d.jobORM.RecordError(ctx, jb.ID, msg), "unable to record error") // TODO
	})
	oracleArgs := libocr3.OCR3OracleArgs[[]byte]{
		BinaryNetworkEndpointFactory: o.peerWrapper.Peer2,
		Database:                     o.db,
		V2Bootstrappers:              o.bootstrapperLocators,
		ContractConfigTracker:        nil, // TODO
		ContractTransmitter:          contractTransmitter,
		LocalConfig: ocrtypes.LocalConfig{
			BlockchainTimeout:                  10 * time.Second,
			ContractConfigConfirmations:        0,
			SkipContractConfigConfirmations:    false,
			ContractConfigTrackerPollInterval:  10 * time.Second,
			ContractTransmitterTransmitTimeout: 10 * time.Second,
			DatabaseTimeout:                    10 * time.Second,
			MinOCR2MaxDurationQuery:            1 * time.Second,
			DevelopmentMode:                    "false",
		},
		Logger:            ocrLogger,
		MetricsRegisterer: prometheus.WrapRegistererWith(map[string]string{"name": fmt.Sprintf("exec-%d", config.ChainSelector())}, prometheus.DefaultRegisterer),
		MonitoringEndpoint: o.monitoringEndpointGen.GenMonitoringEndpoint(
			string(destChainFamily),
			destRelayID.ChainID,
			string(config.OfframpAddress()),
			synchronization.OCR3CCIPExec,
		),
		OffchainConfigDigester: nil, // TODO
		OffchainKeyring:        keybundle,
		OnchainKeyring:         onchainKeyring,
		ReportingPluginFactory: nil, // TODO
	}
	oracle, err := libocr3.NewOracle(oracleArgs)
	if err != nil {
		return nil, err
	}
	return oracle, nil
}

// CreateBootstrapOracle implements types.OracleCreator.
func (o *oracleCreator) CreateBootstrapOracle(config cctypes.OCRConfig) (cctypes.CCIPOracle, error) {
	// Assuming that the chain selector is referring to an evm chain for now.
	// TODO: add an api that returns chain family.
	chainID, err := chainsel.ChainIdFromSelector(config.ChainSelector())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID from selector: %w", err)
	}

	destChainFamily := chaintype.EVM
	destRelayID := types.NewRelayID(string(destChainFamily), fmt.Sprintf("%d", chainID))

	bootstrapperArgs := libocr3.BootstrapperArgs{
		BootstrapperFactory:   o.peerWrapper.Peer2,
		V2Bootstrappers:       o.bootstrapperLocators,
		ContractConfigTracker: nil, // TODO
		Database:              o.db,
		LocalConfig: ocrtypes.LocalConfig{
			BlockchainTimeout:                  10 * time.Second,
			ContractConfigConfirmations:        0,
			SkipContractConfigConfirmations:    false,
			ContractConfigTrackerPollInterval:  10 * time.Second,
			ContractTransmitterTransmitTimeout: 10 * time.Second,
			DatabaseTimeout:                    10 * time.Second,
			MinOCR2MaxDurationQuery:            1 * time.Second,
			DevelopmentMode:                    "false",
		},
		Logger: ocrcommon.NewOCRWrapper(o.lggr, false /* traceLogging */, func(ctx context.Context, msg string) {
			// o.lggr.ErrorIf(d.jobORM.RecordError(ctx, jb.ID, msg), "unable to record error") // TODO
		}),
		MonitoringEndpoint: o.monitoringEndpointGen.GenMonitoringEndpoint(
			string(destChainFamily),
			destRelayID.ChainID,
			string(config.OfframpAddress()),
			synchronization.OCR3CCIPBootstrap,
		),
		OffchainConfigDigester: nil, // TODO
	}
	bootstrapper, err := libocr3.NewBootstrapper(bootstrapperArgs)
	if err != nil {
		return nil, err
	}
	return bootstrapper, nil
}
