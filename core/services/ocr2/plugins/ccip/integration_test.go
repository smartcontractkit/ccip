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
)

func TestIntegration_CCIP(t *testing.T) {
	ccipContracts := testhelpers.SetupCCIPContracts(t, testhelpers.SourceChainID, testhelpers.DestChainID)
	linkUSD := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte(`{"UsdPerLink": "8000000000000000000"}`))
		require.NoError(t, err)
	}))
	ethUSD := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte(`{"UsdPerETH": "1700000000000000000000"}`))
		require.NoError(t, err)
	}))
	wrapped, err := ccipContracts.Source.Router.GetWrappedNative(nil)
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
		linkUSD.URL, ethUSD.URL, ccipContracts.Dest.LinkToken.Address(), wrapped)
	defer linkUSD.Close()
	defer ethUSD.Close()

	nodes, jobParams := ccipContracts.SetUpNodesAndJobs(t, tokenPricesUSDPipeline, 19399)

	geCurrentSeqNum := 1

	t.Run("single ge", func(t *testing.T) {
		tokenAmount := big.NewInt(500000003) // prime number
		gasLimit := big.NewInt(200_003)      // prime number
		// gasPrice := big.NewInt(1e9)          // 1 gwei

		extraArgs, err2 := testhelpers.GetEVMExtraArgsV1(gasLimit, false)
		require.NoError(t, err2)

		sourceBalances, err2 := testhelpers.GetBalances(t, []testhelpers.BalanceReq{
			{Name: testhelpers.SourcePool, Addr: ccipContracts.Source.Pool.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: testhelpers.OnRamp, Addr: ccipContracts.Source.OnRamp.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: testhelpers.SourceRouter, Addr: ccipContracts.Source.Router.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: testhelpers.SourcePrices, Addr: ccipContracts.Source.PriceRegistry.Address(), Getter: ccipContracts.GetSourceLinkBalance},
		})
		require.NoError(t, err2)
		destBalances, err2 := testhelpers.GetBalances(t, []testhelpers.BalanceReq{
			{Name: testhelpers.Receiver, Addr: ccipContracts.Dest.Receivers[0].Receiver.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: testhelpers.DestPool, Addr: ccipContracts.Dest.Pool.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: testhelpers.OffRamp, Addr: ccipContracts.Dest.OffRamp.Address(), Getter: ccipContracts.GetDestLinkBalance},
		})
		require.NoError(t, err2)

		msg := router.ClientEVM2AnyMessage{
			Receiver: testhelpers.MustEncodeAddress(t, ccipContracts.Dest.Receivers[0].Receiver.Address()),
			Data:     []byte("hello"),
			TokenAmounts: []router.ClientEVMTokenAmount{
				{
					Token:  ccipContracts.Source.LinkToken.Address(),
					Amount: tokenAmount,
				},
			},
			FeeToken:  ccipContracts.Source.LinkToken.Address(),
			ExtraArgs: extraArgs,
		}
		fee, err2 := ccipContracts.Source.Router.GetFee(nil, testhelpers.DestChainID, msg)
		require.NoError(t, err2)
		// Currently no overhead and 10gwei dest gas price. So fee is simply (gasLimit * gasPrice)* link/native
		// require.Equal(t, new(big.Int).Mul(gasLimit, gasPrice).String(), fee.String())
		// Approve the fee amount + the token amount
		_, err2 = ccipContracts.Source.LinkToken.Approve(ccipContracts.Source.User, ccipContracts.Source.Router.Address(), new(big.Int).Add(fee, tokenAmount))
		require.NoError(t, err2)
		ccipContracts.Source.Chain.Commit()
		ccipContracts.SendRequest(t, msg)
		// Should eventually see this executed.
		ccipContracts.AllNodesHaveReqSeqNum(t, ccipContracts.Source.OnRamp.Address(), nodes, geCurrentSeqNum)
		ccipContracts.EventuallyReportCommitted(t, ccipContracts.Source.OnRamp.Address(), geCurrentSeqNum)

		executionLogs := ccipContracts.AllNodesHaveExecutedSeqNums(t, ccipContracts.Dest.OffRamp.Address(), nodes, geCurrentSeqNum, geCurrentSeqNum)
		assert.Len(t, executionLogs, 1)
		ccipContracts.AssertExecState(t, executionLogs[0], abihelpers.ExecutionStateSuccess)

		// Asserts
		// 1) The total pool input == total pool output
		// 2) Pool flow equals tokens sent
		// 3) Sent tokens arrive at the receiver

		ccipContracts.AssertBalances(t, []testhelpers.BalanceAssertion{
			{
				Name:     testhelpers.SourcePool,
				Address:  ccipContracts.Source.Pool.Address(),
				Expected: testhelpers.MustAddBigInt(sourceBalances[testhelpers.SourcePool], tokenAmount.String()).String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			},
			{
				Name:     testhelpers.SourcePrices,
				Address:  ccipContracts.Source.PriceRegistry.Address(),
				Expected: sourceBalances[testhelpers.SourcePrices].String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			},
			{
				// Fees end up in the onramp.
				Name:     testhelpers.OnRamp,
				Address:  ccipContracts.Source.OnRamp.Address(),
				Expected: testhelpers.MustAddBigInt(sourceBalances[testhelpers.SourcePrices], fee.String()).String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			},
			{
				Name:     testhelpers.SourceRouter,
				Address:  ccipContracts.Source.Router.Address(),
				Expected: sourceBalances[testhelpers.SourceRouter].String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			},
		})
		ccipContracts.AssertBalances(t, []testhelpers.BalanceAssertion{
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
				Receiver: testhelpers.MustEncodeAddress(t, ccipContracts.Dest.Receivers[0].Receiver.Address()),
				Data:     []byte("hello"),
				TokenAmounts: []router.ClientEVMTokenAmount{
					{
						Token:  ccipContracts.Source.LinkToken.Address(),
						Amount: tokenAmount,
					},
				},
				FeeToken:  ccipContracts.Source.LinkToken.Address(),
				ExtraArgs: extraArgs,
			}
			fee, err2 := ccipContracts.Source.Router.GetFee(nil, testhelpers.DestChainID, msg)
			require.NoError(t, err2)
			// Currently no overhead and 1gwei dest gas price. So fee is simply gasLimit * gasPrice.
			// require.Equal(t, new(big.Int).Mul(txGasLimit, gasPrice).String(), fee.String())
			// Approve the fee amount + the token amount
			_, err2 = ccipContracts.Source.LinkToken.Approve(ccipContracts.Source.User, ccipContracts.Source.Router.Address(), new(big.Int).Add(fee, tokenAmount))
			require.NoError(t, err2)
			tx, err2 := ccipContracts.Source.Router.CcipSend(ccipContracts.Source.User, ccipContracts.Dest.ChainID, msg)
			require.NoError(t, err2)
			txs = append(txs, tx)
		}

		// Send a batch of requests in a single block
		testhelpers.ConfirmTxs(t, txs, ccipContracts.Source.Chain)
		for i := 0; i < n; i++ {
			ccipContracts.AllNodesHaveReqSeqNum(t, ccipContracts.Source.OnRamp.Address(), nodes, geCurrentSeqNum+i)
		}
		// Should see a report with the full range
		ccipContracts.EventuallyReportCommitted(t, ccipContracts.Source.OnRamp.Address(), geCurrentSeqNum+n-1)
		// Should all be executed
		executionLogs := ccipContracts.AllNodesHaveExecutedSeqNums(t, ccipContracts.Dest.OffRamp.Address(), nodes, geCurrentSeqNum, geCurrentSeqNum+n-1)
		for _, execLog := range executionLogs {
			ccipContracts.AssertExecState(t, execLog, abihelpers.ExecutionStateSuccess)
		}

		geCurrentSeqNum += n
	})

	t.Run("ge strict sequencing", func(t *testing.T) {
		// approve the total amount to be sent
		// set revert to true so that the execution gets reverted
		_, err = ccipContracts.Dest.Receivers[1].Receiver.SetRevert(ccipContracts.Dest.User, true)
		require.NoError(t, err, "setting revert to true on the receiver")
		ccipContracts.Dest.Chain.Commit()
		currentBlockNumber := ccipContracts.Dest.Chain.Blockchain().CurrentBlock().Number.Uint64()

		// Test sequence:
		// Send msg1: strict reverts
		// Send msg2, msg3: blocked on manual exec
		// Execute msg1 manually.
		// msg2 and msg2 should go through.
		totalMsgs := 2
		extraArgs, err2 := testhelpers.GetEVMExtraArgsV1(big.NewInt(200_000), true)
		require.NoError(t, err2)
		startNonce, err2 := ccipContracts.Dest.OffRamp.GetSenderNonce(nil, ccipContracts.Source.User.From)
		require.NoError(t, err2)
		msg := router.ClientEVM2AnyMessage{
			Receiver:     testhelpers.MustEncodeAddress(t, ccipContracts.Dest.Receivers[1].Receiver.Address()),
			Data:         []byte("hello"),
			TokenAmounts: []router.ClientEVMTokenAmount{},
			FeeToken:     ccipContracts.Source.LinkToken.Address(),
			ExtraArgs:    extraArgs,
		}
		fee, err2 := ccipContracts.Source.Router.GetFee(nil, testhelpers.DestChainID, msg)
		require.NoError(t, err2)
		// Approve the fee amount
		_, err2 = ccipContracts.Source.LinkToken.Approve(ccipContracts.Source.User, ccipContracts.Source.Router.Address(), big.NewInt(0).Mul(big.NewInt(int64(totalMsgs)), fee))
		require.NoError(t, err2)
		ccipContracts.Source.Chain.Commit()
		txForFailedReq := ccipContracts.SendRequest(t, msg)
		failedReqLog := ccipContracts.AllNodesHaveReqSeqNum(t, ccipContracts.Source.OnRamp.Address(), nodes, geCurrentSeqNum)
		ccipContracts.EventuallyReportCommitted(t, ccipContracts.Source.OnRamp.Address(), geCurrentSeqNum)
		ccipContracts.EventuallyCommitReportAccepted(t, currentBlockNumber)

		// execution status should be failed
		executionLogs := ccipContracts.AllNodesHaveExecutedSeqNums(t, ccipContracts.Dest.OffRamp.Address(), nodes, geCurrentSeqNum, geCurrentSeqNum)
		assert.Len(t, executionLogs, 1)
		ccipContracts.AssertExecState(t, executionLogs[0], abihelpers.ExecutionStateFailure)
		// Nonce should not have incremented
		afterNonce, err2 := ccipContracts.Dest.OffRamp.GetSenderNonce(nil, ccipContracts.Source.User.From)
		require.NoError(t, err2)
		require.Equal(t, startNonce, afterNonce)
		geCurrentSeqNum++

		// flip the revert settings on receiver
		_, err2 = ccipContracts.Dest.Receivers[1].Receiver.SetRevert(ccipContracts.Dest.User, false)
		require.NoError(t, err2, "setting revert to false on the receiver")
		ccipContracts.Dest.Chain.Commit()
		ccipContracts.Source.Chain.Commit()

		// subsequent requests which should not be executed.
		var pendingReqNumbers []int
		for i := 1; i < totalMsgs; i++ {
			ccipContracts.SendRequest(t, msg)
			ccipContracts.AllNodesHaveReqSeqNum(t, ccipContracts.Source.OnRamp.Address(), nodes, geCurrentSeqNum)
			ccipContracts.EventuallyReportCommitted(t, ccipContracts.Source.OnRamp.Address(), geCurrentSeqNum)
			executionLog := ccipContracts.NoNodesHaveExecutedSeqNum(t, ccipContracts.Dest.OffRamp.Address(), nodes, geCurrentSeqNum)
			require.Empty(t, executionLog)
			pendingReqNumbers = append(pendingReqNumbers, geCurrentSeqNum)
			geCurrentSeqNum++
		}

		// manually execute the failed request
		failedSeqNum := ccipContracts.ExecuteMessage(t, failedReqLog, txForFailedReq.Hash(), currentBlockNumber)
		currentBlockNumber = ccipContracts.Dest.Chain.Blockchain().CurrentBlock().Number.Uint64()
		ccipContracts.EventuallyExecutionStateChangedToSuccess(t, []uint64{failedSeqNum}, currentBlockNumber)

		// verify all the pending requests should be successfully executed now
		for _, seqNo := range pendingReqNumbers {
			t.Logf("Verify execution for pending seqNum %d", seqNo)
			ccipContracts.EventuallyExecutionStateChangedToSuccess(t, []uint64{uint64(seqNo)}, 1)
		}
	})

	// Deploy new on ramp,Commit store,off ramp
	// Delete previous jobs
	// Enable new contracts
	// Create new jobs
	// Send a number of requests
	// Verify all requests after the contracts are upgraded
	t.Run("upgrade contracts and verify requests can be sent with upgraded contract", func(t *testing.T) {
		ccipContracts.DeployNewOnRamp(t)
		ccipContracts.DeployNewCommitStore(t)
		ccipContracts.DeployNewOffRamp(t)
		newConfigBlock := ccipContracts.Dest.Chain.Blockchain().CurrentBlock().Number.Int64()
		// delete previous jobs, 1 commit and exec
		for _, node := range nodes {
			err = node.App.DeleteJob(context.Background(), 1)
			require.NoError(t, err)
			err = node.App.DeleteJob(context.Background(), 2)
			require.NoError(t, err)
		}

		// enable the newly deployed contracts
		ccipContracts.EnableOnRamp(t)
		ccipContracts.EnableOffRamp(t)
		ccipContracts.EnableCommitStore(t)

		// create updated jobs
		jobParams = ccipContracts.NewCCIPJobSpecParams(tokenPricesUSDPipeline, newConfigBlock)
		ccipContracts.AddAllJobs(t, jobParams, nodes)

		startSeq := 1
		endSeqNum := 3
		gasLimit := big.NewInt(200_003) // prime number
		gasPrice := big.NewInt(1e9)     // 1 gwei
		tokenAmount := big.NewInt(100)
		for i := startSeq; i <= endSeqNum; i++ {
			t.Logf("sending request for seqnum %d", i)
			ccipContracts.SendMessage(t, gasLimit, gasPrice, tokenAmount, ccipContracts.Dest.Receivers[0].Receiver.Address())
			ccipContracts.Source.Chain.Commit()
			ccipContracts.Dest.Chain.Commit()
			t.Logf("verifying seqnum %d", i)
			ccipContracts.AllNodesHaveReqSeqNum(t, ccipContracts.Source.OnRamp.Address(), nodes, i)
			ccipContracts.EventuallyReportCommitted(t, ccipContracts.Source.OnRamp.Address(), i)
			executionLog := ccipContracts.AllNodesHaveExecutedSeqNums(t, ccipContracts.Dest.OffRamp.Address(), nodes, i, i)
			ccipContracts.AssertExecState(t, executionLog[0], abihelpers.ExecutionStateSuccess)
		}

		geCurrentSeqNum = endSeqNum + 1
	})

	t.Run("pay nops", func(t *testing.T) {
		linkToTransferToOnRamp := big.NewInt(1e18)

		// transfer some link to onramp to pay the nops
		_, err = ccipContracts.Source.LinkToken.Transfer(
			ccipContracts.Source.User, ccipContracts.Source.OnRamp.Address(), linkToTransferToOnRamp)
		require.NoError(t, err)
		ccipContracts.Source.Chain.Commit()

		srcBalReq := []testhelpers.BalanceReq{
			{
				Name:   testhelpers.Sender,
				Addr:   ccipContracts.Source.User.From,
				Getter: ccipContracts.GetSourceWrappedTokenBalance,
			},
			{
				Name:   testhelpers.OnRampNative,
				Addr:   ccipContracts.Source.OnRamp.Address(),
				Getter: ccipContracts.GetSourceWrappedTokenBalance,
			},
			{
				Name:   testhelpers.OnRamp,
				Addr:   ccipContracts.Source.OnRamp.Address(),
				Getter: ccipContracts.GetSourceLinkBalance,
			},
			{
				Name:   testhelpers.SourceRouter,
				Addr:   ccipContracts.Source.Router.Address(),
				Getter: ccipContracts.GetSourceWrappedTokenBalance,
			},
		}

		var nopsAndWeights []evm_2_evm_onramp.EVM2EVMOnRampNopAndWeight
		var totalWeight uint16
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
				Getter: ccipContracts.GetSourceLinkBalance,
			})
		}
		srcBalances, err := testhelpers.GetBalances(t, srcBalReq)
		require.NoError(t, err)

		// set nops on the onramp
		ccipContracts.SetNopsOnRamp(t, nopsAndWeights)

		// send a message
		extraArgs, err := testhelpers.GetEVMExtraArgsV1(big.NewInt(200_000), true)
		require.NoError(t, err)

		// FeeToken is empty, therefore it should use native token
		msg := router.ClientEVM2AnyMessage{
			Receiver:     testhelpers.MustEncodeAddress(t, ccipContracts.Dest.Receivers[1].Receiver.Address()),
			Data:         []byte("hello"),
			TokenAmounts: []router.ClientEVMTokenAmount{},
			ExtraArgs:    extraArgs,
		}
		fee, err := ccipContracts.Source.Router.GetFee(nil, testhelpers.DestChainID, msg)
		require.NoError(t, err)

		// verify message is sent
		ccipContracts.Source.User.Value = fee
		ccipContracts.SendRequest(t, msg)
		ccipContracts.Source.User.Value = nil
		ccipContracts.AllNodesHaveReqSeqNum(t, ccipContracts.Source.OnRamp.Address(), nodes, geCurrentSeqNum)
		ccipContracts.EventuallyReportCommitted(t, ccipContracts.Source.OnRamp.Address(), geCurrentSeqNum)

		executionLogs := ccipContracts.AllNodesHaveExecutedSeqNums(t, ccipContracts.Dest.OffRamp.Address(), nodes, geCurrentSeqNum, geCurrentSeqNum)
		assert.Len(t, executionLogs, 1)
		ccipContracts.AssertExecState(t, executionLogs[0], abihelpers.ExecutionStateSuccess)
		geCurrentSeqNum++

		// get the nop fee
		nopFee, err := ccipContracts.Source.OnRamp.GetNopFeesJuels(nil)
		require.NoError(t, err)
		t.Log("nopFee", nopFee)

		// withdraw fees and verify there is still fund left for nop payment
		_, err = ccipContracts.Source.OnRamp.WithdrawNonLinkFees(
			ccipContracts.Source.User,
			ccipContracts.Source.WrappedNative.Address(),
			ccipContracts.Source.User.From,
		)
		require.NoError(t, err)
		ccipContracts.Source.Chain.Commit()

		// pay nops
		_, err = ccipContracts.Source.OnRamp.PayNops(ccipContracts.Source.User)
		require.NoError(t, err)
		ccipContracts.Source.Chain.Commit()

		srcBalanceAssertions := []testhelpers.BalanceAssertion{
			{
				// Onramp should not have any balance left in wrapped native
				Name:     testhelpers.OnRampNative,
				Address:  ccipContracts.Source.OnRamp.Address(),
				Expected: big.NewInt(0).String(),
				Getter:   ccipContracts.GetSourceWrappedTokenBalance,
			},
			{
				// Onramp should have the remaining link after paying nops
				Name:     testhelpers.OnRamp,
				Address:  ccipContracts.Source.OnRamp.Address(),
				Expected: new(big.Int).Sub(srcBalances[testhelpers.OnRamp], nopFee).String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			},
			{
				Name:     testhelpers.SourceRouter,
				Address:  ccipContracts.Source.Router.Address(),
				Expected: srcBalances[testhelpers.SourceRouter].String(),
				Getter:   ccipContracts.GetSourceWrappedTokenBalance,
			},
			// onRamp's balance (of previously sent fee during message sending) should have been transferred to
			// the owner as a result of WithdrawNonLinkFees
			{
				Name:     testhelpers.Sender,
				Address:  ccipContracts.Source.User.From,
				Expected: fee.String(),
				Getter:   ccipContracts.GetSourceWrappedTokenBalance,
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
				Getter:   ccipContracts.GetSourceLinkBalance,
			})
		}
		ccipContracts.AssertBalances(t, srcBalanceAssertions)
	})
}
