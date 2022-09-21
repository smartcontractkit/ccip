package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	ocrtypes2 "github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/secrets"
	"github.com/smartcontractkit/chainlink/core/scripts/common"
)

type EvmChainConfig struct {
	Owner          *bind.TransactOpts
	Client         *ethclient.Client
	ChainId        *big.Int
	EthUrl         string
	GasSettings    EVMGasSettings
	Logger         logger.Logger
	DeploySettings DeploySettings

	LinkToken       gethcommon.Address
	BridgeTokens    []gethcommon.Address
	TokenPools      []gethcommon.Address
	OnRamp          gethcommon.Address
	OnRampRouter    gethcommon.Address
	BlobVerifier    gethcommon.Address
	OffRampRouter   gethcommon.Address
	TokenSender     gethcommon.Address
	MessageReceiver gethcommon.Address
	ReceiverDapp    gethcommon.Address
	GovernanceDapp  gethcommon.Address
	OffRamp         gethcommon.Address
	Afn             gethcommon.Address
}

var pollPeriod = "1m0s"

// EVMGasSettings specifies the gas configuration for an EVM chain.
type EVMGasSettings struct {
	EIP1559   bool
	GasPrice  *big.Int
	GasTipCap *big.Int
}

// DefaultGasTipFee is the default gas tip fee of 1 gwei.
var DefaultGasTipFee = big.NewInt(1e9)

type DeploySettings struct {
	DeployAFN            bool
	DeployTokenPools     bool
	DeployRamp           bool
	DeployRouter         bool
	DeployBlobVerifier   bool
	DeployGovernanceDapp bool
	DeployedAt           uint64
}

// Rinkeby is configured to work as an onramp for Goerli
var Rinkeby = EvmChainConfig{
	ChainId: big.NewInt(4),
	EthUrl:  secrets.RinkebyEthURL,
	GasSettings: EVMGasSettings{
		EIP1559:   true,
		GasTipCap: DefaultGasTipFee,
	},
	LinkToken:      gethcommon.HexToAddress("0x01be23585060835e02b77ef475b0cc51aa1e0709"),
	BridgeTokens:   []gethcommon.Address{gethcommon.HexToAddress("0x01be23585060835e02b77ef475b0cc51aa1e0709")},
	TokenPools:     []gethcommon.Address{gethcommon.HexToAddress("0x9698EB719d6807348cE0C8269d7dd6f12700fC69")},
	OnRamp:         gethcommon.HexToAddress("0x0a33b4fd85Ed50AB3484eDfd3a7D49783fAd2Cf2"),
	OnRampRouter:   gethcommon.HexToAddress("0x5C535256fCac5493df2DBE54B7E097590A77f6fA"),
	TokenSender:    gethcommon.HexToAddress("0x4a96E9DCD32A5de534A4DEC740f68262ddBe5B58"),
	GovernanceDapp: gethcommon.HexToAddress("0xf7f19036d8fb9e9D0F1bA91CAb6f460781355e83"),
	Afn:            gethcommon.HexToAddress("0xd65C654AbF5D79e1cc4615E62C647a814E98c5D5"),
	DeploySettings: DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployRamp:           false,
		DeployRouter:         false,
		DeployGovernanceDapp: true,
		DeployedAt:           11352595,
	},
}

// Goerli is configured to work as an offRamp for Rinkeby
var Goerli = EvmChainConfig{
	ChainId: big.NewInt(5),
	EthUrl:  secrets.GoerliEthURL,
	GasSettings: EVMGasSettings{
		EIP1559:   true,
		GasTipCap: DefaultGasTipFee,
	},
	LinkToken:       gethcommon.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB"),
	BridgeTokens:    []gethcommon.Address{gethcommon.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB")},
	TokenPools:      []gethcommon.Address{gethcommon.HexToAddress("0x4E4cdb6cEB8Fb2940a97Cf0dBC8581c8A420Ee6f")},
	BlobVerifier:    gethcommon.HexToAddress("0xf32202e2C0B75854851EF7e4e9E9929bBB6C35EC"),
	OffRamp:         gethcommon.HexToAddress("0x235dC05991791f199D766A1b3fB2D1b80E1CfD6B"),
	OffRampRouter:   gethcommon.HexToAddress("0xf36B9824dD342a9F69Ad4FB3599f1A7FC9373608"),
	MessageReceiver: gethcommon.HexToAddress("0x430a78789539e80B6eD1A6F3E52f51280A159BCd"),
	ReceiverDapp:    gethcommon.HexToAddress("0x4308444c67DeeA56644a861460231F5fdBC1Ad0b"),
	GovernanceDapp:  gethcommon.HexToAddress("0xebAd1E6cd98A93C5Df4e8A9FEf7b7d1AE892A764"),
	Afn:             gethcommon.HexToAddress("0x77b04A38e53fCF18edc2220dcBD7635f9E5D8F76"),
	DeploySettings: DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployBlobVerifier:   false,
		DeployRamp:           true,
		DeployRouter:         true,
		DeployGovernanceDapp: true,
		DeployedAt:           7573621,
	},
}

// Sepolia is configured to work as an onramp for Goerli
var Sepolia = EvmChainConfig{
	ChainId: big.NewInt(11155111),
	EthUrl:  secrets.SepoliaEthURL,
	GasSettings: EVMGasSettings{
		EIP1559:   true,
		GasTipCap: DefaultGasTipFee,
	},
	LinkToken:    gethcommon.HexToAddress("0xb227f007804c16546Bd054dfED2E7A1fD5437678"),
	BridgeTokens: []gethcommon.Address{gethcommon.HexToAddress("0xb227f007804c16546Bd054dfED2E7A1fD5437678")},
	TokenPools:   []gethcommon.Address{gethcommon.HexToAddress("")},
	OnRamp:       gethcommon.HexToAddress(""),
	OnRampRouter: gethcommon.HexToAddress(""),
	TokenSender:  gethcommon.HexToAddress(""),
	Afn:          gethcommon.HexToAddress(""),
	DeploySettings: DeploySettings{
		DeployAFN:        false,
		DeployTokenPools: false,
		DeployRamp:       false,
		DeployRouter:     false,
		DeployedAt:       1592304,
	},
}

// BSCTestnet is configured to be an offramp for Rinkeby
var BSCTestnet = EvmChainConfig{
	ChainId: big.NewInt(97),
	EthUrl:  secrets.BSCEthURL,
	GasSettings: EVMGasSettings{
		EIP1559:  false,
		GasPrice: big.NewInt(20e9),
	},
	LinkToken:       gethcommon.HexToAddress("0x84b9b910527ad5c03a9ca831909e21e236ea7b06"),
	BridgeTokens:    []gethcommon.Address{gethcommon.HexToAddress("0x84b9b910527ad5c03a9ca831909e21e236ea7b06")},
	TokenPools:      []gethcommon.Address{gethcommon.HexToAddress("0xc99148f5e687CB16511EFcA6668F2b6eCe63458C")},
	OffRamp:         gethcommon.HexToAddress("0xC0a1fFeAefd1544A454A49f3c4319B11cD4fDf1D"),
	OffRampRouter:   gethcommon.HexToAddress("0x3a9e41F6a28331bcc3a4ca4c58b844Cd2Fd217bb"),
	BlobVerifier:    gethcommon.HexToAddress("0x3755b7B14e9c71C080787e084471e5f51BBD2Cc6"),
	MessageReceiver: gethcommon.HexToAddress("0xB0Fa66B3B165D10ED46F3e33E2a45926d958d391"),
	ReceiverDapp:    gethcommon.HexToAddress("0xbaEf074daeE1F4Cdc48eb9F6877C03fEA3039Fd8"),
	Afn:             gethcommon.HexToAddress("0x77ad33c539Eff83732379B1C0474901Db6EB4B44"),
	DeploySettings: DeploySettings{
		DeployAFN:          false,
		DeployTokenPools:   false,
		DeployBlobVerifier: false,
		DeployRamp:         true,
		DeployRouter:       true,
	},
}

// PolygonMumbai is configured to be an offramp for Rinkeby
var PolygonMumbai = EvmChainConfig{
	ChainId: big.NewInt(80001),
	EthUrl:  secrets.PolygonEthURL,
	GasSettings: EVMGasSettings{
		EIP1559:  false,
		GasPrice: nil,
	},
	LinkToken:       gethcommon.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB"),
	BridgeTokens:    []gethcommon.Address{gethcommon.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB")},
	TokenPools:      []gethcommon.Address{gethcommon.HexToAddress("0xf45818c983DD98792576062F128B4ad6E4b93632")},
	OnRamp:          gethcommon.Address{},
	BlobVerifier:    gethcommon.HexToAddress("0xB16eaA4596a2CedD765B85334448DB6C6Cb5c2FE"),
	MessageReceiver: gethcommon.HexToAddress("0x887F2081E5d3A3780098E3110E8b027848efF01c"),
	ReceiverDapp:    gethcommon.HexToAddress("0x82a91b70A1470976979BE3862615A1A569fBb701"),
	OffRamp:         gethcommon.HexToAddress("0xe3B3001a415072AF66A533376eb3182b1f47f646"),
	Afn:             gethcommon.HexToAddress("0x1c5cE558D50FaaFee9a9da89F5Db20aC7037E3Fb"),
	DeploySettings: DeploySettings{
		DeployAFN:        true,
		DeployTokenPools: true,
	},
}

// BootstrapPeer in the format <BOOTSTRAP-PEER-ID>@<BOOTSTRAP-HOST>:<PORT>
const BootstrapPeer = "12D3KooWQDAsmFP4x8T8tfdgizgZpvVGJQpgPM1A8UcfbPz7jZiv@ccip-b-tcp.staging.org.devnet.tools:5001"

var keyBundleIDs = []string{
	"4e57ae6c96090fe59e837feb9bc4bc265bb9f2328a7fd4b6fde9e803fb6d5665",
	"17a2fd5637323d1665f2a991459cd60a8aa5fbcb084e18f45b82f1fd8dcadb57",
	"c07e9aa1246e97731cc7344699173d60a170cb7ff88cca4543a02005a70c252b",
	"c7ec6a5fcd84bfc91ed2b163f5f0ab52f5c03dd56647387cd45119f5dbf9fe82"}

var transmitterAccounts = map[int64]map[int]string{
	4: {
		0: "0xf658D7d3aEdD70Cbe9050969cfe766eFC8Ef0341",
		1: "0x350c6b57923EBd32d4C928FE4e8A3D4b8b07ac33",
		2: "0x26A8bDD396acA21578d2A12b11477B4E5071fC4b",
		3: "0x19fA8Bd8fB7aB1Cb63615145FD81A4acb14dE09d",
	},
	5: {
		0: "0xdc4FB792d2aa782FF4E0689e9E4030C5f8171807",
		1: "0xea4E47518D611Bb467e0FFa345768C6d352588f7",
		2: "0xfc0bD918A1dCb0d3ca17AE261aD057527A086fE2",
		3: "0xe88ff73814fB891bb0e149F5578796fa41F20242",
	},
	42: {
		0: "0x93A022332C95128Fd48fE44853836E038062509A",
		1: "0x24CA26d8Fd0f45F70F8A63F74cfCaAfa535df46a",
		2: "0x0Af5e07Bee36758863052d1632d81Ded38f59054",
		3: "0x43b326bA02FE60bb8109baD7d1F1d6F3fA1e6858",
	},
	97: {
		0: "0x2E5b8cDc9b0345c1cc12CFc36Da91E613F15652a",
		1: "0x51a2D14737DBE5A2E74ADE092CBAAc2038b9a9F2",
		2: "0x9fDA7a5595eC5BaD4629a69F8E140236Fb0486CA",
		3: "0x150aBCc822C4686B1032E36069d2240189437491",
	},
	80001: {
		0: "0x3FeC3B77e452C2AF6f2DDFa6eE2D1Be812BE943A",
		1: "0xa47eAA24f1Df5Ce53f9599CB86dAe30AB791A3D7",
		2: "0xe52290b55CdfBC07Ad8516151f621198313eCe4D",
		3: "0x28081674f6f66f54CB8b1c74e8e750dA71d48464",
	},
}

func getOraclesForChain(chainID int64) (or []confighelper2.OracleIdentityExtra) {
	var oracles = []confighelper2.OracleIdentityExtra{
		{
			// Node 0
			OracleIdentity: confighelper2.OracleIdentity{
				OnchainPublicKey:  gethcommon.HexToAddress("0x6590f85d9719b4ef1933e81c9f6edcea61c44132").Bytes(),
				OffchainPublicKey: common.ToOffchainPublicKey("0x189c67b68665252873a9adba2a9e35595ac449e845af79902e863b796b674234"),
				PeerID:            "12D3KooWAwAN688dKMPn9b4An4sJZYWZ9kGE6kpbREWqda3c3hPz",
			},
			ConfigEncryptionPublicKey: common.StringTo32Bytes("0x8fc27f57ff9488c7b50b37933a40d6f064c0db6d7c0064d33a914e28326af613"),
		},
		{
			// Node 1
			OracleIdentity: confighelper2.OracleIdentity{
				OnchainPublicKey:  gethcommon.HexToAddress("0xfcea0c0f454b9d6b8a9fc632a65c9b4aa683e7e8").Bytes(),
				OffchainPublicKey: common.ToOffchainPublicKey("0x66f1d0aeb6d29dcb3b0c8033cb03b851fb9ddc3cc9605c8c64914beab923ebbe"),
				PeerID:            "12D3KooWJb2extYFf9n67Sh9dnyb5Gc1xVTCdQjrWBdBg66gdUB5",
			},
			ConfigEncryptionPublicKey: common.StringTo32Bytes("0xe764a3b5d6f167609e5cccfead27665cd8ff2a1f3adf3349b933a6be8b947332"),
		},
		{
			// Node 2
			OracleIdentity: confighelper2.OracleIdentity{
				OnchainPublicKey:  gethcommon.HexToAddress("0x951ac5f47cf795db69e1cc38e0c05b5fbdef2cc0").Bytes(),
				OffchainPublicKey: common.ToOffchainPublicKey("0x4a1a3ca0bc64d7d3cad969bde281d5a299b5b8b0f2ebc58bc43ac882588bda8b"),
				PeerID:            "12D3KooWBTLWnqv5YxUMacUvdRfFJw4LjGcAj56yq9ni2Jj99y27",
			},
			ConfigEncryptionPublicKey: common.StringTo32Bytes("0x696c138605964599a56a33a0db0030ae7fd0644730b8f5ba145c7dfecbdaa64a"),
		},
		{
			// Node 3
			OracleIdentity: confighelper2.OracleIdentity{
				OnchainPublicKey:  gethcommon.HexToAddress("0x5ae86428953108e602767f03ed58cfd4c7d28acb").Bytes(),
				OffchainPublicKey: common.ToOffchainPublicKey("0xc5c254874d4ab6136418bfdb61ac4bf940a7728983fb55cba66ef4227a47d401"),
				PeerID:            "12D3KooWCHu8ja9Y5emXQmR9RLKa6TfUNC2L74EeYK11FXbrWzMq",
			},
			ConfigEncryptionPublicKey: common.StringTo32Bytes("0x2ba36ed1cffac6feed29306c71a4a7d1b13ec0c8b279d572458680bf366b0479"),
		},
	}

	for i := range oracles {
		oracles[i].TransmitAccount = ocrtypes2.Account(transmitterAccounts[chainID][i])
	}
	return oracles
}
