package laneconfig

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/multierr"
)

//go:embed contracts.json
var ExistingContracts []byte

type CommonContracts struct {
	IsNativeFeeToken bool     `json:"is_native_fee_token,omitempty"`
	FeeToken         string   `json:"fee_token"`
	FeeTokenPool     string   `json:"fee_token_pool"`
	BridgeTokens     []string `json:"bridge_tokens"`
	BridgeTokenPools []string `json:"bridge_tokens_pools"`
	AFN              string   `json:"afn"`
	Router           string   `json:"router"`
	PriceRegistry    string   `json:"price_registry"`
}

type LaneConfig struct {
	CommonContracts
	SrcContracts  map[uint64]SourceContracts `json:"src_contracts"`  // key destination chain id
	DestContracts map[uint64]DestContracts   `json:"dest_contracts"` // key source chain id
}

type SourceContracts struct {
	OnRamp string `json:"on_ramp"`
}

type DestContracts struct {
	OffRamp      string `json:"off_ramp"`
	CommitStore  string `json:"commit_store"`
	ReceiverDapp string `json:"receiver_dapp"`
}

type Lane struct {
	NetworkA   string     `json:"network_name"`
	LaneConfig LaneConfig `json:"lane_config"`
}

func (l Lane) Validate() error {
	var laneConfigError error

	if l.NetworkA == "" {
		laneConfigError = multierr.Append(laneConfigError, errors.New("must set network_name"))
	}
	if l.LaneConfig.AFN == "" || !common.IsHexAddress(l.LaneConfig.AFN) {
		laneConfigError = multierr.Append(laneConfigError, errors.New("must set proper address for afn"))
	}
	if l.LaneConfig.FeeTokenPool == "" || !common.IsHexAddress(l.LaneConfig.FeeTokenPool) {
		laneConfigError = multierr.Append(laneConfigError, errors.New("must set proper address for fee_token_pool"))
	}
	if l.LaneConfig.FeeToken == "" || !common.IsHexAddress(l.LaneConfig.FeeToken) {
		laneConfigError = multierr.Append(laneConfigError, errors.New("must set proper address for fee_token"))
	}
	if len(l.LaneConfig.BridgeTokens) < 1 {
		laneConfigError = multierr.Append(laneConfigError, errors.New("must set at least 1 bridge_tokens"))
	}
	for _, token := range l.LaneConfig.BridgeTokens {
		if token == "" || !common.IsHexAddress(token) {
			laneConfigError = multierr.Append(laneConfigError, errors.New("must set proper address for bridge_tokens"))
		}
	}
	if len(l.LaneConfig.BridgeTokenPools) < 1 {
		laneConfigError = multierr.Append(laneConfigError, errors.New("must set at least 1 bridge_tokens_pools"))
	}
	for _, pool := range l.LaneConfig.BridgeTokenPools {
		if pool == "" || !common.IsHexAddress(pool) {
			laneConfigError = multierr.Append(laneConfigError, errors.New("must set proper address for bridge_tokens_pools"))
		}
	}
	if l.LaneConfig.Router == "" || !common.IsHexAddress(l.LaneConfig.Router) {
		laneConfigError = multierr.Append(laneConfigError, errors.New("must set proper address for router"))
	}
	if l.LaneConfig.PriceRegistry == "" || !common.IsHexAddress(l.LaneConfig.PriceRegistry) {
		laneConfigError = multierr.Append(laneConfigError, errors.New("must set proper address for price_registry"))
	}
	return laneConfigError
}

// copyConfig updates l1 with l2 and returns l1
func (l1 *Lane) copyConfig(l2 Lane) {
	if l1.LaneConfig.SrcContracts == nil {
		l1.LaneConfig.SrcContracts = l2.LaneConfig.SrcContracts
	}
	for chain, cfg := range l2.LaneConfig.SrcContracts {
		l1.LaneConfig.SrcContracts[chain] = cfg
	}
	if l1.LaneConfig.DestContracts == nil {
		l1.LaneConfig.DestContracts = l2.LaneConfig.DestContracts
	}

	for chain, cfg := range l2.LaneConfig.DestContracts {
		l1.LaneConfig.DestContracts[chain] = cfg
	}
	l1.LaneConfig.CommonContracts = l2.LaneConfig.CommonContracts
}

var laneMu = &sync.Mutex{}

// ReadLane reads existing lane config from ./contracts.json
func ReadLane(networkA string) (*Lane, error) {
	var existingLanes []Lane
	if len(ExistingContracts) == 0 {
		return nil, nil
	}
	err := json.Unmarshal(ExistingContracts, &existingLanes)
	if err != nil {
		return nil, err
	}
	for _, c := range existingLanes {
		if strings.ToLower(c.NetworkA) == strings.ToLower(networkA) {
			if len(c.LaneConfig.BridgeTokens) != len(c.LaneConfig.BridgeTokenPools) {
				return nil, fmt.Errorf("no of pools and tokens should match")
			}
			return &c, nil
		}
	}
	return nil, nil
}

// UpdateLane reads existing lane config from ./contracts.json and adds/updates provided lane config
// the updated lane config is stored in a temporary file under working dir with the name of tempFileName
// if needed, the contents of ./contracts.json should be manually updated with the newly generated content of tempFileName
func UpdateLane(l1, l2 Lane, tempFileName string) error {
	laneMu.Lock()
	defer laneMu.Unlock()
	err := l1.Validate()
	if err != nil {
		return err
	}
	err = l2.Validate()
	if err != nil {
		return err
	}
	var existingLanes []Lane
	_, err = os.Stat(tempFileName)
	if errors.Is(err, os.ErrNotExist) {
		err := json.Unmarshal(ExistingContracts, &existingLanes)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else {
		err = common.LoadJSON(tempFileName, &existingLanes)
		if err != nil {
			return err
		}
	}

	l1index := -1
	l2index := -1
	for i, c := range existingLanes {
		if c.NetworkA == l1.NetworkA {
			existingLanes[i].copyConfig(l1)
			l1index = i
		}
		if c.NetworkA == l2.NetworkA {
			existingLanes[i].copyConfig(l2)
			l2index = i
		}
	}

	if l1index == -1 {
		existingLanes = append(existingLanes, l1)
	}

	if l2index == -1 {
		existingLanes = append(existingLanes, l2)
	}

	b, err := json.MarshalIndent(existingLanes, "", " ")
	if err != nil {
		return err
	}
	file, err := os.Create(tempFileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(b)
	if err != nil {
		return err
	}
	return nil
}
