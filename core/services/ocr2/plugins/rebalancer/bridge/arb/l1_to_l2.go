package arb

import (
	"context"
	"fmt"
	"math/big"
	"slices"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	chainsel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	ubig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/abstract_arbitrum_token_gateway"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arb_node_interface"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_gateway_router"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_inbox"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_l1_bridge_adapter"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_token_gateway"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/l2_arbitrum_gateway"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/rebalancer"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

var (
	// Multipliers to ensure our L1 -> L2 tx goes through
	// These values match the arbitrum SDK
	// TODO: should these be configurable?
	l2BaseFeeMultiplier     = big.NewInt(3)
	submissionFeeMultiplier = big.NewInt(4)

	// Important events we're going to track
	// Since the deposit finalized topic is emitted from any L2 gateway,
	// we won't be able to specify addresses in the log poller filters.
	// These are emitted on L2
	DepositFinalizedTopic     = l2_arbitrum_gateway.L2ArbitrumGatewayDepositFinalized{}.Topic()
	LiquidityTransferredTopic = rebalancer.RebalancerLiquidityTransferred{}.Topic()

	// Events emitted on L1
	ArbitrumL1ToL2ERC20Sent = arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent{}.Topic()

	nodeInterfaceABI = abihelpers.MustParseABI(arb_node_interface.NodeInterfaceMetaData.ABI)
)

type l1ToL2Bridge struct {
	localSelector  models.NetworkSelector
	remoteSelector models.NetworkSelector

	l1Rebalancer        *rebalancer.Rebalancer
	l2RebalancerAddress common.Address
	l1BridgeAdapter     *arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapter

	// Arbitrum contract wrappers
	l1GatewayRouter *arbitrum_gateway_router.ArbitrumGatewayRouter
	l1Inbox         *arbitrum_inbox.ArbitrumInbox
	l2Gateway       *l2_arbitrum_gateway.L2ArbitrumGateway

	l1Client client.Client
	l2Client client.Client

	l1LogPoller logpoller.LogPoller
	l2LogPoller logpoller.LogPoller

	l1FilterName string
	l2FilterName string

	lggr logger.Logger
}

func NewL1ToL2Bridge(
	lggr logger.Logger,
	localSelector,
	remoteSelector models.NetworkSelector,
	l1RebalancerAddress,
	l2RebalancerAddress,
	l1BridgeAdapterAddress,
	l1GatewayRouterAddress,
	l1InboxAddress common.Address,
	l1Client,
	l2Client client.Client,
	l1LogPoller,
	l2LogPoller logpoller.LogPoller,
) (*l1ToL2Bridge, error) {
	localChain, ok := chainsel.ChainBySelector(uint64(localSelector))
	if !ok {
		return nil, fmt.Errorf("unknown chain selector for local chain: %d", localSelector)
	}
	remoteChain, ok := chainsel.ChainBySelector(uint64(remoteSelector))
	if !ok {
		return nil, fmt.Errorf("unknown chain selector for remote chain: %d", remoteSelector)
	}

	l1GatewayRouter, err := arbitrum_gateway_router.NewArbitrumGatewayRouter(l1GatewayRouterAddress, l1Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate L1 gateway router at %s: %w", l1GatewayRouterAddress, err)
	}

	l1Inbox, err := arbitrum_inbox.NewArbitrumInbox(l1InboxAddress, l1Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate L1 inbox at %s: %w", l1InboxAddress, err)
	}

	l1BridgeAdapter, err := arbitrum_l1_bridge_adapter.NewArbitrumL1BridgeAdapter(l1BridgeAdapterAddress, l1Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate L1 bridge adapter at %s: %w", l1BridgeAdapterAddress, err)
	}

	l1FilterName := fmt.Sprintf("ArbitrumL1ToL2Bridge-%s", l1BridgeAdapterAddress.String())
	err = l1LogPoller.RegisterFilter(logpoller.Filter{
		Addresses: []common.Address{l1BridgeAdapterAddress},
		Name:      l1FilterName,
		EventSigs: []common.Hash{
			ArbitrumL1ToL2ERC20Sent,
		},
		Retention: DurationMonth,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to register L1 log filter: %w", err)
	}

	// figure out which gateway to watch for the token on L2
	l1Rebalancer, err := rebalancer.NewRebalancer(l1RebalancerAddress, l1Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate rebalancer at %s: %w", l1RebalancerAddress, err)
	}

	l1Token, err := l1Rebalancer.ILocalToken(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get local token from rebalancer: %w", err)
	}

	// get the gateway on L1 and then it's counterpart gateway on L2
	// that's the one we need to watch
	l1TokenGateway, err := l1GatewayRouter.GetGateway(nil, l1Token)
	if err != nil {
		return nil, fmt.Errorf("failed to get gateway for token %s: %w", l1Token, err)
	}

	abstractGateway, err := abstract_arbitrum_token_gateway.NewAbstractArbitrumTokenGateway(l1TokenGateway, l1Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate abstract gateway at %s: %w", l1TokenGateway, err)
	}

	l2Gateway, err := abstractGateway.CounterpartGateway(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get counterpart gateway for gateway %s: %w", l1TokenGateway, err)
	}

	l2FilterName := "ArbitrumL2ToL1Bridge-L2Events"
	err = l2LogPoller.RegisterFilter(logpoller.Filter{
		Addresses: []common.Address{
			l2Gateway,           // emits DepositFinalized
			l2RebalancerAddress, // emits LiquidityTransferred
		},
		Name: l2FilterName,
		EventSigs: []common.Hash{
			DepositFinalizedTopic,     // emitted by the gateways
			LiquidityTransferredTopic, // emitted by the rebalancers
		},
		Retention: DurationMonth,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to register L2 log filter: %w", err)
	}

	l2GatewayWrapper, err := l2_arbitrum_gateway.NewL2ArbitrumGateway(l2Gateway, l2Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate l2 arbitrum gateway at %s: %w", l2Gateway, err)
	}

	lggr = lggr.Named("ArbitrumL1ToL2Bridge").With(
		"localSelector", localSelector,
		"remoteSelector", remoteSelector,
		"localChainID", localChain.EvmChainID,
		"remoteChainID", remoteChain.EvmChainID,
		"l1Rebalancer", l1Rebalancer.Address(),
		"l2Rebalancer", l2RebalancerAddress,
		"l1BridgeAdapter", l1BridgeAdapter.Address(),
		"l1GatewayRouter", l1GatewayRouter.Address(),
		"l1Inbox", l1Inbox.Address(),
		"l2Gateway", l2Gateway,
	)
	lggr.Infow("successfully initialized arbitrum L1 -> L2 bridge")

	return &l1ToL2Bridge{
		localSelector:       localSelector,
		remoteSelector:      remoteSelector,
		l1Rebalancer:        l1Rebalancer,
		l2RebalancerAddress: l2RebalancerAddress,
		l1BridgeAdapter:     l1BridgeAdapter,
		l1GatewayRouter:     l1GatewayRouter,
		l1Inbox:             l1Inbox,
		l2Gateway:           l2GatewayWrapper,
		l1Client:            l1Client,
		l2Client:            l2Client,
		l1LogPoller:         l1LogPoller,
		l2LogPoller:         l2LogPoller,
		l1FilterName:        l1FilterName,
		l2FilterName:        l2FilterName,
		lggr:                lggr,
	}, nil
}

func (l *l1ToL2Bridge) GetTransfers(
	ctx context.Context,
	localToken,
	remoteToken models.Address,
) ([]models.PendingTransfer, error) {
	fromTs := time.Now().Add(-24 * time.Hour) // last day
	erc20SentLogs, err := l.l1LogPoller.IndexedLogsCreatedAfter(
		ArbitrumL1ToL2ERC20Sent,
		l.l1BridgeAdapter.Address(),
		1, // topic index 1: localToken field in event
		[]common.Hash{
			common.HexToHash(common.Address(localToken).Hex()),
		},
		fromTs,
		logpoller.Finalized,
		pg.WithParentCtx(ctx),
	)
	// TODO: check if err is sql.ErrNoRows
	if err != nil {
		return nil, fmt.Errorf("failed to get ArbitrumL1ToL2ERC20Sent events from L1 bridge adapter: %w", err)
	}

	depositFinalizedLogs, err := l.l2LogPoller.IndexedLogsCreatedAfter(
		DepositFinalizedTopic,
		l.l2Gateway.Address(),
		3, // topic index 3: to address of deposit on L2
		[]common.Hash{
			common.HexToHash(l.l2RebalancerAddress.Hex()),
		},
		fromTs,
		logpoller.Finalized,
		pg.WithParentCtx(ctx),
	)
	// TODO: check if err is sql.ErrNoRows
	if err != nil {
		return nil, fmt.Errorf("failed to get DepositFinalized events from L2 gateway: %w", err)
	}

	liquidityTransferredLogs, err := l.l2LogPoller.IndexedLogsCreatedAfter(
		LiquidityTransferredTopic,
		l.l2RebalancerAddress,
		2, // topic index 2: fromChainSelector
		[]common.Hash{
			toHash(l.localSelector),
		},
		fromTs,
		logpoller.Finalized,
		pg.WithParentCtx(ctx),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get LiquidityTransferred events from L2 rebalancer: %w", err)
	}

	// the log poller SQL queries return logs sorted by block number and log index already
	// however given that this is an implementation detail we sort again here
	// but based on timestamp
	slices.SortFunc(erc20SentLogs, func(a, b logpoller.Log) int {
		return a.BlockTimestamp.Compare(b.BlockTimestamp)
	})
	slices.SortFunc(depositFinalizedLogs, func(a, b logpoller.Log) int {
		return a.BlockTimestamp.Compare(b.BlockTimestamp)
	})
	slices.SortFunc(liquidityTransferredLogs, func(a, b logpoller.Log) int {
		return a.BlockTimestamp.Compare(b.BlockTimestamp)
	})

	parsedERC20Sent, parsedToLP, err := l.parseL1ToL2Transfers(erc20SentLogs)
	if err != nil {
		return nil, fmt.Errorf("failed to parse L1 -> L2 transfers: %w", err)
	}

	parsedDepositFinalized, err := l.parseDepositFinalized(depositFinalizedLogs)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DepositFinalized logs: %w", err)
	}

	parsedLiquidityTransferred, err := l.parseLiquidityTransferred(liquidityTransferredLogs)
	if err != nil {
		return nil, fmt.Errorf("failed to parse LiquidityTransferred logs: %w", err)
	}

	// unfortunately its not easy to match DepositFinalized events with ERC20Sent events
	// reason being is that arbitrum does not emit any identifying information as part of the DepositFinalized
	// event, such as the l1 to l2 tx id. This is only available as part of the calldata for when the L2 calls
	// submitRetryable on the ArbRetryableTx precompile.
	// e.g https://sepolia.arbiscan.io/tx/0xce0d0d7e74f184fa8cb264b6d9aab5ced159faf3d0d9ae54b67fd40ba9d965a7
	// therefore we're kind of relegated here to doing a simple count check - filter out all of the
	// ERC20Sent logs destined for the rebalancer on L2 and all the DepositFinalized logs that
	// pay out to the rebalancer on L2.
	// this isn't a big deal because we can assume that the earlier ERC20Sent logs on L1
	// are more likely to be finalizedNotExecuted than later ones.
	notReady, ready, readyData, executed, err := l.partitionTransfers(localToken, parsedERC20Sent, parsedDepositFinalized, parsedLiquidityTransferred)
	if err != nil {
		return nil, fmt.Errorf("failed to partition logs into not-ready, ready, finalized and executed states: %w", err)
	}

	return l.toPendingTransfers(notReady, ready, readyData, executed, parsedToLP)
}

func (l *l1ToL2Bridge) toPendingTransfers(
	notReady,
	ready []*arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent,
	readyData [][]byte,
	executed []*arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent,
	parsedToLP map[logKey]logpoller.Log,
) ([]models.PendingTransfer, error) {
	if len(ready) != len(readyData) {
		return nil, fmt.Errorf("length of ready and readyData should be the same: len(ready) = %d, len(readyData) = %d",
			len(ready), len(readyData))
	}
	var transfers []models.PendingTransfer
	for _, transfer := range notReady {
		transfers = append(transfers, models.PendingTransfer{
			Transfer: models.Transfer{
				From:   l.localSelector,
				To:     l.remoteSelector,
				Amount: ubig.New(transfer.Amount),
				Date: parsedToLP[logKey{
					txHash:   transfer.Raw.TxHash,
					logIndex: int64(transfer.Raw.Index),
				}].BlockTimestamp,
				BridgeData: []byte{}, // no finalization data, not ready
			},
			Status: models.TransferStatusNotReady,
			ID:     fmt.Sprintf("%s-%d", transfer.Raw.TxHash.Hex(), transfer.Raw.Index),
		})
	}
	for i, transfer := range ready {
		transfers = append(transfers, models.PendingTransfer{
			Transfer: models.Transfer{
				From:   l.localSelector,
				To:     l.remoteSelector,
				Amount: ubig.New(transfer.Amount),
				Date: parsedToLP[logKey{
					txHash:   transfer.Raw.TxHash,
					logIndex: int64(transfer.Raw.Index),
				}].BlockTimestamp,
				BridgeData: readyData[i], // finalization data since its ready
			},
			Status: models.TransferStatusReady, // ready == finalized for L1 -> L2 transfers due to auto-finalization by the native bridge
			ID:     fmt.Sprintf("%s-%d", transfer.Raw.TxHash.Hex(), transfer.Raw.Index),
		})
	}
	for _, transfer := range executed {
		transfers = append(transfers, models.PendingTransfer{
			Transfer: models.Transfer{
				From:   l.localSelector,
				To:     l.remoteSelector,
				Amount: ubig.New(transfer.Amount),
				Date: parsedToLP[logKey{
					txHash:   transfer.Raw.TxHash,
					logIndex: int64(transfer.Raw.Index),
				}].BlockTimestamp,
				BridgeData: []byte{}, // no finalization data, already executed
			},
			Status: models.TransferStatusExecuted,
			ID:     fmt.Sprintf("%s-%d", transfer.Raw.TxHash.Hex(), transfer.Raw.Index),
		})
	}
	return transfers, nil
}

// precondition: the input logs are already sorted in time-ascending order
func (l *l1ToL2Bridge) partitionTransfers(
	localToken models.Address,
	erc20SentLogs []*arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent,
	depositFinalizedLogs []*l2_arbitrum_gateway.L2ArbitrumGatewayDepositFinalized,
	liquidityTransferredLogs []*rebalancer.RebalancerLiquidityTransferred,
) (
	notReady,
	ready []*arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent,
	readyData [][]byte,
	executed []*arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent,
	err error,
) {
	effectiveERC20Sent, effectiveDepositFinalized := l.getEffectiveEvents(localToken, erc20SentLogs, depositFinalizedLogs)
	// determine ready and not ready first
	if len(effectiveERC20Sent) > len(effectiveDepositFinalized) {
		// more sent than have been finalized
		for i := len(effectiveERC20Sent) - len(effectiveDepositFinalized) + 1; i < len(effectiveERC20Sent); i++ {
			notReady = append(notReady, effectiveERC20Sent[i])
		}
		for i := 0; i < (len(effectiveERC20Sent) - len(effectiveDepositFinalized)); i++ {
			ready = append(ready, effectiveERC20Sent[i])
		}
	} else if len(effectiveERC20Sent) < len(effectiveDepositFinalized) {
		// more finalized than have been sent - should be impossible
		return nil, nil, nil, nil, fmt.Errorf("got more finalized logs than sent - should be impossible: len(sent) = %d, len(finalized) = %d",
			len(effectiveERC20Sent), len(effectiveDepositFinalized))
	} else {
		ready = effectiveERC20Sent
	}
	// figure out if any of the ready have been executed
	ready, executed, err = l.filterExecuted(ready, liquidityTransferredLogs)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to filter executed transfers: %w", err)
	}
	// get the readyData
	// this is just going to be the L1 to L2 tx id that is emitted in the ERC20Sent log itself
	for _, r := range ready {
		readyData = append(readyData, r.OutboundTransferResult)
	}
	return
}

func (l *l1ToL2Bridge) filterExecuted(
	readyCandidates []*arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent,
	liquidityTransferredLogs []*rebalancer.RebalancerLiquidityTransferred,
) (
	ready,
	executed []*arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent,
	err error,
) {
	for _, readyCandidate := range readyCandidates {
		found, err := l.matchingExecutionExists(readyCandidate, liquidityTransferredLogs)
		if err != nil {
			return nil, nil, fmt.Errorf("error checking if ready candidate has been executed: %w", err)
		}
		if !found {
			executed = append(executed, readyCandidate)
		} else {
			ready = append(ready, readyCandidate)
		}
	}
	return
}

// TODO: might be able to optimize this
// map[l2ToL1TxId]bool and check if the l2ToL1TxId exists in the map
func (l *l1ToL2Bridge) matchingExecutionExists(
	readyCandidate *arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent,
	liquidityTransferredLogs []*rebalancer.RebalancerLiquidityTransferred,
) (bool, error) {
	for _, ltLog := range liquidityTransferredLogs {
		// decode the bridge specific data, which should be the l1 -> l2 tx id
		ltL1ToL2TxId, err := unpackUint256(ltLog.BridgeSpecificData)
		if err != nil {
			return false, fmt.Errorf("failed to unpack bridge specific data from LiquidityTransferred log: %w, data: %s",
				err, hexutil.Encode(ltLog.BridgeSpecificData))
		}
		l1ToL2TxId, err := unpackUint256(readyCandidate.OutboundTransferResult)
		if err != nil {
			return false, fmt.Errorf("failed to unpack outbound transfer result from ArbitrumL1ToL2ERC20Sent log: %w, data: %s",
				err, hexutil.Encode(readyCandidate.OutboundTransferResult))
		}
		if l1ToL2TxId.Cmp(ltL1ToL2TxId) == 0 {
			return true, nil
		}
	}
	return false, nil
}

func (l *l1ToL2Bridge) getEffectiveEvents(
	localToken models.Address,
	erc20SentLogs []*arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent,
	depositFinalizedLogs []*l2_arbitrum_gateway.L2ArbitrumGatewayDepositFinalized,
) (
	effectiveERC20Sent []*arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent,
	effectiveDepositFinalized []*l2_arbitrum_gateway.L2ArbitrumGatewayDepositFinalized,
) {
	// filter out ERC20Sent logs not destined for the rebalancer on L2
	// TODO: ideally this would be done in the log poller query but no such query exists
	// at the moment.
	for _, erc20Sent := range erc20SentLogs {
		if erc20Sent.Recipient == l.l2RebalancerAddress {
			effectiveERC20Sent = append(effectiveERC20Sent, erc20Sent)
		}
	}

	// filter out DepositFinalized logs not coming from the l1 bridge adapter
	// and not matching the localToken provided.
	// in theory anyone can bridge any token to the rebalancer on L2 from L1
	// TODO: ideally this would be done in the log poller query but no such query exists
	// at the moment.
	// TODO: should we care about L1 -> L2 bridges not done by the bridge adapter?
	// in theory those are funds that can be injected into the pools.
	for _, depFinalized := range depositFinalizedLogs {
		if depFinalized.From == l.l1BridgeAdapter.Address() &&
			depFinalized.L1Token == common.Address(localToken) {
			effectiveDepositFinalized = append(effectiveDepositFinalized, depFinalized)
		}
	}
	return
}

func (l *l1ToL2Bridge) parseL1ToL2Transfers(lgs []logpoller.Log) (
	[]*arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent,
	map[logKey]logpoller.Log,
	error,
) {
	transfers := make([]*arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterArbitrumL1ToL2ERC20Sent, len(lgs))
	parsedToLPLog := make(map[logKey]logpoller.Log)
	for i, lg := range lgs {
		parsed, err := l.l1BridgeAdapter.ParseArbitrumL1ToL2ERC20Sent(lg.ToGethLog())
		if err != nil {
			// should never happen
			return nil, nil, fmt.Errorf("failed to parse L1 -> L2 transfer log: %w", err)
		}
		transfers[i] = parsed
		parsedToLPLog[logKey{
			txHash:   lg.TxHash,
			logIndex: lg.LogIndex,
		}] = lg
	}
	return transfers, parsedToLPLog, nil
}

func (l *l1ToL2Bridge) parseDepositFinalized(lgs []logpoller.Log) ([]*l2_arbitrum_gateway.L2ArbitrumGatewayDepositFinalized, error) {
	finalized := make([]*l2_arbitrum_gateway.L2ArbitrumGatewayDepositFinalized, len(lgs))
	for i, lg := range lgs {
		parsed, err := l.l2Gateway.ParseDepositFinalized(lg.ToGethLog())
		if err != nil {
			// should never happen
			return nil, fmt.Errorf("failed to parse DepositFinalized log: %w", err)
		}
		finalized[i] = parsed
	}
	return finalized, nil
}

func (l *l1ToL2Bridge) parseLiquidityTransferred(lgs []logpoller.Log) ([]*rebalancer.RebalancerLiquidityTransferred, error) {
	transferred := make([]*rebalancer.RebalancerLiquidityTransferred, len(lgs))
	for i, lg := range lgs {
		parsed, err := l.l1Rebalancer.ParseLiquidityTransferred(lg.ToGethLog())
		if err != nil {
			// should never happen
			return nil, fmt.Errorf("failed to parse LiquidityTransferred log: %w", err)
		}
		transferred[i] = parsed
	}
	return transferred, nil
}

func (l *l1ToL2Bridge) QuorumizedBridgePayload(payloads [][]byte) ([]byte, error) {
	// TODO: decode and take top n-f index after sorting asc gasLimit/maxSubmissionCost/maxFeePerGas
	return payloads[0], nil
}

// GetBridgePayloadAndFee implements bridge.Bridge
// For Arbitrum L1 -> L2 transfers, the bridge specific payload is a tuple of 3 numbers:
// 1. gasLimit
// 2. maxSubmissionCost
// 3. maxFeePerGas
func (l *l1ToL2Bridge) GetBridgePayloadAndFee(
	ctx context.Context,
	transfer models.Transfer,
) ([]byte, *big.Int, error) {
	l1Gateway, err := l.l1GatewayRouter.GetGateway(&bind.CallOpts{
		Context: ctx,
	}, common.Address(transfer.LocalTokenAddress))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get L1 gateway for local token %s: %w",
			transfer.LocalTokenAddress, err)
	}

	l1TokenGateway, err := arbitrum_token_gateway.NewArbitrumTokenGateway(l1Gateway, l.l1Client)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to instantiate L1 token gateway at %s: %w",
			l1Gateway, err)
	}

	// get the counterpart gateway on L2 from the L1 gateway
	// unfortunately we need to instantiate a new wrapper because the counterpartGateway field,
	// although it is public, is not accessible via a getter function on the token gateway interface
	abstractGateway, err := abstract_arbitrum_token_gateway.NewAbstractArbitrumTokenGateway(l1Gateway, l.l1Client)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to instantiate abstract gateway at %s: %w",
			l1Gateway, err)
	}

	l2Gateway, err := abstractGateway.CounterpartGateway(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get counterpart gateway for L1 gateway %s: %w",
			l1Gateway, err)
	}

	retryableData := RetryableData{
		From:                l1Gateway,
		To:                  l2Gateway,
		ExcessFeeRefundAddr: common.Address(transfer.Receiver),
		CallValueRefundAddr: common.Address(transfer.Sender),
		// typically just one
		L2CallValue: big.NewInt(1),
		// 3 seems to work, but not sure if it's the best value
		// you definitely need a non-nil deposit for the NodeInterface call to succeed
		Deposit: big.NewInt(3),
		// MaxSubmissionCost: , // To be filled in
		// GasLimit: , // To be filled in
		// MaxFeePerGas: , // To be filled in
		// Data: , // To be filled in
	}

	// determine the finalizeInboundTransfer calldata
	finalizeInboundTransferCalldata, err := l1TokenGateway.GetOutboundCalldata(
		nil,
		common.Address(transfer.LocalTokenAddress), // L1 token address
		l.l1BridgeAdapter.Address(),                // L1 sender address
		common.Address(transfer.Receiver),          // L2 recipient address
		transfer.Amount.ToInt(),                    // token amount
		[]byte{},                                   // extra data (unused here)
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get finalizeInboundTransfer calldata: %w", err)
	}
	retryableData.Data = finalizeInboundTransferCalldata

	l.lggr.Infow("Constructed RetryableData",
		"from", retryableData.From,
		"to", retryableData.To,
		"excessFeeRefundAddr", retryableData.ExcessFeeRefundAddr,
		"callValueRefundAddr", retryableData.CallValueRefundAddr,
		"l2CallValue", retryableData.L2CallValue,
		"deposit", retryableData.Deposit,
		"data", hexutil.Encode(retryableData.Data))

	l1BaseFee, err := l.l1Client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get L1 base fee: %w", err)
	}

	return l.estimateAll(ctx, retryableData, l1BaseFee)
}

func (l *l1ToL2Bridge) estimateAll(
	ctx context.Context,
	retryableData RetryableData,
	l1BaseFee *big.Int,
) ([]byte, *big.Int, error) {
	l2MaxFeePerGas, err := l.estimateMaxFeePerGasOnL2(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to estimate max fee per gas on L2: %w", err)
	}

	maxSubmissionFee, err := l.estimateMaxSubmissionFee(ctx, l1BaseFee, len(retryableData.Data))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to estimate max submission fee: %w", err)
	}

	gasLimit, err := l.estimateRetryableGasLimit(ctx, retryableData)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to estimate retryable gas limit: %w", err)
	}

	deposit := new(big.Int).Mul(gasLimit, l2MaxFeePerGas)
	deposit = deposit.Add(deposit, maxSubmissionFee)

	l.lggr.Infow("Estimated L1 -> L2 fees",
		"gasLimit", gasLimit,
		"maxSubmissionFee", maxSubmissionFee,
		"l2MaxFeePerGas", l2MaxFeePerGas,
		"deposit", deposit)

	bridgeCalldata, err := l1AdapterABI.Pack("exposeSendERC20Params", arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterSendERC20Params{
		GasLimit:          gasLimit,
		MaxSubmissionCost: maxSubmissionFee,
		MaxFeePerGas:      l2MaxFeePerGas,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to pack bridge calldata for bridge adapter: %w", err)
	}
	bridgeCalldata = bridgeCalldata[4:] // remove method id
	return bridgeCalldata, deposit, nil
}

func (l *l1ToL2Bridge) estimateRetryableGasLimit(ctx context.Context, rd RetryableData) (*big.Int, error) {
	packed, err := nodeInterfaceABI.Pack("estimateRetryableTicket",
		rd.From,
		assets.Ether(1),
		rd.To,
		rd.L2CallValue,
		rd.ExcessFeeRefundAddr,
		rd.CallValueRefundAddr,
		rd.Data,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack estimateRetryableTicket call: %w", err)
	}

	gasLimit, err := l.l2Client.EstimateGas(ctx, ethereum.CallMsg{
		To:   &NodeInterfaceAddress,
		Data: packed,
	})
	if err != nil {
		return nil, fmt.Errorf("error esimtating gas on node interface for estimateRetryableTicket: %s, calldata: %s",
			err, hexutil.Encode(packed))
	}

	// no multiplier on gas limit
	// should be pretty accurate
	return big.NewInt(int64(gasLimit)), nil
}

func (l *l1ToL2Bridge) estimateMaxSubmissionFee(
	ctx context.Context,
	l1BaseFee *big.Int,
	dataLength int,
) (*big.Int, error) {
	submissionFee, err := l.l1Inbox.CalculateRetryableSubmissionFee(&bind.CallOpts{
		Context: ctx,
	}, big.NewInt(int64(dataLength)), l1BaseFee)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate retryable submission fee: %w", err)
	}

	submissionFee = submissionFee.Mul(submissionFee, submissionFeeMultiplier)
	return submissionFee, nil
}

func (l *l1ToL2Bridge) estimateMaxFeePerGasOnL2(ctx context.Context) (*big.Int, error) {
	l2BaseFee, err := l.l2Client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to suggest gas price on L2: %w", err)
	}

	l2BaseFee = l2BaseFee.Mul(l2BaseFee, l2BaseFeeMultiplier)
	return l2BaseFee, nil
}

func (l *l1ToL2Bridge) Close(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (l *l1ToL2Bridge) LocalChainSelector() models.NetworkSelector {
	return l.localSelector
}

func (l *l1ToL2Bridge) RemoteChainSelector() models.NetworkSelector {
	return l.remoteSelector
}

type RetryableData struct {
	// From is the gateway on L1 that will be sending the funds to the L2 gateway.
	From common.Address
	// To is the gateway on L2 that will be receiving the funds and eventually
	// sending them to the final recipient.
	To                common.Address
	L2CallValue       *big.Int
	Deposit           *big.Int
	MaxSubmissionCost *big.Int
	// ExcessFeeRefundAddr is an address on L2 that will be receiving excess fees
	ExcessFeeRefundAddr common.Address
	// CallValueRefundAddr is an address on L1 that will be receiving excess fees
	CallValueRefundAddr common.Address
	GasLimit            *big.Int
	MaxFeePerGas        *big.Int
	// Data is the calldata for the L2 gateway's `finalizeInboundTransfer` method.
	// The final recipient on L2 is specified in this calldata.
	Data []byte
}

func toHash(selector models.NetworkSelector) common.Hash {
	encoded := hexutil.EncodeUint64(uint64(selector))
	return common.HexToHash(encoded)
}
