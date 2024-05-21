package roleuis

type CapabilityAuthor interface {
	// DeployConfigContract deploys a new configuration contract for the given capability.
	DeployConfigContract(capabilityID string, capabilityConfig []byte) (addr string, err error)

	// AddCapabilityToRegistry adds a capability to the canonical capability registry.
	// emits 'CapabilityAdded(capabilityID)'
	AddCapabilityToRegistry(capabilityID string, addr string, cfg []byte) error
}
