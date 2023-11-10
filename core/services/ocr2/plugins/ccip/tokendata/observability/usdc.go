package observability

import (
	"context"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/tokendata/usdc"
)

var (
	latencyBuckets = []float64{
		float64(10 * time.Millisecond),
		float64(25 * time.Millisecond),
		float64(50 * time.Millisecond),
		float64(75 * time.Millisecond),
		float64(100 * time.Millisecond),
		float64(250 * time.Millisecond),
		float64(500 * time.Millisecond),
		float64(750 * time.Millisecond),
		float64(1 * time.Second),
		float64(2 * time.Second),
	}
	labels        = []string{"plugin", "function", "success"}
	usdcHistogram = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "ccip_usdc_reader_request_total",
		Help:    "Latency of calls to the USDC reader",
		Buckets: latencyBuckets,
	}, labels)
)

type metricDetails struct {
	histogram  *prometheus.HistogramVec
	pluginName string
}

type ObservedUSDCTokenDataReader struct {
	usdc.TokenDataReader
	metric metricDetails
}

func NewObservedUSDCTokenDataReader(origin usdc.TokenDataReader, pluginName string) *ObservedUSDCTokenDataReader {
	return &ObservedUSDCTokenDataReader{
		TokenDataReader: *usdc.NewUSDCTokenDataReaderWithHttpClient(origin, NewObservedIHttpClient(&usdc.HttpClient{}, pluginName)),
		metric: metricDetails{
			histogram:  usdcHistogram,
			pluginName: pluginName,
		},
	}
}

func (o *ObservedUSDCTokenDataReader) ReadTokenData(ctx context.Context, msg internal.EVM2EVMOnRampCCIPSendRequestedWithMeta) ([]byte, error) {
	return withObservedContract(o.metric, "ReadTokenData", func() ([]byte, error) {
		return o.TokenDataReader.ReadTokenData(ctx, msg)
	})
}

func withObservedContract[T any](metric metricDetails, function string, contract func() (T, error)) (T, error) {
	contractExecutionStarted := time.Now()
	value, err := contract()
	metric.histogram.
		WithLabelValues(
			metric.pluginName,
			function,
			strconv.FormatBool(err == nil),
		).
		Observe(float64(time.Since(contractExecutionStarted)))
	return value, err
}
