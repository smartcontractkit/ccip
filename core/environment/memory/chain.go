package memory

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

type EVMChain struct {
	Backend     *backends.SimulatedBackend
	DeployerKey *bind.TransactOpts
}

func fundAddress(t *testing.T, from *bind.TransactOpts, to common.Address, amount *big.Int, backend *backends.SimulatedBackend) {
	nonce, err := backend.PendingNonceAt(testutils.Context(t), from.From)
	require.NoError(t, err)
	gp, err := backend.SuggestGasPrice(testutils.Context(t))
	require.NoError(t, err)
	rawTx := gethtypes.NewTx(&gethtypes.LegacyTx{
		Nonce:    nonce,
		GasPrice: gp,
		Gas:      21000,
		To:       &to,
		Value:    amount,
	})
	signedTx, err := from.Signer(from.From, rawTx)
	require.NoError(t, err)
	err = backend.SendTransaction(testutils.Context(t), signedTx)
	require.NoError(t, err)
	backend.Commit()
}

func GenerateChains(t *testing.T, numChains int) map[uint64]EVMChain {
	chains := make(map[uint64]EVMChain)
	for i := 0; i < numChains; i++ {
		chainID := chainsel.TEST_90000001.EvmChainID + uint64(i)
		owner := testutils.MustNewSimTransactor(t)
		backend := backends.NewSimulatedBackend(core.GenesisAlloc{
			owner.From: {Balance: big.NewInt(0).Mul(big.NewInt(100), big.NewInt(params.Ether))}}, 10000000)
		chains[chainID] = EVMChain{
			Backend:     backend,
			DeployerKey: owner,
		}
	}
	return chains
}
