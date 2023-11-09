package observability

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/prometheus/client_golang/prometheus"
	io_prometheus_client "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/tokendata/usdc"
)

// TODO avoid duplicating types from usdc.go
type attestationResponse struct {
	Status      string `json:"status"`
	Attestation string `json:"attestation"`
}

func TestUSDCMonitoring(t *testing.T) {

	lggr := logger.TestLogger(t)
	usdcReader := mocks.NewUSDCReader(t)
	msgBody := []byte{0xb0, 0xd1}
	usdcReader.On("GetLastUSDCMessagePriorToLogIndexInTx", mock.Anything, mock.Anything, mock.Anything).Return(msgBody, nil)

	// Create a fake USDC server.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		response := attestationResponse{
			Status:      "complete",
			Attestation: "720502893578a89a8a87982982ef781c18b193",
		}
		responseBytes, err := json.Marshal(response)
		require.NoError(t, err)
		_, err = w.Write(responseBytes)
		require.NoError(t, err)
	}))
	defer server.Close()
	attestationURI, err := url.ParseRequestURI(server.URL)
	require.NoError(t, err)

	// Service with mock http client.
	usdcService := usdc.NewUSDCTokenDataReader(lggr, usdcReader, attestationURI, 10)
	observedService := NewObservedUSDCTokenDataReader(*usdcService, "plugin")
	require.NotNil(t, observedService)
	msgAndAttestation, err := observedService.ReadTokenData(context.Background(), internal.EVM2EVMOnRampCCIPSendRequestedWithMeta{})
	require.NoError(t, err)
	require.NotNil(t, msgAndAttestation)
	expectedMessageAndAttestation := "0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000002b0d10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000013720502893578a89a8a87982982ef781c18b19300000000000000000000000000"
	require.Equal(t, expectedMessageAndAttestation, hexutil.Encode(msgAndAttestation))

	// Check that the metrics are updated.
	histogram := usdcHistogram
	assert.Equal(t, 0, counterFromHistogramByLabels(t, histogram, "plugin", "XYZMethod"))
	assert.Equal(t, 1, counterFromHistogramByLabels(t, histogram, "plugin", "ReadTokenData"))
	assert.Equal(t, 0, counterFromHistogramByLabels(t, histogram, "plugin", "Get"))
	assert.Equal(t, 1, counterFromHistogramByLabels(t, histogram, "plugin", "GetWithTimeout"))

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
