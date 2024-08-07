package deployment

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"
)

func TestAddressBook(t *testing.T) {
	ab := NewMemoryAddressBook()
	err := ab.Save(1, "0x1", "OnRamp 1.0.0")
	require.NoError(t, err)
	// Duplicate address will error
	err = ab.Save(1, "0x1", "OnRamp 1.0.0")
	require.Error(t, err)
	// Distinct address same TV will not
	err = ab.Save(1, "0x2", "OnRamp 1.0.0")
	require.NoError(t, err)
	// Same address different chain will not error
	err = ab.Save(2, "0x1", "OnRamp 1.0.0")
	require.NoError(t, err)
	// We can save different versions of the same contract
	err = ab.Save(2, "0x2", "OnRamp 1.2.0")
	require.NoError(t, err)

	addresses, err := ab.Addresses()
	require.NoError(t, err)
	assert.DeepEqual(t, addresses, map[uint64]map[string]string{
		1: {
			"0x1": "OnRamp 1.0.0",
			"0x2": "OnRamp 1.0.0",
		},
		2: {
			"0x1": "OnRamp 1.0.0",
			"0x2": "OnRamp 1.2.0",
		},
	})

	// TODO: Further testing of merge etc.
}
