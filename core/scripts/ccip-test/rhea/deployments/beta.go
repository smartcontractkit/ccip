package deployments

import (
	"math/big"
	"time"

	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

var BetaChains = map[rhea.Chain]rhea.EvmDeploymentConfig{
	rhea.AvaxFuji:       {ChainConfig: Beta_AvaxFuji},
	rhea.OptimismGoerli: {ChainConfig: Beta_OptimismGoerli},
	rhea.ArbitrumGoerli: {ChainConfig: Beta_ArbitrumGoerli},
	rhea.Goerli:         {ChainConfig: Beta_Goerli},
	rhea.Sepolia:        {ChainConfig: Beta_Sepolia},
}

var BetaChainMapping = map[rhea.Chain]map[rhea.Chain]rhea.EvmDeploymentConfig{
	rhea.AvaxFuji: {
		rhea.OptimismGoerli: Beta_AvaxFujiToOptimismGoerli,
		rhea.ArbitrumGoerli: Beta_AvaxFujiToArbitrumGoerli,
	},
	rhea.OptimismGoerli: {
		rhea.AvaxFuji:       Beta_OptimismGoerliToAvaxFuji,
		rhea.ArbitrumGoerli: Beta_OptimismGoerliToArbitrumGoerli,
	},
	rhea.ArbitrumGoerli: {
		rhea.AvaxFuji:       Beta_ArbitrumGoerliToAvaxFuji,
		rhea.OptimismGoerli: Beta_ArbitrumGoerliToOptimismGoerli,
	},
}

var Beta_Goerli = rhea.EVMChainConfig{
	ChainId: 5,
	GasSettings: rhea.EVMGasSettings{
		EIP1559:   true,
		GasTipCap: rhea.DefaultGasTipFee,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:         gethcommon.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB"),
			Pool:          gethcommon.HexToAddress(""),
			Price:         rhea.LINK.Price(),
			TokenPoolType: rhea.LockRelease,
			PriceFeed: rhea.PriceFeed{
				Aggregator: gethcommon.HexToAddress("0x15608A8A3E2f65e00fe0ef2C9c78Ada4e4E8172E"),
				Multiplier: big.NewInt(1e10),
			},
		},
		rhea.WETH: {
			Token: gethcommon.HexToAddress("0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6"),
			Pool:  gethcommon.HexToAddress(""),
			Price: rhea.WETH.Price(),
			PriceFeed: rhea.PriceFeed{
				Aggregator: gethcommon.HexToAddress("0xeEF0ee34696a7982b8AE27c67f91f364D91967a0"),
				Multiplier: big.NewInt(1e10),
			},

			TokenPoolType: rhea.LockRelease,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress(""),
	Afn:           gethcommon.HexToAddress(""),
	PriceRegistry: gethcommon.HexToAddress(""),
	DeploySettings: rhea.ChainDeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,
		DeployedAtBlock:     0,
	},
}

var Beta_Sepolia = rhea.EVMChainConfig{
	ChainId: 11155111,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},

	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:         gethcommon.HexToAddress("0x779877A7B0D9E8603169DdbD7836e478b4624789"),
			Pool:          gethcommon.HexToAddress(""),
			Price:         rhea.LINK.Price(),
			TokenPoolType: rhea.LockRelease,
			PriceFeed: rhea.PriceFeed{
				Aggregator: gethcommon.HexToAddress("0x5A2734CC0341ea6564dF3D00171cc99C63B1A7d3"),
				Multiplier: big.NewInt(1e10),
			},
		},
		rhea.WETH: {
			Token:         gethcommon.HexToAddress("0x097D90c9d3E0B50Ca60e1ae45F6A81010f9FB534"),
			Pool:          gethcommon.HexToAddress(""),
			Price:         rhea.WETH.Price(),
			TokenPoolType: rhea.LockRelease,
			PriceFeed: rhea.PriceFeed{
				Aggregator: gethcommon.HexToAddress("0x719E22E3D4b690E5d96cCb40619180B5427F14AE"),
				Multiplier: big.NewInt(1e10),
			},
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress(""),
	Afn:           gethcommon.HexToAddress(""),
	PriceRegistry: gethcommon.HexToAddress(""),
	Confirmations: 4,
	DeploySettings: rhea.ChainDeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,
		DeployedAtBlock:     0,
	},
}

var Beta_OptimismGoerli = rhea.EVMChainConfig{
	ChainId: 420,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: true,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:          gethcommon.HexToAddress("0xdc2CC710e42857672E7907CF474a69B63B93089f"),
			TokenPoolType:  rhea.LockRelease,
			Pool:           gethcommon.HexToAddress("0xd798a5f80ec531d7feafad5e008801a96aa1cdf0"),
			TokenPriceType: rhea.PriceFeeds,
			Price:          rhea.LINK.Price(),
			PriceFeed: rhea.PriceFeed{
				Aggregator: gethcommon.HexToAddress("0x53AFfFfA77006432146b667C67FA77b5D405793b"),
				Multiplier: big.NewInt(1e10),
			},
		},
		rhea.WETH: {
			Token:          gethcommon.HexToAddress("0x4200000000000000000000000000000000000006"),
			Pool:           gethcommon.HexToAddress("0xd53ad873bd0aa553dd37d7608101c30138b3e207"),
			Price:          rhea.WETH.Price(),
			TokenPoolType:  rhea.LockRelease,
			TokenPriceType: rhea.TokenPrices,
			//PriceFeedsAggregator: gethcommon.HexToAddress("0x95Fd25C1238ED3274A53250927B568aF3D80E654"),
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0x6378637c6da4bb73356d0e1b7858d7e238597642"),
	Afn:           gethcommon.HexToAddress("0xb1c6dfae50984367711b81f2d9518ca4c133c2cd"),
	PriceRegistry: gethcommon.HexToAddress("0x8e53ae9420247fe49c01bfd0e8fd362c978499f3"),
	Confirmations: 4,
	DeploySettings: rhea.ChainDeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,
		DeployedAtBlock:     0,
	},
}

var Beta_AvaxFuji = rhea.EVMChainConfig{
	ChainId: 43113,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:          gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846"),
			Pool:           gethcommon.HexToAddress("0xdbd821d220265b8499bfd3b1950a71077e62e537"),
			Price:          rhea.LINK.Price(),
			TokenPoolType:  rhea.LockRelease,
			TokenPriceType: rhea.TokenPrices,
			PriceFeed: rhea.PriceFeed{
				Aggregator: gethcommon.HexToAddress("0x5F4a4f309Aefb6fb0Ab927A0421D0342fF92f194"),
				Multiplier: big.NewInt(1e10),
			},
		},
		rhea.WAVAX: {
			Token:          gethcommon.HexToAddress("0xd00ae08403B9bbb9124bB305C09058E32C39A48c"),
			Pool:           gethcommon.HexToAddress("0x1038d32bbda703b14398f47290751b0c0e3e4190"),
			Price:          rhea.WAVAX.Price(),
			TokenPoolType:  rhea.LockRelease,
			TokenPriceType: rhea.TokenPrices,
			PriceFeed: rhea.PriceFeed{
				Aggregator: gethcommon.HexToAddress("0x6C2441920404835155f33d88faf0545B895871b1"),
				Multiplier: big.NewInt(1e10),
			},
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WAVAX},
	WrappedNative: rhea.WAVAX,
	Router:        gethcommon.HexToAddress("0xe86da2e1fd8d93b7a53a25fbfc55141f13f44fc0"),
	Afn:           gethcommon.HexToAddress("0x5a730b4c573cc8f9bfa8e4c87104a0cadca32206"),
	PriceRegistry: gethcommon.HexToAddress("0x081bb0ab942542cb7b652ebab874e8e153f35443"),
	Confirmations: 1,
	DeploySettings: rhea.ChainDeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,
		DeployedAtBlock:     0,
	},
}

var Beta_ArbitrumGoerli = rhea.EVMChainConfig{
	ChainId: 421613,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: true,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:          gethcommon.HexToAddress("0xd14838A68E8AFBAdE5efb411d5871ea0011AFd28"),
			Price:          rhea.LINK.Price(),
			TokenPoolType:  rhea.LockRelease,
			TokenPriceType: rhea.TokenPrices,
			PriceFeed: rhea.PriceFeed{
				Aggregator: gethcommon.HexToAddress("0xb1D4538B4571d411F07960EF2838Ce337FE1E80E"),
				Multiplier: big.NewInt(1e10),
			},
		},
		rhea.WETH: {
			Token:          gethcommon.HexToAddress("0x32d5D5978905d9c6c2D4C417F0E06Fe768a4FB5a"),
			Price:          rhea.WETH.Price(),
			TokenPoolType:  rhea.LockRelease,
			TokenPriceType: rhea.TokenPrices,
			PriceFeed: rhea.PriceFeed{
				Aggregator: gethcommon.HexToAddress("0xC975dEfb12C5e83F2C7E347831126cF136196447"),
				Multiplier: big.NewInt(1e10),
			},
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Confirmations: 1,
	DeploySettings: rhea.ChainDeploySettings{
		DeployAFN:           true,
		DeployTokenPools:    true,
		DeployRouter:        true,
		DeployPriceRegistry: true,
		DeployedAtBlock:     0,
	},
}

var Beta_OptimismGoerliToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x95907b18815e49074298ed34b17789f7c0a55bab"),
		OnRamp:       gethcommon.HexToAddress("0x9988026ad1f8dcfad2774a30dc9d12877e507ca6"),
		OffRamp:      gethcommon.HexToAddress("0x7ba06058c210005ff07caddc59eab2274e04c7b1"),
		PingPongDapp: gethcommon.HexToAddress("0x1e259733242246c4a2ea14a4408e1c53083f901d"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    7367304,
		},
		CommitOffchainConfig: ccip.CommitOffchainConfig{
			FeeUpdateHeartBeat:    models.MustMakeDuration(24 * time.Hour),
			FeeUpdateDeviationPPB: 5e7,
			MaxGasPrice:           200e9,
		},
		ExecOffchainConfig: ccip.ExecOffchainConfig{
			BatchGasLimit:            5_000_000,
			RelativeBoostPerWaitHour: 0.7,
			MaxGasPrice:              200e9,
		},
	},
}

var Beta_AvaxFujiToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0xd3d3309e6be8abc3e0f81d5744c39c2c399e562f"),
		OnRamp:       gethcommon.HexToAddress("0xf147cf6b7ce39b7e9aa2c5b2ab8c79edd34d6a05"),
		OffRamp:      gethcommon.HexToAddress("0x179115cb633cdad9ea3a5295159662a1339807d7"),
		PingPongDapp: gethcommon.HexToAddress("0x99015fb7d2c75125ce75871f92df95281c915e70"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    20412087,
		},
		CommitOffchainConfig: ccip.CommitOffchainConfig{
			FeeUpdateHeartBeat:    models.MustMakeDuration(24 * time.Hour),
			FeeUpdateDeviationPPB: 5e7,
			MaxGasPrice:           200e9,
		},
		ExecOffchainConfig: ccip.ExecOffchainConfig{
			BatchGasLimit:            5_000_000,
			RelativeBoostPerWaitHour: 0.7,
			MaxGasPrice:              200e9,
		},
	},
}

var Beta_ArbitrumGoerliToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_ArbitrumGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  true,
			DeployRamp:         true,
			DeployPingPongDapp: true,
			DeployedAtBlock:    0,
		},
		CommitOffchainConfig: ccip.CommitOffchainConfig{
			FeeUpdateHeartBeat:    models.MustMakeDuration(24 * time.Hour),
			FeeUpdateDeviationPPB: 5e7,
			MaxGasPrice:           200e9,
		},
		ExecOffchainConfig: ccip.ExecOffchainConfig{
			BatchGasLimit:            5_000_000,
			RelativeBoostPerWaitHour: 0.7,
			MaxGasPrice:              200e9,
		},
	},
}

var Beta_AvaxFujiToArbitrumGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    20412087,
		},
		CommitOffchainConfig: ccip.CommitOffchainConfig{
			FeeUpdateHeartBeat:    models.MustMakeDuration(24 * time.Hour),
			FeeUpdateDeviationPPB: 5e7,
			MaxGasPrice:           200e9,
		},
		ExecOffchainConfig: ccip.ExecOffchainConfig{
			BatchGasLimit:            5_000_000,
			RelativeBoostPerWaitHour: 0.7,
			MaxGasPrice:              200e9,
		},
	},
}

var Beta_ArbitrumGoerliToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_ArbitrumGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  true,
			DeployRamp:         true,
			DeployPingPongDapp: true,
			DeployedAtBlock:    0,
		},
		CommitOffchainConfig: ccip.CommitOffchainConfig{
			FeeUpdateHeartBeat:    models.MustMakeDuration(24 * time.Hour),
			FeeUpdateDeviationPPB: 5e7,
			MaxGasPrice:           200e9,
		},
		ExecOffchainConfig: ccip.ExecOffchainConfig{
			BatchGasLimit:            5_000_000,
			RelativeBoostPerWaitHour: 0.7,
			MaxGasPrice:              200e9,
		},
	},
}

var Beta_OptimismGoerliToArbitrumGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  true,
			DeployRamp:         true,
			DeployPingPongDapp: true,
			DeployedAtBlock:    0,
		},
		CommitOffchainConfig: ccip.CommitOffchainConfig{
			FeeUpdateHeartBeat:    models.MustMakeDuration(24 * time.Hour),
			FeeUpdateDeviationPPB: 5e7,
			MaxGasPrice:           200e9,
		},
		ExecOffchainConfig: ccip.ExecOffchainConfig{
			BatchGasLimit:            5_000_000,
			RelativeBoostPerWaitHour: 0.7,
			MaxGasPrice:              200e9,
		},
	},
}
