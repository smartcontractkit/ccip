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
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/secrets"
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
	DeployAFN           bool
	DeployTokenPools    bool
	DeployRouter        bool
	DeployPriceRegistry bool

	DeployRamp         bool
	DeployCommitStore  bool
	DeployPingPongDapp bool
	DeployedAt         uint64
}

type CustomerSettings struct {
	CacheGoldFeeAddress  gethcommon.Address
	CacheGoldFeeEnforcer gethcommon.Address
}

type Chain string

const (
	Sepolia        Chain = "ethereum-testnet-sepolia"
	AvaxFuji       Chain = "avalanche-testnet-fuji"
	OptimismGoerli Chain = "ethereum-testnet-goerli-optimism-1"
	Goerli         Chain = "ethereum-testnet-goerli"
)

func GetAllChains() []Chain {
	return []Chain{
		Sepolia, AvaxFuji, OptimismGoerli, Goerli,
	}
}

type Token string

const (
	LINK      Token = "Link"
	WETH      Token = "WETH"
	WAVAX     Token = "WAVAX"
	CACHEGOLD Token = "CACHE.gold"
)

func GetAllTokens() []Token {
	return []Token{
		LINK, WETH, WAVAX,
	}
}

type TokenPoolType string

const (
	LockRelease TokenPoolType = "lockRelease"
	BurnMint    TokenPoolType = "burnMint"
)

type EVMChainConfig struct {
	ChainId     uint64
	GasSettings EVMGasSettings

	SupportedTokens map[Token]EVMBridgedToken
	FeeTokens       []Token
	WrappedNative   Token
	Router          gethcommon.Address
	Afn             gethcommon.Address
	PriceRegistry   gethcommon.Address
	AllowList       []gethcommon.Address
	Confirmations   uint32
	CustomerSettings
}

type EVMBridgedToken struct {
	Token                gethcommon.Address
	Pool                 gethcommon.Address
	Price                *big.Int
	PriceFeedsAggregator gethcommon.Address
	TokenPoolType
}

type EVMLaneConfig struct {
	OnRamp      gethcommon.Address
	OffRamp     gethcommon.Address
	CommitStore gethcommon.Address

	TokenSender     gethcommon.Address
	MessageReceiver gethcommon.Address
	ReceiverDapp    gethcommon.Address
	GovernanceDapp  gethcommon.Address
	PingPongDapp    gethcommon.Address
}

type EvmDeploymentConfig struct {
	Owner          *bind.TransactOpts
	Client         *ethclient.Client
	Logger         logger.Logger
	DeploySettings DeploySettings

	ChainConfig EVMChainConfig
	LaneConfig  EVMLaneConfig
}

func (chain *EvmDeploymentConfig) SetupChain(t *testing.T, ownerPrivateKey string) {
	chain.Owner = GetOwner(t, ownerPrivateKey, chain.ChainConfig.ChainId, chain.ChainConfig.GasSettings)
	chain.Client = GetClient(t, secrets.GetRPC(chain.ChainConfig.ChainId))
	chain.Logger = logger.TestLogger(t).Named(helpers.ChainName(int64(chain.ChainConfig.ChainId)))
	chain.Logger.Info("Completed chain setup")
}

func (chain *EvmDeploymentConfig) SetupReadOnlyChain(lggr logger.Logger) error {
	client, err := ethclient.Dial(secrets.GetRPC(chain.ChainConfig.ChainId))
	if err != nil {
		return err
	}
	chain.Logger = lggr
	chain.Client = client

	return nil
}

// GetOwner sets the owner user credentials and ensures a GasTipCap is set for the resulting user.
func GetOwner(t *testing.T, ownerPrivateKey string, chainId uint64, gasSettings EVMGasSettings) *bind.TransactOpts {
	ownerKey, err := crypto.HexToECDSA(ownerPrivateKey)
	require.NoError(t, err)
	user, err := bind.NewKeyedTransactorWithChainID(ownerKey, big.NewInt(int64(chainId)))
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
