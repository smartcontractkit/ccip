package memory

type AddressBook struct {
	AddressesByChain map[uint64]map[string]struct{}
}

func (m *AddressBook) Save(chainSelector uint64, address string) error {
	if _, exists := m.AddressesByChain[chainSelector]; !exists {
		m.AddressesByChain[chainSelector] = make(map[string]struct{})
	} else {
		// Error?
	}
	m.AddressesByChain[chainSelector][address] = struct{}{}
	return nil
}

func (m *AddressBook) Addresses() (map[uint64]map[string]struct{}, error) {
	return m.AddressesByChain, nil
}

func NewMemoryAddressBook() *AddressBook {
	return &AddressBook{
		AddressesByChain: make(map[uint64]map[string]struct{}),
	}
}
