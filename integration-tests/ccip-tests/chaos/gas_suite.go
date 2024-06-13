package chaos

import (
	"context"
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"

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
	SrcRPCClient  *client.RPCClient
	DstRPCClient  *client.RPCClient
	SrcEVMClient  *ethclient.Client
	DstEVMClient  *ethclient.Client
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
	srcEVMClient, err := ethclient.Dial(cfg.SrcGethHTTPURL)
	if err != nil {
		return nil, err
	}
	dstEVMClient, err := ethclient.Dial(cfg.DstGethHTTPURL)
	if err != nil {
		return nil, err
	}
	return &GasSuite{
		t:             t,
		Cfg:           cfg,
		Logger:        l,
		SrcRPCClient:  client.NewRPCClient(cfg.SrcGethHTTPURL),
		DstRPCClient:  client.NewRPCClient(cfg.DstGethHTTPURL),
		SrcEVMClient:  srcEVMClient,
		DstEVMClient:  dstEVMClient,
		GrafanaClient: grafana.NewGrafanaClient(cfg.GrafanaURL, cfg.GrafanaToken),
	}, nil
}

// ChangeBlockGasBaseFee simulates slow or fast gas spike
func (r *GasSuite) ChangeBlockGasBaseFee(chain string, from int64, percentage float64, duration time.Duration, spike bool) {
	go func() {
		err := PostGrafanaAnnotation(
			r.Logger,
			r.GrafanaClient,
			r.Cfg.dashboardUID,
			fmt.Sprintf("gas spike started (src chain), initial price: %d, raise: %.2f", from, percentage),
			[]string{"gas-spike"},
		)
		assert.NoError(r.t, err)
		switch chain {
		case "src":
			err := r.SrcRPCClient.ModulateBaseFeeOverDuration(r.Logger, from, percentage, duration, spike)
			assert.NoError(r.t, err)
		case "dst":
			err := r.DstRPCClient.ModulateBaseFeeOverDuration(r.Logger, from, percentage, duration, spike)
			assert.NoError(r.t, err)
		default:
			r.t.Errorf("chain can be 'src' or 'dst'")
		}
		err = PostGrafanaAnnotation(
			r.Logger,
			r.GrafanaClient,
			r.Cfg.dashboardUID,
			fmt.Sprintf("gas spike ended (src chain), price: %.2f, raise: %.2f", float64(from)*percentage, percentage),
			[]string{"gas-spike"},
		)
		assert.NoError(r.t, err)
	}()
}

// ChangeNextBlockGasLimit changes next block gas limit,
// sets it to percentage of last gasUsed in previous block creating congestion
func (r *GasSuite) ChangeNextBlockGasLimit(startIn time.Duration, wait time.Duration, chain string, percentage float64) {
	go func() {
		time.Sleep(startIn)
		latestBlock, err := r.SrcEVMClient.BlockByNumber(context.Background(), nil)
		assert.NoError(r.t, err)
		newGasLimit := int64(math.Ceil(float64(latestBlock.GasUsed()) * percentage))
		r.Logger.Info().
			Str("Network", chain).
			Int64("GasLimit", newGasLimit).
			Uint64("GasUsed", latestBlock.GasUsed()).
			Msg("Setting next block gas limit")
		err = PostGrafanaAnnotation(
			r.Logger,
			r.GrafanaClient,
			r.Cfg.dashboardUID,
			fmt.Sprintf("changed block gas limit, now: %d, was used in last block: %d, network: %s", newGasLimit, latestBlock.GasUsed(), chain),
			[]string{"gas-limit"},
		)
		assert.NoError(r.t, err)
		switch chain {
		case "src":
			err := r.SrcRPCClient.AnvilSetBlockGasLimit([]interface{}{newGasLimit})
			assert.NoError(r.t, err)
		case "dst":
			err := r.DstRPCClient.AnvilSetBlockGasLimit([]interface{}{newGasLimit})
			assert.NoError(r.t, err)
		default:
			r.t.Errorf("chain can be 'src' or 'dst'")
		}
		time.Sleep(wait)
		r.Logger.Info().
			Str("Network", chain).
			Uint64("GasLimit", latestBlock.GasLimit()).
			Msg("Returning old gas limit")
		switch chain {
		case "src":
			err := r.SrcRPCClient.AnvilSetBlockGasLimit([]interface{}{latestBlock.GasLimit()})
			assert.NoError(r.t, err)
		case "dst":
			err := r.DstRPCClient.AnvilSetBlockGasLimit([]interface{}{latestBlock.GasLimit()})
			assert.NoError(r.t, err)
		default:
			r.t.Errorf("chain can be 'src' or 'dst'")
		}
		err = PostGrafanaAnnotation(
			r.Logger,
			r.GrafanaClient,
			r.Cfg.dashboardUID,
			fmt.Sprintf("changed block gas limit, now: %d, network: %s", newGasLimit, chain),
			[]string{"gas-limit"},
		)
	}()
}
