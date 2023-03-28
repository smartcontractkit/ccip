package deployment_io

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/dione"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
)

func PrettyPrintLanes(env dione.Environment, source *rhea.EvmDeploymentConfig, destination *rhea.EvmDeploymentConfig) {
	WriteChainConfigToFile(env, source)
	WriteChainConfigToFile(env, destination)

	writeLaneConfigToFile(env, source)
	writeLaneConfigToFile(env, destination)
}

func WriteChainConfigToFile(env dione.Environment, chain *rhea.EvmDeploymentConfig) {
	sourceChainConfig := prettyPrint(chain.ChainConfig)
	chain.Logger.Info(string(sourceChainConfig))
	chainName := helpers.ChainName(int64(chain.ChainConfig.ChainId))
	attemptWriteToFile(sourceChainConfig, chainName, "chain", string(env))
}

func writeLaneConfigToFile(env dione.Environment, chain *rhea.EvmDeploymentConfig) {
	destLaneConfig := prettyPrint(chain.LaneConfig)
	chain.Logger.Info(string(destLaneConfig))
	chainName := helpers.ChainName(int64(chain.ChainConfig.ChainId))
	attemptWriteToFile(destLaneConfig, chainName, "lane", string(env))
}

func attemptWriteToFile(file []byte, chainName string, configType string, env string) {
	path := fmt.Sprintf("json/deployments/%s/%s/%s", env, configType, chainName)
	// ignore errors
	_ = os.MkdirAll(path, os.ModePerm)
	fileName := fmt.Sprintf("/%s", time.Now().Format("2006-01-02 15:04:05"))
	// ignore errors
	_ = os.WriteFile(path+fileName, file, 0600)
}

func prettyPrint(i interface{}) []byte {
	s, _ := json.MarshalIndent(i, "", "\t")
	return s
}
