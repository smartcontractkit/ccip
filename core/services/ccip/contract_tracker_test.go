package ccip

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOffchainConfig_Encode_Decode(t *testing.T) {
	tests := map[string]struct {
		want  OffchainConfig
		error bool
	}{
		"Success": {
			want: OffchainConfig{
				SourceIncomingConfirmations: 3,
				DestIncomingConfirmations:   6,
			},
		},
		"Missing value as 0": {
			want: OffchainConfig{
				SourceIncomingConfirmations: 99999999,
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			encode, err := tc.want.Encode()
			got, err := Decode(encode)
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}
}
