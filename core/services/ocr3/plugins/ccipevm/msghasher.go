package ccipevm

import (
	"context"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

var _ cciptypes.MessageHasher = (*MessageHasher)(nil)

type MessageHasher struct{}

func NewMessageHasher() *MessageHasher {
	return &MessageHasher{}
}

func (m MessageHasher) Hash(ctx context.Context, msg cciptypes.CCIPMsg) (cciptypes.Bytes32, error) {
	panic("implement me")
}
