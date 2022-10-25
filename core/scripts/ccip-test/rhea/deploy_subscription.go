package rhea

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_free_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/blob_verifier"
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
)

// DeploySubscriptionContracts will deploy all source and Destination chain contracts using the
// owner key. Only run this of the currently deployed contracts are outdated or
// when initializing a new chain.
func DeploySubscriptionContracts(t *testing.T, ownerKey string, sourceChain *EvmChainConfig, destChain *EvmChainConfig) {
	sourceChain.SetupChain(t, ownerKey)
	destChain.SetupChain(t, ownerKey)
	deploySourceAndDestContracts(t, sourceChain, destChain)
}

func deploySourceAndDestContracts(t *testing.T, source *EvmChainConfig, destination *EvmChainConfig) {
	// After running this code please update the configuration to reflect the newly
	// deployed contract addresses.
	deploySourceContracts(t, source, destination.ChainId)

	source.Logger.Infof("%s contracts fully deployed as source chain", helpers.ChainName(source.ChainId.Int64()))

	deployDestinationContracts(t, destination, source)
	destination.Logger.Infof("%s contracts fully deployed as destination chain", helpers.ChainName(destination.ChainId.Int64()))

	// Deploy onramp sender dapp
	deploySenderDapp(t, source, destination)

	// Deploy governance dapps
	deployGovernanceDapps(t, source, destination)

	DeployPingPongDapps(t, source, destination)
}

func deploySenderDapp(t *testing.T, source *EvmChainConfig, dest *EvmChainConfig) {
	tokenSenderAddress, tx, _, err := subscription_sender_dapp.DeploySubscriptionSenderDapp(source.Owner, source.Client, source.OnRampRouter, dest.ChainId)
	require.NoError(t, err)
	shared.WaitForMined(t, source.Logger, source.Client, tx.Hash(), true)
	source.Logger.Infof("Token sender dapp deployed on %s in tx: %s", tokenSenderAddress.Hex(), helpers.ExplorerLink(source.ChainId.Int64(), tx.Hash()))
	source.TokenSender = tokenSenderAddress

	createDestSubscription(t, dest, dest.ReceiverDapp, []common.Address{source.TokenSender})
}

func deploySourceContracts(t *testing.T, source *EvmChainConfig, destChainId *big.Int) {
	// Updates source.TokenPools if any new contracts are deployed
	deployNativeTokenPool(t, source)
	// Updates source.AFN if any new contracts are deployed
	deployAFN(t, source)
	// Updates source.OnRampRouter if any new contracts are deployed
	deployOnRampRouter(t, source)
	// Updates source.OnRamp if any new contracts are deployed
	deployOnRamp(t, source, destChainId)

	// Skip if we reuse both the onRamp and the token pools
	if source.DeploySettings.DeployRamp || source.DeploySettings.DeployTokenPools {
		setOnRampOnTokenPools(t, source)
	}
}

func deployDestinationContracts(t *testing.T, destClient *EvmChainConfig, sourceClient *EvmChainConfig) {
	// Updates destClient.AFN if any new contracts are deployed
	deployAFN(t, destClient)
	// Updates destClient.TokenPools if any new contracts are deployed
	deployNativeTokenPool(t, destClient)
	// Updates destClient.BlobVerifier if any new contracts are deployed
	deployBlobVerifier(t, destClient, sourceClient)

	// Updates destClient.OffRamp if any new contracts are deployed
	deployOffRamp(t, destClient, sourceClient)
	// Updates destClient.OffRampRouter if any new contracts are deployed
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
	destClient.Logger.Infof("Offramp message receiver deployed on %s in tx: %s", messageReceiverAddress.Hex(), helpers.ExplorerLink(destClient.ChainId.Int64(), tx.Hash()))
	destClient.MessageReceiver = messageReceiverAddress

}

func deployOnRampRouter(t *testing.T, client *EvmChainConfig) *evm_2_any_subscription_onramp_router.EVM2AnySubscriptionOnRampRouter {
	if !client.DeploySettings.DeployRouter {
		client.Logger.Infof("Skipping OnRampRouter deployment, using OnRampRouter on %s", client.OnRampRouter)
		onRampRouter, err := evm_2_any_subscription_onramp_router.NewEVM2AnySubscriptionOnRampRouter(client.OnRamp, client.Client)
		require.NoError(t, err)
		return onRampRouter
	}
	client.Logger.Infof("Deploying OnRampRouter")

	onRampRouterAddress, tx, _, err := evm_2_any_subscription_onramp_router.DeployEVM2AnySubscriptionOnRampRouter(
		client.Owner,
		client.Client,
		evm_2_any_subscription_onramp_router.EVM2AnySubscriptionOnRampRouterInterfaceRouterConfig{
			Fee:      big.NewInt(0),
			FeeToken: client.LinkToken,
			FeeAdmin: client.Owner.From,
		})
	require.NoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	client.OnRampRouter = onRampRouterAddress

	client.Logger.Infof(fmt.Sprintf("OnRampRouter deployed on %s in tx %s", onRampRouterAddress.String(), helpers.ExplorerLink(client.ChainId.Int64(), tx.Hash())))

	onRampRouter, err := evm_2_any_subscription_onramp_router.NewEVM2AnySubscriptionOnRampRouter(client.OnRamp, client.Client)
	require.NoError(t, err)

	// TODO FUNDING? Reverts...
	//fundingAmount := big.NewInt(0)

	//sourceLinkToken, err := link_token_interface.NewLinkToken(client.LinkToken, client.Client)
	//require.NoError(t, err)
	//
	//tx, err = sourceLinkToken.Approve(client.Owner, onRampRouter.Address(), fundingAmount)
	//require.NoError(t, err)
	//shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	//client.Logger.Infof(fmt.Sprintf("Approved link for onramp subscription funding in tx %s", helpers.ExplorerLink(client.ChainId.Int64(), tx.Hash())))
	//
	//// Fund subscription with 0.01 LINK, enough for 1000 txs at 1e13 each
	//tx, err = onRampRouter.FundSubscription(client.Owner, fundingAmount)
	//require.NoError(t, err)
	//shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	//client.Logger.Infof(fmt.Sprintf("Funded onramp subscription in tx %s", helpers.ExplorerLink(client.ChainId.Int64(), tx.Hash())))

	return onRampRouter
}

func deployOnRamp(t *testing.T, client *EvmChainConfig, destChainId *big.Int) *evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRamp {
	if !client.DeploySettings.DeployRamp {
		client.Logger.Infof("Skipping OnRamp deployment, using onRamp on %s", client.OnRamp)
		onRamp, err := evm_2_evm_subscription_onramp.NewEVM2EVMSubscriptionOnRamp(client.OnRamp, client.Client)
		require.NoError(t, err)
		return onRamp
	}

	client.Logger.Infof("Deploying OnRamp: destinationChains %+v, bridgeTokens %+v, poolAddresses %+v", destChainId, client.BridgeTokens, client.TokenPools)
	onRampAddress, tx, _, err := evm_2_evm_subscription_onramp.DeployEVM2EVMSubscriptionOnRamp(
		client.Owner,        // user
		client.Client,       // client
		client.ChainId,      // source chain id
		destChainId,         // destinationChainId
		client.BridgeTokens, // tokens
		client.TokenPools,   // pools
		[]common.Address{},  // allow list
		client.Afn,          // AFN
		evm_2_evm_subscription_onramp.BaseOnRampInterfaceOnRampConfig{
			RelayingFeeJuels: 0,
			MaxDataSize:      1e6,
			MaxTokensLength:  5,
		},
		evm_2_evm_subscription_onramp.AggregateRateLimiterInterfaceRateLimiterConfig{
			Capacity: big.NewInt(1e18),
			Rate:     big.NewInt(1e18),
		},
		client.Owner.From,
		client.OnRampRouter,
	)
	require.NoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)

	onRamp, err := evm_2_evm_subscription_onramp.NewEVM2EVMSubscriptionOnRamp(onRampAddress, client.Client)
	require.NoError(t, err)
	client.Logger.Infof(fmt.Sprintf("Onramp deployed on %s in tx %s", onRampAddress.String(), helpers.ExplorerLink(client.ChainId.Int64(), tx.Hash())))
	client.OnRamp = onRampAddress

	setOnRampOnOnRampRouter(t, client, destChainId)

	// Prices are used by the rate limiter and dictate what tokens are supported
	tx, err = onRamp.SetPrices(client.Owner, []common.Address{client.LinkToken}, []*big.Int{big.NewInt(10)})
	require.NoError(t, err)
	shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)

	return onRamp
}

func deployOffRamp(t *testing.T, destClient *EvmChainConfig, sourceClient *EvmChainConfig) *any_2_evm_free_offramp.EVM2EVMFreeOffRamp {
	if !destClient.DeploySettings.DeployRamp {
		destClient.Logger.Infof("Skipping OffRamp deployment, using offRamp on %s", destClient.OnRamp)
		offRamp, err := any_2_evm_free_offramp.NewEVM2EVMFreeOffRamp(destClient.OffRamp, destClient.Client)
		require.NoError(t, err)
		return offRamp
	}

	destClient.Logger.Infof("Deploying OffRamp")
	offRampAddress, tx, _, err := any_2_evm_free_offramp.DeployEVM2EVMFreeOffRamp(
		destClient.Owner,
		destClient.Client,
		sourceClient.ChainId,
		destClient.ChainId,
		any_2_evm_free_offramp.BaseOffRampInterfaceOffRampConfig{
			OnRampAddress:                           sourceClient.OnRamp,
			ExecutionDelaySeconds:                   60,
			MaxDataSize:                             1e5,
			MaxTokensLength:                         15,
			PermissionLessExecutionThresholdSeconds: 60,
		},
		destClient.BlobVerifier,
		destClient.Afn,
		sourceClient.BridgeTokens,
		destClient.TokenPools,
		any_2_evm_free_offramp.AggregateRateLimiterInterfaceRateLimiterConfig{
			Capacity: big.NewInt(1e18),
			Rate:     big.NewInt(1e18),
		},
		destClient.Owner.From)
	require.NoError(t, err)
	shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)

	destClient.Logger.Infof("OffRamp contract deployed on %s in tx: %s", offRampAddress.Hex(), helpers.ExplorerLink(destClient.ChainId.Int64(), tx.Hash()))
	destClient.OffRamp = offRampAddress
	offRamp, err := any_2_evm_free_offramp.NewEVM2EVMFreeOffRamp(destClient.OffRamp, destClient.Client)
	require.NoError(t, err)

	// Prices are used by the rate limiter and dictate what tokens are supported
	tx, err = offRamp.SetPrices(destClient.Owner, []common.Address{destClient.LinkToken}, []*big.Int{big.NewInt(10)})
	require.NoError(t, err)
	shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)

	return offRamp
}

func deployOffRampRouter(t *testing.T, destClient *EvmChainConfig) *any_2_evm_subscription_offramp_router.Any2EVMSubscriptionOffRampRouter {
	if !destClient.DeploySettings.DeployRouter {
		destClient.Logger.Infof("Skipping OffRampRouter deployment, using OffRampRouter on %s", destClient.OffRampRouter)
		offRampRouter, err := any_2_evm_subscription_offramp_router.NewAny2EVMSubscriptionOffRampRouter(destClient.OffRampRouter, destClient.Client)
		require.NoError(t, err)

		if destClient.DeploySettings.DeployRamp {
			tx, err2 := offRampRouter.AddOffRamp(destClient.Owner, destClient.OffRamp)
			require.NoError(t, err2)
			shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
			destClient.Logger.Infof(fmt.Sprintf("Offramp configured for alraedy deployed router  in tx %s", helpers.ExplorerLink(destClient.ChainId.Int64(), tx.Hash())))
		}

		return offRampRouter
	}

	destClient.Logger.Infof("Deploying OffRampRouter")
	offRampRouterAddress, tx, _, err := any_2_evm_subscription_offramp_router.DeployAny2EVMSubscriptionOffRampRouter(
		destClient.Owner,
		destClient.Client,
		[]common.Address{destClient.OffRamp},
		any_2_evm_subscription_offramp_router.SubscriptionInterfaceSubscriptionConfig{
			SetSubscriptionSenderDelay: 10,
			WithdrawalDelay:            10,
			FeeToken:                   destClient.LinkToken,
		},
	)
	require.NoError(t, err)
	shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
	destClient.OffRampRouter = offRampRouterAddress

	destClient.Logger.Infof(fmt.Sprintf("OffRampRouter deployed on %s in tx %s", offRampRouterAddress.String(), helpers.ExplorerLink(destClient.ChainId.Int64(), tx.Hash())))

	offRampRouter, err := any_2_evm_subscription_offramp_router.NewAny2EVMSubscriptionOffRampRouter(destClient.OffRampRouter, destClient.Client)
	require.NoError(t, err)

	return offRampRouter
}

func deployBlobVerifier(t *testing.T, destClient *EvmChainConfig, sourceClient *EvmChainConfig) *blob_verifier.BlobVerifier {
	if !destClient.DeploySettings.DeployBlobVerifier {
		destClient.Logger.Infof("Skipping BlobVerifier deployment, using BlobVerifier on %s", destClient.BlobVerifier)
		blobVerifier, err := blob_verifier.NewBlobVerifier(destClient.BlobVerifier, destClient.Client)
		require.NoError(t, err)
		return blobVerifier
	}

	destClient.Logger.Infof("Deploying blob verifier")

	blobVerifierAddress, tx, _, err := blob_verifier.DeployBlobVerifier(
		destClient.Owner,     // user
		destClient.Client,    // client
		destClient.ChainId,   // dest chain id
		sourceClient.ChainId, // source chain id
		destClient.Afn,       // AFN address
		blob_verifier.BlobVerifierInterfaceBlobVerifierConfig{
			OnRamps:          []common.Address{sourceClient.OnRamp},
			MinSeqNrByOnRamp: []uint64{1},
		},
	)
	require.NoError(t, err)
	shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
	destClient.Logger.Infof("Blob verifier deployed on %s in tx: %s", blobVerifierAddress.Hex(), helpers.ExplorerLink(destClient.ChainId.Int64(), tx.Hash()))
	destClient.BlobVerifier = blobVerifierAddress

	blobVerifier, err := blob_verifier.NewBlobVerifier(blobVerifierAddress, destClient.Client)
	require.NoError(t, err)
	return blobVerifier
}

func deployReceiverDapp(t *testing.T, destClient *EvmChainConfig) *receiver_dapp.ReceiverDapp {
	destClient.Logger.Infof("Deploying ReceiverDapp")
	receiverDappAddress, tx, _, err := receiver_dapp.DeployReceiverDapp(destClient.Owner, destClient.Client, destClient.OffRampRouter)
	require.NoError(t, err)
	shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
	destClient.Logger.Infof("Offramp receiver dapp deployed on %s in tx: %s", receiverDappAddress.Hex(), helpers.ExplorerLink(destClient.ChainId.Int64(), tx.Hash()))
	destClient.ReceiverDapp = receiverDappAddress

	receiverDapp, err := receiver_dapp.NewReceiverDapp(receiverDappAddress, destClient.Client)
	require.NoError(t, err)

	return receiverDapp
}

func deployNativeTokenPool(t *testing.T, client *EvmChainConfig) []*native_token_pool.NativeTokenPool {
	var pools []*native_token_pool.NativeTokenPool
	var poolAddresses []common.Address

	for i, bridgeToken := range client.BridgeTokens {
		if client.DeploySettings.DeployTokenPools {
			client.Logger.Infof("Deploying token pool for token %s", bridgeToken.Hex())
			tokenPoolAddress, tx, _, err := native_token_pool.DeployNativeTokenPool(client.Owner, client.Client, bridgeToken)
			require.NoError(t, err)
			shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
			client.Logger.Infof("Native token pool deployed on %s in tx %s", tokenPoolAddress, helpers.ExplorerLink(client.ChainId.Int64(), tx.Hash()))
			pool, err := native_token_pool.NewNativeTokenPool(tokenPoolAddress, client.Client)
			require.NoError(t, err)
			fillPoolWithTokens(t, client, pool)
			pools = append(pools, pool)
			poolAddresses = append(poolAddresses, tokenPoolAddress)
		} else {
			if client.TokenPools[i].Hex() == "0x0000000000000000000000000000000000000000" {
				t.Error("deploy new lock unlock pool set to false but no lock unlock pool given in config")
			}
			pool, err := native_token_pool.NewNativeTokenPool(client.TokenPools[i], client.Client)
			require.NoError(t, err)
			client.Logger.Infof("Lock unlock pool loaded from: %s", pool.Address().Hex())
			pools = append(pools, pool)
			poolAddresses = append(poolAddresses, client.TokenPools[i])
		}
	}

	client.TokenPools = poolAddresses
	return pools
}

func DeployPingPongDapps(t *testing.T, sourceClient *EvmChainConfig, destClient *EvmChainConfig) {
	if sourceClient.DeploySettings.DeployPingPongDapp {
		sourceClient.Logger.Infof("Deploying source chain ping pong dapp")

		pingPongDappAddress, tx, _, err := ping_pong_demo.DeployPingPongDemo(sourceClient.Owner, sourceClient.Client, sourceClient.OffRampRouter, sourceClient.OnRampRouter)
		require.NoError(t, err)

		shared.WaitForMined(t, sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
		sourceClient.Logger.Infof("Ping pong deployed on %s in tx: %s", pingPongDappAddress.Hex(), helpers.ExplorerLink(sourceClient.ChainId.Int64(), tx.Hash()))

		sourceClient.PingPongDapp = pingPongDappAddress
	}

	if destClient.DeploySettings.DeployPingPongDapp {
		destClient.Logger.Infof("Deploying destination chain ping pong dapp")

		pingPongDappAddress, tx, _, err := ping_pong_demo.DeployPingPongDemo(destClient.Owner, destClient.Client, destClient.OffRampRouter, destClient.OnRampRouter)
		require.NoError(t, err)

		shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
		destClient.Logger.Infof("Ping pong deployed on %s in tx: %s", pingPongDappAddress.Hex(), helpers.ExplorerLink(destClient.ChainId.Int64(), tx.Hash()))

		destClient.PingPongDapp = pingPongDappAddress
	}

	if sourceClient.DeploySettings.DeployPingPongDapp || destClient.DeploySettings.DeployPingPongDapp {
		pingDapp, err := ping_pong_demo.NewPingPongDemo(sourceClient.PingPongDapp, sourceClient.Client)
		require.NoError(t, err)

		tx, err := pingDapp.SetCounterpart(sourceClient.Owner, destClient.ChainId, destClient.PingPongDapp)
		require.NoError(t, err)
		shared.WaitForMined(t, sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
		sourceClient.Logger.Infof("Ping pong dapp configured in tx: %s", helpers.ExplorerLink(sourceClient.ChainId.Int64(), tx.Hash()))

		pongDapp, err := ping_pong_demo.NewPingPongDemo(destClient.PingPongDapp, destClient.Client)
		require.NoError(t, err)

		tx, err = pongDapp.SetCounterpart(destClient.Owner, sourceClient.ChainId, sourceClient.PingPongDapp)
		require.NoError(t, err)
		shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
		sourceClient.Logger.Infof("Ping pong dapp configured in tx: %s", helpers.ExplorerLink(sourceClient.ChainId.Int64(), tx.Hash()))

		createDestSubscription(t, sourceClient, sourceClient.PingPongDapp, []common.Address{destClient.PingPongDapp})
		createDestSubscription(t, destClient, destClient.PingPongDapp, []common.Address{sourceClient.PingPongDapp})
	} else {
		sourceClient.Logger.Infof("Skipping ping pong deployment")
	}
}

func deployGovernanceDapps(t *testing.T, sourceClient *EvmChainConfig, destClient *EvmChainConfig) {
	feeConfig := governance_dapp.GovernanceDappFeeConfig{
		FeeAmount:           big.NewInt(10),
		SubscriptionManager: sourceClient.Owner.From,
		ChangedAtBlock:      big.NewInt(0),
	}

	if sourceClient.DeploySettings.DeployGovernanceDapp {
		sourceClient.Logger.Infof("Deploying source chain governance dapp")
		governanceDappAddress, tx, _, err := governance_dapp.DeployGovernanceDapp(sourceClient.Owner, sourceClient.Client, sourceClient.OffRampRouter, sourceClient.OnRampRouter, feeConfig)
		require.NoError(t, err)

		shared.WaitForMined(t, sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
		sourceClient.Logger.Infof("GovernanceDapp deployed on %s in tx: %s", governanceDappAddress.Hex(), helpers.ExplorerLink(sourceClient.ChainId.Int64(), tx.Hash()))

		sourceClient.GovernanceDapp = governanceDappAddress
	}

	if destClient.DeploySettings.DeployGovernanceDapp {
		destClient.Logger.Infof("Deploying destination chain governance dapp")
		governanceDappAddress, tx, _, err := governance_dapp.DeployGovernanceDapp(destClient.Owner, destClient.Client, destClient.OffRampRouter, destClient.OnRampRouter, feeConfig)
		require.NoError(t, err)

		shared.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
		destClient.Logger.Infof("GovernanceDapp deployed on %s in tx: %s", governanceDappAddress.Hex(), helpers.ExplorerLink(destClient.ChainId.Int64(), tx.Hash()))

		destClient.GovernanceDapp = governanceDappAddress

		createDestSubscription(t, destClient, destClient.GovernanceDapp, []common.Address{sourceClient.GovernanceDapp})
	}

	if sourceClient.DeploySettings.DeployGovernanceDapp || destClient.DeploySettings.DeployGovernanceDapp {
		governanceDapp, err := governance_dapp.NewGovernanceDapp(sourceClient.GovernanceDapp, sourceClient.Client)
		require.NoError(t, err)

		governanceClone := governance_dapp.GovernanceDappCrossChainClone{
			ChainId:         destClient.ChainId,
			ContractAddress: destClient.GovernanceDapp,
		}

		tx, err := governanceDapp.AddClone(sourceClient.Owner, governanceClone)
		require.NoError(t, err)
		shared.WaitForMined(t, sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
		sourceClient.Logger.Infof("GovernanceDapp configured in tx: %s", helpers.ExplorerLink(sourceClient.ChainId.Int64(), tx.Hash()))
	}
}

func deployAFN(t *testing.T, client *EvmChainConfig) *afn_contract.AFNContract {
	if client.DeploySettings.DeployAFN {
		client.Logger.Infof("Deploying AFN")
		address, tx, _, err := mock_afn_contract.DeployMockAFNContract(client.Owner, client.Client)
		require.NoError(t, err)
		shared.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
		client.Logger.Infof("AFN deployed on %s in tx: %s", address.Hex(), helpers.ExplorerLink(client.ChainId.Int64(), tx.Hash()))
		client.Afn = address

		afn, err := afn_contract.NewAFNContract(address, client.Client)
		require.NoError(t, err)
		return afn
	}
	if client.Afn.Hex() == "0x0000000000000000000000000000000000000000" {
		t.Error("deploy new afn set to false but no afn given in config")
	}
	afn, err := afn_contract.NewAFNContract(client.Afn, client.Client)
	require.NoError(t, err)
	client.Logger.Infof("AFN loaded from: %s", afn.Address().Hex())
	return afn
}
