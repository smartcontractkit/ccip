package rhea

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/shared"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/router"
)

func setOffRampOnTokenPools(t *testing.T, destClient *EvmDeploymentConfig) {
	for _, tokenConfig := range destClient.ChainConfig.SupportedTokens {
		pool, err := lock_release_token_pool.NewLockReleaseTokenPool(tokenConfig.Pool, destClient.Client)
		shared.RequireNoError(t, err)

		rampUpdate := lock_release_token_pool.TokenPoolRampUpdate{
			Ramp:    destClient.LaneConfig.OffRamp,
			Allowed: true,
		}

		// Configure offramp address on pool
		tx, err := pool.ApplyRampUpdates(destClient.Owner, []lock_release_token_pool.TokenPoolRampUpdate{}, []lock_release_token_pool.TokenPoolRampUpdate{rampUpdate})
		shared.RequireNoError(t, err)
		err = shared.WaitForMined(destClient.Logger, destClient.Client, tx.Hash(), true)
		shared.RequireNoError(t, err)
		destClient.Logger.Infof("Offramp pool configured with offramp address: %s", helpers.ExplorerLink(int64(destClient.ChainConfig.EvmChainId), tx.Hash()))
	}
}

func SetPriceRegistryPrices(t *testing.T, client *EvmDeploymentConfig, destChainSelector uint64) {
	priceRegistry, err := price_registry.NewPriceRegistry(client.ChainConfig.PriceRegistry, client.Client)
	shared.RequireNoError(t, err)

	priceUpdates := price_registry.InternalPriceUpdates{
		TokenPriceUpdates: []price_registry.InternalTokenPriceUpdate{},
		DestChainId:       destChainSelector,
		// Set 1e18 units of gas to $2k, being fairly reasonable for eth
		// These values will get auto updated by the DON
		UsdPerUnitGas: big.NewInt(2000e9), // $2000 per eth * 1gwei = 2000e9
	}

	for _, tokenConfig := range client.ChainConfig.SupportedTokens {
		priceUpdates.TokenPriceUpdates = append(priceUpdates.TokenPriceUpdates, price_registry.InternalTokenPriceUpdate{
			SourceToken: tokenConfig.Token,
			UsdPerToken: tokenConfig.Price,
		})
	}

	tx, err := priceRegistry.UpdatePrices(client.Owner, priceUpdates)
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
}

func setOnRampOnRouter(t *testing.T, sourceClient *EvmDeploymentConfig, destChainSelector uint64) {
	sourceClient.Logger.Infof("Setting the onRamp on the Router")
	routerContract, err := router.NewRouter(sourceClient.ChainConfig.Router, sourceClient.Client)
	shared.RequireNoError(t, err)
	sourceClient.Logger.Infof("Registering new onRamp")
	tx, err := routerContract.ApplyRampUpdates(sourceClient.Owner, []router.RouterOnRampUpdate{{DestChainId: destChainSelector, OnRamp: sourceClient.LaneConfig.OnRamp}}, nil)
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
}

func setOnRampOnTokenPools(t *testing.T, sourceClient *EvmDeploymentConfig) {
	for _, tokenConfig := range sourceClient.ChainConfig.SupportedTokens {
		pool, err := lock_release_token_pool.NewLockReleaseTokenPool(tokenConfig.Pool, sourceClient.Client)
		shared.RequireNoError(t, err)

		rampUpdate := lock_release_token_pool.TokenPoolRampUpdate{
			Ramp:    sourceClient.LaneConfig.OnRamp,
			Allowed: true,
		}

		// Configure offramp address on pool
		tx, err := pool.ApplyRampUpdates(sourceClient.Owner, []lock_release_token_pool.TokenPoolRampUpdate{rampUpdate}, []lock_release_token_pool.TokenPoolRampUpdate{})
		shared.RequireNoError(t, err)
		err = shared.WaitForMined(sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
		shared.RequireNoError(t, err)
		sourceClient.Logger.Infof("Onramp pool configured with offramp address: %s", helpers.ExplorerLink(int64(sourceClient.ChainConfig.EvmChainId), tx.Hash()))
	}
}

func setOffRampOnRouter(t *testing.T, sourceChainSelector uint64, client *EvmDeploymentConfig) {
	client.Logger.Infof("Setting the offRamp on the Router")
	routerContract, err := router.NewRouter(client.ChainConfig.Router, client.Client)
	shared.RequireNoError(t, err)

	offRamps, err := routerContract.GetOffRamps(&bind.CallOpts{}, sourceChainSelector)
	shared.RequireNoError(t, err)
	for _, offRamp := range offRamps {
		if offRamp == client.LaneConfig.OffRamp {
			client.Logger.Infof("OffRamp already configured on router. Skipping")
			return
		}
	}

	tx, err := routerContract.ApplyRampUpdates(client.Owner, nil, []router.RouterOffRampUpdate{
		{SourceChainId: sourceChainSelector, OffRamps: []common.Address{client.LaneConfig.OffRamp}}})
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
}

func setPriceRegistryUpdater(t *testing.T, client *EvmDeploymentConfig) {
	priceRegistry, err := price_registry.NewPriceRegistry(client.ChainConfig.PriceRegistry, client.Client)
	shared.RequireNoError(t, err)

	tx, err := priceRegistry.ApplyPriceUpdatersUpdates(client.Owner, []common.Address{client.LaneConfig.CommitStore}, []common.Address{})
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
}

func fillPoolWithTokens(client *EvmDeploymentConfig, pool *lock_release_token_pool.LockReleaseTokenPool, tokenAddress common.Address, tokenName Token) error {
	token, err := link_token_interface.NewLinkToken(tokenAddress, client.Client)
	if err != nil {
		return err
	}

	// fill offramp token pool with 0.5 token
	amount := new(big.Int).Div(tokenName.Multiplier(), big.NewInt(2))
	tx, err := token.Approve(client.Owner, pool.Address(), amount)
	if err != nil {
		return err
	}
	client.Logger.Infof("Approving token to the token pool: %s", helpers.ExplorerLink(int64(client.ChainConfig.EvmChainId), tx.Hash()))
	if err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true); err != nil {
		return err
	}
	tx, err = pool.AddLiquidity(client.Owner, amount)
	if err != nil {
		return err
	}
	client.Logger.Infof("Adding liquidity to the token pool: %s", helpers.ExplorerLink(int64(client.ChainConfig.EvmChainId), tx.Hash()))
	if err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true); err != nil {
		return err
	}
	client.Logger.Infof("Pool filled with tokens: %s", helpers.ExplorerLink(int64(client.ChainConfig.EvmChainId), tx.Hash()))
	return nil
}

func FundPingPong(client *EvmDeploymentConfig, fundingAmount *big.Int, tokenAddress common.Address) error {
	linkToken, err := link_token_interface.NewLinkToken(tokenAddress, client.Client)
	if err != nil {
		return err
	}

	tx, err := linkToken.Transfer(client.Owner, client.LaneConfig.PingPongDapp, fundingAmount)
	if err != nil {
		return err
	}
	if err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true); err != nil {
		return err
	}
	client.Logger.Infof("Ping pong funded with %s in tx: %s", fundingAmount.String(), helpers.ExplorerLink(int64(client.ChainConfig.EvmChainId), tx.Hash()))
	return nil
}

func UpdateDeployedAt(t *testing.T, source *EvmDeploymentConfig, dest *EvmDeploymentConfig) {
	sourceBlock, err := source.Client.BlockNumber(context.Background())
	require.NoError(t, err)

	source.ChainConfig.DeploySettings.DeployedAtBlock = sourceBlock
	source.LaneConfig.DeploySettings.DeployedAtBlock = sourceBlock

	destBlock, err := dest.Client.BlockNumber(context.Background())
	require.NoError(t, err)

	dest.ChainConfig.DeploySettings.DeployedAtBlock = destBlock
	dest.LaneConfig.DeploySettings.DeployedAtBlock = destBlock
}
