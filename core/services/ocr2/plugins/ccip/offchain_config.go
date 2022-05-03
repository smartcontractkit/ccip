package ccip

import (
	"encoding/json"
)

type OffchainConfig struct {
	SourceIncomingConfirmations uint32
	DestIncomingConfirmations   uint32
}

func Decode(encodedConfig []byte) (OffchainConfig, error) {
	var result OffchainConfig
	err := json.Unmarshal(encodedConfig, &result)
	return result, err
}

func (occ OffchainConfig) Encode() ([]byte, error) {
	return json.Marshal(occ)
}
