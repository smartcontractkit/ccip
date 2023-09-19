package usdc

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

var (
	mockOnRampAddress    = utils.RandomAddress()
	mockUSDCTokenAddress = utils.RandomAddress()
)

func TestUSDCService_callAttestationApi(t *testing.T) {
	t.Skipf("Skipping test because it uses the real USDC attestation API")
	usdcMessageHash := "912f22a13e9ccb979b621500f6952b2afd6e75be7eadaed93fc2625fe11c52a2"
	attestationURI, err := url.ParseRequestURI("https://iris-api-sandbox.circle.com")
	require.NoError(t, err)
	usdcService := NewUSDCTokenDataReader(nil, mockUSDCTokenAddress, mockOnRampAddress, attestationURI, 420)

	attestation, err := usdcService.callAttestationApi(context.Background(), [32]byte(common.FromHex(usdcMessageHash)))
	require.NoError(t, err)

	require.Equal(t, AttestationStatusPending, attestation.Status)
	require.Equal(t, "PENDING", attestation.Attestation)
}

func TestUSDCService_callAttestationApiMock(t *testing.T) {
	response := attestationResponse{
		Status:      AttestationStatusSuccess,
		Attestation: "720502893578a89a8a87982982ef781c18b193",
	}

	ts := getMockUSDCEndpoint(t, response)
	defer ts.Close()
	attestationURI, err := url.ParseRequestURI(ts.URL)
	require.NoError(t, err)

	usdcService := NewUSDCTokenDataReader(nil, mockUSDCTokenAddress, mockOnRampAddress, attestationURI, 420)
	attestation, err := usdcService.callAttestationApi(context.Background(), utils.RandomBytes32())
	require.NoError(t, err)

	require.Equal(t, response.Status, attestation.Status)
	require.Equal(t, response.Attestation, attestation.Attestation)
}

func TestUSDCService_callAttestationApiMockError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()
	attestationURI, err := url.ParseRequestURI(ts.URL)
	require.NoError(t, err)

	usdcService := NewUSDCTokenDataReader(nil, mockUSDCTokenAddress, mockOnRampAddress, attestationURI, 420)
	_, err = usdcService.callAttestationApi(context.Background(), utils.RandomBytes32())
	require.Error(t, err)
}

func TestUSDCService_IsAttestationComplete(t *testing.T) {
	response := attestationResponse{
		Status:      AttestationStatusSuccess,
		Attestation: "720502893578a89a8a87982982ef781c18b193",
	}

	attestationBytes, err := hex.DecodeString(response.Attestation)
	require.NoError(t, err)

	ts := getMockUSDCEndpoint(t, response)
	defer ts.Close()

	seqNum := uint64(23825)
	txHash := utils.RandomBytes32()
	logIndex := int64(4)

	eventsClient := ccipdata.MockReader{}
	eventsClient.On("GetSendRequestsBetweenSeqNums",
		mock.Anything,
		mockOnRampAddress,
		seqNum,
		seqNum,
		0,
	).Return([]ccipdata.Event[evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested]{
		{
			Data: evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested{
				Raw: types.Log{
					TxHash: txHash,
					Index:  uint(logIndex),
				},
			},
		},
	}, nil)

	eventsClient.On("GetLastUSDCMessagePriorToLogIndexInTx",
		mock.Anything,
		logIndex,
		common.Hash(txHash),
	).Return(attestationBytes, nil)
	attestationURI, err := url.ParseRequestURI(ts.URL)
	require.NoError(t, err)

	usdcService := NewUSDCTokenDataReader(&eventsClient, mockUSDCTokenAddress, mockOnRampAddress, attestationURI, 420)
	isReady, attestation, err := usdcService.IsTokenDataReady(context.Background(), seqNum)
	require.NoError(t, err)

	require.True(t, isReady)
	require.Equal(t, attestationBytes, attestation)
	require.Equal(t, isReady, response.Status == AttestationStatusSuccess)
}

func getMockUSDCEndpoint(t *testing.T, response attestationResponse) *httptest.Server {
	responseBytes, err := json.Marshal(response)
	require.NoError(t, err)

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(responseBytes)
		require.NoError(t, err)
	}))
}

// Asserts the hard coded event signature matches Keccak256("MessageSent(bytes)")
func TestGetUSDCServiceSourceLPFilters(t *testing.T) {
	chainId := uint64(420)
	usdcService := NewUSDCTokenDataReader(nil, mockUSDCTokenAddress, mockOnRampAddress, nil, chainId)

	filters := usdcService.GetSourceLogPollerFilters()
	expectedTransmitterAddress := messageTransmitterMapping[chainId]

	require.Equal(t, 1, len(filters))
	filter := filters[0]
	require.Equal(t, logpoller.FilterName(MESSAGE_SENT_FILTER_NAME, expectedTransmitterAddress.Hex()), filter.Name)
	hash, err := utils.Keccak256([]byte("MessageSent(bytes)"))
	require.NoError(t, err)
	require.Equal(t, hash, filter.EventSigs[0].Bytes())
	require.Equal(t, expectedTransmitterAddress, filter.Addresses[0])
}
