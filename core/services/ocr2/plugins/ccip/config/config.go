package config

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/core/store/models"
)

// RelayPluginConfig contains the plugin specific variables for the ccip.CCIPRelay plugin.
type RelayPluginConfig struct {
	SourceChainID int64           `json:"sourceChainID"`
	DestChainID   int64           `json:"destChainID"`
	OnRampID      types.Account   `json:"onRampID"`
	PollPeriod    models.Duration `json:"pollPeriod"`
}

// ValidateRelayPluginConfig validates the arguments for the CCIP Relay plugin.
// It will return an error if there is anything wrong with the provided config.
func (rp *RelayPluginConfig) ValidateRelayPluginConfig() error {
	if rp.SourceChainID <= 0 {
		return errors.Errorf("Invalid source chain value %d", rp.SourceChainID)
	}
	if rp.DestChainID <= 0 {
		return errors.Errorf("Invalid destination chain value %d", rp.DestChainID)
	}

	if _, err := hexutil.Decode(string(rp.OnRampID)); err != nil {
		return err
	}

	return nil
}

// ExecutionPluginConfig contains the plugin specific variables for the ccip.CCIPExecution plugin.
type ExecutionPluginConfig struct {
	RelayPluginConfig
	OffRampId string `json:"offRampId"`
}

// ValidateExecutionPluginConfig validates the arguments for the CCIP Execution plugin.
// It will return an error if there is anything wrong with the provided config.
func (ep *ExecutionPluginConfig) ValidateExecutionPluginConfig() error {
	if err := ep.ValidateRelayPluginConfig(); err != nil {
		return err
	}

	if _, err := hexutil.Decode(ep.OffRampId); err != nil {
		return err
	}

	return nil
}
