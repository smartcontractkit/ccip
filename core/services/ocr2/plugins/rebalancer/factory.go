package rebalancer

import (
	"encoding/json"
	"fmt"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/liquiditygraph"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/liquiditymanager"
	models2 "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

type PluginFactory struct{}

func NewPluginFactory() *PluginFactory {
	return &PluginFactory{}
}

func (p PluginFactory) NewReportingPlugin(config ocr3types.ReportingPluginConfig) (ocr3types.ReportingPlugin[models2.ReportMetadata], ocr3types.ReportingPluginInfo, error) {
	var offchainConfig models2.PluginConfig
	if err := json.Unmarshal(config.OffchainConfig, &offchainConfig); err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("invalid config: %w", err)
	}

	liquidityGraph := liquiditygraph.NewDummyGraph()
	liquidityManagerFactory := liquiditymanager.NewBaseLiquidityManagerFactory()

	return NewPlugin(
			offchainConfig.LiquidityManagerNetwork,
			offchainConfig.LiquidityManagerAddress,
			liquidityManagerFactory,
			liquidityGraph,
		),
		ocr3types.ReportingPluginInfo{
			Name: models2.PluginName,
			Limits: ocr3types.ReportingPluginLimits{
				MaxQueryLength:       0,
				MaxObservationLength: 0,
				MaxOutcomeLength:     0,
				MaxReportLength:      0,
				MaxReportCount:       0,
			},
		},
		nil
}
