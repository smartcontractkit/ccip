package smoke

//revive:disable:dot-imports
import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/chainlink-env/environment"
	"github.com/smartcontractkit/chainlink-env/pkg/cdk8s/blockscout"
	"github.com/smartcontractkit/chainlink-env/pkg/helm/chainlink"
	"github.com/smartcontractkit/chainlink-env/pkg/helm/mockserver"
	mockservercfg "github.com/smartcontractkit/chainlink-env/pkg/helm/mockserver-cfg"
	"github.com/smartcontractkit/chainlink-env/pkg/helm/reorg"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	ctfClient "github.com/smartcontractkit/chainlink-testing-framework/client"
	ctfUtils "github.com/smartcontractkit/chainlink-testing-framework/utils"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4"

	"github.com/smartcontractkit/chainlink/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/utils"
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
	jsonABI, err := abi.JSON(strings.NewReader(evm_2_evm_toll_onramp.EVM2EVMTollOnRampABI))
	require.NoError(t, err, "should be able to jsonify abi")
	for _, abiEvent := range jsonABI.Events {
		//fmt.Println(abiEvent.ID)
		if bytes.Equal(data[:4], abiEvent.ID.Bytes()[:4]) {
			// Found a matching error
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
		testScenarios = []TableEntry{
			Entry("CCIP suite on 2 Geths @simulated", networks.NetworkAlpha, networks.NetworkBeta),
		}

		testEnvironment   *environment.Environment
		sourceChainClient blockchain.EVMClient
		destChainClient   blockchain.EVMClient
		chainlinkNodes    []*client.Chainlink
		mockServer        *ctfClient.MockserverClient
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
		By("Deploying the environment")
		evmNodes, err := json.Marshal([]types.NewNode{
			{
				Name:       "primary_0_source",
				EVMChainID: *utils.NewBigI(sourceNetwork.ChainID),
				WSURL:      null.StringFrom("ws://source-chain-ethereum-geth:8546"),
				HTTPURL:    null.StringFrom("http://source-chain-ethereum-geth:8544"),
				SendOnly:   false,
			},
			{
				Name:       "primary_0_dest",
				EVMChainID: *utils.NewBigI(destNetwork.ChainID),
				WSURL:      null.StringFrom("ws://dest-chain-ethereum-geth:8546"),
				HTTPURL:    null.StringFrom("http://dest-chain-ethereum-geth:8544"),
				SendOnly:   false,
			},
		})
		Expect(err).ShouldNot(HaveOccurred())
		// TODO move env set-up in a generic method to be used by all integration-tests
		testEnvironment = environment.New(&environment.Config{
			NamespacePrefix: "smoke-ccip",
		})
		err = testEnvironment.
			AddHelm(mockservercfg.New(nil)).
			AddHelm(mockserver.New(nil)).
			AddHelm(reorg.New(&reorg.Props{
				NetworkName: sourceNetwork.Name,
				NetworkType: "simulated-geth-non-dev",
				Values: map[string]interface{}{
					"geth": map[string]interface{}{
						"genesis": map[string]interface{}{
							"networkId": fmt.Sprint(sourceNetwork.ChainID),
						},
					},
				},
			})).
			// use blockscout for debugging on-chain transactions
			AddChart(blockscout.New(&blockscout.Props{
				Name:    "dest-blockscout",
				WsURL:   "ws://dest-chain-ethereum-geth:8546",
				HttpURL: "http://dest-chain-ethereum-geth:8544",
			})).
			AddChart(blockscout.New(&blockscout.Props{
				Name:    "source-blockscout",
				WsURL:   "ws://source-chain-ethereum-geth:8546",
				HttpURL: "http://source-chain-ethereum-geth:8544",
			})).
			AddHelm(reorg.New(&reorg.Props{
				NetworkName: destNetwork.Name,
				NetworkType: "simulated-geth-non-dev",
				Values: map[string]interface{}{
					"geth": map[string]interface{}{
						"genesis": map[string]interface{}{
							"networkId": fmt.Sprint(destNetwork.ChainID),
						},
					},
				},
			})).Run()
		Expect(err).ShouldNot(HaveOccurred())
		// related https://app.shortcut.com/chainlinklabs/story/38295/creating-an-evm-chain-via-cli-or-api-immediately-polling-the-nodes-and-returning-an-error
		// node must work and reconnect even if network is not working
		time.Sleep(30 * time.Second)
		err = testEnvironment.AddHelm(chainlink.New(0, map[string]interface{}{
			"replicas": 6,
			"env": map[string]interface{}{
				"FEATURE_CCIP":                   "true",
				"FEATURE_OFFCHAIN_REPORTING2":    "true",
				"feature_offchain_reporting":     "false",
				"FEATURE_LOG_POLLER":             "true",
				"GAS_ESTIMATOR_MODE":             "FixedPrice",
				"P2P_NETWORKING_STACK":           "V2",
				"P2PV2_LISTEN_ADDRESSES":         "0.0.0.0:6690",
				"P2PV2_ANNOUNCE_ADDRESSES":       "0.0.0.0:6690",
				"P2PV2_DELTA_DIAL":               "500ms",
				"P2PV2_DELTA_RECONCILE":          "5s",
				"ETH_GAS_LIMIT_DEFAULT":          "1500000",
				"ETH_LOG_POLL_INTERVAL":          "1s",
				"p2p_listen_port":                "0",
				"ETH_FINALITY_DEPTH":             "50",
				"ETH_HEAD_TRACKER_HISTORY_DEPTH": "100",
				// It is not permitted to set both ETH_URL and EVM_NODES,
				// imposing blank values to stop getting the env variable set as default node set up in qa-charts
				"ETH_URL":      "",
				"ETH_CHAIN_ID": "0",
				"EVM_NODES":    string(evmNodes),
			},
		})).Run()
		Expect(err).ShouldNot(HaveOccurred())

		By("Connecting to launched resources")
		sourceChainClient, err = blockchain.NewEVMClient(networks.NetworkAlpha, testEnvironment)
		Expect(err).ShouldNot(HaveOccurred(), "Connecting to blockchain nodes shouldn't fail")
		destChainClient, err = blockchain.NewEVMClient(networks.NetworkBeta, testEnvironment)
		Expect(err).ShouldNot(HaveOccurred(), "Connecting to blockchain nodes shouldn't fail")

		chainlinkNodes, err = client.ConnectChainlinkNodes(testEnvironment)
		Expect(err).ShouldNot(HaveOccurred(), "Connecting to chainlink nodes shouldn't fail")

		mockServer, err = ctfClient.ConnectMockServer(testEnvironment)
		Expect(err).ShouldNot(HaveOccurred(), "Creating mockserver clients shouldn't fail")

		sourceChainClient.ParallelTransactions(true)
		destChainClient.ParallelTransactions(true)

		By("Funding Chainlink nodes for both the chains")
		err = actions.FundChainlinkNodesForChain(chainlinkNodes, sourceChainClient, big.NewFloat(10))
		Expect(err).ShouldNot(HaveOccurred())
		err = actions.FundChainlinkNodesForChain(chainlinkNodes, destChainClient, big.NewFloat(10))
		Expect(err).ShouldNot(HaveOccurred())

		// create node keys
		_, clNodes, err := client.CreateNodeKeysBundle(chainlinkNodes, "evm", destChainClient.GetChainID().String())
		Expect(err).ShouldNot(HaveOccurred())

		// deploy all source contracts
		sourceCCIP := actions.DefaultSourceCCIPModule(sourceChainClient, destChainClient.GetChainID(), big.NewInt(5e17))
		By("Deploying source contracts")
		sourceCCIP.DeployContracts()

		// deploy all destination contracts
		destCCIP := actions.DefaultDestinationCCIPModule(destChainClient, sourceChainClient.GetChainID(), big.NewInt(5e17))
		By("Deploying destination contracts")
		destCCIP.DeployContracts(*sourceCCIP)

		// Setup mock token fee conversion value
		By("Setup mock token fee conversion value")
		actions.SetMockServerWithSameTokenFeeConversionValue(
			map[string]interface{}{
				destCCIP.Common.LinkToken.Address(): "200000000000000000000",
			}, clNodes[1:], mockServer)

		// set up ocr2 jobs
		By("Setting up bootstrap, relay and execute job")
		actions.CreateOCRJobsForCCIP(
			clNodes, sourceCCIP.OnRamp.Address(),
			destCCIP.BlobVerifier.Address(),
			destCCIP.OffRamp.Address(),
			sourceChainClient, destChainClient,
			destCCIP.Common.LinkToken.Address(),
			mockServer,
		)

		// set up ocr2 config
		By("Setting up ocr config in blob verifier and offramp")
		actions.SetOCRConfigs(clNodes[1:], *destCCIP) // first node is the bootstrapper

		// initiate transfer and verify
		By("Initiate the transfer, watch for updated sequence numbers and events logs, " +
			"verify balance in receiving and sending account pre and post transfer")
		actions.InitiateTokenTransfer(*sourceCCIP, *destCCIP)
	},
		testScenarios,
	)
})
