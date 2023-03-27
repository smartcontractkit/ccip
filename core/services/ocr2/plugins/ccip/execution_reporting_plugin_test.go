package ccip_test

import (
	"context"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store_helper"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_offramp_helper"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/mock_afn_contract"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/price_registry"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/simple_message_receiver"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/merklemulti"
	"github.com/smartcontractkit/chainlink/core/utils"
)

type ExecutionContracts struct {
	// Has all the link and 100ETH
	user                       *bind.TransactOpts
	offRamp                    *evm_2_evm_offramp.EVM2EVMOffRamp
	commitStore                *commit_store_helper.CommitStoreHelper
	receiver                   *simple_message_receiver.SimpleMessageReceiver
	linkTokenAddress           common.Address
	destChainID, sourceChainID uint64
	destChain                  *backends.SimulatedBackend
}

func setupContractsForExecution(t *testing.T) ExecutionContracts {
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	destChainID := uint64(1337)
	sourceChainID := uint64(1338)
	destUser, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(0).SetUint64(destChainID))
	require.NoError(t, err)
	destChain := backends.NewSimulatedBackend(core.GenesisAlloc{
		destUser.From: {Balance: big.NewInt(0).Mul(big.NewInt(100), big.NewInt(1000000000000000000))}},
		10*ethconfig.Defaults.Miner.GasCeil) // 80M gas
	// Deploy link token
	destLinkTokenAddress, _, destLinkToken, err := link_token_interface.DeployLinkToken(destUser, destChain)
	require.NoError(t, err)
	destChain.Commit()
	_, err = link_token_interface.NewLinkToken(destLinkTokenAddress, destChain)
	require.NoError(t, err)

	// Deploy destination pool
	destPoolAddress, _, _, err := lock_release_token_pool.DeployLockReleaseTokenPool(destUser, destChain, destLinkTokenAddress)
	require.NoError(t, err)
	destChain.Commit()
	destPool, err := lock_release_token_pool.NewLockReleaseTokenPool(destPoolAddress, destChain)
	require.NoError(t, err)

	// Fund dest pool
	liquidityAmount := big.NewInt(1000000)
	_, err = destLinkToken.Approve(destUser, destPoolAddress, liquidityAmount)
	require.NoError(t, err)
	destChain.Commit()
	_, err = destPool.AddLiquidity(destUser, liquidityAmount)
	require.NoError(t, err)
	destChain.Commit()

	afnAddress, _, _, err := mock_afn_contract.DeployMockAFNContract(
		destUser,
		destChain,
	)
	require.NoError(t, err)
	destChain.Commit()

	destPriceRegistryAddress, _, _, err := price_registry.DeployPriceRegistry(
		destUser,
		destChain,
		price_registry.InternalPriceUpdates{
			TokenPriceUpdates: []price_registry.InternalTokenPriceUpdate{
				{
					SourceToken: destLinkTokenAddress,
					UsdPerToken: big.NewInt(8e18), // 8usd
				},
			},
			DestChainId:   destChainID,
			UsdPerUnitGas: big.NewInt(2000e9), // $2000 per eth * 1gwei = 2000e9
		},
		nil,
		[]common.Address{destLinkTokenAddress},
		60*60*24*14, // two weeks
	)
	require.NoError(t, err)

	onRampAddress := common.HexToAddress("0x01BE23585060835E02B77ef475b0Cc51aA1e0709")
	linkTokenSourceAddress := common.HexToAddress("0x01BE23585060835E02B77ef475b0Cc51aA1e0709")
	commitStoreAddress, _, _, err := commit_store_helper.DeployCommitStoreHelper(
		destUser,  // user
		destChain, // client
		commit_store_helper.CommitStoreStaticConfig{
			ChainId:       1338,
			SourceChainId: 1337,
			OnRamp:        onRampAddress,
		},
		commit_store_helper.CommitStoreDynamicConfig{
			PriceRegistry: destPriceRegistryAddress,
			Afn:           afnAddress, // AFN address
		},
	)
	require.NoError(t, err)
	commitStore, err := commit_store_helper.NewCommitStoreHelper(commitStoreAddress, destChain)
	require.NoError(t, err)

	destChain.Commit()

	routerAddress, _, routerContract, err := router.DeployRouter(destUser, destChain, common.Address{})
	require.NoError(t, err)
	destChain.Commit()
	offRampAddress, _, _, err := evm_2_evm_offramp.DeployEVM2EVMOffRamp(
		destUser,
		destChain,
		evm_2_evm_offramp.EVM2EVMOffRampStaticConfig{
			CommitStore:   commitStore.Address(),
			ChainId:       destChainID,
			SourceChainId: sourceChainID,
			OnRamp:        onRampAddress,
		},
		evm_2_evm_offramp.EVM2EVMOffRampDynamicConfig{
			Router:                                  routerAddress,
			PermissionLessExecutionThresholdSeconds: 1,
			ExecutionDelaySeconds:                   0,
			MaxDataSize:                             1e5,
			MaxTokensLength:                         5,
			Afn:                                     afnAddress,
		},
		[]common.Address{linkTokenSourceAddress},
		[]common.Address{destPoolAddress},
		evm_2_evm_offramp.AggregateRateLimiterRateLimiterConfig{
			Capacity: big.NewInt(1e18),
			Rate:     big.NewInt(1e18),
			Admin:    destUser.From,
		},
	)
	require.NoError(t, err)
	offRamp, err := evm_2_evm_offramp.NewEVM2EVMOffRamp(offRampAddress, destChain)
	require.NoError(t, err)
	_, err = destPool.ApplyRampUpdates(destUser, []lock_release_token_pool.IPoolRampUpdate{}, []lock_release_token_pool.IPoolRampUpdate{{Ramp: offRampAddress, Allowed: true}})
	require.NoError(t, err)
	receiverAddress, _, _, err := simple_message_receiver.DeploySimpleMessageReceiver(destUser, destChain)
	require.NoError(t, err)
	receiver, err := simple_message_receiver.NewSimpleMessageReceiver(receiverAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()
	_, err = routerContract.ApplyRampUpdates(destUser, nil, []router.RouterOffRampUpdate{
		{SourceChainId: sourceChainID, OffRamps: []common.Address{offRampAddress}}})
	require.NoError(t, err)
	destChain.Commit()

	require.NoError(t, err)
	destChain.Commit()
	return ExecutionContracts{user: destUser,
		offRamp:          offRamp,
		commitStore:      commitStore,
		receiver:         receiver,
		linkTokenAddress: destLinkTokenAddress,
		destChainID:      destChainID,
		sourceChainID:    sourceChainID, destChain: destChain}
}

type messageBatch struct {
	msgs        []Message
	allMsgBytes [][]byte
	seqNums     []uint64
	helperMsgs  []evm_2_evm_onramp.InternalEVM2EVMMessage
	proof       merklemulti.Proof[[32]byte]
	root        [32]byte
}

// Message contains the data from a cross chain message
type Message struct {
	SourceChainId  uint64                                  `json:"sourceChainId"`
	SequenceNumber uint64                                  `json:"sequenceNumber"`
	FeeTokenAmount *big.Int                                `json:"feeTokenAmount"`
	Sender         common.Address                          `json:"sender"`
	Nonce          uint64                                  `json:"nonce"`
	GasLimit       *big.Int                                `json:"gasLimit"`
	Strict         bool                                    `json:"strict"`
	Receiver       common.Address                          `json:"receiver"`
	Data           []uint8                                 `json:"data"`
	TokenAmounts   []evm_2_evm_onramp.ClientEVMTokenAmount `json:"tokensAndAmounts"`
	FeeToken       common.Address                          `json:"feeToken"`
	MessageId      [32]byte                                `json:"messageId"`
}

func (e ExecutionContracts) generateMessageBatch(t *testing.T, payloadSize int, nMessages int, nTokensPerMessage int) messageBatch {
	mctx := hasher.NewKeccakCtx()
	maxData := func() []byte {
		var b []byte
		for i := 0; i < payloadSize; i++ {
			b = append(b, 0xa)
		}
		return b
	}
	maxPayload := maxData()
	var leafHashes [][32]byte
	var msgs []Message
	var indices []int
	var tokens []evm_2_evm_onramp.ClientEVMTokenAmount
	var helperMsgs []evm_2_evm_onramp.InternalEVM2EVMMessage
	for i := 0; i < nTokensPerMessage; i++ {
		tokens = append(tokens, evm_2_evm_onramp.ClientEVMTokenAmount{
			Token:  e.linkTokenAddress,
			Amount: big.NewInt(1),
		})
	}
	var seqNums []uint64
	var allMsgBytes [][]byte
	for i := 0; i < nMessages; i++ {
		seqNums = append(seqNums, 1+uint64(i))
		message := Message{
			SourceChainId:  e.sourceChainID,
			SequenceNumber: 1 + uint64(i),
			FeeTokenAmount: big.NewInt(1e9),
			Sender:         e.user.From,
			Nonce:          1 + uint64(i),
			GasLimit:       big.NewInt(100_000),
			Strict:         false,
			Receiver:       e.receiver.Address(),
			Data:           maxPayload,
			TokenAmounts:   tokens,
			FeeToken:       tokens[0].Token,
			MessageId:      utils.Keccak256Fixed([]byte(`MyError(uint256)`)),
		}

		// Unfortunately have to do this to use the helper's gethwrappers.
		helperMsgs = append(helperMsgs, evm_2_evm_onramp.InternalEVM2EVMMessage{
			SourceChainId:  message.SourceChainId,
			SequenceNumber: message.SequenceNumber,
			FeeTokenAmount: message.FeeTokenAmount,
			Sender:         message.Sender,
			Nonce:          message.Nonce,
			GasLimit:       message.GasLimit,
			Strict:         message.Strict,
			Receiver:       message.Receiver,
			Data:           message.Data,
			TokenAmounts:   message.TokenAmounts,
			FeeToken:       message.FeeToken,
			MessageId:      message.MessageId,
		})

		msgs = append(msgs, message)
		indices = append(indices, i)
		msgBytes, err := ccip.MakeMessageArgs().PackValues([]interface{}{message})
		require.NoError(t, err)
		allMsgBytes = append(allMsgBytes, msgBytes)
		leafHashes = append(leafHashes, mctx.Hash(msgBytes))
	}
	tree, err := merklemulti.NewTree(mctx, leafHashes)
	require.NoError(t, err)
	proof := tree.Prove(indices)
	rootLocal, err := merklemulti.VerifyComputeRoot(mctx, leafHashes, proof)
	require.NoError(t, err)
	return messageBatch{allMsgBytes: allMsgBytes, seqNums: seqNums, msgs: msgs, proof: proof, root: rootLocal, helperMsgs: helperMsgs}
}

func TestMaxInternalExecutionReportSize(t *testing.T) {
	// Ensure that given max payload size and max num tokens,
	// Our report size is under the tx size limit.
	c := setupContractsForExecution(t)
	mb := c.generateMessageBatch(t, ccip.MaxPayloadLength, 50, ccip.MaxTokensPerMessage)
	// Ensure execution report size is valid
	executorReport, err := ccip.EncodeExecutionReport(
		mb.seqNums,
		mb.allMsgBytes,
		mb.proof.Hashes,
		mb.proof.SourceFlags,
	)
	require.NoError(t, err)
	t.Log("execution report length", len(executorReport), ccip.MaxExecutionReportLength)
	require.True(t, len(executorReport) <= ccip.MaxExecutionReportLength)

	// Check can get into mempool i.e. tx size limit is respected.
	a := c.offRamp.Address()
	bi, _ := abi.JSON(strings.NewReader(evm_2_evm_offramp_helper.EVM2EVMOffRampHelperABI))
	b, err := bi.Pack("report", []byte(executorReport))
	require.NoError(t, err)
	n, err := c.destChain.NonceAt(context.Background(), c.user.From, nil)
	require.NoError(t, err)
	signedTx, err := c.user.Signer(c.user.From, types.NewTx(&types.LegacyTx{
		To:       &a,
		Nonce:    n,
		GasPrice: big.NewInt(1e9),
		Gas:      10 * ethconfig.Defaults.Miner.GasCeil, // Massive gas limit, 10x normal block size
		Value:    big.NewInt(0),
		Data:     b,
	}))
	require.NoError(t, err)
	pool := core.NewTxPool(core.DefaultTxPoolConfig, params.AllEthashProtocolChanges, c.destChain.Blockchain())
	require.NoError(t, pool.AddLocal(signedTx))
}

func TestInternalExecutionReportEncoding(t *testing.T) {
	// Note could consider some fancier testing here (fuzz/property)
	// but I think that would essentially be testing geth's abi library
	// as our encode/decode is a thin wrapper around that.
	c := setupContractsForExecution(t)
	mb := c.generateMessageBatch(t, 1, 1, 1)
	report := evm_2_evm_offramp.InternalExecutionReport{
		SequenceNumbers: mb.seqNums,
		EncodedMessages: mb.allMsgBytes,
		Proofs:          mb.proof.Hashes,
		ProofFlagBits:   ccip.ProofFlagsToBits(mb.proof.SourceFlags),
	}
	encodeCommitReport, err := ccip.EncodeExecutionReport(report.SequenceNumbers, report.EncodedMessages, report.Proofs, mb.proof.SourceFlags)
	require.NoError(t, err)
	decodeCommitReport, err := ccip.DecodeExecutionReport(encodeCommitReport)
	require.NoError(t, err)
	require.Equal(t, &report, decodeCommitReport)
}
