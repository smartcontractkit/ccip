package rhea

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_free_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_subscription_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/shared"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
)

func setOffRampOnTokenPools(t *testing.T, destClient *EvmDeploymentConfig) {
	for _, tokenConfig := range destClient.ChainConfig.SupportedTokens {
		pool, err := native_token_pool.NewNativeTokenPool(tokenConfig.Pool, destClient.Client)
		require.NoError(t, err)

		// Configure offramp address on pool
		tx, err := pool.SetOffRamp(destClient.Owner, destClient.LaneConfig.OffRamp, true)
		require.NoError(t, err)
		shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
		destClient.Logger.Infof("Offramp pool configured with offramp address: %s", helpers.ExplorerLink(destClient.ChainConfig.ChainId.Int64(), tx.Hash()))
	}
}

func setOnRampOnOnRampRouter(t *testing.T, sourceClient *EvmDeploymentConfig, destChainId *big.Int) {
	sourceClient.Logger.Infof("Setting the onRamp on the onRampRouter")
	onRampRouter, err := evm_2_any_subscription_onramp_router.NewEVM2AnySubscriptionOnRampRouter(sourceClient.ChainConfig.OnRampRouter, sourceClient.Client)
	require.NoError(t, err)
	sourceClient.Logger.Infof("Registering new onRamp")
	tx, err := onRampRouter.SetOnRamp(sourceClient.Owner, destChainId, sourceClient.LaneConfig.OnRamp)
	require.NoError(t, err)
	shared.WaitForMined(t, sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
}

func setOnRampOnTokenPools(t *testing.T, sourceClient *EvmDeploymentConfig) {
	for _, tokenConfig := range sourceClient.ChainConfig.SupportedTokens {
		pool, err := native_token_pool.NewNativeTokenPool(tokenConfig.Pool, sourceClient.Client)
		require.NoError(t, err)

		// Configure offramp address on pool
		tx, err := pool.SetOnRamp(sourceClient.Owner, sourceClient.LaneConfig.OnRamp, true)
		require.NoError(t, err)
		shared.WaitForMined(t, sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
		sourceClient.Logger.Infof("Onramp pool configured with offramp address: %s", helpers.ExplorerLink(sourceClient.ChainConfig.ChainId.Int64(), tx.Hash()))
	}
}

func setOnRampOnCommitStore(t *testing.T, sourceClient *EvmDeploymentConfig, destClient *EvmDeploymentConfig) {
	commitStore, err := commit_store.NewCommitStore(destClient.LaneConfig.CommitStore, destClient.Client)
	require.NoError(t, err)

	config, err := commitStore.GetConfig(&bind.CallOpts{})
	require.NoError(t, err)

	config.OnRamps = append(config.OnRamps, sourceClient.LaneConfig.OnRamp)
	config.MinSeqNrByOnRamp = append(config.MinSeqNrByOnRamp, 1)

	tx, err := commitStore.SetConfig(destClient.Owner, config)
	require.NoError(t, err)
	destClient.Logger.Infof(fmt.Sprintf("Adding new onRamp to commitStore in tx %s", helpers.ExplorerLink(destClient.ChainConfig.ChainId.Int64(), tx.Hash())))
	shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
}

func setOffRampRouterOnOffRamp(t *testing.T, destClient *EvmDeploymentConfig) {
	offRamp, err := any_2_evm_free_offramp.NewEVM2EVMFreeOffRamp(destClient.LaneConfig.OffRamp, destClient.Client)
	require.NoError(t, err)

	tx, err := offRamp.SetRouter(destClient.Owner, destClient.ChainConfig.OffRampRouter)
	require.NoError(t, err)
	shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
	destClient.Logger.Infof(fmt.Sprintf("OffRampRouter set on offRamp in tx %s", helpers.ExplorerLink(destClient.ChainConfig.ChainId.Int64(), tx.Hash())))
}

func setOffRampOnOffRampRouter(t *testing.T, destClient *EvmDeploymentConfig) {
	destClient.Logger.Infof("Setting the offRamp on the offRampRouter")
	offRampRouter, err := any_2_evm_subscription_offramp_router.NewAny2EVMSubscriptionOffRampRouter(destClient.ChainConfig.OffRampRouter, destClient.Client)
	require.NoError(t, err)

	tx, err := offRampRouter.AddOffRamp(destClient.Owner, destClient.LaneConfig.OffRamp)
	require.NoError(t, err)
	shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
}

/*
func revokeOffRampOnOffRampRouter(t *testing.T, destClient *EvmDeploymentConfig, offRamp common.Address) {
	destClient.Logger.Infof("Revoking the offRamp on the offRampRouter")
	offRampRouter, err := any_2_evm_subscription_offramp_router.NewAny2EVMSubscriptionOffRampRouter(destClient.ChainConfig.OffRampRouter, destClient.Client)
	require.NoError(t, err)

	tx, err := offRampRouter.RemoveOffRamp(destClient.Owner, offRamp)
	require.NoError(t, err)
	shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
}
*/

/*
func revokeOffRampOnTokenPools(t *testing.T, destClient *EvmDeploymentConfig, offRamp common.Address) {
	// TODO
}
*/

func createDestSubscription(t *testing.T, client *EvmDeploymentConfig, receiver common.Address, allowedSenders []common.Address) {
	offRampRouter, err := any_2_evm_subscription_offramp_router.NewAny2EVMSubscriptionOffRampRouter(client.ChainConfig.OffRampRouter, client.Client)
	require.NoError(t, err)

	fundingAmount := big.NewInt(1e18)

	linkToken, err := link_token_interface.NewLinkToken(client.ChainConfig.LinkToken, client.Client)
	require.NoError(t, err)

	tx, err := linkToken.Approve(client.Owner, client.ChainConfig.OffRampRouter, fundingAmount)
	require.NoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	client.Logger.Infof(fmt.Sprintf("Approved link for offramp subscription funding in tx %s", helpers.ExplorerLink(client.ChainConfig.ChainId.Int64(), tx.Hash())))

	tx, err = offRampRouter.CreateSubscription(client.Owner, any_2_evm_subscription_offramp_router.SubscriptionInterfaceOffRampSubscription{
		Senders:          allowedSenders,
		Receiver:         receiver,
		StrictSequencing: false,
		Balance:          fundingAmount,
	})
	require.NoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	client.Logger.Infof(fmt.Sprintf("Created offramp subscription in tx %s", helpers.ExplorerLink(client.ChainConfig.ChainId.Int64(), tx.Hash())))
}

func fillPoolWithTokens(t *testing.T, client *EvmDeploymentConfig, pool *native_token_pool.NativeTokenPool) {
	destLinkToken, err := link_token_interface.NewLinkToken(client.ChainConfig.LinkToken, client.Client)
	require.NoError(t, err)

	// fill offramp token pool with 0.5 LINK
	amount := big.NewInt(5e17)
	tx, err := destLinkToken.Transfer(client.Owner, pool.Address(), amount)
	require.NoError(t, err)
	client.Logger.Infof("Transferring token to token pool: %s", helpers.ExplorerLink(client.ChainConfig.ChainId.Int64(), tx.Hash()))
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)

	client.Logger.Infof("Locking tokens in pool")
	tx, err = pool.LockOrBurn(client.Owner, amount)
	require.NoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	client.Logger.Infof("Pool filled with tokens: %s", helpers.ExplorerLink(client.ChainConfig.ChainId.Int64(), tx.Hash()))
}
