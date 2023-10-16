package testconfig

import (
	"fmt"
	"testing"

	"github.com/pelletier/go-toml/v2"

	"github.com/smartcontractkit/ccip/integration-tests/types/config/node"
	"github.com/smartcontractkit/ccip/integration-tests/utils"
)

func TestSmoke(t *testing.T) {
	CreateToml()
}

func CreateToml() {
	chainlinkcfg, err := node.NewBaseConfig().TOMLString()
	cfg := Config{
		Chainlink: &Chainlink{
			Common: &Node{
				Image:      "smartcontract/chainlink:latest",
				Tag:        "latest",
				NodeConfig: chainlinkcfg,
				DBImage:    "postgres",
				DBTag:      "latest",
			},
			NoOfNodes: utils.Ptr(5),
			Nodes: []*Node{
				{
					Name:       "node1",
					Image:      "smartcontract/chainlink:latest",
					Tag:        "latest",
					NodeConfig: chainlinkcfg,
					DBImage:    "postgres",
					DBTag:      "latest",
				},
				{
					Name:       "node2",
					Image:      "smartcontract/chainlink:latest",
					Tag:        "latest",
					NodeConfig: chainlinkcfg,
					DBImage:    "postgres",
					DBTag:      "latest",
				},
			},
		},
		Networks: []string{"SIMULATED", "ROPSTEN"},
		Products: map[string]map[string]TestConfig{
			"CCIP": {
				"load": CCIPTestConfig{
					MsgType: "WithToken",
				},
			},
		},
	}
	d, err := toml.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(d))
}
