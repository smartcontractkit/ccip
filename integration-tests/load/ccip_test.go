package load

import (
	"testing"

	"github.com/rs/zerolog/log"
)

func TestLoadCCIPParallelSendSameSenderReceiver(t *testing.T) {
	var (
		testArgs *loadArgs
	)

	testArgs = PopulateAndValidate(t)
	testArgs.Setup()
	// if the test runs on remote runner
	if testArgs.ccipLoad == nil {
		return
	}
	t.Cleanup(func() {
		log.Info().Msg("Tearing down the environment")
		testArgs.TearDown()
	})
	testArgs.Run()
}

func TestExistingDeploymentLoadCCIPParallelSendSameSenderReceiver(t *testing.T) {
	var (
		testArgs *loadArgs
	)

	testArgs = PopulateAndValidate(t)
	testArgs.ExistingDeployment = true
	testArgs.Setup()
	// if the test runs on remote runner
	if testArgs.ccipLoad == nil {
		return
	}
	t.Cleanup(func() {
		log.Info().Msg("Tearing down the environment")
		testArgs.TearDown()
	})
	testArgs.Run()
}
