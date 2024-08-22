package executable

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var TestAddress = common.HexToAddress("0x1234567890abcdef")

func TestExecutableMCMSProposal_Validate_Success(t *testing.T) {
	proposal := &ExecutableMCMSProposal{
		ExecutableMCMSProposalBase: ExecutableMCMSProposalBase{
			Version:              "1.0",
			ValidUntil:           2004259681,
			Signatures:           []Signature{},
			OverridePreviousRoot: false,
			ChainMetadata: map[string]ExecutableMCMSChainMetadata{
				"chain1": {
					NonceOffset: 1,
					MCMAddress:  TestAddress,
				},
			},
		},
		Transactions: []ChainOperation{
			{
				ChainIdentifier: "chain1",
				Operation: Operation{
					To:    TestAddress,
					Value: 0,
					Data:  "0x",
				},
			},
		},
	}

	err := proposal.Validate()

	assert.NoError(t, err)
}

func TestExecutableMCMSProposal_Validate_InvalidVersion(t *testing.T) {
	proposal := &ExecutableMCMSProposal{
		ExecutableMCMSProposalBase: ExecutableMCMSProposalBase{
			Version:              "",
			ValidUntil:           2004259681,
			Signatures:           []Signature{},
			OverridePreviousRoot: false,
			ChainMetadata: map[string]ExecutableMCMSChainMetadata{
				"chain1": {
					NonceOffset: 1,
					MCMAddress:  TestAddress,
				},
			},
		},
		Transactions: []ChainOperation{
			{
				ChainIdentifier: "chain1",
				Operation: Operation{
					To:    TestAddress,
					Value: 0,
					Data:  "0x",
				},
			},
		},
	}

	err := proposal.Validate()

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "invalid version: ")
}

func TestExecutableMCMSProposal_Validate_InvalidValidUntil(t *testing.T) {
	proposal := &ExecutableMCMSProposal{
		ExecutableMCMSProposalBase: ExecutableMCMSProposalBase{
			Version:              "1.0",
			ValidUntil:           0,
			Signatures:           []Signature{},
			OverridePreviousRoot: false,
			ChainMetadata: map[string]ExecutableMCMSChainMetadata{
				"chain1": {
					NonceOffset: 1,
					MCMAddress:  TestAddress,
				},
			},
		},
		Transactions: []ChainOperation{
			{
				ChainIdentifier: "chain1",
				Operation: Operation{
					To:    TestAddress,
					Value: 0,
					Data:  "0x",
				},
			},
		},
	}

	err := proposal.Validate()

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "invalid valid until: 0")
}

func TestExecutableMCMSProposal_Validate_InvalidChainMetadata(t *testing.T) {
	proposal := &ExecutableMCMSProposal{
		ExecutableMCMSProposalBase: ExecutableMCMSProposalBase{
			Version:              "1.0",
			ValidUntil:           2004259681,
			Signatures:           []Signature{},
			OverridePreviousRoot: false,
			ChainMetadata:        map[string]ExecutableMCMSChainMetadata{},
		},
		Transactions: []ChainOperation{
			{
				ChainIdentifier: "chain1",
				Operation: Operation{
					To:    TestAddress,
					Value: 0,
					Data:  "0x",
				},
			},
		},
	}

	err := proposal.Validate()

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "no chain metadata")
}

func TestExecutableMCMSProposal_Validate_NoTransactions(t *testing.T) {
	proposal := &ExecutableMCMSProposal{
		ExecutableMCMSProposalBase: ExecutableMCMSProposalBase{
			Version:              "1.0",
			ValidUntil:           2004259681,
			Signatures:           []Signature{},
			OverridePreviousRoot: false,
			ChainMetadata: map[string]ExecutableMCMSChainMetadata{
				"chain1": {
					NonceOffset: 1,
					MCMAddress:  TestAddress,
				},
			},
		},
		Transactions: []ChainOperation{},
	}

	err := proposal.Validate()

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "no transactions")
}

func TestExecutableMCMSProposal_Validate_MissingChainMetadataForTransaction(t *testing.T) {
	proposal := &ExecutableMCMSProposal{
		ExecutableMCMSProposalBase: ExecutableMCMSProposalBase{
			Version:              "1.0",
			ValidUntil:           2004259681,
			Signatures:           []Signature{},
			OverridePreviousRoot: false,
			ChainMetadata: map[string]ExecutableMCMSChainMetadata{
				"chain1": {
					NonceOffset: 1,
					MCMAddress:  TestAddress,
				},
			},
		},
		Transactions: []ChainOperation{
			{
				ChainIdentifier: "chain2",
				Operation: Operation{
					To:    TestAddress,
					Value: 0,
					Data:  "0x",
				},
			},
		},
	}

	err := proposal.Validate()

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "missing chain metadata for chain chain2")
}
