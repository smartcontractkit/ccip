package rhea

import (
	"testing"
)

func UpgradeLane(t *testing.T, sourceClient *EvmDeploymentConfig, destClient *EvmDeploymentConfig) {
	if !sourceClient.DeploySettings.DeployRamp || !destClient.DeploySettings.DeployRamp {
		sourceClient.Logger.Errorf("Please set \"DeployRamp\" to true for the given EvmChainConfigs and make sure "+
			"the right ones are set. Source: %d, Dest %d", sourceClient.ChainConfig.ChainId.Int64(), destClient.ChainConfig.ChainId.Int64())
		return
	}

	upgradeOnRamp(t, sourceClient, destClient)
	upgradeOffRamp(t, sourceClient, destClient)

	PrintContractConfig(sourceClient, destClient)
}

func upgradeOnRamp(t *testing.T, sourceClient *EvmDeploymentConfig, destClient *EvmDeploymentConfig) {
	sourceClient.Logger.Infof("Upgrading onRamp")
	deployOnRamp(t, sourceClient, destClient.ChainConfig.ChainId)
	setOnRampOnTokenPools(t, sourceClient)
	setOnRampOnCommitStore(t, sourceClient, destClient)

	sourceClient.Logger.Info("Please deploy new commit jobs")
}

func upgradeOffRamp(t *testing.T, sourceClient *EvmDeploymentConfig, destClient *EvmDeploymentConfig) {
	destClient.Logger.Infof("Upgrading offRamp")
	deployOffRamp(t, destClient, sourceClient)
	setOffRampRouterOnOffRamp(t, destClient)
	setOffRampOnOffRampRouter(t, destClient)
	setOffRampOnTokenPools(t, destClient)

	destClient.Logger.Info("Please deploy new execution jobs")
}

/*
func removeOffRamp(t *testing.T, destClient *EvmDeploymentConfig, offRampAddress common.Address) {
	// Pause contract
	revokeOffRampOnOffRampRouter(t, destClient, offRampAddress)
	revokeOffRampOnTokenPools(t, destClient, offRampAddress)
}
*/
