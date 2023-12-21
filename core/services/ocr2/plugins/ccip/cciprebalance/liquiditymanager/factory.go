package liquiditymanager

import (
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/cciprebalance/models"
)

//go:generate mockery --quiet --name Factory --output ../cciprebalancemocks --filename lm_factory_mock.go --case=underscore
type Factory interface {
	NewLiquidityManager(networkID models.NetworkID, address models.Address) (LiquidityManager, error)
}

type BaseLiquidityManagerFactory struct{}

func (b BaseLiquidityManagerFactory) NewLiquidityManager(networkID models.NetworkID, address models.Address) (LiquidityManager, error) {
	switch networkID {
	case 1: // todo
		return NewEvmLiquidityManager(address), nil
	default:
		return nil, errors.New("not found")
	}
}
