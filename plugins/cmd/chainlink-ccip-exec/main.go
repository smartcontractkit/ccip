package main

import (
	"github.com/hashicorp/go-plugin"

	"github.com/smartcontractkit/chainlink-common/pkg/loop"
)

const (
	loggerName = "PluginCCIPExec"
)

func main() {
	s := loop.MustNewStartedServer(loggerName)
	defer s.Stop()

	p := NewPlugin(s.Logger)
	defer s.Logger.ErrorIfFn(p.Close, "Failed to close")

	s.MustRegister(p)

	stop := make(chan struct{})
	defer close(stop)

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: loop.PluginCCIPExecutionHandshakeConfig(),
		Plugins: map[string]plugin.Plugin{
			loop.CCIPExecutionLOOPName: &loop.ExecutionLoop{
				PluginServer: p,
				BrokerConfig: loop.BrokerConfig{
					StopCh:   stop,
					Logger:   s.Logger,
					GRPCOpts: s.GRPCOpts,
				},
			},
		},
		GRPCServer: s.GRPCOpts.NewServer,
	})
}
