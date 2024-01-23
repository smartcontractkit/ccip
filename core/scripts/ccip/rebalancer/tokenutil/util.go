package tokenutil

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip/rebalancer/multienv"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/weth9"
	"go.uber.org/zap"
)

func DepositWETHOnChain(env multienv.Env, chainID uint64, wethAddress common.Address, amount *big.Int) {
	zap.L().Info("depositing WETH",
		zap.Uint64("chainID", chainID),
		zap.String("wethAddress", wethAddress.Hex()),
		zap.String("amount", amount.String()))
	weth, err := weth9.NewWETH9(wethAddress, env.Clients[chainID])
	helpers.PanicErr(err)

	tx, err := weth.Deposit(&bind.TransactOpts{
		From:   env.Transactors[chainID].From,
		Signer: env.Transactors[chainID].Signer,
		Value:  amount,
	})
	helpers.PanicErr(err)
	helpers.ConfirmTXMined(context.Background(), env.Clients[chainID], tx, int64(chainID),
		"depositing", amount.String(), "wei to", wethAddress.Hex())
}
