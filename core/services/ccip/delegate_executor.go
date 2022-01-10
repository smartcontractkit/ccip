package ccip

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/chains/evm/bulletprooftxmanager"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/message_executor"
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

var _ job.Delegate = (*ExecutionDelegate)(nil)

type ExecutionDelegate struct {
	db          *sqlx.DB
	jobORM      job.ORM
	orm         ORM
	chainSet    evm.ChainSet
	ks          keystore.OCR2
	peerWrapper *ocrcommon.SingletonPeerWrapper
	lggr        logger.Logger
}

// TODO: Register this delegate behind a FF
func NewExecutionDelegate(
	db *sqlx.DB,
	jobORM job.ORM,
	chainSet evm.ChainSet,
	keyStore keystore.OCR2,
	peerWrapper *ocrcommon.SingletonPeerWrapper,
	lggr logger.Logger,

) *ExecutionDelegate {
	return &ExecutionDelegate{
		db:          db,
		jobORM:      jobORM,
		orm:         NewORM(db),
		chainSet:    chainSet,
		ks:          keyStore,
		peerWrapper: peerWrapper,
		lggr:        lggr,
	}
}

func (d ExecutionDelegate) JobType() job.Type {
	return job.CCIPExecution
}

func (d ExecutionDelegate) getOracleArgs(l logger.Logger, jb job.Job, executor *message_executor.MessageExecutor, chain evm.Chain, contractTracker *CCIPContractTracker, offchainConfigDigester evmutil.EVMOffchainConfigDigester, offRamp *single_token_offramp.SingleTokenOffRamp) (*ocr.OracleArgs, error) {
	ta, err := getTransmitterAddress(jb.CCIPExecutionSpec.TransmitterAddress, chain)
	if err != nil {
		return nil, err
	}
	executorABI, err := abi.JSON(strings.NewReader(message_executor.MessageExecutorABI))
	if err != nil {
		return nil, errors.Wrap(err, "could not get contract ABI JSON")
	}
	contractTransmitter := NewExecutionTransmitter(
		executor,
		executorABI,
		NewExecuteTransmitter(chain.TxManager(),
			d.db,
			jb.CCIPExecutionSpec.SourceEVMChainID.ToInt(),
			jb.CCIPExecutionSpec.DestEVMChainID.ToInt(), ta.Address(),
			chain.Config().EvmGasLimitDefault(),
			bulletprooftxmanager.NewQueueingTxStrategy(jb.ExternalJobID,
				chain.Config().OCRDefaultTransactionQueueDepth(), false),
			chain.Client()),
		d.lggr,
	)
	ocrLogger := logger.NewOCRWrapper(l, true, func(msg string) {
		_ = d.jobORM.RecordError(jb.ID, msg)
	})
	key, err := getValidatedKeyBundle(jb.CCIPExecutionSpec.EncryptedOCRKeyBundleID, chain, d.ks)
	if err != nil {
		return nil, err
	}
	if err = validatePeerWrapper(jb.CCIPExecutionSpec.P2PPeerID, chain, d.peerWrapper); err != nil {
		return nil, err
	}
	bootstrapPeers, err := getValidatedBootstrapPeers(jb.CCIPExecutionSpec.P2PBootstrapPeers, chain)
	if err != nil {
		return nil, err
	}

	ocrdb := NewDB(d.db.DB, jb.CCIPExecutionSpec.ExecutorAddress.Address(), d.lggr)
	return &ocr.OracleArgs{
		BinaryNetworkEndpointFactory: d.peerWrapper.Peer2,
		V2Bootstrappers:              bootstrapPeers,
		ContractTransmitter:          contractTransmitter,
		ContractConfigTracker:        contractTracker,
		Database:                     ocrdb,
		LocalConfig: computeLocalConfig(chain.Config(), chain.Config().Dev(),
			jb.CCIPExecutionSpec.BlockchainTimeout.Duration(),
			jb.CCIPExecutionSpec.ContractConfigConfirmations, jb.CCIPExecutionSpec.ContractConfigTrackerPollInterval.Duration()),
		Logger:                 ocrLogger,
		MonitoringEndpoint:     nil, // TODO
		OffchainConfigDigester: offchainConfigDigester,
		OffchainKeyring:        key,
		OnchainKeyring:         key,
		ReportingPluginFactory: NewExecutionReportingPluginFactory(l, d.orm, jb.CCIPExecutionSpec.SourceEVMChainID.ToInt(), jb.CCIPExecutionSpec.DestEVMChainID.ToInt(), jb.CCIPExecutionSpec.ExecutorAddress.Address(), offRamp),
	}, nil
}

func (d ExecutionDelegate) ServicesForSpec(jb job.Job) ([]job.Service, error) {
	if jb.CCIPExecutionSpec == nil {
		return nil, errors.New("no ccip job specified")
	}
	l := d.lggr.With(
		"jobID", jb.ID,
		"externalJobID", jb.ExternalJobID,
		"offRampAddress", jb.CCIPExecutionSpec.OffRampAddress,
		"onRampAddress", jb.CCIPExecutionSpec.OnRampAddress,
		"executorAddress", jb.CCIPExecutionSpec.OnRampAddress,
	)

	destChain, err := d.chainSet.Get(jb.CCIPExecutionSpec.DestEVMChainID.ToInt())
	if err != nil {
		return nil, errors.Wrap(err, "unable to open chain")
	}
	sourceChain, err := d.chainSet.Get(jb.CCIPExecutionSpec.SourceEVMChainID.ToInt())
	if err != nil {
		return nil, errors.Wrap(err, "unable to open chain")
	}
	contract, err := message_executor.NewMessageExecutor(jb.CCIPExecutionSpec.ExecutorAddress.Address(), destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "could not instantiate NewOffchainAggregator")
	}
	singleTokenOffRamp, err := single_token_offramp.NewSingleTokenOffRamp(jb.CCIPExecutionSpec.OffRampAddress.Address(), destChain.Client())
	if err != nil {
		return nil, err
	}
	offchainConfigDigester := evmutil.EVMOffchainConfigDigester{
		ChainID:         maybeRemapChainID(destChain.Config().ChainID()).Uint64(),
		ContractAddress: jb.CCIPExecutionSpec.ExecutorAddress.Address(),
	}
	contractTracker := NewCCIPContractTracker(
		executorTracker{contract},
		destChain.Client(),
		destChain.LogBroadcaster(),
		jb.ID,
		d.lggr,
		d.db,
		destChain,
		destChain.HeadBroadcaster(),
	)

	oracleArgs, err := d.getOracleArgs(l, jb, contract, destChain, contractTracker, offchainConfigDigester, singleTokenOffRamp)
	if err != nil {
		return nil, err
	}
	oracle, err := ocr.NewOracle(*oracleArgs)
	if err != nil {
		return nil, err
	}

	singleTokenOnRamp, err := single_token_onramp.NewSingleTokenOnRamp(jb.CCIPExecutionSpec.OnRampAddress.Address(), sourceChain.Client())
	if err != nil {
		return nil, err
	}

	encodedCCIPConfig, err := contractTracker.GetOffchainConfig()
	if err != nil {
		return nil, errors.Wrap(err, "could not get the latest encoded config")
	}
	// TODO: Its conceivable we may want pull out this log listener into its own job spec so to avoid repeating
	// all the log subscriptions.
	logListener := NewLogListener(l,
		sourceChain.LogBroadcaster(),
		destChain.LogBroadcaster(),
		singleTokenOnRamp,
		singleTokenOffRamp,
		encodedCCIPConfig,
		d.db,
		jb.ID)
	return []job.Service{contractTracker, oracle, logListener}, nil
}

func (d ExecutionDelegate) AfterJobCreated(spec job.Job) {
}

func (d ExecutionDelegate) BeforeJobDeleted(spec job.Job) {
}
