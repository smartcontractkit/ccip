package ccipdata

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client/mocks"
	lpmocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
)

func TestExecutionReportEncoding(t *testing.T) {
	// Note could consider some fancier testing here (fuzz/property)
	// but I think that would essentially be testing geth's abi library
	// as our encode/decode is a thin wrapper around that.
	report := ExecReport{
		Messages:          []internal.EVM2EVMMessage{},
		OffchainTokenData: [][][]byte{{}},
		Proofs:            [][32]byte{testutils.Random32Byte()},
		ProofFlagBits:     big.NewInt(133),
	}
	offRamp := &OffRampV1_0_0{}
	encodeExecutionReport, err := offRamp.EncodeExecutionReport(report)
	require.NoError(t, err)
	decodeCommitReport, err := offRamp.DecodeExecutionReport(encodeExecutionReport)
	require.NoError(t, err)
	require.Equal(t, report, decodeCommitReport)
}

func TestOffRampFilters(t *testing.T) {
	assertFilterRegistration(t, new(lpmocks.LogPoller), func(lp *lpmocks.LogPoller, addr common.Address) Closer {
		c, err := NewOffRampV1_0_0(logger.TestLogger(t), addr, new(mocks.Client), lp, nil)
		require.NoError(t, err)
		return c
	}, 3)
	assertFilterRegistration(t, new(lpmocks.LogPoller), func(lp *lpmocks.LogPoller, addr common.Address) Closer {
		c, err := NewOffRampV1_2_0(logger.TestLogger(t), addr, new(mocks.Client), lp, nil)
		require.NoError(t, err)
		return c
	}, 3)
}
