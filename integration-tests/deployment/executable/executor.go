package executable

import (
	"context"
	"encoding/binary"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/ccip-owner-contracts/gethwrappers"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment/errors"
)

type Executor struct {
	Proposal      *ExecutableMCMSProposal
	Tree          *MerkleTree
	RootMetadatas map[string]gethwrappers.ManyChainMultiSigRootMetadata
	Operations    map[string][]gethwrappers.ManyChainMultiSigOp
	Caller        *Caller
}

func NewProposalExecutor(proposal *ExecutableMCMSProposal, clients map[string]ContractDeployBackend) (*Executor, error) {
	txCounts := calculateTransactionCounts(proposal.Transactions)

	caller, err := NewCaller(mapMCMAddresses(proposal.ChainMetadata), clients)
	if err != nil {
		return nil, err
	}

	currentOpCounts, err := caller.GetCurrentOpCounts()
	if err != nil {
		return nil, err
	}

	rootMetadatas, err := buildRootMetadatas(proposal.ChainMetadata, txCounts, currentOpCounts, proposal.OverridePreviousRoot)
	if err != nil {
		return nil, err
	}

	ops, err := buildOperations(proposal.Transactions, rootMetadatas, txCounts)
	if err != nil {
		return nil, err
	}

	chainIdentifiers := sortedChainIdentifiers(proposal.ChainMetadata)

	tree, err := buildMerkleTree(chainIdentifiers, rootMetadatas, ops)

	return &Executor{
		Proposal:      proposal,
		Tree:          tree,
		RootMetadatas: rootMetadatas,
		Operations:    ops,
		Caller:        caller,
	}, err
}

func (e *Executor) SigningHash() (common.Hash, error) {
	// Convert validUntil to [32]byte
	var validUntilBytes [32]byte
	binary.BigEndian.PutUint32(validUntilBytes[28:], e.Proposal.ValidUntil) // Place the uint32 in the last 4 bytes

	hashToSign := crypto.Keccak256Hash(e.Tree.Root.Bytes(), validUntilBytes[:])
	return toEthSignedMessageHash(hashToSign), nil
}

func toEthSignedMessageHash(messageHash common.Hash) common.Hash {
	// Add the Ethereum signed message prefix
	prefix := []byte("\x19Ethereum Signed Message:\n32")
	data := append(prefix, messageHash.Bytes()...)

	// Hash the prefixed message
	return crypto.Keccak256Hash(data)
}

func (e *Executor) ValidateSignatures() error {
	hash, err := e.SigningHash()
	if err != nil {
		return err
	}

	recoveredSigners := make([]common.Address, len(e.Proposal.Signatures))
	for _, sig := range e.Proposal.Signatures {
		recoveredAddr, err := recoverAddressFromSignature(hash, sig.ToBytes())
		if err != nil {
			return err
		}
		recoveredSigners = append(recoveredSigners, recoveredAddr)
	}

	configs, err := e.Caller.GetConfigs()
	if err != nil {
		return err
	}

	// Validate that all signers are valid
	for chain, config := range configs {
		for _, signer := range recoveredSigners {
			found := false
			for _, mcmsSigner := range config.Signers {
				if mcmsSigner.Addr == signer {
					found = true
					break
				}
			}

			if !found {
				return &errors.ErrInvalidSignature{
					ChainIdentifier:  chain,
					RecoveredAddress: signer,
				}
			}
		}
	}

	return nil
}

func (e *Executor) Execute() error {
	proofs, err := e.Tree.GetProofs()
	if err != nil {
		return err
	}

	setRootTxs := make(map[string]*types.Transaction, 0)
	for chain, metadata := range e.RootMetadatas {
		encodedMetadata, err := metadataEncoder(metadata)
		if err != nil {
			return err
		}

		tx, err := e.Caller.Callers[chain].SetRoot(
			&bind.TransactOpts{},
			[32]byte(e.Tree.Root.Bytes()),
			e.Proposal.ValidUntil,
			metadata,
			mapHashes(proofs[encodedMetadata]),
			mapSignatures(e.Proposal.Signatures),
		)
		setRootTxs[chain] = tx

		if err != nil {
			return err
		}
	}

	// Wait for all SetRoot transactions to be mined
	for chain, tx := range setRootTxs {
		_, err := bind.WaitMined(context.TODO(), e.Caller.Clients[chain], tx)
		if err != nil {
			return err
		}
	}

	// TODO: Implement execute as well
	return nil
}
