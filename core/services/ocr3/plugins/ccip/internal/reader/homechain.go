package reader

import (
	"context"
	"time"

	cciptypes "github.com/smartcontractkit/ccipocr3/ccipocr3-dont-merge"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

type CCIPHomeChainReader struct {
	homeChainReader types.ContractReader
}

type OnChainCapabilityConfig struct {
	// TODO: map to the actual contract ChainConfig
}

func (r *CCIPHomeChainReader) FetchLatestConfig(ctx context.Context) (cciptypes.HomeChainConfig, error) {
	var onChainCapabilityConfig OnChainCapabilityConfig
	err := r.homeChainReader.GetLatestValue(ctx, "CCIPCapabilityConfiguration", "getAllChainConfigs", nil, &onChainCapabilityConfig)
	if err != nil {
		return cciptypes.HomeChainConfig{}, err
	}

	homeChainConfig, err := r.convertOnChainConfigToHomeChainConfig(onChainCapabilityConfig)
	if err != nil {
		return cciptypes.HomeChainConfig{}, err
	}
	return homeChainConfig, err
}

func (r *CCIPHomeChainReader) StartConfigAutoUpdate(ctx context.Context, config *cciptypes.HomeChainConfig, interval time.Duration) error {
	panic("implement me")
}

func (r *CCIPHomeChainReader) Close(ctx context.Context) error {
	return nil
}

func (r *CCIPHomeChainReader) convertOnChainConfigToHomeChainConfig(config OnChainCapabilityConfig) (cciptypes.HomeChainConfig, error) {
	panic("implement me")
}
