package main

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/blob_verifier"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_toll_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/receiver_dapp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/simple_message_receiver"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/toll_sender_dapp"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
)

// deployCCIPContracts will deploy all source and Destination chain contracts using the
// owner key. Only run this of the currently deployed contracts are outdated or
// when initializing a new chain.
func deployCCIPContracts(t *testing.T, ownerKey string, sourceChain *EvmChainConfig, destChain *EvmChainConfig) {
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
	tokenSenderAddress, tx, _, err := toll_sender_dapp.DeployTollSenderDapp(source.Owner, source.Client, source.OnRamp, destination.ChainId, destination.ReceiverDapp)
	require.NoError(t, err)
	WaitForMined(t, destination.Logger, source.Client, tx.Hash(), true)
	source.Logger.Infof("Token sender dapp deployed on %s in tx: %s", tokenSenderAddress.Hex(), helpers.ExplorerLink(source.ChainId.Int64(), tx.Hash()))
	source.TokenSender = tokenSenderAddress

	printContractConfig(source, destination)
}

func deploySourceContracts(t *testing.T, source *EvmChainConfig, offRampChainID *big.Int) {
	// Updates source.TokenPools if any new contracts are deployed
	tokenPools := deployNativeTokenPool(t, source)
	// Updates source.AFN if any new contracts are deployed
	deployAFN(t, source)
	// Updates source.OnRampRouter if any new contracts are deployed
	deployOnRampRouter(t, source)
	// Updates source.OnRamp if any new contracts are deployed
	deployOnRamp(t, source, offRampChainID)

	// Skip if we reuse both the onRamp and the token pools
	if source.DeploySettings.DeployRamp || source.DeploySettings.DeployTokenPools {
		for _, tokenPool := range tokenPools {
			// Configure onramp address on pool
			tx, err := tokenPool.SetOnRamp(source.Owner, source.OnRamp, true)
			require.NoError(t, err)
			source.Logger.Infof("Onramp pool configured with onramp: %s", helpers.ExplorerLink(source.ChainId.Int64(), tx.Hash()))
		}
	}
}

func deployDestinationContracts(t *testing.T, destClient *EvmChainConfig, sourceClient *EvmChainConfig) {
	// Updates source.TokenPools if any new contracts are deployed
	tokenPools := deployNativeTokenPool(t, destClient)
	// Updates source.AFN if any new contracts are deployed
	deployAFN(t, destClient)
	// Updates source.BlobVerifier if any new contracts are deployed
	deployBlobVerifier(t, destClient, sourceClient)

	// Deploy offramp contract message receiver
	messageReceiverAddress, tx, _, err := simple_message_receiver.DeploySimpleMessageReceiver(destClient.Owner, destClient.Client)
	require.NoError(t, err)
	WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
	destClient.Logger.Infof("Offramp message receiver deployed on %s in tx: %s", messageReceiverAddress.Hex(), helpers.ExplorerLink(destClient.ChainId.Int64(), tx.Hash()))
	destClient.MessageReceiver = messageReceiverAddress

	// Updates source.ReceiverDapp if any new contracts are deployed
	deployReceiverDapp(t, destClient)
	// Updates source.OffRamp if any new contracts are deployed
	deployOffRamp(t, destClient, sourceClient)
	// Updates source.OffRampRouter if any new contracts are deployed
	deployOffRampRouter(t, destClient)

	if destClient.DeploySettings.DeployRamp || destClient.DeploySettings.DeployTokenPools {
		for _, tokenPool := range tokenPools {
			// Configure offramp address on pool
			tx, err = tokenPool.SetOffRamp(destClient.Owner, destClient.BlobVerifier, true)
			require.NoError(t, err)
			WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
			destClient.Logger.Infof("Offramp pool configured with offramp address: %s", helpers.ExplorerLink(destClient.ChainId.Int64(), tx.Hash()))
		}
	}
}

func deployOnRampRouter(t *testing.T, client *EvmChainConfig) *evm_2_any_toll_onramp_router.EVM2AnyTollOnRampRouter {
	if !client.DeploySettings.DeployRouter {
		client.Logger.Infof("Skipping OnRampRouter deployment, using OnRampRouter on %s", client.OnRampRouter)
		onRampRouter, err := evm_2_any_toll_onramp_router.NewEVM2AnyTollOnRampRouter(client.OnRamp, client.Client)
		require.NoError(t, err)
		return onRampRouter
	}
	client.Logger.Infof("Deploying OnRampRouter")

	onRampRouterAddress, tx, _, err := evm_2_any_toll_onramp_router.DeployEVM2AnyTollOnRampRouter(client.Owner, client.Client)
	require.NoError(t, err)
	WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	client.OnRampRouter = onRampRouterAddress

	client.Logger.Infof(fmt.Sprintf("OnRampRouter deployed on %s in tx %s", onRampRouterAddress.String(), helpers.ExplorerLink(client.ChainId.Int64(), tx.Hash())))

	onRampRouter, err := evm_2_any_toll_onramp_router.NewEVM2AnyTollOnRampRouter(client.OnRamp, client.Client)
	require.NoError(t, err)
	return onRampRouter
}

func deployOnRamp(t *testing.T, client *EvmChainConfig, destinationChain *big.Int) *evm_2_evm_toll_onramp.EVM2EVMTollOnRamp {
	if !client.DeploySettings.DeployRamp {
		client.Logger.Infof("Skipping OnRamp deployment, using onRamp on %s", client.OnRamp)
		onRamp, err := evm_2_evm_toll_onramp.NewEVM2EVMTollOnRamp(client.OnRamp, client.Client)
		require.NoError(t, err)
		return onRamp
	}

	client.Logger.Infof("Deploying OnRamp: destinationChains %+v, bridgeTokens %+v, poolAddresses %+v", destinationChain, client.BridgeTokens, client.TokenPools)
	onRampAddress, tx, _, err := evm_2_evm_toll_onramp.DeployEVM2EVMTollOnRamp(
		client.Owner,        // user
		client.Client,       // client
		client.ChainId,      // source chain id
		destinationChain,    // destinationChainId
		client.BridgeTokens, // tokens
		client.TokenPools,   // pools
		[]common.Address{},  // allow list
		client.Afn,          // AFN
		evm_2_evm_toll_onramp.BaseOnRampInterfaceOnRampConfig{
			RelayingFeeJuels: 0,
			MaxDataSize:      1e6,
			MaxTokensLength:  5,
		},
		evm_2_evm_toll_onramp.AggregateRateLimiterInterfaceRateLimiterConfig{
			Capacity: big.NewInt(1e18),
			Rate:     big.NewInt(1e18),
		},
		client.Owner.From,
		client.OnRampRouter,
	)
	require.NoError(t, err)
	WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)

	onRamp, err := evm_2_evm_toll_onramp.NewEVM2EVMTollOnRamp(onRampAddress, client.Client)
	require.NoError(t, err)
	client.Logger.Infof(fmt.Sprintf("Onramp deployed on %s in tx %s", onRampAddress.String(), helpers.ExplorerLink(client.ChainId.Int64(), tx.Hash())))
	client.OnRamp = onRampAddress

	onRampRouter, err := evm_2_any_toll_onramp_router.NewEVM2AnyTollOnRampRouterTransactor(client.OnRampRouter, client.Client)
	require.NoError(t, err)
	tx, err = onRampRouter.SetOnRamp(client.Owner, destinationChain, onRampAddress)
	require.NoError(t, err)
	WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)

	_, err = onRamp.SetPrices(client.Owner, []common.Address{client.LinkToken}, []*big.Int{big.NewInt(1)})
	require.NoError(t, err)

	return onRamp
}

func deployOffRamp(t *testing.T, destClient *EvmChainConfig, sourceClient *EvmChainConfig) *any_2_evm_toll_offramp.EVM2EVMTollOffRamp {
	if !destClient.DeploySettings.DeployRamp {
		destClient.Logger.Infof("Skipping OffRamp deployment, using offRamp on %s", destClient.OnRamp)
		offRamp, err := any_2_evm_toll_offramp.NewEVM2EVMTollOffRamp(destClient.OffRamp, destClient.Client)
		require.NoError(t, err)
		return offRamp
	}

	destClient.Logger.Infof("Deploying OffRamp")
	tollOffRampAddress, tx, _, err := any_2_evm_toll_offramp.DeployEVM2EVMTollOffRamp(

		destClient.Owner,
		destClient.Client,
		sourceClient.ChainId,
		destClient.ChainId,
		any_2_evm_toll_offramp.BaseOffRampInterfaceOffRampConfig{
			ExecutionDelaySeconds:                   60,
			MaxDataSize:                             1e5,
			MaxTokensLength:                         15,
			PermissionLessExecutionThresholdSeconds: 60,
		},
		destClient.BlobVerifier,
		destClient.OnRamp,
		destClient.Afn,
		sourceClient.BridgeTokens,
		destClient.TokenPools,
		any_2_evm_toll_offramp.AggregateRateLimiterInterfaceRateLimiterConfig{
			Capacity: big.NewInt(1e18),
			Rate:     big.NewInt(1e18),
		},
		destClient.Owner.From)
	require.NoError(t, err)
	WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)

	destClient.Logger.Infof("OffRamp contract deployed on %s in tx: %s", tollOffRampAddress.Hex(), helpers.ExplorerLink(destClient.ChainId.Int64(), tx.Hash()))
	destClient.OffRamp = tollOffRampAddress
	offRamp, err := any_2_evm_toll_offramp.NewEVM2EVMTollOffRamp(destClient.OffRamp, destClient.Client)
	require.NoError(t, err)

	_, err = offRamp.SetPrices(destClient.Owner, []common.Address{sourceClient.LinkToken}, []*big.Int{big.NewInt(1)})
	require.NoError(t, err)

	return offRamp
}

func deployOffRampRouter(t *testing.T, destClient *EvmChainConfig) *any_2_evm_toll_offramp_router.Any2EVMTollOffRampRouter {
	if !destClient.DeploySettings.DeployRouter {
		destClient.Logger.Infof("Skipping OffRampRouter deployment, using OffRampRouter on %s", destClient.OffRampRouter)
		offRampRouter, err := any_2_evm_toll_offramp_router.NewAny2EVMTollOffRampRouter(destClient.OffRampRouter, destClient.Client)
		require.NoError(t, err)
		return offRampRouter
	}

	destClient.Logger.Infof("Deploying OffRampRouter")
	offRampRouterAddress, tx, _, err := any_2_evm_toll_offramp_router.DeployAny2EVMTollOffRampRouter(destClient.Owner, destClient.Client, []common.Address{destClient.OffRamp})
	require.NoError(t, err)
	WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
	destClient.OffRampRouter = offRampRouterAddress

	destClient.Logger.Infof(fmt.Sprintf("OffRampRouter deployed on %s in tx %s", offRampRouterAddress.String(), helpers.ExplorerLink(destClient.ChainId.Int64(), tx.Hash())))

	offRampRouter, err := any_2_evm_toll_offramp_router.NewAny2EVMTollOffRampRouter(destClient.OffRampRouter, destClient.Client)
	require.NoError(t, err)

	offRamp, err := any_2_evm_toll_offramp.NewEVM2EVMTollOffRamp(destClient.OffRamp, destClient.Client)
	require.NoError(t, err)

	tx, err = offRamp.SetRouter(destClient.Owner, offRampRouterAddress)
	require.NoError(t, err)
	WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
	destClient.Logger.Infof(fmt.Sprintf("OffRampRouter set on offRamp in tx %s", helpers.ExplorerLink(destClient.ChainId.Int64(), tx.Hash())))

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
	WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
	destClient.Logger.Infof("Blob verifier deployed on %s in tx: %s", blobVerifierAddress.Hex(), helpers.ExplorerLink(destClient.ChainId.Int64(), tx.Hash()))
	destClient.BlobVerifier = blobVerifierAddress

	blobVerifier, err := blob_verifier.NewBlobVerifier(blobVerifierAddress, destClient.Client)
	require.NoError(t, err)
	return blobVerifier
}

func deployReceiverDapp(t *testing.T, destClient *EvmChainConfig) *receiver_dapp.ReceiverDapp {
	destClient.Logger.Infof("Deploying ReceiverDapp")
	receiverDappAddress, tx, _, err := receiver_dapp.DeployReceiverDapp(destClient.Owner, destClient.Client, destClient.OnRampRouter)
	require.NoError(t, err)
	WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
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
			WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
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

func deployAFN(t *testing.T, client *EvmChainConfig) *afn_contract.AFNContract {
	if client.DeploySettings.DeployAFN {
		client.Logger.Infof("Deploying AFN")
		address, tx, _, err := afn_contract.DeployAFNContract(
			client.Owner,
			client.Client,
			[]common.Address{client.Owner.From},
			[]*big.Int{big.NewInt(1)},
			big.NewInt(1),
			big.NewInt(1),
		)
		require.NoError(t, err)
		WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
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

func fillPoolWithTokens(t *testing.T, client *EvmChainConfig, pool *native_token_pool.NativeTokenPool) {
	destLinkToken, err := link_token_interface.NewLinkToken(client.LinkToken, client.Client)
	require.NoError(t, err)

	// fill offramp token pool with 0.5 LINK
	amount := big.NewInt(5e17)
	tx, err := destLinkToken.Transfer(client.Owner, pool.Address(), amount)
	require.NoError(t, err)
	client.Logger.Infof("Transferring token to token pool: %s", helpers.ExplorerLink(client.ChainId.Int64(), tx.Hash()))
	WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)

	client.Logger.Infof("Locking tokens in pool")
	tx, err = pool.LockOrBurn(client.Owner, amount)
	require.NoError(t, err)
	WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	client.Logger.Infof("Pool filled with tokens: %s", helpers.ExplorerLink(client.ChainId.Int64(), tx.Hash()))
}

func printContractConfig(source *EvmChainConfig, destination *EvmChainConfig) {
	source.Logger.Infof("Source chain config")
	source.Logger.Infof(`
Source chain config

LinkToken:    common.HexToAddress("%s"),
BridgeTokens: %s,
TokenPools:   %s,
PriceFeeds:   %s,
OnRamp:       common.HexToAddress("%s"),
OnRampRouter: common.HexToAddress("%s"),
TokenSender:  common.HexToAddress("%s"),
Afn:          common.HexToAddress("%s"),
`,
		source.LinkToken,
		source.BridgeTokens,
		source.TokenPools,
		source.OnRamp,
		source.OnRampRouter,
		source.TokenSender,
		source.Afn)

	destination.Logger.Infof(`
Destination chain config

LinkToken:       common.HexToAddress("%s"),
BridgeTokens:    %s,
TokenPools:      %s,
OffRamp:         common.HexToAddress("%s"),
OffRampRouter:   common.HexToAddress("%s"),
BlobVerifier:    common.HexToAddress("%s"),	
MessageReceiver: common.HexToAddress("%s"),
ReceiverDapp:    common.HexToAddress("%s"),
Afn:             common.HexToAddress("%s"),
`,
		destination.LinkToken,
		destination.BridgeTokens,
		destination.TokenPools,
		destination.OffRamp,
		destination.OffRampRouter,
		destination.BlobVerifier,
		destination.MessageReceiver,
		destination.ReceiverDapp,
		destination.Afn)

	PrintJobSpecs(source.OnRamp, destination.BlobVerifier, destination.OffRamp, source.ChainId, destination.ChainId)
}
