package devenv

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/ccip/integration-tests/web/sdk/generated"
	clclient "github.com/smartcontractkit/chainlink/integration-tests/client"
	"github.com/smartcontractkit/chainlink/integration-tests/web/sdk/client"
)

type DON struct {
	Bootstrap Node
	Nodes     []Node
}

type Node struct {
	FMS    client.Client
	NodeId string
}

func NewNode(ctx context.Context, nodeInfo clclient.ChainlinkConfig, jdURL string) error {
	fms, err := client.New(nodeInfo.URL, client.Credentials{
		Email:    nodeInfo.Email,
		Password: nodeInfo.Password,
	})
	if err != nil {
		return fmt.Errorf("failed to create FMS client: %w", err)
	}
	// create feeds manager client
	// TODO: if can be connected during node startup, then we can remove this
	fms.CreateFeedsManager(ctx, generated.CreateFeedsManagerInput{
		Name:      "",
		Uri:       "",
		PublicKey: "",
	})
}
