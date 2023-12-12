package merclib

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"

	"github.com/ethereum/go-ethereum/common/hexutil"

	evmtypes "github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/llo-feeds/generated/verifier_proxy"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/models"
)

const (
	MercuryBatchPath = "/api/v1/reports/bulk"
)

var (
	verifierProxyABI = evmtypes.MustGetABI(verifier_proxy.VerifierProxyMetaData.ABI)
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

// ReportVerifier is an interface for verifying mercury reports
//
//go:generate mockery --name ReportVerifier --output ./mocks/ --case=underscore
type ReportVerifier interface {
	// VerifyReports returns an error if any of the reports are invalid
	// or there was an error verifying them
	VerifyReports(ctx context.Context, signedReports [][]byte) error
}

type mercClient struct {
	verifier ReportVerifier
	creds    *models.MercuryCredentials
	h        Doer
	lggr     logger.Logger
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
	verifier ReportVerifier) *mercClient {
	return &mercClient{
		creds:    creds,
		h:        doer,
		lggr:     lggr,
		verifier: verifier,
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

	req, err := GenerateBatchRequest(ctx, feedIDs, m.creds.URL, m.creds.Username, m.creds.Password)
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

	err = m.verifier.VerifyReports(ctx, func() (sr [][]byte) {
		for _, rwc := range reportsWithContext {
			sr = append(sr, rwc.RawFullReport)
		}
		return sr
	}())
	if err != nil {
		return nil, fmt.Errorf("could not verify report: %w", err)
	}

	return reportsWithContext, nil
}
