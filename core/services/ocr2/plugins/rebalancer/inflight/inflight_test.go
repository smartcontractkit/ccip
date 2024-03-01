package inflight

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	ubig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

func Test_inflight_Add(t *testing.T) {
	type fields struct {
		items map[mapKey]models.Transfer
		lggr  logger.Logger
	}
	type args struct {
		ctx context.Context
		t   models.Transfer
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		assertion func(*testing.T, *inflight)
	}{
		{
			"transfer not in map",
			fields{
				items: map[mapKey]models.Transfer{},
				lggr:  logger.TestLogger(t),
			},
			args{
				testutils.Context(t),
				models.Transfer{
					From:   1,
					To:     2,
					Amount: ubig.NewI(1),
				},
			},
			func(t *testing.T, i *inflight) {
				item, ok := i.items[mapKey{From: 1, To: 2, Amount: "1"}]
				require.True(t, ok)
				require.Equal(t, models.Transfer{
					From:   1,
					To:     2,
					Amount: ubig.NewI(1),
				}, item)
			},
		},
		{
			"transfer in map",
			fields{
				items: map[mapKey]models.Transfer{
					{From: 1, To: 2, Amount: "1"}: {
						From:   1,
						To:     2,
						Amount: ubig.NewI(1),
					},
				},
				lggr: logger.TestLogger(t),
			},
			args{
				testutils.Context(t),
				models.Transfer{
					From:   1,
					To:     2,
					Amount: ubig.NewI(1),
				},
			},
			func(t *testing.T, i *inflight) {
				require.Len(t, i.items, 1)
				item, ok := i.items[mapKey{From: 1, To: 2, Amount: "1"}]
				require.True(t, ok)
				require.Equal(t, models.Transfer{
					From:   1,
					To:     2,
					Amount: ubig.NewI(1),
				}, item)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &inflight{
				items: tt.fields.items,
				lggr:  tt.fields.lggr,
			}
			i.Add(tt.args.ctx, tt.args.t)
			tt.assertion(t, i)
		})
	}
}

func Test_inflight_Expire(t *testing.T) {
	type fields struct {
		items map[mapKey]models.Transfer
		lggr  logger.Logger
	}
	type args struct {
		ctx     context.Context
		pending []models.PendingTransfer
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		before    func(*testing.T, *inflight)
		assertion func(*testing.T, *inflight)
	}{
		{
			"no pending transfers",
			fields{
				items: map[mapKey]models.Transfer{
					{From: 1, To: 2, Amount: "1"}: {
						From:   1,
						To:     2,
						Amount: ubig.NewI(1),
					},
				},
				lggr: logger.TestLogger(t),
			},
			args{
				testutils.Context(t),
				[]models.PendingTransfer{},
			},
			func(t *testing.T, i *inflight) {
				require.Len(t, i.items, 1)
			},
			func(t *testing.T, i *inflight) {
				require.Len(t, i.items, 1)
			},
		},
		{
			"pending transfer",
			fields{
				items: map[mapKey]models.Transfer{
					{From: 1, To: 2, Amount: "1"}: {
						From:   1,
						To:     2,
						Amount: ubig.NewI(1),
					},
				},
				lggr: logger.TestLogger(t),
			},
			args{
				testutils.Context(t),
				[]models.PendingTransfer{
					{
						Transfer: models.Transfer{
							From:   1,
							To:     2,
							Amount: ubig.NewI(1),
						},
					},
				},
			},
			func(t *testing.T, i *inflight) {
				require.Len(t, i.items, 1)
			},
			func(t *testing.T, i *inflight) {
				require.Len(t, i.items, 0)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &inflight{
				items: tt.fields.items,
				lggr:  tt.fields.lggr,
			}
			tt.before(t, i)
			i.Expire(tt.args.ctx, tt.args.pending)
			tt.assertion(t, i)
		})
	}
}

func Test_inflight_GetAll(t *testing.T) {
	type fields struct {
		items map[mapKey]models.Transfer
		lggr  logger.Logger
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []models.Transfer
	}{
		{
			"empty",
			fields{
				items: map[mapKey]models.Transfer{},
				lggr:  logger.TestLogger(t),
			},
			args{
				testutils.Context(t),
			},
			[]models.Transfer{},
		},
		{
			"not empty",
			fields{
				items: map[mapKey]models.Transfer{
					{From: 1, To: 2, Amount: "1"}: {
						From:   1,
						To:     2,
						Amount: ubig.NewI(1),
					},
				},
				lggr: logger.TestLogger(t),
			},
			args{
				testutils.Context(t),
			},
			[]models.Transfer{
				{
					From:   1,
					To:     2,
					Amount: ubig.NewI(1),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &inflight{
				items: tt.fields.items,
				lggr:  tt.fields.lggr,
			}
			got := i.GetAll(tt.args.ctx)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_inflight_IsInflight(t *testing.T) {
	type fields struct {
		items map[mapKey]models.Transfer
		lggr  logger.Logger
	}
	type args struct {
		ctx context.Context
		t   models.Transfer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"not inflight",
			fields{
				items: map[mapKey]models.Transfer{},
				lggr:  logger.TestLogger(t),
			},
			args{
				testutils.Context(t),
				models.Transfer{
					From:   1,
					To:     2,
					Amount: ubig.NewI(1),
				},
			},
			false,
		},
		{
			"inflight",
			fields{
				items: map[mapKey]models.Transfer{
					{From: 1, To: 2, Amount: "1"}: {
						From:   1,
						To:     2,
						Amount: ubig.NewI(1),
					},
				},
				lggr: logger.TestLogger(t),
			},
			args{
				testutils.Context(t),
				models.Transfer{
					From:   1,
					To:     2,
					Amount: ubig.NewI(1),
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &inflight{
				items: tt.fields.items,
				lggr:  tt.fields.lggr,
			}
			got := i.IsInflight(tt.args.ctx, tt.args.t)
			require.Equal(t, tt.want, got)
		})
	}
}
