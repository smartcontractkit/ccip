package cciptypes

import "C"
import (
	"encoding/json"
	"fmt"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

type CommitPluginConfig struct {
	// DestChain is the ccip destination chain configured for the commit plugin DON.
	DestChain ChainSelector `json:"destChain"`

	// PricedTokens is a list of tokens that we want to submit price updates for.
	PricedTokens []types.Account `json:"pricedTokens"`

	// TokenPricesObserver indicates that the node can observe token prices.
	TokenPricesObserver bool `json:"tokenPricesObserver"`

	// NewMsgScanBatchSize is the number of max new messages to scan, typically set to 256.
	NewMsgScanBatchSize int `json:"newMsgScanBatchSize"`
}

func (c CommitPluginConfig) Validate() error {
	if c.DestChain == ChainSelector(0) {
		return fmt.Errorf("destChain not set")
	}

	if len(c.PricedTokens) == 0 {
		return fmt.Errorf("priced tokens not set, at least one priced token is required")
	}

	if c.NewMsgScanBatchSize == 0 {
		return fmt.Errorf("newMsgScanBatchSize not set")
	}

	return nil
}

type ExecutePluginConfig struct {
	// DestChain is the ccip destination chain configured for the execute DON.
	DestChain ChainSelector

	// ObserverInfo is a map of oracle IDs to ObserverInfo.
	ObserverInfo map[commontypes.OracleID]ObserverInfo

	// MessageVisibilityInterval is the time interval for which the messages are visible by the plugin.
	MessageVisibilityInterval time.Duration
}

type ObserverInfo struct {
	// Writer indicates that the node can contribute by sending reports to the destination chain.
	// Being a Writer guarantees that the node can also read from the destination chain.
	Writer bool `json:"writer"`

	// Reads define the chains that the current node can read from.
	Reads []ChainSelector `json:"reads"`
}

type ConsensusObservation struct {
	// FChain defines the FChain value for each chain. FChain is used while forming consensus based on the observations.
	FChain map[ChainSelector]int `json:"fChain"`
	// PricedTokens is a list of tokens that we want to submit price updates for.
	PricedTokens []types.Account `json:"pricedTokens"`
	// NodeSupportedChains is a map of oracle IDs to SupportedChains.
	NodeSupportedChains map[commontypes.OracleID]SupportedChains `json:"nodeSupportedChains"`
}

func (c ConsensusObservation) Validate() error {
	for _, inf := range c.NodeSupportedChains {
		for ch := range inf.Supported.Iter() {
			if _, ok := c.FChain[ch]; !ok {
				return fmt.Errorf("fChain not set for chain %d", ch)
			}
		}
	}
	if len(c.PricedTokens) == 0 {
		return fmt.Errorf("priced tokens not set, at least one priced token is required")
	}

	return nil
}

type HomeChainConfig struct {
	// FChain defines the FChain value for each chain. FChain is used while forming consensus based on the observations.
	FChain map[ChainSelector]int `json:"fChain"`
	// NodeSupportedChains is a map of oracle IDs to SupportedChains.
	NodeSupportedChains map[commontypes.OracleID]SupportedChains `json:"nodeSupportedChains"`
}

func (c *HomeChainConfig) Validate() error {
	for _, inf := range c.NodeSupportedChains {
		for ch := range inf.Supported.Iter() {
			if _, ok := c.FChain[ch]; !ok {
				return fmt.Errorf("fChain not set for chain %d", ch)
			}
		}
	}
	return nil
}
func (c *HomeChainConfig) GetFChain(chain ChainSelector) int {
	return c.FChain[chain]
}

func (c *HomeChainConfig) IsSupported(node commontypes.OracleID, chain ChainSelector) bool {
	supportedChains, ok := c.NodeSupportedChains[node]
	if !ok {
		return false
	}
	return supportedChains.IsSupported(chain)
}

func (c *HomeChainConfig) GetSupportedChains(node commontypes.OracleID) mapset.Set[ChainSelector] {
	supportedChains, ok := c.NodeSupportedChains[node]
	if !ok {
		return mapset.NewSet[ChainSelector]()
	}
	return supportedChains.Supported
}

type SupportedChains struct {
	Supported mapset.Set[ChainSelector] `json:"supported"`
}

func (supportedChains *SupportedChains) IsSupported(chain ChainSelector) bool {
	return supportedChains.Supported.Contains(chain)
}

// UnmarshalJSON to convert the array to Set
func (sc *SupportedChains) UnmarshalJSON(data []byte) error {
	// Define a temporary struct with a slice for Supported
	temp := struct {
		Supported []ChainSelector `json:"supported"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Convert the slice to a mapset.Set
	sc.Supported = mapset.NewSet[ChainSelector]()
	for _, ch := range temp.Supported {
		sc.Supported.Add(ch)
	}

	return nil
}

type OnChainConfig struct {
	Readers []Bytes32 `json:"readers"`
	FChain  uint8     `json:"fChain"`
	Config  []byte    `json:"config"`
}
type OnChainCapabilityConfig struct {
	// Calling function https://github.com/smartcontractkit/ccip/blob/330c5e98f624cfb10108c92fe1e00ced6d345a99/contracts/src/v0.8/ccip/capability/CCIPCapabilityConfiguration.sol#L140
	ChainSelector ChainSelector `json:"chainSelector"`
	ChainConfig   OnChainConfig `json:"chainConfig"`
}
