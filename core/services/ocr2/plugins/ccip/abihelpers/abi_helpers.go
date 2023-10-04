package abihelpers

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/libocr/gethwrappers2/ocr2aggregator"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

// MessageExecutionState defines the execution states of CCIP messages.
type MessageExecutionState uint8

const (
	ExecutionStateUntouched MessageExecutionState = iota
	ExecutionStateInProgress
	ExecutionStateSuccess
	ExecutionStateFailure
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
	commitStoreABI, err := abi.JSON(strings.NewReader(commit_store.CommitStoreABI))
	if err != nil {
		panic(err)
	}
	EventSignatures.ReportAccepted = GetIDOrPanic("ReportAccepted", commitStoreABI)
	EventSignatures.ReportAcceptedMaxSequenceNumberWord = 3

	offRampABI, err := abi.JSON(strings.NewReader(evm_2_evm_offramp.EVM2EVMOffRampABI))
	if err != nil {
		panic(err)
	}
	EventSignatures.ExecutionStateChanged = GetIDOrPanic("ExecutionStateChanged", offRampABI)
	EventSignatures.ExecutionStateChangedSequenceNumberIndex = 1
	EventSignatures.PoolAdded = GetIDOrPanic("PoolAdded", offRampABI)
	EventSignatures.PoolRemoved = GetIDOrPanic("PoolRemoved", offRampABI)

	priceRegistryABI, err := abi.JSON(strings.NewReader(price_registry.PriceRegistryABI))
	if err != nil {
		panic(err)
	}
	EventSignatures.UsdPerUnitGasUpdated = GetIDOrPanic("UsdPerUnitGasUpdated", priceRegistryABI)
	EventSignatures.UsdPerTokenUpdated = GetIDOrPanic("UsdPerTokenUpdated", priceRegistryABI)
	EventSignatures.FeeTokenAdded = GetIDOrPanic("FeeTokenAdded", priceRegistryABI)
	EventSignatures.FeeTokenRemoved = GetIDOrPanic("FeeTokenRemoved", priceRegistryABI)

	CommitReportArgs = commitStoreABI.Events["ReportAccepted"].Inputs

	manuallyExecuteMethod, ok := offRampABI.Methods["manuallyExecute"]
	if !ok {
		panic("missing event 'manuallyExecute'")
	}
	ExecutionReportArgs = manuallyExecuteMethod.Inputs[:1]
}

func MessagesFromExecutionReport(report types.Report) ([]evm_2_evm_offramp.InternalEVM2EVMMessage, error) {
	decodedExecutionReport, err := DecodeExecutionReport(report)
	if err != nil {
		return nil, err
	}
	return decodedExecutionReport.Messages, nil
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

func EncodeExecutionReport(execReport evm_2_evm_offramp.InternalExecutionReport) ([]byte, error) {
	return ExecutionReportArgs.PackValues([]interface{}{&execReport})
}

func DecodeExecutionReport(report []byte) (evm_2_evm_offramp.InternalExecutionReport, error) {
	unpacked, err := ExecutionReportArgs.Unpack(report)
	if err != nil {
		return evm_2_evm_offramp.InternalExecutionReport{}, err
	}
	if len(unpacked) == 0 {
		return evm_2_evm_offramp.InternalExecutionReport{}, errors.New("assumptionViolation: expected at least one element")
	}

	// Must be anonymous struct here
	erStruct, ok := unpacked[0].(struct {
		Messages []struct {
			SourceChainSelector uint64         `json:"sourceChainSelector"`
			Sender              common.Address `json:"sender"`
			Receiver            common.Address `json:"receiver"`
			SequenceNumber      uint64         `json:"sequenceNumber"`
			GasLimit            *big.Int       `json:"gasLimit"`
			Strict              bool           `json:"strict"`
			Nonce               uint64         `json:"nonce"`
			FeeToken            common.Address `json:"feeToken"`
			FeeTokenAmount      *big.Int       `json:"feeTokenAmount"`
			Data                []uint8        `json:"data"`
			TokenAmounts        []struct {
				Token  common.Address `json:"token"`
				Amount *big.Int       `json:"amount"`
			} `json:"tokenAmounts"`
			SourceTokenData [][]byte `json:"sourceTokenData"`
			MessageId       [32]byte `json:"messageId"`
		} `json:"messages"`
		OffchainTokenData [][][]byte  `json:"offchainTokenData"`
		Proofs            [][32]uint8 `json:"proofs"`
		ProofFlagBits     *big.Int    `json:"proofFlagBits"`
	})
	if !ok {
		return evm_2_evm_offramp.InternalExecutionReport{}, fmt.Errorf("got %T", unpacked[0])
	}
	var er evm_2_evm_offramp.InternalExecutionReport
	er.Messages = []evm_2_evm_offramp.InternalEVM2EVMMessage{}

	for _, msg := range erStruct.Messages {
		var tokensAndAmounts []evm_2_evm_offramp.ClientEVMTokenAmount
		for _, tokenAndAmount := range msg.TokenAmounts {
			tokensAndAmounts = append(tokensAndAmounts, evm_2_evm_offramp.ClientEVMTokenAmount{
				Token:  tokenAndAmount.Token,
				Amount: tokenAndAmount.Amount,
			})
		}
		er.Messages = append(er.Messages, evm_2_evm_offramp.InternalEVM2EVMMessage{
			SourceChainSelector: msg.SourceChainSelector,
			SequenceNumber:      msg.SequenceNumber,
			FeeTokenAmount:      msg.FeeTokenAmount,
			Sender:              msg.Sender,
			Nonce:               msg.Nonce,
			GasLimit:            msg.GasLimit,
			Strict:              msg.Strict,
			Receiver:            msg.Receiver,
			Data:                msg.Data,
			TokenAmounts:        tokensAndAmounts,
			SourceTokenData:     msg.SourceTokenData,
			FeeToken:            msg.FeeToken,
			MessageId:           msg.MessageId,
		})
	}

	er.OffchainTokenData = erStruct.OffchainTokenData
	er.Proofs = append(er.Proofs, erStruct.Proofs...)
	// Unpack will populate with big.Int{false, <allocated empty nat>} for 0 values,
	// which is different from the expected big.NewInt(0). Rebuild to the expected value for this case.
	er.ProofFlagBits = new(big.Int).SetBytes(erStruct.ProofFlagBits.Bytes())
	return er, nil
}

// EncodeCommitReport abi encodes an offramp.InternalCommitReport.
func EncodeCommitReport(commitReport commit_store.CommitStoreCommitReport) ([]byte, error) {
	return CommitReportArgs.PackValues([]interface{}{commitReport})
}

// DecodeCommitReport abi decodes a types.Report to an CommitStoreCommitReport
func DecodeCommitReport(report []byte) (commit_store.CommitStoreCommitReport, error) {
	unpacked, err := CommitReportArgs.Unpack(report)
	if err != nil {
		return commit_store.CommitStoreCommitReport{}, err
	}
	if len(unpacked) != 1 {
		return commit_store.CommitStoreCommitReport{}, errors.New("expected single struct value")
	}

	commitReport, ok := unpacked[0].(struct {
		PriceUpdates struct {
			TokenPriceUpdates []struct {
				SourceToken common.Address `json:"sourceToken"`
				UsdPerToken *big.Int       `json:"usdPerToken"`
			} `json:"tokenPriceUpdates"`
			DestChainSelector uint64   `json:"destChainSelector"`
			UsdPerUnitGas     *big.Int `json:"usdPerUnitGas"`
		} `json:"priceUpdates"`
		Interval struct {
			Min uint64 `json:"min"`
			Max uint64 `json:"max"`
		} `json:"interval"`
		MerkleRoot [32]byte `json:"merkleRoot"`
	})
	if !ok {
		return commit_store.CommitStoreCommitReport{}, errors.Errorf("invalid commit report got %T", unpacked[0])
	}

	var tokenPriceUpdates []commit_store.InternalTokenPriceUpdate
	for _, u := range commitReport.PriceUpdates.TokenPriceUpdates {
		tokenPriceUpdates = append(tokenPriceUpdates, commit_store.InternalTokenPriceUpdate{
			SourceToken: u.SourceToken,
			UsdPerToken: u.UsdPerToken,
		})
	}

	return commit_store.CommitStoreCommitReport{
		PriceUpdates: commit_store.InternalPriceUpdates{
			DestChainSelector: commitReport.PriceUpdates.DestChainSelector,
			UsdPerUnitGas:     commitReport.PriceUpdates.UsdPerUnitGas,
			TokenPriceUpdates: tokenPriceUpdates,
		},
		Interval: commit_store.CommitStoreInterval{
			Min: commitReport.Interval.Min,
			Max: commitReport.Interval.Max,
		},
		MerkleRoot: commitReport.MerkleRoot,
	}, nil
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
