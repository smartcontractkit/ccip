package statuschecker

import (
	"context"
	"fmt"
	"strings"
)

//go:generate mockery --quiet --name CCIPTransactionStatusChecker --output ./mocks/ --case=underscore

// TODO: TO REMOVE - START
// replace with actual implementation coming from "github.com/smartcontractkit/chainlink-common/pkg/types"
type TransactionStatus int

const (
	Unknown TransactionStatus = iota
	Unconfirmed
	Finalized
	Failed
	Fatal
)

type TxManager interface {
	GetTransactionStatus(ctx context.Context, transactionID string) (TransactionStatus, error)
}

// TODO: TO REMOVE - END

// CCIPTransactionStatusChecker is an interface that defines the method for checking the status of a transaction.
// CheckMessageStatus checks the status of a transaction for a given message ID.
// It returns a list of transaction statuses, the retry counter, and an error if any occurred during the process.
type CCIPTransactionStatusChecker interface {
	CheckMessageStatus(ctx context.Context, msgID string) (transactionStatuses []TransactionStatus, retryCounter int, err error)
}

type TxmStatusChecker struct {
	txManager TxManager
}

func NewTxmStatusChecker(txManager TxManager) *TxmStatusChecker {
	return &TxmStatusChecker{txManager: txManager}
}

// CheckMessageStatus checks the status of a message by checking the status of all transactions associated with the message ID.
// It returns a slice of all statuses and the number of transactions found (-1 if none).
// The key will follow the format: <msgID>-<counter>. TXM will be queried for each key until a NotFound error is returned.
// The goal is to find all transactions associated with a message ID and snooze messages if they are fatal in the Execution Plugin.
func (tsc *TxmStatusChecker) CheckMessageStatus(ctx context.Context, msgID string) ([]TransactionStatus, int, error) {
	var allStatuses []TransactionStatus
	var counter int
	const maxStatuses = 1000 // Cap the number of statuses to avoid infinite loop

	for {
		transactionID := fmt.Sprintf("%s-%d", msgID, counter)
		status, err := tsc.txManager.GetTransactionStatus(ctx, transactionID)
		if err != nil {
			if strings.Contains(err.Error(), fmt.Sprintf("failed to find transaction with IdempotencyKey %s", transactionID)) {
				break
			}
			return nil, counter - 1, err
		}
		allStatuses = append(allStatuses, status)
		counter++

		// Break the loop if the cap is reached
		if counter >= maxStatuses {
			return allStatuses, counter - 1, fmt.Errorf("maximum number of statuses reached, possible infinite loop")
		}

	}

	return allStatuses, counter - 1, nil
}
