package ccip

import (
	"context"
	"errors"
	"github.com/smartcontractkit/chainlink-common/pkg/loop"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/chaintype"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ethkey"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/v2/core/services/telemetry"
	libocr2 "github.com/smartcontractkit/libocr/offchainreporting2plus"
	libocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"time"
)

// TODO: Probably all need to live in capabilities pkg
type CapabilityID = string

type CapabilityConfiguration struct {
	CapabilityID CapabilityID // Versioned etc.
	Config       []byte
}

type DON struct {
	ID                       uint32
	IsPublic                 bool
	Nodes                    [][]byte
	CapabilityConfigurations []CapabilityConfiguration
}

type CapabilitiesRegistry interface {
	// Extend capabilities.Registry with this to unblock CCIP?
	GetDONsWithCapability(ctx context.Context, capability CapabilityID) map[uint32]DON
}

type RelayGetter interface {
	// Gather all relayers the node supports in a map
	GetIDToRelayerMap() (map[types.RelayID]loop.Relayer, error)
}

type Delegate struct {
	lggr                  logger.Logger
	peerWrapper           *ocrcommon.SingletonPeerWrapper
	monitoringEndpointGen telemetry.MonitoringEndpointGenerator
	ks                    keystore.OCR2
	ethKs                 keystore.Eth
	rg                    RelayGetter
	capabilitiesRegistry  CapabilitiesRegistry
	isNewlyCreatedJob     bool
}

var _ job.Delegate = (*Delegate)(nil)

func NewDelegate(
	lggr logger.Logger,
	peerWrapper *ocrcommon.SingletonPeerWrapper,
	monitoringEndpointGen telemetry.MonitoringEndpointGenerator,
	ks keystore.OCR2,
	ethKs keystore.Eth,
	rg RelayGetter,
	capabilitiesRegistry CapabilitiesRegistry,
) *Delegate {
	return &Delegate{
		lggr:                  lggr.Named("CCIP"),
		peerWrapper:           peerWrapper,
		monitoringEndpointGen: monitoringEndpointGen,
		ks:                    ks,
		ethKs:                 ethKs,
		rg:                    rg,
		isNewlyCreatedJob:     false,
		capabilitiesRegistry:  capabilitiesRegistry,
	}
}

func (d *Delegate) JobType() job.Type {
	return job.CCIP
}

func (d *Delegate) BeforeJobCreated(spec job.Job) {}

func (d *Delegate) AfterJobCreated(jb job.Job) {}

func (d *Delegate) BeforeJobDeleted(spec job.Job) {}

func (d *Delegate) OnDeleteJob(context.Context, job.Job) error { return nil }

// ServicesForSpec satisfies the job.Delegate interface.
func (d *Delegate) ServicesForSpec(ctx context.Context, spec job.Job) ([]job.ServiceCtx, error) {
	// validate the job spec

	// TODO: we could even try to autodetect these and remove from job spec?
	// Safer than auto-detecting transmitter keys.
	ocrKeys := make(map[chaintype.ChainType]ocr2key.KeyBundle)
	for chainType, ocrKeyBundle := range spec.CCIPSpec.OCRKeyBundleIDs {
		keyBundle, err := d.ks.Get(ocrKeyBundle)
		if err != nil {
			return nil, err
		}
		if string(keyBundle.ChainType()) != chainType {
			return nil, errors.New("key bundle chain type does not match")
		}
		ocrKeys[chaintype.ChainType(chainType)] = keyBundle
	}

	transmitterKeys := make(map[string]ethkey.KeyV2)
	for relayID, transmitterID := range spec.CCIPSpec.TransmitterIDs {
		// TODO, need to parse relayID and use appropriate key store for non-EVM chains.
		transmitterKey, err := d.ethKs.Get(ctx, transmitterID)
		if err != nil {
			return nil, err
		}
		transmitterKeys[relayID] = transmitterKey
	}
	// All supported relayers
	relayersByID, err := d.rg.GetIDToRelayerMap()
	if err != nil {
		return nil, err
	}
	// Obtain providers for all supported chains.
	commitProviders := make(map[types.RelayID]types.CCIPCommitProvider)
	execProviders := make(map[types.RelayID]types.CCIPExecProvider)
	for relayID, relayer := range relayersByID {
		// TODO: Probably parse relayConfig to get relayConfig per chain
		// Ditto with plugin config.
		commitProvider, err2 := relayer.NewCCIPCommitProvider(ctx,
			types.RelayArgs{
				ExternalJobID: spec.ExternalJobID,
				JobID:         spec.ID,
				ContractID:    "", // Should be deriving offramp etc from capability config contract
				New:           d.isNewlyCreatedJob,
				RelayConfig:   spec.CCIPSpec.RelayConfig.Bytes(),
				// TODO: Maybe NewCCIPCommitProvider as opposed to NewPluginProvider with a type?
				ProviderType: "CCIPCommit",
			}, types.PluginArgs{
				TransmitterID: "", // Actually seems unused?
				PluginConfig:  spec.CCIPSpec.PluginConfig.Bytes(),
			})
		if err2 != nil {
			return nil, err2
		}
		commitProviders[relayID] = commitProvider
		// Diito with exec.

	}

	ctx, cancel := context.WithCancel(context.Background())
	return []job.ServiceCtx{&ccipService{
		capabilitiesRegistry: d.capabilitiesRegistry,
		commitProviders:      commitProviders,
		execProviders:        execProviders,
		ctx:                  ctx,
		cancel:               cancel,
	},
	}, nil
}

type Oracle struct {
	ocrOracle libocr2.Oracle
	config    libocr2types.ContractConfig
}

type ccipService struct {
	capabilitiesRegistry CapabilitiesRegistry
	commitProviders      map[types.RelayID]types.CCIPCommitProvider
	execProviders        map[types.RelayID]types.CCIPExecProvider
	ctx                  context.Context
	cancel               context.CancelFunc
	endpoint             telemetry.MonitoringEndpointGenerator
	ocrDB                libocr2types.Database
	ocrKeys              map[types.RelayID]ocr2key.KeyBundle

	// State
	dons map[uint32][]Oracle
}

func (c *ccipService) computeChanges(have map[uint32]DON, onchain map[uint32][]Oracle) []libocr2types.ContractConfig {
	return nil
}

func (c *ccipService) Start(ctx context.Context) error {
	tc := time.Tick(5 * time.Second)
	// TODO: Background
	for {
		select {
		case <-tc:
			// Read registry to determine if new OCR instances created
			dons := c.capabilitiesRegistry.GetDONsWithCapability(ctx, "ccip")
			_ = c.computeChanges(dons, c.dons)
			// Start/close various oracles if required using providers, example of a new commit plugin
			// instance for chain evm.123:
			relayID := types.RelayID{
				Network: "evm",
				ChainID: "123",
			}
			commitProvider := c.commitProviders[relayID]
			oracleArgs := libocr2.OCR3OracleArgs[[]byte]{
				BinaryNetworkEndpointFactory: c.peerWrapper.Peer2, // passed in
				V2Bootstrappers:              c.bootstrappers,     // spec
				ContractConfigTracker:        commitProvider.ContractConfigTracker(),
				ContractTransmitter:          ocrcommon.NewOCR3ContractTransmitterAdapter(commitProvider.ContractTransmitter()),
				Database:                     c.ocrDB, // TODO: Investigate this
				LocalConfig:                  lc,
				Logger:                       ocrLogger,
				MonitoringEndpoint:           c.monitoringEndpoint,
				OffchainConfigDigester:       commitProvider.OffchainConfigDigester(),
				OffchainKeyring:              c.ocrKeys[relayID],
				OnchainKeyring:               ocrcommon.NewOCR3OnchainKeyringAdapter(c.ocrKeys[relayID]),
				MetricsRegisterer:            nil, // TODO
			}
			// Can maybe use this? Or our own plugin factory client directly
			//plugin := ocr3.NewLOOPPService(c.lggr, grpcOpts, cmdFn, pluginConfig, providerClientConn, pr, ta, errorLog,
			//	 capabilitiesRegistry, keyValueStore, relayerSet)
			oracleArgs.ReportingPluginFactory = nil
			oracle, _ := libocr2.NewOracle(oracleArgs)
			oracle.Start()

		case <-ctx.Done():
			return nil
		}
	}
	return nil
}

func (c *ccipService) Close() error {
	c.cancel()
	return nil
}
