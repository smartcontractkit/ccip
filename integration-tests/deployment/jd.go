package deployment

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	jobv1 "github.com/smartcontractkit/ccip/integration-tests/deployment/jd/job/v1"
	nodev1 "github.com/smartcontractkit/ccip/integration-tests/deployment/jd/node/v1"
)

type JDConfig struct {
	Server string
	Creds  credentials.TransportCredentials
}

func NewClientConnection(cfg JDConfig) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	// TODO: add auth details
	if cfg.Creds != nil {
		opts = append(opts, grpc.WithTransportCredentials(cfg.Creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	}
	conn, err := grpc.NewClient(cfg.Server, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect Job Distributor service. Err: %w", err)
	}

	return conn, nil
}

type JDClient struct {
	nodeClient nodev1.NodeServiceClient
	jobClient  jobv1.JobServiceClient
}

func (jd JDClient) GetJob(ctx context.Context, in *jobv1.GetJobRequest, opts ...grpc.CallOption) (*jobv1.GetJobResponse, error) {
	return jd.jobClient.GetJob(ctx, in, opts...)
}

func (jd JDClient) GetProposal(ctx context.Context, in *jobv1.GetProposalRequest, opts ...grpc.CallOption) (*jobv1.GetProposalResponse, error) {
	return jd.jobClient.GetProposal(ctx, in, opts...)
}

func (jd JDClient) ListJobs(ctx context.Context, in *jobv1.ListJobsRequest, opts ...grpc.CallOption) (*jobv1.ListJobsResponse, error) {
	return jd.jobClient.ListJobs(ctx, in, opts...)
}

func (jd JDClient) ListProposals(ctx context.Context, in *jobv1.ListProposalsRequest, opts ...grpc.CallOption) (*jobv1.ListProposalsResponse, error) {
	return jd.jobClient.ListProposals(ctx, in, opts...)
}

func (jd JDClient) ProposeJob(ctx context.Context, in *jobv1.ProposeJobRequest, opts ...grpc.CallOption) (*jobv1.ProposeJobResponse, error) {
	return jd.jobClient.ProposeJob(ctx, in, opts...)
}

func (jd JDClient) RevokeJob(ctx context.Context, in *jobv1.RevokeJobRequest, opts ...grpc.CallOption) (*jobv1.RevokeJobResponse, error) {
	return jd.jobClient.RevokeJob(ctx, in, opts...)
}

func (jd JDClient) DeleteJob(ctx context.Context, in *jobv1.DeleteJobRequest, opts ...grpc.CallOption) (*jobv1.DeleteJobResponse, error) {
	return jd.jobClient.DeleteJob(ctx, in, opts...)
}

func (jd JDClient) GetNode(ctx context.Context, in *nodev1.GetNodeRequest, opts ...grpc.CallOption) (*nodev1.GetNodeResponse, error) {
	return jd.nodeClient.GetNode(ctx, in, opts...)
}

func (jd JDClient) ListNodes(ctx context.Context, in *nodev1.ListNodesRequest, opts ...grpc.CallOption) (*nodev1.ListNodesResponse, error) {
	return jd.nodeClient.ListNodes(ctx, in, opts...)
}

func (jd JDClient) ListNodeChainConfigs(ctx context.Context, in *nodev1.ListNodeChainConfigsRequest, opts ...grpc.CallOption) (*nodev1.ListNodeChainConfigsResponse, error) {
	return jd.nodeClient.ListNodeChainConfigs(ctx, in, opts...)
}

func NewJDClient(cfg JDConfig) (OffchainClient, error) {
	conn, err := NewClientConnection(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect Job Distributor service. Err: %w", err)
	}

	return &JDClient{
		nodeClient: nodev1.NewNodeServiceClient(conn),
		jobClient:  jobv1.NewJobServiceClient(conn),
	}, err
}
