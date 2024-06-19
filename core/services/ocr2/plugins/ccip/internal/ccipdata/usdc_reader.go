package ccipdata

import (
	"context"
	"encoding/binary"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
)

const (
	MESSAGE_SENT_FILTER_NAME = "USDC message sent"
	USDC_MESSAGE_NONCE_INDEX = 3
	USDC_MESSAGE_VERSION     = uint32(0)
)

var _ USDCReader = &USDCReaderImpl{}

//go:generate mockery --quiet --name USDCReader --filename usdc_reader_mock.go --case=underscore
type USDCReader interface {
	GetUSDCMessageWithNonce(ctx context.Context, nonce uint64) ([]byte, error)
}

type USDCReaderImpl struct {
	usdcMessageSent    common.Hash
	lp                 logpoller.LogPoller
	filter             logpoller.Filter
	lggr               logger.Logger
	transmitterAddress common.Address
	sourceDomain       uint32
	destDomain         uint32
}

func (u *USDCReaderImpl) Close() error {
	// FIXME Dim pgOpts removed from LogPoller
	return u.lp.UnregisterFilter(context.Background(), u.filter.Name)
}

func (u *USDCReaderImpl) RegisterFilters() error {
	// FIXME Dim pgOpts removed from LogPoller
	return u.lp.RegisterFilter(context.Background(), u.filter)
}

// usdcPayload has to match the onchain event emitted by the USDC message transmitter
type usdcPayload []byte

func (d usdcPayload) AbiString() string {
	return `[{"type": "bytes"}]`
}

func (d usdcPayload) Validate() error {
	if len(d) == 0 {
		return errors.New("must be non-empty")
	}
	return nil
}

func parseUSDCMessageSent(logData []byte) ([]byte, error) {
	decodeAbiStruct, err := abihelpers.DecodeAbiStruct[usdcPayload](logData)
	if err != nil {
		return nil, err
	}
	return decodeAbiStruct, nil
}

func GetExpectedNonceSlotData(sourceDomain, destDomain uint32, nonce uint64) [32]byte {
	// USDC message payload:
	// uint32 _msgVersion,
	// uint32 _msgSourceDomain,
	// uint32 _msgDestinationDomain,
	// uint64 _msgNonce,
	// bytes32 _msgSender,
	// Since it's packed, all of these values contribute to the first slot

	msgVersionBytes := [4]byte{}
	binary.BigEndian.PutUint32(msgVersionBytes[:], USDC_MESSAGE_VERSION)

	sourceDomainBytes := [4]byte{}
	binary.BigEndian.PutUint32(sourceDomainBytes[:], sourceDomain)

	destDomainBytes := [4]byte{}
	binary.BigEndian.PutUint32(destDomainBytes[:], destDomain)

	nonceBytes := [8]byte{}
	binary.BigEndian.PutUint64(nonceBytes[:], nonce)

	senderBytes := [12]byte{}

	return [32]byte(append(append(append(append(
		msgVersionBytes[:],
		sourceDomainBytes[:]...),
		destDomainBytes[:]...),
		nonceBytes[:]...),
		senderBytes[:]...))
}

func (u *USDCReaderImpl) GetUSDCMessageWithNonce(ctx context.Context, nonce uint64) ([]byte, error) {
	expectedSlotData := GetExpectedNonceSlotData(u.sourceDomain, u.destDomain, nonce)
	logs, err := u.lp.IndexedLogsTopicRange(
		ctx,
		u.usdcMessageSent,
		u.transmitterAddress,
		USDC_MESSAGE_NONCE_INDEX,
		expectedSlotData,
		expectedSlotData,
		types.Finalized,
	)
	if err != nil {
		return nil, err
	}

	if len(logs) == 0 {
		return nil, errors.New("no USDC message found")
	}

	if len(logs) > 1 {
		return nil, errors.New("more than one USDC message found, nonces should be unique")
	}

	return parseUSDCMessageSent(logs[0].Data)
}

func NewUSDCReader(lggr logger.Logger, jobID string, transmitter common.Address, lp logpoller.LogPoller, registerFilters bool) (*USDCReaderImpl, error) {
	eventSig := utils.Keccak256Fixed([]byte("MessageSent(bytes)"))

	r := &USDCReaderImpl{
		lggr:            lggr,
		lp:              lp,
		usdcMessageSent: eventSig,
		filter: logpoller.Filter{
			Name:      logpoller.FilterName(MESSAGE_SENT_FILTER_NAME, jobID, transmitter.Hex()),
			EventSigs: []common.Hash{eventSig},
			Addresses: []common.Address{transmitter},
			Retention: CommitExecLogsRetention,
		},
		transmitterAddress: transmitter,
		sourceDomain:       0, // TODO
		destDomain:         1, // TODO
	}

	if registerFilters {
		if err := r.RegisterFilters(); err != nil {
			return nil, fmt.Errorf("register filters: %w", err)
		}
	}
	return r, nil
}

func CloseUSDCReader(lggr logger.Logger, jobID string, transmitter common.Address, lp logpoller.LogPoller) error {
	r, err := NewUSDCReader(lggr, jobID, transmitter, lp, false)
	if err != nil {
		return err
	}
	return r.Close()
}
