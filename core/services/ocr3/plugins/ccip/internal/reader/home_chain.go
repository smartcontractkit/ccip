package reader

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"
)

type HomeChainConfigPoller struct {
	homeChainReader types.ContractReader
	// gets updated by the polling loop
	chainConfigs map[cciptypes.ChainSelector]cciptypes.ChainConfig
	// mapping between each node's peerID and the chains it supports. derived from chainConfigs
	nodeSupportedChains map[libocrtypes.PeerID]mapset.Set[cciptypes.ChainSelector]
	// set of chains that are known to CCIP, derived from chainConfigs
	knownSourceChains mapset.Set[cciptypes.ChainSelector]
	// map of chain to FChain value, derived from chainConfigs
	fChain           map[cciptypes.ChainSelector]int
	lggr             logger.Logger
	mutex            *sync.RWMutex
	backgroundCancel context.CancelFunc
	backgroundCtx    context.Context
	// How frequent will the poller fetch the chain configs
	pollingInterval time.Duration
}

func NewHomeChainConfigPoller(
	homeChainReader types.ContractReader,
	lggr logger.Logger,
	pollingInterval time.Duration,
) *HomeChainConfigPoller {
	ctx, cancel := context.WithCancel(context.Background())
	return &HomeChainConfigPoller{
		homeChainReader:     homeChainReader,
		lggr:                lggr,
		chainConfigs:        map[cciptypes.ChainSelector]cciptypes.ChainConfig{},
		mutex:               &sync.RWMutex{},
		nodeSupportedChains: map[libocrtypes.PeerID]mapset.Set[cciptypes.ChainSelector]{},
		knownSourceChains:   mapset.NewSet[cciptypes.ChainSelector](),
		fChain:              map[cciptypes.ChainSelector]int{},
		backgroundCancel:    cancel,
		backgroundCtx:       ctx,
		pollingInterval:     pollingInterval,
	}
}

func (r *HomeChainConfigPoller) Start(ctx context.Context) error {
	r.mutex.Lock()
	if r.backgroundCancel != nil {
		r.lggr.Errorw("Polling already started")
		// We don't want to return an actual error here as it's already working as expected
		r.mutex.Unlock()
		return fmt.Errorf("polling already started")
	}
	if r.backgroundCtx == nil {
		return fmt.Errorf("backgroundCtx not set")
	}
	//bgCtx, cancelFunc := context.WithCancel(context.Background())
	//r.backgroundCancel = cancelFunc
	r.mutex.Unlock()

	err := r.fetchAndSetConfigs(r.backgroundCtx)
	if err != nil {
		r.lggr.Errorw("Initial fetch of on-chain configs failed", "err", err)
		return fmt.Errorf("initial fetch of on-chain configs failed")
	}

	r.lggr.Infow("Start Polling ChainConfig")
	go r.poll(r.backgroundCtx)
	//go r.poll(bgCtx)
	return nil
}

func (r *HomeChainConfigPoller) poll(ctx context.Context) {
	ticker := time.NewTicker(r.pollingInterval * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			r.lggr.Infow("Polling stopped")
			return
		case <-ticker.C:
			err := r.fetchAndSetConfigs(ctx)
			if err != nil {
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
		r.lggr.Errorw("no on chain configs found")
		return fmt.Errorf("no on chain configs found")
	}
	homeChainConfigs, err := r.convertOnChainConfigToHomeChainConfig(chainConfigInfos)
	if err != nil {
		r.lggr.Errorw("error converting OnChainConfigs to ChainConfig", "err", err)
		return err
	}
	r.lggr.Infow("Setting ChainConfig")
	r.mutex.Lock()
	r.chainConfigs = homeChainConfigs
	r.nodeSupportedChains = createNodesSupportedChains(homeChainConfigs)
	r.knownSourceChains = createKnownChains(homeChainConfigs)
	r.fChain = createFChain(homeChainConfigs)
	r.mutex.Unlock()
	return nil
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
	if len(r.chainConfigs) == 0 {
		return nil, errors.New("no chain configs found")
	}
	return r.chainConfigs, nil

}

func (r *HomeChainConfigPoller) GetSupportedChains(id libocrtypes.PeerID) (mapset.Set[cciptypes.ChainSelector], error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if _, ok := r.nodeSupportedChains[id]; !ok {
		return nil, fmt.Errorf("node %v not found", id)
	}
	return r.nodeSupportedChains[id], nil
}

func (r *HomeChainConfigPoller) GetKnownChains() (mapset.Set[cciptypes.ChainSelector], error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	knownSourceChains := mapset.NewSet[cciptypes.ChainSelector]()
	for chain, _ := range r.chainConfigs {
		knownSourceChains.Add(chain)
	}
	if knownSourceChains.Cardinality() == 0 {
		return nil, fmt.Errorf("no known chain configs")
	}

	return knownSourceChains, nil
}

func (r *HomeChainConfigPoller) GetFChain() (map[cciptypes.ChainSelector]int, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if len(r.fChain) == 0 {
		return nil, fmt.Errorf("no FChain values found")
	}
	return r.fChain, nil
}

func (r *HomeChainConfigPoller) Close() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if r.backgroundCancel == nil {
		return errors.New("cancel function not set")
	}
	r.backgroundCancel()
	return nil
}

func (r *HomeChainConfigPoller) Ready() error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if len(r.chainConfigs) == 0 {
		return errors.New("no chain configs found")
	}
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
	for chain, _ := range chainConfigs {
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
	r.lggr.Infow("Fetching OnChainConfig")
	var chainConfigInfo []ChainConfigInfo
	err := r.homeChainReader.GetLatestValue(ctx, "CCIPCapabilityConfiguration", "getAllChainConfigs", nil, &chainConfigInfo)
	if err != nil {
		return nil, err
	}
	return chainConfigInfo, err
}

func (r *HomeChainConfigPoller) convertOnChainConfigToHomeChainConfig(capabilityConfigs []ChainConfigInfo) (map[cciptypes.ChainSelector]cciptypes.ChainConfig, error) {
	chainConfigs := make(map[cciptypes.ChainSelector]cciptypes.ChainConfig)
	//iterate over configs
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

type OnChainConfig struct {
	Readers []libocrtypes.PeerID `json:"readers"`
	FChain  uint8                `json:"fChain"`
	Config  []byte               `json:"config"`
}
type ChainConfigInfo struct {
	// Calling function https://github.com/smartcontractkit/ccip/blob/330c5e98f624cfb10108c92fe1e00ced6d345a99/contracts/src/v0.8/ccip/capability/CCIPCapabilityConfiguration.sol#L140
	ChainSelector cciptypes.ChainSelector `json:"chainSelector"`
	ChainConfig   OnChainConfig           `json:"chainConfig"`
}

var _ cciptypes.HomeChainPoller = (*HomeChainConfigPoller)(nil)
