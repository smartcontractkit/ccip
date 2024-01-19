package bridge

import (
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
)

// TODO: this functions should be moved out of rebalancer plugin.

func parseLogs[T any](logs []logpoller.Log, parseFunc func(log types.Log) (*T, error)) ([]Event[T], error) {
	parsed := make([]Event[T], 0, len(logs))
	for _, log := range logs {
		data, err := parseFunc(log.ToGethLog())
		if err != nil {
			return nil, fmt.Errorf("cannot parse log: %w", err)
		}
		parsed = append(parsed, Event[T]{
			Data: *data,
			Meta: Meta{
				BlockTimestamp: log.BlockTimestamp,
				BlockNumber:    log.BlockNumber,
				TxHash:         log.TxHash,
				LogIndex:       uint(log.LogIndex),
			},
		})
	}
	return parsed, nil
}

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
