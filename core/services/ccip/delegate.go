package ccip

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/commontypes"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting/types"
	libocr2 "github.com/smartcontractkit/libocr/offchainreporting2plus"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	libocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/loop"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/chaintype"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ethkey"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/v2/core/services/synchronization"
	"github.com/smartcontractkit/chainlink/v2/core/services/telemetry"
)

type RelayGetter interface {
	// Gather all relayers the node supports in a map
	GetIDToRelayerMap() (map[types.RelayID]loop.Relayer, error)
}

type Delegate struct {
	lggr                  commontypes.Logger
	peerWrapper           *ocrcommon.SingletonPeerWrapper
	monitoringEndpointGen telemetry.MonitoringEndpointGenerator
	ks                    keystore.OCR2
	ethKs                 keystore.Eth
	rg                    RelayGetter
	cr                    CapabilityRegistry
	bootstrappers         []commontypes.BootstrapperLocator
	database              ocr3types.Database
	localConfig           ocrtypes.LocalConfig

	isNewlyCreatedJob bool
}

func NewDelegate(
	lggr commontypes.Logger,
	peerWrapper *ocrcommon.SingletonPeerWrapper,
	monitoringEndpointGen telemetry.MonitoringEndpointGenerator,
	ks keystore.OCR2,
	ethKs keystore.Eth,
	rg RelayGetter,
	capabilityRegistry CapabilityRegistry,
) *Delegate {
	return &Delegate{
		lggr:                  lggr,
		peerWrapper:           peerWrapper,
		monitoringEndpointGen: monitoringEndpointGen,
		ks:                    ks,
		ethKs:                 ethKs,
		rg:                    rg,
		isNewlyCreatedJob:     false,
		cr:                    capabilityRegistry,
	}
}

func (d *Delegate) JobType() job.Type {
	return job.CCIP
}

func (d *Delegate) BeforeJobCreated(job job.Job) {}

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
	//relayersByID, err := d.rg.GetIDToRelayerMap()
	//if err != nil {
	//	return nil, err
	//}
	// Obtain providers for all supported chains.
	commitProviders := make(map[types.RelayID]types.CCIPCommitProvider)
	execProviders := make(map[types.RelayID]types.CCIPExecProvider)
	//for relayID, relayer := range relayersByID {
	// TODO: Probably parse relayConfig to get relayConfig per chain
	// Ditto with plugin config.
	//commitProvider, err2 := relayer.NewCCIPCommitProvider(ctx,
	//	types.RelayArgs{
	//		ExternalJobID: spec.ExternalJobID,
	//		JobID:         spec.ID,
	//		ContractID:    "", // Should be deriving offramp etc from capability config contract
	//		New:           d.isNewlyCreatedJob,
	//		RelayConfig:   spec.CCIPSpec.RelayConfig.Bytes(),
	//		// TODO: Maybe NewCCIPCommitProvider as opposed to NewPluginProvider with a type?
	//		ProviderType: "CCIPCommit",
	//	}, types.PluginArgs{
	//		TransmitterID: "", // Actually seems unused?
	//		PluginConfig:  spec.CCIPSpec.PluginConfig.Bytes(),
	//	})
	//if err2 != nil {
	//	return nil, err2
	//}
	//commitProviders[relayID] = commitProvider
	// Diito with exec.
	//}

	ctx, cancel := context.WithCancel(context.Background())
	return []job.ServiceCtx{
		&ccipService{
			d:               d,
			commitProviders: commitProviders,
			execProviders:   execProviders,
			ctx:             ctx,
			cancel:          cancel,
		},
	}, nil
}

type Oracle struct {
	ocrOracle libocr2.Oracle
	config    libocr2types.ContractConfig
}

func (d *Delegate) AfterJobCreated(job job.Job) {}

func (d *Delegate) BeforeJobDeleted(job job.Job) {}

func (d *Delegate) OnDeleteJob(ctx context.Context, jb job.Job) error {
	panic("implement me")
}

type ccipService struct {
	d               *Delegate
	commitProviders map[types.RelayID]types.CCIPCommitProvider
	execProviders   map[types.RelayID]types.CCIPExecProvider
	ctx             context.Context
	cancel          context.CancelFunc
	endpoint        telemetry.MonitoringEndpointGenerator
	ocrDB           libocr2types.Database
	ocrKeys         map[types.RelayID]ocr2key.KeyBundle
	database        ocrtypes.Database
	registrySyncer  CapabilityRegistrySyncer

	// State
	dons map[uint32][]Oracle
}

func (c *ccipService) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case diff := <-c.registrySyncer.Listen(ctx):
			fmt.Println(">>> new diff: ", diff)

			// Start/close various oracles if required using providers, example of a new commit plugin
			// instance for chain evm.123:
			relayID := types.RelayID{
				Network: "evm",
				ChainID: "123",
			}

			commitProvider := c.commitProviders[relayID]

			oracleArgs := libocr2.OCR3OracleArgs[[]byte]{
				BinaryNetworkEndpointFactory: c.d.peerWrapper.Peer2, // passed in
				V2Bootstrappers:              c.d.bootstrappers,     // spec
				ContractConfigTracker:        commitProvider.ContractConfigTracker(),
				ContractTransmitter:          ocrcommon.NewOCR3ContractTransmitterAdapter(commitProvider.ContractTransmitter()),
				Database:                     c.d.database,
				//LocalConfig:                  c.d.localConfig,
				Logger:                 c.d.lggr,
				MonitoringEndpoint:     c.d.monitoringEndpointGen.GenMonitoringEndpoint("", "", "", synchronization.OCR3CCIP),
				OffchainConfigDigester: commitProvider.OffchainConfigDigester(),
				OffchainKeyring:        c.ocrKeys[relayID],
				OnchainKeyring:         ocrcommon.NewOCR3OnchainKeyringAdapter(c.ocrKeys[relayID]),
				MetricsRegisterer:      nil, // TODO
			}
			// Can maybe use this? Or our own plugin factory client directly
			//plugin := ocr3.NewLOOPPService(c.lggr, grpcOpts, cmdFn, pluginConfig, providerClientConn, pr, ta, errorLog,
			//	 capabilitiesRegistry, keyValueStore, relayerSet)
			oracleArgs.ReportingPluginFactory = nil
			oracle, _ := libocr2.NewOracle(oracleArgs)
			oracle.Start()
		}
	}
}

func (c *ccipService) Close() error {
	c.cancel()
	return nil
}

// Ensure Delegate implements job.Delegate
var _ job.Delegate = (*Delegate)(nil)
