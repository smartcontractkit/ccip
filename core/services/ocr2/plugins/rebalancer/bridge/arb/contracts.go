package arb

import (
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
)

type L1Contracts struct {
	RollupAddress        common.Address
	GatewayRouterAddress common.Address
	InboxAddress         common.Address
}

type Contracts struct {
	L1 L1Contracts
}

var (
	AllContracts map[uint64]Contracts
)

func init() {
	AllContracts = map[uint64]Contracts{
		// See https://docs.arbitrum.io/for-devs/useful-addresses
		chainsel.ETHEREUM_TESTNET_SEPOLIA_ARBITRUM_1.Selector: {
			L1: L1Contracts{
				RollupAddress:        common.HexToAddress("0xd80810638dbDF9081b72C1B33c65375e807281C8"),
				GatewayRouterAddress: common.HexToAddress("0xcE18836b233C83325Cc8848CA4487e94C6288264"),
				InboxAddress:         common.HexToAddress("0xaAe29B0366299461418F5324a79Afc425BE5ae21"),
			},
		},
		chainsel.ETHEREUM_MAINNET_ARBITRUM_1.Selector: {
			L1: L1Contracts{
				RollupAddress:        common.HexToAddress("0x5eF0D09d1E6204141B4d37530808eD19f60FBa35"),
				GatewayRouterAddress: common.HexToAddress("0x72Ce9c846789fdB6fC1f34aC4AD25Dd9ef7031ef"),
				InboxAddress:         common.HexToAddress("0x4Dbd4fc535Ac27206064B68FfCf827b0A60BAB3f"),
			},
		},
	}
}
