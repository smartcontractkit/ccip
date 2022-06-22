package ccip

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/any_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/blob_verifier"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/merklemulti"
	"github.com/smartcontractkit/chainlink/core/utils/mathutil"
)

const (
	MessageStateUntouched = iota
	MessageStateInProgress
	MessageStateSuccess
	MessageStateFailure
)

const (
	BatchGasLimit                   = 1_000_000 // TODO: think if a good value for this
	RootSnoozeTime                  = 1 * 3600 * time.Second
	ExecutionMaxInflightTimeSeconds = 180
	// Note user research is required for setting (MaxPayloadLength, MaxTokensPerMessage).
	// TODO: If we really want this to be constant and not dynamic, then we need to wait
	// until we have gas limits per message and ensure the block gas limit constraint is respected
	// as well as the tx size limit.
	MaxNumMessagesInExecutionReport = 50
	MaxPayloadLength                = 1000
	MaxTokensPerMessage             = 5
	MaxExecutionReportLength        = 150_000 // TODO
	MaxGasPrice                     = 200_000 // Gwei, TODO: probably want this to be some dynamic value, a multiplier of the current gas price.
)

var (
	_                     types.ReportingPluginFactory = &ExecutionReportingPluginFactory{}
	_                     types.ReportingPlugin        = &ExecutionReportingPlugin{}
	ErrBlobVerifierIsDown                              = errors.New("blobVerifier is down")
)

func EncodeExecutionReport(seqNums []uint64, tokensPerFeeCoin map[common.Address]uint64, msgs [][]byte, innerProofs [][32]byte, innerProofSourceFlags []bool, outerProofs [][32]byte, outerProofSourceFlags []bool) (types.Report, error) {
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
		tokensPerFeeCoinValues = append(tokensPerFeeCoinValues, big.NewInt(int64(tokensPerFeeCoin[addr])))
	}
	report, err := makeExecutionReportArgs().PackValues([]interface{}{&any_2_evm_toll_offramp.CCIPExecutionReport{
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

func DecodeExecutionReport(report types.Report) (*any_2_evm_toll_offramp.CCIPExecutionReport, error) {
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
		EncodedMessages          [][]byte         `json:"encodedMessages"`
		InnerProofs              [][32]uint8      `json:"innerProofs"`
		InnerProofFlagBits       *big.Int         `json:"innerProofFlagBits"`
		OuterProofs              [][32]uint8      `json:"outerProofs"`
		OuterProofFlagBits       *big.Int         `json:"outerProofFlagBits"`
	})
	if !ok {
		return nil, fmt.Errorf("got %T", unpacked[0])
	}
	if len(erStruct.EncodedMessages) == 0 {
		return nil, errors.New("assumptionViolation: expected at least one element")
	}
	var er any_2_evm_toll_offramp.CCIPExecutionReport
	for _, msg := range erStruct.EncodedMessages {
		er.EncodedMessages = append(er.EncodedMessages, msg)
	}
	for _, proof := range erStruct.InnerProofs {
		er.InnerProofs = append(er.InnerProofs, proof)
	}
	for _, proof := range erStruct.OuterProofs {
		er.OuterProofs = append(er.OuterProofs, proof)
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

type ExecutionReportingPluginFactory struct {
	lggr                logger.Logger
	source, dest        logpoller.LogPoller
	onRamp, offRampAddr common.Address
	offRamp             OffRamp
	blobVerifier        *blob_verifier.BlobVerifier
	builder             BatchBuilder
	onRampSeqParser     func(log logpoller.Log) (uint64, error)
}

func NewExecutionReportingPluginFactory(
	lggr logger.Logger,
	onRamp common.Address,
	blobVerifier *blob_verifier.BlobVerifier,
	source, dest logpoller.LogPoller,
	offRampAddr common.Address,
	offRamp OffRamp,
	builder BatchBuilder,
	onRampSeqParser func(log logpoller.Log) (uint64, error),
) types.ReportingPluginFactory {
	return &ExecutionReportingPluginFactory{lggr: lggr, onRamp: onRamp, blobVerifier: blobVerifier, offRamp: offRamp, source: source, dest: dest, offRampAddr: offRampAddr, builder: builder, onRampSeqParser: onRampSeqParser}
}

type dummyDataSource struct{}

func (d dummyDataSource) GetPrice(address common.Address) (uint64, error) {
	// TODO: Actually query data source. For now just return a juels/ETH value.
	//  0.006 link/eth or 6e15 juels/eth
	return 6000000000000000, nil
}

func (rf *ExecutionReportingPluginFactory) NewReportingPlugin(config types.ReportingPluginConfig) (types.ReportingPlugin, types.ReportingPluginInfo, error) {
	offchainConfig, err := Decode(config.OffchainConfig)
	if err != nil {
		return nil, types.ReportingPluginInfo{}, err
	}
	return &ExecutionReportingPlugin{
			lggr:            rf.lggr.Named("ExecutionReportingPlugin"),
			F:               config.F,
			offRampAddr:     rf.offRampAddr,
			offRamp:         rf.offRamp,
			onRamp:          rf.onRamp,
			blobVerifier:    rf.blobVerifier,
			source:          rf.source,
			dest:            rf.dest,
			offchainConfig:  offchainConfig,
			builder:         NewExecutionBatchBuilder(BatchGasLimit, RootSnoozeTime, rf.blobVerifier, rf.onRamp, rf.offRampAddr, rf.source, rf.dest, rf.builder, offchainConfig, rf.offRamp),
			onRampSeqParser: rf.onRampSeqParser,
			dataSource:      dummyDataSource{},
		}, types.ReportingPluginInfo{
			Name:          "CCIPExecution",
			UniqueReports: true,
			Limits: types.ReportingPluginLimits{
				MaxQueryLength:       0,
				MaxObservationLength: MaxObservationLength,
				MaxReportLength:      MaxExecutionReportLength,
			},
		}, nil
}

type ExecutionReportingPlugin struct {
	lggr         logger.Logger
	F            int
	offRampAddr  common.Address
	onRamp       common.Address
	offRamp      OffRamp
	blobVerifier *blob_verifier.BlobVerifier
	source, dest logpoller.LogPoller
	// We need to synchronize access to the inflight structure
	// as reporting plugin methods may be called from separate goroutines,
	// e.g. reporting vs transmission protocol.
	inFlightMu      sync.RWMutex
	inFlight        []InflightExecutionReport
	offchainConfig  OffchainConfig
	builder         *ExecutionBatchBuilder
	onRampSeqParser func(log logpoller.Log) (uint64, error)
	dataSource      DataSource
}

type DataSource interface {
	GetPrice(token common.Address) (uint64, error)
}

type InflightExecutionReport struct {
	createdAt time.Time
	report    any_2_evm_toll_offramp.CCIPExecutionReport
}

func (r *ExecutionReportingPlugin) Query(ctx context.Context, timestamp types.ReportTimestamp) (types.Query, error) {
	// We don't use a query for this reporting plugin, so we can just leave it empty here
	return types.Query{}, nil
}

func (r *ExecutionReportingPlugin) inflightSeqNums() map[uint64]struct{} {
	r.inFlightMu.RLock()
	defer r.inFlightMu.RUnlock()
	inFlightSeqNums := make(map[uint64]struct{})
	for _, report := range r.inFlight {
		for _, seqNr := range report.report.SequenceNumbers {
			inFlightSeqNums[seqNr] = struct{}{}
		}
	}
	return inFlightSeqNums
}

func (r *ExecutionReportingPlugin) maxGasPrice() uint64 {
	return MaxGasPrice
}

func (r *ExecutionReportingPlugin) Observation(ctx context.Context, timestamp types.ReportTimestamp, query types.Query) (types.Observation, error) {
	lggr := r.lggr.Named("ExecutionObservation")
	if isBlobVerifierDownNow(r.lggr, r.blobVerifier) {
		return nil, ErrBlobVerifierIsDown
	}
	tokens, err := r.offRamp.GetPoolTokens(nil)
	if err != nil {
		return nil, err
	}
	tokensPerFeeCoin := make(map[common.Address]uint64)
	for _, token := range tokens {
		price, err2 := r.dataSource.GetPrice(token)
		if err2 != nil {
			return nil, err2
		}
		tokensPerFeeCoin[token] = price
	}
	// Read and make a copy for the builder.
	r.inFlightMu.RLock()
	inFlight := make([]InflightExecutionReport, len(r.inFlight))
	copy(inFlight[:], r.inFlight[:])
	r.inFlightMu.RUnlock()
	executableSequenceNumbers, err := r.builder.getExecutableSeqNrs(r.maxGasPrice(), tokensPerFeeCoin, inFlight)
	if err != nil {
		return nil, err
	}
	if len(executableSequenceNumbers) == 0 {
		return []byte{}, nil
	}
	lggr.Infof("executable seq nums %v", executableSequenceNumbers)
	return ExecutionObservation{
		SeqNrs:           executableSequenceNumbers,
		TokensPerFeeCoin: tokensPerFeeCoin,
	}.Marshal()
}

func contiguousReqs(lggr logger.Logger, min, max uint64, seqNrs []uint64) bool {
	for i, j := min, 0; i < max && j < len(seqNrs); i, j = i+1, j+1 {
		if seqNrs[j] != i {
			lggr.Errorw("unexpected gap in seq nums", "seq", i)
			return false
		}
	}
	return true
}

func leafsFromIntervals(lggr logger.Logger, seqParsers map[common.Address]func(logpoller.Log) (uint64, error), intervalByOnRamp map[common.Address]blob_verifier.CCIPInterval, srcLogPoller logpoller.LogPoller) (map[common.Address][][32]byte, error) {
	leafsByOnRamp := make(map[common.Address][][32]byte)
	for onRamp, interval := range intervalByOnRamp {
		// Logs are guaranteed to be in order of seq num, since these are finalized logs only
		// and the contract's seq num is auto-incrementing.
		// TODO: Could be different event_sig/index per onRamp
		logs, err := srcLogPoller.LogsDataWordRange(CCIPSendRequested, onRamp, SendRequestedSequenceNumberIndex, logpoller.EvmWord(interval.Min), logpoller.EvmWord(interval.Max), 1)
		if err != nil {
			return nil, err
		}
		var seqNrs []uint64
		for _, log := range logs {
			seqNr, err := seqParsers[onRamp](log)
			if err != nil {
				return nil, err
			}
			seqNrs = append(seqNrs, seqNr)
		}
		if !contiguousReqs(lggr, interval.Min, interval.Max, seqNrs) {
			return nil, errors.Errorf("do not have full range [%v, %v] have %v", interval.Min, interval.Max, seqNrs)
		}
		var leafs [][32]byte
		for _, log := range logs {
			// TODO: Hasher
			ctx := merklemulti.NewKeccakCtx()
			leafs = append(leafs, ctx.HashLeaf(log.Data))
		}
		leafsByOnRamp[onRamp] = leafs
	}
	return leafsByOnRamp, nil
}

// Assumes non-empty report. Messages to execute can be span more than one report, but are assumed to be in order of increasing
// sequence number.
func (r *ExecutionReportingPlugin) buildReport(lggr logger.Logger, finalSeqNums []uint64, tokensPerFeeCoin map[common.Address]uint64) ([]byte, error) {
	rep, err := r.builder.relayedReport(finalSeqNums[0])
	if err != nil {
		return nil, err
	}
	intervalsByOnRamp := make(map[common.Address]blob_verifier.CCIPInterval)
	merkleRootsByOnRamp := make(map[common.Address][32]byte)
	for i, onRamp := range rep.OnRamps {
		intervalsByOnRamp[onRamp] = rep.Intervals[i]
		merkleRootsByOnRamp[onRamp] = rep.MerkleRoots[i]
	}
	interval := intervalsByOnRamp[r.onRamp]
	msgsInRoot, err := r.source.LogsDataWordRange(CCIPSendRequested, r.onRamp, SendRequestedSequenceNumberIndex, EvmWord(interval.Min), EvmWord(interval.Max), int(r.offchainConfig.SourceIncomingConfirmations))
	if err != nil {
		return nil, err
	}
	leafsByOnRamp, err := leafsFromIntervals(r.lggr, map[common.Address]func(log logpoller.Log) (uint64, error){
		r.onRamp: r.onRampSeqParser,
	}, map[common.Address]blob_verifier.CCIPInterval{
		r.onRamp: intervalsByOnRamp[r.onRamp],
	}, r.source)
	if err != nil {
		return nil, err
	}
	var outerTreeLeafs [][32]byte
	var onRampIdx int
	for i, onRamp := range rep.OnRamps {
		if onRamp == r.onRamp {
			onRampIdx = i
		}
		outerTreeLeafs = append(outerTreeLeafs, merkleRootsByOnRamp[onRamp])
	}
	ctx := merklemulti.NewKeccakCtx()
	outerTree := merklemulti.NewTree[[32]byte](ctx, outerTreeLeafs)
	outerProof := outerTree.Prove([]int{onRampIdx})
	innerTree := merklemulti.NewTree[[32]byte](ctx, leafsByOnRamp[r.onRamp])
	var innerIdxs []int
	var encMsgs [][]byte
	var hashes [][32]byte
	for _, seqNum := range finalSeqNums {
		innerIdx := int(seqNum - intervalsByOnRamp[r.onRamp].Min)
		innerIdxs = append(innerIdxs, innerIdx)
		// TODO: Use hasher
		encMsgs = append(encMsgs, msgsInRoot[innerIdx].Data)
		hashes = append(hashes, ctx.HashLeaf(msgsInRoot[innerIdx].Data))
	}
	innerProof := innerTree.Prove(innerIdxs)
	// Double check this verifies before sending.
	res, err := r.blobVerifier.Verify(nil, hashes, innerProof.Hashes, ProofFlagsToBits(innerProof.SourceFlags), outerProof.Hashes, ProofFlagsToBits(outerProof.SourceFlags))
	if err != nil {
		return nil, err
	}
	// No timestamp, means failed to verify root.
	if res.Cmp(big.NewInt(0)) == 0 {
		ir := innerTree.Root()
		or := outerTree.Root()
		r.lggr.Errorf("Root does not verify: our inner root %v our outer root %v contract outer root %v",
			hexutil.Encode(ir[:]), hexutil.Encode(or[:]), rep.RootOfRoots)
		return nil, errors.New("root does not verify")
	}
	er, err := EncodeExecutionReport(finalSeqNums, tokensPerFeeCoin, encMsgs, innerProof.Hashes, innerProof.SourceFlags, outerProof.Hashes, outerProof.SourceFlags)
	if err != nil {
		return nil, err
	}
	return er, nil
}

func median[T uint64](vals []T) T {
	valsCopy := make([]T, len(vals))
	copy(valsCopy[:], vals[:])
	sort.Slice(valsCopy, func(i, j int) bool {
		return valsCopy[i] < valsCopy[j]
	})
	return valsCopy[len(valsCopy)/2]
}

func (r *ExecutionReportingPlugin) Report(ctx context.Context, timestamp types.ReportTimestamp, query types.Query, observations []types.AttributedObservation) (bool, types.Report, error) {
	lggr := r.lggr.Named("Report")
	if isBlobVerifierDownNow(lggr, r.blobVerifier) {
		return false, nil, ErrBlobVerifierIsDown
	}
	actualMaybeObservations := getNonEmptyObservations[ExecutionObservation](r.lggr, observations)
	var actualObservations []ExecutionObservation
	tokens, err := r.offRamp.GetPoolTokens(nil)
	if err != nil {
		return false, nil, err
	}
	priceObservations := make(map[common.Address][]uint64)
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
		// TODO: Implement according to https://github.com/smartcontractkit/ccip-spec/issues/71
		for token, price := range obs.TokensPerFeeCoin {
			priceObservations[token] = append(priceObservations[token], price)
		}
		actualObservations = append(actualObservations, obs)
	}
	// Need at least F+1 observations
	if len(actualObservations) <= r.F {
		lggr.Tracew("Non-empty observations <= F, need at least F+1 to continue")
		return false, nil, nil
	}
	tokensPerFeeCoin := make(map[common.Address]uint64)
	for _, token := range tokens {
		tokensPerFeeCoin[token] = median[uint64](priceObservations[token])
	}
	tally := make(map[uint64]int)
	for _, obs := range actualObservations {
		for _, seq_nr := range obs.SeqNrs {
			tally[seq_nr]++
		}
	}
	// TODO: Will change in https://github.com/smartcontractkit/ccip-spec/issues/71
	var finalSequenceNumbers []uint64
	for seqNr, count := range tally {
		if count > r.F {
			finalSequenceNumbers = append(finalSequenceNumbers, seqNr)
		}
	}
	nextMin, err := r.blobVerifier.SExpectedNextMinByOnRamp(nil, r.onRamp)
	if err != nil {
		return false, nil, err
	}
	if mathutil.Max(finalSequenceNumbers[0], finalSequenceNumbers[1:]...) >= nextMin {
		return false, nil, errors.Errorf("Cannot execute unrelayed seq num. nextMin %v, seqNums %v", nextMin, finalSequenceNumbers)
	}
	report, err := r.buildReport(lggr, finalSequenceNumbers, tokensPerFeeCoin)
	if err != nil {
		return false, nil, err
	}
	r.lggr.Infow("Built report", "onRamp", r.onRamp, "finalSeqNums", finalSequenceNumbers)
	return true, report, nil
}

func (r *ExecutionReportingPlugin) updateInFlight(lggr logger.Logger, er any_2_evm_toll_offramp.CCIPExecutionReport) error {
	r.inFlightMu.Lock()
	defer r.inFlightMu.Unlock()
	// Reap old inflights and check if any messages in the report are inflight.
	var stillInFlight []InflightExecutionReport
	for _, report := range r.inFlight {
		// TODO: Think about if this fails in reorgs
		if report.report.SequenceNumbers[0] == er.SequenceNumbers[0] {
			return errors.Errorf("report is already in flight")
		}
		if time.Since(report.createdAt) < ExecutionMaxInflightTimeSeconds {
			stillInFlight = append(stillInFlight, report)
		} else {
			lggr.Warnw("Inflight report expired, retrying", "min", report.report.SequenceNumbers[0], "max", report.report.SequenceNumbers[len(report.report.SequenceNumbers)-1])
		}
	}
	// Add new inflight
	r.inFlight = append(stillInFlight, InflightExecutionReport{
		createdAt: time.Now(),
		report:    er,
	})
	return nil
}

func (r *ExecutionReportingPlugin) ShouldAcceptFinalizedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	lggr := r.lggr.Named("ShouldAcceptFinalizedReport")
	er, err := DecodeExecutionReport(report)
	if err != nil {
		r.lggr.Errorw("unable to decode report", "err", err)
		return false, nil
	}
	if len(er.SequenceNumbers) == 0 {
		r.lggr.Warnw("Received empty report")
		return false, nil
	}
	lggr.Infof("Seq nums %v", er.SequenceNumbers)
	// If the first message is executed already, this execution report is stale, and we do not accept it.
	stale, err := r.isStaleReport(er.SequenceNumbers[0])
	if err != nil {
		return !stale, err
	}
	if stale {
		return false, err
	}
	if err := r.updateInFlight(lggr, *er); err != nil {
		return false, err
	}
	return true, nil
}

func (r *ExecutionReportingPlugin) ShouldTransmitAcceptedReport(ctx context.Context, timestamp types.ReportTimestamp, report types.Report) (bool, error) {
	parsedReport, err := DecodeExecutionReport(report)
	if err != nil {
		return false, nil
	}
	// If report is not stale we transmit.
	// When the executeTransmitter enqueues the tx for bptxm,
	// we mark it as execution_sent, removing it from the set of inflight messages.
	stale, err := r.isStaleReport(parsedReport.SequenceNumbers[0])
	return !stale, err
}

func (r *ExecutionReportingPlugin) isStaleReport(min uint64) (bool, error) {
	// If the first message is executed already, this execution report is stale.
	msgState, err := r.offRamp.ExecutedMessages(nil, min)
	if err != nil {
		// TODO: do we need to check for not present error?
		return true, err
	}
	if msgState == MessageStateFailure || msgState == MessageStateSuccess {
		return true, nil
	}
	return false, nil
}

func (r *ExecutionReportingPlugin) Close() error {
	return nil
}
