package ccip

import (
	"context"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_offramp_helper"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/merklemulti"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

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

func (e ccipPluginTestHarness) generateMessageBatch(t *testing.T, payloadSize int, nMessages int, nTokensPerMessage int) messageBatch {
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
			Token:  e.feeTokenAddress,
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
			Sender:         e.owner.From,
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
		msgBytes, err := MakeMessageArgs().PackValues([]interface{}{message})
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
	c := setupCcipTestHarness(t)
	mb := c.generateMessageBatch(t, MaxPayloadLength, 50, MaxTokensPerMessage)
	// Ensure execution report size is valid
	executorReport, err := EncodeExecutionReport(
		mb.seqNums,
		mb.allMsgBytes,
		mb.proof.Hashes,
		mb.proof.SourceFlags,
	)
	require.NoError(t, err)
	t.Log("execution report length", len(executorReport), MaxExecutionReportLength)
	require.True(t, len(executorReport) <= MaxExecutionReportLength)

	// Check can get into mempool i.e. tx size limit is respected.
	a := c.offRamp.Address()
	bi, _ := abi.JSON(strings.NewReader(evm_2_evm_offramp_helper.EVM2EVMOffRampHelperABI))
	b, err := bi.Pack("report", []byte(executorReport))
	require.NoError(t, err)
	n, err := c.client.NonceAt(context.Background(), c.owner.From, nil)
	require.NoError(t, err)
	signedTx, err := c.owner.Signer(c.owner.From, types.NewTx(&types.LegacyTx{
		To:       &a,
		Nonce:    n,
		GasPrice: big.NewInt(1e9),
		Gas:      10 * ethconfig.Defaults.Miner.GasCeil, // Massive gas limit, 10x normal block size
		Value:    big.NewInt(0),
		Data:     b,
	}))
	require.NoError(t, err)
	pool := txpool.NewTxPool(txpool.DefaultConfig, params.AllEthashProtocolChanges, c.client.Blockchain())
	require.NoError(t, pool.AddLocal(signedTx))
}

func TestInternalExecutionReportEncoding(t *testing.T) {
	// Note could consider some fancier testing here (fuzz/property)
	// but I think that would essentially be testing geth's abi library
	// as our encode/decode is a thin wrapper around that.
	c := setupCcipTestHarness(t)
	mb := c.generateMessageBatch(t, 1, 1, 1)
	report := evm_2_evm_offramp.InternalExecutionReport{
		SequenceNumbers: mb.seqNums,
		EncodedMessages: mb.allMsgBytes,
		Proofs:          mb.proof.Hashes,
		ProofFlagBits:   ProofFlagsToBits(mb.proof.SourceFlags),
	}
	encodeCommitReport, err := EncodeExecutionReport(report.SequenceNumbers, report.EncodedMessages, report.Proofs, mb.proof.SourceFlags)
	require.NoError(t, err)
	decodeCommitReport, err := DecodeExecutionReport(encodeCommitReport)
	require.NoError(t, err)
	require.Equal(t, &report, decodeCommitReport)
}

func TestExecutionReportToEthTxMetadata(t *testing.T) {
	c := setupCcipTestHarness(t)
	tests := []struct {
		name     string
		msgBatch messageBatch
		err      error
	}{
		{
			"happy flow",
			c.generateMessageBatch(t, MaxPayloadLength, 50, MaxTokensPerMessage),
			nil,
		},
		{
			"invalid msgs",
			func() messageBatch {
				mb := c.generateMessageBatch(t, MaxPayloadLength, 50, MaxTokensPerMessage)
				mb.allMsgBytes[0] = []byte{1, 1, 1, 1}
				return mb
			}(),
			errors.New("abi: cannot marshal in to go type: length insufficient 4 require 32"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			encExecReport, err := EncodeExecutionReport(
				tc.msgBatch.seqNums,
				tc.msgBatch.allMsgBytes,
				tc.msgBatch.proof.Hashes,
				tc.msgBatch.proof.SourceFlags,
			)
			require.NoError(t, err)
			txMeta, err := ExecutionReportToEthTxMeta(encExecReport)
			if tc.err != nil {
				require.Equal(t, tc.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.NotNil(t, txMeta)
			require.Len(t, txMeta.MessageIDs, len(tc.msgBatch.allMsgBytes))
		})
	}
}
