package rhea

import (
	"math/big"

	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/secrets"
)

var GoerliConfig = EvmChainConfig{
	ChainId: big.NewInt(5),
	EthUrl:  secrets.GoerliEthURL,
	GasSettings: EVMGasSettings{
		EIP1559:   true,
		GasTipCap: DefaultGasTipFee,
	},
	LinkToken:       gethcommon.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB"),
	BridgeTokens:    []gethcommon.Address{gethcommon.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB")},
	TokenPools:      []gethcommon.Address{gethcommon.HexToAddress("0x4c10d67E4B8e18a67A7606DEFDCe42CCc281D39B")},
	BlobVerifier:    gethcommon.HexToAddress("0x701Fe16916dd21EFE2f535CA59611D818B017877"),
	OnRamp:          gethcommon.HexToAddress("0xF9a21B587111e7E8745Fb8b13750014f19DB0014"),
	OnRampRouter:    gethcommon.HexToAddress("0xA189971a2c5AcA0DFC5Ee7a2C44a2Ae27b3CF389"),
	TokenSender:     gethcommon.HexToAddress("0xc3e8bB61e1db9adE45F76237d75AAfaCca2066AF"),
	OffRamp:         gethcommon.HexToAddress("0xD42be8Af3761DC2cb547B91e8c1B80067243eCFe"),
	OffRampRouter:   gethcommon.HexToAddress("0xb78d314d32EB4B01C459EDE0774cc3b6AF244Dd7"),
	MessageReceiver: gethcommon.HexToAddress("0xe0D4860bD0429B87f508f0aE8d1789cC0adbbfcA"),
	ReceiverDapp:    gethcommon.HexToAddress("0x84B7B012c95f8A152B44Ab3e952f2dEE424fA8e1"),
	GovernanceDapp:  gethcommon.HexToAddress(""),
	PingPongDapp:    gethcommon.HexToAddress("0x201D1843707764CA2F236bd69E37CCbefF0827D4"),
	Afn:             gethcommon.HexToAddress("0x8a710bBd77661D168D5A6725bD2E514ba1bFf59d"),
	DeploySettings: DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployBlobVerifier:   true,
		DeployRamp:           true,
		DeployRouter:         false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   true,
		DeployedAt:           7722040,
	},
}

var OptimismGoerliConfig = EvmChainConfig{
	ChainId: big.NewInt(420),
	EthUrl:  secrets.OptimismGoerliURL,
	GasSettings: EVMGasSettings{
		EIP1559: false,
	},
	LinkToken:       gethcommon.HexToAddress("0xdc2CC710e42857672E7907CF474a69B63B93089f"),
	BridgeTokens:    []gethcommon.Address{gethcommon.HexToAddress("0xdc2CC710e42857672E7907CF474a69B63B93089f")},
	TokenPools:      []gethcommon.Address{gethcommon.HexToAddress("0xE4aB69C077896252FAFBD49EFD26B5D171A32410")},
	BlobVerifier:    gethcommon.HexToAddress("0x4A1d9c5a7f9f9de7D5d8eC0f96f7213b0AB953d9"),
	OnRamp:          gethcommon.HexToAddress("0x56eDC4D8367932F0e36B966CbBd95dF48E9DB40F"),
	OnRampRouter:    gethcommon.HexToAddress("0xE591bf0A0CF924A0674d7792db046B23CEbF5f34"),
	TokenSender:     gethcommon.HexToAddress("0x51298c07eF8849f89552C2B3184741a759d4B37C"),
	OffRamp:         gethcommon.HexToAddress("0x20D3D6F9CFc3268437Aff3d55e612309Fd7dfB8C"),
	OffRampRouter:   gethcommon.HexToAddress("0x2b7aB40413DA5077E168546eA376920591Aee8E7"),
	MessageReceiver: gethcommon.HexToAddress("0x2321F13659889c2f1e7a62A7700744E36F9C60E5"),
	ReceiverDapp:    gethcommon.HexToAddress("0xA189971a2c5AcA0DFC5Ee7a2C44a2Ae27b3CF389"),
	GovernanceDapp:  gethcommon.HexToAddress(""),
	PingPongDapp:    gethcommon.HexToAddress("0xb6E24bd5376f808a8f4cEf945c96ec5582791255"),
	Afn:             gethcommon.HexToAddress("0x4c10d67E4B8e18a67A7606DEFDCe42CCc281D39B"),
	DeploySettings: DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployBlobVerifier:   false,
		DeployRamp:           false,
		DeployRouter:         false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   true,
		DeployedAt:           1770383,
	},
}

var GoerliAvaxConfig = EvmChainConfig{
	ChainId: big.NewInt(5),
	EthUrl:  secrets.GoerliEthURL,
	GasSettings: EVMGasSettings{
		EIP1559:   true,
		GasTipCap: DefaultGasTipFee,
	},
	LinkToken:       gethcommon.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB"),
	BridgeTokens:    []gethcommon.Address{gethcommon.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB")},
	TokenPools:      []gethcommon.Address{gethcommon.HexToAddress("0x4c10d67E4B8e18a67A7606DEFDCe42CCc281D39B")},
	BlobVerifier:    gethcommon.HexToAddress("0x56eDC4D8367932F0e36B966CbBd95dF48E9DB40F"),
	OnRamp:          gethcommon.HexToAddress("0xA2Ecbbe981d06F096cDC1698C70A7a8Ce34e4852"),
	OnRampRouter:    gethcommon.HexToAddress("0xA189971a2c5AcA0DFC5Ee7a2C44a2Ae27b3CF389"),
	TokenSender:     gethcommon.HexToAddress("0xC5662F413AffaE59d214FC84BE92B469a92c077C"),
	OffRamp:         gethcommon.HexToAddress("0x193eD728A2c7e9E296a9FA05A068757b05c73776"),
	OffRampRouter:   gethcommon.HexToAddress("0xb78d314d32EB4B01C459EDE0774cc3b6AF244Dd7"),
	MessageReceiver: gethcommon.HexToAddress("0x670bAeAa765CA179B82aDAA21947Ff02f819EbC0"),
	ReceiverDapp:    gethcommon.HexToAddress("0x6D984b7515604C27413BEFF5E92b3a1146E84B18"),
	GovernanceDapp:  gethcommon.HexToAddress(""),
	PingPongDapp:    gethcommon.HexToAddress("0x43A2A4C2ECB74FF45Eca704a14111d8f2B1c0fA0"),
	Afn:             gethcommon.HexToAddress("0x8a710bBd77661D168D5A6725bD2E514ba1bFf59d"),
	DeploySettings: DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployBlobVerifier:   false,
		DeployRamp:           true,
		DeployRouter:         false,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   true,
		DeployedAt:           7757359,
	},
}

var AvaxFuji = EvmChainConfig{
	ChainId: big.NewInt(43113),
	EthUrl:  secrets.AvaxFujiURL,
	GasSettings: EVMGasSettings{
		EIP1559: false,
	},
	LinkToken:       gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846"),
	BridgeTokens:    []gethcommon.Address{gethcommon.HexToAddress("0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846")},
	TokenPools:      []gethcommon.Address{gethcommon.HexToAddress("0xb6f1Fe2CDE891eFd5Efd2A563C4C2F2549163718")},
	BlobVerifier:    gethcommon.HexToAddress("0x177e068bc512AD99eC73dB6FEB7c731d9fea0CB3"),
	OnRamp:          gethcommon.HexToAddress("0x7f63086AFf4C189ADf8FA2d955410dbdBF771eFd"),
	OnRampRouter:    gethcommon.HexToAddress("0xc0A2c03115d1B48BAA59f676c108EfE5Ba3ee062"),
	TokenSender:     gethcommon.HexToAddress("0xD6B8378092f590a39C360e8196101290551a66EA"),
	OffRamp:         gethcommon.HexToAddress("0x88A2d74F47a237a62e7A51cdDa67270CE381555e"),
	OffRampRouter:   gethcommon.HexToAddress("0x7d5297c5506ee2A7Ef121Da9bE02b6a6AD30b392"),
	MessageReceiver: gethcommon.HexToAddress("0x4d57C6d8037C65fa66D6231844785a428310a735"),
	ReceiverDapp:    gethcommon.HexToAddress("0x8AB103843ED9D28D2C5DAf5FdB9c3e1CE2B6c876"),
	GovernanceDapp:  gethcommon.HexToAddress(""),
	PingPongDapp:    gethcommon.HexToAddress("0xACD8713E31B2CD1cf936673C4ccb8B5f16156129"),
	Afn:             gethcommon.HexToAddress("0xb2958D1Bd07448865E555FeeFf32b58D254ffB4C"),
	DeploySettings: DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployBlobVerifier:   true,
		DeployRamp:           true,
		DeployRouter:         true,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   true,
		DeployedAt:           14566887,
	},
}

var SepoliaConfig = EvmChainConfig{
	ChainId: big.NewInt(11155111),
	EthUrl:  secrets.SepoliaEthURL,
	GasSettings: EVMGasSettings{
		EIP1559:   true,
		GasTipCap: DefaultGasTipFee,
	},
	LinkToken:       gethcommon.HexToAddress("0xb227f007804c16546Bd054dfED2E7A1fD5437678"),
	BridgeTokens:    []gethcommon.Address{gethcommon.HexToAddress("0xb227f007804c16546Bd054dfED2E7A1fD5437678")},
	TokenPools:      []gethcommon.Address{gethcommon.HexToAddress("")},
	BlobVerifier:    gethcommon.HexToAddress(""),
	OnRamp:          gethcommon.HexToAddress(""),
	OnRampRouter:    gethcommon.HexToAddress(""),
	TokenSender:     gethcommon.HexToAddress(""),
	OffRamp:         gethcommon.HexToAddress(""),
	OffRampRouter:   gethcommon.HexToAddress(""),
	MessageReceiver: gethcommon.HexToAddress(""),
	ReceiverDapp:    gethcommon.HexToAddress(""),
	GovernanceDapp:  gethcommon.HexToAddress(""),
	PingPongDapp:    gethcommon.HexToAddress(""),
	Afn:             gethcommon.HexToAddress(""),
	DeploySettings: DeploySettings{
		DeployAFN:            true,
		DeployTokenPools:     true,
		DeployBlobVerifier:   true,
		DeployRamp:           true,
		DeployRouter:         true,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   true,
		DeployedAt:           1592304,
	},
}
