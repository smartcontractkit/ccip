package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlekSi/pointer"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// This script decodes the constructor arguments of a contract from a hex string
func main() {
	abiFilePath := flag.String("abiPath", "", "Absolute Path to the compiled contract ABI JSON file")
	encodedConstructorArgs := flag.String("encodedConstructorArgs", "", "Hex encoded constructor arguments")
	deploymentTx := flag.String("deploymentTx", "", "Contract Deployment transaction hash")
	rpcURL := flag.String("rpcURL", "", "RPC URL for the chain")
	flag.Parse()
	if pointer.GetString(encodedConstructorArgs) == "" && (pointer.GetString(deploymentTx) == "" || pointer.GetString(rpcURL) == "") {
		log.Fatalf("Encoded constructor arguments or deploymentTx and rpcURL must be provided")
	}

	if !filepath.IsAbs(pointer.GetString(abiFilePath)) {
		log.Fatalf("Abi file path must be absolute")
	}
	var params string
	if pointer.GetString(encodedConstructorArgs) == "" {
		// Get the contract deployment transaction receipt
		client, err := ethclient.Dial(pointer.GetString(rpcURL))
		if err != nil {
			log.Fatalf("Failed to connect to the rpc client: %v", err)
		}
		tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash(pointer.GetString(deploymentTx)))
		if err != nil {
			log.Fatalf("Failed to get transaction receipt: %v", err)
		}
		params = string(tx.Data())
	} else {
		params = pointer.GetString(encodedConstructorArgs)
	}
	// Read the ABI JSON file
	abiFileContent, err := os.ReadFile(pointer.GetString(abiFilePath))
	if err != nil {
		log.Fatalf("Failed to read ABI file: %v", err)
	}

	// Parse the JSON content to extract the ABI and deployed bytecode
	var compiledFile struct {
		ABI              json.RawMessage `json:"abi"`
		DeployedBytecode string          `json:"deployedBytecode"`
	}
	if err = json.Unmarshal(abiFileContent, &compiledFile); err != nil {
		log.Fatalf("Failed to unmarshal ABI file content: %v", err)
	}

	// Parse the ABI
	parsedABI, err := abi.JSON(strings.NewReader(string(compiledFile.ABI)))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	// Decode the constructor arguments from the input hex
	encodedParamsBytes, err := hex.DecodeString(params)
	if err != nil {
		log.Fatalf("Failed to decode hex string: %v", err)
	}

	// fmt.Println(hex.EncodeToString(encodedParamsBytes[0 : 4+(4*32)]))
	encodedParamsBytes = encodedParamsBytes[4+(4*32):]

	// Use the constructor arguments to unpack the values
	decodedArgs, err := parsedABI.Constructor.Inputs.Unpack(encodedParamsBytes)
	if err != nil {
		log.Fatalf("Failed to unpack constructor arguments: %v", err)
	}

	// Create a map to hold the named constructor arguments
	constructorArgsMap := make(map[string]interface{})
	for i, arg := range parsedABI.Constructor.Inputs {
		fmt.Println(arg.Name)
		constructorArgsMap[arg.Name] = decodedArgs[i]
	}

	// Convert decoded arguments to JSON
	decodedArgsJSON, err := json.MarshalIndent(constructorArgsMap, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal decoded arguments to JSON: %v", err)
	}

	fmt.Println("Decoded Constructor Arguments in JSON Format:")
	fmt.Println(string(decodedArgsJSON))
}
