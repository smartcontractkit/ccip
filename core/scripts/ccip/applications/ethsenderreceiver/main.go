package main

import (
	"context"
	"flag"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
	chainsel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip/rebalancer/multienv"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/ether_sender_receiver"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/erc20"
)

func main() {
	switch os.Args[1] {
	case "deploy":
		cmd := flag.NewFlagSet("deploy", flag.ExitOnError)
		chainID := cmd.Uint64("chain-id", 0, "Chain ID")
		routerAddress := cmd.String("router-address", "", "Router address")

		helpers.ParseArgs(cmd, os.Args[2:], "chain-id", "router-address")
		env := multienv.New(false, false)
		_, tx, _, err := ether_sender_receiver.DeployEtherSenderReceiver(
			env.Transactors[*chainID],
			env.Clients[*chainID],
			common.HexToAddress(*routerAddress),
		)
		helpers.PanicErr(err)
		helpers.ConfirmContractDeployed(context.Background(), env.Clients[*chainID], tx, int64(*chainID))
	case "get-fee":
		cmd := flag.NewFlagSet("get-fee", flag.ExitOnError)
		chainID := cmd.Uint64("chain-id", 0, "Chain ID")
		destChainID := cmd.Uint64("dest-chain-id", 0, "Destination chain ID")
		senderReceiverAddress := cmd.String("sender-receiver-address", "", "Sender receiver address")
		// message data
		destReceiver := cmd.String("dest-receiver-address", "", "Destination receiver address")
		destEOA := cmd.String("dest-eoa-address", "", "Destination EOA address")
		tokenAddress := cmd.String("token-address", "", "Token address")
		feeToken := cmd.String("fee-token", "", "Fee token address")
		amount := cmd.String("amount", "", "Amount")

		helpers.ParseArgs(cmd, os.Args[2:],
			"chain-id",
			"dest-chain-id",
			"sender-receiver-address",
			"dest-receiver-address",
			"dest-eoa-address",
			"token-address",
			"fee-token",
			"amount",
		)

		destChain, ok := chainsel.ChainByEvmChainID(*destChainID)
		if !ok {
			panic(fmt.Sprintf("Unknown chain ID: %d", *destChainID))
		}

		env := multienv.New(false, false)
		senderReceiver, err := ether_sender_receiver.NewEtherSenderReceiver(common.HexToAddress(*senderReceiverAddress), env.Clients[*chainID])
		helpers.PanicErr(err)

		receiverBytes, err := utils.ABIEncode(`[{"type": "address"}]`, common.HexToAddress(*destReceiver))
		helpers.PanicErr(err)
		destEOABytes, err := utils.ABIEncode(`[{"type": "address"}]`, common.HexToAddress(*destEOA))
		helpers.PanicErr(err)
		fmt.Println("receiver bytes:", hexutil.Encode(receiverBytes),
			"dest eoa bytes:", hexutil.Encode(destEOABytes), "fee token:", common.HexToAddress(*feeToken))

		msg := ether_sender_receiver.ClientEVM2AnyMessage{
			Receiver: receiverBytes,
			Data:     destEOABytes,
			TokenAmounts: []ether_sender_receiver.ClientEVMTokenAmount{
				{
					Token:  common.HexToAddress(*tokenAddress),
					Amount: decimal.RequireFromString(*amount).BigInt(),
				},
			},
			FeeToken: common.HexToAddress(*feeToken),
			// ExtraArgs: nil, // will be filled in by the contract
		}
		fee, err := senderReceiver.GetFee(nil, destChain.Selector, msg)
		helpers.PanicErr(err)

		fmt.Println("fee is:", fee, "juels/wei")
	case "ccip-send":
		cmd := flag.NewFlagSet("ccip-send", flag.ExitOnError)
		chainID := cmd.Uint64("chain-id", 0, "Chain ID")
		destChainID := cmd.Uint64("dest-chain-id", 0, "Destination chain ID")
		senderReceiverAddress := cmd.String("sender-receiver-address", "", "Sender receiver address")
		// message data
		destReceiver := cmd.String("dest-receiver-address", "", "Destination receiver address")
		destEOA := cmd.String("dest-eoa-address", "", "Destination EOA address")
		tokenAddress := cmd.String("token-address", "", "Token address")
		feeToken := cmd.String("fee-token", "", "Fee token address")
		amount := cmd.String("amount", "", "Amount")

		helpers.ParseArgs(cmd, os.Args[2:],
			"chain-id",
			"dest-chain-id",
			"sender-receiver-address",
			"dest-receiver-address",
			"dest-eoa-address",
			"token-address",
			"fee-token",
			"amount",
		)

		destChain, ok := chainsel.ChainByEvmChainID(*destChainID)
		if !ok {
			panic(fmt.Sprintf("Unknown chain ID: %d", *destChainID))
		}

		env := multienv.New(false, false)
		senderReceiver, err := ether_sender_receiver.NewEtherSenderReceiver(common.HexToAddress(*senderReceiverAddress), env.Clients[*chainID])
		helpers.PanicErr(err)

		receiverBytes, err := utils.ABIEncode(`[{"type": "address"}]`, common.HexToAddress(*destReceiver))
		helpers.PanicErr(err)
		destEOABytes, err := utils.ABIEncode(`[{"type": "address"}]`, common.HexToAddress(*destEOA))
		helpers.PanicErr(err)
		feeTok := common.HexToAddress(*feeToken)
		msg := ether_sender_receiver.ClientEVM2AnyMessage{
			Receiver: receiverBytes,
			Data:     destEOABytes,
			TokenAmounts: []ether_sender_receiver.ClientEVMTokenAmount{
				{
					Token:  common.HexToAddress(*tokenAddress),
					Amount: decimal.RequireFromString(*amount).BigInt(),
				},
			},
			FeeToken: feeTok,
			// ExtraArgs: nil, // will be filled in by the contract
		}
		fee, err := senderReceiver.GetFee(nil, destChain.Selector, msg)
		helpers.PanicErr(err)

		fmt.Println("fee is:", fee, "juels/wei")

		if (feeTok == common.Address{}) {
			totalValue := new(big.Int).Add(fee, decimal.RequireFromString(*amount).BigInt())
			ethBalance, err := env.Clients[*chainID].BalanceAt(context.Background(), env.Transactors[*chainID].From, nil)
			helpers.PanicErr(err)
			if ethBalance.Cmp(totalValue) < 0 {
				panic(fmt.Sprintf("Insufficient balance to send: %s < %s, please get more ETH", ethBalance, totalValue))
			}

			fmt.Println("Sending with value:", totalValue.String(), "total balance:", ethBalance.String())
			// native is the fee token, so send with value.
			tx, err := senderReceiver.CcipSend(&bind.TransactOpts{
				From:   env.Transactors[*chainID].From,
				Signer: env.Transactors[*chainID].Signer,
				Value:  totalValue,
			}, destChain.Selector, msg)
			helpers.PanicErr(err)
			helpers.ConfirmTXMined(context.Background(), env.Clients[*chainID], tx, int64(*chainID),
				"ccip send native, msg value:", totalValue.String(), "fee:", fee.String())
		} else {
			// non-native fee token, so approve first then send.
			erc20Token, err := erc20.NewERC20(feeTok, env.Clients[*chainID])
			helpers.PanicErr(err)

			// check if we have enough to provide allowance
			balance, err := erc20Token.BalanceOf(nil, env.Transactors[*chainID].From)
			helpers.PanicErr(err)
			if balance.Cmp(fee) < 0 {
				panic(fmt.Sprintf("Insufficient balance to provide allowance: %s < %s, please get more tokens", balance, fee))
			}

			// approve
			tx, err := erc20Token.Approve(env.Transactors[*chainID], common.HexToAddress(*senderReceiverAddress), fee)
			helpers.PanicErr(err)
			helpers.ConfirmTXMined(context.Background(), env.Clients[*chainID], tx, int64(*chainID),
				"approving sender receiver to spend fee token, amount:", *amount)

			// send message cross chain
			tx, err = senderReceiver.CcipSend(&bind.TransactOpts{
				From:   env.Transactors[*chainID].From,
				Signer: env.Transactors[*chainID].Signer,
				Value:  decimal.RequireFromString(*amount).BigInt(),
			}, destChain.Selector, msg)
			helpers.PanicErr(err)
			helpers.ConfirmTXMined(context.Background(), env.Clients[*chainID], tx, int64(*chainID), "ccip send non-native")
		}
	}
}
