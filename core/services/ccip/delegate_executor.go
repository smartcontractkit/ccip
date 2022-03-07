package ccip

import (
	"encoding/hex"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	ocr "github.com/smartcontractkit/libocr/offchainreporting2"
	"github.com/smartcontractkit/sqlx"

	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/chains/evm/bulletprooftxmanager"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/message_executor"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_onramp"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/keystore"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/validate"
	"github.com/smartcontractkit/chainlink/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/core/services/pg"
	"github.com/smartcontractkit/chainlink/core/services/relay"
	"github.com/smartcontractkit/chainlink/core/services/relay/types"
	"github.com/smartcontractkit/chainlink/core/services/telemetry"
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
	relayer               types.RelayerCtx
}

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
	relayer types.RelayerCtx,
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
		relayer:               relayer,
	}
}

func (d ExecutionDelegate) JobType() job.Type {
	return job.CCIPExecution
}

func (d ExecutionDelegate) ServicesForSpec(jb job.Job) (services []job.ServiceCtx, err error) {
	spec := jb.CCIPExecutionSpec
	if spec == nil {
		return nil, errors.Errorf("CCIPExecution expects a *job.CCIPExecutionSpec to be present, got %v", jb)
	}
	if !spec.TransmitterID.Valid {
		return nil, errors.New("spec.TransmitterID not valid")
	}

	ocr2Provider, err := d.relayer.NewOCR2Provider(jb.ExternalJobID, &relay.OCR2ProviderArgs{
		ID:              spec.ID,
		ContractID:      spec.ContractID,
		TransmitterID:   spec.TransmitterID,
		Relay:           spec.Relay,
		RelayConfig:     spec.RelayConfig,
		IsBootstrapPeer: false,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error calling 'relayer.NewOCR2Provider'")
	}
	services = append(services, ocr2Provider)

	destChain, err := d.chainSet.Get(spec.DestEVMChainID.ToInt())
	if err != nil {
		return nil, errors.Wrap(err, "unable to open destination chain")
	}
	if !common.IsHexAddress(spec.ExecutorID) {
		return nil, errors.Wrap(err, "spec.ExecutorID is not a valid hex address")
	}
	contract, err := message_executor.NewMessageExecutor(common.HexToAddress(spec.ExecutorID), destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "could not instantiate NewOffchainAggregator")
	}

	lggr := d.lggr.With(
		"OCRLogger", "true",
		"contractID", spec.ContractID,
		"jobName", jb.Name.ValueOrZero(),
		"jobID", jb.ID,
	)
	lggr.Infof("starting job with externalJobId %s, "+
		"offrampContract %s, onrampContract %s",
		jb.ExternalJobID.String(),
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
			bulletprooftxmanager.NewQueueingTxStrategy(jb.ExternalJobID,
				destChain.Config().OCRDefaultTransactionQueueDepth()),
			destChain.Client()),
		d.lggr,
	)

	ocrLogger := logger.NewOCRWrapper(lggr, true, func(msg string) {
		d.lggr.ErrorIf(d.jobORM.RecordError(jb.ID, msg), "unable to record error")
	})

	// Fetch the specified OCR2 key bundle
	var kbID string
	if spec.OCRKeyBundleID.Valid {
		kbID = spec.OCRKeyBundleID.String
	} else if kbID, err = d.cfg.OCR2KeyBundleID(); err != nil {
		return nil, errors.Wrap(err, "error getting OCR2 key bundle id")
	}
	kb, err := d.ks.Get(kbID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get keys from key bundle")
	}

	bootstrapPeers, err := ocrcommon.GetValidatedBootstrapPeers(spec.P2PBootstrapPeers, d.peerWrapper.Config().P2PV2Bootstrappers())
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate bootstrap peers")
	}

	ocrdb := NewDB(d.db.DB, common.HexToAddress(spec.ExecutorID), d.lggr)
	lc := validate.ToLocalConfig(d.cfg, spec.AsOCR2Spec())
	if err = ocr.SanityCheckLocalConfig(lc); err != nil {
		return nil, errors.Wrap(err, "error while checking local config")
	}
	singleTokenOffRamp, err := single_token_offramp.NewSingleTokenOffRamp(common.HexToAddress(spec.OffRampID), destChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed creating a new offramp")
	}
	d.lggr.Infow("OCR2 job using local config",
		"BlockchainTimeout", lc.BlockchainTimeout,
		"ContractConfigConfirmations", lc.ContractConfigConfirmations,
		"ContractConfigTrackerPollInterval", lc.ContractConfigTrackerPollInterval,
		"ContractTransmitterTransmitTimeout", lc.ContractTransmitterTransmitTimeout,
		"DatabaseTimeout", lc.DatabaseTimeout,
	)

	tracker := ocr2Provider.ContractConfigTracker()
	oracle, err := ocr.NewOracle(ocr.OracleArgs{
		BinaryNetworkEndpointFactory: d.peerWrapper.Peer2,
		V2Bootstrappers:              bootstrapPeers,
		ContractTransmitter:          contractTransmitter,
		ContractConfigTracker:        tracker,
		Database:                     ocrdb,
		LocalConfig:                  lc,
		Logger:                       ocrLogger,
		MonitoringEndpoint:           d.monitoringEndpointGen.GenMonitoringEndpoint(spec.ContractID),
		OffchainConfigDigester:       ocr2Provider.OffchainConfigDigester(),
		OffchainKeyring:              kb,
		OnchainKeyring:               kb,
		ReportingPluginFactory: NewExecutionReportingPluginFactory(
			lggr,
			d.ccipORM,
			spec.SourceEVMChainID.ToInt(),
			spec.DestEVMChainID.ToInt(),
			common.HexToAddress(spec.ExecutorID),
			singleTokenOffRamp),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new oracle")
	}
	services = append(services, job.NewServiceAdapter(oracle))

	sourceChain, err := d.chainSet.Get(spec.SourceEVMChainID.ToInt())
	if err != nil {
		return nil, errors.Wrap(err, "unable to open destination chain")
	}
	singleTokenOnRamp, err := single_token_onramp.NewSingleTokenOnRamp(common.HexToAddress(spec.OnRampID), sourceChain.Client())
	if err != nil {
		return nil, errors.Wrap(err, "failed creating a new onramp")
	}
	ccipConfig, err := GetOffchainConfig(tracker)
	if err != nil {
		return nil, errors.Wrap(err, "could not get the latest encoded config")
	}

	// all the log subscriptions.
	logListener := NewLogListener(lggr,
		sourceChain.LogBroadcaster(),
		destChain.LogBroadcaster(),
		singleTokenOnRamp,
		singleTokenOffRamp,
		ccipConfig,
		d.ccipORM,
		jb.ID,
		pg.NewQ(d.db, d.lggr, d.cfg))
	services = append(services, logListener)
	return services, nil
}

func (d ExecutionDelegate) AfterJobCreated(spec job.Job) {
}

func (d ExecutionDelegate) BeforeJobDeleted(spec job.Job) {
}
