package cciptypes

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Address string

func (a Address) Equals(addr2 Address) bool {
	if common.IsHexAddress(string(a)) && common.IsHexAddress(string(addr2)) {
		return common.HexToAddress(string(a)) == common.HexToAddress(string(addr2))
	}
	return a == addr2
}

func (a Address) ToEVM() (common.Address, error) {
	if !common.IsHexAddress(string(a)) {
		return common.Address{}, fmt.Errorf("%s is not a hex evm address", a)
	}
	return common.HexToAddress(string(a)), nil
}

type Hash [32]byte

func (h Hash) String() string {
	return hexutil.Encode(h[:])
}

type BlockMeta struct {
	BlockTimestamp time.Time
	BlockNumber    int64
	TxHash         string
	LogIndex       uint
}
