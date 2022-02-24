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
	Owner                 *bind.TransactOpts
	Client                *ethclient.Client
	ChainId               *big.Int
	LinkToken             common.Address
	SingleTokenOnramp     common.Address
	SingleTokenOfframp    common.Address
	LockUnlockPool        common.Address
	SingleTokenSender     common.Address
	SimpleMessageReceiver common.Address
	SingleTokenReceiver   common.Address
	MessageExecutor       common.Address
	Afn                   common.Address
	EthUrl                string
}

var Kovan = EvmChainConfig{
	ChainId:            big.NewInt(42),
	LinkToken:          common.HexToAddress("0xa36085F69e2889c224210F603D836748e7dC0088"),
	SingleTokenOnramp:  common.HexToAddress("0x7514aB087CcD8f5803CD20034eD9955cA5f2B99B"),
	SingleTokenOfframp: common.Address{},
	LockUnlockPool:     common.HexToAddress("0xc0299FBdAfdE8A91998989e6B36a78fE6b179112"),
	SingleTokenSender:  common.HexToAddress("0x8958F24DF47CE7C518ee4bD55d92406c23d834a1"),
	Afn:                common.HexToAddress("0xA6A3b37ACd10937D5857C7fac93E8BdBAc80424d"),
	EthUrl:             "wss://parity-kovan.eth.devnet.tools/ws",
}

var Rinkeby = EvmChainConfig{
	ChainId:               big.NewInt(4),
	LinkToken:             common.HexToAddress("0x01be23585060835e02b77ef475b0cc51aa1e0709"),
	SingleTokenOnramp:     common.Address{},
	SingleTokenOfframp:    common.HexToAddress("0x5f76AE26CaC2400861a00620af1F7ba8234Fb86F"),
	LockUnlockPool:        common.HexToAddress("0x759bB7c540900f489047adF66322D27E858A6281"),
	SimpleMessageReceiver: common.HexToAddress("0x869f73B2BA20efDb7d38534461fF9cCe901C4AEF"),
	SingleTokenReceiver:   common.HexToAddress("0x6A19BA4584230FC8c3D9Bcd8d54E48976af3f7DB"),
	MessageExecutor:       common.HexToAddress("0x79333B2A23BF56952E65F8aC0019e1B2ec3c8B88"),
	Afn:                   common.HexToAddress("0xb6067e139e90C17765116847D258578c6Db102dd"),
	EthUrl:                "wss://geth-rinkeby.eth.devnet.tools/ws",
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
