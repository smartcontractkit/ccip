package ccip

import (
	"math"
	"math/rand"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

func Test_mergeEpochAndRound(t *testing.T) {
	type args struct {
		epoch uint32
		round uint8
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "zero round and epoch",
			args: args{epoch: 0, round: 0},
			want: 0,
		},
		{
			name: "avg case",
			args: args{
				epoch: 243,
				round: 15,
			},
			want: 62223,
		},
		{
			name: "largest epoch and round",
			args: args{
				epoch: math.MaxUint32,
				round: math.MaxUint8,
			},
			want: 1099511627775,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want,
				mergeEpochAndRound(tt.args.epoch, tt.args.round),
				"mergeEpochAndRound(%v, %v)", tt.args.epoch, tt.args.round)
		})
	}
}

func Test_bytesOfBytesKeccak(t *testing.T) {
	h, err := bytesOfBytesKeccak(nil)
	assert.NoError(t, err)
	assert.Equal(t, [32]byte{}, h)

	h1, err := bytesOfBytesKeccak([][]byte{{0x1}, {0x1}})
	assert.NoError(t, err)
	h2, err := bytesOfBytesKeccak([][]byte{{0x1, 0x1}})
	assert.NoError(t, err)
	assert.NotEqual(t, h1, h2)
}

func Test_contiguousReqs(t *testing.T) {
	testCases := []struct {
		min    uint64
		max    uint64
		seqNrs []uint64
		exp    bool
	}{
		{min: 5, max: 10, seqNrs: []uint64{5, 6, 7, 8, 9, 10}, exp: true},
		{min: 5, max: 10, seqNrs: []uint64{5, 7, 8, 9, 10}, exp: false},
		{min: 5, max: 10, seqNrs: []uint64{5, 6, 7, 8, 9, 10, 11}, exp: false},
		{min: 5, max: 10, seqNrs: []uint64{}, exp: false},
		{min: 1, max: 1, seqNrs: []uint64{1}, exp: true},
	}

	for _, tc := range testCases {
		res := contiguousReqs(logger.NullLogger, tc.min, tc.max, tc.seqNrs)
		assert.Equal(t, tc.exp, res)
	}
}

func Test_getMessageIDsAsHexString(t *testing.T) {
	t.Run("base", func(t *testing.T) {
		hashes := make([]common.Hash, 10)
		for i := range hashes {
			hashes[i] = common.HexToHash(strconv.Itoa(rand.Intn(100000)))
		}

		msgs := make([]evm_2_evm_offramp.InternalEVM2EVMMessage, len(hashes))
		for i := range msgs {
			msgs[i] = evm_2_evm_offramp.InternalEVM2EVMMessage{MessageId: hashes[i]}
		}

		messageIDs := getMessageIDsAsHexString(msgs)
		for i := range messageIDs {
			assert.Equal(t, hashes[i].String(), messageIDs[i])
		}
	})

	t.Run("empty", func(t *testing.T) {
		messageIDs := getMessageIDsAsHexString(nil)
		assert.Empty(t, messageIDs)
	})
}
