package cciprebalance

import (
	"context"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	mocks "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/cciprebalance/cciprebalancemocks"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/cciprebalance/liquiditygraph"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/cciprebalance/models"
	"github.com/smartcontractkit/chainlink/v2/core/utils"
)

func TestService(t *testing.T) {
	ctx := context.Background()

	t.Run("base e2e test with single network", func(t *testing.T) {
		lmFactory := mocks.NewFactory(t)

		oneTwoThree := models.NetworkID(123)
		oneTwoThreeLM := models.Address(utils.RandomAddress())
		mockLM := mocks.NewLiquidityManager(t)

		s := NewService(
			models.Address(utils.RandomAddress()),
			lmFactory,
			map[models.NetworkID]models.Address{
				oneTwoThree: oneTwoThreeLM,
			},
			liquiditygraph.NewDummyGraph(),
		)

		lmFactory.On("NewLiquidityManager", oneTwoThree, oneTwoThreeLM).Return(mockLM, nil)
		mockLM.On("GetLiquidityManagers", ctx).Return(map[models.NetworkID]models.Address{}, nil)
		mockLM.On("GetBalance", ctx).Return(big.NewInt(100), nil)
		mockLM.On("GetPendingTransfers", ctx).Return(nil, nil)

		err := s.Run(ctx)
		assert.NoError(t, err)
	})

	t.Run("large example", func(t *testing.T) {
		networkA := models.NetworkID(100)
		networkB := models.NetworkID(200)
		networkC := models.NetworkID(300)
		networkD := models.NetworkID(400)

		lmA := models.Address(utils.RandomAddress())
		lmB := models.Address(utils.RandomAddress())
		lmC := models.Address(utils.RandomAddress())
		lmD := models.Address(utils.RandomAddress())

		mockLMA := mocks.NewLiquidityManager(t)
		mockLMB := mocks.NewLiquidityManager(t)
		mockLMC := mocks.NewLiquidityManager(t)
		mockLMD := mocks.NewLiquidityManager(t)

		lmFactory := mocks.NewFactory(t)

		s := NewService(
			models.Address(utils.RandomAddress()),
			lmFactory,
			map[models.NetworkID]models.Address{
				networkA: lmA,
				networkB: lmB,
				networkC: lmC,
			},
			liquiditygraph.NewDummyGraph(),
		)

		mockLMA.On("GetPendingTransfers", ctx).Return(nil, nil)
		mockLMB.On("GetPendingTransfers", ctx).Return(nil, nil)
		mockLMC.On("GetPendingTransfers", ctx).Return([]models.Transfer{
			{From: networkA, To: networkB, Amount: big.NewInt(1000)},
		}, nil)
		mockLMD.On("GetPendingTransfers", ctx).Return(nil, nil)

		lmFactory.On("NewLiquidityManager", networkA, lmA).Return(mockLMA, nil)
		lmFactory.On("NewLiquidityManager", networkB, lmB).Return(mockLMB, nil)
		lmFactory.On("NewLiquidityManager", networkC, lmC).Return(mockLMC, nil)
		lmFactory.On("NewLiquidityManager", networkD, lmD).Return(mockLMD, nil)

		mockLMA.On("GetLiquidityManagers", ctx).Return(map[models.NetworkID]models.Address{
			networkB: lmB,
			networkC: lmC,
		}, nil)

		mockLMB.On("GetLiquidityManagers", ctx).Return(map[models.NetworkID]models.Address{
			networkA: lmA,
			networkC: lmC,
			networkD: lmD,
		}, nil)

		mockLMC.On("GetLiquidityManagers", ctx).Return(map[models.NetworkID]models.Address{
			networkA: lmA,
			networkB: lmB,
		}, nil)

		mockLMD.On("GetLiquidityManagers", ctx).Return(map[models.NetworkID]models.Address{
			networkB: lmB,
		}, nil)

		mockLMA.On("GetBalance", ctx).Return(big.NewInt(450), nil)
		mockLMB.On("GetBalance", ctx).Return(big.NewInt(300), nil)
		mockLMC.On("GetBalance", ctx).Return(big.NewInt(100), nil)
		mockLMD.On("GetBalance", ctx).Return(big.NewInt(150), nil)

		mockLMB.On("MoveLiquidity", ctx, networkA, big.NewInt(300)).Return(nil)
		mockLMC.On("MoveLiquidity", ctx, networkA, big.NewInt(100)).Return(nil)
		mockLMD.On("MoveLiquidity", ctx, networkA, big.NewInt(150)).Return(nil)

		err := s.Run(ctx)
		assert.NoError(t, err)

		assert.Len(t, s.pendingTransfers, 3+1) // 3+1 pending
	})
}
