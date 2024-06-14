package ccipcapability

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

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
		return nil, errors.Wrapf(err, "failed to make peer ID from provided spec p2p id: %s", spec.CCIPSpec.P2PKeyID)
	}

	p2pID, err := d.keystore.P2P().Get(peerID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all p2p keys")
	}

	ocrKeys := make(map[chaintype.ChainType]ocr2key.KeyBundle)
	for chainType, bundleAny := range spec.CCIPSpec.OCRKeyBundleIDs {
		ct := chaintype.ChainType(chainType)
		if !chaintype.IsSupportedChainType(ct) {
			return nil, errors.Errorf("unsupported chain type: %s", chainType)
		}

		bundleID, ok := bundleAny.(string)
		if !ok {
			return nil, errors.New("OCRKeyBundleIDs must be a map of chain types to OCR key bundle IDs")
		}

		bundle, err2 := d.keystore.OCR2().Get(bundleID)
		if err2 != nil {
			return nil, errors.Wrapf(err2, "OCR key bundle with ID %s not found", bundleID)
		}

		ocrKeys[ct] = bundle
	}

	relayers, err := d.relayGetter.GetIDToRelayerMap()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all relayers")
	}

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
		case types.NetworkCosmos:
			cosmosKeys, err2 := d.keystore.Cosmos().GetAll()
			if err2 != nil {
				return nil, fmt.Errorf("error getting all cosmos keys: %w", err2)
			}

			transmitterKeys[relayID] = func() (r []string) {
				for _, key := range cosmosKeys {
					r = append(r, key.String())
				}
				return
			}()
		case types.NetworkSolana:
			solKey, err2 := d.keystore.Solana().GetAll()
			if err2 != nil {
				return nil, fmt.Errorf("error getting all solana keys: %w", err2)
			}

			transmitterKeys[relayID] = func() (r []string) {
				for _, key := range solKey {
					r = append(r, key.String())
				}
				return
			}()
		case types.NetworkStarkNet:
			starkKey, err2 := d.keystore.StarkNet().GetAll()
			if err2 != nil {
				return nil, fmt.Errorf("error getting all stark keys: %w", err2)
			}

			transmitterKeys[relayID] = func() (r []string) {
				for _, key := range starkKey {
					r = append(r, key.String())
				}
				return
			}()
		default:
			return nil, errors.Errorf("unsupported network: %s", relayID.Network)
		}
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
