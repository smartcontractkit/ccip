package chaos

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/smartcontractkit/chainlink-testing-framework/grafana"
	"strings"
)

// GrafanaConfig is a basic Grafana client configuration
type GrafanaConfig struct {
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
func (rc *GrafanaConfig) Validate() error {
	urlParams := strings.Split(rc.DashboardURL, "/")
	if len(urlParams) != 4 {
		return fmt.Errorf("invalid Grafana dashboard URL format, must be: /d/6vjVx-1V8/ccip-long-running-tests")
	}
	rc.dashboardUID = urlParams[2]
	return nil
}

// PostGrafanaAnnotation sets Grafana dashboard annotation with text and tags
func PostGrafanaAnnotation(l *zerolog.Logger, grafanaClient *grafana.Client, dashboardUID string, text string, tags []string) error {
	res, _, err := grafanaClient.PostAnnotation(grafana.PostAnnotation{
		DashboardUID: dashboardUID,
		Tags:         tags,
		Text:         fmt.Sprintf("<pre>%s</pre>", text),
	})
	l.Info().Str("DashboardUID", dashboardUID).Any("ResponseBody", res).Msg("Annotated experiment")
	return err
}
