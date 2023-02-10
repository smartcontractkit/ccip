package printing

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/exp/slices"

	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/afn_contract"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/link_token_interface"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/ping_pong_demo"
	"github.com/smartcontractkit/chainlink/core/gethwrappers/generated/router"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/dione"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/rhea"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
)

func PrintCCIPState(source *rhea.EvmDeploymentConfig, destination *rhea.EvmDeploymentConfig) {
	printPoolBalances(source)
	printPoolBalances(destination)

	printSupportedTokensCheck(source, destination)

	printDappSanityCheck(source)
	printDappSanityCheck(destination)

	printRampSanityCheck(source, destination.LaneConfig.OnRamp)
	printRampSanityCheck(destination, source.LaneConfig.OnRamp)

	printPaused(source)
	printPaused(destination)

	printRateLimitingStatus(source)
	printRateLimitingStatus(destination)
}

type CCIPTXStatus struct {
	message      *evm_2_evm_onramp.EVM2EVMOnRampCCIPSendRequested
	commitReport *commit_store.CommitStoreReportAccepted
	execStatus   *evm_2_evm_offramp.EVM2EVMOffRampExecutionStateChanged
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

func printBool(b bool) string {
	if b {
		return "✅"
	}
	return "❌"
}

func printBoolNeutral(b bool) string {
	if b {
		return "✅"
	}
	return "➖"
}

func PrintTxStatuses(source *rhea.EvmDeploymentConfig, destination *rhea.EvmDeploymentConfig) {
	onRamp, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(source.LaneConfig.OnRamp, source.Client)
	helpers.PanicErr(err)

	block, err := source.Client.BlockNumber(context.Background())
	helpers.PanicErr(err)

	sendRequested, err := onRamp.FilterCCIPSendRequested(&bind.FilterOpts{
		Start: block - 9990,
	})
	helpers.PanicErr(err)

	txs := make(map[uint64]*CCIPTXStatus)
	maxSeqNum := uint64(0)
	minSeqNum := uint64(1)
	var seqNums []uint64

	for sendRequested.Next() {
		txs[sendRequested.Event.Message.SequenceNumber] = &CCIPTXStatus{
			message: sendRequested.Event,
		}
		if sendRequested.Event.Message.SequenceNumber > maxSeqNum {
			maxSeqNum = sendRequested.Event.Message.SequenceNumber
		}
		if minSeqNum == 1 {
			minSeqNum = sendRequested.Event.Message.SequenceNumber
		}
		seqNums = append(seqNums, sendRequested.Event.Message.SequenceNumber)
	}

	commitStore, err := commit_store.NewCommitStore(destination.LaneConfig.CommitStore, destination.Client)
	helpers.PanicErr(err)

	block, err = destination.Client.BlockNumber(context.Background())
	helpers.PanicErr(err)

	reports, err := commitStore.FilterReportAccepted(&bind.FilterOpts{
		Start: block - 9990,
	})
	helpers.PanicErr(err)

	for reports.Next() {
		for j := reports.Event.Report.Interval.Min; j <= reports.Event.Report.Interval.Max; j++ {
			if _, ok := txs[j]; ok {
				txs[j].commitReport = reports.Event
			}
		}
	}

	offRamp, err := evm_2_evm_offramp.NewEVM2EVMOffRamp(destination.LaneConfig.OffRamp, destination.Client)
	helpers.PanicErr(err)

	stateChanges, err := offRamp.FilterExecutionStateChanged(
		&bind.FilterOpts{
			Start: block - 9990,
		},
		seqNums,
		[][32]byte{})
	helpers.PanicErr(err)

	for stateChanges.Next() {
		if _, ok := txs[stateChanges.Event.SequenceNumber]; !ok {
			txs[stateChanges.Event.SequenceNumber] = &CCIPTXStatus{}
			if stateChanges.Event.SequenceNumber > maxSeqNum {
				maxSeqNum = stateChanges.Event.SequenceNumber
			}
		}
		txs[stateChanges.Event.SequenceNumber].execStatus = stateChanges.Event
	}

	var sb strings.Builder
	sb.WriteString("\n")
	tableHeaders := []string{"SequenceNumber", "Committed in block", "Execution status", "Executed in block", "Nonce"}
	headerLengths := []int{18, 18, 20, 18, 18}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))

	if minSeqNum > 1 {
		sb.WriteString(fmt.Sprintf("| %18d | %18d | %41s | %18s | \n", 1, minSeqNum-1, "Probably > 10k blocks in the past", ""))
	}

	for i := minSeqNum; i <= maxSeqNum; i++ {
		tx := txs[i]
		committedAt := "-"
		if tx == nil {
			sb.WriteString(fmt.Sprintf("| %18d | %18s | %41s | %18s | \n", i, "TX MISSING", "", ""))
			continue
		}
		if tx.commitReport != nil {
			committedAt = strconv.Itoa(int(tx.commitReport.Raw.BlockNumber))
		}

		if tx.message == nil {
			sb.WriteString(fmt.Sprintf("| %18s | %18s | %20v | %18d | %18s | \n", "MISSING", committedAt, ExecutionStatus(tx.execStatus.State), tx.execStatus.Raw.BlockNumber, "-"))
		} else if tx.execStatus != nil {
			sb.WriteString(fmt.Sprintf("| %18d | %18s | %20v | %18d | %18d | %s \n",
				tx.message.Message.SequenceNumber,
				committedAt,
				ExecutionStatus(tx.execStatus.State),
				tx.execStatus.Raw.BlockNumber,
				tx.message.Message.Nonce,
				helpers.ExplorerLink(int64(destination.ChainConfig.ChainId), tx.execStatus.Raw.TxHash)))
		} else {
			sb.WriteString(fmt.Sprintf("| %18d | %18s | %20v | %18s | %18d | %s \n",
				tx.message.Message.SequenceNumber,
				committedAt,
				"-",
				"-",
				tx.message.Message.Nonce,
				""))
		}
	}
	sb.WriteString(generateSeparator(headerLengths))

	destination.Logger.Info(sb.String())
}

func printDappSanityCheck(source *rhea.EvmDeploymentConfig) {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("Dapp sanity checks for %s\n", helpers.ChainName(int64(source.ChainConfig.ChainId))))

	tableHeaders := []string{"Dapp", "Router Set"}
	headerLengths := []int{30, 14}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))

	//receiverDap, err := receiver_dapp.NewReceiverDapp(source.LaneConfig.ReceiverDapp, source.Client)
	//helpers.PanicErr(err)
	//router, err = receiverDap.SRouter(&bind.CallOpts{})
	//helpers.PanicErr(err)
	//sb.WriteString(fmt.Sprintf("| %-30s | %14s |\n", "Receiver dapp", printBool(router == source.ChainConfig.Router)))

	if source.LaneConfig.PingPongDapp != common.HexToAddress("") {
		pingDapp, err := ping_pong_demo.NewPingPongDemo(source.LaneConfig.PingPongDapp, source.Client)
		helpers.PanicErr(err)
		router, err := pingDapp.GetRouter(&bind.CallOpts{})
		helpers.PanicErr(err)
		sb.WriteString(fmt.Sprintf("| %-30s | %14s |\n", "Ping dapp sender", printBool(router == source.ChainConfig.Router)))
	}

	sb.WriteString(generateSeparator(headerLengths))

	source.Logger.Info(sb.String())
}

func printRampSanityCheck(chain *rhea.EvmDeploymentConfig, sourceOnRamp common.Address) {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("Ramp checks for %s\n", helpers.ChainName(int64(chain.ChainConfig.ChainId))))

	tableHeaders := []string{"Contract", "Config correct"}
	headerLengths := []int{30, 14}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))

	afn, err := afn_contract.NewAFNContract(chain.ChainConfig.Afn, chain.Client)
	helpers.PanicErr(err)
	badSignal, err := afn.BadSignalReceived(&bind.CallOpts{})
	helpers.PanicErr(err)

	sb.WriteString(fmt.Sprintf("| %-30s | %14s |\n", "AFN healthy", printBool(!badSignal)))

	onRamp, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(chain.LaneConfig.OnRamp, chain.Client)
	helpers.PanicErr(err)
	routerAddress, err := onRamp.GetRouter(&bind.CallOpts{})
	helpers.PanicErr(err)
	sb.WriteString(fmt.Sprintf("| %-30s | %14s |\n", "OnRamp Router set", printBool(routerAddress == chain.ChainConfig.Router)))

	offRamp, err := evm_2_evm_offramp.NewEVM2EVMOffRamp(chain.LaneConfig.OffRamp, chain.Client)
	helpers.PanicErr(err)
	routerAddress, err = offRamp.GetRouter(&bind.CallOpts{})
	helpers.PanicErr(err)

	sb.WriteString(fmt.Sprintf("| %-30s | %14s |\n", "OffRamp Router set", printBool(routerAddress == chain.ChainConfig.Router)))

	configDetails, err := offRamp.LatestConfigDetails(&bind.CallOpts{})
	helpers.PanicErr(err)
	sb.WriteString(fmt.Sprintf("| %-30s | %14s |\n", "OffRamp OCR2 configured", printBool(configDetails.ConfigCount != 0)))

	commitStore, err := commit_store.NewCommitStore(chain.LaneConfig.CommitStore, chain.Client)
	helpers.PanicErr(err)

	blobConfigDetails, err := commitStore.LatestConfigDetails(&bind.CallOpts{})
	helpers.PanicErr(err)
	sb.WriteString(fmt.Sprintf("| %-30s | %14s |\n", "CommitStore OCR2 configured", printBool(blobConfigDetails.ConfigCount != 0)))

	router, err := router.NewRouter(chain.ChainConfig.Router, chain.Client)
	helpers.PanicErr(err)

	isRamp, err := router.IsOffRamp(&bind.CallOpts{}, chain.LaneConfig.OffRamp)
	helpers.PanicErr(err)

	sb.WriteString(fmt.Sprintf("| %-30s | %14s |\n", "Router has offRamp Set", printBool(isRamp)))

	sb.WriteString(generateSeparator(headerLengths))

	chain.Logger.Info(sb.String())
}

func printRateLimitingStatus(chain *rhea.EvmDeploymentConfig) {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("Rate limits for %s\n", helpers.ChainName(int64(chain.ChainConfig.ChainId))))

	tableHeaders := []string{"Contract", "Tokens"}
	headerLengths := []int{25, 42}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))

	onRamp, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(chain.LaneConfig.OnRamp, chain.Client)
	helpers.PanicErr(err)
	bucketState, err := onRamp.CalculateCurrentTokenBucketState(&bind.CallOpts{})
	helpers.PanicErr(err)

	sb.WriteString(fmt.Sprintf("| %-25s | %42d |\n", "onramp", bucketState.Tokens))

	offRamp, err := evm_2_evm_offramp.NewEVM2EVMOffRamp(chain.LaneConfig.OffRamp, chain.Client)
	helpers.PanicErr(err)
	offRampBucketState, err := offRamp.CalculateCurrentTokenBucketState(&bind.CallOpts{})
	helpers.PanicErr(err)

	sb.WriteString(fmt.Sprintf("| %-25s | %42d |\n", "offramp", offRampBucketState.Tokens))

	sb.WriteString(generateSeparator(headerLengths))
	chain.Logger.Info(sb.String())
}

func printPaused(chain *rhea.EvmDeploymentConfig) {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("Paused addresses for %s\n", helpers.ChainName(int64(chain.ChainConfig.ChainId))))

	tableHeaders := []string{"Contract", "Address", "Running"}
	headerLengths := []int{25, 42, 14}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))

	for _, tokenConfig := range chain.ChainConfig.SupportedTokens {
		tokenPool, err := lock_release_token_pool.NewLockReleaseTokenPool(tokenConfig.Pool, chain.Client)
		helpers.PanicErr(err)
		paused, err := tokenPool.Paused(&bind.CallOpts{})
		helpers.PanicErr(err)

		sb.WriteString(fmt.Sprintf("| %-25s | %42s | %14s |\n", "token pool", tokenConfig.Pool.Hex(), printBool(!paused)))
	}

	onRamp, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(chain.LaneConfig.OnRamp, chain.Client)
	helpers.PanicErr(err)
	paused, err := onRamp.Paused(&bind.CallOpts{})
	helpers.PanicErr(err)

	sb.WriteString(fmt.Sprintf("| %-25s | %42s | %14s |\n", "onramp", onRamp.Address(), printBool(!paused)))

	offRamp, err := evm_2_evm_offramp.NewEVM2EVMOffRamp(chain.LaneConfig.OffRamp, chain.Client)
	helpers.PanicErr(err)
	paused, err = offRamp.Paused(&bind.CallOpts{})
	helpers.PanicErr(err)

	sb.WriteString(fmt.Sprintf("| %-25s | %42s | %14s |\n", "offramp", offRamp.Address(), printBool(!paused)))

	commitStore, err := commit_store.NewCommitStore(chain.LaneConfig.CommitStore, chain.Client)
	helpers.PanicErr(err)
	paused, err = commitStore.Paused(&bind.CallOpts{})
	helpers.PanicErr(err)

	sb.WriteString(fmt.Sprintf("| %-25s | %42s | %14s |\n", "commitStore", commitStore.Address(), printBool(!paused)))

	sb.WriteString(generateSeparator(headerLengths))
	chain.Logger.Info(sb.String())
}

func PrintNodeBalances(chain *rhea.EvmDeploymentConfig, addresses []common.Address) {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("Node balances for %s\n", helpers.ChainName(int64(chain.ChainConfig.ChainId))))

	tableHeaders := []string{"Sender", "Balance"}
	headerLengths := []int{42, 18}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))

	for _, sender := range addresses {
		balanceAt, err := chain.Client.BalanceAt(context.Background(), sender, nil)
		helpers.PanicErr(err)

		sb.WriteString(fmt.Sprintf("| %42s |   %-16s |\n", sender.Hex(), new(big.Float).Quo(new(big.Float).SetInt(balanceAt), big.NewFloat(1e18)).String()))
	}

	sb.WriteString(generateSeparator(headerLengths))
	chain.Logger.Info(sb.String())
}

func printPoolBalances(chain *rhea.EvmDeploymentConfig) {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("Pool balances for %s\n", helpers.ChainName(int64(chain.ChainConfig.ChainId))))

	tableHeaders := []string{"Token", "Pool", "Balance", "Onramp", "OffRamp", "Price"}
	headerLengths := []int{32, 42, 20, 9, 9, 10}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))

	onRamp, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(chain.LaneConfig.OnRamp, chain.Client)
	helpers.PanicErr(err)

	for tokenName, tokenConfig := range chain.ChainConfig.SupportedTokens {
		tokenPool, err := lock_release_token_pool.NewLockReleaseTokenPool(tokenConfig.Pool, chain.Client)
		helpers.PanicErr(err)

		tokenAddress, err := tokenPool.GetToken(&bind.CallOpts{})
		helpers.PanicErr(err)

		tokenInstance, err := link_token_interface.NewLinkToken(tokenAddress, chain.Client)
		helpers.PanicErr(err)

		price, err := onRamp.GetPricesForTokens(&bind.CallOpts{}, []common.Address{tokenAddress})
		helpers.PanicErr(err)

		balance, err := tokenInstance.BalanceOf(&bind.CallOpts{}, tokenConfig.Pool)
		helpers.PanicErr(err)

		isAllowedOnRamp, err := tokenPool.IsOnRamp(&bind.CallOpts{}, chain.LaneConfig.OnRamp)
		helpers.PanicErr(err)

		isAllowedOffRamp, err := tokenPool.IsOffRamp(&bind.CallOpts{}, chain.LaneConfig.OffRamp)
		helpers.PanicErr(err)

		if tokenAddress != tokenConfig.Token {
			sb.WriteString(fmt.Sprintf("| %-32s | TOKEN CONFIG MISMATCH ❌ | expected %s | pool token %s |\n", tokenName, tokenConfig.Token.Hex(), tokenAddress.Hex()))
		} else {
			sb.WriteString(fmt.Sprintf("| %-32s | %s | %20d | %9s | %9s | %10s |\n", tokenName, tokenConfig.Pool, balance, printBool(isAllowedOnRamp), printBool(isAllowedOffRamp), price[0].String()))
		}
	}

	sb.WriteString(generateSeparator(headerLengths))

	chain.Logger.Info(sb.String())
}

func printSupportedTokensCheck(source *rhea.EvmDeploymentConfig, destination *rhea.EvmDeploymentConfig) {
	sourceRouter, err := router.NewRouter(source.ChainConfig.Router, source.Client)
	helpers.PanicErr(err)

	sourceTokens, err := sourceRouter.GetSupportedTokens(&bind.CallOpts{}, destination.ChainConfig.ChainId)
	helpers.PanicErr(err)

	destRouter, err := router.NewRouter(destination.ChainConfig.Router, destination.Client)
	helpers.PanicErr(err)

	destTokens, err := destRouter.GetSupportedTokens(&bind.CallOpts{}, source.ChainConfig.ChainId)
	helpers.PanicErr(err)

	var sb strings.Builder
	sb.WriteString("\nToken matching\n")

	tableHeaders := []string{"Token", "Source", "Pool", "Dest", "Pool"}
	headerLengths := []int{20, 10, 10, 10, 10}

	sb.WriteString(generateHeader(tableHeaders, headerLengths))

	for _, token := range rhea.GetAllTokens() {
		sourcePool := false
		if val, ok := source.ChainConfig.SupportedTokens[token]; ok && val.Pool != common.HexToAddress("") {
			sourcePool = true
		}
		destPool := false
		if val, ok := destination.ChainConfig.SupportedTokens[token]; ok && val.Pool != common.HexToAddress("") {
			destPool = true
		}
		sourceEnabled := slices.Contains(sourceTokens, source.ChainConfig.SupportedTokens[token].Token)
		destEnabled := slices.Contains(destTokens, destination.ChainConfig.SupportedTokens[token].Token)

		if sourceEnabled || destEnabled {
			sb.WriteString(fmt.Sprintf("| %-20s | %10s | %10s | %10s | %10s |\n", token, printBool(sourceEnabled), printBool(sourcePool), printBool(destEnabled), printBool(destPool)))
		} else {
			sb.WriteString(fmt.Sprintf("| %-20s | %10s | %10s | %10s | %10s |\n", token, printBoolNeutral(sourceEnabled), printBoolNeutral(sourcePool), printBoolNeutral(destEnabled), printBoolNeutral(destPool)))
		}

	}

	sb.WriteString(generateSeparator(headerLengths))

	source.Logger.Info(sb.String())
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
	return strings.Repeat("─", length) + "\n"
}

// PrintJobSpecs prints the job spec for each node and CCIP spec type, as well as a bootstrap spec.
func PrintJobSpecs(env dione.Environment, sourceClient rhea.EvmDeploymentConfig, destClient rhea.EvmDeploymentConfig) {
	don := dione.NewOfflineDON(env, nil)
	// jobparams for the lane
	jobParams := dione.NewCCIPJobSpecParams(sourceClient, destClient)

	bootstrapSpec := jobParams.BootstrapJob(destClient.LaneConfig.CommitStore.Hex())
	specString, err := bootstrapSpec.String()
	helpers.PanicErr(err)
	jobs := fmt.Sprintf("# BootstrapSpec%s", specString)

	commitJobSpec, err := jobParams.CommitJobSpec()
	helpers.PanicErr(err)
	committingChainID := commitJobSpec.OCR2OracleSpec.RelayConfig["chainID"].(uint64)
	executionSpec, err := jobParams.ExecutionJobSpec()
	helpers.PanicErr(err)
	execChainID := executionSpec.OCR2OracleSpec.RelayConfig["chainID"].(uint64)
	for i, oracle := range don.Config.Nodes {
		jobs += fmt.Sprintf("\n// [Node %d]\n", i)
		evmKeyBundle := dione.GetOCRkeysForChainType(oracle.OCRKeys, "evm")
		transmitterIDs := oracle.EthKeys

		// set node specific values
		commitJobSpec.OCR2OracleSpec.OCRKeyBundleID.SetValid(evmKeyBundle.ID)
		commitJobSpec.OCR2OracleSpec.TransmitterID.SetValid(transmitterIDs[fmt.Sprintf("%v", committingChainID)])
		specString, err := commitJobSpec.String()
		helpers.PanicErr(err)
		jobs += fmt.Sprintf("\n# CCIP commit spec%s", specString)

		// set node specific values
		executionSpec.OCR2OracleSpec.OCRKeyBundleID.SetValid(evmKeyBundle.ID)
		executionSpec.OCR2OracleSpec.TransmitterID.SetValid(transmitterIDs[fmt.Sprintf("%v", execChainID)])
		specString, err = executionSpec.String()
		helpers.PanicErr(err)
		jobs += fmt.Sprintf("\n# CCIP execution spec%s", specString)
	}
	fmt.Println(jobs)
}
