package ccipdata

import (
	"context"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/cciptypes"
)

const (
	ManuallyExecute = "manuallyExecute"
)

//go:generate mockery --quiet --name OffRampReader --filename offramp_reader_mock.go --case=underscore
type OffRampReader interface {
	cciptypes.OffRampReader
	//TODO Move to chainlink-common
	GetSendersNonce(ctx context.Context, senders []cciptypes.Address) (map[cciptypes.Address]uint64, error)
}
