package actions

import (
	"context"
	_ "embed"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
	"github.com/smartcontractkit/chainlink-env/environment"
	"github.com/smartcontractkit/chainlink-env/pkg/cdk8s/blockscout"
	"github.com/smartcontractkit/chainlink-env/pkg/helm/chainlink"
	"github.com/smartcontractkit/chainlink-env/pkg/helm/mockserver"
	mockservercfg "github.com/smartcontractkit/chainlink-env/pkg/helm/mockserver-cfg"
	"github.com/smartcontractkit/chainlink-env/pkg/helm/reorg"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	ctfClient "github.com/smartcontractkit/chainlink-testing-framework/client"
	ctfUtils "github.com/smartcontractkit/chainlink-testing-framework/utils"
	"golang.org/x/exp/slices"
	"gopkg.in/guregu/null.v4"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_toll_onramp_router"
	"github.com/smartcontractkit/chainlink/core/services/job"
	ccipPlugin "github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/testhelpers"
	"github.com/smartcontractkit/chainlink/core/services/relay"
	"github.com/smartcontractkit/chainlink/core/store/models"
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
	// TODO dynamic calculation of tollfee for multiple tokens in a msg

	TollFee   = big.NewInt(0).Mul(big.NewInt(12), big.NewInt(1e18)) // "maxCharge":"10784576000000000000" for a msg with two tokens
	UnusedFee = big.NewInt(0).Mul(big.NewInt(11), big.NewInt(1e18)) // for a msg with two tokens

	//go:embed clconfig/ccip-default.txt
	CLConfig                   string
	sourceNetwork, destNetwork = func() (blockchain.EVMNetwork, blockchain.EVMNetwork) {
		if len(networks.SelectedNetworks) < 3 {
			log.Fatal().
				Interface("SELECTED_NETWORKS", networks.SelectedNetworks).
				Msg("Set source and destination network in index 1 & 2 of env variable SELECTED_NETWORKS")
		}
		log.Info().
			Interface("Source Network", networks.SelectedNetworks[1]).
			Interface("Destination Network", networks.SelectedNetworks[2]).
			Msg("SELECTED_NETWORKS")
		return *networks.SelectedNetworks[1], *networks.SelectedNetworks[2]
	}()

	DefaultCCIPCLNodeEnv = func() string {
		ccipTOML, err := client.MarshallTemplate(
			CCIPTOMLEnv{
				Networks: []blockchain.EVMNetwork{sourceNetwork, destNetwork},
			},
			"ccip env toml", CLConfig)
		Expect(err).ShouldNot(HaveOccurred())
		return ccipTOML
	}
)

type CCIPCommon struct {
	ChainClient       blockchain.EVMClient
	Deployer          *ccip.CCIPContractsDeployer
	FeeToken          contracts.LinkToken
	FeeTokenPool      *ccip.NativeTokenPool
	BridgeTokens      []contracts.LinkToken // as of now considering the bridge token is same as link token
	TokenPrices       []*big.Int
	BridgeTokenPools  []*ccip.NativeTokenPool
	RateLimiterConfig ccip.RateLimiterConfig
	AFNConfig         ccip.AFNConfig
	AFN               *ccip.AFN
}

// DeployContracts deploys the contracts which are necessary in both source and dest chain
func (ccipModule *CCIPCommon) DeployContracts(cd *ccip.CCIPContractsDeployer, noOfTokens int) {
	// deploy link token
	token, err := cd.DeployLinkTokenContract()
	Expect(err).ShouldNot(HaveOccurred(), "Deploying Link Token Contract shouldn't fail")
	ccipModule.FeeToken = token

	// deploy bridge token.
	for i := 0; i < noOfTokens; i++ {
		token, err = cd.DeployLinkTokenContract()
		Expect(err).ShouldNot(HaveOccurred(), "Deploying Link Token Contract shouldn't fail")
		ccipModule.BridgeTokens = append(ccipModule.BridgeTokens, token)
	}

	// Set price of the bridge tokens to 1
	ccipModule.TokenPrices = []*big.Int{}
	for range ccipModule.BridgeTokens {
		ccipModule.TokenPrices = append(ccipModule.TokenPrices, big.NewInt(1))
	}

	err = ccipModule.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for Link Token deployments")
	// deploy native token pool
	for _, token = range ccipModule.BridgeTokens {
		ntp, err := cd.DeployNativeTokenPoolContract(token.Address())
		Expect(err).ShouldNot(HaveOccurred(), "Deploying Native TokenPool Contract shouldn't fail")
		ccipModule.BridgeTokenPools = append(ccipModule.BridgeTokenPools, ntp)
	}
	// token pool for fee token
	ccipModule.FeeTokenPool, err = cd.DeployNativeTokenPoolContract(ccipModule.FeeToken.Address())
	Expect(err).ShouldNot(HaveOccurred(), "Deploying Native TokenPool Contract shouldn't fail")

	// deploy AFN
	ccipModule.AFN, err = cd.DeployAFNContract()
	Expect(err).ShouldNot(HaveOccurred(), "Deploying AFN Contract shouldn't fail")

	err = ccipModule.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for contract deployments")
}

func DefaultCCIPModule(chainClient blockchain.EVMClient) *CCIPCommon {
	return &CCIPCommon{
		ChainClient: chainClient,
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
	DestinationChainId *big.Int
	TollOnRampRouter   *ccip.TollOnRampRouter
	TollOnRamp         *ccip.TollOnRamp
	TollSender         *ccip.TollSender
}

// DeploySenderApp deploys TollSenderApp. It is a bit outdated and only accepts feeAmount as zero.
// execution may revert if the feeconfig is set to an amount more than zero in onramp.
func (sourceCCIP *SourceCCIPModule) DeploySenderApp(destCCIP DestCCIPModule) {
	var err error
	sourceCCIP.TollSender, err = sourceCCIP.Common.Deployer.DeployTollSenderDapp(
		sourceCCIP.TollOnRampRouter.EthAddress,
		destCCIP.ReceiverDapp.EthAddress,
		destCCIP.Common.ChainClient.GetChainID())
	Expect(err).ShouldNot(HaveOccurred(), "Toll Sender contract should be deployed successfully")
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for deployments")
}

// DeployContracts deploys all CCIP contracts specific to the source chain
func (sourceCCIP *SourceCCIPModule) DeployContracts() {
	var err error
	sourceCCIP.Common.Deployer, err = ccip.NewCCIPContractsDeployer(sourceCCIP.Common.ChainClient)
	Expect(err).ShouldNot(HaveOccurred(), "contract deployer should be created successfully")
	contractDeployer := sourceCCIP.Common.Deployer
	sourceCCIP.Common.DeployContracts(contractDeployer, len(sourceCCIP.TransferAmount))

	// deploy on ramp router
	sourceCCIP.TollOnRampRouter, err = contractDeployer.DeployTollOnRampRouter()
	Expect(err).ShouldNot(HaveOccurred(), "Deploying onramp router shouldn't fail")
	// wait for all contract deployments before moving on to on-ramp deployment
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for deployments")

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

	// Toll Set up

	// onRamp
	sourceCCIP.TollOnRamp, err = contractDeployer.DeployTollOnRamp(
		sourceCCIP.Common.ChainClient.GetChainID(), sourceCCIP.DestinationChainId,
		tokens, pools, []common.Address{}, sourceCCIP.Common.AFN.EthAddress,
		sourceCCIP.TollOnRampRouter.EthAddress, sourceCCIP.Common.RateLimiterConfig)
	Expect(err).ShouldNot(HaveOccurred(), "Error on OnRamp deployment")
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for OnRamp deployment")

	// Set bridge token prices on the onRamp
	err = sourceCCIP.TollOnRamp.SetTokenPrices(bridgeTokens, sourceCCIP.Common.TokenPrices)
	Expect(err).ShouldNot(HaveOccurred(), "Setting prices shouldn't fail")

	// update onRampRouter with OnRamp address
	err = sourceCCIP.TollOnRampRouter.SetOnRamp(sourceCCIP.DestinationChainId, sourceCCIP.TollOnRamp.EthAddress)
	Expect(err).ShouldNot(HaveOccurred(), "Error setting onramp on the router")

	// update native pool with onRamp address
	for _, pool := range sourceCCIP.Common.BridgeTokenPools {
		err = pool.SetOnRamp(sourceCCIP.TollOnRamp.EthAddress)
		Expect(err).ShouldNot(HaveOccurred(), "Error setting tollonramp on the token pool %s", pool.Address())
	}
	err = sourceCCIP.Common.FeeTokenPool.SetOnRamp(sourceCCIP.TollOnRamp.EthAddress)
	Expect(err).ShouldNot(HaveOccurred(), "Error setting TollOnRamp on the token pool %s", sourceCCIP.Common.FeeTokenPool.Address())

	// The Fee token to be used in on+off ramp
	sourceCCIP.TollFeeAmount = TollFee

	// set a part of sourceCCIP.TollFeeAmount as onRamp fee rest would be used as offramp
	err = sourceCCIP.TollOnRamp.SetFeeConfig([]common.Address{common.HexToAddress(sourceCCIP.Common.FeeToken.Address())}, []*big.Int{big.NewInt(1)})
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for setting OnRamp Fee config")

	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for events")
}

func (sourceCCIP *SourceCCIPModule) CollectBalanceRequirements(model BillingModel) []testhelpers.BalanceReq {
	var balancesReq []testhelpers.BalanceReq
	for _, token := range sourceCCIP.Common.BridgeTokens {
		balancesReq = append(balancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("%s-BridgeToken-%s", testhelpers.Sender, token.Address()),
			Addr:   sourceCCIP.Sender,
			Getter: GetterForLinkToken(token, sourceCCIP.Sender.Hex()),
		})
	}
	for i, pool := range sourceCCIP.Common.BridgeTokenPools {
		balancesReq = append(balancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("%s-TokenPool-%s", testhelpers.Sender, pool.Address()),
			Addr:   pool.EthAddress,
			Getter: GetterForLinkToken(sourceCCIP.Common.BridgeTokens[i], pool.Address()),
		})
	}
	balancesReq = append(balancesReq, testhelpers.BalanceReq{
		Name:   fmt.Sprintf("%s-FeeToken-%s", testhelpers.Sender, sourceCCIP.Common.FeeToken.Address()),
		Addr:   sourceCCIP.Sender,
		Getter: GetterForLinkToken(sourceCCIP.Common.FeeToken, sourceCCIP.Sender.Hex()),
	})
	return balancesReq
}

func (sourceCCIP *SourceCCIPModule) BalanceAssertions(model BillingModel, prevBalances map[string]*big.Int, noOfreq int64) []testhelpers.BalanceAssertion {
	var balAssertions []testhelpers.BalanceAssertion
	for i, token := range sourceCCIP.Common.BridgeTokens {
		name := fmt.Sprintf("%s-BridgeToken-%s", testhelpers.Sender, token.Address())
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     name,
			Address:  sourceCCIP.Sender,
			Getter:   GetterForLinkToken(token, sourceCCIP.Sender.Hex()),
			Expected: bigmath.Sub(prevBalances[name], bigmath.Mul(big.NewInt(noOfreq), sourceCCIP.TransferAmount[i])).String(),
		})
	}
	for i, pool := range sourceCCIP.Common.BridgeTokenPools {
		name := fmt.Sprintf("%s-TokenPool-%s", testhelpers.Sender, pool.Address())
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     fmt.Sprintf("%s-TokenPool-%s", testhelpers.Sender, pool.Address()),
			Address:  pool.EthAddress,
			Getter:   GetterForLinkToken(sourceCCIP.Common.BridgeTokens[i], pool.Address()),
			Expected: bigmath.Add(prevBalances[name], bigmath.Mul(big.NewInt(noOfreq), sourceCCIP.TransferAmount[i])).String(),
		})
	}
	switch model {
	case TOLL:
		name := fmt.Sprintf("%s-FeeToken-%s", testhelpers.Sender, sourceCCIP.Common.FeeToken.Address())
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     name,
			Address:  sourceCCIP.Sender,
			Getter:   GetterForLinkToken(sourceCCIP.Common.FeeToken, sourceCCIP.Sender.Hex()),
			Expected: bigmath.Sub(prevBalances[name], bigmath.Mul(big.NewInt(noOfreq), sourceCCIP.TollFeeAmount)).String(),
		})
	}
	return balAssertions
}

func (sourceCCIP *SourceCCIPModule) AssertEventCCIPSendRequested(model BillingModel, txHash string, currentBlockOnSource uint64, timeout time.Duration) uint64 {
	log.Info().Msg("Waiting for CCIPSendRequested event")
	var seqNum uint64
	switch model {
	case TOLL:
		Eventually(func(g Gomega) bool {
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
	}
	return seqNum
}

func (sourceCCIP *SourceCCIPModule) SendTollRequest(receiver common.Address, tokenAndAmounts []evm_2_any_toll_onramp_router.CCIPEVMTokenAndAmount, data string, feeToken evm_2_any_toll_onramp_router.CCIPEVMTokenAndAmount) string {
	receiverAddr, err := utils.ABIEncode(`[{"type":"address"}]`, receiver)
	Expect(err).ShouldNot(HaveOccurred(), "Failed encoding the receiver address")

	extraArgsV1, err := testhelpers.GetEVMExtraArgsV1(big.NewInt(100_000), false)
	Expect(err).ShouldNot(HaveOccurred(), "Failed encoding the options field")

	// form the message for transfer
	msg := evm_2_any_toll_onramp_router.CCIPEVM2AnyTollMessage{
		Receiver:          receiverAddr,
		Data:              []byte(data),
		TokensAndAmounts:  tokenAndAmounts,
		FeeTokenAndAmount: feeToken,
		ExtraArgs:         extraArgsV1,
	}
	log.Info().Interface("msg details", msg).Msg("ccip message to be sent")

	// initiate the transfer
	sendTx, err := sourceCCIP.TollOnRampRouter.CCIPSend(sourceCCIP.DestinationChainId, msg)
	Expect(err).ShouldNot(HaveOccurred(), "send token should be initiated successfully")
	log.Info().Str("toll send token transaction", sendTx.Hash().String()).Msg("Sending token")
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Failed to wait for events")
	return sendTx.Hash().Hex()
}

func DefaultSourceCCIPModule(chainClient blockchain.EVMClient, destChain *big.Int, transferamount []*big.Int) *SourceCCIPModule {
	return &SourceCCIPModule{
		Common:             DefaultCCIPModule(chainClient),
		TransferAmount:     transferamount,
		DestinationChainId: destChain,
		Sender:             common.HexToAddress(chainClient.GetDefaultWallet().Address()),
	}
}

type DestCCIPModule struct {
	Common            *CCIPCommon
	SourceChainId     *big.Int
	CommitStore       *ccip.CommitStore
	TollOffRamp       *ccip.TollOffRamp
	TollOffRampRouter *ccip.TollOffRampRouter
	ReceiverDapp      *ccip.ReceiverDapp
}

// DeployContracts deploys all CCIP contracts specific to the destination chain
func (destCCIP *DestCCIPModule) DeployContracts(sourceCCIP SourceCCIPModule) {
	var err error
	destCCIP.Common.Deployer, err = ccip.NewCCIPContractsDeployer(destCCIP.Common.ChainClient)
	Expect(err).ShouldNot(HaveOccurred(), "contract deployer should be created successfully")
	contractDeployer := destCCIP.Common.Deployer
	destCCIP.Common.DeployContracts(contractDeployer, len(sourceCCIP.TransferAmount))

	// commitStore responsible for validating the transfer message
	destCCIP.CommitStore, err = contractDeployer.DeployCommitStore(
		destCCIP.SourceChainId,
		destCCIP.Common.ChainClient.GetChainID(),
		destCCIP.Common.AFN.EthAddress,
		commit_store.CommitStoreInterfaceCommitStoreConfig{
			OnRamps:          []common.Address{sourceCCIP.TollOnRamp.EthAddress},
			MinSeqNrByOnRamp: []uint64{1},
		})
	Expect(err).ShouldNot(HaveOccurred(), "Deploying CommitStore shouldn't fail")
	err = destCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for setting destination contracts")

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
		Expect(err).ShouldNot(HaveOccurred())
	}
	// add the fee token and feetoken price for dest
	destTokens = append(destTokens, common.HexToAddress(destCCIP.Common.FeeToken.Address()))
	destCCIP.Common.TokenPrices = append(destCCIP.Common.TokenPrices, big.NewInt(1))

	pools = append(pools, destCCIP.Common.FeeTokenPool.EthAddress)
	err = destCCIP.Common.FeeToken.Transfer(destCCIP.Common.FeeTokenPool.Address(), testhelpers.Link(1000))
	Expect(err).ShouldNot(HaveOccurred())

	// Toll
	// offRamp
	destCCIP.TollOffRamp, err = contractDeployer.DeployTollOffRamp(destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID(),
		destCCIP.CommitStore.EthAddress, sourceCCIP.TollOnRamp.EthAddress, destCCIP.Common.AFN.EthAddress,
		sourceTokens, pools, destCCIP.Common.RateLimiterConfig)
	Expect(err).ShouldNot(HaveOccurred(), "Deploying TollOffRamp shouldn't fail")
	err = destCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for deploying offramp")
	// Set token prices on the offRamp
	err = destCCIP.TollOffRamp.SetTokenPrices(destTokens, destCCIP.Common.TokenPrices)
	Expect(err).ShouldNot(HaveOccurred(), "Setting prices shouldn't fail")

	// OffRampRouter
	destCCIP.TollOffRampRouter, err = contractDeployer.DeployTollOffRampRouter([]common.Address{destCCIP.TollOffRamp.EthAddress})
	Expect(err).ShouldNot(HaveOccurred(), "Deploying TollOffRampRouter shouldn't fail")
	err = destCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for deploying OffRampRouter")
	err = destCCIP.TollOffRamp.SetRouter(destCCIP.TollOffRampRouter.EthAddress)
	Expect(err).ShouldNot(HaveOccurred(), "Error setting router on the offramp")

	// ReceiverDapp
	destCCIP.ReceiverDapp, err = contractDeployer.DeployReceiverDapp(false)
	Expect(err).ShouldNot(HaveOccurred(), "ReceiverDapp contract should be deployed successfully")

	// update pools with offRamp Id
	for _, pool := range destCCIP.Common.BridgeTokenPools {
		err = pool.SetOffRamp(destCCIP.TollOffRamp.EthAddress)
		Expect(err).ShouldNot(HaveOccurred(), "Error setting TollOffRamp on the token pool %s", pool.Address())
	}
	err = destCCIP.Common.FeeTokenPool.SetOffRamp(destCCIP.TollOffRamp.EthAddress)
	Expect(err).ShouldNot(HaveOccurred(), "Error setting TollOffRamp on the token pool %s", destCCIP.Common.FeeTokenPool.Address())

	err = destCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for events on destination contract deployments")
}

func (destCCIP *DestCCIPModule) CollectBalanceRequirements(model BillingModel) []testhelpers.BalanceReq {
	var destBalancesReq []testhelpers.BalanceReq
	for _, token := range destCCIP.Common.BridgeTokens {
		destBalancesReq = append(destBalancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("%s-BridgeToken-%s", testhelpers.Receiver, token.Address()),
			Addr:   destCCIP.ReceiverDapp.EthAddress,
			Getter: GetterForLinkToken(token, destCCIP.ReceiverDapp.Address()),
		})
	}
	for i, pool := range destCCIP.Common.BridgeTokenPools {
		destBalancesReq = append(destBalancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("%s-TokenPool-%s", testhelpers.Receiver, pool.Address()),
			Addr:   pool.EthAddress,
			Getter: GetterForLinkToken(destCCIP.Common.BridgeTokens[i], pool.Address()),
		})
	}
	destBalancesReq = append(destBalancesReq, testhelpers.BalanceReq{
		Name:   fmt.Sprintf("%s-FeeToken-%s", testhelpers.Receiver, destCCIP.Common.FeeToken.Address()),
		Addr:   destCCIP.ReceiverDapp.EthAddress,
		Getter: GetterForLinkToken(destCCIP.Common.FeeToken, destCCIP.ReceiverDapp.Address()),
	})
	return destBalancesReq
}

func (destCCIP *DestCCIPModule) BalanceAssertions(prevBalances map[string]*big.Int, transferAmount []*big.Int, unusedFee *big.Int, noOfReq int64) []testhelpers.BalanceAssertion {
	var balAssertions []testhelpers.BalanceAssertion
	for i, token := range destCCIP.Common.BridgeTokens {
		name := fmt.Sprintf("%s-BridgeToken-%s", testhelpers.Receiver, token.Address())
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     name,
			Address:  destCCIP.ReceiverDapp.EthAddress,
			Getter:   GetterForLinkToken(token, destCCIP.ReceiverDapp.Address()),
			Expected: bigmath.Add(prevBalances[name], bigmath.Mul(big.NewInt(noOfReq), transferAmount[i])).String(),
		})
	}
	for i, pool := range destCCIP.Common.BridgeTokenPools {
		name := fmt.Sprintf("%s-TokenPool-%s", testhelpers.Receiver, pool.Address())
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     fmt.Sprintf("%s-TokenPool-%s", testhelpers.Receiver, pool.Address()),
			Address:  pool.EthAddress,
			Getter:   GetterForLinkToken(destCCIP.Common.BridgeTokens[i], pool.Address()),
			Expected: bigmath.Sub(prevBalances[name], bigmath.Mul(big.NewInt(noOfReq), transferAmount[i])).String(),
		})
	}
	name := fmt.Sprintf("%s-FeeToken-%s", testhelpers.Receiver, destCCIP.Common.FeeToken.Address())
	balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
		Name:     name,
		Address:  destCCIP.ReceiverDapp.EthAddress,
		Getter:   GetterForLinkToken(destCCIP.Common.FeeToken, destCCIP.ReceiverDapp.Address()),
		Expected: bigmath.Add(prevBalances[name], unusedFee.String()).String(),
		Within:   big.NewInt(1e18).String(),
	})
	return balAssertions
}

func (destCCIP *DestCCIPModule) AssertEventExecutionStateChanged(model BillingModel, seqNum uint64, currentBlockOnDest uint64, timeout time.Duration) {
	log.Info().Int64("seqNum", int64(seqNum)).Msg("Waiting for ExecutionStateChanged event")
	switch model {
	case TOLL:
		Eventually(func(g Gomega) ccipPlugin.MessageExecutionState {
			iterator, err := destCCIP.TollOffRamp.FilterExecutionStateChanged([]uint64{seqNum}, currentBlockOnDest)
			g.Expect(err).NotTo(HaveOccurred(), "Error filtering ExecutionStateChanged event for seqNum %d", seqNum)
			g.Expect(iterator.Next()).To(BeTrue(), "No ExecutionStateChanged event found for seqNum %d", seqNum)
			return ccipPlugin.MessageExecutionState(iterator.Event.State)
		}, timeout, "1s").Should(Equal(ccipPlugin.Success))
	}
}

func (destCCIP *DestCCIPModule) AssertEventReportAccepted(onRamp common.Address, seqNum, currentBlockOnDest uint64, timeout time.Duration) {
	log.Info().Int64("seqNum", int64(seqNum)).Msg("Waiting for ReportAccepted event")
	Eventually(func(g Gomega) bool {
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

func (destCCIP *DestCCIPModule) AssertSeqNumberExecuted(onRamp common.Address, seqNumberBefore uint64, timeout time.Duration) {
	log.Info().Int64("seqNum", int64(seqNumberBefore)).Msg("Waiting to be executed")
	Eventually(func(g Gomega) {
		seqNumberAfter, err := destCCIP.CommitStore.GetNextSeqNumber(onRamp)
		g.Expect(err).ShouldNot(HaveOccurred(), "Getting expected seq number should be successful %d", seqNumberBefore)
		g.Expect(seqNumberAfter).Should(BeNumerically(">", seqNumberBefore), "Next Sequence number is not increased")
	}, timeout, "1s").Should(Succeed(), "Error Executing Sequence number %d", seqNumberBefore)
}

func DefaultDestinationCCIPModule(chainClient blockchain.EVMClient, sourceChain *big.Int) *DestCCIPModule {
	return &DestCCIPModule{
		Common:        DefaultCCIPModule(chainClient),
		SourceChainId: sourceChain,
	}
}

func NewCCIPTest(sourceCCIP *SourceCCIPModule, destCCIP *DestCCIPModule, timeout time.Duration) *CCIPTest {
	return &CCIPTest{
		Source:            sourceCCIP,
		Dest:              destCCIP,
		ValidationTimeout: timeout,
	}
}

type CCIPTest struct {
	Source                  *SourceCCIPModule
	Dest                    *DestCCIPModule
	NumberOfTollReq         int
	SourceBalances          map[string]*big.Int
	DestBalances            map[string]*big.Int
	StartBlockOnSource      uint64
	StartBlockOnDestination uint64
	SentTollReqHashes       []string
	ValidationTimeout       time.Duration
}

func (c *CCIPTest) SendTollRequests(noOfRequests int) {
	c.NumberOfTollReq = noOfRequests
	var tokenAndAmounts []evm_2_any_toll_onramp_router.CCIPEVMTokenAndAmount
	for i, token := range c.Source.Common.BridgeTokens {
		tokenAndAmounts = append(tokenAndAmounts, evm_2_any_toll_onramp_router.CCIPEVMTokenAndAmount{
			Token: common.HexToAddress(token.Address()), Amount: c.Source.TransferAmount[i],
		})
		// approve the onramp router so that it can initiate transferring the token
		err := token.Approve(c.Source.TollOnRampRouter.Address(), bigmath.Mul(c.Source.TransferAmount[i], big.NewInt(int64(c.NumberOfTollReq))))
		Expect(err).ShouldNot(HaveOccurred(), "Could not approve permissions for the onRamp router "+
			"on the source link token contract")
	}
	// approve onramp router to use the feetoken
	err := c.Source.Common.FeeToken.Approve(c.Source.TollOnRampRouter.Address(), bigmath.Mul(c.Source.TollFeeAmount, big.NewInt(int64(c.NumberOfTollReq))))
	Expect(err).ShouldNot(HaveOccurred(), "Could not approve permissions for the onRamp router "+
		"on the source link token contract")
	err = c.Source.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Failed to wait for events")

	// collect the balance requirement to verify balances after transfer
	c.SourceBalances, err = testhelpers.GetBalances(c.Source.CollectBalanceRequirements(TOLL))
	Expect(err).ShouldNot(HaveOccurred(), "fetching source balance")
	c.DestBalances, err = testhelpers.GetBalances(c.Dest.CollectBalanceRequirements(TOLL))
	Expect(err).ShouldNot(HaveOccurred(), "fetching dest balance")

	// save the current block numbers to use in various filter log requests
	c.StartBlockOnSource, err = c.Source.Common.ChainClient.LatestBlockNumber(context.Background())
	Expect(err).ShouldNot(HaveOccurred(), "Getting current block should be successful in source chain")
	c.StartBlockOnDestination, err = c.Dest.Common.ChainClient.LatestBlockNumber(context.Background())
	Expect(err).ShouldNot(HaveOccurred(), "Getting current block should be successful in dest chain")

	for i := 1; i <= c.NumberOfTollReq; i++ {
		txHash := c.Source.SendTollRequest(
			c.Dest.ReceiverDapp.EthAddress,
			tokenAndAmounts,
			fmt.Sprintf("msg %d", i),
			evm_2_any_toll_onramp_router.CCIPEVMTokenAndAmount{
				Token: common.HexToAddress(c.Source.Common.FeeToken.Address()), Amount: c.Source.TollFeeAmount,
			},
		)
		c.SentTollReqHashes = append(c.SentTollReqHashes, txHash)
	}
}

func (c *CCIPTest) ValidateTollRequests() {
	for _, txHash := range c.SentTollReqHashes {
		// Verify if
		// - CCIPSendRequested Event log generated,
		// - NextSeqNumber from commitStore got increased
		seqNumber := c.Source.AssertEventCCIPSendRequested(TOLL, txHash, c.StartBlockOnSource, c.ValidationTimeout)
		c.Dest.AssertSeqNumberExecuted(c.Source.TollOnRamp.EthAddress, seqNumber, c.ValidationTimeout)

		// Verify whether commitStore has accepted the report
		c.Dest.AssertEventReportAccepted(c.Source.TollOnRamp.EthAddress, seqNumber, c.StartBlockOnDestination, c.ValidationTimeout)

		// Verify whether the execution state is changed and the transfer is successful
		c.Dest.AssertEventExecutionStateChanged(TOLL, seqNumber, c.StartBlockOnDestination, c.ValidationTimeout)
	}
	// verify the fee amount is deducted from sender, added to receiver token balances and
	// unused fee is returned to receiver fee token account
	AssertBalances(c.Source.BalanceAssertions(TOLL, c.SourceBalances, int64(c.NumberOfTollReq)))
	AssertBalances(c.Dest.BalanceAssertions(c.DestBalances, c.Source.TransferAmount, UnusedFee, int64(c.NumberOfTollReq)))
}

// SetOCRConfigs sets the oracle config in ocr2 contracts
// nil value in execNodes denotes commit and execution jobs are to be set up in same DON
func SetOCRConfigs(commitNodes, execNodes []*client.CLNodesWithKeys, destCCIP DestCCIPModule) {
	signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, err :=
		ccip.NewOffChainAggregatorV2Config(commitNodes)
	Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail while getting the config values for ocr2 type contract")
	err = destCCIP.CommitStore.SetOCRConfig(signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
	Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail while setting commitStore config")
	// if commit and exec job is set up in different DON
	if len(execNodes) > 0 {
		signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, err =
			ccip.NewOffChainAggregatorV2Config(execNodes)
	}
	Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail while getting the config values for ocr2 type contract")
	err = destCCIP.TollOffRamp.SetConfig(signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
	Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail while setting TollOffRamp config")
	err = destCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail while waiting for events on setting ocr2 config")
}

// CreateOCRJobsForCCIP bootstraps the first node and to the other nodes sends ocr jobs that
// sets up ccip-commit and ccip-execution plugin
// nil value in bootstrapExec and execNodes denotes commit and execution jobs are to be set up in same DON
func CreateOCRJobsForCCIP(
	bootstrapCommit *client.CLNodesWithKeys,
	bootstrapExec *client.CLNodesWithKeys,
	commitNodes, execNodes []*client.CLNodesWithKeys,
	tollOnRamp, commitStore, tollOffRamp string,
	sourceChainClient, destChainClient blockchain.EVMClient,
	linkTokenAddr []string,
	mockServer *ctfClient.MockserverClient,
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
	sourceChainID := sourceChainClient.GetChainID()
	destChainID := destChainClient.GetChainID()
	sourceChainName := sourceChainClient.GetNetworkName()
	destChainName := destChainClient.GetNetworkName()
	bootstrapSpec := func(contractID string) *client.OCR2TaskJobSpec {
		return &client.OCR2TaskJobSpec{
			Name:    fmt.Sprintf("bootstrap-%s-%s", destChainName, uuid.NewV4().String()),
			JobType: "bootstrap",
			OCR2OracleSpec: job.OCR2OracleSpec{
				ContractID:                        contractID,
				Relay:                             relay.EVM,
				ContractConfigConfirmations:       1,
				ContractConfigTrackerPollInterval: models.Interval(1 * time.Second),
				RelayConfig: map[string]interface{}{
					"chainID": fmt.Sprintf("\"%s\"", destChainID.String()),
				},
			},
		}
	}

	_, err := bootstrapCommit.Node.MustCreateJob(bootstrapSpec(commitStore))
	Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail creating bootstrap job on bootstrap node")
	if bootstrapExec != nil && len(execNodes) > 0 {
		_, err := bootstrapExec.Node.MustCreateJob(bootstrapSpec(tollOffRamp))
		Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail creating bootstrap job on bootstrap node")
	} else {
		execNodes = commitNodes
	}
	// save the current block numbers. If there is a delay between job start up and ocr config set up, the jobs will
	// replay the log polling from these mentioned block number. The dest block number should ideally be the block number on which
	// contract config is set and the source block number should be the one on which the ccip send request is performed.
	// Here for simplicity we are just taking the current block number just before the job is created.
	currentBlockOnSource, err := sourceChainClient.LatestBlockNumber(context.Background())
	Expect(err).ShouldNot(HaveOccurred(), "Getting current block should be successful in source chain")
	currentBlockOnDest, err := destChainClient.LatestBlockNumber(context.Background())
	Expect(err).ShouldNot(HaveOccurred(), "Getting current block should be successful in dest chain")

	tokenFeeConv := make(map[string]interface{})
	for _, token := range linkTokenAddr {
		tokenFeeConv[token] = "200000000000000000000"
	}
	SetMockServerWithSameTokenFeeConversionValue(tokenFeeConv, execNodes, mockServer)
	p2pBootstrappersCommit := &client.P2PData{
		RemoteIP: bootstrapCommit.Node.RemoteIP(),
		PeerID:   bootstrapCommitP2PId,
	}
	p2pBootstrappersExec := &client.P2PData{
		RemoteIP: bootstrapExec.Node.RemoteIP(),
		PeerID:   bootstrapExecP2PId,
	}
	addCommitJob := func(node *client.Chainlink, nodeTransmitterAddress, nodeOCR2KeyId string) error {
		ocr2SpecCommit := &client.OCR2TaskJobSpec{
			JobType: "offchainreporting2",
			Name:    fmt.Sprintf("ccip-commit-%s-%s", sourceChainName, destChainName),
			OCR2OracleSpec: job.OCR2OracleSpec{
				Relay:                             relay.EVM,
				PluginType:                        job.CCIPCommit,
				ContractID:                        commitStore,
				OCRKeyBundleID:                    null.StringFrom(nodeOCR2KeyId),
				TransmitterID:                     null.StringFrom(nodeTransmitterAddress),
				ContractConfigConfirmations:       1,
				ContractConfigTrackerPollInterval: models.Interval(1 * time.Second),
				P2PV2Bootstrappers: []string{
					p2pBootstrappersCommit.P2PV2Bootstrapper(),
				},
				PluginConfig: map[string]interface{}{
					"sourceChainID":       sourceChainID,
					"destChainID":         destChainID,
					"onRampIDs":           fmt.Sprintf("[\"%s\"]", tollOnRamp),
					"pollPeriod":          `"1s"`,
					"destStartBlock":      currentBlockOnDest,
					"sourceStartBlock":    currentBlockOnSource,
					"inflightCacheExpiry": fmt.Sprintf("\"%s\"", InflightExpiry.String()),
				},
				RelayConfig: map[string]interface{}{
					"chainID": destChainID,
				},
			},
		}
		_, err = node.MustCreateJob(ocr2SpecCommit)
		return err
	}

	addExecJob := func(node *client.CLNodesWithKeys, jobname, nodeTransmitterAddress, nodeOCR2KeyId, offRamp, onRamp string) error {
		tokensPerFeeCoinPipeline := TokenFeeForMultipleTokenAddr(node, linkTokenAddr, mockServer)
		ocr2SpecExec := &client.OCR2TaskJobSpec{
			JobType: "offchainreporting2",
			Name:    jobname,
			OCR2OracleSpec: job.OCR2OracleSpec{
				Relay:                             relay.EVM,
				PluginType:                        job.CCIPExecution,
				ContractID:                        offRamp,
				OCRKeyBundleID:                    null.StringFrom(nodeOCR2KeyId),
				TransmitterID:                     null.StringFrom(nodeTransmitterAddress),
				ContractConfigConfirmations:       1,
				ContractConfigTrackerPollInterval: models.Interval(1 * time.Second),
				P2PV2Bootstrappers: []string{
					p2pBootstrappersExec.P2PV2Bootstrapper(),
				},
				PluginConfig: map[string]interface{}{
					"sourceChainID":    sourceChainID,
					"destChainID":      destChainID,
					"onRampID":         fmt.Sprintf("\"%s\"", onRamp),
					"commitStoreID":    fmt.Sprintf("\"%s\"", commitStore),
					"destStartBlock":   currentBlockOnDest,
					"sourceStartBlock": currentBlockOnSource,
					"tokensPerFeeCoinPipeline": fmt.Sprintf(`"""
%s
"""`, tokensPerFeeCoinPipeline),
					"rootSnoozeTime":      fmt.Sprintf("\"%s\"", RootSnoozeTime.String()),
					"inflightCacheExpiry": fmt.Sprintf("\"%s\"", InflightExpiry.String()),
				},
				RelayConfig: map[string]interface{}{
					"chainID": destChainID,
				},
			},
		}
		_, err = node.Node.MustCreateJob(ocr2SpecExec)
		return err
	}

	for nodeIndex := 0; nodeIndex < len(commitNodes); nodeIndex++ {
		nodeTransmitterAddress := commitNodes[nodeIndex].KeysBundle.EthAddress
		nodeOCR2Key := commitNodes[nodeIndex].KeysBundle.OCR2Key
		nodeOCR2KeyId := nodeOCR2Key.Data.ID
		err := addCommitJob(commitNodes[nodeIndex].Node, nodeTransmitterAddress, nodeOCR2KeyId)
		Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail creating CCIP-Commit OCR Task job on OCR node %d", nodeIndex+1)
	}

	for nodeIndex := 0; nodeIndex < len(execNodes); nodeIndex++ {
		nodeTransmitterAddress := execNodes[nodeIndex].KeysBundle.EthAddress
		nodeOCR2Key := execNodes[nodeIndex].KeysBundle.OCR2Key
		nodeOCR2KeyId := nodeOCR2Key.Data.ID
		err = addExecJob(execNodes[nodeIndex], fmt.Sprintf("ccip-exec-toll-%s-%s", sourceChainName, destChainName),
			nodeTransmitterAddress, nodeOCR2KeyId, tollOffRamp, tollOnRamp)
		Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail creating CCIP-Exec-toll OCR Task job on OCR node %d", nodeIndex+1)
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
				Expect(err).ShouldNot(HaveOccurred(), "Setting mockserver value path shouldn't fail")
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
	CLNodesWithKeys          []*client.CLNodesWithKeys
	CLNodes                  []*client.Chainlink
	execNodeStartIndex       int
	commitNodeStartIndex     int
	numOfAllowedFaultyCommit int
	numOfAllowedFaultyExec   int
	SourceChainClient        blockchain.EVMClient
	DestChainClient          blockchain.EVMClient
	TestEnv                  *environment.Environment
}

func (c CCIPTestEnv) ChaosLabel() {
	for i := c.commitNodeStartIndex; i < len(c.CLNodes); i++ {
		labelSelector := map[string]string{
			"app":      "chainlink-0",
			"instance": strconv.Itoa(i),
		}
		// commit node starts from index 2
		if i >= c.commitNodeStartIndex && i < c.commitNodeStartIndex+c.numOfAllowedFaultyCommit+1 {
			err := c.TestEnv.Client.LabelChaosGroupByLabels(c.TestEnv.Cfg.Namespace, labelSelector, ChaosGroupCommitFaultyPlus)
			Expect(err).ShouldNot(HaveOccurred())
		}
		if i >= c.commitNodeStartIndex && i < c.commitNodeStartIndex+c.numOfAllowedFaultyCommit {
			err := c.TestEnv.Client.LabelChaosGroupByLabels(c.TestEnv.Cfg.Namespace, labelSelector, ChaosGroupCommitFaulty)
			Expect(err).ShouldNot(HaveOccurred())
		}
		if i >= c.execNodeStartIndex && i < c.execNodeStartIndex+c.numOfAllowedFaultyExec+1 {
			err := c.TestEnv.Client.LabelChaosGroupByLabels(c.TestEnv.Cfg.Namespace, labelSelector, ChaosGroupExecutionFaultyPlus)
			Expect(err).ShouldNot(HaveOccurred())
		}
		if i >= c.execNodeStartIndex && i < c.execNodeStartIndex+c.numOfAllowedFaultyExec {
			err := c.TestEnv.Client.LabelChaosGroupByLabels(c.TestEnv.Cfg.Namespace, labelSelector, ChaosGroupExecutionFaulty)
			Expect(err).ShouldNot(HaveOccurred())
		}
	}
}

func DeployEnvironments(
	envconfig *environment.Config,
	clProps map[string]interface{},
) *environment.Environment {
	testEnvironment := environment.New(envconfig)
	err := testEnvironment.
		AddHelm(mockservercfg.New(nil)).
		AddHelm(mockserver.New(nil)).
		AddHelm(reorg.New(&reorg.Props{
			NetworkName: sourceNetwork.Name,
			NetworkType: "simulated-geth-non-dev",
			Values: map[string]interface{}{
				"geth": map[string]interface{}{
					"genesis": map[string]interface{}{
						"networkId": fmt.Sprint(sourceNetwork.ChainID),
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
			NetworkName: destNetwork.Name,
			NetworkType: "simulated-geth-non-dev",
			Values: map[string]interface{}{
				"geth": map[string]interface{}{
					"genesis": map[string]interface{}{
						"networkId": fmt.Sprint(destNetwork.ChainID),
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
		// use blockscout for debugging on-chain transactions
		AddChart(blockscout.New(&blockscout.Props{
			Name:    "dest-blockscout",
			WsURL:   destNetwork.URLs[0],
			HttpURL: destNetwork.HTTPURLs[0],
		})).
		AddChart(blockscout.New(&blockscout.Props{
			Name:    "source-blockscout",
			WsURL:   sourceNetwork.URLs[0],
			HttpURL: sourceNetwork.HTTPURLs[0],
		})).Run()
	Expect(err).ShouldNot(HaveOccurred())
	// related https://app.shortcut.com/chainlinklabs/story/38295/creating-an-evm-chain-via-cli-or-api-immediately-polling-the-nodes-and-returning-an-error
	// node must work and reconnect even if network is not working
	time.Sleep(30 * time.Second)
	err = testEnvironment.AddHelm(chainlink.New(0, clProps)).Run()
	Expect(err).ShouldNot(HaveOccurred())
	return testEnvironment
}

func SetUpNodesAndKeys(
	testEnvironment *environment.Environment,
	nodeFund *big.Float,
) CCIPTestEnv {
	log.Info().Msg("Connecting to launched resources")
	sourceChainClient, err := blockchain.NewEVMClient(&sourceNetwork, testEnvironment)
	Expect(err).ShouldNot(HaveOccurred(), "Connecting to blockchain nodes shouldn't fail")
	destChainClient, err := blockchain.NewEVMClient(&destNetwork, testEnvironment)
	Expect(err).ShouldNot(HaveOccurred(), "Connecting to blockchain nodes shouldn't fail")

	chainlinkNodes, err := client.ConnectChainlinkNodes(testEnvironment)
	Expect(err).ShouldNot(HaveOccurred(), "Connecting to chainlink nodes shouldn't fail")
	Expect(len(chainlinkNodes)).Should(BeNumerically(">", 0), "No CL node found")

	mockServer, err := ctfClient.ConnectMockServer(testEnvironment)
	Expect(err).ShouldNot(HaveOccurred(), "Creating mockserver clients shouldn't fail")

	sourceChainClient.ParallelTransactions(true)
	destChainClient.ParallelTransactions(true)

	log.Info().Msg("creating node keys")
	_, clNodes, err := client.CreateNodeKeysBundle(chainlinkNodes, "evm", destChainClient.GetChainID().String())
	Expect(err).ShouldNot(HaveOccurred())
	Expect(len(clNodes)).Should(BeNumerically(">", 0), "No CL node with keys found")

	log.Info().Msg("Funding Chainlink nodes for both the chains")
	err = FundChainlinkNodesAddresses(chainlinkNodes, sourceChainClient, nodeFund)
	Expect(err).ShouldNot(HaveOccurred())
	err = FundChainlinkNodesAddresses(chainlinkNodes, destChainClient, nodeFund)
	Expect(err).ShouldNot(HaveOccurred())

	return CCIPTestEnv{
		MockServer:        mockServer,
		CLNodesWithKeys:   clNodes,
		CLNodes:           chainlinkNodes,
		SourceChainClient: sourceChainClient,
		DestChainClient:   destChainClient,
		TestEnv:           testEnvironment,
	}
}

func AssertBalances(bas []testhelpers.BalanceAssertion) {
	event := log.Info()
	for _, b := range bas {
		actual := b.Getter(b.Address)
		Expect(actual).ShouldNot(BeNil(), "%v getter return nil", b.Name)
		if b.Within == "" {
			Expect(actual.String()).Should(Equal(b.Expected), "wrong balance for %s got %s want %s", b.Name, actual, b.Expected)
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
			Expect(actual.Cmp(high)).Should(BeNumerically("==", -1),
				"wrong balance for %s got %s outside expected range [%s, %s]", b.Name, actual, low, high)
			Expect(actual.Cmp(low)).Should(BeNumerically("==", 1),
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

func GetterForLinkToken(token contracts.LinkToken, addr string) func(_ common.Address) *big.Int {
	return func(_ common.Address) *big.Int {
		balance, err := token.BalanceOf(context.Background(), addr)
		Expect(err).ShouldNot(HaveOccurred())
		return balance
	}
}

func CCIPDefaultTestSetUp(
	envName string,
	clProps map[string]interface{},
	numOfCommitNodes int,
	commitAndExecOnSameDON bool,
) (*environment.Environment, *SourceCCIPModule, *DestCCIPModule, CCIPTestEnv, func()) {
	testEnvironment := DeployEnvironments(
		&environment.Config{
			TTL:             4 * time.Hour,
			NamespacePrefix: envName,
		}, clProps)
	testSetUp := SetUpNodesAndKeys(testEnvironment, big.NewFloat(10))
	clNodes := testSetUp.CLNodesWithKeys
	mockServer := testSetUp.MockServer
	sourceChainClient := testSetUp.SourceChainClient
	destChainClient := testSetUp.DestChainClient
	sourceChainClient.ParallelTransactions(false)
	destChainClient.ParallelTransactions(false)

	// transfer more than one token
	transferAmounts := []*big.Int{big.NewInt(1e8)}

	// deploy all source contracts
	sourceCCIP := DefaultSourceCCIPModule(sourceChainClient, destChainClient.GetChainID(), transferAmounts)
	sourceCCIP.DeployContracts()

	// deploy all destination contracts
	destCCIP := DefaultDestinationCCIPModule(destChainClient, sourceChainClient.GetChainID())
	destCCIP.DeployContracts(*sourceCCIP)

	// set up ocr2 jobs
	var tokenAddr []string
	for _, token := range destCCIP.Common.BridgeTokens {
		tokenAddr = append(tokenAddr, token.Address())
	}
	tokenAddr = append(tokenAddr, destCCIP.Common.FeeToken.Address())
	bootstrapCommit := clNodes[0]
	var bootstrapExec *client.CLNodesWithKeys
	var execNodes []*client.CLNodesWithKeys
	commitNodes := clNodes[1:]
	testSetUp.commitNodeStartIndex = 1
	testSetUp.execNodeStartIndex = 1
	testSetUp.numOfAllowedFaultyExec = 1
	testSetUp.numOfAllowedFaultyCommit = 1
	if !commitAndExecOnSameDON {
		bootstrapExec = clNodes[1]
		commitNodes = clNodes[2 : 2+numOfCommitNodes]
		execNodes = clNodes[2+numOfCommitNodes:]
		testSetUp.commitNodeStartIndex = 2
		testSetUp.execNodeStartIndex = 7
	}
	CreateOCRJobsForCCIP(
		bootstrapCommit, bootstrapExec, commitNodes, execNodes,
		sourceCCIP.TollOnRamp.Address(),
		destCCIP.CommitStore.Address(),
		destCCIP.TollOffRamp.Address(),
		sourceChainClient, destChainClient,
		tokenAddr,
		mockServer,
	)

	// set up ocr2 config
	SetOCRConfigs(commitNodes, execNodes, *destCCIP) // first node is the bootstrapper

	tearDown := func() {
		sourceChainClient.GasStats().PrintStats()
		destChainClient.GasStats().PrintStats()
		err := TeardownSuite(testEnvironment, ctfUtils.ProjectRoot, testSetUp.CLNodes, nil, sourceChainClient, destChainClient)
		Expect(err).ShouldNot(HaveOccurred(), "Environment teardown shouldn't fail")
	}

	return testEnvironment, sourceCCIP, destCCIP, testSetUp, tearDown
}
