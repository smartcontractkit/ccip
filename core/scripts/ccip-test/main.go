package main

import (
	"fmt"
	"os"
)

func main() {
	// This key is used to deploy all contracts on both source and Dest chains
	ownerKey := os.Getenv("OWNER_KEY")
	if ownerKey == "" {
		panic("must set owner key")
	}

	if len(os.Args) < 2 {
		fmt.Println("expected subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "deploy":
		deployContracts(ownerKey)
	default:
		runCommand(ownerKey, os.Args[1])
	}
}

func runCommand(ownerKey string, command string) {
	// The seed key is used to generate 10 keys from a single key by changing the
	// first character of the given seed with the digits 0-9
	seedKey := os.Getenv("SEED_KEY")
	if seedKey == "" {
		panic("must set seed key")
	}

	// Configures a client to run tests with using the network defaults and given keys.
	// After updating any contracts be sure to update the network defaults to reflect
	// those changes.
	client := NewCcipClient(
		// Source chain
		Kovan,
		// Dest chain
		Rinkeby,
		ownerKey,
		seedKey,
	)

	client.Source.Client.AssureHealth()
	client.Dest.Client.AssureHealth()
	client.UnpauseAll()

	switch command {
	case "setConfig":
		// Set the config to the message executor and the offramp
		client.SetConfig()
	case "externalExecution":
		// Cross chain request with the client manually proving and executing the transaction
		client.ExternalExecutionHappyPath()
	case "noRepeat":
		// Executing the same request twice should fail
		client.ExternalExecutionSubmitOfframpTwiceShouldFail()
	case "don":
		// Cross chain request with DON execution
		client.DonExecutionHappyPath()
	case "batching":
		// Submit 10 txs. This should result in the txs being batched together
		client.ScalingAndBatching()
	case "exceedBucket":
		// Should not be able to send funds greater than the amount in the bucket
		client.NotEnoughFundsInBucketShouldFail()
	case "tryPausedPool":
		// Should fail because the pool is paused
		client.TryGetTokensFromPausedPool()
	case "tryPausedOfframp":
		// Should not be included in a report because the offramp is paused
		client.CrossChainSendPausedOfframpShouldFail()
	case "tryPausedOnramp":
		// Should not succeed because the onramp is paused
		client.CrossChainSendPausedOnrampShouldFail()
	default:
		fmt.Println("[ERROR] unknown command \"" + command + "\"")
		os.Exit(1)
	}
}
