package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/lock_unlock_pool"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/message_executor"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/simple_message_receiver"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_onramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_receiver"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_sender"
	ccip_shared "github.com/smartcontractkit/chainlink/core/scripts/ccip-test/ccip-shared"
)

func main() {
	// This key is used to deploy all contracts on both source and Dest chains
	k := os.Getenv("OWNER_KEY")
	if k == "" {
		panic("must set owner key")
	}

	// This method will deploy all source and Destination chain contracts using the
	// owner key. Only run this of the currently deployed contracts are outdated or
	// when initializing a new chain.
	deploySourceAndDestContracts(
		ccip_shared.Kovan.SetOwner(k).SetupClient(),
		ccip_shared.Rinkeby.SetOwner(k).SetupClient(),
	)
}

func deploySourceAndDestContracts(source *ccip_shared.EvmChainConfig, dest *ccip_shared.EvmChainConfig) {
	// 2 gwei to not use eip1559
	source.Owner.GasTipCap = big.NewInt(2000000000)
	dest.Owner.GasTipCap = big.NewInt(2000000000)

	// After running this code please update the configuration to reflect the newly
	// deployed contract addresses.
	onramp := deployOnramp(source, dest.ChainId, dest.LinkToken)
	fmt.Println("Onramp fully deployed:", onramp.Address().Hex())
	offramp, executor, singleTokenReceiver := deployOfframp(dest, source.ChainId)
	fmt.Println("Offramp fully deployed:", offramp.Address().Hex())

	// Deploy onramp EOA token sender
	eoaTokenSenderAddress, tx, _, err := single_token_sender.DeployEOASingleTokenSender(source.Owner, source.Client, onramp.Address(), singleTokenReceiver)
	ccip_shared.PanicErr(err)
	ccip_shared.WaitForMined(context.Background(), source.Client, tx.Hash(), true)
	fmt.Println("Onramp EOA token sender deployed:", eoaTokenSenderAddress.Hex())

	PrintJobSpecs(onramp.Address(), offramp.Address(), executor.Address())
}

func deployOnramp(source *ccip_shared.EvmChainConfig, offrampChainId *big.Int, offrampLinkTokenAddress common.Address) *single_token_onramp.SingleTokenOnRamp {
	sourcePool := deployLockUnlockPool(source, true)
	afn := deployAFN(source, true)

	// deploy onramp
	onRampAddress, tx, _, err := single_token_onramp.DeploySingleTokenOnRamp(
		source.Owner,                    // user
		source.Client,                   // client
		source.ChainId,                  // source chain id
		source.LinkToken,                // source token
		sourcePool.Address(),            // source pool
		offrampChainId,                  // dest chain id
		offrampLinkTokenAddress,         // remoteToken
		[]common.Address{},              // allow list
		false,                           // enableAllowList
		big.NewInt(1),                   // token bucket rate
		big.NewInt(1000000000000000000), // token bucket capacity, 1 LINK
		afn.Address(),                   // AFN
		// 86400 seconds = one day
		big.NewInt(86400), //maxTimeWithoutAFNSignal
	)
	ccip_shared.PanicErr(err)
	ccip_shared.WaitForMined(context.Background(), source.Client, tx.Hash(), true)

	onRamp, err := single_token_onramp.NewSingleTokenOnRamp(onRampAddress, source.Client)
	ccip_shared.PanicErr(err)
	fmt.Println("Onramp deployed on:", onRampAddress.String())

	// Configure onramp address on pool
	tx, err = sourcePool.SetOnRamp(source.Owner, onRampAddress, true)
	ccip_shared.PanicErr(err)

	fmt.Println("Onramp pool configured with onramp on:", tx.Hash().Hex())
	return onRamp
}

func deployOfframp(dest *ccip_shared.EvmChainConfig, onrampChainId *big.Int) (*single_token_offramp.SingleTokenOffRamp, *message_executor.MessageExecutor, common.Address) {
	pool := deployLockUnlockPool(dest, true)
	fillPoolWithLink(dest, pool)
	afn := deployAFN(dest, true)

	// deploy offramp on Rinkeby
	offrampAddress, tx, _, err := single_token_offramp.DeploySingleTokenOffRamp(
		dest.Owner,                      // user
		dest.Client,                     // client
		onrampChainId,                   // source chain id
		dest.ChainId,                    // dest chain id
		dest.LinkToken,                  // link token address
		pool.Address(),                  // dest pool address
		big.NewInt(1),                   // token bucket rate
		big.NewInt(1000000000000000000), // token bucket capacity
		afn.Address(),                   // AFN address
		// 86400 seconds = one day
		big.NewInt(86400), // max timeout without AFN signal
		big.NewInt(0),     // execution delay in seconds
	)
	ccip_shared.PanicErr(err)
	ccip_shared.WaitForMined(context.Background(), dest.Client, tx.Hash(), true)
	fmt.Println("Offramp deployed on:", offrampAddress.Hex())

	offramp, err := single_token_offramp.NewSingleTokenOffRamp(offrampAddress, dest.Client)
	ccip_shared.PanicErr(err)

	// Configure offramp address on pool
	tx, err = pool.SetOffRamp(dest.Owner, offramp.Address(), true)
	ccip_shared.PanicErr(err)
	fmt.Println("Offramp pool configured with offramp address, tx hash:", tx.Hash().Hex())

	// Deploy offramp contract token receiver
	messageReceiverAddress, tx, _, err := simple_message_receiver.DeploySimpleMessageReceiver(dest.Owner, dest.Client)
	ccip_shared.PanicErr(err)
	ccip_shared.WaitForMined(context.Background(), dest.Client, tx.Hash(), true)
	fmt.Println("Offramp contract message receiver deployed on:", messageReceiverAddress.Hex())

	// Deploy offramp EOA token receiver
	tokenReceiverAddress, tx, _, err := single_token_receiver.DeployEOASingleTokenReceiver(dest.Owner, dest.Client, offramp.Address())
	ccip_shared.PanicErr(err)
	ccip_shared.WaitForMined(context.Background(), dest.Client, tx.Hash(), true)
	fmt.Println("Offramp EOA token receiver deployed on:", tokenReceiverAddress.Hex())
	// Deploy the message executor ocr2 contract
	executorAddress, tx, _, err := message_executor.DeployMessageExecutor(dest.Owner, dest.Client, offramp.Address())
	ccip_shared.PanicErr(err)
	ccip_shared.WaitForMined(context.Background(), dest.Client, tx.Hash(), true)
	fmt.Println("Message executor ocr2 contract deployed on:", executorAddress.Hex())

	executor, err := message_executor.NewMessageExecutor(executorAddress, dest.Client)
	ccip_shared.PanicErr(err)

	return offramp, executor, tokenReceiverAddress
}

func deployLockUnlockPool(client *ccip_shared.EvmChainConfig, deployNew bool) *lock_unlock_pool.LockUnlockPool {
	if deployNew {
		address, tx, _, err := lock_unlock_pool.DeployLockUnlockPool(client.Owner, client.Client, client.LinkToken)
		ccip_shared.PanicErr(err)
		ccip_shared.WaitForMined(context.Background(), client.Client, tx.Hash(), true)
		fmt.Println("Lock/unlock pool deployed on:", address.Hex())
		pool, err := lock_unlock_pool.NewLockUnlockPool(address, client.Client)
		ccip_shared.PanicErr(err)
		return pool
	}
	if client.LockUnlockPool.Hex() == "0x0000000000000000000000000000000000000000" {
		ccip_shared.PanicErr(errors.New("deploy new lock unlock pool set to false but no lock unlock pool given in config"))
	}
	sourcePool, err := lock_unlock_pool.NewLockUnlockPool(client.LockUnlockPool, client.Client)
	ccip_shared.PanicErr(err)
	fmt.Println("Lock unlock pool loaded from:", sourcePool.Address().Hex())
	return sourcePool
}

func deployAFN(client *ccip_shared.EvmChainConfig, deployNew bool) *afn_contract.AFNContract {
	if deployNew {
		address, tx, _, err := afn_contract.DeployAFNContract(
			client.Owner,
			client.Client,
			[]common.Address{client.Owner.From},
			[]*big.Int{big.NewInt(1)},
			big.NewInt(1),
			big.NewInt(1),
		)
		ccip_shared.PanicErr(err)
		ccip_shared.WaitForMined(context.Background(), client.Client, tx.Hash(), true)
		fmt.Println("AFN deployed on:", address.Hex())
		afn, err := afn_contract.NewAFNContract(address, client.Client)
		ccip_shared.PanicErr(err)
		return afn
	}
	if client.Afn.Hex() == "0x0000000000000000000000000000000000000000" {
		ccip_shared.PanicErr(errors.New("deploy new afn set to false but no afn given in config"))
	}
	afn, err := afn_contract.NewAFNContract(client.Afn, client.Client)
	ccip_shared.PanicErr(err)
	fmt.Println("AFN loaded from:", afn.Address().Hex())
	return afn
}

func fillPoolWithLink(client *ccip_shared.EvmChainConfig, pool *lock_unlock_pool.LockUnlockPool) {
	destLinkToken, err := link_token_interface.NewLinkToken(client.LinkToken, client.Client)
	ccip_shared.PanicErr(err)

	// fill offramp token pool with 5 LINK
	amount, _ := new(big.Int).SetString("5000000000000000000", 10)
	tx, err := destLinkToken.Approve(client.Owner, pool.Address(), amount)
	ccip_shared.PanicErr(err)
	ccip_shared.WaitForMined(context.Background(), client.Client, tx.Hash(), true)

	tx, err = pool.LockOrBurn(client.Owner, client.Owner.From, amount)
	ccip_shared.PanicErr(err)
	ccip_shared.WaitForMined(context.Background(), client.Client, tx.Hash(), true)
	fmt.Println("Dest pool filled with funds, tx hash:", tx.Hash().Hex())
}
