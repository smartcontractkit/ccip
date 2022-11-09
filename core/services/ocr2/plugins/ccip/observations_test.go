package ccip

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/blob_verifier"
	"github.com/smartcontractkit/chainlink/core/logger"
)

func TestObservationFilter(t *testing.T) {
	lggr := logger.TestLogger(t)
	obs1 := RelayObservation{IntervalsByOnRamp: map[common.Address]blob_verifier.CCIPInterval{
		common.HexToAddress("0x5431F5F973781809D18643b87B44921b11355d81"): blob_verifier.CCIPInterval{Min: 1, Max: 10},
	}}
	b1, err := obs1.Marshal()
	require.NoError(t, err)
	nonEmpty := getNonEmptyObservations[RelayObservation](lggr, []types.AttributedObservation{{Observation: b1}, {Observation: []byte{}}})
	require.NoError(t, err)
	require.Equal(t, 1, len(nonEmpty))
	assert.Equal(t, nonEmpty[0].IntervalsByOnRamp, obs1.IntervalsByOnRamp)
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
