package common

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
)

// HashedCapabilityID returns the hashed capability id in a manner equivalent to the capability registry.
func HashedCapabilityID(capabilityLabelledName, capabilityVersion string) (r [32]byte, err error) {
	tabi := `[{"type": "string"}, {"type": "string"}]`
	abiEncoded, err := utils.ABIEncode(tabi, capabilityLabelledName, capabilityVersion)
	if err != nil {
		return r, fmt.Errorf("failed to ABI encode capability version and labelled name: %w", err)
	}

	h := crypto.Keccak256(abiEncoded)
	copy(r[:], h)
	return r, nil
}

// MustHashedCapabilityID is a helper function that panics if the HashedCapabilityID function returns an error.
// should only use in tests.
func MustHashedCapabilityID(capabilityLabelledName, capabilityVersion string) [32]byte {
	r, err := HashedCapabilityID(capabilityLabelledName, capabilityVersion)
	if err != nil {
		panic(err)
	}
	return r
}