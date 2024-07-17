package ccip_integration_tests

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/hashicorp/consul/sdk/freeport"
	deployments "github.com/smartcontractkit/chainlink/v2/core/services/ocr3/plugins/deployment"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"testing"
)

type MockJobDistributor struct {
	nodes []*ocr3Node
}

func (m MockJobDistributor) GetJob(ctx context.Context, in *deployments.GetJobRequest, opts ...grpc.CallOption) (*deployments.GetJobResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockJobDistributor) GetProposal(ctx context.Context, in *deployments.GetProposalRequest, opts ...grpc.CallOption) (*deployments.GetProposalResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockJobDistributor) ListJobs(ctx context.Context, in *deployments.ListJobsRequest, opts ...grpc.CallOption) (*deployments.ListJobsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockJobDistributor) ListProposals(ctx context.Context, in *deployments.ListProposalsRequest, opts ...grpc.CallOption) (*deployments.ListProposalsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockJobDistributor) ProposeJob(ctx context.Context, in *deployments.ProposeJobRequest, opts ...grpc.CallOption) (*deployments.ProposeJobResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockJobDistributor) RevokeJob(ctx context.Context, in *deployments.RevokeJobRequest, opts ...grpc.CallOption) (*deployments.RevokeJobResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockJobDistributor) DeleteJob(ctx context.Context, in *deployments.DeleteJobRequest, opts ...grpc.CallOption) (*deployments.DeleteJobResponse, error) {
	//TODO implement me
	panic("implement me")
}

func TestMessageFlow(t *testing.T) {
	numNodes := 4
	ports := freeport.GetN(t, numNodes)
	capabilitiesPorts := freeport.GetN(t, numNodes)
	var nodes []*ocr3Node
	var nodeIds []string
	homeChainUni, universes := createUniverses(t, 4)
	chains := make(map[uint64]bind.ContractBackend)
	for chainID, uni := range universes {
		chains[chainID] = uni.backend
	}
	for i := 0; i < numNodes; i++ {
		memoryNode := setupNodeOCR3(t, ports[i], capabilitiesPorts[i], nil, universes, homeChainUni)
		nodes = append(nodes, memoryNode)
		nodeIds = append(nodeIds, memoryNode.peerID)
	}
	addressBook := deployments.NewMemoryAddressBook()
	jobServiceClient := MockJobDistributor{nodes: nodes}
	require.NoError(t, deployments.DeployNewCCIPToExistingDON(addressBook, nodeIds, chains, nil, jobServiceClient))
	// Do all the in memory testing here
}
