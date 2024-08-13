package ccipdeployment

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/ccip_config"

	owner_helpers "github.com/smartcontractkit/ccip-owner-contracts/gethwrappers"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/arm_proxy_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/mock_arm_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/nonce_manager"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/token_admin_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/weth9"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/capabilities_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/shared/generated/burn_mint_erc677"
)

var (
	// 1.0
	ARMProxy_1_1_0      = "ARMProxy 1.0.0"
	MockARM_1_0_0       = "MockARM 1.0.0"
	LinkToken_1_0_0     = "LinkToken 1.0.0"
	WETH9_1_0_0         = "WETH9 1.0.0"
	MCMS_1_0_0          = "ManyChainMultiSig 1.0.0"
	RBAC_Timelock_1_0_0 = "RBACTimelock 1.0.0"

	// 1.2
	Router_1_2_0 = "Router 1.2.0"
	// 1.5
	TokenAdminRegistry_1_5_0 = "TokenAdminRegistry 1.5.0-dev"
	// 1.6
	CapabilitiesRegistry_1_0_0 = "CapabilitiesRegistry 1.0.0"
	CCIPConfig_1_6_0           = "CCIPConfig 1.6.0-dev"
	EVM2EVMMultiOnRamp_1_6_0   = "EVM2EVMMultiOnRamp 1.6.0-dev"
	EVM2EVMMultiOffRamp_1_6_0  = "EVM2EVMMultiOffRamp 1.6.0-dev"
	NonceManager_1_6_0         = "NonceManager 1.6.0-dev"
	PriceRegistry_1_6_0        = "PriceRegistry 1.6.0-dev"
)

type Contracts interface {
	*capabilities_registry.CapabilitiesRegistry |
		*arm_proxy_contract.ARMProxyContract |
		*ccip_config.CCIPConfig |
		*nonce_manager.NonceManager |
		*price_registry.PriceRegistry |
		*router.Router |
		*token_admin_registry.TokenAdminRegistry |
		*weth9.WETH9 |
		*mock_arm_contract.MockARMContract |
		*owner_helpers.ManyChainMultiSig |
		*owner_helpers.RBACTimelock |
		*evm_2_evm_multi_offramp.EVM2EVMMultiOffRamp |
		*evm_2_evm_multi_onramp.EVM2EVMMultiOnRamp |
		*burn_mint_erc677.BurnMintERC677
}

type ContractDeploy[C Contracts] struct {
	// We just return keep all the deploy return values
	// since some will be empty if there's an error.
	// and we want to avoid repeating that
	Address  common.Address
	Contract C
	Tx       *types.Transaction
	TvStr    string
	Err      error
}

// TODO: pull up to general deployment pkg
func deployContract[C Contracts](
	lggr logger.Logger,
	chain deployment.Chain,
	addressBook deployment.AddressBook,
	deploy func(chain deployment.Chain) ContractDeploy[C],
) (*ContractDeploy[C], error) {
	contractDeploy := deploy(chain)
	if contractDeploy.Err != nil {
		lggr.Errorw("Failed to deploy contract", "err", contractDeploy.Err)
		return nil, contractDeploy.Err
	}
	err := chain.Confirm(contractDeploy.Tx.Hash())
	if err != nil {
		lggr.Errorw("Failed to confirm deployment", "err", err)
		return nil, err
	}
	err = addressBook.Save(chain.Selector, contractDeploy.Address.String(), contractDeploy.TvStr)
	if err != nil {
		lggr.Errorw("Failed to save contract address", "err", err)
		return nil, err
	}
	return &contractDeploy, nil
}

type DeployCCIPContractConfig struct {
	HomeChainSel uint64
	// Existing contracts which we want to skip deployment
	// Leave empty if we want to deploy everything
	// TODO: Add skips to deploy function.
	CCIPOnChainState
}

// TODO: Likely we'll want to further parameterize the deployment
// For example a list of contracts to skip deploying if they already exist.
// Or mock vs real RMN.
// Deployment produces an address book of everything it deployed.
func DeployCCIPContracts(e deployment.Environment, c DeployCCIPContractConfig) (deployment.AddressBook, error) {
	ab := deployment.NewMemoryAddressBook()
	for sel, chain := range e.Chains {
		if c.HomeChainSel == sel {
		}

		// TODO: Still waiting for RMNRemote/RMNHome contracts etc.
		mockARM, err := deployContract(e.Logger, chain, ab,
			func(chain deployment.Chain) ContractDeploy[*mock_arm_contract.MockARMContract] {
				mockARMAddr, tx, mockARM, err2 := mock_arm_contract.DeployMockARMContract(
					chain.DeployerKey,
					chain.Client,
				)
				return ContractDeploy[*mock_arm_contract.MockARMContract]{
					mockARMAddr, mockARM, tx, MockARM_1_0_0, err2,
				}
			})
		if err != nil {
			e.Logger.Errorw("Failed to deploy mockARM", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed mockARM", "addr", mockARM)

		mcm, err := deployContract(e.Logger, chain, ab,
			func(chain deployment.Chain) ContractDeploy[*owner_helpers.ManyChainMultiSig] {
				mcmAddr, tx, mcm, err2 := owner_helpers.DeployManyChainMultiSig(
					chain.DeployerKey,
					chain.Client,
				)
				return ContractDeploy[*owner_helpers.ManyChainMultiSig]{
					mcmAddr, mcm, tx, MCMS_1_0_0, err2,
				}
			})
		if err != nil {
			e.Logger.Errorw("Failed to deploy mcm", "err", err)
			return ab, err
		}
		// TODO: Address soon
		e.Logger.Infow("deployed mcm", "addr", mcm.Address)

		_, err = deployContract(e.Logger, chain, ab,
			func(chain deployment.Chain) ContractDeploy[*owner_helpers.RBACTimelock] {
				timelock, tx, cc, err2 := owner_helpers.DeployRBACTimelock(
					chain.DeployerKey,
					chain.Client,
					big.NewInt(0), // minDelay
					mcm.Address,
					[]common.Address{mcm.Address},            // proposers
					[]common.Address{chain.DeployerKey.From}, //executors
					[]common.Address{mcm.Address},            // cancellers
					[]common.Address{mcm.Address},            // bypassers
				)
				return ContractDeploy[*owner_helpers.RBACTimelock]{
					timelock, cc, tx, RBAC_Timelock_1_0_0, err2,
				}
			})
		if err != nil {
			e.Logger.Errorw("Failed to deploy timelock", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed timelock", "addr", mcm.Address)

		armProxy, err := deployContract(e.Logger, chain, ab,
			func(chain deployment.Chain) ContractDeploy[*arm_proxy_contract.ARMProxyContract] {
				armProxyAddr, tx, armProxy, err2 := arm_proxy_contract.DeployARMProxyContract(
					chain.DeployerKey,
					chain.Client,
					mockARM.Address,
				)
				return ContractDeploy[*arm_proxy_contract.ARMProxyContract]{
					armProxyAddr, armProxy, tx, ARMProxy_1_1_0, err2,
				}
			})
		if err != nil {
			e.Logger.Errorw("Failed to deploy armProxy", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed armProxy", "addr", armProxy.Address)

		weth9, err := deployContract(e.Logger, chain, ab,
			func(chain deployment.Chain) ContractDeploy[*weth9.WETH9] {
				weth9Addr, tx, weth9c, err2 := weth9.DeployWETH9(
					chain.DeployerKey,
					chain.Client,
				)
				return ContractDeploy[*weth9.WETH9]{
					weth9Addr, weth9c, tx, WETH9_1_0_0, err2,
				}
			})
		if err != nil {
			e.Logger.Errorw("Failed to deploy weth9", "err", err)
			return ab, err
		}

		linkToken, err := deployContract(e.Logger, chain, ab,
			func(chain deployment.Chain) ContractDeploy[*burn_mint_erc677.BurnMintERC677] {
				linkTokenAddr, tx, linkToken, err2 := burn_mint_erc677.DeployBurnMintERC677(
					chain.DeployerKey,
					chain.Client,
					"Link Token",
					"LINK",
					uint8(18),
					big.NewInt(0).Mul(big.NewInt(1e9), big.NewInt(1e18)),
				)
				return ContractDeploy[*burn_mint_erc677.BurnMintERC677]{
					linkTokenAddr, linkToken, tx, LinkToken_1_0_0, err2,
				}
			})
		if err != nil {
			e.Logger.Errorw("Failed to deploy linkToken", "err", err)
			return ab, err
		}

		routerContract, err := deployContract(e.Logger, chain, ab,
			func(chain deployment.Chain) ContractDeploy[*router.Router] {
				routerAddr, tx, routerC, err2 := router.DeployRouter(
					chain.DeployerKey,
					chain.Client,
					weth9.Address,
					armProxy.Address,
				)
				return ContractDeploy[*router.Router]{
					routerAddr, routerC, tx, Router_1_2_0, err2,
				}
			})
		if err != nil {
			e.Logger.Errorw("Failed to deploy router", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed router", "addr", routerContract)

		tokenAdminRegistry, err := deployContract(e.Logger, chain, ab,
			func(chain deployment.Chain) ContractDeploy[*token_admin_registry.TokenAdminRegistry] {
				tokenAdminRegistryAddr, tx, tokenAdminRegistry, err2 := token_admin_registry.DeployTokenAdminRegistry(
					chain.DeployerKey,
					chain.Client)
				return ContractDeploy[*token_admin_registry.TokenAdminRegistry]{
					tokenAdminRegistryAddr, tokenAdminRegistry, tx, TokenAdminRegistry_1_5_0, err2,
				}
			})
		if err != nil {
			e.Logger.Errorw("Failed to deploy token admin registry", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed tokenAdminRegistry", "addr", tokenAdminRegistry)

		nonceManager, err := deployContract(e.Logger, chain, ab,
			func(chain deployment.Chain) ContractDeploy[*nonce_manager.NonceManager] {
				nonceManagerAddr, tx, nonceManager, err2 := nonce_manager.DeployNonceManager(
					chain.DeployerKey,
					chain.Client,
					[]common.Address{}, // Need to add onRamp after
				)
				return ContractDeploy[*nonce_manager.NonceManager]{
					nonceManagerAddr, nonceManager, tx, NonceManager_1_6_0, err2,
				}
			})
		if err != nil {
			e.Logger.Errorw("Failed to deploy router", "err", err)
			return ab, err
		}

		priceRegistry, err := deployContract(e.Logger, chain, ab,
			func(chain deployment.Chain) ContractDeploy[*price_registry.PriceRegistry] {
				prAddr, tx, pr, err2 := price_registry.DeployPriceRegistry(
					chain.DeployerKey,
					chain.Client,
					price_registry.PriceRegistryStaticConfig{
						MaxFeeJuelsPerMsg:  big.NewInt(0).Mul(big.NewInt(2e2), big.NewInt(1e18)),
						LinkToken:          linkToken.Address,
						StalenessThreshold: uint32(86400),
					},
					[]common.Address{},              // ramps added after
					[]common.Address{weth9.Address}, // fee tokens
					[]price_registry.PriceRegistryTokenPriceFeedUpdate{},
					[]price_registry.PriceRegistryTokenTransferFeeConfigArgs{}, // TODO: tokens
					[]price_registry.PriceRegistryPremiumMultiplierWeiPerEthArgs{
						{
							Token:                      weth9.Address,
							PremiumMultiplierWeiPerEth: 1e6,
						},
					},
					[]price_registry.PriceRegistryDestChainConfigArgs{},
				)
				return ContractDeploy[*price_registry.PriceRegistry]{
					prAddr, pr, tx, PriceRegistry_1_6_0, err2,
				}
			})
		if err != nil {
			e.Logger.Errorw("Failed to deploy price registry", "err", err)
			return ab, err
		}

		onRamp, err := deployContract(e.Logger, chain, ab,
			func(chain deployment.Chain) ContractDeploy[*evm_2_evm_multi_onramp.EVM2EVMMultiOnRamp] {
				onRampAddr, tx, onRamp, err2 := evm_2_evm_multi_onramp.DeployEVM2EVMMultiOnRamp(
					chain.DeployerKey,
					chain.Client,
					evm_2_evm_multi_onramp.EVM2EVMMultiOnRampStaticConfig{
						ChainSelector:      sel,
						RmnProxy:           routerContract.Address,
						NonceManager:       nonceManager.Address,
						TokenAdminRegistry: tokenAdminRegistry.Address,
					},
					evm_2_evm_multi_onramp.EVM2EVMMultiOnRampDynamicConfig{
						PriceRegistry: priceRegistry.Address,
						FeeAggregator: common.HexToAddress("0x1"), // TODO real fee aggregator
					},
					[]evm_2_evm_multi_onramp.EVM2EVMMultiOnRampDestChainConfigArgs{},
				)
				return ContractDeploy[*evm_2_evm_multi_onramp.EVM2EVMMultiOnRamp]{
					onRampAddr, onRamp, tx, EVM2EVMMultiOnRamp_1_6_0, err2,
				}
			})
		if err != nil {
			e.Logger.Errorw("Failed to deploy onramp", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed onramp", "addr", onRamp.Address)

		offRamp, err := deployContract(e.Logger, chain, ab,
			func(chain deployment.Chain) ContractDeploy[*evm_2_evm_multi_offramp.EVM2EVMMultiOffRamp] {
				offRamp, tx, _, err2 := evm_2_evm_multi_offramp.DeployEVM2EVMMultiOffRamp(
					chain.DeployerKey,
					chain.Client,
					evm_2_evm_multi_offramp.EVM2EVMMultiOffRampStaticConfig{
						ChainSelector:      sel,
						RmnProxy:           routerContract.Address,
						NonceManager:       nonceManager.Address,
						TokenAdminRegistry: tokenAdminRegistry.Address,
					},
					evm_2_evm_multi_offramp.EVM2EVMMultiOffRampDynamicConfig{
						PriceRegistry:                           priceRegistry.Address,
						PermissionLessExecutionThresholdSeconds: uint32(86400),
						MaxTokenTransferGas:                     uint32(200_000),
						MaxPoolReleaseOrMintGas:                 uint32(200_000),
					},
					[]evm_2_evm_multi_offramp.EVM2EVMMultiOffRampSourceChainConfigArgs{},
				)
				return ContractDeploy[*evm_2_evm_multi_offramp.EVM2EVMMultiOffRamp]{
					offRamp, nil, tx, EVM2EVMMultiOffRamp_1_6_0, err2,
				}
			})
		if err != nil {
			e.Logger.Errorw("Failed to deploy offramp", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed offramp", "addr", offRamp)

		// Enable ramps on price registry/nonce manager
		tx, err := priceRegistry.Contract.ApplyAuthorizedCallerUpdates(chain.DeployerKey, price_registry.AuthorizedCallersAuthorizedCallerArgs{
			AddedCallers: []common.Address{offRamp.Address},
		})
		if err := chain.Confirm(tx.Hash()); err != nil {
			e.Logger.Errorw("Failed to confirm price registry authorized caller update", "err", err)
			return ab, err
		}
		tx, err = nonceManager.Contract.ApplyAuthorizedCallerUpdates(chain.DeployerKey, nonce_manager.AuthorizedCallersAuthorizedCallerArgs{
			AddedCallers: []common.Address{offRamp.Address, onRamp.Address},
		})
		if err != nil {
			e.Logger.Errorw("Failed to update nonce manager with ramps", "err", err)
			return ab, err
		}
		if err := chain.Confirm(tx.Hash()); err != nil {
			e.Logger.Errorw("Failed to confirm price registry authorized caller update", "err", err)
			return ab, err
		}
	}
	return ab, nil
}
