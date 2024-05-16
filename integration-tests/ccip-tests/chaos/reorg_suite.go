package chaos

import (
	"fmt"
	"strings"
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
	// GrafanaURL Grafana URL
	GrafanaURL string
	// GrafanaToken Grafana API token
	GrafanaToken string
	// DashboardURL dashboard URL in format "/d/6vjVx-1V8/ccip-long-running-tests"
	DashboardURL string
	// dashboardUID part of DashboardURL to put annotation on
	dashboardUID string
}

// Validate validates ReorgConfig params
func (rc *ReorgConfig) Validate() error {
	if rc.FinalityDelta >= int(rc.SrcFinalityDepth) || rc.FinalityDelta >= int(rc.DstFinalityDepth) {
		return fmt.Errorf(
			"finality delta can't be higher than source or dest chain finality, delta: %d, src: %d, dst: %d",
			rc.FinalityDelta, rc.SrcFinalityDepth, rc.DstFinalityDepth,
		)
	}
	urlParams := strings.Split(rc.DashboardURL, "/")
	if len(urlParams) != 4 {
		return fmt.Errorf("invalid Grafana dashboard URL format, must be: /d/6vjVx-1V8/ccip-long-running-tests")
	}
	rc.dashboardUID = urlParams[2]
	return nil
}

// NewReorgSuite creates new reorg suite with source/dest RPC clients, works only with Geth
func NewReorgSuite(t *testing.T, cfg *ReorgConfig) (*ReorgSuite, error) {
	l := logging.GetTestLogger(t).With().Str("Component", "reorg").Logger()
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return &ReorgSuite{
		Cfg:           cfg,
		Logger:        l,
		SrcClient:     client.NewRPCClient(cfg.SrcGethHTTPURL),
		DstClient:     client.NewRPCClient(cfg.DstGethHTTPURL),
		GrafanaClient: grafana.NewGrafanaClient(cfg.GrafanaURL, cfg.GrafanaToken),
	}, nil
}

// annotate sets dashboard annotation about when block was rewinding on which chain
func (r *ReorgSuite) annotate(text string) {
	res, _, err := r.GrafanaClient.PostAnnotation(grafana.PostAnnotation{
		DashboardUID: r.Cfg.dashboardUID,
		Tags:         []string{"reorg", "rewind_head"},
		Text:         fmt.Sprintf("<pre>%s</pre>", text),
	})
	r.Logger.Warn().Str("DashboardUID", r.Cfg.dashboardUID).Any("ResponseBody", res).Msg("Annotated experiment")
	assert.NoError(r.t, err)
}

// RunReorgBelowFinalityThreshold we rollback both chains, one by one, for N blocks back
// no assertions needed, load test should fail if something went wrong
func (r *ReorgSuite) RunReorgBelowFinalityThreshold(startDelay time.Duration) {
	go func() {
		time.Sleep(startDelay)
		blocksBackSrc := int(r.Cfg.SrcFinalityDepth) - r.Cfg.FinalityDelta
		r.Logger.Warn().
			Str("URL", r.SrcClient.URL).
			Str("Case", "below finality").
			Int("BlocksBack", blocksBackSrc).
			Msg("Rewinding blocks on src chain")
		err := r.SrcClient.GethSetHead(blocksBackSrc)
		assert.NoError(r.t, err)
		r.annotate(fmt.Sprintf("rewinded source chain for %d blocks back, finality is: %d", blocksBackSrc, r.Cfg.SrcFinalityDepth))
		time.Sleep(r.Cfg.ExperimentDuration)

		blocksBackDst := int(r.Cfg.DstFinalityDepth) - r.Cfg.FinalityDelta
		r.Logger.Warn().
			Str("URL", r.SrcClient.URL).
			Str("Case", "below finality").
			Int("BlocksBack", blocksBackDst).
			Msg("Rewinding blocks on dst chain")
		err = r.DstClient.GethSetHead(blocksBackDst)
		assert.NoError(r.t, err)
		r.annotate(fmt.Sprintf("rewinded dest chain for %d blocks back, finality is: %d", blocksBackDst, r.Cfg.DstFinalityDepth))
		time.Sleep(r.Cfg.ExperimentDuration)
	}()
}

// RunReorgAboveFinalityThreshold we rollback both chains, one by one, for N blocks back, above threshold
// asserting there is no messages passing and finality violation is detected
func (r *ReorgSuite) RunReorgAboveFinalityThreshold(startDelay time.Duration) {
	go func() {
		time.Sleep(startDelay)
		blocksBackSrc := int(r.Cfg.SrcFinalityDepth) + r.Cfg.FinalityDelta
		r.Logger.Warn().
			Str("URL", r.SrcClient.URL).
			Str("Case", "above finality").
			Int("BlocksBack", blocksBackSrc).
			Msg("Rewinding blocks on dst chain")
		err := r.SrcClient.GethSetHead(blocksBackSrc)
		assert.NoError(r.t, err)
		r.annotate(fmt.Sprintf("rewinded source chain for %d blocks back, finality is: %d", blocksBackSrc, r.Cfg.SrcFinalityDepth))
		// TODO: assert the interval, no messages should be processed
		time.Sleep(r.Cfg.ExperimentDuration)

		blocksBackDst := int(r.Cfg.DstFinalityDepth) + r.Cfg.FinalityDelta
		r.Logger.Warn().
			Str("URL", r.SrcClient.URL).
			Str("Case", "above finality").
			Int("BlocksBack", blocksBackDst).
			Msg("Rewinding blocks on dst chain")
		err = r.DstClient.GethSetHead(blocksBackDst)
		assert.NoError(r.t, err)
		r.annotate(fmt.Sprintf("rewinded dest chain for %d blocks back, finality is: %d", blocksBackDst, r.Cfg.DstFinalityDepth))
		// TODO: assert the interval, no messages should be processed
		time.Sleep(r.Cfg.ExperimentDuration)
	}()
}
