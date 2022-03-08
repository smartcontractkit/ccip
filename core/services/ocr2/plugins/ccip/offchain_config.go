package ccip

import (
	"context"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	types2 "github.com/smartcontractkit/libocr/offchainreporting2/types"
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

func GetOffchainConfig(t types2.ContractConfigTracker) (OffchainConfig, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	changedInBlock, _, err := t.LatestConfigDetails(ctx)
	if err != nil {
		return OffchainConfig{}, errors.Wrap(err, "could not get block number for latest config change")
	}
	config, err := t.LatestConfig(ctx, changedInBlock)
	if err != nil {
		return OffchainConfig{}, errors.Wrap(err, "could not get latest config")
	}
	publicConfig, err := confighelper.PublicConfigFromContractConfig(false, config)
	if err != nil {
		return OffchainConfig{}, errors.Wrap(err, "could not parse latest config")
	}
	ccipConfig, err := Decode(publicConfig.ReportingPluginConfig)
	if err != nil {
		return OffchainConfig{}, errors.Wrap(err, "could not decode latest config")
	}
	return ccipConfig, nil
}
