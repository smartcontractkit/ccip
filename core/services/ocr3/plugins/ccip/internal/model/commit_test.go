package model

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommitPluginObservation_EncodeAndDecode(t *testing.T) {
	obs := NewCommitPluginObservation(
		[]CCIPMsgBaseDetails{
			{ID: [32]byte{1}, SourceChain: math.MaxUint64, SeqNum: 123},
			{ID: [32]byte{2}, SourceChain: 321, SeqNum: math.MaxUint64},
		},
		[]GasPriceChain{},
		[]TokenPrice{},
		[]SeqNumChain{},
	)

	b, err := obs.Encode()
	assert.NoError(t, err)
	assert.Equal(t, `{"newMsgs":[{"id":[1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],"sourceChain":"18446744073709551615","seqNum":"123"},{"id":[2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],"sourceChain":"321","seqNum":"18446744073709551615"}],"gasPrices":[],"tokenPrices":[],"maxSeqNums":[]}`, string(b))

	obs2, err := DecodeCommitPluginObservation(b)
	assert.NoError(t, err)
	assert.Equal(t, obs, obs2)
}
