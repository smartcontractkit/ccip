package ccip

import (
	"encoding/hex"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/services/offchainreporting2"

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
	cfg         Config
	ks          keystore.OCR2
	peerWrapper *ocrcommon.SingletonPeerWrapper
	lggr        logger.Logger
}

// TODO: Register this delegate behind a FF
func NewExecutionDelegate(
	db *sqlx.DB,
	jobORM job.ORM,
	chainSet evm.ChainSet,
	cfg Config,
	keyStore keystore.OCR2,
	peerWrapper *ocrcommon.SingletonPeerWrapper,
	lggr logger.Logger,

) *ExecutionDelegate {
	return &ExecutionDelegate{
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

func (d ExecutionDelegate) JobType() job.Type {
	return job.CCIPExecution
}

func (d ExecutionDelegate) getOracleArgs(l logger.Logger, jobSpec job.Job, executor *message_executor.MessageExecutor, chain evm.Chain, contractTracker *CCIPContractTracker, offchainConfigDigester evmutil.EVMOffchainConfigDigester, offRamp *single_token_offramp.SingleTokenOffRamp) (*ocr.OracleArgs, error) {
	spec := jobSpec.CCIPExecutionSpec
	if !spec.TransmitterID.Valid {
		return nil, errors.New("spec.TransmitterID not valid")
	}
	bytes, err := hex.DecodeString(strings.TrimPrefix(spec.TransmitterID.String, "0x"))
	if err != nil {
		return nil, errors.Wrap(err, "error parsing spec.TransmitterID ")
	}
	transmitterAddress := common.BytesToAddress(bytes)

	executorABI, err := abi.JSON(strings.NewReader(message_executor.MessageExecutorABI))
	if err != nil {
		return nil, errors.Wrap(err, "could not get contract ABI JSON")
	}
	contractTransmitter := NewExecutionTransmitter(
		executor,
		executorABI,
		NewExecuteTransmitter(chain.TxManager(),
			d.db,
			jobSpec.CCIPExecutionSpec.SourceEVMChainID.ToInt(),
			jobSpec.CCIPExecutionSpec.DestEVMChainID.ToInt(),
			transmitterAddress,
			chain.Config().EvmGasLimitDefault(),
			bulletprooftxmanager.NewQueueingTxStrategy(jobSpec.ExternalJobID,
				chain.Config().OCRDefaultTransactionQueueDepth()),
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

	bootstrapPeers, err := getValidatedBootstrapPeers(jobSpec.CCIPExecutionSpec.P2PBootstrapPeers, chain)
	if err != nil {
		return nil, err
	}

	ocrdb := NewDB(d.db.DB, jobSpec.CCIPExecutionSpec.ExecutorAddress.Address(), d.lggr)

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
		ReportingPluginFactory:       NewExecutionReportingPluginFactory(l, d.orm, jobSpec.CCIPExecutionSpec.SourceEVMChainID.ToInt(), jobSpec.CCIPExecutionSpec.DestEVMChainID.ToInt(), jobSpec.CCIPExecutionSpec.ExecutorAddress.Address(), offRamp),
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
