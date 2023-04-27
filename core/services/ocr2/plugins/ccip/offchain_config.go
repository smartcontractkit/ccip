package ccip

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

type ValidatedOffchainConfig interface {
	Validate() error
}

type CommitOffchainConfig struct {
	SourceIncomingConfirmations uint32
	DestIncomingConfirmations   uint32
	FeeUpdateHeartBeat          models.Duration
	FeeUpdateDeviationPPB       uint32
	MaxGasPrice                 uint64
	InflightCacheExpiry         models.Duration
}

type ExecOffchainConfig struct {
	SourceIncomingConfirmations uint32
	DestIncomingConfirmations   uint32
	BatchGasLimit               uint32
	RelativeBoostPerWaitHour    float64
	MaxGasPrice                 uint64
	InflightCacheExpiry         models.Duration
	RootSnoozeTime              models.Duration
}

func (c CommitOffchainConfig) Validate() error {
	if c.InflightCacheExpiry.Duration() == 0 {
		return errors.New("InflightCacheExpiry not set")
	}
	return nil
}

func (c ExecOffchainConfig) Validate() error {
	if c.InflightCacheExpiry.Duration() == 0 {
		return errors.New("InflightCacheExpiry not set")
	}
	if c.RootSnoozeTime.Duration() == 0 {
		return errors.New("RootSnoozeTime not set")
	}
	return nil
}

type OffchainConfig interface {
	ValidatedOffchainConfig
}

func DecodeOffchainConfig[OCC OffchainConfig](encodedConfig []byte) (OCC, error) {
	var result OCC
	err := json.Unmarshal(encodedConfig, &result)
	if err != nil {
		return result, err
	}
	err = result.Validate()
	if err != nil {
		return result, err
	}
	return result, nil
}

func EncodeOffchainConfig[OCC OffchainConfig](occ OCC) ([]byte, error) {
	return json.Marshal(occ)
}
