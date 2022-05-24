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
	"github.com/smartcontractkit/libocr/gethwrappers/link_token_interface"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/mock_v3_aggregator_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/no_storage_message_receiver"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp_executor_helper"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp_helper"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp_router"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/merklemulti"
)

type ExecutionContracts struct {
	// Has all the link and 100ETH
	user                       *bind.TransactOpts
	executorHelper             *offramp_executor_helper.OffRampExecutorHelper
	offRampHelper              *offramp_helper.OffRampHelper
	receiver                   *no_storage_message_receiver.NoStorageMessageReceiver
	linkTokenAddress           common.Address
	destChainID, sourceChainID *big.Int
	destChain                  *backends.SimulatedBackend
}

func setupContractsForExecution(t *testing.T) ExecutionContracts {
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	destChainID := big.NewInt(1337)
	sourceChainID := big.NewInt(1338)
	destUser, err := bind.NewKeyedTransactorWithChainID(key, destChainID)
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
	destPoolAddress, _, _, err := native_token_pool.DeployNativeTokenPool(destUser, destChain, destLinkTokenAddress,
		native_token_pool.PoolInterfaceBucketConfig{
			Rate:     big.NewInt(1),
			Capacity: big.NewInt(1e9),
		}, native_token_pool.PoolInterfaceBucketConfig{
			Rate:     big.NewInt(1),
			Capacity: big.NewInt(1e9),
		})
	require.NoError(t, err)
	destChain.Commit()
	destPool, err := native_token_pool.NewNativeTokenPool(destPoolAddress, destChain)
	require.NoError(t, err)

	// Fund dest pool
	_, err = destLinkToken.Transfer(destUser, destPoolAddress, big.NewInt(1000000))
	require.NoError(t, err)
	destChain.Commit()
	_, err = destPool.LockOrBurn(destUser, big.NewInt(1000000))
	require.NoError(t, err)
	destChain.Commit()

	afnAddress, _, _, err := afn_contract.DeployAFNContract(
		destUser,
		destChain,
		[]common.Address{destUser.From},
		[]*big.Int{big.NewInt(1)},
		big.NewInt(1),
		big.NewInt(1),
	)
	require.NoError(t, err)
	destChain.Commit()

	// LINK/ETH price
	feedAddress, _, _, err := mock_v3_aggregator_contract.DeployMockV3AggregatorContract(destUser, destChain, 18, big.NewInt(6000000000000000))
	require.NoError(t, err)

	offRampAddress, _, _, err := offramp_helper.DeployOffRampHelper(
		destUser,
		destChain,
		sourceChainID,
		destChainID,
		[]common.Address{destLinkTokenAddress}, // source tokens
		[]common.Address{destPoolAddress},      // dest pool addresses
		[]common.Address{feedAddress},          // feeds
		afnAddress,                             // AFN address
		big.NewInt(86400),                      // max timeout without AFN signal  86400 seconds = one day
		0,                                      // executionDelaySeconds
		5,                                      // maxTokensLength
	)
	require.NoError(t, err)
	offRamp, err := offramp_helper.NewOffRampHelper(offRampAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()
	_, err = destPool.SetOffRamp(destUser, offRampAddress, true)
	require.NoError(t, err)
	receiverAddress, _, _, err := no_storage_message_receiver.DeployNoStorageMessageReceiver(destUser, destChain)
	require.NoError(t, err)
	receiver, err := no_storage_message_receiver.NewNoStorageMessageReceiver(receiverAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()
	routerAddress, _, _, err := offramp_router.DeployOffRampRouter(destUser, destChain, []common.Address{offRampAddress})
	require.NoError(t, err)
	destChain.Commit()
	_, err = offRamp.SetRouter(destUser, routerAddress)
	require.NoError(t, err)
	destChain.Commit()

	executorAddress, _, _, err := offramp_executor_helper.DeployOffRampExecutorHelper(
		destUser,
		destChain,
		offRampAddress,
		false)
	require.NoError(t, err)
	executor, err := offramp_executor_helper.NewOffRampExecutorHelper(executorAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()
	return ExecutionContracts{user: destUser,
		executorHelper:   executor,
		offRampHelper:    offRamp,
		receiver:         receiver,
		linkTokenAddress: destLinkTokenAddress,
		destChainID:      destChainID,
		sourceChainID:    sourceChainID, destChain: destChain}
}

type messageBatch struct {
	msgs       []ccip.Message
	helperMsgs []offramp_helper.CCIPMessage
	proof      merklemulti.Proof[[32]byte]
	root       [32]byte
}

func (e ExecutionContracts) generateMessageBatch(t *testing.T, payloadSize int, nMessages int, nTokensPerMessage int) messageBatch {
	mctx := merklemulti.NewKeccakCtx()
	maxData := func() []byte {
		var b []byte
		for i := 0; i < payloadSize; i++ {
			b = append(b, 0xa)
		}
		return b
	}
	maxPayload := maxData()
	var leafHashes [][32]byte
	var msgs []ccip.Message
	var indices []int
	var tokens []common.Address
	var amounts []*big.Int
	var helperMsgs []offramp_helper.CCIPMessage
	for i := 0; i < nTokensPerMessage; i++ {
		tokens = append(tokens, e.linkTokenAddress)
		amounts = append(amounts, big.NewInt(1))
	}
	for i := 0; i < nMessages; i++ {
		message := ccip.Message{
			SequenceNumber: 1 + uint64(i),
			SourceChainId:  e.sourceChainID,
			Sender:         e.user.From,
			Payload: struct {
				Tokens             []common.Address `json:"tokens"`
				Amounts            []*big.Int       `json:"amounts"`
				DestinationChainId *big.Int         `json:"destinationChainId"`
				Receiver           common.Address   `json:"receiver"`
				Executor           common.Address   `json:"executor"`
				Data               []uint8          `json:"data"`
			}{
				Tokens:             tokens,
				Amounts:            amounts,
				DestinationChainId: e.destChainID,
				Receiver:           e.receiver.Address(),
				Data:               maxPayload,
			},
		}
		// Unfortunately have to do this to use the helper's gethwrappers.
		helperMsgs = append(helperMsgs, offramp_helper.CCIPMessage{
			SequenceNumber: message.SequenceNumber,
			SourceChainId:  message.SourceChainId,
			Sender:         message.Sender,
			Payload: offramp_helper.CCIPMessagePayload{
				Tokens:             message.Payload.Tokens,
				Amounts:            message.Payload.Amounts,
				DestinationChainId: message.Payload.DestinationChainId,
				Receiver:           message.Payload.Receiver,
				Executor:           message.Payload.Executor,
				Data:               message.Payload.Data,
			},
		})
		msgs = append(msgs, message)
		indices = append(indices, i)
		msgBytes, err := ccip.MakeCCIPMsgArgs().PackValues([]interface{}{message})
		require.NoError(t, err)
		leafHashes = append(leafHashes, mctx.HashLeaf(msgBytes))
	}
	tree := merklemulti.NewTree(mctx, leafHashes)
	proof := tree.Prove(indices)
	rootLocal, err := merklemulti.VerifyComputeRoot(mctx, leafHashes, proof)
	require.NoError(t, err)
	return messageBatch{msgs: msgs, proof: proof, root: rootLocal, helperMsgs: helperMsgs}
}

func TestMaxExecutionReportSize(t *testing.T) {
	// Ensure that given max payload size and max num tokens,
	// Our report size is under the tx size limit.
	c := setupContractsForExecution(t)
	mb := c.generateMessageBatch(t, ccip.MaxPayloadLength, ccip.MaxNumMessagesInExecutionReport, ccip.MaxTokensPerMessage)
	// Ensure execution report size is valid
	executorReport, err := ccip.EncodeExecutionReport(
		mb.msgs,
		mb.proof.Hashes,
		mb.proof.SourceFlags,
	)
	require.NoError(t, err)
	t.Log("execution report length", len(executorReport), ccip.MaxExecutionReportLength)
	require.True(t, len(executorReport) <= ccip.MaxExecutionReportLength)

	// Check can get into mempool i.e. tx size limit is respected.
	a := c.executorHelper.Address()
	bi, _ := abi.JSON(strings.NewReader(offramp_executor_helper.OffRampExecutorHelperABI))
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

func TestExecutionReportEncoding(t *testing.T) {
	// Note could consider some fancier testing here (fuzz/property)
	// but I think that would essentially be testing geth's abi library
	// as our encode/decode is a thin wrapper around that.
	c := setupContractsForExecution(t)
	mb := c.generateMessageBatch(t, ccip.MaxPayloadLength, 3, ccip.MaxTokensPerMessage)
	report := ccip.ExecutionReport{
		Messages:      mb.msgs,
		Proofs:        mb.proof.Hashes,
		ProofFlagBits: ccip.ProofFlagsToBits(mb.proof.SourceFlags),
	}
	encodeRelayReport, err := ccip.EncodeExecutionReport(mb.msgs, mb.proof.Hashes, mb.proof.SourceFlags)
	require.NoError(t, err)
	decodeRelayReport, err := ccip.DecodeExecutionReport(encodeRelayReport)
	require.NoError(t, err)
	require.Equal(t, &report, decodeRelayReport)

	executorReport, err := ccip.EncodeExecutionReport([]ccip.Message{}, [][32]byte{}, []bool{})
	require.NoError(t, err)
	_, err = ccip.DecodeExecutionReport(executorReport)
	require.Error(t, err)
	require.Equal(t, "assumptionViolation: expected at least one element", err.Error())
}
