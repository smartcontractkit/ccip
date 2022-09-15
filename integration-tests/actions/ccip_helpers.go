package actions

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"
	ctfClient "github.com/smartcontractkit/chainlink-testing-framework/client"
	"gopkg.in/guregu/null.v4"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/blob_verifier"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_toll_onramp_router"
	"github.com/smartcontractkit/chainlink/core/services/job"
	ccipPlugin "github.com/smartcontractkit/chainlink/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/core/services/relay"
	"github.com/smartcontractkit/chainlink/core/store/models"
	bigmath "github.com/smartcontractkit/chainlink/core/utils/big_math"
	"github.com/smartcontractkit/chainlink/integration-tests/client"
	"github.com/smartcontractkit/chainlink/integration-tests/contracts"
	"github.com/smartcontractkit/chainlink/integration-tests/contracts/ccip"
)

type CCIPCommon struct {
	ChainClient       blockchain.EVMClient
	Deployer          *ccip.CCIPContractsDeployer
	LinkToken         contracts.LinkToken
	BridgeTokens      []contracts.LinkToken // as of now considering the bridge token is same as link token
	BridgeTokenPrices []*big.Int
	NativeTokenPools  []*ccip.NativeTokenPool
	RateLimiterConfig ccip.RateLimiterConfig
	AFNConfig         ccip.AFNConfig
	AFN               *ccip.AFN
	TransferAmount    *big.Int
}

// DeployContracts deploys the contracts which are necessary in both source and dest chain
func (ccipModule *CCIPCommon) DeployContracts(cd *ccip.CCIPContractsDeployer) {
	// deploy link token
	token, err := cd.DeployLinkTokenContract()
	Expect(err).ShouldNot(HaveOccurred(), "Deploying Link Token Contract shouldn't fail")
	ccipModule.LinkToken = token
	// deploy bridge token. as of now keeping the bridge token same as link token,
	ccipModule.BridgeTokens = []contracts.LinkToken{token}
	// Set price of the bridge tokens to 1
	ccipModule.BridgeTokenPrices = []*big.Int{big.NewInt(1)}

	err = ccipModule.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for Link Token deployments")
	// deploy native token pool
	for _, token = range ccipModule.BridgeTokens {
		ntp, err := cd.DeployNativeTokenPoolContract(token.Address())
		Expect(err).ShouldNot(HaveOccurred(), "Deploying Native TokenPool Contract shouldn't fail")
		ccipModule.NativeTokenPools = append(ccipModule.NativeTokenPools, ntp)
		err = ccipModule.ChainClient.WaitForEvents()
		Expect(err).ShouldNot(HaveOccurred(), "Error waiting for Native TokenPool deployments")
	}
	// deploy AFN
	ccipModule.AFN, err = cd.DeployAFNContract(
		ccipModule.AFNConfig.AFNWeightsByParticipants, ccipModule.AFNConfig.ThresholdForBlessing, ccipModule.AFNConfig.ThresholdForBadSignal)
	Expect(err).ShouldNot(HaveOccurred(), "Deploying AFN Contract shouldn't fail")
}

func DefaultCCIPModule(chainClient blockchain.EVMClient, transferamount *big.Int) *CCIPCommon {
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
		TransferAmount: transferamount,
	}
}

type SourceCCIPModule struct {
	Common             *CCIPCommon
	FeeToken           *big.Int
	DestinationChainId *big.Int
	OnRampRouter       *ccip.OnRampRouter
	OnRamp             *ccip.OnRamp
	TollSender         *ccip.TollSender
}

// DeploySenderApp deploys TollSenderApp. It is a bit outdated and only accepts feeAmount as zero.
// execution may revert if the feeconfig is set to an amount more than zero in onramp.
func (sourceCCIP *SourceCCIPModule) DeploySenderApp(destCCIP DestCCIPModule) {
	var err error
	sourceCCIP.TollSender, err = sourceCCIP.Common.Deployer.DeployTollSenderDapp(
		sourceCCIP.OnRampRouter.EthAddress,
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
	sourceCCIP.Common.DeployContracts(contractDeployer)

	// deploy on ramp router
	sourceCCIP.OnRampRouter, err = contractDeployer.DeployOnRampRouter()
	Expect(err).ShouldNot(HaveOccurred(), "Deploying onramp router shouldn't fail")
	// wait for all contract deployments before moving on to on-ramp deployment
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for deployments")

	var tokens, pools []common.Address
	for _, token := range sourceCCIP.Common.BridgeTokens {
		tokens = append(tokens, common.HexToAddress(token.Address()))
	}
	for _, pool := range sourceCCIP.Common.NativeTokenPools {
		pools = append(pools, pool.EthAddress)
	}
	// onRamp
	sourceCCIP.OnRamp, err = contractDeployer.DeployOnRamp(
		sourceCCIP.Common.ChainClient.GetChainID(), sourceCCIP.DestinationChainId,
		tokens, pools, []common.Address{}, sourceCCIP.Common.AFN.EthAddress, sourceCCIP.OnRampRouter.EthAddress, sourceCCIP.Common.RateLimiterConfig)
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for OnRamp deployment")

	// Set token prices on the onRamp
	err = sourceCCIP.OnRamp.SetTokenPrices(tokens, sourceCCIP.Common.BridgeTokenPrices)
	Expect(err).ShouldNot(HaveOccurred(), "Setting prices shouldn't fail")
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for setting prices")

	// update onRampRouter with OnRamp address
	err = sourceCCIP.OnRampRouter.SetOnRamp(sourceCCIP.DestinationChainId, sourceCCIP.OnRamp.EthAddress)
	Expect(err).ShouldNot(HaveOccurred(), "Error setting onramp on the router")
	// update native pool with onRamp address
	for _, pool := range sourceCCIP.Common.NativeTokenPools {
		err = pool.SetOnRamp(sourceCCIP.OnRamp.EthAddress)
		Expect(err).ShouldNot(HaveOccurred(), "Error setting onramp on the token pool %s", pool.Address())
	}
	// The Fee token to be used in on+off ramp
	sourceCCIP.FeeToken = big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1e18))

	// set a part of sourceCCIP.FeeToken as onRamp fee rest would be used as offramp
	err = sourceCCIP.OnRamp.SetFeeConfig(tokens, []*big.Int{big.NewInt(1)})
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for setting OnRamp Fee config")
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for events")
}

func DefaultSourceCCIPModule(chainClient blockchain.EVMClient, destChain *big.Int, transferamount *big.Int) *SourceCCIPModule {
	return &SourceCCIPModule{
		Common:             DefaultCCIPModule(chainClient, transferamount),
		DestinationChainId: destChain,
	}
}

type DestCCIPModule struct {
	Common          *CCIPCommon
	SourceChainId   *big.Int
	BlobVerifier    *ccip.BlobVerifier
	MessageReceiver *ccip.MessageReceiver
	OffRamp         *ccip.OffRamp
	OffRampRouter   *ccip.OffRampRouter
	ReceiverDapp    *ccip.ReceiverDapp
}

// DeployContracts deploys all CCIP contracts specific to the destination chain
func (destCCIP *DestCCIPModule) DeployContracts(sourceCCIP SourceCCIPModule) {
	var err error
	destCCIP.Common.Deployer, err = ccip.NewCCIPContractsDeployer(destCCIP.Common.ChainClient)
	Expect(err).ShouldNot(HaveOccurred(), "contract deployer should be created successfully")
	contractDeployer := destCCIP.Common.Deployer
	destCCIP.Common.DeployContracts(contractDeployer)

	// blobVerifier responsible for validating the transfer message
	destCCIP.BlobVerifier, err = contractDeployer.DeployBlobVerifier(
		destCCIP.SourceChainId,
		destCCIP.Common.ChainClient.GetChainID(),
		destCCIP.Common.AFN.EthAddress,
		blob_verifier.BlobVerifierInterfaceBlobVerifierConfig{
			OnRamps:          []common.Address{sourceCCIP.OnRamp.EthAddress},
			MinSeqNrByOnRamp: []uint64{1},
		})
	Expect(err).ShouldNot(HaveOccurred(), "Deploying BlobVerifier shouldn't fail")
	err = destCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for setting destination contracts")

	var sourceTokens, destTokens, pools []common.Address
	for _, token := range sourceCCIP.Common.BridgeTokens {
		sourceTokens = append(sourceTokens, common.HexToAddress(token.Address()))
	}
	for _, token := range destCCIP.Common.BridgeTokens {
		destTokens = append(destTokens, common.HexToAddress(token.Address()))
	}
	for _, pool := range destCCIP.Common.NativeTokenPools {
		pools = append(pools, pool.EthAddress)
		destCCIP.Common.LinkToken.Transfer(pool.Address(), ccip.HundredCoins)
	}
	// offRamp
	destCCIP.OffRamp, err = contractDeployer.DeployOffRamp(destCCIP.SourceChainId, destCCIP.Common.ChainClient.GetChainID(),
		destCCIP.BlobVerifier.EthAddress, sourceCCIP.OnRamp.EthAddress, destCCIP.Common.AFN.EthAddress,
		sourceTokens, pools, destCCIP.Common.RateLimiterConfig)
	Expect(err).ShouldNot(HaveOccurred(), "Deploying OffRamp shouldn't fail")
	err = destCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for deploying offramp")
	// Set token prices on the offRamp
	err = destCCIP.OffRamp.SetTokenPrices(destTokens, destCCIP.Common.BridgeTokenPrices)
	Expect(err).ShouldNot(HaveOccurred(), "Setting prices shouldn't fail")
	err = destCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for setting prices")
	// OffRampRouter
	destCCIP.OffRampRouter, err = contractDeployer.DeployOffRampRouter([]common.Address{destCCIP.OffRamp.EthAddress})
	Expect(err).ShouldNot(HaveOccurred(), "Deploying OffRampRouter shouldn't fail")
	err = destCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for deploying OffRampRouter")
	// ReceiverDapp
	destCCIP.ReceiverDapp, err = contractDeployer.DeployReceiverDapp(destCCIP.OffRampRouter.EthAddress)
	Expect(err).ShouldNot(HaveOccurred(), "ReceiverDapp contract should be deployed successfully")
	err = destCCIP.OffRamp.SetRouter(destCCIP.OffRampRouter.EthAddress)
	Expect(err).ShouldNot(HaveOccurred(), "Error setting router on the offramp")
	// update pools with offRamp Id
	for _, pool := range destCCIP.Common.NativeTokenPools {
		err = pool.SetOffRamp(destCCIP.OffRamp.EthAddress)
		Expect(err).ShouldNot(HaveOccurred(), "Error setting offramp on the token pool %s", pool.Address())
	}
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Error waiting for events on destination contract deployments")
}

func DefaultDestinationCCIPModule(chainClient blockchain.EVMClient, sourceChain *big.Int, transferamount *big.Int) *DestCCIPModule {
	return &DestCCIPModule{
		Common:        DefaultCCIPModule(chainClient, transferamount),
		SourceChainId: sourceChain,
	}
}

// InitiateTokenTransfer initiates transfer of token with an amount as defined in sourceCCIP module. It waits for
// CCIPSendRequestedEvent, ReportRelayedEvent and ExecutionStateChanged Event to ensure that transfer has taken place
// and verifies senders and receiver's balance pre- and post-transfer
func InitiateTokenTransfer(sourceCCIP SourceCCIPModule, destCCIP DestCCIPModule) {
	// take a note of balance in both accounts before initiating the transfer
	balanceBeforeSrc, err := sourceCCIP.Common.LinkToken.BalanceOf(context.Background(),
		sourceCCIP.Common.ChainClient.GetDefaultWallet().Address())
	Expect(err).ShouldNot(HaveOccurred(), "fetching balance should work")
	balanceBeforeDest, err := destCCIP.Common.LinkToken.BalanceOf(context.Background(),
		destCCIP.Common.ChainClient.GetDefaultWallet().Address())
	Expect(err).ShouldNot(HaveOccurred(), "fetching balance should work")

	// approve the onramp router so that it can initiate transferring the token
	err = sourceCCIP.Common.LinkToken.Approve(sourceCCIP.OnRampRouter.Address(),
		bigmath.Add(sourceCCIP.FeeToken, sourceCCIP.Common.TransferAmount))
	Expect(err).ShouldNot(HaveOccurred(), "Could not approve permissions for the onRamp router "+
		"on the source link token contract")
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Failed to wait for events")

	// save the current block numbers to use in various filter log requests
	currentBlockOnSource, err := sourceCCIP.Common.ChainClient.LatestBlockNumber(context.Background())
	Expect(err).ShouldNot(HaveOccurred(), "Getting current block should be successful in source chain")
	currentBlockOnDest, err := destCCIP.Common.ChainClient.LatestBlockNumber(context.Background())
	Expect(err).ShouldNot(HaveOccurred(), "Getting current block should be successful in dest chain")

	// save the next seq number to compare with NextSeqNumber generated after transfer
	seqNumberBefore, err := destCCIP.BlobVerifier.GetNextSeqNumber(sourceCCIP.OnRamp.EthAddress)
	Expect(err).ShouldNot(HaveOccurred(), "Getting expected seq number should be successful")

	// form the message for transfer
	msg := evm_2_any_toll_onramp_router.CCIPEVM2AnyTollMessage{
		Receiver:       common.HexToAddress(destCCIP.Common.ChainClient.GetDefaultWallet().Address()),
		Data:           []byte("Token transfer by DON"),
		Tokens:         []common.Address{common.HexToAddress(sourceCCIP.Common.LinkToken.Address())},
		Amounts:        []*big.Int{sourceCCIP.Common.TransferAmount},
		FeeToken:       common.HexToAddress(sourceCCIP.Common.LinkToken.Address()),
		FeeTokenAmount: sourceCCIP.FeeToken,
		GasLimit:       big.NewInt(100_000),
	}
	log.Info().Interface("msg details", msg).Msg("ccip message to be sent")

	// initiate the transfer
	sendTx, err := sourceCCIP.OnRampRouter.CCIPSend(destCCIP.Common.ChainClient.GetChainID(), msg)
	log.Info().Str("send token transaction", sendTx.Hash().String()).Msg("Sending token")

	// Open question - If the fee config is set with more than 0 feeAmount sending tokens with tollsender is getting reverted
	/*	sendTx, err := sourceCCIP.TollSender.SendTokens(
			common.HexToAddress(destCCIP.Common.ChainClient.GetDefaultWallet().Address()),
			[]common.Address{common.HexToAddress(sourceCCIP.Common.LinkToken.Address())},
			[]*big.Int{sourceCCIP.Common.TransferAmount},
		)
	*/
	if err != nil {
		err = blockchain.LogRevertReason(err, evm_2_any_toll_onramp_router.EVM2AnyTollOnRampRouterABI)
		Expect(err).ShouldNot(HaveOccurred(), "error in logging revert reason")
	}

	Expect(err).ShouldNot(HaveOccurred(), "send token should be initiated successfully")
	err = sourceCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Failed to wait for events")

	// Verify if
	// - CCIPSendRequested Event log generated,
	// - NextSeqNumber from blobVerifier got increased
	Eventually(func(g Gomega) {
		iterator, err := sourceCCIP.OnRamp.FilterCCIPSendRequested(currentBlockOnSource)
		g.Expect(err).NotTo(HaveOccurred(), "Error filtering CCIPSendRequested event")
		g.Expect(iterator.Next()).To(BeTrue(), "No CCIPSendRequested event found")
		g.Expect(iterator.Event.Raw.TxHash.Hex()).Should(Equal(sendTx.Hash().Hex()), "CCIPSendRequested event not found")
		seqNumberAfter, err := destCCIP.BlobVerifier.GetNextSeqNumber(sourceCCIP.OnRamp.EthAddress)
		g.Expect(err).ShouldNot(HaveOccurred(), "Getting expected seq number should be successful")
		g.Expect(seqNumberAfter).Should(BeNumerically(">", seqNumberBefore), "Next Sequence number is not increased")
	}, "2m", "1s").Should(Succeed(), "Error Relaying report") // 2m timeout is a bit longer. It flakes with 1m timeout

	balanceAfterSrc, err := sourceCCIP.Common.LinkToken.BalanceOf(context.Background(),
		sourceCCIP.Common.ChainClient.GetDefaultWallet().Address())
	Expect(err).ShouldNot(HaveOccurred(), "fetching balance should work")
	log.Info().
		Str("Balance Before in sender account", balanceBeforeSrc.String()).
		Str("balance After sender", balanceAfterSrc.String()).
		Str("Balance Now in receiver", balanceBeforeDest.String()).
		Msg("balance in sender and receiver before and after CCIPSendRequested event")

	// Verify whether blobVerifier has accepted the report
	Eventually(func(g Gomega) []common.Address {
		iterator, err := destCCIP.BlobVerifier.FilterReportAccepted(currentBlockOnDest)
		g.Expect(err).NotTo(HaveOccurred(), "Error filtering ReportAccepted event")
		g.Expect(iterator.Next()).To(BeTrue(), "No ReportAccepted event found")
		return iterator.Event.Report.OnRamps
	}, "1m", "1s").Should(Equal([]common.Address{sourceCCIP.OnRamp.EthAddress}))

	// Verify whether the execution state is changed and the transfer is successful
	Eventually(func(g Gomega) ccipPlugin.MessageExecutionState {
		iterator, err := destCCIP.OffRamp.FilterExecutionStateChanged([]uint64{seqNumberBefore})
		g.Expect(err).NotTo(HaveOccurred(), "Error filtering ExecutionStateChanged event")
		g.Expect(iterator.Next()).To(BeTrue(), "No ExecutionStateChanged event found")
		return ccipPlugin.MessageExecutionState(iterator.Event.State)
	}, "1m", "1s").Should(Equal(ccipPlugin.Success))

	balanceAfterDest, err := destCCIP.Common.LinkToken.BalanceOf(context.Background(),
		destCCIP.Common.ChainClient.GetDefaultWallet().Address())

	log.Info().
		Str("Balance Before receiver", balanceBeforeDest.String()).
		Str("Balance Now in receiver", balanceAfterDest.String()).
		Msg("Change in balance in sender and receiver")

	// verify that sender balance should be deducted by transferAmount + feeTokenAmount
	Expect(bigmath.Equal(
		bigmath.Add(sourceCCIP.Common.TransferAmount, sourceCCIP.FeeToken),
		bigmath.Sub(balanceBeforeSrc, balanceAfterSrc))).
		Should(BeTrue(), "BalanceAfter-BalanceBefore should match the transfer amount + FeeToken for Address %s",
			sourceCCIP.Common.ChainClient.GetDefaultWallet().Address())

	// verify that receiver balance is increased by transferAmount
	// TODO use generic method for balance assertion
	unusedFee := big.NewInt(9.0491e18) // for single transaction
	expectedBalance := bigmath.Add(bigmath.Add(sourceCCIP.Common.TransferAmount, unusedFee), balanceBeforeDest)
	high := big.NewInt(0).Add(expectedBalance, big.NewInt(1e18))
	low := big.NewInt(0).Sub(expectedBalance, big.NewInt(1e18))
	Expect(balanceAfterDest.Cmp(high)).
		Should(BeNumerically("==", -1),
			"Receiver balance diff between BalanceAfter-%v BalanceBefore-%v should be lesser than %v",
			balanceAfterDest, balanceBeforeDest, high)
	Expect(balanceAfterDest.Cmp(low)).
		Should(BeNumerically("==", 1),
			"Receiver balance diff between BalanceAfter-%v BalanceBefore-%v should be greater than %v",
			balanceAfterDest, balanceBeforeDest, low)
}

func PrintSourceContractDetails(sourceCCIP SourceCCIPModule, destCCIP DestCCIPModule) {
	var sourceTokens, destTokens, pools []string
	for _, token := range sourceCCIP.Common.BridgeTokens {
		sourceTokens = append(sourceTokens, token.Address())
	}
	for _, pool := range sourceCCIP.Common.NativeTokenPools {
		pools = append(pools, pool.Address())
	}

	log.Info().
		Str("Link Token", sourceCCIP.Common.LinkToken.Address()).
		Strs("Bridge Token", sourceTokens).
		Strs("Pool", pools).
		Str("AFN", sourceCCIP.Common.AFN.Address()).
		Str("OnRampRouter", sourceCCIP.OnRampRouter.Address()).
		Str("OnRamp", sourceCCIP.OnRamp.Address()).Msg("Source Contracts")
	for _, token := range destCCIP.Common.BridgeTokens {
		destTokens = append(destTokens, token.Address())
	}
	for _, pool := range destCCIP.Common.NativeTokenPools {
		pools = append(pools, pool.Address())
	}

	log.Info().
		Str("Link Token", destCCIP.Common.LinkToken.Address()).
		Strs("Bridge Token", destTokens).
		Strs("Pool", pools).
		Str("AFN", destCCIP.Common.AFN.Address()).
		Str("BlobVerifier", destCCIP.BlobVerifier.Address()).
		Str("OffRampRouter", destCCIP.OffRampRouter.Address()).
		Str("OffRamp", destCCIP.OffRamp.Address()).
		Str("ReceiverDapp", destCCIP.ReceiverDapp.Address()).
		Msg("Destination Contracts")
}

func SetOCRConfigs(chainlinkNodes []*client.CLNodesWithKeys, destCCIP DestCCIPModule) {
	signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig, err :=
		ccip.NewOffChainAggregatorV2Config(chainlinkNodes)
	Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail while getting the config values for ocr2 type contract")
	err = destCCIP.BlobVerifier.SetOCRConfig(signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
	Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail while setting blobverifier config")
	err = destCCIP.OffRamp.SetConfig(signers, transmitters, f, onchainConfig, offchainConfigVersion, offchainConfig)
	Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail while setting OffRamp config")
	err = destCCIP.Common.ChainClient.WaitForEvents()
	Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail while waiting for events on setting ocr2 config")
}

// CreateOCRJobsForCCIP bootstraps the first node and to the other nodes sends ocr jobs that
// sets up ccip-relay and ccip-execution plugin
func CreateOCRJobsForCCIP(
	chainlinkNodes []*client.CLNodesWithKeys,
	onramp, blobVerifier, offRamp string,
	sourceChainClient, destChainClient blockchain.EVMClient,
	destLinkTokenAddr string,
	mockserver *ctfClient.MockserverClient,
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
					"onRampIDs":        []string{fmt.Sprintf("\"%s\"", onramp)},
					"pollPeriod":       `"1s"`,
					"destStartBlock":   currentBlockOnDest,
					"sourceStartBlock": currentBlockOnSource,
				},
				RelayConfig: map[string]interface{}{
					"chainID": fmt.Sprintf("\"%s\"", destChainID.String()),
				},
			},
		}
		_, err = chainlinkNodes[nodeIndex].Node.MustCreateJob(ocr2SpecRelay)
		Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail creating CCIP-Relay OCR Task job on OCR node %d", nodeIndex+1)
		tokenFeeConversionRateURL := fmt.Sprintf("%s/%s", mockserver.Config.ClusterURL,
			nodeContractPair(chainlinkNodes[nodeIndex].KeysBundle.EthAddress, destLinkTokenAddr))
		ocr2SpecExec := &client.OCR2TaskJobSpec{
			JobType: "offchainreporting2",
			Name:    fmt.Sprintf("ccip-exec-%s-%s", sourceChainName, destChainName),
			OCR2OracleSpec: job.OCR2OracleSpec{
				Relay:                             relay.EVM,
				PluginType:                        job.CCIPExecution,
				ContractID:                        offRamp,
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
					"onRampID":         fmt.Sprintf("\"%s\"", onramp),
					"blobVerifierID":   fmt.Sprintf("\"%s\"", blobVerifier),
					"pollPeriod":       `"1s"`,
					"destStartBlock":   currentBlockOnDest,
					"sourceStartBlock": currentBlockOnSource,
					"tokensPerFeeCoinPipeline": fmt.Sprintf(`"""
link [type=http method=GET url="%s"];
link_parse [type=jsonparse path="Data,Result"];
link->link_parse;
merge [type=merge left="{}" right="{\\\"%s\\\":$(link_parse)}"];
"""`, tokenFeeConversionRateURL, destLinkTokenAddr),
				},
				RelayConfig: map[string]interface{}{
					"chainID": fmt.Sprintf("\"%s\"", destChainID.String()),
				},
			},
		}
		_, err = chainlinkNodes[nodeIndex].Node.MustCreateJob(ocr2SpecExec)
		Expect(err).ShouldNot(HaveOccurred(), "Shouldn't fail creating CCIP-Exec OCR Task job on OCR node %d", nodeIndex+1)
	}
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
				log.Info().Str("path", path).Msg("setting exp for path")
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
