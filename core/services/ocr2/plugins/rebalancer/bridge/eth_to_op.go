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
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/op_l2_standard_bridge"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

const (
	ERC20BridgeFinalized = "ERC20BridgeFinalized"
)

var (
	OptimismL2Abi                abi.ABI
	ERC20BridgeFinalizedEventSig common.Hash
	OptimismL2Address            = common.HexToAddress("0x4200000000000000000000000000000000000010")
)

func init() {
	abiL2, err := abi.JSON(strings.NewReader(op_l2_standard_bridge.OpL2StandardBridgeABI))
	if err != nil {
		panic(fmt.Errorf("parse op l2 bridge abi: %w", err))
	}
	OptimismL2Abi = abiL2

	if _, exists := OptimismL2Abi.Events[ERC20BridgeFinalized]; !exists {
		panic(fmt.Errorf("op l2 event %s not found in abi", ERC20BridgeFinalized))
	}
	ERC20BridgeFinalizedEventSig = OptimismL2Abi.Events[ERC20BridgeFinalized].ID
}

// EthereumToOptimism utilizes the optimism standard bridge events to implement the Bridge interface.
// https://docs.optimism.io/builders/dapp-developers/bridging/standard-bridge
type EthereumToOptimism struct {
	// lp is a logpoller instance running on Optimism
	lp logpoller.LogPoller
	// bridgeAddr is the L2 optimism standard bridge address
	bridgeAddr common.Address

	cleanupFunc func(ctx context.Context) error
}

func NewEthereumToOptimism(opLP logpoller.LogPoller, opAddr common.Address) (*EthereumToOptimism, error) {
	bridgeFinalizedFilter := logpoller.Filter{
		Name:      logpoller.FilterName("optimism-l2-%s-%s", ERC20BridgeFinalized, opAddr),
		EventSigs: []common.Hash{ERC20BridgeFinalizedEventSig},
		Addresses: []common.Address{opAddr},
	}
	if err := opLP.RegisterFilter(bridgeFinalizedFilter); err != nil {
		return nil, fmt.Errorf("register filter: %w", err)
	}

	return &EthereumToOptimism{
		lp:         opLP,
		bridgeAddr: opAddr,

		cleanupFunc: func(ctx context.Context) error {
			return opLP.UnregisterFilter(bridgeFinalizedFilter.Name, pg.WithParentCtx(ctx))
		},
	}, nil
}

// PopulateStatusOfTransfers reads all the ERC20BridgeFinalized events that occurred
// on the destination chain after the transfers were initiated.
// If events are found that match some specific transfer
// then this transfer is considered done. The funds should've been transferred to the recipient.
func (e *EthereumToOptimism) PopulateStatusOfTransfers(
	ctx context.Context, token, sender models.Address, transfers []models.Transfer) ([]models.PendingTransfer, error) {

	dateOfFirstTransfer := time.Now()
	for _, tr := range transfers {
		if tr.Date.Before(dateOfFirstTransfer) {
			dateOfFirstTransfer = tr.Date
		}
	}

	logs, err := e.lp.LogsCreatedAfter(
		ERC20BridgeFinalizedEventSig, e.bridgeAddr, dateOfFirstTransfer, logpoller.Finalized, pg.WithParentCtx(ctx))
	if err != nil {
		return nil, fmt.Errorf("get lp logs: %w", err)
	}

	opL2Filterer, err := op_l2_standard_bridge.NewOpL2StandardBridgeFilterer(e.bridgeAddr, nil)
	if err != nil {
		return nil, fmt.Errorf("new optimism l2 filterer: %w", err)
	}

	events, err := parseLogs[op_l2_standard_bridge.OpL2StandardBridgeERC20BridgeFinalized](logs,
		func(log types.Log) (*op_l2_standard_bridge.OpL2StandardBridgeERC20BridgeFinalized, error) {
			return opL2Filterer.ParseERC20BridgeFinalized(log)
		})
	if err != nil {
		return nil, fmt.Errorf("parse logs: %w", err)
	}

	transfersWithStatus := make([]models.PendingTransfer, 0, len(transfers))
	for _, tr := range transfers {
		transferWithStatus := models.NewPendingTransfer(tr)

		// todo: after nonce is added to LiqTransferred event, add a check for extra data
		// expExtraData := abi.Encode(tr.Nonce)
		// expExtraData == ev.Data.ExtraData

		// todo: check if from == lm addr

		for _, ev := range events {
			if ev.Data.From == common.Address(sender) &&
				ev.Data.RemoteToken == common.Address(token) {
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
