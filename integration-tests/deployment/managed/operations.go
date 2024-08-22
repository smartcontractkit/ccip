package managed

import (
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/executable"
)

type ChainOperationDetails struct {
	ContractType string   `json:"contractType"`
	Tags         []string `json:"tags"`
}

type DetailedOperation struct {
	executable.Operation
	ChainOperationDetails
}

type ChainOperation interface {
	GetChainIdentifier() string
}

type DetailedChainOperation struct {
	ChainIdentifier string `json:"chainIdentifier"`
	DetailedOperation
}

func (m DetailedChainOperation) GetChainIdentifier() string {
	return m.ChainIdentifier
}

type DetailedBatchChainOperation struct {
	ChainIdentifier string              `json:"chainIdentifier"`
	Batch           []DetailedOperation `json:"batch"`
}

func (m DetailedBatchChainOperation) GetChainIdentifier() string {
	return m.ChainIdentifier
}
