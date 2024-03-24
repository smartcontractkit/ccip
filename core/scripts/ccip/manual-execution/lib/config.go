package lib

import (
	"fmt"

	"go.uber.org/multierr"

	"manual-execution/helpers"
)

// Config represents configuration fields
type Config struct {
	SrcNodeURL       string `json:"src_rpc"`
	DestNodeURL      string `json:"dest_rpc"`
	DestOwner        string `json:"dest_owner_key"`
	CommitStore      string `json:"commit_store"`
	OffRamp          string `json:"off_ramp"`
	DestStartBlock   uint64 `json:"dest_start_block"`
	SourceChainTx    string `json:"source_chain_tx"`
	CCIPMsgID        string `json:"ccip_msg_id"`
	DestDeployedAt   uint64 `json:"dest_deployed_at"`
	GasLimitOverride uint64 `json:"gas_limit_override"`
}

func (cfg Config) VerifyConfig() error {
	var allErr error
	if cfg.SrcNodeURL == "" {
		allErr = multierr.Append(allErr, fmt.Errorf("must set src_rpc - source chain rpc\n"))
	}
	if cfg.DestNodeURL == "" {
		allErr = multierr.Append(allErr, fmt.Errorf("must set dest_rpc - destination chain rpc\n"))
	}
	if cfg.DestOwner == "" {
		allErr = multierr.Append(allErr, fmt.Errorf("must set dest_owner_key - destination user private key\n"))
	}
	if cfg.SourceChainTx == "" {
		allErr = multierr.Append(allErr, fmt.Errorf("must set source_chain_tx - txHash of ccip-send request\n"))
	}

	if cfg.DestStartBlock == 0 && cfg.DestDeployedAt == 0 {
		allErr = multierr.Append(allErr, fmt.Errorf(`must set either of -
dest_deployed_at - the block number before destination contracts were deployed;
dest_start_block - the block number from which events will be filtered at destination chain.
`))
	}
	if cfg.GasLimitOverride == 0 {
		allErr = multierr.Append(allErr, fmt.Errorf("must set gas_limit_override - new value of gas limit for ccip-send request\n"))
	}
	err := helpers.VerifyAddress(cfg.CommitStore)
	if err != nil {
		allErr = multierr.Append(allErr, fmt.Errorf("check the commit_store address - %v\n", err))
	}
	err = helpers.VerifyAddress(cfg.OffRamp)
	if err != nil {
		allErr = multierr.Append(allErr, fmt.Errorf("check the off_ramp address - %v\n", err))
	}

	return allErr
}
