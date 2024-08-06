package chaos

import (
	"fmt"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-testing-framework/client"
	"github.com/smartcontractkit/chainlink-testing-framework/grafana"
)

// ReorgSuite is a test suite that generates reorgs on source/dest chains
type ReorgSuite struct {
	t             *testing.T
	Cfg           *ReorgConfig
	Logger        *zerolog.Logger
	SrcClient     *client.RPCClient
	DstClient     *client.RPCClient
	GrafanaClient *grafana.Client
}

// ReorgConfig is a configuration for reorg tests
type ReorgConfig struct {
	// SrcGethHTTPURL source chain Geth HTTP URL
	SrcGethHTTPURL string
	// DstGethHTTPURL dest chain Geth HTTP URL
	DstGethHTTPURL string
	// SrcFinalityDepth source chain finality depth
	SrcFinalityDepth uint64
	// DstGethHTTPURL dest chain finality depth
	DstFinalityDepth uint64
	// FinalityDelta blocks to rewind below or above finality
	FinalityDelta int
	// ExperimentDuration experiment duration
	ExperimentDuration time.Duration
	// GrafanaConfig is common Grafana config
	*GrafanaConfig
}

// Validate validates ReorgConfig params
func (rc *ReorgConfig) Validate() error {
	if rc.FinalityDelta >= int(rc.SrcFinalityDepth) || rc.FinalityDelta >= int(rc.DstFinalityDepth) {
		return fmt.Errorf(
			"finality delta can't be higher than source or dest chain finality, delta: %d, src: %d, dst: %d",
			rc.FinalityDelta, rc.SrcFinalityDepth, rc.DstFinalityDepth,
		)
	}
	return rc.GrafanaConfig.Validate()
}

// NewReorgSuite creates new reorg suite with source/dest RPC clients, works only with Geth
func NewReorgSuite(t *testing.T, lggr *zerolog.Logger, cfg *ReorgConfig) (*ReorgSuite, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return &ReorgSuite{
		t:             t,
		Cfg:           cfg,
		Logger:        lggr,
		SrcClient:     client.NewRPCClient(cfg.SrcGethHTTPURL),
		DstClient:     client.NewRPCClient(cfg.DstGethHTTPURL),
		GrafanaClient: grafana.NewGrafanaClient(cfg.GrafanaURL, cfg.GrafanaToken),
	}, nil
}

// RunReorg rollbacks given chain, for N blocks back
func (r *ReorgSuite) RunReorg(client *client.RPCClient, blocksBack int, network string, startDelay time.Duration) {
	go func() {
		time.Sleep(startDelay)
		r.Logger.Info().
			Str("Network", network).
			Str("URL", client.URL).
			Int("BlocksBack", blocksBack).
			Msg(fmt.Sprintf("Rewinding blocks on %s chain", network))
		blockNumber, err := client.BlockNumber()
		assert.NoError(r.t, err)
		r.Logger.Info().
			Int64("Number", blockNumber).
			Str("Network", network).
			Msg("Block number before rewinding")
		err = client.GethSetHead(blocksBack)
		assert.NoError(r.t, err)
		blockNumber, err = client.BlockNumber()
		assert.NoError(r.t, err)
		r.Logger.Info().
			Int64("Number", blockNumber).
			Str("Network", network).
			Msg("Block number after rewinding")
		err = PostGrafanaAnnotation(
			r.Logger,
			r.GrafanaClient,
			r.Cfg.dashboardUID,
			fmt.Sprintf("rewinded %s chain for %d blocks back, finality in source is: %d and finality in destination is: %d",
				network, blocksBack, r.Cfg.SrcFinalityDepth, r.Cfg.SrcFinalityDepth),
			nil,
		)
		assert.NoError(r.t, err)
	}()
}
