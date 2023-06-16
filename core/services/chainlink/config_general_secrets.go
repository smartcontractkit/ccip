package chainlink

import (
	"net/url"

	lgsconfig "github.com/smartcontractkit/chainlink/v2/core/services/legacygasstation/types/config"
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

func (g *generalConfig) LegacyGasStationAuthConfig() *lgsconfig.AuthConfig {
	if g.secrets.LegacyGasStation.AuthConfig == nil {
		return nil
	}
	return &lgsconfig.AuthConfig{
		ClientKey:         string(g.secrets.LegacyGasStation.AuthConfig.ClientKey),
		ClientCertificate: string(g.secrets.LegacyGasStation.AuthConfig.ClientCertificate),
	}
}
