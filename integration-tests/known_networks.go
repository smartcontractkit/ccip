package networks

import (
	"os"
	"strings"
	"time"

	"github.com/smartcontractkit/chainlink-testing-framework/blockchain"

	"github.com/smartcontractkit/chainlink-env/logging"
)

func init() {
	logging.Init()
}

// Pre-configured test networks and their connections
// Some networks with public RPC endpoints are already filled out, but make use of environment variables to use info like
// private RPC endpoints and private keys.
var (
	// SimulatedEVM represents a simulated network
	SimulatedEVM *blockchain.EVMNetwork = blockchain.SimulatedEVMNetwork

	NetworkAlpha = &blockchain.EVMNetwork{
		Name:      "source-chain",
		Simulated: true,
		ChainID:   1337,
		PrivateKeys: []string{
			"ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
		},
		URLs:                      []string{"ws://source-chain-ethereum-geth:8546"},
		HTTPURLs:                  []string{"http://source-chain-ethereum-geth:8544"},
		ChainlinkTransactionLimit: 500000,
		Timeout:                   2 * time.Minute,
		MinimumConfirmations:      1,
		GasEstimationBuffer:       10000,
	}

	NetworkBeta = &blockchain.EVMNetwork{
		Name:      "dest-chain",
		Simulated: true,
		ChainID:   2337,
		PrivateKeys: []string{
			"ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
		},
		URLs:                      []string{"ws://dest-chain-ethereum-geth:8546"},
		HTTPURLs:                  []string{"http://dest-chain-ethereum-geth:8544"},
		ChainlinkTransactionLimit: 500000,
		Timeout:                   2 * time.Minute,
		MinimumConfirmations:      1,
		GasEstimationBuffer:       10000,
	}

	// SepoliaTestnet https://sepolia.dev/
	SepoliaTestnet *blockchain.EVMNetwork = &blockchain.EVMNetwork{
		Name:                      "Sepolia Testnet",
		ChainID:                   11155111,
		URLs:                      strings.Split(os.Getenv("EVM_URLS"), ","),
		Simulated:                 false,
		PrivateKeys:               strings.Split(os.Getenv("EVM_PRIVATE_KEYS"), ","),
		ChainlinkTransactionLimit: 5000,
		Timeout:                   time.Minute,
		MinimumConfirmations:      1,
		GasEstimationBuffer:       1000,
	}

	// GoerliTestnet https://goerli.net/
	GoerliTestnet *blockchain.EVMNetwork = &blockchain.EVMNetwork{
		Name:                      "Goerli Testnet",
		ChainID:                   5,
		URLs:                      strings.Split(os.Getenv("EVM_URLS"), ","),
		Simulated:                 false,
		PrivateKeys:               strings.Split(os.Getenv("EVM_PRIVATE_KEYS"), ","),
		ChainlinkTransactionLimit: 5000,
		Timeout:                   time.Minute * 5,
		MinimumConfirmations:      5,
		GasEstimationBuffer:       1000,
	}

	// KlaytnBaobab https://klaytn.foundation/
	KlaytnBaobab *blockchain.EVMNetwork = &blockchain.EVMNetwork{
		Name:                      "Klaytn Baobab",
		ChainID:                   1001,
		URLs:                      strings.Split(os.Getenv("EVM_URLS"), ","),
		Simulated:                 false,
		PrivateKeys:               strings.Split(os.Getenv("EVM_PRIVATE_KEYS"), ","),
		ChainlinkTransactionLimit: 5000,
		Timeout:                   time.Minute,
		MinimumConfirmations:      1,
		GasEstimationBuffer:       0,
	}

	// MetisStardust https://www.metis.io/
	MetisStardust *blockchain.EVMNetwork = &blockchain.EVMNetwork{
		Name:                      "Metis Stardust",
		ChainID:                   588,
		URLs:                      []string{"wss://stardust-ws.metis.io/"},
		Simulated:                 false,
		PrivateKeys:               strings.Split(os.Getenv("EVM_PRIVATE_KEYS"), ","),
		ChainlinkTransactionLimit: 5000,
		Timeout:                   time.Minute,
		MinimumConfirmations:      1,
		GasEstimationBuffer:       1000,
	}

	// ArbitrumGoerli https://developer.offchainlabs.com/docs/public_chains
	ArbitrumGoerli *blockchain.EVMNetwork = &blockchain.EVMNetwork{
		Name:                      "Arbitrum Goerli",
		ChainID:                   421613,
		URLs:                      strings.Split(os.Getenv("EVM_URLS"), ","),
		Simulated:                 false,
		PrivateKeys:               strings.Split(os.Getenv("EVM_PRIVATE_KEYS"), ","),
		ChainlinkTransactionLimit: 5000,
		Timeout:                   time.Minute,
		MinimumConfirmations:      0,
		GasEstimationBuffer:       0,
	}

	// OptimismGoerli https://dev.optimism.io/kovan-to-goerli/
	OptimismGoerli *blockchain.EVMNetwork = &blockchain.EVMNetwork{
		Name:                      "Optimism Goerli",
		ChainID:                   420,
		URLs:                      strings.Split(os.Getenv("EVM_URLS"), ","),
		Simulated:                 false,
		PrivateKeys:               strings.Split(os.Getenv("EVM_PRIVATE_KEYS"), ","),
		ChainlinkTransactionLimit: 5000,
		Timeout:                   time.Minute,
		MinimumConfirmations:      0,
		GasEstimationBuffer:       0,
	}
)

// GeneralEVM loads general EVM settings from env vars
func GeneralEVM() *blockchain.EVMNetwork {
	return blockchain.LoadNetworkFromEnvironment()
}
