package main

import (
	"math/big"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/dione"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/metis/printing"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea/deployments"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
)

var (
	SOURCE      = deployments.Beta_AvaxFujiToSepolia
	DESTINATION = deployments.Beta_SepoliaToAvaxFuji
	SOURCES     = []*rhea.EvmDeploymentConfig{
		&deployments.Beta_SepoliaToAvaxFuji,
		&deployments.Beta_SepoliaToOptimismGoerli,
		&deployments.Beta_AvaxFujiToOptimismGoerli,
	}
	DESTINATIONS = []*rhea.EvmDeploymentConfig{
		&deployments.Beta_AvaxFujiToSepolia,
		&deployments.Beta_OptimismGoerliToSepolia,
		&deployments.Beta_OptimismGoerliToAvaxFuji,
	}
	ENV = dione.StagingBeta
)

var envToChainConfigs = map[dione.Environment][]rhea.EvmDeploymentConfig{
	dione.StagingAlpha: deployments.Alpha_ChainConfigs,
	dione.StagingBeta:  deployments.Beta_ChainConfigs,
	dione.Production:   deployments.Prod_ChainConfigs,
}

// These functions can be run as a test (prefix with Test) with the following config
// DATABASE_URL
// Use "-v" as a Go tool argument for streaming log output.

// TestDeploy can be run as a test with the following config
// OWNER_KEY  private key used to deploy all contracts and is used as default in all single user tests.
func TestRheaDeploy(t *testing.T) {
	checkOwnerKeyAndSetupChain(t)

	rhea.DeployLanes(t, &SOURCE, &DESTINATION)
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
	client := NewCcipClient(t,
		SOURCE,
		DESTINATION,
		ownerKey,
		seedKey,
	)

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
		// Set the config to the onRamp fee
	case "setAllowList":
		client.setAllowList(t)
	// Set the config to the onRamp AllowList
	case "upgradeLane":
		rhea.UpgradeLane(t, &SOURCE, &DESTINATION)
	case "gov":
		client.ChangeGovernanceParameters(t)
	case "don":
		// Cross chain request with DON execution
		client.DonExecutionHappyPath(t)
	case "batching":
		// Submit 10 txs. This should result in the txs being batched together
		client.ScalingAndBatching(t)
	case "gas":
		client.TestGasVariousTxs(t)
	case "acceptOwnership":
		// Should accept ownership on the destination chain OffRamp & Executor
		client.AcceptOwnership(t)
		// work in progress call, use for any custom scripting
	case "syncTokenPools":
		// Sync EvmChainConfig tokenPools to on-chain on/offRamp: remove deleted, add new BridgeTokens+TokenPools
		client.SyncTokenPools(t)
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
// 3. set ocrconfig for both
// OWNER_KEY  private key used to deploy all contracts and is used as default in all single user tests.
func TestUpdateAllLanes(t *testing.T) {
	ownerKey := checkOwnerKey(t)
	seedKey := os.Getenv("SEED_KEY")
	if seedKey == "" {
		t.Error("must set seed key")
	}
	require.Equal(t, len(SOURCES), len(DESTINATIONS), "number of sources and destinations should match")
	don := dione.NewDON(ENV, logger.TestLogger(t))
	for i, src := range SOURCES {
		dest := DESTINATIONS[i]
		src.SetupChain(t, ownerKey)
		dest.SetupChain(t, ownerKey)
		if !src.DeploySettings.DeployCommitStore || !src.DeploySettings.DeployRamp {
			src.Logger.Errorf("Please set \"DeployRamp and DeployCommitStore\" to true for the given EvmChainConfigs and make sure "+
				"the right ones are set. Source: %d, Dest %d", src.ChainConfig.ChainId, dest.ChainConfig.ChainId)
			continue
		}
		if !dest.DeploySettings.DeployCommitStore || !dest.DeploySettings.DeployRamp {
			dest.Logger.Errorf("Please set \"DeployRamp and DeployCommitStore\" to true for the given EvmChainConfigs and make sure "+
				"the right ones are set. Source: %d, Dest %d", dest.ChainConfig.ChainId, src.ChainConfig.ChainId)
			continue
		}
		rhea.UpgradeLaneTwoWay(t, src, dest)
		don.ClearAllJobs(helpers.ChainName(int64(src.ChainConfig.ChainId)), helpers.ChainName(int64(dest.ChainConfig.ChainId)))
		don.AddTwoWaySpecs(*src, *dest)
		client := NewCcipClient(t, *src, *dest, ownerKey, seedKey)
		client.SetOCR2Config(ENV)
		client = NewCcipClient(t, *dest, *src, ownerKey, seedKey)
		client.SetOCR2Config(ENV)
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
	for _, source := range envToChainConfigs[ENV] {
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
	for _, source := range envToChainConfigs[ENV] {
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
