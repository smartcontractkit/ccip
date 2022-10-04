package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/secrets"
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
	PingPongDapp    gethcommon.Address
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
	DeployPingPongDapp   bool
	DeployedAt           uint64
}

var RinkebyConfig = EvmChainConfig{
	ChainId: big.NewInt(4),
	EthUrl:  secrets.RinkebyEthURL,
	GasSettings: EVMGasSettings{
		EIP1559:   true,
		GasTipCap: DefaultGasTipFee,
	},
	LinkToken:       gethcommon.HexToAddress("0x01be23585060835e02b77ef475b0cc51aa1e0709"),
	BridgeTokens:    []gethcommon.Address{gethcommon.HexToAddress("0x01be23585060835e02b77ef475b0cc51aa1e0709")},
	TokenPools:      []gethcommon.Address{gethcommon.HexToAddress("0x4169578D48aa81129cB9269d658727367ec1e7cC")},
	BlobVerifier:    gethcommon.HexToAddress("0x0fF4367365D7AE16Fa32F1D7ae9827Fb8F602B53"),
	OnRamp:          gethcommon.HexToAddress("0xDE8b0E4d36dFeE71f39a45594617EdFCb2DD53DA"),
	OnRampRouter:    gethcommon.HexToAddress("0xe7606b059c307594fE1DCE50986Fac8Da858E557"),
	OffRamp:         gethcommon.HexToAddress("0xDe35059B8bFfA55604b378109FA994c93B727E2f"),
	OffRampRouter:   gethcommon.HexToAddress("0x61EeaD767B3A2eC057bCc391328901AbFDE17558"),
	MessageReceiver: gethcommon.HexToAddress("0x56304B9d317356F79298DdE661bC48220075e723"),
	ReceiverDapp:    gethcommon.HexToAddress("0x7EB465B48A300288948341bc0D0590Aa1597C5a0"),
	TokenSender:     gethcommon.HexToAddress("0x75506D10Bb15CA81fD9C95603fD082393e6d335d"),
	GovernanceDapp:  gethcommon.HexToAddress(""),
	PingPongDapp:    gethcommon.HexToAddress("0x80aa9e26BddB4d8f2E9CBC46Db3292dbD6b0506E"),
	Afn:             gethcommon.HexToAddress("0x3AaF47808B72d77fC7068DBB2388609A59968910"),
	DeploySettings: DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployBlobVerifier:   true,
		DeployRamp:           true,
		DeployRouter:         true,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   true,
		DeployedAt:           11460200,
	},
}

var GoerliConfig = EvmChainConfig{
	ChainId: big.NewInt(5),
	EthUrl:  secrets.GoerliEthURL,
	GasSettings: EVMGasSettings{
		EIP1559:   true,
		GasTipCap: DefaultGasTipFee,
	},
	LinkToken:       gethcommon.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB"),
	BridgeTokens:    []gethcommon.Address{gethcommon.HexToAddress("0x326C977E6efc84E512bB9C30f76E30c160eD06FB")},
	TokenPools:      []gethcommon.Address{gethcommon.HexToAddress("0x4169578D48aa81129cB9269d658727367ec1e7cC")},
	BlobVerifier:    gethcommon.HexToAddress("0xe7606b059c307594fE1DCE50986Fac8Da858E557"),
	OnRamp:          gethcommon.HexToAddress("0x5856289466c37661b17E717BF03062a3Ca8cF9AB"),
	OnRampRouter:    gethcommon.HexToAddress("0x61EeaD767B3A2eC057bCc391328901AbFDE17558"),
	TokenSender:     gethcommon.HexToAddress("0x804929C12Aa04b1BF435dE068cc8DB85aad65267"),
	OffRamp:         gethcommon.HexToAddress("0xdafF9D7bc634ecE25c3Aea5453Ec1af9b61660F5"),
	OffRampRouter:   gethcommon.HexToAddress("0x55F3c943300859359d6f4c138575aC868be29478"),
	MessageReceiver: gethcommon.HexToAddress("0xDE8b0E4d36dFeE71f39a45594617EdFCb2DD53DA"),
	ReceiverDapp:    gethcommon.HexToAddress("0x0fF4367365D7AE16Fa32F1D7ae9827Fb8F602B53"),
	GovernanceDapp:  gethcommon.HexToAddress(""),
	PingPongDapp:    gethcommon.HexToAddress("0x2681f959bECE0e908A330cfd1be1Dd601866991F"),
	Afn:             gethcommon.HexToAddress("0x3AaF47808B72d77fC7068DBB2388609A59968910"),
	DeploySettings: DeploySettings{
		DeployAFN:            false,
		DeployTokenPools:     false,
		DeployBlobVerifier:   false,
		DeployRamp:           true,
		DeployRouter:         true,
		DeployGovernanceDapp: false,
		DeployPingPongDapp:   true,
		DeployedAt:           7675800,
	},
}

var SepoliaConfig = EvmChainConfig{
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
