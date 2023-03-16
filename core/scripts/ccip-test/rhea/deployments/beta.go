package deployments

import (
	"math/big"

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
			Token: gethcommon.HexToAddress("0x779877A7B0D9E8603169DdbD7836e478b4624789"),
			Pool:  gethcommon.HexToAddress("0xb210d4a3e4d7fb5089d5332d62128e747f081640"),
			Price: big.NewInt(10),
		},
		rhea.WETH: {
			Token: gethcommon.HexToAddress("0x097D90c9d3E0B50Ca60e1ae45F6A81010f9FB534"),
			Pool:  gethcommon.HexToAddress("0xa45dc3abb8924e82d7cc74d723f6e35d0afe172f"),
			Price: big.NewInt(1500),
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	Router:        gethcommon.HexToAddress("0xc4f3c3bb9e58ab406450cc1704f20f94e60b02f3"),
	Afn:           gethcommon.HexToAddress("0x280fd789c4df1ff45b503a71a5700d94049798d6"),
	PriceRegistry: gethcommon.HexToAddress("0x8bf3b19e46489f2efc7fd0c52123e41e4e47d7b3"),
	Confirmations: 4,
}

var Beta_OptimismGoerli = rhea.EVMChainConfig{
	ChainId: 420,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: true,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token: gethcommon.HexToAddress("0xdc2CC710e42857672E7907CF474a69B63B93089f"),
			Pool:  gethcommon.HexToAddress("0xd2b7874dc83ba90ddedd3b5bbd8bf69d70f3e08d"),
			Price: big.NewInt(10),
		},
		rhea.WETH: {
			Token: gethcommon.HexToAddress("0x4200000000000000000000000000000000000006"),
			Pool:  gethcommon.HexToAddress("0x0283f4786e6a1e1748f3b1215768f59eb9f3b19b"),
			Price: big.NewInt(1500),
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WETH},
	Router:        gethcommon.HexToAddress("0xebeb45ef02491dcf4400e26dfe5b25f4fb1bcdf7"),
	Afn:           gethcommon.HexToAddress("0x4f6fc77a5cb17e289dda758dc932f4c23de3e230"),
	PriceRegistry: gethcommon.HexToAddress("0xa92218886f3b4a8e5c59b959f0bd05a3f7138f39"),
	Confirmations: 4,
}

var Beta_AvaxFuji = rhea.EVMChainConfig{
	ChainId: 43113,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token: gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846"),
			Pool:  gethcommon.HexToAddress("0x0529305ea53834a96a624dc04f874051e8bb9bc3"),
			Price: big.NewInt(10),
		},
		rhea.WAVAX: {
			Token: gethcommon.HexToAddress("0xd00ae08403B9bbb9124bB305C09058E32C39A48c"),
			Pool:  gethcommon.HexToAddress("0x0ad506743e69cd070343916b40464b5e6d665a6f"),
			Price: big.NewInt(25),
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK, rhea.WAVAX},
	Router:        gethcommon.HexToAddress("0xb515b3a88ad10148e6e127fc62ba1947d3d286e8"),
	Afn:           gethcommon.HexToAddress("0x828163926583e32f2d5570a8c4d338ee73c72659"),
	PriceRegistry: gethcommon.HexToAddress("0x543ef9e167e5e4d0c8bb41f2154e65b7b6f369a0"),
	Confirmations: 1,
}

var Beta_SepoliaToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0xeb2053029d5dcb0156b2a4c2c6d1feb7b2db7f02"),
		OnRamp:       gethcommon.HexToAddress("0x77455df22d16277622dbde3bf8b2aa26418f10ed"),
		OffRamp:      gethcommon.HexToAddress("0x6ff28baea457664e741d9313267e1bc6de3ab276"),
		ReceiverDapp: gethcommon.HexToAddress("0xd0807c5fd955ed4c160776f4238438b7830a39b4"),
		PingPongDapp: gethcommon.HexToAddress("0x4bcA742cE3c7B8fC2419E25e14EC97D3470734bF"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,

		DeployCommitStore:  false,
		DeployRamp:         false,
		DeployPingPongDapp: false,
		DeployedAt:         2915166,
	},
}

var Beta_AvaxFujiToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x363c3a63ab17affcfbb4ed88d08bde29672ef59b"),
		OnRamp:       gethcommon.HexToAddress("0xc468a11959abafaf3a2872fbf75b6bfe815aa97b"),
		OffRamp:      gethcommon.HexToAddress("0x15e8c905907bf6342e9b4967abaf451e888df764"),
		ReceiverDapp: gethcommon.HexToAddress("0x54be0408e91fab34cd01ba105e6bbc3232d85286"),
		PingPongDapp: gethcommon.HexToAddress("0x7cab0Ec4fB08746d2acBd6D02D334d8f777be670"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,

		DeployCommitStore:  false,
		DeployRamp:         false,
		DeployPingPongDapp: false,
		DeployedAt:         18998250,
	},
}

var Beta_SepoliaToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x6ce4048828bfdc48ee496783895ec0e9e2289a21"),
		OnRamp:       gethcommon.HexToAddress("0x6be5a3a7d0b0133a6d98fb83e5b16d4846ba9ad5"),
		OffRamp:      gethcommon.HexToAddress("0x278979cdebaa46f647f1ccbf13c8c432b55b010e"),
		ReceiverDapp: gethcommon.HexToAddress("0xb6af49084ef002010a42cfb048c93692fc1b9202"),
		PingPongDapp: gethcommon.HexToAddress("0x08e39Fd9D3D067e2929F30B4c30A631F1b0a0aCf"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,

		DeployCommitStore:  false,
		DeployRamp:         false,
		DeployPingPongDapp: false,
		DeployedAt:         2915382,
	},
}

var Beta_OptimismGoerliToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0xc131d20fd2143770689ed72d2cd0332b99c0f2bb"),
		OnRamp:       gethcommon.HexToAddress("0x8e14e3ed8610b1f85bf927a5010ac5ea6e95b214"),
		OffRamp:      gethcommon.HexToAddress("0x74f461b758037893f7ba17e889b2be0336f8c987"),
		ReceiverDapp: gethcommon.HexToAddress("0xc8e6f4225f3821f1b19fccf6e0d78ece42196400"),
		PingPongDapp: gethcommon.HexToAddress("0xa45dC3Abb8924e82d7cC74D723F6E35d0AFE172f"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,

		DeployCommitStore:  false,
		DeployRamp:         false,
		DeployPingPongDapp: false,
		DeployedAt:         5578510,
	},
}

var Beta_OptimismGoerliToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0xe940ecbfff465dc8cbf48f55894bfdee46453031"),
		OnRamp:       gethcommon.HexToAddress("0x2c29abe892a0258b7e4605c06b0a3d78493b15f7"),
		OffRamp:      gethcommon.HexToAddress("0x2250b6fb8aefd36079924709925fef6f755d2521"),
		ReceiverDapp: gethcommon.HexToAddress("0x8b49b6ef445ef7460e82b42804e37b6ed6d01020"),
		PingPongDapp: gethcommon.HexToAddress("0x139FF64C62e3825f79E177C18200e8b456D7553D"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,

		DeployCommitStore:  false,
		DeployRamp:         false,
		DeployPingPongDapp: false,
		DeployedAt:         5577852,
	},
}

var Beta_AvaxFujiToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0xc07699e59ef98ae0aa09a9d2cbeb80192b992c5c"),
		OnRamp:       gethcommon.HexToAddress("0x1fd73af7953b3402b9204e01f137f9878b04122e"),
		OffRamp:      gethcommon.HexToAddress("0x3a9eb812a299196f2b4e05c62cbeb959c7beab59"),
		ReceiverDapp: gethcommon.HexToAddress("0x7f682178faf7fa2feb80ff6845aa8203fe63b903"),
		PingPongDapp: gethcommon.HexToAddress("0xb7eeD42a7A27B861dd009fc8d5B45392Fd3cB13A"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,

		DeployCommitStore:  false,
		DeployRamp:         false,
		DeployPingPongDapp: false,
		DeployedAt:         18998581,
	},
}
