package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func PrintJobSpecs(onramp, offramp, executor common.Address) {
	jobs := fmt.Sprintf(bootstrapTemplate, offramp)
	for i, oracle := range Oracles {
		jobs += "\n" + fmt.Sprintf(relayTemplate, i, offramp, onramp,
			Kovan.ChainId.Int64(), Rinkeby.ChainId.Int64(),
			oracle.OracleIdentity.TransmitAccount, BootstrapPeerID)
		jobs += fmt.Sprintf(executionTemplate, onramp, offramp, executor,
			Kovan.ChainId.Int64(), Rinkeby.ChainId.Int64(),
			oracle.OracleIdentity.TransmitAccount, BootstrapPeerID)
	}
	fmt.Println(jobs)
}

const bootstrapTemplate = `
// Bootstrap Node
# CCIPBootstrapSpec
type                                = "ccip-bootstrap"
name                                = "ccip-bootstrap"
schemaVersion                       = 1
contractAddress                     = "%s"
evmChainID                          = 4
isBootstrapPeer                     = true
contractConfigConfirmations         = 1
contractConfigTrackerPollInterval   = "60s"
`

const relayTemplate = `
// Node %d
# CCIPRelaySpec
type				= "ccip-relay"
name				= "ccip-relay"
schemaVersion		= 1
offRampID			= "%s"
onRampID			= "%s"
sourceEvmChainID	= "%d"
destEvmChainID		= "%d"
ocrKeyBundleID		= "<KEY-BUNDLE-ID>"
transmitterID		= "%s"
p2pBootstrapPeers	= ["%s@<BOOTSTRAP-HOST>:<PORT>"]
relay				= "evm"
`

const executionTemplate = `
# CCIPExecutionSpec
type				= "ccip-execution"
name				= "ccip-execution"
schemaVersion		= 1
onRampID			= "%s"
offRampID			= "%s"
executorID			= "%s"
sourceEvmChainID	= "%d"
destEvmChainID		= "%d"
ocrKeyBundleID		= "<KEY-BUNDLE-ID>"
transmitterID		= "%s"
p2pBootstrapPeers	= ["%s@<BOOTSTRAP-HOST>:<PORT>"]
relay 				= "evm"
`
