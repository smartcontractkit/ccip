package ccip

import (
	"encoding/hex"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/config"
	"github.com/smartcontractkit/chainlink/core/services/offchainreporting2"

	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/chains/evm/bulletprooftxmanager"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_onramp"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/keystore"
	"github.com/smartcontractkit/chainlink/core/services/ocrcommon"
	ocr "github.com/smartcontractkit/libocr/offchainreporting2"
	"github.com/smartcontractkit/libocr/offchainreporting2/chains/evmutil"
	"github.com/smartcontractkit/sqlx"
)

var _ job.Delegate = (*RelayDelegate)(nil)

type Config interface {
	config.OCR2Config
}

type RelayDelegate struct {
	db          *sqlx.DB
	jobORM      job.ORM
	orm         ORM
	chainSet    evm.ChainSet
	cfg         Config
	ks          keystore.OCR2
	peerWrapper *ocrcommon.SingletonPeerWrapper
	lggr        logger.Logger
}

// TODO: Register this delegate behind a FF
func NewRelayDelegate(
	db *sqlx.DB,
	jobORM job.ORM,
	chainSet evm.ChainSet,
	cfg Config,
	keyStore keystore.OCR2,
	peerWrapper *ocrcommon.SingletonPeerWrapper,
	lggr logger.Logger,
) *RelayDelegate {
	return &RelayDelegate{
		db:          db,
		jobORM:      jobORM,
		orm:         NewORM(db),
		chainSet:    chainSet,
		cfg:         cfg,
		ks:          keyStore,
		peerWrapper: peerWrapper,
		lggr:        lggr,
	}
}

func (d RelayDelegate) JobType() job.Type {
	return job.CCIPRelay
}

func (d RelayDelegate) getOracleArgs(l logger.Logger, jobSpec job.Job, offRamp *single_token_offramp.SingleTokenOffRamp, chain evm.Chain, contractTracker *CCIPContractTracker, offchainConfigDigester evmutil.EVMOffchainConfigDigester) (*ocr.OracleArgs, error) {
	spec := jobSpec.CCIPRelaySpec
	if !spec.TransmitterID.Valid {
		return nil, errors.New("spec.TransmitterID not valid")
	}
	bytes, err := hex.DecodeString(strings.TrimPrefix(spec.TransmitterID.String, "0x"))
	transmitterAddress := common.BytesToAddress(bytes)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing spec.TransmitterID ")
	}
	offRampABI, err := abi.JSON(strings.NewReader(single_token_offramp.SingleTokenOffRampABI))
	if err != nil {
		return nil, errors.Wrap(err, "could not get contract ABI JSON")
	}
	contractTransmitter := NewOfframpTransmitter(
		offRamp,
		offRampABI,
		NewRelayTransmitter(chain.TxManager(),
			d.db,
			jobSpec.CCIPRelaySpec.SourceEVMChainID.ToInt(),
			jobSpec.CCIPRelaySpec.DestEVMChainID.ToInt(),
			transmitterAddress,
			chain.Config().EvmGasLimitDefault(),
			bulletprooftxmanager.NewQueueingTxStrategy(jobSpec.ExternalJobID,
				chain.Config().OCRDefaultTransactionQueueDepth(), false),
			chain.Client()),
		d.lggr,
	)
	loggerWith := d.lggr.With(
		"OCRLogger", "true",
		"contractID", spec.ContractID,
		"jobName", jobSpec.Name.ValueOrZero(),
		"jobID", jobSpec.ID,
	)
	ocrLogger := logger.NewOCRWrapper(loggerWith, true, func(msg string) {
		d.lggr.ErrorIf(d.jobORM.RecordError(jobSpec.ID, msg), "unable to record error")
	})

	// Fetch the specified OCR2 key bundle
	var kbID string
	if spec.OCRKeyBundleID.Valid {
		kbID = spec.OCRKeyBundleID.String
	} else if kbID, err = d.cfg.OCR2KeyBundleID(); err != nil {
		return nil, err
	}
	key, err := d.ks.Get(kbID)
	if err != nil {
		return nil, err
	}

	bootstrapPeers, err := getValidatedBootstrapPeers(jobSpec.CCIPRelaySpec.P2PBootstrapPeers, chain)
	if err != nil {
		return nil, err
	}

	lc := offchainreporting2.ToLocalConfig(chain.Config(), spec.AsOCR2Spec())
	if err = ocr.SanityCheckLocalConfig(lc); err != nil {
		return nil, err
	}
	d.lggr.Infow("OCR2 job using local config",
		"BlockchainTimeout", lc.BlockchainTimeout,
		"ContractConfigConfirmations", lc.ContractConfigConfirmations,
		"ContractConfigTrackerPollInterval", lc.ContractConfigTrackerPollInterval,
		"ContractTransmitterTransmitTimeout", lc.ContractTransmitterTransmitTimeout,
		"DatabaseTimeout", lc.DatabaseTimeout,
	)

	ocrdb := NewDB(d.db.DB, jobSpec.CCIPRelaySpec.OffRampAddress.Address(), l)
	return &ocr.OracleArgs{
		BinaryNetworkEndpointFactory: d.peerWrapper.Peer2,
		V2Bootstrappers:              bootstrapPeers,
		ContractTransmitter:          contractTransmitter,
		ContractConfigTracker:        contractTracker,
		Database:                     ocrdb,
		LocalConfig:                  lc,
		Logger:                       ocrLogger,
		MonitoringEndpoint:           nil, // TODO
		OffchainConfigDigester:       offchainConfigDigester,
		OffchainKeyring:              key,
		OnchainKeyring:               key,
		ReportingPluginFactory:       NewRelayReportingPluginFactory(l, d.orm, offRamp),
	}, nil
}

func (d RelayDelegate) ServicesForSpec(jb job.Job) ([]job.Service, error) {
	if jb.CCIPRelaySpec == nil {
		return nil, errors.New("no ccip job specified")
	}
	l := d.lggr.With(
		"jobID", jb.ID,
		"externalJobID", jb.ExternalJobID,
		"offRampContract", jb.CCIPRelaySpec.OffRampAddress,
		"onRampContract", jb.CCIPRelaySpec.OnRampAddress,
	)

	destChain, err := d.chainSet.Get(jb.CCIPRelaySpec.DestEVMChainID.ToInt())
	if err != nil {
		return nil, errors.Wrap(err, "unable to open chain")
	}
	sourceChain, err := d.chainSet.Get(jb.CCIPRelaySpec.SourceEVMChainID.ToInt())
	if err != nil {
		return nil, errors.Wrap(err, "unable to open chain")
	}
	offRamp, err := single_token_offramp.NewSingleTokenOffRamp(jb.CCIPRelaySpec.OffRampAddress.Address(), destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "could not instantiate NewOffchainAggregator")
	}
	contractTracker := NewCCIPContractTracker(
		offrampTracker{offRamp},
		destChain.Client(),
		destChain.LogBroadcaster(),
		jb.ID,
		d.lggr,
		d.db,
		destChain,
		destChain.HeadBroadcaster(),
	)
	offchainConfigDigester := evmutil.EVMOffchainConfigDigester{
		ChainID:         maybeRemapChainID(destChain.Config().ChainID()).Uint64(),
		ContractAddress: jb.CCIPRelaySpec.OffRampAddress.Address(),
	}
	oracleArgs, err := d.getOracleArgs(l, jb, offRamp, destChain, contractTracker, offchainConfigDigester)
	if err != nil {
		return nil, err
	}
	oracle, err := ocr.NewOracle(*oracleArgs)
	if err != nil {
		return nil, err
	}
	singleTokenOnRamp, err := single_token_onramp.NewSingleTokenOnRamp(jb.CCIPRelaySpec.OnRampAddress.Address(), sourceChain.Client())
	if err != nil {
		return nil, err
	}

	encodedCCIPConfig, err := contractTracker.GetOffchainConfig()
	if err != nil {
		return nil, errors.Wrap(err, "could not get the latest encoded config")
	}
	// TODO: Its conceivable we may want pull out this log listener into its own job spec so to avoid repeating
	// All the log subscriptions
	logListener := NewLogListener(l,
		sourceChain.LogBroadcaster(),
		destChain.LogBroadcaster(),
		singleTokenOnRamp,
		offRamp,
		encodedCCIPConfig,
		d.db,
		jb.ID)
	return []job.Service{contractTracker, oracle, logListener}, nil
}

func (d RelayDelegate) AfterJobCreated(spec job.Job) {
}

func (d RelayDelegate) BeforeJobDeleted(spec job.Job) {
}
