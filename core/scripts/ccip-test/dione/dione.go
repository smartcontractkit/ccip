package dione

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"strings"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/chainlink/integration-tests/client"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2/types"
	null2 "gopkg.in/guregu/null.v4"

	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
	"github.com/smartcontractkit/chainlink/core/scripts/common"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/core/services/job"
)

const (
	JSON_FOLDER  = "json"
	CHAIN_FOLDER = "chain"
	NODES_FOLDER = "nodes"

	PollPeriod = "1s"
)

type Environment string

const (
	StagingAlpha Environment = "staging-alpha"
	StagingBeta  Environment = "staging-beta"
	Production   Environment = "prod"
)

type JobType string

const (
	Relay     JobType = "relay"
	Execution JobType = "exec"
	Boostrap  JobType = "bootstrap"
)

type Chain string

const (
	Rinkeby        Chain = "Rinkeby"
	Goerli         Chain = "Goerli"
	OptimismGoerli Chain = "420"
	Sepolia        Chain = "Sepolia"
	AvaxFuji       Chain = "Avax Fuji"
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
	Nodes     []*client.Chainlink
	bootstrap *client.Chainlink
	Config    NodesConfig
	env       Environment
	lggr      logger.Logger
}

func NewDON(env Environment, lggr logger.Logger) DON {
	config := MustReadNodeConfig(env)
	nodes, bootstrap, err := LoadNodes(config)
	common.PanicErr(err)

	return DON{
		Nodes:     nodes,
		bootstrap: bootstrap,
		Config:    config,
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

func (don *DON) FundNodeKeys(chainConfig rhea.EvmChainConfig, ownerPrivKey string, amount *big.Int) {
	nonce, err := chainConfig.Client.PendingNonceAt(context.Background(), chainConfig.Owner.From)
	helpers.PanicErr(err)
	var gasTipCap *big.Int
	if chainConfig.GasSettings.EIP1559 {
		gasTipCap, err = chainConfig.Client.SuggestGasTipCap(context.Background())
		helpers.PanicErr(err)
	}
	gasPrice, err := chainConfig.Client.SuggestGasPrice(context.Background())
	helpers.PanicErr(err)

	ownerKey, err := crypto.HexToECDSA(ownerPrivKey)
	helpers.PanicErr(err)

	for i, node := range don.Config.Nodes {
		to := gethcommon.HexToAddress(node.EthKeys[chainConfig.ChainId.String()])
		if to == gethcommon.HexToAddress("0x") {
			don.lggr.Warnf("Node %2d has no sending key configured. Skipping funding")
			continue
		}
		if chainConfig.GasSettings.EIP1559 {
			sendEthEIP1559(to, chainConfig, nonce+uint64(i), gasTipCap, ownerKey, amount)
		} else {
			sendEth(to, chainConfig, nonce+uint64(i), gasPrice, ownerKey, amount)
		}
		don.lggr.Infof("Sent %s wei to %s", amount.String(), to.Hex())
	}
}

func sendEth(to gethcommon.Address, chainConfig rhea.EvmChainConfig, nonce uint64, gasPrice *big.Int, ownerKey *ecdsa.PrivateKey, amount *big.Int) {
	tx := types.NewTx(
		&types.LegacyTx{
			Nonce:    nonce,
			GasPrice: gasPrice,
			Gas:      21_000,
			To:       &to,
			Value:    amount,
			Data:     []byte{},
		},
	)

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainConfig.ChainId), ownerKey)
	helpers.PanicErr(err)
	err = chainConfig.Client.SendTransaction(context.Background(), signedTx)
	helpers.PanicErr(err)
}

func sendEthEIP1559(to gethcommon.Address, chainConfig rhea.EvmChainConfig, nonce uint64, gasTipCap *big.Int, ownerKey *ecdsa.PrivateKey, amount *big.Int) {
	tx := types.NewTx(
		&types.DynamicFeeTx{
			ChainID:    chainConfig.ChainId,
			Nonce:      nonce,
			GasTipCap:  gasTipCap,
			GasFeeCap:  big.NewInt(2e9),
			Gas:        uint64(21_000),
			To:         &to,
			Value:      amount,
			Data:       []byte{},
			AccessList: types.AccessList{},
		},
	)

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainConfig.ChainId), ownerKey)
	helpers.PanicErr(err)
	err = chainConfig.Client.SendTransaction(context.Background(), signedTx)
	helpers.PanicErr(err)
}

func (don *DON) PopulateOCR2Keys() {
	for i, node := range don.Nodes {
		keys, _, err := node.ReadOCR2Keys()
		common.PanicErr(err)
		don.Config.Nodes[i].OCRKeys = *keys
	}
}

func createKey(c *client.Chainlink, chain string) (*http.Response, error) {
	createUrl := url.URL{
		Path: "/v2/keys/evm",
	}
	query := createUrl.Query()
	query.Set("evmChainID", chain)

	createUrl.RawQuery = query.Encode()
	resp, err := c.APIClient.R().Post(createUrl.String())
	if err != nil {
		return nil, err
	}

	return resp.RawResponse, nil
}

func deleteKnownETHKey(node *client.Chainlink, key string) (*http.Response, error) {
	deleteUrl := url.URL{
		Path: "/v2/keys/evm/" + key,
	}
	query := deleteUrl.Query()
	query.Set("hard", "true")
	deleteUrl.RawQuery = query.Encode()

	resp, err := node.APIClient.R().
		Delete(deleteUrl.String())
	if err != nil {
		return nil, err
	}
	return resp.RawResponse, err
}

func (don *DON) DeleteKnownKey(chainID string) {
	for i, node := range don.Nodes {
		// Only remove a key if it exists
		if key, ok := don.Config.Nodes[i].EthKeys[chainID]; ok {
			resp, err := deleteKnownETHKey(node, key)
			if err != nil {
				don.lggr.Infof("Failed to delete key: %s", resp.Status)
			}
		}
	}
}

func (don *DON) CreateNewEthKeysForChain(chainID *big.Int) {
	for i, node := range don.Nodes {
		_, err := createKey(node, chainID.String())
		common.PanicErr(err)
		don.lggr.Infof("Node [%2d] Created new eth key", i)
	}
}

func (don *DON) PopulatePeerId() {
	for i, node := range don.Nodes {
		p2pkeys, err := node.MustReadP2PKeys()
		common.PanicErr(err)

		don.Config.Nodes[i].PeerID = p2pkeys.Data[0].Attributes.PeerID
	}

	p2pkeys, err := don.bootstrap.MustReadP2PKeys()
	common.PanicErr(err)
	don.Config.Bootstrap.PeerID = p2pkeys.Data[0].Attributes.PeerID
}

func (don *DON) PopulateEthKeys() {
	for i, node := range don.Nodes {
		keys, err := node.MustReadETHKeys()
		if err != nil {
			don.lggr.Infof("Failed getting keys for node %d", i)
		}

		don.Config.Nodes[i].EthKeys = make(map[string]string)
		don.lggr.Infof("Read %d keys for node %2d", len(keys.Data), i)
		for _, key := range keys.Data {
			don.Config.Nodes[i].EthKeys[key.Attributes.ChainID] = key.Attributes.Address
		}
	}
}

func (don *DON) ClearJobSpecs(jobType JobType, source Chain, destination Chain) {
	jobToDelete := fmt.Sprintf("ccip-%s-%s-%s", jobType, source, destination)

	for i, node := range don.Nodes {
		jobs, _, err := node.ReadJobs()
		common.PanicErr(err)

		for _, maps := range jobs.Data {
			jb := maps["attributes"].(map[string]interface{})
			jobName := jb["name"].(string)
			id := maps["id"].(string)

			don.lggr.Infof("Node [%2d]: Job %s: %s", i, id, jobName)

			if jobToDelete == jobName {
				don.lggr.Infof("Node [%2d]:Deleting job %s: %s", i, id, jobName)
				s, err := node.DeleteJob(id)
				common.PanicErr(err)
				don.lggr.Infof(s.Status)
			}
		}
	}
}

func (don *DON) ListJobSpecs() {
	for i, node := range don.Nodes {
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

	don.lggr.Infof("Created job %3s. Status code %s", jb.Data.ID, tx.Status)
}

func (don *DON) GetBootstrapPeerID() string {
	return fmt.Sprintf("%s@%s:5001", don.Config.Bootstrap.PeerID, strings.TrimSuffix(strings.TrimPrefix(don.bootstrap.Config.URL, "https://"), "/"))
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

func (don *DON) AddTwoWaySpecs(chainA rhea.EvmChainConfig, chainB rhea.EvmChainConfig) {
	bootstrapPeerID := don.GetBootstrapPeerID()
	relaySpecAB := generateRelayJobSpecs(&chainA, &chainB, bootstrapPeerID)
	don.AddJobSpecs(relaySpecAB)
	executionSpecAB := generateExecutionJobSpecs(&chainA, &chainB, bootstrapPeerID)
	don.AddJobSpecs(executionSpecAB)
	relaySpecBA := generateRelayJobSpecs(&chainB, &chainA, bootstrapPeerID)
	don.AddJobSpecs(relaySpecBA)
	executionSpecBA := generateExecutionJobSpecs(&chainB, &chainA, bootstrapPeerID)
	don.AddJobSpecs(executionSpecBA)
}

func (don *DON) AddJobSpecs(spec job.Job) {
	chainID := spec.OCR2OracleSpec.RelayConfig["chainID"].(string)

	for i, node := range don.Nodes {
		evmKeyBundle := GetOCRkeysForChainType(don.Config.Nodes[i].OCRKeys, "evm")
		transmitterIDs := don.Config.Nodes[i].EthKeys

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
	}
}

func CreateJob(node *client.Chainlink, spec job.Job) (*client.Job, *http.Response, error) {
	job := &client.Job{}
	specString := RelaySpecToString(spec)

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

func RelaySpecToString(spec job.Job) string {
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

[pluginConfig]
sourceChainID      = %s
destChainID        = %s
onRampIDs          = ["%s"]
pollPeriod         = "%s"
SourceStartBlock   = %d
DestStartBlock     = %d

[relayConfig]
chainID            = %s
`

	return fmt.Sprintf(relayTemplate+"\n",
		spec.Name.String,
		spec.OCR2OracleSpec.ContractID,
		spec.OCR2OracleSpec.OCRKeyBundleID.String,
		spec.OCR2OracleSpec.TransmitterID.String,
		spec.OCR2OracleSpec.PluginConfig["sourceChainID"],
		spec.OCR2OracleSpec.PluginConfig["destChainID"],
		onRamp,
		PollPeriod,
		spec.OCR2OracleSpec.PluginConfig["SourceStartBlock"],
		spec.OCR2OracleSpec.PluginConfig["DestStartBlock"],
		spec.OCR2OracleSpec.PluginConfig["destChainID"],
	)
}

func ExecSpecToString(spec job.Job) string {
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

[pluginConfig]
sourceChainID      = %s
destChainID        = %s
onRampID           = "%s"
blobVerifierID     = "%s"
SourceStartBlock   = %d
DestStartBlock     = %d
tokensPerFeeCoinPipeline = %s

[relayConfig]
chainID            = %s
`

	return fmt.Sprintf(relayTemplate+"\n",
		spec.Name.String,
		spec.OCR2OracleSpec.ContractID,
		spec.OCR2OracleSpec.OCRKeyBundleID.String,
		spec.OCR2OracleSpec.TransmitterID.String,
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

	for _, node := range don.Config.Nodes {
		evmKeys := GetOCRkeysForChainType(node.OCRKeys, "evm")

		oracles = append(oracles,
			confighelper2.OracleIdentityExtra{
				OracleIdentity: confighelper2.OracleIdentity{
					TransmitAccount:   ocr2types.Account(node.EthKeys[chain]),
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
	file, err := json.MarshalIndent(don.Config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(jsonFile+".new.json", file, 0600)
}

func (don *DON) PrintConfig() {
	file, err := json.MarshalIndent(don.Config, "", "  ")
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

func generateRelayJobSpecs(sourceClient *rhea.EvmChainConfig, destClient *rhea.EvmChainConfig, bootstrapPeer string) job.Job {
	return job.Job{
		Name: null2.StringFrom(fmt.Sprintf("ccip-relay-%s-%s", helpers.ChainName(sourceClient.ChainId.Int64()), helpers.ChainName(destClient.ChainId.Int64()))),
		Type: "offchainreporting2",
		OCR2OracleSpec: &job.OCR2OracleSpec{
			PluginType:                  job.CCIPRelay,
			ContractID:                  destClient.BlobVerifier.Hex(),
			Relay:                       "evm",
			RelayConfig:                 map[string]interface{}{"chainID": destClient.ChainId.String()},
			P2PV2Bootstrappers:          []string{bootstrapPeer},
			OCRKeyBundleID:              null2.String{}, // Set per node
			TransmitterID:               null2.String{}, // Set per node
			ContractConfigConfirmations: 2,
			PluginConfig: map[string]interface{}{
				"sourceChainID":    sourceClient.ChainId.String(),
				"destChainID":      destClient.ChainId.String(),
				"onRampIDs":        []string{sourceClient.OnRamp.String()},
				"pollPeriod":       PollPeriod,
				"SourceStartBlock": sourceClient.DeploySettings.DeployedAt,
				"DestStartBlock":   destClient.DeploySettings.DeployedAt,
			},
		},
	}
}

func generateExecutionJobSpecs(sourceClient *rhea.EvmChainConfig, destClient *rhea.EvmChainConfig, bootstrapPeer string) job.Job {
	return job.Job{
		Name: null2.StringFrom(fmt.Sprintf("ccip-exec-%s-%s", helpers.ChainName(sourceClient.ChainId.Int64()), helpers.ChainName(destClient.ChainId.Int64()))),
		Type: "offchainreporting2",
		OCR2OracleSpec: &job.OCR2OracleSpec{
			PluginType:                  job.CCIPExecution,
			ContractID:                  destClient.OffRamp.Hex(),
			Relay:                       "evm",
			RelayConfig:                 map[string]interface{}{"chainID": destClient.ChainId.String()},
			P2PV2Bootstrappers:          []string{bootstrapPeer},
			OCRKeyBundleID:              null2.String{}, // Set per node
			TransmitterID:               null2.String{}, // Set per node
			ContractConfigConfirmations: 2,
			PluginConfig: map[string]interface{}{
				"sourceChainID":            sourceClient.ChainId.String(),
				"destChainID":              destClient.ChainId.String(),
				"onRampID":                 sourceClient.OnRamp.String(),
				"pollPeriod":               PollPeriod,
				"blobVerifierID":           destClient.BlobVerifier.Hex(),
				"SourceStartBlock":         sourceClient.DeploySettings.DeployedAt,
				"DestStartBlock":           destClient.DeploySettings.DeployedAt,
				"tokensPerFeeCoinPipeline": fmt.Sprintf(`"""merge [type=merge left="{}" right="{\\\"%s\\\":\\\"1000000000000000000\\\"}"];"""`, destClient.LinkToken.Hex()),
			},
		},
	}
}
