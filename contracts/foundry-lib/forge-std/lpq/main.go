package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/sqlx"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp"
	logger2 "github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/abihelpers"
	"github.com/smartcontractkit/chainlink/v2/core/services/pg"
)

func pie(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sqlx.Connect("postgres", "postgresql://aurora-clnodes-ccip-prod-1-2.c94fupu7o8i1.us-west-2.rds.amazonaws.com:5432/main-chainlink-clc-ocr2-eth-goerli-ccip-nodes-0?user=temp_dev_ro&password=SrFM4gBoQBAv4v")
	pie(err)

	defer db.Close()

	o := logpoller.NewORM(big.NewInt(420), db, logger2.NullLogger, pg.NewQConfig(true))
	onRampAddr := common.HexToAddress("0xEe55842b1D68224d9eEF238d4736E851db613630")
	logs, err := o.SelectDataWordRange(
		onRampAddr,
		abihelpers.EventSignatures.SendRequested,
		abihelpers.EventSignatures.SendRequestedSequenceNumberWord,
		logpoller.EvmWord(22175),
		logpoller.EvmWord(22175),
		10,
	)
	pie(err)

	for _, log := range logs {
		onRamp, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(onRampAddr, nil)
		pie(err)

		ccipSendRequested, err := onRamp.ParseCCIPSendRequested(gethtypes.Log{
			Topics: log.GetTopics(),
			Data:   log.Data,
		})
		pie(err)

		fmt.Println(ccipSendRequested.Message.SequenceNumber)
	}

	fmt.Println("shit")
}
