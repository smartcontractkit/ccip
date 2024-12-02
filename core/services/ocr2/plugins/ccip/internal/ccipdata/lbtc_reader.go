package ccipdata

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
)

var (
	_ LBTCReader = &LBTCReaderImpl{}
)

const (
	LBTC_DEPOSIT_FILTER_NAME = "LBTC deposited"
	LBTC_PAYLOAD_ABI         = `[{"type": "bytes"}]`
)

type lbtcPayload []byte

func (d lbtcPayload) AbiString() string {
	return LBTC_PAYLOAD_ABI
}

func (d lbtcPayload) Validate() error {
	if len(d) == 0 {
		return errors.New("must be non-empty")
	}
	return nil
}

type LBTCReader interface {
	GetLBTCMessageInTx(ctx context.Context, payloadHash [32]byte, txHash string) ([]byte, error)
	Close() error
}

type LBTCReaderImpl struct {
	eventID            common.Hash
	lp                 logpoller.LogPoller
	filter             logpoller.Filter
	lggr               logger.Logger
	transmitterAddress common.Address

	// shortLivedInMemLogs is a short-lived cache (items expire every few seconds)
	// used to prevent frequent log fetching from the log poller
	shortLivedInMemLogs *cache.Cache
}

func NewLBTCReader(lggr logger.Logger, jobID string, transmitter common.Address, lp logpoller.LogPoller, registerFilters bool) (*LBTCReaderImpl, error) {
	return NewLBTCReaderWithCache(lggr, jobID, transmitter, lp, cache.New(shortLivedInMemLogsCacheExpiration, 2*shortLivedInMemLogsCacheExpiration), registerFilters)
}

func NewLBTCReaderWithCache(lggr logger.Logger, jobID string, transmitter common.Address, lp logpoller.LogPoller, cache *cache.Cache, registerFilters bool) (*LBTCReaderImpl, error) {
	eventSig := utils.Keccak256Fixed([]byte("DepositToBridge(address,bytes32,bytes32,bytes)"))
	r := &LBTCReaderImpl{
		lggr:    lggr,
		lp:      lp,
		eventID: eventSig,
		filter: logpoller.Filter{
			Name:      logpoller.FilterName(LBTC_DEPOSIT_FILTER_NAME, jobID, transmitter.Hex()),
			EventSigs: []common.Hash{eventSig},
			Addresses: []common.Address{transmitter},
			Retention: CommitExecLogsRetention,
		},
		transmitterAddress:  transmitter,
		shortLivedInMemLogs: cache,
	}

	if registerFilters {
		if err := r.RegisterFilters(); err != nil {
			return nil, fmt.Errorf("register filters: %w", err)
		}
	}
	return r, nil
}

func (r *LBTCReaderImpl) GetLBTCMessageInTx(ctx context.Context, payloadHash [32]byte, txHash string) ([]byte, error) {
	var lpLogs []logpoller.Log

	// fetch all the lbtc logs for the provided tx hash
	key := fmt.Sprintf("lbtc-%s", txHash)
	if rawLogs, foundInMem := r.shortLivedInMemLogs.Get(key); foundInMem {
		inMemLogs, ok := rawLogs.([]logpoller.Log)
		if !ok {
			return nil, errors.Errorf("unexpected in-mem logs type %T", rawLogs)
		}
		r.lggr.Debugw("found logs in memory", "key", key, "len", len(inMemLogs))
		lpLogs = inMemLogs
	}
	if len(lpLogs) == 0 {
		r.lggr.Debugw("fetching logs from lp")
		var err error
		lpLogs, err = r.lp.IndexedLogsByTxHash(
			ctx,
			r.eventID,
			r.transmitterAddress,
			common.HexToHash(txHash),
		)
		if err != nil {
			return nil, err
		}
		r.shortLivedInMemLogs.Set(key, lpLogs, cache.DefaultExpiration)
		r.lggr.Debugw("fetched logs from lp", "logs", len(lpLogs))
	}
	for _, log := range lpLogs {
		topics := log.GetTopics()
		if currentPayloadHash := topics[3]; currentPayloadHash == payloadHash {
			return parseLBTCDepositPayload(log.Data)
		}
	}
	return nil, fmt.Errorf("payload with hash=%s not found in logs", hexutil.Encode(payloadHash[:]))
}

func parseLBTCDepositPayload(logData []byte) ([]byte, error) {
	decodeAbiStruct, err := abihelpers.DecodeAbiStruct[lbtcPayload](logData)
	if err != nil {
		return nil, err
	}
	return decodeAbiStruct, nil
}

func (r *LBTCReaderImpl) RegisterFilters() error {
	return r.lp.RegisterFilter(context.Background(), r.filter)
}

func (r *LBTCReaderImpl) Close() error {
	return r.lp.UnregisterFilter(context.Background(), r.filter.Name)
}
