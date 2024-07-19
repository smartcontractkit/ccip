package deployments

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/arm_proxy_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/nonce_manager"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/token_admin_registry"
	kcr "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/capabilities_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/shared/generated/burn_mint_erc677"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/deployment/jobdistributor"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/deployment/ownerhelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/deployment/ownerhelpers/rbactimelock"
	"math/big"
)

func deployAndSaveAddress(
	ab ContractAddressBook,
	chain Chain,
	deploy func(chain Chain) (common.Address, common.Hash, error),
	confirm func(tx common.Hash, chain Chain) error,
) (common.Address, error) {
	addr, hash, err := deploy(chain)
	if err != nil {
		return common.Address{}, err
	}
	if err := confirm(hash, chain); err != nil {
		return common.Address{}, err
	}
	if err := ab.Save(chain.Selector, addr.String()); err != nil {
		// Note if we fail to save but we confirm, then we need to manually retry the save
		// using logs.
		return common.Address{}, err
	}
	return addr, nil
}

// For these top level workflows we update the address after every deployment so that if
// if errors we can resume from wherever we left off.
// TODO: bind.ContractBackend would have to be abstracted for multi-family (OR just have a switch statement
// for the handful of families)
// deployNewCCIPContracts deploys a completely fresh V2 set of contracts
// TODO: confirm each tx after deploying.
func deployNewCCIPContracts(
	addressBook ContractAddressBook,
	chains map[uint64]Chain,
	confirm func(hash common.Hash, chain Chain) error,
) error {
	for chainSelector, chain := range chains {
		linkToken, err := deployAndSaveAddress(addressBook, chain, func(chain Chain) (common.Address, common.Hash, error) {
			linkToken, tx, _, err := burn_mint_erc677.DeployBurnMintERC677(
				chain.Auth,
				chain.Client,
				"link", "LINK", 18, big.NewInt(1e9))
			return linkToken, tx.Hash(), err
		}, confirm)
		if err != nil {
			return err
		}
		tokenAdminRegistry, err := deployAndSaveAddress(addressBook, chain, func(chain Chain) (common.Address, common.Hash, error) {
			tokenAdminRegistry, tx, _, err := token_admin_registry.DeployTokenAdminRegistry(
				chain.Auth,
				chain.Client)
			return tokenAdminRegistry, tx.Hash(), err
		}, confirm)
		if err != nil {
			return err
		}
		rmnProxy, err := deployAndSaveAddress(addressBook, chain, func(chain Chain) (common.Address, common.Hash, error) {
			rmnProxy, tx, _, err := arm_proxy_contract.DeployARMProxyContract(
				chain.Auth,
				chain.Client,
				common.HexToAddress("0x12"))
			return rmnProxy, tx.Hash(), err
		}, confirm)
		if err != nil {
			return err
		}
		router, err := deployAndSaveAddress(addressBook, chain, func(chain Chain) (common.Address, common.Hash, error) {
			router, tx, _, err := router.DeployRouter(
				chain.Auth,
				chain.Client,
				common.HexToAddress("0x1"), rmnProxy)
			return router, tx.Hash(), err
		}, confirm)
		if err != nil {
			return err
		}
		nonceManager, err := deployAndSaveAddress(addressBook, chain, func(chain Chain) (common.Address, common.Hash, error) {
			nonceManager, tx, _, err := nonce_manager.DeployNonceManager(
				chain.Auth,
				chain.Client,
				nil)
			return nonceManager, tx.Hash(), err
		}, confirm)
		if err != nil {
			return err
		}
		priceRegistry, err := deployAndSaveAddress(addressBook, chain, func(chain Chain) (common.Address, common.Hash, error) {
			nonceManager, tx, _, err := price_registry.DeployPriceRegistry(
				chain.Auth,
				chain.Client,
				[]common.Address{}, []common.Address{}, uint32(90000), []price_registry.PriceRegistryTokenPriceFeedUpdate{})
			return nonceManager, tx.Hash(), err
		}, confirm)
		if err != nil {
			return err
		}
		_, err = deployAndSaveAddress(addressBook, chain, func(chain Chain) (common.Address, common.Hash, error) {
			dest := make([]evm_2_evm_multi_onramp.EVM2EVMMultiOnRampDestChainConfigArgs, 0)
			prem := make([]evm_2_evm_multi_onramp.EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthArgs, 0)
			tt := make([]evm_2_evm_multi_onramp.EVM2EVMMultiOnRampTokenTransferFeeConfigArgs, 0)
			onRampAddr, tx, _, err := evm_2_evm_multi_onramp.DeployEVM2EVMMultiOnRamp(chain.Auth,
				chain.Client,
				evm_2_evm_multi_onramp.EVM2EVMMultiOnRampStaticConfig{
					LinkToken:          linkToken,
					ChainSelector:      chainSelector,
					MaxFeeJuelsPerMsg:  big.NewInt(1e18),
					NonceManager:       nonceManager,
					RmnProxy:           rmnProxy,
					TokenAdminRegistry: tokenAdminRegistry,
				},
				evm_2_evm_multi_onramp.EVM2EVMMultiOnRampDynamicConfig{
					Router:        router,
					PriceRegistry: priceRegistry,
					FeeAggregator: common.HexToAddress("0x123"),
				},
				dest,
				prem,
				tt)
			return onRampAddr, tx.Hash(), err
		}, confirm)
		if err != nil {
			return err
		}
		_, err = deployAndSaveAddress(addressBook, chain, func(chain Chain) (common.Address, common.Hash, error) {
			offRampAddr, tx, _, err := evm_2_evm_multi_offramp.DeployEVM2EVMMultiOffRamp(chain.Auth, chain.Client,
				evm_2_evm_multi_offramp.EVM2EVMMultiOffRampStaticConfig{
					ChainSelector:      chainSelector,
					RmnProxy:           rmnProxy,
					TokenAdminRegistry: tokenAdminRegistry,
					NonceManager:       nonceManager,
				},
				evm_2_evm_multi_offramp.EVM2EVMMultiOffRampDynamicConfig{
					Router:                                  router,
					PermissionLessExecutionThresholdSeconds: uint32(86400),
					MaxTokenTransferGas:                     200_000,
					MaxPoolReleaseOrMintGas:                 200_000,
					MessageValidator:                        common.HexToAddress("0x0"),
					PriceRegistry:                           priceRegistry,
				},
				[]evm_2_evm_multi_offramp.EVM2EVMMultiOffRampSourceChainConfigArgs{})
			return offRampAddr, tx.Hash(), err
		}, confirm)
		if err != nil {
			return fmt.Errorf("offramp %s", err)
		}
	}
	for chainSelector, chain := range chains {
		// Wire contracts for each chain
		// Note we need not save the configuration, we regenerate from the chain.
		fmt.Println(chainSelector, chain.Auth)
	}
	return nil
}

// We can implement a JobServiceClient which sends the jobs to an in memory chainlink application.
func deployJobSpecs(nodesToJobs map[string][]CCIPSpec, jobClient jobdistributor.JobServiceClient) error {
	for node, jobs := range nodesToJobs {
		for _, job := range jobs {
			// We shouldn't need to persist this data, thats on the job distributor.
			// It holds the source of truth.
			_, err := jobClient.ProposeJob(context.Background(), &jobdistributor.ProposeJobRequest{
				NodeId: node,
				Spec:   job.ToTOML(),
			})
			if err != nil {
				return err
			}
		}
	}
	// TODO: could read them back
	return nil
}

type Chain struct {
	Selector uint64 // note can map to evm using selector pkg.
	Client   bind.ContractBackend
	Auth     *bind.TransactOpts
}

func DeployCapabilityRegistry(addressBook ContractAddressBook, chain Chain, confirm func(tx common.Hash, chain Chain) error) error {
	// deploy the capability registry on the home chain
	crAddress, tx, _, err := kcr.DeployCapabilitiesRegistry(chain.Auth, chain.Client)
	if err != nil {
		return err
	}
	if err := addressBook.Save(chain.Selector, crAddress.String()); err != nil {
		return err
	}
	if err := confirm(tx.Hash(), chain); err != nil {
		return err
	}
	fmt.Println("Deployed %s to %s", crAddress, chain.Selector)
	return nil
}

// Deploys a brand new CCIP system (on/offchain components) read for onchain messaging.
func DeployNewCCIPToExistingDON(addressBook ContractAddressBook,
	nodesIds []string,
	chains map[uint64]Chain,
	jobServiceClient jobdistributor.JobServiceClient,
	confirm func(tx common.Hash, chain Chain) error,
) error {
	if err := deployNewCCIPContracts(addressBook, chains, confirm); err != nil {
		return err
	}
	// Obtain the deployed state.
	ccipState, err := generateState(chains, addressBook, nodesIds, jobServiceClient)
	if err != nil {
		return err
	}
	// We expect no job specs yet, just the onchain state to be populated.
	// Build the job specs from the onchain state.
	nodesToJobs := make(map[string][]CCIPSpec)
	for _, nodeId := range nodesIds {
		// One spec needed per node in CCIPv2.
		nodesToJobs[nodeId] = []CCIPSpec{
			{
				CapabilityRegistry: ccipState.CapabilityRegistry.Address(),
			},
		}
	}
	// Build the jobs specs
	if err := deployJobSpecs(nodesToJobs, jobServiceClient); err != nil {
		return err
	}
	return nil
}

// Example of changing something through a proposal.
func ProposePremiumMultiplierUpdates(chains map[uint64]Chain,
	addressBook ContractAddressBook,
	opts *bind.TransactOpts,
	chainsToApply []uint64,
	premiumMultiplierUpdates []evm_2_evm_multi_onramp.EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthArgs,
	delay *big.Int,
) (ownerhelpers.SetRootArgs, error) {
	var setRootArgs ownerhelpers.SetRootArgs
	state, err := GenerateOnchainState(chains, addressBook)
	if err != nil {
		return setRootArgs, err
	}
	// We need to gather calldata for each operation.
	// Don't send it, we're just generating a proposal.
	opts.NoSend = true
	batches := make(map[uint64]ownerhelpers.Batch)
	for _, chain := range chainsToApply {
		tx, err := state.EvmOnRampsV160[chain].ApplyPremiumMultiplierWeiPerEthUpdates(opts, premiumMultiplierUpdates)
		if err != nil {
			return setRootArgs, err
		}
		// Note could use reflection to include metadata (like name of method being called)
		batches[chain] = ownerhelpers.Batch{
			Calls: []rbactimelock.RBACTimelockCall{
				{
					Value:  nil,
					Target: state.EvmOnRampsV160[chain].Address(),
					Data:   tx.Data(),
				},
			},
			Delay: delay,
		}
	}
	// Convert these batches into a root to sign.
	return ownerhelpers.GenerateSetRootArgs(batches)
}
