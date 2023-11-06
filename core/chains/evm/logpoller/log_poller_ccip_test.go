package logpoller_test

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/test-go/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/internal/cltest/heavyweight"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

const (
	numberOfReports           = 1000
	numberOfMessagesPerReport = 100
	numberOfMessages          = numberOfReports * numberOfMessagesPerReport
)

func populateDbWithCommitReports(b *testing.B, o *logpoller.DbORM, chainID *big.Int, commitStoreAddress common.Address, commitReportAccepted common.Hash) {
	// Max we can insert per batch
	var logs []logpoller.Log
	for i := 0; i < numberOfReports; i++ {
		data := make([]byte, 64)
		// MinSeqNr
		data = append(data, logpoller.EvmWord(uint64(numberOfMessagesPerReport*i+1)).Bytes()...)
		// MaxSeqNr
		data = append(data, logpoller.EvmWord(uint64(numberOfMessagesPerReport*(i+1))).Bytes()...)

		logs = append(logs, logpoller.Log{
			EvmChainId:     utils.NewBig(chainID),
			LogIndex:       int64(i + 1),
			BlockHash:      utils.RandomBytes32(),
			BlockNumber:    int64(i + 1),
			BlockTimestamp: time.Now(),
			EventSig:       commitReportAccepted,
			Topics:         [][]byte{},
			Address:        commitStoreAddress,
			TxHash:         utils.RandomAddress().Hash(),
			Data:           data,
			CreatedAt:      time.Now(),
		})
	}
	require.NoError(b, o.InsertBlock(utils.RandomAddress().Hash(), int64(10_000), time.Now()))
	require.NoError(b, o.InsertLogs(logs))
}

func populateDbWithExecutionStateChanges(b *testing.B, o *logpoller.DbORM, chainID *big.Int, offrampAddress common.Address, offrampExecuted common.Hash) {
	var logs []logpoller.Log
	for i := 1; i <= numberOfMessages; i++ {
		var topics [][]byte
		for j := 0; j < 5; j++ {
			topics = append(topics, logpoller.EvmWord(uint64(i)).Bytes())
		}

		logs = append(logs, logpoller.Log{
			EvmChainId:     utils.NewBig(chainID),
			LogIndex:       int64(i),
			BlockHash:      utils.RandomBytes32(),
			BlockNumber:    int64(i),
			BlockTimestamp: time.Now(),
			EventSig:       offrampExecuted,
			Topics:         topics,
			Address:        offrampAddress,
			TxHash:         utils.RandomAddress().Hash(),
			Data:           []byte{},
			CreatedAt:      time.Now(),
		})

	}
	require.NoError(b, o.InsertBlock(utils.RandomAddress().Hash(), int64(100_000), time.Now()))
	require.NoError(b, o.InsertLogs(logs))
}

func populateDbWithSomeExecuted(b *testing.B, o *logpoller.DbORM, chainID *big.Int, offrampAddress common.Address, offrampExecuted common.Hash) {
	var logs []logpoller.Log
	for i := 1; i <= numberOfMessages; i += 2 {
		var topics [][]byte
		for j := 0; j < 5; j++ {
			topics = append(topics, logpoller.EvmWord(uint64(i)).Bytes())
		}

		logs = append(logs, logpoller.Log{
			EvmChainId:     utils.NewBig(chainID),
			LogIndex:       int64(i),
			BlockHash:      utils.RandomBytes32(),
			BlockNumber:    int64(i),
			BlockTimestamp: time.Now(),
			EventSig:       offrampExecuted,
			Topics:         topics,
			Address:        offrampAddress,
			TxHash:         utils.RandomAddress().Hash(),
			Data:           []byte{},
			CreatedAt:      time.Now(),
		})

	}
	require.NoError(b, o.InsertBlock(utils.RandomAddress().Hash(), int64(100_000), time.Now()))
	require.NoError(b, o.InsertLogs(logs))
}

func populateDbWithMessages(b *testing.B, o *logpoller.DbORM, chainID *big.Int, onrampAddress common.Address, onrampEvent common.Hash) {
	var logs []logpoller.Log
	for i := 1; i <= numberOfMessages; i++ {
		data := make([]byte, 128)
		// SeqNr
		data = append(data, logpoller.EvmWord(uint64(i)).Bytes()...)

		logs = append(logs, logpoller.Log{
			EvmChainId:     utils.NewBig(chainID),
			LogIndex:       int64(i),
			BlockHash:      utils.RandomBytes32(),
			BlockNumber:    int64(i),
			BlockTimestamp: time.Now(),
			EventSig:       onrampEvent,
			Topics:         [][]byte{},
			Address:        onrampAddress,
			TxHash:         utils.RandomAddress().Hash(),
			Data:           data,
			CreatedAt:      time.Now(),
		})

	}
	require.NoError(b, o.InsertBlock(utils.RandomAddress().Hash(), int64(100_000), time.Now()))
	require.NoError(b, o.InsertLogs(logs))
}

func Benchmark_CreatedAfter(b *testing.B) {
	chainId := big.NewInt(137)
	_, db := heavyweight.FullTestDBV2(b, "logs_scale_created_after", nil)
	o := logpoller.NewORM(chainId, db, logger.TestLogger(b), pgtest.NewQConfig(false))

	commitStoreAddress := common.HexToAddress("0x2ab9a2Dc53736b361b72d900CdF9F78F9406fbbb")
	commitReportAccepted := common.HexToHash("0xe81b49e583122eb290c46fc255c962b9a2dec468816c00fb7a2e6ebc42dc92d4")

	offrampAddress := common.HexToAddress("0x6E225058950f237371261C985Db6bDe26df2200E")
	offrampExecuted := common.HexToHash("0xd4f851956a5d67c3997d1c9205045fef79bae2947fdee7e9e2641abc7391ef65")

	populateDbWithCommitReports(b, o, chainId, commitStoreAddress, commitReportAccepted)
	populateDbWithExecutionStateChanges(b, o, chainId, offrampAddress, offrampExecuted)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		start := time.Now()
		logs, err := o.SelectLogsCreatedAfter(
			commitStoreAddress,
			commitReportAccepted,
			time.Now().Add(-1*time.Hour),
			0,
		)
		fmt.Printf("%d millis\n", time.Since(start).Milliseconds())
		require.NoError(b, err)
		require.Len(b, logs, numberOfReports)
	}
}

func Benchmark_GetExecutionStatesAndMessages(b *testing.B) {
	chainId := big.NewInt(137)
	_, db := heavyweight.FullTestDBV2(b, "msgs", nil)
	o := logpoller.NewORM(chainId, db, logger.TestLogger(b), pgtest.NewQConfig(false))

	commitStoreAddress := utils.RandomAddress()
	commitReportAccepted := common.HexToHash("0xe81b49e583122eb290c46fc255c962b9a2dec468816c00fb7a2e6ebc42dc92d4")

	offrampAddress := utils.RandomAddress()
	offrampExecuted := common.HexToHash("0xd4f851956a5d67c3997d1c9205045fef79bae2947fdee7e9e2641abc7391ef65")

	onrampAddress := utils.RandomAddress()
	onrampEvent := common.HexToHash("0xe2947fdee7e9e2641a")

	populateDbWithCommitReports(b, o, chainId, commitStoreAddress, commitReportAccepted)
	populateDbWithExecutionStateChanges(b, o, chainId, offrampAddress, offrampExecuted)
	populateDbWithMessages(b, o, chainId, onrampAddress, onrampEvent)

	for i := 0; i < b.N; i++ {
		start := time.Now()
		// Get commit reports
		commitReports, err := o.SelectLogsCreatedAfter(
			commitStoreAddress,
			commitReportAccepted,
			time.Now().Add(-1*time.Hour),
			0,
		)
		require.NoError(b, err)
		assert.Len(b, commitReports, numberOfReports)
		fmt.Printf("Commit Reports: %d millis\n", time.Since(start).Milliseconds())

		start = time.Now()
		messages, err := o.SelectLogsDataWordRange(
			onrampAddress,
			onrampEvent,
			0,
			logpoller.EvmWord(0),
			logpoller.EvmWord(numberOfMessages),
			0,
		)
		require.NoError(b, err)
		assert.Len(b, messages, numberOfMessages)
		fmt.Printf("OnRamp messages: %d millis\n", time.Since(start).Milliseconds())

		start = time.Now()
		executionStateChanges, err := o.SelectIndexedLogsTopicRange(
			offrampAddress,
			offrampExecuted,
			2,
			logpoller.EvmWord(0),
			logpoller.EvmWord(numberOfMessages),
			0,
		)
		require.NoError(b, err)
		assert.Len(b, executionStateChanges, numberOfMessages)
		fmt.Printf("Offramp exec state changes: %d millis\n", time.Since(start).Milliseconds())
	}
}

func Benchmark_SingleQueryAllExecuted(b *testing.B) {
	chainId := big.NewInt(137)
	_, db := heavyweight.FullTestDBV2(b, "logs_match_none", nil)
	o := logpoller.NewORM(chainId, db, logger.TestLogger(b), pgtest.NewQConfig(false))

	commitStoreAddress := utils.RandomAddress()
	commitReportAccepted := common.HexToHash("0xe81b49e583122eb290c46fc255c962b9a2dec468816c00fb7a2e6ebc42dc92d4")

	offrampAddress := utils.RandomAddress()
	offrampExecuted := common.HexToHash("0xd4f851956a5d67c3997d1c9205045fef79bae2947fdee7e9e2641abc7391ef65")

	populateDbWithCommitReports(b, o, chainId, commitStoreAddress, commitReportAccepted)
	populateDbWithExecutionStateChanges(b, o, chainId, offrampAddress, offrampExecuted)

	b.Log("Db load, running query")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		start := time.Now()
		logs, err := o.FetchNotExecutedReports(
			commitStoreAddress,
			commitReportAccepted,
			offrampAddress,
			offrampExecuted,
			time.Now().Add(-1*time.Hour),
		)
		fmt.Printf("%d millis\n", time.Since(start).Milliseconds())
		require.NoError(b, err)
		require.Len(b, logs, 0)
	}
}

func Benchmark_SingleQueryNoneExecuted(b *testing.B) {
	chainId := big.NewInt(137)
	_, db := heavyweight.FullTestDBV2(b, "log_match_all", nil)
	o := logpoller.NewORM(chainId, db, logger.TestLogger(b), pgtest.NewQConfig(false))

	commitStoreAddress := common.HexToAddress("0x2ab9a2Dc53736b361b72d900CdF9F78F9406fbbb")
	commitReportAccepted := common.HexToHash("0xe81b49e583122eb290c46fc255c962b9a2dec468816c00fb7a2e6ebc42dc92d4")

	offrampAddress := common.HexToAddress("0x6E225058950f237371261C985Db6bDe26df2200E")
	offrampExecuted := common.HexToHash("0xd4f851956a5d67c3997d1c9205045fef79bae2947fdee7e9e2641abc7391ef65")

	populateDbWithCommitReports(b, o, chainId, commitStoreAddress, commitReportAccepted)

	b.Log("Db load, running query")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		start := time.Now()
		logs, err := o.FetchNotExecutedReports(
			commitStoreAddress,
			commitReportAccepted,
			offrampAddress,
			offrampExecuted,
			time.Now().Add(-1*time.Hour),
		)
		fmt.Printf("%d millis\n", time.Since(start).Milliseconds())
		require.NoError(b, err)
		require.Len(b, logs, numberOfReports)
	}
}

func Benchmark_SingleQuerySomeExecuted(b *testing.B) {
	chainId := big.NewInt(137)
	_, db := heavyweight.FullTestDBV2(b, "log_match_all", nil)
	o := logpoller.NewORM(chainId, db, logger.TestLogger(b), pgtest.NewQConfig(false))

	commitStoreAddress := common.HexToAddress("0x2ab9a2Dc53736b361b72d900CdF9F78F9406fbbb")
	commitReportAccepted := common.HexToHash("0xe81b49e583122eb290c46fc255c962b9a2dec468816c00fb7a2e6ebc42dc92d4")

	offrampAddress := common.HexToAddress("0x6E225058950f237371261C985Db6bDe26df2200E")
	offrampExecuted := common.HexToHash("0xd4f851956a5d67c3997d1c9205045fef79bae2947fdee7e9e2641abc7391ef65")

	populateDbWithCommitReports(b, o, chainId, commitStoreAddress, commitReportAccepted)
	populateDbWithSomeExecuted(b, o, chainId, offrampAddress, offrampExecuted)

	b.Log("Db load, running query")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		start := time.Now()
		logs, err := o.FetchNotExecutedReports(
			commitStoreAddress,
			commitReportAccepted,
			offrampAddress,
			offrampExecuted,
			time.Now().Add(-1*time.Hour),
		)
		fmt.Printf("%d millis\n", time.Since(start).Milliseconds())
		require.NoError(b, err)
		require.Len(b, logs, numberOfReports)
	}
}

// V1
// 168 millis
// Benchmark_AllMessagesExecuted-12    	       1	10373778417 ns/op

// V2
// 93 millis
//
// Benchmark_AllMessagesExecuted
// Benchmark_AllMessagesExecuted-12    	      14	  83441804 ns/op
// Benchmark_AllMessagesExecuted-12    	      14	  80466869 ns/op
// Benchmark_AllMessagesExecuted-12    	      12	  83875597 ns/op
// Benchmark_AllMessagesExecuted-12    	      13	  79230596 ns/op
// Benchmark_AllMessagesExecuted-12    	      13	  83327372 ns/op
func Benchmark_AllMessagesExecuted(b *testing.B) {
	chainId := big.NewInt(137)
	_, db := heavyweight.FullTestDBV2(b, "msgs", nil)
	o := logpoller.NewORM(chainId, db, logger.TestLogger(b), pgtest.NewQConfig(false))

	offrampAddress := utils.RandomAddress()
	offrampExecuted := common.HexToHash("0xd4f851956a5d67c3997d1c9205045fef79bae2947fdee7e9e2641abc7391ef65")

	onrampAddress := utils.RandomAddress()
	onrampEvent := common.HexToHash("0xaffc45517195d6499808c643bd4a7b0ffeedf95bea5852840d7bfcf63f59e821")

	populateDbWithMessages(b, o, chainId, onrampAddress, onrampEvent)
	populateDbWithExecutionStateChanges(b, o, chainId, offrampAddress, offrampExecuted)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		//start := time.Now()
		logs, err := o.FetchNotExecutedMessages(
			onrampAddress,
			onrampEvent,
			offrampAddress,
			offrampExecuted,
			logpoller.EvmWord(0),
			logpoller.EvmWord(numberOfMessages),
		)
		//fmt.Printf("%d millis\n", time.Since(start).Milliseconds())
		require.NoError(b, err)
		require.Len(b, logs, 0)
	}
}

// V1
// 189 millis
// Benchmark_NoneMessagesExecuted-12    	       1	3135554125 ns/op
// V2
// 230 millis
// Benchmark_NoneMessagesExecuted-12    	       1	3809230583 ns/op

func Benchmark_NoneMessagesExecuted(b *testing.B) {
	chainId := big.NewInt(137)
	_, db := heavyweight.FullTestDBV2(b, "msgs", nil)
	o := logpoller.NewORM(chainId, db, logger.TestLogger(b), pgtest.NewQConfig(false))

	offrampAddress := utils.RandomAddress()
	offrampExecuted := common.HexToHash("0xd4f851956a5d67c3997d1c9205045fef79bae2947fdee7e9e2641abc7391ef65")

	onrampAddress := utils.RandomAddress()
	onrampEvent := common.HexToHash("0xaffc45517195d6499808c643bd4a7b0ffeedf95bea5852840d7bfcf63f59e821")

	populateDbWithMessages(b, o, chainId, onrampAddress, onrampEvent)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		start := time.Now()
		logs, err := o.FetchNotExecutedMessages(
			onrampAddress,
			onrampEvent,
			offrampAddress,
			offrampExecuted,
			logpoller.EvmWord(0),
			logpoller.EvmWord(numberOfMessages),
		)
		fmt.Printf("%d millis\n", time.Since(start).Milliseconds())
		require.NoError(b, err)
		require.Len(b, logs, numberOfMessages)
	}
}
