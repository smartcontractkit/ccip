package ccipevm

import (
	"math/rand"
	"testing"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/stretchr/testify/assert"
)

var randomExecuteReport = func(t *testing.T) cciptypes.ExecutePluginReport {
	const numChainReports = 100
	const msgsPerReport = 50
	const numTokensPerMsg = 20

	chainReports := make([]cciptypes.ExecutePluginReportSingleChain, numChainReports)
	for i := 0; i < numChainReports; i++ {
		reportMessages := make([]cciptypes.Message, msgsPerReport)
		for j := 0; j < msgsPerReport; j++ {
			data, err := cciptypes.NewBytesFromString(utils.RandomAddress().String())
			assert.NoError(t, err)

			tokenAmounts := make([]cciptypes.RampTokenAmount, numTokensPerMsg)
			for z := 0; z < numTokensPerMsg; z++ {
				tokenAmounts[z] = cciptypes.RampTokenAmount{
					SourcePoolAddress: utils.RandomAddress().Bytes(),
					DestTokenAddress:  utils.RandomAddress().Bytes(),
					ExtraData:         data,
					Amount:            cciptypes.NewBigInt(utils.RandUint256()),
				}
			}

			reportMessages[j] = cciptypes.Message{
				Header: cciptypes.RampMessageHeader{
					MessageID:           utils.RandomBytes32(),
					SourceChainSelector: cciptypes.ChainSelector(rand.Uint64()),
					DestChainSelector:   cciptypes.ChainSelector(rand.Uint64()),
					SequenceNumber:      cciptypes.SeqNum(rand.Uint64()),
					Nonce:               rand.Uint64(),
					MsgHash:             utils.RandomBytes32(),
					OnRamp:              utils.RandomAddress().Bytes(),
				},
				Sender:         utils.RandomAddress().Bytes(),
				Data:           data,
				Receiver:       utils.RandomAddress().Bytes(),
				ExtraArgs:      data,
				FeeToken:       utils.RandomAddress().Bytes(),
				FeeTokenAmount: cciptypes.NewBigInt(utils.RandUint256()),
				TokenAmounts:   tokenAmounts,
			}
		}

		tokenData := make([][][]byte, numTokensPerMsg)

		chainReports[i] = cciptypes.ExecutePluginReportSingleChain{
			SourceChainSelector: cciptypes.ChainSelector(rand.Uint64()),
			Messages:            reportMessages,
			OffchainTokenData:   tokenData,
			Proofs:              []cciptypes.Bytes32{utils.RandomBytes32(), utils.RandomBytes32()},
			ProofFlagBits:       cciptypes.NewBigInt(utils.RandUint256()),
		}
	}

	return cciptypes.ExecutePluginReport{ChainReports: chainReports}
}

func TestExecutePluginCodecV1(t *testing.T) {
	testCases := []struct {
		name   string
		report func(report cciptypes.ExecutePluginReport) cciptypes.ExecutePluginReport
		expErr bool
	}{
		{
			name:   "base report",
			report: func(report cciptypes.ExecutePluginReport) cciptypes.ExecutePluginReport { return report },
			expErr: false,
		},
		{
			name: "reports have empty msgs",
			report: func(report cciptypes.ExecutePluginReport) cciptypes.ExecutePluginReport {
				report.ChainReports[0].Messages = nil
				report.ChainReports[4].Messages = []cciptypes.Message{}
				return report
			},
			expErr: false,
		},
		{
			name: "reports have empty offchain token data",
			report: func(report cciptypes.ExecutePluginReport) cciptypes.ExecutePluginReport {
				report.ChainReports[0].OffchainTokenData = [][][]byte{}
				report.ChainReports[4].OffchainTokenData[1] = [][]byte{}
				return report
			},
			expErr: false,
		},
	}

	ctx := testutils.Context(t)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			codec := NewExecutePluginCodecV1()
			report := tc.report(randomExecuteReport(t))
			bytes, err := codec.Encode(ctx, report)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			decodedReport, err := codec.Decode(ctx, bytes)
			assert.NoError(t, err)
			assert.Equal(t, report, decodedReport)
		})
	}
}
