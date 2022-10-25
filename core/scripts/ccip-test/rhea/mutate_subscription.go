package rhea

import (
	"testing"
)

func UpgradeLane(t *testing.T, sourceClient *EvmChainConfig, destClient *EvmChainConfig) {
	if !sourceClient.DeploySettings.DeployRamp || !destClient.DeploySettings.DeployRamp {
		sourceClient.Logger.Errorf("Please set \"DeployRamp\" to true for the given EvmChainConfigs and make sure "+
			"the right ones are set. Source: %d, Dest %d", sourceClient.ChainId.Int64(), destClient.ChainId.Int64())
		return
	}

	upgradeOnRamp(t, sourceClient, destClient)
	upgradeOffRamp(t, sourceClient, destClient)

	PrintContractConfig(sourceClient, destClient)
}

func upgradeOnRamp(t *testing.T, sourceClient *EvmChainConfig, destClient *EvmChainConfig) {
	sourceClient.Logger.Infof("Upgrading onRamp")
	deployOnRamp(t, sourceClient, destClient.ChainId)
	setOnRampOnTokenPools(t, sourceClient)
	setOnRampOnBlobVerifier(t, sourceClient, destClient)

	sourceClient.Logger.Info("Please deploy new relay jobs")
}

func upgradeOffRamp(t *testing.T, sourceClient *EvmChainConfig, destClient *EvmChainConfig) {
	destClient.Logger.Infof("Upgrading offRamp")
	deployOffRamp(t, destClient, sourceClient)
	setOffRampRouterOnOffRamp(t, destClient)
	setOffRampOnOffRampRouter(t, destClient)
	setOffRampOnTokenPools(t, destClient)

	destClient.Logger.Info("Please deploy new execution jobs")
}

/*
func removeOffRamp(t *testing.T, destClient *EvmChainConfig, offRampAddress common.Address) {
	// Pause contract
	revokeOffRampOnOffRampRouter(t, destClient, offRampAddress)
	revokeOffRampOnTokenPools(t, destClient, offRampAddress)
}
*/
