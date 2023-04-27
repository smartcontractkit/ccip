package ccip

import (
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

func TestCommitOffchainConfig_Encoding(t *testing.T) {
	tests := map[string]struct {
		want    CommitOffchainConfig
		wantErr error
	}{
		"happy flow": {
			want: CommitOffchainConfig{
				SourceIncomingConfirmations: 3,
				DestIncomingConfirmations:   6,
				FeeUpdateHeartBeat:          models.MustMakeDuration(1 * time.Hour),
				FeeUpdateDeviationPPB:       5e7,
				MaxGasPrice:                 200e9,
				InflightCacheExpiry:         models.MustMakeDuration(23456 * time.Second),
			},
		},
		"missing not required fields": {
			want: CommitOffchainConfig{
				SourceIncomingConfirmations: 99999999,
				InflightCacheExpiry:         models.MustMakeDuration(23456 * time.Second),
			},
		},
		"missing required fields": {
			want: CommitOffchainConfig{
				SourceIncomingConfirmations: 3,
				DestIncomingConfirmations:   6,
				FeeUpdateHeartBeat:          models.MustMakeDuration(1 * time.Hour),
				FeeUpdateDeviationPPB:       5e7,
				MaxGasPrice:                 200e9,
			},
			wantErr: errors.New("InflightCacheExpiry not set"),
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			encode, err := EncodeOffchainConfig(tc.want)
			require.NoError(t, err)
			got, err := DecodeOffchainConfig[CommitOffchainConfig](encode)

			if tc.wantErr != nil {
				require.Error(t, tc.wantErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.want, got)
			}
		})
	}
}

func TestExecOffchainConfig_Encoding(t *testing.T) {
	tests := map[string]struct {
		want    ExecOffchainConfig
		wantErr error
	}{
		"happy flow": {
			want: ExecOffchainConfig{
				SourceIncomingConfirmations: 3,
				DestIncomingConfirmations:   6,
				BatchGasLimit:               5_000_000,
				RelativeBoostPerWaitHour:    0.07,
				MaxGasPrice:                 200e9,
				InflightCacheExpiry:         models.MustMakeDuration(64 * time.Second),
				RootSnoozeTime:              models.MustMakeDuration(128 * time.Minute),
			},
		},
		"missing not required fields": {
			want: ExecOffchainConfig{
				SourceIncomingConfirmations: 99999999,
				InflightCacheExpiry:         models.MustMakeDuration(64 * time.Second),
				RootSnoozeTime:              models.MustMakeDuration(128 * time.Minute),
			},
		},
		"missing required fields": {
			want: ExecOffchainConfig{
				SourceIncomingConfirmations: 99999999,
				InflightCacheExpiry:         models.MustMakeDuration(64 * time.Second),
			},
			wantErr: errors.New("RootSnoozeTime not set"),
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			encode, err := EncodeOffchainConfig(tc.want)
			require.NoError(t, err)
			got, err := DecodeOffchainConfig[ExecOffchainConfig](encode)

			if tc.wantErr != nil {
				require.Error(t, tc.wantErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.want, got)
			}
		})
	}
}
