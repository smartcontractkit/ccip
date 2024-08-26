package persistent

import (
	"fmt"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"strings"
	"testing"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment"
	tc "github.com/smartcontractkit/chainlink/integration-tests/testconfig"
	"github.com/stretchr/testify/require"
)

func NewPersistentEnvironment(t *testing.T, lggr logger.Logger) deployment.Environment {
	config, err := tc.GetConfig([]string{"Smoke"}, tc.OCR)
	require.NoError(t, err, "Error getting config")

	chains, err := NewPersistentChains(t, config)
	require.NoError(t, err, "Error getting chain")

	return deployment.Environment{
		Name: fmt.Sprintf("eth-persistent-%s", strings.Join(func() []string {
			var chainIds []string
			for id := range chains {
				chainIds = append(chainIds, fmt.Sprintf("%d", id))
			}

			return chainIds
		}(), "-")),
		Chains:  chains,
		NodeIDs: []string{}, //TODO add nodes!
		Logger:  lggr,
	}
}
