package smoke

//revive:disable:dot-imports
import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/chainlink-env/environment"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	"github.com/stretchr/testify/require"

	ctfUtils "github.com/smartcontractkit/chainlink-testing-framework/utils"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	networks "github.com/smartcontractkit/chainlink/integration-tests"
	"github.com/smartcontractkit/chainlink/integration-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/client"
)

// not for usual test run. just a utility script to decode the event from topic hash
// remove this later
func TestPrintEvent(t *testing.T) {
	t.Skip()
	//dataError := []byte("0x894882b8")
	data, err := hex.DecodeString("4cd172fb90d81a44670b97a6e2a5a3b01417f33a809b634a5a1764e93d338e1f")
	inputs, err := hex.DecodeString("")
	jsonABI, err := abi.JSON(strings.NewReader(evm_2_evm_toll_onramp.EVM2EVMTollOnRampABI))
	require.NoError(t, err, "should be able to jsonify abi")
	for _, abiEvent := range jsonABI.Events {
		//fmt.Println(abiEvent.ID)
		if bytes.Equal(data[:4], abiEvent.ID.Bytes()[:4]) {
			// Found a matching error
			v, _ := abiEvent.Inputs.Unpack(inputs)
			log.Info().Interface("v", v).Msg("Event Name")
			log.Info().Str("Event", abiEvent.Name).Msg("Event Name")
			return
		}
	}
}

// not for usual test run. just a utility script to decode the revert reason from error id
// remove this later
func TestPrintRevertReason(t *testing.T) {
	t.Skip()
	//dataError := []byte("0x894882b8")
	data, err := hex.DecodeString("894882b8")
	jsonABI, err := abi.JSON(strings.NewReader(any_2_evm_toll_offramp.EVM2EVMTollOffRampABI))
	require.NoError(t, err, "should be able to jsonify abi")
	for k, abiError := range jsonABI.Errors {
		fmt.Println(abiError.ID)
		if bytes.Equal(data[:4], abiError.ID.Bytes()[:4]) {
			// Found a matching error
			v, err := abiError.Unpack(data)
			require.NoError(t, err)
			log.Info().Interface("Error", k).Interface("args - ", v).Msg("Revert Reason")
			fmt.Println(k, v)
			return
		}
	}
}

var _ = FDescribe("CCIP interactions test @ccip", func() {
	var (
		diffNetworkEntries = []TableEntry{
			Entry("CCIP suite on 2 Geths @simulated", networks.NetworkAlpha, networks.NetworkBeta),
		}
		sourceChainClient blockchain.EVMClient
		testEnvironment   *environment.Environment
		chainlinkNodes    []*client.Chainlink
	)

	AfterEach(func() {
		By("Tearing down the environment")
		sourceChainClient.GasStats().PrintStats()
		err := actions.TeardownSuite(testEnvironment, ctfUtils.ProjectRoot, chainlinkNodes, nil, sourceChainClient)
		Expect(err).ShouldNot(HaveOccurred(), "Environment teardown shouldn't fail")
	})

	DescribeTable("CCIP suite on different EVM networks", func(
		sourceNetwork *blockchain.EVMNetwork,
		destNetwork *blockchain.EVMNetwork,
	) {
		var (
			sourceCCIP *actions.SourceCCIPModule
			destCCIP   *actions.DestCCIPModule
		)
		By("Deploying the environment")
		testEnvironment = actions.DeployEnvironments(sourceNetwork, destNetwork)

		By("Setting up chainlink nodes")
		testSetUp := actions.SetUpNodesAndKeys(sourceNetwork, destNetwork, testEnvironment)
		clNodes := testSetUp.CLNodesWithKeys
		mockServer := testSetUp.MockServer
		chainlinkNodes = testSetUp.CLNodes
		sourceChainClient = testSetUp.SourceChainClient
		destChainClient := testSetUp.DestChainClient

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
		By("Setting up bootstrap, relay and execute job")
		var tokenAddr []string
		for _, token := range destCCIP.Common.BridgeTokens {
			tokenAddr = append(tokenAddr, token.Address())
		}
		tokenAddr = append(tokenAddr, destCCIP.Common.FeeToken.Address())

		actions.CreateOCRJobsForCCIP(
			clNodes, sourceCCIP.TollOnRamp.Address(),
			sourceCCIP.SubOnRamp.Address(),
			destCCIP.BlobVerifier.Address(),
			destCCIP.TollOffRamp.Address(),
			destCCIP.SubOffRamp.Address(),
			sourceChainClient, destChainClient,
			tokenAddr,
			mockServer,
		)

		// set up ocr2 config
		By("Setting up ocr config in blob verifier and offramp")
		actions.SetOCRConfigs(clNodes[1:], *destCCIP) // first node is the bootstrapper

		// initiate transfer with toll and verify
		By("Multiple Token transfer with toll, watch for updated sequence numbers and events logs, " +
			"verify balance in receiving and sending account pre and post transfer")
		actions.TokenTransferWithToll(*sourceCCIP, *destCCIP)

		// initiate transfer with subscription and verify
		By("Multiple Token transfer with subscription, watch for updated sequence numbers and events logs, " +
			"verify receiver,sender and subscription balance pre and post transfer")
		actions.CreateAndFundSubscription(*sourceCCIP, *destCCIP, big.NewInt(0).Mul(big.NewInt(80), big.NewInt(1e18)))
		actions.TokenTransferWithSub(*sourceCCIP, *destCCIP)
	},
		diffNetworkEntries,
	)
})
