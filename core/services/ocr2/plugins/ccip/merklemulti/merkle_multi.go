package merklemulti

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/core/utils"
)

// Add additional hash types e.g. [20]byte as needed here.
type Hash interface {
	[32]byte
}

type Ctx[H Hash] interface {
	HashLeaf(l []byte) H
	HashInternal(a, b H) H
	ZeroHash() H
}

type keccakCtx struct{}

func NewKeccakCtx() Ctx[[32]byte] {
	return keccakCtx{}
}

func (k keccakCtx) HashLeaf(l []byte) [32]byte {
	// Note this Keccak256 cannot error https://github.com/golang/crypto/blob/master/sha3/sha3.go#L126
	// if we start supporting hashing algos which do, we can change this API to include an error.
	return utils.Keccak256Fixed(append([]byte{0x00}, l...))
}

func (k keccakCtx) HashInternal(a, b [32]byte) [32]byte {
	if bytes.Compare(a[:], b[:]) < 0 {
		return utils.Keccak256Fixed(append([]byte{0x01}, append(a[:], b[:]...)...))
	}
	return utils.Keccak256Fixed(append([]byte{0x01}, append(b[:], a[:]...)...))
}

// We use empty bytes32 for zeroHash
// on the solidity side, this needs to match.
func (k keccakCtx) ZeroHash() [32]byte {
	var zeroes [32]byte
	return zeroes
}

type singleLayerProof[H Hash] struct {
	nextIndices []int
	subProof    []H
	sourceFlags []bool
}

type Proof[H Hash] struct {
	Hashes      []H    `json:"hashes"`
	SourceFlags []bool `json:"source_flags"`
}

func (p Proof[H]) countSourceFlags(b bool) (count int) {
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

func proveSingleLayer[H Hash](layer []H, indices []int) singleLayerProof[H] {
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
	var subProof []H
	for _, i := range authIndices {
		subProof = append(subProof, layer[i])
	}
	return singleLayerProof[H]{
		nextIndices: nextIndices,
		subProof:    subProof,
		sourceFlags: sourceFlags,
	}
}

type Tree[H Hash] struct {
	layers [][]H
	ctx    Ctx[H]
}

func NewTree[H Hash](ctx Ctx[H], leafHashes []H) *Tree[H] {
	var layer = make([]H, len(leafHashes))
	copy(layer, leafHashes)

	var layers = [][]H{layer}
	var curr int
	for len(layer) > 1 {
		paddedLayer, nextLayer := computeNextLayer(ctx, layer)
		layers[curr] = paddedLayer
		curr++
		layers = append(layers, nextLayer)
		layer = nextLayer
	}
	return &Tree[H]{
		layers: layers,
	}
}

// Revive appears confused with the generics "receiver name t should be consistent with previous receiver name p for invalid-type"
//revive:disable:receiver-naming
func (t *Tree[H]) String() string {
	b := strings.Builder{}
	for _, layer := range t.layers {
		b.WriteString(fmt.Sprintf("%v", layer))
	}
	return b.String()
}

func (t *Tree[H]) Root() H {
	return t.layers[len(t.layers)-1][0]
}

func (t *Tree[H]) Prove(indices []int) Proof[H] {
	var proof Proof[H]
	for _, layer := range t.layers[:len(t.layers)-1] {
		res := proveSingleLayer(layer, indices)
		indices = res.nextIndices
		proof.Hashes = append(proof.Hashes, res.subProof...)
		proof.SourceFlags = append(proof.SourceFlags, res.sourceFlags...)
	}
	return proof
}

func computeNextLayer[H Hash](ctx Ctx[H], layer []H) ([]H, []H) {
	if len(layer) == 1 {
		return layer, layer
	}
	if len(layer)%2 != 0 {
		layer = append(layer, ctx.ZeroHash())
	}
	var nextLayer []H
	for i := 0; i < len(layer); i += 2 {
		nextLayer = append(nextLayer, ctx.HashInternal(layer[i], layer[i+1]))
	}
	return layer, nextLayer
}

func VerifyComputeRoot[H Hash](ctx Ctx[H], leafHashes []H, proof Proof[H]) (H, error) {
	totalHashes := len(leafHashes) + len(proof.Hashes) - 1
	if totalHashes != len(proof.SourceFlags) {
		return ctx.ZeroHash(), errors.Errorf("hashes %d != sourceFlags %d", totalHashes, len(proof.SourceFlags))
	}
	sourceProofCount := proof.countSourceFlags(SourceFromProof)
	if sourceProofCount != len(proof.Hashes) {
		return ctx.ZeroHash(), errors.Errorf("proof source flags %d != proof hashes%d", sourceProofCount, len(proof.Hashes))
	}
	var hashes []H
	for i := 0; i < totalHashes; i++ {
		hashes = append(hashes, leafHashes[0])
	}
	var (
		leafPos  int
		hashPos  int
		proofPos int
	)
	for i := 0; i < totalHashes; i++ {
		var a, b H
		if proof.SourceFlags[i] == SourceFromHashes {
			if leafPos < len(leafHashes) {
				a = leafHashes[leafPos]
				leafPos++
			} else {
				a = hashes[hashPos]
				hashPos++
			}
		} else if proof.SourceFlags[i] == SourceFromProof {
			a = proof.Hashes[proofPos]
			proofPos++
		}
		if leafPos < len(leafHashes) {
			b = leafHashes[leafPos]
			leafPos++
		} else {
			b = hashes[hashPos]
			hashPos++
		}
		hashes[i] = ctx.HashInternal(a, b)
	}
	if totalHashes == 0 {
		return leafHashes[0], nil
	}
	return hashes[totalHashes-1], nil
}
