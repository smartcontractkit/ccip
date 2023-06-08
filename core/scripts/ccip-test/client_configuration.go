package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	ocrconfighelper "github.com/smartcontractkit/libocr/offchainreporting2/confighelper"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slices"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/dione"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/shared"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/arm_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/burn_mint_erc677"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/ping_pong_demo"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/receiver_dapp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	ccipconfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/hasher"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/merklemulti"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/testhelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocrcommon"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

type ocr2Configurer interface {
	SetOCR2Config(
		opts *bind.TransactOpts,
		signers []common.Address,
		transmitters []common.Address,
		f uint8,
		onchainConfig []byte,
		offchainConfigVersion uint64,
		offchainConfig []byte,
	) (*types.Transaction, error)
}

func (client *CCIPClient) wip(t *testing.T, sourceClient *rhea.EvmDeploymentConfig, destClient *rhea.EvmDeploymentConfig) {
}

func (client *CCIPClient) applyFeeTokensUpdates(t *testing.T, sourceClient *rhea.EvmDeploymentConfig) {
	var feeTokens []common.Address

	for _, feeToken := range sourceClient.ChainConfig.FeeTokens {
		feeTokens = append(feeTokens, sourceClient.ChainConfig.SupportedTokens[feeToken].Token)
	}

	tx, err := client.Source.PriceRegistry.ApplyFeeTokensUpdates(client.Source.Owner, feeTokens, []common.Address{})
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
	client.Source.logger.Infof("Added feeTokens: %v to PriceRegistry: %s", feeTokens, client.Source.PriceRegistry.Address().Hex())
}

func (client *CCIPClient) setOnRampFeeConfig(t *testing.T, sourceClient *rhea.EvmDeploymentConfig) {
	var feeTokenConfig []evm_2_evm_onramp.EVM2EVMOnRampFeeTokenConfigArgs

	for _, feeToken := range sourceClient.ChainConfig.FeeTokens {
		feeTokenConfig = append(feeTokenConfig, evm_2_evm_onramp.EVM2EVMOnRampFeeTokenConfigArgs{
			Token:               sourceClient.ChainConfig.SupportedTokens[feeToken].Token,
			Multiplier:          1e18,
			NetworkFeeAmountUSD: big.NewInt(1e18),
			DestGasOverhead:     0,
		})
	}

	tx, err := client.Source.OnRamp.SetFeeTokenConfig(client.Source.Owner, feeTokenConfig)
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
}

func (client *CCIPClient) SetDynamicConfigOnRamp(t *testing.T) {
	config := evm_2_evm_onramp.EVM2EVMOnRampDynamicConfig{
		Router:          client.Source.Router.Address(),
		PriceRegistry:   client.Source.PriceRegistry.Address(),
		MaxDataSize:     rhea.MAX_DATA_SIZE,
		MaxTokensLength: rhea.MAX_TOKEN_LENGTH,
		MaxGasLimit:     rhea.MAX_TX_GAS_LIMIT,
		Arm:             client.Source.ARM.Address(),
	}
	tx, err := client.Source.OnRamp.SetDynamicConfig(client.Source.Owner, config)
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
}

func (client *CCIPClient) setAllowListEnabled(t *testing.T) {
	tx, err := client.Source.OnRamp.SetAllowListEnabled(client.Source.Owner, true)
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
}

func (client *CCIPClient) setAllowList(t *testing.T) {
	isEnabled, err := client.Source.OnRamp.GetAllowListEnabled(&bind.CallOpts{})
	shared.RequireNoError(t, err)
	if !isEnabled {
		client.setAllowListEnabled(t)
	}
	currentAllowList, err := client.Source.OnRamp.GetAllowList(&bind.CallOpts{})
	shared.RequireNoError(t, err)

	var toRemove []common.Address
	for _, addr := range currentAllowList {
		if !slices.Contains(client.Source.AllowList, addr) {
			toRemove = append(toRemove, addr)
		}
	}
	var toAdd []common.Address
	for _, addr := range client.Source.AllowList {
		if !slices.Contains(currentAllowList, addr) {
			toAdd = append(toAdd, addr)
		}
	}
	if len(toRemove) == 0 && len(toAdd) == 0 {
		t.Logf("Nothing to add or remove from allowlist: chainId=%d", client.Source.ChainId)
		return
	}
	t.Logf("ApplyAllowListUpdates: chainId=%d, toRemove=%v, toAdd=%v", client.Source.ChainId, toRemove, toAdd)

	tx, err := client.Source.OnRamp.ApplyAllowListUpdates(client.Source.Owner, toRemove, toAdd)
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
}

func (client *CCIPClient) setRateLimiterConfig(t *testing.T) {
	tx, err := client.Source.OnRamp.SetRateLimiterConfig(client.Source.Owner, evm_2_evm_onramp.RateLimiterConfig{
		Capacity: rhea.UsdToRateLimitValue(rhea.RATE_LIMIT_CAPACITY_DOLLAR),
		Rate:     rhea.UsdToRateLimitValue(rhea.RATE_LIMIT_RATE_DOLLAR),
	})
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)

	tx, err = client.Dest.OffRamp.SetRateLimiterConfig(client.Dest.Owner, evm_2_evm_offramp.RateLimiterConfig{
		Capacity: rhea.UsdToRateLimitValue(rhea.RATE_LIMIT_CAPACITY_DOLLAR),
		Rate:     rhea.UsdToRateLimitValue(rhea.RATE_LIMIT_RATE_DOLLAR),
	})
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
}

func (client *CCIPClient) startPingPong(t *testing.T) {
	tx, err := client.Source.PingPongDapp.StartPingPong(client.Source.Owner)
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
	client.Source.logger.Infof("Ping pong started in tx %s", helpers.ExplorerLink(int64(client.Source.ChainId), tx.Hash()))
}

func (client *CCIPClient) setPingPongPaused(t *testing.T, paused bool) {
	tx, err := client.Source.PingPongDapp.SetPaused(client.Source.Owner, paused)
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
}

func (client *CCIPClient) reemitEvents(t *testing.T, destConfig rhea.EvmDeploymentConfig) {
	currentOnRamp, err := client.Source.Router.GetOnRamp(nil, client.Dest.ChainSelector)
	shared.RequireNoError(t, err)
	require.NotEqual(t, common.Address{}, currentOnRamp)
	// reemit Router.OnRampSet (feeds Atlas's `ccip_onramps` table)
	tx, err := client.Source.Router.ApplyRampUpdates(
		client.Source.Owner,
		[]router.RouterOnRamp{{DestChainSelector: client.Dest.ChainSelector, OnRamp: currentOnRamp}},
		[]router.RouterOffRamp{}, []router.RouterOffRamp{},
	)
	shared.RequireNoError(t, err)
	client.Source.logger.Infof("Router.OnRampSet %v for destChainSelector=%d in tx %s", currentOnRamp, client.Dest.ChainSelector, helpers.ExplorerLink(int64(client.Source.ChainId), tx.Hash()))
	err = shared.WaitForMined(client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)

	// reemit OffRamp.ConfigSet (feeds Atlas's `ccip_offramps` table)
	offRampConfigIt, err := client.Dest.OffRamp.FilterConfigSet0(&bind.FilterOpts{Context: context.Background(), Start: destConfig.LaneConfig.DeploySettings.DeployedAtBlock - 100})
	shared.RequireNoError(t, err)
	var offRampConfigEvent *evm_2_evm_offramp.EVM2EVMOffRampConfigSet0
	for offRampConfigIt.Next() {
		offRampConfigEvent = offRampConfigIt.Event
	}
	require.NotEqual(t, nil, offRampConfigEvent)
	tx, err = client.Dest.OffRamp.SetOCR2Config(
		client.Dest.Owner,
		offRampConfigEvent.Signers,
		offRampConfigEvent.Transmitters,
		offRampConfigEvent.F,
		offRampConfigEvent.OnchainConfig,
		offRampConfigEvent.OffchainConfigVersion,
		offRampConfigEvent.OffchainConfig,
	)
	shared.RequireNoError(t, err)
	client.Dest.logger.Infof("OffRamp.SetOCR2Config %+v in tx %s", offRampConfigEvent, helpers.ExplorerLink(int64(client.Dest.ChainId), tx.Hash()))
	err = shared.WaitForMined(client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)

	// reemit CommitStore.ConfigSet (feeds Atlas's `ccip_arms` table)
	commitStoreConfigIt, err := client.Dest.CommitStore.FilterConfigSet0(&bind.FilterOpts{Context: context.Background(), Start: destConfig.LaneConfig.DeploySettings.DeployedAtBlock - 100})
	shared.RequireNoError(t, err)
	var commitStoreConfigEvent *commit_store.CommitStoreConfigSet0
	for commitStoreConfigIt.Next() {
		commitStoreConfigEvent = commitStoreConfigIt.Event
	}
	require.NotEqual(t, nil, commitStoreConfigEvent)
	tx, err = client.Dest.CommitStore.SetOCR2Config(
		client.Dest.Owner,
		commitStoreConfigEvent.Signers,
		commitStoreConfigEvent.Transmitters,
		commitStoreConfigEvent.F,
		commitStoreConfigEvent.OnchainConfig,
		commitStoreConfigEvent.OffchainConfigVersion,
		commitStoreConfigEvent.OffchainConfig,
	)
	shared.RequireNoError(t, err)
	client.Dest.logger.Infof("CommitStore.SetOCR2Config %+v in tx %s", commitStoreConfigEvent, helpers.ExplorerLink(int64(client.Dest.ChainId), tx.Hash()))
	err = shared.WaitForMined(client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
}

// Uncurses the ARM contract on the source chain.
func (client *CCIPClient) uncurseSourceARM(t *testing.T) {
	owner := client.Source.Owner
	arm := client.Source.ARM

	configDetails, err := arm.GetConfigDetails(nil)
	shared.RequireNoError(t, err)

	// Quite the hack, doesn't look at the curses hash, suffers from TOCTOU, etc.
	var forceUncurseRecords []arm_contract.ARMUnvoteToCurseRecord
	for _, voter := range configDetails.Config.Voters {
		forceUncurseRecords = append(forceUncurseRecords, arm_contract.ARMUnvoteToCurseRecord{
			CurseVoteAddr: voter.CurseVoteAddr,
			CursesHash:    [32]byte{},
			ForceUnvote:   true,
		})
	}

	tx, err := arm.OwnerUnvoteToCurse(owner, forceUncurseRecords)
	shared.RequireNoError(t, err)
	client.Source.logger.Infof("ARM.OwnerUnvoteToCurse %+v in tx %s", forceUncurseRecords, helpers.ExplorerLink(int64(client.Source.ChainId), tx.Hash()))
	err = shared.WaitForMined(client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
}

type Client struct {
	Owner            *bind.TransactOpts
	Users            []*bind.TransactOpts
	Client           *ethclient.Client
	ChainId          uint64
	ChainSelector    uint64
	LinkToken        *burn_mint_erc677.BurnMintERC677
	LinkTokenAddress common.Address
	WrappedNative    *burn_mint_erc677.BurnMintERC677
	SupportedTokens  map[rhea.Token]EVMBridgedToken
	PingPongDapp     *ping_pong_demo.PingPongDemo
	ARM              *arm_contract.ARMContract
	PriceRegistry    *price_registry.PriceRegistry
	Router           *router.Router
	TunableValues    rhea.TunableChainValues
	AllowList        []common.Address
	logger           logger.Logger
	t                *testing.T
}

type EVMBridgedToken struct {
	Token                common.Address
	Pool                 *lock_release_token_pool.LockReleaseTokenPool
	Price                *big.Int
	PriceFeedsAggregator common.Address
	rhea.TokenPoolType
}

type SourceClient struct {
	Client
	OnRamp *evm_2_evm_onramp.EVM2EVMOnRamp
}

func NewSourceClient(t *testing.T, config rhea.EvmConfig, laneConfig rhea.EVMLaneConfig) SourceClient {
	LinkToken, err := burn_mint_erc677.NewBurnMintERC677(config.ChainConfig.SupportedTokens[rhea.LINK].Token, config.Client)
	shared.RequireNoError(t, err)

	wrappedNative, err := burn_mint_erc677.NewBurnMintERC677(config.ChainConfig.SupportedTokens[config.ChainConfig.WrappedNative].Token, config.Client)
	shared.RequireNoError(t, err)

	supportedTokens := map[rhea.Token]EVMBridgedToken{}
	for token, tokenConfig := range config.ChainConfig.SupportedTokens {
		tokenPool, err2 := lock_release_token_pool.NewLockReleaseTokenPool(tokenConfig.Pool, config.Client)
		require.NoError(t, err2)
		supportedTokens[token] = EVMBridgedToken{
			Token:         tokenConfig.Token,
			Pool:          tokenPool,
			Price:         tokenConfig.Price,
			TokenPoolType: tokenConfig.TokenPoolType,
		}
	}

	arm, err := arm_contract.NewARMContract(config.ChainConfig.ARM, config.Client)
	shared.RequireNoError(t, err)
	onRamp, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(laneConfig.OnRamp, config.Client)
	shared.RequireNoError(t, err)
	router, err := router.NewRouter(config.ChainConfig.Router, config.Client)
	shared.RequireNoError(t, err)
	pingPongDapp, err := ping_pong_demo.NewPingPongDemo(laneConfig.PingPongDapp, config.Client)
	shared.RequireNoError(t, err)
	priceRegistry, err := price_registry.NewPriceRegistry(config.ChainConfig.PriceRegistry, config.Client)
	shared.RequireNoError(t, err)

	return SourceClient{
		Client: Client{
			Client:           config.Client,
			ChainId:          config.ChainConfig.EvmChainId,
			ChainSelector:    rhea.GetCCIPChainSelector(config.ChainConfig.EvmChainId),
			LinkTokenAddress: config.ChainConfig.SupportedTokens[rhea.LINK].Token,
			LinkToken:        LinkToken,
			WrappedNative:    wrappedNative,
			ARM:              arm,
			PriceRegistry:    priceRegistry,
			SupportedTokens:  supportedTokens,
			PingPongDapp:     pingPongDapp,
			Router:           router,
			AllowList:        config.ChainConfig.AllowList,
			TunableValues:    config.ChainConfig.TunableChainValues,
			logger:           config.Logger,
			t:                t,
		},
		OnRamp: onRamp,
	}
}

type DestClient struct {
	Client
	CommitStore  *commit_store.CommitStore
	ReceiverDapp *receiver_dapp.ReceiverDapp
	OffRamp      *evm_2_evm_offramp.EVM2EVMOffRamp
}

func NewDestinationClient(t *testing.T, config rhea.EvmConfig, laneConfig rhea.EVMLaneConfig) DestClient {
	linkToken, err := burn_mint_erc677.NewBurnMintERC677(config.ChainConfig.SupportedTokens[rhea.LINK].Token, config.Client)
	shared.RequireNoError(t, err)
	wrappedNative, err := burn_mint_erc677.NewBurnMintERC677(config.ChainConfig.SupportedTokens[config.ChainConfig.WrappedNative].Token, config.Client)
	shared.RequireNoError(t, err)

	supportedTokens := map[rhea.Token]EVMBridgedToken{}
	for token, tokenConfig := range config.ChainConfig.SupportedTokens {
		tokenPool, err2 := lock_release_token_pool.NewLockReleaseTokenPool(tokenConfig.Pool, config.Client)
		require.NoError(t, err2)
		supportedTokens[token] = EVMBridgedToken{
			Token:         tokenConfig.Token,
			Pool:          tokenPool,
			Price:         tokenConfig.Price,
			TokenPoolType: tokenConfig.TokenPoolType,
		}
	}

	arm, err := arm_contract.NewARMContract(config.ChainConfig.ARM, config.Client)
	shared.RequireNoError(t, err)
	commitStore, err := commit_store.NewCommitStore(laneConfig.CommitStore, config.Client)
	shared.RequireNoError(t, err)
	offRamp, err := evm_2_evm_offramp.NewEVM2EVMOffRamp(laneConfig.OffRamp, config.Client)
	shared.RequireNoError(t, err)
	receiverDapp, err := receiver_dapp.NewReceiverDapp(laneConfig.ReceiverDapp, config.Client)
	shared.RequireNoError(t, err)
	router, err := router.NewRouter(config.ChainConfig.Router, config.Client)
	shared.RequireNoError(t, err)
	pingPongDapp, err := ping_pong_demo.NewPingPongDemo(laneConfig.PingPongDapp, config.Client)
	shared.RequireNoError(t, err)
	priceRegistry, err := price_registry.NewPriceRegistry(config.ChainConfig.PriceRegistry, config.Client)
	shared.RequireNoError(t, err)

	return DestClient{
		Client: Client{
			Client:           config.Client,
			ChainId:          config.ChainConfig.EvmChainId,
			ChainSelector:    rhea.GetCCIPChainSelector(config.ChainConfig.EvmChainId),
			LinkTokenAddress: config.ChainConfig.SupportedTokens[rhea.LINK].Token,
			LinkToken:        linkToken,
			WrappedNative:    wrappedNative,
			SupportedTokens:  supportedTokens,
			PingPongDapp:     pingPongDapp,
			ARM:              arm,
			PriceRegistry:    priceRegistry,
			logger:           config.Logger,
			Router:           router,
			TunableValues:    config.ChainConfig.TunableChainValues,
			AllowList:        config.ChainConfig.AllowList,
			t:                t,
		},
		CommitStore:  commitStore,
		ReceiverDapp: receiverDapp,
		OffRamp:      offRamp,
	}
}

// CCIPClient contains a source chain and destination chain client and implements many methods
// that are useful for testing CCIP functionality on chain.
type CCIPClient struct {
	Source SourceClient
	Dest   DestClient
}

// NewCcipClient returns a new CCIPClient with initialised source and destination clients.
func NewCcipClient(
	t *testing.T,
	sourceConfig rhea.EvmDeploymentConfig,
	destConfig rhea.EvmDeploymentConfig,
	ownerKey string,
	seedKey string,
) CCIPClient {
	return NewCcipClientByLane(t, sourceConfig.OnlyEvmConfig(), sourceConfig.LaneConfig, destConfig.OnlyEvmConfig(), destConfig.LaneConfig, ownerKey, seedKey)
}

// NewUpgradeLaneCcipClient creates client that point to Upgrade Lanes and Upgrade Router
func NewUpgradeLaneCcipClient(
	t *testing.T,
	sourceConfig rhea.EvmDeploymentConfig,
	destConfig rhea.EvmDeploymentConfig,
	ownerKey string,
	seedKey string,
) CCIPClient {
	client := NewCcipClientByLane(t, sourceConfig.OnlyEvmConfig(), sourceConfig.UpgradeLaneConfig, destConfig.OnlyEvmConfig(), destConfig.UpgradeLaneConfig, ownerKey, seedKey)
	applyUpgradeRouterAddress(t, sourceConfig, &client.Source.Client)
	applyUpgradeRouterAddress(t, destConfig, &client.Dest.Client)
	return client
}

func NewCcipClientByLane(
	t *testing.T,
	sourceConfig rhea.EvmConfig,
	sourceLane rhea.EVMLaneConfig,
	destConfig rhea.EvmConfig,
	destLane rhea.EVMLaneConfig,
	ownerKey string,
	seedKey string,
) CCIPClient {
	source := NewSourceClient(t, sourceConfig, sourceLane)
	source.SetOwnerAndUsers(t, ownerKey, seedKey, sourceConfig.ChainConfig.GasSettings)
	dest := NewDestinationClient(t, destConfig, destLane)
	dest.SetOwnerAndUsers(t, ownerKey, seedKey, destConfig.ChainConfig.GasSettings)

	return CCIPClient{
		Source: source,
		Dest:   dest,
	}
}

func applyUpgradeRouterAddress(t *testing.T, config rhea.EvmDeploymentConfig, client *Client) {
	upgradeRouter, err := router.NewRouter(config.ChainConfig.UpgradeRouter, config.Client)
	shared.RequireNoError(t, err)
	client.Router = upgradeRouter
}

func GetSetupChain(t *testing.T, ownerPrivateKey string, chain rhea.EvmDeploymentConfig) *rhea.EvmDeploymentConfig {
	chain.SetupChain(t, ownerPrivateKey)
	return &chain
}

// SetOwnerAndUsers sets the owner and 10 users on a given client. It also set the proper
// gas parameters on these users.
func (client *Client) SetOwnerAndUsers(t *testing.T, ownerPrivateKey string, seedKey string, gasSettings rhea.EVMGasSettings) {
	client.Owner = rhea.GetOwner(t, ownerPrivateKey, client.ChainId, gasSettings)

	var users []*bind.TransactOpts
	seedKeyWithoutFirstChar := seedKey[1:]
	fmt.Println("--- Addresses of the seed key")
	for i := 0; i <= 9; i++ {
		_, err := hex.DecodeString(strconv.Itoa(i) + seedKeyWithoutFirstChar)
		require.NoError(t, err)
		key, err := crypto.HexToECDSA(strconv.Itoa(i) + seedKeyWithoutFirstChar)
		require.NoError(t, err)
		user, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(0).SetUint64(client.ChainId))
		require.NoError(t, err)
		rhea.SetGasFees(user, gasSettings)
		users = append(users, user)
		fmt.Println(user.From.Hex())
	}
	fmt.Println("---")

	client.Users = users
}

func (client *Client) ApproveLinkFrom(t *testing.T, user *bind.TransactOpts, approvedFor common.Address, amount *big.Int) {
	client.logger.Warnf("Approving %d link for %s", amount.Int64(), approvedFor.Hex())
	tx, err := client.LinkToken.Approve(user, approvedFor, amount)
	require.NoError(t, err)

	err = shared.WaitForMined(client.logger, client.Client, tx.Hash(), true)
	require.NoError(t, err)
	client.logger.Warnf("Link approved %s", helpers.ExplorerLink(int64(client.ChainId), tx.Hash()))
}

func (client *Client) ApproveLink(t *testing.T, approvedFor common.Address, amount *big.Int) {
	client.ApproveLinkFrom(t, client.Owner, approvedFor, amount)
}

func (client *CCIPClient) WaitForCommit(DestBlockNum uint64) error {
	client.Dest.logger.Infof("Waiting for commit")

	commitEvent := make(chan *commit_store.CommitStoreReportAccepted)
	sub, err := client.Dest.CommitStore.WatchReportAccepted(
		&bind.WatchOpts{
			Context: context.Background(),
			Start:   &DestBlockNum,
		},
		commitEvent,
	)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	select {
	case event := <-commitEvent:
		client.Dest.logger.Infof("Commit in tx %s", helpers.ExplorerLink(int64(client.Dest.ChainId), event.Raw.TxHash))
		return nil
	case err = <-sub.Err():
		return err
	}
}

func (client *CCIPClient) WaitForExecution(DestBlockNum uint64, sequenceNumber uint64) error {
	client.Dest.logger.Infof("Waiting for execution...")

	events := make(chan *evm_2_evm_offramp.EVM2EVMOffRampExecutionStateChanged)
	sub, err := client.Dest.OffRamp.WatchExecutionStateChanged(
		&bind.WatchOpts{
			Context: context.Background(),
			Start:   &DestBlockNum,
		},
		events,
		[]uint64{sequenceNumber},
		[][32]byte{})
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	select {
	case event := <-events:
		client.Dest.logger.Infof("Execution in tx %s", helpers.ExplorerLink(int64(client.Dest.ChainId), event.Raw.TxHash))
		return nil
	case err = <-sub.Err():
		return err
	}
}

func (client *CCIPClient) ExecuteManually(t *testing.T, destDeploySettings *rhea.EvmDeploymentConfig) {
	txHash := os.Getenv("CCIP_SEND_TX")
	require.NotEmpty(t, txHash, "must set txhash for which manual execution is needed")
	require.NotEmpty(t, os.Getenv("LOG_INDEX"), "must set log index of CCIPSendRequested event emitted by onRamp contract")
	logIndex, err := strconv.ParseUint(os.Getenv("LOG_INDEX"), 10, 64)
	require.NoError(t, err, "log index of CCIPSendRequested event must be an int")
	require.NotEmpty(t, os.Getenv("SOURCE_START_BLOCK"), "must set the block number of source chain prior to ccip-send tx")

	sendReqReceipt, err := client.Source.Client.Client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	require.NoErrorf(t, err, "fetching receipt for ccip-send tx - ", txHash)
	destLatestBlock, err := client.Dest.Client.Client.BlockNumber(context.Background())
	require.NoError(t, err, "fetching latest block number in destination chain")
	args := testhelpers.ManualExecArgs{
		SourceChainID:      client.Source.ChainId,
		DestChainID:        client.Dest.ChainId,
		DestUser:           client.Dest.Owner,
		SourceChain:        client.Source.Client.Client,
		DestChain:          client.Dest.Client.Client,
		SourceStartBlock:   sendReqReceipt.BlockNumber,
		DestLatestBlockNum: destLatestBlock,
		DestDeployedAt:     destDeploySettings.LaneConfig.DeploySettings.DeployedAtBlock,
		SendReqLogIndex:    uint(logIndex),
		SendReqTxHash:      txHash,
		CommitStore:        client.Dest.CommitStore.Address().String(),
		OnRamp:             client.Source.OnRamp.Address().String(),
		OffRamp:            client.Dest.OffRamp.Address().String(),
	}
	tx, err := args.ExecuteManually()
	require.NoErrorf(t, err, "executing manually for tx %s in source chain", txHash)
	fmt.Println("Manual execution successful", client.Dest.Owner.From, tx.Hash(), err)
}

//func (client CCIPClient) ExternalExecutionHappyPath(t *testing.T) {
//	ctx := context.Background()
//	offrampBlockNumber := GetCurrentBlockNumber(client.Dest.Client.Client)
//	onrampBlockNumber := GetCurrentBlockNumber(client.Source.Client.Client)
//
//	amount, _ := new(big.Int).SetString("10", 10)
//	client.Source.ApproveLink(t, client.Source.OnRamRouter.Address(), amount)
//
//	onrampRequest := client.SendToOnrampWithExecution(client.Source, client.Source.Owner, client.Dest.Owner.From, amount, common.HexToAddress("0x0000000000000000000000000000000000000000"))
//	sequenceNumber := onrampRequest.Message.SequenceNumber
//
//	// Gets the report that our transaction is included in
//	client.Dest.logger.Info("Getting report")
//	report, err := client.GetReportForSequenceNumber(ctx, sequenceNumber, offrampBlockNumber)
//	require.NoError(t, err)
//
//	// Get all requests included in the given report
//	client.Dest.logger.Info("Getting recent cross chain requests")
//	requests := client.GetCrossChainSendRequestsForRange(ctx, t, report, onrampBlockNumber)
//
//	// Generate the proof
//	client.Dest.logger.Info("Generating proof")
//	proof := client.ValidateMerkleRoot(t, onrampRequest, requests, report)
//
//	// Execute the transaction on the offramp
//	client.Dest.logger.Info("Executing offramp TX")
//	tx, err := client.ExecuteOffRampTransaction(t, proof, onrampRequest.Raw.Data)
//	require.NoError(t, err)
//
//	WaitForMined(t, client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
//	client.Dest.logger.Infof("Cross chain tx sent %s", helpers.ExplorerLink(client.Dest.ChainId.Int64(), tx.Hash()))
//}

// ScalingAndBatching should scale so that we see batching on the nodes
func (client *CCIPClient) ScalingAndBatching(t *testing.T) {
	amount := big.NewInt(10)
	toAddress := common.HexToAddress("0x57359120D900fab8cE74edC2c9959b21660d3887")
	DestBlockNum := GetCurrentBlockNumber(client.Dest.Client.Client)
	var seqNum uint64

	var wg sync.WaitGroup
	for _, user := range client.Source.Users {
		wg.Add(1)
		go func(user *bind.TransactOpts) {
			defer wg.Done()
			client.Source.ApproveLinkFrom(t, user, client.Source.Router.Address(), amount)
			crossChainRequest := client.SendCrossChainMessage(client.Source, user, toAddress, amount)
			client.Source.logger.Info("Don executed tx submitted with sequence number: ", crossChainRequest.Message.SequenceNumber)
			seqNum = crossChainRequest.Message.SequenceNumber
		}(user)
	}
	wg.Wait()
	require.NoError(t, client.WaitForCommit(DestBlockNum), "waiting for commit")
	require.NoError(t, client.WaitForExecution(DestBlockNum, seqNum), "waiting for execution")
	client.Source.logger.Info("Sent 10 txs to onramp.")
}

func (client *CCIPClient) SendCrossChainMessage(source SourceClient, from *bind.TransactOpts, toAddress common.Address, amount *big.Int) *evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested {
	SourceBlockNumber := GetCurrentBlockNumber(source.Client.Client)
	token := router.ClientEVMTokenAmount{
		Token:  client.Source.LinkTokenAddress,
		Amount: amount,
	}
	extraArgsV1, err := testhelpers.GetEVMExtraArgsV1(big.NewInt(100_000), false)
	helpers.PanicErr(err)

	tx, err := source.Router.CcipSend(from, client.Dest.ChainSelector, router.ClientEVM2AnyMessage{
		Receiver:     toAddress.Bytes(),
		Data:         nil,
		TokenAmounts: []router.ClientEVMTokenAmount{token},
		FeeToken:     common.Address{},
		ExtraArgs:    extraArgsV1,
	})
	helpers.PanicErr(err)
	source.logger.Infof("Send tokens tx %s", helpers.ExplorerLink(int64(source.ChainId), tx.Hash()))

	return WaitForCrossChainSendRequest(source, SourceBlockNumber, tx.Hash())
}

func GetCurrentBlockNumber(chain *ethclient.Client) uint64 {
	blockNumber, err := chain.BlockNumber(context.Background())
	helpers.PanicErr(err)
	return blockNumber
}

func (client *CCIPClient) ValidateMerkleRoot(
	t *testing.T,
	request *evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested,
	reportRequests []*evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested,
	report commit_store.CommitStoreCommitReport,
) merklemulti.Proof[[32]byte] {
	mctx := hasher.NewKeccakCtx()
	var leafHashes [][32]byte
	for _, req := range reportRequests {
		leafHashes = append(leafHashes, mctx.Hash(req.Raw.Data))
	}

	tree, err := merklemulti.NewTree(mctx, leafHashes)
	require.NoError(t, err)

	require.Equal(t, tree.Root(), report.MerkleRoot)

	exists, err := client.Dest.CommitStore.GetMerkleRoot(nil, tree.Root())
	require.NoError(t, err)
	if exists.Uint64() < 1 {
		panic("Path is not present in the offramp")
	}
	index := request.Message.SequenceNumber - report.Interval.Min
	client.Dest.logger.Info("index is ", index)
	return tree.Prove([]int{int(index)})
}

// SendToOnrampWithExecution executes a cross chain transactions using the onramp interface.
func (client *CCIPClient) SendToOnrampWithExecution(t *testing.T, source SourceClient, from *bind.TransactOpts, toAddress common.Address, amount *big.Int) *evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested {
	SourceBlockNumber := GetCurrentBlockNumber(source.Client.Client)

	senderAndReceiver, err := utils.ABIEncode(`[{"type":"address"}, {"type":"address"}]`, source.Owner.From, source.Owner.From)
	helpers.PanicErr(err)

	extraArgsV1, err := testhelpers.GetEVMExtraArgsV1(big.NewInt(3e5), false)
	helpers.PanicErr(err)

	payload := router.ClientEVM2AnyMessage{
		TokenAmounts: []router.ClientEVMTokenAmount{},
		Receiver:     testhelpers.MustEncodeAddress(t, toAddress),
		Data:         senderAndReceiver,
		ExtraArgs:    extraArgsV1,
		FeeToken:     client.Source.LinkTokenAddress,
	}
	source.logger.Infof("Send tx with payload %+v", payload)

	tx, err := source.Router.CcipSend(from, client.Dest.ChainSelector, payload)
	if err != nil {
		t.Log(err.Error())
		printRevertReason(err, router.RouterABI)
	}
	helpers.PanicErr(err)
	source.logger.Infof("Send tokens tx %s", helpers.ExplorerLink(int64(source.ChainId), tx.Hash()))
	return WaitForCrossChainSendRequest(source, SourceBlockNumber, tx.Hash())
}

func printRevertReason(errorData interface{}, abiString string) {
	dataError := errorData.(rpc.DataError)
	data, err := hex.DecodeString(dataError.ErrorData().(string)[2:])
	helpers.PanicErr(err)
	jsonABI, err := abi.JSON(strings.NewReader(abiString))
	helpers.PanicErr(err)
	for k, abiError := range jsonABI.Errors {
		if bytes.Equal(data[:4], abiError.ID.Bytes()[:4]) {
			// Found a matching error
			v, err := abiError.Unpack(data)
			helpers.PanicErr(err)
			fmt.Printf("Error \"%v\" args \"%v\"\n", k, v)
			return
		}
	}
}

// WaitForCrossChainSendRequest checks on chain for a successful onramp send event with the given tx hash.
// If not immediately found it will keep retrying in intervals of the globally specified RetryTiming.
func WaitForCrossChainSendRequest(source SourceClient, fromBlockNum uint64, txhash common.Hash) *evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested {
	filter := bind.FilterOpts{Start: fromBlockNum}
	source.logger.Infof("Waiting for cross chain send... ")

	for {
		iterator, err := source.OnRamp.FilterCCIPSendRequested(&filter)
		helpers.PanicErr(err)
		for iterator.Next() {
			if iterator.Event.Raw.TxHash.Hex() == txhash.Hex() {
				source.logger.Infof("Cross chain send event found in tx: %s ", helpers.ExplorerLink(int64(source.ChainId), txhash))
				return iterator.Event
			}
		}
		time.Sleep(shared.RetryTiming)
	}
}

func (client *CCIPClient) SetOCR2Config(env dione.Environment) {
	verifierOCRConfig, err := client.Dest.CommitStore.LatestConfigDetails(&bind.CallOpts{})
	helpers.PanicErr(err)
	if verifierOCRConfig.BlockNumber != 0 {
		client.Dest.logger.Infof("CommitStore OCR config already found: %+v", verifierOCRConfig.ConfigDigest)
		client.Dest.logger.Infof("The new config will overwrite the current one.")
	}

	rampOCRConfig, err := client.Dest.OffRamp.LatestConfigDetails(&bind.CallOpts{})
	helpers.PanicErr(err)
	if rampOCRConfig.BlockNumber != 0 {
		client.Dest.logger.Infof("OffRamp OCR config already found: %+v", rampOCRConfig.ConfigDigest)
		client.Dest.logger.Infof("The new config will overwrite the current one.")
	}
	if client.Dest.Client.ChainId == 1337 || client.Source.Client.ChainId == 1337 {
		env = dione.Prod_Swift
	}
	don := dione.NewOfflineDON(env, client.Dest.logger)
	faults := len(don.Config.Nodes) / 3

	tx, err := client.setOCRConfig(client.Dest.CommitStore, client.getCommitStoreOffChainConfig(), client.getCommitStoreOnchainConfig(), faults, don.GenerateOracleIdentities(client.Dest.ChainId))
	helpers.PanicErr(err)
	client.Dest.logger.Infof("Config set on commitStore %s", helpers.ExplorerLink(int64(client.Dest.ChainId), tx.Hash()))

	tx, err = client.setOCRConfig(client.Dest.OffRamp, client.getOffRampOffChainConfig(), client.getOffRampOnchainConfig(), faults, don.GenerateOracleIdentities(client.Dest.ChainId))
	helpers.PanicErr(err)
	client.Dest.logger.Infof("Config set on offramp %s", helpers.ExplorerLink(int64(client.Dest.ChainId), tx.Hash()))
}

func (client *CCIPClient) getCommitStoreOffChainConfig() []byte {
	if client.Source.TunableValues.FinalityDepth == 0 || client.Dest.TunableValues.FinalityDepth == 0 {
		panic("Please set the tunable chain values")
	}

	commitPluginConfig := ccipconfig.CommitOffchainConfig{
		SourceFinalityDepth:   client.Source.TunableValues.FinalityDepth,
		FeeUpdateHeartBeat:    client.Dest.TunableValues.FeeUpdateHeartBeat,
		FeeUpdateDeviationPPB: client.Dest.TunableValues.FeeUpdateDeviationPPB,
		MaxGasPrice:           client.Dest.TunableValues.MaxGasPrice,
		InflightCacheExpiry:   client.Dest.TunableValues.InflightCacheExpiry,
	}

	encodedOffchainConfig, err := ccipconfig.EncodeOffchainConfig(commitPluginConfig)
	helpers.PanicErr(err)

	return encodedOffchainConfig
}

func (client *CCIPClient) getCommitStoreOnchainConfig() []byte {
	commitStoreOnchainConfig := ccipconfig.CommitOnchainConfig{
		PriceRegistry: client.Dest.PriceRegistry.Address(),
		Arm:           client.Dest.ARM.Address(),
	}

	encodedCommitStoreOnchainConfig, err := abihelpers.EncodeAbiStruct(commitStoreOnchainConfig)
	helpers.PanicErr(err)

	return encodedCommitStoreOnchainConfig
}

func (client *CCIPClient) getOffRampOffChainConfig() []byte {
	if client.Source.TunableValues.FinalityDepth == 0 || client.Dest.TunableValues.FinalityDepth == 0 {
		panic("Please set the tunable chain values")
	}
	execPluginConfig := ccipconfig.ExecOffchainConfig{
		SourceFinalityDepth:         client.Source.TunableValues.FinalityDepth,
		DestFinalityDepth:           client.Dest.TunableValues.FinalityDepth,
		DestOptimisticConfirmations: client.Dest.TunableValues.OptimisticConfirmations,
		BatchGasLimit:               client.Dest.TunableValues.BatchGasLimit,
		RelativeBoostPerWaitHour:    client.Dest.TunableValues.RelativeBoostPerWaitHour,
		MaxGasPrice:                 client.Dest.TunableValues.MaxGasPrice,
		InflightCacheExpiry:         client.Dest.TunableValues.InflightCacheExpiry,
		RootSnoozeTime:              client.Dest.TunableValues.RootSnoozeTime,
	}

	encodedOffRampConfig, err := ccipconfig.EncodeOffchainConfig(execPluginConfig)
	helpers.PanicErr(err)

	return encodedOffRampConfig
}

func (client *CCIPClient) getOffRampOnchainConfig() []byte {
	offRampOnchainConfig := ccipconfig.ExecOnchainConfig{
		PermissionLessExecutionThresholdSeconds: rhea.PERMISSIONLESS_EXEC_THRESHOLD_SEC,
		Router:                                  client.Dest.Router.Address(),
		Arm:                                     client.Dest.ARM.Address(),
		PriceRegistry:                           client.Dest.PriceRegistry.Address(),
		MaxTokensLength:                         rhea.MAX_TOKEN_LENGTH,
		MaxDataSize:                             rhea.MAX_DATA_SIZE,
	}

	encodedOffRampOnchainConfig, err := abihelpers.EncodeAbiStruct(offRampOnchainConfig)
	helpers.PanicErr(err)

	return encodedOffRampOnchainConfig
}

func (client *CCIPClient) setOCRConfig(ocrConf ocr2Configurer, pluginOffchainConfig []byte, onchainConfig []byte, faults int, identities []ocrconfighelper.OracleIdentityExtra) (*types.Transaction, error) {
	// Simple transmission schedule of 1 node per stage.
	// sum(transmissionSchedule) should equal number of nodes.
	var transmissionSchedule []int
	for i := 0; i < len(identities); i++ {
		transmissionSchedule = append(transmissionSchedule, 1)
	}
	signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, err := ocrconfighelper.ContractSetConfigArgsForTests(
		2*time.Minute, // deltaProgress
		5*time.Second, // deltaResend
		1*time.Minute, // deltaRound
		5*time.Second, // deltaGrace
		client.Dest.TunableValues.InflightCacheExpiry.Duration(), // deltaStage
		3,
		transmissionSchedule,
		identities,
		pluginOffchainConfig,
		100*time.Millisecond, // query not used
		35*time.Second,       // observation TODO: shorten once db issues resolved
		10*time.Second,       // report TODO: shorten once db issues resolved
		5*time.Second,        // shouldAccept
		10*time.Second,       // shouldTransmit TODO: shorten once db issues resolved
		faults,
		onchainConfig,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create args for ocr config tx")
	}
	signerAddresses, err := ocrcommon.OnchainPublicKeyToAddress(signers)
	if err != nil {
		return nil, err
	}
	transmitterAddresses, err := ocrcommon.AccountToAddress(transmitters)
	if err != nil {
		return nil, err
	}

	tx, err := ocrConf.SetOCR2Config(
		client.Dest.Owner,
		signerAddresses,
		transmitterAddresses,
		f,
		onchainConfig,
		offchainConfigVersion,
		offchainConfig,
	)
	if err != nil {
		return nil, err
	}
	err = shared.WaitForMined(client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
	if err != nil {
		return tx, err
	}

	return tx, nil
}

func (client *CCIPClient) AcceptOwnership(t *testing.T) {
	tx, err := client.Dest.CommitStore.AcceptOwnership(client.Dest.Owner)
	require.NoError(t, err)
	err = shared.WaitForMined(client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
	helpers.PanicErr(err)

	tx, err = client.Dest.OffRamp.AcceptOwnership(client.Dest.Owner)
	require.NoError(t, err)
	err = shared.WaitForMined(client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
	helpers.PanicErr(err)
}

func (client *CCIPClient) syncPoolsOnOnRamp() {
	registeredTokens, err := client.Source.OnRamp.GetSupportedTokens(&bind.CallOpts{})
	require.NoError(client.Source.t, err)

	// We only want to add tokens that are supported on both chains.
	var wantedSourceTokens []common.Address
	var wantedSourceTokenConfig []EVMBridgedToken

	for _, token := range rhea.GetAllTokens() {
		if sourceConfig, ok := client.Source.SupportedTokens[token]; ok {
			if destConfig, ok := client.Dest.SupportedTokens[token]; ok {
				if sourceConfig.TokenPoolType == rhea.FeeTokenOnly || destConfig.TokenPoolType == rhea.FeeTokenOnly {
					continue
				}
				wantedSourceTokens = append(wantedSourceTokens, sourceConfig.Token)
				wantedSourceTokenConfig = append(wantedSourceTokenConfig, sourceConfig)
				client.Source.logger.Infof("Wanted token: %s", token)
			}
		}
	}

	var poolsToRemove, poolsToAdd []evm_2_evm_onramp.InternalPoolUpdate

	// remove registered tokenPools not present in config
	for _, token := range registeredTokens {
		if !slices.Contains(wantedSourceTokens, token) {
			pool, err := client.Source.OnRamp.GetPoolBySourceToken(&bind.CallOpts{}, token)
			require.NoError(client.Source.t, err)
			poolsToRemove = append(poolsToRemove, evm_2_evm_onramp.InternalPoolUpdate{
				Token: token,
				Pool:  pool,
			})
		}
	}
	// add tokenPools present in config and not in the ramp
	for i, wantedSourceToken := range wantedSourceTokens {
		if !slices.Contains(registeredTokens, wantedSourceToken) {
			poolsToAdd = append(poolsToAdd, evm_2_evm_onramp.InternalPoolUpdate{
				Token: wantedSourceToken,
				Pool:  wantedSourceTokenConfig[i].Pool.Address(),
			})
		}
	}

	if len(poolsToAdd) > 0 || len(poolsToRemove) > 0 {
		// Pools to add should be the SECOND argument and poolsToRemove the first
		// Since our deployments are still based on a swapped order, until we deploy new onRamps
		// this order needs to be maintained to be compatible with the deployed code.
		tx, err := client.Source.OnRamp.ApplyPoolUpdates(client.Source.Owner, poolsToRemove, poolsToAdd)
		require.NoError(client.Source.t, err)
		err = shared.WaitForMined(client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
		require.NoError(client.Source.t, err)
		client.Source.logger.Infof("synced(add=%s, remove=%s) from onRamp=%s: tx=%s", poolsToAdd, poolsToRemove, client.Source.OnRamp.Address(), tx.Hash())
	}
}

func (client *CCIPClient) syncPoolsOffOnRamp() {
	registeredTokens, err := client.Dest.OffRamp.GetSupportedTokens(&bind.CallOpts{})
	require.NoError(client.Dest.t, err)

	// We only want to add tokens that are supported on both chains.
	var wantedSourceTokens []common.Address
	var wantedDestTokenConfig []EVMBridgedToken

	for _, token := range rhea.GetAllTokens() {
		if sourceConfig, ok := client.Source.SupportedTokens[token]; ok {
			if destConfig, ok := client.Dest.SupportedTokens[token]; ok {
				if sourceConfig.TokenPoolType == rhea.FeeTokenOnly || destConfig.TokenPoolType == rhea.FeeTokenOnly {
					continue
				}
				wantedSourceTokens = append(wantedSourceTokens, sourceConfig.Token)
				wantedDestTokenConfig = append(wantedDestTokenConfig, destConfig)
				client.Dest.logger.Infof("Wanted token: %s", token)
			}
		}
	}

	var poolsToRemove, poolsToAdd []evm_2_evm_offramp.InternalPoolUpdate

	// remove registered tokenPools not present in config
	for _, token := range registeredTokens {
		if !slices.Contains(wantedSourceTokens, token) {
			pool, err := client.Dest.OffRamp.GetPoolBySourceToken(&bind.CallOpts{}, token)
			require.NoError(client.Dest.t, err)
			poolsToRemove = append(poolsToRemove, evm_2_evm_offramp.InternalPoolUpdate{
				Token: token,
				Pool:  pool,
			})
		}
	}
	// add tokenPools present in config and not yet registered
	for i, wantedSourceToken := range wantedSourceTokens {
		if !slices.Contains(registeredTokens, wantedSourceToken) {
			poolsToAdd = append(poolsToAdd, evm_2_evm_offramp.InternalPoolUpdate{
				Token: wantedSourceToken,
				Pool:  wantedDestTokenConfig[i].Pool.Address(),
			})
		}
	}

	if len(poolsToAdd) > 0 || len(poolsToRemove) > 0 {
		tx, err := client.Dest.OffRamp.ApplyPoolUpdates(client.Dest.Owner, poolsToRemove, poolsToAdd)
		require.NoError(client.Dest.t, err)
		err = shared.WaitForMined(client.Dest.logger, client.Dest.Client.Client, tx.Hash(), true)
		require.NoError(client.Dest.t, err)
		client.Dest.logger.Infof("synced(add=%s, remove=%s) from offRamp=%s: tx=%s", poolsToAdd, poolsToRemove, client.Dest.OffRamp.Address(), tx.Hash())
	}
}

func syncPrices(client *Client, otherChainTokens map[rhea.Token]EVMBridgedToken) {
	// We only want to add token prices that are supported on both chains.
	var wantedTokens []common.Address
	var prices []*big.Int

	for _, token := range rhea.GetAllTokens() {
		if sourceConfig, ok := client.SupportedTokens[token]; ok {
			if _, ok := otherChainTokens[token]; ok {
				wantedTokens = append(wantedTokens, sourceConfig.Token)
				prices = append(prices, sourceConfig.Price)
				client.logger.Infof("Wanted token: %s", token)
			}
		}
	}

	if len(wantedTokens) == 0 {
		client.logger.Info("No tokens found for this lane")
		return
	}

	tokenPrices, err := client.PriceRegistry.GetTokenPrices(&bind.CallOpts{}, wantedTokens)
	require.NoError(client.t, err)
	for i, price := range prices {
		// If a price difference is found update all prices and return
		if price.Cmp(tokenPrices[i].Value) != 0 {
			priceUpdates := price_registry.InternalPriceUpdates{
				TokenPriceUpdates: []price_registry.InternalTokenPriceUpdate{},
				DestChainSelector: 0,
				UsdPerUnitGas:     big.NewInt(0),
			}

			for i, token := range wantedTokens {
				priceUpdates.TokenPriceUpdates = append(priceUpdates.TokenPriceUpdates, price_registry.InternalTokenPriceUpdate{
					SourceToken: token,
					// USD per Token.
					UsdPerToken: prices[i],
				})
			}

			tx, err2 := client.PriceRegistry.UpdatePrices(client.Owner, priceUpdates)
			require.NoError(client.t, err2)
			err2 = shared.WaitForMined(client.logger, client.Client, tx.Hash(), true)
			require.NoError(client.t, err2)
			client.logger.Infof("updatePrices(tokens=%s, prices=%s) for registry=%s: tx=%s", wantedTokens, prices, client.PriceRegistry.Address(), tx.Hash())
			return
		}
	}
	client.logger.Info("Prices already set correctly")
}

func (client *CCIPClient) syncOnRampOnPools() error {
	for tokenName, tokenConfig := range client.Source.SupportedTokens {
		// Only add tokens that are supported on both chains
		// and not marked as feeTokenOnly
		if destConfig, ok := client.Dest.SupportedTokens[tokenName]; !ok || destConfig.TokenPoolType == rhea.FeeTokenOnly || tokenConfig.TokenPoolType == rhea.FeeTokenOnly {
			continue
		}

		isOnRamp, err := tokenConfig.Pool.IsOnRamp(&bind.CallOpts{}, client.Source.OnRamp.Address())
		if err != nil {
			return errors.Wrapf(err, "failed loading onRamp data for token %s", tokenName)
		}

		if !isOnRamp {
			rampUpdate := lock_release_token_pool.TokenPoolRampUpdate{
				Ramp:    client.Source.OnRamp.Address(),
				Allowed: true,
			}

			tx, err := tokenConfig.Pool.ApplyRampUpdates(client.Source.Owner, []lock_release_token_pool.TokenPoolRampUpdate{rampUpdate}, []lock_release_token_pool.TokenPoolRampUpdate{})
			if err != nil {
				return errors.Wrapf(err, "failed to apply ramp update for token %s", tokenName)
			}
			err = shared.WaitForMined(client.Source.Client.logger, client.Source.Client.Client, tx.Hash(), true)
			if err != nil {
				return errors.Wrapf(err, "failed to apply ramp update for token %s", tokenName)
			}
			client.Source.logger.Infof("Setting onRamp for token %s", tokenName)
		} else {
			client.Source.logger.Infof("OnRamp already set for token %s", tokenName)
		}
	}
	return nil
}

func (client *CCIPClient) syncOffRampOnPools() error {
	for tokenName, tokenConfig := range client.Dest.SupportedTokens {
		// Only add tokens that are supported on both chains
		if destConfig, ok := client.Source.SupportedTokens[tokenName]; !ok || destConfig.TokenPoolType == rhea.FeeTokenOnly || tokenConfig.TokenPoolType == rhea.FeeTokenOnly {
			continue
		}
		isOffRamp, err := tokenConfig.Pool.IsOffRamp(&bind.CallOpts{}, client.Dest.OffRamp.Address())
		if err != nil {
			return errors.Wrapf(err, "failed loading offRamp data for token %s", tokenName)
		}

		if !isOffRamp {
			rampUpdate := lock_release_token_pool.TokenPoolRampUpdate{
				Ramp:    client.Dest.OffRamp.Address(),
				Allowed: true,
			}

			tx, err := tokenConfig.Pool.ApplyRampUpdates(client.Dest.Owner, []lock_release_token_pool.TokenPoolRampUpdate{}, []lock_release_token_pool.TokenPoolRampUpdate{rampUpdate})
			if err != nil {
				return errors.Wrapf(err, "failed to apply ramp update for token %s", tokenName)
			}
			err = shared.WaitForMined(client.Dest.Client.logger, client.Dest.Client.Client, tx.Hash(), true)
			if err != nil {
				return errors.Wrapf(err, "failed to apply ramp update for token %s", tokenName)
			}
			client.Dest.logger.Infof("Setting offRamp for token %s", tokenName)
		} else {
			client.Dest.logger.Infof("OnRamp already set for token %s", tokenName)
		}
	}
	return nil
}

func (client *CCIPClient) SyncTokenPools() {
	// onRamp maps source tokens to source pools
	client.syncPoolsOnOnRamp()
	syncPrices(&client.Source.Client, client.Dest.SupportedTokens)
	err := client.syncOnRampOnPools()
	require.NoError(client.Source.t, err)

	// offRamp maps *source* tokens to *dest* pools
	client.syncPoolsOffOnRamp()
	syncPrices(&client.Dest.Client, client.Source.SupportedTokens)
	err = client.syncOffRampOnPools()
	require.NoError(client.Dest.t, err)
}

func (client *CCIPClient) ccipSendBasicTx(t *testing.T) {
	msg := client.getBasicTx(t, client.Source.LinkTokenAddress, false)

	/////////////////////////////////
	// ADD TOKENS AND/OR DATA HERE //
	/////////////////////////////////

	DATA := []byte("")
	TOKENS := []rhea.Token{rhea.LINK}
	AMOUNTS := []*big.Int{big.NewInt(100)}

	/////////////////////////////////
	// END TOKENS AND/OR DATA HERE //
	/////////////////////////////////

	if len(TOKENS) != len(AMOUNTS) {
		t.Error("Tokens and amounts need to be the same length")
		t.FailNow()
	}

	addToFeeApprove := big.NewInt(0)

	for i, token := range TOKENS {
		msg.TokenAmounts = append(msg.TokenAmounts, router.ClientEVMTokenAmount{
			Token:  client.Source.SupportedTokens[token].Token,
			Amount: AMOUNTS[i],
		})

		if token == rhea.LINK {
			addToFeeApprove = AMOUNTS[i]
			continue
		}

		client.Source.logger.Infof("Approving %d %s", AMOUNTS[i], token)

		ERC20, err := burn_mint_erc677.NewBurnMintERC677(client.Source.SupportedTokens[token].Token, client.Source.Client.Client)
		require.NoError(t, err)

		tx, err := ERC20.Approve(client.Source.Owner, client.Source.Router.Address(), AMOUNTS[i])
		require.NoError(t, err)
		err = shared.WaitForMined(client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
		shared.RequireNoError(t, err)
	}

	msg.Data = DATA

	fee, err := client.Source.Router.GetFee(&bind.CallOpts{}, client.Dest.ChainSelector, msg)
	shared.RequireNoError(t, err)

	// If link was sent, add it to the fee for the approval
	fee = new(big.Int).Add(fee, addToFeeApprove)

	client.Source.ApproveLinkFrom(t, client.Source.Owner, client.Source.Router.Address(), fee)

	sourceBlockNumber := GetCurrentBlockNumber(client.Source.Client.Client)
	DestBlockNum := GetCurrentBlockNumber(client.Dest.Client.Client)

	tx, err := client.Source.Router.CcipSend(client.Source.Owner, client.Dest.ChainSelector, msg)
	shared.RequireNoError(t, err)
	client.Source.logger.Warnf("Message sent for max %d gas %s", tx.Gas(), helpers.ExplorerLink(int64(client.Source.ChainId), tx.Hash()))

	sendRequested := WaitForCrossChainSendRequest(client.Source, sourceBlockNumber, tx.Hash())
	require.NoError(t, client.WaitForCommit(DestBlockNum), "waiting for commit")
	require.NoError(t, client.WaitForExecution(DestBlockNum, sendRequested.Message.SequenceNumber), "waiting for execution")
}

func (client *CCIPClient) TestGasVariousTxs(t *testing.T) {
	client.sendLinkTx(t, client.Source.Owner, false)
	// we wait in between txs to make sure they're not batched
	time.Sleep(5 * time.Second)
	client.sendWrappedNativeTx(t, client.Source.Owner, false)
	time.Sleep(5 * time.Second)
	client.sendNativeTx(t, client.Source.Owner, false)
	time.Sleep(5 * time.Second)
	client.sendLinkTx(t, client.Source.Owner, true)
	time.Sleep(5 * time.Second)
	client.sendWrappedNativeTx(t, client.Source.Owner, true)
	time.Sleep(5 * time.Second)
	client.sendNativeTx(t, client.Source.Owner, true)
}

func (client *CCIPClient) getBasicTx(t *testing.T, feeToken common.Address, includesToken bool) router.ClientEVM2AnyMessage {
	msg := router.ClientEVM2AnyMessage{
		Receiver:     testhelpers.MustEncodeAddress(t, client.Dest.Owner.From),
		Data:         []byte{},
		TokenAmounts: []router.ClientEVMTokenAmount{},
		FeeToken:     feeToken,
		ExtraArgs:    []byte{},
	}
	if includesToken {
		msg.TokenAmounts = append(msg.TokenAmounts, router.ClientEVMTokenAmount{
			Token:  client.Source.LinkTokenAddress,
			Amount: big.NewInt(100),
		})
	}
	return msg
}

func (client *CCIPClient) sendLinkTx(t *testing.T, from *bind.TransactOpts, token bool) *evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested {
	msg := client.getBasicTx(t, client.Source.LinkTokenAddress, token)

	fee, err := client.Source.Router.GetFee(&bind.CallOpts{}, client.Dest.ChainSelector, msg)
	shared.RequireNoError(t, err)
	if token {
		fee.Add(fee, msg.TokenAmounts[0].Amount)
	}

	client.Source.ApproveLinkFrom(t, from, client.Source.Router.Address(), fee)

	sourceBlockNumber := GetCurrentBlockNumber(client.Source.Client.Client)
	tx, err := client.Source.Router.CcipSend(from, client.Dest.ChainSelector, msg)
	shared.RequireNoError(t, err)
	client.Source.logger.Warnf("Message sent for max %d gas %s", tx.Gas(), helpers.ExplorerLink(int64(client.Source.ChainId), tx.Hash()))

	return WaitForCrossChainSendRequest(client.Source, sourceBlockNumber, tx.Hash())
}

func (client *CCIPClient) sendWrappedNativeTx(t *testing.T, from *bind.TransactOpts, token bool) *evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested {
	msg := client.getBasicTx(t, client.Source.WrappedNative.Address(), token)

	fee, err := client.Source.Router.GetFee(&bind.CallOpts{}, client.Dest.ChainSelector, msg)
	shared.RequireNoError(t, err)

	tx, err := client.Source.WrappedNative.Approve(from, client.Source.Router.Address(), fee)
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Source.logger, client.Source.Client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
	if token {
		client.Source.ApproveLinkFrom(t, from, client.Source.Router.Address(), msg.TokenAmounts[0].Amount)
	}

	sourceBlockNumber := GetCurrentBlockNumber(client.Source.Client.Client)
	tx, err = client.Source.Router.CcipSend(from, client.Dest.ChainSelector, msg)
	shared.RequireNoError(t, err)
	client.Source.logger.Warnf("Message sent for max %d gas %s", tx.Gas(), helpers.ExplorerLink(int64(client.Source.ChainId), tx.Hash()))
	return WaitForCrossChainSendRequest(client.Source, sourceBlockNumber, tx.Hash())
}

func (client *CCIPClient) sendNativeTx(t *testing.T, from *bind.TransactOpts, token bool) *evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested {
	msg := client.getBasicTx(t, common.Address{}, token)

	fee, err := client.Source.Router.GetFee(&bind.CallOpts{}, client.Dest.ChainSelector, msg)
	shared.RequireNoError(t, err)

	if token {
		client.Source.ApproveLinkFrom(t, from, client.Source.Router.Address(), msg.TokenAmounts[0].Amount)
	}
	from.Value = fee

	sourceBlockNumber := GetCurrentBlockNumber(client.Source.Client.Client)
	tx, err := client.Source.Router.CcipSend(from, client.Dest.ChainSelector, msg)
	shared.RequireNoError(t, err)
	from.Value = big.NewInt(0)
	client.Source.logger.Warnf("Message sent for max %d gas %s", tx.Gas(), helpers.ExplorerLink(int64(client.Source.ChainId), tx.Hash()))
	return WaitForCrossChainSendRequest(client.Source, sourceBlockNumber, tx.Hash())
}

func FundPingPong(t *testing.T, chain rhea.EvmDeploymentConfig, minimumBalance *big.Int) {
	linkToken, err := burn_mint_erc677.NewBurnMintERC677(chain.ChainConfig.SupportedTokens[rhea.LINK].Token, chain.Client)
	require.NoError(t, err)

	balance, err := linkToken.BalanceOf(&bind.CallOpts{}, chain.LaneConfig.PingPongDapp)
	require.NoError(t, err)

	if balance.Cmp(minimumBalance) == -1 {
		t.Logf("❌ %s balance of %s link, which is lower than the set minimum. Funding...", ccip.ChainName(int64(chain.ChainConfig.EvmChainId)), dione.EthBalanceToString(balance))
		_, err := linkToken.Transfer(chain.Owner, chain.LaneConfig.PingPongDapp, minimumBalance)
		require.NoError(t, err)
	} else {
		t.Logf("✅ %s balance of %s link ", ccip.ChainName(int64(chain.ChainConfig.EvmChainId)), dione.EthBalanceToString(balance))
	}
}
