package ccipcapability

import (
	"context"

	"github.com/smartcontractkit/chainlink/v2/core/services"
	cctypes "github.com/smartcontractkit/chainlink/v2/core/services/ccipcapability/types"
)

var _ cctypes.HomeChainReader = (*homeChainReader)(nil)
var _ services.ServiceCtx = (*homeChainReader)(nil)

type homeChainReader struct{}

// Close implements services.Service.
func (h *homeChainReader) Close() error {
	panic("unimplemented")
}

// HealthReport implements services.Service.
func (h *homeChainReader) HealthReport() map[string]error {
	panic("unimplemented")
}

// Name implements services.Service.
func (h *homeChainReader) Name() string {
	panic("unimplemented")
}

// Ready implements services.Service.
func (h *homeChainReader) Ready() error {
	panic("unimplemented")
}

// Start implements services.Service.
func (h *homeChainReader) Start(context.Context) error {
	panic("unimplemented")
}

// GetAllChainConfigs implements HomeChainReader.
func (h *homeChainReader) GetAllChainConfigs(ctx context.Context) (map[uint64]cctypes.ChainConfig, error) {
	panic("unimplemented")
}

// GetOCRConfigs implements HomeChainReader.
func (h *homeChainReader) GetOCRConfigs(ctx context.Context, donID uint32, pluginType cctypes.PluginType) ([]cctypes.OCRConfig, error) {
	panic("unimplemented")
}

// IsHealthy implements HomeChainReader.
func (h *homeChainReader) IsHealthy() bool {
	panic("unimplemented")
}
