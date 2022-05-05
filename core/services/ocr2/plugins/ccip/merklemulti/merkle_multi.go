package merklemulti

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/core/utils"
)

type Hash []byte

type Ctx interface {
	HashLeaf(l []byte) Hash
	HashInternal(a, b Hash) Hash
	ZeroHash() Hash
}

type keccakCtx struct{}

func NewKeccakCtx() Ctx {
	return keccakCtx{}
}

func (k keccakCtx) HashLeaf(l []byte) Hash {
	// Note this Keccak256 cannot error https://github.com/golang/crypto/blob/master/sha3/sha3.go#L126
	// if we start supporting hashing algos which do, we can change this API to include an error.
	h, _ := utils.Keccak256(append([]byte{0x00}, l...))
	return h
}

func (k keccakCtx) HashInternal(a, b Hash) Hash {
	if bytes.Compare(a, b) < 0 {
		h, _ := utils.Keccak256(append([]byte{0x01}, append(a, b...)...))
		return h
	}
	h, _ := utils.Keccak256(append([]byte{0x01}, append(b, a...)...))
	return h
}

// We use empty bytes32 for zeroHash
// on the solidity side, this needs to match.
func (k keccakCtx) ZeroHash() Hash {
	var zeroes [32]byte
	return zeroes[:]
}

type singleLayerProof struct {
	nextIndices []int
	subProof    []Hash
	sourceFlags []bool
}

type Proof struct {
	Hashes      []Hash `json:"hashes"`
	SourceFlags []bool `json:"source_flags"`
}

func (p Proof) countSourceFlags(b bool) (count int) {
	for _, flag := range p.SourceFlags {
		if flag == b {
			count++
		}
	}
	return
}

const (
	SourceFromHashes = true
	SourceFromProof  = false
)

func parentIndex(idx int) int {
	return idx / 2
}

func siblingIndex(idx int) int {
	return idx ^ 1
}

func proveSingleLayer(layer []Hash, indices []int) singleLayerProof {
	var (
		authIndices []int
		nextIndices []int
		sourceFlags []bool
	)
	j := 0
	for j < len(indices) {
		x := indices[j]
		nextIndices = append(nextIndices, parentIndex(x))
		if j+1 < len(indices) && indices[j+1] == siblingIndex(x) {
			j++
			sourceFlags = append(sourceFlags, SourceFromHashes)
		} else {
			authIndices = append(authIndices, siblingIndex(x))
			sourceFlags = append(sourceFlags, SourceFromProof)
		}
		j++
	}
	var subProof []Hash
	for _, i := range authIndices {
		subProof = append(subProof, layer[i])
	}
	return singleLayerProof{
		nextIndices: nextIndices,
		subProof:    subProof,
		sourceFlags: sourceFlags,
	}
}

type Tree struct {
	layers [][]Hash
	ctx    Ctx
}

func NewTree(ctx Ctx, leafHashes []Hash) *Tree {
	var layer = make([]Hash, len(leafHashes))
	copy(layer, leafHashes)

	var layers = [][]Hash{layer}
	var curr int
	for len(layer) > 1 {
		paddedLayer, nextLayer := computeNextLayer(ctx, layer)
		layers[curr] = paddedLayer
		curr++
		layers = append(layers, nextLayer)
		layer = nextLayer
	}
	return &Tree{
		layers: layers,
	}
}

func (t *Tree) String() string {
	b := strings.Builder{}
	for _, layer := range t.layers {
		b.WriteString(fmt.Sprintf("%v", layer))
	}
	return b.String()
}

func (t *Tree) Root() Hash {
	return t.layers[len(t.layers)-1][0]
}

func (t *Tree) Prove(indices []int) Proof {
	var proof Proof
	for _, layer := range t.layers[:len(t.layers)-1] {
		res := proveSingleLayer(layer, indices)
		indices = res.nextIndices
		proof.Hashes = append(proof.Hashes, res.subProof...)
		proof.SourceFlags = append(proof.SourceFlags, res.sourceFlags...)
	}
	return proof
}

func computeNextLayer(ctx Ctx, layer []Hash) ([]Hash, []Hash) {
	if len(layer) == 1 {
		return layer, layer
	}
	if len(layer)%2 != 0 {
		layer = append(layer, ctx.ZeroHash())
	}
	var nextLayer []Hash
	for i := 0; i < len(layer); i += 2 {
		nextLayer = append(nextLayer, ctx.HashInternal(layer[i], layer[i+1]))
	}
	return layer, nextLayer
}

func VerifyComputeRoot(ctx Ctx, leaves []Hash, proof Proof) (Hash, error) {
	totalHashes := len(leaves) + len(proof.Hashes) - 1
	if totalHashes != len(proof.SourceFlags) {
		return nil, errors.Errorf("hashes %d != sourceFlags %d", totalHashes, len(proof.SourceFlags))
	}
	sourceProofCount := proof.countSourceFlags(SourceFromProof)
	if sourceProofCount != len(proof.Hashes) {
		return nil, errors.Errorf("proof source flags %d != proof hashes%d", sourceProofCount, len(proof.Hashes))
	}
	var hashes []Hash
	for i := 0; i < totalHashes; i++ {
		hashes = append(hashes, leaves[0])
	}
	var (
		leafPos  int
		hashPos  int
		proofPos int
	)
	for i := 0; i < totalHashes; i++ {
		var a, b Hash
		if proof.SourceFlags[i] == SourceFromHashes {
			if leafPos < len(leaves) {
				a = leaves[leafPos]
				leafPos++
			} else {
				a = hashes[hashPos]
				hashPos++
			}
		} else if proof.SourceFlags[i] == SourceFromProof {
			a = proof.Hashes[proofPos]
			proofPos++
		}
		if leafPos < len(leaves) {
			b = leaves[leafPos]
			leafPos++
		} else {
			b = hashes[hashPos]
			hashPos++
		}
		hashes[i] = ctx.HashInternal(a, b)
	}
	if totalHashes == 0 {
		return leaves[0], nil
	}
	return hashes[totalHashes-1], nil
}
