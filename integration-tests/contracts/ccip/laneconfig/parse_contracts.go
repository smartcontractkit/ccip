package laneconfig

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/rs/zerolog/log"
)

//go:embed contracts.json
var ExistingContracts []byte

type LaneConfig struct {
	OtherNetwork     string   `json:"OtherNetwork"`
	FeeToken         string   `json:"FeeToken"`
	FeeTokenPool     string   `json:"FeeTokenPool"`
	BridgeTokens     []string `json:"BridgeTokens"`
	BridgeTokenPools []string `json:"BridgeTokenPools"`
	AFN              string   `json:"AFN"`
	Router           string   `json:"Router"`
}

type Lane struct {
	NetworkA   string     `json:"NetworkA"`
	LaneConfig LaneConfig `json:"LaneConfig"`
}

// ReadLane reads existing lane config from ./contracts.json
func ReadLane(networkA, networkB string, laneMu *sync.Mutex) (*Lane, error) {
	laneMu.Lock()
	defer laneMu.Unlock()
	var existingLanes []Lane
	err := json.Unmarshal(ExistingContracts, &existingLanes)
	if err != nil {
		return nil, err
	}
	for _, c := range existingLanes {
		if strings.ToLower(c.NetworkA) == strings.ToLower(networkA) &&
			strings.ToLower(c.LaneConfig.OtherNetwork) == strings.ToLower(networkB) {
			if len(c.LaneConfig.BridgeTokens) != len(c.LaneConfig.BridgeTokenPools) {
				return nil, fmt.Errorf("no of pools and tokens should match")
			}
			return &c, nil
		}
	}
	return nil, nil
}

// UpdateLane reads existing lane config from ./contracts.json and adds/updates provided lane config
// the updated lane config is stored in a temporary file under working dir with the name of tmpcontracts.json
// if needed, the contents of ./contracts.json should be manually updated with the newly generated content of tmpcontracts.json
func UpdateLane(l Lane, laneMu *sync.Mutex) error {
	laneMu.Lock()
	defer laneMu.Unlock()
	var existingLanes []Lane
	err := json.Unmarshal(ExistingContracts, &existingLanes)
	if err != nil {
		return err
	}
	log.Info().Msgf("Updating lane %s -> %s", l.NetworkA, l.LaneConfig.OtherNetwork)
	index := -1
	for i, c := range existingLanes {
		if c.NetworkA == l.NetworkA && c.LaneConfig.OtherNetwork == l.LaneConfig.OtherNetwork {
			existingLanes[i].LaneConfig = l.LaneConfig
			index = i
		}
	}

	if index == -1 {
		existingLanes = append(existingLanes, l)
	}
	b, err := json.MarshalIndent(existingLanes, "", " ")
	if err != nil {
		return err
	}
	file, err := os.Create("./tmpcontracts.json")
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
