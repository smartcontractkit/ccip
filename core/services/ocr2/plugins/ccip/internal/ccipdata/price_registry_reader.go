package ccipdata

import cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"

const (
	COMMIT_PRICE_UPDATES = "Commit price updates"
	FEE_TOKEN_ADDED      = "Fee token added"
	FEE_TOKEN_REMOVED    = "Fee token removed"
	ExecPluginLabel      = "exec"
)

//go:generate mockery --quiet --name PriceRegistryReader --filename price_registry_reader_mock.go --case=underscore
type PriceRegistryReader interface {
	cciptypes.PriceRegistryReader
}