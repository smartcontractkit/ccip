package ccip

import (
	"reflect"
	"testing"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_onramp"
)

func TestOverheadGas(t *testing.T) {
	type test struct {
		merkleGasShare uint64
		geMsg          evm_2_evm_onramp.InternalEVM2EVMMessage
		want           uint64
	}

	// Only Data and TokenAmounts are used from the messages
	// And only the length is used so the contents doesn't matter.
	tests := []test{
		{
			geMsg: evm_2_evm_onramp.InternalEVM2EVMMessage{
				Data:         []byte{},
				TokenAmounts: []evm_2_evm_onramp.ClientEVMTokenAmount{},
			},
			want: 27760,
		},
		{
			geMsg: evm_2_evm_onramp.InternalEVM2EVMMessage{
				Data: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
				TokenAmounts: []evm_2_evm_onramp.ClientEVMTokenAmount{
					{},
				},
			},
			want: 71288,
		},
	}

	for _, tc := range tests {
		got := overheadGas(tc.geMsg)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestMaxGasOverHeadGas(t *testing.T) {
	type test struct {
		numMsgs int
		geMsg   evm_2_evm_onramp.InternalEVM2EVMMessage
		want    uint64
	}

	// Only Data and TokenAmounts are used from the messages
	// And only the length is used so the contents doesn't matter.
	tests := []test{
		{
			numMsgs: 6,
			geMsg: evm_2_evm_onramp.InternalEVM2EVMMessage{
				Data:         []byte{},
				TokenAmounts: []evm_2_evm_onramp.ClientEVMTokenAmount{},
			},
			want: 31856,
		},
		{
			numMsgs: 3,
			geMsg: evm_2_evm_onramp.InternalEVM2EVMMessage{
				Data: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
				TokenAmounts: []evm_2_evm_onramp.ClientEVMTokenAmount{
					{},
				},
			},
			want: 74872,
		},
	}

	for _, tc := range tests {
		got := maxGasOverHeadGas(tc.numMsgs, tc.geMsg)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
