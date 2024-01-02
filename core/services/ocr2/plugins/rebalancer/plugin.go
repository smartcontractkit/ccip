package rebalancer

import (
	"context"
	"fmt"
	"math/big"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/liquiditygraph"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/liquiditymanager"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/liquidityrebalancer"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

type Plugin struct {
	liquidityManagers       map[models.NetworkID]models.Address
	liquidityManagerFactory liquiditymanager.Factory
	pendingTransfers        []models.PendingTransfer
	liquidityGraph          liquiditygraph.LiquidityGraph
	liquidityRebalancer     liquidityrebalancer.Rebalancer
}

func NewPlugin(
	liquidityManagerNetwork models.NetworkID,
	liquidityManagerAddress models.Address,
	liquidityManagerFactory liquiditymanager.Factory,
	liquidityGraph liquiditygraph.LiquidityGraph,
	liquidityRebalancer liquidityrebalancer.Rebalancer,
) *Plugin {
	return &Plugin{
		liquidityManagers: map[models.NetworkID]models.Address{
			liquidityManagerNetwork: liquidityManagerAddress,
		},
		liquidityManagerFactory: liquidityManagerFactory,
		pendingTransfers:        make([]models.PendingTransfer, 0), // todo: thread-safe
		liquidityGraph:          liquidityGraph,
		liquidityRebalancer:     liquidityRebalancer,
	}
}

func (p *Plugin) Query(_ context.Context, _ ocr3types.OutcomeContext) (ocrtypes.Query, error) {
	return ocrtypes.Query{}, nil
}

func (p *Plugin) Observation(ctx context.Context, outctx ocr3types.OutcomeContext, query ocrtypes.Query) (ocrtypes.Observation, error) {
	if err := p.syncGraphEdges(ctx); err != nil {
		return ocrtypes.Observation{}, fmt.Errorf("sync graph edges: %w", err)
	}

	// todo: return the graph balances after syncing them
	if err := p.syncGraphBalances(ctx); err != nil {
		return ocrtypes.Observation{}, fmt.Errorf("sync graph balances: %w", err)
	}

	// todo: return the pending transfers after loading them
	if err := p.loadPendingTransfers(ctx); err != nil {
		return ocrtypes.Observation{}, fmt.Errorf("load pending transfers: %w", err)
	}

	networks := p.liquidityGraph.GetNodes()
	liquidityPerChain := make([]models.ChainLiquidity, 0, len(networks))
	for _, network := range networks {
		liq, err := p.liquidityGraph.GetWeight(network)
		if err != nil {
			return ocrtypes.Observation{}, fmt.Errorf("get network %v weight: %w", network, err)
		}
		liquidityPerChain = append(liquidityPerChain, models.NewChainLiquidity(network, liq))
	}

	return models.NewObservation(liquidityPerChain, p.pendingTransfers).Encode(), nil
}

func (p *Plugin) ValidateObservation(outctx ocr3types.OutcomeContext, query ocrtypes.Query, ao ocrtypes.AttributedObservation) error {
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
	// todo: consensus on observations
	outcome := models.NewObservation(nil, nil).Encode()
	return outcome, fmt.Errorf("not implemented")
}

func (p *Plugin) Reports(seqNr uint64, outcome ocr3types.Outcome) ([]ocr3types.ReportWithInfo[models.ReportMetadata], error) {
	obs, err := models.DecodeObservation(outcome)
	if err != nil {
		return nil, fmt.Errorf("decode outcome: %w", err)
	}

	transfersToReachBalance, err := p.liquidityRebalancer.ComputeTransfersToBalance(
		p.liquidityGraph, obs.PendingTransfers)
	if err != nil {
		return nil, fmt.Errorf("compute transfers to reach balance: %w", err)
	}

	// group transfers by source chain
	transfersBySourceChain := make(map[models.NetworkID][]models.Transfer)
	for _, tr := range transfersToReachBalance {
		transfersBySourceChain[tr.From] = append(transfersBySourceChain[tr.From], tr)
	}

	var reports []ocr3types.ReportWithInfo[models.ReportMetadata]
	for sourceChain, transfers := range transfersBySourceChain {
		lmAddress, exists := p.liquidityManagers[sourceChain]
		if !exists {
			return nil, fmt.Errorf("liquidity manager for %v does not exist", sourceChain)
		}

		reportMeta := models.NewReportMetadata(transfers, lmAddress, sourceChain)
		reports = append(reports, ocr3types.ReportWithInfo[models.ReportMetadata]{
			Report: reportMeta.Encode(),
			Info:   reportMeta,
		})
	}

	return reports, nil
}

func (p *Plugin) ShouldAcceptAttestedReport(ctx context.Context, u uint64, r ocr3types.ReportWithInfo[models.ReportMetadata]) (bool, error) {
	reportMeta, err := models.DecodeReportMetadata(r.Report)
	if err != nil {
		return false, fmt.Errorf("decode report metadata: %w", err)
	}

	fmt.Println(reportMeta.Transfers)
	// todo: check if reportMeta.transfers are valid

	return true, nil
}

func (p *Plugin) ShouldTransmitAcceptedReport(ctx context.Context, u uint64, r ocr3types.ReportWithInfo[models.ReportMetadata]) (bool, error) {
	newPendingTransfers := make([]models.PendingTransfer, 0, len(r.Info.Transfers))
	for _, tr := range r.Info.Transfers {
		//if slices.Contains(p.pendingTransfers, tr) { // todo: use a struct for this ops
		//	return false, nil
		//}
		newPendingTransfers = append(newPendingTransfers, models.NewPendingTransfer(tr))
	}

	p.pendingTransfers = append(p.pendingTransfers, newPendingTransfers...)
	return true, nil
}

func (p *Plugin) Close() error {
	// todo: init a ctx with timeout

	for networkID, lmAddr := range p.liquidityManagers {
		// todo: lmCloser := liquidityManagerFactory.NewLiquidityManagerCloser(); lmCloser.Close()
		lm, err := p.liquidityManagerFactory.NewLiquidityManager(networkID, lmAddr)
		if err != nil {
			return err
		}

		if err := lm.Close(context.TODO()); err != nil {
			return err
		}
	}

	return nil
}

func (p *Plugin) syncGraphEdges(ctx context.Context) error {
	// todo: if there wasn't any change to the graph stop earlier

	p.liquidityGraph.Reset()

	type qItem struct {
		networkID models.NetworkID
		lmAddress models.Address
	}

	seen := mapset.NewSet[qItem]()
	queue := mapset.NewSet[qItem]()
	for networkID, lmAddress := range p.liquidityManagers {
		elem := qItem{networkID: networkID, lmAddress: lmAddress}
		queue.Add(elem)
		seen.Add(elem)
	}

	for queue.Cardinality() > 0 {
		elem, ok := queue.Pop()
		if !ok {
			return errors.New("unexpected internal error, there is a bug in the algorithm")
		}

		p.liquidityGraph.AddNode(elem.networkID, big.NewInt(0)) // TODO: investigate fetching the balance here.

		lm, err := p.liquidityManagerFactory.NewLiquidityManager(elem.networkID, elem.lmAddress)
		if err != nil {
			return fmt.Errorf("init liquidity manager: %w", err)
		}

		destinationLMs, err := lm.GetLiquidityManagers(ctx)
		if err != nil {
			return fmt.Errorf("get %v destination liquidity managers: %w", elem.networkID, err)
		}

		for destNetworkID, lmAddr := range destinationLMs {
			p.liquidityGraph.AddEdge(elem.networkID, destNetworkID)

			newElem := qItem{networkID: destNetworkID, lmAddress: lmAddr}
			if !seen.Contains(newElem) {
				queue.Add(newElem)
				seen.Add(newElem)

				if _, exists := p.liquidityManagers[destNetworkID]; !exists {
					p.liquidityManagers[destNetworkID] = lmAddr
				}
			}
		}
	}

	return nil
}

func (p *Plugin) syncGraphBalances(ctx context.Context) error {
	for _, networkID := range p.liquidityGraph.GetNodes() {
		lmAddr, exists := p.liquidityManagers[networkID]
		if !exists {
			return fmt.Errorf("liquidity manager for network %v was not found", networkID)
		}

		lm, err := p.liquidityManagerFactory.NewLiquidityManager(networkID, lmAddr)
		if err != nil {
			return fmt.Errorf("init liquidity manager: %w", err)
		}

		balance, err := lm.GetBalance(ctx)
		if err != nil {
			return fmt.Errorf("get %v balance: %w", networkID, err)
		}

		p.liquidityGraph.SetWeight(networkID, balance)
	}

	return nil
}

func (p *Plugin) loadPendingTransfers(ctx context.Context) error {
	// todo: do not load pending transfers all the time

	p.pendingTransfers = make([]models.PendingTransfer, 0)

	for networkID, lmAddress := range p.liquidityManagers {
		lm, err := p.liquidityManagerFactory.NewLiquidityManager(networkID, lmAddress)
		if err != nil {
			return fmt.Errorf("init liquidity manager: %w", err)
		}

		pendingTransfers, err := lm.GetPendingTransfers(ctx)
		if err != nil {
			return fmt.Errorf("get pending %v transfers: %w", networkID, err)
		}

		p.pendingTransfers = append(p.pendingTransfers, pendingTransfers...)
	}

	return nil
}
