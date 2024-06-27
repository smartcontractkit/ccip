package reader

import (
	"context"
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

type HomeChain interface {
	GetChainConfig(chainSelector cciptypes.ChainSelector) (ChainConfig, error)
	GetAllChainConfigs() (map[cciptypes.ChainSelector]ChainConfig, error)
	// GetSupportedChainsForPeer Gets all chain selectors that the peerID can read/write from/to
	GetSupportedChainsForPeer(id libocrtypes.PeerID) (mapset.Set[cciptypes.ChainSelector], error)
	// GetKnownCCIPChains Gets all chain selectors that are known to CCIP
	GetKnownCCIPChains() (mapset.Set[cciptypes.ChainSelector], error)
	// GetFChain Gets the FChain value for each chain
	GetFChain() (map[cciptypes.ChainSelector]int, error)
	// GetOCRConfigs Gets the OCR3Configs for a given donID and pluginType
	GetOCRConfigs(ctx context.Context, donID uint32, pluginType uint8) ([]OCR3ConfigWithMeta, error)
	services.Service
}

type state struct {
	// gets updated by the polling loop
	chainConfigs map[cciptypes.ChainSelector]ChainConfig
	// mapping between each node's peerID and the chains it supports. derived from chainConfigs
	nodeSupportedChains map[libocrtypes.PeerID]mapset.Set[cciptypes.ChainSelector]
	// set of chains that are known to CCIP, derived from chainConfigs
	knownSourceChains mapset.Set[cciptypes.ChainSelector]
	// map of chain to FChain value, derived from chainConfigs
	fChain map[cciptypes.ChainSelector]int
}

type homeChainPoller struct {
	stopCh          services.StopChan
	sync            services.StateMachine
	homeChainReader types.ContractReader
	lggr            logger.Logger
	mutex           *sync.RWMutex
	state           state
	failedPolls     uint
	// How frequently the poller fetches the chain configs
	pollingDuration time.Duration
}

const MaxFailedPolls = 10

func NewHomeChainConfigPoller(
	homeChainReader types.ContractReader,
	lggr logger.Logger,
	pollingInterval time.Duration,
) HomeChain {
	return &homeChainPoller{
		stopCh:          make(chan struct{}),
		homeChainReader: homeChainReader,
		state:           state{},
		mutex:           &sync.RWMutex{},
		failedPolls:     0,
		lggr:            lggr,
		pollingDuration: pollingInterval,
	}
}

func (r *homeChainPoller) Start(ctx context.Context) error {
	err := r.fetchAndSetConfigs(ctx)
	if err != nil {
		// Just log, don't return error as we want to keep polling
		r.lggr.Errorw("Initial fetch of on-chain configs failed", "err", err)
	}
	r.lggr.Infow("Start Polling ChainConfig")
	return r.sync.StartOnce(r.Name(), func() error {
		go r.poll()
		return nil
	})
}

func (r *homeChainPoller) poll() {
	ctx, cancel := r.stopCh.NewCtx()
	defer cancel()
	ticker := time.NewTicker(r.pollingDuration)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			r.mutex.Lock()
			r.failedPolls = 0
			r.mutex.Unlock()
			return
		case <-ticker.C:
			if err := r.fetchAndSetConfigs(ctx); err != nil {
				r.mutex.Lock()
				r.failedPolls++
				r.mutex.Unlock()
				r.lggr.Errorw("fetching and setting configs failed", "failedPolls", r.failedPolls, "err", err)
			}
		}
	}
}

func (r *homeChainPoller) fetchAndSetConfigs(ctx context.Context) error {
	var chainConfigInfos []ChainConfigInfo
	err := r.homeChainReader.GetLatestValue(
		ctx, "CCIPCapabilityConfiguration", "getAllChainConfigs", nil, &chainConfigInfos,
	)
	if err != nil {
		return err
	}
	if len(chainConfigInfos) == 0 {
		// That's a legitimate case if there are no chain configs on chain yet
		r.lggr.Warnw("no on chain configs found")
		return nil
	}
	homeChainConfigs, err := convertOnChainConfigToHomeChainConfig(chainConfigInfos)
	if err != nil {
		r.lggr.Errorw("error converting OnChainConfigs to ChainConfig", "err", err)
		return err
	}
	r.setState(homeChainConfigs)
	return nil
}

func (r *homeChainPoller) setState(chainConfigs map[cciptypes.ChainSelector]ChainConfig) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	s := &r.state
	s.chainConfigs = chainConfigs
	s.nodeSupportedChains = createNodesSupportedChains(chainConfigs)
	s.knownSourceChains = createKnownChains(chainConfigs)
	s.fChain = createFChain(chainConfigs)
}

func (r *homeChainPoller) GetChainConfig(chainSelector cciptypes.ChainSelector) (ChainConfig, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	s := r.state
	if chainConfig, ok := s.chainConfigs[chainSelector]; ok {
		return chainConfig, nil
	}
	return ChainConfig{}, fmt.Errorf("chain config not found for chain %v", chainSelector)
}

func (r *homeChainPoller) GetAllChainConfigs() (map[cciptypes.ChainSelector]ChainConfig, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.state.chainConfigs, nil
}

func (r *homeChainPoller) GetSupportedChainsForPeer(
	id libocrtypes.PeerID,
) (mapset.Set[cciptypes.ChainSelector], error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	s := r.state
	if _, ok := s.nodeSupportedChains[id]; !ok {
		// empty set to denote no chains supported
		return mapset.NewSet[cciptypes.ChainSelector](), nil
	}
	return s.nodeSupportedChains[id], nil
}

func (r *homeChainPoller) GetKnownCCIPChains() (mapset.Set[cciptypes.ChainSelector], error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	knownSourceChains := mapset.NewSet[cciptypes.ChainSelector]()
	for chain := range r.state.chainConfigs {
		knownSourceChains.Add(chain)
	}

	return knownSourceChains, nil
}

func (r *homeChainPoller) GetFChain() (map[cciptypes.ChainSelector]int, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.state.fChain, nil
}

func (r *homeChainPoller) GetOCRConfigs(
	ctx context.Context, donID uint32, pluginType uint8,
) ([]OCR3ConfigWithMeta, error) {
	var ocrConfigs []OCR3ConfigWithMeta
	err := r.homeChainReader.GetLatestValue(ctx, "CCIPCapabilityConfiguration", "getOCRConfig", map[string]any{
		"donId":      donID,
		"pluginType": pluginType,
	}, &ocrConfigs)
	if err != nil {
		return nil, fmt.Errorf("error fetching OCR configs: %w", err)
	}

	return ocrConfigs, nil
}

func (r *homeChainPoller) Close() error {
	err := r.sync.StopOnce(r.Name(), func() error {
		close(r.stopCh)
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to stop %s: %w", r.Name(), err)
	}
	// give it twice the polling duration to ensure the poller is caught up and stopped
	ticker := time.NewTicker(r.pollingDuration * 4)
	defer ticker.Stop()
	for {
		// Make sure it's closed gracefully (Ready returns an error once it's not ready)
		err := r.Ready()
		if err != nil {
			return nil
		}
		select {
		case <-ticker.C:
			return fmt.Errorf("HomeChainReader did not close gracefully")
		default:
		}
	}
}

func (r *homeChainPoller) Ready() error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	return r.sync.Ready()
}

func (r *homeChainPoller) HealthReport() map[string]error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if r.failedPolls >= MaxFailedPolls {
		r.sync.SvcErrBuffer.Append(fmt.Errorf("polling failed %d times in a row", MaxFailedPolls))
	}
	return map[string]error{r.Name(): r.sync.Healthy()}
}

func (r *homeChainPoller) Name() string {
	return "homeChainPoller"
}

func createFChain(chainConfigs map[cciptypes.ChainSelector]ChainConfig) map[cciptypes.ChainSelector]int {
	fChain := map[cciptypes.ChainSelector]int{}
	for chain, config := range chainConfigs {
		fChain[chain] = config.FChain
	}
	return fChain
}

func createKnownChains(chainConfigs map[cciptypes.ChainSelector]ChainConfig) mapset.Set[cciptypes.ChainSelector] {
	knownChains := mapset.NewSet[cciptypes.ChainSelector]()
	for chain := range chainConfigs {
		knownChains.Add(chain)
	}
	return knownChains
}

func createNodesSupportedChains(
	chainConfigs map[cciptypes.ChainSelector]ChainConfig,
) map[libocrtypes.PeerID]mapset.Set[cciptypes.ChainSelector] {
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

func convertOnChainConfigToHomeChainConfig(
	capabilityConfigs []ChainConfigInfo,
) (map[cciptypes.ChainSelector]ChainConfig, error) {
	chainConfigs := make(map[cciptypes.ChainSelector]ChainConfig)
	for _, capabilityConfig := range capabilityConfigs {
		chainSelector := capabilityConfig.ChainSelector
		config := capabilityConfig.ChainConfig

		chainConfigs[chainSelector] = ChainConfig{
			FChain:         int(config.FChain),
			SupportedNodes: mapset.NewSet(config.Readers...),
		}
	}
	return chainConfigs, nil
}

// HomeChainConfigMapper This is a 1-1 mapping between the config that we get from the contract to make
// se/deserializing easier
type HomeChainConfigMapper struct {
	Readers []libocrtypes.PeerID `json:"readers"`
	FChain  uint8                `json:"fChain"`
	Config  []byte               `json:"config"`
}

// ChainConfigInfo This is a 1-1 mapping between the config that we get from the contract to make
// se/deserializing easier
type ChainConfigInfo struct {
	// nolint:lll // don't split up the long url
	// Calling function https://github.com/smartcontractkit/ccip/blob/330c5e98f624cfb10108c92fe1e00ced6d345a99/contracts/src/v0.8/ccip/capability/CCIPCapabilityConfiguration.sol#L140
	ChainSelector cciptypes.ChainSelector `json:"chainSelector"`
	ChainConfig   HomeChainConfigMapper   `json:"chainConfig"`
}

// ChainConfig will live on the home chain and will be used to update chain configuration like F value and supported
// nodes dynamically.
type ChainConfig struct {
	// FChain defines the FChain value for the chain. FChain is used while forming consensus based on the observations.
	FChain int `json:"fChain"`
	// SupportedNodes is a map of PeerIDs to SupportedChains.
	SupportedNodes mapset.Set[libocrtypes.PeerID] `json:"supportedNodes"`
	// Config is the chain specific configuration.
	Config []byte `json:"config"`
}

// OCR3Config mirrors CCIPCapabilityConfiguration.sol's OCR3Config struct
type OCR3Config struct {
	PluginType            uint8                   `json:"pluginType"`
	ChainSelector         cciptypes.ChainSelector `json:"chainSelector"`
	F                     uint8                   `json:"F"`
	OffchainConfigVersion uint64                  `json:"offchainConfigVersion"`
	OfframpAddress        []byte                  `json:"offrampAddress"`
	BootstrapP2PIds       [][32]byte              `json:"bootstrapP2PIds"`
	P2PIds                [][32]byte              `json:"p2pIds"`
	Signers               [][]byte                `json:"signers"`
	Transmitters          [][]byte                `json:"transmitters"`
	OffchainConfig        []byte                  `json:"offchainConfig"`
}

// OCR3ConfigWithmeta mirrors CCIPCapabilityConfiguration.sol's OCR3ConfigWithMeta struct
type OCR3ConfigWithMeta struct {
	Config       OCR3Config `json:"config"`
	ConfigCount  uint64     `json:"configCount"`
	ConfigDigest [32]byte   `json:"configDigest"`
}

var _ HomeChain = (*homeChainPoller)(nil)
