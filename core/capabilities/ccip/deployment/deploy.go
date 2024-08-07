package deployment

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink/v2/core/deployment"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/arm_proxy_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/mock_arm_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/token_admin_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/weth9"
)

var (
	ARMProxy_1_1_0             = "ARMProxy 1.0.0"
	MockARM_1_0_0              = "MockARM 1.0.0"
	LinkToken_1_0_0            = "LinkToken 1.0.0"
	TokenAdminRegistry_1_0_0   = "TokenAdminRegistry 1.0.0"
	WETH9_1_0_0                = "WETH9 1.0.0"
	Router_1_0_0               = "Router 1.0.0"
	CapabilitiesRegistry_1_0_0 = "CapabilitiesRegistry 1.0.0"
	EVM2EVMMultiOnRamp_1_6_0   = "EVM2EVMMultiOnRamp 1.6.0-dev"
	EVM2EVMMultiOffRamp_1_6_0  = "EVM2EVMMultiOffRamp 1.6.0-dev"
	PriceRegistry_1_0_0        = "PriceRegistry 1.0.0"
	NonceManager_1_0_0         = "NonceManager 1.0.0"
)

// TODO: pull up to general deployment pkg
func deployContract(
	lggr logger.Logger,
	deploy func() (common.Address, string, common.Hash, error),
	confirm func(common.Hash) error,
	save func(address common.Address, tv string) error,
) (common.Address, error) {
	contractAddr, tvStr, tx, err := deploy()
	if err != nil {
		lggr.Errorw("Failed to deploy contract", "err", err)
		return common.Address{}, err
	}
	err = confirm(tx)
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

func GenerateJobSpecs(capReg common.Address) (CCIPSpec, error) {
	return CCIPSpec{}, nil
}

type DeployCCIPContractConfig struct {
	Weth9s map[uint64]common.Address
	// TODO: More params as needed
}

// TODO: Likely we'll want to further parameterize the deployment
// For example a list of contracts to skip deploying if they already exist.
// Or mock vs real RMN.
// Deployment produces an address book of everything it deployed.
func DeployCCIPContracts(e deployment.Environment, c DeployCCIPContractConfig) (deployment.AddressBook, error) {
	ab := deployment.NewMemoryAddressBook()
	for _, chain := range e.Chains {
		saveToChain := func(addr common.Address, tv string) error {
			return ab.Save(chain.Selector, addr.String(), tv)
		}

		// TODO: Still waiting for RMNRemote/RMNHome contracts etc.
		mockARM, err := deployContract(e.Logger,
			func() (common.Address, string, common.Hash, error) {
				mockARM, tx, _, err := mock_arm_contract.DeployMockARMContract(
					chain.DeployerKey,
					chain.Client,
				)
				return mockARM, MockARM_1_0_0, tx.Hash(), err
			}, chain.Confirm, saveToChain)
		if err != nil {
			e.Logger.Errorw("Failed to deploy mockARM", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed mockARM", "addr", mockARM)

		armProxy, err := deployContract(e.Logger,
			func() (common.Address, string, common.Hash, error) {
				armProxy, tx, _, err := arm_proxy_contract.DeployARMProxyContract(
					chain.DeployerKey,
					chain.Client,
					mockARM,
				)
				return armProxy, ARMProxy_1_1_0, tx.Hash(), err
			}, chain.Confirm, saveToChain)
		if err != nil {
			e.Logger.Errorw("Failed to deploy armProxy", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed armProxy", "addr", armProxy)

		weth9, err := deployContract(e.Logger,
			func() (common.Address, string, common.Hash, error) {
				weth9, tx, _, err := weth9.DeployWETH9(
					chain.DeployerKey,
					chain.Client,
				)
				return weth9, WETH9_1_0_0, tx.Hash(), err
			}, chain.Confirm, saveToChain)
		if err != nil {
			e.Logger.Errorw("Failed to deploy weth9", "err", err)
			return ab, err
		}

		routerAddr, err := deployContract(e.Logger,
			func() (common.Address, string, common.Hash, error) {
				router, tx, _, err := router.DeployRouter(
					chain.DeployerKey,
					chain.Client,
					weth9,
					armProxy,
				)
				return router, Router_1_0_0, tx.Hash(), err
			}, chain.Confirm, saveToChain)
		if err != nil {
			e.Logger.Errorw("Failed to deploy router", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed router", "addr", routerAddr)

		tokenAdminRegistry, err := deployContract(e.Logger,
			func() (common.Address, string, common.Hash, error) {
				tokenAdminRegistry, tx, _, err := token_admin_registry.DeployTokenAdminRegistry(
					chain.DeployerKey,
					chain.Client)
				return tokenAdminRegistry, TokenAdminRegistry_1_0_0, tx.Hash(), err
			}, chain.Confirm, saveToChain)
		if err != nil {
			e.Logger.Errorw("Failed to deploy token admin registry", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed tokenAdminRegistry", "addr", tokenAdminRegistry)
	}
	return ab, nil
}
