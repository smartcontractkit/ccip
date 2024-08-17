package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Khan/genqlient/graphql"

	"github.com/smartcontractkit/chainlink/integration-tests/web/sdk/client/internal/doer"
	"github.com/smartcontractkit/chainlink/integration-tests/web/sdk/internal/generated"
)

type Client struct {
	gqlClient   graphql.Client
	credentials Credentials
	endpoints   endpoints
	cookie      string
}

type endpoints struct {
	Sessions string
	Query    string
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Config struct {
	BaseURI string `json:"baseURI"`
	Credentials
}

func New(cfg Config) (*Client, error) { //nolint:revive
	endpoints := endpoints{
		Sessions: cfg.BaseURI + "/sessions",
		Query:    cfg.BaseURI + "/query",
	}
	c := &Client{
		endpoints:   endpoints,
		credentials: cfg.Credentials,
	}

	if err := c.login(); err != nil {
		return nil, fmt.Errorf("failed to login to node: %w", err)
	}

	c.gqlClient = graphql.NewClient(
		c.endpoints.Query,
		doer.NewAuthed(c.cookie),
	)

	return c, nil
}

func (c *Client) GetCSAKeys(ctx context.Context) (*generated.GetCSAKeysResponse, error) {
	return generated.GetCSAKeys(ctx, c.gqlClient)
}

func (c *Client) GetJob(ctx context.Context, id string) (*generated.GetJobResponse, error) {
	return generated.GetJob(ctx, c.gqlClient, id)
}

func (c *Client) ListJobs(ctx context.Context, offset, limit int) (*generated.ListJobsResponse, error) {
	return generated.ListJobs(ctx, c.gqlClient, offset, limit)
}

func (c *Client) GetJobProposal(ctx context.Context, id string) (*generated.GetJobProposalResponse, error) {
	return generated.GetJobProposal(ctx, c.gqlClient, id)
}

func (c *Client) GetBridge(ctx context.Context, id string) (*generated.GetBridgeResponse, error) {
	return generated.GetBridge(ctx, c.gqlClient, id)
}

func (c *Client) ListBridges(ctx context.Context, offset, limit int) (*generated.ListBridgesResponse, error) {
	return generated.ListBridges(ctx, c.gqlClient, offset, limit)
}

func (c *Client) GetFeedsManager(ctx context.Context, id string) (*generated.GetFeedsManagerResponse, error) {
	return generated.GetFeedsManager(ctx, c.gqlClient, id)
}

func (c *Client) ListFeedsManagers(ctx context.Context) (*generated.ListFeedsManagersResponse, error) {
	return generated.ListFeedsManagers(ctx, c.gqlClient)
}

func (c *Client) CreateFeedsManager(ctx context.Context, cmd generated.CreateFeedsManagerInput) (*generated.CreateFeedsManagerResponse, error) {
	return generated.CreateFeedsManager(ctx, c.gqlClient, cmd)
}

func (c *Client) login() error {
	b, err := json.Marshal(c.credentials)
	if err != nil {
		return fmt.Errorf("failed to marshal credentials: %w", err)
	}

	payload := strings.NewReader(string(b))

	req, err := http.NewRequest("POST", c.endpoints.Sessions, payload)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	cookieHeader := res.Header.Get("Set-Cookie")
	if cookieHeader == "" {
		return fmt.Errorf("no cookie found in header")
	}

	c.cookie = strings.Split(cookieHeader, ";")[0]
	return nil
}
