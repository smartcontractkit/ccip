package ccip_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/smartcontractkit/libocr/gethwrappers/link_token_interface"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/mock_v3_aggregator_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp_executor_helper"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp_helper"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp_router"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/simple_message_receiver"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/merklemulti"
)

func TestExecutionReportEncoding(t *testing.T) {
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	destChainID := big.NewInt(1337)
	sourceChainID := big.NewInt(1338)
	destUser, err := bind.NewKeyedTransactorWithChainID(key, destChainID)
	destChain := backends.NewSimulatedBackend(core.GenesisAlloc{
		destUser.From: {Balance: big.NewInt(0).Mul(big.NewInt(100), big.NewInt(1000000000000000000))}},
		ethconfig.Defaults.Miner.GasCeil)
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
	_, err = destLinkToken.Approve(destUser, destPoolAddress, big.NewInt(1000000))
	require.NoError(t, err)
	destChain.Commit()
	_, err = destPool.LockOrBurn(destUser, destUser.From, big.NewInt(1000000))
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
	receiverAddress, _, _, err := simple_message_receiver.DeploySimpleMessageReceiver(destUser, destChain)
	require.NoError(t, err)
	destChain.Commit()
	routerAddress, _, _, err := offramp_router.DeployOffRampRouter(destUser, destChain, []common.Address{offRampAddress})
	require.NoError(t, err)
	destChain.Commit()
	_, err = offRamp.SetRouter(destUser, routerAddress)
	require.NoError(t, err)
	destChain.Commit()

	mctx := merklemulti.NewKeccakCtx()
	var leafHashes []merklemulti.Hash
	var msgs []ccip.Message
	for i := 0; i < 3; i++ {
		message := ccip.Message{
			SequenceNumber: 10 + uint64(i),
			SourceChainId:  sourceChainID,
			Sender:         destUser.From,
			Payload: struct {
				Tokens             []common.Address `json:"tokens"`
				Amounts            []*big.Int       `json:"amounts"`
				DestinationChainId *big.Int         `json:"destinationChainId"`
				Receiver           common.Address   `json:"receiver"`
				Executor           common.Address   `json:"executor"`
				Data               []uint8          `json:"data"`
			}{
				Tokens:             []common.Address{destLinkTokenAddress},
				Amounts:            []*big.Int{big.NewInt(100)},
				DestinationChainId: destChainID,
				Receiver:           receiverAddress,
				Data:               []byte("hello"),
			},
		}
		msgs = append(msgs, message)
		msgBytes, err := ccip.MakeCCIPMsgArgs().PackValues([]interface{}{message})
		require.NoError(t, err)
		leafHashes = append(leafHashes, mctx.HashLeaf(msgBytes))
	}
	tree := merklemulti.NewTree(mctx, leafHashes)
	require.NoError(t, err)
	proof := tree.Prove([]int{0, 1, 2})

	rootLocal, err := merklemulti.VerifyComputeRoot(mctx, leafHashes, proof)
	require.NoError(t, err)
	var root [32]byte
	copy(root[:], tree.Root()[:])
	require.Equal(t, []byte(rootLocal[:]), root[:])

	report := offramp.CCIPRelayReport{
		MerkleRoot:        root,
		MinSequenceNumber: 10,
		MaxSequenceNumber: 12,
	}
	encodeRelayReport, err := ccip.EncodeRelayReport(&report)
	require.NoError(t, err)
	decodeRelayReport, err := ccip.DecodeRelayReport(encodeRelayReport)
	require.NoError(t, err)
	require.Equal(t, &report, decodeRelayReport)

	// RelayReport that Message
	_, err = offRamp.Report(destUser, encodeRelayReport)
	require.NoError(t, err)
	destChain.Commit()

	// Now execute that Message via the executor
	t.Log(offRampAddress)
	executorAddress, _, _, err := offramp_executor_helper.DeployOffRampExecutorHelper(
		destUser,
		destChain,
		offRampAddress,
		false)
	require.NoError(t, err)
	executor, err := offramp_executor_helper.NewOffRampExecutorHelper(executorAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()

	executorReport, err := ccip.EncodeExecutionReport(
		msgs,
		proof.Hashes,
		proof.SourceFlags,
	)
	require.NoError(t, err)
	er, err := ccip.DecodeExecutionReport(executorReport)
	require.NoError(t, err)

	var helperMsgs []offramp_helper.CCIPMessage
	for _, message := range msgs {
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
	}

	executionReport := offramp_helper.CCIPExecutionReport{
		Messages:       helperMsgs,
		Proofs:         er.Proofs,
		ProofFlagsBits: er.ProofFlagBits,
	}

	generatedRoot, err := offRamp.MerkleRoot(nil, executionReport)
	require.NoError(t, err)
	require.Equal(t, root, generatedRoot)
	tx, err := executor.Report(destUser, executorReport)
	require.NoError(t, err)
	destChain.Commit()
	res, err := destChain.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	assert.Equal(t, uint64(1), res.Status)
}

func TestExecutionReportInvariance(t *testing.T) {
	message := ccip.Message{
		SequenceNumber: 2e18,
		SourceChainId:  big.NewInt(9999999999999999),
		Sender:         common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB2"),
		Payload: struct {
			Tokens             []common.Address `json:"tokens"`
			Amounts            []*big.Int       `json:"amounts"`
			DestinationChainId *big.Int         `json:"destinationChainId"`
			Receiver           common.Address   `json:"receiver"`
			Executor           common.Address   `json:"executor"`
			Data               []uint8          `json:"data"`
		}{
			[]common.Address{common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB3")},
			// 1e18 * 2e9 to test values larger than int64
			[]*big.Int{big.NewInt(1e18), new(big.Int).Mul(big.NewInt(1e18), big.NewInt(2e9))},
			big.NewInt(11110),
			common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB4"),
			common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB5"),
			[]uint8{23, 255, 0, 1},
		},
	}
	var h [32]byte
	report, err := ccip.EncodeExecutionReport([]ccip.Message{message, message, message}, []merklemulti.Hash{h[:]}, []bool{true})
	require.NoError(t, err)
	er, err := ccip.DecodeExecutionReport(report)
	require.NoError(t, err)
	require.Len(t, er.Messages, 3)
	require.Equal(t, message, er.Messages[0])
	require.Equal(t, message, er.Messages[2])
}

func TestDecodeEmptyExecutionReport(t *testing.T) {
	executorReport, err := ccip.EncodeExecutionReport([]ccip.Message{}, []merklemulti.Hash{}, []bool{})
	require.NoError(t, err)
	_, err = ccip.DecodeExecutionReport(executorReport)
	require.Error(t, err)
	require.Equal(t, "assumptionViolation: expected at least one element", err.Error())
}
