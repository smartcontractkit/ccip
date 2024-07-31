package deployment

import (
	"github.com/smartcontractkit/chainlink/v2/core/environment"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/token_admin_registry"
)

func DeployCCIPContracts(e environment.Environment) error {
	for _, chain := range e.Chains {
		// For example deploy token admin registry to all chains
		// And save the address
		tokenAdminRegistry, tx, _, err := token_admin_registry.DeployTokenAdminRegistry(
			chain.DeployerKey,
			chain.Client)
		if err != nil {
			e.Logger.Errorw("Failed to deploy token admin registry", "err", err)
			return err
		}
		e.Logger.Infow("Deployed token admin registry", "address", tokenAdminRegistry.String(), "tx", tx.Hash(), "from", chain.DeployerKey.From.Hex())
		if err := chain.Confirm(tx.Hash()); err != nil {
			e.Logger.Errorw("Failed to confirm registry deployment", "err", err)
			return err
		}
		// Note the address book also serves as checkpointing mechanism for the deployment.
		err = e.AddressBook.Save(chain.Selector, tokenAdminRegistry.String())
		if err != nil {
			return err
		}
	}
	return nil
}
