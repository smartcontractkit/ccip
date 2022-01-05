// Note lifted from https://github.com/smartcontractkit/chainlink/pull/4809/files
// TODO: Pull into common ocr library
package ccip

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink/core/chains"
	"github.com/smartcontractkit/chainlink/core/chains/evm"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/message_executor"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/eth"
	httypes "github.com/smartcontractkit/chainlink/core/services/headtracker/types"
	"github.com/smartcontractkit/chainlink/core/services/log"
	"github.com/smartcontractkit/chainlink/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/core/services/postgres"
	"github.com/smartcontractkit/chainlink/core/utils"
	"github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2/types"

	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

var (
	_ ocrtypes.ContractConfigTracker = &CCIPContractTracker{}
	_ httypes.HeadTrackable          = &CCIPContractTracker{}

	OCRContractConfigSet = getEventTopic("ConfigSet")
)

type LatestConfigDetails struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}

type ConfigSet struct {
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   []gethCommon.Address
	Transmitters              []gethCommon.Address
	F                         uint8
	OnchainConfig             []byte
	EncodedConfigVersion      uint64
	Encoded                   []byte
	Raw                       gethTypes.Log
}

type OffchainConfig struct {
	SourceIncomingConfirmations uint32
	DestIncomingConfirmations   uint32
}

type OCR2 interface {
	Address() gethCommon.Address
	ParseLog(log gethTypes.Log) (generated.AbigenLog, error)
	LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails, error)
	ParseConfigSet(log gethTypes.Log) (ConfigSet, error)
}

type offrampTracker struct {
	*single_token_offramp.SingleTokenOffRamp
}

func (ot offrampTracker) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails, error) {
	dets, err := ot.SingleTokenOffRamp.LatestConfigDetails(opts)
	return LatestConfigDetails(dets), err
}

func (ot offrampTracker) ParseConfigSet(log gethTypes.Log) (ConfigSet, error) {
	c, err := ot.SingleTokenOffRamp.ParseConfigSet(log)
	return ConfigSet(*c), err
}

type executorTracker struct {
	*message_executor.MessageExecutor
}

func (ot executorTracker) LatestConfigDetails(opts *bind.CallOpts) (LatestConfigDetails, error) {
	dets, err := ot.MessageExecutor.LatestConfigDetails(opts)
	return LatestConfigDetails(dets), err
}

func (ot executorTracker) ParseConfigSet(log gethTypes.Log) (ConfigSet, error) {
	c, err := ot.MessageExecutor.ParseConfigSet(log)
	return ConfigSet(*c), err
}

// CCIPContractTracker complies with ContractConfigTracker interface and
// handles log events related to the contract more generally
type CCIPContractTracker struct {
	utils.StartStopOnce

	ethClient       eth.Client
	contract        OCR2
	logBroadcaster  log.Broadcaster
	jobID           int32
	logger          logger.Logger
	gdb             *gorm.DB
	blockTranslator ocrcommon.BlockTranslator
	chain           evm.Chain

	// HeadBroadcaster
	headBroadcaster  httypes.HeadBroadcaster
	unsubscribeHeads func()

	// Start/Stop lifecycle
	ctx             context.Context
	ctxCancel       context.CancelFunc
	wg              sync.WaitGroup
	unsubscribeLogs func()

	// ContractConfig
	configsMB utils.Mailbox
	chConfigs chan ocrtypes.ContractConfig

	// LatestBlockHeight
	latestBlockHeight   int64
	latestBlockHeightMu sync.RWMutex
}

// NewCCIPContractTracker makes a new CCIPContractTracker
func NewCCIPContractTracker(
	contract OCR2,
	ethClient eth.Client,
	logBroadcaster log.Broadcaster,
	jobID int32,
	logger logger.Logger,
	gdb *gorm.DB,
	chain evm.Chain,
	headBroadcaster httypes.HeadBroadcaster,
) (o *CCIPContractTracker) {
	ctx, cancel := context.WithCancel(context.Background())
	return &CCIPContractTracker{
		utils.StartStopOnce{},
		ethClient,
		contract,
		logBroadcaster,
		jobID,
		logger,
		gdb,
		ocrcommon.NewBlockTranslator(chain.Config(), ethClient, logger),
		chain,
		headBroadcaster,
		nil,
		ctx,
		cancel,
		sync.WaitGroup{},
		nil,
		// Should only ever be 0 or 1 config in the mailbox, put a sanity bound of 100.
		*utils.NewMailbox(100),
		make(chan ocrtypes.ContractConfig),
		-1,
		sync.RWMutex{},
	}
}

// Start must be called before logs can be delivered
// It ought to be called before starting OCR
func (t *CCIPContractTracker) Start() error {
	return t.StartOnce("CCIPContractTracker", func() (err error) {
		t.logger.Infow("CCIPContractTracker: registering for config set logs")
		t.unsubscribeLogs = t.logBroadcaster.Register(t, log.ListenerOpts{
			Contract: t.contract.Address(),
			ParseLog: t.contract.ParseLog,
			LogsWithTopics: map[gethCommon.Hash][][]log.Topic{
				single_token_offramp.SingleTokenOffRampConfigSet{}.Topic(): nil,
			},
			NumConfirmations: 1,
		})

		var latestHead *eth.Head
		latestHead, t.unsubscribeHeads = t.headBroadcaster.Subscribe(t)
		if latestHead != nil {
			t.setLatestBlockHeight(*latestHead)
		}

		t.wg.Add(1)
		go t.processLogs()
		return nil
	})
}

// Close should be called after teardown of the OCR job relying on this tracker
func (t *CCIPContractTracker) Close() error {
	return t.StopOnce("CCIPContractTracker", func() error {
		t.ctxCancel()
		t.wg.Wait()
		t.unsubscribeHeads()
		t.unsubscribeLogs()
		close(t.chConfigs)
		return nil
	})
}

// Connect conforms to HeadTrackable
func (t *CCIPContractTracker) Connect(*eth.Head) error { return nil }

// OnNewLongestChain conformed to HeadTrackable and updates latestBlockHeight
func (t *CCIPContractTracker) OnNewLongestChain(_ context.Context, h eth.Head) {
	t.setLatestBlockHeight(h)
}

func (t *CCIPContractTracker) setLatestBlockHeight(h eth.Head) {
	var num int64
	if h.L1BlockNumber.Valid {
		num = h.L1BlockNumber.Int64
	} else {
		num = h.Number
	}
	t.latestBlockHeightMu.Lock()
	defer t.latestBlockHeightMu.Unlock()
	if num > t.latestBlockHeight {
		t.latestBlockHeight = num
	}
}

func (t *CCIPContractTracker) getLatestBlockHeight() int64 {
	t.latestBlockHeightMu.RLock()
	defer t.latestBlockHeightMu.RUnlock()
	return t.latestBlockHeight
}

func (t *CCIPContractTracker) processLogs() {
	defer t.wg.Done()
	for {
		select {
		case <-t.configsMB.Notify():
			// NOTE: libocr could take an arbitrary amount of time to process a
			// new config. To avoid blocking the log broadcaster, we use this
			// background thread to deliver them and a mailbox as the buffer.
			for {
				x, exists := t.configsMB.Retrieve()
				if !exists {
					break
				}
				cc, ok := x.(ocrtypes.ContractConfig)
				if !ok {
					panic(fmt.Sprintf("expected ocrtypes.ContractConfig but got %T", x))
				}
				select {
				case t.chConfigs <- cc:
				case <-t.ctx.Done():
					return
				}
			}
		case <-t.ctx.Done():
			return
		}
	}
}

// HandleLog complies with LogListener interface
// It is not thread safe
func (t *CCIPContractTracker) HandleLog(lb log.Broadcast) {
	t.logger.Infow("CCIPContractTracker: config set log received", "log", lb.String())
	was, err := t.logBroadcaster.WasAlreadyConsumed(t.gdb, lb)
	if err != nil {
		t.logger.Errorw("OCRContract: could not determine if log was already consumed", "error", err)
		return
	} else if was {
		return
	}

	raw := lb.RawLog()
	if raw.Address != t.contract.Address() {
		t.logger.Errorf("log address of 0x%x does not match configured contract address of 0x%x", raw.Address, t.contract.Address())
		t.logger.ErrorIfCalling(func() error { return t.logBroadcaster.MarkConsumed(t.gdb, lb) })
		return
	}
	topics := raw.Topics
	if len(topics) == 0 {
		t.logger.ErrorIfCalling(func() error { return t.logBroadcaster.MarkConsumed(t.gdb, lb) })
		return
	}

	var consumed bool
	switch topics[0] {
	case single_token_offramp.SingleTokenOffRampConfigSet{}.Topic():
		configSet, err := t.contract.ParseConfigSet(raw)
		if err != nil {
			t.logger.Errorw("could not parse config set", "err", err)
			t.logger.ErrorIfCalling(func() error { return t.logBroadcaster.MarkConsumed(t.gdb, lb) })
			return
		}
		configSet.Raw = raw
		cc := ContractConfigFromConfigSetEvent(configSet)

		wasOverCapacity := t.configsMB.Deliver(cc)
		if wasOverCapacity {
			t.logger.Error("config mailbox is over capacity - dropped the oldest unprocessed item")
		}
	default:
		logger.Debugw("CCIPContractTracker: got unrecognised log topic", "topic", topics[0])
	}
	if !consumed {
		ctx, cancel := postgres.DefaultQueryCtx()
		defer cancel()
		t.logger.ErrorIfCalling(func() error { return t.logBroadcaster.MarkConsumed(t.gdb.WithContext(ctx), lb) })
	}
}

// IsLaterThan returns true if the first log was emitted "after" the second log
// from the blockchain's point of view
func IsLaterThan(incoming gethTypes.Log, existing gethTypes.Log) bool {
	return incoming.BlockNumber > existing.BlockNumber ||
		(incoming.BlockNumber == existing.BlockNumber && incoming.TxIndex > existing.TxIndex) ||
		(incoming.BlockNumber == existing.BlockNumber && incoming.TxIndex == existing.TxIndex && incoming.Index > existing.Index)
}

// IsV2Job complies with LogListener interface
func (t *CCIPContractTracker) IsV2Job() bool {
	return true
}

// JobID complies with LogListener interface
func (t *CCIPContractTracker) JobID() int32 {
	return t.jobID
}

// Notify returns a channel that can wake up the contract tracker to let it
// know when a new config is available
func (t *CCIPContractTracker) Notify() <-chan struct{} {
	return nil
}

// LatestConfigDetails queries the eth node
func (t *CCIPContractTracker) LatestConfigDetails(ctx context.Context) (changedInBlock uint64, configDigest ocrtypes.ConfigDigest, err error) {
	var cancel context.CancelFunc
	ctx, cancel = utils.CombinedContext(t.ctx, ctx)
	defer cancel()

	opts := bind.CallOpts{Context: ctx, Pending: false}
	result, err := t.contract.LatestConfigDetails(&opts)
	if err != nil {
		return 0, configDigest, errors.Wrap(err, "error getting LatestConfigDetails")
	}

	t.logger.Infow("CCIPContractTracker: latest config details", "digest", hexutil.Encode(result.ConfigDigest[:]))
	configDigest, err = ocrtypes.BytesToConfigDigest(result.ConfigDigest[:])
	if err != nil {
		return 0, configDigest, errors.Wrap(err, fmt.Sprintf("error getting LatestConfigDetails %v", t.contract.Address()))
	}
	return uint64(result.BlockNumber), configDigest, err
}

// Return the latest configuration
func (t *CCIPContractTracker) LatestConfig(ctx context.Context, changedInBlock uint64) (ocrtypes.ContractConfig, error) {
	fromBlock, toBlock := t.blockTranslator.NumberToQueryRange(ctx, changedInBlock)
	q := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: []gethCommon.Address{t.contract.Address()},
		Topics: [][]gethCommon.Hash{
			{OCRContractConfigSet},
		},
	}
	var cancel context.CancelFunc
	ctx, cancel = utils.CombinedContext(t.ctx, ctx)
	defer cancel()

	logs, err := t.ethClient.FilterLogs(ctx, q)
	if err != nil {
		return ocrtypes.ContractConfig{}, err
	}
	if len(logs) == 0 {
		return ocrtypes.ContractConfig{}, errors.Errorf("ConfigFromLogs: OCRContract with address 0x%x has no logs", t.contract.Address())
	}

	latest, err := t.contract.ParseConfigSet(logs[len(logs)-1])
	if err != nil {
		return ocrtypes.ContractConfig{}, errors.Wrap(err, "ConfigFromLogs failed to ParseConfigSet")
	}
	latest.Raw = logs[len(logs)-1]
	if latest.Raw.Address != t.contract.Address() {
		return ocrtypes.ContractConfig{}, errors.Errorf("log address of 0x%x does not match configured contract address of 0x%x", latest.Raw.Address, t.contract.Address())
	}

	cc := ContractConfigFromConfigSetEvent(latest)
	t.logger.Infow("CCIPContractTracker: latest config", "digest", hexutil.Encode(cc.ConfigDigest[:]))
	return cc, err
}

func ContractConfigFromConfigSetEvent(changed ConfigSet) ocrtypes.ContractConfig {
	var transmitAccounts []ocrtypes.Account
	for _, addr := range changed.Transmitters {
		transmitAccounts = append(transmitAccounts, ocrtypes.Account(addr.Hex()))
	}
	var signers []ocrtypes.OnchainPublicKey
	for _, addr := range changed.Signers {
		addr := addr
		signers = append(signers, addr[:])
	}
	return ocrtypes.ContractConfig{
		ConfigDigest:          changed.ConfigDigest,
		ConfigCount:           changed.ConfigCount,
		Signers:               signers,
		Transmitters:          transmitAccounts,
		F:                     changed.F,
		OnchainConfig:         changed.OnchainConfig,
		OffchainConfigVersion: changed.EncodedConfigVersion,
		OffchainConfig:        changed.Encoded,
	}
}

// LatestBlockHeight queries the eth node for the most recent header
func (t *CCIPContractTracker) LatestBlockHeight(ctx context.Context) (blockheight uint64, err error) {
	// We skip confirmation checking anyway on Optimism so there's no need to
	// care about the block height; we have no way of getting the L1 block
	// height anyway
	if t.chain.Config().ChainType() == chains.Optimism {
		return 0, nil
	}
	latestBlockHeight := t.getLatestBlockHeight()
	if latestBlockHeight >= 0 {
		return uint64(latestBlockHeight), nil
	}

	t.logger.Debugw("CCIPContractTracker: still waiting for first head, falling back to on-chain lookup")

	var cancel context.CancelFunc
	ctx, cancel = utils.CombinedContext(t.ctx, ctx)
	defer cancel()

	h, err := t.ethClient.HeadByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	if h == nil {
		return 0, errors.New("got nil head")
	}

	if h.L1BlockNumber.Valid {
		return uint64(h.L1BlockNumber.Int64), nil
	}

	return uint64(h.Number), nil
}

func (t *CCIPContractTracker) GetOffchainConfig() (OffchainConfig, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	changedInBlock, _, err := t.LatestConfigDetails(ctx)
	if err != nil {
		return OffchainConfig{}, errors.Wrap(err, "could not get block number for latest config change")
	}
	config, err := t.LatestConfig(ctx, changedInBlock)
	if err != nil {
		return OffchainConfig{}, errors.Wrap(err, "could not get latest config")
	}
	publicConfig, err := confighelper.PublicConfigFromContractConfig(false, config)
	if err != nil {
		return OffchainConfig{}, errors.Wrap(err, "could not parse latest config")
	}
	ccipConfig, err := Decode(publicConfig.ReportingPluginConfig)
	if err != nil {
		return OffchainConfig{}, errors.Wrap(err, "could not decode latest config")
	}
	return ccipConfig, nil
}

func Decode(encodedConfig []byte) (OffchainConfig, error) {
	var result OffchainConfig
	err := json.Unmarshal(encodedConfig, &result)
	return result, err
}

func (occ OffchainConfig) Encode() ([]byte, error) {
	return json.Marshal(occ)
}

func getEventTopic(name string) gethCommon.Hash {
	abi, err := abi.JSON(strings.NewReader(single_token_offramp.SingleTokenOffRampABI))
	if err != nil {
		panic("could not parse singletoken ABI: " + err.Error())
	}
	event, exists := abi.Events[name]
	if !exists {
		panic(fmt.Sprintf("abi.Events was missing %s", name))
	}
	return event.ID
}
