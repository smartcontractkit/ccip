package dione

import (
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/smartcontractkit/chainlink/integration-tests/client"

	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
	"github.com/smartcontractkit/chainlink/core/scripts/common"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
)

const (
	PollPeriod = time.Second
)

type Environment string

const (
	StagingAlpha Environment = "staging-alpha"
	StagingBeta  Environment = "staging-beta"
	Production   Environment = "prod"
)

type JobType string

const (
	Commit    JobType = "commit"
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

type NodesConfig struct {
	Bootstrap NodeConfig
	Nodes     []NodeConfig
}

type NodeConfig struct {
	EthKeys map[string]string
	PeerID  string
	OCRKeys client.OCR2Keys
}

type DON struct {
	Nodes     []*client.Chainlink
	bootstrap *client.Chainlink
	OfflineDON
}

func NewDON(env Environment, lggr logger.Logger) DON {
	creds, err := ReadCredentials(env)
	common.PanicErr(err)
	nodes, bootstrap, err := creds.DialNodes()
	common.PanicErr(err)

	return DON{
		Nodes:      nodes,
		bootstrap:  bootstrap,
		OfflineDON: NewOfflineDON(env, lggr),
	}
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

func (don *DON) ClearJobSpecs(jobType JobType, source string, destination string) {
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

func (don *DON) LoadCurrentNodeParams() {
	don.PopulateOCR2Keys()
	don.PopulateEthKeys()
	don.PopulatePeerId()
	don.PrintConfig()
}

func (don *DON) ClearAllJobs(chainA string, chainB string) {
	don.ClearJobSpecs(Commit, chainA, chainB)
	don.ClearJobSpecs(Execution, chainA, chainB)
	don.ClearJobSpecs(Commit, chainB, chainA)
	don.ClearJobSpecs(Execution, chainB, chainA)
}

func (don *DON) AddTwoWaySpecs(chainA rhea.EvmDeploymentConfig, chainB rhea.EvmDeploymentConfig) {
	jobParamsAB := NewCCIPJobSpecParams(chainA, chainB)
	relaySpecAB, err := jobParamsAB.CommitJobSpec()
	if err != nil {
		don.lggr.Errorf("commit jobspec error %v", err)
	}
	don.AddJobSpec(relaySpecAB)
	executionSpecAB, err := jobParamsAB.ExecutionJobSpec()
	if err != nil {
		don.lggr.Errorf("exec jobspec error %v", err)
	}
	don.AddJobSpec(executionSpecAB)
	jobParamsBA := NewCCIPJobSpecParams(chainB, chainA)
	relaySpecBA, err := jobParamsBA.CommitJobSpec()
	if err != nil {
		don.lggr.Errorf("commit jobspec error %v", err)
	}
	don.AddJobSpec(relaySpecBA)
	executionSpecBA, err := jobParamsBA.ExecutionJobSpec()
	if err != nil {
		don.lggr.Errorf("exec jobspec error %v", err)
	}
	don.AddJobSpec(executionSpecBA)
}

func (don *DON) AddMissingSpecs(chainA rhea.EvmDeploymentConfig, chainB rhea.EvmDeploymentConfig) {
	jobsAdded := 0
	for i, node := range don.Nodes {
		jobs, _, err := node.ReadJobs()
		common.PanicErr(err)

		lookingForCommit := fmt.Sprintf("ccip-%s-%s-%s", Commit, helpers.ChainName(int64(chainA.ChainConfig.ChainId)), helpers.ChainName(int64(chainB.ChainConfig.ChainId)))
		lookingForExec := fmt.Sprintf("ccip-%s-%s-%s", Execution, helpers.ChainName(int64(chainA.ChainConfig.ChainId)), helpers.ChainName(int64(chainB.ChainConfig.ChainId)))
		don.lggr.Infof("Checking node #%d for [%s] and ", i, lookingForCommit)

		commitFound, execFound := false, false
		for _, maps := range jobs.Data {
			jb := maps["attributes"].(map[string]interface{})
			jobName := jb["name"].(string)

			if jobName == lookingForCommit {
				commitFound = true
			}
			if jobName == lookingForExec {
				execFound = true
			}
		}
		jobParamsAB := NewCCIPJobSpecParams(chainA, chainB)

		if !commitFound {
			don.lggr.Infof("Found missing job [%s] on node #%d", lookingForCommit, i)

			relaySpecAB, err := jobParamsAB.CommitJobSpec()
			if err != nil {
				don.lggr.Errorf("commit jobspec error %v", err)
			}
			don.AddSingleJob(node, relaySpecAB, i)
			jobsAdded++
		}

		if !execFound {
			don.lggr.Infof("Found missing job [%s] on node #%d", lookingForExec, i)
			executionSpecAB, err := jobParamsAB.ExecutionJobSpec()
			if err != nil {
				don.lggr.Errorf("exec jobspec error %v", err)
			}
			don.AddSingleJob(node, executionSpecAB, i)
			jobsAdded++
		}
	}
	don.lggr.Infof("Added %d missing jobs", jobsAdded)

}

func (don *DON) AddJobSpec(spec *client.OCR2TaskJobSpec) {
	for i, node := range don.Nodes {
		don.AddSingleJob(node, spec, i)
	}
}

func (don *DON) AddSingleJob(node *client.Chainlink, spec *client.OCR2TaskJobSpec, nodeIndex int) {
	chainID := spec.OCR2OracleSpec.RelayConfig["chainID"].(uint64)
	evmKeyBundle := GetOCRkeysForChainType(don.Config.Nodes[nodeIndex].OCRKeys, "evm")
	transmitterIDs := don.Config.Nodes[nodeIndex].EthKeys

	// set node specific values
	spec.OCR2OracleSpec.OCRKeyBundleID.SetValid(evmKeyBundle.ID)
	spec.OCR2OracleSpec.TransmitterID.SetValid(transmitterIDs[strconv.FormatUint(chainID, 10)])

	specString, err := spec.String()
	common.PanicErr(err)

	don.lggr.Infof(specString)
	jb, tx, err := node.CreateJobRaw(specString)
	common.PanicErr(err)

	don.lggr.Infof("Created job %3s. Status code %s", jb.Data.ID, tx.Status)
}
