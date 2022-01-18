// TODO: I think we might be able to make the bootstrap job type just generic for all genocr jobs?
package ccip

import (
	"time"

	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/config"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/ocrcommon"
	ocrcommontypes "github.com/smartcontractkit/libocr/commontypes"
	ocr "github.com/smartcontractkit/libocr/offchainreporting2"
	"github.com/smartcontractkit/libocr/offchainreporting2/chains/evmutil"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/smartcontractkit/sqlx"
)

type DelegateBootstrap struct {
	bootstrappers []ocrcommontypes.BootstrapperLocator
	db            *sqlx.DB
	jobORM        job.ORM
	orm           ORM
	chainSet      evm.ChainSet
	peerWrapper   *ocrcommon.SingletonPeerWrapper
	lggr          logger.Logger
}

// TODO: Register this delegate behind a FF
func NewDelegateBootstrap(
	db *sqlx.DB,
	jobORM job.ORM,
	ccipORM ORM,
	chainSet evm.ChainSet,
	peerWrapper *ocrcommon.SingletonPeerWrapper,
	lggr logger.Logger,
) *DelegateBootstrap {
	return &DelegateBootstrap{
		db:          db,
		jobORM:      jobORM,
		orm:         ccipORM,
		chainSet:    chainSet,
		peerWrapper: peerWrapper,
		lggr:        lggr,
	}
}

func (d DelegateBootstrap) JobType() job.Type {
	return job.CCIPBootstrap
}

func (d DelegateBootstrap) ServicesForSpec(jb job.Job) ([]job.Service, error) {
	if jb.CCIPBootstrapSpec == nil {
		return nil, errors.New("no bootstrap job specified")
	}
	l := d.lggr.With(
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

	ocrdb := NewDB(d.db.DB, jb.CCIPBootstrapSpec.ContractAddress.Address(), d.lggr)
	contractTracker := NewCCIPContractTracker(
		offrampTracker{offRamp},
		c.Client(),
		c.LogBroadcaster(),
		jb.ID,
		d.lggr,
		c,
		c.HeadBroadcaster(),
	)
	ocrLogger := logger.NewOCRWrapper(l, true, func(msg string) {
		_ = d.jobORM.RecordError(jb.ID, msg)
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

// Fallback to config if explicit spec parameters are not set
func computeLocalConfig(config config.OCR2Config, dev bool, bt time.Duration, confs uint16, poll time.Duration) ocrtypes.LocalConfig {
	var blockchainTimeout time.Duration
	if bt != 0 {
		blockchainTimeout = bt
	} else {
		blockchainTimeout = config.OCR2BlockchainTimeout()
	}

	var contractConfirmations uint16
	if confs != 0 {
		contractConfirmations = confs
	} else {
		contractConfirmations = config.OCR2ContractConfirmations()
	}

	var contractConfigTrackerPollInterval time.Duration
	if poll != 0 {
		contractConfigTrackerPollInterval = poll
	} else {
		contractConfigTrackerPollInterval = config.OCR2ContractPollInterval()
	}

	lc := ocrtypes.LocalConfig{
		BlockchainTimeout:                  blockchainTimeout,
		ContractConfigConfirmations:        contractConfirmations,
		ContractConfigTrackerPollInterval:  contractConfigTrackerPollInterval,
		ContractTransmitterTransmitTimeout: config.OCR2ContractTransmitterTransmitTimeout(),
		DatabaseTimeout:                    config.OCR2DatabaseTimeout(),
	}
	if dev {
		// Skips config validation so we can use any config parameters we want.
		// For example to lower contractConfigTrackerPollInterval to speed up tests.
		lc.DevelopmentMode = ocrtypes.EnableDangerousDevelopmentMode
	}
	return lc
}
