package deployments

import (
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
)

const (
	BATCH_GAS_LIMIT                     = 5_000_000
	FEE_UPDATE_HEARTBEAT                = 24 * time.Hour
	FEE_UPDATE_DEVIATION_PPB            = 10e7 // 10%
	FEE_UPDATE_DEVIATION_PPB_FAST_CHAIN = 20e7 // 20%
	// This boosts the fee paid every 10x every 15s, since fees artificially low (0.1 on testnet) and
	// we have source finality artificially low. On fast chains transactions become available for execution
	// within 30s - 1min, and we want to avoid waiting for a full root snooze so we make sure they
	// are boosted back to cost immediately.
	RELATIVE_BOOST_PER_WAIT_HOUR = 2400
	INFLIGHT_CACHE_EXPIRY        = 3 * time.Minute
	ROOT_SNOOZE_TIME             = 5 * time.Minute
)

func getBlockConfirmations(chain rhea.Chain) uint32 {
	// NOTE most of these is still way artificially low but we aim for quick iteration on testnet.
	// Optimism, polygon and arbitrum in particular are known for decent sized reorgs so we set
	// those higher than others.
	var blockConfirmationPerChain = map[rhea.Chain]uint32{
		rhea.Goerli:         4,
		rhea.Sepolia:        4,
		rhea.OptimismGoerli: 10,
		rhea.AvaxFuji:       2, // Should be 1 theoretically, air on the side of caution.
		rhea.PolygonMumbai:  10,
		rhea.ArbitrumGoerli: 10,
		rhea.Quorum:         4,
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
		rhea.PolygonMumbai:  200e9,
		rhea.ArbitrumGoerli: 200e9,
		rhea.Quorum:         200e9,
	}

	if val, ok := maxGasPricePerChain[chain]; ok {
		return val
	}
	panic(fmt.Sprintf("Max gas price for %s not found", chain))
}
