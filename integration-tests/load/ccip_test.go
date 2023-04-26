package load

import (
	"context"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/chainlink-testing-framework/utils"
	"github.com/smartcontractkit/wasp"
)

func TestLoadCCIPStableRPS(t *testing.T) {
	t.Parallel()
	lggr := utils.GetTestLogger(t)
	testArgs := NewLoadArgs(t, lggr, context.Background())
	testArgs.Setup()
	// if the test runs on remote runner
	if testArgs.TestSetupArgs.Env != nil && testArgs.TestSetupArgs.Env.K8Env.WillUseRemoteRunner() {
		return
	}
	t.Cleanup(func() {
		log.Info().Msg("Tearing down the environment")
		testArgs.TestSetupArgs.TearDown()
	})
	testArgs.TriggerLoad()
	testArgs.Wait()
}

func TestLoadCCIPSequentialLaneAdd(t *testing.T) {
	t.Parallel()
	lggr := utils.GetTestLogger(t)
	testArgs := NewLoadArgs(t, lggr, context.Background())
	testArgs.TestCfg.SequentialLaneAddition = true
	if len(testArgs.TestCfg.NetworkPairs) <= 1 {
		t.Skip("Skipping the test as there are not enough network pairs to run the test")
	}
	testArgs.Setup()
	// if the test runs on remote runner
	if testArgs.TestSetupArgs.Env != nil && testArgs.TestSetupArgs.Env.K8Env.WillUseRemoteRunner() {
		return
	}
	t.Cleanup(func() {
		log.Info().Msg("Tearing down the environment")
		testArgs.TestSetupArgs.TearDown()
	})
	testArgs.TriggerLoad()
	testArgs.AddMoreLanesToRun()
	testArgs.Wait()
}

func TestLoadCCIPIncrementalLoad(t *testing.T) {
	t.Parallel()
	lggr := utils.GetTestLogger(t)
	testArgs := NewLoadArgs(t, lggr, context.Background())
	testArgs.TestCfg.SequentialLaneAddition = true
	testArgs.Setup()
	// if the test runs on remote runner
	if testArgs.TestSetupArgs.Env != nil && testArgs.TestSetupArgs.Env.K8Env.WillUseRemoteRunner() {
		return
	}
	t.Cleanup(func() {
		log.Info().Msg("Tearing down the environment")
		testArgs.TestSetupArgs.TearDown()
	})
	stepDuration := testArgs.TestCfg.TestDuration / 3
	schedules := wasp.CombineAndRepeat(3,
		wasp.Line(1, testArgs.TestCfg.Load.LoadRPS, stepDuration))
	testArgs.TriggerLoad(schedules...)
	testArgs.Wait()
}
