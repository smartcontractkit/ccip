package models

const PluginName = "liquidityBalance"

type PluginConfig struct {
	F                       int       `json:"f"`
	LiquidityManagerAddress Address   `json:"liquidityManagerAddress"`
	LiquidityManagerNetwork NetworkID `json:"liquidityManagerNetwork"`
}
