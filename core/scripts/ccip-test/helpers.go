package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const RetryTiming = 5 * time.Second
const CrossChainTimout = 5 * time.Minute
const TxInclusionTimout = 3 * time.Minute

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func WaitForMined(ctx context.Context, client *ethclient.Client, hash common.Hash, shouldSucceed bool) {
	maxIterations := TxInclusionTimout / RetryTiming
	for i := 0; i < int(maxIterations); i++ {
		fmt.Println("[MINING] waiting for tx to be mined...")
		receipt, _ := client.TransactionReceipt(ctx, hash)

		if receipt != nil {
			if shouldSucceed && receipt.Status == 0 {
				fmt.Println("[MINING] ERROR tx reverted!", hash.Hex())
				panic(receipt)
			} else if !shouldSucceed && receipt.Status != 0 {
				fmt.Println("[MINING] ERROR expected tx to revert!", hash.Hex())
				panic(receipt)
			}
			fmt.Println("[MINING] tx mined", hash.Hex(), "successful", shouldSucceed)
			return
		}

		time.Sleep(RetryTiming)
	}
	panic("No tx found within the given timeout")
}
