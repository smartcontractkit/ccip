package commit

import (
	"context"
	"fmt"
	"sort"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/ccipocr3/internal/codec"
	"github.com/smartcontractkit/ccipocr3/internal/libs/hashlib"
	"github.com/smartcontractkit/ccipocr3/internal/libs/merklemulti"
	"github.com/smartcontractkit/ccipocr3/internal/libs/slicelib"
	"github.com/smartcontractkit/ccipocr3/internal/model"
	"github.com/smartcontractkit/ccipocr3/internal/reader"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

// Plugin implements the main ocr3 ccip commit plugin logic.
// To learn more about the plugin lifecycle, see the ocr3types.ReportingPlugin interface.
type Plugin struct {
	nodeID      commontypes.OracleID
	cfg         model.CommitPluginConfig
	ccipReader  reader.CCIP
	reportCodec codec.Commit
	lggr        logger.Logger

	// readableChains is the set of chains that the plugin can read from.
	readableChains mapset.Set[model.ChainSelector]
	// knownSourceChains is the set of chains that the plugin knows about.
	knownSourceChains mapset.Set[model.ChainSelector]
}

// TODO: background service for home chain config polling

func NewPlugin(
	_ context.Context,
	nodeID commontypes.OracleID,
	cfg model.CommitPluginConfig,
	ccipReader reader.CCIP,
	reportCodec codec.Commit,
	lggr logger.Logger,
) *Plugin {
	knownSourceChains := mapset.NewSet[model.ChainSelector](cfg.Reads...)
	for _, inf := range cfg.ObserverInfo {
		knownSourceChains = knownSourceChains.Union(mapset.NewSet(inf.Reads...))
	}

	return &Plugin{
		nodeID:      nodeID,
		cfg:         cfg,
		ccipReader:  ccipReader,
		reportCodec: reportCodec,
		lggr:        lggr,

		readableChains:    mapset.NewSet(cfg.Reads...),
		knownSourceChains: knownSourceChains,
	}
}

// Query phase is not used.
func (p *Plugin) Query(_ context.Context, _ ocr3types.OutcomeContext) (types.Query, error) {
	return types.Query{}, nil
}

// Observation phase is used to discover max chain sequence numbers, new messages, gas and token prices.
//
// Max Chain Sequence Numbers:
//
//	It is the sequence number of the last known committed message for each known source chain.
//	If there was a previous outcome we start with the max sequence numbers of the previous outcome.
//	We then read the sequence numbers from the destination chain and override when the on-chain sequence number
//	is greater than previous outcome or when previous outcome did not contain a sequence number for a known source chain.
//
// New Messages:
//
//	We discover new ccip messages only for the chains that the current node is allowed to read from based on the
//	previously discovered max chain sequence numbers. For each chain we scan for new messages
//	in the [max_sequence_number+1, max_sequence_number+1+p.cfg.NewMsgScanBatchSize] range.
//
// Gas Prices:
//
//	TODO
//
// Token Prices:
//
//	TODO
func (p *Plugin) Observation(ctx context.Context, outctx ocr3types.OutcomeContext, _ types.Query) (types.Observation, error) {
	maxSeqNumsPerChain, err := p.observeMaxSeqNumsPerChain(ctx, outctx.PreviousOutcome)
	if err != nil {
		return types.Observation{}, fmt.Errorf("observe max sequence numbers per chain: %w", err)
	}

	newMsgs, err := p.observeNewMsgs(ctx, maxSeqNumsPerChain)
	if err != nil {
		return types.Observation{}, fmt.Errorf("observe new messages: %w", err)
	}

	// TODO: The code below is related to token and gas prices and should be cleaned up in relevant PRs...
	knownSourceChainsSlice := p.knownSourceChains.ToSlice()
	sort.Slice(knownSourceChainsSlice, func(i, j int) bool { return knownSourceChainsSlice[i] < knownSourceChainsSlice[j] })

	// Find the gas prices for each chain.
	gasPricesVals, err := p.ccipReader.GasPrices(ctx, knownSourceChainsSlice)
	if err != nil {
		return nil, fmt.Errorf("get gas prices: %w", err)
	}
	p.lggr.Debugw("reading gas prices", "chains",
		len(knownSourceChainsSlice), "gasPrices", len(gasPricesVals))
	gasPrices := make([]model.GasPriceChain, 0, len(knownSourceChainsSlice))
	for i, ch := range knownSourceChainsSlice {
		p.lggr.Debugw("gas price", "chain", ch, "price", gasPricesVals[i])
		gasPrices = append(gasPrices, model.NewGasPriceChain(gasPricesVals[i], ch))
	}

	// Find the token prices.
	tokenPrices := make([]model.TokenPrice, 0) // TODO: token prices ...

	p.lggr.Infow("submitting observation", "observedNewMsgs", len(newMsgs), "gasPrices", len(gasPrices), "tokenPrices", len(tokenPrices))
	return model.NewCommitPluginObservation(newMsgs, gasPrices, tokenPrices, maxSeqNumsPerChain).Encode()
}

func (p *Plugin) ValidateObservation(_ ocr3types.OutcomeContext, _ types.Query, ao types.AttributedObservation) error {
	obs, err := model.DecodeCommitPluginObservation(ao.Observation)
	if err != nil {
		return fmt.Errorf("decode commit plugin observation: %w", err)
	}

	if err := p.validateObservedSequenceNumbers(obs.NewMsgs, obs.MaxSeqNums); err != nil {
		return fmt.Errorf("validate sequence numbers: %w", err)
	}

	if err := p.validateObserverReadingEligibility(ao.Observer, obs.NewMsgs); err != nil {
		return fmt.Errorf("validate observer %d reading eligibility: %w", ao.Observer, err)
	}

	if err := p.validateObservedGasAndTokenPrices(obs.GasPrices, obs.TokenPrices); err != nil {
		return fmt.Errorf("validate gas and token prices: %w", err)
	}

	return nil
}

func (p *Plugin) ObservationQuorum(_ ocr3types.OutcomeContext, _ types.Query) (ocr3types.Quorum, error) {
	// Across all chains we require at least 2F+1 observations.
	return ocr3types.QuorumTwoFPlusOne, nil
}

// Outcome phase is used to construct the final outcome based on the observations of multiple followers.
//
// The outcome contains:
//   - Max Sequence Numbers: The max sequence number for each source chain.
//   - Merkle Roots: One merkle tree root per source chain. The leaves of the tree are the IDs of the observed messages.
//     The merkle root data type contains information about the chain and the sequence numbers range.
func (p *Plugin) Outcome(_ ocr3types.OutcomeContext, _ types.Query, aos []types.AttributedObservation) (ocr3types.Outcome, error) {
	decodedObservations := make([]model.CommitPluginObservation, 0)
	for _, ao := range aos {
		obs, err := model.DecodeCommitPluginObservation(ao.Observation)
		if err != nil {
			return ocr3types.Outcome{}, fmt.Errorf("decode commit plugin observation: %w", err)
		}
		decodedObservations = append(decodedObservations, obs)
	}

	maxSeqNumsConsensus, err := p.maxSeqNumsConsensus(decodedObservations)
	if err != nil {
		return ocr3types.Outcome{}, fmt.Errorf("max sequence numbers consensus: %w", err)
	}
	p.lggr.Debugw("max sequence numbers consensus", "maxSeqNumsConsensus", maxSeqNumsConsensus)

	merkleRoots, err := p.newMsgsConsensus(maxSeqNumsConsensus, decodedObservations)
	if err != nil {
		return ocr3types.Outcome{}, fmt.Errorf("new messages consensus: %w", err)
	}
	p.lggr.Debugw("new messages consensus", "merkleRoots", merkleRoots)

	return model.NewCommitPluginOutcome(maxSeqNumsConsensus, merkleRoots).Encode()
}

func (p *Plugin) Reports(seqNr uint64, outcome ocr3types.Outcome) ([]ocr3types.ReportWithInfo[[]byte], error) {
	outc, err := model.DecodeCommitPluginOutcome(outcome)
	if err != nil {
		return nil, fmt.Errorf("decode commit plugin outcome: %w", err)
	}

	// todo: include token price updates
	// todo: include gas price updates
	priceUpdates := make([]model.TokenPriceUpdate, 0)

	/*
		Once token/gas prices are implemented, we would want to probably check if outc.MerkleRoots is empty or not
		and only create a report if outc.MerkleRoots is non-empty OR gas/token price timer has expired
	*/
	rep := model.NewCommitPluginReport(outc.MerkleRoots, priceUpdates)
	encodedReport, err := p.reportCodec.Encode(context.Background(), rep)
	if err != nil {
		return nil, fmt.Errorf("encode commit plugin report: %w", err)
	}

	return []ocr3types.ReportWithInfo[[]byte]{{Report: encodedReport, Info: nil}}, nil
}

func (p *Plugin) ShouldAcceptAttestedReport(ctx context.Context, u uint64, r ocr3types.ReportWithInfo[[]byte]) (bool, error) {
	decodedReport, err := p.reportCodec.Decode(ctx, r.Report)
	if err != nil {
		return false, fmt.Errorf("decode commit plugin report: %w", err)
	}

	isEmpty := decodedReport.IsEmpty()
	if isEmpty {
		p.lggr.Infow("skipping empty report")
		return false, nil
	}

	return true, nil
}

func (p *Plugin) ShouldTransmitAcceptedReport(ctx context.Context, u uint64, r ocr3types.ReportWithInfo[[]byte]) (bool, error) {
	if !p.cfg.Writer {
		p.lggr.Debugw("not a writer, skipping report transmission")
		return false, nil
	}

	decodedReport, err := p.reportCodec.Decode(ctx, r.Report)
	if err != nil {
		return false, fmt.Errorf("decode commit plugin report: %w", err)
	}

	p.lggr.Debugw("transmitting report",
		"roots", len(decodedReport.MerkleRoots), "priceUpdates", len(decodedReport.PriceUpdates))

	// todo: if report is stale -> do not transmit
	return true, nil
}

func (p *Plugin) Close() error {
	timeout := 10 * time.Second
	ctx, cf := context.WithTimeout(context.Background(), timeout)
	defer cf()

	if err := p.ccipReader.Close(ctx); err != nil {
		return fmt.Errorf("close ccip reader: %w", err)
	}
	return nil
}

func (p *Plugin) observeMaxSeqNumsPerChain(ctx context.Context, previousOutcomeBytes []byte) ([]model.SeqNumChain, error) {
	// If there is a previous outcome, start with the sequence numbers of it.
	seqNumPerChain := make(map[model.ChainSelector]model.SeqNum)
	if previousOutcomeBytes != nil {
		p.lggr.Debugw("observing based on previous outcome")
		prevOutcome, err := model.DecodeCommitPluginOutcome(previousOutcomeBytes)
		if err != nil {
			return nil, fmt.Errorf("decode commit plugin previous outcome: %w", err)
		}
		p.lggr.Debugw("previous outcome decoded", "outcome", prevOutcome.String())

		for _, seqNumChain := range prevOutcome.MaxSeqNums {
			if seqNumChain.SeqNum > seqNumPerChain[seqNumChain.ChainSel] {
				seqNumPerChain[seqNumChain.ChainSel] = seqNumChain.SeqNum
			}
		}
		p.lggr.Debugw("discovered sequence numbers from prev outcome", "seqNumPerChain", seqNumPerChain)
	}

	// If reading destination chain is supported find the latest sequence numbers per chain from the onchain state.
	if p.readableChains.Contains(p.cfg.DestChain) {
		p.lggr.Debugw("reading sequence numbers from destination")
		onChainSeqNums, err := p.ccipReader.NextSeqNum(ctx, p.knownSourceChainsSlice())
		if err != nil {
			return nil, fmt.Errorf("get next seq nums: %w", err)
		}
		p.lggr.Debugw("discovered sequence numbers from destination", "onChainSeqNums", onChainSeqNums)

		// Update the seq nums if the on-chain sequence number is greater than previous outcome.
		for i, ch := range p.knownSourceChainsSlice() {
			if onChainSeqNums[i] > seqNumPerChain[ch] {
				seqNumPerChain[ch] = onChainSeqNums[i]
				p.lggr.Debugw("updated sequence number", "chain", ch, "seqNum", onChainSeqNums[i])
			}
		}
	}

	maxChainSeqNums := make([]model.SeqNumChain, 0)
	for ch, seqNum := range seqNumPerChain {
		maxChainSeqNums = append(maxChainSeqNums, model.NewSeqNumChain(ch, seqNum))
	}

	return maxChainSeqNums, nil
}

func (p *Plugin) observeNewMsgs(ctx context.Context, maxSeqNumsPerChain []model.SeqNumChain) ([]model.CCIPMsgBaseDetails, error) {
	// Find the new msgs for each supported chain based on the discovered max sequence numbers.
	observedNewMsgs := make([]model.CCIPMsgBaseDetails, 0)
	for _, seqNumChain := range maxSeqNumsPerChain {
		if !p.readableChains.Contains(seqNumChain.ChainSel) {
			p.lggr.Debugw("reading chain is not supported", "chain", seqNumChain.ChainSel)
			continue
		}

		minSeqNum := seqNumChain.SeqNum + 1
		maxSeqNum := minSeqNum + model.SeqNum(p.cfg.NewMsgScanBatchSize)
		p.lggr.Debugw("scanning for new messages",
			"chain", seqNumChain.ChainSel, "minSeqNum", minSeqNum, "maxSeqNum", maxSeqNum)

		newMsgs, err := p.ccipReader.MsgsBetweenSeqNums(
			ctx, []model.ChainSelector{seqNumChain.ChainSel}, model.NewSeqNumRange(minSeqNum, maxSeqNum))
		if err != nil {
			return nil, fmt.Errorf("get messages between seq nums: %w", err)
		}

		if len(newMsgs) > 0 {
			p.lggr.Debugw("discovered new messages", "chain", seqNumChain.ChainSel, "newMsgs", len(newMsgs))
		} else {
			p.lggr.Debugw("no new messages discovered", "chain", seqNumChain.ChainSel)
		}

		for _, msg := range newMsgs {
			observedNewMsgs = append(observedNewMsgs, msg.CCIPMsgBaseDetails)
		}
	}

	return observedNewMsgs, nil
}

func (p *Plugin) newMsgsConsensus(maxSeqNums []model.SeqNumChain, observations []model.CommitPluginObservation) ([]model.MerkleRootChain, error) {
	maxSeqNumsPerChain := make(map[model.ChainSelector]model.SeqNum)
	for _, seqNumChain := range maxSeqNums {
		maxSeqNumsPerChain[seqNumChain.ChainSel] = seqNumChain.SeqNum
	}

	// Gather all messages from all observations.
	msgsFromObservations := make([]model.CCIPMsgBaseDetails, 0)
	for _, obs := range observations {
		msgsFromObservations = append(msgsFromObservations, obs.NewMsgs...)
	}
	p.lggr.Debugw("total observed messages across all followers", "msgs", len(msgsFromObservations))

	// Filter out messages less than or equal to the max sequence numbers.
	msgsFromObservations = slicelib.Filter(msgsFromObservations, func(msg model.CCIPMsgBaseDetails) bool {
		maxSeqNum, ok := maxSeqNumsPerChain[msg.SourceChain]
		if !ok {
			return false
		}
		return msg.SeqNum > maxSeqNum
	})
	p.lggr.Debugw("observed messages after filtering", "msgs", len(msgsFromObservations))

	// Group messages by source chain.
	sourceChains, groupedMsgs := slicelib.GroupBy(
		msgsFromObservations,
		func(msg model.CCIPMsgBaseDetails) model.ChainSelector { return msg.SourceChain },
	)

	// Come to consensus on the observed messages by source chain.
	consensusBySourceChain := make(map[model.ChainSelector]observedMsgsConsensus)
	for _, sourceChain := range sourceChains { // note: we iterate using sourceChains slice for deterministic order.
		observedMsgs, ok := groupedMsgs[sourceChain]
		if !ok {
			p.lggr.Panicw("source chain not found in grouped messages", "sourceChain", sourceChain)
		}

		msgsConsensus, err := p.newMsgsConsensusForChain(sourceChain, observedMsgs)
		if err != nil {
			return nil, fmt.Errorf("calculate observed msgs consensus: %w", err)
		}

		if msgsConsensus.isEmpty() {
			p.lggr.Debugw("no consensus on observed messages", "sourceChain", sourceChain)
			continue
		}
		consensusBySourceChain[sourceChain] = msgsConsensus
		p.lggr.Debugw("observed messages consensus", "sourceChain", sourceChain, "consensus", msgsConsensus)
	}

	merkleRoots := make([]model.MerkleRootChain, 0)
	for sourceChain, consensus := range consensusBySourceChain {
		merkleRoots = append(
			merkleRoots,
			model.NewMerkleRootChain(sourceChain, consensus.seqNumRange, consensus.merkleRoot),
		)
	}
	return merkleRoots, nil
}

// Given a list of observed msgs
//   - Keep the messages that were observed by at least 2f_chain+1 followers.
//   - Starting from the first message (min seq num), keep adding the messages to the merkle tree until a gap is found.
func (p *Plugin) newMsgsConsensusForChain(chainSel model.ChainSelector, observedMsgs []model.CCIPMsgBaseDetails) (observedMsgsConsensus, error) {
	fChain, ok := p.cfg.FChain[chainSel]
	if !ok {
		return observedMsgsConsensus{}, fmt.Errorf("fchain not found for chain %d", chainSel)
	}
	p.lggr.Debugw("observed messages consensus",
		"chain", chainSel, "fChain", fChain, "observedMsgs", len(observedMsgs))

	// Reach consensus on the observed msgs sequence numbers.
	msgSeqNums := make(map[model.SeqNum]int)
	for _, msg := range observedMsgs {
		msgSeqNums[msg.SeqNum]++
		// TODO: message data might be spoofed, validate the message data
	}
	p.lggr.Debugw("observed message counts", "chain", chainSel, "msgSeqNums", msgSeqNums)

	// Filter out msgs not observed by at least 2f_chain+1 followers.
	msgSeqNumsQuorum := mapset.NewSet[model.SeqNum]()
	for seqNum, count := range msgSeqNums {
		if count >= 2*fChain+1 {
			msgSeqNumsQuorum.Add(seqNum)
		}
	}
	if msgSeqNumsQuorum.Cardinality() == 0 {
		return observedMsgsConsensus{}, nil
	}

	// Come to consensus on the observed messages sequence numbers range.
	msgSeqNumsQuorumSlice := msgSeqNumsQuorum.ToSlice()
	sort.Slice(msgSeqNumsQuorumSlice, func(i, j int) bool { return msgSeqNumsQuorumSlice[i] < msgSeqNumsQuorumSlice[j] })
	seqNumConsensusRange := model.NewSeqNumRange(msgSeqNumsQuorumSlice[0], msgSeqNumsQuorumSlice[0])
	for _, seqNum := range msgSeqNumsQuorumSlice[1:] {
		if seqNum != seqNumConsensusRange.End()+1 {
			break // Found a gap in the sequence numbers.
		}
		seqNumConsensusRange.SetEnd(seqNum)
	}

	msgsBySeqNum := make(map[model.SeqNum]model.CCIPMsgBaseDetails)
	for _, msg := range observedMsgs {
		msgsBySeqNum[msg.SeqNum] = msg
	}

	treeLeaves := make([][32]byte, 0)
	for seqNum := seqNumConsensusRange.Start(); seqNum <= seqNumConsensusRange.End(); seqNum++ {
		msg, ok := msgsBySeqNum[seqNum]
		if !ok {
			return observedMsgsConsensus{}, fmt.Errorf("msg not found in map for seq num %d", seqNum)
		}
		treeLeaves = append(treeLeaves, msg.ID)
	}

	p.lggr.Debugw("constructing merkle tree", "chain", chainSel, "treeLeaves", len(treeLeaves))
	tree, err := merklemulti.NewTree(hashlib.NewKeccakCtx(), treeLeaves)
	if err != nil {
		return observedMsgsConsensus{}, fmt.Errorf("construct merkle tree from %d leaves: %w", len(treeLeaves), err)
	}

	return observedMsgsConsensus{
		seqNumRange: seqNumConsensusRange,
		merkleRoot:  tree.Root(),
	}, nil
}

// maxSeqNumsConsensus groups the observed max seq nums across all followers per chain.
// Orders the sequence numbers and selects the one at the index of destination chain fChain.
//
// For example:
//
//	seqNums: [1, 1, 1, 10, 10, 10, 10, 10, 10]
//	fChain: 4
//	result: 10
//
// Selecting seqNums[fChain] ensures:
//   - At least one honest node has seen this value, so adversary cannot bias the value lower which would cause reverts
//   - If an honest oracle reports sorted_min[f] which happens to be stale i.e. that oracle has a delayed view
//     of the chain, then the report will revert onchain but still succeed upon retry
//   - We minimize the risk of naturally hitting the error condition minSeqNum > maxSeqNum due to oracles
//     delayed views of the chain (would be an issue with taking sorted_mins[-f])
func (p *Plugin) maxSeqNumsConsensus(observations []model.CommitPluginObservation) ([]model.SeqNumChain, error) {
	fChain, ok := p.cfg.FChain[p.cfg.DestChain]
	if !ok {
		return nil, fmt.Errorf("fchain not found for chain %d", p.cfg.DestChain)
	}

	observedSeqNumsPerChain := make(map[model.ChainSelector][]model.SeqNum)
	for _, obs := range observations {
		for _, maxSeqNum := range obs.MaxSeqNums {
			if _, exists := observedSeqNumsPerChain[maxSeqNum.ChainSel]; !exists {
				observedSeqNumsPerChain[maxSeqNum.ChainSel] = make([]model.SeqNum, 0)
			}
			observedSeqNumsPerChain[maxSeqNum.ChainSel] = append(observedSeqNumsPerChain[maxSeqNum.ChainSel], maxSeqNum.SeqNum)
		}
	}

	maxSeqNumsConsensus := make([]model.SeqNumChain, 0, len(observedSeqNumsPerChain))
	for ch, observedSeqNums := range observedSeqNumsPerChain {
		sort.Slice(observedSeqNums, func(i, j int) bool { return observedSeqNums[i] < observedSeqNums[j] })
		maxSeqNumsConsensus[ch] = model.NewSeqNumChain(ch, observedSeqNums[fChain])
	}

	return maxSeqNumsConsensus, nil
}

// validateObservedSequenceNumbers checks if the sequence numbers of the provided messages are unique for each chain and
// that they match the observed max sequence numbers.
func (p *Plugin) validateObservedSequenceNumbers(msgs []model.CCIPMsgBaseDetails, maxSeqNums []model.SeqNumChain) error {
	seqNums := make(map[model.ChainSelector]mapset.Set[model.SeqNum], len(msgs))

	// MaxSeqNums must be unique for each chain.
	maxSeqNumsMap := make(map[model.ChainSelector]model.SeqNum)
	for _, maxSeqNum := range maxSeqNums {
		if _, exists := maxSeqNumsMap[maxSeqNum.ChainSel]; exists {
			return fmt.Errorf("duplicate max sequence number for chain %d", maxSeqNum.ChainSel)
		}
		maxSeqNumsMap[maxSeqNum.ChainSel] = maxSeqNum.SeqNum
	}

	for _, msg := range msgs {
		// The same sequence number must not appear more than once for the same chain and must be valid.
		knownSeqNums, exists := seqNums[msg.SourceChain]
		if !exists {
			seqNums[msg.SourceChain] = mapset.NewSet(msg.SeqNum)
			continue
		}
		if knownSeqNums.Contains(msg.SeqNum) {
			return fmt.Errorf("duplicate sequence number %d for chain %d", msg.SeqNum, msg.SourceChain)
		}
		seqNums[msg.SourceChain].Add(msg.SeqNum)

		// The observed msg sequence number cannot be less than or equal to the max observed sequence number.
		maxSeqNum, exists := maxSeqNumsMap[msg.SourceChain]
		if !exists {
			return fmt.Errorf("max sequence number observation not found for chain %d", msg.SourceChain)
		}
		if maxSeqNum <= msg.SeqNum {
			return fmt.Errorf("max sequence number %d must be greater than observed sequence number %d for chain %d",
				maxSeqNum, msg.SeqNum, msg.SourceChain)
		}
	}

	return nil
}

// validateObserverReadingEligibility checks if the observer is eligible to observe the messages it observed.
func (p *Plugin) validateObserverReadingEligibility(observer commontypes.OracleID, msgs []model.CCIPMsgBaseDetails) error {
	if len(msgs) == 0 {
		return nil
	}

	observerInfo, exists := p.cfg.ObserverInfo[observer]
	if !exists {
		return fmt.Errorf("observer not found in config")
	}

	observerReadChains := mapset.NewSet(observerInfo.Reads...)
	p.lggr.Debugw("validating observation", "observer", observer,
		"observerReadChains", observerReadChains, "msgs", len(msgs))

	for _, msg := range msgs {
		p.lggr.Debugw("validating message", "msg", msg, "observer", observer)
		// Observer must be able to read the chain that the message is coming from.
		if !observerReadChains.Contains(msg.SourceChain) {
			return fmt.Errorf("observer not allowed to read chain %d", msg.SourceChain)
		}

		if msg.SeqNum <= 0 {
			return fmt.Errorf("sequence number must be positive")
		}
	}

	return nil
}

// validateGasAndTokenPrices checks if the provided gas and token prices are valid.
func (p *Plugin) validateObservedGasAndTokenPrices(gasPrices []model.GasPriceChain, tokenPrices []model.TokenPrice) error {
	// Duplicate gas prices must not appear for the same chain and must not be empty.
	gasPriceChains := mapset.NewSet[model.ChainSelector]()
	for _, g := range gasPrices {
		if gasPriceChains.Contains(g.ChainSel) {
			return fmt.Errorf("duplicate gas price for chain %d", g.ChainSel)
		}
		gasPriceChains.Add(g.ChainSel)
		if g.GasPrice == nil {
			return fmt.Errorf("gas price must not be nil")
		}
	}

	// Duplicate token prices must not appear for the same token and must not be empty.
	tokensWithPrice := mapset.NewSet[types.Account]()
	for _, t := range tokenPrices {
		if tokensWithPrice.Contains(t.TokenID) {
			return fmt.Errorf("duplicate token price for token: %s", t.TokenID)
		}
		tokensWithPrice.Add(t.TokenID)
		if t.Price == nil {
			return fmt.Errorf("token price must not be nil")
		}
	}

	return nil
}

func (p *Plugin) knownSourceChainsSlice() []model.ChainSelector {
	knownSourceChainsSlice := p.knownSourceChains.ToSlice()
	sort.Slice(knownSourceChainsSlice, func(i, j int) bool { return knownSourceChainsSlice[i] < knownSourceChainsSlice[j] })
	return slicelib.Filter(knownSourceChainsSlice, func(ch model.ChainSelector) bool { return ch != p.cfg.DestChain })
}

type observedMsgsConsensus struct {
	seqNumRange model.SeqNumRange
	merkleRoot  [32]byte
}

func (o observedMsgsConsensus) isEmpty() bool {
	return o.seqNumRange.Start() == 0 && o.seqNumRange.End() == 0 && o.merkleRoot == [32]byte{}
}

// Interface compatibility checks.
var _ ocr3types.ReportingPlugin[[]byte] = &Plugin{}
