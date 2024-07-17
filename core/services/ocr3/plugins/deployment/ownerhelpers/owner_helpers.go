package ownerhelpers

import (
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/deployment/ownerhelpers/manychainmultisig"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/deployment/ownerhelpers/rbactimelock"
	"math/big"
)

// A timelock batch for a given chain.
type Batch struct {
	Calls []rbactimelock.RBACTimelockCall
	Delay *big.Int
	// TODO: salt/pred
}

type SetRootArgs struct {
	// keccak256(abi.encode(root, validUntil)) is what is signed by MCMS
	// signers.
	Root       [32]byte
	ValidUntil uint32
	// To be filled in by signers.
	Signatures manychainmultisig.ManyChainMultiSigSignature
	// Other calldata
	MetadataProof [][32]byte // Merkle proof of metadata being in the root
	RootMetadata  manychainmultisig.ManyChainMultiSigRootMetadata
}

// This thing should be signed.
func (s SetRootArgs) SigningPayload() ([]byte, error) {
	return utils.ABIEncode(`[{"type":"bytes32"},{"type":"uint32"}]`, s.Root, s.ValidUntil)
}

func GenerateSetRootArgs(batches map[uint64]Batch) (SetRootArgs, error) {
	// One of the leaves stores root metadata.
	return SetRootArgs{}, nil
}
