package dione

import (
	"fmt"

	null2 "gopkg.in/guregu/null.v4"

	"github.com/smartcontractkit/chainlink/integration-tests/client"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/core/services/job"
)

func CommitSpecToString(spec job.Job) string {
	onRamp := spec.OCR2OracleSpec.PluginConfig["onRampIDs"].([]string)[0]

	const commitTemplate = `
# CCIP Commit spec
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

	return fmt.Sprintf(commitTemplate+"\n",
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
	const execTemplate = `
# CCIP Execution spec
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
commitStoreID     = "%s"
SourceStartBlock   = %d
DestStartBlock     = %d
tokensPerFeeCoinPipeline = %s

[relayConfig]
chainID            = %s
`

	return fmt.Sprintf(execTemplate+"\n",
		spec.Name.String,
		spec.OCR2OracleSpec.ContractID,
		spec.OCR2OracleSpec.OCRKeyBundleID.String,
		spec.OCR2OracleSpec.TransmitterID.String,
		spec.OCR2OracleSpec.PluginConfig["sourceChainID"],
		spec.OCR2OracleSpec.PluginConfig["destChainID"],
		spec.OCR2OracleSpec.PluginConfig["onRampID"],
		spec.OCR2OracleSpec.PluginConfig["commitStoreID"],
		spec.OCR2OracleSpec.PluginConfig["SourceStartBlock"],
		spec.OCR2OracleSpec.PluginConfig["DestStartBlock"],
		spec.OCR2OracleSpec.PluginConfig["tokensPerFeeCoinPipeline"],
		spec.OCR2OracleSpec.PluginConfig["destChainID"],
	)
}

func GetOCRkeysForChainType(OCRKeys client.OCR2Keys, chainType string) client.OCR2KeyData {
	for _, key := range OCRKeys.Data {
		if key.Attributes.ChainType == chainType {
			return key
		}
	}

	panic("Keys not found for chain")
}

func generateCommitJobSpecs(sourceClient *rhea.EvmDeploymentConfig, destClient *rhea.EvmDeploymentConfig) job.Job {
	return job.Job{
		Name: null2.StringFrom(fmt.Sprintf("ccip-relay-%s-%s", helpers.ChainName(sourceClient.ChainConfig.ChainId.Int64()), helpers.ChainName(destClient.ChainConfig.ChainId.Int64()))),
		Type: "offchainreporting2",
		OCR2OracleSpec: &job.OCR2OracleSpec{
			PluginType:                  job.CCIPRelay,
			ContractID:                  destClient.LaneConfig.CommitStore.Hex(),
			Relay:                       "evm",
			RelayConfig:                 map[string]interface{}{"chainID": destClient.ChainConfig.ChainId.String()},
			P2PV2Bootstrappers:          []string{},     // Set in env vars
			OCRKeyBundleID:              null2.String{}, // Set per node
			TransmitterID:               null2.String{}, // Set per node
			ContractConfigConfirmations: 2,
			PluginConfig: map[string]interface{}{
				"sourceChainID":    sourceClient.ChainConfig.ChainId.String(),
				"destChainID":      destClient.ChainConfig.ChainId.String(),
				"onRampIDs":        []string{sourceClient.LaneConfig.OnRamp.String()},
				"pollPeriod":       PollPeriod,
				"SourceStartBlock": sourceClient.DeploySettings.DeployedAt,
				"DestStartBlock":   destClient.DeploySettings.DeployedAt,
			},
		},
	}
}

func generateExecutionJobSpecs(sourceClient *rhea.EvmDeploymentConfig, destClient *rhea.EvmDeploymentConfig) job.Job {
	return job.Job{
		Name: null2.StringFrom(fmt.Sprintf("ccip-exec-%s-%s", helpers.ChainName(sourceClient.ChainConfig.ChainId.Int64()), helpers.ChainName(destClient.ChainConfig.ChainId.Int64()))),
		Type: "offchainreporting2",
		OCR2OracleSpec: &job.OCR2OracleSpec{
			PluginType:                  job.CCIPExecution,
			ContractID:                  destClient.LaneConfig.OffRamp.Hex(),
			Relay:                       "evm",
			RelayConfig:                 map[string]interface{}{"chainID": destClient.ChainConfig.ChainId.String()},
			P2PV2Bootstrappers:          []string{},     // Set in env vars
			OCRKeyBundleID:              null2.String{}, // Set per node
			TransmitterID:               null2.String{}, // Set per node
			ContractConfigConfirmations: 2,
			PluginConfig: map[string]interface{}{
				"sourceChainID":            sourceClient.ChainConfig.ChainId.String(),
				"destChainID":              destClient.ChainConfig.ChainId.String(),
				"onRampID":                 sourceClient.LaneConfig.OnRamp.String(),
				"pollPeriod":               PollPeriod,
				"commitStoreID":            destClient.LaneConfig.CommitStore.Hex(),
				"SourceStartBlock":         sourceClient.DeploySettings.DeployedAt,
				"DestStartBlock":           destClient.DeploySettings.DeployedAt,
				"tokensPerFeeCoinPipeline": fmt.Sprintf(`"""merge [type=merge left="{}" right="{\\\"%s\\\":\\\"1000000000000000000\\\"}"];"""`, destClient.ChainConfig.LinkToken.Hex()),
			},
		},
	}
}
