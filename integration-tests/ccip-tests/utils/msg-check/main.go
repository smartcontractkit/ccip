package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var networkNameMapping = map[string]string{
	"polygon-mainnet": "POLY",
	"binance_smart_chain-mainnet": "BSC",
	"ethereum-mainnet": "ETH",
	"avalanche-mainnet": "AVAX",
	"ethereum-mainnet-optimism-1": "OPT",
	"wemix-mainnet": "WEMIX",
	"ethereum-mainnet-kroma-1": "KROMA",
	"ethereum-mainnet-arbitrum-1": "ARB",
	"ethereum-mainnet-base-1": "BASE",

	// testnet mappings
	"polygon-testnet-mumbai": "POLY_TEST",
	"ethereum-testnet-sepolia-arbitrum-1": "ARBSEP",
	"ethereum-testnet-sepolia": "SEP",
	"ethereum-testnet-sepolia-optimism-1": "OPTSEP",

	// Add more reversed mappings here
}

func main() {
	// Define and parse the command-line flag
	var messageId string
	flag.StringVar(&messageId, "msg", "", "Message ID to query")
	flag.Parse()

	if messageId == "" {
		log.Fatal("You must provide a message ID using the -msg flag.")
	}

	tokenMap, err := LoadTokenMapFromToml("tokens.toml")
	if err != nil {
		fmt.Println("Error loading token map:", err)
		return
	}

	// Generate the query string
	queryURL := GenerateMessageQueryString(messageId)
	fmt.Println("Querying the API with URL:", queryURL)

	// Query the API
	response, err := QueryMessageAPI(queryURL)
	if err != nil {
		log.Fatalf("Failed to query the API: %v", err)
	}


	if len(response.Data.AllCcipMessages.Nodes) > 0 {
		node := response.Data.AllCcipMessages.Nodes[0]

		layout := "2006-01-02T15:04:05"
		sendTime, _ := time.Parse(layout, node.SendTimestamp)
		sendFinal, _ := time.Parse(layout, node.SendFinalized)
		commitTime, _ := time.Parse(layout, node.CommitBlockTimestamp)
		blessTime, _ := time.Parse(layout, node.BlessBlockTimestamp)
		receiptTime, _ := time.Parse(layout, node.ReceiptTimestamp)
		receiptFinal, _ := time.Parse(layout, node.ReceiptFinalized)
	
		sendToSendFinalDuration := sendFinal.Sub(sendTime)
		sendFinalToCommitDuration := commitTime.Sub(sendFinal)
		commitToBlessDuration := blessTime.Sub(commitTime)
		blessToReceiptDuration := receiptTime.Sub(blessTime)
		receiptToFinalDuration := receiptFinal.Sub(receiptTime)
		
		sourceNetworkMapped := networkNameMapping[node.SourceNetworkName]
		destNetworkMapped := networkNameMapping[node.DestNetworkName]
		
		fmt.Printf("Message ID: %s\n", node.MessageId)
		resolvedFeeToken, exists := GetTokenName( tokenMap, node.FeeToken)
		if exists {
			fmt.Println("Fee token: ", resolvedFeeToken)
		} else {
			fmt.Println("Fee token: ", node.FeeToken)
		}
		fmt.Println("Send timestamp:", node.SendTimestamp)
		fmt.Println("Lane: ", sourceNetworkMapped, "to", destNetworkMapped)
		fmt.Printf("%-24s %s\n", "Send to SendFinal:", sendToSendFinalDuration)
		fmt.Printf("%-24s %s\n", "SendFinal to Commit:", sendFinalToCommitDuration)
		fmt.Printf("%-24s %s\n", "Commit to Bless:", commitToBlessDuration)
		fmt.Printf("%-24s %s\n", "Bless to Receipt:", blessToReceiptDuration)
		fmt.Printf("%-24s %s\n", "Receipt to ReceiptFinal:", receiptToFinalDuration)

	} else {
		fmt.Println("No data found for the provided message ID")
	}


}
