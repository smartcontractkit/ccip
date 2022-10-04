package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/any_2_evm_subscription_offramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/blob_verifier"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_any_subscription_onramp_router"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_subscription_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/native_token_pool"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/ping_pong_demo"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/receiver_dapp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/subscription_sender_dapp"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
)

func printCCIPState(source *EvmChainConfig, destination *EvmChainConfig) {
	printPoolBalances(source)
	printPoolBalances(destination)

	printDappSanityCheck(source)
	printDappSanityCheck(destination)
	printRampSanityCheck(source)
	printRampSanityCheck(destination)

	printSourceSubscriptionBalances(source)
	printDestinationSubscriptionBalances(destination)

	printPaused(source)
	printPaused(destination)

	printRateLimitingStatus(source)
	printRateLimitingStatus(destination)

	printTxStatuses(source, destination)

	printContractConfig(source, destination)
}

type CCIPTXStatus struct {
	message     *evm_2_evm_subscription_onramp.EVM2EVMSubscriptionOnRampCCIPSendRequested
	relayReport *blob_verifier.BlobVerifierReportAccepted
	execStatus  *any_2_evm_subscription_offramp.EVM2EVMSubscriptionOffRampExecutionStateChanged
}

type ExecutionStatus uint8

const (
	Untouched  ExecutionStatus = 0
	InProgress ExecutionStatus = 1
	Success    ExecutionStatus = 2
	Failed     ExecutionStatus = 3
)

func (e ExecutionStatus) String() string {
	switch e {
	case Untouched:
		return "Untouched"
	case InProgress:
		return "InProgress"
	case Success:
		return "Success"
	case Failed:
		return "Failed"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

func printTxStatuses(source *EvmChainConfig, destination *EvmChainConfig) {
	onRamp, err := evm_2_evm_subscription_onramp.NewEVM2EVMSubscriptionOnRamp(source.OnRamp, source.Client)
	helpers.PanicErr(err)

	sendRequested, err := onRamp.FilterCCIPSendRequested(&bind.FilterOpts{
		Start: source.DeploySettings.DeployedAt,
	})
	helpers.PanicErr(err)

	txs := make(map[uint64]*CCIPTXStatus)
	maxSeqNum := uint64(0)
	var seqNums []uint64

	for sendRequested.Next() {
		txs[sendRequested.Event.Message.SequenceNumber] = &CCIPTXStatus{
			message: sendRequested.Event,
		}
		if sendRequested.Event.Message.SequenceNumber > maxSeqNum {
			maxSeqNum = sendRequested.Event.Message.SequenceNumber
		}
		seqNums = append(seqNums, sendRequested.Event.Message.SequenceNumber)
	}

	blobVerifier, err := blob_verifier.NewBlobVerifier(destination.BlobVerifier, destination.Client)
	helpers.PanicErr(err)

	reports, err := blobVerifier.FilterReportAccepted(&bind.FilterOpts{
		Start: destination.DeploySettings.DeployedAt,
	})
	helpers.PanicErr(err)

	for reports.Next() {
		for i, interval := range reports.Event.Report.Intervals {
			if reports.Event.Report.OnRamps[i] != source.OnRamp {
				continue
			}
			for j := interval.Min; j <= interval.Max; j++ {
				txs[j].relayReport = reports.Event
			}
		}
	}

	offRamp, err := any_2_evm_subscription_offramp.NewEVM2EVMSubscriptionOffRamp(destination.OffRamp, destination.Client)
	helpers.PanicErr(err)

	stateChanges, err := offRamp.FilterExecutionStateChanged(&bind.FilterOpts{
		Start: destination.DeploySettings.DeployedAt,
	}, seqNums)
	helpers.PanicErr(err)

	for stateChanges.Next() {
		txs[stateChanges.Event.SequenceNumber].execStatus = stateChanges.Event
	}

	var sb strings.Builder
	sb.WriteString("\n")
	tableHeaders := []string{"SequenceNumber", "Relayed in block", "Execution status", "Executed in block", "Nonce"}
	headerLengths := []int{18, 18, 20, 18, 18}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))

	for i := uint64(1); i <= maxSeqNum; i++ {
		tx := txs[i]
		relayedAt := "-"
		if tx.relayReport != nil {
			relayedAt = strconv.Itoa(int(tx.relayReport.Raw.BlockNumber))
		}

		if tx.message == nil {
			sb.WriteString(fmt.Sprintf("| %18s | %18s | %20v | %18d | %18s | \n", "MISSING", relayedAt, ExecutionStatus(tx.execStatus.State), tx.execStatus.Raw.BlockNumber, "-"))
		} else if tx.execStatus != nil {
			sb.WriteString(fmt.Sprintf("| %18d | %18s | %20v | %18d | %18d | %s \n",
				tx.message.Message.SequenceNumber,
				relayedAt,
				ExecutionStatus(tx.execStatus.State),
				tx.execStatus.Raw.BlockNumber,
				tx.message.Message.Nonce,
				helpers.ExplorerLink(destination.ChainId.Int64(), tx.execStatus.Raw.TxHash)))
		} else {
			sb.WriteString(fmt.Sprintf("| %18d | %18s | %20v | %18s | %18d | %s \n",
				tx.message.Message.SequenceNumber,
				relayedAt,
				"-",
				"-",
				tx.message.Message.Nonce,
				""))
		}
	}
	sb.WriteString(generateSeparator(headerLengths))

	destination.Logger.Info(sb.String())
}

func printDappSanityCheck(source *EvmChainConfig) {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("Dapp sanity checks for %s\n", helpers.ChainName(source.ChainId.Int64())))

	tableHeaders := []string{"Dapp", "Router Set"}
	headerLengths := []int{30, 15}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))

	onRamp, err := evm_2_evm_subscription_onramp.NewEVM2EVMSubscriptionOnRamp(source.OnRamp, source.Client)
	helpers.PanicErr(err)
	router, err := onRamp.GetRouter(&bind.CallOpts{})
	helpers.PanicErr(err)
	sb.WriteString(fmt.Sprintf("| %-30s | %15t |\n", "OnRamp", router == source.OnRampRouter))

	offRamp, err := any_2_evm_subscription_offramp.NewEVM2EVMSubscriptionOffRamp(source.OffRamp, source.Client)
	helpers.PanicErr(err)
	router, err = offRamp.GetRouter(&bind.CallOpts{})
	helpers.PanicErr(err)

	sb.WriteString(fmt.Sprintf("| %-30s | %15t |\n", "OffRamp", router == source.OffRampRouter))

	sb.WriteString(generateSeparator(headerLengths))

	senderDapp, err := subscription_sender_dapp.NewSubscriptionSenderDapp(source.TokenSender, source.Client)
	helpers.PanicErr(err)
	router, err = senderDapp.IOnRampRouter(&bind.CallOpts{})
	helpers.PanicErr(err)
	sb.WriteString(fmt.Sprintf("| %-30s | %15t |\n", "Sender dapp", router == source.OnRampRouter))

	receiverDap, err := receiver_dapp.NewReceiverDapp(source.ReceiverDapp, source.Client)
	helpers.PanicErr(err)
	router, err = receiverDap.SRouter(&bind.CallOpts{})
	helpers.PanicErr(err)
	sb.WriteString(fmt.Sprintf("| %-30s | %15t |\n", "Receiver dapp", router == source.OffRampRouter))

	pingDapp, err := ping_pong_demo.NewPingPongDemo(source.PingPongDapp, source.Client)
	helpers.PanicErr(err)
	receiver, sender, err := pingDapp.GetRouters(&bind.CallOpts{})
	helpers.PanicErr(err)
	sb.WriteString(fmt.Sprintf("| %-30s | %15t |\n", "Ping dapp receiver", receiver == source.OffRampRouter))
	sb.WriteString(fmt.Sprintf("| %-30s | %15t |\n", "Ping dapp sender", sender == source.OnRampRouter))

	sb.WriteString(generateSeparator(headerLengths))

	source.Logger.Info(sb.String())
}

func printRampSanityCheck(chain *EvmChainConfig) {
	afn, err := afn_contract.NewAFNContract(chain.Afn, chain.Client)
	helpers.PanicErr(err)

	badSignal, err := afn.BadSignalReceived(&bind.CallOpts{})
	helpers.PanicErr(err)

	chain.Logger.Infof("AFN healthy: %t", !badSignal)
}

func printSourceSubscriptionBalances(source *EvmChainConfig) {
	onRampRouter, err := evm_2_any_subscription_onramp_router.NewEVM2AnySubscriptionOnRampRouter(source.OnRampRouter, source.Client)
	helpers.PanicErr(err)

	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("Source subscription balances for %s\n", helpers.ChainName(source.ChainId.Int64())))

	tableHeaders := []string{"Sender", "Address", "Balance", "#Txs funded"}
	headerLengths := []int{20, 42, 22, 22}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))

	fee, err := onRampRouter.GetFee(&bind.CallOpts{})
	helpers.PanicErr(err)

	possibleTxs := "∞"
	balance, err := onRampRouter.GetBalance(&bind.CallOpts{}, source.Owner.From)
	helpers.PanicErr(err)
	if fee.Int64() != 0 {
		possibleTxs = big.NewInt(0).Div(balance, fee).String()
	}
	sb.WriteString(fmt.Sprintf("| %-20s | %42s | %22d | %22s |\n", "owner", source.Owner.From.Hex(), balance, possibleTxs))

	balance, err = onRampRouter.GetBalance(&bind.CallOpts{}, source.TokenSender)
	helpers.PanicErr(err)
	if fee.Int64() != 0 {
		possibleTxs = big.NewInt(0).Div(balance, fee).String()
	}
	sb.WriteString(fmt.Sprintf("| %-20s | %42s | %22d | %22s |\n", "sender dapp", source.TokenSender.Hex(), balance, possibleTxs))

	balance, err = onRampRouter.GetBalance(&bind.CallOpts{}, source.GovernanceDapp)
	helpers.PanicErr(err)
	if fee.Int64() != 0 {
		possibleTxs = big.NewInt(0).Div(balance, fee).String()
	}
	sb.WriteString(fmt.Sprintf("| %-20s | %42s | %22d | %22s |\n", "governance dapp", source.GovernanceDapp.Hex(), balance, possibleTxs))

	balance, err = onRampRouter.GetBalance(&bind.CallOpts{}, source.PingPongDapp)
	helpers.PanicErr(err)
	if fee.Int64() != 0 {
		possibleTxs = big.NewInt(0).Div(balance, fee).String()
	}
	sb.WriteString(fmt.Sprintf("| %-20s | %42s | %22d | %22s |\n", "ping pong dapp", source.PingPongDapp.Hex(), balance, possibleTxs))

	sb.WriteString(generateSeparator(headerLengths))

	sb.WriteString(fmt.Sprintf("| %-20s | %92d |\n", "relay fee", fee))

	sb.WriteString(generateSeparator(headerLengths))

	source.Logger.Info(sb.String())
}

func printDestinationSubscriptionBalances(destination *EvmChainConfig) {
	offRampRouter, err := any_2_evm_subscription_offramp_router.NewAny2EVMSubscriptionOffRampRouter(destination.OffRampRouter, destination.Client)
	helpers.PanicErr(err)

	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("Destination subscription balances for %s\n", helpers.ChainName(destination.ChainId.Int64())))

	tableHeaders := []string{"Receiver", "Address", "Balance", "Allowed senders"}
	headerLengths := []int{20, 42, 22, 44}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))

	subscription, err := offRampRouter.GetSubscription(&bind.CallOpts{}, destination.ReceiverDapp)
	helpers.PanicErr(err)
	sb.WriteString(fmt.Sprintf("| %-20s | %42s | %22d | %42v |\n", "receiver dapp", destination.ReceiverDapp.Hex(), subscription.Balance, subscription.Senders))

	subscription, err = offRampRouter.GetSubscription(&bind.CallOpts{}, destination.MessageReceiver)
	helpers.PanicErr(err)
	sb.WriteString(fmt.Sprintf("| %-20s | %42s | %22d | %42v |\n", "message receiver", destination.MessageReceiver.Hex(), subscription.Balance, subscription.Senders))

	subscription, err = offRampRouter.GetSubscription(&bind.CallOpts{}, destination.GovernanceDapp)
	helpers.PanicErr(err)
	sb.WriteString(fmt.Sprintf("| %-20s | %42s | %22d | %42v |\n", "governance dapp", destination.GovernanceDapp.Hex(), subscription.Balance, subscription.Senders))

	subscription, err = offRampRouter.GetSubscription(&bind.CallOpts{}, destination.PingPongDapp)
	helpers.PanicErr(err)
	sb.WriteString(fmt.Sprintf("| %-20s | %42s | %22d | %42v |\n", "ping pong dapp", destination.PingPongDapp.Hex(), subscription.Balance, subscription.Senders))

	sb.WriteString(generateSeparator(headerLengths))

	destination.Logger.Info(sb.String())
}

func printRateLimitingStatus(chain *EvmChainConfig) {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("Rate limits for %s\n", helpers.ChainName(chain.ChainId.Int64())))

	tableHeaders := []string{"Contract", "Tokens"}
	headerLengths := []int{25, 42}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))

	onRamp, err := evm_2_evm_subscription_onramp.NewEVM2EVMSubscriptionOnRamp(chain.OnRamp, chain.Client)
	helpers.PanicErr(err)
	bucketState, err := onRamp.CalculateCurrentTokenBucketState(&bind.CallOpts{})
	helpers.PanicErr(err)

	sb.WriteString(fmt.Sprintf("| %-25s | %42d |\n", "onramp", bucketState.Tokens))

	offRamp, err := any_2_evm_subscription_offramp.NewEVM2EVMSubscriptionOffRamp(chain.OffRamp, chain.Client)
	helpers.PanicErr(err)
	offRampBucketState, err := offRamp.CalculateCurrentTokenBucketState(&bind.CallOpts{})
	helpers.PanicErr(err)

	sb.WriteString(fmt.Sprintf("| %-25s | %42d |\n", "offramp", offRampBucketState.Tokens))

	sb.WriteString(generateSeparator(headerLengths))
	chain.Logger.Info(sb.String())
}

func printPaused(chain *EvmChainConfig) {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("Paused addresses for %s\n", helpers.ChainName(chain.ChainId.Int64())))

	tableHeaders := []string{"Contract", "Address", "Running"}
	headerLengths := []int{25, 42, 20}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))

	for _, pool := range chain.TokenPools {
		tokenPool, err := native_token_pool.NewNativeTokenPool(pool, chain.Client)
		helpers.PanicErr(err)
		paused, err := tokenPool.Paused(&bind.CallOpts{})
		helpers.PanicErr(err)

		sb.WriteString(fmt.Sprintf("| %-25s | %42s | %20t |\n", "token pool", pool.Hex(), !paused))
	}

	onRamp, err := evm_2_evm_subscription_onramp.NewEVM2EVMSubscriptionOnRamp(chain.OnRamp, chain.Client)
	helpers.PanicErr(err)
	paused, err := onRamp.Paused(&bind.CallOpts{})
	helpers.PanicErr(err)

	sb.WriteString(fmt.Sprintf("| %-25s | %42s | %20t |\n", "onramp", onRamp.Address(), !paused))

	offRamp, err := any_2_evm_subscription_offramp.NewEVM2EVMSubscriptionOffRamp(chain.OffRamp, chain.Client)
	helpers.PanicErr(err)
	paused, err = offRamp.Paused(&bind.CallOpts{})
	helpers.PanicErr(err)

	sb.WriteString(fmt.Sprintf("| %-25s | %42s | %20t |\n", "offramp", offRamp.Address(), !paused))

	blobVerifier, err := blob_verifier.NewBlobVerifier(chain.BlobVerifier, chain.Client)
	helpers.PanicErr(err)
	paused, err = blobVerifier.Paused(&bind.CallOpts{})
	helpers.PanicErr(err)

	sb.WriteString(fmt.Sprintf("| %-25s | %42s | %20t |\n", "blobverifier", blobVerifier.Address(), !paused))

	sb.WriteString(generateSeparator(headerLengths))
	chain.Logger.Info(sb.String())
}

func printPoolBalances(chain *EvmChainConfig) {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("Pool balances for %s\n", helpers.ChainName(chain.ChainId.Int64())))

	tableHeaders := []string{"Token", "Pool", "Balance", "Onramp", "OffRamp"}
	headerLengths := []int{42, 42, 20, 10, 10}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))
	for _, pool := range chain.TokenPools {
		tokenPool, err := native_token_pool.NewNativeTokenPool(pool, chain.Client)
		helpers.PanicErr(err)

		tokenAddress, err := tokenPool.GetToken(&bind.CallOpts{})
		helpers.PanicErr(err)

		token, err := link_token_interface.NewLinkToken(tokenAddress, chain.Client)
		helpers.PanicErr(err)

		balance, err := token.BalanceOf(&bind.CallOpts{}, pool)
		helpers.PanicErr(err)

		isAllowedOnRamp, err := tokenPool.IsOnRamp(&bind.CallOpts{}, chain.OnRamp)
		helpers.PanicErr(err)

		isAllowedOffRamp, err := tokenPool.IsOffRamp(&bind.CallOpts{}, chain.OffRamp)
		helpers.PanicErr(err)

		sb.WriteString(fmt.Sprintf("| %s | %s | %20d | %10t | %10t |\n", tokenAddress.Hex(), pool.Hex(), balance, isAllowedOnRamp, isAllowedOffRamp))
	}

	sb.WriteString(generateSeparator(headerLengths))

	chain.Logger.Info(sb.String())
}

func generateHeader(headers []string, headerLengths []int) string {
	var sb strings.Builder

	sb.WriteString(generateSeparator(headerLengths))
	sb.WriteString("|")
	for i, header := range headers {
		sb.WriteString(fmt.Sprintf(" %-"+strconv.Itoa(headerLengths[i])+"s |", header))
	}
	sb.WriteString("\n")
	sb.WriteString(generateSeparator(headerLengths))

	return sb.String()
}

func generateSeparator(headerLengths []int) string {
	length := 1

	for _, headerLength := range headerLengths {
		length += headerLength + 3
	}
	return strings.Repeat("-", length) + "\n"
}

func printContractConfig(source *EvmChainConfig, destination *EvmChainConfig) {
	source.Logger.Infof(`
Source chain config

LinkToken:      common.HexToAddress("%s"),
BridgeTokens:   %s,
TokenPools:     %s,
OnRamp:         common.HexToAddress("%s"),
OnRampRouter:   common.HexToAddress("%s"),
TokenSender:    common.HexToAddress("%s"),
Afn:            common.HexToAddress("%s"),
GovernanceDapp: common.HexToAddress("%s"),
PingPongDapp:   common.HexToAddress("%s"),
	
`,
		source.LinkToken,
		source.BridgeTokens,
		source.TokenPools,
		source.OnRamp,
		source.OnRampRouter,
		source.TokenSender,
		source.Afn,
		source.GovernanceDapp,
		source.PingPongDapp)

	destination.Logger.Infof(`
Destination chain config

LinkToken:       common.HexToAddress("%s"),
BridgeTokens:    %s,
TokenPools:      %s,
OffRamp:         common.HexToAddress("%s"),
OffRampRouter:   common.HexToAddress("%s"),
BlobVerifier:    common.HexToAddress("%s"),	
MessageReceiver: common.HexToAddress("%s"),
ReceiverDapp:    common.HexToAddress("%s"),
Afn:             common.HexToAddress("%s"),
GovernanceDapp:  common.HexToAddress("%s"),
PingPongDapp:    common.HexToAddress("%s"),
`,
		destination.LinkToken,
		destination.BridgeTokens,
		destination.TokenPools,
		destination.OffRamp,
		destination.OffRampRouter,
		destination.BlobVerifier,
		destination.MessageReceiver,
		destination.ReceiverDapp,
		destination.Afn,
		destination.GovernanceDapp,
		destination.PingPongDapp)

	PrintJobSpecs(source.OnRamp, destination.BlobVerifier, destination.OffRamp, source.ChainId, destination.ChainId, destination.LinkToken, source.DeploySettings.DeployedAt, destination.DeploySettings.DeployedAt)
}

// PrintJobSpecs prints the job spec for each node and CCIP spec type, as well as a bootstrap spec.
func PrintJobSpecs(onramp, blobVerifier, offRamp common.Address, sourceChainID, destChainID *big.Int, destLinkToken common.Address, sourceReplayFrom, destReplayFrom uint64) {
	jobs := fmt.Sprintf(bootstrapTemplate+"\n", helpers.ChainName(destChainID.Int64()), blobVerifier, destChainID)
	don := NewDON(Staging, nil)

	for i, oracle := range don.config.Nodes {
		jobs += fmt.Sprintf("// [Node %d]\n", i)
		jobs += fmt.Sprintf(relayTemplate+"\n",
			helpers.ChainName(sourceChainID.Int64())+"-"+helpers.ChainName(destChainID.Int64()),
			blobVerifier,
			GetOCRkeysForChainType(oracle.OCRKeys, "evm").ID,
			oracle.EthKeys[destChainID.String()],
			don.GetBootstrapPeerID(),
			sourceChainID,
			destChainID,
			onramp,
			pollPeriod,
			sourceReplayFrom,
			destReplayFrom,
			destChainID,
		)
		jobs += fmt.Sprintf(executionTemplate+"\n",
			helpers.ChainName(sourceChainID.Int64())+"-"+helpers.ChainName(destChainID.Int64()),
			offRamp,
			GetOCRkeysForChainType(oracle.OCRKeys, "evm").ID,
			oracle.EthKeys[destChainID.String()],
			don.GetBootstrapPeerID(),
			sourceChainID,
			destChainID,
			onramp,
			blobVerifier,
			sourceReplayFrom,
			destReplayFrom,
			destLinkToken,
			destChainID,
		)
	}
	fmt.Println(jobs)
}

const bootstrapTemplate = `
// Bootstrap Node
# BootstrapSpec
type                               = "bootstrap"
name                               = "bootstrap-%s"
relay                              = "evm"
schemaVersion                      = 1
contractID                         = "%s"
contractConfigConfirmations        = 1
contractConfigTrackerPollInterval  = "60s"
[relayConfig]
chainID                            = %s
`

const relayTemplate = `
# CCIPRelaySpec
type               = "offchainreporting2"
name               = "ccip-relay-%s"
pluginType         = "ccip-relay"
relay              = "evm"
schemaVersion      = 1
contractID         = "%s"
ocrKeyBundleID     = "%s"
transmitterID      = "%s"
p2pv2Bootstrappers  = ["%s"]

[pluginConfig]
sourceChainID      = %d
destChainID        = %d
onRampIDs          = ["%s"]
pollPeriod         = "%s"
SourceStartBlock   = %d
DestStartBlock     = %d

[relayConfig]
chainID            = %d
`

const executionTemplate = `
# CCIPExecutionSpec
type              = "offchainreporting2"
name              = "ccip-exec-%s"
pluginType        = "ccip-execution"
relay             = "evm"
schemaVersion     = 1
contractID        = "%s"
ocrKeyBundleID    = "%s"
transmitterID     = "%s"
p2pv2Bootstrappers = ["%s"]

[pluginConfig]
sourceChainID     = %d
destChainID       = %d
onRampID          = "%s"
blobVerifierID    = "%s"
SourceStartBlock  = %d
DestStartBlock    = %d
tokensPerFeeCoinPipeline = """merge [type=merge left="{}" right="{\\\"%s\\\":\\\"1000000000000000000\\\"}"];"""

[relayConfig]
chainID           = %d
`
