package rebalancer

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap/zapcore"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
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
		//twoNodesOneRoundNothingInflight(t),
		twoNodesTwoRoundsNothingInflight(t),
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

			for _, round := range tc.rounds {
				for i, n := range nodes {
					discoverer := discoverermocks.NewDiscoverer(t)
					n.discovererFactory.On("NewDiscoverer", n.plugin.rootNetwork, n.plugin.rootAddress).Return(discoverer, nil).Maybe()
					discoverer.On("Discover", mock.Anything).Return(round.discoveredGraphPerNode[i](), nil).Maybe()

					// for each neighbor of this node get the bridge
					observedGraph := round.discoveredGraphPerNode[i]()
					edges, err := observedGraph.GetEdges()
					assert.NoError(t, err)
					for _, edge := range edges {
						bridge, ok := n.bridges[[2]models.NetworkSelector{edge.Source, edge.Dest}]
						assert.True(t, ok, "the test case is wrong, bridge is not defined %d->%d", edge.Source, edge.Dest)
						n.bridgeFactory.On("NewBridge", edge.Source, edge.Dest).Return(bridge, nil).Maybe()
						bridge.On("GetBridgePayloadAndFee", mock.Anything, mock.Anything).Return(nil, big.NewInt(10), nil).Maybe()
						// use round.pendingTransfersPerNode[i] to get bridge inflight transfers and statuses
						bridge.On("GetTransfers", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil).Maybe()
						bridge.On("QuorumizedBridgePayload", mock.Anything, mock.Anything).Return(nil, nil).Maybe()
					}

					// todo
					rbA := n.rebalancers[networkA]
					rbA.On("GetLatestSequenceNumber", mock.Anything).Return(uint64(0), nil).Maybe()
					rbB := n.rebalancers[networkA]
					rbB.On("GetLatestSequenceNumber", mock.Anything).Return(uint64(0), nil).Maybe()

					n.rbFactory.On("NewRebalancer", networkA, mock.Anything).Return(rbA, nil).Maybe()
					n.rbFactory.On("NewRebalancer", networkB, mock.Anything).Return(rbB, nil).Maybe()

				}

				transmitted, notAccepted, notTransmitted, err := ocr3Runner.RunRound(ctx)
				if round.expErr {
					assert.Error(t, err)
					continue
				}
				assert.NoError(t, err)
				t.Log(transmitted, notAccepted, notTransmitted)
			}
		})
	}
}

func twoNodesOneRoundNothingInflight(t *testing.T) testCase {
	return testCase{
		name:     "two nodes one round nothing inflight",
		numNodes: 2,
		f:        1,
		rounds: []roundData{
			{
				discoveredGraphPerNode: []func() graph.Graph{
					func() graph.Graph {
						g := graph.NewGraph()
						g.AddNetwork(networkA, graph.Data{
							Liquidity:         big.NewInt(1000),
							TokenAddress:      tokenX,
							RebalancerAddress: rebalancerA,
							XChainRebalancers: nil,
							NetworkSelector:   networkA,
						})
						g.AddNetwork(networkB, graph.Data{
							Liquidity:         big.NewInt(2000),
							TokenAddress:      tokenX,
							RebalancerAddress: rebalancerB,
							XChainRebalancers: nil,
							NetworkSelector:   networkB,
						})
						assert.NoError(t, g.AddConnection(networkA, networkB))
						assert.NoError(t, g.AddConnection(networkB, networkA))
						return g
					},
					func() graph.Graph {
						g := graph.NewGraph()
						g.AddNetwork(networkA, graph.Data{
							Liquidity:         big.NewInt(1000),
							TokenAddress:      tokenX,
							RebalancerAddress: rebalancerA,
							XChainRebalancers: nil,
							NetworkSelector:   networkA,
						})
						g.AddNetwork(networkB, graph.Data{
							Liquidity:         big.NewInt(2000),
							TokenAddress:      tokenX,
							RebalancerAddress: rebalancerB,
							XChainRebalancers: nil,
							NetworkSelector:   networkB,
						})
						assert.NoError(t, g.AddConnection(networkA, networkB))
						assert.NoError(t, g.AddConnection(networkB, networkA))
						return g
					},
				},
				pendingTransfersPerNode: [][]models.PendingTransfer{
					{},
					{},
				},
				expTransmitted:    "[]",
				expNotTransmitted: "[]",
				expNotAccepted:    "[]",
			},
		},
	}
}

func twoNodesTwoRoundsNothingInflight(t *testing.T) testCase {
	return testCase{
		name:     "two nodes one round nothing inflight",
		numNodes: 2,
		f:        1,
		rounds: []roundData{
			{ // round 1
				discoveredGraphPerNode: []func() graph.Graph{
					func() graph.Graph {
						g := graph.NewGraph()
						g.AddNetwork(networkA, graph.Data{
							Liquidity:         big.NewInt(1000),
							TokenAddress:      tokenX,
							RebalancerAddress: rebalancerA,
							XChainRebalancers: nil,
							NetworkSelector:   networkA,
						})
						g.AddNetwork(networkB, graph.Data{
							Liquidity:         big.NewInt(2000),
							TokenAddress:      tokenX,
							RebalancerAddress: rebalancerB,
							XChainRebalancers: nil,
							NetworkSelector:   networkB,
						})
						assert.NoError(t, g.AddConnection(networkA, networkB))
						assert.NoError(t, g.AddConnection(networkB, networkA))
						return g
					},
					func() graph.Graph {
						g := graph.NewGraph()
						g.AddNetwork(networkA, graph.Data{
							Liquidity:         big.NewInt(1000),
							TokenAddress:      tokenX,
							RebalancerAddress: rebalancerA,
							XChainRebalancers: nil,
							NetworkSelector:   networkA,
						})
						g.AddNetwork(networkB, graph.Data{
							Liquidity:         big.NewInt(2000),
							TokenAddress:      tokenX,
							RebalancerAddress: rebalancerB,
							XChainRebalancers: nil,
							NetworkSelector:   networkB,
						})
						assert.NoError(t, g.AddConnection(networkA, networkB))
						assert.NoError(t, g.AddConnection(networkB, networkA))
						return g
					},
				},
				pendingTransfersPerNode: [][]models.PendingTransfer{
					{},
					{},
				},
				expTransmitted:    "[]",
				expNotTransmitted: "[]",
				expNotAccepted:    "[]",
			},
			{ // round 2
				discoveredGraphPerNode: []func() graph.Graph{
					func() graph.Graph {
						g := graph.NewGraph()
						g.AddNetwork(networkA, graph.Data{
							Liquidity:         big.NewInt(1000),
							TokenAddress:      tokenX,
							RebalancerAddress: rebalancerA,
							XChainRebalancers: nil,
							NetworkSelector:   networkA,
						})
						g.AddNetwork(networkB, graph.Data{
							Liquidity:         big.NewInt(2000),
							TokenAddress:      tokenX,
							RebalancerAddress: rebalancerB,
							XChainRebalancers: nil,
							NetworkSelector:   networkB,
						})
						assert.NoError(t, g.AddConnection(networkA, networkB))
						assert.NoError(t, g.AddConnection(networkB, networkA))
						return g
					},
					func() graph.Graph {
						g := graph.NewGraph()
						g.AddNetwork(networkA, graph.Data{
							Liquidity:         big.NewInt(1000),
							TokenAddress:      tokenX,
							RebalancerAddress: rebalancerA,
							XChainRebalancers: nil,
							NetworkSelector:   networkA,
						})
						g.AddNetwork(networkB, graph.Data{
							Liquidity:         big.NewInt(2000),
							TokenAddress:      tokenX,
							RebalancerAddress: rebalancerB,
							XChainRebalancers: nil,
							NetworkSelector:   networkB,
						})
						assert.NoError(t, g.AddConnection(networkA, networkB))
						assert.NoError(t, g.AddConnection(networkB, networkA))
						return g
					},
				},
				pendingTransfersPerNode: [][]models.PendingTransfer{
					{},
					{},
				},
				expTransmitted:    "[]",
				expNotTransmitted: "[]",
				expNotAccepted:    "[]",
			},
		},
	}
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
	}
}

type roundData struct {
	discoveredGraphPerNode  []func() graph.Graph
	pendingTransfersPerNode [][]models.PendingTransfer

	expTransmitted    string
	expNotAccepted    string
	expNotTransmitted string
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
		models.NetworkSelector(1),
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
)
