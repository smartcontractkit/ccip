package ccipevm

import (
	"math/big"
	"math/rand"
	"testing"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
)

var randomReport = func() cciptypes.CommitPluginReport {
	return cciptypes.CommitPluginReport{
		MerkleRoots: []cciptypes.MerkleRootChain{
			{
				ChainSel: cciptypes.ChainSelector(rand.Uint64()),
				SeqNumsRange: cciptypes.NewSeqNumRange(
					cciptypes.SeqNum(rand.Uint64()),
					cciptypes.SeqNum(rand.Uint64()),
				),
				MerkleRoot: utils.RandomBytes32(),
			},
			{
				ChainSel: cciptypes.ChainSelector(rand.Uint64()),
				SeqNumsRange: cciptypes.NewSeqNumRange(
					cciptypes.SeqNum(rand.Uint64()),
					cciptypes.SeqNum(rand.Uint64()),
				),
				MerkleRoot: utils.RandomBytes32(),
			},
		},
		PriceUpdates: cciptypes.PriceUpdates{
			TokenPriceUpdates: []cciptypes.TokenPrice{
				{
					TokenID: types.Account(utils.RandomAddress().String()),
					Price:   cciptypes.NewBigInt(utils.RandUint256()),
				},
			},
			GasPriceUpdates: []cciptypes.GasPriceChain{
				{GasPrice: cciptypes.NewBigInt(utils.RandUint256()), ChainSel: cciptypes.ChainSelector(rand.Uint64())},
				{GasPrice: cciptypes.NewBigInt(utils.RandUint256()), ChainSel: cciptypes.ChainSelector(rand.Uint64())},
				{GasPrice: cciptypes.NewBigInt(utils.RandUint256()), ChainSel: cciptypes.ChainSelector(rand.Uint64())},
			},
		},
	}
}

func TestCommitPluginCodec(t *testing.T) {
	testCases := []struct {
		name   string
		report func(report cciptypes.CommitPluginReport) cciptypes.CommitPluginReport
		expErr bool
	}{
		{
			name: "base report",
			report: func(report cciptypes.CommitPluginReport) cciptypes.CommitPluginReport {
				return report
			},
		},
		{
			name: "empty token address",
			report: func(report cciptypes.CommitPluginReport) cciptypes.CommitPluginReport {
				report.PriceUpdates.TokenPriceUpdates[0].TokenID = ""
				return report
			},
			expErr: true,
		},
		{
			name: "empty merkle root",
			report: func(report cciptypes.CommitPluginReport) cciptypes.CommitPluginReport {
				report.MerkleRoots[0].MerkleRoot = cciptypes.Bytes32{}
				return report
			},
		},
		{
			name: "zero token price",
			report: func(report cciptypes.CommitPluginReport) cciptypes.CommitPluginReport {
				report.PriceUpdates.TokenPriceUpdates[0].Price = cciptypes.NewBigInt(big.NewInt(0))
				return report
			},
		},
		{
			name: "zero gas price",
			report: func(report cciptypes.CommitPluginReport) cciptypes.CommitPluginReport {
				report.PriceUpdates.GasPriceUpdates[0].GasPrice = cciptypes.NewBigInt(big.NewInt(0))
				return report
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			report := tc.report(randomReport())
			commitCodec := NewCommitPluginCodecV1()
			ctx := testutils.Context(t)
			encodedReport, err := commitCodec.Encode(ctx, report)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			decodedReport, err := commitCodec.Decode(ctx, encodedReport)
			require.NoError(t, err)
			require.Equal(t, report, decodedReport)
		})
	}
}

func BenchmarkCommitPluginCodec_Encode(b *testing.B) {
	commitCodec := NewCommitPluginCodecV1()
	ctx := testutils.Context(b)

	rep := randomReport()
	for i := 0; i < b.N; i++ {
		_, err := commitCodec.Encode(ctx, rep)
		require.NoError(b, err)
	}
}

func BenchmarkCommitPluginCodec_Decode(b *testing.B) {
	commitCodec := NewCommitPluginCodecV1()
	ctx := testutils.Context(b)
	encodedReport, err := commitCodec.Encode(ctx, randomReport())
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		_, err := commitCodec.Decode(ctx, encodedReport)
		require.NoError(b, err)
	}
}