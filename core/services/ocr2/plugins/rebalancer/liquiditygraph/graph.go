package liquiditygraph

import (
	"math/big"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

// LiquidityGraph contains graphs functionality that is used by the service.
// Graph operations should be thread-safe.
type LiquidityGraph interface {
	Reset()
	GetNodes() []models.NetworkID
	AddNode(n models.NetworkID, v *big.Int)
	SetWeight(n models.NetworkID, v *big.Int)
	GetWeight(n models.NetworkID) (*big.Int, error)
	AddEdge(from, to models.NetworkID)
	IsEmpty() bool
}
