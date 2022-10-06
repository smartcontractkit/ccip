package load

import (
	"testing"
	"time"

	"github.com/onsi/gomega"
	"github.com/smartcontractkit/chainlink-testing-framework/client"
	"github.com/smartcontractkit/chainlink-testing-framework/logging"
	"github.com/stretchr/testify/require"

	networks "github.com/smartcontractkit/chainlink/integration-tests"
	"github.com/smartcontractkit/chainlink/integration-tests/actions"
)

func TestCCIPLoadParallelSendBetweenSameSenderAndReceiver(t *testing.T) {
	sourceNetwork := networks.NetworkAlpha
	destNetwork := networks.NetworkBeta
	gomega.RegisterTestingT(t)
	logging.Init()
	source, dest, tearDown := actions.CCIPDefaultTestSetUpForLoad(sourceNetwork, destNetwork, "load-ccip-sim-geth")
	ccipLoad := NewCCIPLoad(source, dest, actions.SUB, 5*time.Minute, 100000)
	defer func() {
		tearDown()
		ccipLoad.PrintStats(t)
	}()
	ccipLoad.BeforeAllCall()
	interval := 10 * time.Second
	tests := []struct {
		name        string
		rps         int
		duration    time.Duration
		callTimeout time.Duration
		ccipTimeout time.Duration
		msgType     string
	}{
		// TODO - https://app.shortcut.com/chainlinklabs/story/54689
		//{"RPS=1,Duration=10s,TokenTransfer", 1, 10 * time.Second, 10 * time.Minute, 5 * time.Minute, TokenTransfer},
		{"RPS=2,Duration=10m,DataOnlyTransfer", 2, 10 * time.Minute, 10 * time.Minute, 3 * time.Minute, DataOnlyTransfer},
		{"RPS=3,Duration=10m,DataOnlyTransfer", 3, 10 * time.Minute, 10 * time.Minute, 3 * time.Minute, DataOnlyTransfer},
		{"RPS=4,Duration=10m,DataOnlyTransfer", 4, 10 * time.Minute, 25 * time.Minute, 3 * time.Minute, DataOnlyTransfer},
	}
	for _, params := range tests {
		t.Run(params.name, func(t *testing.T) {
			ccipLoad.CallTimeOut = params.ccipTimeout
			loadgen, err := client.NewLoadGenerator(&client.LoadGeneratorConfig{
				RPS:                  params.rps,
				Gun:                  ccipLoad,
				Duration:             params.duration,
				CallTimeout:          params.callTimeout,
				CallFailThreshold:    10,
				CallTimeoutThreshold: 5,
				SharedData:           params.msgType,
			})

			require.NoError(t, err)
			loadgen.Run()
			_, failed := loadgen.Wait()
			require.False(t, failed)
			time.Sleep(interval)
		})
	}
}
