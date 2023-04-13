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
	rhea.Goerli:         {ChainConfig: Beta_Goerli},
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
		rhea.Goerli:         Beta_OptimismGoerliToGoerli,
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
	rhea.Goerli: {
		rhea.OptimismGoerli: Beta_GoerliToOptimismGoerli,
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
			Pool:          gethcommon.HexToAddress("0xb97cd2f3c35f360559dcc489ef80aeeb4fc558b3"),
			Price:         rhea.LINK.Price(),
			TokenPoolType: rhea.LockRelease,
			PriceFeed: rhea.PriceFeed{
				Aggregator: gethcommon.HexToAddress("0x15608A8A3E2f65e00fe0ef2C9c78Ada4e4E8172E"),
				Multiplier: big.NewInt(1e10),
			},
		},
		rhea.WETH: {
			Token: gethcommon.HexToAddress("0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6"),
			Pool:  gethcommon.HexToAddress("0x38c93ba510903fe75c164ff50b0c5c2dbb3e96ca"),
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
	Router:        gethcommon.HexToAddress("0x5783d3e5976c4fb4f8843f7d71389a67e19d39a6"),
	Afn:           gethcommon.HexToAddress("0x87c4ee46ab7c3cfe9cb1598a1e4e5ab7356688ed"),
	PriceRegistry: gethcommon.HexToAddress("0x7dc83a160bee27f68fb0992fe585cbab8544a5d0"),
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
			Pool:          gethcommon.HexToAddress("0x4c82fb2c3bd2b9afa15c7f66b7bd168900bc8d98"),
			Price:         rhea.LINK.Price(),
			TokenPoolType: rhea.LockRelease,
			PriceFeed: rhea.PriceFeed{
				Aggregator: gethcommon.HexToAddress("0x5A2734CC0341ea6564dF3D00171cc99C63B1A7d3"),
				Multiplier: big.NewInt(1e10),
			},
		},
		rhea.WETH: {
			Token:         gethcommon.HexToAddress("0x097D90c9d3E0B50Ca60e1ae45F6A81010f9FB534"),
			Pool:          gethcommon.HexToAddress("0x91da067b52e73d6f7e33ef2367b05a5157f96fb9"),
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
	Router:        gethcommon.HexToAddress("0xee83867f8039e457054391bcb716395a6b8c428e"),
	Afn:           gethcommon.HexToAddress("0xee1af1b4a9cc059f8baa9c5c9db97fcc8caa3188"),
	PriceRegistry: gethcommon.HexToAddress("0x027291826799343563288c9c8ca3ffef79ecb791"),
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
		EIP1559: false,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:          gethcommon.HexToAddress("0xdc2CC710e42857672E7907CF474a69B63B93089f"),
			Pool:           gethcommon.HexToAddress("0xd798a5f80ec531d7feafad5e008801a96aa1cdf0"),
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
			Pool:           gethcommon.HexToAddress("0xd53ad873bd0aa553dd37d7608101c30138b3e207"),
			Price:          rhea.WETH.Price(),
			TokenPoolType:  rhea.LockRelease,
			TokenPriceType: rhea.TokenPrices,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0xb1495929adb3335594f01fdea3f4c8195f0c4a02"),
	Afn:           gethcommon.HexToAddress("0x91f053f2656f289446203c6118edf958417e4f30"),
	PriceRegistry: gethcommon.HexToAddress("0x3adac3e7a5618c73303fbe4cf0aa42b491e70dce"),
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
			Pool:           gethcommon.HexToAddress("0x58a5654272a89c0d89197ebb87b17423f5ecab92"),
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
			Pool:           gethcommon.HexToAddress("0x99b75c6b0159dcdf4649b75e4c01b8195cad6fcb"),
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
	Router:        gethcommon.HexToAddress("0x28ad7c2990241cd40080e5c638b358230dbee6f0"),
	Afn:           gethcommon.HexToAddress("0x4a9ee459bd6777bfbf05b7ce7f5cf4e4efb8ba08"),
	PriceRegistry: gethcommon.HexToAddress("0x12254ac902b724ceba8a297c711f4d8792ea4414"),
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
			Pool:           gethcommon.HexToAddress("0x25d997d8618e1299418b3d905e40bc353ec89f61"),
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
			Pool:           gethcommon.HexToAddress("0xcc22cd1ad8ae43389eb3577ae576efb99e66be25"),
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
	Router:        gethcommon.HexToAddress("0x3431aa74d9468e3c40ecfb6f3059de4cecf3565f"),
	Afn:           gethcommon.HexToAddress("0x90410a109c952074645b66e149eaac70a91d4f50"),
	PriceRegistry: gethcommon.HexToAddress("0xe5bd0b59ddb5ba47c1d6d8631276209f4d23c6e1"),
	Confirmations: 1,
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
		CommitStore:  gethcommon.HexToAddress("0x6fe86c4f2223536b3b0881922f18bf0bde981ec0"),
		OnRamp:       gethcommon.HexToAddress("0x8bd92cf0ec06a034f76996c190f59495191a224e"),
		OffRamp:      gethcommon.HexToAddress("0x8172fe7a3ea68043231d838fe4558e11c1682c7e"),
		PingPongDapp: gethcommon.HexToAddress("0x0395128fde723b9f7e96dfaedca6e66c0e564b07"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    7946108,
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
		CommitStore:  gethcommon.HexToAddress("0xc55f5c76f01322c1423c0db55d499491b5c8dfc5"),
		OnRamp:       gethcommon.HexToAddress("0xcb1ca6ec34db6f9924c13890d1e3419a31782414"),
		OffRamp:      gethcommon.HexToAddress("0x0d2071ebd17e7f94dac1f8cc50986410454c953c"),
		PingPongDapp: gethcommon.HexToAddress("0x93c085c1f7770427941bf5983e38be451e29b5d5"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    20883681,
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
		CommitStore:  gethcommon.HexToAddress("0x622cb640f52bffa68b78b2bd12c1940ca4899621"),
		OnRamp:       gethcommon.HexToAddress("0x44225eb3b73b1b52dd2ecd258f9b63418ec6bf79"),
		OffRamp:      gethcommon.HexToAddress("0x51158ca439fea9e809bc063cfa6701747b05254e"),
		PingPongDapp: gethcommon.HexToAddress("0x0dbbacafffadd79ac0be0f119fc3b3e46d11e35c"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    16008952,
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
		CommitStore:  gethcommon.HexToAddress("0x5a155e047f3bf7cfe00f828adb3961db9616787c"),
		OnRamp:       gethcommon.HexToAddress("0xe7a9afee8b72411e10f0f15abf84568804dc7202"),
		OffRamp:      gethcommon.HexToAddress("0xc26529293b466c66ef2cd197854a4017ad6adc89"),
		PingPongDapp: gethcommon.HexToAddress("0x32e15870663ee32059a400ec147f88f9d908eb08"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    20883825,
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
		CommitStore:  gethcommon.HexToAddress("0x80c2aa80f202fefdfeef80f516cfd89768c54057"),
		OnRamp:       gethcommon.HexToAddress("0x45c93e0d328572b0e3e64321a49f7405ba6b668f"),
		OffRamp:      gethcommon.HexToAddress("0x2d3d988ee473d144bcccf8c486385e52de9ab395"),
		PingPongDapp: gethcommon.HexToAddress("0x844eeb0a27e24d4d780900dbc88ff3b8bb5b6aa5"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    16009507,
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
		CommitStore:  gethcommon.HexToAddress("0xd80b34a6e98a197f6d2e9e45c915095b8d0a49d0"),
		OnRamp:       gethcommon.HexToAddress("0x14a5b2aaceb95bacb0c30a359fe67611611fbb4d"),
		OffRamp:      gethcommon.HexToAddress("0x21daa4964397ca38607240f6e94d974688e582cb"),
		PingPongDapp: gethcommon.HexToAddress("0xd1032eaedb01ea8fca6a50e771b24842b4a339f9"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    7946529,
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

var Beta_AvaxFujiToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x76e9eaa2b5d45acd73425b0ae47c1237848684a2"),
		OnRamp:       gethcommon.HexToAddress("0x593dccf4f68987ec97c03ee7043143778caf25ee"),
		OffRamp:      gethcommon.HexToAddress("0x68fd744e29da60172f42f4e6334c4cf8a2c8854b"),
		PingPongDapp: gethcommon.HexToAddress("0x0fdc66d7d5d970e44883665e61b09fd8fca98b92"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    20883312,
		},
		CommitOffchainConfig: ccip.CommitOffchainConfig{
			FeeUpdateHeartBeat:    models.MustMakeDuration(24 * time.Hour),
			FeeUpdateDeviationPPB: 5e7,
			MaxGasPrice:           200e9,
		},
		ExecOffchainConfig: ccip.ExecOffchainConfig{
			BatchGasLimit:            5_000_000,
			RelativeBoostPerWaitHour: 0.07,
			MaxGasPrice:              200e9,
		},
	},
}

var Beta_SepoliaToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x248d240b2338349e71db5a8f8e29fa5e55d4ef1b"),
		OnRamp:       gethcommon.HexToAddress("0xd1289a6a5c40b92f7420515ef52d51f0d5f27c0e"),
		OffRamp:      gethcommon.HexToAddress("0xb75ea7e398f3df1eab0836a0bafbdef69b93519d"),
		PingPongDapp: gethcommon.HexToAddress("0x3375faee2987e9c67b46dfa6d6a7cfeb72c544e4"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3277038,
		},
		CommitOffchainConfig: ccip.CommitOffchainConfig{
			FeeUpdateHeartBeat:    models.MustMakeDuration(24 * time.Hour),
			FeeUpdateDeviationPPB: 5e7,
			MaxGasPrice:           200e9,
		},
		ExecOffchainConfig: ccip.ExecOffchainConfig{
			BatchGasLimit:            5_000_000,
			RelativeBoostPerWaitHour: 0.07,
			MaxGasPrice:              200e9,
		},
	},
}

var Beta_OptimismGoerliToGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0xda757342597297a5594186f883a737a74ae8e52d"),
		OnRamp:       gethcommon.HexToAddress("0x963f7948dd7bcc70b0dde93ebd2759f0cd2cc164"),
		OffRamp:      gethcommon.HexToAddress("0xf4fa7a336a70e8b19b2dc4ad1cbb14c2986248f0"),
		PingPongDapp: gethcommon.HexToAddress("0x6446f15d2111724c9a3e63fee79bf6171623698f"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    7946824,
		},
		CommitOffchainConfig: ccip.CommitOffchainConfig{
			FeeUpdateHeartBeat:    models.MustMakeDuration(24 * time.Hour),
			FeeUpdateDeviationPPB: 5e7,
			MaxGasPrice:           200e9,
		},
		ExecOffchainConfig: ccip.ExecOffchainConfig{
			BatchGasLimit:            5_000_000,
			RelativeBoostPerWaitHour: 0.07,
			MaxGasPrice:              200e9,
		},
	},
}

var Beta_GoerliToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_Goerli,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x20224cb68df994a32e2ea26d22098c9a7e0e9667"),
		OnRamp:       gethcommon.HexToAddress("0xcac61d16e62f2c626ccfb4aeaa0ea73457e9b895"),
		OffRamp:      gethcommon.HexToAddress("0xf037523a2a9beafad74e77446deb75d4635853f6"),
		PingPongDapp: gethcommon.HexToAddress("0x219f7b02fa0efac53dd51e9cb050639305378477"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    8818791,
		},
		CommitOffchainConfig: ccip.CommitOffchainConfig{
			FeeUpdateHeartBeat:    models.MustMakeDuration(24 * time.Hour),
			FeeUpdateDeviationPPB: 5e7,
			MaxGasPrice:           200e9,
		},
		ExecOffchainConfig: ccip.ExecOffchainConfig{
			BatchGasLimit:            5_000_000,
			RelativeBoostPerWaitHour: 0.07,
			MaxGasPrice:              200e9,
		},
	},
}

var Beta_OptimismGoerliToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x891410316013b95021bc2abdf1febf45d66647b0"),
		OnRamp:       gethcommon.HexToAddress("0x27f0ee4f431c25ccfc01f66b8a009ceb0e6d9935"),
		OffRamp:      gethcommon.HexToAddress("0xa734577cd70a12313cf1bf8aecb93e3c7d4377da"),
		PingPongDapp: gethcommon.HexToAddress("0xbfb2fef68bb872eac19de9e2061e9f99e7803baa"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    7947157,
		},
		CommitOffchainConfig: ccip.CommitOffchainConfig{
			FeeUpdateHeartBeat:    models.MustMakeDuration(24 * time.Hour),
			FeeUpdateDeviationPPB: 5e7,
			MaxGasPrice:           200e9,
		},
		ExecOffchainConfig: ccip.ExecOffchainConfig{
			BatchGasLimit:            5_000_000,
			RelativeBoostPerWaitHour: 0.07,
			MaxGasPrice:              200e9,
		},
	},
}

var Beta_SepoliaToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x50d289b2ff6605209988092424bf354488a52468"),
		OnRamp:       gethcommon.HexToAddress("0x2eb5554ea058ca44ee18a731b651ae06b5261b7a"),
		OffRamp:      gethcommon.HexToAddress("0xfe01bbad74159b184f5a7351cdd3faddc68ceb89"),
		PingPongDapp: gethcommon.HexToAddress("0x3859c817b9795cb82ced4b1447dc23310e98ce1d"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3277280,
		},
		CommitOffchainConfig: ccip.CommitOffchainConfig{
			FeeUpdateHeartBeat:    models.MustMakeDuration(24 * time.Hour),
			FeeUpdateDeviationPPB: 5e7,
			MaxGasPrice:           200e9,
		},
		ExecOffchainConfig: ccip.ExecOffchainConfig{
			BatchGasLimit:            5_000_000,
			RelativeBoostPerWaitHour: 0.07,
			MaxGasPrice:              200e9,
		},
	},
}

var Beta_ArbitrumGoerliToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_ArbitrumGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0xd97db6611816aa31d085a46a9731e97937a15b6f"),
		OnRamp:       gethcommon.HexToAddress("0x21560b4acaedb8aa2dd935618f15da43197bdc12"),
		OffRamp:      gethcommon.HexToAddress("0x7323c8101e472535647f84199700e687559ce8ea"),
		PingPongDapp: gethcommon.HexToAddress("0x0ad2dc37e1a373b4d66563faf636021d96330c4a"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    16241152,
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

var Beta_SepoliaToArbitrumGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x7ce9b67c7ceff8a0638f71ddfe10179f69649064"),
		OnRamp:       gethcommon.HexToAddress("0x6653c76f928b8f1442e25c8accf479ffb0d4ce96"),
		OffRamp:      gethcommon.HexToAddress("0x71ae425e2631e91464afdb9558caf6f7d66a3634"),
		PingPongDapp: gethcommon.HexToAddress("0xc3a11ef77e260e9bff04b6b2e0a4f1619436a9c2"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3282515,
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
