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
}

var BetaChainMapping = map[rhea.Chain]map[rhea.Chain]rhea.EvmDeploymentConfig{
	rhea.AvaxFuji: {
		rhea.OptimismGoerli: Beta_AvaxFujiToOptimismGoerli,
	},
	rhea.OptimismGoerli: {
		rhea.AvaxFuji: Beta_OptimismGoerliToAvaxFuji,
	},
}

var Beta_OptimismGoerli = rhea.EVMChainConfig{
	EvmChainId:    420,
	ChainSelector: 420,
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
	TunableChainValues: rhea.TunableChainValues{
		BlockConfirmations:       getBlockConfirmations(rhea.OptimismGoerli),
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
		DeployedAtBlock:     0,
	},
}

var Beta_AvaxFuji = rhea.EVMChainConfig{
	EvmChainId:    43113,
	ChainSelector: 43113,
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
	TunableChainValues: rhea.TunableChainValues{
		BlockConfirmations:       getBlockConfirmations(rhea.AvaxFuji),
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
		DeployPriceRegistry: false,
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
	},
}
