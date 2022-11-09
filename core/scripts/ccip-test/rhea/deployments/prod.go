package deployments

import (
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
)

// Chains

var Prod_Goerli = rhea.EVMChainConfig{
	ChainId: big.NewInt(5),
	GasSettings: rhea.EVMGasSettings{
		EIP1559:   true,
		GasTipCap: rhea.DefaultGasTipFee,
	},
	LinkToken: gethcommon.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB"),
	SupportedTokens: map[gethcommon.Address]rhea.EVMBridgedToken{
		gethcommon.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB"): {
			Pool:  gethcommon.HexToAddress("0x4c10d67E4B8e18a67A7606DEFDCe42CCc281D39B"),
			Price: big.NewInt(1),
		},
		gethcommon.HexToAddress("0x5680dC17bD191EE04d048719b57983335c5E6153"): {
			Pool:  gethcommon.HexToAddress("0x1fce171011B16F3b0D16198e3F59FD72c091f43B"),
			Price: big.NewInt(1),
		},
	},
	OnRampRouter:  gethcommon.HexToAddress("0xA189971a2c5AcA0DFC5Ee7a2C44a2Ae27b3CF389"),
	OffRampRouter: gethcommon.HexToAddress("0xb78d314d32EB4B01C459EDE0774cc3b6AF244Dd7"),
	Afn:           gethcommon.HexToAddress("0x8a710bBd77661D168D5A6725bD2E514ba1bFf59d"),
}

var Prod_OptimismGoerli = rhea.EVMChainConfig{
	ChainId: big.NewInt(420),
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	LinkToken: gethcommon.HexToAddress("0xdc2CC710e42857672E7907CF474a69B63B93089f"),
	SupportedTokens: map[gethcommon.Address]rhea.EVMBridgedToken{
		gethcommon.HexToAddress("0xdc2CC710e42857672E7907CF474a69B63B93089f"): {
			Pool:  gethcommon.HexToAddress("0xE4aB69C077896252FAFBD49EFD26B5D171A32410"),
			Price: big.NewInt(1),
		},
		gethcommon.HexToAddress("0xfe628556155F681dd897e3FD029e5ED699a9248E"): {
			Pool:  gethcommon.HexToAddress("0xc5CCb84C3d8eAD52C081dDB24e7Add615c0c9Daf"),
			Price: big.NewInt(1),
		},
	},
	OnRampRouter:  gethcommon.HexToAddress("0xE591bf0A0CF924A0674d7792db046B23CEbF5f34"),
	OffRampRouter: gethcommon.HexToAddress("0x2b7aB40413DA5077E168546eA376920591Aee8E7"),
	Afn:           gethcommon.HexToAddress("0x4c10d67E4B8e18a67A7606DEFDCe42CCc281D39B"),
}

var Prod_AvaxFuji = rhea.EVMChainConfig{
	ChainId: big.NewInt(43113),
	GasSettings: rhea.EVMGasSettings{
		EIP1559: false,
	},
	LinkToken: gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846"),
	SupportedTokens: map[gethcommon.Address]rhea.EVMBridgedToken{
		gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846"): {
			Pool:  gethcommon.HexToAddress("0xb6f1Fe2CDE891eFd5Efd2A563C4C2F2549163718"),
			Price: big.NewInt(1),
		},
		gethcommon.HexToAddress("0x3C3de1Dd82eA10B664C693C9a3c19645Ab9635EB"): {
			Pool:  gethcommon.HexToAddress("0x43A2A4C2ECB74FF45Eca704a14111d8f2B1c0fA0"),
			Price: big.NewInt(1),
		},
	},
	OnRampRouter:  gethcommon.HexToAddress("0xc0A2c03115d1B48BAA59f676c108EfE5Ba3ee062"),
	OffRampRouter: gethcommon.HexToAddress("0x7d5297c5506ee2A7Ef121Da9bE02b6a6AD30b392"),
	Afn:           gethcommon.HexToAddress("0xb2958D1Bd07448865E555FeeFf32b58D254ffB4C"),
}

// Lanes

var Prod_GoerliToOptimism = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Goerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:          gethcommon.HexToAddress("0xD3b06cEbF099CE7DA4AcCf578aaebFDBd6e88a93"),
		OffRamp:         gethcommon.HexToAddress("0x07BFa2C37050d35825804a95cB698b80c7528c54"),
		CommitStore:     gethcommon.HexToAddress("0x701Fe16916dd21EFE2f535CA59611D818B017877"),
		TokenSender:     gethcommon.HexToAddress("0xc3e8bB61e1db9adE45F76237d75AAfaCca2066AF"),
		MessageReceiver: gethcommon.HexToAddress("0xe0D4860bD0429B87f508f0aE8d1789cC0adbbfcA"),
		ReceiverDapp:    gethcommon.HexToAddress("0x84B7B012c95f8A152B44Ab3e952f2dEE424fA8e1"),
		GovernanceDapp:  gethcommon.HexToAddress(""),
		PingPongDapp:    gethcommon.HexToAddress("0x201D1843707764CA2F236bd69E37CCbefF0827D4"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployRouter:         false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           7910091,
	},
}

var Prod_OptimismToGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:          gethcommon.HexToAddress("0x95B2D6e4119cC7d1832A7190b0126D034B642c6C"),
		OffRamp:         gethcommon.HexToAddress("0xFc47e4a7153312D30B9dA9706DdF3EeaB8324b3B"),
		CommitStore:     gethcommon.HexToAddress("0x4A1d9c5a7f9f9de7D5d8eC0f96f7213b0AB953d9"),
		TokenSender:     gethcommon.HexToAddress("0x51298c07eF8849f89552C2B3184741a759d4B37C"),
		MessageReceiver: gethcommon.HexToAddress("0x2321F13659889c2f1e7a62A7700744E36F9C60E5"),
		ReceiverDapp:    gethcommon.HexToAddress("0xA189971a2c5AcA0DFC5Ee7a2C44a2Ae27b3CF389"),
		GovernanceDapp:  gethcommon.HexToAddress(""),
		PingPongDapp:    gethcommon.HexToAddress("0xb6E24bd5376f808a8f4cEf945c96ec5582791255"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployRouter:         false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           2504001,
	},
}

var Prod_GoerliToAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_Goerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:          gethcommon.HexToAddress("0x25C819038b44183E746eA6dD12C49BF3E9e402aC"),
		OffRamp:         gethcommon.HexToAddress("0xF6FfEDeCE09B16e3D8055038BE22d5635A53F787"),
		CommitStore:     gethcommon.HexToAddress("0x56eDC4D8367932F0e36B966CbBd95dF48E9DB40F"),
		TokenSender:     gethcommon.HexToAddress("0xC5662F413AffaE59d214FC84BE92B469a92c077C"),
		MessageReceiver: gethcommon.HexToAddress("0x670bAeAa765CA179B82aDAA21947Ff02f819EbC0"),
		ReceiverDapp:    gethcommon.HexToAddress("0x6D984b7515604C27413BEFF5E92b3a1146E84B18"),
		GovernanceDapp:  gethcommon.HexToAddress(""),
		PingPongDapp:    gethcommon.HexToAddress("0x43A2A4C2ECB74FF45Eca704a14111d8f2B1c0fA0"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployRouter:         false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           7910311,
	},
}

var Prod_AvaxFujiToGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:          gethcommon.HexToAddress("0xA5007c13B7dF93e7647ffd671B9982cDced4E7Ff"),
		OffRamp:         gethcommon.HexToAddress("0x35BA9712D65EFb35f5092A6C3E832E38B6d17ccc"),
		CommitStore:     gethcommon.HexToAddress("0x177e068bc512AD99eC73dB6FEB7c731d9fea0CB3"),
		TokenSender:     gethcommon.HexToAddress("0xD6B8378092f590a39C360e8196101290551a66EA"),
		MessageReceiver: gethcommon.HexToAddress("0x4d57C6d8037C65fa66D6231844785a428310a735"),
		ReceiverDapp:    gethcommon.HexToAddress("0x8AB103843ED9D28D2C5DAf5FdB9c3e1CE2B6c876"),
		GovernanceDapp:  gethcommon.HexToAddress(""),
		PingPongDapp:    gethcommon.HexToAddress("0xACD8713E31B2CD1cf936673C4ccb8B5f16156129"),
	},
	DeploySettings: rhea.DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployRouter:         false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           15510013,
	},
}

var Prod_OptimismGoerlitoAvaxFuji = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_OptimismGoerli,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:          gethcommon.HexToAddress("0xCB07846c37aDaf8E90BD53f648Ac25B0bAaF970B"),
		OffRamp:         gethcommon.HexToAddress("0xa75b54b29Df4d38454f7Da5B6dF2f0a6b2c16514"),
		CommitStore:     gethcommon.HexToAddress("0xD9AC310783242A17e347Ee334a90B6cF3411f384"),
		TokenSender:     gethcommon.HexToAddress("0xD42be8Af3761DC2cb547B91e8c1B80067243eCFe"),
		MessageReceiver: gethcommon.HexToAddress("0x6154b0a8Ada0Da450E4226bf8772b3A1B756A152"),
		ReceiverDapp:    gethcommon.HexToAddress("0x201D1843707764CA2F236bd69E37CCbefF0827D4"),
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
		DeployedAt:           2517979,
	},
}

var Prod_AvaxFujitoOptimismGoerli = rhea.EvmDeploymentConfig{
	ChainConfig: Prod_AvaxFuji,
	LaneConfig: rhea.EVMLaneConfig{
		OnRamp:          gethcommon.HexToAddress("0x6B6b328cb1467D906389a1BBE54359c56000422c"),
		OffRamp:         gethcommon.HexToAddress("0x3E0df884042C21E83276abB368B4388a17f78A82"),
		CommitStore:     gethcommon.HexToAddress("0x6BEcd9eb4Df6Bf59152344fBcdC7919B9f38C6Ef"),
		TokenSender:     gethcommon.HexToAddress("0xa75b54b29Df4d38454f7Da5B6dF2f0a6b2c16514"),
		MessageReceiver: gethcommon.HexToAddress("0x651EF69F635f5017E281C2Ee09CDd4436560C89E"),
		ReceiverDapp:    gethcommon.HexToAddress("0x6eA3dE96a33617c3620b7c33c22656f860DDC255"),
		GovernanceDapp:  gethcommon.HexToAddress(""),
		PingPongDapp:    gethcommon.HexToAddress("0x1E357fc2a4AaB4Ec3382D5231F8A86E966Da3F28"),
	},

	DeploySettings: rhea.DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployCommitStore:    false,
		DeployRamp:           false,
		DeployRouter:         false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   false,
		DeployedAt:           15537046,
	},
}
