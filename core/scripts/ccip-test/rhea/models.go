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

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/secrets"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
)

// DefaultGasTipFee is the default gas tip fee of 1 gwei.
var DefaultGasTipFee = big.NewInt(1e9)

// EVMGasSettings specifies the gas configuration for an EVM chain.
type EVMGasSettings struct {
	EIP1559   bool
	GasPrice  *big.Int
	GasTipCap *big.Int
}

type ChainDeploySettings struct {
	DeployARM           bool
	DeployTokenPools    bool
	DeployRouter        bool
	DeployUpgradeRouter bool
	DeployPriceRegistry bool
	DeployedAtBlock     uint64
}

type LaneDeploySettings struct {
	DeployRamp         bool
	DeployCommitStore  bool
	DeployPingPongDapp bool
	DeployedAtBlock    uint64
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
	PolygonMumbai  Chain = "polygon-testnet-mumbai"
	ArbitrumGoerli Chain = "ethereum-testnet-goerli-arbitrum-1"
)

func GetAllChains() []Chain {
	return []Chain{
		Sepolia, AvaxFuji, OptimismGoerli, Goerli, PolygonMumbai, ArbitrumGoerli,
	}
}

var evmChainIdToChainSelector = map[uint64]uint64{
	11155111: 16015286601757825753,
	420:      2664363617261496610,
	421613:   6101244977088475029,
	43113:    14767482510784806043,
}

func GetCCIPChainSelector(EVMChainId uint64) uint64 {
	selector, ok := evmChainIdToChainSelector[EVMChainId]
	if !ok {
		panic(fmt.Sprintf("no chain selector for %d", EVMChainId))
	}
	return selector
}

type Token string

const (
	LINK      Token = "Link"
	WETH      Token = "WETH"
	WAVAX     Token = "WAVAX"
	CACHEGOLD Token = "CACHE.gold"
	ANZ       Token = "ANZ"
	InsurAce  Token = "InsurAce"
)

func GetAllTokens() []Token {
	return []Token{
		LINK, WETH, WAVAX, CACHEGOLD, ANZ, InsurAce,
	}
}

var tokenSymbols = map[Token]string{
	LINK:      "LINK",
	WETH:      "wETH",
	WAVAX:     "wAVAX",
	CACHEGOLD: "CGT",
	ANZ:       "A$DC",
	InsurAce:  "INSUR",
}

func (token Token) Symbol() string {
	return tokenSymbols[token]
}

var tokenDecimalMultiplier = map[Token]uint8{
	LINK:      18,
	WETH:      18,
	WAVAX:     18,
	CACHEGOLD: 8,
	ANZ:       6,
	InsurAce:  18,
}

func (token Token) Decimals() uint8 {
	return tokenDecimalMultiplier[token]
}

// Price is a mapping from Token to (dollar/1e18) price per wei
// This means a coin that costs $2000 and has 18 decimals precision
// will have a value of 2000e18
func (token Token) Price() *big.Int {
	// Token prices in $ per whole coin
	var TokenPrices = map[Token]*big.Float{
		LINK:      big.NewFloat(10),
		WETH:      big.NewFloat(1500),
		WAVAX:     big.NewFloat(25),
		CACHEGOLD: big.NewFloat(60),
		ANZ:       big.NewFloat(1),
		InsurAce:  big.NewFloat(0.08),
	}

	tokenValue := big.NewInt(0)
	new(big.Float).Mul(TokenPrices[token], big.NewFloat(1e18)).Int(tokenValue)

	// Multiply by 1e18 and divide by the token multiplier so a token with fewer decimals
	// becomes worth more per base unit if the full token price is the same.
	return new(big.Int).Quo(new(big.Int).Mul(tokenValue, big.NewInt(1e18)), token.Multiplier())
}

func (token Token) Multiplier() *big.Int {
	return new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(token.Decimals())), nil)
}

type TokenPoolType string

const (
	LockRelease TokenPoolType = "lockRelease"
	BurnMint    TokenPoolType = "burnMint"
	Wrapped     TokenPoolType = "wrapped"
)

type EVMChainConfig struct {
	EvmChainId  uint64
	GasSettings EVMGasSettings

	SupportedTokens    map[Token]EVMBridgedToken
	FeeTokens          []Token
	WrappedNative      Token
	Router             gethcommon.Address
	UpgradeRouter      gethcommon.Address
	ARM                gethcommon.Address
	PriceRegistry      gethcommon.Address
	AllowList          []gethcommon.Address
	DeploySettings     ChainDeploySettings
	TunableChainValues TunableChainValues
	CustomerSettings
}

type TunableChainValues struct {
	FinalityDepth            uint32
	OptimisticConfirmations  uint32
	BatchGasLimit            uint32
	RelativeBoostPerWaitHour float64
	FeeUpdateHeartBeat       models.Duration
	FeeUpdateDeviationPPB    uint32
	MaxGasPrice              uint64
	InflightCacheExpiry      models.Duration
	RootSnoozeTime           models.Duration
}

type EVMBridgedToken struct {
	ChainId uint64
	Token   gethcommon.Address
	Pool    gethcommon.Address
	TokenPriceType
	Price *big.Int
	PriceFeed
	TokenPoolType
	PoolAllowList []gethcommon.Address // empty slice indicates allowList is not enabled
}

type TokenPriceType string

const (
	TokenPrices TokenPriceType = "TokenPrices"
	PriceFeeds  TokenPriceType = "PriceFeeds"
)

type PriceFeed struct {
	Aggregator gethcommon.Address
	Multiplier *big.Int
}

type EVMLaneConfig struct {
	OnRamp      gethcommon.Address
	OffRamp     gethcommon.Address
	CommitStore gethcommon.Address

	ReceiverDapp   gethcommon.Address
	PingPongDapp   gethcommon.Address
	DeploySettings LaneDeploySettings
}

type EvmDeploymentConfig struct {
	Owner  *bind.TransactOpts
	Client *ethclient.Client
	Logger logger.Logger

	ChainConfig       EVMChainConfig
	LaneConfig        EVMLaneConfig
	UpgradeLaneConfig EVMLaneConfig
}

type EvmConfig struct {
	Owner       *bind.TransactOpts
	Client      *ethclient.Client
	Logger      logger.Logger
	ChainConfig *EVMChainConfig
}

func (chain *EvmDeploymentConfig) OnlyEvmConfig() EvmConfig {
	return EvmConfig{
		Owner:       chain.Owner,
		Client:      chain.Client,
		Logger:      chain.Logger,
		ChainConfig: &chain.ChainConfig,
	}
}

func (chain *EvmDeploymentConfig) SetupChain(t *testing.T, ownerPrivateKey string) {
	chain.Owner = GetOwner(t, ownerPrivateKey, chain.ChainConfig.EvmChainId, chain.ChainConfig.GasSettings)
	chain.Client = GetClient(t, secrets.GetRPC(chain.ChainConfig.EvmChainId))
	chain.Logger = logger.TestLogger(t).Named(ccip.ChainName(int64(chain.ChainConfig.EvmChainId)))
	chain.Logger.Info("Completed chain setup")
}

func (chain *EvmDeploymentConfig) SetupReadOnlyChain(lggr logger.Logger) error {
	client, err := ethclient.Dial(secrets.GetRPC(chain.ChainConfig.EvmChainId))
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
