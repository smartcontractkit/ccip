package deployments

import (
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip/rhea"

	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

var BetaChains = map[rhea.Chain]rhea.EvmDeploymentConfig{
	rhea.AvaxFuji:       {ChainConfig: Beta_AvaxFuji},
	rhea.OptimismGoerli: {ChainConfig: Beta_OptimismGoerli},
	rhea.Sepolia:        {ChainConfig: Beta_Sepolia},
	rhea.ArbitrumGoerli: {ChainConfig: Beta_ArbitrumGoerli},
	rhea.BASEGoerli:     {ChainConfig: Beta_BaseGoerli},
	rhea.BSCTestnet:     {ChainConfig: Beta_BSCTestnet},
	rhea.AvaxAnzSubnet:  {ChainConfig: Beta_AvaxANZTestnet},
}

var BetaChainMapping = map[rhea.Chain]map[rhea.Chain]rhea.EvmDeploymentConfig{
	rhea.AvaxFuji: {
		rhea.OptimismGoerli: Beta_AvaxFujiToOptimismGoerli,
		rhea.ArbitrumGoerli: Beta_AvaxFujiToArbitrumGoerli,
		rhea.Sepolia:        Beta_AvaxFujiToSepolia,
		rhea.BASEGoerli:     Beta_AvaxFujiToBaseGoerli,
		rhea.BSCTestnet:     Beta_AvaxFujiToBscTestnet,
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
		rhea.AvaxAnzSubnet:  Beta_SepoliaToAvaxANZTestnet,
		rhea.BASEGoerli:     Beta_SepoliaToBaseGoerli,
	},
	rhea.BASEGoerli: {
		rhea.AvaxFuji: Beta_BaseGoerliToAvaxFuji,
		rhea.Sepolia:  Beta_BaseGoerliToSepolia,
	},
	rhea.BSCTestnet: {
		rhea.AvaxFuji: Beta_BscTestnetToAvaxFuji,
	},
	rhea.AvaxAnzSubnet: {
		rhea.Sepolia: Beta_AvaxANZTestnetToSepolia,
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
			Pool:          gethcommon.HexToAddress("0xb26776e14f69d8bbbe2403c0aebbb2c108cda039"),
			Price:         rhea.LINK.Price(),
			Decimals:      rhea.LINK.Decimals(),
			TokenPoolType: rhea.LockRelease,
			PriceFeed: rhea.PriceFeed{
				Aggregator: gethcommon.HexToAddress("0x5A2734CC0341ea6564dF3D00171cc99C63B1A7d3"),
				Multiplier: big.NewInt(1e10),
			},
		},
		rhea.WETH: {
			Token:         gethcommon.HexToAddress("0x097D90c9d3E0B50Ca60e1ae45F6A81010f9FB534"),
			Price:         rhea.WETH.Price(),
			Decimals:      rhea.WETH.Decimals(),
			TokenPoolType: rhea.FeeTokenOnly,
			PriceFeed: rhea.PriceFeed{
				Aggregator: gethcommon.HexToAddress("0x719E22E3D4b690E5d96cCb40619180B5427F14AE"),
				Multiplier: big.NewInt(1e10),
			},
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0x77a23a44caf6eb8e6a3dfba286964192bcf62704"),
	ARM:           gethcommon.HexToAddress("0x0e5ce290a70358d31fcd9e3ba29a95822a414ae8"),
	ARMProxy:      gethcommon.HexToAddress("0xe628d44b64fedaf9175b368bc50ccc17f2cc7f9a"),
	PriceRegistry: gethcommon.HexToAddress("0x0ac19ec9069d229586e533b087a05faa2a0a049a"),
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
			Pool:           gethcommon.HexToAddress("0x87b2d4aa134ddf4871aee33764736931c6a55eda"),
			TokenPoolType:  rhea.LockRelease,
			TokenPriceType: rhea.PriceFeeds,
			Price:          rhea.LINK.Price(),
			Decimals:       rhea.LINK.Decimals(),
			PriceFeed: rhea.PriceFeed{
				Aggregator: gethcommon.HexToAddress("0x53AFfFfA77006432146b667C67FA77b5D405793b"),
				Multiplier: big.NewInt(1e10),
			},
		},
		rhea.WETH: {
			Token:          gethcommon.HexToAddress("0x4200000000000000000000000000000000000006"),
			Price:          rhea.WETH.Price(),
			Decimals:       rhea.WETH.Decimals(),
			TokenPoolType:  rhea.FeeTokenOnly,
			TokenPriceType: rhea.TokenPrices,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0xdb12517bc81108f332cfcbe6739c7ab9a4621979"),
	UpgradeRouter: gethcommon.HexToAddress(""),
	ARM:           gethcommon.HexToAddress("0x270dc4218d51bcaa422af27fa9fef7b1025f4822"),
	ARMProxy:      gethcommon.HexToAddress("0xbf93cf175dad548528829d2558445eea200131c9"),
	PriceRegistry: gethcommon.HexToAddress("0x4f060c6e639ecdccd7ba6c040622a1d65c6aee4e"),
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
		DeployARM:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployUpgradeRouter: false,
		DeployPriceRegistry: false,
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
			Pool:           gethcommon.HexToAddress("0xbaa8f352c3eed1451f6cbea3f21e7ede259d203e"),
			Price:          rhea.LINK.Price(),
			Decimals:       rhea.LINK.Decimals(),
			TokenPoolType:  rhea.LockRelease,
			TokenPriceType: rhea.TokenPrices,
			PriceFeed: rhea.PriceFeed{
				Aggregator: gethcommon.HexToAddress("0x5F4a4f309Aefb6fb0Ab927A0421D0342fF92f194"),
				Multiplier: big.NewInt(1e10),
			},
		},
		rhea.WAVAX: {
			Token:          gethcommon.HexToAddress("0xd00ae08403B9bbb9124bB305C09058E32C39A48c"),
			Price:          rhea.WAVAX.Price(),
			Decimals:       rhea.WAVAX.Decimals(),
			TokenPoolType:  rhea.FeeTokenOnly,
			TokenPriceType: rhea.TokenPrices,
			PriceFeed: rhea.PriceFeed{
				Aggregator: gethcommon.HexToAddress("0x6C2441920404835155f33d88faf0545B895871b1"),
				Multiplier: big.NewInt(1e10),
			},
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WAVAX},
	WrappedNative: rhea.WAVAX,
	Router:        gethcommon.HexToAddress("0x42415975c9cbcc0a73f8fa90018b55ed3b239273"),
	UpgradeRouter: gethcommon.HexToAddress(""),
	ARM:           gethcommon.HexToAddress("0x7103f9da75f04dd452b257aaca23842895ae1920"),
	ARMProxy:      gethcommon.HexToAddress("0xb140855a247c139710d17ef0eef71e08b7515952"),
	PriceRegistry: gethcommon.HexToAddress("0x08ebebf6d75aa188bcdf310da231ff07c792eed0"),
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
		DeployARM:           false,
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
			Pool:           gethcommon.HexToAddress("0x35a7cdde653e9bd51a12d95f2db74729f4d77f90"),
			Price:          rhea.LINK.Price(),
			Decimals:       rhea.LINK.Decimals(),
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
			Decimals:       rhea.WETH.Decimals(),
			TokenPoolType:  rhea.FeeTokenOnly,
			TokenPriceType: rhea.TokenPrices,
			PriceFeed: rhea.PriceFeed{
				Aggregator: gethcommon.HexToAddress("0xC975dEfb12C5e83F2C7E347831126cF136196447"),
				Multiplier: big.NewInt(1e10),
			},
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0x4599d97cf3bef6dc61f14d2c275640d1dd47d637"),
	UpgradeRouter: gethcommon.HexToAddress(""),
	ARM:           gethcommon.HexToAddress("0xae4d2c1476db9f88a402805d9937035126c6478e"),
	ARMProxy:      gethcommon.HexToAddress("0xb8a9fdc9c9d504641415d0263b6df0e34ebff296"),
	PriceRegistry: gethcommon.HexToAddress("0x7faddb7dc5bd30b1fc0df403e788dc6b546bf709"),
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
		DeployARM:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployUpgradeRouter: false,
		DeployPriceRegistry: false,
	},
}

var Beta_BaseGoerli = rhea.EVMChainConfig{
	EvmChainId: 84531,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:          gethcommon.HexToAddress("0xd886e2286fd1073df82462ea1822119600af80b6"),
			Pool:           gethcommon.HexToAddress("0x0ff6b6f3ad10d66600fd5cc25b98542a05aa7bc2"),
			TokenPoolType:  rhea.Wrapped,
			TokenPriceType: rhea.TokenPrices,
			Price:          rhea.LINK.Price(),
			Decimals:       rhea.LINK.Decimals(),
		},
		rhea.WETH: {
			Token:          gethcommon.HexToAddress("0x4200000000000000000000000000000000000006"),
			Price:          rhea.WETH.Price(),
			Decimals:       rhea.WETH.Decimals(),
			TokenPoolType:  rhea.FeeTokenOnly,
			TokenPriceType: rhea.TokenPrices,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0x11BEe8AD23bA3Fd56fcbD88467D5C76375fD03ef"),
	ARM:           gethcommon.HexToAddress("0x903f08730dcffc883104c07dd63a86df93472026"),
	ARMProxy:      gethcommon.HexToAddress("0x2ae6d5495fc20226f433be50e37d59c05d186aaa"),
	PriceRegistry: gethcommon.HexToAddress("0xd84e2f21d0bbbe0be697cb4a400af367b39fbcd4"),
	TunableChainValues: rhea.TunableChainValues{
		FinalityDepth:            getFinalityDepth(rhea.BASEGoerli),
		OptimisticConfirmations:  getOptimisticConfirmations(rhea.BASEGoerli),
		BatchGasLimit:            BATCH_GAS_LIMIT,
		RelativeBoostPerWaitHour: RELATIVE_BOOST_PER_WAIT_HOUR,
		FeeUpdateHeartBeat:       models.MustMakeDuration(FEE_UPDATE_HEARTBEAT),
		FeeUpdateDeviationPPB:    FEE_UPDATE_DEVIATION_PPB_FAST_CHAIN,
		MaxGasPrice:              getMaxGasPrice(rhea.BASEGoerli),
		InflightCacheExpiry:      models.MustMakeDuration(INFLIGHT_CACHE_EXPIRY),
		RootSnoozeTime:           models.MustMakeDuration(ROOT_SNOOZE_TIME),
	},
	DeploySettings: rhea.ChainDeploySettings{
		DeployARM:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,
	},
}

var Beta_BSCTestnet = rhea.EVMChainConfig{
	EvmChainId: 97,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:          gethcommon.HexToAddress("0x84b9B910527Ad5C03A9Ca831909E21e236EA7b06"),
			Pool:           gethcommon.HexToAddress("0xd84e2f21d0bbbe0be697cb4a400af367b39fbcd4"),
			TokenPoolType:  rhea.LockRelease,
			TokenPriceType: rhea.TokenPrices,
			Price:          rhea.LINK.Price(),
			Decimals:       rhea.LINK.Decimals(),
		},
		rhea.WBNB: {
			Token:          gethcommon.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"),
			Price:          rhea.WBNB.Price(),
			Decimals:       rhea.WBNB.Decimals(),
			TokenPoolType:  rhea.FeeTokenOnly,
			TokenPriceType: rhea.TokenPrices,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WBNB},
	WrappedNative: rhea.WBNB,
	Router:        gethcommon.HexToAddress("0x25d997d8618e1299418b3d905e40bc353ec89f61"),
	ARM:           gethcommon.HexToAddress("0x3b80b7ef5c00eb892cbe72800c028c47ad6380ef"),
	ARMProxy:      gethcommon.HexToAddress("0x11bee8ad23ba3fd56fcbd88467d5c76375fd03ef"),
	PriceRegistry: gethcommon.HexToAddress("0xc095b651f3a483d74c42f319309e08af676ddcae"),
	TunableChainValues: rhea.TunableChainValues{
		FinalityDepth:            getFinalityDepth(rhea.BSCTestnet),
		OptimisticConfirmations:  getOptimisticConfirmations(rhea.BSCTestnet),
		BatchGasLimit:            BATCH_GAS_LIMIT,
		RelativeBoostPerWaitHour: RELATIVE_BOOST_PER_WAIT_HOUR,
		FeeUpdateHeartBeat:       models.MustMakeDuration(FEE_UPDATE_HEARTBEAT),
		FeeUpdateDeviationPPB:    FEE_UPDATE_DEVIATION_PPB_FAST_CHAIN,
		MaxGasPrice:              getMaxGasPrice(rhea.BSCTestnet),
		InflightCacheExpiry:      models.MustMakeDuration(INFLIGHT_CACHE_EXPIRY),
		RootSnoozeTime:           models.MustMakeDuration(ROOT_SNOOZE_TIME),
	},
	DeploySettings: rhea.ChainDeploySettings{
		DeployARM:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,
	},
}

var Beta_AvaxANZTestnet = rhea.EVMChainConfig{
	EvmChainId: 76578,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:          gethcommon.HexToAddress("0x25d997d8618e1299418b3d905e40bc353ec89f61"),
			Pool:           gethcommon.HexToAddress("0xc095b651f3a483d74c42f319309e08af676ddcae"),
			TokenPoolType:  rhea.Wrapped,
			TokenPriceType: rhea.TokenPrices,
			Price:          rhea.LINK.Price(),
			Decimals:       rhea.LINK.Decimals(),
		},
		rhea.WCBS: {
			Token:          gethcommon.HexToAddress("0xD84E2f21D0bBbe0be697cB4A400AF367b39FbCD4"),
			Price:          rhea.WCBS.Price(),
			Decimals:       rhea.WCBS.Decimals(),
			TokenPoolType:  rhea.FeeTokenOnly,
			TokenPriceType: rhea.TokenPrices,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WCBS},
	WrappedNative: rhea.WCBS,
	Router:        gethcommon.HexToAddress("0xcc22cd1ad8ae43389eb3577ae576efb99e66be25"),
	ARM:           gethcommon.HexToAddress("0x5058af17e36899aa9073c2fe777f3e79ae06c566"),
	ARMProxy:      gethcommon.HexToAddress("0x90410a109c952074645b66e149eaac70a91d4f50"),
	PriceRegistry: gethcommon.HexToAddress("0x338f33f149c9257284a37144e37b1d5a62507a0e"),
	TunableChainValues: rhea.TunableChainValues{
		FinalityDepth:            getFinalityDepth(rhea.AvaxAnzSubnet),
		OptimisticConfirmations:  getOptimisticConfirmations(rhea.AvaxAnzSubnet),
		BatchGasLimit:            BATCH_GAS_LIMIT,
		RelativeBoostPerWaitHour: RELATIVE_BOOST_PER_WAIT_HOUR,
		FeeUpdateHeartBeat:       models.MustMakeDuration(FEE_UPDATE_HEARTBEAT),
		FeeUpdateDeviationPPB:    FEE_UPDATE_DEVIATION_PPB_FAST_CHAIN,
		MaxGasPrice:              getMaxGasPrice(rhea.AvaxAnzSubnet),
		InflightCacheExpiry:      models.MustMakeDuration(INFLIGHT_CACHE_EXPIRY),
		RootSnoozeTime:           models.MustMakeDuration(ROOT_SNOOZE_TIME),
	},
	DeploySettings: rhea.ChainDeploySettings{
		DeployARM:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,
	},
}

var Beta_OptimismGoerliToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x9e107c2051a10cbe0231848023fe7f0bc06c90d4"),
		OffRamp:      gethcommon.HexToAddress("0xb770866b1b092eac9a59423c325f76f823cfbc62"),
		CommitStore:  gethcommon.HexToAddress("0x8b71158cc257f1bb46d76652aa9a37cd726f6ef5"),
		PingPongDapp: gethcommon.HexToAddress("0x362b4f65c64bad166e40d226554e3f258944e35e"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    11780150,
		},
	},
}

var Beta_AvaxFujiToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x1e1f3d8ac7df65fccfcc52dbf03929cee95430ac"),
		OffRamp:      gethcommon.HexToAddress("0xcfafd9dadde8b0d0000a4372e7116562430c4cc6"),
		CommitStore:  gethcommon.HexToAddress("0x80d2c88e42d3a704460adb3a10d95e0302aaf0e2"),
		PingPongDapp: gethcommon.HexToAddress("0xf3253d3ee39b50d3407d9922fe99edb29df815b6"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    24012308,
		},
	},
}

var Beta_ArbitrumGoerliToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_ArbitrumGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x68bfcdd4a6346a4d3dfc0c88dc554e1cf574b94c"),
		OffRamp:      gethcommon.HexToAddress("0x0e45d94bfbc1225bea89b8be46774b167b90c541"),
		CommitStore:  gethcommon.HexToAddress("0xf1cdcf045c6403339ad5b518b6b240b55847abcb"),
		PingPongDapp: gethcommon.HexToAddress("0xafb579a78be6f4d17136bb85a17b20a18bbee70c"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    29964791,
		},
	},
}

var Beta_AvaxFujiToArbitrumGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xff238b3a206b12fd9a437de9f0f351c27d53ba76"),
		OffRamp:      gethcommon.HexToAddress("0xf786080a801a1ea0683d58308cc344a92a6533d4"),
		CommitStore:  gethcommon.HexToAddress("0xf45b41e4af72c86046620f53bef9d210dd9066d1"),
		PingPongDapp: gethcommon.HexToAddress("0x6a01d2212b6ecb9afbae048355e74b989cdeb6ab"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    24012476,
		},
	},
}

var Beta_ArbitrumGoerliToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_ArbitrumGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress(""),
		OffRamp:      gethcommon.HexToAddress(""),
		CommitStore:  gethcommon.HexToAddress(""),
		PingPongDapp: gethcommon.HexToAddress(""),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    25966637,
		},
	},
}

var Beta_OptimismGoerliToArbitrumGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress(""),
		OffRamp:      gethcommon.HexToAddress(""),
		CommitStore:  gethcommon.HexToAddress(""),
		PingPongDapp: gethcommon.HexToAddress(""),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    10656674,
		},
	},
}

var Beta_AvaxFujiToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xcf5688404c589dae8fe53a8a65e9794b35703338"),
		OffRamp:      gethcommon.HexToAddress("0x2395c1c5153483ed256b1115878d7ab2b25a6a18"),
		CommitStore:  gethcommon.HexToAddress("0x1e1ae8a4fc42371b675573a0dcd4590e7ba9bb32"),
		PingPongDapp: gethcommon.HexToAddress("0xa082d42bb1a3badd7a7ec7f9b45b9c951ba03dfe"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    23684025,
		},
	},
}

var Beta_SepoliaToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x04b1b3d66c1bff50e2b8ac9f46ba061e535e64e5"),
		OffRamp:      gethcommon.HexToAddress("0xa9c6552677f6bb82afb0b783aa03b6d9812df3cc"),
		CommitStore:  gethcommon.HexToAddress("0x9e50655259f419c06382afd26ad2b3d44f5262b9"),
		PingPongDapp: gethcommon.HexToAddress("0x51896d2d175857a53afc4d4d4b81fdb2b20d9491"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3804307,
		},
	},
}

var Beta_OptimismGoerliToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x9eb32f3fb9eb0f764a0b16167a033b4e0f624fee"),
		OffRamp:      gethcommon.HexToAddress("0xb66b33513abab52611d9ba025fdb90e839096091"),
		CommitStore:  gethcommon.HexToAddress("0x682c35dc98c2a11829d76bdc0917b28156fb1375"),
		PingPongDapp: gethcommon.HexToAddress("0xe08f435e7d0c51e1e9d1384bf699a9eae6bea701"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    11399246,
		},
	},
}

var Beta_SepoliaToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x50d91e49e1075f021d45ad2d27293be3e08ffc4c"),
		OffRamp:      gethcommon.HexToAddress("0x5bffbf67bb4f4f6837b03017153383076caf34bc"),
		CommitStore:  gethcommon.HexToAddress("0xc307ccd8bf12b5de46443d283f0a23316cfd068d"),
		PingPongDapp: gethcommon.HexToAddress("0x95c47b6f1e8c63634f23f4a87f211c0fbcbefea3"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3804366,
		},
	},
}

var Beta_ArbitrumGoerliToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_ArbitrumGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress(""),
		OffRamp:      gethcommon.HexToAddress(""),
		CommitStore:  gethcommon.HexToAddress(""),
		PingPongDapp: gethcommon.HexToAddress(""),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    26819726,
		},
	},
}

var Beta_SepoliaToArbitrumGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress(""),
		OffRamp:      gethcommon.HexToAddress(""),
		CommitStore:  gethcommon.HexToAddress(""),
		PingPongDapp: gethcommon.HexToAddress(""),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3724273,
		},
	},
}

var Beta_AvaxFujiToBaseGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x9a25ab1f6007a45eedd159d35360bb66e8b06f85"),
		OffRamp:      gethcommon.HexToAddress("0x63eefe6734c606b60a25fb3fb766f1f73fd5b785"),
		CommitStore:  gethcommon.HexToAddress("0x4157f0856703a670cd1eab154bbae7f48410aef4"),
		PingPongDapp: gethcommon.HexToAddress("0x852eb3ae1d95e90c5f6b55d6559f58495c669468"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    24534985,
		},
	},
}

var Beta_BaseGoerliToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_BaseGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x479a6ff9a2335180d92306d76754bbbdca751bf0"),
		OffRamp:      gethcommon.HexToAddress("0xd84a16dd4235aebca9515dc3cb5cd1665f74cb2e"),
		CommitStore:  gethcommon.HexToAddress("0x37004c1245a2d5541377e87ca29699492a4114d5"),
		PingPongDapp: gethcommon.HexToAddress("0x67eee3b47c841da57c46eba88177725fa018a4aa"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    7628508,
		},
	},
}

var Beta_SepoliaToBaseGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xe3d9613aabd4739015fa16b80b05d89c55edbfdc"),
		OffRamp:      gethcommon.HexToAddress("0x1eff05a31ba79f40ac8fc7084ffe9cd30e6074d1"),
		CommitStore:  gethcommon.HexToAddress("0x28f6189d6316c2fbc09c4dc20273d92de0f56492"),
		PingPongDapp: gethcommon.HexToAddress("0xdf8831ddfb14561d05679f08086ec5432fcf8502"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    4113802,
		},
	},
}

var Beta_BaseGoerliToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_BaseGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xdae257e1acf6665eb897bfbf8dd1bf62825e87dd"),
		OffRamp:      gethcommon.HexToAddress("0x0d52c2472dc6f37fbee59552b1165deafe3b9a4b"),
		CommitStore:  gethcommon.HexToAddress("0xd58cc39f4273497f686d94dc3e614bc26f5003d9"),
		PingPongDapp: gethcommon.HexToAddress("0xcaf37aa8cad1bb7964e4a918660d2f341030d740"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    8590522,
		},
	},
}

var Beta_AvaxFujiToBscTestnet = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0xc8ab64477b14baa744ffba0d783847d05b8da0dc"),
		OffRamp:      gethcommon.HexToAddress("0x59cf18505b5ac6a7e4aa5213d2d5d34d9e0004bd"),
		CommitStore:  gethcommon.HexToAddress("0x56fe1342714a4a3f0ad65b1b19cbe4cdacc042b4"),
		PingPongDapp: gethcommon.HexToAddress("0x1d2b308e029ec587729f1cab2ca9663cafec1f5f"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    24674403,
		},
	},
}

var Beta_BscTestnetToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_BSCTestnet,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress("0x3431aa74d9468e3c40ecfb6f3059de4cecf3565f"),
		OffRamp:      gethcommon.HexToAddress("0x11ac1567069e540920ee5d5ea4b3d44bc6c6d10b"),
		CommitStore:  gethcommon.HexToAddress("0x99ce75105d6a882af40cd5f6166a9564b3003a07"),
		PingPongDapp: gethcommon.HexToAddress("0x51158ca439fea9e809bc063cfa6701747b05254e"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    32089112,
		},
	},
}

var Beta_AvaxANZTestnetToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_AvaxANZTestnet,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress(""),
		OffRamp:      gethcommon.HexToAddress(""),
		CommitStore:  gethcommon.HexToAddress(""),
		PingPongDapp: gethcommon.HexToAddress(""),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    0,
		},
	},
}
var Beta_SepoliaToAvaxANZTestnet = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:       gethcommon.HexToAddress(""),
		OffRamp:      gethcommon.HexToAddress(""),
		CommitStore:  gethcommon.HexToAddress(""),
		PingPongDapp: gethcommon.HexToAddress(""),
		DeploySettings: rhea.LaneDeploySettings{
			DeployLane:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    0,
		},
	},
}
