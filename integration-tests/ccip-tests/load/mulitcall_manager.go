package load

import (
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
)

type MultiCallManager struct {
	Client   blockchain.EVMClient
	LoadGens map[string]bool
}
