package ccip

import (
	"fmt"
	"math/big"
	"reflect"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_onramp"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/job"
	"github.com/smartcontractkit/chainlink/core/services/log"
	"github.com/smartcontractkit/chainlink/core/services/postgres"
	"github.com/smartcontractkit/chainlink/core/utils"
	"github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	"gorm.io/gorm"
)

var (
	_ log.Listener = &LogListener{}
	_ job.Service  = &LogListener{}
)

type LogListener struct {
	utils.StartStopOnce

	logger                     logger.Logger
	sourceChainLogBroadcaster  log.Broadcaster
	destChainLogBroadcaster    log.Broadcaster
	singleTokenOnRamp          *single_token_onramp.SingleTokenOnRamp
	singleTokenOffRamp         *single_token_offramp.SingleTokenOffRamp
	sourceChainId, destChainId *big.Int
	// this can get overwritten by on-chain changes but doesn't need mutexes
	// because this is a single goroutine service.
	offchainConfig OffchainConfig
	jobID          int32

	unsubscribeLogsOnRamp  func()
	unsubscribeLogsOffRamp func()

	db  *gorm.DB
	orm ORM

	wgShutdown sync.WaitGroup
	mbLogs     *utils.Mailbox
	chStop     chan struct{}
}

func NewLogListener(
	l logger.Logger,
	sourceChainLogBroadcaster log.Broadcaster,
	destChainLogBroadcaster log.Broadcaster,
	singleTokenOnRamp *single_token_onramp.SingleTokenOnRamp,
	singleTokenOffRamp *single_token_offramp.SingleTokenOffRamp,
	offchainConfig OffchainConfig,
	db *gorm.DB,
	jobID int32,
) *LogListener {
	return &LogListener{
		logger:                    l,
		sourceChainLogBroadcaster: sourceChainLogBroadcaster,
		destChainLogBroadcaster:   destChainLogBroadcaster,
		jobID:                     jobID,
		db:                        db,
		orm:                       NewORM(postgres.UnwrapGormDB(db)),
		singleTokenOnRamp:         singleTokenOnRamp,
		singleTokenOffRamp:        singleTokenOffRamp,
		offchainConfig:            offchainConfig,
		mbLogs:                    utils.NewMailbox(10000),
		chStop:                    make(chan struct{}),
	}
}

// Start complies with job.Service
func (l *LogListener) Start() error {
	return l.StartOnce("CCIP_LogListener", func() error {
		sourceChainId, err := l.singleTokenOnRamp.CHAINID(nil)
		if err != nil {
			return err
		}
		destChainId, err := l.singleTokenOffRamp.CHAINID(nil)
		if err != nil {
			return err
		}
		l.sourceChainId = sourceChainId
		l.destChainId = destChainId
		l.subscribeSourceChainLogBroadcaster()
		l.subscribeDestChainLogBroadcaster()
		l.wgShutdown.Add(1)
		l.logger.Infow("CCIP_LogListener: Starting", "onRamp", l.singleTokenOnRamp.Address(), "offRamp", l.singleTokenOffRamp.Address())
		go l.run()

		return nil
	})
}

func (l *LogListener) subscribeSourceChainLogBroadcaster() {
	l.unsubscribeLogsOnRamp = l.sourceChainLogBroadcaster.Register(l, log.ListenerOpts{
		Contract: l.singleTokenOnRamp.Address(),
		LogsWithTopics: map[common.Hash][][]log.Topic{
			// Both relayer and executor save to db
			single_token_onramp.SingleTokenOnRampCrossChainSendRequested{}.Topic(): {},
		},
		ParseLog:         l.singleTokenOnRamp.ParseLog,
		NumConfirmations: uint64(l.offchainConfig.SourceIncomingConfirmations),
	})
}

func (l *LogListener) subscribeDestChainLogBroadcaster() {
	l.unsubscribeLogsOffRamp = l.destChainLogBroadcaster.Register(l, log.ListenerOpts{
		Contract: l.singleTokenOffRamp.Address(),
		LogsWithTopics: map[common.Hash][][]log.Topic{
			// Both relayer and executor mark as report_confirmed state
			single_token_offramp.SingleTokenOffRampReportAccepted{}.Topic(): {},
			// Both relayer and executor mark as execution_confirmed state
			single_token_offramp.SingleTokenOffRampCrossChainMessageExecuted{}.Topic(): {},
			// The offramp listens to config changed
			single_token_offramp.SingleTokenOffRampConfigSet{}.Topic(): {},
		},
		ParseLog:         l.singleTokenOffRamp.ParseLog,
		NumConfirmations: uint64(l.offchainConfig.DestIncomingConfirmations),
	})
}

// Close complies with job.Service
func (l *LogListener) Close() error {
	return l.StopOnce("CCIP_LogListener", func() error {
		close(l.chStop)
		l.wgShutdown.Wait()
		return nil
	})
}

func (l *LogListener) HandleLog(lb log.Broadcast) {
	wasOverCapacity := l.mbLogs.Deliver(lb)
	if wasOverCapacity {
		l.logger.Error("CCIP_LogListener: log mailbox is over capacity - dropped the oldest log")
	}
}

func (l *LogListener) run() {
	for {
		select {
		case <-l.chStop:
			l.unsubscribeLogsOffRamp()
			l.unsubscribeLogsOnRamp()
			l.wgShutdown.Done()
			return
		case <-l.mbLogs.Notify():
			l.handleReceivedLogs()
		}
	}
}

func (l *LogListener) handleReceivedLogs() {
	for {
		i, exists := l.mbLogs.Retrieve()
		if !exists {
			return
		}
		lb, ok := i.(log.Broadcast)
		if !ok {
			panic(errors.Errorf("CCIP_LogListener: invariant violation, expected log.Broadcast but got %T", lb))
		}

		logObj := lb.DecodedLog()
		if logObj == nil || reflect.ValueOf(logObj).IsNil() {
			l.logger.Error("CCIP_LogListener: HandleLog: ignoring nil value")
			return
		}

		// TODO: think about a way to do a single switch
		var logBroadcaster log.Broadcaster
		switch logObj.(type) {
		case *single_token_onramp.SingleTokenOnRampCrossChainSendRequested:
			logBroadcaster = l.sourceChainLogBroadcaster
		case *single_token_offramp.SingleTokenOffRampCrossChainMessageExecuted, *single_token_offramp.SingleTokenOffRampReportAccepted, *single_token_offramp.SingleTokenOffRampConfigSet:
			logBroadcaster = l.destChainLogBroadcaster
		default:
			l.logger.Warnf("CCIP_LogListener: unexpected log type %T", logObj)
		}

		ctx, cancel := postgres.DefaultQueryCtx()
		wasConsumed, err := logBroadcaster.WasAlreadyConsumed(l.db.WithContext(ctx), lb)
		cancel()
		if err != nil {
			l.logger.Errorw("CCIP_LogListener: could not determine if log was already consumed", "error", err)
			return
		} else if wasConsumed {
			return
		}

		switch log := logObj.(type) {
		case *single_token_onramp.SingleTokenOnRampCrossChainSendRequested:
			l.handleCrossChainSendRequested(log, lb)
		case *single_token_offramp.SingleTokenOffRampCrossChainMessageExecuted:
			l.handleCrossChainMessageExecuted(log, lb)
		case *single_token_offramp.SingleTokenOffRampReportAccepted:
			l.handleCrossChainReportRelayed(log, lb)
		case *single_token_offramp.SingleTokenOffRampConfigSet:
			if err := l.updateIncomingConfirmationsConfig(lb.RawLog()); err != nil {
				l.logger.Errorw("could not parse config set", "err", err)
			}
		default:
			l.logger.Warnf("CCIP_LogListener: unexpected log type %T", logObj)
		}
	}
}

func (l *LogListener) updateIncomingConfirmationsConfig(log types.Log) error {
	offrampConfigSet, err := l.singleTokenOffRamp.ParseConfigSet(log)
	if err != nil {
		return err
	}
	contractConfig := ContractConfigFromConfigSetEvent(ConfigSet(*offrampConfigSet))
	publicConfig, err := confighelper.PublicConfigFromContractConfig(false, contractConfig)
	if err != nil {
		return err
	}
	ccipConfig, err := Decode(publicConfig.ReportingPluginConfig)
	if err != nil {
		return err
	}
	if l.offchainConfig.SourceIncomingConfirmations != ccipConfig.SourceIncomingConfirmations {
		l.offchainConfig.SourceIncomingConfirmations = ccipConfig.SourceIncomingConfirmations
		l.unsubscribeLogsOnRamp()
		l.subscribeSourceChainLogBroadcaster()
	}

	if l.offchainConfig.DestIncomingConfirmations != ccipConfig.DestIncomingConfirmations {
		l.offchainConfig.DestIncomingConfirmations = ccipConfig.DestIncomingConfirmations
		l.unsubscribeLogsOffRamp()
		l.subscribeDestChainLogBroadcaster()
	}
	return nil
}

func (l *LogListener) handleCrossChainMessageExecuted(executed *single_token_offramp.SingleTokenOffRampCrossChainMessageExecuted, lb log.Broadcast) {
	l.logger.Infow("CCIP_LogListener: cross chain request executed",
		"seqNum", fmt.Sprintf("%0x", executed.SequenceNumber),
		"jobID", lb.JobID(),
	)
	err := l.orm.UpdateRequestStatus(l.sourceChainId, l.destChainId, executed.SequenceNumber, executed.SequenceNumber, RequestStatusExecutionConfirmed)
	if err != nil {
		// We can replay the logs if needed
		l.logger.Errorw("failed to save CCIP request", "error", err)
		return
	}
	ctx, cancel := postgres.DefaultQueryCtx()
	defer cancel()
	if err := l.destChainLogBroadcaster.MarkConsumed(l.db.WithContext(ctx), lb); err != nil {
		l.logger.Errorw("CCIP_LogListener: failed mark consumed", "err", err)
	}
}

func (l *LogListener) handleCrossChainReportRelayed(relayed *single_token_offramp.SingleTokenOffRampReportAccepted, lb log.Broadcast) {
	l.logger.Infow("CCIP_LogListener: cross chain report relayed",
		"minSeqNum", fmt.Sprintf("%0x", relayed.Report.MinSequenceNumber),
		"maxSeqNum", fmt.Sprintf("%0x", relayed.Report.MaxSequenceNumber),
		"jobID", lb.JobID(),
	)

	// TODO: should be in the same tx
	err := l.orm.UpdateRequestStatus(l.sourceChainId, l.destChainId, relayed.Report.MinSequenceNumber, relayed.Report.MaxSequenceNumber, RequestStatusRelayConfirmed)
	if err != nil {
		// We can replay the logs if needed
		l.logger.Errorw("failed to save CCIP request", "error", err)
		return
	}
	err = l.orm.SaveRelayReport(RelayReport{
		Root:      relayed.Report.MerkleRoot[:],
		MinSeqNum: *utils.NewBig(relayed.Report.MinSequenceNumber),
		MaxSeqNum: *utils.NewBig(relayed.Report.MaxSequenceNumber),
	})
	if err != nil {
		// We can replay the logs if needed
		l.logger.Errorw("failed to save CCIP report", "error", err)
		return
	}
	ctx, cancel := postgres.DefaultQueryCtx()
	defer cancel()
	if err := l.destChainLogBroadcaster.MarkConsumed(l.db.WithContext(ctx), lb); err != nil {
		l.logger.Errorw("CCIP_LogListener: failed mark consumed", "err", err)
	}
}

// We assume a bounded Message size which is enforced on-chain,
// TODO: add Message bounds to onramp and include assertion offchain as well.
func (l *LogListener) handleCrossChainSendRequested(request *single_token_onramp.SingleTokenOnRampCrossChainSendRequested, lb log.Broadcast) {
	l.logger.Infow("CCIP_LogListener: cross chain send request received",
		"requestId", fmt.Sprintf("%0x", request.Message.SequenceNumber),
		"sender", request.Message.Sender,
		"receiver", request.Message.Payload.Receiver,
		"sourceChainId", request.Message.SourceChainId,
		"destChainId", request.Message.DestinationChainId,
		"tokens", request.Message.Payload.Tokens,
		"amounts", request.Message.Payload.Amounts,
		"options", request.Message.Payload.Options,
		"jobID", lb.JobID(),
	)

	var tokens []string
	for _, token := range request.Message.Payload.Tokens {
		tokens = append(tokens, token.String())
	}
	var amounts []string
	for _, amount := range request.Message.Payload.Amounts {
		amounts = append(amounts, amount.String())
	}
	err := l.orm.SaveRequest(&Request{
		SeqNum:        *utils.NewBig(request.Message.SequenceNumber),
		SourceChainID: request.Message.SourceChainId.String(),
		DestChainID:   request.Message.DestinationChainId.String(),
		Sender:        request.Message.Sender,
		Receiver:      request.Message.Payload.Receiver,
		Data:          request.Message.Payload.Data,
		Tokens:        tokens,
		Amounts:       amounts,
		Executor:      request.Message.Payload.Executor,
		Options:       request.Message.Payload.Options,
		Raw:           request.Raw.Data,
		Status:        RequestStatusUnstarted,
	})
	if err != nil {
		// We can replay the logs if needed
		l.logger.Errorw("failed to save CCIP request", "error", err)
		return
	}

	ctx, cancel := postgres.DefaultQueryCtx()
	defer cancel()
	if err := l.sourceChainLogBroadcaster.MarkConsumed(l.db.WithContext(ctx), lb); err != nil {
		l.logger.Errorw("CCIP_LogListener: failed mark consumed", "err", err)
	}
}

// JobID complies with log.Listener
func (l *LogListener) JobID() int32 {
	return l.jobID
}
