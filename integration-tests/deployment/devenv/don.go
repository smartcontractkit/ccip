package devenv

import (
	"context"
	"fmt"

	"github.com/AlekSi/pointer"

	nodev1 "github.com/smartcontractkit/ccip/integration-tests/deployment/jd/node/v1"
	"github.com/smartcontractkit/ccip/integration-tests/deployment/jd/shared/ptypes"
	clclient "github.com/smartcontractkit/chainlink/integration-tests/client"
	"github.com/smartcontractkit/chainlink/integration-tests/web/sdk/client"
)

type NodeInfo struct {
	CLConfig    clclient.ChainlinkConfig
	IsBootstrap bool
	Name        string
}
type DON struct {
	Bootstrap []Node
	Nodes     []Node
}

func NewRegisteredDON(ctx context.Context, nodeInfo []NodeInfo, jd JobDistributor) (*DON, error) {
	don := &DON{
		Bootstrap: make([]Node, 0),
		Nodes:     make([]Node, 0),
	}
	for i, info := range nodeInfo {
		node, err := NewNode(info.CLConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to create node %d: %w", i, err)
		}
		if info.Name == "" {
			info.Name = fmt.Sprintf("node-%d", i)
		}
		// node Labels so that it's easier to query them
		nodeLabels := make([]*ptypes.Label, 0)
		if info.IsBootstrap {
			nodeLabels = append(nodeLabels, &ptypes.Label{
				Key:   "bootstrap",
				Value: pointer.ToString("true"),
			})
		} else {
			nodeLabels = append(nodeLabels, &ptypes.Label{
				Key:   "bootstrap",
				Value: pointer.ToString("false"),
			})
		}
		// Register the node in Job distributor
		registerResponse, err := node.RegisterNodeToJobDistributor(ctx, jd.NodeServiceClient, nodeLabels, info.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to register node %w", err)
		}
		node.NodeId = registerResponse.GetId()
		if info.IsBootstrap {
			don.Bootstrap = append(don.Bootstrap, *node)
		} else {
			don.Nodes = append(don.Nodes, *node)
		}
	}
}

type Node struct {
	FMS    client.Client
	NodeId string
}

func NewNode(nodeInfo clclient.ChainlinkConfig) (*Node, error) {
	fmsClient, err := client.New(nodeInfo.URL, client.Credentials{
		Email:    nodeInfo.Email,
		Password: nodeInfo.Password,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create FMS client: %w", err)
	}
	return &Node{
		FMS: fmsClient,
	}, nil
}

func (n *Node) CreateFeedsManager(ctx context.Context, jobID string) error {
	fms, err := n.FMS.CreateFeedsManager(ctx, jobID)
	if err != nil {
		return fmt.Errorf("failed to create feeds manager: %w", err)
	}

	return nil
}

func (n *Node) GetCSAKeys(ctx context.Context) (*string, error) {
	nodeCSAResult, err := n.FMS.GetCSAKeys(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get csa keypair for node %w", err)
	}
	if nodeCSAResult.GetCsaKeys().Results == nil || len(nodeCSAResult.GetCsaKeys().Results) == 0 {
		return nil, fmt.Errorf("failed to get csa keypair for node: %w", err)
	}
	nodeCSA := nodeCSAResult.GetCsaKeys().Results[0].GetPublicKey()
	return &nodeCSA, nil
}

func (n *Node) RegisterNodeToJobDistributor(ctx context.Context, jd nodev1.NodeServiceClient, labels []*ptypes.Label, name string) (*nodev1.Node, error) {
	nodeCSAResult, err := n.GetCSAKeys(ctx)
	if err != nil || nodeCSAResult == nil {
		return nil, fmt.Errorf("failed to get csa keypair for node %s: %w", name, err)
	}

	nodeCSA := *nodeCSAResult
	registerResponse, err := jd.RegisterNode(ctx, &nodev1.RegisterNodeRequest{
		PublicKey: nodeCSA,
		Labels:    labels,
		Name:      name,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to register node %s:%w", name, err)
	}
	if registerResponse.GetNode() == nil {
		return nil, fmt.Errorf("failed to register node %s returned null response", name)
	}
	return registerResponse.GetNode(), nil
}
