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

	"github.com/smartcontractkit/chainlink/core/assets"
	"github.com/smartcontractkit/chainlink/core/chains/evm/gas"
	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_ge_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_ge_onramp"
	"github.com/smartcontractkit/chainlink/core/logger"
)

func MessagesFromGEExecutionReport(report types.Report) ([]uint64, [][]byte, error) {
	geReport, err := DecodeGEExecutionReport(report)
	if err != nil {
		return nil, nil, err
	}
	return geReport.SequenceNumbers, geReport.EncodedMessages, nil
}

func DecodeGEExecutionReport(report types.Report) (*evm_2_evm_ge_offramp.GEExecutionReport, error) {
	unpacked, err := makeExecutionReportArgs().Unpack(report)
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
		FeeUpdates               []struct {
			SourceFeeToken common.Address `json:"sourceFeeToken"`
			DestChainId    uint64         `json:"destChainId"`
			LinkPerUnitGas *big.Int       `json:"linkPerUnitGas"`
		} `json:"feeUpdates"`
		EncodedMessages    [][]byte    `json:"encodedMessages"`
		InnerProofs        [][32]uint8 `json:"innerProofs"`
		InnerProofFlagBits *big.Int    `json:"innerProofFlagBits"`
		OuterProofs        [][32]uint8 `json:"outerProofs"`
		OuterProofFlagBits *big.Int    `json:"outerProofFlagBits"`
	})
	if !ok {
		return nil, fmt.Errorf("got %T", unpacked[0])
	}
	var er evm_2_evm_ge_offramp.GEExecutionReport
	er.EncodedMessages = append(er.EncodedMessages, erStruct.EncodedMessages...)
	er.InnerProofs = append(er.InnerProofs, erStruct.InnerProofs...)
	er.OuterProofs = append(er.OuterProofs, erStruct.OuterProofs...)

	er.FeeUpdates = []evm_2_evm_ge_offramp.GEFeeUpdate{}

	for _, feeUpdate := range erStruct.FeeUpdates {
		er.FeeUpdates = append(er.FeeUpdates, evm_2_evm_ge_offramp.GEFeeUpdate{
			SourceFeeToken: feeUpdate.SourceFeeToken,
			DestChainId:    feeUpdate.DestChainId,
			LinkPerUnitGas: feeUpdate.LinkPerUnitGas,
		})
	}

	er.SequenceNumbers = erStruct.SequenceNumbers
	// Unpack will populate with big.Int{false, <allocated empty nat>} for 0 values,
	// which is different from the expected big.NewInt(0). Rebuild to the expected value for this case.
	er.InnerProofFlagBits = big.NewInt(erStruct.InnerProofFlagBits.Int64())
	er.OuterProofFlagBits = big.NewInt(erStruct.OuterProofFlagBits.Int64())
	er.TokenPerFeeCoinAddresses = erStruct.TokenPerFeeCoinAddresses
	er.TokenPerFeeCoin = erStruct.TokenPerFeeCoin
	return &er, nil
}

func EncodeGEExecutionReport(seqNums []uint64,
	tokensPerFeeCoin map[common.Address]*big.Int,
	msgs [][]byte,
	innerProofs [][32]byte,
	innerProofSourceFlags []bool,
	outerProofs [][32]byte,
	outerProofSourceFlags []bool,
	feeUpdates []evm_2_evm_ge_offramp.GEFeeUpdate,
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
	report, err := makeExecutionReportArgs().PackValues([]interface{}{&evm_2_evm_ge_offramp.GEExecutionReport{
		SequenceNumbers:          seqNums,
		FeeUpdates:               feeUpdates,
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
	_ types.ReportingPluginFactory = &GEExecutionReportingPluginFactory{}
	_ types.ReportingPlugin        = &GEExecutionReportingPlugin{}
)

type GEExecutionPluginConfig struct {
	onRamp                               *evm_2_evm_ge_onramp.EVM2EVMGEOnRamp
	offRamp                              *evm_2_evm_ge_offramp.EVM2EVMGEOffRamp
	commitStore                          *commit_store.CommitStore
	source, dest                         logpoller.LogPoller
	eventSignatures                      EventSignatures
	priceGetter                          PriceGetter
	snoozeTime                           time.Duration
	inflightCacheExpiry                  time.Duration
	sourceGasEstimator, destGasEstimator gas.Estimator
	sourceChainID                        uint64
	builder                              BatchBuilder
	leafHasher                           LeafHasher[[32]byte]
	lggr                                 logger.Logger
	gasLimit                             uint64
}

type GEExecutionReportingPluginFactory struct {
	config GEExecutionPluginConfig
}

func NewGEExecutionReportingPluginFactory(
	config GEExecutionPluginConfig,
) types.ReportingPluginFactory {
	return &GEExecutionReportingPluginFactory{config: config}
}

func (rf *GEExecutionReportingPluginFactory) NewReportingPlugin(config types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	offchainConfig, err := Decode(config.OffchainConfig)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	return &GEExecutionReportingPlugin{
			lggr:           rf.config.lggr.Named("GEExecutionReportingPlugin"),
			F:              config.F,
			offchainConfig: offchainConfig,
			config:         rf.config,
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

type GEExecutionReportingPlugin struct {
	lggr   logger.Logger
	F      int
	config GEExecutionPluginConfig
	// We need to synchronize access to the inflight structure
	// as reporting plugin methods may be called from separate goroutines,
	// e.g. reporting vs transmission protocol.
	inFlightMu          sync.RWMutex
	inFlight            []InflightExecutionReport
	inflightCacheExpiry time.Duration
	offchainConfig      OffchainConfig
	snoozedRoots        map[[32]byte]time.Time
}

type GEQuery struct {
	TokenPrices  map[common.Address]*big.Int `json:"tokenPrices"` // TODO: We should simplify this to just link for toll as well.
	DestGasPrice *big.Int                    `json:"destGasPrice"`
}

// expect percentMultiplier to be [0, 100]
func (r *GEExecutionReportingPlugin) tokenPrices(percentMultiplier *big.Int) (map[common.Address]*big.Int, error) {
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

func (r *GEExecutionReportingPlugin) Query(ctx context.Context, timestamp types.ReportTimestamp) (types.Query, error) {
	// The leader queries an overestimated set of token prices, which are used by all the followers
	// to compute message executability, ensuring that the set of executable messages is deterministic.
	tokensPerFeeCoin, err := r.tokenPrices(big.NewInt(TokenPriceBufferPercent))
	if err != nil {
		return nil, err
	}
	// In the context of GE CCIP, latency is much less important than cost, so generally we'd prefer to wait vs bump at all.
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
	return json.Marshal(GEQuery{TokenPrices: tokensPerFeeCoin, DestGasPrice: destGasPrice.ToInt()})
}

func (r *GEExecutionReportingPlugin) Observation(ctx context.Context, timestamp types.ReportTimestamp, query types.Query) (types.Observation, error) {
	// GEQuery contains the tokenPricesPerFeeCoin
	lggr := r.lggr.Named("GEExecutionObservation")
	if isCommitStoreDownNow(lggr, r.config.commitStore) {
		return nil, ErrCommitStoreIsDown
	}
	// Expire any inflight reports.
	r.expireInflight(lggr)

	var q GEQuery
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
	// Observe a source chain price for GE pricing.
	// TODO: 1559 support
	sourceGasPrice, _, err := r.config.sourceGasEstimator.GetLegacyGas(ctx, nil, BatchGasLimit, assets.NewWei(big.NewInt(int64(MaxGasPrice))))
	if err != nil {
		return nil, err
	}
	return GEExecutionObservation{
		SeqNrs:           executableSequenceNumbers, // Note can be empty
		TokensPerFeeCoin: followerTokensPerFeeCoin,
		SourceGasPrice:   sourceGasPrice.ToInt(),
	}.Marshal()
}

func (r *GEExecutionReportingPlugin) generateFeeUpdate(token common.Address, sourceGasPrice *big.Int, juelsPerFeeCoin *big.Int) []evm_2_evm_ge_offramp.GEFeeUpdate {
	// TODO: Check gas fee updated logs
	linkPerUnitGas := big.NewInt(0).Div(big.NewInt(0).Mul(sourceGasPrice, juelsPerFeeCoin), big.NewInt(1e18))
	return []evm_2_evm_ge_offramp.GEFeeUpdate{
		{
			SourceFeeToken: token,
			// Since this gas fee update will be sent to the destination chain, this plugins
			// source chain will be the feeUpdaters destination chain.
			DestChainId: r.config.sourceChainID,
			// (juels/eth) * (wei / gas) / (1 eth / 1e18 wei) = juels/gas
			// TODO: Think more about this offchain/onchain computation split
			LinkPerUnitGas: linkPerUnitGas,
		},
	}
}

func (r *GEExecutionReportingPlugin) getExecutedSeqNrsInRange(min, max uint64) (map[uint64]struct{}, error) {
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

func (r *GEExecutionReportingPlugin) getExecutableSeqNrs(
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
	sourceTokens, err := r.config.offRamp.GetPoolTokens(nil)
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
		r.snoozedRoots[unexpiredReport.MerkleRoots[idx]] = time.Now().Add(r.config.snoozeTime)
	}
	return []uint64{}, nil
}

func (r *GEExecutionReportingPlugin) parseGESeqNr(log gethtypes.Log) (uint64, error) {
	s, err := r.config.onRamp.ParseCCIPSendRequested(log)
	if err != nil {
		return 0, err
	}
	return s.Message.SequenceNumber, nil
}

// Assumes non-empty report. Messages to execute can span more than one report, but are assumed to be in order of increasing
// sequence number.
func (r *GEExecutionReportingPlugin) buildReport(lggr logger.Logger, finalSeqNums []uint64, tokensPerFeeCoin map[common.Address]*big.Int, sourceGasPrice *big.Int) ([]byte, error) {
	execTokens, err := r.config.offRamp.GetDestinationTokens(nil)
	if err != nil {
		return nil, err
	}
	// TODO: Hack assume link is the first token
	linkToken := execTokens[0]
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
			r.parseGESeqNr,
			r.config.leafHasher.HashLeaf,
		)
		if err != nil {
			return nil, err
		}
	}
	gasFeeUpdates := r.generateFeeUpdate(linkToken, sourceGasPrice, tokensPerFeeCoin[linkToken])
	if len(gasFeeUpdates) == 0 && len(finalSeqNums) == 0 {
		return nil, errors.New("No report needed")
	}
	if len(finalSeqNums) != 0 {
		return EncodeGEExecutionReport(finalSeqNums,
			tokensPerFeeCoin,
			me.encMsgs,
			me.innerProofs,
			me.innerProofSourceFlags,
			me.outerProofs,
			me.outerProofSourceFlags,
			gasFeeUpdates,
		)
	}
	lggr.Infow("Building execution report fee update only", "feeUpdates", gasFeeUpdates)
	return EncodeGEExecutionReport(finalSeqNums, tokensPerFeeCoin, nil, nil, nil, nil, nil, gasFeeUpdates)
}

func (r *GEExecutionReportingPlugin) Report(ctx context.Context, timestamp types.ReportTimestamp, query types.Query, observations []types.AttributedObservation) (bool, types.Report, error) {
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
	var sourceGasObservations []*big.Int
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
		if obs.SourceGasPrice == nil {
			lggr.Errorw("Expect source gas price to be present")
			continue
		}
		// If it has all the prices, add each price to observations
		for token, price := range obs.TokensPerFeeCoin {
			priceObservations[token] = append(priceObservations[token], price)
		}
		// Add source gas price
		sourceGasObservations = append(sourceGasObservations, obs.SourceGasPrice)
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
	var q GEQuery
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
	report, err := r.buildReport(lggr, finalSequenceNumbers, medianTokensPerFeeCoin, median(sourceGasObservations))
	if err != nil {
		return false, nil, err
	}
	lggr.Infow("Built report",
		"onRampAddr", r.config.onRamp.Address(),
		"finalSeqNums", finalSequenceNumbers,
		"executionPrices", medianTokensPerFeeCoin)
	return true, report, nil
}

func (r *GEExecutionReportingPlugin) expireInflight(lggr logger.Logger) {
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

func (r *GEExecutionReportingPlugin) addToInflight(lggr logger.Logger, seqNrs []uint64, encMsgs [][]byte) error {
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
	r.inFlight = append(r.inFlight, InflightExecutionReport{
		createdAt:   time.Now(),
		seqNrs:      seqNrs,
		encMessages: encMsgs,
	})
	return nil
}

func (r *GEExecutionReportingPlugin) ShouldAcceptFinalizedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	lggr := r.lggr.Named("ShouldAcceptFinalizedReport")
	seqNrs, encMsgs, err := MessagesFromGEExecutionReport(report)
	if err != nil {
		lggr.Errorw("unable to decode report", "err", err)
		return false, nil
	}
	if len(seqNrs) > 0 {
		lggr.Infof("Seq nums %v", seqNrs)
		// If the first message is executed already, this execution report is stale, and we do not accept it.
		stale, err2 := r.isStaleReport(seqNrs[0])
		if err2 != nil {
			return !stale, err2
		}
		if stale {
			return false, nil
		}
	}
	// Else just assume in flight
	if err = r.addToInflight(lggr, seqNrs, encMsgs); err != nil {
		return false, err
	}
	return true, nil
}

func (r *GEExecutionReportingPlugin) ShouldTransmitAcceptedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	seqNrs, _, err := MessagesFromGEExecutionReport(report)
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

func (r *GEExecutionReportingPlugin) isStaleReport(min uint64) (bool, error) {
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

func (r *GEExecutionReportingPlugin) Close() error {
	return nil
}
