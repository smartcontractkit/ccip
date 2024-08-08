package deployment

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// TODO: To move to ccip-owner-contracts
type ManyChainMultiSigOp struct {
	ChainId  *big.Int
	MultiSig common.Address
	Nonce    *big.Int
	To       common.Address
	Value    *big.Int
	Data     []byte
}

type Proposal struct {
	// keccak256(abi.encode(root, validUntil)) is what is signed by MCMS
	// signers.
	ValidUntil uint32
	// Leaves are the items in the proposal.
	// Uses these to generate the root as well as display whats in the root.
	// These Ops may be destined for distinct chains.
	Ops []ManyChainMultiSigOp
}

func (p Proposal) String() string {
	// TODO
	return ""
}

// Services as input to CI/Async tasks
type MigrationOutput struct {
	JobSpecs    map[string][]string
	Proposals   []Proposal
	AddressBook AddressBook
}
