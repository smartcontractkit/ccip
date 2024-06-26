package mocks

import (
	"context"
	"encoding/json"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type CommitPluginJSONReportCodec struct{}

func NewCommitPluginJSONReportCodec() *CommitPluginJSONReportCodec {
	return &CommitPluginJSONReportCodec{}
}

func (c CommitPluginJSONReportCodec) Encode(ctx context.Context, report cciptypes.CommitPluginReport) ([]byte, error) {
	return json.Marshal(report)
}

func (c CommitPluginJSONReportCodec) Decode(ctx context.Context, bytes []byte) (cciptypes.CommitPluginReport, error) {
	report := cciptypes.CommitPluginReport{}
	err := json.Unmarshal(bytes, &report)
	return report, err
}

type ExecutePluginJSONReportCodec struct{}

func NewExecutePluginJSONReportCodec() *ExecutePluginJSONReportCodec {
	return &ExecutePluginJSONReportCodec{}
}

func (c ExecutePluginJSONReportCodec) Encode(_ context.Context, report cciptypes.ExecutePluginReport) ([]byte, error) {
	return json.Marshal(report)
}

func (c ExecutePluginJSONReportCodec) Decode(_ context.Context, bytes []byte) (cciptypes.ExecutePluginReport, error) {
	report := cciptypes.ExecutePluginReport{}
	err := json.Unmarshal(bytes, &report)
	return report, err
}
