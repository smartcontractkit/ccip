package rhea

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_free_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_subscription_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_subscription_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/governance_dapp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/mock_afn_contract"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/ping_pong_demo"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/receiver_dapp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/simple_message_receiver"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/subscription_sender_dapp"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/shared"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
)

// DeploySubscriptionContracts will deploy all source and Destination chain contracts using the
// owner key. Only run this of the currently deployed contracts are outdated or
// when initializing a new chain.
func DeploySubscriptionContracts(t *testing.T, source *EvmDeploymentConfig, destination *EvmDeploymentConfig) {
	// After running this code please update the configuration to reflect the newly
	// deployed contract addresses.
	deploySourceContracts(t, source, destination.ChainConfig.ChainId)

	source.Logger.Infof("%s contracts fully deployed as source chain", helpers.ChainName(source.ChainConfig.ChainId.Int64()))

	deployDestinationContracts(t, destination, source)
	destination.Logger.Infof("%s contracts fully deployed as destination chain", helpers.ChainName(destination.ChainConfig.ChainId.Int64()))

	// Deploy onramp sender dapp
	deploySenderDapp(t, source, destination)

	// Deploy governance dapps
	deployGovernanceDapps(t, source, destination)

	DeployPingPongDapps(t, source, destination)
}

func deploySenderDapp(t *testing.T, source *EvmDeploymentConfig, dest *EvmDeploymentConfig) {
	tokenSenderAddress, tx, _, err := subscription_sender_dapp.DeploySubscriptionSenderDapp(source.Owner, source.Client, source.ChainConfig.OnRampRouter, dest.ChainConfig.ChainId)
	require.NoError(t, err)
	shared.WaitForMined(t, source.Logger, source.Client, tx.Hash(), true)
	source.Logger.Infof("Token sender dapp deployed on %s in tx: %s", tokenSenderAddress.Hex(), helpers.ExplorerLink(source.ChainConfig.ChainId.Int64(), tx.Hash()))
	source.LaneConfig.TokenSender = tokenSenderAddress

	createDestSubscription(t, dest, dest.LaneConfig.ReceiverDapp, []common.Address{source.LaneConfig.TokenSender})
}

func deploySourceContracts(t *testing.T, source *EvmDeploymentConfig, destChainId *big.Int) {
	// Updates source.TokenPools if any new contracts are deployed
	deployNativeTokenPool(t, source)
	// Updates source.AFN if any new contracts are deployed
	deployAFN(t, source)
	// Updates source.ChainConfig.OnRampRouter if any new contracts are deployed
	deployOnRampRouter(t, source)
	// Updates source.OnRamp if any new contracts are deployed
	deployOnRamp(t, source, destChainId)

	// Skip if we reuse both the onRamp and the token pools
	if source.DeploySettings.DeployRamp || source.DeploySettings.DeployTokenPools {
		setOnRampOnTokenPools(t, source)
	}
}

func deployDestinationContracts(t *testing.T, destClient *EvmDeploymentConfig, sourceClient *EvmDeploymentConfig) {
	// Updates destClient.ChainConfig.Afn if any new contracts are deployed
	deployAFN(t, destClient)
	// Updates destclient.ChainConfig.TokenPools if any new contracts are deployed
	deployNativeTokenPool(t, destClient)
	// Updates destClient.LaneConfig.CommitStore if any new contracts are deployed
	deployCommitStore(t, destClient, sourceClient)

	// Updates destClient.LaneConfig.OffRamp if any new contracts are deployed
	deployOffRamp(t, destClient, sourceClient)
	// Updates destClient.ChainConfig.OffRampRouter if any new contracts are deployed
	deployOffRampRouter(t, destClient)

	if destClient.DeploySettings.DeployRamp || destClient.DeploySettings.DeployRouter {
		setOffRampRouterOnOffRamp(t, destClient)
	}

	if destClient.DeploySettings.DeployRamp || destClient.DeploySettings.DeployTokenPools {
		setOffRampOnTokenPools(t, destClient)
	}

	// Updates destClient.ReceiverDapp if any new contracts are deployed
	deployReceiverDapp(t, destClient)

	// Deploy offramp contract message receiver
	messageReceiverAddress, tx, _, err := simple_message_receiver.DeploySimpleMessageReceiver(destClient.Owner, destClient.Client)
	require.NoError(t, err)
	shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
	destClient.Logger.Infof("Offramp message receiver deployed on %s in tx: %s", messageReceiverAddress.Hex(), helpers.ExplorerLink(destClient.ChainConfig.ChainId.Int64(), tx.Hash()))
	destClient.LaneConfig.MessageReceiver = messageReceiverAddress

}

func deployOnRampRouter(t *testing.T, client *EvmDeploymentConfig) *evm_2_any_subscription_onramp_router.EVM2AnySubscriptionOnRampRouter {
	if !client.DeploySettings.DeployRouter {
		client.Logger.Infof("Skipping OnRampRouter deployment, using OnRampRouter on %s", client.ChainConfig.OnRampRouter)
		onRampRouter, err := evm_2_any_subscription_onramp_router.NewEVM2AnySubscriptionOnRampRouter(client.LaneConfig.OnRamp, client.Client)
		require.NoError(t, err)
		return onRampRouter
	}
	client.Logger.Infof("Deploying OnRampRouter")

	onRampRouterAddress, tx, _, err := evm_2_any_subscription_onramp_router.DeployEVM2AnySubscriptionOnRampRouter(
		client.Owner,
		client.Client,
		evm_2_any_subscription_onramp_router.EVM2AnySubscriptionOnRampRouterInterfaceRouterConfig{
			Fee:      big.NewInt(0),
			FeeToken: client.ChainConfig.LinkToken,
			FeeAdmin: client.Owner.From,
		})
	require.NoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	client.ChainConfig.OnRampRouter = onRampRouterAddress

	client.Logger.Infof(fmt.Sprintf("OnRampRouter deployed on %s in tx %s", onRampRouterAddress.String(), helpers.ExplorerLink(client.ChainConfig.ChainId.Int64(), tx.Hash())))

	onRampRouter, err := evm_2_any_subscription_onramp_router.NewEVM2AnySubscriptionOnRampRouter(client.LaneConfig.OnRamp, client.Client)
	require.NoError(t, err)

	// TODO FUNDING? Reverts...
	//fundingAmount := big.NewInt(0)

	//sourceLinkToken, err := link_token_interface.NewLinkToken(client.ChainConfig.LinkToken, client.Client)
	//require.NoError(t, err)
	//
	//tx, err = sourceLinkToken.Approve(client.Owner, onRampRouter.Address(), fundingAmount)
	//require.NoError(t, err)
	//shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	//client.Logger.Infof(fmt.Sprintf("Approved link for onramp subscription funding in tx %s", helpers.ExplorerLink(Client.ChainConfig.ChainId.Int64(), tx.Hash())))
	//
	//// Fund subscription with 0.01 LINK, enough for 1000 txs at 1e13 each
	//tx, err = onRampRouter.FundSubscription(client.Owner, fundingAmount)
	//require.NoError(t, err)
	//shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	//client.Logger.Infof(fmt.Sprintf("Funded onramp subscription in tx %s", helpers.ExplorerLink(Client.ChainConfig.ChainId.Int64(), tx.Hash())))

	return onRampRouter
}

func deployOnRamp(t *testing.T, client *EvmDeploymentConfig, destChainId *big.Int) *evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRamp {
	if !client.DeploySettings.DeployRamp {
		client.Logger.Infof("Skipping OnRamp deployment, using onRamp on %s", client.LaneConfig.OnRamp)
		onRamp, err := evm_2_evm_subscription_onramp.NewEVM2EVMSubscriptionOnRamp(client.LaneConfig.OnRamp, client.Client)
		require.NoError(t, err)
		return onRamp
	}

	var bridgeTokens, tokenPools []common.Address
	for token, tokenConfig := range client.ChainConfig.SupportedTokens {
		bridgeTokens = append(bridgeTokens, token)
		tokenPools = append(tokenPools, tokenConfig.Pool)
	}

	client.Logger.Infof("Deploying OnRamp: destinationChains %+v, bridgeTokens %+v, poolAddresses %+v", destChainId, bridgeTokens, tokenPools)
	onRampAddress, tx, _, err := evm_2_evm_subscription_onramp.DeployEVM2EVMSubscriptionOnRamp(
		client.Owner,               // user
		client.Client,              // client
		client.ChainConfig.ChainId, // source chain id
		destChainId,                // destinationChainId
		bridgeTokens,               // tokens
		tokenPools,                 // pools
		[]common.Address{},         // allow list
		client.ChainConfig.Afn,     // AFN
		evm_2_evm_subscription_onramp.BaseOnRampInterfaceOnRampConfig{
			CommitFeeJuels:  0,
			MaxDataSize:     1e6,
			MaxTokensLength: 5,
			MaxGasLimit:     ccip.GasLimitPerTx,
		},
		evm_2_evm_subscription_onramp.AggregateRateLimiterInterfaceRateLimiterConfig{
			Capacity: new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e9)),
			Rate:     new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e5)),
		},
		client.Owner.From,
		client.ChainConfig.OnRampRouter,
	)
	require.NoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)

	onRamp, err := evm_2_evm_subscription_onramp.NewEVM2EVMSubscriptionOnRamp(onRampAddress, client.Client)
	require.NoError(t, err)
	client.Logger.Infof(fmt.Sprintf("Onramp deployed on %s in tx %s", onRampAddress.String(), helpers.ExplorerLink(client.ChainConfig.ChainId.Int64(), tx.Hash())))
	client.LaneConfig.OnRamp = onRampAddress

	setOnRampOnOnRampRouter(t, client, destChainId)

	// Prices are used by the rate limiter and dictate what tokens are supported
	tx, err = onRamp.SetPrices(client.Owner, []common.Address{client.ChainConfig.LinkToken}, []*big.Int{big.NewInt(10)})
	require.NoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)

	return onRamp
}

func deployOffRamp(t *testing.T, destClient *EvmDeploymentConfig, sourceClient *EvmDeploymentConfig) *any_2_evm_free_offramp.EVM2EVMFreeOffRamp {
	if !destClient.DeploySettings.DeployRamp {
		destClient.Logger.Infof("Skipping OffRamp deployment, using offRamp on %s", destClient.LaneConfig.OnRamp)
		offRamp, err := any_2_evm_free_offramp.NewEVM2EVMFreeOffRamp(destClient.LaneConfig.OffRamp, destClient.Client)
		require.NoError(t, err)
		return offRamp
	}
	var sourceTokens, tokenPools []common.Address
	for _, tokenConfig := range destClient.ChainConfig.SupportedTokens {
		tokenPools = append(tokenPools, tokenConfig.Pool)
	}
	for token := range sourceClient.ChainConfig.SupportedTokens {
		sourceTokens = append(sourceTokens, token)
	}

	destClient.Logger.Infof("Deploying OffRamp")
	offRampAddress, tx, _, err := any_2_evm_free_offramp.DeployEVM2EVMFreeOffRamp(
		destClient.Owner,
		destClient.Client,
		sourceClient.ChainConfig.ChainId,
		destClient.ChainConfig.ChainId,
		any_2_evm_free_offramp.BaseOffRampInterfaceOffRampConfig{
			ExecutionDelaySeconds:                   60,
			MaxDataSize:                             1e5,
			MaxTokensLength:                         15,
			PermissionLessExecutionThresholdSeconds: 60,
		},
		sourceClient.LaneConfig.OnRamp,
		destClient.LaneConfig.CommitStore,
		destClient.ChainConfig.Afn,
		sourceTokens,
		tokenPools,
		any_2_evm_free_offramp.AggregateRateLimiterInterfaceRateLimiterConfig{
			Capacity: new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e9)),
			Rate:     new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e5)),
		},
		destClient.Owner.From)
	require.NoError(t, err)
	shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)

	destClient.Logger.Infof("OffRamp contract deployed on %s in tx: %s", offRampAddress.Hex(), helpers.ExplorerLink(destClient.ChainConfig.ChainId.Int64(), tx.Hash()))
	destClient.LaneConfig.OffRamp = offRampAddress
	offRamp, err := any_2_evm_free_offramp.NewEVM2EVMFreeOffRamp(destClient.LaneConfig.OffRamp, destClient.Client)
	require.NoError(t, err)

	// Prices are used by the rate limiter and dictate what tokens are supported
	tx, err = offRamp.SetPrices(destClient.Owner, []common.Address{destClient.ChainConfig.LinkToken}, []*big.Int{big.NewInt(10)})
	require.NoError(t, err)
	shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)

	return offRamp
}

func deployOffRampRouter(t *testing.T, destClient *EvmDeploymentConfig) *any_2_evm_subscription_offramp_router.Any2EVMSubscriptionOffRampRouter {
	if !destClient.DeploySettings.DeployRouter {
		destClient.Logger.Infof("Skipping OffRampRouter deployment, using OffRampRouter on %s", destClient.ChainConfig.OffRampRouter)
		offRampRouter, err := any_2_evm_subscription_offramp_router.NewAny2EVMSubscriptionOffRampRouter(destClient.ChainConfig.OffRampRouter, destClient.Client)
		require.NoError(t, err)

		if destClient.DeploySettings.DeployRamp {
			tx, err2 := offRampRouter.AddOffRamp(destClient.Owner, destClient.LaneConfig.OffRamp)
			require.NoError(t, err2)
			shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
			destClient.Logger.Infof(fmt.Sprintf("Offramp configured for alraedy deployed router  in tx %s", helpers.ExplorerLink(destClient.ChainConfig.ChainId.Int64(), tx.Hash())))
		}

		return offRampRouter
	}

	destClient.Logger.Infof("Deploying OffRampRouter")
	offRampRouterAddress, tx, _, err := any_2_evm_subscription_offramp_router.DeployAny2EVMSubscriptionOffRampRouter(
		destClient.Owner,
		destClient.Client,
		[]common.Address{destClient.LaneConfig.OffRamp},
		any_2_evm_subscription_offramp_router.SubscriptionInterfaceSubscriptionConfig{
			SetSubscriptionSenderDelay: 10,
			WithdrawalDelay:            10,
			FeeToken:                   destClient.ChainConfig.LinkToken,
		},
	)
	require.NoError(t, err)
	shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
	destClient.ChainConfig.OffRampRouter = offRampRouterAddress

	destClient.Logger.Infof(fmt.Sprintf("OffRampRouter deployed on %s in tx %s", offRampRouterAddress.String(), helpers.ExplorerLink(destClient.ChainConfig.ChainId.Int64(), tx.Hash())))

	offRampRouter, err := any_2_evm_subscription_offramp_router.NewAny2EVMSubscriptionOffRampRouter(destClient.ChainConfig.OffRampRouter, destClient.Client)
	require.NoError(t, err)

	return offRampRouter
}

func deployCommitStore(t *testing.T, destClient *EvmDeploymentConfig, sourceClient *EvmDeploymentConfig) *commit_store.CommitStore {
	if !destClient.DeploySettings.DeployCommitStore {
		destClient.Logger.Infof("Skipping CommitStore deployment, using CommitStore on %s", destClient.LaneConfig.CommitStore)
		commitStore, err := commit_store.NewCommitStore(destClient.LaneConfig.CommitStore, destClient.Client)
		require.NoError(t, err)
		return commitStore
	}

	destClient.Logger.Infof("Deploying commitStore")

	commitStoreAddress, tx, _, err := commit_store.DeployCommitStore(
		destClient.Owner,                 // user
		destClient.Client,                // client
		destClient.ChainConfig.ChainId,   // dest chain id
		sourceClient.ChainConfig.ChainId, // source chain id
		destClient.ChainConfig.Afn,       // AFN address
		commit_store.CommitStoreInterfaceCommitStoreConfig{
			OnRamps:          []common.Address{sourceClient.LaneConfig.OnRamp},
			MinSeqNrByOnRamp: []uint64{1},
		},
	)
	require.NoError(t, err)
	shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
	destClient.Logger.Infof("commitStore deployed on %s in tx: %s", commitStoreAddress.Hex(), helpers.ExplorerLink(destClient.ChainConfig.ChainId.Int64(), tx.Hash()))
	destClient.LaneConfig.CommitStore = commitStoreAddress

	commitStore, err := commit_store.NewCommitStore(commitStoreAddress, destClient.Client)
	require.NoError(t, err)
	return commitStore
}

func deployReceiverDapp(t *testing.T, destClient *EvmDeploymentConfig) *receiver_dapp.ReceiverDapp {
	destClient.Logger.Infof("Deploying ReceiverDapp")
	receiverDappAddress, tx, _, err := receiver_dapp.DeployReceiverDapp(destClient.Owner, destClient.Client, destClient.ChainConfig.OffRampRouter)
	require.NoError(t, err)
	shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
	destClient.Logger.Infof("Offramp receiver dapp deployed on %s in tx: %s", receiverDappAddress.Hex(), helpers.ExplorerLink(destClient.ChainConfig.ChainId.Int64(), tx.Hash()))
	destClient.LaneConfig.ReceiverDapp = receiverDappAddress

	receiverDapp, err := receiver_dapp.NewReceiverDapp(receiverDappAddress, destClient.Client)
	require.NoError(t, err)

	return receiverDapp
}

func deployNativeTokenPool(t *testing.T, client *EvmDeploymentConfig) []*native_token_pool.NativeTokenPool {
	var pools []*native_token_pool.NativeTokenPool

	for token, tokenConfig := range client.ChainConfig.SupportedTokens {
		if client.DeploySettings.DeployTokenPools {
			client.Logger.Infof("Deploying token pool for token %s", token.Hex())
			tokenPoolAddress, tx, _, err := native_token_pool.DeployNativeTokenPool(client.Owner, client.Client, token)
			require.NoError(t, err)
			shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
			client.Logger.Infof("Native token pool deployed on %s in tx %s", tokenPoolAddress, helpers.ExplorerLink(client.ChainConfig.ChainId.Int64(), tx.Hash()))
			pool, err := native_token_pool.NewNativeTokenPool(tokenPoolAddress, client.Client)
			require.NoError(t, err)
			fillPoolWithTokens(t, client, pool)
			pools = append(pools, pool)
			client.ChainConfig.SupportedTokens[token] = EVMBridgedToken{
				Pool:  tokenPoolAddress,
				Price: big.NewInt(1),
			}
		} else {
			if tokenConfig.Pool.Hex() == "0x0000000000000000000000000000000000000000" {
				t.Error("deploy new lock unlock pool set to false but no lock unlock pool given in config")
			}
			pool, err := native_token_pool.NewNativeTokenPool(tokenConfig.Pool, client.Client)
			require.NoError(t, err)
			client.Logger.Infof("Lock unlock pool loaded from: %s", pool.Address().Hex())
			pools = append(pools, pool)
		}
	}

	return pools
}

func DeployPingPongDapps(t *testing.T, sourceClient *EvmDeploymentConfig, destClient *EvmDeploymentConfig) {
	if sourceClient.DeploySettings.DeployPingPongDapp {
		sourceClient.Logger.Infof("Deploying source chain ping pong dapp")

		pingPongDappAddress, tx, _, err := ping_pong_demo.DeployPingPongDemo(sourceClient.Owner, sourceClient.Client, sourceClient.ChainConfig.OffRampRouter, sourceClient.ChainConfig.OnRampRouter)
		require.NoError(t, err)

		shared.WaitForMined(t, sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
		sourceClient.Logger.Infof("Ping pong deployed on %s in tx: %s", pingPongDappAddress.Hex(), helpers.ExplorerLink(sourceClient.ChainConfig.ChainId.Int64(), tx.Hash()))

		sourceClient.LaneConfig.PingPongDapp = pingPongDappAddress
	}

	if destClient.DeploySettings.DeployPingPongDapp {
		destClient.Logger.Infof("Deploying destination chain ping pong dapp")

		pingPongDappAddress, tx, _, err := ping_pong_demo.DeployPingPongDemo(destClient.Owner, destClient.Client, destClient.ChainConfig.OffRampRouter, destClient.ChainConfig.OnRampRouter)
		require.NoError(t, err)

		shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
		destClient.Logger.Infof("Ping pong deployed on %s in tx: %s", pingPongDappAddress.Hex(), helpers.ExplorerLink(destClient.ChainConfig.ChainId.Int64(), tx.Hash()))

		destClient.LaneConfig.PingPongDapp = pingPongDappAddress
	}

	if sourceClient.DeploySettings.DeployPingPongDapp || destClient.DeploySettings.DeployPingPongDapp {
		pingDapp, err := ping_pong_demo.NewPingPongDemo(sourceClient.LaneConfig.PingPongDapp, sourceClient.Client)
		require.NoError(t, err)

		tx, err := pingDapp.SetCounterpart(sourceClient.Owner, destClient.ChainConfig.ChainId, destClient.LaneConfig.PingPongDapp)
		require.NoError(t, err)
		shared.WaitForMined(t, sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
		sourceClient.Logger.Infof("Ping pong dapp configured in tx: %s", helpers.ExplorerLink(sourceClient.ChainConfig.ChainId.Int64(), tx.Hash()))

		pongDapp, err := ping_pong_demo.NewPingPongDemo(destClient.LaneConfig.PingPongDapp, destClient.Client)
		require.NoError(t, err)

		tx, err = pongDapp.SetCounterpart(destClient.Owner, sourceClient.ChainConfig.ChainId, sourceClient.LaneConfig.PingPongDapp)
		require.NoError(t, err)
		shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
		sourceClient.Logger.Infof("Ping pong dapp configured in tx: %s", helpers.ExplorerLink(sourceClient.ChainConfig.ChainId.Int64(), tx.Hash()))

		createDestSubscription(t, sourceClient, sourceClient.LaneConfig.PingPongDapp, []common.Address{destClient.LaneConfig.PingPongDapp})
		createDestSubscription(t, destClient, destClient.LaneConfig.PingPongDapp, []common.Address{sourceClient.LaneConfig.PingPongDapp})
	} else {
		sourceClient.Logger.Infof("Skipping ping pong deployment")
	}
}

func deployGovernanceDapps(t *testing.T, sourceClient *EvmDeploymentConfig, destClient *EvmDeploymentConfig) {
	feeConfig := governance_dapp.GovernanceDappFeeConfig{
		FeeAmount:           big.NewInt(10),
		SubscriptionManager: sourceClient.Owner.From,
		ChangedAtBlock:      big.NewInt(0),
	}

	if sourceClient.DeploySettings.DeployGovernanceDapp {
		sourceClient.Logger.Infof("Deploying source chain governance dapp")
		governanceDappAddress, tx, _, err := governance_dapp.DeployGovernanceDapp(sourceClient.Owner, sourceClient.Client, sourceClient.ChainConfig.OffRampRouter, sourceClient.ChainConfig.OnRampRouter, feeConfig)
		require.NoError(t, err)

		shared.WaitForMined(t, sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
		sourceClient.Logger.Infof("GovernanceDapp deployed on %s in tx: %s", governanceDappAddress.Hex(), helpers.ExplorerLink(sourceClient.ChainConfig.ChainId.Int64(), tx.Hash()))

		sourceClient.LaneConfig.GovernanceDapp = governanceDappAddress
	}

	if destClient.DeploySettings.DeployGovernanceDapp {
		destClient.Logger.Infof("Deploying destination chain governance dapp")
		governanceDappAddress, tx, _, err := governance_dapp.DeployGovernanceDapp(destClient.Owner, destClient.Client, destClient.ChainConfig.OffRampRouter, destClient.ChainConfig.OnRampRouter, feeConfig)
		require.NoError(t, err)

		shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
		destClient.Logger.Infof("GovernanceDapp deployed on %s in tx: %s", governanceDappAddress.Hex(), helpers.ExplorerLink(destClient.ChainConfig.ChainId.Int64(), tx.Hash()))

		destClient.LaneConfig.GovernanceDapp = governanceDappAddress

		createDestSubscription(t, destClient, destClient.LaneConfig.GovernanceDapp, []common.Address{sourceClient.LaneConfig.GovernanceDapp})
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
		sourceClient.Logger.Infof("GovernanceDapp configured in tx: %s", helpers.ExplorerLink(sourceClient.ChainConfig.ChainId.Int64(), tx.Hash()))
	}
}

func deployAFN(t *testing.T, client *EvmDeploymentConfig) *afn_contract.AFNContract {
	if client.DeploySettings.DeployAFN {
		client.Logger.Infof("Deploying AFN")
		address, tx, _, err := mock_afn_contract.DeployMockAFNContract(client.Owner, client.Client)
		require.NoError(t, err)
		shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
		client.Logger.Infof("AFN deployed on %s in tx: %s", address.Hex(), helpers.ExplorerLink(client.ChainConfig.ChainId.Int64(), tx.Hash()))
		client.ChainConfig.Afn = address

		afn, err := afn_contract.NewAFNContract(address, client.Client)
		require.NoError(t, err)
		return afn
	}
	if client.ChainConfig.Afn.Hex() == "0x0000000000000000000000000000000000000000" {
		t.Error("deploy new afn set to false but no afn given in config")
	}
	afn, err := afn_contract.NewAFNContract(client.ChainConfig.Afn, client.Client)
	require.NoError(t, err)
	client.Logger.Infof("AFN loaded from: %s", afn.Address().Hex())
	return afn
}
