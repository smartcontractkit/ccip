package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
)

// PrintJobSpecs prints the job spec for each node and CCIP spec type, as well as a bootstrap spec.
func PrintJobSpecs(onramp, blobVerifier, offRamp common.Address, sourceChainID, destChainID *big.Int) {
	jobs := fmt.Sprintf(bootstrapTemplate+"\n", helpers.ChainName(destChainID.Int64()), blobVerifier, destChainID)
	oracles := getOraclesForChain(destChainID.Int64())
	for i, oracle := range oracles {
		jobs += fmt.Sprintf("// [Node %d]\n", i)
		jobs += fmt.Sprintf(relayTemplate+"\n",
			helpers.ChainName(sourceChainID.Int64())+"-"+helpers.ChainName(destChainID.Int64()),
			blobVerifier,
			keyBundleIDs[i],
			oracle.OracleIdentity.TransmitAccount,
			BootstrapPeer,
			sourceChainID,
			destChainID,
			onramp,
			pollPeriod,
			destChainID,
		)
		jobs += fmt.Sprintf(executionTemplate+"\n",
			helpers.ChainName(sourceChainID.Int64())+"-"+helpers.ChainName(destChainID.Int64()),
			offRamp,
			keyBundleIDs[i],
			oracle.OracleIdentity.TransmitAccount,
			BootstrapPeer,
			sourceChainID,
			destChainID,
			onramp,
			blobVerifier,
			destChainID,
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
sourceChainID      = %d
destChainID        = %d
onRampIDs          = ["%s"]
pollPeriod         = "%s"

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
sourceChainID     = %d
destChainID       = %d
onRampID          = "%s"
blobVerifierID    = "%s"

[relayConfig]
chainID           = "%d"
`
