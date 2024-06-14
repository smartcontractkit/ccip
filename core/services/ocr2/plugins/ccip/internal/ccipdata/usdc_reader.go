package ccipdata

import (
	"context"
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
	USDC_MESSAGE_NONCE_INDEX = 2
)

var _ USDCReader = &USDCReaderImpl{}

//go:generate mockery --quiet --name USDCReader --filename usdc_reader_mock.go --case=underscore
type USDCReader interface {
	GetUSDCMessageWithNonce(ctx context.Context, nonce [32]byte) ([]byte, error)
}

type USDCReaderImpl struct {
	usdcMessageSent    common.Hash
	lp                 logpoller.LogPoller
	filter             logpoller.Filter
	lggr               logger.Logger
	transmitterAddress common.Address
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

func (u *USDCReaderImpl) GetUSDCMessageWithNonce(ctx context.Context, nonce [32]byte) ([]byte, error) {
	logs, err := u.lp.IndexedLogsTopicRange(
		ctx,
		u.usdcMessageSent,
		u.transmitterAddress,
		USDC_MESSAGE_NONCE_INDEX,
		nonce,
		nonce,
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
