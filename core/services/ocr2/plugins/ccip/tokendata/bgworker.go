package tokendata

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
)

type msgResult struct {
	TokenAmountIndex int
	Err              error
	Data             []byte
}

type Worker interface {
	// AddJobsFromMsgs will include the provided msgs for background processing.
	AddJobsFromMsgs(ctx context.Context, msgs []internal.EVM2EVMOnRampCCIPSendRequestedWithMeta)

	// GetMsgTokenData returns the token data for the provided msg. If data are not ready it keeps waiting
	// until they get ready. Important: Make sure to pass a proper context with timeout.
	GetMsgTokenData(ctx context.Context, msg internal.EVM2EVMOnRampCCIPSendRequestedWithMeta) ([][]byte, error)

	GetReaders() map[common.Address]Reader
}

type BackgroundWorker struct {
	tokenDataReaders map[common.Address]Reader
	numWorkers       int
	jobsChan         chan internal.EVM2EVMOnRampCCIPSendRequestedWithMeta
	resultsCache     *resultsCache
}

func (w *BackgroundWorker) AddJobsFromMsgs(ctx context.Context, msgs []internal.EVM2EVMOnRampCCIPSendRequestedWithMeta) {
	go func() {
		for _, msg := range msgs {
			select {
			case <-ctx.Done():
				return
			default:
				if len(msg.TokenAmounts) > 0 {
					w.jobsChan <- msg
				}
			}
		}
	}()
}

func (w *BackgroundWorker) GetReaders() map[common.Address]Reader {
	return w.tokenDataReaders
}

func (w *BackgroundWorker) GetMsgTokenData(ctx context.Context, msg internal.EVM2EVMOnRampCCIPSendRequestedWithMeta) ([][]byte, error) {
	res, err := w.getMsgTokenData(ctx, msg.SequenceNumber)
	if err != nil {
		return nil, err
	}

	tokenDatas := make([][]byte, len(msg.TokenAmounts))
	for _, r := range res {
		if r.Err != nil {
			return nil, r.Err
		}
		if r.TokenAmountIndex < 0 || r.TokenAmountIndex >= len(msg.TokenAmounts) {
			return nil, fmt.Errorf("token data index incosistency")
		}
		tokenDatas[r.TokenAmountIndex] = r.Data
	}

	return tokenDatas, nil
}

func NewBackgroundWorker(ctx context.Context, tokenDataReaders map[common.Address]Reader, numWorkers int) *BackgroundWorker {
	w := &BackgroundWorker{
		tokenDataReaders: tokenDataReaders,
		numWorkers:       numWorkers,
		jobsChan:         make(chan internal.EVM2EVMOnRampCCIPSendRequestedWithMeta),
		resultsCache:     newResultsCache(),
	}

	w.spawnWorkers(ctx)
	return w
}

func (w *BackgroundWorker) spawnWorkers(ctx context.Context) {
	for i := 0; i < w.numWorkers; i++ {
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case msg := <-w.jobsChan:
					res := w.work(ctx, msg)
					w.resultsCache.add(msg.SequenceNumber, res)
				}
			}
		}()
	}
}

func (w *BackgroundWorker) work(ctx context.Context, msg internal.EVM2EVMOnRampCCIPSendRequestedWithMeta) []msgResult {
	results := make([]msgResult, 0, len(msg.TokenAmounts))

	cachedTokenData := make(map[int]msgResult) // tokenAmount index -> token data
	if cachedData, exists := w.resultsCache.get(msg.SequenceNumber); exists {
		for _, r := range cachedData {
			cachedTokenData[r.TokenAmountIndex] = r
		}
	}

	for i, token := range msg.TokenAmounts {
		offchainTokenDataProvider, exists := w.tokenDataReaders[token.Token]
		if !exists {
			// No token data required
			continue
		}

		// if the result exists in the cache and there wasn't any error keep the existing result
		if cachedResult, exists := cachedTokenData[i]; exists && cachedResult.Err == nil {
			results = append(results, cachedResult)
			continue
		}

		// if there was any error or if the data do not exist in the cache make a call to the provider
		tknData, err := offchainTokenDataProvider.ReadTokenData(ctx, msg, i)
		results = append(results, msgResult{
			TokenAmountIndex: i,
			Err:              err,
			Data:             tknData,
		})
	}

	return results
}

func (w *BackgroundWorker) getMsgTokenData(ctx context.Context, seqNum uint64) ([]msgResult, error) {
	if msgTokenData, exists := w.resultsCache.get(seqNum); exists {
		return msgTokenData, nil
	}

	// wait until the results are ready or until context timeout is reached
	tick := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			return nil, context.DeadlineExceeded
		case <-tick.C:
			if msgTokenData, exists := w.resultsCache.get(seqNum); exists {
				return msgTokenData, nil
			}
		}
	}
}

type resultsCache struct {
	results   map[uint64][]msgResult
	resultsMu *sync.RWMutex
}

func newResultsCache() *resultsCache {
	return &resultsCache{
		results:   make(map[uint64][]msgResult),
		resultsMu: &sync.RWMutex{},
	}
}

func (c *resultsCache) add(msgSeqNum uint64, results []msgResult) {
	c.resultsMu.Lock()
	defer c.resultsMu.Unlock()
	c.results[msgSeqNum] = results
}

func (c *resultsCache) get(msgSeqNum uint64) ([]msgResult, bool) {
	c.resultsMu.RLock()
	defer c.resultsMu.RUnlock()
	v, exists := c.results[msgSeqNum]
	return v, exists
}
