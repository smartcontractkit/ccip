package integration_tests

import (
	"context"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/deployment/jobdistributor"
	"google.golang.org/grpc"
)

type MockJobDistributor struct {
	nodes []*ocr3Node
}

func NewJobDistributor(nodes []*ocr3Node) MockJobDistributor {
	return MockJobDistributor{
		nodes: nodes,
	}
}

func (m MockJobDistributor) GetJob(ctx context.Context, in *jobdistributor.GetJobRequest, opts ...grpc.CallOption) (*jobdistributor.GetJobResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockJobDistributor) GetProposal(ctx context.Context, in *jobdistributor.GetProposalRequest, opts ...grpc.CallOption) (*jobdistributor.GetProposalResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockJobDistributor) ListJobs(ctx context.Context, in *jobdistributor.ListJobsRequest, opts ...grpc.CallOption) (*jobdistributor.ListJobsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockJobDistributor) ListProposals(ctx context.Context, in *jobdistributor.ListProposalsRequest, opts ...grpc.CallOption) (*jobdistributor.ListProposalsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockJobDistributor) ProposeJob(ctx context.Context, in *jobdistributor.ProposeJobRequest, opts ...grpc.CallOption) (*jobdistributor.ProposeJobResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockJobDistributor) RevokeJob(ctx context.Context, in *jobdistributor.RevokeJobRequest, opts ...grpc.CallOption) (*jobdistributor.RevokeJobResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockJobDistributor) DeleteJob(ctx context.Context, in *jobdistributor.DeleteJobRequest, opts ...grpc.CallOption) (*jobdistributor.DeleteJobResponse, error) {
	//TODO implement me
	panic("implement me")
}
