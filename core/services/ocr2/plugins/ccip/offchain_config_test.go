package ccip

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

func TestCommitOffchainConfig_Encoding(t *testing.T) {
	tests := map[string]struct {
		want  CommitOffchainConfig
		error bool
	}{
		"happy flow": {
			want: CommitOffchainConfig{
				SourceIncomingConfirmations: 3,
				DestIncomingConfirmations:   6,
				FeeUpdateHeartBeat:          models.MustMakeDuration(1 * time.Hour),
				FeeUpdateDeviationPPB:       5e7,
				MaxGasPrice:                 200e9,
			},
		},
		"missing fields": {
			want: CommitOffchainConfig{
				SourceIncomingConfirmations: 99999999,
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			encode, err := EncodeOffchainConfig(tc.want)
			require.NoError(t, err)
			got, err := DecodeOffchainConfig[CommitOffchainConfig](encode)
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}
}

func TestExecOffchainConfig_Encoding(t *testing.T) {
	tests := map[string]struct {
		want  ExecOffchainConfig
		error bool
	}{
		"happy flow": {
			want: ExecOffchainConfig{
				SourceIncomingConfirmations: 3,
				DestIncomingConfirmations:   6,
				BatchGasLimit:               5_000_000,
				RelativeBoostPerWaitHour:    0.07,
				MaxGasPrice:                 200e9,
			},
		},
		"missing fields": {
			want: ExecOffchainConfig{
				SourceIncomingConfirmations: 99999999,
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			encode, err := EncodeOffchainConfig(tc.want)
			require.NoError(t, err)
			got, err := DecodeOffchainConfig[ExecOffchainConfig](encode)
			require.NoError(t, err)
			require.Equal(t, tc.want, got)
		})
	}
}
