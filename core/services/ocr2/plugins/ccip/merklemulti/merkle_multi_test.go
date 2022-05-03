package merklemulti

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gonum.org/v1/gonum/stat/combin"

	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/merklemulti/fixtures"
)

func TestSpecFixture(t *testing.T) {
	var proofLeaves []Hash
	for _, leaf := range fixtures.TestVectorLeaves {
		proofLeaves = append(proofLeaves, hexutil.MustDecode(fmt.Sprintf("0x%s", leaf)))
	}
	var proofHashes []Hash
	for _, proofHash := range fixtures.TestVectorProof {
		proofHashes = append(proofHashes, hexutil.MustDecode(fmt.Sprintf("0x%s", proofHash)))
	}
	computedRoot, err := verifyComputeRoot(NewKeccakCtx(), proofLeaves, Proof{
		Hashes: proofHashes, SourceFlags: fixtures.TestVectorSourceFlags,
	})
	require.NoError(t, err)
	assert.Equal(t, hexutil.MustDecode(fmt.Sprintf("0x%s", fixtures.TestVectorExpectedRoot)), []byte(computedRoot))
}

func TestPadding(t *testing.T) {
	ctx := NewKeccakCtx()
	tr4 := NewTree(ctx, []Hash{[]byte{0xa}, []byte{0xb}, []byte{0xc}})
	assert.Equal(t, 4, len(tr4.layers[0]))
	tr8 := NewTree(ctx, []Hash{[]byte{0xa}, []byte{0xb}, []byte{0xc}, []byte{0xd}, []byte{0xe}})
	assert.Equal(t, 6, len(tr8.layers[0]))
	assert.Equal(t, 4, len(tr8.layers[1]))
	p := tr8.Prove([]int{0})
	h, err := verifyComputeRoot(ctx, []Hash{[]byte{0xa}}, p)
	require.NoError(t, err)
	assert.Equal(t, tr8.Root(), h)
	expected := ctx.hashInternal(ctx.hashInternal(ctx.hashInternal([]byte{0xa}, []byte{0xb}), ctx.hashInternal([]byte{0xc}, []byte{0xd})), ctx.hashInternal(ctx.hashInternal([]byte{0xe}, ctx.zeroHash), ctx.zeroHash))
	assert.Equal(t, expected, tr8.Root())
}

func TestMerkleMultiProof(t *testing.T) {
	ctx := NewKeccakCtx()
	leaves := []Hash{[]byte{0xa}, []byte{0xb}, []byte{0xc}, []byte{0xd}, []byte{0xe}, []byte{0xf}}
	expectedRoots := []Hash{
		[]byte{0xa},
		ctx.hashInternal([]byte{0xa}, []byte{0xb}),
		ctx.hashInternal(ctx.hashInternal([]byte{0xa}, []byte{0xb}), ctx.hashInternal([]byte{0xc}, ctx.zeroHash)),
		ctx.hashInternal(ctx.hashInternal([]byte{0xa}, []byte{0xb}), ctx.hashInternal([]byte{0xc}, []byte{0xd})),
		ctx.hashInternal(ctx.hashInternal(ctx.hashInternal([]byte{0xa}, []byte{0xb}), ctx.hashInternal([]byte{0xc}, []byte{0xd})), ctx.hashInternal(ctx.hashInternal([]byte{0xe}, ctx.zeroHash), ctx.zeroHash)),
		ctx.hashInternal(ctx.hashInternal(ctx.hashInternal([]byte{0xa}, []byte{0xb}), ctx.hashInternal([]byte{0xc}, []byte{0xd})), ctx.hashInternal(ctx.hashInternal([]byte{0xe}, []byte{0xf}), ctx.zeroHash)),
	}
	// For every size tree from 0..len(leaves)
	for len_ := 1; len_ <= len(leaves); len_++ {
		tr := NewTree(NewKeccakCtx(), leaves[:len_])
		t.Log(tr)
		expectedRoot := expectedRoots[len_-1]
		require.Equal(t, tr.Root(), expectedRoot)
		// Prove every subset of its leaves
		for k := 1; k <= len_; k++ {
			gen := combin.NewCombinationGenerator(len_, k)
			for gen.Next() {
				leaveIndices := gen.Combination(nil)
				t.Log("indices", leaveIndices)
				proof := tr.Prove(leaveIndices)
				var leavesToProve []Hash
				for _, idx := range leaveIndices {
					leavesToProve = append(leavesToProve, leaves[idx])
				}
				root, err := verifyComputeRoot(ctx, leavesToProve, proof)
				require.NoError(t, err)
				assert.Equal(t, expectedRoot, root)
			}
		}
	}
}
