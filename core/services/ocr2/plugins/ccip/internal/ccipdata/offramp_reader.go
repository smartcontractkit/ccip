package ccipdata

import (
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
)

const (
	ManuallyExecute = "manuallyExecute"
)

// DefaultMsgVisibilityInterval is the default time interval for which a message is considered visible.
// The default will be set if the offramp's offchain config does not specify a visibility interval.
var DefaultMsgVisibilityInterval = 8 * time.Hour

//go:generate mockery --quiet --name OffRampReader --filename offramp_reader_mock.go --case=underscore
type OffRampReader interface {
	cciptypes.OffRampReader
}
