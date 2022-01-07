package ccip

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_onramp"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/bulletprooftxmanager"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/keystore"
	"github.com/smartcontractkit/chainlink/core/services/ocrcommon"
	ocr "github.com/smartcontractkit/libocr/offchainreporting2"
	"github.com/smartcontractkit/libocr/offchainreporting2/chains/evmutil"
	"github.com/smartcontractkit/sqlx"
)

var _ job.Delegate = (*RelayDelegate)(nil)

type RelayDelegate struct {
	db          *sqlx.DB
	jobORM      job.ORM
	orm         ORM
	chainSet    evm.ChainSet
	keyStore    keystore.OCR2
	peerWrapper *ocrcommon.SingletonPeerWrapper
	lggr        logger.Logger
}

// TODO: Register this delegate behind a FF
func NewRelayDelegate(
	db *sqlx.DB,
	jobORM job.ORM,
	chainSet evm.ChainSet,
	keyStore keystore.OCR2,
	peerWrapper *ocrcommon.SingletonPeerWrapper,
	lggr logger.Logger,
) *RelayDelegate {
	return &RelayDelegate{
		db:          db,
		jobORM:      jobORM,
		orm:         NewORM(db),
		chainSet:    chainSet,
		keyStore:    keyStore,
		peerWrapper: peerWrapper,
		lggr:        lggr,
	}
}

func (d RelayDelegate) JobType() job.Type {
	return job.CCIPRelay
}

func (d RelayDelegate) getOracleArgs(l logger.Logger, jb job.Job, offRamp *single_token_offramp.SingleTokenOffRamp, chain evm.Chain, contractTracker *CCIPContractTracker, offchainConfigDigester evmutil.EVMOffchainConfigDigester) (*ocr.OracleArgs, error) {
	ta, err := getTransmitterAddress(jb.CCIPRelaySpec.TransmitterAddress, chain)
	if err != nil {
		return nil, err
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
			jb.CCIPRelaySpec.SourceEVMChainID.ToInt(),
			jb.CCIPRelaySpec.DestEVMChainID.ToInt(), ta.Address(),
			chain.Config().EvmGasLimitDefault(),
			bulletprooftxmanager.NewQueueingTxStrategy(jb.ExternalJobID,
				chain.Config().OCRDefaultTransactionQueueDepth(), false),
			chain.Client()),
		d.lggr,
	)
	ocrLogger := logger.NewOCRWrapper(l, true, func(msg string) {
		d.jobORM.RecordError(jb.ID, msg)
	})
	key, err := getValidatedKeyBundle(jb.CCIPRelaySpec.EncryptedOCRKeyBundleID, chain, d.keyStore)
	if err != nil {
		return nil, err
	}
	if err = validatePeerWrapper(jb.CCIPRelaySpec.P2PPeerID, chain, d.peerWrapper); err != nil {
		return nil, err
	}
	bootstrapPeers, err := getValidatedBootstrapPeers(jb.CCIPRelaySpec.P2PBootstrapPeers, chain)
	if err != nil {
		return nil, err
	}

	ocrdb := NewDB(d.db.DB, jb.CCIPRelaySpec.OffRampAddress.Address(), l)
	return &ocr.OracleArgs{
		BinaryNetworkEndpointFactory: d.peerWrapper.Peer2,
		V2Bootstrappers:              bootstrapPeers,
		ContractTransmitter:          contractTransmitter,
		ContractConfigTracker:        contractTracker,
		Database:                     ocrdb,
		LocalConfig: computeLocalConfig(chain.Config(), chain.Config().Dev(),
			jb.CCIPRelaySpec.BlockchainTimeout.Duration(),
			jb.CCIPRelaySpec.ContractConfigConfirmations, jb.CCIPRelaySpec.ContractConfigTrackerPollInterval.Duration()),
		Logger:                 ocrLogger,
		MonitoringEndpoint:     nil, // TODO
		OffchainConfigDigester: offchainConfigDigester,
		OffchainKeyring:        key,
		OnchainKeyring:         key,
		ReportingPluginFactory: NewRelayReportingPluginFactory(l, d.orm, offRamp),
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
