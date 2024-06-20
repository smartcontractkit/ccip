package commit

import (
	"context"
	"reflect"
	"testing"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/stretchr/testify/mock"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/ccipocr3/internal/libs/testhelpers"
	"github.com/smartcontractkit/ccipocr3/internal/mocks"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

func TestPlugin(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Test(t)

	testCases := []struct {
		name                  string
		description           string
		nodes                 []nodeSetup
		expErr                func(*testing.T, error)
		expOutcome            cciptypes.CommitPluginOutcome
		expTransmittedReports []cciptypes.CommitPluginReport
		initialOutcome        cciptypes.CommitPluginOutcome
	}{
		{
			name:        "EmptyOutcome",
			description: "Empty observations are returned by all nodes which leads to an empty outcome.",
			nodes:       setupEmptyOutcome(ctx, t, lggr),
			expErr:      func(t *testing.T, err error) { assert.Equal(t, testhelpers.ErrEmptyOutcome, err) },
		},
		{
			name: "AllNodesReadAllChains",
			description: "Nodes observe the latest sequence numbers and new messages after those sequence numbers. " +
				"They also observe gas prices. In this setup all nodes can read all chains.",
			nodes: setupAllNodesReadAllChains(ctx, t, lggr),
			expOutcome: cciptypes.CommitPluginOutcome{
				MaxSeqNums: []cciptypes.SeqNumChain{
					{ChainSel: chainA, SeqNum: 10},
					{ChainSel: chainB, SeqNum: 20},
				},
				MerkleRoots: []cciptypes.MerkleRootChain{
					{ChainSel: chainB, MerkleRoot: cciptypes.Bytes32{}, SeqNumsRange: cciptypes.NewSeqNumRange(21, 22)},
				},
				TokenPrices: []cciptypes.TokenPrice{},
				GasPrices: []cciptypes.GasPriceChain{
					{ChainSel: chainA, GasPrice: cciptypes.NewBigIntFromInt64(1000)},
					{ChainSel: chainB, GasPrice: cciptypes.NewBigIntFromInt64(20_000)},
				},
			},
			expTransmittedReports: []cciptypes.CommitPluginReport{
				{
					MerkleRoots: []cciptypes.MerkleRootChain{
						{ChainSel: chainB, SeqNumsRange: cciptypes.NewSeqNumRange(21, 22)},
					},
					PriceUpdates: cciptypes.PriceUpdates{
						TokenPriceUpdates: []cciptypes.TokenPrice{},
						GasPriceUpdates: []cciptypes.GasPriceChain{
							{ChainSel: chainA, GasPrice: cciptypes.NewBigIntFromInt64(1000)},
							{ChainSel: chainB, GasPrice: cciptypes.NewBigIntFromInt64(20_000)},
						},
					},
				},
			},
			initialOutcome: cciptypes.CommitPluginOutcome{
				MaxSeqNums: []cciptypes.SeqNumChain{
					{ChainSel: chainA, SeqNum: 10},
					{ChainSel: chainB, SeqNum: 20},
				},
				MerkleRoots: []cciptypes.MerkleRootChain{},
				TokenPrices: []cciptypes.TokenPrice{},
				GasPrices:   []cciptypes.GasPriceChain{},
			},
		},
		{
			name:        "NodesDoNotAgreeOnMsgs",
			description: "Nodes do not agree on messages which leads to an outcome with empty merkle roots.",
			nodes:       setupNodesDoNotAgreeOnMsgs(ctx, t, lggr),
			expOutcome: cciptypes.CommitPluginOutcome{
				MaxSeqNums: []cciptypes.SeqNumChain{
					{ChainSel: chainA, SeqNum: 10},
					{ChainSel: chainB, SeqNum: 20},
				},
				MerkleRoots: []cciptypes.MerkleRootChain{},
				TokenPrices: []cciptypes.TokenPrice{},
				GasPrices: []cciptypes.GasPriceChain{
					{ChainSel: chainA, GasPrice: cciptypes.NewBigIntFromInt64(1000)},
					{ChainSel: chainB, GasPrice: cciptypes.NewBigIntFromInt64(20_000)},
				},
			},
			expTransmittedReports: []cciptypes.CommitPluginReport{
				{
					MerkleRoots: []cciptypes.MerkleRootChain{},
					PriceUpdates: cciptypes.PriceUpdates{
						TokenPriceUpdates: []cciptypes.TokenPrice{},
						GasPriceUpdates: []cciptypes.GasPriceChain{
							{ChainSel: chainA, GasPrice: cciptypes.NewBigIntFromInt64(1000)},
							{ChainSel: chainB, GasPrice: cciptypes.NewBigIntFromInt64(20_000)},
						},
					},
				},
			},
			initialOutcome: cciptypes.CommitPluginOutcome{
				MaxSeqNums: []cciptypes.SeqNumChain{
					{ChainSel: chainA, SeqNum: 10},
					{ChainSel: chainB, SeqNum: 20},
				},
				MerkleRoots: []cciptypes.MerkleRootChain{},
				TokenPrices: []cciptypes.TokenPrice{},
				GasPrices:   []cciptypes.GasPriceChain{},
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
			o, err := tc.initialOutcome.Encode()
			require.NoError(t, err)
			runner := testhelpers.NewOCR3Runner(nodes, nodeIDs, o)

			res, err := runner.RunRound(ctx)
			if tc.expErr != nil {
				tc.expErr(t, err)
			} else {
				assert.NoError(t, err)
			}

			if !reflect.DeepEqual(tc.expOutcome, cciptypes.CommitPluginOutcome{}) {
				outcome, err := cciptypes.DecodeCommitPluginOutcome(res.Outcome)
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

			assert.Equal(t, len(tc.expTransmittedReports), len(res.Transmitted))
			for i, exp := range tc.expTransmittedReports {
				actual, err := nodesSetup[0].reportCodec.Decode(ctx, res.Transmitted[i].Report)
				assert.NoError(t, err)
				assert.Equal(t, exp.PriceUpdates, actual.PriceUpdates)
				assert.Equal(t, len(exp.MerkleRoots), len(actual.MerkleRoots))
				for j, expRoot := range exp.MerkleRoots {
					assert.Equal(t, expRoot.ChainSel, actual.MerkleRoots[j].ChainSel)
					assert.Equal(t, expRoot.SeqNumsRange, actual.MerkleRoots[j].SeqNumsRange)
				}
			}
		})
	}
}

func setupEmptyOutcome(ctx context.Context, t *testing.T, lggr logger.Logger) []nodeSetup {
	cfg := cciptypes.CommitPluginConfig{
		DestChain:           chainC,
		PricedTokens:        []types.Account{tokenX},
		TokenPricesObserver: false,
		NewMsgScanBatchSize: 256,
	}

	homeChainConfig := cciptypes.HomeChainConfig{
		FChain: map[cciptypes.ChainSelector]int{chainC: 1},
		NodeSupportedChains: map[cciptypes.P2PID]cciptypes.SupportedChains{
			cciptypes.GetP2pID(1): {Supported: mapset.NewSet[cciptypes.ChainSelector](chainC)},
			cciptypes.GetP2pID(2): {Supported: mapset.NewSet[cciptypes.ChainSelector](chainC)},
			cciptypes.GetP2pID(3): {Supported: mapset.NewSet[cciptypes.ChainSelector](chainC)},
		},
	}

	oracleIDToP2pID := GetP2pIDs(1, 2, 3)
	return []nodeSetup{
		newNode(ctx, t, lggr, 1, cfg, homeChainConfig, oracleIDToP2pID),
		newNode(ctx, t, lggr, 2, cfg, homeChainConfig, oracleIDToP2pID),
		newNode(ctx, t, lggr, 3, cfg, homeChainConfig, oracleIDToP2pID),
	}
}

func setupAllNodesReadAllChains(ctx context.Context, t *testing.T, lggr logger.Logger) []nodeSetup {
	cfg := cciptypes.CommitPluginConfig{
		DestChain:           chainC,
		PricedTokens:        []types.Account{tokenX},
		TokenPricesObserver: false,
		NewMsgScanBatchSize: 256,
	}

	homeChainConfig := cciptypes.HomeChainConfig{
		FChain: map[cciptypes.ChainSelector]int{
			chainA: 1,
			chainB: 1,
			chainC: 1,
		},
		NodeSupportedChains: map[cciptypes.P2PID]cciptypes.SupportedChains{
			cciptypes.GetP2pID(1): {Supported: mapset.NewSet[cciptypes.ChainSelector](chainA, chainB, chainC)},
			cciptypes.GetP2pID(2): {Supported: mapset.NewSet[cciptypes.ChainSelector](chainA, chainB, chainC)},
			cciptypes.GetP2pID(3): {Supported: mapset.NewSet[cciptypes.ChainSelector](chainA, chainB, chainC)},
		},
	}
	oracleIDToP2pID := GetP2pIDs(1, 2, 3)
	n1 := newNode(ctx, t, lggr, 1, cfg, homeChainConfig, oracleIDToP2pID)
	n2 := newNode(ctx, t, lggr, 2, cfg, homeChainConfig, oracleIDToP2pID)
	n3 := newNode(ctx, t, lggr, 3, cfg, homeChainConfig, oracleIDToP2pID)
	nodes := []nodeSetup{n1, n2, n3}

	for _, n := range nodes {
		// then they fetch new msgs, there is nothing new on chainA
		n.ccipReader.On(
			"MsgsBetweenSeqNums",
			ctx,
			chainA,
			cciptypes.NewSeqNumRange(11, cciptypes.SeqNum(11+cfg.NewMsgScanBatchSize)),
		).Return([]cciptypes.CCIPMsg{}, nil)

		// and there are two new message on chainB
		n.ccipReader.On(
			"MsgsBetweenSeqNums",
			ctx,
			chainB,
			cciptypes.NewSeqNumRange(21, cciptypes.SeqNum(21+cfg.NewMsgScanBatchSize)),
		).Return([]cciptypes.CCIPMsg{
			{CCIPMsgBaseDetails: cciptypes.CCIPMsgBaseDetails{ID: cciptypes.Bytes32{1}, SourceChain: chainB, SeqNum: 21}},
			{CCIPMsgBaseDetails: cciptypes.CCIPMsgBaseDetails{ID: cciptypes.Bytes32{2}, SourceChain: chainB, SeqNum: 22}},
		}, nil)

		n.ccipReader.On("GasPrices", ctx, []cciptypes.ChainSelector{chainA, chainB}).
			Return([]cciptypes.BigInt{
				cciptypes.NewBigIntFromInt64(1000),
				cciptypes.NewBigIntFromInt64(20_000),
			}, nil)

		// all nodes observe the same sequence numbers 10 for chainA and 20 for chainB
		n.ccipReader.On("NextSeqNum", ctx, []cciptypes.ChainSelector{chainA, chainB}).
			Return([]cciptypes.SeqNum{10, 20}, nil)

	}

	return nodes
}

func setupNodesDoNotAgreeOnMsgs(ctx context.Context, t *testing.T, lggr logger.Logger) []nodeSetup {
	cfg := cciptypes.CommitPluginConfig{
		DestChain:           chainC,
		PricedTokens:        []types.Account{tokenX},
		TokenPricesObserver: false,
		NewMsgScanBatchSize: 256,
	}

	homeChainConfig := cciptypes.HomeChainConfig{
		FChain: map[cciptypes.ChainSelector]int{
			chainA: 1,
			chainB: 1,
			chainC: 1,
		},
		NodeSupportedChains: map[cciptypes.P2PID]cciptypes.SupportedChains{
			cciptypes.GetP2pID(1): {Supported: mapset.NewSet[cciptypes.ChainSelector](chainA, chainB, chainC)},
			cciptypes.GetP2pID(2): {Supported: mapset.NewSet[cciptypes.ChainSelector](chainA, chainB, chainC)},
			cciptypes.GetP2pID(3): {Supported: mapset.NewSet[cciptypes.ChainSelector](chainA, chainB, chainC)},
		},
	}

	oracleIDToP2pID := GetP2pIDs(1, 2, 3)
	n1 := newNode(ctx, t, lggr, 1, cfg, homeChainConfig, oracleIDToP2pID)
	n2 := newNode(ctx, t, lggr, 2, cfg, homeChainConfig, oracleIDToP2pID)
	n3 := newNode(ctx, t, lggr, 3, cfg, homeChainConfig, oracleIDToP2pID)
	nodes := []nodeSetup{n1, n2, n3}

	for i, n := range nodes {
		// all nodes observe the same sequence numbers 10 for chainA and 20 for chainB
		n.ccipReader.On("NextSeqNum", ctx, []cciptypes.ChainSelector{chainA, chainB}).
			Return([]cciptypes.SeqNum{10, 20}, nil)

		// then they fetch new msgs, there is nothing new on chainA
		n.ccipReader.On(
			"MsgsBetweenSeqNums",
			ctx,
			chainA,
			cciptypes.NewSeqNumRange(11, cciptypes.SeqNum(11+cfg.NewMsgScanBatchSize)),
		).Return([]cciptypes.CCIPMsg{}, nil)

		// and there are two new message on chainB
		n.ccipReader.On(
			"MsgsBetweenSeqNums",
			ctx,
			chainB,
			cciptypes.NewSeqNumRange(
				21,
				cciptypes.SeqNum(21+cfg.NewMsgScanBatchSize),
			),
		).Return([]cciptypes.CCIPMsg{
			{CCIPMsgBaseDetails: cciptypes.CCIPMsgBaseDetails{ID: cciptypes.Bytes32{1, byte(i)}, SourceChain: chainB, SeqNum: 21 + cciptypes.SeqNum(i*10)}},
			{CCIPMsgBaseDetails: cciptypes.CCIPMsgBaseDetails{ID: cciptypes.Bytes32{2, byte(i)}, SourceChain: chainB, SeqNum: 22 + cciptypes.SeqNum(i*20)}},
		}, nil)

		n.ccipReader.On("GasPrices", ctx, []cciptypes.ChainSelector{chainA, chainB}).
			Return([]cciptypes.BigInt{
				cciptypes.NewBigIntFromInt64(1000),
				cciptypes.NewBigIntFromInt64(20_000),
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

func newNode(ctx context.Context, t *testing.T, lggr logger.Logger, id int, cfg cciptypes.CommitPluginConfig, homeChainConfig cciptypes.HomeChainConfig, oracleIDToP2pID map[commontypes.OracleID]cciptypes.P2PID) nodeSetup {
	ccipReader := mocks.NewCCIPReader()
	priceReader := mocks.NewTokenPricesReader()
	reportCodec := mocks.NewCommitPluginJSONReportCodec()
	msgHasher := mocks.NewMessageHasher()
	homeChainPoller := mocks.NewHomeChainPollerMock()

	homeChainPoller.On("Start", mock.Anything).Return(nil)
	homeChainPoller.On("GetConfig").Return(homeChainConfig)

	node1 := NewPlugin(
		context.Background(),
		commontypes.OracleID(id),
		oracleIDToP2pID,
		cfg,
		ccipReader,
		priceReader,
		reportCodec,
		msgHasher,
		lggr,
		homeChainPoller,
	)

	return nodeSetup{
		node:        node1,
		ccipReader:  ccipReader,
		priceReader: priceReader,
		reportCodec: reportCodec,
		msgHasher:   msgHasher,
	}
}
func GetP2pIDs(ids ...int) map[commontypes.OracleID]cciptypes.P2PID {
	res := make(map[commontypes.OracleID]cciptypes.P2PID)
	for _, id := range ids {
		res[commontypes.OracleID(id)] = cciptypes.GetP2pID(id)
	}
	return res
}

var (
	chainA = cciptypes.ChainSelector(1)
	chainB = cciptypes.ChainSelector(2)
	chainC = cciptypes.ChainSelector(3)

	tokenX = types.Account("tk_xxx")
)
