package actions

import (
	"context"
	_ "embed"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/smartcontractkit/chainlink-env/environment"
	"github.com/smartcontractkit/chainlink-env/pkg/helm/chainlink"
	"github.com/smartcontractkit/chainlink-env/pkg/helm/mockserver"
	mockservercfg "github.com/smartcontractkit/chainlink-env/pkg/helm/mockserver-cfg"
	"github.com/smartcontractkit/chainlink-env/pkg/helm/reorg"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	ctfClient "github.com/smartcontractkit/chainlink-testing-framework/client"
	ctfUtils "github.com/smartcontractkit/chainlink-testing-framework/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/multierr"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/price_registry"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/router"
	ccipPlugin "github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/testhelpers"
	"github.com/smartcontractkit/chainlink/core/utils"
	bigmath "github.com/smartcontractkit/chainlink/core/utils/big_math"
	networks "github.com/smartcontractkit/chainlink/integration-tests"
	"github.com/smartcontractkit/chainlink/integration-tests/client"
	"github.com/smartcontractkit/chainlink/integration-tests/contracts/ccip"
	"github.com/smartcontractkit/chainlink/integration-tests/contracts/ccip/laneconfig"
)

const (
	ChaosGroupExecution           = "ExecutionNodesAll"      // all execution nodes
	ChaosGroupCommit              = "CommitNodesAll"         // all commit nodes
	ChaosGroupCommitFaultyPlus    = "CommitMajority"         // >f number of nodes
	ChaosGroupCommitFaulty        = "CommitMinority"         //  f number of nodes
	ChaosGroupExecutionFaultyPlus = "ExecutionNodesMajority" // > f number of nodes
	ChaosGroupExecutionFaulty     = "ExecutionNodesMinority" //  f number of nodes
	RootSnoozeTime                = 10 * time.Second
	InflightExpiry                = 10 * time.Second
)

type CCIPTOMLEnv struct {
	Networks []blockchain.EVMNetwork
}

var (
	UnusedFee = big.NewInt(0).Mul(big.NewInt(11), big.NewInt(1e18)) // for a msg with two tokens

	//go:embed clconfig/ccip-default.txt
	CLConfig           string
	NetworkA, NetworkB = func() (blockchain.EVMNetwork, blockchain.EVMNetwork) {
		if len(networks.SelectedNetworks) < 3 {
			log.Fatal().
				Interface("SELECTED_NETWORKS", networks.SelectedNetworks).
				Msg("Set source and destination network in index 1 & 2 of env variable SELECTED_NETWORKS")
		}
		log.Info().
			Interface("Source Network", networks.SelectedNetworks[1]).
			Interface("Destination Network", networks.SelectedNetworks[2]).
			Msg("SELECTED_NETWORKS")
		return networks.SelectedNetworks[1], networks.SelectedNetworks[2]
	}()

	DefaultCCIPCLNodeEnv = func(t *testing.T) string {
		ccipTOML, err := client.MarshallTemplate(
			CCIPTOMLEnv{
				Networks: []blockchain.EVMNetwork{NetworkA, NetworkB},
			},
			"ccip env toml", CLConfig)
		require.NoError(t, err)
		fmt.Println("Configuration", ccipTOML)
		return ccipTOML
	}

	networkAName = strings.ReplaceAll(strings.ToLower(NetworkA.Name), " ", "-")
	networkBName = strings.ReplaceAll(strings.ToLower(NetworkB.Name), " ", "-")
)

type CCIPCommon struct {
	ChainClient              blockchain.EVMClient
	Deployer                 *ccip.CCIPContractsDeployer
	FeeToken                 *ccip.LinkToken
	FeeTokenPool             *ccip.LockReleaseTokenPool
	BridgeTokens             []*ccip.LinkToken // as of now considering the bridge token is same as link token
	TokenPrices              []*big.Int
	BridgeTokenPools         []*ccip.LockReleaseTokenPool
	RateLimiterConfig        ccip.RateLimiterConfig
	AFNConfig                ccip.AFNConfig
	AFN                      *ccip.AFN
	Router                   *ccip.Router
	PriceRegistry            *ccip.PriceRegistry
	WrappedNative            common.Address
	deployedNewPool          bool
	deployedNewPriceRegistry bool
	deploy                   *errgroup.Group
}

func (ccipModule *CCIPCommon) CopyAddresses(ctx context.Context, chainClient blockchain.EVMClient) *CCIPCommon {
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
		FeeTokenPool: &ccip.LockReleaseTokenPool{
			EthAddress: ccipModule.FeeTokenPool.EthAddress,
		},
		BridgeTokens:      tokens,
		TokenPrices:       ccipModule.TokenPrices,
		BridgeTokenPools:  pools,
		RateLimiterConfig: ccipModule.RateLimiterConfig,
		AFNConfig: ccip.AFNConfig{
			AFNWeightsByParticipants: map[string]*big.Int{
				chainClient.GetDefaultWallet().Address(): big.NewInt(1),
			},
			ThresholdForBlessing:  big.NewInt(1),
			ThresholdForBadSignal: big.NewInt(1),
		},
		AFN: &ccip.AFN{
			EthAddress: ccipModule.AFN.EthAddress,
		},
		Router: &ccip.Router{
			EthAddress: ccipModule.Router.EthAddress,
		},
		WrappedNative: ccipModule.WrappedNative,
		deploy:        grp,
	}
}

func (ccipModule *CCIPCommon) LoadContractAddresses(conf *laneconfig.Lane) {
	if conf != nil {
		if common.IsHexAddress(conf.LaneConfig.FeeToken) {
			ccipModule.FeeToken = &ccip.LinkToken{
				EthAddress: common.HexToAddress(conf.LaneConfig.FeeToken),
			}
		}
		if conf.LaneConfig.IsNativeFeeToken {
			ccipModule.FeeToken = &ccip.LinkToken{
				EthAddress: common.HexToAddress("0x0"),
			}
		}

		if common.IsHexAddress(conf.LaneConfig.FeeTokenPool) {
			ccipModule.FeeTokenPool = &ccip.LockReleaseTokenPool{
				EthAddress: common.HexToAddress(conf.LaneConfig.FeeTokenPool),
			}
		}
		if common.IsHexAddress(conf.LaneConfig.Router) {
			ccipModule.Router = &ccip.Router{
				EthAddress: common.HexToAddress(conf.LaneConfig.Router),
			}
		}
		if common.IsHexAddress(conf.LaneConfig.AFN) {
			ccipModule.AFN = &ccip.AFN{
				EthAddress: common.HexToAddress(conf.LaneConfig.AFN),
			}
		}
		if common.IsHexAddress(conf.LaneConfig.PriceRegistry) {
			ccipModule.PriceRegistry = &ccip.PriceRegistry{
				EthAddress: common.HexToAddress(conf.LaneConfig.PriceRegistry),
			}
		}
		if len(conf.LaneConfig.BridgeTokens) > 0 {
			var tokens []*ccip.LinkToken
			for _, token := range conf.LaneConfig.BridgeTokens {
				if common.IsHexAddress(token) {
					tokens = append(tokens, &ccip.LinkToken{
						EthAddress: common.HexToAddress(token),
					})
				}
			}
			ccipModule.BridgeTokens = tokens
		}
		if len(conf.LaneConfig.BridgeTokenPools) > 0 {
			var pools []*ccip.LockReleaseTokenPool
			for _, pool := range conf.LaneConfig.BridgeTokenPools {
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

// DeployContracts deploys the contracts which are necessary in both source and dest chain
// This reuses common contracts for bidirectional lanes
func (ccipModule *CCIPCommon) DeployContracts(noOfTokens int, conf *laneconfig.Lane, reuseContracts bool) error {
	var err error
	cd := ccipModule.Deployer

	// if REUSE_CCIP_CONTRACTS is true use existing contract addresses instead of deploying new ones. Used for testnet deployments
	if reuseContracts && !ccipModule.ChainClient.NetworkSimulated() {
		ccipModule.LoadContractAddresses(conf)
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
	if ccipModule.FeeTokenPool == nil {
		// token pool for fee token
		ccipModule.FeeTokenPool, err = cd.DeployLockReleaseTokenPoolContract(ccipModule.FeeToken.Address())
		if err != nil {
			return fmt.Errorf("deploying fee Token pool shouldn't fail %+v", err)
		}
		ccipModule.deployedNewPool = true
	} else {
		pool, err := cd.NewLockReleaseTokenPoolContract(ccipModule.FeeTokenPool.EthAddress)
		if err != nil {
			return fmt.Errorf("getting new fee token pool contract shouldn't fail %+v", err)
		}
		ccipModule.FeeTokenPool = pool
	}
	var btAddresses, btpAddresses []string
	if len(ccipModule.BridgeTokens) != noOfTokens {
		// deploy bridge token.
		for i := 0; i < noOfTokens; i++ {
			token, err := cd.DeployLinkTokenContract()
			if err != nil {
				return fmt.Errorf("deploying bridge token contract shouldn't fail %+v", err)
			}
			ccipModule.BridgeTokens = append(ccipModule.BridgeTokens, token)
			btAddresses = append(btAddresses, token.Address())
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
			btAddresses = append(btAddresses, token.Address())
		}
		ccipModule.BridgeTokens = tokens
	}
	if len(ccipModule.BridgeTokenPools) != noOfTokens {
		// deploy native token pool
		for _, token := range ccipModule.BridgeTokens {
			ntp, err := cd.DeployLockReleaseTokenPoolContract(token.Address())
			if err != nil {
				return fmt.Errorf("deploying bridge Token pool shouldn't fail %+v", err)
			}
			ccipModule.BridgeTokenPools = append(ccipModule.BridgeTokenPools, ntp)
			btpAddresses = append(btpAddresses, ntp.Address())
		}
	} else {
		var pools []*ccip.LockReleaseTokenPool
		for _, pool := range ccipModule.BridgeTokenPools {
			newPool, err := cd.NewLockReleaseTokenPoolContract(pool.EthAddress)
			if err != nil {
				return fmt.Errorf("getting new bridge token pool contract shouldn't fail %+v", err)
			}
			pools = append(pools, newPool)
			btpAddresses = append(btpAddresses, newPool.Address())
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
		router, err := cd.NewRouter(ccipModule.Router.EthAddress)
		if err != nil {
			return fmt.Errorf("getting new router contract shouldn't fail %+v", err)
		}
		ccipModule.Router = router
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
		ccipModule.deployedNewPriceRegistry = true
	} else {
		ccipModule.PriceRegistry, err = cd.NewPriceRegistry(ccipModule.PriceRegistry.EthAddress)
		if err != nil {
			return fmt.Errorf("getting new PriceRegistry contract shouldn't fail %+v", err)
		}
	}

	log.Info().Msg("finished deploying common contracts")
	return nil
}

func DefaultCCIPModule(ctx context.Context, chainClient blockchain.EVMClient) *CCIPCommon {
	grp, _ := errgroup.WithContext(ctx)
	return &CCIPCommon{
		ChainClient: chainClient,
		deploy:      grp,
		RateLimiterConfig: ccip.RateLimiterConfig{
			Rate:     ccip.HundredCoins,
			Capacity: ccip.HundredCoins,
		},
		AFNConfig: ccip.AFNConfig{
			AFNWeightsByParticipants: map[string]*big.Int{
				chainClient.GetDefaultWallet().Address(): big.NewInt(1),
			},
			ThresholdForBlessing:  big.NewInt(1),
			ThresholdForBadSignal: big.NewInt(1),
		},
	}
}

type SourceCCIPModule struct {
	Common             *CCIPCommon
	Sender             common.Address
	TransferAmount     []*big.Int
	DestinationChainId uint64
	OnRamp             *ccip.OnRamp
}

func (sourceCCIP *SourceCCIPModule) LoadContracts(conf *laneconfig.Lane) {
	if conf != nil {
		cfg, ok := conf.LaneConfig.SrcContracts[sourceCCIP.DestinationChainId]
		if ok {
			if common.IsHexAddress(cfg.OnRamp) {
				sourceCCIP.OnRamp = &ccip.OnRamp{
					EthAddress: common.HexToAddress(cfg.OnRamp),
				}
			}
		}
	}
}

// DeployContracts deploys all CCIP contracts specific to the source chain
func (sourceCCIP *SourceCCIPModule) DeployContracts(reuse bool, lane *laneconfig.Lane) error {
	var err error
	contractDeployer := sourceCCIP.Common.Deployer
	log.Info().Msg("Deploying source chain specific contracts")
	err = sourceCCIP.Common.deploy.Wait()
	if err != nil {
		return errors.WithStack(err)
	}

	if reuse && !sourceCCIP.Common.ChainClient.NetworkSimulated() {
		sourceCCIP.LoadContracts(lane)
	}
	if sourceCCIP.Common.deployedNewPriceRegistry {
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

		sourceCCIP.OnRamp, err = contractDeployer.DeployOnRamp(
			sourceCCIP.Common.ChainClient.GetChainID().Uint64(),
			sourceCCIP.DestinationChainId,
			[]common.Address{},
			tokensAndPools,
			sourceCCIP.Common.AFN.EthAddress,
			sourceCCIP.Common.Router.EthAddress,
			sourceCCIP.Common.PriceRegistry.EthAddress,
			sourceCCIP.Common.RateLimiterConfig,
			[]evm_2_evm_onramp.IEVM2EVMOnRampFeeTokenConfigArgs{
				{
					Token:           common.HexToAddress(sourceCCIP.Common.FeeToken.Address()),
					FeeAmount:       big.NewInt(0),
					DestGasOverhead: 0,
					Multiplier:      1e18,
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

func (sourceCCIP *SourceCCIPModule) AssertEventCCIPSendRequested(t *testing.T, txHash string, currentBlockOnSource uint64, timeout time.Duration) (uint64, [32]byte) {
	log.Info().Msg("Waiting for CCIPSendRequested event")
	var seqNum uint64
	var msgId [32]byte
	gom := NewWithT(t)
	gom.Eventually(func(g Gomega) bool {
		iterator, err := sourceCCIP.OnRamp.FilterCCIPSendRequested(currentBlockOnSource)
		g.Expect(err).NotTo(HaveOccurred(), "Error filtering CCIPSendRequested event")
		for iterator.Next() {
			if strings.EqualFold(iterator.Event.Raw.TxHash.Hex(), txHash) {
				seqNum = iterator.Event.Message.SequenceNumber
				msgId = iterator.Event.Message.MessageId
				return true
			}
		}
		return false
	}, timeout, "1s").Should(BeTrue(), "No CCIPSendRequested event found with txHash %s", txHash)

	return seqNum, msgId
}

func (sourceCCIP *SourceCCIPModule) SendRequest(
	t *testing.T,
	receiver common.Address,
	tokenAndAmounts []router.ClientEVMTokenAmount,
	data string,
	feeToken common.Address,
) (string, *big.Int) {
	receiverAddr, err := utils.ABIEncode(`[{"type":"address"}]`, receiver)
	assert.NoError(t, err, "Failed encoding the receiver address")

	extraArgsV1, err := testhelpers.GetEVMExtraArgsV1(big.NewInt(100_000), false)
	assert.NoError(t, err, "Failed encoding the options field")

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
	assert.NoError(t, err, "calculating fee")
	log.Info().Int64("fee", fee.Int64()).Msg("calculated fee")

	// if the token address is 0x0 it will use Native as fee token no need to approve
	if feeToken != common.HexToAddress("0x0") {
		// if any of the bridge tokens are same as feetoken approve the token amount as well
		// otherwise the fee amount approval of the feetoken will overwrite previous transfer amountapproval for same bridge token
		for index, bt := range sourceCCIP.Common.BridgeTokens {
			if bt.Address() == sourceCCIP.Common.FeeToken.Address() {
				fee = big.NewInt(0).Add(fee, sourceCCIP.TransferAmount[index])
			}
		}
		// Approve the fee amount
		err = sourceCCIP.Common.FeeToken.Approve(sourceCCIP.Common.Router.Address(), fee)
		require.NoError(t, err, "approving fee for ge router")
		require.NoError(t, sourceCCIP.Common.ChainClient.WaitForEvents(), "error waiting for events")
	}

	var sendTx *types.Transaction
	// initiate the transfer
	// if the token address is 0x0 it will use Native as fee token and the fee amount should be mentioned in bind.TransactOpts's value
	if feeToken != common.HexToAddress("0x0") {
		sendTx, err = sourceCCIP.Common.Router.CCIPSend(sourceCCIP.DestinationChainId, msg, nil)
		require.NoError(t, err, "send token should be initiated successfully")
	} else {
		sendTx, err = sourceCCIP.Common.Router.CCIPSend(sourceCCIP.DestinationChainId, msg, fee)
		require.NoError(t, err, "send token should be initiated successfully")
	}

	log.Info().Str("Send token transaction", sendTx.Hash().String()).Msg("Sending token")
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	assert.NoError(t, err, "Failed to wait for events")
	return sendTx.Hash().Hex(), fee
}

func DefaultSourceCCIPModule(chainClient blockchain.EVMClient, destChain uint64, transferAmount []*big.Int, ccipCommon *CCIPCommon) (*SourceCCIPModule, error) {
	sourceCCIP := &SourceCCIPModule{
		Common:             ccipCommon,
		TransferAmount:     transferAmount,
		DestinationChainId: destChain,
		Sender:             common.HexToAddress(chainClient.GetDefaultWallet().Address()),
	}
	var err error
	sourceCCIP.Common.Deployer, err = ccip.NewCCIPContractsDeployer(chainClient)
	if err != nil {
		return nil, fmt.Errorf("contract deployer should be created successfully %+v", err)
	}

	return sourceCCIP, nil
}

type DestCCIPModule struct {
	Common        *CCIPCommon
	SourceChainId uint64
	CommitStore   *ccip.CommitStore
	ReceiverDapp  *ccip.ReceiverDapp
	OffRamp       *ccip.OffRamp
	WrappedNative common.Address
}

func (destCCIP *DestCCIPModule) LoadContracts(conf *laneconfig.Lane) {
	if conf != nil {
		cfg, ok := conf.LaneConfig.DestContracts[destCCIP.SourceChainId]
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
	lane *laneconfig.Lane,
	reuse bool,
) error {
	var err error
	contractDeployer := destCCIP.Common.Deployer
	log.Info().Msg("Deploying destination chain specific contracts")

	err = destCCIP.Common.deploy.Wait()
	if err != nil {
		return errors.WithStack(err)
	}

	if reuse && !destCCIP.Common.ChainClient.NetworkSimulated() {
		destCCIP.LoadContracts(lane)
	}
	if destCCIP.Common.deployedNewPriceRegistry {
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
		if destCCIP.Common.deployedNewPool {
			err = token.Transfer(pool.Address(), testhelpers.Link(10))
			if err != nil {
				return fmt.Errorf("transferring token to dest pool shouldn't fail %+v", err)
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

func (destCCIP *DestCCIPModule) AssertEventExecutionStateChanged(t *testing.T, seqNum uint64, msgID [32]byte, currentBlockOnDest uint64, timeout time.Duration) {
	log.Info().Int64("seqNum", int64(seqNum)).
		Msgf("Waiting for ExecutionStateChanged event for lane %d-->%d",
			destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID())
	gom := NewWithT(t)
	gom.Eventually(func(g Gomega) ccipPlugin.MessageExecutionState {
		iterator, err := destCCIP.OffRamp.FilterExecutionStateChanged([]uint64{seqNum}, [][32]byte{msgID}, currentBlockOnDest)
		g.Expect(err).NotTo(HaveOccurred(),
			"filtering ExecutionStateChanged event for seqNum %d for lane %d-->%d",
			seqNum, destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID())
		g.Expect(iterator.Next()).To(BeTrue(),
			"no ExecutionStateChanged event found for seqNum %d for lane %d-->%d",
			seqNum, destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID())
		return ccipPlugin.MessageExecutionState(iterator.Event.State)
	}, timeout, "1s").Should(Equal(ccipPlugin.Success), "execution state lane %d-->%d",
		destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID())
}

func (destCCIP *DestCCIPModule) AssertEventReportAccepted(t *testing.T, onRamp common.Address, seqNum, currentBlockOnDest uint64, timeout time.Duration) {
	log.Info().Int64("seqNum", int64(seqNum)).Msgf("Waiting for ReportAccepted event lane %d-->%d", destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID())
	gom := NewWithT(t)
	gom.Eventually(func(g Gomega) bool {
		iterator, err := destCCIP.CommitStore.FilterReportAccepted(currentBlockOnDest)
		g.Expect(err).NotTo(HaveOccurred(),
			"filtering ReportAccepted event lane %d-->%d",
			destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID())
		for iterator.Next() {
			if iterator.Event.Report.Interval.Min <= seqNum && iterator.Event.Report.Interval.Max >= seqNum {
				return true
			}
		}
		return false
	}, timeout, "1s").Should(BeTrue(),
		"no ReportAccepted Event found for onRamp %s and seq num %d lane %d-->%d",
		onRamp.Hex(), seqNum, destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID())
}

func (destCCIP *DestCCIPModule) AssertSeqNumberExecuted(t *testing.T, onRamp common.Address, seqNumberBefore uint64, timeout time.Duration) {
	log.Info().Int64("seqNum", int64(seqNumberBefore)).
		Msgf("Waiting to be executed lane %d-->%d", destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID())
	gom := NewWithT(t)
	gom.Eventually(func(g Gomega) bool {
		seqNumberAfter, err := destCCIP.CommitStore.GetNextSeqNumber()
		if err != nil {
			return false
		}
		if seqNumberAfter > seqNumberBefore {
			return true
		}
		return false
	}, timeout, "1s").Should(BeTrue(),
		"error executing sequence number %d lane %d-->%d",
		seqNumberBefore, destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID())
}

func DefaultDestinationCCIPModule(chainClient blockchain.EVMClient, sourceChain uint64, ccipCommon *CCIPCommon) (*DestCCIPModule, error) {
	destCCIP := &DestCCIPModule{
		Common:        ccipCommon,
		SourceChainId: sourceChain,
	}
	var err error
	destCCIP.Common.Deployer, err = ccip.NewCCIPContractsDeployer(chainClient)
	if err != nil {
		return nil, err
	}
	return destCCIP, nil
}

type CCIPLane struct {
	t                       *testing.T
	SourceNetworkName       string
	DestNetworkName         string
	SourceChain             blockchain.EVMClient
	DestChain               blockchain.EVMClient
	Source                  *SourceCCIPModule
	Dest                    *DestCCIPModule
	TestEnv                 *CCIPTestEnv
	NumberOfReq             int
	SourceBalances          map[string]*big.Int
	DestBalances            map[string]*big.Int
	StartBlockOnSource      uint64
	StartBlockOnDestination uint64
	SentReqHashes           []string
	TotalFee                *big.Int // total fee for all the requests. Used for balance validation.
	ValidationTimeout       time.Duration
	context                 context.Context
	commonContractsWg       *sync.WaitGroup
}

func (lane *CCIPLane) WriteLaneConfig() error {
	var btAddresses, btpAddresses []string

	if lane.Dest.Common.ChainClient.NetworkSimulated() ||
		lane.Source.Common.ChainClient.NetworkSimulated() {
		return nil
	}
	for i, bt := range lane.Source.Common.BridgeTokens {
		btAddresses = append(btAddresses, bt.Address())
		btpAddresses = append(btpAddresses, lane.Source.Common.BridgeTokenPools[i].Address())
	}
	l1 := laneconfig.Lane{
		NetworkA: lane.Source.Common.ChainClient.GetNetworkName(),
		LaneConfig: laneconfig.LaneConfig{
			CommonContracts: laneconfig.CommonContracts{
				FeeToken:         lane.Source.Common.FeeToken.Address(),
				FeeTokenPool:     lane.Source.Common.FeeTokenPool.Address(),
				BridgeTokens:     btAddresses,
				BridgeTokenPools: btpAddresses,
				AFN:              lane.Source.Common.AFN.Address(),
				Router:           lane.Source.Common.Router.Address(),
				PriceRegistry:    lane.Source.Common.PriceRegistry.Address(),
			},
			SrcContracts: map[uint64]laneconfig.SourceContracts{
				lane.Source.DestinationChainId: {OnRamp: lane.Source.OnRamp.Address()},
			},
		},
	}
	btAddresses, btpAddresses = []string{}, []string{}
	for i, bt := range lane.Dest.Common.BridgeTokens {
		btAddresses = append(btAddresses, bt.Address())
		btpAddresses = append(btpAddresses, lane.Dest.Common.BridgeTokenPools[i].Address())
	}
	l2 := laneconfig.Lane{
		NetworkA: lane.Dest.Common.ChainClient.GetNetworkName(),
		LaneConfig: laneconfig.LaneConfig{
			CommonContracts: laneconfig.CommonContracts{
				FeeToken:         lane.Dest.Common.FeeToken.Address(),
				FeeTokenPool:     lane.Dest.Common.FeeTokenPool.Address(),
				BridgeTokens:     btAddresses,
				BridgeTokenPools: btpAddresses,
				AFN:              lane.Dest.Common.AFN.Address(),
				Router:           lane.Dest.Common.Router.Address(),
				PriceRegistry:    lane.Dest.Common.PriceRegistry.Address(),
			},
			DestContracts: map[uint64]laneconfig.DestContracts{
				lane.Dest.SourceChainId: {
					OffRamp:      lane.Dest.OffRamp.Address(),
					CommitStore:  lane.Dest.CommitStore.Address(),
					ReceiverDapp: lane.Dest.ReceiverDapp.Address(),
				},
			},
		},
	}
	return laneconfig.UpdateLane(l1, l2, fmt.Sprintf("./tmp_%s.json", lane.t.Name()))
}

func (lane *CCIPLane) RecordStateBeforeTransfer() {
	var err error
	// collect the balance assert.ment to verify balances after transfer
	lane.SourceBalances, err = testhelpers.GetBalances(lane.Source.CollectBalanceRequirements(lane.t))
	require.NoError(lane.t, err, "fetching source balance")
	lane.DestBalances, err = testhelpers.GetBalances(lane.Dest.CollectBalanceRequirements(lane.t))
	require.NoError(lane.t, err, "fetching dest balance")

	// save the current block numbers to use in various filter log requests
	lane.StartBlockOnSource, err = lane.Source.Common.ChainClient.LatestBlockNumber(context.Background())
	require.NoError(lane.t, err, "Getting current block should be successful in source chain")
	lane.StartBlockOnDestination, err = lane.Dest.Common.ChainClient.LatestBlockNumber(context.Background())
	require.NoError(lane.t, err, "Getting current block should be successful in dest chain")
	lane.TotalFee = big.NewInt(0)
	lane.NumberOfReq = 0
	lane.SentReqHashes = []string{}
}

func (lane *CCIPLane) SendRequests(noOfRequests int) []string {
	t := lane.t
	lane.NumberOfReq += noOfRequests
	var tokenAndAmounts []router.ClientEVMTokenAmount
	for i, token := range lane.Source.Common.BridgeTokens {
		tokenAndAmounts = append(tokenAndAmounts, router.ClientEVMTokenAmount{
			Token: common.HexToAddress(token.Address()), Amount: lane.Source.TransferAmount[i],
		})
		// approve the onramp router so that it can initiate transferring the token

		err := token.Approve(lane.Source.Common.Router.Address(), bigmath.Mul(lane.Source.TransferAmount[i], big.NewInt(int64(noOfRequests))))
		require.NoError(t, err, "Could not approve permissions for the onRamp router "+
			"on the source link token contract")
	}

	err := lane.Source.Common.ChainClient.WaitForEvents()
	require.NoError(t, err, "Failed to wait for events")
	var txs []string
	for i := 1; i <= noOfRequests; i++ {
		txHash, fee := lane.Source.SendRequest(
			t, lane.Dest.ReceiverDapp.EthAddress,
			tokenAndAmounts,
			fmt.Sprintf("msg %d", i),
			common.HexToAddress(lane.Source.Common.FeeToken.Address()),
		)
		txs = append(txs, txHash)
		lane.SentReqHashes = append(lane.SentReqHashes, txHash)
		lane.TotalFee = bigmath.Add(lane.TotalFee, fee)
	}
	return txs
}

func (lane *CCIPLane) ValidateRequests() {
	for _, txHash := range lane.SentReqHashes {
		lane.ValidateRequestByTxHash(txHash)
	}
	// verify the fee amount is deducted from sender, added to receiver token balances and
	// unused fee is returned to receiver fee token account
	AssertBalances(lane.t, lane.Source.BalanceAssertions(lane.t, lane.SourceBalances, int64(lane.NumberOfReq), lane.TotalFee))
	AssertBalances(lane.t, lane.Dest.BalanceAssertions(lane.t, lane.DestBalances, lane.Source.TransferAmount, int64(lane.NumberOfReq)))
}

func (lane *CCIPLane) ValidateRequestByTxHash(txHash string) {
	t := lane.t
	// Verify if
	// - CCIPSendRequested Event log generated,
	// - NextSeqNumber from commitStore got increased
	seqNumber, msgId := lane.Source.AssertEventCCIPSendRequested(t, txHash, lane.StartBlockOnSource, lane.ValidationTimeout)
	lane.Dest.AssertSeqNumberExecuted(t, lane.Source.OnRamp.EthAddress, seqNumber, lane.ValidationTimeout)

	// Verify whether commitStore has accepted the report
	lane.Dest.AssertEventReportAccepted(t, lane.Source.OnRamp.EthAddress, seqNumber, lane.StartBlockOnDestination, lane.ValidationTimeout)

	// Verify whether the execution state is changed and the transfer is successful
	lane.Dest.AssertEventExecutionStateChanged(t, seqNumber, msgId, lane.StartBlockOnDestination, lane.ValidationTimeout)
}

func (lane *CCIPLane) SoakRun(interval, duration time.Duration) (int, int) {
	t := lane.t
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
			wg.Add(1)
			log.Info().
				Int("Req No", numOfReq).
				Msgf("Token transfer with for lane %s --> %s", lane.SourceNetworkName, lane.DestNetworkName)
			txs := lane.SendRequests(1)
			assert.NotEmpty(t, txs)
			go func(txHash string) {
				defer wg.Done()
				lane.ValidateRequestByTxHash(txHash)
				reqSuccess++
			}(txs[0])
		case <-ctx.Done():
			log.Warn().
				Msgf("Soak Test duration completed for lane %s --> %s. Completing validation for triggered requests",
					lane.SourceNetworkName, lane.DestNetworkName)
			timeout = true
			wg.Wait()
			return numOfReq, reqSuccess
		}
	}
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
) error {
	var err error
	env := lane.TestEnv
	sourceChainClient := lane.SourceChain
	destChainClient := lane.DestChain

	if sourceCommon == nil {
		sourceCommon = DefaultCCIPModule(lane.context, sourceChainClient)
	}

	if destCommon == nil {
		destCommon = DefaultCCIPModule(lane.context, destChainClient)
	}

	lane.Source, err = DefaultSourceCCIPModule(sourceChainClient, destChainClient.GetChainID().Uint64(), transferAmounts, sourceCommon)
	if err != nil {
		return errors.WithStack(err)
	}
	lane.Dest, err = DefaultDestinationCCIPModule(destChainClient, sourceChainClient.GetChainID().Uint64(), destCommon)
	if err != nil {
		return errors.WithStack(err)
	}

	// if REUSE_CCIP_CONTRACTS= true, read already deployed contract addresses and
	// reuse contracts in tests instead of deploying the contracts.
	reuseEnv := os.Getenv("REUSE_CCIP_CONTRACTS")
	reuse := false
	if reuseEnv != "" {
		reuse, err = strconv.ParseBool(reuseEnv)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	srcConf, err := laneconfig.ReadLane(
		lane.Source.Common.ChainClient.GetNetworkName())
	if err != nil {
		return errors.WithStack(err)
	}

	destConf, err := laneconfig.ReadLane(
		lane.Dest.Common.ChainClient.GetNetworkName())
	if err != nil {
		return errors.WithStack(err)
	}

	// deploy all common contracts in parallel
	lane.Source.Common.deploy.Go(func() error {
		return lane.Source.Common.DeployContracts(len(lane.Source.TransferAmount), srcConf, reuse)
	})

	lane.Dest.Common.deploy.Go(func() error {
		return lane.Dest.Common.DeployContracts(len(lane.Source.TransferAmount), destConf, reuse)
	})

	// deploy all source contracts
	err = lane.Source.DeployContracts(reuse, srcConf)
	if err != nil {
		return errors.WithStack(err)
	}
	// deploy all destination contracts
	err = lane.Dest.DeployContracts(*lane.Source, lane.commonContractsWg, destConf, reuse)
	if err != nil {
		return errors.WithStack(err)
	}
	err = lane.WriteLaneConfig()
	if err != nil {
		return errors.WithStack(err)
	}

	// if lane is being set up for already configured CL nodes and contracts
	// no further action is necessary
	if !configureCLNodes {
		return nil
	}

	// if lane is being set up for already configured CL nodes and contracts
	// no further action is necessary
	if !configureCLNodes {
		return nil
	}

	// wait for the CL nodes to be ready before moving ahead with job creation
	<-env.clNodeWithKeyReady
	clNodesWithKeys := env.CLNodesWithKeys
	mockServer := env.MockServer
	// set up ocr2 jobs
	var tokenAddr []string
	for _, token := range lane.Dest.Common.BridgeTokens {
		tokenAddr = append(tokenAddr, token.Address())
	}
	clNodes, exists := clNodesWithKeys[lane.Dest.Common.ChainClient.GetChainID().String()]
	if !exists {
		return fmt.Errorf("could not find CL nodes for %s", lane.Dest.Common.ChainClient.GetChainID().String())
	}

	tokenAddr = append(tokenAddr, lane.Dest.Common.FeeToken.Address(), lane.Source.Common.WrappedNative.Hex())

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

	err = CreateOCRJobsForCCIP(
		lane.context,
		bootstrapCommit, bootstrapExec, commitNodes, execNodes,
		lane.Source.OnRamp.EthAddress,
		lane.Dest.CommitStore.EthAddress,
		lane.Dest.OffRamp.EthAddress,
		sourceChainClient, destChainClient,
		tokenAddr,
		mockServer, newBootstrap,
	)
	if err != nil {
		return errors.WithStack(err)
	}

	// set up ocr2 config
	err = SetOCR2Configs(commitNodes, execNodes, *lane.Dest)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// SetOCR2Configs sets the oracle config in ocr2 contracts
// nil value in execNodes denotes commit and execution jobs are to be set up in same DON
func SetOCR2Configs(commitNodes, execNodes []*client.CLNodesWithKeys, destCCIP DestCCIPModule) error {
	signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, err :=
		ccip.NewOffChainAggregatorV2Config(commitNodes)
	if err != nil {
		return errors.WithStack(err)
	}

	err = destCCIP.CommitStore.SetOCR2Config(signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
	if err != nil {
		return errors.WithStack(err)
	}
	// if commit and exec job is set up in different DON
	if len(execNodes) > 0 {
		signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, err =
			ccip.NewOffChainAggregatorV2Config(execNodes)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	if destCCIP.OffRamp != nil {
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
	linkTokenAddr []string,
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
	currentBlockOnSource, err := sourceChainClient.LatestBlockNumber(context.Background())
	if err != nil {
		return fmt.Errorf("getting current block should be successful in source chain %+v", err)
	}
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
		SourceStartBlock:   currentBlockOnSource,
		DestStartBlock:     currentBlockOnDest,
		RelayInflight:      InflightExpiry,
		ExecInflight:       InflightExpiry,
		RootSnooze:         RootSnoozeTime,
		P2PV2Bootstrappers: []string{p2pBootstrappersCommit.P2PV2Bootstrapper()},
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
	for _, token := range linkTokenAddr {
		tokenFeeConv[token] = "200000000000000000000"
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
	SourceChainClient        blockchain.EVMClient
	DestChainClient          blockchain.EVMClient
	K8Env                    *environment.Environment
	clNodeWithKeyReady       chan struct{} // denotes if the chainlink nodes are deployed, keys are created and ready to be used for job creation
}

func (c *CCIPTestEnv) ChaosLabel(t *testing.T) {
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
}

func (c *CCIPTestEnv) SetUpNodesAndKeys(
	ctx context.Context,
	nodeFund *big.Float,
	source, dest blockchain.EVMClient,
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
	log.Info().Msg("creating node keys")
	grp, _ := errgroup.WithContext(ctx)
	populateKeys := func(chain blockchain.EVMClient) {
		grp.Go(func() error {
			_, clNodes, err := client.CreateNodeKeysBundle(chainlinkNodes, "evm", chain.GetChainID().String())
			if err != nil {
				return errors.WithStack(err)
			}
			if len(clNodes) == 0 {
				return fmt.Errorf("no CL node with keys found")
			}
			mu.Lock()
			defer mu.Unlock()
			nodesWithKeys[chain.GetChainID().String()] = clNodes
			return nil
		})
	}

	populateKeys(source)
	populateKeys(dest)

	log.Info().Msg("Funding Chainlink nodes for both the chains")

	fund := func(chain blockchain.EVMClient) {
		grp.Go(func() error {
			cfg := chain.GetNetworkConfig()
			if cfg == nil {
				return fmt.Errorf("blank network config")
			}
			c, err := blockchain.ConcurrentEVMClient(*cfg, c.K8Env, chain)
			if err != nil {
				return fmt.Errorf("getting concurrent evmclient %+v", err)
			}
			err = FundChainlinkNodesAddresses(chainlinkNodes[1:], c, nodeFund)
			if err != nil {
				return fmt.Errorf("funding nodes %+v", err)
			}
			return nil
		})
	}

	fund(source)
	fund(dest)

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
// The deployment of chainlink nodes and mockserver are done in a go routine and the corresponding
// WaitGroup is returned. To proceed forward with the tests you need to wait for the WaitGroup returned.
// For Example refer to CCIPDefaultTestSetUp function
func DeployEnvironments(
	t *testing.T,
	envconfig *environment.Config,
	clProps map[string]interface{},
) *CCIPTestEnv {
	testEnvironment := environment.New(envconfig)
	c := &CCIPTestEnv{
		K8Env: testEnvironment,
	}
	if NetworkA.Simulated {
		if !NetworkB.Simulated {
			t.Fatalf("both the networks should be simulated")
		}
		testEnvironment.
			AddHelm(reorg.New(&reorg.Props{
				NetworkName: NetworkA.Name,
				NetworkType: "simulated-geth-non-dev",
				Values: map[string]interface{}{
					"geth": map[string]interface{}{
						"genesis": map[string]interface{}{
							"networkId": fmt.Sprint(NetworkA.ChainID),
						},
						"tx": map[string]interface{}{
							"replicas": "2",
						},
						"miner": map[string]interface{}{
							"replicas": "0",
						},
					},
					"bootnode": map[string]interface{}{
						"replicas": "1",
					},
				},
			})).
			AddHelm(reorg.New(&reorg.Props{
				NetworkName: NetworkB.Name,
				NetworkType: "simulated-geth-non-dev",
				Values: map[string]interface{}{
					"geth": map[string]interface{}{
						"genesis": map[string]interface{}{
							"networkId": fmt.Sprint(NetworkB.ChainID),
						},
						"tx": map[string]interface{}{
							"replicas": "2",
						},
						"miner": map[string]interface{}{
							"replicas": "0",
						},
					},
					"bootnode": map[string]interface{}{
						"replicas": "1",
					},
				},
			}))
	}

	// skip adding blockscout for simplified deployments
	// uncomment the following to debug on-chain transactions
	/*
		testEnvironment.AddChart(blockscout.New(&blockscout.Props{
			Name:    "dest-blockscout",
			WsURL:   NetworkB.URLs[0],
			HttpURL: NetworkB.HTTPURLs[0],
		}))

		testEnvironment.AddChart(blockscout.New(&blockscout.Props{
			Name:    "source-blockscout",
			WsURL:   NetworkA.URLs[0],
			HttpURL: NetworkA.HTTPURLs[0],
		}))
	*/

	err := testEnvironment.Run()
	require.NoError(t, err)

	if testEnvironment.WillUseRemoteRunner() {
		return c
	}

	err = testEnvironment.
		AddHelm(mockservercfg.New(nil)).
		AddHelm(mockserver.New(nil)).
		AddHelm(chainlink.New(0, clProps)).
		Run()

	require.NoError(t, err)
	return c
}

func AssertBalances(t *testing.T, bas []testhelpers.BalanceAssertion) {
	event := log.Info()
	for _, b := range bas {
		actual := b.Getter(b.Address)
		require.NotNil(t, actual, "%v getter return nil", b.Name)
		if b.Within == "" {
			require.Equal(t, b.Expected, actual.String(), "wrong balance for %s got %s want %s", b.Name, actual, b.Expected)
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
			require.Equal(t, -1, actual.Cmp(high),
				"wrong balance for %s got %s outside expected range [%s, %s]", b.Name, actual, low, high)
			require.Equal(t, 1, actual.Cmp(low),
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

// CCIPDefaultTestSetUp sets up CCIP Lanes as a pre-condition for ccip tests.
// if configureCLNode is set as true, it :
// 1. Deploys K8 env with CL nodes (and simulated geth if NetworkA and NetworkB are simulated).
// 2. Deploys (Reuses) all CCIP contracts for source and destination.
// 3. Creates CCIP jobs on CL node.
// 4. Configures CCIP OCR2 contracts.
//
// if configureCLNode is set as false, it assumes:
// 1. contracts are already deployed on live networks
// 2. CL nodes are set up and configured with existing contracts
// 3. No k8 env deployment is needed
// It reuses already deployed contracts from the addresses provided in ../contracts/ccip/laneconfig/contracts.json
//
// If bidirectional is true it sets up two-way lanes between NetworkA and NetworkB. Same CL nodes are used for both the lanes.
// If bidirectional is false only one way lane is set up.
//
// Returns -
// 1. CCIPLane for NetworkA --> NetworkB
// 2. If bidirectional is true, CCIPLane for NetworkB --> NetworkA
// 3. If configureCLNode is true, the tearDown func to call when environment needs to be destroyed
func CCIPDefaultTestSetUp(
	t *testing.T,
	envName string,
	clProps map[string]interface{},
	transferAmounts []*big.Int,
	numOfCommitNodes int,
	commitAndExecOnSameDON, bidirectional, configureCLNode bool,
) (*CCIPLane, *CCIPLane, func()) {
	parent, cancel := context.WithCancel(context.Background())
	defer cancel()
	var allErrors, err error
	setUpFuncs, ctx := errgroup.WithContext(parent)
	var ccipEnv *CCIPTestEnv
	var k8Env *environment.Environment
	if configureCLNode {
		// deploy the env if configureCLNode is true
		ccipEnv = DeployEnvironments(
			t,
			&environment.Config{
				NamespacePrefix: strings.ToLower(fmt.Sprintf("%s-%s-%s", envName, networkAName, networkBName)),
				Test:            t,
			}, clProps)
		k8Env = ccipEnv.K8Env
		if ccipEnv.K8Env.WillUseRemoteRunner() {
			return nil, nil, nil
		}
	}

	sourceChainClientA2B, err := blockchain.NewEVMClient(NetworkA, k8Env)
	require.NoError(t, err, "Connecting to blockchain nodes shouldn't fail")
	sourceChainClientA2B.ParallelTransactions(true)

	destChainClientA2B, err := blockchain.NewEVMClient(NetworkB, k8Env)
	require.NoError(t, err, "Connecting to blockchain nodes shouldn't fail")
	destChainClientA2B.ParallelTransactions(true)

	// For the reverse lane another set of clients(sourceChainClient,destChainClient)
	// are required with new header subscriptions(otherwise transactions
	// on one lane will keep on waiting for transactions on other lane for the same network)
	// Currently for simulated network clients(from same network) created with NewEVMClient does not sync nonce
	// ConcurrentEVMClient is a work-around for that.
	sourceChainClientB2A, err := blockchain.ConcurrentEVMClient(NetworkB, k8Env, destChainClientA2B)
	require.NoError(t, err, "Connecting to blockchain nodes shouldn't fail")
	sourceChainClientB2A.ParallelTransactions(true)

	destChainClientB2A, err := blockchain.ConcurrentEVMClient(NetworkA, k8Env, sourceChainClientA2B)
	require.NoError(t, err, "Connecting to blockchain nodes shouldn't fail")
	destChainClientB2A.ParallelTransactions(true)

	setUpFuncs.Go(func() error {
		if !configureCLNode {
			return nil
		}
		ccipEnv.clNodeWithKeyReady = make(chan struct{})
		err = ccipEnv.SetUpNodesAndKeys(ctx, big.NewFloat(1), sourceChainClientA2B, destChainClientA2B)
		if err != nil {
			allErrors = multierr.Append(allErrors, fmt.Errorf("setting up nodes and keys shouldn't fail; err -  %+v", err))
		} else {
			// sends value in channel to denote job creation for the lanes can be started
			ccipEnv.clNodeWithKeyReady <- struct{}{}
			if bidirectional {
				ccipEnv.clNodeWithKeyReady <- struct{}{}
			}
		}
		return err
	})

	ccipLaneA2B := &CCIPLane{
		t:                 t,
		TestEnv:           ccipEnv,
		SourceChain:       sourceChainClientA2B,
		DestChain:         destChainClientA2B,
		SourceNetworkName: networkAName,
		DestNetworkName:   networkBName,
		ValidationTimeout: 2 * time.Minute,
		SentReqHashes:     []string{},
		TotalFee:          big.NewInt(0),
		SourceBalances:    make(map[string]*big.Int),
		DestBalances:      make(map[string]*big.Int),
		context:           ctx,
		commonContractsWg: &sync.WaitGroup{},
	}

	var ccipLaneB2A *CCIPLane

	if bidirectional {
		ccipLaneB2A = &CCIPLane{
			t:                 t,
			TestEnv:           ccipEnv,
			SourceNetworkName: networkBName,
			DestNetworkName:   networkAName,
			SourceChain:       sourceChainClientB2A,
			DestChain:         destChainClientB2A,
			ValidationTimeout: 2 * time.Minute,
			SourceBalances:    make(map[string]*big.Int),
			DestBalances:      make(map[string]*big.Int),
			SentReqHashes:     []string{},
			TotalFee:          big.NewInt(0),
			context:           ctx,
			commonContractsWg: &sync.WaitGroup{},
		}
	}

	ccipLaneA2B.commonContractsWg.Add(1)
	setUpFuncs.Go(func() error {
		log.Info().Msg("Setting up lane A to B")
		err := ccipLaneA2B.DeployNewCCIPLane(numOfCommitNodes, commitAndExecOnSameDON, nil, nil,
			transferAmounts, true, configureCLNode)
		if err != nil {
			allErrors = multierr.Append(allErrors, fmt.Errorf("deploying lane A to B; err - %+v", err))
		}
		return err
	})

	if ccipLaneB2A != nil {
		ccipLaneB2A.commonContractsWg.Add(1)
	}

	setUpFuncs.Go(func() error {
		if bidirectional {
			ccipLaneA2B.commonContractsWg.Wait()
			srcCommon := ccipLaneA2B.Dest.Common.CopyAddresses(ccipLaneB2A.context, ccipLaneB2A.SourceChain)
			destCommon := ccipLaneA2B.Source.Common.CopyAddresses(ccipLaneB2A.context, ccipLaneB2A.DestChain)
			log.Info().Msg("Setting up lane B to A")
			err := ccipLaneB2A.DeployNewCCIPLane(numOfCommitNodes, commitAndExecOnSameDON, srcCommon, destCommon,
				transferAmounts, false, configureCLNode)
			if err != nil {
				allErrors = multierr.Append(allErrors, fmt.Errorf("deploying lane A to B; err -  %+v", err))
			}
			return err
		}
		return nil
	})

	tearDown := func() {
		if configureCLNode {
			err := TeardownSuite(t, ccipEnv.K8Env, ctfUtils.ProjectRoot, ccipEnv.CLNodes, nil,
				zapcore.ErrorLevel, ccipLaneA2B.SourceChain, ccipLaneA2B.DestChain)
			require.NoError(t, err, "Environment teardown shouldn't fail")
		}
	}

	errs := make(chan error, 1)
	go func() {
		errs <- setUpFuncs.Wait()
	}()

	// wait for either context to get cancelled or all the error-groups to finish execution
	for {
		select {
		case <-ctx.Done():
			// if the context is cancelled it means there has been some error
			// allErrors helps to print a stacktrace
			require.NoError(t, allErrors)
			return ccipLaneA2B, ccipLaneB2A, tearDown
		case err := <-errs:
			// check if there has been any error while waiting for the error groups
			// to finish execution
			require.NoError(t, err)
			return ccipLaneA2B, ccipLaneB2A, tearDown
		}
	}
}

// CCIPLaneOnExistingDeployment is same as CCIPDefaultTestSetUp
// except it's called when
// 1. contracts are already deployed on live networks
// 2. CL nodes are set up and configured with existing contracts
// 3. No k8 env deployment is needed
// It reuses already deployed contracts from the addresses provided in ../contracts/ccip/laneconfig/contracts.json
// Returns -
// CCIPLane for NetworkA --> NetworkB
// CCIPLane for NetworkB --> NetworkA
func CCIPLaneOnExistingDeployment(
	t *testing.T,
	transferAmounts []*big.Int,
	bidirectional bool,
) (*CCIPLane, *CCIPLane) {
	if NetworkA.Simulated || NetworkB.Simulated {
		t.Fatalf("test cannot run in simulated env")
	}
	ccipLaneA2B, ccipLaneB2A, _ := CCIPDefaultTestSetUp(
		t, "", nil, transferAmounts, 0, false, bidirectional, false)

	return ccipLaneA2B, ccipLaneB2A
}
