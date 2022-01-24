// TODO: I think we might be able to make the bootstrap job type just generic for all genocr jobs?
package ccip

import (
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/core/services/offchainreporting2"
	"github.com/smartcontractkit/chainlink/core/services/relay/types"
	ocrcommontypes "github.com/smartcontractkit/libocr/commontypes"
	ocr "github.com/smartcontractkit/libocr/offchainreporting2"
	"github.com/smartcontractkit/sqlx"
)

type DelegateBootstrap struct {
	bootstrappers []ocrcommontypes.BootstrapperLocator
	db            *sqlx.DB
	jobORM        job.ORM
	chainSet      evm.ChainSet
	cfg           Config
	peerWrapper   *ocrcommon.SingletonPeerWrapper
	lggr          logger.Logger
	relayer       types.Relayer
}

// TODO: Register this delegate behind a FF
func NewDelegateBootstrap(
	db *sqlx.DB,
	jobORM job.ORM,
	chainSet evm.ChainSet,
	peerWrapper *ocrcommon.SingletonPeerWrapper,
	lggr logger.Logger,
	cfg Config,
	relayer types.Relayer,
) *DelegateBootstrap {
	return &DelegateBootstrap{
		db:          db,
		jobORM:      jobORM,
		chainSet:    chainSet,
		peerWrapper: peerWrapper,
		lggr:        lggr,
		cfg:         cfg,
		relayer:     relayer,
	}
}

func (d DelegateBootstrap) JobType() job.Type {
	return job.CCIPBootstrap
}

func (d DelegateBootstrap) ServicesForSpec(jobSpec job.Job) (services []job.Service, err error) {
	spec := jobSpec.CCIPBootstrapSpec
	if spec == nil {
		return nil, errors.Errorf("CCIPBootstrap expects a *job.CCIPBootstrapSpec to be present, got %v", jobSpec)
	}
	l := d.lggr.With(
		"jobID", jobSpec.ID,
		"externalJobID", jobSpec.ExternalJobID,
		"coordinatorAddress", spec.ContractAddress,
	)
	ocr2Spec := spec.AsOCR2Spec()
	ocr2Provider, err := d.relayer.NewOCR2Provider(jobSpec.ExternalJobID, &ocr2Spec)
	if err != nil {
		return nil, errors.Wrap(err, "error calling 'relayer.NewOCR2Provider'")
	}
	services = append(services, ocr2Provider)

	lc := offchainreporting2.ToLocalConfig(d.cfg, ocr2Spec)
	if err = ocr.SanityCheckLocalConfig(lc); err != nil {
		return nil, err
	}

	ocrdb := NewDB(d.db.DB, spec.ContractAddress.Address(), d.lggr)
	ocrLogger := logger.NewOCRWrapper(l, true, func(msg string) {
		_ = d.jobORM.RecordError(jobSpec.ID, msg)
	})

	bootstrapNode, err := ocr.NewBootstrapper(ocr.BootstrapperArgs{
		BootstrapperFactory:    d.peerWrapper.Peer2,
		ContractConfigTracker:  ocr2Provider.ContractConfigTracker(),
		Database:               ocrdb,
		LocalConfig:            lc,
		Logger:                 ocrLogger,
		OffchainConfigDigester: ocr2Provider.OffchainConfigDigester(),
	})
	if err != nil {
		return nil, err
	}
	services = append(services, bootstrapNode)
	return services, nil
}

func (d DelegateBootstrap) AfterJobCreated(spec job.Job) {
}

func (d DelegateBootstrap) BeforeJobDeleted(spec job.Job) {
}
