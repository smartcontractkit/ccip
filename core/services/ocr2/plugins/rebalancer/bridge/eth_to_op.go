package bridge

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/optimism_l2"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

const (
	ERC20BridgeFinalized = "ERC20BridgeFinalized"
)

type EthereumToOptimism struct {
	optimismLP   logpoller.LogPoller
	optimismAddr common.Address

	opL2ABI     abi.ABI
	cleanupFunc func(ctx context.Context) error
}

func NewEthereumToOptimism(opLP logpoller.LogPoller, opAddr common.Address) (*EthereumToOptimism, error) {
	abiL2, err := abi.JSON(strings.NewReader(optimism_l2.OptimismL2ABI))
	if err != nil {
		return nil, fmt.Errorf("parse op l2 bridge abi: %w", err)
	}

	bridgeFinalizedFilter := logpoller.Filter{
		Name:      logpoller.FilterName("optimism-erc20-bridge-finalized-%s", opAddr),
		EventSigs: []common.Hash{abiL2.Events[ERC20BridgeFinalized].ID},
		Addresses: []common.Address{opAddr},
		Retention: 30 * 24 * time.Hour,
	}

	if err := opLP.RegisterFilter(bridgeFinalizedFilter); err != nil {
		return nil, fmt.Errorf("register filter: %w", err)
	}

	return &EthereumToOptimism{
		optimismLP:   opLP,
		optimismAddr: opAddr,

		opL2ABI: abiL2,
		cleanupFunc: func(ctx context.Context) error {
			return opLP.UnregisterFilter(bridgeFinalizedFilter.Name, pg.WithParentCtx(ctx))
		},
	}, nil
}

// PopulateStatusOfTransfers reads all the ERC20BridgeFinalized events that occurred
// after the transfers started. If events are found that match some specific transfer
// then this transfer is considered done. The funds should've been in the recipient's wallet.
func (e *EthereumToOptimism) PopulateStatusOfTransfers(ctx context.Context, token, sender common.Address, transfers []models.Transfer) ([]models.PendingTransfer, error) {
	eventSig := e.opL2ABI.Events[ERC20BridgeFinalized].ID

	createdAfter := time.Now()
	for _, tr := range transfers {
		if tr.Date.Before(createdAfter) {
			createdAfter = tr.Date
		}
	}

	logs, err := e.optimismLP.LogsCreatedAfter(
		eventSig, e.optimismAddr, createdAfter, logpoller.Finalized, pg.WithParentCtx(ctx))
	if err != nil {
		return nil, fmt.Errorf("get lp logs: %w", err)
	}

	events, err := parseLogs[optimism_l2.OptimismL2ERC20BridgeFinalized](logs,
		func(log types.Log) (*optimism_l2.OptimismL2ERC20BridgeFinalized, error) {
			opL2, err := optimism_l2.NewOptimismL2Filterer(e.optimismAddr, nil)
			if err != nil {
				return nil, err
			}
			return opL2.ParseERC20BridgeFinalized(log)
		})
	if err != nil {
		return nil, fmt.Errorf("parse logs: %w", err)
	}

	transfersWithStatus := make([]models.PendingTransfer, 0, len(transfers))

	for _, tr := range transfers {
		transferWithStatus := models.NewPendingTransfer(tr)

		expExtraData, err := e.opL2ABI.Pack("", tr.ID) // todo: test
		if err != nil {
			return nil, fmt.Errorf("pack transfer id: %w", err)
		}

		for _, ev := range events {
			// todo: check if from == lm addr
			if string(expExtraData) == string(ev.Data.ExtraData) &&
				ev.Data.From == sender &&
				ev.Data.RemoteToken == token {
				transferWithStatus.Status = models.TransferStatusExecuted
			}
		}

		transfersWithStatus = append(transfersWithStatus, transferWithStatus)
	}

	return transfersWithStatus, nil
}

func (e *EthereumToOptimism) Close(ctx context.Context) error {
	return e.cleanupFunc(ctx)
}

// todo: copied from ccip, move and re-use from a lib
func parseLogs[T any](logs []logpoller.Log, parseFunc func(log types.Log) (*T, error)) ([]Event[T], error) {
	parsed := make([]Event[T], 0, len(logs))
	for _, log := range logs {
		data, err := parseFunc(log.ToGethLog())
		if err != nil {
			return nil, fmt.Errorf("cannot parse log: %w", err)
		}
		parsed = append(parsed, Event[T]{
			Data: *data,
			Meta: Meta{
				BlockTimestamp: log.BlockTimestamp,
				BlockNumber:    log.BlockNumber,
				TxHash:         log.TxHash,
				LogIndex:       uint(log.LogIndex),
			},
		})
	}
	return parsed, nil
}

type Event[T any] struct {
	Data T
	Meta
}

type Meta struct {
	BlockTimestamp time.Time
	BlockNumber    int64
	TxHash         common.Hash
	LogIndex       uint
}
