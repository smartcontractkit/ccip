package environment

type AddressBook interface {
	// TODO: Need manualTV override
	Save(chainSelector uint64, address string) error
	Addresses() (map[uint64]map[string]struct{}, error)
	// Allows for merging address books
	Merge(other AddressBook) error
}

type AddressBookMap struct {
	AddressesByChain map[uint64]map[string]struct{}
}

func (m *AddressBookMap) Save(chainSelector uint64, address string) error {
	if _, exists := m.AddressesByChain[chainSelector]; !exists {
		m.AddressesByChain[chainSelector] = make(map[string]struct{})
	} else {
		// Error?
	}
	m.AddressesByChain[chainSelector][address] = struct{}{}
	return nil
}

func (m *AddressBookMap) Addresses() (map[uint64]map[string]struct{}, error) {
	return m.AddressesByChain, nil
}

// Attention this will mutate existing book
func (m *AddressBookMap) Merge(ab AddressBook) error {
	addresses, err := ab.Addresses()
	if err != nil {
		return err
	}
	for chain, chainAddresses := range addresses {
		for address := range chainAddresses {
			return m.Save(chain, address)
		}
	}
	return nil
}

func NewMemoryAddressBook() *AddressBookMap {
	return &AddressBookMap{
		AddressesByChain: make(map[uint64]map[string]struct{}),
	}
}
