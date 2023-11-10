package observability

import (
	"context"
	"strconv"
	"time"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/tokendata/usdc"
)

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

func (o *ObservedIHttpClient) Get(ctx context.Context, url string) ([]byte, int, error) {
	return withObservedHttpClient(o.metric, "Get", func() ([]byte, int, error) {
		return o.IHttpClient.Get(ctx, url)
	})
}

func (o *ObservedIHttpClient) GetWithTimeout(ctx context.Context, url string, timeout time.Duration) ([]byte, int, error) {
	return withObservedHttpClient(o.metric, "GetWithTimeout", func() ([]byte, int, error) {
		return o.IHttpClient.GetWithTimeout(ctx, url, timeout)
	})
}

func withObservedHttpClient[T any](metric metricDetails, function string, contract func() (T, int, error)) (T, int, error) {
	contractExecutionStarted := time.Now()
	value, status, err := contract()
	metric.histogram.
		WithLabelValues(
			metric.pluginName,
			function,
			strconv.FormatInt(int64(status), 10),
			strconv.FormatBool(err == nil),
		).
		Observe(float64(time.Since(contractExecutionStarted)))
	return value, status, err
}
