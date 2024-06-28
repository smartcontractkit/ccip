package ccipevm

import (
	"math/rand"
	"testing"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/require"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
)

var randomReport = cciptypes.CommitPluginReport{
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

func TestCommitPluginCodec(t *testing.T) {
	commitCodec := NewCommitPluginCodecV1()
	ctx := testutils.Context(t)
	encodedReport, err := commitCodec.Encode(ctx, randomReport)
	require.NoError(t, err)
	decodedReport, err := commitCodec.Decode(ctx, encodedReport)
	require.NoError(t, err)
	require.Equal(t, randomReport, decodedReport)
}

func BenchmarkCommitPluginCodec_Encode(b *testing.B) {
	commitCodec := NewCommitPluginCodecV1()
	ctx := testutils.Context(b)

	for i := 0; i < b.N; i++ {
		_, err := commitCodec.Encode(ctx, randomReport)
		require.NoError(b, err)
	}
}

func BenchmarkCommitPluginCodec_Decode(b *testing.B) {
	commitCodec := NewCommitPluginCodecV1()
	ctx := testutils.Context(b)
	encodedReport, err := commitCodec.Encode(ctx, randomReport)
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		_, err := commitCodec.Decode(ctx, encodedReport)
		require.NoError(b, err)
	}
}
