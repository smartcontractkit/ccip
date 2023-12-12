package merclib

import (
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/chains/evmutil"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/ocr2key"
)

func XXXTestOnlyMercuryHandlerWithReportSigning(
	t *testing.T,
	prices map[[32]byte]*big.Int,
	ocr2KeyBundles []ocr2key.KeyBundle,
	feedIDToConfigDigest map[[32]byte][32]byte,
	f int,
	sharedSecret string) http.Handler {
	mux := http.NewServeMux()
	handler := &xxxTestOnlyMercuryHandlerWithSignedReports{
		ocr2Keys:             ocr2KeyBundles,
		feedIDToConfigDigest: feedIDToConfigDigest,
		t:                    t,
		prices:               prices,
		f:                    f,
		sharedSecret:         sharedSecret,
	}
	mux.Handle(MercuryBatchPath, handler)
	return mux
}

type xxxTestOnlyMercuryHandlerWithSignedReports struct {
	ocr2Keys             []ocr2key.KeyBundle
	feedIDToConfigDigest map[[32]byte][32]byte
	t                    *testing.T
	prices               map[[32]byte]*big.Int
	f                    int
	sharedSecret         string
}

func (x *xxxTestOnlyMercuryHandlerWithSignedReports) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	x.t.Log("received request to MercuryBatchPath: ", r.URL.Path, "headers:", r.Header, "query params:", r.URL.Query())
	// parse the feed IDs from the request query params
	feedIDsStr, ok := r.URL.Query()["feedIDs"]
	if !ok {
		w.Write([]byte("missing feedIDs query param"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var feedIDs [][32]byte
	fids := strings.Split(feedIDsStr[0], ",")
	for _, fid := range fids {
		decoded := hexutil.MustDecode(fid)
		var feedID [32]byte
		copy(feedID[:], decoded)
		if _, ok := x.prices[feedID]; !ok {
			w.Write([]byte(fmt.Sprintf("missing price for feedID %s", fid)))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		feedIDs = append(feedIDs, feedID)
	}
	// parse the timestamp but don't really use it
	// its in unix milliseconds after the epoch
	_, ok = r.URL.Query()["timestamp"]
	if !ok {
		w.Write([]byte("missing timestamp query param"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// verify the hmac signature
	authorizationTimestamp := r.Header.Get("X-Authorization-Timestamp")
	if authorizationTimestamp == "" {
		w.Write([]byte("missing X-Authorization-Timestamp header"))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	authTs, err := strconv.ParseInt(authorizationTimestamp, 10, 64)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("invalid X-Authorization-Timestamp header: %s", err)))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	authorizationSignature := r.Header.Get("X-Authorization-Signature-SHA256")
	if authorizationSignature == "" {
		w.Write([]byte("missing X-Authorization-Signature-SHA256 header"))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	username := r.Header.Get("Authorization")
	if username == "" {
		w.Write([]byte("missing Authorization header"))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	hmacString := GenerateHMAC(
		http.MethodGet,
		r.URL.RequestURI(),
		[]byte(""),
		username,
		authTs,
		x.sharedSecret,
		"testserver")
	if hmacString != authorizationSignature {
		w.Write([]byte(fmt.Sprintf("invalid signature, expected '%s' got '%s'", hmacString, authorizationSignature)))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// generate the report
	var reports []MercuryV03Report
	var (
		obsTs = uint32(time.Now().UTC().Unix())
		// set expiry to be 1 hour later
		expiresTs = obsTs + 3600
	)
	for _, feedID := range feedIDs {
		configDigest, ok := x.feedIDToConfigDigest[feedID]
		if !ok {
			w.Write([]byte(fmt.Sprintf("missing config digest for feedID %s", feedID)))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fullReport, err := generateReport(feedID, configDigest, obsTs, expiresTs, x.prices[feedID], x.ocr2Keys, x.f)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("failed to generate report: %s", err)))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		reports = append(reports, MercuryV03Report{
			FeedID:                hexutil.Encode(feedID[:]),
			ValidFromTimestamp:    obsTs,
			ObservationsTimestamp: obsTs,
			FullReport:            hexutil.Encode(fullReport),
		})
	}
	v3Response := MercuryV03Response{
		Reports: reports,
	}
	jsonified, err := json.Marshal(v3Response)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("failed to marshal response: %s", err)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(jsonified)
	w.Header().Add("Content-Type", "application/json")
}

func generateReport(
	feedID, configDigest [32]byte,
	obsTs uint32,
	expiresTs uint32,
	price *big.Int,
	ocr2Keys []ocr2key.KeyBundle,
	f int) ([]byte, error) {
	r, err := EncodeReportDataV3(feedID, obsTs, obsTs, big.NewInt(0), big.NewInt(0),
		expiresTs, price, price, price)
	if err != nil {
		return nil, fmt.Errorf("failed to encode report data: %s", err)
	}

	rawCtx, reportCtx := createReportContext(configDigest)
	rawRs, rawSs, rawVs, err := signReport(r, ocr2Keys, reportCtx, f)
	if err != nil {
		return nil, fmt.Errorf("failed to sign report: %s", err)
	}

	fullReport, err := EncodeFullReport(rawCtx, r, rawRs, rawSs, rawVs)
	if err != nil {
		return nil, fmt.Errorf("failed to encode full report: %s", err)
	}

	return fullReport, nil
}

func createReportContext(configDigest [32]byte) (rawCtx [3][32]byte, ctx ocrtypes.ReportContext) {
	reportContext := ocrtypes.ReportContext{
		ReportTimestamp: ocrtypes.ReportTimestamp{
			ConfigDigest: configDigest,
			Epoch:        uint32(rand.Int()),
			Round:        uint8(rand.Int()),
		},
		ExtraHash: testutils.Random32Byte(),
	}
	return evmutil.RawReportContext(reportContext), reportContext
}

func signReport(r []byte, kbs []ocr2key.KeyBundle, reportCtx ocrtypes.ReportContext, f int) (rawRs [][32]byte, rawSs [][32]byte, rawVs [32]byte, err error) {
	var signedCount int
	for i, kb := range kbs {
		if signedCount == (f + 1) {
			break
		}
		var sig []byte
		sig, err = kb.Sign(reportCtx, r)
		if err != nil {
			return
		}
		if len(sig) != 65 {
			err = fmt.Errorf("expected sig length 65, got %d", len(sig))
			return
		}
		var (
			r, s [32]byte
		)
		copy(r[:], sig[:32])
		copy(s[:], sig[32:64])
		rawVs[i] = sig[64]
		rawRs = append(rawRs, r)
		rawSs = append(rawSs, s)
		signedCount++
	}
	return
}
