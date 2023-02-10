package rhea

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/fee_manager"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/router"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/shared"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
)

func setOffRampOnTokenPools(t *testing.T, destClient *EvmDeploymentConfig) {
	for _, tokenConfig := range destClient.ChainConfig.SupportedTokens {
		pool, err := lock_release_token_pool.NewLockReleaseTokenPool(tokenConfig.Pool, destClient.Client)
		shared.RequireNoError(t, err)

		// Configure offramp address on pool
		tx, err := pool.SetOffRamp(destClient.Owner, destClient.LaneConfig.OffRamp, true)
		shared.RequireNoError(t, err)
		shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
		destClient.Logger.Infof("Offramp pool configured with offramp address: %s", helpers.ExplorerLink(int64(destClient.ChainConfig.ChainId), tx.Hash()))
	}
}

func setFeeManagerPrices(t *testing.T, client *EvmDeploymentConfig, destChainId uint64) {
	feeManager, err := fee_manager.NewFeeManager(client.ChainConfig.FeeManager, client.Client)
	shared.RequireNoError(t, err)

	for _, feeToken := range client.ChainConfig.FeeTokens {
		tx, err := feeManager.UpdateFees(client.Owner, []fee_manager.InternalFeeUpdate{
			{
				SourceFeeToken:              client.ChainConfig.SupportedTokens[feeToken].Token,
				DestChainId:                 destChainId,
				FeeTokenBaseUnitsPerUnitGas: big.NewInt(1e18),
			},
		})
		shared.RequireNoError(t, err)
		shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	}
}

func setOnRampOnRouter(t *testing.T, sourceClient *EvmDeploymentConfig, destChainId uint64) {
	sourceClient.Logger.Infof("Setting the onRamp on the Router")
	router, err := router.NewRouter(sourceClient.ChainConfig.Router, sourceClient.Client)
	shared.RequireNoError(t, err)
	sourceClient.Logger.Infof("Registering new onRamp")
	tx, err := router.SetOnRamp(sourceClient.Owner, destChainId, sourceClient.LaneConfig.OnRamp)
	shared.RequireNoError(t, err)
	shared.WaitForMined(t, sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
}

func setOnRampOnTokenPools(t *testing.T, sourceClient *EvmDeploymentConfig) {
	for _, tokenConfig := range sourceClient.ChainConfig.SupportedTokens {
		pool, err := lock_release_token_pool.NewLockReleaseTokenPool(tokenConfig.Pool, sourceClient.Client)
		shared.RequireNoError(t, err)

		// Configure offramp address on pool
		tx, err := pool.SetOnRamp(sourceClient.Owner, sourceClient.LaneConfig.OnRamp, true)
		shared.RequireNoError(t, err)
		shared.WaitForMined(t, sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
		sourceClient.Logger.Infof("Onramp pool configured with offramp address: %s", helpers.ExplorerLink(int64(sourceClient.ChainConfig.ChainId), tx.Hash()))
	}
}

func setRouterOnOffRamp(t *testing.T, destClient *EvmDeploymentConfig) {
	offRamp, err := evm_2_evm_offramp.NewEVM2EVMOffRamp(destClient.LaneConfig.OffRamp, destClient.Client)
	shared.RequireNoError(t, err)
	tx, err := offRamp.SetRouter(destClient.Owner, destClient.ChainConfig.Router)
	shared.RequireNoError(t, err)
	shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
	destClient.Logger.Infof(fmt.Sprintf("Router set on offRamp in tx %s", helpers.ExplorerLink(int64(destClient.ChainConfig.ChainId), tx.Hash())))
}

func setOffRampOnRouter(t *testing.T, client *EvmDeploymentConfig) {
	client.Logger.Infof("Setting the offRamp on the Router")
	router, err := router.NewRouter(client.ChainConfig.Router, client.Client)
	shared.RequireNoError(t, err)

	isOffRamp, err := router.IsOffRamp(&bind.CallOpts{}, client.LaneConfig.OffRamp)
	shared.RequireNoError(t, err)
	if isOffRamp {
		client.Logger.Infof("OffRamp already configured on router. Skipping")
		return
	}

	tx, err := router.AddOffRamp(client.Owner, client.LaneConfig.OffRamp)
	shared.RequireNoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
}

func setFeeManagerUpdater(t *testing.T, client *EvmDeploymentConfig) {
	feeManager, err := fee_manager.NewFeeManager(client.ChainConfig.FeeManager, client.Client)
	shared.RequireNoError(t, err)

	tx, err := feeManager.SetFeeUpdater(client.Owner, client.LaneConfig.OffRamp)
	shared.RequireNoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
}

/*
func revokeOffRampOnOffRampRouter(t *testing.T, destClient *EvmDeploymentConfig, offRamp common.Address) {
	destClient.Logger.Infof("Revoking the offRamp on the offRampRouter")
	offRampRouter, err := any_2_evm_subscription_offramp_router.NewAny2EVMSubscriptionOffRampRouter(destClient.ChainConfig.OffRampRouter, destClient.Client)
	shared.RequireNoError(t, err)

	tx, err := offRampRouter.RemoveOffRamp(destClient.Owner, offRamp)
	shared.RequireNoError(t, err)
	shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
}
*/

/*
func revokeOffRampOnTokenPools(t *testing.T, destClient *EvmDeploymentConfig, offRamp common.Address) {
	// TODO
}
*/

func fillPoolWithTokens(t *testing.T, client *EvmDeploymentConfig, pool *lock_release_token_pool.LockReleaseTokenPool, tokenAddress common.Address) {
	token, err := link_token_interface.NewLinkToken(tokenAddress, client.Client)
	shared.RequireNoError(t, err)

	// fill offramp token pool with 0.5 token
	amount := big.NewInt(5e17)
	tx, err := token.Approve(client.Owner, pool.Address(), amount)
	shared.RequireNoError(t, err)
	client.Logger.Infof("Approving token to the token pool: %s", helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash()))
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)

	tx, err = pool.AddLiquidity(client.Owner, amount)
	shared.RequireNoError(t, err)
	client.Logger.Infof("Adding liquidity to the token pool: %s", helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash()))
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)

	client.Logger.Infof("Pool filled with tokens: %s", helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash()))
}
