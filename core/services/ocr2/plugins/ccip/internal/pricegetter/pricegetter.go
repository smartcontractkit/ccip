package pricegetter

import "github.com/smartcontractkit/chainlink-common/pkg/types/cciptypes"

//go:generate mockery --quiet --name PriceGetter --output . --filename mock.go --inpackage --case=underscore
type PriceGetter interface {
	cciptypes.PriceGetter
}
