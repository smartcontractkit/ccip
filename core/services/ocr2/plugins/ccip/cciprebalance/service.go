package cciprebalance

import (
	"context"
	"fmt"
	"math/big"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/cciprebalance/liquiditygraph"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/cciprebalance/liquiditymanager"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/cciprebalance/models"
)

type Service struct {
	assetAddress     models.Address
	lmFactory        liquiditymanager.Factory
	lms              map[models.NetworkID]models.Address
	pendingTransfers []models.Transfer
	graph            liquiditygraph.LiquidityGraph
}

func NewService(
	assetAddress models.Address,
	lmFactory liquiditymanager.Factory,
	lms map[models.NetworkID]models.Address,
	liquidityGraph liquiditygraph.LiquidityGraph,
) *Service {
	return &Service{
		assetAddress:     assetAddress,
		lmFactory:        lmFactory,
		lms:              lms,
		pendingTransfers: make([]models.Transfer, 0),
		graph:            liquidityGraph,
	}
}

func (s *Service) Run(ctx context.Context) error {
	if err := s.syncGraphEdges(ctx); err != nil {
		return fmt.Errorf("sync graph edges: %w", err)
	}

	if err := s.syncGraphBalances(ctx); err != nil {
		return fmt.Errorf("sync graph balances: %w", err)
	}

	if err := s.loadPendingTransfers(ctx); err != nil {
		return fmt.Errorf("load pending transfers: %w", err)
	}

	if err := s.balanceLiquidity(ctx); err != nil {
		return fmt.Errorf("balance liquidity: %w", err)
	}

	return nil
}

func (s *Service) syncGraphEdges(ctx context.Context) error {
	// todo: if there wasn't any change to the graph stop earlier

	s.graph.Reset()

	type qItem struct {
		networkID models.NetworkID
		lmAddress models.Address
	}

	seen := mapset.NewSet[qItem]()
	queue := mapset.NewSet[qItem]()
	for networkID, lmAddress := range s.lms {
		elem := qItem{networkID: networkID, lmAddress: lmAddress}
		queue.Add(elem)
		seen.Add(elem)
	}

	for queue.Cardinality() > 0 {
		elem, ok := queue.Pop()
		if !ok {
			return errors.New("unexpected internal error, there is a bug in the algorithm")
		}

		s.graph.AddNode(elem.networkID, big.NewInt(0))

		lm, err := s.lmFactory.NewLiquidityManager(elem.networkID, elem.lmAddress) // todo: this is not testable
		if err != nil {
			return fmt.Errorf("init liquidity manager: %w", err)
		}

		destinationLMs, err := lm.GetLiquidityManagers(ctx)
		if err != nil {
			return fmt.Errorf("get %v destination liquidity managers: %w", elem.networkID, err)
		}

		for destNetworkID, lmAddr := range destinationLMs {
			s.graph.AddEdge(elem.networkID, destNetworkID)

			newElem := qItem{networkID: destNetworkID, lmAddress: lmAddr}
			if !seen.Contains(newElem) {
				queue.Add(newElem)
				seen.Add(newElem)

				if _, exists := s.lms[destNetworkID]; !exists {
					s.lms[destNetworkID] = lmAddr
				}
			}
		}
	}

	return nil
}

func (s *Service) syncGraphBalances(ctx context.Context) error {
	for _, networkID := range s.graph.GetNodes() {
		lmAddr, exists := s.lms[networkID]
		if !exists {
			return fmt.Errorf("liquidity manager for network %v was not found", networkID)
		}

		lm, err := s.lmFactory.NewLiquidityManager(networkID, lmAddr)
		if err != nil {
			return fmt.Errorf("init liquidity manager: %w", err)
		}

		balance, err := lm.GetBalance(ctx)
		if err != nil {
			return fmt.Errorf("get %v balance: %w", networkID, err)
		}

		s.graph.SetWeight(networkID, balance)
	}

	return nil
}

func (s *Service) loadPendingTransfers(ctx context.Context) error {
	// todo: do not load pending transfers all the time

	s.pendingTransfers = make([]models.Transfer, 0)

	for networkID, lmAddress := range s.lms {
		lm, err := s.lmFactory.NewLiquidityManager(networkID, lmAddress)
		if err != nil {
			return fmt.Errorf("init liquidity manager: %w", err)
		}

		pendingTransfers, err := lm.GetPendingTransfers(ctx)
		if err != nil {
			return fmt.Errorf("get pending %v transfers: %w", networkID, err)
		}

		s.pendingTransfers = append(s.pendingTransfers, pendingTransfers...)
	}

	return nil
}

func (s *Service) balanceLiquidity(ctx context.Context) error {
	transfersToBalance, err := s.graph.ComputeTransfersToBalance(s.pendingTransfers)
	if err != nil {
		return fmt.Errorf("compute transfers to balance: %w", err)
	}

	for _, transfer := range transfersToBalance {
		lmAddress, exists := s.lms[transfer.From]
		if !exists {
			return fmt.Errorf("liquidity manager for %v does not exist", transfer.From)
		}

		lm, err := s.lmFactory.NewLiquidityManager(transfer.From, lmAddress)
		if err != nil {
			return fmt.Errorf("init liquidity manager: %w", err)
		}

		if err := lm.MoveLiquidity(ctx, transfer.To, transfer.Amount); err != nil {
			return fmt.Errorf("move liquidity %v: %w", transfer, err)
		}
		s.pendingTransfers = append(s.pendingTransfers, transfer)
	}

	return nil
}
