package deployments

import (
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

var ProdChains = map[rhea.Chain]rhea.EvmDeploymentConfig{
	rhea.AvaxFuji:       {ChainConfig: Prod_AvaxFuji},
	rhea.OptimismGoerli: {ChainConfig: Prod_OptimismGoerli},
	rhea.Sepolia:        {ChainConfig: Prod_Sepolia},
	rhea.Goerli:         {ChainConfig: Prod_Goerli},
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
		rhea.Goerli:   Prod_OptimismGoerliToGoerli,
	},
	rhea.Goerli: {
		rhea.OptimismGoerli: Prod_GoerliToOptimismGoerli,
	},
}

var Prod_Goerli = rhea.EVMChainConfig{
	EvmChainId: 5,
	GasSettings: rhea.EVMGasSettings{
		EIP1559:   true,
		GasTipCap: rhea.DefaultGasTipFee,
	},
	AllowList: []gethcommon.Address{},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:         gethcommon.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB"),
			Pool:          gethcommon.HexToAddress("0x9Ef131613079733Da157D7EB8FFB41f1D7CA869F"),
			Price:         rhea.LINK.Price(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.WETH: {
			Token:         gethcommon.HexToAddress("0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6"),
			Pool:          gethcommon.HexToAddress("0xCc4fa89b93E3203C96B3e4137153bBE13e7f9255"),
			Price:         rhea.WETH.Price(),
			TokenPoolType: rhea.LockRelease,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0x7d870741ca453ed5fa6808f8197664566d80c39e"),
	ARM:           gethcommon.HexToAddress("0xb192F3b22f24d6DaD266DA5ECd7361D6e6534B49"),
	PriceRegistry: gethcommon.HexToAddress("0x99d13D346f4D35F5139A5D51671563eb9a7e09a6"),
	TunableChainValues: rhea.TunableChainValues{
		FinalityDepth:            getFinalityDepth(rhea.Goerli),
		OptimisticConfirmations:  getOptimisticConfirmations(rhea.Goerli),
		BatchGasLimit:            BATCH_GAS_LIMIT,
		RelativeBoostPerWaitHour: RELATIVE_BOOST_PER_WAIT_HOUR,
		FeeUpdateHeartBeat:       models.MustMakeDuration(FEE_UPDATE_HEARTBEAT),
		FeeUpdateDeviationPPB:    FEE_UPDATE_DEVIATION_PPB,
		MaxGasPrice:              getMaxGasPrice(rhea.Goerli),
		InflightCacheExpiry:      models.MustMakeDuration(INFLIGHT_CACHE_EXPIRY),
		RootSnoozeTime:           models.MustMakeDuration(ROOT_SNOOZE_TIME),
	},
	DeploySettings: rhea.ChainDeploySettings{
		DeployARM:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,
		DeployedAtBlock:     0,
	},
}

var Prod_Sepolia = rhea.EVMChainConfig{
	EvmChainId: 11155111,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	AllowList: []gethcommon.Address{},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:         gethcommon.HexToAddress("0x779877A7B0D9E8603169DdbD7836e478b4624789"),
			Pool:          gethcommon.HexToAddress("0x401dA48dB998Fa1A1ba108eDFe06334aB271F501"),
			Price:         rhea.LINK.Price(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.WETH: {
			Token:         gethcommon.HexToAddress("0x097D90c9d3E0B50Ca60e1ae45F6A81010f9FB534"),
			Pool:          gethcommon.HexToAddress("0x5e7AdeF4C8fe5145d813b1EE5c55233A3EeCa0B4"),
			Price:         rhea.WETH.Price(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.CACHEGOLD: {
			Token:         gethcommon.HexToAddress("0x997BCCAE553112CD023592691d41687a3f1EfA7C"),
			Pool:          gethcommon.HexToAddress("0x7500A909CeBEc7F211ae021FaF4720dE3ca13d18"),
			Price:         rhea.CACHEGOLD.Price(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.ANZ: {
			Token:         gethcommon.HexToAddress("0x92eA346B7a2AaB84e6AaB03b80E2421eeFB04685"),
			Pool:          gethcommon.HexToAddress("0xB70cdf1876eB92A99FD7c24205D87693877Aed3E"),
			Price:         rhea.ANZ.Price(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.InsurAce: {
			Token:         gethcommon.HexToAddress("0xb7c8bCA891143221a34DB60A26639785C4839040"),
			Pool:          gethcommon.HexToAddress("0xEC6d1eC94D518be47DA1cb35F5d43286558d8B62"),
			Price:         rhea.InsurAce.Price(),
			TokenPoolType: rhea.LockRelease,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0x0a36795b3006f50088c11ea45b960a1b0406f03b"),
	ARM:           gethcommon.HexToAddress("0xDB81c131193263314762CfCE59B9A057ae7dbB41"),
	PriceRegistry: gethcommon.HexToAddress("0xd824DA3583F02183FaC25669315B004E977780FD"),
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
		DeployARM:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,
		DeployedAtBlock:     0,
	},
}

var Prod_OptimismGoerli = rhea.EVMChainConfig{
	EvmChainId: 420,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: true,
	},
	AllowList: []gethcommon.Address{},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:         gethcommon.HexToAddress("0xdc2CC710e42857672E7907CF474a69B63B93089f"),
			Pool:          gethcommon.HexToAddress("0x4Dca657257f6392922e1834183A27daEaD2c8D62"),
			Price:         rhea.LINK.Price(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.WETH: {
			Token:         gethcommon.HexToAddress("0x4200000000000000000000000000000000000006"),
			Pool:          gethcommon.HexToAddress("0x26dA7Ab49296c31D0AcD3BB93Ed80fEF6943d488"),
			Price:         rhea.WETH.Price(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.CACHEGOLD: {
			Token:         gethcommon.HexToAddress("0xa6446C6f492f31A33bC68249ae59F8871123a777"),
			Pool:          gethcommon.HexToAddress("0xe05e9822466b6fda15b03f02fb47cabe7943ca45"),
			Price:         rhea.CACHEGOLD.Price(),
			TokenPoolType: rhea.BurnMint,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0xec6d1ec94d518be47da1cb35f5d43286558d8b62"),
	ARM:           gethcommon.HexToAddress("0x8E8cD3608AFce92b902B74F577B1429Ce7BcCa96"),
	PriceRegistry: gethcommon.HexToAddress("0x737bbAEb317993e4450fD19E844bDC145aef5adA"),
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
	CustomerSettings: rhea.CustomerSettings{
		CacheGoldFeeAddress:  gethcommon.HexToAddress("0x8264AcEE321ac02549aff7fA05A4Ae7a2e92A6f1"),
		CacheGoldFeeEnforcer: gethcommon.HexToAddress("0x194E7a932663f11AC0790bfC44dBdd8339f0ED65"),
	},
	DeploySettings: rhea.ChainDeploySettings{
		DeployARM:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,
		DeployedAtBlock:     0,
	},
}

var Prod_AvaxFuji = rhea.EVMChainConfig{
	EvmChainId: 43113,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	AllowList: []gethcommon.Address{},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:         gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846"),
			Pool:          gethcommon.HexToAddress("0x4A78C7d84f1E58A532fE569a53f3B14F7e2Cce2d"),
			Price:         rhea.LINK.Price(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.WAVAX: {
			Token:         gethcommon.HexToAddress("0xd00ae08403B9bbb9124bB305C09058E32C39A48c"),
			Pool:          gethcommon.HexToAddress("0x153494a5e36072C2769E2eF49674C313cb3596A3"),
			Price:         rhea.WAVAX.Price(),
			TokenPoolType: rhea.LockRelease,
		},
		rhea.CACHEGOLD: {
			Token:         gethcommon.HexToAddress("0xD16eD805F3eCe986d9541afaD3E59De2F3732517"),
			Pool:          gethcommon.HexToAddress("0xd48b9213583074f518d8f4336fdf35370d450132"),
			Price:         rhea.CACHEGOLD.Price(),
			TokenPoolType: rhea.BurnMint,
		},
		rhea.ANZ: {
			Token:         gethcommon.HexToAddress("0x169d58fd58d598dd7106082b0a43d430d2fec75f"),
			Pool:          gethcommon.HexToAddress("0x169d58fd58d598dd7106082b0a43d430d2fec75f"),
			Price:         rhea.ANZ.Price(),
			TokenPoolType: rhea.Wrapped,
		},
		rhea.InsurAce: {
			Token:         gethcommon.HexToAddress("0xa5f97bc69bf06e7c37b93265c5457420a92c5f4b"),
			Pool:          gethcommon.HexToAddress("0xa5f97bc69bf06e7c37b93265c5457420a92c5f4b"),
			Price:         rhea.InsurAce.Price(),
			TokenPoolType: rhea.Wrapped,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WAVAX},
	WrappedNative: rhea.WAVAX,
	Router:        gethcommon.HexToAddress("0xb352e636f4093e4f5a4ac903064881491926aaa9"),
	ARM:           gethcommon.HexToAddress("0x5fCC6941D1685C7115e257CDcBda258aF85F0C83"),
	PriceRegistry: gethcommon.HexToAddress("0x8D8cEE7D59D967b12A40330AE9F6CC15578073bb"),
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
	CustomerSettings: rhea.CustomerSettings{
		CacheGoldFeeAddress:  gethcommon.HexToAddress("0x8264AcEE321ac02549aff7fA05A4Ae7a2e92A6f1"),
		CacheGoldFeeEnforcer: gethcommon.HexToAddress("0x194E7a932663f11AC0790bfC44dBdd8339f0ED65"),
	},
	DeploySettings: rhea.ChainDeploySettings{
		DeployARM:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,
		DeployedAtBlock:     0,
	},
}

var Prod_SepoliaToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x8febc74c26129c8d7e60288c6dccc75eb494aa3c"),
		OnRamp:       gethcommon.HexToAddress("0xd5685740e7c25315ea6712645ee25de7d8712e16"),
		OffRamp:      gethcommon.HexToAddress("0x0a750ca77369e03613d7640548f4b2b1c695c3bb"),
		PingPongDapp: gethcommon.HexToAddress("0x6ac3e353d1ddda24d5a5416024d6e436b8817a4e"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3127829,
		},
	},
}

var Prod_AvaxFujiToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x4b025af64e5676d9607cf0b52eed216b7915c0e5"),
		OnRamp:       gethcommon.HexToAddress("0x2b54bf278431e8a0299f6048e6adf4c28ba04c2d"),
		OffRamp:      gethcommon.HexToAddress("0xe4620ce35bac283a21d4b0f10347fe2be2a70569"),
		PingPongDapp: gethcommon.HexToAddress("0xd5d0ae131b609500b912daf7b631ff84b8c5d61c"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    20040846,
		},
	},
}

var Prod_SepoliaToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x92c55b159f45648957f32c8a017ac7d62b16e1f7"),
		OnRamp:       gethcommon.HexToAddress("0x24c79d28e0380230265f772bc3babc7a4ed9c9f4"),
		OffRamp:      gethcommon.HexToAddress("0xd2ca6d383917259dc2445bd6470dbad386d4b67a"),
		PingPongDapp: gethcommon.HexToAddress("0x1925a076933c5a587a72dc2c5ff3d737bbc80fa3"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3127868,
		},
	},
}

var Prod_OptimismGoerliToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x6f344d08ae21382c2bae9a5b69bc233b93c6b953"),
		OnRamp:       gethcommon.HexToAddress("0x87bec45564b337384ad9cc23be49e3f2c813a8c2"),
		OffRamp:      gethcommon.HexToAddress("0x36cbeb4723adb23d24169f1ffcc023e8cfa37288"),
		PingPongDapp: gethcommon.HexToAddress("0xa512a49ea3dfe90e4bed31bd8edfc750e08ee3fe"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    6939370,
		},
	},
}

var Prod_OptimismGoerliToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0xc1c76a8c5bfde1be034bbcd930c668726e7c1987"),
		OnRamp:       gethcommon.HexToAddress("0x674fcad1a94f611a6a15b995e05707011c29606b"),
		OffRamp:      gethcommon.HexToAddress("0x1d507e9c72cb99538db5fb05e515b9e01ba7e290"),
		PingPongDapp: gethcommon.HexToAddress("0x7cf8b3149949ea4608239eebfe0352b884fd272b"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    6939519,
		},
	},
}

var Prod_AvaxFujiToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x4e0867356cb9f557b496d7242239b0ccc630adec"),
		OnRamp:       gethcommon.HexToAddress("0x5b1e3838f69bb0724be96d39b44bb5b1e53f87aa"),
		OffRamp:      gethcommon.HexToAddress("0x3211973c6e09945c27d9d05818814c3cbe0d621e"),
		PingPongDapp: gethcommon.HexToAddress("0xfe7fb8b832261a8487acbca838a2e9e64ea1acc5"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    20041225,
		},
	},
}

var Prod_OptimismGoerliToGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x7d2b06444d77b585b3ee65a796e9ba02aaab104a"),
		OnRamp:       gethcommon.HexToAddress("0xf4f64a9afd1a3df2999838fb5459fbde1e8df073"),
		OffRamp:      gethcommon.HexToAddress("0xb2883bfa88181e9f1a6233f19e18167ec57fcff0"),
		PingPongDapp: gethcommon.HexToAddress("0xafacc40e59bb7c7a59c8c160c4893249648ad3c0"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    6939814,
		},
	},
}

var Prod_GoerliToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Goerli,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0xcbcb01b21c085d0f726a9dcb3fd691eb09be8fda"),
		OnRamp:       gethcommon.HexToAddress("0xc913a04c9c2b07a19adc6f714d155c2e8d516f2e"),
		OffRamp:      gethcommon.HexToAddress("0x4dcf71b5abb94da8e74e271e11c366e122d22e7b"),
		PingPongDapp: gethcommon.HexToAddress("0xffc5a11dc009bc36ffe76929fa2d83a2ea4ba1c3"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    8686987,
		},
	},
}
