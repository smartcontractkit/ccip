package ccip_integration_tests

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

const chainID = 1337

type TestSetupData[T any] struct {
	ContractAddr common.Address
	Contract     *T
	SimulatedBE  *backends.SimulatedBackend
	Auth         *bind.TransactOpts
}

//func (data *TestSetupData[T]) GetContract() (*T, error) {
//	contract, ok := data.contract.(*T)
//	if !ok {
//		return nil, fmt.Errorf("failed to cast contract to the expected type")
//	}
//	return contract, nil
//}

//func GetContract[T any](data *TestSetupData) (*T, error) {
//	contract, ok := data.contract.(*T)
//	if !ok {
//		return nil, fmt.Errorf("failed to cast contract to the expected type")
//	}
//	return contract, nil
//}

type DeployFunc[T any] func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *T, error)

type DeployFuncWithCapabilities[T any] func(auth *bind.TransactOpts, backend bind.ContractBackend, capabilityRegistry common.Address) (common.Address, *types.Transaction, *T, error)

type NewFunc[T any] func(address common.Address, backend bind.ContractBackend) (*T, error)

func SetupTest[T any](t *testing.T, _ context.Context, deployFunc DeployFunc[T], newFunc NewFunc[T]) *TestSetupData[T] {
	// Generate a new key pair for the simulated account
	privateKey, err := crypto.GenerateKey()
	assert.NoError(t, err)
	// Set up the genesis account with balance
	blnc, ok := big.NewInt(0).SetString("999999999999999999999999999999999999", 10)
	assert.True(t, ok)
	alloc := map[common.Address]core.GenesisAccount{crypto.PubkeyToAddress(privateKey.PublicKey): {Balance: blnc}}
	simulatedBackend := backends.NewSimulatedBackend(alloc, 0)
	// Create a transactor
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	assert.NoError(t, err)
	auth.GasLimit = uint64(0)

	// Deploy the contract using the provided deployFunc
	address, tx, _, err := deployFunc(auth, simulatedBackend)
	assert.NoError(t, err)
	simulatedBackend.Commit()

	t.Logf("contract deployed: addr=%s tx=%s", address.Hex(), tx.Hash())

	// Setup contract client using the provided newFunc
	contract, err := newFunc(address, simulatedBackend)
	assert.NoError(t, err)

	return &TestSetupData[T]{
		ContractAddr: address,
		Contract:     contract,
		SimulatedBE:  simulatedBackend,
		Auth:         auth,
	}
}

func SetupTestWithCapability[T any](t *testing.T, _ context.Context, deployFunc DeployFuncWithCapabilities[T], newFunc NewFunc[T], capabilityRegistry common.Address) *TestSetupData[T] {
	// Generate a new key pair for the simulated account
	privateKey, err := crypto.GenerateKey()
	assert.NoError(t, err)
	// Set up the genesis account with balance
	blnc, ok := big.NewInt(0).SetString("999999999999999999999999999999999999", 10)
	assert.True(t, ok)
	alloc := map[common.Address]core.GenesisAccount{crypto.PubkeyToAddress(privateKey.PublicKey): {Balance: blnc}}
	simulatedBackend := backends.NewSimulatedBackend(alloc, 0)
	// Create a transactor
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	assert.NoError(t, err)
	auth.GasLimit = uint64(0)

	// Deploy the contract using the provided deployFunc
	address, tx, _, err := deployFunc(auth, simulatedBackend, capabilityRegistry)
	assert.NoError(t, err)
	simulatedBackend.Commit()

	t.Logf("contract deployed: addr=%s tx=%s", address.Hex(), tx.Hash())

	// Setup contract client using the provided newFunc
	contract, err := newFunc(address, simulatedBackend)
	assert.NoError(t, err)

	return &TestSetupData[T]{
		ContractAddr: address,
		Contract:     contract,
		SimulatedBE:  simulatedBackend,
		Auth:         auth,
	}
}
