package abihelpers

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/test-go/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
)

func TestProofFlagToBits(t *testing.T) {
	genFlags := func(indexesSet []int, size int) []bool {
		bools := make([]bool, size)
		for _, indexSet := range indexesSet {
			bools[indexSet] = true
		}
		return bools
	}
	tt := []struct {
		flags    []bool
		expected *big.Int
	}{
		{
			[]bool{true, false, true},
			big.NewInt(5),
		},
		{
			[]bool{true, true, false}, // Note the bits are reversed, slightly easier to implement.
			big.NewInt(3),
		},
		{
			[]bool{false, true, true},
			big.NewInt(6),
		},
		{
			[]bool{false, false, false},
			big.NewInt(0),
		},
		{
			[]bool{true, true, true},
			big.NewInt(7),
		},
		{
			genFlags([]int{266}, 300),
			big.NewInt(0).SetBit(big.NewInt(0), 266, 1),
		},
	}
	for _, tc := range tt {
		tc := tc
		a := ProofFlagsToBits(tc.flags)
		assert.Equal(t, tc.expected.String(), a.String())
	}
}

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
