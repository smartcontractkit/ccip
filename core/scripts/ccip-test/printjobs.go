package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// PrintJobSpecs prints the job spec for each node and CCIP spec type, as well as a bootstrap spec.
func PrintJobSpecs(onramp, offramp, executor common.Address, sourceChainID, destChainID *big.Int) {
	jobs := fmt.Sprintf(bootstrapTemplate, offramp)
	for i, oracle := range Oracles {
		jobs += "\n" + fmt.Sprintf(relayTemplate, i,
			oracle.OracleIdentity.TransmitAccount, BootstrapPeerID,
			offramp, onramp, sourceChainID, destChainID, destChainID,
		)
		jobs += fmt.Sprintf(executionTemplate,
			oracle.OracleIdentity.TransmitAccount, BootstrapPeerID,
			onramp, offramp, executor, sourceChainID, destChainID, destChainID,
		)
	}
	fmt.Println(jobs)
}

const bootstrapTemplate = `
// Bootstrap Node
# BootstrapSpec
type                               = "bootstrap"
name                               = "bootstrap"
relay                              = "evm"
schemaVersion                      = 1
contractID                         = "%s"
contractConfigConfirmations        = 1
contractConfigTrackerPollInterval  = "60s"
[relayConfig]
chainID                            = 4
`

const relayTemplate = `
// Node %d
# CCIPRelaySpec
type               = "offchainreporting2"
name               = "ccip-relay"
pluginType         = "ccip-relay"
relay              = "evm"
schemaVersion      = 1
ocrKeyBundleID     = "<KEY-BUNDLE-ID>"
transmitterID      = "%s"
p2pBootstrapPeers  = ["%s@<BOOTSTRAP-HOST>:<PORT>"]

[pluginConfig]
offRampID          = "%s"
onRampID           = "%s"
sourceChainID      = "%d"
destChainID        = "%d"

[relayConfig]
chainID            = "%d"
`

const executionTemplate = `
# CCIPExecutionSpec
type              = "ccip-execution"
name              = "ccip-execution"
pluginType        = "ccip-execution"
relay             = "evm"
schemaVersion     = 1
ocrKeyBundleID    = "<KEY-BUNDLE-ID>"
transmitterID     = "%s"
p2pBootstrapPeers = ["%s@<BOOTSTRAP-HOST>:<PORT>"]

[pluginConfig]
onRampID          = "%s"
offRampID         = "%s"
executorID        = "%s"
sourceChainID     = "%d"
destChainID       = "%d"

[relayConfig]
chainID           = "%d"
`
