package rhea

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip/shared"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/cache_gold_child"
)

func DeployCacheGoldTokenAndPool(t *testing.T, client *EvmDeploymentConfig) {
	if client.ChainConfig.CustomerSettings.CacheGoldFeeAddress == common.HexToAddress("") ||
		client.ChainConfig.CustomerSettings.CacheGoldFeeEnforcer == common.HexToAddress("") {
		client.Logger.Infof("Cannot deploy Cache.gold token because no fee address is set.")
		return
	}

	tokenAddress, tx, _, err := cache_gold_child.DeployCacheGoldChild(client.Owner, client.Client)
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
	client.Logger.Infof("CACHE.gold token instance deployed on %s in tx: %s", tokenAddress.Hex(), helpers.ExplorerLink(int64(client.ChainConfig.EvmChainId), tx.Hash()))

	poolAddress, err := deployBurnMintTokenPool(client, CACHEGOLD, tokenAddress, []common.Address{})
	shared.RequireNoError(t, err)

	cacheGoldToken, err := cache_gold_child.NewCacheGoldChild(tokenAddress, client.Client)
	shared.RequireNoError(t, err)

	tx, err = cacheGoldToken.Initialize(client.Owner, client.ChainConfig.CustomerSettings.CacheGoldFeeAddress, client.ChainConfig.CustomerSettings.CacheGoldFeeEnforcer, poolAddress, client.Owner.From, common.Address{})
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
	client.Logger.Infof("CACHE.gold token initialized in tx: %s", helpers.ExplorerLink(int64(client.ChainConfig.EvmChainId), tx.Hash()))

	tx, err = cacheGoldToken.SetTransferFeeExempt(client.Owner, poolAddress)
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
	client.Logger.Infof("CACHE.gold token pool set fee exempt in tx: %s", helpers.ExplorerLink(int64(client.ChainConfig.EvmChainId), tx.Hash()))
}

func UpdateCacheGoldPool(t *testing.T, client *EvmDeploymentConfig) {
	config := client.ChainConfig.SupportedTokens[CACHEGOLD]

	cacheGoldToken, err := cache_gold_child.NewCacheGoldChild(config.Token, client.Client)
	shared.RequireNoError(t, err)

	tx, err := cacheGoldToken.SetFxManager(client.Owner, config.Pool)
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)

}
