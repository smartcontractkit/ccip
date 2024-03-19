package ccipcommon

import (
	"context"
	"encoding/hex"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/cciptypes"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
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

// TODO Matt
// GetChainTokens returns union of all tokens supported on the destination chain, including fee tokens from the provided price registry
// and the bridgeable tokens from all the offramps living on the chain.
func GetChainTokens(ctx context.Context, offRamps []ccipdata.OffRampReader, priceRegistry cciptypes.PriceRegistryReader) (fee, bridged []cciptypes.Address, err error) {
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

	for _, o := range offRamps {
		offRamp := o
		eg.Go(func() error {
			tokens, err := offRamp.GetTokens(ctx)
			if err != nil {
				return fmt.Errorf("get dest bridgeable tokens from offramp %s: %w", offRamp.Address(), err)
			}
			lock.Lock()
			destBridgeableTokens = append(destBridgeableTokens, tokens.DestinationTokens...)
			lock.Unlock()
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, nil, err
	}

	// Same token can be returned by multiple offRamps, we need to dedup them
	existingTokens := make(map[cciptypes.Address]bool)
	var uniqueBridgeableTokens []cciptypes.Address

	for _, token := range destBridgeableTokens {
		if _, ok := existingTokens[token]; !ok {
			existingTokens[token] = true
			uniqueBridgeableTokens = append(uniqueBridgeableTokens, token)
		}
	}

	return destFeeTokens, uniqueBridgeableTokens, nil
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
