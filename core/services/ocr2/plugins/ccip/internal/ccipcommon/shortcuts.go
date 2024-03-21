package ccipcommon

import (
	"context"
	"encoding/hex"
	"fmt"
	"sort"
	"sync"

	"golang.org/x/sync/errgroup"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
)

const (
	offRampBatchSizeLimit = 30
)

func GetMessageIDsAsHexString(messages []cciptypes.EVM2EVMMessage) []string {
	messageIDs := make([]string, 0, len(messages))
	for _, m := range messages {
		messageIDs = append(messageIDs, "0x"+hex.EncodeToString(m.MessageID[:]))
	}
	return messageIDs
}

type BackfillArgs struct {
	SourceLP, DestLP                 logpoller.LogPoller
	SourceStartBlock, DestStartBlock uint64
}

func GetSortedChainTokens(ctx context.Context, offRamps []ccipdata.OffRampReader, priceRegistry cciptypes.PriceRegistryReader) (chainTokens []cciptypes.Address, err error) {
	return getSortedChainTokensWithBatchLimit(ctx, offRamps, priceRegistry, offRampBatchSizeLimit)
}

// GetChainTokens returns union of all tokens supported on the destination chain, including fee tokens from the provided price registry
// and the bridgeable tokens from all the offRamps living on the chain.
func getSortedChainTokensWithBatchLimit(ctx context.Context, offRamps []ccipdata.OffRampReader, priceRegistry cciptypes.PriceRegistryReader, batchSize uint) (chainTokens []cciptypes.Address, err error) {
	if batchSize == 0 {
		return nil, fmt.Errorf("batch size must be greater than 0")
	}

	eg := new(errgroup.Group)

	var destFeeTokens []cciptypes.Address
	var destBridgeableTokens []cciptypes.Address
	lock := &sync.RWMutex{}

	eg.Go(func() error {
		tokens, err := priceRegistry.GetFeeTokens(ctx)
		if err != nil {
			return fmt.Errorf("get dest fee tokens: %w", err)
		}
		destFeeTokens = tokens
		return nil
	})

	var batchCounter uint = 0
	for _, o := range offRamps {
		offRamp := o
		eg.Go(func() error {
			tokens, err := offRamp.GetTokens(ctx)
			if err != nil {
				return fmt.Errorf("get dest bridgeable tokens: %w", err)
			}
			lock.Lock()
			destBridgeableTokens = append(destBridgeableTokens, tokens.DestinationTokens...)
			lock.Unlock()
			return nil
		})

		batchCounter++
		if batchCounter == batchSize {
			if err := eg.Wait(); err != nil {
				return nil, err
			}
			batchCounter = 0
		}
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	// same token can be returned by multiple offRamps, and fee token can overlap with bridgeable tokens,
	// we need to dedup them to arrive at chain token set
	uniqueBridgeableTokens := FlattenUniqueSlice(destFeeTokens, destBridgeableTokens)

	// return the tokens in deterministic order to aid with testing and debugging
	sort.Slice(uniqueBridgeableTokens, func(i, j int) bool {
		return uniqueBridgeableTokens[i] < uniqueBridgeableTokens[j]
	})

	return uniqueBridgeableTokens, nil
}

// GetDestinationTokens returns the destination chain fee tokens from the provided price registry
// and the bridgeable tokens from the offramp.
func GetDestinationTokens(ctx context.Context, offRamp ccipdata.OffRampReader, priceRegistry cciptypes.PriceRegistryReader) (fee, bridged []cciptypes.Address, err error) {
	eg := new(errgroup.Group)

	var destFeeTokens []cciptypes.Address
	var destBridgeableTokens []cciptypes.Address

	eg.Go(func() error {
		tokens, err := priceRegistry.GetFeeTokens(ctx)
		if err != nil {
			return fmt.Errorf("get dest fee tokens: %w", err)
		}
		destFeeTokens = tokens
		return nil
	})

	eg.Go(func() error {
		tokens, err := offRamp.GetTokens(ctx)
		if err != nil {
			return fmt.Errorf("get dest bridgeable tokens: %w", err)
		}
		destBridgeableTokens = tokens.DestinationTokens
		return nil
	})

	if err := eg.Wait(); err != nil {
		return nil, nil, err
	}

	return destFeeTokens, destBridgeableTokens, nil
}

// FlattenUniqueSlice returns a flattened slice that contains unique elements by preserving their order.
func FlattenUniqueSlice[T comparable](slices ...[]T) []T {
	seen := make(map[T]struct{})
	flattened := make([]T, 0)

	for _, sl := range slices {
		for _, el := range sl {
			if _, exists := seen[el]; !exists {
				flattened = append(flattened, el)
				seen[el] = struct{}{}
			}
		}
	}
	return flattened
}
