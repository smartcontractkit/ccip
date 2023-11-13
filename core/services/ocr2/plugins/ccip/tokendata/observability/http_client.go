package observability

import (
	"context"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

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
	usdcClientHistogram = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "ccip_usdc_client_request_total",
		Help:    "Latency of calls to the USDC client",
		Buckets: latencyBuckets,
	}, []string{"plugin", "status", "success"})
)

type metricDetails struct {
	histogram  *prometheus.HistogramVec
	pluginName string
}

type ObservedIHttpClient struct {
	usdc.IHttpClient
	metric metricDetails
}

func NewObservedIHttpClient(origin usdc.IHttpClient, pluginName string) *ObservedIHttpClient {
	return &ObservedIHttpClient{
		IHttpClient: origin,
		metric: metricDetails{
			histogram:  usdcClientHistogram,
			pluginName: pluginName,
		},
	}
}

func (o *ObservedIHttpClient) Get(ctx context.Context, url string, timeout time.Duration) ([]byte, int, error) {
	return withObservedHttpClient(o.metric, func() ([]byte, int, error) {
		return o.IHttpClient.Get(ctx, url, timeout)
	})
}

func withObservedHttpClient[T any](metric metricDetails, contract func() (T, int, error)) (T, int, error) {
	contractExecutionStarted := time.Now()
	value, status, err := contract()
	metric.histogram.
		WithLabelValues(
			metric.pluginName,
			strconv.FormatInt(int64(status), 10),
			strconv.FormatBool(err == nil),
		).
		Observe(float64(time.Since(contractExecutionStarted)))
	return value, status, err
}
