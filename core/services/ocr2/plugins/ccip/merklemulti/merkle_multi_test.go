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

var (
	ctx              = NewKeccakCtx()
	a, b, c, d, e, f = ctx.HashLeaf([]byte{0xa}), ctx.HashLeaf([]byte{0xb}), ctx.HashLeaf([]byte{0xc}), ctx.HashLeaf([]byte{0xd}), ctx.HashLeaf([]byte{0xe}), ctx.HashLeaf([]byte{0xf})
)

func TestSpecFixture(t *testing.T) {
	var leafHashes [][32]byte
	for _, leaf := range fixtures.TestVectorLeaves {
		var leaf32 [32]byte
		copy(leaf32[:], hexutil.MustDecode(fmt.Sprintf("0x%s", leaf)))
		leafHashes = append(leafHashes, leaf32)
	}
	var proofHashes [][32]byte
	for _, proofHash := range fixtures.TestVectorProof {
		var proofHash32 [32]byte
		copy(proofHash32[:], hexutil.MustDecode(fmt.Sprintf("0x%s", proofHash)))
		proofHashes = append(proofHashes, proofHash32)
	}
	computedRoot, err := VerifyComputeRoot(NewKeccakCtx(), leafHashes, Proof[[32]byte]{
		Hashes: proofHashes, SourceFlags: fixtures.TestVectorSourceFlags,
	})
	require.NoError(t, err)
	assert.Equal(t, hexutil.MustDecode(fmt.Sprintf("0x%s", fixtures.TestVectorExpectedRoot)), computedRoot[:])
}

func TestPadding(t *testing.T) {
	tr4 := NewTree(ctx, [][32]byte{a, b, c})
	assert.Equal(t, 4, len(tr4.layers[0]))
	tr8 := NewTree(ctx, [][32]byte{a, b, c, d, e})
	assert.Equal(t, 6, len(tr8.layers[0]))
	assert.Equal(t, 4, len(tr8.layers[1]))
	p := tr8.Prove([]int{0})
	h, err := VerifyComputeRoot(ctx, [][32]byte{a}, p)
	require.NoError(t, err)
	assert.Equal(t, tr8.Root(), h)
	expected := ctx.HashInternal(ctx.HashInternal(ctx.HashInternal(a, b), ctx.HashInternal(c, d)), ctx.HashInternal(ctx.HashInternal(e, ctx.ZeroHash()), ctx.ZeroHash()))
	assert.Equal(t, expected, tr8.Root())
}

func TestMerkleMultiProofSecondPreimage(t *testing.T) {
	tr := NewTree(ctx, [][32]byte{a, b})
	root, err := VerifyComputeRoot(ctx, [][32]byte{a}, tr.Prove([]int{0}))
	require.NoError(t, err)
	assert.Equal(t, root, tr.Root())
	tr2 := NewTree(ctx, [][32]byte{ctx.HashLeaf(append(a[:], b[:]...))})
	assert.NotEqual(t, tr2.Root(), tr.Root())
}

func TestMerkleMultiProof(t *testing.T) {
	ctx := NewKeccakCtx()
	leafHashes := [][32]byte{a, b, c, d, e, f}
	expectedRoots := [][32]byte{
		a,
		ctx.HashInternal(a, b),
		ctx.HashInternal(ctx.HashInternal(a, b), ctx.HashInternal(c, ctx.ZeroHash())),
		ctx.HashInternal(ctx.HashInternal(a, b), ctx.HashInternal(c, d)),
		ctx.HashInternal(ctx.HashInternal(ctx.HashInternal(a, b), ctx.HashInternal(c, d)), ctx.HashInternal(ctx.HashInternal(e, ctx.ZeroHash()), ctx.ZeroHash())),
		ctx.HashInternal(ctx.HashInternal(ctx.HashInternal(a, b), ctx.HashInternal(c, d)), ctx.HashInternal(ctx.HashInternal(e, f), ctx.ZeroHash())),
	}
	// For every size tree from 0..len(leaves)
	for len_ := 1; len_ <= len(leafHashes); len_++ {
		tr := NewTree(NewKeccakCtx(), leafHashes[:len_])
		expectedRoot := expectedRoots[len_-1]
		require.Equal(t, tr.Root(), expectedRoot)
		// Prove every subset of its leaves
		for k := 1; k <= len_; k++ {
			gen := combin.NewCombinationGenerator(len_, k)
			for gen.Next() {
				leaveIndices := gen.Combination(nil)
				proof := tr.Prove(leaveIndices)
				var leavesToProve [][32]byte
				for _, idx := range leaveIndices {
					leavesToProve = append(leavesToProve, leafHashes[idx])
				}
				root, err := VerifyComputeRoot(ctx, leavesToProve, proof)
				require.NoError(t, err)
				assert.Equal(t, expectedRoot, root)
			}
		}
	}
}
