package liquidityrebalancer

import (
	"fmt"
	"math/big"
	"sort"

	mapset "github.com/deckarep/golang-set/v2"

	ubig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/graph"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

// TargetBalanceRebalancer tries to reach balance using a target balance that is configured on each network.
// If the network balance falls below the target, balance is incremented using the other networks until reaching target.
type TargetBalanceRebalancer struct {
	lggr logger.Logger
}

func NewTargetBalanceRebalancer(lggr logger.Logger) *TargetBalanceRebalancer {
	return &TargetBalanceRebalancer{
		lggr: lggr,
	}
}

func (r *TargetBalanceRebalancer) ComputeTransfersToBalance(
	originalGraph graph.Graph,
	nonExecutedTransfers []UnexecutedTransfer,
) ([]models.ProposedTransfer, error) {
	// copy the graph before making changes
	g, err := r.copyGraph(originalGraph)
	if err != nil {
		return nil, fmt.Errorf("copy graph: %w", err)
	}

	for _, tr := range nonExecutedTransfers {
		liqTo, err := g.GetLiquidity(tr.ToNetwork())
		if err != nil {
			return nil, err
		}
		g.SetLiquidity(tr.ToNetwork(), big.NewInt(0).Add(liqTo, tr.TransferAmount()))

		switch tr.TransferStatus() {
		case models.TransferStatusProposed, models.TransferStatusInflight:
			liqFrom, err := g.GetLiquidity(tr.FromNetwork())
			if err != nil {
				return nil, err
			}
			g.SetLiquidity(tr.FromNetwork(), big.NewInt(0).Sub(liqFrom, tr.TransferAmount()))
		}
	}

	proposedTransfers := make([]models.ProposedTransfer, 0)

	// for each network of the graph compute target offset
	requiredTokensToReachTarget := make(map[models.NetworkSelector]*big.Int)
	for _, net := range g.GetNetworks() {
		data, err := g.GetData(net)
		if err != nil {
			return nil, err
		}
		requiredTokensToReachTarget[net] = big.NewInt(0).Sub(data.TargetLiquidity, data.Liquidity)
		r.lggr.Debugf("required tokens to reach target for %d is: %s", net, requiredTokensToReachTarget[net])
	}

	// find networks with balance below target
	for net, requiredTokens := range requiredTokensToReachTarget {
		if requiredTokens.Cmp(big.NewInt(0)) <= 0 {
			continue
		}
		donors, err := r.find1HopDonors(g, net, requiredTokens, requiredTokensToReachTarget)
		if err != nil {
			return nil, fmt.Errorf("find network=%d 1 hop donors for %s tokens: %w", net, requiredTokens, err)
		}

		for _, d := range donors {
			r.lggr.Debugf("%d donates to %d: %s", d.net, net, d.liq)
			proposedTransfers = append(proposedTransfers, models.ProposedTransfer{
				From: d.net, To: net, Amount: ubig.New(d.liq)})

			// increase liquidity of network
			liqBefore, err := g.GetLiquidity(net)
			if err != nil {
				return nil, err
			}
			g.SetLiquidity(net, big.NewInt(0).Add(liqBefore, d.liq))

			// decrease liquidity of donor
			liqBefore, err = g.GetLiquidity(d.net)
			if err != nil {
				return nil, err
			}
			g.SetLiquidity(d.net, big.NewInt(0).Sub(liqBefore, d.liq))
		}

		// todo: after 1hop there might still be networks below target
		//       try to reach targets using 2hops.
	}

	sort.Slice(proposedTransfers, func(i, j int) bool {
		return proposedTransfers[i].From < proposedTransfers[j].From ||
			(proposedTransfers[i].From == proposedTransfers[j].From && proposedTransfers[i].To < proposedTransfers[j].To)
	})

	return proposedTransfers, nil
}

type donor struct {
	net models.NetworkSelector
	liq *big.Int
}

// find1HopDonors finds networks that can increase liquidity of
func (r *TargetBalanceRebalancer) find1HopDonors(
	g graph.Graph,
	donateTo models.NetworkSelector,
	requiredAmount *big.Int,
	tokensToReachTarget map[models.NetworkSelector]*big.Int,
) ([]donor, error) {
	allEdges, err := g.GetEdges()
	if err != nil {
		return nil, fmt.Errorf("get edges: %w", err)
	}

	donors := make([]donor, 0)
	donorsSet := mapset.NewSet[models.NetworkSelector]()
	fundsRaised := big.NewInt(0)

	// find potential donors
	potentialDonors := make([]donor, 0)
	for _, edge := range allEdges {
		if edge.Dest != donateTo {
			// we only care about the target network
			continue
		}

		if donorsSet.Contains(edge.Source) {
			// cannot have the same donor twice
			continue
		}

		tokensOffset, exists := tokensToReachTarget[edge.Source]
		if !exists {
			return nil, fmt.Errorf("net %d does not exist in the tokens to target offset", edge.Source)
		}

		potentialDonors = append(potentialDonors, donor{net: edge.Source, liq: tokensOffset})
		donorsSet.Add(edge.Source)
	}

	// order potential donors by offset to target
	sort.Slice(potentialDonors, func(i, j int) bool {
		if potentialDonors[i].liq.Cmp(potentialDonors[j].liq) == 0 {
			return potentialDonors[i].net < potentialDonors[j].net
		}
		return potentialDonors[i].liq.Cmp(potentialDonors[j].liq) < 0
	})

	for _, potentialDonor := range potentialDonors {
		tokensToBalance := potentialDonor.liq
		hasEnoughTokens := tokensToBalance.Cmp(big.NewInt(0)) < 0
		if !hasEnoughTokens {
			continue
		}

		// donate everything
		donatedAmount := big.NewInt(0).Mul(tokensToBalance, big.NewInt(-1))
		// increment the raised funds
		fundsRaised = big.NewInt(0).Add(fundsRaised, donatedAmount)
		// in case we surpassed target give refund to donator
		if refund := big.NewInt(0).Sub(fundsRaised, requiredAmount); refund.Cmp(big.NewInt(0)) > 0 {
			donatedAmount = big.NewInt(0).Sub(donatedAmount, refund)
			fundsRaised = big.NewInt(0).Sub(fundsRaised, refund)
		}
		// add donor to the donors set and include it in the results
		donorsSet.Add(potentialDonor.net)
		donors = append(donors, donor{net: potentialDonor.net, liq: donatedAmount})
		// if all the funds are raised stop
		if fundsRaised.Cmp(requiredAmount) >= 0 {
			break
		}
	}

	return donors, nil
}

func (r *TargetBalanceRebalancer) copyGraph(originalG graph.Graph) (graph.Graph, error) {
	edges, err := originalG.GetEdges()
	if err != nil {
		return nil, err
	}

	g := graph.NewGraph()

	for _, edge := range edges {
		sourceData, err := originalG.GetData(edge.Source)
		if err != nil {
			return nil, err
		}

		destData, err := originalG.GetData(edge.Dest)
		if err != nil {
			return nil, err
		}

		g.AddNetwork(edge.Source, sourceData)
		g.AddNetwork(edge.Dest, destData)
		if err := g.AddConnection(edge.Source, edge.Dest); err != nil {
			return nil, err
		}
	}
	return g, nil
}
