package messagehasher

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/message_hasher"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/ccipevm"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/assert"
)

func TestMessageHasher(t *testing.T) {
	// Deploy messageHasher contract
	ctx := testutils.Context(t)
	d := testSetup(t, ctx)

	// Setup random msg data
	metadataHash := utils.RandomBytes32()

	sourceTokenData := make([]byte, rand.Intn(2048))
	_, err := rand.Read(sourceTokenData)
	assert.NoError(t, err)

	sourceChain := rand.Uint64()
	seqNum := rand.Uint64()
	chainFeeLimit := rand.Uint64()
	nonce := rand.Uint64()
	strict := rand.Intn(2) == 1
	feeTokenAmount := rand.Uint64()

	data := make([]byte, rand.Intn(2048))
	_, err = rand.Read(data)
	assert.NoError(t, err)

	sourceTokenDatas := make([][]byte, rand.Intn(10))
	for i := range sourceTokenDatas {
		sourceTokenDatas[i] = sourceTokenData
	}

	numTokenAmounts := rand.Intn(50)
	tokenAmounts := make([]cciptypes.TokenAmount, 0, numTokenAmounts)
	for i := 0; i < numTokenAmounts; i++ {
		tokenAmounts = append(tokenAmounts, cciptypes.TokenAmount{
			Token:  types.Account(utils.RandomAddress().String()),
			Amount: big.NewInt(0).SetUint64(rand.Uint64()),
		})
	}
	ccipMsg := cciptypes.CCIPMsg{
		CCIPMsgBaseDetails: cciptypes.CCIPMsgBaseDetails{
			SourceChain: cciptypes.ChainSelector(sourceChain),
			SeqNum:      cciptypes.SeqNum(seqNum),
		},
		ChainFeeLimit:   cciptypes.NewBigInt(big.NewInt(0).SetUint64(chainFeeLimit)),
		Nonce:           nonce,
		Sender:          types.Account(utils.RandomAddress().String()),
		Receiver:        types.Account(utils.RandomAddress().String()),
		Strict:          strict,
		FeeToken:        types.Account(utils.RandomAddress().String()),
		FeeTokenAmount:  cciptypes.NewBigInt(big.NewInt(0).SetUint64(feeTokenAmount)),
		Data:            data,
		TokenAmounts:    tokenAmounts,
		SourceTokenData: sourceTokenDatas,
	}

	evmTokenAmounts := make([]message_hasher.ClientEVMTokenAmount, 0, len(ccipMsg.TokenAmounts))
	for _, ta := range ccipMsg.TokenAmounts {
		evmTokenAmounts = append(evmTokenAmounts, message_hasher.ClientEVMTokenAmount{
			Token:  common.HexToAddress(string(ta.Token)),
			Amount: ta.Amount,
		})
	}
	evmMsg := message_hasher.InternalEVM2EVMMessage{
		SourceChainSelector: uint64(ccipMsg.SourceChain),
		Sender:              common.HexToAddress(string(ccipMsg.Sender)),
		Receiver:            common.HexToAddress(string(ccipMsg.Receiver)),
		SequenceNumber:      uint64(ccipMsg.SeqNum),
		GasLimit:            ccipMsg.ChainFeeLimit.Int,
		Strict:              ccipMsg.Strict,
		Nonce:               ccipMsg.Nonce,
		FeeToken:            common.HexToAddress(string(ccipMsg.FeeToken)),
		FeeTokenAmount:      ccipMsg.FeeTokenAmount.Int,
		Data:                ccipMsg.Data,
		TokenAmounts:        evmTokenAmounts,
		SourceTokenData:     ccipMsg.SourceTokenData,
	}

	h, err := d.contract.Hash(&bind.CallOpts{Context: ctx}, evmMsg, metadataHash)
	assert.NoError(t, err)

	evmMsgHasher := ccipevm.NewMessageHasherV1(metadataHash)
	h2, err := evmMsgHasher.Hash(ctx, ccipMsg)
	assert.NoError(t, err)

	assert.Equal(t, fmt.Sprintf("%x", h), strings.TrimPrefix(h2.String(), "0x"))
}

type testSetupData struct {
	contractAddr common.Address
	contract     *message_hasher.MessageHasher
	sb           *backends.SimulatedBackend
	auth         *bind.TransactOpts
}

const chainID = 1337

func testSetup(t *testing.T, ctx context.Context) *testSetupData {
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
	address, _, _, err := message_hasher.DeployMessageHasher(auth, simulatedBackend)
	assert.NoError(t, err)
	simulatedBackend.Commit()

	// Setup contract client
	contract, err := message_hasher.NewMessageHasher(address, simulatedBackend)
	assert.NoError(t, err)

	return &testSetupData{
		contractAddr: address,
		contract:     contract,
		sb:           simulatedBackend,
		auth:         auth,
	}
}
