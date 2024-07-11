package ccipevm

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_multi_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
)

// ExecutePluginCodecV1 is a codec for encoding and decoding execute plugin reports.
// Compatible with:
// - "EVM2EVMMultiOffRamp 1.6.0-dev"
type ExecutePluginCodecV1 struct {
	executeReportEventInputs abi.Arguments
}

func NewExecutePluginCodecV1() *ExecutePluginCodecV1 {
	abiParsed, err := abi.JSON(strings.NewReader(evm_2_evm_multi_offramp.EVM2EVMMultiOffRampABI))
	if err != nil {
		panic(fmt.Errorf("parse multi offramp abi: %s", err))
	}
	eventInputs := abihelpers.MustGetMethodInputs("manuallyExecute", abiParsed)
	if len(eventInputs) == 0 {
		panic("no inputs found for method: manuallyExecute")
	}

	return &ExecutePluginCodecV1{
		executeReportEventInputs: eventInputs[:1],
	}
}

func (e *ExecutePluginCodecV1) Encode(ctx context.Context, report cciptypes.ExecutePluginReport) ([]byte, error) {
	evmReport := make([]evm_2_evm_multi_offramp.InternalExecutionReportSingleChain, 0, len(report.ChainReports))

	for _, chainReport := range report.ChainReports {
		if chainReport.ProofFlagBits.IsEmpty() {
			return nil, fmt.Errorf("proof flag bits are empty")
		}

		evmProofs := make([][32]byte, 0, len(chainReport.Proofs))
		for _, proof := range chainReport.Proofs {
			evmProofs = append(evmProofs, proof)
		}

		evmMessages := make([]evm_2_evm_multi_offramp.InternalAny2EVMRampMessage, 0, len(chainReport.Messages))
		for _, message := range chainReport.Messages {
			receiver := common.BytesToAddress(message.Receiver)

			tokenAmounts := make([]evm_2_evm_multi_offramp.InternalRampTokenAmount, 0, len(message.TokenAmounts))
			for _, tokenAmount := range message.TokenAmounts {
				if tokenAmount.Amount.IsEmpty() {
					return nil, fmt.Errorf("empty amount for token: %s", tokenAmount.DestTokenAddress)
				}

				tokenAmounts = append(tokenAmounts, evm_2_evm_multi_offramp.InternalRampTokenAmount{
					SourcePoolAddress: tokenAmount.SourcePoolAddress,
					DestTokenAddress:  tokenAmount.DestTokenAddress,
					ExtraData:         tokenAmount.ExtraData,
					Amount:            tokenAmount.Amount.Int,
				})
			}

			evmMessages = append(evmMessages, evm_2_evm_multi_offramp.InternalAny2EVMRampMessage{
				Header: evm_2_evm_multi_offramp.InternalRampMessageHeader{
					MessageId:           message.Header.MessageID,
					SourceChainSelector: uint64(message.Header.SourceChainSelector),
					DestChainSelector:   uint64(message.Header.DestChainSelector),
					SequenceNumber:      uint64(message.Header.SequenceNumber),
					Nonce:               message.Header.Nonce,
				},
				Sender:       message.Sender,
				Data:         message.Data,
				Receiver:     receiver,
				GasLimit:     big.NewInt(0), // todo
				TokenAmounts: tokenAmounts,
			})
		}

		evmChainReport := evm_2_evm_multi_offramp.InternalExecutionReportSingleChain{
			SourceChainSelector: uint64(chainReport.SourceChainSelector),
			Messages:            evmMessages,
			OffchainTokenData:   chainReport.OffchainTokenData,
			Proofs:              evmProofs,
			ProofFlagBits:       chainReport.ProofFlagBits.Int,
		}
		evmReport = append(evmReport, evmChainReport)
	}

	return e.executeReportEventInputs.PackValues([]interface{}{&evmReport})
}

func (e *ExecutePluginCodecV1) Decode(ctx context.Context, bytes []byte) (cciptypes.ExecutePluginReport, error) {
	panic("implement me")
}

// Ensure ExecutePluginCodec implements the ExecutePluginCodec interface
var _ cciptypes.ExecutePluginCodec = (*ExecutePluginCodecV1)(nil)
