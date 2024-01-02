package rebalancer

import (
	"context"
	"encoding/json"
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

	if err := p.syncGraphBalances(ctx); err != nil {
		return ocrtypes.Observation{}, fmt.Errorf("sync graph balances: %w", err)
	}

	if err := p.loadPendingTransfers(ctx); err != nil {
		return ocrtypes.Observation{}, fmt.Errorf("load pending transfers: %w", err)
	}

	transfersToBalance, err := p.liquidityRebalancer.ComputeTransfersToBalance(p.liquidityGraph, p.pendingTransfers)
	if err != nil {
		return ocrtypes.Observation{}, fmt.Errorf("compute transfers to balance: %w", err)
	}

	if len(transfersToBalance) == 0 {
		return nil, nil
	}
	return json.Marshal(transfersToBalance)
}

func (p *Plugin) ValidateObservation(outctx ocr3types.OutcomeContext, query ocrtypes.Query, ao ocrtypes.AttributedObservation) error {
	var transfers []models.Transfer
	if err := json.Unmarshal(ao.Observation, &transfers); err != nil {
		return fmt.Errorf("invalid observation: %w", err)
	}

	if len(transfers) == 0 {
		return errors.New("empty observation")
	}

	return nil
}

func (p *Plugin) ObservationQuorum(outctx ocr3types.OutcomeContext, query ocrtypes.Query) (ocr3types.Quorum, error) {
	return ocr3types.QuorumTwoFPlusOne, nil
}

func (p *Plugin) Outcome(outctx ocr3types.OutcomeContext, query ocrtypes.Query, aos []ocrtypes.AttributedObservation) (ocr3types.Outcome, error) {
	// todo: consensus on observations
	return ocr3types.Outcome{}, fmt.Errorf("not implemented")
}

func (p *Plugin) Reports(seqNr uint64, outcome ocr3types.Outcome) ([]ocr3types.ReportWithInfo[models.ReportMetadata], error) {
	var transfersToReachBalance []models.Transfer
	if err := json.Unmarshal(outcome, &transfersToReachBalance); err != nil {
		return nil, fmt.Errorf("parse outcome: %w", err)
	}

	var reports []ocr3types.ReportWithInfo[models.ReportMetadata]

	for _, transfer := range transfersToReachBalance {
		lmAddress, exists := p.liquidityManagers[transfer.From]
		if !exists {
			return nil, fmt.Errorf("liquidity manager for %v does not exist", transfer.From)
		}

		reportMeta := models.ReportMetadata{
			Transfer:                transfer,
			LiquidityManagerAddress: lmAddress,
			NetworkID:               transfer.From,
		}

		b, err := json.Marshal(reportMeta)
		if err != nil {
			return nil, fmt.Errorf("encode report meta: %w", err)
		}

		reports = append(reports, ocr3types.ReportWithInfo[models.ReportMetadata]{
			Report: b,
			Info:   reportMeta,
		})
	}

	return reports, nil
}

func (p *Plugin) ShouldAcceptAttestedReport(ctx context.Context, u uint64, r ocr3types.ReportWithInfo[models.ReportMetadata]) (bool, error) {
	return true, nil
}

func (p *Plugin) ShouldTransmitAcceptedReport(ctx context.Context, u uint64, r ocr3types.ReportWithInfo[models.ReportMetadata]) (bool, error) {
	p.pendingTransfers = append(p.pendingTransfers, models.PendingTransfer{
		Transfer: r.Info.Transfer,
		Status:   models.TransferStatusNotReady,
	})
	return true, nil
}

func (p *Plugin) Close() error {
	for networkID, lmAddr := range p.liquidityManagers {
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
