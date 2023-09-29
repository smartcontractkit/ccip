package usdc_test

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/tokendata/usdc"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

type attestationResponse struct {
	Status      string `json:"status"`
	Attestation string `json:"attestation"`
}

func TestUSDCReader_ReadTokenData(t *testing.T) {
	response := attestationResponse{
		Status:      "complete",
		Attestation: "0x9049623e91719ef2aa63c55f357be2529b0e7122ae552c18aff8db58b4633c4d3920ff03d3a6d1ddf11f06bf64d7fd60d45447ac81f527ba628877dc5ca759651b08ffae25a6d3b1411749765244f0a1c131cbfe04430d687a2e12fd9d2e6dc08e118ad95d94ad832332cf3c4f7a4f3da0baa803b7be024b02db81951c0f0714de1b",
	}
	// Message is the bytes itself from MessageSend(bytes message)
	// i.e. ABI parsed.
	message := "0x0000000000000001000000020000000000048d71000000000000000000000000eb08f243e5d3fcff26a9e38ae5520a669f4019d000000000000000000000000023a04d5935ed8bc8e3eb78db3541f0abfb001c6e0000000000000000000000006cb3ed9b441eb674b58495c8b3324b59faff5243000000000000000000000000000000005425890298aed601595a70ab815c96711a31bc65000000000000000000000000ab4f961939bfe6a93567cc57c59eed7084ce2131000000000000000000000000000000000000000000000000000000000000271000000000000000000000000035e08285cfed1ef159236728f843286c55fc0861"
	expectedMessageAndAttestation := "0x00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000000f80000000000000001000000020000000000048d71000000000000000000000000eb08f243e5d3fcff26a9e38ae5520a669f4019d000000000000000000000000023a04d5935ed8bc8e3eb78db3541f0abfb001c6e0000000000000000000000006cb3ed9b441eb674b58495c8b3324b59faff5243000000000000000000000000000000005425890298aed601595a70ab815c96711a31bc65000000000000000000000000ab4f961939bfe6a93567cc57c59eed7084ce2131000000000000000000000000000000000000000000000000000000000000271000000000000000000000000035e08285cfed1ef159236728f843286c55fc0861000000000000000000000000000000000000000000000000000000000000000000000000000000829049623e91719ef2aa63c55f357be2529b0e7122ae552c18aff8db58b4633c4d3920ff03d3a6d1ddf11f06bf64d7fd60d45447ac81f527ba628877dc5ca759651b08ffae25a6d3b1411749765244f0a1c131cbfe04430d687a2e12fd9d2e6dc08e118ad95d94ad832332cf3c4f7a4f3da0baa803b7be024b02db81951c0f0714de1b000000000000000000000000000000000000000000000000000000000000"
	lggr := logger.TestLogger(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		messageHash := utils.Keccak256Fixed(hexutil.MustDecode(message))
		expectedUrl := "/v1/attestations/0x" + hex.EncodeToString(messageHash[:])
		require.Equal(t, expectedUrl, r.URL.Path)

		responseBytes, err2 := json.Marshal(response)
		require.NoError(t, err2)

		_, err2 = w.Write(responseBytes)
		require.NoError(t, err2)
	}))

	defer ts.Close()

	seqNum := uint64(23825)
	txHash := utils.RandomBytes32()
	logIndex := int64(4)

	usdcReader := ccipdata.MockUSDCReader{}
	usdcReader.On("GetLastUSDCMessagePriorToLogIndexInTx",
		mock.Anything,
		logIndex,
		common.Hash(txHash),
	).Return(hexutil.MustDecode(message), nil)
	attestationURI, err := url.ParseRequestURI(ts.URL)
	require.NoError(t, err)

	usdcService := usdc.NewUSDCTokenDataReader(lggr, &usdcReader, attestationURI)
	msgAndAttestation, err := usdcService.ReadTokenData(context.Background(), internal.EVM2EVMOnRampCCIPSendRequestedWithMeta{
		InternalEVM2EVMMessage: evm_2_evm_offramp.InternalEVM2EVMMessage{
			SequenceNumber: seqNum,
		},
		TxHash:   txHash,
		LogIndex: uint(logIndex),
	})
	require.NoError(t, err)
	// Expected attestation for parsed body.
	require.Equal(t, expectedMessageAndAttestation, hexutil.Encode(msgAndAttestation))
}
