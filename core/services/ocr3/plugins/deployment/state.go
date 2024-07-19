package deployments

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_onramp"
	type_and_version "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/type_and_version_interface_wrapper"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/capabilities_registry"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/deployment/jobdistributor"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/deployment/ownerhelpers/manychainmultisig"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/deployment/ownerhelpers/rbactimelock"
	"reflect"
)

// Onchain state always derivable from an address book.
// Offchain state always derivable from a list of nodeIds.
// Note can translate this into Go struct needed for MCMS/Docs/UI.
type CCIPOnChainState struct {
	// Populated go bindings for the appropriate version for all contracts.
	// We would hold 2 versions of each contract here. Once we upgrade we can phase out the old one.
	// When generating bindings, make sure the package name corresponds to the version.
	EvmOnRampsV160 map[uint64]*evm_2_evm_multi_onramp.EVM2EVMMultiOnRamp
	// TODO: all other CCIP contracts.
	// Analogous to https://github.com/smartcontractkit/rddtool-ccip/blob/d975a799e1981e98c350684365f8e32a74881803/types/cross_chain/ccipState.go#L15
	// except its all derivable from the chain using the bare minimum persisted offchain.
	Timelocks map[uint64]rbactimelock.Rbactimelock
	Mcms      map[uint64]manychainmultisig.Manychainmultisig

	// Only lives on the home chain.
	CapabilityRegistry *capabilities_registry.CapabilitiesRegistry
}

type CCIPOffChainState struct {
	// Offchain state
	NodesToJobSpecs map[string][]CCIPSpec
}

type CCIPState struct {
	CCIPOnChainState
	CCIPOffChainState
}

type CCIPSpec struct {
	// Can use core node type here I think
	CapabilityRegistry common.Address
}

func (CCIPSpec) ToTOML() string {
	return ""
}

func NewCCIPSpecFromTOML(toml string) CCIPSpec {
	return CCIPSpec{}
}

// Serialize from generated go bindings. Avoids having to define custom structs for each version.
func toJSON(v interface{}) (string, error) {
	vType := reflect.TypeOf(v)
	vValue := reflect.ValueOf(v)

	// Create a new struct type with the same fields and add JSON tags
	fields := make([]reflect.StructField, vType.NumField())
	for i := 0; i < vType.NumField(); i++ {
		field := vType.Field(i)
		field.Tag = reflect.StructTag(fmt.Sprintf(`json:"%s"`, field.Name))
		fields[i] = field
	}

	// Create a new struct with the fields
	newStructType := reflect.StructOf(fields)
	newStructValue := reflect.New(newStructType).Elem()

	// Copy values from the original struct to the new struct
	for i := 0; i < vType.NumField(); i++ {
		newStructValue.Field(i).Set(vValue.Field(i))
	}

	// Marshal the new struct to JSON
	jsonData, err := json.Marshal(newStructValue.Interface())
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func GenerateOnchainState(chains map[uint64]Chain, addressBook ContractAddressBook) (CCIPOnChainState, error) {
	var state CCIPOnChainState
	// Get all the onchain state
	for chainSelector, addresses := range addressBook.Addresses() {
		for address := range addresses {
			// we assume all contract support the type and version interface.
			// this allow us to load the appropriate binding.
			// TODO: make this family agnostic.
			tv, err := type_and_version.NewTypeAndVersionInterface(common.HexToAddress(address), chains[chainSelector].Client)
			if err != nil {
				return state, err
			}
			tvStr, err := tv.TypeAndVersion(nil)
			if err != nil {
				return state, err
			}
			switch tvStr {
			case "CapabilitiesRegistry 1.0.0":
				cr, err := capabilities_registry.NewCapabilitiesRegistry(common.HexToAddress(address), chains[chainSelector].Client)
				if err != nil {
					return state, err
				}
				state.CapabilityRegistry = cr
			case "EVM2MultiOnRamp 1.6.0":
				onRamp, err := evm_2_evm_multi_onramp.NewEVM2EVMMultiOnRamp(common.HexToAddress(address), chains[chainSelector].Client)
				if err != nil {
					return state, err
				}
				state.EvmOnRampsV160[chainSelector] = onRamp
				// TODO: add each contract as needed.
			default:
				return state, fmt.Errorf("unknown contract %s", tv)
			}
		}
	}
	return state, nil
}

func generateOffchainState(nodeIds []string, client jobdistributor.JobServiceClient) (CCIPOffChainState, error) {
	var state CCIPOffChainState
	// Get all the offchain state.
	jobs, err := client.ListJobs(context.Background(), &jobdistributor.ListJobsRequest{Filter: &jobdistributor.ListJobsRequest_Filter{NodeIds: nodeIds}})
	if err != nil {
		return CCIPOffChainState{}, err
	}
	// Look up associate proposal for each job to get the spec.
	var jobIds []string
	jobsToNodes := make(map[string]string)
	for _, job := range jobs.Jobs {
		jobsToNodes[job.NodeId] = job.Id
		jobIds = append(jobIds, job.Id)
	}
	proposals, err := client.ListProposals(context.Background(), &jobdistributor.ListProposalsRequest{Filter: &jobdistributor.ListProposalsRequest_Filter{JobIds: jobIds}})
	if err != nil {
		return CCIPOffChainState{}, err
	}
	for _, proposal := range proposals.Proposals {
		spec := NewCCIPSpecFromTOML(proposal.Spec)
		// TODO: should probably include nodeId in the proposal?
		state.NodesToJobSpecs[jobsToNodes[proposal.JobId]] = append(state.NodesToJobSpecs[jobsToNodes[proposal.JobId]], spec)
	}
	return state, nil

}

func generateState(rpcs map[uint64]Chain, addressBook ContractAddressBook, nodeIds []string, client jobdistributor.JobServiceClient) (CCIPState, error) {
	var state CCIPState
	onChainState, err := GenerateOnchainState(rpcs, addressBook)
	if err != nil {
		return state, err
	}
	offChainState, err := generateOffchainState(nodeIds, client)
	if err != nil {
		return state, err
	}
	return CCIPState{CCIPOnChainState: onChainState, CCIPOffChainState: offChainState}, nil
}
