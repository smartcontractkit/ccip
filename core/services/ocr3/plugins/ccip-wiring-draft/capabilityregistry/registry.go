package capabilityregistry

type CanonicalCapabilityRegistry interface {
	AddDON(don DON) error    // emit 'DONCreated'
	RemoveDON(id int) error  // emit 'DONRemoved'
	UpdateDON(don DON) error // emit 'DONUpdated'
	GetDON(id int) (DON, error)
	GetAllDons() ([]DON, error)
}

type CapabilityConfigManager interface {
	SetConfig(capID string, config []byte) error // emit 'CapabilityConfigurationSet'
}

type Registry struct {
	dons []DON
}

type DON struct {
	ID           int
	IsPublic     bool
	Nodes        []Node
	Capabilities []Capability // cannot have the same capability twice
}

type Node struct {
	Address        string
	NodeOperatorID int
	P2PID          [32]byte
	CapabilityIDs  []string
}

type NodeOperator struct {
	Name      string
	AdminAddr string
	Nodes     []Node
}

// @contract
type Capability struct {
	Config       []byte
	Typ          string
	Version      string
	Removed      bool
	ResponseType string // REPORT or OBSERVATION_IDENTICAL

	// emits CapabilityConfigurationSet()
	// getCapabilityConfiguration()
}
