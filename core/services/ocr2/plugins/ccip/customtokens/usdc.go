package customtokens

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/ccipevents"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

type USDCService struct {
	sourceChainEvents ccipevents.Client
	attestationApi    string
	sourceChainId     uint64
	SourceUSDCToken   common.Address
	OnRampAddress     common.Address

	// Cache of sequence number -> attestation attempt
	attestationCache map[uint64]USDCAttestationAttempt
}

type USDCAttestationAttempt struct {
	USDCMessageBody         []byte
	USDCMessageHash         [32]byte
	CCIPSendTxHash          common.Hash
	CCIPSendLogIndex        int64
	USDCAttestationResponse AttestationResponse
}

type AttestationResponse struct {
	Status      AttestationStatus `json:"status"`
	Attestation string            `json:"attestation"`
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

type AttestationStatus string

const (
	USDCAttestationStatusSuccess   AttestationStatus = "complete"
	USDCAttestationStatusPending   AttestationStatus = "pending_confirmations"
	USDCAttestationStatusUnchecked AttestationStatus = "unchecked"
)

func NewUSDCService(sourceChainEvents ccipevents.Client, onRampAddress common.Address, usdcAttestationApi string, sourceChainId uint64) *USDCService {
	return &USDCService{
		sourceChainEvents: sourceChainEvents,
		attestationApi:    usdcAttestationApi,
		sourceChainId:     sourceChainId,
		SourceUSDCToken:   USDCTokenMapping[sourceChainId],
		OnRampAddress:     onRampAddress,
		attestationCache:  make(map[uint64]USDCAttestationAttempt),
	}
}

func (usdc *USDCService) IsAttestationComplete(ctx context.Context, seqNum uint64) (bool, string, error) {
	response, err := usdc.GetUpdatedAttestation(ctx, seqNum)
	if err != nil {
		return false, "", err
	}

	if response.Status == USDCAttestationStatusSuccess {
		return true, response.Attestation, nil
	}
	return false, "", nil
}

func (usdc *USDCService) GetUpdatedAttestation(ctx context.Context, seqNum uint64) (AttestationResponse, error) {
	// Try to get information from cache to reduce the number of database and external calls
	attestationAttempt, ok := usdc.attestationCache[seqNum]
	if ok && attestationAttempt.USDCAttestationResponse.Status == USDCAttestationStatusSuccess {
		// If successful, return the cached response
		return attestationAttempt.USDCAttestationResponse, nil
	}

	// If no attempt for this message id exists, get the required data to make one
	if !ok {
		var err error
		attestationAttempt, err = usdc.getAttemptInfoFromCCIPMessageId(ctx, seqNum)
		if err != nil {
			return AttestationResponse{}, err
		}
		// Save the attempt in the cache in case the external call fails
		usdc.attestationCache[seqNum] = attestationAttempt
	}

	response, err := usdc.callAttestationApi(ctx, attestationAttempt.USDCMessageHash)
	if err != nil {
		return AttestationResponse{}, err
	}

	// Save the response in the cache
	attestationAttempt.USDCAttestationResponse = response
	usdc.attestationCache[seqNum] = attestationAttempt

	return response, nil
}

func (usdc *USDCService) getAttemptInfoFromCCIPMessageId(ctx context.Context, seqNum uint64) (USDCAttestationAttempt, error) {
	// Get the CCIP message send event from the log poller
	ccipSendRequests, err := usdc.sourceChainEvents.GetSendRequestsBetweenSeqNums(ctx, usdc.OnRampAddress, seqNum, seqNum, 0)
	if err != nil {
		return USDCAttestationAttempt{}, err
	}

	if len(ccipSendRequests) != 1 {
		return USDCAttestationAttempt{}, fmt.Errorf("expected 1 CCIP send request, got %d", len(ccipSendRequests))
	}

	ccipSendRequest := ccipSendRequests[0]

	ccipSendTxHash := ccipSendRequest.Data.Raw.TxHash
	ccipSendLogIndex := int64(ccipSendRequest.Data.Raw.Index)

	// Get the USDC message body
	usdcMessageBody, err := usdc.sourceChainEvents.GetLastUSDCMessagePriorToLogIndexInTx(ctx, ccipSendLogIndex, ccipSendTxHash)
	if err != nil {
		return USDCAttestationAttempt{}, err
	}

	return USDCAttestationAttempt{
		USDCMessageBody:  usdcMessageBody,
		USDCMessageHash:  utils.Keccak256Fixed(usdcMessageBody),
		CCIPSendTxHash:   ccipSendTxHash,
		CCIPSendLogIndex: ccipSendLogIndex,
		USDCAttestationResponse: AttestationResponse{
			Status:      USDCAttestationStatusUnchecked,
			Attestation: "",
		},
	}, nil
}

func (usdc *USDCService) callAttestationApi(ctx context.Context, usdcMessageHash [32]byte) (AttestationResponse, error) {
	fullAttestationUrl := fmt.Sprintf("%s/%s/%s/0x%s", usdc.attestationApi, version, attestationPath, hex.EncodeToString(usdcMessageHash[:]))
	req, err := http.NewRequestWithContext(ctx, "GET", fullAttestationUrl, nil)
	if err != nil {
		return AttestationResponse{}, err
	}
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return AttestationResponse{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return AttestationResponse{}, err
	}

	var response AttestationResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return AttestationResponse{}, err
	}
	return response, nil
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
