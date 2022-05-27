package main

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/mock_v3_aggregator_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp_executor"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/onramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/receiver_dapp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/sender_dapp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/simple_message_receiver"
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
	onRamp := deploySourceContracts(t, source, destination.ChainId)
	source.Logger.Infof("%s contracts fully deployed as source chain", helpers.ChainName(source.ChainId.Int64()))

	tokenReceiver := deployDestinationContracts(t, destination, source.ChainId, source.BridgeTokens)
	destination.Logger.Infof("%s contracts fully deployed as destination chain", helpers.ChainName(destination.ChainId.Int64()))

	// Deploy onramp sender dapp
	tokenSenderAddress, tx, _, err := sender_dapp.DeploySenderDapp(source.Owner, source.Client, onRamp.Address(), destination.ChainId, tokenReceiver)
	require.NoError(t, err)
	WaitForMined(t, destination.Logger, source.Client, tx.Hash(), true)
	source.Logger.Infof("Token sender dapp deployed on %s in tx: %s", tokenSenderAddress.Hex(), helpers.ExplorerLink(source.ChainId.Int64(), tx.Hash()))
	source.TokenSender = tokenSenderAddress

	printContractConfig(source, destination)
}

func deploySourceContracts(t *testing.T, source *EvmChainConfig, offRampChainID *big.Int) *onramp.OnRamp {
	tokenPools := deployNativeTokenPool(t, source)
	afn := deployAFN(t, source)
	feedAddresses := deployPriceFeed(t, source)

	var tokenPoolAddresses []common.Address
	for _, tokenPool := range tokenPools {
		tokenPoolAddresses = append(tokenPoolAddresses, tokenPool.Address())
	}

	onRamp := deployOnRamp(t, source, offRampChainID, tokenPoolAddresses, feedAddresses, afn.Address())

	for _, tokenPool := range tokenPools {
		// Configure onramp address on pool
		tx, err := tokenPool.SetOnRamp(source.Owner, onRamp.Address(), true)
		require.NoError(t, err)
		source.Logger.Infof("Onramp pool configured with onramp: %s", helpers.ExplorerLink(source.ChainId.Int64(), tx.Hash()))
	}

	return onRamp
}

func deployDestinationContracts(t *testing.T, client *EvmChainConfig, onRampChainId *big.Int, sourceBridgeTokens []common.Address) common.Address {
	tokenPools := deployNativeTokenPool(t, client)
	afn := deployAFN(t, client)
	feedAddresses := deployPriceFeed(t, client)

	var tokenPoolAddresses []common.Address
	for _, tokenPool := range tokenPools {
		tokenPoolAddresses = append(tokenPoolAddresses, tokenPool.Address())
	}

	offRamp := deployOffRamp(t, client, onRampChainId, tokenPoolAddresses, feedAddresses, afn.Address(), sourceBridgeTokens)

	for _, tokenPool := range tokenPools {
		// Configure offramp address on pool
		tx, err := tokenPool.SetOffRamp(client.Owner, offRamp.Address(), true)
		require.NoError(t, err)
		WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
		client.Logger.Infof("Offramp pool configured with offramp address: %s", helpers.ExplorerLink(client.ChainId.Int64(), tx.Hash()))
	}

	// Deploy offramp contract token receiver
	messageReceiverAddress, tx, _, err := simple_message_receiver.DeploySimpleMessageReceiver(client.Owner, client.Client)
	require.NoError(t, err)
	WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	client.Logger.Infof("Offramp contract message receiver deployed on %s in tx: %s", messageReceiverAddress.Hex(), helpers.ExplorerLink(client.ChainId.Int64(), tx.Hash()))
	client.MessageReceiver = messageReceiverAddress

	// Deploy offramp token receiver dapp
	tokenReceiverAddress, tx, _, err := receiver_dapp.DeployReceiverDapp(client.Owner, client.Client, offRamp.Address(), client.LinkToken)
	require.NoError(t, err)
	WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	client.Logger.Infof("Offramp token receiver dapp deployed on %s in tx: %s", tokenReceiverAddress.Hex(), helpers.ExplorerLink(client.ChainId.Int64(), tx.Hash()))
	client.TokenReceiver = tokenReceiverAddress

	// Deploy the message executor contract
	executorAddress, tx, _, err := offramp_executor.DeployOffRampExecutor(client.Owner, client.Client, offRamp.Address(), false)
	require.NoError(t, err)
	WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	client.Logger.Infof("OffRamp executor contract deployed on %s in tx: %s", executorAddress.Hex(), helpers.ExplorerLink(client.ChainId.Int64(), tx.Hash()))
	client.OffRampExecutor = executorAddress

	return tokenReceiverAddress
}

func deployOnRamp(t *testing.T, client *EvmChainConfig, destinationChain *big.Int, poolAddresses []common.Address, feedAddresses []common.Address, afn common.Address) *onramp.OnRamp {
	client.Logger.Infof("Deploying onramp: destinationChains %+v, bridgeTokens %+v, poolAddresses %+v, priceFeeds %+v", destinationChain, client.BridgeTokens, poolAddresses, feedAddresses)
	onRampAddress, tx, _, err := onramp.DeployOnRamp(
		client.Owner,                  // user
		client.Client,                 // client
		client.ChainId,                // source chain id
		destinationChain,              // destinationChainId
		client.BridgeTokens,           // tokens
		poolAddresses,                 // pools
		feedAddresses,                 // Feeds
		[]common.Address{},            // allow list
		afn,                           // AFN
		big.NewInt(defaultAFNTimeout), // max timeout without AFN signal
		onramp.OnRampInterfaceOnRampConfig{
			RelayingFeeJuels: 0,
			MaxDataSize:      1e6,
			MaxTokensLength:  5,
		},
	)
	require.NoError(t, err)
	WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)

	onRamp, err := onramp.NewOnRamp(onRampAddress, client.Client)
	require.NoError(t, err)
	client.Logger.Infof(fmt.Sprintf("Onramp deployed on %s in tx %s", onRampAddress.String(), helpers.ExplorerLink(client.ChainId.Int64(), tx.Hash())))
	client.OnRamp = onRampAddress

	return onRamp
}

func deployOffRamp(t *testing.T, client *EvmChainConfig, sourceChain *big.Int, poolAddresses []common.Address, feedAddresses []common.Address, afn common.Address, sourceBridgeTokens []common.Address) *offramp.OffRamp {
	client.Logger.Infof("Deploying offramp: bridgeTokens %+v, poolAddresses %+v, priceFeeds %+v", client.BridgeTokens, client.TokenPools, client.PriceFeeds)
	offrampAddress, tx, _, err := offramp.DeployOffRamp(
		client.Owner,                  // user
		client.Client,                 // client
		sourceChain,                   // source chain id
		client.ChainId,                // dest chain id
		sourceBridgeTokens,            // source tokens
		poolAddresses,                 // dest pool addresses
		feedAddresses,                 // Feeds
		afn,                           // AFN address
		big.NewInt(defaultAFNTimeout), // max timeout without AFN signal
		offramp.OffRampInterfaceOffRampConfig{
			ExecutionFeeJuels:     0,
			ExecutionDelaySeconds: 0,
			MaxDataSize:           1e6,
			MaxTokensLength:       5,
		},
	)
	require.NoError(t, err)
	WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
	client.Logger.Infof("Offramp deployed on %s in tx: %s", offrampAddress.Hex(), helpers.ExplorerLink(client.ChainId.Int64(), tx.Hash()))
	client.OffRamp = offrampAddress

	offRamp, err := offramp.NewOffRamp(offrampAddress, client.Client)
	require.NoError(t, err)
	return offRamp
}

func deployNativeTokenPool(t *testing.T, client *EvmChainConfig) []*native_token_pool.NativeTokenPool {
	var pools []*native_token_pool.NativeTokenPool
	var poolAddresses []common.Address

	for i, bridgeToken := range client.BridgeTokens {
		if client.DeploySettings.DeployTokenPools {
			tenCoins := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(10))
			client.Logger.Infof("Deploying token pool for token %s", bridgeToken.Hex())
			lockConfig := native_token_pool.PoolInterfaceBucketConfig{
				Rate:     tenCoins,
				Capacity: tenCoins,
			}
			releaseConfig := native_token_pool.PoolInterfaceBucketConfig{
				Rate:     tenCoins,
				Capacity: tenCoins,
			}
			tokenPoolAddress, tx, _, err := native_token_pool.DeployNativeTokenPool(client.Owner, client.Client, bridgeToken, lockConfig, releaseConfig)
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

func deployPriceFeed(t *testing.T, client *EvmChainConfig) []common.Address {
	var priceFeeds []common.Address

	for _, feed := range client.PriceFeeds {
		if client.DeploySettings.DeployPriceFeeds {
			address, tx, _, err := mock_v3_aggregator_contract.DeployMockV3AggregatorContract(client.Owner, client.Client, 18, big.NewInt(6e12))
			require.NoError(t, err)
			WaitForMined(t, client.Logger, client.Client, tx.Hash(), true)
			client.Logger.Infof("Mock feed deployed on %s in tx: %s", address, helpers.ExplorerLink(client.ChainId.Int64(), tx.Hash()))
			priceFeeds = append(priceFeeds, address)
		} else {
			if feed.Hex() == "0x0000000000000000000000000000000000000000" {
				t.Error("deploy new price feed set to false but no price feed given in config")
			}
			priceFeeds = append(priceFeeds, feed)
		}
	}

	client.PriceFeeds = priceFeeds
	return priceFeeds
}

func fillPoolWithTokens(t *testing.T, client *EvmChainConfig, pool *native_token_pool.NativeTokenPool) {
	destLinkToken, err := link_token_interface.NewLinkToken(client.LinkToken, client.Client)
	require.NoError(t, err)

	// fill offramp token pool with 0.05 LINK
	amount := big.NewInt(5e16)
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
TokenSenders: %s,
OnRamp:       common.HexToAddress("%s"),
OffRamp:      common.Address{},
Afn:          common.HexToAddress("%s"),
`, source.LinkToken, source.BridgeTokens, source.TokenPools, source.PriceFeeds, source.TokenSender, source.OnRamp, source.Afn)

	destination.Logger.Infof(`
Destination chain config	
	
LinkToken:       common.HexToAddress("%s"),
BridgeTokens:    %s,
TokenPools:      %s,
PriceFeeds:      %s,
OnRamp:          common.Address{},
OffRamp:         common.HexToAddress("%s"),
MessageReceiver: common.HexToAddress("%s"),
TokenReceiver:   common.HexToAddress("%s"),
MessageExecutor: common.HexToAddress("%s"),
Afn:             common.HexToAddress("%s"),
`, destination.LinkToken, destination.BridgeTokens, destination.TokenPools, destination.PriceFeeds, destination.OffRamp, destination.MessageReceiver, destination.TokenReceiver, destination.OffRampExecutor, destination.Afn)

	PrintJobSpecs(source.OnRamp, destination.OffRamp, destination.OffRampExecutor, source.ChainId, destination.ChainId)

}
