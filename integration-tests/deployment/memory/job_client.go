package memory

import (
	"context"

	"google.golang.org/grpc"

	jobv1 "github.com/smartcontractkit/chainlink/integration-tests/deployment/jd/job/v1"
	nodev1 "github.com/smartcontractkit/chainlink/integration-tests/deployment/jd/node/v1"

	"github.com/smartcontractkit/chainlink/v2/core/capabilities/ccip/validate"
)

type JobClient struct {
	Nodes map[string]Node
}

func (j JobClient) ArchiveNode(ctx context.Context, in *nodev1.ArchiveNodeRequest, opts ...grpc.CallOption) (*nodev1.ArchiveNodeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (j JobClient) CreateNode(ctx context.Context, in *nodev1.CreateNodeRequest, opts ...grpc.CallOption) (*nodev1.CreateNodeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (j JobClient) GetNode(ctx context.Context, in *nodev1.GetNodeRequest, opts ...grpc.CallOption) (*nodev1.GetNodeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (j JobClient) ListNodes(ctx context.Context, in *nodev1.ListNodesRequest, opts ...grpc.CallOption) (*nodev1.ListNodesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (j JobClient) ListNodeChainConfigs(ctx context.Context, in *nodev1.ListNodeChainConfigsRequest, opts ...grpc.CallOption) (*nodev1.ListNodeChainConfigsResponse, error) {
	n := j.Nodes[in.Filter.NodeId]
	// TODO: I think we can pull it from the feeds manager.
	return &nodev1.ListNodeChainConfigsResponse{
		ChainConfigs: []*nodev1.ChainConfig{
			{
				Ocr2Config: &nodev1.OCR2Config{
					P2PKeyBundle: &nodev1.OCR2Config_P2PKeyBundle{
						PeerId:    n.Keys.PeerID.String(),
						PublicKey: "",
					},
					OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{
						BundleId: n.Keys.OCRKeyBundle.ID(),
					},
					IsBootstrap: n.IsBoostrap,
					Multiaddr:   n.Addr.String(),
				},
			},
		},
	}, nil
}

func (j JobClient) UnarchiveNode(ctx context.Context, in *nodev1.UnarchiveNodeRequest, opts ...grpc.CallOption) (*nodev1.UnarchiveNodeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (j JobClient) UpdateNode(ctx context.Context, in *nodev1.UpdateNodeRequest, opts ...grpc.CallOption) (*nodev1.UpdateNodeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (j JobClient) GetJob(ctx context.Context, in *jobv1.GetJobRequest, opts ...grpc.CallOption) (*jobv1.GetJobResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (j JobClient) GetProposal(ctx context.Context, in *jobv1.GetProposalRequest, opts ...grpc.CallOption) (*jobv1.GetProposalResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (j JobClient) ListJobs(ctx context.Context, in *jobv1.ListJobsRequest, opts ...grpc.CallOption) (*jobv1.ListJobsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (j JobClient) ListProposals(ctx context.Context, in *jobv1.ListProposalsRequest, opts ...grpc.CallOption) (*jobv1.ListProposalsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (j JobClient) ProposeJob(ctx context.Context, in *jobv1.ProposeJobRequest, opts ...grpc.CallOption) (*jobv1.ProposeJobResponse, error) {
	n := j.Nodes[in.NodeId]
	// TODO: Use FMS
	jb, err := validate.ValidatedCCIPSpec(in.Spec)
	if err != nil {
		return nil, err
	}
	err = n.App.AddJobV2(ctx, &jb)
	if err != nil {
		return nil, err
	}
	return &jobv1.ProposeJobResponse{Proposal: &jobv1.Proposal{
		Id:      "",
		Version: 0,
		// Auto approve for now
		Status:             jobv1.ProposalStatus_PROPOSAL_STATUS_APPROVED,
		DeliveryStatus:     jobv1.ProposalDeliveryStatus_PROPOSAL_DELIVERY_STATUS_DELIVERED,
		Spec:               in.Spec,
		JobId:              jb.ExternalJobID.String(),
		CreatedAt:          nil,
		UpdatedAt:          nil,
		AckedAt:            nil,
		ResponseReceivedAt: nil,
	}}, nil
}

func (j JobClient) RevokeJob(ctx context.Context, in *jobv1.RevokeJobRequest, opts ...grpc.CallOption) (*jobv1.RevokeJobResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (j JobClient) DeleteJob(ctx context.Context, in *jobv1.DeleteJobRequest, opts ...grpc.CallOption) (*jobv1.DeleteJobResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewMemoryJobClient(nodesByPeerID map[string]Node) *JobClient {
	return &JobClient{nodesByPeerID}
}
