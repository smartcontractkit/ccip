package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp"
)

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Need a node URL
	// NOTE: this node needs to run in archive mode
	ec, err := ethclient.Dial("wss://kovan.infura.io/ws/v3/myapikey")
	panicErr(err)
	// txHash in question
	txHash := "0x578aea01b3ed827ec743fcd6179f81b8b0ac6b3d1ba2ea281d6fc4ec89498088"
	tx, _, err := ec.TransactionByHash(context.Background(), common.HexToHash(txHash))
	panicErr(err)
	re, err := ec.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	panicErr(err)

	// Random requester
	requester := common.HexToAddress("9ca9d2d5e04012c9ed24c0e513c9bfaa4a2dd77f")
	call := ethereum.CallMsg{
		From:     requester,
		To:       tx.To(),
		Data:     tx.Data(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
	}
	_, err = ec.CallContract(context.Background(), call, re.BlockNumber)
	if err == nil {
		panic("no error calling contract")
	}

	b, err := json.Marshal(err)
	panicErr(err)
	var callErr struct {
		Data string `json:"data"`
	}
	err = json.Unmarshal(b, &callErr)
	panicErr(err)

	data, err := hex.DecodeString(callErr.Data[2:])
	panicErr(err)
	offrampABI, err := abi.JSON(strings.NewReader(any_2_evm_toll_offramp.EVM2EVMTollOffRampABI))
	panicErr(err)

	for k, abiError := range offrampABI.Errors {
		if bytes.Equal(data[:4], abiError.ID.Bytes()[:4]) {
			// Found a matching error
			v, err := abiError.Unpack(data)
			panicErr(err)
			fmt.Printf("Error is %v args %v\n", k, v)
		}
	}
}
