package toll

//
//import (
//	"fmt"
//	"math/big"
//	"testing"
//
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/stretchr/testify/require"
//
//	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/afn_contract"
//	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp"
//	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp_router"
//	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
//	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_toll_onramp_router"
//	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
//	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/link_token_interface"
//	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/native_token_pool"
//	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/receiver_dapp"
//	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/simple_message_receiver"
//	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/toll_sender_dapp"
//	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test"
//	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
//)
//
//// deployCCIPContracts will deploy all source and Destination chain contracts using the
//// owner key. Only run this of the currently deployed contracts are outdated or
//// when initializing a new chain.
//func deployCCIPContracts(t *testing.T, ownerKey string, sourceChain *main.EvmChainConfig, destChain *main.EvmChainConfig) {
//	sourceChain.SetupChain(t, ownerKey)
//	destChain.SetupChain(t, ownerKey)
//	deploySourceAndDestContracts(t, sourceChain, destChain)
//}
//
//func deploySourceAndDestContracts(t *testing.T, source *main.EvmChainConfig, destination *main.EvmChainConfig) {
//	// After running this code please update the configuration to reflect the newly
//	// deployed contract addresses.
//	deploySourceContracts(t, source, destination.ChainConfig.ChainId)
//	source.Logger.Infof("%s contracts fully deployed as source chain", helpers.ChainName(source.ChainConfig.ChainId.Int64()))
//
//	deployDestinationContracts(t, destination, source)
//	destination.Logger.Infof("%s contracts fully deployed as destination chain", helpers.ChainName(destination.ChainConfig.ChainId.Int64()))
//
//	// Deploy onramp sender dapp
//	tokenSenderAddress, tx, _, err := toll_sender_dapp.DeployTollSenderDapp(source.Owner, source.Client, source.OnRamp, destination.ChainConfig.ChainId, destination.ReceiverDapp)
//	require.NoError(t, err)
//	main.WaitForMined(t, destination.Logger, source.Client, tx.Hash(), true)
//	source.Logger.Infof("Token sender dapp deployed on %s in tx: %s", tokenSenderAddress.Hex(), helpers.ExplorerLink(source.ChainConfig.ChainId.Int64(), tx.Hash()))
//	source.LaneConfig.TokenSender = tokenSenderAddress
//
//	printContractConfig(source, destination)
//}
//
//func deploySourceContracts(t *testing.T, source *main.EvmChainConfig, offRampChainID *big.Int) {
//	// Updates source.TokenPools if any new contracts are deployed
//	tokenPools := deployNativeTokenPool(t, source)
//	// Updates source.AFN if any new contracts are deployed
//	deployAFN(t, source)
//	// Updates source.ChainConfig.OnRampRouter if any new contracts are deployed
//	deployOnRampRouter(t, source)
//	// Updates source.OnRamp if any new contracts are deployed
//	deployOnRamp(t, source, offRampChainID)
//
//	// Skip if we reuse both the onRamp and the token pools
//	if source.DeploySettings.DeployRamp || source.DeploySettings.DeployTokenPools {
//		for _, tokenPool := range tokenPools {
//			// Configure onramp address on pool
//			tx, err := tokenPool.SetOnRamp(source.Owner, source.OnRamp, true)
//			require.NoError(t, err)
//			source.Logger.Infof("Onramp pool configured with onramp: %s", helpers.ExplorerLink(source.ChainConfig.ChainId.Int64(), tx.Hash()))
//		}
//	}
//}
//
//func deployDestinationContracts(t *testing.T, destClient *main.EvmChainConfig, sourceClient *main.EvmChainConfig) {
//	// Updates source.TokenPools if any new contracts are deployed
//	tokenPools := deployNativeTokenPool(t, destClient)
//	// Updates source.AFN if any new contracts are deployed
//	deployAFN(t, destClient)
//	// Updates source.CommitStore if any new contracts are deployed
//	deployCommitStore(t, destClient, sourceClient)
//
//	// Deploy offramp contract message receiver
//	messageReceiverAddress, tx, _, err := simple_message_receiver.DeploySimpleMessageReceiver(destClient.Owner, destClient.Client)
//	require.NoError(t, err)
//	main.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
//	destClient.Logger.Infof("Offramp message receiver deployed on %s in tx: %s", messageReceiverAddress.Hex(), helpers.ExplorerLink(destClient.ChainConfig.ChainId.Int64(), tx.Hash()))
//	destClient.MessageReceiver = messageReceiverAddress
//
//	// Updates source.ReceiverDapp if any new contracts are deployed
//	deployReceiverDapp(t, destClient)
//	// Updates source.OffRamp if any new contracts are deployed
//	deployOffRamp(t, destClient, sourceClient)
//	// Updates source.OffRampRouter if any new contracts are deployed
//	deployOffRampRouter(t, destClient)
//
//	if destClient.DeploySettings.DeployRamp || destClient.DeploySettings.DeployTokenPools {
//		for _, tokenPool := range tokenPools {
//			// Configure offramp address on pool
//			tx, err = tokenPool.SetOffRamp(destClient.Owner, destClient.LaneConfig.CommitStore, true)
//			require.NoError(t, err)
//			main.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
//			destClient.Logger.Infof("Offramp pool configured with offramp address: %s", helpers.ExplorerLink(destClient.ChainConfig.ChainId.Int64(), tx.Hash()))
//		}
//	}
//}
//
//func deployOnRampRouter(t *testing.T, client *main.EvmChainConfig) *evm_2_any_toll_onramp_router.EVM2AnyTollOnRampRouter {
//	if !client.DeploySettings.DeployRouter {
//		client.Logger.Infof("Skipping OnRampRouter deployment, using OnRampRouter on %s", Client.LaneConfig.OnRampRouter)
//		onRampRouter, err := evm_2_any_toll_onramp_router.NewEVM2AnyTollOnRampRouter(Client.LaneConfig.OnRamp, client.Client)
//		require.NoError(t, err)
//		return onRampRouter
//	}
//	client.Logger.Infof("Deploying OnRampRouter")
//
//	onRampRouterAddress, tx, _, err := evm_2_any_toll_onramp_router.DeployEVM2AnyTollOnRampRouter(client.Owner, client.Client)
//	require.NoError(t, err)
//	main.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
//	Client.LaneConfig.OnRampRouter = onRampRouterAddress
//
//	client.Logger.Infof(fmt.Sprintf("OnRampRouter deployed on %s in tx %s", onRampRouterAddress.String(), helpers.ExplorerLink(Client.ChainConfig.ChainId.Int64(), tx.Hash())))
//
//	onRampRouter, err := evm_2_any_toll_onramp_router.NewEVM2AnyTollOnRampRouter(Client.LaneConfig.OnRamp, client.Client)
//	require.NoError(t, err)
//	return onRampRouter
//}
//
//func deployOnRamp(t *testing.T, client *main.EvmChainConfig, destinationChain *big.Int) *evm_2_evm_toll_onramp.EVM2EVMTollOnRamp {
//	if !client.DeploySettings.DeployRamp {
//		client.Logger.Infof("Skipping OnRamp deployment, using onRamp on %s", Client.LaneConfig.OnRamp)
//		onRamp, err := evm_2_evm_toll_onramp.NewEVM2EVMTollOnRamp(Client.LaneConfig.OnRamp, client.Client)
//		require.NoError(t, err)
//		return onRamp
//	}
//
//	client.Logger.Infof("Deploying OnRamp: destinationChains %+v, bridgeTokens %+v, poolAddresses %+v", destinationChain, client.ChainConfig.BridgeTokens, client.ChainConfig.TokenPools)
//	onRampAddress, tx, _, err := evm_2_evm_toll_onramp.DeployEVM2EVMTollOnRamp(
//		client.Owner,        // user
//		client.Client,       // client
//		Client.ChainConfig.ChainId,      // source chain id
//		destinationChain,    // destinationChainId
//		client.ChainConfig.BridgeTokens, // tokens
//		client.ChainConfig.TokenPools,   // pools
//		[]common.Address{},  // allow list
//		client.ChainConfig.Afn,          // AFN
//		evm_2_evm_toll_onramp.BaseOnRampInterfaceOnRampConfig{
//			CommitFeeJuels: 0,
//			MaxDataSize:      1e6,
//			MaxTokensLength:  5,
//			MaxGasLimit:      ccip.BatchGasLimit,
//		},
//		evm_2_evm_toll_onramp.AggregateRateLimiterInterfaceRateLimiterConfig{
//			Capacity: big.NewInt(1e18),
//			Rate:     big.NewInt(1e18),
//		},
//		client.Owner.From,
//		Client.LaneConfig.OnRampRouter,
//	)
//	require.NoError(t, err)
//	main.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
//
//	onRamp, err := evm_2_evm_toll_onramp.NewEVM2EVMTollOnRamp(onRampAddress, client.Client)
//	require.NoError(t, err)
//	client.Logger.Infof(fmt.Sprintf("Onramp deployed on %s in tx %s", onRampAddress.String(), helpers.ExplorerLink(Client.ChainConfig.ChainId.Int64(), tx.Hash())))
//	Client.LaneConfig.OnRamp = onRampAddress
//
//	onRampRouter, err := evm_2_any_toll_onramp_router.NewEVM2AnyTollOnRampRouterTransactor(Client.LaneConfig.OnRampRouter, client.Client)
//	require.NoError(t, err)
//	tx, err = onRampRouter.SetOnRamp(client.Owner, destinationChain, onRampAddress)
//	require.NoError(t, err)
//	main.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
//
//	_, err = onRamp.SetPrices(client.Owner, []common.Address{client.ChainConfig.LinkToken}, []*big.Int{big.NewInt(1)})
//	require.NoError(t, err)
//
//	return onRamp
//}
//
//func deployOffRamp(t *testing.T, destClient *main.EvmChainConfig, sourceClient *main.EvmChainConfig) *any_2_evm_toll_offramp.EVM2EVMTollOffRamp {
//	if !destClient.DeploySettings.DeployRamp {
//		destClient.Logger.Infof("Skipping OffRamp deployment, using offRamp on %s", destClient.LaneConfig.OnRamp)
//		offRamp, err := any_2_evm_toll_offramp.NewEVM2EVMTollOffRamp(destClient.LaneConfig.OffRamp, destClient.Client)
//		require.NoError(t, err)
//		return offRamp
//	}
//
//	destClient.Logger.Infof("Deploying OffRamp")
//	tollOffRampAddress, tx, _, err := any_2_evm_toll_offramp.DeployEVM2EVMTollOffRamp(
//		destClient.Owner,
//		destClient.Client,
//		sourceClient.ChainConfig.ChainId,
//		destClient.ChainConfig.ChainId,
//		any_2_evm_toll_offramp.BaseOffRampInterfaceOffRampConfig{
//			OnRampAddress:                           sourceClient.LaneConfig.OnRamp,
//			ExecutionDelaySeconds:                   60,
//			MaxDataSize:                             1e5,
//			MaxTokensLength:                         15,
//			PermissionLessExecutionThresholdSeconds: 60,
//		},
//		destClient.LaneConfig.CommitStore,
//		destClient.ChainConfig.Afn,
//		sourceclient.ChainConfig.BridgeTokens,
//		destclient.ChainConfig.TokenPools,
//		any_2_evm_toll_offramp.AggregateRateLimiterInterfaceRateLimiterConfig{
//			Capacity: big.NewInt(1e18),
//			Rate:     big.NewInt(1e18),
//		},
//		destClient.Owner.From)
//	require.NoError(t, err)
//	main.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
//
//	destClient.Logger.Infof("OffRamp contract deployed on %s in tx: %s", tollOffRampAddress.Hex(), helpers.ExplorerLink(destClient.ChainConfig.ChainId.Int64(), tx.Hash()))
//	destClient.LaneConfig.OffRamp = tollOffRampAddress
//	offRamp, err := any_2_evm_toll_offramp.NewEVM2EVMTollOffRamp(destClient.LaneConfig.OffRamp, destClient.Client)
//	require.NoError(t, err)
//
//	_, err = offRamp.SetPrices(destClient.Owner, []common.Address{destclient.ChainConfig.LinkToken}, []*big.Int{big.NewInt(1)})
//	require.NoError(t, err)
//
//	return offRamp
//}
//
//func deployOffRampRouter(t *testing.T, destClient *main.EvmChainConfig) *any_2_evm_toll_offramp_router.Any2EVMTollOffRampRouter {
//	if !destClient.DeploySettings.DeployRouter {
//		destClient.Logger.Infof("Skipping OffRampRouter deployment, using OffRampRouter on %s", destClient.ChainConfig.OffRampRouter)
//		offRampRouter, err := any_2_evm_toll_offramp_router.NewAny2EVMTollOffRampRouter(destClient.ChainConfig.OffRampRouter, destClient.Client)
//		require.NoError(t, err)
//		return offRampRouter
//	}
//
//	destClient.Logger.Infof("Deploying OffRampRouter")
//	offRampRouterAddress, tx, _, err := any_2_evm_toll_offramp_router.DeployAny2EVMTollOffRampRouter(destClient.Owner, destClient.Client, []common.Address{destClient.LaneConfig.OffRamp})
//	require.NoError(t, err)
//	main.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
//	destClient.ChainConfig.OffRampRouter = offRampRouterAddress
//
//	destClient.Logger.Infof(fmt.Sprintf("OffRampRouter deployed on %s in tx %s", offRampRouterAddress.String(), helpers.ExplorerLink(destClient.ChainConfig.ChainId.Int64(), tx.Hash())))
//
//	offRampRouter, err := any_2_evm_toll_offramp_router.NewAny2EVMTollOffRampRouter(destClient.ChainConfig.OffRampRouter, destClient.Client)
//	require.NoError(t, err)
//
//	offRamp, err := any_2_evm_toll_offramp.NewEVM2EVMTollOffRamp(destClient.LaneConfig.OffRamp, destClient.Client)
//	require.NoError(t, err)
//
//	tx, err = offRamp.SetRouter(destClient.Owner, offRampRouterAddress)
//	require.NoError(t, err)
//	main.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
//	destClient.Logger.Infof(fmt.Sprintf("OffRampRouter set on offRamp in tx %s", helpers.ExplorerLink(destClient.ChainConfig.ChainId.Int64(), tx.Hash())))
//
//	return offRampRouter
//}
//
//func deployCommitStore(t *testing.T, destClient *main.EvmChainConfig, sourceClient *main.EvmChainConfig) *commit_store.CommitStore {
//	if !destClient.DeploySettings.DeployCommitStore {
//		destClient.Logger.Infof("Skipping CommitStore deployment, using CommitStore on %s", destClient.LaneConfig.CommitStore)
//		commitStore, err := commit_store.NewCommitStore(destClient.LaneConfig.CommitStore, destClient.Client)
//		require.NoError(t, err)
//		return commitStore
//	}
//
//	destClient.Logger.Infof("Deploying blob verifier")
//
//	commitStoreAddress, tx, _, err := commit_store.DeployCommitStore(
//		destClient.Owner,     // user
//		destClient.Client,    // client
//		destClient.ChainConfig.ChainId,   // dest chain id
//		sourceClient.ChainConfig.ChainId, // source chain id
//		destClient.ChainConfig.Afn,       // AFN address
//		commit_store.CommitStoreInterfaceCommitStoreConfig{
//			OnRamps:          []common.Address{sourceClient.LaneConfig.OnRamp},
//			MinSeqNrByOnRamp: []uint64{1},
//		},
//	)
//	require.NoError(t, err)
//	main.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
//	destClient.Logger.Infof("Blob verifier deployed on %s in tx: %s", commitStoreAddress.Hex(), helpers.ExplorerLink(destClient.ChainConfig.ChainId.Int64(), tx.Hash()))
//	destClient.LaneConfig.CommitStore = commitStoreAddress
//
//	commitStore, err := commit_store.NewCommitStore(commitStoreAddress, destClient.Client)
//	require.NoError(t, err)
//	return commitStore
//}
//
//func deployReceiverDapp(t *testing.T, destClient *main.EvmChainConfig) *receiver_dapp.ReceiverDapp {
//	destClient.Logger.Infof("Deploying ReceiverDapp")
//	receiverDappAddress, tx, _, err := receiver_dapp.DeployReceiverDapp(destClient.Owner, destClient.Client, destClient.LaneConfig.OnRampRouter)
//	require.NoError(t, err)
//	main.WaitForMined(t, destClient.Logger, destClient.Client, tx.Hash(), true)
//	destClient.Logger.Infof("Offramp receiver dapp deployed on %s in tx: %s", receiverDappAddress.Hex(), helpers.ExplorerLink(destClient.ChainConfig.ChainId.Int64(), tx.Hash()))
//	destClient.ReceiverDapp = receiverDappAddress
//
//	receiverDapp, err := receiver_dapp.NewReceiverDapp(receiverDappAddress, destClient.Client)
//	require.NoError(t, err)
//	return receiverDapp
//}
//
//func deployNativeTokenPool(t *testing.T, client *main.EvmChainConfig) []*native_token_pool.NativeTokenPool {
//	var pools []*native_token_pool.NativeTokenPool
//	var poolAddresses []common.Address
//
//	for i, bridgeToken := range client.ChainConfig.BridgeTokens {
//		if client.DeploySettings.DeployTokenPools {
//			client.Logger.Infof("Deploying token pool for token %s", bridgeToken.Hex())
//			tokenPoolAddress, tx, _, err := native_token_pool.DeployNativeTokenPool(client.Owner, client.Client, bridgeToken)
//			require.NoError(t, err)
//			main.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
//			client.Logger.Infof("Native token pool deployed on %s in tx %s", tokenPoolAddress, helpers.ExplorerLink(Client.ChainConfig.ChainId.Int64(), tx.Hash()))
//			pool, err := native_token_pool.NewNativeTokenPool(tokenPoolAddress, client.Client)
//			require.NoError(t, err)
//			fillPoolWithTokens(t, client, pool)
//			pools = append(pools, pool)
//			poolAddresses = append(poolAddresses, tokenPoolAddress)
//		} else {
//			if client.ChainConfig.TokenPools[i].Hex() == "0x0000000000000000000000000000000000000000" {
//				t.Error("deploy new lock unlock pool set to false but no lock unlock pool given in config")
//			}
//			pool, err := native_token_pool.NewNativeTokenPool(client.ChainConfig.TokenPools[i], client.Client)
//			require.NoError(t, err)
//			client.Logger.Infof("Lock unlock pool loaded from: %s", pool.Address().Hex())
//			pools = append(pools, pool)
//			poolAddresses = append(poolAddresses, client.ChainConfig.TokenPools[i])
//		}
//	}
//
//	client.ChainConfig.TokenPools = poolAddresses
//	return pools
//}
//
//func deployAFN(t *testing.T, client *main.EvmChainConfig) *afn_contract.AFNContract {
//	if client.DeploySettings.DeployAFN {
//		client.Logger.Infof("Deploying AFN")
//		address, tx, _, err := afn_contract.DeployAFNContract(
//			client.Owner,
//			client.Client,
//			[]common.Address{client.Owner.From},
//			[]*big.Int{big.NewInt(1)},
//			big.NewInt(1),
//			big.NewInt(1),
//		)
//		require.NoError(t, err)
//		main.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
//		client.Logger.Infof("AFN deployed on %s in tx: %s", address.Hex(), helpers.ExplorerLink(Client.ChainConfig.ChainId.Int64(), tx.Hash()))
//		client.ChainConfig.Afn = address
//
//		afn, err := afn_contract.NewAFNContract(address, client.Client)
//		require.NoError(t, err)
//		return afn
//	}
//	if client.ChainConfig.Afn.Hex() == "0x0000000000000000000000000000000000000000" {
//		t.Error("deploy new afn set to false but no afn given in config")
//	}
//	afn, err := afn_contract.NewAFNContract(client.ChainConfig.Afn, client.Client)
//	require.NoError(t, err)
//	client.Logger.Infof("AFN loaded from: %s", afn.Address().Hex())
//	return afn
//}
//
//func fillPoolWithTokens(t *testing.T, client *main.EvmChainConfig, pool *native_token_pool.NativeTokenPool) {
//	destLinkToken, err := link_token_interface.NewLinkToken(client.ChainConfig.LinkToken, client.Client)
//	require.NoError(t, err)
//
//	// fill offramp token pool with 0.5 LINK
//	amount := big.NewInt(5e17)
//	tx, err := destLinkToken.Transfer(client.Owner, pool.Address(), amount)
//	require.NoError(t, err)
//	client.Logger.Infof("Transferring token to token pool: %s", helpers.ExplorerLink(Client.ChainConfig.ChainId.Int64(), tx.Hash()))
//	main.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
//
//	client.Logger.Infof("Locking tokens in pool")
//	tx, err = pool.LockOrBurn(client.Owner, amount)
//	require.NoError(t, err)
//	main.WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
//	client.Logger.Infof("Pool filled with tokens: %s", helpers.ExplorerLink(Client.ChainConfig.ChainId.Int64(), tx.Hash()))
//}
//
//func printContractConfig(source *main.EvmChainConfig, destination *main.EvmChainConfig) {
//	source.Logger.Infof("Source chain config")
//	source.Logger.Infof(`
//Source chain config
//
//LinkToken:    common.HexToAddress("%s"),
//BridgeTokens: %s,
//TokenPools:   %s,
//OnRamp:       common.HexToAddress("%s"),
//OnRampRouter: common.HexToAddress("%s"),
//TokenSender:  common.HexToAddress("%s"),
//Afn:          common.HexToAddress("%s"),
//`,
//		source.LinkToken,
//		source.BridgeTokens,
//		source.TokenPools,
//		source.OnRamp,
//		source.ChainConfig.OnRampRouter,
//		source.LaneConfig.TokenSender,
//		source.Afn)
//
//	destination.Logger.Infof(`
//Destination chain config
//
//LinkToken:       common.HexToAddress("%s"),
//BridgeTokens:    %s,
//TokenPools:      %s,
//OffRamp:         common.HexToAddress("%s"),
//OffRampRouter:   common.HexToAddress("%s"),
//CommitStore:    common.HexToAddress("%s"),
//MessageReceiver: common.HexToAddress("%s"),
//ReceiverDapp:    common.HexToAddress("%s"),
//Afn:             common.HexToAddress("%s"),
//`,
//		destination.LinkToken,
//		destination.BridgeTokens,
//		destination.TokenPools,
//		destination.OffRamp,
//		destination.OffRampRouter,
//		destination.CommitStore,
//		destination.MessageReceiver,
//		destination.ReceiverDapp,
//		destination.Afn)
//
//	main.PrintJobSpecs(source.OnRamp, destination.CommitStore, destination.OffRamp, source.ChainConfig.ChainId, destination.ChainConfig.ChainId)
//}
