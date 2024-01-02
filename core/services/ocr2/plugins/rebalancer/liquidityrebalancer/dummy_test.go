package liquidityrebalancer

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/liquiditygraph"
)

func TestDummyRebalancerComputeTransfersToBalance(t *testing.T) {
	g := liquiditygraph.NewGraph()
	g.AddNode(10, big.NewInt(1000))
	g.AddNode(20, big.NewInt(500))
	g.AddNode(30, big.NewInt(200))
	g.AddNode(40, big.NewInt(300))

	r := NewDummyRebalancer()
	transfers, err := r.ComputeTransfersToBalance(g, nil)
	assert.NoError(t, err)
	assert.Len(t, transfers, 3)
}
