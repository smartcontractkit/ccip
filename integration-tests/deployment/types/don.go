package types

import (
	"github.com/smartcontractkit/ccip/integration-tests/client"
)

type DON struct {
	Bootstrap client.ChainlinkConfig
	Nodes     []client.ChainlinkConfig
}
