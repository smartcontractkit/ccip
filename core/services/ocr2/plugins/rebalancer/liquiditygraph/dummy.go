package liquiditygraph

import (
	"fmt"
	"math/big"
	"sort"
	"strings"

	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

type DummyGraph struct {
	networksGraph  map[models.NetworkID][]models.NetworkID
	networkBalance map[models.NetworkID]*big.Int
}

func NewDummyGraph() *DummyGraph {
	return &DummyGraph{
		networksGraph:  map[models.NetworkID][]models.NetworkID{},
		networkBalance: map[models.NetworkID]*big.Int{},
	}
}

func (g *DummyGraph) GetNodes() []models.NetworkID {
	networks := make([]models.NetworkID, 0, len(g.networkBalance))
	for networkID := range g.networkBalance {
		networks = append(networks, networkID)
	}
	return networks
}

func (g *DummyGraph) Reset() {
	g.networksGraph = make(map[models.NetworkID][]models.NetworkID)
	g.networkBalance = make(map[models.NetworkID]*big.Int)
}

func (g *DummyGraph) AddNode(n models.NetworkID, v *big.Int) {
	g.networkBalance[n] = v
	g.networksGraph[n] = make([]models.NetworkID, 0)
}

func (g *DummyGraph) SetWeight(n models.NetworkID, v *big.Int) {
	g.networkBalance[n] = v
}

func (g *DummyGraph) AddEdge(from, to models.NetworkID) {
	g.networksGraph[from] = append(g.networksGraph[from], to)
}

func (g *DummyGraph) ComputeTransfersToBalance(inflightTransfers []models.Transfer) ([]models.Transfer, error) {
	// selects the node with the highest balance
	// and moves all the liquidity from the other nodes to it
	// inflightTransfers are ignored

	if len(g.networkBalance) == 0 {
		return nil, fmt.Errorf("empty graph")
	}

	keys := maps.Keys(g.networkBalance)
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	luckyNode := keys[0]
	maxV := g.networkBalance[luckyNode]

	for _, k := range keys {
		if g.networkBalance[k].Cmp(maxV) > 0 {
			luckyNode = k
			maxV = g.networkBalance[k]
		}
	}

	transfers := make([]models.Transfer, 0)
	for node, w := range g.networkBalance {
		if node == luckyNode {
			continue
		}
		if w.BitLen() == 0 {
			continue
		}

		transfers = append(transfers, models.Transfer{
			From:   node,
			To:     luckyNode,
			Amount: w,
		})
	}

	return transfers, nil
}

func (g *DummyGraph) String() string {
	sb := strings.Builder{}

	sb.WriteString("~~~ NODES ~~~\n")
	for n, w := range g.networkBalance {
		sb.WriteString(fmt.Sprintf("[%d] %s\n", n, w))
	}

	sb.WriteString("\n~~~ LINKS ~~~\n")
	for n, nbs := range g.networksGraph {
		sb.WriteString(fmt.Sprintf("%d: ", n))
		for _, nb := range nbs {
			sb.WriteString(fmt.Sprintf(" %d ", nb))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
