package deployment

import (
	"github.com/smartcontractkit/chainlink/integration-tests/web/sdk/client"
)

type DON struct {
	Bootstrap client.Client
	Nodes     []client.Client
}

type DonConfig struct {
	Bootstrap client.Credentials
}
