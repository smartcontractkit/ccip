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
	computedRoot, err := VerifyComputeRoot(NewKeccakCtx(), proofLeaves, Proof{
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
	h, err := VerifyComputeRoot(ctx, []Hash{[]byte{0xa}}, p)
	require.NoError(t, err)
	assert.Equal(t, tr8.Root(), h)
	expected := ctx.HashInternal(ctx.HashInternal(ctx.HashInternal([]byte{0xa}, []byte{0xb}), ctx.HashInternal([]byte{0xc}, []byte{0xd})), ctx.HashInternal(ctx.HashInternal([]byte{0xe}, ctx.ZeroHash()), ctx.ZeroHash()))
	assert.Equal(t, expected, tr8.Root())
}

func TestMerkleMultiProofSecondPreimage(t *testing.T) {
	ctx := NewKeccakCtx()
	leaves := []Hash{ctx.HashLeaf([]byte{0xa}), ctx.HashLeaf([]byte{0xb})}
	tr := NewTree(ctx, leaves)
	root, err := VerifyComputeRoot(ctx, []Hash{ctx.HashLeaf([]byte{0xa})}, tr.Prove([]int{0}))
	require.NoError(t, err)
	assert.Equal(t, root, tr.Root())

	secondPreimage := append(ctx.HashLeaf([]byte{0xa}), ctx.HashLeaf([]byte{0xb})...)
	leaves2 := []Hash{ctx.HashLeaf(secondPreimage)}
	tr2 := NewTree(ctx, leaves2)
	assert.NotEqual(t, tr2.Root(), tr.Root())
}

func TestMerkleMultiProof(t *testing.T) {
	ctx := NewKeccakCtx()
	leaves := []Hash{[]byte{0xa}, []byte{0xb}, []byte{0xc}, []byte{0xd}, []byte{0xe}, []byte{0xf}}
	expectedRoots := []Hash{
		[]byte{0xa},
		ctx.HashInternal([]byte{0xa}, []byte{0xb}),
		ctx.HashInternal(ctx.HashInternal([]byte{0xa}, []byte{0xb}), ctx.HashInternal([]byte{0xc}, ctx.ZeroHash())),
		ctx.HashInternal(ctx.HashInternal([]byte{0xa}, []byte{0xb}), ctx.HashInternal([]byte{0xc}, []byte{0xd})),
		ctx.HashInternal(ctx.HashInternal(ctx.HashInternal([]byte{0xa}, []byte{0xb}), ctx.HashInternal([]byte{0xc}, []byte{0xd})), ctx.HashInternal(ctx.HashInternal([]byte{0xe}, ctx.ZeroHash()), ctx.ZeroHash())),
		ctx.HashInternal(ctx.HashInternal(ctx.HashInternal([]byte{0xa}, []byte{0xb}), ctx.HashInternal([]byte{0xc}, []byte{0xd})), ctx.HashInternal(ctx.HashInternal([]byte{0xe}, []byte{0xf}), ctx.ZeroHash())),
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
				root, err := VerifyComputeRoot(ctx, leavesToProve, proof)
				require.NoError(t, err)
				assert.Equal(t, expectedRoot, root)
			}
		}
	}
}
