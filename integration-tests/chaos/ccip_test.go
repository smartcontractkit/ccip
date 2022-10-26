package chaos_test

//revive:disable:dot-imports
import (
	"math/big"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/chainlink-env/chaos"
	"github.com/smartcontractkit/chainlink-env/environment"
	a "github.com/smartcontractkit/chainlink-env/pkg/alias"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"

	networks "github.com/smartcontractkit/chainlink/integration-tests"
	"github.com/smartcontractkit/chainlink/integration-tests/actions"
)

var _ = Describe("CCIP chaos test @chaos-ccip", func() {
	var (
		testScenarios = []TableEntry{
			Entry("Must catch up pending requests if majority of relay nodes goes down for 1m @chaos-ccip-fail-relay-majority",
				networks.NetworkAlpha,
				networks.NetworkBeta,
				chaos.NewFailPods,
				&chaos.Props{
					LabelsSelector: &map[string]*string{actions.ChaosGroupRelayFaultyPlus: a.Str("1")},
					DurationStr:    "1m",
				}, true,
				big.NewInt(1.78e18),
			),
			// TODO bug - https://app.shortcut.com/chainlinklabs/story/55732
			// All parallel execution fails
			PEntry("Must catch up pending requests if majority of execution nodes goes down for 1m @chaos-ccip-fail-execution-majority",
				networks.NetworkAlpha,
				networks.NetworkBeta,
				chaos.NewFailPods,
				&chaos.Props{
					LabelsSelector: &map[string]*string{actions.ChaosGroupExecutionFaultyPlus: a.Str("1")},
					DurationStr:    "1m",
				}, true,
				big.NewInt(2.41e18),
			),
			PEntry("Must continue ocr2 if minority of relay nodes get killed @chaos-ccip-fail-relay-minority",
				networks.NetworkAlpha,
				networks.NetworkBeta,
				chaos.NewFailPods,
				&chaos.Props{
					LabelsSelector: &map[string]*string{actions.ChaosGroupRelayFaulty: a.Str("1")},
					DurationStr:    "3m",
				}, false,
				big.NewInt(1.78e18),
			),
			PEntry("Must continue ocr2 if minority of execution nodes get killed @chaos-ccip-fail-execution-minority",
				networks.NetworkAlpha,
				networks.NetworkBeta,
				chaos.NewFailPods,
				&chaos.Props{
					LabelsSelector: &map[string]*string{actions.ChaosGroupExecutionFaulty: a.Str("1")},
					DurationStr:    "2m",
				}, false,
				big.NewInt(2.41e18),
			),
		}
		tearDown         func()
		testEnvironment  *environment.Environment
		source           *actions.SourceCCIPModule
		dest             *actions.DestCCIPModule
		testSetup        actions.CCIPTestEnv
		numOfSubRequests = 5
		numOfRelayNodes  = 5
	)

	AfterEach(func() {
		tearDown()
	})

	DescribeTable("OCR chaos on different EVM networks", func(
		sourceNetwork *blockchain.EVMNetwork,
		destNetwork *blockchain.EVMNetwork,
		chaosFunc chaos.ManifestFunc,
		chaosProps *chaos.Props,
		waitForChaosRecovery bool,
		subCost *big.Int,
	) {
		By("Deploy and set up default chaos test environment")
		testEnvironment, source, dest, testSetup, tearDown = actions.CCIPDefaultTestSetUp(sourceNetwork, destNetwork, "chaos-ccip",
			map[string]interface{}{
				"replicas": "12",
				"env":      actions.DefaultCCIPCLNodeEnv,
				"db": map[string]interface{}{
					"stateful": true,
					"capacity": "10Gi",
					"resources": map[string]interface{}{
						"requests": map[string]interface{}{
							"cpu":    "250m",
							"memory": "256Mi",
						},
						"limits": map[string]interface{}{
							"cpu":    "250m",
							"memory": "256Mi",
						},
					},
				},
			}, numOfRelayNodes, false)

		testSetup.ChaosLabel()

		cTest := actions.NewCCIPTest(source, dest, big.NewInt(0).Mul(big.NewInt(80), big.NewInt(1e18)), subCost, 1*time.Minute)
		// apply chaos
		chaosId, err := testEnvironment.Chaos.Run(chaosFunc(testEnvironment.Cfg.Namespace, chaosProps))
		Expect(err).ShouldNot(HaveOccurred())
		// Send the ccip-request while the chaos is at play
		cTest.SendSubRequests(numOfSubRequests)
		if waitForChaosRecovery {
			// wait for chaos to be recovered before further validation
			testEnvironment.Chaos.WaitForAllRecovered(chaosId)
		} else {
			log.Info().Msg("proceeding without waiting for chaos recovery")
		}
		cTest.ValidateSubRequests()
	},
		testScenarios,
	)
})
