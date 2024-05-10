package commit

import (
	"context"
	"fmt"
	"sort"

	mapset "github.com/deckarep/golang-set/v2"
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

// Plugin implements the main ocr3 plugin logic.
type Plugin struct {
	nodeID     commontypes.OracleID
	cfg        model.CommitPluginConfig
	ccipReader reader.CCIP
	lggr       logger.Logger

	// computed helper fields
	readChains mapset.Set[model.ChainSelector]
}

func NewPlugin(
	_ context.Context,
	nodeID commontypes.OracleID,
	cfg model.CommitPluginConfig,
	ccipReader reader.CCIP,
	lggr logger.Logger,
) *Plugin {
	return &Plugin{
		nodeID:     nodeID,
		cfg:        cfg,
		ccipReader: ccipReader,
		lggr:       lggr,

		readChains: mapset.NewSet(cfg.Reads...),
	}
}

func (p *Plugin) Query(ctx context.Context, outctx ocr3types.OutcomeContext) (types.Query, error) {
	return types.Query{}, nil
}

func (p *Plugin) Observation(ctx context.Context, outctx ocr3types.OutcomeContext, query types.Query) (types.Observation, error) {
	knownSourceChains := mapset.NewSet[model.ChainSelector](p.cfg.Reads...)
	seqNumPerChain := make(map[model.ChainSelector]model.SeqNum)

	// If there is a previous outcome, find latest sequence numbers per chain from it.
	if outctx.PreviousOutcome != nil {
		p.lggr.Debugw("observing based on previous outcome")
		prevOutcome, err := model.DecodeCommitPluginOutcome(outctx.PreviousOutcome)
		if err != nil {
			return types.Observation{}, fmt.Errorf("decode commit plugin previous outcome: %w", err)
		}
		p.lggr.Debugw("previous outcome decoded", "outcome", prevOutcome.String())

		for _, seqNumChain := range prevOutcome.MaxSequenceNumbers {
			if seqNumChain.SeqNum > seqNumPerChain[seqNumChain.ChainSel] {
				seqNumPerChain[seqNumChain.ChainSel] = seqNumChain.SeqNum
			}
			knownSourceChains.Add(seqNumChain.ChainSel)
		}
		p.lggr.Debugw("discovered source chains from prev outcome", "chains", knownSourceChains.ToSlice())
		p.lggr.Debugw("discovered sequence numbers from prev outcome", "seqNumPerChain", seqNumPerChain)
	}

	knownSourceChainsSlice := knownSourceChains.ToSlice()
	sort.Slice(knownSourceChains, func(i, j int) bool { return knownSourceChainsSlice[i] < knownSourceChainsSlice[j] })

	// If reading destination chain is supported find the latest sequence numbers per chain from the onchain state.
	if p.readChains.Contains(p.cfg.DestChain) {
		p.lggr.Debugw("reading sequence numbers from destination")
		onChainSeqNums, err := p.ccipReader.NextSeqNum(ctx, knownSourceChainsSlice)
		if err != nil {
			return types.Observation{}, fmt.Errorf("get next seq nums: %w", err)
		}
		p.lggr.Debugw("discovered sequence numbers from destination", "onChainSeqNums", onChainSeqNums)

		for i, ch := range knownSourceChainsSlice {
			if onChainSeqNums[i] > seqNumPerChain[ch] {
				seqNumPerChain[ch] = onChainSeqNums[i]
				p.lggr.Debugw("updated sequence number", "chain", ch, "seqNum", onChainSeqNums[i])
			}
		}
	}

	// Find the new msgs for each supported chain based on the discovered sequence numbers.
	observedNewMsgs := make([]model.CCIPMsgBaseDetails, 0)
	for ch, seqNum := range seqNumPerChain {
		if !p.readChains.Contains(ch) {
			p.lggr.Debugw("reading chain is not supported", "chain", ch)
			continue
		}

		minSeqNum := seqNum + 1
		maxSeqNum := minSeqNum + model.SeqNum(p.cfg.NewMsgScanBatchSize)
		p.lggr.Debugw("scanning for new messages",
			"chain", ch, "minSeqNum", minSeqNum, "maxSeqNum", maxSeqNum)

		newMsgs, err := p.ccipReader.MsgsBetweenSeqNums(
			ctx, []model.ChainSelector{ch}, model.NewSeqNumRange(minSeqNum, maxSeqNum))
		if err != nil {
			return nil, fmt.Errorf("get messages between seq nums: %w", err)
		}

		if len(newMsgs) > 0 {
			p.lggr.Debugw("discovered new messages", "chain", ch, "newMsgs", len(newMsgs))
		} else {
			p.lggr.Debugw("no new messages discovered", "chain", ch)
		}

		for _, msg := range newMsgs {
			observedNewMsgs = append(observedNewMsgs, msg.CCIPMsgBaseDetails)
		}
	}

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

	p.lggr.Infow("submitting observation", "observedNewMsgs", len(observedNewMsgs), "gasPrices", len(gasPrices), "tokenPrices", len(tokenPrices))
	return model.NewCommitPluginObservation(observedNewMsgs, gasPrices, tokenPrices).Encode()
}

func (p *Plugin) ValidateObservation(_ ocr3types.OutcomeContext, _ types.Query, ao types.AttributedObservation) error {
	obs, err := model.DecodeCommitPluginObservation(ao.Observation)
	if err != nil {
		return fmt.Errorf("decode commit plugin observation: %w", err)
	}

	if err := p.validateObservedSequenceNumbersUniqueness(obs.NewMsgs); err != nil {
		return fmt.Errorf("validate sequence numbers uniqueness: %w", err)
	}

	if err := p.validateObserverReadingEligibility(ao.Observer, obs.NewMsgs); err != nil {
		return fmt.Errorf("validate observer %d eligibility: %w", ao.Observer, err)
	}

	if err := p.validateObservedGasAndTokenPrices(obs.GasPrices, obs.TokenPrices); err != nil {
		return fmt.Errorf("validate gas and token prices: %w", err)
	}

	return nil
}

func (p *Plugin) ObservationQuorum(_ ocr3types.OutcomeContext, _ types.Query) (ocr3types.Quorum, error) {
	// across all chains we require at least 2f+1 observations.
	return ocr3types.QuorumTwoFPlusOne, nil
}

func (p *Plugin) Outcome(_ ocr3types.OutcomeContext, _ types.Query, aos []types.AttributedObservation) (ocr3types.Outcome, error) {
	observerSeqNums := make(map[commontypes.OracleID]map[model.ChainSelector]mapset.Set[model.SeqNum])
	msgsFromObservations := make([]model.CCIPMsgBaseDetails, 0)

	p.lggr.Debugw("calculating outcome", "observations", len(aos))
	for _, ao := range aos {
		obs, err := model.DecodeCommitPluginObservation(ao.Observation)
		if err != nil {
			return ocr3types.Outcome{}, fmt.Errorf("decode commit plugin observation: %w", err)
		}

		p.lggr.Debugw("processing observation", "observer", ao.Observer, "msgs", len(obs.NewMsgs))
		// Ignore observations that have duplicate sequence numbers coming from the same sender for the same chain.
		for _, msg := range obs.NewMsgs {
			if _, exists := observerSeqNums[ao.Observer]; !exists {
				observerSeqNums[ao.Observer] = map[model.ChainSelector]mapset.Set[model.SeqNum]{}
			}
			if _, exists := observerSeqNums[ao.Observer][msg.SourceChain]; !exists {
				observerSeqNums[ao.Observer][msg.SourceChain] = mapset.NewSet[model.SeqNum]()
			}
			if observerSeqNums[ao.Observer][msg.SourceChain].Contains(msg.SeqNum) {
				p.lggr.Warnw("duplicate follower sequence number in observation",
					"observer", ao.Observer, "chain", msg.SourceChain, "seqNum", msg.SeqNum)
				continue
			}
			observerSeqNums[ao.Observer][msg.SourceChain].Add(msg.SeqNum)
		}

		msgsFromObservations = append(msgsFromObservations, obs.NewMsgs...)
	}
	p.lggr.Debugw("total observed messages across all followers", "msgs", len(msgsFromObservations))

	// Group messages by source chain.
	sourceChains, groupedMsgs := slicelib.GroupBy(
		msgsFromObservations, func(msg model.CCIPMsgBaseDetails) model.ChainSelector { return msg.SourceChain })
	for _, sourceChain := range sourceChains {
		p.lggr.Debugw("grouped messages by source chain", "sourceChain", sourceChain, "msgs", len(groupedMsgs[sourceChain]))
	}

	// Come to consensus on the observed messages by source chain.
	consensusBySourceChain := make(map[model.ChainSelector]observedMsgsConsensus)
	for _, sourceChain := range sourceChains {
		observedMsgs, ok := groupedMsgs[sourceChain]
		if !ok {
			p.lggr.Panicw("source chain not found in grouped messages", "sourceChain", sourceChain)
		}

		msgsConsensus, err := p.observedMsgsConsensus(sourceChain, observedMsgs)
		if err != nil {
			return ocr3types.Outcome{}, fmt.Errorf("calculate observed msgs consensus: %w", err)
		}
		consensusBySourceChain[sourceChain] = msgsConsensus
		p.lggr.Debugw("observed messages consensus", "sourceChain", sourceChain, "consensus", msgsConsensus)

		/*
			# TODO: Intention of the expiry is to prevent outctx.max_committed
			# from getting permanently out of sync with the dest chain if something goes wrong. Maybe there's a better way?
			# If its expired, we update from the consensus destination chain.
			if expired(outctx.max_committed_seq_nr_by_source):
				max_committed_seq_nr_by_source = calculate_consensus_committed_seq_nr(observations_by_source[chain], quorum)
		*/
	}

	// Construct the outcome.
	maxSeqNums := make([]model.SeqNumChain, 0)
	merkleRoots := make([]model.MerkleRootChain, 0)
	for sourceChain, consensus := range consensusBySourceChain {
		maxSeqNums = append(maxSeqNums, model.NewSeqNumChain(sourceChain, consensus.seqNumRange.End()))
		merkleRoots = append(merkleRoots, model.NewMerkleRootChain(sourceChain, consensus.merkleRoot))
	}
	return model.NewCommitPluginOutcome(maxSeqNums, merkleRoots).Encode()
}

func (p *Plugin) Reports(seqNr uint64, outcome ocr3types.Outcome) ([]ocr3types.ReportWithInfo[[]byte], error) {
	/*
		# Only a single report is needed containing a batch of gas price updates for the chain.
		# and a list of roots by source. We only include the gas price batch update
		# if a timer has expired - ie periodically we batch write all gas prices.
	*/
	panic("implement me")
}

func (p *Plugin) ShouldAcceptAttestedReport(ctx context.Context, u uint64, r ocr3types.ReportWithInfo[[]byte]) (bool, error) {
	panic("implement me")
}

func (p *Plugin) ShouldTransmitAcceptedReport(ctx context.Context, u uint64, r ocr3types.ReportWithInfo[[]byte]) (bool, error) {
	/*
		if dest not in self.chains:
			# Can't write, skip
			return False
		return self.chains[dest].is_stale(report):
	*/
	panic("implement me")
}

func (p *Plugin) Close() error {
	panic("implement me")
}

func (p *Plugin) observedMsgsConsensus(chainSel model.ChainSelector, observedMsgs []model.CCIPMsgBaseDetails) (observedMsgsConsensus, error) {
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
	}
	p.lggr.Debugw("observed message counts", "chain", chainSel, "msgSeqNums", msgSeqNums)

	// Filter out msgs not observed by at least 2f_chain+1 followers.
	msgSeqNumsQuorum := mapset.NewSet[model.SeqNum]()
	for seqNum, count := range msgSeqNums {
		if count >= 2*fChain+1 {
			msgSeqNumsQuorum.Add(seqNum)
		}
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
		msgsBySeqNum[msg.SeqNum] = msg // todo: validate that all msgs are the same
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

	// TODO: gas price consensus
	// TODO: token prices consensus

	return observedMsgsConsensus{
		seqNumRange: seqNumConsensusRange,
		merkleRoot:  tree.Root(),
	}, nil
}

// validateObservedSequenceNumbersUniqueness checks if the sequence numbers of the provided messages are unique for each chain.
func (p *Plugin) validateObservedSequenceNumbersUniqueness(msgs []model.CCIPMsgBaseDetails) error {
	seqNums := make(map[model.ChainSelector]mapset.Set[model.SeqNum], len(msgs))

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
	tokensWithPrice := mapset.NewSet[string]()
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

type observedMsgsConsensus struct {
	seqNumRange model.SeqNumRange
	merkleRoot  [32]byte
}

// Interface compatibility checks.
var _ ocr3types.ReportingPlugin[[]byte] = &Plugin{}
