package smoke_test

import (
	"testing"

	"github.com/smartcontractkit/chainlink-env/environment"
	"github.com/smartcontractkit/chainlink-env/pkg/helm/chainlink"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/networks"
	"github.com/smartcontractkit/chainlink-testing-framework/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/smartcontractkit/ccip/integration-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/client"
)

func TestReturnFunds(t *testing.T) {
	envName := "smoke-ccip-38703"
	env := environment.New(
		&environment.Config{
			Namespace:        envName,
			NoManifestUpdate: true,
		})
	err := env.AddHelm(chainlink.New(0, nil)).Run()
	require.NoError(t, err, "error running env")
	chainlinkNodes, err := client.ConnectChainlinkNodes(env)
	require.NoError(t, err, "error connecting to chainlink nodes")
	for _, node := range chainlinkNodes {
		jobs, _, _ := node.ReadJobs()
		for _, maps := range jobs.Data {
			id := maps["id"].(string)
			node.DeleteJob(id)
		}
	}
	var chains []blockchain.EVMClient
	for _, network := range networks.SelectedNetworks[4:] {
		if network.Simulated {
			continue
		}
		ec, err := blockchain.NewEVMClient(network, env)
		require.NoError(t, err, "Connecting to blockchain nodes shouldn't fail")
		chains = append(chains, ec)
	}
	err = actions.TeardownSuite(t, env, utils.ProjectRoot, chainlinkNodes, nil,
		zapcore.ErrorLevel, chains...)
	require.NoError(t, err, "Environment teardown shouldn't fail")
}
