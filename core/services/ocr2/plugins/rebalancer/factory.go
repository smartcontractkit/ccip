package rebalancer

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/liquiditygraph"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/liquiditymanager"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/liquidityrebalancer"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

const (
	PluginName = "Rebalancer"
)

type PluginFactory struct {
	lggr      logger.Logger
	config    models.PluginConfig
	lmFactory liquiditymanager.Factory
}

func NewPluginFactory(lggr logger.Logger, pluginConfigBytes []byte, lmFactory liquiditymanager.Factory) (*PluginFactory, error) {
	var pluginConfig models.PluginConfig
	if err := json.Unmarshal(pluginConfigBytes, &pluginConfig); err != nil {
		return nil, fmt.Errorf("failed to unmarshal plugin config: %w", err)
	}
	return &PluginFactory{
		lggr:      lggr.Named(PluginName),
		config:    pluginConfig,
		lmFactory: lmFactory,
	}, nil
}

func (p PluginFactory) buildRebalancer() (liquidityrebalancer.Rebalancer, error) {
	switch p.config.RebalancerConfig.Type {
	case models.RebalancerTypeDummy:
		return liquidityrebalancer.NewDummyRebalancer(), nil
	case models.RebalancerTypeRandom:
		return liquidityrebalancer.NewRandomRebalancer(
			p.config.RebalancerConfig.RandomRebalancerConfig.MaxNumTransfers,
			p.config.RebalancerConfig.RandomRebalancerConfig.CheckSourceDestEqual,
			p.lggr), nil
	default:
		return nil, fmt.Errorf("invalid rebalancer type %s", p.config.RebalancerConfig.Type)
	}
}

func (p PluginFactory) NewReportingPlugin(config ocr3types.ReportingPluginConfig) (ocr3types.ReportingPlugin[models.ReportMetadata], ocr3types.ReportingPluginInfo, error) {
	liquidityRebalancer, err := p.buildRebalancer()
	if err != nil {
		return nil, ocr3types.ReportingPluginInfo{}, fmt.Errorf("failed to build rebalancer: %w", err)
	}

	liquidityGraph := liquiditygraph.NewGraph()

	closePluginTimeout := 30 * time.Second
	if p.config.ClosePluginTimeoutSec > 0 {
		closePluginTimeout = time.Duration(p.config.ClosePluginTimeoutSec) * time.Second
	}

	return NewPlugin(
			config.F,
			closePluginTimeout,
			p.config.LiquidityManagerNetwork,
			p.config.LiquidityManagerAddress,
			p.lmFactory,
			liquidityGraph,
			liquidityRebalancer,
			p.lggr,
		),
		ocr3types.ReportingPluginInfo{
			Name: models.PluginName,
			Limits: ocr3types.ReportingPluginLimits{
				MaxQueryLength:       maxQueryLength,
				MaxObservationLength: maxObservationLength,
				MaxOutcomeLength:     maxOutcomeLength,
				MaxReportLength:      maxReportLength,
				MaxReportCount:       maxReportCount,
			},
		},
		nil
}

const (
	// maxQueryLength should be 0 as no queries are performed
	maxQueryLength = 0
	// maxObservationLength should be 10 kilobytes
	maxObservationLength = 10 * 1024
	// maxOutcomeLength should be 10 kilobytes
	maxOutcomeLength = 10 * 1024
	// maxReportLength should be 10 kilobytes
	maxReportLength = 10 * 1024
	// maxReportCount should be 100
	maxReportCount = 100
)
