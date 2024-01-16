package factory

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	mocks2 "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
)

func TestOnRamp(t *testing.T) {
	for _, versionStr := range []string{ccipdata.V1_0_0, ccipdata.V1_1_0, ccipdata.V1_2_0, ccipdata.V1_3_0} {
		lggr := logger.TestLogger(t)
		addr := utils.RandomAddress()
		lp := mocks2.NewLogPoller(t)

		sourceSelector := uint64(1000)
		destSelector := uint64(2000)

		expFilterNames := []string{
			logpoller.FilterName(ccipdata.COMMIT_CCIP_SENDS, addr),
		}
		versionFinder := newMockVersionFinder(ccipconfig.EVM2EVMOnRamp, *semver.MustParse(versionStr), nil)

		lp.On("RegisterFilter", mock.Anything).Return(nil).Times(len(expFilterNames))
		_, err := NewOnRampReader(lggr, versionFinder, sourceSelector, destSelector, addr, lp, nil)
		assert.NoError(t, err)

		for _, f := range expFilterNames {
			lp.On("UnregisterFilter", f).Return(nil)
		}
		err = CloseOnRampReader(lggr, versionFinder, sourceSelector, destSelector, addr, lp, nil)
		assert.NoError(t, err)
	}
}
