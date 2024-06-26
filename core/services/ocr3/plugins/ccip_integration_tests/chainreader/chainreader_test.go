package chainreader

import (
	"context"
	_ "embed"
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-common/pkg/codec"
	query2 "github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	helpers "github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/ccip_integration_tests"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
	"github.com/stretchr/testify/assert"
)

func TestChainReader(t *testing.T) {
	ctx := testutils.Context(t)
	const (
		ContractNameAlias = "myCoolContract"

		FnAliasGetCount = "myCoolFunction"
		FnGetCount      = "getEventCount"

		FnAliasGetNumbers = "GetNumbers"
		FnGetNumbers      = "getNumbers"

		FnAliasGetPerson = "GetPerson"
		FnGetPerson      = "getPerson"

		EventNameAlias = "myCoolEvent"
		EventName      = "SimpleEvent"
	)

	// Initialize chainReader
	cfg := evmtypes.ChainReaderConfig{
		Contracts: map[string]evmtypes.ChainContractReader{
			ContractNameAlias: {
				ContractABI: ChainreaderMetaData.ABI,
				Configs: map[string]*evmtypes.ChainReaderDefinition{
					EventNameAlias: {
						ChainSpecificName:       EventName,
						ReadType:                evmtypes.Event,
						ConfidenceConfirmations: map[string]int{"0.0": 0, "1.0": 0},
					},
					FnAliasGetCount: {
						ChainSpecificName: FnGetCount,
					},
					FnAliasGetNumbers: {
						ChainSpecificName:   FnGetNumbers,
						OutputModifications: codec.ModifiersConfig{},
					},
					FnAliasGetPerson: {
						ChainSpecificName: FnGetPerson,
						OutputModifications: codec.ModifiersConfig{
							&codec.RenameModifierConfig{
								Fields: map[string]string{"Name": "NameField"}, // solidity name -> go struct name
							},
						},
					},
				},
			},
		},
	}

	d := helpers.SetupChainReader[Chainreader](t, ctx, DeployChainreader, NewChainreader, cfg, ContractNameAlias)
	cr := *d.ChainReader

	emitEvents(t, d, ctx) // Calls the contract to emit events

	// (hack) Sometimes LP logs are missing, commit several times and wait few seconds to make it work.
	for i := 0; i < 100; i++ {
		d.SimulatedBE.Commit()
	}
	time.Sleep(5 * time.Second)

	t.Run("simple contract read", func(t *testing.T) {
		var cnt big.Int
		err := cr.GetLatestValue(ctx, ContractNameAlias, FnAliasGetCount, map[string]interface{}{}, &cnt)
		assert.NoError(t, err)
		assert.Equal(t, int64(10), cnt.Int64())
	})

	t.Run("read array", func(t *testing.T) {
		var nums []big.Int
		err := cr.GetLatestValue(ctx, ContractNameAlias, FnAliasGetNumbers, map[string]interface{}{}, &nums)
		assert.NoError(t, err)
		assert.Len(t, nums, 10)
		for i := 1; i <= 10; i++ {
			assert.Equal(t, int64(i), nums[i-1].Int64())
		}
	})

	t.Run("read struct", func(t *testing.T) {
		person := struct {
			NameField string
			Age       *big.Int // WARN: specifying a wrong data type e.g. int instead of *big.Int fails silently with a default value of 0
		}{}
		err := cr.GetLatestValue(ctx, ContractNameAlias, FnAliasGetPerson, map[string]interface{}{}, &person)
		assert.NoError(t, err)
		assert.Equal(t, "Dim", person.NameField)
		assert.Equal(t, int64(18), person.Age.Int64())
	})

	t.Run("read events", func(t *testing.T) {
		var myDataType *big.Int
		seq, err := cr.QueryKey(
			ctx,
			ContractNameAlias,
			query2.KeyFilter{
				Key:         EventNameAlias,
				Expressions: []query2.Expression{},
			},
			query2.LimitAndSort{},
			myDataType,
		)
		assert.NoError(t, err)
		assert.Equal(t, 10, len(seq), "expected 10 events from chain reader")
		for _, v := range seq {
			// TODO: for some reason log poller does not populate event data
			t.Logf("(chain reader) got event: (data=%v) (hash=%x)", v.Data, v.Hash)
		}
	})
}

func emitEvents(t *testing.T, d *helpers.TestSetupData[Chainreader], ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(2)

	// Start emitting events
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			contract := d.Contract
			_, err := contract.EmitEvent(d.Auth)
			assert.NoError(t, err)
			d.SimulatedBE.Commit()
		}
	}()

	// Listen events using go-ethereum lib
	go func() {
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(0),
			Addresses: []common.Address{d.ContractAddr},
		}
		logs := make(chan types.Log)
		sub, err := d.SimulatedBE.SubscribeFilterLogs(ctx, query, logs)
		assert.NoError(t, err)

		numLogs := 0
		defer wg.Done()
		for {
			// Wait for the events
			select {
			case err := <-sub.Err():
				assert.NoError(t, err, "got an unexpected error")
			case vLog := <-logs:
				assert.Equal(t, d.ContractAddr, vLog.Address, "got an unexpected address")
				t.Logf("(geth) got new log (cnt=%d) (data=%x) (topics=%s)", numLogs, vLog.Data, vLog.Topics)
				numLogs++
				if numLogs == 10 {
					return
				}
			}
		}
	}()

	wg.Wait() // wait for all the events to be consumed
}
