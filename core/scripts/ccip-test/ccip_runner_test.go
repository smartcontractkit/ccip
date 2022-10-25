package main

import (
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/dione"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/metis"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
)

var (
	SOURCE      = rhea.Staging_Beta_GoerliToAvaxFuji
	DESTINATION = rhea.Staging_Beta_AvaxFujiToGoerli
	ENV         = dione.StagingBeta
)

// These functions can be run as a test (prefix with Test) with the following config
// DATABASE_URL
// Use "-v" as a Go tool argument for streaming log output.

// TestPrintState can be run as a test with the following config
// OWNER_KEY  private key used to deploy all contracts and is used as default in all single user tests.
func TestMetisPrintState(t *testing.T) {
	ownerKey := os.Getenv("OWNER_KEY")
	if ownerKey == "" {
		t.Log("No command given, skipping ccip-test-script. This is intended behaviour for automated testing.")
		t.SkipNow()
	}
	SOURCE.SetupChain(t, ownerKey)
	DESTINATION.SetupChain(t, ownerKey)

	metis.PrintCCIPState(&SOURCE, &DESTINATION)

	rhea.PrintContractConfig(&SOURCE, &DESTINATION)
}

// TestMetisPrintNodeBalances can be run as a test with the following config
// OWNER_KEY  private key used to deploy all contracts and is used as default in all single user tests.
func TestMetisPrintNodeBalances(t *testing.T) {
	ownerKey := os.Getenv("OWNER_KEY")
	if ownerKey == "" {
		t.Log("No command given, skipping ccip-test-script. This is intended behaviour for automated testing.")
		t.SkipNow()
	}

	SOURCE.SetupChain(t, ownerKey)
	DESTINATION.SetupChain(t, ownerKey)

	don := dione.NewOfflineDON(ENV, logger.TestLogger(t))

	var sourceKeys, destKeys []common.Address

	for _, node := range don.Config.Nodes {
		sourceKeys = append(sourceKeys, common.HexToAddress(node.EthKeys[SOURCE.ChainId.String()]))
		destKeys = append(destKeys, common.HexToAddress(node.EthKeys[DESTINATION.ChainId.String()]))
	}
	metis.PrintNodeBalances(&SOURCE, sourceKeys)
	metis.PrintNodeBalances(&DESTINATION, destKeys)
}

// TestDeploySubscription can be run as a test with the following config
// OWNER_KEY  private key used to deploy all contracts and is used as default in all single user tests.
func TestRheaDeploySubscription(t *testing.T) {
	ownerKey := os.Getenv("OWNER_KEY")
	if ownerKey == "" {
		t.Log("No key given, this test will be skipped. This is intended behaviour for automated testing.")
		t.SkipNow()
	}
	rhea.DeploySubscriptionContracts(t, ownerKey, &SOURCE, &DESTINATION)
	rhea.PrintContractConfig(&SOURCE, &DESTINATION)
}

// TestDione can be run as a test with the following config
// OWNER_KEY  private key used to deploy all contracts and is used as default in all single user tests.
func TestDione(t *testing.T) {
	ownerKey := os.Getenv("OWNER_KEY")
	if ownerKey == "" {
		t.Log("No key given, this test will be skipped. This is intended behaviour for automated testing.")
		t.SkipNow()
	}
	SOURCE.SetupChain(t, ownerKey)
	DESTINATION.SetupChain(t, ownerKey)

	don := dione.NewDON(ENV, logger.TestLogger(t))
	//don := dione.NewOfflineDON(ENV, logger.TestLogger(t))
	don.WriteToFile()

	//don.FundNodeKeys(DESTINATION, ownerKey, big.NewInt(1e18))
	//don.FundNodeKeys(DESTINATION, ownerKey, big.NewInt(9e17))
	//don.DeleteKnownKey("4")
	//don.PopulateEthKeys()
	//don.PrintConfig()
	//don.ClearAllJobs(dione.Goerli, dione.AvaxFuji)
	//don.AddTwoWaySpecs(SOURCE, DESTINATION)
	//don.CreateNewEthKeysForChain(OptimismGoerliConfig.ChainId)
	//don.WIP()
	//don.ClearAllJobs(Rinkeby, Goerli)
	//don.WriteToFile()
}

// TestCCIP can be run as a test with the following config
// OWNER_KEY  private key used to deploy all contracts and is used as default in all single user tests.
// SEED_KEY   private key used for multi-user tests. Not needed when using the "deploy" command.
// COMMAND    what function to run e.g. "deploy", "setConfig", or "externalExecution".
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
	// The seed key is used to generate 10 keys from a single key by changing the
	// first character of the given seed with the digits 0-9
	seedKey := os.Getenv("SEED_KEY")
	if seedKey == "" {
		t.Error("must set seed key")
	}

	runCommand(t, ownerKey, seedKey, command)
}

func runCommand(t *testing.T, ownerKey string, seedKey string, command string) {
	// Configures a client to run tests with using the network defaults and given keys.
	// After updating any contracts be sure to update the network defaults to reflect
	// those changes.
	client := NewCcipClient(t,
		SOURCE,
		DESTINATION,
		ownerKey,
		seedKey,
	)

	SOURCE.SetupChain(t, ownerKey)
	DESTINATION.SetupChain(t, ownerKey)

	// Auto unpauses all contracts if they're paused.
	//client.UnpauseAll()

	switch command {
	// Deploys a new set of PingPong contracts, configures them to talk to each other
	// and creates destination chain subscriptions for both.
	case "deployPingPong":
		rhea.DeployPingPongDapps(t, &SOURCE, &DESTINATION)
		// Starts and unpauses the PingPong dapp that is on the `source` chain.
	case "startPingPong":
		client.startPingPong(t)
		// Stops the PingPong dapp by pausing the source chain dapp.
	case "stopPingPong":
		client.setPingPongPaused(t, true)
	case "fundPingPong":
		client.fundPingPong(t)
	case "printSpecs":
		metis.PrintJobSpecs(ENV, SOURCE.OnRamp, DESTINATION.BlobVerifier, DESTINATION.OffRamp, SOURCE.ChainId, DESTINATION.ChainId, DESTINATION.LinkToken, SOURCE.DeploySettings.DeployedAt, DESTINATION.DeploySettings.DeployedAt)
	case "setConfig":
		// Set the config to the blobVerifier and the offramp
		client.SetOCRConfig(ENV)
	case "upgradeLane":
		rhea.UpgradeLane(t, &SOURCE, &DESTINATION)
	case "dapp":
		client.SendDappTx(t)
		// Sends a new config to the governance dapp, spreading it to all configured chains
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
	case "tryPausedOnramp":
		// Should not succeed because the onramp is paused
		client.CrossChainSendPausedOnrampShouldFail(t)
	case "acceptOwnership":
		// Should accept ownership on the destination chain OffRamp & Executor
		client.AcceptOwnership(t)
		// work in progress call, use for any custom scripting
	case "wip":
		client.wip(t, GetSetupChain(t, ownerKey, SOURCE), GetSetupChain(t, ownerKey, DESTINATION))
	case "":
		t.Log("No command given, exit successfully")
		t.SkipNow()
	default:
		t.Errorf("Unknown command \"%s\"", command)
	}
}
