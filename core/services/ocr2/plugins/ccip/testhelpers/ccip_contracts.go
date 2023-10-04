package testhelpers

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2/types"
	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/arm_proxy_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/commit_store_helper"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/custom_token_pool"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/maybe_revert_message_receiver"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/mock_arm_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/weth9"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/shared/generated/burn_mint_erc677"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/hashlib"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/merklemulti"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

var (
	// Source
	SourcePool   = "source pool"
	SourcePrices = "source prices"
	OnRamp       = "onramp"
	OnRampNative = "onramp-native"
	SourceRouter = "source router"

	// Dest
	OffRamp  = "offramp"
	DestPool = "dest pool"

	Receiver            = "receiver"
	Sender              = "sender"
	Link                = func(amount int64) *big.Int { return new(big.Int).Mul(big.NewInt(1e18), big.NewInt(amount)) }
	HundredLink         = Link(100)
	LinkUSDValue        = func(amount int64) *big.Int { return new(big.Int).Mul(big.NewInt(1e18), big.NewInt(amount)) }
	SourceChainID       = uint64(1000)
	SourceChainSelector = uint64(11787463284727550157)
	DestChainID         = uint64(1337)
	DestChainSelector   = uint64(3379446385462418246)
)

type MaybeRevertReceiver struct {
	Receiver *maybe_revert_message_receiver.MaybeRevertMessageReceiver
	Strict   bool
}

type Common struct {
	ChainID           uint64
	ChainSelector     uint64
	User              *bind.TransactOpts
	Chain             *backends.SimulatedBackend
	LinkToken         *link_token_interface.LinkToken
	Pool              *lock_release_token_pool.LockReleaseTokenPool
	CustomPool        *custom_token_pool.CustomTokenPool
	CustomToken       *link_token_interface.LinkToken
	WrappedNative     *weth9.WETH9
	WrappedNativePool *lock_release_token_pool.LockReleaseTokenPool
	ARM               *mock_arm_contract.MockARMContract
	ARMProxy          *arm_proxy_contract.ARMProxyContract
	PriceRegistry     *price_registry.PriceRegistry
}

type SourceChain struct {
	Common
	Router *router.Router
	OnRamp *evm_2_evm_onramp.EVM2EVMOnRamp
}

type DestinationChain struct {
	Common

	CommitStoreHelper *commit_store_helper.CommitStoreHelper
	CommitStore       *commit_store.CommitStore
	Router            *router.Router
	OffRamp           *evm_2_evm_offramp.EVM2EVMOffRamp
	Receivers         []MaybeRevertReceiver
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
	Getter   func(t *testing.T, addr common.Address) *big.Int
	Within   string
}

type BalanceReq struct {
	Name   string
	Addr   common.Address
	Getter func(t *testing.T, addr common.Address) *big.Int
}

type CCIPContracts struct {
	Source  SourceChain
	Dest    DestinationChain
	Oracles []confighelper.OracleIdentityExtra

	commitOCRConfig, execOCRConfig *OCR2Config
}

func (c *CCIPContracts) DeployNewOffRamp(t *testing.T) {
	prevOffRamp := common.HexToAddress("")
	if c.Dest.OffRamp != nil {
		prevOffRamp = c.Dest.OffRamp.Address()
	}
	offRampAddress, _, _, err := evm_2_evm_offramp.DeployEVM2EVMOffRamp(
		c.Dest.User,
		c.Dest.Chain,
		evm_2_evm_offramp.EVM2EVMOffRampStaticConfig{
			CommitStore:         c.Dest.CommitStore.Address(),
			ChainSelector:       c.Dest.ChainSelector,
			SourceChainSelector: c.Source.ChainSelector,
			OnRamp:              c.Source.OnRamp.Address(),
			PrevOffRamp:         prevOffRamp,
			ArmProxy:            c.Dest.ARMProxy.Address(),
		},
		[]common.Address{c.Source.LinkToken.Address()}, // source tokens
		[]common.Address{c.Dest.Pool.Address()},        // pools
		evm_2_evm_offramp.RateLimiterConfig{
			IsEnabled: true,
			Capacity:  LinkUSDValue(100),
			Rate:      LinkUSDValue(1),
		},
	)
	require.NoError(t, err)
	c.Dest.Chain.Commit()

	c.Dest.OffRamp, err = evm_2_evm_offramp.NewEVM2EVMOffRamp(offRampAddress, c.Dest.Chain)
	require.NoError(t, err)

	c.Dest.Chain.Commit()
	c.Source.Chain.Commit()
}

func (c *CCIPContracts) EnableOffRamp(t *testing.T) {
	_, err := c.Dest.Pool.ApplyRampUpdates(c.Dest.User,
		[]lock_release_token_pool.TokenPoolRampUpdate{},
		[]lock_release_token_pool.TokenPoolRampUpdate{{Ramp: c.Dest.OffRamp.Address(), Allowed: true,
			RateLimiterConfig: lock_release_token_pool.RateLimiterConfig{
				IsEnabled: true,
				Capacity:  HundredLink,
				Rate:      big.NewInt(1e18),
			},
		}},
	)

	require.NoError(t, err)
	c.Dest.Chain.Commit()

	_, err = c.Dest.Router.ApplyRampUpdates(c.Dest.User, nil, nil, []router.RouterOffRamp{{SourceChainSelector: SourceChainSelector, OffRamp: c.Dest.OffRamp.Address()}})
	require.NoError(t, err)
	c.Dest.Chain.Commit()

	onChainConfig := c.CreateDefaultExecOnchainConfig(t)
	offChainConfig := c.CreateDefaultExecOffchainConfig(t)

	c.SetupExecOCR2Config(t, onChainConfig, offChainConfig)
}

func (c *CCIPContracts) EnableCommitStore(t *testing.T) {
	onChainConfig := c.CreateDefaultCommitOnchainConfig(t)
	offChainConfig := c.CreateDefaultCommitOffchainConfig(t)

	c.SetupCommitOCR2Config(t, onChainConfig, offChainConfig)

	_, err := c.Dest.PriceRegistry.ApplyPriceUpdatersUpdates(c.Dest.User, []common.Address{c.Dest.CommitStore.Address()}, []common.Address{})
	require.NoError(t, err)
	c.Dest.Chain.Commit()
}

func (c *CCIPContracts) DeployNewOnRamp(t *testing.T) {
	t.Log("Deploying new onRamp")
	// find the last onRamp
	prevOnRamp := common.HexToAddress("")
	if c.Source.OnRamp != nil {
		prevOnRamp = c.Source.OnRamp.Address()
	}
	onRampAddress, _, _, err := evm_2_evm_onramp.DeployEVM2EVMOnRamp(
		c.Source.User,  // user
		c.Source.Chain, // client
		evm_2_evm_onramp.EVM2EVMOnRampStaticConfig{
			LinkToken:         c.Source.LinkToken.Address(),
			ChainSelector:     c.Source.ChainSelector,
			DestChainSelector: c.Dest.ChainSelector,
			DefaultTxGasLimit: 200_000,
			MaxNopFeesJuels:   big.NewInt(0).Mul(big.NewInt(100_000_000), big.NewInt(1e18)),
			PrevOnRamp:        prevOnRamp,
			ArmProxy:          c.Source.ARM.Address(), // ARM
		},
		evm_2_evm_onramp.EVM2EVMOnRampDynamicConfig{
			Router:                          c.Source.Router.Address(),
			MaxTokensLength:                 5,
			DestGasOverhead:                 350_000,
			DestGasPerPayloadByte:           16,
			DestDataAvailabilityOverheadGas: 33_596,
			DestGasPerDataAvailabilityByte:  16,
			DestDataAvailabilityMultiplier:  6840, // 0.684
			PriceRegistry:                   c.Source.PriceRegistry.Address(),
			MaxDataSize:                     1e5,
			MaxGasLimit:                     4_000_000,
		},
		[]evm_2_evm_onramp.InternalPoolUpdate{
			{
				Token: c.Source.LinkToken.Address(),
				Pool:  c.Source.Pool.Address(),
			},
		},
		evm_2_evm_onramp.RateLimiterConfig{
			IsEnabled: true,
			Capacity:  LinkUSDValue(100),
			Rate:      LinkUSDValue(1),
		},
		[]evm_2_evm_onramp.EVM2EVMOnRampFeeTokenConfigArgs{
			{
				Token:                  c.Source.LinkToken.Address(),
				NetworkFeeUSD:          1_00,
				MinTokenTransferFeeUSD: 1_00,
				MaxTokenTransferFeeUSD: 5000_00,
				GasMultiplier:          1e18,
				PremiumMultiplier:      9e17,
				Enabled:                true,
			},
			{
				Token:                  c.Source.WrappedNative.Address(),
				NetworkFeeUSD:          1_00,
				MinTokenTransferFeeUSD: 1_00,
				MaxTokenTransferFeeUSD: 5000_00,
				GasMultiplier:          1e18,
				PremiumMultiplier:      1e18,
				Enabled:                true,
			},
		},
		[]evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfigArgs{
			{
				Token:             c.Source.LinkToken.Address(),
				Ratio:             5_0, // 5 bps
				DestGasOverhead:   34_000,
				DestBytesOverhead: 0,
			},
		},
		[]evm_2_evm_onramp.EVM2EVMOnRampNopAndWeight{},
	)

	require.NoError(t, err)
	c.Source.Chain.Commit()
	c.Dest.Chain.Commit()
	c.Source.OnRamp, err = evm_2_evm_onramp.NewEVM2EVMOnRamp(onRampAddress, c.Source.Chain)
	require.NoError(t, err)
	c.Source.Chain.Commit()
	c.Dest.Chain.Commit()
}

func (c *CCIPContracts) EnableOnRamp(t *testing.T) {
	t.Log("Setting onRamp on source pool")
	_, err := c.Source.Pool.ApplyRampUpdates(
		c.Source.User,
		[]lock_release_token_pool.TokenPoolRampUpdate{
			{
				Ramp:    c.Source.OnRamp.Address(),
				Allowed: true,
				RateLimiterConfig: lock_release_token_pool.RateLimiterConfig{
					IsEnabled: true,
					Capacity:  HundredLink,
					Rate:      big.NewInt(1e18),
				},
			},
		},
		[]lock_release_token_pool.TokenPoolRampUpdate{},
	)

	require.NoError(t, err)
	c.Source.Chain.Commit()

	t.Log("Setting onRamp on source router")
	_, err = c.Source.Router.ApplyRampUpdates(c.Source.User, []router.RouterOnRamp{{DestChainSelector: c.Dest.ChainSelector, OnRamp: c.Source.OnRamp.Address()}}, nil, nil)
	require.NoError(t, err)
	c.Source.Chain.Commit()
	c.Dest.Chain.Commit()
}

func (c *CCIPContracts) DeployNewCommitStore(t *testing.T) {
	commitStoreAddress, _, _, err := commit_store_helper.DeployCommitStoreHelper(
		c.Dest.User,  // user
		c.Dest.Chain, // client
		commit_store_helper.CommitStoreStaticConfig{
			ChainSelector:       c.Dest.ChainSelector,
			SourceChainSelector: c.Source.ChainSelector,
			OnRamp:              c.Source.OnRamp.Address(),
			ArmProxy:            c.Dest.ARMProxy.Address(),
		},
	)
	require.NoError(t, err)
	c.Dest.Chain.Commit()
	c.Dest.CommitStoreHelper, err = commit_store_helper.NewCommitStoreHelper(commitStoreAddress, c.Dest.Chain)
	require.NoError(t, err)
	// since CommitStoreHelper derives from CommitStore, it's safe to instantiate both on same address
	c.Dest.CommitStore, err = commit_store.NewCommitStore(commitStoreAddress, c.Dest.Chain)
	require.NoError(t, err)
}

func (c *CCIPContracts) DeployNewPriceRegistry(t *testing.T) {
	t.Log("Deploying new Price Registry")
	destPricesAddress, _, _, err := price_registry.DeployPriceRegistry(
		c.Dest.User,
		c.Dest.Chain,
		[]common.Address{c.Dest.CommitStore.Address()},
		[]common.Address{c.Dest.LinkToken.Address()},
		60*60*24*14, // two weeks
	)
	require.NoError(t, err)
	c.Source.Chain.Commit()
	c.Dest.Chain.Commit()
	c.Dest.PriceRegistry, err = price_registry.NewPriceRegistry(destPricesAddress, c.Dest.Chain)
	require.NoError(t, err)
	t.Logf("New Price Registry deployed at %s", destPricesAddress.String())

	priceUpdates := price_registry.InternalPriceUpdates{
		TokenPriceUpdates: []price_registry.InternalTokenPriceUpdate{
			{
				SourceToken: c.Dest.LinkToken.Address(),
				UsdPerToken: big.NewInt(8e18), // 8usd
			},
			{
				SourceToken: c.Dest.WrappedNative.Address(),
				UsdPerToken: big.NewInt(1e18), // 1usd
			},
		},
		DestChainSelector: c.Source.ChainSelector,
		UsdPerUnitGas:     big.NewInt(2000e9), // $2000 per eth * 1gwei = 2000e9
	}
	_, err = c.Dest.PriceRegistry.UpdatePrices(c.Dest.User, priceUpdates)
	require.NoError(t, err)

	c.Source.Chain.Commit()
	c.Dest.Chain.Commit()
}

func (c *CCIPContracts) SetNopsOnRamp(t *testing.T, nopsAndWeights []evm_2_evm_onramp.EVM2EVMOnRampNopAndWeight) {
	tx, err := c.Source.OnRamp.SetNops(c.Source.User, nopsAndWeights)
	require.NoError(t, err)
	c.Source.Chain.Commit()
	_, err = bind.WaitMined(context.Background(), c.Source.Chain, tx)
	require.NoError(t, err)
}

func (c *CCIPContracts) GetSourceLinkBalance(t *testing.T, addr common.Address) *big.Int {
	return GetBalance(t, c.Source.Chain, c.Source.LinkToken.Address(), addr)
}

func (c *CCIPContracts) GetDestLinkBalance(t *testing.T, addr common.Address) *big.Int {
	return GetBalance(t, c.Dest.Chain, c.Dest.LinkToken.Address(), addr)
}

func (c *CCIPContracts) GetSourceWrappedTokenBalance(t *testing.T, addr common.Address) *big.Int {
	return GetBalance(t, c.Source.Chain, c.Source.WrappedNative.Address(), addr)
}

func (c *CCIPContracts) GetDestWrappedTokenBalance(t *testing.T, addr common.Address) *big.Int {
	return GetBalance(t, c.Dest.Chain, c.Dest.WrappedNative.Address(), addr)
}

func (c *CCIPContracts) AssertBalances(t *testing.T, bas []BalanceAssertion) {
	for _, b := range bas {
		actual := b.Getter(t, b.Address)
		t.Log("Checking balance for", b.Name, "at", b.Address.Hex(), "got", actual)
		require.NotNil(t, actual, "%v getter return nil", b.Name)
		if b.Within == "" {
			require.Equal(t, b.Expected, actual.String(), "wrong balance for %s got %s want %s", b.Name, actual, b.Expected)
		} else {
			bi, _ := big.NewInt(0).SetString(b.Expected, 10)
			withinI, _ := big.NewInt(0).SetString(b.Within, 10)
			high := big.NewInt(0).Add(bi, withinI)
			low := big.NewInt(0).Sub(bi, withinI)
			require.Equal(t, -1, actual.Cmp(high), "wrong balance for %s got %s outside expected range [%s, %s]", b.Name, actual, low, high)
			require.Equal(t, 1, actual.Cmp(low), "wrong balance for %s got %s outside expected range [%s, %s]", b.Name, actual, low, high)
		}
	}
}

func AccountToAddress(accounts []ocr2types.Account) (addresses []common.Address, err error) {
	for _, signer := range accounts {
		bytes, err := hexutil.Decode(string(signer))
		if err != nil {
			return []common.Address{}, errors.Wrap(err, fmt.Sprintf("given address is not valid %s", signer))
		}
		if len(bytes) != 20 {
			return []common.Address{}, errors.Errorf("address is not 20 bytes %s", signer)
		}
		addresses = append(addresses, common.BytesToAddress(bytes))
	}
	return addresses, nil
}

func OnchainPublicKeyToAddress(publicKeys []ocrtypes.OnchainPublicKey) (addresses []common.Address, err error) {
	for _, signer := range publicKeys {
		if len(signer) != 20 {
			return []common.Address{}, errors.Errorf("address is not 20 bytes %s", signer)
		}
		addresses = append(addresses, common.BytesToAddress(signer))
	}
	return addresses, nil
}

func (c *CCIPContracts) DeriveOCR2Config(t *testing.T, oracles []confighelper.OracleIdentityExtra, rawOnchainConfig []byte, rawOffchainConfig []byte) *OCR2Config {
	signers, transmitters, threshold, onchainConfig, offchainConfigVersion, offchainConfig, err := confighelper.ContractSetConfigArgsForTests(
		2*time.Second,        // deltaProgress
		1*time.Second,        // deltaResend
		1*time.Second,        // deltaRound
		500*time.Millisecond, // deltaGrace
		2*time.Second,        // deltaStage
		3,
		[]int{1, 1, 1, 1},
		oracles,
		rawOffchainConfig,
		50*time.Millisecond, // Max duration query
		1*time.Second,       // Max duration observation
		100*time.Millisecond,
		100*time.Millisecond,
		100*time.Millisecond,
		1, // faults
		rawOnchainConfig,
	)
	require.NoError(t, err)
	lggr := logger.TestLogger(t)
	lggr.Infow("Setting Config on Oracle Contract",
		"signers", signers,
		"transmitters", transmitters,
		"threshold", threshold,
		"onchainConfig", onchainConfig,
		"encodedConfigVersion", offchainConfigVersion,
	)
	signerAddresses, err := OnchainPublicKeyToAddress(signers)
	require.NoError(t, err)
	transmitterAddresses, err := AccountToAddress(transmitters)
	require.NoError(t, err)

	return &OCR2Config{
		Signers:               signerAddresses,
		Transmitters:          transmitterAddresses,
		F:                     threshold,
		OnchainConfig:         onchainConfig,
		OffchainConfigVersion: offchainConfigVersion,
		OffchainConfig:        offchainConfig,
	}
}

func (c *CCIPContracts) SetupCommitOCR2Config(t *testing.T, commitOnchainConfig, commitOffchainConfig []byte) {
	c.commitOCRConfig = c.DeriveOCR2Config(t, c.Oracles, commitOnchainConfig, commitOffchainConfig)
	// Set the DON on the commit store
	_, err := c.Dest.CommitStore.SetOCR2Config(
		c.Dest.User,
		c.commitOCRConfig.Signers,
		c.commitOCRConfig.Transmitters,
		c.commitOCRConfig.F,
		c.commitOCRConfig.OnchainConfig,
		c.commitOCRConfig.OffchainConfigVersion,
		c.commitOCRConfig.OffchainConfig,
	)
	require.NoError(t, err)
	c.Dest.Chain.Commit()
}

func (c *CCIPContracts) SetupExecOCR2Config(t *testing.T, execOnchainConfig, execOffchainConfig []byte) {
	c.execOCRConfig = c.DeriveOCR2Config(t, c.Oracles, execOnchainConfig, execOffchainConfig)
	// Same DON on the offramp
	_, err := c.Dest.OffRamp.SetOCR2Config(
		c.Dest.User,
		c.execOCRConfig.Signers,
		c.execOCRConfig.Transmitters,
		c.execOCRConfig.F,
		c.execOCRConfig.OnchainConfig,
		c.execOCRConfig.OffchainConfigVersion,
		c.execOCRConfig.OffchainConfig,
	)
	require.NoError(t, err)
	c.Dest.Chain.Commit()
}

func (c *CCIPContracts) SetupOnchainConfig(t *testing.T, commitOnchainConfig, commitOffchainConfig, execOnchainConfig, execOffchainConfig []byte) int64 {
	// Note We do NOT set the payees, payment is done in the OCR2Base implementation
	blockBeforeConfig, err := c.Dest.Chain.BlockByNumber(context.Background(), nil)
	require.NoError(t, err)

	c.SetupCommitOCR2Config(t, commitOnchainConfig, commitOffchainConfig)
	c.SetupExecOCR2Config(t, execOnchainConfig, execOffchainConfig)

	return blockBeforeConfig.Number().Int64()
}

func (c *CCIPContracts) SetupLockAndMintTokenPool(
	sourceTokenAddress common.Address,
	wrappedTokenName,
	wrappedTokenSymbol string) (common.Address, *burn_mint_erc677.BurnMintERC677, error) {
	// Deploy dest token & pool
	destTokenAddress, _, _, err := burn_mint_erc677.DeployBurnMintERC677(c.Dest.User, c.Dest.Chain, wrappedTokenName, wrappedTokenSymbol, 18, big.NewInt(0))
	if err != nil {
		return [20]byte{}, nil, err
	}
	c.Dest.Chain.Commit()

	destToken, err := burn_mint_erc677.NewBurnMintERC677(destTokenAddress, c.Dest.Chain)
	if err != nil {
		return [20]byte{}, nil, err
	}

	destPoolAddress, _, destPool, err := burn_mint_token_pool.DeployBurnMintTokenPool(
		c.Dest.User,
		c.Dest.Chain,
		destTokenAddress,
		[]common.Address{}, // pool originalSender allowList
		c.Dest.ARMProxy.Address(),
	)
	if err != nil {
		return [20]byte{}, nil, err
	}
	c.Dest.Chain.Commit()

	_, err = destToken.GrantMintAndBurnRoles(c.Dest.User, destPoolAddress)
	if err != nil {
		return [20]byte{}, nil, err
	}

	_, err = destPool.ApplyRampUpdates(c.Dest.User, nil, []burn_mint_token_pool.TokenPoolRampUpdate{
		{Ramp: c.Dest.OffRamp.Address(), Allowed: true,
			RateLimiterConfig: burn_mint_token_pool.RateLimiterConfig{
				IsEnabled: true,
				Capacity:  HundredLink,
				Rate:      big.NewInt(1e18),
			},
		},
	})
	if err != nil {
		return [20]byte{}, nil, err
	}
	c.Dest.Chain.Commit()

	sourcePoolAddress, _, sourcePool, err := lock_release_token_pool.DeployLockReleaseTokenPool(
		c.Source.User,
		c.Source.Chain,
		sourceTokenAddress,
		[]common.Address{}, // empty allowList at deploy time indicates pool has no original sender restrictions
		c.Source.ARMProxy.Address(),
		true,
	)
	if err != nil {
		return [20]byte{}, nil, err
	}
	c.Source.Chain.Commit()

	// set onRamp as valid caller for source pool
	_, err = sourcePool.ApplyRampUpdates(c.Source.User, []lock_release_token_pool.TokenPoolRampUpdate{
		{
			Ramp:    c.Source.OnRamp.Address(),
			Allowed: true,
			RateLimiterConfig: lock_release_token_pool.RateLimiterConfig{
				IsEnabled: true,
				Capacity:  HundredLink,
				Rate:      big.NewInt(1e18),
			},
		},
	}, nil)
	if err != nil {
		return [20]byte{}, nil, err
	}
	c.Source.Chain.Commit()

	wrappedNativeAddress, err := c.Source.Router.GetWrappedNative(nil)
	if err != nil {
		return [20]byte{}, nil, err
	}

	// native token is used as fee token
	_, err = c.Source.PriceRegistry.UpdatePrices(c.Source.User, price_registry.InternalPriceUpdates{
		TokenPriceUpdates: []price_registry.InternalTokenPriceUpdate{
			{
				SourceToken: wrappedNativeAddress,
				UsdPerToken: big.NewInt(1e18), // 1usd
			},
		},
		DestChainSelector: c.Dest.ChainSelector,
		UsdPerUnitGas:     big.NewInt(2000e9), // $2000 per eth * 1gwei = 2000e9,
	})
	if err != nil {
		return [20]byte{}, nil, err
	}
	c.Source.Chain.Commit()
	_, err = c.Source.PriceRegistry.ApplyFeeTokensUpdates(c.Source.User, []common.Address{wrappedNativeAddress}, nil)
	if err != nil {
		return [20]byte{}, nil, err
	}
	c.Source.Chain.Commit()

	// add new token pool created above
	_, err = c.Source.OnRamp.ApplyPoolUpdates(c.Source.User, nil, []evm_2_evm_onramp.InternalPoolUpdate{
		{
			Token: sourceTokenAddress,
			Pool:  sourcePoolAddress,
		},
	})
	if err != nil {
		return [20]byte{}, nil, err
	}

	_, err = c.Source.PriceRegistry.UpdatePrices(c.Source.User, price_registry.InternalPriceUpdates{
		TokenPriceUpdates: []price_registry.InternalTokenPriceUpdate{
			{
				SourceToken: sourceTokenAddress,
				UsdPerToken: big.NewInt(5),
			},
		},
		DestChainSelector: c.Source.ChainSelector,
		UsdPerUnitGas:     big.NewInt(0),
	})
	if err != nil {
		return [20]byte{}, nil, err
	}
	c.Source.Chain.Commit()

	_, err = c.Dest.OffRamp.ApplyPoolUpdates(c.Dest.User, nil, []evm_2_evm_offramp.InternalPoolUpdate{
		{
			Token: sourceTokenAddress,
			Pool:  destPoolAddress,
		},
	})
	if err != nil {
		return [20]byte{}, nil, err
	}
	c.Dest.Chain.Commit()

	_, err = c.Dest.PriceRegistry.UpdatePrices(c.Dest.User, price_registry.InternalPriceUpdates{
		TokenPriceUpdates: []price_registry.InternalTokenPriceUpdate{
			{
				SourceToken: destPoolAddress,
				UsdPerToken: big.NewInt(5),
			},
		},
		DestChainSelector: 0,
		UsdPerUnitGas:     big.NewInt(0),
	})
	if err != nil {
		return [20]byte{}, nil, err
	}
	c.Dest.Chain.Commit()

	return sourcePoolAddress, destToken, err
}

func (c *CCIPContracts) SendMessage(t *testing.T, gasLimit, tokenAmount *big.Int, receiverAddr common.Address) {
	extraArgs, err := GetEVMExtraArgsV1(gasLimit, false)
	require.NoError(t, err)
	msg := router.ClientEVM2AnyMessage{
		Receiver: MustEncodeAddress(t, receiverAddr),
		Data:     []byte("hello"),
		TokenAmounts: []router.ClientEVMTokenAmount{
			{
				Token:  c.Source.LinkToken.Address(),
				Amount: tokenAmount,
			},
		},
		FeeToken:  c.Source.LinkToken.Address(),
		ExtraArgs: extraArgs,
	}
	fee, err := c.Source.Router.GetFee(nil, c.Dest.ChainSelector, msg)
	require.NoError(t, err)
	// Currently no overhead and 1gwei dest gas price. So fee is simply gasLimit * gasPrice.
	// require.Equal(t, new(big.Int).Mul(gasLimit, gasPrice).String(), fee.String())
	// Approve the fee amount + the token amount
	_, err = c.Source.LinkToken.Approve(c.Source.User, c.Source.Router.Address(), new(big.Int).Add(fee, tokenAmount))
	require.NoError(t, err)
	c.Source.Chain.Commit()
	c.SendRequest(t, msg)
}

func GetBalances(t *testing.T, brs []BalanceReq) (map[string]*big.Int, error) {
	m := make(map[string]*big.Int)
	for _, br := range brs {
		m[br.Name] = br.Getter(t, br.Addr)
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

func SetupCCIPContracts(t *testing.T, sourceChainID, sourceChainSelector, destChainID, destChainSelector uint64) CCIPContracts {
	sourceChain, sourceUser := SetupChain(t)
	destChain, destUser := SetupChain(t)

	armSourceAddress, _, _, err := mock_arm_contract.DeployMockARMContract(
		sourceUser,
		sourceChain,
	)
	require.NoError(t, err)
	sourceARM, err := mock_arm_contract.NewMockARMContract(armSourceAddress, sourceChain)
	require.NoError(t, err)
	armProxySourceAddress, _, _, err := arm_proxy_contract.DeployARMProxyContract(
		sourceUser,
		sourceChain,
		armSourceAddress,
	)
	require.NoError(t, err)
	sourceARMProxy, err := arm_proxy_contract.NewARMProxyContract(armProxySourceAddress, sourceChain)
	require.NoError(t, err)
	sourceChain.Commit()

	armDestAddress, _, _, err := mock_arm_contract.DeployMockARMContract(
		destUser,
		destChain,
	)
	require.NoError(t, err)
	armProxyDestAddress, _, _, err := arm_proxy_contract.DeployARMProxyContract(
		destUser,
		destChain,
		armDestAddress,
	)
	require.NoError(t, err)
	destChain.Commit()
	destARM, err := mock_arm_contract.NewMockARMContract(armDestAddress, destChain)
	require.NoError(t, err)
	destARMProxy, err := arm_proxy_contract.NewARMProxyContract(armProxyDestAddress, destChain)
	require.NoError(t, err)

	// Deploy link token and pool on source chain
	sourceLinkTokenAddress, _, _, err := link_token_interface.DeployLinkToken(sourceUser, sourceChain)
	require.NoError(t, err)
	sourceChain.Commit()
	sourceLinkToken, err := link_token_interface.NewLinkToken(sourceLinkTokenAddress, sourceChain)
	require.NoError(t, err)
	sourcePoolAddress, _, _, err := lock_release_token_pool.DeployLockReleaseTokenPool(
		sourceUser,
		sourceChain,
		sourceLinkTokenAddress,
		[]common.Address{},
		armProxySourceAddress,
		true,
	)
	require.NoError(t, err)
	sourceChain.Commit()
	sourcePool, err := lock_release_token_pool.NewLockReleaseTokenPool(sourcePoolAddress, sourceChain)
	require.NoError(t, err)

	// Deploy link token and pool on destination chain
	destLinkTokenAddress, _, _, err := link_token_interface.DeployLinkToken(destUser, destChain)
	require.NoError(t, err)
	destChain.Commit()
	destLinkToken, err := link_token_interface.NewLinkToken(destLinkTokenAddress, destChain)
	require.NoError(t, err)
	destPoolAddress, _, _, err := lock_release_token_pool.DeployLockReleaseTokenPool(
		destUser,
		destChain,
		destLinkTokenAddress,
		[]common.Address{},
		armProxyDestAddress,
		true,
	)
	require.NoError(t, err)
	destChain.Commit()
	destPool, err := lock_release_token_pool.NewLockReleaseTokenPool(destPoolAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()

	// Float the offramp pool
	o, err := destPool.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, destUser.From.String(), o.String())
	_, err = destLinkToken.Approve(destUser, destPoolAddress, Link(200))
	require.NoError(t, err)
	_, err = destPool.AddLiquidity(destUser, Link(200))
	require.NoError(t, err)
	destChain.Commit()

	// Deploy custom token pool source
	sourceCustomTokenAddress, _, _, err := link_token_interface.DeployLinkToken(sourceUser, sourceChain) // Just re-use this, it's an ERC20.
	require.NoError(t, err)
	sourceCustomToken, err := link_token_interface.NewLinkToken(sourceCustomTokenAddress, sourceChain)
	require.NoError(t, err)
	destChain.Commit()

	// Deploy custom token pool dest
	destCustomTokenAddress, _, _, err := link_token_interface.DeployLinkToken(destUser, destChain) // Just re-use this, it's an ERC20.
	require.NoError(t, err)
	destCustomToken, err := link_token_interface.NewLinkToken(destCustomTokenAddress, destChain)
	require.NoError(t, err)
	destChain.Commit()

	// Create router
	sourceWeth9addr, _, _, err := weth9.DeployWETH9(sourceUser, sourceChain)
	require.NoError(t, err)
	sourceWrapped, err := weth9.NewWETH9(sourceWeth9addr, sourceChain)
	require.NoError(t, err)
	sourceWeth9PoolAddress, _, _, err := lock_release_token_pool.DeployLockReleaseTokenPool(
		sourceUser,
		sourceChain,
		sourceWeth9addr,
		[]common.Address{},
		armProxySourceAddress,
		true,
	)
	require.NoError(t, err)
	sourceChain.Commit()

	sourceWeth9Pool, err := lock_release_token_pool.NewLockReleaseTokenPool(sourceWeth9PoolAddress, sourceChain)
	require.NoError(t, err)

	sourceRouterAddress, _, _, err := router.DeployRouter(sourceUser, sourceChain, sourceWeth9addr, armProxySourceAddress)
	require.NoError(t, err)
	sourceRouter, err := router.NewRouter(sourceRouterAddress, sourceChain)
	require.NoError(t, err)
	sourceChain.Commit()

	// Deploy and configure onramp
	sourcePricesAddress, _, _, err := price_registry.DeployPriceRegistry(
		sourceUser,
		sourceChain,
		nil,
		[]common.Address{sourceLinkTokenAddress, sourceWeth9addr},
		60*60*24*14, // two weeks
	)
	require.NoError(t, err)

	srcPriceRegistry, err := price_registry.NewPriceRegistry(sourcePricesAddress, sourceChain)
	require.NoError(t, err)

	prices := price_registry.InternalPriceUpdates{
		TokenPriceUpdates: []price_registry.InternalTokenPriceUpdate{
			{
				SourceToken: sourceLinkTokenAddress,
				UsdPerToken: big.NewInt(8e18), // 8usd
			},
			{
				SourceToken: sourceWeth9addr,
				UsdPerToken: new(big.Int).Mul(big.NewInt(1e18), big.NewInt(2)), // TODO make this 2000USD and once we figure out the fee and exec cost discrepancy
			},
		},
		DestChainSelector: destChainSelector,
		UsdPerUnitGas:     big.NewInt(2000e9), // $2000 per eth * 1gwei = 2000e9
	}

	_, err = srcPriceRegistry.UpdatePrices(sourceUser, prices)
	require.NoError(t, err)

	onRampAddress, _, _, err := evm_2_evm_onramp.DeployEVM2EVMOnRamp(
		sourceUser,  // user
		sourceChain, // client
		evm_2_evm_onramp.EVM2EVMOnRampStaticConfig{
			LinkToken:         sourceLinkTokenAddress,
			ChainSelector:     sourceChainSelector,
			DestChainSelector: destChainSelector,
			DefaultTxGasLimit: 200_000,
			MaxNopFeesJuels:   big.NewInt(0).Mul(big.NewInt(100_000_000), big.NewInt(1e18)),
			PrevOnRamp:        common.HexToAddress(""),
			ArmProxy:          armProxySourceAddress, // ARM
		},
		evm_2_evm_onramp.EVM2EVMOnRampDynamicConfig{
			Router:                          sourceRouterAddress,
			MaxTokensLength:                 5,
			DestGasOverhead:                 350_000,
			DestGasPerPayloadByte:           16,
			DestDataAvailabilityOverheadGas: 33_596,
			DestGasPerDataAvailabilityByte:  16,
			DestDataAvailabilityMultiplier:  6840, // 0.684
			PriceRegistry:                   sourcePricesAddress,
			MaxDataSize:                     1e5,
			MaxGasLimit:                     4_000_000,
		},
		[]evm_2_evm_onramp.InternalPoolUpdate{
			{
				Token: sourceLinkTokenAddress,
				Pool:  sourcePoolAddress,
			},
			{
				Token: sourceWeth9addr,
				Pool:  sourceWeth9PoolAddress,
			},
		},
		evm_2_evm_onramp.RateLimiterConfig{
			IsEnabled: true,
			Capacity:  LinkUSDValue(100),
			Rate:      LinkUSDValue(1),
		},
		[]evm_2_evm_onramp.EVM2EVMOnRampFeeTokenConfigArgs{
			{
				Token:                  sourceLinkTokenAddress,
				NetworkFeeUSD:          1_00,
				MinTokenTransferFeeUSD: 1_00,
				MaxTokenTransferFeeUSD: 5000_00,
				GasMultiplier:          1e18,
				PremiumMultiplier:      9e17,
				Enabled:                true,
			},
			{
				Token:                  sourceWeth9addr,
				NetworkFeeUSD:          1_00,
				MinTokenTransferFeeUSD: 1_00,
				MaxTokenTransferFeeUSD: 5000_00,
				GasMultiplier:          1e18,
				PremiumMultiplier:      1e18,
				Enabled:                true,
			},
		},
		[]evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfigArgs{
			{
				Token:             sourceLinkTokenAddress,
				Ratio:             5_0, // 5 bps
				DestGasOverhead:   34_000,
				DestBytesOverhead: 0,
			},
		},
		[]evm_2_evm_onramp.EVM2EVMOnRampNopAndWeight{},
	)
	require.NoError(t, err)
	onRamp, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(onRampAddress, sourceChain)
	require.NoError(t, err)
	_, err = sourcePool.ApplyRampUpdates(sourceUser,
		[]lock_release_token_pool.TokenPoolRampUpdate{{Ramp: onRampAddress, Allowed: true,
			RateLimiterConfig: lock_release_token_pool.RateLimiterConfig{
				IsEnabled: true,
				Capacity:  HundredLink,
				Rate:      big.NewInt(1e18),
			},
		}},
		[]lock_release_token_pool.TokenPoolRampUpdate{},
	)
	require.NoError(t, err)
	_, err = sourceWeth9Pool.ApplyRampUpdates(sourceUser,
		[]lock_release_token_pool.TokenPoolRampUpdate{{Ramp: onRampAddress, Allowed: true,
			RateLimiterConfig: lock_release_token_pool.RateLimiterConfig{
				IsEnabled: true,
				Capacity:  HundredLink,
				Rate:      big.NewInt(1e18),
			},
		}},
		[]lock_release_token_pool.TokenPoolRampUpdate{},
	)
	require.NoError(t, err)
	sourceChain.Commit()
	_, err = sourceRouter.ApplyRampUpdates(sourceUser, []router.RouterOnRamp{{DestChainSelector: destChainSelector, OnRamp: onRampAddress}}, nil, nil)
	require.NoError(t, err)
	sourceChain.Commit()

	destWeth9addr, _, _, err := weth9.DeployWETH9(destUser, destChain)
	require.NoError(t, err)
	destWrapped, err := weth9.NewWETH9(destWeth9addr, destChain)
	require.NoError(t, err)
	destWrappedPoolAddress, _, _, err := lock_release_token_pool.DeployLockReleaseTokenPool(
		destUser,
		destChain,
		destWeth9addr,
		[]common.Address{},
		armProxyDestAddress,
		true,
	)
	require.NoError(t, err)
	destWrappedPool, err := lock_release_token_pool.NewLockReleaseTokenPool(destWrappedPoolAddress, destChain)
	require.NoError(t, err)

	// Deploy and configure ge offramp.
	destPricesAddress, _, _, err := price_registry.DeployPriceRegistry(
		destUser,
		destChain,
		nil,
		[]common.Address{destLinkTokenAddress},
		60*60*24*14, // two weeks
	)
	require.NoError(t, err)
	destPriceRegistry, err := price_registry.NewPriceRegistry(destPricesAddress, destChain)
	require.NoError(t, err)

	destPrices := price_registry.InternalPriceUpdates{
		TokenPriceUpdates: []price_registry.InternalTokenPriceUpdate{
			{SourceToken: destLinkTokenAddress, UsdPerToken: big.NewInt(8e18)},   // 8usd
			{SourceToken: destCustomTokenAddress, UsdPerToken: big.NewInt(5e18)}, // 5usd
			{SourceToken: destWeth9addr, UsdPerToken: big.NewInt(2e18)},          // 2usd
		},
		DestChainSelector: sourceChainSelector,
		UsdPerUnitGas:     big.NewInt(2000e9), // $2000 per eth * 1gwei = 2000e9
	}

	_, err = destPriceRegistry.UpdatePrices(destUser, destPrices)
	require.NoError(t, err)

	// Deploy commit store.
	commitStoreAddress, _, _, err := commit_store_helper.DeployCommitStoreHelper(
		destUser,  // user
		destChain, // client
		commit_store_helper.CommitStoreStaticConfig{
			ChainSelector:       destChainSelector,
			SourceChainSelector: sourceChainSelector,
			OnRamp:              onRamp.Address(),
			ArmProxy:            destARMProxy.Address(),
		},
	)
	require.NoError(t, err)
	destChain.Commit()
	commitStore, err := commit_store.NewCommitStore(commitStoreAddress, destChain)
	require.NoError(t, err)
	commitStoreHelper, err := commit_store_helper.NewCommitStoreHelper(commitStoreAddress, destChain)
	require.NoError(t, err)

	// Create dest router
	destRouterAddress, _, _, err := router.DeployRouter(destUser, destChain, destWeth9addr, armProxyDestAddress)
	require.NoError(t, err)
	destChain.Commit()
	destRouter, err := router.NewRouter(destRouterAddress, destChain)
	require.NoError(t, err)

	offRampAddress, _, _, err := evm_2_evm_offramp.DeployEVM2EVMOffRamp(
		destUser,
		destChain,
		evm_2_evm_offramp.EVM2EVMOffRampStaticConfig{
			CommitStore:         commitStore.Address(),
			ChainSelector:       destChainSelector,
			SourceChainSelector: sourceChainSelector,
			OnRamp:              onRampAddress,
			PrevOffRamp:         common.HexToAddress(""),
			ArmProxy:            armProxyDestAddress,
		},
		[]common.Address{sourceLinkTokenAddress},
		[]common.Address{destPoolAddress},
		evm_2_evm_offramp.RateLimiterConfig{
			IsEnabled: true,
			Capacity:  LinkUSDValue(100),
			Rate:      LinkUSDValue(1),
		},
	)
	require.NoError(t, err)
	offRamp, err := evm_2_evm_offramp.NewEVM2EVMOffRamp(offRampAddress, destChain)
	require.NoError(t, err)
	_, err = destPool.ApplyRampUpdates(destUser,
		[]lock_release_token_pool.TokenPoolRampUpdate{},
		[]lock_release_token_pool.TokenPoolRampUpdate{{Ramp: offRampAddress, Allowed: true,
			RateLimiterConfig: lock_release_token_pool.RateLimiterConfig{
				IsEnabled: true,
				Capacity:  HundredLink,
				Rate:      big.NewInt(1e18),
			},
		}},
	)
	require.NoError(t, err)
	destChain.Commit()
	_, err = destPriceRegistry.ApplyPriceUpdatersUpdates(destUser, []common.Address{commitStoreAddress}, []common.Address{})
	require.NoError(t, err)
	_, err = destRouter.ApplyRampUpdates(destUser, nil,
		nil, []router.RouterOffRamp{{SourceChainSelector: sourceChainSelector, OffRamp: offRampAddress}})
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
			ChainID:           sourceChainID,
			ChainSelector:     sourceChainSelector,
			User:              sourceUser,
			Chain:             sourceChain,
			LinkToken:         sourceLinkToken,
			Pool:              sourcePool,
			CustomPool:        nil,
			CustomToken:       sourceCustomToken,
			ARM:               sourceARM,
			ARMProxy:          sourceARMProxy,
			PriceRegistry:     srcPriceRegistry,
			WrappedNative:     sourceWrapped,
			WrappedNativePool: sourceWeth9Pool,
		},
		Router: sourceRouter,
		OnRamp: onRamp,
	}
	dest := DestinationChain{
		Common: Common{
			ChainID:           destChainID,
			ChainSelector:     destChainSelector,
			User:              destUser,
			Chain:             destChain,
			LinkToken:         destLinkToken,
			Pool:              destPool,
			CustomPool:        nil,
			CustomToken:       destCustomToken,
			ARM:               destARM,
			ARMProxy:          destARMProxy,
			PriceRegistry:     destPriceRegistry,
			WrappedNative:     destWrapped,
			WrappedNativePool: destWrappedPool,
		},
		CommitStoreHelper: commitStoreHelper,
		CommitStore:       commitStore,
		Router:            destRouter,
		OffRamp:           offRamp,
		Receivers:         []MaybeRevertReceiver{{Receiver: revertingMessageReceiver1, Strict: false}, {Receiver: revertingMessageReceiver2, Strict: true}},
	}

	return CCIPContracts{
		Source: source,
		Dest:   dest,
	}
}

func (c *CCIPContracts) SendRequest(t *testing.T, msg router.ClientEVM2AnyMessage) *types.Transaction {
	tx, err := c.Source.Router.CcipSend(c.Source.User, c.Dest.ChainSelector, msg)
	require.NoError(t, err)
	ConfirmTxs(t, []*types.Transaction{tx}, c.Source.Chain)
	return tx
}

func (c *CCIPContracts) AssertExecState(t *testing.T, log logpoller.Log, state abihelpers.MessageExecutionState, offRampOpts ...common.Address) {
	var offRamp *evm_2_evm_offramp.EVM2EVMOffRamp
	var err error
	if len(offRampOpts) > 0 {
		offRamp, err = evm_2_evm_offramp.NewEVM2EVMOffRamp(offRampOpts[0], c.Dest.Chain)
		require.NoError(t, err)
	} else {
		require.NotNil(t, c.Dest.OffRamp, "no offRamp configured")
		offRamp = c.Dest.OffRamp
	}
	executionStateChanged, err := offRamp.ParseExecutionStateChanged(log.ToGethLog())
	require.NoError(t, err)
	if abihelpers.MessageExecutionState(executionStateChanged.State) != state {
		t.Log("Execution failed")
		t.Fail()
	}
}

func GetEVMExtraArgsV1(gasLimit *big.Int, strict bool) ([]byte, error) {
	EVMV1Tag := []byte{0x97, 0xa6, 0x57, 0xc9}

	encodedArgs, err := utils.ABIEncode(`[{"type":"uint256"},{"type":"bool"}]`, gasLimit, strict)
	if err != nil {
		return nil, err
	}

	return append(EVMV1Tag, encodedArgs...), nil
}

type ManualExecArgs struct {
	SourceChainID, DestChainID uint64
	DestUser                   *bind.TransactOpts
	SourceChain, DestChain     bind.ContractBackend
	SourceStartBlock           *big.Int // the block in/after which failed ccip-send transaction was triggered
	destStartBlock             uint64   // the start block for filtering ReportAccepted event (including the failed seq num)
	// in destination chain. if not provided to be derived by ApproxDestStartBlock method
	DestLatestBlockNum uint64 // current block number in destination
	DestDeployedAt     uint64 // destination block number for the initial destination contract deployment.
	// Can be any number before the tx was reverted in destination chain. Preferably this needs to be set up with
	// a value greater than zero to avoid performance issue in locating approximate destination block
	SendReqLogIndex uint   // log index of the CCIPSendRequested log in source chain
	SendReqTxHash   string // tx hash of the ccip-send transaction for which execution was reverted
	CommitStore     string
	OnRamp          string
	OffRamp         string
	seqNr           uint64
}

// ApproxDestStartBlock attempts to locate a block in destination chain with timestamp closest to the timestamp of the block
// in source chain in which ccip-send transaction was included
// it uses binary search to locate the block with the closest timestamp
// if the block located has a timestamp greater than the timestamp of mentioned source block
// it just returns the first block found with lesser timestamp of the source block
// providing a value of args.DestDeployedAt ensures better performance by reducing the range of block numbers to be traversed
func (args *ManualExecArgs) ApproxDestStartBlock() error {
	sourceBlockHdr, err := args.SourceChain.HeaderByNumber(context.Background(), args.SourceStartBlock)
	if err != nil {
		return err
	}
	sendTxTime := sourceBlockHdr.Time
	maxBlockNum := args.DestLatestBlockNum
	// setting this to an approx value of 1000 considering destination chain would have at least 1000 blocks before the transaction started
	minBlockNum := args.DestDeployedAt
	closestBlockNum := uint64(math.Floor((float64(maxBlockNum) + float64(minBlockNum)) / 2))
	var closestBlockHdr *types.Header
	closestBlockHdr, err = args.DestChain.HeaderByNumber(context.Background(), big.NewInt(int64(closestBlockNum)))
	if err != nil {
		return err
	}
	// to reduce the number of RPC calls increase the value of blockOffset
	blockOffset := uint64(10)
	for {
		blockNum := closestBlockHdr.Number.Uint64()
		if minBlockNum > maxBlockNum {
			break
		}
		timeDiff := math.Abs(float64(closestBlockHdr.Time - sendTxTime))
		// break if the difference in timestamp is lesser than 1 minute
		if timeDiff < 60 {
			break
		} else if closestBlockHdr.Time > sendTxTime {
			maxBlockNum = blockNum - 1
		} else {
			minBlockNum = blockNum + 1
		}
		closestBlockNum = uint64(math.Floor((float64(maxBlockNum) + float64(minBlockNum)) / 2))
		closestBlockHdr, err = args.DestChain.HeaderByNumber(context.Background(), big.NewInt(int64(closestBlockNum)))
		if err != nil {
			return err
		}
	}

	for closestBlockHdr.Time > sendTxTime {
		closestBlockNum = closestBlockNum - blockOffset
		if closestBlockNum <= 0 {
			return fmt.Errorf("approx destination blocknumber not found")
		}
		closestBlockHdr, err = args.DestChain.HeaderByNumber(context.Background(), big.NewInt(int64(closestBlockNum)))
		if err != nil {
			return err
		}
	}
	args.destStartBlock = closestBlockHdr.Number.Uint64()
	fmt.Println("using approx destination start block number", args.destStartBlock)
	return nil
}

func (args *ManualExecArgs) FindSeqNrFromCCIPSendRequested() (uint64, error) {
	var seqNr uint64
	onRampContract, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(common.HexToAddress(args.OnRamp), args.SourceChain)
	if err != nil {
		return seqNr, err
	}
	iterator, err := onRampContract.FilterCCIPSendRequested(&bind.FilterOpts{
		Start: args.SourceStartBlock.Uint64(),
	})
	if err != nil {
		return seqNr, err
	}
	for iterator.Next() {
		if iterator.Event.Raw.Index == args.SendReqLogIndex &&
			iterator.Event.Raw.TxHash.Hex() == args.SendReqTxHash {
			seqNr = iterator.Event.Message.SequenceNumber
			break
		}
	}
	if seqNr == 0 {
		return seqNr,
			fmt.Errorf("no CCIPSendRequested logs found for logIndex %d starting from block number %d", args.SendReqLogIndex, args.SourceStartBlock)
	}
	return seqNr, nil
}

func (args *ManualExecArgs) ExecuteManually() (*types.Transaction, error) {
	if args.SourceChainID == 0 ||
		args.DestChainID == 0 ||
		args.DestUser == nil {
		return nil, fmt.Errorf("chain ids and owners are mandatory for source and dest chain")
	}
	if !common.IsHexAddress(args.CommitStore) ||
		!common.IsHexAddress(args.OffRamp) ||
		!common.IsHexAddress(args.OnRamp) {
		return nil, fmt.Errorf("contract addresses must be valid hex address")
	}
	if args.SendReqTxHash == "" || args.SendReqLogIndex < 1 {
		return nil, fmt.Errorf("log index for CCIPSendRequested event and tx hash of ccip-send request are required")
	}
	if args.SourceStartBlock == nil {
		return nil, fmt.Errorf("must provide the value of source block in/after which ccip-send tx was included")
	}
	// locate seq nr from CCIPSendRequested log
	seqNr, err := args.FindSeqNrFromCCIPSendRequested()
	if err != nil {
		return nil, err
	}
	commitStore, err := commit_store.NewCommitStore(common.HexToAddress(args.CommitStore), args.DestChain)
	if err != nil {
		return nil, err
	}
	if args.destStartBlock < 1 {
		err = args.ApproxDestStartBlock()
		if err != nil {
			return nil, err
		}
	}
	iterator, err := commitStore.FilterReportAccepted(&bind.FilterOpts{Start: args.destStartBlock})
	if err != nil {
		return nil, err
	}

	var commitReport *commit_store.CommitStoreCommitReport
	for iterator.Next() {
		if iterator.Event.Report.Interval.Min <= seqNr && iterator.Event.Report.Interval.Max >= seqNr {
			commitReport = &iterator.Event.Report
			fmt.Println("Found root")
			break
		}
	}
	if commitReport == nil {
		return nil, fmt.Errorf("unable to find seq num %d in commit report", seqNr)
	}
	args.seqNr = seqNr
	return args.execute(commitReport)
}

func (args *ManualExecArgs) execute(report *commit_store.CommitStoreCommitReport) (*types.Transaction, error) {
	log.Info().Msg("Executing request manually")
	seqNr := args.seqNr
	// Build a merkle tree for the report
	mctx := hashlib.NewKeccakCtx()
	leafHasher := ccipdata.NewLeafHasherV1_2_0(args.SourceChainID, args.DestChainID, common.HexToAddress(args.OnRamp), mctx, &evm_2_evm_onramp.EVM2EVMOnRamp{})
	onRampContract, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(common.HexToAddress(args.OnRamp), args.SourceChain)
	if err != nil {
		return nil, err
	}
	var leaves [][32]byte
	var curr, prove int
	var msgs []evm_2_evm_offramp.InternalEVM2EVMMessage
	var manualExecGasLimits []*big.Int
	var tokenData [][][]byte
	sendRequestedIterator, err := onRampContract.FilterCCIPSendRequested(&bind.FilterOpts{
		Start: args.SourceStartBlock.Uint64(),
	})
	if err != nil {
		return nil, err
	}
	for sendRequestedIterator.Next() {
		if sendRequestedIterator.Event.Message.SequenceNumber <= report.Interval.Max &&
			sendRequestedIterator.Event.Message.SequenceNumber >= report.Interval.Min {
			fmt.Println("Found seq num", sendRequestedIterator.Event.Message.SequenceNumber, report.Interval)
			hash, err2 := leafHasher.HashLeaf(sendRequestedIterator.Event.Raw)
			if err2 != nil {
				return nil, err2
			}
			leaves = append(leaves, hash)
			if sendRequestedIterator.Event.Message.SequenceNumber == seqNr {
				fmt.Printf("Found proving %d %+v\n", curr, sendRequestedIterator.Event.Message)
				msg, err2 := ccipdata.DecodeOffRampMessageV1_2_0(sendRequestedIterator.Event.Raw.Data)
				if err2 != nil {
					return nil, err2
				}
				msgs = append(msgs, *msg)
				manualExecGasLimits = append(manualExecGasLimits, msg.GasLimit)
				var msgTokenData [][]byte
				for range sendRequestedIterator.Event.Message.TokenAmounts {
					msgTokenData = append(msgTokenData, []byte{})
				}

				tokenData = append(tokenData, msgTokenData)
				prove = curr
			}
			curr++
		}
	}
	sendRequestedIterator.Close()
	if msgs == nil {
		return nil, fmt.Errorf("unable to find msg with seqNr %d", seqNr)
	}
	tree, err := merklemulti.NewTree(mctx, leaves)
	if err != nil {
		return nil, err
	}
	if tree.Root() != report.MerkleRoot {
		return nil, fmt.Errorf("root doesn't match")
	}

	proof, err := tree.Prove([]int{prove})
	if err != nil {
		return nil, err
	}

	offRampProof := evm_2_evm_offramp.InternalExecutionReport{
		Messages:          msgs,
		OffchainTokenData: tokenData,
		Proofs:            proof.Hashes,
		ProofFlagBits:     abihelpers.ProofFlagsToBits(proof.SourceFlags),
	}
	offRamp, err := evm_2_evm_offramp.NewEVM2EVMOffRamp(common.HexToAddress(args.OffRamp), args.DestChain)
	if err != nil {
		return nil, err
	}
	// Execute.
	return offRamp.ManuallyExecute(args.DestUser, offRampProof, manualExecGasLimits)
}

func (c *CCIPContracts) ExecuteMessage(
	t *testing.T,
	req logpoller.Log,
	txHash common.Hash,
	destStartBlock uint64,
) uint64 {
	t.Log("Executing request manually")
	sendReqReceipt, err := c.Source.Chain.TransactionReceipt(context.Background(), txHash)
	require.NoError(t, err)
	args := ManualExecArgs{
		SourceChainID:      c.Source.ChainID,
		DestChainID:        c.Dest.ChainID,
		DestUser:           c.Dest.User,
		SourceChain:        c.Source.Chain,
		DestChain:          c.Dest.Chain,
		SourceStartBlock:   sendReqReceipt.BlockNumber,
		destStartBlock:     destStartBlock,
		DestLatestBlockNum: c.Dest.Chain.Blockchain().CurrentBlock().Number.Uint64(),
		SendReqLogIndex:    uint(req.LogIndex),
		SendReqTxHash:      txHash.String(),
		CommitStore:        c.Dest.CommitStore.Address().String(),
		OnRamp:             c.Source.OnRamp.Address().String(),
		OffRamp:            c.Dest.OffRamp.Address().String(),
	}
	tx, err := args.ExecuteManually()
	require.NoError(t, err)
	c.Dest.Chain.Commit()
	c.Source.Chain.Commit()
	rec, err := c.Dest.Chain.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.Equal(t, uint64(1), rec.Status, "manual execution failed")
	t.Logf("Manual Execution completed for seqNum %d", args.seqNr)
	return args.seqNr
}

func GetBalance(t *testing.T, chain bind.ContractBackend, tokenAddr common.Address, addr common.Address) *big.Int {
	token, err := link_token_interface.NewLinkToken(tokenAddr, chain)
	require.NoError(t, err)
	bal, err := token.BalanceOf(nil, addr)
	require.NoError(t, err)
	return bal
}
