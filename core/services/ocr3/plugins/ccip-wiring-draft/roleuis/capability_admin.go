package roleuis

type CapabilityAdmin interface {
	// SetConfig will set the configuration for the capability with the given ID.
	SetConfig(capabilityID string, addr string, cfg []byte) error
}
