package models

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type NetworkLiquidity struct {
	Network   NetworkID
	Liquidity *big.Int
}

func NewNetworkLiquidity(chain NetworkID, liq *big.Int) NetworkLiquidity {
	return NetworkLiquidity{
		Network:   chain,
		Liquidity: liq,
	}
}

type Observation struct {
	LiquidityPerChain []NetworkLiquidity
	PendingTransfers  []PendingTransfer
}

func NewObservation(liqPerChain []NetworkLiquidity, pendingTransfers []PendingTransfer) Observation {
	return Observation{
		LiquidityPerChain: liqPerChain,
		PendingTransfers:  pendingTransfers,
	}
}

func (o Observation) Encode() []byte {
	b, err := json.Marshal(o)
	if err != nil {
		panic(fmt.Errorf("observation %#v encoding unexpected internal error: %w", o, err))
	}
	return b
}

func DecodeObservation(b []byte) (Observation, error) {
	var obs Observation
	err := json.Unmarshal(b, &obs)
	return obs, err
}
