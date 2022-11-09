package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_subscription_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/secrets"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
)

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Need a node URL
	// NOTE: this node needs to run in archive mode
	ethUrl := secrets.GetRPC(big.NewInt(5))
	txHash := "0x97be8559164442595aba46b5f849c23257905b78e72ee43d9b998b28eee78b84"
	requester := "0xe88ff73814fb891bb0e149f5578796fa41f20242"

	ec, ethErr := ethclient.Dial(ethUrl)
	panicErr(ethErr)
	errorString, contractAddress := getErrorForTx(ec, txHash, requester)
	// Some nodes prepend "Reverted "
	trimmed := strings.TrimPrefix(errorString, "Reverted ")
	// Remove 0x before hex decoding
	data, err := hex.DecodeString(trimmed[2:])
	panicErr(err)

	contractABIs := getABIForContract(ec, contractAddress)

	for _, contractABI := range contractABIs {
		parsedAbi, err2 := abi.JSON(strings.NewReader(contractABI))
		panicErr(err2)

		for k, abiError := range parsedAbi.Errors {
			if bytes.Equal(data[:4], abiError.ID.Bytes()[:4]) {
				// Found a matching error
				v, err3 := abiError.Unpack(data)
				panicErr(err3)
				fmt.Printf("Error is \"%v\" args %v\n", k, v)
				return
			}
		}
	}

	fmt.Printf("Cannot match error with contract ABI. Error code \"%v\"\n", trimmed)
}

// getABIForContract. Since contracts interact with other contracts we return all ABIs we expect the given
// contract to interact with
func getABIForContract(client *ethclient.Client, contractAddress common.Address) []string {
	contractType, _, err := ccip.TypeAndVersion(contractAddress, client)
	panicErr(err)

	switch contractType {
	// TOLL
	case ccip.EVM2EVMTollOnRamp:
		return []string{evm_2_evm_toll_onramp.EVM2EVMTollOnRampABI}
	case ccip.EVM2EVMTollOffRamp:
		return []string{any_2_evm_toll_offramp.EVM2EVMTollOffRampABI, any_2_evm_toll_offramp_router.Any2EVMTollOffRampRouterABI, commit_store.CommitStoreABI}
		// SUBSCRIPTION
	case ccip.EVM2EVMSubscriptionOnRamp:
		return []string{evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampABI}
	case ccip.EVM2EVMSubscriptionOffRamp:
		return []string{any_2_evm_subscription_offramp.EVM2EVMSubscriptionOffRampABI, any_2_evm_subscription_offramp_router.Any2EVMSubscriptionOffRampRouterABI, commit_store.CommitStoreABI}
		// SHARED
	case ccip.CommitStore:
		return []string{commit_store.CommitStoreABI}
	}
	panic("Contract not found")
}

func getErrorForTx(client *ethclient.Client, txHash string, requester string) (string, common.Address) {
	tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash(txHash))
	panicErr(err)
	re, err := client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	panicErr(err)

	call := ethereum.CallMsg{
		From:     common.HexToAddress(requester),
		To:       tx.To(),
		Data:     tx.Data(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
	}
	_, err = client.CallContract(context.Background(), call, re.BlockNumber)
	if err == nil {
		panic("no error calling contract")
	}

	return parseError(err), *tx.To()
}

func parseError(txError error) string {
	b, err := json.Marshal(txError)
	panicErr(err)
	var callErr struct {
		Code    int
		Data    string `json:"data"`
		Message string `json:"message"`
	}
	err = json.Unmarshal(b, &callErr)
	panicErr(err)

	if callErr.Data == "" && strings.Contains(callErr.Message, "missing trie node") {
		panic("Use an archive node")
	}
	return callErr.Data
}
