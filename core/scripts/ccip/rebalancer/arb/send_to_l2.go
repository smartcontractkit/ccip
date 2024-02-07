package arb

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip/rebalancer/multienv"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/assets"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/erc20"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arb_node_interface"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_gateway_router"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_inbox"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_l1_bridge_adapter"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_l1_gateway_router"
)

func SendToL2(
	env multienv.Env,
	l1ChainID,
	l2ChainID uint64,
	l1BridgeAdapterAddress,
	l1TokenAddress,
	l2Recipient common.Address,
	amount *big.Int,
) {
	// do some basic checks before proceeding
	token, err := erc20.NewERC20(l1TokenAddress, env.Clients[l1ChainID])
	helpers.PanicErr(err)

	// check if we have enough balance otherwise approve will fail
	balance, err := token.BalanceOf(nil, env.Transactors[l1ChainID].From)
	helpers.PanicErr(err)
	if balance.Cmp(amount) < 0 {
		panic(fmt.Sprintf("Insufficient balance, get more tokens or specify smaller amount: %s < %s", balance, amount))
	}

	gatewayRouter, err := arbitrum_gateway_router.NewArbitrumGatewayRouter(ArbitrumContracts[l1ChainID]["L1GatewayRouter"], env.Clients[l1ChainID])
	helpers.PanicErr(err)

	depositFunc := createDepositFunc(
		env.Transactors[l1ChainID].From,
		ArbitrumContracts[l1ChainID]["L1GatewayRouter"],
		l1TokenAddress,
		l2Recipient,
		amount,
	)

	params, _, _, _, _, value := populateFunctionParams(env, l1ChainID, l2ChainID, gatewayRouter, token, depositFunc)

	// call the L1 adapter to send the funds to L2
	// first approve the L1 adapter to spend the tokens
	// check allowance so we don't approve unnecessarily
	allowance, err := token.Allowance(nil, env.Transactors[l1ChainID].From, l1BridgeAdapterAddress)
	helpers.PanicErr(err)
	if allowance.Cmp(amount) < 0 {
		tx, err2 := token.Approve(env.Transactors[l1ChainID], l1BridgeAdapterAddress, amount)
		helpers.PanicErr(err2)
		helpers.ConfirmTXMined(context.Background(), env.Clients[l1ChainID], tx, int64(l1ChainID),
			"Approve", amount.String(), "to", l1BridgeAdapterAddress.String())

		// check allowance
		allowance, err2 = token.Allowance(nil, env.Transactors[l1ChainID].From, l1BridgeAdapterAddress)
		helpers.PanicErr(err2)
		if allowance.Cmp(amount) < 0 {
			panic(fmt.Sprintf("Allowance failed, expected %s, got %s", amount, allowance))
		}
	} else {
		fmt.Println("Allowance already set to", allowance, "for", l1BridgeAdapterAddress.String())
	}

	// transact with the bridge adapter to send funds cross-chain
	l1AdapterABI, err := arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterMetaData.GetAbi()
	helpers.PanicErr(err)
	bridgeCalldata, err := l1AdapterABI.Pack("exposeSendERC20Params", arbitrum_l1_bridge_adapter.ArbitrumL1BridgeAdapterSendERC20Params{
		GasLimit:          params.GasLimit,
		MaxSubmissionCost: params.MaxSubmissionCost,
		MaxFeePerGas:      params.MaxFeePerGas,
	})
	helpers.PanicErr(err)
	bridgeCalldata = bridgeCalldata[4:] // remove the method id
	sendERC20Calldata, err := l1AdapterABI.Pack("sendERC20", l1TokenAddress, common.HexToAddress("0x0"), l2Recipient, amount, bridgeCalldata)
	helpers.PanicErr(err)

	fmt.Println("Sending ERC20 to L2:", "\n",
		"l1TokenAddress:", l1TokenAddress.String(), "\n",
		"l2Recipient:", l2Recipient.String(), "\n",
		"amount:", amount, "\n",
		"calldata:", hexutil.Encode(bridgeCalldata), "\n",
		"value:", value)

	gasPrice, err := env.Clients[l1ChainID].SuggestGasPrice(context.Background())
	helpers.PanicErr(err)
	nonce, err := env.Clients[l1ChainID].PendingNonceAt(context.Background(), env.Transactors[l1ChainID].From)
	helpers.PanicErr(err)
	rawTx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &l1BridgeAdapterAddress,
		GasPrice: gasPrice,
		Gas:      1e6,
		Value:    value,
		Data:     sendERC20Calldata,
	})
	signedTx, err := env.Transactors[l1ChainID].Signer(env.Transactors[l1ChainID].From, rawTx)
	helpers.PanicErr(err)
	err = env.Clients[l1ChainID].SendTransaction(context.Background(), signedTx)
	helpers.PanicErr(err)
	helpers.ConfirmTXMined(context.Background(), env.Clients[l1ChainID], signedTx, int64(l1ChainID),
		"Calling SendERC20, amount:", amount.String(), ", on adapter at:", l1BridgeAdapterAddress.String())
}

type L1ToL2MessageGasParams struct {
	GasLimit          *big.Int
	MaxSubmissionCost *big.Int
	MaxFeePerGas      *big.Int
	Deposit           *big.Int
}

func populateFunctionParams(
	env multienv.Env,
	l1ChainID,
	l2ChainID uint64,
	gatewayRouter *arbitrum_gateway_router.ArbitrumGatewayRouter,
	l1Token *erc20.ERC20,
	f dataFunc,
) (
	L1ToL2MessageGasParams,
	RetryableData,
	[]byte,
	common.Address,
	common.Address,
	*big.Int,
) {
	// call initially with the dummy values to trigger the revert
	data, to, from, value := f(big.NewInt(1), big.NewInt(1), big.NewInt(1))

	l1Client := env.Clients[l1ChainID]
	gasPrice, err := l1Client.SuggestGasPrice(context.Background())
	helpers.PanicErr(err)

	gatewayToApprove, err := gatewayRouter.GetGateway(nil, l1Token.Address())
	helpers.PanicErr(err)

	// check allowance so we don't approve unnecessarily
	allowance, err := l1Token.Allowance(nil, env.Transactors[l1ChainID].From, gatewayToApprove)
	helpers.PanicErr(err)
	if allowance.Cmp(value) < 0 {
		tx, err2 := l1Token.Approve(env.Transactors[l1ChainID], gatewayToApprove, value)
		helpers.PanicErr(err2)
		helpers.ConfirmTXMined(context.Background(), env.Clients[l1ChainID], tx, int64(l1ChainID),
			"Approve", value.String(), "to gateway at", gatewayToApprove.String())

		// check allowance
		allowance, err2 = l1Token.Allowance(nil, env.Transactors[l1ChainID].From, gatewayToApprove)
		helpers.PanicErr(err2)
		if allowance.Cmp(value) < 0 {
			panic(fmt.Sprintf("Allowance failed, expected %s, got %s", value, allowance))
		}
	} else {
		fmt.Println("Allowance already set to", allowance)
	}

	fmt.Println("Making eth_call to trigger revert with data:", hexutil.Encode(data),
		"value:", value.String(), "to:", to.String(), "from:", from.String())
	_, err = l1Client.CallContract(context.Background(), ethereum.CallMsg{
		To:       &to,
		Data:     data,
		From:     from,
		Value:    value,
		GasPrice: new(big.Int).Mul(gasPrice, big.NewInt(2)),
	}, nil)
	if err == nil {
		panic(fmt.Sprintf("Expected an error when calling with dummy values: %v", err))
	}
	jErr, err := client.ExtractRPCError(err)
	helpers.PanicErr(err)
	if jErr == nil {
		panic(fmt.Sprintf("Expected an error with a JSON-RPC error code: %+v", err))
	}

	fmt.Println("Got JSON-RPC error, code:", jErr.Code, "message:", jErr.Message, "data:", jErr.Data)

	errData, ok := jErr.Data.(string)
	if !ok {
		panic(fmt.Sprintf("Expected error data to be a string: %+v", jErr.Data))
	}

	baseFee, err := l1Client.SuggestGasPrice(context.Background())
	helpers.PanicErr(err)

	rd := parseRetryableError(errData)
	estimates := estimateAll(env, l1ChainID, l2ChainID, rd, baseFee)

	// form real data for the transaction
	data, to, from, value = f(estimates.GasLimit, estimates.MaxFeePerGas, estimates.MaxSubmissionCost)

	return estimates, rd, data, to, from, value
}

func estimateAll(env multienv.Env, l1ChainID, l2ChainID uint64, rd RetryableData, l1BaseFee *big.Int) L1ToL2MessageGasParams {
	data := rd.Data

	// NOTE: this is the L2 base fee, not L1
	l2Client := env.Clients[l2ChainID]
	maxFeePerGas := estimateMaxFeePerGas(l2Client)

	maxSubmissionFee := estimateSubmissionFee(env.Clients[l1ChainID], l1ChainID, l1BaseFee, uint64(len(data)))

	gasLimit := estimateRetryableGasLimit(env, l2Client, l2ChainID, rd)

	deposit := new(big.Int).Mul(gasLimit, maxFeePerGas)
	deposit = deposit.Add(deposit, maxSubmissionFee)

	fmt.Println("estimated L1 -> L2 fees:", "\n",
		"gasLimit:", gasLimit, "\n",
		"maxSubmissionFee:", maxSubmissionFee, "\n",
		"maxFeePerGas:", maxFeePerGas, "\n",
		"deposit:", deposit)

	return L1ToL2MessageGasParams{
		GasLimit:          gasLimit,
		MaxSubmissionCost: maxSubmissionFee,
		MaxFeePerGas:      maxFeePerGas,
		Deposit:           deposit,
	}
}

func estimateRetryableGasLimit(env multienv.Env, l2Client *ethclient.Client, l2ChainID uint64, rd RetryableData) *big.Int {
	nodeInterfaceABI, err := arb_node_interface.NodeInterfaceMetaData.GetAbi()
	helpers.PanicErr(err)

	packed, err := nodeInterfaceABI.Pack("estimateRetryableTicket",
		rd.From,
		assets.Ether(1).ToInt(), // this is what is done in the SDK, not sure why yet
		rd.To,
		rd.L2CallValue,
		rd.ExcessFeeRefundAddr,
		rd.CallValueRefundAddr,
		rd.Data)
	helpers.PanicErr(err)

	l2Balance, err := l2Client.BalanceAt(context.Background(), env.Transactors[l2ChainID].From, nil)
	helpers.PanicErr(err)
	if l2Balance.Cmp(rd.Deposit) < 0 {
		panic(fmt.Sprintf("Insufficient balance on L2: %s < %s", l2Balance, rd.Deposit))
	}

	fmt.Println("calling node interface with calldata:", hexutil.Encode(packed), "value:", rd.Deposit)
	nodeInterfaceAddr := ArbitrumContracts[l2ChainID]["NodeInterface"]
	gasLimit, err := l2Client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &nodeInterfaceAddr,
		Data: packed,
	})
	helpers.PanicErr(err)

	// no percent increase on gas limit
	// should be pretty accurate
	return big.NewInt(int64(gasLimit))
}

func estimateMaxFeePerGas(l2Client *ethclient.Client) *big.Int {
	l2BaseFee, err := l2Client.SuggestGasPrice(context.Background())
	helpers.PanicErr(err)
	// base fee on L2 is bumped by 200% by the arbitrum sdk (i.e 3x)
	l2BaseFee = new(big.Int).Mul(l2BaseFee, big.NewInt(3))
	return l2BaseFee
}

func estimateSubmissionFee(l1Client *ethclient.Client, l1ChainID uint64, l1BaseFee *big.Int, calldataSize uint64) *big.Int {
	inbox, err := arbitrum_inbox.NewArbitrumInbox(ArbitrumContracts[l1ChainID]["L1Inbox"], l1Client)
	helpers.PanicErr(err)

	submissionFee, err := inbox.CalculateRetryableSubmissionFee(nil, big.NewInt(int64(calldataSize)), l1BaseFee)
	helpers.PanicErr(err)

	// submission fee is bumped by 300% (i.e 4x) by the arbitrum sdk
	// do the same here
	submissionFee = submissionFee.Mul(submissionFee, big.NewInt(4))

	return submissionFee
}

type RetryableData struct {
	From                common.Address
	To                  common.Address
	L2CallValue         *big.Int
	Deposit             *big.Int
	MaxSubmissionCost   *big.Int
	ExcessFeeRefundAddr common.Address
	CallValueRefundAddr common.Address
	GasLimit            *big.Int
	MaxFeePerGas        *big.Int
	Data                []byte
}

func parseRetryableError(errData string) RetryableData {
	dataBytes, err := hexutil.Decode(errData)
	helpers.PanicErr(err)

	retryableIfaces, err := utils.ABIDecode(
		`[{"type":"address"},{"type":"address"},{"type":"uint256"},{"type":"uint256"},{"type":"uint256"},{"type":"address"},{"type":"address"},{"type":"uint256"},{"type":"uint256"},{"type":"bytes"}]`,
		dataBytes[4:])
	helpers.PanicErr(err)

	if len(retryableIfaces) != 10 {
		panic(fmt.Sprintf("Expected 10 inputs, got %d", len(retryableIfaces)))
	}

	from := *abi.ConvertType(retryableIfaces[0], new(common.Address)).(*common.Address)
	to := *abi.ConvertType(retryableIfaces[1], new(common.Address)).(*common.Address)
	l2CallValue := *abi.ConvertType(retryableIfaces[2], new(*big.Int)).(**big.Int)
	deposit := *abi.ConvertType(retryableIfaces[3], new(*big.Int)).(**big.Int)
	maxSubmissionCost := *abi.ConvertType(retryableIfaces[4], new(*big.Int)).(**big.Int)
	excessFeeRefundAddress := *abi.ConvertType(retryableIfaces[5], new(common.Address)).(*common.Address)
	callValueRefundAddress := *abi.ConvertType(retryableIfaces[6], new(common.Address)).(*common.Address)
	gasLimit := *abi.ConvertType(retryableIfaces[7], new(*big.Int)).(**big.Int)
	maxFeePerGas := *abi.ConvertType(retryableIfaces[8], new(*big.Int)).(**big.Int)
	data := *abi.ConvertType(retryableIfaces[9], new([]byte)).(*[]byte)

	fmt.Println("Successfully parsed retryable error:", "\n",
		"from:", from, "\n",
		"to:", to, "\n",
		"l2CallValue:", l2CallValue, "\n",
		"deposit:", deposit, "\n",
		"maxSubmissionCost:", maxSubmissionCost, "\n",
		"excessFeeRefundAddress:", excessFeeRefundAddress, "\n",
		"callValueRefundAddress:", callValueRefundAddress, "\n",
		"gasLimit:", gasLimit, "\n",
		"maxFeePerGas:", maxFeePerGas, "\n",
		"data:", hexutil.Encode(data))

	return RetryableData{
		From:                from,
		To:                  to,
		L2CallValue:         l2CallValue,
		Deposit:             deposit,
		MaxSubmissionCost:   maxSubmissionCost,
		ExcessFeeRefundAddr: excessFeeRefundAddress,
		CallValueRefundAddr: callValueRefundAddress,
		GasLimit:            gasLimit,
		MaxFeePerGas:        maxFeePerGas,
		Data:                data,
	}
}

/**
* Function that will internally make an L1->L2 transaction
* Will initially be called with dummy values to trigger a special revert containing
* the real params. Then called again with the real params to form the final data to be submitted
 */
type dataFunc func(gasLimit, maxFeePerGas, maxSubmissionCost *big.Int) (data []byte, to, from common.Address, value *big.Int)

func createDepositFunc(fromAddress, l1GatewayRouter, l1TokenAddress, l2Recipient common.Address, amount *big.Int) dataFunc {
	return func(gasLimit, maxFeePerGas, maxSubmissionCost *big.Int) (calldata []byte, to, from common.Address, value *big.Int) {
		innerData, err := utils.ABIEncode(`[{"type":"uint256"}, {"type": "bytes"}]`, maxSubmissionCost, []byte(""))
		helpers.PanicErr(err)

		grABI, err := arbitrum_l1_gateway_router.ArbitrumL1GatewayRouterMetaData.GetAbi()
		helpers.PanicErr(err)

		calldata, err = grABI.Pack("outboundTransferCustomRefund", l1TokenAddress, l2Recipient, l2Recipient, amount, gasLimit, maxFeePerGas, innerData)
		helpers.PanicErr(err)

		to = l1GatewayRouter
		from = fromAddress

		value = new(big.Int).Mul(gasLimit, maxFeePerGas)
		value = value.Add(value, maxSubmissionCost)

		return
	}
}
