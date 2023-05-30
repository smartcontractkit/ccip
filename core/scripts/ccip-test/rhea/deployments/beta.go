package deployments

import (
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

var BetaChains = map[rhea.Chain]rhea.EvmDeploymentConfig{
	rhea.AvaxFuji:       {ChainConfig: Beta_AvaxFuji},
	rhea.OptimismGoerli: {ChainConfig: Beta_OptimismGoerli},
	rhea.Sepolia:        {ChainConfig: Beta_Sepolia},
	rhea.ArbitrumGoerli: {ChainConfig: Beta_ArbitrumGoerli},
}

var BetaChainMapping = map[rhea.Chain]map[rhea.Chain]rhea.EvmDeploymentConfig{
	rhea.AvaxFuji: {
		rhea.OptimismGoerli: Beta_AvaxFujiToOptimismGoerli,
		rhea.ArbitrumGoerli: Beta_AvaxFujiToArbitrumGoerli,
		rhea.Sepolia:        Beta_AvaxFujiToSepolia,
	},
	rhea.OptimismGoerli: {
		rhea.AvaxFuji:       Beta_OptimismGoerliToAvaxFuji,
		rhea.ArbitrumGoerli: Beta_OptimismGoerliToArbitrumGoerli,
		rhea.Sepolia:        Beta_OptimismGoerliToSepolia,
	},
	rhea.ArbitrumGoerli: {
		rhea.AvaxFuji:       Beta_ArbitrumGoerliToAvaxFuji,
		rhea.OptimismGoerli: Beta_ArbitrumGoerliToOptimismGoerli,
		rhea.Sepolia:        Beta_ArbitrumGoerliToSepolia,
	},
	rhea.Sepolia: {
		rhea.AvaxFuji:       Beta_SepoliaToAvaxFuji,
		rhea.OptimismGoerli: Beta_SepoliaToOptimismGoerli,
		rhea.ArbitrumGoerli: Beta_SepoliaToArbitrumGoerli,
	},
}

var Beta_Sepolia = rhea.EVMChainConfig{
	EvmChainId: 11155111,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:         gethcommon.HexToAddress("0x779877A7B0D9E8603169DdbD7836e478b4624789"),
			Pool:          gethcommon.HexToAddress("0xecfa5e914c3d449b4a9b4622dc097ff50e604d52"),
			Price:         rhea.LINK.Price(),
			TokenPoolType: rhea.LockRelease,
			PriceFeed: rhea.PriceFeed{
				Aggregator: gethcommon.HexToAddress("0x5A2734CC0341ea6564dF3D00171cc99C63B1A7d3"),
				Multiplier: big.NewInt(1e10),
			},
		},
		rhea.WETH: {
			Token:         gethcommon.HexToAddress("0x097D90c9d3E0B50Ca60e1ae45F6A81010f9FB534"),
			Pool:          gethcommon.HexToAddress("0xefbad0e7668949704369a7d0e4f9d82bf91093ff"),
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
	Router:        gethcommon.HexToAddress("0x7cbdb9cf6420dd3fda3ce258e9933b998bbe99ac"),
	UpgradeRouter: gethcommon.HexToAddress("0xe436cb4795a31a22222c382a3c64c05029e43295"),
	Afn:           gethcommon.HexToAddress("0x44973d00af062093c1d2d64c6201dbc2de36b4b6"),
	PriceRegistry: gethcommon.HexToAddress("0x4fc8590a8450d98907c861e7031dc7372c2c4df7"),
	TunableChainValues: rhea.TunableChainValues{
		FinalityDepth:            getFinalityDepth(rhea.Sepolia),
		OptimisticConfirmations:  getOptimisticConfirmations(rhea.Sepolia),
		BatchGasLimit:            BATCH_GAS_LIMIT,
		RelativeBoostPerWaitHour: RELATIVE_BOOST_PER_WAIT_HOUR,
		FeeUpdateHeartBeat:       models.MustMakeDuration(FEE_UPDATE_HEARTBEAT),
		FeeUpdateDeviationPPB:    FEE_UPDATE_DEVIATION_PPB,
		MaxGasPrice:              getMaxGasPrice(rhea.Sepolia),
		InflightCacheExpiry:      models.MustMakeDuration(INFLIGHT_CACHE_EXPIRY),
		RootSnoozeTime:           models.MustMakeDuration(ROOT_SNOOZE_TIME),
	},
	DeploySettings: rhea.ChainDeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployUpgradeRouter: false,
		DeployPriceRegistry: false,
	},
}

var Beta_OptimismGoerli = rhea.EVMChainConfig{
	EvmChainId: 420,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:          gethcommon.HexToAddress("0xdc2CC710e42857672E7907CF474a69B63B93089f"),
			Pool:           gethcommon.HexToAddress("0x68bfcdd4a6346a4d3dfc0c88dc554e1cf574b94c"),
			TokenPoolType:  rhea.LockRelease,
			TokenPriceType: rhea.PriceFeeds,
			Price:          rhea.LINK.Price(),
			PriceFeed: rhea.PriceFeed{
				Aggregator: gethcommon.HexToAddress("0x53AFfFfA77006432146b667C67FA77b5D405793b"),
				Multiplier: big.NewInt(1e10),
			},
		},
		rhea.WETH: {
			Token:          gethcommon.HexToAddress("0x4200000000000000000000000000000000000006"),
			Pool:           gethcommon.HexToAddress("0xf1cdcf045c6403339ad5b518b6b240b55847abcb"),
			Price:          rhea.WETH.Price(),
			TokenPoolType:  rhea.LockRelease,
			TokenPriceType: rhea.TokenPrices,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0xa76549590250cf2b13acd848db6ae9e398767121"),
	UpgradeRouter: gethcommon.HexToAddress("0xec3a5b1710f1dd343d710c6be90856fea760252d"),
	Afn:           gethcommon.HexToAddress("0x7faddb7dc5bd30b1fc0df403e788dc6b546bf709"),
	PriceRegistry: gethcommon.HexToAddress("0x20ec331f5bcd9e54f794a163f859cc2e5e4f64d8"),
	TunableChainValues: rhea.TunableChainValues{
		FinalityDepth:            getFinalityDepth(rhea.OptimismGoerli),
		OptimisticConfirmations:  getOptimisticConfirmations(rhea.OptimismGoerli),
		BatchGasLimit:            BATCH_GAS_LIMIT,
		RelativeBoostPerWaitHour: RELATIVE_BOOST_PER_WAIT_HOUR,
		FeeUpdateHeartBeat:       models.MustMakeDuration(FEE_UPDATE_HEARTBEAT),
		FeeUpdateDeviationPPB:    FEE_UPDATE_DEVIATION_PPB,
		MaxGasPrice:              getMaxGasPrice(rhea.OptimismGoerli),
		InflightCacheExpiry:      models.MustMakeDuration(INFLIGHT_CACHE_EXPIRY),
		RootSnoozeTime:           models.MustMakeDuration(ROOT_SNOOZE_TIME),
	},
	DeploySettings: rhea.ChainDeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,
		DeployUpgradeRouter: false,
	},
}

var Beta_AvaxFuji = rhea.EVMChainConfig{
	EvmChainId: 43113,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:          gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846"),
			Pool:           gethcommon.HexToAddress("0x4353997d8dbed095dc49fe0db1de5792a626c360"),
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
			Pool:           gethcommon.HexToAddress("0x410a23e05a28abe33ceba36975e3d99152de6c63"),
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
	Router:        gethcommon.HexToAddress("0x2b2a31372fc7b8508d8505a7e36c2fd32723c09e"),
	UpgradeRouter: gethcommon.HexToAddress("0x589034ff232a32fa7dab47b0e17729c0faed2df7"),
	Afn:           gethcommon.HexToAddress("0xe08f435e7d0c51e1e9d1384bf699a9eae6bea701"),
	PriceRegistry: gethcommon.HexToAddress("0xad5649c1fcdbb6304e1d5f4e81d8d324fb6fabee"),
	TunableChainValues: rhea.TunableChainValues{
		FinalityDepth:            getFinalityDepth(rhea.AvaxFuji),
		OptimisticConfirmations:  getOptimisticConfirmations(rhea.AvaxFuji),
		BatchGasLimit:            BATCH_GAS_LIMIT,
		RelativeBoostPerWaitHour: RELATIVE_BOOST_PER_WAIT_HOUR,
		FeeUpdateHeartBeat:       models.MustMakeDuration(FEE_UPDATE_HEARTBEAT),
		FeeUpdateDeviationPPB:    FEE_UPDATE_DEVIATION_PPB,
		MaxGasPrice:              getMaxGasPrice(rhea.AvaxFuji),
		InflightCacheExpiry:      models.MustMakeDuration(INFLIGHT_CACHE_EXPIRY),
		RootSnoozeTime:           models.MustMakeDuration(ROOT_SNOOZE_TIME),
	},
	DeploySettings: rhea.ChainDeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployUpgradeRouter: false,
		DeployPriceRegistry: false,
	},
}

var Beta_ArbitrumGoerli = rhea.EVMChainConfig{
	EvmChainId: 421613,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: true,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:          gethcommon.HexToAddress("0xd14838A68E8AFBAdE5efb411d5871ea0011AFd28"),
			Pool:           gethcommon.HexToAddress("0xac04bdb191ce949604d8dc851ea436daf1281f5c"),
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
			Pool:           gethcommon.HexToAddress("0x91a60744dcb88d555309980d9bcc8d391c215cdb"),
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
	Router:        gethcommon.HexToAddress("0x32bb20f8c8869afdef1db14c54e5dc89738e914a"),
	UpgradeRouter: gethcommon.HexToAddress("0xf74d62403b6ec0085ce7421320675b88f1ec8a53"),
	Afn:           gethcommon.HexToAddress("0x0ad826d1c3522abe300493139dfb4cd2b863cd64"),
	PriceRegistry: gethcommon.HexToAddress("0x11408549e500b6a0cf8b2ea7d006ab32adb2df4b"),
	TunableChainValues: rhea.TunableChainValues{
		FinalityDepth:            getFinalityDepth(rhea.ArbitrumGoerli),
		OptimisticConfirmations:  getOptimisticConfirmations(rhea.ArbitrumGoerli),
		BatchGasLimit:            BATCH_GAS_LIMIT,
		RelativeBoostPerWaitHour: RELATIVE_BOOST_PER_WAIT_HOUR,
		FeeUpdateHeartBeat:       models.MustMakeDuration(FEE_UPDATE_HEARTBEAT),
		FeeUpdateDeviationPPB:    FEE_UPDATE_DEVIATION_PPB,
		MaxGasPrice:              getMaxGasPrice(rhea.ArbitrumGoerli),
		InflightCacheExpiry:      models.MustMakeDuration(INFLIGHT_CACHE_EXPIRY),
		RootSnoozeTime:           models.MustMakeDuration(ROOT_SNOOZE_TIME),
	},
	DeploySettings: rhea.ChainDeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployUpgradeRouter: false,
		DeployPriceRegistry: false,
	},
}

var Beta_OptimismGoerliToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x3ab3c715e79ef433fc615b0fea8fbc4151c4dfe1"),
		OffRamp:      gethcommon.HexToAddress("0xb9c06e0029a3ee3d1e770fd9218993c6880493a2"),
		CommitStore:  gethcommon.HexToAddress("0x6a519df34aa2cb3dfb57284ac840432b53182c3a"),
		PingPongDapp: gethcommon.HexToAddress("0x7880293c10f86a5ae9bc271cc3c031396ca499f2"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    9970189,
		},
	},
}

var Beta_AvaxFujiToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x246885339751ec09a07587833c7d1aefec80b372"),
		OffRamp:      gethcommon.HexToAddress("0xfef71d030b530502737bd64dfea0a167580b97a5"),
		CommitStore:  gethcommon.HexToAddress("0xc428d000c9e10aea137d89593cbde85055e3e584"),
		PingPongDapp: gethcommon.HexToAddress("0x28d49961a83b679c7f50a96e19533ebfe14d7c13"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    22426883,
		},
	},
}

var Beta_ArbitrumGoerliToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_ArbitrumGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xa5f819e24c4abfc4200987a5fabf562a904fed92"),
		OffRamp:      gethcommon.HexToAddress("0x5aba4c199cf737866a718eda7810878db75ecf66"),
		CommitStore:  gethcommon.HexToAddress("0x77455df22d16277622dbde3bf8b2aa26418f10ed"),
		PingPongDapp: gethcommon.HexToAddress("0xcd4bc3d174ec4dafad35df7765aa03bf7de78388"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    21801909,
		},
	},
}

var Beta_AvaxFujiToArbitrumGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x02482b253e7727f44d4e42dcbfea09ec0211a95b"),
		OffRamp:      gethcommon.HexToAddress("0x9734ca95c70703368cabd67cb108a57db5f8fe0f"),
		CommitStore:  gethcommon.HexToAddress("0xb7131b89e3380f7cb6aaafe92f3c2fddc9b392c6"),
		PingPongDapp: gethcommon.HexToAddress("0x0e8931475cd45ab2c796fdad5d530c34fb3edecc"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    22251404,
		},
	},
}

var Beta_ArbitrumGoerliToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_ArbitrumGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xf0046682ec5a3427cb37fe3e3d2dafea8eb409b4"),
		OffRamp:      gethcommon.HexToAddress("0x3a9eb812a299196f2b4e05c62cbeb959c7beab59"),
		CommitStore:  gethcommon.HexToAddress("0xf987ad0f725b95c7d565ef28a4b5aa441187a14d"),
		PingPongDapp: gethcommon.HexToAddress("0xc07699e59ef98ae0aa09a9d2cbeb80192b992c5c"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    22063933,
		},
	},
}

var Beta_OptimismGoerliToArbitrumGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xefbad0e7668949704369a7d0e4f9d82bf91093ff"),
		OffRamp:      gethcommon.HexToAddress("0xc56e536229f73cb711bd546c24ad255665c04047"),
		CommitStore:  gethcommon.HexToAddress("0xf248de56e32d8f6fe8d8fba42d39020d8bfc89fc"),
		PingPongDapp: gethcommon.HexToAddress("0x65647829a2d9f4726eb826ee3b5ce1764eb59a93"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    9819412,
		},
	},
}

var Beta_AvaxFujiToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x02b963404d58dd7fc725e5266dc837b8df179424"),
		OffRamp:      gethcommon.HexToAddress("0x6f080341b74c7a6bc10af38bccb735728d2156fe"),
		CommitStore:  gethcommon.HexToAddress("0xcebe2538df3ab9b80a6a461d886922618bea763b"),
		PingPongDapp: gethcommon.HexToAddress("0x7e0eb057bbe5681666d1aac8d1c2ec909d9dc412"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    22427198,
		},
	},
}

var Beta_SepoliaToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x2b2a31372fc7b8508d8505a7e36c2fd32723c09e"),
		OffRamp:      gethcommon.HexToAddress("0x0a504a58f3711662ae07a053451524b73b4f6c99"),
		CommitStore:  gethcommon.HexToAddress("0xabe9631a964a76eafe278a508f462741dcf708c3"),
		PingPongDapp: gethcommon.HexToAddress("0x4752c8696cec308abd92abc3465de5e384ec0738"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3583064,
		},
	},
}

var Beta_OptimismGoerliToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xa2a7e89a684784dd73e80bdcf2f2368ef7c7a690"),
		OffRamp:      gethcommon.HexToAddress("0xaa0090ab7944648c7d3a950f714b6acd3338774b"),
		CommitStore:  gethcommon.HexToAddress("0xfd7676202c2e64bb2d030e7f8bb519fad378780f"),
		PingPongDapp: gethcommon.HexToAddress("0x701ccf1c98bd3284649dc6cf08c69410d90c8590"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    9768160,
		},
	},
}

var Beta_SepoliaToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xf248de56e32d8f6fe8d8fba42d39020d8bfc89fc"),
		OffRamp:      gethcommon.HexToAddress("0xf117c308b9d5f2df16871255c04697309c8b7b0f"),
		CommitStore:  gethcommon.HexToAddress("0x3eae3950ff71db1c6eeee5a777be03c36a517fee"),
		PingPongDapp: gethcommon.HexToAddress("0x244209b8cd80ac7c9e43f8fdcbcd8435be1e2729"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3552407,
		},
	},
}

var Beta_ArbitrumGoerliToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_ArbitrumGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xe54562659c62d5fa3b93b78192985ac9e7b04b05"),
		OffRamp:      gethcommon.HexToAddress("0x94e5fceaade02b20da2e3ede15740f1137a8a023"),
		CommitStore:  gethcommon.HexToAddress("0x07828f55f1eff4e972890a183e3ab787316ed31e"),
		PingPongDapp: gethcommon.HexToAddress("0x8bd92cf0ec06a034f76996c190f59495191a224e"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    22949258,
		},
	},
	UpgradeLaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x0"),
		OffRamp:      gethcommon.HexToAddress("0x0"),
		CommitStore:  gethcommon.HexToAddress("0x0"),
		PingPongDapp: gethcommon.HexToAddress("0x7fd2e811b6507fe4b255d37f709c250804aaa810"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    22949258,
		},
	},
}

var Beta_SepoliaToArbitrumGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xce41acebd9d929e144057ddee824a857164f508e"),
		OffRamp:      gethcommon.HexToAddress("0x593f8596cd42e64ead05a04dc297eb873b55b6e3"),
		CommitStore:  gethcommon.HexToAddress("0xbd7f4bb9dbc4e8f11b49e09f570cefe3440b7356"),
		PingPongDapp: gethcommon.HexToAddress("0x3dd6806335cc17757a75259f2a33f80b931e72b9"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3587929,
		},
	},
	UpgradeLaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x0"),
		OffRamp:      gethcommon.HexToAddress("0x0"),
		CommitStore:  gethcommon.HexToAddress("0x0"),
		PingPongDapp: gethcommon.HexToAddress("0x95e69e03a151478ec20fe349be5332274bf7cb00"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3587929,
		},
	},
}
