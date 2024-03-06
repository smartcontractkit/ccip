package cache

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/ccipdata/mocks"
)

func Test_ArmChainStateCacheCapabilities(t *testing.T) {
	ctx := tests.Context(t)
	lggr := logger.TestLogger(t)
	mockCommitStore := mocks.NewCommitStoreReader(t)
	mockOnRamp := mocks.NewOnRampReader(t)

	chainState := newArmChainWithCustomEviction(lggr, mockOnRamp, mockCommitStore, 10*time.Hour)

	// Chain is not cursed
	mockCommitStore.On("IsDown", ctx).Return(false, nil).Once()
	mockOnRamp.On("IsSourceCursed", ctx).Return(false, nil).Once()
	assert.NoError(t, chainState.ValidateNotCursed(ctx))

	// Chain is cursed, but cache is stale
	mockCommitStore.On("IsDown", ctx).Return(true, nil).Once()
	mockOnRamp.On("IsSourceCursed", ctx).Return(true, nil).Once()
	assert.NoError(t, chainState.ValidateNotCursed(ctx))

	// Enforce cache refresh
	err := chainState.ForceValidateNotCursed(ctx)
	assert.Error(t, err)
	assert.Equal(t, ccip.ErrChainPausedOrCursed, err)
}

func Test_ArmChainState(t *testing.T) {
	ctx := tests.Context(t)

	testCases := []struct {
		name            string
		commitStoreDown bool
		commitStoreErr  error
		onRampCursed    bool
		onRampErr       error
		expectedErr     bool
	}{
		{
			name:        "Neither down nor cursed",
			expectedErr: false,
		},
		{
			name:            "CommitStore is down",
			commitStoreDown: true,
			expectedErr:     true,
		},
		{
			name:         "OnRamp is cursed",
			onRampCursed: true,
			expectedErr:  true,
		},
		{
			name:           "CommitStore error",
			commitStoreErr: errors.New("commit store error"),
			expectedErr:    true,
		},
		{
			name:        "OnRamp error",
			onRampErr:   errors.New("onramp error"),
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCommitStore := mocks.NewCommitStoreReader(t)
			mockOnRamp := mocks.NewOnRampReader(t)

			mockCommitStore.On("IsDown", ctx).Return(tc.commitStoreDown, tc.commitStoreErr)
			mockOnRamp.On("IsSourceCursed", ctx).Return(tc.onRampCursed, tc.onRampErr)

			chainState := NewArmChainState(logger.TestLogger(t), mockOnRamp, mockCommitStore)
			err := chainState.ValidateNotCursed(ctx)

			if tc.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
