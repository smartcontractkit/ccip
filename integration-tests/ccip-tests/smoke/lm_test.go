package smoke

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-testing-framework/logging"
	"github.com/smartcontractkit/chainlink/integration-tests/ccip-tests/testconfig"
	"github.com/smartcontractkit/chainlink/integration-tests/ccip-tests/testsetups"
)

func TestLmBasic(t *testing.T) {
	t.Parallel()
	log := logging.GetTestLogger(t)
	TestCfg := testsetups.NewCCIPTestConfig(t, log, testconfig.Smoke)
	require.NotNil(t, TestCfg.TestGroupInput.MsgDetails.DestGasLimit)
	//gasLimit := big.NewInt(*TestCfg.TestGroupInput.MsgDetails.DestGasLimit)
	_ = testsetups.LMDefaultTestSetup(t, log, "smoke-lm", TestCfg)
}
