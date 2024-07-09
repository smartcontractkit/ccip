package ocrimpls

import (
	"context"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

var _ ocr3types.ContractTransmitter[[]byte] = &contractTransmitter{}

type contractTransmitter struct {
	cw          types.ChainWriter
	fromAccount ocrtypes.Account
}

func NewContractTransmitter(
	cw types.ChainWriter,
	fromAccount ocrtypes.Account,
) ocr3types.ContractTransmitter[[]byte] {
	return &contractTransmitter{cw: cw, fromAccount: fromAccount}
}

// FromAccount implements ocr3types.ContractTransmitter.
func (c *contractTransmitter) FromAccount() (ocrtypes.Account, error) {
	return c.fromAccount, nil
}

// Transmit implements ocr3types.ContractTransmitter.
func (c *contractTransmitter) Transmit(
	ctx context.Context,
	configDigest ocrtypes.ConfigDigest,
	seqNr uint64,
	reportWithInfo ocr3types.ReportWithInfo[[]byte],
	aos []ocrtypes.AttributedOnchainSignature,
) error {
	panic("unimplemented")
}
