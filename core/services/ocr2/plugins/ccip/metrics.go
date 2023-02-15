package ccip

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type skipReason string

const (
	// reasonSnoozed describes when a report is skipped due to being snoozed.
	reasonSnoozed skipReason = "snoozed"

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
