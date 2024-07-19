package ccipreader

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/pkg/crconsts"
	ccipreaderpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/headtracker"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/ccip_reader_tester"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/relay/evm"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/services/relay/evm/types"
)

func TestCCIPReader_CommitReportsGTETimestamp(t *testing.T) {}

func TestCCIPReader_ExecutedMessageRanges(t *testing.T) {
	lggr := logger.TestLogger(t)
	ctx := context.Background()
	const chainS1 = cciptypes.ChainSelector(1)
	const chainD = cciptypes.ChainSelector(2)
	s := testSetup(t, ctx, chainS1, nil)

	cfg := evmtypes.ChainReaderConfig{
		Contracts: map[string]evmtypes.ChainContractReader{
			crconsts.ContractNameOffRamp: {
				ContractPollingFilter: evmtypes.ContractPollingFilter{
					GenericEventNames: []string{crconsts.EventNameExecutionStateChanged},
				},
				ContractABI: ccip_reader_tester.CCIPReaderTesterABI,
				Configs: map[string]*evmtypes.ChainReaderDefinition{
					crconsts.EventNameExecutionStateChanged: {
						ChainSpecificName: crconsts.EventNameExecutionStateChanged,
						ReadType:          evmtypes.Event,
					},
				},
			},
		},
	}

	cr, err := evm.NewChainReaderService(ctx, lggr, s.lp, s.cl, cfg)
	assert.NoError(t, err)
	err = cr.Bind(ctx, []types.BoundContract{
		{
			Address: s.contractAddr.String(),
			Name:    crconsts.ContractNameOffRamp,
			Pending: false,
		},
	})
	assert.NoError(t, err)
	err = cr.Start(ctx)
	assert.NoError(t, err)

	contractReaders := map[cciptypes.ChainSelector]types.ContractReader{chainS1: cr}
	contractWriters := make(map[cciptypes.ChainSelector]types.ChainWriter)
	reader := ccipreaderpkg.NewCCIPReader(lggr, contractReaders, contractWriters, chainD)

	_, err = s.contract.EmitExecutionStateChanged(
		s.auth,
		uint64(chainS1),
		14,
		cciptypes.Bytes32{1, 0, 0, 1},
		1,
		[]byte{1, 2, 3, 4},
	)
	assert.NoError(t, err)

	_, err = s.contract.EmitExecutionStateChanged(
		s.auth,
		uint64(chainS1),
		15,
		cciptypes.Bytes32{1, 0, 0, 2},
		1,
		[]byte{1, 2, 3, 4, 5},
	)
	assert.NoError(t, err)

	s.sb.Commit()
	time.Sleep(5 * time.Second)

	executedRanges, err := reader.ExecutedMessageRanges(
		ctx,
		chainS1,
		chainD,
		cciptypes.NewSeqNumRange(14, 15),
	)
	require.NoError(t, err)
	require.Len(t, executedRanges, 2)
}

func TestCCIPReader_MsgsBetweenSeqNums(t *testing.T) {
	lggr := logger.TestLogger(t)
	ctx := context.Background()
	const chainS1 = cciptypes.ChainSelector(1)
	const chainD = cciptypes.ChainSelector(2)
	s := testSetup(t, ctx, chainS1, nil)

	cfg := evmtypes.ChainReaderConfig{
		Contracts: map[string]evmtypes.ChainContractReader{
			crconsts.ContractNameOnRamp: {
				ContractPollingFilter: evmtypes.ContractPollingFilter{
					GenericEventNames: []string{crconsts.EventNameCCIPSendRequested},
				},
				ContractABI: ccip_reader_tester.CCIPReaderTesterABI,
				Configs: map[string]*evmtypes.ChainReaderDefinition{
					crconsts.EventNameCCIPSendRequested: {
						ChainSpecificName: crconsts.EventNameCCIPSendRequested,
						ReadType:          evmtypes.Event,
					},
				},
			},
		},
	}

	cr, err := evm.NewChainReaderService(ctx, lggr, s.lp, s.cl, cfg)
	assert.NoError(t, err)
	err = cr.Bind(ctx, []types.BoundContract{
		{
			Address: s.contractAddr.String(),
			Name:    crconsts.ContractNameOnRamp,
			Pending: false,
		},
	})
	assert.NoError(t, err)
	err = cr.Start(ctx)
	assert.NoError(t, err)

	contractReaders := map[cciptypes.ChainSelector]types.ContractReader{chainS1: cr}
	contractWriters := make(map[cciptypes.ChainSelector]types.ChainWriter)
	reader := ccipreaderpkg.NewCCIPReader(lggr, contractReaders, contractWriters, chainD)

	_, err = s.contract.EmitCCIPSendRequested(s.auth, uint64(chainD), ccip_reader_tester.CCIPReaderTesterEVM2AnyRampMessage{
		Header: ccip_reader_tester.CCIPReaderTesterRampMessageHeader{
			MessageId:           [32]byte{1, 0, 0, 0, 0},
			SourceChainSelector: uint64(chainS1),
			DestChainSelector:   uint64(chainD),
			SequenceNumber:      10,
		},
		Sender: common.Address{},
	})
	assert.NoError(t, err)

	_, err = s.contract.EmitCCIPSendRequested(s.auth, uint64(chainD), ccip_reader_tester.CCIPReaderTesterEVM2AnyRampMessage{
		Header: ccip_reader_tester.CCIPReaderTesterRampMessageHeader{
			MessageId:           [32]byte{1, 0, 0, 0, 1},
			SourceChainSelector: uint64(chainS1),
			DestChainSelector:   uint64(chainD),
			SequenceNumber:      15,
		},
		Sender: common.Address{},
	})
	assert.NoError(t, err)

	s.sb.Commit()
	time.Sleep(5 * time.Second)

	msgs, err := reader.MsgsBetweenSeqNums(
		ctx,
		chainS1,
		cciptypes.NewSeqNumRange(5, 20),
	)
	require.NoError(t, err)
	require.Len(t, msgs, 2)
	require.Equal(t, cciptypes.SeqNum(10), msgs[0].Header.SequenceNumber)
	require.Equal(t, cciptypes.SeqNum(15), msgs[1].Header.SequenceNumber)
	for _, msg := range msgs {
		require.Equal(t, chainS1, msg.Header.SourceChainSelector)
		require.Equal(t, chainD, msg.Header.DestChainSelector)
	}
}

func TestCCIPReader_NextSeqNum(t *testing.T) {
	lggr := logger.TestLogger(t)
	ctx := context.Background()

	const chainS1 = cciptypes.ChainSelector(1)
	const chainS2 = cciptypes.ChainSelector(2)
	const chainS3 = cciptypes.ChainSelector(3)
	const chainD = cciptypes.ChainSelector(4)

	onChainSeqNums := map[cciptypes.ChainSelector]cciptypes.SeqNum{
		chainS1: 10,
		chainS2: 20,
		chainS3: 30,
	}

	s := testSetup(t, ctx, chainD, onChainSeqNums)

	cfg := evmtypes.ChainReaderConfig{
		Contracts: map[string]evmtypes.ChainContractReader{
			crconsts.ContractNameOffRamp: {
				ContractABI: ccip_reader_tester.CCIPReaderTesterABI,
				Configs: map[string]*evmtypes.ChainReaderDefinition{
					crconsts.FunctionNameGetSourceChainConfig: {
						ChainSpecificName: "getSourceChainConfig",
						ReadType:          evmtypes.Method,
					},
				},
			},
		},
	}

	cr, err := evm.NewChainReaderService(ctx, lggr, s.lp, s.cl, cfg)
	assert.NoError(t, err)
	err = cr.Bind(ctx, []types.BoundContract{
		{
			Address: s.contractAddr.String(),
			Name:    crconsts.ContractNameOffRamp,
			Pending: false,
		},
	})
	assert.NoError(t, err)
	err = cr.Start(ctx)
	assert.NoError(t, err)

	contractReaders := map[cciptypes.ChainSelector]types.ContractReader{chainD: cr}
	contractWriters := make(map[cciptypes.ChainSelector]types.ChainWriter)
	reader := ccipreaderpkg.NewCCIPReader(lggr, contractReaders, contractWriters, chainD)

	seqNums, err := reader.NextSeqNum(ctx, []cciptypes.ChainSelector{chainS1, chainS2, chainS3})
	assert.NoError(t, err)
	assert.Len(t, seqNums, 3)
	assert.Equal(t, cciptypes.SeqNum(10), seqNums[0])
	assert.Equal(t, cciptypes.SeqNum(20), seqNums[1])
	assert.Equal(t, cciptypes.SeqNum(30), seqNums[2])
}

func TestCCIPReader_GasPrices(t *testing.T) {}

func TestCCIPReader_Sync(t *testing.T) {}

func testSetup(t *testing.T, ctx context.Context, readerChain cciptypes.ChainSelector, onChainSeqNums map[cciptypes.ChainSelector]cciptypes.SeqNum) *testSetupData {
	const chainID = 1337

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

	// Deploy the contract
	address, _, _, err := ccip_reader_tester.DeployCCIPReaderTester(auth, simulatedBackend)
	assert.NoError(t, err)
	simulatedBackend.Commit()

	// Setup contract client
	contract, err := ccip_reader_tester.NewCCIPReaderTester(address, simulatedBackend)
	assert.NoError(t, err)

	lggr := logger.TestLogger(t)
	db := pgtest.NewSqlxDB(t)
	lpOpts := logpoller.Opts{
		PollPeriod:               time.Millisecond,
		FinalityDepth:            0,
		BackfillBatchSize:        10,
		RpcBatchSize:             10,
		KeepFinalizedBlocksDepth: 100000,
	}
	cl := client.NewSimulatedBackendClient(t, simulatedBackend, big.NewInt(0).SetUint64(uint64(readerChain)))
	headTracker := headtracker.NewSimulatedHeadTracker(cl, lpOpts.UseFinalityTag, lpOpts.FinalityDepth)
	lp := logpoller.NewLogPoller(logpoller.NewORM(big.NewInt(0).SetUint64(uint64(readerChain)), db, lggr),
		cl,
		lggr,
		headTracker,
		lpOpts,
	)
	assert.NoError(t, lp.Start(ctx))

	for sourceChain, seqNum := range onChainSeqNums {
		_, err := contract.SetSourceChainConfig(auth, uint64(sourceChain), ccip_reader_tester.CCIPReaderTesterSourceChainConfig{
			IsEnabled: true,
			MinSeqNr:  uint64(seqNum),
		})
		assert.NoError(t, err)
		simulatedBackend.Commit()
		scc, err := contract.GetSourceChainConfig(&bind.CallOpts{Context: ctx}, uint64(sourceChain))
		assert.NoError(t, err)
		assert.Equal(t, seqNum, cciptypes.SeqNum(scc.MinSeqNr))
	}

	return &testSetupData{
		contractAddr: address,
		contract:     contract,
		sb:           simulatedBackend,
		auth:         auth,
		lp:           lp,
		cl:           cl,
	}
}

type testSetupData struct {
	contractAddr common.Address
	contract     *ccip_reader_tester.CCIPReaderTester
	sb           *backends.SimulatedBackend
	auth         *bind.TransactOpts
	lp           logpoller.LogPoller
	cl           client.Client
}
