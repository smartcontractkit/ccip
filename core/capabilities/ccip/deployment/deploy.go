package deployment

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink/v2/core/environment"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/arm_proxy_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/mock_arm_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/token_admin_registry"
)

type Proposal struct {
}

func (p Proposal) String() string {
	return ""
}

func GenerateAcceptOwnershipProposal(e environment.Environment, state CCIPOnChainState) Proposal {
	return Proposal{}
}

// TODO: pull up to environment pkg
func deployContract(
	lggr logger.Logger,
	deploy func() (common.Address, common.Hash, error),
	confirm func(common.Hash) error,
	save func(address common.Address) error,
) (common.Address, error) {
	contractAddr, tx, err := deploy()
	if err != nil {
		lggr.Errorw("Failed to deploy contract", "err", err)
		return common.Address{}, err
	}
	err = confirm(tx)
	if err != nil {
		lggr.Errorw("Failed to confirm deployment", "err", err)
		return common.Address{}, err
	}
	err = save(contractAddr)
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

func GenerateJobSpecs(capReg common.Address) CCIPSpec {
	return CCIPSpec{}
}

type DeployCCIPContractConfig struct {
	Weth9s map[uint64]common.Address
	// TODO: More params as needed
}

// TODO: Likely we'll want to further parameterize the deployment
// For example a list of contracts to skip deploying if they already exist.
// Or mock vs real RMN.
// Deployment produces an address book of everything it deployed.
func DeployCCIPContracts(e environment.Environment, c DeployCCIPContractConfig) (environment.AddressBook, error) {
	ab := environment.NewMemoryAddressBook()
	for _, chain := range e.Chains {
		saveToChain := func(addr common.Address) error {
			return ab.Save(chain.Selector, addr.String())
		}

		// TODO: Still waiting for RMNRemote/RMNHome contracts etc.
		mockARM, err := deployContract(e.Logger,
			func() (common.Address, common.Hash, error) {
				mockARM, tx, _, err := mock_arm_contract.DeployMockARMContract(
					chain.DeployerKey,
					chain.Client,
				)
				return mockARM, tx.Hash(), err
			}, chain.Confirm, saveToChain)
		if err != nil {
			e.Logger.Errorw("Failed to deploy mockARM", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed mockARM", "addr", mockARM)

		armProxy, err := deployContract(e.Logger,
			func() (common.Address, common.Hash, error) {
				mockARM, tx, _, err := arm_proxy_contract.DeployARMProxyContract(
					chain.DeployerKey,
					chain.Client,
					mockARM,
				)
				return mockARM, tx.Hash(), err
			}, chain.Confirm, saveToChain)
		if err != nil {
			e.Logger.Errorw("Failed to deploy armProxy", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed armProxy", "addr", armProxy)

		//weth9, err := deployContract(e.Logger,
		//	func() (common.Address, common.Hash, error) {
		//		weth9, tx, _, err := weth9.DeployWETH9(
		//			chain.DeployerKey,
		//			chain.Client,
		//		)
		//		return weth9, tx.Hash(), err
		//	}, chain.Confirm, saveToChain)
		//if err != nil {
		//	e.Logger.Errorw("Failed to deploy weth9", "err", err)
		//	return err
		//}

		routerAddr, err := deployContract(e.Logger,
			func() (common.Address, common.Hash, error) {
				router, tx, _, err := router.DeployRouter(
					chain.DeployerKey,
					chain.Client,
					common.HexToAddress("0x0"),
					armProxy,
				)
				return router, tx.Hash(), err
			}, chain.Confirm, saveToChain)
		if err != nil {
			e.Logger.Errorw("Failed to deploy router", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed router", "addr", routerAddr)

		tokenAdminRegistry, err := deployContract(e.Logger,
			func() (common.Address, common.Hash, error) {
				tokenAdminRegistry, tx, _, err := token_admin_registry.DeployTokenAdminRegistry(
					chain.DeployerKey,
					chain.Client)
				return tokenAdminRegistry, tx.Hash(), err
			}, chain.Confirm, saveToChain)
		if err != nil {
			e.Logger.Errorw("Failed to deploy token admin registry", "err", err)
			return ab, err
		}
		e.Logger.Infow("deployed tokenAdminRegistry", "addr", tokenAdminRegistry)
	}
	return ab, nil
}
