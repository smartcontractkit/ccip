package main

import (
	"math/big"
	"os"
	"testing"

	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/dione"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/metis/printing"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea/deployment_io"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea/deployments"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
)

var (
	// Change these values
	sourceChain = rhea.OptimismGoerli
	destChain   = rhea.AvaxFuji
	ENV         = dione.StagingBeta

	// These will automatically populate or error if the lane doesn't exist
	SOURCE      = laneMapping[ENV][sourceChain][destChain]
	DESTINATION = laneMapping[ENV][destChain][sourceChain]
)

var laneMapping = map[dione.Environment]map[rhea.Chain]map[rhea.Chain]rhea.EvmDeploymentConfig{
	dione.StagingAlpha: deployments.AlphaChainMapping,
	dione.StagingBeta:  deployments.BetaChainMapping,
	dione.Production:   deployments.ProdChainMapping,
}

var chainMapping = map[dione.Environment]map[rhea.Chain]rhea.EvmDeploymentConfig{
	dione.StagingAlpha: deployments.AlphaChains,
	dione.StagingBeta:  deployments.BetaChains,
	dione.Production:   deployments.ProdChains,
}

// These functions can be run as a test (prefix with Test) with the following config
// DATABASE_URL
// Use "-v" as a Go tool argument for streaming log output.

// TestDeploy can be run as a test with the following config
// OWNER_KEY  private key used to deploy all contracts and is used as default in all single user tests.
func TestRheaDeploy(t *testing.T) {
	checkOwnerKeyAndSetupChain(t)
	rhea.DeployToNewChain(t, &SOURCE)
	rhea.DeployToNewChain(t, &DESTINATION)
	rhea.DeployLanes(t, &SOURCE, &DESTINATION)
	deployment_io.PrettyPrintLanes(ENV, &SOURCE, &DESTINATION)
}

// TestDione can be run as a test with the following config
// OWNER_KEY  private key used to deploy all contracts and is used as default in all single user tests.
func TestDione(t *testing.T) {
	checkOwnerKeyAndSetupChain(t)

	don := dione.NewDON(ENV, logger.TestLogger(t))
	don.ClearAllJobs(helpers.ChainName(int64(SOURCE.ChainConfig.ChainId)), helpers.ChainName(int64(DESTINATION.ChainConfig.ChainId)))
	don.AddTwoWaySpecs(SOURCE, DESTINATION)

	// Sometimes jobs don't get added correctly. This script looks for missing jobs
	// and attempts to add them.
	don.AddMissingSpecs(DESTINATION, SOURCE)
	don.AddMissingSpecs(SOURCE, DESTINATION)
}

// TestCCIP can be run as a test with the following config
// OWNER_KEY  private key used to deploy all contracts and is used as default in all single user tests.
// SEED_KEY   private key used for multi-user tests. Not needed when using the "deploy" command.
// COMMAND    what function to run e.g. "deploy", "setConfig", or "gas".
func TestCCIP(t *testing.T) {
	ownerKey := checkOwnerKeyAndSetupChain(t)
	command := os.Getenv("COMMAND")
	if command == "" {
		t.Log("No command given, skipping ccip-test-script. This is intended behaviour for automated testing.")
		t.SkipNow()
	}
	// The seed key is used to generate 10 keys from a single key by changing the
	// first character of the given seed with the digits 0-9
	seedKey := os.Getenv("SEED_KEY")
	if seedKey == "" {
		t.Error("must set seed key")
	}

	// Configures a client to run tests with using the network defaults and given keys.
	// After updating any contracts be sure to update the network defaults to reflect
	// those changes.
	client := NewCcipClient(t, SOURCE, DESTINATION, ownerKey, seedKey)

	switch command {
	// Deploys a new set of PingPong contracts, configures them to talk to each other
	case "deployPingPong":
		rhea.DeployPingPongDapps(t, &SOURCE, &DESTINATION)
		// Starts and unpauses the PingPong dapp that is on the `source` chain.
	case "startPingPong":
		client.startPingPong(t)
		// Stops the PingPong dapp by pausing the source chain dapp.
	case "stopPingPong":
		client.setPingPongPaused(t, true)
	case "fundPingPong":
		client.fundPingPong(t, &SOURCE, &DESTINATION)
	case "printSpecs":
		printing.PrintJobSpecs(ENV, SOURCE, DESTINATION)
	case "setConfig":
		// Set the config to the commitStore and the offramp
		client.SetOCR2Config(ENV)
	case "setOnRampFeeConfig":
		client.setOnRampFeeConfig(t, &SOURCE)
	case "applyFeeTokensUpdates":
		client.applyFeeTokensUpdates(t, &SOURCE)
	case "batching":
		// Submit 10 txs. This should result in the txs being batched together
		client.ScalingAndBatching(t)
	case "gas":
		client.TestGasVariousTxs(t)
	case "wip":
		client.wip(t, &SOURCE, &DESTINATION)
	case "":
		t.Log("No command given, exit successfully")
		t.SkipNow()
	default:
		t.Errorf("Unknown command \"%s\"", command)
	}
}

// TestUpdateAllLanes
// 1. updates all the available lanes with new offramp, onramp, commit store
// 2. creates new jobs
// 3. set ocrConfig for both
// OWNER_KEY  private key used to deploy all contracts and is used as default in all single user tests.
func TestUpdateAllLanes(t *testing.T) {
	ownerKey := checkOwnerKey(t)
	if _, ok := laneMapping[ENV]; !ok {
		t.Error("set environment not supported")
	}

	don := dione.NewDON(ENV, logger.TestLogger(t))

	// Potential todo: remove old deployment artifact permissions
	// Optimizations:
	// 		Concurrent chain contracts deployment before any lange deployment
	// 		Concurrent lane contract deployment for non-intersecting lanes
	// 		Concurrent lane contract deployment within a bidirectional deploy
	// 		Not waiting for mining, self incrementing the nonce

	// 		Downsides: less control and worse retry experience
	// 			As failures should be very rare this is probably worth it
	upgradeLane := func(source, dest rhea.EvmDeploymentConfig) {
		if !source.LaneConfig.DeploySettings.DeployCommitStore || !source.LaneConfig.DeploySettings.DeployRamp {
			source.Logger.Warnf("Please set \"DeployRamp and DeployCommitStore\" to true for the given EvmChainConfigs and make sure "+
				"the right ones are set. Source: %d, Dest %d", source.ChainConfig.ChainId, dest.ChainConfig.ChainId)
			return
		}
		if !dest.LaneConfig.DeploySettings.DeployCommitStore || !dest.LaneConfig.DeploySettings.DeployRamp {
			dest.Logger.Warnf("Please set \"DeployRamp and DeployCommitStore\" to true for the given EvmChainConfigs and make sure "+
				"the right ones are set. Source: %d, Dest %d", dest.ChainConfig.ChainId, source.ChainConfig.ChainId)
			return
		}
		if source.ChainConfig.DeploySettings.DeployRouter || dest.ChainConfig.DeploySettings.DeployRouter {
			dest.Logger.Warnf("Routers should never be set to true Source: %d, Dest %d", dest.ChainConfig.ChainId, source.ChainConfig.ChainId)
			return
		}
		// Removes any old job specs
		don.ClearAllJobs(helpers.ChainName(int64(source.ChainConfig.ChainId)), helpers.ChainName(int64(dest.ChainConfig.ChainId)))
		// Deploys the new contracts and updates `source` and `dest`
		rhea.DeployLanes(t, &source, &dest)
		// Prints the new config and writes them to file
		deployment_io.PrettyPrintLanes(ENV, &source, &dest)
		// Add new job specs
		don.AddTwoWaySpecs(source, dest)
		// Set the OCR2 config on the source contracts
		client := NewCcipClient(t, source, dest, ownerKey, ownerKey)
		client.SetOCR2Config(ENV)
		// Set the OCR2 config on the destination contracts
		client = NewCcipClient(t, dest, source, ownerKey, ownerKey)
		client.SetOCR2Config(ENV)
		// Starts the ping pong dapp
		client.startPingPong(t)
	}

	// This script only deploys new lane contracts. Please deploy any new chain contracts
	// and update the config before running this.

	DoForEachBidirectionalLane(t, ownerKey, upgradeLane)
}

func DoForEachLane(t *testing.T, ownerKey string, f func(source rhea.EvmDeploymentConfig, destination rhea.EvmDeploymentConfig)) {
	for sourceChain, sourceMap := range laneMapping[ENV] {
		for destChain, _ := range sourceMap {
			t.Logf("Running function for lane %s -> %s", sourceChain, destChain)

			source := laneMapping[ENV][sourceChain][destChain]
			dest := laneMapping[ENV][destChain][sourceChain]

			source.SetupChain(t, ownerKey)
			dest.SetupChain(t, ownerKey)

			f(source, dest)
		}
	}
}

func DoForEachBidirectionalLane(t *testing.T, ownerKey string, f func(source rhea.EvmDeploymentConfig, destination rhea.EvmDeploymentConfig)) {
	completed := make(map[rhea.Chain]map[rhea.Chain]interface{})

	for sourceChain, sourceMap := range laneMapping[ENV] {
		for destChain, _ := range sourceMap {
			// Skip if we already processed the lane from the other side
			if destMap, ok := completed[destChain]; ok {
				if _, ok := destMap[sourceChain]; ok {
					continue
				}
			}

			t.Logf("Running function for lane %s <-> %s", sourceChain, destChain)

			source := laneMapping[ENV][sourceChain][destChain]
			dest := laneMapping[ENV][destChain][sourceChain]

			source.SetupChain(t, ownerKey)
			dest.SetupChain(t, ownerKey)

			f(source, dest)

			if _, ok := completed[sourceChain]; !ok {
				completed[sourceChain] = make(map[rhea.Chain]interface{})
			}
			if _, ok := completed[destChain]; !ok {
				completed[destChain] = make(map[rhea.Chain]interface{})
			}
			completed[sourceChain][destChain] = true
			completed[destChain][sourceChain] = true
		}
	}
}

// TestPrintNodeBalances can be run as a test with the following config
// OWNER_KEY  private key used to deploy all contracts and is used as default in all single user tests.
func TestPrintNodeBalances(t *testing.T) {
	checkOwnerKeyAndSetupChain(t)

	don := dione.NewOfflineDON(ENV, logger.TestLogger(t))

	printing.PrintNodeBalances(&SOURCE, don.GetSendingKeys(SOURCE.ChainConfig.ChainId))
	printing.PrintNodeBalances(&DESTINATION, don.GetSendingKeys(DESTINATION.ChainConfig.ChainId))
}

func TestFundNodes(t *testing.T) {
	key := checkOwnerKeyAndSetupChain(t)

	don := dione.NewOfflineDON(ENV, logger.TestLogger(t))
	don.FundNodeKeys(&SOURCE, key, big.NewInt(4e18), big.NewInt(4e18))
}

// TestPrintAllNodeBalancesPerEnv can be run as a test with the following config
// OWNER_KEY  private key used to deploy all contracts and is used as default in all single user tests.
// It will print the node balances for all chains where the given `env` is deployed
func TestPrintAllNodeBalancesPerEnv(t *testing.T) {
	ownerKey := checkOwnerKey(t)

	for _, source := range chainMapping[ENV] {
		source.SetupChain(t, ownerKey)
		don := dione.NewOfflineDON(ENV, logger.TestLogger(t))
		printing.PrintNodeBalances(&source, don.GetSendingKeys(source.ChainConfig.ChainId))
	}
}

// TestFundAllNodesPerEnv can be run as a test with the following config
// OWNER_KEY  private key used to deploy all contracts and is used as default in all single user tests.
// It will fund the node balances for all chains where the given `env` is deployed
func TestFundAllNodesPerEnv(t *testing.T) {
	ownerKey := checkOwnerKey(t)
	for _, source := range chainMapping[ENV] {
		source.SetupChain(t, ownerKey)
		don := dione.NewOfflineDON(ENV, logger.TestLogger(t))
		don.FundNodeKeys(&source, ownerKey, big.NewInt(1e18), big.NewInt(4e18))
	}
}

func checkOwnerKeyAndSetupChain(t *testing.T) string {
	ownerKey := checkOwnerKey(t)
	SOURCE.SetupChain(t, ownerKey)
	DESTINATION.SetupChain(t, ownerKey)

	return ownerKey
}

func checkOwnerKey(t *testing.T) string {
	ownerKey := os.Getenv("OWNER_KEY")
	if ownerKey == "" {
		t.Log("No key given, this test will be skipped. This is intended behaviour for automated testing.")
		t.SkipNow()
	}

	return ownerKey
}

// This ALWAYS uses the production env
func Test__PROD__SetAllowListAllLanes(t *testing.T) {
	ownerKey := checkOwnerKey(t)

	// Simply comment out the lanes that are not needed.
	allProdLanes := []*rhea.EvmDeploymentConfig{
		&deployments.Prod_SepoliaToOptimismGoerli,
		&deployments.Prod_SepoliaToAvaxFuji,

		&deployments.Prod_AvaxFujiToSepolia,
		&deployments.Prod_AvaxFujiToOptimismGoerli,

		&deployments.Prod_OptimismGoerliToAvaxFuji,
		&deployments.Prod_OptimismGoerliToSepolia,
	}

	for _, lane := range allProdLanes {
		lane.SetupChain(t, ownerKey)
		client := CCIPClient{Source: NewSourceClient(t, *lane)}
		client.Source.Owner = rhea.GetOwner(t, ownerKey, client.Source.ChainId, lane.ChainConfig.GasSettings)

		client.setAllowList(t)
	}
}
