package liquiditygraph

import (
	"math/big"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/cciprebalance/models"
)

// LiquidityGraph contains graphs functionality that is used by the service.
type LiquidityGraph interface {
	Reset()
	GetNodes() []models.NetworkID
	AddNode(n models.NetworkID, v *big.Int)
	SetWeight(n models.NetworkID, v *big.Int)
	AddEdge(from, to models.NetworkID)
	ComputeTransfersToBalance(inflightTransfers []models.Transfer) ([]models.Transfer, error)
}
