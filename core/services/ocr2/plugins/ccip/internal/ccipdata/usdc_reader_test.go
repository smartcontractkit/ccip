package ccipdata_test

import (
	"testing"

	lpmocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
)

func TestUSDCReaderFilters(t *testing.T) {
	ccipdata.AssertFilterRegistration(t, new(lpmocks.LogPoller), 1)
}
