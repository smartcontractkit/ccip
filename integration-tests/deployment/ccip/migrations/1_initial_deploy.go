package migrations

import (
	"github.com/ethereum/go-ethereum/common"
	deployment2 "github.com/smartcontractkit/ccip/integration-tests/deployment"

	ccipdeployment "github.com/smartcontractkit/chainlink/v2/core/capabilities/ccip/deployment"
)

// We expect the migration input to be unique per migration.
// TODO: Maybe there's a generics approach here?
func Apply0001(env deployment2.Environment, c ccipdeployment.DeployCCIPContractConfig) (deployment2.MigrationOutput, error) {
	ab, err := ccipdeployment.DeployCCIPContracts(env, c)
	if err != nil {
		// If we fail here, just throw away the addresses.
		// TODO: if expensive could consider partial recovery
		env.Logger.Errorw("Failed to deploy CCIP contracts", "err", err, "addresses", ab)
		return deployment2.MigrationOutput{}, err
	}
	state, err := ccipdeployment.GenerateOnchainState(env, ab)
	if err != nil {
		return deployment2.MigrationOutput{}, err
	}
	js, err := ccipdeployment.GenerateJobSpecs(common.Address{})
	if err != nil {
		return deployment2.MigrationOutput{}, err
	}
	proposal, err := ccipdeployment.GenerateAcceptOwnershipProposal(env, env.AllChainSelectors(), state)
	if err != nil {
		return deployment2.MigrationOutput{}, err
	}
	return deployment2.MigrationOutput{
		Proposals:   []deployment2.Proposal{proposal},
		AddressBook: ab,
		JobSpecs: map[string][]string{
			"chain-layer": {js.String()},
		},
	}, nil
}
