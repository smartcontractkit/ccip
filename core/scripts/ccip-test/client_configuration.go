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

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
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
	"github.com/smartcontractkit/chainlink/core/logger"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/merklemulti"
	"github.com/smartcontractkit/chainlink/core/services/ocrcommon"
)

type Client struct {
	Owner            *bind.TransactOpts
	Users            []*bind.TransactOpts
	Client           *ethclient.Client
	ChainId          *big.Int
	LinkToken        *link_token_interface.LinkToken
	LinkTokenAddress common.Address
	TokenPools       []*native_token_pool.NativeTokenPool
	Afn              *afn_contract.AFNContract
	logger           logger.Logger
	t                *testing.T
}

type SourceClient struct {
	Client
	OnRamp       *evm_2_evm_toll_onramp.EVM2EVMTollOnRamp
	OnRampRouter *evm_2_any_toll_onramp_router.EVM2AnyTollOnRampRouter
	SenderDapp   *toll_sender_dapp.TollSenderDapp
}

func NewSourceClient(t *testing.T, config EvmChainConfig) SourceClient {
	client := GetClient(t, config.EthUrl)
	LinkToken, err := link_token_interface.NewLinkToken(config.LinkToken, client)
	require.NoError(t, err)
	var tokenPools []*native_token_pool.NativeTokenPool
	for _, poolAddress := range config.TokenPools {
		tokenPool, err2 := native_token_pool.NewNativeTokenPool(poolAddress, client)
		require.NoError(t, err2)
		tokenPools = append(tokenPools, tokenPool)
	}

	afn, err := afn_contract.NewAFNContract(config.Afn, client)
	require.NoError(t, err)
	onRamp, err := evm_2_evm_toll_onramp.NewEVM2EVMTollOnRamp(config.OnRamp, client)
	require.NoError(t, err)
	tokenSender, err := toll_sender_dapp.NewTollSenderDapp(config.TokenSender, client)
	require.NoError(t, err)
	onRampRouter, err := evm_2_any_toll_onramp_router.NewEVM2AnyTollOnRampRouter(config.OnRampRouter, client)
	require.NoError(t, err)

	return SourceClient{
		Client: Client{
			Client:           client,
			ChainId:          config.ChainId,
			LinkTokenAddress: config.LinkToken,
			LinkToken:        LinkToken,
			Afn:              afn,
			TokenPools:       tokenPools,
			logger:           logger.TestLogger(t).Named(helpers.ChainName(config.ChainId.Int64())),
			t:                t,
		},
		OnRamp:       onRamp,
		OnRampRouter: onRampRouter,
		SenderDapp:   tokenSender,
	}
}

type DestClient struct {
	Client
	BlobVerifier    *blob_verifier.BlobVerifier
	MessageReceiver *simple_message_receiver.SimpleMessageReceiver
	ReceiverDapp    *receiver_dapp.ReceiverDapp
	OffRamp         *any_2_evm_toll_offramp.EVM2EVMTollOffRamp
	OffRampRouter   *any_2_evm_toll_offramp_router.Any2EVMTollOffRampRouter
}

func NewDestinationClient(t *testing.T, config EvmChainConfig) DestClient {
	client := GetClient(t, config.EthUrl)
	LinkToken, err := link_token_interface.NewLinkToken(config.LinkToken, client)
	require.NoError(t, err)

	var tokenPools []*native_token_pool.NativeTokenPool
	for _, poolAddress := range config.TokenPools {
		tokenPool, err2 := native_token_pool.NewNativeTokenPool(poolAddress, client)
		require.NoError(t, err2)
		tokenPools = append(tokenPools, tokenPool)
	}

	afn, err := afn_contract.NewAFNContract(config.Afn, client)
	require.NoError(t, err)
	blobVerifier, err := blob_verifier.NewBlobVerifier(config.BlobVerifier, client)
	require.NoError(t, err)
	offRamp, err := any_2_evm_toll_offramp.NewEVM2EVMTollOffRamp(config.OffRamp, client)
	require.NoError(t, err)
	messageReceiver, err := simple_message_receiver.NewSimpleMessageReceiver(config.MessageReceiver, client)
	require.NoError(t, err)
	receiverDapp, err := receiver_dapp.NewReceiverDapp(config.ReceiverDapp, client)
	require.NoError(t, err)
	offRampRouter, err := any_2_evm_toll_offramp_router.NewAny2EVMTollOffRampRouter(config.OffRampRouter, client)
	require.NoError(t, err)

	return DestClient{
		Client: Client{
			Client:           client,
			ChainId:          config.ChainId,
			LinkTokenAddress: config.LinkToken,
			LinkToken:        LinkToken,
			TokenPools:       tokenPools,
			Afn:              afn,
			logger:           logger.TestLogger(t).Named(helpers.ChainName(config.ChainId.Int64())),
			t:                t,
		},
		BlobVerifier:    blobVerifier,
		OffRampRouter:   offRampRouter,
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
func NewCcipClient(t *testing.T, sourceConfig EvmChainConfig, destConfig EvmChainConfig, ownerKey string, seedKey string) CCIPClient {
	source := NewSourceClient(t, sourceConfig)
	source.SetOwnerAndUsers(t, ownerKey, seedKey, sourceConfig.GasSettings)
	dest := NewDestinationClient(t, destConfig)
	dest.SetOwnerAndUsers(t, ownerKey, seedKey, destConfig.GasSettings)

	return CCIPClient{
		Source: source,
		Dest:   dest,
	}
}

func GetSetupChain(t *testing.T, ownerPrivateKey string, chain EvmChainConfig) *EvmChainConfig {
	chain.SetupChain(t, ownerPrivateKey)
	return &chain
}

func (chain *EvmChainConfig) SetupChain(t *testing.T, ownerPrivateKey string) {
	chain.Owner = GetOwner(t, ownerPrivateKey, chain.ChainId, chain.GasSettings)
	chain.Client = GetClient(t, chain.EthUrl)
	chain.Logger = logger.TestLogger(t).Named(helpers.ChainName(chain.ChainId.Int64()))

	require.Equal(t, len(chain.BridgeTokens), len(chain.TokenPools))
	chain.Logger.Info("Completed chain setup")
}

// GetOwner sets the owner user credentials and ensures a GasTipCap is set for the resulting user.
func GetOwner(t *testing.T, ownerPrivateKey string, chainId *big.Int, gasSettings EVMGasSettings) *bind.TransactOpts {
	ownerKey, err := crypto.HexToECDSA(ownerPrivateKey)
	require.NoError(t, err)
	user, err := bind.NewKeyedTransactorWithChainID(ownerKey, chainId)
	require.NoError(t, err)
	fmt.Println("--- Owner address ")
	fmt.Println(user.From.Hex())
	SetGasFees(user, gasSettings)

	return user
}

// GetClient dials a given EVM client url and returns the resulting client.
func GetClient(t *testing.T, ethUrl string) *ethclient.Client {
	client, err := ethclient.Dial(ethUrl)
	require.NoError(t, err)
	return client
}

// SetOwnerAndUsers sets the owner and 10 users on a given client. It also set the proper
// gas parameters on these users.
func (client *Client) SetOwnerAndUsers(t *testing.T, ownerPrivateKey string, seedKey string, gasSettings EVMGasSettings) {
	client.Owner = GetOwner(t, ownerPrivateKey, client.ChainId, gasSettings)

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
		SetGasFees(user, gasSettings)
		users = append(users, user)
		fmt.Println(user.From.Hex())
	}
	fmt.Println("---")

	client.Users = users
}

func (client *Client) ApproveLinkFrom(t *testing.T, user *bind.TransactOpts, approvedFor common.Address, amount *big.Int) {
	client.logger.Warnf("Approving %d link for %s", amount.Int64(), approvedFor.Hex())
	tx, err := client.LinkToken.Approve(user, approvedFor, amount)
	require.NoError(t, err)

	WaitForMined(client.t, client.logger, client.Client, tx.Hash(), true)
	client.logger.Warnf("Link approved %s", helpers.ExplorerLink(client.ChainId.Int64(), tx.Hash()))
}

func (client *Client) ApproveLink(t *testing.T, approvedFor common.Address, amount *big.Int) {
	client.ApproveLinkFrom(t, client.Owner, approvedFor, amount)
}

func (client CCIPClient) SendMessage(t *testing.T) {
	// ABI encoded message
	bytes, err := hex.DecodeString("00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000005626c616e6b000000000000000000000000000000000000000000000000000000")
	require.NoError(t, err)

	msg := evm_2_any_toll_onramp_router.CCIPEVM2AnyTollMessage{
		Receiver:       client.Dest.MessageReceiver.Address(),
		Data:           bytes,
		Tokens:         []common.Address{client.Source.LinkTokenAddress},
		Amounts:        []*big.Int{big.NewInt(1)},
		FeeToken:       client.Source.LinkTokenAddress,
		FeeTokenAmount: DefaultGasTipFee,
		GasLimit:       big.NewInt(0),
	}

	tx, err := client.Source.OnRampRouter.CcipSend(client.Source.Owner, client.Dest.ChainId, msg)
	require.NoError(t, err)
	WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
}

func (client CCIPClient) DonExecutionHappyPath(t *testing.T) {
	client.Source.logger.Infof("Starting cross chain tx with DON execution")
	sourceFeeAmount, err := client.Source.OnRamp.GetRequiredFee(&bind.CallOpts{Pending: false}, client.Source.LinkTokenAddress)
	require.NoError(t, err)
	client.Source.logger.Infof("Required onRamp fee is set to %d", sourceFeeAmount)

	destFeeAmount := big.NewInt(5e17)

	tokenAmount := big.NewInt(100)
	totalAmount := (&big.Int{}).Add(sourceFeeAmount, tokenAmount)
	totalAmount = (&big.Int{}).Add(totalAmount, destFeeAmount)
	client.Source.ApproveLink(t, client.Source.OnRampRouter.Address(), totalAmount)
	client.SetBlobVerifierConfig(t)
	client.Source.logger.Infof("Approved %d tokens for sending amount + %d for source fee + %d for destination fee. Total %d approved", tokenAmount, sourceFeeAmount, destFeeAmount, totalAmount)

	DestBlockNum := GetCurrentBlockNumber(client.Dest.Client.Client)
	crossChainRequest := client.SendToOnrampWithExecution(t, client.Source, client.Source.Owner, client.Dest.MessageReceiver.Address(), tokenAmount, (&big.Int{}).Add(sourceFeeAmount, destFeeAmount))
	client.Source.logger.Infof("Don executed tx submitted with sequence number: %d", crossChainRequest.Message.SequenceNumber)
	client.Source.logger.Infof("Waiting for Destination funds transfer...")

	events := make(chan *any_2_evm_toll_offramp.EVM2EVMTollOffRampExecutionStateChanged)
	sub, err := client.Dest.OffRamp.WatchExecutionStateChanged(
		&bind.WatchOpts{
			Context: context.Background(),
			Start:   &DestBlockNum,
		},
		events,
		[]uint64{crossChainRequest.Message.SequenceNumber})
	require.NoError(t, err)
	defer sub.Unsubscribe()

	select {
	case event := <-events:
		client.Dest.logger.Infof("found Destination execution in transaction %s", event.Raw.TxHash.Hex())
		return
	case err := <-sub.Err():
		panic(err)
	}
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

func (client CCIPClient) CrossChainSendPausedOnrampShouldFail(t *testing.T) {
	client.PauseOnramp()
	amount := big.NewInt(100)
	client.Source.ApproveLink(t, client.Source.SenderDapp.Address(), amount)
	client.Source.Owner.GasLimit = 1e6
	tx, err := client.Source.SenderDapp.SendTokens(client.Source.Owner, client.Dest.Owner.From, []common.Address{client.Source.LinkTokenAddress}, []*big.Int{amount})
	require.NoError(t, err)
	WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), false)
}

//func (client CCIPClient) CrossChainSendPausedOfframpShouldFail(t *testing.T) {
//	client.PauseBlobVerifier()
//	ctx := context.Background()
//	offrampBlockNumber := GetCurrentBlockNumber(client.Dest.Client.Client)
//
//	amount, _ := new(big.Int).SetString("10", 10)
//	client.Source.ApproveLink(t, client.Source.SenderDapp.Address(), amount)
//	onrampRequest := client.SendToDappWithExecution(client.Source, client.Source.Owner, client.Dest.Owner.From, amount, common.HexToAddress("0x0000000000000000000000000000000000000000"))
//
//	client.Dest.logger.Info("Waiting for report...")
//	_, err := client.GetReportForSequenceNumber(ctx, onrampRequest.Message.SequenceNumber, offrampBlockNumber)
//	if err.Error() == "No report found within the given time" {
//		client.Dest.logger.Info("Success, no oracle report sent to paused offramp.")
//	} else {
//		panic("report found despite paused offramp")
//	}
//}

func (client CCIPClient) NotEnoughFundsInBucketShouldFail(t *testing.T) {
	amount := big.NewInt(2e18) // 2 LINK, bucket size is 1 LINK
	client.Source.ApproveLink(t, client.Source.SenderDapp.Address(), amount)
	client.Source.Owner.GasLimit = 1e6
	tx, err := client.Source.SenderDapp.SendTokens(client.Source.Owner, client.Dest.Owner.From, []common.Address{client.Source.LinkTokenAddress}, []*big.Int{amount})
	require.NoError(t, err)
	WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), false)
}

//func (client CCIPClient) ExternalExecutionSubmitOfframpTwiceShouldFail(t *testing.T) {
//	ctx := context.Background()
//	offrampBlockNumber := GetCurrentBlockNumber(client.Dest.Client.Client)
//	onrampBlockNumber := GetCurrentBlockNumber(client.Source.Client.Client)
//
//	amount, _ := new(big.Int).SetString("10", 10)
//	client.Source.ApproveLink(t, client.Source.SenderDapp.Address(), amount)
//
//	onrampRequest := client.SendToDappWithExecution(client.Source, client.Source.Owner, client.Dest.Owner.From, amount, common.HexToAddress("0x0000000000000000000000000000000000000000"))
//
//	// Gets the report that our transaction is included in
//	client.Dest.logger.Info("Getting report")
//	report, err := client.GetReportForSequenceNumber(ctx, onrampRequest.Message.SequenceNumber, offrampBlockNumber)
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
//	client.Dest.logger.Info("Executing first offramp TX - should succeed")
//	tx, err := client.ExecuteOffRampTransaction(t, proof, onrampRequest.Raw.Data)
//	require.NoError(t, err)
//	WaitForMined(t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
//
//	// Execute the transaction on the offramp
//	client.Dest.logger.Info("Executing second offramp TX - should fail")
//	client.Dest.Owner.GasLimit = 1e6
//	tx, err = client.ExecuteOffRampTransaction(t, proof, onrampRequest.Raw.Data)
//	require.NoError(t, err)
//	WaitForMined(t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), false)
//}

// ScalingAndBatching should scale so that we see batching on the nodes
func (client CCIPClient) ScalingAndBatching(t *testing.T) {
	amount := big.NewInt(10)
	toAddress := common.HexToAddress("0x57359120D900fab8cE74edC2c9959b21660d3887")

	var wg sync.WaitGroup
	for _, user := range client.Source.Users {
		wg.Add(1)
		go func(user *bind.TransactOpts) {
			defer wg.Done()
			client.Source.ApproveLinkFrom(t, user, client.Source.SenderDapp.Address(), amount)
			crossChainRequest := client.SendToDappWithExecution(client.Source, user, toAddress, amount)
			client.Source.logger.Info("Don executed tx submitted with sequence number: ", crossChainRequest.Message.SequenceNumber)
		}(user)
	}
	wg.Wait()
	client.Source.logger.Info("Sent 10 txs to onramp.")
}

//func (client CCIPClient) ExecuteOffRampTransaction(t *testing.T, proof merklemulti.Proof[[32]byte], encodedMessage []byte) (*types.Transaction, error) {
//	decodedMsg, err := ccip.DecodeCCIPMessage(encodedMessage)
//	require.NoError(t, err)
//	_, err = ccip.MakeCCIPMsgArgs().PackValues([]interface{}{*decodedMsg})
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
//	tx, err := client.Dest.BlobVerifier.ExecuteTransaction(client.Dest.Owner, report, false)
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
//	report blob_verifier.CCIPRelayReport,
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

//// GetReportForSequenceNumber return the offramp.CCIPRelayReport for a given ccip requests sequence number.
//func (client CCIPClient) GetReportForSequenceNumber(ctx context.Context, sequenceNumber uint64, minBlockNumber uint64) (blob_verifier.CCIPRelayReport, error) {
//	client.Dest.logger.Infof("Looking for sequenceNumber %d", sequenceNumber)
//	report, err := client.Dest.OffRamp.GetLastReport(&bind.CallOpts{Context: ctx, Pending: false})
//	if err != nil {
//		return blob_verifier.CCIPRelayReport{}, err
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
//			report, err = client.Dest.BlobVerifier.GetLastReport(&bind.CallOpts{Context: ctx, Pending: false})
//			if err != nil {
//				return blob_verifier.CCIPRelayReport{}, err
//			}
//			client.Dest.logger.Infof("Last report found for range %d-%d", report.MinSequenceNumber, report.MaxSequenceNumber)
//			if sequenceNumber >= report.MinSequenceNumber && sequenceNumber <= report.MaxSequenceNumber {
//				return report, nil
//			}
//			time.Sleep(RetryTiming)
//		}
//		return blob_verifier.CCIPRelayReport{}, errors.New("No report found within the given timeout")
//	}
//
//	// it is in a past report, start looking at the earliest block number possible, the one
//	// before we started the entire transaction on the onramp.
//	reports, err := client.Dest.BlobVerifier.FilterReportAccepted(&bind.FilterOpts{
//		Start:   minBlockNumber,
//		End:     nil,
//		Context: ctx,
//	})
//	if err != nil {
//		return blob_verifier.CCIPRelayReport{}, err
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
//	return blob_verifier.CCIPRelayReport{}, errors.New("No report found for given sequence number")
//}

func (client CCIPClient) SetBlobVerifierConfig(t *testing.T) {
	config := blob_verifier.BlobVerifierInterfaceBlobVerifierConfig{
		OnRamps:          []common.Address{client.Source.OnRamp.Address()},
		MinSeqNrByOnRamp: []uint64{3},
	}
	tx, err := client.Dest.BlobVerifier.SetConfig(client.Dest.Owner, config)
	require.NoError(t, err)
	WaitForMined(t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
}

func GetCurrentBlockNumber(chain *ethclient.Client) uint64 {
	blockNumber, err := chain.BlockNumber(context.Background())
	helpers.PanicErr(err)
	return blockNumber
}

func (client CCIPClient) ValidateMerkleRoot(
	t *testing.T,
	request *evm_2_evm_toll_onramp.EVM2EVMTollOnRampCCIPSendRequested,
	reportRequests []*evm_2_evm_toll_onramp.EVM2EVMTollOnRampCCIPSendRequested,
	report blob_verifier.CCIPRelayReport,
) merklemulti.Proof[[32]byte] {
	mctx := merklemulti.NewKeccakCtx()
	var leafHashes [][32]byte
	for _, req := range reportRequests {
		leafHashes = append(leafHashes, mctx.HashLeaf(req.Raw.Data))
	}

	tree := merklemulti.NewTree(mctx, leafHashes)
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

	exists, err := client.Dest.BlobVerifier.GetMerkleRoot(nil, tree.Root())
	require.NoError(t, err)
	if exists.Uint64() < 1 {
		panic("Path is not present in the offramp")
	}
	index := request.Message.SequenceNumber - report.Intervals[rootIndex].Min
	client.Dest.logger.Info("index is ", index)
	return tree.Prove([]int{int(index)})
}

func (client CCIPClient) TryGetTokensFromPausedPool() {
	client.PauseOnrampPool()

	paused, err := client.Source.TokenPools[0].Paused(nil)
	helpers.PanicErr(err)
	if !paused {
		helpers.PanicErr(errors.New("Should be paused"))
	}

	client.Source.Owner.GasLimit = 2e6
	tx, err := client.Source.TokenPools[0].LockOrBurn(client.Source.Owner, big.NewInt(1000))
	helpers.PanicErr(err)
	WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), false)
}

// SendToDappWithExecution executes a cross chain transactions using the sender dapp interface.
func (client CCIPClient) SendToDappWithExecution(source SourceClient, from *bind.TransactOpts, toAddress common.Address, amount *big.Int) *evm_2_evm_toll_onramp.EVM2EVMTollOnRampCCIPSendRequested {
	SourceBlockNumber := GetCurrentBlockNumber(source.Client.Client)

	tx, err := source.SenderDapp.SendTokens(from, toAddress, []common.Address{client.Source.LinkTokenAddress}, []*big.Int{amount})
	helpers.PanicErr(err)
	source.logger.Infof("Send tokens tx %s", helpers.ExplorerLink(source.ChainId.Int64(), tx.Hash()))
	WaitForMined(source.t, source.logger, source.Client.Client, tx.Hash(), true)

	return WaitForCrossChainSendRequest(source, SourceBlockNumber, tx.Hash())
}

// SendToOnrampWithExecution executes a cross chain transactions using the onramp interface.
func (client CCIPClient) SendToOnrampWithExecution(t *testing.T, source SourceClient, from *bind.TransactOpts, toAddress common.Address, amount *big.Int, feeTokenAmount *big.Int) *evm_2_evm_toll_onramp.EVM2EVMTollOnRampCCIPSendRequested {
	SourceBlockNumber := GetCurrentBlockNumber(source.Client.Client)
	payload := evm_2_any_toll_onramp_router.CCIPEVM2AnyTollMessage{
		Tokens:         []common.Address{},
		Amounts:        []*big.Int{},
		Receiver:       toAddress,
		Data:           []byte{},
		FeeToken:       source.LinkTokenAddress,
		FeeTokenAmount: feeTokenAmount,
		GasLimit:       big.NewInt(0),
	}
	source.logger.Infof("Send tx with payload %+v", payload)

	tx, err := source.OnRampRouter.CcipSend(from, client.Dest.ChainId, payload)
	if err != nil {
		t.Log(err.Error())
		printRevertReason(err, evm_2_any_toll_onramp_router.EVM2AnyTollOnRampRouterABI)
	}
	helpers.PanicErr(err)
	source.logger.Infof("Send tokens tx %s", helpers.ExplorerLink(source.ChainId.Int64(), tx.Hash()))
	return WaitForCrossChainSendRequest(source, SourceBlockNumber, tx.Hash())
}

func (client CCIPClient) manuallyExecuteTx(event evm_2_evm_toll_onramp.CCIPEVM2EVMTollEvent) {
	packedPayload, err := ccip.MakeCCIPMsgArgs().PackValues([]interface{}{event})
	helpers.PanicErr(err)
	report := any_2_evm_toll_offramp.CCIPExecutionReport{
		SequenceNumbers:          []uint64{event.SequenceNumber},
		TokenPerFeeCoinAddresses: []common.Address{client.Source.LinkTokenAddress},
		TokenPerFeeCoin:          []*big.Int{big.NewInt(1)},
		EncodedMessages:          [][]byte{packedPayload},
		InnerProofs:              [][32]byte{},
		InnerProofFlagBits:       big.NewInt(0),
		OuterProofs:              [][32]byte{},
		OuterProofFlagBits:       big.NewInt(0),
	}

	tx, err := client.Dest.OffRamp.Execute(client.Dest.Owner, report, true)
	WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)

	if err != nil {
		printRevertReason(err, evm_2_any_toll_onramp_router.EVM2AnyTollOnRampRouterABI)
	}

	helpers.PanicErr(err)
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
func WaitForCrossChainSendRequest(source SourceClient, fromBlockNum uint64, txhash common.Hash) *evm_2_evm_toll_onramp.EVM2EVMTollOnRampCCIPSendRequested {
	filter := bind.FilterOpts{Start: fromBlockNum}
	for {
		iterator, err := source.OnRamp.FilterCCIPSendRequested(&filter)
		helpers.PanicErr(err)
		for iterator.Next() {
			if iterator.Event.Raw.TxHash.Hex() == txhash.Hex() {
				source.logger.Infof("Cross chain send event found in tx: %s ", helpers.ExplorerLink(source.ChainId.Int64(), txhash))
				return iterator.Event
			}
		}
		time.Sleep(RetryTiming)
	}
}

func (client CCIPClient) PauseOfframpPool() {
	paused, err := client.Dest.TokenPools[0].Paused(nil)
	helpers.PanicErr(err)
	if paused {
		return
	}
	client.Dest.logger.Info("pausing offramp pool...")
	tx, err := client.Dest.TokenPools[0].Pause(client.Dest.Owner)
	helpers.PanicErr(err)
	client.Dest.logger.Info("Offramp pool paused, tx hash: %s", tx.Hash())
	WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
}

func (client CCIPClient) PauseOnrampPool() {
	paused, err := client.Source.TokenPools[0].Paused(nil)
	helpers.PanicErr(err)
	if paused {
		return
	}
	client.Source.logger.Info("pausing onramp pool...")
	tx, err := client.Source.TokenPools[0].Pause(client.Source.Owner)
	helpers.PanicErr(err)
	client.Source.logger.Info("Onramp pool paused, tx hash: %s", tx.Hash())
	WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
}

func (client CCIPClient) UnpauseOfframpPool() {
	paused, err := client.Dest.TokenPools[0].Paused(nil)
	helpers.PanicErr(err)
	if !paused {
		return
	}
	client.Dest.logger.Info("unpausing offramp pool...")
	tx, err := client.Dest.TokenPools[0].Unpause(client.Dest.Owner)
	helpers.PanicErr(err)
	client.Dest.logger.Info("Offramp pool unpaused, tx hash: %s", tx.Hash())
	WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
}

func (client CCIPClient) UnpauseOnrampPool() {
	paused, err := client.Source.TokenPools[0].Paused(nil)
	helpers.PanicErr(err)
	if !paused {
		return
	}
	client.Source.logger.Info("unpausing onramp pool...")
	tx, err := client.Source.TokenPools[0].Unpause(client.Source.Owner)
	helpers.PanicErr(err)
	client.Source.logger.Info("Onramp pool unpaused, tx hash: %s", tx.Hash())
	WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
}

func (client CCIPClient) PauseOnramp() {
	paused, err := client.Source.OnRamp.Paused(nil)
	helpers.PanicErr(err)
	if paused {
		return
	}
	client.Source.logger.Info("pausing onramp...")
	tx, err := client.Source.OnRamp.Pause(client.Source.Owner)
	helpers.PanicErr(err)
	client.Source.logger.Info("Onramp paused, tx hash: %s", tx.Hash())
	WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
}

func (client CCIPClient) PauseBlobVerifier() {
	paused, err := client.Dest.BlobVerifier.Paused(nil)
	helpers.PanicErr(err)
	if paused {
		return
	}
	client.Dest.logger.Info("pausing offramp...")
	tx, err := client.Dest.BlobVerifier.Pause(client.Dest.Owner)
	helpers.PanicErr(err)
	client.Dest.logger.Info("Offramp paused, tx hash: %s", tx.Hash())
	WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
}

func (client CCIPClient) UnpauseOnramp() {
	paused, err := client.Source.OnRamp.Paused(nil)
	helpers.PanicErr(err)
	if !paused {
		return
	}
	client.Source.logger.Info("unpausing onramp...")
	tx, err := client.Source.OnRamp.Unpause(client.Source.Owner)
	helpers.PanicErr(err)
	client.Source.logger.Info("Onramp unpaused, tx hash: %s", tx.Hash())
	WaitForMined(client.Source.t, client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
}

func (client CCIPClient) UnpauseBlobVerifier() {
	paused, err := client.Dest.BlobVerifier.Paused(nil)
	helpers.PanicErr(err)
	if !paused {
		return
	}
	client.Dest.logger.Info("unpausing offramp...")
	tx, err := client.Dest.BlobVerifier.Unpause(client.Dest.Owner)
	helpers.PanicErr(err)
	client.Dest.logger.Info("Offramp unpaused, tx hash: %s", tx.Hash())
	WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
}

func (client CCIPClient) UnpauseAll() {
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer wg.Done()
		client.UnpauseOnramp()
	}()
	go func() {
		defer wg.Done()
		client.UnpauseBlobVerifier()
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

func (client CCIPClient) SetOCRConfig() {
	verifierOCRConfig, err := client.Dest.BlobVerifier.LatestConfigDetails(&bind.CallOpts{})
	helpers.PanicErr(err)
	if verifierOCRConfig.BlockNumber != 0 {
		client.Dest.logger.Infof("BlobVerifier OCR config already found: %+v", verifierOCRConfig.ConfigDigest)
		client.Dest.logger.Infof("The new config will overwrite the current one.")
	}

	rampOCRConfig, err := client.Dest.OffRamp.LatestConfigDetails(&bind.CallOpts{})
	helpers.PanicErr(err)
	if rampOCRConfig.BlockNumber != 0 {
		client.Dest.logger.Infof("OffRamp OCR config already found: %+v", rampOCRConfig.ConfigDigest)
		client.Dest.logger.Infof("The new config will overwrite the current one.")
	}

	ccipConfig, err := ccip.OffchainConfig{
		SourceIncomingConfirmations: 1,
		DestIncomingConfirmations:   1,
	}.Encode()
	helpers.PanicErr(err)

	signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, err := confighelper2.ContractSetConfigArgsForTests(
		60*time.Second, // deltaProgress
		1*time.Second,  // deltaResend
		20*time.Second, // deltaRound
		2*time.Second,  // deltaGrace
		30*time.Second, // deltaStage
		3,
		[]int{1, 2, 3, 4}, // Transmission schedule: 1 oracle in first deltaStage, 2 in the second and so on.
		getOraclesForChain(client.Dest.ChainId.Int64()),
		ccipConfig,
		1*time.Second,
		10*time.Second,
		20*time.Second,
		10*time.Second,
		10*time.Second,
		1, // faults
		nil,
	)
	helpers.PanicErr(err)

	signerAddresses, err := ocrcommon.OnchainPublicKeyToAddress(signers)
	helpers.PanicErr(err)
	transmitterAddresses, err := ocrcommon.AccountToAddress(transmitters)
	helpers.PanicErr(err)

	tx, err := client.Dest.BlobVerifier.SetConfig0(
		client.Dest.Owner,
		signerAddresses,
		transmitterAddresses,
		f,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)
	helpers.PanicErr(err)
	WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
	client.Dest.logger.Infof("Config set on blob verifier %s", helpers.ExplorerLink(client.Dest.ChainId.Int64(), tx.Hash()))

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
	WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
	client.Dest.logger.Infof("Config set on offramp %s", helpers.ExplorerLink(client.Dest.ChainId.Int64(), tx.Hash()))
}

func (client CCIPClient) AcceptOwnership(t *testing.T) {
	tx, err := client.Dest.BlobVerifier.AcceptOwnership(client.Dest.Owner)
	require.NoError(t, err)
	WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)

	tx, err = client.Dest.OffRamp.AcceptOwnership(client.Dest.Owner)
	require.NoError(t, err)
	WaitForMined(client.Dest.t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
}
