package launcher

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/keystone/generated/keystone_capability_registry"
	cctypes "github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/types"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/p2pkey"
)

type diffResult struct {
	added   map[uint32]keystone_capability_registry.CapabilityRegistryDONInfo
	removed map[uint32]keystone_capability_registry.CapabilityRegistryDONInfo
	updated map[uint32]keystone_capability_registry.CapabilityRegistryDONInfo
}

func diff(
	capabilityVersion,
	capabilityLabelledName string,
	oldState,
	newState cctypes.RegistryState,
) (diffResult, error) {
	ccipCapability, err := checkCapabilityPresence(capabilityVersion, capabilityLabelledName, newState)
	if err != nil {
		return diffResult{}, err
	}

	newCCIPDONs, err := filterCCIPDONs(ccipCapability, newState)
	if err != nil {
		return diffResult{}, err
	}

	currCCIPDONs, err := filterCCIPDONs(ccipCapability, oldState)
	if err != nil {
		return diffResult{}, err
	}

	// compare curr with new and launch or update OCR instances as needed
	added, removed, updated, err := compareDONs(currCCIPDONs, newCCIPDONs)
	if err != nil {
		return diffResult{}, err
	}

	return diffResult{
		added:   added,
		removed: removed,
		updated: updated,
	}, nil
}

func compareDONs(
	currCCIPDONs,
	newCCIPDONs map[uint32]keystone_capability_registry.CapabilityRegistryDONInfo,
) (
	added, removed, updated map[uint32]keystone_capability_registry.CapabilityRegistryDONInfo,
	err error,
) {
	added = make(map[uint32]keystone_capability_registry.CapabilityRegistryDONInfo)
	removed = make(map[uint32]keystone_capability_registry.CapabilityRegistryDONInfo)
	updated = make(map[uint32]keystone_capability_registry.CapabilityRegistryDONInfo)

	for id, don := range newCCIPDONs {
		if currDONState, ok := currCCIPDONs[id]; !ok {
			// Not in current state, so mark as added.
			added[id] = don
		} else {
			// If its in the current state and the config count for the DON has changed, mark as updated.
			// Since the registry returns the full state we need to compare the config count.
			if don.ConfigCount > currDONState.ConfigCount {
				updated[id] = don
			}
		}
	}

	for id, don := range currCCIPDONs {
		if _, ok := newCCIPDONs[id]; !ok {
			// In current state but not in latest registry state, so should remove.
			removed[id] = don
		}
	}

	return added, removed, updated, nil
}

func filterCCIPDONs(
	ccipCapability keystone_capability_registry.CapabilityRegistryCapability,
	state cctypes.RegistryState,
) (map[uint32]keystone_capability_registry.CapabilityRegistryDONInfo, error) {
	ccipDONs := make(map[uint32]keystone_capability_registry.CapabilityRegistryDONInfo)
	for _, don := range state.DONs {
		// CCIP DONs should only have one capability, CCIP.
		var found bool
		for _, donCapabilities := range don.CapabilityConfigurations {
			if donCapabilities.CapabilityId == hashedCapabilityId(ccipCapability.Version, ccipCapability.LabelledName) {
				ccipDONs[don.Id] = don
				found = true
			}
		}
		if found && len(don.CapabilityConfigurations) > 1 {
			return nil, fmt.Errorf("found more than one capability (actual: %d) in the CCIP DON %d",
				len(don.CapabilityConfigurations), don.Id)
		}
	}

	return ccipDONs, nil
}

func checkCapabilityPresence(
	capabilityVersion,
	capabilityLabelledName string,
	state cctypes.RegistryState,
) (keystone_capability_registry.CapabilityRegistryCapability, error) {
	// Sanity check to make sure the capability registry has the capability we are looking for.
	var ccipCapability keystone_capability_registry.CapabilityRegistryCapability
	for _, capability := range state.Capabilities {
		if capability.LabelledName == capabilityLabelledName &&
			capability.Version == capabilityVersion {
			ccipCapability = capability
			break
		}
	}

	if ccipCapability.LabelledName == "" {
		return keystone_capability_registry.CapabilityRegistryCapability{},
			fmt.Errorf("unable to find capability with name %s and version %s in capability registry state",
				capabilityLabelledName, capabilityVersion)
	}

	return ccipCapability, nil
}

func hashedCapabilityId(capabilityVersion, capabilityLabelledName string) (r [32]byte) {
	capVersionBytes := []byte(capabilityVersion)
	capLabelledNameBytes := []byte(capabilityLabelledName)
	var capVersionBytes32, capLabelledNameBytes32 [32]byte
	copy(capVersionBytes32[:], capVersionBytes)
	copy(capLabelledNameBytes32[:], capLabelledNameBytes)
	h := crypto.Keccak256(capVersionBytes32[:], capLabelledNameBytes32[:])
	copy(r[:], h)
	return r
}

// isMemberOfDON returns true if and only if the given p2pID is a member of the given DON.
func isMemberOfDON(don keystone_capability_registry.CapabilityRegistryDONInfo, p2pID p2pkey.KeyV2) bool {
	var found bool
	for _, node := range don.NodeP2PIds {
		if node == p2pID.PeerID() {
			found = true
			break
		}
	}
	return found
}
