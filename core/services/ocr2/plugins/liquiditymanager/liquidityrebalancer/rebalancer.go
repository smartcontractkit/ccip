package liquidityrebalancer

import (
	"math/big"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/liquiditymanager/graph"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/liquiditymanager/models"
)

// UnexecutedTransfer represents a transfer of liquidity from one network to another.
// The rebalancing algorithms will all use this interface to fetch information
// about transfers that are unexecuted.
type UnexecutedTransfer interface {
	FromNetwork() models.NetworkSelector
	ToNetwork() models.NetworkSelector
	TransferAmount() *big.Int
	TransferStatus() models.TransferStatus
}

//go:generate mockery --quiet --name Rebalancer --output ../mocks --filename rebalancer_mock.go --case=underscore
type Rebalancer interface {
	// ComputeTransfersToBalance computes the transfers needed to balance the
	// liquidity across the provided graph. The rebalancer will also take into account
	// currently unexecuted transfers to avoid proposing transfers that would be
	// redundant.
	ComputeTransfersToBalance(
		g graph.Graph,
		unexecuted []UnexecutedTransfer,
	) ([]models.ProposedTransfer, error)
}
