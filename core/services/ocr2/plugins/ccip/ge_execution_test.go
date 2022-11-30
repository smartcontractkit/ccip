package ccip

import (
	"reflect"
	"testing"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_ge_onramp"
)

func TestOverheadGasGE(t *testing.T) {
	type test struct {
		merkleGasShare uint64
		geMsg          evm_2_evm_ge_onramp.CCIPEVM2EVMGEMessage
		want           uint64
	}

	// Only Data and TokensAndAmounts are used from the messages
	// And only the length is used so the contents doesn't matter.
	tests := []test{
		{
			merkleGasShare: 0,
			geMsg: evm_2_evm_ge_onramp.CCIPEVM2EVMGEMessage{
				Data:             []byte{},
				TokensAndAmounts: []evm_2_evm_ge_onramp.CCIPEVMTokenAndAmount{},
			},
			want: 63260,
		},
		{
			merkleGasShare: 4_000,
			geMsg: evm_2_evm_ge_onramp.CCIPEVM2EVMGEMessage{
				Data: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
				TokensAndAmounts: []evm_2_evm_ge_onramp.CCIPEVMTokenAndAmount{
					{},
				},
			},
			want: 110788,
		},
	}

	for _, tc := range tests {
		got := overheadGasGE(tc.merkleGasShare, tc.geMsg)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestMaxGasOverHeadGasGE(t *testing.T) {
	type test struct {
		numMsgs int
		geMsg   evm_2_evm_ge_onramp.CCIPEVM2EVMGEMessage
		want    uint64
	}

	// Only Data and TokensAndAmounts are used from the messages
	// And only the length is used so the contents doesn't matter.
	tests := []test{
		{
			numMsgs: 6,
			geMsg: evm_2_evm_ge_onramp.CCIPEVM2EVMGEMessage{
				Data:             []byte{},
				TokensAndAmounts: []evm_2_evm_ge_onramp.CCIPEVMTokenAndAmount{},
			},
			want: 67356,
		},
		{
			numMsgs: 3,
			geMsg: evm_2_evm_ge_onramp.CCIPEVM2EVMGEMessage{
				Data: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
				TokensAndAmounts: []evm_2_evm_ge_onramp.CCIPEVMTokenAndAmount{
					{},
				},
			},
			want: 110372,
		},
	}

	for _, tc := range tests {
		got := maxGasOverHeadGasGE(tc.numMsgs, tc.geMsg)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
