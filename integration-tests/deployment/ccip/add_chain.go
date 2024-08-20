package ccipdeployment

import (
	"github.com/smartcontractkit/ccip/integration-tests/deployment"

	"github.com/smartcontractkit/chainlink-testing-framework/k8s/environment"
)

func AddChain(
	e environment.Environment,
	ab deployment.AddressBook,
	chainSel uint64,
) error {
	// Enable inbound to the new chain
}
