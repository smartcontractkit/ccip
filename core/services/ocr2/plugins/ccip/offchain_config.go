package ccip

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

type OffchainConfig interface {
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

func (c CommitOffchainConfig) Validate() error {
	if c.SourceIncomingConfirmations == 0 {
		return errors.New("must set SourceIncomingConfirmations")
	}
	if c.DestIncomingConfirmations == 0 {
		return errors.New("must set DestIncomingConfirmations")
	}
	if c.FeeUpdateHeartBeat.Duration() == 0 {
		return errors.New("must set FeeUpdateHeartBeat")
	}
	if c.FeeUpdateDeviationPPB == 0 {
		return errors.New("must set FeeUpdateDeviationPPB")
	}
	if c.MaxGasPrice == 0 {
		return errors.New("must set MaxGasPrice")
	}
	if c.InflightCacheExpiry.Duration() == 0 {
		return errors.New("must set InflightCacheExpiry")
	}

	return nil
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

func (c ExecOffchainConfig) Validate() error {
	if c.SourceIncomingConfirmations == 0 {
		return errors.New("must set SourceIncomingConfirmations")
	}
	if c.DestIncomingConfirmations == 0 {
		return errors.New("must set DestIncomingConfirmations")
	}
	if c.BatchGasLimit == 0 {
		return errors.New("must set BatchGasLimit")
	}
	if c.RelativeBoostPerWaitHour == 0 {
		return errors.New("must set RelativeBoostPerWaitHour")
	}
	if c.MaxGasPrice == 0 {
		return errors.New("must set MaxGasPrice")
	}
	if c.InflightCacheExpiry.Duration() == 0 {
		return errors.New("must set InflightCacheExpiry")
	}
	if c.RootSnoozeTime.Duration() == 0 {
		return errors.New("must set RootSnoozeTime")
	}

	return nil
}

func DecodeOffchainConfig[T OffchainConfig](encodedConfig []byte) (T, error) {
	var result T
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

func EncodeOffchainConfig[T OffchainConfig](occ T) ([]byte, error) {
	return json.Marshal(occ)
}
