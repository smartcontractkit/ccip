package ccipdeployment

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	nodev1 "github.com/smartcontractkit/chainlink/integration-tests/deployment/jd/node/v1"

	owner_helpers "github.com/smartcontractkit/ccip-owner-contracts/gethwrappers"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment"
	"github.com/smartcontractkit/chainlink/v2/core/capabilities/ccip/validate"
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
	"github.com/smartcontractkit/chainlink/v2/core/services/relay"
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
	EVM2EVMMultiOnRamp_1_6_0   = "EVM2EVMMultiOnRamp 1.6.0-dev"
	EVM2EVMMultiOffRamp_1_6_0  = "EVM2EVMMultiOffRamp 1.6.0-dev"
	NonceManager_1_6_0         = "NonceManager 1.6.0-dev"
	PriceRegistry_1_6_0        = "PriceRegistry 1.6.0-dev"

	CapabilityVersion = "1.0.0"
)

// TODO: pull up to general deployment pkg
func deployContract(
	lggr logger.Logger,
	deploy func() (common.Address, string, *types.Transaction, error),
	confirm func(common.Hash) error,
	save func(address common.Address, tv string) error,
) (common.Address, error) {
	contractAddr, tvStr, tx, err := deploy()
	if err != nil {
		lggr.Errorw("Failed to deploy contract", "err", err)
		return common.Address{}, err
	}
	err = confirm(tx.Hash())
	if err != nil {
		lggr.Errorw("Failed to confirm deployment", "err", err)
		return common.Address{}, err
	}
	err = save(contractAddr, tvStr)
	if err != nil {
		lggr.Errorw("Failed to save contract address", "err", err)
		return common.Address{}, err
	}
	return contractAddr, nil
}

type CCIPSpec struct{}

func (s CCIPSpec) String() string {
	return ""
}

// In our case, the only address needed is the cap registry which is actually an env var.
// and will pre-exist for our deployment. So the job specs only depend on the environment operators.
func NewCCIPJobSpecs(nodeIds []string, oc deployment.OffchainClient) (map[string][]string, error) {
	// Generate a set of brand new job specs for CCIP for a specific environment
	// (including NOPs) and new addresses.
	// We want to assign one CCIP capability job to each node. And node with
	// an addr we'll list as bootstrapper.
	// Find the bootstrap nodes
	bootstrapMp := make(map[string]struct{})
	for _, node := range nodeIds {
		// TODO: Filter should accept multiple nodes
		nodeChainConfigs, err := oc.ListNodeChainConfigs(context.Background(), &nodev1.ListNodeChainConfigsRequest{Filter: &nodev1.ListNodeChainConfigsRequest_Filter{
			NodeId: node,
		}})
		if err != nil {
			return nil, err
		}
		for _, chainConfig := range nodeChainConfigs.ChainConfigs {
			if chainConfig.Ocr2Config.IsBootstrap {
				bootstrapMp[fmt.Sprintf("%s@%s",
					// p2p_12D3... -> 12D3...
					chainConfig.Ocr2Config.P2PKeyBundle.PeerId[4:], chainConfig.Ocr2Config.Multiaddr)] = struct{}{}
			}
		}
	}
	var bootstraps []string
	for b := range bootstrapMp {
		bootstraps = append(bootstraps, b)
	}
	nodesToJobSpecs := make(map[string][]string)
	for _, node := range nodeIds {
		// TODO: Filter should accept multiple.
		nodeChainConfigs, err := oc.ListNodeChainConfigs(context.Background(), &nodev1.ListNodeChainConfigsRequest{Filter: &nodev1.ListNodeChainConfigsRequest_Filter{
			NodeId: node,
		}})
		if err != nil {
			return nil, err
		}
		spec, err := validate.NewCCIPSpecToml(validate.SpecArgs{
			P2PV2Bootstrappers:     bootstraps,
			CapabilityVersion:      CapabilityVersion,
			CapabilityLabelledName: "CCIP",
			OCRKeyBundleIDs: map[string]string{
				// TODO: Validate that that all EVM chains are using the same keybundle.
				relay.NetworkEVM: nodeChainConfigs.ChainConfigs[0].Ocr2Config.OcrKeyBundle.BundleId,
			},
			// TODO: validate that all EVM chains are using the same keybundle
			P2PKeyID:     nodeChainConfigs.ChainConfigs[0].Ocr2Config.P2PKeyBundle.PeerId,
			RelayConfigs: nil,
			PluginConfig: map[string]any{},
		})
		if err != nil {
			return nil, err
		}
		nodesToJobSpecs[node] = append(nodesToJobSpecs[node], spec)
	}
	return nodesToJobSpecs, nil
}

type DeployCCIPContractConfig struct {
	// Existing contracts which we want to skip deployment
	// Leave empty if we want to deploy everything
	// TODO: Add skips to deploy function.
	CCIPOnChainState
}

func DeployCapReg(lggr logger.Logger, chains map[uint64]deployment.Chain, chainSel uint64) (deployment.AddressBook, error) {
	ab := deployment.NewMemoryAddressBook()
	chain := chains[chainSel]
	saveToChain := func(addr common.Address, tv string) error {
		return ab.Save(chain.Selector, addr.String(), tv)
	}
	capRegAddr, err := deployContract(lggr,
		func() (common.Address, string, *types.Transaction, error) {
			cr, tx, _, err2 := capabilities_registry.DeployCapabilitiesRegistry(
				chain.DeployerKey,
				chain.Client,
			)
			return cr, CapabilitiesRegistry_1_0_0, tx, err2
		}, chain.Confirm, saveToChain)
	if err != nil {
		lggr.Errorw("Failed to deploy capreg", "err", err)
		return ab, err
	}
	lggr.Infow("deployed capreg", "addr", capRegAddr)
	return ab, nil
}

// TODO: Likely we'll want to further parameterize the deployment
// For example a list of contracts to skip deploying if they already exist.
// Or mock vs real RMN.
// Deployment produces an address book of everything it deployed.
func DeployCCIPContracts(e deployment.Environment, c DeployCCIPContractConfig) (deployment.AddressBook, error) {
	ab := deployment.NewMemoryAddressBook()

	for sel, chain := range e.Chains {
		saveToChain := func(addr common.Address, tv string) error {
			return ab.Save(chain.Selector, addr.String(), tv)
		}

		// TODO: Still waiting for RMNRemote/RMNHome contracts etc.
		mockARM, err := deployContract(e.Logger,
			func() (common.Address, string, *types.Transaction, error) {
				mockARM, tx, _, err2 := mock_arm_contract.DeployMockARMContract(
					chain.DeployerKey,
					chain.Client,
				)
				return mockARM, MockARM_1_0_0, tx, err2
			}, chain.Confirm, saveToChain)
		if err != nil {
			e.Logger.Errorw("Failed to deploy mockARM", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed mockARM", "addr", mockARM)

		mcmAddr, err := deployContract(e.Logger,
			func() (common.Address, string, *types.Transaction, error) {
				mcm, tx, _, err2 := owner_helpers.DeployManyChainMultiSig(
					chain.DeployerKey,
					chain.Client,
				)
				return mcm, MCMS_1_0_0, tx, err2
			}, chain.Confirm, saveToChain)
		if err != nil {
			e.Logger.Errorw("Failed to deploy mcm", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed mcm", "addr", mcmAddr)

		_, err = deployContract(e.Logger,
			func() (common.Address, string, *types.Transaction, error) {
				timelock, tx, _, err2 := owner_helpers.DeployRBACTimelock(
					chain.DeployerKey,
					chain.Client,
					big.NewInt(0), // minDelay
					mcmAddr,
					[]common.Address{mcmAddr},                // proposers
					[]common.Address{chain.DeployerKey.From}, //executors
					[]common.Address{mcmAddr},                // cancellers
					[]common.Address{mcmAddr},                // bypassers
				)
				return timelock, RBAC_Timelock_1_0_0, tx, err2
			}, chain.Confirm, saveToChain)
		if err != nil {
			e.Logger.Errorw("Failed to deploy timelock", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed timelock", "addr", mcmAddr)

		armProxy, err := deployContract(e.Logger,
			func() (common.Address, string, *types.Transaction, error) {
				armProxy, tx, _, err2 := arm_proxy_contract.DeployARMProxyContract(
					chain.DeployerKey,
					chain.Client,
					mockARM,
				)
				return armProxy, ARMProxy_1_1_0, tx, err2
			}, chain.Confirm, saveToChain)
		if err != nil {
			e.Logger.Errorw("Failed to deploy armProxy", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed armProxy", "addr", armProxy)

		weth9, err := deployContract(e.Logger,
			func() (common.Address, string, *types.Transaction, error) {
				weth9, tx, _, err2 := weth9.DeployWETH9(
					chain.DeployerKey,
					chain.Client,
				)
				return weth9, WETH9_1_0_0, tx, err2
			}, chain.Confirm, saveToChain)
		if err != nil {
			e.Logger.Errorw("Failed to deploy weth9", "err", err)
			return ab, err
		}

		linkTokenAddr, err := deployContract(e.Logger,
			func() (common.Address, string, *types.Transaction, error) {
				linkToken, tx, _, err2 := burn_mint_erc677.DeployBurnMintERC677(
					chain.DeployerKey,
					chain.Client,
					"Link Token",
					"LINK",
					uint8(18),
					big.NewInt(0).Mul(big.NewInt(1e9), big.NewInt(1e18)),
				)
				return linkToken, LinkToken_1_0_0, tx, err2
			}, chain.Confirm, saveToChain)
		if err != nil {
			e.Logger.Errorw("Failed to deploy linkToken", "err", err)
			return ab, err
		}

		routerAddr, err := deployContract(e.Logger,
			func() (common.Address, string, *types.Transaction, error) {
				router, tx, _, err2 := router.DeployRouter(
					chain.DeployerKey,
					chain.Client,
					weth9,
					armProxy,
				)
				return router, Router_1_2_0, tx, err2
			}, chain.Confirm, saveToChain)
		if err != nil {
			e.Logger.Errorw("Failed to deploy router", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed router", "addr", routerAddr)

		tokenAdminRegistry, err := deployContract(e.Logger,
			func() (common.Address, string, *types.Transaction, error) {
				tokenAdminRegistry, tx, _, err2 := token_admin_registry.DeployTokenAdminRegistry(
					chain.DeployerKey,
					chain.Client)
				return tokenAdminRegistry, TokenAdminRegistry_1_5_0, tx, err2
			}, chain.Confirm, saveToChain)
		if err != nil {
			e.Logger.Errorw("Failed to deploy token admin registry", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed tokenAdminRegistry", "addr", tokenAdminRegistry)

		nonceManagerAddr, err := deployContract(e.Logger,
			func() (common.Address, string, *types.Transaction, error) {
				nonceManager, tx, _, err2 := nonce_manager.DeployNonceManager(
					chain.DeployerKey,
					chain.Client,
					[]common.Address{}, // Need to add onRamp after
				)
				return nonceManager, NonceManager_1_6_0, tx, err2
			}, chain.Confirm, saveToChain)
		if err != nil {
			e.Logger.Errorw("Failed to deploy router", "err", err)
			return ab, err
		}

		priceRegistryAddr, err := deployContract(e.Logger,
			func() (common.Address, string, *types.Transaction, error) {
				pr, tx, _, err2 := price_registry.DeployPriceRegistry(
					chain.DeployerKey,
					chain.Client,
					price_registry.PriceRegistryStaticConfig{
						MaxFeeJuelsPerMsg:  big.NewInt(0).Mul(big.NewInt(2e2), big.NewInt(1e18)),
						LinkToken:          linkTokenAddr,
						StalenessThreshold: uint32(86400),
					},
					[]common.Address{},      // ramps added after
					[]common.Address{weth9}, // fee tokens
					[]price_registry.PriceRegistryTokenPriceFeedUpdate{},
					[]price_registry.PriceRegistryTokenTransferFeeConfigArgs{}, // TODO: tokens
					[]price_registry.PriceRegistryPremiumMultiplierWeiPerEthArgs{
						{
							Token:                      weth9,
							PremiumMultiplierWeiPerEth: 1e6,
						},
					},
					[]price_registry.PriceRegistryDestChainConfigArgs{},
				)
				return pr, PriceRegistry_1_6_0, tx, err2
			}, chain.Confirm, saveToChain)
		if err != nil {
			e.Logger.Errorw("Failed to deploy price registry", "err", err)
			return ab, err
		}

		onRampAddr, err := deployContract(e.Logger,
			func() (common.Address, string, *types.Transaction, error) {
				onRamp, tx, _, err2 := evm_2_evm_multi_onramp.DeployEVM2EVMMultiOnRamp(
					chain.DeployerKey,
					chain.Client,
					evm_2_evm_multi_onramp.EVM2EVMMultiOnRampStaticConfig{
						ChainSelector:      sel,
						RmnProxy:           routerAddr,
						NonceManager:       nonceManagerAddr,
						TokenAdminRegistry: tokenAdminRegistry,
					},
					evm_2_evm_multi_onramp.EVM2EVMMultiOnRampDynamicConfig{
						PriceRegistry: priceRegistryAddr,
						FeeAggregator: common.HexToAddress("0x1"), // TODO real fee aggregator
					},
					[]evm_2_evm_multi_onramp.EVM2EVMMultiOnRampDestChainConfigArgs{},
				)
				return onRamp, EVM2EVMMultiOnRamp_1_6_0, tx, err2
			}, chain.Confirm, saveToChain)
		if err != nil {
			e.Logger.Errorw("Failed to deploy onramp", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed onramp", "addr", tokenAdminRegistry)

		offRampAddr, err := deployContract(e.Logger,
			func() (common.Address, string, *types.Transaction, error) {
				offRamp, tx, _, err2 := evm_2_evm_multi_offramp.DeployEVM2EVMMultiOffRamp(
					chain.DeployerKey,
					chain.Client,
					evm_2_evm_multi_offramp.EVM2EVMMultiOffRampStaticConfig{
						ChainSelector:      sel,
						RmnProxy:           routerAddr,
						NonceManager:       nonceManagerAddr,
						TokenAdminRegistry: tokenAdminRegistry,
					},
					evm_2_evm_multi_offramp.EVM2EVMMultiOffRampDynamicConfig{
						PriceRegistry:                           priceRegistryAddr,
						PermissionLessExecutionThresholdSeconds: uint32(86400),
						MaxTokenTransferGas:                     uint32(200_000),
						MaxPoolReleaseOrMintGas:                 uint32(200_000),
					},
					[]evm_2_evm_multi_offramp.EVM2EVMMultiOffRampSourceChainConfigArgs{},
				)
				return offRamp, EVM2EVMMultiOffRamp_1_6_0, tx, err2
			}, chain.Confirm, saveToChain)
		if err != nil {
			e.Logger.Errorw("Failed to deploy offramp", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed offramp", "addr", offRampAddr)

		// Enable ramps on price registry/nonce manager
		pr, err := price_registry.NewPriceRegistry(priceRegistryAddr, chain.Client)
		if err != nil {
			e.Logger.Errorw("Failed to create price registry", "err", err)
			return ab, err
		}
		tx, err := pr.ApplyAuthorizedCallerUpdates(chain.DeployerKey, price_registry.AuthorizedCallersAuthorizedCallerArgs{
			AddedCallers: []common.Address{offRampAddr},
		})
		if err := chain.Confirm(tx.Hash()); err != nil {
			e.Logger.Errorw("Failed to confirm price registry authorized caller update", "err", err)
			return ab, err
		}
		nm, err := nonce_manager.NewNonceManager(nonceManagerAddr, chain.Client)
		if err != nil {
			e.Logger.Errorw("Failed to create nonce manager", "err", err)
			return ab, err
		}
		tx, err = nm.ApplyAuthorizedCallerUpdates(chain.DeployerKey, nonce_manager.AuthorizedCallersAuthorizedCallerArgs{
			AddedCallers: []common.Address{offRampAddr, onRampAddr},
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
