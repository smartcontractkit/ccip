package ccip_test

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	"github.com/test-go/testify/assert"

	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/router"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/testhelpers"
)

var (
	sourceChainID = uint64(1000)
	destChainID   = uint64(1337)
)

func TestIntegration_CCIP(t *testing.T) {
	ccipContracts := testhelpers.SetupCCIPContracts(t, sourceChainID, destChainID)
	bootstrapNodePort := int64(19399)
	ctx := context.Background()
	// Starts nodes and configures them in the OCR contracts.
	bootstrapNode, nodes, configBlock := testhelpers.SetupAndStartNodes(ctx, t, &ccipContracts, bootstrapNodePort)
	linkEth := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`{"JuelsPerETH": "200000000000000000000"}`))
		require.NoError(t, err)
	}))
	tokensPerFeeCoinPipeline := fmt.Sprintf(`
// Price 1 
link [type=http method=GET url="%s"];
link_parse [type=jsonparse path="JuelsPerETH"];
link->link_parse;
merge [type=merge left="{}" right="{\\\"%s\\\":$(link_parse)}"];`,
		linkEth.URL, ccipContracts.Dest.LinkToken.Address())
	jobParams := ccipContracts.NewCCIPJobSpecParams(tokensPerFeeCoinPipeline, configBlock)
	defer linkEth.Close()

	jobParams.RelayInflight = 2 * time.Second
	jobParams.ExecInflight = 2 * time.Second
	jobParams.RootSnooze = 1 * time.Second

	// Add the bootstrap job
	bootstrapNode.AddBootstrapJob(t, jobParams.BootstrapJob(ccipContracts.Dest.CommitStore.Address().Hex()))
	testhelpers.AddAllJobs(t, jobParams, ccipContracts, nodes)

	// Replay for bootstrap.
	bc, err := bootstrapNode.App.GetChains().EVM.Get(big.NewInt(0).SetUint64(destChainID))
	require.NoError(t, err)
	require.NoError(t, bc.LogPoller().Replay(context.Background(), configBlock))

	geCurrentSeqNum := 1

	ccipContracts.Dest.Chain.Commit()

	t.Run("single ge", func(t *testing.T) {
		tokenAmount := big.NewInt(500000003) // prime number
		gasLimit := big.NewInt(200_003)      // prime number
		gasPrice := big.NewInt(1e9)          // 1 gwei

		eventSignatures := ccip.GetEventSignatures()
		extraArgs, err := testhelpers.GetEVMExtraArgsV1(gasLimit, false)
		require.NoError(t, err)

		sourceBalances, err := testhelpers.GetBalances([]testhelpers.BalanceReq{
			{Name: testhelpers.SourcePool, Addr: ccipContracts.Source.Pool.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: testhelpers.OnRamp, Addr: ccipContracts.Source.OnRamp.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: testhelpers.SourceRouter, Addr: ccipContracts.Source.Router.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: testhelpers.SourceFeeManager, Addr: ccipContracts.Source.FeeManager.Address(), Getter: ccipContracts.GetSourceLinkBalance},
		})
		require.NoError(t, err)
		destBalances, err := testhelpers.GetBalances([]testhelpers.BalanceReq{
			{Name: testhelpers.Receiver, Addr: ccipContracts.Dest.Receivers[0].Receiver.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: testhelpers.DestPool, Addr: ccipContracts.Dest.Pool.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: testhelpers.OffRamp, Addr: ccipContracts.Dest.OffRamp.Address(), Getter: ccipContracts.GetDestLinkBalance},
		})
		require.NoError(t, err)

		msg := router.ClientEVM2AnyMessage{
			Receiver: testhelpers.MustEncodeAddress(t, ccipContracts.Dest.Receivers[0].Receiver.Address()),
			Data:     []byte("hello"),
			TokensAndAmounts: []router.CommonEVMTokenAndAmount{
				{
					Token:  ccipContracts.Source.LinkToken.Address(),
					Amount: tokenAmount,
				},
			},
			FeeToken:  ccipContracts.Source.LinkToken.Address(),
			ExtraArgs: extraArgs,
		}
		fee, err := ccipContracts.Source.Router.GetFee(nil, destChainID, msg)
		require.NoError(t, err)
		// Currently no overhead and 1gwei dest gas price. So fee is simply gasLimit * gasPrice.
		require.Equal(t, new(big.Int).Mul(gasLimit, gasPrice).String(), fee.String())
		// Approve the fee amount + the token amount
		_, err = ccipContracts.Source.LinkToken.Approve(ccipContracts.Source.User, ccipContracts.Source.Router.Address(), new(big.Int).Add(fee, tokenAmount))
		require.NoError(t, err)
		ccipContracts.Source.Chain.Commit()
		testhelpers.SendRequest(t, ccipContracts, msg)
		// Should eventually see this executed.
		testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, eventSignatures, ccipContracts.Source.OnRamp.Address(), nodes, geCurrentSeqNum)
		testhelpers.EventuallyReportCommitted(t, ccipContracts, ccipContracts.Source.OnRamp.Address(), geCurrentSeqNum)

		executionLogs := testhelpers.AllNodesHaveExecutedSeqNums(t, ccipContracts, eventSignatures, ccipContracts.Dest.OffRamp.Address(), nodes, geCurrentSeqNum, geCurrentSeqNum)
		assert.Len(t, executionLogs, 1)
		testhelpers.AssertExecState(t, ccipContracts, executionLogs[0], ccip.Success)

		// Asserts
		// 1) The total pool input == total pool output
		// 2) Pool flow equals tokens sent
		// 3) Sent tokens arrive at the receiver

		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     testhelpers.SourcePool,
				Address:  ccipContracts.Source.Pool.Address(),
				Expected: testhelpers.MustAddBigInt(sourceBalances[testhelpers.SourcePool], tokenAmount.String()).String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			},
			{
				Name:     testhelpers.SourceFeeManager,
				Address:  ccipContracts.Source.FeeManager.Address(),
				Expected: testhelpers.MustAddBigInt(sourceBalances[testhelpers.SourceFeeManager], fee.String()).String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			},
			{
				Name:     testhelpers.OnRamp,
				Address:  ccipContracts.Source.OnRamp.Address(),
				Expected: sourceBalances[testhelpers.OnRamp].String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			},
			{
				Name:     testhelpers.SourceRouter,
				Address:  ccipContracts.Source.Router.Address(),
				Expected: sourceBalances[testhelpers.SourceRouter].String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			},
		})
		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     testhelpers.Receiver,
				Address:  ccipContracts.Dest.Receivers[0].Receiver.Address(),
				Expected: testhelpers.MustAddBigInt(destBalances[testhelpers.Receiver], tokenAmount.String()).String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			},
			{
				Name:     testhelpers.DestPool,
				Address:  ccipContracts.Dest.Pool.Address(),
				Expected: testhelpers.MustSubBigInt(destBalances[testhelpers.DestPool], tokenAmount.String()).String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			},
			{
				Name:     testhelpers.OffRamp,
				Address:  ccipContracts.Dest.OffRamp.Address(),
				Expected: destBalances[testhelpers.OffRamp].String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			},
		})
		geCurrentSeqNum++
	})

	t.Run("multiple batches ge", func(t *testing.T) {
		tokenAmount := big.NewInt(500000003)
		gasLimit := big.NewInt(250_000)
		gasPrice := big.NewInt(1e9) // 1 gwei

		eventSignatures := ccip.GetEventSignatures()
		var txs []*gethtypes.Transaction
		// Enough to require batched executions as gasLimit per tx is 250k -> 500k -> 750k ....
		// The actual gas usage of executing 15 messages is higher than the gas limit for
		// a single tx. This means that when batching is turned off, and we simply include
		// all txs without checking gas, this also fails.
		n := 15
		for i := 0; i < n; i++ {
			txGasLimit := new(big.Int).Mul(gasLimit, big.NewInt(int64(i+1)))
			extraArgs, err := testhelpers.GetEVMExtraArgsV1(txGasLimit, false)
			require.NoError(t, err)
			msg := router.ClientEVM2AnyMessage{
				Receiver: testhelpers.MustEncodeAddress(t, ccipContracts.Dest.Receivers[0].Receiver.Address()),
				Data:     []byte("hello"),
				TokensAndAmounts: []router.CommonEVMTokenAndAmount{
					{
						Token:  ccipContracts.Source.LinkToken.Address(),
						Amount: tokenAmount,
					},
				},
				FeeToken:  ccipContracts.Source.LinkToken.Address(),
				ExtraArgs: extraArgs,
			}
			fee, err := ccipContracts.Source.Router.GetFee(nil, destChainID, msg)
			require.NoError(t, err)
			// Currently no overhead and 1gwei dest gas price. So fee is simply gasLimit * gasPrice.
			require.Equal(t, new(big.Int).Mul(txGasLimit, gasPrice).String(), fee.String())
			// Approve the fee amount + the token amount
			_, err = ccipContracts.Source.LinkToken.Approve(ccipContracts.Source.User, ccipContracts.Source.Router.Address(), new(big.Int).Add(fee, tokenAmount))
			require.NoError(t, err)
			tx, err := ccipContracts.Source.Router.CcipSend(ccipContracts.Source.User, ccipContracts.Dest.ChainID, msg)
			require.NoError(t, err)
			txs = append(txs, tx)
		}

		// Send a batch of requests in a single block
		testhelpers.ConfirmTxs(t, txs, ccipContracts.Source.Chain)
		var reqs []logpoller.Log
		for i := 0; i < n; i++ {
			reqs = append(reqs, testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, eventSignatures, ccipContracts.Source.OnRamp.Address(), nodes, geCurrentSeqNum+i))
		}
		// Should see a report with the full range
		testhelpers.EventuallyReportCommitted(t, ccipContracts, ccipContracts.Source.OnRamp.Address(), geCurrentSeqNum+n-1)
		// Should all be executed
		executionLogs := testhelpers.AllNodesHaveExecutedSeqNums(t, ccipContracts, eventSignatures, ccipContracts.Dest.OffRamp.Address(), nodes, geCurrentSeqNum, geCurrentSeqNum+n-1)
		for _, execLog := range executionLogs {
			testhelpers.AssertExecState(t, ccipContracts, execLog, ccip.Success)
		}

		geCurrentSeqNum += n
	})

	t.Run("ge strict sequencing", func(t *testing.T) {
		// approve the total amount to be sent
		// set revert to true so that the execution gets reverted
		_, err = ccipContracts.Dest.Receivers[1].Receiver.SetRevert(ccipContracts.Dest.User, true)
		require.NoError(t, err, "setting revert to true on the receiver")
		ccipContracts.Dest.Chain.Commit()
		currentBlockNumber := ccipContracts.Dest.Chain.Blockchain().CurrentBlock().Number().Uint64()

		// Test sequence:
		// Send msg1: strict reverts
		// Send msg2, msg3: blocked on manual exec
		// Execute msg1 manually.
		// msg2 and msg2 should go through.
		totalMsgs := 2
		extraArgs, err := testhelpers.GetEVMExtraArgsV1(big.NewInt(200_000), true)
		require.NoError(t, err)
		msg := router.ClientEVM2AnyMessage{
			Receiver:         testhelpers.MustEncodeAddress(t, ccipContracts.Dest.Receivers[1].Receiver.Address()),
			Data:             []byte("hello"),
			TokensAndAmounts: []router.CommonEVMTokenAndAmount{},
			FeeToken:         ccipContracts.Source.LinkToken.Address(),
			ExtraArgs:        extraArgs,
		}
		fee, err := ccipContracts.Source.Router.GetFee(nil, destChainID, msg)
		require.NoError(t, err)
		// Approve the fee amount
		_, err = ccipContracts.Source.LinkToken.Approve(ccipContracts.Source.User, ccipContracts.Source.Router.Address(), big.NewInt(0).Mul(big.NewInt(int64(totalMsgs)), fee))
		require.NoError(t, err)
		ccipContracts.Source.Chain.Commit()
		eventSignatures := ccip.GetEventSignatures()
		testhelpers.SendRequest(t, ccipContracts, msg)
		failedReqLog := testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, eventSignatures, ccipContracts.Source.OnRamp.Address(), nodes, geCurrentSeqNum)
		testhelpers.EventuallyReportCommitted(t, ccipContracts, ccipContracts.Source.OnRamp.Address(), geCurrentSeqNum)
		reportForFailedReq := testhelpers.EventuallyCommitReportAccepted(t, ccipContracts, currentBlockNumber)

		// execution status should be failed
		executionLogs := testhelpers.AllNodesHaveExecutedSeqNums(t, ccipContracts, eventSignatures, ccipContracts.Dest.OffRamp.Address(), nodes, geCurrentSeqNum, geCurrentSeqNum)
		assert.Len(t, executionLogs, 1)
		testhelpers.AssertExecState(t, ccipContracts, executionLogs[0], ccip.Failure)
		geCurrentSeqNum++

		// subsequent requests which should not be executed.
		var pendingReqNumbers []int
		for i := 1; i < totalMsgs; i++ {
			testhelpers.SendRequest(t, ccipContracts, msg)
			testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, eventSignatures, ccipContracts.Source.OnRamp.Address(), nodes, geCurrentSeqNum)
			testhelpers.EventuallyReportCommitted(t, ccipContracts, ccipContracts.Source.OnRamp.Address(), geCurrentSeqNum)
			executionLog := testhelpers.NoNodesHaveExecutedSeqNum(t, ccipContracts, eventSignatures, ccipContracts.Dest.OffRamp.Address(), nodes, geCurrentSeqNum)
			require.Empty(t, executionLog)
			pendingReqNumbers = append(pendingReqNumbers, geCurrentSeqNum)
			geCurrentSeqNum++
		}

		// flip the revert settings on receiver
		_, err = ccipContracts.Dest.Receivers[1].Receiver.SetRevert(ccipContracts.Dest.User, false)
		require.NoError(t, err, "setting revert to false on the receiver")
		ccipContracts.Dest.Chain.Commit()
		ccipContracts.Source.Chain.Commit()

		// manually execute the failed request
		currentBlockNumber = ccipContracts.Dest.Chain.Blockchain().CurrentBlock().Number().Uint64()
		require.NoError(t, err)
		failedSeqNum := testhelpers.ExecuteMessage(t, ccipContracts, failedReqLog, []logpoller.Log{failedReqLog}, reportForFailedReq)
		testhelpers.EventuallyExecutionStateChangedToSuccess(t, ccipContracts, []uint64{failedSeqNum}, currentBlockNumber)

		// verify all the pending requests should be successfully executed now
		for _, seqNo := range pendingReqNumbers {
			t.Logf("Verify execution for pending seqNum %d", seqNo)
			testhelpers.EventuallyExecutionStateChangedToSuccess(t, ccipContracts, []uint64{uint64(seqNo)}, 1)
		}
	})

	// Deploy new on ramp,Commit store,off ramp
	// Delete previous jobs
	// Enable new contracts
	// Create new jobs
	// Send a number of requests
	// Verify all requests after the contracts are upgraded
	t.Run("upgrade contracts and verify requests can be sent with upgraded contract", func(t *testing.T) {
		ccipContracts.DeployNewOnRamp()
		ccipContracts.DeployNewCommitStore()
		ccipContracts.DeployNewOffRamp()
		newConfigBlock := ccipContracts.Dest.Chain.Blockchain().CurrentBlock().Number().Int64()
		// delete previous jobs, 1 commit and exec
		for _, node := range nodes {
			err = node.App.DeleteJob(context.Background(), 1)
			require.NoError(t, err)
			err = node.App.DeleteJob(context.Background(), 2)
			require.NoError(t, err)
		}

		// enable the newly deployed contracts
		ccipContracts.EnableOnRamp()
		ccipContracts.EnableOffRamp()
		ccipContracts.EnableCommitStore()

		// create updated jobs
		jobParams = ccipContracts.NewCCIPJobSpecParams(tokensPerFeeCoinPipeline, newConfigBlock)
		testhelpers.AddAllJobs(t, jobParams, ccipContracts, nodes)

		startSeq := 1
		endSeqNum := 3
		eventSignatures := ccip.GetEventSignatures()
		gasLimit := big.NewInt(200_003) // prime number
		gasPrice := big.NewInt(1e9)     // 1 gwei
		tokenAmount := big.NewInt(100)
		for i := startSeq; i <= endSeqNum; i++ {
			t.Logf("sending request for seqnum %d", i)
			testhelpers.SendMessage(gasLimit, gasPrice, tokenAmount, ccipContracts.Dest.Receivers[0].Receiver.Address(), ccipContracts)
			ccipContracts.Source.Chain.Commit()
			ccipContracts.Dest.Chain.Commit()
			t.Logf("verifying seqnum %d", i)
			testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, eventSignatures, ccipContracts.Source.OnRamp.Address(), nodes, i)
			testhelpers.EventuallyReportCommitted(t, ccipContracts, ccipContracts.Source.OnRamp.Address(), i)
			executionLog := testhelpers.AllNodesHaveExecutedSeqNums(t, ccipContracts, eventSignatures, ccipContracts.Dest.OffRamp.Address(), nodes, i, i)
			testhelpers.AssertExecState(t, ccipContracts, executionLog[0], ccip.Success)
		}

		geCurrentSeqNum = endSeqNum + 1
	})
}
