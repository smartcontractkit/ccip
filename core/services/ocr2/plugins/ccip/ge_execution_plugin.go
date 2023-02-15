package ccip

import (
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

	"github.com/smartcontractkit/chainlink/core/assets"
	"github.com/smartcontractkit/chainlink/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/fee_manager"
	"github.com/smartcontractkit/chainlink/core/logger"
)

type FeeUpdate = evm_2_evm_offramp.InternalFeeUpdate

func MessagesFromExecutionReport(report types.Report) ([]uint64, [][]byte, []FeeUpdate, error) {
	decodeExecutionReport, err := DecodeExecutionReport(report)
	if err != nil {
		return nil, nil, nil, err
	}
	return decodeExecutionReport.SequenceNumbers, decodeExecutionReport.EncodedMessages, decodeExecutionReport.FeeUpdates, nil
}

func DecodeExecutionReport(report types.Report) (*evm_2_evm_offramp.InternalExecutionReport, error) {
	unpacked, err := makeExecutionReportArgs().Unpack(report)
	if err != nil {
		return nil, err
	}
	if len(unpacked) == 0 {
		return nil, errors.New("assumptionViolation: expected at least one element")
	}

	// Must be anonymous struct here
	erStruct, ok := unpacked[0].(struct {
		SequenceNumbers []uint64 `json:"sequenceNumbers"`
		FeeUpdates      []struct {
			SourceFeeToken              common.Address `json:"sourceFeeToken"`
			DestChainId                 uint64         `json:"destChainId"`
			FeeTokenBaseUnitsPerUnitGas *big.Int       `json:"feeTokenBaseUnitsPerUnitGas"`
		} `json:"feeUpdates"`
		EncodedMessages [][]byte    `json:"encodedMessages"`
		Proofs          [][32]uint8 `json:"proofs"`
		ProofFlagBits   *big.Int    `json:"proofFlagBits"`
	})
	if !ok {
		return nil, fmt.Errorf("got %T", unpacked[0])
	}
	var er evm_2_evm_offramp.InternalExecutionReport
	er.EncodedMessages = append(er.EncodedMessages, erStruct.EncodedMessages...)
	er.Proofs = append(er.Proofs, erStruct.Proofs...)
	er.FeeUpdates = []evm_2_evm_offramp.InternalFeeUpdate{}

	for _, feeUpdate := range erStruct.FeeUpdates {
		er.FeeUpdates = append(er.FeeUpdates, FeeUpdate{
			SourceFeeToken:              feeUpdate.SourceFeeToken,
			DestChainId:                 feeUpdate.DestChainId,
			FeeTokenBaseUnitsPerUnitGas: feeUpdate.FeeTokenBaseUnitsPerUnitGas,
		})
	}

	er.SequenceNumbers = erStruct.SequenceNumbers
	// Unpack will populate with big.Int{false, <allocated empty nat>} for 0 values,
	// which is different from the expected big.NewInt(0). Rebuild to the expected value for this case.
	er.ProofFlagBits = big.NewInt(erStruct.ProofFlagBits.Int64())
	return &er, nil
}

func EncodeExecutionReport(seqNums []uint64,
	msgs [][]byte,
	proofs [][32]byte,
	proofSourceFlags []bool,
	feeUpdates []evm_2_evm_offramp.InternalFeeUpdate,
) (types.Report, error) {
	report, err := makeExecutionReportArgs().PackValues([]interface{}{&evm_2_evm_offramp.InternalExecutionReport{
		SequenceNumbers: seqNums,
		FeeUpdates:      feeUpdates,
		EncodedMessages: msgs,
		Proofs:          proofs,
		ProofFlagBits:   ProofFlagsToBits(proofSourceFlags),
	}})
	if err != nil {
		return nil, err
	}
	return report, nil
}

var (
	_ types.ReportingPluginFactory = &ExecutionReportingPluginFactory{}
	_ types.ReportingPlugin        = &ExecutionReportingPlugin{}
)

type ExecutionPluginConfig struct {
	onRamp                               *evm_2_evm_onramp.EVM2EVMOnRamp
	offRamp                              *evm_2_evm_offramp.EVM2EVMOffRamp
	commitStore                          *commit_store.CommitStore
	feeManager                           *fee_manager.FeeManager
	source, dest                         logpoller.LogPoller
	eventSignatures                      EventSignatures
	priceGetter                          PriceGetter
	snoozeTime                           time.Duration
	inflightCacheExpiry                  time.Duration
	sourceGasEstimator, destGasEstimator gas.Estimator
	sourceChainID                        uint64
	builder                              BatchBuilderInterface
	leafHasher                           LeafHasherInterface[[32]byte]
	lggr                                 logger.Logger
	gasLimit                             uint64
}

type ExecutionReportingPluginFactory struct {
	config ExecutionPluginConfig
}

func NewExecutionReportingPluginFactory(config ExecutionPluginConfig) types.ReportingPluginFactory {
	return &ExecutionReportingPluginFactory{config: config}
}

func (rf *ExecutionReportingPluginFactory) NewReportingPlugin(config types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	offchainConfig, err := Decode(config.OffchainConfig)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}

	execTokens, err := rf.config.offRamp.GetDestinationTokens(nil)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	// TODO: Hack assume link is the first token
	linkToken := execTokens[0]

	return &ExecutionReportingPlugin{
			lggr:           rf.config.lggr.Named("ExecutionReportingPlugin"),
			F:              config.F,
			offchainConfig: offchainConfig,
			config:         rf.config,
			snoozedRoots:   make(map[[32]byte]time.Time),
			linkToken:      linkToken,
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

type ExecutionReportingPlugin struct {
	lggr   logger.Logger
	F      int
	config ExecutionPluginConfig
	// We need to synchronize access to the inflight structure
	// as reporting plugin methods may be called from separate goroutines,
	// e.g. reporting vs transmission protocol.
	inFlightMu         sync.RWMutex
	inFlight           []InflightInternalExecutionReport
	inFlightFeeUpdates [][]FeeUpdate
	offchainConfig     OffchainConfig
	snoozedRoots       map[[32]byte]time.Time
	linkToken          common.Address
}

type Query struct {
	TokenPrices  map[common.Address]*big.Int `json:"tokenPrices"`
	DestGasPrice *big.Int                    `json:"destGasPrice"`
}

// expect percentMultiplier to be [0, 100]
func (r *ExecutionReportingPlugin) tokenPrices(percentMultiplier *big.Int) (map[common.Address]*big.Int, error) {
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

func (r *ExecutionReportingPlugin) Query(ctx context.Context, timestamp types.ReportTimestamp) (types.Query, error) {
	// The leader queries an overestimated set of token prices, which are used by all the followers
	// to compute message executability, ensuring that the set of executable messages is deterministic.
	tokensPerFeeCoin, err := r.tokenPrices(big.NewInt(TokenPriceBufferPercent))
	if err != nil {
		return nil, err
	}
	// In the context of CCIP, latency is much less important than cost, so generally we'd prefer to wait vs bump at all.
	// Options:
	// - Disable bumping entirely. Execute messages up to max loss given current price and simply wait until it is included. Means that we can potentially
	// block execution for all jobs until that tx goes though.
	// - Have very very relaxed bumping settings and a relatively low cap (say 1k gwei). The worst case is that we decide to execute a max loss tx,
	// prices remain high for a long time to invoke bumping and increase our loss up to the bumped cap. Benefit is we will unblock the jobs ourselves.
	// Should be possible to ensure the max bumped loss is incurred with some extremely low probability (something much worse than a luna type meltdown of 4hr 8k gwei spike).
	// TODO: Switch between 1559 and non-1559 here based on chain (or wrap estimator at a higher level).
	destGasPrice, _, err := r.config.destGasEstimator.GetLegacyGas(ctx, nil, BatchGasLimit, assets.NewWei(big.NewInt(int64(MaxGasPrice))))
	if err != nil {
		return nil, err
	}
	return json.Marshal(Query{TokenPrices: tokensPerFeeCoin, DestGasPrice: destGasPrice.ToInt()})
}

func (r *ExecutionReportingPlugin) Observation(ctx context.Context, timestamp types.ReportTimestamp, query types.Query) (types.Observation, error) {
	// Query contains the tokenPricesPerFeeCoin
	lggr := r.lggr.Named("ExecutionObservation")
	if isCommitStoreDownNow(lggr, r.config.commitStore) {
		return nil, ErrCommitStoreIsDown
	}
	// Expire any inflight reports.
	r.expireInflight(lggr)

	var q Query
	if err := json.Unmarshal(query, &q); err != nil {
		return nil, err
	}
	// Read and make a copy for the builder.
	r.inFlightMu.RLock()
	inFlight := make([]InflightInternalExecutionReport, len(r.inFlight))
	copy(inFlight[:], r.inFlight[:])
	r.inFlightMu.RUnlock()

	batchBuilderStart := time.Now()
	// IMPORTANT: We build executable set based on the leaders token prices, ensuring consistency across followers.
	executableSequenceNumbers, err := r.getExecutableSeqNrs(q.DestGasPrice, q.TokenPrices, inFlight)
	lggr.Infof("Batch building took %d ms", time.Since(batchBuilderStart).Milliseconds())
	if err != nil {
		return nil, err
	}
	lggr.Infof("executable seq nums %v %x", executableSequenceNumbers, r.config.eventSignatures.SendRequested)
	followerTokensPerFeeCoin, err := r.tokenPrices(big.NewInt(TokenPriceBufferPercent))
	if err != nil {
		return nil, err
	}
	// Observe a source chain price for pricing.
	// TODO: 1559 support
	sourceGasPriceWei, _, err := r.config.sourceGasEstimator.GetLegacyGas(ctx, nil, BatchGasLimit, assets.NewWei(big.NewInt(int64(MaxGasPrice))))
	if err != nil {
		return nil, err
	}
	sourceGasPrice := sourceGasPriceWei.ToInt()

	if canSkip, err := r.canSkipFeeUpdate(r.calculateFeeTokenBaseUnitsPerUnitGas(sourceGasPrice, followerTokensPerFeeCoin[r.linkToken])); err != nil {
		return nil, err
	} else if canSkip {
		sourceGasPrice = nil // vote skip
	}

	return ExecutionObservation{
		SeqNrs:           executableSequenceNumbers, // Note can be empty
		TokensPerFeeCoin: followerTokensPerFeeCoin,
		SourceGasPrice:   sourceGasPrice,
	}.Marshal()
}

func (r *ExecutionReportingPlugin) canSkipFeeUpdate(feeTokenBaseUnitsPerUnitGas *big.Int) (bool, error) {
	token := r.linkToken
	chainID := r.config.sourceChainID

	var latestUpdateTimestamp time.Time
	var latestUpdate FeeUpdate

	logsWithinHeartBeat, err := r.config.dest.LogsCreatedAfter(GasFeeUpdated, r.config.feeManager.Address(), time.Now().Add(-r.offchainConfig.FeeUpdateHeartBeat.Duration()))
	if err != nil {
		return false, err
	}
	for _, log := range logsWithinHeartBeat {
		parsed, err := r.config.feeManager.ParseGasFeeUpdated(log.GetGethLog())
		if err != nil {
			return false, err
		}
		if ts := time.Unix(int64(parsed.Timestamp), 0); parsed.DestChain == chainID && parsed.Token == token && !ts.Before(latestUpdateTimestamp) {
			latestUpdateTimestamp = ts
			latestUpdate = FeeUpdate{
				SourceFeeToken:              token,
				DestChainId:                 parsed.DestChain,
				FeeTokenBaseUnitsPerUnitGas: parsed.FeeTokenBaseUnitsPerUnitGas,
			}
		}
	}

	r.inFlightMu.RLock()
	for i, inflight := range r.inFlightFeeUpdates {
		for _, update := range inflight {
			if update.DestChainId == chainID && update.SourceFeeToken == token && !r.inFlight[i].createdAt.Before(latestUpdateTimestamp) {
				latestUpdateTimestamp = r.inFlight[i].createdAt
				latestUpdate = update
			}
		}
	}
	r.inFlightMu.RUnlock()

	if time.Since(latestUpdateTimestamp) > r.offchainConfig.FeeUpdateHeartBeat.Duration() {
		return false, nil
	}

	deviation := big.NewInt(0).Sub(feeTokenBaseUnitsPerUnitGas, latestUpdate.FeeTokenBaseUnitsPerUnitGas)
	deviation.Mul(deviation, big.NewInt(1e9))
	deviation.Div(deviation, latestUpdate.FeeTokenBaseUnitsPerUnitGas) // deviation_parts_per_billion = ((x2 - x1) / x1) * 1e9

	// can skip if latest feeUpdate for this plugin's sourceChainId, linkToken, is within heartbeat and deviation
	return deviation.CmpAbs(big.NewInt(int64(r.offchainConfig.FeeUpdateDeviationPPB))) <= 0, nil
}

func (r *ExecutionReportingPlugin) calculateFeeTokenBaseUnitsPerUnitGas(sourceGasPrice *big.Int, juelsPerFeeCoin *big.Int) (feeTokenBaseUnitsPerUnitGas *big.Int) {
	// (juels/eth) * (wei / gas) / (1 eth / 1e18 wei) = juels/gas
	// TODO: Think more about this offchain/onchain computation split
	feeTokenBaseUnitsPerUnitGas = big.NewInt(0).Mul(sourceGasPrice, juelsPerFeeCoin)
	return feeTokenBaseUnitsPerUnitGas.Div(feeTokenBaseUnitsPerUnitGas, big.NewInt(1e18))
}

func (r *ExecutionReportingPlugin) generateFeeUpdate(sourceGasPrice *big.Int, juelsPerFeeCoin *big.Int) []FeeUpdate {
	// if sourceGasPrice got majority votes as nil, skip generating feeUpdate
	if sourceGasPrice == nil {
		return nil
	}

	return []FeeUpdate{
		{
			SourceFeeToken: r.linkToken,
			// Since this gas fee update will be sent to the destination chain, this plugins
			// source chain will be the feeUpdaters destination chain.
			DestChainId:                 r.config.sourceChainID,
			FeeTokenBaseUnitsPerUnitGas: r.calculateFeeTokenBaseUnitsPerUnitGas(sourceGasPrice, juelsPerFeeCoin),
		},
	}
}

func (r *ExecutionReportingPlugin) getExecutedSeqNrsInRange(min, max uint64) (map[uint64]struct{}, error) {
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

func (r *ExecutionReportingPlugin) getExecutableSeqNrs(
	maxGasPrice *big.Int,
	tokensPerFeeCoin map[common.Address]*big.Int,
	inflight []InflightInternalExecutionReport,
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

	supportedDestTokenAmounts := make([]common.Address, 0, len(srcToDst))
	for _, destTokenAmounts := range srcToDst {
		supportedDestTokenAmounts = append(supportedDestTokenAmounts, destTokenAmounts)
	}

	destTokenPrices, err := r.config.offRamp.GetPricesForTokens(nil, supportedDestTokenAmounts)
	if err != nil {
		return nil, err
	}

	pricePerDestToken := make(map[common.Address]*big.Int)
	for i, destToken := range supportedDestTokenAmounts {
		pricePerDestToken[destToken] = destTokenPrices[i]
	}

	for _, unexpiredReport := range unexpiredReports {
		snoozeUntil, haveSnoozed := r.snoozedRoots[unexpiredReport.MerkleRoot]
		if haveSnoozed && time.Now().Before(snoozeUntil) {
			incSkippedRequests(reasonSnoozed)
			continue
		}
		blessed, err := r.config.commitStore.IsBlessed(nil, unexpiredReport.MerkleRoot)
		if err != nil {
			return nil, err
		}
		if !blessed {
			r.lggr.Infow("report is accepted but not blessed", "report", hexutil.Encode(unexpiredReport.MerkleRoot[:]))
			incSkippedRequests(reasonNotBlessed)
			continue
		}
		// Check this root for executable messages
		srcLogs, err := r.config.source.LogsDataWordRange(r.config.eventSignatures.SendRequested, r.config.onRamp.Address(), r.config.eventSignatures.SendRequestedSequenceNumberIndex, logpoller.EvmWord(unexpiredReport.Interval.Min), logpoller.EvmWord(unexpiredReport.Interval.Max), int(r.offchainConfig.SourceIncomingConfirmations))
		if err != nil {
			return nil, err
		}
		if len(srcLogs) != int(unexpiredReport.Interval.Max-unexpiredReport.Interval.Min+1) {
			return nil, errors.Errorf("unexpected missing msgs in committed root %x have %d want %d", unexpiredReport.MerkleRoot, len(srcLogs), int(unexpiredReport.Interval.Max-unexpiredReport.Interval.Min+1))
		}
		// TODO: Reorg risk here? I.e. 1 message in a batch, we see its executed so we snooze forever,
		// then it gets reorged out and we'll never retry.
		executedMp, err := r.getExecutedSeqNrsInRange(unexpiredReport.Interval.Min, unexpiredReport.Interval.Max)
		if err != nil {
			return nil, err
		}

		batch, allMessagesExecuted := r.config.builder.BuildBatch(srcToDst, srcLogs, executedMp, r.config.gasLimit, maxGasPrice, tokensPerFeeCoin, inflight, allowedTokenAmount, pricePerDestToken)
		// If all messages are already executed, snooze the root for the PERMISSIONLESS_EXECUTION_THRESHOLD_SECONDS,
		// so it will never be considered again.
		if allMessagesExecuted {
			r.lggr.Infof("Snoozing root %s forever since there are no executable txs anymore %v", hex.EncodeToString(unexpiredReport.MerkleRoot[:]), executedMp)
			r.snoozedRoots[unexpiredReport.MerkleRoot] = time.Now().Add(PERMISSIONLESS_EXECUTION_THRESHOLD)
			incSkippedRequests(reasonAllExecuted)
			continue
		}
		if len(batch) != 0 {
			return batch, nil
		}
		r.snoozedRoots[unexpiredReport.MerkleRoot] = time.Now().Add(r.config.snoozeTime)
	}
	return []uint64{}, nil
}

func (r *ExecutionReportingPlugin) parseSeqNr(log gethtypes.Log) (uint64, error) {
	s, err := r.config.onRamp.ParseCCIPSendRequested(log)
	if err != nil {
		return 0, err
	}
	return s.Message.SequenceNumber, nil
}

// Assumes non-empty report. Messages to execute can span more than one report, but are assumed to be in order of increasing
// sequence number.
func (r *ExecutionReportingPlugin) buildReport(lggr logger.Logger, finalSeqNums []uint64, tokensPerFeeCoin map[common.Address]*big.Int, sourceGasPrice *big.Int) ([]byte, error) {
	var err error
	var me *MessageExecution
	if len(finalSeqNums) > 0 {
		me, err = buildExecution(
			lggr,
			r.config.source,
			r.config.dest,
			r.config.onRamp.Address(),
			finalSeqNums,
			r.config.commitStore,
			int(r.offchainConfig.SourceIncomingConfirmations),
			r.config.eventSignatures,
			r.parseSeqNr,
			r.config.leafHasher.HashLeaf,
		)
		if err != nil {
			return nil, err
		}
	}
	gasFeeUpdates := r.generateFeeUpdate(sourceGasPrice, tokensPerFeeCoin[r.linkToken])
	if len(gasFeeUpdates) == 0 && len(finalSeqNums) == 0 {
		return nil, errors.New("No report needed")
	}
	if len(finalSeqNums) != 0 {
		return EncodeExecutionReport(finalSeqNums,
			me.encMsgs,
			me.proofs,
			me.proofSourceFlags,
			gasFeeUpdates,
		)
	}
	lggr.Infow("Building execution report fee update only", "feeUpdates", gasFeeUpdates)
	return EncodeExecutionReport(finalSeqNums, nil, nil, nil, gasFeeUpdates)
}

func (r *ExecutionReportingPlugin) Report(ctx context.Context, timestamp types.ReportTimestamp, query types.Query, observations []types.AttributedObservation) (bool, types.Report, error) {
	lggr := r.lggr.Named("Report")
	if isCommitStoreDownNow(lggr, r.config.commitStore) {
		return false, nil, ErrCommitStoreIsDown
	}
	actualMaybeObservations := getNonEmptyObservations[ExecutionObservation](lggr, observations)
	var actualObservations []ExecutionObservation
	tokens, err := r.config.offRamp.GetDestinationTokens(nil)
	if err != nil {
		return false, nil, err
	}
	priceObservations := make(map[common.Address][]*big.Int)
	var sourceGasObservations []*big.Int
	var sourceGasPriceNilCount int
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
		if obs.SourceGasPrice == nil {
			sourceGasPriceNilCount++
		} else {
			// Add only non-nil source gas price
			sourceGasObservations = append(sourceGasObservations, obs.SourceGasPrice)
		}
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
	// and, eventually we'll elect a new one.
	var q Query
	if err = json.Unmarshal(query, &q); err != nil {
		return false, nil, err
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

	var sourceGasPrice *big.Int
	// skip gasPrice feeUpdate by leaving it nil if majority voted so by sending nil gasPrice observations
	if sourceGasPriceNilCount <= len(sourceGasObservations) {
		sourceGasPrice = median(sourceGasObservations)
	}

	// Important we actually execute based on the medianTokensPrices, which we ensure
	// is <= than prices used to determine executability.
	report, err := r.buildReport(lggr, finalSequenceNumbers, medianTokensPerFeeCoin, sourceGasPrice)
	if err != nil {
		return false, nil, err
	}
	lggr.Infow("Built report",
		"onRampAddr", r.config.onRamp.Address(),
		"finalSeqNums", finalSequenceNumbers,
		"executionPrices", medianTokensPerFeeCoin)
	return true, report, nil
}

func (r *ExecutionReportingPlugin) expireInflight(lggr logger.Logger) {
	r.inFlightMu.Lock()
	defer r.inFlightMu.Unlock()
	// Reap old inflight txs and check if any messages in the report are inflight.
	var stillInFlight []InflightInternalExecutionReport
	var stillInFlightFeeUpdates [][]FeeUpdate
	for i, report := range r.inFlight {
		if time.Since(report.createdAt) > r.config.inflightCacheExpiry {
			// Happy path: inflight report was successfully transmitted onchain, we remove it from inflight and onchain state reflects inflight.
			// Sad path: inflight report reverts onchain, we remove it from inflight, onchain state does not reflect the change so we retry.
			lggr.Infow("Inflight report expired", "seqNums", report.seqNrs)
		} else {
			stillInFlight = append(stillInFlight, report)
			stillInFlightFeeUpdates = append(stillInFlightFeeUpdates, r.inFlightFeeUpdates[i])
		}
	}
	r.inFlight = stillInFlight
	r.inFlightFeeUpdates = stillInFlightFeeUpdates
}

func (r *ExecutionReportingPlugin) addToInflight(lggr logger.Logger, seqNrs []uint64, encMsgs [][]byte, feeUpdates []FeeUpdate) error {
	r.inFlightMu.Lock()
	defer r.inFlightMu.Unlock()
	for _, report := range r.inFlight {
		// TODO: Think about if this fails in reorgs
		if (len(report.seqNrs) > 0 && len(seqNrs) > 0) && (report.seqNrs[0] == seqNrs[0]) {
			return errors.Errorf("report is already in flight")
		}
	}
	// Otherwise not already in flight, add it.
	lggr.Infow("Added report to inflight",
		"seqNums", seqNrs)
	r.inFlight = append(r.inFlight, InflightInternalExecutionReport{
		createdAt:   time.Now(),
		seqNrs:      seqNrs,
		encMessages: encMsgs,
	})
	r.inFlightFeeUpdates = append(r.inFlightFeeUpdates, feeUpdates)
	return nil
}

func (r *ExecutionReportingPlugin) ShouldAcceptFinalizedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	lggr := r.lggr.Named("ShouldAcceptFinalizedReport")
	seqNrs, encMsgs, feeUpdates, err := MessagesFromExecutionReport(report)
	if err != nil {
		lggr.Errorw("unable to decode report", "err", err)
		return false, nil
	}
	lggr.Infof("Seq nums %v", seqNrs)
	// If the first message is executed already, this execution report is stale, and we do not accept it.
	stale, err := r.isStaleReport(seqNrs, feeUpdates)
	if err != nil {
		return !stale, err
	}
	if stale {
		return false, nil
	}
	// Else just assume in flight
	if err = r.addToInflight(lggr, seqNrs, encMsgs, feeUpdates); err != nil {
		return false, err
	}
	return true, nil
}

func (r *ExecutionReportingPlugin) ShouldTransmitAcceptedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	seqNrs, _, feeUpdates, err := MessagesFromExecutionReport(report)
	if err != nil {
		return false, nil
	}
	// If report is not stale we transmit.
	// When the executeTransmitter enqueues the tx for tx manager,
	// we mark it as execution_sent, removing it from the set of inflight messages.
	stale, err := r.isStaleReport(seqNrs, feeUpdates)
	return !stale, err
}

func (r *ExecutionReportingPlugin) isStaleReport(seqNrs []uint64, feeUpdates []FeeUpdate) (bool, error) {
	// If the first message is executed already, this execution report is stale.
	if len(seqNrs) > 0 {
		min := seqNrs[0]
		msgState, err := r.config.offRamp.GetExecutionState(nil, min)
		if err != nil {
			// TODO: do we need to check for not present error?
			return true, err
		}
		if msgState == MessageStateFailure || msgState == MessageStateSuccess {
			return true, nil
		}
	} else if len(feeUpdates) > 0 {
		update := feeUpdates[0]
		fee, err := r.config.feeManager.GetFeeTokenBaseUnitsPerUnitGas(nil, update.SourceFeeToken, update.DestChainId)
		if err != nil {
			return true, err
		}
		// for feeUpdates-only reports, if the fee is equal our report, it is stale (not needed or already executed)
		return fee.Cmp(update.FeeTokenBaseUnitsPerUnitGas) == 0, nil
	}
	return false, nil
}

func (r *ExecutionReportingPlugin) Close() error {
	return nil
}
