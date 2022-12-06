package ccip_test

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"

	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_toll_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/ge_router"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/testhelpers"
)

var (
	sourceChainID = big.NewInt(1000)
	destChainID   = big.NewInt(1337)
)

func TestIntegration_CCIP(t *testing.T) {
	ccipContracts := testhelpers.SetupCCIPContracts(t, sourceChainID, destChainID)
	bootstrapNodePort := int64(19598)
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
		linkEth.URL, ccipContracts.DestLinkToken.Address())
	jobParams := ccipContracts.NewCCIPJobSpecParams(tokensPerFeeCoinPipeline, configBlock)
	defer linkEth.Close()

	// Add the bootstrap job
	bootstrapNode.AddBootstrapJob(t, jobParams.BootstrapJob(ccipContracts.CommitStore.Address().Hex()))
	testhelpers.AddAllJobs(t, jobParams, ccipContracts, nodes)

	// Replay for bootstrap.
	bc, err := bootstrapNode.App.GetChains().EVM.Get(destChainID)
	require.NoError(t, err)
	require.NoError(t, bc.LogPoller().Replay(context.Background(), configBlock))

	tollCurrentSeqNum := 1
	geCurrentSeqNum := 1

	ccipContracts.DestChain.Commit()

	t.Run("single ge", func(t *testing.T) {
		tokenAmount := big.NewInt(500000003) // prime number
		gasLimit := big.NewInt(200_003)      // prime number
		gasPrice := big.NewInt(1e9)          // 1 gwei

		eventSignatures := ccip.GetGEEventSignatures()
		extraArgs, err := testhelpers.GetEVMExtraArgsV1(gasLimit, false)
		require.NoError(t, err)

		sourceBalances, err := testhelpers.GetBalances([]testhelpers.BalanceReq{
			{Name: testhelpers.SourcePool, Addr: ccipContracts.SourcePool.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: testhelpers.GEOnRamp, Addr: ccipContracts.GEOnRamp.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: testhelpers.SourceGERouter, Addr: ccipContracts.SourceGERouter.Address(), Getter: ccipContracts.GetSourceLinkBalance},
		})
		require.NoError(t, err)
		destBalances, err := testhelpers.GetBalances([]testhelpers.BalanceReq{
			{Name: testhelpers.Receiver, Addr: ccipContracts.Receivers[0].Receiver.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: testhelpers.DestPool, Addr: ccipContracts.DestPool.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: testhelpers.GEOffRamp, Addr: ccipContracts.GEOffRamp.Address(), Getter: ccipContracts.GetDestLinkBalance},
		})
		require.NoError(t, err)

		msg := ge_router.GEConsumerEVM2AnyGEMessage{
			Receiver: testhelpers.MustEncodeAddress(t, ccipContracts.Receivers[0].Receiver.Address()),
			Data:     []byte("hello"),
			TokensAndAmounts: []ge_router.CommonEVMTokenAndAmount{
				{
					Token:  ccipContracts.SourceLinkToken.Address(),
					Amount: tokenAmount,
				},
			},
			FeeToken:  ccipContracts.SourceLinkToken.Address(),
			ExtraArgs: extraArgs,
		}
		fee, err := ccipContracts.SourceGERouter.GetFee(nil, destChainID, msg)
		require.NoError(t, err)
		// Currently no overhead and 1gwei dest gas price. So fee is simply gasLimit * gasPrice.
		require.Equal(t, new(big.Int).Mul(gasLimit, gasPrice).String(), fee.String())
		// Approve the fee amount + the token amount
		_, err = ccipContracts.SourceLinkToken.Approve(ccipContracts.SourceUser, ccipContracts.SourceGERouter.Address(), new(big.Int).Add(fee, tokenAmount))
		require.NoError(t, err)
		ccipContracts.SourceChain.Commit()
		testhelpers.SendGERequest(t, ccipContracts, msg)
		// Should eventually see this executed.
		testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, eventSignatures, ccipContracts.GEOnRamp.Address(), nodes, geCurrentSeqNum)
		testhelpers.EventuallyReportCommitted(t, ccipContracts, ccipContracts.GEOnRamp.Address(), geCurrentSeqNum, geCurrentSeqNum)

		executionLogs := testhelpers.AllNodesHaveExecutedSeqNums(t, ccipContracts, eventSignatures, ccipContracts.GEOffRamp.Address(), nodes, geCurrentSeqNum, geCurrentSeqNum)
		assert.Len(t, executionLogs, 1)
		testhelpers.AssertGEExecSuccess(t, ccipContracts, executionLogs[0])

		// Asserts
		// 1) The total pool input == total pool output
		// 2) Pool flow equals tokens sent + fees
		// 3) Sent tokens arrive at the receiver
		// 4) Fees are kept in the offRamp
		poolInOutFlow := new(big.Int).Add(tokenAmount, fee).String()

		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     testhelpers.SourcePool,
				Address:  ccipContracts.SourcePool.Address(),
				Expected: testhelpers.MustAddBigInt(sourceBalances[testhelpers.SourcePool], poolInOutFlow).String(), // Tokens & fee both locked in the pool
				Getter:   ccipContracts.GetSourceLinkBalance,
			},
			{
				Name:     testhelpers.GEOnRamp,
				Address:  ccipContracts.GEOnRamp.Address(),
				Expected: sourceBalances[testhelpers.GEOnRamp].String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			},
			{
				Name:     testhelpers.SourceGERouter,
				Address:  ccipContracts.SourceGERouter.Address(),
				Expected: sourceBalances[testhelpers.SourceGERouter].String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			},
		})
		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     testhelpers.Receiver,
				Address:  ccipContracts.Receivers[0].Receiver.Address(),
				Expected: testhelpers.MustAddBigInt(destBalances[testhelpers.Receiver], tokenAmount.String()).String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			},
			{
				Name:     testhelpers.DestPool,
				Address:  ccipContracts.DestPool.Address(),
				Expected: testhelpers.MustSubBigInt(destBalances[testhelpers.DestPool], poolInOutFlow).String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			},
			{
				Name:     testhelpers.GEOffRamp,
				Address:  ccipContracts.GEOffRamp.Address(),
				Expected: testhelpers.MustAddBigInt(destBalances[testhelpers.GEOffRamp], fee.String()).String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			},
		})
		geCurrentSeqNum++
	})

	t.Run("multiple batches ge", func(t *testing.T) {
		tokenAmount := big.NewInt(500000003)
		gasLimit := big.NewInt(250_000)
		gasPrice := big.NewInt(1e9) // 1 gwei

		eventSignatures := ccip.GetGEEventSignatures()
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
			msg := ge_router.GEConsumerEVM2AnyGEMessage{
				Receiver: testhelpers.MustEncodeAddress(t, ccipContracts.Receivers[0].Receiver.Address()),
				Data:     []byte("hello"),
				TokensAndAmounts: []ge_router.CommonEVMTokenAndAmount{
					{
						Token:  ccipContracts.SourceLinkToken.Address(),
						Amount: tokenAmount,
					},
				},
				FeeToken:  ccipContracts.SourceLinkToken.Address(),
				ExtraArgs: extraArgs,
			}
			fee, err := ccipContracts.SourceGERouter.GetFee(nil, destChainID, msg)
			require.NoError(t, err)
			// Currently no overhead and 1gwei dest gas price. So fee is simply gasLimit * gasPrice.
			require.Equal(t, new(big.Int).Mul(txGasLimit, gasPrice).String(), fee.String())
			// Approve the fee amount + the token amount
			_, err = ccipContracts.SourceLinkToken.Approve(ccipContracts.SourceUser, ccipContracts.SourceGERouter.Address(), new(big.Int).Add(fee, tokenAmount))
			require.NoError(t, err)
			tx, err := ccipContracts.SourceGERouter.CcipSend(ccipContracts.SourceUser, ccipContracts.DestChainID, msg)
			require.NoError(t, err)
			txs = append(txs, tx)
		}

		// Send a batch of requests in a single block
		testhelpers.ConfirmTxs(t, txs, ccipContracts.SourceChain)
		// All nodes should have all 3.
		var reqs []logpoller.Log
		for i := 0; i < n; i++ {
			reqs = append(reqs, testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, eventSignatures, ccipContracts.GEOnRamp.Address(), nodes, geCurrentSeqNum+i))
		}
		// Should see a report with the full range
		testhelpers.EventuallyReportCommitted(t, ccipContracts, ccipContracts.GEOnRamp.Address(), geCurrentSeqNum, geCurrentSeqNum+n-1)
		// Should all be executed
		executionLogs := testhelpers.AllNodesHaveExecutedSeqNums(t, ccipContracts, eventSignatures, ccipContracts.GEOffRamp.Address(), nodes, geCurrentSeqNum, geCurrentSeqNum+n-1)
		for _, execLog := range executionLogs {
			testhelpers.AssertGEExecSuccess(t, ccipContracts, execLog)
		}

		geCurrentSeqNum += n
	})

	t.Run("single auto-execute toll", func(t *testing.T) {
		eventSignatures := ccip.GetTollEventSignatures()
		// Approve router to take source token.
		tokenAmount := big.NewInt(100)
		// Example sending a msg with 100k callback execution on eth mainnet.
		// Gas price 200e9 wei/gas * 1e5 gas * (2e20 juels/eth / 1e18wei/eth)
		// 4e18 juels = 4 link, which does not include gas used outside the callback.
		// Gas outside the callback for 1 token is ~100k in the worst case.
		feeTokenAmount := new(big.Int).Mul(big.NewInt(10), big.NewInt(1e18))
		_, err = ccipContracts.SourceLinkToken.Approve(ccipContracts.SourceUser, ccipContracts.TollOnRampRouter.Address(), new(big.Int).Add(tokenAmount, feeTokenAmount))
		require.NoError(t, err)
		ccipContracts.SourceChain.Commit()

		sourceBalances, err := testhelpers.GetBalances([]testhelpers.BalanceReq{
			{Name: testhelpers.SourcePool, Addr: ccipContracts.SourcePool.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: testhelpers.TollOnRamp, Addr: ccipContracts.TollOnRamp.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: testhelpers.TollOnRampRouter, Addr: ccipContracts.TollOnRampRouter.Address(), Getter: ccipContracts.GetSourceLinkBalance},
		})
		require.NoError(t, err)
		destBalances, err := testhelpers.GetBalances([]testhelpers.BalanceReq{
			{Name: testhelpers.Receiver, Addr: ccipContracts.Receivers[0].Receiver.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: testhelpers.DestPool, Addr: ccipContracts.DestPool.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: testhelpers.TollOffRamp, Addr: ccipContracts.TollOffRamp.Address(), Getter: ccipContracts.GetDestLinkBalance},
		})
		require.NoError(t, err)

		testhelpers.SendTollRequest(t, ccipContracts, "hey DON, execute for me",
			[]evm_2_any_toll_onramp_router.CommonEVMTokenAndAmount{{
				Token:  ccipContracts.SourceLinkToken.Address(),
				Amount: tokenAmount,
			}},
			evm_2_any_toll_onramp_router.CommonEVMTokenAndAmount{
				Token:  ccipContracts.SourceLinkToken.Address(),
				Amount: feeTokenAmount,
			},
			big.NewInt(100_000),
			ccipContracts.Receivers[0].Receiver.Address())
		testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, eventSignatures, ccipContracts.TollOnRamp.Address(), nodes, tollCurrentSeqNum)
		testhelpers.EventuallyReportCommitted(t, ccipContracts, ccipContracts.TollOnRamp.Address(), tollCurrentSeqNum, tollCurrentSeqNum)
		executionLogs := testhelpers.AllNodesHaveExecutedSeqNums(t, ccipContracts, eventSignatures, ccipContracts.TollOffRamp.Address(), nodes, tollCurrentSeqNum, tollCurrentSeqNum)
		assert.Len(t, executionLogs, 1)
		testhelpers.AssertTollExecSuccess(t, ccipContracts, executionLogs[0])

		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     testhelpers.SourcePool,
				Address:  ccipContracts.SourcePool.Address(),
				Expected: testhelpers.MustAddBigInt(sourceBalances[testhelpers.SourcePool], "10000000000000000099").String(), // 10e18 + 100 transfer - 1 fee
				Getter:   ccipContracts.GetSourceLinkBalance,
			},
			{
				Name:     testhelpers.TollOnRamp,
				Address:  ccipContracts.TollOnRamp.Address(),
				Expected: sourceBalances[testhelpers.TollOnRamp].String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			},
			{
				Name:     testhelpers.TollOnRampRouter,
				Address:  ccipContracts.TollOnRampRouter.Address(),
				Expected: testhelpers.MustAddBigInt(sourceBalances[testhelpers.TollOnRampRouter], "1").String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			},
		})
		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     testhelpers.Receiver,
				Address:  ccipContracts.Receivers[0].Receiver.Address(),
				Expected: testhelpers.MustAddBigInt(destBalances[testhelpers.Receiver], "9049107200000000099").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
				Within:   "1000000000000000000"}, // Roughly 200k gas * 200e9 wei/gas * (2e20 link/eth / 1e18wei/eth)
			{
				Name:     testhelpers.DestPool,
				Address:  ccipContracts.DestPool.Address(),
				Expected: testhelpers.MustSubBigInt(destBalances[testhelpers.DestPool], "10000000000000000099").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // We lose 10 link from the pool
			{
				Name:     testhelpers.TollOffRamp,
				Address:  ccipContracts.TollOffRamp.Address(),
				Expected: testhelpers.MustAddBigInt(destBalances[testhelpers.TollOffRamp], "965804400000000000").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
				Within:   "100000000000000", // To account for change in the number of contract optimizations
			},
		})
		tollCurrentSeqNum++
	})

	t.Run("batch auto-execute toll", func(t *testing.T) {
		eventSignatures := ccip.GetTollEventSignatures()
		sourceBalances, err := testhelpers.GetBalances([]testhelpers.BalanceReq{
			{Name: testhelpers.SourcePool, Addr: ccipContracts.SourcePool.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: testhelpers.TollOnRampRouter, Addr: ccipContracts.TollOnRampRouter.Address(), Getter: ccipContracts.GetSourceLinkBalance},
		})
		require.NoError(t, err, "fetching source balance")
		destBalances, err := testhelpers.GetBalances([]testhelpers.BalanceReq{
			{Name: testhelpers.Receiver, Addr: ccipContracts.Receivers[0].Receiver.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: testhelpers.DestPool, Addr: ccipContracts.DestPool.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: testhelpers.TollOffRampRouter, Addr: ccipContracts.TollOffRampRouter.Address(), Getter: ccipContracts.GetDestLinkBalance},
		})
		require.NoError(t, err, "fetching dest balance")

		tokenAmount := big.NewInt(100)
		feeTokenAmount := new(big.Int).Mul(big.NewInt(10), big.NewInt(1e18))
		var txs []*gethtypes.Transaction
		n := 3
		for i := 0; i < n; i++ {
			_, err = ccipContracts.SourceLinkToken.Approve(ccipContracts.SourceUser, ccipContracts.TollOnRampRouter.Address(), new(big.Int).Add(tokenAmount, feeTokenAmount))
			require.NoError(t, err)
			txs = append(txs, testhelpers.QueueTollRequest(t, ccipContracts, fmt.Sprintf("batch request %d", tollCurrentSeqNum+i), []evm_2_any_toll_onramp_router.CommonEVMTokenAndAmount{{
				Token:  ccipContracts.SourceLinkToken.Address(),
				Amount: tokenAmount,
			}}, evm_2_any_toll_onramp_router.CommonEVMTokenAndAmount{
				Token:  ccipContracts.SourceLinkToken.Address(),
				Amount: feeTokenAmount,
			}, big.NewInt(100_000), ccipContracts.Receivers[0].Receiver.Address()))
		}
		// Send a batch of requests in a single block
		testhelpers.ConfirmTxs(t, txs, ccipContracts.SourceChain)
		// All nodes should have all 3.
		var reqs []logpoller.Log
		for i := 0; i < n; i++ {
			reqs = append(reqs, testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, eventSignatures, ccipContracts.TollOnRamp.Address(), nodes, tollCurrentSeqNum+i))
		}
		// Should see a report with the full range
		testhelpers.EventuallyReportCommitted(t, ccipContracts, ccipContracts.TollOnRamp.Address(), tollCurrentSeqNum, tollCurrentSeqNum+n-1)
		// Should all be executed
		executionLogs := testhelpers.AllNodesHaveExecutedSeqNums(t, ccipContracts, eventSignatures, ccipContracts.TollOffRamp.Address(), nodes, tollCurrentSeqNum, tollCurrentSeqNum+n-1)
		for _, execLog := range executionLogs {
			testhelpers.AssertTollExecSuccess(t, ccipContracts, execLog)
		}

		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     testhelpers.SourcePool,
				Address:  ccipContracts.SourcePool.Address(),
				Expected: testhelpers.MustAddBigInt(sourceBalances[testhelpers.SourcePool], "30000000000000000297").String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			}, // (10e18 + 100 - 1)*3
			{
				Name:     testhelpers.TollOnRampRouter,
				Address:  ccipContracts.TollOnRampRouter.Address(),
				Expected: testhelpers.MustAddBigInt(sourceBalances[testhelpers.TollOnRampRouter], "3").String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			},
		})
		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     testhelpers.Receiver,
				Address:  ccipContracts.Receivers[0].Receiver.Address(),
				Expected: testhelpers.MustAddBigInt(destBalances[testhelpers.Receiver], "27225848400000000297").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
				Within:   "1000000000000000000",
			}, // 3 toll fees +/- 1 link
			{
				Name:     testhelpers.DestPool,
				Address:  ccipContracts.DestPool.Address(),
				Expected: testhelpers.MustSubBigInt(destBalances[testhelpers.DestPool], "30000000000000000297").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			},
			{
				Name:     testhelpers.TollOffRampRouter,
				Address:  ccipContracts.TollOffRamp.Address(),
				Expected: testhelpers.MustAddBigInt(destBalances[testhelpers.TollOffRampRouter], "2852678400000000000").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
				Within:   "1000000000000000000",
			}, // +/- 1 link
		})
		tollCurrentSeqNum += n
	})

	t.Run("upgrade contracts while transactions are pending", func(t *testing.T) {
		eventSignatures := ccip.GetTollEventSignatures()
		ccipContracts.DeployNewTollOnRamp()
		ccipContracts.DeployNewTollOffRamp()
		newConfigBlock := ccipContracts.DestChain.Blockchain().CurrentBlock().Number().Int64()

		// delete previous jobs, 1 commit and a toll exec & GE exec
		for _, node := range nodes {
			err = node.App.DeleteJob(context.Background(), 1)
			require.NoError(t, err)
			err = node.App.DeleteJob(context.Background(), 2)
			require.NoError(t, err)
			err = node.App.DeleteJob(context.Background(), 3)
			require.NoError(t, err)
		}
		// create updated jobs
		jobParams = ccipContracts.NewCCIPJobSpecParams(tokensPerFeeCoinPipeline, newConfigBlock)
		testhelpers.AddAllJobs(t, jobParams, ccipContracts, nodes)

		// keep sending a number of send requests all of which would be in pending state
		currentSeqNum := atomic.NewInt32(1) // start with 1 as it's a new onramp
		startSeq := 1
		ticker := time.NewTicker(1 * time.Second)
		ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
		defer cancel()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ticker.C:
					// Approve router to take source token.
					tokenAmount := big.NewInt(100)
					feeTokenAmount := new(big.Int).Mul(big.NewInt(20), big.NewInt(1e18))
					_, err = ccipContracts.SourceLinkToken.Approve(
						ccipContracts.SourceUser,
						ccipContracts.TollOnRampRouter.Address(),
						new(big.Int).Add(tokenAmount, feeTokenAmount))
					require.NoError(t, err)

					t.Logf("sending request for seqnum %d", currentSeqNum.Load())
					currentSeqNum.Inc()
					testhelpers.SendTollRequest(t, ccipContracts,
						"hey DON, execute for me",
						[]evm_2_any_toll_onramp_router.CommonEVMTokenAndAmount{{
							Token:  ccipContracts.SourceLinkToken.Address(),
							Amount: tokenAmount,
						}},
						evm_2_any_toll_onramp_router.CommonEVMTokenAndAmount{
							Token:  ccipContracts.SourceLinkToken.Address(),
							Amount: feeTokenAmount,
						},
						big.NewInt(300_000),
						ccipContracts.Receivers[0].Receiver.Address())
					ccipContracts.SourceChain.Commit()
					ccipContracts.DestChain.Commit()
				case <-ctx.Done():
					return
				}
			}
		}()

		// now enable the newly deployed on/offRamp
		ccipContracts.EnableTollOnRamp()
		ccipContracts.EnableTollOffRamp()
		// wait for all requests to get triggered
		wg.Wait()
		// verify if all seqNums were delivered
		endSeqNum := int(currentSeqNum.Load())
		for i := startSeq; i < endSeqNum; i++ {
			t.Logf("verifying seqnum %d", i)
			testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, eventSignatures, ccipContracts.TollOnRamp.Address(), nodes, i)
			testhelpers.EventuallyReportCommitted(t, ccipContracts, ccipContracts.TollOnRamp.Address(), i, i)
			executionLog := testhelpers.AllNodesHaveExecutedSeqNums(t, ccipContracts, eventSignatures, ccipContracts.TollOffRamp.Address(), nodes, i, i)
			testhelpers.AssertTollExecSuccess(t, ccipContracts, executionLog[0])
		}
		tollCurrentSeqNum = endSeqNum
	})
}
