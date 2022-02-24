package ccip_test

import (
	"bytes"
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/lib/pq"
	"github.com/smartcontractkit/libocr/gethwrappers/link_token_interface"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/internal/cltest/heavyweight"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/lock_unlock_pool"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/message_executor_helper"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/simple_message_receiver"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp_helper"
	"github.com/smartcontractkit/chainlink/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ccip/abihelpers"
	lastreportermocks "github.com/smartcontractkit/chainlink/core/services/ccip/mocks/lastreporter"
	"github.com/smartcontractkit/chainlink/core/utils"
)

func TestExecutionReportEncoding(t *testing.T) {
	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	destUser, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
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
	destPoolAddress, _, _, err := lock_unlock_pool.DeployLockUnlockPool(destUser, destChain, destLinkTokenAddress)
	require.NoError(t, err)
	destChain.Commit()
	destPool, err := lock_unlock_pool.NewLockUnlockPool(destPoolAddress, destChain)
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

	offRampAddress, _, _, err := single_token_offramp_helper.DeploySingleTokenOffRampHelper(
		destUser,
		destChain,
		big.NewInt(1337),
		big.NewInt(1338),
		destLinkTokenAddress,
		destPoolAddress,
		big.NewInt(1),    // token bucket rate
		big.NewInt(1000), // token bucket capacity
		afnAddress,       // AFN address
		// 86400 seconds = one day
		big.NewInt(86400), // max timeout without AFN signal
		big.NewInt(0),     // execution delay in seconds
	)
	require.NoError(t, err)
	offRamp, err := single_token_offramp_helper.NewSingleTokenOffRampHelper(offRampAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()
	_, err = destPool.SetOffRamp(destUser, offRampAddress, true)
	require.NoError(t, err)
	receiverAddress, _, _, err := simple_message_receiver.DeploySimpleMessageReceiver(destUser, destChain)
	require.NoError(t, err)
	destChain.Commit()

	m := ccip.Request{
		SeqNum:        *utils.NewBigI(10),
		SourceChainID: "1337",
		DestChainID:   "1338",
		Sender:        destUser.From,
		Receiver:      receiverAddress,
		Data:          []byte("hello"),
		Tokens:        []string{destLinkTokenAddress.String()},
		Amounts:       []string{"100"},
		Options:       []byte{},
	}
	msgBytes, err := abihelpers.MakeCCIPMsgArgs().PackValues([]interface{}{m.ToMessage()})
	require.NoError(t, err)
	r, proof := ccip.GenerateMerkleProof(2, [][]byte{msgBytes}, 0)
	var root [32]byte
	copy(root[:], r[:])
	rootLocal := ccip.GenerateMerkleRoot(msgBytes, proof)
	require.True(t, bytes.Equal(rootLocal[:], r[:]))

	out, err := ccip.EncodeRelayReport(&single_token_offramp.CCIPRelayReport{
		MerkleRoot:        root,
		MinSequenceNumber: big.NewInt(10),
		MaxSequenceNumber: big.NewInt(10),
	})
	require.NoError(t, err)
	_, err = ccip.DecodeRelayReport(out)
	require.NoError(t, err)

	// RelayReport that Message
	tx, err := offRamp.Report(destUser, out)
	require.NoError(t, err)
	destChain.Commit()

	// Now execute that Message via the executor
	t.Log(offRampAddress)
	executorAddress, _, _, err := message_executor_helper.DeployMessageExecutorHelper(
		destUser,
		destChain,
		offRampAddress)
	require.NoError(t, err)
	executor, err := message_executor_helper.NewMessageExecutorHelper(executorAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()

	executorReport, err := ccip.EncodeExecutionReport([]ccip.ExecutableMessage{
		{
			Proof:   proof.PathForExecute(),
			Message: m.ToMessage(),
			Index:   proof.Index(),
		},
	})
	require.NoError(t, err)
	ems, err := ccip.DecodeExecutionReport(executorReport)
	require.NoError(t, err)
	t.Log(ems)

	generatedRoot, err := offRamp.GenerateMerkleRoot(nil, proof.PathForExecute(), ccip.HashLeaf(msgBytes), proof.Index())
	require.NoError(t, err)
	require.Equal(t, root, generatedRoot)
	tx, err = executor.Report(destUser, executorReport)
	require.NoError(t, err)
	destChain.Commit()
	res, err := destChain.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	assert.Equal(t, uint64(1), res.Status)
}

func TestExecutionPlugin(t *testing.T) {
	// Avoid using txdb: it has bugs and currently has savepoints disabled (to be able to use with gorm)
	// and so any ctx cancellation poisons the tx.
	_, db := heavyweight.FullTestDB(t, "executor_plugin", true, false)
	lggr := logger.TestLogger(t)
	orm := ccip.NewORM(db, lggr, pgtest.NewPGCfg(false))
	lr := new(lastreportermocks.OffRampLastReporter)
	executor := common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB5")
	rf := ccip.NewExecutionReportingPluginFactory(logger.TestLogger(t), orm, big.NewInt(1), big.NewInt(2), executor, lr)
	rp, _, err := rf.NewReportingPlugin(types.ReportingPluginConfig{F: 1})
	require.NoError(t, err)
	sid, did := big.NewInt(1), big.NewInt(2)
	// Observe with nothing in the db should error with no observations
	obs, err := rp.Observation(context.Background(), types.ReportTimestamp{}, types.Query{})
	require.Error(t, err)
	require.Len(t, obs, 0)

	// Observe with a non-relay-confirmed request should still return no requests
	req := ccip.Request{
		SeqNum:        *utils.NewBigI(2),
		SourceChainID: sid.String(),
		DestChainID:   did.String(),
		Sender:        common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB4"),
		Data:          []byte("hello"),
		Tokens:        pq.StringArray{},
		Amounts:       pq.StringArray{},
		Executor:      executor,
		Options:       []byte{},
	}
	b, err := abihelpers.MakeCCIPMsgArgs().PackValues([]interface{}{req.ToMessage()})
	require.NoError(t, err)
	req.Raw = b
	require.NoError(t, orm.SaveRequest(&req))
	obs, err = rp.Observation(context.Background(), types.ReportTimestamp{}, types.Query{})
	require.Error(t, err)
	require.Equal(t, "no requests for oracle execution", err.Error())
	require.Len(t, obs, 0)

	// We should see an error if the latest report doesn't have a higher seq num
	lr.On("GetLastReport", mock.Anything).Return(getLastReportMock(1)).Once()
	require.NoError(t, orm.UpdateRequestSetStatus(sid, did, []*big.Int{big.NewInt(2)}, ccip.RequestStatusRelayConfirmed))
	obs, err = rp.Observation(context.Background(), types.ReportTimestamp{}, types.Query{})
	require.Error(t, err)
	// Should succeed if we do have a higher seq num
	lr.On("GetLastReport", mock.Anything).Return(getLastReportMock(2)).Once()
	obs, err = rp.Observation(context.Background(), types.ReportTimestamp{}, types.Query{})
	require.NoError(t, err)
	var o ccip.Observation
	require.NoError(t, json.Unmarshal(obs, &o))
	require.Equal(t, "2", o.MinSeqNum.String())
	require.Equal(t, "2", o.MaxSeqNum.String())

	// If all the nodes report the same, this should succeed
	// First add the relay report
	root, _ := ccip.GenerateMerkleProof(32, [][]byte{b}, 0)
	require.NoError(t, orm.SaveRelayReport(ccip.RelayReport{Root: root[:], MinSeqNum: *utils.NewBigI(2), MaxSeqNum: *utils.NewBigI(2)}))
	lr.On("GetLastReport", mock.Anything).Return(getLastReportMock(2)).Once()
	finalizeReport, rep, err := rp.Report(context.Background(), types.ReportTimestamp{}, types.Query{}, []types.AttributedObservation{
		{Observation: obs}, {Observation: obs}, {Observation: obs}, {Observation: obs},
	})
	require.NoError(t, err)
	require.True(t, finalizeReport)
	executableMessages, err := ccip.DecodeExecutionReport(rep)
	require.NoError(t, err)
	// Should see our one message there
	require.Len(t, executableMessages, 1)
	require.Equal(t, "2", executableMessages[0].Message.SequenceNumber.String())

	// If we have < F observations, we should not get a report
	finalizeReport, rep, err = rp.Report(context.Background(), types.ReportTimestamp{}, types.Query{}, []types.AttributedObservation{
		{Observation: nil}, {Observation: nil}, {Observation: nil}, {Observation: obs},
	})
	require.False(t, finalizeReport)
	// With F=1, that means a single value cannot corrupt our report
	var fakeObs = ccip.Observation{
		MinSeqNum: *utils.NewBigI(10000),
		MaxSeqNum: *utils.NewBigI(10000),
	}
	b, err = json.Marshal(fakeObs)
	require.NoError(t, err)
	lr.On("GetLastReport", mock.Anything).Return(getLastReportMock(2)).Once()
	finalizeReport, rep, err = rp.Report(context.Background(), types.ReportTimestamp{}, types.Query{}, []types.AttributedObservation{
		{Observation: obs}, {Observation: obs}, {Observation: obs}, {Observation: b},
	})
	require.NoError(t, err)
	// Still our message 2 despite the fakeObs
	executableMessages, err = ccip.DecodeExecutionReport(rep)
	require.NoError(t, err)
	require.Len(t, executableMessages, 1)
	require.Equal(t, "2", executableMessages[0].Message.SequenceNumber.String())

	// Should not accept or transmit if the report is stale
	orm.UpdateRequestSetStatus(sid, did, []*big.Int{big.NewInt(2)}, ccip.RequestStatusExecutionConfirmed)
	accept, err := rp.ShouldAcceptFinalizedReport(context.Background(), types.ReportTimestamp{}, rep)
	require.NoError(t, err)
	require.False(t, accept)
	accept, err = rp.ShouldTransmitAcceptedReport(context.Background(), types.ReportTimestamp{}, rep)
	require.NoError(t, err)
	require.False(t, accept)

	// Ensure observing and reporting works with batches.
	// Let's save a batch of seqnums {3,4,5}
	var leaves [][]byte
	for i := 3; i < 6; i++ {
		req := ccip.Request{
			SeqNum:        *utils.NewBigI(int64(i)),
			SourceChainID: sid.String(),
			DestChainID:   did.String(),
			Sender:        common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB4"),
			Data:          []byte("hello"),
			Tokens:        pq.StringArray{},
			Amounts:       pq.StringArray{},
			Executor:      executor,
			Options:       []byte{},
		}
		b, err := abihelpers.MakeCCIPMsgArgs().PackValues([]interface{}{req.ToMessage()})
		require.NoError(t, err)
		req.Raw = b
		require.NoError(t, orm.SaveRequest(&req))
		leaves = append(leaves, b)
	}
	require.NoError(t, orm.UpdateRequestStatus(sid, did, big.NewInt(3), big.NewInt(5), ccip.RequestStatusRelayConfirmed))
	lr.On("GetLastReport", mock.Anything).Return(getLastReportMock(5)).Once()
	obs, err = rp.Observation(context.Background(), types.ReportTimestamp{}, types.Query{})
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(obs, &o))
	require.Equal(t, "3", o.MinSeqNum.String())
	require.Equal(t, "5", o.MaxSeqNum.String())

	// Let's put 2 in one report and 1 in a different report then assert the execution report makes sense
	root1, _ := ccip.GenerateMerkleProof(32, [][]byte{leaves[0]}, 0)
	require.NoError(t, orm.SaveRelayReport(ccip.RelayReport{Root: root1[:], MinSeqNum: *utils.NewBigI(3), MaxSeqNum: *utils.NewBigI(3)}))
	root2, _ := ccip.GenerateMerkleProof(32, [][]byte{leaves[1], leaves[2]}, 0)
	require.NoError(t, orm.SaveRelayReport(ccip.RelayReport{Root: root2[:], MinSeqNum: *utils.NewBigI(4), MaxSeqNum: *utils.NewBigI(5)}))
	lr.On("GetLastReport", mock.Anything).Return(getLastReportMock(5)).Once()
	finalizeReport, rep, err = rp.Report(context.Background(), types.ReportTimestamp{}, types.Query{}, []types.AttributedObservation{
		{Observation: obs}, {Observation: obs}, {Observation: obs}, {Observation: obs},
	})
	require.NoError(t, err)
	msgs, err := ccip.DecodeExecutionReport(rep)
	require.NoError(t, err)
	require.Len(t, msgs, 3)
	rootLeaf1 := ccip.GenerateMerkleRoot(leaves[0], ccip.NewMerkleProof(int(msgs[0].Index.Int64()), msgs[0].Proof))
	rootLeaf2 := ccip.GenerateMerkleRoot(leaves[1], ccip.NewMerkleProof(int(msgs[1].Index.Int64()), msgs[1].Proof))
	rootLeaf3 := ccip.GenerateMerkleRoot(leaves[1], ccip.NewMerkleProof(int(msgs[1].Index.Int64()), msgs[1].Proof))
	require.True(t, bytes.Equal(rootLeaf1[:], root1[:]))
	require.True(t, bytes.Equal(rootLeaf2[:], root2[:]))
	require.True(t, bytes.Equal(rootLeaf3[:], root2[:]))
}

func getLastReportMock(maxSequenceNumber int64) (single_token_offramp.CCIPRelayReport, error) {
	maxSequenceNumberBig := big.NewInt(maxSequenceNumber)
	return single_token_offramp.CCIPRelayReport{
		MaxSequenceNumber: maxSequenceNumberBig,
	}, nil
}
