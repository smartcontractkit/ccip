package ccipcommon

import (
	"context"
	"encoding/hex"
	"fmt"

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

// IsDown returns true if the commitStore is down (paused or destination cursed) or if the source chain is cursed
// Both RPCs are called in parallel to save some time. These calls cannot be batches because they aim different chains.
func IsDown(ctx context.Context, commitStore ccipdata.CommitStoreReader, onRamp ccipdata.OnRampReader) (bool, error) {
	eg := new(errgroup.Group)

	var isDown bool
	var isCursed bool

	eg.Go(func() error {
		var err error
		isDown, err = commitStore.IsDown(ctx)
		if err != nil {
			return fmt.Errorf("is down: %w", err)
		}
		return nil
	})

	eg.Go(func() error {
		var err error
		isCursed, err = onRamp.IsSourceCursed(ctx)
		if err != nil {
			return fmt.Errorf("is cursed: %w", err)
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return false, err
	}

	return isDown || isCursed, nil
}
