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
	. "github.com/onsi/gomega"
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
	"golang.org/x/exp/slices"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_toll_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_ge_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/fee_manager"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/ge_router"
	ccipPlugin "github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/testhelpers"
	"github.com/smartcontractkit/chainlink/core/utils"
	bigmath "github.com/smartcontractkit/chainlink/core/utils/big_math"
	networks "github.com/smartcontractkit/chainlink/integration-tests"
	"github.com/smartcontractkit/chainlink/integration-tests/client"
	"github.com/smartcontractkit/chainlink/integration-tests/contracts"
	"github.com/smartcontractkit/chainlink/integration-tests/contracts/ccip"
)

type BillingModel string

const (
	TOLL BillingModel = "toll"
	GE   BillingModel = "GE"

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
	// TODO dynamic calculation of tollFee for multiple tokens in a msg

	TollFee   = big.NewInt(0).Mul(big.NewInt(12), big.NewInt(1e18)) // "maxCharge":"10784576000000000000" for a msg with two tokens
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
	ChainClient       blockchain.EVMClient
	Deployer          *ccip.CCIPContractsDeployer
	FeeToken          contracts.LinkToken
	FeeTokenPool      *ccip.LockReleaseTokenPool
	BridgeTokens      []contracts.LinkToken // as of now considering the bridge token is same as link token
	TokenPrices       []*big.Int
	BridgeTokenPools  []*ccip.LockReleaseTokenPool
	RateLimiterConfig ccip.RateLimiterConfig
	AFNConfig         ccip.AFNConfig
	AFN               *ccip.AFN
	GERouter          *ccip.GERouter
	deployed          chan struct{}
}

func (ccipModule *CCIPCommon) CopyAddresses(chainClient blockchain.EVMClient) *CCIPCommon {
	var pools []*ccip.LockReleaseTokenPool
	for _, pool := range ccipModule.BridgeTokenPools {
		pools = append(pools, &ccip.LockReleaseTokenPool{EthAddress: pool.EthAddress})
	}
	var tokens []contracts.LinkToken
	for _, token := range ccipModule.BridgeTokens {
		tokens = append(tokens, token)
	}
	return &CCIPCommon{
		ChainClient: chainClient,
		Deployer:    nil,
		FeeToken:    ccipModule.FeeToken,
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
		GERouter: &ccip.GERouter{
			EthAddress: ccipModule.GERouter.EthAddress,
		},
		deployed: make(chan struct{}, 1),
	}
}

// DeployContracts deploys the contracts which are necessary in both source and dest chain
// This reuses common contracts for bidirectional lanes
func (ccipModule *CCIPCommon) DeployContracts(t *testing.T, noOfTokens int) {
	var err error
	cd := ccipModule.Deployer
	if ccipModule.FeeToken == nil {
		// deploy link token
		token, err := cd.DeployLinkTokenContract()
		require.NoError(t, err, "Deploying Link Token Contract shouldn't fail")
		ccipModule.FeeToken = token
		require.NoError(t, ccipModule.ChainClient.WaitForEvents(), "error in waiting for feetoken deployment")
		// token pool for fee token
		ccipModule.FeeTokenPool, err = cd.DeployLockReleaseTokenPoolContract(ccipModule.FeeToken.Address())
		require.NoError(t, err, "Deploying Native TokenPool Contract shouldn't fail")
	} else {
		token, err := cd.NewLinkTokenContract(common.HexToAddress(ccipModule.FeeToken.Address()))
		require.NoError(t, err, "Instantiating Link Token Contract shouldn't fail")
		ccipModule.FeeToken = token
		pool, err := cd.NewLockReleaseTokenPoolContract(ccipModule.FeeTokenPool.EthAddress)
		require.NoError(t, err, "Instantiating Native TokenPool Contract shouldn't fail")
		ccipModule.FeeTokenPool = pool
	}
	if len(ccipModule.BridgeTokens) != noOfTokens {
		// deploy bridge token.
		for i := 0; i < noOfTokens; i++ {
			token, err := cd.DeployLinkTokenContract()
			require.NoError(t, err, "Deploying Link Token Contract shouldn't fail")
			ccipModule.BridgeTokens = append(ccipModule.BridgeTokens, token)
		}

		require.NoError(t, ccipModule.ChainClient.WaitForEvents(), "Error waiting for Link Token deployments")
		// deploy native token pool
		for _, token := range ccipModule.BridgeTokens {
			ntp, err := cd.DeployLockReleaseTokenPoolContract(token.Address())
			require.NoError(t, err, "Deploying Native TokenPool Contract shouldn't fail")
			ccipModule.BridgeTokenPools = append(ccipModule.BridgeTokenPools, ntp)
		}
	} else {
		var tokens []contracts.LinkToken
		for _, token := range ccipModule.BridgeTokens {
			newToken, err := cd.NewLinkTokenContract(common.HexToAddress(token.Address()))
			require.NoError(t, err, "Instantiating Link Token Contract shouldn't fail")
			tokens = append(tokens, newToken)
		}
		ccipModule.BridgeTokens = tokens
		var pools []*ccip.LockReleaseTokenPool
		for _, pool := range ccipModule.BridgeTokenPools {
			newPool, err := cd.NewLockReleaseTokenPoolContract(pool.EthAddress)
			require.NoError(t, err, "Instantiating Native TokenPool Contract shouldn't fail")
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
		require.NoError(t, err, "Deploying AFN Contract shouldn't fail")
	} else {
		afn, err := cd.NewAFNContract(ccipModule.AFN.EthAddress)
		require.NoError(t, err, "Instantiating AFN Contract shouldn't fail")
		ccipModule.AFN = afn
	}
	if ccipModule.GERouter == nil {
		ccipModule.GERouter, err = cd.DeployGERouter([]common.Address{})
		require.NoError(t, err, "Error on GERouter deployment on source chain")
		require.NoError(t, ccipModule.ChainClient.WaitForEvents(), "Error waiting for common contract deployment")
	} else {
		router, err := cd.NewGERouter(ccipModule.GERouter.EthAddress)
		require.NoError(t, err, "Instantiating GERouter Contract shouldn't fail")
		ccipModule.GERouter = router
	}

	ccipModule.deployed <- struct{}{}
	log.Info().Msg("finished deploying common contracts")
}

func DefaultCCIPModule(chainClient blockchain.EVMClient) *CCIPCommon {
	return &CCIPCommon{
		ChainClient: chainClient,
		deployed:    make(chan struct{}, 1),
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
	TollFeeAmount      *big.Int
	DestinationChainId uint64

	// toll
	TollOnRampRouter *ccip.TollOnRampRouter
	TollOnRamp       *ccip.TollOnRamp

	// GE
	GEOnRamp *ccip.GEOnRamp
}

// DeployContracts deploys all CCIP contracts specific to the source chain
func (sourceCCIP *SourceCCIPModule) DeployContracts(t *testing.T, model BillingModel) {
	var err error
	contractDeployer := sourceCCIP.Common.Deployer
	log.Info().Msg("Deploying source chain specific contracts")
	<-sourceCCIP.Common.deployed

	var tokens, pools, bridgeTokens []common.Address
	for _, token := range sourceCCIP.Common.BridgeTokens {
		tokens = append(tokens, common.HexToAddress(token.Address()))
	}
	bridgeTokens = tokens
	tokens = append(tokens, common.HexToAddress(sourceCCIP.Common.FeeToken.Address()))

	for _, pool := range sourceCCIP.Common.BridgeTokenPools {
		pools = append(pools, pool.EthAddress)
	}
	pools = append(pools, sourceCCIP.Common.FeeTokenPool.EthAddress)

	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	require.NoError(t, err, "Error waiting for OnRamp deployment")
	switch model {
	case TOLL:
		// deploy on ramp router
		sourceCCIP.TollOnRampRouter, err = contractDeployer.DeployTollOnRampRouter()
		require.NoError(t, err, "Deploying onramp router shouldn't fail")

		// onRamp
		sourceCCIP.TollOnRamp, err = contractDeployer.DeployTollOnRamp(
			sourceCCIP.Common.ChainClient.GetChainID().Uint64(),
			sourceCCIP.DestinationChainId,
			tokens,
			pools,
			[]common.Address{},
			sourceCCIP.Common.AFN.EthAddress,
			sourceCCIP.TollOnRampRouter.EthAddress,
			sourceCCIP.Common.RateLimiterConfig)
		require.NoError(t, err, "Error on OnRamp deployment")

		err = sourceCCIP.Common.ChainClient.WaitForEvents()
		require.NoError(t, err, "Error waiting for OnRamp deployment")

		// Set bridge token prices on the onRamp
		err = sourceCCIP.TollOnRamp.SetTokenPrices(bridgeTokens, sourceCCIP.Common.TokenPrices)
		require.NoError(t, err, "Setting prices shouldn't fail")

		// update onRampRouter with OnRamp address
		err = sourceCCIP.TollOnRampRouter.SetOnRamp(sourceCCIP.DestinationChainId, sourceCCIP.TollOnRamp.EthAddress)
		require.NoError(t, err, "Error setting onramp on the router")

		// The Fee token to be used in on+off ramp
		sourceCCIP.TollFeeAmount = TollFee

		// set a part of sourceCCIP.TollFeeAmount as onRamp fee rest would be used as offramp
		err = sourceCCIP.TollOnRamp.SetFeeConfig([]common.Address{common.HexToAddress(sourceCCIP.Common.FeeToken.Address())}, []*big.Int{big.NewInt(1)})
		require.NoError(t, err, "setting OnRamp Fee config")

	case GE:
		// GE Set up
		sourceFeeManager, err := contractDeployer.DeployFeeManager([]fee_manager.GEFeeUpdate{
			{
				SourceFeeToken:              common.HexToAddress(sourceCCIP.Common.FeeToken.Address()),
				DestChainId:                 sourceCCIP.DestinationChainId,
				FeeTokenBaseUnitsPerUnitGas: big.NewInt(1e9), // 1 gwei
			},
		})
		require.NoError(t, err, "Error on FeeManager deployment")

		err = sourceCCIP.Common.ChainClient.WaitForEvents()
		require.NoError(t, err, "Error waiting for events")
		sourceCCIP.GEOnRamp, err = contractDeployer.DeployGEOnRamp(
			sourceCCIP.Common.ChainClient.GetChainID().Uint64(),
			sourceCCIP.DestinationChainId,
			tokens,
			pools,
			[]common.Address{},
			sourceCCIP.Common.AFN.EthAddress,
			sourceCCIP.Common.GERouter.EthAddress,
			sourceFeeManager.EthAddress,
			sourceCCIP.Common.RateLimiterConfig,
			[]evm_2_evm_ge_onramp.IEVM2EVMGEOnRampFeeTokenConfigArgs{
				{
					Token:           common.HexToAddress(sourceCCIP.Common.FeeToken.Address()),
					FeeAmount:       big.NewInt(0),
					DestGasOverhead: 0,
					Multiplier:      1e18,
				}})

		require.NoError(t, err, "Error on GEOnRamp deployment")

		err = sourceCCIP.Common.ChainClient.WaitForEvents()
		require.NoError(t, err, "Error waiting for events")

		// Set bridge token prices on the onRamp
		err = sourceCCIP.GEOnRamp.SetTokenPrices(bridgeTokens, sourceCCIP.Common.TokenPrices)
		require.NoError(t, err, "Setting prices shouldn't fail")

		// update onRampRouter with OnRamp address
		err = sourceCCIP.Common.GERouter.SetOnRamp(sourceCCIP.DestinationChainId, sourceCCIP.GEOnRamp.EthAddress)
		require.NoError(t, err, "Error setting onramp on the router")

		err = sourceCCIP.Common.ChainClient.WaitForEvents()
		require.NoError(t, err, "Error waiting for events")
	}

	// update native pool with onRamp address
	for _, pool := range sourceCCIP.Common.BridgeTokenPools {
		if model == TOLL {
			err = pool.SetOnRamp(sourceCCIP.TollOnRamp.EthAddress)
			require.NoError(t, err, "Error setting tollOnramp on the token pool %s", pool.Address())
		}
		if model == GE {
			err = pool.SetOnRamp(sourceCCIP.GEOnRamp.EthAddress)
			require.NoError(t, err, "Error setting GEOnRamp on the token pool %s", pool.Address())
		}
	}
	if model == TOLL {
		err = sourceCCIP.Common.FeeTokenPool.SetOnRamp(sourceCCIP.TollOnRamp.EthAddress)
		require.NoError(t, err, "Error setting TollOnRamp on the token pool %s", sourceCCIP.Common.FeeTokenPool.Address())
	}
	if model == GE {
		err = sourceCCIP.Common.FeeTokenPool.SetOnRamp(sourceCCIP.GEOnRamp.EthAddress)
		require.NoError(t, err, "Error setting GEOnRamp on the token pool %s", sourceCCIP.Common.FeeTokenPool.Address())
	}

	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	require.NoError(t, err, "Error waiting for events")
}

func (sourceCCIP *SourceCCIPModule) CollectBalanceRequirements(t *testing.T, model BillingModel) []testhelpers.BalanceReq {
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
	balancesReq = append(balancesReq, testhelpers.BalanceReq{
		Name:   fmt.Sprintf("%s-FeeToken-%s", testhelpers.Sender, sourceCCIP.Sender.Hex()),
		Addr:   sourceCCIP.Sender,
		Getter: GetterForLinkToken(t, sourceCCIP.Common.FeeToken, sourceCCIP.Sender.Hex()),
	})
	if model == GE {
		balancesReq = append(balancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("%s-GERouter-%s", testhelpers.Sender, sourceCCIP.Common.GERouter.Address()),
			Addr:   sourceCCIP.Common.GERouter.EthAddress,
			Getter: GetterForLinkToken(t, sourceCCIP.Common.FeeToken, sourceCCIP.Common.GERouter.Address()),
		})
		balancesReq = append(balancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("%s-GEOnRamp-%s", testhelpers.Sender, sourceCCIP.GEOnRamp.Address()),
			Addr:   sourceCCIP.GEOnRamp.EthAddress,
			Getter: GetterForLinkToken(t, sourceCCIP.Common.FeeToken, sourceCCIP.GEOnRamp.Address()),
		})
		balancesReq = append(balancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("%s-FeeTokenPool-%s", testhelpers.Sender, sourceCCIP.Common.FeeTokenPool.Address()),
			Addr:   sourceCCIP.Common.FeeTokenPool.EthAddress,
			Getter: GetterForLinkToken(t, sourceCCIP.Common.FeeToken, sourceCCIP.Common.FeeTokenPool.Address()),
		})
	}
	return balancesReq
}

func (sourceCCIP *SourceCCIPModule) BalanceAssertions(t *testing.T, model BillingModel, prevBalances map[string]*big.Int, noOfReq int64, totalGEFee *big.Int) []testhelpers.BalanceAssertion {
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
	switch model {
	case TOLL:
		name := fmt.Sprintf("%s-FeeToken-%s", testhelpers.Sender, sourceCCIP.Sender.Hex())
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     name,
			Address:  sourceCCIP.Sender,
			Getter:   GetterForLinkToken(t, sourceCCIP.Common.FeeToken, sourceCCIP.Sender.Hex()),
			Expected: bigmath.Sub(prevBalances[name], bigmath.Mul(big.NewInt(noOfReq), sourceCCIP.TollFeeAmount)).String(),
		})
	case GE:
		name := fmt.Sprintf("%s-FeeToken-%s", testhelpers.Sender, sourceCCIP.Sender.Hex())
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     name,
			Address:  sourceCCIP.Sender,
			Getter:   GetterForLinkToken(t, sourceCCIP.Common.FeeToken, sourceCCIP.Sender.Hex()),
			Expected: bigmath.Sub(prevBalances[name], totalGEFee).String(),
		})
		name = fmt.Sprintf("%s-FeeTokenPool-%s", testhelpers.Sender, sourceCCIP.Common.FeeTokenPool.Address())
		// TODO add balance assertion for fees going to the feeManager
		//balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
		//	Name:     name,
		//	Address:  sourceCCIP.Common.FeeTokenPool.EthAddress,
		//	Getter:   GetterForLinkToken(t, sourceCCIP.Common.FeeToken, sourceCCIP.Common.FeeTokenPool.Address()),
		//	Expected: bigmath.Add(prevBalances[name], totalGEFee).String(),
		//})
		name = fmt.Sprintf("%s-GERouter-%s", testhelpers.Sender, sourceCCIP.Common.GERouter.Address())
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     fmt.Sprintf("%s-GERouter-%s", testhelpers.Sender, sourceCCIP.Common.GERouter.Address()),
			Address:  sourceCCIP.Common.GERouter.EthAddress,
			Getter:   GetterForLinkToken(t, sourceCCIP.Common.FeeToken, sourceCCIP.Common.GERouter.Address()),
			Expected: prevBalances[name].String(),
		})
		name = fmt.Sprintf("%s-GEOnRamp-%s", testhelpers.Sender, sourceCCIP.GEOnRamp.Address())
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     fmt.Sprintf("%s-GEOnRamp-%s", testhelpers.Sender, sourceCCIP.GEOnRamp.Address()),
			Address:  sourceCCIP.GEOnRamp.EthAddress,
			Getter:   GetterForLinkToken(t, sourceCCIP.Common.FeeToken, sourceCCIP.GEOnRamp.Address()),
			Expected: prevBalances[name].String(),
		})
	}
	return balAssertions
}

func (sourceCCIP *SourceCCIPModule) AssertEventCCIPSendRequested(t *testing.T, model BillingModel, txHash string, currentBlockOnSource uint64, timeout time.Duration) (uint64, [32]byte) {
	log.Info().Msg("Waiting for CCIPSendRequested event")
	var seqNum uint64
	var msgId [32]byte
	gom := NewWithT(t)
	switch model {
	case TOLL:
		gom.Eventually(func(g Gomega) bool {
			iterator, err := sourceCCIP.TollOnRamp.FilterCCIPSendRequested(currentBlockOnSource)
			g.Expect(err).NotTo(HaveOccurred(), "Error filtering CCIPSendRequested event")
			for iterator.Next() {
				if strings.EqualFold(iterator.Event.Raw.TxHash.Hex(), txHash) {
					seqNum = iterator.Event.Message.SequenceNumber
					return true
				}
			}
			return false
		}, timeout, "1s").Should(BeTrue(), "No CCIPSendRequested event found with txHash %s", txHash)
	case GE:
		gom.Eventually(func(g Gomega) bool {
			iterator, err := sourceCCIP.GEOnRamp.FilterCCIPSendRequested(currentBlockOnSource)
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
	}
	return seqNum, msgId
}

func (sourceCCIP *SourceCCIPModule) SendTollRequest(t *testing.T, receiver common.Address, tokenAndAmounts []evm_2_any_toll_onramp_router.CommonEVMTokenAndAmount, data string, feeToken evm_2_any_toll_onramp_router.CommonEVMTokenAndAmount) string {
	receiverAddr, err := utils.ABIEncode(`[{"type":"address"}]`, receiver)
	require.NoError(t, err, "Failed encoding the receiver address")

	extraArgsV1, err := testhelpers.GetEVMExtraArgsV1(big.NewInt(100_000), false)
	require.NoError(t, err, "Failed encoding the options field")

	// form the message for transfer
	msg := evm_2_any_toll_onramp_router.TollConsumerEVM2AnyTollMessage{
		Receiver:          receiverAddr,
		Data:              []byte(data),
		TokensAndAmounts:  tokenAndAmounts,
		FeeTokenAndAmount: feeToken,
		ExtraArgs:         extraArgsV1,
	}
	log.Info().Interface("toll msg details", msg).Msg("ccip message to be sent")

	// initiate the transfer
	sendTx, err := sourceCCIP.TollOnRampRouter.CCIPSend(sourceCCIP.DestinationChainId, msg)
	require.NoError(t, err, "send token should be initiated successfully")
	log.Info().Str("toll send token transaction", sendTx.Hash().String()).Msg("Sending token")
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	require.NoError(t, err, "Failed to wait for events")
	return sendTx.Hash().Hex()
}

func (sourceCCIP *SourceCCIPModule) SendGERequest(
	t *testing.T,
	receiver common.Address,
	tokenAndAmounts []ge_router.CommonEVMTokenAndAmount,
	data string,
	feeToken common.Address,
) (string, *big.Int) {
	receiverAddr, err := utils.ABIEncode(`[{"type":"address"}]`, receiver)
	require.NoError(t, err, "Failed encoding the receiver address")

	extraArgsV1, err := testhelpers.GetEVMExtraArgsV1(big.NewInt(100_000), false)
	require.NoError(t, err, "Failed encoding the options field")

	// form the message for transfer
	msg := ge_router.GEConsumerEVM2AnyGEMessage{
		Receiver:         receiverAddr,
		Data:             []byte(data),
		TokensAndAmounts: tokenAndAmounts,
		FeeToken:         feeToken,
		ExtraArgs:        extraArgsV1,
	}
	log.Info().Interface("ge msg details", msg).Msg("ccip message to be sent")
	fee, err := sourceCCIP.Common.GERouter.GetFee(sourceCCIP.DestinationChainId, msg)
	require.NoError(t, err, "calculating fee")
	log.Info().Int64("fee", fee.Int64()).Msg("calculated fee")

	// Approve the fee amount
	err = sourceCCIP.Common.FeeToken.Approve(sourceCCIP.Common.GERouter.Address(), fee)
	require.NoError(t, err, "approving fee for ge router")
	require.NoError(t, sourceCCIP.Common.ChainClient.WaitForEvents(), "error waiting for events")

	// initiate the transfer
	sendTx, err := sourceCCIP.Common.GERouter.CCIPSend(sourceCCIP.DestinationChainId, msg)
	require.NoError(t, err, "send token should be initiated successfully")

	log.Info().Str("GE send token transaction", sendTx.Hash().String()).Msg("Sending token")
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	require.NoError(t, err, "Failed to wait for events")
	return sendTx.Hash().Hex(), fee
}

func DefaultSourceCCIPModule(t *testing.T, chainClient blockchain.EVMClient, destChain uint64, transferAmount []*big.Int, ccipCommon *CCIPCommon) *SourceCCIPModule {
	sourceCCIP := &SourceCCIPModule{
		Common:             ccipCommon,
		TransferAmount:     transferAmount,
		DestinationChainId: destChain,
		Sender:             common.HexToAddress(chainClient.GetDefaultWallet().Address()),
	}
	var err error
	sourceCCIP.Common.Deployer, err = ccip.NewCCIPContractsDeployer(chainClient)
	require.NoError(t, err, "contract deployer should be created successfully")
	return sourceCCIP
}

type DestCCIPModule struct {
	Common            *CCIPCommon
	SourceChainId     uint64
	CommitStore       *ccip.CommitStore
	TollOffRamp       *ccip.TollOffRamp
	TollOffRampRouter *ccip.TollOffRampRouter
	ReceiverDapp      *ccip.ReceiverDapp
	GEOffRamp         *ccip.GEOffRamp
}

// DeployContracts deploys all CCIP contracts specific to the destination chain
func (destCCIP *DestCCIPModule) DeployContracts(t *testing.T, sourceCCIP SourceCCIPModule, wg *sync.WaitGroup, model BillingModel) {
	var err error
	contractDeployer := destCCIP.Common.Deployer
	log.Info().Msg("Deploying destination chain specific contracts")

	<-destCCIP.Common.deployed
	var onramps []common.Address
	var MinSeqNrByOnRamp []uint64

	if model == TOLL {
		onramps = append(onramps, sourceCCIP.TollOnRamp.EthAddress)
		MinSeqNrByOnRamp = append(MinSeqNrByOnRamp, 1)
	}
	if model == GE {
		onramps = append(onramps, sourceCCIP.GEOnRamp.EthAddress)
		MinSeqNrByOnRamp = append(MinSeqNrByOnRamp, 1)
	}
	// commitStore responsible for validating the transfer message
	destCCIP.CommitStore, err = contractDeployer.DeployCommitStore(
		destCCIP.SourceChainId,
		destCCIP.Common.ChainClient.GetChainID().Uint64(),
		destCCIP.Common.AFN.EthAddress,
		commit_store.ICommitStoreCommitStoreConfig{
			OnRamps:          onramps,
			MinSeqNrByOnRamp: MinSeqNrByOnRamp,
		})
	require.NoError(t, err, "Deploying CommitStore shouldn't fail")
	err = destCCIP.Common.ChainClient.WaitForEvents()
	require.NoError(t, err, "Error waiting for setting destination contracts")

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
		err = token.Transfer(pool.Address(), testhelpers.Link(1000))
		require.NoError(t, err)
	}
	// add the fee token and fee token price for dest
	destTokens = append(destTokens, common.HexToAddress(destCCIP.Common.FeeToken.Address()))
	destCCIP.Common.TokenPrices = append(destCCIP.Common.TokenPrices, big.NewInt(1))

	pools = append(pools, destCCIP.Common.FeeTokenPool.EthAddress)
	err = destCCIP.Common.FeeToken.Transfer(destCCIP.Common.FeeTokenPool.Address(), testhelpers.Link(1000))
	require.NoError(t, err)
	switch model {
	case TOLL:
		// Toll
		// offRamp
		destCCIP.TollOffRamp, err = contractDeployer.DeployTollOffRamp(destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID().Uint64(),
			destCCIP.CommitStore.EthAddress, sourceCCIP.TollOnRamp.EthAddress, destCCIP.Common.AFN.EthAddress,
			sourceTokens, pools, destCCIP.Common.RateLimiterConfig)
		require.NoError(t, err, "Deploying TollOffRamp shouldn't fail")
		err = destCCIP.Common.ChainClient.WaitForEvents()
		require.NoError(t, err, "Error waiting for deploying offramp")
		// Set token prices on the offRamp
		err = destCCIP.TollOffRamp.SetTokenPrices(destTokens, destCCIP.Common.TokenPrices)
		require.NoError(t, err, "Setting prices shouldn't fail")

		// OffRampRouter
		destCCIP.TollOffRampRouter, err = contractDeployer.DeployTollOffRampRouter([]common.Address{destCCIP.TollOffRamp.EthAddress})
		require.NoError(t, err, "Deploying TollOffRampRouter shouldn't fail")
		err = destCCIP.Common.ChainClient.WaitForEvents()
		require.NoError(t, err, "Error waiting for deploying OffRampRouter")
		err = destCCIP.TollOffRamp.SetRouter(destCCIP.TollOffRampRouter.EthAddress)
		require.NoError(t, err, "Error setting router on the offramp")
	case GE:
		// GE
		destFeeManager, err := contractDeployer.DeployFeeManager([]fee_manager.GEFeeUpdate{
			{
				SourceFeeToken:              common.HexToAddress(sourceCCIP.Common.FeeToken.Address()),
				DestChainId:                 destCCIP.SourceChainId,
				FeeTokenBaseUnitsPerUnitGas: big.NewInt(200e9), // (2e20 juels/eth) * (1 gwei / gas) / (1 eth/1e18)
			},
		})
		require.NoError(t, err, "Error on FeeManager deployment")

		err = destCCIP.Common.ChainClient.WaitForEvents()
		require.NoError(t, err, "Error waiting for events on destination contract deployments")

		destCCIP.GEOffRamp, err = contractDeployer.DeployGEOffRamp(destCCIP.SourceChainId, sourceCCIP.DestinationChainId,
			destCCIP.CommitStore.EthAddress, sourceCCIP.GEOnRamp.EthAddress,
			destCCIP.Common.AFN.EthAddress, common.HexToAddress(destCCIP.Common.FeeToken.Address()),
			destFeeManager.EthAddress, sourceTokens, pools, destCCIP.Common.RateLimiterConfig, big.NewInt(0))
		require.NoError(t, err, "Deploying GEOffRamp shouldn't fail")
		err = destCCIP.Common.ChainClient.WaitForEvents()
		require.NoError(t, err, "Error waiting for deploying GEOffRamp")

		// OffRamp can update
		err = destFeeManager.SetFeeUpdater(destCCIP.GEOffRamp.EthAddress)
		require.NoError(t, err, "setting GEOffRamp as fee updater shouldn't fail")

		_, err = destCCIP.Common.GERouter.AddOffRamp(destCCIP.GEOffRamp.EthAddress)
		require.NoError(t, err, "setting GEOffRamp as fee updater shouldn't fail")
		err = destCCIP.Common.ChainClient.WaitForEvents()
		require.NoError(t, err, "Error waiting for events on destination contract deployments")
		err = destCCIP.GEOffRamp.SetRouter(destCCIP.Common.GERouter.EthAddress)
		require.NoError(t, err, "Error setting router on the offramp")

		err = destCCIP.GEOffRamp.SetTokenPrices(destTokens, destCCIP.Common.TokenPrices)
		require.NoError(t, err, "Setting prices shouldn't fail")
	}

	// ReceiverDapp
	destCCIP.ReceiverDapp, err = contractDeployer.DeployReceiverDapp(false)
	require.NoError(t, err, "ReceiverDapp contract should be deployed successfully")

	// update pools with offRamp id
	for _, pool := range destCCIP.Common.BridgeTokenPools {
		if model == TOLL {
			err = pool.SetOffRamp(destCCIP.TollOffRamp.EthAddress)
			require.NoError(t, err, "Error setting TollOffRamp on the token pool %s", pool.Address())
		}
		if model == GE {
			err = pool.SetOffRamp(destCCIP.GEOffRamp.EthAddress)
			require.NoError(t, err, "Error setting GEOffRamp on the token pool %s", pool.Address())
		}
	}
	if model == TOLL {
		err = destCCIP.Common.FeeTokenPool.SetOffRamp(destCCIP.TollOffRamp.EthAddress)
		require.NoError(t, err, "Error setting TollOffRamp on the token pool %s", destCCIP.Common.FeeTokenPool.Address())
	}
	if model == GE {
		err = destCCIP.Common.FeeTokenPool.SetOffRamp(destCCIP.GEOffRamp.EthAddress)
		require.NoError(t, err, "Error setting GEOffRamp on the token pool %s", destCCIP.Common.FeeTokenPool.Address())
	}
	err = destCCIP.Common.ChainClient.WaitForEvents()
	require.NoError(t, err, "Error waiting for events on destination contract deployments")
}

func (destCCIP *DestCCIPModule) CollectBalanceRequirements(t *testing.T, model BillingModel) []testhelpers.BalanceReq {
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
	destBalancesReq = append(destBalancesReq, testhelpers.BalanceReq{
		Name:   fmt.Sprintf("%s-FeeToken-%s", testhelpers.Receiver, destCCIP.ReceiverDapp.Address()),
		Addr:   destCCIP.ReceiverDapp.EthAddress,
		Getter: GetterForLinkToken(t, destCCIP.Common.FeeToken, destCCIP.ReceiverDapp.Address()),
	})
	if model == GE {
		destBalancesReq = append(destBalancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("%s-GEOffRamp-%s", testhelpers.Receiver, destCCIP.GEOffRamp.Address()),
			Addr:   destCCIP.GEOffRamp.EthAddress,
			Getter: GetterForLinkToken(t, destCCIP.Common.FeeToken, destCCIP.GEOffRamp.Address()),
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
	model BillingModel,
	prevBalances map[string]*big.Int,
	transferAmount []*big.Int,
	unusedFee *big.Int,
	noOfReq int64,
	totalGEFee *big.Int,
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
	if model == TOLL {
		name := fmt.Sprintf("%s-FeeToken-%s", testhelpers.Receiver, destCCIP.ReceiverDapp.Address())
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     name,
			Address:  destCCIP.ReceiverDapp.EthAddress,
			Getter:   GetterForLinkToken(t, destCCIP.Common.FeeToken, destCCIP.ReceiverDapp.Address()),
			Expected: bigmath.Add(prevBalances[name], unusedFee.String()).String(),
			Within:   big.NewInt(1e18).String(),
		})
	}

	if model == GE {
		// TODO: FOR GE ALL FEES REMAIN ON SOURCE
		//name := fmt.Sprintf("%s-GEOffRamp-%s", testhelpers.Receiver, destCCIP.GEOffRamp.Address())
		//balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
		//	Name:     name,
		//	Address:  destCCIP.GEOffRamp.EthAddress,
		//	Getter:   GetterForLinkToken(t, destCCIP.Common.FeeToken, destCCIP.GEOffRamp.Address()),
		//	Expected: bigmath.Add(prevBalances[name], totalGEFee).String(),
		//})
		//name = fmt.Sprintf("%s-FeeTokenPool-%s", testhelpers.Receiver, destCCIP.Common.FeeTokenPool.Address())
		//balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
		//	Name:     name,
		//	Address:  destCCIP.Common.FeeTokenPool.EthAddress,
		//	Getter:   GetterForLinkToken(t, destCCIP.Common.FeeToken, destCCIP.Common.FeeTokenPool.Address()),
		//	Expected: bigmath.Sub(prevBalances[name], totalGEFee).String(),
		//})
	}
	return balAssertions
}

func (destCCIP *DestCCIPModule) AssertEventExecutionStateChanged(t *testing.T, model BillingModel, seqNum uint64, msgID [32]byte, currentBlockOnDest uint64, timeout time.Duration) {
	log.Info().Int64("seqNum", int64(seqNum)).Msg("Waiting for ExecutionStateChanged event")
	gom := NewWithT(t)
	switch model {
	case TOLL:
		gom.Eventually(func(g Gomega) ccipPlugin.MessageExecutionState {
			iterator, err := destCCIP.TollOffRamp.FilterExecutionStateChanged([]uint64{seqNum}, currentBlockOnDest)
			g.Expect(err).NotTo(HaveOccurred(), "Error filtering ExecutionStateChanged event for seqNum %d", seqNum)
			g.Expect(iterator.Next()).To(BeTrue(), "No ExecutionStateChanged event found for seqNum %d", seqNum)
			return ccipPlugin.MessageExecutionState(iterator.Event.State)
		}, timeout, "1s").Should(Equal(ccipPlugin.Success))
	case GE:
		gom.Eventually(func(g Gomega) ccipPlugin.MessageExecutionState {
			iterator, err := destCCIP.GEOffRamp.FilterExecutionStateChanged([]uint64{seqNum}, [][32]byte{msgID}, currentBlockOnDest)
			g.Expect(err).NotTo(HaveOccurred(), "Error filtering ExecutionStateChanged event for seqNum %d", seqNum)
			g.Expect(iterator.Next()).To(BeTrue(), "No ExecutionStateChanged event found for seqNum %d", seqNum)
			return ccipPlugin.MessageExecutionState(iterator.Event.State)
		}, timeout, "1s").Should(Equal(ccipPlugin.Success))
	}
}

func (destCCIP *DestCCIPModule) AssertEventReportAccepted(t *testing.T, onRamp common.Address, seqNum, currentBlockOnDest uint64, timeout time.Duration) {
	log.Info().Int64("seqNum", int64(seqNum)).Msg("Waiting for ReportAccepted event")
	gom := NewWithT(t)
	gom.Eventually(func(g Gomega) bool {
		iterator, err := destCCIP.CommitStore.FilterReportAccepted(currentBlockOnDest)
		g.Expect(err).NotTo(HaveOccurred(), "Error filtering ReportAccepted event")
		for iterator.Next() {
			if slices.Contains(iterator.Event.Report.OnRamps, onRamp) {
				for _, ints := range iterator.Event.Report.Intervals {
					if ints.Min <= seqNum && ints.Max >= seqNum {
						return true
					}
				}
			}
		}
		return false
	}, timeout, "1s").Should(BeTrue(), "No ReportAccepted Event found for onRamp %s and seq num %d", onRamp.Hex(), seqNum)
}

func (destCCIP *DestCCIPModule) AssertSeqNumberExecuted(t *testing.T, onRamp common.Address, seqNumberBefore uint64, timeout time.Duration) {
	log.Info().Int64("seqNum", int64(seqNumberBefore)).Msg("Waiting to be executed")
	gom := NewWithT(t)
	gom.Eventually(func(g Gomega) bool {
		seqNumberAfter, err := destCCIP.CommitStore.GetNextSeqNumber(onRamp)
		if err != nil {
			return false
		}
		if seqNumberAfter > seqNumberBefore {
			return true
		}
		return false
	}, timeout, "1s").Should(BeTrue(), "Error Executing Sequence number %d", seqNumberBefore)
}

func DefaultDestinationCCIPModule(t *testing.T, chainClient blockchain.EVMClient, sourceChain uint64, ccipCommon *CCIPCommon) *DestCCIPModule {
	destCCIP := &DestCCIPModule{
		Common:        ccipCommon,
		SourceChainId: sourceChain,
	}
	var err error
	destCCIP.Common.Deployer, err = ccip.NewCCIPContractsDeployer(chainClient)
	require.NoError(t, err, "contract deployer should be created successfully")
	return destCCIP
}

type CCIPLane struct {
	t                       *testing.T
	SourceNetworkName       string
	DestNetworkName         string
	Source                  *SourceCCIPModule
	Dest                    *DestCCIPModule
	TestEnv                 *CCIPTestEnv
	Ready                   chan struct{}
	commonContractsDeployed chan struct{}
	NumberOfTollReq         int
	NumberOfGEReq           int
	SourceBalances          map[string]*big.Int
	DestBalances            map[string]*big.Int
	StartBlockOnSource      uint64
	StartBlockOnDestination uint64
	SentTollReqHashes       []string
	SentGEReqHashes         []string
	TotalGEFee              *big.Int
	ValidationTimeout       time.Duration
}

func (lane *CCIPLane) IsLaneDeployed() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	for {
		select {
		case <-lane.Ready:
			return nil
		case <-ctx.Done():
			return fmt.Errorf("waited too long for the lane set up")
		}
	}
}

func (lane *CCIPLane) RecordStateBeforeTollTransfer() {
	var err error
	// collect the balance requirement to verify balances after transfer
	lane.SourceBalances, err = testhelpers.GetBalances(lane.Source.CollectBalanceRequirements(lane.t, TOLL))
	require.NoError(lane.t, err, "fetching source balance")
	lane.DestBalances, err = testhelpers.GetBalances(lane.Dest.CollectBalanceRequirements(lane.t, TOLL))
	require.NoError(lane.t, err, "fetching dest balance")

	// save the current block numbers to use in various filter log requests
	lane.StartBlockOnSource, err = lane.Source.Common.ChainClient.LatestBlockNumber(context.Background())
	require.NoError(lane.t, err, "Getting current block should be successful in source chain")
	lane.StartBlockOnDestination, err = lane.Dest.Common.ChainClient.LatestBlockNumber(context.Background())
	require.NoError(lane.t, err, "Getting current block should be successful in dest chain")
}

func (lane *CCIPLane) SendTollRequests(noOfRequests int) {
	t := lane.t
	lane.NumberOfTollReq += noOfRequests
	var tokenAndAmounts []evm_2_any_toll_onramp_router.CommonEVMTokenAndAmount
	for i, token := range lane.Source.Common.BridgeTokens {
		tokenAndAmounts = append(tokenAndAmounts, evm_2_any_toll_onramp_router.CommonEVMTokenAndAmount{
			Token: common.HexToAddress(token.Address()), Amount: lane.Source.TransferAmount[i],
		})
		// approve the onramp router so that it can initiate transferring the token
		err := token.Approve(lane.Source.TollOnRampRouter.Address(), bigmath.Mul(lane.Source.TransferAmount[i], big.NewInt(int64(noOfRequests))))
		require.NoError(t, err, "Could not approve permissions for the onRamp router "+
			"on the source link token contract")
	}
	// approve onramp router to use the fee token
	err := lane.Source.Common.FeeToken.Approve(lane.Source.TollOnRampRouter.Address(), bigmath.Mul(lane.Source.TollFeeAmount, big.NewInt(int64(noOfRequests))))
	require.NoError(t, err, "Could not approve permissions for the onRamp router "+
		"on the source link token contract")
	err = lane.Source.Common.ChainClient.WaitForEvents()
	require.NoError(t, err, "Failed to wait for events")

	for i := 1; i <= noOfRequests; i++ {
		txHash := lane.Source.SendTollRequest(
			t, lane.Dest.ReceiverDapp.EthAddress,
			tokenAndAmounts,
			fmt.Sprintf("msg %d", i),
			evm_2_any_toll_onramp_router.CommonEVMTokenAndAmount{
				Token: common.HexToAddress(lane.Source.Common.FeeToken.Address()), Amount: lane.Source.TollFeeAmount,
			},
		)
		lane.SentTollReqHashes = append(lane.SentTollReqHashes, txHash)
	}
}

func (lane *CCIPLane) ValidateTollRequests() {
	t := lane.t
	for _, txHash := range lane.SentTollReqHashes {
		// Verify if
		// - CCIPSendRequested Event log generated,
		// - NextSeqNumber from commitStore got increased
		seqNumber, _ := lane.Source.AssertEventCCIPSendRequested(t, TOLL, txHash, lane.StartBlockOnSource, lane.ValidationTimeout)
		lane.Dest.AssertSeqNumberExecuted(t, lane.Source.TollOnRamp.EthAddress, seqNumber, lane.ValidationTimeout)

		// Verify whether commitStore has accepted the report
		lane.Dest.AssertEventReportAccepted(t, lane.Source.TollOnRamp.EthAddress, seqNumber, lane.StartBlockOnDestination, lane.ValidationTimeout)

		// Verify whether the execution state is changed and the transfer is successful
		lane.Dest.AssertEventExecutionStateChanged(t, TOLL, seqNumber, [32]byte{}, lane.StartBlockOnDestination, lane.ValidationTimeout)
	}
	// verify the fee amount is deducted from sender, added to receiver token balances and
	// unused fee is returned to receiver fee token account
	AssertBalances(lane.t, lane.Source.BalanceAssertions(lane.t, TOLL, lane.SourceBalances, int64(lane.NumberOfTollReq), lane.TotalGEFee))
	AssertBalances(lane.t, lane.Dest.BalanceAssertions(lane.t, TOLL, lane.DestBalances, lane.Source.TransferAmount, UnusedFee, int64(lane.NumberOfTollReq), lane.TotalGEFee))

}

func (lane *CCIPLane) RecordStateBeforeGETransfer() {
	var err error
	// collect the balance requirement to verify balances after transfer
	lane.SourceBalances, err = testhelpers.GetBalances(lane.Source.CollectBalanceRequirements(lane.t, GE))
	require.NoError(lane.t, err, "fetching source balance")
	lane.DestBalances, err = testhelpers.GetBalances(lane.Dest.CollectBalanceRequirements(lane.t, GE))
	require.NoError(lane.t, err, "fetching dest balance")

	// save the current block numbers to use in various filter log requests
	lane.StartBlockOnSource, err = lane.Source.Common.ChainClient.LatestBlockNumber(context.Background())
	require.NoError(lane.t, err, "Getting current block should be successful in source chain")
	lane.StartBlockOnDestination, err = lane.Dest.Common.ChainClient.LatestBlockNumber(context.Background())
	require.NoError(lane.t, err, "Getting current block should be successful in dest chain")
}

func (lane *CCIPLane) SendGERequests(noOfRequests int) []string {
	t := lane.t
	lane.NumberOfGEReq += noOfRequests
	var tokenAndAmounts []ge_router.CommonEVMTokenAndAmount
	for i, token := range lane.Source.Common.BridgeTokens {
		tokenAndAmounts = append(tokenAndAmounts, ge_router.CommonEVMTokenAndAmount{
			Token: common.HexToAddress(token.Address()), Amount: lane.Source.TransferAmount[i],
		})
		// approve the onramp router so that it can initiate transferring the token

		err := token.Approve(lane.Source.Common.GERouter.Address(), bigmath.Mul(lane.Source.TransferAmount[i], big.NewInt(int64(noOfRequests))))
		require.NoError(t, err, "Could not approve permissions for the onRamp router "+
			"on the source link token contract")
	}

	err := lane.Source.Common.ChainClient.WaitForEvents()
	require.NoError(t, err, "Failed to wait for events")
	var txs []string
	for i := 1; i <= noOfRequests; i++ {
		txHash, fee := lane.Source.SendGERequest(
			t, lane.Dest.ReceiverDapp.EthAddress,
			tokenAndAmounts,
			fmt.Sprintf("msg %d", i),
			common.HexToAddress(lane.Source.Common.FeeToken.Address()),
		)
		lane.SentGEReqHashes = append(lane.SentGEReqHashes, txHash)
		txs = append(txs, txHash)
		lane.TotalGEFee = bigmath.Add(lane.TotalGEFee, fee)
	}
	return txs
}

func (lane *CCIPLane) ValidateGERequests() {
	for _, txHash := range lane.SentGEReqHashes {
		lane.ValidateGERequestByTxHash(txHash)
	}
	// verify the fee amount is deducted from sender, added to receiver token balances and
	// unused fee is returned to receiver fee token account
	AssertBalances(lane.t, lane.Source.BalanceAssertions(lane.t, GE, lane.SourceBalances, int64(lane.NumberOfGEReq), lane.TotalGEFee))
	AssertBalances(lane.t, lane.Dest.BalanceAssertions(lane.t, GE, lane.DestBalances, lane.Source.TransferAmount, UnusedFee, int64(lane.NumberOfGEReq), lane.TotalGEFee))

}

func (lane *CCIPLane) ValidateGERequestByTxHash(txHash string) {
	t := lane.t
	// Verify if
	// - CCIPSendRequested Event log generated,
	// - NextSeqNumber from commitStore got increased
	seqNumber, msgId := lane.Source.AssertEventCCIPSendRequested(t, GE, txHash, lane.StartBlockOnSource, lane.ValidationTimeout)
	lane.Dest.AssertSeqNumberExecuted(t, lane.Source.GEOnRamp.EthAddress, seqNumber, lane.ValidationTimeout)

	// Verify whether commitStore has accepted the report
	lane.Dest.AssertEventReportAccepted(t, lane.Source.GEOnRamp.EthAddress, seqNumber, lane.StartBlockOnDestination, lane.ValidationTimeout)

	// Verify whether the execution state is changed and the transfer is successful
	lane.Dest.AssertEventExecutionStateChanged(t, GE, seqNumber, msgId, lane.StartBlockOnDestination, lane.ValidationTimeout)
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
	lane.RecordStateBeforeGETransfer()
	for {
		select {
		case <-ticker.C:
			if timeout {
				break
			}
			numOfReq++
			wg.Add(1)
			log.Info().Int("Req No", numOfReq).Msgf("Token transfer with GE for lane %s --> %s", lane.SourceNetworkName, lane.DestNetworkName)
			txs := lane.SendGERequests(1)
			require.NotEmpty(t, txs)
			go func(txHash string) {
				defer wg.Done()
				lane.ValidateGERequestByTxHash(txHash)
				reqSuccess++
			}(txs[0])
		case <-ctx.Done():
			log.Warn().Msgf("Soak Test duration completed for lane %s --> %s. Completing validation for triggered requests", lane.SourceNetworkName, lane.DestNetworkName)
			timeout = true
			wg.Wait()
			return numOfReq, reqSuccess
		}
	}
}

// SetOCR2Configs sets the oracle config in ocr2 contracts
// nil value in execNodes denotes commit and execution jobs are to be set up in same DON
func SetOCR2Configs(t *testing.T, commitNodes, execNodes []*client.CLNodesWithKeys, destCCIP DestCCIPModule) {
	signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, err :=
		ccip.NewOffChainAggregatorV2Config(commitNodes)
	require.NoError(t, err, "Shouldn't fail while getting the config values for ocr2 type contract")
	err = destCCIP.CommitStore.SetOCR2Config(signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
	require.NoError(t, err, "Shouldn't fail while setting commitStore config")
	// if commit and exec job is set up in different DON
	if len(execNodes) > 0 {
		signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, err =
			ccip.NewOffChainAggregatorV2Config(execNodes)
		require.NoError(t, err, "Shouldn't fail while getting the config values for ocr2 type contract")
	}
	if destCCIP.TollOffRamp != nil {
		err = destCCIP.TollOffRamp.SetOCR2Config(signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
		require.NoError(t, err, "Shouldn't fail while setting TollOffRamp config")
	}
	if destCCIP.GEOffRamp != nil {
		err = destCCIP.GEOffRamp.SetOCR2Config(signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
		require.NoError(t, err, "Shouldn't fail while setting GEOffRamp config")
	}
	err = destCCIP.Common.ChainClient.WaitForEvents()
	require.NoError(t, err, "Shouldn't fail while waiting for events on setting ocr2 config")
}

// CreateOCRJobsForCCIP bootstraps the first node and to the other nodes sends ocr jobs that
// sets up ccip-commit and ccip-execution plugin
// nil value in bootstrapExec and execNodes denotes commit and execution jobs are to be set up in same DON
func CreateOCRJobsForCCIP(
	t *testing.T,
	bootstrapCommit *client.CLNodesWithKeys,
	bootstrapExec *client.CLNodesWithKeys,
	commitNodes, execNodes []*client.CLNodesWithKeys,
	tollOnRamp, geOnRamp, commitStore, tollOffRamp, geOffRamp common.Address,
	sourceChainClient, destChainClient blockchain.EVMClient,
	linkTokenAddr []string,
	mockServer *ctfClient.MockserverClient,
	newBootStrap bool,
) {
	bootstrapCommitP2PIds := bootstrapCommit.KeysBundle.P2PKeys
	bootstrapCommitP2PId := bootstrapCommitP2PIds.Data[0].Attributes.PeerID
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
	require.NoError(t, err, "Getting current block should be successful in source chain")
	currentBlockOnDest, err := destChainClient.LatestBlockNumber(context.Background())
	require.NoError(t, err, "Getting current block should be successful in dest chain")
	var onRamps []common.Address
	if tollOnRamp != common.HexToAddress("0x0") {
		onRamps = append(onRamps, tollOnRamp)
	}
	if geOnRamp != common.HexToAddress("0x0") {
		onRamps = append(onRamps, geOnRamp)
	}
	jobParams := testhelpers.CCIPJobSpecParams{
		OnRampsOnCommit:    onRamps,
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
		require.NoError(t, err, "Shouldn't fail creating bootstrap job on bootstrap node")
		if bootstrapExec != nil && len(execNodes) > 0 {
			_, err := bootstrapExec.Node.MustCreateJob(jobParams.BootstrapJob(geOffRamp.Hex()))
			require.NoError(t, err, "Shouldn't fail creating bootstrap job on bootstrap node")
		}
	}

	if len(execNodes) == 0 {
		execNodes = commitNodes
	}

	tokenFeeConv := make(map[string]interface{})
	for _, token := range linkTokenAddr {
		tokenFeeConv[token] = "200000000000000000000"
	}
	SetMockServerWithSameTokenFeeConversionValue(t, tokenFeeConv, execNodes, mockServer)

	ocr2SpecCommit, err := jobParams.CommitJobSpec()
	require.NoError(t, err)
	var ocr2SpecExecForToll, ocr2SpecExecForGE *client.OCR2TaskJobSpec
	if tollOffRamp != common.HexToAddress("0x0") {
		jobParams.OffRamp = tollOffRamp
		jobParams.OnRampForExecution = tollOnRamp
		jobParams.P2PV2Bootstrappers = []string{p2pBootstrappersExec.P2PV2Bootstrapper()}
		ocr2SpecExecForToll, err = jobParams.ExecutionJobSpec()
		require.NoError(t, err)
		ocr2SpecExecForToll.Name = fmt.Sprintf("%s-toll", ocr2SpecExecForToll.Name)
	}

	if geOffRamp != common.HexToAddress("0x0") {
		jobParams.OffRamp = geOffRamp
		jobParams.OnRampForExecution = geOnRamp
		jobParams.P2PV2Bootstrappers = []string{p2pBootstrappersExec.P2PV2Bootstrapper()}
		ocr2SpecExecForGE, err = jobParams.ExecutionJobSpec()
		require.NoError(t, err)
		ocr2SpecExecForGE.Name = fmt.Sprintf("%s-ge", ocr2SpecExecForGE.Name)
	}

	for nodeIndex := 0; nodeIndex < len(commitNodes); nodeIndex++ {
		nodeTransmitterAddress := commitNodes[nodeIndex].KeysBundle.EthAddress
		nodeOCR2Key := commitNodes[nodeIndex].KeysBundle.OCR2Key
		nodeOCR2KeyId := nodeOCR2Key.Data.ID
		ocr2SpecCommit.OCR2OracleSpec.OCRKeyBundleID.SetValid(nodeOCR2KeyId)
		ocr2SpecCommit.OCR2OracleSpec.TransmitterID.SetValid(nodeTransmitterAddress)

		_, err = commitNodes[nodeIndex].Node.MustCreateJob(ocr2SpecCommit)
		require.NoError(t, err, "Shouldn't fail creating CCIP-Commit OCR Task job on OCR node %d job name %s",
			nodeIndex+1, ocr2SpecCommit.Name)
	}

	for nodeIndex := 0; nodeIndex < len(execNodes); nodeIndex++ {
		tokensPerFeeCoinPipeline := TokenFeeForMultipleTokenAddr(execNodes[nodeIndex], linkTokenAddr, mockServer)
		nodeTransmitterAddress := execNodes[nodeIndex].KeysBundle.EthAddress
		nodeOCR2Key := execNodes[nodeIndex].KeysBundle.OCR2Key
		nodeOCR2KeyId := nodeOCR2Key.Data.ID
		if ocr2SpecExecForToll != nil {
			ocr2SpecExecForToll.OCR2OracleSpec.PluginConfig["tokensPerFeeCoinPipeline"] = fmt.Sprintf(`"""
%s
"""`, tokensPerFeeCoinPipeline)
			ocr2SpecExecForToll.OCR2OracleSpec.OCRKeyBundleID.SetValid(nodeOCR2KeyId)
			ocr2SpecExecForToll.OCR2OracleSpec.TransmitterID.SetValid(nodeTransmitterAddress)

			_, err = execNodes[nodeIndex].Node.MustCreateJob(ocr2SpecExecForToll)
			require.NoError(t, err, "Shouldn't fail creating CCIP-Exec-toll OCR Task job on OCR node %d job name %s",
				nodeIndex+1, ocr2SpecExecForToll.Name)
		}

		if ocr2SpecExecForGE != nil {
			ocr2SpecExecForGE.OCR2OracleSpec.PluginConfig["tokensPerFeeCoinPipeline"] = fmt.Sprintf(`"""
%s
"""`, tokensPerFeeCoinPipeline)
			ocr2SpecExecForGE.OCR2OracleSpec.OCRKeyBundleID.SetValid(nodeOCR2KeyId)
			ocr2SpecExecForGE.OCR2OracleSpec.TransmitterID.SetValid(nodeTransmitterAddress)

			_, err = execNodes[nodeIndex].Node.MustCreateJob(ocr2SpecExecForGE)
			require.NoError(t, err, "Shouldn't fail creating CCIP-Exec-ge OCR Task job on OCR node %d job name %s",
				nodeIndex+1, ocr2SpecExecForGE.Name)
		}
	}
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
	t *testing.T,
	tokenValueAddress map[string]interface{},
	chainlinkNodes []*client.CLNodesWithKeys,
	mockserver *ctfClient.MockserverClient,
) {
	var valueAdditions sync.WaitGroup
	for tokenAddr, value := range tokenValueAddress {
		for _, n := range chainlinkNodes {
			valueAdditions.Add(1)
			nodeTokenPairID := nodeContractPair(n.KeysBundle.EthAddress, tokenAddr)
			path := fmt.Sprintf("/%s", nodeTokenPairID)
			go func(path string) {
				defer valueAdditions.Done()
				err := mockserver.SetAnyValuePath(path, value)
				require.NoError(t, err, "Setting mockserver value path shouldn't fail")
			}(path)
		}
	}
	valueAdditions.Wait()
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
	SourceChainClient        blockchain.EVMClient
	DestChainClient          blockchain.EVMClient
	K8Env                    *environment.Environment
}

func (c CCIPTestEnv) ChaosLabel(t *testing.T) {
	for i := c.commitNodeStartIndex; i < len(c.CLNodes); i++ {
		labelSelector := map[string]string{
			"app":      "chainlink-0",
			"instance": strconv.Itoa(i),
		}
		// commit node starts from index 2
		if i >= c.commitNodeStartIndex && i < c.commitNodeStartIndex+c.numOfAllowedFaultyCommit+1 {
			err := c.K8Env.Client.LabelChaosGroupByLabels(c.K8Env.Cfg.Namespace, labelSelector, ChaosGroupCommitFaultyPlus)
			require.NoError(t, err)
		}
		if i >= c.commitNodeStartIndex && i < c.commitNodeStartIndex+c.numOfAllowedFaultyCommit {
			err := c.K8Env.Client.LabelChaosGroupByLabels(c.K8Env.Cfg.Namespace, labelSelector, ChaosGroupCommitFaulty)
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

func DeployEnvironments(
	t *testing.T,
	envconfig *environment.Config,
	clProps map[string]interface{},
) *environment.Environment {
	testEnvironment := environment.New(envconfig)

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
							"replicas": "1",
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
							"replicas": "1",
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
	testEnvironment.
		AddHelm(mockservercfg.New(nil)).
		AddHelm(mockserver.New(nil))

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
		return testEnvironment
	}

	err = testEnvironment.AddHelm(chainlink.New(0, clProps)).Run()
	require.NoError(t, err)

	return testEnvironment
}

func SetUpNodesAndKeys(
	t *testing.T,
	testEnvironment *environment.Environment,
	nodeFund *big.Float,
) CCIPTestEnv {
	log.Info().Msg("Connecting to launched resources")

	sourceChainClient, err := blockchain.NewEVMClient(NetworkA, testEnvironment)
	require.NoError(t, err, "Connecting to blockchain nodes shouldn't fail")
	sourceChainClient.ParallelTransactions(true)

	destChainClient, err := blockchain.NewEVMClient(NetworkB, testEnvironment)
	require.NoError(t, err, "Connecting to blockchain nodes shouldn't fail")
	destChainClient.ParallelTransactions(true)

	chainlinkNodes, err := client.ConnectChainlinkNodes(testEnvironment)
	require.NoError(t, err, "Connecting to chainlink nodes shouldn't fail")
	require.Greater(t, len(chainlinkNodes), 0, "No CL node found")

	mockServer, err := ctfClient.ConnectMockServer(testEnvironment)
	require.NoError(t, err, "Creating mockserver clients shouldn't fail")

	nodesWithKeys := make(map[string][]*client.CLNodesWithKeys)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	log.Info().Msg("creating node keys")
	populateKeys := func(chain blockchain.EVMClient) {
		defer wg.Done()
		_, clNodes, err := client.CreateNodeKeysBundle(chainlinkNodes, "evm", chain.GetChainID().String())
		require.NoError(t, err)
		require.Greater(t, len(clNodes), 0, "No CL node with keys found")
		mu.Lock()
		defer mu.Unlock()
		nodesWithKeys[chain.GetChainID().String()] = clNodes
	}

	wg.Add(1)
	go populateKeys(destChainClient)

	wg.Add(1)
	go populateKeys(sourceChainClient)

	log.Info().Msg("Funding Chainlink nodes for both the chains")

	fund := func(chain blockchain.EVMClient) {
		defer wg.Done()
		err = FundChainlinkNodesAddresses(chainlinkNodes[1:], chain, nodeFund)
		require.NoError(t, err)
	}

	wg.Add(1)
	go fund(sourceChainClient)

	wg.Add(1)
	go fund(destChainClient)

	wg.Wait()

	return CCIPTestEnv{
		MockServer:        mockServer,
		CLNodesWithKeys:   nodesWithKeys,
		CLNodes:           chainlinkNodes,
		SourceChainClient: sourceChainClient,
		DestChainClient:   destChainClient,
		K8Env:             testEnvironment,
	}
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

func GetterForLinkToken(t *testing.T, token contracts.LinkToken, addr string) func(_ common.Address) *big.Int {
	return func(_ common.Address) *big.Int {
		balance, err := token.BalanceOf(context.Background(), addr)
		require.NoError(t, err)
		return balance
	}
}

func CCIPDefaultTestSetUp(
	t *testing.T,
	envName string,
	clProps map[string]interface{},
	transferAmounts []*big.Int,
	numOfCommitNodes int,
	commitAndExecOnSameDON bool,
	bidirectional bool,
	model BillingModel,
) (*CCIPLane, *CCIPLane, func()) {
	testEnvironment := DeployEnvironments(
		t,
		&environment.Config{
			NamespacePrefix: strings.ToLower(fmt.Sprintf("%s-%s-%s-%s", envName, model, networkAName, networkBName)),
			Test:            t,
		}, clProps)

	if testEnvironment.WillUseRemoteRunner() {
		return nil, nil, nil
	}
	testSetUpA2B := SetUpNodesAndKeys(t, testEnvironment, big.NewFloat(1))

	sourceChainClient, err := blockchain.ConcurrentEVMClient(NetworkB, testEnvironment, testSetUpA2B.DestChainClient)
	require.NoError(t, err, "Connecting to blockchain nodes shouldn't fail")
	sourceChainClient.ParallelTransactions(true)

	destChainClient, err := blockchain.ConcurrentEVMClient(NetworkA, testEnvironment, testSetUpA2B.SourceChainClient)
	require.NoError(t, err, "Connecting to blockchain nodes shouldn't fail")
	destChainClient.ParallelTransactions(true)

	testSetUpB2A := CCIPTestEnv{
		MockServer:        testSetUpA2B.MockServer,
		CLNodesWithKeys:   testSetUpA2B.CLNodesWithKeys,
		CLNodes:           testSetUpA2B.CLNodes,
		SourceChainClient: sourceChainClient,
		DestChainClient:   destChainClient,
		K8Env:             testEnvironment,
	}

	ccipLaneA2B := &CCIPLane{
		t:                       t,
		TestEnv:                 &testSetUpA2B,
		SourceNetworkName:       networkAName,
		DestNetworkName:         networkBName,
		ValidationTimeout:       2 * time.Minute,
		Ready:                   make(chan struct{}, 1),
		commonContractsDeployed: make(chan struct{}, 1),
		SentGEReqHashes:         []string{},
		SentTollReqHashes:       []string{},
		TotalGEFee:              big.NewInt(0),
	}

	ccipLaneB2A := &CCIPLane{
		t:                       t,
		TestEnv:                 &testSetUpB2A,
		SourceNetworkName:       networkBName,
		DestNetworkName:         networkAName,
		ValidationTimeout:       2 * time.Minute,
		Ready:                   make(chan struct{}, 1),
		commonContractsDeployed: make(chan struct{}, 1),
		SourceBalances:          make(map[string]*big.Int),
		DestBalances:            make(map[string]*big.Int),
		SentGEReqHashes:         []string{},
		SentTollReqHashes:       []string{},
		TotalGEFee:              big.NewInt(0),
	}
	// This WaitGroup is for waiting on the deployment of common contracts for lane A to B
	// so that these contracts can be reused for lane B to A
	// one group is added for each chain
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		log.Info().Msg("Setting up lane A to B")
		ccipLaneA2B.DeployNewCCIPLane(numOfCommitNodes, commitAndExecOnSameDON, nil, nil,
			transferAmounts, true, wg, model)
	}()

	wg.Wait()
	go func() {
		if bidirectional {
			srcCommon := ccipLaneA2B.Dest.Common.CopyAddresses(testSetUpB2A.SourceChainClient)
			destCommon := ccipLaneA2B.Source.Common.CopyAddresses(testSetUpB2A.DestChainClient)
			log.Info().Msg("Setting up lane B to A")
			ccipLaneB2A.DeployNewCCIPLane(numOfCommitNodes, commitAndExecOnSameDON, srcCommon, destCommon,
				transferAmounts, false, nil, model)
		}
	}()

	tearDown := func() {
		err := TeardownSuite(t, testEnvironment, ctfUtils.ProjectRoot, testSetUpA2B.CLNodes, nil,
			testSetUpA2B.SourceChainClient, testSetUpA2B.DestChainClient)
		require.NoError(t, err, "Environment teardown shouldn't fail")
	}

	return ccipLaneA2B, ccipLaneB2A, tearDown
}

func (lane *CCIPLane) DeployNewCCIPLane(
	numOfCommitNodes int,
	commitAndExecOnSameDON bool,
	sourceCommon *CCIPCommon,
	destCommon *CCIPCommon,
	transferAmounts []*big.Int,
	newBootstrap bool,
	wg *sync.WaitGroup,
	model BillingModel,
) {
	env := lane.TestEnv
	sourceChainClient := env.SourceChainClient
	destChainClient := env.DestChainClient
	clNodesWithKeys := env.CLNodesWithKeys
	mockServer := env.MockServer
	t := lane.t
	// deploy all source contracts
	if sourceCommon == nil {
		sourceCommon = DefaultCCIPModule(sourceChainClient)
	}

	if destCommon == nil {
		destCommon = DefaultCCIPModule(destChainClient)
	}

	lane.Source = DefaultSourceCCIPModule(t, sourceChainClient, destChainClient.GetChainID().Uint64(), transferAmounts, sourceCommon)
	lane.Dest = DefaultDestinationCCIPModule(t, destChainClient, sourceChainClient.GetChainID().Uint64(), destCommon)

	go func() {
		lane.Source.Common.DeployContracts(t, len(lane.Source.TransferAmount))
	}()
	go func() {
		lane.Dest.Common.DeployContracts(t, len(lane.Source.TransferAmount))
	}()

	// deploy all source contracts
	lane.Source.DeployContracts(t, model)

	// deploy all destination contracts
	lane.Dest.DeployContracts(t, *lane.Source, wg, model)

	// set up ocr2 jobs
	var tokenAddr []string
	for _, token := range lane.Dest.Common.BridgeTokens {
		tokenAddr = append(tokenAddr, token.Address())
	}
	clNodes, exists := clNodesWithKeys[lane.Dest.Common.ChainClient.GetChainID().String()]
	require.True(t, exists)

	tokenAddr = append(tokenAddr, lane.Dest.Common.FeeToken.Address())
	// first node is the bootstrapper
	bootstrapCommit := clNodes[0]
	var bootstrapExec *client.CLNodesWithKeys
	var execNodes []*client.CLNodesWithKeys
	commitNodes := clNodes[1:]
	env.commitNodeStartIndex = 1
	env.execNodeStartIndex = 1
	env.numOfAllowedFaultyExec = 1
	env.numOfAllowedFaultyCommit = 1
	if !commitAndExecOnSameDON {
		bootstrapExec = clNodes[1] // for a set-up of different commit and execution nodes second node is the bootstrapper for execution nodes
		commitNodes = clNodes[2 : 2+numOfCommitNodes]
		execNodes = clNodes[2+numOfCommitNodes:]
		env.commitNodeStartIndex = 2
		env.execNodeStartIndex = 7
	}
	if model == TOLL {
		CreateOCRJobsForCCIP(
			t,
			bootstrapCommit, bootstrapExec, commitNodes, execNodes,
			lane.Source.TollOnRamp.EthAddress,
			common.HexToAddress("0x0"),
			lane.Dest.CommitStore.EthAddress,
			lane.Dest.TollOffRamp.EthAddress,
			common.HexToAddress("0x0"),
			sourceChainClient, destChainClient,
			tokenAddr,
			mockServer, newBootstrap,
		)
	}
	if model == GE {
		CreateOCRJobsForCCIP(
			t,
			bootstrapCommit, bootstrapExec, commitNodes, execNodes,
			common.HexToAddress("0x0"),
			lane.Source.GEOnRamp.EthAddress,
			lane.Dest.CommitStore.EthAddress,
			common.HexToAddress("0x0"),
			lane.Dest.GEOffRamp.EthAddress,
			sourceChainClient, destChainClient,
			tokenAddr,
			mockServer, newBootstrap,
		)
	}

	// set up ocr2 config
	SetOCR2Configs(t, commitNodes, execNodes, *lane.Dest)
	lane.Ready <- struct{}{}
}
