package load

import (
	"testing"

	"github.com/rs/zerolog/log"
)

func TestLoadCCIPParallelSendSameSenderReceiver(t *testing.T) {
	var (
		testArgs *loadArgs
	)

	t.Cleanup(func() {
		log.Info().Msg("Tearing down the environment")
		testArgs.TearDown()
	})
	testArgs = PopulateAndValidate(t)
	testArgs.Setup()
	testArgs.Run()
}
