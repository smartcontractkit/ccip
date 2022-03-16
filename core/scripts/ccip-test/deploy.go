package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/message_executor"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/mock_v3_aggregator_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/onramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/receiver_dapp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/sender_dapp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/simple_message_receiver"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
)

// deployContracts will deploy all source and Destination chain contracts using the
// owner key. Only run this of the currently deployed contracts are outdated or
// when initializing a new chain.
func deployContracts(ownerKey string) {
	source := Rinkeby
	source.Owner = GetOwner(ownerKey, source.ChainId, source.GasSettings)
	source.Client = GetClient(source.EthUrl)
	dest := Kovan
	dest.Owner = GetOwner(ownerKey, dest.ChainId, dest.GasSettings)
	dest.Client = GetClient(dest.EthUrl)

	deploySourceAndDestContracts(source, dest)
}

func deploySourceAndDestContracts(source EvmChainConfig, dest EvmChainConfig) {
	SetGasFees(source.Owner, source.GasSettings)
	SetGasFees(dest.Owner, dest.GasSettings)

	// After running this code please update the configuration to reflect the newly
	// deployed contract addresses.
	onRamp := deployOnramp(source, dest.ChainId)
	fmt.Println("Onramp fully deployed:", onRamp.Address().Hex())
	offRamp, executor, tokenReceiver := deployOfframp(dest, source.ChainId)
	fmt.Println("Offramp fully deployed:", offRamp.Address().Hex())

	// Deploy onramp sender dapp
	tokenSenderAddress, tx, _, err := sender_dapp.DeploySenderDapp(source.Owner, source.Client, onRamp.Address(), dest.ChainId, tokenReceiver)
	helpers.PanicErr(err)
	WaitForMined(context.Background(), source.Client, tx.Hash(), true)
	fmt.Println("Onramp token sender dapp deployed:", tokenSenderAddress.Hex())

	PrintJobSpecs(onRamp.Address(), offRamp.Address(), executor.Address(), source.ChainId, dest.ChainId)
}

func deployOnramp(source EvmChainConfig, offrampChainId *big.Int) *onramp.OnRamp {
	sourcePool := deployNativeTokenPool(source, true)
	afn := deployAFN(source, true)
	feedAddress := deployPriceFeed(source, true)

	// deploy onramp
	onRampAddress, tx, _, err := onramp.DeployOnRamp(
		source.Owner,                           // user
		source.Client,                          // client
		source.ChainId,                         // source chain id
		[]*big.Int{offrampChainId},             // destinationChainIds
		[]common.Address{source.LinkToken},     // tokens
		[]common.Address{sourcePool.Address()}, // pools
		[]common.Address{feedAddress},          // Feeds
		[]common.Address{},                     // allow list
		afn.Address(),                          // AFN
		big.NewInt(86400),                      //maxTimeWithoutAFNSignal 86400 seconds = one day
		onramp.OnRampInterfaceOnRampConfig{
			RelayingFeeJuels: 0,
			MaxDataSize:      1e6,
			MaxTokensLength:  5,
		},
	)
	helpers.PanicErr(err)
	WaitForMined(context.Background(), source.Client, tx.Hash(), true)

	onRamp, err := onramp.NewOnRamp(onRampAddress, source.Client)
	helpers.PanicErr(err)
	fmt.Println("Onramp deployed on:", onRampAddress.String())

	// Configure onramp address on pool
	tx, err = sourcePool.SetOnRamp(source.Owner, onRampAddress, true)
	helpers.PanicErr(err)

	fmt.Println("Onramp pool configured with onramp on:", tx.Hash().Hex())
	return onRamp
}

func deployOfframp(dest EvmChainConfig, onrampChainId *big.Int) (*offramp.OffRamp, *message_executor.MessageExecutor, common.Address) {
	pool := deployNativeTokenPool(dest, true)
	afn := deployAFN(dest, true)
	feedAddress := deployPriceFeed(dest, true)

	// deploy offramp on Rinkeby
	offrampAddress, tx, _, err := offramp.DeployOffRamp(
		dest.Owner,                       // user
		dest.Client,                      // client
		onrampChainId,                    // source chain id
		dest.ChainId,                     // dest chain id
		[]common.Address{dest.LinkToken}, // source tokens
		[]common.Address{pool.Address()}, // dest pool addresses
		[]common.Address{feedAddress},    // Feeds
		afn.Address(),                    // AFN address
		big.NewInt(86400),                // max timeout without AFN signal  86400 seconds = one day
		offramp.OffRampInterfaceOffRampConfig{
			ExecutionFeeJuels:     0,
			ExecutionDelaySeconds: 0,
			MaxDataSize:           1e6,
			MaxTokensLength:       5,
		},
	)
	helpers.PanicErr(err)
	WaitForMined(context.Background(), dest.Client, tx.Hash(), true)
	fmt.Println("Offramp deployed on:", offrampAddress.Hex())

	offRamp, err := offramp.NewOffRamp(offrampAddress, dest.Client)
	helpers.PanicErr(err)

	fillPoolWithLink(dest, pool)

	// Configure offramp address on pool
	tx, err = pool.SetOffRamp(dest.Owner, offRamp.Address(), true)
	helpers.PanicErr(err)
	fmt.Println("Offramp pool configured with offramp address, tx hash:", tx.Hash().Hex())

	// Deploy offramp contract token receiver
	messageReceiverAddress, tx, _, err := simple_message_receiver.DeploySimpleMessageReceiver(dest.Owner, dest.Client)
	helpers.PanicErr(err)
	WaitForMined(context.Background(), dest.Client, tx.Hash(), true)
	fmt.Println("Offramp contract message receiver deployed on:", messageReceiverAddress.Hex())

	// Deploy offramp token receiver dapp
	tokenReceiverAddress, tx, _, err := receiver_dapp.DeployReceiverDapp(dest.Owner, dest.Client, offRamp.Address(), dest.LinkToken)
	helpers.PanicErr(err)
	WaitForMined(context.Background(), dest.Client, tx.Hash(), true)
	fmt.Println("Offramp token receiver dapp deployed on:", tokenReceiverAddress.Hex())
	// Deploy the message executor contract
	executorAddress, tx, _, err := message_executor.DeployMessageExecutor(dest.Owner, dest.Client, offRamp.Address(), false)
	helpers.PanicErr(err)
	WaitForMined(context.Background(), dest.Client, tx.Hash(), true)
	fmt.Println("Message executor contract deployed on:", executorAddress.Hex())

	executor, err := message_executor.NewMessageExecutor(executorAddress, dest.Client)
	helpers.PanicErr(err)

	return offRamp, executor, tokenReceiverAddress
}

func deployNativeTokenPool(client EvmChainConfig, deployNew bool) *native_token_pool.NativeTokenPool {
	if deployNew {
		tenLink := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(10))
		address, tx, _, err := native_token_pool.DeployNativeTokenPool(client.Owner, client.Client, client.LinkToken,
			native_token_pool.PoolInterfaceBucketConfig{
				Rate:     big.NewInt(10),
				Capacity: tenLink,
			},
			native_token_pool.PoolInterfaceBucketConfig{
				Rate:     big.NewInt(10),
				Capacity: tenLink,
			})
		helpers.PanicErr(err)
		WaitForMined(context.Background(), client.Client, tx.Hash(), true)
		fmt.Println("Native token pool deployed on:", address.Hex())
		pool, err := native_token_pool.NewNativeTokenPool(address, client.Client)
		helpers.PanicErr(err)
		return pool
	}
	if client.TokenPool.Hex() == "0x0000000000000000000000000000000000000000" {
		helpers.PanicErr(errors.New("deploy new lock unlock pool set to false but no lock unlock pool given in config"))
	}
	sourcePool, err := native_token_pool.NewNativeTokenPool(client.TokenPool, client.Client)
	helpers.PanicErr(err)
	fmt.Println("Lock unlock pool loaded from:", sourcePool.Address().Hex())
	return sourcePool
}

func deployAFN(client EvmChainConfig, deployNew bool) *afn_contract.AFNContract {
	if deployNew {
		address, tx, _, err := afn_contract.DeployAFNContract(
			client.Owner,
			client.Client,
			[]common.Address{client.Owner.From},
			[]*big.Int{big.NewInt(1)},
			big.NewInt(1),
			big.NewInt(1),
		)
		helpers.PanicErr(err)
		WaitForMined(context.Background(), client.Client, tx.Hash(), true)
		fmt.Println("AFN deployed on:", address.Hex())
		afn, err := afn_contract.NewAFNContract(address, client.Client)
		helpers.PanicErr(err)
		return afn
	}
	if client.Afn.Hex() == "0x0000000000000000000000000000000000000000" {
		helpers.PanicErr(errors.New("deploy new afn set to false but no afn given in config"))
	}
	afn, err := afn_contract.NewAFNContract(client.Afn, client.Client)
	helpers.PanicErr(err)
	fmt.Println("AFN loaded from:", afn.Address().Hex())
	return afn
}

func deployPriceFeed(client EvmChainConfig, deployNew bool) common.Address {
	if deployNew {
		address, tx, _, err := mock_v3_aggregator_contract.DeployMockV3AggregatorContract(client.Owner, client.Client, 18, big.NewInt(6e12))
		helpers.PanicErr(err)
		WaitForMined(context.Background(), client.Client, tx.Hash(), true)
		fmt.Println("Mock feed deployed on:", address.Hex())
		return address
	}
	if client.PriceFeed.Hex() == "0x0000000000000000000000000000000000000000" {
		helpers.PanicErr(errors.New("deploy new price feed set to false but no price feed given in config"))
	}
	return client.PriceFeed
}

func fillPoolWithLink(client EvmChainConfig, pool *native_token_pool.NativeTokenPool) {
	destLinkToken, err := link_token_interface.NewLinkToken(client.LinkToken, client.Client)
	helpers.PanicErr(err)

	// fill offramp token pool with 0.5 LINK
	amount := big.NewInt(5e17)
	tx, err := destLinkToken.Approve(client.Owner, pool.Address(), amount)
	helpers.PanicErr(err)
	fmt.Println("Approving LINK on token pool:", tx.Hash().Hex())
	WaitForMined(context.Background(), client.Client, tx.Hash(), true)

	tx, err = pool.LockOrBurn(client.Owner, client.Owner.From, amount)
	helpers.PanicErr(err)
	WaitForMined(context.Background(), client.Client, tx.Hash(), true)
	fmt.Println("Dest pool filled with funds, tx hash:", tx.Hash().Hex())
}
