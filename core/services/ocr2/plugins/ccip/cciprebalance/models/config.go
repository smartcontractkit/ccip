package models

const PluginName = "liquidityBalance"

type PluginConfig struct {
	LiquidityManagerAddress Address   `json:"liquidityManagerAddress"`
	LiquidityManagerNetwork NetworkID `json:"liquidityManagerNetwork"`
}
