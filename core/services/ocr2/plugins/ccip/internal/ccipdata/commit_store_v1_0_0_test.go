package ccipdata

import (
	"math/big"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

func TestCommitReportEncoding(t *testing.T) {
	report := CommitStoreReport{
		TokenPrices: []TokenPrice{
			{
				Token: utils.RandomAddress(),
				Value: big.NewInt(9e18),
			},
		},
		GasPrices: []GasPrice{
			{
				DestChainSelector: rand.Uint64(),
				Value:             big.NewInt(2000e9),
			},
		},
		MerkleRoot: [32]byte{123},
		Interval:   CommitStoreInterval{Min: 1, Max: 10},
	}

	c := CommitStoreV1_0_0{}
	encodedReport, err := c.EncodeCommitReport(report)
	require.NoError(t, err)

	decodedReport, err := c.DecodeCommitReport(encodedReport)
	require.NoError(t, err)
	require.Equal(t, report, decodedReport)
}
