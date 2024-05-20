package corenode

type DonManagerUI interface {
	// CreateDon allows a node operator to create a new DON.
	// This will register all the provided dons in the canonical capability registry in eth mainnet.
	// emits 'DONCreated'.
	CreateDon(cfg []byte) error

	// UpdateDON
	// RemoveDON
	// ...
}

type CapabilityRegistryTracker interface {
	FilterDONCreatedEvents() error
	FilterDONUpdatedEvents() error
	FilterDONRemovedEvents() error
	FilterCapabilityConfigSetEvents() error
}
