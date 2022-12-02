package deployments

import (
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
)

var Alpha_AvaxFuji = rhea.EVMChainConfig{
	ChainId: big.NewInt(43113),
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	LinkToken: gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846"),
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token: gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846"),
			Pool:  gethcommon.HexToAddress("0xb6f1Fe2CDE891eFd5Efd2A563C4C2F2549163718"),
			Price: big.NewInt(1),
		},
		rhea.SNX: {
			Token: gethcommon.HexToAddress("0x3C3de1Dd82eA10B664C693C9a3c19645Ab9635EB"),
			Pool:  gethcommon.HexToAddress("0x43A2A4C2ECB74FF45Eca704a14111d8f2B1c0fA0"),
			Price: big.NewInt(1),
		},
	},
	Router: gethcommon.HexToAddress("0xc0A2c03115d1B48BAA59f676c108EfE5Ba3ee062"),
	Afn:    gethcommon.HexToAddress("0xb2958D1Bd07448865E555FeeFf32b58D254ffB4C"),
}

var Alpha_OptimismGoerli = rhea.EVMChainConfig{
	ChainId: big.NewInt(420),
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	LinkToken: gethcommon.HexToAddress("0xdc2CC710e42857672E7907CF474a69B63B93089f"),
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token: gethcommon.HexToAddress("0xdc2CC710e42857672E7907CF474a69B63B93089f"),
			Pool:  gethcommon.HexToAddress("0xE4aB69C077896252FAFBD49EFD26B5D171A32410"),
			Price: big.NewInt(1),
		},
		rhea.SNX: {
			Token: gethcommon.HexToAddress("0xfe628556155F681dd897e3FD029e5ED699a9248E"),
			Pool:  gethcommon.HexToAddress("0xc5CCb84C3d8eAD52C081dDB24e7Add615c0c9Daf"),
			Price: big.NewInt(1),
		},
	},
	Router: gethcommon.HexToAddress("0xE591bf0A0CF924A0674d7792db046B23CEbF5f34"),
	Afn:    gethcommon.HexToAddress("0x4c10d67E4B8e18a67A7606DEFDCe42CCc281D39B"),
}

var Staging_Alpha_OptimismGoerlitoAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Alpha_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:     gethcommon.HexToAddress("0xf9B7595D64a380fFa605A1d11BFf5cd629FB7189"),
		OnRamp:          gethcommon.HexToAddress("0x4a827De1b7bB0F56c8Cd046dc8eA72C9f412f22c"),
		TokenSender:     gethcommon.HexToAddress("0x0cA18254C9DFB652F0d6A3b6C88aBAc3793EDdf5"),
		OffRamp:         gethcommon.HexToAddress("0xbAcf5cb76B2AbC6b754bCffAe8209C76bAE731aA"),
		MessageReceiver: gethcommon.HexToAddress("0xeB59fefaFbE89EC09a546172eddE3300c9889B14"),
		ReceiverDapp:    gethcommon.HexToAddress("0x86000BFF3465C579dbA5703B2DBA6117ce022576"),
		GovernanceDapp:  gethcommon.HexToAddress(""),
		PingPongDapp:    gethcommon.HexToAddress("0xdf19B70440051A6497aB48B86E291746cdFeF89A"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployRouter:         false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           2297721,
	},
}

var Staging_Alpha_AvaxFujitoOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Alpha_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:     gethcommon.HexToAddress("0x84B7B012c95f8A152B44Ab3e952f2dEE424fA8e1"),
		OnRamp:          gethcommon.HexToAddress("0x65120aF1C7Ecaa90294758AafbB87226D2b3B798"),
		TokenSender:     gethcommon.HexToAddress("0x7854E73C73e7F9bb5b0D5B4861E997f4C6E8dcC6"),
		OffRamp:         gethcommon.HexToAddress("0x832c8f2666adBeA842ef30C90DeB59225Bcd67aa"),
		MessageReceiver: gethcommon.HexToAddress("0x75d642e8050d075C225ca3ED818C39ba7A6D6B76"),
		ReceiverDapp:    gethcommon.HexToAddress("0x6154b0a8Ada0Da450E4226bf8772b3A1B756A152"),
		GovernanceDapp:  gethcommon.HexToAddress(""),
		PingPongDapp:    gethcommon.HexToAddress("0x35a926bc94654627443e436Bb3D197D62821cF05"),
	},

	DeploySettings: rhea.DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployRouter:         false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           15036940,
	},
}
