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

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp_helper"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store_helper"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_ge_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/simple_message_receiver"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/merklemulti"
)

type ExecutionContracts struct {
	// Has all the link and 100ETH
	user                       *bind.TransactOpts
	offRamp                    *evm_2_evm_toll_offramp.EVM2EVMTollOffRamp
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
	destPoolAddress, _, _, err := native_token_pool.DeployNativeTokenPool(destUser, destChain, destLinkTokenAddress)
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
	commitStoreAddress, _, _, err := commit_store_helper.DeployCommitStoreHelper(
		destUser,   // user
		destChain,  // client
		1338,       // dest chain id
		1337,       // source chain id
		afnAddress, // AFN address
		commit_store_helper.ICommitStoreCommitStoreConfig{
			OnRamps:          []common.Address{onRampAddress},
			MinSeqNrByOnRamp: []uint64{1},
		},
	)
	require.NoError(t, err)
	commitStore, err := commit_store_helper.NewCommitStoreHelper(commitStoreAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()
	offRampAddress, _, _, err := evm_2_evm_toll_offramp.DeployEVM2EVMTollOffRamp(
		destUser,
		destChain,
		sourceChainID,
		destChainID,
		evm_2_evm_toll_offramp.IBaseOffRampOffRampConfig{
			ExecutionDelaySeconds: 0,
			MaxDataSize:           1e12,
			MaxTokensLength:       5,
		},
		onRampAddress,
		commitStore.Address(),
		afnAddress,
		[]common.Address{linkTokenSourceAddress},
		[]common.Address{destPoolAddress},
		evm_2_evm_toll_offramp.IAggregateRateLimiterRateLimiterConfig{
			Capacity: big.NewInt(1e18),
			Rate:     big.NewInt(1e18),
		},
		destUser.From,
	)
	require.NoError(t, err)
	offRamp, err := evm_2_evm_toll_offramp.NewEVM2EVMTollOffRamp(offRampAddress, destChain)
	require.NoError(t, err)
	_, err = destPool.SetOffRamp(destUser, offRampAddress, true)
	require.NoError(t, err)
	receiverAddress, _, _, err := simple_message_receiver.DeploySimpleMessageReceiver(destUser, destChain)
	require.NoError(t, err)
	receiver, err := simple_message_receiver.NewSimpleMessageReceiver(receiverAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()
	routerAddress, _, _, err := any_2_evm_toll_offramp_router.DeployAny2EVMTollOffRampRouter(destUser, destChain, []common.Address{offRampAddress})
	require.NoError(t, err)
	destChain.Commit()
	_, err = offRamp.SetRouter(destUser, routerAddress)
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
	helperMsgs  []evm_2_evm_toll_onramp.TollEVM2EVMTollMessage
	proof       merklemulti.Proof[[32]byte]
	root        [32]byte
}

// Message contains the data from a cross chain message
type Message struct {
	SourceChainId     uint64                                          `json:"sourceChainId"`
	SequenceNumber    uint64                                          `json:"sequenceNumber"`
	Sender            common.Address                                  `json:"sender"`
	Receiver          common.Address                                  `json:"receiver"`
	Data              []uint8                                         `json:"data"`
	TokensAndAmounts  []evm_2_evm_toll_onramp.CommonEVMTokenAndAmount `json:"tokensAndAmounts"`
	FeeTokenAndAmount evm_2_evm_toll_onramp.CommonEVMTokenAndAmount   `json:"feeTokenAndAmount"`
	GasLimit          *big.Int                                        `json:"gasLimit"`
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
	var tokens []evm_2_evm_toll_onramp.CommonEVMTokenAndAmount
	var helperMsgs []evm_2_evm_toll_onramp.TollEVM2EVMTollMessage
	for i := 0; i < nTokensPerMessage; i++ {
		tokens = append(tokens, evm_2_evm_toll_onramp.CommonEVMTokenAndAmount{
			Token:  e.linkTokenAddress,
			Amount: big.NewInt(1),
		})
	}
	var seqNums []uint64
	var allMsgBytes [][]byte
	for i := 0; i < nMessages; i++ {
		seqNums = append(seqNums, 1+uint64(i))
		message := Message{
			SequenceNumber:    1 + uint64(i),
			SourceChainId:     e.sourceChainID,
			Sender:            e.user.From,
			TokensAndAmounts:  tokens,
			Receiver:          e.receiver.Address(),
			Data:              maxPayload,
			FeeTokenAndAmount: tokens[0],
			GasLimit:          big.NewInt(100_000),
		}

		// Unfortunately have to do this to use the helper's gethwrappers.
		helperMsgs = append(helperMsgs, evm_2_evm_toll_onramp.TollEVM2EVMTollMessage{
			SequenceNumber:    message.SequenceNumber,
			SourceChainId:     message.SourceChainId,
			Sender:            message.Sender,
			TokensAndAmounts:  message.TokensAndAmounts,
			Receiver:          message.Receiver,
			Data:              message.Data,
			FeeTokenAndAmount: message.FeeTokenAndAmount,
			GasLimit:          message.GasLimit,
		})
		msgs = append(msgs, message)
		indices = append(indices, i)
		msgBytes, err := ccip.MakeTollCCIPMsgArgs().PackValues([]interface{}{message})
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

func TestMaxExecutionReportSize(t *testing.T) {
	// Ensure that given max payload size and max num tokens,
	// Our report size is under the tx size limit.
	c := setupContractsForExecution(t)
	mb := c.generateMessageBatch(t, ccip.MaxPayloadLength, 50, ccip.MaxTokensPerMessage)
	ctx := hasher.NewKeccakCtx()
	outerTree, err := merklemulti.NewTree(ctx, [][32]byte{mb.root})
	require.NoError(t, err)
	outerProof := outerTree.Prove([]int{0})
	// Ensure execution report size is valid
	executorReport, err := ccip.EncodeGEExecutionReport(
		mb.seqNums,
		map[common.Address]*big.Int{},
		mb.allMsgBytes,
		mb.proof.Hashes,
		mb.proof.SourceFlags,
		outerProof.Hashes,
		outerProof.SourceFlags,
		nil,
	)
	require.NoError(t, err)
	t.Log("execution report length", len(executorReport), ccip.MaxExecutionReportLength)
	require.True(t, len(executorReport) <= ccip.MaxExecutionReportLength)

	// Check can get into mempool i.e. tx size limit is respected.
	a := c.offRamp.Address()
	bi, _ := abi.JSON(strings.NewReader(any_2_evm_toll_offramp_helper.EVM2EVMTollOffRampHelperABI))
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
	ctx := hasher.NewKeccakCtx()
	outerTree, err := merklemulti.NewTree(ctx, [][32]byte{mb.root})
	require.NoError(t, err)
	outerProof := outerTree.Prove([]int{0})
	report := evm_2_evm_ge_offramp.GEExecutionReport{
		SequenceNumbers:          mb.seqNums,
		TokenPerFeeCoin:          []*big.Int{},
		TokenPerFeeCoinAddresses: []common.Address{},
		FeeUpdates:               []evm_2_evm_ge_offramp.GEFeeUpdate{},
		EncodedMessages:          mb.allMsgBytes,
		InnerProofs:              mb.proof.Hashes,
		InnerProofFlagBits:       ccip.ProofFlagsToBits(mb.proof.SourceFlags),
		OuterProofs:              outerProof.Hashes,
		OuterProofFlagBits:       ccip.ProofFlagsToBits(outerProof.SourceFlags),
	}
	encodeCommitReport, err := ccip.EncodeGEExecutionReport(report.SequenceNumbers, map[common.Address]*big.Int{}, report.EncodedMessages, report.InnerProofs, mb.proof.SourceFlags, report.OuterProofs, outerProof.SourceFlags, nil)
	require.NoError(t, err)
	decodeCommitReport, err := ccip.DecodeGEExecutionReport(encodeCommitReport)
	require.NoError(t, err)
	require.Equal(t, &report, decodeCommitReport)
}
