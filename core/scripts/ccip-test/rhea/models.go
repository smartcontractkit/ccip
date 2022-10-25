package rhea

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/test-go/testify/require"

	"github.com/smartcontractkit/chainlink/core/logger"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
)

// DefaultGasTipFee is the default gas tip fee of 1 gwei.
var DefaultGasTipFee = big.NewInt(1e9)

// EVMGasSettings specifies the gas configuration for an EVM chain.
type EVMGasSettings struct {
	EIP1559   bool
	GasPrice  *big.Int
	GasTipCap *big.Int
}

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

func (chain *EvmChainConfig) SetupChain(t *testing.T, ownerPrivateKey string) {
	chain.Owner = GetOwner(t, ownerPrivateKey, chain.ChainId, chain.GasSettings)
	chain.Client = GetClient(t, chain.EthUrl)
	chain.Logger = logger.TestLogger(t).Named(helpers.ChainName(chain.ChainId.Int64()))

	require.Equal(t, len(chain.BridgeTokens), len(chain.TokenPools))
	chain.Logger.Info("Completed chain setup")
}

// GetOwner sets the owner user credentials and ensures a GasTipCap is set for the resulting user.
func GetOwner(t *testing.T, ownerPrivateKey string, chainId *big.Int, gasSettings EVMGasSettings) *bind.TransactOpts {
	ownerKey, err := crypto.HexToECDSA(ownerPrivateKey)
	require.NoError(t, err)
	user, err := bind.NewKeyedTransactorWithChainID(ownerKey, chainId)
	require.NoError(t, err)
	fmt.Println("--- Owner address ")
	fmt.Println(user.From.Hex())
	SetGasFees(user, gasSettings)

	return user
}

// GetClient dials a given EVM client url and returns the resulting client.
func GetClient(t *testing.T, ethUrl string) *ethclient.Client {
	client, err := ethclient.Dial(ethUrl)
	require.NoError(t, err)
	return client
}

// SetGasFees configures the chain client with the given EVMGasSettings. This method is needed for EIP txs
// to function because of the geth-only tip fee method.
func SetGasFees(owner *bind.TransactOpts, config EVMGasSettings) {
	if config.EIP1559 {
		// to not use geth-only tip fee method when EIP1559 is enabled
		// https://github.com/ethereum/go-ethereum/pull/23484
		owner.GasTipCap = config.GasTipCap
	} else {
		owner.GasPrice = config.GasPrice
	}
}

func PrintContractConfig(source *EvmChainConfig, destination *EvmChainConfig) {
	source.Logger.Infof(`
Source chain config

LinkToken:      common.HexToAddress("%s"),
BridgeTokens:   %s,
TokenPools:     %s,
OnRamp:         common.HexToAddress("%s"),
OnRampRouter:   common.HexToAddress("%s"),
TokenSender:    common.HexToAddress("%s"),
Afn:            common.HexToAddress("%s"),
GovernanceDapp: common.HexToAddress("%s"),
PingPongDapp:   common.HexToAddress("%s"),
	
`,
		source.LinkToken,
		source.BridgeTokens,
		source.TokenPools,
		source.OnRamp,
		source.OnRampRouter,
		source.TokenSender,
		source.Afn,
		source.GovernanceDapp,
		source.PingPongDapp)

	destination.Logger.Infof(`
Destination chain config

LinkToken:       common.HexToAddress("%s"),
BridgeTokens:    %s,
TokenPools:      %s,
OffRamp:         common.HexToAddress("%s"),
OffRampRouter:   common.HexToAddress("%s"),
BlobVerifier:    common.HexToAddress("%s"),	
MessageReceiver: common.HexToAddress("%s"),
ReceiverDapp:    common.HexToAddress("%s"),
Afn:             common.HexToAddress("%s"),
GovernanceDapp:  common.HexToAddress("%s"),
PingPongDapp:    common.HexToAddress("%s"),
`,
		destination.LinkToken,
		destination.BridgeTokens,
		destination.TokenPools,
		destination.OffRamp,
		destination.OffRampRouter,
		destination.BlobVerifier,
		destination.MessageReceiver,
		destination.ReceiverDapp,
		destination.Afn,
		destination.GovernanceDapp,
		destination.PingPongDapp)
}
