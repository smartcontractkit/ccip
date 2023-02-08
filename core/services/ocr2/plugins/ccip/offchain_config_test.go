package ccip

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/store/models"
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
				FeeUpdateHeartBeat:          models.MustMakeDuration(1 * time.Hour),
				FeeUpdateDeviationPPB:       5e7,
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
