package deployment

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink/v2/core/environment"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/token_admin_registry"
)

// TODO: pull up to environment pkg
func deployContract(
	lggr logger.Logger,
	deploy func() (string, common.Hash, error),
	confirm func(common.Hash) error,
	save func(string) error,
) error {
	contractAddr, tx, err := deploy()
	if err != nil {
		lggr.Errorw("Failed to deploy contract", "err", err)
		return err
	}
	err = confirm(tx)
	if err != nil {
		lggr.Errorw("Failed to confirm deployment", "err", err)
		return err
	}
	err = save(contractAddr)
	if err != nil {
		lggr.Errorw("Failed to save contract address", "err", err)
		return err
	}
	return nil
}

func DeployCCIPContracts(e environment.Environment) error {
	for _, chain := range e.Chains {
		// For example deploy token admin registry to all chains
		// And save the address
		err := deployContract(e.Logger,
			func() (string, common.Hash, error) {
				tokenAdminRegistry, tx, _, err := token_admin_registry.DeployTokenAdminRegistry(
					chain.DeployerKey,
					chain.Client)
				return tokenAdminRegistry.String(), tx.Hash(), err
			},
			chain.Confirm,
			func(addr string) error {
				return e.AddressBook.Save(chain.Selector, addr)
			},
		)
		if err != nil {
			e.Logger.Errorw("Failed to deploy token admin registry", "err", err)
			return err
		}
	}
	return nil
}
