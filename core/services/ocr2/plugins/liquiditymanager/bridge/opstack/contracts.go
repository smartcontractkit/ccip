package opstack

import (
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
)

var (
	// Optimism contract addresses: https://docs.optimism.io/chain/addresses
	AllContracts map[uint64]map[string]common.Address
)

func init() {
	AllContracts = map[uint64]map[string]common.Address{
		chainsel.ETHEREUM_TESTNET_SEPOLIA.Selector: {
			"L1StandardBridgeProxy": common.HexToAddress("0xFBb0621E0B23b5478B630BD55a5f21f67730B0F1"),
		},
		chainsel.ETHEREUM_TESTNET_SEPOLIA_OPTIMISM_1.Selector: {
			"L2StandardBridge": common.HexToAddress("0x4200000000000000000000000000000000000010"),
		},
	}
}
