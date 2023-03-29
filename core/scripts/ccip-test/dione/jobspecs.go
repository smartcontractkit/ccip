package dione

import (
	"fmt"
	"strings"

	"github.com/smartcontractkit/chainlink/integration-tests/client"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/testhelpers"
)

// NewCCIPJobSpecParams returns set of parameters needed for setting up ccip jobs for sourceClient --> destClient
func NewCCIPJobSpecParams(sourceClient rhea.EvmDeploymentConfig, destClient rhea.EvmDeploymentConfig) testhelpers.CCIPJobSpecParams {
	var pipelineTokens []rhea.EVMBridgedToken
	for _, feeTokenName := range destClient.ChainConfig.FeeTokens {
		if token, ok := destClient.ChainConfig.SupportedTokens[feeTokenName]; ok {
			pipelineTokens = append(pipelineTokens, token)
		}
	}
	pipelineTokens = append(pipelineTokens, sourceClient.ChainConfig.SupportedTokens[sourceClient.ChainConfig.WrappedNative])
	return testhelpers.CCIPJobSpecParams{
		OffRamp:                destClient.LaneConfig.OffRamp,
		OnRamp:                 sourceClient.LaneConfig.OnRamp,
		CommitStore:            destClient.LaneConfig.CommitStore,
		SourceChainName:        ccip.ChainName(int64(sourceClient.ChainConfig.ChainId)),
		DestChainName:          ccip.ChainName(int64(destClient.ChainConfig.ChainId)),
		SourceChainId:          sourceClient.ChainConfig.ChainId,
		DestChainId:            destClient.ChainConfig.ChainId,
		TokenPricesUSDPipeline: GetTokenPricesUSDPipeline(pipelineTokens),
		PollPeriod:             PollPeriod,
		SourceStartBlock:       sourceClient.LaneConfig.DeploySettings.DeployedAtBlock,
		DestStartBlock:         destClient.LaneConfig.DeploySettings.DeployedAtBlock,
		P2PV2Bootstrappers:     []string{}, // Set in env vars
	}
}

func GetTokenPricesUSDPipeline(pipelineTokens []rhea.EVMBridgedToken) string {
	tokenPricesUSDPipeline := "merge [type=merge left=\"{}\" right=\"{"
	for _, token := range pipelineTokens {
		tokenPricesUSDPipeline += fmt.Sprintf(`\\\"%s\\\":\\\"%s\\\",`, token.Token.Hex(), token.Price)
	}
	tokenPricesUSDPipeline = strings.TrimSuffix(tokenPricesUSDPipeline, ",")
	tokenPricesUSDPipeline += "}\"];"
	return tokenPricesUSDPipeline
}

func GetOCRkeysForChainType(OCRKeys client.OCR2Keys, chainType string) client.OCR2KeyData {
	for _, key := range OCRKeys.Data {
		if key.Attributes.ChainType == chainType {
			return key
		}
	}

	panic("Keys not found for chain")
}
