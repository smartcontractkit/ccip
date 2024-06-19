package ccipcapability

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/loop"
	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/launcher"
	"github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/oraclecreator"
	cctypes "github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/types"
	"github.com/smartcontractkit/chainlink/v2/core/services/job"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/chaintype"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/p2pkey"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/v2/core/services/pipeline"
	"github.com/smartcontractkit/chainlink/v2/plugins"
)

type RelayGetter interface {
	GetIDToRelayerMap() (map[types.RelayID]loop.Relayer, error)
}

type Delegate struct {
	lggr            logger.Logger
	registrarConfig plugins.RegistrarConfig
	pipelineRunner  pipeline.Runner
	relayGetter     RelayGetter
	capRegistry     cctypes.CapabilityRegistry
	keystore        keystore.Master
	ds              sqlutil.DataSource
	peerWrapper     *ocrcommon.SingletonPeerWrapper

	isNewlyCreatedJob bool
}

func NewDelegate(
	lggr logger.Logger,
	registrarConfig plugins.RegistrarConfig,
	pipelineRunner pipeline.Runner,
	relayGetter RelayGetter,
	registrySyncer cctypes.CapabilityRegistry,
	keystore keystore.Master,
	ds sqlutil.DataSource,
	peerWrapper *ocrcommon.SingletonPeerWrapper,
) *Delegate {
	return &Delegate{
		lggr:            lggr,
		registrarConfig: registrarConfig,
		pipelineRunner:  pipelineRunner,
		relayGetter:     relayGetter,
		capRegistry:     registrySyncer,
		ds:              ds,
		keystore:        keystore,
		peerWrapper:     peerWrapper,
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

	// TODO: implement
	hcr := &homeChainReader{}

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
	)

	return []job.ServiceCtx{
		hcr,
		launcher.New(
			spec.CCIPSpec.CapabilityVersion,
			spec.CCIPBootstrapSpec.CapabilityLabelledName,
			p2pID,
			d.capRegistry,
			d.lggr,
			hcr,
			oracleCreator,
		),
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
		case types.NetworkEVM:
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
