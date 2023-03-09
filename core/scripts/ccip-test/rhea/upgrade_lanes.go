package rhea

import (
	"testing"

	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
)

func UpgradeLaneTwoWay(t *testing.T, sourceClient *EvmDeploymentConfig, destClient *EvmDeploymentConfig) {
	UpgradeLane(t, sourceClient, destClient)
	UpgradeLane(t, destClient, sourceClient)
	sourceClient.Logger.Infof("Upgraded Lane %s <--> %s",
		helpers.ChainName(int64(sourceClient.ChainConfig.ChainId)),
		helpers.ChainName(int64(destClient.ChainConfig.ChainId)),
	)
	prettyPrintLanes(sourceClient, destClient)
}

func UpgradeLane(t *testing.T, sourceClient *EvmDeploymentConfig, destClient *EvmDeploymentConfig) {
	if !sourceClient.DeploySettings.DeployRamp || !destClient.DeploySettings.DeployRamp || !destClient.DeploySettings.DeployCommitStore {
		sourceClient.Logger.Errorf("Please set \"DeployRamp and DeployCommitStore\" to true for the given EvmChainConfigs and make sure "+
			"the right ones are set. Source: %d, Dest %d", sourceClient.ChainConfig.ChainId, destClient.ChainConfig.ChainId)
		return
	}

	upgradeOnRamp(t, sourceClient, destClient)
	upgradeCommitStore(t, sourceClient, destClient)
	upgradeOffRamp(t, sourceClient, destClient)
}

func upgradeOnRamp(t *testing.T, sourceClient *EvmDeploymentConfig, destClient *EvmDeploymentConfig) {
	sourceClient.Logger.Infof("Upgrading onRamp")
	deployOnRamp(t, sourceClient, destClient.ChainConfig.ChainId, destClient.ChainConfig.SupportedTokens)
	setOnRampOnTokenPools(t, sourceClient)

	sourceClient.Logger.Info("Please deploy new commit jobs")
}

func upgradeCommitStore(t *testing.T, sourceClient *EvmDeploymentConfig, destClient *EvmDeploymentConfig) {
	destClient.Logger.Infof("Upgrading commit store")
	deployCommitStore(t, destClient, sourceClient.ChainConfig.ChainId, sourceClient.LaneConfig.OnRamp)
	destClient.Logger.Info("Please deploy new commit jobs and set OCR2 config")
}

func upgradeOffRamp(t *testing.T, sourceClient *EvmDeploymentConfig, destClient *EvmDeploymentConfig) {
	destClient.Logger.Infof("Upgrading offRamp")
	deployOffRamp(t, destClient, sourceClient.ChainConfig.ChainId, sourceClient.ChainConfig.SupportedTokens, sourceClient.LaneConfig.OnRamp)
	setOffRampOnRouter(t, sourceClient.ChainConfig.ChainId, destClient)
	setOffRampOnTokenPools(t, destClient)

	destClient.Logger.Info("Please deploy new execution jobs and set OCR2 config")
}

/*
func removeOffRamp(t *testing.T, destClient *EvmDeploymentConfig, offRampAddress common.Address) {
	// Pause contract
	revokeOffRampOnRouter(t, destClient, offRampAddress)
	revokeOffRampOnTokenPools(t, destClient, offRampAddress)
}
*/
