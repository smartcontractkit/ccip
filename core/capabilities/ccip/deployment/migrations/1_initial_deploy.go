package migrations

import (
	"github.com/ethereum/go-ethereum/common"

	ccipdeployment "github.com/smartcontractkit/chainlink/v2/core/capabilities/ccip/deployment"
	"github.com/smartcontractkit/chainlink/v2/core/deployment"
)

// We expect the migration input to be unique per migration.
// TODO: Maybe there's a generics approach here?
func Apply0001(env deployment.Environment, c ccipdeployment.DeployCCIPContractConfig) (deployment.MigrationOutput, error) {
	ab, err := ccipdeployment.DeployCCIPContracts(env, c)
	if err != nil {
		// If we fail here, just throw away the addresses.
		// TODO: if expensive could consider partial recovery
		env.Logger.Errorw("Failed to deploy CCIP contracts", "err", err, "addresses", ab)
		return deployment.MigrationOutput{}, err
	}
	state, err := ccipdeployment.GenerateOnchainState(env, ab)
	if err != nil {
		return deployment.MigrationOutput{}, err
	}
	js, err := ccipdeployment.GenerateJobSpecs(common.Address{})
	if err != nil {
		return deployment.MigrationOutput{}, err
	}
	proposal, err := ccipdeployment.GenerateAcceptOwnershipProposal(env, env.AllChainSelectors(), state)
	if err != nil {
		return deployment.MigrationOutput{}, err
	}
	return deployment.MigrationOutput{
		Proposals:   []deployment.Proposal{proposal},
		AddressBook: ab,
		JobSpecs: map[string][]string{
			"chain-layer": {js.String()},
		},
	}, nil
}
