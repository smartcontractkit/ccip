package merclib

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func GenerateHMAC(method, path string, body []byte, clientId string, timestamp int64, userSecret, caller string) string {
	serverBodyHash := sha256.New()
	serverBodyHash.Write(body)
	serverBodyHashString := fmt.Sprintf("%s %s %s %s %d",
		method,
		path,
		hex.EncodeToString(serverBodyHash.Sum(nil)),
		clientId,
		timestamp)
	signedMessage := hmac.New(sha256.New, []byte(userSecret))
	signedMessage.Write([]byte(serverBodyHashString))
	userHmac := hex.EncodeToString(signedMessage.Sum(nil))
	return userHmac
}

func GenerateAuthHeaders(method, pathAndParams, clientId, userSecret string) http.Header {
	header := http.Header{}
	timestamp := time.Now().UTC().UnixMilli()
	hmacString := GenerateHMAC(method, pathAndParams, []byte(""), clientId, timestamp, userSecret, "merclib")

	header.Add("Authorization", clientId)
	header.Add("X-Authorization-Timestamp", strconv.FormatInt(timestamp, 10))
	header.Add("X-Authorization-Signature-SHA256", hmacString)
	// TODO: check if we should have another header for CCIP specifically, similar to automation
	return header
}

func GenerateBatchRequest(ctx context.Context, feedIDs []string, mercuryURL, clientId, sharedSecret string) (*http.Request, error) {
	tsNow := time.Now().UTC().UnixMilli()
	urlValues := url.Values{
		"feedIDs":   {strings.Join(feedIDs, ",")},
		"timestamp": {strconv.FormatInt(tsNow, 10)},
	}
	u, err := url.Parse(mercuryURL)
	if err != nil {
		return nil, fmt.Errorf("parsing mercury URL: %w", err)
	}
	u.Path = MercuryBatchPath
	u.RawQuery = urlValues.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header = GenerateAuthHeaders(http.MethodGet, req.URL.RequestURI(), clientId, sharedSecret)
	return req, nil
}
