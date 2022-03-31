package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
)

// PrintJobSpecs prints the job spec for each node and CCIP spec type, as well as a bootstrap spec.
func PrintJobSpecs(onramp, offramp, executor common.Address, sourceChainID, destChainID *big.Int) {
	jobs := fmt.Sprintf(bootstrapTemplate, helpers.ChainName(destChainID.Int64()), offramp, destChainID)
	oracles := getOraclesForChain(destChainID.Int64())
	for i, oracle := range oracles {
		jobs += "\n" + fmt.Sprintf(relayTemplate, i, helpers.ChainName(sourceChainID.Int64())+"-"+helpers.ChainName(destChainID.Int64()), offramp, keyBundleIDs[i],
			oracle.OracleIdentity.TransmitAccount, BootstrapPeer,
			onramp, sourceChainID, destChainID, destChainID,
		)
		jobs += fmt.Sprintf(executionTemplate, helpers.ChainName(sourceChainID.Int64())+"-"+helpers.ChainName(destChainID.Int64()), executor, keyBundleIDs[i],
			oracle.OracleIdentity.TransmitAccount, BootstrapPeer,
			onramp, offramp, sourceChainID, destChainID, destChainID,
		)
	}
	fmt.Println(jobs)
}

const bootstrapTemplate = `
// Bootstrap Node
# BootstrapSpec
type                               = "bootstrap"
name                               = "bootstrap-%s"
relay                              = "evm"
schemaVersion                      = 1
contractID                         = "%s"
contractConfigConfirmations        = 1
contractConfigTrackerPollInterval  = "60s"
[relayConfig]
chainID                            = %s
`

const relayTemplate = `
// Node %d
# CCIPRelaySpec
type               = "offchainreporting2"
name               = "ccip-relay-%s"
pluginType         = "ccip-relay"
relay              = "evm"
schemaVersion      = 1
contractID         = "%s"
ocrKeyBundleID     = "%s"
transmitterID      = "%s"
p2pBootstrapPeers  = ["%s"]

[pluginConfig]
onRampID           = "%s"
sourceChainID      = %d
destChainID        = %d

[relayConfig]
chainID            = "%d"
`

const executionTemplate = `
# CCIPExecutionSpec
type              = "offchainreporting2"
name              = "ccip-exec-%s"
pluginType        = "ccip-execution"
relay             = "evm"
schemaVersion     = 1
contractID        = "%s"
ocrKeyBundleID    = "%s"
transmitterID     = "%s"
p2pBootstrapPeers = ["%s"]

[pluginConfig]
onRampID          = "%s"
offRampID         = "%s"
sourceChainID     = %d
destChainID       = %d

[relayConfig]
chainID           = "%d"
`
