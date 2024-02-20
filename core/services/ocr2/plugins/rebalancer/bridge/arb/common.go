package arb

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/rebalancer"
)

type logKey struct {
	txHash   common.Hash
	logIndex int64
}

func parseLiquidityTransferred(parseFunc func(gethtypes.Log) (*rebalancer.RebalancerLiquidityTransferred, error), lgs []logpoller.Log) ([]*rebalancer.RebalancerLiquidityTransferred, map[logKey]logpoller.Log, error) {
	transferred := make([]*rebalancer.RebalancerLiquidityTransferred, len(lgs))
	toLP := make(map[logKey]logpoller.Log)
	for i, lg := range lgs {
		parsed, err := parseFunc(lg.ToGethLog())
		if err != nil {
			// should never happen
			return nil, nil, fmt.Errorf("parse LiquidityTransferred log: %w", err)
		}
		transferred[i] = parsed
		toLP[logKey{
			txHash:   lg.TxHash,
			logIndex: lg.LogIndex,
		}] = lg
	}
	return transferred, toLP, nil
}

func unpackUint256(data []byte) (*big.Int, error) {
	ifaces, err := utils.ABIDecode(`[{"type": "uint256"}]`, data)
	if err != nil {
		return nil, fmt.Errorf("decode uint256: %w", err)
	}
	if len(ifaces) != 1 {
		return nil, fmt.Errorf("expected 1 argument, got %d", len(ifaces))
	}
	ret := *abi.ConvertType(ifaces[0], new(*big.Int)).(**big.Int)
	return ret, nil
}
