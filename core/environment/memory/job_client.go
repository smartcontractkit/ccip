package memory

import (
	"context"
	"github.com/smartcontractkit/chainlink/v2/core/services/chainlink"
)

type JobClient struct {
	app chainlink.Application
}

func NewMemoryJobClient(app chainlink.Application) *JobClient {
	return &JobClient{app}
}

func (m JobClient) ProposeJob(ctx context.Context, nodeId string, spec string) (string, error) {
	// Use proto once ready.
	return "", nil
}
