package main

import (
	"context"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/ccipexec"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/loop"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

type Plugin struct {
	loop.Plugin
	stop services.StopChan
}

func NewPlugin(lggr logger.Logger) *Plugin {
	return &Plugin{Plugin: loop.Plugin{Logger: lggr}, stop: make(services.StopChan)}
}

func (p *Plugin) NewExecutionFactory(ctx context.Context, srcProvider types.CCIPExecProvider, dstProvider types.CCIPExecProvider, srcChainID int64, dstChainID int64, sourceTokenAddress string) (types.ReportingPluginFactory, error) {
	return ccipexec.NewExecutionReportingPluginFactoryV2(ctx, nil, sourceTokenAddress, srcChainID, dstChainID, srcProvider, dstProvider)
}
