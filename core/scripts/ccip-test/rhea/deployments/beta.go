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
	Afn:           gethcommon.HexToAddress("0x1df877d5cefb303a94b4fbe3f3c3089fb8332db3"),
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
	Afn:           gethcommon.HexToAddress("0xb15fa4de18ae593c8d53520dc7e2679c9d0cd2be"),
	PriceRegistry: gethcommon.HexToAddress("0x20ec331f5bcd9e54f794a163f859cc2e5e4f64d8"),
	TunableChainValues: rhea.TunableChainValues{
		FinalityDepth:            getFinalityDepth(rhea.OptimismGoerli),
		OptimisticConfirmations:  getOptimisticConfirmations(rhea.OptimismGoerli),
		BatchGasLimit:            BATCH_GAS_LIMIT,
		RelativeBoostPerWaitHour: RELATIVE_BOOST_PER_WAIT_HOUR,
		FeeUpdateHeartBeat:       models.MustMakeDuration(FEE_UPDATE_HEARTBEAT),
		FeeUpdateDeviationPPB:    FEE_UPDATE_DEVIATION_PPB_FAST_CHAIN,
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
	Afn:           gethcommon.HexToAddress("0x829135c138d7ac262a0504b17cf4176896e491e6"),
	PriceRegistry: gethcommon.HexToAddress("0xad5649c1fcdbb6304e1d5f4e81d8d324fb6fabee"),
	TunableChainValues: rhea.TunableChainValues{
		FinalityDepth:            getFinalityDepth(rhea.AvaxFuji),
		OptimisticConfirmations:  getOptimisticConfirmations(rhea.AvaxFuji),
		BatchGasLimit:            BATCH_GAS_LIMIT,
		RelativeBoostPerWaitHour: RELATIVE_BOOST_PER_WAIT_HOUR,
		FeeUpdateHeartBeat:       models.MustMakeDuration(FEE_UPDATE_HEARTBEAT),
		FeeUpdateDeviationPPB:    FEE_UPDATE_DEVIATION_PPB_FAST_CHAIN,
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
	Afn:           gethcommon.HexToAddress("0xca3186cf799f07c68694737ec45026bee3b4d9c2"),
	PriceRegistry: gethcommon.HexToAddress("0x11408549e500b6a0cf8b2ea7d006ab32adb2df4b"),
	TunableChainValues: rhea.TunableChainValues{
		FinalityDepth:            getFinalityDepth(rhea.ArbitrumGoerli),
		OptimisticConfirmations:  getOptimisticConfirmations(rhea.ArbitrumGoerli),
		BatchGasLimit:            BATCH_GAS_LIMIT,
		RelativeBoostPerWaitHour: RELATIVE_BOOST_PER_WAIT_HOUR,
		FeeUpdateHeartBeat:       models.MustMakeDuration(FEE_UPDATE_HEARTBEAT),
		FeeUpdateDeviationPPB:    FEE_UPDATE_DEVIATION_PPB_FAST_CHAIN,
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
		OnRamp:       gethcommon.HexToAddress("0x0a324534145c12d1d7b0bedcb848d22cda68d3be"),
		OffRamp:      gethcommon.HexToAddress("0x55d80cb2c68b24dbcfd96e4881f769d1af0b31a2"),
		CommitStore:  gethcommon.HexToAddress("0xadf3e27e415b0ec24c4ee926aae4dfdb8ac80b24"),
		PingPongDapp: gethcommon.HexToAddress("0x981fba55ee9ff40b0d3101a95a0a67d14b274afa"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    10273267,
		},
	},
}

var Beta_AvaxFujiToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xa89492a4cf20a83c43bb58c6269f300b94031316"),
		OffRamp:      gethcommon.HexToAddress("0x8b802b71605a15f98257f60fe5b53e756861bb93"),
		CommitStore:  gethcommon.HexToAddress("0xb7ea88547075dcbfb96de9b8d0d9e79bb1557cc4"),
		PingPongDapp: gethcommon.HexToAddress("0x5db83d90e687e7700a8efb94162ddbf9a23bcedb"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    22715575,
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
		OnRamp:       gethcommon.HexToAddress("0xbfb2fef68bb872eac19de9e2061e9f99e7803baa"),
		OffRamp:      gethcommon.HexToAddress("0xf5aadb5aefe9e63b5fa273f31ef86f3dab039e37"),
		CommitStore:  gethcommon.HexToAddress("0xf9a449aeff10ffbeb4e6410bb531831b98258861"),
		PingPongDapp: gethcommon.HexToAddress("0x343750b5fd3cefbb230e1c2222b42f0db0484ecf"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    24214036,
		},
	},
}

var Beta_OptimismGoerliToArbitrumGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x93c085c1f7770427941bf5983e38be451e29b5d5"),
		OffRamp:      gethcommon.HexToAddress("0xd3e90bb91a40a75e196d054307051141eb6779fb"),
		CommitStore:  gethcommon.HexToAddress("0x5845c5b86e21288dbe65117a7db5795bcaf824fd"),
		PingPongDapp: gethcommon.HexToAddress("0xc2d6c2319a8c142549f6d3b624309b7cc9bc7dbb"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    10273948,
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
		OnRamp:       gethcommon.HexToAddress("0x7923e2b721e28c420227d244d39bd33a8392f9d0"),
		OffRamp:      gethcommon.HexToAddress("0x0d78bf94d9c5ae628abf919492ceb950d76c0be9"),
		CommitStore:  gethcommon.HexToAddress("0x69959fc99b375712d2e3a503bc3599b1354c5f42"),
		PingPongDapp: gethcommon.HexToAddress("0x0db171377d230372f7393860271bd47df53700ad"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    10273761,
		},
	},
}

var Beta_SepoliaToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x270dc4218d51bcaa422af27fa9fef7b1025f4822"),
		OffRamp:      gethcommon.HexToAddress("0x4f060c6e639ecdccd7ba6c040622a1d65c6aee4e"),
		CommitStore:  gethcommon.HexToAddress("0xa3fcf93e44c9fe6ea368b1bc6b9002e3677574c3"),
		PingPongDapp: gethcommon.HexToAddress("0xfcc95cbd997acb6966a461c9adf3227ccb66ca10"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3629174,
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
}
