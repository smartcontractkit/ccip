package bridge

import (
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

type EthereumToOptimism struct{}

func NewEthereumToOptimism() *EthereumToOptimism {
	return &EthereumToOptimism{}
}

func (e *EthereumToOptimism) PopulateStatusOfTransfers(transfers []models.Transfer) ([]models.PendingTransfer, error) {
	return nil, nil // todo
}
