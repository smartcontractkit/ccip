package liquidityrebalancer

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"

	ubig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/graph"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

var eth = models.NetworkSelector(1)
var opt = models.NetworkSelector(2)
var arb = models.NetworkSelector(3)

func TestTargetBalanceRebalancer_ComputeTransfersToBalance_arb_eth_opt(t *testing.T) {
	type transfer struct {
		from   models.NetworkSelector
		to     models.NetworkSelector
		am     int64
		status models.TransferStatus
	}

	testCases := []struct {
		name             string
		balances         map[models.NetworkSelector]int64
		targets          map[models.NetworkSelector]int64
		pendingTransfers []transfer
		expTransfers     []transfer
	}{
		{
			name:             "all above target",
			balances:         map[models.NetworkSelector]int64{eth: 1400, arb: 1000, opt: 1100},
			targets:          map[models.NetworkSelector]int64{eth: 1000, arb: 1000, opt: 1000},
			pendingTransfers: []transfer{},
			expTransfers:     []transfer{},
		},
		{
			name:             "arb below target",
			balances:         map[models.NetworkSelector]int64{eth: 1400, arb: 800, opt: 1100},
			targets:          map[models.NetworkSelector]int64{eth: 1000, arb: 1000, opt: 1000},
			pendingTransfers: []transfer{},
			expTransfers:     []transfer{{from: eth, to: arb, am: 200}},
		},
		{
			name:             "opt below target",
			balances:         map[models.NetworkSelector]int64{eth: 1400, arb: 1000, opt: 900},
			targets:          map[models.NetworkSelector]int64{eth: 1000, arb: 1000, opt: 1000},
			pendingTransfers: []transfer{},
			expTransfers:     []transfer{{from: eth, to: opt, am: 100}},
		},
		{
			name:             "eth below target",
			balances:         map[models.NetworkSelector]int64{eth: 900, arb: 1000, opt: 1300},
			targets:          map[models.NetworkSelector]int64{eth: 1000, arb: 1000, opt: 1000},
			pendingTransfers: []transfer{},
			expTransfers:     []transfer{{from: opt, to: eth, am: 100}},
		},
		{
			name:             "both opt and arb below target",
			balances:         map[models.NetworkSelector]int64{eth: 1500, arb: 800, opt: 900},
			targets:          map[models.NetworkSelector]int64{eth: 1000, arb: 1000, opt: 1000},
			pendingTransfers: []transfer{},
			expTransfers:     []transfer{{from: eth, to: opt, am: 100}, {from: eth, to: arb, am: 200}},
		},
		{
			name:             "arb below target but there is no full funding to reach target",
			balances:         map[models.NetworkSelector]int64{eth: 1100, arb: 800, opt: 1050},
			targets:          map[models.NetworkSelector]int64{eth: 1000, arb: 1000, opt: 1000},
			pendingTransfers: []transfer{},
			expTransfers:     []transfer{{from: eth, to: arb, am: 100}}, // transfer is made but without reaching target
		},
		{
			name:             "eth below target and requires two transfers to reach target",
			balances:         map[models.NetworkSelector]int64{eth: 800, arb: 1100, opt: 1150},
			targets:          map[models.NetworkSelector]int64{eth: 1000, arb: 1000, opt: 1000},
			pendingTransfers: []transfer{},
			expTransfers:     []transfer{{from: opt, to: eth, am: 150}, {from: arb, to: eth, am: 50}},
		},
		{
			name:             "eth below with two sources to reach target the highest one is selected",
			balances:         map[models.NetworkSelector]int64{eth: 900, arb: 2000, opt: 1800},
			targets:          map[models.NetworkSelector]int64{eth: 1000, arb: 1000, opt: 1000},
			pendingTransfers: []transfer{},
			expTransfers:     []transfer{{from: arb, to: eth, am: 100}},
		},
		{
			name:             "eth below with two sources to reach target the highest one is selected - reversed",
			balances:         map[models.NetworkSelector]int64{eth: 900, arb: 1800, opt: 2000},
			targets:          map[models.NetworkSelector]int64{eth: 1000, arb: 1000, opt: 1000},
			pendingTransfers: []transfer{},
			expTransfers:     []transfer{{from: opt, to: eth, am: 100}},
		},
		{
			name:             "eth below with two sources to reach target should be deterministic",
			balances:         map[models.NetworkSelector]int64{eth: 700, arb: 1400, opt: 1400},
			targets:          map[models.NetworkSelector]int64{eth: 1000, arb: 1000, opt: 1000},
			pendingTransfers: []transfer{},
			expTransfers:     []transfer{{from: opt, to: eth, am: 300}},
		},
		//{
		//	name:             "arb below target but there is inflight transfer that reaches balance",
		//	balances:         map[models.NetworkSelector]int64{eth: 1000, arb: 800, opt: 1000},
		//	targets:          map[models.NetworkSelector]int64{eth: 1000, arb: 1000, opt: 1000},
		//	pendingTransfers: []transfer{{from: eth, to: arb, am: 250}},
		//	expTransfers:     []transfer{},
		//},
		//{
		//	name:             "arb below target but there is inflight transfer (not seen on-chain) that reaches balance",
		//	balances:         map[models.NetworkSelector]int64{eth: 1000, arb: 800, opt: 1000},
		//	targets:          map[models.NetworkSelector]int64{eth: 1000, arb: 1000, opt: 1000},
		//	pendingTransfers: []transfer{{from: eth, to: arb, am: 250, status: models.TransferStatusProposed}},
		//	expTransfers:     []transfer{}, // we propose a transfer from balance that is inflight
		//},
	}

	lggr := logger.TestLogger(t)
	lggr.SetLogLevel(zapcore.DebugLevel)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			g := graph.NewGraph()
			for net, b := range tc.balances {
				g.AddNetwork(net, graph.Data{
					Liquidity:       big.NewInt(b),
					NetworkSelector: net,
					TargetLiquidity: big.NewInt(tc.targets[net]),
				})
			}
			assert.NoError(t, g.AddConnection(eth, arb))
			assert.NoError(t, g.AddConnection(arb, eth))
			assert.NoError(t, g.AddConnection(eth, opt))
			assert.NoError(t, g.AddConnection(opt, eth))

			r := NewTargetBalanceRebalancer(lggr)

			unexecuted := make([]UnexecutedTransfer, 0, len(tc.pendingTransfers))
			for _, tr := range tc.pendingTransfers {
				unexecuted = append(unexecuted, models.Transfer{
					From:   tr.from,
					To:     tr.to,
					Amount: ubig.New(big.NewInt(tr.am)),
				})
			}
			transfersToBalance, err := r.ComputeTransfersToBalance(g, unexecuted)
			assert.NoError(t, err)

			assert.Len(t, transfersToBalance, len(tc.expTransfers))
			for i, tr := range tc.expTransfers {
				assert.Equal(t, tr.from, transfersToBalance[i].From)
				assert.Equal(t, tr.to, transfersToBalance[i].To)
				assert.Equal(t, tr.am, transfersToBalance[i].Amount.Int64())
			}
		})
	}
}
