package ccip

import (
	"encoding/json"

	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

type CommitOffchainConfig struct {
	SourceIncomingConfirmations uint32
	DestIncomingConfirmations   uint32
	FeeUpdateHeartBeat          models.Duration
	FeeUpdateDeviationPPB       uint32
	MaxGasPrice                 uint64
}

type ExecOffchainConfig struct {
	SourceIncomingConfirmations uint32
	DestIncomingConfirmations   uint32
	BatchGasLimit               uint32
	RelativeBoostPerWaitHour    float64
	MaxGasPrice                 uint64
}

type OffchainConfig interface {
	ExecOffchainConfig | CommitOffchainConfig
}

func DecodeOffchainConfig[OCC OffchainConfig](encodedConfig []byte) (OCC, error) {
	var result OCC
	err := json.Unmarshal(encodedConfig, &result)
	return result, err
}

func EncodeOffchainConfig[OCC OffchainConfig](occ OCC) ([]byte, error) {
	return json.Marshal(occ)
}
