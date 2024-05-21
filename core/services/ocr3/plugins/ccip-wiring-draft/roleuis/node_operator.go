package roleuis

type NodeOperator interface {
	// AddCapability allows a node operator to add a new capability to his stack.
	// The capability is defined in the provided configuration and should be some
	// capability that the core node already supports. i.e. already added by the CapabilityAuthor
	AddCapability(capabilityID string, cfg []byte) (id int, err error)

	// PublishNode allows a node operator to publish a new node that will run the provided capability.
	// This will add the new node in the capability registry in ethereum mainnet
	// contract 'Canonical Capability Registry' and 'NodeAdded' event is emitted.
	// The capability and it's config must already exist. i.e. added by the AddCapability method.
	PublishNode(capabilityID int, cfg []byte) error
}
