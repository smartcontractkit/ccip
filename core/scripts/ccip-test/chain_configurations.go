package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	ocrtypes2 "github.com/smartcontractkit/libocr/offchainreporting2/types"
)

type EvmChainConfig struct {
	Owner           *bind.TransactOpts
	Client          *ethclient.Client
	ChainId         *big.Int
	LinkToken       common.Address
	OnRamp          common.Address
	OffRamp         common.Address
	TokenPool       common.Address
	TokenSender     common.Address
	MessageReceiver common.Address
	TokenReceiver   common.Address
	MessageExecutor common.Address
	Afn             common.Address
	PriceFeed       common.Address
	EthUrl          string
	GasSettings     EVMGasSettings
}

// EVMGasSettings specifies the gas configuration for an EVM chain.
type EVMGasSettings struct {
	EIP1559   bool
	GasPrice  *big.Int
	GasTipCap *big.Int
}

// DefaultGasTipFee is the default gas tip fee of 2gwei.
var DefaultGasTipFee = big.NewInt(2e9)

// Rinkeby is configured to work as an onramp for Kovan
var Rinkeby = EvmChainConfig{
	ChainId:         big.NewInt(4),
	LinkToken:       common.HexToAddress("0x01be23585060835e02b77ef475b0cc51aa1e0709"),
	OnRamp:          common.Address{},
	OffRamp:         common.HexToAddress("0x1F10C5d25A08B21b3DCf9af95caEd54cE71a970B"),
	TokenPool:       common.HexToAddress("0xdA4318e7ec689f985E19F05d0351492244FFBa65"),
	MessageReceiver: common.HexToAddress("0xEb53a2B96aA765b078DCD6AB291E18BBA856f6Df"),
	TokenReceiver:   common.HexToAddress("0xaA8C18271dD9Dea20186b2Df8685156EA5eE43f5"),
	MessageExecutor: common.HexToAddress("0x11aF9a660AA8DF43D575541312caAE5d929f43F7"),
	Afn:             common.HexToAddress("0xb6067e139e90C17765116847D258578c6Db102dd"),
	PriceFeed:       common.HexToAddress("0xA2E0614f90BD107BF72faE4AD9251b63E18bA6f2"),
	EthUrl:          "wss://geth-rinkeby.eth.devnet.tools/ws",
	GasSettings: EVMGasSettings{
		EIP1559:   true,
		GasTipCap: DefaultGasTipFee,
	},
}

// Kovan is configured to work as an offramp for Rinkeby
var Kovan = EvmChainConfig{
	ChainId:     big.NewInt(42),
	LinkToken:   common.HexToAddress("0xa36085F69e2889c224210F603D836748e7dC0088"),
	OnRamp:      common.HexToAddress("0xB412Afaa0526979d85dcadf7Db0ADA40acCAaFa5"),
	OffRamp:     common.Address{},
	TokenPool:   common.HexToAddress("0x676f9261ecB2fd1B9d047BaC003bD30be61fDf24"),
	TokenSender: common.HexToAddress("0x8016b65003284abEA8405fBec65BF86BC7e460E0"),
	Afn:         common.HexToAddress("0xcB6f8a746db85f60a58Eba211E476601fd40A999"),
	PriceFeed:   common.HexToAddress("0xE96B14ddd215F2203601f2C5dc574DBBf02c19d9"),
	EthUrl:      "wss://parity-kovan.eth.devnet.tools/ws",
	GasSettings: EVMGasSettings{
		EIP1559:   true,
		GasTipCap: DefaultGasTipFee,
	},
}

// BSCTestnet is configured to be an offramp for PolygonMumbai
var BSCTestnet = EvmChainConfig{
	ChainId:         big.NewInt(97),
	LinkToken:       common.HexToAddress("0x84b9b910527ad5c03a9ca831909e21e236ea7b06"),
	OnRamp:          common.Address{},
	OffRamp:         common.HexToAddress("0xD4495b27DF53A15E513c43045C52f12f0EBa1f5e"),
	TokenPool:       common.HexToAddress("0x0924011a856483E47565d54BBC65cA9E21E8EE42"),
	MessageReceiver: common.HexToAddress("0x9E8C0E213e5B781D82134b8030fBfdEE379D54E7"),
	TokenReceiver:   common.HexToAddress("0x98Ed3e482a6133D8098A7D010808503F4224c6b3"),
	MessageExecutor: common.HexToAddress("0x7054E6420d2041c84b845c2f60537C7974a059d8"),
	Afn:             common.HexToAddress("0xCa3186CF799F07c68694737eC45026Bee3B4D9C2"),
	PriceFeed:       common.HexToAddress("0xCe078EA05aEA5342f3b8856004F8959930d00b69"),
	EthUrl:          "wss://binance-testnet.eth.devnet.tools/ws",
	GasSettings: EVMGasSettings{
		EIP1559:  false,
		GasPrice: big.NewInt(2e10),
	},
}

// PolygonMumbai is configured to be an onramp for BSCTestnet
var PolygonMumbai = EvmChainConfig{
	ChainId:     big.NewInt(80001),
	LinkToken:   common.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB"),
	OnRamp:      common.HexToAddress("0xD4495b27DF53A15E513c43045C52f12f0EBa1f5e"),
	OffRamp:     common.Address{},
	TokenPool:   common.HexToAddress("0x0924011a856483E47565d54BBC65cA9E21E8EE42"),
	TokenSender: common.HexToAddress("0x37816F4367064A17c000483c55c8a846Cab33051"),
	Afn:         common.HexToAddress("0xCa3186CF799F07c68694737eC45026Bee3B4D9C2"),
	PriceFeed:   common.HexToAddress("0xCe078EA05aEA5342f3b8856004F8959930d00b69"),
	EthUrl:      "wss://link-matic.getblock.io/testnet/axej8woh-seej-6ash-4Yu7-eyib1495dhno/",
	GasSettings: EVMGasSettings{
		EIP1559:  false,
		GasPrice: nil,
	},
}

const BootstrapPeerID = "12D3KooWNEeby2qn8hHhNEmX8WzWTRSoP46KqyfPtjhq6duRB1Zy"

func toOffchainPublicKey(s string) (key ocrtypes2.OffchainPublicKey) {
	copy(key[:], hexutil.MustDecode(s)[:])
	return
}

var Oracles = []confighelper2.OracleIdentityExtra{
	{
		// Node 0
		OracleIdentity: confighelper2.OracleIdentity{
			OnchainPublicKey:  common.HexToAddress("0x9546f9162e4dc7b4a03e55e988c5fda8dfe27cb1").Bytes(),
			TransmitAccount:   ocrtypes2.Account("0x0021B2310DB6679998ac483d841Ea72F691c9B50"),
			OffchainPublicKey: toOffchainPublicKey("0x3cdb3f0e649007f07bf0a72413304c06e686963399ec47b9bf836a6db3240765"),
			PeerID:            "12D3KooWBTk3X89nfcpX7VWxRmi3gQBoJ8GHDnUDhWsMWz9JbZMt",
		},
		ConfigEncryptionPublicKey: stringTo32Bytes("0x3b82e5e7d4be5a65c4115754342a53aab85a4180b96da3463724d357d171ae0c"),
	},
	{
		// Node 1
		OracleIdentity: confighelper2.OracleIdentity{
			OnchainPublicKey:  common.HexToAddress("0x63349e339807a177d2caa7409e0bc87bc031724f").Bytes(),
			TransmitAccount:   ocrtypes2.Account("0x256FF5d0406fE8a55B92135d3470707C23077F15"),
			OffchainPublicKey: toOffchainPublicKey("0xe3b4c7afbd138ff69660055b439a8d2e4c11175eec1a1206a66dad6ddce60a38"),
			PeerID:            "12D3KooWBMbX3cgV8QWBi92W9eMMjEpN1vB8wgdJ8jdtnZLPkvPd",
		},
		ConfigEncryptionPublicKey: stringTo32Bytes("0x2d631eb41938bc2e4589e8ab7d598dc8fa8e7c1e6f86c18b972699a5540e150e"),
	},
	{
		// Node 2
		OracleIdentity: confighelper2.OracleIdentity{
			OnchainPublicKey:  common.HexToAddress("0xd0da342c51f2790f9a21695e9efa45e3757b1337").Bytes(),
			TransmitAccount:   ocrtypes2.Account("0x645e882A796893Ba829179937d775b3a784b35A7"),
			OffchainPublicKey: toOffchainPublicKey("0x26d224b04e429a1f89b8d4d509aa6492c8d4985a78b869d6362d71cf7584423b"),
			PeerID:            "12D3KooWCDqaBN9PQaLk7ZxEYGanXusw9fyDbsTYRRRe3jwHk8xZ",
		},
		ConfigEncryptionPublicKey: stringTo32Bytes("0x3381bde66e3368bd8d100218beb59eb8232c8a2110dd3e4cbda179be6409ac27"),
	},
	{
		// Node 3
		OracleIdentity: confighelper2.OracleIdentity{
			OnchainPublicKey:  common.HexToAddress("0x89f8e7edd34432188220a736d3e64f6af2bf46a1").Bytes(),
			TransmitAccount:   ocrtypes2.Account("0x714546e24F8F7Ea328076718E4534D5a37F0c86B"),
			OffchainPublicKey: toOffchainPublicKey("0x2f7156e538da9f0fa259d63e0a0d587936342a98f469871bba0b546d3464d320"),
			PeerID:            "12D3KooWGEX8ynS2PWLNSmY6wHNrjSMTJm7SN2rtsxcYcALcvH3L",
		},
		ConfigEncryptionPublicKey: stringTo32Bytes("0x9c3de2854cd228f069660ae7d46dbf3b33550655d2ed1d6bbc3eeb11a0d73d25"),
	},
}

func stringTo32Bytes(s string) [32]byte {
	var b [32]byte
	copy(b[:], hexutil.MustDecode(s))
	return b
}
