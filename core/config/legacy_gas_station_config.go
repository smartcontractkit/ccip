package config

import "github.com/smartcontractkit/chainlink/v2/core/services/legacygasstation/types/config"

type LegacyGasStation interface {
	AuthConfig() *config.AuthConfig
}
