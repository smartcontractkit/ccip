package usdc

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/tokendata"
)

type IHttpClient interface {
	// Get issue a GET request to the given url and return the response body.
	Get(ctx context.Context, url string) ([]byte, error)
	GetWithTimeout(ctx context.Context, url string, timeout time.Duration) ([]byte, error)
}

type HttpClient struct {
}

func (s *HttpClient) Get(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

func (s *HttpClient) GetWithTimeout(ctx context.Context, url string, timeout time.Duration) ([]byte, error) {
	// Use a timeout to guard against attestation API hanging, causing observation timeout and failing to make any progress.
	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(timeoutCtx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, tokendata.ErrTimeout
		}
		return nil, err
	}
	defer res.Body.Close()

	// Explicitly signal if the API is being rate limited
	if res.StatusCode == http.StatusTooManyRequests {
		return nil, tokendata.ErrRateLimit
	}

	return io.ReadAll(res.Body)
}
