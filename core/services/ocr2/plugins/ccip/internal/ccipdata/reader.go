package ccipdata

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

type Event[T any] struct {
	Data T
	Meta
}

type Meta struct {
	BlockTimestamp time.Time
	BlockNumber    int64
	TxHash         common.Hash
	LogIndex       uint
}

type Closer interface {
	Close(qopts ...pg.QOpt) error
}

// Client can be used to fetch CCIP related parsed on-chain data.
//
//go:generate mockery --quiet --name Reader --output . --filename reader_mock.go --inpackage --case=underscore
type Reader interface {
	// LatestBlock returns the latest known/parsed block of the underlying implementation.
	LatestBlock(ctx context.Context) (int64, error)
}
