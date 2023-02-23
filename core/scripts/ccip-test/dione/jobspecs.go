package dione

import (
	"fmt"
	"sort"
	"strings"

	"github.com/smartcontractkit/chainlink/integration-tests/client"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/testhelpers"
)

// NewCCIPJobSpecParams returns set of parameters needed for setting up ccip jobs for sourceClient --> destClient
func NewCCIPJobSpecParams(sourceClient rhea.EvmDeploymentConfig, destClient rhea.EvmDeploymentConfig) testhelpers.CCIPJobSpecParams {
	return testhelpers.CCIPJobSpecParams{
		OffRamp:                  destClient.LaneConfig.OffRamp,
		OnRamp:                   sourceClient.LaneConfig.OnRamp,
		CommitStore:              destClient.LaneConfig.CommitStore,
		SourceChainName:          helpers.ChainName(int64(sourceClient.ChainConfig.ChainId)),
		DestChainName:            helpers.ChainName(int64(destClient.ChainConfig.ChainId)),
		SourceChainId:            sourceClient.ChainConfig.ChainId,
		DestChainId:              destClient.ChainConfig.ChainId,
		TokensPerFeeCoinPipeline: GetTokensPerFeeCoinPipeline(destClient.ChainConfig.SupportedTokens),
		PollPeriod:               PollPeriod,
		SourceStartBlock:         sourceClient.DeploySettings.DeployedAt,
		DestStartBlock:           destClient.DeploySettings.DeployedAt,
		P2PV2Bootstrappers:       []string{}, // Set in env vars
	}
}

func GetTokensPerFeeCoinPipeline(supportedTokens map[rhea.Token]rhea.EVMBridgedToken) string {
	var sortedTokensNames []string
	for tokenName := range supportedTokens {
		sortedTokensNames = append(sortedTokensNames, string(tokenName))
	}
	sort.Strings(sortedTokensNames)

	tokensPerFeeCoinPipeline := "merge [type=merge left=\"{}\" right=\"{"
	for _, tokenName := range sortedTokensNames {
		token := supportedTokens[rhea.Token(tokenName)]
		tokensPerFeeCoinPipeline += fmt.Sprintf(`\\\"%s\\\":\\\"1000000000000000000\\\",`, token.Token.Hex())
	}
	tokensPerFeeCoinPipeline = strings.TrimSuffix(tokensPerFeeCoinPipeline, ",")
	tokensPerFeeCoinPipeline += "}\"];"
	return tokensPerFeeCoinPipeline
}

func GetOCRkeysForChainType(OCRKeys client.OCR2Keys, chainType string) client.OCR2KeyData {
	for _, key := range OCRKeys.Data {
		if key.Attributes.ChainType == chainType {
			return key
		}
	}

	panic("Keys not found for chain")
}
