package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var networkNameMapping = map[string]string{
	"POLY": "polygon-mainnet",
	"BSC": "binance_smart_chain-mainnet",
	"ETH": "ethereum-mainnet",
	"AVAX": "avalanche-mainnet",
	"OPT": "ethereum-mainnet-optimism-1",
	"WEMIX": "wemix-mainnet",
	"KROMA": "ethereum-mainnet-kroma-1",
	"ARB": "ethereum-mainnet-arbitrum-1",
	"BASE": "ethereum-mainnet-base-1",

	"POLY_TEST": "polygon-testnet-mumbai",

	// Add more mappings here
}

func main() {

	fmt.Println("Network Name Mapping:", networkNameMapping)
	// Parse the command-line flags
	args := ParseFlags()
	fmt.Println("Parsed command-line arguments:", args)
	if args.Sender == "" {
		fmt.Println("The -sender flag is required.")
		os.Exit(1)
	}

	config, err := LoadConfig("networks.toml")
	if err != nil {
		fmt.Printf("Error loading TOML configuration: %v\n", err)
		return
	}

	resolvedFeeToken := resolveTokenAddress( args.Source, args.FeeToken, &config)
	//fmt.Println("Resolved fee token address:", strings.ToLower(resolvedFeeToken))
	url := GenerateQueryString(args.Sender, args.Receiver, args.Source, args.Dest, args.MessageId, resolvedFeeToken, args.First, args.Offset )
	fmt.Println("Querying API with URL:", url)
	apiResponse, nil := QueryTransactionsAPI(url)

	file, err := os.Create("transactions.csv")
	if err != nil {   
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Create a new map for reverse lookup
	var reversedNetworkNameMapping = make(map[string]string)

	// Fill the reversed map with values from networkNameMapping
	for shortCode, descriptiveName := range networkNameMapping {
		reversedNetworkNameMapping[descriptiveName] = shortCode
	}

	shortCode, exists := reversedNetworkNameMapping["ethereum-mainnet"]
	if exists {
		fmt.Println(shortCode) // This will print "ETH"
	} else {
		fmt.Println("Short code not found")
	}

	// Updated header to include destination fields
	header := []string{	"Source Network Name",  "Dest Network Name", "Fee Token", "Message URL", "Source Scan URL", 
						"Dest Scan URL", "Message ID", "Source Transaction Hash", "Dest Transaction Hash"}
	if err := writer.Write(header); err != nil {
		fmt.Println("Error writing header to CSV file:", err)
		return
	}

	for _, node := range apiResponse.Data.AllCcipTransactionsFlats.Nodes {
		networkName := node.SourceNetworkName
		tokenAddress := node.FeeToken

		networkConfig, exists := config.Networks[networkName]
		if !exists {
			fmt.Println("Network not found in config")
			return
		}

		tokenName, exists := findTokenNameByAddressInsensitive(networkConfig.Tokens, tokenAddress)
		if !exists {
			// fmt.Println("Token not found in network config")
			// return
			tokenName = tokenAddress
		}
	
		sourceScanURL := constructScanURL(node.SourceNetworkName, node.TransactionHash, &config)
		destScanURL := constructScanURL(node.DestNetworkName, node.DestTransactionHash, &config)
		messageURL := constructMessageURL(node.MessageID)
		record := []string{
			reversedNetworkNameMapping[node.SourceNetworkName],
			reversedNetworkNameMapping[node.DestNetworkName],
			tokenName,
			sourceScanURL,
			destScanURL,
			messageURL,
			node.TransactionHash,
			node.DestTransactionHash,
			node.MessageID,
		}
		if err := writer.Write(record); err != nil {
			fmt.Println("Error writing record to CSV file:", err)
			return
		}
		fmt.Println("-----")
		fmt.Printf("[msg](%s)\n", messageURL)
		fmt.Printf("[%s](%s)\n", reversedNetworkNameMapping[node.SourceNetworkName], sourceScanURL)
		fmt.Printf("[%s](%s)\n", reversedNetworkNameMapping[node.DestNetworkName], destScanURL)
		
		
	}

	fmt.Println("-----")
	fmt.Println("CSV file 'transactions.csv' created successfully with scan URLs for both source and destination.")

}
