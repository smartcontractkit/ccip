package rhea

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/governance_dapp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/ping_pong_demo"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/receiver_dapp"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/shared"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
)

// DeployGELanes will deploy all source and Destination chain contracts using the
// owner key. Only run this of the currently deployed contracts are outdated or
// when initializing a new chain.
func DeployGELanes(t *testing.T, source *EvmDeploymentConfig, destination *EvmDeploymentConfig) {
	sourceChainId, destChainId := source.ChainConfig.ChainId, destination.ChainConfig.ChainId
	// After running this code please update the configuration to reflect the newly
	// deployed contract addresses.
	DeployToNewChain(t, source)
	DeployToNewChain(t, destination)

	// Deploy onRamps on both chains
	deploySourceContracts(t, source, destChainId, destination.ChainConfig.SupportedTokens)
	deploySourceContracts(t, destination, sourceChainId, source.ChainConfig.SupportedTokens)

	// Deploy commitStores and offRamps on both chains
	prettyPrintLanes(source, destination)
	deployDestinationContracts(t, destination, sourceChainId, source.LaneConfig.OnRamp, source.ChainConfig.SupportedTokens)
	deployDestinationContracts(t, source, destChainId, destination.LaneConfig.OnRamp, destination.ChainConfig.SupportedTokens)

	SetFeeManagerPrices(t, source, destChainId)
	SetFeeManagerPrices(t, destination, sourceChainId)

	deployGovernanceDapps(t, source, destination)

	DeployPingPongDapps(t, source, destination)

	prettyPrintLanes(source, destination)
}

func prettyPrintLanes(source *EvmDeploymentConfig, destination *EvmDeploymentConfig) {
	source.Logger.Info(prettyPrint(source.ChainConfig))
	source.Logger.Info(prettyPrint(source.LaneConfig))

	destination.Logger.Info(prettyPrint(destination.ChainConfig))
	destination.Logger.Info(prettyPrint(destination.LaneConfig))
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return "\n" + string(s)
}

func deploySourceContracts(t *testing.T, source *EvmDeploymentConfig, destChainId uint64, destSupportedTokens map[Token]EVMBridgedToken) {
	// Updates source.OnRamp if any new contracts are deployed
	deployOnRamp(t, source, destChainId, destSupportedTokens)

	// Skip if we reuse both the onRamp and the token pools
	if source.DeploySettings.DeployRamp || source.DeploySettings.DeployTokenPools {
		setOnRampOnTokenPools(t, source)
	}
	source.Logger.Infof("%s contracts deployed as source chain", helpers.ChainName(int64(source.ChainConfig.ChainId)))
}

func deployDestinationContracts(t *testing.T, client *EvmDeploymentConfig, sourceChainId uint64, onRamp common.Address, supportedTokens map[Token]EVMBridgedToken) {
	// Updates destClient.LaneConfig.CommitStore if any new contracts are deployed
	deployCommitStore(t, client, sourceChainId, onRamp)

	// Updates destClient.LaneConfig.OffRamp if any new contracts are deployed
	deployOffRamp(t, client, sourceChainId, supportedTokens, onRamp)

	setFeeManagerUpdater(t, client)

	if client.DeploySettings.DeployRamp || client.DeploySettings.DeployTokenPools {
		setOffRampOnTokenPools(t, client)
	}

	// Updates destClient.ReceiverDapp if any new contracts are deployed
	deployReceiverDapp(t, client)
	client.Logger.Infof("%s contracts fully deployed as destination chain", helpers.ChainName(int64(client.ChainConfig.ChainId)))
}

func deployOnRamp(t *testing.T, client *EvmDeploymentConfig, destChainId uint64, destSupportedTokens map[Token]EVMBridgedToken) *evm_2_evm_onramp.EVM2EVMOnRamp {
	if !client.DeploySettings.DeployRamp {
		client.Logger.Infof("Skipping OnRamp deployment, using onRamp on %s", client.LaneConfig.OnRamp)
		onRamp, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(client.LaneConfig.OnRamp, client.Client)
		shared.RequireNoError(t, err)
		return onRamp
	}

	var bridgeTokens, tokenPools []common.Address
	for token, tokenConfig := range client.ChainConfig.SupportedTokens {
		if _, ok := destSupportedTokens[token]; !ok {
			// If the token is not supported on the destination chain we
			// should not enable it for this ramp. If we enable the token,
			// txs could be sent but not executed, keeping the tokens in limbo.
			continue
		}

		bridgeTokens = append(bridgeTokens, tokenConfig.Token)
		tokenPools = append(tokenPools, tokenConfig.Pool)
	}

	var feeTokenConfig []evm_2_evm_onramp.IEVM2EVMOnRampFeeTokenConfigArgs

	for _, feeToken := range client.ChainConfig.FeeTokens {
		feeTokenConfig = append(feeTokenConfig, evm_2_evm_onramp.IEVM2EVMOnRampFeeTokenConfigArgs{
			Token:           client.ChainConfig.SupportedTokens[feeToken].Token,
			Multiplier:      1e18,
			FeeAmount:       big.NewInt(100e9),
			DestGasOverhead: 0,
		})
	}

	client.Logger.Infof("Deploying OnRamp: destinationChains %+v, bridgeTokens %+v, poolAddresses %+v", destChainId, bridgeTokens, tokenPools)
	onRampAddress, tx, _, err := evm_2_evm_onramp.DeployEVM2EVMOnRamp(
		client.Owner,               // user
		client.Client,              // client
		client.ChainConfig.ChainId, // source chain id
		destChainId,                // destinationChainId
		bridgeTokens,               // tokens
		tokenPools,                 // pools
		[]common.Address{},         // allow list
		client.ChainConfig.Afn,     // AFN
		evm_2_evm_onramp.IEVM2EVMOnRampOnRampConfig{
			MaxDataSize:     1e6,
			MaxTokensLength: 5,
			MaxGasLimit:     ccip.GasLimitPerTx,
		},
		evm_2_evm_onramp.IAggregateRateLimiterRateLimiterConfig{
			Capacity: new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e9)),
			Rate:     new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e5)),
			Admin:    client.Owner.From,
		},
		client.ChainConfig.Router,
		client.ChainConfig.FeeManager,
		feeTokenConfig,
	)
	shared.RequireNoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)

	onRamp, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(onRampAddress, client.Client)
	shared.RequireNoError(t, err)
	client.Logger.Infof(fmt.Sprintf("Onramp deployed on %s in tx %s", onRampAddress.String(), helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash())))
	client.LaneConfig.OnRamp = onRampAddress

	setOnRampOnRouter(t, client, destChainId)

	// Prices are used by the rate limiter and dictate what tokens are supported
	tx, err = onRamp.SetPrices(client.Owner, []common.Address{client.ChainConfig.SupportedTokens[LINK].Token}, []*big.Int{big.NewInt(10)})
	shared.RequireNoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)

	return onRamp
}

func deployOffRamp(t *testing.T, client *EvmDeploymentConfig, sourceChainId uint64, sourceTokens map[Token]EVMBridgedToken, onRamp common.Address) *evm_2_evm_offramp.EVM2EVMOffRamp {
	if !client.DeploySettings.DeployRamp {
		client.Logger.Infof("Skipping OffRamp deployment, using offRamp on %s", client.LaneConfig.OnRamp)
		offRamp, err := evm_2_evm_offramp.NewEVM2EVMOffRamp(client.LaneConfig.OffRamp, client.Client)
		shared.RequireNoError(t, err)
		return offRamp
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
		sourceChainId,
		client.ChainConfig.ChainId,
		onRamp,
		evm_2_evm_offramp.IEVM2EVMOffRampOffRampConfig{
			Router:                                  client.ChainConfig.Router,
			CommitStore:                             client.LaneConfig.CommitStore,
			FeeManager:                              client.ChainConfig.FeeManager,
			ExecutionDelaySeconds:                   60,
			MaxDataSize:                             1e5,
			MaxTokensLength:                         15,
			PermissionLessExecutionThresholdSeconds: 60,
		},
		client.ChainConfig.Afn,
		syncedSourceTokens,
		syncedDestPools,
		evm_2_evm_offramp.IAggregateRateLimiterRateLimiterConfig{
			Capacity: new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e9)),
			Rate:     new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e5)),
			Admin:    client.Owner.From,
		},
	)
	shared.RequireNoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)

	client.Logger.Infof("OffRamp contract deployed on %s in tx: %s", offRampAddress.Hex(), helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash()))
	client.LaneConfig.OffRamp = offRampAddress
	offRamp, err := evm_2_evm_offramp.NewEVM2EVMOffRamp(client.LaneConfig.OffRamp, client.Client)
	shared.RequireNoError(t, err)

	// Prices are used by the rate limiter and dictate what tokens are supported
	tx, err = offRamp.SetPrices(client.Owner, []common.Address{client.ChainConfig.SupportedTokens[LINK].Token}, []*big.Int{big.NewInt(10)})
	shared.RequireNoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)

	client.Logger.Infof(fmt.Sprintf("Offramp configured for already deployed router in tx %s", helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash())))

	return offRamp
}

func deployCommitStore(t *testing.T, client *EvmDeploymentConfig, sourceChainId uint64, onRamp common.Address) *commit_store.CommitStore {
	if !client.DeploySettings.DeployCommitStore {
		client.Logger.Infof("Skipping CommitStore deployment, using CommitStore on %s", client.LaneConfig.CommitStore)
		commitStore, err := commit_store.NewCommitStore(client.LaneConfig.CommitStore, client.Client)
		shared.RequireNoError(t, err)
		return commitStore
	}

	client.Logger.Infof("Deploying commitStore")
	commitStoreAddress, tx, _, err := commit_store.DeployCommitStore(
		client.Owner,  // user
		client.Client, // client
		commit_store.ICommitStoreCommitStoreConfig{
			ChainId:       client.ChainConfig.ChainId,
			SourceChainId: sourceChainId,
			OnRamp:        onRamp,
		},
		client.ChainConfig.Afn, // AFN address
		1,                      // Minimum sequence number
	)
	shared.RequireNoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	client.Logger.Infof("CommitStore deployed on %s in tx: %s", commitStoreAddress.Hex(), helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash()))
	client.LaneConfig.CommitStore = commitStoreAddress

	commitStore, err := commit_store.NewCommitStore(commitStoreAddress, client.Client)
	shared.RequireNoError(t, err)
	return commitStore
}

func deployReceiverDapp(t *testing.T, client *EvmDeploymentConfig) *receiver_dapp.ReceiverDapp {
	client.Logger.Infof("Deploying ReceiverDapp")
	receiverDappAddress, tx, _, err := receiver_dapp.DeployReceiverDapp(client.Owner, client.Client, client.ChainConfig.Router)
	shared.RequireNoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	client.Logger.Infof("Offramp receiver dapp deployed on %s in tx: %s", receiverDappAddress.Hex(), helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash()))
	client.LaneConfig.ReceiverDapp = receiverDappAddress

	receiverDapp, err := receiver_dapp.NewReceiverDapp(receiverDappAddress, client.Client)
	shared.RequireNoError(t, err)

	return receiverDapp
}

func DeployPingPongDapps(t *testing.T, sourceClient *EvmDeploymentConfig, destClient *EvmDeploymentConfig) {
	fundingAmount := big.NewInt(1e18)

	if sourceClient.DeploySettings.DeployPingPongDapp {
		feeToken := sourceClient.ChainConfig.SupportedTokens[WETH].Token
		sourceClient.Logger.Infof("Deploying source chain ping pong dapp")

		pingPongDappAddress, tx, _, err := ping_pong_demo.DeployPingPongDemo(
			sourceClient.Owner,
			sourceClient.Client,
			sourceClient.ChainConfig.Router,
			feeToken,
		)
		shared.RequireNoError(t, err)

		shared.WaitForMined(t, sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
		sourceClient.Logger.Infof("Ping pong deployed on %s in tx: %s", pingPongDappAddress.Hex(), helpers.ExplorerLink(int64(sourceClient.ChainConfig.ChainId), tx.Hash()))

		sourceClient.LaneConfig.PingPongDapp = pingPongDappAddress
		FundPingPong(t, sourceClient, fundingAmount, feeToken)
	}

	if destClient.DeploySettings.DeployPingPongDapp {
		feeToken := destClient.ChainConfig.SupportedTokens[WAVAX].Token
		destClient.Logger.Infof("Deploying destination chain ping pong dapp")

		pingPongDappAddress, tx, _, err := ping_pong_demo.DeployPingPongDemo(
			destClient.Owner,
			destClient.Client,
			destClient.ChainConfig.Router,
			feeToken,
		)
		shared.RequireNoError(t, err)

		shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
		destClient.Logger.Infof("Ping pong deployed on %s in tx: %s", pingPongDappAddress.Hex(), helpers.ExplorerLink(int64(destClient.ChainConfig.ChainId), tx.Hash()))

		destClient.LaneConfig.PingPongDapp = pingPongDappAddress
		FundPingPong(t, destClient, fundingAmount, feeToken)
	}

	if sourceClient.DeploySettings.DeployPingPongDapp || destClient.DeploySettings.DeployPingPongDapp {
		pingDapp, err := ping_pong_demo.NewPingPongDemo(sourceClient.LaneConfig.PingPongDapp, sourceClient.Client)
		shared.RequireNoError(t, err)

		tx, err := pingDapp.SetCounterpart(sourceClient.Owner, destClient.ChainConfig.ChainId, destClient.LaneConfig.PingPongDapp)
		shared.RequireNoError(t, err)
		shared.WaitForMined(t, sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
		sourceClient.Logger.Infof("Ping pong dapp configured in tx: %s", helpers.ExplorerLink(int64(sourceClient.ChainConfig.ChainId), tx.Hash()))

		pongDapp, err := ping_pong_demo.NewPingPongDemo(destClient.LaneConfig.PingPongDapp, destClient.Client)
		shared.RequireNoError(t, err)

		tx, err = pongDapp.SetCounterpart(destClient.Owner, sourceClient.ChainConfig.ChainId, sourceClient.LaneConfig.PingPongDapp)
		shared.RequireNoError(t, err)
		shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
		sourceClient.Logger.Infof("Ping pong dapp configured in tx: %s", helpers.ExplorerLink(int64(sourceClient.ChainConfig.ChainId), tx.Hash()))
	} else {
		sourceClient.Logger.Infof("Skipping ping pong deployment")
	}
}

func FundPingPong(t *testing.T, client *EvmDeploymentConfig, fundingAmount *big.Int, tokenAddress common.Address) {
	pingDapp, err := ping_pong_demo.NewPingPongDemo(client.LaneConfig.PingPongDapp, client.Client)
	require.NoError(t, err)

	linkToken, err := link_token_interface.NewLinkToken(tokenAddress, client.Client)
	require.NoError(t, err)

	tx, err := linkToken.Approve(client.Owner, client.LaneConfig.PingPongDapp, fundingAmount)
	require.NoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)

	tx, err = pingDapp.Fund(client.Owner, fundingAmount)
	require.NoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	client.Logger.Infof("Ping pong funded with %s in tx: %s", fundingAmount.String(), helpers.ExplorerLink(int64(client.ChainConfig.ChainId), tx.Hash()))
}

func deployGovernanceDapps(t *testing.T, sourceClient *EvmDeploymentConfig, destClient *EvmDeploymentConfig) {
	feeConfig := governance_dapp.GovernanceDappFeeConfig{
		FeeAmount:      big.NewInt(10),
		ChangedAtBlock: big.NewInt(0),
	}

	if sourceClient.DeploySettings.DeployGovernanceDapp {
		sourceClient.Logger.Infof("Deploying source chain governance dapp")
		governanceDappAddress, tx, _, err := governance_dapp.DeployGovernanceDapp(
			sourceClient.Owner,
			sourceClient.Client,
			sourceClient.ChainConfig.Router,
			feeConfig,
			destClient.ChainConfig.SupportedTokens[LINK].Token)
		require.NoError(t, err)

		shared.WaitForMined(t, sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
		sourceClient.Logger.Infof("GovernanceDapp deployed on %s in tx: %s", governanceDappAddress.Hex(), helpers.ExplorerLink(int64(sourceClient.ChainConfig.ChainId), tx.Hash()))

		sourceClient.LaneConfig.GovernanceDapp = governanceDappAddress
	}

	if destClient.DeploySettings.DeployGovernanceDapp {
		destClient.Logger.Infof("Deploying destination chain governance dapp")
		governanceDappAddress, tx, _, err := governance_dapp.DeployGovernanceDapp(
			destClient.Owner,
			destClient.Client,
			destClient.ChainConfig.Router,
			feeConfig,
			destClient.ChainConfig.SupportedTokens[LINK].Token)
		require.NoError(t, err)

		shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
		destClient.Logger.Infof("GovernanceDapp deployed on %s in tx: %s", governanceDappAddress.Hex(), helpers.ExplorerLink(int64(destClient.ChainConfig.ChainId), tx.Hash()))

		destClient.LaneConfig.GovernanceDapp = governanceDappAddress
	}

	if sourceClient.DeploySettings.DeployGovernanceDapp || destClient.DeploySettings.DeployGovernanceDapp {
		governanceDapp, err := governance_dapp.NewGovernanceDapp(sourceClient.LaneConfig.GovernanceDapp, sourceClient.Client)
		require.NoError(t, err)

		governanceClone := governance_dapp.GovernanceDappCrossChainClone{
			ChainId:         destClient.ChainConfig.ChainId,
			ContractAddress: destClient.LaneConfig.GovernanceDapp,
		}

		tx, err := governanceDapp.AddClone(sourceClient.Owner, governanceClone)
		require.NoError(t, err)
		shared.WaitForMined(t, sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
		sourceClient.Logger.Infof("GovernanceDapp configured in tx: %s", helpers.ExplorerLink(int64(sourceClient.ChainConfig.ChainId), tx.Hash()))
	}
}
