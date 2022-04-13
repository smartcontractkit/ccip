package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	ocrtypes2 "github.com/smartcontractkit/libocr/offchainreporting2/types"

	"github.com/smartcontractkit/chainlink/core/logger"
)

type EvmChainConfig struct {
	Owner          *bind.TransactOpts
	Client         *ethclient.Client
	ChainId        *big.Int
	EthUrl         string
	GasSettings    EVMGasSettings
	Logger         logger.Logger
	DeploySettings DeploySettings

	LinkToken       common.Address
	BridgeTokens    []common.Address
	TokenPools      []common.Address
	PriceFeeds      []common.Address
	OnRamp          common.Address
	OnRampRouter    common.Address
	OffRamp         common.Address
	OffRampRouter   common.Address
	TokenSenders    []common.Address
	MessageReceiver common.Address
	TokenReceiver   common.Address
	MessageExecutor common.Address
	Afn             common.Address
}

// EVMGasSettings specifies the gas configuration for an EVM chain.
type EVMGasSettings struct {
	EIP1559   bool
	GasPrice  *big.Int
	GasTipCap *big.Int
}

// DefaultGasTipFee is the default gas tip fee of 2gwei.
var DefaultGasTipFee = big.NewInt(2e9)

type DeploySettings struct {
	DeployAFN        bool
	DeployTokenPools bool
	DeployPriceFeeds bool
}

// Rinkeby is configured to work as an onramp for Kovan
var Rinkeby = EvmChainConfig{
	ChainId: big.NewInt(4),
	EthUrl:  "wss://geth-rinkeby.eth.devnet.tools/ws",
	GasSettings: EVMGasSettings{
		EIP1559:   true,
		GasTipCap: DefaultGasTipFee,
	},
	LinkToken:     common.HexToAddress("0x01be23585060835e02b77ef475b0cc51aa1e0709"),
	BridgeTokens:  []common.Address{common.HexToAddress("0x01be23585060835e02b77ef475b0cc51aa1e0709")},
	TokenPools:    []common.Address{common.HexToAddress("0x232bc33b22501fC010835e872EA47969F0afb97A")},
	PriceFeeds:    []common.Address{common.HexToAddress("0x8eDcf07BfF3ee235B6b8Ba7FAe907E385c24e175")},
	OnRamp:        common.HexToAddress("0xda92A4766e1cFF0Aa77470fEccd095e471D8295F"),
	OnRampRouter:  common.HexToAddress(""),
	OffRamp:       common.Address{},
	OffRampRouter: common.HexToAddress(""),
	TokenSenders:  []common.Address{common.HexToAddress("0xC2F8a293ca968c449795d47337497F26b2813434"), common.HexToAddress("0x621686CC559AaC74af08C3cc829c696751c33000"), common.HexToAddress("0x496bBBD6840b31d3C54517ab81D3688eb962EBcD")},
	Afn:           common.HexToAddress("0xA4E4fd20121f268674d2C5c771e4cd27eAD0C543"),
	DeploySettings: DeploySettings{
		DeployAFN:        false,
		DeployTokenPools: false,
		DeployPriceFeeds: false,
	},
}

// Kovan is configured to be an offramp for Rinkeby
var Kovan = EvmChainConfig{
	ChainId: big.NewInt(42),
	EthUrl:  "wss://parity-kovan.eth.devnet.tools/ws",
	GasSettings: EVMGasSettings{
		EIP1559:   true,
		GasTipCap: DefaultGasTipFee,
	},
	LinkToken:       common.HexToAddress("0xa36085F69e2889c224210F603D836748e7dC0088"),
	BridgeTokens:    []common.Address{common.HexToAddress("0xa36085F69e2889c224210F603D836748e7dC0088")},
	TokenPools:      []common.Address{common.HexToAddress("0x8E63731A427d2D3E4b9358cDc7d75ff2650A6989")},
	PriceFeeds:      []common.Address{common.HexToAddress("0x7c132A947feC9C1aA17Db2d11eBc562243d4f331")},
	OnRamp:          common.Address{},
	OnRampRouter:    common.HexToAddress(""),
	OffRamp:         common.HexToAddress("0x5B48C82ED450e20E04E6133D26ddcA1eFF33D062"),
	OffRampRouter:   common.HexToAddress(""),
	MessageReceiver: common.HexToAddress("0x1a224Ab562D640aDE0FBc203bE619D5Cd3CEf935"),
	TokenReceiver:   common.HexToAddress("0xC4ED6c4F56ef4bD105e398BA320cF08e99C36E06"),
	MessageExecutor: common.HexToAddress("0x97403d2377e02eFc2D6e60D129435c0BA0917c9A"),
	Afn:             common.HexToAddress("0xE0271CC39005530A228E91bdF9454deEfdE47DCd"),
	DeploySettings: DeploySettings{
		DeployAFN:        false,
		DeployTokenPools: false,
		DeployPriceFeeds: false,
	},
}

// BSCTestnet is configured to be an offramp for Rinkeby
var BSCTestnet = EvmChainConfig{
	ChainId: big.NewInt(97),
	EthUrl:  "wss://binance-testnet.eth.devnet.tools/ws",
	GasSettings: EVMGasSettings{
		EIP1559:  false,
		GasPrice: big.NewInt(2e10),
	},
	LinkToken:       common.HexToAddress("0x84b9b910527ad5c03a9ca831909e21e236ea7b06"),
	BridgeTokens:    []common.Address{common.HexToAddress("0x84b9b910527ad5c03a9ca831909e21e236ea7b06")},
	TokenPools:      []common.Address{common.HexToAddress("0x6beeDC9AD88818d5B7658cD8F2ebd6BCa5302B22")},
	PriceFeeds:      []common.Address{common.HexToAddress("0x4082CE8081140b8369dD4c65B9ac7d360b47eeC5")},
	OnRamp:          common.Address{},
	OffRamp:         common.HexToAddress("0x57c9639Fd47Ba84F0ac5DCe3d6c368D8eDDEEf94"),
	MessageReceiver: common.HexToAddress("0xE2e54DF7E28e1Fa2ebAebBc9A0796A884Eb7B9BB"),
	TokenReceiver:   common.HexToAddress("0xD9B7CD57A2ADB88A28D48D61c3552DA748247a22"),
	MessageExecutor: common.HexToAddress("0x1e936E6e617E20ba3E73D7BF02b5A8b0bBb2673A"),
	Afn:             common.HexToAddress("0xaa9DFC20A535886974acdD23192D85f6A4c9a8D1"),
	DeploySettings: DeploySettings{
		DeployAFN:        false,
		DeployTokenPools: false,
		DeployPriceFeeds: false,
	},
}

// PolygonMumbai is configured to be an offramp for Rinkeby
var PolygonMumbai = EvmChainConfig{
	ChainId: big.NewInt(80001),
	EthUrl:  "wss://link-matic.getblock.io/testnet/axej8woh-seej-6ash-4Yu7-eyib1495dhno/",
	GasSettings: EVMGasSettings{
		EIP1559:  false,
		GasPrice: nil,
	},
	LinkToken:       common.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB"),
	BridgeTokens:    []common.Address{common.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB")},
	PriceFeeds:      []common.Address{common.HexToAddress("0xB42E8f41F21b3ce8C857a45972106ccE57cea0Fa")},
	TokenPools:      []common.Address{common.HexToAddress("0xf45818c983DD98792576062F128B4ad6E4b93632")},
	OnRamp:          common.Address{},
	OffRamp:         common.HexToAddress("0xB16eaA4596a2CedD765B85334448DB6C6Cb5c2FE"),
	MessageReceiver: common.HexToAddress("0x887F2081E5d3A3780098E3110E8b027848efF01c"),
	TokenReceiver:   common.HexToAddress("0x82a91b70A1470976979BE3862615A1A569fBb701"),
	MessageExecutor: common.HexToAddress("0xe3B3001a415072AF66A533376eb3182b1f47f646"),
	Afn:             common.HexToAddress("0x1c5cE558D50FaaFee9a9da89F5Db20aC7037E3Fb"),
	DeploySettings: DeploySettings{
		DeployAFN:        false,
		DeployTokenPools: false,
		DeployPriceFeeds: false,
	},
}

// BootstrapPeer in the format <BOOTSTRAP-PEER-ID>@<BOOTSTRAP-HOST>:<PORT>
const BootstrapPeer = "12D3KooWQDAsmFP4x8T8tfdgizgZpvVGJQpgPM1A8UcfbPz7jZiv@ccip-b-tcp.staging.org.devnet.tools:5001"

func toOffchainPublicKey(s string) (key ocrtypes2.OffchainPublicKey) {
	copy(key[:], hexutil.MustDecode(s)[:])
	return
}

var keyBundleIDs = []string{
	"1f1e25712701487c3f151e0eefbda7d3e1c6eb4786ccaf21ad2ad80cb7eefc75",
	"d81a2ee830974527f8fb45c3b8e7c021a85ab9f38f26aafe2bad90a5bfc0f3a3",
	"f212a4890a7db7fd8b6c6d76674aad221a74d38c37e1253078feaf8e3a8b88df",
	"52695bb2d815885932b89c55ada7a5c344b7e7df43bb8b37fad10c7013f71833"}

var transmitterAccounts = map[int64]map[int]string{
	4: {
		0: "0xf658D7d3aEdD70Cbe9050969cfe766eFC8Ef0341",
		1: "0x350c6b57923EBd32d4C928FE4e8A3D4b8b07ac33",
		2: "0x26A8bDD396acA21578d2A12b11477B4E5071fC4b",
		3: "0x19fA8Bd8fB7aB1Cb63615145FD81A4acb14dE09d",
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
				OnchainPublicKey:  common.HexToAddress("0x6590f85d9719b4ef1933e81c9f6edcea61c44132").Bytes(),
				OffchainPublicKey: toOffchainPublicKey("0x189c67b68665252873a9adba2a9e35595ac449e845af79902e863b796b674234"),
				PeerID:            "12D3KooWAwAN688dKMPn9b4An4sJZYWZ9kGE6kpbREWqda3c3hPz",
			},
			ConfigEncryptionPublicKey: stringTo32Bytes("0x8fc27f57ff9488c7b50b37933a40d6f064c0db6d7c0064d33a914e28326af613"),
		},
		{
			// Node 1
			OracleIdentity: confighelper2.OracleIdentity{
				OnchainPublicKey:  common.HexToAddress("0xfcea0c0f454b9d6b8a9fc632a65c9b4aa683e7e8").Bytes(),
				OffchainPublicKey: toOffchainPublicKey("0x66f1d0aeb6d29dcb3b0c8033cb03b851fb9ddc3cc9605c8c64914beab923ebbe"),
				PeerID:            "12D3KooWJb2extYFf9n67Sh9dnyb5Gc1xVTCdQjrWBdBg66gdUB5",
			},
			ConfigEncryptionPublicKey: stringTo32Bytes("0xe764a3b5d6f167609e5cccfead27665cd8ff2a1f3adf3349b933a6be8b947332"),
		},
		{
			// Node 2
			OracleIdentity: confighelper2.OracleIdentity{
				OnchainPublicKey:  common.HexToAddress("0x951ac5f47cf795db69e1cc38e0c05b5fbdef2cc0").Bytes(),
				OffchainPublicKey: toOffchainPublicKey("0x4a1a3ca0bc64d7d3cad969bde281d5a299b5b8b0f2ebc58bc43ac882588bda8b"),
				PeerID:            "12D3KooWBTLWnqv5YxUMacUvdRfFJw4LjGcAj56yq9ni2Jj99y27",
			},
			ConfigEncryptionPublicKey: stringTo32Bytes("0x696c138605964599a56a33a0db0030ae7fd0644730b8f5ba145c7dfecbdaa64a"),
		},
		{
			// Node 3
			OracleIdentity: confighelper2.OracleIdentity{
				OnchainPublicKey:  common.HexToAddress("0x5ae86428953108e602767f03ed58cfd4c7d28acb").Bytes(),
				OffchainPublicKey: toOffchainPublicKey("0xc5c254874d4ab6136418bfdb61ac4bf940a7728983fb55cba66ef4227a47d401"),
				PeerID:            "12D3KooWCHu8ja9Y5emXQmR9RLKa6TfUNC2L74EeYK11FXbrWzMq",
			},
			ConfigEncryptionPublicKey: stringTo32Bytes("0x2ba36ed1cffac6feed29306c71a4a7d1b13ec0c8b279d572458680bf366b0479"),
		},
	}

	for i := range oracles {
		oracles[i].TransmitAccount = ocrtypes2.Account(transmitterAccounts[chainID][i])
	}
	return oracles
}

func stringTo32Bytes(s string) [32]byte {
	var b [32]byte
	copy(b[:], hexutil.MustDecode(s))
	return b
}
