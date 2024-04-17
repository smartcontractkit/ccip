package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp_1_2_0"
)

type txDetails struct {
	SequenceNumber uint64
	TxHash         string
	MessageId      string
}

func main() {
	seqNumsInput := flag.String("seqnums", "", "Enter sequence numbers separated by comma (e.g., 1,2,3)")
	onRampAddress := flag.String("onRamp", "", "Enter on-ramp address")
	startBlock := flag.Uint64("startblock", 0, "Enter start block as starting block num to look for txs")
	rpcURL := flag.String("rpc", "", "Enter RPC endpoint")

	// Parse the flags
	flag.Parse()
	seqNumbers := make(map[uint64]txDetails)
	log.SetOutput(os.Stdout)
	if *seqNumsInput == "" {
		log.Fatalf("Please provide sequence numbers")
	}
	if !common.IsHexAddress(*onRampAddress) {
		log.Fatalf("Please provide valid on-ramp address")
	}
	if *startBlock == 0 {
		log.Fatalf("Please provide start block")
	}
	if *rpcURL == "" {
		log.Fatalf("Please provide RPC URL")
	}

	if *seqNumsInput != "" {
		for _, numStr := range strings.Split(*seqNumsInput, ",") {
			num, err := strconv.ParseUint(strings.TrimSpace(numStr), 10, 64)
			if err != nil {
				fmt.Printf("Error parsing '%s' as uint64: %v\n", numStr, err)
				return
			}
			seqNumbers[num] = txDetails{}
		}
	}

	ec, err := ethclient.Dial(*rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to the rpc client: %v", err)
	}
	onRamp, err := evm_2_evm_onramp_1_2_0.NewEVM2EVMOnRamp(common.HexToAddress(*onRampAddress), ec)
	if err != nil {
		log.Fatalf("Failed to instantiate the on-ramp contract: %v", err)
	}
	sendRequested, err := onRamp.FilterCCIPSendRequested(&bind.FilterOpts{
		Start: *startBlock,
	})
	if err != nil {
		log.Fatalf("Failed to filter CCIPSendRequested events: %v", err)
	}

	for sendRequested.Next() {
		if _, exist := seqNumbers[sendRequested.Event.Message.SequenceNumber]; exist {
			log.Printf("Found sequence number %d in tx %s\n", sendRequested.Event.Message.SequenceNumber, sendRequested.Event.Raw.TxHash.String())
			seqNumbers[sendRequested.Event.Message.SequenceNumber] = txDetails{
				SequenceNumber: sendRequested.Event.Message.SequenceNumber,
				TxHash:         sendRequested.Event.Raw.TxHash.String(),
				MessageId:      fmt.Sprintf("%x", sendRequested.Event.Message.MessageId[:]),
			}
		}
	}
	log.Println("Tx Details found for Sequence numbers:", seqNumbers)
}
