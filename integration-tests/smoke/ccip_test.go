package smoke

//revive:disable:dot-imports
import (
	"math/big"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/chainlink-env/environment"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	ctfUtils "github.com/smartcontractkit/chainlink-testing-framework/utils"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/client"
)

func TestSmokeCCIP(t *testing.T) {
	t.Parallel()
	var (
		sourceChainClient blockchain.EVMClient
		destChainClient   blockchain.EVMClient
		testEnvironment   *environment.Environment
		chainlinkNodes    []*client.Chainlink
		sourceCCIP        *actions.SourceCCIPModule
		destCCIP          *actions.DestCCIPModule
	)

	t.Cleanup(func() {
		log.Info().Msg("Tearing down the environment")
		err := actions.TeardownSuite(t, testEnvironment, ctfUtils.ProjectRoot, chainlinkNodes,
			nil, destChainClient, sourceChainClient)
		require.NoError(t, err, "Environment teardown shouldn't fail")
	})

	log.Info().Msg("Deploying the environment")
	testEnvironment = actions.DeployEnvironments(
		t, &environment.Config{NamespacePrefix: "smoke-ccip"},
		map[string]interface{}{
			"replicas": "6",
			"toml":     actions.DefaultCCIPCLNodeEnv(t),
			"env": map[string]interface{}{
				"CL_DEV": "true",
			},
		})

	log.Info().Msg("Setting up chainlink nodes")
	testSetUp := actions.SetUpNodesAndKeys(t, testEnvironment, big.NewFloat(10))
	clNodes := testSetUp.CLNodesWithKeys
	mockServer := testSetUp.MockServer
	chainlinkNodes = testSetUp.CLNodes
	sourceChainClient = testSetUp.SourceChainClient
	destChainClient = testSetUp.DestChainClient

	// transfer more than one token
	transferAmounts := []*big.Int{big.NewInt(5e17), big.NewInt(5e17)}

	// deploy all source contracts
	sourceCCIP = actions.DefaultSourceCCIPModule(sourceChainClient, destChainClient.GetChainID().Uint64(), transferAmounts)
	log.Info().Msg("Deploying source contracts")
	sourceCCIP.DeployContracts(t)

	// deploy all destination contracts
	destCCIP = actions.DefaultDestinationCCIPModule(destChainClient, sourceChainClient.GetChainID().Uint64())
	log.Info().Msg("Deploying destination contracts")
	destCCIP.DeployContracts(t, *sourceCCIP)

	// set up ocr2 jobs
	log.Info().Msg("Setting up bootstrap, commit and execute job")
	var tokenAddr []string
	for _, token := range destCCIP.Common.BridgeTokens {
		tokenAddr = append(tokenAddr, token.Address())
	}
	tokenAddr = append(tokenAddr, destCCIP.Common.FeeToken.Address())
	actions.CreateOCRJobsForCCIP(
		t, clNodes[0], nil, clNodes[1:], nil,
		sourceCCIP.TollOnRamp.EthAddress,
		sourceCCIP.GEOnRamp.EthAddress,
		destCCIP.CommitStore.EthAddress,
		destCCIP.TollOffRamp.EthAddress,
		destCCIP.GEOffRamp.EthAddress,
		sourceChainClient, destChainClient,
		tokenAddr,
		mockServer,
	)

	// set up ocr2 config
	log.Info().Msg("Setting up ocr config in commit store and offramp")
	actions.SetOCRConfigs(t, clNodes[1:], nil, *destCCIP) // first node is the bootstrapper

	ccipTest := actions.NewCCIPTest(t, sourceCCIP, destCCIP, time.Minute)

	// initiate transfer with GE and verify
	t.Run("Multiple Token transfer with GE", func(t *testing.T) {
		ccipTest.SendGERequests(1)
		ccipTest.ValidateGERequests()
	})

	// initiate transfer with toll and verify
	t.Run("Multiple Token transfer with toll", func(t *testing.T) {
		ccipTest.SendTollRequests(1)
		ccipTest.ValidateTollRequests()
	})
}
