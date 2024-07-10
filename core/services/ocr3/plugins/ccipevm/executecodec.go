package ccipevm

import (
	"context"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// ExecutePluginCodecV1 is a codec for encoding and decoding execute plugin reports.
// Compatible with:
// - "EVM2EVMMultiOffRamp 1.6.0-dev"
type ExecutePluginCodecV1 struct{}

func NewExecutePluginCodecV1() *ExecutePluginCodecV1 {
	return &ExecutePluginCodecV1{}
}

func (e ExecutePluginCodecV1) Encode(ctx context.Context, report cciptypes.ExecutePluginReport) ([]byte, error) {
	panic("implement me")
}

func (e ExecutePluginCodecV1) Decode(ctx context.Context, bytes []byte) (cciptypes.ExecutePluginReport, error) {
	panic("implement me")
}

// Ensure ExecutePluginCodec implements the ExecutePluginCodec interface
var _ cciptypes.ExecutePluginCodec = (*ExecutePluginCodecV1)(nil)
