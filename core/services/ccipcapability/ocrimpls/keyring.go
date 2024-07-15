package ocrimpls

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

var _ ocr3types.OnchainKeyring[[]byte] = &oc3Keyring[[]byte]{}

type oc3Keyring[RI any] struct {
	core types.OnchainKeyring
	lggr logger.Logger
}

func NewOnchainKeyring[RI any](keyring types.OnchainKeyring, lggr logger.Logger) *oc3Keyring[RI] {
	return &oc3Keyring[RI]{
		core: keyring,
		lggr: lggr.Named("OCR3Keyring"),
	}
}

func (w *oc3Keyring[RI]) PublicKey() types.OnchainPublicKey {
	return w.core.PublicKey()
}

func (w *oc3Keyring[RI]) MaxSignatureLength() int {
	return w.core.MaxSignatureLength()
}

func (w *oc3Keyring[RI]) Sign(configDigest types.ConfigDigest, seqNr uint64, r ocr3types.ReportWithInfo[RI]) (signature []byte, err error) {
	epoch, round := uint64ToUint32AndUint8(seqNr)
	rCtx := types.ReportContext{
		ReportTimestamp: types.ReportTimestamp{
			ConfigDigest: configDigest,
			Epoch:        epoch,
			Round:        round,
		},
	}

	w.lggr.Debugw("signing report", "configDigest", configDigest.Hex(), "seqNr", seqNr, "report", hexutil.Encode(r.Report))

	return w.core.Sign(rCtx, r.Report)
}

func (w *oc3Keyring[RI]) Verify(key types.OnchainPublicKey, configDigest types.ConfigDigest, seqNr uint64, r ocr3types.ReportWithInfo[RI], signature []byte) bool {
	epoch, round := uint64ToUint32AndUint8(seqNr)
	rCtx := types.ReportContext{
		ReportTimestamp: types.ReportTimestamp{
			ConfigDigest: configDigest,
			Epoch:        epoch,
			Round:        round,
		},
	}

	w.lggr.Debugw("verifying report", "configDigest", configDigest.Hex(), "seqNr", seqNr, "report", hexutil.Encode(r.Report))

	return w.core.Verify(key, rCtx, r.Report, signature)
}
