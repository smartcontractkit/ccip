package laneconfig

import (
	_ "embed"
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/multierr"
)

var (
	//go:embed contracts.json
	ExistingContracts []byte
	laneMu            = &sync.Mutex{}
)

type CommonContracts struct {
	IsNativeFeeToken bool     `json:"is_native_fee_token,omitempty"`
	IsMockARM        bool     `json:"is_mock_arm,omitempty"`
	FeeToken         string   `json:"fee_token"`
	BridgeTokens     []string `json:"bridge_tokens"`
	BridgeTokenPools []string `json:"bridge_tokens_pools"`
	ARM              string   `json:"arm"`
	Router           string   `json:"router"`
	PriceRegistry    string   `json:"price_registry"`
	WrappedNative    string   `json:"wrapped_native"`
}

type SourceContracts struct {
	OnRamp     string `json:"on_ramp"`
	DepolyedAt uint64 `json:"deployed_at"`
}

type DestContracts struct {
	OffRamp      string `json:"off_ramp"`
	CommitStore  string `json:"commit_store"`
	ReceiverDapp string `json:"receiver_dapp"`
}

type LaneConfig struct {
	CommonContracts
	SrcContractsMu  *sync.Mutex                `json:"-"`
	SrcContracts    map[string]SourceContracts `json:"src_contracts"` // key destination chain id
	DestContractsMu *sync.Mutex                `json:"-"`
	DestContracts   map[string]DestContracts   `json:"dest_contracts"` // key source chain id
}

func (l *LaneConfig) Validate() error {
	var laneConfigError error

	if l.ARM == "" || !common.IsHexAddress(l.ARM) {
		laneConfigError = multierr.Append(laneConfigError, errors.New("must set proper address for arm"))
	}

	if l.FeeToken != "" && !common.IsHexAddress(l.FeeToken) {
		laneConfigError = multierr.Append(laneConfigError, errors.New("must set proper address for fee_token"))
	}

	for _, token := range l.BridgeTokens {
		if token != "" && !common.IsHexAddress(token) {
			laneConfigError = multierr.Append(laneConfigError, errors.New("must set proper address for bridge_tokens"))
		}
	}

	for _, pool := range l.BridgeTokenPools {
		if pool != "" && !common.IsHexAddress(pool) {
			laneConfigError = multierr.Append(laneConfigError, errors.New("must set proper address for bridge_tokens_pools"))
		}
	}
	if l.Router == "" || !common.IsHexAddress(l.Router) {
		laneConfigError = multierr.Append(laneConfigError, errors.New("must set proper address for router"))
	}
	if l.PriceRegistry == "" || !common.IsHexAddress(l.PriceRegistry) {
		laneConfigError = multierr.Append(laneConfigError, errors.New("must set proper address for price_registry"))
	}
	if l.WrappedNative == "" || !common.IsHexAddress(l.WrappedNative) {
		laneConfigError = multierr.Append(laneConfigError, errors.New("must set proper address for wrapped_native"))
	}
	return laneConfigError
}

type Lanes struct {
	LaneConfigs map[string]*LaneConfig `json:"lane_configs"`
}

func (l *Lanes) ReadLaneConfig(networkA string) *LaneConfig {
	laneMu.Lock()
	defer laneMu.Unlock()
	cfg, ok := l.LaneConfigs[networkA]
	if !ok {
		l.LaneConfigs[networkA] = &LaneConfig{
			SrcContracts:    make(map[string]SourceContracts),
			DestContracts:   make(map[string]DestContracts),
			SrcContractsMu:  &sync.Mutex{},
			DestContractsMu: &sync.Mutex{},
		}
		return l.LaneConfigs[networkA]
	}
	if cfg.SrcContractsMu == nil {
		l.LaneConfigs[networkA].SrcContractsMu = &sync.Mutex{}
	}
	if cfg.DestContractsMu == nil {
		l.LaneConfigs[networkA].DestContractsMu = &sync.Mutex{}
	}
	return l.LaneConfigs[networkA]
}

// CopyCommonContracts copies network config for common contracts from fromNetwork to toNetwork
// if the toNetwork already exists, it does nothing
// If reuse is set to false, it only retains the token contracts
func (l *Lanes) CopyCommonContracts(fromNetwork, toNetwork string, reuse, isTokenTransfer bool) {
	laneMu.Lock()
	defer laneMu.Unlock()
	// if the toNetwork already exists, return
	if _, ok := l.LaneConfigs[toNetwork]; ok {
		return
	}
	existing, ok := l.LaneConfigs[fromNetwork]
	if !ok {
		l.LaneConfigs[toNetwork] = &LaneConfig{
			SrcContracts:    make(map[string]SourceContracts),
			DestContracts:   make(map[string]DestContracts),
			SrcContractsMu:  &sync.Mutex{},
			DestContractsMu: &sync.Mutex{},
		}
		return
	}
	cfg := &LaneConfig{
		SrcContracts:    make(map[string]SourceContracts),
		SrcContractsMu:  &sync.Mutex{},
		DestContractsMu: &sync.Mutex{},
		DestContracts:   make(map[string]DestContracts),
		CommonContracts: CommonContracts{
			FeeToken:         existing.FeeToken,
			WrappedNative:    existing.WrappedNative,
			IsNativeFeeToken: existing.IsNativeFeeToken,
		},
	}
	// if reuse is set to true, it copies all the common contracts except the router
	if reuse {
		cfg.CommonContracts.PriceRegistry = existing.PriceRegistry
		cfg.CommonContracts.ARM = existing.ARM
		cfg.CommonContracts.IsNativeFeeToken = existing.IsNativeFeeToken
		cfg.CommonContracts.IsMockARM = existing.IsMockARM
	}
	// if it is a token transfer, it copies the bridge token contracts
	if isTokenTransfer {
		cfg.CommonContracts.BridgeTokens = existing.BridgeTokens
		if reuse {
			cfg.CommonContracts.BridgeTokenPools = existing.BridgeTokenPools
		}
	}
	l.LaneConfigs[toNetwork] = cfg
}

func (l *Lanes) WriteLaneConfig(networkA string, cfg *LaneConfig) error {
	laneMu.Lock()
	defer laneMu.Unlock()
	if l.LaneConfigs == nil {
		l.LaneConfigs = make(map[string]*LaneConfig)
	}
	err := cfg.Validate()
	if err != nil {
		return err
	}
	l.LaneConfigs[networkA] = cfg
	return nil
}

func ReadLanesFromExistingDeployment() (*Lanes, error) {
	var lanes Lanes
	if err := json.Unmarshal(ExistingContracts, &lanes); err != nil {
		return nil, err
	}
	return &lanes, nil
}

func CreateDeploymentJSON(path string) (*Lanes, error) {
	existingLanes := Lanes{
		LaneConfigs: make(map[string]*LaneConfig),
	}
	err := WriteLanesToJSON(path, &existingLanes)
	return &existingLanes, err
}

func WriteLanesToJSON(path string, lanes *Lanes) error {
	b, err := json.MarshalIndent(lanes, "", "  ")
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(b)
	return err
}
