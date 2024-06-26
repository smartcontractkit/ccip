package ccipcapability

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	ocr3reader "github.com/smartcontractkit/ccipocr3/pkg/reader"
	"github.com/smartcontractkit/chainlink-common/pkg/loop"
	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
	"github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/config"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/ccip_capability_configuration"
	kcr "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/capabilities_registry"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/launcher"
	"github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/oraclecreator"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/chaintype"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/p2pkey"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/v2/core/services/pipeline"
	"github.com/smartcontractkit/chainlink/v2/core/services/registrysyncer"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay"
	evmrelaytypes "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/services/telemetry"
	"github.com/smartcontractkit/chainlink/v2/plugins"
)

type RelayGetter interface {
	GetIDToRelayerMap() (map[types.RelayID]loop.Relayer, error)
}

type Delegate struct {
	lggr                  logger.Logger
	registrarConfig       plugins.RegistrarConfig
	pipelineRunner        pipeline.Runner
	relayGetter           RelayGetter
	keystore              keystore.Master
	ds                    sqlutil.DataSource
	peerWrapper           *ocrcommon.SingletonPeerWrapper
	monitoringEndpointGen telemetry.MonitoringEndpointGenerator
	registrySyncer        registrysyncer.Syncer
	capabilityConfig      config.Capabilities

	isNewlyCreatedJob bool
}

func NewDelegate(
	lggr logger.Logger,
	registrarConfig plugins.RegistrarConfig,
	pipelineRunner pipeline.Runner,
	relayGetter RelayGetter,
	registrySyncer registrysyncer.Syncer,
	keystore keystore.Master,
	ds sqlutil.DataSource,
	peerWrapper *ocrcommon.SingletonPeerWrapper,
	monitoringEndpointGen telemetry.MonitoringEndpointGenerator,
) *Delegate {
	return &Delegate{
		lggr:                  lggr,
		registrarConfig:       registrarConfig,
		pipelineRunner:        pipelineRunner,
		relayGetter:           relayGetter,
		registrySyncer:        registrySyncer,
		ds:                    ds,
		keystore:              keystore,
		peerWrapper:           peerWrapper,
		monitoringEndpointGen: monitoringEndpointGen,
	}
}

func (d *Delegate) JobType() job.Type {
	return job.CCIP
}

func (d *Delegate) BeforeJobCreated(job.Job) {
	// This is only called first time the job is created
	d.isNewlyCreatedJob = true
}

func (d *Delegate) ServicesForSpec(ctx context.Context, spec job.Job) (services []job.ServiceCtx, err error) {
	// In general there should only be one P2P key but the node may have multiple.
	// The job spec should specify the correct P2P key to use.
	peerID, err := p2pkey.MakePeerID(spec.CCIPSpec.P2PKeyID)
	if err != nil {
		return nil, fmt.Errorf("failed to make peer ID from provided spec p2p id (%s): %w", spec.CCIPSpec.P2PKeyID, err)
	}

	p2pID, err := d.keystore.P2P().Get(peerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get all p2p keys: %w", err)
	}

	ocrKeys, err := d.getOCRKeys(spec.CCIPSpec.OCRKeyBundleIDs)
	if err != nil {
		return nil, err
	}

	relayers, err := d.relayGetter.GetIDToRelayerMap()
	if err != nil {
		return nil, fmt.Errorf("failed to get all relayers: %w", err)
	}

	transmitterKeys, err := d.getTransmitterKeys(ctx, relayers)
	if err != nil {
		return nil, err
	}

	// NOTE: we can use the same DB for all plugin instances,
	// since all queries are scoped by config digest.
	ocrDB := ocr2.NewDB(d.ds, spec.ID, 0, d.lggr)

	oracleCreator := oraclecreator.New(
		ocrKeys,
		transmitterKeys,
		relayers,
		d.peerWrapper,
		spec.ExternalJobID,
		spec.ID,
		d.isNewlyCreatedJob,
		spec.CCIPSpec.RelayConfigs,
		spec.CCIPSpec.PluginConfig,
		ocrDB,
		d.lggr,
		d.monitoringEndpointGen,
	)

	homeChainContractReader, err := d.getHomeChainContractReader(relayers,
		spec.CCIPSpec.CapabilityLabelledName,
		spec.CCIPSpec.CapabilityVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to get home chain contract reader: %w", err)
	}

	hcr := ocr3reader.NewHomeChainReader(
		homeChainContractReader,
		d.lggr.Named("HomeChainReader"),
		12*time.Second,
	)

	capLauncher := launcher.New(
		spec.CCIPSpec.CapabilityVersion,
		spec.CCIPSpec.CapabilityLabelledName,
		p2pID,
		d.lggr,
		hcr,
		oracleCreator,
	)

	// register the capability launcher with the registry syncer
	d.registrySyncer.AddLauncher(capLauncher)

	return []job.ServiceCtx{
		hcr,
		capLauncher,
	}, nil
}

func (d *Delegate) AfterJobCreated(spec job.Job) {}

func (d *Delegate) BeforeJobDeleted(spec job.Job) {}

func (d *Delegate) OnDeleteJob(ctx context.Context, spec job.Job) error {
	// TODO: shut down needed services?
	return nil
}

func (d *Delegate) getOCRKeys(ocrKeyBundleIDs job.JSONConfig) (map[chaintype.ChainType]ocr2key.KeyBundle, error) {
	ocrKeys := make(map[chaintype.ChainType]ocr2key.KeyBundle)
	for chainType, bundleIDRaw := range ocrKeyBundleIDs {
		ct := chaintype.ChainType(chainType)
		if !chaintype.IsSupportedChainType(ct) {
			return nil, fmt.Errorf("unsupported chain type: %s", chainType)
		}

		bundleID, ok := bundleIDRaw.(string)
		if !ok {
			return nil, fmt.Errorf("OCRKeyBundleIDs must be a map of chain types to OCR key bundle IDs, got: %T", bundleIDRaw)
		}

		bundle, err2 := d.keystore.OCR2().Get(bundleID)
		if err2 != nil {
			return nil, fmt.Errorf("OCR key bundle with ID %s not found: %w", bundleID, err2)
		}

		ocrKeys[ct] = bundle
	}
	return ocrKeys, nil
}

func (d *Delegate) getTransmitterKeys(ctx context.Context, relayers map[types.RelayID]loop.Relayer) (map[types.RelayID][]string, error) {
	transmitterKeys := make(map[types.RelayID][]string)
	for relayID := range relayers {
		switch relayID.Network {
		case relay.NetworkEVM:
			ethKeys, err2 := d.keystore.Eth().GetAll(ctx)
			if err2 != nil {
				return nil, fmt.Errorf("error getting all eth keys: %w", err2)
			}

			transmitterKeys[relayID] = func() (r []string) {
				for _, key := range ethKeys {
					r = append(r, key.String())
				}
				return
			}()
		default:
			return nil, fmt.Errorf("unsupported network: %s", relayID.Network)
		}
	}
	return transmitterKeys, nil
}

func (d *Delegate) getHomeChainContractReader(
	relayers map[types.RelayID]loop.Relayer,
	capabilityLabelledName,
	capabilityVersion string,
) (types.ContractReader, error) {
	// home chain is where the capability registry is deployed,
	// which should be set correctly in toml config.
	homeChainRelayID := d.capabilityConfig.ExternalRegistry().RelayID()
	homeChainRelayer, ok := relayers[homeChainRelayID]
	if !ok {
		return nil, fmt.Errorf("home chain relayer not found, chain id: %s", homeChainRelayID.String())
	}

	cccChainReaderConfig := homeChainReaderConfig()
	encoded, err := json.Marshal(cccChainReaderConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal CCC chain reader config: %w", err)
	}

	reader, err := homeChainRelayer.NewContractReader(context.Background(), encoded)
	if err != nil {
		return nil, fmt.Errorf("failed to create home chain contract reader: %w", err)
	}

	reader, err = bindReader(reader, d.capabilityConfig.ExternalRegistry().Address(), capabilityLabelledName, capabilityVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to bind home chain contract reader: %w", err)
	}

	return reader, nil
}

func hashedCapabilityId(capabilityLabelledName, capabilityVersion string) (r [32]byte, err error) {
	tabi := `[{"type": "string"}, {"type": "string"}]`
	abiEncoded, err := utils.ABIEncode(tabi, capabilityLabelledName, capabilityVersion)
	if err != nil {
		return r, fmt.Errorf("failed to ABI encode capability version and labelled name: %w", err)
	}

	h := crypto.Keccak256(abiEncoded)
	copy(r[:], h)
	return r, nil
}

func bindReader(reader types.ContractReader, capRegAddress, capabilityLabelledName, capabilityVersion string) (types.ContractReader, error) {
	err := reader.Bind(context.Background(), []types.BoundContract{
		{
			Address: capRegAddress,
			Name:    "CapabilitiesRegistry",
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to bind home chain contract reader: %w", err)
	}

	hid, err := hashedCapabilityId(capabilityLabelledName, capabilityVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to hash capability id: %w", err)
	}

	var ccipCapabilityInfo kcr.CapabilitiesRegistryCapabilityInfo
	err = reader.GetLatestValue(context.Background(), "CapabilitiesRegistry", "getCapability", map[string]any{
		"hashedId": hid,
	}, &ccipCapabilityInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to get CCIP capability info from chain reader: %w", err)
	}

	// bind the ccip capability configuration contract
	err = reader.Bind(context.Background(), []types.BoundContract{
		{
			Address: ccipCapabilityInfo.ConfigurationContract.String(),
			Name:    "CCIPCapabilityConfiguration",
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to bind CCIP capability configuration contract: %w", err)
	}

	return reader, nil
}

func homeChainReaderConfig() evmrelaytypes.ChainReaderConfig {
	return evmrelaytypes.ChainReaderConfig{
		Contracts: map[string]evmrelaytypes.ChainContractReader{
			"CapabilitiesRegistry": {
				ContractABI: kcr.CapabilitiesRegistryABI,
				Configs: map[string]*evmrelaytypes.ChainReaderDefinition{
					"getCapability": {
						ChainSpecificName: "getCapability",
					},
				},
			},
			"CCIPCapabilityConfiguration": {
				ContractABI: ccip_capability_configuration.CCIPCapabilityConfigurationABI,
				Configs: map[string]*evmrelaytypes.ChainReaderDefinition{
					"getAllChainConfigs": {
						ChainSpecificName: "getAllChainConfigs",
					},
					"getOCRConfig": {
						ChainSpecificName: "getOCRConfig",
					},
				},
			},
		},
	}
}
