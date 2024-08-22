package managed

import "github.com/smartcontractkit/chainlink/integration-tests/deployment/executable"

type MCMSOnlyChainMetadata struct {
	executable.ExecutableMCMSChainMetadata
}

// MCMSOnlyProposal is a struct where the target contract is an MCMS contract
// with no forwarder contracts. This type does not support any type of atomic contract
// call batching, as the MCMS contract natively doesn't support batching
type MCMSOnlyProposal struct {
	baseMCMSProposal

	// Operations to be executed
	Transactions []DetailedChainOperation `json:"transactions"`
}

func (m *MCMSOnlyProposal) ToExecutableMCMSProposal() (executable.ExecutableMCMSProposal, error) {
	raw := m.baseMCMSProposal.ToExecutableMCMSProposal()

	for _, t := range m.Transactions {
		raw.Transactions = append(raw.Transactions, executable.ChainOperation{
			ChainIdentifier: t.ChainIdentifier,
			Operation:       t.Operation,
		})
	}

	return raw, nil
}
