package actions

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
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
	"gopkg.in/guregu/null.v4"

	"github.com/smartcontractkit/chainlink/core/chains/evm/types"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/blob_verifier"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_subscription_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_toll_onramp_router"
	evm_2_evm_subscription_onramp2 "github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_subscription_onramp"
	"github.com/smartcontractkit/chainlink/core/services/job"
	ccipPlugin "github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip/testhelpers"
	"github.com/smartcontractkit/chainlink/core/services/relay"
	"github.com/smartcontractkit/chainlink/core/store/models"
	"github.com/smartcontractkit/chainlink/core/utils"
	bigmath "github.com/smartcontractkit/chainlink/core/utils/big_math"
	"github.com/smartcontractkit/chainlink/integration-tests/client"
	"github.com/smartcontractkit/chainlink/integration-tests/contracts"
	"github.com/smartcontractkit/chainlink/integration-tests/contracts/ccip"
)

type BillingModel string

const (
	SUB  BillingModel = "subscription"
	TOLL BillingModel = "toll"
)

var (
	// TODO dynamic calculation of tollfee for multiple tokens in a msg

	TollFee   = big.NewInt(0).Mul(big.NewInt(12), big.NewInt(1e18)) // "maxCharge":"10784576000000000000" for a msg with two tokens
	UnusedFee = big.NewInt(0).Mul(big.NewInt(11), big.NewInt(1e18)) // for a msg with two tokens
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

	err = ccipModule.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for Native TokenPool deployments")

	// deploy AFN
	ccipModule.AFN, err = cd.DeployAFNContract()
	Expect(err).ShouldNot(HaveOccurred(), "Deploying AFN Contract shouldn't fail")
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
	SubOnRampFee       *big.Int
	DestinationChainId *big.Int
	TollOnRampRouter   *ccip.TollOnRampRouter
	TollOnRamp         *ccip.TollOnRamp
	TollSender         *ccip.TollSender
	SubOnRamp          *ccip.SubOnRamp
	SubOnRampRouter    *ccip.SubOnRampRouter
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

	// Subscription Set up
	sourceCCIP.SubOnRampRouter, err = contractDeployer.DeploySubOnRampRouter(common.HexToAddress(sourceCCIP.Common.FeeToken.Address()))
	Expect(err).ShouldNot(HaveOccurred(), "Error on SubOnRampRouter deployment")
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for SubOnRampRouter deployment")
	sourceCCIP.SubOnRamp, err = contractDeployer.DeploySubOnRamp(
		sourceCCIP.Common.ChainClient.GetChainID(),
		sourceCCIP.DestinationChainId,
		tokens, pools, []common.Address{}, sourceCCIP.Common.AFN.EthAddress,
		sourceCCIP.SubOnRampRouter.EthAddress,
		evm_2_evm_subscription_onramp2.BaseOnRampInterfaceOnRampConfig{RelayingFeeJuels: 0, MaxDataSize: 10e12, MaxTokensLength: 5},
		evm_2_evm_subscription_onramp2.AggregateRateLimiterInterfaceRateLimiterConfig{Capacity: ccip.HundredCoins, Rate: bigmath.Mul(big.NewInt(1e18), big.NewInt(10))})
	Expect(err).ShouldNot(HaveOccurred(), "Error on SubOnRamp deployment")
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for SubOnRamp deployment")

	sourceCCIP.SubOnRampFee = big.NewInt(1)
	err = sourceCCIP.SubOnRampRouter.SetFee(sourceCCIP.SubOnRampFee)
	Expect(err).ShouldNot(HaveOccurred(), "Error setting fee on SubOnRampRouter")

	// update onRampRouter with OnRamp address
	err = sourceCCIP.SubOnRampRouter.SetOnRamp(sourceCCIP.DestinationChainId, sourceCCIP.SubOnRamp.EthAddress)
	Expect(err).ShouldNot(HaveOccurred(), "Error setting SubOnRamp on the router")

	err = sourceCCIP.SubOnRamp.SetTokenPrices(bridgeTokens, sourceCCIP.Common.TokenPrices)
	Expect(err).ShouldNot(HaveOccurred(), "Setting prices shouldn't fail")

	// update native pool with onRamp address
	for _, pool := range sourceCCIP.Common.BridgeTokenPools {
		err = pool.SetOnRamp(sourceCCIP.TollOnRamp.EthAddress)
		Expect(err).ShouldNot(HaveOccurred(), "Error setting tollonramp on the token pool %s", pool.Address())
		err = pool.SetOnRamp(sourceCCIP.SubOnRamp.EthAddress)
		Expect(err).ShouldNot(HaveOccurred(), "Error setting subonramp on the token pool %s", pool.Address())
	}
	err = sourceCCIP.Common.FeeTokenPool.SetOnRamp(sourceCCIP.TollOnRamp.EthAddress)
	Expect(err).ShouldNot(HaveOccurred(), "Error setting TollOnRamp on the token pool %s", sourceCCIP.Common.FeeTokenPool.Address())
	err = sourceCCIP.Common.FeeTokenPool.SetOnRamp(sourceCCIP.SubOnRamp.EthAddress)
	Expect(err).ShouldNot(HaveOccurred(), "Error setting SubOnRamp on the token pool %s", sourceCCIP.Common.FeeTokenPool.Address())

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
	if model == SUB {
		balancesReq = append(balancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("%s-Subscription-%s", testhelpers.Sender, sourceCCIP.Sender.Hex()),
			Addr:   sourceCCIP.Sender,
			Getter: sourceCCIP.SubscriptionBalance,
		})
	}
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
	case SUB:
		name := fmt.Sprintf("%s-Subscription-%s", testhelpers.Sender, sourceCCIP.Sender.Hex())
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     name,
			Address:  sourceCCIP.Sender,
			Getter:   sourceCCIP.SubscriptionBalance,
			Expected: bigmath.Sub(prevBalances[name], bigmath.Mul(big.NewInt(noOfreq), sourceCCIP.SubOnRampFee)).String(),
		})
	}
	return balAssertions
}

func (sourceCCIP *SourceCCIPModule) SubscriptionBalance(addr common.Address) *big.Int {
	bal, err := sourceCCIP.SubOnRampRouter.GetBalance(addr)
	Expect(err).ShouldNot(HaveOccurred())
	return bal
}

func (sourceCCIP *SourceCCIPModule) AssertEventCCIPSendRequested(model BillingModel, txHash string, seqNum, currentBlockOnSource uint64) {
	log.Info().Msg("Waiting for CCIPSendRequested event")
	switch model {
	case TOLL:
		Eventually(func(g Gomega) {
			iterator, err := sourceCCIP.TollOnRamp.FilterCCIPSendRequested(currentBlockOnSource)
			g.Expect(err).NotTo(HaveOccurred(), "Error filtering CCIPSendRequested event")
			g.Expect(iterator.Next()).To(BeTrue(), "No CCIPSendRequested event found")
			g.Expect(iterator.Event.Raw.TxHash.Hex()).Should(Equal(txHash), "txhash on CCIPSendRequested")
			g.Expect(iterator.Event.Message.SequenceNumber).Should(Equal(seqNum), "SequenceNumber on CCIPSendRequested")
		}, "1m", "1s").Should(Succeed(), "CCIPSendRequested")
	case SUB:
		Eventually(func(g Gomega) {
			iterator, err := sourceCCIP.SubOnRamp.FilterCCIPSendRequested(currentBlockOnSource)
			g.Expect(err).NotTo(HaveOccurred(), "Error filtering CCIPSendRequested event")
			g.Expect(iterator.Next()).To(BeTrue(), "No CCIPSendRequested event found")
			g.Expect(iterator.Event.Raw.TxHash.Hex()).Should(Equal(txHash), "txhash on CCIPSendRequested")
			g.Expect(iterator.Event.Message.SequenceNumber).Should(Equal(seqNum), "SequenceNumber on CCIPSendRequested")
		}, "1m", "1s").Should(Succeed(), "CCIPSendRequested")
	}
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
	BlobVerifier      *ccip.BlobVerifier
	TollOffRamp       *ccip.TollOffRamp
	TollOffRampRouter *ccip.TollOffRampRouter
	SubOffRamp        *ccip.SubOffRamp
	SubOffRampRouter  *ccip.SubOffRampRouter
	ReceiverDapp      *ccip.ReceiverDapp
}

// DeployContracts deploys all CCIP contracts specific to the destination chain
func (destCCIP *DestCCIPModule) DeployContracts(sourceCCIP SourceCCIPModule) {
	var err error
	destCCIP.Common.Deployer, err = ccip.NewCCIPContractsDeployer(destCCIP.Common.ChainClient)
	Expect(err).ShouldNot(HaveOccurred(), "contract deployer should be created successfully")
	contractDeployer := destCCIP.Common.Deployer
	destCCIP.Common.DeployContracts(contractDeployer, len(sourceCCIP.TransferAmount))

	// blobVerifier responsible for validating the transfer message
	destCCIP.BlobVerifier, err = contractDeployer.DeployBlobVerifier(
		destCCIP.SourceChainId,
		destCCIP.Common.ChainClient.GetChainID(),
		destCCIP.Common.AFN.EthAddress,
		blob_verifier.BlobVerifierInterfaceBlobVerifierConfig{
			OnRamps:          []common.Address{sourceCCIP.TollOnRamp.EthAddress, sourceCCIP.SubOnRamp.EthAddress},
			MinSeqNrByOnRamp: []uint64{1, 1},
		})
	Expect(err).ShouldNot(HaveOccurred(), "Deploying BlobVerifier shouldn't fail")
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
		err = token.Transfer(pool.Address(), ccip.HundredCoins)
		Expect(err).ShouldNot(HaveOccurred())
	}
	// add the fee token and feetoken price for dest
	destTokens = append(destTokens, common.HexToAddress(destCCIP.Common.FeeToken.Address()))
	destCCIP.Common.TokenPrices = append(destCCIP.Common.TokenPrices, big.NewInt(1))

	pools = append(pools, destCCIP.Common.FeeTokenPool.EthAddress)
	err = destCCIP.Common.FeeToken.Transfer(destCCIP.Common.FeeTokenPool.Address(), ccip.HundredCoins)
	Expect(err).ShouldNot(HaveOccurred())

	// Toll
	// offRamp
	destCCIP.TollOffRamp, err = contractDeployer.DeployTollOffRamp(destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID(),
		destCCIP.BlobVerifier.EthAddress, sourceCCIP.TollOnRamp.EthAddress, destCCIP.Common.AFN.EthAddress,
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

	// subscription
	destCCIP.SubOffRamp, err = contractDeployer.DeploySubOffRamp(
		destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID(), destCCIP.BlobVerifier.EthAddress,
		sourceCCIP.SubOnRamp.EthAddress, destCCIP.Common.AFN.EthAddress, sourceTokens, pools,
		destCCIP.Common.RateLimiterConfig,
		any_2_evm_subscription_offramp.BaseOffRampInterfaceOffRampConfig{ExecutionDelaySeconds: 0, MaxDataSize: 10e12, MaxTokensLength: 15})
	Expect(err).ShouldNot(HaveOccurred(), "Deploying SubOffRamp shouldn't fail")
	err = destCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for deploying SubOffRamp")
	err = destCCIP.SubOffRamp.SetTokenPrices(destTokens, destCCIP.Common.TokenPrices)
	destCCIP.SubOffRampRouter, err = contractDeployer.DeploySubOffRampRouter(
		[]common.Address{destCCIP.SubOffRamp.EthAddress}, common.HexToAddress(destCCIP.Common.FeeToken.Address()))
	Expect(err).ShouldNot(HaveOccurred(), "Deploying SubOffRampRouter shouldn't fail")
	err = destCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for deploying SubOffRampRouter")
	err = destCCIP.SubOffRamp.SetRouter(destCCIP.SubOffRampRouter.EthAddress)
	Expect(err).ShouldNot(HaveOccurred(), "Error setting router on the offramp")

	// ReceiverDapp
	destCCIP.ReceiverDapp, err = contractDeployer.DeployReceiverDapp(false)
	Expect(err).ShouldNot(HaveOccurred(), "ReceiverDapp contract should be deployed successfully")

	// update pools with offRamp Id
	for _, pool := range destCCIP.Common.BridgeTokenPools {
		err = pool.SetOffRamp(destCCIP.TollOffRamp.EthAddress)
		Expect(err).ShouldNot(HaveOccurred(), "Error setting TollOffRamp on the token pool %s", pool.Address())
		err = pool.SetOffRamp(destCCIP.SubOffRamp.EthAddress)
		Expect(err).ShouldNot(HaveOccurred(), "Error setting SubOffRamp on the token pool %s", pool.Address())
	}
	err = destCCIP.Common.FeeTokenPool.SetOffRamp(destCCIP.TollOffRamp.EthAddress)
	Expect(err).ShouldNot(HaveOccurred(), "Error setting TollOffRamp on the token pool %s", destCCIP.Common.FeeTokenPool.Address())
	err = destCCIP.Common.FeeTokenPool.SetOffRamp(destCCIP.SubOffRamp.EthAddress)
	Expect(err).ShouldNot(HaveOccurred(), "Error setting SubOffRamp on the token pool %s", destCCIP.Common.FeeTokenPool.Address())

	err = sourceCCIP.Common.ChainClient.WaitForEvents()
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
	if model == SUB {
		destBalancesReq = append(destBalancesReq, testhelpers.BalanceReq{
			Name:   fmt.Sprintf("%s-Subscription-%s", testhelpers.Receiver, destCCIP.ReceiverDapp.Address()),
			Addr:   destCCIP.ReceiverDapp.EthAddress,
			Getter: destCCIP.SubscriptionBalance,
		})
	}
	return destBalancesReq
}

func (destCCIP *DestCCIPModule) SubscriptionBalance(addr common.Address) *big.Int {
	sub, err := destCCIP.SubOffRampRouter.GetSubscription(addr)
	Expect(err).ShouldNot(HaveOccurred())
	return sub
}

func (destCCIP *DestCCIPModule) BalanceAssertions(model BillingModel, prevBalances map[string]*big.Int, transferAmount []*big.Int, unusedFee *big.Int, noOfReq int64) []testhelpers.BalanceAssertion {
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
	if model == SUB {
		name := fmt.Sprintf("%s-Subscription-%s", testhelpers.Receiver, destCCIP.ReceiverDapp.Address())
		log.Info().Interface("PrevBalance", prevBalances[name]).Msg(name)
		balAssertions = append(balAssertions, testhelpers.BalanceAssertion{
			Name:     name,
			Address:  destCCIP.ReceiverDapp.EthAddress,
			Getter:   destCCIP.SubscriptionBalance,
			Expected: bigmath.Sub(prevBalances[name], bigmath.Mul(big.NewInt(noOfReq), big.NewInt(0.79e18))).String(),
			Within:   big.NewInt(0.1e18).String(),
		})
	}
	return balAssertions
}

func (destCCIP *DestCCIPModule) AssertEventExecutionStateChanged(model BillingModel, prevSeqNum uint64, currentBlockOnDest uint64) {
	log.Info().Msg("Waiting for ExecutionStateChanged event")
	switch model {
	case TOLL:
		Eventually(func(g Gomega) ccipPlugin.MessageExecutionState {
			iterator, err := destCCIP.TollOffRamp.FilterExecutionStateChanged([]uint64{prevSeqNum}, currentBlockOnDest)
			g.Expect(err).NotTo(HaveOccurred(), "Error filtering ExecutionStateChanged event")
			g.Expect(iterator.Next()).To(BeTrue(), "No ExecutionStateChanged event found")
			return ccipPlugin.MessageExecutionState(iterator.Event.State)
		}, "1m", "1s").Should(Equal(ccipPlugin.Success))
	case SUB:
		Eventually(func(g Gomega) ccipPlugin.MessageExecutionState {
			iterator, err := destCCIP.SubOffRamp.FilterExecutionStateChanged([]uint64{prevSeqNum}, currentBlockOnDest)
			g.Expect(err).NotTo(HaveOccurred(), "Error filtering ExecutionStateChanged event")
			g.Expect(iterator.Next()).To(BeTrue(), "No ExecutionStateChanged event found")
			return ccipPlugin.MessageExecutionState(iterator.Event.State)
		}, "1m", "1s").Should(Equal(ccipPlugin.Success), "Execution should be successful")
	}
}

func (destCCIP *DestCCIPModule) AssertEventReportAccepted(onRamp []common.Address, currentBlockOnDest uint64) {
	log.Info().Msg("Waiting for ReportAccepted event")
	Eventually(func(g Gomega) []common.Address {
		iterator, err := destCCIP.BlobVerifier.FilterReportAccepted(currentBlockOnDest)
		g.Expect(err).NotTo(HaveOccurred(), "Error filtering ReportAccepted event")
		g.Expect(iterator.Next()).To(BeTrue(), "No ReportAccepted event found")
		return iterator.Event.Report.OnRamps
	}, "1m", "1s").Should(ContainElements(onRamp))
}

func (destCCIP *DestCCIPModule) AssertSeqNumberExecuted(onRamp common.Address, seqNumberBefore uint64) {
	log.Info().Msg("Waiting for SeqNumber to be executed")
	Eventually(func(g Gomega) {
		seqNumberAfter, err := destCCIP.BlobVerifier.GetNextSeqNumber(onRamp)
		g.Expect(err).ShouldNot(HaveOccurred(), "Getting expected seq number should be successful")
		g.Expect(seqNumberAfter).Should(BeNumerically(">", seqNumberBefore), "Next Sequence number is not increased")
	}, "1m", "1s").Should(Succeed(), "Error Executing Sequence number")
}

func DefaultDestinationCCIPModule(chainClient blockchain.EVMClient, sourceChain *big.Int) *DestCCIPModule {
	return &DestCCIPModule{
		Common:        DefaultCCIPModule(chainClient),
		SourceChainId: sourceChain,
	}
}

// TokenTransferWithToll initiates transfer of token with an amount as defined in sourceCCIP module. It waits for
// CCIPSendRequestedEvent, ReportRelayedEvent and ExecutionStateChanged Event to ensure that transfer has taken place
// and verifies senders and receiver's balance pre- and post-transfer
func TokenTransferWithToll(sourceCCIP SourceCCIPModule, destCCIP DestCCIPModule) {
	var sourceTokens []common.Address
	for i, token := range sourceCCIP.Common.BridgeTokens {
		sourceTokens = append(sourceTokens, common.HexToAddress(token.Address()))
		// approve the onramp router so that it can initiate transferring the token
		err := token.Approve(sourceCCIP.TollOnRampRouter.Address(), sourceCCIP.TransferAmount[i])
		Expect(err).ShouldNot(HaveOccurred(), "Could not approve permissions for the onRamp router "+
			"on the source link token contract")
	}
	// approve onramp router to use the feetoken
	err := sourceCCIP.Common.FeeToken.Approve(sourceCCIP.TollOnRampRouter.Address(), sourceCCIP.TollFeeAmount)
	Expect(err).ShouldNot(HaveOccurred(), "Could not approve permissions for the onRamp router "+
		"on the source link token contract")
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Failed to wait for events")

	// collect the balance requirement to verify balances after transfer
	sourceBalances, err := testhelpers.GetBalances(sourceCCIP.CollectBalanceRequirements(TOLL))
	Expect(err).ShouldNot(HaveOccurred(), "fetching source balance")
	destBalances, err := testhelpers.GetBalances(destCCIP.CollectBalanceRequirements(TOLL))
	Expect(err).ShouldNot(HaveOccurred(), "fetching dest balance")

	// save the current block numbers to use in various filter log requests
	currentBlockOnSource, err := sourceCCIP.Common.ChainClient.LatestBlockNumber(context.Background())
	Expect(err).ShouldNot(HaveOccurred(), "Getting current block should be successful in source chain")
	currentBlockOnDest, err := destCCIP.Common.ChainClient.LatestBlockNumber(context.Background())
	Expect(err).ShouldNot(HaveOccurred(), "Getting current block should be successful in dest chain")

	// save the next seq number to compare with NextSeqNumber generated after transfer
	seqNumberBefore, err := destCCIP.BlobVerifier.GetNextSeqNumber(sourceCCIP.TollOnRamp.EthAddress)
	Expect(err).ShouldNot(HaveOccurred(), "Getting expected seq number should be successful")

	receiver, err := utils.ABIEncode(`[{"type":"address"}]`, destCCIP.ReceiverDapp.EthAddress)
	Expect(err).ShouldNot(HaveOccurred(), "Failed encoding the receiver address")

	// form the message for transfer
	msg := evm_2_any_toll_onramp_router.CCIPEVM2AnyTollMessage{
		Receiver:       receiver,
		Data:           []byte("Token transfer by DON"),
		Tokens:         sourceTokens,
		Amounts:        sourceCCIP.TransferAmount,
		FeeToken:       common.HexToAddress(sourceCCIP.Common.FeeToken.Address()),
		FeeTokenAmount: sourceCCIP.TollFeeAmount,
		GasLimit:       big.NewInt(100_000),
	}
	log.Info().Interface("msg details", msg).Msg("ccip message to be sent")

	// initiate the transfer
	sendTx, err := sourceCCIP.TollOnRampRouter.CCIPSend(destCCIP.Common.ChainClient.GetChainID(), msg)
	Expect(err).ShouldNot(HaveOccurred(), "send token should be initiated successfully")
	log.Info().Str("send token transaction", sendTx.Hash().String()).Msg("Sending token")
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Failed to wait for events")

	// Verify if
	// - CCIPSendRequested Event log generated,
	// - NextSeqNumber from blobVerifier got increased
	sourceCCIP.AssertEventCCIPSendRequested(TOLL, sendTx.Hash().Hex(), seqNumberBefore, currentBlockOnSource)
	destCCIP.AssertSeqNumberExecuted(sourceCCIP.TollOnRamp.EthAddress, seqNumberBefore)

	// Verify whether blobVerifier has accepted the report
	destCCIP.AssertEventReportAccepted([]common.Address{sourceCCIP.TollOnRamp.EthAddress}, currentBlockOnDest)

	// Verify whether the execution state is changed and the transfer is successful
	destCCIP.AssertEventExecutionStateChanged(TOLL, seqNumberBefore, currentBlockOnDest)

	// verify the fee amount is deducted from sender, added to receiver token balances and
	// unused fee is returned to receiver fee token account
	AssertBalances(sourceCCIP.BalanceAssertions(TOLL, sourceBalances, 1))
	AssertBalances(destCCIP.BalanceAssertions(TOLL, destBalances, sourceCCIP.TransferAmount, UnusedFee, 1))
}

func CreateAndFundSubscription(sourceCCIP SourceCCIPModule, destCCIP DestCCIPModule, subscriptionBalance *big.Int, noOfReq int64) {
	relayFee := big.NewInt(0).Mul(sourceCCIP.SubOnRampFee, big.NewInt(noOfReq)) // for noOfReq requests
	err := sourceCCIP.Common.FeeToken.Approve(sourceCCIP.SubOnRampRouter.Address(), relayFee)
	Expect(err).ShouldNot(HaveOccurred(), "Error approving fee token")
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Failed to wait for events")
	err = destCCIP.Common.FeeToken.Approve(destCCIP.SubOffRampRouter.Address(), subscriptionBalance)
	Expect(err).ShouldNot(HaveOccurred(), "Error approving fee token")
	err = destCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Failed to wait for events")
	err = sourceCCIP.SubOnRampRouter.FundSubscription(relayFee)
	Expect(err).ShouldNot(HaveOccurred(), "Error funding subscription on SubOnRampRouter")
	err = destCCIP.SubOffRampRouter.CreateSubscription([]common.Address{sourceCCIP.Sender}, destCCIP.ReceiverDapp.EthAddress, false, subscriptionBalance)
	Expect(err).ShouldNot(HaveOccurred(), "Error creating subscription on SubOffRampRouter")
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Failed to wait for creating and funding subscription")
}

func TokenTransferWithSub(sourceCCIP SourceCCIPModule, destCCIP DestCCIPModule) {
	var sourceTokens []common.Address
	for i, token := range sourceCCIP.Common.BridgeTokens {
		sourceTokens = append(sourceTokens, common.HexToAddress(token.Address()))
		// approve the onramp router so that it can initiate transferring the token
		err := token.Approve(sourceCCIP.SubOnRampRouter.Address(), sourceCCIP.TransferAmount[i])
		Expect(err).ShouldNot(HaveOccurred(), "Could not approve permissions for the onRamp router "+
			"on the source link token contract")
	}
	err := sourceCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Failed to wait for events")

	// collect the balance requirement to verify balances after transfer
	sourceBalances, err := testhelpers.GetBalances(sourceCCIP.CollectBalanceRequirements(SUB))
	Expect(err).ShouldNot(HaveOccurred(), "fetching source balance")
	destBalances, err := testhelpers.GetBalances(destCCIP.CollectBalanceRequirements(SUB))
	Expect(err).ShouldNot(HaveOccurred(), "fetching dest balance")

	// save the current block numbers to use in various filter log requests
	currentBlockOnSource, err := sourceCCIP.Common.ChainClient.LatestBlockNumber(context.Background())
	Expect(err).ShouldNot(HaveOccurred(), "Getting current block should be successful in source chain")
	currentBlockOnDest, err := destCCIP.Common.ChainClient.LatestBlockNumber(context.Background())
	Expect(err).ShouldNot(HaveOccurred(), "Getting current block should be successful in dest chain")

	// save the next seq number to compare with NextSeqNumber generated after transfer
	seqNumber, err := destCCIP.BlobVerifier.GetNextSeqNumber(sourceCCIP.SubOnRamp.EthAddress)
	Expect(err).ShouldNot(HaveOccurred(), "Getting expected seq number should be successful")

	receiver, err := utils.ABIEncode(`[{"type":"address"}]`, destCCIP.ReceiverDapp.EthAddress)
	Expect(err).ShouldNot(HaveOccurred(), "Failed encoding the receiver address")

	// form the message for transfer
	msg := evm_2_any_subscription_onramp_router.CCIPEVM2AnySubscriptionMessage{
		Receiver: receiver,
		Data:     []byte("Token transfer subscription model"),
		Tokens:   sourceTokens,
		Amounts:  sourceCCIP.TransferAmount,
		GasLimit: big.NewInt(100_000),
	}
	log.Info().Interface("msg details", msg).Msg("ccip message to be sent")

	// initiate the transfer
	sendTx, err := sourceCCIP.SubOnRampRouter.CCIPSend(destCCIP.Common.ChainClient.GetChainID(), msg)
	Expect(err).ShouldNot(HaveOccurred(), "CCIPSend should be initiated successfully")
	log.Info().Str("send token transaction", sendTx.Hash().String()).Msg("Sending token")
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Failed to wait for events")

	// Verify if
	// - CCIPSendRequested Event log generated,
	// - NextSeqNumber from blobVerifier got increased
	sourceCCIP.AssertEventCCIPSendRequested(SUB, sendTx.Hash().Hex(), seqNumber, currentBlockOnSource)
	destCCIP.AssertSeqNumberExecuted(sourceCCIP.SubOnRamp.EthAddress, seqNumber)

	// Verify whether blobVerifier has accepted the report
	destCCIP.AssertEventReportAccepted([]common.Address{sourceCCIP.SubOnRamp.EthAddress}, currentBlockOnDest)

	// Verify whether the execution state is changed and the transfer is successful
	destCCIP.AssertEventExecutionStateChanged(SUB, seqNumber, currentBlockOnDest)

	// no change forwarding for sub
	unusedFee := big.NewInt(0)
	AssertBalances(sourceCCIP.BalanceAssertions(SUB, sourceBalances, 1))
	AssertBalances(destCCIP.BalanceAssertions(SUB, destBalances, sourceCCIP.TransferAmount, unusedFee, 1))
}

func SetOCRConfigs(chainlinkNodes []*client.CLNodesWithKeys, destCCIP DestCCIPModule) {
	signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, err :=
		ccip.NewOffChainAggregatorV2Config(chainlinkNodes)
	Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail while getting the config values for ocr2 type contract")
	err = destCCIP.BlobVerifier.SetOCRConfig(signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
	Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail while setting blobverifier config")
	err = destCCIP.TollOffRamp.SetConfig(signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
	Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail while setting TollOffRamp config")
	err = destCCIP.SubOffRamp.SetConfig(signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
	Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail while setting SubOffRamp config")
	err = destCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail while waiting for events on setting ocr2 config")
}

// CreateOCRJobsForCCIP bootstraps the first node and to the other nodes sends ocr jobs that
// sets up ccip-relay and ccip-execution plugin
func CreateOCRJobsForCCIP(
	chainlinkNodes []*client.CLNodesWithKeys,
	tollOnRamp, subOnRamp, blobVerifier, tollOffRamp, subOffRamp string,
	sourceChainClient, destChainClient blockchain.EVMClient,
	linkTokenAddr []string,
	mockServer *ctfClient.MockserverClient,
) {
	bootstrapNodeWithKey := chainlinkNodes[0]
	bootstrapNode := chainlinkNodes[0].Node
	bootstrapP2PIds := bootstrapNodeWithKey.KeysBundle.P2PKeys
	bootstrapP2PId := bootstrapP2PIds.Data[0].Attributes.PeerID
	sourceChainID := sourceChainClient.GetChainID()
	destChainID := destChainClient.GetChainID()
	sourceChainName := sourceChainClient.GetNetworkName()
	destChainName := destChainClient.GetNetworkName()
	bootstrapSpec := &client.OCR2TaskJobSpec{
		Name:    fmt.Sprintf("bootstrap-%s-%s", destChainName, uuid.NewV4().String()),
		JobType: "bootstrap",
		OCR2OracleSpec: job.OCR2OracleSpec{
			ContractID:                        blobVerifier,
			Relay:                             relay.EVM,
			ContractConfigConfirmations:       1,
			ContractConfigTrackerPollInterval: models.Interval(1 * time.Second),
			RelayConfig: map[string]interface{}{
				"chainID": fmt.Sprintf("\"%s\"", destChainID.String()),
			},
		},
	}
	_, err := bootstrapNode.MustCreateJob(bootstrapSpec)
	Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail creating bootstrap job on bootstrap node")
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
	SetMockServerWithSameTokenFeeConversionValue(tokenFeeConv, chainlinkNodes[1:], mockServer)

	for nodeIndex := 1; nodeIndex < len(chainlinkNodes); nodeIndex++ {
		nodeTransmitterAddress := chainlinkNodes[nodeIndex].KeysBundle.EthAddress
		nodeOCR2Key := chainlinkNodes[nodeIndex].KeysBundle.OCR2Key
		nodeOCR2KeyId := nodeOCR2Key.Data.ID
		ocr2SpecRelay := &client.OCR2TaskJobSpec{
			JobType: "offchainreporting2",
			Name:    fmt.Sprintf("ccip-relay-%s-%s", sourceChainName, destChainName),
			OCR2OracleSpec: job.OCR2OracleSpec{
				Relay:                             relay.EVM,
				PluginType:                        job.CCIPRelay,
				ContractID:                        blobVerifier,
				OCRKeyBundleID:                    null.StringFrom(nodeOCR2KeyId),
				TransmitterID:                     null.StringFrom(nodeTransmitterAddress),
				ContractConfigConfirmations:       1,
				ContractConfigTrackerPollInterval: models.Interval(1 * time.Second),
				P2PV2Bootstrappers: []string{
					client.P2PData{
						RemoteIP: bootstrapNode.RemoteIP(),
						PeerID:   bootstrapP2PId,
					}.P2PV2Bootstrapper(),
				},
				PluginConfig: map[string]interface{}{
					"sourceChainID":    sourceChainID,
					"destChainID":      destChainID,
					"onRampIDs":        fmt.Sprintf("[\"%s\",\"%s\"]", tollOnRamp, subOnRamp),
					"pollPeriod":       `"1s"`,
					"destStartBlock":   currentBlockOnDest,
					"sourceStartBlock": currentBlockOnSource,
				},
				RelayConfig: map[string]interface{}{
					"chainID": destChainID,
				},
			},
		}
		_, err = chainlinkNodes[nodeIndex].Node.MustCreateJob(ocr2SpecRelay)
		Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail creating CCIP-Relay OCR Task job on OCR node %d", nodeIndex+1)

		tokensPerFeeCoinPipeline := TokenFeeForMultipleTokenAddr(chainlinkNodes[nodeIndex], linkTokenAddr, mockServer)
		ocr2SpecExec := &client.OCR2TaskJobSpec{
			JobType: "offchainreporting2",
			Name:    fmt.Sprintf("ccip-exec-toll-%s-%s", sourceChainName, destChainName),
			OCR2OracleSpec: job.OCR2OracleSpec{
				Relay:                             relay.EVM,
				PluginType:                        job.CCIPExecution,
				ContractID:                        tollOffRamp,
				OCRKeyBundleID:                    null.StringFrom(nodeOCR2KeyId),
				TransmitterID:                     null.StringFrom(nodeTransmitterAddress),
				ContractConfigConfirmations:       1,
				ContractConfigTrackerPollInterval: models.Interval(1 * time.Second),
				P2PV2Bootstrappers: []string{
					client.P2PData{
						RemoteIP: bootstrapNode.RemoteIP(),
						PeerID:   bootstrapP2PId,
					}.P2PV2Bootstrapper(),
				},
				PluginConfig: map[string]interface{}{
					"sourceChainID":    sourceChainID,
					"destChainID":      destChainID,
					"onRampID":         fmt.Sprintf("\"%s\"", tollOnRamp),
					"blobVerifierID":   fmt.Sprintf("\"%s\"", blobVerifier),
					"pollPeriod":       `"1s"`,
					"destStartBlock":   currentBlockOnDest,
					"sourceStartBlock": currentBlockOnSource,
					"tokensPerFeeCoinPipeline": fmt.Sprintf(`"""
		%s
		"""`, tokensPerFeeCoinPipeline),
				},
				RelayConfig: map[string]interface{}{
					"chainID": destChainID,
				},
			},
		}
		_, err = chainlinkNodes[nodeIndex].Node.MustCreateJob(ocr2SpecExec)
		Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail creating CCIP-Exec-toll OCR Task job on OCR node %d", nodeIndex+1)

		ocr3SpecExec := &client.OCR2TaskJobSpec{
			JobType: "offchainreporting2",
			Name:    fmt.Sprintf("ccip-exec-sub-%s-%s", sourceChainName, destChainName),
			OCR2OracleSpec: job.OCR2OracleSpec{
				Relay:                             relay.EVM,
				PluginType:                        job.CCIPExecution,
				ContractID:                        subOffRamp,
				OCRKeyBundleID:                    null.StringFrom(nodeOCR2KeyId),
				TransmitterID:                     null.StringFrom(nodeTransmitterAddress),
				ContractConfigConfirmations:       1,
				ContractConfigTrackerPollInterval: models.Interval(1 * time.Second),
				P2PV2Bootstrappers: []string{
					client.P2PData{
						RemoteIP: bootstrapNode.RemoteIP(),
						PeerID:   bootstrapP2PId,
					}.P2PV2Bootstrapper(),
				},
				PluginConfig: map[string]interface{}{
					"sourceChainID":    sourceChainID,
					"destChainID":      destChainID,
					"onRampID":         fmt.Sprintf("\"%s\"", subOnRamp),
					"blobVerifierID":   fmt.Sprintf("\"%s\"", blobVerifier),
					"pollPeriod":       `"0.5s"`,
					"destStartBlock":   currentBlockOnDest,
					"sourceStartBlock": currentBlockOnSource,
					"tokensPerFeeCoinPipeline": fmt.Sprintf(`"""
%s
"""`, tokensPerFeeCoinPipeline),
				},
				RelayConfig: map[string]interface{}{
					"chainID": destChainID,
				},
			},
		}
		_, err = chainlinkNodes[nodeIndex].Node.MustCreateJob(ocr3SpecExec)
		Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail creating CCIP-Exec-sub OCR Task job on OCR node %d", nodeIndex+1)
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

type CCIPTest struct {
	MockServer        *ctfClient.MockserverClient
	CLNodesWithKeys   []*client.CLNodesWithKeys
	CLNodes           []*client.Chainlink
	SourceChainClient blockchain.EVMClient
	DestChainClient   blockchain.EVMClient
	TestEnv           *environment.Environment
}

func DeployEnvironments(sourceNetwork *blockchain.EVMNetwork, destNetwork *blockchain.EVMNetwork, envconfig *environment.Config) *environment.Environment {
	evmNodes, err := json.Marshal([]types.NewNode{
		{
			Name:       "primary_0_source",
			EVMChainID: *utils.NewBigI(sourceNetwork.ChainID),
			WSURL:      null.StringFrom(sourceNetwork.URLs[0]),
			HTTPURL:    null.StringFrom(sourceNetwork.HTTPURLs[0]),
			SendOnly:   false,
		},
		{
			Name:       "primary_0_dest",
			EVMChainID: *utils.NewBigI(destNetwork.ChainID),
			WSURL:      null.StringFrom(destNetwork.URLs[0]),
			HTTPURL:    null.StringFrom(destNetwork.HTTPURLs[0]),
			SendOnly:   false,
		},
	})
	Expect(err).ShouldNot(HaveOccurred())

	testEnvironment := environment.New(envconfig)
	err = testEnvironment.
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
	err = testEnvironment.AddHelm(chainlink.New(0, map[string]interface{}{
		"replicas": 6,
		"env": map[string]interface{}{
			"FEATURE_CCIP":                   "true",
			"FEATURE_OFFCHAIN_REPORTING2":    "true",
			"feature_offchain_reporting":     "false",
			"FEATURE_LOG_POLLER":             "true",
			"GAS_ESTIMATOR_MODE":             "FixedPrice",
			"P2P_NETWORKING_STACK":           "V2",
			"P2PV2_LISTEN_ADDRESSES":         "0.0.0.0:6690",
			"P2PV2_ANNOUNCE_ADDRESSES":       "0.0.0.0:6690",
			"P2PV2_DELTA_DIAL":               "500ms",
			"P2PV2_DELTA_RECONCILE":          "5s",
			"ETH_GAS_LIMIT_DEFAULT":          "5000000",
			"ETH_LOG_POLL_INTERVAL":          "1s",
			"p2p_listen_port":                "0",
			"ETH_FINALITY_DEPTH":             "50",
			"ETH_HEAD_TRACKER_HISTORY_DEPTH": "100",
			// It is not permitted to set both ETH_URL and EVM_NODES,
			// imposing blank values to stop getting the env variable set as default node set up in qa-charts
			"ETH_URL":      "",
			"ETH_CHAIN_ID": "0",
			"EVM_NODES":    string(evmNodes),
		},
	})).Run()
	Expect(err).ShouldNot(HaveOccurred())
	return testEnvironment
}

func SetUpNodesAndKeys(
	sourceNetwork *blockchain.EVMNetwork,
	destNetwork *blockchain.EVMNetwork,
	testEnvironment *environment.Environment,
	nodeFund *big.Float,
) CCIPTest {
	log.Info().Msg("Connecting to launched resources")
	sourceChainClient, err := blockchain.NewEVMClient(sourceNetwork, testEnvironment)
	Expect(err).ShouldNot(HaveOccurred(), "Connecting to blockchain nodes shouldn't fail")
	destChainClient, err := blockchain.NewEVMClient(destNetwork, testEnvironment)
	Expect(err).ShouldNot(HaveOccurred(), "Connecting to blockchain nodes shouldn't fail")

	chainlinkNodes, err := client.ConnectChainlinkNodes(testEnvironment)
	Expect(err).ShouldNot(HaveOccurred(), "Connecting to chainlink nodes shouldn't fail")

	mockServer, err := ctfClient.ConnectMockServer(testEnvironment)
	Expect(err).ShouldNot(HaveOccurred(), "Creating mockserver clients shouldn't fail")

	sourceChainClient.ParallelTransactions(true)
	destChainClient.ParallelTransactions(true)

	log.Info().Msg("Funding Chainlink nodes for both the chains")
	err = FundChainlinkNodesForChain(chainlinkNodes, sourceChainClient, nodeFund)
	Expect(err).ShouldNot(HaveOccurred())
	err = FundChainlinkNodesForChain(chainlinkNodes, destChainClient, nodeFund)
	Expect(err).ShouldNot(HaveOccurred())

	// create node keys
	_, clNodes, err := client.CreateNodeKeysBundle(chainlinkNodes, "evm", destChainClient.GetChainID().String())
	Expect(err).ShouldNot(HaveOccurred())
	return CCIPTest{
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

func CCIPDefaultTestSetUpForLoad(sourceNetwork *blockchain.EVMNetwork, destNetwork *blockchain.EVMNetwork, envName string) (*SourceCCIPModule, *DestCCIPModule, func()) {
	testEnvironment := DeployEnvironments(sourceNetwork, destNetwork, &environment.Config{
		TTL:             1 * time.Hour,
		NamespacePrefix: "load-ccip",
	})
	testSetUp := SetUpNodesAndKeys(sourceNetwork, destNetwork, testEnvironment, big.NewFloat(10))
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

	CreateOCRJobsForCCIP(
		clNodes, sourceCCIP.TollOnRamp.Address(),
		sourceCCIP.SubOnRamp.Address(),
		destCCIP.BlobVerifier.Address(),
		destCCIP.TollOffRamp.Address(),
		destCCIP.SubOffRamp.Address(),
		sourceChainClient, destChainClient,
		tokenAddr,
		mockServer,
	)

	// set up ocr2 config
	SetOCRConfigs(clNodes[1:], *destCCIP) // first node is the bootstrapper
	tearDown := func() {
		sourceChainClient.GasStats().PrintStats()
		destChainClient.GasStats().PrintStats()
		err := TeardownSuite(testEnvironment, ctfUtils.ProjectRoot, testSetUp.CLNodes, nil, sourceChainClient)
		Expect(err).ShouldNot(HaveOccurred(), "Environment teardown shouldn't fail")
	}
	return sourceCCIP, destCCIP, tearDown
}
