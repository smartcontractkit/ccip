package memory

import (
	"context"

	"github.com/google/uuid"

	"github.com/smartcontractkit/chainlink/v2/core/services/feeds"
)

type JobClient struct {
	Nodes map[string]Node
}

func NewMemoryJobClient(nodesByPeerID map[string]Node) *JobClient {
	return &JobClient{nodesByPeerID}
}

// TODO: Use interface once ready.
func (m JobClient) ProposeJob(ctx context.Context, nodeId string, spec string) (int64, error) {
	jobProposalID, err := m.Nodes[nodeId].App.GetFeedsService().ProposeJob(ctx, &feeds.ProposeJobArgs{
		FeedsManagerID: 0,
		RemoteUUID:     uuid.New(),
		Multiaddrs:     nil,
		Version:        0,
		Spec:           spec,
	})
	return jobProposalID, err
}

func (m JobClient) GetJob(ctx context.Context, nodeId string, jobID int64) (string, error) {
	jobProposal, err := m.Nodes[nodeId].App.GetFeedsService().GetSpec(ctx, jobID)
	if err != nil {
		return "", err
	}
	return jobProposal.Definition, nil
}
