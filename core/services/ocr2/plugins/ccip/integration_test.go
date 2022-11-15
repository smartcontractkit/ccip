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

	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"

	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_subscription_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_toll_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/subscription_sender_dapp"
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

	// Add the bootstrap job
	bootstrapNode.AddBootstrapJob(t, fmt.Sprintf(`
type               	= "bootstrap"
relay 				= "evm"
schemaVersion      	= 1
name               	= "boot"
contractID    	    = "%s"
contractConfigConfirmations = 1
contractConfigTrackerPollInterval = "1s"
[relayConfig]
chainID = %s
`, ccipContracts.CommitStore.Address(), destChainID))

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

	defer linkEth.Close()
	spec := ccipContracts.NewCCIPJobSpecParams(tokensPerFeeCoinPipeline)

	for i, node := range nodes {
		spec.AddCCIPCommitJob(t, fmt.Sprintf("ccip-commit-%d", i), node, configBlock)
		spec.AddCCIPTollExecutionJob(t, fmt.Sprintf("ccip-executor-toll-%d", i), node, configBlock)
		spec.AddCCIPSubExecutionJob(t, fmt.Sprintf("ccip-executor-subscription-%d", i), node, configBlock)
	}
	// Replay for bootstrap.
	bc, err := bootstrapNode.App.GetChains().EVM.Get(destChainID)
	require.NoError(t, err)
	require.NoError(t, bc.LogPoller().Replay(context.Background(), configBlock))

	tollCurrentSeqNum := 1
	subCurrentSeqNum := 1
	// Create src sub and dst sub, funded for 10 msgs with 100k callback and 1 token.
	// Needs to be sufficient to cover default gas price of 200gwei.
	// Costs ~7 link for 100k callback @ 200gwei.
	for _, receiver := range ccipContracts.Receivers {
		commitFee := big.NewInt(0).Mul(ccipContracts.SubOnRampFee, big.NewInt(10))
		_, err = ccipContracts.SourceLinkToken.Approve(ccipContracts.SourceUser, ccipContracts.SubOnRampRouter.Address(), commitFee)
		require.NoError(t, err)
		_, err = ccipContracts.SubOnRampRouter.FundSubscription(ccipContracts.SourceUser, commitFee)
		require.NoError(t, err)
		ccipContracts.SourceChain.Commit()
		// if using senderDapp
		_, err = ccipContracts.SourceLinkToken.Transfer(ccipContracts.SourceUser, ccipContracts.SubSenderApp.Address(), commitFee)
		require.NoError(t, err)
		_, err = ccipContracts.SourceLinkToken.Approve(ccipContracts.SourceUser, ccipContracts.SubSenderApp.Address(), commitFee)
		require.NoError(t, err)
		ccipContracts.SourceChain.Commit()
		_, err = ccipContracts.SubSenderApp.FundSubscription(ccipContracts.SourceUser, ccipContracts.SourceLinkToken.Address(), commitFee)
		require.NoError(t, err)
		ccipContracts.SourceChain.Commit()

		subscriptionBalance := big.NewInt(0).Mul(big.NewInt(80), big.NewInt(1e18))
		_, err = ccipContracts.DestLinkToken.Approve(ccipContracts.DestUser, ccipContracts.SubOffRampRouter.Address(), subscriptionBalance)
		require.NoError(t, err)
		_, err = ccipContracts.SubOffRampRouter.CreateSubscription(ccipContracts.DestUser, any_2_evm_subscription_offramp_router.SubscriptionInterfaceOffRampSubscription{
			Senders:          []common.Address{ccipContracts.SourceUser.From, ccipContracts.SubSenderApp.Address()},
			Receiver:         receiver.Receiver.Address(),
			StrictSequencing: receiver.Strict,
			Balance:          subscriptionBalance,
		})
		require.NoError(t, err)
	}

	ccipContracts.DestChain.Commit()

	t.Run("single auto-execute toll", func(t *testing.T) {
		// Approve router to take source token.
		tokenAmount := big.NewInt(100)
		// Example sending a msg with 100k callback execution on eth mainnet.
		// Gas price 200e9 wei/gas * 1e5 gas * (2e20 juels/eth / 1e18wei/eth)
		// 4e18 juels = 4 link, which does not include gas used outside the callback.
		// Gas outside the callback for 1 token is ~100k in the worst case.
		feeTokenAmount := big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18))
		_, err = ccipContracts.SourceLinkToken.Approve(ccipContracts.SourceUser, ccipContracts.TollOnRampRouter.Address(), big.NewInt(0).Add(tokenAmount, feeTokenAmount))
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

		testhelpers.SendRequest(t, ccipContracts, "hey DON, execute for me",
			[]evm_2_any_toll_onramp_router.CCIPEVMTokenAndAmount{{
				Token:  ccipContracts.SourceLinkToken.Address(),
				Amount: tokenAmount,
			}},
			evm_2_any_toll_onramp_router.CCIPEVMTokenAndAmount{
				Token:  ccipContracts.SourceLinkToken.Address(),
				Amount: feeTokenAmount,
			}, big.NewInt(100_000),
			ccipContracts.Receivers[0].Receiver.Address())
		testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, ccip.CCIPTollSendRequested, ccipContracts.TollOnRamp.Address(), nodes, tollCurrentSeqNum)
		testhelpers.EventuallyReportCommitted(t, ccipContracts, ccipContracts.TollOnRamp.Address(), tollCurrentSeqNum, tollCurrentSeqNum)
		executionLog := testhelpers.AllNodesHaveExecutedSeqNum(t, ccipContracts, ccipContracts.TollOffRamp.Address(), nodes, tollCurrentSeqNum)
		testhelpers.AssertTollExecSuccess(t, ccipContracts, executionLog)

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
				Getter:   ccipContracts.GetSourceLinkBalance},
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

	t.Run("single auto-execute subscription custom token", func(t *testing.T) {
		tokenAmount := big.NewInt(100)
		_, err = ccipContracts.SourceCustomToken.Approve(ccipContracts.SourceUser, ccipContracts.SubOnRampRouter.Address(), tokenAmount)
		require.NoError(t, err)
		ccipContracts.SourceChain.Commit()
		testhelpers.SendSubRequest(t, ccipContracts, "hey DON, execute for me", []evm_2_any_subscription_onramp_router.CCIPEVMTokenAndAmount{{
			Token:  ccipContracts.SourceCustomToken.Address(),
			Amount: tokenAmount,
		}}, big.NewInt(100_000), ccipContracts.Receivers[0].Receiver.Address())
		executionLog := testhelpers.AllNodesHaveExecutedSeqNum(t, ccipContracts, ccipContracts.SubOffRamp.Address(), nodes, subCurrentSeqNum)
		subCurrentSeqNum++
		testhelpers.AssertSubExecSuccess(t, ccipContracts, executionLog)
	})

	t.Run("single auto-execute subscription", func(t *testing.T) {
		tokenAmount := big.NewInt(100)
		_, err = ccipContracts.SourceLinkToken.Approve(ccipContracts.SourceUser, ccipContracts.SubOnRampRouter.Address(), tokenAmount)
		require.NoError(t, err)
		ccipContracts.SourceChain.Commit()

		sourceBalances, err := testhelpers.GetBalances([]testhelpers.BalanceReq{
			{Name: testhelpers.SourcePool, Addr: ccipContracts.SourcePool.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: testhelpers.SubOnRamp, Addr: ccipContracts.SubOnRamp.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: testhelpers.SubOnRampRouter, Addr: ccipContracts.SubOnRampRouter.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: testhelpers.SourceSub, Addr: ccipContracts.SourceUser.From, Getter: ccipContracts.GetSourceSubBalance},
		})
		require.NoError(t, err, "fetching source balance")
		destBalances, err := testhelpers.GetBalances([]testhelpers.BalanceReq{
			{Name: testhelpers.Receiver, Addr: ccipContracts.Receivers[0].Receiver.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: testhelpers.DestPool, Addr: ccipContracts.DestPool.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: testhelpers.SubOffRampRouter, Addr: ccipContracts.SubOffRampRouter.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: testhelpers.DestSub, Addr: ccipContracts.Receivers[0].Receiver.Address(), Getter: ccipContracts.GetDestSubBalance},
		})
		require.NoError(t, err, "fetching dest balance")

		testhelpers.SendSubRequest(t, ccipContracts, "hey DON, execute for me", []evm_2_any_subscription_onramp_router.CCIPEVMTokenAndAmount{{
			Token:  ccipContracts.SourceLinkToken.Address(),
			Amount: tokenAmount,
		}}, big.NewInt(100_000), ccipContracts.Receivers[0].Receiver.Address())
		testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, ccip.CCIPSubSendRequested, ccipContracts.SubOnRamp.Address(), nodes, subCurrentSeqNum)
		testhelpers.EventuallyReportCommitted(t, ccipContracts, ccipContracts.SubOnRamp.Address(), subCurrentSeqNum, subCurrentSeqNum)
		executionLog := testhelpers.AllNodesHaveExecutedSeqNum(t, ccipContracts, ccipContracts.SubOffRamp.Address(), nodes, subCurrentSeqNum)
		testhelpers.AssertSubExecSuccess(t, ccipContracts, executionLog)

		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     testhelpers.SourcePool,
				Address:  ccipContracts.SourcePool.Address(),
				Expected: testhelpers.MustAddBigInt(sourceBalances[testhelpers.SourcePool], "100").String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			}, // 100 transfer
			{
				Name: testhelpers.SubOnRamp, Address: ccipContracts.SubOnRamp.Address(),
				Expected: sourceBalances[testhelpers.SubOnRamp].String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			},
			{
				Name: testhelpers.SubOnRampRouter, Address: ccipContracts.SubOnRampRouter.Address(),
				Expected: sourceBalances[testhelpers.SubOnRampRouter].String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			}, // No change, internal account of fee to us.
			{
				Name:     testhelpers.SourceSub,
				Address:  ccipContracts.SourceUser.From,
				Expected: testhelpers.MustSubBigInt(sourceBalances[testhelpers.SourceSub], "1").String(),
				Getter:   ccipContracts.GetSourceSubBalance,
			}, // Pays 1 in fee
		})
		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     testhelpers.Receiver,
				Address:  ccipContracts.Receivers[0].Receiver.Address(),
				Expected: testhelpers.MustAddBigInt(destBalances[testhelpers.Receiver], "100").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // Full amount gets transferred
			{
				Name:     testhelpers.DestPool,
				Address:  ccipContracts.DestPool.Address(),
				Expected: testhelpers.MustSubBigInt(destBalances[testhelpers.DestPool], "100").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // We lose 100 link from the pool
			{
				Name:     testhelpers.SubOffRampRouter,
				Address:  ccipContracts.SubOffRampRouter.Address(),
				Expected: destBalances[testhelpers.SubOffRampRouter].String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // Gas reimbursement for nop
			{
				Name:     testhelpers.DestSub,
				Address:  ccipContracts.Receivers[0].Receiver.Address(),
				Expected: testhelpers.MustSubBigInt(destBalances[testhelpers.DestSub], "617786400000000000").String(),
				Getter:   ccipContracts.GetDestSubBalance,
				Within:   "100000000000000000",
			}, // Costs ~0.65 link. +/- 0.1
		})
		subCurrentSeqNum++
	})

	t.Run("single auto-execute subscription by senderDapp", func(t *testing.T) {
		tokenAmount := big.NewInt(100)

		// approve the token amount for senderdapp
		_, err = ccipContracts.SourceLinkToken.Approve(ccipContracts.SourceUser, ccipContracts.SubSenderApp.Address(), tokenAmount)
		require.NoError(t, err)
		ccipContracts.SourceChain.Commit()

		sourceBalances, err := testhelpers.GetBalances([]testhelpers.BalanceReq{
			{Name: testhelpers.SourcePool, Addr: ccipContracts.SourcePool.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: testhelpers.SubOnRamp, Addr: ccipContracts.SubOnRamp.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: testhelpers.SubOnRampRouter, Addr: ccipContracts.SubOnRampRouter.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: testhelpers.SourceSub, Addr: ccipContracts.SubSenderApp.Address(), Getter: ccipContracts.GetSourceSubBalance},
		})
		require.NoError(t, err, "fetching source balance")
		destBalances, err := testhelpers.GetBalances([]testhelpers.BalanceReq{
			{Name: testhelpers.Receiver, Addr: ccipContracts.Receivers[0].Receiver.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: testhelpers.DestPool, Addr: ccipContracts.DestPool.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: testhelpers.SubOffRampRouter, Addr: ccipContracts.SubOffRampRouter.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: testhelpers.DestSub, Addr: ccipContracts.Receivers[0].Receiver.Address(), Getter: ccipContracts.GetDestSubBalance},
		})
		require.NoError(t, err, "fetching dest balance")

		testhelpers.SendSubRequestByDapp(t, ccipContracts, "hey DON, execute for me", []subscription_sender_dapp.CCIPEVMTokenAndAmount{{
			Token:  ccipContracts.SourceLinkToken.Address(),
			Amount: tokenAmount,
		}}, big.NewInt(100_000), ccipContracts.Receivers[0].Receiver.Address())
		testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, ccip.CCIPSubSendRequested, ccipContracts.SubOnRamp.Address(), nodes, subCurrentSeqNum)
		testhelpers.EventuallyReportCommitted(t, ccipContracts, ccipContracts.SubOnRamp.Address(), subCurrentSeqNum, subCurrentSeqNum)
		executionLog := testhelpers.AllNodesHaveExecutedSeqNum(t, ccipContracts, ccipContracts.SubOffRamp.Address(), nodes, subCurrentSeqNum)
		testhelpers.AssertSubExecSuccess(t, ccipContracts, executionLog)

		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     testhelpers.SourcePool,
				Address:  ccipContracts.SourcePool.Address(),
				Expected: testhelpers.MustAddBigInt(sourceBalances[testhelpers.SourcePool], "100").String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			}, // 100 transfer
			{
				Name: testhelpers.SubOnRamp, Address: ccipContracts.SubOnRamp.Address(),
				Expected: sourceBalances[testhelpers.SubOnRamp].String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			},
			{
				Name: testhelpers.SubOnRampRouter, Address: ccipContracts.SubOnRampRouter.Address(),
				Expected: sourceBalances[testhelpers.SubOnRampRouter].String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			}, // No change, internal account of fee to us.
			{
				Name:     testhelpers.SourceSub,
				Address:  ccipContracts.SubSenderApp.Address(),
				Expected: testhelpers.MustSubBigInt(sourceBalances[testhelpers.SourceSub], "1").String(),
				Getter:   ccipContracts.GetSourceSubBalance,
			}, // Pays 1 in fee
		})
		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     testhelpers.Receiver,
				Address:  ccipContracts.Receivers[0].Receiver.Address(),
				Expected: testhelpers.MustAddBigInt(destBalances[testhelpers.Receiver], "100").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // Full amount gets transferred
			{
				Name:     testhelpers.DestPool,
				Address:  ccipContracts.DestPool.Address(),
				Expected: testhelpers.MustSubBigInt(destBalances[testhelpers.DestPool], "100").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // We lose 100 link from the pool
			{
				Name:     testhelpers.SubOffRampRouter,
				Address:  ccipContracts.SubOffRampRouter.Address(),
				Expected: destBalances[testhelpers.SubOffRampRouter].String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // Gas reimbursement for nop
			{
				Name:     testhelpers.DestSub,
				Address:  ccipContracts.Receivers[0].Receiver.Address(),
				Expected: testhelpers.MustSubBigInt(destBalances[testhelpers.DestSub], "627786400000000000").String(),
				Getter:   ccipContracts.GetDestSubBalance,
				Within:   "100000000000000000",
			}, // Costs ~0.65 link. +/- 0.1
		})
		subCurrentSeqNum++
	})

	t.Run("batch auto-execute toll", func(t *testing.T) {
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
		feeTokenAmount := big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18))
		var txs []*gethtypes.Transaction
		n := 3
		for i := 0; i < n; i++ {
			_, err = ccipContracts.SourceLinkToken.Approve(ccipContracts.SourceUser, ccipContracts.TollOnRampRouter.Address(), big.NewInt(0).Add(tokenAmount, feeTokenAmount))
			require.NoError(t, err)
			txs = append(txs, testhelpers.QueueRequest(t, ccipContracts, fmt.Sprintf("batch request %d", tollCurrentSeqNum+i), []evm_2_any_toll_onramp_router.CCIPEVMTokenAndAmount{{
				Token:  ccipContracts.SourceLinkToken.Address(),
				Amount: tokenAmount,
			}}, evm_2_any_toll_onramp_router.CCIPEVMTokenAndAmount{
				Token:  ccipContracts.SourceLinkToken.Address(),
				Amount: feeTokenAmount,
			}, big.NewInt(100_000), ccipContracts.Receivers[0].Receiver.Address()))
		}
		// Send a batch of requests in a single block
		testhelpers.ConfirmTxs(t, txs, ccipContracts.SourceChain)
		// All nodes should have all 3.
		var reqs []logpoller.Log
		for i := 0; i < n; i++ {
			reqs = append(reqs, testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, ccip.CCIPTollSendRequested, ccipContracts.TollOnRamp.Address(), nodes, tollCurrentSeqNum+i))
		}
		// Should see a report with the full range
		testhelpers.EventuallyReportCommitted(t, ccipContracts, ccipContracts.TollOnRamp.Address(), tollCurrentSeqNum, tollCurrentSeqNum+n-1)
		// Should all be executed
		for i := range reqs {
			executionLog := testhelpers.AllNodesHaveExecutedSeqNum(t, ccipContracts, ccipContracts.TollOffRamp.Address(), nodes, tollCurrentSeqNum+i)
			testhelpers.AssertTollExecSuccess(t, ccipContracts, executionLog)
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

	t.Run("batch auto-execute subscription", func(t *testing.T) {
		sourceBalances, err := testhelpers.GetBalances([]testhelpers.BalanceReq{
			{Name: testhelpers.SourcePool, Addr: ccipContracts.SourcePool.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: testhelpers.SourceSub, Addr: ccipContracts.SourceUser.From, Getter: ccipContracts.GetSourceSubBalance},
		})
		require.NoError(t, err, "fetching source balance")
		destBalances, err := testhelpers.GetBalances([]testhelpers.BalanceReq{
			{Name: testhelpers.Receiver, Addr: ccipContracts.Receivers[0].Receiver.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: testhelpers.DestPool, Addr: ccipContracts.DestPool.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: testhelpers.DestSub, Addr: ccipContracts.Receivers[0].Receiver.Address(), Getter: ccipContracts.GetDestSubBalance},
		})
		require.NoError(t, err, "fetching dest balance")

		tokenAmount := big.NewInt(100)
		var txs []*gethtypes.Transaction
		n := 3
		for i := 0; i < n; i++ {
			_, err = ccipContracts.SourceLinkToken.Approve(ccipContracts.SourceUser, ccipContracts.SubOnRampRouter.Address(), tokenAmount)
			require.NoError(t, err)
			txs = append(txs, testhelpers.QueueSubRequest(t, ccipContracts, "hey DON, execute for me", []evm_2_any_subscription_onramp_router.CCIPEVMTokenAndAmount{{
				Token:  ccipContracts.SourceLinkToken.Address(),
				Amount: tokenAmount,
			}}, big.NewInt(100_000), ccipContracts.Receivers[0].Receiver.Address()))
		}
		ccipContracts.SourceChain.Commit()
		// Send a batch of requests in a single block
		testhelpers.ConfirmTxs(t, txs, ccipContracts.SourceChain)
		var reqs []logpoller.Log
		for i := 0; i < n; i++ {
			reqs = append(reqs, testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, ccip.CCIPSubSendRequested, ccipContracts.SubOnRamp.Address(), nodes, subCurrentSeqNum+i))
		}
		// Should see a report with the full range
		testhelpers.EventuallyReportCommitted(t, ccipContracts, ccipContracts.SubOnRamp.Address(), subCurrentSeqNum, subCurrentSeqNum+n-1)
		// Should all be executed
		for i := range reqs {
			executionLog := testhelpers.AllNodesHaveExecutedSeqNum(t, ccipContracts, ccipContracts.SubOffRamp.Address(), nodes, subCurrentSeqNum+i)
			testhelpers.AssertSubExecSuccess(t, ccipContracts, executionLog)
		}
		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     testhelpers.SourcePool,
				Address:  ccipContracts.SourcePool.Address(),
				Expected: testhelpers.MustAddBigInt(sourceBalances[testhelpers.SourcePool], "300").String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			}, // 100 transfer
			{
				Name:     testhelpers.SourceSub,
				Address:  ccipContracts.SourceUser.From,
				Expected: testhelpers.MustSubBigInt(sourceBalances[testhelpers.SourceSub], "3").String(),
				Getter:   ccipContracts.GetSourceSubBalance,
			}, // Pays 1 in fee
		})
		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     testhelpers.Receiver,
				Address:  ccipContracts.Receivers[0].Receiver.Address(),
				Expected: testhelpers.MustAddBigInt(destBalances[testhelpers.Receiver], "300").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // Full amount gets transferred
			{
				Name:     testhelpers.DestPool,
				Address:  ccipContracts.DestPool.Address(),
				Expected: testhelpers.MustSubBigInt(destBalances[testhelpers.DestPool], "300").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // We lose 100 link from the pool
			{
				Name:     testhelpers.DestSub,
				Address:  ccipContracts.Receivers[0].Receiver.Address(),
				Expected: testhelpers.MustSubBigInt(destBalances[testhelpers.DestSub], "1864160000000000000").String(),
				Getter:   ccipContracts.GetDestSubBalance,
				Within:   "1000000000000000000"}, // Costs ~0.65 link. Varies slightly due to variable calldata encoding gas costs.
		})
		subCurrentSeqNum += n
	})

	t.Run("single strict sequencing auto-execute subscription", func(t *testing.T) {
		sourceBalances, err := testhelpers.GetBalances([]testhelpers.BalanceReq{
			{Name: testhelpers.SourcePool, Addr: ccipContracts.SourcePool.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: testhelpers.SourceSub, Addr: ccipContracts.SourceUser.From, Getter: ccipContracts.GetSourceSubBalance},
		})
		require.NoError(t, err, "fetching source balance")
		destBalances, err := testhelpers.GetBalances([]testhelpers.BalanceReq{
			{Name: testhelpers.Receiver, Addr: ccipContracts.Receivers[1].Receiver.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: testhelpers.DestPool, Addr: ccipContracts.DestPool.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: testhelpers.DestSub, Addr: ccipContracts.Receivers[1].Receiver.Address(), Getter: ccipContracts.GetDestSubBalance},
		})
		require.NoError(t, err, "fetching dest balance")
		tokenAmount := big.NewInt(100)
		_, err = ccipContracts.SourceLinkToken.Approve(ccipContracts.SourceUser, ccipContracts.SubOnRampRouter.Address(), tokenAmount)
		require.NoError(t, err)
		ccipContracts.SourceChain.Commit()
		testhelpers.SendSubRequest(t, ccipContracts, "hey DON, execute for me", []evm_2_any_subscription_onramp_router.CCIPEVMTokenAndAmount{{
			Token:  ccipContracts.SourceLinkToken.Address(),
			Amount: tokenAmount,
		}}, big.NewInt(100_000), ccipContracts.Receivers[1].Receiver.Address())
		testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, ccip.CCIPSubSendRequested, ccipContracts.SubOnRamp.Address(), nodes, subCurrentSeqNum)
		testhelpers.EventuallyReportCommitted(t, ccipContracts, ccipContracts.SubOnRamp.Address(), subCurrentSeqNum, subCurrentSeqNum)
		executionLog := testhelpers.AllNodesHaveExecutedSeqNum(t, ccipContracts, ccipContracts.SubOffRamp.Address(), nodes, subCurrentSeqNum)
		testhelpers.AssertSubExecSuccess(t, ccipContracts, executionLog)
		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     testhelpers.SourcePool,
				Address:  ccipContracts.SourcePool.Address(),
				Expected: testhelpers.MustAddBigInt(sourceBalances[testhelpers.SourcePool], "100").String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			}, // 100 transfer
			{
				Name:     testhelpers.SourceSub,
				Address:  ccipContracts.SourceUser.From,
				Expected: testhelpers.MustSubBigInt(sourceBalances[testhelpers.SourceSub], "1").String(),
				Getter:   ccipContracts.GetSourceSubBalance,
			}, // Pays 1 in fee
		})
		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     testhelpers.Receiver,
				Address:  ccipContracts.Receivers[1].Receiver.Address(),
				Expected: testhelpers.MustAddBigInt(destBalances[testhelpers.Receiver], "100").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // Full amount gets transferred
			{
				Name:     testhelpers.DestPool,
				Address:  ccipContracts.DestPool.Address(),
				Expected: testhelpers.MustSubBigInt(destBalances[testhelpers.DestPool], "100").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // We lose 100 link from the pool
			{
				Name:     testhelpers.DestSub,
				Address:  ccipContracts.Receivers[1].Receiver.Address(),
				Expected: testhelpers.MustSubBigInt(destBalances[testhelpers.DestSub], "654720000000000000").String(),
				Getter:   ccipContracts.GetDestSubBalance,
				Within:   "100000000000000000",
			}, // Costs ~0.65 link. Varies slightly due to variable calldata encoding gas costs.
		})
		subCurrentSeqNum++
	})

	t.Run("strict sequencing - subsequent messages are not executed until reverted message is manually executed", func(t *testing.T) {
		// token amounts for diff requests
		tokenAmounts := []*big.Int{
			big.NewInt(40),
			big.NewInt(10),
			big.NewInt(20),
			big.NewInt(30),
		}
		totalTokenAmount := big.NewInt(100) // total amount sent in this scenario
		strTotalAmount := totalTokenAmount.String()
		sourceBalances, err := testhelpers.GetBalances([]testhelpers.BalanceReq{
			{Name: testhelpers.SourcePool, Addr: ccipContracts.SourcePool.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: testhelpers.Sender, Addr: ccipContracts.SourceUser.From, Getter: ccipContracts.GetSourceLinkBalance},
			{Name: testhelpers.SourceSub, Addr: ccipContracts.SourceUser.From, Getter: ccipContracts.GetSourceSubBalance},
		})
		require.NoError(t, err, "fetching source balance")
		destBalances, err := testhelpers.GetBalances([]testhelpers.BalanceReq{
			{Name: testhelpers.Receiver, Addr: ccipContracts.Receivers[1].Receiver.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: testhelpers.DestPool, Addr: ccipContracts.DestPool.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: testhelpers.DestSub, Addr: ccipContracts.Receivers[1].Receiver.Address(), Getter: ccipContracts.GetDestSubBalance},
		})
		require.NoError(t, err, "fetching dest balance")

		// approve the total amount to be sent
		_, err = ccipContracts.SourceLinkToken.Approve(ccipContracts.SourceUser, ccipContracts.SubOnRampRouter.Address(), totalTokenAmount)
		require.NoError(t, err, "approving link token for onramp router")
		ccipContracts.SourceChain.Commit()

		// set revert to true so that the execution gets reverted
		_, err = ccipContracts.Receivers[1].Receiver.SetRevert(ccipContracts.DestUser, true)
		require.NoError(t, err, "setting revert to true on the receiver")
		ccipContracts.DestChain.Commit()
		currentBlockNumber := ccipContracts.DestChain.Blockchain().CurrentBlock().Number().Uint64()

		// attempt sending a request and the execution should be reverted
		testhelpers.SendSubRequest(t, ccipContracts, "hey DON, execute for me", []evm_2_any_subscription_onramp_router.CCIPEVMTokenAndAmount{{
			Token:  ccipContracts.SourceLinkToken.Address(),
			Amount: tokenAmounts[0],
		}}, big.NewInt(100_000), ccipContracts.Receivers[1].Receiver.Address())
		failedReq := testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, ccip.CCIPSubSendRequested, ccipContracts.SubOnRamp.Address(), nodes, subCurrentSeqNum)
		testhelpers.EventuallyReportCommitted(t, ccipContracts, ccipContracts.SubOnRamp.Address(), subCurrentSeqNum, subCurrentSeqNum)
		reportForFailedReq := testhelpers.EventuallyCommitReportAccepted(t, ccipContracts, currentBlockNumber)
		executionLog := testhelpers.AllNodesHaveExecutedSeqNum(t, ccipContracts, ccipContracts.SubOffRamp.Address(), nodes, subCurrentSeqNum)

		// the transaction should get reverted and the execution status should be failed
		testhelpers.AssertSubExecFailure(t, ccipContracts, executionLog)
		subCurrentSeqNum++

		// flip the revert settings on receiver
		_, err = ccipContracts.Receivers[1].Receiver.SetRevert(ccipContracts.DestUser, false)
		require.NoError(t, err, "setting revert to false on the receiver")
		ccipContracts.DestChain.Commit()
		ccipContracts.SourceChain.Commit()

		// send a bunch of subsequent ones which should not be executed
		var pendingReqNumbers []int
		for i := 1; i <= 3; i++ {
			testhelpers.SendSubRequest(t, ccipContracts, "hey DON, execute for me", []evm_2_any_subscription_onramp_router.CCIPEVMTokenAndAmount{{
				Token:  ccipContracts.SourceLinkToken.Address(),
				Amount: tokenAmounts[i],
			}}, big.NewInt(100_000), ccipContracts.Receivers[1].Receiver.Address())
			testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, ccip.CCIPSubSendRequested, ccipContracts.SubOnRamp.Address(), nodes, subCurrentSeqNum)
			testhelpers.EventuallyReportCommitted(t, ccipContracts, ccipContracts.SubOnRamp.Address(), subCurrentSeqNum, subCurrentSeqNum)
			executionLog := testhelpers.NoNodesHaveExecutedSeqNum(t, ccipContracts, ccipContracts.SubOffRamp.Address(), nodes, subCurrentSeqNum)
			require.Empty(t, executionLog)
			pendingReqNumbers = append(pendingReqNumbers, subCurrentSeqNum)
			subCurrentSeqNum++
		}

		// manually execute the failed request
		ccipContracts.DestChain.Commit()
		ccipContracts.SourceChain.Commit()
		currentBlockNumber = ccipContracts.DestChain.Blockchain().CurrentBlock().Number().Uint64()
		failedSeqNum := testhelpers.ExecuteSubMessage(t, ccipContracts, failedReq, []logpoller.Log{failedReq}, reportForFailedReq)
		testhelpers.EventuallyExecutionStateChangedToSuccess(t, ccipContracts, []uint64{failedSeqNum}, currentBlockNumber)

		// verify all the pending requests should be successfully executed now
		for _, seqNo := range pendingReqNumbers {
			t.Logf("Verify execution for pending seq Number %d", seqNo)
			testhelpers.EventuallyExecutionStateChangedToSuccess(t, ccipContracts, []uint64{uint64(seqNo)}, currentBlockNumber)
		}

		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     testhelpers.SourcePool,
				Address:  ccipContracts.SourcePool.Address(),
				Expected: testhelpers.MustAddBigInt(sourceBalances[testhelpers.SourcePool], strTotalAmount).String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			}, // 100 transfer
			{
				Name:     testhelpers.SourceSub,
				Address:  ccipContracts.SourceUser.From,
				Expected: testhelpers.MustSubBigInt(sourceBalances[testhelpers.SourceSub], "4").String(),
				Getter:   ccipContracts.GetSourceSubBalance,
			}, // Pays 4 in fee
			{
				Name:     testhelpers.Sender,
				Address:  ccipContracts.SourceUser.From,
				Expected: testhelpers.MustSubBigInt(sourceBalances[testhelpers.Sender], strTotalAmount).String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			}, // 100 transfer
		})
		t.Logf("destBalances[DestSub] %v", destBalances[testhelpers.DestSub])
		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     testhelpers.Receiver,
				Address:  ccipContracts.Receivers[1].Receiver.Address(),
				Expected: testhelpers.MustAddBigInt(destBalances[testhelpers.Receiver], strTotalAmount).String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // Full amount gets transferred
			{
				Name:     testhelpers.DestPool,
				Address:  ccipContracts.DestPool.Address(),
				Expected: testhelpers.MustSubBigInt(destBalances[testhelpers.DestPool], strTotalAmount).String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // We lose 100 link from the pool

			{
				Name:     testhelpers.DestSub,
				Address:  ccipContracts.Receivers[1].Receiver.Address(),
				Expected: testhelpers.MustSubBigInt(destBalances[testhelpers.DestSub], "2371120000000000000").String(),
				Getter:   ccipContracts.GetDestSubBalance,
				Within:   "100000000000000000",
			}, // Costs ~0.77 link per transfer for 3 auto req. Varies slightly due to variable calldata encoding gas costs.
		})
	})

	t.Run("upgrade contracts while transactions are pending", func(t *testing.T) {
		ccipContracts.DeployNewTollOnRamp()
		ccipContracts.DeployNewTollOffRamp()
		newConfigBlock := ccipContracts.DestChain.Blockchain().CurrentBlock().Number().Int64()
		// create updated jobs
		spec := ccipContracts.NewCCIPJobSpecParams(tokensPerFeeCoinPipeline)
		for i, node := range nodes {
			err = node.App.DeleteJob(context.Background(), 1)
			require.NoError(t, err)
			err = node.App.DeleteJob(context.Background(), 2)
			require.NoError(t, err)
			spec.AddCCIPCommitJob(t, fmt.Sprintf("ccip-commit-new-%d", i), node, newConfigBlock)
			spec.AddCCIPTollExecutionJob(t, fmt.Sprintf("ccip-executor-toll-new-%d", i), node, configBlock)
		}
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
					feeTokenAmount := big.NewInt(0).Mul(big.NewInt(20), big.NewInt(1e18))
					_, err = ccipContracts.SourceLinkToken.Approve(
						ccipContracts.SourceUser,
						ccipContracts.TollOnRampRouter.Address(),
						big.NewInt(0).Add(tokenAmount, feeTokenAmount))
					require.NoError(t, err)

					t.Logf("sending request for seqnum %d", currentSeqNum.Load())
					currentSeqNum.Inc()
					testhelpers.SendRequest(t, ccipContracts,
						"hey DON, execute for me",
						[]evm_2_any_toll_onramp_router.CCIPEVMTokenAndAmount{{
							Token:  ccipContracts.SourceLinkToken.Address(),
							Amount: tokenAmount,
						}},
						evm_2_any_toll_onramp_router.CCIPEVMTokenAndAmount{
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
			testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, ccip.CCIPTollSendRequested,
				ccipContracts.TollOnRamp.Address(), nodes, i)
			testhelpers.EventuallyReportCommitted(t, ccipContracts, ccipContracts.TollOnRamp.Address(), i, i)
			executionLog := testhelpers.AllNodesHaveExecutedSeqNum(t, ccipContracts,
				ccipContracts.TollOffRamp.Address(), nodes, i)
			testhelpers.AssertTollExecSuccess(t, ccipContracts, executionLog)
		}
		tollCurrentSeqNum = endSeqNum
	})
}
