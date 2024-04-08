package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

)

type ApiTransactionsResponse struct {
	Data struct {
		AllCcipTransactionsFlats struct {
			Nodes []struct {
				TransactionHash     string `json:"transactionHash"`
				FeeToken            string `json:"feeToken"`
				DestTransactionHash string `json:"destTransactionHash"`
				SourceNetworkName   string `json:"sourceNetworkName"`
				DestNetworkName     string `json:"destNetworkName"`
				MessageID           string `json:"messageId"`
				BlockTimestamp      string `json:"blockTimestamp"`
			} `json:"nodes"`
		} `json:"allCcipTransactionsFlats"`
	} `json:"data"`
}

type ChainlinkResponse struct {
    Data struct {
        AllCcipMessages struct {
            Nodes []Node `json:"nodes"`
        } `json:"allCcipMessages"`
    } `json:"data"`
}

type Node struct {
    Arm                   string        `json:"arm"`
    BlessBlockNumber      int           `json:"blessBlockNumber"`
    BlessBlockTimestamp   string        `json:"blessBlockTimestamp"`
    BlessLogIndex         int           `json:"blessLogIndex"`
    BlockTimestamp        string        `json:"blockTimestamp"`
    CommitBlockNumber     int           `json:"commitBlockNumber"`
    BlessTransactionHash  string        `json:"blessTransactionHash"`
    CommitBlockTimestamp  string        `json:"commitBlockTimestamp"`
    CommitLogIndex        int           `json:"commitLogIndex"`
    CommitStore           string        `json:"commitStore"`
    CommitTransactionHash string        `json:"commitTransactionHash"`
    Data                  string        `json:"data"`
    DestChainId           string        `json:"destChainId"`
    DestRouterAddress     string        `json:"destRouterAddress"`
    DestNetworkName       string        `json:"destNetworkName"`
    TokenAmounts          []TokenAmount `json:"tokenAmounts"`
    Votes                 int           `json:"votes"`
    Strict                bool          `json:"strict"`
    SourceChainId         string        `json:"sourceChainId"`
    State                 int           `json:"state"`
    SourceNetworkName     string        `json:"sourceNetworkName"`
    SequenceNumber        int           `json:"sequenceNumber"`
    Sender                string        `json:"sender"`
    SendTransactionHash   string        `json:"sendTransactionHash"`
    SendTimestamp         string        `json:"sendTimestamp"`
    SendLogIndex          int           `json:"sendLogIndex"`
    SendBlock             int           `json:"sendBlock"`
    SendFinalized         string        `json:"sendFinalized"`
    RouterAddress         string        `json:"routerAddress"`
    Receiver              string        `json:"receiver"`
    Root                  string        `json:"root"`
    ReceiptTransactionHash string      `json:"receiptTransactionHash"`
    ReceiptTimestamp      string        `json:"receiptTimestamp"`
    ReceiptLogIndex       int           `json:"receiptLogIndex"`
    ReceiptBlock          int           `json:"receiptBlock"`
    ReceiptFinalized      string        `json:"receiptFinalized"`
    OnrampAddress         string        `json:"onrampAddress"`
    OfframpAddress        string        `json:"offrampAddress"`
    Nonce                 int           `json:"nonce"`
    Min                   string        `json:"min"`
    MessageId             string        `json:"messageId"`
    Max                   string        `json:"max"`
    Info                  Info          `json:"info"`
    GasLimit              string        `json:"gasLimit"`
    FeeTokenAmount        string        `json:"feeTokenAmount"`
    FeeToken              string        `json:"feeToken"`
}

// TokenAmount represents each object within the "tokenAmounts" array
type TokenAmount struct {
    Token  string `json:"token"`
    Amount string `json:"amount"`
}

// Info represents the "info" object within a Node
type Info struct {
    Data              string        `json:"data"`
    Nonce             int           `json:"nonce"`
    Sender            string        `json:"sender"`
    Strict            bool          `json:"strict"`
    FeeToken          string        `json:"feeToken"`
    GasLimit          int           `json:"gasLimit"`
    Receiver          string        `json:"receiver"`
    MessageId         string        `json:"messageId"`
    TokenAmounts      []TokenAmount `json:"tokenAmounts"`
    FeeTokenAmount    int64         `json:"feeTokenAmount"`
    SequenceNumber    int           `json:"sequenceNumber"`
    SourceTokenData   []string      `json:"sourceTokenData"`
    SourceChainSelector int64       `json:"sourceChainSelector"`
}

func GenerateQueryString(sender string, receiver string, sourceNetworkName string,
	destNetworkName string, messageId string, feeToken string,
	first int, offset int) string {
	baseURL := "https://ccip.chain.link/api/query/LATEST_TRANSACTIONS_QUERY"
	queryParams := url.Values{}
	//queryParams.Add("query", "LATEST_TRANSACTIONS_QUERY")

	// Construct the condition map dynamically based on provided arguments
	// Construct the condition map dynamically based on provided arguments
condition := make(map[string]interface{})

// Add sender to the condition map if it's not empty
if sender != "" {
    condition["sender"] = sender
}

// Add receiver to the condition map if it's not empty
if receiver != "" {
    condition["receiver"] = receiver
}

// Add sourceNetworkName to the condition map if it's not empty
if sourceNetworkName != "" {
    condition["sourceNetworkName"] = sourceNetworkName
}

// Add destNetworkName to the condition map if it's not empty
if destNetworkName != "" {
    condition["destNetworkName"] = destNetworkName
}

// Add messageId to the condition map if it's not empty
if messageId != "" {
    condition["messageId"] = messageId
}

// Add feeToken to the condition map if it's not empty
if feeToken != "" {
    condition["feeToken"] = feeToken
}

	// Convert the condition map and other variables into a JSON string
	variables := map[string]interface{}{
		"first":     first,
		"offset":    offset,
		"condition": condition,
	}
	variablesJSON, err := json.Marshal(variables)
	if err != nil {
		fmt.Println("Error marshalling variables to JSON:", err)
		return ""
	}

	queryParams.Add("variables", string(variablesJSON))

	return baseURL + "?" + queryParams.Encode()
}

func QueryTransactionsAPI(url string) (*ApiTransactionsResponse, error) {

	apiResponse := ApiTransactionsResponse{}
	method := "GET"
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error making request to API: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return &apiResponse, nil
}

func QueryMessageAPI(url string) (*ChainlinkResponse, error) {
	
	apiResponse := ChainlinkResponse{}
	method := "GET"
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error making request to API: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return &apiResponse, nil
}
