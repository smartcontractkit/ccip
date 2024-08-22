package executable

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/ccip-owner-contracts/gethwrappers"
)

var MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_OP = crypto.Keccak256Hash([]byte("MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_OP"))
var MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_METADATA = crypto.Keccak256Hash([]byte("MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_METADATA"))

type ContractDeployBackend interface {
	bind.ContractBackend
	bind.DeployBackend
}

type Operation struct {
	To    common.Address
	Data  string
	Value uint64
}

type ChainOperation struct {
	ChainIdentifier string
	Operation
}

func recoverAddressFromSignature(hash common.Hash, sig []byte) (common.Address, error) {
	// The signature should be 65 bytes, and the last byte is the recovery id (v).
	if len(sig) != 65 {
		return common.Address{}, fmt.Errorf("invalid signature length")
	}

	// Adjust the recovery id (v) if needed. Ethereum signatures expect 27 or 28.
	// But `crypto.SigToPub` expects 0 or 1.
	sig[64] -= 27

	// Recover the public key from the signature and the message hash
	pubKey, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		return common.Address{}, err
	}

	// Derive the Ethereum address from the public key
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return recoveredAddr, nil
}

func mapMCMAddresses(metadatas map[string]ExecutableMCMSChainMetadata) map[string]common.Address {
	m := make(map[string]common.Address)
	for k, v := range metadatas {
		m[k] = v.MCMAddress
	}
	return m
}

func mapSignatures(signatures []Signature) []gethwrappers.ManyChainMultiSigSignature {
	sigs := make([]gethwrappers.ManyChainMultiSigSignature, len(signatures))
	for i, sig := range signatures {
		sigs[i] = sig.ToGethSignature()
	}
	return sigs
}

func mapHashes(hashes []common.Hash) [][32]byte {
	m := make([][32]byte, len(hashes))
	for i, h := range hashes {
		m[i] = hashToBytes32(h)
	}
	return m
}

func hashToBytes32(hash common.Hash) [32]byte {
	return [32]byte(hash)
}
