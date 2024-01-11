package v1_0_0

import (
	"math/big"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
)

func TestCommitReportEncoding(t *testing.T) {
	report := ccipdata.CommitStoreReport{
		TokenPrices: []ccipdata.TokenPrice{
			{
				Token: utils.RandomAddress(),
				Value: big.NewInt(9e18),
			},
		},
		GasPrices: []ccipdata.GasPrice{
			{
				DestChainSelector: rand.Uint64(),
				Value:             big.NewInt(2000e9),
			},
		},
		MerkleRoot: [32]byte{123},
		Interval:   ccipdata.CommitStoreInterval{Min: 1, Max: 10},
	}

	c, err := NewCommitStore(logger.TestLogger(t), utils.RandomAddress(), nil, mocks.NewLogPoller(t), nil)
	assert.NoError(t, err)

	encodedReport, err := c.EncodeCommitReport(report)
	require.NoError(t, err)
	assert.Greater(t, len(encodedReport), 0)

	decodedReport, err := c.DecodeCommitReport(encodedReport)
	require.NoError(t, err)
	require.Equal(t, report.TokenPrices, decodedReport.TokenPrices)
	require.Equal(t, report, decodedReport)
}
