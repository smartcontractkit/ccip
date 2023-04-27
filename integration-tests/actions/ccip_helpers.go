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
	envclient "github.com/smartcontractkit/chainlink-env/client"
	"github.com/smartcontractkit/chainlink-env/pkg/helm/mockserver"
	mockservercfg "github.com/smartcontractkit/chainlink-env/pkg/helm/mockserver-cfg"
	"go.uber.org/atomic"

	"github.com/smartcontractkit/chainlink-env/environment"
	"github.com/smartcontractkit/chainlink-env/pkg/helm/chainlink"
	"github.com/smartcontractkit/chainlink-env/pkg/helm/reorg"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	ctfClient "github.com/smartcontractkit/chainlink-testing-framework/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink/integration-tests/client"
	"github.com/smartcontractkit/chainlink/integration-tests/contracts/ccip"
	"github.com/smartcontractkit/chainlink/integration-tests/contracts/ccip/laneconfig"
	"github.com/smartcontractkit/chainlink/integration-tests/testreporters"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/price_registry"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/router"
	ccipPlugin "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/testhelpers"
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
	RootSnoozeTime                = 1 * time.Minute
	RootSnoozeTimeSimulated       = 10 * time.Second
	InflightExpiry                = 1 * time.Minute
	InflightExpirySimulated       = 45 * time.Second
)

type CCIPTOMLEnv struct {
	Networks []blockchain.EVMNetwork
}

var (
	//go:embed clconfig/ccip-default.txt
	CLConfig             string
	DefaultCCIPCLNodeEnv = func(t *testing.T, networks []blockchain.EVMNetwork) string {
		ccipTOML, err := client.MarshallTemplate(
			CCIPTOMLEnv{
				Networks: networks,
			},
			"ccip env toml", CLConfig)
		require.NoError(t, err)
		fmt.Println("Configuration", ccipTOML)
		return ccipTOML
	}

	NetworkName = func(name string) string {
		return strings.ReplaceAll(strings.ToLower(name), " ", "-")
	}

	GethLabel = func(name string) string {
		return fmt.Sprintf("%s-ethereum-geth", name)
	}
	// ApprovedAmountToRouter is the default amount which gets approved for router so that it can transfer token and use the fee token for fee payment
	ApprovedAmountToRouter           = new(big.Int).Mul(big.NewInt(1e18), big.NewInt(100))
	ApprovedFeeAmountToRouter        = new(big.Int).Mul(big.NewInt(int64(FeeMultiplier)), big.NewInt(1e9))
	FeeMultiplier             uint64 = 12e17
	LinkToUSD                        = big.NewInt(8e18)
	WrappedNativeToUSD               = new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1.7e3))
)

type CCIPCommon struct {
	ChainClient        blockchain.EVMClient
	Deployer           *ccip.CCIPContractsDeployer
	FeeToken           *ccip.LinkToken
	FeeTokenPool       *ccip.LockReleaseTokenPool
	BridgeTokens       []*ccip.LinkToken // as of now considering the bridge token is same as link token
	TokenPrices        []*big.Int
	BridgeTokenPools   []*ccip.LockReleaseTokenPool
	RateLimiterConfig  ccip.RateLimiterConfig
	AFN                *ccip.AFN
	Router             *ccip.Router
	PriceRegistry      *ccip.PriceRegistry
	WrappedNative      common.Address
	ExistingDeployment bool
	deploy             *errgroup.Group
	poolFunds          *big.Int
}

func (ccipModule *CCIPCommon) CopyAddresses(ctx context.Context, chainClient blockchain.EVMClient, existingDeployment bool) *CCIPCommon {
	var pools []*ccip.LockReleaseTokenPool
	for _, pool := range ccipModule.BridgeTokenPools {
		pools = append(pools, &ccip.LockReleaseTokenPool{EthAddress: pool.EthAddress})
	}
	var tokens []*ccip.LinkToken
	for _, token := range ccipModule.BridgeTokens {
		tokens = append(tokens, &ccip.LinkToken{
			EthAddress: token.EthAddress,
		})
	}
	grp, _ := errgroup.WithContext(ctx)
	return &CCIPCommon{
		ChainClient: chainClient,
		Deployer:    nil,
		FeeToken: &ccip.LinkToken{
			EthAddress: ccipModule.FeeToken.EthAddress,
		},
		ExistingDeployment: existingDeployment,
		FeeTokenPool: &ccip.LockReleaseTokenPool{
			EthAddress: ccipModule.FeeTokenPool.EthAddress,
		},
		BridgeTokens:      tokens,
		TokenPrices:       ccipModule.TokenPrices,
		BridgeTokenPools:  pools,
		RateLimiterConfig: ccipModule.RateLimiterConfig,
		AFN: &ccip.AFN{
			EthAddress: ccipModule.AFN.EthAddress,
		},
		Router: &ccip.Router{
			EthAddress: ccipModule.Router.EthAddress,
		},
		WrappedNative: ccipModule.WrappedNative,
		deploy:        grp,
		poolFunds:     ccipModule.poolFunds,
	}
}

func (ccipModule *CCIPCommon) LoadContractAddresses(conf *laneconfig.LaneConfig) {
	if conf != nil {
		if common.IsHexAddress(conf.FeeToken) {
			ccipModule.FeeToken = &ccip.LinkToken{
				EthAddress: common.HexToAddress(conf.FeeToken),
			}
		}
		if conf.IsNativeFeeToken {
			ccipModule.FeeToken = &ccip.LinkToken{
				EthAddress: common.HexToAddress("0x0"),
			}
		}

		if common.IsHexAddress(conf.FeeTokenPool) {
			ccipModule.FeeTokenPool = &ccip.LockReleaseTokenPool{
				EthAddress: common.HexToAddress(conf.FeeTokenPool),
			}
		}
		if common.IsHexAddress(conf.Router) {
			ccipModule.Router = &ccip.Router{
				EthAddress: common.HexToAddress(conf.Router),
			}
		}
		if common.IsHexAddress(conf.AFN) {
			ccipModule.AFN = &ccip.AFN{
				EthAddress: common.HexToAddress(conf.AFN),
			}
		}
		if common.IsHexAddress(conf.PriceRegistry) {
			ccipModule.PriceRegistry = &ccip.PriceRegistry{
				EthAddress: common.HexToAddress(conf.PriceRegistry),
			}
		}
		if common.IsHexAddress(conf.WrappedNative) {
			ccipModule.WrappedNative = common.HexToAddress(conf.WrappedNative)
		}
		if len(conf.BridgeTokens) > 0 {
			var tokens []*ccip.LinkToken
			for _, token := range conf.BridgeTokens {
				if common.IsHexAddress(token) {
					tokens = append(tokens, &ccip.LinkToken{
						EthAddress: common.HexToAddress(token),
					})
				}
			}
			ccipModule.BridgeTokens = tokens
		}
		if len(conf.BridgeTokenPools) > 0 {
			var pools []*ccip.LockReleaseTokenPool
			for _, pool := range conf.BridgeTokenPools {
				if common.IsHexAddress(pool) {
					pools = append(pools, &ccip.LockReleaseTokenPool{
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
	for _, token := range ccipModule.BridgeTokens {
		err := token.Approve(ccipModule.Router.Address(), ApprovedAmountToRouter)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	if ccipModule.FeeToken.EthAddress != common.HexToAddress("0x0") {
		isApproved := false
		for _, token := range ccipModule.BridgeTokens {
			if token.EthAddress == ccipModule.FeeToken.EthAddress {
				isApproved = true
				break
			}
		}
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
	err := ccipModule.ChainClient.WaitForEvents()
	if err != nil {
		return errors.WithStack(err)
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

// DeployContracts deploys the contracts which are necessary in both source and dest chain
// This reuses common contracts for bidirectional lanes
func (ccipModule *CCIPCommon) DeployContracts(noOfTokens int, conf *laneconfig.LaneConfig) error {
	var err error
	cd := ccipModule.Deployer

	ccipModule.LoadContractAddresses(conf)
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
	if ccipModule.FeeTokenPool == nil {
		// token pool for fee token
		ccipModule.FeeTokenPool, err = cd.DeployLockReleaseTokenPoolContract(ccipModule.FeeToken.Address())
		if err != nil {
			return fmt.Errorf("deploying fee Token pool shouldn't fail %+v", err)
		}
	} else {
		pool, err := cd.NewLockReleaseTokenPoolContract(ccipModule.FeeTokenPool.EthAddress)
		if err != nil {
			return fmt.Errorf("getting new fee token pool contract shouldn't fail %+v", err)
		}
		ccipModule.FeeTokenPool = pool
	}

	if len(ccipModule.BridgeTokens) < noOfTokens {
		// deploy bridge token.
		for i := len(ccipModule.BridgeTokens); i < noOfTokens; i++ {
			token, err := cd.DeployLinkTokenContract()
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
		var tokens []*ccip.LinkToken
		for _, token := range ccipModule.BridgeTokens {
			newToken, err := cd.NewLinkTokenContract(common.HexToAddress(token.Address()))
			if err != nil {
				return fmt.Errorf("getting new bridge token contract shouldn't fail %+v", err)
			}
			tokens = append(tokens, newToken)
		}
		ccipModule.BridgeTokens = tokens
	}
	if len(ccipModule.BridgeTokenPools) < noOfTokens {
		// deploy native token pool
		for i := len(ccipModule.BridgeTokenPools); i < noOfTokens; i++ {
			token := ccipModule.BridgeTokens[i]
			btp, err := cd.DeployLockReleaseTokenPoolContract(token.Address())
			if err != nil {
				return fmt.Errorf("deploying bridge Token pool shouldn't fail %+v", err)
			}
			ccipModule.BridgeTokenPools = append(ccipModule.BridgeTokenPools, btp)
		}
	} else {
		var pools []*ccip.LockReleaseTokenPool
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
		ccipModule.TokenPrices = append(ccipModule.TokenPrices, big.NewInt(1))
	}
	if ccipModule.AFN == nil {
		ccipModule.AFN, err = cd.DeployAFNContract()
		if err != nil {
			return fmt.Errorf("deploying AFN shouldn't fail %+v", err)
		}
	} else {
		afn, err := cd.NewAFNContract(ccipModule.AFN.EthAddress)
		if err != nil {
			return fmt.Errorf("getting new AFN contract shouldn't fail %+v", err)
		}
		ccipModule.AFN = afn
	}
	if ccipModule.Router == nil {
		weth9addr, err := cd.DeployWrappedNative()
		if err != nil {
			return fmt.Errorf("deploying wrapped native shouldn't fail %+v", err)
		}

		err = ccipModule.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("waiting for deploying wrapped native shouldn't fail %+v", err)
		}
		ccipModule.WrappedNative = *weth9addr

		ccipModule.Router, err = cd.DeployRouter(ccipModule.WrappedNative)
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
		ccipModule.PriceRegistry, err = cd.DeployPriceRegistry(price_registry.InternalPriceUpdates{
			TokenPriceUpdates: []price_registry.InternalTokenPriceUpdate{},
			DestChainId:       0,
			UsdPerUnitGas:     big.NewInt(0),
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

func DefaultCCIPModule(ctx context.Context, chainClient blockchain.EVMClient, existingDeployment bool) *CCIPCommon {
	grp, _ := errgroup.WithContext(ctx)
	return &CCIPCommon{
		ChainClient: chainClient,
		deploy:      grp,
		RateLimiterConfig: ccip.RateLimiterConfig{
			Rate:     ccip.HundredCoins,
			Capacity: ccip.HundredCoins,
		},
		ExistingDeployment: existingDeployment,
		poolFunds:          testhelpers.Link(10),
	}
}

type SourceCCIPModule struct {
	Common                     *CCIPCommon
	Sender                     common.Address
	TransferAmount             []*big.Int
	DestinationChainId         uint64
	OnRamp                     *ccip.OnRamp
	SrcStartBlock              uint64
	CCIPSendRequestedWatcherMu *sync.Mutex
	CCIPSendRequestedWatcher   map[string]*evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested
	SeqNumToMsgIDMu            *sync.Mutex
	SeqNumToMsgID              map[uint64][32]byte
}

func (sourceCCIP *SourceCCIPModule) LoadContracts(conf *laneconfig.LaneConfig) {
	if conf != nil {
		cfg, ok := conf.SrcContracts[sourceCCIP.DestinationChainId]
		if ok {
			if common.IsHexAddress(cfg.OnRamp) {
				sourceCCIP.OnRamp = &ccip.OnRamp{
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
	err = sourceCCIP.Common.deploy.Wait()
	if err != nil {
		return errors.WithStack(err)
	}

	sourceCCIP.Common.ApproveTokens()
	sourceCCIP.LoadContracts(lane)

	if !sourceCCIP.Common.ExistingDeployment {
		// ensure token is a feeToken
		err = sourceCCIP.Common.PriceRegistry.AddFeeToken(common.HexToAddress(sourceCCIP.Common.FeeToken.Address()))
		if err != nil {
			return fmt.Errorf("addFeeTokens shouldn't fail %+v", err)
		}
		err = sourceCCIP.Common.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("waiting for addFeeTokens shouldn't fail %+v", err)
		}

		// update PriceRegistry
		err = sourceCCIP.Common.PriceRegistry.UpdatePrices(price_registry.InternalPriceUpdates{
			DestChainId:   sourceCCIP.DestinationChainId,
			UsdPerUnitGas: big.NewInt(2000e9), // $2000 per eth * 1gwei = 2000e9
			TokenPriceUpdates: []price_registry.InternalTokenPriceUpdate{
				{
					SourceToken: common.HexToAddress(sourceCCIP.Common.FeeToken.Address()),
					/// USD per full fee token, in base units 1e18.
					/// Example:
					///   * 1 USDC = 1.00 USD per token -> 1e18
					///   * 1 LINK = 5.00 USD per token -> 5e18
					///   * 1 ETH = 2,000 USD per token -> 2_000e18
					UsdPerToken: big.NewInt(5e18),
				},
			},
		})
		if err != nil {
			return fmt.Errorf("feeupdates shouldn't fail %+v", err)
		}
		err = sourceCCIP.Common.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("waiting for feeupdates shouldn't fail %+v", err)
		}
	}

	if sourceCCIP.OnRamp == nil {
		var tokensAndPools []evm_2_evm_onramp.EVM2EVMOnRampTokenAndPool
		var tokens []common.Address
		for i, token := range sourceCCIP.Common.BridgeTokens {
			tokens = append(tokens, token.EthAddress)
			tokensAndPools = append(tokensAndPools, evm_2_evm_onramp.EVM2EVMOnRampTokenAndPool{
				Token: token.EthAddress,
				Pool:  sourceCCIP.Common.BridgeTokenPools[i].EthAddress,
			})
		}
		tokensAndPools = append(tokensAndPools, evm_2_evm_onramp.EVM2EVMOnRampTokenAndPool{
			Token: common.HexToAddress(sourceCCIP.Common.FeeToken.Address()),
			Pool:  sourceCCIP.Common.FeeTokenPool.EthAddress,
		})

		sourceCCIP.SrcStartBlock, err = sourceCCIP.Common.ChainClient.LatestBlockNumber(context.Background())
		if err != nil {
			return fmt.Errorf("getting latest block number shouldn't fail %+v", err)
		}

		sourceCCIP.OnRamp, err = contractDeployer.DeployOnRamp(
			sourceCCIP.Common.ChainClient.GetChainID().Uint64(),
			sourceCCIP.DestinationChainId,
			[]common.Address{},
			tokensAndPools,
			sourceCCIP.Common.AFN.EthAddress,
			sourceCCIP.Common.Router.EthAddress,
			sourceCCIP.Common.PriceRegistry.EthAddress,
			sourceCCIP.Common.RateLimiterConfig,
			[]evm_2_evm_onramp.EVM2EVMOnRampFeeTokenConfigArgs{
				{
					Token:           common.HexToAddress(sourceCCIP.Common.FeeToken.Address()),
					FeeAmount:       big.NewInt(0),
					DestGasOverhead: 0,
					Multiplier:      FeeMultiplier, // the low multiplier is for testing purposes. This keeps accounts from running out of funds
				}},
			sourceCCIP.Common.FeeToken.EthAddress,
		)

		if err != nil {
			return fmt.Errorf("onRamp deployment shouldn't fail %+v", err)
		}

		err = sourceCCIP.Common.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("waiting for onRamp deployment shouldn't fail %+v", err)
		}

		// Set bridge token prices on the onRamp
		err = sourceCCIP.OnRamp.SetTokenPrices(tokens, sourceCCIP.Common.TokenPrices)
		if err != nil {
			return fmt.Errorf("setting prices shouldn't fail %+v", err)
		}

		// update source Router with OnRamp address
		err = sourceCCIP.Common.Router.SetOnRamp(sourceCCIP.DestinationChainId, sourceCCIP.OnRamp.EthAddress)
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
		err = sourceCCIP.Common.FeeTokenPool.SetOnRamp(sourceCCIP.OnRamp.EthAddress)
		if err != nil {
			return fmt.Errorf("setting OnRamp on the fee token pool shouldn't fail %+v", err)
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

func (sourceCCIP *SourceCCIPModule) CollectBalanceRequirements(t *testing.T) []testhelpers.BalanceReq {
	var balancesReq []testhelpers.BalanceReq
	for _, token := range sourceCCIP.Common.BridgeTokens {
		balancesReq = append(balancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("%s-BridgeToken-%s", testhelpers.Sender, token.Address()),
			Addr:   sourceCCIP.Sender,
			Getter: GetterForLinkToken(t, token, sourceCCIP.Sender.Hex()),
		})
	}
	for i, pool := range sourceCCIP.Common.BridgeTokenPools {
		balancesReq = append(balancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("%s-TokenPool-%s", testhelpers.Sender, pool.Address()),
			Addr:   pool.EthAddress,
			Getter: GetterForLinkToken(t, sourceCCIP.Common.BridgeTokens[i], pool.Address()),
		})
	}

	if sourceCCIP.Common.FeeToken.Address() != common.HexToAddress("0x0").String() {
		balancesReq = append(balancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("%s-FeeToken-%s-Address-%s", testhelpers.Sender, sourceCCIP.Common.FeeToken.Address(), sourceCCIP.Sender.Hex()),
			Addr:   sourceCCIP.Sender,
			Getter: GetterForLinkToken(t, sourceCCIP.Common.FeeToken, sourceCCIP.Sender.Hex()),
		})
		balancesReq = append(balancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("%s-Router-%s", testhelpers.Sender, sourceCCIP.Common.Router.Address()),
			Addr:   sourceCCIP.Common.Router.EthAddress,
			Getter: GetterForLinkToken(t, sourceCCIP.Common.FeeToken, sourceCCIP.Common.Router.Address()),
		})
		balancesReq = append(balancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("%s-OnRamp-%s", testhelpers.Sender, sourceCCIP.OnRamp.Address()),
			Addr:   sourceCCIP.OnRamp.EthAddress,
			Getter: GetterForLinkToken(t, sourceCCIP.Common.FeeToken, sourceCCIP.OnRamp.Address()),
		})
		balancesReq = append(balancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("%s-Prices-%s", testhelpers.Sender, sourceCCIP.Common.PriceRegistry.Address()),
			Addr:   sourceCCIP.Common.PriceRegistry.EthAddress,
			Getter: GetterForLinkToken(t, sourceCCIP.Common.FeeToken, sourceCCIP.Common.PriceRegistry.Address()),
		})
	}
	return balancesReq
}

func (sourceCCIP *SourceCCIPModule) BalanceAssertions(t *testing.T, prevBalances map[string]*big.Int, noOfReq int64, totalFee *big.Int) []testhelpers.BalanceAssertion {
	var balAssertions []testhelpers.BalanceAssertion
	for i, token := range sourceCCIP.Common.BridgeTokens {
		name := fmt.Sprintf("%s-BridgeToken-%s", testhelpers.Sender, token.Address())
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     name,
			Address:  sourceCCIP.Sender,
			Getter:   GetterForLinkToken(t, token, sourceCCIP.Sender.Hex()),
			Expected: bigmath.Sub(prevBalances[name], bigmath.Mul(big.NewInt(noOfReq), sourceCCIP.TransferAmount[i])).String(),
		})
	}
	for i, pool := range sourceCCIP.Common.BridgeTokenPools {
		name := fmt.Sprintf("%s-TokenPool-%s", testhelpers.Sender, pool.Address())
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     fmt.Sprintf("%s-TokenPool-%s", testhelpers.Sender, pool.Address()),
			Address:  pool.EthAddress,
			Getter:   GetterForLinkToken(t, sourceCCIP.Common.BridgeTokens[i], pool.Address()),
			Expected: bigmath.Add(prevBalances[name], bigmath.Mul(big.NewInt(noOfReq), sourceCCIP.TransferAmount[i])).String(),
		})
	}

	if sourceCCIP.Common.FeeToken.Address() != common.HexToAddress("0x0").String() {
		name := fmt.Sprintf("%s-FeeToken-%s-Address-%s", testhelpers.Sender, sourceCCIP.Common.FeeToken.Address(), sourceCCIP.Sender.Hex())
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     name,
			Address:  sourceCCIP.Sender,
			Getter:   GetterForLinkToken(t, sourceCCIP.Common.FeeToken, sourceCCIP.Sender.Hex()),
			Expected: bigmath.Sub(prevBalances[name], totalFee).String(),
		})
		name = fmt.Sprintf("%s-Prices-%s", testhelpers.Sender, sourceCCIP.Common.PriceRegistry.Address())
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     name,
			Address:  sourceCCIP.Common.PriceRegistry.EthAddress,
			Getter:   GetterForLinkToken(t, sourceCCIP.Common.FeeToken, sourceCCIP.Common.PriceRegistry.Address()),
			Expected: prevBalances[name].String(),
		})
		name = fmt.Sprintf("%s-Router-%s", testhelpers.Sender, sourceCCIP.Common.Router.Address())
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     fmt.Sprintf("%s-Router-%s", testhelpers.Sender, sourceCCIP.Common.Router.Address()),
			Address:  sourceCCIP.Common.Router.EthAddress,
			Getter:   GetterForLinkToken(t, sourceCCIP.Common.FeeToken, sourceCCIP.Common.Router.Address()),
			Expected: prevBalances[name].String(),
		})
		name = fmt.Sprintf("%s-OnRamp-%s", testhelpers.Sender, sourceCCIP.OnRamp.Address())
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     fmt.Sprintf("%s-OnRamp-%s", testhelpers.Sender, sourceCCIP.OnRamp.Address()),
			Address:  sourceCCIP.OnRamp.EthAddress,
			Getter:   GetterForLinkToken(t, sourceCCIP.Common.FeeToken, sourceCCIP.OnRamp.Address()),
			Expected: bigmath.Add(prevBalances[name], totalFee).String(),
		})
	}
	return balAssertions
}

func (sourceCCIP *SourceCCIPModule) AssertEventCCIPSendRequested(
	lggr zerolog.Logger,
	reqNo int64,
	txHash string,
	timeout time.Duration,
	prevEventAt time.Time,
	reports *testreporters.CCIPLaneStats,
) (*evm_2_evm_onramp.InternalEVM2EVMMessage, time.Time, error) {
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
				hdr, err := sourceCCIP.Common.ChainClient.HeaderByNumber(context.Background(), big.NewInt(int64(sendRequested.Raw.BlockNumber)))
				receivedAt := time.Now()
				if err == nil {
					receivedAt = hdr.Timestamp
				}
				sentMsg := sendRequested.Message
				seqNum := sentMsg.SequenceNumber
				reports.UpdatePhaseStats(reqNo, seqNum, testreporters.CCIPSendRe, receivedAt.Sub(prevEventAt), testreporters.Success)
				return &sentMsg, receivedAt, nil
			}
		case <-ctx.Done():
			reports.UpdatePhaseStats(reqNo, 0, testreporters.CCIPSendRe, time.Since(prevEventAt), testreporters.Failure)
			return nil, time.Now(), fmt.Errorf("CCIPSendRequested event is not found for tx %s", txHash)
		}
	}
}

func (sourceCCIP *SourceCCIPModule) SendRequest(
	receiver common.Address,
	tokenAndAmounts []router.ClientEVMTokenAmount,
	data string,
	feeToken common.Address,
) (common.Hash, time.Duration, *big.Int, error) {
	receiverAddr, err := utils.ABIEncode(`[{"type":"address"}]`, receiver)
	var d time.Duration
	if err != nil {
		return common.Hash{}, d, nil, fmt.Errorf("failed encoding the receiver address: %+v", err)
	}

	extraArgsV1, err := testhelpers.GetEVMExtraArgsV1(big.NewInt(100_000), false)
	if err != nil {
		return common.Hash{}, d, nil, fmt.Errorf("failed encoding the options field: %+v", err)
	}

	// form the message for transfer
	msg := router.ClientEVM2AnyMessage{
		Receiver:     receiverAddr,
		Data:         []byte(data),
		TokenAmounts: tokenAndAmounts,
		FeeToken:     feeToken,
		ExtraArgs:    extraArgsV1,
	}
	log.Info().Interface("ge msg details", msg).Msg("ccip message to be sent")
	fee, err := sourceCCIP.Common.Router.GetFee(sourceCCIP.DestinationChainId, msg)
	if err != nil {
		return common.Hash{}, d, nil, fmt.Errorf("failed getting the fee: %+v", err)
	}
	log.Info().Int64("fee", fee.Int64()).Msg("calculated fee")

	var sendTx *types.Transaction
	timeNow := time.Now()
	// initiate the transfer
	// if the token address is 0x0 it will use Native as fee token and the fee amount should be mentioned in bind.TransactOpts's value
	if feeToken != common.HexToAddress("0x0") {
		sendTx, err = sourceCCIP.Common.Router.CCIPSend(sourceCCIP.DestinationChainId, msg, nil)
		if err != nil {
			return common.Hash{}, time.Since(timeNow), nil, fmt.Errorf("failed initiating the transfer ccip-send: %+v", err)
		}
	} else {
		sendTx, err = sourceCCIP.Common.Router.CCIPSend(sourceCCIP.DestinationChainId, msg, fee)
		if err != nil {
			return common.Hash{}, time.Since(timeNow), nil, fmt.Errorf("failed initiating the transfer ccip-send: %+v", err)
		}
	}

	log.Info().Str("Send token transaction", sendTx.Hash().String()).Msg("Sending token")
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	if err != nil {
		return common.Hash{}, time.Since(timeNow), nil, fmt.Errorf("failed waiting for events: %+v", err)
	}
	return sendTx.Hash(), time.Since(timeNow), fee, nil
}

func DefaultSourceCCIPModule(chainClient blockchain.EVMClient, destChain uint64, transferAmount []*big.Int, ccipCommon *CCIPCommon) (*SourceCCIPModule, error) {
	sourceCCIP := &SourceCCIPModule{
		Common:                     ccipCommon,
		TransferAmount:             transferAmount,
		DestinationChainId:         destChain,
		Sender:                     common.HexToAddress(chainClient.GetDefaultWallet().Address()),
		CCIPSendRequestedWatcher:   make(map[string]*evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested),
		CCIPSendRequestedWatcherMu: &sync.Mutex{},
	}
	var err error
	sourceCCIP.Common.Deployer, err = ccip.NewCCIPContractsDeployer(chainClient)
	if err != nil {
		return nil, fmt.Errorf("contract deployer should be created successfully %+v", err)
	}

	return sourceCCIP, nil
}

type DestCCIPModule struct {
	Common                  *CCIPCommon
	SourceChainId           uint64
	CommitStore             *ccip.CommitStore
	ReceiverDapp            *ccip.ReceiverDapp
	OffRamp                 *ccip.OffRamp
	WrappedNative           common.Address
	ReportAcceptedWatcherMu *sync.Mutex
	ReportAcceptedWatcher   map[uint64]*types.Log
	ExecStateChangedMu      *sync.Mutex
	ExecStateChangedWatcher map[uint64]*evm_2_evm_offramp.EVM2EVMOffRampExecutionStateChanged
	NextSeqNumToCommit      *atomic.Uint64
}

func (destCCIP *DestCCIPModule) LoadContracts(conf *laneconfig.LaneConfig) {
	if conf != nil {
		cfg, ok := conf.DestContracts[destCCIP.SourceChainId]
		if ok {
			if common.IsHexAddress(cfg.OffRamp) {
				destCCIP.OffRamp = &ccip.OffRamp{
					EthAddress: common.HexToAddress(cfg.OffRamp),
				}
			}
			if common.IsHexAddress(cfg.CommitStore) {
				destCCIP.CommitStore = &ccip.CommitStore{
					EthAddress: common.HexToAddress(cfg.CommitStore),
				}
			}
			if common.IsHexAddress(cfg.ReceiverDapp) {
				destCCIP.ReceiverDapp = &ccip.ReceiverDapp{
					EthAddress: common.HexToAddress(cfg.ReceiverDapp),
				}
			}
		}
	}
}

// DeployContracts deploys all CCIP contracts specific to the destination chain
func (destCCIP *DestCCIPModule) DeployContracts(
	sourceCCIP SourceCCIPModule,
	wg *sync.WaitGroup,
	lane *laneconfig.LaneConfig,
) error {
	var err error
	contractDeployer := destCCIP.Common.Deployer
	log.Info().Msg("Deploying destination chain specific contracts")

	err = destCCIP.Common.deploy.Wait()
	if err != nil {
		return errors.WithStack(err)
	}

	destCCIP.LoadContracts(lane)
	if !destCCIP.Common.ExistingDeployment {
		// ensure token is a feeToken
		err = destCCIP.Common.PriceRegistry.AddFeeToken(common.HexToAddress(destCCIP.Common.FeeToken.Address()))
		if err != nil {
			return fmt.Errorf("addFeeTokens shouldn't fail %+v", err)
		}
		err = destCCIP.Common.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("waiting for addFeeTokens shouldn't fail %+v", err)
		}

		// update PriceRegistry
		err = destCCIP.Common.PriceRegistry.UpdatePrices(price_registry.InternalPriceUpdates{
			DestChainId:   destCCIP.SourceChainId,
			UsdPerUnitGas: big.NewInt(2000e9), // $2000 per eth * 1gwei = 2000e9
			TokenPriceUpdates: []price_registry.InternalTokenPriceUpdate{
				{
					SourceToken: common.HexToAddress(destCCIP.Common.FeeToken.Address()),
					/// USD per full fee token, in base units 1e18.
					/// Example:
					///   * 1 USDC = 1.00 USD per token -> 1e18
					///   * 1 LINK = 5.00 USD per token -> 5e18
					///   * 1 ETH = 2,000 USD per token -> 2_000e18
					UsdPerToken: big.NewInt(5e18),
				},
			},
		})
		if err != nil {
			return fmt.Errorf("priceupdates shouldn't fail %+v", err)
		}

		err = destCCIP.Common.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("waiting for priceupdates shouldn't fail %+v", err)
		}
	}

	if destCCIP.CommitStore == nil {
		// commitStore responsible for validating the transfer message
		destCCIP.CommitStore, err = contractDeployer.DeployCommitStore(
			destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID().Uint64(),
			destCCIP.Common.AFN.EthAddress, sourceCCIP.OnRamp.EthAddress,
			destCCIP.Common.PriceRegistry.EthAddress,
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

	// notify that all common contracts and commit store has been deployed so that the set-up in reverse lane can be triggered.
	if wg != nil {
		wg.Done()
	}

	var sourceTokens, destTokens, pools []common.Address

	for _, token := range sourceCCIP.Common.BridgeTokens {
		sourceTokens = append(sourceTokens, common.HexToAddress(token.Address()))
	}
	sourceTokens = append(sourceTokens, common.HexToAddress(sourceCCIP.Common.FeeToken.Address()))

	for i, token := range destCCIP.Common.BridgeTokens {
		destTokens = append(destTokens, common.HexToAddress(token.Address()))
		pool := destCCIP.Common.BridgeTokenPools[i]
		pools = append(pools, pool.EthAddress)
		if !destCCIP.Common.ExistingDeployment {
			err = pool.AddLiquidity(token, destCCIP.Common.poolFunds)
			if err != nil {
				return fmt.Errorf("adding liquidity token to dest pool shouldn't fail %+v", err)
			}
			err = destCCIP.Common.ChainClient.WaitForEvents()
			if err != nil {
				return fmt.Errorf("waiting for adding liquidity token to dest pool shouldn't fail %+v", err)
			}
		}
	}
	// add the fee token and fee token price for dest
	destTokens = append(destTokens, common.HexToAddress(destCCIP.Common.FeeToken.Address()))
	destCCIP.Common.TokenPrices = append(destCCIP.Common.TokenPrices, big.NewInt(1))

	pools = append(pools, destCCIP.Common.FeeTokenPool.EthAddress)

	if destCCIP.OffRamp == nil {
		destCCIP.OffRamp, err = contractDeployer.DeployOffRamp(destCCIP.SourceChainId, sourceCCIP.DestinationChainId,
			destCCIP.CommitStore.EthAddress, sourceCCIP.OnRamp.EthAddress,
			destCCIP.Common.AFN.EthAddress, destCCIP.Common.Router.EthAddress,
			sourceTokens, pools, destCCIP.Common.RateLimiterConfig)
		if err != nil {
			return fmt.Errorf("deploying offramp shouldn't fail %+v", err)
		}
		err = destCCIP.Common.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("waiting for offramp deployment shouldn't fail %+v", err)
		}

		// apply offramp updates
		_, err = destCCIP.Common.Router.AddOffRamp(destCCIP.OffRamp.EthAddress, destCCIP.SourceChainId)
		if err != nil {
			return fmt.Errorf("setting offramp as fee updater shouldn't fail %+v", err)
		}
		err = destCCIP.Common.ChainClient.WaitForEvents()
		if err != nil {
			return fmt.Errorf("waiting for events on destination contract shouldn't fail %+v", err)
		}

		err = destCCIP.OffRamp.SetTokenPrices(destTokens, destCCIP.Common.TokenPrices)
		if err != nil {
			return fmt.Errorf("setting offramp token prices shouldn't fail %+v", err)
		}

		// update pools with offRamp id
		for _, pool := range destCCIP.Common.BridgeTokenPools {
			err = pool.SetOffRamp(destCCIP.OffRamp.EthAddress)
			if err != nil {
				return fmt.Errorf("setting offramp on the bridge token pool shouldn't fail %+v", err)
			}
		}

		err = destCCIP.Common.FeeTokenPool.SetOffRamp(destCCIP.OffRamp.EthAddress)
		if err != nil {
			return fmt.Errorf("setting offramp on the fee token pool shouldn't fail %+v", err)
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

func (destCCIP *DestCCIPModule) CollectBalanceRequirements(t *testing.T) []testhelpers.BalanceReq {
	var destBalancesReq []testhelpers.BalanceReq
	for _, token := range destCCIP.Common.BridgeTokens {
		destBalancesReq = append(destBalancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("%s-BridgeToken-%s", testhelpers.Receiver, token.Address()),
			Addr:   destCCIP.ReceiverDapp.EthAddress,
			Getter: GetterForLinkToken(t, token, destCCIP.ReceiverDapp.Address()),
		})
	}
	for i, pool := range destCCIP.Common.BridgeTokenPools {
		destBalancesReq = append(destBalancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("%s-TokenPool-%s", testhelpers.Receiver, pool.Address()),
			Addr:   pool.EthAddress,
			Getter: GetterForLinkToken(t, destCCIP.Common.BridgeTokens[i], pool.Address()),
		})
	}
	if destCCIP.Common.FeeToken.Address() != common.HexToAddress("0x0").String() {
		destBalancesReq = append(destBalancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("%s-FeeToken-%s-Address-%s", testhelpers.Receiver, destCCIP.Common.FeeToken.Address(), destCCIP.ReceiverDapp.Address()),
			Addr:   destCCIP.ReceiverDapp.EthAddress,
			Getter: GetterForLinkToken(t, destCCIP.Common.FeeToken, destCCIP.ReceiverDapp.Address()),
		})
		destBalancesReq = append(destBalancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("%s-OffRamp-%s", testhelpers.Receiver, destCCIP.OffRamp.Address()),
			Addr:   destCCIP.OffRamp.EthAddress,
			Getter: GetterForLinkToken(t, destCCIP.Common.FeeToken, destCCIP.OffRamp.Address()),
		})
		destBalancesReq = append(destBalancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("%s-FeeTokenPool-%s", testhelpers.Receiver, destCCIP.Common.FeeTokenPool.Address()),
			Addr:   destCCIP.Common.FeeTokenPool.EthAddress,
			Getter: GetterForLinkToken(t, destCCIP.Common.FeeToken, destCCIP.Common.FeeTokenPool.Address()),
		})
	}
	return destBalancesReq
}

func (destCCIP *DestCCIPModule) BalanceAssertions(
	t *testing.T,
	prevBalances map[string]*big.Int,
	transferAmount []*big.Int,
	noOfReq int64,
) []testhelpers.BalanceAssertion {
	var balAssertions []testhelpers.BalanceAssertion
	for i, token := range destCCIP.Common.BridgeTokens {
		name := fmt.Sprintf("%s-BridgeToken-%s", testhelpers.Receiver, token.Address())
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     name,
			Address:  destCCIP.ReceiverDapp.EthAddress,
			Getter:   GetterForLinkToken(t, token, destCCIP.ReceiverDapp.Address()),
			Expected: bigmath.Add(prevBalances[name], bigmath.Mul(big.NewInt(noOfReq), transferAmount[i])).String(),
		})
	}
	for i, pool := range destCCIP.Common.BridgeTokenPools {
		name := fmt.Sprintf("%s-TokenPool-%s", testhelpers.Receiver, pool.Address())
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     fmt.Sprintf("%s-TokenPool-%s", testhelpers.Receiver, pool.Address()),
			Address:  pool.EthAddress,
			Getter:   GetterForLinkToken(t, destCCIP.Common.BridgeTokens[i], pool.Address()),
			Expected: bigmath.Sub(prevBalances[name], bigmath.Mul(big.NewInt(noOfReq), transferAmount[i])).String(),
		})
	}
	if destCCIP.Common.FeeToken.Address() != common.HexToAddress("0x0").String() {
		name := fmt.Sprintf("%s-OffRamp-%s", testhelpers.Receiver, destCCIP.OffRamp.Address())
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     name,
			Address:  destCCIP.OffRamp.EthAddress,
			Getter:   GetterForLinkToken(t, destCCIP.Common.FeeToken, destCCIP.OffRamp.Address()),
			Expected: prevBalances[name].String(),
		})
		name = fmt.Sprintf("%s-FeeTokenPool-%s", testhelpers.Receiver, destCCIP.Common.FeeTokenPool.Address())
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     name,
			Address:  destCCIP.Common.FeeTokenPool.EthAddress,
			Getter:   GetterForLinkToken(t, destCCIP.Common.FeeToken, destCCIP.Common.FeeTokenPool.Address()),
			Expected: prevBalances[name].String(),
		})

		name = fmt.Sprintf("%s-FeeToken-%s-Address-%s", testhelpers.Receiver, destCCIP.Common.FeeToken.Address(), destCCIP.ReceiverDapp.Address())
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     name,
			Address:  destCCIP.ReceiverDapp.EthAddress,
			Getter:   GetterForLinkToken(t, destCCIP.Common.FeeToken, destCCIP.ReceiverDapp.Address()),
			Expected: prevBalances[name].String(),
		})
	}
	return balAssertions
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
				hdr, err := destCCIP.Common.ChainClient.HeaderByNumber(
					context.Background(), big.NewInt(int64(vLogs.BlockNumber)))
				if err == nil {
					receivedAt = hdr.Timestamp
				}
				if e.State == ccipPlugin.MessageStateSuccess {
					reports.UpdatePhaseStats(reqNo, seqNum, testreporters.ExecStateChanged, receivedAt.Sub(timeNow), testreporters.Success)
					return nil
				} else {
					reports.UpdatePhaseStats(reqNo, seqNum, testreporters.ExecStateChanged, time.Since(timeNow), testreporters.Failure)
					return fmt.Errorf("ExecutionStateChanged event state changed to %i for seq num %v for lane %d-->%d",
						e.State, seqNum, destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID())
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
) (time.Time, error) {
	lggr.Info().Int64("seqNum", int64(seqNum)).Msg("Waiting for ReportAccepted event")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			destCCIP.ReportAcceptedWatcherMu.Lock()
			vLogs, ok := destCCIP.ReportAcceptedWatcher[seqNum]
			destCCIP.ReportAcceptedWatcherMu.Unlock()
			receivedAt := time.Now().UTC()
			if ok && vLogs != nil {
				hdr, err := destCCIP.Common.ChainClient.HeaderByNumber(
					context.Background(), big.NewInt(int64(vLogs.BlockNumber)))
				if err == nil {
					receivedAt = hdr.Timestamp
				}

				reports.UpdatePhaseStats(reqNo, seqNum, testreporters.Commit, receivedAt.Sub(prevEventAt), testreporters.Success)
				return receivedAt, nil
			}
		case <-ctx.Done():
			reports.UpdatePhaseStats(reqNo, seqNum, testreporters.Commit, time.Since(prevEventAt), testreporters.Failure)
			return time.Now().UTC(), fmt.Errorf("ReportAccepted is not found for seq num %d lane %d-->%d",
				seqNum, destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID())
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

func DefaultDestinationCCIPModule(chainClient blockchain.EVMClient, sourceChain uint64, ccipCommon *CCIPCommon) (*DestCCIPModule, error) {
	destCCIP := &DestCCIPModule{
		Common:                  ccipCommon,
		SourceChainId:           sourceChain,
		ReportAcceptedWatcherMu: &sync.Mutex{},
		ReportAcceptedWatcher:   make(map[uint64]*types.Log),
		ExecStateChangedMu:      &sync.Mutex{},
		ExecStateChangedWatcher: make(map[uint64]*evm_2_evm_offramp.EVM2EVMOffRampExecutionStateChanged),
		NextSeqNumToCommit:      atomic.NewUint64(1),
	}

	var err error
	destCCIP.Common.Deployer, err = ccip.NewCCIPContractsDeployer(chainClient)
	if err != nil {
		return nil, err
	}
	return destCCIP, nil
}

type CCIPRequest struct {
	reqNo                   int64
	txHash                  string
	txConfirmationTimestamp time.Time
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
	SourceBalances          map[string]*big.Int
	DestBalances            map[string]*big.Int
	StartBlockOnSource      uint64
	StartBlockOnDestination uint64
	SentReqs                map[int64]CCIPRequest
	TotalFee                *big.Int // total fee for all the requests. Used for balance validation.
	ValidationTimeout       time.Duration
	Context                 context.Context
	CommonContractsWg       *sync.WaitGroup
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
		FeeTokenPool:     lane.Source.Common.FeeTokenPool.Address(),
		BridgeTokens:     btAddresses,
		BridgeTokenPools: btpAddresses,
		AFN:              lane.Source.Common.AFN.Address(),
		Router:           lane.Source.Common.Router.Address(),
		PriceRegistry:    lane.Source.Common.PriceRegistry.Address(),
		WrappedNative:    lane.Source.Common.WrappedNative.Hex(),
	}
	if lane.SrcNetworkLaneCfg.SrcContracts == nil {
		lane.SrcNetworkLaneCfg.SrcContracts = map[uint64]laneconfig.SourceContracts{
			lane.Source.DestinationChainId: {
				OnRamp:     lane.Source.OnRamp.Address(),
				DepolyedAt: lane.Source.SrcStartBlock,
			},
		}
	} else {
		lane.SrcNetworkLaneCfg.SrcContracts[lane.Source.DestinationChainId] = laneconfig.SourceContracts{
			OnRamp:     lane.Source.OnRamp.Address(),
			DepolyedAt: lane.Source.SrcStartBlock,
		}
	}
	btAddresses, btpAddresses = []string{}, []string{}
	for i, bt := range lane.Dest.Common.BridgeTokens {
		btAddresses = append(btAddresses, bt.Address())
		btpAddresses = append(btpAddresses, lane.Dest.Common.BridgeTokenPools[i].Address())
	}
	lane.DstNetworkLaneCfg.CommonContracts = laneconfig.CommonContracts{
		FeeToken:         lane.Dest.Common.FeeToken.Address(),
		FeeTokenPool:     lane.Dest.Common.FeeTokenPool.Address(),
		BridgeTokens:     btAddresses,
		BridgeTokenPools: btpAddresses,
		AFN:              lane.Dest.Common.AFN.Address(),
		Router:           lane.Dest.Common.Router.Address(),
		PriceRegistry:    lane.Dest.Common.PriceRegistry.Address(),
		WrappedNative:    lane.Dest.Common.WrappedNative.Hex(),
	}
	if lane.DstNetworkLaneCfg.DestContracts == nil {
		lane.DstNetworkLaneCfg.DestContracts = map[uint64]laneconfig.DestContracts{
			lane.Dest.SourceChainId: {
				OffRamp:      lane.Dest.OffRamp.Address(),
				CommitStore:  lane.Dest.CommitStore.Address(),
				ReceiverDapp: lane.Dest.ReceiverDapp.Address(),
			},
		}
	} else {
		lane.DstNetworkLaneCfg.DestContracts[lane.Dest.SourceChainId] = laneconfig.DestContracts{
			OffRamp:      lane.Dest.OffRamp.Address(),
			CommitStore:  lane.Dest.CommitStore.Address(),
			ReceiverDapp: lane.Dest.ReceiverDapp.Address(),
		}
	}
}

func (lane *CCIPLane) RecordStateBeforeTransfer() {
	var err error
	// collect the balance assert.ment to verify balances after transfer
	lane.SourceBalances, err = testhelpers.GetBalances(lane.Source.CollectBalanceRequirements(lane.Test))
	require.NoError(lane.Test, err, "fetching source balance")
	lane.DestBalances, err = testhelpers.GetBalances(lane.Dest.CollectBalanceRequirements(lane.Test))
	require.NoError(lane.Test, err, "fetching dest balance")

	// save the current block numbers to use in various filter log requests
	lane.StartBlockOnSource, err = lane.Source.Common.ChainClient.LatestBlockNumber(context.Background())
	require.NoError(lane.Test, err, "Getting current block should be successful in source chain")
	lane.StartBlockOnDestination, err = lane.Dest.Common.ChainClient.LatestBlockNumber(context.Background())
	require.NoError(lane.Test, err, "Getting current block should be successful in dest chain")
	lane.TotalFee = big.NewInt(0)
	lane.NumberOfReq = 0
	lane.SentReqs = make(map[int64]CCIPRequest)
}

func (lane *CCIPLane) SendRequests(noOfRequests int) (map[int64]CCIPRequest, error) {
	var tokenAndAmounts []router.ClientEVMTokenAmount
	for i, token := range lane.Source.Common.BridgeTokens {
		tokenAndAmounts = append(tokenAndAmounts, router.ClientEVMTokenAmount{
			Token: common.HexToAddress(token.Address()), Amount: lane.Source.TransferAmount[i],
		})
	}

	err := lane.Source.Common.ChainClient.WaitForEvents()
	if err != nil {
		return nil, fmt.Errorf("could not wait for events: %+v", err)
	}

	txMap := make(map[int64]CCIPRequest)
	for i := 1; i <= noOfRequests; i++ {
		txHash, txConfirmationDur, fee, err := lane.Source.SendRequest(
			lane.Dest.ReceiverDapp.EthAddress,
			tokenAndAmounts,
			fmt.Sprintf("msg %d", i),
			common.HexToAddress(lane.Source.Common.FeeToken.Address()),
		)
		if err != nil {
			lane.Reports.UpdatePhaseStats(int64(lane.NumberOfReq+i), 0,
				testreporters.TX, txConfirmationDur, testreporters.Failure)
			return nil, fmt.Errorf("could not send request: %+v", err)
		}
		txConfirmationTimestamp := time.Now().UTC()
		rcpt, err := lane.Source.Common.ChainClient.GetTxReceipt(txHash)
		if err == nil {
			hdr, err := lane.Source.Common.ChainClient.HeaderByNumber(context.Background(), rcpt.BlockNumber)
			if err == nil {
				txConfirmationTimestamp = hdr.Timestamp
			}
		}
		txMap[int64(lane.NumberOfReq+i)] = CCIPRequest{
			txHash:                  txHash.Hex(),
			txConfirmationTimestamp: txConfirmationTimestamp,
		}
		lane.SentReqs[int64(lane.NumberOfReq+i)] = txMap[int64(lane.NumberOfReq+i)]
		lane.Reports.UpdatePhaseStats(int64(lane.NumberOfReq+i), 0,
			testreporters.TX, txConfirmationDur, testreporters.Success)
		lane.TotalFee = bigmath.Add(lane.TotalFee, fee)
	}
	lane.NumberOfReq += noOfRequests

	return txMap, nil
}

func (lane *CCIPLane) ValidateRequests() {
	for i, tx := range lane.SentReqs {
		require.NoError(lane.Test, lane.ValidateRequestByTxHash(tx.txHash, tx.txConfirmationTimestamp, i),
			"validating request events by tx hash")
	}
	// Asserting balances reliably work only for simulated private chains. The testnet contract balances might get updated by other transactions
	// verify the fee amount is deducted from sender, added to receiver token balances and
	// unused fee is returned to receiver fee token account
	AssertBalances(lane.Test, lane.Source.BalanceAssertions(lane.Test, lane.SourceBalances, int64(lane.NumberOfReq), lane.TotalFee))
	AssertBalances(lane.Test, lane.Dest.BalanceAssertions(lane.Test, lane.DestBalances, lane.Source.TransferAmount, int64(lane.NumberOfReq)))
}

func (lane *CCIPLane) ValidateRequestByTxHash(txHash string, txConfirmattion time.Time, reqNo int64) error {
	// Verify if
	// - CCIPSendRequested Event log generated,
	// - NextSeqNumber from commitStore got increased
	msg, receivedAt, err := lane.Source.AssertEventCCIPSendRequested(
		lane.Logger, reqNo, txHash, lane.ValidationTimeout, txConfirmattion, lane.Reports)
	if err != nil || msg == nil {
		return fmt.Errorf("could not validate CCIPSendRequested event: %+v", err)
	}
	seqNumber := msg.SequenceNumber

	err = lane.Dest.AssertSeqNumberExecuted(lane.Logger, reqNo, seqNumber, lane.ValidationTimeout, receivedAt, lane.Reports)
	if err != nil {
		return fmt.Errorf("could not validate seq number increase at commit store: %+v", err)
	}

	// Verify whether commitStore has accepted the report
	receivedAt, err = lane.Dest.AssertEventReportAccepted(lane.Logger, reqNo, seqNumber, lane.ValidationTimeout, receivedAt, lane.Reports)
	if err != nil {
		return fmt.Errorf("could not validate ReportAccepted event: %+v", err)
	}

	// Verify whether the execution state is changed and the transfer is successful
	err = lane.Dest.AssertEventExecutionStateChanged(lane.Logger, reqNo, seqNumber, lane.ValidationTimeout, receivedAt, lane.Reports)
	if err != nil {
		return fmt.Errorf("could not validate ExecutionStateChanged event: %+v", err)
	}
	return nil
}

func (lane *CCIPLane) SoakRun(interval, duration time.Duration) (int, int) {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	ticker := time.NewTicker(interval)
	numOfReq := 0
	reqSuccess := 0
	wg := &sync.WaitGroup{}
	timeout := false
	lane.RecordStateBeforeTransfer()
	for {
		select {
		case <-ticker.C:
			if timeout {
				break
			}
			numOfReq++
			lane.Logger.Info().
				Int("Req No", numOfReq).
				Msg("CCIP transfer")
			txs, err := lane.SendRequests(1)
			if err == nil {
				for reqNo, req := range txs {
					wg.Add(1)
					go func(txHash string, reqNo int64, txConfirmationTime time.Time) {
						defer wg.Done()
						if lane.ValidateRequestByTxHash(txHash, txConfirmationTime, reqNo) == nil {
							reqSuccess++
						}
					}(req.txHash, reqNo, req.txConfirmationTimestamp.In(time.Now().Location()))
				}
			}
		case <-ctx.Done():
			lane.Logger.Warn().
				Msg("Soak Test duration completed. Completing validation for triggered requests")
			timeout = true
			wg.Wait()
			return numOfReq, reqSuccess
		}
	}
}

func (lane *CCIPLane) StartEventWatchers() error {
	sendReqEvent := make(chan *evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested)
	sub, err := lane.Source.OnRamp.Instance.WatchCCIPSendRequested(nil, sendReqEvent)
	if err != nil {
		return err
	}
	lane.Subscriptions = append(lane.Subscriptions, sub)
	go func() {
		for {
			e := <-sendReqEvent
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
				lane.Dest.ReportAcceptedWatcher[i] = &e.Raw
			}
			lane.Dest.ReportAcceptedWatcherMu.Unlock()
		}
	}()

	execStateChangedEvent := make(chan *evm_2_evm_offramp.EVM2EVMOffRampExecutionStateChanged)
	sub, err = lane.Dest.OffRamp.Instance.WatchExecutionStateChanged(nil, execStateChangedEvent, nil, nil)
	if err != nil {
		return err
	}

	lane.Subscriptions = append(lane.Subscriptions, sub)

	go func() {
		for {
			e := <-execStateChangedEvent
			lane.Dest.ExecStateChangedMu.Lock()
			lane.Dest.ExecStateChangedWatcher[e.SequenceNumber] = e
			lane.Dest.ExecStateChangedMu.Unlock()
		}
	}()
	return nil
}

func (lane *CCIPLane) CleanUp() {
	lane.Logger.Info().Msg("Cleaning up lane")
	for _, sub := range lane.Subscriptions {
		sub.Unsubscribe()
	}
	require.NoError(lane.Test, lane.DestChain.Close())
	require.NoError(lane.Test, lane.SourceChain.Close())
}

// DeployNewCCIPLane sets up a lane and initiates lane.Source and lane.Destination
// If configureCLNodes is true it sets up jobs and contract config for the lane
func (lane *CCIPLane) DeployNewCCIPLane(
	numOfCommitNodes int,
	commitAndExecOnSameDON bool,
	sourceCommon *CCIPCommon,
	destCommon *CCIPCommon,
	transferAmounts []*big.Int,
	newBootstrap bool,
	configureCLNodes bool,
	existingDeployment bool,
) error {
	var err error
	env := lane.TestEnv
	sourceChainClient := lane.SourceChain
	destChainClient := lane.DestChain

	if sourceCommon == nil {
		sourceCommon = DefaultCCIPModule(lane.Context, sourceChainClient, existingDeployment)
	}

	if destCommon == nil {
		destCommon = DefaultCCIPModule(lane.Context, destChainClient, existingDeployment)
	}

	lane.Source, err = DefaultSourceCCIPModule(sourceChainClient, destChainClient.GetChainID().Uint64(), transferAmounts, sourceCommon)
	if err != nil {
		return errors.WithStack(err)
	}
	lane.Dest, err = DefaultDestinationCCIPModule(destChainClient, sourceChainClient.GetChainID().Uint64(), destCommon)
	if err != nil {
		return errors.WithStack(err)
	}

	srcConf := lane.SrcNetworkLaneCfg

	destConf := lane.DstNetworkLaneCfg

	// deploy all common contracts in parallel
	lane.Source.Common.deploy.Go(func() error {
		return lane.Source.Common.DeployContracts(len(lane.Source.TransferAmount), srcConf)
	})

	lane.Dest.Common.deploy.Go(func() error {
		return lane.Dest.Common.DeployContracts(len(lane.Source.TransferAmount), destConf)
	})

	// deploy all source contracts
	err = lane.Source.DeployContracts(srcConf)
	if err != nil {
		return errors.WithStack(err)
	}
	// deploy all destination contracts
	err = lane.Dest.DeployContracts(*lane.Source, lane.CommonContractsWg, destConf)
	if err != nil {
		return errors.WithStack(err)
	}
	lane.UpdateLaneConfig()

	// if lane is being set up for already configured CL nodes and contracts
	// no further action is necessary
	if !configureCLNodes {
		return nil
	}

	if env == nil {
		return errors.WithStack(errors.New("test environment not set"))
	}
	// wait for the CL nodes to be ready before moving ahead with job creation
	err = env.CLNodeWithKeyReady.Wait()
	if err != nil {
		return errors.WithStack(err)
	}
	clNodesWithKeys := env.CLNodesWithKeys
	mockServer := env.MockServer
	// set up ocr2 jobs
	tokenUSDMap := make(map[string]string)
	for _, token := range lane.Dest.Common.BridgeTokens {
		tokenUSDMap[token.Address()] = LinkToUSD.String()
	}
	clNodes, exists := clNodesWithKeys[lane.Dest.Common.ChainClient.GetChainID().String()]
	if !exists {
		return fmt.Errorf("could not find CL nodes for %s", lane.Dest.Common.ChainClient.GetChainID().String())
	}

	tokenUSDMap[lane.Dest.Common.FeeToken.Address()] = LinkToUSD.String()
	tokenUSDMap[lane.Source.Common.WrappedNative.Hex()] = WrappedNativeToUSD.String()

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
		bootstrapExec = clNodes[1] // for a set-up of different commit and execution nodes second node is the bootstrapper for execution nodes
		commitNodes = clNodes[2 : 2+numOfCommitNodes]
		execNodes = clNodes[2+numOfCommitNodes:]
		env.commitNodeStartIndex = 2
		env.execNodeStartIndex = 7
		env.numOfCommitNodes = len(commitNodes)
		env.numOfExecNodes = len(execNodes)
	}

	// set up ocr2 config
	err = SetOCR2Configs(commitNodes, execNodes, *lane.Dest)
	if err != nil {
		return errors.WithStack(err)
	}
	err = CreateOCRJobsForCCIP(
		lane.Context,
		bootstrapCommit, bootstrapExec, commitNodes, execNodes,
		lane.Source.OnRamp.EthAddress,
		lane.Dest.CommitStore.EthAddress,
		lane.Dest.OffRamp.EthAddress,
		sourceChainClient, destChainClient,
		lane.Source.SrcStartBlock,
		tokenUSDMap,
		mockServer, newBootstrap,
	)
	if err != nil {
		return errors.WithStack(err)
	}
	// start event watchers
	err = lane.StartEventWatchers()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// SetOCR2Configs sets the oracle config in ocr2 contracts
// nil value in execNodes denotes commit and execution jobs are to be set up in same DON
func SetOCR2Configs(commitNodes, execNodes []*client.CLNodesWithKeys, destCCIP DestCCIPModule) error {
	signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, err :=
		ccip.NewOffChainAggregatorV2Config(commitNodes, ccipPlugin.CommitOffchainConfig{
			SourceIncomingConfirmations: 1,
			DestIncomingConfirmations:   1,
			MaxGasPrice:                 200e9,
		}, ccipPlugin.CommitOnchainConfig{
			PriceRegistry: destCCIP.Common.PriceRegistry.EthAddress,
			Afn:           destCCIP.Common.AFN.EthAddress,
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
		signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, err =
			ccip.NewOffChainAggregatorV2Config(nodes, ccipPlugin.ExecOffchainConfig{
				SourceIncomingConfirmations: 1,
				DestIncomingConfirmations:   1,
				BatchGasLimit:               5_000_000,
				RelativeBoostPerWaitHour:    0.7,
				MaxGasPrice:                 200e9,
			}, ccipPlugin.ExecOnchainConfig{
				// FIXME Replace with real values
				PermissionLessExecutionThresholdSeconds: 60,
				Router:                                  destCCIP.Common.Router.EthAddress,
				Afn:                                     destCCIP.Common.AFN.EthAddress,
				MaxTokensLength:                         5,
				MaxDataSize:                             1e5,
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

// CreateOCRJobsForCCIP bootstraps the first node and to the other nodes sends ocr jobs that
// sets up ccip-commit and ccip-execution plugin
// nil value in bootstrapExec and execNodes denotes commit and execution jobs are to be set up in same DON
func CreateOCRJobsForCCIP(
	ctx context.Context,
	bootstrapCommit *client.CLNodesWithKeys,
	bootstrapExec *client.CLNodesWithKeys,
	commitNodes, execNodes []*client.CLNodesWithKeys,
	onRamp, commitStore, offRamp common.Address,
	sourceChainClient, destChainClient blockchain.EVMClient,
	srcStartBlock uint64,
	tokenUSDMap map[string]string,
	mockServer *ctfClient.MockserverClient,
	newBootStrap bool,
) error {
	bootstrapCommitP2PId := bootstrapCommit.KeysBundle.P2PKeys.Data[0].Attributes.PeerID
	var bootstrapExecP2PId string
	if bootstrapExec == nil {
		bootstrapExec = bootstrapCommit
		bootstrapExecP2PId = bootstrapCommitP2PId
	} else {
		bootstrapExecP2PId = bootstrapExec.KeysBundle.P2PKeys.Data[0].Attributes.PeerID
	}
	p2pBootstrappersCommit := &client.P2PData{
		RemoteIP: bootstrapCommit.Node.RemoteIP(),
		PeerID:   bootstrapCommitP2PId,
	}

	p2pBootstrappersExec := &client.P2PData{
		RemoteIP: bootstrapExec.Node.RemoteIP(),
		PeerID:   bootstrapExecP2PId,
	}
	// save the current block numbers. If there is a delay between job start up and ocr config set up, the jobs will
	// replay the log polling from these mentioned block number. The dest block number should ideally be the block number on which
	// contract config is set and the source block number should be the one on which the ccip send request is performed.
	// Here for simplicity we are just taking the current block number just before the job is created.
	currentBlockOnDest, err := destChainClient.LatestBlockNumber(context.Background())
	if err != nil {
		return fmt.Errorf("getting current block should be successful in destination chain %+v", err)
	}

	jobParams := testhelpers.CCIPJobSpecParams{
		OnRamp:             onRamp,
		OffRamp:            offRamp,
		CommitStore:        commitStore,
		SourceChainName:    sourceChainClient.GetNetworkName(),
		DestChainName:      destChainClient.GetNetworkName(),
		SourceChainId:      sourceChainClient.GetChainID().Uint64(),
		DestChainId:        destChainClient.GetChainID().Uint64(),
		PollPeriod:         time.Second,
		SourceStartBlock:   srcStartBlock,
		DestStartBlock:     currentBlockOnDest,
		RelayInflight:      InflightExpiry,
		ExecInflight:       InflightExpiry,
		RootSnooze:         RootSnoozeTime,
		P2PV2Bootstrappers: []string{p2pBootstrappersCommit.P2PV2Bootstrapper()},
	}

	if destChainClient.NetworkSimulated() {
		jobParams.RootSnooze = RootSnoozeTimeSimulated
		jobParams.ExecInflight = InflightExpirySimulated
		jobParams.RelayInflight = InflightExpirySimulated
	}

	if newBootStrap {
		_, err = bootstrapCommit.Node.MustCreateJob(jobParams.BootstrapJob(commitStore.Hex()))
		if err != nil {
			return fmt.Errorf("shouldn't fail creating bootstrap job on bootstrap node %+v", err)
		}
		if bootstrapExec != nil && len(execNodes) > 0 {
			_, err := bootstrapExec.Node.MustCreateJob(jobParams.BootstrapJob(offRamp.Hex()))
			if err != nil {
				return fmt.Errorf("shouldn't fail creating bootstrap job on bootstrap node %+v", err)
			}
		}
	}

	if len(execNodes) == 0 {
		execNodes = commitNodes
	}

	tokenFeeConv := make(map[string]interface{})
	var linkTokenAddr []string
	for token, value := range tokenUSDMap {
		tokenFeeConv[token] = value
		linkTokenAddr = append(linkTokenAddr, token)
	}
	err = SetMockServerWithSameTokenFeeConversionValue(ctx, tokenFeeConv, execNodes, mockServer)
	if err != nil {
		return errors.WithStack(err)
	}
	err = SetMockServerWithSameTokenFeeConversionValue(ctx, tokenFeeConv, commitNodes, mockServer)
	if err != nil {
		return errors.WithStack(err)
	}

	ocr2SpecCommit, err := jobParams.CommitJobSpec()
	if err != nil {
		return errors.WithStack(err)
	}

	jobParams.P2PV2Bootstrappers = []string{p2pBootstrappersExec.P2PV2Bootstrapper()}
	ocr2SpecExec, err := jobParams.ExecutionJobSpec()
	if err != nil {
		return errors.WithStack(err)
	}
	ocr2SpecExec.Name = fmt.Sprintf("%s", ocr2SpecExec.Name)

	for i, node := range commitNodes {
		tokenPricesUSDPipeline := TokenFeeForMultipleTokenAddr(node, linkTokenAddr, mockServer)

		ocr2SpecCommit.OCR2OracleSpec.OCRKeyBundleID.SetValid(node.KeysBundle.OCR2Key.Data.ID)
		ocr2SpecCommit.OCR2OracleSpec.TransmitterID.SetValid(node.KeysBundle.EthAddress)
		ocr2SpecCommit.OCR2OracleSpec.PluginConfig["tokenPricesUSDPipeline"] = fmt.Sprintf(`"""
%s
"""`, tokenPricesUSDPipeline)

		_, err = node.Node.MustCreateJob(ocr2SpecCommit)
		if err != nil {
			return fmt.Errorf("shouldn't fail creating CCIP-Commit job on OCR node %d job name %s - %+v", i+1, ocr2SpecCommit.Name, err)
		}
	}

	if ocr2SpecExec != nil {
		for i, node := range execNodes {
			tokenPricesUSDPipeline := TokenFeeForMultipleTokenAddr(node, linkTokenAddr, mockServer)

			ocr2SpecExec.OCR2OracleSpec.PluginConfig["tokenPricesUSDPipeline"] = fmt.Sprintf(`"""
%s
"""`, tokenPricesUSDPipeline)
			ocr2SpecExec.OCR2OracleSpec.OCRKeyBundleID.SetValid(node.KeysBundle.OCR2Key.Data.ID)
			ocr2SpecExec.OCR2OracleSpec.TransmitterID.SetValid(node.KeysBundle.EthAddress)

			_, err = node.Node.MustCreateJob(ocr2SpecExec)
			if err != nil {
				return fmt.Errorf("shouldn't fail creating CCIP-Exec job on OCR node %d job name %s - %+v", i+1, ocr2SpecCommit.Name, err)
			}
		}
	}
	return nil
}

func TokenFeeForMultipleTokenAddr(node *client.CLNodesWithKeys, linkTokenAddr []string, mockserver *ctfClient.MockserverClient) string {
	source := ""
	right := ""
	for i, addr := range linkTokenAddr {
		url := fmt.Sprintf("%s/%s", mockserver.Config.ClusterURL,
			nodeContractPair(node.KeysBundle.EthAddress, addr))
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

// SetMockServerWithSameTokenFeeConversionValue sets the mock responses in mockserver that are read by chainlink nodes
// to simulate different price feed value.
func SetMockServerWithSameTokenFeeConversionValue(
	ctx context.Context,
	tokenValueAddress map[string]interface{},
	chainlinkNodes []*client.CLNodesWithKeys,
	mockserver *ctfClient.MockserverClient,
) error {
	valueAdditions, _ := errgroup.WithContext(ctx)
	for tokenAddr, value := range tokenValueAddress {
		for _, n := range chainlinkNodes {
			nodeTokenPairID := nodeContractPair(n.KeysBundle.EthAddress, tokenAddr)
			path := fmt.Sprintf("/%s", nodeTokenPairID)
			func(path string) {
				valueAdditions.Go(func() error {
					return mockserver.SetAnyValuePath(path, value)
				})
			}(path)
		}
	}
	return valueAdditions.Wait()
}

func nodeContractPair(nodeAddr, contractAddr string) string {
	return fmt.Sprintf("node_%s_contract_%s", nodeAddr[2:12], contractAddr[2:12])
}

type CCIPTestEnv struct {
	MockServer               *ctfClient.MockserverClient
	CLNodesWithKeys          map[string][]*client.CLNodesWithKeys // key - network chain-id
	CLNodes                  []*client.Chainlink
	execNodeStartIndex       int
	commitNodeStartIndex     int
	numOfAllowedFaultyCommit int
	numOfAllowedFaultyExec   int
	numOfCommitNodes         int
	numOfExecNodes           int
	K8Env                    *environment.Environment
	CLNodeWithKeyReady       *errgroup.Group // denotes if the chainlink nodes are deployed, keys are created and ready to be used for job creation
}

func (c *CCIPTestEnv) ChaosLabel(t *testing.T, srcChain, destChain string) {
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
			assert.NoError(t, err)
		}
		if i >= c.commitNodeStartIndex && i < c.commitNodeStartIndex+c.numOfAllowedFaultyCommit {
			err := c.K8Env.Client.LabelChaosGroupByLabels(c.K8Env.Cfg.Namespace, labelSelector, ChaosGroupCommitFaulty)
			assert.NoError(t, err)
		}
		if i >= c.execNodeStartIndex && i < c.execNodeStartIndex+c.numOfExecNodes {
			err := c.K8Env.Client.LabelChaosGroupByLabels(c.K8Env.Cfg.Namespace, labelSelector, ChaosGroupExecution)
			require.NoError(t, err)
		}
		if i >= c.execNodeStartIndex && i < c.execNodeStartIndex+c.numOfAllowedFaultyExec+1 {
			err := c.K8Env.Client.LabelChaosGroupByLabels(c.K8Env.Cfg.Namespace, labelSelector, ChaosGroupExecutionFaultyPlus)
			assert.NoError(t, err)
		}
		if i >= c.execNodeStartIndex && i < c.execNodeStartIndex+c.numOfAllowedFaultyExec {
			err := c.K8Env.Client.LabelChaosGroupByLabels(c.K8Env.Cfg.Namespace, labelSelector, ChaosGroupExecutionFaulty)
			assert.NoError(t, err)
		}
	}

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

func (c *CCIPTestEnv) SetUpNodesAndKeys(
	ctx context.Context,
	nodeFund *big.Float,
	chains []blockchain.EVMClient,
) error {
	log.Info().Msg("Connecting to launched resources")
	chainlinkNodes, err := client.ConnectChainlinkNodes(c.K8Env)
	if err != nil {
		return errors.WithStack(err)
	}
	if len(chainlinkNodes) == 0 {
		return fmt.Errorf("no CL node found")
	}

	mockServer, err := ctfClient.ConnectMockServer(c.K8Env)
	if err != nil {
		return fmt.Errorf("creating mockserver clients shouldn't fail %+v", err)
	}

	nodesWithKeys := make(map[string][]*client.CLNodesWithKeys)
	mu := &sync.Mutex{}
	grp, _ := errgroup.WithContext(ctx)
	populateKeys := func(chain blockchain.EVMClient) {
		grp.Go(func() error {
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
		})
	}

	fund := func(ec blockchain.EVMClient) {
		grp.Go(func() error {
			cfg := ec.GetNetworkConfig()
			if cfg == nil {
				return fmt.Errorf("blank network config")
			}
			c, err := blockchain.ConcurrentEVMClient(*cfg, c.K8Env, ec)
			if err != nil {
				return fmt.Errorf("getting concurrent evmclient chain %s %+v", c.GetNetworkName(), err)
			}
			defer func() {
				if c != nil {
					c.Close()
				}
			}()
			err = FundChainlinkNodesAddresses(chainlinkNodes[1:], c, nodeFund)
			if err != nil {
				return fmt.Errorf("funding nodes for chain %s %+v", c.GetNetworkName(), err)
			}
			return nil
		})
	}

	for _, chain := range chains {
		log.Info().Msg("creating node keys")
		populateKeys(chain)
		log.Info().Msg("Funding Chainlink nodes for both the chains")
		fund(chain)
	}

	err = grp.Wait()
	if err != nil {
		return errors.WithStack(err)
	}

	c.MockServer = mockServer
	c.CLNodesWithKeys = nodesWithKeys
	c.CLNodes = chainlinkNodes
	return nil
}

// DeployEnvironments deploys K8 env for CCIP tests. For tests running on simulated geth it deploys -
// 1. two simulated geth network in non-dev mode
// 2. mockserver ( to set mock price feed details)
// 3. chainlink nodes
func DeployEnvironments(
	t *testing.T,
	envconfig *environment.Config,
	clProps map[string]interface{},
	gethResource map[string]interface{},
	networks []blockchain.EVMNetwork,
) *CCIPTestEnv {
	testEnvironment := environment.New(envconfig)
	c := &CCIPTestEnv{
		K8Env: testEnvironment,
	}
	numOfTxNodes := 1
	for _, network := range networks {
		if network.Simulated {
			testEnvironment.
				AddHelm(reorg.New(&reorg.Props{
					NetworkName: network.Name,
					NetworkType: "simulated-geth-non-dev",
					Values: map[string]interface{}{
						"geth": map[string]interface{}{
							"genesis": map[string]interface{}{
								"networkId": fmt.Sprint(network.ChainID),
							},
							"tx": map[string]interface{}{
								"replicas":  strconv.Itoa(numOfTxNodes),
								"resources": gethResource,
							},
							"miner": map[string]interface{}{
								"replicas":  "0",
								"resources": gethResource,
							},
						},
						"bootnode": map[string]interface{}{
							"replicas": "1",
						},
					},
				}))
		}
	}

	err := testEnvironment.
		AddHelm(mockservercfg.New(nil)).
		AddHelm(mockserver.New(nil)).
		Run()
	require.NoError(t, err)
	if testEnvironment.WillUseRemoteRunner() {
		return c
	}
	urlFinder := func(network blockchain.EVMNetwork) ([]string, []string) {
		if !network.Simulated {
			return network.URLs, network.HTTPURLs
		}
		networkName := network.Name
		var internalWsURLs, internalHttpURLs []string
		for i := 0; i < numOfTxNodes; i++ {
			podName := fmt.Sprintf("%s-ethereum-geth:%d", networkName, i)
			txNodeInternalWs, err := testEnvironment.Fwd.FindPort(podName, "geth", "ws-rpc").As(envclient.RemoteConnection, envclient.WS)
			require.NoError(t, err, "Error finding WS ports")
			internalWsURLs = append(internalWsURLs, txNodeInternalWs)
			txNodeInternalHttp, err := testEnvironment.Fwd.FindPort(podName, "geth", "http-rpc").As(envclient.RemoteConnection, envclient.HTTP)
			require.NoError(t, err, "Error finding HTTP ports")
			internalHttpURLs = append(internalHttpURLs, txNodeInternalHttp)
		}
		return internalWsURLs, internalHttpURLs
	}
	var nets []blockchain.EVMNetwork
	for i := range networks {
		nets = append(nets, networks[i])
		nets[i].URLs, nets[i].HTTPURLs = urlFinder(networks[i])
		// skip adding blockscout for simplified deployments
		// uncomment the following to debug on-chain transactions
		/*
			testEnvironment.AddChart(blockscout.New(&blockscout.Props{
					Name:    fmt.Sprintf("%s-blockscout", networks[i].Name),
					WsURL:   networks[i].URLs[0],
					HttpURL: networks[i].HTTPURLs[0],
				}))
		*/
	}

	clProps["toml"] = DefaultCCIPCLNodeEnv(t, nets)
	err = testEnvironment.
		AddHelm(chainlink.New(0, clProps)).
		Run()
	require.NoError(t, err)
	return c
}

func AssertBalances(t *testing.T, bas []testhelpers.BalanceAssertion) {
	event := log.Info()
	for _, b := range bas {
		actual := b.Getter(b.Address)
		assert.NotNil(t, actual, "%v getter return nil", b.Name)
		if b.Within == "" {
			assert.Equal(t, b.Expected, actual.String(), "wrong balance for %s got %s want %s", b.Name, actual, b.Expected)
			event.Interface(b.Name, struct {
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
			event.Interface(b.Name, struct {
				ExpRange string
				Actual   string
			}{
				ExpRange: fmt.Sprintf("[%s, %s]", low, high),
				Actual:   actual.String(),
			})
		}
	}
	event.Msg("balance assertions succeeded")
}

func GetterForLinkToken(t *testing.T, token *ccip.LinkToken, addr string) func(_ common.Address) *big.Int {
	return func(_ common.Address) *big.Int {
		balance, err := token.BalanceOf(context.Background(), addr)
		assert.NoError(t, err)
		return balance
	}
}
