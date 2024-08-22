package executable

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type MerkleTree struct {
	// Root hash of the merkle tree
	Root common.Hash

	// Layers of the merkle tree, starting from the leaves
	Layers [][]common.Hash
}

func NewMerkleTree(leaves []common.Hash) *MerkleTree {
	tree := &MerkleTree{
		Layers: make([][]common.Hash, 0),
	}

	currHashes := leaves
	for len(currHashes) > 1 {
		// If the number of hashes is odd, duplicate the last hash
		if len(currHashes)%2 != 0 {
			currHashes = append(currHashes, currHashes[len(currHashes)-1])
		}

		// Append the current layer to the tree
		tree.Layers = append(tree.Layers, currHashes)

		// Calculate the parent hashes
		tempHashes := make([]common.Hash, len(currHashes)/2)
		for i := 0; i < len(currHashes); i += 2 {
			// Sort the pair of hashes before hashing
			var parentHash common.Hash
			if currHashes[i].String() > currHashes[i+1].String() {
				parentHash = crypto.Keccak256Hash(currHashes[i+1].Bytes(), currHashes[i].Bytes())
			} else {
				parentHash = crypto.Keccak256Hash(currHashes[i].Bytes(), currHashes[i+1].Bytes())
			}
			tempHashes[i/2] = parentHash
		}

		// Set the current hashes to the parent hashes
		currHashes = tempHashes
	}

	// Append the root hash to the tree
	tree.Root = currHashes[0]
	return tree
}

func (t *MerkleTree) GetProof(hash common.Hash) ([]common.Hash, error) {
	proof := make([]common.Hash, 0)

	targetHash := hash
	for i := 0; i < len(t.Layers); i++ {
		found := false
		for j, h := range t.Layers[i] {
			if h == targetHash {
				// Get the sibling hash
				siblingIdx := j ^ 1
				siblingHash := t.Layers[i][siblingIdx]
				proof = append(proof, siblingHash)

				// Get next target hash by sorting the pair of hashes and hashing them
				if targetHash.String() > siblingHash.String() {
					targetHash = crypto.Keccak256Hash(siblingHash.Bytes(), targetHash.Bytes())
				} else {
					targetHash = crypto.Keccak256Hash(targetHash.Bytes(), siblingHash.Bytes())
				}

				// Move to the next layer
				found = true
				break
			}
		}

		if !found {
			// If the hash is not found in the current layer, it is not in the tree
			// THIS SHOULD NEVER HAPPEN
			return nil, &ErrMerkleTreeNodeNotFound{
				TargetHash: targetHash,
			}
		}
	}

	return proof, nil
}

func (t *MerkleTree) GetProofs() (map[common.Hash][]common.Hash, error) {
	proofs := make(map[common.Hash][]common.Hash)

	for _, leaf := range t.Layers[0] {
		proof, err := t.GetProof(leaf)
		if err != nil {
			// THIS SHOULD NEVER HAPPEN
			return nil, err
		}

		proofs[leaf] = proof
	}

	return proofs, nil
}

type ErrMerkleTreeNodeNotFound struct {
	TargetHash common.Hash
}

func (e *ErrMerkleTreeNodeNotFound) Error() string {
	return "merkle tree does not contain hash: " + e.TargetHash.String()
}
