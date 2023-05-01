package ccip

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type skipReason string

const (
	// reasonNotBlessed describes when a report is skipped due to not being blessed.
	reasonNotBlessed skipReason = "not blessed"

	// reasonAllExecuted describes when a report is skipped due to messages being all executed.
	reasonAllExecuted skipReason = "all executed"
)

var (
	metricReportSkipped = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "ccip_unexpired_report_skipped",
		Help: "Times report is skipped for the possible reasons",
	}, []string{"reason"})
)

func incSkippedRequests(reason skipReason) {
	metricReportSkipped.WithLabelValues(string(reason)).Inc()
}

// ChainName returns the name of the EVM network based on its chainID
func ChainName(chainID int64) string {
	switch chainID {
	case 1:
		return "ethereum-mainnet"
	case 4:
		return "ethereum-testnet-rinkeby"
	case 5:
		return "ethereum-testnet-goerli"
	case 10:
		return "ethereum-mainnet-optimism-1"
	case 42:
		return "ethereum-testnet-kovan"
	case 56:
		return "binance_smart_chain-mainnet"
	case 97:
		return "binance_smart_chain-testnet"
	case 137:
		return "polygon-mainnet"
	case 420:
		return "ethereum-testnet-goerli-optimism-1"
	case 4002:
		return "fantom-testnet"
	case 42161:
		return "ethereum-mainnet-arbitrum-1"
	case 43113:
		return "avalanche-testnet-fuji"
	case 43114:
		return "avalanche-mainnet"
	case 80001:
		return "polygon-testnet-mumbai"
	case 11155111:
		return "ethereum-testnet-sepolia"
	default: // Unknown chain, return chainID as string
		return strconv.FormatInt(chainID, 10)
	}
}
