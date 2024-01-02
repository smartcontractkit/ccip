package models

import (
	"encoding/json"
	"fmt"
	"math/big"
)

// TODO: rename chain -> net
type ChainLiquidity struct {
	Chain     NetworkID
	Liquidity *big.Int
}

func NewChainLiquidity(chain NetworkID, liq *big.Int) ChainLiquidity {
	return ChainLiquidity{
		Chain:     chain,
		Liquidity: liq,
	}
}

type Observation struct {
	LiquidityPerChain []ChainLiquidity
	PendingTransfers  []PendingTransfer
}

func NewObservation(liqPerChain []ChainLiquidity, pendingTransfers []PendingTransfer) Observation {
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
