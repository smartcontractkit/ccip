package ccip

import (
	"math/big"
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
	execPluginReportsCount = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ccip_execution_observation_reports_count",
		Help: "Number of reports that are being processed by Execution Plugin during single observation",
	}, []string{"plugin", "source", "dest"})
	messagesProcessed = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ccip_number_of_messages_processed",
		Help: "Number of messages processed by the plugin during different OCR phases",
	}, []string{"plugin", "source", "dest", "ocr_phase"})
	sequenceNumberCounter = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ccip_sequence_number_counter",
		Help: "Sequence number of the last message processed by the plugin",
	}, []string{"plugin", "source", "dest"})
)

func incSkippedRequests(reason skipReason) {
	metricReportSkipped.WithLabelValues(string(reason)).Inc()
}

type ocrPhase string

const (
	Observation ocrPhase = "observation"
	Report      ocrPhase = "report"
)

type PluginMetricsCollector interface {
	NumberOfMessagesProcessed(phase ocrPhase, count int)
	NumberOfMessagesBasedOnInterval(phase ocrPhase, seqNrMin, seqNrMax uint64)
	UnexpiredCommitRoots(count int)
	SequenceNumber(seqNr uint64)
}

type pluginMetricsCollector struct {
	pluginName   string
	source, dest string
}

func NewPluginMetricsCollector(pluginLabel string, sourceChainId, destChainId *big.Int) *pluginMetricsCollector {
	return &pluginMetricsCollector{
		pluginName: pluginLabel,
		source:     sourceChainId.String(),
		dest:       destChainId.String(),
	}
}

func (p *pluginMetricsCollector) NumberOfMessagesProcessed(phase ocrPhase, count int) {
	messagesProcessed.
		WithLabelValues(p.pluginName, p.source, p.dest, string(phase)).
		Set(float64(count))
}

func (p *pluginMetricsCollector) NumberOfMessagesBasedOnInterval(phase ocrPhase, seqNrMin, seqNrMax uint64) {
	messagesProcessed.
		WithLabelValues(p.pluginName, p.source, p.dest, string(phase)).
		Set(float64(seqNrMax - seqNrMin + 1))
}

func (p *pluginMetricsCollector) UnexpiredCommitRoots(count int) {
	execPluginReportsCount.
		WithLabelValues(p.pluginName, p.source, p.dest).
		Set(float64(count))
}

func (p *pluginMetricsCollector) SequenceNumber(seqNr uint64) {
	// Don't publish price reports
	if seqNr == 0 {
		return
	}

	sequenceNumberCounter.
		WithLabelValues(p.pluginName, p.source, p.dest).
		Set(float64(seqNr))
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
	case 1111:
		return "wemix-mainnet"
	case 1112:
		return "wemix-testnet"
	case 255:
		return "ethereum-mainnet-kroma-1"
	case 2358:
		return "ethereum-testnet-sepolia-kroma-1"
	case 4002:
		return "fantom-testnet"
	case 8453:
		return "ethereum-mainnet-base-1"
	case 84531:
		return "ethereum-testnet-goerli-base-1"
	case 84532:
		return "ethereum-testnet-sepolia-base-1"
	case 42161:
		return "ethereum-mainnet-arbitrum-1"
	case 421613:
		return "ethereum-testnet-goerli-arbitrum-1"
	case 421614:
		return "ethereum-testnet-sepolia-arbitrum-1"
	case 43113:
		return "avalanche-testnet-fuji"
	case 43114:
		return "avalanche-mainnet"
	case 76578:
		return "avalanche-testnet-anz-subnet"
	case 80001:
		return "polygon-testnet-mumbai"
	case 11155111:
		return "ethereum-testnet-sepolia"
	case 11155420:
		return "ethereum-testnet-sepolia-optimism-1"
	default: // Unknown chain, return chainID as string
		return strconv.FormatInt(chainID, 10)
	}
}
