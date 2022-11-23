package ccip

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/hasher"
)

func TestTollHasher(t *testing.T) {
	sourceChainId, destChainId := big.NewInt(1), big.NewInt(4)
	onRampAddress := common.HexToAddress("0x5550000000000000000000000000000000000001")

	hashingCtx := hasher.NewKeccakCtx()

	hasher := NewTollLeafHasher(sourceChainId, destChainId, onRampAddress, hashingCtx)

	message := evm_2_evm_toll_onramp.CCIPEVM2EVMTollMessage{
		SourceChainId:     sourceChainId,
		SequenceNumber:    1337,
		Sender:            common.HexToAddress("0x1110000000000000000000000000000000000001"),
		Receiver:          common.HexToAddress("0x2220000000000000000000000000000000000001"),
		Data:              []byte{},
		TokensAndAmounts:  []evm_2_evm_toll_onramp.CCIPEVMTokenAndAmount{{Token: common.HexToAddress("0x4440000000000000000000000000000000000001"), Amount: big.NewInt(12345678900)}},
		GasLimit:          big.NewInt(100),
		FeeTokenAndAmount: evm_2_evm_toll_onramp.CCIPEVMTokenAndAmount{Token: common.HexToAddress("0x3330000000000000000000000000000000000001"), Amount: big.NewInt(987654321)},
	}

	hash, err := hasher.HashLeaf(generateTollLog(t, message))
	require.NoError(t, err)

	// NOTE: Must match spec
	require.Equal(t, "d4504baca27221b294969ab5a2989e2121cb3577a209b85d7d83371b3429df4d", hex.EncodeToString(hash[:]))

	message = evm_2_evm_toll_onramp.CCIPEVM2EVMTollMessage{
		SourceChainId:  sourceChainId,
		SequenceNumber: 1337,
		Sender:         common.HexToAddress("0x1110000000000000000000000000000000000001"),
		Receiver:       common.HexToAddress("0x2220000000000000000000000000000000000001"),
		Data:           []byte("foo bar baz"),
		TokensAndAmounts: []evm_2_evm_toll_onramp.CCIPEVMTokenAndAmount{
			{Token: common.HexToAddress("0x4440000000000000000000000000000000000001"), Amount: big.NewInt(12345678900)},
			{Token: common.HexToAddress("0x6660000000000000000000000000000000000001"), Amount: big.NewInt(4204242)},
		},
		GasLimit: big.NewInt(100),
		FeeTokenAndAmount: evm_2_evm_toll_onramp.CCIPEVMTokenAndAmount{
			Token: common.HexToAddress("0x3330000000000000000000000000000000000001"), Amount: big.NewInt(987654321)},
	}

	hash, err = hasher.HashLeaf(generateTollLog(t, message))
	require.NoError(t, err)

	// NOTE: Must match spec
	require.Equal(t, "c0b2bac538afab5af9c654028ff27f3a3cc5aa9e1082efc70656b8467dd41fb2", hex.EncodeToString(hash[:]))
}

func generateTollLog(t *testing.T, message evm_2_evm_toll_onramp.CCIPEVM2EVMTollMessage) types.Log {
	pack, err := MakeTollCCIPMsgArgs().Pack(message)
	require.NoError(t, err)

	return types.Log{
		Topics: []common.Hash{CCIPTollSendRequested},
		Data:   pack,
	}
}

func TestMetaDataHash(t *testing.T) {
	sourceChainId, destChainId := big.NewInt(1), big.NewInt(4)
	onRampAddress := common.HexToAddress("0x5550000000000000000000000000000000000001")
	ctx := hasher.NewKeccakCtx()
	hash := getMetaDataHash(ctx, ctx.Hash([]byte("EVM2EVMSubscriptionMessagePlus")), sourceChainId, onRampAddress, destChainId)
	require.Equal(t, "e8b93c9d01a7a72ec6c7235e238701cf1511b267a31fdb78dd342649ee58c08d", hex.EncodeToString(hash[:]))
}
