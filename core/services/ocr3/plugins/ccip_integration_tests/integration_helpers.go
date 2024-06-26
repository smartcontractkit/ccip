package ccip_integration_tests

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	types2 "github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	logger2 "github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
	"github.com/stretchr/testify/assert"
)

const chainID = 1337

type TestSetupData[T any] struct {
	ContractAddr common.Address
	Contract     *T
	SimulatedBE  *backends.SimulatedBackend
	Auth         *bind.TransactOpts
	ChainReader  *evm.ChainReaderService
	ChainID      int
}

type DeployFunc[T any] func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *T, error)

type NewFunc[T any] func(address common.Address, backend bind.ContractBackend) (*T, error)

func SetupBackendWithAuth(t *testing.T) (*backends.SimulatedBackend, *bind.TransactOpts) {
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

	return simulatedBackend, auth
}

func SetupChainReader(t *testing.T, simulatedBackend *backends.SimulatedBackend, address common.Address, chainReaderConfig evmtypes.ChainReaderConfig, contractName string) *evm.ChainReaderService {
	lggr := logger2.NullLogger
	db := pgtest.NewSqlxDB(t)
	lpOpts := logpoller.Opts{
		PollPeriod:               time.Millisecond,
		FinalityDepth:            1,
		BackfillBatchSize:        1,
		RpcBatchSize:             1,
		KeepFinalizedBlocksDepth: 10000,
	}
	cl := client.NewSimulatedBackendClient(t, simulatedBackend, big.NewInt(chainID))
	lp := logpoller.NewLogPoller(logpoller.NewORM(big.NewInt(chainID), db, lggr), cl, lggr, lpOpts)
	assert.NoError(t, lp.Start(context.Background()))

	cr, err := evm.NewChainReaderService(context.Background(), lggr, lp, cl, chainReaderConfig)
	assert.NoError(t, err)
	err = cr.Bind(context.Background(), []types2.BoundContract{
		{
			Address: address.String(),
			Name:    contractName,
			Pending: false,
		},
	})
	assert.NoError(t, err)

	err = cr.Start(context.Background())
	assert.NoError(t, err)
	for {
		if err := cr.Ready(); err == nil {
			break
		}
	}
	simulatedBackend.Commit()
	return &cr
}
