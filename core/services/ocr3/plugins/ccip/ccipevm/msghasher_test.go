package ccipevm

import (
	"context"
	"testing"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/stretchr/testify/assert"
)

func TestMessageHasher_Hash(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name   string
		msg    cciptypes.CCIPMsg
		exp    string
		expErr bool
	}{
		{
			name: "empty msg",
			msg: cciptypes.CCIPMsg{
				ChainFeeLimit:  cciptypes.NewBigIntFromInt64(0),
				FeeTokenAmount: cciptypes.NewBigIntFromInt64(0),
			},
			exp:    "0xcc876be60dc969a6b64da96b122acf2087f9d91ab10a6dd96e582dffedc4c540",
			expErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := NewMessageHasher([32]byte{})
			hash, err := m.Hash(ctx, tc.msg)
			if tc.expErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.exp, hash.String())
		})
	}
}
