package customtokens

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type USDCService struct {
	attestationApi string
}

type USDCAttestationResponse struct {
	Status      USDCAttestationStatus `json:"status"`
	Attestation string                `json:"attestation"`
}

const (
	version         = "v1"
	attestationPath = "attestations"
)

type USDCAttestationStatus string

const (
	USDCAttestationStatusSuccess USDCAttestationStatus = "complete"
	USDCAttestationStatusPending USDCAttestationStatus = "pending_confirmations"
)

func NewUSDCService(usdcAttestationApi string) *USDCService {
	return &USDCService{attestationApi: usdcAttestationApi}
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
