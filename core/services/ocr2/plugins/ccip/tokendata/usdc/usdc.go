package usdc

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/tokendata"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

type TokenDataReader struct {
	sourceChainEvents  ccipdata.Reader
	attestationApi     *url.URL
	messageTransmitter common.Address
	sourceToken        common.Address
	onRampAddress      common.Address

	// Cache of sequence number -> usdc message body
	usdcMessageHashCache      map[uint64][32]byte
	usdcMessageHashCacheMutex sync.Mutex
}

type attestationResponse struct {
	Status      AttestationStatus `json:"status"`
	Attestation string            `json:"attestation"`
}

// Hard coded mapping of chain id to USDC token addresses
// Will be removed in favour of more flexible solution.
var tokenMapping = map[uint64]common.Address{
	// Mainnet
	1:     common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"), // Ethereum
	10:    common.HexToAddress("0x0b2c639c533813f4aa9d7837caf62653d097ff85"), // Optimism
	42161: common.HexToAddress("0xaf88d065e77c8cc2239327c5edb3a432268e5831"), // Arbitrum
	43114: common.HexToAddress("0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e"), // Avalanche

	// Testnets
	5:      common.HexToAddress("0xca6b4c00831ffb77afe22e734a6101b268b7fcbe"), // Goerli
	420:    common.HexToAddress("0xe05606174bac4A6364B31bd0eCA4bf4dD368f8C6"), // Optimism Goerli
	43113:  common.HexToAddress("0x5425890298aed601595a70ab815c96711a31bc65"), // Avalanche Fuji
	421613: common.HexToAddress("0xfd064A18f3BF249cf1f87FC203E90D8f650f2d63"), // Arbitrum Goerli
}

func GetUSDCTokenAddress(chain uint64) (common.Address, error) {
	if tokenAddress, ok := tokenMapping[chain]; ok {
		return tokenAddress, nil
	}
	return common.Address{}, errors.New("token not found")
}

var messageTransmitterMapping = map[uint64]common.Address{
	// Mainnet
	1:     common.HexToAddress("0x0a992d191deec32afe36203ad87d7d289a738f81"), // Ethereum
	10:    common.HexToAddress("0x4d41f22c5a0e5c74090899e5a8fb597a8842b3e8"), // Optimism
	42161: common.HexToAddress("0xC30362313FBBA5cf9163F0bb16a0e01f01A896ca"), // Arbitrum
	43114: common.HexToAddress("0x8186359af5f57fbb40c6b14a588d2a59c0c29880"), // Avalanche

	// Testnets
	5:      common.HexToAddress("0xca6b4c00831ffb77afe22e734a6101b268b7fcbe"), // Goerli
	420:    common.HexToAddress("0x9ff9a4da6f2157a9c82ce756f8fd7e0d75be8895"), // Optimism Goerli
	43113:  common.HexToAddress("0xa9fb1b3009dcb79e2fe346c16a604b8fa8ae0a79"), // Avalanche Fuji
	421613: common.HexToAddress("0x109bc137cb64eab7c0b1dddd1edf341467dc2d35"), // Arbitrum Goerli
}

func GetUSDCMessageTransmitterAddress(chain uint64) (common.Address, error) {
	if transmitterAddress, ok := messageTransmitterMapping[chain]; ok {
		return transmitterAddress, nil
	}
	return common.Address{}, errors.New("usdc transmitter not found")
}

const (
	version                  = "v1"
	attestationPath          = "attestations"
	MESSAGE_SENT_FILTER_NAME = "USDC message sent"
)

type AttestationStatus string

const (
	AttestationStatusSuccess AttestationStatus = "complete"
	AttestationStatusPending AttestationStatus = "pending_confirmations"
)

var _ tokendata.Reader = &TokenDataReader{}

func NewUSDCTokenDataReader(sourceChainEvents ccipdata.Reader, usdcTokenAddress, onRampAddress common.Address, usdcAttestationApi *url.URL, sourceChainId uint64) *TokenDataReader {
	return &TokenDataReader{
		sourceChainEvents:    sourceChainEvents,
		attestationApi:       usdcAttestationApi,
		messageTransmitter:   messageTransmitterMapping[sourceChainId],
		onRampAddress:        onRampAddress,
		sourceToken:          usdcTokenAddress,
		usdcMessageHashCache: make(map[uint64][32]byte),
	}
}

func (s *TokenDataReader) ReadTokenData(ctx context.Context, msg internal.EVM2EVMOnRampCCIPSendRequestedWithMeta) (attestation []byte, err error) {
	response, err := s.getUpdatedAttestation(ctx, msg)
	if err != nil {
		return []byte{}, err
	}

	if response.Status == AttestationStatusSuccess {
		attestationBytes, err := hex.DecodeString(response.Attestation)
		if err != nil {
			return nil, fmt.Errorf("decode response attestation: %w", err)
		}
		return attestationBytes, nil
	}
	return []byte{}, tokendata.ErrNotReady
}

func (s *TokenDataReader) getUpdatedAttestation(ctx context.Context, msg internal.EVM2EVMOnRampCCIPSendRequestedWithMeta) (attestationResponse, error) {
	messageBody, err := s.getUSDCMessageBody(ctx, msg)
	if err != nil {
		return attestationResponse{}, errors.Wrap(err, "failed getting the USDC message body")
	}

	response, err := s.callAttestationApi(ctx, messageBody)
	if err != nil {
		return attestationResponse{}, errors.Wrap(err, "failed calling usdc attestation API ")
	}

	return response, nil
}

func (s *TokenDataReader) getUSDCMessageBody(ctx context.Context, msg internal.EVM2EVMOnRampCCIPSendRequestedWithMeta) ([32]byte, error) {
	s.usdcMessageHashCacheMutex.Lock()
	defer s.usdcMessageHashCacheMutex.Unlock()

	if body, ok := s.usdcMessageHashCache[msg.SequenceNumber]; ok {
		return body, nil
	}

	usdcMessageBody, err := s.sourceChainEvents.GetLastUSDCMessagePriorToLogIndexInTx(ctx, int64(msg.LogIndex), msg.TxHash)
	if err != nil {
		return [32]byte{}, err
	}

	msgBodyHash := utils.Keccak256Fixed(usdcMessageBody)

	// Save the attempt in the cache in case the external call fails
	s.usdcMessageHashCache[msg.SequenceNumber] = msgBodyHash
	return msgBodyHash, nil
}

func (s *TokenDataReader) callAttestationApi(ctx context.Context, usdcMessageHash [32]byte) (attestationResponse, error) {
	fullAttestationUrl := fmt.Sprintf("%s/%s/%s/0x%x", s.attestationApi, version, attestationPath, usdcMessageHash)
	req, err := http.NewRequestWithContext(ctx, "GET", fullAttestationUrl, nil)
	if err != nil {
		return attestationResponse{}, err
	}
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return attestationResponse{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return attestationResponse{}, err
	}

	var response attestationResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return attestationResponse{}, err
	}
	return response, nil
}

func (s *TokenDataReader) GetSourceLogPollerFilters() []logpoller.Filter {
	return []logpoller.Filter{
		{
			Name:      logpoller.FilterName(MESSAGE_SENT_FILTER_NAME, s.messageTransmitter.Hex()),
			EventSigs: []common.Hash{abihelpers.EventSignatures.USDCMessageSent},
			Addresses: []common.Address{s.messageTransmitter},
		},
	}
}
