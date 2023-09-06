package customtokens

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

func TestUSDCService_TryGetAttestation(t *testing.T) {
	t.Skipf("Skipping test because it uses the real USDC attestation API")
	usdcMessageHash := "0x912f22a13e9ccb979b621500f6952b2afd6e75be7eadaed93fc2625fe11c52a2"
	usdcService := NewUSDCService("https://iris-api-sandbox.circle.com", 420)

	attestation, err := usdcService.TryGetAttestation(usdcMessageHash)
	require.NoError(t, err)

	require.Equal(t, USDCAttestationStatusPending, attestation.Status)
	require.Equal(t, "PENDING", attestation.Attestation)
}

func TestUSDCService_TryGetAttestationMock(t *testing.T) {
	response := USDCAttestationResponse{
		Status:      USDCAttestationStatusSuccess,
		Attestation: "720502893578a89a8a87982982ef781c18b193",
	}

	ts := getMockUSDCEndpoint(t, response)
	defer ts.Close()

	usdcService := NewUSDCService(ts.URL, 420)
	attestation, err := usdcService.TryGetAttestation("0x912f22a13e9ccb979b621500f6952b2afd6e75be7eadaed93fc2625fe11c52a2")
	require.NoError(t, err)

	require.Equal(t, response.Status, attestation.Status)
	require.Equal(t, response.Attestation, attestation.Attestation)
}

func TestUSDCService_TryGetAttestationMockError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	usdcService := NewUSDCService(ts.URL, 420)
	_, err := usdcService.TryGetAttestation("0x912f22a13e9ccb979b621500f6952b2afd6e75be7eadaed93fc2625fe11c52a2")
	require.Error(t, err)
}

func TestUSDCService_IsAttestationComplete(t *testing.T) {
	response := USDCAttestationResponse{
		Status:      USDCAttestationStatusSuccess,
		Attestation: "720502893578a89a8a87982982ef781c18b193",
	}

	ts := getMockUSDCEndpoint(t, response)
	defer ts.Close()

	usdcService := NewUSDCService(ts.URL, 420)
	isReady, attestation, err := usdcService.IsAttestationComplete("0x912f22a13e9ccb979b621500f6952b2afd6e75be7eadaed93fc2625fe11c52a2")
	require.NoError(t, err)

	require.True(t, isReady)
	require.Equal(t, response.Attestation, attestation)
}

func getMockUSDCEndpoint(t *testing.T, response USDCAttestationResponse) *httptest.Server {
	responseBytes, err := json.Marshal(response)
	require.NoError(t, err)

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(responseBytes)
		require.NoError(t, err)
	}))
}

// Asserts the hard coded event signature matches Keccak256("MessageSent(bytes)")
func TestGetUSDCServiceSourceLPFilters(t *testing.T) {
	usdcTokenAddress := common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")
	filters := GetUSDCServiceSourceLPFilters(usdcTokenAddress)

	require.Equal(t, 1, len(filters))
	filter := filters[0]
	require.Equal(t, logpoller.FilterName(USDC_MESSAGE_SENT_FILTER_NAME, usdcTokenAddress.Hex()), filter.Name)
	hash, err := utils.Keccak256([]byte("MessageSent(bytes)"))
	require.NoError(t, err)
	require.Equal(t, hash, filter.EventSigs[0].Bytes())
	require.Equal(t, usdcTokenAddress, filter.Addresses[0])
}
