package ccipcommon

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"

	"golang.org/x/sync/errgroup"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
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

func IsTxRevertError(err error) bool {
	if err == nil {
		return false
	}

	// Geth eth_call reverts with "execution reverted"
	// Nethermind, Parity, OpenEthereum eth_call reverts with "VM execution error"
	// See: https://github.com/ethereum/go-ethereum/issues/21886
	return strings.Contains(err.Error(), "execution reverted") || strings.Contains(err.Error(), "VM execution error")
}

func SelectorToBytes(chainSelector uint64) [32]byte {
	var b [32]byte
	binary.BigEndian.PutUint64(b[:], chainSelector)
	return b
}
