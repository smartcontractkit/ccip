package load

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"

	networks "github.com/smartcontractkit/chainlink/integration-tests"
)

var _ = Describe("CCIP load test - parallel send between same sender and receiver address", func() {
	var (
		diffNetworkEntries = []TableEntry{
			Entry("CCIP suite on 2 Geths @simulated", networks.NetworkAlpha, networks.NetworkBeta),
		}

		testArgs *loadArgs
	)

	AfterEach(func() {
		testArgs.TearDown()
	})

	DescribeTable("CCIP suite on different EVM networks", func(
		sourceNetwork *blockchain.EVMNetwork,
		destNetwork *blockchain.EVMNetwork,
	) {
		By("Initialize ccip load test parameters")
		testArgs = PopulateAndValidate()

		By("Deploy and set up default load test environment")
		testArgs.Setup(sourceNetwork, destNetwork)

		By("Run the load sequence")
		testArgs.Run()
	},
		diffNetworkEntries,
	)
})
