package deployments

import (
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
)

// Chains

var Beta_ChainConfigs = []rhea.EvmDeploymentConfig{
	{ChainConfig: Beta_Goerli},
	{ChainConfig: Beta_AvaxFuji},
}

var Beta_Goerli = rhea.EVMChainConfig{
	ChainId: 5,
	GasSettings: rhea.EVMGasSettings{
		EIP1559:   true,
		GasTipCap: rhea.DefaultGasTipFee,
	},
	LinkToken: gethcommon.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB"),
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token: gethcommon.HexToAddress("0x326c977e6efc84e512bb9c30f76e30c160ed06fb"),
			Pool:  gethcommon.HexToAddress("0xa579e20583a02b389e2354f359f6a22d7a8a72c3"),
			Price: big.NewInt(1),
		},
		rhea.Custom: {
			Token: gethcommon.HexToAddress("0x419a4c8c9ba74bd1fdfb355f7b02848f758dd9ce"),
			Pool:  gethcommon.HexToAddress("0xe77fc010fe72098c396846e55ddccedcce4322df"),
			Price: big.NewInt(1),
		},
	},
	Router:     gethcommon.HexToAddress("0x14d66e53299174b8ef117d33285e1c814ed652de"),
	Afn:        gethcommon.HexToAddress("0xb1cc4adbbfa4216541d9ccc7ba2240d129986efb"),
	FeeManager: gethcommon.HexToAddress("0x8aed531563568dde9ff437d86c65b637535a5da3"),
}

var Beta_AvaxFuji = rhea.EVMChainConfig{
	ChainId: 43113,
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	SupportedTokens: map[rhea.Token]rhea.EVMBridgedToken{
		rhea.LINK: {
			Token: gethcommon.HexToAddress("0x0b9d5d9136855f6fec3c0993fee6e9ce8a297846"),
			Pool:  gethcommon.HexToAddress("0xbbf534d89d9640e3886db25fe1ffe603fe160d75"),
			Price: big.NewInt(1),
		},
		rhea.Custom: {
			Token: gethcommon.HexToAddress("0x1d22f1dd850980d738a4dbd71588f07eeca10dfe"),
			Pool:  gethcommon.HexToAddress("0xb7ce1b61d30c68776a94eff0af92af9b77ff6521"),
			Price: big.NewInt(1),
		},
	},
	LinkToken:  gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846"),
	Router:     gethcommon.HexToAddress("0x9213967a47fc3f15a16a0b813208e8ccb63dbba6"),
	Afn:        gethcommon.HexToAddress("0xb5fb34580ff11fd2ab6cb4b1182aa2cd589b9234"),
	FeeManager: gethcommon.HexToAddress("0x6f5b9393032cb06dabb09d8b08b7b08e045413fa"),
}

// Lanes

var Beta_GoerliToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_Goerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:          gethcommon.HexToAddress("0x433165f192f7665e8f0cd7a35f79330fd43da0d6"),
		OffRamp:         gethcommon.HexToAddress("0x2efd2a210f8e8f14067ce2017cd0b71664895011"),
		CommitStore:     gethcommon.HexToAddress("0xded799abdb9fdeee0fa738bbcef565996a169a0e"),
		TokenSender:     gethcommon.HexToAddress("0x0f30449bccaccaa7221b3f7c3304c4aad68068e8"),
		MessageReceiver: gethcommon.HexToAddress("0x848683aaf65d62cd326ba6e49f2a6417f7f6eea7"),
		ReceiverDapp:    gethcommon.HexToAddress("0x9eb17D0aB50A79f40654d090c7Db276BFD12A7C2"),
		GovernanceDapp:  gethcommon.HexToAddress("0x688b32d38770a1690b1a063339ace12960d8ad35"),
		PingPongDapp:    gethcommon.HexToAddress("0xA3547d2a5d39eDd8B704D38d49a58cF2fdd7B0B7"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployRouter:         false,
		DeployGovernanceDapp: false,
		DeployFeeManager:     false,
		DeployPingPongDapp:   false,
		DeployedAt:           8092017,
	},
}

var Beta_AvaxFujiToGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:          gethcommon.HexToAddress("0x17a5746c9cf7eaf23533f060f395b2e38eb976ea"),
		OffRamp:         gethcommon.HexToAddress("0x439db5c0c194b07f7b5e6e7db8e39c264036a1df"),
		CommitStore:     gethcommon.HexToAddress("0x82d96373fb24ce812b051db4b53e490a20cfbbff"),
		TokenSender:     gethcommon.HexToAddress("0x99ce75105d6a882af40cd5f6166a9564b3003a07"),
		MessageReceiver: gethcommon.HexToAddress("0x3b80b7ef5c00eb892cbe72800c028c47ad6380ef"),
		ReceiverDapp:    gethcommon.HexToAddress("0x41aEe73f38e58618d6D7901Fa28cF1c9899c076d"),
		GovernanceDapp:  gethcommon.HexToAddress("0x0d52c2472dc6f37fbee59552b1165deafe3b9a4b"),
		PingPongDapp:    gethcommon.HexToAddress("0x155fd0E4852fE4709022d530Fc86F31a183A790B"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployRouter:         false,
		DeployGovernanceDapp: false,
		DeployFeeManager:     false,
		DeployPingPongDapp:   false,
		DeployedAt:           16658457,
	},
}
