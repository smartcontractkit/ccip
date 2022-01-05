package ccip

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func bytes32Equal(l, r [32]byte) bool {
	return bytes.Equal(l[:], r[:])
}

func TestMerkleProver(t *testing.T) {
	zhs := computeZeroHashes(2)
	require.Equal(t, 2, len(zhs))
	var zh [32]byte
	assert.True(t, bytes32Equal(zh, zhs[0]))
	assert.True(t, bytes32Equal(hashInternal(zh, zh), zhs[1]))
	zhs = computeZeroHashes(32)
	require.Equal(t, 32, len(zhs))

	leaves := make([][]byte, 2)
	leaves[0] = []byte{0xaa}
	leaves[1] = []byte{0xbb}

	// With a tree height of 2 and 2 elements, the root should simply be
	// h(h(leaf0) || h(leaf1))
	root, proof := GenerateMerkleProof(2, leaves, 0)
	assert.True(t, bytes32Equal(root, hashInternal(HashLeaf(leaves[0]), HashLeaf(leaves[1]))))
	assert.True(t, bytes32Equal(root, GenerateMerkleRoot(leaves[0], proof)))

	// With a tree height of 3 and 2 elements, we expect
	// h((h(leaf0) || h(leaf1)) || h(0 || 0))
	root, proof = GenerateMerkleProof(3, leaves, 0)
	assert.True(t, bytes32Equal(root,
		hashInternal(hashInternal(HashLeaf(leaves[0]), HashLeaf(leaves[1])), hashInternal(zh, zh))))
	assert.True(t, bytes32Equal(root, GenerateMerkleRoot(leaves[0], proof)))

	// One element tree height 2
	root, proof = GenerateMerkleProof(2, leaves[:1], 0)
	assert.True(t, bytes32Equal(root,
		hashInternal(HashLeaf(leaves[0]), zh)))
	assert.True(t, bytes32Equal(zh, proof.path[0]))
	assert.True(t, bytes32Equal(zh, proof.path[0]))
	assert.True(t, bytes32Equal(root, GenerateMerkleRoot(leaves[0], proof)))
}
