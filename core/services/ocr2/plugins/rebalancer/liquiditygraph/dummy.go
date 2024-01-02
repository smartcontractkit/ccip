package liquiditygraph

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

type Graph struct {
	networksGraph  map[models.NetworkID][]models.NetworkID
	networkBalance map[models.NetworkID]*big.Int
}

func NewGraph() *Graph {
	return &Graph{
		networksGraph:  map[models.NetworkID][]models.NetworkID{},
		networkBalance: map[models.NetworkID]*big.Int{},
	}
}

func (g *Graph) GetNodes() []models.NetworkID {
	networks := make([]models.NetworkID, 0, len(g.networkBalance))
	for networkID := range g.networkBalance {
		networks = append(networks, networkID)
	}
	return networks
}

func (g *Graph) Reset() {
	g.networksGraph = make(map[models.NetworkID][]models.NetworkID)
	g.networkBalance = make(map[models.NetworkID]*big.Int)
}

func (g *Graph) AddNode(n models.NetworkID, v *big.Int) {
	g.networkBalance[n] = v
	g.networksGraph[n] = make([]models.NetworkID, 0)
}

func (g *Graph) SetWeight(n models.NetworkID, v *big.Int) {
	g.networkBalance[n] = v
}

func (g *Graph) AddEdge(from, to models.NetworkID) {
	g.networksGraph[from] = append(g.networksGraph[from], to)
}

func (g *Graph) IsEmpty() bool {
	return len(g.networkBalance) == 0
}

func (g *Graph) GetWeight(n models.NetworkID) (*big.Int, error) {
	w, exists := g.networkBalance[n]
	if !exists {
		return nil, fmt.Errorf("network balance not found")
	}
	return w, nil
}

func (g *Graph) String() string {
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
