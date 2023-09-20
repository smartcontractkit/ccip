package tokendata

import (
	"context"
	"sync"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal"
)

type CachedReader struct {
	Reader

	cache      map[uint64][]byte
	cacheMutex sync.Mutex
}

func NewCachedReader(reader Reader) *CachedReader {
	return &CachedReader{
		Reader: reader,
		cache:  make(map[uint64][]byte),
	}
}

func (r *CachedReader) ReadTokenData(ctx context.Context, msg internal.EVM2EVMOnRampCCIPSendRequestedWithMeta) ([]byte, error) {
	r.cacheMutex.Lock()
	defer r.cacheMutex.Unlock()

	if data, ok := r.cache[msg.SequenceNumber]; ok {
		return data, nil
	}

	tokenData, err := r.Reader.ReadTokenData(ctx, msg)
	if err != nil {
		return []byte{}, err
	}

	r.cache[msg.SequenceNumber] = tokenData

	return tokenData, nil
}
