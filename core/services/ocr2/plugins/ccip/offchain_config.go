package ccip

import (
	"encoding/json"

	"github.com/smartcontractkit/chainlink/core/store/models"
)

type OffchainConfig struct {
	SourceIncomingConfirmations uint32
	DestIncomingConfirmations   uint32
	FeeUpdateHeartBeat          models.Duration
	FeeUpdateDeviationPPB       uint32
}

func Decode(encodedConfig []byte) (OffchainConfig, error) {
	var result OffchainConfig
	err := json.Unmarshal(encodedConfig, &result)
	return result, err
}

func (occ OffchainConfig) Encode() ([]byte, error) {
	return json.Marshal(occ)
}
