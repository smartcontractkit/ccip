package ccipdata

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
)

func TestExecutionReportEncoding(t *testing.T) {
	// Note could consider some fancier testing here (fuzz/property)
	// but I think that would essentially be testing geth's abi library
	// as our encode/decode is a thin wrapper around that.
	report := evm_2_evm_offramp.InternalExecutionReport{
		Messages:          []evm_2_evm_offramp.InternalEVM2EVMMessage{},
		OffchainTokenData: [][][]byte{{}},
		Proofs:            [][32]byte{testutils.Random32Byte()},
		ProofFlagBits:     big.NewInt(133),
	}
	encodeExecutionReport, err := EncodeExecutionReport(evm_2_evm_offramp.InternalExecutionReport{
		Messages:          report.Messages,
		OffchainTokenData: report.OffchainTokenData,
		Proofs:            report.Proofs,
		ProofFlagBits:     report.ProofFlagBits,
	})
	require.NoError(t, err)
	decodeCommitReport, err := DecodeExecutionReport(encodeExecutionReport)
	require.NoError(t, err)
	require.Equal(t, report, decodeCommitReport)
}
