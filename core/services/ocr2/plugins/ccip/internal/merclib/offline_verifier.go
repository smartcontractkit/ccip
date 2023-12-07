package merclib

import (
	"bytes"
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

var _ ReportVerifier = &offlineVerifier{}

type offlineVerifier struct {
	ocr2PubKeys []ocrtypes.OnchainPublicKey
	f           int
}

// NewOfflineVerifier returns a new ReportVerifier instance
// that verifies reports offline by utilizing the provided
// OCR2 onchain public keys.
// NOTE: verification here is EVM-only.
func NewOfflineVerifier(ocr2PubKeys []ocrtypes.OnchainPublicKey, f int) *offlineVerifier {
	return &offlineVerifier{ocr2PubKeys: ocr2PubKeys}
}

func (o *offlineVerifier) VerifyReports(ctx context.Context, signedReports [][]byte) error {
	for _, signedReport := range signedReports {
		if err := verifySingle(signedReport, o.f, o.ocr2PubKeys); err != nil {
			return err
		}
	}
	return nil
}

// verifySingle verifies a single signed report
// This is based off of the Verifier.sol contract:
// function verify(
//
//	  bytes calldata signedReport,
//	  address sender
//	) external override returns (bytes memory verifierResponse)
func verifySingle(signedReport []byte, f int, pubkeys []ocrtypes.OnchainPublicKey) error {
	// decode into raw report constituents
	fullReport, err := DecodeFullReport(signedReport)
	if err != nil {
		return fmt.Errorf("decoding full report: %w", err)
	}

	if err := validateReport(fullReport.RawRs, fullReport.RawSs, f); err != nil {
		return fmt.Errorf("validating report: %w", err)
	}

	var signedCount int
	for i := range fullReport.RawRs {
		combinedSig := combineSignature(fullReport.RawRs[i], fullReport.RawSs[i], fullReport.RawVs[i])
		pubkey := ecrecover(
			fullReport.ReportContext,
			fullReport.ReportBlob,
			combinedSig)
		if pubkey == nil {
			continue
		}
		for _, ocr2Pubkey := range pubkeys {
			if bytes.Equal(pubkey, ocr2Pubkey) {
				signedCount++
				break
			}
		}
	}

	if signedCount != (f + 1) {
		return fmt.Errorf("expected %d signatures, got %d", f+1, signedCount)
	}

	return nil
}

func validateReport(rs [][32]byte, ss [][32]byte, f int) error {
	expectedNumSignatures := f + 1
	if len(rs) != expectedNumSignatures {
		return fmt.Errorf("expected %d signatures, got %d", expectedNumSignatures, len(rs))
	}
	if len(rs) != len(ss) {
		return fmt.Errorf("mismatched rs and ss: len(rs)=%d, len(ss)=%d", len(rs), len(ss))
	}
	return nil
}

func reportToSigData(rawReportContext [3][32]byte, report ocrtypes.Report) []byte {
	sigData := crypto.Keccak256(report)
	sigData = append(sigData, rawReportContext[0][:]...)
	sigData = append(sigData, rawReportContext[1][:]...)
	sigData = append(sigData, rawReportContext[2][:]...)
	return crypto.Keccak256(sigData)
}

func ecrecover(
	reportCtx [3][32]byte,
	report ocrtypes.Report,
	signature []byte) []byte {
	hash := reportToSigData(reportCtx, report)
	authorPubkey, err := crypto.SigToPub(hash, signature)
	if err != nil {
		return nil
	}
	authorAddress := crypto.PubkeyToAddress(*authorPubkey)
	return authorAddress.Bytes()
}

func combineSignature(rs [32]byte, ss [32]byte, v byte) (sig []byte) {
	sig = append(sig, rs[:]...)
	sig = append(sig, ss[:]...)
	sig = append(sig, v)
	return sig
}
