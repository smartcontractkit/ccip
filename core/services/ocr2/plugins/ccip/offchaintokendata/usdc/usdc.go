package usdc

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/ccipevents"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/offchaintokendata"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

type OffchainTokenDataService struct {
	sourceChainEvents  ccipevents.Client
	attestationApi     string
	messageTransmitter common.Address
	sourceToken        common.Address
	onRampAddress      common.Address

	// Cache of sequence number -> attestation attempt
	attestationCache map[uint64]AttestationAttempt
}

type AttestationAttempt struct {
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
var TokenMapping = map[uint64]common.Address{
	420:    common.HexToAddress("0xe05606174bac4A6364B31bd0eCA4bf4dD368f8C6"),
	43113:  common.HexToAddress("0x5425890298aed601595a70ab815c96711a31bc65"),
	80001:  common.HexToAddress("0x9999f7fea5938fd3b1e26a12c3f2fb024e194f97"),
	84531:  common.HexToAddress("0xf175520c52418dfe19c8098071a252da48cd1c19"),
	421613: common.HexToAddress("0xfd064A18f3BF249cf1f87FC203E90D8f650f2d63"),
}

var messageTransmitterMapping = map[uint64]common.Address{
	420: common.HexToAddress("0x8c5261668696ce22758910d05bab8f186d6eb247"),
}

const (
	version                  = "v1"
	attestationPath          = "attestations"
	MESSAGE_SENT_FILTER_NAME = "USDC message sent"
)

type AttestationStatus string

const (
	AttestationStatusSuccess   AttestationStatus = "complete"
	AttestationStatusPending   AttestationStatus = "pending_confirmations"
	AttestationStatusUnchecked AttestationStatus = "unchecked"
)

var _ offchaintokendata.Provider = &OffchainTokenDataService{}

func NewUSDCOffchainTokenDataService(sourceChainEvents ccipevents.Client, usdcTokenAddress, onRampAddress common.Address, usdcAttestationApi string, sourceChainId uint64) *OffchainTokenDataService {
	return &OffchainTokenDataService{
		sourceChainEvents:  sourceChainEvents,
		attestationApi:     usdcAttestationApi,
		messageTransmitter: messageTransmitterMapping[sourceChainId],
		onRampAddress:      onRampAddress,
		sourceToken:        usdcTokenAddress,
		attestationCache:   make(map[uint64]AttestationAttempt),
	}
}

func (usdc *OffchainTokenDataService) IsAttestationComplete(ctx context.Context, seqNum uint64) (success bool, attestation []byte, err error) {
	response, err := usdc.GetUpdatedAttestation(ctx, seqNum)
	if err != nil {
		return false, []byte{}, err
	}

	if response.Status == AttestationStatusSuccess {
		attestationBytes, err := hex.DecodeString(response.Attestation)
		if err != nil {
			return false, nil, err
		}
		return true, attestationBytes, nil
	}
	return false, []byte{}, nil
}

func (usdc *OffchainTokenDataService) GetUpdatedAttestation(ctx context.Context, seqNum uint64) (AttestationResponse, error) {
	// Try to get information from cache to reduce the number of database and external calls
	attestationAttempt, ok := usdc.attestationCache[seqNum]
	if ok && attestationAttempt.USDCAttestationResponse.Status == AttestationStatusSuccess {
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

func (usdc *OffchainTokenDataService) getAttemptInfoFromCCIPMessageId(ctx context.Context, seqNum uint64) (AttestationAttempt, error) {
	// Get the CCIP message send event from the log poller
	ccipSendRequests, err := usdc.sourceChainEvents.GetSendRequestsBetweenSeqNums(ctx, usdc.onRampAddress, seqNum, seqNum, 0)
	if err != nil {
		return AttestationAttempt{}, err
	}

	if len(ccipSendRequests) != 1 {
		return AttestationAttempt{}, fmt.Errorf("expected 1 CCIP send request, got %d", len(ccipSendRequests))
	}

	ccipSendRequest := ccipSendRequests[0]

	ccipSendTxHash := ccipSendRequest.Data.Raw.TxHash
	ccipSendLogIndex := int64(ccipSendRequest.Data.Raw.Index)

	// Get the USDC message body
	usdcMessageBody, err := usdc.sourceChainEvents.GetLastUSDCMessagePriorToLogIndexInTx(ctx, ccipSendLogIndex, ccipSendTxHash)
	if err != nil {
		return AttestationAttempt{}, err
	}

	return AttestationAttempt{
		USDCMessageBody:  usdcMessageBody,
		USDCMessageHash:  utils.Keccak256Fixed(usdcMessageBody),
		CCIPSendTxHash:   ccipSendTxHash,
		CCIPSendLogIndex: ccipSendLogIndex,
		USDCAttestationResponse: AttestationResponse{
			Status:      AttestationStatusUnchecked,
			Attestation: "",
		},
	}, nil
}

func (usdc *OffchainTokenDataService) callAttestationApi(ctx context.Context, usdcMessageHash [32]byte) (AttestationResponse, error) {
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

func (usdc *OffchainTokenDataService) GetSourceToken() common.Address {
	return usdc.sourceToken
}

func (usdc *OffchainTokenDataService) GetSourceLogPollerFilters() []logpoller.Filter {
	return []logpoller.Filter{
		{
			Name:      logpoller.FilterName(MESSAGE_SENT_FILTER_NAME, usdc.messageTransmitter.Hex()),
			EventSigs: []common.Hash{abihelpers.EventSignatures.USDCMessageSent},
			Addresses: []common.Address{usdc.messageTransmitter},
		},
	}
}
