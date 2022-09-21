package ccip_test

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp_router"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/testhelpers"
)

var (
	sourceChainID = big.NewInt(1000)
	destChainID   = big.NewInt(1337)
)

func mustAddBigInt(a *big.Int, b string) *big.Int {
	bi, _ := big.NewInt(0).SetString(b, 10)
	return big.NewInt(0).Add(a, bi)
}

func mustSubBigInt(a *big.Int, b string) *big.Int {
	bi, _ := big.NewInt(0).SetString(b, 10)
	return big.NewInt(0).Sub(a, bi)
}

var (
	// Source
	SourcePool       = "source pool"
	SourceSub        = "source sub"
	TollOnRampRouter = "toll onramp router"
	TollOnRamp       = "toll onramp"
	SubOnRamp        = "sub onramp"
	SubOnRampRouter  = "sub onramp router"

	// Dest
	TollOffRampRouter = "toll offramp router"
	TollOffRamp       = "toll offramp"
	SubOffRampRouter  = "sub offramp router"
	DestPool          = "dest pool"
	DestSub           = "dest sub"
	Receiver          = "receiver"
	Sender            = "sender"
)

func TestIntegration_CCIP(t *testing.T) {
	ccipContracts := testhelpers.SetupCCIPContracts(t, sourceChainID, destChainID)
	bootstrapNodePort := int64(19599)
	ctx := context.Background()
	// Starts nodes and configures them in the OCR contracts.
	bootstrapNode, nodes, configBlock := testhelpers.SetupAndStartNodes(ctx, t, ccipContracts, bootstrapNodePort)

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
`, ccipContracts.BlobVerifier.Address(), destChainID))

	linkEth := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`{"JuelsPerETH": "200000000000000000000"}`))
		require.NoError(t, err)
	}))
	defer linkEth.Close()
	// For each node add a relayer and executor job.
	for i, node := range nodes {
		node.AddJob(t, fmt.Sprintf(`
type                = "offchainreporting2"
pluginType          = "ccip-relay"
relay               = "evm"
schemaVersion      	= 1
name               	= "ccip-relay-%d"
contractID 			= "%s"
ocrKeyBundleID      = "%s"
transmitterID 		= "%s"
contractConfigConfirmations = 1
contractConfigTrackerPollInterval = "1s"

[pluginConfig]
onRampIDs            = ["%s", "%s"]
sourceChainID       = %s
destChainID         = %s
pollPeriod          = "1s"
destStartBlock      = %d

[relayConfig]
chainID             = "%s"

`, i, ccipContracts.BlobVerifier.Address(), node.KeyBundle.ID(), node.Transmitter, ccipContracts.TollOnRamp.Address(), ccipContracts.SubOnRamp.Address(), sourceChainID, destChainID, configBlock, destChainID))
		node.AddJob(t, fmt.Sprintf(`
type                = "offchainreporting2"
pluginType          = "ccip-execution"
relay               = "evm"
schemaVersion       = 1
name                = "ccip-executor-toll-%d"
contractID          = "%s"
ocrKeyBundleID      = "%s"
transmitterID       = "%s"
contractConfigConfirmations = 1
contractConfigTrackerPollInterval = "1s"

[pluginConfig]
onRampID            = "%s"
blobVerifierID      = "%s"
sourceChainID       = %s
destChainID         = %s
pollPeriod          = "1s"
destStartBlock      = %d
tokensPerFeeCoinPipeline = """
// Price 1 
link [type=http method=GET url="%s"];
link_parse [type=jsonparse path="JuelsPerETH"];
link->link_parse;
merge [type=merge left="{}" right="{\\\"%s\\\":$(link_parse)}"];
"""

[relayConfig]
chainID             = "%s"

`, i, ccipContracts.TollOffRamp.Address(), node.KeyBundle.ID(), node.Transmitter, ccipContracts.TollOnRamp.Address(), ccipContracts.BlobVerifier.Address(), sourceChainID, destChainID, configBlock, linkEth.URL, ccipContracts.DestLinkToken.Address(), destChainID))
		node.AddJob(t, fmt.Sprintf(`
type                = "offchainreporting2"
pluginType          = "ccip-execution"
relay               = "evm"
schemaVersion       = 1
name                = "ccip-executor-subscription-%d"
contractID 			= "%s"
ocrKeyBundleID      = "%s"
transmitterID       = "%s"
contractConfigConfirmations = 1
contractConfigTrackerPollInterval = "1s"

[pluginConfig]
onRampID            = "%s"
blobVerifierID      = "%s"
sourceChainID       = %s
destChainID         = %s
pollPeriod          = "1s"
destStartBlock      = %d
tokensPerFeeCoinPipeline = """
link [type=http method=GET url="%s"];
link_parse [type=jsonparse path="JuelsPerETH"];
link->link_parse;
merge [type=merge left="{}" right="{\\\"%s\\\":$(link_parse)}"];
"""

[relayConfig]
chainID             = "%s"

`, i, ccipContracts.SubOffRamp.Address(), node.KeyBundle.ID(), node.Transmitter, ccipContracts.SubOnRamp.Address(), ccipContracts.BlobVerifier.Address(), sourceChainID, destChainID, configBlock, linkEth.URL, ccipContracts.DestLinkToken.Address(), destChainID))
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
		relayFee := big.NewInt(0).Mul(ccipContracts.SubOnRampFee, big.NewInt(10))
		_, err = ccipContracts.SourceLinkToken.Approve(ccipContracts.SourceUser, ccipContracts.SubOnRampRouter.Address(), relayFee)
		require.NoError(t, err)
		_, err = ccipContracts.SubOnRampRouter.FundSubscription(ccipContracts.SourceUser, relayFee)
		require.NoError(t, err)
		ccipContracts.SourceChain.Commit()
		subscriptionBalance := big.NewInt(0).Mul(big.NewInt(80), big.NewInt(1e18))
		_, err = ccipContracts.DestLinkToken.Approve(ccipContracts.DestUser, ccipContracts.SubOffRampRouter.Address(), subscriptionBalance)
		require.NoError(t, err)
		_, err = ccipContracts.SubOffRampRouter.CreateSubscription(ccipContracts.DestUser, any_2_evm_subscription_offramp_router.SubscriptionInterfaceOffRampSubscription{
			Senders:          []common.Address{ccipContracts.SourceUser.From},
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

		sourceBalances := ccipContracts.GetBalances([]testhelpers.BalanceReq{
			{Name: SourcePool, Addr: ccipContracts.SourcePool.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: TollOnRamp, Addr: ccipContracts.TollOnRamp.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: TollOnRampRouter, Addr: ccipContracts.TollOnRampRouter.Address(), Getter: ccipContracts.GetSourceLinkBalance},
		})
		destBalances := ccipContracts.GetBalances([]testhelpers.BalanceReq{
			{Name: Receiver, Addr: ccipContracts.Receivers[0].Receiver.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: DestPool, Addr: ccipContracts.DestPool.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: TollOffRamp, Addr: ccipContracts.TollOffRamp.Address(), Getter: ccipContracts.GetDestLinkBalance},
		})

		testhelpers.SendRequest(t, ccipContracts, "hey DON, execute for me",
			[]common.Address{ccipContracts.SourceLinkToken.Address()},
			[]*big.Int{tokenAmount}, ccipContracts.SourceLinkToken.Address(),
			feeTokenAmount,
			big.NewInt(100_000))
		testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, ccip.CCIPTollSendRequested, ccipContracts.TollOnRamp.Address(), nodes, tollCurrentSeqNum)
		testhelpers.EventuallyReportRelayed(t, ccipContracts, ccipContracts.TollOnRamp.Address(), tollCurrentSeqNum, tollCurrentSeqNum)
		executionLog := testhelpers.AllNodesHaveExecutedSeqNum(t, ccipContracts, ccipContracts.TollOffRamp.Address(), nodes, tollCurrentSeqNum)
		testhelpers.AssertTollExecSuccess(t, ccipContracts, executionLog)

		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     SourcePool,
				Address:  ccipContracts.SourcePool.Address(),
				Expected: mustAddBigInt(sourceBalances[SourcePool], "10000000000000000099").String(), // 10e18 + 100 transfer - 1 fee
				Getter:   ccipContracts.GetSourceLinkBalance,
				Within:   ""},
			{
				Name:     TollOnRamp,
				Address:  ccipContracts.TollOnRamp.Address(),
				Expected: sourceBalances[TollOnRamp].String(),
				Getter:   ccipContracts.GetSourceLinkBalance},
			{
				Name:     TollOnRampRouter,
				Address:  ccipContracts.TollOnRampRouter.Address(),
				Expected: mustAddBigInt(sourceBalances[TollOnRampRouter], "1").String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			},
		})
		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     Receiver,
				Address:  ccipContracts.Receivers[0].Receiver.Address(),
				Expected: mustAddBigInt(destBalances[Receiver], "9049107200000000099").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
				Within:   "1000000000000000000"}, // Roughly 200k gas * 200e9 wei/gas * (2e20 link/eth / 1e18wei/eth)
			{
				Name:     DestPool,
				Address:  ccipContracts.DestPool.Address(),
				Expected: mustSubBigInt(destBalances[DestPool], "10000000000000000099").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // We lose 10 link from the pool
			{
				Name:     TollOffRamp,
				Address:  ccipContracts.TollOffRamp.Address(),
				Expected: mustAddBigInt(destBalances[TollOffRamp], "950954400000000000").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			},
		})
		tollCurrentSeqNum++
	})

	t.Run("single auto-execute subscription", func(t *testing.T) {
		tokenAmount := big.NewInt(100)
		_, err = ccipContracts.SourceLinkToken.Approve(ccipContracts.SourceUser, ccipContracts.SubOnRampRouter.Address(), tokenAmount)
		require.NoError(t, err)
		ccipContracts.SourceChain.Commit()

		sourceBalances := ccipContracts.GetBalances([]testhelpers.BalanceReq{
			{Name: SourcePool, Addr: ccipContracts.SourcePool.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: SubOnRamp, Addr: ccipContracts.SubOnRamp.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: SubOnRampRouter, Addr: ccipContracts.SubOnRampRouter.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: SourceSub, Addr: ccipContracts.SourceUser.From, Getter: ccipContracts.GetSourceSubBalance},
		})
		destBalances := ccipContracts.GetBalances([]testhelpers.BalanceReq{
			{Name: Receiver, Addr: ccipContracts.Receivers[0].Receiver.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: DestPool, Addr: ccipContracts.DestPool.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: SubOffRampRouter, Addr: ccipContracts.SubOffRampRouter.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: DestSub, Addr: ccipContracts.Receivers[0].Receiver.Address(), Getter: ccipContracts.GetDestSubBalance},
		})
		testhelpers.SendSubRequest(t, ccipContracts, "hey DON, execute for me", []common.Address{ccipContracts.SourceLinkToken.Address()},
			[]*big.Int{tokenAmount}, big.NewInt(100_000), ccipContracts.Receivers[0].Receiver.Address())
		testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, ccip.CCIPSubSendRequested, ccipContracts.SubOnRamp.Address(), nodes, subCurrentSeqNum)
		testhelpers.EventuallyReportRelayed(t, ccipContracts, ccipContracts.SubOnRamp.Address(), subCurrentSeqNum, subCurrentSeqNum)
		executionLog := testhelpers.AllNodesHaveExecutedSeqNum(t, ccipContracts, ccipContracts.SubOffRamp.Address(), nodes, subCurrentSeqNum)
		testhelpers.AssertSubExecSuccess(t, ccipContracts, executionLog)

		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     SourcePool,
				Address:  ccipContracts.SourcePool.Address(),
				Expected: mustAddBigInt(sourceBalances[SourcePool], "100").String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			}, // 100 transfer
			{
				Name: SubOnRamp, Address: ccipContracts.SubOnRamp.Address(),
				Expected: sourceBalances[SubOnRamp].String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			},
			{
				Name: SubOnRampRouter, Address: ccipContracts.SubOnRampRouter.Address(),
				Expected: sourceBalances[SubOnRampRouter].String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			}, // No change, internal account of fee to us.
			{
				Name:     SourceSub,
				Address:  ccipContracts.SourceUser.From,
				Expected: mustSubBigInt(sourceBalances[SourceSub], "1").String(),
				Getter:   ccipContracts.GetSourceSubBalance,
			}, // Pays 1 in fee
		})
		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     Receiver,
				Address:  ccipContracts.Receivers[0].Receiver.Address(),
				Expected: mustAddBigInt(destBalances[Receiver], "100").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // Full amount gets transferred
			{
				Name:     DestPool,
				Address:  ccipContracts.DestPool.Address(),
				Expected: mustSubBigInt(destBalances[DestPool], "100").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // We lose 100 link from the pool
			{
				Name:     SubOffRampRouter,
				Address:  ccipContracts.SubOffRampRouter.Address(),
				Expected: destBalances[SubOffRampRouter].String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // Gas reimbursement for nop
			{
				Name:     DestSub,
				Address:  ccipContracts.Receivers[0].Receiver.Address(),
				Expected: mustSubBigInt(destBalances[DestSub], "617786400000000000").String(),
				Getter:   ccipContracts.GetDestSubBalance,
				Within:   "100000000000000000",
			}, // Costs ~0.65 link. +/- 0.1
		})
		subCurrentSeqNum++
	})

	t.Run("batch auto-execute toll", func(t *testing.T) {
		sourceBalances := ccipContracts.GetBalances([]testhelpers.BalanceReq{
			{Name: SourcePool, Addr: ccipContracts.SourcePool.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: TollOnRampRouter, Addr: ccipContracts.TollOnRampRouter.Address(), Getter: ccipContracts.GetSourceLinkBalance},
		})
		destBalances := ccipContracts.GetBalances([]testhelpers.BalanceReq{
			{Name: Receiver, Addr: ccipContracts.Receivers[0].Receiver.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: DestPool, Addr: ccipContracts.DestPool.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: TollOffRampRouter, Addr: ccipContracts.TollOffRampRouter.Address(), Getter: ccipContracts.GetDestLinkBalance},
		})
		tokenAmount := big.NewInt(100)
		feeTokenAmount := big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18))
		var txs []*gethtypes.Transaction
		n := 3
		for i := 0; i < n; i++ {
			_, err = ccipContracts.SourceLinkToken.Approve(ccipContracts.SourceUser, ccipContracts.TollOnRampRouter.Address(), big.NewInt(0).Add(tokenAmount, feeTokenAmount))
			require.NoError(t, err)
			txs = append(txs, testhelpers.QueueRequest(t, ccipContracts, fmt.Sprintf("batch request %d", tollCurrentSeqNum+i), []common.Address{ccipContracts.SourceLinkToken.Address()}, []*big.Int{tokenAmount},
				ccipContracts.SourceLinkToken.Address(), feeTokenAmount, big.NewInt(100_000)))
		}
		// Send a batch of requests in a single block
		testhelpers.ConfirmTxs(t, txs, ccipContracts.SourceChain)
		// All nodes should have all 3.
		var reqs []logpoller.Log
		for i := 0; i < n; i++ {
			reqs = append(reqs, testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, ccip.CCIPTollSendRequested, ccipContracts.TollOnRamp.Address(), nodes, tollCurrentSeqNum+i))
		}
		// Should see a report with the full range
		testhelpers.EventuallyReportRelayed(t, ccipContracts, ccipContracts.TollOnRamp.Address(), tollCurrentSeqNum, tollCurrentSeqNum+n-1)
		// Should all be executed
		for i := range reqs {
			executionLog := testhelpers.AllNodesHaveExecutedSeqNum(t, ccipContracts, ccipContracts.TollOffRamp.Address(), nodes, tollCurrentSeqNum+i)
			testhelpers.AssertTollExecSuccess(t, ccipContracts, executionLog)
		}
		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     SourcePool,
				Address:  ccipContracts.SourcePool.Address(),
				Expected: mustAddBigInt(sourceBalances[SourcePool], "30000000000000000297").String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			}, // (10e18 + 100 - 1)*3
			{
				Name:     TollOnRampRouter,
				Address:  ccipContracts.TollOnRampRouter.Address(),
				Expected: mustAddBigInt(sourceBalances[TollOnRampRouter], "3").String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			},
		})
		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     Receiver,
				Address:  ccipContracts.Receivers[0].Receiver.Address(),
				Expected: mustAddBigInt(destBalances[Receiver], "27225848400000000297").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
				Within:   "1000000000000000000",
			}, // 3 toll fees +/- 1 link
			{
				Name:     DestPool,
				Address:  ccipContracts.DestPool.Address(),
				Expected: mustSubBigInt(destBalances[DestPool], "30000000000000000297").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			},
			{
				Name:     TollOffRampRouter,
				Address:  ccipContracts.TollOffRamp.Address(),
				Expected: mustAddBigInt(destBalances[TollOffRampRouter], "2852678400000000000").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
				Within:   "1000000000000000000",
			}, // +/- 1 link
		})
		tollCurrentSeqNum += n
	})

	t.Run("batch auto-execute subscription", func(t *testing.T) {
		sourceBalances := ccipContracts.GetBalances([]testhelpers.BalanceReq{
			{Name: SourcePool, Addr: ccipContracts.SourcePool.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: SourceSub, Addr: ccipContracts.SourceUser.From, Getter: ccipContracts.GetSourceSubBalance},
		})
		destBalances := ccipContracts.GetBalances([]testhelpers.BalanceReq{
			{Name: Receiver, Addr: ccipContracts.Receivers[0].Receiver.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: DestPool, Addr: ccipContracts.DestPool.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: DestSub, Addr: ccipContracts.Receivers[0].Receiver.Address(), Getter: ccipContracts.GetDestSubBalance},
		})
		tokenAmount := big.NewInt(100)
		var txs []*gethtypes.Transaction
		n := 3
		for i := 0; i < n; i++ {
			_, err = ccipContracts.SourceLinkToken.Approve(ccipContracts.SourceUser, ccipContracts.SubOnRampRouter.Address(), tokenAmount)
			require.NoError(t, err)
			txs = append(txs, testhelpers.QueueSubRequest(t, ccipContracts, "hey DON, execute for me", []common.Address{ccipContracts.SourceLinkToken.Address()},
				[]*big.Int{tokenAmount}, big.NewInt(100_000), ccipContracts.Receivers[0].Receiver.Address()))
		}
		ccipContracts.SourceChain.Commit()
		// Send a batch of requests in a single block
		testhelpers.ConfirmTxs(t, txs, ccipContracts.SourceChain)
		var reqs []logpoller.Log
		for i := 0; i < n; i++ {
			reqs = append(reqs, testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, ccip.CCIPSubSendRequested, ccipContracts.SubOnRamp.Address(), nodes, subCurrentSeqNum+i))
		}
		// Should see a report with the full range
		testhelpers.EventuallyReportRelayed(t, ccipContracts, ccipContracts.SubOnRamp.Address(), subCurrentSeqNum, subCurrentSeqNum+n-1)
		// Should all be executed
		for i := range reqs {
			executionLog := testhelpers.AllNodesHaveExecutedSeqNum(t, ccipContracts, ccipContracts.SubOffRamp.Address(), nodes, subCurrentSeqNum+i)
			testhelpers.AssertSubExecSuccess(t, ccipContracts, executionLog)
		}
		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     SourcePool,
				Address:  ccipContracts.SourcePool.Address(),
				Expected: mustAddBigInt(sourceBalances[SourcePool], "300").String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			}, // 100 transfer
			{
				Name:     SourceSub,
				Address:  ccipContracts.SourceUser.From,
				Expected: mustSubBigInt(sourceBalances[SourceSub], "3").String(),
				Getter:   ccipContracts.GetSourceSubBalance,
			}, // Pays 1 in fee
		})
		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     Receiver,
				Address:  ccipContracts.Receivers[0].Receiver.Address(),
				Expected: mustAddBigInt(destBalances[Receiver], "300").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // Full amount gets transferred
			{
				Name:     DestPool,
				Address:  ccipContracts.DestPool.Address(),
				Expected: mustSubBigInt(destBalances[DestPool], "300").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // We lose 100 link from the pool
			{
				Name:     DestSub,
				Address:  ccipContracts.Receivers[0].Receiver.Address(),
				Expected: mustSubBigInt(destBalances[DestSub], "1864160000000000000").String(),
				Getter:   ccipContracts.GetDestSubBalance,
				Within:   "1000000000000000000"}, // Costs ~0.65 link. Varies slightly due to variable calldata encoding gas costs.
		})
		subCurrentSeqNum += n
	})

	t.Run("single strict sequencing auto-execute subscription", func(t *testing.T) {
		sourceBalances := ccipContracts.GetBalances([]testhelpers.BalanceReq{
			{Name: SourcePool, Addr: ccipContracts.SourcePool.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: SourceSub, Addr: ccipContracts.SourceUser.From, Getter: ccipContracts.GetSourceSubBalance},
		})
		destBalances := ccipContracts.GetBalances([]testhelpers.BalanceReq{
			{Name: Receiver, Addr: ccipContracts.Receivers[1].Receiver.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: DestPool, Addr: ccipContracts.DestPool.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: DestSub, Addr: ccipContracts.Receivers[1].Receiver.Address(), Getter: ccipContracts.GetDestSubBalance},
		})
		tokenAmount := big.NewInt(100)
		_, err = ccipContracts.SourceLinkToken.Approve(ccipContracts.SourceUser, ccipContracts.SubOnRampRouter.Address(), tokenAmount)
		require.NoError(t, err)
		ccipContracts.SourceChain.Commit()
		testhelpers.SendSubRequest(t, ccipContracts, "hey DON, execute for me", []common.Address{ccipContracts.SourceLinkToken.Address()},
			[]*big.Int{tokenAmount}, big.NewInt(100_000), ccipContracts.Receivers[1].Receiver.Address())
		testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, ccip.CCIPSubSendRequested, ccipContracts.SubOnRamp.Address(), nodes, subCurrentSeqNum)
		testhelpers.EventuallyReportRelayed(t, ccipContracts, ccipContracts.SubOnRamp.Address(), subCurrentSeqNum, subCurrentSeqNum)
		executionLog := testhelpers.AllNodesHaveExecutedSeqNum(t, ccipContracts, ccipContracts.SubOffRamp.Address(), nodes, subCurrentSeqNum)
		testhelpers.AssertSubExecSuccess(t, ccipContracts, executionLog)
		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     SourcePool,
				Address:  ccipContracts.SourcePool.Address(),
				Expected: mustAddBigInt(sourceBalances[SourcePool], "100").String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			}, // 100 transfer
			{
				Name: SourceSub, Address: ccipContracts.SourceUser.From,
				Expected: mustSubBigInt(sourceBalances[SourceSub], "1").String(),
				Getter:   ccipContracts.GetSourceSubBalance,
			}, // Pays 1 in fee
		})
		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     Receiver,
				Address:  ccipContracts.Receivers[1].Receiver.Address(),
				Expected: mustAddBigInt(destBalances[Receiver], "100").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // Full amount gets transferred
			{
				Name:     DestPool,
				Address:  ccipContracts.DestPool.Address(),
				Expected: mustSubBigInt(destBalances[DestPool], "100").String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // We lose 100 link from the pool
			{
				Name: DestSub, Address: ccipContracts.Receivers[1].Receiver.Address(),
				Expected: mustSubBigInt(destBalances[DestSub], "654720000000000000").String(),
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
		sourceBalances := ccipContracts.GetBalances([]testhelpers.BalanceReq{
			{Name: SourcePool, Addr: ccipContracts.SourcePool.Address(), Getter: ccipContracts.GetSourceLinkBalance},
			{Name: Sender, Addr: ccipContracts.SourceUser.From, Getter: ccipContracts.GetSourceLinkBalance},
			{Name: SourceSub, Addr: ccipContracts.SourceUser.From, Getter: ccipContracts.GetSourceSubBalance},
		})
		destBalances := ccipContracts.GetBalances([]testhelpers.BalanceReq{
			{Name: Receiver, Addr: ccipContracts.Receivers[1].Receiver.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: DestPool, Addr: ccipContracts.DestPool.Address(), Getter: ccipContracts.GetDestLinkBalance},
			{Name: DestSub, Addr: ccipContracts.Receivers[1].Receiver.Address(), Getter: ccipContracts.GetDestSubBalance},
		})

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
		testhelpers.SendSubRequest(t, ccipContracts, "hey DON, execute for me", []common.Address{ccipContracts.SourceLinkToken.Address()},
			[]*big.Int{tokenAmounts[0]}, big.NewInt(100_000), ccipContracts.Receivers[1].Receiver.Address())
		failedReq := testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, ccip.CCIPSubSendRequested, ccipContracts.SubOnRamp.Address(), nodes, subCurrentSeqNum)
		testhelpers.EventuallyReportRelayed(t, ccipContracts, ccipContracts.SubOnRamp.Address(), subCurrentSeqNum, subCurrentSeqNum)
		reportForFailedReq := testhelpers.EventuallyRelayReportAccepted(t, ccipContracts, currentBlockNumber)
		executionLog := testhelpers.AllNodesHaveExecutedSeqNum(t, ccipContracts, ccipContracts.SubOffRamp.Address(), nodes, subCurrentSeqNum)

		// the transaction should get reverted and the execution status should be failed
		testhelpers.AssertSubExecFailure(t, ccipContracts, executionLog)
		subCurrentSeqNum++

		// flip the revert settings on receiver
		_, err = ccipContracts.Receivers[1].Receiver.SetRevert(ccipContracts.DestUser, false)
		require.NoError(t, err, "setting revert to false on the receiver")
		ccipContracts.DestChain.Commit()

		// send a bunch of subsequent ones and it should not be executed
		var pendingReqNumbers []int
		for i := 1; i <= 3; i++ {
			testhelpers.SendSubRequest(t, ccipContracts, "hey DON, execute for me", []common.Address{ccipContracts.SourceLinkToken.Address()},
				[]*big.Int{tokenAmounts[i]}, big.NewInt(100_000), ccipContracts.Receivers[1].Receiver.Address())
			testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, ccip.CCIPSubSendRequested, ccipContracts.SubOnRamp.Address(), nodes, subCurrentSeqNum)
			testhelpers.EventuallyReportRelayed(t, ccipContracts, ccipContracts.SubOnRamp.Address(), subCurrentSeqNum, subCurrentSeqNum)
			executionLog := testhelpers.NoNodesHaveExecutedSeqNum(t, ccipContracts, ccipContracts.SubOffRamp.Address(), nodes, subCurrentSeqNum)
			require.Empty(t, executionLog)
			pendingReqNumbers = append(pendingReqNumbers, subCurrentSeqNum)
			subCurrentSeqNum++
		}

		// manually execute the failed request
		failedSeqNum := testhelpers.ExecuteSubMessage(t, ccipContracts, failedReq, []logpoller.Log{failedReq}, reportForFailedReq)
		testhelpers.EventuallyExecutionStateChangedToSuccess(t, ccipContracts, []uint64{failedSeqNum})

		// verify all of the pending requests should be successfully executed now
		for _, seqNo := range pendingReqNumbers {
			t.Logf("Verify execution for pending seq Number %d", seqNo)
			executionLog := testhelpers.AllNodesHaveExecutedSeqNum(t, ccipContracts, ccipContracts.SubOffRamp.Address(), nodes, seqNo)
			testhelpers.AssertSubExecSuccess(t, ccipContracts, executionLog)
		}

		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     SourcePool,
				Address:  ccipContracts.SourcePool.Address(),
				Expected: mustAddBigInt(sourceBalances[SourcePool], strTotalAmount).String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			}, // 100 transfer
			{
				Name:     SourceSub,
				Address:  ccipContracts.SourceUser.From,
				Expected: mustSubBigInt(sourceBalances[SourceSub], "4").String(),
				Getter:   ccipContracts.GetSourceSubBalance,
			}, // Pays 4 in fee
			{
				Name:     Sender,
				Address:  ccipContracts.SourceUser.From,
				Expected: mustSubBigInt(sourceBalances[Sender], strTotalAmount).String(),
				Getter:   ccipContracts.GetSourceLinkBalance,
			}, // 100 transfer
		})
		t.Logf("destBalances[DestSub] %v", destBalances[DestSub])
		ccipContracts.AssertBalances([]testhelpers.BalanceAssertion{
			{
				Name:     Receiver,
				Address:  ccipContracts.Receivers[1].Receiver.Address(),
				Expected: mustAddBigInt(destBalances[Receiver], strTotalAmount).String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // Full amount gets transferred
			{
				Name:     DestPool,
				Address:  ccipContracts.DestPool.Address(),
				Expected: mustSubBigInt(destBalances[DestPool], strTotalAmount).String(),
				Getter:   ccipContracts.GetDestLinkBalance,
			}, // We lose 100 link from the pool

			{
				Name:     DestSub,
				Address:  ccipContracts.Receivers[1].Receiver.Address(),
				Expected: mustSubBigInt(destBalances[DestSub], "2311120000000000000").String(),
				Getter:   ccipContracts.GetDestSubBalance,
				Within:   "100000000000000000",
			}, // Costs ~0.77 link per transfer for 3 auto req. Varies slightly due to variable calldata encoding gas costs.
		})
	})
}
