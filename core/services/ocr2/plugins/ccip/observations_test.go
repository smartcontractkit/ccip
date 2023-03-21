package ccip

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/logger"
)

func TestObservationFilter(t *testing.T) {
	lggr := logger.TestLogger(t)
	obs1 := CommitObservation{Interval: commit_store.CommitStoreInterval{Min: 1, Max: 10}}
	b1, err := obs1.Marshal()
	require.NoError(t, err)
	nonEmpty := getNonEmptyObservations[CommitObservation](lggr, []types.AttributedObservation{{Observation: b1}, {Observation: []byte{}}})
	require.NoError(t, err)
	require.Equal(t, 1, len(nonEmpty))
	assert.Equal(t, nonEmpty[0].Interval, obs1.Interval)
}

func TestObservationSize(t *testing.T) {
	testParams := gopter.DefaultTestParameters()
	testParams.MinSuccessfulTests = 100
	p := gopter.NewProperties(testParams)
	p.Property("bounded observation size", prop.ForAll(func(min, max uint64) bool {
		o := ExecutionObservation{SeqNrs: []uint64{min, max}}
		b, err := o.Marshal()
		require.NoError(t, err)
		return len(b) <= MaxObservationLength
	}, gen.UInt64(), gen.UInt64()))
	p.TestingRun(t)
}
