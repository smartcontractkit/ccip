package test_env

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ks "github.com/kurtosis-tech/kurtosis/api/golang/core/lib/services"
	"github.com/kurtosis-tech/kurtosis/api/golang/core/lib/starlark_run_config"
	"github.com/kurtosis-tech/kurtosis/api/golang/engine/lib/kurtosis_context"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/docker/test_env"
	"github.com/smartcontractkit/chainlink-testing-framework/logging"
	tc "github.com/testcontainers/testcontainers-go"
	"gopkg.in/yaml.v2"
)

var kurtosisPackage = "github.com/Tofel/ethereum-package/"

type KurtosisEvmNode struct {
	ExternalHttpUrl   string
	InternalHttpUrl   string
	ExternalWsUrl     string
	InternalWsUrl     string
	EVMClient         blockchain.EVMClient
	EVMNetwork        blockchain.EVMNetwork
	DockerNetworkName string
	configFile        string
	t                 *testing.T
	l                 zerolog.Logger
	tc                *tc.Container
}

func (n *KurtosisEvmNode) GetInternalHttpUrl() string {
	return n.InternalHttpUrl
}

func (n *KurtosisEvmNode) GetInternalWsUrl() string {
	return n.InternalWsUrl
}

func (n *KurtosisEvmNode) GetEVMClient() blockchain.EVMClient {
	return n.EVMClient
}

func (n *KurtosisEvmNode) WithTestLogger(t *testing.T) test_env.NonDevNode {
	n.t = t
	n.l = logging.GetTestLogger(t)
	return n
}

func (n *KurtosisEvmNode) Start() error {
	kurtosisCtx, err := kurtosis_context.NewKurtosisContextFromLocalEngine()
	if err != nil {
		return err
	}
	ctx := context.Background()

	chainConfig, err := n.readKurtosisConfig(n.configFile)
	if err != nil {
		return err
	}

	enclaveName := strings.Replace(n.configFile, ".yaml", "", 1)
	enclaveCtx, err := kurtosisCtx.CreateEnclave(ctx, enclaveName)
	if err != nil {
		return err
	}

	packageArgsFileBytes, err := os.ReadFile(n.configFile)
	if err != nil {
		return err
	}

	packageArgs := string(packageArgsFileBytes)

	if err := n.validateSerializedArgs(packageArgs); err != nil {
		return err
	}

	starlarkRunConfig := starlark_run_config.NewRunStarlarkConfig(starlark_run_config.WithSerializedParams(packageArgs))
	// _, err = enclaveCtx.RunStarlarkRemotePackageBlocking(ctx, kurtosisPackage, starlarkRunConfig)
	_, err = enclaveCtx.RunStarlarkPackageBlocking(ctx, "/Users/btofel/Desktop/repos/ethereum-package", starlarkRunConfig)
	if err != nil {
		return err
	}
	// fmt.Print(starlarkRunResult.RunOutput)

	services, err := enclaveCtx.GetExistingAndHistoricalServiceIdentifiers(ctx)
	if err != nil {
		return err
	}
	rpcs := make(map[string]map[string]string)

	elNodeFound := false
	containerName := ""

	var primaryElNodeKey string
	if len(chainConfig.Participants) != 1 {
		return fmt.Errorf("exactly one participant is supported in kurtosis chain config, but %d were found", len(chainConfig.Participants))
	}

	// TODO here we could add support for a situation, when we have different network participants running different clients
	// which could get complicated, because each participant could have multiple nodes (bootstrap, primary, the rest)
	if chainConfig.Participants[0].Count == 1 {
		primaryElNodeKey = "el-1"
	} else {
		primaryElNodeKey = "el-2"
	}

	for _, serviceName := range services.GetOrderedListOfNames() {
		if strings.Contains(serviceName, primaryElNodeKey) {
			elNodeFound = true
			serviceCtx, err := enclaveCtx.GetServiceContext(serviceName)
			if err != nil {
				return err
			}

			serviceUUID := string(serviceCtx.GetServiceUUID())
			containerName = fmt.Sprintf("%s--%s", serviceName, serviceUUID)

			publicPorts := serviceCtx.GetPublicPorts()
			privatePorts := serviceCtx.GetPrivatePorts()

			// assume same key is used to represent ws port in public and private ports map
			for _, key := range wsKeys {
				if _, ok := publicPorts[key]; ok {
					wsKey = key
					break
				}
			}

			if wsKey == "" {
				return errors.New(fmt.Sprintf("failed to find any key representing ws public port in %v", publicPorts))
			}

			rpcs["private"] = make(map[string]string)
			rpcs["public"] = make(map[string]string)

			getPorts := func(portMap map[string]*ks.PortSpec) (httpPort, wsPort uint16, e error) {
				if v, ok := portMap[wsKey]; ok {
					wsPort = v.GetNumber()
				} else {
					e = fmt.Errorf("no public port for %s key and service %s", wsKey, serviceName)
					return
				}
				// if there's no rpc/http port specified it means it is the same as ws port
				if v, ok := portMap[rpcKey]; ok {
					httpPort = v.GetNumber()
					return
				} else {
					if v, ok := portMap[wsKey]; ok {
						httpPort = v.GetNumber()
						return
					} else {
						e = fmt.Errorf("no public port for %s key and service %s", wsKey, serviceName)
						return
					}
				}
			}

			//TODO for external host we'd need some method to get it reliably, like we do in case of testcontainers-go
			host := "localhost"

			getPrivateEndpoint := func(protocol string, port uint16) string {
				return fmt.Sprintf("%s://%s--%s:%d", protocol, serviceName, serviceUUID, port)
			}

			getPublicEndpoint := func(protocol string, port uint16) string {
				return fmt.Sprintf("%s://%s:%d", protocol, host, port)
			}

			httpPort, wsPort, err := getPorts(publicPorts)
			if err != nil {
				return err
			}

			rpcs["public"][wsKey] = getPublicEndpoint("ws", wsPort)
			rpcs["public"][rpcKey] = getPublicEndpoint("http", httpPort)

			httpPort, wsPort, err = getPorts(privatePorts)
			if err != nil {
				return err
			}

			rpcs["private"][wsKey] = getPrivateEndpoint("ws", wsPort)
			rpcs["private"][rpcKey] = getPrivateEndpoint("http", httpPort)
		}
	}

	if !elNodeFound {
		return errors.New("no execution layer node found")
	}

	chainId, err := strconv.Atoi(chainConfig.NetworkParams.NetworkId)
	if err != nil {
		return err
	}

	n.ExternalHttpUrl = rpcs["public"][rpcKey]
	n.InternalHttpUrl = rpcs["private"][rpcKey]
	n.ExternalWsUrl = rpcs["public"][wsKey]
	n.InternalWsUrl = rpcs["private"][wsKey]

	n.EVMNetwork = blockchain.EVMNetwork{
		Name:            enclaveName,
		ChainID:         int64(chainId),
		URLs:            []string{rpcs["private"][wsKey]},
		HTTPURLs:        []string{rpcs["private"][rpcKey]},
		FinalityTag:     true,
		Simulated:       false,
		SupportsEIP1559: false,
		Timeout: blockchain.JSONStrDuration{
			Duration: 4 * time.Minute,
		},
		SimulationType: "kurtosis",
		// one of the hardcoded keys
		PrivateKeys: []string{"bcdf20249abf0ed6d944c0288fad489e33f66b3960d9e6229c1cd214ed3bbe31"},
	}
	n.DockerNetworkName = fmt.Sprintf("kt-%s", enclaveName)

	// Reuse existing container to wrap it in testcontainers-go
	req := tc.GenericContainerRequest{
		ContainerRequest: tc.ContainerRequest{
			Name: containerName,
		},
		Reuse:        true,
		ProviderType: tc.ProviderDocker,
	}

	container, err := tc.GenericContainer(context.Background(), req)
	if err != nil {
		return err
	}

	n.tc = &container

	return nil
}

func (n *KurtosisEvmNode) ConnectToClient() error {
	networkCfg := n.EVMNetwork
	networkCfg.URLs = []string{n.ExternalWsUrl}
	networkCfg.HTTPURLs = []string{n.ExternalHttpUrl}
	networkCfg.PrivateKeys = n.EVMNetwork.PrivateKeys

	ec, err := blockchain.NewEVMClientFromNetwork(networkCfg, n.l)
	if err != nil {
		return err
	}

	_, err = ec.BalanceAt(context.Background(), common.HexToAddress("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"))
	if err != nil {
		return err
	}
	n.EVMClient = ec
	//TODO add support for multiclient here? like we do in case of Besu?

	// to make sure all the pending txs are done
	err = ec.WaitForEvents()
	if err != nil {
		return err
	}

	return nil
}

type kurtosisParticipant struct {
	Count int `yaml:"count"`
}

type kurtosisConfig struct {
	Participants  []kurtosisParticipant `yaml:"participants"`
	NetworkParams struct {
		NetworkId string `yaml:"network_id"`
	} `yaml:"network_params"`
}

var (
	rpcKey = "rpc"
	wsKey  string
	wsKeys = []string{"ws-rpc", "ws"}
)

func (c *KurtosisEvmNode) readKurtosisConfig(filePath string) (kurtosisConfig, error) {
	var network kurtosisConfig

	// Open YAML file
	file, err := os.Open(filePath)
	if err != nil {
		return network, err
	}
	defer file.Close()

	if file != nil {
		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(&network); err != nil {
			return network, err
		}
	}

	return network, nil
}

func (c *KurtosisEvmNode) validateSerializedArgs(serializedArgs string) error {
	var result interface{}
	var jsonError error
	if jsonError = json.Unmarshal([]byte(serializedArgs), &result); jsonError == nil {
		return nil
	}
	var yamlError error
	if yamlError = yaml.Unmarshal([]byte(serializedArgs), &result); yamlError == nil {
		return nil
	}
	return fmt.Errorf("invalid serialized args: %s", jsonError.Error()+yamlError.Error())
}

type KurtosisPrivateChain struct {
	blockchain.EVMNetwork
	configFile  string
	primaryNode *KurtosisEvmNode
}

func (k *KurtosisPrivateChain) GetPrimaryNode() test_env.NonDevNode {
	return k.primaryNode
}

func (k *KurtosisPrivateChain) GetNodes() []test_env.NonDevNode {
	return []test_env.NonDevNode{k.primaryNode}
}
func (k *KurtosisPrivateChain) GetNetworkConfig() *blockchain.EVMNetwork {
	return &k.EVMNetwork
}
func (k *KurtosisPrivateChain) GetDockerNetworks() []string {
	return []string{fmt.Sprintf("kt-%s", k.Name)}
}

func NewPrivateKurtosisChain(configFile string) test_env.PrivateChain {
	return &KurtosisPrivateChain{
		configFile: configFile,
		primaryNode: &KurtosisEvmNode{
			configFile: configFile,
		},
	}
}
