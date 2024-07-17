package deployments

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/deployment/ownerhelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/deployment/ownerhelpers/rbactimelock"
	"math/big"
)

// For these top level workflows we update the address after every deployment so that if
// if errors we can resume from wherever we left off.
// TODO: bind.ContractBackend would have to be abstracted for multi-family (OR just have a switch statement
// for the handful of families)
// deployNewCCIPContracts deploys a completely fresh V2 set of contracts
func deployNewCCIPContracts(addressBook ContractAddressBook, chains map[uint64]bind.ContractBackend, auth *bind.TransactOpts) error {
	for chainSelector, rpc := range chains {
		// Deploy relevant contracts per chain
		offRampAddr, _, _, _ := evm_2_evm_multi_offramp.DeployEVM2EVMMultiOffRamp(auth, rpc,
			evm_2_evm_multi_offramp.EVM2EVMMultiOffRampStaticConfig{},
			evm_2_evm_multi_offramp.EVM2EVMMultiOffRampDynamicConfig{},
			[]evm_2_evm_multi_offramp.EVM2EVMMultiOffRampSourceChainConfigArgs{})
		// TODO: Confirm TX
		// Once confirmed save the address
		if err := addressBook.Save(chainSelector, offRampAddr.String()); err != nil {
			return err
		}
		// Rest of deployments
	}
	for chainSelector, rpc := range chains {
		// Wire contracts for each chain
		// Note we need not save the configuration, we regenerate from the chain.
		fmt.Println(chainSelector, rpc)
	}
	return nil
}

// We can implement a JobServiceClient which sends the jobs to an in memory chainlink application.
func deployJobSpecs(nodesToJobs map[string][]CCIPSpec, jobClient JobServiceClient) error {
	for node, jobs := range nodesToJobs {
		for _, job := range jobs {
			// We shouldn't need to persist this data, thats on the job distributor.
			// It holds the source of truth.
			_, err := jobClient.ProposeJob(context.Background(), &ProposeJobRequest{
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

// Deploys a brand new CCIP system (on/offchain components) read for onchain messaging.
func DeployNewCCIPToExistingDON(addressBook ContractAddressBook,
	nodesIds []string,
	chains map[uint64]bind.ContractBackend,
	deployerKey *bind.TransactOpts,
	jobServiceClient JobServiceClient) error {
	if err := deployNewCCIPContracts(addressBook, chains, deployerKey); err != nil {
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
func ProposePremiumMultiplierUpdates(rpcs map[uint64]bind.ContractBackend,
	addressBook ContractAddressBook,
	opts *bind.TransactOpts,
	chainsToApply []uint64,
	premiumMultiplierUpdates []evm_2_evm_multi_onramp.EVM2EVMMultiOnRampPremiumMultiplierWeiPerEthArgs,
	delay *big.Int,
) (ownerhelpers.SetRootArgs, error) {
	var setRootArgs ownerhelpers.SetRootArgs
	state, err := generateOnchainState(rpcs, addressBook)
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
