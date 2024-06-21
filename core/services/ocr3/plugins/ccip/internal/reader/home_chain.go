package reader

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"
)

type HomeChainPoller interface {
	GetChainConfig(chainSelector cciptypes.ChainSelector) (cciptypes.ChainConfig, error)
	GetAllChainConfigs() (map[cciptypes.ChainSelector]cciptypes.ChainConfig, error)
	// GetSupportedChainsForPeer Gets all chain selectors that the peerID can read/write from/to
	GetSupportedChainsForPeer(id libocrtypes.PeerID) (mapset.Set[cciptypes.ChainSelector], error)
	// GetKnownCCIPChains Gets all chain selectors that are known to CCIP
	GetKnownCCIPChains() (mapset.Set[cciptypes.ChainSelector], error)
	// GetFChain Gets the FChain value for each chain
	GetFChain() (map[cciptypes.ChainSelector]int, error)
	services.Service
}

type HomeChainConfigPoller struct {
	stopCh services.StopChan
	services.StateMachine

	homeChainReader types.ContractReader
	// gets updated by the polling loop
	chainConfigs map[cciptypes.ChainSelector]cciptypes.ChainConfig
	// mapping between each node's peerID and the chains it supports. derived from chainConfigs
	nodeSupportedChains map[libocrtypes.PeerID]mapset.Set[cciptypes.ChainSelector]
	// set of chains that are known to CCIP, derived from chainConfigs
	knownSourceChains mapset.Set[cciptypes.ChainSelector]
	// map of chain to FChain value, derived from chainConfigs
	fChain map[cciptypes.ChainSelector]int
	lggr   logger.Logger
	mutex  *sync.RWMutex
	// How frequent will the poller fetch the chain configs
	pollingInterval time.Duration
}

func NewHomeChainConfigPoller(
	homeChainReader types.ContractReader,
	lggr logger.Logger,
	pollingInterval time.Duration,
) *HomeChainConfigPoller {
	return &HomeChainConfigPoller{
		stopCh:              make(chan struct{}),
		homeChainReader:     homeChainReader,
		lggr:                lggr,
		chainConfigs:        map[cciptypes.ChainSelector]cciptypes.ChainConfig{},
		mutex:               &sync.RWMutex{},
		nodeSupportedChains: map[libocrtypes.PeerID]mapset.Set[cciptypes.ChainSelector]{},
		knownSourceChains:   mapset.NewSet[cciptypes.ChainSelector](),
		fChain:              map[cciptypes.ChainSelector]int{},
		pollingInterval:     pollingInterval,
	}
}

func (r *HomeChainConfigPoller) Start(ctx context.Context) error {
	err := r.fetchAndSetConfigs(ctx)
	if err != nil {
		// Just log, don't return error as we want to keep polling
		r.lggr.Errorw("Initial fetch of on-chain configs failed", "err", err)
	}
	r.lggr.Infow("Start Polling ChainConfig")
	return r.StartOnce(r.Name(), func() error {
		go r.poll()
		return nil
	})
}

func (r *HomeChainConfigPoller) poll() {
	ctx, cancel := r.stopCh.NewCtx()
	defer cancel()
	ticker := time.NewTicker(r.pollingInterval * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := r.fetchAndSetConfigs(ctx); err != nil {
				r.lggr.Errorw("Fetching and setting configs failed", "err", err)
			}
		}
	}
}

func (r *HomeChainConfigPoller) fetchAndSetConfigs(ctx context.Context) error {
	chainConfigInfos, err := r.fetchOnChainConfig(ctx)
	if err != nil {
		r.lggr.Errorw("Fetching on-chain configs failed", "err", err)
		return err
	}
	if len(chainConfigInfos) == 0 {
		// That's a legitimate case if there are no chain configs on chain yet
		r.lggr.Warnw("no on chain configs found")
	}
	homeChainConfigs, err := r.convertOnChainConfigToHomeChainConfig(chainConfigInfos)
	if err != nil {
		r.lggr.Errorw("error converting OnChainConfigs to ChainConfig", "err", err)
		return err
	}
	r.lggr.Infow("Setting ChainConfig")
	r.setChainConfigs(homeChainConfigs)
	return nil
}

func (r *HomeChainConfigPoller) setChainConfigs(chainConfigs map[cciptypes.ChainSelector]cciptypes.ChainConfig) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.chainConfigs = chainConfigs
	r.nodeSupportedChains = createNodesSupportedChains(chainConfigs)
	r.knownSourceChains = createKnownChains(chainConfigs)
	r.fChain = createFChain(chainConfigs)
}

func (r *HomeChainConfigPoller) GetChainConfig(chainSelector cciptypes.ChainSelector) (cciptypes.ChainConfig, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if _, ok := r.chainConfigs[chainSelector]; !ok {
		return cciptypes.ChainConfig{}, fmt.Errorf("chain config not found for chain %v", chainSelector)
	}
	return r.chainConfigs[chainSelector], nil
}

func (r *HomeChainConfigPoller) GetAllChainConfigs() (map[cciptypes.ChainSelector]cciptypes.ChainConfig, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.chainConfigs, nil

}

func (r *HomeChainConfigPoller) GetSupportedChainsForPeer(id libocrtypes.PeerID) (mapset.Set[cciptypes.ChainSelector], error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if _, ok := r.nodeSupportedChains[id]; !ok {
		// empty set to denote no chains supported
		return mapset.NewSet[cciptypes.ChainSelector](), nil
	}
	return r.nodeSupportedChains[id], nil
}

func (r *HomeChainConfigPoller) GetKnownCCIPChains() (mapset.Set[cciptypes.ChainSelector], error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	knownSourceChains := mapset.NewSet[cciptypes.ChainSelector]()
	for chain := range r.chainConfigs {
		knownSourceChains.Add(chain)
	}

	return knownSourceChains, nil
}

func (r *HomeChainConfigPoller) GetFChain() (map[cciptypes.ChainSelector]int, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.fChain, nil
}

func (r *HomeChainConfigPoller) Close() error {
	return r.StopOnce(r.Name(), func() error {
		close(r.stopCh)
		return nil
	})
}

func (r *HomeChainConfigPoller) Ready() error {
	return nil
}

func (r *HomeChainConfigPoller) HealthReport() map[string]error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if len(r.chainConfigs) == 0 {
		return map[string]error{"ChainConfig": errors.New("no chain configs found")}
	}
	return nil
}

func (r *HomeChainConfigPoller) Name() string {
	return "HomeChainConfigPoller"
}

func createFChain(chainConfigs map[cciptypes.ChainSelector]cciptypes.ChainConfig) map[cciptypes.ChainSelector]int {
	fChain := map[cciptypes.ChainSelector]int{}
	for chain, config := range chainConfigs {
		fChain[chain] = config.FChain
	}
	return fChain
}

func createKnownChains(chainConfigs map[cciptypes.ChainSelector]cciptypes.ChainConfig) mapset.Set[cciptypes.ChainSelector] {
	knownChains := mapset.NewSet[cciptypes.ChainSelector]()
	for chain := range chainConfigs {
		knownChains.Add(chain)
	}
	return knownChains
}

func createNodesSupportedChains(chainConfigs map[cciptypes.ChainSelector]cciptypes.ChainConfig) map[libocrtypes.PeerID]mapset.Set[cciptypes.ChainSelector] {
	nodeSupportedChains := map[libocrtypes.PeerID]mapset.Set[cciptypes.ChainSelector]{}
	for chainSelector, config := range chainConfigs {
		for _, p2pID := range config.SupportedNodes.ToSlice() {
			if _, ok := nodeSupportedChains[p2pID]; !ok {
				nodeSupportedChains[p2pID] = mapset.NewSet[cciptypes.ChainSelector]()
			}
			//add chain to SupportedChains
			nodeSupportedChains[p2pID].Add(chainSelector)
		}
	}
	return nodeSupportedChains
}

func (r *HomeChainConfigPoller) fetchOnChainConfig(ctx context.Context) ([]ChainConfigInfo, error) {
	r.lggr.Infow("Fetching HomeChainConfigMapper")
	var chainConfigInfo []ChainConfigInfo
	err := r.homeChainReader.GetLatestValue(ctx, "CCIPCapabilityConfiguration", "getAllChainConfigs", nil, &chainConfigInfo)
	if err != nil {
		return nil, err
	}
	return chainConfigInfo, err
}

func (r *HomeChainConfigPoller) convertOnChainConfigToHomeChainConfig(capabilityConfigs []ChainConfigInfo) (map[cciptypes.ChainSelector]cciptypes.ChainConfig, error) {
	chainConfigs := make(map[cciptypes.ChainSelector]cciptypes.ChainConfig)
	for _, capabilityConfig := range capabilityConfigs {
		chainSelector := capabilityConfig.ChainSelector
		config := capabilityConfig.ChainConfig

		chainConfigs[chainSelector] = cciptypes.ChainConfig{
			FChain:         int(config.FChain),
			SupportedNodes: mapset.NewSet(config.Readers...),
		}
	}
	return chainConfigs, nil
}

// HomeChainConfigMapper This is a 1-1 mapping between the config that we get from the contract to make se/deserializing easier
type HomeChainConfigMapper struct {
	Readers []libocrtypes.PeerID `json:"readers"`
	FChain  uint8                `json:"fChain"`
	Config  []byte               `json:"config"`
}

// ChainConfigInfo This is a 1-1 mapping between the config that we get from the contract to make se/deserializing easier
type ChainConfigInfo struct {
	// Calling function https://github.com/smartcontractkit/ccip/blob/330c5e98f624cfb10108c92fe1e00ced6d345a99/contracts/src/v0.8/ccip/capability/CCIPCapabilityConfiguration.sol#L140
	ChainSelector cciptypes.ChainSelector `json:"chainSelector"`
	ChainConfig   HomeChainConfigMapper   `json:"chainConfig"`
}

var _ HomeChainPoller = (*HomeChainConfigPoller)(nil)
