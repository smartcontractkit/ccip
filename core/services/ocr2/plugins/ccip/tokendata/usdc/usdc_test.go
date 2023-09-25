package usdc

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

var (
	mockOnRampAddress    = utils.RandomAddress()
	mockUSDCTokenAddress = utils.RandomAddress()
	mockMsgTransmitter   = utils.RandomAddress()
)

func TestUSDCReader_callAttestationApi(t *testing.T) {
	t.Skipf("Skipping test because it uses the real USDC attestation API")
	usdcMessageHash := "912f22a13e9ccb979b621500f6952b2afd6e75be7eadaed93fc2625fe11c52a2"
	attestationURI, err := url.ParseRequestURI("https://iris-api-sandbox.circle.com")
	require.NoError(t, err)
	lggr := logger.TestLogger(t)
	usdcReader, err := ccipdata.NewUSDCReader(lggr, mockMsgTransmitter, nil)
	require.NoError(t, err)
	usdcService := NewUSDCTokenDataReader(lggr, usdcReader, attestationURI)

	attestation, err := usdcService.callAttestationApi(context.Background(), [32]byte(common.FromHex(usdcMessageHash)))
	require.NoError(t, err)

	require.Equal(t, attestationStatusPending, attestation.Status)
	require.Equal(t, "PENDING", attestation.Attestation)
}

func TestUSDCReader_callAttestationApiMock(t *testing.T) {
	response := attestationResponse{
		Status:      attestationStatusSuccess,
		Attestation: "720502893578a89a8a87982982ef781c18b193",
	}

	ts := getMockUSDCEndpoint(t, response)
	defer ts.Close()
	attestationURI, err := url.ParseRequestURI(ts.URL)
	require.NoError(t, err)

	lggr := logger.TestLogger(t)
	usdcReader, err := ccipdata.NewUSDCReader(lggr, mockMsgTransmitter, nil)
	require.NoError(t, err)
	usdcService := NewUSDCTokenDataReader(lggr, usdcReader, attestationURI)
	attestation, err := usdcService.callAttestationApi(context.Background(), utils.RandomBytes32())
	require.NoError(t, err)

	require.Equal(t, response.Status, attestation.Status)
	require.Equal(t, response.Attestation, attestation.Attestation)
}

func TestUSDCReader_callAttestationApiMockError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()
	attestationURI, err := url.ParseRequestURI(ts.URL)
	require.NoError(t, err)

	lggr := logger.TestLogger(t)
	usdcReader, err := ccipdata.NewUSDCReader(lggr, mockMsgTransmitter, nil)
	require.NoError(t, err)
	usdcService := NewUSDCTokenDataReader(lggr, usdcReader, attestationURI)
	_, err = usdcService.callAttestationApi(context.Background(), utils.RandomBytes32())
	require.Error(t, err)
}

func getMockUSDCEndpoint(t *testing.T, response attestationResponse) *httptest.Server {
	responseBytes, err := json.Marshal(response)
	require.NoError(t, err)

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(responseBytes)
		require.NoError(t, err)
	}))
}

func TestGetUSDCMessageBody(t *testing.T) {

	sourceChainEventsMock := ccipdata.MockUSDCReader{}
	sourceChainEventsMock.On("GetLastUSDCMessagePriorToLogIndexInTx", mock.Anything, mock.Anything, mock.Anything).Return(expectedBody, nil)

	lggr := logger.TestLogger(t)
	usdcReader, err := ccipdata.NewUSDCReader(lggr, mockMsgTransmitter, nil)
	require.NoError(t, err)
	usdcService := NewUSDCTokenDataReader(lggr, usdcReader, attestationURI)

	// Make the first call and assert the underlying function is called
	body, err := usdcService.getUSDCMessageBody(context.Background(), internal.EVM2EVMOnRampCCIPSendRequestedWithMeta{})
	require.NoError(t, err)
	require.Equal(t, body, parsedBody)

	sourceChainEventsMock.AssertNumberOfCalls(t, "GetLastUSDCMessagePriorToLogIndexInTx", 1)

	// Make another call and assert that the cache is used
	body, err = usdcService.getUSDCMessageBody(context.Background(), internal.EVM2EVMOnRampCCIPSendRequestedWithMeta{})
	require.NoError(t, err)
	require.Equal(t, body, parsedBody)
	sourceChainEventsMock.AssertNumberOfCalls(t, "GetLastUSDCMessagePriorToLogIndexInTx", 1)
}
