package rhea

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/governance_dapp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/ping_pong_demo"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/shared"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
)

// DeployLanes will deploy all source and Destination chain contracts using the
// owner key. Only run this of the currently deployed contracts are outdated or
// when initializing a new chain.
func DeployLanes(t *testing.T, source *EvmDeploymentConfig, destination *EvmDeploymentConfig) {
	sourceChainId, destChainId := source.ChainConfig.ChainId, destination.ChainConfig.ChainId
	// After running this code please update the configuration to reflect the newly
	// deployed contract addresses.
	// Deploy onRamps on both chains
	deploySourceContracts(t, source, destChainId, destination.ChainConfig.SupportedTokens)
	deploySourceContracts(t, destination, sourceChainId, source.ChainConfig.SupportedTokens)

	// Deploy commitStores and offRamps on both chains
	deployDestinationContracts(t, destination, sourceChainId, source.LaneConfig.OnRamp, source.ChainConfig.SupportedTokens)
	deployDestinationContracts(t, source, destChainId, destination.LaneConfig.OnRamp, destination.ChainConfig.SupportedTokens)

	SetPriceRegistryPrices(t, source, destChainId)
	SetPriceRegistryPrices(t, destination, sourceChainId)

	DeployPingPongDapps(t, source, destination)

	UpdateDeployedAt(t, source, destination)
}

func deploySourceContracts(t *testing.T, source *EvmDeploymentConfig, destChainId uint64, destSupportedTokens map[Token]EVMBridgedToken) {
	if source.LaneConfig.DeploySettings.DeployRamp {
		// Updates source.OnRamp if any new contracts are deployed
		deployOnRamp(t, source, destChainId, destSupportedTokens)
		setOnRampPrices(t, source)
		setOnRampOnRouter(t, source, destChainId)
	}

	// Skip if we reuse both the onRamp and the token pools
	if source.LaneConfig.DeploySettings.DeployRamp || source.ChainConfig.DeploySettings.DeployTokenPools {
		setOnRampOnTokenPools(t, source)
	}
	source.Logger.Infof("%s contracts deployed as source chain", helpers.ChainName(int64(source.ChainConfig.ChainId)))
}

func deployDestinationContracts(t *testing.T, client *EvmDeploymentConfig, sourceChainId uint64, onRamp common.Address, supportedTokens map[Token]EVMBridgedToken) {
	// Updates destClient.LaneConfig.CommitStore if any new contracts are deployed
	deployCommitStore(t, client, sourceChainId, onRamp)

	if client.LaneConfig.DeploySettings.DeployCommitStore || client.ChainConfig.DeploySettings.DeployPriceRegistry {
		setPriceRegistryUpdater(t, client)
	}

	if client.LaneConfig.DeploySettings.DeployRamp {
		// Updates destClient.LaneConfig.OffRamp if any new contracts are deployed
		deployOffRamp(t, client, sourceChainId, supportedTokens, onRamp)
		setOffRampPrices(t, client)
	}

	if client.LaneConfig.DeploySettings.DeployRamp || client.ChainConfig.DeploySettings.DeployTokenPools {
		setOffRampOnTokenPools(t, client)
	}

	if client.LaneConfig.DeploySettings.DeployRamp || client.ChainConfig.DeploySettings.DeployRouter {
		setOffRampOnRouter(t, sourceChainId, client)
	}

	client.Logger.Infof("%s contracts fully deployed as destination chain", helpers.ChainName(int64(client.ChainConfig.ChainId)))
}

func deployOnRamp(t *testing.T, client *EvmDeploymentConfig, destChainId uint64, destSupportedTokens map[Token]EVMBridgedToken) {
	if !client.LaneConfig.DeploySettings.DeployRamp {
		client.Logger.Infof("Skipping OnRamp deployment, using onRamp on %s", client.LaneConfig.OnRamp)
		return
	}

	var tokensAndPools []evm_2_evm_onramp.EVM2EVMOnRampTokenAndPool
	for token, tokenConfig := range client.ChainConfig.SupportedTokens {
		if _, ok := destSupportedTokens[token]; !ok {
			// If the token is not supported on the destination chain we
			// should not enable it for this ramp. If we enable the token,
			// txs could be sent but not executed, keeping the tokens in limbo.
			continue
		}

		tokensAndPools = append(tokensAndPools, evm_2_evm_onramp.EVM2EVMOnRampTokenAndPool{
			Token: tokenConfig.Token,
			Pool:  tokenConfig.Pool,
		})
	}

	var feeTokenConfig []evm_2_evm_onramp.EVM2EVMOnRampFeeTokenConfigArgs

	for _, feeToken := range client.ChainConfig.FeeTokens {
		tokenConfig := client.ChainConfig.SupportedTokens[feeToken]
		multiplier := uint64(1e18)
		// Let link cost 10% of the non-link fee. This helps with our ping pong running out of funds.
		if feeToken == LINK {
			multiplier = 1e17
		}

		feeTokenConfig = append(feeTokenConfig, evm_2_evm_onramp.EVM2EVMOnRampFeeTokenConfigArgs{
			Token:           tokenConfig.Token,
			Multiplier:      multiplier,
			FeeAmount:       big.NewInt(100e9),
			DestGasOverhead: 5_000,
		})
	}

	client.Logger.Infof("Deploying OnRamp: destinationChains %+v, tokensAndPools %+v", destChainId, tokensAndPools)
	onRampAddress, tx, _, err := evm_2_evm_onramp.DeployEVM2EVMOnRamp(
		client.Owner,  // user
		client.Client, // client
		evm_2_evm_onramp.EVM2EVMOnRampStaticConfig{
			LinkToken:         client.ChainConfig.SupportedTokens[LINK].Token,
			ChainId:           client.ChainConfig.ChainId,
			DestChainId:       destChainId,
			DefaultTxGasLimit: 200_000,
		},
		evm_2_evm_onramp.EVM2EVMOnRampDynamicConfig{
			Router:          client.ChainConfig.Router,
			PriceRegistry:   client.ChainConfig.PriceRegistry,
			MaxDataSize:     1e6,
			MaxTokensLength: 5,
			MaxGasLimit:     ccip.GasLimitPerTx,
			Afn:             client.ChainConfig.Afn,
		},
		tokensAndPools,
		[]common.Address{}, // allow list
		evm_2_evm_onramp.AggregateRateLimiterRateLimiterConfig{
			Capacity: new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e9)),
			Rate:     new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e5)),
			Admin:    client.Owner.From,
		},
		feeTokenConfig,
		[]evm_2_evm_onramp.EVM2EVMOnRampNopAndWeight{},
	)
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)

	client.Logger.Infof(fmt.Sprintf("Onramp deployed on %s in tx %s", onRampAddress.String(), helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash())))
	client.LaneConfig.OnRamp = onRampAddress
}

func deployOffRamp(t *testing.T, client *EvmDeploymentConfig, sourceChainId uint64, sourceTokens map[Token]EVMBridgedToken, onRamp common.Address) {
	if !client.LaneConfig.DeploySettings.DeployRamp {
		client.Logger.Infof("Skipping OffRamp deployment, using offRamp on %s", client.LaneConfig.OnRamp)
		return
	}

	var syncedSourceTokens []common.Address
	var syncedDestPools []common.Address

	for tokenName, tokenConfig := range sourceTokens {
		if _, ok := client.ChainConfig.SupportedTokens[tokenName]; ok {
			syncedSourceTokens = append(syncedSourceTokens, tokenConfig.Token)
			syncedDestPools = append(syncedDestPools, client.ChainConfig.SupportedTokens[tokenName].Pool)
		} else {
			client.Logger.Warnf("Token %s not supported by destination chain", tokenName)
		}
	}

	client.Logger.Infof("Deploying OffRamp")
	offRampAddress, tx, _, err := evm_2_evm_offramp.DeployEVM2EVMOffRamp(
		client.Owner,
		client.Client,
		evm_2_evm_offramp.EVM2EVMOffRampStaticConfig{
			CommitStore:   client.LaneConfig.CommitStore,
			ChainId:       client.ChainConfig.ChainId,
			SourceChainId: sourceChainId,
			OnRamp:        onRamp,
		},
		evm_2_evm_offramp.EVM2EVMOffRampDynamicConfig{
			Router:                                  client.ChainConfig.Router,
			ExecutionDelaySeconds:                   60,
			MaxDataSize:                             1e5,
			MaxTokensLength:                         15,
			PermissionLessExecutionThresholdSeconds: 60,
			Afn:                                     client.ChainConfig.Afn,
		},
		syncedSourceTokens,
		syncedDestPools,
		evm_2_evm_offramp.AggregateRateLimiterRateLimiterConfig{
			Capacity: new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e9)),
			Rate:     new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e5)),
			Admin:    client.Owner.From,
		},
	)
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)

	client.Logger.Infof("OffRamp contract deployed on %s in tx: %s", offRampAddress.Hex(), helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash()))
	client.LaneConfig.OffRamp = offRampAddress
	offRamp, err := evm_2_evm_offramp.NewEVM2EVMOffRamp(client.LaneConfig.OffRamp, client.Client)
	shared.RequireNoError(t, err)

	// Prices are used by the rate limiter and dictate what tokens are supported
	tx, err = offRamp.SetPrices(client.Owner, []common.Address{client.ChainConfig.SupportedTokens[LINK].Token}, []*big.Int{big.NewInt(10)})
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)

	client.Logger.Infof(fmt.Sprintf("Offramp configured for already deployed router in tx %s", helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash())))
}

func deployCommitStore(t *testing.T, client *EvmDeploymentConfig, sourceChainId uint64, onRamp common.Address) {
	if !client.LaneConfig.DeploySettings.DeployCommitStore {
		client.Logger.Infof("Skipping CommitStore deployment, using CommitStore on %s", client.LaneConfig.CommitStore)
		return
	}

	client.Logger.Infof("Deploying commitStore")
	commitStoreAddress, tx, _, err := commit_store.DeployCommitStore(
		client.Owner,  // user
		client.Client, // client
		commit_store.CommitStoreStaticConfig{
			ChainId:       client.ChainConfig.ChainId,
			SourceChainId: sourceChainId,
			OnRamp:        onRamp,
		},
		commit_store.CommitStoreDynamicConfig{
			PriceRegistry: client.ChainConfig.PriceRegistry,
			Afn:           client.ChainConfig.Afn,
		},
	)
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
	client.Logger.Infof("CommitStore deployed on %s in tx: %s", commitStoreAddress.Hex(), helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash()))
	client.LaneConfig.CommitStore = commitStoreAddress
}

func DeployPingPongDapps(t *testing.T, sourceClient *EvmDeploymentConfig, destClient *EvmDeploymentConfig) {
	fundingAmount := big.NewInt(1e18)

	if sourceClient.LaneConfig.DeploySettings.DeployPingPongDapp {
		feeToken := sourceClient.ChainConfig.SupportedTokens[LINK].Token
		sourceClient.Logger.Infof("Deploying source chain ping pong dapp")

		pingPongDappAddress, tx, _, err := ping_pong_demo.DeployPingPongDemo(
			sourceClient.Owner,
			sourceClient.Client,
			sourceClient.ChainConfig.Router,
			feeToken,
		)
		shared.RequireNoError(t, err)

		err = shared.WaitForMined(sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
		shared.RequireNoError(t, err)
		sourceClient.Logger.Infof("Ping pong deployed on %s in tx: %s", pingPongDappAddress.Hex(), helpers.ExplorerLink(int64(sourceClient.ChainConfig.ChainId), tx.Hash()))

		sourceClient.LaneConfig.PingPongDapp = pingPongDappAddress
		err = FundPingPong(sourceClient, fundingAmount, feeToken)
		shared.RequireNoError(t, err)
	}

	if destClient.LaneConfig.DeploySettings.DeployPingPongDapp {
		feeToken := destClient.ChainConfig.SupportedTokens[LINK].Token
		destClient.Logger.Infof("Deploying destination chain ping pong dapp")

		pingPongDappAddress, tx, _, err := ping_pong_demo.DeployPingPongDemo(
			destClient.Owner,
			destClient.Client,
			destClient.ChainConfig.Router,
			feeToken,
		)
		shared.RequireNoError(t, err)

		err = shared.WaitForMined(destClient.Logger, destClient.Client, tx.Hash(), true)
		shared.RequireNoError(t, err)
		destClient.Logger.Infof("Ping pong deployed on %s in tx: %s", pingPongDappAddress.Hex(), helpers.ExplorerLink(int64(destClient.ChainConfig.ChainId), tx.Hash()))

		destClient.LaneConfig.PingPongDapp = pingPongDappAddress
		err = FundPingPong(destClient, fundingAmount, feeToken)
		shared.RequireNoError(t, err)
	}

	if sourceClient.LaneConfig.DeploySettings.DeployPingPongDapp || destClient.LaneConfig.DeploySettings.DeployPingPongDapp {
		pingDapp, err := ping_pong_demo.NewPingPongDemo(sourceClient.LaneConfig.PingPongDapp, sourceClient.Client)
		shared.RequireNoError(t, err)

		tx, err := pingDapp.SetCounterpart(sourceClient.Owner, destClient.ChainConfig.ChainId, destClient.LaneConfig.PingPongDapp)
		shared.RequireNoError(t, err)
		err = shared.WaitForMined(sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
		shared.RequireNoError(t, err)
		sourceClient.Logger.Infof("Ping pong dapp configured in tx: %s", helpers.ExplorerLink(int64(sourceClient.ChainConfig.ChainId), tx.Hash()))

		pongDapp, err := ping_pong_demo.NewPingPongDemo(destClient.LaneConfig.PingPongDapp, destClient.Client)
		shared.RequireNoError(t, err)

		tx, err = pongDapp.SetCounterpart(destClient.Owner, sourceClient.ChainConfig.ChainId, sourceClient.LaneConfig.PingPongDapp)
		shared.RequireNoError(t, err)
		err = shared.WaitForMined(destClient.Logger, destClient.Client, tx.Hash(), true)
		shared.RequireNoError(t, err)
		destClient.Logger.Infof("Ping pong dapp configured in tx: %s", helpers.ExplorerLink(int64(destClient.ChainConfig.ChainId), tx.Hash()))
	} else {
		sourceClient.Logger.Infof("Skipping ping pong deployment")
	}
}

func deployGovernanceDapps(t *testing.T, sourceClient *EvmDeploymentConfig, destClient *EvmDeploymentConfig) {
	feeConfig := governance_dapp.GovernanceDappFeeConfig{
		FeeAmount:      big.NewInt(10),
		ChangedAtBlock: big.NewInt(0),
	}

	sourceClient.Logger.Infof("Deploying source chain governance dapp")
	governanceDappAddress, tx, _, err := governance_dapp.DeployGovernanceDapp(
		sourceClient.Owner,
		sourceClient.Client,
		sourceClient.ChainConfig.Router,
		feeConfig,
		destClient.ChainConfig.SupportedTokens[LINK].Token)
	require.NoError(t, err)

	err = shared.WaitForMined(sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
	sourceClient.Logger.Infof("GovernanceDapp deployed on %s in tx: %s", governanceDappAddress.Hex(), helpers.ExplorerLink(int64(sourceClient.ChainConfig.ChainId), tx.Hash()))

	sourceClient.LaneConfig.GovernanceDapp = governanceDappAddress

	destClient.Logger.Infof("Deploying destination chain governance dapp")
	governanceDappAddress, tx, _, err = governance_dapp.DeployGovernanceDapp(
		destClient.Owner,
		destClient.Client,
		destClient.ChainConfig.Router,
		feeConfig,
		destClient.ChainConfig.SupportedTokens[LINK].Token)
	require.NoError(t, err)

	err = shared.WaitForMined(destClient.Logger, destClient.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
	destClient.Logger.Infof("GovernanceDapp deployed on %s in tx: %s", governanceDappAddress.Hex(), helpers.ExplorerLink(int64(destClient.ChainConfig.ChainId), tx.Hash()))

	destClient.LaneConfig.GovernanceDapp = governanceDappAddress

	governanceDapp, err := governance_dapp.NewGovernanceDapp(sourceClient.LaneConfig.GovernanceDapp, sourceClient.Client)
	require.NoError(t, err)

	governanceClone := governance_dapp.GovernanceDappCrossChainClone{
		ChainId:         destClient.ChainConfig.ChainId,
		ContractAddress: destClient.LaneConfig.GovernanceDapp,
	}

	tx, err = governanceDapp.AddClone(sourceClient.Owner, governanceClone)
	require.NoError(t, err)
	err = shared.WaitForMined(sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
	sourceClient.Logger.Infof("GovernanceDapp configured in tx: %s", helpers.ExplorerLink(int64(sourceClient.ChainConfig.ChainId), tx.Hash()))
}
