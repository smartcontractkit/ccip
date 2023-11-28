package merclib

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
)

// XXXTestOnlyMercuryHandler returns a http.Handler that can be used to mock
// the Mercury API.
func XXXTestOnlyMercuryHandler(t *testing.T, prices map[[32]byte]*big.Int, sharedSecret string) http.Handler {
	mux := http.NewServeMux()
	mux.Handle(MercuryBatchPath, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Log("received request to MercuryBatchPath: ", r.URL.Path, "headers:", r.Header)
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
			if _, ok := prices[feedID]; !ok {
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

		// TODO: is this necessary for testing?
		// TODO: pull out common hmac code
		//bodyHash := sha256.New()
		//bodyHash.Write(nil)
		//hashString := fmt.Sprintf("%s %s %s %s %s",
		//	http.MethodGet,
		//	r.URL.Path,
		//	hex.EncodeToString(bodyHash.Sum(nil)),
		//	username,
		//	authorizationTimestamp)
		//signedMessage := hmac.New(sha256.New, []byte(sharedSecret))
		//signedMessage.Write([]byte(hashString))
		//expectedSignature := hex.EncodeToString(signedMessage.Sum(nil))
		//if authorizationSignature != expectedSignature {
		//	w.Write([]byte(fmt.Sprintf("invalid signature: %s, expected: %s", authorizationSignature, expectedSignature)))
		//	w.WriteHeader(http.StatusUnauthorized)
		//	return
		//}

		// generate the report
		var reports []MercuryV03Report
		var (
			ts = uint32(time.Now().UTC().Unix())
		)
		for _, feedID := range feedIDs {
			r, err := EncodeReportDataV3(feedID, ts, ts, big.NewInt(0), big.NewInt(0),
				ts, prices[feedID], prices[feedID], prices[feedID])
			if err != nil {
				w.Write([]byte(fmt.Sprintf("failed to encode report data: %s", err)))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			fullReport, err := EncodeFullReport(
				randomReportContext(), r, randomSliceOf32ByteArrays(), randomSliceOf32ByteArrays(), testutils.Random32Byte())
			if err != nil {
				w.Write([]byte(fmt.Sprintf("failed to encode full report: %s", err)))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			reports = append(reports, MercuryV03Report{
				FeedID:                hexutil.Encode(feedID[:]),
				ValidFromTimestamp:    ts,
				ObservationsTimestamp: ts,
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
	}))

	return mux
}

func randomReportContext() [3][32]byte {
	return [3][32]byte{
		testutils.Random32Byte(),
		testutils.Random32Byte(),
		testutils.Random32Byte(),
	}
}

func randomSliceOf32ByteArrays() [][32]byte {
	return [][32]byte{
		testutils.Random32Byte(),
		testutils.Random32Byte(),
		testutils.Random32Byte(),
	}
}
