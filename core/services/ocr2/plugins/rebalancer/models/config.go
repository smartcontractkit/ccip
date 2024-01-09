package models

const PluginName = "liquidityRebalancer"

type PluginConfig struct {
	LiquidityManagerAddress Address         `json:"liquidityManagerAddress"`
	LiquidityManagerNetwork NetworkSelector `json:"liquidityManagerNetwork"`
	ClosePluginTimeoutSec   int             `json:"closePluginTimeoutSec"`
}
