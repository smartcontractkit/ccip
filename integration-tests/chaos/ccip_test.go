package chaos_test

//revive:disable:dot-imports
import (
	"math/big"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/smartcontractkit/chainlink-env/environment"

	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/chainlink-env/chaos"
	a "github.com/smartcontractkit/chainlink-env/pkg/alias"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
)

var _ = Describe("CCIP chaos test @chaos-ccip", Ordered, func() {
	var (
		tearDown         func()
		numOfCommitNodes = 5
		testEnvironment  *environment.Environment
		source           *actions.SourceCCIPModule
		dest             *actions.DestCCIPModule
		testSetup        actions.CCIPTestEnv
		totalReq         = 20
	)

	AfterAll(func() {
		tearDown()
	})

	BeforeAll(func() {
		testEnvironment, source, dest, testSetup, tearDown = actions.CCIPDefaultTestSetUp("chaos-ccip",
			map[string]interface{}{
				"replicas": "12",
				"env":      actions.DefaultCCIPCLNodeEnv(),
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
			}, numOfCommitNodes, false)
		actions.CreateAndFundSubscription(*source, *dest, big.NewInt(0).Mul(big.NewInt(100), big.NewInt(1e18)), int64(totalReq))
	})

	Describe("CCIP chaos @chaos-ccip-ocr", func() {
		var (
			testScenarios = []TableEntry{
				Entry("Must catch up pending requests if majority of commit nodes goes down for 1m @chaos-ccip-fail-commit-majority",
					chaos.NewFailPods,
					&chaos.Props{
						LabelsSelector: &map[string]*string{actions.ChaosGroupCommitFaultyPlus: a.Str("1")},
						DurationStr:    "1m",
					}, true,
					big.NewInt(2e18),
				),
				Entry("Must catch up pending requests if majority of execution nodes goes down for 1m @chaos-ccip-fail-execution-majority",
					chaos.NewFailPods,
					&chaos.Props{
						LabelsSelector: &map[string]*string{actions.ChaosGroupExecutionFaultyPlus: a.Str("1")},
						DurationStr:    "1m",
					}, true,
					big.NewInt(2.5e18),
				),
				Entry("Must continue ocr2 if minority of commit nodes get killed @chaos-ccip-fail-commit-minority",
					chaos.NewFailPods,
					&chaos.Props{
						LabelsSelector: &map[string]*string{actions.ChaosGroupCommitFaulty: a.Str("1")},
						DurationStr:    "90s",
					}, false,
					big.NewInt(2.3e18),
				),
				Entry("Must continue ocr2 if minority of execution nodes get killed @chaos-ccip-fail-execution-minority",
					chaos.NewFailPods,
					&chaos.Props{
						LabelsSelector: &map[string]*string{actions.ChaosGroupExecutionFaulty: a.Str("1")},
						DurationStr:    "90s",
					}, false,
					big.NewInt(2.3e18),
				),
			}
			numOfSubRequests = 5
			chaosId          string
			err              error
		)
		AfterEach(func() {
			if chaosId != "" {
				testEnvironment.Chaos.Stop(chaosId)
			}
		})

		DescribeTable("CCIP chaos scenarios with multiple sequential requests", func(
			chaosFunc chaos.ManifestFunc,
			chaosProps *chaos.Props,
			waitForChaosRecovery bool,
			subCost *big.Int,
		) {
			testSetup.ChaosLabel()
			cTest := actions.NewCCIPTest(source, dest, big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18)), subCost, 1*time.Minute)
			// apply chaos
			chaosId, err = testEnvironment.Chaos.Run(chaosFunc(testEnvironment.Cfg.Namespace, chaosProps))
			Expect(err).ShouldNot(HaveOccurred())
			// Send the ccip-request while the chaos is at play
			cTest.SendSubRequests(numOfSubRequests, false)
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
})
