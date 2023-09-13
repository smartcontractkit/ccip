package handler

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
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/arm_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/usdc_token_pool"
)

// RevertReasonFromErrorCodeString attempts to decode an error code string
func (h *BaseHandler) RevertReasonFromErrorCodeString(errorCodeString string) (string, error) {
	errorCodeString = strings.TrimPrefix(errorCodeString, "0x")
	return DecodeErrorStringFromABI(errorCodeString)
}

// RevertReasonFromTx attempts to fetch more info on failed TX
func (h *BaseHandler) RevertReasonFromTx(txHash string) (string, error) {
	// Need a node URL
	// NOTE: this node needs to run in archive mode
	ethUrl := h.cfg.NodeURL
	if ethUrl == "" {
		panicErr(errors.New("you must define ETH_NODE env variable"))
	}
	requester := h.cfg.FromAddress

	ec, err := ethclient.Dial(ethUrl)
	panicErr(err)
	errorString, _ := GetErrorForTx(ec, txHash, requester)

	return DecodeErrorStringFromABI(errorString)
}

func DecodeErrorStringFromABI(errorString string) (string, error) {
	contractABIs := getAllABIs()

	// Sanitize error string
	errorString = strings.TrimPrefix(errorString, "Reverted ")
	errorString = strings.TrimPrefix(errorString, "0x")

	data, err := hex.DecodeString(errorString)
	if err != nil {
		return "", errors.Wrap(err, "error decoding error string")
	}

	for _, contractABI := range contractABIs {
		parsedAbi, err2 := abi.JSON(strings.NewReader(contractABI))
		if err2 != nil {
			return "", errors.Wrap(err2, "error loading ABI")
		}

		for errorName, abiError := range parsedAbi.Errors {
			if bytes.Equal(data[:4], abiError.ID.Bytes()[:4]) {
				// Found a matching error
				v, err3 := abiError.Unpack(data)
				if err3 != nil {
					return "", errors.Wrap(err3, "error unpacking data")
				}

				// If exec error, the actual error is within the revert reason
				if errorName == "ExecutionError" || errorName == "TokenRateLimitError" || errorName == "TokenHandlingError" {
					// Get the inner type, which is `bytes`
					fmt.Printf("Error is \"%v\" inner error: ", errorName)
					errorBytes := v.([]interface{})[0].([]byte)
					if len(errorBytes) < 4 {
						return "[reverted without error code]", nil
					}
					return DecodeErrorStringFromABI(hex.EncodeToString(errorBytes))
				}
				return fmt.Sprintf("error is \"%v\" args %v\n", errorName, v), nil
			}
		}
	}

	if len(errorString) > 8 && errorString[:8] == "4e487b71" {
		fmt.Println("Assertion failure")
		indicator := errorString[len(errorString)-2:]
		switch indicator {
		case "01":
			return fmt.Sprintf("If you call assert with an argument that evaluates to false."), nil
		case "11":
			return fmt.Sprintf("If an arithmetic operation results in underflow or overflow outside of an unchecked { ... } block."), nil
		case "12":
			return fmt.Sprintf("If you divide or modulo by zero (e.g. 5 / 0 or 23 modulo 0)."), nil
		case "21":
			return fmt.Sprintf("If you convert a value that is too big or negative into an enum type."), nil
		case "31":
			return fmt.Sprintf("If you call .pop() on an empty array."), nil
		case "32":
			return fmt.Sprintf("If you access an array, bytesN or an array slice at an out-of-bounds or negative index (i.e. x[i] where i >= x.length or i < 0)."), nil
		case "41":
			return fmt.Sprintf("If you allocate too much memory or create an array that is too large."), nil
		case "51":
			return fmt.Sprintf("If you call a zero-initialized variable of internal function type."), nil
		default:
			return fmt.Sprintf("This is a revert produced by an assertion failure. Exact code not found \"%s\"", indicator), nil
		}
	}

	stringErr, err := abi.UnpackRevert(data)
	if err == nil {
		return fmt.Sprintf("string error: %s", stringErr), nil
	}

	return "", errors.Errorf("Cannot match error with contract ABI. Error code \"%v\"\n", errorString)
}

func getAllABIs() []string {
	return []string{
		arm_contract.ARMContractABI,
		lock_release_token_pool.LockReleaseTokenPoolABI,
		burn_mint_token_pool.BurnMintTokenPoolABI,
		usdc_token_pool.USDCTokenPoolABI,
		commit_store.CommitStoreABI,
		price_registry.PriceRegistryABI,
		evm_2_evm_onramp.EVM2EVMOnRampABI,
		evm_2_evm_offramp.EVM2EVMOffRampABI,
		router.RouterABI,
	}
}

func GetErrorForTx(client *ethclient.Client, txHash string, requester string) (string, error) {
	tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash(txHash))
	if err != nil {
		return "", errors.Wrap(err, "error getting transaction from hash")
	}
	re, err := client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		return "", errors.Wrap(err, "error getting transaction receipt")
	}

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

	return parseError(err)
}

func parseError(txError error) (string, error) {
	b, err := json.Marshal(txError)
	if err != nil {
		return "", err
	}
	var callErr struct {
		Code    int
		Data    string `json:"data"`
		Message string `json:"message"`
	}
	if json.Unmarshal(b, &callErr) != nil {
		return "", err
	}

	if callErr.Data == "" && strings.Contains(callErr.Message, "missing trie node") {
		return "", errors.Errorf("please use an archive node")
	}

	return callErr.Data, nil
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
