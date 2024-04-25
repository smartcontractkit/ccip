package codec

import (
	"context"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/ccip/internal/model"
)

type Commit interface {
	Encode(context.Context, model.CommitPluginReport) ([]byte, error)
	Decode(context.Context, []byte) (model.CommitPluginReport, error)
}

type Execute interface {
	Encode(context.Context, model.ExecutePluginReport) ([]byte, error)
	Decode(context.Context, []byte) (model.ExecutePluginReport, error)
}
