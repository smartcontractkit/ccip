package lbtc

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"golang.org/x/time/rate"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/tokendata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/tokendata/http"
)

// TODO: double check the validty of default values for lombard's API after checking docs
const (
	apiVersion                = "v1"
	attestationPath           = "deposits/getByHash"
	defaultAttestationTimeout = 5 * time.Second

	// defaultCoolDownDurationSec defines the default time to wait after getting rate limited.
	// this value is only used if the 429 response does not contain the Retry-After header
	defaultCoolDownDuration = 5 * time.Minute

	// maxCoolDownDuration defines the maximum duration we can wait till firing the next request
	maxCoolDownDuration = 10 * time.Minute

	// defaultRequestInterval defines the rate in requests per second that the attestation API can be called.
	// this is set according to the APIs documentated 10 requests per second rate limit.
	defaultRequestInterval = 100 * time.Millisecond

	// APIIntervalRateLimitDisabled is a special value to disable the rate limiting.
	APIIntervalRateLimitDisabled = -1
	// APIIntervalRateLimitDefault is a special value to select the default rate limit interval.
	APIIntervalRateLimitDefault = 0
)

type attestationStatus string

const (
	attestationStatusUnspecified     attestationStatus = "NOTARIZATION_STATUS_UNSPECIFIED"
	attestationStatusPending         attestationStatus = "NOTARIZATION_STATUS_PENDING"
	attestationStatusSubmitted       attestationStatus = "NOTARIZATION_STATUS_SUBMITTED"
	attestationStatusSessionApproved attestationStatus = "NOTARIZATION_STATUS_SESSION_APPROVED"
	attestationStatusFailed          attestationStatus = "NOTARIZATION_STATUS_FAILED"
)

var (
	ErrUnknownResponse = errors.New("unexpected response from attestation API")
)

type TokenDataReader struct {
	lggr                  logger.Logger
	httpClient            http.IHttpClient
	attestationApi        *url.URL
	attestationApiTimeout time.Duration
	lbtcTokenAddress      common.Address
	rate                  *rate.Limiter

	// coolDownUntil defines whether requests are blocked or not.
	coolDownUntil time.Time
	coolDownMu    *sync.RWMutex
}

type messageAttestationResponse struct {
	MessageHash string            `json:"message_hash"`
	Status      attestationStatus `json:"status"`
	Attestation string            `json:"attestation"`
}

// TODO: Adjust after checking API docs
type attestationResponse struct {
	Attestations []messageAttestationResponse `json:"attestations"`
}

// TODO: Implement encoding/decoding

var _ tokendata.Reader = &TokenDataReader{}

func NewLBTCTokenDataReader(
	lggr logger.Logger,
	lbtcAttestationApi *url.URL,
	lbtcAttestationApiTimeoutSeconds int,
	lbtcTokenAddress common.Address,
	requestInterval time.Duration,
) *TokenDataReader {
	timeout := time.Duration(lbtcAttestationApiTimeoutSeconds) * time.Second
	if lbtcAttestationApiTimeoutSeconds == 0 {
		timeout = defaultAttestationTimeout
	}

	if requestInterval == APIIntervalRateLimitDisabled {
		requestInterval = 0
	} else if requestInterval == APIIntervalRateLimitDefault {
		requestInterval = defaultRequestInterval
	}

	return &TokenDataReader{
		lggr:                  lggr,
		httpClient:            http.NewObservedIHttpClient(&http.HttpClient{}),
		attestationApi:        lbtcAttestationApi,
		attestationApiTimeout: timeout,
		lbtcTokenAddress:      lbtcTokenAddress,
		coolDownMu:            &sync.RWMutex{},
		rate:                  rate.NewLimiter(rate.Every(requestInterval), 1),
	}
}

func NewLBTCTokenDataReaderWithHttpClient(
	origin TokenDataReader,
	httpClient http.IHttpClient,
	lbtcTokenAddress common.Address,
	requestInterval time.Duration,
) *TokenDataReader {
	return &TokenDataReader{
		lggr:                  origin.lggr,
		httpClient:            httpClient,
		attestationApi:        origin.attestationApi,
		attestationApiTimeout: origin.attestationApiTimeout,
		coolDownMu:            origin.coolDownMu,
		lbtcTokenAddress:      lbtcTokenAddress,
		rate:                  rate.NewLimiter(rate.Every(requestInterval), 1),
	}
}

// ReadTokenData queries the LBTC attestation API.
func (s *TokenDataReader) ReadTokenData(ctx context.Context, msg cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta, tokenIndex int) ([]byte, error) {
	if tokenIndex < 0 || tokenIndex >= len(msg.TokenAmounts) {
		return nil, fmt.Errorf("token index out of bounds")
	}

	if s.inCoolDownPeriod() {
		// rate limiting cool-down period, we prevent new requests from being sent
		return nil, tokendata.ErrRequestsBlocked
	}

	if s.rate != nil {
		// Wait blocks until it the attestation API can be called or the
		// context is Done.
		if waitErr := s.rate.Wait(ctx); waitErr != nil {
			return nil, fmt.Errorf("lbtc rate limiting error: %w", waitErr)
		}
	}

	messageBody, err := s.getLBTCMessageBody(ctx, msg, tokenIndex)
	if err != nil {
		return []byte{}, errors.Wrap(err, "failed getting the LBTC message body")
	}

	msgID := hexutil.Encode(msg.MessageID[:])
	messageBodyHash := sha256.Sum256(messageBody)
	messageBodyHashHex := hexutil.Encode(messageBodyHash[:])
	s.lggr.Infow("Calling attestation API", "messageBodyHash", messageBodyHashHex, "messageID", msgID)

	attestationResp, err := s.callAttestationApi(ctx, messageBodyHash)
	if err != nil {
		return nil, err
	}
	if attestationResp.Attestations == nil || len(attestationResp.Attestations) == 0 {
		return nil, errors.New("attestation response is empty")
	}
	if len(attestationResp.Attestations) > 1 {
		s.lggr.Warnw("Multiple attestations received, expected one", "attestations", attestationResp.Attestations)
	}
	var attestation messageAttestationResponse
	for _, attestationCandidate := range attestationResp.Attestations {
		if attestationCandidate.MessageHash == messageBodyHashHex {
			attestation = attestationCandidate
		}
	}
	s.lggr.Infow("Got response from attestation API", "messageID", msgID,
		"attestationStatus", attestation.Status, "attestation", attestation)
	switch attestation.Status {
	case attestationStatusSessionApproved:
		messageAndAttestation, err := encodeMessageAndAttestation(messageBody, attestation.Attestation)
		if err != nil {
			return nil, fmt.Errorf("failed to encode messageAndAttestation : %w", err)
		}
		return messageAndAttestation, nil
	case attestationStatusPending:
		return nil, tokendata.ErrNotReady
	case attestationStatusSubmitted:
		return nil, tokendata.ErrNotReady
	default:
		s.lggr.Errorw("Unexpected response from attestation API", "attestation", attestation)
		return nil, ErrUnknownResponse
	}
}

func (s *TokenDataReader) getLBTCMessageBody(ctx context.Context, msg cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta, tokenIndex int) ([]byte, error) {
	return nil, nil
}

func (s *TokenDataReader) callAttestationApi(ctx context.Context, lbtcMessageHash [32]byte) (attestationResponse, error) {
	_, _, _, err := s.httpClient.Get(ctx, "", s.attestationApiTimeout)
	switch {
	case errors.Is(err, tokendata.ErrRateLimit):
		s.setCoolDownPeriod(defaultCoolDownDuration)
		return attestationResponse{}, tokendata.ErrRateLimit
	case err != nil:
		return attestationResponse{}, err
	}
	return attestationResponse{}, nil
}

func encodeMessageAndAttestation(messageBody []byte, attestation string) ([]byte, error) {
	return nil, nil
}

func (s *TokenDataReader) setCoolDownPeriod(d time.Duration) {
	s.coolDownMu.Lock()
	if d > maxCoolDownDuration {
		d = maxCoolDownDuration
	}
	s.coolDownUntil = time.Now().Add(d)
	s.coolDownMu.Unlock()
}

func (s *TokenDataReader) inCoolDownPeriod() bool {
	s.coolDownMu.RLock()
	defer s.coolDownMu.RUnlock()
	return time.Now().Before(s.coolDownUntil)
}

func (s *TokenDataReader) Close() error {
	return nil
}
