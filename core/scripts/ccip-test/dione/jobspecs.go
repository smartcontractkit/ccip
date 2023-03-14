package dione

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/smartcontractkit/chainlink/integration-tests/client"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/testhelpers"
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
		SourceChainName:        helpers.ChainName(int64(sourceClient.ChainConfig.ChainId)),
		DestChainName:          helpers.ChainName(int64(destClient.ChainConfig.ChainId)),
		SourceChainId:          sourceClient.ChainConfig.ChainId,
		DestChainId:            destClient.ChainConfig.ChainId,
		TokenPricesUSDPipeline: GetTokenPricesUSDPipeline(pipelineTokens),
		PollPeriod:             PollPeriod,
		SourceStartBlock:       sourceClient.DeploySettings.DeployedAt,
		DestStartBlock:         destClient.DeploySettings.DeployedAt,
		P2PV2Bootstrappers:     []string{}, // Set in env vars
	}
}

func GetTokenPricesUSDPipeline(pipelineTokens []rhea.EVMBridgedToken) string {
	tokenPricesUSDPipeline := "merge [type=merge left=\"{}\" right=\"{"
	for _, token := range pipelineTokens {
		price := new(big.Int).Mul(token.Price, big.NewInt(1e18))
		tokenPricesUSDPipeline += fmt.Sprintf(`\\\"%s\\\":\\\"%s\\\",`, token.Token.Hex(), price)
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
