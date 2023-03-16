package ccip

import (
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_onramp"
)

func TestOverheadGas(t *testing.T) {
	// Only Data and TokenAmounts are used from the messages
	// And only the length is used so the contents doesn't matter.
	tests := []struct {
		geMsg evm_2_evm_onramp.InternalEVM2EVMMessage
		want  uint64
	}{
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
	// Only Data and TokenAmounts are used from the messages
	// And only the length is used so the contents doesn't matter.
	tests := []struct {
		numMsgs int
		geMsg   evm_2_evm_onramp.InternalEVM2EVMMessage
		want    uint64
	}{
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

func TestComputeExecCost(t *testing.T) {
	tests := []struct {
		name            string
		gasLimit        *big.Int
		execGasEstimate *big.Int
		tokenPriceUSD   *big.Int
		execCostUsd     *big.Int
	}{
		{
			"happy flow",
			big.NewInt(3_000_000),
			big.NewInt(2e10),
			big.NewInt(6e18),
			big.NewInt(384e15),
		},
		{
			"low usd price",
			big.NewInt(3_000_000),
			big.NewInt(2e10),
			big.NewInt(6e15),
			big.NewInt(384e12),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msg := &evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested{
				Message: evm_2_evm_onramp.InternalEVM2EVMMessage{
					GasLimit: tc.gasLimit,
				},
			}
			execCostUsd := computeExecCost(msg, tc.execGasEstimate, tc.tokenPriceUSD)
			require.Equal(t, tc.execCostUsd, execCostUsd)
		})
	}
}

func TestWaitBoostedFee(t *testing.T) {
	tests := []struct {
		name         string
		sendTimeDiff time.Duration
		fee          *big.Int
		boostedFee   *big.Int
	}{
		{
			"boosted fee 15hr",
			time.Hour * 14,
			big.NewInt(6e18),
			big.NewInt(0).Mul(big.NewInt(2), big.NewInt(594e16)),
		},
		{
			"boosted fee 5hr",
			time.Hour * 5,
			big.NewInt(6e18),
			big.NewInt(0).Add(big.NewInt(8e18), big.NewInt(1e17+1024)),
		},
		{
			"boosted fee 1hr",
			time.Hour * 1,
			big.NewInt(6e18),
			big.NewInt(642e16),
		},
		{
			"boosted fee 10s",
			time.Second * 10,
			big.NewInt(6e18),
			big.NewInt(6e18),
		},
		{
			"boosted fee 2m",
			time.Minute * 2,
			big.NewInt(6e18),
			big.NewInt(6e18),
		},
		{
			"boosted fee 25m",
			time.Minute * 25,
			big.NewInt(6e18),
			big.NewInt(6e18),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			boosted := waitBoostedFee(time.Now().Add(-tc.sendTimeDiff), tc.fee)
			require.Equal(t, tc.boostedFee, boosted)
		})
	}
}
