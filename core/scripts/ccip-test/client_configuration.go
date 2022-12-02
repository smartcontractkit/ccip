package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slices"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_ge_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_ge_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/ge_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/governance_dapp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/ping_pong_demo"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/receiver_dapp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/simple_message_receiver"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/subscription_sender_dapp"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/dione"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/shared"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/merklemulti"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/testhelpers"
	"github.com/smartcontractkit/chainlink/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/core/utils"
)

func (client *CCIPClient) wip(t *testing.T, sourceClient *rhea.EvmDeploymentConfig, destClient *rhea.EvmDeploymentConfig) {
}

func (client *CCIPClient) setRateLimiterConfig(t *testing.T) {
	tx, err := client.Source.OnRamp.SetRateLimiterConfig(client.Source.Owner, evm_2_evm_ge_onramp.AggregateRateLimiterInterfaceRateLimiterConfig{
		Rate:     new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e5)),
		Capacity: new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e9)),
	})
	require.NoError(t, err)
	shared.WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), true)

	tx, err = client.Dest.OffRamp.SetRateLimiterConfig(client.Dest.Owner, evm_2_evm_ge_offramp.AggregateRateLimiterInterfaceRateLimiterConfig{
		Rate:     new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e5)),
		Capacity: new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e9)),
	})
	require.NoError(t, err)
	shared.WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
}

func (client *CCIPClient) startPingPong(t *testing.T) {
	tx, err := client.Source.PingPongDapp.StartPingPong(client.Source.Owner)
	require.NoError(t, err)
	shared.WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
}

func (client *CCIPClient) setPingPongPaused(t *testing.T, paused bool) {
	tx, err := client.Source.PingPongDapp.SetPaused(client.Source.Owner, paused)
	require.NoError(t, err)
	shared.WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
}

func (client *CCIPClient) fundPingPong(t *testing.T) {
	// TODO fund with ge:
	// send tokens, approve them from dapp

	//fundingAmount := big.NewInt(1e18)
	//client.Dest.ApproveLinkFrom(t, client.Dest.Owner, client.Dest.OffRampRouter.Address(), fundingAmount)
	//tx, err := client.Dest.OffRampRouter.FundSubscription(client.Dest.Owner, client.Dest.PingPongDapp.Address(), fundingAmount)
	//require.NoError(t, err)
	//shared.WaitForMined(t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
	//client.Dest.logger.Infof(fmt.Sprintf("Ping pong funded %s", helpers.ExplorerLink(client.Dest.ChainId.Int64(), tx.Hash())))
}

type Client struct {
	Owner            *bind.TransactOpts
	Users            []*bind.TransactOpts
	Client           *ethclient.Client
	ChainId          *big.Int
	LinkToken        *link_token_interface.LinkToken
	LinkTokenAddress common.Address
	SupportedTokens  map[rhea.Token]EVMBridgedToken
	GovernanceDapp   *governance_dapp.GovernanceDapp
	PingPongDapp     *ping_pong_demo.PingPongDemo
	Afn              *afn_contract.AFNContract
	Router           *ge_router.GERouter
	logger           logger.Logger
	t                *testing.T
}

type EVMBridgedToken struct {
	Token common.Address
	Pool  *native_token_pool.NativeTokenPool
	Price *big.Int
}

type SourceClient struct {
	Client
	OnRamp     *evm_2_evm_ge_onramp.EVM2EVMGEOnRamp
	SenderDapp *subscription_sender_dapp.SubscriptionSenderDapp
}

func NewSourceClient(t *testing.T, config rhea.EvmDeploymentConfig) SourceClient {
	LinkToken, err := link_token_interface.NewLinkToken(config.ChainConfig.LinkToken, config.Client)
	require.NoError(t, err)

	supportedTokens := map[rhea.Token]EVMBridgedToken{}
	for token, tokenConfig := range config.ChainConfig.SupportedTokens {
		tokenPool, err2 := native_token_pool.NewNativeTokenPool(tokenConfig.Pool, config.Client)
		require.NoError(t, err2)
		supportedTokens[token] = EVMBridgedToken{
			Token: tokenConfig.Token,
			Pool:  tokenPool,
			Price: tokenConfig.Price,
		}
	}

	afn, err := afn_contract.NewAFNContract(config.ChainConfig.Afn, config.Client)
	require.NoError(t, err)
	onRamp, err := evm_2_evm_ge_onramp.NewEVM2EVMGEOnRamp(config.LaneConfig.OnRamp, config.Client)
	require.NoError(t, err)
	senderDapp, err := subscription_sender_dapp.NewSubscriptionSenderDapp(config.LaneConfig.TokenSender, config.Client)
	require.NoError(t, err)
	router, err := ge_router.NewGERouter(config.ChainConfig.Router, config.Client)
	require.NoError(t, err)
	governanceDapp, err := governance_dapp.NewGovernanceDapp(config.LaneConfig.GovernanceDapp, config.Client)
	require.NoError(t, err)
	pingPongDapp, err := ping_pong_demo.NewPingPongDemo(config.LaneConfig.PingPongDapp, config.Client)
	require.NoError(t, err)

	return SourceClient{
		Client: Client{
			Client:           config.Client,
			ChainId:          config.ChainConfig.ChainId,
			LinkTokenAddress: config.ChainConfig.LinkToken,
			LinkToken:        LinkToken,
			Afn:              afn,
			SupportedTokens:  supportedTokens,
			GovernanceDapp:   governanceDapp,
			PingPongDapp:     pingPongDapp,
			Router:           router,
			logger:           config.Logger,
			t:                t,
		},
		OnRamp:     onRamp,
		SenderDapp: senderDapp,
	}
}

type DestClient struct {
	Client
	CommitStore     *commit_store.CommitStore
	MessageReceiver *simple_message_receiver.SimpleMessageReceiver
	ReceiverDapp    *receiver_dapp.ReceiverDapp
	OffRamp         *evm_2_evm_ge_offramp.EVM2EVMGEOffRamp
}

func NewDestinationClient(t *testing.T, config rhea.EvmDeploymentConfig) DestClient {
	LinkToken, err := link_token_interface.NewLinkToken(config.ChainConfig.LinkToken, config.Client)
	require.NoError(t, err)

	supportedTokens := map[rhea.Token]EVMBridgedToken{}
	for token, tokenConfig := range config.ChainConfig.SupportedTokens {
		tokenPool, err2 := native_token_pool.NewNativeTokenPool(tokenConfig.Pool, config.Client)
		require.NoError(t, err2)
		supportedTokens[token] = EVMBridgedToken{
			Token: tokenConfig.Token,
			Pool:  tokenPool,
			Price: tokenConfig.Price,
		}
	}

	afn, err := afn_contract.NewAFNContract(config.ChainConfig.Afn, config.Client)
	require.NoError(t, err)
	commitStore, err := commit_store.NewCommitStore(config.LaneConfig.CommitStore, config.Client)
	require.NoError(t, err)
	offRamp, err := evm_2_evm_ge_offramp.NewEVM2EVMGEOffRamp(config.LaneConfig.OffRamp, config.Client)
	require.NoError(t, err)
	messageReceiver, err := simple_message_receiver.NewSimpleMessageReceiver(config.LaneConfig.MessageReceiver, config.Client)
	require.NoError(t, err)
	receiverDapp, err := receiver_dapp.NewReceiverDapp(config.LaneConfig.ReceiverDapp, config.Client)
	require.NoError(t, err)
	router, err := ge_router.NewGERouter(config.ChainConfig.Router, config.Client)
	require.NoError(t, err)
	governanceDapp, err := governance_dapp.NewGovernanceDapp(config.LaneConfig.GovernanceDapp, config.Client)
	require.NoError(t, err)
	pingPongDapp, err := ping_pong_demo.NewPingPongDemo(config.LaneConfig.PingPongDapp, config.Client)
	require.NoError(t, err)

	return DestClient{
		Client: Client{
			Client:           config.Client,
			ChainId:          config.ChainConfig.ChainId,
			LinkTokenAddress: config.ChainConfig.LinkToken,
			LinkToken:        LinkToken,
			SupportedTokens:  supportedTokens,
			GovernanceDapp:   governanceDapp,
			PingPongDapp:     pingPongDapp,
			Afn:              afn,
			logger:           config.Logger,
			Router:           router,
			t:                t,
		},
		CommitStore:     commitStore,
		MessageReceiver: messageReceiver,
		ReceiverDapp:    receiverDapp,
		OffRamp:         offRamp,
	}
}

// CCIPClient contains a source chain and destination chain client and implements many methods
// that are useful for testing CCIP functionality on chain.
type CCIPClient struct {
	Source SourceClient
	Dest   DestClient
}

// NewCcipClient returns a new CCIPClient with initialised source and destination clients.
func NewCcipClient(t *testing.T, sourceConfig rhea.EvmDeploymentConfig, destConfig rhea.EvmDeploymentConfig, ownerKey string, seedKey string) CCIPClient {
	source := NewSourceClient(t, sourceConfig)
	source.SetOwnerAndUsers(t, ownerKey, seedKey, sourceConfig.ChainConfig.GasSettings)
	dest := NewDestinationClient(t, destConfig)
	dest.SetOwnerAndUsers(t, ownerKey, seedKey, destConfig.ChainConfig.GasSettings)

	return CCIPClient{
		Source: source,
		Dest:   dest,
	}
}

func GetSetupChain(t *testing.T, ownerPrivateKey string, chain rhea.EvmDeploymentConfig) *rhea.EvmDeploymentConfig {
	chain.SetupChain(t, ownerPrivateKey)
	return &chain
}

// SetOwnerAndUsers sets the owner and 10 users on a given client. It also set the proper
// gas parameters on these users.
func (client *Client) SetOwnerAndUsers(t *testing.T, ownerPrivateKey string, seedKey string, gasSettings rhea.EVMGasSettings) {
	client.Owner = rhea.GetOwner(t, ownerPrivateKey, client.ChainId, gasSettings)

	var users []*bind.TransactOpts
	seedKeyWithoutFirstChar := seedKey[1:]
	fmt.Println("--- Addresses of the seed key")
	for i := 0; i <= 9; i++ {
		_, err := hex.DecodeString(strconv.Itoa(i) + seedKeyWithoutFirstChar)
		require.NoError(t, err)
		key, err := crypto.HexToECDSA(strconv.Itoa(i) + seedKeyWithoutFirstChar)
		require.NoError(t, err)
		user, err := bind.NewKeyedTransactorWithChainID(key, client.ChainId)
		require.NoError(t, err)
		rhea.SetGasFees(user, gasSettings)
		users = append(users, user)
		fmt.Println(user.From.Hex())
	}
	fmt.Println("---")

	client.Users = users
}

func (client *Client) TypeAndVersion(addr common.Address) (ccip.ContractType, semver.Version, error) {
	return ccip.TypeAndVersion(addr, client.Client)
}

func (client *Client) ApproveLinkFrom(t *testing.T, user *bind.TransactOpts, approvedFor common.Address, amount *big.Int) {
	client.logger.Warnf("Approving %d link for %s", amount.Int64(), approvedFor.Hex())
	tx, err := client.LinkToken.Approve(user, approvedFor, amount)
	require.NoError(t, err)

	shared.WaitForMined(client.t, client.logger, client.Client, tx.Hash(), true)
	client.logger.Warnf("Link approved %s", helpers.ExplorerLink(client.ChainId.Int64(), tx.Hash()))
}

func (client *Client) ApproveLink(t *testing.T, approvedFor common.Address, amount *big.Int) {
	client.ApproveLinkFrom(t, client.Owner, approvedFor, amount)
}

func (client *CCIPClient) ChangeGovernanceParameters(t *testing.T) {
	feeConfig := governance_dapp.GovernanceDappFeeConfig{
		FeeAmount:      big.NewInt(10),
		ChangedAtBlock: big.NewInt(0),
	}
	DestBlockNum := GetCurrentBlockNumber(client.Dest.Client.Client)
	sourceBlockNum := GetCurrentBlockNumber(client.Source.Client.Client)

	tx, err := client.Source.GovernanceDapp.VoteForNewFeeConfig(client.Source.Owner, feeConfig)
	require.NoError(t, err)
	sendRequest := WaitForCrossChainSendRequest(client.Source, sourceBlockNum, tx.Hash())
	client.WaitForCommit(t, DestBlockNum)
	client.WaitForExecution(t, DestBlockNum, sendRequest.Message.SequenceNumber)
}

func (client *CCIPClient) SendMessage(t *testing.T) {
	DestBlockNum := GetCurrentBlockNumber(client.Dest.Client.Client)

	// ABI encoded message
	bts, err := hex.DecodeString("00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000005626c616e6b000000000000000000000000000000000000000000000000000000")
	require.NoError(t, err)

	token := ge_router.CCIPEVMTokenAndAmount{
		Token:  client.Source.LinkTokenAddress,
		Amount: big.NewInt(1),
	}
	extraArgsV1, err := testhelpers.GetEVMExtraArgsV1(big.NewInt(3e5), false)
	require.NoError(t, err)

	msg := ge_router.CCIPEVM2AnyGEMessage{
		Receiver:         testhelpers.MustEncodeAddress(t, client.Dest.MessageReceiver.Address()),
		Data:             bts,
		TokensAndAmounts: []ge_router.CCIPEVMTokenAndAmount{token},
		ExtraArgs:        extraArgsV1,
		FeeToken:         client.Source.LinkTokenAddress,
	}

	tx, err := client.Source.Router.CcipSend(client.Source.Owner, client.Dest.ChainId, msg)
	require.NoError(t, err)
	shared.WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
	client.WaitForCommit(t, DestBlockNum)
}

func (client *CCIPClient) DonExecutionHappyPath(t *testing.T) {
	client.Source.logger.Infof("Starting cross chain tx with DON execution")

	tokenAmount := big.NewInt(500)
	client.Source.ApproveLink(t, client.Source.Router.Address(), tokenAmount)

	DestBlockNum := GetCurrentBlockNumber(client.Dest.Client.Client)
	crossChainRequest := client.SendToOnrampWithExecution(t, client.Source, client.Source.Owner, client.Dest.ReceiverDapp.Address(), tokenAmount)
	client.Source.logger.Infof("Don executed tx submitted with sequence number: %d", crossChainRequest.Message.SequenceNumber)

	client.WaitForCommit(t, DestBlockNum)
	client.WaitForExecution(t, DestBlockNum, crossChainRequest.Message.SequenceNumber)
}

func (client *CCIPClient) WaitForCommit(t *testing.T, DestBlockNum uint64) {
	client.Dest.logger.Infof("Waiting for commit")

	commitEvent := make(chan *commit_store.CommitStoreReportAccepted)
	sub, err := client.Dest.CommitStore.WatchReportAccepted(
		&bind.WatchOpts{
			Context: context.Background(),
			Start:   &DestBlockNum,
		},
		commitEvent,
	)
	require.NoError(t, err)
	defer sub.Unsubscribe()

	select {
	case event := <-commitEvent:
		client.Dest.logger.Infof("Commit in tx %s", helpers.ExplorerLink(client.Dest.ChainId.Int64(), event.Raw.TxHash))
		return
	case err = <-sub.Err():
		panic(err)
	}
}

func (client *CCIPClient) WaitForExecution(t *testing.T, DestBlockNum uint64, sequenceNumber uint64) {
	client.Dest.logger.Infof("Waiting for execution...")

	events := make(chan *evm_2_evm_ge_offramp.EVM2EVMGEOffRampExecutionStateChanged)
	sub, err := client.Dest.OffRamp.WatchExecutionStateChanged(
		&bind.WatchOpts{
			Context: context.Background(),
			Start:   &DestBlockNum,
		},
		events,
		[]uint64{sequenceNumber},
		[][32]byte{})
	require.NoError(t, err)
	defer sub.Unsubscribe()

	select {
	case event := <-events:
		client.Dest.logger.Infof("Execution in tx %s", helpers.ExplorerLink(client.Dest.ChainId.Int64(), event.Raw.TxHash))
		return
	case err = <-sub.Err():
		panic(err)
	}
}

func (client *CCIPClient) ExecuteManually(seqNr uint64) error {
	// Find the seq num
	// Find the corresponding commit report
	end := uint64(11436244)
	reportIterator, err := client.Dest.CommitStore.FilterReportAccepted(&bind.FilterOpts{
		Start: end - 10000,
		End:   &end,
	})
	if err != nil {
		return err
	}
	var onRampIdx int
	var report *commit_store.CCIPCommitReport
	for reportIterator.Next() {
		for i, onRamp := range reportIterator.Event.Report.OnRamps {
			if onRamp == client.Source.OnRamp.Address() {
				if reportIterator.Event.Report.Intervals[i].Min <= seqNr && reportIterator.Event.Report.Intervals[i].Max >= seqNr {
					onRampIdx = i
					report = &reportIterator.Event.Report
					fmt.Println("Found root")
					break
				}
			}
		}
	}
	reportIterator.Close()
	if report == nil {
		return errors.New("unable to find seq num")
	}
	ctx := hasher.NewKeccakCtx()
	leafHasher := ccip.NewGELeafHasher(client.Source.ChainId, client.Dest.ChainId, client.Source.OnRamp.Address(), ctx)
	// Get all seqNrs in that range.
	end = uint64(7651526)
	sendRequestedIterator, err := client.Source.OnRamp.FilterCCIPSendRequested(&bind.FilterOpts{
		Start: end - 10000,
		End:   &end,
	})
	if err != nil {
		return err
	}
	var leaves [][32]byte
	var curr, prove int
	var originalMsg []byte
	for sendRequestedIterator.Next() {
		// Assume in order?
		if sendRequestedIterator.Event.Message.SequenceNumber <= report.Intervals[onRampIdx].Max && sendRequestedIterator.Event.Message.SequenceNumber >= report.Intervals[onRampIdx].Min {
			fmt.Println("Found seq num", sendRequestedIterator.Event.Message.SequenceNumber, report.Intervals[onRampIdx])
			hash, err2 := leafHasher.HashLeaf(sendRequestedIterator.Event.Raw)
			if err2 != nil {
				return err2
			}
			leaves = append(leaves, hash)
			if sendRequestedIterator.Event.Message.SequenceNumber == seqNr {
				fmt.Printf("Found proving %d %+v\n", curr, sendRequestedIterator.Event.Message)
				originalMsg = sendRequestedIterator.Event.Raw.Data
				prove = curr
			}
			curr++
		}
	}
	sendRequestedIterator.Close()
	if originalMsg == nil {
		return errors.New("unable to find")
	}
	tree, err := merklemulti.NewTree(ctx, leaves)
	if err != nil {
		return err
	}
	innerProof := tree.Prove([]int{prove})
	if tree.Root() != report.MerkleRoots[onRampIdx] {
		return errors.New("inner root doesn't match")
	}
	outerTree, err := merklemulti.NewTree(ctx, report.MerkleRoots)
	if err != nil {
		return err
	}
	if outerTree.Root() != report.RootOfRoots {
		return errors.New("outer root doesn't match")
	}
	outerProof := outerTree.Prove([]int{onRampIdx})
	executionReport := evm_2_evm_ge_offramp.CCIPExecutionReport{
		SequenceNumbers:          []uint64{seqNr},
		TokenPerFeeCoinAddresses: []common.Address{client.Dest.LinkTokenAddress},
		TokenPerFeeCoin:          []*big.Int{big.NewInt(1)},
		EncodedMessages:          [][]byte{originalMsg},
		InnerProofs:              innerProof.Hashes,
		InnerProofFlagBits:       ccip.ProofFlagsToBits(innerProof.SourceFlags),
		OuterProofs:              outerProof.Hashes,
		OuterProofFlagBits:       ccip.ProofFlagsToBits(outerProof.SourceFlags),
	}
	tx, err := client.Dest.OffRamp.ManuallyExecute(client.Dest.Owner, executionReport)
	if err != nil {
		fmt.Printf("%+v err %v\n", executionReport, err)
		return err
	}
	fmt.Println(client.Dest.Owner.From, tx.Hash(), err)
	return nil
}

//func (client CCIPClient) ExternalExecutionHappyPath(t *testing.T) {
//	ctx := context.Background()
//	offrampBlockNumber := GetCurrentBlockNumber(client.Dest.Client.Client)
//	onrampBlockNumber := GetCurrentBlockNumber(client.Source.Client.Client)
//
//	amount, _ := new(big.Int).SetString("10", 10)
//	client.Source.ApproveLink(t, client.Source.OnRamRouter.Address(), amount)
//
//	onrampRequest := client.SendToOnrampWithExecution(client.Source, client.Source.Owner, client.Dest.Owner.From, amount, common.HexToAddress("0x0000000000000000000000000000000000000000"))
//	sequenceNumber := onrampRequest.Message.SequenceNumber
//
//	// Gets the report that our transaction is included in
//	client.Dest.logger.Info("Getting report")
//	report, err := client.GetReportForSequenceNumber(ctx, sequenceNumber, offrampBlockNumber)
//	require.NoError(t, err)
//
//	// Get all requests included in the given report
//	client.Dest.logger.Info("Getting recent cross chain requests")
//	requests := client.GetCrossChainSendRequestsForRange(ctx, t, report, onrampBlockNumber)
//
//	// Generate the proof
//	client.Dest.logger.Info("Generating proof")
//	proof := client.ValidateMerkleRoot(t, onrampRequest, requests, report)
//
//	// Execute the transaction on the offramp
//	client.Dest.logger.Info("Executing offramp TX")
//	tx, err := client.ExecuteOffRampTransaction(t, proof, onrampRequest.Raw.Data)
//	require.NoError(t, err)
//
//	WaitForMined(t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
//	client.Dest.logger.Infof("Cross chain tx sent %s", helpers.ExplorerLink(client.Dest.ChainId.Int64(), tx.Hash()))
//}

func (client *CCIPClient) SendDappTx(t *testing.T) {
	amount := big.NewInt(500)
	destBlockNumber := GetCurrentBlockNumber(client.Dest.Client.Client)

	client.Source.ApproveLink(t, client.Source.SenderDapp.Address(), amount)
	crossChainRequest := client.SendToDappWithExecution(t, client.Source, client.Source.Owner, client.Dest.Owner.From, amount)
	client.WaitForCommit(t, destBlockNumber)
	client.WaitForExecution(t, destBlockNumber, crossChainRequest.Message.SequenceNumber)
}

// ScalingAndBatching should scale so that we see batching on the nodes
func (client *CCIPClient) ScalingAndBatching(t *testing.T) {
	amount := big.NewInt(10)
	toAddress := common.HexToAddress("0x57359120D900fab8cE74edC2c9959b21660d3887")
	DestBlockNum := GetCurrentBlockNumber(client.Dest.Client.Client)
	var seqNum uint64

	var wg sync.WaitGroup
	for _, user := range client.Source.Users {
		wg.Add(1)
		go func(user *bind.TransactOpts) {
			defer wg.Done()
			client.Source.ApproveLinkFrom(t, user, client.Source.SenderDapp.Address(), amount)
			crossChainRequest := client.SendToDappWithExecution(t, client.Source, user, toAddress, amount)
			client.Source.logger.Info("Don executed tx submitted with sequence number: ", crossChainRequest.Message.SequenceNumber)
			seqNum = crossChainRequest.Message.SequenceNumber
		}(user)
	}
	wg.Wait()
	client.WaitForCommit(t, DestBlockNum)
	client.WaitForExecution(t, DestBlockNum, seqNum)
	client.Source.logger.Info("Sent 10 txs to onramp.")
}

//func (client CCIPClient) ExecuteOffRampTransaction(t *testing.T, proof merklemulti.Proof[[32]byte], encodedMessage []byte) (*types.Transaction, error) {
//	decodedMsg, err := ccip.DecodeCCIPMessage(encodedMessage)
//	require.NoError(t, err)
//	_, err = ccip.MakeTollCCIPMsgArgs().PackValues([]interface{}{*decodedMsg})
//	require.NoError(t, err)
//
//	client.Dest.logger.Infof("Cross chain message %+v", decodedMsg)
//
//	report := any_2_evm_toll_offramp.CCIPExecutionReport{
//		Messages:       []any_2_evm_toll_offramp.CCIPAny2EVMTollMessage{*decodedMsg},
//		Proofs:         proof.Hashes,
//		ProofFlagsBits: ccip.ProofFlagsToBits(proof.SourceFlags),
//	}
//
//	tx, err := client.Dest.CommitStore.ExecuteTransaction(client.Dest.Owner, report, false)
//	if err != nil {
//		reason, err2 := evmclient.ExtractRevertReasonFromRPCError(err)
//		require.NoError(t, err2)
//		client.Dest.logger.Errorf("Extracting revert reason \"%s\" err \"%s\"", reason, err)
//	}
//
//	return tx, errors.Wrap(err, "Executing offramp tx")
//}

//func (client CCIPClient) GetCrossChainSendRequestsForRange(
//	ctx context.Context,
//	t *testing.T,
//	report commit_store.CCIPCommitReport,
//	onrampBlockNumber uint64) []*evm_2_evm_toll_onramp.EVM2EVMTollOnRampCCIPSendRequested {
//	// Get the other transactions in the proof, we look 1000 blocks back for transaction
//	// should be fine? Needs fine-tuning after improved batching strategies are developed
//	// in milestone 4
//	reqsIterator, err := client.Source.OnRamp.FilterCCIPSendRequested(&bind.FilterOpts{
//		Context: ctx,
//		Start:   onrampBlockNumber - 1000,
//	})
//	require.NoError(t, err)
//
//	var requests []*evm_2_evm_toll_onramp.EVM2EVMTollOnRampCCIPSendRequested
//	var minFound = report.MaxSequenceNumber
//
//	for reqsIterator.Next() {
//		num := reqsIterator.Event.Message.SequenceNumber
//		if num < minFound {
//			minFound = num
//		}
//		if num >= report.MinSequenceNumber && num <= report.MaxSequenceNumber {
//			requests = append(requests, reqsIterator.Event)
//		}
//	}
//
//	// TODO: Even if this check passes, we may not have fetched all necessary requests if
//	// minFound == report.MinSequenceNumber
//	if minFound > report.MinSequenceNumber {
//		t.Log("Not all cross chain requests found in the last 1000 blocks")
//		t.FailNow()
//	}
//
//	return requests
//}

//// GetReportForSequenceNumber return the offramp.CCIPCommitReport for a given ccip requests sequence number.
//func (client CCIPClient) GetReportForSequenceNumber(ctx context.Context, sequenceNumber uint64, minBlockNumber uint64) (commit_store.CCIPCommitReport, error) {
//	client.Dest.logger.Infof("Looking for sequenceNumber %d", sequenceNumber)
//	report, err := client.Dest.OffRamp.GetLastReport(&bind.CallOpts{Context: ctx, Pending: false})
//	if err != nil {
//		return commit_store.CCIPCommitReport{}, err
//	}
//
//	client.Dest.logger.Infof("Last report found for range %d-%d", report.MinSequenceNumber, report.MaxSequenceNumber)
//	// our tx is in the latest report
//	if sequenceNumber >= report.MinSequenceNumber && sequenceNumber <= report.MaxSequenceNumber {
//		return report, nil
//	}
//	// report isn't out yet, it will be in a future report
//	if sequenceNumber > report.MaxSequenceNumber {
//		maxIterations := CrossChainTimout / RetryTiming
//		for i := 0; i < int(maxIterations); i++ {
//			report, err = client.Dest.CommitStore.GetLastReport(&bind.CallOpts{Context: ctx, Pending: false})
//			if err != nil {
//				return commit_store.CCIPCommitReport{}, err
//			}
//			client.Dest.logger.Infof("Last report found for range %d-%d", report.MinSequenceNumber, report.MaxSequenceNumber)
//			if sequenceNumber >= report.MinSequenceNumber && sequenceNumber <= report.MaxSequenceNumber {
//				return report, nil
//			}
//			time.Sleep(RetryTiming)
//		}
//		return commit_store.CCIPCommitReport{}, errors.New("No report found within the given timeout")
//	}
//
//	// it is in a past report, start looking at the earliest block number possible, the one
//	// before we started the entire transaction on the onramp.
//	reports, err := client.Dest.CommitStore.FilterReportAccepted(&bind.FilterOpts{
//		Start:   minBlockNumber,
//		End:     nil,
//		Context: ctx,
//	})
//	if err != nil {
//		return commit_store.CCIPCommitReport{}, err
//	}
//
//	for reports.Next() {
//		report = reports.Event.Report
//		if sequenceNumber >= report.MinSequenceNumber && sequenceNumber <= report.MaxSequenceNumber {
//			return report, nil
//		}
//	}
//
//	// Somehow the transaction was not included in any report within blocks produced after
//	// the transaction was initialized but the sequence number is lower than we are currently at
//	return commit_store.CCIPCommitReport{}, errors.New("No report found for given sequence number")
//}

func (client *CCIPClient) SetCommitStoreConfig(t *testing.T) {
	config := commit_store.CommitStoreInterfaceCommitStoreConfig{
		OnRamps:          []common.Address{client.Source.OnRamp.Address()},
		MinSeqNrByOnRamp: []uint64{3},
	}
	tx, err := client.Dest.CommitStore.SetConfig(client.Dest.Owner, config)
	require.NoError(t, err)
	shared.WaitForMined(t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
}

func GetCurrentBlockNumber(chain *ethclient.Client) uint64 {
	blockNumber, err := chain.BlockNumber(context.Background())
	helpers.PanicErr(err)
	return blockNumber
}

func (client *CCIPClient) ValidateMerkleRoot(
	t *testing.T,
	request *evm_2_evm_ge_onramp.EVM2EVMGEOnRampCCIPSendRequested,
	reportRequests []*evm_2_evm_ge_onramp.EVM2EVMGEOnRampCCIPSendRequested,
	report commit_store.CCIPCommitReport,
) merklemulti.Proof[[32]byte] {
	mctx := hasher.NewKeccakCtx()
	var leafHashes [][32]byte
	for _, req := range reportRequests {
		leafHashes = append(leafHashes, mctx.Hash(req.Raw.Data))
	}

	tree, err := merklemulti.NewTree(mctx, leafHashes)
	require.NoError(t, err)
	rootIndex := -1
	for i, root := range report.MerkleRoots {
		if tree.Root() == root {
			rootIndex = i
		}

	}
	if rootIndex < 0 {
		t.Log("Merkle root does not match any root in the report")
		t.FailNow()
	}

	exists, err := client.Dest.CommitStore.GetMerkleRoot(nil, tree.Root())
	require.NoError(t, err)
	if exists.Uint64() < 1 {
		panic("Path is not present in the offramp")
	}
	index := request.Message.SequenceNumber - report.Intervals[rootIndex].Min
	client.Dest.logger.Info("index is ", index)
	return tree.Prove([]int{int(index)})
}

// SendToDappWithExecution executes a cross chain transactions using the sender dapp interface.
func (client *CCIPClient) SendToDappWithExecution(t *testing.T, source SourceClient, from *bind.TransactOpts, toAddress common.Address, amount *big.Int) *evm_2_evm_ge_onramp.EVM2EVMGEOnRampCCIPSendRequested {
	SourceBlockNumber := GetCurrentBlockNumber(source.Client.Client)
	token := subscription_sender_dapp.CCIPEVMTokenAndAmount{
		Token:  client.Source.LinkTokenAddress,
		Amount: amount,
	}
	extraArgsV1, err := testhelpers.GetEVMExtraArgsV1(big.NewInt(100_000), false)
	helpers.PanicErr(err)

	tx, err := source.SenderDapp.SendMessage(from, subscription_sender_dapp.CCIPEVM2AnySubscriptionMessage{
		Receiver:         testhelpers.MustEncodeAddress(t, toAddress),
		TokensAndAmounts: []subscription_sender_dapp.CCIPEVMTokenAndAmount{token},
		ExtraArgs:        extraArgsV1,
	})
	helpers.PanicErr(err)
	source.logger.Infof("Send tokens tx %s", helpers.ExplorerLink(source.ChainId.Int64(), tx.Hash()))

	return WaitForCrossChainSendRequest(source, SourceBlockNumber, tx.Hash())
}

// SendToOnrampWithExecution executes a cross chain transactions using the onramp interface.
func (client *CCIPClient) SendToOnrampWithExecution(t *testing.T, source SourceClient, from *bind.TransactOpts, toAddress common.Address, amount *big.Int) *evm_2_evm_ge_onramp.EVM2EVMGEOnRampCCIPSendRequested {
	SourceBlockNumber := GetCurrentBlockNumber(source.Client.Client)

	senderAndReceiver, err := utils.ABIEncode(`[{"type":"address"}, {"type":"address"}]`, source.Owner.From, source.Owner.From)
	helpers.PanicErr(err)

	extraArgsV1, err := testhelpers.GetEVMExtraArgsV1(big.NewInt(3e5), false)
	helpers.PanicErr(err)

	payload := ge_router.CCIPEVM2AnyGEMessage{
		TokensAndAmounts: []ge_router.CCIPEVMTokenAndAmount{},
		Receiver:         testhelpers.MustEncodeAddress(t, toAddress),
		Data:             senderAndReceiver,
		ExtraArgs:        extraArgsV1,
		FeeToken:         client.Source.LinkTokenAddress,
	}
	source.logger.Infof("Send tx with payload %+v", payload)

	tx, err := source.Router.CcipSend(from, client.Dest.ChainId, payload)
	if err != nil {
		t.Log(err.Error())
		printRevertReason(err, ge_router.GERouterABI)
	}
	helpers.PanicErr(err)
	source.logger.Infof("Send tokens tx %s", helpers.ExplorerLink(source.ChainId.Int64(), tx.Hash()))
	return WaitForCrossChainSendRequest(source, SourceBlockNumber, tx.Hash())
}

func printRevertReason(errorData interface{}, abiString string) {
	dataError := errorData.(rpc.DataError)
	data, err := hex.DecodeString(dataError.ErrorData().(string)[2:])
	helpers.PanicErr(err)
	jsonABI, err := abi.JSON(strings.NewReader(abiString))
	helpers.PanicErr(err)
	for k, abiError := range jsonABI.Errors {
		if bytes.Equal(data[:4], abiError.ID.Bytes()[:4]) {
			// Found a matching error
			v, err := abiError.Unpack(data)
			helpers.PanicErr(err)
			fmt.Printf("Error \"%v\" args \"%v\"\n", k, v)
			return
		}
	}
}

// WaitForCrossChainSendRequest checks on chain for a successful onramp send event with the given tx hash.
// If not immediately found it will keep retrying in intervals of the globally specified RetryTiming.
func WaitForCrossChainSendRequest(source SourceClient, fromBlockNum uint64, txhash common.Hash) *evm_2_evm_ge_onramp.EVM2EVMGEOnRampCCIPSendRequested {
	filter := bind.FilterOpts{Start: fromBlockNum}
	source.logger.Infof("Waiting for cross chain send... ")

	for {
		iterator, err := source.OnRamp.FilterCCIPSendRequested(&filter)
		helpers.PanicErr(err)
		for iterator.Next() {
			if iterator.Event.Raw.TxHash.Hex() == txhash.Hex() {
				source.logger.Infof("Cross chain send event found in tx: %s ", helpers.ExplorerLink(source.ChainId.Int64(), txhash))
				return iterator.Event
			}
		}
		time.Sleep(shared.RetryTiming)
	}
}

func (client *CCIPClient) PauseOfframpPool() {
	for _, tokenConfig := range client.Dest.SupportedTokens {
		paused, err := tokenConfig.Pool.Paused(nil)
		helpers.PanicErr(err)
		if paused {
			return
		}
		client.Dest.logger.Info("pausing offramp pool...")
		tx, err := tokenConfig.Pool.Pause(client.Dest.Owner)
		helpers.PanicErr(err)
		client.Dest.logger.Info("Offramp pool paused, tx hash: %s", tx.Hash())
		shared.WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
	}
}
func (client *CCIPClient) PauseOnrampPool() {
	for _, tokenConfig := range client.Source.SupportedTokens {
		paused, err := tokenConfig.Pool.Paused(nil)
		helpers.PanicErr(err)
		if paused {
			return
		}
		client.Source.logger.Info("pausing onramp pool...")
		tx, err := tokenConfig.Pool.Pause(client.Source.Owner)
		helpers.PanicErr(err)
		client.Source.logger.Info("Onramp pool paused, tx hash: %s", tx.Hash())
		shared.WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
	}
}

func (client *CCIPClient) UnpauseOfframpPool() {
	for _, tokenConfig := range client.Dest.SupportedTokens {
		paused, err := tokenConfig.Pool.Paused(nil)
		helpers.PanicErr(err)
		if !paused {
			return
		}
		client.Dest.logger.Info("unpausing offramp pool...")
		tx, err := tokenConfig.Pool.Unpause(client.Dest.Owner)
		helpers.PanicErr(err)
		client.Dest.logger.Info("Offramp pool unpaused, tx hash: %s", tx.Hash())
		shared.WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
	}
}

func (client *CCIPClient) UnpauseOnrampPool() {
	for _, tokenConfig := range client.Source.SupportedTokens {
		paused, err := tokenConfig.Pool.Paused(nil)
		helpers.PanicErr(err)
		if !paused {
			return
		}
		client.Source.logger.Info("unpausing onramp pool...")
		tx, err := tokenConfig.Pool.Unpause(client.Source.Owner)
		helpers.PanicErr(err)
		client.Source.logger.Info("Onramp pool unpaused, tx hash: %s", tx.Hash())
		shared.WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
	}
}

func (client *CCIPClient) PauseOnramp() {
	paused, err := client.Source.OnRamp.Paused(nil)
	helpers.PanicErr(err)
	if paused {
		return
	}
	client.Source.logger.Info("pausing onramp...")
	tx, err := client.Source.OnRamp.Pause(client.Source.Owner)
	helpers.PanicErr(err)
	client.Source.logger.Info("Onramp paused, tx hash: %s", tx.Hash())
	shared.WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
}

func (client *CCIPClient) PauseCommitStore() {
	paused, err := client.Dest.CommitStore.Paused(nil)
	helpers.PanicErr(err)
	if paused {
		return
	}
	client.Dest.logger.Info("pausing offramp...")
	tx, err := client.Dest.CommitStore.Pause(client.Dest.Owner)
	helpers.PanicErr(err)
	client.Dest.logger.Info("Offramp paused, tx hash: %s", tx.Hash())
	shared.WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
}

func (client *CCIPClient) UnpauseOnramp() {
	paused, err := client.Source.OnRamp.Paused(nil)
	helpers.PanicErr(err)
	if !paused {
		return
	}
	client.Source.logger.Info("unpausing onramp...")
	tx, err := client.Source.OnRamp.Unpause(client.Source.Owner)
	helpers.PanicErr(err)
	client.Source.logger.Info("Onramp unpaused, tx hash: %s", tx.Hash())
	shared.WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
}

func (client *CCIPClient) UnpauseCommitStore() {
	paused, err := client.Dest.CommitStore.Paused(nil)
	helpers.PanicErr(err)
	if !paused {
		return
	}
	client.Dest.logger.Info("unpausing offramp...")
	tx, err := client.Dest.CommitStore.Unpause(client.Dest.Owner)
	helpers.PanicErr(err)
	client.Dest.logger.Info("Offramp unpaused, tx hash: %s", tx.Hash())
	shared.WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
}

func (client *CCIPClient) UnpauseAll() {
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer wg.Done()
		client.UnpauseOnramp()
	}()
	go func() {
		defer wg.Done()
		client.UnpauseCommitStore()
	}()
	go func() {
		defer wg.Done()
		client.UnpauseOnrampPool()
	}()
	go func() {
		defer wg.Done()
		client.UnpauseOfframpPool()
	}()
	wg.Wait()
}

func (client *CCIPClient) SetOCRConfig(env dione.Environment) {
	verifierOCRConfig, err := client.Dest.CommitStore.LatestConfigDetails(&bind.CallOpts{})
	helpers.PanicErr(err)
	if verifierOCRConfig.BlockNumber != 0 {
		client.Dest.logger.Infof("CommitStore OCR config already found: %+v", verifierOCRConfig.ConfigDigest)
		client.Dest.logger.Infof("The new config will overwrite the current one.")
	}

	rampOCRConfig, err := client.Dest.OffRamp.LatestConfigDetails(&bind.CallOpts{})
	helpers.PanicErr(err)
	if rampOCRConfig.BlockNumber != 0 {
		client.Dest.logger.Infof("OffRamp OCR config already found: %+v", rampOCRConfig.ConfigDigest)
		client.Dest.logger.Infof("The new config will overwrite the current one.")
	}

	ccipConfig, err := ccip.OffchainConfig{
		SourceIncomingConfirmations: 10,
		DestIncomingConfirmations:   10,
	}.Encode()
	helpers.PanicErr(err)

	don := dione.NewOfflineDON(env, client.Dest.logger)
	faults := len(don.Config.Nodes) / 3

	signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, err := confighelper2.ContractSetConfigArgsForTests(
		70*time.Second, // deltaProgress
		5*time.Second,  // deltaResend
		30*time.Second, // deltaRound
		2*time.Second,  // deltaGrace
		40*time.Second, // deltaStage
		3,
		[]int{1, 1, 2, 3}, // Transmission schedule: 1 oracle in first deltaStage, 2 in the second and so on.
		don.GenerateOracleIdentities(client.Dest.ChainId.String()),
		ccipConfig,
		5*time.Second,
		32*time.Second,
		20*time.Second,
		10*time.Second,
		10*time.Second,
		faults,
		nil,
	)
	helpers.PanicErr(err)

	signerAddresses, err := ocrcommon.OnchainPublicKeyToAddress(signers)
	helpers.PanicErr(err)
	transmitterAddresses, err := ocrcommon.AccountToAddress(transmitters)
	helpers.PanicErr(err)

	tx, err := client.Dest.CommitStore.SetConfig0(
		client.Dest.Owner,
		signerAddresses,
		transmitterAddresses,
		f,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)
	helpers.PanicErr(err)
	shared.WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
	client.Dest.logger.Infof("Config set on commitStore %s", helpers.ExplorerLink(client.Dest.ChainId.Int64(), tx.Hash()))

	tx, err = client.Dest.OffRamp.SetConfig(
		client.Dest.Owner,
		signerAddresses,
		transmitterAddresses,
		f,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)
	helpers.PanicErr(err)
	shared.WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
	client.Dest.logger.Infof("Config set on offramp %s", helpers.ExplorerLink(client.Dest.ChainId.Int64(), tx.Hash()))
}

func (client *CCIPClient) AcceptOwnership(t *testing.T) {
	tx, err := client.Dest.CommitStore.AcceptOwnership(client.Dest.Owner)
	require.NoError(t, err)
	shared.WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)

	tx, err = client.Dest.OffRamp.AcceptOwnership(client.Dest.Owner)
	require.NoError(t, err)
	shared.WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
}

type tokenPoolRegistry interface {
	Address() common.Address
	GetPoolTokens(opts *bind.CallOpts) ([]common.Address, error)
	GetPoolBySourceToken(opts *bind.CallOpts, token common.Address) (common.Address, error)
	RemovePool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)
	AddPool(opts *bind.TransactOpts, token common.Address, pool common.Address) (*types.Transaction, error)
}

type aggregateRateLimiter interface {
	Address() common.Address
	GetPricesForTokens(opts *bind.CallOpts, tokens []common.Address) ([]*big.Int, error)
	SetPrices(opts *bind.TransactOpts, tokens []common.Address, prices []*big.Int) (*types.Transaction, error)
}

func syncPools(client *Client, registry tokenPoolRegistry, bridgeTokens map[rhea.Token]EVMBridgedToken, txOpts *bind.TransactOpts) []*types.Transaction {
	registeredTokens, err := registry.GetPoolTokens(&bind.CallOpts{})
	require.NoError(client.t, err)

	pendingTxs := make([]*types.Transaction, 0)
	// remove registered tokenPools not present in config
	for _, token := range registeredTokens {
		found := false
		for _, bridgedToken := range bridgeTokens {
			if bridgedToken.Token == token {
				found = true
				break
			}
		}
		if !found {
			pool, err := registry.GetPoolBySourceToken(&bind.CallOpts{}, token)
			require.NoError(client.t, err)
			tx, err := registry.RemovePool(txOpts, token, pool)
			require.NoError(client.t, err)
			client.logger.Infof("removePool(token=%s, pool=%s) from registry=%s: tx=%s", token, pool, registry.Address(), tx.Hash())
			pendingTxs = append(pendingTxs, tx)           // queue txs for wait
			txOpts.Nonce.Add(txOpts.Nonce, big.NewInt(1)) // increment nonce
		}
	}
	// add tokenPools present in config and not yet registered
	for _, tokenConfig := range bridgeTokens {
		// remove tokenPools not present in config
		if !slices.Contains(registeredTokens, tokenConfig.Token) {
			pool := tokenConfig.Pool.Address()
			tx, err := registry.AddPool(txOpts, tokenConfig.Token, pool)
			require.NoError(client.t, err)
			client.logger.Infof("addPool(token=%s, pool=%s) from registry=%s: tx=%s", tokenConfig.Token, pool, registry.Address(), tx.Hash())
			pendingTxs = append(pendingTxs, tx)           // queue txs for wait
			txOpts.Nonce.Add(txOpts.Nonce, big.NewInt(1)) // increment nonce
		}
	}
	return pendingTxs
}

func syncPrices(client *Client, limiter aggregateRateLimiter, txOpts *bind.TransactOpts) *types.Transaction {
	// sync tokenPrices if needed
	if len(client.SupportedTokens) == 0 {
		return nil
	}

	var tokens []common.Address
	var prices []*big.Int
	for _, tokenConfig := range client.SupportedTokens {
		tokens = append(tokens, tokenConfig.Token)
		prices = append(prices, tokenConfig.Price)
	}

	limiterTokenPrices, err := limiter.GetPricesForTokens(&bind.CallOpts{}, tokens)
	require.NoError(client.t, err)
	i := 0
	for _, tokenConfig := range client.SupportedTokens {
		// on first difference, setPrices then return
		if tokenConfig.Price.Cmp(limiterTokenPrices[i]) != 0 {
			tx, err2 := limiter.SetPrices(txOpts, tokens, prices)
			require.NoError(client.t, err2)
			client.logger.Infof("setPrices(tokens=%s, prices=%s) for limiter=%s: tx=%s", tokens, prices, limiter.Address(), tx.Hash())
			txOpts.Nonce.Add(txOpts.Nonce, big.NewInt(1)) // increment nonce
			return tx
		}
		i++
	}
	return nil
}

func waitPendingTxs(client *Client, pendingTxs *[]*types.Transaction) {
	// wait for all queued txs
	for _, tx := range *pendingTxs {
		shared.WaitForMined(client.t, client.logger, client.Client, tx.Hash(), true)
	}
	*pendingTxs = (*pendingTxs)[:0] // clear pending txs
}

func (client *CCIPClient) SyncTokenPools(t *testing.T) {
	// use local txOpts, so we can cache/increment nonce manually before waiting on all txs
	sourceTxOpts := *client.Source.Owner
	sourceTxOpts.GasLimit = 120_000 // hardcode gasLimit (enough for each tx here), to avoid race from mis-estimating
	sourcePendingNonce, err := client.Source.Client.Client.PendingNonceAt(context.Background(), client.Source.Owner.From)
	require.NoError(t, err)
	sourceTxOpts.Nonce = big.NewInt(int64(sourcePendingNonce))

	// onRamp maps source tokens to source pools
	sourcePendingTxs := syncPools(&client.Source.Client, client.Source.OnRamp, client.Source.SupportedTokens, &sourceTxOpts)

	// same as above, for offRamp
	destTxOpts := *client.Dest.Owner
	destTxOpts.GasLimit = 120_000 // hardcode gasLimit (enough for each tx here), to avoid race from mis-estimating
	destPendingNonce, err := client.Dest.Client.Client.PendingNonceAt(context.Background(), client.Dest.Owner.From)
	require.NoError(t, err)
	destTxOpts.Nonce = big.NewInt(int64(destPendingNonce))

	// offRamp maps *source* tokens to *dest* pools
	destPendingTxs := syncPools(&client.Dest.Client, client.Dest.OffRamp, client.Source.SupportedTokens, &destTxOpts)

	waitPendingTxs(&client.Source.Client, &sourcePendingTxs)
	waitPendingTxs(&client.Dest.Client, &destPendingTxs)

	if tx := syncPrices(&client.Source.Client, client.Source.OnRamp, &sourceTxOpts); tx != nil {
		sourcePendingTxs = append(sourcePendingTxs, tx)
	}
	if tx := syncPrices(&client.Dest.Client, client.Dest.OffRamp, &destTxOpts); tx != nil {
		destPendingTxs = append(destPendingTxs, tx)
	}

	waitPendingTxs(&client.Source.Client, &sourcePendingTxs)
	waitPendingTxs(&client.Dest.Client, &destPendingTxs)
}
