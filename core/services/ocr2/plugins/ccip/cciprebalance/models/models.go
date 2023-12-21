package models

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Address common.Address

type NetworkID uint64

type Transfer struct {
	From   NetworkID
	To     NetworkID
	Amount *big.Int
}

func (t Transfer) String() string {
	return fmt.Sprintf("%v->%v %s", t.From, t.To, t.Amount.String())
}

type ReportMetadata struct {
	Transfer                Transfer
	LiquidityManagerAddress Address
	NetworkID               NetworkID
}
