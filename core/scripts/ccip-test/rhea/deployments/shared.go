package deployments

import (
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
)

const (
	BATCH_GAS_LIMIT              = 5_000_000
	FEE_UPDATE_HEARTBEAT         = 24 * time.Hour
	FEE_UPDATE_DEVIATION_PPB     = 5e7
	RELATIVE_BOOST_PER_WAIT_HOUR = 0.7
	INFLIGHT_CACHE_EXPIRY        = 3 * time.Minute
	ROOT_SNOOZE_TIME             = 10 * time.Minute
)

func getBlockConfirmations(chain rhea.Chain) uint32 {
	var blockConfirmationPerChain = map[rhea.Chain]uint32{
		rhea.Goerli:         4,
		rhea.Sepolia:        4,
		rhea.OptimismGoerli: 4,
		rhea.AvaxFuji:       1,
	}

	if val, ok := blockConfirmationPerChain[chain]; ok {
		return val
	}
	panic(fmt.Sprintf("Block confirmation for %s not found", chain))
}

func getMaxGasPrice(chain rhea.Chain) uint64 {
	var maxGasPricePerChain = map[rhea.Chain]uint64{
		rhea.Goerli:         200e9,
		rhea.Sepolia:        200e9,
		rhea.OptimismGoerli: 200e9,
		rhea.AvaxFuji:       200e9,
	}

	if val, ok := maxGasPricePerChain[chain]; ok {
		return val
	}
	panic(fmt.Sprintf("Max gas price for %s not found", chain))
}
