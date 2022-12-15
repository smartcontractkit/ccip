package rhea

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/gas_fee_cache"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/ge_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/mock_afn_contract"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/shared"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
)

func DeployToNewChain(t *testing.T, client *EvmDeploymentConfig) {
	// Updates client.AFN if any new contracts are deployed
	deployAFN(t, client)
	// Updates client.TokenPools if any new contracts are deployed
	deployNativeTokenPool(t, client)
	// Updates client.ChainConfig.Router if any new contracts are deployed
	deployRouter(t, client)
	// Update client.GasFeeCache if any new contracts are deployed
	deployGasFeeCache(t, client)
}

func deployAFN(t *testing.T, client *EvmDeploymentConfig) {
	if !client.DeploySettings.DeployAFN {
		if client.ChainConfig.Afn.Hex() == "0x0000000000000000000000000000000000000000" {
			t.Error("deploy new afn set to false but no afn given in config")
		}
		client.Logger.Infof("Skipping AFN deployment, using AFN on %s", client.ChainConfig.Afn)
	}

	client.Logger.Infof("Deploying AFN")
	address, tx, _, err := mock_afn_contract.DeployMockAFNContract(client.Owner, client.Client)
	require.NoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	client.Logger.Infof("AFN deployed on %s in tx: %s", address.Hex(), helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash()))
	client.ChainConfig.Afn = address
}

func deployNativeTokenPool(t *testing.T, client *EvmDeploymentConfig) {
	for tokenName, tokenConfig := range client.ChainConfig.SupportedTokens {
		if client.DeploySettings.DeployTokenPools {
			client.Logger.Infof("Deploying token pool for %s token", tokenName)
			tokenPoolAddress, tx, _, err := native_token_pool.DeployNativeTokenPool(client.Owner, client.Client, tokenConfig.Token)
			require.NoError(t, err)
			shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
			client.Logger.Infof("Native token pool deployed on %s in tx %s", tokenPoolAddress, helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash()))
			pool, err := native_token_pool.NewNativeTokenPool(tokenPoolAddress, client.Client)
			require.NoError(t, err)
			fillPoolWithTokens(t, client, pool)
			client.ChainConfig.SupportedTokens[tokenName] = EVMBridgedToken{
				Pool:  tokenPoolAddress,
				Price: big.NewInt(1),
			}
		} else {
			if tokenConfig.Pool.Hex() == "0x0000000000000000000000000000000000000000" {
				t.Error("deploy new lock unlock pool set to false but no lock unlock pool given in config")
			}
			pool, err := native_token_pool.NewNativeTokenPool(tokenConfig.Pool, client.Client)
			require.NoError(t, err)
			client.Logger.Infof("Skipping Pool deployment, using Pool on %s", pool.Address().Hex())
		}
	}
}

// deployRouter always uses an empty list of offRamps. Ramps should be set in the offRamp deployment step.
func deployRouter(t *testing.T, client *EvmDeploymentConfig) {
	if !client.DeploySettings.DeployRouter {
		client.Logger.Infof("Skipping Router deployment, using Router on %s", client.ChainConfig.Router)
		return
	}

	client.Logger.Infof("Deploying Router")
	routerAddress, tx, _, err := ge_router.DeployGERouter(client.Owner, client.Client, []common.Address{})
	require.NoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	client.ChainConfig.Router = routerAddress

	client.Logger.Infof(fmt.Sprintf("Router deployed on %s in tx %s", routerAddress.String(), helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash())))
}

// deployGasFeeCache GasFeeCache is deployed without any feeUpdaters
func deployGasFeeCache(t *testing.T, client *EvmDeploymentConfig) {
	if !client.DeploySettings.DeployGasFeeCache {
		client.Logger.Infof("Skipping GasFeeCache deployment, using GasFeeCache on %s", client.ChainConfig.GasFeeCache)
		return
	}

	client.Logger.Infof("Deploying GasFeeCache")
	gasFeeCache, tx, _, err := gas_fee_cache.DeployGasFeeCache(
		client.Owner,
		client.Client,
		[]gas_fee_cache.GEFeeUpdate{},
		[]common.Address{},
		big.NewInt(1e18),
	)
	require.NoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	client.ChainConfig.GasFeeCache = gasFeeCache

	client.Logger.Infof(fmt.Sprintf("GasFeeCache deployed on %s in tx %s", gasFeeCache.String(), helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash())))
}
