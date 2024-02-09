package arb

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/abstract_arbitrum_token_gateway"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arb_node_interface"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_gateway_router"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_inbox"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_l1_bridge_adapter"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_token_gateway"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/bridge"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

var (
	_ bridge.Bridge = (*l1ToL2Bridge)(nil)

	// Multipliers to ensure our L1 -> L2 tx goes through
	// These values match the arbitrum SDK
	l2BaseFeeMultiplier     = big.NewInt(3)
	submissionFeeMultiplier = big.NewInt(4)

	nodeInterfaceABI = abihelpers.MustParseABI(arb_node_interface.NodeInterfaceMetaData.ABI)
)

type l1ToL2Bridge struct {
	localSelector  models.NetworkSelector
	remoteSelector models.NetworkSelector

	l1BridgeAdapterAddress common.Address

	// Arbitrum contract addresses
	l1GatewayRouter *arbitrum_gateway_router.ArbitrumGatewayRouter
	l2GatewayRouter *arbitrum_gateway_router.ArbitrumGatewayRouter
	l1Inbox         *arbitrum_inbox.ArbitrumInbox

	l1Client client.Client
	l2Client client.Client

	lggr logger.Logger
}

func (l *l1ToL2Bridge) GetTransfers(
	ctx context.Context,
	localToken,
	remoteToken models.Address,
) ([]models.PendingTransfer, error) {
	//TODO implement me
	panic("implement me")
}

// GetBridgeSpecificPayload implements bridge.Bridge
// For Arbitrum L1 -> L2 transfers, the bridge specific payload is a tuple of 3 numbers:
// 1. gasLimit
// 2. maxSubmissionCost
// 3. maxFeePerGas
func (l *l1ToL2Bridge) GetBridgeSpecificPayload(
	ctx context.Context,
	transfer models.Transfer,
) ([]byte, error) {
	l1Gateway, err := l.l1GatewayRouter.GetGateway(&bind.CallOpts{
		Context: ctx,
	}, common.Address(transfer.LocalTokenAddress))
	if err != nil {
		return nil, fmt.Errorf("failed to get L1 gateway for local token %s: %w",
			transfer.LocalTokenAddress, err)
	}

	l1TokenGateway, err := arbitrum_token_gateway.NewArbitrumTokenGateway(l1Gateway, l.l1Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate L1 token gateway at %s: %w",
			l1Gateway, err)
	}

	// get the counterpart gateway on L2 from the L1 gateway
	// unfortunately we need to instantiate a new wrapper because the counterpartGateway field,
	// although it is public, is not accessible via a getter function on the token gateway interface
	abstractGateway, err := abstract_arbitrum_token_gateway.NewAbstractArbitrumTokenGateway(l1Gateway, l.l1Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate abstract gateway at %s: %w",
			l1Gateway, err)
	}

	l2Gateway, err := abstractGateway.CounterpartGateway(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get counterpart gateway for L1 gateway %s: %w",
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
		l.l1BridgeAdapterAddress,                   // L1 sender address
		common.Address(transfer.Receiver),          // L2 recipient address
		transfer.Amount,                            // token amount
		[]byte{},                                   // extra data (unused here)
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get finalizeInboundTransfer calldata: %w", err)
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
		return nil, fmt.Errorf("failed to get L1 base fee: %w", err)
	}

	return l.estimateAll(ctx, retryableData, l1BaseFee)
}

func (l *l1ToL2Bridge) estimateAll(
	ctx context.Context,
	retryableData RetryableData,
	l1BaseFee *big.Int,
) ([]byte, error) {
	l2MaxFeePerGas, err := l.estimateMaxFeePerGasOnL2(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate max fee per gas on L2: %w", err)
	}

	maxSubmissionFee, err := l.estimateMaxSubmissionFee(ctx, l1BaseFee, len(retryableData.Data))
	if err != nil {
		return nil, fmt.Errorf("failed to estimate max submission fee: %w", err)
	}

	gasLimit, err := l.estimateRetryableGasLimit(ctx, retryableData)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate retryable gas limit: %w", err)
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
		return nil, fmt.Errorf("failed to pack bridge calldata for bridge adapter: %w", err)
	}
	bridgeCalldata = bridgeCalldata[4:] // remove method id
	return bridgeCalldata, nil
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
