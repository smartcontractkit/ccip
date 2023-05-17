package deployments

import (
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
)

const (
	BATCH_GAS_LIMIT              = 5_000_000
	FEE_UPDATE_HEARTBEAT         = 24 * time.Hour
	FEE_UPDATE_DEVIATION_PPB     = 10e7
	RELATIVE_BOOST_PER_WAIT_HOUR = 600 // 10x every minute, fees artificially low (0.1 on testnet)
	INFLIGHT_CACHE_EXPIRY        = 3 * time.Minute
	ROOT_SNOOZE_TIME             = 7 * time.Minute
)

func getFinalityDepth(chain rhea.Chain) uint32 {
	// Note for testnets these are intentionally low to provide quicker testing.
	var finalityDepthPerChain = map[rhea.Chain]uint32{
		rhea.Goerli:         4,
		rhea.Sepolia:        4,
		rhea.OptimismGoerli: 4,
		rhea.AvaxFuji:       1,
		rhea.PolygonMumbai:  4,
		rhea.ArbitrumGoerli: 1,
	}

	if val, ok := finalityDepthPerChain[chain]; ok {
		return val
	}
	panic(fmt.Sprintf("Finality depth for %s not found", chain))
}

func getOptimisticConfirmations(chain rhea.Chain) uint32 {
	var optimisticConfirmations = map[rhea.Chain]uint32{
		rhea.Goerli:         4,
		rhea.Sepolia:        4,
		rhea.OptimismGoerli: 4,
		rhea.AvaxFuji:       1,
		rhea.PolygonMumbai:  4,
		rhea.ArbitrumGoerli: 1,
	}

	if val, ok := optimisticConfirmations[chain]; ok {
		return val
	}
	panic(fmt.Sprintf("Optimistic confirmations for %s not found", chain))
}

func getMaxGasPrice(chain rhea.Chain) uint64 {
	var maxGasPricePerChain = map[rhea.Chain]uint64{
		rhea.Goerli:         200e9,
		rhea.Sepolia:        200e9,
		rhea.OptimismGoerli: 200e9,
		rhea.AvaxFuji:       200e9,
		rhea.PolygonMumbai:  200e9,
		rhea.ArbitrumGoerli: 200e9,
	}

	if val, ok := maxGasPricePerChain[chain]; ok {
		return val
	}
	panic(fmt.Sprintf("Max gas price for %s not found", chain))
}
