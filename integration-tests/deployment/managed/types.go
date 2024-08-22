package managed

import "github.com/smartcontractkit/chainlink/integration-tests/deployment/executable"

type MCMSProposal interface {
	// All Managed MCMSProposal types should implement this function to ensure
	// it can be converted to an executable set of transactions in a way that can
	// be validated
	ToExecutableMCMSProposal() (executable.ExecutableMCMSProposal, error)

	// the signing step is a proposal-agnostic step that converts the proposal to
	// an executable equivalent, as a result this function although straightforward
	// needs to be included here
	AddSignature(sig executable.Signature)

	// All proposal types must implement a validate function to ensure its constructed
	// in a way that can be signed and later executed.
	Validate() error
}

type MCMSProposalType string

const (
	MCMSOnly         MCMSProposalType = "mcms-only"
	MCMSWithTimelock MCMSProposalType = "mcms-with-timelock"
)

var MCMSProposalTypeMap = map[string]MCMSProposalType{
	"mcms-only":          MCMSOnly,
	"mcms-with-timelock": MCMSWithTimelock,
}
