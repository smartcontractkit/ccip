package ccip

import (
	"math/big"

	"golang.org/x/crypto/sha3"
)

type MerkleProof struct {
	index int
	path  [][32]byte
}

func NewMerkleProof(index int, path [][32]byte) MerkleProof {
	return MerkleProof{
		index: index,
		path:  path,
	}
}

func (mp MerkleProof) PathForExecute() [][32]byte {
	return mp.path
}

func (mp MerkleProof) Index() *big.Int {
	return big.NewInt(int64(mp.index))
}

func GenerateMerkleProof(treeHeight int, leaves [][]byte, index int) ([32]byte, MerkleProof) {
	zhs := ComputeZeroHashes(treeHeight)
	var level [][32]byte
	for _, leaf := range leaves {
		level = append(level, HashLeaf(leaf))
	}
	levelIndex := index
	var path [][32]byte
	// Go level by level up the tree starting from the bottom.
	// Record the path of sibling nodes for the Index node required
	// to get the top.
	for height := 0; height < treeHeight-1; height++ {
		// If we have an odd number of level elements
		if len(level)%2 == 1 {
			level = append(level, zhs[height])
		}
		pathIndex := levelIndex + 1
		if levelIndex%2 == 1 {
			pathIndex = levelIndex - 1
		}
		path = append(path, level[pathIndex])
		// Floor division here
		// E.g. [0, 1, 2, 3]
		// maps to [0, 1] on the next level.
		// So 0,1 -> 0 and 2,3 -> 1
		levelIndex /= 2
		// Compute the next level by hashing each pair of nodes.
		// (we know there is an even number of them)
		var newLevel [][32]byte
		for i := 0; i < len(level)-1; i += 2 {
			newLevel = append(newLevel, HashInternal(level[i], level[i+1]))
		}
		level = newLevel
	}
	if len(level) != 1 {
		panic("invalid")
	}
	return level[0], MerkleProof{path: path, index: index}
}

func GenerateMerkleRoot(leaf []byte, proof MerkleProof) [32]byte {
	// Make a deep copy of the path
	var path [][32]byte
	for _, p := range proof.path {
		var pc [32]byte
		copy(pc[:], p[:])
		path = append(path, pc)
	}
	index := proof.index
	h := HashLeaf(leaf)
	var l, r [32]byte
	for {
		if len(path) == 0 {
			break
		}
		if index%2 == 0 {
			// if Index is even then our Index is on the left
			// and the Proof element is on the right
			l = h
			r = path[0]
		} else {
			// if Index is odd then our Index is on the right
			// and the Proof element is on the left
			l = path[0]
			r = h
		}
		path = path[1:] // done with that Proof element
		h = HashInternal(l, r)
		index >>= 1
	}
	return h
}

func HashInternal(l, r [32]byte) [32]byte {
	hash := sha3.NewLegacyKeccak256()
	// Ignore errors
	hash.Write([]byte{0x01})
	hash.Write(l[:])
	hash.Write(r[:])
	var res [32]byte
	copy(res[:], hash.Sum(nil))
	return res
}

func HashLeaf(b []byte) [32]byte {
	hash := sha3.NewLegacyKeccak256()
	// Ignore errors
	hash.Write([]byte{0x00})
	hash.Write(b)
	var r [32]byte
	copy(r[:], hash.Sum(nil))
	return r
}

func ComputeZeroHashes(height int) [][32]byte {
	// Pre-compute all-zero trees for each depth
	// i.e. [0x00, H(0x00), H(H(0x00)||H(0x00)), ...]
	var zeroHashes = make([][32]byte, height)
	for i := 0; i < height-1; i++ {
		if i == 0 {
			var zh [32]byte
			zeroHashes[i] = zh
		}
		zeroHashes[i+1] = HashInternal(zeroHashes[i], zeroHashes[i])
	}
	return zeroHashes
}
