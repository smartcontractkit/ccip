package ccip

import (
	"github.com/pkg/errors"
)

const (
	MaxObservationLength = 250_000 // plugins's Observation should make sure to cap to this limit
)

var ErrCommitStoreIsDown = errors.New("commitStoreReader is down")
