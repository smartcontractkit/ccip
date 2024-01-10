package ccipdata

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	lpmocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
)

// NewSimulation returns a client and a simulated backend.
func NewSimulation(t *testing.T) (*bind.TransactOpts, *client.SimulatedBackendClient) {
	user := testutils.MustNewSimTransactor(t)
	simulatedBackend := backends.NewSimulatedBackend(map[common.Address]core.GenesisAccount{
		user.From: {
			Balance: big.NewInt(0).Mul(big.NewInt(3), big.NewInt(1e18)),
		},
	}, 10e6)
	simulatedBackendClient := client.NewSimulatedBackendClient(t, simulatedBackend, testutils.SimulatedChainID)
	return user, simulatedBackendClient
}

// AssertNonRevert Verify that a transaction was not reverted.
func AssertNonRevert(t *testing.T, tx *types.Transaction, bc *client.SimulatedBackendClient, user *bind.TransactOpts) {
	require.NotNil(t, tx, "Transaction should not be nil")
	receipt, err := bc.TransactionReceipt(user.Context, tx.Hash())
	require.NoError(t, err)
	require.NotEqual(t, uint64(0), receipt.Status, "Transaction should not have reverted")
}

func AssertFilterRegistration(t *testing.T, lp *lpmocks.LogPoller, buildCloser func(lp *lpmocks.LogPoller, addr common.Address) Closer, numFilter int) {
	// Expected filter properties for a closer:
	// - Should be the same filter set registered that is unregistered
	// - Should be registered to the address specified
	// - Number of events specific to this component should be registered
	addr := common.HexToAddress("0x1234")
	var filters []logpoller.Filter

	lp.On("RegisterFilter", mock.Anything).Run(func(args mock.Arguments) {
		f := args.Get(0).(logpoller.Filter)
		require.Equal(t, len(f.Addresses), 1)
		require.Equal(t, f.Addresses[0], addr)
		filters = append(filters, f)
	}).Return(nil).Times(numFilter)

	c := buildCloser(lp, addr)
	for _, filter := range filters {
		lp.On("UnregisterFilter", filter.Name).Return(nil)
	}

	require.NoError(t, c.Close())
	lp.AssertExpectations(t)
}

// AssertAllFieldsAreRequired asserts that setting any struct field to zero value invalidates the value.
// NOTE: the value argument must be valid.
func AssertAllFieldsAreRequired[T config.OffchainConfig](t *testing.T, value T) {
	// Iterate over all fields in the struct using the reflection API.
	// Make a copy of the value and try setting the corresponding field to zero value.
	// Then try to validate the value and assert that:
	//   1. The modified copy became invalid.
	//   2. The validation error contains the field name.
	require.NoError(t, value.Validate())

	src := reflect.ValueOf(value)
	if src.Kind() != reflect.Struct {
		t.Fatalf("expected struct, got %s", src.Kind())
	}
	fields := reflect.VisibleFields(src.Type())
	for _, field := range fields {
		valueCopy := value
		dst := reflect.ValueOf(&valueCopy)
		f := src.FieldByIndex(field.Index)
		dst.Elem().FieldByIndex(field.Index).Set(reflect.Zero(f.Type()))
		err := valueCopy.Validate()
		require.Contains(t, err.Error(), field.Name)
	}
}
