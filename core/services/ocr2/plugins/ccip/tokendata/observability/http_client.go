package observability

import (
	"context"
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
			histogram:  usdcHistogram,
			pluginName: pluginName,
		},
	}
}

func (o *ObservedIHttpClient) Get(ctx context.Context, url string) ([]byte, error) {
	return withObservedContract(o.metric, "Get", func() ([]byte, error) {
		return o.IHttpClient.Get(ctx, url)
	})
}

func (o *ObservedIHttpClient) GetWithTimeout(ctx context.Context, url string, timeout time.Duration) ([]byte, error) {
	return withObservedContract(o.metric, "GetWithTimeout", func() ([]byte, error) {
		return o.IHttpClient.GetWithTimeout(ctx, url, timeout)
	})
}
