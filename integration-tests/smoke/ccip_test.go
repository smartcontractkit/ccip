package smoke

//revive:disable:dot-imports
import (
	"math/big"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/smartcontractkit/chainlink-env/environment"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	ctfUtils "github.com/smartcontractkit/chainlink-testing-framework/utils"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/client"
)

var _ = Describe("CCIP interactions test @ccip", func() {
	var (
		sourceChainClient blockchain.EVMClient
		destChainClient   blockchain.EVMClient
		testEnvironment   *environment.Environment
		chainlinkNodes    []*client.Chainlink
	)

	AfterEach(func() {
		By("Tearing down the environment")
		sourceChainClient.GasStats().PrintStats()
		err := actions.TeardownSuite(testEnvironment, ctfUtils.ProjectRoot, chainlinkNodes, nil, destChainClient, sourceChainClient)
		Expect(err).ShouldNot(HaveOccurred(), "Environment teardown shouldn't fail")
	})

	It("Deliver message with token in toll based model", func() {
		var (
			sourceCCIP *actions.SourceCCIPModule
			destCCIP   *actions.DestCCIPModule
		)
		By("Deploying the environment")
		testEnvironment = actions.DeployEnvironments(
			&environment.Config{NamespacePrefix: "smoke-ccip"},
			map[string]interface{}{
				"replicas": "6",
				"toml":     actions.DefaultCCIPCLNodeEnv(),
				"env": map[string]interface{}{
					"CL_DEV": "true",
				},
			})

		By("Setting up chainlink nodes")
		testSetUp := actions.SetUpNodesAndKeys(testEnvironment, big.NewFloat(10))
		clNodes := testSetUp.CLNodesWithKeys
		mockServer := testSetUp.MockServer
		chainlinkNodes = testSetUp.CLNodes
		sourceChainClient = testSetUp.SourceChainClient
		destChainClient = testSetUp.DestChainClient

		// transfer more than one token
		transferAmounts := []*big.Int{big.NewInt(5e17), big.NewInt(5e17)}

		// deploy all source contracts
		sourceCCIP = actions.DefaultSourceCCIPModule(sourceChainClient, destChainClient.GetChainID(), transferAmounts)
		By("Deploying source contracts")
		sourceCCIP.DeployContracts()

		// deploy all destination contracts
		destCCIP = actions.DefaultDestinationCCIPModule(destChainClient, sourceChainClient.GetChainID())
		By("Deploying destination contracts")
		destCCIP.DeployContracts(*sourceCCIP)

		// set up ocr2 jobs
		By("Setting up bootstrap, commit and execute job")
		var tokenAddr []string
		for _, token := range destCCIP.Common.BridgeTokens {
			tokenAddr = append(tokenAddr, token.Address())
		}
		tokenAddr = append(tokenAddr, destCCIP.Common.FeeToken.Address())
		actions.CreateOCRJobsForCCIP(
			clNodes[0], nil, clNodes[1:], nil,
			sourceCCIP.TollOnRamp.EthAddress,
			destCCIP.CommitStore.EthAddress,
			destCCIP.TollOffRamp.EthAddress,
			sourceChainClient, destChainClient,
			tokenAddr,
			mockServer,
		)

		// set up ocr2 config
		By("Setting up ocr config in commit store and offramp")
		actions.SetOCRConfigs(clNodes[1:], nil, *destCCIP) // first node is the bootstrapper

		ccipTest := actions.NewCCIPTest(sourceCCIP, destCCIP, time.Minute)

		// initiate transfer with toll and verify
		By("Multiple Token transfer with toll, watch for updated sequence numbers and events logs, " +
			"verify balance in receiving and sending account pre and post transfer")
		ccipTest.SendTollRequests(1)
		ccipTest.ValidateTollRequests()
	})
})
