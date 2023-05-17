package ccip_test

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	"github.com/test-go/testify/assert"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/testhelpers"
	integrationtesthelpers "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/testhelpers/integration"
)

func TestIntegration_CCIP(t *testing.T) {
	ccipTH := integrationtesthelpers.SetupCCIPIntegrationTH(t, testhelpers.SourceChainID, testhelpers.DestChainID)
	linkUSD := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte(`{"UsdPerLink": "8000000000000000000"}`))
		require.NoError(t, err)
	}))
	ethUSD := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte(`{"UsdPerETH": "1700000000000000000000"}`))
		require.NoError(t, err)
	}))
	wrapped, err := ccipTH.Source.Router.GetWrappedNative(nil)
	require.NoError(t, err)
	tokenPricesUSDPipeline := fmt.Sprintf(`
// Price 1
link [type=http method=GET url="%s"];
link_parse [type=jsonparse path="UsdPerLink"];
link->link_parse;
eth [type=http method=GET url="%s"];
eth_parse [type=jsonparse path="UsdPerETH"];
eth->eth_parse;
merge [type=merge left="{}" right="{\\\"%s\\\":$(link_parse), \\\"%s\\\":$(eth_parse)}"];`,
		linkUSD.URL, ethUSD.URL, ccipTH.Dest.LinkToken.Address(), wrapped)
	defer linkUSD.Close()
	defer ethUSD.Close()

	jobParams := ccipTH.SetUpNodesAndJobs(t, tokenPricesUSDPipeline, 19399)

	geCurrentSeqNum := 1

	t.Run("single ge", func(t *testing.T) {
		tokenAmount := big.NewInt(500000003) // prime number
		gasLimit := big.NewInt(200_003)      // prime number
		// gasPrice := big.NewInt(1e9)          // 1 gwei

		extraArgs, err2 := testhelpers.GetEVMExtraArgsV1(gasLimit, false)
		require.NoError(t, err2)

		sourceBalances, err2 := testhelpers.GetBalances(t, []testhelpers.BalanceReq{
			{Name: testhelpers.SourcePool, Addr: ccipTH.Source.Pool.Address(), Getter: ccipTH.GetSourceLinkBalance},
			{Name: testhelpers.OnRamp, Addr: ccipTH.Source.OnRamp.Address(), Getter: ccipTH.GetSourceLinkBalance},
			{Name: testhelpers.SourceRouter, Addr: ccipTH.Source.Router.Address(), Getter: ccipTH.GetSourceLinkBalance},
			{Name: testhelpers.SourcePrices, Addr: ccipTH.Source.PriceRegistry.Address(), Getter: ccipTH.GetSourceLinkBalance},
		})
		require.NoError(t, err2)
		destBalances, err2 := testhelpers.GetBalances(t, []testhelpers.BalanceReq{
			{Name: testhelpers.Receiver, Addr: ccipTH.Dest.Receivers[0].Receiver.Address(), Getter: ccipTH.GetDestLinkBalance},
			{Name: testhelpers.DestPool, Addr: ccipTH.Dest.Pool.Address(), Getter: ccipTH.GetDestLinkBalance},
			{Name: testhelpers.OffRamp, Addr: ccipTH.Dest.OffRamp.Address(), Getter: ccipTH.GetDestLinkBalance},
		})
		require.NoError(t, err2)

		msg := router.ClientEVM2AnyMessage{
			Receiver: testhelpers.MustEncodeAddress(t, ccipTH.Dest.Receivers[0].Receiver.Address()),
			Data:     []byte("hello"),
			TokenAmounts: []router.ClientEVMTokenAmount{
				{
					Token:  ccipTH.Source.LinkToken.Address(),
					Amount: tokenAmount,
				},
			},
			FeeToken:  ccipTH.Source.LinkToken.Address(),
			ExtraArgs: extraArgs,
		}
		fee, err2 := ccipTH.Source.Router.GetFee(nil, testhelpers.DestChainID, msg)
		require.NoError(t, err2)
		// Currently no overhead and 10gwei dest gas price. So fee is simply (gasLimit * gasPrice)* link/native
		// require.Equal(t, new(big.Int).Mul(gasLimit, gasPrice).String(), fee.String())
		// Approve the fee amount + the token amount
		_, err2 = ccipTH.Source.LinkToken.Approve(ccipTH.Source.User, ccipTH.Source.Router.Address(), new(big.Int).Add(fee, tokenAmount))
		require.NoError(t, err2)
		ccipTH.Source.Chain.Commit()
		ccipTH.SendRequest(t, msg)
		// Should eventually see this executed.
		ccipTH.AllNodesHaveReqSeqNum(t, geCurrentSeqNum)
		ccipTH.EventuallyReportCommitted(t, ccipTH.Source.OnRamp.Address(), geCurrentSeqNum)

		executionLogs := ccipTH.AllNodesHaveExecutedSeqNums(t, geCurrentSeqNum, geCurrentSeqNum)
		assert.Len(t, executionLogs, 1)
		ccipTH.AssertExecState(t, executionLogs[0], abihelpers.ExecutionStateSuccess)

		// Asserts
		// 1) The total pool input == total pool output
		// 2) Pool flow equals tokens sent
		// 3) Sent tokens arrive at the receiver

		ccipTH.AssertBalances(t, []testhelpers.BalanceAssertion{
			{
				Name:     testhelpers.SourcePool,
				Address:  ccipTH.Source.Pool.Address(),
				Expected: testhelpers.MustAddBigInt(sourceBalances[testhelpers.SourcePool], tokenAmount.String()).String(),
				Getter:   ccipTH.GetSourceLinkBalance,
			},
			{
				Name:     testhelpers.SourcePrices,
				Address:  ccipTH.Source.PriceRegistry.Address(),
				Expected: sourceBalances[testhelpers.SourcePrices].String(),
				Getter:   ccipTH.GetSourceLinkBalance,
			},
			{
				// Fees end up in the onramp.
				Name:     testhelpers.OnRamp,
				Address:  ccipTH.Source.OnRamp.Address(),
				Expected: testhelpers.MustAddBigInt(sourceBalances[testhelpers.SourcePrices], fee.String()).String(),
				Getter:   ccipTH.GetSourceLinkBalance,
			},
			{
				Name:     testhelpers.SourceRouter,
				Address:  ccipTH.Source.Router.Address(),
				Expected: sourceBalances[testhelpers.SourceRouter].String(),
				Getter:   ccipTH.GetSourceLinkBalance,
			},
		})
		ccipTH.AssertBalances(t, []testhelpers.BalanceAssertion{
			{
				Name:     testhelpers.Receiver,
				Address:  ccipTH.Dest.Receivers[0].Receiver.Address(),
				Expected: testhelpers.MustAddBigInt(destBalances[testhelpers.Receiver], tokenAmount.String()).String(),
				Getter:   ccipTH.GetDestLinkBalance,
			},
			{
				Name:     testhelpers.DestPool,
				Address:  ccipTH.Dest.Pool.Address(),
				Expected: testhelpers.MustSubBigInt(destBalances[testhelpers.DestPool], tokenAmount.String()).String(),
				Getter:   ccipTH.GetDestLinkBalance,
			},
			{
				Name:     testhelpers.OffRamp,
				Address:  ccipTH.Dest.OffRamp.Address(),
				Expected: destBalances[testhelpers.OffRamp].String(),
				Getter:   ccipTH.GetDestLinkBalance,
			},
		})
		geCurrentSeqNum++
	})

	t.Run("multiple batches ge", func(t *testing.T) {
		tokenAmount := big.NewInt(500000003)
		gasLimit := big.NewInt(250_000)
		// gasPrice := big.NewInt(1e9) // 1 gwei

		var txs []*gethtypes.Transaction
		// Enough to require batched executions as gasLimit per tx is 250k -> 500k -> 750k ....
		// The actual gas usage of executing 15 messages is higher than the gas limit for
		// a single tx. This means that when batching is turned off, and we simply include
		// all txs without checking gas, this also fails.
		n := 15
		for i := 0; i < n; i++ {
			txGasLimit := new(big.Int).Mul(gasLimit, big.NewInt(int64(i+1)))
			extraArgs, err2 := testhelpers.GetEVMExtraArgsV1(txGasLimit, false)
			require.NoError(t, err2)
			msg := router.ClientEVM2AnyMessage{
				Receiver: testhelpers.MustEncodeAddress(t, ccipTH.Dest.Receivers[0].Receiver.Address()),
				Data:     []byte("hello"),
				TokenAmounts: []router.ClientEVMTokenAmount{
					{
						Token:  ccipTH.Source.LinkToken.Address(),
						Amount: tokenAmount,
					},
				},
				FeeToken:  ccipTH.Source.LinkToken.Address(),
				ExtraArgs: extraArgs,
			}
			fee, err2 := ccipTH.Source.Router.GetFee(nil, testhelpers.DestChainID, msg)
			require.NoError(t, err2)
			// Currently no overhead and 1gwei dest gas price. So fee is simply gasLimit * gasPrice.
			// require.Equal(t, new(big.Int).Mul(txGasLimit, gasPrice).String(), fee.String())
			// Approve the fee amount + the token amount
			_, err2 = ccipTH.Source.LinkToken.Approve(ccipTH.Source.User, ccipTH.Source.Router.Address(), new(big.Int).Add(fee, tokenAmount))
			require.NoError(t, err2)
			tx, err2 := ccipTH.Source.Router.CcipSend(ccipTH.Source.User, ccipTH.Dest.ChainID, msg)
			require.NoError(t, err2)
			txs = append(txs, tx)
		}

		// Send a batch of requests in a single block
		testhelpers.ConfirmTxs(t, txs, ccipTH.Source.Chain)
		for i := 0; i < n; i++ {
			ccipTH.AllNodesHaveReqSeqNum(t, geCurrentSeqNum+i)
		}
		// Should see a report with the full range
		ccipTH.EventuallyReportCommitted(t, ccipTH.Source.OnRamp.Address(), geCurrentSeqNum+n-1)
		// Should all be executed
		executionLogs := ccipTH.AllNodesHaveExecutedSeqNums(t, geCurrentSeqNum, geCurrentSeqNum+n-1)
		for _, execLog := range executionLogs {
			ccipTH.AssertExecState(t, execLog, abihelpers.ExecutionStateSuccess)
		}

		geCurrentSeqNum += n
	})

	t.Run("ge strict sequencing", func(t *testing.T) {
		// approve the total amount to be sent
		// set revert to true so that the execution gets reverted
		_, err = ccipTH.Dest.Receivers[1].Receiver.SetRevert(ccipTH.Dest.User, true)
		require.NoError(t, err, "setting revert to true on the receiver")
		ccipTH.Dest.Chain.Commit()
		currentBlockNumber := ccipTH.Dest.Chain.Blockchain().CurrentBlock().Number.Uint64()

		// Test sequence:
		// Send msg1: strict reverts
		// Send msg2, msg3: blocked on manual exec
		// Execute msg1 manually.
		// msg2 and msg2 should go through.
		totalMsgs := 2
		extraArgs, err2 := testhelpers.GetEVMExtraArgsV1(big.NewInt(200_000), true)
		require.NoError(t, err2)
		startNonce, err2 := ccipTH.Dest.OffRamp.GetSenderNonce(nil, ccipTH.Source.User.From)
		require.NoError(t, err2)
		msg := router.ClientEVM2AnyMessage{
			Receiver:     testhelpers.MustEncodeAddress(t, ccipTH.Dest.Receivers[1].Receiver.Address()),
			Data:         []byte("hello"),
			TokenAmounts: []router.ClientEVMTokenAmount{},
			FeeToken:     ccipTH.Source.LinkToken.Address(),
			ExtraArgs:    extraArgs,
		}
		fee, err2 := ccipTH.Source.Router.GetFee(nil, testhelpers.DestChainID, msg)
		require.NoError(t, err2)
		// Approve the fee amount
		_, err2 = ccipTH.Source.LinkToken.Approve(ccipTH.Source.User, ccipTH.Source.Router.Address(), big.NewInt(0).Mul(big.NewInt(int64(totalMsgs)), fee))
		require.NoError(t, err2)
		ccipTH.Source.Chain.Commit()
		txForFailedReq := ccipTH.SendRequest(t, msg)
		failedReqLog := ccipTH.AllNodesHaveReqSeqNum(t, geCurrentSeqNum)
		ccipTH.EventuallyReportCommitted(t, ccipTH.Source.OnRamp.Address(), geCurrentSeqNum)
		ccipTH.EventuallyCommitReportAccepted(t, currentBlockNumber)

		// execution status should be failed
		executionLogs := ccipTH.AllNodesHaveExecutedSeqNums(t, geCurrentSeqNum, geCurrentSeqNum)
		assert.Len(t, executionLogs, 1)
		ccipTH.AssertExecState(t, executionLogs[0], abihelpers.ExecutionStateFailure)
		// Nonce should not have incremented
		afterNonce, err2 := ccipTH.Dest.OffRamp.GetSenderNonce(nil, ccipTH.Source.User.From)
		require.NoError(t, err2)
		require.Equal(t, startNonce, afterNonce)
		geCurrentSeqNum++

		// flip the revert settings on receiver
		_, err2 = ccipTH.Dest.Receivers[1].Receiver.SetRevert(ccipTH.Dest.User, false)
		require.NoError(t, err2, "setting revert to false on the receiver")
		ccipTH.Dest.Chain.Commit()
		ccipTH.Source.Chain.Commit()

		// subsequent requests which should not be executed.
		var pendingReqNumbers []int
		for i := 1; i < totalMsgs; i++ {
			ccipTH.SendRequest(t, msg)
			ccipTH.AllNodesHaveReqSeqNum(t, geCurrentSeqNum)
			ccipTH.EventuallyReportCommitted(t, ccipTH.Source.OnRamp.Address(), geCurrentSeqNum)
			executionLog := ccipTH.NoNodesHaveExecutedSeqNum(t, geCurrentSeqNum)
			require.Empty(t, executionLog)
			pendingReqNumbers = append(pendingReqNumbers, geCurrentSeqNum)
			geCurrentSeqNum++
		}

		// manually execute the failed request
		failedSeqNum := ccipTH.ExecuteMessage(t, failedReqLog, txForFailedReq.Hash(), currentBlockNumber)
		currentBlockNumber = ccipTH.Dest.Chain.Blockchain().CurrentBlock().Number.Uint64()
		ccipTH.EventuallyExecutionStateChangedToSuccess(t, []uint64{failedSeqNum}, currentBlockNumber)

		// verify all the pending requests should be successfully executed now
		for _, seqNo := range pendingReqNumbers {
			t.Logf("Verify execution for pending seqNum %d", seqNo)
			ccipTH.EventuallyExecutionStateChangedToSuccess(t, []uint64{uint64(seqNo)}, 1)
		}
	})

	// Deploy new on ramp,Commit store,off ramp
	// create new jobs
	// Send a number of requests
	// Verify all requests after the contracts are upgraded
	t.Run("upgrade contracts and verify requests can be sent with upgraded contract", func(t *testing.T) {
		ccipTH.DeployNewOnRamp(t)
		ccipTH.DeployNewCommitStore(t)
		ccipTH.DeployNewOffRamp(t)
		newConfigBlock := ccipTH.Dest.Chain.Blockchain().CurrentBlock().Number.Int64()
		// delete previous jobs, 1 commit and exec
		for _, node := range ccipTH.Nodes {
			err = node.App.DeleteJob(context.Background(), 1)
			require.NoError(t, err)
			err = node.App.DeleteJob(context.Background(), 2)
			require.NoError(t, err)
		}

		// enable the newly deployed contracts
		ccipTH.EnableOnRamp(t)
		ccipTH.EnableOffRamp(t)
		ccipTH.EnableCommitStore(t)

		// create updated jobs
		jobParams = ccipTH.NewCCIPJobSpecParams(tokenPricesUSDPipeline, newConfigBlock)
		ccipTH.AddAllJobs(t, jobParams)

		startSeq := 1
		endSeqNum := 3
		gasLimit := big.NewInt(200_003) // prime number
		gasPrice := big.NewInt(1e9)     // 1 gwei
		tokenAmount := big.NewInt(100)
		for i := startSeq; i <= endSeqNum; i++ {
			t.Logf("sending request for seqnum %d", i)
			ccipTH.SendMessage(t, gasLimit, gasPrice, tokenAmount, ccipTH.Dest.Receivers[0].Receiver.Address())
			ccipTH.Source.Chain.Commit()
			ccipTH.Dest.Chain.Commit()
			t.Logf("verifying seqnum %d", i)
			ccipTH.AllNodesHaveReqSeqNum(t, i)
			ccipTH.EventuallyReportCommitted(t, ccipTH.Source.OnRamp.Address(), i)
			executionLog := ccipTH.AllNodesHaveExecutedSeqNums(t, i, i)
			ccipTH.AssertExecState(t, executionLog[0], abihelpers.ExecutionStateSuccess)
		}

		geCurrentSeqNum = endSeqNum + 1
	})

	t.Run("pay nops", func(t *testing.T) {
		linkToTransferToOnRamp := big.NewInt(1e18)

		// transfer some link to onramp to pay the nops
		_, err = ccipTH.Source.LinkToken.Transfer(
			ccipTH.Source.User, ccipTH.Source.OnRamp.Address(), linkToTransferToOnRamp)
		require.NoError(t, err)
		ccipTH.Source.Chain.Commit()

		srcBalReq := []testhelpers.BalanceReq{
			{
				Name:   testhelpers.Sender,
				Addr:   ccipTH.Source.User.From,
				Getter: ccipTH.GetSourceWrappedTokenBalance,
			},
			{
				Name:   testhelpers.OnRampNative,
				Addr:   ccipTH.Source.OnRamp.Address(),
				Getter: ccipTH.GetSourceWrappedTokenBalance,
			},
			{
				Name:   testhelpers.OnRamp,
				Addr:   ccipTH.Source.OnRamp.Address(),
				Getter: ccipTH.GetSourceLinkBalance,
			},
			{
				Name:   testhelpers.SourceRouter,
				Addr:   ccipTH.Source.Router.Address(),
				Getter: ccipTH.GetSourceWrappedTokenBalance,
			},
		}

		var nopsAndWeights []evm_2_evm_onramp.EVM2EVMOnRampNopAndWeight
		var totalWeight uint16
		nodes := ccipTH.Nodes
		for i := range nodes {
			// For now set the transmitter addresses to be the same as the payee addresses
			nodes[i].PaymentReceiver = nodes[i].Transmitter
			nopsAndWeights = append(nopsAndWeights, evm_2_evm_onramp.EVM2EVMOnRampNopAndWeight{
				Nop:    nodes[i].PaymentReceiver,
				Weight: 5,
			})
			totalWeight += 5
			srcBalReq = append(srcBalReq, testhelpers.BalanceReq{
				Name:   fmt.Sprintf("node %d", i),
				Addr:   nodes[i].PaymentReceiver,
				Getter: ccipTH.GetSourceLinkBalance,
			})
		}
		srcBalances, err := testhelpers.GetBalances(t, srcBalReq)
		require.NoError(t, err)

		// set nops on the onramp
		ccipTH.SetNopsOnRamp(t, nopsAndWeights)

		// send a message
		extraArgs, err := testhelpers.GetEVMExtraArgsV1(big.NewInt(200_000), true)
		require.NoError(t, err)

		// FeeToken is empty, therefore it should use native token
		msg := router.ClientEVM2AnyMessage{
			Receiver:     testhelpers.MustEncodeAddress(t, ccipTH.Dest.Receivers[1].Receiver.Address()),
			Data:         []byte("hello"),
			TokenAmounts: []router.ClientEVMTokenAmount{},
			ExtraArgs:    extraArgs,
		}
		fee, err := ccipTH.Source.Router.GetFee(nil, testhelpers.DestChainID, msg)
		require.NoError(t, err)

		// verify message is sent
		ccipTH.Source.User.Value = fee
		ccipTH.SendRequest(t, msg)
		ccipTH.Source.User.Value = nil
		ccipTH.AllNodesHaveReqSeqNum(t, geCurrentSeqNum)
		ccipTH.EventuallyReportCommitted(t, ccipTH.Source.OnRamp.Address(), geCurrentSeqNum)

		executionLogs := ccipTH.AllNodesHaveExecutedSeqNums(t, geCurrentSeqNum, geCurrentSeqNum)
		assert.Len(t, executionLogs, 1)
		ccipTH.AssertExecState(t, executionLogs[0], abihelpers.ExecutionStateSuccess)
		geCurrentSeqNum++

		// get the nop fee
		nopFee, err := ccipTH.Source.OnRamp.GetNopFeesJuels(nil)
		require.NoError(t, err)
		t.Log("nopFee", nopFee)

		// withdraw fees and verify there is still fund left for nop payment
		_, err = ccipTH.Source.OnRamp.WithdrawNonLinkFees(
			ccipTH.Source.User,
			ccipTH.Source.WrappedNative.Address(),
			ccipTH.Source.User.From,
		)
		require.NoError(t, err)
		ccipTH.Source.Chain.Commit()

		// pay nops
		_, err = ccipTH.Source.OnRamp.PayNops(ccipTH.Source.User)
		require.NoError(t, err)
		ccipTH.Source.Chain.Commit()

		srcBalanceAssertions := []testhelpers.BalanceAssertion{
			{
				// Onramp should not have any balance left in wrapped native
				Name:     testhelpers.OnRampNative,
				Address:  ccipTH.Source.OnRamp.Address(),
				Expected: big.NewInt(0).String(),
				Getter:   ccipTH.GetSourceWrappedTokenBalance,
			},
			{
				// Onramp should have the remaining link after paying nops
				Name:     testhelpers.OnRamp,
				Address:  ccipTH.Source.OnRamp.Address(),
				Expected: new(big.Int).Sub(srcBalances[testhelpers.OnRamp], nopFee).String(),
				Getter:   ccipTH.GetSourceLinkBalance,
			},
			{
				Name:     testhelpers.SourceRouter,
				Address:  ccipTH.Source.Router.Address(),
				Expected: srcBalances[testhelpers.SourceRouter].String(),
				Getter:   ccipTH.GetSourceWrappedTokenBalance,
			},
			// onRamp's balance (of previously sent fee during message sending) should have been transferred to
			// the owner as a result of WithdrawNonLinkFees
			{
				Name:     testhelpers.Sender,
				Address:  ccipTH.Source.User.From,
				Expected: fee.String(),
				Getter:   ccipTH.GetSourceWrappedTokenBalance,
			},
		}

		// the nodes should be paid according to the weights assigned
		for i, node := range nodes {
			paymentWeight := float64(nopsAndWeights[i].Weight) / float64(totalWeight)
			paidInFloat := paymentWeight * float64(nopFee.Int64())
			paid, _ := new(big.Float).SetFloat64(paidInFloat).Int64()
			bal := new(big.Int).Add(
				new(big.Int).SetInt64(paid),
				srcBalances[fmt.Sprintf("node %d", i)]).String()
			srcBalanceAssertions = append(srcBalanceAssertions, testhelpers.BalanceAssertion{
				Name:     fmt.Sprintf("node %d", i),
				Address:  node.PaymentReceiver,
				Expected: bal,
				Getter:   ccipTH.GetSourceLinkBalance,
			})
		}
		ccipTH.AssertBalances(t, srcBalanceAssertions)
	})
}
