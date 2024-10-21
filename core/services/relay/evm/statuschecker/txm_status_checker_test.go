package statuschecker

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
)

func Test_CheckMessageStatus(t *testing.T) {
	testutils.SkipShort(t, "")
	ctx := context.Background()
	mockTxManager := mocks.NewMockEvmTxManager(t)
	checker := NewTxmStatusChecker(mockTxManager.GetTransactionStatus)

	msgID := "test-message-id"

	// Define test cases
	testCases := []struct {
		name            string
		setupMock       func()
		expectedStatus  []TransactionStatusWithError
		expectedCounter int
		expectedError   error
	}{
		{
			name: "No transactions found",
			setupMock: func() {
				mockTxManager.Mock = mock.Mock{}
				mockTxManager.On("GetTransactionStatus", ctx, "test-message-id-0").Return(types.Unknown, errors.New("failed to find transaction with IdempotencyKey test-message-id-0"))
			},
			expectedStatus:  []TransactionStatusWithError{},
			expectedCounter: -1,
			expectedError:   nil,
		},
		{
			name: "Single transaction found",
			setupMock: func() {
				mockTxManager.Mock = mock.Mock{}
				mockTxManager.On("GetTransactionStatus", ctx, "test-message-id-0").Return(types.Finalized, nil)
				mockTxManager.On("GetTransactionStatus", ctx, "test-message-id-1").Return(types.Unknown, errors.New("failed to find transaction with IdempotencyKey test-message-id-1"))
			},
			expectedStatus:  []TransactionStatusWithError{{Status: types.Finalized, Error: nil}},
			expectedCounter: 0,
			expectedError:   nil,
		},
		{
			name: "Multiple transactions found",
			setupMock: func() {
				mockTxManager.Mock = mock.Mock{}
				mockTxManager.On("GetTransactionStatus", ctx, "test-message-id-0").Return(types.Finalized, nil)
				mockTxManager.On("GetTransactionStatus", ctx, "test-message-id-1").Return(types.Failed, errors.New("dummy error"))
				mockTxManager.On("GetTransactionStatus", ctx, "test-message-id-2").Return(types.Unknown, errors.New("failed to find transaction with IdempotencyKey test-message-id-2"))
			},
			expectedStatus: []TransactionStatusWithError{
				{Status: types.Finalized, Error: nil},
				{Status: types.Failed, Error: errors.New("dummy error")},
			},
			expectedCounter: 1,
			expectedError:   nil,
		},
		{
			name: "Unknown status without nil (in progress)",
			setupMock: func() {
				mockTxManager.Mock = mock.Mock{}
				mockTxManager.On("GetTransactionStatus", ctx, "test-message-id-0").Return(types.Unknown, nil)
				mockTxManager.On("GetTransactionStatus", ctx, "test-message-id-1").Return(types.Unknown, errors.New("failed to find transaction with IdempotencyKey test-message-id-1"))
			},
			expectedStatus: []TransactionStatusWithError{
				{Status: types.Unknown, Error: nil},
			},
			expectedCounter: 0,
			expectedError:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()
			statuses, counter, err := checker.CheckMessageStatus(ctx, msgID)
			assert.Equal(t, tc.expectedStatus, statuses)
			assert.Equal(t, tc.expectedCounter, counter)
			assert.Equal(t, tc.expectedError, err)
			mockTxManager.AssertExpectations(t)
		})
	}
}

func Test_FailForMoreThan1000Retries(t *testing.T) {
	ctx := context.Background()
	mockTxManager := mocks.NewMockEvmTxManager(t)
	checker := NewTxmStatusChecker(mockTxManager.GetTransactionStatus)

	for i := 0; i < 1000; i++ {
		mockTxManager.On("GetTransactionStatus", ctx, fmt.Sprintf("test-message-id-%d", i)).Return(types.Finalized, nil)
	}

	msgID := "test-message-id"
	_, _, err := checker.CheckMessageStatus(ctx, msgID)
	assert.EqualError(t, err, "maximum number of statuses reached, possible infinite loop")
}
