package ccipdata

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
)

// NewSimulation returns a client and a simulated backend.
func NewSimulation(t *testing.T) (*bind.TransactOpts, *client.SimulatedBackendClient) {
	user := testutils.MustNewSimTransactor(t)
	sim := backends.NewSimulatedBackend(map[common.Address]core.GenesisAccount{
		user.From: {
			Balance: big.NewInt(0).Mul(big.NewInt(3), big.NewInt(1e18)),
		},
	}, 10e6)
	balance, err := sim.BalanceAt(user.Context, user.From, nil)
	require.NoError(t, err)
	require.Equal(t, big.NewInt(0).Mul(big.NewInt(3), big.NewInt(1e18)), balance)
	ec := client.NewSimulatedBackendClient(t, sim, testutils.SimulatedChainID)
	return user, ec
}

// AssertNonRevert Verify that a transaction was not reverted.
func AssertNonRevert(t *testing.T, tx *types.Transaction, bc *client.SimulatedBackendClient, user *bind.TransactOpts) {
	receipt, err := bc.TransactionReceipt(user.Context, tx.Hash())
	require.NoError(t, err)
	require.NotEqual(t, uint64(0), receipt.Status, "Transaction should not have reverted")
}
