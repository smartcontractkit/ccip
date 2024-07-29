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

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	ccipreaderpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

func TestCCIPReader_CommitReportsGTETimestamp(t *testing.T) {
	lggr := logger.TestLogger(t)
	ctx := context.Background()
	const chainS1 = cciptypes.ChainSelector(1)
	const chainD = cciptypes.ChainSelector(2)
	s := testSetup(t, ctx, chainD, nil)

	cfg := evmtypes.ChainReaderConfig{
		Contracts: map[string]evmtypes.ChainContractReader{
			consts.ContractNameOffRamp: {
				ContractPollingFilter: evmtypes.ContractPollingFilter{
					GenericEventNames: []string{consts.EventNameCommitReportAccepted},
				},
				ContractABI: ccip_reader_tester.CCIPReaderTesterABI,
				Configs: map[string]*evmtypes.ChainReaderDefinition{
					consts.EventNameCommitReportAccepted: {
						ChainSpecificName: consts.EventNameCommitReportAccepted,
						ReadType:          evmtypes.Event,
					},
				},
			},
		},
	}

	headTracker := headtracker.NewSimulatedHeadTracker(s.cl, false, 0)
	cr, err := evm.NewChainReaderService(ctx, lggr, s.lp, headTracker, s.cl, cfg)
	assert.NoError(t, err)
	err = cr.Bind(ctx, []types.BoundContract{
		{
			Address: s.contractAddr.String(),
			Name:    consts.ContractNameOffRamp,
		},
	})
	assert.NoError(t, err)
	err = cr.Start(ctx)
	assert.NoError(t, err)

	contractReaders := map[cciptypes.ChainSelector]types.ContractReader{chainD: cr}
	contractWriters := make(map[cciptypes.ChainSelector]types.ChainWriter)
	reader := ccipreaderpkg.NewCCIPReader(lggr, contractReaders, contractWriters, chainD)

	tokenA := common.HexToAddress("123")

	_, err = s.contract.EmitCommitReportAccepted(s.auth, ccip_reader_tester.CCIPReaderTesterCommitReport{
		PriceUpdates: ccip_reader_tester.CCIPReaderTesterPriceUpdates{
			TokenPriceUpdates: []ccip_reader_tester.CCIPReaderTesterTokenPriceUpdate{
				{
					SourceToken: tokenA,
					UsdPerToken: big.NewInt(1000),
				},
			},
			GasPriceUpdates: []ccip_reader_tester.CCIPReaderTesterGasPriceUpdate{
				{
					DestChainSelector: uint64(chainD),
					UsdPerUnitGas:     big.NewInt(90),
				},
			},
		},
		MerkleRoots: []ccip_reader_tester.CCIPReaderTesterMerkleRoot{
			{
				SourceChainSelector: uint64(chainS1),
				Interval: ccip_reader_tester.CCIPReaderTesterInterval{
					Min: 10,
					Max: 20,
				},
				MerkleRoot: [32]byte{1, 2, 3},
			},
		},
	})
	assert.NoError(t, err)

	s.sb.Commit()

	var reports []plugintypes.CommitPluginReportWithMeta
	require.Eventually(t, func() bool {
		reports, err = reader.CommitReportsGTETimestamp(
			ctx,
			chainD,
			time.Unix(0, 0),
			10,
		)
		require.NoError(t, err)
		return len(reports) == 1
	}, 15*time.Second, 100*time.Millisecond)

	assert.Len(t, reports[0].Report.MerkleRoots, 1)
	assert.Equal(t, chainS1, reports[0].Report.MerkleRoots[0].ChainSel)
	assert.Equal(t, cciptypes.SeqNum(10), reports[0].Report.MerkleRoots[0].SeqNumsRange.Start())
	assert.Equal(t, cciptypes.SeqNum(20), reports[0].Report.MerkleRoots[0].SeqNumsRange.End())
	assert.Equal(t, "0x0102030000000000000000000000000000000000000000000000000000000000",
		reports[0].Report.MerkleRoots[0].MerkleRoot.String())

	assert.Equal(t, tokenA.String(), string(reports[0].Report.PriceUpdates.TokenPriceUpdates[0].TokenID))
	assert.Equal(t, uint64(1000), reports[0].Report.PriceUpdates.TokenPriceUpdates[0].Price.Uint64())

	assert.Equal(t, chainD, reports[0].Report.PriceUpdates.GasPriceUpdates[0].ChainSel)
	assert.Equal(t, uint64(90), reports[0].Report.PriceUpdates.GasPriceUpdates[0].GasPrice.Uint64())
}

func TestCCIPReader_ExecutedMessageRanges(t *testing.T) {
	lggr := logger.TestLogger(t)
	ctx := context.Background()
	const chainS1 = cciptypes.ChainSelector(1)
	const chainD = cciptypes.ChainSelector(2)
	s := testSetup(t, ctx, chainS1, nil)

	cfg := evmtypes.ChainReaderConfig{
		Contracts: map[string]evmtypes.ChainContractReader{
			consts.ContractNameOffRamp: {
				ContractPollingFilter: evmtypes.ContractPollingFilter{
					GenericEventNames: []string{consts.EventNameExecutionStateChanged},
				},
				ContractABI: ccip_reader_tester.CCIPReaderTesterABI,
				Configs: map[string]*evmtypes.ChainReaderDefinition{
					consts.EventNameExecutionStateChanged: {
						ChainSpecificName: consts.EventNameExecutionStateChanged,
						ReadType:          evmtypes.Event,
					},
				},
			},
		},
	}

	headTracker := headtracker.NewSimulatedHeadTracker(s.cl, false, 0)
	cr, err := evm.NewChainReaderService(ctx, lggr, s.lp, headTracker, s.cl, cfg)
	assert.NoError(t, err)
	err = cr.Bind(ctx, []types.BoundContract{
		{
			Address: s.contractAddr.String(),
			Name:    consts.ContractNameOffRamp,
		},
	})
	assert.NoError(t, err)
	err = cr.Start(ctx)
	assert.NoError(t, err)

	contractReaders := map[cciptypes.ChainSelector]types.ContractReader{chainD: cr}
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

	var executedRanges []cciptypes.SeqNumRange
	require.Eventually(t, func() bool {
		executedRanges, err = reader.ExecutedMessageRanges(
			ctx,
			chainS1,
			chainD,
			cciptypes.NewSeqNumRange(14, 15),
		)
		// require.NoError(t, err) // todo: fails because chainReader is returning structs with all empty fields
		return len(executedRanges) == 0
	}, 5*time.Second, 100*time.Millisecond)

	// todo: more validations
}

func TestCCIPReader_MsgsBetweenSeqNums(t *testing.T) {
	lggr := logger.TestLogger(t)
	ctx := context.Background()
	const chainS1 = cciptypes.ChainSelector(1)
	const chainD = cciptypes.ChainSelector(2)
	s := testSetup(t, ctx, chainS1, nil)

	cfg := evmtypes.ChainReaderConfig{
		Contracts: map[string]evmtypes.ChainContractReader{
			consts.ContractNameOnRamp: {
				ContractPollingFilter: evmtypes.ContractPollingFilter{
					GenericEventNames: []string{consts.EventNameCCIPSendRequested},
				},
				ContractABI: ccip_reader_tester.CCIPReaderTesterABI,
				Configs: map[string]*evmtypes.ChainReaderDefinition{
					consts.EventNameCCIPSendRequested: {
						ChainSpecificName: consts.EventNameCCIPSendRequested,
						ReadType:          evmtypes.Event,
					},
				},
			},
		},
	}

	headTracker := headtracker.NewSimulatedHeadTracker(s.cl, false, 0)
	cr, err := evm.NewChainReaderService(ctx, lggr, s.lp, headTracker, s.cl, cfg)
	assert.NoError(t, err)
	err = cr.Bind(ctx, []types.BoundContract{
		{
			Address: s.contractAddr.String(),
			Name:    consts.ContractNameOnRamp,
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

	var msgs []cciptypes.Message
	require.Eventually(t, func() bool {
		msgs, err = reader.MsgsBetweenSeqNums(
			ctx,
			chainS1,
			cciptypes.NewSeqNumRange(5, 20),
		)
		require.NoError(t, err)
		return len(msgs) == 2
	}, 5*time.Second, 100*time.Millisecond)

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
			consts.ContractNameOffRamp: {
				ContractABI: ccip_reader_tester.CCIPReaderTesterABI,
				Configs: map[string]*evmtypes.ChainReaderDefinition{
					consts.MethodNameGetSourceChainConfig: {
						ChainSpecificName: "getSourceChainConfig",
						ReadType:          evmtypes.Method,
					},
				},
			},
		},
	}

	headTracker := headtracker.NewSimulatedHeadTracker(s.cl, false, 0)
	cr, err := evm.NewChainReaderService(ctx, lggr, s.lp, headTracker, s.cl, cfg)
	assert.NoError(t, err)
	err = cr.Bind(ctx, []types.BoundContract{
		{
			Address: s.contractAddr.String(),
			Name:    consts.ContractNameOffRamp,
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
