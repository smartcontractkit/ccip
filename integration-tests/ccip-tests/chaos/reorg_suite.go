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

// ReorgSuite is a test suite that generates reorgs on source/dest chains
type ReorgSuite struct {
	t             *testing.T
	Cfg           *ReorgConfig
	Logger        zerolog.Logger
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
func NewReorgSuite(t *testing.T, cfg *ReorgConfig) (*ReorgSuite, error) {
	l := logging.GetTestLogger(t).With().Str("Component", "reorg").Logger()
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return &ReorgSuite{
		t:             t,
		Cfg:           cfg,
		Logger:        l,
		SrcClient:     client.NewRPCClient(cfg.SrcGethHTTPURL),
		DstClient:     client.NewRPCClient(cfg.DstGethHTTPURL),
		GrafanaClient: grafana.NewGrafanaClient(cfg.GrafanaURL, cfg.GrafanaToken),
	}, nil
}

// RunReorgBelowFinalityThreshold we rollback both chains, one by one, for N blocks back
// no assertions needed, load test should fail if something went wrong
func (r *ReorgSuite) RunReorgBelowFinalityThreshold(startDelay time.Duration) {
	go func() {
		time.Sleep(startDelay)
		blocksBackSrc := int(r.Cfg.SrcFinalityDepth) - r.Cfg.FinalityDelta
		r.Logger.Info().
			Str("URL", r.SrcClient.URL).
			Str("Case", "below finality").
			Int("BlocksBack", blocksBackSrc).
			Msg("Rewinding blocks on src chain")

		blockNumber, err := r.SrcClient.BlockNumber()
		assert.NoError(r.t, err)
		r.Logger.Info().
			Int64("Number", blockNumber).
			Msg("Current block number")
		err = r.SrcClient.GethSetHead(blocksBackSrc)
		assert.NoError(r.t, err)
		blockNumber, err = r.SrcClient.BlockNumber()
		assert.NoError(r.t, err)
		r.Logger.Info().
			Int64("Number", blockNumber).
			Msg("Block number after rewinding:")
		//err = PostGrafanaAnnotation(
		//	r.Logger,
		//	r.GrafanaClient,
		//	r.Cfg.dashboardUID,
		//	fmt.Sprintf("rewinded source chain for %d blocks back, finality is: %d", blocksBackSrc, r.Cfg.SrcFinalityDepth),
		//	nil,
		//)
		//require.NoError(r.t, err)

		//time.Sleep(r.Cfg.ExperimentDuration)
		time.Sleep(1 * time.Minute)
		blocksBackDest := int(r.Cfg.SrcFinalityDepth) - r.Cfg.FinalityDelta
		r.Logger.Info().
			Str("URL", r.SrcClient.URL).
			Str("Case", "below finality").
			Int("BlocksBack", blocksBackDest).
			Msg("Rewinding blocks on src chain")

		blockNumber, err = r.DstClient.BlockNumber()
		assert.NoError(r.t, err)
		r.Logger.Info().
			Int64("Number", blockNumber).
			Msg("Current block number")
		err = r.DstClient.GethSetHead(blocksBackDest)
		assert.NoError(r.t, err)
		blockNumber, err = r.DstClient.BlockNumber()
		assert.NoError(r.t, err)
		r.Logger.Info().
			Int64("Number", blockNumber).
			Msg("Block number after rewinding:")
		//err = PostGrafanaAnnotation(
		//	r.Logger,
		//	r.GrafanaClient,
		//	r.Cfg.dashboardUID,
		//	fmt.Sprintf("rewinded source chain for %d blocks back, finality is: %d", blocksBackSrc, r.Cfg.SrcFinalityDepth),
		//	nil,
		//)
		//require.NoError(r.t, err)

		//time.Sleep(r.Cfg.ExperimentDuration)
		time.Sleep(1 * time.Minute)
	}()
}

// RunReorgAboveFinalityThreshold we rollback both chains, one by one, for N blocks back, above threshold
// asserting there is no messages passing and finality violation is detected
func (r *ReorgSuite) RunReorgAboveFinalityThreshold(startDelay time.Duration) {
	go func() {
		time.Sleep(startDelay)
		//blocksBackSrc := int(r.Cfg.SrcFinalityDepth) + r.Cfg.FinalityDelta
		//r.Logger.Info().
		//	Str("URL", r.SrcClient.URL).
		//	Str("Case", "above finality").
		//	Int("BlocksBack", blocksBackSrc).
		//	Msg("Rewinding blocks on dst chain")
		//blockNumber, err := r.SrcClient.BlockNumber()
		//assert.NoError(r.t, err)
		//r.Logger.Info().
		//	Int64("Number", blockNumber).
		//	Msg("Block number before rewinding:")
		//err = r.SrcClient.GethSetHead(blocksBackSrc)
		//assert.NoError(r.t, err)
		//blockNumber, err = r.SrcClient.BlockNumber()
		//assert.NoError(r.t, err)
		//r.Logger.Info().
		//	Int64("Number", blockNumber).
		//	Msg("Block number after rewinding:")
		//err = PostGrafanaAnnotation(
		//	r.Logger,
		//	r.GrafanaClient,
		//	r.Cfg.dashboardUID,
		//	fmt.Sprintf("rewinded source chain for %d blocks back, finality is: %d", blocksBackSrc, r.Cfg.SrcFinalityDepth),
		//	nil,
		//)
		//require.NoError(r.t, err)
		//time.Sleep(1 * time.Minute)
		//time.Sleep(r.Cfg.ExperimentDuration)

		blocksBackDst := int(r.Cfg.DstFinalityDepth) + r.Cfg.FinalityDelta
		r.Logger.Info().
			Str("URL", r.DstClient.URL).
			Str("Case", "above finality").
			Int("BlocksBack", blocksBackDst).
			Msg("Rewinding blocks on dst chain")
		blockNumber, err := r.DstClient.BlockNumber()
		assert.NoError(r.t, err)
		r.Logger.Info().
			Int64("Number", blockNumber).
			Msg("Block number before rewinding:")
		err = r.DstClient.GethSetHead(blocksBackDst)
		assert.NoError(r.t, err)
		blockNumber, err = r.DstClient.BlockNumber()
		assert.NoError(r.t, err)
		r.Logger.Info().
			Int64("Number", blockNumber).
			Msg("Block number after rewinding:")
		//err = PostGrafanaAnnotation(
		//	r.Logger,
		//	r.GrafanaClient,
		//	r.Cfg.dashboardUID,
		//	fmt.Sprintf("rewinded dest chain for %d blocks back, finality is: %d", blocksBackDst, r.Cfg.DstFinalityDepth),
		//	nil,
		//)
		//require.NoError(r.t, err)
		time.Sleep(1 * time.Minute)
		//time.Sleep(r.Cfg.ExperimentDuration)
	}()
}
