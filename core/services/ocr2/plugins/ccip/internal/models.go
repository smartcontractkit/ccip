package internal

import (
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
)

// EVM2EVMOnRampCCIPSendRequestedWithMeta helper struct to hold the send request and some metadata
type EVM2EVMOnRampCCIPSendRequestedWithMeta struct {
	ccipdata.EVM2EVMMessage
	BlockTimestamp time.Time
	Executed       bool
	Finalized      bool
	LogIndex       uint
	TxHash         common.Hash
}
