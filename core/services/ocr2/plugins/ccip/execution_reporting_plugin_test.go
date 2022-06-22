package ccip_test

import (
	"context"
	"math/big"
	"strings"
	"testing"
	"time"

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

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/any_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/any_2_evm_toll_offramp_router"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/blob_verifier_helper"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/no_storage_message_receiver"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp_helper"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/merklemulti"
)

type ExecutionContracts struct {
	// Has all the link and 100ETH
	user                       *bind.TransactOpts
	offRampHelper              *offramp_helper.OffRampHelper
	blobVerifier               *blob_verifier_helper.BlobVerifierHelper
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

	onRampAddress := common.HexToAddress("0x01BE23585060835E02B77ef475b0Cc51aA1e0709")
	linkTokenSourceAddress := common.HexToAddress("0x01BE23585060835E02B77ef475b0Cc51aA1e0709")
	blobVerifierAddress, _, _, err := blob_verifier_helper.DeployBlobVerifierHelper(
		destUser,          // user
		destChain,         // client
		big.NewInt(1338),  // dest chain id
		afnAddress,        // AFN address
		big.NewInt(86400), // max timeout without AFN signal  86400 seconds = one day
		blob_verifier_helper.BlobVerifierInterfaceBlobVerifierConfig{
			SourceChainId:    big.NewInt(1337),
			OnRamps:          []common.Address{onRampAddress},
			MinSeqNrByOnRamp: []uint64{1},
		},
	)
	require.NoError(t, err)
	blobVerifier, err := blob_verifier_helper.NewBlobVerifierHelper(blobVerifierAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()
	offRampAddress, _, _, err := any_2_evm_toll_offramp.DeployAny2EVMTollOffRamp(destUser,
		destChain, destChainID, any_2_evm_toll_offramp.TollOffRampInterfaceOffRampConfig{
			SourceChainId:         sourceChainID,
			ExecutionDelaySeconds: 0,
			MaxDataSize:           1e12,
			MaxTokensLength:       5,
		},
		blobVerifier.Address(),
		onRampAddress,
		afnAddress,
		[]common.Address{linkTokenSourceAddress},
		[]common.Address{destPoolAddress},
		big.NewInt(time.Now().Unix()*2),
	)
	require.NoError(t, err)
	offRamp, err := any_2_evm_toll_offramp.NewAny2EVMTollOffRamp(offRampAddress, destChain)
	require.NoError(t, err)
	_, err = destPool.SetOffRamp(destUser, offRampAddress, true)
	require.NoError(t, err)
	receiverAddress, _, _, err := no_storage_message_receiver.DeployNoStorageMessageReceiver(destUser, destChain)
	require.NoError(t, err)
	receiver, err := no_storage_message_receiver.NewNoStorageMessageReceiver(receiverAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()
	routerAddress, _, _, err := any_2_evm_toll_offramp_router.DeployAny2EVMTollOffRampRouter(destUser, destChain, []common.Address{offRampAddress})
	require.NoError(t, err)
	destChain.Commit()
	_, err = offRamp.SetRouter(destUser, routerAddress)
	require.NoError(t, err)
	destChain.Commit()

	executorAddress, _, _, err := offramp_helper.DeployOffRampHelper(
		destUser,
		destChain,
		offRampAddress,
		false)
	require.NoError(t, err)
	executor, err := offramp_helper.NewOffRampHelper(executorAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()
	return ExecutionContracts{user: destUser,
		offRampHelper:    executor,
		blobVerifier:     blobVerifier,
		receiver:         receiver,
		linkTokenAddress: destLinkTokenAddress,
		destChainID:      destChainID,
		sourceChainID:    sourceChainID, destChain: destChain}
}

type messageBatch struct {
	msgs        []ccip.Message
	allMsgBytes [][]byte
	seqNums     []uint64
	helperMsgs  []evm_2_evm_toll_onramp.CCIPEVM2EVMTollEvent
	proof       merklemulti.Proof[[32]byte]
	root        [32]byte
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
	var helperMsgs []evm_2_evm_toll_onramp.CCIPEVM2EVMTollEvent
	for i := 0; i < nTokensPerMessage; i++ {
		tokens = append(tokens, e.linkTokenAddress)
		amounts = append(amounts, big.NewInt(1))
	}
	var seqNums []uint64
	var allMsgBytes [][]byte
	for i := 0; i < nMessages; i++ {
		seqNums = append(seqNums, 1+uint64(i))
		message := ccip.Message{
			SequenceNumber: 1 + uint64(i),
			SourceChainId:  e.sourceChainID,
			Sender:         e.user.From,
			Tokens:         tokens,
			Amounts:        amounts,
			Receiver:       e.receiver.Address(),
			Data:           maxPayload,
			FeeToken:       tokens[0],
			FeeTokenAmount: big.NewInt(4),
			GasLimit:       big.NewInt(100_000),
		}

		// Unfortunately have to do this to use the helper's gethwrappers.
		helperMsgs = append(helperMsgs, evm_2_evm_toll_onramp.CCIPEVM2EVMTollEvent{
			SequenceNumber: message.SequenceNumber,
			SourceChainId:  message.SourceChainId,
			Sender:         message.Sender,
			Tokens:         message.Tokens,
			Amounts:        message.Amounts,
			Receiver:       message.Receiver,
			Data:           message.Data,
			FeeToken:       message.FeeToken,
			FeeTokenAmount: message.FeeTokenAmount,
			GasLimit:       message.GasLimit,
		})
		msgs = append(msgs, message)
		indices = append(indices, i)
		msgBytes, err := ccip.MakeCCIPMsgArgs().PackValues([]interface{}{message})
		require.NoError(t, err)
		allMsgBytes = append(allMsgBytes, msgBytes)
		leafHashes = append(leafHashes, mctx.HashLeaf(msgBytes))
	}
	tree := merklemulti.NewTree(mctx, leafHashes)
	proof := tree.Prove(indices)
	rootLocal, err := merklemulti.VerifyComputeRoot(mctx, leafHashes, proof)
	require.NoError(t, err)
	return messageBatch{allMsgBytes: allMsgBytes, seqNums: seqNums, msgs: msgs, proof: proof, root: rootLocal, helperMsgs: helperMsgs}
}

func TestMaxExecutionReportSize(t *testing.T) {
	// Ensure that given max payload size and max num tokens,
	// Our report size is under the tx size limit.
	c := setupContractsForExecution(t)
	mb := c.generateMessageBatch(t, ccip.MaxPayloadLength, ccip.MaxNumMessagesInExecutionReport, ccip.MaxTokensPerMessage)
	ctx := merklemulti.NewKeccakCtx()
	outerTree := merklemulti.NewTree(ctx, [][32]byte{mb.root})
	outerProof := outerTree.Prove([]int{0})
	// Ensure execution report size is valid
	executorReport, err := ccip.EncodeExecutionReport(
		mb.seqNums,
		map[common.Address]uint64{},
		mb.allMsgBytes,
		mb.proof.Hashes,
		mb.proof.SourceFlags,
		outerProof.Hashes,
		outerProof.SourceFlags,
	)
	require.NoError(t, err)
	t.Log("execution report length", len(executorReport), ccip.MaxExecutionReportLength)
	require.True(t, len(executorReport) <= ccip.MaxExecutionReportLength)

	// Check can get into mempool i.e. tx size limit is respected.
	a := c.offRampHelper.Address()
	bi, _ := abi.JSON(strings.NewReader(offramp_helper.OffRampHelperABI))
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
	mb := c.generateMessageBatch(t, 1, 1, 1)
	ctx := merklemulti.NewKeccakCtx()
	outerTree := merklemulti.NewTree(ctx, [][32]byte{mb.root})
	outerProof := outerTree.Prove([]int{0})
	report := any_2_evm_toll_offramp.CCIPExecutionReport{
		SequenceNumbers:          mb.seqNums,
		TokenPerFeeCoin:          []*big.Int{},
		TokenPerFeeCoinAddresses: []common.Address{},
		EncodedMessages:          mb.allMsgBytes,
		InnerProofs:              mb.proof.Hashes,
		InnerProofFlagBits:       ccip.ProofFlagsToBits(mb.proof.SourceFlags),
		OuterProofs:              outerProof.Hashes,
		OuterProofFlagBits:       ccip.ProofFlagsToBits(outerProof.SourceFlags),
	}
	encodeRelayReport, err := ccip.EncodeExecutionReport(report.SequenceNumbers, map[common.Address]uint64{}, report.EncodedMessages, report.InnerProofs, mb.proof.SourceFlags, report.OuterProofs, outerProof.SourceFlags)
	require.NoError(t, err)
	decodeRelayReport, err := ccip.DecodeExecutionReport(encodeRelayReport)
	require.NoError(t, err)
	require.Equal(t, &report, decodeRelayReport)
}
