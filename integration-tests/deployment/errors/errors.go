package errors

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// ErrInvalidChainID is the error for an invalid chain ID.
type ErrInvalidChainID struct {
	ReceivedChainID string
}

// Error returns the error message.
func (e *ErrInvalidChainID) Error() string {
	return fmt.Sprintf("invalid chain ID: %s", e.ReceivedChainID)
}

type ErrInvalidDescription struct {
	ReceivedDescription string
}

func (e *ErrInvalidDescription) Error() string {
	return fmt.Sprint("invalid description: ", e.ReceivedDescription)
}

// ErrInvalidMinDelay is the error for when the received min delay is invalid.
type ErrInvalidMinDelay struct {
	ReceivedMinDelay string
}

// Error returns the error message.
func (e *ErrInvalidMinDelay) Error() string {
	return fmt.Sprintf("invalid min delay: %s", e.ReceivedMinDelay)
}

type ErrInvalidOperation struct {
	Message string
}

func (e *ErrInvalidOperation) Error() string {
	return fmt.Sprintf("invalid operation: %s", e.Message)
}

// ErrInvalidProposalType is used when an invalid proposal type is received.
type ErrInvalidProposalType struct {
	ReceivedProposalType string
}

func (e *ErrInvalidProposalType) Error() string {
	return fmt.Sprintf("invalid proposal type: %s", e.ReceivedProposalType)
}

// ErrInvalidTimelockOperation is the error for an invalid timelock operation.
type ErrInvalidTimelockOperation struct {
	ReceivedTimelockOperation string
}

// Error returns the error message.
func (e *ErrInvalidTimelockOperation) Error() string {
	return fmt.Sprintf("invalid timelock operation: %s", e.ReceivedTimelockOperation)
}

type ErrInvalidValidUntil struct {
	ReceivedValidUntil uint32
}

func (e *ErrInvalidValidUntil) Error() string {
	return fmt.Sprintf("invalid valid until: %v", e.ReceivedValidUntil)
}

type ErrInvalidVersion struct {
	ReceivedVersion string
}

func (e *ErrInvalidVersion) Error() string {
	return fmt.Sprintf("invalid version: %s", e.ReceivedVersion)
}

// ErrMissingChainDetails is the error for missing chain metadata.
type ErrMissingChainDetails struct {
	Parameter       string
	ChainIdentifier string
}

// Error returns the error message.
func (e *ErrMissingChainDetails) Error() string {
	return fmt.Sprintf("missing %s for chain %s", e.Parameter, e.ChainIdentifier)
}

// ErrMissingChainClient is the error for missing chain client.
type ErrMissingChainClient struct {
	ChainIdentifier string
}

// Error returns the error message.
func (e *ErrMissingChainClient) Error() string {
	return fmt.Sprintf("missing chain client for chain %s", e.ChainIdentifier)
}

type ErrNoChainMetadata struct {
}

func (e *ErrNoChainMetadata) Error() string {
	return "no chain metadata"
}

type ErrNoTransactions struct {
}

func (e *ErrNoTransactions) Error() string {
	return "no transactions"
}

type ErrInvalidSignature struct {
	ChainIdentifier  string
	MCMSAddress      common.Address
	RecoveredAddress common.Address
}

func (e *ErrInvalidSignature) Error() string {
	return fmt.Sprintf("invalid signature: received signature for address %s is not a signer on MCMS %s on chain %s", e.RecoveredAddress, e.MCMSAddress, e.ChainIdentifier)
}

type ErrInvalidMCMSConfig struct {
	Reason string
}

func (e *ErrInvalidMCMSConfig) Error() string {
	return fmt.Sprintf("invalid MCMS config: %s", e.Reason)
}
