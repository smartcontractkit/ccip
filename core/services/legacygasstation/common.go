package legacygasstation

import (
	"context"

	"github.com/ethereum/go-ethereum"

	txmgrtypes "github.com/smartcontractkit/chainlink/v2/common/txmgr/types"
	"github.com/smartcontractkit/chainlink/v2/core/assets"
	"github.com/smartcontractkit/chainlink/v2/core/services/legacygasstation/types"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

type Config interface {
	EvmGasLimitDefault() uint32
	EvmFinalityDepth() uint32
	EvmMaxGasPriceWei() *assets.Wei
}

type ORM interface {
	SelectBySourceChainIDAndStatus(sourceChainID uint64, status types.Status, qopts ...pg.QOpt) (txs []types.LegacyGaslessTx, err error)
	SelectByDestChainIDAndStatus(destChainID uint64, status types.Status, qopts ...pg.QOpt) (txs []types.LegacyGaslessTx, err error)
	SelectBySourceChainIDAndEthTxStates(sourceChainID uint64, states []txmgrtypes.TxState, qopts ...pg.QOpt) ([]types.LegacyGaslessTxPlus, error)
	InsertLegacyGaslessTx(tx types.LegacyGaslessTx, qopts ...pg.QOpt) error
	UpdateLegacyGaslessTx(tx types.LegacyGaslessTx, qopts ...pg.QOpt) error
}

type EthClient interface {
	EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error)
}
