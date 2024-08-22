package managed

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment/executable"
)

func TestMCMSOnlyProposal_ToExecutableMCMSProposal(t *testing.T) {
	proposal := MCMSOnlyProposal{
		baseMCMSProposal: baseMCMSProposal{
			ExecutableMCMSProposal: executable.ExecutableMCMSProposal{
				ExecutableMCMSProposalBase: executable.ExecutableMCMSProposalBase{
					Version:    "1.0.0",
					ValidUntil: 2004259681,
					Signatures: []executable.Signature{},
					ChainMetadata: map[string]executable.ExecutableMCMSChainMetadata{
						TestChain: {
							NonceOffset: 1,
							MCMAddress:  TestAddress,
						},
					},
				},
			},
			Description: "Sample description",
		},
		Transactions: []DetailedChainOperation{
			{
				ChainIdentifier: TestChain,
				DetailedOperation: DetailedOperation{
					ChainOperationDetails: ChainOperationDetails{
						ContractType: "Sample contract",
						Tags:         []string{"tag1", "tag2"},
					},
					Operation: executable.Operation{
						To:    TestAddress,
						Value: 0,
						Data:  "0x",
					},
				},
			},
		},
	}

	expectedProposal := executable.ExecutableMCMSProposal{
		ExecutableMCMSProposalBase: executable.ExecutableMCMSProposalBase{
			Version:    "1.0.0",
			ValidUntil: 2004259681,
			Signatures: []executable.Signature{},
			ChainMetadata: map[string]executable.ExecutableMCMSChainMetadata{
				TestChain: {
					NonceOffset: 1,
					MCMAddress:  TestAddress,
				},
			},
		},
		Transactions: []executable.ChainOperation{
			{
				ChainIdentifier: TestChain,
				Operation: executable.Operation{
					To:    TestAddress,
					Value: 0,
					Data:  "0x",
				},
			},
		},
	}

	result, err := proposal.ToExecutableMCMSProposal()
	assert.NoError(t, err)
	assert.Equal(t, expectedProposal, result)
}
