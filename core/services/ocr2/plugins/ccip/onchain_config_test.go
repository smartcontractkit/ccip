package ccip

import (
	"math/big"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func randomAddress() common.Address {
	return common.BigToAddress(big.NewInt(rand.Int63()))
}

func TestCommitOnchainConfig(t *testing.T) {
	tests := []struct {
		name string
		want CommitOnchainConfig
	}{
		{
			name: "encodes and decodes config with all fields set",
			want: CommitOnchainConfig{
				PriceRegistry: randomAddress(),
				Afn:           randomAddress(),
			},
		},
		{
			name: "encodes and decodes config with missing fields",
			want: CommitOnchainConfig{
				PriceRegistry: randomAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded, err := EncodeAbiStruct(tt.want)
			require.NoError(t, err)

			decoded, err := DecodeAbiStruct(encoded, &CommitOnchainConfig{})
			require.NoError(t, err)
			require.Equal(t, tt.want, decoded)
		})
	}
}

func TestExecOnchainConfig(t *testing.T) {
	tests := []struct {
		name string
		want ExecOnchainConfig
	}{
		{
			name: "encodes and decodes config with all fields set",
			want: ExecOnchainConfig{
				PermissionLessExecutionThresholdSeconds: rand.Uint32(),
				Router:                                  randomAddress(),
				Afn:                                     randomAddress(),
				MaxTokensLength:                         uint16(rand.Uint32()),
				MaxDataSize:                             rand.Uint32(),
			},
		},
		{
			name: "encodes and decodes config with missing fields",
			want: ExecOnchainConfig{
				PermissionLessExecutionThresholdSeconds: rand.Uint32(),
				MaxDataSize:                             rand.Uint32(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded, err := EncodeAbiStruct(tt.want)
			require.NoError(t, err)

			decoded, err := DecodeAbiStruct(encoded, &ExecOnchainConfig{})
			require.NoError(t, err)
			require.Equal(t, tt.want, decoded)
		})
	}
}
