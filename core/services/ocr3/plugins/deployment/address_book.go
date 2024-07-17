package deployments

// For persistent environments (testnet/mainnet, we can have a persistent address book
// that just writes to a file after every addition).
type ContractAddressBook interface {
	Save(chainSelector uint64, address string) error
	Addresses() map[uint64]map[string]struct{}
}

type MemoryAddressBook struct {
	AddressesByChain map[uint64]map[string]struct{}
}

func (m *MemoryAddressBook) Save(chainSelector uint64, address string) error {
	if _, exists := m.AddressesByChain[chainSelector]; !exists {
		m.AddressesByChain[chainSelector] = make(map[string]struct{})
	} else {
		// Error?
	}
	m.AddressesByChain[chainSelector][address] = struct{}{}
	return nil
}

func (m *MemoryAddressBook) Addresses() map[uint64]map[string]struct{} {
	return m.AddressesByChain
}

func NewMemoryAddressBook() *MemoryAddressBook {
	return &MemoryAddressBook{
		AddressesByChain: make(map[uint64]map[string]struct{}),
	}
}

type FileAddressBook struct {
}

func (f FileAddressBook) Save(chainSelector uint64, address string) error {
	//TODO implement me
	panic("implement me")
}

func (f FileAddressBook) Addresses() map[uint64]map[string]struct{} {
	//TODO implement me
	panic("implement me")
}

func NewFileAddressBook(file string) *FileAddressBook {
	return &FileAddressBook{}
}
