package ccip_test

import (
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/test-go/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/testhelpers"
	integrationtesthelpers "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/testhelpers/integration"
)

func Test_CLOSpecApprovalFlow(t *testing.T) {
	ccipTH := integrationtesthelpers.SetupCCIPIntegrationTH(t, testhelpers.SourceChainID, testhelpers.SourceChainSelector, testhelpers.DestChainID, testhelpers.DestChainSelector)

	linkUSD := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte(`{"UsdPerLink": "8000000000000000000"}`))
		require.NoError(t, err)
	}))
	ethUSD := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte(`{"UsdPerETH": "1700000000000000000000"}`))
		require.NoError(t, err)
	}))
	wrapped, err1 := ccipTH.Source.Router.GetWrappedNative(nil)
	require.NoError(t, err1)
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

	// Create initial job specs
	jobParams := ccipTH.SetUpNodesAndJobs(t, tokenPricesUSDPipeline, 19399)
	ccipTH.SetupFeedsManager(t)

	// Propose and approve new specs
	ccipTH.ApproveJobSpecs(t, jobParams, tokenPricesUSDPipeline)
	//ccipTH.ApproveJobSpecs(t, jobParams)

	// Sanity check that CCIP works after CLO flow
	currentSeqNum := 1

	extraArgs, err := testhelpers.GetEVMExtraArgsV1(big.NewInt(200_003), false)
	require.NoError(t, err)

	msg := router.ClientEVM2AnyMessage{
		Receiver:     testhelpers.MustEncodeAddress(t, ccipTH.Dest.Receivers[0].Receiver.Address()),
		Data:         []byte("hello"),
		TokenAmounts: []router.ClientEVMTokenAmount{},
		FeeToken:     ccipTH.Source.LinkToken.Address(),
		ExtraArgs:    extraArgs,
	}
	fee, err := ccipTH.Source.Router.GetFee(nil, testhelpers.DestChainSelector, msg)
	require.NoError(t, err)

	_, err = ccipTH.Source.LinkToken.Approve(ccipTH.Source.User, ccipTH.Source.Router.Address(), new(big.Int).Set(fee))
	require.NoError(t, err)
	ccipTH.Source.Chain.Commit()

	ccipTH.SendRequest(t, msg)
	ccipTH.AllNodesHaveReqSeqNum(t, currentSeqNum)
	ccipTH.EventuallyReportCommitted(t, currentSeqNum)

	executionLogs := ccipTH.AllNodesHaveExecutedSeqNums(t, currentSeqNum, currentSeqNum)
	assert.Len(t, executionLogs, 1)
	ccipTH.AssertExecState(t, executionLogs[0], testhelpers.ExecutionStateSuccess)
}
