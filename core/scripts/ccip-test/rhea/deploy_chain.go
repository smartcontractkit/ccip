package rhea

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/shared"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/burn_mint_erc677"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/legacy_burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/mock_afn_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/router"
)

func DeployToNewChain(client *EvmDeploymentConfig) error {
	// Updates client.AFN if any new contracts are deployed
	err := deployAFN(client)
	if err != nil {
		return errors.Wrap(err, "afn deployment failed")
	}
	// Updates client.TokenPools if any new contracts are deployed
	err = DeployTokenPools(client)
	if err != nil {
		return errors.Wrap(err, "pool deployment failed")
	}
	// Updates client.ChainConfig.Router if any new contracts are deployed
	err = deployRouter(client)
	if err != nil {
		return errors.Wrap(err, "router deployment failed")
	}
	// Updates client.ChainConfig.UpgradeRouter if any new contracts are deployed
	err = deployUpgradeRouter(client)
	if err != nil {
		return errors.Wrap(err, "upgrade router deployment failed")
	}
	// Update client.PriceRegistry if any new contracts are deployed
	err = deployPriceRegistry(client)
	if err != nil {
		return errors.Wrap(err, "price registry deployment failed")
	}
	return nil
}

func DeployUpgradeRouters(source *EvmDeploymentConfig, dest *EvmDeploymentConfig) error {
	err := deployUpgradeRouter(source)
	if err != nil {
		return errors.Wrap(err, "upgrade router in source chain deployment failed")
	}
	err = deployUpgradeRouter(dest)
	if err != nil {
		return errors.Wrap(err, "upgrade router in dest chain deployment failed")
	}
	return nil
}

func deployAFN(client *EvmDeploymentConfig) error {
	if !client.ChainConfig.DeploySettings.DeployAFN {
		if client.ChainConfig.Afn.Hex() == "0x0000000000000000000000000000000000000000" {
			return fmt.Errorf("deploy new afn set to false but no afn given in config")
		}
		client.Logger.Infof("Skipping AFN deployment, using AFN on %s", client.ChainConfig.Afn)
		return nil
	}

	client.Logger.Infof("Deploying AFN")
	address, tx, _, err := mock_afn_contract.DeployMockAFNContract(client.Owner, client.Client)
	if err != nil {
		return err
	}
	if err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true); err != nil {
		return err
	}
	client.Logger.Infof("AFN deployed on %s in tx: %s", address.Hex(), helpers.ExplorerLink(int64(client.ChainConfig.EvmChainId), tx.Hash()))
	client.ChainConfig.Afn = address
	return nil
}

func DeployTokenPools(client *EvmDeploymentConfig) error {
	for tokenName, tokenConfig := range client.ChainConfig.SupportedTokens {
		if err := deployPool(client, tokenName, tokenConfig); err != nil {
			return errors.Wrapf(err, "failed %s", tokenName)
		}
	}
	return nil
}

func deployPool(client *EvmDeploymentConfig, tokenName Token, tokenConfig EVMBridgedToken) error {
	if tokenConfig.TokenPoolType == FeeTokenOnly {
		client.Logger.Infof("Skipping pool deployment for fee only token")
		return nil
	}
	// Only deploy a new pool if there is no current pool address given
	// and the deploySetting indicate a new pool should be deployed.
	if client.ChainConfig.DeploySettings.DeployTokenPools && tokenConfig.Pool == common.HexToAddress("") {
		client.Logger.Infof("Deploying token pool for %s token", tokenName)
		var poolAddress, tokenAddress common.Address
		var err error
		switch tokenConfig.TokenPoolType {
		case LockRelease:
			poolAddress, err = deployLockReleaseTokenPool(client, tokenName, tokenConfig.Token, tokenConfig.PoolAllowList)
		case BurnMint:
			poolAddress, err = deployBurnMintTokenPool(client, tokenName, tokenConfig.Token, tokenConfig.PoolAllowList)
		case Legacy:
			poolAddress, err = deployLegacyBurnMintTokenPool(client, tokenName, tokenConfig.Token, tokenConfig.PoolAllowList)
		case Wrapped:
			tokenAddress, poolAddress, err = deployWrappedTokenPool(client, tokenName, tokenConfig.PoolAllowList)
			// Since we also deployed the token we need to set it
			tokenConfig.Token = tokenAddress
		default:
			return fmt.Errorf("unknown pool type %s", tokenConfig.TokenPoolType)
		}
		if err != nil {
			return err
		}
		client.ChainConfig.SupportedTokens[tokenName] = EVMBridgedToken{
			Token:         tokenConfig.Token,
			Pool:          poolAddress,
			Price:         tokenConfig.Price,
			TokenPoolType: tokenConfig.TokenPoolType,
			PoolAllowList: tokenConfig.PoolAllowList,
		}
		return nil
	}

	// If no pools should be deployed but there is no pool address set fail.
	if tokenConfig.Pool == common.HexToAddress("") {
		return fmt.Errorf("deploy new %s pool set to false but no %s pool given in config", tokenName, tokenConfig.TokenPoolType)
	}
	client.Logger.Infof("Skipping %s Pool deployment, using Pool on %s", tokenName, tokenConfig.Pool)

	return setPoolAllowList(client, tokenConfig.Pool, tokenConfig.PoolAllowList, tokenName)
}

func deployLockReleaseTokenPool(client *EvmDeploymentConfig, tokenName Token, tokenAddress common.Address, poolAllowList []common.Address) (common.Address, error) {
	tokenPoolAddress, tx, _, err := lock_release_token_pool.DeployLockReleaseTokenPool(
		client.Owner,
		client.Client,
		tokenAddress,
		poolAllowList,
		lock_release_token_pool.RateLimiterConfig{
			IsEnabled: false,
			Capacity:  new(big.Int).Mul(tokenName.Multiplier(), big.NewInt(1e9)),
			Rate:      new(big.Int).Mul(tokenName.Multiplier(), big.NewInt(1e5)),
		})
	if err != nil {
		return common.Address{}, err
	}
	if err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true); err != nil {
		return common.Address{}, err
	}
	client.Logger.Infof("Lock/release pool for %s deployed on %s in tx %s", tokenName, tokenPoolAddress, helpers.ExplorerLink(int64(client.ChainConfig.EvmChainId), tx.Hash()))
	pool, err := lock_release_token_pool.NewLockReleaseTokenPool(tokenPoolAddress, client.Client)
	if err != nil {
		return common.Address{}, err
	}
	err = fillPoolWithTokens(client, pool, tokenAddress, tokenName)
	return tokenPoolAddress, err
}

func deployLegacyBurnMintTokenPool(client *EvmDeploymentConfig, tokenName Token, tokenAddress common.Address, poolAllowList []common.Address) (common.Address, error) {
	client.Logger.Infof("Deploying token pool for %s token", tokenName)
	tokenPoolAddress, tx, _, err := legacy_burn_mint_token_pool.DeployLegacyBurnMintTokenPool(
		client.Owner,
		client.Client,
		tokenAddress,
		poolAllowList,
		legacy_burn_mint_token_pool.RateLimiterConfig{
			IsEnabled: false,
			Capacity:  new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e9)),
			Rate:      new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e5)),
		})
	if err != nil {
		return common.Address{}, err
	}
	if err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true); err != nil {
		return common.Address{}, err
	}
	client.Logger.Infof("Legacy burn/mint pool for %s deployed on %s in tx %s", tokenName, tokenPoolAddress, helpers.ExplorerLink(int64(client.ChainConfig.EvmChainId), tx.Hash()))
	return tokenPoolAddress, nil
}

func deployBurnMintTokenPool(client *EvmDeploymentConfig, tokenName Token, tokenAddress common.Address, poolAllowList []common.Address) (common.Address, error) {
	client.Logger.Infof("Deploying token pool for %s token", tokenName)
	tokenPoolAddress, tx, _, err := burn_mint_token_pool.DeployBurnMintTokenPool(
		client.Owner,
		client.Client,
		tokenAddress,
		poolAllowList,
		burn_mint_token_pool.RateLimiterConfig{
			IsEnabled: false,
			Capacity:  new(big.Int).Mul(tokenName.Multiplier(), big.NewInt(1e9)),
			Rate:      new(big.Int).Mul(tokenName.Multiplier(), big.NewInt(1e5)),
		})
	if err != nil {
		return common.Address{}, err
	}
	if err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true); err != nil {
		return common.Address{}, err
	}
	client.Logger.Infof("Burn/mint pool for %s deployed on %s in tx %s", tokenName, tokenPoolAddress, helpers.ExplorerLink(int64(client.ChainConfig.EvmChainId), tx.Hash()))
	return tokenPoolAddress, nil
}

func deployWrappedTokenPool(client *EvmDeploymentConfig, tokenName Token, poolAllowList []common.Address) (common.Address, common.Address, error) {
	client.Logger.Infof("Deploying token pool for %s token", tokenName)
	if tokenName.Symbol() == "" {
		return common.Address{}, common.Address{}, fmt.Errorf("no token symbol given for wrapped token pool %s", tokenName)
	}

	tokenAddress, tx, _, err := burn_mint_erc677.DeployBurnMintERC677(client.Owner, client.Client, string(tokenName), tokenName.Symbol(), tokenName.Decimals())
	if err != nil {
		return common.Address{}, common.Address{}, err
	}
	if err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true); err != nil {
		return common.Address{}, common.Address{}, err
	}
	client.Logger.Infof("New %s token deployed on %s in tx %s", tokenName, tokenAddress, helpers.ExplorerLink(int64(client.ChainConfig.EvmChainId), tx.Hash()))

	poolAddress, err := deployBurnMintTokenPool(client, tokenName, tokenAddress, poolAllowList)
	if err != nil {
		return common.Address{}, common.Address{}, err
	}

	token, err := burn_mint_erc677.NewBurnMintERC677(tokenAddress, client.Client)
	if err != nil {
		return common.Address{}, common.Address{}, err
	}

	tx, err = token.GrantMintAndBurnRoles(client.Owner, poolAddress)
	if err != nil {
		return common.Address{}, common.Address{}, err
	}
	if err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true); err != nil {
		return common.Address{}, common.Address{}, err
	}

	return tokenAddress, poolAddress, nil
}

// deployRouter always uses an empty list of offRamps. Ramps should be set in the offRamp deployment step.
func deployRouter(client *EvmDeploymentConfig) error {
	if !client.ChainConfig.DeploySettings.DeployRouter {
		client.Logger.Infof("Skipping Router deployment, using Router on %s", client.ChainConfig.Router)
		return nil
	}

	client.Logger.Infof("Deploying Router")
	nativeFeeToken := common.Address{}
	if client.ChainConfig.WrappedNative != "" {
		nativeFeeToken = client.ChainConfig.SupportedTokens[client.ChainConfig.WrappedNative].Token
	}

	routerAddress, tx, _, err := router.DeployRouter(client.Owner, client.Client, nativeFeeToken)
	if err != nil {
		return err
	}
	if err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true); err != nil {
		return err
	}
	client.ChainConfig.Router = routerAddress

	client.Logger.Infof(fmt.Sprintf("Router deployed on %s in tx %s", routerAddress.String(), helpers.ExplorerLink(int64(client.ChainConfig.EvmChainId), tx.Hash())))
	return nil
}

// deployUpgradeRouter always uses an empty list of offRamps. Ramps should be set in the offRamp deployment step.
func deployUpgradeRouter(client *EvmDeploymentConfig) error {
	if !client.ChainConfig.DeploySettings.DeployUpgradeRouter {
		client.Logger.Infof("Skipping Upgrade Router deployment, using Router on %s", client.ChainConfig.UpgradeRouter)
		return nil
	}

	client.Logger.Infof("Deploying Router")
	nativeFeeToken := common.Address{}
	if client.ChainConfig.WrappedNative != "" {
		nativeFeeToken = client.ChainConfig.SupportedTokens[client.ChainConfig.WrappedNative].Token
	}

	routerAddress, tx, _, err := router.DeployRouter(client.Owner, client.Client, nativeFeeToken)
	if err != nil {
		return err
	}
	if err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true); err != nil {
		return err
	}
	client.ChainConfig.UpgradeRouter = routerAddress

	client.Logger.Infof(fmt.Sprintf("Router deployed on %s in tx %s", routerAddress.String(), helpers.ExplorerLink(int64(client.ChainConfig.EvmChainId), tx.Hash())))
	return nil
}

// deployPriceRegistry Prices is deployed without any feeUpdaters
func deployPriceRegistry(client *EvmDeploymentConfig) error {
	if !client.ChainConfig.DeploySettings.DeployPriceRegistry {
		client.Logger.Infof("Skipping PriceRegistry deployment, using PriceRegistry on %s", client.ChainConfig.PriceRegistry)
		return nil
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
			DestChainSelector: 0,
			UsdPerUnitGas:     big.NewInt(0),
		},
		[]common.Address{},
		feeTokens,
		60*60*24*14, // two weeks
	)
	if err != nil {
		return err
	}
	if err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true); err != nil {
		return err
	}
	client.ChainConfig.PriceRegistry = priceRegistry

	client.Logger.Infof(fmt.Sprintf("PriceRegistry deployed on %s in tx %s", priceRegistry.String(), helpers.ExplorerLink(int64(client.ChainConfig.EvmChainId), tx.Hash())))
	return nil
}
