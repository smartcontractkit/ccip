package managed

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	owner "github.com/smartcontractkit/ccip-owner-contracts/gethwrappers"

	"github.com/smartcontractkit/chainlink/integration-tests/deployment/errors"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment/executable"
)

const (
	ZERO_HASH = "0x0000000000000000000000000000000000000000000000000000000000000000"
)

type MCMSWithTimelockChainMetadata struct {
	executable.ExecutableMCMSChainMetadata
	TimelockAddress common.Address `json:"timelockAddress"`
}

type TimelockOperation string

const (
	Schedule TimelockOperation = "schedule"
	Cancel   TimelockOperation = "cancel"
	Bypass   TimelockOperation = "bypass"
)

type MCMSWithTimelockProposal struct {
	baseMCMSProposal

	Operation TimelockOperation `json:"operation"` // Always 'schedule', 'cancel', or 'bypass'

	// TODO: this should be configurable as a human-readable string
	// i.e. 1d, 1w, 1m, 1y
	MinDelay string `json:"minDelay"`

	// Map of chain identifier to chain metadata
	ChainMetadata map[string]MCMSWithTimelockChainMetadata `json:"chainMetadata"`

	// Operations to be executed
	Transactions []DetailedBatchChainOperation `json:"transactions"`
}

func (m *MCMSWithTimelockProposal) Validate() error {
	if err := m.baseMCMSProposal.Validate(); err != nil {
		return err
	}

	switch m.Operation {
	case Schedule, Cancel, Bypass:
		break
	default:
		return &errors.ErrInvalidTimelockOperation{
			ReceivedTimelockOperation: string(m.Operation),
		}
	}

	_, err := time.ParseDuration(m.MinDelay)
	if err != nil {
		return &errors.ErrInvalidMinDelay{
			ReceivedMinDelay: m.MinDelay,
		}
	}

	return nil
}

func (m *MCMSWithTimelockProposal) ToExecutableMCMSProposal() (executable.ExecutableMCMSProposal, error) {
	raw := m.baseMCMSProposal.ToExecutableMCMSProposal()

	predecessorMap := make(map[string][32]byte)
	for chain := range m.ChainMetadata {
		predecessorMap[chain] = [32]byte(common.FromHex(ZERO_HASH))
	}

	for _, t := range m.Transactions {
		calls := make([]owner.RBACTimelockCall, 0)
		for _, op := range t.Batch {
			calls = append(calls, owner.RBACTimelockCall{
				Target: op.To,
				Data:   common.FromHex(op.Data),
				Value:  big.NewInt(int64(op.Value)),
			})
		}
		predecessor := predecessorMap[t.ChainIdentifier]
		salt := [32]byte(common.FromHex(ZERO_HASH))
		delay, _ := time.ParseDuration(m.MinDelay)

		abi, err := owner.RBACTimelockMetaData.GetAbi()
		if err != nil {
			return executable.ExecutableMCMSProposal{}, err
		}
		data, err := abi.Pack("scheduleBatch", calls, predecessor, salt, big.NewInt(int64(delay.Seconds())))
		if err != nil {
			return executable.ExecutableMCMSProposal{}, err
		}

		raw.Transactions = append(raw.Transactions, executable.ChainOperation{
			ChainIdentifier: t.ChainIdentifier,

			Operation: executable.Operation{
				To:    m.ChainMetadata[t.ChainIdentifier].TimelockAddress,
				Data:  common.Bytes2Hex(data),
				Value: 0, // TODO: is this right?
			},
		})

		predecessorMap[t.ChainIdentifier], err = hashOperationBatch(calls, predecessor, salt)
		if err != nil {
			return executable.ExecutableMCMSProposal{}, err
		}
	}

	return raw, nil
}

// hashOperationBatch replicates the hash calculation from Solidity
// TODO: see if there's an easier way to do this using the gethwrappers
func hashOperationBatch(calls []owner.RBACTimelockCall, predecessor, salt [32]byte) ([32]byte, error) {
	// Encode the calls using RLP encoding
	encodedCalls, err := rlp.EncodeToBytes(calls)
	if err != nil {
		return [32]byte{}, err
	}

	// Encode the entire data (calls, predecessor, salt) using ABI encoding
	encoded := crypto.Keccak256(
		append(encodedCalls, append(predecessor[:], salt[:]...)...),
	)

	// Return the hash as a [32]byte array
	return [32]byte(crypto.Keccak256Hash(encoded).Bytes()), nil
}
