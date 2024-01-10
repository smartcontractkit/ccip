package ccip

import (
	"math/big"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type skipReason string

const (
	// ReasonNotBlessed describes when a report is skipped due to not being blessed.
	ReasonNotBlessed skipReason = "not blessed"

	// ReasonAllExecuted describes when a report is skipped due to messages being all executed.
	ReasonAllExecuted skipReason = "all executed"
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
	}, []string{"plugin", "source", "dest", "ocrPhase"})
	sequenceNumberCounter = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ccip_sequence_number_counter",
		Help: "Sequence number of the last message processed by the plugin",
	}, []string{"plugin", "source", "dest"})
)

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

func IncSkippedRequests(reason skipReason) {
	metricReportSkipped.WithLabelValues(string(reason)).Inc()
}
