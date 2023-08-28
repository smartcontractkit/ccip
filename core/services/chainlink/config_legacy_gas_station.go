package chainlink

import (
	"github.com/smartcontractkit/chainlink/v2/core/config/toml"
	"github.com/smartcontractkit/chainlink/v2/core/services/legacygasstation/types/config"
)

type legacyGasStationConfig struct {
	s toml.LegacyGasStationSecrets
}

func (l *legacyGasStationConfig) AuthConfig() *config.AuthConfig {
	if l.s.AuthConfig == nil {
		return nil
	}
	if len(l.s.AuthConfig.ClientKey) > 0 && len(l.s.AuthConfig.ClientCertificate) > 0 {
		return &config.AuthConfig{
			ClientKey:         string(l.s.AuthConfig.ClientKey),
			ClientCertificate: string(l.s.AuthConfig.ClientCertificate),
		}
	}
	return nil
}

type AuthConfig struct {
	ClientKey         string
	ClientCertificate string
}
