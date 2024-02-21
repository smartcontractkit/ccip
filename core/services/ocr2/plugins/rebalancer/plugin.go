package rebalancer

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"go.uber.org/multierr"

	ubig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/bridge"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/discoverer"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/graph"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/liquiditymanager"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/liquidityrebalancer"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

const (
	cacheExpiryTime = 1 * time.Minute
)

type Plugin struct {
	f                       int
	rootNetwork             models.NetworkSelector
	rootAddress             models.Address
	closePluginTimeout      time.Duration
	liquidityManagerFactory liquiditymanager.Factory
	discovererFactory       discoverer.Factory
	bridgeFactory           bridge.Factory
	mu                      sync.RWMutex
	rebalancerGraph         graph.Graph
	liquidityRebalancer     liquidityrebalancer.Rebalancer
	inflightCache           InflightCache
	lggr                    logger.Logger
}

func NewPlugin(
	f int,
	closePluginTimeout time.Duration,
	rootNetwork models.NetworkSelector,
	rootAddress models.Address,
	liquidityManagerFactory liquiditymanager.Factory,
	discovererFactory discoverer.Factory,
	bridgeFactory bridge.Factory,
	liquidityRebalancer liquidityrebalancer.Rebalancer,
	lggr logger.Logger,
) *Plugin {
	return &Plugin{
		f:                       f,
		rootNetwork:             rootNetwork,
		rootAddress:             rootAddress,
		closePluginTimeout:      closePluginTimeout,
		liquidityManagerFactory: liquidityManagerFactory,
		discovererFactory:       discovererFactory,
		bridgeFactory:           bridgeFactory,
		rebalancerGraph:         graph.NewGraph(),
		liquidityRebalancer:     liquidityRebalancer,
		lggr:                    lggr,
		mu:                      sync.RWMutex{},
		inflightCache:           NewInflightCache(cacheExpiryTime),
	}
}

func (p *Plugin) Query(_ context.Context, outcomeCtx ocr3types.OutcomeContext) (ocrtypes.Query, error) {
	p.lggr.Infow("in query", "seqNr", outcomeCtx.SeqNr)
	return ocrtypes.Query{}, nil
}

func (p *Plugin) Observation(ctx context.Context, outcomeCtx ocr3types.OutcomeContext, _ ocrtypes.Query) (ocrtypes.Observation, error) {
	lggr := p.lggr.With("seqNr", outcomeCtx.SeqNr, "phase", "Observation")
	lggr.Infow("in observation", "seqNr", outcomeCtx.SeqNr)

	p.inflightCache.Expire(lggr)

	inflight := p.inflightCache.Get()

	if err := p.syncGraphEdges(ctx); err != nil {
		return ocrtypes.Observation{}, fmt.Errorf("sync graph edges: %w", err)
	}

	networkLiquidities, err := p.syncGraphBalances(ctx)
	if err != nil {
		return ocrtypes.Observation{}, fmt.Errorf("sync graph balances: %w", err)
	}

	pendingTransfers, err := p.loadPendingTransfers(ctx)
	if err != nil {
		return ocrtypes.Observation{}, fmt.Errorf("load pending transfers: %w", err)
	}

	edges, err := p.rebalancerGraph.GetEdges()
	if err != nil {
		return ocrtypes.Observation{}, fmt.Errorf("get edges: %w", err)
	}

	resolvedTransfers, err := p.resolveProposedTransfers(ctx, outcomeCtx)
	if err != nil {
		return ocrtypes.Observation{}, fmt.Errorf("resolve proposed transfers: %w", err)
	}

	configDigests := make([]models.ConfigDigestWithMeta, 0)
	for _, net := range p.rebalancerGraph.GetNetworks() {
		data, err := p.rebalancerGraph.GetData(net)
		if err != nil {
			return nil, fmt.Errorf("get rb %d data: %w", net, err)
		}
		configDigests = append(configDigests, models.ConfigDigestWithMeta{
			Digest:         data.ConfigDigest,
			NetworkSel:     data.NetworkSelector,
			RebalancerAddr: data.RebalancerAddress,
		})
	}

	lggr.Infow("finished observing",
		"networkLiquidities", networkLiquidities,
		"pendingTransfers", pendingTransfers,
		"edges", edges,
		"resolvedTransfers", resolvedTransfers,
		"inflight", inflight,
	)

	return models.NewObservation(networkLiquidities, resolvedTransfers, pendingTransfers, inflight, edges, configDigests).Encode(), nil
}

func (p *Plugin) ValidateObservation(outctx ocr3types.OutcomeContext, query ocrtypes.Query, ao ocrtypes.AttributedObservation) error {
	p.lggr.Infow("in validate observation", "seqNr", outctx.SeqNr, "phase", "ValidateObservation")

	_, err := models.DecodeObservation(ao.Observation)
	if err != nil {
		return fmt.Errorf("invalid observation: %w", err)
	}

	// todo: consider adding more validations

	return nil
}

func (p *Plugin) ObservationQuorum(outctx ocr3types.OutcomeContext, query ocrtypes.Query) (ocr3types.Quorum, error) {
	return ocr3types.QuorumTwoFPlusOne, nil
}

func (p *Plugin) Outcome(outctx ocr3types.OutcomeContext, query ocrtypes.Query, aos []ocrtypes.AttributedObservation) (ocr3types.Outcome, error) {
	lggr := p.lggr.With("seqNr", outctx.SeqNr, "numObservations", len(aos), "phase", "Outcome")
	lggr.Infow("in outcome")

	// Gather all the observations.
	observations := make([]models.Observation, 0, len(aos))
	for _, encodedObs := range aos {
		obs, err := models.DecodeObservation(encodedObs.Observation)
		if err != nil {
			return ocr3types.Outcome{}, fmt.Errorf("decode observation: %w", err)
		}
		lggr.Debugw("decoded observation", "observation", obs, "oracleID", encodedObs.Observer)
		observations = append(observations, obs)
	}

	// Come to a consensus based on the observations of all the different nodes.
	medianLiquidityPerChain := computeMedianLiquidityPerChain(observations)
	graphEdges := computeGraphEdgesConsensus(observations, p.f)

	pendingTransfers, err := computePendingTransfersConsensus(observations, p.f)
	if err != nil {
		return ocr3types.Outcome{}, fmt.Errorf("compute pending transfers consensus: %w", err)
	}

	configDigests, err := computeConfigDigestsConsensus(observations, p.f)
	if err != nil {
		return ocr3types.Outcome{}, fmt.Errorf("compute config digests consensus: %w", err)
	}

	// Compute a new graph with the median liquidities and the edges of the quorum of nodes.
	g, err := computeMedianGraph(lggr, graphEdges, medianLiquidityPerChain)
	if err != nil {
		return nil, fmt.Errorf("compute median graph: %w", err)
	}

	resolvedTransfersQuorum, err := computeResolvedTransfersQuorum(lggr, observations, p.f, p.bridgeFactory)
	if err != nil {
		return nil, fmt.Errorf("compute resolved transfers quorum: %w", err)
	}

	inflightQuorum, err := computeInflightTransfersQuorum(lggr, observations, p.f)
	if err != nil {
		return nil, fmt.Errorf("compute inflight transfers quorum: %w", err)
	}

	pendingTransfers = removeInflightTransfers(lggr, pendingTransfers, inflightQuorum)

	lggr.Infow("computing transfers to reach balance",
		"pendingTransfers", pendingTransfers,
		"liquidityGraph", g,
		"resolvedTransfersQuorum", resolvedTransfersQuorum,
	)
	inflightTransfers := append(pendingTransfers, toPending(resolvedTransfersQuorum)...)
	inflightTransfers = append(inflightTransfers, toPending(inflightQuorum)...)
	proposedTransfers, err := p.liquidityRebalancer.ComputeTransfersToBalance(g, inflightTransfers)
	if err != nil {
		return nil, fmt.Errorf("compute transfers to reach balance: %w", err)
	}

	outcome := models.NewOutcome(proposedTransfers, resolvedTransfersQuorum, pendingTransfers, inflightQuorum, configDigests).Encode()
	lggr.Infow("finished computing outcome",
		"medianLiquidityPerChain", medianLiquidityPerChain,
		"pendingTransfers", pendingTransfers,
		"proposedTransfers", proposedTransfers,
		"resolvedTransfers", resolvedTransfersQuorum,
		"inflightTransfers", inflightTransfers,
		"outcomeEncoded", outcome,
	)

	return outcome, nil
}

func (p *Plugin) Reports(seqNr uint64, outcome ocr3types.Outcome) ([]ocr3types.ReportWithInfo[models.Report], error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	lggr := p.lggr.With("seqNr", seqNr, "phase", "Reports")
	lggr.Infow("in reports")

	decodedOutcome, err := models.DecodeOutcome(outcome)
	if err != nil {
		return nil, fmt.Errorf("decode outcome: %w", err)
	}

	// get all incoming and outgoing transfers for each network
	// incoming transfers will need to be finalized
	// outgoing transfers will need to be executed
	incomingAndOutgoing := make(map[models.NetworkSelector][]models.Transfer)
	for _, networkID := range p.rebalancerGraph.GetNetworks() {
		for _, outgoing := range decodedOutcome.ResolvedTransfers {
			if outgoing.From == networkID {
				incomingAndOutgoing[networkID] = append(incomingAndOutgoing[networkID], outgoing)
			}
		}
		for _, incoming := range decodedOutcome.PendingTransfers {
			if incoming.To == networkID &&
				(incoming.Status == models.TransferStatusReady ||
					incoming.Status == models.TransferStatusFinalized) {
				incomingAndOutgoing[networkID] = append(incomingAndOutgoing[networkID], incoming.Transfer)
			}
		}
	}

	lggr = lggr.With(
		"incomingAndOutgoing", incomingAndOutgoing,
		"resolvedTransfers", decodedOutcome.ResolvedTransfers,
		"pendingTransfers", decodedOutcome.PendingTransfers,
		"proposedTransfers", decodedOutcome.ProposedTransfers,
	)
	lggr.Infow("got incoming and outgoing transfers")

	configDigestsMap := map[models.NetworkSelector]map[models.Address]types.ConfigDigest{}
	for _, cd := range decodedOutcome.ConfigDigests {
		_, found := configDigestsMap[cd.NetworkSel]
		if found {
			return nil, fmt.Errorf("found duplicate config digest for %v", cd.NetworkSel)
		}
		configDigestsMap[cd.NetworkSel] = map[models.Address]types.ConfigDigest{
			cd.RebalancerAddr: cd.Digest.ConfigDigest,
		}
	}

	var reports []ocr3types.ReportWithInfo[models.Report]
	for networkID, transfers := range incomingAndOutgoing {
		lmAddress, err := p.rebalancerGraph.GetRebalancerAddress(networkID)
		if err != nil {
			return nil, fmt.Errorf("liquidity manager for %v does not exist", networkID)
		}

		configDigests, found := configDigestsMap[networkID]
		if !found {
			return nil, fmt.Errorf("cannot find config digest for %v", networkID)
		}
		configDigest, found := configDigests[lmAddress]
		if !found {
			return nil, fmt.Errorf("cannot find config digest for %v:%s", networkID, lmAddress)
		}

		report := models.NewReport(transfers, lmAddress, networkID, configDigest)
		encoded, err := report.OnchainEncode()
		if err != nil {
			return nil, fmt.Errorf("encode report metadata for onchain usage: %w", err)
		}
		reports = append(reports, ocr3types.ReportWithInfo[models.Report]{
			Report: encoded,
			Info:   report,
		})
	}

	lggr.Infow("generated reports", "numReports", len(reports))
	return reports, nil
}

func (p *Plugin) ShouldAcceptAttestedReport(ctx context.Context, seqNr uint64, r ocr3types.ReportWithInfo[models.Report]) (bool, error) {
	lggr := p.lggr.With("seqNr", seqNr, "reportMeta", r.Info, "reportHex", hexutil.Encode(r.Report), "reportLen", len(r.Report), "phase", "ShouldAcceptAttestedReport")
	lggr.Infow("in should accept attested report")

	report, instructions, err := models.DecodeReport(r.Info.NetworkID, r.Info.RebalancerAddress, r.Report)
	if err != nil {
		return false, fmt.Errorf("failed to decode report: %w", err)
	}

	lggr = lggr.With(
		"networkSelector", report.NetworkID,
		"rebalancerAddress", report.RebalancerAddress,
		"transfers", len(report.Transfers),
	)
	lggr.Infow("decoded report")

	// report with no instructions should not be transmitted.
	if len(report.Transfers) == 0 {
		lggr.Infow("report has no transfers, returning false")
		return false, nil
	}

	if p.isStaleReport(ctx, lggr, seqNr, report.NetworkID, report.RebalancerAddress, report.Transfers) {
		lggr.Infow("report is stale, returning false")
		return false, nil
	}

	lggr.Infow("report is not stale, accepting",
		"transfers", len(report.Transfers),
		"sendInstructions", instructions.SendLiquidityParams,
		"receiveInstructions", instructions.ReceiveLiquidityParams)

	p.inflightCache.Add(lggr, report.Transfers)

	return true, nil
}

func (p *Plugin) ShouldTransmitAcceptedReport(ctx context.Context, seqNr uint64, r ocr3types.ReportWithInfo[models.Report]) (bool, error) {
	lggr := p.lggr.With("seqNr", seqNr, "reportMeta", r.Info, "reportHex", hexutil.Encode(r.Report), "reportLen", len(r.Report), "phase", "ShouldTransmitAcceptedReport")
	lggr.Infow("in should transmit accepted report")

	report, instructions, err := models.DecodeReport(r.Info.NetworkID, r.Info.RebalancerAddress, r.Report)
	if err != nil {
		return false, fmt.Errorf("failed to decode report: %w", err)
	}

	lggr = lggr.With(
		"networkSelector", report.NetworkID,
		"rebalancerAddress", report.RebalancerAddress,
		"transfers", len(report.Transfers))
	lggr.Infow("decoded report",
		"sendInstructions", instructions.SendLiquidityParams,
		"receiveInstructions", instructions.ReceiveLiquidityParams)

	// report with no instructions should not be transmitted.
	if len(report.Transfers) == 0 {
		lggr.Infow("report has no transfers, returning false")
		return false, nil
	}

	if p.isStaleReport(ctx, lggr, seqNr, report.NetworkID, report.RebalancerAddress, report.Transfers) {
		lggr.Infow("report is stale, returning false")
		return false, nil
	}

	lggr.Infow("report is not stale, transmitting")

	return true, nil
}

func (p *Plugin) isStaleReport(
	ctx context.Context,
	lggr logger.Logger,
	seqNr uint64,
	networkID models.NetworkSelector,
	rebalancerAddress models.Address,
	transfers []models.Transfer,
) bool {
	// check sequence number to see if its already transmitted onchain.
	rebalancer, err := p.liquidityManagerFactory.NewRebalancer(networkID, rebalancerAddress)
	if err != nil {
		lggr.Warnw("failed to get rebalancer", "err", err)
		return true
	}

	onchainSeqNr, err := rebalancer.GetLatestSequenceNumber(ctx)
	if err != nil {
		lggr.Warnw("failed to get latest sequence number", "err", err)
		return true
	}

	if onchainSeqNr >= seqNr {
		lggr.Infow("report already transmitted onchain, report is stale, should not be transmitted", "onchainSeqNr", onchainSeqNr)
		return true
	}

	lggr.Infow("onchain sequence number < current", "onchainSeqNr", onchainSeqNr)

	// check that the instructions will not cause failures onchain.
	// e.g send instructions when there is not enough liquidity.
	currentBalance, err := rebalancer.GetBalance(ctx)
	if err != nil {
		lggr.Warnw("failed to get balance", "err", err)
		return true
	}

	lggr.Infow("checking if there is enough balance onchain to send", "currentBalance", currentBalance.String())

	for _, transfer := range transfers {
		if transfer.From != networkID {
			continue
		}

		if currentBalance.Cmp(transfer.Amount.ToInt()) < 0 {
			lggr.Warnw("not enough balance onchain to send", "amount", transfer.Amount, "remaining", currentBalance.String())
			return true
		}
		currentBalance = currentBalance.Sub(currentBalance, transfer.Amount.ToInt())
	}

	if currentBalance.Cmp(big.NewInt(0)) < 0 {
		lggr.Warnw("not enough balance onchain to send", "remaining", currentBalance.String())
		return true
	}

	lggr.Infow("enough balance onchain to send", "currentBalance", currentBalance.String())

	return false
}

func (p *Plugin) Close() error {
	p.lggr.Infow("closing plugin")
	ctx, cf := context.WithTimeout(context.Background(), p.closePluginTimeout)
	defer cf()

	var errs []error
	for _, networkID := range p.rebalancerGraph.GetNetworks() {
		rebalancerAddress, err := p.rebalancerGraph.GetRebalancerAddress(networkID)
		if err != nil {
			errs = append(errs, fmt.Errorf("get rebalancer address for %v: %w", networkID, err))
			continue
		}

		rb, err := p.liquidityManagerFactory.GetRebalancer(networkID, rebalancerAddress)
		if err != nil {
			errs = append(errs, fmt.Errorf("get rebalancer (%d, %v): %w", networkID, rebalancerAddress, err))
			continue
		}

		if err := rb.Close(ctx); err != nil {
			errs = append(errs, fmt.Errorf("close rebalancer (%d, %v): %w", networkID, rebalancerAddress, err))
			continue
		}
	}

	return multierr.Combine(errs...)
}

func (p *Plugin) syncGraphEdges(ctx context.Context) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	// todo: if there wasn't any change to the graph stop earlier
	p.lggr.Infow("syncing graph edges")

	p.lggr.Infow("discovering rebalancers")
	discoverer, err := p.discovererFactory.NewDiscoverer(p.rootNetwork, p.rootAddress)
	if err != nil {
		return fmt.Errorf("init discoverer: %w", err)
	}

	g, err := discoverer.Discover(ctx)
	if err != nil {
		return fmt.Errorf("discovering rebalancers: %w", err)
	}

	p.rebalancerGraph = g

	p.lggr.Infow("finished syncing graph edges", "graph", g.String())

	return nil
}

func (p *Plugin) syncGraphBalances(ctx context.Context) ([]models.NetworkLiquidity, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	networks := p.rebalancerGraph.GetNetworks()
	p.lggr.Infow("syncing graph balances", "networks", networks)

	networkLiquidities := make([]models.NetworkLiquidity, 0, len(networks))
	for _, networkID := range networks {
		lmAddr, err := p.rebalancerGraph.GetRebalancerAddress(networkID)
		if err != nil {
			return nil, fmt.Errorf("liquidity manager for network %v was not found", networkID)
		}

		lm, err := p.liquidityManagerFactory.NewRebalancer(networkID, lmAddr)
		if err != nil {
			return nil, fmt.Errorf("init liquidity manager: %w", err)
		}

		balance, err := lm.GetBalance(ctx)
		if err != nil {
			return nil, fmt.Errorf("get %v balance: %w", networkID, err)
		}

		p.rebalancerGraph.SetLiquidity(networkID, balance)
		networkLiquidities = append(networkLiquidities, models.NewNetworkLiquidity(networkID, balance))
	}

	return networkLiquidities, nil
}

func (p *Plugin) loadPendingTransfers(ctx context.Context) ([]models.PendingTransfer, error) {
	p.lggr.Infow("loading pending transfers")

	pendingTransfers := make([]models.PendingTransfer, 0)
	for _, networkID := range p.rebalancerGraph.GetNetworks() {
		neighbors, ok := p.rebalancerGraph.GetNeighbors(networkID)
		if !ok {
			p.lggr.Warnw("no neighbors found for network", "network", networkID)
			continue
		}

		// todo: figure out what to do with this
		// dateToStartLookingFrom := time.Now().Add(-10 * 24 * time.Hour)

		// if mostRecentTransfer, exists := p.pendingTransfers.LatestNetworkTransfer(networkID); exists {
		// 	dateToStartLookingFrom = mostRecentTransfer.Date
		// }

		for _, neighbor := range neighbors {
			bridge, err := p.bridgeFactory.NewBridge(networkID, neighbor)
			if err != nil {
				return nil, fmt.Errorf("init bridge: %w", err)
			}

			if bridge == nil {
				p.lggr.Warnw("no bridge found for network pair", "sourceNetwork", networkID, "destNetwork", neighbor)
				continue
			}

			localToken, err := p.rebalancerGraph.GetTokenAddress(networkID)
			if err != nil {
				return nil, fmt.Errorf("get local token address for %v: %w", networkID, err)
			}
			remoteToken, err := p.rebalancerGraph.GetTokenAddress(neighbor)
			if err != nil {
				return nil, fmt.Errorf("get remote token address for %v: %w", neighbor, err)
			}

			netPendingTransfers, err := bridge.GetTransfers(ctx, localToken, remoteToken)
			if err != nil {
				return nil, fmt.Errorf("get pending transfers: %w", err)
			}

			p.lggr.Infow("loaded pending transfers", "network", networkID, "pendingTransfers", netPendingTransfers)
			pendingTransfers = append(pendingTransfers, netPendingTransfers...)
		}
	}

	return pendingTransfers, nil
}

func (p *Plugin) resolveProposedTransfers(ctx context.Context, outcomeCtx ocr3types.OutcomeContext) ([]models.Transfer, error) {
	p.lggr.Infow("resolving proposed transfers", "seqNr", outcomeCtx.SeqNr, "prevSeqNr", outcomeCtx.SeqNr-1)

	if len(outcomeCtx.PreviousOutcome) == 0 {
		return nil, nil
	}

	outcome, err := models.DecodeOutcome(outcomeCtx.PreviousOutcome)
	if err != nil {
		return nil, fmt.Errorf("decode previous outcome: %w", err)
	}

	var resolvedTransfers []models.Transfer
	for _, proposedTransfer := range outcome.ProposedTransfers {
		bridge, err := p.bridgeFactory.NewBridge(proposedTransfer.From, proposedTransfer.To)
		if err != nil {
			return nil, fmt.Errorf("init bridge: %w", err)
		}

		fromNetRebalancer, err := p.rebalancerGraph.GetRebalancerAddress(proposedTransfer.From)
		if err != nil {
			return nil, fmt.Errorf("get rebalancer address for %v: %w", proposedTransfer.From, err)
		}

		fromNetToken, err := p.rebalancerGraph.GetTokenAddress(proposedTransfer.From)
		if err != nil {
			return nil, fmt.Errorf("get token address for %v: %w", proposedTransfer.From, err)
		}

		toNetRebalancer, err := p.rebalancerGraph.GetRebalancerAddress(proposedTransfer.To)
		if err != nil {
			return nil, fmt.Errorf("get rebalancer address for %v: %w", proposedTransfer.To, err)
		}

		toNetToken, err := p.rebalancerGraph.GetTokenAddress(proposedTransfer.To)
		if err != nil {
			return nil, fmt.Errorf("get token address for %v: %w", proposedTransfer.To, err)
		}

		resolvedTransfer := models.Transfer{
			From:               proposedTransfer.From,
			To:                 proposedTransfer.To,
			Amount:             proposedTransfer.Amount,
			Sender:             fromNetRebalancer,
			Receiver:           toNetRebalancer,
			LocalTokenAddress:  fromNetToken,
			RemoteTokenAddress: toNetToken,
			// BridgeData: nil, // will be filled in below
			// NativeBridgeFee: big.NewInt(0), // will be filled in below
		}

		bridgePayload, bridgeFee, err := bridge.GetBridgePayloadAndFee(ctx, resolvedTransfer)
		if err != nil {
			p.lggr.Warnw("failed to get bridge payload and fee", "proposedTransfer", proposedTransfer, "err", err)
			// return nil, fmt.Errorf("get bridge payload and fee: %w", err)
			continue
		}
		resolvedTransfer.BridgeData = bridgePayload
		resolvedTransfer.NativeBridgeFee = ubig.New(bridgeFee)
		resolvedTransfers = append(resolvedTransfers, resolvedTransfer)
	}

	p.lggr.Infow("finished resolving proposed transfers", "resolvedTransfers", resolvedTransfers)

	return resolvedTransfers, nil
}

// bigIntSortedMiddle returns the middle number after sorting the provided numbers. nil is returned if the provided slice is empty.
// If length of the provided slice is even, the right-hand-side value of the middle 2 numbers is returned.
// The objective of this function is to always pick within the range of values reported by honest nodes when we have 2f+1 values.
// todo: move to libs
func bigIntSortedMiddle(vals []*big.Int) *big.Int {
	if len(vals) == 0 {
		return nil
	}

	valsCopy := make([]*big.Int, len(vals))
	copy(valsCopy[:], vals[:])
	sort.Slice(valsCopy, func(i, j int) bool {
		return valsCopy[i].Cmp(valsCopy[j]) == -1
	})
	return valsCopy[len(valsCopy)/2]
}

func toPending(ts []models.Transfer) []models.PendingTransfer {
	var pts []models.PendingTransfer
	for _, t := range ts {
		pts = append(pts, models.NewPendingTransfer(t))
	}
	return pts
}
