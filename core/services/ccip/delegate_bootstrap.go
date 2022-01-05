// TODO: I think we might be able to make the bootstrap job type just generic for all genocr jobs?
package ccip

import (
	"context"

	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/core/services/postgres"
	ocrcommontypes "github.com/smartcontractkit/libocr/commontypes"
	ocr "github.com/smartcontractkit/libocr/offchainreporting2"
	"github.com/smartcontractkit/libocr/offchainreporting2/chains/evmutil"
	"gorm.io/gorm"
)

type DelegateBootstrap struct {
	bootstrappers []ocrcommontypes.BootstrapperLocator
	db            *gorm.DB
	jobORM        job.ORM
	orm           ORM
	chainSet      evm.ChainSet
	peerWrapper   *ocrcommon.SingletonPeerWrapper
}

// TODO: Register this delegate behind a FF
func NewDelegateBootstrap(
	db *gorm.DB,
	jobORM job.ORM,
	chainSet evm.ChainSet,
	peerWrapper *ocrcommon.SingletonPeerWrapper,
) *DelegateBootstrap {
	return &DelegateBootstrap{
		db:          db,
		jobORM:      jobORM,
		orm:         NewORM(postgres.UnwrapGormDB(db)),
		chainSet:    chainSet,
		peerWrapper: peerWrapper,
	}
}

func (d DelegateBootstrap) JobType() job.Type {
	return job.CCIPBootstrap
}

func (d DelegateBootstrap) ServicesForSpec(jb job.Job) ([]job.Service, error) {
	if jb.CCIPBootstrapSpec == nil {
		return nil, errors.New("no bootstrap job specified")
	}
	l := logger.Default.With(
		"jobID", jb.ID,
		"externalJobID", jb.ExternalJobID,
		"coordinatorAddress", jb.CCIPBootstrapSpec.ContractAddress,
	)

	c, err := d.chainSet.Get(jb.CCIPBootstrapSpec.EVMChainID.ToInt())
	if err != nil {
		return nil, errors.Wrap(err, "unable to open chain")
	}
	// Bootstrap could either be an offramp or an executor, should work in both cases
	offRamp, err := single_token_offramp.NewSingleTokenOffRamp(jb.CCIPBootstrapSpec.ContractAddress.Address(), c.Client())
	if err != nil {
		return nil, errors.Wrap(err, "could not instantiate NewOffchainAggregator")
	}

	gormdb, errdb := d.db.DB()
	if errdb != nil {
		return nil, errors.Wrap(errdb, "unable to open sql db")
	}
	ocrdb := NewDB(gormdb, jb.CCIPBootstrapSpec.ContractAddress.Address())
	contractTracker := NewCCIPContractTracker(
		offrampTracker{offRamp},
		c.Client(),
		c.LogBroadcaster(),
		jb.ID,
		logger.Default,
		d.db,
		c,
		c.HeadBroadcaster(),
	)
	ocrLogger := logger.NewOCRWrapper(l, true, func(msg string) {
		d.jobORM.RecordError(context.Background(), jb.ID, msg)
	})
	offchainConfigDigester := evmutil.EVMOffchainConfigDigester{
		ChainID:         maybeRemapChainID(c.Config().ChainID()).Uint64(),
		ContractAddress: jb.CCIPBootstrapSpec.ContractAddress.Address(),
	}
	bootstrapNode, err := ocr.NewBootstrapper(ocr.BootstrapperArgs{
		BootstrapperFactory:   d.peerWrapper.Peer2,
		ContractConfigTracker: contractTracker,
		Database:              ocrdb,
		LocalConfig: computeLocalConfig(c.Config(), c.Config().Dev(),
			jb.CCIPBootstrapSpec.BlockchainTimeout.Duration(),
			jb.CCIPBootstrapSpec.ContractConfigConfirmations, jb.CCIPBootstrapSpec.ContractConfigTrackerPollInterval.Duration()),
		Logger:                 ocrLogger,
		MonitoringEndpoint:     nil, // TODO
		OffchainConfigDigester: offchainConfigDigester,
	})
	if err != nil {
		return nil, err
	}
	return []job.Service{contractTracker, bootstrapNode}, nil
}

func (d DelegateBootstrap) AfterJobCreated(spec job.Job) {
}

func (d DelegateBootstrap) BeforeJobDeleted(spec job.Job) {
}
