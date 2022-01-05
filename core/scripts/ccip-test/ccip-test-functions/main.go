package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/services/ccip/abihelpers"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2/confighelper"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_offramp"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/single_token_onramp"
	ccipshared "github.com/smartcontractkit/chainlink/core/scripts/ccip-test/ccip-shared"
	"github.com/smartcontractkit/chainlink/core/services/ccip"
)

type ccipClient ccipshared.CcipClient

func main() {
	// This key is used to deploy all contracts on both source and Dest chains
	k := os.Getenv("OWNER_KEY")
	if k == "" {
		panic("must set owner key")
	}

	// The seed key is used to generate 10 keys from a single key by changing the
	// first character of the given seed with the digits 0-9
	seedKey := os.Getenv("SEED_KEY")
	if seedKey == "" {
		panic("must set seed key")
	}

	// Configures a client to run tests with using the network defaults and given keys.
	// After updating any contracts be sure to update the network defaults to reflect
	// those changes.
	client := ccipClient(ccipshared.NewCcipClient(
		// Source chain
		ccipshared.Kovan.SetOwnerAndUsers(k, seedKey),
		// Dest chain
		ccipshared.Rinkeby.SetOwnerAndUsers(k, seedKey)))

	client.Source.Client.AssureHealth()
	client.Dest.Client.AssureHealth()
	client.unpauseAll()

	// Set the config to the message executor and the offramp
	//client.setConfig()

	// Cross chain request with the client manually proving and executing the transaction
	//client.externalExecutionHappyPath()

	// Executing the same request twice should fail
	//client.externalExecutionSubmitOfframpTwiceShouldFail()

	// Cross chain request with DON execution
	//client.donExecutionHappyPath()

	// Submit 10 txs. This should result in the txs being batched together
	//client.scalingAndBatching()

	// Should not be able to send funds greater than the amount in the bucket
	//client.notEnoughFundsInBucketShouldFail()

	//client.tryGetTokensFromPausedPool()

	client.crossChainSendPausedOfframpShouldFail()

	//client.crossChainSendPausedOnrampShouldFail()
}

func (client ccipClient) sendMessage() {
	// ABI encoded message
	bytes, _ := hex.DecodeString("00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000005626c616e6b000000000000000000000000000000000000000000000000000000")

	msg := single_token_onramp.CCIPMessagePayload{
		Receiver: client.Dest.SimpleMessageReceiver.Address(),
		Data:     bytes,
		Tokens:   []common.Address{client.Source.LinkToken.Address()},
		Amounts:  []*big.Int{big.NewInt(1)},
		Options:  []byte{},
		Executor: client.Dest.MessageExecutor.Address(),
	}

	client.Source.ApproveLink(client.Source.LockUnlockPool.Address(), big.NewInt(1))
	tx, err := client.Source.SingleTokenOnramp.RequestCrossChainSend(client.Source.Owner, msg)
	ccipshared.PanicErr(err)
	ccipshared.WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), true)
}

func (client ccipClient) donExecutionHappyPath() {
	amount := big.NewInt(100)
	client.Source.ApproveLink(client.Source.SingleTokenSender.Address(), amount)
	DestBlockNum := getCurrentBlockNumber(client.Dest.Client.Client)
	crossChainRequest := client.sendToOnrampWithExecution(client.Source, client.Source.Owner, client.Dest.Owner.From, amount, client.Dest.MessageExecutor.Address())
	fmt.Println("Don executed tx submitted with sequence number: ", crossChainRequest.Message.SequenceNumber.Int64())
	fmt.Println("Waiting for Destination funds transfer...")

	events := make(chan *single_token_offramp.SingleTokenOffRampCrossChainMessageExecuted)
	sub, err := client.Dest.SingleTokenOfframp.WatchCrossChainMessageExecuted(
		&bind.WatchOpts{
			Context: context.Background(),
			Start:   &DestBlockNum,
		},
		events,
		[]*big.Int{crossChainRequest.Message.SequenceNumber})
	ccipshared.PanicErr(err)
	defer sub.Unsubscribe()

	select {
	case event := <-events:
		fmt.Printf("found Destination execution in transaction: %s\n", event.Raw.TxHash.Hex())
		return
	case err := <-sub.Err():
		panic(err)
	}
}

func (client ccipClient) externalExecutionHappyPath() {
	ctx := context.Background()
	offrampBlockNumber := getCurrentBlockNumber(client.Dest.Client.Client)
	onrampBlockNumber := getCurrentBlockNumber(client.Source.Client.Client)

	amount, _ := new(big.Int).SetString("10", 10)
	client.Source.ApproveLink(client.Source.SingleTokenSender.Address(), amount)

	onrampRequest := client.sendToOnrampWithExecution(client.Source, client.Source.Owner, client.Dest.Owner.From, amount, common.HexToAddress("0x0000000000000000000000000000000000000000"))
	sequenceNumber := onrampRequest.Message.SequenceNumber.Int64()

	// Gets the report that our transaction is included in
	fmt.Println("Getting report")
	report := client.getReportForSequenceNumber(ctx, sequenceNumber, offrampBlockNumber)

	// Get all requests included in the given report
	fmt.Println("Getting recent cross chain requests")
	requests := client.getCrossChainSendRequestsForRange(ctx, report, onrampBlockNumber)

	// Generate the proof
	fmt.Println("Generating proof")
	proof := client.validateMerkleRoot(onrampRequest, requests, report)

	// Execute the transaction on the offramp
	client.Dest.Owner.GasLimit = 2e9
	fmt.Println("Executing offramp TX")
	tx, err := client.executeOfframpTransaction(proof, onrampRequest.Raw.Data)
	ccipshared.PanicErr(err)
	ccipshared.WaitForMined(ctx, client.Dest.Client.Client, tx.Hash(), true)
}

func (client ccipClient) crossChainSendPausedOnrampShouldFail() {
	client.pauseOnramp()
	amount := big.NewInt(100)
	client.Source.ApproveLink(client.Source.SingleTokenSender.Address(), amount)
	client.Source.Owner.GasLimit = 1e6
	tx, err := client.Source.SingleTokenSender.SendTokens(client.Source.Owner, client.Dest.Owner.From, amount, client.Dest.MessageExecutor.Address())
	ccipshared.PanicErr(err)
	ccipshared.WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), false)
}

func (client ccipClient) crossChainSendPausedOfframpShouldFail() {
	client.pauseOfframp()
	ctx := context.Background()
	offrampBlockNumber := getCurrentBlockNumber(client.Dest.Client.Client)

	amount, _ := new(big.Int).SetString("10", 10)
	client.Source.ApproveLink(client.Source.SingleTokenSender.Address(), amount)

	onrampRequest := client.sendToOnrampWithExecution(client.Source, client.Source.Owner, client.Dest.Owner.From, amount, common.HexToAddress("0x0000000000000000000000000000000000000000"))
	sequenceNumber := onrampRequest.Message.SequenceNumber.Int64()

	fmt.Println("Waiting for report...")
	result := make(chan single_token_offramp.CCIPRelayReport)
	go func() {
		result <- client.getReportForSequenceNumber(ctx, sequenceNumber, offrampBlockNumber)
	}()

	select {
	case r := <-result:
		panic(fmt.Errorf("report found despite paused offramp: %+v", r))
	case <-time.After(time.Minute):
		fmt.Println("Success, no oracle report sent to paused offramp.")
	}
}

func (client ccipClient) notEnoughFundsInBucketShouldFail() {
	amount := big.NewInt(2e18) // 2 LINK, bucket size is 1 LINK
	client.Source.ApproveLink(client.Source.SingleTokenSender.Address(), amount)
	client.Source.Owner.GasLimit = 1e6
	tx, err := client.Source.SingleTokenSender.SendTokens(client.Source.Owner, client.Dest.Owner.From, amount, client.Dest.MessageExecutor.Address())
	ccipshared.PanicErr(err)
	ccipshared.WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), false)
}

func (client ccipClient) externalExecutionSubmitOfframpTwiceShouldFail() {
	ctx := context.Background()
	offrampBlockNumber := getCurrentBlockNumber(client.Dest.Client.Client)
	onrampBlockNumber := getCurrentBlockNumber(client.Source.Client.Client)

	amount, _ := new(big.Int).SetString("10", 10)
	client.Source.ApproveLink(client.Source.SingleTokenSender.Address(), amount)

	onrampRequest := client.sendToOnrampWithExecution(client.Source, client.Source.Owner, client.Dest.Owner.From, amount, common.HexToAddress("0x0000000000000000000000000000000000000000"))
	sequenceNumber := onrampRequest.Message.SequenceNumber.Int64()

	// Gets the report that our transaction is included in
	fmt.Println("Getting report")
	report := client.getReportForSequenceNumber(ctx, sequenceNumber, offrampBlockNumber)

	// Get all requests included in the given report
	fmt.Println("Getting recent cross chain requests")
	requests := client.getCrossChainSendRequestsForRange(ctx, report, onrampBlockNumber)

	// Generate the proof
	fmt.Println("Generating proof")
	proof := client.validateMerkleRoot(onrampRequest, requests, report)

	// Execute the transaction on the offramp
	fmt.Println("Executing first offramp TX - should succeed")
	tx, err := client.executeOfframpTransaction(proof, onrampRequest.Raw.Data)
	ccipshared.PanicErr(err)
	ccipshared.WaitForMined(ctx, client.Dest.Client.Client, tx.Hash(), true)

	// Execute the transaction on the offramp
	fmt.Println("Executing second offramp TX - should fail")
	client.Dest.Owner.GasLimit = 1e6
	tx, err = client.executeOfframpTransaction(proof, onrampRequest.Raw.Data)
	ccipshared.PanicErr(err)
	ccipshared.WaitForMined(ctx, client.Dest.Client.Client, tx.Hash(), false)
}

// Scale so that we see batching on the nodes
func (client ccipClient) scalingAndBatching() {
	amount := big.NewInt(10)
	toAddress := common.HexToAddress("0x57359120D900fab8cE74edC2c9959b21660d3887")

	var wg sync.WaitGroup
	for _, user := range client.Source.Users {
		wg.Add(1)
		go func(user *bind.TransactOpts) {
			defer wg.Done()
			client.Source.ApproveLinkFrom(user, client.Source.SingleTokenSender.Address(), amount)
			crossChainRequest := client.sendToOnrampWithExecution(client.Source, user, toAddress, amount, client.Dest.MessageExecutor.Address())
			fmt.Println("Don executed tx submitted with sequence number: ", crossChainRequest.Message.SequenceNumber.Int64())
		}(user)
	}
	wg.Wait()
	fmt.Println("Sent 10 txs to onramp.")
}

func (client ccipClient) executeOfframpTransaction(proof ccip.MerkleProof, encodedMessage []byte) (*types.Transaction, error) {
	decodedMsg, err := abihelpers.DecodeCCIPMessage(encodedMessage)
	ccipshared.PanicErr(err)
	_, err = abihelpers.MakeCCIPMsgArgs().PackValues([]interface{}{*decodedMsg})
	ccipshared.PanicErr(err)

	tx, err := client.Dest.SingleTokenOfframp.ExecuteTransaction(client.Dest.Owner, proof.PathForExecute(), *decodedMsg, proof.Index())
	return tx, errors.Wrap(err, "executing offramp tx")
}

func (client ccipClient) getCrossChainSendRequestsForRange(
	ctx context.Context,
	report single_token_offramp.CCIPRelayReport,
	onrampBlockNumber uint64) []*single_token_onramp.SingleTokenOnRampCrossChainSendRequested {
	// Get the other transactions in the proof, we look 1000 blocks back for transaction
	// should be fine? Needs fine-tuning after improved batching strategies are developed
	// in milestone 4
	reqsIterator, err := client.Source.SingleTokenOnramp.FilterCrossChainSendRequested(&bind.FilterOpts{
		Context: ctx,
		Start:   onrampBlockNumber - 1000,
	})
	ccipshared.PanicErr(err)

	var requests []*single_token_onramp.SingleTokenOnRampCrossChainSendRequested

	var minFound = report.MaxSequenceNumber.Int64()
	for reqsIterator.Next() {
		num := reqsIterator.Event.Message.SequenceNumber.Int64()
		if num < minFound {
			minFound = num
		}
		if num >= report.MinSequenceNumber.Int64() && num <= report.MaxSequenceNumber.Int64() {
			requests = append(requests, reqsIterator.Event)
		}
	}

	// TODO: Even if this check passes, we may not have fetched all necessary requests if
	// minFound == report.MinSequenceNumber
	if minFound > report.MinSequenceNumber.Int64() {
		ccipshared.PanicErr(errors.New("Not all cross chain requests found in the last 1000 blocks"))
	}

	return requests
}

func (client ccipClient) getReportForSequenceNumber(ctx context.Context, sequenceNumber int64, minBlockNumber uint64) single_token_offramp.CCIPRelayReport {
	report, err := client.Dest.SingleTokenOfframp.GetLastReport(&bind.CallOpts{
		Pending: false,
	})
	ccipshared.PanicErr(err)

	// our tx is in the latest report
	if sequenceNumber >= report.MinSequenceNumber.Int64() && sequenceNumber <= report.MaxSequenceNumber.Int64() {
		return report
	}
	// report isn't out yet, it will be in a future report
	if sequenceNumber > report.MaxSequenceNumber.Int64() {
		for {
			report, err = client.Dest.SingleTokenOfframp.GetLastReport(&bind.CallOpts{
				Pending: false,
				Context: ctx,
			})
			ccipshared.PanicErr(err)
			if sequenceNumber >= report.MinSequenceNumber.Int64() && sequenceNumber <= report.MaxSequenceNumber.Int64() {
				return report
			}
			time.Sleep(ccipshared.RetryTiming)
		}
	}

	// it is in a past report, start looking at the earliest block number possible, the one
	// before we started the entire transaction on the onramp.
	reports, err := client.Dest.SingleTokenOfframp.FilterReportAccepted(&bind.FilterOpts{
		Start:   minBlockNumber,
		End:     nil,
		Context: ctx,
	})
	ccipshared.PanicErr(err)

	for reports.Next() {
		report = reports.Event.Report
		if sequenceNumber >= report.MinSequenceNumber.Int64() && sequenceNumber <= report.MaxSequenceNumber.Int64() {
			return report
		}
	}

	// Somehow the transaction was not included in any report within blocks produced after
	// the transaction was initialized but the sequence number is lower than we are currently at
	ccipshared.PanicErr(errors.New("No report found"))
	return single_token_offramp.CCIPRelayReport{}
}

func getCurrentBlockNumber(chain *ethclient.Client) uint64 {
	blockNumber, err := chain.BlockNumber(context.Background())
	ccipshared.PanicErr(err)
	return blockNumber
}

func (client ccipClient) validateMerkleRoot(
	request *single_token_onramp.SingleTokenOnRampCrossChainSendRequested,
	reportRequests []*single_token_onramp.SingleTokenOnRampCrossChainSendRequested,
	report single_token_offramp.CCIPRelayReport,
) ccip.MerkleProof {
	var leaves [][]byte
	for _, req := range reportRequests {
		leaves = append(leaves, req.Raw.Data)
	}

	index := big.NewInt(0).Sub(request.Message.SequenceNumber, report.MinSequenceNumber)
	fmt.Println("index is", index)
	root, proof := ccip.GenerateMerkleProof(32, leaves, int(index.Int64()))
	if !reflect.DeepEqual(root[:], report.MerkleRoot[:]) {
		ccipshared.PanicErr(errors.New("Merkle root does not match the report"))
	}

	genRoot := ccip.GenerateMerkleRoot(leaves[int(index.Int64())], proof)
	if !reflect.DeepEqual(root[:], genRoot[:]) {
		ccipshared.PanicErr(errors.New("Root does not verify"))
	}

	exists, err := client.Dest.SingleTokenOfframp.GetMerkleRoot(nil, root)
	ccipshared.PanicErr(err)
	if exists.Uint64() < 1 {
		ccipshared.PanicErr(errors.New("Proof is not present in the offramp"))
	}
	return proof
}

func (client ccipClient) tryGetTokensFromPausedPool() {
	client.pauseOnrampPool()

	paused, err := client.Source.LockUnlockPool.Paused(nil)
	ccipshared.PanicErr(err)
	if !paused {
		ccipshared.PanicErr(errors.New("Should be paused"))
	}

	client.Source.Owner.GasLimit = 2e6
	tx, err := client.Source.LockUnlockPool.LockOrBurn(client.Source.Owner, client.Source.Owner.From, big.NewInt(1000))
	ccipshared.PanicErr(err)
	ccipshared.WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), false)
}

func (client ccipClient) sendToOnrampWithExecution(Source ccipshared.SourceClient, from *bind.TransactOpts, toAddress common.Address, amount *big.Int, executor common.Address) *single_token_onramp.SingleTokenOnRampCrossChainSendRequested {
	ctx := context.Background()
	SourceBlockNumber := getCurrentBlockNumber(Source.Client.Client)

	tx, err := Source.SingleTokenSender.SendTokens(from, toAddress, amount, executor)
	ccipshared.PanicErr(err)
	fmt.Println("send tokens hash ", tx.Hash())
	ccipshared.WaitForMined(ctx, Source.Client.Client, tx.Hash(), true)

	return waitForCrossChainSendRequest(Source, SourceBlockNumber, tx.Hash())
}

func waitForCrossChainSendRequest(Source ccipshared.SourceClient, fromBlockNum uint64, txhash common.Hash) *single_token_onramp.SingleTokenOnRampCrossChainSendRequested {
	filter := bind.FilterOpts{Start: fromBlockNum}
	for {
		iterator, err := Source.SingleTokenOnramp.FilterCrossChainSendRequested(&filter)
		ccipshared.PanicErr(err)
		for iterator.Next() {
			if iterator.Event.Raw.TxHash.Hex() == txhash.Hex() {
				fmt.Println("cross chain send event found in tx: ", txhash.Hex())
				return iterator.Event
			}
		}
		time.Sleep(ccipshared.RetryTiming)
	}
}

func (client ccipClient) pauseOfframpPool() {
	paused, err := client.Dest.LockUnlockPool.Paused(nil)
	ccipshared.PanicErr(err)
	if paused {
		return
	}
	fmt.Println("pausing offramp pool...")
	tx, err := client.Dest.LockUnlockPool.Pause(client.Dest.Owner)
	ccipshared.PanicErr(err)
	fmt.Println("Offramp pool paused, tx hash:", tx.Hash())
	ccipshared.WaitForMined(context.Background(), client.Dest.Client.Client, tx.Hash(), true)
}

func (client ccipClient) pauseOnrampPool() {
	paused, err := client.Source.LockUnlockPool.Paused(nil)
	ccipshared.PanicErr(err)
	if paused {
		return
	}
	fmt.Println("pausing onramp pool...")
	tx, err := client.Source.LockUnlockPool.Pause(client.Source.Owner)
	ccipshared.PanicErr(err)
	fmt.Println("Onramp pool paused, tx hash:", tx.Hash())
	ccipshared.WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), true)
}

func (client ccipClient) unpauseOfframpPool() {
	paused, err := client.Dest.LockUnlockPool.Paused(nil)
	ccipshared.PanicErr(err)
	if !paused {
		return
	}
	fmt.Println("unpausing offramp pool...")
	tx, err := client.Dest.LockUnlockPool.Unpause(client.Dest.Owner)
	ccipshared.PanicErr(err)
	fmt.Println("Offramp pool unpaused, tx hash:", tx.Hash())
	ccipshared.WaitForMined(context.Background(), client.Dest.Client.Client, tx.Hash(), true)
}

func (client ccipClient) unpauseOnrampPool() {
	paused, err := client.Source.LockUnlockPool.Paused(nil)
	ccipshared.PanicErr(err)
	if !paused {
		return
	}
	fmt.Println("unpausing onramp pool...")
	tx, err := client.Source.LockUnlockPool.Unpause(client.Source.Owner)
	ccipshared.PanicErr(err)
	fmt.Println("Onramp pool unpaused, tx hash:", tx.Hash())
	ccipshared.WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), true)
}

func (client ccipClient) pauseOnramp() {
	paused, err := client.Source.SingleTokenOnramp.Paused(nil)
	ccipshared.PanicErr(err)
	if paused {
		return
	}
	fmt.Println("pausing onramp...")
	tx, err := client.Source.SingleTokenOnramp.Pause(client.Source.Owner)
	ccipshared.PanicErr(err)
	fmt.Println("Onramp paused, tx hash:", tx.Hash())
	ccipshared.WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), true)
}

func (client ccipClient) pauseOfframp() {
	paused, err := client.Dest.SingleTokenOfframp.Paused(nil)
	ccipshared.PanicErr(err)
	if paused {
		return
	}
	fmt.Println("pausing offramp...")
	tx, err := client.Dest.SingleTokenOfframp.Pause(client.Dest.Owner)
	ccipshared.PanicErr(err)
	fmt.Println("Offramp paused, tx hash:", tx.Hash())
	ccipshared.WaitForMined(context.Background(), client.Dest.Client.Client, tx.Hash(), true)
}

func (client ccipClient) unpauseOnramp() {
	paused, err := client.Source.SingleTokenOnramp.Paused(nil)
	ccipshared.PanicErr(err)
	if !paused {
		return
	}
	fmt.Println("unpausing onramp...")
	tx, err := client.Source.SingleTokenOnramp.Unpause(client.Source.Owner)
	ccipshared.PanicErr(err)
	fmt.Println("Onramp unpaused, tx hash:", tx.Hash())
	ccipshared.WaitForMined(context.Background(), client.Source.Client.Client, tx.Hash(), true)
}

func (client ccipClient) unpauseOfframp() {
	paused, err := client.Dest.SingleTokenOfframp.Paused(nil)
	ccipshared.PanicErr(err)
	if !paused {
		return
	}
	fmt.Println("unpausing offramp...")
	tx, err := client.Dest.SingleTokenOfframp.Unpause(client.Dest.Owner)
	ccipshared.PanicErr(err)
	fmt.Println("Offramp unpaused, tx hash:", tx.Hash())
	ccipshared.WaitForMined(context.Background(), client.Dest.Client.Client, tx.Hash(), true)
}

func (client ccipClient) unpauseAll() {
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer wg.Done()
		client.unpauseOnramp()
	}()
	go func() {
		defer wg.Done()
		client.unpauseOfframp()
	}()
	go func() {
		defer wg.Done()
		client.unpauseOnrampPool()
	}()
	go func() {
		defer wg.Done()
		client.unpauseOfframpPool()
	}()
	wg.Wait()
}

func (client ccipClient) setConfig() {
	ccipConfig, err := ccip.OffchainConfig{
		SourceIncomingConfirmations: 0,
		DestIncomingConfirmations:   0,
	}.Encode()
	ccipshared.PanicErr(err)
	signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, err := confighelper2.ContractSetConfigArgs(
		60*time.Second, // deltaProgress
		1*time.Second,  // deltaResend
		20*time.Second, // deltaRound
		2*time.Second,  // deltaGrace
		30*time.Second, // deltaStage
		3,
		[]int{1, 2, 3, 4}, // Transmission schedule: 1 oracle in first deltaStage, 2 in the second and so on.
		ccipshared.Oracles,
		ccipConfig,
		1*time.Second,
		10*time.Second,
		20*time.Second,
		10*time.Second,
		10*time.Second,
		1, // faults
		nil,
	)
	ccipshared.PanicErr(err)

	ctx := context.Background()

	tx, err := client.Dest.SingleTokenOfframp.SetConfig(
		client.Dest.Owner,
		signers,
		transmitters,
		f,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)
	ccipshared.PanicErr(err)
	ccipshared.WaitForMined(ctx, client.Dest.Client.Client, tx.Hash(), true)
	fmt.Println("Config set on offramp. Tx hash:", tx.Hash().Hex())

	tx, err = client.Dest.MessageExecutor.SetConfig(
		client.Dest.Owner,
		signers,
		transmitters,
		f,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)
	ccipshared.PanicErr(err)
	ccipshared.WaitForMined(ctx, client.Dest.Client.Client, tx.Hash(), true)
	fmt.Println("Config set on message executor. Tx hash:", tx.Hash().Hex())
}
