package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

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
    Token  string      `json:"token"`
    //Amount string    `json:"amount"` 
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

func GenerateMessageQueryString(messageId string ) string {
	baseURL := "https://ccip.chain.link/api/query/MESSAGE_DETAILS_QUERY"
	queryParams := url.Values{}
	//queryParams.Add("query", "MESSAGE_DETAILS_QUERY")

	// Convert the condition map and other variables into a JSON string
	variables := map[string]interface{}{
		"messageId":     messageId,
	}

	variablesJSON, err := json.Marshal(variables)
	if err != nil {
		fmt.Println("Error marshalling variables to JSON:", err)
		return ""
	}

	queryParams.Add("variables", string(variablesJSON))
	return baseURL + "?" + queryParams.Encode()

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
