package ccipdata

import (
	"context"
	"crypto/sha256"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	evmclimocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/headtracker"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	lpmocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	types2 "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	ubig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
)

func TestLBTCParse(t *testing.T) {
	encodedPayload, err := hexutil.Decode("0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000e45c70a5050000000000000000000000000000000000000000000000000000000000aa36a7000000000000000000000000845f8e3c214d8d0e4d83fc094f302aa26a12a0bc0000000000000000000000000000000000000000000000000000000000014a34000000000000000000000000845f8e3c214d8d0e4d83fc094f302aa26a12a0bc00000000000000000000000062f10ce5b727edf787ea45776bd050308a61150800000000000000000000000000000000000000000000000000000000000003e6000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000")
	require.NoError(t, err)
	payload, err := parseLBTCDepositPayload(encodedPayload)
	require.NoError(t, err)
	expected := "0x5c70a5050000000000000000000000000000000000000000000000000000000000aa36a7000000000000000000000000845f8e3c214d8d0e4d83fc094f302aa26a12a0bc0000000000000000000000000000000000000000000000000000000000014a34000000000000000000000000845f8e3c214d8d0e4d83fc094f302aa26a12a0bc00000000000000000000000062f10ce5b727edf787ea45776bd050308a61150800000000000000000000000000000000000000000000000000000000000003e60000000000000000000000000000000000000000000000000000000000000006"
	assert.Equal(t, expected, hexutil.Encode(payload))
}

func Test_MockLogPoller(t *testing.T) {
	lggr := logger.TestLogger(t)
	payload := []byte("0x1111")
	payloadHash := sha256.Sum256(payload)
	t.Run("found one", func(t *testing.T) {
		lp := lpmocks.NewLogPoller(t)
		reader, err := NewLBTCReader(lggr, "job_1", utils.RandomAddress(), lp, false)
		require.NoError(t, err)
		lp.On("IndexedLogsByTxHash", mock.Anything, reader.eventID, reader.transmitterAddress, mock.Anything).
			Return([]logpoller.Log{
				LogWithPayload(t, 20, payload),
			}, nil)

		data, err := reader.GetLBTCMessageInTx(context.Background(), payloadHash, "0x0001")
		assert.NoError(t, err)
		assert.Equal(t, payload, data)
	})

	t.Run("found multiple", func(t *testing.T) {
		lp := lpmocks.NewLogPoller(t)
		reader, err := NewLBTCReader(lggr, "job_1", utils.RandomAddress(), lp, false)
		require.NoError(t, err)
		lp.On("IndexedLogsByTxHash", mock.Anything, reader.eventID, reader.transmitterAddress, mock.Anything).
			Return([]logpoller.Log{
				LogWithPayload(t, 10, []byte("0x1110")),
				LogWithPayload(t, 20, payload),
				LogWithPayload(t, 30, []byte("0x2222")),
			}, nil)

		data, err := reader.GetLBTCMessageInTx(context.Background(), payloadHash, "0x0001")
		assert.NoError(t, err)
		assert.Equal(t, payload, data)
	})

	t.Run("found multiple none match", func(t *testing.T) {
		lp := lpmocks.NewLogPoller(t)
		reader, err := NewLBTCReader(lggr, "job_1", utils.RandomAddress(), lp, false)
		require.NoError(t, err)
		lp.On("IndexedLogsByTxHash", mock.Anything, reader.eventID, reader.transmitterAddress, mock.Anything).
			Return([]logpoller.Log{
				LogWithPayload(t, 10, []byte("0x1110")),
				LogWithPayload(t, 30, []byte("0x2222")),
			}, nil)

		data, err := reader.GetLBTCMessageInTx(context.Background(), payloadHash, "0x0001")
		assert.Nil(t, data)
		assert.Errorf(t, err, "payload with hash=%s not found in logs", payloadHash)
	})

	t.Run("no logs found", func(t *testing.T) {
		lp := lpmocks.NewLogPoller(t)
		reader, err := NewLBTCReader(lggr, "job_1", utils.RandomAddress(), lp, false)
		require.NoError(t, err)
		lp.On("IndexedLogsByTxHash", mock.Anything, reader.eventID, reader.transmitterAddress, mock.Anything).
			Return([]logpoller.Log{}, nil)

		data, err := reader.GetLBTCMessageInTx(context.Background(), payloadHash, "0x0001")
		assert.Nil(t, data)
		assert.Errorf(t, err, "payload with hash=%s not found in logs", payloadHash)
	})

	t.Run("cache hit", func(t *testing.T) {
		rCache := cache.New(cache.NoExpiration, cache.NoExpiration)
		err := rCache.Add("lbtc-0x0001", []logpoller.Log{LogWithPayload(t, 20, payload)}, cache.NoExpiration)
		require.NoError(t, err)
		r, err := NewLBTCReaderWithCache(lggr, "job_1", utils.RandomAddress(), nil, rCache, false)
		require.NoError(t, err)
		data, err := r.GetLBTCMessageInTx(context.Background(), payloadHash, "0x0001")
		assert.NoError(t, err)
		assert.Equal(t, payload, data)
	})
}

func Test_SimulatedLogPoller_FoundMultiple(t *testing.T) {
	lggr := logger.TestLogger(t)
	chainID := testutils.NewRandomEVMChainID()
	db := pgtest.NewSqlxDB(t)
	o := logpoller.NewORM(chainID, db, lggr)

	transmitter := utils.RandomAddress()
	payload := []byte("0x1111")
	payloadHash := sha256.Sum256(payload)
	logs := []types.Log{
		EthLogWithPayload(t, 10, transmitter, []byte("0x2222")),
		EthLogWithPayload(t, 20, utils.RandomAddress(), payload),
		EthLogWithPayload(t, 30, transmitter, payload),
	}

	ec := evmclimocks.NewClient(t)
	head := types2.NewHead(big.NewInt(1), common.Hash{}, common.Hash{}, 0, ubig.New(chainID))
	ec.On("HeadByNumber", mock.Anything, mock.Anything).Return(&head, nil)
	ec.On("FilterLogs", mock.Anything, mock.Anything).Return(logs, nil)
	ec.On("ConfiguredChainID").Return(chainID, nil)

	lpOpts := logpoller.Opts{
		PollPeriod:               time.Hour,
		FinalityDepth:            1,
		BackfillBatchSize:        1,
		RpcBatchSize:             1,
		KeepFinalizedBlocksDepth: 100,
	}
	headTracker := headtracker.NewSimulatedHeadTracker(ec, lpOpts.UseFinalityTag, lpOpts.FinalityDepth)
	lp := logpoller.NewLogPoller(o, ec, lggr, headTracker, lpOpts)
	lp.PollAndSaveLogs(context.Background(), 1)

	reader, err := NewLBTCReader(lggr, "job_1", transmitter, lp, true)
	require.NoError(t, err)

	data, err := reader.GetLBTCMessageInTx(context.Background(), payloadHash, common.Hash{}.Hex())
	assert.NoError(t, err)
	assert.Equal(t, payload, data)
}

func EthLogWithPayload(t *testing.T, logIndex uint, transmitter common.Address, payload []byte) types.Log {
	encodedPayload, err := abihelpers.ABIEncode(LBTC_PAYLOAD_ABI, payload)
	require.NoError(t, err)
	payloadHash := sha256.Sum256(payload)
	topics := make([]common.Hash, 4)
	topics[0] = crypto.Keccak256Hash([]byte("DepositToBridge(address,bytes32,bytes32,bytes)"))
	topics[3] = common.BytesToHash(payloadHash[:])
	return types.Log{
		Address:     transmitter,
		Topics:      topics,
		Data:        encodedPayload,
		BlockNumber: 1,
		TxHash:      common.Hash{},
		TxIndex:     1,
		BlockHash:   common.Hash{},
		Index:       logIndex,
		Removed:     false,
	}
}

func LogWithPayload(t *testing.T, index int64, payload []byte) logpoller.Log {
	payloadHash := sha256.Sum256(payload)
	topics := make([][]byte, 4)
	topics[3] = payloadHash[:]
	logData, err := abihelpers.ABIEncode(LBTC_PAYLOAD_ABI, payload)
	require.NoError(t, err)
	return logpoller.Log{
		LogIndex: index,
		Topics:   topics,
		Data:     logData,
	}
}
