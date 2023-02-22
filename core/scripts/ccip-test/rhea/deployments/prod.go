package deployments

import (
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
)

// Chains
var Prod_ChainConfigs = []rhea.EvmDeploymentConfig{
	{ChainConfig: Prod_Sepolia},
	{ChainConfig: Prod_AvaxFuji},
	{ChainConfig: Prod_OptimismGoerli},
}

var Prod_Sepolia = rhea.EVMChainConfig{
	ChainId: 11155111,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:                gethcommon.HexToAddress("0x779877A7B0D9E8603169DdbD7836e478b4624789"),
			Pool:                 gethcommon.HexToAddress("0x1e65be3083be02fef86531205fc68bc17288fe65"),
			Price:                big.NewInt(10),
			PriceFeedsAggregator: gethcommon.HexToAddress("0xc59E3633BAAC79493d908e63626716e204A45EdF"),
		},
		rhea.WETH: {
			Token: gethcommon.HexToAddress("0x097D90c9d3E0B50Ca60e1ae45F6A81010f9FB534"),
			Pool:  gethcommon.HexToAddress(""),
			Price: big.NewInt(1500),
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0x428c4dc89b6bf908b82d77c9cbcea786ea8cc7d0"),
	Afn:           gethcommon.HexToAddress("0x89d17571db7c9540eeb36760e3c749c8fb984569"),
	PriceRegistry: gethcommon.HexToAddress("0x1f0e1ef0928c32fbb36c07054893dca7b1c9cd75"),
}

var Prod_OptimismGoerli = rhea.EVMChainConfig{
	ChainId: 420,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: true,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token: gethcommon.HexToAddress("0xdc2CC710e42857672E7907CF474a69B63B93089f"),
			Pool:  gethcommon.HexToAddress("0x6ceabcbe4f904ff5ad0d827a9346863bdcf565ce"),
			Price: big.NewInt(10),
		},
		rhea.WETH: {
			Token: gethcommon.HexToAddress("0x4200000000000000000000000000000000000006"),
			Pool:  gethcommon.HexToAddress("0xde8d0f47a71ea3fdfbd3162271652f2847939097"),
			Price: big.NewInt(1500),
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0xf01ebcd92bc2a7c4e58b5c1527d6814f47443232"),
	Afn:           gethcommon.HexToAddress("0x6d0e94e9f06cafb6031410de7df19fe2f286566c"),
	PriceRegistry: gethcommon.HexToAddress("0x583e761181801f2b43efb1a9489db1a7c78fb60f"),
}

var Prod_AvaxFuji = rhea.EVMChainConfig{
	ChainId: 43113,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:                gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846"),
			Pool:                 gethcommon.HexToAddress("0x7ea5f675272bbc19830e24b0f098626498246f4e"),
			Price:                big.NewInt(10),
			PriceFeedsAggregator: gethcommon.HexToAddress("0x34C4c526902d88a3Aa98DB8a9b802603EB1E3470"),
		},
		rhea.WAVAX: {
			Token:                gethcommon.HexToAddress("0xd00ae08403B9bbb9124bB305C09058E32C39A48c"),
			Pool:                 gethcommon.HexToAddress("0x7d870741ca453ed5fa6808f8197664566d80c39e"),
			Price:                big.NewInt(25),
			PriceFeedsAggregator: gethcommon.HexToAddress("0x6C2441920404835155f33d88faf0545B895871b1"),
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WAVAX},
	WrappedNative: rhea.WAVAX,
	Router:        gethcommon.HexToAddress("0x3e4472e98997564378070308fce2a958a66194df"),
	Afn:           gethcommon.HexToAddress("0x0fa77a2df96f59ab60f440af790ed74eb7d16128"),
	PriceRegistry: gethcommon.HexToAddress("0x45dccaec07a3241a1492e23ea055fb8d501d07e2"),
}

// Lanes
var Prod_SepoliaToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0xc3184d8c25630ad8f3bf486c0f2d704b09263d69"),
		OnRamp:       gethcommon.HexToAddress("0x1041d5b06b9f1e73a4d6ae4f4d2a3d5c9e0aa88e"),
		OffRamp:      gethcommon.HexToAddress("0x95ab1853c803c740e7b095776b217f0e8cbd2e16"),
		ReceiverDapp: gethcommon.HexToAddress("0xa98fa8a008371b9408195e52734b1768c0d1cb5c"),
		PingPongDapp: gethcommon.HexToAddress("0x177e068bc512ad99ec73db6feb7c731d9fea0cb3"),
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
		DeployedAt:           2808574,
	},
}

var Prod_AvaxFujiToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0xc3bf411c3b304ab823aaa05e9a2d1b0c8dc0bff8"),
		OnRamp:       gethcommon.HexToAddress("0x5779309e9bec2f2baae053f779c941c4c1ae08cb"),
		OffRamp:      gethcommon.HexToAddress("0x41172db9cc50a6861910b5c62b323756b05ea9c3"),
		ReceiverDapp: gethcommon.HexToAddress("0x09daf7ec6601e08a1add371011aab0b8a7953dbf"),
		PingPongDapp: gethcommon.HexToAddress("0x2dfe0cf6b49b4f3577c09fb1581ee8e1c9b088aa"),
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
		DeployedAt:           18523284,
	},
}

var Prod_SepoliaToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0xaddd890488aea14bbc8bb3d400a68d265b14c777"),
		OnRamp:       gethcommon.HexToAddress("0x063de5a3f43ee3b8ef2099775f69f4ed03712923"),
		OffRamp:      gethcommon.HexToAddress("0x4d5f14561e949127bf435d2ac884f512f4c058fc"),
		ReceiverDapp: gethcommon.HexToAddress("0x9227faeb95c68cde77783577c25b26568b517d4e"),
		PingPongDapp: gethcommon.HexToAddress("0xd13a553a78c297a4dc3b4f606c289b8dde78933c"),
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
		DeployedAt:           2808895,
	},
}

var Prod_OptimismGoerliToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x81bbb265c82d424711ce504bda91b7c290f0e5f5"),
		OnRamp:       gethcommon.HexToAddress("0x92c55b159f45648957f32c8a017ac7d62b16e1f7"),
		OffRamp:      gethcommon.HexToAddress("0xa0e0b23552dccca01087b6536645f6ab6e433072"),
		ReceiverDapp: gethcommon.HexToAddress("0x754617445703796d9e28631d6109935b15450b28"),
		PingPongDapp: gethcommon.HexToAddress("0x39dd66df77c65fa0c0c9b56c715c157dac666c7a"),
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
		DeployedAt:           4874057,
	},
}

var Prod_OptimismGoerliToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x50f4dc0953544082d77d1e0dafc4a70d7373e934"),
		OnRamp:       gethcommon.HexToAddress("0x3dddbe9463598a014fb2b2bce875743777c2da46"),
		OffRamp:      gethcommon.HexToAddress("0x0a750ca77369e03613d7640548f4b2b1c695c3bb"),
		ReceiverDapp: gethcommon.HexToAddress("0x6ac3e353d1ddda24d5a5416024d6e436b8817a4e"),
		PingPongDapp: gethcommon.HexToAddress("0xb5Ba0846a222e8Cc1189480bde801C38F9f392b9"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,

		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployPingPongDapp:   true,
		DeployGovernanceDapp: false,
		DeployedAt:           4874027,
	},
}

var Prod_AvaxFujiToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x88bf5c5c752fc65ba25ae2356a728edb78041f36"),
		OnRamp:       gethcommon.HexToAddress("0x2d6d77554cf42a81e4abb69857ad9b4ef78d6127"),
		OffRamp:      gethcommon.HexToAddress("0x40f4795eee96dc3e11f931e7b4132a930a3245d0"),
		ReceiverDapp: gethcommon.HexToAddress("0x315d63683c0a8ad3acc8604cd38735a6c19fa8f3"),
		PingPongDapp: gethcommon.HexToAddress("0x9BD470FC8Bea5d3e6985889C442522896b7Db650"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,

		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployPingPongDapp:   true,
		DeployGovernanceDapp: false,
		DeployedAt:           18504934,
	},
}
