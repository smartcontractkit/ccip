package deployments

import (
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
)

var Alpha_ChainConfigs = []rhea.EvmDeploymentConfig{
	{ChainConfig: Alpha_AvaxFuji},
	{ChainConfig: Alpha_OptimismGoerli},
	{ChainConfig: Alpha_Sepolia},
}

var Alpha_Sepolia = rhea.EVMChainConfig{
	ChainId: 11155111,
	GasSettings: rhea.EVMGasSettings{
		EIP1559:   true,
		GasTipCap: rhea.DefaultGasTipFee,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token: gethcommon.HexToAddress("0xb227f007804c16546Bd054dfED2E7A1fD5437678"),
			Pool:  gethcommon.HexToAddress("0xAE32FD8Ae148BD88E3da6FaE8Cd7561Eed3ec5Cc"),
			Price: big.NewInt(1),
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK},
	Router:        gethcommon.HexToAddress("0x7B9e7A97ca47B7a5501c8EB07aBF78Fb37bB4738"),
	Afn:           gethcommon.HexToAddress("0x521c7694c158d2d3D2FC79a93d2aDa5673d8226a"),
	PriceRegistry: gethcommon.HexToAddress("0x43a75aDEb32A06C5489d32a946dD8F2cC18Db9E8"),
	Confirmations: 4,
}

var Alpha_AvaxFuji = rhea.EVMChainConfig{
	ChainId: 43113,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token: gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846"),
			Pool:  gethcommon.HexToAddress("0xA92218886F3b4A8e5C59B959f0Bd05A3f7138F39"),
			Price: big.NewInt(1),
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK},
	Router:        gethcommon.HexToAddress("0x0e9210C8601723d9C546298db86D8925D5884eD1"),
	Afn:           gethcommon.HexToAddress("0xEbEB45ef02491dcF4400E26Dfe5b25f4fB1BCDf7"),
	PriceRegistry: gethcommon.HexToAddress("0x6F21947b14037f541a030ba9ca801986DC0ca9E9"),
	Confirmations: 1,
}

var Alpha_OptimismGoerli = rhea.EVMChainConfig{
	ChainId: 420,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: true,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token: gethcommon.HexToAddress("0xdc2CC710e42857672E7907CF474a69B63B93089f"),
			Pool:  gethcommon.HexToAddress("0x25c53f77e4f6FC85CbA2a892Ac62A44C770389cC"),
			Price: big.NewInt(1),
		},
	},
	FeeTokens:     []rhea.Token{rhea.LINK},
	Router:        gethcommon.HexToAddress("0x9CdA5b77eA23459eBaf2e3092c570a6B5605850A"),
	Afn:           gethcommon.HexToAddress("0xdAe257e1ACF6665eB897BFbf8Dd1bF62825E87dD"),
	PriceRegistry: gethcommon.HexToAddress("0x0D52c2472DC6f37FBeE59552b1165deafe3b9a4B"),
	Confirmations: 4,
}

var Alpha_OptimismGoerliToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Alpha_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x039ad8B76Efc29F1b233d04B413a0DCACBe7E901"),
		OnRamp:       gethcommon.HexToAddress("0xbdaA9d42aD7560F5BD673E55da776966C1E73ACa"),
		OffRamp:      gethcommon.HexToAddress("0x2C2AD296F773fd2f59b0e67AC3EdAd5f3a456008"),
		ReceiverDapp: gethcommon.HexToAddress("0x0D8d5E6b3d34cf7bF7093c18d976AEd86e699855"),
		PingPongDapp: gethcommon.HexToAddress("0x6E97B1DC16023D0f7C3F60342cAD3d30C3BC9971"),
	},

	DeploySettings: rhea.DeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,

		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           4103083,
	},
}

var Alpha_AvaxFujiToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Alpha_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0xaD427c949EA402E70c681F2723F843D6e7968DB7"),
		OnRamp:       gethcommon.HexToAddress("0x45fEef0276358C869f1e5ad152FB664914830d21"),
		OffRamp:      gethcommon.HexToAddress("0xB84b195542c17a7b6bAc25289Fc6755Dd7908b70"),
		ReceiverDapp: gethcommon.HexToAddress("0x2c8433DDa4B2e492Be07c90e83849db2a6f5A89e"),
		PingPongDapp: gethcommon.HexToAddress("0x93eE84a98e7454a6CEfEdAb3d3Def3f8daedc9Ae"),
	},

	DeploySettings: rhea.DeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,

		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           17987033,
	},
}

var Alpha_SepoliaToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Alpha_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x0f30449bcCaCCaA7221B3f7C3304c4AaD68068E8"),
		OnRamp:       gethcommon.HexToAddress("0xb7CE1b61D30C68776A94eFf0AF92aF9b77FF6521"),
		OffRamp:      gethcommon.HexToAddress("0x848683AaF65d62Cd326BA6e49F2a6417F7f6EEA7"),
		ReceiverDapp: gethcommon.HexToAddress("0x4bFd0226BEDed9Ff665bD07e42524905503F7b34"),
		PingPongDapp: gethcommon.HexToAddress("0x9CdA5b77eA23459eBaf2e3092c570a6B5605850A"),
	},

	DeploySettings: rhea.DeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,

		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           2651082,
	},
}

var Alpha_AvaxFujiToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Alpha_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0xB39FC12e7eDa9af06c3c2E8406C9d590Cf3C850B"),
		OnRamp:       gethcommon.HexToAddress("0x16B05c98D1e9EAD51BccfCbFD514e48edBBD8DF9"),
		OffRamp:      gethcommon.HexToAddress("0xF06230a439b21794569D61C86d3824bC724295fC"),
		ReceiverDapp: gethcommon.HexToAddress("0x80b5c0B162F57AAaceD888CE898068d621756ee6"),
		PingPongDapp: gethcommon.HexToAddress("0x4abaadc0cef614149e0722060babcfb99051f0d6"),
	},

	DeploySettings: rhea.DeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,

		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           17835508,
	},
}

var Alpha_OptimismGoerliToSepolia = rhea.EvmDeploymentConfig{
	ChainConfig: Alpha_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x58aE643a115c0A11f7A2115FA7f396ff1cC49234"),
		OnRamp:       gethcommon.HexToAddress("0x5058Af17E36899Aa9073c2FE777F3E79ae06c566"),
		OffRamp:      gethcommon.HexToAddress("0xcC22cd1AD8aE43389eB3577AE576EFB99E66BE25"),
		ReceiverDapp: gethcommon.HexToAddress("0x99cE75105D6A882Af40CD5F6166A9564b3003a07"),
		PingPongDapp: gethcommon.HexToAddress("0xc136114F379b812345bb7e467ECDdb6D0c87De8b"),
	},

	DeploySettings: rhea.DeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,

		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           3996202,
	},
}

var Alpha_SepoliaToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Alpha_Sepolia,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x6E97B1DC16023D0f7C3F60342cAD3d30C3BC9971"),
		OnRamp:       gethcommon.HexToAddress("0xec30e7025731995eA53B2A7d3AA3F72D92D61cc7"),
		OffRamp:      gethcommon.HexToAddress("0xc888A46C756Dca1F37a01DD246C496ba2386Eeb6"),
		ReceiverDapp: gethcommon.HexToAddress("0x12c164d0778E215873A062cEE2814507417339cB"),
		PingPongDapp: gethcommon.HexToAddress("0xB3F3f362FbeD49fA0086B434051C822B55BaADbD"),
	},

	DeploySettings: rhea.DeploySettings{
		DeployAFN:           false,
		DeployTokenPools:    false,
		DeployRouter:        false,
		DeployPriceRegistry: false,

		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           2651606,
	},
}
