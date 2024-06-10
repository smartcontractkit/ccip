package reader

import (
	"context"
	"errors"
	"sync"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	cciptypes "github.com/smartcontractkit/ccipocr3/ccipocr3-dont-merge"
	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

type HomeChainConfigPoller struct {
	homeChainReader  types.ContractReader
	homeChainConfig  cciptypes.HomeChainConfig
	lggr             logger.Logger
	mutex            sync.RWMutex
	backgroundCancel context.CancelFunc
	p2pIdToOracleId  map[cciptypes.Bytes32]commontypes.OracleID
}

func NewHomeChainConfigPoller(
	homeChainReader types.ContractReader,
	lggr logger.Logger,
	p2pIdToOracleId map[cciptypes.Bytes32]commontypes.OracleID,
) *HomeChainConfigPoller {
	return &HomeChainConfigPoller{
		homeChainReader: homeChainReader,
		lggr:            lggr,
		p2pIdToOracleId: p2pIdToOracleId,
	}
}

func (r *HomeChainConfigPoller) StartPolling(ctx context.Context, interval time.Duration) {

	r.mutex.Lock()
	if r.backgroundCancel != nil {
		r.lggr.Errorw("Polling already started")
		return
	}
	bgCtx, cancelFunc := context.WithCancel(ctx)
	r.backgroundCancel = cancelFunc
	r.mutex.Unlock()
	r.lggr.Infow("Start Polling HomeChainConfig")
	go func() {
		ticker := time.NewTicker(interval)
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
}

func (r *HomeChainConfigPoller) GetConfig() cciptypes.HomeChainConfig {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.homeChainConfig
}

func (r *HomeChainConfigPoller) GetSupportedChains(oracleID commontypes.OracleID) mapset.Set[cciptypes.ChainSelector] {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.homeChainConfig.NodeSupportedChains[oracleID].Supported
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

func (r *HomeChainConfigPoller) Close(ctx context.Context) error {
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
