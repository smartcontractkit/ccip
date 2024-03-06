package opstack

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip/rebalancer/multienv"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/optimism_portal"
)

func FinalizeL1(
	env multienv.Env,
	l1ChainID,
	l2ChainID uint64,
	l1BridgeAdapterAddress,
	optimismPortalAddress common.Address,
	l2TxHash common.Hash,
) {
	l2Client, ok := env.Clients[l2ChainID]
	if !ok {
		panic(fmt.Sprintf("No L2 client found for chain ID %d", l2ChainID))
	}

	receipt, err := l2Client.TransactionReceipt(context.Background(), l2TxHash)
	helpers.PanicErr(err)

	messagePassedLog := getMessagePassedLog(receipt.Logs)
	if messagePassedLog == nil {
		panic(fmt.Sprintf("No message passed log found in receipt %s", receipt.TxHash.String()))
	}

	messagePassed := parseMessagePassedLog(messagePassedLog)

	lowLevelMsg := toLowLevelMessage(messagePassed)

	portal, err := optimism_portal.NewOptimismPortal(optimismPortalAddress, env.Clients[l1ChainID])
	helpers.PanicErr(err)

	tx, err := portal.FinalizeWithdrawalTransaction(
		env.Transactors[l1ChainID],
		optimism_portal.TypesWithdrawalTransaction{
			Nonce:    lowLevelMsg.MessageNonce,
			Sender:   lowLevelMsg.Sender,
			Target:   lowLevelMsg.Target,
			Value:    lowLevelMsg.Value,
			GasLimit: lowLevelMsg.MinGasLimit,
			Data:     lowLevelMsg.Message,
		},
	)
	helpers.PanicErr(err)

	helpers.ConfirmTXMined(context.Background(), env.Clients[l1ChainID], tx, int64(l1ChainID), "FinalizeWithdrawalTransaction")
}
