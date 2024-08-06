package deployment

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	chainsel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink/v2/core/environment"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/arm_proxy_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/nonce_manager"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/token_admin_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/weth9"
	type_and_version "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/type_and_version_interface_wrapper"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/capabilities_registry"
)

// Onchain state always derivable from an address book.
// Offchain state always derivable from a list of nodeIds.
// Note can translate this into Go struct needed for MCMS/Docs/UI.
type CCIPOnChainState struct {
	// Populated go bindings for the appropriate version for all contracts.
	// We would hold 2 versions of each contract here. Once we upgrade we can phase out the old one.
	// When generating bindings, make sure the package name corresponds to the version.
	EvmOnRampsV160       map[uint64]*evm_2_evm_multi_onramp.EVM2EVMMultiOnRamp
	EvmOffRampsV160      map[uint64]*evm_2_evm_multi_offramp.EVM2EVMMultiOffRamp
	PriceRegistries      map[uint64]*price_registry.PriceRegistry
	ArmProxies           map[uint64]*arm_proxy_contract.ARMProxyContract
	NonceManagers        map[uint64]*nonce_manager.NonceManager
	TokenAdminRegistries map[uint64]*token_admin_registry.TokenAdminRegistry
	Routers              map[uint64]*router.Router
	Weth9s               map[uint64]*weth9.WETH9

	// Only lives on the home chain.
	CapabilityRegistry *capabilities_registry.CapabilitiesRegistry
}

type CCIPSnapShot struct {
	Chains map[string]Chain `json:"chains"`
}

type Chain struct {
	TokenAdminRegistry       common.Address   `json:"tokenAdminRegistry"`
	TokenAdminRegistryTokens []common.Address `json:"tokenAdminRegistryTokens"`
}

func (s CCIPOnChainState) Snapshot(chains []uint64) (CCIPSnapShot, error) {
	snapshot := CCIPSnapShot{
		Chains: make(map[string]Chain),
	}
	for _, chainSelector := range chains {
		chainid, _ := chainsel.ChainIdFromSelector(chainSelector)
		chainName, _ := chainsel.NameFromChainId(chainid)
		var c Chain
		if ta, ok := s.TokenAdminRegistries[chainSelector]; ok {
			tokens, err := ta.GetAllConfiguredTokens(nil, 0, 10)
			if err != nil {
				return snapshot, err
			}
			c.TokenAdminRegistry = ta.Address()
			c.TokenAdminRegistryTokens = tokens
		}
		snapshot.Chains[chainName] = c
	}
	return snapshot, nil
}

func SnapshotState(e environment.Environment, ab environment.AddressBook) (CCIPSnapShot, error) {
	state, err := GenerateOnchainState(e, ab)
	if err != nil {
		return CCIPSnapShot{}, err
	}
	return state.Snapshot(e.AllChainSelectors())
}

func GenerateOnchainState(e environment.Environment, ab environment.AddressBook) (CCIPOnChainState, error) {
	state := CCIPOnChainState{
		EvmOnRampsV160:       make(map[uint64]*evm_2_evm_multi_onramp.EVM2EVMMultiOnRamp),
		EvmOffRampsV160:      make(map[uint64]*evm_2_evm_multi_offramp.EVM2EVMMultiOffRamp),
		PriceRegistries:      make(map[uint64]*price_registry.PriceRegistry),
		ArmProxies:           make(map[uint64]*arm_proxy_contract.ARMProxyContract),
		NonceManagers:        make(map[uint64]*nonce_manager.NonceManager),
		TokenAdminRegistries: make(map[uint64]*token_admin_registry.TokenAdminRegistry),
		Routers:              make(map[uint64]*router.Router),
		Weth9s:               make(map[uint64]*weth9.WETH9),
	}
	// Get all the onchain state
	addresses, err := ab.Addresses()
	if err != nil {
		return state, errors.Wrap(err, "could not get addresses")
	}
	for chainSelector, addresses := range addresses {
		for address := range addresses {
			tv, err := type_and_version.NewTypeAndVersionInterface(common.HexToAddress(address), e.Chains[chainSelector].Client)
			if err != nil {
				return state, errors.Wrap(err, "could not create tv interface")
			}
			tvStr, err := tv.TypeAndVersion(nil)
			if err != nil {
				// TODO: there are some contracts which dont like the link token/weth9
				return state, errors.Wrap(err, fmt.Sprintf("could not call tv version, does the contract %s implement it?", address))
			}
			switch tvStr {
			case "CapabilitiesRegistry 1.0.0":
				cr, err := capabilities_registry.NewCapabilitiesRegistry(common.HexToAddress(address), e.Chains[chainSelector].Client)
				if err != nil {
					return state, err
				}
				state.CapabilityRegistry = cr
			case "EVM2EVMMultiOnRamp 1.6.0-dev":
				onRamp, err := evm_2_evm_multi_onramp.NewEVM2EVMMultiOnRamp(common.HexToAddress(address), e.Chains[chainSelector].Client)
				if err != nil {
					return state, err
				}
				state.EvmOnRampsV160[chainSelector] = onRamp
			case "EVM2EVMMultiOffRamp 1.6.0-dev":
				offRamp, err := evm_2_evm_multi_offramp.NewEVM2EVMMultiOffRamp(common.HexToAddress(address), e.Chains[chainSelector].Client)
				if err != nil {
					return state, err
				}
				state.EvmOffRampsV160[chainSelector] = offRamp
			case "ARMProxy 1.0.0":
				armProxy, err := arm_proxy_contract.NewARMProxyContract(common.HexToAddress(address), e.Chains[chainSelector].Client)
				if err != nil {
					return state, err
				}
				state.ArmProxies[chainSelector] = armProxy
			case "NonceManager 1.6.0-dev":
				nm, err := nonce_manager.NewNonceManager(common.HexToAddress(address), e.Chains[chainSelector].Client)
				if err != nil {
					return state, err
				}
				state.NonceManagers[chainSelector] = nm
			case "TokenAdminRegistry 1.5.0-dev":
				tm, err := token_admin_registry.NewTokenAdminRegistry(common.HexToAddress(address), e.Chains[chainSelector].Client)
				if err != nil {
					return state, err
				}
				state.TokenAdminRegistries[chainSelector] = tm
			case "Router 1.2.0":
				r, err := router.NewRouter(common.HexToAddress(address), e.Chains[chainSelector].Client)
				if err != nil {
					return state, err
				}
				state.Routers[chainSelector] = r
			case "PriceRegistry 1.6.0-dev":
				pr, err := price_registry.NewPriceRegistry(common.HexToAddress(address), e.Chains[chainSelector].Client)
				if err != nil {
					return state, err
				}
				state.PriceRegistries[chainSelector] = pr
			default:
				return state, fmt.Errorf("unknown contract %s", tvStr)
			}
		}
	}
	return state, nil
}
