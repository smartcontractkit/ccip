package cciptypes

import (
	"encoding/hex"
	"strings"
)

type Address string

type Hash [32]byte

func (h Hash) String() string {
	res := hex.EncodeToString(h[:])
	if !strings.HasPrefix(res, "0x") {
		return "0x" + res
	}
	return res
}

type TxMeta struct {
	BlockTimestampUnixMilli int64
	BlockNumber             uint64
	TxHash                  string
	LogIndex                uint64
}
