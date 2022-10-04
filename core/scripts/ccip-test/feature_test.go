package main

import (
	"os"
	"testing"

	"github.com/smartcontractkit/chainlink/core/logger"
)

var SOURCE = GoerliConfig
var DESTINATION = RinkebyConfig

// FullFeatureCCIP can be run as a test (prefix with Test) with the following config
// Env vars:
// OWNER_KEY  private key used to deploy all contracts and is used as default in all single user tests.
// SEED_KEY   private key used for multi-user tests. Not needed when using the "deploy" command.
// COMMAND    what function to run e.g. "deploy", "setConfig", or "externalExecution".
//
// Use "-v" as a Go tool argument for streaming log output.
func TestPrintState(t *testing.T) {
	ownerKey := os.Getenv("OWNER_KEY")
	if ownerKey == "" {
		t.Log("No command given, skipping ccip-test-script. This is intended behaviour for automated testing.")
		t.SkipNow()
	}
	printCCIPState(
		GetSetupChain(t, ownerKey, SOURCE),
		GetSetupChain(t, ownerKey, DESTINATION))
}

func TestCCIP(t *testing.T) {
	ownerKey := os.Getenv("OWNER_KEY")
	command := os.Getenv("COMMAND")
	if ownerKey == "" {
		if command == "" {
			t.Log("No command given, skipping ccip-test-script. This is intended behaviour for automated testing.")
			t.SkipNow()
		}
		t.Log("Must set owner key")
		t.FailNow()
	}

	switch command {
	case "":
		t.Log("No command given, exit successfully")
		t.SkipNow()
	case "deploy":
		deploySubscriptionContracts(t, ownerKey,
			&SOURCE,
			&DESTINATION)
	case "readOCRKeys":
		don := NewDON(Staging, logger.TestLogger(t))
		don.WIP()
		//don.ListJobSpecs()
		//don.WriteConfig()
	default:
		runCommand(t, ownerKey, command)
	}
}

func runCommand(t *testing.T, ownerKey string, command string) {
	// The seed key is used to generate 10 keys from a single key by changing the
	// first character of the given seed with the digits 0-9
	seedKey := os.Getenv("SEED_KEY")
	if seedKey == "" {
		t.Error("must set seed key")
	}

	// Configures a client to run tests with using the network defaults and given keys.
	// After updating any contracts be sure to update the network defaults to reflect
	// those changes.
	client := NewCcipClient(t,
		SOURCE,
		DESTINATION,
		ownerKey,
		seedKey,
	)

	// Auto unpauses all contracts if they're paused.
	client.UnpauseAll()

	switch command {
	// Deploys a new set of PingPong contracts, configures them to talk to each other
	// and creates destination chain subscriptions for both.
	case "deployPingPong":
		deployPingPongDapps(t, GetSetupChain(t, ownerKey, SOURCE), GetSetupChain(t, ownerKey, DESTINATION))

		// Starts and unpauses the PingPong dapp that is on the `source` chain.
	case "startPingPong":
		client.startPingPong(t)

		// Stops the PingPong dapp by pausing the source chain dapp.
	case "stopPingPong":
		client.setPingPongPaused(t, true)

	case "setConfig":
		// Set the config to the blobVerifier and the offramp
		client.SetOCRConfig()
	case "dapp":
		client.SendDappTx(t)
	case "gov":
		client.ChangeGovernanceParameters(t)
	case "don":
		// Cross chain request with DON execution
		client.DonExecutionHappyPath(t)
	case "batching":
		// Submit 10 txs. This should result in the txs being batched together
		client.ScalingAndBatching(t)
	case "exceedBucket":
		// Should not be able to send funds greater than the amount in the bucket
		client.NotEnoughFundsInBucketShouldFail(t)
	case "tryPausedPool":
		// Should fail because the pool is paused
		client.TryGetTokensFromPausedPool()
	//case "tryPausedOfframp":
	//	// Should not be included in a report because the offramp is paused
	//	client.CrossChainSendPausedOfframpShouldFail(t)
	case "tryPausedOnramp":
		// Should not succeed because the onramp is paused
		client.CrossChainSendPausedOnrampShouldFail(t)
	case "acceptOwnership":
		// Should accept ownership on the destination chain OffRamp & Executor
		client.AcceptOwnership(t)
		//case "externalExecution":
	//	// Cross chain request with the client manually proving and executing the transaction
	//	client.ExternalExecutionHappyPath(t)
	//case "noRepeat":
	//	// Executing the same request twice should fail
	//	client.ExternalExecutionSubmitOfframpTwiceShouldFail(t)
	case "wip":
		client.wip(t, GetSetupChain(t, ownerKey, SOURCE), GetSetupChain(t, ownerKey, DESTINATION))
	default:
		t.Errorf("Unknown command \"%s\"", command)
	}
}
