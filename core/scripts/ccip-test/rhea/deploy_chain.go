package rhea

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/mock_afn_contract"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/price_registry"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/router"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/shared"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
)

func DeployToNewChain(t *testing.T, client *EvmDeploymentConfig) {
	// Updates client.AFN if any new contracts are deployed
	deployAFN(t, client)
	// Updates client.TokenPools if any new contracts are deployed
	deployLockReleaseTokenPool(t, client)
	// Updates client.ChainConfig.Router if any new contracts are deployed
	deployRouter(t, client)
	// Update client.PriceRegistry if any new contracts are deployed
	deployPriceRegistry(t, client)
}

func deployAFN(t *testing.T, client *EvmDeploymentConfig) {
	if !client.DeploySettings.DeployAFN {
		if client.ChainConfig.Afn.Hex() == "0x0000000000000000000000000000000000000000" {
			t.Error("deploy new afn set to false but no afn given in config")
		}
		client.Logger.Infof("Skipping AFN deployment, using AFN on %s", client.ChainConfig.Afn)
		return
	}

	client.Logger.Infof("Deploying AFN")
	address, tx, _, err := mock_afn_contract.DeployMockAFNContract(client.Owner, client.Client)
	shared.RequireNoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	client.Logger.Infof("AFN deployed on %s in tx: %s", address.Hex(), helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash()))
	client.ChainConfig.Afn = address
}

func deployLockReleaseTokenPool(t *testing.T, client *EvmDeploymentConfig) {
	for tokenName, tokenConfig := range client.ChainConfig.SupportedTokens {
		if client.DeploySettings.DeployTokenPools {
			client.Logger.Infof("Deploying token pool for %s token", tokenName)
			tokenPoolAddress, tx, _, err := lock_release_token_pool.DeployLockReleaseTokenPool(client.Owner, client.Client, tokenConfig.Token)
			shared.RequireNoError(t, err)
			shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
			client.Logger.Infof("Native token pool deployed on %s in tx %s", tokenPoolAddress, helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash()))
			pool, err := lock_release_token_pool.NewLockReleaseTokenPool(tokenPoolAddress, client.Client)
			shared.RequireNoError(t, err)
			fillPoolWithTokens(t, client, pool, tokenConfig.Token)
			client.ChainConfig.SupportedTokens[tokenName] = EVMBridgedToken{
				Token:                tokenConfig.Token,
				Pool:                 tokenPoolAddress,
				Price:                big.NewInt(1),
				PriceFeedsAggregator: tokenConfig.PriceFeedsAggregator,
			}
		} else {
			if tokenConfig.Pool.Hex() == "0x0000000000000000000000000000000000000000" {
				t.Error("deploy new lock unlock pool set to false but no lock unlock pool given in config")
			}
			pool, err := lock_release_token_pool.NewLockReleaseTokenPool(tokenConfig.Pool, client.Client)
			shared.RequireNoError(t, err)
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
	nativeFeeToken := common.HexToAddress("0x0")
	if client.ChainConfig.WrappedNative != "" {
		nativeFeeToken = client.ChainConfig.SupportedTokens[client.ChainConfig.WrappedNative].Token
	}

	routerAddress, tx, _, err := router.DeployRouter(client.Owner, client.Client, nativeFeeToken)
	shared.RequireNoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	client.ChainConfig.Router = routerAddress

	client.Logger.Infof(fmt.Sprintf("Router deployed on %s in tx %s", routerAddress.String(), helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash())))
}

// deployPriceRegistry Prices is deployed without any feeUpdaters
func deployPriceRegistry(t *testing.T, client *EvmDeploymentConfig) {
	if !client.DeploySettings.DeployPriceRegistry {
		client.Logger.Infof("Skipping PriceRegistry deployment, using Prices on %s", client.ChainConfig.PriceRegistry)
		return
	}

	feeTokens := make([]common.Address, len(client.ChainConfig.FeeTokens))
	for i, token := range client.ChainConfig.FeeTokens {
		feeTokens[i] = client.ChainConfig.SupportedTokens[token].Token
	}

	client.Logger.Infof("Deploying PriceRegistry")
	priceRegistry, tx, _, err := price_registry.DeployPriceRegistry(
		client.Owner,
		client.Client,
		price_registry.InternalPriceUpdates{
			// No updates needed, these should all be done by the DON upon
			// starting OCR.
			TokenPriceUpdates: []price_registry.InternalTokenPriceUpdate{},
			// 0 signals that the UsdPerUnitGas should not be used and this is
			// not an update for gas fee prices.
			DestChainId:   0,
			UsdPerUnitGas: big.NewInt(0),
		},
		[]common.Address{},
		feeTokens,
		60*60*24*14, // two weeks
	)
	shared.RequireNoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	client.ChainConfig.PriceRegistry = priceRegistry

	client.Logger.Infof(fmt.Sprintf("PriceRegistry deployed on %s in tx %s", priceRegistry.String(), helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash())))
}
