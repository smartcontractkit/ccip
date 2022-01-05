package ccip_shared

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/lock_unlock_pool"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/message_executor"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/simple_message_receiver"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_onramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_receiver"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_sender"
)

type Client struct {
	Owner            *bind.TransactOpts
	Users            []*bind.TransactOpts
	Client           *ethclient.Client
	ChainId          *big.Int
	LinkToken        *link_token_interface.LinkToken
	LinkTokenAddress common.Address
	LockUnlockPool   *lock_unlock_pool.LockUnlockPool
	Afn              *afn_contract.AFNContract
}

type SourceClient struct {
	Client
	SingleTokenOnramp *single_token_onramp.SingleTokenOnRamp
	SingleTokenSender *single_token_sender.EOASingleTokenSender
}

func NewSourceClient(config *EvmChainConfig) SourceClient {
	LinkToken, err := link_token_interface.NewLinkToken(config.LinkToken, config.Client)
	PanicErr(err)
	lockUnlockPool, err := lock_unlock_pool.NewLockUnlockPool(config.LockUnlockPool, config.Client)
	PanicErr(err)
	afn, err := afn_contract.NewAFNContract(config.Afn, config.Client)
	PanicErr(err)
	singleTokenOnramp, err := single_token_onramp.NewSingleTokenOnRamp(config.SingleTokenOnramp, config.Client)
	PanicErr(err)
	simpleTokenSender, err := single_token_sender.NewEOASingleTokenSender(config.SingleTokenSender, config.Client)
	PanicErr(err)

	return SourceClient{
		Client: Client{
			Client:           config.Client,
			Users:            config.Users,
			ChainId:          config.ChainId,
			LinkTokenAddress: config.LinkToken,
			LinkToken:        LinkToken,
			Afn:              afn,
			LockUnlockPool:   lockUnlockPool,
		},
		SingleTokenOnramp: singleTokenOnramp,
		SingleTokenSender: simpleTokenSender,
	}
}

type DestClient struct {
	Client
	SingleTokenOfframp    *single_token_offramp.SingleTokenOffRamp
	SimpleMessageReceiver *simple_message_receiver.SimpleMessageReceiver
	SingleTokenReceiver   *single_token_receiver.EOASingleTokenReceiver
	MessageExecutor       *message_executor.MessageExecutor
}

func NewDestinationClient(config *EvmChainConfig) DestClient {
	LinkToken, err := link_token_interface.NewLinkToken(config.LinkToken, config.Client)
	PanicErr(err)
	lockUnlockPool, err := lock_unlock_pool.NewLockUnlockPool(config.LockUnlockPool, config.Client)
	PanicErr(err)
	afn, err := afn_contract.NewAFNContract(config.Afn, config.Client)
	PanicErr(err)
	singleTokenOfframp, err := single_token_offramp.NewSingleTokenOffRamp(config.SingleTokenOfframp, config.Client)
	PanicErr(err)
	messageExecutor, err := message_executor.NewMessageExecutor(config.MessageExecutor, config.Client)
	PanicErr(err)
	simpleMessageReceiver, err := simple_message_receiver.NewSimpleMessageReceiver(config.SimpleMessageReceiver, config.Client)
	PanicErr(err)
	singleTokenReceiver, err := single_token_receiver.NewEOASingleTokenReceiver(config.SingleTokenReceiver, config.Client)
	PanicErr(err)

	return DestClient{
		Client: Client{
			Owner:            config.Owner,
			Users:            config.Users,
			Client:           config.Client,
			ChainId:          config.ChainId,
			LinkTokenAddress: config.LinkToken,
			LinkToken:        LinkToken,
			LockUnlockPool:   lockUnlockPool,
			Afn:              afn,
		},
		SingleTokenOfframp:    singleTokenOfframp,
		SimpleMessageReceiver: simpleMessageReceiver,
		SingleTokenReceiver:   singleTokenReceiver,
		MessageExecutor:       messageExecutor,
	}
}

type CcipClient struct {
	Source SourceClient
	Dest   DestClient
}

func NewCcipClient(sourceConfig *EvmChainConfig, destConfig *EvmChainConfig) CcipClient {
	// to not use geth-only tip fee method
	// https://github.com/ethereum/go-ethereum/pull/23484
	var twoGwei = big.NewInt(2e9)
	sourceConfig.Owner.GasTipCap = twoGwei
	destConfig.Owner.GasTipCap = twoGwei
	for _, user := range sourceConfig.Users {
		user.GasTipCap = twoGwei
	}
	for _, user := range destConfig.Users {
		user.GasTipCap = twoGwei
	}

	source := NewSourceClient(sourceConfig.SetupClient())
	dest := NewDestinationClient(destConfig.SetupClient())

	return CcipClient{
		Source: source,
		Dest:   dest,
	}
}

func (client Client) AssureHealth() {
	standardAfnTimeout := int64(86400)
	status, err := client.Afn.GetLastHeartbeat(&bind.CallOpts{
		Pending: false,
		Context: nil,
	})
	PanicErr(err)
	timeNow := time.Now().Unix()

	if timeNow > status.Timestamp.Int64()+standardAfnTimeout {
		tx, err := client.Afn.VoteGood(client.Owner, big.NewInt(status.Round.Int64()+1))
		PanicErr(err)
		WaitForMined(context.Background(), client.Client, tx.Hash(), true)
		fmt.Printf("[HEALTH] client with chainId %d set healthy for %d hours\n", client.ChainId.Int64(), standardAfnTimeout/60/60)
	} else {
		fmt.Printf("[HEALTH] client with chainId %d is already healthy for %d more hours\n", client.ChainId.Int64(), (standardAfnTimeout-(timeNow-status.Timestamp.Int64()))/60/60)
	}
}

func (client Client) ApproveLinkFrom(user *bind.TransactOpts, approvedFor common.Address, amount *big.Int) {
	ctx := context.Background()
	tx, err := client.LinkToken.Approve(user, approvedFor, amount)
	PanicErr(err)

	WaitForMined(ctx, client.Client, tx.Hash(), true)
	fmt.Println("approve tx hash", tx.Hash().Hex())
}

func (client Client) ApproveLink(approvedFor common.Address, amount *big.Int) {
	client.ApproveLinkFrom(client.Owner, approvedFor, amount)
}
