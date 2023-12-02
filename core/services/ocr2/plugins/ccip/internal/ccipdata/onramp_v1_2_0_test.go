package ccipdata

import (
	"context"
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/hashlib"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

func TestHasherV1_2_0(t *testing.T) {
	sourceChainSelector, destChainSelector := uint64(1), uint64(4)
	onRampAddress := common.HexToAddress("0x5550000000000000000000000000000000000001")
	onRampABI := abihelpers.MustParseABI(evm_2_evm_onramp.EVM2EVMOnRampABI)

	hashingCtx := hashlib.NewKeccakCtx()
	ramp, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(onRampAddress, nil)
	require.NoError(t, err)
	hasher := NewLeafHasherV1_2_0(sourceChainSelector, destChainSelector, onRampAddress, hashingCtx, ramp)

	message := evm_2_evm_onramp.InternalEVM2EVMMessage{
		SourceChainSelector: sourceChainSelector,
		Sender:              common.HexToAddress("0x1110000000000000000000000000000000000001"),
		Receiver:            common.HexToAddress("0x2220000000000000000000000000000000000001"),
		SequenceNumber:      1337,
		GasLimit:            big.NewInt(100),
		Strict:              false,
		Nonce:               1337,
		FeeToken:            common.Address{},
		FeeTokenAmount:      big.NewInt(1),
		Data:                []byte{},
		TokenAmounts:        []evm_2_evm_onramp.ClientEVMTokenAmount{{Token: common.HexToAddress("0x4440000000000000000000000000000000000001"), Amount: big.NewInt(12345678900)}},
		SourceTokenData:     [][]byte{},
		MessageId:           [32]byte{},
	}

	data, err := onRampABI.Events[CCIPSendRequestedEventNameV1_2_0].Inputs.Pack(message)
	require.NoError(t, err)
	hash, err := hasher.HashLeaf(types.Log{Topics: []common.Hash{CCIPSendRequestEventSigV1_2_0}, Data: data})
	require.NoError(t, err)

	// NOTE: Must match spec
	require.Equal(t, "46ad031bfb052db2e4a2514fed8dc480b98e5ce4acb55d5640d91407e0d8a3e9", hex.EncodeToString(hash[:]))

	message = evm_2_evm_onramp.InternalEVM2EVMMessage{
		SourceChainSelector: sourceChainSelector,
		Sender:              common.HexToAddress("0x1110000000000000000000000000000000000001"),
		Receiver:            common.HexToAddress("0x2220000000000000000000000000000000000001"),
		SequenceNumber:      1337,
		GasLimit:            big.NewInt(100),
		Strict:              false,
		Nonce:               1337,
		FeeToken:            common.Address{},
		FeeTokenAmount:      big.NewInt(1e12),
		Data:                []byte("foo bar baz"),
		TokenAmounts: []evm_2_evm_onramp.ClientEVMTokenAmount{
			{Token: common.HexToAddress("0x4440000000000000000000000000000000000001"), Amount: big.NewInt(12345678900)},
			{Token: common.HexToAddress("0x6660000000000000000000000000000000000001"), Amount: big.NewInt(4204242)},
		},
		SourceTokenData: [][]byte{{0x2, 0x1}},
		MessageId:       [32]byte{},
	}

	data, err = onRampABI.Events[CCIPSendRequestedEventNameV1_2_0].Inputs.Pack(message)
	require.NoError(t, err)
	hash, err = hasher.HashLeaf(types.Log{Topics: []common.Hash{CCIPSendRequestEventSigV1_2_0}, Data: data})
	require.NoError(t, err)

	// NOTE: Must match spec
	require.Equal(t, "4362a13a42e52ff5ce4324e7184dc7aa41704c3146bc842d35d95b94b32a78b6", hex.EncodeToString(hash[:]))
}

func TestLogPollerClient_GetSendRequestsBetweenSeqNums(t *testing.T) {
	onRampAddr := utils.RandomAddress()
	seqNum := uint64(100)
	limit := uint64(10)
	lggr := logger.TestLogger(t)

	lp := mocks.NewLogPoller(t)
	onRampV2, err := NewOnRampV1_2_0(lggr, 1, 1, onRampAddr, lp, nil)
	require.NoError(t, err)
	lp.On("LogsDataWordRange",
		onRampV2.sendRequestedEventSig,
		onRampAddr,
		onRampV2.sendRequestedSeqNumberWord,
		abihelpers.EvmWord(seqNum),
		abihelpers.EvmWord(seqNum+limit),
		logpoller.Finalized,
		mock.Anything,
	).Return([]logpoller.Log{}, nil)

	events, err := onRampV2.GetSendRequestsBetweenSeqNums(context.Background(), seqNum, seqNum+limit)
	assert.NoError(t, err)
	assert.Empty(t, events)
	lp.AssertExpectations(t)
}
