package ccipdeployment

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	chainsel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment/executable"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/managed"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment"
)

// TODO: Pull up to deploy
func SimTransactOpts() *bind.TransactOpts {
	return &bind.TransactOpts{Signer: func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
		return transaction, nil
	}, From: common.HexToAddress("0x0"), NoSend: true, GasLimit: 1_000_000}
}

func GenerateAcceptOwnershipProposal(
	e deployment.Environment,
	chains []uint64,
	ab deployment.AddressBook,
) (managed.MCMSWithTimelockProposal, error) {
	state, err := LoadOnchainState(e, ab)
	if err != nil {
		return managed.MCMSWithTimelockProposal{}, err
	}
	// TODO: Just onramp as an example
	var batches []managed.DetailedBatchChainOperation
	metaDataPerChain := make(map[string]managed.MCMSWithTimelockChainMetadata)
	for _, sel := range chains {
		chain, _ := chainsel.ChainBySelector(sel)
		acceptOnRamp, err := state.Chains[sel].OnRamp.AcceptOwnership(SimTransactOpts())
		if err != nil {
			return managed.MCMSWithTimelockProposal{}, err
		}
		metaDataPerChain[chain.Name] = managed.MCMSWithTimelockChainMetadata{
			ExecutableMCMSChainMetadata: executable.ExecutableMCMSChainMetadata{
				NonceOffset: 0,
				MCMAddress:  state.Chains[sel].McmAddr,
			},
			TimelockAddress: state.Chains[sel].TimelockAddr,
		}
		batches = append(batches, managed.DetailedBatchChainOperation{
			ChainIdentifier: chain.Name,
			Batch: []managed.DetailedOperation{
				{
					// Enable the source in on ramp
					Operation: executable.Operation{
						To:    state.Chains[sel].OnRamp.Address(),
						Data:  hexutil.Encode(acceptOnRamp.Data()),
						Value: 0,
					},
				},
			},
		})
	}
	// TODO: Real valid until.
	return managed.MCMSWithTimelockProposal{
		Operation:     managed.Schedule,
		MinDelay:      "1h",
		ChainMetadata: metaDataPerChain,
		Transactions:  batches,
	}, nil
}
