package testhelpers

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/onsi/gomega"
	"github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_toll_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/custom_token_pool"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_toll_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_ge_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_ge_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_toll_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/fee_manager"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/ge_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/maybe_revert_message_receiver"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/mock_afn_contract"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/merklemulti"
	"github.com/smartcontractkit/chainlink/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/core/utils"
)

var (
	// Source
	SourcePool       = "source pool"
	TollOnRampRouter = "toll onramp router"
	TollOnRamp       = "toll onramp"
	GEOnRamp         = "ge onramp"
	SourceGERouter   = "source ge router"

	// Dest
	TollOffRampRouter = "toll offramp router"
	TollOffRamp       = "toll offramp"
	GEOffRamp         = "ge offramp"
	DestGERouter      = "dest ge router"
	DestPool          = "dest pool"
	Receiver          = "receiver"
	Sender            = "sender"
	Link              = func(amount int64) *big.Int { return new(big.Int).Mul(big.NewInt(1e18), big.NewInt(amount)) }
	HundredLink       = Link(100)
)

type MaybeRevertReceiver struct {
	Receiver *maybe_revert_message_receiver.MaybeRevertMessageReceiver
	Strict   bool
}

type CCIPContracts struct {
	t         *testing.T
	Source    SourceChain
	Dest      DestinationChain
	OCRConfig *OCR2Config
}

type Common struct {
	ChainID     uint64
	User        *bind.TransactOpts
	Chain       *backends.SimulatedBackend
	LinkToken   *link_token_interface.LinkToken
	Pool        *native_token_pool.NativeTokenPool
	CustomPool  *custom_token_pool.CustomTokenPool
	CustomToken *link_token_interface.LinkToken
	AFN         *mock_afn_contract.MockAFNContract
}

type SourceChain struct {
	Common

	// Toll contracts
	TollOnRampFees   map[common.Address]*big.Int
	TollOnRampRouter *evm_2_any_toll_onramp_router.EVM2AnyTollOnRampRouter
	TollOnRamp       *evm_2_evm_toll_onramp.EVM2EVMTollOnRamp

	// GE
	GERouter *ge_router.GERouter
	GEOnRamp *evm_2_evm_ge_onramp.EVM2EVMGEOnRamp
}

type DestinationChain struct {
	Common

	CommitStore *commit_store.CommitStore
	// Toll contracts
	TollOffRampRouter *any_2_evm_toll_offramp_router.Any2EVMTollOffRampRouter
	TollOffRamp       *evm_2_evm_toll_offramp.EVM2EVMTollOffRamp

	// GE
	GERouter  *ge_router.GERouter
	GEOffRamp *evm_2_evm_ge_offramp.EVM2EVMGEOffRamp
	Receivers []MaybeRevertReceiver
}

type OCR2Config struct {
	Signers               []common.Address
	Transmitters          []common.Address
	F                     uint8
	OnchainConfig         []byte
	OffchainConfigVersion uint64
	OffchainConfig        []byte
}

type BalanceAssertion struct {
	Name     string
	Address  common.Address
	Expected string
	Getter   func(addr common.Address) *big.Int
	Within   string
}

type BalanceReq struct {
	Name   string
	Addr   common.Address
	Getter func(addr common.Address) *big.Int
}

func (c *CCIPContracts) GetSourceLinkBalance(addr common.Address) *big.Int {
	bal, err := c.Source.LinkToken.BalanceOf(nil, addr)
	require.NoError(c.t, err)
	return bal
}

func (c *CCIPContracts) GetDestLinkBalance(addr common.Address) *big.Int {
	bal, err := c.Dest.LinkToken.BalanceOf(nil, addr)
	require.NoError(c.t, err)
	return bal
}

func (c *CCIPContracts) AssertBalances(bas []BalanceAssertion) {
	for _, b := range bas {
		actual := b.Getter(b.Address)
		require.NotNil(c.t, actual, "%v getter return nil", b.Name)
		if b.Within == "" {
			assert.Equal(c.t, b.Expected, actual.String(), "wrong balance for %s got %s want %s", b.Name, actual, b.Expected)
		} else {
			bi, _ := big.NewInt(0).SetString(b.Expected, 10)
			withinI, _ := big.NewInt(0).SetString(b.Within, 10)
			high := big.NewInt(0).Add(bi, withinI)
			low := big.NewInt(0).Sub(bi, withinI)
			assert.Equal(c.t, -1, actual.Cmp(high), "wrong balance for %s got %s outside expected range [%s, %s]", b.Name, actual, low, high)
			assert.Equal(c.t, 1, actual.Cmp(low), "wrong balance for %s got %s outside expected range [%s, %s]", b.Name, actual, low, high)
		}
	}
}

func (c *CCIPContracts) DeployNewTollOffRamp() {
	offRampAddress, _, _, err := evm_2_evm_toll_offramp.DeployEVM2EVMTollOffRamp(
		c.Dest.User,
		c.Dest.Chain,
		c.Source.ChainID,
		c.Dest.ChainID,
		evm_2_evm_toll_offramp.IBaseOffRampOffRampConfig{
			ExecutionDelaySeconds:                   60,
			MaxDataSize:                             1e5,
			MaxTokensLength:                         15,
			PermissionLessExecutionThresholdSeconds: 60,
		},
		c.Source.TollOnRamp.Address(),
		c.Dest.CommitStore.Address(),
		c.Dest.AFN.Address(),
		[]common.Address{c.Source.LinkToken.Address()},
		[]common.Address{c.Dest.Pool.Address()},
		evm_2_evm_toll_offramp.IAggregateRateLimiterRateLimiterConfig{
			Capacity: HundredLink,
			Rate:     big.NewInt(1e18),
		},
		c.Dest.User.From,
	)
	require.NoError(c.t, err)
	c.Dest.Chain.Commit()

	c.Dest.TollOffRamp, err = evm_2_evm_toll_offramp.NewEVM2EVMTollOffRamp(offRampAddress, c.Dest.Chain)
	require.NoError(c.t, err)

	_, err = c.Dest.TollOffRamp.SetPrices(c.Dest.User, []common.Address{c.Dest.LinkToken.Address()}, []*big.Int{big.NewInt(1)})
	require.NoError(c.t, err)
	c.Dest.Chain.Commit()
	c.Source.Chain.Commit()
}

func (c *CCIPContracts) EnableTollOffRamp() {
	_, err := c.Dest.Pool.SetOffRamp(c.Dest.User, c.Dest.TollOffRamp.Address(), true)
	require.NoError(c.t, err)
	c.Dest.Chain.Commit()

	_, err = c.Dest.TollOffRamp.SetRouter(c.Dest.User, c.Dest.TollOffRampRouter.Address())
	require.NoError(c.t, err)
	c.Dest.Chain.Commit()

	_, err = c.Dest.TollOffRampRouter.AddOffRamp(c.Dest.User, c.Dest.TollOffRamp.Address())
	require.NoError(c.t, err)
	c.Dest.Chain.Commit()

	_, err = c.Dest.TollOffRamp.SetOCR2Config(
		c.Dest.User,
		c.OCRConfig.Signers,
		c.OCRConfig.Transmitters,
		c.OCRConfig.F,
		c.OCRConfig.OnchainConfig,
		c.OCRConfig.OffchainConfigVersion,
		c.OCRConfig.OffchainConfig,
	)
	require.NoError(c.t, err)
	c.Source.Chain.Commit()
	c.Dest.Chain.Commit()
}

func (c *CCIPContracts) DeployNewTollOnRamp() {
	c.t.Log("Deploying new toll onRamp")
	onRampAddress, _, _, err := evm_2_evm_toll_onramp.DeployEVM2EVMTollOnRamp(
		c.Source.User,    // user
		c.Source.Chain,   // client
		c.Source.ChainID, // source chain id
		c.Dest.ChainID,   // destinationChainIds
		[]common.Address{c.Source.LinkToken.Address()}, // tokens
		[]common.Address{c.Source.Pool.Address()},      // pools
		[]common.Address{},                             // allow list
		c.Source.AFN.Address(),                         // AFN
		evm_2_evm_toll_onramp.IBaseOnRampOnRampConfig{
			CommitFeeJuels:  0,
			MaxDataSize:     1e12,
			MaxTokensLength: 5,
			MaxGasLimit:     ccip.GasLimitPerTx,
		},
		evm_2_evm_toll_onramp.IAggregateRateLimiterRateLimiterConfig{
			Capacity: HundredLink,
			Rate:     big.NewInt(1e18),
		},
		c.Source.User.From,
		c.Source.TollOnRampRouter.Address(),
	)
	require.NoError(c.t, err)
	c.Source.TollOnRamp, err = evm_2_evm_toll_onramp.NewEVM2EVMTollOnRamp(onRampAddress, c.Source.Chain)
	require.NoError(c.t, err)
	c.Source.Chain.Commit()

	_, err = c.Source.TollOnRamp.SetFeeConfig(c.Source.User, evm_2_evm_toll_onramp.IEVM2EVMTollOnRampFeeConfig{
		Fees:      []*big.Int{big.NewInt(1)},
		FeeTokens: []common.Address{c.Source.LinkToken.Address()},
	})
	require.NoError(c.t, err)

	_, err = c.Source.TollOnRamp.SetPrices(c.Source.User, []common.Address{c.Source.LinkToken.Address()}, []*big.Int{big.NewInt(1)})
	require.NoError(c.t, err)

	c.Source.Chain.Commit()
	c.Dest.Chain.Commit()
}

func (c *CCIPContracts) EnableTollOnRamp() {
	c.t.Log("Setting toll onRamp on source pool")
	_, err := c.Source.Pool.SetOnRamp(c.Source.User, c.Source.TollOnRamp.Address(), true)
	require.NoError(c.t, err)
	c.Source.Chain.Commit()

	c.t.Log("Setting toll onRamp on source router")
	_, err = c.Source.TollOnRampRouter.SetOnRamp(c.Source.User, c.Dest.ChainID, c.Source.TollOnRamp.Address())
	require.NoError(c.t, err)
	c.Source.Chain.Commit()

	c.t.Log("Enabling toll onRamp on commitStore")
	config, err := c.Dest.CommitStore.GetCommitStoreConfig(&bind.CallOpts{})
	require.NoError(c.t, err)

	config.OnRamps = append(config.OnRamps, c.Source.TollOnRamp.Address())
	config.MinSeqNrByOnRamp = append(config.MinSeqNrByOnRamp, 1)

	_, err = c.Dest.CommitStore.SetCommitStoreConfig(c.Dest.User, config)
	require.NoError(c.t, err)

	c.Source.Chain.Commit()
	c.Dest.Chain.Commit()
}

func (c *CCIPContracts) DeriveOCR2Config(oracles []confighelper.OracleIdentityExtra, reportingPluginConfig []byte) {
	signers, transmitters, threshold, onchainConfig, offchainConfigVersion, offchainConfig, err := confighelper.ContractSetConfigArgsForTests(
		2*time.Second,        // deltaProgress
		1*time.Second,        // deltaResend
		1*time.Second,        // deltaRound
		500*time.Millisecond, // deltaGrace
		2*time.Second,        // deltaStage
		3,
		[]int{1, 1, 1, 1},
		oracles,
		reportingPluginConfig,
		50*time.Millisecond, // Max duration query
		1*time.Second,       // Max duration observation
		100*time.Millisecond,
		100*time.Millisecond,
		100*time.Millisecond,
		1, // faults
		nil,
	)
	require.NoError(c.t, err)
	lggr := logger.TestLogger(c.t)
	lggr.Infow("Setting Config on Oracle Contract",
		"signers", signers,
		"transmitters", transmitters,
		"threshold", threshold,
		"onchainConfig", onchainConfig,
		"encodedConfigVersion", offchainConfigVersion,
	)
	signerAddresses, err := ocrcommon.OnchainPublicKeyToAddress(signers)
	require.NoError(c.t, err)
	transmitterAddresses, err := ocrcommon.AccountToAddress(transmitters)
	require.NoError(c.t, err)

	c.OCRConfig = &OCR2Config{
		Signers:               signerAddresses,
		Transmitters:          transmitterAddresses,
		F:                     threshold,
		OnchainConfig:         onchainConfig,
		OffchainConfigVersion: offchainConfigVersion,
		OffchainConfig:        offchainConfig,
	}
}

func (c *CCIPContracts) SetupOnchainConfig(oracles []confighelper.OracleIdentityExtra, reportingPluginConfig []byte) int64 {
	// Note We do NOT set the payees, payment is done in the OCR2Base implementation
	// Set the offramp offchainConfig.
	c.DeriveOCR2Config(oracles, reportingPluginConfig)
	blockBeforeConfig, err := c.Dest.Chain.BlockByNumber(context.Background(), nil)
	require.NoError(c.t, err)
	// Set the DON on the offramp
	_, err = c.Dest.CommitStore.SetOCR2Config(
		c.Dest.User,
		c.OCRConfig.Signers,
		c.OCRConfig.Transmitters,
		c.OCRConfig.F,
		c.OCRConfig.OnchainConfig,
		c.OCRConfig.OffchainConfigVersion,
		c.OCRConfig.OffchainConfig,
	)
	require.NoError(c.t, err)
	c.Dest.Chain.Commit()

	// Same DON on the toll offramp
	_, err = c.Dest.TollOffRamp.SetOCR2Config(
		c.Dest.User,
		c.OCRConfig.Signers,
		c.OCRConfig.Transmitters,
		c.OCRConfig.F,
		c.OCRConfig.OnchainConfig,
		c.OCRConfig.OffchainConfigVersion,
		c.OCRConfig.OffchainConfig,
	)
	require.NoError(c.t, err)
	c.Dest.Chain.Commit()
	// Same DON on the GE offramp
	_, err = c.Dest.GEOffRamp.SetOCR2Config(
		c.Dest.User,
		c.OCRConfig.Signers,
		c.OCRConfig.Transmitters,
		c.OCRConfig.F,
		c.OCRConfig.OnchainConfig,
		c.OCRConfig.OffchainConfigVersion,
		c.OCRConfig.OffchainConfig,
	)
	require.NoError(c.t, err)
	c.Dest.Chain.Commit()

	return blockBeforeConfig.Number().Int64()
}

func (c *CCIPContracts) NewCCIPJobSpecParams(tokensPerFeeCoinPipeline string, configBlock int64) CCIPJobSpecParams {
	return CCIPJobSpecParams{
		OnRampsOnCommit:          []common.Address{c.Source.TollOnRamp.Address(), c.Source.GEOnRamp.Address()},
		CommitStore:              c.Dest.CommitStore.Address(),
		SourceChainId:            c.Source.ChainID,
		DestChainId:              c.Dest.ChainID,
		SourceChainName:          "SimulatedSource",
		DestChainName:            "SimulatedDest",
		TokensPerFeeCoinPipeline: tokensPerFeeCoinPipeline,
		PollPeriod:               time.Second,
		DestStartBlock:           uint64(configBlock),
	}
}

func GetBalances(brs []BalanceReq) (map[string]*big.Int, error) {
	m := make(map[string]*big.Int)
	for _, br := range brs {
		m[br.Name] = br.Getter(br.Addr)
		if m[br.Name] == nil {
			return nil, fmt.Errorf("%v getter return nil", br.Name)
		}
	}
	return m, nil
}

func MustAddBigInt(a *big.Int, b string) *big.Int {
	bi, _ := big.NewInt(0).SetString(b, 10)
	return big.NewInt(0).Add(a, bi)
}

func MustSubBigInt(a *big.Int, b string) *big.Int {
	bi, _ := big.NewInt(0).SetString(b, 10)
	return big.NewInt(0).Sub(a, bi)
}

func MustEncodeAddress(t *testing.T, address common.Address) []byte {
	bts, err := utils.ABIEncode(`[{"type":"address"}]`, address)
	require.NoError(t, err)
	return bts
}

func SetupCCIPContracts(t *testing.T, sourceChainID, destChainID uint64) CCIPContracts {
	sourceChain, sourceUser := SetupChain(t)
	destChain, destUser := SetupChain(t)

	// Deploy link token and pool on source chain
	sourceLinkTokenAddress, _, _, err := link_token_interface.DeployLinkToken(sourceUser, sourceChain)
	require.NoError(t, err)
	sourceChain.Commit()
	sourceLinkToken, err := link_token_interface.NewLinkToken(sourceLinkTokenAddress, sourceChain)
	require.NoError(t, err)
	sourcePoolAddress, _, _, err := native_token_pool.DeployNativeTokenPool(sourceUser,
		sourceChain,
		sourceLinkTokenAddress)
	require.NoError(t, err)
	sourceChain.Commit()
	sourcePool, err := native_token_pool.NewNativeTokenPool(sourcePoolAddress, sourceChain)
	require.NoError(t, err)

	// Deploy link token and pool on destination chain
	destLinkTokenAddress, _, _, err := link_token_interface.DeployLinkToken(destUser, destChain)
	require.NoError(t, err)
	destChain.Commit()
	destLinkToken, err := link_token_interface.NewLinkToken(destLinkTokenAddress, destChain)
	require.NoError(t, err)
	destPoolAddress, _, _, err := native_token_pool.DeployNativeTokenPool(destUser, destChain, destLinkTokenAddress)
	require.NoError(t, err)
	destChain.Commit()
	destPool, err := native_token_pool.NewNativeTokenPool(destPoolAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()

	// Float the offramp pool
	o, err := destPool.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, destUser.From.String(), o.String())
	_, err = destLinkToken.Transfer(destUser, destPoolAddress, Link(200))
	require.NoError(t, err)
	destChain.Commit()

	// Deploy custom token pool source
	sourceCustomTokenAddress, _, _, err := link_token_interface.DeployLinkToken(sourceUser, sourceChain) // Just re-use this, its an ERC20.
	require.NoError(t, err)
	sourceCustomToken, err := link_token_interface.NewLinkToken(sourceCustomTokenAddress, sourceChain)
	require.NoError(t, err)
	destChain.Commit()

	// Deploy custom token pool dest
	destCustomTokenAddress, _, _, err := link_token_interface.DeployLinkToken(destUser, destChain) // Just re-use this, its an ERC20.
	require.NoError(t, err)
	destCustomToken, err := link_token_interface.NewLinkToken(destCustomTokenAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()

	afnSourceAddress, _, _, err := mock_afn_contract.DeployMockAFNContract(
		sourceUser,
		sourceChain,
	)
	require.NoError(t, err)
	sourceChain.Commit()
	sourceAFN, err := mock_afn_contract.NewMockAFNContract(afnSourceAddress, sourceChain)
	require.NoError(t, err)

	// Create toll onramp router
	onRampRouterAddress, _, _, err := evm_2_any_toll_onramp_router.DeployEVM2AnyTollOnRampRouter(sourceUser, sourceChain)
	require.NoError(t, err)
	sourceChain.Commit()

	// Create ge router
	sourceGERouterAddress, _, _, err := ge_router.DeployGERouter(sourceUser, sourceChain, []common.Address{})
	require.NoError(t, err)
	sourceGERouter, err := ge_router.NewGERouter(sourceGERouterAddress, sourceChain)
	require.NoError(t, err)
	sourceChain.Commit()

	// Deploy and configure toll onramp.
	onRampAddress, _, _, err := evm_2_evm_toll_onramp.DeployEVM2EVMTollOnRamp(
		sourceUser,                               // user
		sourceChain,                              // client
		sourceChainID,                            // source chain id
		destChainID,                              // destinationChainIds
		[]common.Address{sourceLinkTokenAddress}, // tokens
		[]common.Address{sourcePoolAddress},      // pools
		[]common.Address{},                       // allow list
		afnSourceAddress,                         // AFN
		evm_2_evm_toll_onramp.IBaseOnRampOnRampConfig{
			CommitFeeJuels:  0,
			MaxDataSize:     1e12,
			MaxTokensLength: 5,
			MaxGasLimit:     ccip.GasLimitPerTx,
		},
		evm_2_evm_toll_onramp.IAggregateRateLimiterRateLimiterConfig{
			Capacity: HundredLink,
			Rate:     big.NewInt(1e18),
		},
		sourceUser.From,
		onRampRouterAddress,
	)
	require.NoError(t, err)
	tollOnRamp, err := evm_2_evm_toll_onramp.NewEVM2EVMTollOnRamp(onRampAddress, sourceChain)
	require.NoError(t, err)
	tollOnRampFees := map[common.Address]*big.Int{
		sourceLinkTokenAddress: big.NewInt(1), // 1 juel commit fee.
	}
	_, err = tollOnRamp.SetFeeConfig(sourceUser, evm_2_evm_toll_onramp.IEVM2EVMTollOnRampFeeConfig{
		Fees:      []*big.Int{tollOnRampFees[sourceLinkTokenAddress]},
		FeeTokens: []common.Address{sourceLinkTokenAddress},
	})
	require.NoError(t, err)
	_, err = sourcePool.SetOnRamp(sourceUser, onRampAddress, true)
	require.NoError(t, err)
	sourceChain.Commit()
	_, err = tollOnRamp.SetPrices(sourceUser, []common.Address{sourceLinkTokenAddress}, []*big.Int{big.NewInt(1)})
	require.NoError(t, err)
	tollOnRampRouter, err := evm_2_any_toll_onramp_router.NewEVM2AnyTollOnRampRouter(onRampRouterAddress, sourceChain)
	require.NoError(t, err)
	_, err = tollOnRampRouter.SetOnRamp(sourceUser, destChainID, onRampAddress)
	require.NoError(t, err)
	sourceChain.Commit()

	// Deploy and configure GE onramp
	sourceFeeManagerAddress, _, _, err := fee_manager.DeployFeeManager(sourceUser, sourceChain, []fee_manager.GEFeeUpdate{
		{
			Token:          sourceLinkTokenAddress,
			ChainId:        destChainID,
			LinkPerUnitGas: big.NewInt(1e9), // 1 gwei
		},
	}, nil, big.NewInt(1e18))
	require.NoError(t, err)
	geOnRampAddress, _, _, err := evm_2_evm_ge_onramp.DeployEVM2EVMGEOnRamp(
		sourceUser,                               // user
		sourceChain,                              // client
		sourceChainID,                            // source chain id
		destChainID,                              // destinationChainIds
		[]common.Address{sourceLinkTokenAddress}, // tokens
		[]common.Address{sourcePoolAddress},      // pools
		[]common.Address{},                       // allow list
		afnSourceAddress,                         // AFN
		evm_2_evm_ge_onramp.IBaseOnRampOnRampConfig{
			CommitFeeJuels:  0,
			MaxDataSize:     1e12,
			MaxTokensLength: 5,
			MaxGasLimit:     ccip.GasLimitPerTx,
		},
		evm_2_evm_ge_onramp.IAggregateRateLimiterRateLimiterConfig{
			Capacity: HundredLink,
			Rate:     big.NewInt(1e18),
		},
		sourceUser.From,
		sourceGERouterAddress,
		evm_2_evm_ge_onramp.IEVM2EVMGEOnRampDynamicFeeConfig{
			LinkToken:       sourceLinkTokenAddress,
			FeeAmount:       big.NewInt(0),
			DestGasOverhead: big.NewInt(0),
			Multiplier:      big.NewInt(1e18),
			FeeManager:      sourceFeeManagerAddress,
			DestChainId:     destChainID,
		},
	)
	require.NoError(t, err)
	geOnRamp, err := evm_2_evm_ge_onramp.NewEVM2EVMGEOnRamp(geOnRampAddress, sourceChain)
	require.NoError(t, err)
	_, err = sourcePool.SetOnRamp(sourceUser, geOnRampAddress, true)
	require.NoError(t, err)
	sourceChain.Commit()
	_, err = geOnRamp.SetPrices(sourceUser, []common.Address{sourceLinkTokenAddress}, []*big.Int{big.NewInt(1)})
	require.NoError(t, err)
	_, err = sourceGERouter.SetOnRamp(sourceUser, destChainID, geOnRampAddress)
	require.NoError(t, err)
	sourceChain.Commit()

	afnDestAddress, _, _, err := mock_afn_contract.DeployMockAFNContract(
		destUser,
		destChain,
	)
	require.NoError(t, err)
	destChain.Commit()
	destAFN, err := mock_afn_contract.NewMockAFNContract(afnDestAddress, destChain)
	require.NoError(t, err)

	// Deploy commit store.
	commitStoreAddress, _, _, err := commit_store.DeployCommitStore(
		destUser,    // user
		destChain,   // client
		destChainID, // dest chain id
		sourceChainID,
		afnDestAddress, // AFN address
		commit_store.ICommitStoreCommitStoreConfig{
			OnRamps:          []common.Address{tollOnRamp.Address(), geOnRamp.Address()},
			MinSeqNrByOnRamp: []uint64{1, 1},
		},
	)
	require.NoError(t, err)
	commitStore, err := commit_store.NewCommitStore(commitStoreAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()

	// Deploy and configure toll offramp.
	tollOffRampAddress, _, _, err := evm_2_evm_toll_offramp.DeployEVM2EVMTollOffRamp(
		destUser,
		destChain,
		sourceChainID,
		destChainID,
		evm_2_evm_toll_offramp.IBaseOffRampOffRampConfig{
			ExecutionDelaySeconds: 0,
			MaxDataSize:           1e12,
			MaxTokensLength:       5,
		},
		onRampAddress,
		commitStore.Address(),
		afnDestAddress,
		[]common.Address{sourceLinkTokenAddress},
		[]common.Address{destPoolAddress},
		evm_2_evm_toll_offramp.IAggregateRateLimiterRateLimiterConfig{
			Capacity: HundredLink,
			Rate:     big.NewInt(1e18),
		},
		sourceUser.From,
	)
	require.NoError(t, err)
	tollOffRamp, err := evm_2_evm_toll_offramp.NewEVM2EVMTollOffRamp(tollOffRampAddress, destChain)
	require.NoError(t, err)
	_, err = destPool.SetOffRamp(destUser, tollOffRampAddress, true)
	require.NoError(t, err)
	// Create offRampAddr router
	offRampRouterAddress, _, _, err := any_2_evm_toll_offramp_router.DeployAny2EVMTollOffRampRouter(destUser, destChain, []common.Address{tollOffRampAddress})
	require.NoError(t, err)
	destChain.Commit()
	tollOffRampRouter, err := any_2_evm_toll_offramp_router.NewAny2EVMTollOffRampRouter(offRampRouterAddress, destChain)
	require.NoError(t, err)
	_, err = tollOffRamp.SetRouter(destUser, offRampRouterAddress)
	require.NoError(t, err)
	_, err = tollOffRamp.SetPrices(destUser, []common.Address{destLinkTokenAddress}, []*big.Int{big.NewInt(1)})
	require.NoError(t, err)

	// Deploy and configure ge offramp.
	destFeeManagerAddress, _, _, err := fee_manager.DeployFeeManager(destUser, destChain, []fee_manager.GEFeeUpdate{{
		Token:          destLinkTokenAddress,
		ChainId:        sourceChainID,
		LinkPerUnitGas: big.NewInt(200e9), // (2e20 juels/eth) * (1 gwei / gas) / (1 eth/1e18)
	}}, nil, big.NewInt(1e18))
	require.NoError(t, err)
	destFeeManager, err := fee_manager.NewFeeManager(destFeeManagerAddress, destChain)
	require.NoError(t, err)
	geOffRampAddress, _, _, err := evm_2_evm_ge_offramp.DeployEVM2EVMGEOffRamp(
		destUser,
		destChain,
		sourceChainID,
		destChainID,
		evm_2_evm_ge_offramp.IEVM2EVMGEOffRampGEOffRampConfig{
			GasOverhead:                             big.NewInt(0),
			FeeManager:                              destFeeManagerAddress,
			PermissionLessExecutionThresholdSeconds: 1,
			ExecutionDelaySeconds:                   0,
			MaxDataSize:                             1e12,
			MaxTokensLength:                         5,
		},
		geOnRampAddress,
		commitStore.Address(),
		afnDestAddress,
		[]common.Address{sourceLinkTokenAddress},
		[]common.Address{destPoolAddress},
		evm_2_evm_ge_offramp.IAggregateRateLimiterRateLimiterConfig{
			Capacity: HundredLink,
			Rate:     big.NewInt(1e18),
		},
		sourceUser.From,
		destLinkTokenAddress,
	)
	require.NoError(t, err)
	geOffRamp, err := evm_2_evm_ge_offramp.NewEVM2EVMGEOffRamp(geOffRampAddress, destChain)
	require.NoError(t, err)
	_, err = destPool.SetOffRamp(destUser, geOffRampAddress, true)
	require.NoError(t, err)
	destChain.Commit()
	// OffRamp can update
	_, err = destFeeManager.SetFeeUpdater(destUser, geOffRampAddress)
	require.NoError(t, err)

	// Create dest ge router
	destGERouterAddress, _, _, err := ge_router.DeployGERouter(destUser, destChain, []common.Address{geOffRampAddress})
	require.NoError(t, err)
	destChain.Commit()
	destGERouter, err := ge_router.NewGERouter(destGERouterAddress, destChain)
	require.NoError(t, err)
	_, err = geOffRamp.SetRouter(destUser, destGERouterAddress)
	require.NoError(t, err)
	_, err = geOffRamp.SetPrices(destUser, []common.Address{destLinkTokenAddress}, []*big.Int{big.NewInt(1)})
	require.NoError(t, err)

	// Deploy 2 revertable (one SS one non-SS)
	revertingMessageReceiver1Address, _, _, err := maybe_revert_message_receiver.DeployMaybeRevertMessageReceiver(destUser, destChain, false)
	require.NoError(t, err)
	revertingMessageReceiver1, _ := maybe_revert_message_receiver.NewMaybeRevertMessageReceiver(revertingMessageReceiver1Address, destChain)
	revertingMessageReceiver2Address, _, _, err := maybe_revert_message_receiver.DeployMaybeRevertMessageReceiver(destUser, destChain, false)
	require.NoError(t, err)
	revertingMessageReceiver2, _ := maybe_revert_message_receiver.NewMaybeRevertMessageReceiver(revertingMessageReceiver2Address, destChain)
	// Need to commit here, or we will hit the block gas limit when deploying the executor
	sourceChain.Commit()
	destChain.Commit()

	// Ensure we have at least finality blocks.
	for i := 0; i < 50; i++ {
		sourceChain.Commit()
		destChain.Commit()
	}

	source := SourceChain{
		Common: Common{
			ChainID:     sourceChainID,
			User:        sourceUser,
			Chain:       sourceChain,
			LinkToken:   sourceLinkToken,
			Pool:        sourcePool,
			CustomPool:  nil,
			CustomToken: sourceCustomToken,
			AFN:         sourceAFN,
		},
		TollOnRampFees:   tollOnRampFees,
		TollOnRampRouter: tollOnRampRouter,
		TollOnRamp:       tollOnRamp,
		GERouter:         sourceGERouter,
		GEOnRamp:         geOnRamp,
	}
	dest := DestinationChain{
		Common: Common{
			ChainID:     destChainID,
			User:        destUser,
			Chain:       destChain,
			LinkToken:   destLinkToken,
			Pool:        destPool,
			CustomPool:  nil,
			CustomToken: destCustomToken,
			AFN:         destAFN,
		},
		CommitStore:       commitStore,
		TollOffRampRouter: tollOffRampRouter,
		TollOffRamp:       tollOffRamp,
		GERouter:          destGERouter,
		GEOffRamp:         geOffRamp,
		Receivers:         []MaybeRevertReceiver{{Receiver: revertingMessageReceiver1, Strict: false}, {Receiver: revertingMessageReceiver2, Strict: true}},
	}

	return CCIPContracts{
		t:      t,
		Source: source,
		Dest:   dest,
	}
}

func SendGERequest(t *testing.T, ccipContracts CCIPContracts, msg ge_router.GEConsumerEVM2AnyGEMessage) {
	tx, err := ccipContracts.Source.GERouter.CcipSend(ccipContracts.Source.User, ccipContracts.Dest.ChainID, msg)
	require.NoError(t, err)
	ConfirmTxs(t, []*types.Transaction{tx}, ccipContracts.Source.Chain)
}

func AssertGEExecState(t *testing.T, ccipContracts CCIPContracts, log logpoller.Log, state ccip.MessageExecutionState) {
	executionStateChanged, err := ccipContracts.Dest.GEOffRamp.ParseExecutionStateChanged(log.GetGethLog())
	require.NoError(t, err)
	if ccip.MessageExecutionState(executionStateChanged.State) != state {
		t.Log("Execution failed")
		t.Fail()
	}
}

func EventuallyExecutionStateChangedToSuccess(t *testing.T, ccipContracts CCIPContracts, seqNum []uint64, blockNum uint64) {
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		it, err := ccipContracts.Dest.GEOffRamp.FilterExecutionStateChanged(&bind.FilterOpts{Start: blockNum}, seqNum, [][32]byte{})
		require.NoError(t, err)
		for it.Next() {
			if ccip.MessageExecutionState(it.Event.State) == ccip.Success {
				return true
			}
		}
		ccipContracts.Source.Chain.Commit()
		ccipContracts.Dest.Chain.Commit()
		return false
	}, 1*time.Minute, time.Second).
		Should(gomega.BeTrue(), "ExecutionStateChanged Event")
}

func QueueTollRequest(
	t *testing.T,
	ccipContracts CCIPContracts,
	msgPayload string,
	tokens []evm_2_any_toll_onramp_router.CommonEVMTokenAndAmount,
	feeToken evm_2_any_toll_onramp_router.CommonEVMTokenAndAmount,
	gasLimit *big.Int,
	receiver common.Address,
) *types.Transaction {
	extraArgs, err := GetEVMExtraArgsV1(gasLimit, false)
	require.NoError(t, err)
	msg := evm_2_any_toll_onramp_router.TollConsumerEVM2AnyTollMessage{
		Receiver:          MustEncodeAddress(t, receiver),
		Data:              []byte(msgPayload),
		TokensAndAmounts:  tokens,
		FeeTokenAndAmount: feeToken,
		ExtraArgs:         extraArgs,
	}
	tx, err := ccipContracts.Source.TollOnRampRouter.CcipSend(ccipContracts.Source.User, ccipContracts.Dest.ChainID, msg)
	require.NoError(t, err)
	return tx
}

func SendTollRequest(t *testing.T, ccipContracts CCIPContracts, msgPayload string, tokens []evm_2_any_toll_onramp_router.CommonEVMTokenAndAmount, feeToken evm_2_any_toll_onramp_router.CommonEVMTokenAndAmount, gasLimit *big.Int, receiver common.Address) {
	tx := QueueTollRequest(t, ccipContracts, msgPayload, tokens, feeToken, gasLimit, receiver)
	ConfirmTxs(t, []*types.Transaction{tx}, ccipContracts.Source.Chain)
}

func AssertTollExecSuccess(t *testing.T, ccipContracts CCIPContracts, log logpoller.Log) {
	executionStateChanged, err := ccipContracts.Dest.TollOffRamp.ParseExecutionStateChanged(log.GetGethLog())
	require.NoError(t, err)
	if ccip.MessageExecutionState(executionStateChanged.State) != ccip.Success {
		t.Log("Execution failed")
		t.Fail()
	}
}

func EventuallyReportCommitted(t *testing.T, ccipContracts CCIPContracts, onRamp common.Address, min, max int) {
	gomega.NewGomegaWithT(t).Eventually(func() bool {
		minSeqNum, err := ccipContracts.Dest.CommitStore.GetExpectedNextSequenceNumber(nil, onRamp)
		require.NoError(t, err)
		ccipContracts.Source.Chain.Commit()
		ccipContracts.Dest.Chain.Commit()
		t.Log("min seq num reported", minSeqNum)
		return minSeqNum > uint64(max)
	}, testutils.WaitTimeout(t), 1*time.Second).Should(gomega.BeTrue(), "report has not been committed")
}

func GetEVMExtraArgsV1(gasLimit *big.Int, strict bool) ([]byte, error) {
	EVMV1Tag := []byte{0x97, 0xa6, 0x57, 0xc9}

	encodedArgs, err := utils.ABIEncode(`[{"type":"uint256"},{"type":"bool"}]`, gasLimit, strict)
	if err != nil {
		return nil, err
	}

	return append(EVMV1Tag, encodedArgs...), nil
}

func ExecuteGEMessage(
	t *testing.T,
	ccipContracts CCIPContracts,
	req logpoller.Log,
	allReqs []logpoller.Log,
	report commit_store.InternalCommitReport,
) uint64 {
	t.Log("Executing request manually")
	// Build full tree for report
	mctx := hasher.NewKeccakCtx()
	leafHasher := ccip.NewGELeafHasher(ccipContracts.Source.ChainID, ccipContracts.Dest.ChainID, ccipContracts.Source.GEOnRamp.Address(), mctx)

	var leafHashes [][32]byte
	for _, otherReq := range allReqs {
		hash, err := leafHasher.HashLeaf(otherReq.GetGethLog())
		require.NoError(t, err)
		leafHashes = append(leafHashes, hash)
	}
	intervalsByOnRamp := make(map[common.Address]commit_store.InternalInterval)
	merkleRootsByOnRamp := make(map[common.Address][32]byte)
	for i, onRamp := range report.OnRamps {
		intervalsByOnRamp[onRamp] = report.Intervals[i]
		merkleRootsByOnRamp[onRamp] = report.MerkleRoots[i]
	}
	interval := intervalsByOnRamp[ccipContracts.Source.GEOnRamp.Address()]
	decodedMsg, err := ccip.DecodeGEMessage(req.Data)
	require.NoError(t, err)
	innerIdx := int(decodedMsg.SequenceNumber - interval.Min)
	innerTree, err := merklemulti.NewTree(mctx, leafHashes)
	require.NoError(t, err)
	innerProof := innerTree.Prove([]int{innerIdx})
	var onRampIdx int
	var outerTreeLeafs [][32]byte
	for i, onRamp := range report.OnRamps {
		if onRamp == ccipContracts.Source.GEOnRamp.Address() {
			onRampIdx = i
		}
		outerTreeLeafs = append(outerTreeLeafs, merkleRootsByOnRamp[onRamp])
	}
	outerTree, err := merklemulti.NewTree(mctx, outerTreeLeafs)
	require.NoError(t, err)
	require.Equal(t, outerTree.Root(), report.RootOfRoots, "Roots donot match")

	outerProof := outerTree.Prove([]int{onRampIdx})

	offRampProof := evm_2_evm_ge_offramp.GEExecutionReport{
		SequenceNumbers:          []uint64{decodedMsg.SequenceNumber},
		TokenPerFeeCoinAddresses: []common.Address{ccipContracts.Source.LinkToken.Address()},
		TokenPerFeeCoin:          []*big.Int{big.NewInt(1)},
		EncodedMessages:          [][]byte{req.Data},
		InnerProofs:              innerProof.Hashes,
		InnerProofFlagBits:       ccip.ProofFlagsToBits(innerProof.SourceFlags),
		OuterProofs:              outerProof.Hashes,
		OuterProofFlagBits:       ccip.ProofFlagsToBits(outerProof.SourceFlags),
	}

	// Execute.
	tx, err := ccipContracts.Dest.GEOffRamp.ManuallyExecute(ccipContracts.Dest.User, offRampProof)
	require.NoError(t, err, "Executing manually")
	ccipContracts.Dest.Chain.Commit()
	ccipContracts.Source.Chain.Commit()
	rec, err := ccipContracts.Dest.Chain.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.Equal(t, uint64(1), rec.Status, "manual execution failed")
	t.Logf("Manual Execution completed for seqNum %d", decodedMsg.SequenceNumber)
	return decodedMsg.SequenceNumber
}
