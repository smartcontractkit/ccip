package load

import (
	"context"
	"fmt"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-testing-framework/client"
	"github.com/smartcontractkit/chainlink-testing-framework/grafana"

	"github.com/smartcontractkit/chainlink/integration-tests/ccip-tests/testconfig"
)

// GasSuite is a test suite that generates gas chaos
type GasSuite struct {
	t             *testing.T
	Ctx           context.Context
	Cfg           *GasSuiteConfig
	Logger        zerolog.Logger
	RPCClient     *client.RPCClient
	EVMClient     *ethclient.Client
	GrafanaClient *grafana.Client
}

// GasSuiteConfig is a configuration for gas chaos tests
type GasSuiteConfig struct {
	NetworkName string
	// GethHTTPURL dest chain Geth HTTP URL
	GethHTTPURL string
	GasProfile  []*testconfig.GasProfile
	// GrafanaConfig is a common Grafana config used for annotating experiments
	*GrafanaConfig
}

// Validate validates ReorgConfig params
func (rc *GasSuiteConfig) Validate() error {
	return rc.GrafanaConfig.Validate()
}

// NewGasSuite creates new gas suite with source/dest RPC clients, works only with Anvil
func NewGasSuite(t *testing.T, l zerolog.Logger, cfg *GasSuiteConfig) (*GasSuite, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	dstEVMClient, err := ethclient.Dial(cfg.GethHTTPURL)
	if err != nil {
		return nil, err
	}
	return &GasSuite{
		t:             t,
		Cfg:           cfg,
		Logger:        l,
		RPCClient:     client.NewRPCClient(cfg.GethHTTPURL),
		EVMClient:     dstEVMClient,
		GrafanaClient: grafana.NewGrafanaClient(cfg.GrafanaURL, cfg.GrafanaToken),
	}, nil
}

func (r *GasSuite) Close() {
	r.EVMClient.Close()
}

// ModulateBaseFeeOverDuration simulates slow or fast gas spike or drop
func (r *GasSuite) ModulateBaseFeeOverDuration(ctx context.Context) {
	go func() {
		defer r.Close()
		select {
		case <-ctx.Done():
			r.Logger.Info().
				Str("Network", r.Cfg.NetworkName).
				Msg("Exiting ModulateBaseFeeOverDuration")
			return
		default:
			for _, profile := range r.Cfg.GasProfile {
				from := pointer.GetInt64(profile.StartGasPrice)
				duration := profile.Duration.Duration()
				percentage := pointer.GetFloat64(profile.DeviationPercentage)
				spike := pointer.GetBool(profile.Spike)
				startText := fmt.Sprintf("gas spike started, initial price: %d, raise: %.2f, network %s", from, percentage, r.Cfg.NetworkName)
				endText := fmt.Sprintf("gas spike ended, price: %.2f, raise: %.2f, network %s", float64(from)*percentage, percentage, r.Cfg.NetworkName)
				if !spike {
					startText = fmt.Sprintf("gas drop started, initial price: %d, drop: %.2f, network %s", from, percentage, r.Cfg.NetworkName)
					endText = fmt.Sprintf("gas drop ended, price: %.2f, drop: %.2f, network %s", float64(from)*percentage, percentage, r.Cfg.NetworkName)
				}
				err := PostGrafanaAnnotation(
					r.Logger,
					r.GrafanaClient,
					r.Cfg.dashboardUID,
					startText,
					[]string{"gas-change", fmt.Sprintf("network:%s", r.Cfg.NetworkName)},
				)
				assert.NoError(r.t, err)
				r.Logger.Info().
					Int64("Starting Base Fee", from).
					Float64("Percentage", percentage).
					Dur("Duration", duration).
					Bool("Spike", spike).
					Str("Network", r.Cfg.NetworkName).
					Msg("Modulating base fee over duration started")
				err = r.RPCClient.ModulateBaseFeeOverDuration(r.Logger, from, percentage, duration, spike)
				assert.NoError(r.t, err)
				r.Logger.Info().
					Int64("Starting Base Fee", from).
					Float64("Percentage", percentage).
					Dur("Duration", duration).
					Bool("Spike", spike).
					Str("Network", r.Cfg.NetworkName).
					Msg("Modulating base fee over duration ended")
				err = PostGrafanaAnnotation(
					r.Logger,
					r.GrafanaClient,
					r.Cfg.dashboardUID,
					endText,
					[]string{"gas-change", fmt.Sprintf("network:%s", r.Cfg.NetworkName)},
				)
				assert.NoError(r.t, err)
			}
		}
	}()
}
