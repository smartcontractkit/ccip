package commit

import (
	"context"
	"testing"

	"github.com/smartcontractkit/ccipocr3/internal/libs/testhelpers"
	"github.com/smartcontractkit/ccipocr3/internal/mocks"
	"github.com/smartcontractkit/ccipocr3/internal/model"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

func TestPlugin(t *testing.T) {
	testCases := []struct {
		name        string
		description string
		nodes       []nodeSetup
		expErr      func(*testing.T, error)
	}{
		{
			name:        "EmptyOutcome",
			description: "Empty observations are returned by all nodes which leads to an empty outcome.",
			nodes:       setupEmptyOutcome(t, context.Background(), logger.Test(t)),
			expErr:      func(t *testing.T, err error) { assert.Equal(t, testhelpers.ErrEmptyOutcome, err) },
		},
	}

	ctx := context.Background()
	lggr := logger.Test(t)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log("-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-")
			t.Logf(">>> [%s]\n", tc.name)
			t.Logf(">>> %s\n", tc.description)
			defer t.Log("-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-")

			nodesSetup := setupEmptyOutcome(t, ctx, lggr)
			nodes := make([]ocr3types.ReportingPlugin[[]byte], 0, len(nodesSetup))
			for _, n := range nodesSetup {
				nodes = append(nodes, n.node)
			}
			runner := testhelpers.NewOCR3Runner(nodes)

			_, err := runner.RunRound(ctx)
			tc.expErr(t, err)
		})
	}
}

func setupEmptyOutcome(t *testing.T, ctx context.Context, lggr logger.Logger) []nodeSetup {
	return []nodeSetup{
		newNode(t, ctx, lggr, 1),
		newNode(t, ctx, lggr, 2),
		newNode(t, ctx, lggr, 3),
	}
}

type nodeSetup struct {
	node        *Plugin
	ccipReader  *mocks.CCIPReader
	priceReader *mocks.TokenPricesReader
	reportCodec *mocks.CommitPluginJSONReportCodec
	msgHasher   *mocks.MessageHasher
}

func newNode(t *testing.T, ctx context.Context, lggr logger.Logger, id int) nodeSetup {
	ccipReader := mocks.NewCCIPReader()
	priceReader := mocks.NewTokenPricesReader()
	reportCodec := mocks.NewCommitPluginJSONReportCodec()
	msgHasher := mocks.NewMessageHasher()

	node1 := NewPlugin(
		context.Background(),
		commontypes.OracleID(id),
		model.CommitPluginConfig{
			Writer:    false,
			Reads:     nil,
			DestChain: chainC,
			FChain: map[model.ChainSelector]int{
				chainC: 1,
			},
			ObserverInfo:        nil,
			PricedTokens:        []types.Account{tokenX},
			TokenPricesObserver: false,
			NewMsgScanBatchSize: 256,
		},
		ccipReader,
		priceReader,
		reportCodec,
		msgHasher,
		lggr,
	)

	return nodeSetup{
		node:        node1,
		ccipReader:  ccipReader,
		priceReader: priceReader,
		reportCodec: reportCodec,
		msgHasher:   msgHasher,
	}
}

var (
	chainA = model.ChainSelector(1)
	chainB = model.ChainSelector(2)
	chainC = model.ChainSelector(3)

	tokenX = types.Account("tk_xxx")
	tokenY = types.Account("tk_yyy")
	tokenZ = types.Account("tk_zzz")
)
