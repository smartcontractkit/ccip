package rhea

import (
	"testing"

	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
)

func UpgradeLaneTwoWay(t *testing.T, source *EvmDeploymentConfig, destination *EvmDeploymentConfig) {
	UpgradeLane(t, source, destination)
	UpgradeLane(t, destination, source)
	source.Logger.Infof("Upgraded Lane %s <--> %s",
		helpers.ChainName(int64(source.ChainConfig.ChainId)),
		helpers.ChainName(int64(destination.ChainConfig.ChainId)),
	)
	prettyPrintLanes(source, destination)
}

func UpgradeLane(t *testing.T, source *EvmDeploymentConfig, destination *EvmDeploymentConfig) {
	if !source.DeploySettings.DeployRamp || !destination.DeploySettings.DeployRamp || !destination.DeploySettings.DeployCommitStore {
		source.Logger.Errorf("Please set \"DeployRamp and DeployCommitStore\" to true for the given EvmChainConfigs and make sure "+
			"the right ones are set. Source: %d, Dest %d", source.ChainConfig.ChainId, destination.ChainConfig.ChainId)
		return
	}

	upgradeOnRamp(t, source, destination)
	upgradeCommitStore(t, source, destination)
	upgradeOffRamp(t, source, destination)

	prettyPrintLanes(source, destination)
}

func upgradeOnRamp(t *testing.T, source *EvmDeploymentConfig, destination *EvmDeploymentConfig) {
	source.Logger.Infof("Upgrading onRamp")
	deployOnRamp(t, source, destination.ChainConfig.ChainId, destination.ChainConfig.SupportedTokens)
	setOnRampOnTokenPools(t, source)

	source.Logger.Info("Please deploy new commit jobs")
}

func upgradeCommitStore(t *testing.T, sourceClient *EvmDeploymentConfig, destClient *EvmDeploymentConfig) {
	destClient.Logger.Infof("Upgrading commit store")
	deployCommitStore(t, destClient, sourceClient.ChainConfig.ChainId, sourceClient.LaneConfig.OnRamp)
	destClient.Logger.Info("Please deploy new commit jobs and set OCR2 config")
}

func upgradeOffRamp(t *testing.T, source *EvmDeploymentConfig, destination *EvmDeploymentConfig) {
	destination.Logger.Infof("Upgrading offRamp")
	deployOffRamp(t, destination, source.ChainConfig.ChainId, source.ChainConfig.SupportedTokens, source.LaneConfig.OnRamp)
	setOffRampOnRouter(t, source.ChainConfig.ChainId, destination)
	setOffRampOnTokenPools(t, destination)

	destination.Logger.Info("Please deploy new execution jobs and set OCR2 config")
}

/*
func removeOffRamp(t *testing.T, destClient *EvmDeploymentConfig, offRampAddress common.Address) {
	// Pause contract
	revokeOffRampOnRouter(t, destClient, offRampAddress)
	revokeOffRampOnTokenPools(t, destClient, offRampAddress)
}
*/
