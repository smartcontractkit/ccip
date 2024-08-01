package persistent

import (
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"
	"os"
	"testing"
)

func TestAddressBook(t *testing.T) {
	fn := "blah.json"
	_, err := os.Create(fn)
	require.NoError(t, err)
	os.WriteFile(fn, []byte("{}"), 0644)
	defer func() {
		require.NoError(t, os.Remove(fn))
	}()

	a := NewAddressBook(fn)
	addrs, err := a.Addresses()
	require.NoError(t, err)
	require.Equal(t, 0, len(addrs))

	err = a.Save(1, "0x1")
	addrs, err = a.Addresses()
	require.NoError(t, err)
	assert.DeepEqual(t, map[uint64]map[string]struct{}{1: {"0x1": {}}}, addrs)

	err = a.Save(1, "0x2")
	addrs, err = a.Addresses()
	require.NoError(t, err)
	assert.DeepEqual(t, map[uint64]map[string]struct{}{
		1: {"0x1": {}, "0x2": {}},
	}, addrs)

	// TODO: Maybe prevent chain collisions?
	err = a.Save(2, "0x2")
	addrs, err = a.Addresses()
	require.NoError(t, err)
	assert.DeepEqual(t, map[uint64]map[string]struct{}{
		1: {"0x1": {}, "0x2": {}},
		2: {"0x2": {}},
	}, addrs)
}
