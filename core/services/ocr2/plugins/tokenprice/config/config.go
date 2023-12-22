package config

import (
	"errors"
	"fmt"
	"math/big"
	"net/url"
)

type PluginConfig struct {
	CCIPPriceServiceURL string `json:"ccipPriceServiceURL"`
	DestinationChainID  string `json:"destinationChainID"`
}

func ValidatePluginConfig(config PluginConfig) error {
	if config.CCIPPriceServiceURL == "" {
		return errors.New("ccipPriceServiceURL must not be empty in config")
	}
	_, err := url.Parse(config.CCIPPriceServiceURL)
	if err != nil {
		return fmt.Errorf("error parsing url: %w", err)
	}
	if config.DestinationChainID == "" {
		return errors.New("destinationChainID must not be empty in config")
	}
	// parse as big int
	_, ok := big.NewInt(0).SetString(config.DestinationChainID, 10)
	if !ok {
		return errors.New("error parsing destinationChainID as big int")
	}
	return nil
}
