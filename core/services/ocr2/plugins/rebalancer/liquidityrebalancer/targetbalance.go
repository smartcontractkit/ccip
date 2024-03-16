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
type TargetBalanceRebalancer struct {
	lggr logger.Logger
}

func NewTargetBalanceRebalancer(lggr logger.Logger) *TargetBalanceRebalancer {
	return &TargetBalanceRebalancer{
		lggr: lggr,
	}
}

func (r *TargetBalanceRebalancer) ComputeTransfersToBalance(
	graphNow graph.Graph,
	nonExecutedTransfers []UnexecutedTransfer,
) ([]models.ProposedTransfer, error) {
	// filter out executed transfers - graph state should already reflect this transfers
	filtered := make([]UnexecutedTransfer, 0, len(nonExecutedTransfers))
	for _, tr := range nonExecutedTransfers {
		if tr.TransferStatus() != models.TransferStatusExecuted {
			filtered = append(filtered, tr)
		}
	}
	nonExecutedTransfers = filtered

	graphLater, err := r.getExpectedGraph(graphNow, nonExecutedTransfers)
	if err != nil {
		return nil, fmt.Errorf("copy graph: %w", err)
	}

	reqFundingNow, reqFundingLater, err := r.getRequiredTokensFunding(graphNow, graphLater)
	if err != nil {
		return nil, fmt.Errorf("compute tokens funding requirements: %w", err)
	}

	proposedTransfers := make([]models.ProposedTransfer, 0)
	for net := range reqFundingNow {
		fundingNow := reqFundingNow[net]
		fundingLater := reqFundingLater[net]

		if fundingNow.Cmp(big.NewInt(0)) <= 0 {
			r.lggr.Debugf("net %d does not require funding, donatable tokens: %d", net, big.NewInt(0).Abs(fundingNow))
			continue // no tokens required, already in target
		}
		if fundingLater.Cmp(big.NewInt(0)) <= 0 {
			r.lggr.Debugf("net %d does not require funding, donatable tokens will soon be: %d", net, big.NewInt(0).Abs(fundingLater))
			continue
		}

		r.lggr.Debugf("net %d requires %d token donations to reach target", net, fundingLater)
		donors, err := r.find1HopDonors(graphLater, net, fundingLater, reqFundingNow, reqFundingLater)
		if err != nil {
			return nil, fmt.Errorf("find 1 hop donors for network %d for %s tokens: %w", net, fundingLater, err)
		}

		for _, d := range donors {
			r.lggr.Debugf("network %d donates to %d: %s tokens", d.net, net, d.amount)
			proposedTransfers = append(proposedTransfers, models.ProposedTransfer{From: d.net, To: net, Amount: ubig.New(d.amount)})

			// apply changes to the intermediate state to prevent invalid donations
			liqBefore, err := graphLater.GetLiquidity(net)
			if err != nil {
				return nil, err
			}
			graphLater.SetLiquidity(net, big.NewInt(0).Add(liqBefore, d.amount))

			liqBefore, err = graphLater.GetLiquidity(d.net)
			if err != nil {
				return nil, err
			}
			graphLater.SetLiquidity(d.net, big.NewInt(0).Sub(liqBefore, d.amount))
		}

		// todo: after 1hop there might still be networks below target, try to reach targets using 2hops.
		//
	}

	sort.Slice(proposedTransfers, func(i, j int) bool {
		if proposedTransfers[i].From == proposedTransfers[j].From {
			return proposedTransfers[i].To < proposedTransfers[j].To
		}
		return proposedTransfers[i].From < proposedTransfers[j].From
	})

	return proposedTransfers, nil
}

// getRequiredTokensFunding computes for each network the required funding.
// Negative funding means that this network can keep donating tokens until reaching zero.
func (r *TargetBalanceRebalancer) getRequiredTokensFunding(
	graphNow, graphLater graph.Graph,
) (reqFundingNow, reqFundingLater map[models.NetworkSelector]*big.Int, err error) {
	reqFundingNow = make(map[models.NetworkSelector]*big.Int)
	reqFundingLater = make(map[models.NetworkSelector]*big.Int)

	for _, net := range graphNow.GetNetworks() {
		dataNow, err := graphNow.GetData(net)
		if err != nil {
			return nil, nil, fmt.Errorf("get data now of net %d: %w", net, err)
		}

		dataLater, err := graphLater.GetData(net)
		if err != nil {
			return nil, nil, fmt.Errorf("get data later of net %d: %w", net, err)
		}

		reqFundingNow[net] = big.NewInt(0).Sub(dataNow.TargetLiquidity, dataNow.Liquidity)
		reqFundingLater[net] = big.NewInt(0).Sub(dataLater.TargetLiquidity, dataLater.Liquidity)
		r.lggr.Debugf("current required funding for %d is: %s (expected to become: %s)",
			net, reqFundingNow[net], reqFundingLater[net])
	}

	return reqFundingNow, reqFundingLater, nil
}

// find1HopDonors finds networks that can increase liquidity of the target network with a single bridge transaction.
func (r *TargetBalanceRebalancer) find1HopDonors(
	graphLater graph.Graph, // the networks graph state after all transfers are applied
	donateTo models.NetworkSelector, // target network
	requiredAmount *big.Int, // the required tokens amount
	reqFundingNow map[models.NetworkSelector]*big.Int, // the token funding requirements for each network
	reqFundingLater map[models.NetworkSelector]*big.Int, // the token funding requirements after all pending txs are applied
) ([]netAmount, error) {
	allEdges, err := graphLater.GetEdges()
	if err != nil {
		return nil, fmt.Errorf("get edges: %w", err)
	}

	donors := make([]netAmount, 0)
	seenDonors := mapset.NewSet[models.NetworkSelector]()
	fundsRaised := big.NewInt(0)

	potentialDonors := make([]netAmount, 0)
	for _, edge := range allEdges {
		if edge.Dest != donateTo {
			// we only care about the target network
			continue
		}

		if seenDonors.Contains(edge.Source) {
			// cannot have the same donor twice
			continue
		}

		fundingNow, exists := reqFundingNow[edge.Source]
		if !exists {
			return nil, fmt.Errorf("net %d does not exist in the tokens to target offset", edge.Source)
		}
		funding := fundingNow

		fundingLater, exists := reqFundingLater[edge.Source]
		if !exists {
			return nil, fmt.Errorf("net %d does not exist in the tokens to target offset", edge.Source)
		}

		// If the balance is expected to become lower, we consider the lower balance to prevent a race condition in the donations.
		// If the balance is expected to become higher, we do not consider it since the funds are not available yet.
		if fundingNow.Cmp(fundingLater) < 0 {
			funding = fundingLater
		}

		potentialDonors = append(potentialDonors, newNetAmount(edge.Source, funding))
		seenDonors.Add(edge.Source)
	}

	// order potential donors by offset to target
	sort.Slice(potentialDonors, func(i, j int) bool {
		if potentialDonors[i].amount.Cmp(potentialDonors[j].amount) == 0 {
			return potentialDonors[i].net < potentialDonors[j].net
		}
		return potentialDonors[i].amount.Cmp(potentialDonors[j].amount) < 0
	})

	for _, potentialDonor := range potentialDonors {
		tokensToBalance := potentialDonor.amount
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

		donors = append(donors, newNetAmount(potentialDonor.net, donatedAmount))
		// if all the funds are raised stop
		if fundsRaised.Cmp(requiredAmount) >= 0 {
			break
		}
	}

	return donors, nil
}

func (r *TargetBalanceRebalancer) getExpectedGraph(
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

// helper struct that holds a network, amount pair
type netAmount struct {
	net    models.NetworkSelector
	amount *big.Int
}

func newNetAmount(net models.NetworkSelector, am *big.Int) netAmount {
	return netAmount{net: net, amount: am}
}
