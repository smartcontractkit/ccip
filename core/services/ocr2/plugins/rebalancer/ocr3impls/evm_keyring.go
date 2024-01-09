package ocr3impls

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"io"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

var _ ocr3types.OnchainKeyring[any] = &evmKeyring[any]{}

var curve = secp256k1.S256()

type evmKeyring[RI any] struct {
	privateKey ecdsa.PrivateKey
}

func NewEVMKeyring[RI any](material io.Reader) (*evmKeyring[RI], error) {
	ecdsaKey, err := ecdsa.GenerateKey(curve, material)
	if err != nil {
		return nil, err
	}
	return &evmKeyring[RI]{privateKey: *ecdsaKey}, nil
}

// XXX: PublicKey returns the address of the public key not the public key itself
func (ok *evmKeyring[RI]) PublicKey() ocrtypes.OnchainPublicKey {
	address := ok.signingAddress()
	return address[:]
}

func (ok *evmKeyring[RI]) reportToSigData(configDigest ocrtypes.ConfigDigest, seqNr [32]byte, rwi ocr3types.ReportWithInfo[RI]) []byte {
	// reportContext consists of:
	// reportContext[0]: ConfigDigest
	// reportContext[1]: 24 byte padding, 8 byte sequence number
	var rawReportContext [3][32]byte
	copy(rawReportContext[0][:], configDigest[:])
	copy(rawReportContext[1][:], seqNr[:])
	// TODO: ExtraHash unused in OCR3 (?)
	sigData := crypto.Keccak256(rwi.Report)
	sigData = append(sigData, rawReportContext[0][:]...)
	sigData = append(sigData, rawReportContext[1][:]...)
	sigData = append(sigData, rawReportContext[2][:]...)
	return crypto.Keccak256(sigData)
}

func (ok *evmKeyring[RI]) Sign(configDigest ocrtypes.ConfigDigest, seqNr uint64, rwi ocr3types.ReportWithInfo[RI]) ([]byte, error) {
	formattedSeqNr, err := formatSequenceNumber(seqNr)
	if err != nil {
		return nil, fmt.Errorf("failed to format sequence number: %w", err)
	}
	return crypto.Sign(ok.reportToSigData(configDigest, formattedSeqNr, rwi), &ok.privateKey)

}

func (ok *evmKeyring[RI]) Verify(publicKey ocrtypes.OnchainPublicKey, configDigest ocrtypes.ConfigDigest, seqNr uint64, rwi ocr3types.ReportWithInfo[RI], signature []byte) bool {
	formattedSeqNr, err := formatSequenceNumber(seqNr)
	if err != nil {
		return false
	}
	hash := ok.reportToSigData(configDigest, formattedSeqNr, rwi)
	authorPubkey, err := crypto.SigToPub(hash, signature)
	if err != nil {
		return false
	}
	authorAddress := crypto.PubkeyToAddress(*authorPubkey)
	return bytes.Equal(publicKey[:], authorAddress[:])
}

func (ok *evmKeyring[RI]) MaxSignatureLength() int {
	return 65
}

func (ok *evmKeyring[RI]) signingAddress() common.Address {
	return crypto.PubkeyToAddress(*(&ok.privateKey).Public().(*ecdsa.PublicKey))
}

func (ok *evmKeyring[RI]) Marshal() ([]byte, error) {
	return crypto.FromECDSA(&ok.privateKey), nil
}

func (ok *evmKeyring[RI]) Unmarshal(in []byte) error {
	privateKey, err := crypto.ToECDSA(in)
	if err != nil {
		return err
	}
	ok.privateKey = *privateKey
	return nil
}
