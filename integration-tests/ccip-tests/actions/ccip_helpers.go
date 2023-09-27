package actions

import (
	"context"
	_ "embed"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	chainselectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink-env/environment"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	ctfClient "github.com/smartcontractkit/chainlink-testing-framework/client"

	"github.com/smartcontractkit/chainlink/integration-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/ccip-tests/contracts"
	"github.com/smartcontractkit/chainlink/integration-tests/ccip-tests/contracts/laneconfig"
	"github.com/smartcontractkit/chainlink/integration-tests/ccip-tests/testreporters"
	"github.com/smartcontractkit/chainlink/integration-tests/client"
	"github.com/smartcontractkit/chainlink/integration-tests/docker/test_env"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/arm_contract"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/router"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	ccipConfig "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/config"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/testhelpers"
	integrationtesthelpers "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/testhelpers/integration"
	"github.com/smartcontractkit/chainlink/v2/core/store/models"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
	bigmath "github.com/smartcontractkit/chainlink/v2/core/utils/big_math"
)

const (
	ChaosGroupExecution           = "ExecutionNodesAll"      // all execution nodes
	ChaosGroupCommit              = "CommitNodesAll"         // all commit nodes
	ChaosGroupCommitFaultyPlus    = "CommitMajority"         // >f number of nodes
	ChaosGroupCommitFaulty        = "CommitMinority"         //  f number of nodes
	ChaosGroupExecutionFaultyPlus = "ExecutionNodesMajority" // > f number of nodes
	ChaosGroupExecutionFaulty     = "ExecutionNodesMinority" //  f number of nodes
	ChaosGroupCCIPGeth            = "CCIPGeth"               // both source and destination simulated geth networks
	ChaosGroupNetworkACCIPGeth    = "CCIPNetworkAGeth"
	ChaosGroupNetworkBCCIPGeth    = "CCIPNetworkBGeth"
	RootSnoozeTimeSimulated       = 1 * time.Minute
	InflightExpirySimulated       = 1 * time.Minute
	// we keep the finality timeout high as it's out of our control
	FinalityTimeout        = 1 * time.Hour
	TokenTransfer   string = "WithToken"

	DataOnlyTransfer string = "WithoutToken"
)

type CCIPTOMLEnv struct {
	Networks []blockchain.EVMNetwork
}

var (
	NetworkName = func(name string) string {
		return strings.ReplaceAll(strings.ToLower(name), " ", "-")
	}

	GethLabel = func(name string) string {
		return fmt.Sprintf("%s-ethereum-geth", name)
	}
	// ApprovedAmountToRouter is the default amount which gets approved for router so that it can transfer token and use the fee token for fee payment
	ApprovedAmountToRouter           = new(big.Int).Mul(big.NewInt(1e18), big.NewInt(200))
	ApprovedFeeAmountToRouter        = new(big.Int).Mul(big.NewInt(int64(GasFeeMultiplier)), big.NewInt(1e5))
	GasFeeMultiplier          uint64 = 12e17
	LinkToUSD                        = big.NewInt(6e18)
	WrappedNativeToUSD               = new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1.7e3))
)

type CCIPCommon struct {
	ChainClient        blockchain.EVMClient
	Deployer           *contracts.CCIPContractsDeployer
	FeeToken           *contracts.LinkToken
	BridgeTokens       []*contracts.ERC20Token // as of now considering the bridge token is same as link token
	TokenPrices        []*big.Int
	BridgeTokenPools   []*contracts.LockReleaseTokenPool
	RateLimiterConfig  contracts.RateLimiterConfig
	ARMContract        *common.Address
	ARM                *contracts.ARM // populate only if the ARM contracts is not a mock and can be used to verify various ARM events
	Router             *contracts.Router
	PriceRegistry      *contracts.PriceRegistry
	WrappedNative      common.Address
	ExistingDeployment bool
	poolFunds          *big.Int
	gasUpdateWatcherMu *sync.Mutex
	gasUpdateWatcher   map[uint64]*big.Int // key - destchain id; value - timestamp of update
	priceUpdateSubs    []event.Subscription
	connectionIssues   *atomic.Bool
	connectionRestored *atomic.Bool
}

func (ccipModule *CCIPCommon) ConnectionRestored() {
	for {
		select {
		case <-ccipModule.ChainClient.ConnectionRestored():
			ccipModule.connectionRestored.Store(true)
			ccipModule.connectionIssues.Store(false)
		case <-ccipModule.ChainClient.ConnectionIssue():
			ccipModule.connectionIssues.Store(true)
			ccipModule.connectionRestored.Store(false)
		}
	}
}

func (ccipModule *CCIPCommon) StopWatchingPriceUpdates() {
	for _, sub := range ccipModule.priceUpdateSubs {
		sub.Unsubscribe()
	}
}

func (ccipModule *CCIPCommon) Copy(logger zerolog.Logger, chainClient blockchain.EVMClient) (*CCIPCommon, error) {
	newCD, err := contracts.NewCCIPContractsDeployer(logger, chainClient)
	if err != nil {
		return nil, err
	}
	var arm *contracts.ARM
	if ccipModule.ARM != nil {
		arm, err = newCD.NewARMContract(*ccipModule.ARMContract)
		if err != nil {
			return nil, err
		}
	}
	var pools []*contracts.LockReleaseTokenPool
	for i := range ccipModule.BridgeTokenPools {
		pool, err := newCD.NewLockReleaseTokenPoolContract(common.HexToAddress(ccipModule.BridgeTokenPools[i].Address()))
		if err != nil {
			return nil, err
		}
		pools = append(pools, pool)
	}
	var tokens []*contracts.ERC20Token
	for i := range ccipModule.BridgeTokens {
		token, err := newCD.NewERC20TokenContract(common.HexToAddress(ccipModule.BridgeTokens[i].Address()))
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	newCommon := &CCIPCommon{
		ChainClient:        chainClient,
		Deployer:           newCD,
		BridgeTokens:       tokens,
		TokenPrices:        ccipModule.TokenPrices,
		BridgeTokenPools:   pools,
		RateLimiterConfig:  ccipModule.RateLimiterConfig,
		ARMContract:        ccipModule.ARMContract,
		ARM:                arm,
		WrappedNative:      ccipModule.WrappedNative,
		ExistingDeployment: ccipModule.ExistingDeployment,
		poolFunds:          ccipModule.poolFunds,
		gasUpdateWatcherMu: &sync.Mutex{},
		gasUpdateWatcher:   make(map[uint64]*big.Int),
	}
	newCommon.FeeToken, err = newCommon.Deployer.NewLinkTokenContract(common.HexToAddress(ccipModule.FeeToken.Address()))
	if err != nil {
		return nil, err
	}
	newCommon.PriceRegistry, err = newCommon.Deployer.NewPriceRegistry(common.HexToAddress(ccipModule.PriceRegistry.Address()))
	if err != nil {
		return nil, err
	}
	newCommon.Router, err = newCommon.Deployer.NewRouter(common.HexToAddress(ccipModule.Router.Address()))
	if err != nil {
		return nil, err
	}
	return newCommon, nil
}

func (ccipModule *CCIPCommon) LoadContractAddresses(conf *laneconfig.LaneConfig) {
	if conf != nil {
		if common.IsHexAddress(conf.FeeToken) {
			ccipModule.FeeToken = &contracts.LinkToken{
				EthAddress: common.HexToAddress(conf.FeeToken),
			}
		}
		if conf.IsNativeFeeToken {
			ccipModule.FeeToken = &contracts.LinkToken{
				EthAddress: common.HexToAddress("0x0"),
			}
		}

		if common.IsHexAddress(conf.Router) {
			ccipModule.Router = &contracts.Router{
				EthAddress: common.HexToAddress(conf.Router),
			}
		}
		if common.IsHexAddress(conf.ARM) {
			addr := common.HexToAddress(conf.ARM)
			ccipModule.ARMContract = &addr
			if !conf.IsMockARM {
				ccipModule.ARM = &contracts.ARM{
					EthAddress: addr,
				}
			}
		}
		if common.IsHexAddress(conf.PriceRegistry) {
			ccipModule.PriceRegistry = &contracts.PriceRegistry{
				EthAddress: common.HexToAddress(conf.PriceRegistry),
			}
		}
		if common.IsHexAddress(conf.WrappedNative) {
			ccipModule.WrappedNative = common.HexToAddress(conf.WrappedNative)
		}
		if len(conf.BridgeTokens) > 0 {
			var tokens []*contracts.ERC20Token
			for _, token := range conf.BridgeTokens {
				if common.IsHexAddress(token) {
					tokens = append(tokens, &contracts.ERC20Token{
						ContractAddress: common.HexToAddress(token),
					})
				}
			}
			ccipModule.BridgeTokens = tokens
		}
		if len(conf.BridgeTokenPools) > 0 {
			var pools []*contracts.LockReleaseTokenPool
			for _, pool := range conf.BridgeTokenPools {
				if common.IsHexAddress(pool) {
					pools = append(pools, &contracts.LockReleaseTokenPool{
						EthAddress: common.HexToAddress(pool),
					})
				}
			}
			ccipModule.BridgeTokenPools = pools
		}
	}
}

// ApproveTokens approve tokens for router - usually a massive amount of tokens enough to cover all the ccip transfers
// to be triggered by the test
func (ccipModule *CCIPCommon) ApproveTokens() error {
	isApproved := false
	for _, token := range ccipModule.BridgeTokens {
		err := token.Approve(ccipModule.Router.Address(), ApprovedAmountToRouter)
		if err != nil {
			return errors.WithStack(err)
		}
		if token.ContractAddress == ccipModule.FeeToken.EthAddress {
			isApproved = true
		}
	}
	if ccipModule.FeeToken.EthAddress != common.HexToAddress("0x0") {
		if !isApproved {
			err := ccipModule.FeeToken.Approve(ccipModule.Router.Address(), ApprovedFeeAmountToRouter)
			if err != nil {
				return errors.WithStack(err)
			}
		} else {
			err := ccipModule.FeeToken.Approve(ccipModule.Router.Address(), new(big.Int).Add(ApprovedAmountToRouter, ApprovedFeeAmountToRouter))
			if err != nil {
				return errors.WithStack(err)
			}
		}
	}

	return nil
}

func (ccipModule *CCIPCommon) CleanUp() error {
	if !ccipModule.ExistingDeployment {
		for i, pool := range ccipModule.BridgeTokenPools {
			bal, err := ccipModule.BridgeTokens[i].BalanceOf(context.Background(), pool.Address())
			if err != nil {
				return fmt.Errorf("error in getting pool balance %+v", err)
			}
			if bal.Cmp(big.NewInt(0)) == 0 {
				continue
			}
			err = pool.RemoveLiquidity(bal)
			if err != nil {
				return fmt.Errorf("error in removing liquidity %+v", err)
			}
		}
		err := ccipModule.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("error in waiting for events %+v", err)
		}
	}
	return nil
}

func (ccipModule *CCIPCommon) WaitForPriceUpdates(
	lggr zerolog.Logger,
	timeout time.Duration,
	destChainId uint64,
) error {
	destChainSelector, err := chainselectors.SelectorFromChainId(destChainId)
	if err != nil {
		return err
	}
	// check if price is already updated
	price, err := ccipModule.PriceRegistry.Instance.GetDestinationChainGasPrice(nil, destChainSelector)
	if err != nil {
		return err
	}
	if price.Timestamp > 0 && price.Value.Cmp(big.NewInt(0)) > 0 {
		lggr.Info().
			Str("Price Registry", ccipModule.PriceRegistry.Address()).
			Uint64("dest chain", destChainId).
			Str("source chain", ccipModule.ChainClient.GetNetworkName()).
			Msg("Price already updated")
		return nil
	}
	// if not, wait for price update
	lggr.Info().Msgf("Waiting for UsdPerUnitGas for dest chain %d Price Registry %s", destChainId, ccipModule.PriceRegistry.Address())
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	for {
		select {
		case <-ticker.C:
			ccipModule.gasUpdateWatcherMu.Lock()
			timestampOfUpdate, ok := ccipModule.gasUpdateWatcher[destChainId]
			ccipModule.gasUpdateWatcherMu.Unlock()
			if ok && timestampOfUpdate.Cmp(big.NewInt(0)) == 1 {
				lggr.Info().
					Str("Price Registry", ccipModule.PriceRegistry.Address()).
					Uint64("dest chain", destChainId).
					Str("source chain", ccipModule.ChainClient.GetNetworkName()).
					Msg("Price updated")
				return nil
			}
		case <-ctx.Done():
			return fmt.Errorf("UsdPerUnitGasUpdated is not found for chain %d", destChainId)
		}
	}
}

func (ccipModule *CCIPCommon) WatchForPriceUpdates() error {
	gasUpdateEvent := make(chan *price_registry.PriceRegistryUsdPerUnitGasUpdated)
	sub, err := ccipModule.PriceRegistry.Instance.WatchUsdPerUnitGasUpdated(nil, gasUpdateEvent, nil)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case e := <-gasUpdateEvent:
				destChain, err := chainselectors.ChainIdFromSelector(e.DestChain)
				if err != nil {
					continue
				}
				ccipModule.gasUpdateWatcherMu.Lock()
				ccipModule.gasUpdateWatcher[destChain] = e.Timestamp
				ccipModule.gasUpdateWatcherMu.Unlock()
				log.Info().
					Str("source_chain", ccipModule.ChainClient.GetNetworkName()).
					Uint64("dest_chain", destChain).
					Str("price_registry", ccipModule.PriceRegistry.Address()).
					Msgf("UsdPerUnitGasUpdated event received for dest chain %d source chain %s",
						destChain, ccipModule.ChainClient.GetNetworkName())
			case <-sub.Err():
				return
			}
		}
	}()
	ccipModule.priceUpdateSubs = append(ccipModule.priceUpdateSubs, sub)

	return nil
}

// DeployContracts deploys the contracts which are necessary in both source and dest chain
// This reuses common contracts for bidirectional lanes
func (ccipModule *CCIPCommon) DeployContracts(noOfTokens int,
	tokenDeployerFns []blockchain.ContractDeployer,
	conf *laneconfig.LaneConfig) error {
	var err error
	cd := ccipModule.Deployer

	ccipModule.LoadContractAddresses(conf)
	if ccipModule.ARM != nil {
		arm, err := cd.NewARMContract(ccipModule.ARM.EthAddress)
		if err != nil {
			return fmt.Errorf("getting new ARM contract shouldn't fail %+v", err)
		}
		ccipModule.ARM = arm
	} else {
		// deploy a mock ARM contract
		if ccipModule.ARMContract == nil {
			ccipModule.ARMContract, err = cd.DeployMockARMContract()
			if err != nil {
				return fmt.Errorf("deploying mock ARM contract shouldn't fail %+v", err)
			}
			err = ccipModule.ChainClient.WaitForEvents()
			if err != nil {
				return fmt.Errorf("error in waiting for mock ARM deployment %+v", err)
			}
		}
	}
	if ccipModule.FeeToken == nil {
		// deploy link token
		token, err := cd.DeployLinkTokenContract()
		if err != nil {
			return fmt.Errorf("deploying fee token contract shouldn't fail %+v", err)
		}

		ccipModule.FeeToken = token
		err = ccipModule.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("error in waiting for feetoken deployment %+v", err)
		}
	} else {
		token, err := cd.NewLinkTokenContract(common.HexToAddress(ccipModule.FeeToken.Address()))
		if err != nil {
			return fmt.Errorf("getting fee token contract shouldn't fail %+v", err)
		}
		ccipModule.FeeToken = token
	}

	if len(ccipModule.BridgeTokens) == 0 {
		// deploy bridge token.
		for i := len(ccipModule.BridgeTokens); i < noOfTokens; i++ {
			var token *contracts.ERC20Token
			var err error
			if len(tokenDeployerFns) != noOfTokens {
				// we deploy link token and cast it to ERC20Token
				linkToken, err := cd.DeployLinkTokenContract()
				if err != nil {
					return fmt.Errorf("deploying bridge token contract shouldn't fail %+v", err)
				}
				token, err = cd.NewERC20TokenContract(common.HexToAddress(linkToken.Address()))
			} else {
				token, err = cd.DeployERC20TokenContract(tokenDeployerFns[i])
			}

			if err != nil {
				return fmt.Errorf("deploying bridge token contract shouldn't fail %+v", err)
			}

			ccipModule.BridgeTokens = append(ccipModule.BridgeTokens, token)
		}
		err = ccipModule.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("error in waiting for bridge token deployment %+v", err)
		}
	} else {
		var tokens []*contracts.ERC20Token
		for _, token := range ccipModule.BridgeTokens {
			newToken, err := cd.NewERC20TokenContract(common.HexToAddress(token.Address()))
			if err != nil {
				return fmt.Errorf("getting new bridge token contract shouldn't fail %+v", err)
			}
			tokens = append(tokens, newToken)
		}
		ccipModule.BridgeTokens = tokens
	}
	if len(ccipModule.BridgeTokenPools) == 0 {
		// deploy native token pool
		for i := len(ccipModule.BridgeTokenPools); i < noOfTokens; i++ {
			token := ccipModule.BridgeTokens[i]
			btp, err := cd.DeployLockReleaseTokenPoolContract(token.Address(), *ccipModule.ARMContract)
			if err != nil {
				return fmt.Errorf("deploying bridge Token pool shouldn't fail %+v", err)
			}
			ccipModule.BridgeTokenPools = append(ccipModule.BridgeTokenPools, btp)
			err = btp.AddLiquidity(token.Approve, token.Address(), ccipModule.poolFunds)
			if err != nil {
				return fmt.Errorf("adding liquidity token to dest pool shouldn't fail %+v", err)
			}
		}
	} else {
		var pools []*contracts.LockReleaseTokenPool
		for _, pool := range ccipModule.BridgeTokenPools {
			newPool, err := cd.NewLockReleaseTokenPoolContract(pool.EthAddress)
			if err != nil {
				return fmt.Errorf("getting new bridge token pool contract shouldn't fail %+v", err)
			}
			pools = append(pools, newPool)
		}
		ccipModule.BridgeTokenPools = pools
	}
	// Set price of the bridge tokens to 1
	ccipModule.TokenPrices = []*big.Int{}
	for range ccipModule.BridgeTokens {
		ccipModule.TokenPrices = append(ccipModule.TokenPrices, big.NewInt(6e18))
	}

	if ccipModule.WrappedNative == common.HexToAddress("0x0") {
		weth9addr, err := cd.DeployWrappedNative()
		if err != nil {
			return fmt.Errorf("deploying wrapped native shouldn't fail %+v", err)
		}

		err = ccipModule.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("waiting for deploying wrapped native shouldn't fail %+v", err)
		}
		ccipModule.WrappedNative = *weth9addr
	}

	if ccipModule.Router == nil {
		ccipModule.Router, err = cd.DeployRouter(ccipModule.WrappedNative, *ccipModule.ARMContract)
		if err != nil {
			return fmt.Errorf("deploying router shouldn't fail %+v", err)
		}
		err = ccipModule.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("error in waiting for router deployment %+v", err)
		}
	} else {
		r, err := cd.NewRouter(ccipModule.Router.EthAddress)
		if err != nil {
			return fmt.Errorf("getting new router contract shouldn't fail %+v", err)
		}
		ccipModule.Router = r
	}
	if ccipModule.PriceRegistry == nil {
		// we will update the price updates later based on source and dest PriceUpdates
		ccipModule.PriceRegistry, err = cd.DeployPriceRegistry([]common.Address{
			common.HexToAddress(ccipModule.FeeToken.Address()),
			common.HexToAddress(ccipModule.WrappedNative.Hex()),
		})
		if err != nil {
			return fmt.Errorf("deploying PriceRegistry shouldn't fail %+v", err)
		}
		err = ccipModule.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("error in waiting for PriceRegistry deployment %+v", err)
		}
	} else {
		ccipModule.PriceRegistry, err = cd.NewPriceRegistry(ccipModule.PriceRegistry.EthAddress)
		if err != nil {
			return fmt.Errorf("getting new PriceRegistry contract shouldn't fail %+v", err)
		}
	}

	log.Info().Msg("finished deploying common contracts")
	return nil
}

func DefaultCCIPModule(logger zerolog.Logger, chainClient blockchain.EVMClient, existingDeployment bool) (*CCIPCommon, error) {
	cd, err := contracts.NewCCIPContractsDeployer(logger, chainClient)
	if err != nil {
		return nil, err
	}
	return &CCIPCommon{
		ChainClient: chainClient,
		Deployer:    cd,
		RateLimiterConfig: contracts.RateLimiterConfig{
			Rate:     contracts.HundredCoins,
			Capacity: contracts.HundredCoins,
		},
		ExistingDeployment: existingDeployment,
		poolFunds:          testhelpers.Link(1000),
		gasUpdateWatcherMu: &sync.Mutex{},
		gasUpdateWatcher:   make(map[uint64]*big.Int),
	}, nil
}

type SourceCCIPModule struct {
	Common                     *CCIPCommon
	Sender                     common.Address
	TransferAmount             []*big.Int
	DestinationChainId         uint64
	DestNetworkName            string
	OnRamp                     *contracts.OnRamp
	SrcStartBlock              uint64
	CCIPSendRequestedWatcherMu *sync.Mutex
	CCIPSendRequestedWatcher   map[string]*evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested
	NewFinalizedBlockNum       atomic.Uint64
	NewFinalizedBlockTimestamp atomic.Time
}

func (sourceCCIP *SourceCCIPModule) PayCCIPFeeToOwnerAddress() error {
	isNativeFee := sourceCCIP.Common.FeeToken.EthAddress == common.HexToAddress("0x0")
	if isNativeFee {
		err := sourceCCIP.OnRamp.WithdrawNonLinkFees(sourceCCIP.Common.WrappedNative)
		if err != nil {
			return err
		}
	} else {
		err := sourceCCIP.OnRamp.SetNops()
		if err != nil {
			return err
		}
		err = sourceCCIP.OnRamp.PayNops()
		if err != nil {
			return err
		}
	}
	return nil
}

func (sourceCCIP *SourceCCIPModule) LoadContracts(conf *laneconfig.LaneConfig) {
	if conf != nil {
		cfg, ok := conf.SrcContracts[sourceCCIP.DestNetworkName]
		if ok {
			if common.IsHexAddress(cfg.OnRamp) {
				sourceCCIP.OnRamp = &contracts.OnRamp{
					EthAddress: common.HexToAddress(cfg.OnRamp),
				}
			}
			if cfg.DepolyedAt > 0 {
				sourceCCIP.SrcStartBlock = cfg.DepolyedAt
			}
		}
	}
}

// DeployContracts deploys all CCIP contracts specific to the source chain
func (sourceCCIP *SourceCCIPModule) DeployContracts(lane *laneconfig.LaneConfig) error {
	var err error
	contractDeployer := sourceCCIP.Common.Deployer
	log.Info().Msg("Deploying source chain specific contracts")

	err = sourceCCIP.Common.ApproveTokens()
	if err != nil {
		return err
	}

	sourceCCIP.LoadContracts(lane)
	// update transfer amount array length to be equal to the number of tokens
	// each index in TransferAmount array corresponds to the amount to be transferred for the token at the same index in BridgeTokens array
	if len(sourceCCIP.TransferAmount) != len(sourceCCIP.Common.BridgeTokens) && len(sourceCCIP.TransferAmount) > 0 {
		sourceCCIP.TransferAmount = sourceCCIP.TransferAmount[:len(sourceCCIP.Common.BridgeTokens)]
	}
	sourceChainSelector, err := chainselectors.SelectorFromChainId(sourceCCIP.Common.ChainClient.GetChainID().Uint64())
	if err != nil {
		return errors.WithStack(err)
	}
	destChainSelector, err := chainselectors.SelectorFromChainId(sourceCCIP.DestinationChainId)
	if err != nil {
		return errors.WithStack(err)
	}

	if sourceCCIP.OnRamp == nil {
		var tokensAndPools []evm_2_evm_onramp.InternalPoolUpdate
		var tokenTransferFeeConfig []evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfigArgs
		for i, token := range sourceCCIP.Common.BridgeTokens {
			tokensAndPools = append(tokensAndPools, evm_2_evm_onramp.InternalPoolUpdate{
				Token: token.ContractAddress,
				Pool:  sourceCCIP.Common.BridgeTokenPools[i].EthAddress,
			})
			tokenTransferFeeConfig = append(tokenTransferFeeConfig, evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfigArgs{
				Token:             token.ContractAddress,
				Ratio:             5_0, // 5 bps
				DestGasOverhead:   34_000,
				DestBytesOverhead: 0,
			})
		}

		sourceCCIP.SrcStartBlock, err = sourceCCIP.Common.ChainClient.LatestBlockNumber(context.Background())
		if err != nil {
			return fmt.Errorf("getting latest block number shouldn't fail %+v", err)
		}

		sourceCCIP.OnRamp, err = contractDeployer.DeployOnRamp(
			sourceChainSelector,
			destChainSelector,
			tokensAndPools,
			*sourceCCIP.Common.ARMContract,
			sourceCCIP.Common.Router.EthAddress,
			sourceCCIP.Common.PriceRegistry.EthAddress,
			sourceCCIP.Common.RateLimiterConfig,
			[]evm_2_evm_onramp.EVM2EVMOnRampFeeTokenConfigArgs{
				{
					Token:                  common.HexToAddress(sourceCCIP.Common.FeeToken.Address()),
					NetworkFeeUSD:          1_00,
					MinTokenTransferFeeUSD: 1_00,
					MaxTokenTransferFeeUSD: 5000_00,
					GasMultiplier:          GasFeeMultiplier,
					PremiumMultiplier:      1e18,
					Enabled:                true,
				},
				{
					Token:                  sourceCCIP.Common.WrappedNative,
					NetworkFeeUSD:          1_00,
					MinTokenTransferFeeUSD: 1_00,
					MaxTokenTransferFeeUSD: 5000_00,
					GasMultiplier:          GasFeeMultiplier,
					PremiumMultiplier:      1e18,
					Enabled:                true,
				},
			},
			tokenTransferFeeConfig,
			sourceCCIP.Common.FeeToken.EthAddress,
		)

		if err != nil {
			return fmt.Errorf("onRamp deployment shouldn't fail %+v", err)
		}

		err = sourceCCIP.Common.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("waiting for onRamp deployment shouldn't fail %+v", err)
		}

		// update source Router with OnRamp address
		err = sourceCCIP.Common.Router.SetOnRamp(destChainSelector, sourceCCIP.OnRamp.EthAddress)
		if err != nil {
			return fmt.Errorf("setting onramp on the router shouldn't fail %+v", err)
		}

		err = sourceCCIP.Common.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("waiting for events shouldn't fail %+v", err)
		}

		// update native pool with onRamp address
		for _, pool := range sourceCCIP.Common.BridgeTokenPools {
			err = pool.SetOnRamp(sourceCCIP.OnRamp.EthAddress)
			if err != nil {
				return fmt.Errorf("setting OnRamp on the bridge token pool shouldn't fail %+v", err)
			}
		}

		err = sourceCCIP.Common.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("waiting for events shouldn't fail %+v", err)
		}
	} else {
		sourceCCIP.OnRamp, err = contractDeployer.NewOnRamp(sourceCCIP.OnRamp.EthAddress)
		if err != nil {
			return fmt.Errorf("getting new onramp contractshouldn't fail %+v", err)
		}
	}

	return nil
}

func (sourceCCIP *SourceCCIPModule) CollectBalanceRequirements() []testhelpers.BalanceReq {
	var balancesReq []testhelpers.BalanceReq
	for _, token := range sourceCCIP.Common.BridgeTokens {
		balancesReq = append(balancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("BridgeToken-%s-Address-%s", token.Address(), sourceCCIP.Sender.Hex()),
			Addr:   sourceCCIP.Sender,
			Getter: GetterForLinkToken(token.BalanceOf, sourceCCIP.Sender.Hex()),
		})
	}
	for i, pool := range sourceCCIP.Common.BridgeTokenPools {
		balancesReq = append(balancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("BridgeToken-%s-TokenPool-%s", sourceCCIP.Common.BridgeTokens[i].Address(), pool.Address()),
			Addr:   pool.EthAddress,
			Getter: GetterForLinkToken(sourceCCIP.Common.BridgeTokens[i].BalanceOf, pool.Address()),
		})
	}

	if sourceCCIP.Common.FeeToken.Address() != common.HexToAddress("0x0").String() {
		balancesReq = append(balancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("FeeToken-%s-Address-%s", sourceCCIP.Common.FeeToken.Address(), sourceCCIP.Sender.Hex()),
			Addr:   sourceCCIP.Sender,
			Getter: GetterForLinkToken(sourceCCIP.Common.FeeToken.BalanceOf, sourceCCIP.Sender.Hex()),
		})
		balancesReq = append(balancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("FeeToken-%s-Router-%s", sourceCCIP.Common.FeeToken.Address(), sourceCCIP.Common.Router.Address()),
			Addr:   sourceCCIP.Common.Router.EthAddress,
			Getter: GetterForLinkToken(sourceCCIP.Common.FeeToken.BalanceOf, sourceCCIP.Common.Router.Address()),
		})
		balancesReq = append(balancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("FeeToken-%s-OnRamp-%s", sourceCCIP.Common.FeeToken.Address(), sourceCCIP.OnRamp.Address()),
			Addr:   sourceCCIP.OnRamp.EthAddress,
			Getter: GetterForLinkToken(sourceCCIP.Common.FeeToken.BalanceOf, sourceCCIP.OnRamp.Address()),
		})
		balancesReq = append(balancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("FeeToken-%s-Prices-%s", sourceCCIP.Common.FeeToken.Address(), sourceCCIP.Common.PriceRegistry.Address()),
			Addr:   sourceCCIP.Common.PriceRegistry.EthAddress,
			Getter: GetterForLinkToken(sourceCCIP.Common.FeeToken.BalanceOf, sourceCCIP.Common.PriceRegistry.Address()),
		})
	}
	return balancesReq
}

func (sourceCCIP *SourceCCIPModule) UpdateBalance(
	noOfReq int64,
	totalFee *big.Int,
	balances *BalanceSheet,
) {
	if len(sourceCCIP.TransferAmount) > 0 {
		for i, token := range sourceCCIP.Common.BridgeTokens {
			name := fmt.Sprintf("BridgeToken-%s-Address-%s", token.Address(), sourceCCIP.Sender.Hex())
			balances.Update(name, BalanceItem{
				Address:  sourceCCIP.Sender,
				Getter:   GetterForLinkToken(token.BalanceOf, sourceCCIP.Sender.Hex()),
				AmtToSub: bigmath.Mul(big.NewInt(noOfReq), sourceCCIP.TransferAmount[i]),
			})
		}
		for i, pool := range sourceCCIP.Common.BridgeTokenPools {
			name := fmt.Sprintf("BridgeToken-%s-TokenPool-%s", sourceCCIP.Common.BridgeTokens[i].Address(), pool.Address())
			balances.Update(name, BalanceItem{
				Address:  pool.EthAddress,
				Getter:   GetterForLinkToken(sourceCCIP.Common.BridgeTokens[i].BalanceOf, pool.Address()),
				AmtToAdd: bigmath.Mul(big.NewInt(noOfReq), sourceCCIP.TransferAmount[i]),
			})
		}
	}
	if sourceCCIP.Common.FeeToken.Address() != common.HexToAddress("0x0").String() {
		name := fmt.Sprintf("FeeToken-%s-Address-%s", sourceCCIP.Common.FeeToken.Address(), sourceCCIP.Sender.Hex())
		balances.Update(name, BalanceItem{
			Address:  sourceCCIP.Sender,
			Getter:   GetterForLinkToken(sourceCCIP.Common.FeeToken.BalanceOf, sourceCCIP.Sender.Hex()),
			AmtToSub: totalFee,
		})
		name = fmt.Sprintf("FeeToken-%s-Prices-%s", sourceCCIP.Common.FeeToken.Address(), sourceCCIP.Common.PriceRegistry.Address())
		balances.Update(name, BalanceItem{
			Address: sourceCCIP.Common.PriceRegistry.EthAddress,
			Getter:  GetterForLinkToken(sourceCCIP.Common.FeeToken.BalanceOf, sourceCCIP.Common.PriceRegistry.Address()),
		})
		name = fmt.Sprintf("FeeToken-%s-Router-%s", sourceCCIP.Common.FeeToken.Address(), sourceCCIP.Common.Router.Address())
		balances.Update(name, BalanceItem{
			Address: sourceCCIP.Common.Router.EthAddress,
			Getter:  GetterForLinkToken(sourceCCIP.Common.FeeToken.BalanceOf, sourceCCIP.Common.Router.Address()),
		})
		name = fmt.Sprintf("FeeToken-%s-OnRamp-%s", sourceCCIP.Common.FeeToken.Address(), sourceCCIP.OnRamp.Address())
		balances.Update(name, BalanceItem{
			Address:  sourceCCIP.OnRamp.EthAddress,
			Getter:   GetterForLinkToken(sourceCCIP.Common.FeeToken.BalanceOf, sourceCCIP.OnRamp.Address()),
			AmtToAdd: totalFee,
		})
	}
}

func (sourceCCIP *SourceCCIPModule) AssertSendRequestedLogFinalized(
	lggr zerolog.Logger,
	reqNo int64,
	seqNum uint64,
	SendRequested *evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested,
	prevEventAt time.Time,
	reports *testreporters.CCIPLaneStats,
) (time.Time, uint64, error) {
	if sourceCCIP.Common.ChainClient.NetworkSimulated() {
		return prevEventAt, 0, nil
	}
	lggr.Info().Msg("Waiting for CCIPSendRequested event log to be finalized")
	finalizedBlockNum, finalizedAt, err := sourceCCIP.Common.ChainClient.WaitForFinalizedTx(SendRequested.Raw.TxHash)
	if err != nil {
		reports.UpdatePhaseStats(reqNo, seqNum, testreporters.SourceLogFinalized, time.Since(prevEventAt), testreporters.Failure)
		return time.Time{}, 0, fmt.Errorf("error waiting for CCIPSendRequested event log to be finalized - %+v", err)
	}
	reports.UpdatePhaseStats(reqNo, seqNum, testreporters.SourceLogFinalized, finalizedAt.Sub(prevEventAt), testreporters.Success,
		testreporters.TransactionStats{
			TxHash:           SendRequested.Raw.TxHash.Hex(),
			FinalizedByBlock: finalizedBlockNum.String(),
			FinalizedAt:      finalizedAt.String(),
		})
	return finalizedAt, finalizedBlockNum.Uint64(), nil
}

func (sourceCCIP *SourceCCIPModule) AssertEventCCIPSendRequested(
	lggr zerolog.Logger,
	reqNo int64,
	txHash string,
	timeout time.Duration,
	prevEventAt time.Time,
	reports *testreporters.CCIPLaneStats,
) (*evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested, time.Time, error) {
	lggr.Info().Msg("Waiting for CCIPSendRequested event")
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	for {
		select {
		case <-ticker.C:
			sourceCCIP.CCIPSendRequestedWatcherMu.Lock()
			sendRequested, ok := sourceCCIP.CCIPSendRequestedWatcher[txHash]
			sourceCCIP.CCIPSendRequestedWatcherMu.Unlock()
			if ok && sendRequested != nil {
				hdr, err := sourceCCIP.Common.ChainClient.HeaderByNumber(ctx, big.NewInt(int64(sendRequested.Raw.BlockNumber)))
				receivedAt := time.Now().UTC()
				if err == nil {
					receivedAt = hdr.Timestamp
				}
				sentMsg := sendRequested.Message
				seqNum := sentMsg.SequenceNumber
				reports.UpdatePhaseStats(reqNo, seqNum, testreporters.CCIPSendRe, receivedAt.Sub(prevEventAt), testreporters.Success)
				return sendRequested, receivedAt, nil
			}
		case <-ctx.Done():
			reports.UpdatePhaseStats(reqNo, 0, testreporters.CCIPSendRe, time.Since(prevEventAt), testreporters.Failure)
			return nil, time.Now(), fmt.Errorf("CCIPSendRequested event is not found for tx %s", txHash)
		}
	}
}

func (sourceCCIP *SourceCCIPModule) SendRequest(
	receiver common.Address,
	msgType,
	data string,
	feeToken common.Address,
) (common.Hash, time.Duration, *big.Int, error) {
	var tokenAndAmounts []router.ClientEVMTokenAmount
	if msgType == TokenTransfer {
		for i, token := range sourceCCIP.Common.BridgeTokens {
			tokenAndAmounts = append(tokenAndAmounts, router.ClientEVMTokenAmount{
				Token: common.HexToAddress(token.Address()), Amount: sourceCCIP.TransferAmount[i],
			})
		}
	}
	receiverAddr, err := utils.ABIEncode(`[{"type":"address"}]`, receiver)
	var d time.Duration
	if err != nil {
		return common.Hash{}, d, nil, fmt.Errorf("failed encoding the receiver address: %+v", err)
	}

	extraArgsV1, err := testhelpers.GetEVMExtraArgsV1(big.NewInt(100_000), false)
	if err != nil {
		return common.Hash{}, d, nil, fmt.Errorf("failed encoding the options field: %+v", err)
	}
	destChainSelector, err := chainselectors.SelectorFromChainId(sourceCCIP.DestinationChainId)
	if err != nil {
		return common.Hash{}, d, nil, fmt.Errorf("failed getting the chain selector: %+v", err)
	}
	// form the message for transfer
	msg := router.ClientEVM2AnyMessage{
		Receiver:     receiverAddr,
		Data:         []byte(data),
		TokenAmounts: tokenAndAmounts,
		FeeToken:     feeToken,
		ExtraArgs:    extraArgsV1,
	}
	log.Info().Interface("msg details", msg).Msg("ccip message to be sent")
	fee, err := sourceCCIP.Common.Router.GetFee(destChainSelector, msg)
	if err != nil {
		reason, _ := blockchain.RPCErrorFromError(err)
		if reason != "" {
			return common.Hash{}, d, nil, fmt.Errorf("failed getting the fee: %s", reason)
		}
		return common.Hash{}, d, nil, fmt.Errorf("failed getting the fee: %+v", err)
	}
	log.Info().Str("fee", fee.String()).Msg("calculated fee")

	var sendTx *types.Transaction
	timeNow := time.Now()

	// initiate the transfer
	// if the token address is 0x0 it will use Native as fee token and the fee amount should be mentioned in bind.TransactOpts's value
	if feeToken != common.HexToAddress("0x0") {
		sendTx, err = sourceCCIP.Common.Router.CCIPSend(destChainSelector, msg, nil)
		if err != nil {
			return common.Hash{}, time.Since(timeNow), nil, fmt.Errorf("failed initiating the transfer ccip-send: %+v", err)
		}
	} else {
		sendTx, err = sourceCCIP.Common.Router.CCIPSend(destChainSelector, msg, fee)
		if err != nil {
			return common.Hash{}, time.Since(timeNow), nil, fmt.Errorf("failed initiating the transfer ccip-send: %+v", err)
		}
	}

	log.Info().
		Str("Network", sourceCCIP.Common.ChainClient.GetNetworkName()).
		Str("Send token transaction", sendTx.Hash().String()).
		Str("lane", fmt.Sprintf("%s-->%s", sourceCCIP.Common.ChainClient.GetNetworkName(), sourceCCIP.DestNetworkName)).
		Msg("Sending token")
	return sendTx.Hash(), time.Since(timeNow), fee, nil
}

func DefaultSourceCCIPModule(logger zerolog.Logger, chainClient blockchain.EVMClient, destChainId uint64, destChain string, transferAmount []*big.Int, ccipCommon *CCIPCommon) (*SourceCCIPModule, error) {
	cmn, err := ccipCommon.Copy(logger, chainClient)
	if err != nil {
		return nil, err
	}
	return &SourceCCIPModule{
		Common:                     cmn,
		TransferAmount:             transferAmount,
		DestinationChainId:         destChainId,
		DestNetworkName:            destChain,
		Sender:                     common.HexToAddress(chainClient.GetDefaultWallet().Address()),
		CCIPSendRequestedWatcher:   make(map[string]*evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested),
		CCIPSendRequestedWatcherMu: &sync.Mutex{},
	}, nil
}

type DestCCIPModule struct {
	Common                  *CCIPCommon
	SourceChainId           uint64
	SourceNetworkName       string
	CommitStore             *contracts.CommitStore
	ReceiverDapp            *contracts.ReceiverDapp
	OffRamp                 *contracts.OffRamp
	WrappedNative           common.Address
	ReportAcceptedWatcherMu *sync.Mutex
	ReportAcceptedWatcher   map[uint64]*commit_store.CommitStoreReportAccepted
	ExecStateChangedMu      *sync.Mutex
	ExecStateChangedWatcher map[uint64]*evm_2_evm_offramp.EVM2EVMOffRampExecutionStateChanged
	ReportBlessedWatcherMu  *sync.Mutex
	ReportBlessedWatcher    map[[32]byte]*types.Log
	NextSeqNumToCommit      *atomic.Uint64
}

func (destCCIP *DestCCIPModule) LoadContracts(conf *laneconfig.LaneConfig) {
	if conf != nil {
		cfg, ok := conf.DestContracts[destCCIP.SourceNetworkName]
		if ok {
			if common.IsHexAddress(cfg.OffRamp) {
				destCCIP.OffRamp = &contracts.OffRamp{
					EthAddress: common.HexToAddress(cfg.OffRamp),
				}
			}
			if common.IsHexAddress(cfg.CommitStore) {
				destCCIP.CommitStore = &contracts.CommitStore{
					EthAddress: common.HexToAddress(cfg.CommitStore),
				}
			}
			if common.IsHexAddress(cfg.ReceiverDapp) {
				destCCIP.ReceiverDapp = &contracts.ReceiverDapp{
					EthAddress: common.HexToAddress(cfg.ReceiverDapp),
				}
			}
		}
	}
}

// DeployContracts deploys all CCIP contracts specific to the destination chain
func (destCCIP *DestCCIPModule) DeployContracts(
	sourceCCIP SourceCCIPModule,
	lane *laneconfig.LaneConfig,
) error {
	var err error
	contractDeployer := destCCIP.Common.Deployer
	log.Info().Msg("Deploying destination chain specific contracts")
	destCCIP.LoadContracts(lane)
	sourceChainSelector, err := chainselectors.SelectorFromChainId(destCCIP.SourceChainId)
	if err != nil {
		return errors.WithStack(err)
	}
	destChainSelector, err := chainselectors.SelectorFromChainId(destCCIP.Common.ChainClient.GetChainID().Uint64())
	if err != nil {
		return errors.WithStack(err)
	}

	if destCCIP.CommitStore == nil {
		// commitStore responsible for validating the transfer message
		destCCIP.CommitStore, err = contractDeployer.DeployCommitStore(
			sourceChainSelector, destChainSelector, sourceCCIP.OnRamp.EthAddress,
			*destCCIP.Common.ARMContract,
		)
		if err != nil {
			return fmt.Errorf("deploying commitstore shouldn't fail %+v", err)
		}
		err = destCCIP.Common.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("waiting for commitstore deployment shouldn't fail %+v", err)
		}

		// CommitStore can update
		err = destCCIP.Common.PriceRegistry.AddPriceUpdater(destCCIP.CommitStore.EthAddress)
		if err != nil {
			return fmt.Errorf("setting commitstore as fee updater shouldn't fail %+v", err)
		}
		err = destCCIP.Common.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("waiting for setting commitstore as fee updater shouldn't fail %+v", err)
		}
	} else {
		destCCIP.CommitStore, err = contractDeployer.NewCommitStore(destCCIP.CommitStore.EthAddress)
		if err != nil {
			return fmt.Errorf("getting new commitstore shouldn't fail %+v", err)
		}
	}

	var sourceTokens, destTokens, pools []common.Address

	for _, token := range sourceCCIP.Common.BridgeTokens {
		sourceTokens = append(sourceTokens, common.HexToAddress(token.Address()))
	}

	for i, token := range destCCIP.Common.BridgeTokens {
		destTokens = append(destTokens, common.HexToAddress(token.Address()))
		pool := destCCIP.Common.BridgeTokenPools[i]
		pools = append(pools, pool.EthAddress)
	}

	if destCCIP.OffRamp == nil {
		destCCIP.OffRamp, err = contractDeployer.DeployOffRamp(
			sourceChainSelector, destChainSelector,
			destCCIP.CommitStore.EthAddress, sourceCCIP.OnRamp.EthAddress,
			sourceTokens, pools, destCCIP.Common.RateLimiterConfig, *destCCIP.Common.ARMContract)
		if err != nil {
			return fmt.Errorf("deploying offramp shouldn't fail %+v", err)
		}
		err = destCCIP.Common.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("waiting for offramp deployment shouldn't fail %+v", err)
		}

		// apply offramp updates
		_, err = destCCIP.Common.Router.AddOffRamp(destCCIP.OffRamp.EthAddress, sourceChainSelector)
		if err != nil {
			return fmt.Errorf("setting offramp as fee updater shouldn't fail %+v", err)
		}
		err = destCCIP.Common.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("waiting for events on destination contract shouldn't fail %+v", err)
		}

		// update pools with offRamp id
		for _, pool := range destCCIP.Common.BridgeTokenPools {
			err = pool.SetOffRamp(destCCIP.OffRamp.EthAddress)
			if err != nil {
				return fmt.Errorf("setting offramp on the bridge token pool shouldn't fail %+v", err)
			}
		}

		err = destCCIP.Common.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("waiting for events on destination contract shouldn't fail %+v", err)
		}
	} else {
		destCCIP.OffRamp, err = contractDeployer.NewOffRamp(destCCIP.OffRamp.EthAddress)
		if err != nil {
			return fmt.Errorf("getting new offramp shouldn't fail %+v", err)
		}
	}
	if destCCIP.ReceiverDapp == nil {
		// ReceiverDapp
		destCCIP.ReceiverDapp, err = contractDeployer.DeployReceiverDapp(destCCIP.Common.Router.EthAddress)
		if err != nil {
			return fmt.Errorf("receiverDapp contract should be deployed successfully %+v", err)
		}
		err = destCCIP.Common.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("waiting for events on destination contract deployments %+v", err)
		}
	} else {
		destCCIP.ReceiverDapp, err = contractDeployer.NewReceiverDapp(destCCIP.ReceiverDapp.EthAddress)
		if err != nil {
			return fmt.Errorf("getting new receiverDapp shouldn't fail %+v", err)
		}
	}
	return nil
}

func (destCCIP *DestCCIPModule) CollectBalanceRequirements() []testhelpers.BalanceReq {
	var destBalancesReq []testhelpers.BalanceReq
	for _, token := range destCCIP.Common.BridgeTokens {
		destBalancesReq = append(destBalancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("BridgeToken-%s-Address-%s", token.Address(), destCCIP.ReceiverDapp.Address()),
			Addr:   destCCIP.ReceiverDapp.EthAddress,
			Getter: GetterForLinkToken(token.BalanceOf, destCCIP.ReceiverDapp.Address()),
		})
	}
	for i, pool := range destCCIP.Common.BridgeTokenPools {
		destBalancesReq = append(destBalancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("BridgeToken-%s-TokenPool-%s", destCCIP.Common.BridgeTokens[i].Address(), pool.Address()),
			Addr:   pool.EthAddress,
			Getter: GetterForLinkToken(destCCIP.Common.BridgeTokens[i].BalanceOf, pool.Address()),
		})
	}
	if destCCIP.Common.FeeToken.Address() != common.HexToAddress("0x0").String() {
		destBalancesReq = append(destBalancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("FeeToken-%s-Address-%s", destCCIP.Common.FeeToken.Address(), destCCIP.ReceiverDapp.Address()),
			Addr:   destCCIP.ReceiverDapp.EthAddress,
			Getter: GetterForLinkToken(destCCIP.Common.FeeToken.BalanceOf, destCCIP.ReceiverDapp.Address()),
		})
		destBalancesReq = append(destBalancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("FeeToken-%s-OffRamp-%s", destCCIP.Common.FeeToken.Address(), destCCIP.OffRamp.Address()),
			Addr:   destCCIP.OffRamp.EthAddress,
			Getter: GetterForLinkToken(destCCIP.Common.FeeToken.BalanceOf, destCCIP.OffRamp.Address()),
		})
	}
	return destBalancesReq
}

func (destCCIP *DestCCIPModule) UpdateBalance(
	transferAmount []*big.Int,
	noOfReq int64,
	balance *BalanceSheet,
) {
	if len(transferAmount) > 0 {
		for i, token := range destCCIP.Common.BridgeTokens {
			name := fmt.Sprintf("BridgeToken-%s-Address-%s", token.Address(), destCCIP.ReceiverDapp.Address())
			balance.Update(name, BalanceItem{
				Address:  destCCIP.ReceiverDapp.EthAddress,
				Getter:   GetterForLinkToken(token.BalanceOf, destCCIP.ReceiverDapp.Address()),
				AmtToAdd: bigmath.Mul(big.NewInt(noOfReq), transferAmount[i]),
			})
		}
		for i, pool := range destCCIP.Common.BridgeTokenPools {
			name := fmt.Sprintf("BridgeToken-%s-TokenPool-%s", destCCIP.Common.BridgeTokens[i].Address(), pool.Address())
			balance.Update(name, BalanceItem{
				Address:  pool.EthAddress,
				Getter:   GetterForLinkToken(destCCIP.Common.BridgeTokens[i].BalanceOf, pool.Address()),
				AmtToSub: bigmath.Mul(big.NewInt(noOfReq), transferAmount[i]),
			})
		}
	}
	if destCCIP.Common.FeeToken.Address() != common.HexToAddress("0x0").String() {
		name := fmt.Sprintf("FeeToken-%s-OffRamp-%s", destCCIP.Common.FeeToken.Address(), destCCIP.OffRamp.Address())
		balance.Update(name, BalanceItem{
			Address: destCCIP.OffRamp.EthAddress,
			Getter:  GetterForLinkToken(destCCIP.Common.FeeToken.BalanceOf, destCCIP.OffRamp.Address()),
		})

		name = fmt.Sprintf("FeeToken-%s-Address-%s", destCCIP.Common.FeeToken.Address(), destCCIP.ReceiverDapp.Address())
		balance.Update(name, BalanceItem{
			Address: destCCIP.ReceiverDapp.EthAddress,
			Getter:  GetterForLinkToken(destCCIP.Common.FeeToken.BalanceOf, destCCIP.ReceiverDapp.Address()),
		})
	}
}

func (destCCIP *DestCCIPModule) AssertEventExecutionStateChanged(
	lggr zerolog.Logger,
	reqNo int64,
	seqNum uint64,
	timeout time.Duration,
	timeNow time.Time,
	reports *testreporters.CCIPLaneStats,
) error {
	lggr.Info().Int64("seqNum", int64(seqNum)).Msg("Waiting for ExecutionStateChanged event")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			destCCIP.ExecStateChangedMu.Lock()
			e, ok := destCCIP.ExecStateChangedWatcher[seqNum]
			destCCIP.ExecStateChangedMu.Unlock()
			if ok && e != nil {
				vLogs := e.Raw
				receivedAt := time.Now().UTC()
				hdr, err := destCCIP.Common.ChainClient.HeaderByNumber(ctx, big.NewInt(int64(vLogs.BlockNumber)))
				if err == nil {
					receivedAt = hdr.Timestamp
				}
				receipt, err := destCCIP.Common.ChainClient.GetTxReceipt(vLogs.TxHash)
				if err != nil {
					lggr.Warn().Msg("Failed to get receipt for ExecStateChanged event")
				}
				if abihelpers.MessageExecutionState(e.State) == abihelpers.ExecutionStateSuccess {
					reports.UpdatePhaseStats(reqNo, seqNum, testreporters.ExecStateChanged, receivedAt.Sub(timeNow),
						testreporters.Success,
						testreporters.TransactionStats{
							TxHash:  vLogs.TxHash.Hex(),
							GasUsed: receipt.GasUsed,
						})
					return nil
				} else {
					reports.UpdatePhaseStats(reqNo, seqNum, testreporters.ExecStateChanged, time.Since(timeNow), testreporters.Failure)
					return fmt.Errorf("ExecutionStateChanged event state changed to %d with data %x for seq num %v for lane %d-->%d",
						e.State, e.ReturnData, seqNum, destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID())
				}
			}
		case <-ctx.Done():
			reports.UpdatePhaseStats(reqNo, seqNum, testreporters.ExecStateChanged, time.Since(timeNow), testreporters.Failure)
			return fmt.Errorf("ExecutionStateChanged event not found for seq num %v for lane %d-->%d",
				seqNum, destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID())
		}
	}
}

func (destCCIP *DestCCIPModule) AssertEventReportAccepted(
	lggr zerolog.Logger,
	reqNo int64,
	seqNum uint64,
	timeout time.Duration,
	prevEventAt time.Time,
	reports *testreporters.CCIPLaneStats,
) (*commit_store.CommitStoreCommitReport, time.Time, error) {
	lggr.Info().Int64("seqNum", int64(seqNum)).Msg("Waiting for ReportAccepted event")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			destCCIP.ReportAcceptedWatcherMu.Lock()
			reportAccepted, ok := destCCIP.ReportAcceptedWatcher[seqNum]
			destCCIP.ReportAcceptedWatcherMu.Unlock()
			if ok && reportAccepted != nil {
				receivedAt := time.Now().UTC()
				hdr, err := destCCIP.Common.ChainClient.HeaderByNumber(ctx, big.NewInt(int64(reportAccepted.Raw.BlockNumber)))
				if err == nil {
					receivedAt = hdr.Timestamp
				}

				totalTime := receivedAt.Sub(prevEventAt)
				// we cannot calculate the exact time at which block was finalized
				// as a result sometimes we get a time which is slightly after the block was marked as finalized
				// in such cases we get a negative time difference between finalized and report accepted if the commit
				// has happened almost immediately after block being finalized
				// in such cases we set the time difference to 1 second
				if totalTime < 0 {
					lggr.Warn().
						Uint64("seqNum", seqNum).
						Time("finalized at", prevEventAt).
						Time("ReportAccepted at", receivedAt).
						Msg("ReportAccepted event received before finalized timestamp")
					totalTime = time.Second
				}
				receipt, err := destCCIP.Common.ChainClient.GetTxReceipt(reportAccepted.Raw.TxHash)
				if err != nil {
					lggr.Warn().Msg("Failed to get receipt for ReportAccepted event")
				}
				reports.UpdatePhaseStats(reqNo, seqNum, testreporters.Commit, totalTime, testreporters.Success,
					testreporters.TransactionStats{
						GasUsed:    receipt.GasUsed,
						TxHash:     reportAccepted.Raw.TxHash.String(),
						CommitRoot: fmt.Sprintf("%x", reportAccepted.Report.MerkleRoot),
					})
				return &reportAccepted.Report, receivedAt, nil
			}
		case <-ctx.Done():
			reports.UpdatePhaseStats(reqNo, seqNum, testreporters.Commit, time.Since(prevEventAt), testreporters.Failure)
			return nil, time.Now().UTC(), fmt.Errorf("ReportAccepted is not found for seq num %d lane %d-->%d",
				seqNum, destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID())
		}
	}
}

func (destCCIP *DestCCIPModule) AssertReportBlessed(
	lggr zerolog.Logger,
	reqNo int64,
	seqNum uint64,
	timeout time.Duration,
	CommitReport commit_store.CommitStoreCommitReport,
	prevEventAt time.Time,
	reports *testreporters.CCIPLaneStats,
) (time.Time, error) {
	if destCCIP.Common.ARM == nil {
		lggr.Info().Interface("commit store interval", CommitReport.Interval).Hex("Root", CommitReport.MerkleRoot[:]).Msg("Skipping ReportBlessed check for mock ARM")
		return prevEventAt, nil
	}
	lggr.Info().Interface("commit store interval", CommitReport.Interval).Msg("Waiting for Report To be blessed")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			destCCIP.ReportBlessedWatcherMu.Lock()
			vLogs, ok := destCCIP.ReportBlessedWatcher[CommitReport.MerkleRoot]
			destCCIP.ReportBlessedWatcherMu.Unlock()
			receivedAt := time.Now().UTC()
			if ok && vLogs != nil {
				hdr, err := destCCIP.Common.ChainClient.HeaderByNumber(ctx, big.NewInt(int64(vLogs.BlockNumber)))
				if err == nil {
					receivedAt = hdr.Timestamp
				}
				receipt, err := destCCIP.Common.ChainClient.GetTxReceipt(vLogs.TxHash)
				if err != nil {
					lggr.Fatal().Err(err).Msg("Failed to get receipt for ReportBlessed event")
				}
				reports.UpdatePhaseStats(reqNo, seqNum, testreporters.ReportBlessed, receivedAt.Sub(prevEventAt), testreporters.Success,
					testreporters.TransactionStats{
						GasUsed:    receipt.GasUsed,
						TxHash:     vLogs.TxHash.String(),
						CommitRoot: fmt.Sprintf("%x", CommitReport.MerkleRoot),
					})
				return receivedAt, nil
			}
		case <-ctx.Done():
			reports.UpdatePhaseStats(reqNo, seqNum, testreporters.ReportBlessed, time.Since(prevEventAt), testreporters.Failure)
			return time.Now().UTC(), fmt.Errorf("ReportBlessed is not found for interval %+v lane %d-->%d",
				CommitReport.Interval, destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID())
		}
	}
}

func (destCCIP *DestCCIPModule) AssertSeqNumberExecuted(
	lggr zerolog.Logger,
	reqNo int64,
	seqNumberBefore uint64,
	timeout time.Duration,
	timeNow time.Time,
	reports *testreporters.CCIPLaneStats,
) error {
	lggr.Info().Int64("seqNum", int64(seqNumberBefore)).Msg("Waiting to be executed")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			if destCCIP.NextSeqNumToCommit.Load() > seqNumberBefore {
				return nil
			}
			seqNumberAfter, err := destCCIP.CommitStore.Instance.GetExpectedNextSequenceNumber(nil)
			if err != nil {
				reports.UpdatePhaseStats(reqNo, seqNumberBefore, testreporters.Commit, time.Since(timeNow), testreporters.Failure)
				return fmt.Errorf("error %+v in GetNextExpectedSeqNumber by commitStore for seqNum %d lane %d-->%d",
					err, seqNumberBefore+1, destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID())
			}
			if seqNumberAfter > seqNumberBefore {
				destCCIP.NextSeqNumToCommit.Store(seqNumberAfter)
				return nil
			}
		case <-ctx.Done():
			reports.UpdatePhaseStats(reqNo, seqNumberBefore, testreporters.Commit, time.Since(timeNow), testreporters.Failure)
			return fmt.Errorf("sequence number is not increased for seq num %d lane %d-->%d",
				seqNumberBefore, destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID())
		}
	}
}

func DefaultDestinationCCIPModule(logger zerolog.Logger, chainClient blockchain.EVMClient, sourceChainId uint64, sourceChain string, ccipCommon *CCIPCommon) (*DestCCIPModule, error) {
	cmn, err := ccipCommon.Copy(logger, chainClient)
	if err != nil {
		return nil, err
	}
	return &DestCCIPModule{
		Common:                  cmn,
		SourceChainId:           sourceChainId,
		SourceNetworkName:       sourceChain,
		ReportAcceptedWatcherMu: &sync.Mutex{},
		ReportAcceptedWatcher:   make(map[uint64]*commit_store.CommitStoreReportAccepted),
		ExecStateChangedMu:      &sync.Mutex{},
		ExecStateChangedWatcher: make(map[uint64]*evm_2_evm_offramp.EVM2EVMOffRampExecutionStateChanged),
		ReportBlessedWatcherMu:  &sync.Mutex{},
		ReportBlessedWatcher:    make(map[[32]byte]*types.Log),
		NextSeqNumToCommit:      atomic.NewUint64(1),
	}, nil
}

type CCIPRequest struct {
	txHash                  string
	txConfirmationTimestamp time.Time
}

func CCIPRequestFromTxHash(txHash common.Hash, chainClient blockchain.EVMClient) (CCIPRequest, *types.Receipt, error) {
	txConfirmationTimestamp := time.Now().UTC()
	rcpt, err := chainClient.GetTxReceipt(txHash)
	if err != nil {
		return CCIPRequest{}, nil, err
	}

	hdr, err := chainClient.HeaderByNumber(context.Background(), rcpt.BlockNumber)
	if err != nil {
		return CCIPRequest{}, nil, err
	}
	txConfirmationTimestamp = hdr.Timestamp

	return CCIPRequest{
		txHash:                  txHash.Hex(),
		txConfirmationTimestamp: txConfirmationTimestamp,
	}, rcpt, nil
}

type CCIPLane struct {
	Test                    *testing.T
	Logger                  zerolog.Logger
	SourceNetworkName       string
	DestNetworkName         string
	SourceChain             blockchain.EVMClient
	DestChain               blockchain.EVMClient
	Source                  *SourceCCIPModule
	Dest                    *DestCCIPModule
	TestEnv                 *CCIPTestEnv
	NumberOfReq             int
	Reports                 *testreporters.CCIPLaneStats
	Balance                 *BalanceSheet
	StartBlockOnSource      uint64
	StartBlockOnDestination uint64
	SentReqs                map[int64]CCIPRequest
	TotalFee                *big.Int // total fee for all the requests. Used for balance validation.
	ValidationTimeout       time.Duration
	Context                 context.Context
	SrcNetworkLaneCfg       *laneconfig.LaneConfig
	DstNetworkLaneCfg       *laneconfig.LaneConfig
	Subscriptions           []event.Subscription
}

func (lane *CCIPLane) UpdateLaneConfig() {
	var btAddresses, btpAddresses []string

	for i, bt := range lane.Source.Common.BridgeTokens {
		btAddresses = append(btAddresses, bt.Address())
		btpAddresses = append(btpAddresses, lane.Source.Common.BridgeTokenPools[i].Address())
	}
	lane.SrcNetworkLaneCfg.CommonContracts = laneconfig.CommonContracts{
		FeeToken:         lane.Source.Common.FeeToken.Address(),
		BridgeTokens:     btAddresses,
		BridgeTokenPools: btpAddresses,
		ARM:              lane.Source.Common.ARMContract.Hex(),
		Router:           lane.Source.Common.Router.Address(),
		PriceRegistry:    lane.Source.Common.PriceRegistry.Address(),
		WrappedNative:    lane.Source.Common.WrappedNative.Hex(),
	}
	if lane.Source.Common.ARM == nil {
		lane.SrcNetworkLaneCfg.CommonContracts.IsMockARM = true
	}
	lane.SrcNetworkLaneCfg.SrcContractsMu.Lock()
	lane.SrcNetworkLaneCfg.SrcContracts[lane.Source.DestNetworkName] = laneconfig.SourceContracts{
		OnRamp:     lane.Source.OnRamp.Address(),
		DepolyedAt: lane.Source.SrcStartBlock,
	}
	lane.SrcNetworkLaneCfg.SrcContractsMu.Unlock()
	btAddresses, btpAddresses = []string{}, []string{}

	for i, bt := range lane.Dest.Common.BridgeTokens {
		btAddresses = append(btAddresses, bt.Address())
		btpAddresses = append(btpAddresses, lane.Dest.Common.BridgeTokenPools[i].Address())
	}
	lane.DstNetworkLaneCfg.CommonContracts = laneconfig.CommonContracts{
		FeeToken:         lane.Dest.Common.FeeToken.Address(),
		BridgeTokens:     btAddresses,
		BridgeTokenPools: btpAddresses,
		ARM:              lane.Dest.Common.ARMContract.Hex(),
		Router:           lane.Dest.Common.Router.Address(),
		PriceRegistry:    lane.Dest.Common.PriceRegistry.Address(),
		WrappedNative:    lane.Dest.Common.WrappedNative.Hex(),
	}
	if lane.Dest.Common.ARM == nil {
		lane.DstNetworkLaneCfg.CommonContracts.IsMockARM = true
	}
	lane.DstNetworkLaneCfg.DestContractsMu.Lock()
	lane.DstNetworkLaneCfg.DestContracts[lane.Dest.SourceNetworkName] = laneconfig.DestContracts{
		OffRamp:      lane.Dest.OffRamp.Address(),
		CommitStore:  lane.Dest.CommitStore.Address(),
		ReceiverDapp: lane.Dest.ReceiverDapp.Address(),
	}
	lane.DstNetworkLaneCfg.DestContractsMu.Unlock()
}

func (lane *CCIPLane) RecordStateBeforeTransfer() {
	// collect the balance assert.ment to verify balances after transfer
	bal, err := testhelpers.GetBalances(lane.Test, lane.Source.CollectBalanceRequirements())
	require.NoError(lane.Test, err, "fetching source balance")
	lane.Balance.RecordBalance(bal)

	bal, err = testhelpers.GetBalances(lane.Test, lane.Dest.CollectBalanceRequirements())
	require.NoError(lane.Test, err, "fetching dest balance")
	lane.Balance.RecordBalance(bal)

	// save the current block numbers to use in various filter log requests
	lane.StartBlockOnSource, err = lane.Source.Common.ChainClient.LatestBlockNumber(context.Background())
	require.NoError(lane.Test, err, "Getting current block should be successful in source chain")
	lane.StartBlockOnDestination, err = lane.Dest.Common.ChainClient.LatestBlockNumber(context.Background())
	require.NoError(lane.Test, err, "Getting current block should be successful in dest chain")
	lane.TotalFee = big.NewInt(0)
	lane.NumberOfReq = 0
	lane.SentReqs = make(map[int64]CCIPRequest)
}

func (lane *CCIPLane) AddToSentReqs(txHash common.Hash) (*types.Receipt, error) {
	request, rcpt, err := CCIPRequestFromTxHash(txHash, lane.Source.Common.ChainClient)
	if err != nil {
		return rcpt, fmt.Errorf("could not get request from tx hash %s: %+v", txHash.Hex(), err)
	}
	lane.SentReqs[int64(lane.NumberOfReq+1)] = request
	lane.NumberOfReq++
	return rcpt, nil
}

func (lane *CCIPLane) SendRequests(noOfRequests int, msgType string) error {
	for i := 1; i <= noOfRequests; i++ {
		msg := fmt.Sprintf("msg %d", i)
		txHash, txConfirmationDur, fee, err := lane.Source.SendRequest(
			lane.Dest.ReceiverDapp.EthAddress,
			msgType, msg,
			common.HexToAddress(lane.Source.Common.FeeToken.Address()),
		)
		if err != nil {
			lane.Reports.UpdatePhaseStats(int64(lane.NumberOfReq+i), 0,
				testreporters.TX, txConfirmationDur, testreporters.Failure)
			return fmt.Errorf("could not send request: %+v", err)
		}
		err = lane.Source.Common.ChainClient.WaitForEvents()
		if err != nil {
			lane.Reports.UpdatePhaseStats(int64(lane.NumberOfReq+i), 0,
				testreporters.TX, txConfirmationDur, testreporters.Failure)
			return fmt.Errorf("could not send request: %+v", err)
		}

		noOfTokens := len(lane.Source.TransferAmount)
		if msgType == DataOnlyTransfer {
			noOfTokens = 0
		}
		rcpt, err := lane.AddToSentReqs(txHash)
		if err != nil {
			lane.Reports.UpdatePhaseStats(int64(lane.NumberOfReq+i), 0,
				testreporters.TX, txConfirmationDur, testreporters.Failure)
			return err
		}
		lane.Reports.UpdatePhaseStats(int64(lane.NumberOfReq+i), 0,
			testreporters.TX, txConfirmationDur, testreporters.Success, testreporters.TransactionStats{
				Fee:                fee.String(),
				GasUsed:            rcpt.GasUsed,
				TxHash:             txHash.Hex(),
				NoOfTokensSent:     noOfTokens,
				MessageBytesLength: len([]byte(msg)),
			})
		lane.TotalFee = bigmath.Add(lane.TotalFee, fee)
	}

	return nil
}

func (lane *CCIPLane) ValidateRequests() {
	for i, tx := range lane.SentReqs {
		require.NoError(lane.Test, lane.ValidateRequestByTxHash(tx.txHash, tx.txConfirmationTimestamp, i),
			"validating request events by tx hash")
	}
	// Asserting balances reliably work only for simulated private chains. The testnet contract balances might get updated by other transactions
	// verify the fee amount is deducted from sender, added to receiver token balances and
	if len(lane.Source.TransferAmount) > 0 {
		lane.Source.UpdateBalance(int64(lane.NumberOfReq), lane.TotalFee, lane.Balance)
		lane.Dest.UpdateBalance(lane.Source.TransferAmount, int64(lane.NumberOfReq), lane.Balance)
	}
}

func (lane *CCIPLane) ValidateRequestByTxHash(txHash string, txConfirmattion time.Time, reqNo int64) error {
	msgLog, ccipSendReqGenAt, err := lane.Source.AssertEventCCIPSendRequested(
		lane.Logger, reqNo, txHash, lane.ValidationTimeout, txConfirmattion, lane.Reports)
	if err != nil || msgLog == nil {
		return fmt.Errorf("could not validate CCIPSendRequested event: %+v", err)
	}
	seqNumber := msgLog.Message.SequenceNumber

	sourceLogFinalizedAt, _, err := lane.Source.AssertSendRequestedLogFinalized(lane.Logger, reqNo, seqNumber, msgLog, ccipSendReqGenAt, lane.Reports)
	if err != nil {
		return fmt.Errorf("could not finalize CCIPSendRequested event: %+v", err)
	}
	err = lane.Dest.AssertSeqNumberExecuted(lane.Logger, reqNo, seqNumber, lane.ValidationTimeout, sourceLogFinalizedAt, lane.Reports)
	if err != nil {
		return fmt.Errorf("could not validate seq number increase at commit store: %+v", err)
	}

	// Verify whether commitStore has accepted the report
	commitReport, reportAcceptedAt, err := lane.Dest.AssertEventReportAccepted(lane.Logger, reqNo, seqNumber, lane.ValidationTimeout, sourceLogFinalizedAt, lane.Reports)
	if err != nil || commitReport == nil {
		return fmt.Errorf("could not validate ReportAccepted event: %+v", err)
	}

	reportBlessedAt, err := lane.Dest.AssertReportBlessed(lane.Logger, reqNo, seqNumber, lane.ValidationTimeout, *commitReport, reportAcceptedAt, lane.Reports)
	if err != nil {
		return fmt.Errorf("could not validate ReportBlessed event: %+v", err)
	}
	// Verify whether the execution state is changed and the transfer is successful
	err = lane.Dest.AssertEventExecutionStateChanged(lane.Logger, reqNo, seqNumber, lane.ValidationTimeout, reportBlessedAt, lane.Reports)
	if err != nil {
		return fmt.Errorf("could not validate ExecutionStateChanged event: %+v", err)
	}
	return nil
}

func (lane *CCIPLane) StartEventWatchers() error {
	if !lane.Source.Common.ChainClient.NetworkSimulated() &&
		lane.Source.Common.ChainClient.GetNetworkConfig().FinalityDepth == 0 {
		err := lane.Source.Common.ChainClient.PollFinality()
		if err != nil {
			return err
		}
	}

	sendReqEvent := make(chan *evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested)
	sub, err := lane.Source.OnRamp.Instance.WatchCCIPSendRequested(nil, sendReqEvent)
	if err != nil {
		return err
	}
	lane.Subscriptions = append(lane.Subscriptions, sub)
	go func() {
		for {
			e := <-sendReqEvent
			lane.Logger.Info().Msgf("CCIPSendRequested event received for seq number %d", e.Message.SequenceNumber)
			lane.Source.CCIPSendRequestedWatcherMu.Lock()
			lane.Source.CCIPSendRequestedWatcher[e.Raw.TxHash.Hex()] = e
			lane.Source.CCIPSendRequestedWatcherMu.Unlock()
		}
	}()
	reportAcceptedEvent := make(chan *commit_store.CommitStoreReportAccepted)
	sub, err = lane.Dest.CommitStore.Instance.WatchReportAccepted(nil, reportAcceptedEvent)
	if err != nil {
		return err
	}

	lane.Subscriptions = append(lane.Subscriptions, sub)

	go func() {
		for {
			e := <-reportAcceptedEvent
			lane.Dest.ReportAcceptedWatcherMu.Lock()
			for i := e.Report.Interval.Min; i <= e.Report.Interval.Max; i++ {
				lane.Dest.ReportAcceptedWatcher[i] = e
			}
			lane.Dest.ReportAcceptedWatcherMu.Unlock()
		}
	}()

	if lane.Dest.Common.ARM != nil {
		reportBlessedEvent := make(chan *arm_contract.ARMContractTaggedRootBlessed)
		sub, err = lane.Dest.Common.ARM.Instance.WatchTaggedRootBlessed(nil, reportBlessedEvent, nil)
		if err != nil {
			return err
		}

		lane.Subscriptions = append(lane.Subscriptions, sub)

		go func() {
			for {
				e := <-reportBlessedEvent
				lane.Logger.Info().Msgf("TaggedRootBlessed event received for root %x", e.TaggedRoot.Root)
				lane.Dest.ReportBlessedWatcherMu.Lock()
				if e.TaggedRoot.CommitStore == lane.Dest.CommitStore.EthAddress {
					lane.Dest.ReportBlessedWatcher[e.TaggedRoot.Root] = &e.Raw
				}
				lane.Dest.ReportBlessedWatcherMu.Unlock()
			}
		}()
	}
	execStateChangedEvent := make(chan *evm_2_evm_offramp.EVM2EVMOffRampExecutionStateChanged)
	sub, err = lane.Dest.OffRamp.Instance.WatchExecutionStateChanged(nil, execStateChangedEvent, nil, nil)
	if err != nil {
		return err
	}

	lane.Subscriptions = append(lane.Subscriptions, sub)

	go func() {
		for {
			e := <-execStateChangedEvent
			lane.Logger.Info().Msgf("Execution state changed event received for seq number %d", e.SequenceNumber)
			lane.Dest.ExecStateChangedMu.Lock()
			lane.Dest.ExecStateChangedWatcher[e.SequenceNumber] = e
			lane.Dest.ExecStateChangedMu.Unlock()
		}
	}()
	return nil
}

func (lane *CCIPLane) CleanUp(clearFees bool) error {
	lane.Logger.Info().Msg("Cleaning up lane")
	if !lane.Source.Common.ChainClient.NetworkSimulated() &&
		lane.Source.Common.ChainClient.GetNetworkConfig().FinalityDepth == 0 {
		lane.Source.Common.ChainClient.CancelFinalityPolling()
	}
	for _, sub := range lane.Subscriptions {
		sub.Unsubscribe()
	}
	// recover fees from onRamp contract
	if clearFees && !lane.Source.Common.ChainClient.NetworkSimulated() {
		err := lane.Source.PayCCIPFeeToOwnerAddress()
		if err != nil {
			return err
		}
	}
	err := lane.Dest.Common.ChainClient.Close()
	if err != nil {
		return err
	}
	return lane.Source.Common.ChainClient.Close()
}

// DeployNewCCIPLane sets up a lane and initiates lane.Source and lane.Destination
// If configureCLNodes is true it sets up jobs and contract config for the lane
func (lane *CCIPLane) DeployNewCCIPLane(
	numOfCommitNodes int,
	commitAndExecOnSameDON bool,
	sourceCommon *CCIPCommon,
	destCommon *CCIPCommon,
	transferAmounts []*big.Int,
	bootstrapAdded *atomic.Bool,
	configureCLNodes bool,
	jobErrGroup *errgroup.Group,
) (*laneconfig.LaneConfig, *laneconfig.LaneConfig, error) {
	var err error
	env := lane.TestEnv
	sourceChainClient := lane.SourceChain
	destChainClient := lane.DestChain

	if sourceCommon == nil {
		return nil, nil, errors.WithStack(fmt.Errorf("common contracts for source chain %s not found", sourceChainClient.GetChainID().String()))
	}

	if destCommon == nil {
		return nil, nil, errors.WithStack(fmt.Errorf("common contracts for destination chain %s not found", destChainClient.GetChainID().String()))
	}

	lane.Source, err = DefaultSourceCCIPModule(
		lane.Logger,
		sourceChainClient, destChainClient.GetChainID().Uint64(),
		destChainClient.GetNetworkName(), transferAmounts, sourceCommon)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	lane.Dest, err = DefaultDestinationCCIPModule(
		lane.Logger,
		destChainClient, sourceChainClient.GetChainID().Uint64(),
		sourceChainClient.GetNetworkName(), destCommon)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	srcConf := lane.SrcNetworkLaneCfg

	destConf := lane.DstNetworkLaneCfg

	// deploy all source contracts
	err = lane.Source.DeployContracts(srcConf)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	// deploy all destination contracts
	err = lane.Dest.DeployContracts(*lane.Source, destConf)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	lane.UpdateLaneConfig()

	// if lane is being set up for already configured CL nodes and contracts
	// no further action is necessary
	if !configureCLNodes {
		return lane.SrcNetworkLaneCfg, lane.DstNetworkLaneCfg, nil
	}

	if env == nil {
		return lane.SrcNetworkLaneCfg, lane.DstNetworkLaneCfg, errors.WithStack(errors.New("test environment not set"))
	}
	// wait for the CL nodes to be ready before moving ahead with job creation
	err = env.CLNodeWithKeyReady.Wait()
	if err != nil {
		return lane.SrcNetworkLaneCfg, lane.DstNetworkLaneCfg, errors.WithStack(err)
	}
	clNodesWithKeys := env.CLNodesWithKeys
	// set up ocr2 jobs
	clNodes, exists := clNodesWithKeys[lane.Dest.Common.ChainClient.GetChainID().String()]
	if !exists {
		return lane.SrcNetworkLaneCfg, lane.DstNetworkLaneCfg, fmt.Errorf("could not find CL nodes for %s", lane.Dest.Common.ChainClient.GetChainID().String())
	}

	// first node is the bootstrapper
	bootstrapCommit := clNodes[0]
	var bootstrapExec *client.CLNodesWithKeys
	var execNodes []*client.CLNodesWithKeys
	commitNodes := clNodes[1:]
	env.commitNodeStartIndex = 1
	env.execNodeStartIndex = 1
	env.numOfAllowedFaultyExec = 1
	env.numOfAllowedFaultyCommit = 1
	env.numOfCommitNodes = numOfCommitNodes
	env.numOfExecNodes = numOfCommitNodes
	if !commitAndExecOnSameDON {
		if len(clNodes) < 11 {
			return lane.SrcNetworkLaneCfg, lane.DstNetworkLaneCfg, fmt.Errorf("not enough CL nodes for separate commit and execution nodes")
		}
		bootstrapExec = clNodes[1] // for a set-up of different commit and execution nodes second node is the bootstrapper for execution nodes
		commitNodes = clNodes[2 : 2+numOfCommitNodes]
		execNodes = clNodes[2+numOfCommitNodes:]
		env.commitNodeStartIndex = 2
		env.execNodeStartIndex = 7
		env.numOfCommitNodes = len(commitNodes)
		env.numOfExecNodes = len(execNodes)
	} else {
		execNodes = commitNodes
	}
	// save the current block numbers. If there is a delay between job start up and ocr config set up, the jobs will
	// replay the log polling from these mentioned block number. The dest block number should ideally be the block number on which
	// contract config is set and the source block number should be the one on which the ccip send request is performed.
	// Here for simplicity we are just taking the current block number just before the job is created.
	currentBlockOnDest, err := destChainClient.LatestBlockNumber(context.Background())
	if err != nil {
		return lane.SrcNetworkLaneCfg, lane.DstNetworkLaneCfg, fmt.Errorf("getting current block should be successful in destination chain %+v", err)
	}

	tokenUSDMap := make(map[string]string)
	for _, token := range lane.Dest.Common.BridgeTokens {
		tokenUSDMap[token.Address()] = LinkToUSD.String()
	}

	tokenUSDMap[lane.Dest.Common.FeeToken.Address()] = LinkToUSD.String()
	tokenUSDMap[lane.Source.Common.WrappedNative.Hex()] = WrappedNativeToUSD.String()
	tokenUSDMap[lane.Dest.Common.WrappedNative.Hex()] = WrappedNativeToUSD.String()
	lane.Logger.Info().Interface("tokenUSDMap", tokenUSDMap).Msg("tokenUSDMap")

	jobParams := integrationtesthelpers.CCIPJobSpecParams{
		OffRamp:                lane.Dest.OffRamp.EthAddress,
		CommitStore:            lane.Dest.CommitStore.EthAddress,
		SourceChainName:        sourceChainClient.GetNetworkName(),
		DestChainName:          destChainClient.GetNetworkName(),
		DestEvmChainId:         destChainClient.GetChainID().Uint64(),
		SourceStartBlock:       lane.Source.SrcStartBlock,
		TokenPricesUSDPipeline: StaticTokenFeeForMultipleTokenAddr(tokenUSDMap),
		DestStartBlock:         currentBlockOnDest,
	}

	if !bootstrapAdded.Load() {
		bootstrapAdded.Store(true)
		err := CreateBootstrapJob(jobParams, bootstrapCommit, bootstrapExec)
		if err != nil {
			return lane.SrcNetworkLaneCfg, lane.DstNetworkLaneCfg, errors.WithStack(err)
		}
	}

	bootstrapCommitP2PId := bootstrapCommit.KeysBundle.P2PKeys.Data[0].Attributes.PeerID
	var bootstrapExecP2PId string
	var p2pBootstrappersExec, p2pBootstrappersCommit *client.P2PData
	if bootstrapExec == nil {
		bootstrapExec = bootstrapCommit
		bootstrapExecP2PId = bootstrapCommitP2PId
	} else {
		bootstrapExecP2PId = bootstrapExec.KeysBundle.P2PKeys.Data[0].Attributes.PeerID
		p2pBootstrappersExec = &client.P2PData{
			InternalIP: bootstrapExec.Node.InternalIP(),
			PeerID:     bootstrapExecP2PId,
		}
	}
	p2pBootstrappersCommit = &client.P2PData{
		InternalIP: bootstrapCommit.Node.InternalIP(),
		PeerID:     bootstrapCommitP2PId,
	}

	jobParams.P2PV2Bootstrappers = []string{p2pBootstrappersCommit.P2PV2Bootstrapper()}

	// set up ocr2 config
	err = SetOCR2Configs(commitNodes, execNodes, *lane.Dest)
	if err != nil {
		return lane.SrcNetworkLaneCfg, lane.DstNetworkLaneCfg, errors.WithStack(err)
	}

	err = CreateOCR2CCIPCommitJobs(lane.Logger, jobParams, commitNodes, env.nodeMutexes, jobErrGroup)
	if err != nil {
		return lane.SrcNetworkLaneCfg, lane.DstNetworkLaneCfg, errors.WithStack(err)
	}
	if p2pBootstrappersExec != nil {
		jobParams.P2PV2Bootstrappers = []string{p2pBootstrappersExec.P2PV2Bootstrapper()}
	}

	err = CreateOCR2CCIPExecutionJobs(lane.Logger, jobParams, execNodes, env.nodeMutexes, jobErrGroup)
	if err != nil {
		return lane.SrcNetworkLaneCfg, lane.DstNetworkLaneCfg, errors.WithStack(err)
	}

	lane.Dest.Common.ChainClient.ParallelTransactions(false)
	lane.Source.Common.ChainClient.ParallelTransactions(false)

	return lane.SrcNetworkLaneCfg, lane.DstNetworkLaneCfg, nil
}

// SetOCR2Configs sets the oracle config in ocr2 contracts
// nil value in execNodes denotes commit and execution jobs are to be set up in same DON
func SetOCR2Configs(commitNodes, execNodes []*client.CLNodesWithKeys, destCCIP DestCCIPModule) error {
	rootSnooze := models.MustMakeDuration(7 * time.Minute)
	inflightExpiry := models.MustMakeDuration(3 * time.Minute)
	if destCCIP.Common.ChainClient.NetworkSimulated() {
		rootSnooze = models.MustMakeDuration(RootSnoozeTimeSimulated)
		inflightExpiry = models.MustMakeDuration(InflightExpirySimulated)
	}
	signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, err := contracts.NewOffChainAggregatorV2Config(commitNodes, ccipConfig.CommitOffchainConfig{
		SourceFinalityDepth:   1,
		DestFinalityDepth:     1,
		FeeUpdateHeartBeat:    models.MustMakeDuration(10 * time.Second), // reduce the heartbeat to 10 sec for faster fee updates
		FeeUpdateDeviationPPB: 1e6,
		MaxGasPrice:           200e9,
		InflightCacheExpiry:   inflightExpiry,
	}, ccipConfig.CommitOnchainConfig{
		PriceRegistry: destCCIP.Common.PriceRegistry.EthAddress,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	err = destCCIP.CommitStore.SetOCR2Config(signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
	if err != nil {
		return errors.WithStack(err)
	}

	nodes := commitNodes
	// if commit and exec job is set up in different DON
	if len(execNodes) > 0 {
		nodes = execNodes
	}
	if destCCIP.OffRamp != nil {
		signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, err = contracts.NewOffChainAggregatorV2Config(nodes, ccipConfig.ExecOffchainConfig{
			SourceFinalityDepth:         1,
			DestOptimisticConfirmations: 1,
			DestFinalityDepth:           1,
			BatchGasLimit:               5_000_000,
			RelativeBoostPerWaitHour:    0.7,
			MaxGasPrice:                 200e9,
			InflightCacheExpiry:         inflightExpiry,
			RootSnoozeTime:              rootSnooze,
		}, ccipConfig.ExecOnchainConfig{
			PermissionLessExecutionThresholdSeconds: 60 * 30,
			Router:                                  destCCIP.Common.Router.EthAddress,
			PriceRegistry:                           destCCIP.Common.PriceRegistry.EthAddress,
			MaxTokensLength:                         5,
			MaxDataSize:                             50000,
		})
		if err != nil {
			return errors.WithStack(err)
		}
		err = destCCIP.OffRamp.SetOCR2Config(signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return destCCIP.Common.ChainClient.WaitForEvents()
}

func CreateBootstrapJob(
	jobParams integrationtesthelpers.CCIPJobSpecParams,
	bootstrapCommit *client.CLNodesWithKeys,
	bootstrapExec *client.CLNodesWithKeys,
) error {
	_, err := bootstrapCommit.Node.MustCreateJob(jobParams.BootstrapJob(jobParams.CommitStore.Hex()))
	if err != nil {
		return fmt.Errorf("shouldn't fail creating bootstrap job on bootstrap node %+v", err)
	}
	if bootstrapExec != nil {
		_, err := bootstrapExec.Node.MustCreateJob(jobParams.BootstrapJob(jobParams.OffRamp.Hex()))
		if err != nil {
			return fmt.Errorf("shouldn't fail creating bootstrap job on bootstrap node %+v", err)
		}
	}
	return nil
}

func CreateOCR2CCIPCommitJobs(
	lggr zerolog.Logger,
	jobParams integrationtesthelpers.CCIPJobSpecParams,
	commitNodes []*client.CLNodesWithKeys,
	mutexes []*sync.Mutex,
	group *errgroup.Group,
) error {
	ocr2SpecCommit, err := jobParams.CommitJobSpec()
	if err != nil {
		return errors.WithStack(err)
	}
	createJob := func(index int, node *client.CLNodesWithKeys, ocr2SpecCommit client.OCR2TaskJobSpec, mu *sync.Mutex) error {
		mu.Lock()
		defer mu.Unlock()
		ocr2SpecCommit.OCR2OracleSpec.OCRKeyBundleID.SetValid(node.KeysBundle.OCR2Key.Data.ID)
		ocr2SpecCommit.OCR2OracleSpec.TransmitterID.SetValid(node.KeysBundle.EthAddress)
		lggr.Info().Msgf("Creating CCIP-Commit job on OCR node %d job name %s", index+1, ocr2SpecCommit.Name)
		_, err = node.Node.MustCreateJob(&ocr2SpecCommit)
		if err != nil {
			return fmt.Errorf("shouldn't fail creating CCIP-Commit job on OCR node %d job name %s - %+v", index+1, ocr2SpecCommit.Name, err)
		}
		return nil
	}
	for i, node := range commitNodes {
		node := node
		i := i
		group.Go(func() error {
			return createJob(i, node, *ocr2SpecCommit, mutexes[i])
		})
	}
	return nil
}

func CreateOCR2CCIPExecutionJobs(
	lggr zerolog.Logger,
	jobParams integrationtesthelpers.CCIPJobSpecParams,
	execNodes []*client.CLNodesWithKeys,
	mutexes []*sync.Mutex,
	group *errgroup.Group,
) error {
	ocr2SpecExec, err := jobParams.ExecutionJobSpec()
	if err != nil {
		return errors.WithStack(err)
	}
	createJob := func(index int, node *client.CLNodesWithKeys, ocr2SpecExec client.OCR2TaskJobSpec, mu *sync.Mutex) error {
		mu.Lock()
		defer mu.Unlock()
		ocr2SpecExec.OCR2OracleSpec.OCRKeyBundleID.SetValid(node.KeysBundle.OCR2Key.Data.ID)
		ocr2SpecExec.OCR2OracleSpec.TransmitterID.SetValid(node.KeysBundle.EthAddress)
		lggr.Info().Msgf("Creating CCIP-Exec job on OCR node %d job name %s", index+1, ocr2SpecExec.Name)
		_, err = node.Node.MustCreateJob(&ocr2SpecExec)
		if err != nil {
			return fmt.Errorf("shouldn't fail creating CCIP-Exec job on OCR node %d job name %s - %+v", index+1,
				ocr2SpecExec.Name, err)
		}
		return nil
	}
	if ocr2SpecExec != nil {
		for i, node := range execNodes {
			node := node
			i := i
			group.Go(func() error {
				return createJob(i, node, *ocr2SpecExec, mutexes[i])
			})
		}
	}
	return nil
}

// TODO : keep it if there is a better mockserver implementation is found
func _(tokenAddr []string, mockserver *ctfClient.MockserverClient) string {
	source := ""
	right := ""
	for i, addr := range tokenAddr {
		url := fmt.Sprintf("%s/%s", mockserver.Config.ClusterURL, addr)
		source = source + fmt.Sprintf(`
token%d [type=http method=GET url="%s"];
token%d_parse [type=jsonparse path="Data,Result"];
token%d->token%d_parse;`, i+1, url, i+1, i+1, i+1)
		right = right + fmt.Sprintf(` \\\"%s\\\":$(token%d_parse),`, addr, i+1)
	}
	right = right[:len(right)-1]
	source = fmt.Sprintf(`%s
merge [type=merge left="{}" right="{%s}"];`, source, right)

	return source
}

func StaticTokenFeeForMultipleTokenAddr(tokenUSD map[string]string) string {
	right := ""
	for addr, value := range tokenUSD {
		right = right + fmt.Sprintf(`\\"%s\\":\\"%s\\",`, addr, value)
	}
	right = right[:len(right)-1]
	source := fmt.Sprintf(`merge [type=merge left="{}" right="{%s}"];`, right)

	return source
}

// SetMockServerWithSameTokenFeeConversionValue sets the mock responses in mockserver that are read by chainlink nodes
// to simulate different price feed value.
func SetMockServerWithSameTokenFeeConversionValue(
	ctx context.Context,
	tokenValueAddress map[string]interface{},
	mockserver *ctfClient.MockserverClient,
) error {
	valueAdditions, _ := errgroup.WithContext(ctx)
	for tokenAddr, value := range tokenValueAddress {
		path := fmt.Sprintf("/%s", tokenAddr)
		tokenValue := value
		valueAdditions.Go(func() error {
			log.Info().Str("path", path).
				Str("value", fmt.Sprintf("%v", tokenValue)).
				Msg(fmt.Sprintf("Setting mockserver response"))
			return mockserver.SetAnyValuePath(path, tokenValue)
		})
	}
	return valueAdditions.Wait()
}

type CCIPTestEnv struct {
	LocalCluster             *test_env.CLClusterTestEnv
	CLNodesWithKeys          map[string][]*client.CLNodesWithKeys // key - network chain-id
	CLNodes                  []*client.ChainlinkK8sClient
	nodeMutexes              []*sync.Mutex
	execNodeStartIndex       int
	commitNodeStartIndex     int
	numOfAllowedFaultyCommit int
	numOfAllowedFaultyExec   int
	numOfCommitNodes         int
	numOfExecNodes           int
	K8Env                    *environment.Environment
	CLNodeWithKeyReady       *errgroup.Group // denotes if keys are created in chainlink node and ready to be used for job creation
}

func (c *CCIPTestEnv) ChaosLabelForGeth(t *testing.T, srcChain, destChain string) {
	err := c.K8Env.Client.LabelChaosGroupByLabels(c.K8Env.Cfg.Namespace, map[string]string{
		"app": GethLabel(srcChain),
	}, ChaosGroupNetworkACCIPGeth)
	require.NoError(t, err)

	err = c.K8Env.Client.LabelChaosGroupByLabels(c.K8Env.Cfg.Namespace, map[string]string{
		"app": GethLabel(destChain),
	}, ChaosGroupNetworkBCCIPGeth)
	require.NoError(t, err)

	gethNetworksLabels := []string{GethLabel(srcChain), GethLabel(destChain)}
	for _, gethNetworkLabel := range gethNetworksLabels {
		err := c.K8Env.Client.AddLabel(c.K8Env.Cfg.Namespace,
			fmt.Sprintf("app=%s", gethNetworkLabel),
			fmt.Sprintf("geth=%s", ChaosGroupCCIPGeth))
		require.NoError(t, err)
	}
}

func (c *CCIPTestEnv) ChaosLabelForCLNodes(t *testing.T) {
	for i := c.commitNodeStartIndex; i < len(c.CLNodes); i++ {
		labelSelector := map[string]string{
			"app":      "chainlink-0",
			"instance": strconv.Itoa(i),
		}
		// commit node starts from index 2
		if i >= c.commitNodeStartIndex && i < c.commitNodeStartIndex+c.numOfCommitNodes {
			err := c.K8Env.Client.LabelChaosGroupByLabels(c.K8Env.Cfg.Namespace, labelSelector, ChaosGroupCommit)
			require.NoError(t, err)
		}
		if i >= c.commitNodeStartIndex && i < c.commitNodeStartIndex+c.numOfAllowedFaultyCommit+1 {
			err := c.K8Env.Client.LabelChaosGroupByLabels(c.K8Env.Cfg.Namespace, labelSelector, ChaosGroupCommitFaultyPlus)
			require.NoError(t, err)
		}
		if i >= c.commitNodeStartIndex && i < c.commitNodeStartIndex+c.numOfAllowedFaultyCommit {
			err := c.K8Env.Client.LabelChaosGroupByLabels(c.K8Env.Cfg.Namespace, labelSelector, ChaosGroupCommitFaulty)
			require.NoError(t, err)
		}
		if i >= c.execNodeStartIndex && i < c.execNodeStartIndex+c.numOfExecNodes {
			err := c.K8Env.Client.LabelChaosGroupByLabels(c.K8Env.Cfg.Namespace, labelSelector, ChaosGroupExecution)
			require.NoError(t, err)
		}
		if i >= c.execNodeStartIndex && i < c.execNodeStartIndex+c.numOfAllowedFaultyExec+1 {
			err := c.K8Env.Client.LabelChaosGroupByLabels(c.K8Env.Cfg.Namespace, labelSelector, ChaosGroupExecutionFaultyPlus)
			require.NoError(t, err)
		}
		if i >= c.execNodeStartIndex && i < c.execNodeStartIndex+c.numOfAllowedFaultyExec {
			err := c.K8Env.Client.LabelChaosGroupByLabels(c.K8Env.Cfg.Namespace, labelSelector, ChaosGroupExecutionFaulty)
			require.NoError(t, err)
		}
	}
}

func (c *CCIPTestEnv) SetUpNodesAndKeys(
	ctx context.Context,
	nodeFund *big.Float,
	chains []blockchain.EVMClient,
	logger zerolog.Logger,
) error {
	chainlinkNodes := make([]*client.ChainlinkClient, 0)

	//var err error
	if c.LocalCluster != nil {
		// for local cluster, fetch the values from the local cluster
		for _, chainlinkNode := range c.LocalCluster.CLNodes {
			chainlinkNodes = append(chainlinkNodes, chainlinkNode.API)
			c.nodeMutexes = append(c.nodeMutexes, &sync.Mutex{})
		}
	} else {
		// in case of k8s, we need to connect to the chainlink nodes
		log.Info().Msg("Connecting to launched resources")
		chainlinkK8sNodes, err := client.ConnectChainlinkNodes(c.K8Env)
		if err != nil {
			return errors.WithStack(err)
		}
		if len(chainlinkK8sNodes) == 0 {
			return fmt.Errorf("no CL node found")
		}

		for _, chainlinkNode := range chainlinkK8sNodes {
			chainlinkNode.ChainlinkClient.SetLogger(logger)
			chainlinkNode.ChainlinkClient.AddRetryAttempt(3)
			chainlinkNodes = append(chainlinkNodes, chainlinkNode.ChainlinkClient)
			c.nodeMutexes = append(c.nodeMutexes, &sync.Mutex{})
		}
		c.CLNodes = chainlinkK8sNodes
	}

	nodesWithKeys := make(map[string][]*client.CLNodesWithKeys)
	mu := &sync.Mutex{}
	//grp, _ := errgroup.WithContext(ctx)
	populateKeys := func(chain blockchain.EVMClient) error {
		log.Info().Str("chain id", chain.GetChainID().String()).Msg("creating node keys for chain")
		_, clNodes, err := client.CreateNodeKeysBundle(chainlinkNodes, "evm", chain.GetChainID().String())
		if err != nil {
			return errors.WithStack(err)
		}
		if len(clNodes) == 0 {
			return fmt.Errorf("no CL node with keys found for chain %s", chain.GetNetworkName())
		}
		mu.Lock()
		defer mu.Unlock()
		nodesWithKeys[chain.GetChainID().String()] = clNodes
		return nil
	}

	fund := func(ec blockchain.EVMClient) error {
		cfg := ec.GetNetworkConfig()
		if cfg == nil {
			return fmt.Errorf("blank network config")
		}
		c1, err := blockchain.ConcurrentEVMClient(*cfg, c.K8Env, ec, logger)
		if err != nil {
			return fmt.Errorf("getting concurrent evmclient chain %s %+v", ec.GetNetworkName(), err)
		}
		defer func() {
			if c1 != nil {
				c1.Close()
			}
		}()
		log.Info().Str("chain id", c1.GetChainID().String()).Msg("Funding Chainlink nodes for chain")
		err = actions.FundChainlinkNodesAddresses(chainlinkNodes[1:], c1, nodeFund)
		if err != nil {
			return fmt.Errorf("funding nodes for chain %s %+v", c1.GetNetworkName(), err)
		}
		return nil
	}

	for _, chain := range chains {
		err := populateKeys(chain)
		if err != nil {
			return err
		}
		err = fund(chain)
		if err != nil {
			return err
		}
	}

	c.CLNodesWithKeys = nodesWithKeys
	return nil
}

func AssertBalances(t *testing.T, bas []testhelpers.BalanceAssertion) {
	logEvent := log.Info()
	for _, b := range bas {
		actual := b.Getter(t, b.Address)
		assert.NotNil(t, actual, "%v getter return nil", b.Name)
		if b.Within == "" {
			assert.Equal(t, b.Expected, actual.String(), "wrong balance for %s got %s want %s", b.Name, actual, b.Expected)
			logEvent.Interface(b.Name, struct {
				Exp    string
				Actual string
			}{
				Exp:    b.Expected,
				Actual: actual.String(),
			})
		} else {
			bi, _ := big.NewInt(0).SetString(b.Expected, 10)
			withinI, _ := big.NewInt(0).SetString(b.Within, 10)
			high := big.NewInt(0).Add(bi, withinI)
			low := big.NewInt(0).Sub(bi, withinI)
			assert.Equal(t, -1, actual.Cmp(high),
				"wrong balance for %s got %s outside expected range [%s, %s]", b.Name, actual, low, high)
			assert.Equal(t, 1, actual.Cmp(low),
				"wrong balance for %s got %s outside expected range [%s, %s]", b.Name, actual, low, high)
			logEvent.Interface(b.Name, struct {
				ExpRange string
				Actual   string
			}{
				ExpRange: fmt.Sprintf("[%s, %s]", low, high),
				Actual:   actual.String(),
			})
		}
	}
	logEvent.Msg("balance assertions succeeded")
}

type BalFunc func(ctx context.Context, addr string) (*big.Int, error)

func GetterForLinkToken(getBalance BalFunc, addr string) func(t *testing.T, _ common.Address) *big.Int {
	return func(t *testing.T, _ common.Address) *big.Int {
		balance, err := getBalance(context.Background(), addr)
		assert.NoError(t, err)
		return balance
	}
}

type BalanceItem struct {
	Address         common.Address
	Getter          func(t *testing.T, addr common.Address) *big.Int
	PreviousBalance *big.Int
	AmtToAdd        *big.Int
	AmtToSub        *big.Int
}

type BalanceSheet struct {
	mu          *sync.Mutex
	Items       map[string]BalanceItem
	PrevBalance map[string]*big.Int
}

func (b *BalanceSheet) Update(key string, item BalanceItem) {
	b.mu.Lock()
	defer b.mu.Unlock()
	prev, ok := b.Items[key]
	if !ok {
		b.Items[key] = item
		return
	}
	amtToAdd, amtToSub := big.NewInt(0), big.NewInt(0)
	if prev.AmtToAdd != nil {
		amtToAdd = prev.AmtToAdd
	}
	if prev.AmtToSub != nil {
		amtToSub = prev.AmtToSub
	}
	if item.AmtToAdd != nil {
		amtToAdd = new(big.Int).Add(amtToAdd, item.AmtToAdd)
	}
	if item.AmtToSub != nil {
		amtToSub = new(big.Int).Add(amtToSub, item.AmtToSub)
	}

	b.Items[key] = BalanceItem{
		Address:  item.Address,
		Getter:   item.Getter,
		AmtToAdd: amtToAdd,
		AmtToSub: amtToSub,
	}
}

func (b *BalanceSheet) RecordBalance(bal map[string]*big.Int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	for key, value := range bal {
		if _, ok := b.PrevBalance[key]; !ok {
			b.PrevBalance[key] = value
		}
	}
}

func (b *BalanceSheet) Verify(t *testing.T) {
	var balAssertions []testhelpers.BalanceAssertion
	for key, item := range b.Items {
		prevBalance, ok := b.PrevBalance[key]
		require.Truef(t, ok, "previous balance is not captured for %s", key)
		exp := prevBalance
		if item.AmtToAdd != nil {
			exp = new(big.Int).Add(exp, item.AmtToAdd)
		}
		if item.AmtToSub != nil {
			exp = new(big.Int).Sub(exp, item.AmtToSub)
		}
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     key,
			Address:  item.Address,
			Getter:   item.Getter,
			Expected: exp.String(),
		})
	}
	AssertBalances(t, balAssertions)
}

func NewBalanceSheet() *BalanceSheet {
	return &BalanceSheet{
		mu:          &sync.Mutex{},
		Items:       make(map[string]BalanceItem),
		PrevBalance: make(map[string]*big.Int),
	}
}
