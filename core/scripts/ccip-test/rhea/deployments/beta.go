package deployments

import (
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
)

// Chains

var Beta_Goerli = rhea.EVMChainConfig{
	ChainId: big.NewInt(5),
	GasSettings: rhea.EVMGasSettings{
		EIP1559:   true,
		GasTipCap: rhea.DefaultGasTipFee,
	},
	LinkToken: gethcommon.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB"),
	SupportedTokens: map[gethcommon.Address]rhea.EVMBridgedToken{
		gethcommon.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB"): {
			Pool:  gethcommon.HexToAddress("0x8742E2aC5a5f9c8aAC465Bd8b6Ce1BD54F4d85a4"),
			Price: big.NewInt(1),
		},
		gethcommon.HexToAddress("0x419A4c8C9bA74Bd1fdfb355f7b02848f758DD9Ce"): {
			Pool:  gethcommon.HexToAddress("0xea50De1CA43f136aF11fb06F2393E47E68B47A8E"),
			Price: big.NewInt(1),
		},
	},
	OnRampRouter:  gethcommon.HexToAddress("0xdeE922E69c5142653281163590c92e459bE5Ace8"),
	OffRampRouter: gethcommon.HexToAddress("0xc690CbC8999c2A464606db88657a883a7De26Daf"),
	Afn:           gethcommon.HexToAddress("0x98165Aff4546c25a74e2E86D4e497Aa5b36034BE"),
}

var Beta_AvaxFuji = rhea.EVMChainConfig{
	ChainId: big.NewInt(43113),
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	SupportedTokens: map[gethcommon.Address]rhea.EVMBridgedToken{
		gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846"): {
			Pool:  gethcommon.HexToAddress("0x866faB92E04bAE5EDa238A9cbFf1e56E09508Ade"),
			Price: big.NewInt(1),
		},
		gethcommon.HexToAddress("0x1D22f1dd850980D738A4dBD71588f07eECa10dfE"): {
			Pool:  gethcommon.HexToAddress("0x0924011a856483E47565d54BBC65cA9E21E8EE42"),
			Price: big.NewInt(1),
		},
	},
	LinkToken:     gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846"),
	OnRampRouter:  gethcommon.HexToAddress("0x338F33f149C9257284a37144E37B1D5A62507a0E"),
	OffRampRouter: gethcommon.HexToAddress("0x5058Af17E36899Aa9073c2FE777F3E79ae06c566"),
	Afn:           gethcommon.HexToAddress("0xD886E2286Fd1073df82462ea1822119600Af80b6"),
}

// Lanes

var Beta_GoerliToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_Goerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:          gethcommon.HexToAddress("0x9213967a47FC3F15A16A0b813208e8Ccb63Dbba6"),
		OffRamp:         gethcommon.HexToAddress("0x82D96373fB24Ce812B051db4B53E490a20CFbBfF"),
		CommitStore:     gethcommon.HexToAddress("0x33BBb9c3Ee0f80F1777E973D3814f52740019A86"),
		TokenSender:     gethcommon.HexToAddress("0x0f30449bcCaCCaA7221B3f7C3304c4AaD68068E8"),
		MessageReceiver: gethcommon.HexToAddress("0x848683AaF65d62Cd326BA6e49F2a6417F7f6EEA7"),
		ReceiverDapp:    gethcommon.HexToAddress("0x95074a4903719940516ED44aE101dbd1BFe101d2"),
		GovernanceDapp:  gethcommon.HexToAddress(""),
		PingPongDapp:    gethcommon.HexToAddress("0x276e6b380C7123e494973416eA48480ec489c445"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployRouter:         false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           7786108,
	},
}

var Beta_AvaxFujiToGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Beta_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:          gethcommon.HexToAddress("0x05EBCE331e0201BaeA852A7c3c6f7e51A893D3F0"),
		OffRamp:         gethcommon.HexToAddress("0x11BEe8AD23bA3Fd56fcbD88467D5C76375fD03ef"),
		CommitStore:     gethcommon.HexToAddress("0x0fF6b6F3Ad10D66600Fd5CC25b98542A05Aa7Bc2"),
		TokenSender:     gethcommon.HexToAddress("0x99cE75105D6A882Af40CD5F6166A9564b3003a07"),
		MessageReceiver: gethcommon.HexToAddress("0x3B80b7Ef5c00Eb892CBe72800C028C47AD6380EF"),
		ReceiverDapp:    gethcommon.HexToAddress("0x25d997d8618e1299418b3D905E40bC353ec89F61"),
		GovernanceDapp:  gethcommon.HexToAddress(""),
		PingPongDapp:    gethcommon.HexToAddress("0x01325475d63d77040968B341590c67f87daA82De"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployRouter:         false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           14710571,
	},
}
