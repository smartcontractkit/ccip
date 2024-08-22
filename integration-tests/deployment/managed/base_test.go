package managed

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment/executable"
)

var TestAddress = common.HexToAddress("0x1234567890abcdef")
var TestChain = "chain1"

func TestBaseMCMSProposal_Validate(t *testing.T) {
	proposal := baseMCMSProposal{
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

		Transactions: []ChainOperation{
			DetailedChainOperation{
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

	err := proposal.Validate()
	assert.NoError(t, err)
}

func TestBaseMCMSProposal_Validate_InvalidBase(t *testing.T) {
	proposal := baseMCMSProposal{
		ExecutableMCMSProposal: executable.ExecutableMCMSProposal{
			ExecutableMCMSProposalBase: executable.ExecutableMCMSProposalBase{
				Version:    "",
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

		Transactions: []ChainOperation{
			DetailedChainOperation{
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

	err := proposal.Validate()
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "invalid version: ")
}

func TestBaseMCMSProposal_Validate_InvalidDescription(t *testing.T) {
	proposal := baseMCMSProposal{
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
		Description: "",

		Transactions: []ChainOperation{
			DetailedChainOperation{
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

	err := proposal.Validate()
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "invalid description: ")
}

func TestBaseMCMSProposal_Validate_MissingChainMetadataForTransaction(t *testing.T) {
	proposal := baseMCMSProposal{
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

		Transactions: []ChainOperation{
			DetailedChainOperation{
				ChainIdentifier: "chain2",
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

	err := proposal.Validate()
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "missing chain metadata for chain chain2")
}

func TestBaseMCMSProposal_AddSignature(t *testing.T) {
	proposal := baseMCMSProposal{
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

		Transactions: []ChainOperation{
			DetailedChainOperation{
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

	sig := executable.Signature{
		R: common.HexToHash("0x1234567890abcdef"),
		S: common.HexToHash("0x1234567890abcdef"),
		V: 27,
	}

	proposal.AddSignature(sig)
	assert.Len(t, proposal.Signatures, 1)
	assert.Equal(t, proposal.Signatures[0], sig)
}
