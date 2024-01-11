package liquiditymanager

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/shared/generated/graph_ocr3"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/liquiditygraph"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

type EvmLiquidityManager struct {
	client      graph_ocr3.GraphOCR3Interface
	lp          logpoller.LogPoller
	lmAbi       abi.ABI
	addr        common.Address
	ec          client.Client
	cleanupFunc func() error
	networkID   models.NetworkID
}

func NewEvmLiquidityManager(
	address models.Address,
	networkID models.NetworkID,
	ec client.Client,
	lp logpoller.LogPoller,
) (*EvmLiquidityManager, error) {
	lmClient, err := graph_ocr3.NewGraphOCR3(common.Address(address), ec)
	if err != nil {
		return nil, fmt.Errorf("init liquidity manager: %w", err)
	}

	lmAbi, err := abi.JSON(strings.NewReader(graph_ocr3.GraphOCR3ABI))
	if err != nil {
		return nil, fmt.Errorf("new lm abi: %w", err)
	}

	// lpFilter := logpoller.Filter{
	// 	Name: fmt.Sprintf("lm-liquidity-transferred-%s", address),
	// 	EventSigs: []common.Hash{
	// 		lmAbi.Events["LiquidityTransferred"].ID,
	// 	},
	// 	Addresses: []common.Address{common.Address(address)},
	// }

	// if err := lp.RegisterFilter(lpFilter); err != nil {
	// 	return nil, fmt.Errorf("register filter: %w", err)
	// }

	return &EvmLiquidityManager{
		client: lmClient,
		lp:     lp,
		lmAbi:  lmAbi,
		ec:     ec,
		addr:   common.Address(address),
		cleanupFunc: func() error {
			// return lp.UnregisterFilter(lpFilter.Name)
			return nil
		},
		networkID: networkID,
	}, nil
}

func (e EvmLiquidityManager) MoveLiquidity(ctx context.Context, chainID models.NetworkID, amount *big.Int) error {
	return nil
}

func (e EvmLiquidityManager) GetLiquidityManagers(ctx context.Context) (map[models.NetworkID]models.Address, error) {
	neighbors, err := e.client.GetNeighbors(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("get neighbors: %w", err)
	}
	neighborsMap := make(map[models.NetworkID]models.Address)
	for _, neighbor := range neighbors {
		neighborsMap[models.NetworkID(neighbor.ChainId.Int64())] = models.Address(neighbor.ContractAddress)
	}
	return neighborsMap, nil
}

func (e EvmLiquidityManager) GetBalance(ctx context.Context) (*big.Int, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err != nil {
		return nil, fmt.Errorf("failed to generate random number: %w", err)
	}
	return r, nil
}

func (e EvmLiquidityManager) GetPendingTransfers(ctx context.Context) ([]models.PendingTransfer, error) {
	return nil, nil
}

func (e EvmLiquidityManager) Discover(ctx context.Context, lmFactory Factory) (*Registry, liquiditygraph.LiquidityGraph, error) {
	g := liquiditygraph.NewGraph()
	lms := NewRegistry()

	type qItem struct {
		networkID models.NetworkID
		lmAddress models.Address
	}

	seen := mapset.NewSet[qItem]()
	queue := mapset.NewSet[qItem]()

	elem := qItem{networkID: e.networkID, lmAddress: models.Address(e.client.Address())}
	queue.Add(elem)
	seen.Add(elem)

	for queue.Cardinality() > 0 {
		elem, ok := queue.Pop()
		if !ok {
			return nil, nil, fmt.Errorf("unexpected internal error, there is a bug in the algorithm")
		}

		// TODO: investigate fetching the balance here.
		g.AddNetwork(elem.networkID, big.NewInt(0))

		lm, err := lmFactory.NewLiquidityManager(elem.networkID, elem.lmAddress)
		if err != nil {
			return nil, nil, fmt.Errorf("init liquidity manager: %w", err)
		}

		lms.Add(elem.networkID, elem.lmAddress)

		destinationLMs, err := lm.GetLiquidityManagers(ctx)
		if err != nil {
			return nil, nil, fmt.Errorf("get %v destination liquidity managers: %w", elem.networkID, err)
		}

		if destinationLMs == nil {
			continue
		}

		for destNetworkID, lmAddr := range destinationLMs {
			g.AddConnection(elem.networkID, destNetworkID)

			newElem := qItem{networkID: destNetworkID, lmAddress: lmAddr}
			if !seen.Contains(newElem) {
				queue.Add(newElem)
				seen.Add(newElem)

				if _, exists := lms.Get(destNetworkID); !exists {
					lms.Add(destNetworkID, lmAddr)
				}
			}
		}
	}

	return lms, g, nil
}

func (e EvmLiquidityManager) Close(ctx context.Context) error {
	return nil
}
