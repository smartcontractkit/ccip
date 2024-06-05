package reader

import (
	"context"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/commontypes"
)

type HomeChainReader interface {
	FetchLatestConfig(ctx context.Context) (*HomeChainConfig, error)
	StartConfigAutoUpdate(ctx context.Context, config *HomeChainConfig, interval time.Duration) error
	// Close closes any open resources.
	Close(ctx context.Context) error
}

type CCIPHomeChainReader struct {
	homeChainReader types.ContractReader
}

type OnChainCapabilityConfig struct {
	// TODO: map to the actual contract ChainConfig
}

func (r *CCIPHomeChainReader) FetchLatestConfig(ctx context.Context) (*HomeChainConfig, error) {
	var onChainCapabilityConfig OnChainCapabilityConfig
	err := r.homeChainReader.GetLatestValue(ctx, "CCIPCapabilityConfiguration", "getAllChainConfigs", nil, &onChainCapabilityConfig)
	if err != nil {
		return nil, err
	}

	homeChainConfig, err := r.convertOnChainConfigToHomeChainConfig(onChainCapabilityConfig)
	if err != nil {
		return nil, err
	}
	return &homeChainConfig, err
}

func (r *CCIPHomeChainReader) StartConfigAutoUpdate(ctx context.Context, config *HomeChainConfig, interval time.Duration) error {
	panic("implement me")
}

func (r *CCIPHomeChainReader) Close(ctx context.Context) error {
	return nil
}

func (r *CCIPHomeChainReader) convertOnChainConfigToHomeChainConfig(config OnChainCapabilityConfig) (HomeChainConfig, error) {
	panic("implement me")
}

type HomeChainConfig struct {
	// FChain defines the FChain value for each chain. FChain is used while forming consensus based on the observations.
	FChain map[cciptypes.ChainSelector]int `json:"fChain"`
	// ObserverInfo is a map of oracle IDs to ObserverInfo.
	NodeSupportedChains map[commontypes.OracleID]SupportedChains `json:"nodeSupportedChains"`
}

func (c *HomeChainConfig) GetFChain(chain cciptypes.ChainSelector) int {
	return c.FChain[chain]
}

func (c *HomeChainConfig) IsSupported(node commontypes.OracleID, chain cciptypes.ChainSelector) bool {
	supportedChains, ok := c.NodeSupportedChains[node]
	if !ok {
		return false
	}
	return supportedChains.IsSupported(chain)
}

func (c *HomeChainConfig) GetSupportedChains(node commontypes.OracleID) mapset.Set[cciptypes.ChainSelector] {
	supportedChains, ok := c.NodeSupportedChains[node]
	if !ok {
		return mapset.NewSet[cciptypes.ChainSelector]()
	}
	return supportedChains.Supported
}

type SupportedChains struct {
	Supported mapset.Set[cciptypes.ChainSelector] `json:"supported"`
}

func (supportedChains *SupportedChains) IsSupported(chain cciptypes.ChainSelector) bool {
	return supportedChains.Supported.Contains(chain)
}
