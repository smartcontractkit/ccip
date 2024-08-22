package managed

import (
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/errors"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/executable"
)

// BaseMCMSProposal is the base struct for all MCMS proposals
// Note: this type should never be utilized directly which is why it is private
type baseMCMSProposal struct {
	executable.ExecutableMCMSProposal

	// This is intended to be displayed as-is to signers, to give them
	// context for the change. File authors should templatize strings for
	// this purpose in their pipelines.
	Description string `json:"description"`

	// Operations to be executed
	Transactions []ChainOperation `json:"transactions"`
}

func (m *baseMCMSProposal) Validate() error {
	if err := m.ExecutableMCMSProposal.ExecutableMCMSProposalBase.Validate(); err != nil {
		return err
	}

	if m.Description == "" {
		return &errors.ErrInvalidDescription{
			ReceivedDescription: m.Description,
		}
	}

	// Validate all chains in transactions have an entry in chain metadata
	for _, t := range m.Transactions {
		if _, ok := m.ChainMetadata[t.GetChainIdentifier()]; !ok {
			return &errors.ErrMissingChainDetails{
				ChainIdentifier: t.GetChainIdentifier(),
				Parameter:       "chain metadata",
			}
		}
	}
	return nil
}

func (m *baseMCMSProposal) AddSignature(sig executable.Signature) {
	m.Signatures = append(m.Signatures, sig)
}

func (m *baseMCMSProposal) ToExecutableMCMSProposal() executable.ExecutableMCMSProposal {
	raw := executable.ExecutableMCMSProposal{
		ExecutableMCMSProposalBase: m.ExecutableMCMSProposalBase,
		Transactions:               make([]executable.ChainOperation, 0),
	}

	for k, v := range m.ChainMetadata {
		raw.ChainMetadata[k] = executable.ExecutableMCMSChainMetadata{
			NonceOffset: v.NonceOffset,
			MCMAddress:  v.MCMAddress,
		}
	}

	return raw
}
