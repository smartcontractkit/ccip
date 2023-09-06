package customtokens

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
)

type USDCService struct {
	attestationApi  string
	sourceChainId   uint64
	SourceUSDCToken common.Address
}

type USDCAttestationResponse struct {
	Status      USDCAttestationStatus `json:"status"`
	Attestation string                `json:"attestation"`
}

// Hard coded mapping of chain id to USDC token addresses
// Will be removed in favour of more flexible solution.
var USDCTokenMapping = map[uint64]common.Address{
	420:    common.HexToAddress("0xe05606174bac4A6364B31bd0eCA4bf4dD368f8C6"),
	43113:  common.HexToAddress("0x5425890298aed601595a70ab815c96711a31bc65"),
	80001:  common.HexToAddress("0x9999f7fea5938fd3b1e26a12c3f2fb024e194f97"),
	84531:  common.HexToAddress("0xf175520c52418dfe19c8098071a252da48cd1c19"),
	421613: common.HexToAddress("0xfd064A18f3BF249cf1f87FC203E90D8f650f2d63"),
}

const (
	version                       = "v1"
	attestationPath               = "attestations"
	eventSignatureString          = "8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b036"
	USDC_MESSAGE_SENT_FILTER_NAME = "USDC message sent"
)

var USDC_MESSAGE_SENT = common.HexToHash(eventSignatureString)

type USDCAttestationStatus string

const (
	USDCAttestationStatusSuccess USDCAttestationStatus = "complete"
	USDCAttestationStatusPending USDCAttestationStatus = "pending_confirmations"
)

func NewUSDCService(usdcAttestationApi string, sourceChainId uint64) *USDCService {
	return &USDCService{attestationApi: usdcAttestationApi, sourceChainId: sourceChainId, SourceUSDCToken: USDCTokenMapping[sourceChainId]}
}

func (usdc *USDCService) TryGetAttestation(messageHash string) (USDCAttestationResponse, error) {
	fullAttestationUrl := fmt.Sprintf("%s/%s/%s/%s", usdc.attestationApi, version, attestationPath, messageHash)
	req, err := http.NewRequest("GET", fullAttestationUrl, nil)
	if err != nil {
		return USDCAttestationResponse{}, err
	}
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return USDCAttestationResponse{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return USDCAttestationResponse{}, err
	}

	var response USDCAttestationResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return USDCAttestationResponse{}, err
	}

	return response, nil
}

func (usdc *USDCService) IsAttestationComplete(messageHash string) (bool, string, error) {
	response, err := usdc.TryGetAttestation(messageHash)
	if err != nil {
		return false, "", err
	}
	if response.Status == USDCAttestationStatusSuccess {
		return true, response.Attestation, nil
	}
	return false, "", nil
}

func GetUSDCServiceSourceLPFilters(usdcTokenAddress common.Address) []logpoller.Filter {
	return []logpoller.Filter{
		{
			Name:      logpoller.FilterName(USDC_MESSAGE_SENT_FILTER_NAME, usdcTokenAddress.Hex()),
			EventSigs: []common.Hash{common.HexToHash(eventSignatureString)},
			Addresses: []common.Address{usdcTokenAddress},
		},
	}
}
