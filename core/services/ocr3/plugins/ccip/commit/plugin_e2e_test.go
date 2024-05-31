package commit

import (
	"context"
	"math/big"
	"reflect"
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
	ctx := context.Background()
	lggr := logger.Test(t)

	testCases := []struct {
		name        string
		description string
		nodes       []nodeSetup
		expErr      func(*testing.T, error)
		expOutcome  model.CommitPluginOutcome
	}{
		{
			name:        "EmptyOutcome",
			description: "Empty observations are returned by all nodes which leads to an empty outcome.",
			nodes:       setupEmptyOutcome(t, ctx, lggr),
			expErr:      func(t *testing.T, err error) { assert.Equal(t, testhelpers.ErrEmptyOutcome, err) },
		},
		{
			name: "AllNodesReadAllChains",
			description: "Nodes observe the latest sequence numbers and new messages after those sequence numbers. " +
				"They also observe gas prices. In this setup all nodes can read all chains.",
			nodes: setupAllNodesReadAllChains(t, ctx, lggr),
			expOutcome: model.CommitPluginOutcome{
				MaxSeqNums: []model.SeqNumChain{
					{ChainSel: chainA, SeqNum: 10},
					{ChainSel: chainB, SeqNum: 20},
				},
				MerkleRoots: []model.MerkleRootChain{
					{ChainSel: chainB, MerkleRoot: model.Bytes32{}, SeqNumsRange: model.NewSeqNumRange(21, 22)},
				},
				TokenPrices: []model.TokenPrice{},
				GasPrices: []model.GasPriceChain{
					{ChainSel: chainA, GasPrice: model.BigInt{Int: big.NewInt(1000)}},
					{ChainSel: chainB, GasPrice: model.BigInt{Int: big.NewInt(20000)}},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log("-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-")
			t.Logf(">>> [%s]\n", tc.name)
			t.Logf(">>> %s\n", tc.description)
			defer t.Log("-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-")

			nodesSetup := tc.nodes
			nodes := make([]ocr3types.ReportingPlugin[[]byte], 0, len(nodesSetup))
			for _, n := range nodesSetup {
				nodes = append(nodes, n.node)
			}

			nodeIDs := make([]commontypes.OracleID, 0, len(nodesSetup))
			for _, n := range nodesSetup {
				nodeIDs = append(nodeIDs, n.node.nodeID)
			}
			runner := testhelpers.NewOCR3Runner(nodes, nodeIDs)

			res, err := runner.RunRound(ctx)
			if tc.expErr != nil {
				tc.expErr(t, err)
			} else {
				assert.NoError(t, err)
			}

			if !reflect.DeepEqual(tc.expOutcome, model.CommitPluginOutcome{}) {
				outcome, err := model.DecodeCommitPluginOutcome(res.Outcome)
				assert.NoError(t, err)
				assert.Equal(t, tc.expOutcome.TokenPrices, outcome.TokenPrices)
				assert.Equal(t, tc.expOutcome.MaxSeqNums, outcome.MaxSeqNums)
				assert.Equal(t, tc.expOutcome.GasPrices, outcome.GasPrices)

				assert.Equal(t, len(tc.expOutcome.MerkleRoots), len(outcome.MerkleRoots))
				for i, exp := range tc.expOutcome.MerkleRoots {
					assert.Equal(t, exp.ChainSel, outcome.MerkleRoots[i].ChainSel)
					assert.Equal(t, exp.SeqNumsRange, outcome.MerkleRoots[i].SeqNumsRange)
				}
			}
		})
	}
}

func setupEmptyOutcome(t *testing.T, ctx context.Context, lggr logger.Logger) []nodeSetup {
	cfg := model.CommitPluginConfig{
		DestChain: chainC,
		FChain: map[model.ChainSelector]int{
			chainC: 1,
		},
		ObserverInfo:        map[commontypes.OracleID]model.ObserverInfo{},
		PricedTokens:        []types.Account{tokenX},
		TokenPricesObserver: false,
		NewMsgScanBatchSize: 256,
	}

	return []nodeSetup{
		newNode(t, ctx, lggr, 1, cfg),
		newNode(t, ctx, lggr, 2, cfg),
		newNode(t, ctx, lggr, 3, cfg),
	}
}

func setupAllNodesReadAllChains(t *testing.T, ctx context.Context, lggr logger.Logger) []nodeSetup {
	cfg := model.CommitPluginConfig{
		DestChain: chainC,
		FChain: map[model.ChainSelector]int{
			chainA: 1,
			chainB: 1,
			chainC: 1,
		},
		ObserverInfo: map[commontypes.OracleID]model.ObserverInfo{
			1: {Writer: true, Reads: []model.ChainSelector{chainA, chainB, chainC}},
			2: {Writer: true, Reads: []model.ChainSelector{chainA, chainB, chainC}},
			3: {Writer: true, Reads: []model.ChainSelector{chainA, chainB, chainC}},
		},
		PricedTokens:        []types.Account{tokenX},
		TokenPricesObserver: false,
		NewMsgScanBatchSize: 256,
	}

	n1 := newNode(t, ctx, lggr, 1, cfg)
	n2 := newNode(t, ctx, lggr, 2, cfg)
	n3 := newNode(t, ctx, lggr, 3, cfg)
	nodes := []nodeSetup{n1, n2, n3}

	for _, n := range nodes {
		// all nodes observe the same sequence numbers 10 for chainA and 20 for chainB
		n.ccipReader.On("NextSeqNum", ctx, []model.ChainSelector{chainA, chainB}).
			Return([]model.SeqNum{10, 20}, nil)

		// then they fetch new msgs, there is nothing new on chainA
		n.ccipReader.On(
			"MsgsBetweenSeqNums",
			ctx,
			chainA,
			model.NewSeqNumRange(11, model.SeqNum(11+cfg.NewMsgScanBatchSize)),
		).Return([]model.CCIPMsg{}, nil)

		// and there are two new message on chainB
		n.ccipReader.On(
			"MsgsBetweenSeqNums",
			ctx,
			chainB,
			model.NewSeqNumRange(21, model.SeqNum(21+cfg.NewMsgScanBatchSize)),
		).Return([]model.CCIPMsg{
			{CCIPMsgBaseDetails: model.CCIPMsgBaseDetails{ID: model.Bytes32{1}, SourceChain: chainB, SeqNum: 21}},
			{CCIPMsgBaseDetails: model.CCIPMsgBaseDetails{ID: model.Bytes32{2}, SourceChain: chainB, SeqNum: 22}},
		}, nil)

		n.ccipReader.On("GasPrices", ctx, []model.ChainSelector{chainA, chainB}).
			Return([]model.BigInt{
				{big.NewInt(1000)},
				{big.NewInt(20000)},
			}, nil)
	}

	return nodes
}

type nodeSetup struct {
	node        *Plugin
	ccipReader  *mocks.CCIPReader
	priceReader *mocks.TokenPricesReader
	reportCodec *mocks.CommitPluginJSONReportCodec
	msgHasher   *mocks.MessageHasher
}

func newNode(t *testing.T, ctx context.Context, lggr logger.Logger, id int, cfg model.CommitPluginConfig) nodeSetup {
	ccipReader := mocks.NewCCIPReader()
	priceReader := mocks.NewTokenPricesReader()
	reportCodec := mocks.NewCommitPluginJSONReportCodec()
	msgHasher := mocks.NewMessageHasher()

	node1 := NewPlugin(
		context.Background(),
		commontypes.OracleID(id),
		cfg,
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
