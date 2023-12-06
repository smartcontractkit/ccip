package merclib

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	evmclient "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/llo_feeds"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/models"
)

const (
	MercuryBatchPath = "/api/v1/reports/bulk"
)

var (
	verifierProxyABI = evmtypes.MustGetABI(llo_feeds.LLOVerifierProxyMetaData.ABI)
)

// MercuryV03Response represents the response from Mercury from the batch endpoint
type MercuryV03Response struct {
	Reports []MercuryV03Report `json:"reports"`
}

// MercuryV03Report represents a single report returned by Mercury
// See https://www.notion.so/chainlink/RFC-User-Data-Streams-Integrator-Guide-MASTER-6051b91c97fd4b9c8fdaebd548fd0785?pvs=4#d266d26349284bef8e4dbbb5460b3bce
type MercuryV03Report struct {
	// FeedID is the ID of the feed that this report is for
	// Feed IDs are unique per feed, and are randomly generated.
	// See https://www.notion.so/chainlink/Schema-and-Feed-ID-Registry-4a0203da5e124d7092297a1359cb642f?pvs=4#fd0db0902a13427faa7518d3f4fa7207
	// for more information. The underlying type is a bytes32, this field is a bytes32
	// hex encoded with a leading 0x.
	FeedID string `json:"feedID"`
	// ValidFromTimestamp is the report's earliest applicable timestamp.
	ValidFromTimestamp uint32 `json:"validFromTimestamp"`
	// ObservationsTimestamp is the report's latest applicable timestamp
	ObservationsTimestamp uint32 `json:"observationsTimestamp"`
	// FullReport is the full mercury DON report, encoded as a hex string, with a leading 0x.
	FullReport string `json:"fullReport"`
}

// Doer is an interface for sending HTTP requests
//
//go:generate mockery --name Doer --output ./mocks/ --case=underscore
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

type mercClient struct {
	ethClient            evmclient.Client
	verifierProxyAddress common.Address
	wrappedNativeAddress common.Address
	// fromAddress is the address that is used to call the verifier proxy contract
	// in practice this is the OCR2 onchain key as an ethereum address
	fromAddress common.Address
	creds       *models.MercuryCredentials
	h           Doer
	lggr        logger.Logger
}

// MercuryClient is an interface for fetching offchain token data from Mercury
//
//go:generate mockery --name MercuryClient --output ./mocks/ --case=underscore
type MercuryClient interface {
	// BatchFetchPrices returns the prices for the given feed IDs
	BatchFetchPrices(ctx context.Context, feedIDs [][32]byte) ([]*ReportWithContext, error)
}

// NewMercuryClient returns a new MercuryClient instance
func NewMercuryClient(
	creds *models.MercuryCredentials,
	doer Doer,
	lggr logger.Logger,
	ethClient evmclient.Client,
	fromAddress,
	verifierProxyAddress,
	wrappedNativeAddress common.Address) *mercClient {
	return &mercClient{
		creds:                creds,
		h:                    doer,
		lggr:                 lggr,
		ethClient:            ethClient,
		verifierProxyAddress: verifierProxyAddress,
		wrappedNativeAddress: wrappedNativeAddress,
		fromAddress:          fromAddress,
	}
}

func (m *mercClient) BatchFetchPrices(ctx context.Context, fids [][32]byte) ([]*ReportWithContext, error) {
	feedIDs := func() (r []string) {
		for _, fid := range fids {
			r = append(r, hexutil.Encode(fid[:]))
		}
		return r
	}()

	// sort the feed IDs for deterministic batching
	slices.Sort(feedIDs)

	req, err := m.generateRequest(ctx, feedIDs)
	if err != nil {
		return nil, fmt.Errorf("generating request: %w", err)
	}

	resp, err := m.h.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 (%d) response: %s", resp.StatusCode, string(body))
	}

	var r MercuryV03Response
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, fmt.Errorf("unmarshaling response: %w, body: %s", err, string(body))
	}

	var reportsWithContext []*ReportWithContext
	for _, report := range r.Reports {
		decoded, err := hexutil.Decode(report.FullReport)
		if err != nil {
			return nil, fmt.Errorf("decoding full mercury report: %w", err)
		}
		rwc, err := DecodeFullReportAndReportData(decoded)
		if err != nil {
			return nil, fmt.Errorf("decoding full mercury report: %w", err)
		}
		reportsWithContext = append(reportsWithContext, rwc)
	}

	err = m.verifyReports(ctx, func() (sr [][]byte) {
		for _, rwc := range reportsWithContext {
			sr = append(sr, rwc.FullReport.ReportBlob)
		}
		return sr
	}())
	if err != nil {
		return nil, fmt.Errorf("could not verify report onchain: %w", err)
	}

	return reportsWithContext, nil
}

func (m *mercClient) generateRequest(ctx context.Context, feedIDs []string) (*http.Request, error) {
	// mostly cribbed from streams_lookup.go
	tsNow := time.Now().UTC().UnixMilli()
	params := fmt.Sprintf("?%s=%s&%s=%d",
		// feedIDs param
		"feedIDs", strings.Join(feedIDs, ","),
		// timestamp param
		"timestamp", tsNow,
	)
	reqURL := fmt.Sprintf("%s%s%s",
		cleanMercURL(m.creds.URL),
		MercuryBatchPath,
		params,
	)
	m.lggr.Debugw("generating request URL",
		"mercuryUsername", m.creds.Username,
		"feedIDs", feedIDs,
		"reqURL", reqURL)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	sig := m.generateHMAC(http.MethodGet, MercuryBatchPath+params, nil, m.creds.Username, m.creds.Password, tsNow)
	req.Header.Set("Authorization", m.creds.Username)
	req.Header.Set("X-Authorization-Timestamp", fmt.Sprintf("%d", tsNow))
	req.Header.Set("X-Authorization-Signature-SHA256", sig)
	req.Header.Set("Content-Type", "application/json")
	// TODO: check if we should have another header for CCIP specifically, similar to automation

	return req, nil
}

func (m *mercClient) generateHMAC(method, path string, body []byte, clientId, secret string, ts int64) (sig string) {
	bodyHash := sha256.New()
	bodyHash.Write(body)
	hashString := fmt.Sprintf("%s %s %s %s %d",
		method,
		path,
		hex.EncodeToString(bodyHash.Sum(nil)),
		clientId,
		ts)
	signedMessage := hmac.New(sha256.New, []byte(secret))
	signedMessage.Write([]byte(hashString))
	return hex.EncodeToString(signedMessage.Sum(nil))
}

func (m *mercClient) verifyReports(ctx context.Context, signedReports [][]byte) error {
	// payment options don't really matter since an eth_call won't pay anything in the end
	// NOTE: the fromAddress passed into the client must have 100% discount on the mercury fee manager contract.
	calldata, err := verifierProxyABI.Pack("verifyBulk", signedReports, m.wrappedNativeAddress.Bytes())
	if err != nil {
		return fmt.Errorf("failed to pack verifyBulk: %w", err)
	}
	callMsg := ethereum.CallMsg{
		From:     m.fromAddress,
		Data:     calldata,
		To:       &m.verifierProxyAddress,
		GasPrice: nil,
		Gas:      0,
	}
	m.lggr.Debugw("calling verifier contract",
		"verifierProxyAddress", m.verifierProxyAddress.String(), "calldata", hexutil.Encode(calldata), "callMsg", callMsg)
	_, err = m.ethClient.CallContract(ctx, callMsg, nil)
	if err != nil {
		return fmt.Errorf("failed to call verifier contract at %s: %w, calldata: %s", m.verifierProxyAddress.String(), err, hexutil.Encode(calldata))
	}

	// simulation passing means verification succeeded
	return nil
}

func cleanMercURL(url string) string {
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}
	return url
}
