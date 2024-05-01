package liquidityrebalancer

import (
	"fmt"
	"math/big"
	"sort"

	mapset "github.com/deckarep/golang-set/v2"

	ubig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/liquiditymanager/graph"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/liquiditymanager/models"
)

// MinLiquidityRebalancer tries to reach balance using a target minimum liquidity that is configured on each network.
type MinLiquidityRebalancer struct {
	lggr logger.Logger
}

func NewMinLiquidityRebalancer(lggr logger.Logger) *MinLiquidityRebalancer {
	return &MinLiquidityRebalancer{
		lggr: lggr,
	}
}

func (r *MinLiquidityRebalancer) ComputeTransfersToBalance(
	graphNow graph.Graph,
	nonExecutedTransfers []UnexecutedTransfer,
) ([]models.ProposedTransfer, error) {
	nonExecutedTransfers = r.filterUnexecutedTransfers(nonExecutedTransfers)

	r.lggr.Debugf("computing the expected graph after non executed transfers get applied")
	graphLater, err := r.getExpectedGraph(graphNow, nonExecutedTransfers)
	if err != nil {
		return nil, fmt.Errorf("copy graph: %w", err)
	}

	r.lggr.Debugf("finding networks that require funding")
	networksRequiringFunding, liqDiffsNow, liqDiffsLater, err := r.findNetworksRequiringFunding(graphNow, graphLater)
	if err != nil {
		return nil, fmt.Errorf("find networks that require funding: %w", err)
	}

	r.lggr.Debugf("computing transfers to reach balance using a direct transfer from one network to another")
	proposedTransfers := make([]models.ProposedTransfer, 0)
	for _, net := range networksRequiringFunding {
		potentialTransfers, err2 := r.oneHopTransfers(graphLater, net, liqDiffsNow, liqDiffsLater)
		if err2 != nil {
			return nil, fmt.Errorf("find 1 hop transfers for network %d: %w", net, err2)
		}
		netProposedTransfers, err2 := r.acceptTransfers(graphLater, potentialTransfers, liqDiffsLater[net])
		if err2 != nil {
			return nil, fmt.Errorf("accepting transfers: %w", err2)
		}
		proposedTransfers = append(proposedTransfers, netProposedTransfers...)
	}

	r.lggr.Debugf("finding networks that still require funding")
	networksRequiringFunding, liqDiffsNow, liqDiffsLater, err = r.findNetworksRequiringFunding(graphNow, graphLater)
	if err != nil {
		return nil, fmt.Errorf("find networks that require funding: %w", err)
	}

	r.lggr.Debugf("computing transfers to reach balance with an initial transfer to an intermediate network")
	for _, net := range networksRequiringFunding {
		transfers, err2 := r.twoHopTransfers(graphLater, net, liqDiffsNow, liqDiffsLater)
		if err2 != nil {
			return nil, fmt.Errorf("find 2 hops transfers for network %d: %w", net, err2)
		}
		netProposedTransfers, err2 := r.acceptTransfers(graphLater, transfers, liqDiffsLater[net])
		if err2 != nil {
			return nil, fmt.Errorf("accepting 2hop transfers: %w", err2)
		}
		proposedTransfers = append(proposedTransfers, netProposedTransfers...)
	}

	proposedTransfers = r.mergeProposedTransfers(proposedTransfers)

	r.lggr.Debugf("sorting proposed transfers for determinism")
	sort.Slice(proposedTransfers, func(i, j int) bool {
		if proposedTransfers[i].From == proposedTransfers[j].From {
			return proposedTransfers[i].To < proposedTransfers[j].To
		}
		return proposedTransfers[i].From < proposedTransfers[j].From
	})

	return proposedTransfers, nil
}

func (r *MinLiquidityRebalancer) findNetworksRequiringFunding(graphNow, graphLater graph.Graph) (
	nets []models.NetworkSelector,
	liqDiffsNow, liqDiffsLater map[models.NetworkSelector]*big.Int,
	err error,
) {
	liqDiffsNow, liqDiffsLater, err = r.getTargetLiquidityDifferences(graphNow, graphLater)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("compute tokens funding requirements: %w", err)
	}

	res := make([]models.NetworkSelector, 0, len(liqDiffsNow))
	for net := range liqDiffsNow {
		diffNow := liqDiffsNow[net]
		diffLater := liqDiffsLater[net]

		if diffNow.Cmp(big.NewInt(0)) <= 0 || diffLater.Cmp(big.NewInt(0)) <= 0 {
			r.lggr.Debugf("net %d does not require funding, transferrable tokens: %d (*%d)", net, big.NewInt(0).Abs(diffNow), big.NewInt(0).Abs(diffLater))
			continue
		}

		r.lggr.Debugf("net %d requires funding, %s tokens to reach target", net, diffLater)
		res = append(res, net)
	}

	sort.Slice(res, func(i, j int) bool { return liqDiffsLater[res[i]].Cmp(liqDiffsLater[res[j]]) > 0 })
	return res, liqDiffsNow, liqDiffsLater, nil
}

func (r *MinLiquidityRebalancer) filterUnexecutedTransfers(nonExecutedTransfers []UnexecutedTransfer) []UnexecutedTransfer {
	r.lggr.Debugf("filtering out executed transfers")
	filtered := make([]UnexecutedTransfer, 0, len(nonExecutedTransfers))
	for _, tr := range nonExecutedTransfers {
		if tr.TransferStatus() != models.TransferStatusExecuted {
			filtered = append(filtered, tr)
		}
	}
	return filtered
}

// getTargetLiquidityDifferences computes for each network the difference between
// the target liquidity of the network and the actual liquidity.
// It does this on both the current liquidity graph (graphNow) and the liquidity graph
// after all pending transfers have been successfully executed (graphLater).
// A negative number indicates that there is a liquidity shortage for the network,
// while a positive number indicates a liquidity surplus for the network.
func (r *MinLiquidityRebalancer) getTargetLiquidityDifferences(
	graphNow, graphLater graph.Graph,
) (liqDiffsNow, liqDiffsLater map[models.NetworkSelector]*big.Int, err error) {
	liqDiffsNow = make(map[models.NetworkSelector]*big.Int)
	liqDiffsLater = make(map[models.NetworkSelector]*big.Int)

	for _, net := range graphNow.GetNetworks() {
		dataNow, err := graphNow.GetData(net)
		if err != nil {
			return nil, nil, fmt.Errorf("get data now of net %d: %w", net, err)
		}

		dataLater, err := graphLater.GetData(net)
		if err != nil {
			return nil, nil, fmt.Errorf("get data later of net %d: %w", net, err)
		}

		if dataNow.MinimumLiquidity.Cmp(big.NewInt(0)) == 0 {
			// automated rebalancing is disabled if target is set to 0
			liqDiffsNow[net] = big.NewInt(0)
			liqDiffsLater[net] = big.NewInt(0)
			continue
		}

		liqDiffsNow[net] = big.NewInt(0).Sub(dataNow.MinimumLiquidity, dataNow.Liquidity)
		liqDiffsLater[net] = big.NewInt(0).Sub(dataLater.MinimumLiquidity, dataLater.Liquidity)
	}

	return liqDiffsNow, liqDiffsLater, nil
}

func (r *MinLiquidityRebalancer) oneHopTransfers(
	graphLater graph.Graph, // the networks graph state after all transfers are applied
	targetNetwork models.NetworkSelector,
	liqDiffsNow map[models.NetworkSelector]*big.Int, // the token funding requirements for each network
	liqDiffsLater map[models.NetworkSelector]*big.Int, // the token funding requirements after all pending txs are applied
) ([]models.ProposedTransfer, error) {
	allEdges, err := graphLater.GetEdges()
	if err != nil {
		return nil, fmt.Errorf("get edges: %w", err)
	}

	potentialTransfers := make([]models.ProposedTransfer, 0)
	seenNetworks := mapset.NewSet[models.NetworkSelector]()

	for _, edge := range allEdges {
		if edge.Dest != targetNetwork {
			// we only care about the target network
			continue
		}

		if seenNetworks.Contains(edge.Source) {
			// cannot have the same sender twice
			continue
		}

		diffNow, exists := liqDiffsNow[edge.Source]
		if !exists {
			return nil, fmt.Errorf("net %d does not exist in the tokens to target offset", edge.Source)
		}
		diff := diffNow

		diffLater, exists := liqDiffsLater[edge.Source]
		if !exists {
			return nil, fmt.Errorf("net %d does not exist in the tokens to target offset", edge.Source)
		}

		// If the balance is expected to become lower, we consider the lower balance to prevent a race condition in the transfers.
		if diffNow.Cmp(diffLater) < 0 {
			diff = diffLater
		}

		transferAmount := big.NewInt(0).Sub(big.NewInt(0), diff)
		if transferAmount.Cmp(big.NewInt(0)) <= 0 {
			continue
		}

		potentialTransfers = append(potentialTransfers, newTransfer(edge.Source, targetNetwork, transferAmount))
		seenNetworks.Add(edge.Source)
	}

	return potentialTransfers, nil
}

// twoHopTransfers finds networks that can increase liquidity of the target network with an intermediate network.
func (r *MinLiquidityRebalancer) twoHopTransfers(
	graphLater graph.Graph, // the networks graph state after all transfers are applied
	targetNetwork models.NetworkSelector,
	reqFundingNow map[models.NetworkSelector]*big.Int, // the token funding requirements for each network
	reqFundingLater map[models.NetworkSelector]*big.Int, // the token funding requirements after all pending txs are applied
) ([]models.ProposedTransfer, error) {
	potentialTransfers := make([]models.ProposedTransfer, 0)
	seenNetworks := mapset.NewSet[models.NetworkSelector]()

	for _, net := range graphLater.GetNetworks() {
		if net == targetNetwork {
			continue
		}
		if seenNetworks.Contains(net) {
			// cannot have the same sender twice
			continue
		}

		neibs, ok := graphLater.GetNeighbors(net)
		if !ok {
			return nil, fmt.Errorf("get neighbors of %d failed", net)
		}
		neibsSet := mapset.NewSet[models.NetworkSelector](neibs...)
		if neibsSet.Contains(targetNetwork) {
			// since the target network is a direct network we can transfer using 1hop
			continue
		}

		for _, neib := range neibs {
			intermNeibs, ok := graphLater.GetNeighbors(neib)
			if !ok {
				return nil, fmt.Errorf("get intermediate neighbors of %d failed", net)
			}
			finalNeibsSet := mapset.NewSet[models.NetworkSelector](intermNeibs...)
			if finalNeibsSet.Contains(targetNetwork) {
				fundingNow, exists := reqFundingNow[net]
				if !exists {
					return nil, fmt.Errorf("net %d does not exist in the tokens to target offset", net)
				}
				funding := fundingNow

				fundingLater, exists := reqFundingLater[net]
				if !exists {
					return nil, fmt.Errorf("net %d does not exist in the tokens to target offset", net)
				}

				// If the balance is expected to decrease, consider the lower balance to prevent a race condition in the transfers.
				if fundingNow.Cmp(fundingLater) < 0 {
					funding = fundingLater
				}

				transferAmount := big.NewInt(0).Sub(big.NewInt(0), funding)
				if transferAmount.Cmp(big.NewInt(0)) <= 0 {
					continue
				}

				seenNetworks.Add(net)
				potentialTransfers = append(potentialTransfers, newTransfer(net, neib, transferAmount))
			}
		}
	}

	return potentialTransfers, nil
}

// apply changes to the intermediate state to prevent invalid transfers
func (r *MinLiquidityRebalancer) acceptTransfers(graphLater graph.Graph, potentialTransfers []models.ProposedTransfer, requiredAmount *big.Int) ([]models.ProposedTransfer, error) {
	// sort by amount,sender,receiver
	sort.Slice(potentialTransfers, func(i, j int) bool {
		if potentialTransfers[i].Amount.Cmp(potentialTransfers[j].Amount) == 0 {
			if potentialTransfers[i].From == potentialTransfers[j].From {
				return potentialTransfers[i].To < potentialTransfers[j].To
			}
			return potentialTransfers[i].From < potentialTransfers[j].From
		}
		return potentialTransfers[i].Amount.Cmp(potentialTransfers[j].Amount) > 0
	})

	fundsRaised := big.NewInt(0)
	proposedTransfers := make([]models.ProposedTransfer, 0, len(potentialTransfers))
	skip := false
	for _, d := range potentialTransfers {
		if skip {
			r.lggr.Debugf("skipping transfer: %s", d)
			continue
		}

		senderData, err := graphLater.GetData(d.From)
		if err != nil {
			return nil, fmt.Errorf("get liquidity of sender %d: %w", d.From, err)
		}
		availableAmount := big.NewInt(0).Sub(senderData.Liquidity, senderData.MinimumLiquidity)
		if availableAmount.Cmp(big.NewInt(0)) <= 0 {
			r.lggr.Debugf("no more tokens to transfer, skipping transfer: %s", d)
			continue
		}

		if availableAmount.Cmp(d.Amount.ToInt()) < 0 {
			d.Amount = ubig.New(availableAmount)
			r.lggr.Debugf("reducing transfer amount since sender balance has dropped: %s", d)
		}

		// increment the raised funds
		fundsRaised = big.NewInt(0).Add(fundsRaised, d.Amount.ToInt())

		// in case we raised more than target amount give refund to the sender
		if refund := big.NewInt(0).Sub(fundsRaised, requiredAmount); refund.Cmp(big.NewInt(0)) > 0 {
			d.Amount = ubig.New(big.NewInt(0).Sub(d.Amount.ToInt(), refund))
			fundsRaised = big.NewInt(0).Sub(fundsRaised, refund)
		}
		r.lggr.Debugf("accepting transfer: %s", d)
		proposedTransfers = append(proposedTransfers, models.ProposedTransfer{From: d.From, To: d.To, Amount: d.Amount})

		r.lggr.Debugf("applying transfer to future graph state")
		liqBefore, err := graphLater.GetLiquidity(d.To)
		if err != nil {
			return nil, fmt.Errorf("get liquidity of transfer receiver %d: %w", d.To, err)
		}
		graphLater.SetLiquidity(d.To, big.NewInt(0).Add(liqBefore, d.Amount.ToInt()))

		liqBefore, err = graphLater.GetLiquidity(d.From)
		if err != nil {
			return nil, fmt.Errorf("get liquidity of sender %d: %w", d.From, err)
		}
		graphLater.SetLiquidity(d.From, big.NewInt(0).Sub(liqBefore, d.Amount.ToInt()))

		if fundsRaised.Cmp(requiredAmount) >= 0 {
			r.lggr.Debugf("all funds raised skipping further transfers")
			skip = true
		}
	}

	return proposedTransfers, nil
}

// getExpectedGraph returns the a copy of the graph instance with all the non executed transfers applied.
func (r *MinLiquidityRebalancer) getExpectedGraph(
	g graph.Graph,
	nonExecutedTransfers []UnexecutedTransfer,
) (graph.Graph, error) {
	edges, err := g.GetEdges()
	if err != nil {
		return nil, err
	}

	expG := graph.NewGraph()
	for _, edge := range edges {
		sourceData, err := g.GetData(edge.Source)
		if err != nil {
			return nil, err
		}

		destData, err := g.GetData(edge.Dest)
		if err != nil {
			return nil, err
		}

		expG.AddNetwork(edge.Source, sourceData)
		expG.AddNetwork(edge.Dest, destData)
		if err := expG.AddConnection(edge.Source, edge.Dest); err != nil {
			return nil, err
		}
	}

	for _, tr := range nonExecutedTransfers {
		liqTo, err := expG.GetLiquidity(tr.ToNetwork())
		if err != nil {
			return nil, err
		}
		expG.SetLiquidity(tr.ToNetwork(), big.NewInt(0).Add(liqTo, tr.TransferAmount()))

		switch tr.TransferStatus() {
		case models.TransferStatusProposed, models.TransferStatusInflight:
			liqFrom, err := expG.GetLiquidity(tr.FromNetwork())
			if err != nil {
				return nil, err
			}
			expG.SetLiquidity(tr.FromNetwork(), big.NewInt(0).Sub(liqFrom, tr.TransferAmount()))
		}
	}

	return expG, nil
}

// mergeProposedTransfers merges multiple transfers with same sender and recipient into a single transfer.
func (r *MinLiquidityRebalancer) mergeProposedTransfers(transfers []models.ProposedTransfer) []models.ProposedTransfer {
	sums := make(map[[2]models.NetworkSelector]*big.Int)
	for _, tr := range transfers {
		k := [2]models.NetworkSelector{tr.From, tr.To}
		if _, exists := sums[k]; !exists {
			sums[k] = tr.TransferAmount()
			continue
		}
		sums[k] = big.NewInt(0).Add(sums[k], tr.TransferAmount())
	}

	merged := make([]models.ProposedTransfer, 0, len(transfers))
	for k, v := range sums {
		merged = append(merged, models.ProposedTransfer{From: k[0], To: k[1], Amount: ubig.New(v)})
	}
	return merged
}

func newTransfer(from, to models.NetworkSelector, amount *big.Int) models.ProposedTransfer {
	return models.ProposedTransfer{
		From:   from,
		To:     to,
		Amount: ubig.New(amount),
	}
}
