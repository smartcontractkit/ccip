package deployment

import (
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/managed"
)

// Services as input to CI/Async tasks
type ChangesetOutput struct {
	JobSpecs    map[string][]string
	Proposals   []managed.MCMSWithTimelockProposal
	AddressBook AddressBook
}
