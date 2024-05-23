package ccip

import "context"

type CapabilityRegistry interface {
	// GetDONsWithCapability returns a map of all the DONs that have the given capability.
	GetDONsWithCapability(ctx context.Context, capabilityID string) (map[DonID]DONCapability, error)
}

type DonID uint32

type DONCapability struct {
	ID       DonID
	IsPublic bool
	Nodes    [][]byte
	Config   CapabilityConfiguration
}

type CapabilityConfiguration struct {
	CapabilityID string
	Config       []byte
}
