package ccip

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/logger"
)

func DecodeTollExecutionReport(report types.Report) (*evm_2_evm_toll_offramp.TollExecutionReport, error) {
	unpacked, err := makeTollExecutionReportArgs().Unpack(report)
	if err != nil {
		return nil, err
	}
	if len(unpacked) == 0 {
		return nil, errors.New("assumptionViolation: expected at least one element")
	}

	// Must be anonymous struct here
	erStruct, ok := unpacked[0].(struct {
		SequenceNumbers          []uint64         `json:"sequenceNumbers"`
		TokenPerFeeCoinAddresses []common.Address `json:"tokenPerFeeCoinAddresses"`
		TokenPerFeeCoin          []*big.Int       `json:"tokenPerFeeCoin"`
		EncodedMessages          [][]byte         `json:"encodedMessages"`
		InnerProofs              [][32]uint8      `json:"innerProofs"`
		InnerProofFlagBits       *big.Int         `json:"innerProofFlagBits"`
		OuterProofs              [][32]uint8      `json:"outerProofs"`
		OuterProofFlagBits       *big.Int         `json:"outerProofFlagBits"`
	})
	if !ok {
		return nil, fmt.Errorf("got %T", unpacked[0])
	}
	var er evm_2_evm_toll_offramp.TollExecutionReport
	er.EncodedMessages = append(er.EncodedMessages, erStruct.EncodedMessages...)
	er.InnerProofs = append(er.InnerProofs, erStruct.InnerProofs...)
	er.OuterProofs = append(er.OuterProofs, erStruct.OuterProofs...)
	er.SequenceNumbers = erStruct.SequenceNumbers
	// Unpack will populate with big.Int{false, <allocated empty nat>} for 0 values,
	// which is different from the expected big.NewInt(0). Rebuild to the expected value for this case.
	er.InnerProofFlagBits = big.NewInt(erStruct.InnerProofFlagBits.Int64())
	er.OuterProofFlagBits = big.NewInt(erStruct.OuterProofFlagBits.Int64())
	er.TokenPerFeeCoinAddresses = erStruct.TokenPerFeeCoinAddresses
	er.TokenPerFeeCoin = erStruct.TokenPerFeeCoin
	return &er, nil
}

func MessagesFromTollExecutionReport(report types.Report) ([]uint64, [][]byte, error) {
	tollReport, err := DecodeTollExecutionReport(report)
	if err != nil {
		return nil, nil, err
	}
	return tollReport.SequenceNumbers, tollReport.EncodedMessages, nil
}

func EncodeTollExecutionReport(seqNums []uint64,
	tokensPerFeeCoin map[common.Address]*big.Int,
	msgs [][]byte,
	innerProofs [][32]byte,
	innerProofSourceFlags []bool,
	outerProofs [][32]byte,
	outerProofSourceFlags []bool,
) (types.Report, error) {
	var tokensPerFeeCoinAddresses []common.Address
	var tokensPerFeeCoinValues []*big.Int
	for addr := range tokensPerFeeCoin {
		tokensPerFeeCoinAddresses = append(tokensPerFeeCoinAddresses, addr)
	}
	// Sort the addresses for determinism.
	sort.Slice(tokensPerFeeCoinAddresses, func(i, j int) bool {
		return bytes.Compare(tokensPerFeeCoinAddresses[i].Bytes(), tokensPerFeeCoinAddresses[j].Bytes()) < 0
	})
	for _, addr := range tokensPerFeeCoinAddresses {
		tokensPerFeeCoinValues = append(tokensPerFeeCoinValues, tokensPerFeeCoin[addr])
	}
	report, err := makeTollExecutionReportArgs().PackValues([]interface{}{&evm_2_evm_toll_offramp.TollExecutionReport{
		SequenceNumbers:          seqNums,
		EncodedMessages:          msgs,
		TokenPerFeeCoinAddresses: tokensPerFeeCoinAddresses,
		TokenPerFeeCoin:          tokensPerFeeCoinValues,
		InnerProofs:              innerProofs,
		InnerProofFlagBits:       ProofFlagsToBits(innerProofSourceFlags),
		OuterProofs:              outerProofs,
		OuterProofFlagBits:       ProofFlagsToBits(outerProofSourceFlags),
	}})
	if err != nil {
		return nil, err
	}
	return report, nil
}

var (
	_ types.ReportingPluginFactory = &TollExecutionReportingPluginFactory{}
	_ types.ReportingPlugin        = &TollExecutionReportingPlugin{}
)

type TollExecutionPluginConfig struct {
	lggr                logger.Logger
	source, dest        logpoller.LogPoller
	offRamp             *evm_2_evm_toll_offramp.EVM2EVMTollOffRamp
	onRamp              *evm_2_evm_toll_onramp.EVM2EVMTollOnRamp
	commitStore         *commit_store.CommitStore
	builder             *TollBatchBuilder
	eventSignatures     EventSignatures
	priceGetter         PriceGetter
	leafHasher          LeafHasher[[32]byte]
	rootSnoozeTime      time.Duration
	inflightCacheExpiry time.Duration
	sourceChainID       uint64
	gasLimit            uint64
}

type TollExecutionReportingPluginFactory struct {
	config TollExecutionPluginConfig
}

func NewTollExecutionReportingPluginFactory(
	config TollExecutionPluginConfig,
) types.ReportingPluginFactory {
	return &TollExecutionReportingPluginFactory{config: config}
}

func (rf *TollExecutionReportingPluginFactory) NewReportingPlugin(config types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	offchainConfig, err := Decode(config.OffchainConfig)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	return &TollExecutionReportingPlugin{
			lggr:           rf.config.lggr.Named("TollExecutionReportingPlugin"),
			F:              config.F,
			config:         rf.config,
			offchainConfig: offchainConfig,
			snoozedRoots:   make(map[[32]byte]time.Time),
		}, types.ReportingPluginInfo{
			Name:          "CCIPExecution",
			UniqueReports: true,
			Limits: types.ReportingPluginLimits{
				MaxQueryLength:       MaxQueryLength,
				MaxObservationLength: MaxObservationLength,
				MaxReportLength:      MaxExecutionReportLength,
			},
		}, nil
}

type TollExecutionReportingPlugin struct {
	lggr   logger.Logger
	F      int
	config TollExecutionPluginConfig
	// We need to synchronize access to the inflight structure
	// as reporting plugin methods may be called from separate goroutines,
	// e.g. reporting vs transmission protocol.
	inFlightMu          sync.RWMutex
	inFlight            []InflightExecutionReport
	inflightCacheExpiry time.Duration
	offchainConfig      OffchainConfig
	snoozedRoots        map[[32]byte]time.Time
}

type TollQuery struct {
	TokenPrices map[common.Address]*big.Int `json:"tokenPrices"`
}

// expect percentMultiplier to be [0, 100]
func (r *TollExecutionReportingPlugin) tokenPrices(percentMultiplier *big.Int) (map[common.Address]*big.Int, error) {
	tokensPerFeeCoin := make(map[common.Address]*big.Int)
	executionFeeTokens, err := r.config.offRamp.GetDestinationTokens(nil)
	if err != nil {
		return nil, err
	}
	prices, err := r.config.priceGetter.TokensPerFeeCoin(context.Background(), executionFeeTokens)
	if err != nil {
		return nil, err
	}
	for token, price := range prices {
		buffer := big.NewInt(0).Div(price, percentMultiplier)
		tokensPerFeeCoin[token] = big.NewInt(0).Add(price, buffer)
	}
	return tokensPerFeeCoin, nil
}

func (r *TollExecutionReportingPlugin) Query(ctx context.Context, timestamp types.ReportTimestamp) (types.Query, error) {
	// The leader queries an overestimated set of token prices, which are used by all the followers
	// to compute message executability, ensuring that the set of executable messages is deterministic.
	tokensPerFeeCoin, err := r.tokenPrices(big.NewInt(TokenPriceBufferPercent))
	if err != nil {
		return nil, err
	}
	return json.Marshal(TollQuery{TokenPrices: tokensPerFeeCoin})
}

func (r *TollExecutionReportingPlugin) Observation(ctx context.Context, timestamp types.ReportTimestamp, query types.Query) (types.Observation, error) {
	// GEQuery contains the tokenPricesPerFeeCoin
	lggr := r.lggr.Named("Observation")
	if isCommitStoreDownNow(lggr, r.config.commitStore) {
		return nil, ErrCommitStoreIsDown
	}
	// Expire any inflight reports.
	r.expireInflight(lggr)

	var q TollQuery
	if err := json.Unmarshal(query, &q); err != nil {
		return nil, err
	}
	// Read and make a copy for the builder.
	r.inFlightMu.RLock()
	inFlight := make([]InflightExecutionReport, len(r.inFlight))
	copy(inFlight[:], r.inFlight[:])
	r.inFlightMu.RUnlock()

	batchBuilderStart := time.Now()
	// IMPORTANT: We build executable set based on the leaders token prices, ensuring consistency across followers.
	executableSequenceNumbers, err := r.getExecutableSeqNrs(big.NewInt(MaxGasPrice), q.TokenPrices, inFlight)
	lggr.Infof("Batch building took %d ms", time.Since(batchBuilderStart).Milliseconds())
	if err != nil {
		return nil, err
	}
	if len(executableSequenceNumbers) == 0 {
		return nil, errors.New("No observations")
	}
	lggr.Infof("executable seq nums %v %x", executableSequenceNumbers, r.config.eventSignatures.SendRequested)
	followerTokensPerFeeCoin, err := r.tokenPrices(big.NewInt(TokenPriceBufferPercent))
	if err != nil {
		return nil, err
	}
	return TollExecutionObservation{
		SeqNrs:           executableSequenceNumbers, // Note can be empty
		TokensPerFeeCoin: followerTokensPerFeeCoin,
	}.Marshal()
}

func (r *TollExecutionReportingPlugin) getExecutedSeqNrsInRange(min, max uint64) (map[uint64]struct{}, error) {
	// Should be able to keep this log constant across msg types.
	executedLogs, err := r.config.dest.IndexedLogsTopicRange(r.config.eventSignatures.ExecutionStateChanged, r.config.offRamp.Address(), r.config.eventSignatures.ExecutionStateChangedSequenceNumberIndex, logpoller.EvmWord(min), logpoller.EvmWord(max), int(r.offchainConfig.DestIncomingConfirmations))
	if err != nil {
		return nil, err
	}
	executedMp := make(map[uint64]struct{})
	for _, executedLog := range executedLogs {
		exec, err := r.config.offRamp.ParseExecutionStateChanged(gethtypes.Log{Data: executedLog.Data, Topics: executedLog.GetTopics()})
		if err != nil {
			return nil, err
		}
		executedMp[exec.SequenceNumber] = struct{}{}
	}
	return executedMp, nil
}

func (r *TollExecutionReportingPlugin) getExecutableSeqNrs(
	maxGasPrice *big.Int,
	tokensPerFeeCoin map[common.Address]*big.Int,
	inflight []InflightExecutionReport,
) ([]uint64, error) {
	unexpiredReports, err := getUnexpiredCommitReports(r.config.dest, r.config.commitStore)
	if err != nil {
		return nil, err
	}
	r.lggr.Infow("unexpired roots", "n", len(unexpiredReports))
	if len(unexpiredReports) == 0 {
		return []uint64{}, nil
	}

	// This could result in slightly different values on each call as
	// the function returns the allowed amount at the time of the last block.
	// Since this will only increase over time, the highest observed value will
	// always be the lower bound of what would be available on chain
	// since we already account for inflight txs.
	bucket, err := r.config.offRamp.CalculateCurrentTokenBucketState(nil)
	if err != nil {
		return nil, err
	}
	allowedTokenAmount := bucket.Tokens

	// TODO don't build on every batch builder call but only change on changing configuration
	srcToDst := make(map[common.Address]common.Address)
	sourceTokens, err := r.config.offRamp.GetSupportedTokens(nil)
	if err != nil {
		return nil, err
	}
	for _, sourceToken := range sourceTokens {
		dst, err2 := r.config.offRamp.GetDestinationToken(nil, sourceToken)
		if err2 != nil {
			return nil, err2
		}
		srcToDst[sourceToken] = dst
	}

	supportedDestTokensAndAmounts := make([]common.Address, 0, len(srcToDst))
	for _, destTokensAndAmounts := range srcToDst {
		supportedDestTokensAndAmounts = append(supportedDestTokensAndAmounts, destTokensAndAmounts)
	}

	destTokenPrices, err := r.config.offRamp.GetPricesForTokens(nil, supportedDestTokensAndAmounts)
	if err != nil {
		return nil, err
	}
	pricePerDestToken := make(map[common.Address]*big.Int)
	for i, destToken := range supportedDestTokensAndAmounts {
		pricePerDestToken[destToken] = destTokenPrices[i]
	}

	for _, unexpiredReport := range unexpiredReports {
		var idx int
		var found bool
		for i, onRamp := range unexpiredReport.OnRamps {
			if onRamp == r.config.onRamp.Address() {
				idx = i
				found = true
			}
		}
		if !found {
			continue
		}
		snoozeUntil, haveSnoozed := r.snoozedRoots[unexpiredReport.MerkleRoots[idx]]
		if haveSnoozed && time.Now().Before(snoozeUntil) {
			continue
		}
		blessed, err := r.config.commitStore.IsBlessed(nil, unexpiredReport.RootOfRoots)
		if err != nil {
			return nil, err
		}
		if !blessed {
			r.lggr.Infow("report is accepted but not blessed", "report", hexutil.Encode(unexpiredReport.RootOfRoots[:]))
			continue
		}
		// Check this root for executable messages
		srcLogs, err := r.config.source.LogsDataWordRange(r.config.eventSignatures.SendRequested, r.config.onRamp.Address(), r.config.eventSignatures.SendRequestedSequenceNumberIndex, logpoller.EvmWord(unexpiredReport.Intervals[idx].Min), logpoller.EvmWord(unexpiredReport.Intervals[idx].Max), int(r.offchainConfig.SourceIncomingConfirmations))
		if err != nil {
			return nil, err
		}
		if len(srcLogs) != int(unexpiredReport.Intervals[idx].Max-unexpiredReport.Intervals[idx].Min+1) {
			return nil, errors.Errorf("unexpected missing msgs in committed root %x have %d want %d", unexpiredReport.MerkleRoots[idx], len(srcLogs), int(unexpiredReport.Intervals[idx].Max-unexpiredReport.Intervals[idx].Min+1))
		}
		// TODO: Reorg risk here? I.e. 1 message in a batch, we see its executed so we snooze forever,
		// then it gets reorged out and we'll never retry.
		executedMp, err := r.getExecutedSeqNrsInRange(unexpiredReport.Intervals[idx].Min, unexpiredReport.Intervals[idx].Max)
		if err != nil {
			return nil, err
		}

		batch, allMessagesExecuted := r.config.builder.BuildBatch(srcToDst, srcLogs, executedMp, r.config.gasLimit, maxGasPrice, tokensPerFeeCoin, inflight, allowedTokenAmount, pricePerDestToken)
		// If all messages are already executed, snooze the root for the PERMISSIONLESS_EXECUTION_THRESHOLD_SECONDS,
		// so it will never be considered again.
		if allMessagesExecuted {
			r.lggr.Infof("Snoozing root %s forever since there are no executable txs anymore %v", hex.EncodeToString(unexpiredReport.MerkleRoots[idx][:]), executedMp)
			r.snoozedRoots[unexpiredReport.MerkleRoots[idx]] = time.Now().Add(PERMISSIONLESS_EXECUTION_THRESHOLD)
			continue
		}
		if len(batch) != 0 {
			return batch, nil
		}
		r.snoozedRoots[unexpiredReport.MerkleRoots[idx]] = time.Now().Add(r.config.rootSnoozeTime)
	}
	return []uint64{}, nil
}

func (r *TollExecutionReportingPlugin) parseTollSeqNr(log gethtypes.Log) (uint64, error) {
	s, err := r.config.onRamp.ParseCCIPSendRequested(log)
	if err != nil {
		return 0, err
	}
	return s.Message.SequenceNumber, nil
}

// Assumes non-empty report. Messages to execute can span more than one report, but are assumed to be in order of increasing
// sequence number.
func (r *TollExecutionReportingPlugin) buildReport(lggr logger.Logger, finalSeqNums []uint64, tokensPerFeeCoin map[common.Address]*big.Int) ([]byte, error) {
	me, err := buildExecution(
		lggr,
		r.config.source,
		r.config.dest,
		r.config.onRamp.Address(),
		finalSeqNums,
		r.config.commitStore,
		int(r.offchainConfig.SourceIncomingConfirmations),
		r.config.eventSignatures,
		r.parseTollSeqNr,
		r.config.leafHasher.HashLeaf,
	)
	if err != nil {
		return nil, err
	}
	return EncodeTollExecutionReport(finalSeqNums, tokensPerFeeCoin, me.encMsgs,
		me.innerProofs, me.innerProofSourceFlags, me.outerProofs, me.outerProofSourceFlags)
}

func (r *TollExecutionReportingPlugin) Report(ctx context.Context, timestamp types.ReportTimestamp, query types.Query, observations []types.AttributedObservation) (bool, types.Report, error) {
	lggr := r.lggr.Named("Report")
	if isCommitStoreDownNow(lggr, r.config.commitStore) {
		return false, nil, ErrCommitStoreIsDown
	}
	actualMaybeObservations := getNonEmptyObservations[GEExecutionObservation](lggr, observations)
	var actualObservations []GEExecutionObservation
	tokens, err := r.config.offRamp.GetDestinationTokens(nil)
	if err != nil {
		return false, nil, err
	}
	priceObservations := make(map[common.Address][]*big.Int)
	for _, obs := range actualMaybeObservations {
		hasAllPrices := true
		for _, token := range tokens {
			if _, ok := obs.TokensPerFeeCoin[token]; !ok {
				hasAllPrices = false
				break
			}
		}
		if !hasAllPrices {
			continue
		}
		// If it has all the prices, add each price to observations
		for token, price := range obs.TokensPerFeeCoin {
			priceObservations[token] = append(priceObservations[token], price)
		}
		// Add source gas price
		actualObservations = append(actualObservations, obs)
	}
	// Need at least F+1 observations
	if len(actualObservations) <= r.F {
		lggr.Tracew("Non-empty observations <= F, need at least F+1 to continue")
		return false, nil, nil
	}
	// If we have sufficient observations, only build a report if
	// the leaders prices is >= the median of the followers prices.
	// Note we accept that this can result in the leader stalling progress,
	// by setting an extremely high set of prices, but a malicious leader always had that ability
	// and eventually we'll elect a new one.
	var q TollQuery
	if err2 := json.Unmarshal(query, &q); err2 != nil {
		return false, nil, err2
	}
	medianTokensPerFeeCoin := make(map[common.Address]*big.Int)
	for _, token := range tokens {
		medianTokensPerFeeCoin[token] = median(priceObservations[token])
		if medianTokensPerFeeCoin[token].Cmp(q.TokenPrices[token]) == 1 {
			// Leader specified a price which is too low, reject this report.
			// TODO: Error or not here?
			lggr.Warnw("Leader price is too low, skipping report", "leader", q.TokenPrices[token], "followers", medianTokensPerFeeCoin[token])
			return false, nil, nil
		}
	}
	// If we make it here, the leader has specified a valid set of prices.
	tally := make(map[uint64]int)
	for _, obs := range actualObservations {
		for _, seqNr := range obs.SeqNrs {
			tally[seqNr]++
		}
	}
	var finalSequenceNumbers []uint64
	for seqNr, count := range tally {
		// Note spec deviation - I think it's ok to rely on the batch builder for capping the number of messages vs capping in two places/ways?
		if count > r.F {
			finalSequenceNumbers = append(finalSequenceNumbers, seqNr)
		}
	}
	// buildReport expects sorted sequence numbers (tally map is non-deterministic).
	sort.Slice(finalSequenceNumbers, func(i, j int) bool {
		return finalSequenceNumbers[i] < finalSequenceNumbers[j]
	})
	// Important we actually execute based on the medianTokensPrices, which we ensure
	// is <= than prices used to determine executability.
	report, err := r.buildReport(lggr, finalSequenceNumbers, medianTokensPerFeeCoin)
	if err != nil {
		return false, nil, err
	}
	lggr.Infow("Built report",
		"onRampAddr", r.config.onRamp,
		"finalSeqNums", finalSequenceNumbers,
		"executionPrices", medianTokensPerFeeCoin)
	return true, report, nil
}

func (r *TollExecutionReportingPlugin) expireInflight(lggr logger.Logger) {
	r.inFlightMu.Lock()
	defer r.inFlightMu.Unlock()
	// Reap old inflight txs and check if any messages in the report are inflight.
	var stillInFlight []InflightExecutionReport
	for _, report := range r.inFlight {
		if time.Since(report.createdAt) > r.inflightCacheExpiry {
			// Happy path: inflight report was successfully transmitted onchain, we remove it from inflight and onchain state reflects inflight.
			// Sad path: inflight report reverts onchain, we remove it from inflight, onchain state does not reflect the change so we retry.
			lggr.Infow("Inflight report expired", "seqNums", report.seqNrs)
		} else {
			stillInFlight = append(stillInFlight, report)
		}
	}
	r.inFlight = stillInFlight
}

func (r *TollExecutionReportingPlugin) addToInflight(lggr logger.Logger, seqNrs []uint64, encMsgs [][]byte) error {
	r.inFlightMu.Lock()
	defer r.inFlightMu.Unlock()
	for _, report := range r.inFlight {
		// TODO: Think about if this fails in reorgs
		if report.seqNrs[0] == seqNrs[0] {
			return errors.Errorf("report is already in flight")
		}
	}
	// Otherwise not already in flight, add it.
	lggr.Infow("Added report to inflight",
		"seqNums", seqNrs)
	r.inFlight = append(r.inFlight, InflightExecutionReport{
		createdAt:   time.Now(),
		seqNrs:      seqNrs,
		encMessages: encMsgs,
	})
	return nil
}

func (r *TollExecutionReportingPlugin) ShouldAcceptFinalizedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	lggr := r.lggr.Named("ShouldAcceptFinalizedReport")
	seqNrs, encMsgs, err := MessagesFromTollExecutionReport(report)
	if err != nil {
		lggr.Errorw("unable to decode report", "err", err)
		return false, nil
	}
	lggr.Infof("Seq nums %v", seqNrs)
	// If the first message is executed already, this execution report is stale, and we do not accept it.
	stale, err2 := r.isStaleReport(seqNrs[0])
	if err2 != nil {
		return !stale, err2
	}
	if stale {
		return false, nil
	}
	// Else just assume in flight
	if err = r.addToInflight(lggr, seqNrs, encMsgs); err != nil {
		return false, err
	}
	return true, nil
}

func (r *TollExecutionReportingPlugin) ShouldTransmitAcceptedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	seqNrs, _, err := MessagesFromTollExecutionReport(report)
	if err != nil {
		return false, nil
	}
	// If report is not stale we transmit.
	// When the executeTransmitter enqueues the tx for tx manager,
	// we mark it as execution_sent, removing it from the set of inflight messages.
	if len(seqNrs) > 0 {
		stale, err := r.isStaleReport(seqNrs[0])
		return !stale, err
	}
	// TODO: how to check for staleness on a purely fee update report?
	return false, nil
}

func (r *TollExecutionReportingPlugin) isStaleReport(min uint64) (bool, error) {
	// If the first message is executed already, this execution report is stale.
	msgState, err := r.config.offRamp.GetExecutionState(nil, min)
	if err != nil {
		// TODO: do we need to check for not present error?
		return true, err
	}
	if msgState == MessageStateFailure || msgState == MessageStateSuccess {
		return true, nil
	}
	return false, nil
}

func (r *TollExecutionReportingPlugin) Close() error {
	return nil
}
