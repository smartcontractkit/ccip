package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/chainlink/integration-tests/client"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	"github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/core/services/job"
)

const (
	JSON_FOLDER  = "json"
	CHAIN_FOLDER = "chain"
	NODES_FOLDER = "nodes"
)

type Environment string

const (
	Staging    Environment = "staging"
	Production Environment = "prod"
)

type JobType string

const (
	Relay     JobType = "relay"
	Execution JobType = "exec"
	Boostrap  JobType = "bootstrap"
)

type Chain string

const (
	Rinkeby Chain = "Rinkeby"
	Goerli  Chain = "Goerli"
	Sepolia Chain = "Sepolia"
)

type ChainConfig struct {
	ChainID  uint64
	RpcUrl   string
	EIP1559  bool
	GasPrice uint64
}

func ReadChainConfig(chain Chain) (ChainConfig, error) {
	configFile := fmt.Sprintf("%s/%s/%s.json", JSON_FOLDER, CHAIN_FOLDER, chain)
	jsonFile, err := os.Open(configFile)
	if err != nil {
		return ChainConfig{}, err
	}
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return ChainConfig{}, err
	}

	var config ChainConfig
	err = json.Unmarshal(byteValue, &config)
	return config, err
}

type NodesConfig struct {
	Bootstrap NodeConfig
	Nodes     []NodeConfig
}

type NodeConfig struct {
	Authentication client.ChainlinkConfig
	EthKeys        map[string]string
	PeerID         string
	OCRKeys        client.OCR2Keys
}

type DON struct {
	nodes     []*client.Chainlink
	bootstrap *client.Chainlink
	config    NodesConfig
	env       Environment
	lggr      logger.Logger
}

func NewDON(env Environment, lggr logger.Logger) DON {
	config := MustReadNodeConfig(env)
	nodes, bootstrap, err := LoadNodes(config)
	common.PanicErr(err)

	return DON{
		nodes:     nodes,
		bootstrap: bootstrap,
		config:    config,
		env:       env,
		lggr:      lggr,
	}
}

func GetNodesFileLocation(env Environment) string {
	return fmt.Sprintf("%s/%s/%s.json", JSON_FOLDER, NODES_FOLDER, env)
}

func MustReadNodeConfig(env Environment) NodesConfig {
	config, err := ReadNodeConfig(env)
	common.PanicErr(err)
	return config
}

func ReadNodeConfig(env Environment) (NodesConfig, error) {
	configFile := GetNodesFileLocation(env)
	jsonFile, err := os.Open(configFile)
	if err != nil {
		return NodesConfig{}, err
	}
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return NodesConfig{}, err
	}

	var config NodesConfig
	err = json.Unmarshal(byteValue, &config)

	return config, err
}

func LoadNodes(configs NodesConfig) (nodes []*client.Chainlink, bootstrap *client.Chainlink, err error) {
	for _, config := range configs.Nodes {
		chainlinkNode, err2 := client.NewChainlink(&config.Authentication)
		if err2 != nil {
			return []*client.Chainlink{}, &client.Chainlink{}, err2
		}
		nodes = append(nodes, chainlinkNode)
	}

	bootstrap, err = client.NewChainlink(&configs.Bootstrap.Authentication)
	if err != nil {
		return []*client.Chainlink{}, &client.Chainlink{}, err
	}

	return nodes, bootstrap, nil
}

func (don *DON) PopulateOCR2Keys() {
	for i, node := range don.nodes {
		keys, _, err := node.ReadOCR2Keys()
		common.PanicErr(err)
		don.config.Nodes[i].OCRKeys = *keys
	}
}

func (don *DON) WIP() {

}

//func createKey(c *client.Chainlink, chain string) (*http.Response, error) {
//	createUrl := url.URL{
//		Path: "/v2/keys/evm",
//	}
//	query := createUrl.Query()
//	query.Set("evmChainID", chain)
//
//	createUrl.RawQuery = query.Encode()
//
//	resp, err := c.APIClient.R().Get(createUrl.String())
//	if err != nil {
//		return nil, err
//	}
//
//	return resp.RawResponse, nil
//}
//
//func readKey(c *client.Chainlink) (*client.TxKeys, *http.Response, error) {
//	txKeys := &client.TxKeys{}
//	log.Info().Str("Node URL", c.Config.URL).Msg("Reading Tx Keys")
//	resp, err := c.APIClient.R().
//		SetResult(txKeys).
//		Get("/v2/keys/evm")
//	if err != nil {
//		return nil, nil, err
//	}
//	return txKeys, resp.RawResponse, err
//}
//func deleteKey(c *client.Chainlink, keyString string) (*http.Response, error) {
//	log.Info().Str("Node URL", c.Config.URL).Str("ID", keyString).Msg("Deleting Tx Key")
//	resp, err := c.APIClient.R().
//		SetPathParams(map[string]string{
//			"id": keyString,
//		}).
//		Delete("/v2/keys/evm/{id}")
//	if err != nil {
//		return nil, err
//	}
//	return resp.RawResponse, err
//}

//func (cli *Client) DeleteETHKey(c *cli.Context) (err error) {
//if !c.Args().Present() {
//	return cli.errorOut(errors.New("Must pass the address of the key to be deleted"))
//}
//address := c.Args().Get(0)
//deleteUrl := url.URL{
//	Path: "/v2/keys/evm/" + address,
//}
//query := deleteUrl.Query()
//
//if c.Bool("hard") && !confirmAction(c) {
//	return nil
//}

func (don *DON) PopulatePeerId() {
	for i, node := range don.nodes {
		p2pkeys, err := node.MustReadP2PKeys()
		common.PanicErr(err)

		don.config.Nodes[i].PeerID = p2pkeys.Data[0].Attributes.PeerID
	}

	p2pkeys, err := don.bootstrap.MustReadP2PKeys()
	common.PanicErr(err)
	don.config.Bootstrap.PeerID = p2pkeys.Data[0].Attributes.PeerID
}

func (don *DON) PopulateEthKeys() {
	for i, node := range don.nodes {
		keys, err := node.MustReadETHKeys()
		common.PanicErr(err)

		don.config.Nodes[i].EthKeys = make(map[string]string)
		for _, key := range keys.Data {
			don.config.Nodes[i].EthKeys[key.Attributes.ChainID] = key.Attributes.Address
		}
	}
}

func (don *DON) ClearJobSpecs(jobType JobType, source Chain, destination Chain) {
	jobToDelete := fmt.Sprintf("ccip-%s-%s-%s", jobType, source, destination)

	for i, node := range don.nodes {
		jobs, _, err := node.ReadJobs()
		common.PanicErr(err)

		for _, maps := range jobs.Data {
			jb := maps["attributes"].(map[string]interface{})
			jobName := jb["name"].(string)
			id := maps["id"].(string)

			don.lggr.Infof("Node [%2d]: Job %s: %s", i, id, jobName)

			if jobToDelete == jobName {
				don.lggr.Infof("Node [%2d]:Deleting job %s: %s", i, id, jobName)
				_, err = node.DeleteJob(id)
				common.PanicErr(err)
			}
		}
	}
}

func (don *DON) ListJobSpecs() {
	for i, node := range don.nodes {
		jobs, _, err := node.ReadJobs()
		common.PanicErr(err)

		for _, maps := range jobs.Data {
			jb := maps["attributes"].(map[string]interface{})
			jobName := jb["name"].(string)
			id := maps["id"].(string)

			don.lggr.Infof("Node [%2d]: Job %3s: %-28s %+v", i, id, jobName, jb)
		}
	}
}

func (don *DON) AddRawJobSpec(node *client.Chainlink, spec string) {
	jb, tx, err := node.CreateJobRaw(spec)
	common.PanicErr(err)

	if tx.StatusCode == 200 {
		don.lggr.Infof("Created job %3s", jb.Data.ID)
	} else {
		panic(tx.Status)
	}
}

func (don *DON) GetBootstrapPeerID() string {
	return fmt.Sprintf("%s@%s:5001", don.config.Bootstrap.PeerID, strings.TrimSuffix(strings.TrimPrefix(don.bootstrap.Config.URL, "https://"), "/"))
}

func (don *DON) LoadCurrentNodeParams() {
	don.PopulateOCR2Keys()
	don.PopulateEthKeys()
	don.PopulatePeerId()
	don.PrintConfig()
}

func (don *DON) ClearAllJobs(chainA Chain, chainB Chain) {
	don.ClearJobSpecs(Relay, chainA, chainB)
	don.ClearJobSpecs(Execution, chainA, chainB)
	don.ClearJobSpecs(Relay, chainB, chainA)
	don.ClearJobSpecs(Execution, chainB, chainA)
}

func (don *DON) AddTwoWaySpecs(chainA EvmChainConfig, chainB EvmChainConfig) {
	relaySpecAB := generateRelayJobSpecs(&chainA, &chainB, don.GetBootstrapPeerID())
	don.AddJobSpecs(relaySpecAB)
	executionSpecAB := generateExecutionJobSpecs(&chainA, &chainB, don.GetBootstrapPeerID())
	don.AddJobSpecs(executionSpecAB)
	relaySpecBA := generateRelayJobSpecs(&chainB, &chainA, don.GetBootstrapPeerID())
	don.AddJobSpecs(relaySpecBA)
	executionSpecBA := generateExecutionJobSpecs(&chainB, &chainA, don.GetBootstrapPeerID())
	don.AddJobSpecs(executionSpecBA)
}

func (don *DON) AddJobSpecs(spec client.OCR2TaskJobSpec) {
	chainID := spec.OCR2OracleSpec.RelayConfig["chainID"].(string)

	for i, node := range don.nodes {
		evmKeyBundle := GetOCRkeysForChainType(don.config.Nodes[i].OCRKeys, "evm")
		transmitterIDs := don.config.Nodes[i].EthKeys

		spec.OCR2OracleSpec.OCRKeyBundleID.SetValid(evmKeyBundle.ID)
		spec.OCR2OracleSpec.TransmitterID.SetValid(transmitterIDs[chainID])

		var specString string
		if spec.OCR2OracleSpec.PluginType == job.CCIPRelay {
			specString = RelaySpecToString(spec)
		} else {
			specString = ExecSpecToString(spec)
		}
		don.lggr.Infof(specString)
		don.AddRawJobSpec(node, specString)

		//job, tx, err := node.CreateJob(&spec)
		//common.PanicErr(err)
		//
		//if tx.StatusCode == 200 {
		//	don.lggr.Infof("Created job %3s", job.Data.ID)
		//} else {
		//	panic(tx.Status)
		//}

		//return
	}
}

func CreateJob(node *client.Chainlink, spec client.OCR2TaskJobSpec) (*client.Job, *http.Response, error) {
	job := &client.Job{}
	specString := RelaySpecToString(spec)

	log.Info().Str("Node URL", node.Config.URL).Str("Type", spec.Type()).Msg("Creating Job")
	resp, err := node.APIClient.R().
		SetBody(&client.JobForm{
			TOML: specString,
		}).
		SetResult(&job).
		Post("/v2/jobs")
	if err != nil {
		return nil, nil, err
	}
	return job, resp.RawResponse, err
}

func RelaySpecToString(spec client.OCR2TaskJobSpec) string {
	bootstrapper := spec.OCR2OracleSpec.P2PV2Bootstrappers[0]
	onRamp := spec.OCR2OracleSpec.PluginConfig["onRampIDs"].([]string)[0]

	const relayTemplate = `
# CCIPRelaySpec
type               = "offchainreporting2"
name               = "%s"
pluginType         = "ccip-relay"
relay              = "evm"
schemaVersion      = 1
contractID         = "%s"
ocrKeyBundleID     = "%s"
transmitterID      = "%s"
p2pv2Bootstrappers  = ["%s"]

[pluginConfig]
sourceChainID      = %s
destChainID        = %s
onRampIDs          = ["%s"]
pollPeriod         = "%s"
SourceStartBlock   = %d
DestStartBlock     = %d

[relayConfig]
chainID            = "%s"
`

	return fmt.Sprintf(relayTemplate+"\n",
		spec.Name,
		spec.OCR2OracleSpec.ContractID,
		spec.OCR2OracleSpec.OCRKeyBundleID.String,
		spec.OCR2OracleSpec.TransmitterID.String,
		bootstrapper,
		spec.OCR2OracleSpec.PluginConfig["sourceChainID"],
		spec.OCR2OracleSpec.PluginConfig["destChainID"],
		onRamp,
		pollPeriod,
		spec.OCR2OracleSpec.PluginConfig["SourceStartBlock"],
		spec.OCR2OracleSpec.PluginConfig["DestStartBlock"],
		spec.OCR2OracleSpec.PluginConfig["destChainID"],
	)
}

func ExecSpecToString(spec client.OCR2TaskJobSpec) string {
	bootstrapper := spec.OCR2OracleSpec.P2PV2Bootstrappers[0]

	const relayTemplate = `
# CCIPExecutionSpec
type               = "offchainreporting2"
name               = "%s"
pluginType         = "ccip-execution"
relay              = "evm"
schemaVersion      = 1
contractID         = "%s"
ocrKeyBundleID     = "%s"
transmitterID      = "%s"
p2pv2Bootstrappers  = ["%s"]

[pluginConfig]
sourceChainID      = %s
destChainID        = %s
onRampID           = "%s"
blobVerifierID     = "%s"
SourceStartBlock   = %d
DestStartBlock     = %d
tokensPerFeeCoinPipeline = %s

[relayConfig]
chainID            = "%s"
`

	return fmt.Sprintf(relayTemplate+"\n",
		spec.Name,
		spec.OCR2OracleSpec.ContractID,
		spec.OCR2OracleSpec.OCRKeyBundleID.String,
		spec.OCR2OracleSpec.TransmitterID.String,
		bootstrapper,
		spec.OCR2OracleSpec.PluginConfig["sourceChainID"],
		spec.OCR2OracleSpec.PluginConfig["destChainID"],
		spec.OCR2OracleSpec.PluginConfig["onRampID"],
		spec.OCR2OracleSpec.PluginConfig["blobVerifierID"],
		spec.OCR2OracleSpec.PluginConfig["SourceStartBlock"],
		spec.OCR2OracleSpec.PluginConfig["DestStartBlock"],
		spec.OCR2OracleSpec.PluginConfig["tokensPerFeeCoinPipeline"],
		spec.OCR2OracleSpec.PluginConfig["destChainID"],
	)
}

func (don *DON) GenerateOracleIdentities(chain string) []confighelper2.OracleIdentityExtra {
	var oracles []confighelper2.OracleIdentityExtra

	for _, node := range don.config.Nodes {
		evmKeys := GetOCRkeysForChainType(node.OCRKeys, "evm")

		oracles = append(oracles,
			confighelper2.OracleIdentityExtra{
				OracleIdentity: confighelper2.OracleIdentity{
					TransmitAccount:   types.Account(node.EthKeys[chain]),
					OnchainPublicKey:  gethcommon.HexToAddress(strings.TrimPrefix(evmKeys.Attributes.OnChainPublicKey, "ocr2on_evm_")).Bytes(),
					OffchainPublicKey: common.ToOffchainPublicKey("0x" + strings.TrimPrefix(evmKeys.Attributes.OffChainPublicKey, "ocr2off_evm_")),
					PeerID:            node.PeerID,
				},
				ConfigEncryptionPublicKey: common.StringTo32Bytes("0x" + strings.TrimPrefix(evmKeys.Attributes.ConfigPublicKey, "ocr2cfg_evm_")),
			})
	}
	return oracles
}

func (don *DON) WriteConfig() error {
	jsonFile := GetNodesFileLocation(don.env)
	file, err := json.MarshalIndent(don.config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(jsonFile+".new.json", file, 0600)
}

func (don *DON) PrintConfig() {
	file, err := json.MarshalIndent(don.config, "", "  ")
	common.PanicErr(err)

	don.lggr.Infof(string(file))
}

func GetOCRkeysForChainType(OCRKeys client.OCR2Keys, chainType string) client.OCR2KeyData {
	for _, key := range OCRKeys.Data {
		if key.Attributes.ChainType == chainType {
			return key
		}
	}

	panic("Keys not found for chain")
}
