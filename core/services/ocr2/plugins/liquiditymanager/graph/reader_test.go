package graph

import (
	"math/big"
	"testing"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/liquiditymanager/models"
)

func TestLiquidityGraph_FindPath(t *testing.T) {
	tests := []struct {
		name  string
		setup func(graph *liquidityGraph)
		from  models.NetworkSelector
		to    models.NetworkSelector
		hops  int
		check func(data ...Data) bool
		want  []models.NetworkSelector
	}{
		{
			name:  "no path found case",
			setup: func(graph *liquidityGraph) {},
			from:  1,
			to:    2,
			hops:  5,
			check: func(data ...Data) bool { return true },
			want:  []models.NetworkSelector{},
		},
		{
			name: "path found case",
			setup: func(graph *liquidityGraph) {
				d1 := Data{NetworkSelector: 1, Liquidity: big.NewInt(100)}
				d2 := Data{NetworkSelector: 2, Liquidity: big.NewInt(100)}
				_ = graph.addNetwork(d1.NetworkSelector, d1)
				_ = graph.addNetwork(d2.NetworkSelector, d2)
				_ = graph.addConnection(d1.NetworkSelector, d2.NetworkSelector)
			},
			from:  1,
			to:    2,
			hops:  5,
			check: func(data ...Data) bool { return true },
			want:  []models.NetworkSelector{2},
		},
		{
			name: "path found case - spoke",
			setup: func(graph *liquidityGraph) {
				eth := Data{NetworkSelector: 1, Liquidity: big.NewInt(100)}
				opt := Data{NetworkSelector: 2, Liquidity: big.NewInt(100)}
				arb := Data{NetworkSelector: 3, Liquidity: big.NewInt(100)}
				base := Data{NetworkSelector: 4, Liquidity: big.NewInt(100)}
				celo := Data{NetworkSelector: 5, Liquidity: big.NewInt(100)}
				_ = graph.addNetwork(eth.NetworkSelector, eth)
				_ = graph.addNetwork(opt.NetworkSelector, opt)
				_ = graph.addNetwork(arb.NetworkSelector, arb)
				_ = graph.addNetwork(base.NetworkSelector, base)
				_ = graph.addNetwork(celo.NetworkSelector, celo)
				_ = graph.addConnection(eth.NetworkSelector, opt.NetworkSelector)
				_ = graph.addConnection(opt.NetworkSelector, eth.NetworkSelector)
				_ = graph.addConnection(eth.NetworkSelector, arb.NetworkSelector)
				_ = graph.addConnection(arb.NetworkSelector, eth.NetworkSelector)
				_ = graph.addConnection(arb.NetworkSelector, base.NetworkSelector)
				_ = graph.addConnection(base.NetworkSelector, arb.NetworkSelector)
				_ = graph.addConnection(base.NetworkSelector, celo.NetworkSelector)
				_ = graph.addConnection(celo.NetworkSelector, base.NetworkSelector)
			},
			from:  2,
			to:    5,
			hops:  4,
			check: func(data ...Data) bool { return true },
			want:  []models.NetworkSelector{models.NetworkSelector(1), models.NetworkSelector(3)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &liquidityGraph{
				adj:  make(map[models.NetworkSelector][]models.NetworkSelector),
				data: make(map[models.NetworkSelector]Data),
			}
			tt.setup(g)
			if got := g.FindPath(tt.from, tt.to, tt.hops, tt.check); !equalPath(got, tt.want) {
				t.Errorf("liquidityGraph.findPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Checks whether two paths are equal. Here path is defined as slice of models.NetworkSelector
func equalPath(path1, path2 []models.NetworkSelector) bool {
	if len(path1) != len(path2) {
		return false
	}

	for i := range path1 {
		if path1[i] != path2[i] {
			return false
		}
	}

	return true
}
