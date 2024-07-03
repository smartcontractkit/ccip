package oraclecreator

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/google/uuid"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/libocr/commontypes"
	libocr3 "github.com/smartcontractkit/libocr/offchainreporting2plus"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/loop"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/ocrimpls"
	cctypes "github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/types"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/chaintype"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
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
	panic("unimplemented")
}
