package metatx_test

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/onsi/gomega"
	"github.com/smartcontractkit/chainlink/core/assets"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/meta_erc20"
	"github.com/smartcontractkit/chainlink/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/core/services/metatx"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/testhelpers"
	"github.com/stretchr/testify/require"
	"github.com/test-go/testify/assert"
)

var (
	sourceChainID = uint64(1000)
	destChainID   = uint64(1337)
)

func TestMetaERC20CrossChain(t *testing.T) {
	ccipContracts := testhelpers.SetupCCIPContracts(t, sourceChainID, destChainID)

	holder1Key := cltest.MustGenerateRandomKey(t)
	holder1Transactor, err := bind.NewKeyedTransactorWithChainID(holder1Key.ToEcdsaPrivKey(), ccipContracts.Source.Chain.Blockchain().Config().ChainID)
	require.NoError(t, err)

	relayKey := cltest.MustGenerateRandomKey(t)
	relayTransactor, err := bind.NewKeyedTransactorWithChainID(relayKey.ToEcdsaPrivKey(), ccipContracts.Source.Chain.Blockchain().Config().ChainID)
	require.NoError(t, err)

	holder2Key := cltest.MustGenerateRandomKey(t)
	holder2Transactor, err := bind.NewKeyedTransactorWithChainID(holder2Key.ToEcdsaPrivKey(), ccipContracts.Dest.Chain.Blockchain().Config().ChainID)
	require.NoError(t, err)

	var (
		holder1 = holder1Transactor
		holder2 = holder2Transactor
		relay   = relayTransactor
	)

	sourceMetaWrapperAddress, _, sourceMetaWrapper, err := meta_erc20.DeployMetaERC20(
		ccipContracts.Source.User, ccipContracts.Source.Chain, assets.Ether(int64(1e18)).ToInt(),
		ccipContracts.Source.LinkToken.Address(), ccipContracts.Source.Router.Address())
	require.NoError(t, err)

	ccipContracts.Source.Chain.Commit()

	destMetaWrapperAddress, _, destMetaWrapper, err := meta_erc20.DeployMetaERC20(
		ccipContracts.Dest.User, ccipContracts.Dest.Chain, assets.Ether(int64(0)).ToInt(),
		ccipContracts.Dest.LinkToken.Address(), ccipContracts.Dest.Router.Address())
	require.NoError(t, err)

	ccipContracts.Dest.Chain.Commit()

	sourcePoolAddress, _ := ccipContracts.SetUpNewMintAndBurnPool(sourceMetaWrapperAddress, destMetaWrapperAddress)

	ccipContracts.Dest.Chain.Commit()

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
merge [type=merge left="{}" right="{\\\"%s\\\":$(link_parse), \\\"%s\\\":$(link_parse)}"];`,
		linkEth.URL, ccipContracts.Dest.LinkToken.Address(),
		destMetaWrapperAddress)
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

	t.Run("single cross-chain meta transfer", func(t *testing.T) {

		// transfer MetaERC20 from owner to holder1
		_, err = sourceMetaWrapper.Transfer(ccipContracts.Source.User, holder1.From, assets.Ether(1).ToInt())
		require.NoError(t, err)
		ccipContracts.Source.Chain.Commit()

		holder1Bal, err := sourceMetaWrapper.BalanceOf(nil, holder1.From)
		require.NoError(t, err)
		require.Equal(t, assets.Ether(1).ToInt(), holder1Bal)

		// transfer MetaERC20 from owner to pool
		_, err = sourceMetaWrapper.Transfer(ccipContracts.Source.User, sourcePoolAddress, assets.Ether(10).ToInt())
		require.NoError(t, err)
		ccipContracts.Source.Chain.Commit()

		sourcePoolBal, err := sourceMetaWrapper.BalanceOf(nil, sourcePoolAddress)
		require.NoError(t, err)
		require.Equal(t, assets.Ether(10).ToInt(), sourcePoolBal)

		// transfer MetaERC20 from owner to MetaERC20 contract
		_, err = sourceMetaWrapper.Transfer(ccipContracts.Source.User, sourceMetaWrapperAddress, assets.Ether(100).ToInt())
		require.NoError(t, err)
		ccipContracts.Source.Chain.Commit()

		sourceMetaWrapperBal, err := sourceMetaWrapper.BalanceOf(nil, sourceMetaWrapperAddress)
		require.NoError(t, err)
		require.Equal(t, assets.Ether(100).ToInt(), sourceMetaWrapperBal)

		// transfer LINK from owner to relay
		_, err = ccipContracts.Source.LinkToken.Transfer(ccipContracts.Source.User, relay.From, assets.Ether(1).ToInt())
		require.NoError(t, err)
		ccipContracts.Source.Chain.Commit()

		relayLinkBal, err := ccipContracts.Source.LinkToken.BalanceOf(nil, relay.From)
		require.NoError(t, err)
		require.Equal(t, assets.Ether(1).ToInt(), relayLinkBal)

		// meta transfer from holder1 to holder2
		deadline := big.NewInt(int64(ccipContracts.Source.Chain.Blockchain().CurrentHeader().Time + uint64(time.Hour)))
		msg := meta_erc20.MetaERC20MetaTransferMessage{
			Owner:    holder1.From,
			To:       holder2.From,
			Amount:   assets.Ether(1).ToInt(),
			Deadline: deadline,
			ChainId:  ccipContracts.Dest.ChainID,
			GasLimit: 200_000,
		}

		v, r, s, err := metatx.SignMetaTransfer(
			sourceMetaWrapper,
			holder1Key.ToEcdsaPrivKey(),
			msg,
		)
		require.NoError(t, err)

		n, err := ccipContracts.Source.Chain.NonceAt(testutils.Context(t), ccipContracts.Source.User.From, nil)
		require.NoError(t, err)

		tx := types.NewTransaction(
			n, relay.From,
			assets.Ether(1).ToInt(),
			21000,
			assets.GWei(1).ToInt(),
			nil)
		signedTx, err := ccipContracts.Source.User.Signer(ccipContracts.Source.User.From, tx)
		require.NoError(t, err)
		err = ccipContracts.Source.Chain.SendTransaction(testutils.Context(t), signedTx)
		require.NoError(t, err)
		ccipContracts.Source.Chain.Commit()

		relayNativeBalance, err := ccipContracts.Source.Chain.BalanceAt(context.Background(), relay.From, nil)
		require.NoError(t, err)
		require.Equal(t, assets.Ether(1).ToInt(), relayNativeBalance)

		metaErc20, err := meta_erc20.NewMetaERC20(sourceMetaWrapperAddress, ccipContracts.Source.Chain)
		require.NoError(t, err)

		_, err = metaErc20.ApproveRouter(ccipContracts.Source.User, assets.Ether(2).ToInt())
		require.NoError(t, err)
		ccipContracts.Source.Chain.Commit()

		_, err = ccipContracts.Source.LinkToken.Approve(ccipContracts.Source.User, sourceMetaWrapperAddress, assets.Ether(200).ToInt())
		require.NoError(t, err)
		ccipContracts.Source.Chain.Commit()

		_, err = metaErc20.Fund(ccipContracts.Source.User, assets.Ether(200).ToInt())
		require.NoError(t, err)
		ccipContracts.Source.Chain.Commit()

		_, err = metaErc20.MetaTransfer(relay, msg, v, r, s)
		require.NoError(t, err)
		ccipContracts.Source.Chain.Commit()

		gomega.NewWithT(t).Eventually(func() bool {
			ccipContracts.Dest.Chain.Commit()
			holder2Balance, err := destMetaWrapper.BalanceOf(nil, holder2.From)
			require.NoError(t, err)
			return holder2Balance.Cmp(assets.Ether(1).ToInt()) == 0
		}, testutils.WaitTimeout(t), 5*time.Second).Should(gomega.BeTrue())

		eventSignatures := ccip.GetEventSignatures()

		// Should eventually see this executed.
		// TODO: make sure blocks are commited before this check
		testhelpers.AllNodesHaveReqSeqNum(t, ccipContracts, eventSignatures, ccipContracts.Source.OnRamp.Address(), nodes, geCurrentSeqNum)
		testhelpers.EventuallyReportCommitted(t, ccipContracts, ccipContracts.Source.OnRamp.Address(), geCurrentSeqNum)

		executionLogs := testhelpers.AllNodesHaveExecutedSeqNums(t, ccipContracts, eventSignatures, ccipContracts.Dest.OffRamp.Address(), nodes, geCurrentSeqNum, geCurrentSeqNum)
		assert.Len(t, executionLogs, 1)
		testhelpers.AssertExecState(t, ccipContracts, executionLogs[0], ccip.Success)
	})
}
