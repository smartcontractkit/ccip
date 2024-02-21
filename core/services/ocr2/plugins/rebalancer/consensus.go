package rebalancer

import (
	"fmt"
	"math/big"
	"sort"
	"time"

	ubig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/bridge"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/graph"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

func computePendingTransfersConsensus(observations []models.Observation, f int) ([]models.PendingTransfer, error) {
	eventFromHash := make(map[[32]byte]models.PendingTransfer)
	counts := make(map[[32]byte]int)
	for _, obs := range observations {
		for _, tr := range obs.PendingTransfers {
			h, err := tr.Hash()
			if err != nil {
				return nil, fmt.Errorf("hash %v: %w", tr, err)
			}
			counts[h]++
			eventFromHash[h] = tr
		}
	}

	var quorumEvents []models.PendingTransfer
	for h, count := range counts {
		if count >= f+1 {
			ev, exists := eventFromHash[h]
			if !exists {
				return nil, fmt.Errorf("internal issue, event from hash %v not found", h)
			}
			quorumEvents = append(quorumEvents, ev)
		}
	}

	return quorumEvents, nil
}

func computeConfigDigestsConsensus(observations []models.Observation, f int) ([]models.ConfigDigestWithMeta, error) {
	key := func(meta models.ConfigDigestWithMeta) string {
		return fmt.Sprintf("%d-%s-%s", meta.NetworkSel, meta.RebalancerAddr, meta.Digest.Hex())
	}
	counts := make(map[string]int)
	cds := make(map[string]models.ConfigDigestWithMeta)
	for _, obs := range observations {
		for _, cd := range obs.ConfigDigests {
			k := key(cd)
			counts[k]++
			if counts[k] == 1 {
				cds[k] = cd
			}
		}
	}

	var quorumCds []models.ConfigDigestWithMeta
	for k, count := range counts {
		if count >= f+1 {
			cd, exists := cds[k]
			if !exists {
				return nil, fmt.Errorf("internal issue, config digest by key %s not found", k)
			}
			quorumCds = append(quorumCds, cd)
		}
	}

	// sort by network id for deterministic results
	sort.Slice(quorumCds, func(i, j int) bool {
		return quorumCds[i].NetworkSel < quorumCds[j].NetworkSel
	})

	return quorumCds, nil
}

func computeGraphEdgesConsensus(observations []models.Observation, f int) []models.Edge {
	counts := make(map[models.Edge]int)
	for _, obs := range observations {
		for _, edge := range obs.Edges {
			counts[edge]++
		}
	}

	var quorumEdges []models.Edge
	for edge, count := range counts {
		if count >= f+1 {
			quorumEdges = append(quorumEdges, edge)
		}
	}

	return quorumEdges
}

func computeMedianLiquidityPerChain(observations []models.Observation) []models.NetworkLiquidity {
	liqObsPerChain := make(map[models.NetworkSelector][]*big.Int)
	for _, ob := range observations {
		for _, chainLiq := range ob.LiquidityPerChain {
			liqObsPerChain[chainLiq.Network] = append(liqObsPerChain[chainLiq.Network], chainLiq.Liquidity.ToInt())
		}
	}

	medians := make([]models.NetworkLiquidity, 0, len(liqObsPerChain))
	for chainID, liqs := range liqObsPerChain {
		medians = append(medians, models.NewNetworkLiquidity(chainID, bigIntSortedMiddle(liqs)))
	}
	// sort by network id for deterministic results
	sort.Slice(medians, func(i, j int) bool {
		return medians[i].Network < medians[j].Network
	})
	return medians
}

func computeMedianGraph(lggr logger.Logger, edges []models.Edge, medianLiquidities []models.NetworkLiquidity) (graph.Graph, error) {
	g, err := graph.NewGraphFromEdges(edges)
	if err != nil {
		return nil, fmt.Errorf("new graph from edges: %w", err)
	}

	for _, medianLiq := range medianLiquidities {
		if !g.SetLiquidity(medianLiq.Network, medianLiq.Liquidity.ToInt()) {
			lggr.Errorw("median liquidity on network not found on edges quorum", "net", medianLiq.Network)
		}
	}

	return g, nil
}

func computeResolvedTransfersQuorum(lggr logger.Logger, observations []models.Observation, f int, bridgeFactory bridge.Factory) ([]models.Transfer, error) {
	// assumption: there shouldn't be more than 1 transfer for a (from, to) pair from a single oracle's observation.
	// otherwise they can be collapsed into a single transfer.
	// TODO: we can check for this in ValidateObservation
	type key struct {
		From               models.NetworkSelector
		To                 models.NetworkSelector
		Amount             string
		Sender             models.Address
		Receiver           models.Address
		LocalTokenAddress  models.Address
		RemoteTokenAddress models.Address
	}
	counts := make(map[key][]models.Transfer)
	for _, obs := range observations {
		lggr.Debugw("observed transfers", "transfers", obs.ResolvedTransfers)
		for _, tr := range obs.ResolvedTransfers {
			lggr.Debugw("inserting resolved transfer into mapping", "transfer", tr)
			k := key{
				From:               tr.From,
				To:                 tr.To,
				Amount:             tr.Amount.String(),
				Sender:             tr.Sender,
				Receiver:           tr.Receiver,
				LocalTokenAddress:  tr.LocalTokenAddress,
				RemoteTokenAddress: tr.RemoteTokenAddress,
			}
			counts[k] = append(counts[k], tr)
		}
	}

	lggr.Debugw("resolved transfers counts", "counts", len(counts))

	var quorumTransfers []models.Transfer
	for k, transfers := range counts {
		if len(transfers) >= f+1 {
			lggr.Debugw("quorum reached on transfer", "transfer", k, "votes", len(transfers))
			// need to compute the "medianized" bridge payload
			// only the bridge knows how to do this so we need to delegate it to them
			// the native bridge fee can also be medianized, no need for the bridge to do that
			var (
				bridgeFees     []*big.Int
				bridgePayloads [][]byte
				datesUnix      []*big.Int
			)
			for _, tr := range transfers {
				bridgeFees = append(bridgeFees, tr.NativeBridgeFee.ToInt())
				bridgePayloads = append(bridgePayloads, tr.BridgeData)
				datesUnix = append(datesUnix, big.NewInt(tr.Date.Unix()))
			}
			medianizedNativeFee := bigIntSortedMiddle(bridgeFees)
			medianizedDateUnix := bigIntSortedMiddle(datesUnix)
			bridge, err := bridgeFactory.NewBridge(k.From, k.To)
			if err != nil {
				return nil, fmt.Errorf("init bridge: %w", err)
			}
			quorumizedBridgePayload, err := bridge.QuorumizedBridgePayload(bridgePayloads, f)
			if err != nil {
				return nil, fmt.Errorf("quorumized bridge payload: %w", err)
			}
			quorumTransfer := models.Transfer{
				From:               k.From,
				To:                 k.To,
				Amount:             transfers[0].Amount,
				Date:               time.Unix(medianizedDateUnix.Int64(), 0), // medianized, not in the key
				Sender:             transfers[0].Sender,
				Receiver:           transfers[0].Receiver,
				LocalTokenAddress:  transfers[0].LocalTokenAddress,
				RemoteTokenAddress: transfers[0].RemoteTokenAddress,
				BridgeData:         quorumizedBridgePayload,       // "quorumized", not in the key
				NativeBridgeFee:    ubig.New(medianizedNativeFee), // medianized, not in the key
			}
			quorumTransfers = append(quorumTransfers, quorumTransfer)
		} else {
			lggr.Debugw("dropping transfer, not enough votes on it", "transfer", k, "votes", len(transfers))
		}
	}

	return quorumTransfers, nil
}

func computeInflightTransfersQuorum(lggr logger.Logger, observations []models.Observation, f int) ([]models.Transfer, error) {
	type key struct {
		From               models.NetworkSelector
		To                 models.NetworkSelector
		Amount             string
		Sender             models.Address
		Receiver           models.Address
		LocalTokenAddress  models.Address
		RemoteTokenAddress models.Address
	}
	counts := make(map[key][]models.Transfer)
	for _, obs := range observations {
		lggr.Infow("observed transfers", "transfers", obs.InflightTransfers)
		for _, tr := range obs.InflightTransfers {
			lggr.Infow("inserting inflight transfer into mapping", "transfer", tr)
			k := key{
				From:               tr.From,
				To:                 tr.To,
				Amount:             tr.Amount.String(),
				Sender:             tr.Sender,
				Receiver:           tr.Receiver,
				LocalTokenAddress:  tr.LocalTokenAddress,
				RemoteTokenAddress: tr.RemoteTokenAddress,
			}
			counts[k] = append(counts[k], tr)
		}
	}

	lggr.Infow("inflight transfers counts", "counts", len(counts))

	var quorumTransfers []models.Transfer
	for k, transfers := range counts {
		if len(transfers) >= f+1 {
			quorumTransfer := models.Transfer{
				From:               k.From,
				To:                 k.To,
				Amount:             transfers[0].Amount,
				Date:               transfers[0].Date,
				Sender:             transfers[0].Sender,
				Receiver:           transfers[0].Receiver,
				LocalTokenAddress:  transfers[0].LocalTokenAddress,
				RemoteTokenAddress: transfers[0].RemoteTokenAddress,
				BridgeData:         transfers[0].BridgeData,
				NativeBridgeFee:    transfers[0].NativeBridgeFee,
			}
			quorumTransfers = append(quorumTransfers, quorumTransfer)
			lggr.Infow("quorum reached on transfer", "transfer", k, "votes", len(transfers), "quorumTransfer", quorumTransfer)
		} else {
			lggr.Infow("dropping transfer, not enough votes on it", "transfer", k, "votes", len(transfers))
		}
	}

	// sort by network id for deterministic results
	sort.Slice(quorumTransfers, func(i, j int) bool {
		return quorumTransfers[i].From < quorumTransfers[j].From
	})

	return quorumTransfers, nil
}

func removeInflightTransfers(lggr logger.Logger, pendingTransfers []models.PendingTransfer, inflightTransfers []models.Transfer) []models.PendingTransfer {
	var filtered []models.PendingTransfer
	for _, pending := range pendingTransfers {
		var found bool
		for _, inflight := range inflightTransfers {
			if pending.Transfer.From == inflight.From && pending.Transfer.To == inflight.To {
				lggr.Infow("removing inflight transfer from pending transfers", "inflight", inflight, "pending", pending)
				found = true
				break
			}
		}
		if !found {
			filtered = append(filtered, pending)
		}
	}
	return filtered
}
