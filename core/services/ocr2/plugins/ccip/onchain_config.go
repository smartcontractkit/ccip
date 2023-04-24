package ccip

import (
	"time"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_offramp"
)

type CommitOnchainConfig commit_store.CommitStoreDynamicConfig

func (d CommitOnchainConfig) AbiString() string {
	return `
	[
		{
			"components": [
				{"name": "priceRegistry", "type": "address"},
				{"name": "afn", "type": "address"}
			],
			"type": "tuple"
		}
	]`
}

type ExecOnchainConfig evm_2_evm_offramp.EVM2EVMOffRampDynamicConfig

func (d ExecOnchainConfig) AbiString() string {
	return `
	[
		{
			"components": [
				{"name": "permissionLessExecutionThresholdSeconds", "type": "uint32"},
				{"name": "router", "type": "address"},
				{"name": "afn", "type": "address"},
				{"name": "maxTokensLength", "type": "uint16"},
				{"name": "maxDataSize", "type": "uint32"}
			],
			"type": "tuple"
		}
	]`
}

func (d ExecOnchainConfig) PermissionLessExecutionThresholdDuration() time.Duration {
	return time.Duration(d.PermissionLessExecutionThresholdSeconds) * time.Second
}
