package leafer

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp_1_0_0"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp_1_2_0"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/v1_0_0"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/v1_2_0"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/v1_5_0"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/pkg/hashlib"
)

// LeafHasher converts a CCIPSendRequested event into something that can be hashed and hashes it.
type LeafHasher interface {
	HashLeaf(log types.Log) ([32]byte, error)
}

// HasherVersion maps to the contract versions.
type HasherVersion string

const (
	V1_0_0 HasherVersion = "v1_0_0"
	V1_2_0 HasherVersion = "v1_2_0"
	V1_5_0 HasherVersion = "v1_5_0"
)

// MakeLeafHasher is a factory function to construct the onramp implementing the HashLeaf function for a given version.
func MakeLeafHasher(ver HasherVersion, cl bind.ContractBackend, sourceChainSelector uint64, destChainSelector uint64, onRampId common.Address, ctx hashlib.Ctx[[32]byte]) (LeafHasher, error) {
	switch ver {
	case V1_0_0:
		or, err := evm_2_evm_onramp_1_0_0.NewEVM2EVMOnRamp(onRampId, cl)
		if err != nil {
			return nil, err
		}
		h := v1_0_0.NewLeafHasher(sourceChainSelector, destChainSelector, onRampId, ctx, or)
		return h, nil
	case V1_2_0:
		or, err := evm_2_evm_onramp_1_2_0.NewEVM2EVMOnRamp(onRampId, cl)
		if err != nil {
			return nil, err
		}
		h := v1_2_0.NewLeafHasher(sourceChainSelector, destChainSelector, onRampId, ctx, or)
		return h, nil
	case V1_5_0:
		or, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(onRampId, cl)
		if err != nil {
			return nil, err
		}
		h := v1_5_0.NewLeafHasher(sourceChainSelector, destChainSelector, onRampId, ctx, or)
		return h, nil
	default:
		return nil, fmt.Errorf("unknown version %q", ver)
	}
}
