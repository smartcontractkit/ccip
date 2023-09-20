package tokendata

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
)

func TestCachedReader_ReadTokenData(t *testing.T) {
	mockReader := MockReader{}
	cachedReader := NewCachedReader(&mockReader)

	msgData := []byte("msgData")
	mockReader.On("ReadTokenData", mock.Anything, mock.Anything).Return(msgData, nil)

	msg := internal.EVM2EVMOnRampCCIPSendRequestedWithMeta{}

	// Call ReadTokenData twice, expect only one call to underlying reader
	data, err := cachedReader.ReadTokenData(nil, msg)
	require.NoError(t, err)
	require.Equal(t, msgData, data)

	// First time, calls the underlying reader
	mockReader.AssertNumberOfCalls(t, "ReadTokenData", 1)

	data, err = cachedReader.ReadTokenData(nil, msg)
	require.NoError(t, err)
	require.Equal(t, msgData, data)

	// Second time, get data from cache
	mockReader.AssertNumberOfCalls(t, "ReadTokenData", 1)
}
