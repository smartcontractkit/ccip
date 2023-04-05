package load

import (
	"context"
	"testing"

	"github.com/rs/zerolog/log"
)

func TestLoadCCIPStableRPS(t *testing.T) {
	testArgs := NewLoadArgs(t, context.Background())
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
