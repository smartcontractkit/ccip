package models

import (
	"errors"
	"fmt"
	"slices"
)

const PluginName = "liquidityRebalancer"

type PluginConfig struct {
	LiquidityManagerAddress Address          `json:"liquidityManagerAddress"`
	LiquidityManagerNetwork NetworkSelector  `json:"liquidityManagerNetwork,string"`
	ClosePluginTimeoutSec   int              `json:"closePluginTimeoutSec"`
	RebalancerConfig        RebalancerConfig `json:"rebalancerConfig"`
	Tokens                  []string         `json:"tokens,omitempty"`
}

func (pc PluginConfig) TokenIDs() []TokenID {
	tokens := make([]TokenID, 0, len(pc.Tokens))
	for _, token := range pc.Tokens {
		tokens = append(tokens, NewTokenID(token))
	}
	return tokens
}

type RebalancerConfig struct {
	Type string `json:"type"`
}

func ValidateRebalancerConfig(config RebalancerConfig) error {
	if config.Type == "" {
		return errors.New("rebalancerType must be provided")
	}

	if !RebalancerIsSupported(config.Type) {
		return fmt.Errorf("rebalancerType %s is not supported, supported types are %+v", config.Type, AllRebalancerTypes)
	}

	return nil
}

const (
	RebalancerTypeMinLiquidity = "min-liquidity"
	RebalancerTypePingPong     = "ping-pong"
)

var (
	AllRebalancerTypes = []string{
		RebalancerTypePingPong,
		RebalancerTypeMinLiquidity,
	}
)

func RebalancerIsSupported(rebalancerType string) bool {
	return slices.Contains(AllRebalancerTypes, rebalancerType)
}
