package abihelpers

import (
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/gethwrappers2/ocr2aggregator"

	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

// TODO: Deprecate in favour of version specific types
var EventSignatures struct {
	// CommitStore
	ReportAccepted common.Hash
	// OffRamp
	ExecutionStateChanged common.Hash
	PoolAdded             common.Hash
	PoolRemoved           common.Hash

	// PriceRegistry
	UsdPerUnitGasUpdated common.Hash
	UsdPerTokenUpdated   common.Hash
	FeeTokenAdded        common.Hash
	FeeTokenRemoved      common.Hash

	// offset || priceUpdatesOffset || minSeqNum || maxSeqNum || merkleRoot
	ReportAcceptedMaxSequenceNumberWord int
	// sig || seqNum || messageId || ...
	ExecutionStateChangedSequenceNumberIndex int
}

var (
	MessageArgs         abi.Arguments
	TokenAmountsArgs    abi.Arguments
	CommitReportArgs    abi.Arguments
	ExecutionReportArgs abi.Arguments
)

func GetIDOrPanic(name string, abi2 abi.ABI) common.Hash {
	event, ok := abi2.Events[name]
	if !ok {
		panic(fmt.Sprintf("missing event %s", name))
	}
	return event.ID
}

func init() {
	//commitStoreABI, err := abi.JSON(strings.NewReader(commit_store.CommitStoreABI))
	//if err != nil {
	//	panic(err)
	//}
	//EventSignatures.ReportAccepted = GetIDOrPanic("ReportAccepted", commitStoreABI)
	//EventSignatures.ReportAcceptedMaxSequenceNumberWord = 3

}

// ProofFlagsToBits transforms a list of boolean proof flags to a *big.Int
// encoded number.
func ProofFlagsToBits(proofFlags []bool) *big.Int {
	encodedFlags := big.NewInt(0)
	for i := 0; i < len(proofFlags); i++ {
		if proofFlags[i] {
			encodedFlags.SetBit(encodedFlags, i, 1)
		}
	}
	return encodedFlags
}

type AbiDefined interface {
	AbiString() string
}

type AbiDefinedValid interface {
	AbiDefined
	Validate() error
}

func EncodeAbiStruct[T AbiDefined](decoded T) ([]byte, error) {
	return utils.ABIEncode(decoded.AbiString(), decoded)
}

func DecodeAbiStruct[T AbiDefinedValid](encoded []byte) (T, error) {
	var empty T

	decoded, err := utils.ABIDecode(empty.AbiString(), encoded)
	if err != nil {
		return empty, err
	}

	converted := abi.ConvertType(decoded[0], &empty)
	if casted, ok := converted.(*T); ok {
		return *casted, (*casted).Validate()
	}
	return empty, fmt.Errorf("can't cast from %T to %T", converted, empty)
}

func EvmWord(i uint64) common.Hash {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return common.BigToHash(big.NewInt(0).SetBytes(b))
}

func DecodeOCR2Config(encoded []byte) (*ocr2aggregator.OCR2AggregatorConfigSet, error) {
	unpacked := new(ocr2aggregator.OCR2AggregatorConfigSet)
	abiPointer, err := ocr2aggregator.OCR2AggregatorMetaData.GetAbi()
	if err != nil {
		return unpacked, err
	}
	defaultABI := *abiPointer
	err = defaultABI.UnpackIntoInterface(unpacked, "ConfigSet", encoded)
	if err != nil {
		return unpacked, errors.Wrap(err, "failed to unpack log data")
	}
	return unpacked, nil
}
