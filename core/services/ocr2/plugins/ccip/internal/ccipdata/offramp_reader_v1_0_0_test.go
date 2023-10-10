package ccipdata

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

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

	lp := lpmocks.NewLogPoller(t)
	lp.On("RegisterFilter", mock.Anything).Return(nil)
	offRamp, err := NewOffRampV1_0_0(logger.TestLogger(t), randomAddress(), nil, lp, nil)
	require.NoError(t, err)

	encodeExecutionReport, err := offRamp.EncodeExecutionReport(report)
	require.NoError(t, err)
	decodeCommitReport, err := offRamp.DecodeExecutionReport(encodeExecutionReport)
	require.NoError(t, err)
	require.Equal(t, report.Proofs, decodeCommitReport.Proofs)
	// require.Equal(t, report, decodeCommitReport) // TODO: fails because some fields are not supported on v1_0_0
}
