package liquidityrebalancer

import (
	"fmt"
	"math/big"
	"sort"

	mapset "github.com/deckarep/golang-set/v2"

	big2 "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/liquiditymanager/graph"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/liquiditymanager/models"
)

// TargetMinBalancer tries to reach balance using a target and minimum liquidity that is configured on each network.
type TargetMinBalancer struct {
	lggr logger.Logger
}

func NewTargetMinBalancer(lggr logger.Logger) *TargetMinBalancer {
	return &TargetMinBalancer{
		lggr: lggr,
	}
}

type determineTransfersFunc func(graphLater graph.Graph, targetNetwork models.NetworkSelector, networkFunds map[models.NetworkSelector]*Funds) ([]models.ProposedTransfer, error)

func (r *TargetMinBalancer) ComputeTransfersToBalance(graphNow graph.Graph, nonExecutedTransfers []UnexecutedTransfer) ([]models.ProposedTransfer, error) {
	nonExecutedTransfers = filterUnexecutedTransfers(nonExecutedTransfers)

	var proposedTransfers []models.ProposedTransfer
	// 4 rounds of rebalancing alternate between 1 hop and 2 hop transfers
	for i := 0; i < 5; i++ {
		fmt.Println()
		r.lggr.Debug("Round ", i)
		r.lggr.Debugf("nonExecutedTransfers: %v", nonExecutedTransfers)
		var currentProposed []models.ProposedTransfer
		transfersFunc := r.oneHopTransfers
		if i%2 != 0 {
			transfersFunc = r.twoHopTransfers
		}
		currentProposed, err := r.rebalancingRound(graphNow, nonExecutedTransfers, transfersFunc)
		if err != nil {
			return nil, err
		}
		r.lggr.Debugf("current round proposed transfers: %v", currentProposed)
		for _, t := range currentProposed {
			// put proposed in nonExecutedTransfers to carryover to next round
			nonExecutedTransfers = append(nonExecutedTransfers, t)
		}
		proposedTransfers = append(proposedTransfers, currentProposed...)

	}

	r.lggr.Debugf("merging proposed transfers")
	proposedTransfers = mergeProposedTransfers(proposedTransfers)
	r.lggr.Debugf("sorting proposed transfers for determinism")
	sort.Sort(models.ProposedTransfers(proposedTransfers))

	return proposedTransfers, nil
}

func (r *TargetMinBalancer) rebalancingRound(graphNow graph.Graph, nonExecutedTransfers []UnexecutedTransfer, transfersFunc determineTransfersFunc) ([]models.ProposedTransfer, error) {
	var err error
	graphLater, err := getExpectedGraph(graphNow, nonExecutedTransfers)
	if err != nil {
		return nil, fmt.Errorf("get expected graph: %w", err)
	}

	r.lggr.Debugf("finding networks that require funding")
	networksRequiringFunding, networkFunds, err := r.findNetworksRequiringFunding(graphNow, graphLater)
	if err != nil {
		return nil, fmt.Errorf("find networks that require funding: %w", err)
	}

	proposedTransfers := make([]models.ProposedTransfer, 0)
	for _, net := range networksRequiringFunding {
		r.lggr.Debugf("finding transfers for network %v", net)
		var potentialTransfers []models.ProposedTransfer
		potentialTransfers, err = transfersFunc(graphLater, net, networkFunds)
		if err != nil {
			return nil, fmt.Errorf("finding transfers for network %v: %w", net, err)
		}

		netProposedTransfers, err := r.applyProposedTransfers(graphLater, potentialTransfers, networkFunds[net].LiqDiffLater)
		if err != nil {
			return nil, fmt.Errorf("applying transfers: %w", err)
		}
		proposedTransfers = append(proposedTransfers, netProposedTransfers...)
	}

	return proposedTransfers, nil
}

func (r *TargetMinBalancer) findNetworksRequiringFunding(graphNow, graphLater graph.Graph) ([]models.NetworkSelector, map[models.NetworkSelector]*Funds, error) {
	mapNetworkFunds := make(map[models.NetworkSelector]*Funds)
	liqDiffsNow, liqDiffsLater, err := getTargetLiquidityDifferences(graphNow, graphLater)
	if err != nil {
		return nil, nil, fmt.Errorf("compute tokens funding requirements: %w", err)
	}

	//TODO: LM-23 Create minTokenTransfer config to filter-out small rebalance txs
	// check that the transfer is not tiny, we should only transfer if it is significant. What is too tiny?
	// we could prevent this by only making a network requiring funding if its below X% of the target

	res := make([]models.NetworkSelector, 0, len(liqDiffsNow))
	for net := range liqDiffsNow {
		diffLater := liqDiffsLater[net]

		//use min here for transferable. because we don't know when the transfers will complete and want to avoid issues
		transferableAmount, ataErr := availableTransferableAmount(graphNow, graphLater, net)
		if ataErr != nil {
			r.lggr.Debugf("error getting available transferrable amount for net %d: %v", net, ataErr)
			continue
		}
		mapNetworkFunds[net] = &Funds{
			LiqDiffNow:      liqDiffsNow[net],
			LiqDiffLater:    liqDiffsLater[net],
			AvailableAmount: transferableAmount,
		}

		if diffLater.Cmp(big.NewInt(0)) <= 0 {
			r.lggr.Debugf("net %v does not require funding, transferrable tokens: %d", net, transferableAmount)
			continue
		}

		r.lggr.Debugf("net %v requires funding, %s tokens to reach target, transferrable tokens: %d", net, diffLater, transferableAmount)
		res = append(res, net)
	}

	sort.Slice(res, func(i, j int) bool { return liqDiffsLater[res[i]].Cmp(liqDiffsLater[res[j]]) > 0 })
	return res, mapNetworkFunds, nil
}

func (r *TargetMinBalancer) oneHopTransfers(graphLater graph.Graph, targetNetwork models.NetworkSelector, networkFunds map[models.NetworkSelector]*Funds) ([]models.ProposedTransfer, error) {
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

		destDiffLater := networkFunds[edge.Dest].LiqDiffLater

		r.lggr.Debugf("checking transfer from %v to %v", edge.Source, edge.Dest)

		//source network available transferable amount
		srcData, dErr := graphLater.GetData(edge.Source)
		if dErr != nil {
			return nil, fmt.Errorf("error during GetData for %v in graphLater: %v", targetNetwork, dErr)
		}
		srcAvailableAmount := big.NewInt(0).Sub(srcData.Liquidity, srcData.MinimumLiquidity)
		srcAmountToTarget := big.NewInt(0).Sub(srcData.Liquidity, srcData.TargetLiquidity)

		if srcAmountToTarget.Cmp(big.NewInt(0)) < 0 || srcAvailableAmount.Cmp(big.NewInt(0)) < 0 {
			//r.lggr.Debugf("source network %v is too low, skipping transfer: srcAmountToTarget %v, srcAvailableAmount %v", edge.Source, srcAmountToTarget, srcAvailableAmount)
			continue
		}

		transferAmount := destDiffLater
		if transferAmount.Cmp(srcAmountToTarget) > 0 {
			// if transferAmount > srcAmountToTarget take less
			r.lggr.Debugf("source network %v doesn't have %v, taking the available %v instead", edge.Source, transferAmount, srcAmountToTarget)
			transferAmount = srcAmountToTarget
		}
		if transferAmount.Cmp(big.NewInt(0)) <= 0 {
			//r.lggr.Debugf("transfer %v->%v amount is 0 or less, skipping transfer: %v", edge.Source, targetNetwork, transferAmount)
			continue
		}
		if srcAmountToTarget.Cmp(transferAmount) < 0 || srcAvailableAmount.Cmp(transferAmount) < 0 {
			// source network doesn't have enough to cover
			//r.lggr.Debugf("source network %v liquidity too low, skipping transfer: srcAmountToTarget %v, srcAvailableAmount %v", edge.Source, srcAmountToTarget, srcAvailableAmount)
			continue
		}

		newAmount := big.NewInt(0).Sub(networkFunds[edge.Source].AvailableAmount, transferAmount)
		if newAmount.Cmp(big.NewInt(0)) < 0 {
			r.lggr.Debugf("source network %v doesn't have enough available liquidity, skipping transfer %v only have %v available", edge.Source, transferAmount, networkFunds[edge.Source].AvailableAmount)
			continue
		} else {
			networkFunds[edge.Source].AvailableAmount = newAmount
		}
		potentialTransfers = append(potentialTransfers, newTransfer(edge.Source, targetNetwork, transferAmount))
		seenNetworks.Add(edge.Source)
	}

	return potentialTransfers, nil
}

// twoHopTransfers finds networks that can increase liquidity of the target network with an intermediate network.
func (r *TargetMinBalancer) twoHopTransfers(graphLater graph.Graph, targetNetwork models.NetworkSelector, networkFunds map[models.NetworkSelector]*Funds) ([]models.ProposedTransfer, error) {
	potentialTransfers := make([]models.ProposedTransfer, 0)
	seenNetworks := mapset.NewSet[models.NetworkSelector]()

	for _, src := range graphLater.GetNetworks() {
		if src == targetNetwork {
			continue
		}
		if seenNetworks.Contains(src) {
			// cannot have the same sender twice
			continue
		}

		neighbors, ok := graphLater.GetNeighbors(src, false)
		if !ok {
			return nil, fmt.Errorf("get neighbors of %d failed", src)
		}
		neighborsSet := mapset.NewSet[models.NetworkSelector](neighbors...)
		if neighborsSet.Contains(targetNetwork) {
			// since the target network is a direct network we can transfer using 1hop
			continue
		}

		for _, middle := range neighbors {
			intermediateNeighbors, ok := graphLater.GetNeighbors(middle, false)
			if !ok {
				return nil, fmt.Errorf("get intermediate neighbors of %d failed", src)
			}
			finalNeighborsSet := mapset.NewSet[models.NetworkSelector](intermediateNeighbors...)
			r.lggr.Debugf("checking transfer from %v to %v to %v", src, middle, targetNetwork)

			if finalNeighborsSet.Contains(targetNetwork) {
				fundingDest := networkFunds[targetNetwork].LiqDiffLater

				//source network available transferable amount
				srcData, dErr := graphLater.GetData(src)
				if dErr != nil {
					return nil, fmt.Errorf("error during GetData for %v in graphLater: %v", targetNetwork, dErr)
				}
				srcAvailableAmount := big.NewInt(0).Sub(srcData.Liquidity, srcData.MinimumLiquidity)
				srcAmountToTarget := big.NewInt(0).Sub(srcData.Liquidity, srcData.TargetLiquidity)

				//middle network available transferable amount
				middleData, dErr := graphLater.GetData(middle)
				if dErr != nil {
					return nil, fmt.Errorf("error during GetData for %v in graphLater: %v", targetNetwork, dErr)
				}
				middleAvailableAmount := big.NewInt(0).Sub(middleData.Liquidity, middleData.MinimumLiquidity)
				middleAmountToTarget := big.NewInt(0).Sub(middleData.Liquidity, middleData.TargetLiquidity)

				transferAmount := fundingDest
				if transferAmount.Cmp(srcAmountToTarget) > 0 {
					// if transferAmount > srcAmountToTarget take less
					transferAmount = srcAmountToTarget
				}
				if transferAmount.Cmp(big.NewInt(0)) <= 0 {
					//r.lggr.Debugf("transfer %v->%v amount is 0 or less, skipping transfer: %v", src, targetNetwork, transferAmount)
					continue
				}

				if srcAmountToTarget.Cmp(transferAmount) < 0 || srcAvailableAmount.Cmp(transferAmount) < 0 {
					// source network doesn't have enough to cover
					//r.lggr.Debugf("source network %v liquidity too low, skipping transfer: srcAmountToTarget %v, srcAvailableAmount %v", src, srcAmountToTarget, srcAvailableAmount)
					continue
				}
				if middleAvailableAmount.Cmp(transferAmount) < 0 {
					// middle hop doesn't have enough available liquidity
					r.lggr.Debugf("middle network %v liquidity too low, skipping transfer: middleAmountToTarget %v, middleAvailableAmount %v", middle, middleAmountToTarget, middleAvailableAmount)
					continue
				}

				seenNetworks.Add(src)
				potentialTransfers = append(potentialTransfers, newTransfer(src, middle, transferAmount))
			}
		}
	}

	return potentialTransfers, nil
}

// applyProposedTransfers applies the proposed transfers to the graph.
// increments the raised funds and gives a refund to the sender if more funds have been raised than the required amount.
// It updates the liquidity of the sender and receiver networks in the graph. It stops further transfers if all funds have been raised.
func (r *TargetMinBalancer) applyProposedTransfers(graphLater graph.Graph, potentialTransfers []models.ProposedTransfer, requiredAmount *big.Int) ([]models.ProposedTransfer, error) {
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
			return nil, fmt.Errorf("get liquidity of sender %v: %w", d.From, err)
		}
		availableAmount := big.NewInt(0).Sub(senderData.Liquidity, senderData.MinimumLiquidity)
		if availableAmount.Cmp(big.NewInt(0)) <= 0 {
			r.lggr.Debugf("no more tokens to transfer, skipping transfer: %s", d)
			continue
		}

		if availableAmount.Cmp(d.Amount.ToInt()) < 0 {
			d.Amount = big2.New(availableAmount)
			r.lggr.Debugf("reducing transfer amount since sender balance has dropped: %s", d)
		}

		// increment the raised funds
		fundsRaised = big.NewInt(0).Add(fundsRaised, d.Amount.ToInt())

		// in case we raised more than target amount give refund to the sender
		if refund := big.NewInt(0).Sub(fundsRaised, requiredAmount); refund.Cmp(big.NewInt(0)) > 0 {
			d.Amount = big2.New(big.NewInt(0).Sub(d.Amount.ToInt(), refund))
			fundsRaised = big.NewInt(0).Sub(fundsRaised, refund)
		}
		r.lggr.Debugf("applying transfer: %v", d)
		proposedTransfers = append(proposedTransfers, models.ProposedTransfer{From: d.From, To: d.To, Amount: d.Amount})

		liqBefore, err := graphLater.GetLiquidity(d.To)
		if err != nil {
			return nil, fmt.Errorf("get liquidity of transfer receiver %v: %w", d.To, err)
		}
		graphLater.SetLiquidity(d.To, big.NewInt(0).Add(liqBefore, d.Amount.ToInt()))

		liqBefore, err = graphLater.GetLiquidity(d.From)
		if err != nil {
			return nil, fmt.Errorf("get liquidity of sender %v: %w", d.From, err)
		}
		graphLater.SetLiquidity(d.From, big.NewInt(0).Sub(liqBefore, d.Amount.ToInt()))

		if fundsRaised.Cmp(requiredAmount) >= 0 {
			r.lggr.Debugf("all funds raised skipping further transfers")
			skip = true
		}
	}

	return proposedTransfers, nil
}
