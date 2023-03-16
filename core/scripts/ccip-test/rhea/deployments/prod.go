package deployments

import (
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
)

var ProdChains = map[rhea.Chain]rhea.EvmDeploymentConfig{
	rhea.AvaxFuji:       {ChainConfig: Prod_AvaxFuji},
	rhea.OptimismGoerli: {ChainConfig: Prod_OptimismGoerli},
	rhea.Sepolia:        {ChainConfig: Prod_Sepolia},
}

var ProdChainMapping = map[rhea.Chain]map[rhea.Chain]rhea.EvmDeploymentConfig{
	rhea.Sepolia: {
		rhea.AvaxFuji:       Prod_SepoliaToAvaxFuji,
		rhea.OptimismGoerli: Prod_SepoliaToOptimismGoerli,
	},
	rhea.AvaxFuji: {
		rhea.Sepolia:        Prod_AvaxFujiToSepolia,
		rhea.OptimismGoerli: Prod_AvaxFujiToOptimismGoerli,
	},
	rhea.OptimismGoerli: {
		rhea.Sepolia:  Prod_OptimismGoerliToSepolia,
		rhea.AvaxFuji: Prod_OptimismGoerliToAvaxFuji,
	},
}

var Prod_Sepolia = rhea.EVMChainConfig{
	ChainId: 11155111,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:                gethcommon.HexToAddress(""),
			Pool:                 gethcommon.HexToAddress(""),
			Price:                big.NewInt(10),
			PriceFeedsAggregator: gethcommon.HexToAddress(""),
			TokenPoolType:        rhea.LockRelease,
		},
		rhea.WETH: {
			Token:         gethcommon.HexToAddress(""),
			Pool:          gethcommon.HexToAddress(""),
			Price:         big.NewInt(1500),
			TokenPoolType: rhea.LockRelease,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress(""),
	Afn:           gethcommon.HexToAddress(""),
	PriceRegistry: gethcommon.HexToAddress(""),
	Confirmations: 4,
}

var Prod_OptimismGoerli = rhea.EVMChainConfig{
	ChainId: 420,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: true,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:         gethcommon.HexToAddress(""),
			Pool:          gethcommon.HexToAddress(""),
			Price:         big.NewInt(10),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.WETH: {
			Token:         gethcommon.HexToAddress(""),
			Pool:          gethcommon.HexToAddress(""),
			Price:         big.NewInt(1500),
			TokenPoolType: rhea.LockRelease,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress(""),
	Afn:           gethcommon.HexToAddress(""),
	PriceRegistry: gethcommon.HexToAddress(""),
	Confirmations: 4,
}

var Prod_AvaxFuji = rhea.EVMChainConfig{
	ChainId: 43113,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:                gethcommon.HexToAddress(""),
			Pool:                 gethcommon.HexToAddress(""),
			Price:                big.NewInt(10),
			PriceFeedsAggregator: gethcommon.HexToAddress(""),
			TokenPoolType:        rhea.LockRelease,
		},
		rhea.WAVAX: {
			Token:                gethcommon.HexToAddress(""),
			Pool:                 gethcommon.HexToAddress(""),
			Price:                big.NewInt(25),
			PriceFeedsAggregator: gethcommon.HexToAddress(""),
			TokenPoolType:        rhea.LockRelease,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WAVAX},
	WrappedNative: rhea.WAVAX,
	Router:        gethcommon.HexToAddress(""),
	Afn:           gethcommon.HexToAddress(""),
	PriceRegistry: gethcommon.HexToAddress(""),
	Confirmations: 1,
}

// Lanes
var Prod_SepoliaToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress(""),
		OnRamp:       gethcommon.HexToAddress(""),
		OffRamp:      gethcommon.HexToAddress(""),
		ReceiverDapp: gethcommon.HexToAddress(""),
		PingPongDapp: gethcommon.HexToAddress(""),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,

		DeployCommitStore:  false,
		DeployRamp:         false,
		DeployPingPongDapp: false,
		DeployedAt:         2808574,
	},
}

var Prod_AvaxFujiToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress(""),
		OnRamp:       gethcommon.HexToAddress(""),
		OffRamp:      gethcommon.HexToAddress(""),
		ReceiverDapp: gethcommon.HexToAddress(""),
		PingPongDapp: gethcommon.HexToAddress(""),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,

		DeployCommitStore:  false,
		DeployRamp:         false,
		DeployPingPongDapp: false,
		DeployedAt:         18523284,
	},
}

var Prod_SepoliaToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress(""),
		OnRamp:       gethcommon.HexToAddress(""),
		OffRamp:      gethcommon.HexToAddress(""),
		ReceiverDapp: gethcommon.HexToAddress(""),
		PingPongDapp: gethcommon.HexToAddress(""),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,

		DeployCommitStore:  false,
		DeployRamp:         false,
		DeployPingPongDapp: false,
		DeployedAt:         2808895,
	},
}

var Prod_OptimismGoerliToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress(""),
		OnRamp:       gethcommon.HexToAddress(""),
		OffRamp:      gethcommon.HexToAddress(""),
		ReceiverDapp: gethcommon.HexToAddress(""),
		PingPongDapp: gethcommon.HexToAddress(""),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,

		DeployCommitStore:  false,
		DeployRamp:         false,
		DeployPingPongDapp: false,
		DeployedAt:         4874057,
	},
}

var Prod_OptimismGoerliToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress(""),
		OnRamp:       gethcommon.HexToAddress(""),
		OffRamp:      gethcommon.HexToAddress(""),
		ReceiverDapp: gethcommon.HexToAddress(""),
		PingPongDapp: gethcommon.HexToAddress(""),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,

		DeployCommitStore:  false,
		DeployRamp:         false,
		DeployPingPongDapp: true,
		DeployedAt:         4874027,
	},
}

var Prod_AvaxFujiToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress(""),
		OnRamp:       gethcommon.HexToAddress(""),
		OffRamp:      gethcommon.HexToAddress(""),
		ReceiverDapp: gethcommon.HexToAddress(""),
		PingPongDapp: gethcommon.HexToAddress(""),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,

		DeployCommitStore:  false,
		DeployRamp:         false,
		DeployPingPongDapp: true,
		DeployedAt:         18504934,
	},
}
