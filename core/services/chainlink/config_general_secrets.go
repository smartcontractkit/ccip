package chainlink

import (
	"net/url"

	"github.com/smartcontractkit/chainlink/v2/core/config/toml"
)

func (g *generalConfig) DatabaseURL() url.URL {
	if g.secrets.Database.URL == nil {
		return url.URL{}
	}
	return *g.secrets.Database.URL.URL()
}

func (g *generalConfig) DatabaseBackupURL() *url.URL {
	return g.secrets.Database.BackupURL.URL()
}

func (g *generalConfig) ExplorerAccessKey() string {
	if g.secrets.Explorer.AccessKey == nil {
		return ""
	}
	return string(*g.secrets.Explorer.AccessKey)
}

func (g *generalConfig) ExplorerSecret() string {
	if g.secrets.Explorer.Secret == nil {
		return ""
	}
	return string(*g.secrets.Explorer.Secret)
}

func (g *generalConfig) LegacyGasStationAuthConfig() *toml.LegacyGasStationAuthConfig {
	if g.secrets.LegacyGasStation.AuthConfig == nil {
		return nil
	}
	return &toml.LegacyGasStationAuthConfig{
		ClientKey:         g.secrets.LegacyGasStation.AuthConfig.ClientKey,
		ClientCertificate: g.secrets.LegacyGasStation.AuthConfig.ClientCertificate,
	}
}
