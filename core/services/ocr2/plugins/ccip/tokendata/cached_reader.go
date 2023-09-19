package tokendata

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum/common"
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

func (r *CachedReader) ReadTokenData(ctx context.Context, seqNum uint64, logIndex uint, txHash common.Hash) ([]byte, error) {
	r.cacheMutex.Lock()
	defer r.cacheMutex.Unlock()

	if data, ok := r.cache[seqNum]; ok {
		return data, nil
	}

	tokenData, err := r.Reader.ReadTokenData(ctx, seqNum, logIndex, txHash)
	if err != nil {
		return []byte{}, err
	}

	r.cache[seqNum] = tokenData

	return tokenData, nil
}
