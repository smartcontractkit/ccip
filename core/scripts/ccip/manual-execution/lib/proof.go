package lib

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	chainselectors "github.com/smartcontractkit/chain-selectors"

	"manual-execution/helpers"
)

type ClientData struct {
	SourceChain   *ethclient.Client
	SourceChainId *big.Int
	DestChain     *ethclient.Client
	DestChainId   *big.Int
}

type ProofData struct {
	SrcStartBlock   *big.Int
	DestStartBlock  uint64
	DestLatestBlock uint64
	SourceChainId   *big.Int
	DestChainId     *big.Int

	SeqData
}

type SeqData struct {
	OnRamp common.Address
	SeqNum uint64
	MsgID  [32]byte
}

func dial(ctx context.Context, url string) (client *ethclient.Client, id *big.Int, err error) {
	client, err = ethclient.Dial(url)
	if err != nil {
		err = fmt.Errorf("unable to dial node: %w", err)
		return
	}
	id, err = client.ChainID(ctx)
	if err != nil {
		err = fmt.Errorf("unable to lookup chain ID: %w", err)
		return
	}

	return
}

func GetClients(ctx context.Context, sourceURL, destURL string) (ClientData, error) {
	var clients ClientData
	var err error

	clients.SourceChain, clients.DestChainId, err = dial(ctx, destURL)
	if err != nil {
		return ClientData{}, fmt.Errorf("unable to dial source URL (%s): %w", sourceURL, err)
	}

	clients.DestChain, clients.DestChainId, err = dial(ctx, destURL)
	if err != nil {
		return ClientData{}, fmt.Errorf("unable to dial destination URL (%s): %w", destURL, err)
	}

	return clients, nil
}

// seqNumFromCCIPSendRequested is a helper to process logs in order to find sequence information.
func seqNumFromCCIPSendRequested(cl ClientData, sourceChainTx, CCIPMsgID string, logs []*types.Log) (SeqData, error) {
	var result SeqData

	abi, err := abi.JSON(strings.NewReader(helpers.OnRampABI))
	if err != nil {
		return SeqData{}, fmt.Errorf("unable to parse OnRampABI: %w", err)
	}
	var topic0 common.Hash
	for name, abiEvent := range abi.Events {
		if name == "CCIPSendRequested" {
			topic0 = abiEvent.ID
			break
		}
	}
	if topic0 == (common.Hash{}) {
		return SeqData{}, fmt.Errorf("no CCIPSendRequested event found in ABI")
	}
	var sendRequestedLogs []types.Log
	for _, sendReqLog := range logs {
		if sendReqLog.Topics[0] == topic0 && sendReqLog.TxHash == common.HexToHash(sourceChainTx) {
			result.OnRamp = sendReqLog.Address
			sendRequestedLogs = append(sendRequestedLogs, *sendReqLog)
		}
	}

	if len(sendRequestedLogs) == 0 {
		return SeqData{}, fmt.Errorf("no CCIPSendRequested logs found for in txReceipt for txhash %s", sourceChainTx)
	}
	onRampContract := bind.NewBoundContract(result.OnRamp, abi, cl.SourceChain, cl.SourceChain, cl.SourceChain)

	for _, sendReqLog := range sendRequestedLogs {
		var event helpers.SendRequestedEvent

		err = onRampContract.UnpackLog(&event, "CCIPSendRequested", sendReqLog)
		if err != nil {
			return SeqData{}, fmt.Errorf("unable to unpack onRampContract: %w", err)
		}

		if CCIPMsgID != "" &&
			"0x"+hex.EncodeToString(event.Message.MessageId[:]) != CCIPMsgID {
			continue
		}

		result.SeqNum = event.Message.SequenceNumber
		result.MsgID = event.Message.MessageId

		return result, nil
	}

	return SeqData{}, fmt.Errorf("send request not found in logs")
}

func approxDestStartBlock(cl ClientData, srcStartBlock *big.Int, destDeployedAt, destLatestBlock uint64) (uint64, error) {
	sourceBlockHdr, err := cl.SourceChain.HeaderByNumber(context.Background(), srcStartBlock)
	if err != nil {
		return 0, fmt.Errorf("unable to fetch source start block (%d) header: %w", srcStartBlock.Uint64(), err)
	}
	sendTxTime := sourceBlockHdr.Time
	maxBlockNum := destLatestBlock
	// setting this to an approx value of 1000 considering destination chain would have at least 1000 blocks before the transaction started
	minBlockNum := destDeployedAt
	closestBlockNum := uint64(math.Floor((float64(maxBlockNum) + float64(minBlockNum)) / 2))
	var closestBlockHdr *types.Header
	closestBlockHdr, err = cl.DestChain.HeaderByNumber(context.Background(), big.NewInt(int64(closestBlockNum)))
	if err != nil {
		return 0, fmt.Errorf("unable to fetch dest closest block (%d) header: %w", closestBlockNum, err)
	}
	// to reduce the number of RPC calls increase the value of blockOffset
	blockOffset := uint64(10)
	for {
		blockNum := closestBlockHdr.Number.Uint64()
		if minBlockNum > maxBlockNum {
			break
		}
		timeDiff := math.Abs(float64(closestBlockHdr.Time - sendTxTime))
		// break if the difference in timestamp is lesser than 1 minute
		if timeDiff < 60 {
			break
		} else if closestBlockHdr.Time > sendTxTime {
			maxBlockNum = blockNum - 1
		} else {
			minBlockNum = blockNum + 1
		}
		closestBlockNum = uint64(math.Floor((float64(maxBlockNum) + float64(minBlockNum)) / 2))
		closestBlockHdr, err = cl.DestChain.HeaderByNumber(context.Background(), big.NewInt(int64(closestBlockNum)))
		if err != nil {
			return 0, fmt.Errorf("unable to fetch dest closest block (%d) header: %w", closestBlockNum, err)
		}
	}

	for {
		if closestBlockHdr.Time <= sendTxTime {
			break
		}
		closestBlockNum = closestBlockNum - blockOffset
		if closestBlockNum <= 0 {
			return 0, fmt.Errorf("approx destination blocknumber not found")
		}
		closestBlockHdr, err = cl.DestChain.HeaderByNumber(context.Background(), big.NewInt(int64(closestBlockNum)))
		if err != nil {
			return 0, fmt.Errorf("unable to fetch dest closest block (%d) header: %w", closestBlockNum, err)
		}
	}

	result := closestBlockHdr.Number.Uint64()
	log.Printf("using approx destination start block number %d for filtering event", result)
	return result, nil
}

func GetChainData(cl ClientData, sourceChainTx, CCIPMsgID string, destStartBlock, destDeployedAt, numBlockLookback uint64) (ProofData, error) {
	result := ProofData{
		SourceChainId: cl.SourceChainId,
		DestChainId:   cl.DestChainId,
	}

	txReceipt, err := cl.SourceChain.TransactionReceipt(context.Background(), common.HexToHash(sourceChainTx))
	if err != nil {
		return ProofData{}, fmt.Errorf("unable to get txn receipt for src chain tx (%s): %w", sourceChainTx, err)
	}
	result.SrcStartBlock = big.NewInt(0).Sub(txReceipt.BlockNumber, big.NewInt(int64(numBlockLookback)))
	result.DestLatestBlock, err = cl.DestChain.BlockNumber(context.Background())
	if err != nil {
		return ProofData{}, fmt.Errorf("unable to resolve destination latest block: %w", err)
	}

	result.SeqData, err = seqNumFromCCIPSendRequested(cl, sourceChainTx, CCIPMsgID, txReceipt.Logs)
	if err != nil {
		return ProofData{}, fmt.Errorf("failed to lookup srq number: %w", err)
	}

	if destStartBlock < 1 {
		result.DestStartBlock, err = approxDestStartBlock(cl, result.SrcStartBlock, destDeployedAt, result.DestLatestBlock)
		if err != nil {
			return ProofData{}, fmt.Errorf("unable to approximate start block: %w", err)
		}
	} else {
		result.DestStartBlock = destStartBlock
	}

	return result, nil

}

func getCCIPChainSelector(chainId uint64) uint64 {
	selector, err := chainselectors.SelectorFromChainId(chainId)
	if err != nil {
		panic(fmt.Sprintf("no chain selector for %d", chainId))
	}
	return selector
}

func MakeExecutionReport(cfg Config, clients ClientData, proofInput ProofData) (helpers.InternalExecutionReport, error) {
	acceptedReportIterator, err := helpers.FilterReportAccepted(clients.DestChain, &bind.FilterOpts{Start: proofInput.DestStartBlock}, cfg.CommitStore)
	if err != nil {
		return helpers.InternalExecutionReport{}, fmt.Errorf("unable to filter accepted reports: %w", err)
	}

	var commitReport *helpers.ICommitStoreCommitReport
	for acceptedReportIterator.Next() {
		eventReport, err := acceptedReportIterator.CommitStoreReportAcceptedFromLog()
		if err != nil {
			return helpers.InternalExecutionReport{}, fmt.Errorf("unable to get commit store report: %w", err)
		}

		if eventReport.Report.Interval.Min <= proofInput.SeqNum && eventReport.Report.Interval.Max >= proofInput.SeqNum {
			commitReport = &eventReport.Report
			log.Println("Found root")
			break
		}
	}
	if commitReport == nil {
		return helpers.InternalExecutionReport{}, fmt.Errorf("unable to find root commit report for seq num %d", proofInput.SeqNum)
	}
	log.Println("Executing request manually")
	// Build a merkle tree for the report
	mctx := helpers.NewKeccakCtx()
	leafHasher := helpers.NewLeafHasher(
		getCCIPChainSelector(proofInput.SourceChainId.Uint64()),
		getCCIPChainSelector(proofInput.DestChainId.Uint64()),
		proofInput.OnRamp,
		mctx,
	)

	var leaves [][32]byte
	var curr, prove int
	var tokenData [][][]byte
	var msgs []helpers.InternalEVM2EVMMessage

	sendRequestedIterator, err := helpers.FilterCCIPSendRequested(clients.SourceChain, &bind.FilterOpts{
		Start: proofInput.SrcStartBlock.Uint64(),
	}, proofInput.OnRamp.Hex())
	if err != nil {
		return helpers.InternalExecutionReport{}, err
	}

	for sendRequestedIterator.Next() {
		event, err := sendRequestedIterator.SendRequestedEventFromLog()
		if err != nil {
			return helpers.InternalExecutionReport{}, fmt.Errorf("failed to get send request event: %w", err)
		}
		if event.Message.SequenceNumber <= commitReport.Interval.Max &&
			event.Message.SequenceNumber >= commitReport.Interval.Min {
			log.Println("Found seq num in commit report", event.Message.SequenceNumber, commitReport.Interval)
			hash, err := leafHasher.HashLeaf(sendRequestedIterator.Raw)
			if err != nil {
				return helpers.InternalExecutionReport{}, fmt.Errorf("failed to hash leaf: %w", err)
			}
			leaves = append(leaves, hash)
			if event.Message.SequenceNumber == proofInput.SeqNum && event.Message.MessageId == proofInput.MsgID {
				log.Printf("Found proving %d %+v\n\n", curr, event.Message)
				msgs = append(msgs, event.Message)

				var msgTokenData [][]byte
				for range event.Message.TokenAmounts {
					msgTokenData = append(msgTokenData, []byte{})
				}

				tokenData = append(tokenData, msgTokenData)
				prove = curr
			}
			curr++
		}
	}

	sendRequestedIterator.Close()
	if len(msgs) == 0 {
		return helpers.InternalExecutionReport{}, fmt.Errorf("unable to find msg with SeqNum %d", proofInput.SeqNum)
	}

	expectedNumberOfLeaves := int(commitReport.Interval.Max) - int(commitReport.Interval.Min) + 1
	if len(leaves) != expectedNumberOfLeaves {
		return helpers.InternalExecutionReport{}, fmt.Errorf("not enough leaves gather to build a commit root - want %d got %d. Please set NumberOfBlocks const to a higher value", expectedNumberOfLeaves, len(leaves))
	}

	tree, err := helpers.NewTree(mctx, leaves)
	if err != nil {
		return helpers.InternalExecutionReport{}, fmt.Errorf("failed to build new tree: %w", err)
	}
	if tree.Root() != commitReport.MerkleRoot {
		return helpers.InternalExecutionReport{}, fmt.Errorf("invalid proofInput: root doesn't match")
	}

	proof := tree.Prove([]int{prove})
	offRampProof := helpers.InternalExecutionReport{
		Messages:          msgs,
		Proofs:            proof.Hashes,
		OffchainTokenData: tokenData,
		ProofFlagBits:     helpers.ProofFlagsToBits(proof.SourceFlags),
	}

	return offRampProof, nil
}
