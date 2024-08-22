package executable

import (
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment/errors"
)

type ExecutableMCMSProposalBase struct {
	Version              string      `json:"version"`
	ValidUntil           uint32      `json:"validUntil"`
	Signatures           []Signature `json:"signatures"`
	OverridePreviousRoot bool        `json:"overridePreviousRoot"`

	// Map of chain identifier to chain metadata
	ChainMetadata map[string]ExecutableMCMSChainMetadata `json:"chainMetadata"`
}

type ExecutableMCMSChainMetadata struct {
	NonceOffset uint64         `json:"nonceOffset"`
	MCMAddress  common.Address `json:"mcmAddress"`
}

func (m ExecutableMCMSProposalBase) Validate() error {
	if m.Version == "" {
		return &errors.ErrInvalidVersion{
			ReceivedVersion: m.Version,
		}
	}

	// Get the current Unix timestamp as an int64
	currentTime := time.Now().Unix()

	if m.ValidUntil <= uint32(currentTime) {
		// ValidUntil is a Unix timestamp, so it should be greater than the current time
		return &errors.ErrInvalidValidUntil{
			ReceivedValidUntil: m.ValidUntil,
		}
	}

	if len(m.ChainMetadata) == 0 {
		return &errors.ErrNoChainMetadata{}
	}

	return nil
}

type ExecutableMCMSProposal struct {
	ExecutableMCMSProposalBase

	// Operations to be executed
	Transactions []ChainOperation `json:"transactions"`
}

func (m *ExecutableMCMSProposal) Validate() error {
	if err := m.ExecutableMCMSProposalBase.Validate(); err != nil {
		return err
	}

	if len(m.Transactions) == 0 {
		return &errors.ErrNoTransactions{}
	}

	// Validate all chains in transactions have an entry in chain metadata
	for _, t := range m.Transactions {
		if _, ok := m.ChainMetadata[t.ChainIdentifier]; !ok {
			return &errors.ErrMissingChainDetails{
				ChainIdentifier: t.ChainIdentifier,
				Parameter:       "chain metadata",
			}
		}
	}

	return nil
}

func (m *ExecutableMCMSProposal) ToExecutor(clients map[string]ContractDeployBackend) (*Executor, error) {
	// Create a new executor
	executor, err := NewProposalExecutor(m, clients)
	if err != nil {
		return nil, err
	}

	return executor, nil
}
