package usdc

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/tokendata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/tokendata/http"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

const (
	apiVersion                = "v1"
	attestationPath           = "attestations"
	defaultAttestationTimeout = 5 * time.Second
)

type attestationStatus string

const (
	attestationStatusSuccess attestationStatus = "complete"
	attestationStatusPending attestationStatus = "pending_confirmations"
)

// messageAndAttestation has to match the onchain struct `MessageAndAttestation` in the
// USDC token pool.
type messageAndAttestation struct {
	Message     []byte
	Attestation []byte
}

func (m messageAndAttestation) AbiString() string {
	return `
	[{
		"components": [
			{"name": "message", "type": "bytes"},
			{"name": "attestation", "type": "bytes"}
		],
		"type": "tuple"
	}]`
}

func (m messageAndAttestation) Validate() error {
	if len(m.Message) == 0 {
		return errors.New("message must be non-empty")
	}
	if len(m.Attestation) == 0 {
		return errors.New("attestation must be non-empty")
	}
	return nil
}

type TokenDataReader struct {
	lggr                  logger.Logger
	usdcReader            ccipdata.USDCReader
	httpClient            http.IHttpClient
	attestationApi        *url.URL
	attestationApiTimeout time.Duration
}

type attestationResponse struct {
	Status      attestationStatus `json:"status"`
	Attestation string            `json:"attestation"`
}

var _ tokendata.Reader = &TokenDataReader{}

func NewUSDCTokenDataReader(lggr logger.Logger, usdcReader ccipdata.USDCReader, usdcAttestationApi *url.URL, usdcAttestationApiTimeoutSeconds int) *TokenDataReader {
	timeout := time.Duration(usdcAttestationApiTimeoutSeconds) * time.Second
	if usdcAttestationApiTimeoutSeconds == 0 {
		timeout = defaultAttestationTimeout
	}
	return &TokenDataReader{
		lggr:                  lggr,
		usdcReader:            usdcReader,
		httpClient:            http.NewObservedIHttpClient(&http.HttpClient{}),
		attestationApi:        usdcAttestationApi,
		attestationApiTimeout: timeout,
	}
}

func NewUSDCTokenDataReaderWithHttpClient(origin TokenDataReader, httpClient http.IHttpClient) *TokenDataReader {
	return &TokenDataReader{
		lggr:                  origin.lggr,
		usdcReader:            origin.usdcReader,
		httpClient:            httpClient,
		attestationApi:        origin.attestationApi,
		attestationApiTimeout: origin.attestationApiTimeout,
	}
}

func (s *TokenDataReader) ReadTokenData(ctx context.Context, msg internal.EVM2EVMOnRampCCIPSendRequestedWithMeta) (messageAndAttestation []byte, err error) {
	messageBody, err := s.getUSDCMessageBody(ctx, msg)
	if err != nil {
		return []byte{}, errors.Wrap(err, "failed getting the USDC message body")
	}

	s.lggr.Infow("Calling attestation API", "messageBodyHash", hexutil.Encode(messageBody[:]), "messageID", hexutil.Encode(msg.MessageId[:]))

	// The attestation API expects the hash of the message body
	attestationResp, err := s.callAttestationApi(ctx, utils.Keccak256Fixed(messageBody))
	if err != nil {
		return []byte{}, errors.Wrap(err, "failed calling usdc attestation API ")
	}

	if attestationResp.Status != attestationStatusSuccess {
		return []byte{}, tokendata.ErrNotReady
	}

	// The USDC pool needs a combination of the message body and the attestation
	messageAndAttestation, err = encodeMessageAndAttestation(messageBody, attestationResp.Attestation)
	if err != nil {
		return nil, fmt.Errorf("failed to encode messageAndAttestation : %w", err)
	}

	return messageAndAttestation, nil
}

// encodeMessageAndAttestation encodes the message body and attestation into a single byte array
// that is readable onchain.
func encodeMessageAndAttestation(messageBody []byte, attestation string) ([]byte, error) {
	attestationBytes, err := hex.DecodeString(strings.TrimPrefix(attestation, "0x"))
	if err != nil {
		return nil, fmt.Errorf("failed to decode response attestation: %w", err)
	}

	return abihelpers.EncodeAbiStruct[messageAndAttestation](messageAndAttestation{
		Message:     messageBody,
		Attestation: attestationBytes,
	})
}

func (s *TokenDataReader) getUSDCMessageBody(ctx context.Context, msg internal.EVM2EVMOnRampCCIPSendRequestedWithMeta) ([]byte, error) {
	parsedMsgBody, err := s.usdcReader.GetLastUSDCMessagePriorToLogIndexInTx(ctx, int64(msg.LogIndex), msg.TxHash)
	if err != nil {
		return []byte{}, err
	}
	s.lggr.Infow("Got USDC message body", "messageBody", hexutil.Encode(parsedMsgBody), "messageID", hexutil.Encode(msg.MessageId[:]))
	return parsedMsgBody, nil
}

func (s *TokenDataReader) callAttestationApi(ctx context.Context, usdcMessageHash [32]byte) (attestationResponse, error) {
	fullAttestationUrl := fmt.Sprintf("%s/%s/%s/0x%x", s.attestationApi, apiVersion, attestationPath, usdcMessageHash)
	body, _, err := s.httpClient.Get(ctx, fullAttestationUrl, s.attestationApiTimeout)
	if err != nil {
		return attestationResponse{}, err
	}
	var response attestationResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return attestationResponse{}, err
	}
	if response.Status == "" {
		return attestationResponse{}, fmt.Errorf("invalid attestation response: %s", string(body))
	}
	return response, nil
}

func (s *TokenDataReader) Close(qopts ...pg.QOpt) error {
	return s.usdcReader.Close(qopts...)
}
