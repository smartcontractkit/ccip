package ccip

import (
	"reflect"
	"testing"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_ge_onramp"
)

func TestOverheadGasGE(t *testing.T) {
	type test struct {
		merkleGasShare uint64
		geMsg          evm_2_evm_ge_onramp.GEEVM2EVMGEMessage
		want           uint64
	}

	// Only Data and TokensAndAmounts are used from the messages
	// And only the length is used so the contents doesn't matter.
	tests := []test{
		{
			geMsg: evm_2_evm_ge_onramp.GEEVM2EVMGEMessage{
				Data:             []byte{},
				TokensAndAmounts: []evm_2_evm_ge_onramp.CommonEVMTokenAndAmount{},
			},
			want: 27760,
		},
		{
			geMsg: evm_2_evm_ge_onramp.GEEVM2EVMGEMessage{
				Data: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
				TokensAndAmounts: []evm_2_evm_ge_onramp.CommonEVMTokenAndAmount{
					{},
				},
			},
			want: 71288,
		},
	}

	for _, tc := range tests {
		got := overheadGasGE(tc.geMsg)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestMaxGasOverHeadGasGE(t *testing.T) {
	type test struct {
		numMsgs int
		geMsg   evm_2_evm_ge_onramp.GEEVM2EVMGEMessage
		want    uint64
	}

	// Only Data and TokensAndAmounts are used from the messages
	// And only the length is used so the contents doesn't matter.
	tests := []test{
		{
			numMsgs: 6,
			geMsg: evm_2_evm_ge_onramp.GEEVM2EVMGEMessage{
				Data:             []byte{},
				TokensAndAmounts: []evm_2_evm_ge_onramp.CommonEVMTokenAndAmount{},
			},
			want: 37772,
		},
		{
			numMsgs: 3,
			geMsg: evm_2_evm_ge_onramp.GEEVM2EVMGEMessage{
				Data: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
				TokensAndAmounts: []evm_2_evm_ge_onramp.CommonEVMTokenAndAmount{
					{},
				},
			},
			want: 86705,
		},
	}

	for _, tc := range tests {
		got := maxGasOverHeadGasGE(tc.numMsgs, tc.geMsg)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
