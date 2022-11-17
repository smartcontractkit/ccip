package load

import (
	"os"
	"strconv"
	"time"

	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/smartcontractkit/chainlink-testing-framework/client"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
)

type loadArgs struct {
	rps         int
	duration    time.Duration
	ccipTimeout time.Duration
	loadTimeOut time.Duration
	msgType     string
	envTear     func()
	ccipLoad    *CCIPE2ELoad
	loadGen     *client.LoadGenerator
}

// PopulateAndValidate collects all loadArgs
func PopulateAndValidate() *loadArgs {
	inputRps := os.Getenv("LOAD_TEST_TPS")
	inputTimeout := os.Getenv("LOAD_TEST_CALLTIMEOUT")
	inputDuration := os.Getenv("LOAD_TEST_DURATION")
	inputMsgType := os.Getenv("LOAD_TEST_MSG_TYPE")
	p := &loadArgs{
		rps:         4,
		duration:    10 * time.Minute,
		ccipTimeout: 5 * time.Minute,
		loadTimeOut: 18 * time.Minute,
		msgType:     DataOnlyTransfer,
	}
	if inputRps != "" {
		rps, err := strconv.Atoi(inputRps)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(rps).Should(BeNumerically("<", 16), "rps too high")
		p.rps = rps
	}
	if inputTimeout != "" {
		d, err := strconv.Atoi(inputTimeout)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(d).Should(BeNumerically(">", 1), "invalid timeout")
		Expect(d).Should(BeNumerically("<", 10), "invalid timeout")
		p.ccipTimeout = time.Duration(d) * time.Minute
		p.loadTimeOut = time.Duration(d*3) * time.Minute
	}
	if inputDuration != "" {
		d, err := strconv.Atoi(inputDuration)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(d).Should(BeNumerically("<", 90), "invalid duration")
		p.duration = time.Duration(d) * time.Minute
	}
	if inputMsgType != "" {
		Expect(inputMsgType).Should(BeElementOf([]string{DataOnlyTransfer, TokenTransfer}), "invalid msg type")
		p.msgType = inputMsgType
	}
	return p
}

func (loadArgs *loadArgs) Setup() {
	_, source, dest, _, tearDown := actions.CCIPDefaultTestSetUp("load-ccip",
		map[string]interface{}{
			"replicas": "6",
			"toml":     actions.DefaultCCIPCLNodeEnv(),
			"env": map[string]interface{}{
				"CL_DEV": "true",
			},
		}, 5, true)
	loadArgs.envTear = tearDown
	ccipLoad := NewCCIPLoad(source, dest, actions.SUB, loadArgs.ccipTimeout, 100000)
	ccipLoad.BeforeAllCall()
	loadgen, err := client.NewLoadGenerator(&client.LoadGeneratorConfig{
		RPS:         loadArgs.rps,
		Gun:         ccipLoad,
		Duration:    loadArgs.duration,
		CallTimeout: loadArgs.loadTimeOut,
		SharedData:  loadArgs.msgType,
	})
	Expect(err).ShouldNot(HaveOccurred(), "initiating loadgen")
	loadArgs.ccipLoad = ccipLoad
	loadArgs.loadGen = loadgen
}

func (loadArgs *loadArgs) Run() {
	loadArgs.loadGen.Run()
	_, failed := loadArgs.loadGen.Wait()
	Expect(failed).Should(BeFalse(), "load run is failed")
}

func (loadArgs *loadArgs) TearDown() {
	defer loadArgs.envTear()
	testFailed := ginkgo.CurrentSpecReport().Failed()
	loadArgs.ccipLoad.PrintStats(testFailed, loadArgs.rps, loadArgs.duration.Minutes())
}
