package deployments

import (
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
)

var BetaChains = map[rhea.Chain]rhea.EvmDeploymentConfig{
	rhea.AvaxFuji:       {ChainConfig: Beta_AvaxFuji},
	rhea.OptimismGoerli: {ChainConfig: Beta_OptimismGoerli},
	rhea.Sepolia:        {ChainConfig: Beta_Sepolia},
}

var BetaChainMapping = map[rhea.Chain]map[rhea.Chain]rhea.EvmDeploymentConfig{
	rhea.Sepolia: {
		rhea.AvaxFuji:       Beta_SepoliaToAvaxFuji,
		rhea.OptimismGoerli: Beta_SepoliaToOptimismGoerli,
	},
	rhea.AvaxFuji: {
		rhea.Sepolia:        Beta_AvaxFujiToSepolia,
		rhea.OptimismGoerli: Beta_AvaxFujiToOptimismGoerli,
	},
	rhea.OptimismGoerli: {
		rhea.Sepolia:  Beta_OptimismGoerliToSepolia,
		rhea.AvaxFuji: Beta_OptimismGoerliToAvaxFuji,
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
			Pool:          gethcommon.HexToAddress("0xcd4dd77e0218e29540a4e68b4f695e2a96c871ba"),
			Price:         TokenPrices[rhea.LINK],
			TokenPoolType: rhea.LockRelease,
		},
		rhea.WETH: {
			Token:         gethcommon.HexToAddress("0x097D90c9d3E0B50Ca60e1ae45F6A81010f9FB534"),
			Pool:          gethcommon.HexToAddress("0xec2814587f35f7e55c8d5c949ee2fcf90f4766c9"),
			Price:         TokenPrices[rhea.WETH],
			TokenPoolType: rhea.LockRelease,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0xc4f3c3bb9e58ab406450cc1704f20f94e60b02f3"),
	Afn:           gethcommon.HexToAddress("0xfa9d801c46ab9bc2885efe280251ef12861bc373"),
	PriceRegistry: gethcommon.HexToAddress("0xa1df55dafc3be9cbae1fcc6fba3c76a684becab1"),
	Confirmations: 4,
	DeploySettings: rhea.ChainDeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,
		DeployedAtBlock:     2915166,
	},
}

var Beta_OptimismGoerli = rhea.EVMChainConfig{
	ChainId: 420,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: true,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:         gethcommon.HexToAddress("0xdc2CC710e42857672E7907CF474a69B63B93089f"),
			Pool:          gethcommon.HexToAddress("0x578373d72e9d6a259a3755b6E0e88cD4065B8dEB"),
			Price:         TokenPrices[rhea.LINK],
			TokenPoolType: rhea.LockRelease,
		},
		rhea.WETH: {
			Token:         gethcommon.HexToAddress("0x4200000000000000000000000000000000000006"),
			Pool:          gethcommon.HexToAddress("0xED175544847688ad9990Ab27080Bc257025cDf51"),
			Price:         TokenPrices[rhea.WETH],
			TokenPoolType: rhea.LockRelease,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	WrappedNative: rhea.WETH,
	Router:        gethcommon.HexToAddress("0xebeb45ef02491dcf4400e26dfe5b25f4fb1bcdf7"),
	Afn:           gethcommon.HexToAddress("0x4C795713e8c1C048c256c3f6Ec1DFe7f55834D72"),
	PriceRegistry: gethcommon.HexToAddress("0xF05814a007661EFbc8AdBac9ac1FDa6Ee574B5A1"),
	Confirmations: 4,
	DeploySettings: rhea.ChainDeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,
		DeployedAtBlock:     6816841,
	},
}

var Beta_AvaxFuji = rhea.EVMChainConfig{
	ChainId: 43113,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token:         gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846"),
			Pool:          gethcommon.HexToAddress("0x48354c7bfb2e6ce52e0c01279bf1101e642fcf76"),
			Price:         TokenPrices[rhea.LINK],
			TokenPoolType: rhea.LockRelease,
		},
		rhea.WAVAX: {
			Token:         gethcommon.HexToAddress("0xd00ae08403B9bbb9124bB305C09058E32C39A48c"),
			Pool:          gethcommon.HexToAddress("0x5217b367b1447b7aa79d8b1d432631d3e8a49b96"),
			Price:         TokenPrices[rhea.WAVAX],
			TokenPoolType: rhea.LockRelease,
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WAVAX},
	WrappedNative: rhea.WAVAX,
	Router:        gethcommon.HexToAddress("0xb515b3a88ad10148e6e127fc62ba1947d3d286e8"),
	Afn:           gethcommon.HexToAddress("0xd134c084ca8198cf93faef5f1e9318980e1a659f"),
	PriceRegistry: gethcommon.HexToAddress("0x70f2c49dc068870e844a608c3a04075c98b0550a"),
	Confirmations: 1,
	DeploySettings: rhea.ChainDeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,
		DeployedAtBlock:     19908550,
	},
}

var Beta_SepoliaToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x14c9d65b534e2f7bb860835883204d7463594fee"),
		OnRamp:       gethcommon.HexToAddress("0x35daa1c01b88e336b4ac2854dd371f502e051b89"),
		OffRamp:      gethcommon.HexToAddress("0xf801d51c6146e8c2afa33b09ecfb23f8db8752b8"),
		PingPongDapp: gethcommon.HexToAddress("0x099FdE93C4c2D0DdBe65e4beaE2b0B861DA85857"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    2915166,
		},
	},
}

var Beta_AvaxFujiToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x4ae1dabcbe0fe25e74a489c85d6d6aff07afc06c"),
		OnRamp:       gethcommon.HexToAddress("0x7cddcbe782047a8c4b0e6615e0467f55602fa4fe"),
		OffRamp:      gethcommon.HexToAddress("0x7cbdb9cf6420dd3fda3ce258e9933b998bbe99ac"),
		PingPongDapp: gethcommon.HexToAddress("0x8B9d4513813784b9925A7dC6b3Bb941a62f43c98"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    19908550,
		},
	},
}

var Beta_SepoliaToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0xecb44034f7cca8fb81ecaf6e3e374bc2174e6fb2"),
		OnRamp:       gethcommon.HexToAddress("0xf0dd5fde16c06c7ac6f152ba853306955f28dc0b"),
		OffRamp:      gethcommon.HexToAddress("0x8080b29b11d742c490bec91f2e55d652ec9d9383"),
		PingPongDapp: gethcommon.HexToAddress("0x48230efa6fba0b20cac752a8c3b95e17e7bfb408"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    3109217,
		},
	},
}

var Beta_OptimismGoerliToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x64e916ec137fa79d42acf450190da405678113b4"),
		OnRamp:       gethcommon.HexToAddress("0x59b9878b282cbd84bb14d977b54e13e686368dd0"),
		OffRamp:      gethcommon.HexToAddress("0x2ef9eda52af810b72385b2c05b18bdddddac8fac"),
		PingPongDapp: gethcommon.HexToAddress("0x84ddb8512655406a6dfa52ab46ba4bdad562d4fe"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    6816841,
		},
	},
}

var Beta_OptimismGoerliToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x938df67884ac3839abad96795108f3cd2b799f18"),
		OnRamp:       gethcommon.HexToAddress("0xe184affeb2eccabe0413a42287687ccef035fe25"),
		OffRamp:      gethcommon.HexToAddress("0xd6ced0a2392fba09c214d89701de11bf8bce7741"),
		PingPongDapp: gethcommon.HexToAddress("0x41550e33f7abfedd1380dc50c34ce97adcd0836d"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    6816841,
		},
	},
}

var Beta_AvaxFujiToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0xe8cbcdb79bc4ec90ea3262086938c0f7587cc613"),
		OnRamp:       gethcommon.HexToAddress("0x403285d4838eeb9be1eacac6994536cc751bcc86"),
		OffRamp:      gethcommon.HexToAddress("0x5135db05256a71e3776c794eab6e3141e9a37e27"),
		PingPongDapp: gethcommon.HexToAddress("0xac5c3c15a0d6e02508e44c99cf0570bf1ab253d0"),
		DeploySettings: rhea.LaneDeploySettings{
			DeployCommitStore:  false,
			DeployRamp:         false,
			DeployPingPongDapp: false,
			DeployedAtBlock:    19942877,
		},
	},
}
