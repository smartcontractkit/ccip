package evmliquiditymanager

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/liquiditymanager/generated/liquiditymanager"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/liquiditymanager/models"
)

var _ LiquidityManager = &EvmLiquidityManager{}

type EvmLiquidityManager struct {
	rebalancer liquiditymanager.LiquidityManagerInterface
	addr       common.Address
	networkSel models.NetworkSelector
	lggr       logger.Logger
}

func NewEvmLiquidityManager(
	address models.Address,
	net models.NetworkSelector,
	ec client.Client,
	lggr logger.Logger,
) (*EvmLiquidityManager, error) {
	rebal, err := liquiditymanager.NewLiquidityManager(common.Address(address), ec)
	if err != nil {
		return nil, fmt.Errorf("unable to instantiate liquidity manager wrapper: %w", err)
	}

	return &EvmLiquidityManager{
		rebalancer: rebal,
		addr:       common.Address(address),
		networkSel: net,
		lggr:       lggr.Named("EvmRebalancer"),
	}, nil
}

func (e *EvmLiquidityManager) GetRebalancers(ctx context.Context) (map[models.NetworkSelector]models.Address, error) {
	lms, err := e.rebalancer.GetAllCrossChainRebalancers(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("get all cross chain rebalancers: %w", err)
	}
	ret := make(map[models.NetworkSelector]models.Address)
	for _, lm := range lms {
		ret[models.NetworkSelector(lm.RemoteChainSelector)] = models.Address(lm.RemoteRebalancer)
	}
	return ret, nil
}

func (e *EvmLiquidityManager) GetBalance(ctx context.Context) (*big.Int, error) {
	return e.rebalancer.GetLiquidity(&bind.CallOpts{Context: ctx})
}

func (e *EvmLiquidityManager) Close(ctx context.Context) error {
	return nil
}

// ConfigDigest implements Rebalancer.
func (e *EvmLiquidityManager) ConfigDigest(ctx context.Context) (types.ConfigDigest, error) {
	cdae, err := e.rebalancer.LatestConfigDigestAndEpoch(&bind.CallOpts{Context: ctx})
	if err != nil {
		return ocrtypes.ConfigDigest{}, fmt.Errorf("latest config digest and epoch: %w", err)
	}
	return ocrtypes.ConfigDigest(cdae.ConfigDigest), nil
}

func (e *EvmLiquidityManager) GetTokenAddress(ctx context.Context) (models.Address, error) {
	tokenAddress, err := e.rebalancer.ILocalToken(&bind.CallOpts{
		Context: ctx,
	})
	return models.Address(tokenAddress), err
}

func (e *EvmLiquidityManager) GetLatestSequenceNumber(ctx context.Context) (uint64, error) {
	cdae, err := e.rebalancer.LatestConfigDigestAndEpoch(&bind.CallOpts{Context: ctx})
	return cdae.SequenceNumber, err
}
