package deployment

import (
	"fmt"
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
			return err
		}
		if err := chain.Confirm(tx.Hash()); err != nil {
			return err
		}
		fmt.Println("Saving", chain.Selector, tokenAdminRegistry.String())
		err = e.AddressBook.Save(chain.Selector, tokenAdminRegistry.String())
		if err != nil {
			return err
		}
	}
	return nil
}
