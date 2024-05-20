package capabilityregistry

type CapabilityAuthorUI interface {
	// DeployConfigContract will deploy the config contract that contains the config of the capability.
	DeployConfigContract(cap Capability) error

	// AddCapability can be called after capability config contract is deployed to add the capability
	// to the canonical capability registry. `CapabilityAdded` event is emitted.
	// The capability is now ready to be used by nodes.
	AddCapability(cap Capability) error
}
