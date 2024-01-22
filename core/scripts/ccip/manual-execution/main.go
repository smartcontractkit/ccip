package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"

	"manual-execution/helpers"
	"manual-execution/lib"
)

const NumberOfBlocks = 5000

func main() {
	configPath := flag.String("configFile", "./config.json", "config for manually executing a failed ccip message "+
		"which has been successfully committed but failed to get executed")
	flag.Parse()

	if *configPath == "" {
		log.Println("config json is required")
		os.Exit(1)
	}
	cData, err := os.ReadFile(*configPath)
	if err != nil {
		log.Println("unable to read the json at ", *configPath, "error - ", err)
		os.Exit(1)
	}
	var cfg lib.Config
	err = json.Unmarshal(cData, &cfg)
	if err != nil {
		log.Println("unable to marshal the json at ", *configPath, "error - ", err, `sample json
{
	"src_rpc": "",
	"dest_rpc": "",
	"dest_owner_key": "",
	"commit_store": "",
	"off_ramp": "",
	"dest_start_block": "",
	"ccip_send_tx": "",
	"source_start_block": "",
	"dest_deployed_at": 0,
	"gas_limit_override": 0,
}`)
		os.Exit(1)
	}
	// mandatory fields check
	err = cfg.VerifyConfig()
	if err != nil {
		log.Println("config validation failed: \n", err)
		os.Exit(1)
	}

	err = run(cfg)
	if err != nil {
		log.Println("manual execution was not successful - ", err)
		os.Exit(1)
	}
}

// run strings together library calls to create a report, then executes the report.
func run(cfg lib.Config) error {
	clients, err := lib.GetClients(context.Background(), cfg.SrcNodeURL, cfg.DestNodeURL)
	if err != nil {
		return fmt.Errorf("unable to initialize clients: %w", err)
	}

	proofInput, err := lib.GetChainData(clients, cfg.SourceChainTx, cfg.CCIPMsgID, cfg.DestStartBlock, cfg.DestDeployedAt, NumberOfBlocks)
	if err != nil {
		return fmt.Errorf("unable to get chain data: %w", err)
	}

	report, err := lib.MakeExecutionReport(cfg, clients, proofInput)
	if err != nil {
		return fmt.Errorf("unable to make offRamp execution report: %w", err)
	}

	return execute(cfg, clients, proofInput, report)
}

// execute report on the destination chain.
func execute(cfg lib.Config, clients lib.ClientData, proofReport lib.ProofData, report helpers.InternalExecutionReport) error {
	ownerKey, err := crypto.HexToECDSA(cfg.DestOwner)
	if err != nil {
		return fmt.Errorf("unable to get destination owner: %w", err)
	}

	destUser, err := bind.NewKeyedTransactorWithChainID(ownerKey, clients.DestChainId)
	if err != nil {
		return err
	}
	log.Println("--- Owner address---/n", destUser.From.Hex())

	// Execute.
	gasLimitOverrides := make([]*big.Int, len(report.Messages))
	for i := range report.Messages {
		gasLimitOverrides[i] = big.NewInt(int64(cfg.GasLimitOverride))
	}

	tx, err := helpers.ManuallyExecute(clients.DestChain, destUser, cfg.OffRamp, report, gasLimitOverrides)
	if err != nil {
		return err
	}
	// wait for tx confirmation
	err = helpers.WaitForSuccessfulTxReceipt(clients.DestChain, tx.Hash())
	if err != nil {
		return err
	}

	// check if the message got successfully delivered
	changed, err := helpers.FilterExecutionStateChanged(clients.DestChain, &bind.FilterOpts{
		Start: proofReport.DestStartBlock,
	}, cfg.OffRamp, []uint64{proofReport.SeqNum}, [][32]byte{proofReport.MsgID})
	if err != nil {
		return err
	}
	if changed != 2 {
		return fmt.Errorf("manual execution did not result in ExecutionStateChanged as success")
	}
	return nil
}
