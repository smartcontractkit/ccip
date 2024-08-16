package deployment

import (
	"github.com/smartcontractkit/ccip/integration-tests/client"
)

type DON struct {
	Bootstrap CoreNode
	Nodes     []CoreNode
}

type CoreNode struct {
	WebConfig client.ChainlinkConfig
	CSAKey    string
	Cookie    string
}
