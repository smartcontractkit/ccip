package cciptypes

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Address string

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

type TxMeta struct {
	BlockTimestampUnixMilli int64
	BlockNumber             uint64
	TxHash                  string
	LogIndex                uint64
}
