package roleuis

type DonManager interface {
	// CreateDon allows a node operator to create a new DON.
	// This will register all the provided dons in the canonical capability registry in eth mainnet.
	// emits 'DONCreated'.
	CreateDon(cfg []byte) error
	UpdateDon(cfg []byte) error
	RemoveDon(cfg []byte) error
}
