package ccip_test

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/test-go/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/mock_v3_aggregator_contract"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/pricegetter"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/testhelpers"
	integrationtesthelpers "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/testhelpers/integration"
)

func Test_CLOSpecApprovalFlow(t *testing.T) {
	ccipTH := integrationtesthelpers.SetupCCIPIntegrationTH(t, testhelpers.SourceChainID, testhelpers.SourceChainSelector, testhelpers.DestChainID, testhelpers.DestChainSelector)
	//tokenPricesUSDPipeline, linkUSD, ethUSD := ccipTH.CreatePricesPipeline(t)
	//defer linkUSD.Close()
	//defer ethUSD.Close()

	// Create initial job specs
	//fmt.Printf("token prices pipeline: %s", tokenPricesUSDPipeline)

	// Use price registry to find token prices.
	srcLinkAddr := ccipTH.Source.LinkToken.Address()
	dstLinkAddr := ccipTH.Dest.LinkToken.Address()

	srcNativeAddr, err := ccipTH.Source.Router.GetWrappedNative(nil)
	require.NoError(t, err)
	dstNativeAddr := ccipTH.Dest.WrappedNative.Address()

	fmt.Printf("=> src native at: %s\n", srcNativeAddr)
	fmt.Printf("=> dst native at: %s\n", dstNativeAddr)

	fmt.Printf("=> src LINK at: %s\n", srcLinkAddr)
	fmt.Printf("=> dst LINK at: %s\n", dstLinkAddr)

	// Set up the aggregators here to avoid modifying ccipTH.
	aggSrcNatAddr, _, aggSrcNat, err := mock_v3_aggregator_contract.DeployMockV3AggregatorContract(ccipTH.Source.User, ccipTH.Source.Chain, 18, big.NewInt(2e18))
	require.NoError(t, err)
	_, err = aggSrcNat.UpdateRoundData(ccipTH.Source.User, big.NewInt(50), big.NewInt(17000000), big.NewInt(1000), big.NewInt(1000))
	require.NoError(t, err)
	ccipTH.Source.Chain.Commit()

	aggSrcLnkAddr, _, aggSrcLnk, err := mock_v3_aggregator_contract.DeployMockV3AggregatorContract(ccipTH.Source.User, ccipTH.Source.Chain, 18, big.NewInt(3e18))
	require.NoError(t, err)
	ccipTH.Dest.Chain.Commit()
	_, err = aggSrcLnk.UpdateRoundData(ccipTH.Source.User, big.NewInt(50), big.NewInt(8000000), big.NewInt(1000), big.NewInt(1000))
	require.NoError(t, err)
	ccipTH.Source.Chain.Commit()

	aggDstLnkAddr, _, aggDstLnk, err := mock_v3_aggregator_contract.DeployMockV3AggregatorContract(ccipTH.Dest.User, ccipTH.Dest.Chain, 18, big.NewInt(3e18))
	require.NoError(t, err)
	ccipTH.Dest.Chain.Commit()
	_, err = aggDstLnk.UpdateRoundData(ccipTH.Dest.User, big.NewInt(50), big.NewInt(8000000), big.NewInt(1000), big.NewInt(1000))
	require.NoError(t, err)
	ccipTH.Dest.Chain.Commit()

	// Check content is ok on aggregator.
	tmp, err := aggDstLnk.LatestRoundData(&bind.CallOpts{})
	require.NoError(t, err)
	fmt.Printf("=> CHECK round: %s\n", tmp.RoundId.String())
	fmt.Printf("=> CHECK answer: %s\n", tmp.Answer.String())

	tokenPricesConfig := pricegetter.DynamicPriceGetterConfig{
		AggregatorPrices: map[common.Address]pricegetter.AggregatorPriceConfig{
			srcLinkAddr: {
				ChainID:         ccipTH.Source.ChainID,
				ContractAddress: aggSrcLnkAddr,
			},
			srcNativeAddr: {
				ChainID:         ccipTH.Source.ChainID,
				ContractAddress: aggSrcNatAddr,
			},
			dstLinkAddr: {
				ChainID:         ccipTH.Dest.ChainID,
				ContractAddress: aggDstLnkAddr,
			},
			//ccipTH.Dest.WrappedNative.Address(): {
			//	ChainID:         ccipTH.Dest.ChainID,
			//	ContractAddress: aggDstNat,
			//},
		},
		StaticPrices: map[common.Address]pricegetter.StaticPriceConfig{},
		//StaticPrices: map[common.Address]pricegetter.StaticPriceConfig{
		//	dstLinkAddr: {
		//		ChainID: ccipTH.Dest.ChainID,
		//		Price:   8000000000000000000,
		//	},
		//},
	}
	tokenPricesConfigBytes, err := json.MarshalIndent(tokenPricesConfig, "", " ")
	//tokenPricesConfigBytes, err := json.Marshal(tokenPricesConfig)
	require.NoError(t, err)
	tokenPricesConfigJson := string(tokenPricesConfigBytes)
	fmt.Printf("Token prices config:\n%s\n", tokenPricesConfigJson)

	jobParams := ccipTH.SetUpNodesAndJobs(t, tokenPricesConfigJson, "http://blah.com")
	ccipTH.SetupFeedsManager(t)

	// Propose and approve new specs
	ccipTH.ApproveJobSpecs(t, jobParams, tokenPricesConfigJson)
	// TODO generate one more run with propose & approve
	// ccipTH.ApproveJobSpecs(t, jobParams)

	// Sanity check that CCIP works after CLO flow
	currentSeqNum := 1

	extraArgs, err := testhelpers.GetEVMExtraArgsV1(big.NewInt(200_003), false)
	require.NoError(t, err)

	msg := router.ClientEVM2AnyMessage{
		Receiver:     testhelpers.MustEncodeAddress(t, ccipTH.Dest.Receivers[0].Receiver.Address()),
		Data:         utils.RandomAddress().Bytes(),
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
