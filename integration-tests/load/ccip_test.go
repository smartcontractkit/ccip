package load

import (
	"context"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/chainlink-testing-framework/utils"
)

func TestLoadCCIPStableRPS(t *testing.T) {
	lggr := utils.GetTestLogger(t)
	testArgs := NewLoadArgs(t, lggr, context.Background())
	testArgs.SetupStableLoad()
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
