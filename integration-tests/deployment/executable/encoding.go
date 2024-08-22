package executable

import (
	"math/big"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/ccip-owner-contracts/gethwrappers"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment/errors"
)

func calculateTransactionCounts(transactions []ChainOperation) map[string]uint64 {
	txCounts := make(map[string]uint64)
	for _, tx := range transactions {
		txCounts[tx.ChainIdentifier]++
	}
	return txCounts
}

func buildRootMetadatas(
	chainMetadata map[string]ExecutableMCMSChainMetadata,
	txCounts map[string]uint64,
	currentOpCounts map[string]big.Int,
	overridePreviousRoot bool,
) (map[string]gethwrappers.ManyChainMultiSigRootMetadata, error) {
	rootMetadatas := make(map[string]gethwrappers.ManyChainMultiSigRootMetadata)

	for chainID, metadata := range chainMetadata {
		stringChainId, err := big.NewInt(0).SetString(chainID, 10)
		if !err {
			return nil, &errors.ErrInvalidChainID{
				ReceivedChainID: chainID,
			}
		}

		currentNonce, ok := currentOpCounts[chainID]
		if !ok {
			return nil, &errors.ErrMissingChainDetails{
				ChainIdentifier: chainID,
				Parameter:       "current op count",
			}
		}

		currentTxCount, ok := txCounts[chainID]
		if !ok {
			return nil, &errors.ErrMissingChainDetails{
				ChainIdentifier: chainID,
				Parameter:       "transaction count",
			}
		}

		rootMetadatas[chainID] = gethwrappers.ManyChainMultiSigRootMetadata{
			ChainId:              stringChainId,
			MultiSig:             metadata.MCMAddress,
			PreOpCount:           big.NewInt(currentNonce.Int64() + int64(metadata.NonceOffset)),                         // TODO: handle overflow
			PostOpCount:          big.NewInt(currentNonce.Int64() + int64(metadata.NonceOffset) + int64(currentTxCount)), // TODO: handle overflow
			OverridePreviousRoot: overridePreviousRoot,
		}
	}
	return rootMetadatas, nil
}

func buildOperations(
	transactions []ChainOperation,
	rootMetadatas map[string]gethwrappers.ManyChainMultiSigRootMetadata,
	txCounts map[string]uint64,
) (map[string][]gethwrappers.ManyChainMultiSigOp, error) {
	ops := make(map[string][]gethwrappers.ManyChainMultiSigOp)
	chainIdx := make(map[string]uint32, len(rootMetadatas))

	for _, tx := range transactions {
		rootMetadata := rootMetadatas[tx.ChainIdentifier]
		if _, ok := ops[tx.ChainIdentifier]; !ok {
			ops[tx.ChainIdentifier] = make([]gethwrappers.ManyChainMultiSigOp, txCounts[tx.ChainIdentifier])
			chainIdx[tx.ChainIdentifier] = 0
		}

		op := gethwrappers.ManyChainMultiSigOp{
			ChainId:  rootMetadata.ChainId,
			MultiSig: rootMetadata.MultiSig,
			Nonce:    big.NewInt(rootMetadata.PreOpCount.Int64() + int64(chainIdx[tx.ChainIdentifier])),
			To:       tx.To,
			Data:     common.FromHex(tx.Data),
			Value:    big.NewInt(int64(tx.Value)),
		}

		ops[tx.ChainIdentifier][chainIdx[tx.ChainIdentifier]] = op
		chainIdx[tx.ChainIdentifier]++
	}

	return ops, nil
}

func sortedChainIdentifiers(chainMetadata map[string]ExecutableMCMSChainMetadata) []string {
	chainIdentifiers := make([]string, 0, len(chainMetadata))
	for chainID := range chainMetadata {
		chainIdentifiers = append(chainIdentifiers, chainID)
	}
	sort.Strings(chainIdentifiers)
	return chainIdentifiers
}

func buildMerkleTree(
	chainIdentifiers []string,
	rootMetadatas map[string]gethwrappers.ManyChainMultiSigRootMetadata,
	ops map[string][]gethwrappers.ManyChainMultiSigOp,
) (*MerkleTree, error) {
	hashLeaves := make([]common.Hash, 0)

	for _, chainID := range chainIdentifiers {
		encodedRootMetadata, err := metadataEncoder(rootMetadatas[chainID])
		if err != nil {
			return nil, err
		}
		hashLeaves = append(hashLeaves, encodedRootMetadata)

		for _, op := range ops[chainID] {
			encodedOp, err := txEncoder(op)
			if err != nil {
				return nil, err
			}
			hashLeaves = append(hashLeaves, encodedOp)
		}
	}

	// sort the hashes and sort the pairs
	sort.Slice(hashLeaves, func(i, j int) bool {
		return hashLeaves[i].String() < hashLeaves[j].String()
	})

	return NewMerkleTree(hashLeaves), nil
}

func metadataEncoder(rootMetadata gethwrappers.ManyChainMultiSigRootMetadata) (common.Hash, error) {
	// ABI definition for the metadata tuple
	abiJSON := `[{
		"name": "metadata",
		"type": "function",
		"inputs": [
			{"name": "domainSeparator", "type": "bytes32"},
			{"name": "metadata", "type": "tuple(uint256 chainId, address multisig, uint48 preOpCount, uint48 postOpCount, bool overridePreviousRoot)"}
		]
	}]`

	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return common.Hash{}, err
	}

	// Pack the data using abi.Pack
	packed, err := parsedABI.Pack("metadata",
		MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_METADATA,
		rootMetadata,
	)
	if err != nil {
		return common.Hash{}, err
	}

	return crypto.Keccak256Hash(packed), nil
}

func txEncoder(op gethwrappers.ManyChainMultiSigOp) (common.Hash, error) {
	// ABI definition for the transaction tuple
	abiJSON := `[{
		"name": "tx",
		"type": "function",
		"inputs": [
			{"name": "domainSeparator", "type": "bytes32"},
			{"name": "tx", "type": "tuple(uint256 chainId, address multisig, uint64 nonce, address to, uint256 value, bytes data)"}
		]
	}]`

	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return common.Hash{}, err
	}

	// Pack the data using abi.Pack
	packed, err := parsedABI.Pack("tx",
		MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_OP,
		op,
	)
	if err != nil {
		return common.Hash{}, err
	}

	return crypto.Keccak256Hash(packed), nil
}
