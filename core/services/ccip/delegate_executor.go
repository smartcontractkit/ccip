package ccip

import (
	"encoding/hex"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/services/offchainreporting2"
	"github.com/smartcontractkit/chainlink/core/services/telemetry"

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
	db                    *sqlx.DB
	jobORM                job.ORM
	ccipORM               ORM
	peerWrapper           *ocrcommon.SingletonPeerWrapper
	monitoringEndpointGen telemetry.MonitoringEndpointGenerator
	chainSet              evm.ChainSet
	cfg                   Config
	lggr                  logger.Logger
	ks                    keystore.OCR2
}

// TODO: Register this delegate behind a FF
func NewExecutionDelegate(
	db *sqlx.DB,
	jobORM job.ORM,
	ccipORM ORM,
	peerWrapper *ocrcommon.SingletonPeerWrapper,
	monitoringEndpointGen telemetry.MonitoringEndpointGenerator,
	chainSet evm.ChainSet,
	lggr logger.Logger,
	cfg Config,
	ks keystore.OCR2,
) *ExecutionDelegate {
	return &ExecutionDelegate{
		db:                    db,
		jobORM:                jobORM,
		ccipORM:               ccipORM,
		peerWrapper:           peerWrapper,
		monitoringEndpointGen: monitoringEndpointGen,
		chainSet:              chainSet,
		cfg:                   cfg,
		lggr:                  lggr,
		ks:                    ks,
	}
}

func (d ExecutionDelegate) JobType() job.Type {
	return job.CCIPExecution
}

func (d ExecutionDelegate) ServicesForSpec(jobSpec job.Job) (services []job.Service, err error) {
	spec := jobSpec.CCIPExecutionSpec
	if spec == nil {
		return nil, errors.Errorf("CCIPExecution expects a *job.CCIPExecutionSpec to be present, got %v", jobSpec)
	}
	if !spec.TransmitterID.Valid {
		return nil, errors.New("spec.TransmitterID not valid")
	}

	destChain, err := d.chainSet.Get(spec.DestEVMChainID.ToInt())
	if err != nil {
		return nil, errors.Wrap(err, "unable to open chain")
	}
	contract, err := message_executor.NewMessageExecutor(common.HexToAddress(spec.ExecutorID), destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "could not instantiate NewOffchainAggregator")
	}
	contractTracker := NewCCIPContractTracker(
		executorTracker{contract},
		destChain.Client(),
		destChain.LogBroadcaster(),
		jobSpec.ID,
		d.lggr,
		destChain,
		destChain.HeadBroadcaster(),
	)
	services = append(services, contractTracker)

	loggerWith := d.lggr.With(
		"OCRLogger", "true",
		"contractID", spec.ContractID,
		"jobName", jobSpec.Name.ValueOrZero(),
		"jobID", jobSpec.ID,
	)
	loggerWith.Infof("starting job with externalJobId %s, "+
		"offrampContract %s, onrampContract %s",
		jobSpec.ExternalJobID.String(),
		spec.OffRampID,
		spec.OnRampID,
	)

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
		contract,
		executorABI,
		NewExecuteTransmitter(destChain.TxManager(),
			d.db,
			spec.SourceEVMChainID.ToInt(),
			spec.DestEVMChainID.ToInt(),
			transmitterAddress,
			destChain.Config().EvmGasLimitDefault(),
			bulletprooftxmanager.NewQueueingTxStrategy(jobSpec.ExternalJobID,
				destChain.Config().OCRDefaultTransactionQueueDepth()),
			destChain.Client()),
		d.lggr,
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
	kb, err := d.ks.Get(kbID)
	if err != nil {
		return nil, err
	}

	bootstrapPeers, err := ocrcommon.GetValidatedBootstrapPeers(spec.P2PBootstrapPeers, d.peerWrapper.Config().P2PV2Bootstrappers())
	if err != nil {
		return nil, err
	}

	ocrdb := NewDB(d.db.DB, common.HexToAddress(spec.ExecutorID), d.lggr)
	lc := offchainreporting2.ToLocalConfig(destChain.Config(), spec.AsOCR2Spec())
	if err = ocr.SanityCheckLocalConfig(lc); err != nil {
		return nil, err
	}
	singleTokenOffRamp, err := single_token_offramp.NewSingleTokenOffRamp(common.HexToAddress(spec.OffRampID), destChain.Client())
	if err != nil {
		return nil, err
	}
	d.lggr.Infow("OCR2 job using local config",
		"BlockchainTimeout", lc.BlockchainTimeout,
		"ContractConfigConfirmations", lc.ContractConfigConfirmations,
		"ContractConfigTrackerPollInterval", lc.ContractConfigTrackerPollInterval,
		"ContractTransmitterTransmitTimeout", lc.ContractTransmitterTransmitTimeout,
		"DatabaseTimeout", lc.DatabaseTimeout,
	)

	oracle, err := ocr.NewOracle(ocr.OracleArgs{
		BinaryNetworkEndpointFactory: d.peerWrapper.Peer2,
		V2Bootstrappers:              bootstrapPeers,
		ContractTransmitter:          contractTransmitter,
		ContractConfigTracker:        contractTracker,
		Database:                     ocrdb,
		LocalConfig:                  lc,
		Logger:                       ocrLogger,
		MonitoringEndpoint:           d.monitoringEndpointGen.GenMonitoringEndpoint(spec.ContractID),
		OffchainConfigDigester: evmutil.EVMOffchainConfigDigester{
			ChainID:         maybeRemapChainID(destChain.Config().ChainID()).Uint64(),
			ContractAddress: common.HexToAddress(spec.ExecutorID),
		},
		OffchainKeyring: kb,
		OnchainKeyring:  kb,
		ReportingPluginFactory: NewExecutionReportingPluginFactory(
			loggerWith,
			d.ccipORM,
			spec.SourceEVMChainID.ToInt(),
			spec.DestEVMChainID.ToInt(),
			common.HexToAddress(spec.ExecutorID),
			singleTokenOffRamp),
	})
	if err != nil {
		return nil, err
	}
	services = append(services, oracle)

	sourceChain, err := d.chainSet.Get(spec.SourceEVMChainID.ToInt())
	if err != nil {
		return nil, errors.Wrap(err, "unable to open chain")
	}
	singleTokenOnRamp, err := single_token_onramp.NewSingleTokenOnRamp(common.HexToAddress(spec.OnRampID), sourceChain.Client())
	if err != nil {
		return nil, err
	}
	encodedCCIPConfig, err := contractTracker.GetOffchainConfig()
	if err != nil {
		return nil, errors.Wrap(err, "could not get the latest encoded config")
	}

	// TODO: Its conceivable we may want pull out this log listener into its own job spec so to avoid repeating
	// all the log subscriptions.
	logListener := NewLogListener(loggerWith,
		sourceChain.LogBroadcaster(),
		destChain.LogBroadcaster(),
		singleTokenOnRamp,
		singleTokenOffRamp,
		encodedCCIPConfig,
		d.ccipORM,
		jobSpec.ID)
	services = append(services, logListener)
	return services, nil
}

func (d ExecutionDelegate) AfterJobCreated(spec job.Job) {
}

func (d ExecutionDelegate) BeforeJobDeleted(spec job.Job) {
}
