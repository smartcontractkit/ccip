package persistent

import (
	"fmt"
	"github.com/AlekSi/pointer"
	"github.com/smartcontractkit/ccip/integration-tests/deployment"
	ctfClient "github.com/smartcontractkit/chainlink-testing-framework/client"
	ctftestenv "github.com/smartcontractkit/chainlink-testing-framework/docker/test_env"
)

func NewMocks(config DONConfig) (*deployment.Mocks, error) {
	if config.NewDON != nil {
		return &deployment.Mocks{
			KillGrave: ctftestenv.NewKillgrave(config.NewDON.Options.Networks, "", ctftestenv.WithLogStream(config.NewDON.Options.LogStream)),
		}, nil
	}

	mockserverURL := pointer.GetString(config.ExistingDON.MockServerURL)
	if mockserverURL == "" {
		return nil, fmt.Errorf("mockserver URL is required for existing DON")
	}
	return &deployment.Mocks{
		MockServer: ctfClient.NewMockserverClient(&ctfClient.MockserverConfig{
			LocalURL:   mockserverURL,
			ClusterURL: mockserverURL,
		}),
	}, nil
}
