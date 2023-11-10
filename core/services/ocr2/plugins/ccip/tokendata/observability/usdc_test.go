package observability

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	io_prometheus_client "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/tokendata/usdc"
)

var (
	pluginName = "testplugin"
)

type expected struct {
	pluginName string
	function   string
	success    bool
	count      int
}

func TestUSDCMonitoring(t *testing.T) {

	tests := []struct {
		name     string
		server   *httptest.Server
		requests int
		expected []expected
	}{
		{
			name:     "success",
			server:   newSuccessServer(t),
			requests: 5,
			expected: []expected{
				{pluginName, "ReadTokenData", true, 5},
				{pluginName, "ReadTokenData", false, 0},
				{pluginName, "GetWithTimeout", true, 5},
				{pluginName, "GetWithTimeout", false, 0},
			},
		},
		{
			name:     "rate_limited",
			server:   newRateLimitedServer(),
			requests: 26,
			expected: []expected{
				{pluginName, "ReadTokenData", true, 0},
				{pluginName, "ReadTokenData", false, 26},
				{pluginName, "GetWithTimeout", true, 0},
				{pluginName, "GetWithTimeout", false, 26},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testMonitoring(t, test.name, test.server, test.requests, test.expected, logger.TestLogger(t))
		})
	}

}

func testMonitoring(t *testing.T, name string, server *httptest.Server, requests int, expected []expected, log logger.Logger) {
	server.Start()
	defer server.Close()
	attestationURI, err := url.ParseRequestURI(server.URL)
	require.NoError(t, err)

	// Define test histogram (avoid side effects from other tests if using the real usdcHistogram).
	histogram := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "test_histogram_" + name,
		Help:    "Latency of calls to the USDC reader",
		Buckets: latencyBuckets,
	}, []string{"plugin", "function", "success"})

	// Mock USDC reader.
	usdcReader := mocks.NewUSDCReader(t)
	msgBody := []byte{0xb0, 0xd1}
	usdcReader.On("GetLastUSDCMessagePriorToLogIndexInTx", mock.Anything, mock.Anything, mock.Anything).Return(msgBody, nil)

	// Service with mock http client.
	usdcService := usdc.NewUSDCTokenDataReader(log, usdcReader, attestationURI, 0)
	observedHttpClient := &ObservedIHttpClient{
		IHttpClient: &usdc.HttpClient{},
		metric:      metricDetails{histogram, pluginName},
	}
	observedService := &ObservedUSDCTokenDataReader{
		TokenDataReader: *usdc.NewUSDCTokenDataReaderWithHttpClient(*usdcService, observedHttpClient),
		metric:          metricDetails{histogram, pluginName},
	}
	require.NotNil(t, observedService)

	for i := 0; i < requests; i++ {
		//msgAndAttestation, err := observedService.ReadTokenData(context.Background(), internal.EVM2EVMOnRampCCIPSendRequestedWithMeta{})
		_, _ = observedService.ReadTokenData(context.Background(), internal.EVM2EVMOnRampCCIPSendRequestedWithMeta{})
		//require.NoError(t, err)
		//require.NotNil(t, msgAndAttestation)
		//expectedMessageAndAttestation := "0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000002b0d10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000013720502893578a89a8a87982982ef781c18b19300000000000000000000000000"
		//require.Equal(t, expectedMessageAndAttestation, hexutil.Encode(msgAndAttestation))
	}

	// Check that the metrics are updated.
	//histogram := usdcHistogram
	assert.Equal(t, 0, counterFromHistogramByLabels(t, histogram, pluginName, "XYZMethod", "true"))
	assert.Equal(t, 0, counterFromHistogramByLabels(t, histogram, pluginName, "Get", "true"))
	assert.Equal(t, 0, counterFromHistogramByLabels(t, histogram, pluginName, "Get", "false"))
	for _, e := range expected {
		assert.Equal(t, e.count, counterFromHistogramByLabels(t, histogram, e.pluginName, e.function, strconv.FormatBool(e.success)))
	}
}

func counterFromHistogramByLabels(t *testing.T, histogramVec *prometheus.HistogramVec, labels ...string) int {
	observer, err := histogramVec.GetMetricWithLabelValues(labels...)
	require.NoError(t, err)

	metricCh := make(chan prometheus.Metric, 1)
	observer.(prometheus.Histogram).Collect(metricCh)
	close(metricCh)

	metric := <-metricCh
	pb := &io_prometheus_client.Metric{}
	err = metric.Write(pb)
	require.NoError(t, err)

	return int(pb.GetHistogram().GetSampleCount())
}

func newSuccessServer(t *testing.T) *httptest.Server {
	return httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		response := struct {
			Status      string `json:"status"`
			Attestation string `json:"attestation"`
		}{
			Status:      "complete",
			Attestation: "720502893578a89a8a87982982ef781c18b193",
		}
		responseBytes, err := json.Marshal(response)
		require.NoError(t, err)
		_, err = w.Write(responseBytes)
		require.NoError(t, err)
	}))
}

func newRateLimitedServer() *httptest.Server {
	return httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
	}))
}
