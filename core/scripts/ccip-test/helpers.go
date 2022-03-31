package main

import (
	"context"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/logger"
)

const RetryTiming = 5 * time.Second
const CrossChainTimout = 5 * time.Minute
const TxInclusionTimout = 3 * time.Minute

// WaitForMined wait for a tx to be included on chain. It will panic when
// the tx is reverted/successful based on the shouldSucceed parameter.
func WaitForMined(t *testing.T, lggr logger.Logger, client ethereum.TransactionReader, hash common.Hash, shouldSucceed bool) {
	maxIterations := TxInclusionTimout / RetryTiming
	for i := 0; i < int(maxIterations); i++ {
		lggr.Info("[MINING] waiting for tx to be mined...")
		receipt, _ := client.TransactionReceipt(context.Background(), hash)

		if receipt != nil {
			if shouldSucceed && receipt.Status == 0 {
				lggr.Infof("[MINING] ERROR tx reverted %s", hash.Hex())
				panic(receipt)
			} else if !shouldSucceed && receipt.Status != 0 {
				lggr.Infof("[MINING] ERROR expected tx to revert %s", hash.Hex())
				panic(receipt)
			}
			lggr.Infof("[MINING] tx mined %s successful %t", hash.Hex(), shouldSucceed)
			return
		}

		time.Sleep(RetryTiming)
	}
	t.Error("No tx found within the given timeout")
}

// SetGasFees configures the chain client with the given EVMGasSettings. This method is needed for EIP txs
// to function because of the geth-only tip fee method.
func SetGasFees(owner *bind.TransactOpts, config EVMGasSettings) {
	if config.EIP1559 {
		// to not use geth-only tip fee method when EIP1559 is enabled
		// https://github.com/ethereum/go-ethereum/pull/23484
		owner.GasTipCap = config.GasTipCap
	} else {
		owner.GasPrice = config.GasPrice
	}
}
