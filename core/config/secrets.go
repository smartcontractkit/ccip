package config

import (
	"github.com/smartcontractkit/chainlink/v2/core/services/legacygasstation/types"
	ocr2models "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/models"
)

type Secrets interface {
	MercuryCredentials(credName string) *ocr2models.MercuryCredentials
	LegacyGasStationAuthConfig() *types.AuthConfig
	VRFPassword() string
}
