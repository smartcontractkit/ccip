package deployments

import (
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
)

var Alpha_ChainConfigs = []rhea.EvmDeploymentConfig{
	{ChainConfig: Alpha_AvaxFuji},
	{ChainConfig: Alpha_OptimismGoerli},
}

var Alpha_AvaxFuji = rhea.EVMChainConfig{
	ChainId: 43113,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	LinkToken: gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846"),
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token: gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846"),
			Pool:  gethcommon.HexToAddress("0x9af4a72eab994338cae5deb1f2cdaf49e3c4f11e"),
			Price: big.NewInt(1),
		},
		rhea.SNX: {
			Token: gethcommon.HexToAddress("0x3C3de1Dd82eA10B664C693C9a3c19645Ab9635EB"),
			Pool:  gethcommon.HexToAddress("0xd2c0b212f2366ca5f02fcd91407d64dedca58422"),
			Price: big.NewInt(1),
		},
	},
	Router:      gethcommon.HexToAddress("0x6486906bb2d85a6c0ccef2a2831c11a2059ebfea"),
	Afn:         gethcommon.HexToAddress("0x1f350718e015eb20e5065c09f4a7a3f66888aeed"),
	GasFeeCache: gethcommon.HexToAddress("0xfa41f1986e0443e6e2780ee9fdcf09697a2e3517"),
}

var Alpha_OptimismGoerli = rhea.EVMChainConfig{
	ChainId: 420,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	LinkToken: gethcommon.HexToAddress("0xdc2CC710e42857672E7907CF474a69B63B93089f"),
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token: gethcommon.HexToAddress("0xdc2CC710e42857672E7907CF474a69B63B93089f"),
			Pool:  gethcommon.HexToAddress("0x3cc9364260d80f09ccac1ee6b07366db598900e6"),
			Price: big.NewInt(1),
		},
		rhea.SNX: {
			Token: gethcommon.HexToAddress("0xfe628556155F681dd897e3FD029e5ED699a9248E"),
			Pool:  gethcommon.HexToAddress("0x3382b044d5a3ff656ffb62daaced78084c209e71"),
			Price: big.NewInt(1),
		},
	},
	Router:      gethcommon.HexToAddress("0x114a20a10b43d4115e5aeef7345a1a71d2a60c57"),
	Afn:         gethcommon.HexToAddress("0x8af4204e30565df93352fe8e1de78925f6664da7"),
	GasFeeCache: gethcommon.HexToAddress("0xe25db981c1bc20fb3ed93774d095bb9aa4792234"),
}

var Staging_Alpha_OptimismGoerlitoAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Alpha_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:     gethcommon.HexToAddress("0xb50e48d32bb3ae0afbf97630d7e495e1a2d2ebe3"),
		OnRamp:          gethcommon.HexToAddress("0x260af9b83e0d2bb6c9015fc9f0bff8858a0cce68"),
		TokenSender:     gethcommon.HexToAddress(""),
		OffRamp:         gethcommon.HexToAddress("0xd3b06cebf099ce7da4accf578aaebfdbd6e88a93"),
		MessageReceiver: gethcommon.HexToAddress(""),
		ReceiverDapp:    gethcommon.HexToAddress("0xaee2dac4cecceb4c51f1191b9481c87b27c283dc"),
		GovernanceDapp:  gethcommon.HexToAddress("0x6b38cc6fa938d5ab09bdf0cfe580e226fdd793ce"),
		PingPongDapp:    gethcommon.HexToAddress("0xd2C0b212F2366CA5f02FCd91407D64DedcA58422"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:         false,
		DeployTokenPools:  false,
		DeployRouter:      false,
		DeployGasFeeCache: false,

		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           3258713,
	},
}

var Staging_Alpha_AvaxFujitoOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Alpha_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:     gethcommon.HexToAddress("0x552a6D7131F3425e48f313003aA2fd08596d9663"),
		OnRamp:          gethcommon.HexToAddress("0xfa95612b6d1d98c983ef3880ea9ee53fa8cde48c"),
		TokenSender:     gethcommon.HexToAddress(""),
		OffRamp:         gethcommon.HexToAddress("0xf8ba25c2bf50233602b280f180a21d99b89f173c"),
		MessageReceiver: gethcommon.HexToAddress(""),
		ReceiverDapp:    gethcommon.HexToAddress("0x40cac95c005572c2fb43835334d1f3f020ed6dc7"),
		GovernanceDapp:  gethcommon.HexToAddress("0x6eb91dc6d5263d2551e3360abfaabbe08ac29be1"),
		PingPongDapp:    gethcommon.HexToAddress("0x37687BB2a0582Ddf05C1a289876835ca9c44114b"),
	},

	DeploySettings: rhea.DeploySettings{
		DeployAFN:         false,
		DeployTokenPools:  false,
		DeployRouter:      false,
		DeployGasFeeCache: false,

		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           16656664,
	},
}
