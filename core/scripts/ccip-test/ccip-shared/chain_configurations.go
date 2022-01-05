package ccip_shared

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	confighelper2 "github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	ocrtypes2 "github.com/smartcontractkit/libocr/offchainreporting2/types"
)

type EvmChainConfig struct {
	Owner                 *bind.TransactOpts
	Users                 []*bind.TransactOpts
	Client                *ethclient.Client
	ChainId               *big.Int
	LinkToken             common.Address
	SingleTokenOnramp     common.Address
	SingleTokenOfframp    common.Address
	LockUnlockPool        common.Address
	SimpleMessageReceiver common.Address
	SingleTokenSender     common.Address
	SingleTokenReceiver   common.Address
	MessageExecutor       common.Address
	Afn                   common.Address
	EthUrl                string
}

var Kovan = EvmChainConfig{
	LinkToken:          common.HexToAddress("0xa36085F69e2889c224210F603D836748e7dC0088"),
	SingleTokenOfframp: common.Address{},
	SingleTokenOnramp:  common.HexToAddress("0x7590f49f088a7B74596712A1DB4C39D9FAB00347"),
	SingleTokenSender:  common.HexToAddress("0xCfBc79f4042Be99a7292Ff33466cEE6652F40485"),
	LockUnlockPool:     common.HexToAddress("0x0C710E14226A43301028cEf2D0D492fDE7E3024A"),
	Afn:                common.HexToAddress("0xB86654A84CF21a913f39d9Da126C27f56Df07166"),
	EthUrl:             "wss://parity-kovan.eth.devnet.tools/ws",
	ChainId:            big.NewInt(42),
}

var Rinkeby = EvmChainConfig{
	LinkToken:             common.HexToAddress("0x01be23585060835e02b77ef475b0cc51aa1e0709"),
	SingleTokenOfframp:    common.HexToAddress("0xC3376eD981978E0C107a52f0488f1d131B9ECCAc"),
	SingleTokenOnramp:     common.Address{},
	LockUnlockPool:        common.HexToAddress("0x33079B10A1417EF666040BF5aAF5623FCc90FAFe"),
	SimpleMessageReceiver: common.HexToAddress("0x0389eF5B01822F673cFb87cdf3D8f97E0FaDBf77"),
	SingleTokenReceiver:   common.HexToAddress("0xF47C5C5cEeE3F77954Fa2eA58690e44fD6658B9F"),
	MessageExecutor:       common.HexToAddress("0x21d5C93B2A22Bdc315F1760E92960Cda23D93f3E"),
	Afn:                   common.HexToAddress("0x1E275452a2bD9154EC0F46aE21881E47Aed03E3e"),
	EthUrl:                "wss://geth-rinkeby.eth.devnet.tools/ws",
	ChainId:               big.NewInt(4),
}

const BootstrapPeerID = "12D3KooWFHTQLnS1dzmRoqit8zyLx7ost7sm8pSjFQSByfjsoqyT"

var Oracles = []confighelper2.OracleIdentityExtra{
	{
		// Node 0
		OracleIdentity: confighelper2.OracleIdentity{
			OnchainPublicKey:  common.HexToAddress("0x69B8fADd511A2BE6d90A5dA5F617EB48cE3FA132").Bytes(),
			TransmitAccount:   ocrtypes2.Account("0x1b9aC605d2b2E2E9Db4cac561181Ec10A938390c"),
			OffchainPublicKey: hexutil.MustDecode("0x17992ca120fe8a3075e6c8b3e8c93f06fc3fc5dc5f989d54ec14def8cf080d06"),
			PeerID:            "12D3KooWPRpNDEzJKJevcwhdjKvTWEBV4o9RFJ8FmzPf9ErsPtBM",
		},
		ConfigEncryptionPublicKey: stringTo32Bytes("0x69a21497b875787e4810d2d825aefca5f9ee6dc3e97f51b93b33de67300c402f"),
	},
	{
		// Node 1
		OracleIdentity: confighelper2.OracleIdentity{
			OnchainPublicKey:  common.HexToAddress("0x51A4282729AFE2A7967ab24ff707AffCe1dcc678").Bytes(),
			TransmitAccount:   ocrtypes2.Account("0x2FF79Fff751a157054629eECF2B32aE671d72Bf8"),
			OffchainPublicKey: hexutil.MustDecode("0xd7f949bb2ff6242f2d5158b2f54eb0b629904dddfaa9d699736e7265eb87bb2f"),
			PeerID:            "12D3KooWAPnKdfa3wPobf3FdErZu1VAKKMmuoEHwvmcjnSQhYSvD",
		},
		ConfigEncryptionPublicKey: stringTo32Bytes("0x4320cf6a9be0ffdd4e44787551bfda49950288c31d6854ba5f243e9ea23e5278"),
	},
	{
		// Node 2
		OracleIdentity: confighelper2.OracleIdentity{
			OnchainPublicKey:  common.HexToAddress("0x9d51eeF5292d2fFE9bEa7c263CF1fe18e9f35148").Bytes(),
			TransmitAccount:   ocrtypes2.Account("0xC81C5cccfcA5B95526609575235D55077A25F105"),
			OffchainPublicKey: hexutil.MustDecode("0x65b165e268405827411a79384bae8648f7725d701bc4d8373fdd55838802e4f6"),
			PeerID:            "12D3KooWJnTuDhN1GCSbxWjNW51P7z1QRbC3VJbtY6wuL1VUXkQu",
		},
		ConfigEncryptionPublicKey: stringTo32Bytes("0xc4717af64f5e4235c07e893159c522b12dc0809982f09f519d873d6194129a43"),
	},
	{
		// Node 3
		OracleIdentity: confighelper2.OracleIdentity{
			OnchainPublicKey:  common.HexToAddress("0xaaeB8784265a6ee8181729dDD0Aea99c60814482").Bytes(),
			TransmitAccount:   ocrtypes2.Account("0x1FD884B9088d2013B6c2EC2F9640F551578e2f1C"),
			OffchainPublicKey: hexutil.MustDecode("0x61c4c6a6e9a2ac020e87e2e7e8c88e32373a503a1ae7d1a651b1ac08bb7c31f5"),
			PeerID:            "12D3KooWSDzVm7Kv3xSHB17aUQ5UvBJay2cxC7XfTjGsqGn7MDK7",
		},
		ConfigEncryptionPublicKey: stringTo32Bytes("0xe0b8876b62cb1c5c827be6f6dc271ce7702f06a8bf7d5c289486b2a6c8a21e19"),
	},
}

func stringTo32Bytes(s string) [32]byte {
	var b [32]byte
	copy(b[:], hexutil.MustDecode(s))
	return b
}

func (config *EvmChainConfig) SetupClient() *EvmChainConfig {
	client, err := ethclient.Dial(config.EthUrl)
	PanicErr(err)
	config.Client = client
	return config
}

func (config *EvmChainConfig) SetOwnerAndUsers(ownerPrivateKey string, seedKey string) *EvmChainConfig {
	config.SetOwner(ownerPrivateKey)

	var users []*bind.TransactOpts
	seedKeyWithoutFirstChar := seedKey[1:]
	fmt.Println("--- Addresses of the seed key")
	for i := 0; i <= 9; i++ {
		key, err := crypto.HexToECDSA(strconv.Itoa(i) + seedKeyWithoutFirstChar)
		PanicErr(err)
		user, err := bind.NewKeyedTransactorWithChainID(key, config.ChainId)
		PanicErr(err)
		users = append(users, user)
		fmt.Println(user.From.Hex())
	}
	fmt.Println("---")

	config.Users = users

	return config
}

func (config *EvmChainConfig) SetOwner(ownerPrivateKey string) *EvmChainConfig {
	ownerKey, err := crypto.HexToECDSA(ownerPrivateKey)
	PanicErr(err)
	user, err := bind.NewKeyedTransactorWithChainID(ownerKey, config.ChainId)
	PanicErr(err)
	fmt.Println("--- Owner address ")
	fmt.Println(user.From.Hex())
	config.Owner = user
	return config
}
