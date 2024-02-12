package liquiditymanager

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/rebalancer"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

var _ Rebalancer = &EvmRebalancer{}

type EvmRebalancer struct {
	client      OnchainRebalancer
	lp          logpoller.LogPoller
	lmAbi       abi.ABI
	addr        common.Address
	networkSel  models.NetworkSelector
	ec          client.Client
	cleanupFunc func() error
	lggr        logger.Logger
}

func NewEvmRebalancer(
	address models.Address,
	net models.NetworkSelector,
	ec client.Client,
	lp logpoller.LogPoller,
	lggr logger.Logger) (*EvmRebalancer, error) {
	client, err := NewConcreteRebalancer(common.Address(address), ec)
	if err != nil {
		return nil, fmt.Errorf("new concrete rebalancer: %w", err)
	}

	lmAbi, err := abi.JSON(strings.NewReader(rebalancer.RebalancerABI))
	if err != nil {
		return nil, fmt.Errorf("new rebalancer abi: %w", err)
	}

	lpFilter := logpoller.Filter{
		Name: fmt.Sprintf("%d-lm-liquidity-transferred-%s", time.Now().UnixNano(), common.Address(address)),
		EventSigs: []common.Hash{
			lmAbi.Events["LiquidityTransferred"].ID,
		},
		Addresses: []common.Address{common.Address(address)},
	}

	if err := lp.RegisterFilter(lpFilter); err != nil {
		return nil, fmt.Errorf("register filter: %w", err)
	}

	return &EvmRebalancer{
		client:     client,
		lp:         lp,
		lmAbi:      lmAbi,
		ec:         ec,
		addr:       common.Address(address),
		networkSel: net,
		cleanupFunc: func() error {
			return lp.UnregisterFilter(lpFilter.Name)
		},
		lggr: lggr.Named("EvmRebalancer"),
	}, nil
}

func (e *EvmRebalancer) GetRebalancers(ctx context.Context) (map[models.NetworkSelector]models.Address, error) {
	return e.client.GetAllCrossChainRebalancers(ctx)
}

func (e *EvmRebalancer) GetBalance(ctx context.Context) (*big.Int, error) {
	return e.client.GetLiquidity(ctx)
}

func (e *EvmRebalancer) GetPendingTransfers(ctx context.Context, since time.Time) ([]models.PendingTransfer, error) {
	logs, err := e.lp.LogsCreatedAfter(
		e.lmAbi.Events["LiquidityTransferred"].ID,
		e.addr,
		since,
		logpoller.Finalized,
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("get logs created after: %w", err)
	}

	pendingTransfers := make([]models.PendingTransfer, 0, len(logs))

	for _, log := range logs {
		liqTransferred, err2 := e.client.ParseLiquidityTransferred(log.ToGethLog())
		if err2 != nil {
			return nil, fmt.Errorf("invalid log: %w", err2)
		}

		tr := models.NewPendingTransfer(models.NewTransfer(
			models.NetworkSelector(liqTransferred.FromChainSelector()),
			models.NetworkSelector(liqTransferred.ToChainSelector()),
			liqTransferred.Amount(),
			log.BlockTimestamp,
			[]byte{}, // TODO: fill in bridge data
		))
		// tr.Status = models.TransferStatusExecuted // todo: determine the status
		pendingTransfers = append(pendingTransfers, tr)
	}

	return pendingTransfers, nil
}

func (e *EvmRebalancer) Close(ctx context.Context) error {
	return e.cleanupFunc()
}

// ConfigDigest implements Rebalancer.
func (e *EvmRebalancer) ConfigDigest(ctx context.Context) (types.ConfigDigest, error) {
	return e.client.GetConfigDigest(ctx)
}

func (e *EvmRebalancer) GetTokenAddress(ctx context.Context) (models.Address, error) {
	return e.client.GetTokenAddress(ctx)
}
