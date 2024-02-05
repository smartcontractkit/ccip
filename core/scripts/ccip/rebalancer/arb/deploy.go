package arb

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip/rebalancer/bridgeutil"
)

var (
	// Arbitrum Contracts
	// See https://docs.arbitrum.io/for-devs/useful-addresses
	ArbitrumContracts map[uint64]map[string]common.Address
)

func init() {
	ArbitrumContracts = map[uint64]map[string]common.Address{
		bridgeutil.SepoliaChainID: {
			"L1GatewayRouter": common.HexToAddress("0xcE18836b233C83325Cc8848CA4487e94C6288264"),
			"L1Outbox":        common.HexToAddress("0x65f07C7D521164a4d5DaC6eB8Fac8DA067A3B78F"),
			"Rollup":          common.HexToAddress("0xd80810638dbDF9081b72C1B33c65375e807281C8"),
			"WETH":            common.HexToAddress("0x7b79995e5f793A07Bc00c21412e50Ecae098E7f9"),
		},
		bridgeutil.ArbitrumSepoliaChainID: {
			"L2GatewayRouter": common.HexToAddress("0x9fDD1C4E4AA24EEc1d913FABea925594a20d43C7"),
			"NodeInterface":   common.HexToAddress("0x00000000000000000000000000000000000000C8"),
			"WETH":            common.HexToAddress("0x980B62Da83eFf3D4576C647993b0c1D7faf17c73"),
		},
	}
}
