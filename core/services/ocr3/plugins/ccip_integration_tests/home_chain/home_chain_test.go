package home_chain

import (
	"context"
	_ "embed"
	"math/big"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	types2 "github.com/smartcontractkit/chainlink-common/pkg/types"
	query2 "github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	logger2 "github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
	"github.com/stretchr/testify/assert"
)

/*
 *
 * solc --abi --bin ../../../../../../contracts/src/v0.8/ccip/capability/CCIPCapabilityConfiguration.sol -o build # use solc 0.8.24
 * abigen --abi build/CCIPCapabilityConfiguration.abi --bin build/CCIPCapabilityConfiguration.bin --pkg=main --out=CCIPCapabilityConfiguration.go
 *
 */

//go:embed build/CCIPCapabilityConfiguration.abi
var contractABI string

//go:embed build/CCIPCapabilityConfiguration.bin
var contractBytecode string

func TestHomeChain(t *testing.T) {
	// Generate a new key pair for the simulated account
	privateKey, err := crypto.GenerateKey()
	assert.NoError(t, err)
	// Set up the genesis account with balance
	blnc, ok := big.NewInt(0).SetString("999999999999999999999999999999999999", 10)
	assert.True(t, ok)
	alloc := map[common.Address]core.GenesisAccount{crypto.PubkeyToAddress(privateKey.PublicKey): {Balance: blnc}}
	simulatedBackend := backends.NewSimulatedBackend(alloc, 0)
	// Create a transactor
	const chainID = 1337
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
	assert.NoError(t, err)
	auth.GasLimit = uint64(0)

	// Deploy the contract
	parsed, err := abi.JSON(strings.NewReader(contractABI))
	assert.NoError(t, err)
	address, tx, _, err := bind.DeployContract(auth, parsed, common.FromHex(contractBytecode), simulatedBackend)
	assert.NoError(t, err)
	simulatedBackend.Commit()
	h, err := bind.WaitMined(context.Background(), simulatedBackend, tx)
	assert.NoError(t, err)
	t.Logf("contract deployed: addr=%s tx=%s block=%s", address.Hex(), tx.Hash(), h.BlockNumber.String())

	// Setup contract client
	contract, err := NewMain(address, simulatedBackend)
	assert.NoError(t, err)

	// Set up an event watcher
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		Addresses: []common.Address{address},
	}
	logs := make(chan types.Log)
	sub, err := simulatedBackend.SubscribeFilterLogs(context.Background(), query, logs)
	assert.NoError(t, err)

	// Initialize chainReader
	cfg := evmtypes.ChainReaderConfig{
		Contracts: map[string]evmtypes.ChainContractReader{
			"myContract": {
				ContractABI: contractABI,
				Configs: map[string]*evmtypes.ChainReaderDefinition{
					"EventCounter": {
						ChainSpecificName:       "SimpleEvent",
						ReadType:                evmtypes.Event,
						ConfidenceConfirmations: map[string]int{"0.0": 0, "1.0": 0},
					},
					"GetCounter": {
						ChainSpecificName: "getEventCount",
					},
				},
			},
		},
	}

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

	cr, err := evm.NewChainReaderService(context.Background(), lggr, lp, cl, cfg)
	assert.NoError(t, err)
	err = cr.Bind(context.Background(), []types2.BoundContract{
		{
			Address: address.String(),
			Name:    "myContract",
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

	var wg sync.WaitGroup
	wg.Add(2)

	// Listen events using go-ethereum lib
	go func() {
		numLogs := 0
		defer wg.Done()
		for {
			// Wait for the events
			select {
			case err := <-sub.Err():
				assert.NoError(t, err, "got an unexpected error")
			case vLog := <-logs:
				t.Logf("%d: got log: %s %s %x", numLogs, address.Hex(), vLog.Address.Hex(), vLog.Data)
				numLogs++

				if numLogs == 10 {
					return
				}
			}
		}
	}()

	// Start emitting events
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			tx, err := contract.EmitEvent(auth)
			assert.NoError(t, err)
			simulatedBackend.Commit()
			rcp, err := bind.WaitMined(context.Background(), simulatedBackend, tx)
			assert.NoError(t, err)
			assert.Equal(t, uint64(1), rcp.Status)
			t.Logf(">>> event emitted: %d @ block %s (tx: %s)", i, rcp.BlockNumber.String(), tx.Hash().Hex())
		}
	}()
	wg.Wait()

	simulatedBackend.Commit()
	simulatedBackend.Commit()
	simulatedBackend.Commit()
	simulatedBackend.Commit()
	simulatedBackend.Commit()
	simulatedBackend.Commit()
	simulatedBackend.Commit()
	simulatedBackend.Commit()
	time.Sleep(5 * time.Second)

	// Now read the contract using chain reader
	var cnt big.Int
	err = cr.GetLatestValue(
		context.Background(),
		"myContract",
		"GetCounter",
		map[string]interface{}{},
		&cnt,
	)
	assert.NoError(t, err)
	t.Logf("got cnt: %s", cnt.String())

	// Also read the events using chain reader

	var myDataType *big.Int
	seq, err := cr.QueryKey(
		context.Background(),
		"myContract",
		query2.KeyFilter{
			Key:         "EventCounter",
			Expressions: []query2.Expression{},
		},
		query2.LimitAndSort{},
		myDataType,
	)
	assert.NoError(t, err)
	assert.Len(t, seq, 10)
	for _, v := range seq {
		t.Logf("got event: %#v %s", v, cnt.String())
	}
}
