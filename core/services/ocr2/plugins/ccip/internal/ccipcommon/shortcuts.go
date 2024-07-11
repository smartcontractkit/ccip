package ccipcommon

import (
	"encoding/binary"
	"encoding/hex"
	"strings"
	"time"

	"github.com/avast/retry-go/v4"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
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

func SelectorToBytes(chainSelector uint64) [16]byte {
	var b [16]byte
	binary.BigEndian.PutUint64(b[:], chainSelector)
	return b
}

// RetryUntilSuccess repeatedly calls fn until it returns a nil error. After each failed call there is an exponential
// backoff applied, between initialDelay and maxDelay.
func RetryUntilSuccess[T any](fn func() (T, error), initialDelay time.Duration, maxDelay time.Duration) (T, error) {
	return retry.DoWithData(
		fn,
		retry.Delay(initialDelay),
		retry.MaxDelay(maxDelay),
		retry.DelayType(retry.BackOffDelay),
		retry.UntilSucceeded(),
	)
}
