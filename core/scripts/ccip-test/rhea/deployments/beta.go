package deployments

import (
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
)

// Chains
var Beta_ChainConfigs = []rhea.EvmDeploymentConfig{
	{ChainConfig: Beta_Sepolia},
	{ChainConfig: Beta_AvaxFuji},
	{ChainConfig: Beta_OptimismGoerli},
}

var Beta_Sepolia = rhea.EVMChainConfig{
	ChainId: 11155111,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:                gethcommon.HexToAddress("0x779877A7B0D9E8603169DdbD7836e478b4624789"),
			Pool:                 gethcommon.HexToAddress(""),
			Price:                big.NewInt(1),
			PriceFeedsAggregator: gethcommon.HexToAddress("0xc59E3633BAAC79493d908e63626716e204A45EdF"),
			TokenPoolType:        rhea.LockRelease,
		},
		//rhea.WETH: {
		//	Token:                gethcommon.HexToAddress("")
		//	Pool:                 gethcommon.HexToAddress(""),
		//	Price:                big.NewInt(1),
		//	PriceFeedsAggregator: gethcommon.HexToAddress("0x694AA1769357215DE4FAC081bf1f309aDC325306"),
		//},
	},
	FeeTokens:     []rhea.Token{rhea.LINK},
	Router:        gethcommon.HexToAddress(""),
	Afn:           gethcommon.HexToAddress(""),
	PriceRegistry: gethcommon.HexToAddress(""),
	Confirmations: 4,
}

var Beta_OptimismGoerli = rhea.EVMChainConfig{
	ChainId: 420,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: true,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:                gethcommon.HexToAddress("0xdc2CC710e42857672E7907CF474a69B63B93089f"),
			Pool:                 gethcommon.HexToAddress(""),
			Price:                big.NewInt(1),
			PriceFeedsAggregator: gethcommon.HexToAddress("0x69C5297001f38cCBE30a81359da06E5256bd28B9"),
			TokenPoolType:        rhea.LockRelease,
		},
		rhea.WETH: {
			Token:                gethcommon.HexToAddress("0x4200000000000000000000000000000000000006"),
			Pool:                 gethcommon.HexToAddress(""),
			Price:                big.NewInt(1),
			PriceFeedsAggregator: gethcommon.HexToAddress("0x69C5297001f38cCBE30a81359da06E5256bd28B9"),
			TokenPoolType:        rhea.LockRelease,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	Router:        gethcommon.HexToAddress(""),
	Afn:           gethcommon.HexToAddress(""),
	PriceRegistry: gethcommon.HexToAddress(""),
	Confirmations: 4,
}

var Beta_AvaxFuji = rhea.EVMChainConfig{
	ChainId: 43113,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:                gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846"),
			Pool:                 gethcommon.HexToAddress(""),
			Price:                big.NewInt(1),
			PriceFeedsAggregator: gethcommon.HexToAddress("0x34C4c526902d88a3Aa98DB8a9b802603EB1E3470"),
			TokenPoolType:        rhea.LockRelease,
		},
		rhea.WAVAX: {
			Token:                gethcommon.HexToAddress("0xd00ae08403B9bbb9124bB305C09058E32C39A48c"),
			Pool:                 gethcommon.HexToAddress(""),
			Price:                big.NewInt(1),
			PriceFeedsAggregator: gethcommon.HexToAddress("0x6C2441920404835155f33d88faf0545B895871b1"),
			TokenPoolType:        rhea.LockRelease,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WAVAX},
	Router:        gethcommon.HexToAddress(""),
	Afn:           gethcommon.HexToAddress(""),
	PriceRegistry: gethcommon.HexToAddress(""),
	Confirmations: 1,
}

var Beta_SepoliaToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_Sepolia,
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

		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployPingPongDapp:   false,
		DeployGovernanceDapp: false,
		DeployedAt:           0,
	},
}

var Beta_AvaxFujiToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_AvaxFuji,
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

		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployPingPongDapp:   false,
		DeployGovernanceDapp: false,
		DeployedAt:           0,
	},
}

var Beta_SepoliaToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_Sepolia,
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

		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployPingPongDapp:   false,
		DeployGovernanceDapp: false,
		DeployedAt:           0,
	},
}

var Beta_OptimismGoerliToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_OptimismGoerli,
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

		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployPingPongDapp:   false,
		DeployGovernanceDapp: false,
		DeployedAt:           0,
	},
}

var Beta_OptimismGoerliToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_OptimismGoerli,
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

		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployPingPongDapp:   false,
		DeployGovernanceDapp: false,
		DeployedAt:           0,
	},
}

var Beta_AvaxFujiToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_AvaxFuji,
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

		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployPingPongDapp:   false,
		DeployGovernanceDapp: false,
		DeployedAt:           0,
	},
}
