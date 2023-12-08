package ccip_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/test-go/testify/require"

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

	jobParams := ccipTH.SetUpNodesAndJobs(t, tokenPricesUSDPipeline, 19399)

	ccipTH.ApproveJobSpecs(t, jobParams)
}
