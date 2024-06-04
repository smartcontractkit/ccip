package ccipexec

import (
	"bytes"
	"context"
	"encoding/binary"
	"math"
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
	txmmocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipcalc"
	ccipdatamocks "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/prices"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/testhelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/tokendata"
)

type testCase struct {
	name                     string
	reqs                     []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta
	inflight                 *big.Int
	tokenLimit, destGasPrice *big.Int
	srcPrices, dstPrices     map[cciptypes.Address]*big.Int
	offRampNoncesBySender    map[cciptypes.Address]uint64
	srcToDestTokens          map[cciptypes.Address]cciptypes.Address
	expectedSeqNrs           []ccip.ObservedMessage
	expectedStates           []messageExecStatus
	mockTxm                  func(m *MockTxmFake)
}

func Test_validateSendRequests(t *testing.T) {
	testCases := []struct {
		name             string
		seqNums          []uint64
		providedInterval cciptypes.CommitStoreInterval
		expErr           bool
	}{
		{
			name:             "zero interval no seq nums",
			seqNums:          nil,
			providedInterval: cciptypes.CommitStoreInterval{Min: 0, Max: 0},
			expErr:           true,
		},
		{
			name:             "exp 1 seq num got none",
			seqNums:          nil,
			providedInterval: cciptypes.CommitStoreInterval{Min: 1, Max: 1},
			expErr:           true,
		},
		{
			name:             "exp 10 seq num got none",
			seqNums:          nil,
			providedInterval: cciptypes.CommitStoreInterval{Min: 1, Max: 10},
			expErr:           true,
		},
		{
			name:             "got 1 seq num as expected",
			seqNums:          []uint64{1},
			providedInterval: cciptypes.CommitStoreInterval{Min: 1, Max: 1},
			expErr:           false,
		},
		{
			name:             "got 5 seq num as expected",
			seqNums:          []uint64{11, 12, 13, 14, 15},
			providedInterval: cciptypes.CommitStoreInterval{Min: 11, Max: 15},
			expErr:           false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sendReqs := make([]cciptypes.EVM2EVMMessageWithTxMeta, 0, len(tc.seqNums))
			for _, seqNum := range tc.seqNums {
				sendReqs = append(sendReqs, cciptypes.EVM2EVMMessageWithTxMeta{
					EVM2EVMMessage: cciptypes.EVM2EVMMessage{SequenceNumber: seqNum},
				})
			}
			err := validateSendRequests(sendReqs, tc.providedInterval)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

type delayedTokenDataWorker struct {
	delay time.Duration
	tokendata.Worker
}

func (m delayedTokenDataWorker) GetMsgTokenData(ctx context.Context, msg cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta) ([][]byte, error) {
	time.Sleep(m.delay)
	return nil, ctx.Err()
}

func TestExecutionReportingPlugin_getTokenDataWithCappedLatency(t *testing.T) {
	testCases := []struct {
		name               string
		allowedWaitingTime time.Duration
		workerLatency      time.Duration
		expErr             bool
	}{
		{
			name:               "happy flow",
			allowedWaitingTime: 10 * time.Millisecond,
			workerLatency:      time.Nanosecond,
			expErr:             false,
		},
		{
			name:               "worker takes long to reply",
			allowedWaitingTime: 10 * time.Millisecond,
			workerLatency:      20 * time.Millisecond,
			expErr:             true,
		},
	}

	ctx := testutils.Context(t)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bs := &BaseBatchingStrategy{}
			tokenDataWorker := delayedTokenDataWorker{delay: tc.workerLatency}

			msg := cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{
				EVM2EVMMessage: cciptypes.EVM2EVMMessage{TokenAmounts: make([]cciptypes.TokenAmount, 1)},
			}

			_, _, err := bs.getTokenDataWithTimeout(ctx, msg, tc.allowedWaitingTime, tokenDataWorker)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestBatchingStrategies(t *testing.T) {
	sender1 := ccipcalc.HexToAddress("0xa")
	destNative := ccipcalc.HexToAddress("0xb")
	srcNative := ccipcalc.HexToAddress("0xc")

	msg1 := createTestMessage(1, sender1, 1, srcNative, big.NewInt(1e9), false, nil)

	msg2 := msg1
	msg2.Executed = true

	msg3 := msg1
	msg3.Executed = true
	msg3.Finalized = true

	msg4 := msg1
	msg4.TokenAmounts = []cciptypes.TokenAmount{
		{Token: srcNative, Amount: big.NewInt(100)},
	}

	msg5 := msg4
	msg5.SequenceNumber = msg5.SequenceNumber + 1
	msg5.Nonce = msg5.Nonce + 1

	msgId1 := generateMessageIDFromInt(1)
	msgId2 := generateMessageIDFromInt(2)
	msgId3 := generateMessageIDFromInt(3)

	testCases := []testCase{
		{
			name:                  "single message no tokens",
			reqs:                  []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{msg1},
			inflight:              big.NewInt(0),
			tokenLimit:            big.NewInt(0),
			destGasPrice:          big.NewInt(10),
			srcPrices:             map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1)},
			dstPrices:             map[cciptypes.Address]*big.Int{destNative: big.NewInt(1)},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 0},
			expectedSeqNrs:        []ccip.ObservedMessage{{SeqNr: uint64(1)}},
			expectedStates:        []messageExecStatus{newMessageExecState(msg1.SequenceNumber, msg1.MessageID, AddedToBatch)},
		},
		{
			name:                  "gasPriceEstimator returns error",
			reqs:                  []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{msg1},
			inflight:              big.NewInt(0),
			tokenLimit:            big.NewInt(0),
			destGasPrice:          big.NewInt(10),
			srcPrices:             map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1)},
			dstPrices:             map[cciptypes.Address]*big.Int{destNative: big.NewInt(1)},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 0},
		},
		{
			name:                  "executed non finalized messages should be skipped",
			reqs:                  []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{msg2},
			inflight:              big.NewInt(0),
			tokenLimit:            big.NewInt(0),
			destGasPrice:          big.NewInt(10),
			srcPrices:             map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1)},
			dstPrices:             map[cciptypes.Address]*big.Int{destNative: big.NewInt(1)},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 0},
			expectedStates:        []messageExecStatus{newMessageExecState(msg2.SequenceNumber, msg2.MessageID, AlreadyExecuted)},
		},
		{
			name:                  "finalized executed log",
			reqs:                  []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{msg3},
			inflight:              big.NewInt(0),
			tokenLimit:            big.NewInt(0),
			destGasPrice:          big.NewInt(10),
			srcPrices:             map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1)},
			dstPrices:             map[cciptypes.Address]*big.Int{destNative: big.NewInt(1)},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 0},
			expectedStates:        []messageExecStatus{newMessageExecState(msg3.SequenceNumber, msg3.MessageID, AlreadyExecuted)},
		},
		{
			name:                  "dst token price does not exist",
			reqs:                  []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{msg1},
			inflight:              big.NewInt(0),
			tokenLimit:            big.NewInt(0),
			destGasPrice:          big.NewInt(10),
			srcPrices:             map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1)},
			dstPrices:             map[cciptypes.Address]*big.Int{},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 0},
			expectedStates:        []messageExecStatus{newMessageExecState(msg1.SequenceNumber, msg1.MessageID, TokenNotInDestTokenPrices)},
		},
		{
			name:                  "src token price does not exist",
			reqs:                  []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{msg1},
			inflight:              big.NewInt(0),
			tokenLimit:            big.NewInt(0),
			destGasPrice:          big.NewInt(10),
			srcPrices:             map[cciptypes.Address]*big.Int{},
			dstPrices:             map[cciptypes.Address]*big.Int{destNative: big.NewInt(1)},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 0},
			expectedStates:        []messageExecStatus{newMessageExecState(msg1.SequenceNumber, msg1.MessageID, TokenNotInSrcTokenPrices)},
		},
		{
			name:         "message with tokens is not executed if limit is reached",
			reqs:         []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{msg4},
			inflight:     big.NewInt(0),
			tokenLimit:   big.NewInt(2),
			destGasPrice: big.NewInt(10),
			srcPrices:    map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1e18)},
			dstPrices:    map[cciptypes.Address]*big.Int{destNative: big.NewInt(1e18)},
			srcToDestTokens: map[cciptypes.Address]cciptypes.Address{
				srcNative: destNative,
			},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 0},
			expectedStates:        []messageExecStatus{newMessageExecState(msg4.SequenceNumber, msg4.MessageID, AggregateTokenLimitExceeded)},
		},
		{
			name:         "message with tokens is not executed if limit is reached when inflight is full",
			reqs:         []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{msg5},
			inflight:     new(big.Int).Mul(big.NewInt(1e18), big.NewInt(100)),
			tokenLimit:   big.NewInt(19),
			destGasPrice: big.NewInt(10),
			srcPrices:    map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1e18)},
			dstPrices:    map[cciptypes.Address]*big.Int{destNative: big.NewInt(1e18)},
			srcToDestTokens: map[cciptypes.Address]cciptypes.Address{
				srcNative: destNative,
			},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 1},
			expectedStates:        []messageExecStatus{newMessageExecState(msg5.SequenceNumber, msg5.MessageID, AggregateTokenLimitExceeded)},
		},
		{
			name:                  "skip when nonce doesn't match chain value",
			reqs:                  []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{msg1},
			inflight:              big.NewInt(0),
			tokenLimit:            big.NewInt(0),
			destGasPrice:          big.NewInt(10),
			srcPrices:             map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1)},
			dstPrices:             map[cciptypes.Address]*big.Int{destNative: big.NewInt(1)},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 123},
			expectedStates:        []messageExecStatus{newMessageExecState(msg1.SequenceNumber, msg1.MessageID, InvalidNonce)},
		},
		{
			name:                  "skip when nonce not found",
			reqs:                  []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{msg1},
			inflight:              big.NewInt(0),
			tokenLimit:            big.NewInt(0),
			destGasPrice:          big.NewInt(10),
			srcPrices:             map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1)},
			dstPrices:             map[cciptypes.Address]*big.Int{destNative: big.NewInt(1)},
			offRampNoncesBySender: map[cciptypes.Address]uint64{},
			expectedStates:        []messageExecStatus{newMessageExecState(msg1.SequenceNumber, msg1.MessageID, MissingNonce)},
		},
		{
			name: "unordered messages",
			reqs: []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{
				{
					EVM2EVMMessage: cciptypes.EVM2EVMMessage{
						SequenceNumber: 10,
						FeeTokenAmount: big.NewInt(1e9),
						Sender:         sender1,
						Nonce:          0,
						GasLimit:       big.NewInt(1),
						Data:           bytes.Repeat([]byte{'a'}, 1000),
						FeeToken:       srcNative,
						MessageID:      [32]byte{},
					},
					BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
				},
			},
			inflight:              big.NewInt(0),
			tokenLimit:            big.NewInt(0),
			destGasPrice:          big.NewInt(10),
			srcPrices:             map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1)},
			dstPrices:             map[cciptypes.Address]*big.Int{destNative: big.NewInt(1)},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 0},
			expectedSeqNrs:        []ccip.ObservedMessage{{SeqNr: uint64(10)}},
			expectedStates: []messageExecStatus{
				newMessageExecState(10, [32]byte{}, AddedToBatch),
			},
		},
		{
			name: "unordered messages not blocked by nonce",
			reqs: []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{
				{
					EVM2EVMMessage: cciptypes.EVM2EVMMessage{
						SequenceNumber: 9,
						FeeTokenAmount: big.NewInt(1e9),
						Sender:         sender1,
						Nonce:          5,
						GasLimit:       big.NewInt(1),
						Data:           bytes.Repeat([]byte{'a'}, 1000),
						FeeToken:       srcNative,
						MessageID:      [32]byte{},
					},
					BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
				},
				{
					EVM2EVMMessage: cciptypes.EVM2EVMMessage{
						SequenceNumber: 10,
						FeeTokenAmount: big.NewInt(1e9),
						Sender:         sender1,
						Nonce:          0,
						GasLimit:       big.NewInt(1),
						Data:           bytes.Repeat([]byte{'a'}, 1000),
						FeeToken:       srcNative,
						MessageID:      [32]byte{},
					},
					BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
				},
			},
			inflight:              big.NewInt(0),
			tokenLimit:            big.NewInt(0),
			destGasPrice:          big.NewInt(10),
			srcPrices:             map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1)},
			dstPrices:             map[cciptypes.Address]*big.Int{destNative: big.NewInt(1)},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 3},
			expectedSeqNrs:        []ccip.ObservedMessage{{SeqNr: uint64(10)}},
			expectedStates: []messageExecStatus{
				newMessageExecState(9, [32]byte{}, InvalidNonce),
				newMessageExecState(10, [32]byte{}, AddedToBatch),
			},
		},
	}

	bestEffortTestCases := []testCase{
		{
			name: "skip when batch gas limit is reached",
			reqs: []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{
				{
					EVM2EVMMessage: cciptypes.EVM2EVMMessage{
						SequenceNumber: 10,
						FeeTokenAmount: big.NewInt(1e9),
						Sender:         sender1,
						Nonce:          1,
						GasLimit:       big.NewInt(1),
						Data:           bytes.Repeat([]byte{'a'}, 1000),
						FeeToken:       srcNative,
						MessageID:      [32]byte{},
					},
					BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
				},
				{
					EVM2EVMMessage: cciptypes.EVM2EVMMessage{
						SequenceNumber: 11,
						FeeTokenAmount: big.NewInt(1e9),
						Sender:         sender1,
						Nonce:          2,
						GasLimit:       big.NewInt(math.MaxInt64),
						Data:           bytes.Repeat([]byte{'a'}, 1000),
						FeeToken:       srcNative,
						MessageID:      [32]byte{},
					},
					BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
				},
				{
					EVM2EVMMessage: cciptypes.EVM2EVMMessage{
						SequenceNumber: 12,
						FeeTokenAmount: big.NewInt(1e9),
						Sender:         sender1,
						Nonce:          3,
						GasLimit:       big.NewInt(1),
						Data:           bytes.Repeat([]byte{'a'}, 1000),
						FeeToken:       srcNative,
						MessageID:      [32]byte{},
					},
					BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
				},
			},
			inflight:              big.NewInt(0),
			tokenLimit:            big.NewInt(0),
			destGasPrice:          big.NewInt(10),
			srcPrices:             map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1)},
			dstPrices:             map[cciptypes.Address]*big.Int{destNative: big.NewInt(1)},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 0},
			expectedSeqNrs:        []ccip.ObservedMessage{{SeqNr: uint64(10)}},
			expectedStates: []messageExecStatus{
				newMessageExecState(10, [32]byte{}, AddedToBatch),
				newMessageExecState(11, [32]byte{}, InsufficientRemainingBatchGas),
				newMessageExecState(12, [32]byte{}, InvalidNonce),
			},
		},
		{
			name: "some messages skipped after hitting max batch data len",
			reqs: []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{
				{
					EVM2EVMMessage: cciptypes.EVM2EVMMessage{
						SequenceNumber: 10,
						FeeTokenAmount: big.NewInt(1e9),
						Sender:         sender1,
						Nonce:          1,
						GasLimit:       big.NewInt(1),
						Data:           bytes.Repeat([]byte{'a'}, 1000),
						FeeToken:       srcNative,
						MessageID:      [32]byte{},
					},
					BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
				},
				{
					EVM2EVMMessage: cciptypes.EVM2EVMMessage{
						SequenceNumber: 11,
						FeeTokenAmount: big.NewInt(1e9),
						Sender:         sender1,
						Nonce:          2,
						GasLimit:       big.NewInt(1),
						Data:           bytes.Repeat([]byte{'a'}, MaxDataLenPerBatch-500), // skipped from batch
						FeeToken:       srcNative,
						MessageID:      [32]byte{},
					},
					BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
				},
				{
					EVM2EVMMessage: cciptypes.EVM2EVMMessage{
						SequenceNumber: 12,
						FeeTokenAmount: big.NewInt(1e9),
						Sender:         sender1,
						Nonce:          3,
						GasLimit:       big.NewInt(1),
						Data:           bytes.Repeat([]byte{'a'}, 1000),
						FeeToken:       srcNative,
						MessageID:      [32]byte{},
					},
					BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
				},
			},
			inflight:              big.NewInt(0),
			tokenLimit:            big.NewInt(0),
			destGasPrice:          big.NewInt(10),
			srcPrices:             map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1)},
			dstPrices:             map[cciptypes.Address]*big.Int{destNative: big.NewInt(1)},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 0},
			expectedSeqNrs:        []ccip.ObservedMessage{{SeqNr: uint64(10)}},
			expectedStates: []messageExecStatus{
				newMessageExecState(10, [32]byte{}, AddedToBatch),
				newMessageExecState(11, [32]byte{}, InsufficientRemainingBatchDataLength),
				newMessageExecState(12, [32]byte{}, InvalidNonce),
			},
		},
	}

	// TODO: add following scenarios
	// 1 message, no ZKO (multiple answers and none of them is ZKO) => batch with 1st
	// 1 message, ZKO (with multiple answers and 1 is ZKO) => empty batch/no batch
	// 1 message, no ZKO (with multiple answers and none is ZKO) => empty batch/no batch
	// 2 messages, 1st is pending, 2nd is not ZKO => batch with 1st
	// 2 messages, 1st is ZKO, 2nd is ZKO => empty batch/no batch
	// 3 messages, 1st is not ZKO, 2nd is not ZKO, 3rd is not ZKO => batch with 1st
	// 3 messages, 1st is ZKO, 2nd is not ZKO, 3rd is not ZKO => batch with 2nd
	// 3 messages, 1st is ZKO, 2nd is ZKO, 3rd is not ZKO => batch with 3rd
	specificZkOverflowTestCases := []testCase{
		// {
		// 	name: "batch size is 1",
		// 	reqs: []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{
		// 		{
		// 			EVM2EVMMessage: cciptypes.EVM2EVMMessage{
		// 				SequenceNumber: 10,
		// 				FeeTokenAmount: big.NewInt(1e9),
		// 				Sender:         sender1,
		// 				Nonce:          1,
		// 				GasLimit:       big.NewInt(1),
		// 				Data:           bytes.Repeat([]byte{'a'}, 1000),
		// 				FeeToken:       srcNative,
		// 				MessageID:      [32]byte{},
		// 			},
		// 			BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
		// 		},
		// 		{
		// 			EVM2EVMMessage: cciptypes.EVM2EVMMessage{
		// 				SequenceNumber: 11,
		// 				FeeTokenAmount: big.NewInt(1e9),
		// 				Sender:         sender1,
		// 				Nonce:          2,
		// 				GasLimit:       big.NewInt(1),
		// 				Data:           bytes.Repeat([]byte{'a'}, 1000),
		// 				FeeToken:       srcNative,
		// 				MessageID:      [32]byte{},
		// 			},
		// 			BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
		// 		},
		// 		{
		// 			EVM2EVMMessage: cciptypes.EVM2EVMMessage{
		// 				SequenceNumber: 12,
		// 				FeeTokenAmount: big.NewInt(1e9),
		// 				Sender:         sender1,
		// 				Nonce:          3,
		// 				GasLimit:       big.NewInt(1),
		// 				Data:           bytes.Repeat([]byte{'a'}, 1000),
		// 				FeeToken:       srcNative,
		// 				MessageID:      [32]byte{},
		// 			},
		// 			BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
		// 		},
		// 	},
		// 	inflight:              big.NewInt(0),
		// 	tokenLimit:            big.NewInt(0),
		// 	destGasPrice:          big.NewInt(10),
		// 	srcPrices:             map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1)},
		// 	dstPrices:             map[cciptypes.Address]*big.Int{destNative: big.NewInt(1)},
		// 	offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 0},
		// 	expectedSeqNrs:        []ccip.ObservedMessage{{SeqNr: uint64(10)}},
		// 	expectedStates: []messageExecStatus{
		// 		newMessageExecState(10, [32]byte{}, AddedToBatch),
		// 	},
		// },
		{
			name: "snooze failed message and add next message to batch",
			reqs: []cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{
				{
					EVM2EVMMessage: cciptypes.EVM2EVMMessage{
						SequenceNumber: 10,
						FeeTokenAmount: big.NewInt(1e9),
						Sender:         sender1,
						Nonce:          0,
						GasLimit:       big.NewInt(1),
						Data:           bytes.Repeat([]byte{'a'}, 1000),
						FeeToken:       srcNative,
						MessageID:      msgId1,
					},
					BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
				},
				{
					EVM2EVMMessage: cciptypes.EVM2EVMMessage{
						SequenceNumber: 11,
						FeeTokenAmount: big.NewInt(1e9),
						Sender:         sender1,
						Nonce:          0,
						GasLimit:       big.NewInt(1),
						Data:           bytes.Repeat([]byte{'a'}, 1000),
						FeeToken:       srcNative,
						MessageID:      msgId2,
					},
					BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
				},
				{
					EVM2EVMMessage: cciptypes.EVM2EVMMessage{
						SequenceNumber: 12,
						FeeTokenAmount: big.NewInt(1e9),
						Sender:         sender1,
						Nonce:          0,
						GasLimit:       big.NewInt(1),
						Data:           bytes.Repeat([]byte{'a'}, 1000),
						FeeToken:       srcNative,
						MessageID:      msgId3,
					},
					BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
				},
			},
			inflight:              big.NewInt(0),
			tokenLimit:            big.NewInt(0),
			destGasPrice:          big.NewInt(10),
			srcPrices:             map[cciptypes.Address]*big.Int{srcNative: big.NewInt(1)},
			dstPrices:             map[cciptypes.Address]*big.Int{destNative: big.NewInt(1)},
			offRampNoncesBySender: map[cciptypes.Address]uint64{sender1: 0},
			expectedSeqNrs:        []ccip.ObservedMessage{{SeqNr: uint64(11)}},
			expectedStates: []messageExecStatus{
				newMessageExecState(10, msgId1, "zk_check_failed"),
				newMessageExecState(11, msgId2, AddedToBatch),
			},
			mockTxm: func(m *MockTxmFake) {
				m.On("FindTxsByIdempotencyPrefix", mock.Anything, hexutil.Encode(msgId1[:])).Return([]status{Failed}, nil)
				m.On("FindTxsByIdempotencyPrefix", mock.Anything, hexutil.Encode(msgId2[:])).Return([]status{}, nil)
				m.On("FindTxsByIdempotencyPrefix", mock.Anything, hexutil.Encode(msgId3[:])).Return([]status{}, nil)
			},
		},
	}

	t.Run("BestEffortBatchingStrategy", func(t *testing.T) {
		strategy := &BestEffortBatchingStrategy{}
		runBatchingStrategyTests(t, strategy, 300_000, append(testCases, bestEffortTestCases...))
	})

	t.Run("ZKOverflowBatchingStrategy", func(t *testing.T) {
		mockedTxManager := new(txmmocks.MockEvmTxManager)
		mockedTxManagerFake := new(MockTxmFake)
		strategy := &ZKOverflowBatchingStrategy{
			txManager:     mockedTxManager,
			txManagerFake: mockedTxManagerFake,
		}
		runBatchingStrategyTests(t, strategy, 1_000_000, append(testCases, specificZkOverflowTestCases...))
	})
}

// Function to set up and run tests for a given batching strategy
func runBatchingStrategyTests(t *testing.T, strategy BatchingStrategy, availableGas uint64, testCases []testCase) {

	destNative := ccipcalc.HexToAddress("0xb")

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			offRamp, _ := testhelpers.NewFakeOffRamp(t)
			lggr := logger.TestLogger(t)

			offRamp.SetSenderNonces(tc.offRampNoncesBySender)

			gasPriceEstimator := prices.NewMockGasPriceEstimatorExec(t)
			if tc.expectedSeqNrs != nil {
				gasPriceEstimator.On("EstimateMsgCostUSD", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(big.NewInt(0), nil)
			}

			if tc.expectedStates == nil && tc.expectedSeqNrs == nil {
				gasPriceEstimator.On("EstimateMsgCostUSD", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(big.NewInt(0), errors.New("error"))
			}

			// Mock calls to reader.
			mockOffRampReader := ccipdatamocks.NewOffRampReader(t)
			mockOffRampReader.On("ListSenderNonces", mock.Anything, mock.Anything).Return(tc.offRampNoncesBySender, nil).Maybe()

			// Mock calls to TXM
			if tc.mockTxm != nil {
				tc.mockTxm(strategy.(*ZKOverflowBatchingStrategy).txManagerFake.(*MockTxmFake))
			}

			// default case for ZKOverflowBatchingStrategy
			if strategyType := reflect.TypeOf(strategy); tc.mockTxm == nil && strategyType == reflect.TypeOf(&ZKOverflowBatchingStrategy{}) {
				strategy.(*ZKOverflowBatchingStrategy).txManagerFake.(*MockTxmFake).On("FindTxsByIdempotencyPrefix", mock.Anything, mock.Anything).Return([]status{}, nil)
			}

			batchContext := &BatchContext{
				ctx:                        context.Background(),
				report:                     commitReportWithSendRequests{sendRequestsWithMeta: tc.reqs},
				lggr:                       lggr,
				availableDataLen:           MaxDataLenPerBatch,
				availableGas:               availableGas,
				expectedNonces:             make(map[cciptypes.Address]uint64),
				sendersNonce:               tc.offRampNoncesBySender,
				sourceTokenPricesUSD:       tc.srcPrices,
				destTokenPricesUSD:         tc.dstPrices,
				gasPrice:                   tc.destGasPrice,
				sourceToDestToken:          tc.srcToDestTokens,
				inflightAggregateValue:     tc.inflight,
				aggregateTokenLimit:        tc.tokenLimit,
				tokenDataRemainingDuration: 5 * time.Second,
				tokenDataWorker:            tokendata.NewBackgroundWorker(map[cciptypes.Address]tokendata.Reader{}, 10, 5*time.Second, time.Hour),
				gasPriceEstimator:          gasPriceEstimator,
				destWrappedNative:          destNative,
				offchainConfig: cciptypes.ExecOffchainConfig{
					DestOptimisticConfirmations: 1,
					BatchGasLimit:               300_000,
					RelativeBoostPerWaitHour:    1,
				},
			}

			seqNrs, execStates := strategy.BuildBatch(batchContext)

			runAssertions(t, tc, seqNrs, execStates)
		})
	}
}

// Utility function to run common assertions
func runAssertions(t *testing.T, tc testCase, seqNrs []ccip.ObservedMessage, execStates []messageExecStatus) {
	if tc.expectedSeqNrs == nil {
		assert.Len(t, seqNrs, 0)
	} else {
		assert.Equal(t, tc.expectedSeqNrs, seqNrs)
	}

	if tc.expectedStates == nil {
		assert.Len(t, execStates, 0)
	} else {
		assert.Equal(t, tc.expectedStates, execStates)
	}
}

func createTestMessage(seqNr uint64, sender cciptypes.Address, nonce uint64, feeToken cciptypes.Address, feeAmount *big.Int, executed bool, data []byte) cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta {
	return cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta{
		EVM2EVMMessage: cciptypes.EVM2EVMMessage{
			SequenceNumber: seqNr,
			FeeTokenAmount: feeAmount,
			Sender:         sender,
			Nonce:          nonce,
			GasLimit:       big.NewInt(1),
			Strict:         false,
			Receiver:       "",
			Data:           data,
			TokenAmounts:   nil,
			FeeToken:       feeToken,
			MessageID:      [32]byte{},
		},
		BlockTimestamp: time.Date(2010, 1, 1, 12, 12, 12, 0, time.UTC),
		Executed:       executed,
	}
}

func generateMessageIDFromInt(input uint32) [32]byte {
	var messageID [32]byte
	binary.LittleEndian.PutUint32(messageID[:], input)
	return messageID
}

type MockTxmFake struct {
	mock.Mock
}

func (t *MockTxmFake) FindTxsByIdempotencyPrefix(ctx context.Context, msgIdPrefix string) ([]status, error) {
	args := t.Called(ctx, msgIdPrefix)
	return args.Get(0).([]status), args.Error(1)
}
