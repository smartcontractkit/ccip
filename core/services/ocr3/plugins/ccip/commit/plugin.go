package commit

import (
	"context"
	"fmt"
	"sort"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/ccipocr3/internal/libs/slicelib"
	"github.com/smartcontractkit/ccipocr3/internal/model"
	"github.com/smartcontractkit/ccipocr3/internal/reader"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

// Plugin implements the main ocr3 plugin logic.
type Plugin struct {
	nodeID     model.NodeID
	cfg        model.CommitPluginConfig
	ccipReader reader.CCIP
	lggr       logger.Logger
}

func NewPlugin(
	_ context.Context,
	nodeID model.NodeID,
	cfg model.CommitPluginConfig,
	ccipReader reader.CCIP,
) *Plugin {
	return &Plugin{
		nodeID:     nodeID,
		cfg:        cfg,
		ccipReader: ccipReader,
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
		prevOutcome, err := model.DecodeCommitPluginOutcome(outctx.PreviousOutcome)
		if err != nil {
			return types.Observation{}, fmt.Errorf("decode commit plugin previous outcome: %w", err)
		}

		for _, seqNumChain := range prevOutcome.SequenceNumbers {
			if seqNumChain.SeqNum > seqNumPerChain[seqNumChain.ChainSel] {
				seqNumPerChain[seqNumChain.ChainSel] = seqNumChain.SeqNum
			}
			knownSourceChains.Add(seqNumChain.ChainSel)
		}
	}

	knownSourceChainsSlice := knownSourceChains.ToSlice()
	sort.Slice(knownSourceChains, func(i, j int) bool { return knownSourceChainsSlice[i] < knownSourceChainsSlice[j] })

	// If reading destination chain is supported find the latest sequence numbers per chain from the onchain state.
	if p.canRead(p.cfg.DestChain) {
		onChainSeqNums, err := p.ccipReader.NextSeqNum(ctx, knownSourceChainsSlice)
		if err != nil {
			return types.Observation{}, fmt.Errorf("get next seq nums: %w", err)
		}

		for i, ch := range knownSourceChainsSlice {
			if onChainSeqNums[i] > seqNumPerChain[ch] {
				seqNumPerChain[ch] = onChainSeqNums[i]
			}
		}
	}

	// Find the new msgs for each supported chain based on the discovered sequence numbers.
	observedNewMsgs := make([]model.CCIPMsgBaseDetails, 0)
	for ch, seqNum := range seqNumPerChain {
		if !p.canRead(ch) {
			continue
		}

		minSeqNum := seqNum + 1
		maxSeqNum := minSeqNum + model.SeqNum(p.cfg.NewMsgScanBatchSize)

		newMsgs, err := p.ccipReader.MsgsBetweenSeqNums(
			ctx, []model.ChainSelector{ch}, model.NewSeqNumRange(minSeqNum, maxSeqNum))
		if err != nil {
			return nil, fmt.Errorf("get messages between seq nums: %w", err)
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
	gasPrices := make([]model.GasPriceChain, 0, len(knownSourceChainsSlice))
	for i, ch := range knownSourceChainsSlice {
		gasPrices = append(gasPrices, model.NewGasPriceChain(gasPricesVals[i], ch))
	}

	// Find the token prices.
	tokenPrices := make([]model.TokenPrice, 0) // TODO: token prices ...

	return model.NewCommitPluginObservation(p.nodeID, observedNewMsgs, gasPrices, tokenPrices).Encode()
}

func (p *Plugin) ValidateObservation(outctx ocr3types.OutcomeContext, _ types.Query, ao types.AttributedObservation) error {
	_, err := model.DecodeCommitPluginObservation(ao.Observation)
	return err
}

func (p *Plugin) ObservationQuorum(outctx ocr3types.OutcomeContext, query types.Query) (ocr3types.Quorum, error) {
	// across all chains we require at least 2f+1 observations.
	return ocr3types.QuorumTwoFPlusOne, nil
}

func (p *Plugin) Outcome(outctx ocr3types.OutcomeContext, query types.Query, aos []types.AttributedObservation) (ocr3types.Outcome, error) {
	msgsFromObservations := make([]model.CCIPMsgBaseDetails, 0)
	for _, ao := range aos {
		parsedObservation, err := model.DecodeCommitPluginObservation(ao.Observation)
		if err != nil {
			p.lggr.Errorw("decode commit plugin observation", "err", err)
			return ocr3types.Outcome{}, err
		}
		msgsFromObservations = append(msgsFromObservations, parsedObservation.NewMsgs...)
	}

	sourceChains, groupedMsgs := slicelib.GroupBy(
		msgsFromObservations, func(msg model.CCIPMsgBaseDetails) model.ChainSelector { return msg.SourceChain })
	for _, sourceChain := range sourceChains {
		p.lggr.Debugf("for source chain %d we got %d msg observations", len(groupedMsgs[sourceChain]))
	}

	return ocr3types.Outcome{}, fmt.Errorf("implement me")
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

func (p *Plugin) canRead(ch model.ChainSelector) bool {
	for _, ch := range p.cfg.Reads {
		if ch == ch {
			return true
		}
	}
	return false
}

// Interface compatibility checks.
var _ ocr3types.ReportingPlugin[[]byte] = &Plugin{}
