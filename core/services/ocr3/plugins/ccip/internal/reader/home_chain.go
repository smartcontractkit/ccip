package reader

import (
	"context"
	"log"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	cciptypes "github.com/smartcontractkit/ccipocr3/ccipocr3-dont-merge"
	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

type HomeChainConfigPoller struct {
	homeChainReader types.ContractReader
	homeChainConfig cciptypes.HomeChainConfig
	p2pIdToOracleId map[cciptypes.Bytes32]commontypes.OracleID
}

func (r *HomeChainConfigPoller) StartPolling(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	log.Println("Start Polling HomeChainConfig")
	for {
		select {
		case <-ctx.Done():
			log.Println("Polling stopped")
			return
		case <-ticker.C:
			homeConfig, err := r.fetchLatestConfig(ctx)
			if err != nil {
				log.Println("Error polling DB:", err)
			}
			r.homeChainConfig = homeConfig
		}
	}
}

func (r *HomeChainConfigPoller) GetConfig() cciptypes.HomeChainConfig {
	return r.homeChainConfig
}

func (r *HomeChainConfigPoller) GetSupportedChains(oracleID commontypes.OracleID) mapset.Set[cciptypes.ChainSelector] {
	return r.homeChainConfig.NodeSupportedChains[oracleID].Supported
}

func (r *HomeChainConfigPoller) fetchLatestConfig(ctx context.Context) (cciptypes.HomeChainConfig, error) {
	println("Fetching latest config")
	var onChainCapabilityConfig []cciptypes.OnChainCapabilityConfig
	err := r.homeChainReader.GetLatestValue(ctx, "CCIPCapabilityConfiguration", "getAllChainConfigs", nil, &onChainCapabilityConfig)
	if err != nil {
		return cciptypes.HomeChainConfig{}, err
	}

	homeChainConfig, err := r.convertOnChainConfigToHomeChainConfig(onChainCapabilityConfig)
	if err != nil {
		return cciptypes.HomeChainConfig{}, err
	}
	r.homeChainConfig = homeChainConfig
	return r.homeChainConfig, err
}

func (r *HomeChainConfigPoller) Close(ctx context.Context) error {
	return nil
}

func (r *HomeChainConfigPoller) convertOnChainConfigToHomeChainConfig(capabilityConfigs []cciptypes.OnChainCapabilityConfig) (cciptypes.HomeChainConfig, error) {
	var fChain = make(map[cciptypes.ChainSelector]int)
	// NodeSupportedChains is a map of oracle IDs to SupportedChains.
	var nodeSupportedChains = make(map[commontypes.OracleID]cciptypes.SupportedChains)
	//iterate over configs
	for _, capabilityConfig := range capabilityConfigs {
		//iterate over readers
		chainSelector := cciptypes.ChainSelector(capabilityConfig.ChainSelector)
		config := capabilityConfig.ChainConfig

		fChain[chainSelector] = int(config.FChain)
		for _, p2pId := range config.Readers {
			oracleID := r.p2pIdToOracleId[p2pId]
			if _, ok := nodeSupportedChains[oracleID]; !ok {
				nodeSupportedChains[oracleID] = cciptypes.SupportedChains{
					Supported: mapset.NewSet[cciptypes.ChainSelector](),
				}
			}
			//add chain to SupportedChains
			nodeSupportedChains[oracleID].Supported.Add(chainSelector)
		}
	}
	return cciptypes.HomeChainConfig{
		FChain:              fChain,
		NodeSupportedChains: nodeSupportedChains,
	}, nil
}
