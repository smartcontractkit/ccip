package load

import (
	"testing"

	"github.com/rs/zerolog/log"

	"github.com/smartcontractkit/chainlink/integration-tests/testsetups"
)

func TestLoadCCIPParallelSendSameSenderReceiver(t *testing.T) {
	testArgs := &loadArgs{
		t:       t,
		TestCfg: testsetups.NewCCIPTestConfig(t, testsetups.Load),
	}
	testArgs.Setup()
	// if the test runs on remote runner
	if testArgs.ccipLoad == nil {
		return
	}
	t.Cleanup(func() {
		log.Info().Msg("Tearing down the environment")
		testArgs.TestSetupArgs.TearDown()
	})
	testArgs.Run()
}
