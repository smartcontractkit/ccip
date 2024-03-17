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
	r.lggr.Debugf("filtering out executed transfers")
	filtered := make([]UnexecutedTransfer, 0, len(nonExecutedTransfers))
	for _, tr := range nonExecutedTransfers {
		if tr.TransferStatus() != models.TransferStatusExecuted {
			filtered = append(filtered, tr)
		}
	}
	nonExecutedTransfers = filtered

	r.lggr.Debugf("computing the expected graph after non executed transfers get applied")
	graphLater, err := r.getExpectedGraph(graphNow, nonExecutedTransfers)
	if err != nil {
		return nil, fmt.Errorf("copy graph: %w", err)
	}

	r.lggr.Debugf("finding networks that require funding")
	networksRequiringFunding, reqFundingNow, reqFundingLater, err := r.findNetworksRequiringFunding(graphNow, graphLater)
	if err != nil {
		return nil, fmt.Errorf("find networks that require funding: %w", err)
	}

	r.lggr.Debugf("computing transfers to reach balance using a direct transfer from one network to another")
	proposedTransfers := make([]models.ProposedTransfer, 0)
	for _, net := range networksRequiringFunding {
		potentialDonations, err2 := r.find1hopPotentialDonations(graphLater, net, reqFundingNow, reqFundingLater)
		if err2 != nil {
			return nil, fmt.Errorf("find 1 hop donations for network %d: %w", net, err2)
		}
		netProposedTransfers, err2 := r.acceptDonations(graphLater, potentialDonations, reqFundingLater[net])
		if err2 != nil {
			return nil, fmt.Errorf("accepting donations: %w", err2)
		}
		proposedTransfers = append(proposedTransfers, netProposedTransfers...)
	}

	r.lggr.Debugf("finding networks that still require funding")
	networksRequiringFunding, reqFundingNow, reqFundingLater, err = r.findNetworksRequiringFunding(graphNow, graphLater)
	if err != nil {
		return nil, fmt.Errorf("find networks that require funding: %w", err)
	}

	r.lggr.Debugf("computing transfers to reach balance with an initial transfer to an intermediate network")
	for _, net := range networksRequiringFunding {
		donations, err2 := r.find2hopPotentialDonations(graphLater, net, reqFundingNow, reqFundingLater)
		if err2 != nil {
			return nil, fmt.Errorf("find 2 hops donations for network %d: %w", net, err2)
		}
		netProposedTransfers, err2 := r.acceptDonations(graphLater, donations, reqFundingLater[net])
		if err2 != nil {
			return nil, fmt.Errorf("accepting 2hop donations: %w", err2)
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

func (r *TargetBalanceRebalancer) findNetworksRequiringFunding(graphNow, graphLater graph.Graph) (
	nets []models.NetworkSelector,
	reqFundingNow, reqFundingLater map[models.NetworkSelector]*big.Int,
	err error,
) {
	reqFundingNow, reqFundingLater, err = r.getRequiredTokensFunding(graphNow, graphLater)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("compute tokens funding requirements: %w", err)
	}

	res := make([]models.NetworkSelector, 0, len(reqFundingNow))
	for net := range reqFundingNow {
		fundingNow := reqFundingNow[net]
		fundingLater := reqFundingLater[net]

		if fundingNow.Cmp(big.NewInt(0)) <= 0 || fundingLater.Cmp(big.NewInt(0)) <= 0 {
			r.lggr.Debugf("net %d does not require funding, donatable tokens: %d (*%d)", net, big.NewInt(0).Abs(fundingNow), big.NewInt(0).Abs(fundingLater))
			continue
		}

		r.lggr.Debugf("net %d requires funding, %d tokens to reach target", net, fundingLater)
		res = append(res, net)
	}

	return res, reqFundingNow, reqFundingLater, nil
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

		if dataNow.TargetLiquidity.Cmp(big.NewInt(0)) == 0 {
			// automated rebalancing is disabled if target is set to 0
			reqFundingNow[net] = big.NewInt(0)
			reqFundingLater[net] = big.NewInt(0)
			continue
		}

		reqFundingNow[net] = big.NewInt(0).Sub(dataNow.TargetLiquidity, dataNow.Liquidity)
		reqFundingLater[net] = big.NewInt(0).Sub(dataLater.TargetLiquidity, dataLater.Liquidity)
	}

	return reqFundingNow, reqFundingLater, nil
}

func (r *TargetBalanceRebalancer) find1hopPotentialDonations(
	graphLater graph.Graph, // the networks graph state after all transfers are applied
	donateTo models.NetworkSelector, // target network
	reqFundingNow map[models.NetworkSelector]*big.Int, // the token funding requirements for each network
	reqFundingLater map[models.NetworkSelector]*big.Int, // the token funding requirements after all pending txs are applied
) ([]donation, error) {
	allEdges, err := graphLater.GetEdges()
	if err != nil {
		return nil, fmt.Errorf("get edges: %w", err)
	}

	potentialDonations := make([]donation, 0)
	seenDonors := mapset.NewSet[models.NetworkSelector]()

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

		donationAmount := big.NewInt(0).Sub(big.NewInt(0), funding)
		if donationAmount.Cmp(big.NewInt(0)) <= 0 {
			continue
		}

		potentialDonations = append(potentialDonations, newDonation(edge.Source, donateTo, donationAmount))
		seenDonors.Add(edge.Source)
	}

	return potentialDonations, nil
}

// request2HopDonations finds networks that can increase liquidity of the target network with an intermediate network.
func (r *TargetBalanceRebalancer) find2hopPotentialDonations(
	graphLater graph.Graph, // the networks graph state after all transfers are applied
	donateTo models.NetworkSelector, // target network
	reqFundingNow map[models.NetworkSelector]*big.Int, // the token funding requirements for each network
	reqFundingLater map[models.NetworkSelector]*big.Int, // the token funding requirements after all pending txs are applied
) ([]donation, error) {
	potentialDonations := make([]donation, 0)
	seenDonors := mapset.NewSet[models.NetworkSelector]()

	for _, net := range graphLater.GetNetworks() {
		if net == donateTo {
			continue
		}
		if seenDonors.Contains(net) {
			// cannot have the same donor twice
			continue
		}

		neibs, ok := graphLater.GetNeighbors(net)
		if !ok {
			return nil, fmt.Errorf("get neighbors of %d failed", net)
		}
		neibsSet := mapset.NewSet[models.NetworkSelector](neibs...)
		if neibsSet.Contains(donateTo) {
			// since the target network is a direct network we can donate using 1hop
			continue
		}

		for _, neib := range neibs {
			intermNeibs, ok := graphLater.GetNeighbors(neib)
			if !ok {
				return nil, fmt.Errorf("get intermediate neighbors of %d failed", net)
			}
			finalNeibsSet := mapset.NewSet[models.NetworkSelector](intermNeibs...)
			if finalNeibsSet.Contains(donateTo) {
				fundingNow, exists := reqFundingNow[net]
				if !exists {
					return nil, fmt.Errorf("net %d does not exist in the tokens to target offset", net)
				}
				funding := fundingNow

				fundingLater, exists := reqFundingLater[net]
				if !exists {
					return nil, fmt.Errorf("net %d does not exist in the tokens to target offset", net)
				}

				// If the balance is expected to become lower, we consider the lower balance to prevent a race condition in the donations.
				// If the balance is expected to become higher, we do not consider it since the funds are not available yet.
				if fundingNow.Cmp(fundingLater) < 0 {
					funding = fundingLater
				}

				donationAmount := big.NewInt(0).Sub(big.NewInt(0), funding)
				if donationAmount.Cmp(big.NewInt(0)) <= 0 {
					continue
				}

				seenDonors.Add(net)
				potentialDonations = append(potentialDonations, newDonation(net, neib, donationAmount))
			}
		}
	}

	return potentialDonations, nil
}

// apply changes to the intermediate state to prevent invalid donations
func (r *TargetBalanceRebalancer) acceptDonations(graphLater graph.Graph, potentialDonations []donation, requiredAmount *big.Int) ([]models.ProposedTransfer, error) {
	// sort by amount,donor,receiver
	sort.Slice(potentialDonations, func(i, j int) bool {
		if potentialDonations[i].amount.Cmp(potentialDonations[j].amount) == 0 {
			if potentialDonations[i].donor == potentialDonations[j].donor {
				return potentialDonations[i].receiver < potentialDonations[j].receiver
			}
			return potentialDonations[i].donor < potentialDonations[j].donor
		}
		return potentialDonations[i].amount.Cmp(potentialDonations[j].amount) > 0
	})

	fundsRaised := big.NewInt(0)
	proposedTransfers := make([]models.ProposedTransfer, 0, len(potentialDonations))
	skip := false
	for _, d := range potentialDonations {
		if skip {
			r.lggr.Debugf("skipping donation: %s", d)
			continue
		}

		// increment the raised funds
		fundsRaised = big.NewInt(0).Add(fundsRaised, d.amount)

		// in case we raised more than target amount give refund to the donor
		if refund := big.NewInt(0).Sub(fundsRaised, requiredAmount); refund.Cmp(big.NewInt(0)) > 0 {
			d.amount = big.NewInt(0).Sub(d.amount, refund)
			fundsRaised = big.NewInt(0).Sub(fundsRaised, refund)
		}
		r.lggr.Debugf("accepting donation: %s", d)
		proposedTransfers = append(proposedTransfers, models.ProposedTransfer{From: d.donor, To: d.receiver, Amount: ubig.New(d.amount)})

		r.lggr.Debugf("applying donation to future graph state")
		liqBefore, err := graphLater.GetLiquidity(d.receiver)
		if err != nil {
			return nil, fmt.Errorf("get liquidity of donation receiver %d: %w", d.receiver, err)
		}
		graphLater.SetLiquidity(d.receiver, big.NewInt(0).Add(liqBefore, d.amount))

		liqBefore, err = graphLater.GetLiquidity(d.donor)
		if err != nil {
			return nil, fmt.Errorf("get liquidity of donor %d: %w", d.donor, err)
		}
		graphLater.SetLiquidity(d.donor, big.NewInt(0).Sub(liqBefore, d.amount))

		if fundsRaised.Cmp(requiredAmount) >= 0 {
			r.lggr.Debugf("all funds raised skipping further donations")
			skip = true
		}
	}

	return proposedTransfers, nil
}

// getExpectedGraph returns the a copy of the graph instance with all the non executed transfers applied.
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

// mergeProposedTransfers merges multiple transfers with same sender and recipient into a single transfer.
func (r *TargetBalanceRebalancer) mergeProposedTransfers(transfers []models.ProposedTransfer) []models.ProposedTransfer {
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

type donation struct {
	donor    models.NetworkSelector
	receiver models.NetworkSelector
	amount   *big.Int
}

func newDonation(donor, receiver models.NetworkSelector, amount *big.Int) donation {
	return donation{
		donor:    donor,
		receiver: receiver,
		amount:   amount,
	}
}

func (d donation) String() string {
	return fmt.Sprintf("%d donates %s tokens to %d", d.donor, d.amount.String(), d.receiver)
}
