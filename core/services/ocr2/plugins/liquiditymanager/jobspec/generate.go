package lmjobspec

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// core/cmd/liquiditymanager_configure_commands.go

type JobSpecOptions struct {
	Name                    string
	ContractID              string
	OcrKeyBundleID          string
	TransmitterID           string
	RelayFromBlock          int64
	FollowerChains          string
	LiquidityManagerAddress common.Address
	NetworkSelector         uint64
	Type                    string
}

func (o *JobSpecOptions) Defaults() {
	if len(o.Name) == 0 {
		o.Name = "liquiditymanager"
	}
	if len(o.Type) == 0 {
		o.Type = "ping-pong"
	}
}

func NodeJobSpec(opts *JobSpecOptions) string {
	opts.Defaults()
	jobSpec := fmt.Sprintf(
		`
type                 	= "offchainreporting2"
schemaVersion        	= 1
name                 	= "%s"
maxTaskDuration      	= "30s"
contractID           	= "%s"
ocrKeyBundleID       	= "%s"
relay                	= "evm"
pluginType           	= "liquiditymanager"
transmitterID        	= "%s"
forwardingAllowed       = false
contractConfigTrackerPollInterval = "5s"

[relayConfig]
chainID              	= 1337
# This is the fromBlock for the main chain
fromBlock               = %d
[relayConfig.fromBlocks]
# these are the fromBlock values for the follower chains
%s

[pluginConfig]
liquidityManagerAddress = "%s"
liquidityManagerNetwork = "%d"
closePluginTimeoutSec = 10
[pluginConfig.rebalancerConfig]
type = "%s"
`,
		opts.Name,
		opts.ContractID,
		opts.OcrKeyBundleID,
		opts.TransmitterID,
		opts.RelayFromBlock,
		opts.FollowerChains,
		opts.LiquidityManagerAddress.Hex(),
		opts.NetworkSelector,
		opts.Type,
	)
	return jobSpec
}

func BootstrapJobSpec(l1ChainID int64, contractID string, fromBlock int64) string {
	return fmt.Sprintf(
		`
type          					  = "bootstrap"
schemaVersion 					  = 1
name          					  = "bootstrap-chainID-%d"
id            					  = "1"
contractID    					  = "%s"
relay         					  = "evm"
contractConfigTrackerPollInterval = "1m"

[relayConfig]
chainID       = %d
fromBlock     = %d
`,
		l1ChainID,
		contractID,
		l1ChainID,
		fromBlock,
	)
}
