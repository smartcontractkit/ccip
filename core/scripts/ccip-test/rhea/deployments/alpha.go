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
	LinkToken: gethcommon.HexToAddress("0xb227f007804c16546Bd054dfED2E7A1fD5437678"),
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token: gethcommon.HexToAddress("0xb227f007804c16546Bd054dfED2E7A1fD5437678"),
			Pool:  gethcommon.HexToAddress("0xAE32FD8Ae148BD88E3da6FaE8Cd7561Eed3ec5Cc"),
			Price: big.NewInt(1),
		},
	},
	Router:     gethcommon.HexToAddress("0x7B9e7A97ca47B7a5501c8EB07aBF78Fb37bB4738"),
	Afn:        gethcommon.HexToAddress("0x521c7694c158d2d3D2FC79a93d2aDa5673d8226a"),
	FeeManager: gethcommon.HexToAddress("0x43a75aDEb32A06C5489d32a946dD8F2cC18Db9E8"),
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
			Pool:  gethcommon.HexToAddress("0x1E77120aBD6Eeb6c2801AbC3e6799e3cbc8f67bd"),
			Price: big.NewInt(1),
		},
	},
	Router:     gethcommon.HexToAddress("0x64A16E59AECF0d85843B85e2ecca217fCef1f2eA"),
	Afn:        gethcommon.HexToAddress("0xa70fe070090C4748924B0E5368388059270b73ce"),
	FeeManager: gethcommon.HexToAddress("0xF037523a2a9BEAFAD74e77446DEB75d4635853F6"),
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
			Pool:  gethcommon.HexToAddress("0xD886E2286Fd1073df82462ea1822119600Af80b6"),
			Price: big.NewInt(1),
		},
	},
	Router:     gethcommon.HexToAddress("0x11BEe8AD23bA3Fd56fcbD88467D5C76375fD03ef"),
	Afn:        gethcommon.HexToAddress("0x2aE6d5495fc20226F433be50e37D59c05D186AaA"),
	FeeManager: gethcommon.HexToAddress("0xD84E2f21D0bBbe0be697cB4A400AF367b39FbCD4"),
}

var Alpha_OptimismGoerliToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Alpha_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0x8742E2aC5a5f9c8aAC465Bd8b6Ce1BD54F4d85a4"),
		OnRamp:       gethcommon.HexToAddress("0x7B9e7A97ca47B7a5501c8EB07aBF78Fb37bB4738"),
		OffRamp:      gethcommon.HexToAddress("0x223F9cC58370139199B525e8E51B18129B648466"),
		ReceiverDapp: gethcommon.HexToAddress("0x17a5746c9cf7eAf23533F060F395B2E38eb976ea"),
		PingPongDapp: gethcommon.HexToAddress("0x33BBb9c3Ee0f80F1777E973D3814f52740019A86"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:        false,
		DeployTokenPools: false,
		DeployRouter:     false,
		DeployFeeManager: false,

		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           3997869,
	},
}

var Alpha_AvaxFujiToOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Alpha_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		CommitStore:  gethcommon.HexToAddress("0xFe6B3cE5beC09f58a24A097Fc017228df4888C39"),
		OnRamp:       gethcommon.HexToAddress("0x07603aaCfe7a5f3764f6E2B528507D6AE86eD472"),
		OffRamp:      gethcommon.HexToAddress("0x89FE16bc1555b6baF36836bF21e4BB6bC16CDeDC"),
		ReceiverDapp: gethcommon.HexToAddress("0x178Ed8d99D49eb870118EFefa889d4a2D6d7A5Bc"),
		PingPongDapp: gethcommon.HexToAddress("0xd2B7874dC83BA90DDEDD3B5bbD8bf69d70f3e08D"),
	},

	DeploySettings: rhea.DeploySettings{
		DeployAFN:        false,
		DeployTokenPools: false,
		DeployRouter:     false,
		DeployFeeManager: false,

		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           17842056,
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
		DeployAFN:        false,
		DeployTokenPools: false,
		DeployRouter:     false,
		DeployFeeManager: false,

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
		DeployAFN:        false,
		DeployTokenPools: false,
		DeployRouter:     false,
		DeployFeeManager: false,

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
		DeployAFN:        false,
		DeployTokenPools: false,
		DeployRouter:     false,
		DeployFeeManager: false,

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
		DeployAFN:        false,
		DeployTokenPools: false,
		DeployRouter:     false,
		DeployFeeManager: false,

		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           2651606,
	},
}
