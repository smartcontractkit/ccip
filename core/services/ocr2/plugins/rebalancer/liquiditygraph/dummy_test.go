package liquiditygraph

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDummyGraph(t *testing.T) {
	g := NewDummyGraph()
	g.AddNode(10, big.NewInt(1000))
	g.AddNode(20, big.NewInt(500))
	g.AddNode(30, big.NewInt(200))
	g.AddNode(40, big.NewInt(300))
	transfers, err := g.ComputeTransfersToBalance(nil)
	assert.NoError(t, err)
	assert.Len(t, transfers, 3)
}
