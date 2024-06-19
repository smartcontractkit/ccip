package reader

import (
	"context"
	"errors"
	"sync"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type HomeChainConfigPoller struct {
	homeChainReader  types.ContractReader
	homeChainConfig  cciptypes.HomeChainConfig
	lggr             logger.Logger
	mutex            sync.RWMutex
	backgroundCancel context.CancelFunc
}

func NewHomeChainConfigPoller(
	homeChainReader types.ContractReader,
	lggr logger.Logger,
) *HomeChainConfigPoller {
	return &HomeChainConfigPoller{
		homeChainReader: homeChainReader,
		lggr:            lggr,
	}
}

func (r *HomeChainConfigPoller) Start(ctx context.Context) error {

	r.mutex.Lock()
	if r.backgroundCancel != nil {
		r.lggr.Errorw("Polling already started")
		// We don't want to return an actual error here as it's already working as expected
		return nil
	}
	bgCtx, cancelFunc := context.WithCancel(ctx)
	r.backgroundCancel = cancelFunc
	r.mutex.Unlock()
	r.lggr.Infow("Start Polling HomeChainConfig")
	go func() {
		ticker := time.NewTicker(12 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				r.lggr.Infow("Polling stopped")
				return
			case <-bgCtx.Done():
				r.lggr.Infow("Polling stopped")
				return
			case <-ticker.C:
				onChainCapabilityConfigs, err := r.fetchOnChainConfig(ctx)
				if err != nil {
					r.lggr.Errorw("Fetching on chain configs failed", "err", err)
					continue
				}
				if onChainCapabilityConfigs == nil || len(onChainCapabilityConfigs) == 0 {
					r.lggr.Errorw("No on chain configs found")
					continue
				}
				homeChainConfig, err := r.convertOnChainConfigToHomeChainConfig(onChainCapabilityConfigs)
				if err != nil {
					r.lggr.Errorw("error converting OnChainConfigs to HomeChainConfig", "err", err)
					continue
				}
				r.lggr.Infow("Setting HomeChainConfig")
				r.mutex.Lock()
				r.homeChainConfig = homeChainConfig
				r.mutex.Unlock()
			}
		}
	}()

	return nil
}

func (r *HomeChainConfigPoller) GetConfig() cciptypes.HomeChainConfig {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.homeChainConfig
}

func (r *HomeChainConfigPoller) GetSupportedChains(p2pID cciptypes.P2PID) mapset.Set[cciptypes.ChainSelector] {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.homeChainConfig.NodeSupportedChains[p2pID].Supported
}

func (r *HomeChainConfigPoller) fetchOnChainConfig(ctx context.Context) ([]cciptypes.OnChainCapabilityConfig, error) {
	r.lggr.Infow("Fetching OnChainConfig")
	var onChainCapabilityConfig []cciptypes.OnChainCapabilityConfig
	err := r.homeChainReader.GetLatestValue(ctx, "CCIPCapabilityConfiguration", "getAllChainConfigs", nil, &onChainCapabilityConfig)
	if err != nil {
		return nil, err
	}
	return onChainCapabilityConfig, err
}

func (r *HomeChainConfigPoller) Close() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if r.backgroundCancel == nil {
		return errors.New("Cancel function not set")

	}
	r.backgroundCancel()
	return nil
}

func (r *HomeChainConfigPoller) convertOnChainConfigToHomeChainConfig(capabilityConfigs []cciptypes.OnChainCapabilityConfig) (cciptypes.HomeChainConfig, error) {
	fChain := make(map[cciptypes.ChainSelector]int)
	// NodeSupportedChains is a map of oracle IDs to SupportedChains.
	var nodeSupportedChains = make(map[cciptypes.P2PID]cciptypes.SupportedChains)
	//iterate over configs
	for _, capabilityConfig := range capabilityConfigs {
		chainSelector := capabilityConfig.ChainSelector
		config := capabilityConfig.ChainConfig

		fChain[chainSelector] = int(config.FChain)
		//iterate over readers
		for _, p2pID := range config.Readers {
			if _, ok := nodeSupportedChains[p2pID]; !ok {
				nodeSupportedChains[p2pID] = cciptypes.SupportedChains{
					Supported: mapset.NewSet[cciptypes.ChainSelector](),
				}
			}
			//add chain to SupportedChains
			nodeSupportedChains[p2pID].Supported.Add(chainSelector)
		}
	}
	return cciptypes.HomeChainConfig{
		FChain:              fChain,
		NodeSupportedChains: nodeSupportedChains,
	}, nil
}
