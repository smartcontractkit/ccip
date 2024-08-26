package ccipdeployment

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	chainsel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/executable"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/managed"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/capabilities_registry"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry"
)

// AddChain deploys chain contracts for a new chain
// and generates 3 proposals to connect that new chain to all existing chains.
// We testing in between each proposal.
func NewChainInbound(
	e deployment.Environment,
	ab deployment.AddressBook,
	homeChainSel uint64,
	newChainSel uint64,
	sources []uint64,
) ([]managed.MCMSWithTimelockProposal, deployment.AddressBook, error) {
	// 1. Deploy contracts to new chain and wire them.
	newAddresses, err := DeployChainContracts(e, e.Chains[newChainSel], deployment.NewMemoryAddressBook())
	if err != nil {
		return nil, ab, err
	}
	if err := ab.Merge(newAddresses); err != nil {
		return nil, ab, err
	}
	state, err := LoadOnchainState(e, ab)
	if err != nil {
		return nil, ab, err
	}

	// 2. Generate proposal which enables new destination (from test router) on all source chains.
	var batches []managed.DetailedBatchChainOperation
	metaDataPerChain := make(map[string]managed.MCMSWithTimelockChainMetadata)
	for _, source := range sources {
		chain, _ := chainsel.ChainBySelector(source)
		enableOnRampDest, err := state.Chains[source].OnRamp.ApplyDestChainConfigUpdates(SimTransactOpts(), []onramp.OnRampDestChainConfigArgs{
			{
				DestChainSelector: newChainSel,
				Router:            state.Chains[source].TestRouter.Address(),
			},
		})
		if err != nil {
			return nil, ab, err
		}
		enablePriceRegDest, err := state.Chains[source].PriceRegistry.ApplyDestChainConfigUpdates(
			SimTransactOpts(),
			[]price_registry.PriceRegistryDestChainConfigArgs{
				{
					DestChainSelector: newChainSel,
					DestChainConfig:   defaultPriceRegistryDestChainConfig(),
				},
			})
		if err != nil {
			return nil, ab, err
		}
		initialPrices, err := state.Chains[source].PriceRegistry.UpdatePrices(
			SimTransactOpts(),
			price_registry.InternalPriceUpdates{
				TokenPriceUpdates: []price_registry.InternalTokenPriceUpdate{},
				GasPriceUpdates: []price_registry.InternalGasPriceUpdate{
					{
						DestChainSelector: newChainSel,
						// TODO: parameterize
						UsdPerUnitGas: big.NewInt(2e12),
					},
				}})
		if err != nil {
			return nil, ab, err
		}
		batches = append(batches, managed.DetailedBatchChainOperation{
			ChainIdentifier: chain.Name,
			Batch: []managed.DetailedOperation{
				{
					// Enable the source in on ramp
					Operation: executable.Operation{
						To:    state.Chains[source].OnRamp.Address(),
						Data:  hexutil.Encode(enableOnRampDest.Data()),
						Value: 0,
					},
				},
				{
					// Set initial dest prices to unblock testing.
					Operation: executable.Operation{
						To:    state.Chains[source].PriceRegistry.Address(),
						Data:  hexutil.Encode(initialPrices.Data()),
						Value: 0,
					},
				},
				{
					// Set initial dest prices to unblock testing.
					Operation: executable.Operation{
						To:    state.Chains[source].PriceRegistry.Address(),
						Data:  hexutil.Encode(enablePriceRegDest.Data()),
						Value: 0,
					},
				},
			},
		})
		metaDataPerChain[chain.Name] = managed.MCMSWithTimelockChainMetadata{
			ExecutableMCMSChainMetadata: executable.ExecutableMCMSChainMetadata{
				NonceOffset: 0,
				MCMAddress:  state.Chains[source].McmAddr,
			},
			TimelockAddress: state.Chains[source].TimelockAddr,
		}
	}

	// Home chain new don.
	// - Add new DONs for destination to home chain
	nodes, err := deployment.NodeInfo(e.NodeIDs, e.Offchain)
	if err != nil {
		return nil, ab, err
	}
	newDONArgs, err := BuildAddDONArgs(e.Logger, state.Chains[newChainSel].OffRamp, e.Chains[newChainSel], nodes)
	if err != nil {
		return nil, ab, err
	}
	addDON, err := state.Chains[homeChainSel].CapabilityRegistry.AddDON(SimTransactOpts(),
		nodes.PeerIDs(newChainSel), []capabilities_registry.CapabilitiesRegistryCapabilityConfiguration{
			{
				CapabilityId: CCIPCapabilityId,
				Config:       newDONArgs,
			},
		}, false, false, nodes.DefaultF())
	if err != nil {
		return nil, ab, err
	}
	homeChain, _ := chainsel.ChainBySelector(homeChainSel)
	metaDataPerChain[homeChain.Name] = managed.MCMSWithTimelockChainMetadata{
		ExecutableMCMSChainMetadata: executable.ExecutableMCMSChainMetadata{
			NonceOffset: 0,
			MCMAddress:  state.Chains[homeChainSel].McmAddr,
		},
		TimelockAddress: state.Chains[homeChainSel].TimelockAddr,
	}
	batches = append(batches, managed.DetailedBatchChainOperation{
		ChainIdentifier: homeChain.Name,
		Batch: []managed.DetailedOperation{
			{
				// Enable the source in on ramp
				Operation: executable.Operation{
					To:    state.Chains[homeChainSel].CapabilityRegistry.Address(),
					Data:  hexutil.Encode(addDON.Data()),
					Value: 0,
				},
			},
		},
	})
	newDestProposal := managed.MCMSWithTimelockProposal{
		Operation:     managed.Schedule,
		MinDelay:      "1h",
		ChainMetadata: metaDataPerChain,
		Transactions:  batches,
	}

	// New chain we can configure directly with deployer key first.
	var offRampEnables []offramp.OffRampSourceChainConfigArgs
	for _, source := range sources {
		offRampEnables = append(offRampEnables, offramp.OffRampSourceChainConfigArgs{
			Router:              state.Chains[newChainSel].Router.Address(),
			SourceChainSelector: source,
			IsEnabled:           true,
			OnRamp:              common.LeftPadBytes(state.Chains[source].OnRamp.Address().Bytes(), 32),
		})
	}
	tx, err := state.Chains[newChainSel].OffRamp.ApplySourceChainConfigUpdates(e.Chains[newChainSel].DeployerKey, offRampEnables)
	if err := deployment.ConfirmIfNoError(e.Chains[newChainSel], tx, err); err != nil {
		return nil, ab, err
	}

	// We won't actually be able to setOCR3Config on the remote until the first proposal goes through.
	// TODO: Outbound
	return []managed.MCMSWithTimelockProposal{newDestProposal}, ab, nil
}

//func ApplyInboundChainProposal(
//	e deployment.Environment,
//	ab deployment.AddressBook,
//	proposal managed.MCMSWithTimelockProposal,
//) (deployment.AddressBook, error) {
//	state, err := LoadOnchainState(e, ab)
//	if err != nil {
//		return ab, err
//	}
//
//	// Apply the proposal.
//})

// 1. Deploy contracts
// 2. Proposal 1 (allow for inbound testing)
// - Enables new destination in onramps using test router
// - Enables the sources in the offramp and real router.
// - Sets initial prices for destination in price reg.
// - Add new DONs for destination to home chain
// - SetOCR3Config(s) on destination offramp.
// 3. At this point should be able to test from all sources
// and ensure that its writing those source prices to the new chain.
// 4. Proposal 2 (allow for outbound testing)
// -  Add new destinations on onramp/price reg can use real router.
// No initial prices needed because DON updating them.
// - Add new sources to the remote offramps (test router).
// - Add ChainConfig to home chain so existing OCR instances become aware of the source.
// 5. Now we can test the other direction.
// 6 . Proposal 3 move onramp/offramps on existing chains to real router.
