package persistent

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"sync"
)

type AddressBook struct {
	filePath string
	mu       sync.Mutex
}

// Save stores an address in the address book.
func (m *AddressBook) Save(chainSelector uint64, address string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	addressesByChain, err := m.loadFromFile()
	if err != nil {
		return err
	}

	if _, exists := addressesByChain[chainSelector]; !exists {
		addressesByChain[chainSelector] = make(map[string]struct{})
	}
	addressesByChain[chainSelector][address] = struct{}{}

	return m.saveToFile(addressesByChain)
}

// Addresses returns all addresses.
func (m *AddressBook) Addresses() (map[uint64]map[string]struct{}, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.loadFromFile()
}

// saveToFile writes the address book to the file.
func (m *AddressBook) saveToFile(addressesByChain map[uint64]map[string]struct{}) error {
	// Make json friendly
	addressLists := make(map[uint64][]string)
	for chain, addressMp := range addressesByChain {
		for addr := range addressMp {
			addressLists[chain] = append(addressLists[chain], addr)
		}
	}
	data, err := json.Marshal(addressLists)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(m.filePath, data, 0644)
}

// loadFromFile loads the address book from the file.
func (m *AddressBook) loadFromFile() (map[uint64]map[string]struct{}, error) {
	file, err := os.Open(m.filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return make(map[uint64]map[string]struct{}), err
		}
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	addressLists := make(map[uint64][]string)
	err = json.Unmarshal(data, &addressLists)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to unmarshal address book %s %s", m.filePath, data))
	}
	addressesByChain := make(map[uint64]map[string]struct{})
	for chain, addresses := range addressLists {
		addressesByChain[chain] = make(map[string]struct{})
		for _, address := range addresses {
			addressesByChain[chain][address] = struct{}{}
		}
	}
	return addressesByChain, nil
}

// NewFileBackedAddressBook creates a new AddressBook with file storage.
func NewAddressBook(filePath string) *AddressBook {
	return &AddressBook{
		filePath: filePath,
	}
}
