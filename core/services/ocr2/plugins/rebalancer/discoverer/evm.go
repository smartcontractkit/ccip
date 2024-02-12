package discoverer

import (
	"context"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/rebalancer"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/graph"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

type evmDiscoverer struct {
	evmClients       map[models.NetworkSelector]evmDep
	masterRebalancer models.Address
	masterSelector   models.NetworkSelector
}

func (e *evmDiscoverer) Discover(ctx context.Context) (graph.Graph, error) {
	getData := func(ctx context.Context, selector models.NetworkSelector, rebalancerAddress models.Address) (graph.Data, error) {
		dep, ok := e.evmClients[selector]
		if !ok {
			return graph.Data{}, fmt.Errorf("no client for master chain %+v", selector)
		}
		rebal, err := rebalancer.NewRebalancer(common.Address(e.masterRebalancer), dep.ethClient)
		if err != nil {
			return graph.Data{}, fmt.Errorf("new rebalancer: %w", err)
		}
		liquidity, err := rebal.GetLiquidity(&bind.CallOpts{
			Context: ctx,
		})
		if err != nil {
			return graph.Data{}, fmt.Errorf("get liquidity: %w", err)
		}
		token, err := rebal.ILocalToken(&bind.CallOpts{
			Context: ctx,
		})
		if err != nil {
			return graph.Data{}, fmt.Errorf("get token: %w", err)
		}
		return graph.Data{
			Liquidity:         liquidity,
			TokenAddress:      models.Address(token),
			RebalancerAddress: rebalancerAddress,
		}, nil
	}

	getNeighbors := func(ctx context.Context, selector models.NetworkSelector, rebalancerAddress models.Address) ([]dataItem, error) {
		dep, ok := e.evmClients[selector]
		if !ok {
			return []dataItem{}, fmt.Errorf("no client for master chain %+v", selector)
		}
		rebal, err := rebalancer.NewRebalancer(common.Address(e.masterRebalancer), dep.ethClient)
		if err != nil {
			return []dataItem{}, fmt.Errorf("new rebalancer: %w", err)
		}
		xchainRebalancers, err := rebal.GetAllCrossChainRebalancers(&bind.CallOpts{
			Context: ctx,
		})
		if err != nil {
			return []dataItem{}, fmt.Errorf("get all cross chain rebalancers: %w", err)
		}
		var neighbors []dataItem
		for _, v := range xchainRebalancers {
			neighbors = append(neighbors, dataItem{
				networkSelector:   models.NetworkSelector(v.RemoteChainSelector),
				rebalancerAddress: models.Address(v.RemoteRebalancer),
			})
		}
		return neighbors, nil
	}

	return discover(ctx, e.masterSelector, e.masterRebalancer, getNeighbors, getData)
}

type dataItem struct {
	networkSelector   models.NetworkSelector
	rebalancerAddress models.Address
}

func discover(
	ctx context.Context,
	startNetwork models.NetworkSelector,
	startAddress models.Address,
	getNeighbors func(ctx context.Context, network models.NetworkSelector, rebalancerAddress models.Address) ([]dataItem, error),
	getData func(ctx context.Context, key models.NetworkSelector, rebalancerAddress models.Address) (graph.Data, error),
) (graph.Graph, error) {
	g := graph.NewGraph()

	seen := mapset.NewSet[dataItem]()
	queue := mapset.NewSet[dataItem]()

	start := dataItem{
		networkSelector:   startNetwork,
		rebalancerAddress: startAddress,
	}
	queue.Add(start)
	seen.Add(start)
	for queue.Cardinality() > 0 {
		elem, ok := queue.Pop()
		if !ok {
			return nil, fmt.Errorf("unexpected internal error")
		}

		val, err := getData(ctx, elem.networkSelector, elem.rebalancerAddress)
		if err != nil {
			return nil, fmt.Errorf("could not get value for vertex %+v: %w", elem, err)
		}
		g.AddNetwork(elem.networkSelector, val)

		neighbors, err := getNeighbors(ctx, elem.networkSelector, elem.rebalancerAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to get neighbors on vertex %+v: %w", elem, err)
		}

		for _, neighbor := range neighbors {
			if !g.HasNetwork(neighbor.networkSelector) {
				val2, err := getData(ctx, neighbor.networkSelector, neighbor.rebalancerAddress)
				if err != nil {
					return nil, fmt.Errorf("could not get value for vertex %+v: %w", elem, err)
				}
				g.AddNetwork(neighbor.networkSelector, val2)
			}

			if err := g.AddConnection(elem.networkSelector, neighbor.networkSelector); err != nil {
				return nil, fmt.Errorf("error adding connection from %+v to %+v: %w", elem, neighbor, err)
			}

			if !seen.Contains(neighbor) {
				queue.Add(neighbor)
				seen.Add(neighbor)
			}
		}
	}

	return g, nil
}
