package testconfig

import (
	"fmt"
	"testing"
	"time"

	"github.com/pelletier/go-toml/v2"

	node2 "github.com/smartcontractkit/ccip/integration-tests/ccip-tests/types/config/node"
	"github.com/smartcontractkit/ccip/integration-tests/utils"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

func TestSmoke(t *testing.T) {
	CreateToml()
}

func CreateToml() {
	chainlinkcfg, err := node2.NewConfigFromToml(node2.CCIPTOML)
	tomlString, err := chainlinkcfg.TOMLString()
	cfg := Config{map[string]ProductTest{
		CCIP: {
			Networks: []string{"SIMULATED_1", "SIMULATED_2"},
			Chainlink: &Chainlink{
				Common: &Node{
					Name:       "node1",
					Image:      "chainlink",
					Tag:        "latest",
					NodeConfig: tomlString,
					DBImage:    "postgres",
					DBTag:      "latest",
				},
			},
			Groups: map[string]Group{
				"load": &CCIPTestConfig{
					KeepEnvAlive:           utils.Ptr(false),
					MsgType:                "WithToken",
					PhaseTimeout:           models.MustNewDuration(50 * time.Minute),
					TestDuration:           models.MustNewDuration(10 * time.Minute),
					LocalCluster:           utils.Ptr(false),
					ExistingDeployment:     utils.Ptr(false),
					ReuseContracts:         utils.Ptr(true),
					SequentialLaneAddition: utils.Ptr(false),
					NodeFunding:            10,
					RequestPerUnitTime:     []int64{2},
					TimeUnit:               models.MustNewDuration(1 * time.Second),
					NetworkPairs:           []string{"SIMULATED_1,SIMULATED_2", "SIMULATED_2,SIMULATED_1"},
					NoOfNetworks:           2,
					NoOfLanesPerPair:       1,
				},
				"smoke": &CCIPTestConfig{
					KeepEnvAlive:           utils.Ptr(false),
					MsgType:                "WithToken",
					PhaseTimeout:           models.MustNewDuration(10 * time.Minute),
					LocalCluster:           utils.Ptr(true),
					ExistingDeployment:     utils.Ptr(false),
					ReuseContracts:         utils.Ptr(true),
					SequentialLaneAddition: utils.Ptr(false),
					NodeFunding:            10,
					NoOfNetworks:           2,
					NoOfLanesPerPair:       1,
				},
			},
		},
	}}
	d, err := toml.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(d))
}
