package ccip

var (
	// NoopMetricsCollector is a no-op implementation of PluginMetricsCollector
	NoopMetricsCollector PluginMetricsCollector = noop{}
)

type noop struct{}

func (d noop) NumberOfMessagesProcessed(ocrPhase, int) {
}

func (d noop) NumberOfMessagesBasedOnInterval(ocrPhase, uint64, uint64) {
}

func (d noop) UnexpiredCommitRoots(int) {
}

func (d noop) SequenceNumber(uint64) {
}
