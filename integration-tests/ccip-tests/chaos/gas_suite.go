package chaos

import (
	"fmt"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-testing-framework/client"
	"github.com/smartcontractkit/chainlink-testing-framework/grafana"
	"github.com/smartcontractkit/chainlink-testing-framework/logging"
)

// GasSuite is a test suite that generates gas chaos
type GasSuite struct {
	t             *testing.T
	Cfg           *GasSuiteConfig
	Logger        zerolog.Logger
	SrcClient     *client.RPCClient
	DstClient     *client.RPCClient
	GrafanaClient *grafana.Client
}

// GasSuiteConfig is a configuration for gas chaos tests
type GasSuiteConfig struct {
	// SrcGethHTTPURL source chain Geth HTTP URL
	SrcGethHTTPURL string
	// DstGethHTTPURL dest chain Geth HTTP URL
	DstGethHTTPURL string
	// GrafanaConfig is a common Grafana config used for annotating experiments
	*GrafanaConfig
}

// Validate validates ReorgConfig params
func (rc *GasSuiteConfig) Validate() error {
	return rc.GrafanaConfig.Validate()
}

// NewGasSuite creates new gas suite with source/dest RPC clients, works only with Anvil
func NewGasSuite(t *testing.T, cfg *GasSuiteConfig) (*GasSuite, error) {
	l := logging.GetTestLogger(t).With().Str("Component", "reorg").Logger()
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return &GasSuite{
		Cfg:           cfg,
		Logger:        l,
		SrcClient:     client.NewRPCClient(cfg.SrcGethHTTPURL),
		DstClient:     client.NewRPCClient(cfg.DstGethHTTPURL),
		GrafanaClient: grafana.NewGrafanaClient(cfg.GrafanaURL, cfg.GrafanaToken),
	}, nil
}

// RaiseGas simulates slow or fast gas spike
func (r *GasSuite) RaiseGas(chain string, from int64, percentage float64, duration time.Duration, spike bool) {
	go func() {
		err := PostGrafanaAnnotation(
			r.Logger,
			r.GrafanaClient,
			r.Cfg.dashboardUID,
			fmt.Sprintf("gas spike started (src chain), price: %d", from),
			[]string{"gas-spike"},
		)
		assert.NoError(r.t, err)
		switch chain {
		case "src":
			err := r.SrcClient.ModulateBaseFeeOverDuration(r.Logger, from, percentage, duration, spike)
			assert.NoError(r.t, err)
		case "dst":
			err := r.DstClient.ModulateBaseFeeOverDuration(r.Logger, from, percentage, duration, spike)
			assert.NoError(r.t, err)
		default:
			r.t.Errorf("chain can be 'src' or 'dst'")
		}
		err = PostGrafanaAnnotation(
			r.Logger,
			r.GrafanaClient,
			r.Cfg.dashboardUID,
			fmt.Sprintf("gas spike ended (src chain), price: %.2f", float64(from)*percentage),
			[]string{"gas-spike"},
		)
		assert.NoError(r.t, err)
	}()
}
