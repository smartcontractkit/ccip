package ccipcommon

import (
	"context"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
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
	feeTokens, tokensPerOffRamp, err := getTokensPerOffRamp(ctx, offRamps, priceRegistry, offRampBatchSizeLimit)
	if err != nil {
		return nil, err
	}

	return AggregateChainTokens(feeTokens, tokensPerOffRamp), nil
}

func GetTokensPerOffRamp(ctx context.Context, offRamps []ccipdata.OffRampReader, priceRegistry cciptypes.PriceRegistryReader) (feeTokens []cciptypes.Address, tokensPerOffRamp map[cciptypes.Address][]cciptypes.Address, err error) {
	return getTokensPerOffRamp(ctx, offRamps, priceRegistry, offRampBatchSizeLimit)
}

func getTokensPerOffRamp(ctx context.Context, offRamps []ccipdata.OffRampReader, priceRegistry cciptypes.PriceRegistryReader, batchSize int) (feeTokens []cciptypes.Address, tokensPerOffRamp map[cciptypes.Address][]cciptypes.Address, err error) {
	if batchSize == 0 {
		return nil, nil, fmt.Errorf("batch size must be greater than 0")
	}

	eg := new(errgroup.Group)
	eg.SetLimit(batchSize)

	mu := &sync.RWMutex{}

	eg.Go(func() error {
		tokens, err := priceRegistry.GetFeeTokens(ctx)
		if err != nil {
			return fmt.Errorf("get dest fee tokens: %w", err)
		}
		feeTokens = tokens
		return nil
	})

	tokensPerOffRamp = make(map[cciptypes.Address][]cciptypes.Address)

	for _, o := range offRamps {
		offRamp := o
		eg.Go(func() error {
			tokens, err := offRamp.GetTokens(ctx)
			if err != nil {
				return fmt.Errorf("get dest bridgeable tokens: %w", err)
			}

			offRampAddr, err := offRamp.Address(ctx)
			if err != nil {
				return fmt.Errorf("get offramp address: %w", err)
			}

			mu.Lock()
			tokensPerOffRamp[offRampAddr] = tokens.DestinationTokens
			mu.Unlock()
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, nil, err
	}

	return feeTokens, tokensPerOffRamp, nil
}

func AggregateChainTokens(feeTokens []cciptypes.Address, tokensPerOffRamp map[cciptypes.Address][]cciptypes.Address) []cciptypes.Address {
	var destBridgeableTokens []cciptypes.Address
	for _, tokens := range tokensPerOffRamp {
		destBridgeableTokens = append(destBridgeableTokens, tokens...)
	}

	// same token can be returned by multiple offRamps, and fee token can overlap with bridgeable tokens,
	// we need to dedup them to arrive at chain token set
	chainTokens := FlattenUniqueSlice(feeTokens, destBridgeableTokens)

	// return the tokens in deterministic order to aid with testing and debugging
	sort.Slice(chainTokens, func(i, j int) bool {
		return chainTokens[i] < chainTokens[j]
	})

	return chainTokens
}

// DedupTokensPerOffRamp dedups tokens in the provided tokensPerOffRamp, and returns it.
// If a token exists in leader Offramp, it will be removed from the other OffRamps.
// If a token exists in multiple non-leader OffRamps, it will be removed from all but one of the OffRamps, no ordering is guaranteed.
func DedupTokensPerOffRamp(leaderOffRampAddr cciptypes.Address, tokensPerOffRamp map[cciptypes.Address][]cciptypes.Address) map[cciptypes.Address][]cciptypes.Address {
	tokenExists := make(map[cciptypes.Address]bool)
	for _, tokens := range tokensPerOffRamp[leaderOffRampAddr] {
		tokenExists[tokens] = true
	}

	// dedup leader lane tokens
	tokensPerOffRamp[leaderOffRampAddr] = FlattenUniqueSlice(tokensPerOffRamp[leaderOffRampAddr])

	// dedup non-leader lane tokens
	for offRampAddr, laneTokens := range tokensPerOffRamp {
		var dedupedTokens []cciptypes.Address
		for _, token := range laneTokens {
			if !tokenExists[token] {
				dedupedTokens = append(dedupedTokens, token)
				tokenExists[token] = true
			}
		}
		if len(dedupedTokens) == 0 {
			delete(tokensPerOffRamp, offRampAddr)
		} else {
			tokensPerOffRamp[offRampAddr] = dedupedTokens
		}
	}

	return tokensPerOffRamp
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

func IsTxRevertError(err error) bool {
	if err == nil {
		return false
	}

	// Geth eth_call reverts with "execution reverted"
	// Nethermind, Parity, OpenEthereum eth_call reverts with "VM execution error"
	// See: https://github.com/ethereum/go-ethereum/issues/21886
	return strings.Contains(err.Error(), "execution reverted") || strings.Contains(err.Error(), "VM execution error")
}
