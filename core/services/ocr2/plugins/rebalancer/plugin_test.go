package rebalancer

import (
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2/types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap/zapcore"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	ubig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	bridgemocks "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/bridge/mocks"
	discoverermocks "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/discoverer/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/graph"
	liquiditymanagermocks "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/liquiditymanager/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/liquidityrebalancer"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
	rebalancermocks "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/rebalancermocks"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/testhelpers"
)

func TestPlugin_E2EWithMocks(t *testing.T) {
	ctx := testutils.Context(t)
	lggr := logger.TestLogger(t)
	lggr.SetLogLevel(zapcore.ErrorLevel)

	testCases := []testCase{
		twoNodesFourRounds(t),
	}

	for _, tc := range testCases {
		tc.validate(t)
		t.Run(tc.name, func(t *testing.T) {
			// init the nodes and the ocr3 runner
			nodes := make([]node, tc.numNodes)
			plugins := make([]ocr3types.ReportingPlugin[models.Report], tc.numNodes)
			for i := range nodes {
				nodes[i] = newNode(t, lggr, tc.f)
				plugins[i] = nodes[i].plugin
			}
			ocr3Runner := testhelpers.NewOCR3Runner[models.Report](plugins)

			for numRound, round := range tc.rounds {
				for i, n := range nodes {
					t.Logf(">>> running round: %d", numRound)
					// the node will first discover the graph, let's mock the observed graph
					discoverer := discoverermocks.NewDiscoverer(t)
					n.discovererFactory.
						On("NewDiscoverer", n.plugin.rootNetwork, n.plugin.rootAddress).
						Return(discoverer, nil).Maybe()
					discoverer.
						On("Discover", mock.Anything).
						Return(round.discoveredGraphPerNode[i](), nil).Maybe()

					// the node will now try to load the pending transfers of all the available bridges
					// let's mock the pending transfers
					observedGraph := round.discoveredGraphPerNode[i]()
					edges, err := observedGraph.GetEdges()
					assert.NoError(t, err)
					for _, edge := range edges {
						bridge, ok := n.bridges[[2]models.NetworkSelector{edge.Source, edge.Dest}]
						assert.True(t, ok, "the test case is wrong, bridge is not defined %d->%d", edge.Source, edge.Dest)
						n.bridgeFactory.On("NewBridge", edge.Source, edge.Dest).Return(bridge, nil).Maybe()

						pendingTransfers := make([]models.PendingTransfer, 0)
						for _, tr := range round.pendingTransfersPerNode[i] {
							if tr.From == edge.Source && tr.To == edge.Dest {
								pendingTransfers = append(pendingTransfers, tr)
							}
						}
						bridge.
							On("GetTransfers", mock.Anything, mock.Anything, mock.Anything).
							Return(pendingTransfers, nil).Maybe()

						bridge.
							On("GetBridgePayloadAndFee", mock.Anything, mock.Anything).
							Return(nil, big.NewInt(10), nil).Maybe()

						bridge.
							On("QuorumizedBridgePayload", mock.Anything, mock.Anything).
							Return(nil, nil).Maybe()
					}

					for net, seqNum := range round.seqNumPerRebalancer {
						rb, exists := n.rebalancers[net]
						assert.True(t, exists, "test case is wrong, seq num of rebalancer is not defined")
						rb.On("GetLatestSequenceNumber", mock.Anything).Return(seqNum, nil).Maybe()
						n.rbFactory.On("NewRebalancer", net, mock.Anything).Return(rb, nil).Maybe()
					}
				}

				transmitted, notAccepted, notTransmitted, outcome, err := ocr3Runner.RunRound(ctx)
				if round.expErr {
					assert.Error(t, err)
					continue
				}
				assertOutcomeEqual(t, round.expOutcome, outcome)
				assertReportsSlicesEqual(t, round.expTransmitted, transmitted)
				assertReportsSlicesEqual(t, round.expNotAccepted, notAccepted)
				assertReportsSlicesEqual(t, round.expNotTransmitted, notTransmitted)
			}
		})
	}
}

func twoNodesFourRounds(t *testing.T) testCase {
	g := graph.NewGraph()
	g.AddNetwork(networkA, graph.Data{
		Liquidity:         big.NewInt(1000),
		TokenAddress:      tokenX,
		RebalancerAddress: rebalancerA,
		XChainRebalancers: nil,
		NetworkSelector:   networkA,
		ConfigDigest:      cfgDigest1,
	})
	g.AddNetwork(networkB, graph.Data{
		Liquidity:         big.NewInt(2000),
		TokenAddress:      tokenX,
		RebalancerAddress: rebalancerB,
		XChainRebalancers: nil,
		NetworkSelector:   networkB,
		ConfigDigest:      cfgDigest2,
	})
	assert.NoError(t, g.AddConnection(networkA, networkB))
	assert.NoError(t, g.AddConnection(networkB, networkA))

	return testCase{
		name:     "two nodes four rounds nothing inflight",
		numNodes: 2,
		f:        1,
		rounds: []roundData{
			{ // round 1 - new transfers to reach balance in the outcome
				discoveredGraphPerNode: []func() graph.Graph{
					func() graph.Graph { return g },
					func() graph.Graph { return g },
				},
				pendingTransfersPerNode: [][]models.PendingTransfer{{}, {}},
				expTransmitted:          []ocr3types.ReportWithInfo[models.Report]{},
				expNotTransmitted:       []ocr3types.ReportWithInfo[models.Report]{},
				expNotAccepted:          []ocr3types.ReportWithInfo[models.Report]{},
				expOutcome: models.NewOutcome(
					[]models.ProposedTransfer{
						{From: networkA, To: networkB, Amount: ubig.New(big.NewInt(1000))},
					},
					nil,
					nil,
					[]models.ConfigDigestWithMeta{{Digest: cfgDigest1, NetworkSel: networkA}, {Digest: cfgDigest2, NetworkSel: networkB}}),
				seqNumPerRebalancer: map[models.NetworkSelector]uint64{
					networkA: 1,
					networkB: 2,
				},
			},
			{ // round 2 - the transfers of the previous outcome are included in the report
				discoveredGraphPerNode: []func() graph.Graph{
					func() graph.Graph { return g },
					func() graph.Graph { return g },
				},
				pendingTransfersPerNode: [][]models.PendingTransfer{{}, {}},
				expTransmitted: []ocr3types.ReportWithInfo[models.Report]{
					{
						Info: models.Report{
							Transfers:               []models.Transfer{models.NewTransfer(networkA, networkB, big.NewInt(1000), time.Time{}, nil)},
							LiquidityManagerAddress: rebalancerA,
							NetworkID:               networkA,
							ConfigDigest:            cfgDigest1,
						},
					},
				},
				expNotTransmitted: []ocr3types.ReportWithInfo[models.Report]{},
				expNotAccepted:    []ocr3types.ReportWithInfo[models.Report]{},
				expOutcome: models.NewOutcome(
					nil,
					[]models.Transfer{{From: networkA, To: networkB, Amount: ubig.New(big.NewInt(1000))}},
					nil,
					[]models.ConfigDigestWithMeta{{Digest: cfgDigest1, NetworkSel: networkA}, {Digest: cfgDigest2, NetworkSel: networkB}}),
				seqNumPerRebalancer: map[models.NetworkSelector]uint64{
					networkA: 1,
					networkB: 2,
				},
			},
			{ // round 3 - nothing new
				discoveredGraphPerNode: []func() graph.Graph{
					func() graph.Graph { return g },
					func() graph.Graph { return g },
				},
				pendingTransfersPerNode: [][]models.PendingTransfer{{}, {}},
				expTransmitted:          []ocr3types.ReportWithInfo[models.Report]{},
				expNotTransmitted:       []ocr3types.ReportWithInfo[models.Report]{},
				expNotAccepted:          []ocr3types.ReportWithInfo[models.Report]{},
				expOutcome: models.NewOutcome(
					[]models.ProposedTransfer{
						// TODO: this slice should be empty for this test to pass.
						// right now the plugin will propose the same transfer again because it's missing an inflight cache.
						{From: networkA, To: networkB, Amount: ubig.New(big.NewInt(1000))},
					},
					nil,
					nil,
					[]models.ConfigDigestWithMeta{{Digest: cfgDigest1, NetworkSel: networkA}, {Digest: cfgDigest2, NetworkSel: networkB}}),
				seqNumPerRebalancer: map[models.NetworkSelector]uint64{
					networkA: 2,
					networkB: 3,
				},
			},
		},
	}
}

func assertReportsSlicesEqual(t *testing.T, r1, r2 []ocr3types.ReportWithInfo[models.Report]) {
	assert.Equal(t, len(r1), len(r2))
	for i := range r1 {
		assertReportsEqual(t, r1[i], r2[i])
	}
}

func assertReportsEqual(t *testing.T, r1, r2 ocr3types.ReportWithInfo[models.Report]) {
	assertTransfersEqual(t, r1.Info.Transfers, r2.Info.Transfers)
	assert.Equal(t, r1.Info.NetworkID, r2.Info.NetworkID)
	assert.Equal(t, r1.Info.LiquidityManagerAddress, r2.Info.LiquidityManagerAddress)
	assert.Equal(t, r1.Info.ConfigDigest.Hex(), r2.Info.ConfigDigest.Hex())
}

func assertTransfersEqual(t *testing.T, a, b []models.Transfer) {
	assert.Equal(t, len(a), len(b))
	for i := range a {
		assert.Equal(t, a[i].From, b[i].From)
		assert.Equal(t, a[i].To, b[i].To)
		assert.Equal(t, a[i].Amount, b[i].Amount)
	}
}

func assertPendingTransfersEqual(t *testing.T, a, b []models.PendingTransfer) {
	assert.Equal(t, len(a), len(b))
	for i := range a {
		assert.Equal(t, a[i].From, b[i].From)
		assert.Equal(t, a[i].To, b[i].To)
		assert.Equal(t, a[i].Amount, b[i].Amount)
	}
}

func assertProposedTransfersEqual(t *testing.T, a, b []models.ProposedTransfer) {
	assert.Equal(t, len(a), len(b))
	for i := range a {
		assert.Equal(t, a[i].From, b[i].From)
		assert.Equal(t, a[i].To, b[i].To)
		assert.Equal(t, a[i].Amount, b[i].Amount)
	}
}

func assertOutcomeEqual(t *testing.T, exp models.Outcome, got []byte) {
	decodedOutcome := models.Outcome{}
	err := json.Unmarshal(got, &decodedOutcome)
	assert.NoError(t, err)

	assert.Equal(t, exp.ConfigDigests, decodedOutcome.ConfigDigests)
	assertTransfersEqual(t, exp.ResolvedTransfers, decodedOutcome.ResolvedTransfers)
	assertPendingTransfersEqual(t, exp.PendingTransfers, decodedOutcome.PendingTransfers)
	assertProposedTransfersEqual(t, exp.ProposedTransfers, decodedOutcome.ProposedTransfers)
}

type testCase struct {
	name     string
	numNodes int
	f        int
	rounds   []roundData
}

func (tc *testCase) validate(t *testing.T) {
	assert.Positive(t, len(tc.rounds))
	assert.Positive(t, tc.numNodes)
	assert.NotEmpty(t, tc.name)

	for _, r := range tc.rounds {
		assert.Equal(t, len(r.discoveredGraphPerNode), tc.numNodes, "you should define discovered graph per node")
		assert.Equal(t, len(r.pendingTransfersPerNode), tc.numNodes, "you should define pending transfers per node")
		assert.Positive(t, len(r.seqNumPerRebalancer), "you should define the seq nums of the rebalancers")
	}
}

type roundData struct {
	discoveredGraphPerNode  []func() graph.Graph
	pendingTransfersPerNode [][]models.PendingTransfer
	seqNumPerRebalancer     map[models.NetworkSelector]uint64
	expOutcome              models.Outcome

	expTransmitted    []ocr3types.ReportWithInfo[models.Report]
	expNotAccepted    []ocr3types.ReportWithInfo[models.Report]
	expNotTransmitted []ocr3types.ReportWithInfo[models.Report]
	expErr            bool
}

type node struct {
	plugin            *Plugin
	rbFactory         *rebalancermocks.Factory
	discovererFactory *discoverermocks.Factory
	bridgeFactory     *bridgemocks.Factory
	rebalancers       map[models.NetworkSelector]*liquiditymanagermocks.Rebalancer
	bridges           map[[2]models.NetworkSelector]*bridgemocks.Bridge
}

func newNode(t *testing.T, lggr logger.Logger, f int) node {
	lmFactory := rebalancermocks.NewFactory(t)
	discovererFactory := discoverermocks.NewFactory(t)
	bridgeFactory := bridgemocks.NewFactory(t)
	rebalancerAlg := liquidityrebalancer.NewPingPong()

	node1 := NewPlugin(
		f,
		time.Minute,
		networkA,
		models.Address(utils.RandomAddress()),
		lmFactory,
		discovererFactory,
		bridgeFactory,
		rebalancerAlg,
		lggr,
	)

	bridgeMocks := make(map[[2]models.NetworkSelector]*bridgemocks.Bridge)
	for _, b := range bridges {
		bridgeMocks[b] = bridgemocks.NewBridge(t)
	}

	return node{
		plugin:            node1,
		rbFactory:         lmFactory,
		discovererFactory: discovererFactory,
		bridgeFactory:     bridgeFactory,
		bridges:           bridgeMocks,
		rebalancers: map[models.NetworkSelector]*liquiditymanagermocks.Rebalancer{
			networkA: liquiditymanagermocks.NewRebalancer(t),
			networkB: liquiditymanagermocks.NewRebalancer(t),
			networkC: liquiditymanagermocks.NewRebalancer(t),
			networkD: liquiditymanagermocks.NewRebalancer(t), // todo: loop
		},
	}
}

// test helper variables

var (
	networkA = models.NetworkSelector(1)
	networkB = models.NetworkSelector(2)
	networkC = models.NetworkSelector(3)
	networkD = models.NetworkSelector(4)

	rebalancerA = models.Address(common.HexToAddress("0xa"))
	rebalancerB = models.Address(common.HexToAddress("0xb"))
	rebalancerC = models.Address(common.HexToAddress("0xc"))

	tokenX = models.Address(common.HexToAddress("0x1"))
	tokenY = models.Address(common.HexToAddress("0x2"))
	tokenZ = models.Address(common.HexToAddress("0x3"))

	bridgeAB = [2]models.NetworkSelector{networkA, networkB}
	bridgeBA = [2]models.NetworkSelector{networkB, networkA}

	bridges = [][2]models.NetworkSelector{
		bridgeAB,
		bridgeBA,
	}

	cfgDigest1 = models.ConfigDigest{ConfigDigest: ocrtypes.ConfigDigest([32]byte{1})}
	cfgDigest2 = models.ConfigDigest{ConfigDigest: ocrtypes.ConfigDigest([32]byte{2})}
	cfgDigest3 = models.ConfigDigest{ConfigDigest: ocrtypes.ConfigDigest([32]byte{3})}
)
