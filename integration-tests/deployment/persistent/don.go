package persistent

import (
	"fmt"
	"github.com/AlekSi/pointer"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/ccip/integration-tests/ccip-tests/testconfig"
	"github.com/smartcontractkit/ccip/integration-tests/ccip-tests/testsetups"
	"github.com/smartcontractkit/ccip/integration-tests/docker/test_env"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	ctfClient "github.com/smartcontractkit/chainlink-testing-framework/client"
	"github.com/smartcontractkit/chainlink-testing-framework/docker"
	ctftestenv "github.com/smartcontractkit/chainlink-testing-framework/docker/test_env"
	"github.com/smartcontractkit/chainlink-testing-framework/logging"
	"github.com/smartcontractkit/chainlink-testing-framework/logstream"
	"github.com/smartcontractkit/chainlink/integration-tests/client"
	"github.com/smartcontractkit/chainlink/integration-tests/deployment"
)

type ExistingDONConfig struct {
	testconfig.CLCluster
}

type NewDONHooks interface {
	PostStartupHook([]*client.ChainlinkClient) error
}

type NewDONConfig struct {
	testconfig.ChainlinkDeployment
	Chains        map[uint64]deployment.Chain
	DockerOptions DockerOptions
	NewDONHooks
}

type DockerOptions struct {
	DockerNetworks []string
	LogStream      *logstream.LogStream
}

type DONConfig struct {
	ExistingDON *ExistingDONConfig
	NewDON      *NewDONConfig
}

type DON struct {
	ClClients []*client.ChainlinkK8sClient
	Keys      map[uint64][]client.NodeKeysBundle
	// we use Mockserver in k8s
	MockServer *ctfClient.MockserverClient
	// we use Killgrave in Docker
	KillGrave *ctftestenv.Killgrave
}

func NewNodes(donConfig DONConfig) (DON, error) {
	if donConfig.NewDON == nil && donConfig.ExistingDON == nil {
		return DON{}, fmt.Errorf("no DON config provided, you need to provide either an existing or new DON config")
	}

	if donConfig.NewDON != nil && donConfig.ExistingDON != nil {
		return DON{}, fmt.Errorf("both new and existing DON config provided, you need to provide either an existing or new DON config")
	}

	//TODO I will need also chain config here

	if donConfig.NewDON != nil {
		return NewDON(*donConfig.NewDON)
	}

	return ExistingNodes(*donConfig.ExistingDON)
}

func ExistingNodes(config ExistingDONConfig) (DON, error) {
	noOfNodes := pointer.GetInt(config.NoOfNodes)
	namespace := pointer.GetString(config.Name)

	don := DON{}

	for i := 0; i < noOfNodes; i++ {
		cfg := config.NodeConfigs[i]
		if cfg == nil {
			return don, fmt.Errorf("node %d config is nil", i+1)
		}
		clClient, err := client.NewChainlinkK8sClient(cfg, cfg.InternalIP, namespace)
		if err != nil {
			return don, errors.Wrapf(err, "failed to create chainlink client: %w for node %d config %v", i+1, cfg)
		}
		clClient.ChainlinkClient.WithRetryCount(3)
		don.ClClients = append(don.ClClients, clClient)

		// TODO no idea if that's required for existing DON
		//ocr2Keys, err := clClient.ChainlinkClient.MustReadOCR2Keys()
		//if err != nil {
		//	return don, errors.Wrapf(err, "failed to read OCR2 keys for node %d", i+1)
		//}
		//
		//p2pKeys, err := clClient.ChainlinkClient.MustReadP2PKeys()
		//if err != nil {
		//	return don, errors.Wrapf(err, "failed to read P2P keys for node %d", i+1)
		//}

		// read peer id somehow

	}

	// per chain, if required
	//txKeys, err := clClient.ChainlinkClient.ReadTxKeys()

	//TODO add mockserver

	return don, nil
}

//TODO how to support non-evm here? Solana, Cosmos, Aptos, Starknet? I guess by passing chainConfig as part of DON config?
//TODO add some hooks, for example for performing other operations on nodes once they started, I see CCIP has shitloads of them
//TODO do not forget about mockserver

// for now we won't support starting a new k8s don, I am not sure we should even ever add it
func NewDON(newDonConfig NewDONConfig) (DON, error) {
	don := DON{}

	// maybe we should validate this and return err if not set instead of generating here
	if len(newDonConfig.DockerOptions.DockerNetworks) == 0 {
		dockerNetwork, err := docker.CreateNetwork(logging.GetLogger(nil, "CORE_DOCKER_ENV_LOG_LEVEL"))
		if err != nil {
			return don, errors.Wrap(err, "failed to create docker network")
		}
		// TODO should we return it?
		newDonConfig.DockerOptions.DockerNetworks = []string{dockerNetwork.Name}
	}

	var evmNetworks []blockchain.EVMNetwork
	for _, chain := range newDonConfig.Chains {
		evmNetwork := chain.EVMNetwork.EVMNetworkData()
		evmNetwork.HTTPURLs = chain.EVMNetwork.PrivateHttpUrls()
		evmNetwork.URLs = chain.EVMNetwork.PrivateWsUrls()
		evmNetworks = append(evmNetworks, evmNetwork)
	}

	clCluster := test_env.ClCluster{}
	noOfNodes := pointer.GetInt(newDonConfig.NoOfNodes)
	// if individual nodes are specified, then deploy them with specified configs
	// TODO probably best to put it in a reusable method, use it here and also in integration-tests/ccip-tests/testsetups/test_env.go
	if len(newDonConfig.Nodes) > 0 {
		for _, clNode := range newDonConfig.Nodes {
			toml, _, err := testsetups.SetNodeConfig(
				evmNetworks,
				clNode.BaseConfigTOML,
				clNode.CommonChainConfigTOML,
				clNode.ChainConfigTOMLByChain,
			)
			if err != nil {
				return don, errors.Wrapf(err, "failed to create node config")
			}

			node, err := test_env.NewClNode(
				newDonConfig.DockerOptions.DockerNetworks,
				pointer.GetString(clNode.ChainlinkImage.Image),
				pointer.GetString(clNode.ChainlinkImage.Version),
				toml,
				newDonConfig.DockerOptions.LogStream,
				test_env.WithPgDBOptions(
					ctftestenv.WithPostgresImageName(clNode.DBImage),
					ctftestenv.WithPostgresImageVersion(clNode.DBTag),
				),
			)
			if err != nil {
				return don, errors.Wrapf(err, "failed to build new chainlink node")
			}
			// node.SetTestLogger(t)
			clCluster.Nodes = append(clCluster.Nodes, node)
		}
	} else {
		// if no individual nodes are specified, then deploy the number of nodes specified in the env input with common config
		for i := 0; i < noOfNodes; i++ {
			toml, _, err := testsetups.SetNodeConfig(
				evmNetworks,
				newDonConfig.Common.BaseConfigTOML,
				newDonConfig.Common.CommonChainConfigTOML,
				newDonConfig.Common.ChainConfigTOMLByChain,
			)
			if err != nil {
				return don, errors.Wrapf(err, "failed to create node config")
			}
			node, err := test_env.NewClNode(
				newDonConfig.DockerOptions.DockerNetworks,
				pointer.GetString(newDonConfig.Common.ChainlinkImage.Image),
				pointer.GetString(newDonConfig.Common.ChainlinkImage.Version),
				toml,
				newDonConfig.DockerOptions.LogStream,
				test_env.WithPgDBOptions(
					ctftestenv.WithPostgresImageName(newDonConfig.Common.DBImage),
					ctftestenv.WithPostgresImageVersion(newDonConfig.Common.DBTag),
				),
			)
			if err != nil {
				return don, errors.Wrapf(err, "failed to build new chainlink node")
			}
			//node.SetTestLogger(t)
			clCluster.Nodes = append(clCluster.Nodes, node)
		}
	}

	//TODO maybe a pre-start hook here?

	startErr := clCluster.Start()
	if startErr != nil {
		return don, errors.Wrap(startErr, "failed to start chainlink cluster")
	}

	var chainlinkNodes []*client.ChainlinkClient
	for _, node := range clCluster.Nodes {
		chainlinkNodes = append(chainlinkNodes, node.API.WithRetryCount(3))
	}

	don.Keys = make(map[uint64][]client.NodeKeysBundle)

	for chainId := range newDonConfig.Chains {
		_, clNodes, err := client.CreateNodeKeysBundle(chainlinkNodes, "evm", fmt.Sprint(chainId))
		if err != nil {
			return don, errors.Wrapf(err, "failed to create node keys for chain %d", chainId)
		}
		don.Keys[chainId] = func() []client.NodeKeysBundle {
			var keys []client.NodeKeysBundle
			for _, clNode := range clNodes {
				keys = append(keys, clNode.KeysBundle)
			}
			return keys
		}()
	}

	for _, clClient := range chainlinkNodes {
		don.ClClients = append(don.ClClients, &client.ChainlinkK8sClient{
			ChainlinkClient: clClient,
		})
	}

	if newDonConfig.NewDONHooks != nil {
		err := newDonConfig.NewDONHooks.PostStartupHook(chainlinkNodes)
		if err != nil {
			return don, errors.Wrap(err, "failed to execute post setup hook")
		}
	}

	don.KillGrave = ctftestenv.NewKillgrave(newDonConfig.DockerOptions.DockerNetworks, "", ctftestenv.WithLogStream(newDonConfig.DockerOptions.LogStream))

	return don, nil
}
