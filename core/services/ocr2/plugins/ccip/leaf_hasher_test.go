package ccip

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_subscription_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/hasher"
)

func TestSubscriptionHasher(t *testing.T) {
	sourceChainId, destChainId := big.NewInt(1), big.NewInt(4)
	onRampAddress := common.HexToAddress("0x5550000000000000000000000000000000000001")

	hashingCtx := hasher.NewKeccakCtx()

	hasher := NewSubscriptionLeafHasher(sourceChainId, destChainId, onRampAddress, hashingCtx)

	message := evm_2_evm_subscription_onramp.CCIPEVM2EVMSubscriptionMessage{
		SourceChainId:  sourceChainId,
		SequenceNumber: 1337,
		Sender:         common.HexToAddress("0x1110000000000000000000000000000000000001"),
		Receiver:       common.HexToAddress("0x2220000000000000000000000000000000000001"),
		Nonce:          666,
		Data:           []byte{},
		Tokens:         []common.Address{common.HexToAddress("0x4440000000000000000000000000000000000001")},
		Amounts:        []*big.Int{big.NewInt(12345678900)},
		GasLimit:       big.NewInt(100),
	}

	hash, err := hasher.HashLeaf(generateSubscriptionLog(t, message))
	require.NoError(t, err)

	// NOTE: Must match spec
	require.Equal(t, "cae032f60dc29a4d98e135908afa3f562674954c9d3378606e8b0473d27e94c9", hex.EncodeToString(hash[:]))

	message = evm_2_evm_subscription_onramp.CCIPEVM2EVMSubscriptionMessage{
		SourceChainId:  sourceChainId,
		SequenceNumber: 1337,
		Sender:         common.HexToAddress("0x1110000000000000000000000000000000000001"),
		Receiver:       common.HexToAddress("0x2220000000000000000000000000000000000001"),
		Nonce:          210,
		Data:           []byte("foo bar baz"),
		Tokens:         []common.Address{common.HexToAddress("0x4440000000000000000000000000000000000001"), common.HexToAddress("0x6660000000000000000000000000000000000001")},
		Amounts:        []*big.Int{big.NewInt(12345678900), big.NewInt(4204242)},
		GasLimit:       big.NewInt(100),
	}

	hash, err = hasher.HashLeaf(generateSubscriptionLog(t, message))
	require.NoError(t, err)

	// NOTE: Must match spec
	require.Equal(t, "aef2f373966c54aec50e619bacd6e66275f660c5b5ff3bd53b00386d345bcfa9", hex.EncodeToString(hash[:]))
}

func generateSubscriptionLog(t *testing.T, message evm_2_evm_subscription_onramp.CCIPEVM2EVMSubscriptionMessage) types.Log {
	pack, err := MakeSubscriptionCCIPMsgArgs().Pack(message)
	require.NoError(t, err)

	return types.Log{
		Topics: []common.Hash{CCIPSubSendRequested},
		Data:   pack,
	}
}

func TestTollHasher(t *testing.T) {
	sourceChainId, destChainId := big.NewInt(1), big.NewInt(4)
	onRampAddress := common.HexToAddress("0x5550000000000000000000000000000000000001")

	hashingCtx := hasher.NewKeccakCtx()

	hasher := NewTollLeafHasher(sourceChainId, destChainId, onRampAddress, hashingCtx)

	message := evm_2_evm_toll_onramp.CCIPEVM2EVMTollMessage{
		SourceChainId:  sourceChainId,
		SequenceNumber: 1337,
		Sender:         common.HexToAddress("0x1110000000000000000000000000000000000001"),
		Receiver:       common.HexToAddress("0x2220000000000000000000000000000000000001"),
		Data:           []byte{},
		Tokens:         []common.Address{common.HexToAddress("0x4440000000000000000000000000000000000001")},
		Amounts:        []*big.Int{big.NewInt(12345678900)},
		GasLimit:       big.NewInt(100),
		FeeToken:       common.HexToAddress("0x3330000000000000000000000000000000000001"),
		FeeTokenAmount: big.NewInt(987654321),
	}

	hash, err := hasher.HashLeaf(generateTollLog(t, message))
	require.NoError(t, err)

	// NOTE: Must match spec
	require.Equal(t, "9c014cce73a389409d5dbc863cb4d0054e61698bafb21eb88cafd670ee45ed12", hex.EncodeToString(hash[:]))

	message = evm_2_evm_toll_onramp.CCIPEVM2EVMTollMessage{
		SourceChainId:  sourceChainId,
		SequenceNumber: 1337,
		Sender:         common.HexToAddress("0x1110000000000000000000000000000000000001"),
		Receiver:       common.HexToAddress("0x2220000000000000000000000000000000000001"),
		Data:           []byte("foo bar baz"),
		Tokens:         []common.Address{common.HexToAddress("0x4440000000000000000000000000000000000001"), common.HexToAddress("0x6660000000000000000000000000000000000001")},
		Amounts:        []*big.Int{big.NewInt(12345678900), big.NewInt(4204242)},
		GasLimit:       big.NewInt(100),
		FeeToken:       common.HexToAddress("0x3330000000000000000000000000000000000001"),
		FeeTokenAmount: big.NewInt(987654321),
	}

	hash, err = hasher.HashLeaf(generateTollLog(t, message))
	require.NoError(t, err)

	// NOTE: Must match spec
	require.Equal(t, "b70e53658377bb46b430d3ca5bbfed10c1e97d82dd8feb0af896224b4bf890c8", hex.EncodeToString(hash[:]))
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
