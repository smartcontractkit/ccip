package opstack

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip/rebalancer/multienv"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/optimism_l2_bridge_adapter"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/shared/generated/erc20"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
)

var (
	l2AdapterABI = abihelpers.MustParseABI(optimism_l2_bridge_adapter.OptimismL2BridgeAdapterABI)
)

func WithdrawFromL2(
	env multienv.Env,
	l2ChainID uint64,
	l2BridgeAdapterAddress common.Address,
	amount *big.Int,
	l1ToAddress,
	l2TokenAddress common.Address,
) {
	token, err := erc20.NewERC20(l2TokenAddress, env.Clients[l2ChainID])
	helpers.PanicErr(err)

	// check if we have enough balance
	balance, err := token.BalanceOf(nil, env.Transactors[l2ChainID].From)
	helpers.PanicErr(err)
	if balance.Cmp(amount) < 0 {
		panic(fmt.Sprintf("not enough balance to withdraw, get more tokens or specify less amount, bal: %s, want: %s",
			balance.String(), amount.String()))
	}

	// l2Adapter, err := optimism_l2_bridge_adapter.NewOptimismL2BridgeAdapter(l2BridgeAdapterAddress, env.Clients[l2ChainID])
	// helpers.PanicErr(err)

	// Approve the adapter to receive the tokens
	tx, err := token.Approve(env.Transactors[l2ChainID], l2BridgeAdapterAddress, amount)
	helpers.PanicErr(err)
	helpers.ConfirmTXMined(context.Background(), env.Clients[l2ChainID], tx, int64(l2ChainID))

	// check the approval
	allowance, err := token.Allowance(nil, env.Transactors[l2ChainID].From, l2BridgeAdapterAddress)
	helpers.PanicErr(err)
	if allowance.Cmp(amount) < 0 {
		panic(fmt.Sprintf("approval failed, allowance: %s, expected amount: %s", allowance.String(), amount.String()))
	}

	// at this point we should be able to withdraw the tokens to L1
	nonce, err := env.Clients[l2ChainID].PendingNonceAt(context.Background(), env.Transactors[l2ChainID].From)
	helpers.PanicErr(err)
	gasPrice, err := env.Clients[l2ChainID].SuggestGasPrice(context.Background())
	helpers.PanicErr(err)

	calldata, err := l2AdapterABI.Pack("sendERC20", l2TokenAddress, common.HexToAddress("0x0"), l1ToAddress, amount, []byte{})
	helpers.PanicErr(err)

	rawTx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      500_000,
		To:       &l2BridgeAdapterAddress,
		Data:     calldata,
	})
	tx, err = env.Transactors[l2ChainID].Signer(env.Transactors[l2ChainID].From, rawTx)
	helpers.PanicErr(err)
	err = env.Clients[l2ChainID].SendTransaction(context.Background(), tx)
	helpers.PanicErr(err)
	// tx, err = l2Adapter.SendERC20(env.Transactors[l2ChainID],
	// 	l2TokenAddress,
	// 	common.HexToAddress("0x0"), // not needed
	// 	l1ToAddress,
	// 	amount,
	// 	[]byte{}, /* bridgeSpecificData, unused for optimism L2 adapter */
	// )
	// helpers.PanicErr(err)
	helpers.ConfirmTXMined(context.Background(), env.Clients[l2ChainID], tx, int64(l2ChainID), "WithdrawFromL2")
}
