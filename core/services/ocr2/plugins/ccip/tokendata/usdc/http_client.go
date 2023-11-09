package usdc

import (
	"context"
	"io"
	"net/http"
)

type IHttpClient interface {
	// Get issue a GET request to the given url and return the response body.
	Get(ctx context.Context, url string) ([]byte, error)
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
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
